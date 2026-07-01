// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApiKeysRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiKeyId(v int64) *ListApiKeysRequest
	GetApiKeyId() *int64
	SetDescription(v string) *ListApiKeysRequest
	GetDescription() *string
	SetMaxResults(v int32) *ListApiKeysRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListApiKeysRequest
	GetNextToken() *string
	SetOrder(v string) *ListApiKeysRequest
	GetOrder() *string
	SetOrderBy(v string) *ListApiKeysRequest
	GetOrderBy() *string
	SetWorkspaceId(v string) *ListApiKeysRequest
	GetWorkspaceId() *string
}

type ListApiKeysRequest struct {
	// The API key ID for exact match.
	//
	// example:
	//
	// 3076140
	ApiKeyId *int64 `json:"apiKeyId,omitempty" xml:"apiKeyId,omitempty"`
	// The keyword for fuzzy match by description.
	//
	// example:
	//
	// test
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 30
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The token used to retrieve more results. This parameter is not required for the first query. For subsequent queries, use the token obtained from the previous response.
	//
	// example:
	//
	// w9Z+S5+TZyw=
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The sort order. Valid values:
	//
	// - DESC (default)
	//
	// - ASC
	//
	// example:
	//
	// ASC
	Order *string `json:"order,omitempty" xml:"order,omitempty"`
	// The field by which to sort results. Valid values:
	//
	// - apiKeyId (default)
	//
	// - gmtCreate
	//
	// example:
	//
	// gmtCreate
	OrderBy *string `json:"orderBy,omitempty" xml:"orderBy,omitempty"`
	// The workspace ID for exact match.
	//
	// example:
	//
	// ws-ac3ef438bec22dc5
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s ListApiKeysRequest) String() string {
	return dara.Prettify(s)
}

func (s ListApiKeysRequest) GoString() string {
	return s.String()
}

func (s *ListApiKeysRequest) GetApiKeyId() *int64 {
	return s.ApiKeyId
}

func (s *ListApiKeysRequest) GetDescription() *string {
	return s.Description
}

func (s *ListApiKeysRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListApiKeysRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListApiKeysRequest) GetOrder() *string {
	return s.Order
}

func (s *ListApiKeysRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *ListApiKeysRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListApiKeysRequest) SetApiKeyId(v int64) *ListApiKeysRequest {
	s.ApiKeyId = &v
	return s
}

func (s *ListApiKeysRequest) SetDescription(v string) *ListApiKeysRequest {
	s.Description = &v
	return s
}

func (s *ListApiKeysRequest) SetMaxResults(v int32) *ListApiKeysRequest {
	s.MaxResults = &v
	return s
}

func (s *ListApiKeysRequest) SetNextToken(v string) *ListApiKeysRequest {
	s.NextToken = &v
	return s
}

func (s *ListApiKeysRequest) SetOrder(v string) *ListApiKeysRequest {
	s.Order = &v
	return s
}

func (s *ListApiKeysRequest) SetOrderBy(v string) *ListApiKeysRequest {
	s.OrderBy = &v
	return s
}

func (s *ListApiKeysRequest) SetWorkspaceId(v string) *ListApiKeysRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ListApiKeysRequest) Validate() error {
	return dara.Validate(s)
}
