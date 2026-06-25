// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGroupsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFilter(v string) *ListGroupsShrinkRequest
	GetFilter() *string
	SetLabelsShrink(v string) *ListGroupsShrinkRequest
	GetLabelsShrink() *string
	SetOrder(v string) *ListGroupsShrinkRequest
	GetOrder() *string
	SetPageNumber(v string) *ListGroupsShrinkRequest
	GetPageNumber() *string
	SetPageSize(v string) *ListGroupsShrinkRequest
	GetPageSize() *string
	SetSort(v string) *ListGroupsShrinkRequest
	GetSort() *string
	SetTrafficMode(v string) *ListGroupsShrinkRequest
	GetTrafficMode() *string
	SetWorkspaceId(v string) *ListGroupsShrinkRequest
	GetWorkspaceId() *string
}

type ListGroupsShrinkRequest struct {
	// The filter name. Fuzzy match is supported.
	//
	// example:
	//
	// foo
	Filter *string `json:"Filter,omitempty" xml:"Filter,omitempty"`
	// The user-defined labels.
	LabelsShrink *string `json:"Labels,omitempty" xml:"Labels,omitempty"`
	// The sort order of the results.
	//
	// example:
	//
	// desc
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// The current page number of the service group list. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of service groups to display on each page in a paged query. Settings for paging default to 100.
	//
	// example:
	//
	// 20
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The field by which to sort the results.
	//
	// example:
	//
	// CreateTime
	Sort *string `json:"Sort,omitempty" xml:"Sort,omitempty"`
	// The traffic allocation method used to filter service groups.
	//
	// example:
	//
	// auto
	TrafficMode *string `json:"TrafficMode,omitempty" xml:"TrafficMode,omitempty"`
	// The workspace ID.
	//
	// example:
	//
	// 123***
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListGroupsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListGroupsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListGroupsShrinkRequest) GetFilter() *string {
	return s.Filter
}

func (s *ListGroupsShrinkRequest) GetLabelsShrink() *string {
	return s.LabelsShrink
}

func (s *ListGroupsShrinkRequest) GetOrder() *string {
	return s.Order
}

func (s *ListGroupsShrinkRequest) GetPageNumber() *string {
	return s.PageNumber
}

func (s *ListGroupsShrinkRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *ListGroupsShrinkRequest) GetSort() *string {
	return s.Sort
}

func (s *ListGroupsShrinkRequest) GetTrafficMode() *string {
	return s.TrafficMode
}

func (s *ListGroupsShrinkRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListGroupsShrinkRequest) SetFilter(v string) *ListGroupsShrinkRequest {
	s.Filter = &v
	return s
}

func (s *ListGroupsShrinkRequest) SetLabelsShrink(v string) *ListGroupsShrinkRequest {
	s.LabelsShrink = &v
	return s
}

func (s *ListGroupsShrinkRequest) SetOrder(v string) *ListGroupsShrinkRequest {
	s.Order = &v
	return s
}

func (s *ListGroupsShrinkRequest) SetPageNumber(v string) *ListGroupsShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *ListGroupsShrinkRequest) SetPageSize(v string) *ListGroupsShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListGroupsShrinkRequest) SetSort(v string) *ListGroupsShrinkRequest {
	s.Sort = &v
	return s
}

func (s *ListGroupsShrinkRequest) SetTrafficMode(v string) *ListGroupsShrinkRequest {
	s.TrafficMode = &v
	return s
}

func (s *ListGroupsShrinkRequest) SetWorkspaceId(v string) *ListGroupsShrinkRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ListGroupsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
