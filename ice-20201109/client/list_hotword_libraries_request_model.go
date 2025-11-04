// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHotwordLibrariesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *ListHotwordLibrariesRequest
	GetEndTime() *string
	SetMaxResults(v int32) *ListHotwordLibrariesRequest
	GetMaxResults() *int32
	SetName(v string) *ListHotwordLibrariesRequest
	GetName() *string
	SetNextToken(v string) *ListHotwordLibrariesRequest
	GetNextToken() *string
	SetPageNo(v int64) *ListHotwordLibrariesRequest
	GetPageNo() *int64
	SetPageSize(v int64) *ListHotwordLibrariesRequest
	GetPageSize() *int64
	SetSortBy(v string) *ListHotwordLibrariesRequest
	GetSortBy() *string
	SetStartTime(v string) *ListHotwordLibrariesRequest
	GetStartTime() *string
	SetUsageScenario(v string) *ListHotwordLibrariesRequest
	GetUsageScenario() *string
}

type ListHotwordLibrariesRequest struct {
	// The end of the time range to query.
	//
	// example:
	//
	// 2020-12-26T04:11:10Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The maximum number of entries to return.
	//
	// Default value: 10. Valid values: 1 to 100.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The name of the hotword library.
	//
	// example:
	//
	// 热词库使用名称
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// ****73f33c91-d59383e8280b****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNo *int64 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page. Default value: 10. Valid values: 1 to 100.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The sorting order. By default, the query results are sorted by creation time in descending order.
	//
	// example:
	//
	// CreationTime:Desc
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// The beginning of the time range to query.
	//
	// example:
	//
	// 2020-12-26T04:11:10Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The usage scenario of the hotword library. Valid values:
	//
	// 	- ASR: Automatic Speech Recognition
	//
	// 	- StructuredMediaAssets: structured media analysis
	//
	// 	- VideoTranslation: Video translation. This field cannot be modified after the hotword library is created.
	//
	// example:
	//
	// ASR
	UsageScenario *string `json:"UsageScenario,omitempty" xml:"UsageScenario,omitempty"`
}

func (s ListHotwordLibrariesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListHotwordLibrariesRequest) GoString() string {
	return s.String()
}

func (s *ListHotwordLibrariesRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *ListHotwordLibrariesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListHotwordLibrariesRequest) GetName() *string {
	return s.Name
}

func (s *ListHotwordLibrariesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListHotwordLibrariesRequest) GetPageNo() *int64 {
	return s.PageNo
}

func (s *ListHotwordLibrariesRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListHotwordLibrariesRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *ListHotwordLibrariesRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *ListHotwordLibrariesRequest) GetUsageScenario() *string {
	return s.UsageScenario
}

func (s *ListHotwordLibrariesRequest) SetEndTime(v string) *ListHotwordLibrariesRequest {
	s.EndTime = &v
	return s
}

func (s *ListHotwordLibrariesRequest) SetMaxResults(v int32) *ListHotwordLibrariesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListHotwordLibrariesRequest) SetName(v string) *ListHotwordLibrariesRequest {
	s.Name = &v
	return s
}

func (s *ListHotwordLibrariesRequest) SetNextToken(v string) *ListHotwordLibrariesRequest {
	s.NextToken = &v
	return s
}

func (s *ListHotwordLibrariesRequest) SetPageNo(v int64) *ListHotwordLibrariesRequest {
	s.PageNo = &v
	return s
}

func (s *ListHotwordLibrariesRequest) SetPageSize(v int64) *ListHotwordLibrariesRequest {
	s.PageSize = &v
	return s
}

func (s *ListHotwordLibrariesRequest) SetSortBy(v string) *ListHotwordLibrariesRequest {
	s.SortBy = &v
	return s
}

func (s *ListHotwordLibrariesRequest) SetStartTime(v string) *ListHotwordLibrariesRequest {
	s.StartTime = &v
	return s
}

func (s *ListHotwordLibrariesRequest) SetUsageScenario(v string) *ListHotwordLibrariesRequest {
	s.UsageScenario = &v
	return s
}

func (s *ListHotwordLibrariesRequest) Validate() error {
	return dara.Validate(s)
}
