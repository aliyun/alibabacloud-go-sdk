// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTasksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *ListTasksRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListTasksRequest
	GetPageSize() *int32
	SetProvisionedProductId(v string) *ListTasksRequest
	GetProvisionedProductId() *string
	SetSortBy(v string) *ListTasksRequest
	GetSortBy() *string
	SetSortOrder(v string) *ListTasksRequest
	GetSortOrder() *string
}

type ListTasksRequest struct {
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// pp-bp1ddg1n2a****
	ProvisionedProductId *string `json:"ProvisionedProductId,omitempty" xml:"ProvisionedProductId,omitempty"`
	// example:
	//
	// CreateTime
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// example:
	//
	// Desc
	SortOrder *string `json:"SortOrder,omitempty" xml:"SortOrder,omitempty"`
}

func (s ListTasksRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTasksRequest) GoString() string {
	return s.String()
}

func (s *ListTasksRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListTasksRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListTasksRequest) GetProvisionedProductId() *string {
	return s.ProvisionedProductId
}

func (s *ListTasksRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *ListTasksRequest) GetSortOrder() *string {
	return s.SortOrder
}

func (s *ListTasksRequest) SetPageNumber(v int32) *ListTasksRequest {
	s.PageNumber = &v
	return s
}

func (s *ListTasksRequest) SetPageSize(v int32) *ListTasksRequest {
	s.PageSize = &v
	return s
}

func (s *ListTasksRequest) SetProvisionedProductId(v string) *ListTasksRequest {
	s.ProvisionedProductId = &v
	return s
}

func (s *ListTasksRequest) SetSortBy(v string) *ListTasksRequest {
	s.SortBy = &v
	return s
}

func (s *ListTasksRequest) SetSortOrder(v string) *ListTasksRequest {
	s.SortOrder = &v
	return s
}

func (s *ListTasksRequest) Validate() error {
	return dara.Validate(s)
}
