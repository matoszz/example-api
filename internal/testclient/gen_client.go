// Code generated by github.com/Yamashou/gqlgenc, DO NOT EDIT.

package testclient

import (
	"context"
	"net/http"
	"time"

	"github.com/Yamashou/gqlgenc/client"
	"go.infratographer.com/x/gidx"
)

type TestClient interface {
	GetTodo(ctx context.Context, id gidx.PrefixedID, httpRequestOptions ...client.HTTPRequestOption) (*GetTodo, error)
	ListTodos(ctx context.Context, id gidx.PrefixedID, orderBy *TodoOrder, httpRequestOptions ...client.HTTPRequestOption) (*ListTodos, error)
	TodoCreate(ctx context.Context, input CreateTodoInput, httpRequestOptions ...client.HTTPRequestOption) (*TodoCreate, error)
	TodoDelete(ctx context.Context, id gidx.PrefixedID, httpRequestOptions ...client.HTTPRequestOption) (*TodoDelete, error)
	TodoUpdate(ctx context.Context, id gidx.PrefixedID, input UpdateTodoInput, httpRequestOptions ...client.HTTPRequestOption) (*TodoUpdate, error)
}

type Client struct {
	Client *client.Client
}

func NewClient(cli *http.Client, baseURL string, options ...client.HTTPRequestOption) TestClient {
	return &Client{Client: client.NewClient(cli, baseURL, options...)}
}

type Query struct {
	Todo     Todo     "json:\"todo\" graphql:\"todo\""
	Entities []Entity "json:\"_entities\" graphql:\"_entities\""
	Service  Service  "json:\"_service\" graphql:\"_service\""
}
type Mutation struct {
	TodoCreate TodoCreatePayload "json:\"todoCreate\" graphql:\"todoCreate\""
	TodoUpdate TodoUpdatePayload "json:\"todoUpdate\" graphql:\"todoUpdate\""
	TodoDelete TodoDeletePayload "json:\"todoDelete\" graphql:\"todoDelete\""
}
type GetTodo struct {
	Todo struct {
		ID          gidx.PrefixedID "json:\"id\" graphql:\"id\""
		Name        string          "json:\"name\" graphql:\"name\""
		Description *string         "json:\"description\" graphql:\"description\""
		Tenant      struct {
			ID gidx.PrefixedID "json:\"id\" graphql:\"id\""
		} "json:\"tenant\" graphql:\"tenant\""
		CreatedAt time.Time "json:\"createdAt\" graphql:\"createdAt\""
		UpdatedAt time.Time "json:\"updatedAt\" graphql:\"updatedAt\""
	} "json:\"todo\" graphql:\"todo\""
}
type ListTodos struct {
	Entities []*struct {
		Todo struct {
			Edges []*struct {
				Node *struct {
					ID   gidx.PrefixedID "json:\"id\" graphql:\"id\""
					Name string          "json:\"name\" graphql:\"name\""
				} "json:\"node\" graphql:\"node\""
			} "json:\"edges\" graphql:\"edges\""
		} "json:\"todo\" graphql:\"todo\""
	} "json:\"_entities\" graphql:\"_entities\""
}
type TodoCreate struct {
	TodoCreate struct {
		Todo struct {
			ID          gidx.PrefixedID "json:\"id\" graphql:\"id\""
			Name        string          "json:\"name\" graphql:\"name\""
			Description *string         "json:\"description\" graphql:\"description\""
			Tenant      struct {
				ID gidx.PrefixedID "json:\"id\" graphql:\"id\""
			} "json:\"tenant\" graphql:\"tenant\""
			CreatedAt time.Time "json:\"createdAt\" graphql:\"createdAt\""
			UpdatedAt time.Time "json:\"updatedAt\" graphql:\"updatedAt\""
		} "json:\"todo\" graphql:\"todo\""
	} "json:\"todoCreate\" graphql:\"todoCreate\""
}
type TodoDelete struct {
	TodoDelete struct {
		DeletedID gidx.PrefixedID "json:\"deletedID\" graphql:\"deletedID\""
	} "json:\"todoDelete\" graphql:\"todoDelete\""
}
type TodoUpdate struct {
	TodoUpdate struct {
		Todo struct {
			ID          gidx.PrefixedID "json:\"id\" graphql:\"id\""
			Name        string          "json:\"name\" graphql:\"name\""
			Description *string         "json:\"description\" graphql:\"description\""
			CreatedAt   time.Time       "json:\"createdAt\" graphql:\"createdAt\""
			UpdatedAt   time.Time       "json:\"updatedAt\" graphql:\"updatedAt\""
		} "json:\"todo\" graphql:\"todo\""
	} "json:\"todoUpdate\" graphql:\"todoUpdate\""
}

