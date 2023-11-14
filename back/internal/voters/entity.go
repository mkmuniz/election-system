package voters

type Voter struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Cpf   string `json:"cpf"`
	State string `json:"state"`
}
