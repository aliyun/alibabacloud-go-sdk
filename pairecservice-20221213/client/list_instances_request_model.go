// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListInstancesRequest
	GetInstanceId() *string
	SetOrder(v string) *ListInstancesRequest
	GetOrder() *string
	SetPageNumber(v int32) *ListInstancesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListInstancesRequest
	GetPageSize() *int32
	SetSortBy(v string) *ListInstancesRequest
	GetSortBy() *string
	SetType(v string) *ListInstancesRequest
	GetType() *string
}

type ListInstancesRequest struct {
	// The ID of the instance. You can use this parameter to perform a fuzzy search for instances.
	//
	// example:
	//
	// pairec-test1
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The sort order. Valid values: `Asc` (ascending) and `Desc` (descending).
	//
	// example:
	//
	// Desc
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The page number.
	//
	// example:
	//
	// 50
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The field by which to sort the results.
	//
	// example:
	//
	// Type
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// The instance type. Valid values:
	//
	// - `basic`: Basic edition
	//
	// - `high-level`: High-level edition
	//
	// - `advanced`: Advanced edition
	//
	// - `standard`: Standard edition
	//
	// example:
	//
	// basic
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListInstancesRequest) GoString() string {
	return s.String()
}

func (s *ListInstancesRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListInstancesRequest) GetOrder() *string {
	return s.Order
}

func (s *ListInstancesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListInstancesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListInstancesRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *ListInstancesRequest) GetType() *string {
	return s.Type
}

func (s *ListInstancesRequest) SetInstanceId(v string) *ListInstancesRequest {
	s.InstanceId = &v
	return s
}

func (s *ListInstancesRequest) SetOrder(v string) *ListInstancesRequest {
	s.Order = &v
	return s
}

func (s *ListInstancesRequest) SetPageNumber(v int32) *ListInstancesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListInstancesRequest) SetPageSize(v int32) *ListInstancesRequest {
	s.PageSize = &v
	return s
}

func (s *ListInstancesRequest) SetSortBy(v string) *ListInstancesRequest {
	s.SortBy = &v
	return s
}

func (s *ListInstancesRequest) SetType(v string) *ListInstancesRequest {
	s.Type = &v
	return s
}

func (s *ListInstancesRequest) Validate() error {
	return dara.Validate(s)
}
