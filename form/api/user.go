package api

type User struct {
	ID   uint   `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
}

type UserPosts struct {
	ID    uint   `json:"id" binding:"required"`
	Name  string `json:"name" binding:"required"`
	Posts []Post `json:"posts"`
}
