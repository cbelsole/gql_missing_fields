// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package gql_missing_fields

type NewTodo struct {
	User1 UserReq  `json:"user1"`
	User2 *UserReq `json:"user2"`
	Todo  TodoReq  `json:"todo"`
}

type Todo struct {
	ID   string `json:"id"`
	Text string `json:"text"`
	Done bool   `json:"done"`
	User User   `json:"user"`
}

type TodoReq struct {
	NotRequiredTodo *string `json:"notRequiredTodo"`
	RequiredTodo    string  `json:"requiredTodo"`
}

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type UserReq struct {
	RequiredUser    string  `json:"requiredUser"`
	NotRequiredUser *string `json:"notRequiredUser"`
}
