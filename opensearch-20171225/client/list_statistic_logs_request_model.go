// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListStatisticLogsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetColumns(v string) *ListStatisticLogsRequest
	GetColumns() *string
	SetDistinct(v bool) *ListStatisticLogsRequest
	GetDistinct() *bool
	SetPageNumber(v int32) *ListStatisticLogsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListStatisticLogsRequest
	GetPageSize() *int32
	SetQuery(v string) *ListStatisticLogsRequest
	GetQuery() *string
	SetSortBy(v string) *ListStatisticLogsRequest
	GetSortBy() *string
	SetStartTime(v int32) *ListStatisticLogsRequest
	GetStartTime() *int32
	SetStopTime(v int32) *ListStatisticLogsRequest
	GetStopTime() *int32
}

type ListStatisticLogsRequest struct {
	// The fields to query. Format: columns=wordsTopPv.
	//
	// For more information, see [Metrics in statistical reports](https://help.aliyun.com/document_detail/187665.html).
	//
	// example:
	//
	// wordsTopPv
	Columns *string `json:"columns,omitempty" xml:"columns,omitempty"`
	// Specifies whether to use the distinct clause.
	//
	// example:
	//
	// true
	Distinct *bool `json:"distinct,omitempty" xml:"distinct,omitempty"`
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The number of entries per page. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The content of the query clause.
	//
	// example:
	//
	// "default:\\"OpenSearch\\""
	Query *string `json:"query,omitempty" xml:"query,omitempty"`
	// The content of the sort clause.
	//
	// example:
	//
	// "-id"
	SortBy *string `json:"sortBy,omitempty" xml:"sortBy,omitempty"`
	// The beginning of the time range to query. The default value is the timestamp of 00:00:00 on the current day.
	//
	// example:
	//
	// 1582214400
	StartTime *int32 `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// The end of the time range to query. The default value is the timestamp of 24:00:00 on the current day.
	//
	// example:
	//
	// 1682222400
	StopTime *int32 `json:"stopTime,omitempty" xml:"stopTime,omitempty"`
}

func (s ListStatisticLogsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListStatisticLogsRequest) GoString() string {
	return s.String()
}

func (s *ListStatisticLogsRequest) GetColumns() *string {
	return s.Columns
}

func (s *ListStatisticLogsRequest) GetDistinct() *bool {
	return s.Distinct
}

func (s *ListStatisticLogsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListStatisticLogsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListStatisticLogsRequest) GetQuery() *string {
	return s.Query
}

func (s *ListStatisticLogsRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *ListStatisticLogsRequest) GetStartTime() *int32 {
	return s.StartTime
}

func (s *ListStatisticLogsRequest) GetStopTime() *int32 {
	return s.StopTime
}

func (s *ListStatisticLogsRequest) SetColumns(v string) *ListStatisticLogsRequest {
	s.Columns = &v
	return s
}

func (s *ListStatisticLogsRequest) SetDistinct(v bool) *ListStatisticLogsRequest {
	s.Distinct = &v
	return s
}

func (s *ListStatisticLogsRequest) SetPageNumber(v int32) *ListStatisticLogsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListStatisticLogsRequest) SetPageSize(v int32) *ListStatisticLogsRequest {
	s.PageSize = &v
	return s
}

func (s *ListStatisticLogsRequest) SetQuery(v string) *ListStatisticLogsRequest {
	s.Query = &v
	return s
}

func (s *ListStatisticLogsRequest) SetSortBy(v string) *ListStatisticLogsRequest {
	s.SortBy = &v
	return s
}

func (s *ListStatisticLogsRequest) SetStartTime(v int32) *ListStatisticLogsRequest {
	s.StartTime = &v
	return s
}

func (s *ListStatisticLogsRequest) SetStopTime(v int32) *ListStatisticLogsRequest {
	s.StopTime = &v
	return s
}

func (s *ListStatisticLogsRequest) Validate() error {
	return dara.Validate(s)
}
