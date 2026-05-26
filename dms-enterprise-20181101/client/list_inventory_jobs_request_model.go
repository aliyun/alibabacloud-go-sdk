// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInventoryJobsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOffset(v int32) *ListInventoryJobsRequest
	GetOffset() *int32
	SetQuery(v string) *ListInventoryJobsRequest
	GetQuery() *string
	SetSize(v int32) *ListInventoryJobsRequest
	GetSize() *int32
	SetSortBy(v string) *ListInventoryJobsRequest
	GetSortBy() *string
	SetSortOrder(v string) *ListInventoryJobsRequest
	GetSortOrder() *string
	SetStatus(v int32) *ListInventoryJobsRequest
	GetStatus() *int32
}

type ListInventoryJobsRequest struct {
	// example:
	//
	// 0
	Offset *int32 `json:"Offset,omitempty" xml:"Offset,omitempty"`
	// example:
	//
	// task_1001
	Query *string `json:"Query,omitempty" xml:"Query,omitempty"`
	// example:
	//
	// 20
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
	// example:
	//
	// confidence
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// example:
	//
	// desc
	SortOrder *string `json:"SortOrder,omitempty" xml:"SortOrder,omitempty"`
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListInventoryJobsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListInventoryJobsRequest) GoString() string {
	return s.String()
}

func (s *ListInventoryJobsRequest) GetOffset() *int32 {
	return s.Offset
}

func (s *ListInventoryJobsRequest) GetQuery() *string {
	return s.Query
}

func (s *ListInventoryJobsRequest) GetSize() *int32 {
	return s.Size
}

func (s *ListInventoryJobsRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *ListInventoryJobsRequest) GetSortOrder() *string {
	return s.SortOrder
}

func (s *ListInventoryJobsRequest) GetStatus() *int32 {
	return s.Status
}

func (s *ListInventoryJobsRequest) SetOffset(v int32) *ListInventoryJobsRequest {
	s.Offset = &v
	return s
}

func (s *ListInventoryJobsRequest) SetQuery(v string) *ListInventoryJobsRequest {
	s.Query = &v
	return s
}

func (s *ListInventoryJobsRequest) SetSize(v int32) *ListInventoryJobsRequest {
	s.Size = &v
	return s
}

func (s *ListInventoryJobsRequest) SetSortBy(v string) *ListInventoryJobsRequest {
	s.SortBy = &v
	return s
}

func (s *ListInventoryJobsRequest) SetSortOrder(v string) *ListInventoryJobsRequest {
	s.SortOrder = &v
	return s
}

func (s *ListInventoryJobsRequest) SetStatus(v int32) *ListInventoryJobsRequest {
	s.Status = &v
	return s
}

func (s *ListInventoryJobsRequest) Validate() error {
	return dara.Validate(s)
}
