// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstanceSnapshotRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOrder(v string) *ListInstanceSnapshotRequest
	GetOrder() *string
	SetPageNumber(v int64) *ListInstanceSnapshotRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *ListInstanceSnapshotRequest
	GetPageSize() *int64
	SetSortBy(v string) *ListInstanceSnapshotRequest
	GetSortBy() *string
}

type ListInstanceSnapshotRequest struct {
	// example:
	//
	// DESC
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// gmtCreate
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
}

func (s ListInstanceSnapshotRequest) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceSnapshotRequest) GoString() string {
	return s.String()
}

func (s *ListInstanceSnapshotRequest) GetOrder() *string {
	return s.Order
}

func (s *ListInstanceSnapshotRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListInstanceSnapshotRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListInstanceSnapshotRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *ListInstanceSnapshotRequest) SetOrder(v string) *ListInstanceSnapshotRequest {
	s.Order = &v
	return s
}

func (s *ListInstanceSnapshotRequest) SetPageNumber(v int64) *ListInstanceSnapshotRequest {
	s.PageNumber = &v
	return s
}

func (s *ListInstanceSnapshotRequest) SetPageSize(v int64) *ListInstanceSnapshotRequest {
	s.PageSize = &v
	return s
}

func (s *ListInstanceSnapshotRequest) SetSortBy(v string) *ListInstanceSnapshotRequest {
	s.SortBy = &v
	return s
}

func (s *ListInstanceSnapshotRequest) Validate() error {
	return dara.Validate(s)
}
