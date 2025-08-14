// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSourcesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFilterState(v bool) *ListSourcesRequest
	GetFilterState() *bool
	SetPageNo(v string) *ListSourcesRequest
	GetPageNo() *string
	SetPageSize(v string) *ListSourcesRequest
	GetPageSize() *string
	SetSortBy(v string) *ListSourcesRequest
	GetSortBy() *string
	SetSortByModifiedTime(v string) *ListSourcesRequest
	GetSortByModifiedTime() *string
	SetSourceLocationName(v string) *ListSourcesRequest
	GetSourceLocationName() *string
	SetSourceName(v string) *ListSourcesRequest
	GetSourceName() *string
	SetSourceType(v string) *ListSourcesRequest
	GetSourceType() *string
}

type ListSourcesRequest struct {
	// Specifies whether to ignore sources marked as deleted.
	//
	// example:
	//
	// true
	FilterState *bool `json:"FilterState,omitempty" xml:"FilterState,omitempty"`
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNo *string `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The sorting order. By default, the query results are sorted by creation time in descending order. Valid values: asc and desc.
	//
	// example:
	//
	// asc
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// The sorting order by modification time. Valid values: asc and desc.
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
	// The name of the source.
	//
	// example:
	//
	// MyVodSource
	SourceName *string `json:"SourceName,omitempty" xml:"SourceName,omitempty"`
	// The source type. Valid values: vodSource and liveSource.
	//
	// example:
	//
	// vodSource
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
}

func (s ListSourcesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSourcesRequest) GoString() string {
	return s.String()
}

func (s *ListSourcesRequest) GetFilterState() *bool {
	return s.FilterState
}

func (s *ListSourcesRequest) GetPageNo() *string {
	return s.PageNo
}

func (s *ListSourcesRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *ListSourcesRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *ListSourcesRequest) GetSortByModifiedTime() *string {
	return s.SortByModifiedTime
}

func (s *ListSourcesRequest) GetSourceLocationName() *string {
	return s.SourceLocationName
}

func (s *ListSourcesRequest) GetSourceName() *string {
	return s.SourceName
}

func (s *ListSourcesRequest) GetSourceType() *string {
	return s.SourceType
}

func (s *ListSourcesRequest) SetFilterState(v bool) *ListSourcesRequest {
	s.FilterState = &v
	return s
}

func (s *ListSourcesRequest) SetPageNo(v string) *ListSourcesRequest {
	s.PageNo = &v
	return s
}

func (s *ListSourcesRequest) SetPageSize(v string) *ListSourcesRequest {
	s.PageSize = &v
	return s
}

func (s *ListSourcesRequest) SetSortBy(v string) *ListSourcesRequest {
	s.SortBy = &v
	return s
}

func (s *ListSourcesRequest) SetSortByModifiedTime(v string) *ListSourcesRequest {
	s.SortByModifiedTime = &v
	return s
}

func (s *ListSourcesRequest) SetSourceLocationName(v string) *ListSourcesRequest {
	s.SourceLocationName = &v
	return s
}

func (s *ListSourcesRequest) SetSourceName(v string) *ListSourcesRequest {
	s.SourceName = &v
	return s
}

func (s *ListSourcesRequest) SetSourceType(v string) *ListSourcesRequest {
	s.SourceType = &v
	return s
}

func (s *ListSourcesRequest) Validate() error {
	return dara.Validate(s)
}
