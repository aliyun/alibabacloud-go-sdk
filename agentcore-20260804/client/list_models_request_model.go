// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListModelsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConnectionId(v string) *ListModelsRequest
	GetConnectionId() *string
	SetMaxResults(v int32) *ListModelsRequest
	GetMaxResults() *int32
	SetModelName(v string) *ListModelsRequest
	GetModelName() *string
	SetNextToken(v string) *ListModelsRequest
	GetNextToken() *string
}

type ListModelsRequest struct {
	// The model connection ID used to filter models.
	//
	// example:
	//
	// mc-1
	ConnectionId *string `json:"connectionId,omitempty" xml:"connectionId,omitempty"`
	// The number of results per page. Valid values: 0 to 100. If this parameter is not set or set to 0, the default value 10 is used.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The upstream model name.
	//
	// example:
	//
	// qwen-max
	ModelName *string `json:"modelName,omitempty" xml:"modelName,omitempty"`
	// The pagination token. Pass the token returned from the previous query. An empty response indicates that no more pages are available.
	//
	// example:
	//
	// bW9kZWwtbWFuYWdlbWVudC1vZmZzZXQ6bW9kZWw6MTA
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s ListModelsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListModelsRequest) GoString() string {
	return s.String()
}

func (s *ListModelsRequest) GetConnectionId() *string {
	return s.ConnectionId
}

func (s *ListModelsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListModelsRequest) GetModelName() *string {
	return s.ModelName
}

func (s *ListModelsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListModelsRequest) SetConnectionId(v string) *ListModelsRequest {
	s.ConnectionId = &v
	return s
}

func (s *ListModelsRequest) SetMaxResults(v int32) *ListModelsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListModelsRequest) SetModelName(v string) *ListModelsRequest {
	s.ModelName = &v
	return s
}

func (s *ListModelsRequest) SetNextToken(v string) *ListModelsRequest {
	s.NextToken = &v
	return s
}

func (s *ListModelsRequest) Validate() error {
	return dara.Validate(s)
}
