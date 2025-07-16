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
	SetPageNumber(v int64) *DescribeActiveOperationTasksResponseBody
	GetPageNumber() *int64
	SetPageSize(v int64) *DescribeActiveOperationTasksResponseBody
	GetPageSize() *int64
	SetRequestId(v string) *DescribeActiveOperationTasksResponseBody
	GetRequestId() *string
	SetTotalRecordCount(v int64) *DescribeActiveOperationTasksResponseBody
	GetTotalRecordCount() *int64
}

type DescribeActiveOperationTasksResponseBody struct {
	Items []*DescribeActiveOperationTasksResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 12
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// xxxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 23
	TotalRecordCount *int64 `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
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

func (s *DescribeActiveOperationTasksResponseBody) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeActiveOperationTasksResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeActiveOperationTasksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeActiveOperationTasksResponseBody) GetTotalRecordCount() *int64 {
	return s.TotalRecordCount
}

func (s *DescribeActiveOperationTasksResponseBody) SetItems(v []*DescribeActiveOperationTasksResponseBodyItems) *DescribeActiveOperationTasksResponseBody {
	s.Items = v
	return s
}

func (s *DescribeActiveOperationTasksResponseBody) SetPageNumber(v int64) *DescribeActiveOperationTasksResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeActiveOperationTasksResponseBody) SetPageSize(v int64) *DescribeActiveOperationTasksResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeActiveOperationTasksResponseBody) SetRequestId(v string) *DescribeActiveOperationTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeActiveOperationTasksResponseBody) SetTotalRecordCount(v int64) *DescribeActiveOperationTasksResponseBody {
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
	AllowCancel *string `json:"AllowCancel,omitempty" xml:"AllowCancel,omitempty"`
	// example:
	//
	// 0
	AllowChange *string `json:"AllowChange,omitempty" xml:"AllowChange,omitempty"`
	// example:
	//
	// Risk
	ChangeLevel *string `json:"ChangeLevel,omitempty" xml:"ChangeLevel,omitempty"`
	// example:
	//
	// Risk repairment
	ChangeLevelEn *string `json:"ChangeLevelEn,omitempty" xml:"ChangeLevelEn,omitempty"`
	ChangeLevelZh *string `json:"ChangeLevelZh,omitempty" xml:"ChangeLevelZh,omitempty"`
	// example:
	//
	// 2021-06-15T16:00:00Z
	CreatedTime *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// example:
	//
	// cn-shanghai-et-a
	CurrentAVZ *string `json:"CurrentAVZ,omitempty" xml:"CurrentAVZ,omitempty"`
	// example:
	//
	// polarx
	DbType *string `json:"DbType,omitempty" xml:"DbType,omitempty"`
	// example:
	//
	// 2.0
	DbVersion *string `json:"DbVersion,omitempty" xml:"DbVersion,omitempty"`
	// example:
	//
	// 2021-06-15T16:00:00Z
	Deadline *string `json:"Deadline,omitempty" xml:"Deadline,omitempty"`
	// example:
	//
	// 1
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// --
	Impact *string `json:"Impact,omitempty" xml:"Impact,omitempty"`
	// example:
	//
	// Transient instance disconnection
	ImpactEn *string `json:"ImpactEn,omitempty" xml:"ImpactEn,omitempty"`
	ImpactZh *string `json:"ImpactZh,omitempty" xml:"ImpactZh,omitempty"`
	// example:
	//
	// xxx
	InsComment *string `json:"InsComment,omitempty" xml:"InsComment,omitempty"`
	// example:
	//
	// pxc-dd
	InsName *string `json:"InsName,omitempty" xml:"InsName,omitempty"`
	// example:
	//
	// 2021-06-15T16:00:00Z
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// example:
	//
	// 03:00:00
	PrepareInterval *string `json:"PrepareInterval,omitempty" xml:"PrepareInterval,omitempty"`
	// example:
	//
	// cn-shanghai-et15-b01
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// example:
	//
	// --
	ResultInfo *string `json:"ResultInfo,omitempty" xml:"ResultInfo,omitempty"`
	// example:
	//
	// 2021-06-15T16:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// 3
	Status      *int64    `json:"Status,omitempty" xml:"Status,omitempty"`
	SubInsNames []*string `json:"SubInsNames,omitempty" xml:"SubInsNames,omitempty" type:"Repeated"`
	// example:
	//
	// 2021-06-15T16:00:00Z
	SwitchTime *string `json:"SwitchTime,omitempty" xml:"SwitchTime,omitempty"`
	// example:
	//
	// rds_apsaradb_transfer
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	// example:
	//
	// --
	TaskTypeEn *string `json:"TaskTypeEn,omitempty" xml:"TaskTypeEn,omitempty"`
	// example:
	//
	// --
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

func (s *DescribeActiveOperationTasksResponseBodyItems) GetId() *int64 {
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

func (s *DescribeActiveOperationTasksResponseBodyItems) GetStatus() *int64 {
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

func (s *DescribeActiveOperationTasksResponseBodyItems) SetId(v int64) *DescribeActiveOperationTasksResponseBodyItems {
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

func (s *DescribeActiveOperationTasksResponseBodyItems) SetStatus(v int64) *DescribeActiveOperationTasksResponseBodyItems {
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
