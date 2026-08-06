// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPromptsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFrameworkType(v string) *ListPromptsRequest
	GetFrameworkType() *string
	SetOrder(v string) *ListPromptsRequest
	GetOrder() *string
	SetPageNumber(v string) *ListPromptsRequest
	GetPageNumber() *string
	SetPageSize(v string) *ListPromptsRequest
	GetPageSize() *string
	SetSortBy(v string) *ListPromptsRequest
	GetSortBy() *string
	SetWorkspaceId(v string) *ListPromptsRequest
	GetWorkspaceId() *string
}

type ListPromptsRequest struct {
	// The prompt template framework type.
	//
	// example:
	//
	// ICIO
	FrameworkType *string `json:"FrameworkType,omitempty" xml:"FrameworkType,omitempty"`
	// The sorting order for the specified field during paging. Default value: ASC.
	//
	// - ASC: ascending order.
	//
	// - DESC: descending order.
	//
	// example:
	//
	// ASC
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// The page number, starting from 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The page size. Default value: 20.
	//
	// example:
	//
	// 20
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The field used for sorting. Valid values:
	//
	// - Name: the run name.
	//
	// - GmtCreateTime (default): the run creation time.
	//
	// example:
	//
	// GmtCreateTime
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// The workspace ID. For information about how to obtain the workspace ID, see [ListWorkspaces](https://help.aliyun.com/document_detail/449124.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// 145883
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListPromptsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPromptsRequest) GoString() string {
	return s.String()
}

func (s *ListPromptsRequest) GetFrameworkType() *string {
	return s.FrameworkType
}

func (s *ListPromptsRequest) GetOrder() *string {
	return s.Order
}

func (s *ListPromptsRequest) GetPageNumber() *string {
	return s.PageNumber
}

func (s *ListPromptsRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *ListPromptsRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *ListPromptsRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListPromptsRequest) SetFrameworkType(v string) *ListPromptsRequest {
	s.FrameworkType = &v
	return s
}

func (s *ListPromptsRequest) SetOrder(v string) *ListPromptsRequest {
	s.Order = &v
	return s
}

func (s *ListPromptsRequest) SetPageNumber(v string) *ListPromptsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListPromptsRequest) SetPageSize(v string) *ListPromptsRequest {
	s.PageSize = &v
	return s
}

func (s *ListPromptsRequest) SetSortBy(v string) *ListPromptsRequest {
	s.SortBy = &v
	return s
}

func (s *ListPromptsRequest) SetWorkspaceId(v string) *ListPromptsRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ListPromptsRequest) Validate() error {
	return dara.Validate(s)
}
