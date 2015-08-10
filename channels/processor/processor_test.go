package processor_test

import (
	"github.com/myshkin5/writersandchannels/channels/processor"

	"math/rand"
	"sync/atomic"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Processor", func() {
	var (
		input chan interface{}
		proc  *processor.Processor
	)

	BeforeEach(func() {
		input = make(chan interface{})

		proc = processor.New(input)
	})

	It("returns from processing when its input channel closes", func() {
		var stopTime int64
		go func() {
			proc.Process()
			atomic.StoreInt64(&stopTime, time.Now().UnixNano())
		}()

		random := rand.New(rand.NewSource(time.Now().Unix()))

		// Sleep a small amount of time, close the channel then verify stopTime is set shortly thereafter
		time.Sleep(time.Millisecond * time.Duration(random.Intn(100)))

		close(input)

		Eventually(func() int64 { return atomic.LoadInt64(&stopTime) }).Should(BeNumerically(">", 0))
		Expect(time.Now().UnixNano() - atomic.LoadInt64(&stopTime)).To(BeNumerically("<", 20*time.Millisecond))
	})

	It("closes the output chan before processing completes", func() {
		close(input)

		proc.Process()

		select {
		case _, ok := <-proc.Output():
			Expect(ok).To(BeFalse())
		default:
			Fail("output chan was not closed")
		}
	})

	It("sends items that it receives on the input to its output", func() {

	})
})
