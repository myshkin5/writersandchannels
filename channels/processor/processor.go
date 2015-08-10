package processor

type Processor struct {
	input  <-chan interface{}
	output chan interface{}
}

func New(input <-chan interface{}) *Processor {
	return &Processor{
		input:  input,
		output: make(chan interface{}),
	}
}

func (p *Processor) Process() {
	for _ = range p.input {
	}
	close(p.output)
}

func (p *Processor) Output() <-chan interface{} {
	return p.output
}
