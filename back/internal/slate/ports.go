package slate

type Repository interface {
	FindAll() ([]*Slate, error)
}

type Service struct {
	repository Repository
}

func NewService(repository Repository) *Service {
	return &Service{
		repository: repository,
	}
}
