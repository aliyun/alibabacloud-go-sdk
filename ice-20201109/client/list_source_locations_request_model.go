// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSourceLocationsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFilterState(v bool) *ListSourceLocationsRequest
	GetFilterState() *bool
	SetPageNo(v int32) *ListSourceLocationsRequest
	GetPageNo() *int32
	SetPageSize(v int32) *ListSourceLocationsRequest
	GetPageSize() *int32
	SetSortBy(v string) *ListSourceLocationsRequest
	GetSortBy() *string
	SetSortByModifiedTime(v string) *ListSourceLocationsRequest
	GetSortByModifiedTime() *string
	SetSourceLocationName(v string) *ListSourceLocationsRequest
	GetSourceLocationName() *string
}

type ListSourceLocationsRequest struct {
	// Specifies whether to ignore source locations marked as deleted. A value of true means ignoring source locations marked as deleted.
	//
	// example:
	//
	// true
	FilterState *bool `json:"FilterState,omitempty" xml:"FilterState,omitempty"`
	// 	- The page number.
	//
	// 	- Default value: 1.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page. Default value: 20. Valid values: 1 to 100.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The sorting order. By default, the query results are sorted by creation time in descending order.
	//
	// example:
	//
	// desc
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// The sorting order of the source locations based on the time when they were last modified.
	//
	// example:
	//
	// desc
	SortByModifiedTime *string `json:"SortByModifiedTime,omitempty" xml:"SortByModifiedTime,omitempty"`
	// The name of the source location.
	//
	// example:
	//
	// MySourceLocation
	SourceLocationName *string `json:"SourceLocationName,omitempty" xml:"SourceLocationName,omitempty"`
}

func (s ListSourceLocationsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSourceLocationsRequest) GoString() string {
	return s.String()
}

func (s *ListSourceLocationsRequest) GetFilterState() *bool {
	return s.FilterState
}

func (s *ListSourceLocationsRequest) GetPageNo() *int32 {
	return s.PageNo
}

func (s *ListSourceLocationsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListSourceLocationsRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *ListSourceLocationsRequest) GetSortByModifiedTime() *string {
	return s.SortByModifiedTime
}

func (s *ListSourceLocationsRequest) GetSourceLocationName() *string {
	return s.SourceLocationName
}

func (s *ListSourceLocationsRequest) SetFilterState(v bool) *ListSourceLocationsRequest {
	s.FilterState = &v
	return s
}

func (s *ListSourceLocationsRequest) SetPageNo(v int32) *ListSourceLocationsRequest {
	s.PageNo = &v
	return s
}

func (s *ListSourceLocationsRequest) SetPageSize(v int32) *ListSourceLocationsRequest {
	s.PageSize = &v
	return s
}

func (s *ListSourceLocationsRequest) SetSortBy(v string) *ListSourceLocationsRequest {
	s.SortBy = &v
	return s
}

func (s *ListSourceLocationsRequest) SetSortByModifiedTime(v string) *ListSourceLocationsRequest {
	s.SortByModifiedTime = &v
	return s
}

func (s *ListSourceLocationsRequest) SetSourceLocationName(v string) *ListSourceLocationsRequest {
	s.SourceLocationName = &v
	return s
}

func (s *ListSourceLocationsRequest) Validate() error {
	return dara.Validate(s)
}
