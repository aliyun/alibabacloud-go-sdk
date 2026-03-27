// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBizTracesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListBizTracesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListBizTracesRequest
	GetNextToken() *string
	SetWorkspace(v string) *ListBizTracesRequest
	GetWorkspace() *string
}

type ListBizTracesRequest struct {
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// aa9d0e569b88098a0e3155c29b473201a
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// example:
	//
	// default-cms-xxxxxx-cn-beijing
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s ListBizTracesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListBizTracesRequest) GoString() string {
	return s.String()
}

func (s *ListBizTracesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListBizTracesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListBizTracesRequest) GetWorkspace() *string {
	return s.Workspace
}

func (s *ListBizTracesRequest) SetMaxResults(v int32) *ListBizTracesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListBizTracesRequest) SetNextToken(v string) *ListBizTracesRequest {
	s.NextToken = &v
	return s
}

func (s *ListBizTracesRequest) SetWorkspace(v string) *ListBizTracesRequest {
	s.Workspace = &v
	return s
}

func (s *ListBizTracesRequest) Validate() error {
	return dara.Validate(s)
}
