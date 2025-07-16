// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLabelTablesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLabelTableIdsShrink(v string) *ListLabelTablesShrinkRequest
	GetLabelTableIdsShrink() *string
	SetName(v string) *ListLabelTablesShrinkRequest
	GetName() *string
	SetOrder(v string) *ListLabelTablesShrinkRequest
	GetOrder() *string
	SetOwner(v string) *ListLabelTablesShrinkRequest
	GetOwner() *string
	SetPageNumber(v int64) *ListLabelTablesShrinkRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *ListLabelTablesShrinkRequest
	GetPageSize() *int64
	SetProjectId(v string) *ListLabelTablesShrinkRequest
	GetProjectId() *string
	SetSortBy(v string) *ListLabelTablesShrinkRequest
	GetSortBy() *string
}

type ListLabelTablesShrinkRequest struct {
	LabelTableIdsShrink *string `json:"LabelTableIds,omitempty" xml:"LabelTableIds,omitempty"`
	// example:
	//
	// label_table1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// Asc
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// example:
	//
	// 1231432432****
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// example:
	//
	// 10
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 1
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// project1
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// example:
	//
	// GmtModifiedTime
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
}

func (s ListLabelTablesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListLabelTablesShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListLabelTablesShrinkRequest) GetLabelTableIdsShrink() *string {
	return s.LabelTableIdsShrink
}

func (s *ListLabelTablesShrinkRequest) GetName() *string {
	return s.Name
}

func (s *ListLabelTablesShrinkRequest) GetOrder() *string {
	return s.Order
}

func (s *ListLabelTablesShrinkRequest) GetOwner() *string {
	return s.Owner
}

func (s *ListLabelTablesShrinkRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListLabelTablesShrinkRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListLabelTablesShrinkRequest) GetProjectId() *string {
	return s.ProjectId
}

func (s *ListLabelTablesShrinkRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *ListLabelTablesShrinkRequest) SetLabelTableIdsShrink(v string) *ListLabelTablesShrinkRequest {
	s.LabelTableIdsShrink = &v
	return s
}

func (s *ListLabelTablesShrinkRequest) SetName(v string) *ListLabelTablesShrinkRequest {
	s.Name = &v
	return s
}

func (s *ListLabelTablesShrinkRequest) SetOrder(v string) *ListLabelTablesShrinkRequest {
	s.Order = &v
	return s
}

func (s *ListLabelTablesShrinkRequest) SetOwner(v string) *ListLabelTablesShrinkRequest {
	s.Owner = &v
	return s
}

func (s *ListLabelTablesShrinkRequest) SetPageNumber(v int64) *ListLabelTablesShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *ListLabelTablesShrinkRequest) SetPageSize(v int64) *ListLabelTablesShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListLabelTablesShrinkRequest) SetProjectId(v string) *ListLabelTablesShrinkRequest {
	s.ProjectId = &v
	return s
}

func (s *ListLabelTablesShrinkRequest) SetSortBy(v string) *ListLabelTablesShrinkRequest {
	s.SortBy = &v
	return s
}

func (s *ListLabelTablesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
