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
	// The list of O&M tasks.
	Items []*DescribeActiveOperationTasksResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// The page number. The value must be greater than 0. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: 25. Maximum value: 100.
	//
	// example:
	//
	// 25
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// D4D4BE8A-DD46-440A-BFCD-EE31DA81****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of task records returned.
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
	// Indicates whether the task can be canceled. A value of 1 indicates that the task can be canceled. A value of 0 indicates that the task cannot be canceled.
	//
	// example:
	//
	// 1
	AllowCancel *string `json:"AllowCancel,omitempty" xml:"AllowCancel,omitempty"`
	// Indicates whether the task time can be modified. A value of 1 indicates that the time can be modified. A value of 0 indicates that the time cannot be modified.
	//
	// example:
	//
	// 1
	AllowChange *string `json:"AllowChange,omitempty" xml:"AllowChange,omitempty"`
	// The event level code. S1 indicates system O&M. S0 indicates risk recovery.
	//
	// example:
	//
	// S1
	ChangeLevel *string `json:"ChangeLevel,omitempty" xml:"ChangeLevel,omitempty"`
	// The event level in English.
	//
	// example:
	//
	// System maintenance
	ChangeLevelEn *string `json:"ChangeLevelEn,omitempty" xml:"ChangeLevelEn,omitempty"`
	// The event level in Chinese.
	//
	// example:
	//
	// 系统运维
	ChangeLevelZh *string `json:"ChangeLevelZh,omitempty" xml:"ChangeLevelZh,omitempty"`
	// The creation time. The time is in UTC and follows the format of YYYY-MM-DDTHH:mm:ssZ.
	//
	// example:
	//
	// 2018-05-30T14:30:00Z
	CreatedTime *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// The current zone.
	//
	// example:
	//
	// cn-beijing-h
	CurrentAVZ *string `json:"CurrentAVZ,omitempty" xml:"CurrentAVZ,omitempty"`
	// The database type, such as mysql, pgsql, or mssql.
	//
	// example:
	//
	// mysql
	DbType *string `json:"DbType,omitempty" xml:"DbType,omitempty"`
	// The Milvus version number.
	//
	// example:
	//
	// 5.7
	DbVersion *string `json:"DbVersion,omitempty" xml:"DbVersion,omitempty"`
	// The latest deadline by which the task execution time can be adjusted. The time is in UTC and follows the format of YYYY-MM-DDTHH:mm:ssZ.
	//
	// example:
	//
	// 2018-05-30T23:59:59Z
	Deadline *string `json:"Deadline,omitempty" xml:"Deadline,omitempty"`
	// The task ID.
	//
	// example:
	//
	// 11111
	Id *int32 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The event impact.
	//
	// example:
	//
	// TransientDisconnection
	Impact *string `json:"Impact,omitempty" xml:"Impact,omitempty"`
	// The event impact in English.
	//
	// example:
	//
	// Transient instance disconnection
	ImpactEn *string `json:"ImpactEn,omitempty" xml:"ImpactEn,omitempty"`
	// The event impact in Chinese.
	//
	// example:
	//
	// Instance interruption
	ImpactZh *string `json:"ImpactZh,omitempty" xml:"ImpactZh,omitempty"`
	// The instance alias or instance description.
	//
	// example:
	//
	// test
	InsComment *string `json:"InsComment,omitempty" xml:"InsComment,omitempty"`
	// The instance name.
	//
	// example:
	//
	// rm-wz96h8jujh512****
	InsName *string `json:"InsName,omitempty" xml:"InsName,omitempty"`
	// The modification time. The time is in UTC and follows the format of YYYY-MM-DDTHH:mm:ssZ.
	//
	// example:
	//
	// 2018-05-30T14:30:00Z
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// The preparation time required between the start time and the switchover time. The format is HH:mm:ss.
	//
	// example:
	//
	// 04:00:00
	PrepareInterval *string `json:"PrepareInterval,omitempty" xml:"PrepareInterval,omitempty"`
	// The region ID of the pending event.
	//
	// example:
	//
	// cn-beijing
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The execution result information.
	//
	// example:
	//
	// userCancel
	ResultInfo *string `json:"ResultInfo,omitempty" xml:"ResultInfo,omitempty"`
	// The time when the backend executes the task. The time is in UTC and follows the format of YYYY-MM-DDTHH:mm:ssZ.
	//
	// example:
	//
	// 2018-05-30T00:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The task status. Valid values:
	//
	// 	- **3**: pending.
	//
	// 	- **4**: in progress.
	//
	// 	- **5**: succeeded.
	//
	// 	- **6**: failed.
	//
	// 	- **7**: canceled.
	//
	// example:
	//
	// 3
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The instance shards.
	SubInsNames []*string `json:"SubInsNames,omitempty" xml:"SubInsNames,omitempty" type:"Repeated"`
	// The time when the backend initiates the switchover. The time is in UTC and follows the format of YYYY-MM-DDTHH:mm:ssZ.
	//
	// example:
	//
	// 2018-05-30T14:30:00Z
	SwitchTime *string `json:"SwitchTime,omitempty" xml:"SwitchTime,omitempty"`
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
	// The task type. Valid values:
	//
	// 	- **rds_apsaradb_ha**: primary/secondary node switch.
	//
	// 	- **rds_apsaradb_transfer**: instance migration.
	//
	// 	- **rds_apsaradb_upgrade**: minor engine version update.
	//
	// 	- **rds_apsaradb_maxscale**: proxy minor version upgrade.
	//
	// example:
	//
	// rds_apsaradb_upgrade
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
