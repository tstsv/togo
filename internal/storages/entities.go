package storages

// Task reflects tasks in DB
type Task struct {
	ID          string `json:"id"`
	Content     string `json:"content"`
	UserID      string `json:"user_id"`
	CreatedDate string `json:"created_date"`
}

// User reflects users data from DB
type User struct {
	ID       string `json:"id"`
	Username string `json:"user_name"`
	Password string `json:"password"`
	MaxTodo  int64  `json:"max_todo"`
}
