// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDataMaskingTasksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *DescribeDataMaskingTasksResponseBody
	GetCurrentPage() *int32
	SetItems(v []*DescribeDataMaskingTasksResponseBodyItems) *DescribeDataMaskingTasksResponseBody
	GetItems() []*DescribeDataMaskingTasksResponseBodyItems
	SetPageSize(v int32) *DescribeDataMaskingTasksResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeDataMaskingTasksResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeDataMaskingTasksResponseBody
	GetTotalCount() *int32
}

type DescribeDataMaskingTasksResponseBody struct {
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// A list of de-identification tasks.
	Items []*DescribeDataMaskingTasksResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
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

func (s DescribeDataMaskingTasksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDataMaskingTasksResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDataMaskingTasksResponseBody) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeDataMaskingTasksResponseBody) GetItems() []*DescribeDataMaskingTasksResponseBodyItems {
	return s.Items
}

func (s *DescribeDataMaskingTasksResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeDataMaskingTasksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDataMaskingTasksResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeDataMaskingTasksResponseBody) SetCurrentPage(v int32) *DescribeDataMaskingTasksResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *DescribeDataMaskingTasksResponseBody) SetItems(v []*DescribeDataMaskingTasksResponseBodyItems) *DescribeDataMaskingTasksResponseBody {
	s.Items = v
	return s
}