const GetTodoDocument = `query GetTodo ($id: ID!) {
	todo(id: $id) {
		id
		name
		description
		tenant {
			id
		}
		createdAt
		updatedAt
	}
}
`

func (c *Client) GetTodo(ctx context.Context, id gidx.PrefixedID, httpRequestOptions ...client.HTTPRequestOption) (*GetTodo, error) {
	vars := map[string]interface{}{
		"id": id,
	}

	var res GetTodo
	if err := c.Client.Post(ctx, "GetTodo", GetTodoDocument, &res, vars, httpRequestOptions...); err != nil {
		return nil, err
	}

	return &res, nil
}

const ListTodosDocument = `query ListTodos ($id: ID!, $orderBy: TodoOrder) {
	_entities(representations: [{__typename:"Tenant",id:$id}]) {
		... on Tenant {
			todo(orderBy: $orderBy) {
				edges {
					node {
						id
						name
					}
				}
			}
		}
	}
}
`

func (c *Client) ListTodos(ctx context.Context, id gidx.PrefixedID, orderBy *TodoOrder, httpRequestOptions ...client.HTTPRequestOption) (*ListTodos, error) {
	vars := map[string]interface{}{
		"id":      id,
		"orderBy": orderBy,
	}

	var res ListTodos
	if err := c.Client.Post(ctx, "ListTodos", ListTodosDocument, &res, vars, httpRequestOptions...); err != nil {
		return nil, err
	}

	return &res, nil
}

const TodoCreateDocument = `mutation TodoCreate ($input: CreateTodoInput!) {
	todoCreate(input: $input) {
		todo {
			id
			name
			description
			tenant {
				id
			}
			createdAt
			updatedAt
		}
	}
}
`

func (c *Client) TodoCreate(ctx context.Context, input CreateTodoInput, httpRequestOptions ...client.HTTPRequestOption) (*TodoCreate, error) {
	vars := map[string]interface{}{
		"input": input,
	}

	var res TodoCreate
	if err := c.Client.Post(ctx, "TodoCreate", TodoCreateDocument, &res, vars, httpRequestOptions...); err != nil {
		return nil, err
	}

	return &res, nil
}

const TodoDeleteDocument = `mutation TodoDelete ($id: ID!) {
	todoDelete(id: $id) {
		deletedID
	}
}
`

func (c *Client) TodoDelete(ctx context.Context, id gidx.PrefixedID, httpRequestOptions ...client.HTTPRequestOption) (*TodoDelete, error) {
	vars := map[string]interface{}{
		"id": id,
	}

	var res TodoDelete
	if err := c.Client.Post(ctx, "TodoDelete", TodoDeleteDocument, &res, vars, httpRequestOptions...); err != nil {
		return nil, err
	}

	return &res, nil
}

const TodoUpdateDocument = `mutation TodoUpdate ($id: ID!, $input: UpdateTodoInput!) {
	todoUpdate(id: $id, input: $input) {
		todo {
			id
			name
			description
			createdAt
			updatedAt
		}
	}
}
`

func (c *Client) TodoUpdate(ctx context.Context, id gidx.PrefixedID, input UpdateTodoInput, httpRequestOptions ...client.HTTPRequestOption) (*TodoUpdate, error) {
	vars := map[string]interface{}{
		"id":    id,
		"input": input,
	}

	var res TodoUpdate
	if err := c.Client.Post(ctx, "TodoUpdate", TodoUpdateDocument, &res, vars, httpRequestOptions...); err != nil {
		return nil, err
	}

	return &res, nil
}
