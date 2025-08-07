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
	Items []*DescribeActiveOperationTasksResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// FAF88508-D5F8-52B1-8824-262601769E31
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 1
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
	// example:
	//
	// 0
	AllowCancel *int64 `json:"AllowCancel,omitempty" xml:"AllowCancel,omitempty"`
	// example:
	//
	// 0
	AllowChange *int64 `json:"AllowChange,omitempty" xml:"AllowChange,omitempty"`
	// example:
	//
	// S0
	ChangeLevel *string `json:"ChangeLevel,omitempty" xml:"ChangeLevel,omitempty"`
	// example:
	//
	// System maintenance
	ChangeLevelEn *string `json:"ChangeLevelEn,omitempty" xml:"ChangeLevelEn,omitempty"`
	ChangeLevelZh *string `json:"ChangeLevelZh,omitempty" xml:"ChangeLevelZh,omitempty"`
	// example:
	//
	// 2020-06-09T22:00:42Z
	CreatedTime *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// example:
	//
	// cn-beijing-h
	CurrentAVZ *string `json:"CurrentAVZ,omitempty" xml:"CurrentAVZ,omitempty"`
	// example:
	//
	// pc-*****************
	DBClusterId *string   `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	DBNodeIds   []*string `json:"DBNodeIds,omitempty" xml:"DBNodeIds,omitempty" type:"Repeated"`
	// example:
	//
	// MySQL
	DBType *string `json:"DBType,omitempty" xml:"DBType,omitempty"`
	// example:
	//
	// 8.0
	DBVersion *string `json:"DBVersion,omitempty" xml:"DBVersion,omitempty"`
	// example:
	//
	// 2020-06-11T15:59:59Z
	Deadline *string `json:"Deadline,omitempty" xml:"Deadline,omitempty"`
	// example:
	//
	// TransientDisconnection
	Impact *string `json:"Impact,omitempty" xml:"Impact,omitempty"`
	// example:
	//
	// Transient instance disconnection
	ImpactEn *string `json:"ImpactEn,omitempty" xml:"ImpactEn,omitempty"`
	ImpactZh *string `json:"ImpactZh,omitempty" xml:"ImpactZh,omitempty"`
	// example:
	//
	// test
	InsComment *string `json:"InsComment,omitempty" xml:"InsComment,omitempty"`
	// example:
	//
	// 2020-06-09T22:00:42Z
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// example:
	//
	// 04:00:00
	PrepareInterval *string `json:"PrepareInterval,omitempty" xml:"PrepareInterval,omitempty"`
	// example:
	//
	// cn-hangzhou
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// example:
	//
	// userCancel
	ResultInfo *string `json:"ResultInfo,omitempty" xml:"ResultInfo,omitempty"`
	// example:
	//
	// 2023-05-19T02:48:17Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// 3
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 2020-06-09T22:00:00Z
	SwitchTime *string `json:"SwitchTime,omitempty" xml:"SwitchTime,omitempty"`
	// example:
	//
	// 107202351
	TaskId *int32 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// example:
	//
	// {
	//
	//       "Action": "UpgradeDBInstance"
	//
	// }
	TaskParams *string `json:"TaskParams,omitempty" xml:"TaskParams,omitempty"`
	// example:
	//
	// DatabaseSoftwareUpgrading
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	// example:
	//
	// Minor version update
	TaskTypeEn *string `json:"TaskTypeEn,omitempty" xml:"TaskTypeEn,omitempty"`
	TaskTypeZh *string `json:"TaskTypeZh,omitempty" xml:"TaskTypeZh,omitempty"`
}

func (s DescribeActiveOperationTasksResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeActiveOperationTasksResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeActiveOperationTasksResponseBodyItems) GetAllowCancel() *int64 {
	return s.AllowCancel
}

func (s *DescribeActiveOperationTasksResponseBodyItems) GetAllowChange() *int64 {
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

func (s *DescribeActiveOperationTasksResponseBodyItems) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeActiveOperationTasksResponseBodyItems) GetDBNodeIds() []*string {
	return s.DBNodeIds
}

func (s *DescribeActiveOperationTasksResponseBodyItems) GetDBType() *string {
	return s.DBType
}

func (s *DescribeActiveOperationTasksResponseBodyItems) GetDBVersion() *string {
	return s.DBVersion
}

func (s *DescribeActiveOperationTasksResponseBodyItems) GetDeadline() *string {
	return s.Deadline
}

func (s *DescribeActiveOperationTasksResponseBodyItems) GetImpact() *string {
	return s.Impact
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

func (s *DescribeActiveOperationTasksResponseBodyItems) GetSwitchTime() *string {
	return s.SwitchTime
}

func (s *DescribeActiveOperationTasksResponseBodyItems) GetTaskId() *int32 {
	return s.TaskId
}

func (s *DescribeActiveOperationTasksResponseBodyItems) GetTaskParams() *string {
	return s.TaskParams
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

func (s *DescribeActiveOperationTasksResponseBodyItems) SetAllowCancel(v int64) *DescribeActiveOperationTasksResponseBodyItems {
	s.AllowCancel = &v
	return s
}

func (s *DescribeActiveOperationTasksResponseBodyItems) SetAllowChange(v int64) *DescribeActiveOperationTasksResponseBodyItems {
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

func (s *DescribeActiveOperationTasksResponseBodyItems) SetDBClusterId(v string) *DescribeActiveOperationTasksResponseBodyItems {
	s.DBClusterId = &v
	return s
}

func (s *DescribeActiveOperationTasksResponseBodyItems) SetDBNodeIds(v []*string) *DescribeActiveOperationTasksResponseBodyItems {
	s.DBNodeIds = v
	return s
}

func (s *DescribeActiveOperationTasksResponseBodyItems) SetDBType(v string) *DescribeActiveOperationTasksResponseBodyItems {
	s.DBType = &v
	return s
}

func (s *DescribeActiveOperationTasksResponseBodyItems) SetDBVersion(v string) *DescribeActiveOperationTasksResponseBodyItems {
	s.DBVersion = &v
	return s
}

func (s *DescribeActiveOperationTasksResponseBodyItems) SetDeadline(v string) *DescribeActiveOperationTasksResponseBodyItems {
	s.Deadline = &v
	return s
}

func (s *DescribeActiveOperationTasksResponseBodyItems) SetImpact(v string) *DescribeActiveOperationTasksResponseBodyItems {
	s.Impact = &v
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

func (s *DescribeActiveOperationTasksResponseBodyItems) SetSwitchTime(v string) *DescribeActiveOperationTasksResponseBodyItems {
	s.SwitchTime = &v
	return s
}

func (s *DescribeActiveOperationTasksResponseBodyItems) SetTaskId(v int32) *DescribeActiveOperationTasksResponseBodyItems {
	s.TaskId = &v
	return s
}

func (s *DescribeActiveOperationTasksResponseBodyItems) SetTaskParams(v string) *DescribeActiveOperationTasksResponseBodyItems {
	s.TaskParams = &v
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