func (s *DescribeDataMaskingTasksResponseBody) SetPageSize(v int32) *DescribeDataMaskingTasksResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeDataMaskingTasksResponseBody) SetRequestId(v string) *DescribeDataMaskingTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDataMaskingTasksResponseBody) SetTotalCount(v int32) *DescribeDataMaskingTasksResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeDataMaskingTasksResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDataMaskingTasksResponseBodyItems struct {
	// The member account to which the desensitization target belongs.
	//
	// example:
	//
	// 192479427903xxxx
	DstMemberAccount *int64 `json:"DstMemberAccount,omitempty" xml:"DstMemberAccount,omitempty"`
	// The destination path.
	DstPath *string `json:"DstPath,omitempty" xml:"DstPath,omitempty"`
	// The service to which the data to be de-identified belongs. Valid values: **1**, **2**, **3**, **4**, and **5**. The value 1 indicates MaxCompute. The value 2 indicates OSS. The value 3 indicates AnalyticDB for MySQL. The value 4 indicates Tablestore. The value 5 indicates ApsaraDB RDS.
	//
	// example:
	//
	// 5
	DstType *int32 `json:"DstType,omitempty" xml:"DstType,omitempty"`
	// The type of the service to which the de-identified data belongs. Valid values: **MaxCompute, OSS, ADS, OTS, and RDS**.
	//
	// example:
	//
	// RDS
	DstTypeCode *string `json:"DstTypeCode,omitempty" xml:"DstTypeCode,omitempty"`
	// The time when the de-identification task is created. The value is a UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1582992000000
	GmtCreate *int64 `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// Indicates whether the de-identification task is running.
	//
	// example:
	//
	// false
	HasUnfinishProcess *bool `json:"HasUnfinishProcess,omitempty" xml:"HasUnfinishProcess,omitempty"`
	// The task ID.
	//
	// example:
	//
	// 1
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// Indicates whether the source table is de-identified.
	//
	// example:
	//
	// false
	OriginalTable *bool `json:"OriginalTable,omitempty" xml:"OriginalTable,omitempty"`
	// The user who created the de-identification task.
	//
	// example:
	//
	// owner
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The number of times that the de-identification task is run.
	//
	// example:
	//
	// 1
	RunCount *int32 `json:"RunCount,omitempty" xml:"RunCount,omitempty"`
	// The member account to which the desensitization source belongs.
	//
	// example:
	//
	// 192479427903xxxx
	SrcMemberAccount *int64 `json:"SrcMemberAccount,omitempty" xml:"SrcMemberAccount,omitempty"`
	// The source path.
	SrcPath *string `json:"SrcPath,omitempty" xml:"SrcPath,omitempty"`
	// The type of the service to which the data to be de-identified belongs. Valid values: **1**, **2**, **3**, **4**, and **5**. The value 1 indicates MaxCompute. The value 2 indicates OSS. The value 3 indicates AnalyticDB for MySQL. The value 4 indicates Tablestore. The value 5 indicates ApsaraDB RDS.
	//
	// example:
	//
	// 5
	SrcType *int32 `json:"SrcType,omitempty" xml:"SrcType,omitempty"`
	// The type of the service to which the data to be de-identified belongs. Valid values: **MaxCompute, OSS, ADS, OTS, and RDS**.
	//
	// example:
	//
	// RDS
	SrcTypeCode *string `json:"SrcTypeCode,omitempty" xml:"SrcTypeCode,omitempty"`
	// The status of the task. Valid values:
	//
	// 	- **0**: disabled
	//
	// 	- **1**: enabled
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
	// The name of the task.
	//
	// example:
	//
	// Task name
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	// The mode in which the de-identification task is run. Valid values:
	//
	// 	- **1**: manual
	//
	// 	- **2**: scheduled
	//
	// 	- **3**: manual and scheduled
	//
	// example:
	//
	// 1
	TriggerType *int32 `json:"TriggerType,omitempty" xml:"TriggerType,omitempty"`
}

func (s DescribeDataMaskingTasksResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeDataMaskingTasksResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeDataMaskingTasksResponseBodyItems) GetDstMemberAccount() *int64 {
	return s.DstMemberAccount
}

func (s *DescribeDataMaskingTasksResponseBodyItems) GetDstPath() *string {
	return s.DstPath
}

func (s *DescribeDataMaskingTasksResponseBodyItems) GetDstType() *int32 {
	return s.DstType
}

func (s *DescribeDataMaskingTasksResponseBodyItems) GetDstTypeCode() *string {
	return s.DstTypeCode
}

func (s *DescribeDataMaskingTasksResponseBodyItems) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *DescribeDataMaskingTasksResponseBodyItems) GetHasUnfinishProcess() *bool {
	return s.HasUnfinishProcess
}

func (s *DescribeDataMaskingTasksResponseBodyItems) GetId() *int64 {
	return s.Id
}

func (s *DescribeDataMaskingTasksResponseBodyItems) GetOriginalTable() *bool {
	return s.OriginalTable
}

func (s *DescribeDataMaskingTasksResponseBodyItems) GetOwner() *string {
	return s.Owner
}

func (s *DescribeDataMaskingTasksResponseBodyItems) GetRunCount() *int32 {
	return s.RunCount
}

func (s *DescribeDataMaskingTasksResponseBodyItems) GetSrcMemberAccount() *int64 {
	return s.SrcMemberAccount
}

func (s *DescribeDataMaskingTasksResponseBodyItems) GetSrcPath() *string {
	return s.SrcPath
}

func (s *DescribeDataMaskingTasksResponseBodyItems) GetSrcType() *int32 {
	return s.SrcType
}

func (s *DescribeDataMaskingTasksResponseBodyItems) GetSrcTypeCode() *string {
	return s.SrcTypeCode
}

func (s *DescribeDataMaskingTasksResponseBodyItems) GetStatus() *int32 {
	return s.Status
}

func (s *DescribeDataMaskingTasksResponseBodyItems) GetTaskId() *string {
	return s.TaskId
}

func (s *DescribeDataMaskingTasksResponseBodyItems) GetTaskName() *string {
	return s.TaskName
}

func (s *DescribeDataMaskingTasksResponseBodyItems) GetTriggerType() *int32 {
	return s.TriggerType
}

func (s *DescribeDataMaskingTasksResponseBodyItems) SetDstMemberAccount(v int64) *DescribeDataMaskingTasksResponseBodyItems {
	s.DstMemberAccount = &v
	return s
}

func (s *DescribeDataMaskingTasksResponseBodyItems) SetDstPath(v string) *DescribeDataMaskingTasksResponseBodyItems {
	s.DstPath = &v
	return s
}

func (s *DescribeDataMaskingTasksResponseBodyItems) SetDstType(v int32) *DescribeDataMaskingTasksResponseBodyItems {
	s.DstType = &v
	return s
}

func (s *DescribeDataMaskingTasksResponseBodyItems) SetDstTypeCode(v string) *DescribeDataMaskingTasksResponseBodyItems {
	s.DstTypeCode = &v
	return s
}

func (s *DescribeDataMaskingTasksResponseBodyItems) SetGmtCreate(v int64) *DescribeDataMaskingTasksResponseBodyItems {
	s.GmtCreate = &v
	return s
}

func (s *DescribeDataMaskingTasksResponseBodyItems) SetHasUnfinishProcess(v bool) *DescribeDataMaskingTasksResponseBodyItems {
	s.HasUnfinishProcess = &v
	return s
}

func (s *DescribeDataMaskingTasksResponseBodyItems) SetId(v int64) *DescribeDataMaskingTasksResponseBodyItems {
	s.Id = &v
	return s
}

func (s *DescribeDataMaskingTasksResponseBodyItems) SetOriginalTable(v bool) *DescribeDataMaskingTasksResponseBodyItems {
	s.OriginalTable = &v
	return s
}

func (s *DescribeDataMaskingTasksResponseBodyItems) SetOwner(v string) *DescribeDataMaskingTasksResponseBodyItems {
	s.Owner = &v
	return s
}

func (s *DescribeDataMaskingTasksResponseBodyItems) SetRunCount(v int32) *DescribeDataMaskingTasksResponseBodyItems {
	s.RunCount = &v
	return s
}

func (s *DescribeDataMaskingTasksResponseBodyItems) SetSrcMemberAccount(v int64) *DescribeDataMaskingTasksResponseBodyItems {
	s.SrcMemberAccount = &v
	return s
}

func (s *DescribeDataMaskingTasksResponseBodyItems) SetSrcPath(v string) *DescribeDataMaskingTasksResponseBodyItems {
	s.SrcPath = &v
	return s
}

func (s *DescribeDataMaskingTasksResponseBodyItems) SetSrcType(v int32) *DescribeDataMaskingTasksResponseBodyItems {
	s.SrcType = &v
	return s
}

func (s *DescribeDataMaskingTasksResponseBodyItems) SetSrcTypeCode(v string) *DescribeDataMaskingTasksResponseBodyItems {
	s.SrcTypeCode = &v
	return s
}

func (s *DescribeDataMaskingTasksResponseBodyItems) SetStatus(v int32) *DescribeDataMaskingTasksResponseBodyItems {
	s.Status = &v
	return s
}

func (s *DescribeDataMaskingTasksResponseBodyItems) SetTaskId(v string) *DescribeDataMaskingTasksResponseBodyItems {
	s.TaskId = &v
	return s
}

func (s *DescribeDataMaskingTasksResponseBodyItems) SetTaskName(v string) *DescribeDataMaskingTasksResponseBodyItems {
	s.TaskName = &v
	return s
}

func (s *DescribeDataMaskingTasksResponseBodyItems) SetTriggerType(v int32) *DescribeDataMaskingTasksResponseBodyItems {
	s.TriggerType = &v
	return s
}

func (s *DescribeDataMaskingTasksResponseBodyItems) Validate() error {
	return dara.Validate(s)
}
