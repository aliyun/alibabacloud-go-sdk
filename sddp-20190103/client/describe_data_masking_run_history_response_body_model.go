// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDataMaskingRunHistoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *DescribeDataMaskingRunHistoryResponseBody
	GetCurrentPage() *int32
	SetItems(v []*DescribeDataMaskingRunHistoryResponseBodyItems) *DescribeDataMaskingRunHistoryResponseBody
	GetItems() []*DescribeDataMaskingRunHistoryResponseBodyItems
	SetPageSize(v int32) *DescribeDataMaskingRunHistoryResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeDataMaskingRunHistoryResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeDataMaskingRunHistoryResponseBody
	GetTotalCount() *int32
}

type DescribeDataMaskingRunHistoryResponseBody struct {
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// A list of data masking task details.
	Items []*DescribeDataMaskingRunHistoryResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 769FB3C1-F4C9-4******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeDataMaskingRunHistoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDataMaskingRunHistoryResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDataMaskingRunHistoryResponseBody) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeDataMaskingRunHistoryResponseBody) GetItems() []*DescribeDataMaskingRunHistoryResponseBodyItems {
	return s.Items
}

func (s *DescribeDataMaskingRunHistoryResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeDataMaskingRunHistoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDataMaskingRunHistoryResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeDataMaskingRunHistoryResponseBody) SetCurrentPage(v int32) *DescribeDataMaskingRunHistoryResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *DescribeDataMaskingRunHistoryResponseBody) SetItems(v []*DescribeDataMaskingRunHistoryResponseBodyItems) *DescribeDataMaskingRunHistoryResponseBody {
	s.Items = v
	return s
}

