// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDataMaskingTasksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *DescribeDataMaskingTasksRequest
	GetCurrentPage() *int32
	SetDstType(v int32) *DescribeDataMaskingTasksRequest
	GetDstType() *int32
	SetEndTime(v int64) *DescribeDataMaskingTasksRequest
	GetEndTime() *int64
	SetLang(v string) *DescribeDataMaskingTasksRequest
	GetLang() *string
	SetPageSize(v int32) *DescribeDataMaskingTasksRequest
	GetPageSize() *int32
	SetSearchKey(v string) *DescribeDataMaskingTasksRequest
	GetSearchKey() *string
	SetStartTime(v int64) *DescribeDataMaskingTasksRequest
	GetStartTime() *int64
}

type DescribeDataMaskingTasksRequest struct {
	// The page number of the page to return.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The service to which the data to be de-identified belongs. Valid values include **1**, **2**, **3**, **4**, and **5**. The value 1 indicates MaxCompute. The value 2 indicates Object Storage Service (OSS). The value 3 indicates AnalyticDB for MySQL. The value 4 indicates Tablestore. The value 5 indicates ApsaraDB RDS.
	//
	// example:
	//
	// 2
	DstType *int32 `json:"DstType,omitempty" xml:"DstType,omitempty"`
	// The end of the time range during which the de-identification tasks to be queried are created. The value is a UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1583856000000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The language of the content within the request and response. Default value: **zh_cn**. Valid values:
	//
	// 	- **zh_cn**: Chinese
	//
	// 	- **en_us**: English
	//
	// example:
	//
	// zh_cn
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The number of entries to return on each page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The keyword used to query the de-identification tasks, which can be the task name or ID.
	//
	// example:
	//
	// test
	SearchKey *string `json:"SearchKey,omitempty" xml:"SearchKey,omitempty"`
	// The beginning of the time range during which the de-identification tasks to be queried are created. The value is a UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1582992000000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDataMaskingTasksRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDataMaskingTasksRequest) GoString() string {
	return s.String()
}

func (s *DescribeDataMaskingTasksRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeDataMaskingTasksRequest) GetDstType() *int32 {
	return s.DstType
}

func (s *DescribeDataMaskingTasksRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeDataMaskingTasksRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeDataMaskingTasksRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeDataMaskingTasksRequest) GetSearchKey() *string {
	return s.SearchKey
}

func (s *DescribeDataMaskingTasksRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeDataMaskingTasksRequest) SetCurrentPage(v int32) *DescribeDataMaskingTasksRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeDataMaskingTasksRequest) SetDstType(v int32) *DescribeDataMaskingTasksRequest {
	s.DstType = &v
	return s
}

func (s *DescribeDataMaskingTasksRequest) SetEndTime(v int64) *DescribeDataMaskingTasksRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDataMaskingTasksRequest) SetLang(v string) *DescribeDataMaskingTasksRequest {
	s.Lang = &v
	return s
}

func (s *DescribeDataMaskingTasksRequest) SetPageSize(v int32) *DescribeDataMaskingTasksRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDataMaskingTasksRequest) SetSearchKey(v string) *DescribeDataMaskingTasksRequest {
	s.SearchKey = &v
	return s
}

func (s *DescribeDataMaskingTasksRequest) SetStartTime(v int64) *DescribeDataMaskingTasksRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDataMaskingTasksRequest) Validate() error {
	return dara.Validate(s)
}
