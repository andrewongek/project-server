package server

type Server interface {
	Start() error
	Stop()
}

type server struct {
	// Kafka Conusmer
}

func NewServer() (Server, error) {
	return &server{}, nil 
}

func (s *server) Start() error {
	return nil
}

func (s *server) Stop() {

}


