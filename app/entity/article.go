package entity

type ArticlesListRequest struct {
	Author string `json:"author"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

type ArticlesListResponse struct {
	Articles []*Articles
	Page     *Pagination
}

type Articles struct {
	Id        string `json:"id" gorm:"id" `
	Author    string `json:"author" gorm:"author" `
	Title     string `json:"title" gorm:"title" `
	Body      string `json:"body" gorm:"body" `
	CreatedAt string `json:"created" gorm:"created" `
}
