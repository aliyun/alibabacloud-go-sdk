// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLabelTablesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLabelTableIds(v []*string) *ListLabelTablesRequest
	GetLabelTableIds() []*string
	SetName(v string) *ListLabelTablesRequest
	GetName() *string
	SetOrder(v string) *ListLabelTablesRequest
	GetOrder() *string
	SetOwner(v string) *ListLabelTablesRequest
	GetOwner() *string
	SetPageNumber(v int64) *ListLabelTablesRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *ListLabelTablesRequest
	GetPageSize() *int64
	SetProjectId(v string) *ListLabelTablesRequest
	GetProjectId() *string
	SetSortBy(v string) *ListLabelTablesRequest
	GetSortBy() *string
}

type ListLabelTablesRequest struct {
	// The IDs of the label tables by which to filter the results.
	LabelTableIds []*string `json:"LabelTableIds,omitempty" xml:"LabelTableIds,omitempty" type:"Repeated"`
	// The name of the label table.
	//
	// example:
	//
	// label_table1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The sort order. Valid values:
	//
	// - `Asc`: ascending order.
	//
	// - `Desc`: descending order.
	//
	// example:
	//
	// Asc
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// The Alibaba Cloud account ID of the owner.
	//
	// example:
	//
	// 1231432432****
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The page number. Page numbers start from 1. Default value: 1.
	//
	// example:
	//
	// 10
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Default value: 10.
	//
	// example:
	//
	// 1
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The project ID. You can call the `ListProjects` operation to obtain this ID.
	//
	// example:
	//
	// project1
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The field by which to sort the results.
	//
	// example:
	//
	// GmtModifiedTime
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
}

func (s ListLabelTablesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListLabelTablesRequest) GoString() string {
	return s.String()
}

func (s *ListLabelTablesRequest) GetLabelTableIds() []*string {
	return s.LabelTableIds
}

func (s *ListLabelTablesRequest) GetName() *string {
	return s.Name
}

func (s *ListLabelTablesRequest) GetOrder() *string {
	return s.Order
}

func (s *ListLabelTablesRequest) GetOwner() *string {
	return s.Owner
}

func (s *ListLabelTablesRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListLabelTablesRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListLabelTablesRequest) GetProjectId() *string {
	return s.ProjectId
}

func (s *ListLabelTablesRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *ListLabelTablesRequest) SetLabelTableIds(v []*string) *ListLabelTablesRequest {
	s.LabelTableIds = v
	return s
}

func (s *ListLabelTablesRequest) SetName(v string) *ListLabelTablesRequest {
	s.Name = &v
	return s
}

func (s *ListLabelTablesRequest) SetOrder(v string) *ListLabelTablesRequest {
	s.Order = &v
	return s
}

func (s *ListLabelTablesRequest) SetOwner(v string) *ListLabelTablesRequest {
	s.Owner = &v
	return s
}

func (s *ListLabelTablesRequest) SetPageNumber(v int64) *ListLabelTablesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListLabelTablesRequest) SetPageSize(v int64) *ListLabelTablesRequest {
	s.PageSize = &v
	return s
}

func (s *ListLabelTablesRequest) SetProjectId(v string) *ListLabelTablesRequest {
	s.ProjectId = &v
	return s
}

func (s *ListLabelTablesRequest) SetSortBy(v string) *ListLabelTablesRequest {
	s.SortBy = &v
	return s
}

func (s *ListLabelTablesRequest) Validate() error {
	return dara.Validate(s)
}
