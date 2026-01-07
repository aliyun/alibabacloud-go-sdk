// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataAgentWorkspaceMemberRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDMSUnit(v string) *ListDataAgentWorkspaceMemberRequest
	GetDMSUnit() *string
	SetMaxResults(v int32) *ListDataAgentWorkspaceMemberRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListDataAgentWorkspaceMemberRequest
	GetNextToken() *string
	SetOrder(v string) *ListDataAgentWorkspaceMemberRequest
	GetOrder() *string
	SetOrderBy(v string) *ListDataAgentWorkspaceMemberRequest
	GetOrderBy() *string
	SetPageNumber(v int64) *ListDataAgentWorkspaceMemberRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *ListDataAgentWorkspaceMemberRequest
	GetPageSize() *int64
	SetSearchMemberId(v string) *ListDataAgentWorkspaceMemberRequest
	GetSearchMemberId() *string
	SetSearchRoleName(v string) *ListDataAgentWorkspaceMemberRequest
	GetSearchRoleName() *string
	SetWorkspaceId(v string) *ListDataAgentWorkspaceMemberRequest
	GetWorkspaceId() *string
}

type ListDataAgentWorkspaceMemberRequest struct {
	// example:
	//
	// cn-hangzhou
	DMSUnit *string `json:"DMSUnit,omitempty" xml:"DMSUnit,omitempty"`
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// NesLoK****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// Desc
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// example:
	//
	// gmt_create
	OrderBy *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 20273
	SearchMemberId *string `json:"SearchMemberId,omitempty" xml:"SearchMemberId,omitempty"`
	// example:
	//
	// yunqi
	SearchRoleName *string `json:"SearchRoleName,omitempty" xml:"SearchRoleName,omitempty"`
	// example:
	//
	// 20923*****7291
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListDataAgentWorkspaceMemberRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDataAgentWorkspaceMemberRequest) GoString() string {
	return s.String()
}

func (s *ListDataAgentWorkspaceMemberRequest) GetDMSUnit() *string {
	return s.DMSUnit
}

func (s *ListDataAgentWorkspaceMemberRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListDataAgentWorkspaceMemberRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListDataAgentWorkspaceMemberRequest) GetOrder() *string {
	return s.Order
}

func (s *ListDataAgentWorkspaceMemberRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *ListDataAgentWorkspaceMemberRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListDataAgentWorkspaceMemberRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListDataAgentWorkspaceMemberRequest) GetSearchMemberId() *string {
	return s.SearchMemberId
}

func (s *ListDataAgentWorkspaceMemberRequest) GetSearchRoleName() *string {
	return s.SearchRoleName
}

func (s *ListDataAgentWorkspaceMemberRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListDataAgentWorkspaceMemberRequest) SetDMSUnit(v string) *ListDataAgentWorkspaceMemberRequest {
	s.DMSUnit = &v
	return s
}

func (s *ListDataAgentWorkspaceMemberRequest) SetMaxResults(v int32) *ListDataAgentWorkspaceMemberRequest {
	s.MaxResults = &v
	return s
}

func (s *ListDataAgentWorkspaceMemberRequest) SetNextToken(v string) *ListDataAgentWorkspaceMemberRequest {
	s.NextToken = &v
	return s
}

func (s *ListDataAgentWorkspaceMemberRequest) SetOrder(v string) *ListDataAgentWorkspaceMemberRequest {
	s.Order = &v
	return s
}

func (s *ListDataAgentWorkspaceMemberRequest) SetOrderBy(v string) *ListDataAgentWorkspaceMemberRequest {
	s.OrderBy = &v
	return s
}

func (s *ListDataAgentWorkspaceMemberRequest) SetPageNumber(v int64) *ListDataAgentWorkspaceMemberRequest {
	s.PageNumber = &v
	return s
}

func (s *ListDataAgentWorkspaceMemberRequest) SetPageSize(v int64) *ListDataAgentWorkspaceMemberRequest {
	s.PageSize = &v
	return s
}

func (s *ListDataAgentWorkspaceMemberRequest) SetSearchMemberId(v string) *ListDataAgentWorkspaceMemberRequest {
	s.SearchMemberId = &v
	return s
}

func (s *ListDataAgentWorkspaceMemberRequest) SetSearchRoleName(v string) *ListDataAgentWorkspaceMemberRequest {
	s.SearchRoleName = &v
	return s
}

func (s *ListDataAgentWorkspaceMemberRequest) SetWorkspaceId(v string) *ListDataAgentWorkspaceMemberRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ListDataAgentWorkspaceMemberRequest) Validate() error {
	return dara.Validate(s)
}
