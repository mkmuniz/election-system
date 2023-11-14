package slate

type Slate struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Votes uint   `json:"votes"`
}
