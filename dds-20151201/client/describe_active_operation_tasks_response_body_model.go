// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeActiveOperationTasksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v []*DescribeActiveOperationTasksResponseBodyItems) *DescribeActiveOperationTasksResponseBody
	GetItems() []*DescribeActiveOperationTasksResponseBodyItems
	SetPageNumber(v int32) *DescribeActiveOperationTasksResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeActiveOperationTasksResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeActiveOperationTasksResponseBody
	GetRequestId() *string
	SetTotalRecordCount(v int32) *DescribeActiveOperationTasksResponseBody
	GetTotalRecordCount() *int32
}

type DescribeActiveOperationTasksResponseBody struct {
	// The O\\&M tasks.
	Items []*DescribeActiveOperationTasksResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// The number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 111E7B16-0A87-4CBA-B271-F34AD61E099F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of returned entries.
	//
	// example:
	//
	// 727
	TotalRecordCount *int32 `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
}

func (s DescribeActiveOperationTasksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeActiveOperationTasksResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeActiveOperationTasksResponseBody) GetItems() []*DescribeActiveOperationTasksResponseBodyItems {
	return s.Items
}

func (s *DescribeActiveOperationTasksResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeActiveOperationTasksResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeActiveOperationTasksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeActiveOperationTasksResponseBody) GetTotalRecordCount() *int32 {
	return s.TotalRecordCount
}

func (s *DescribeActiveOperationTasksResponseBody) SetItems(v []*DescribeActiveOperationTasksResponseBodyItems) *DescribeActiveOperationTasksResponseBody {
	s.Items = v
	return s
}

func (s *DescribeActiveOperationTasksResponseBody) SetPageNumber(v int32) *DescribeActiveOperationTasksResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeActiveOperationTasksResponseBody) SetPageSize(v int32) *DescribeActiveOperationTasksResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeActiveOperationTasksResponseBody) SetRequestId(v string) *DescribeActiveOperationTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeActiveOperationTasksResponseBody) SetTotalRecordCount(v int32) *DescribeActiveOperationTasksResponseBody {
	s.TotalRecordCount = &v
	return s
}

func (s *DescribeActiveOperationTasksResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeActiveOperationTasksResponseBodyItems struct {
	// N/A
	//
	// example:
	//
	// ***
	AllowCancel *string `json:"AllowCancel,omitempty" xml:"AllowCancel,omitempty"`
	// Indicates whether the modification operation is allowed.
	//
	// 	- **0**: The modification operation is not allowed.
	//
	// 	- **1**: The modification operation is allowed.
	//
	// example:
	//
	// 0
	AllowChange *string `json:"AllowChange,omitempty" xml:"AllowChange,omitempty"`
	// The change level of the O\\&M task.
	//
	// example:
	//
	// ***
	ChangeLevel *string `json:"ChangeLevel,omitempty" xml:"ChangeLevel,omitempty"`
	// N/A
	//
	// example:
	//
	// ***
	ChangeLevelEn *string `json:"ChangeLevelEn,omitempty" xml:"ChangeLevelEn,omitempty"`
	// The task type in English.
	//
	// example:
	//
	// ***
	ChangeLevelZh *string `json:"ChangeLevelZh,omitempty" xml:"ChangeLevelZh,omitempty"`
	// The time when the task was created.
	//
	// example:
	//
	// 2021-07-14 10:48:43
	CreatedTime *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// N/A
	//
	// example:
	//
	// ***
	CurrentAVZ *string `json:"CurrentAVZ,omitempty" xml:"CurrentAVZ,omitempty"`
	// The type of the database engine.
	//
	// example:
	//
	// mongoDb
	DbType *string `json:"DbType,omitempty" xml:"DbType,omitempty"`
	// The version of the database engine.
	//
	// example:
	//
	// 5.6
	DbVersion *string `json:"DbVersion,omitempty" xml:"DbVersion,omitempty"`
	// The end time of the O\\&M task.
	//
	// example:
	//
	// 1646014421633
	Deadline *string `json:"Deadline,omitempty" xml:"Deadline,omitempty"`
	// The task ID.
	//
	// example:
	//
	// ***
	Id *int32 `json:"Id,omitempty" xml:"Id,omitempty"`
	// N/A
	//
	// example:
	//
	// ***
	ImpactEn *string `json:"ImpactEn,omitempty" xml:"ImpactEn,omitempty"`
	// N/A
	//
	// example:
	//
	// ***
	ImpactZh *string `json:"ImpactZh,omitempty" xml:"ImpactZh,omitempty"`
	// The description of the instance.
	//
	// example:
	//
	// ***
	InsComment *string `json:"InsComment,omitempty" xml:"InsComment,omitempty"`
	// The ID of the node.
	//
	// example:
	//
	// ***
	InsName *string `json:"InsName,omitempty" xml:"InsName,omitempty"`
	// The time when the task is modified. The time follows the ISO 8601 standard in the *yyyy-mm-dd*T*hh:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2022-12-24T09:48:01Z
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// The required preparation period between the task start time and the switchover time. The time is displayed in the *HH:mm:ss	- format.
	//
	// example:
	//
	// ***
	PrepareInterval *string `json:"PrepareInterval,omitempty" xml:"PrepareInterval,omitempty"`
	// The region of the instance.
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The result information. The value of this parameter can be ignored.
	//
	// example:
	//
	// ***
	ResultInfo *string `json:"ResultInfo,omitempty" xml:"ResultInfo,omitempty"`
	// The start time of the task. The time follows the ISO 8601 standard in the *yyyy-mm-dd	- T*hh:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2022-12-24T06:01:05Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// N/A
	//
	// example:
	//
	// ***
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The subinstances.
	SubInsNames []*string `json:"SubInsNames,omitempty" xml:"SubInsNames,omitempty" type:"Repeated"`
	// The switchover point in time in which disconnection may occur. The time follows the ISO 8601 standard in the *yyyy-mm-dd*T*hh:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2022-11-24T11:20:00Z
	SwitchTime *string `json:"SwitchTime,omitempty" xml:"SwitchTime,omitempty"`
	// The task type.
	//
	// example:
	//
	// ***
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	// N/A
	//
	// example:
	//
	// ***
	TaskTypeEn *string `json:"TaskTypeEn,omitempty" xml:"TaskTypeEn,omitempty"`
	// The task type in Chinese.
	//
	// example:
	//
	// ***
	TaskTypeZh *string `json:"TaskTypeZh,omitempty" xml:"TaskTypeZh,omitempty"`
}

func (s DescribeActiveOperationTasksResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeActiveOperationTasksResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeActiveOperationTasksResponseBodyItems) GetAllowCancel() *string {
	return s.AllowCancel
}

func (s *DescribeActiveOperationTasksResponseBodyItems) GetAllowChange() *string {
	return s.AllowChange
}

func (s *DescribeActiveOperationTasksResponseBodyItems) GetChangeLevel() *string {
	return s.ChangeLevel
}

func (s *DescribeActiveOperationTasksResponseBodyItems) GetChangeLevelEn() *string {
	return s.ChangeLevelEn
}

func (s *DescribeActiveOperationTasksResponseBodyItems) GetChangeLevelZh() *string {
	return s.ChangeLevelZh
}

func (s *DescribeActiveOperationTasksResponseBodyItems) GetCreatedTime() *string {
	return s.CreatedTime
}

func (s *DescribeActiveOperationTasksResponseBodyItems) GetCurrentAVZ() *string {
	return s.CurrentAVZ
}

func (s *DescribeActiveOperationTasksResponseBodyItems) GetDbType() *string {
	return s.DbType
}

func (s *DescribeActiveOperationTasksResponseBodyItems) GetDbVersion() *string {
	return s.DbVersion
}

func (s *DescribeActiveOperationTasksResponseBodyItems) GetDeadline() *string {
	return s.Deadline
}

func (s *DescribeActiveOperationTasksResponseBodyItems) GetId() *int32 {
	return s.Id
}

func (s *DescribeActiveOperationTasksResponseBodyItems) GetImpactEn() *string {
	return s.ImpactEn
}

func (s *DescribeActiveOperationTasksResponseBodyItems) GetImpactZh() *string {
	return s.ImpactZh
}

func (s *DescribeActiveOperationTasksResponseBodyItems) GetInsComment() *string {
	return s.InsComment
}

func (s *DescribeActiveOperationTasksResponseBodyItems) GetInsName() *string {
	return s.InsName
}

func (s *DescribeActiveOperationTasksResponseBodyItems) GetModifiedTime() *string {
	return s.ModifiedTime
}

func (s *DescribeActiveOperationTasksResponseBodyItems) GetPrepareInterval() *string {
	return s.PrepareInterval
}

func (s *DescribeActiveOperationTasksResponseBodyItems) GetRegion() *string {
	return s.Region
}

func (s *DescribeActiveOperationTasksResponseBodyItems) GetResultInfo() *string {
	return s.ResultInfo
}

func (s *DescribeActiveOperationTasksResponseBodyItems) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeActiveOperationTasksResponseBodyItems) GetStatus() *int32 {
	return s.Status
}

func (s *DescribeActiveOperationTasksResponseBodyItems) GetSubInsNames() []*string {
	return s.SubInsNames
}

func (s *DescribeActiveOperationTasksResponseBodyItems) GetSwitchTime() *string {
	return s.SwitchTime
}

func (s *DescribeActiveOperationTasksResponseBodyItems) GetTaskType() *string {
	return s.TaskType
}

func (s *DescribeActiveOperationTasksResponseBodyItems) GetTaskTypeEn() *string {
	return s.TaskTypeEn
}

func (s *DescribeActiveOperationTasksResponseBodyItems) GetTaskTypeZh() *string {
	return s.TaskTypeZh
}

func (s *DescribeActiveOperationTasksResponseBodyItems) SetAllowCancel(v string) *DescribeActiveOperationTasksResponseBodyItems {
	s.AllowCancel = &v
	return s
}

func (s *DescribeActiveOperationTasksResponseBodyItems) SetAllowChange(v string) *DescribeActiveOperationTasksResponseBodyItems {
	s.AllowChange = &v
	return s
}

func (s *DescribeActiveOperationTasksResponseBodyItems) SetChangeLevel(v string) *DescribeActiveOperationTasksResponseBodyItems {
	s.ChangeLevel = &v
	return s
}

func (s *DescribeActiveOperationTasksResponseBodyItems) SetChangeLevelEn(v string) *DescribeActiveOperationTasksResponseBodyItems {
	s.ChangeLevelEn = &v
	return s
}

func (s *DescribeActiveOperationTasksResponseBodyItems) SetChangeLevelZh(v string) *DescribeActiveOperationTasksResponseBodyItems {
	s.ChangeLevelZh = &v
	return s
}

func (s *DescribeActiveOperationTasksResponseBodyItems) SetCreatedTime(v string) *DescribeActiveOperationTasksResponseBodyItems {
	s.CreatedTime = &v
	return s
}

func (s *DescribeActiveOperationTasksResponseBodyItems) SetCurrentAVZ(v string) *DescribeActiveOperationTasksResponseBodyItems {
	s.CurrentAVZ = &v
	return s
}

func (s *DescribeActiveOperationTasksResponseBodyItems) SetDbType(v string) *DescribeActiveOperationTasksResponseBodyItems {
	s.DbType = &v
	return s
}

func (s *DescribeActiveOperationTasksResponseBodyItems) SetDbVersion(v string) *DescribeActiveOperationTasksResponseBodyItems {
	s.DbVersion = &v
	return s
}

func (s *DescribeActiveOperationTasksResponseBodyItems) SetDeadline(v string) *DescribeActiveOperationTasksResponseBodyItems {
	s.Deadline = &v
	return s
}

func (s *DescribeActiveOperationTasksResponseBodyItems) SetId(v int32) *DescribeActiveOperationTasksResponseBodyItems {
	s.Id = &v
	return s
}

func (s *DescribeActiveOperationTasksResponseBodyItems) SetImpactEn(v string) *DescribeActiveOperationTasksResponseBodyItems {
	s.ImpactEn = &v
	return s
}

func (s *DescribeActiveOperationTasksResponseBodyItems) SetImpactZh(v string) *DescribeActiveOperationTasksResponseBodyItems {
	s.ImpactZh = &v
	return s
}

func (s *DescribeActiveOperationTasksResponseBodyItems) SetInsComment(v string) *DescribeActiveOperationTasksResponseBodyItems {
	s.InsComment = &v
	return s
}

func (s *DescribeActiveOperationTasksResponseBodyItems) SetInsName(v string) *DescribeActiveOperationTasksResponseBodyItems {
	s.InsName = &v
	return s
}

func (s *DescribeActiveOperationTasksResponseBodyItems) SetModifiedTime(v string) *DescribeActiveOperationTasksResponseBodyItems {
	s.ModifiedTime = &v
	return s
}

func (s *DescribeActiveOperationTasksResponseBodyItems) SetPrepareInterval(v string) *DescribeActiveOperationTasksResponseBodyItems {
	s.PrepareInterval = &v
	return s
}

func (s *DescribeActiveOperationTasksResponseBodyItems) SetRegion(v string) *DescribeActiveOperationTasksResponseBodyItems {
	s.Region = &v
	return s
}

func (s *DescribeActiveOperationTasksResponseBodyItems) SetResultInfo(v string) *DescribeActiveOperationTasksResponseBodyItems {
	s.ResultInfo = &v
	return s
}

func (s *DescribeActiveOperationTasksResponseBodyItems) SetStartTime(v string) *DescribeActiveOperationTasksResponseBodyItems {
	s.StartTime = &v
	return s
}

func (s *DescribeActiveOperationTasksResponseBodyItems) SetStatus(v int32) *DescribeActiveOperationTasksResponseBodyItems {
	s.Status = &v
	return s
}

func (s *DescribeActiveOperationTasksResponseBodyItems) SetSubInsNames(v []*string) *DescribeActiveOperationTasksResponseBodyItems {
	s.SubInsNames = v
	return s
}

func (s *DescribeActiveOperationTasksResponseBodyItems) SetSwitchTime(v string) *DescribeActiveOperationTasksResponseBodyItems {
	s.SwitchTime = &v
	return s
}

func (s *DescribeActiveOperationTasksResponseBodyItems) SetTaskType(v string) *DescribeActiveOperationTasksResponseBodyItems {
	s.TaskType = &v
	return s
}

func (s *DescribeActiveOperationTasksResponseBodyItems) SetTaskTypeEn(v string) *DescribeActiveOperationTasksResponseBodyItems {
	s.TaskTypeEn = &v
	return s
}

func (s *DescribeActiveOperationTasksResponseBodyItems) SetTaskTypeZh(v string) *DescribeActiveOperationTasksResponseBodyItems {
	s.TaskTypeZh = &v
	return s
}

func (s *DescribeActiveOperationTasksResponseBodyItems) Validate() error {
	return dara.Validate(s)
}
