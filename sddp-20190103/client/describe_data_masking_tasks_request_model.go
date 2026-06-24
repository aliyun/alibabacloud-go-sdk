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
	// The page number to return.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The product that the destination data source belongs to. Valid values:
	//
	// - **1**: MaxCompute.
	//
	// - **2**: OSS.
	//
	// - **3**: ADS.
	//
	// - **4**: OTS.
	//
	// - **5**: RDS.
	//
	// - **6**: SELF_DB.
	//
	// example:
	//
	// 2
	DstType *int32 `json:"DstType,omitempty" xml:"DstType,omitempty"`
	// The end time for creating the data masking task. The value is a UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1583856000000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The language of the request and response. Default value: **zh_cn**. Valid values:
	//
	// - **zh_cn**: Chinese.
	//
	// - **en_us**: English.
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
	// A keyword to search for tasks. You can search by task name or task ID.
	//
	// example:
	//
	// test
	SearchKey *string `json:"SearchKey,omitempty" xml:"SearchKey,omitempty"`
	// The start time for creating the task. The value is a UNIX timestamp. Unit: milliseconds.
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