func (s *DescribeDataMaskingRunHistoryResponseBody) SetPageSize(v int32) *DescribeDataMaskingRunHistoryResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeDataMaskingRunHistoryResponseBody) SetRequestId(v string) *DescribeDataMaskingRunHistoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDataMaskingRunHistoryResponseBody) SetTotalCount(v int32) *DescribeDataMaskingRunHistoryResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeDataMaskingRunHistoryResponseBody) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDataMaskingRunHistoryResponseBodyItems struct {
	// The number of data conflicts. This is the number of rows to be inserted into the destination table that conflict with existing data.
	//
	// example:
	//
	// 0
	ConflictCount *int64 `json:"ConflictCount,omitempty" xml:"ConflictCount,omitempty"`
	// The type of service to which the masked data is destined. Valid values: **1*	- for MaxCompute, **2*	- for OSS, **3*	- for ADS, **4*	- for OTS, and **5*	- for RDS.
	//
	// example:
	//
	// 2
	DstType *int32 `json:"DstType,omitempty" xml:"DstType,omitempty"`
	// The type of the destination service. Valid values include **MaxCompute, OSS, ADS, OTS, and RDS**.
	//
	// example:
	//
	// OSS
	DstTypeCode *string `json:"DstTypeCode,omitempty" xml:"DstTypeCode,omitempty"`
	// The time when the execution ended. This is a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1582251233000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The error code returned when the task fails. This parameter has a value only if the task fails.
	//
	// example:
	//
	// masking_task_not_found
	FailCode *string `json:"FailCode,omitempty" xml:"FailCode,omitempty"`
	// The reason the task failed.
	//
	// example:
	//
	// error
	FailMsg *string `json:"FailMsg,omitempty" xml:"FailMsg,omitempty"`
	// Indicates whether a download file is available.
	//
	// - **1**: Yes.
	//
	// - **0**: No.
	//
	// example:
	//
	// 1
	HasDownloadFile *int32 `json:"HasDownloadFile,omitempty" xml:"HasDownloadFile,omitempty"`
	// The number of created subtasks.
	//
	// example:
	//
	// 4
	HasSubProcess *int32 `json:"HasSubProcess,omitempty" xml:"HasSubProcess,omitempty"`
	// The ID of the execution record.
	//
	// example:
	//
	// 1
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The number of masked rows.
	//
	// example:
	//
	// 100
	MaskingCount *int64 `json:"MaskingCount,omitempty" xml:"MaskingCount,omitempty"`
	// The execution progress.
	//
	// example:
	//
	// 100
	Percentage *int32 `json:"Percentage,omitempty" xml:"Percentage,omitempty"`
	// The number of times the task has been executed.
	//
	// example:
	//
	// 1
	RunIndex *int32 `json:"RunIndex,omitempty" xml:"RunIndex,omitempty"`
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
	// The type of the source service. Valid values include **MaxCompute, OSS, ADS, OTS, and RDS**.
	//
	// example:
	//
	// OSS
	SrcTypeCode *string `json:"SrcTypeCode,omitempty" xml:"SrcTypeCode,omitempty"`
	// The time when the execution started. This is a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1582251233000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The execution status of the task. Valid values:
	//
	// - -**1**: pending.
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
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the task.
	//
	// example:
	//
	// mt4HBgtw1B******
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The execution method. Valid values:
	//
	// - **1**: manual.
	//
	// - **2**: scheduled.
	//
	// example:
	//
	// 1
	Type *int32 `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeDataMaskingRunHistoryResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeDataMaskingRunHistoryResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeDataMaskingRunHistoryResponseBodyItems) GetConflictCount() *int64 {
	return s.ConflictCount
}

func (s *DescribeDataMaskingRunHistoryResponseBodyItems) GetDstType() *int32 {
	return s.DstType
}

func (s *DescribeDataMaskingRunHistoryResponseBodyItems) GetDstTypeCode() *string {
	return s.DstTypeCode
}

func (s *DescribeDataMaskingRunHistoryResponseBodyItems) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeDataMaskingRunHistoryResponseBodyItems) GetFailCode() *string {
	return s.FailCode
}

func (s *DescribeDataMaskingRunHistoryResponseBodyItems) GetFailMsg() *string {
	return s.FailMsg
}

func (s *DescribeDataMaskingRunHistoryResponseBodyItems) GetHasDownloadFile() *int32 {
	return s.HasDownloadFile
}

func (s *DescribeDataMaskingRunHistoryResponseBodyItems) GetHasSubProcess() *int32 {
	return s.HasSubProcess
}

func (s *DescribeDataMaskingRunHistoryResponseBodyItems) GetId() *int64 {
	return s.Id
}

func (s *DescribeDataMaskingRunHistoryResponseBodyItems) GetMaskingCount() *int64 {
	return s.MaskingCount
}

func (s *DescribeDataMaskingRunHistoryResponseBodyItems) GetPercentage() *int32 {
	return s.Percentage
}

func (s *DescribeDataMaskingRunHistoryResponseBodyItems) GetRunIndex() *int32 {
	return s.RunIndex
}

func (s *DescribeDataMaskingRunHistoryResponseBodyItems) GetSrcTableName() *string {
	return s.SrcTableName
}

func (s *DescribeDataMaskingRunHistoryResponseBodyItems) GetSrcType() *int32 {
	return s.SrcType
}

func (s *DescribeDataMaskingRunHistoryResponseBodyItems) GetSrcTypeCode() *string {
	return s.SrcTypeCode
}

func (s *DescribeDataMaskingRunHistoryResponseBodyItems) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeDataMaskingRunHistoryResponseBodyItems) GetStatus() *int32 {
	return s.Status
}

func (s *DescribeDataMaskingRunHistoryResponseBodyItems) GetTaskId() *string {
	return s.TaskId
}

func (s *DescribeDataMaskingRunHistoryResponseBodyItems) GetType() *int32 {
	return s.Type
}

func (s *DescribeDataMaskingRunHistoryResponseBodyItems) SetConflictCount(v int64) *DescribeDataMaskingRunHistoryResponseBodyItems {
	s.ConflictCount = &v
	return s
}

func (s *DescribeDataMaskingRunHistoryResponseBodyItems) SetDstType(v int32) *DescribeDataMaskingRunHistoryResponseBodyItems {
	s.DstType = &v
	return s
}

func (s *DescribeDataMaskingRunHistoryResponseBodyItems) SetDstTypeCode(v string) *DescribeDataMaskingRunHistoryResponseBodyItems {
	s.DstTypeCode = &v
	return s
}

func (s *DescribeDataMaskingRunHistoryResponseBodyItems) SetEndTime(v int64) *DescribeDataMaskingRunHistoryResponseBodyItems {
	s.EndTime = &v
	return s
}

func (s *DescribeDataMaskingRunHistoryResponseBodyItems) SetFailCode(v string) *DescribeDataMaskingRunHistoryResponseBodyItems {
	s.FailCode = &v
	return s
}

func (s *DescribeDataMaskingRunHistoryResponseBodyItems) SetFailMsg(v string) *DescribeDataMaskingRunHistoryResponseBodyItems {
	s.FailMsg = &v
	return s
}

func (s *DescribeDataMaskingRunHistoryResponseBodyItems) SetHasDownloadFile(v int32) *DescribeDataMaskingRunHistoryResponseBodyItems {
	s.HasDownloadFile = &v
	return s
}

func (s *DescribeDataMaskingRunHistoryResponseBodyItems) SetHasSubProcess(v int32) *DescribeDataMaskingRunHistoryResponseBodyItems {
	s.HasSubProcess = &v
	return s
}

func (s *DescribeDataMaskingRunHistoryResponseBodyItems) SetId(v int64) *DescribeDataMaskingRunHistoryResponseBodyItems {
	s.Id = &v
	return s
}

func (s *DescribeDataMaskingRunHistoryResponseBodyItems) SetMaskingCount(v int64) *DescribeDataMaskingRunHistoryResponseBodyItems {
	s.MaskingCount = &v
	return s
}

func (s *DescribeDataMaskingRunHistoryResponseBodyItems) SetPercentage(v int32) *DescribeDataMaskingRunHistoryResponseBodyItems {
	s.Percentage = &v
	return s
}

func (s *DescribeDataMaskingRunHistoryResponseBodyItems) SetRunIndex(v int32) *DescribeDataMaskingRunHistoryResponseBodyItems {
	s.RunIndex = &v
	return s
}

func (s *DescribeDataMaskingRunHistoryResponseBodyItems) SetSrcTableName(v string) *DescribeDataMaskingRunHistoryResponseBodyItems {
	s.SrcTableName = &v
	return s
}

func (s *DescribeDataMaskingRunHistoryResponseBodyItems) SetSrcType(v int32) *DescribeDataMaskingRunHistoryResponseBodyItems {
	s.SrcType = &v
	return s
}

func (s *DescribeDataMaskingRunHistoryResponseBodyItems) SetSrcTypeCode(v string) *DescribeDataMaskingRunHistoryResponseBodyItems {
	s.SrcTypeCode = &v
	return s
}

func (s *DescribeDataMaskingRunHistoryResponseBodyItems) SetStartTime(v int64) *DescribeDataMaskingRunHistoryResponseBodyItems {
	s.StartTime = &v
	return s
}

func (s *DescribeDataMaskingRunHistoryResponseBodyItems) SetStatus(v int32) *DescribeDataMaskingRunHistoryResponseBodyItems {
	s.Status = &v
	return s
}

func (s *DescribeDataMaskingRunHistoryResponseBodyItems) SetTaskId(v string) *DescribeDataMaskingRunHistoryResponseBodyItems {
	s.TaskId = &v
	return s
}

func (s *DescribeDataMaskingRunHistoryResponseBodyItems) SetType(v int32) *DescribeDataMaskingRunHistoryResponseBodyItems {
	s.Type = &v
	return s
}

func (s *DescribeDataMaskingRunHistoryResponseBodyItems) Validate() error {
	return dara.Validate(s)
}
