// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRoutineBuildsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageIndex(v int32) *ListRoutineBuildsRequest
	GetPageIndex() *int32
	SetPageSize(v int32) *ListRoutineBuildsRequest
	GetPageSize() *int32
	SetRoutineName(v string) *ListRoutineBuildsRequest
	GetRoutineName() *string
	SetSortBy(v string) *ListRoutineBuildsRequest
	GetSortBy() *string
	SetSortOrder(v string) *ListRoutineBuildsRequest
	GetSortOrder() *string
	SetStatus(v string) *ListRoutineBuildsRequest
	GetStatus() *string
}

type ListRoutineBuildsRequest struct {
	// The page number for a paged query. The value must be greater than or equal to 1.
	//
	// example:
	//
	// 1
	PageIndex *int32 `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	// The number of entries per page for a paged query. Valid values: 1 to 500.
	//
	// example:
	//
	// 500
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ER name.
	//
	// example:
	//
	// test-routine
	RoutineName *string `json:"RoutineName,omitempty" xml:"RoutineName,omitempty"`
	// The field used for sorting. By default, results are sorted by purchase time. Valid values:
	//
	// - CreateTime: purchase time.
	//
	// - ExpireTime: expiration time.
	//
	// example:
	//
	// CreateTime
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// The sort order. Default value: desc. Valid values:
	//
	// - asc: ascending order.
	//
	// - desc: descending order.
	//
	// example:
	//
	// asc
	SortOrder *string `json:"SortOrder,omitempty" xml:"SortOrder,omitempty"`
	// The status of the build task. Valid values:
	//
	// - int: initialization
	//
	// - pending: preparing
	//
	// - building: building
	//
	// - succeed: build succeeded
	//
	// - failed: build failed
	//
	// - canceled: canceled
	//
	// example:
	//
	// canceled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListRoutineBuildsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListRoutineBuildsRequest) GoString() string {
	return s.String()
}

func (s *ListRoutineBuildsRequest) GetPageIndex() *int32 {
	return s.PageIndex
}

func (s *ListRoutineBuildsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListRoutineBuildsRequest) GetRoutineName() *string {
	return s.RoutineName
}

func (s *ListRoutineBuildsRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *ListRoutineBuildsRequest) GetSortOrder() *string {
	return s.SortOrder
}

func (s *ListRoutineBuildsRequest) GetStatus() *string {
	return s.Status
}

func (s *ListRoutineBuildsRequest) SetPageIndex(v int32) *ListRoutineBuildsRequest {
	s.PageIndex = &v
	return s
}

func (s *ListRoutineBuildsRequest) SetPageSize(v int32) *ListRoutineBuildsRequest {
	s.PageSize = &v
	return s
}

func (s *ListRoutineBuildsRequest) SetRoutineName(v string) *ListRoutineBuildsRequest {
	s.RoutineName = &v
	return s
}

func (s *ListRoutineBuildsRequest) SetSortBy(v string) *ListRoutineBuildsRequest {
	s.SortBy = &v
	return s
}

func (s *ListRoutineBuildsRequest) SetSortOrder(v string) *ListRoutineBuildsRequest {
	s.SortOrder = &v
	return s
}

func (s *ListRoutineBuildsRequest) SetStatus(v string) *ListRoutineBuildsRequest {
	s.Status = &v
	return s
}

func (s *ListRoutineBuildsRequest) Validate() error {
	return dara.Validate(s)
}
