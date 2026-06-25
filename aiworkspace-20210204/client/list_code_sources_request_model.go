// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCodeSourcesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDisplayName(v string) *ListCodeSourcesRequest
	GetDisplayName() *string
	SetOrder(v string) *ListCodeSourcesRequest
	GetOrder() *string
	SetPageNumber(v int32) *ListCodeSourcesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListCodeSourcesRequest
	GetPageSize() *int32
	SetSortBy(v string) *ListCodeSourcesRequest
	GetSortBy() *string
	SetWorkspaceId(v string) *ListCodeSourcesRequest
	GetWorkspaceId() *string
}

type ListCodeSourcesRequest struct {
	// The display name of the code source configuration. Fuzzy match is supported.
	//
	// example:
	//
	// MyDataSource
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The sort order. Valid values:
	//
	// - ASC (default): Ascending order.
	//
	// - DESC: Descending order.
	//
	// example:
	//
	// desc
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// The page number. The value starts from 1. The default value is 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. The default value is 20.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The field to use for sorting. Valid values:
	//
	// - GmtModifyTime: The time when the code source was last modified.
	//
	// - DisplayName: The display name.
	//
	// - CodeSourceId: The code source ID.
	//
	// - GmtCreateTime (default): The time when the code source was created.
	//
	// example:
	//
	// GmtModifyTime
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// The workspace ID. For more information, see [ListWorkspaces](https://help.aliyun.com/document_detail/449124.html).
	//
	// example:
	//
	// 1234
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListCodeSourcesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListCodeSourcesRequest) GoString() string {
	return s.String()
}

func (s *ListCodeSourcesRequest) GetDisplayName() *string {
	return s.DisplayName
}

func (s *ListCodeSourcesRequest) GetOrder() *string {
	return s.Order
}

func (s *ListCodeSourcesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListCodeSourcesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListCodeSourcesRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *ListCodeSourcesRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListCodeSourcesRequest) SetDisplayName(v string) *ListCodeSourcesRequest {
	s.DisplayName = &v
	return s
}

func (s *ListCodeSourcesRequest) SetOrder(v string) *ListCodeSourcesRequest {
	s.Order = &v
	return s
}

func (s *ListCodeSourcesRequest) SetPageNumber(v int32) *ListCodeSourcesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListCodeSourcesRequest) SetPageSize(v int32) *ListCodeSourcesRequest {
	s.PageSize = &v
	return s
}

func (s *ListCodeSourcesRequest) SetSortBy(v string) *ListCodeSourcesRequest {
	s.SortBy = &v
	return s
}

func (s *ListCodeSourcesRequest) SetWorkspaceId(v string) *ListCodeSourcesRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ListCodeSourcesRequest) Validate() error {
	return dara.Validate(s)
}
