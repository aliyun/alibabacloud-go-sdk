// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSessionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLabels(v string) *ListSessionsRequest
	GetLabels() *string
	SetOrder(v string) *ListSessionsRequest
	GetOrder() *string
	SetPageNumber(v int64) *ListSessionsRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *ListSessionsRequest
	GetPageSize() *int64
	SetSortBy(v string) *ListSessionsRequest
	GetSortBy() *string
}

type ListSessionsRequest struct {
	// example:
	//
	// system.product.type=DLC
	Labels *string `json:"Labels,omitempty" xml:"Labels,omitempty"`
	// example:
	//
	// ASC
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 20
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// GmtModified
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
}

func (s ListSessionsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSessionsRequest) GoString() string {
	return s.String()
}

func (s *ListSessionsRequest) GetLabels() *string {
	return s.Labels
}

func (s *ListSessionsRequest) GetOrder() *string {
	return s.Order
}

func (s *ListSessionsRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListSessionsRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListSessionsRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *ListSessionsRequest) SetLabels(v string) *ListSessionsRequest {
	s.Labels = &v
	return s
}

func (s *ListSessionsRequest) SetOrder(v string) *ListSessionsRequest {
	s.Order = &v
	return s
}

func (s *ListSessionsRequest) SetPageNumber(v int64) *ListSessionsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListSessionsRequest) SetPageSize(v int64) *ListSessionsRequest {
	s.PageSize = &v
	return s
}

func (s *ListSessionsRequest) SetSortBy(v string) *ListSessionsRequest {
	s.SortBy = &v
	return s
}

func (s *ListSessionsRequest) Validate() error {
	return dara.Validate(s)
}
