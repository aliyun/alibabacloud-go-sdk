// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListModelLimitsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int64) *ListModelLimitsRequest
	GetMaxResults() *int64
	SetModel(v string) *ListModelLimitsRequest
	GetModel() *string
	SetName(v string) *ListModelLimitsRequest
	GetName() *string
	SetNextToken(v string) *ListModelLimitsRequest
	GetNextToken() *string
	SetWorkspaceId(v string) *ListModelLimitsRequest
	GetWorkspaceId() *string
}

type ListModelLimitsRequest struct {
	// The maximum number of results to return. Valid values: 0 to 200.
	//
	// example:
	//
	// 20
	MaxResults *int64 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The model for exact match.
	//
	// example:
	//
	// qwen-plus
	Model *string `json:"model,omitempty" xml:"model,omitempty"`
	// The model name for fuzzy match.
	//
	// example:
	//
	// qwen-plus
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The pagination token.
	//
	// example:
	//
	// lwytFRtLdNk=
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The workspace ID.
	//
	// example:
	//
	// ws-8af73c50f5596193
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s ListModelLimitsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListModelLimitsRequest) GoString() string {
	return s.String()
}

func (s *ListModelLimitsRequest) GetMaxResults() *int64 {
	return s.MaxResults
}

func (s *ListModelLimitsRequest) GetModel() *string {
	return s.Model
}

func (s *ListModelLimitsRequest) GetName() *string {
	return s.Name
}

func (s *ListModelLimitsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListModelLimitsRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListModelLimitsRequest) SetMaxResults(v int64) *ListModelLimitsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListModelLimitsRequest) SetModel(v string) *ListModelLimitsRequest {
	s.Model = &v
	return s
}

func (s *ListModelLimitsRequest) SetName(v string) *ListModelLimitsRequest {
	s.Name = &v
	return s
}

func (s *ListModelLimitsRequest) SetNextToken(v string) *ListModelLimitsRequest {
	s.NextToken = &v
	return s
}

func (s *ListModelLimitsRequest) SetWorkspaceId(v string) *ListModelLimitsRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ListModelLimitsRequest) Validate() error {
	return dara.Validate(s)
}
