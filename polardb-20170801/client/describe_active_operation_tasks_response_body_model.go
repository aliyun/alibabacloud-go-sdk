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
	// The list of O\\&M tasks.
	Items []*DescribeActiveOperationTasksResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// The page number. The value must be greater than 0 and cannot exceed the maximum value of the Integer data type. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page. Valid values:
	//
	// - **30**
	//
	// - **50**
	//
	// - **100**
	//
	// Default value: **30**.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// FAF88508-D5F8-52B1-8824-262601769E31
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of task records returned.
	//
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

type DescribeActiveOperationTasksResponseBodyItems struct {
	// Indicates whether cancellation is allowed. Valid values:
	//
	// - 1: indicates that users are allowed to cancel the task.
	//
	// - 0: indicates that cancellation is not allowed.
	//
	// example:
	//
	// 0
	AllowCancel *int64 `json:"AllowCancel,omitempty" xml:"AllowCancel,omitempty"`
	// Indicates whether time modification is allowed. Valid values:
	//
	// - 1: indicates that users are allowed to modify the time.
	//
	// - 0: indicates that users are not allowed to modify the time.
	//
	// example:
	//
	// 0
	AllowChange *int64 `json:"AllowChange,omitempty" xml:"AllowChange,omitempty"`
	// The event level code. Valid values:
	//
	// - S1: system maintenance.
	//
	// - S0: threat fix.
	//
	// example:
	//
	// S0
	ChangeLevel *string `json:"ChangeLevel,omitempty" xml:"ChangeLevel,omitempty"`
	// The event level (English).
	//
	// example:
	//
	// System maintenance
	ChangeLevelEn *string `json:"ChangeLevelEn,omitempty" xml:"ChangeLevelEn,omitempty"`
	// The event level (Chinese).
	//
	// example:
	//
	// 系统运维
	ChangeLevelZh *string `json:"ChangeLevelZh,omitempty" xml:"ChangeLevelZh,omitempty"`
	// The creation time. Format: YYYY-MM-DDTHH:mm:ssZ.
	//
	// example:
	//
	// 2020-06-09T22:00:42Z
	CreatedTime *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// The current zone.
	//
	// example:
	//
	// cn-beijing-h
	CurrentAVZ *string `json:"CurrentAVZ,omitempty" xml:"CurrentAVZ,omitempty"`
	// The cluster ID.
	//
	// example:
	//
	// pc-*****************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The list of node IDs.
	DBNodeIds []*string `json:"DBNodeIds,omitempty" xml:"DBNodeIds,omitempty" type:"Repeated"`
	// The database engine type. Valid values:
	//
	// - **MySQL**
	//
	// - **PostgreSQL**
	//
	// - **Oracle**
	//
	// example:
	//
	// MySQL
	DBType *string `json:"DBType,omitempty" xml:"DBType,omitempty"`
	// The kernel version number.
	//
	// example:
	//
	// 8.0
	DBVersion *string `json:"DBVersion,omitempty" xml:"DBVersion,omitempty"`
	// The latest deadline for the adjustable range of task execution time. Format: YYYY-MM-DDTHH:mm:ssZ.
	//
	// example:
	//
	// 2020-06-11T15:59:59Z
	Deadline *string `json:"Deadline,omitempty" xml:"Deadline,omitempty"`
	// The event impact.
	//
	// example:
	//
	// TransientDisconnection
	Impact *string `json:"Impact,omitempty" xml:"Impact,omitempty"`
	// The event impact (English).
	//
	// example:
	//
	// Transient instance disconnection
	ImpactEn *string `json:"ImpactEn,omitempty" xml:"ImpactEn,omitempty"`
	// The event impact (Chinese).
	//
	// example:
	//
	// 集群闪断
	ImpactZh *string `json:"ImpactZh,omitempty" xml:"ImpactZh,omitempty"`
	// The cluster alias or cluster comment.
	//
	// example:
	//
	// test
	InsComment *string `json:"InsComment,omitempty" xml:"InsComment,omitempty"`
	// The modification time. Format: YYYY-MM-DDTHH:mm:ssZ.
	//
	// example:
	//
	// 2020-06-09T22:00:42Z
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// The preparation time required between the start time and the switch time. Format: HH:mm:ss.
	//
	// example:
	//
	// 04:00:00
	PrepareInterval *string `json:"PrepareInterval,omitempty" xml:"PrepareInterval,omitempty"`
	// The region code.
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The execution result information.
	//
	// example:
	//
	// userCancel
	ResultInfo *string `json:"ResultInfo,omitempty" xml:"ResultInfo,omitempty"`
	// The time when the background task is executed. Format: YYYY-MM-DDTHH:mm:ssZ.
	//
	// example:
	//
	// 2023-05-19T02:48:17Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The task status. Valid values:
	//
	// - 0: indicates waiting for notification.
	//
	// - 1: indicates retry notification.
	//
	// - 2: indicates waiting for user-specified time.
	//
	// - 3: indicates waiting for processing.
	//
	// - 4: indicates in progress.
	//
	// - 5: indicates successfully completed.
	//
	// - 6: indicates failed.
	//
	// - 7: indicates canceled.
	//
	// example:
	//
	// 3
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The time when the background initiates the switch operation. Format: YYYY-MM-DDTHH:mm:ssZ.
	//
	// example:
	//
	// 2020-06-09T22:00:00Z
	SwitchTime *string `json:"SwitchTime,omitempty" xml:"SwitchTime,omitempty"`
	// The task ID.
	//
	// example:
	//
	// 107202351
	TaskId *int32 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The task parameters.
	//
	// example:
	//
	// {
	//
	//       "Action": "UpgradeDBInstance"
	//
	// }
	TaskParams *string `json:"TaskParams,omitempty" xml:"TaskParams,omitempty"`
	// The type of the pending event task. Valid values:
	//
	// - **DatabaseSoftwareUpgrading**: database software upgrade.
	//
	// - **DatabaseHardwareMaintenance**: hardware maintenance and upgrade.
	//
	// - **DatabaseStorageUpgrading**: database storage upgrade.
	//
	// - **DatabaseProxyUpgrading**: proxy minor version upgrade.
	//
	// - **all**: returns all types of pending events.
	//
	// example:
	//
	// DatabaseSoftwareUpgrading
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	// The task reason in English.
	//
	// example:
	//
	// Minor version update
	TaskTypeEn *string `json:"TaskTypeEn,omitempty" xml:"TaskTypeEn,omitempty"`
	// The task reason in Chinese.
	//
	// example:
	//
	// 小版本升级
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
