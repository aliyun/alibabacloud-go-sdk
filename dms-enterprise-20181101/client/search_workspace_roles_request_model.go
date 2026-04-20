// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchWorkspaceRolesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *SearchWorkspaceRolesRequest
	GetClientToken() *string
	SetMaxResults(v int32) *SearchWorkspaceRolesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *SearchWorkspaceRolesRequest
	GetNextToken() *string
	SetSearchKey(v string) *SearchWorkspaceRolesRequest
	GetSearchKey() *string
	SetWorkspaceId(v string) *SearchWorkspaceRolesRequest
	GetWorkspaceId() *string
}

type SearchWorkspaceRolesRequest struct {
	// example:
	//
	// 123
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// f056501ada12c1cc
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// test
	SearchKey *string `json:"SearchKey,omitempty" xml:"SearchKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 12****
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s SearchWorkspaceRolesRequest) String() string {
	return dara.Prettify(s)
}

func (s SearchWorkspaceRolesRequest) GoString() string {
	return s.String()
}

func (s *SearchWorkspaceRolesRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *SearchWorkspaceRolesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *SearchWorkspaceRolesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *SearchWorkspaceRolesRequest) GetSearchKey() *string {
	return s.SearchKey
}

func (s *SearchWorkspaceRolesRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *SearchWorkspaceRolesRequest) SetClientToken(v string) *SearchWorkspaceRolesRequest {
	s.ClientToken = &v
	return s
}

func (s *SearchWorkspaceRolesRequest) SetMaxResults(v int32) *SearchWorkspaceRolesRequest {
	s.MaxResults = &v
	return s
}

func (s *SearchWorkspaceRolesRequest) SetNextToken(v string) *SearchWorkspaceRolesRequest {
	s.NextToken = &v
	return s
}

func (s *SearchWorkspaceRolesRequest) SetSearchKey(v string) *SearchWorkspaceRolesRequest {
	s.SearchKey = &v
	return s
}

func (s *SearchWorkspaceRolesRequest) SetWorkspaceId(v string) *SearchWorkspaceRolesRequest {
	s.WorkspaceId = &v
	return s
}

func (s *SearchWorkspaceRolesRequest) Validate() error {
	return dara.Validate(s)
}
