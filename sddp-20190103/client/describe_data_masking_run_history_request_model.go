// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDataMaskingRunHistoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *DescribeDataMaskingRunHistoryRequest
	GetCurrentPage() *int32
	SetDstType(v int32) *DescribeDataMaskingRunHistoryRequest
	GetDstType() *int32
	SetEndTime(v int64) *DescribeDataMaskingRunHistoryRequest
	GetEndTime() *int64
	SetLang(v string) *DescribeDataMaskingRunHistoryRequest
	GetLang() *string
	SetMainProcessId(v int64) *DescribeDataMaskingRunHistoryRequest
	GetMainProcessId() *int64
	SetPageSize(v int32) *DescribeDataMaskingRunHistoryRequest
	GetPageSize() *int32
	SetSrcTableName(v string) *DescribeDataMaskingRunHistoryRequest
	GetSrcTableName() *string
	SetSrcType(v int32) *DescribeDataMaskingRunHistoryRequest
	GetSrcType() *int32
	SetStartTime(v int64) *DescribeDataMaskingRunHistoryRequest
	GetStartTime() *int64
	SetStatus(v int32) *DescribeDataMaskingRunHistoryRequest
	GetStatus() *int32
	SetTaskId(v string) *DescribeDataMaskingRunHistoryRequest
	GetTaskId() *string
}

type DescribeDataMaskingRunHistoryRequest struct {
	// The page number to return.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The type of service to which the masked data is destined. Valid values: **1*	- for MaxCompute, **2*	- for OSS, **3*	- for ADS, **4*	- for OTS, and **5*	- for RDS.
	//
	// example:
	//
	// 2
	DstType *int32 `json:"DstType,omitempty" xml:"DstType,omitempty"`
	// The end time to query for task executions. This is a UNIX timestamp in milliseconds.
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
	// The ID of the main task.
	//
	// > If a task has subtasks, this parameter specifies the ID of the main task. Otherwise, this parameter is empty.
	//
	// example:
	//
	// 366731
	MainProcessId *int64 `json:"MainProcessId,omitempty" xml:"MainProcessId,omitempty"`
	// The number of entries to return on each page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The name of the source table.
	//
	// example:
	//
	// add
	SrcTableName *string `json:"SrcTableName,omitempty" xml:"SrcTableName,omitempty"`
	// The type of service to which the source data belongs. Valid values: **1*	- for MaxCompute, **2*	- for OSS, **3*	- for ADS, **4*	- for OTS, and **5*	- for RDS.
	//
	// example:
	//
	// 2
	SrcType *int32 `json:"SrcType,omitempty" xml:"SrcType,omitempty"`
	// The start time to query for task executions. This is a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1582992000000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The execution status of the task. Valid values:
	//
	// - **-1**: pending.
	//
	// - **0**: running.
	//
	// - **1**: successful.
	//
	// - **2**: failed.
	//
	// - **3**: stopped by user.
	//
	// - **4**: partially failed.
	//
	// example:
	//
	// 0
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the data masking task.
	//
	// example:
	//
	// mt4HBgtw1B******
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DescribeDataMaskingRunHistoryRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDataMaskingRunHistoryRequest) GoString() string {
	return s.String()
}

func (s *DescribeDataMaskingRunHistoryRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeDataMaskingRunHistoryRequest) GetDstType() *int32 {
	return s.DstType
}

func (s *DescribeDataMaskingRunHistoryRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeDataMaskingRunHistoryRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeDataMaskingRunHistoryRequest) GetMainProcessId() *int64 {
	return s.MainProcessId
}

func (s *DescribeDataMaskingRunHistoryRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeDataMaskingRunHistoryRequest) GetSrcTableName() *string {
	return s.SrcTableName
}

func (s *DescribeDataMaskingRunHistoryRequest) GetSrcType() *int32 {
	return s.SrcType
}

func (s *DescribeDataMaskingRunHistoryRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeDataMaskingRunHistoryRequest) GetStatus() *int32 {
	return s.Status
}

func (s *DescribeDataMaskingRunHistoryRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *DescribeDataMaskingRunHistoryRequest) SetCurrentPage(v int32) *DescribeDataMaskingRunHistoryRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeDataMaskingRunHistoryRequest) SetDstType(v int32) *DescribeDataMaskingRunHistoryRequest {
	s.DstType = &v
	return s
}

func (s *DescribeDataMaskingRunHistoryRequest) SetEndTime(v int64) *DescribeDataMaskingRunHistoryRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDataMaskingRunHistoryRequest) SetLang(v string) *DescribeDataMaskingRunHistoryRequest {
	s.Lang = &v
	return s
}

func (s *DescribeDataMaskingRunHistoryRequest) SetMainProcessId(v int64) *DescribeDataMaskingRunHistoryRequest {
	s.MainProcessId = &v
	return s
}

func (s *DescribeDataMaskingRunHistoryRequest) SetPageSize(v int32) *DescribeDataMaskingRunHistoryRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDataMaskingRunHistoryRequest) SetSrcTableName(v string) *DescribeDataMaskingRunHistoryRequest {
	s.SrcTableName = &v
	return s
}

func (s *DescribeDataMaskingRunHistoryRequest) SetSrcType(v int32) *DescribeDataMaskingRunHistoryRequest {
	s.SrcType = &v
	return s
}

func (s *DescribeDataMaskingRunHistoryRequest) SetStartTime(v int64) *DescribeDataMaskingRunHistoryRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDataMaskingRunHistoryRequest) SetStatus(v int32) *DescribeDataMaskingRunHistoryRequest {
	s.Status = &v
	return s
}

func (s *DescribeDataMaskingRunHistoryRequest) SetTaskId(v string) *DescribeDataMaskingRunHistoryRequest {
	s.TaskId = &v
	return s
}

func (s *DescribeDataMaskingRunHistoryRequest) Validate() error {
	return dara.Validate(s)
}
