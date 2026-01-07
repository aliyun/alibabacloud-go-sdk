// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataAgentWorkspaceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDMSUnit(v string) *ListDataAgentWorkspaceRequest
	GetDMSUnit() *string
	SetMaxResults(v int32) *ListDataAgentWorkspaceRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListDataAgentWorkspaceRequest
	GetNextToken() *string
	SetOrder(v string) *ListDataAgentWorkspaceRequest
	GetOrder() *string
	SetOrderBy(v string) *ListDataAgentWorkspaceRequest
	GetOrderBy() *string
	SetPageNumber(v string) *ListDataAgentWorkspaceRequest
	GetPageNumber() *string
	SetPageSize(v string) *ListDataAgentWorkspaceRequest
	GetPageSize() *string
	SetWorkspaceName(v string) *ListDataAgentWorkspaceRequest
	GetWorkspaceName() *string
	SetWorkspaceType(v string) *ListDataAgentWorkspaceRequest
	GetWorkspaceType() *string
}

type ListDataAgentWorkspaceRequest struct {
	// example:
	//
	// cn-hangzhou
	DMSUnit *string `json:"DMSUnit,omitempty" xml:"DMSUnit,omitempty"`
	// example:
	//
	// no use
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// no use
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// DESC
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// example:
	//
	// CreateTime
	OrderBy *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// workspaceTest
	WorkspaceName *string `json:"WorkspaceName,omitempty" xml:"WorkspaceName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// MY
	WorkspaceType *string `json:"WorkspaceType,omitempty" xml:"WorkspaceType,omitempty"`
}

func (s ListDataAgentWorkspaceRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDataAgentWorkspaceRequest) GoString() string {
	return s.String()
}

func (s *ListDataAgentWorkspaceRequest) GetDMSUnit() *string {
	return s.DMSUnit
}

func (s *ListDataAgentWorkspaceRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListDataAgentWorkspaceRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListDataAgentWorkspaceRequest) GetOrder() *string {
	return s.Order
}

func (s *ListDataAgentWorkspaceRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *ListDataAgentWorkspaceRequest) GetPageNumber() *string {
	return s.PageNumber
}

func (s *ListDataAgentWorkspaceRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *ListDataAgentWorkspaceRequest) GetWorkspaceName() *string {
	return s.WorkspaceName
}

func (s *ListDataAgentWorkspaceRequest) GetWorkspaceType() *string {
	return s.WorkspaceType
}

func (s *ListDataAgentWorkspaceRequest) SetDMSUnit(v string) *ListDataAgentWorkspaceRequest {
	s.DMSUnit = &v
	return s
}

func (s *ListDataAgentWorkspaceRequest) SetMaxResults(v int32) *ListDataAgentWorkspaceRequest {
	s.MaxResults = &v
	return s
}

func (s *ListDataAgentWorkspaceRequest) SetNextToken(v string) *ListDataAgentWorkspaceRequest {
	s.NextToken = &v
	return s
}

func (s *ListDataAgentWorkspaceRequest) SetOrder(v string) *ListDataAgentWorkspaceRequest {
	s.Order = &v
	return s
}

func (s *ListDataAgentWorkspaceRequest) SetOrderBy(v string) *ListDataAgentWorkspaceRequest {
	s.OrderBy = &v
	return s
}

func (s *ListDataAgentWorkspaceRequest) SetPageNumber(v string) *ListDataAgentWorkspaceRequest {
	s.PageNumber = &v
	return s
}

func (s *ListDataAgentWorkspaceRequest) SetPageSize(v string) *ListDataAgentWorkspaceRequest {
	s.PageSize = &v
	return s
}

func (s *ListDataAgentWorkspaceRequest) SetWorkspaceName(v string) *ListDataAgentWorkspaceRequest {
	s.WorkspaceName = &v
	return s
}

func (s *ListDataAgentWorkspaceRequest) SetWorkspaceType(v string) *ListDataAgentWorkspaceRequest {
	s.WorkspaceType = &v
	return s
}

func (s *ListDataAgentWorkspaceRequest) Validate() error {
	return dara.Validate(s)
}
