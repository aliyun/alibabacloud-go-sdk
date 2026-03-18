// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBackupsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*DescribeBackupsResponseBodyData) *DescribeBackupsResponseBody
	GetData() []*DescribeBackupsResponseBodyData
	SetErrCode(v string) *DescribeBackupsResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *DescribeBackupsResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *DescribeBackupsResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *DescribeBackupsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeBackupsResponseBody
	GetSuccess() *bool
	SetTotal(v int32) *DescribeBackupsResponseBody
	GetTotal() *int32
}

type DescribeBackupsResponseBody struct {
	Data []*DescribeBackupsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// Success
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// example:
	//
	// Invalid params: [instance not exists].
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// 32A44F0D-BFF6-5664-999A-218BBDE7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 1
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeBackupsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBackupsResponseBody) GetData() []*DescribeBackupsResponseBodyData {
	return s.Data
}

func (s *DescribeBackupsResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *DescribeBackupsResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *DescribeBackupsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DescribeBackupsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeBackupsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeBackupsResponseBody) GetTotal() *int32 {
	return s.Total
}

func (s *DescribeBackupsResponseBody) SetData(v []*DescribeBackupsResponseBodyData) *DescribeBackupsResponseBody {
	s.Data = v
	return s
}

func (s *DescribeBackupsResponseBody) SetErrCode(v string) *DescribeBackupsResponseBody {
	s.ErrCode = &v
	return s
}

func (s *DescribeBackupsResponseBody) SetErrMessage(v string) *DescribeBackupsResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *DescribeBackupsResponseBody) SetHttpStatusCode(v int32) *DescribeBackupsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeBackupsResponseBody) SetRequestId(v string) *DescribeBackupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeBackupsResponseBody) SetSuccess(v bool) *DescribeBackupsResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeBackupsResponseBody) SetTotal(v int32) *DescribeBackupsResponseBody {
	s.Total = &v
	return s
}

func (s *DescribeBackupsResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeBackupsResponseBodyData struct {
	// example:
	//
	// 1742179028000
	BackupFinishedTime *int64 `json:"BackupFinishedTime,omitempty" xml:"BackupFinishedTime,omitempty"`
	// example:
	//
	// 1742179018000
	BackupStartTime *int64 `json:"BackupStartTime,omitempty" xml:"BackupStartTime,omitempty"`
	// example:
	//
	// bt-12sui21312dd
	BackupTaskId *string `json:"BackupTaskId,omitempty" xml:"BackupTaskId,omitempty"`
	// example:
	//
	// FullBackup
	BackupType  *string `json:"BackupType,omitempty" xml:"BackupType,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 1742189008000
	ExpireTime *int64 `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// example:
	//
	// c-d4be777ff5e8cXXX
	InstanceId       *string                                          `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceSnapshot *DescribeBackupsResponseBodyDataInstanceSnapshot `json:"InstanceSnapshot,omitempty" xml:"InstanceSnapshot,omitempty" type:"Struct"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// 100
	Size *int64 `json:"Size,omitempty" xml:"Size,omitempty"`
	// example:
	//
	// RUNNING
	Status   *string                                    `json:"Status,omitempty" xml:"Status,omitempty"`
	SubTasks []*DescribeBackupsResponseBodyDataSubTasks `json:"SubTasks,omitempty" xml:"SubTasks,omitempty" type:"Repeated"`
}

func (s DescribeBackupsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupsResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeBackupsResponseBodyData) GetBackupFinishedTime() *int64 {
	return s.BackupFinishedTime
}

func (s *DescribeBackupsResponseBodyData) GetBackupStartTime() *int64 {
	return s.BackupStartTime
}

func (s *DescribeBackupsResponseBodyData) GetBackupTaskId() *string {
	return s.BackupTaskId
}

func (s *DescribeBackupsResponseBodyData) GetBackupType() *string {
	return s.BackupType
}

func (s *DescribeBackupsResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *DescribeBackupsResponseBodyData) GetExpireTime() *int64 {
	return s.ExpireTime
}

func (s *DescribeBackupsResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeBackupsResponseBodyData) GetInstanceSnapshot() *DescribeBackupsResponseBodyDataInstanceSnapshot {
	return s.InstanceSnapshot
}

func (s *DescribeBackupsResponseBodyData) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeBackupsResponseBodyData) GetSize() *int64 {
	return s.Size
}

func (s *DescribeBackupsResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *DescribeBackupsResponseBodyData) GetSubTasks() []*DescribeBackupsResponseBodyDataSubTasks {
	return s.SubTasks
}

func (s *DescribeBackupsResponseBodyData) SetBackupFinishedTime(v int64) *DescribeBackupsResponseBodyData {
	s.BackupFinishedTime = &v
	return s
}

func (s *DescribeBackupsResponseBodyData) SetBackupStartTime(v int64) *DescribeBackupsResponseBodyData {
	s.BackupStartTime = &v
	return s
}

func (s *DescribeBackupsResponseBodyData) SetBackupTaskId(v string) *DescribeBackupsResponseBodyData {
	s.BackupTaskId = &v
	return s
}

func (s *DescribeBackupsResponseBodyData) SetBackupType(v string) *DescribeBackupsResponseBodyData {
	s.BackupType = &v
	return s
}

func (s *DescribeBackupsResponseBodyData) SetDescription(v string) *DescribeBackupsResponseBodyData {
	s.Description = &v
	return s
}

func (s *DescribeBackupsResponseBodyData) SetExpireTime(v int64) *DescribeBackupsResponseBodyData {
	s.ExpireTime = &v
	return s
}

func (s *DescribeBackupsResponseBodyData) SetInstanceId(v string) *DescribeBackupsResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *DescribeBackupsResponseBodyData) SetInstanceSnapshot(v *DescribeBackupsResponseBodyDataInstanceSnapshot) *DescribeBackupsResponseBodyData {
	s.InstanceSnapshot = v
	return s
}

func (s *DescribeBackupsResponseBodyData) SetRegionId(v string) *DescribeBackupsResponseBodyData {
	s.RegionId = &v
	return s
}

func (s *DescribeBackupsResponseBodyData) SetSize(v int64) *DescribeBackupsResponseBodyData {
	s.Size = &v
	return s
}

func (s *DescribeBackupsResponseBodyData) SetStatus(v string) *DescribeBackupsResponseBodyData {
	s.Status = &v
	return s
}

func (s *DescribeBackupsResponseBodyData) SetSubTasks(v []*DescribeBackupsResponseBodyDataSubTasks) *DescribeBackupsResponseBodyData {
	s.SubTasks = v
	return s
}

func (s *DescribeBackupsResponseBodyData) Validate() error {
	if s.InstanceSnapshot != nil {
		if err := s.InstanceSnapshot.Validate(); err != nil {
			return err
		}
	}
	if s.SubTasks != nil {
		for _, item := range s.SubTasks {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeBackupsResponseBodyDataInstanceSnapshot struct {
	// example:
	//
	// c-37708ec80b5****
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// example:
	//
	// 3.3.13-1.0-1.7.2
	MinorVersion *string                                                      `json:"MinorVersion,omitempty" xml:"MinorVersion,omitempty"`
	NodeGroups   []*DescribeBackupsResponseBodyDataInstanceSnapshotNodeGroups `json:"NodeGroups,omitempty" xml:"NodeGroups,omitempty" type:"Repeated"`
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// rg-aekzd7frphchx3a
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// example:
	//
	// shared_data
	RunMode *string `json:"RunMode,omitempty" xml:"RunMode,omitempty"`
	// example:
	//
	// standard
	SpecType *string                                                `json:"SpecType,omitempty" xml:"SpecType,omitempty"`
	Tags     []*DescribeBackupsResponseBodyDataInstanceSnapshotTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// example:
	//
	// 3.3
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
	// VPC ID。
	//
	// example:
	//
	// vpc-2ze0cez8106f2n85c2d7i
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s DescribeBackupsResponseBodyDataInstanceSnapshot) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupsResponseBodyDataInstanceSnapshot) GoString() string {
	return s.String()
}

func (s *DescribeBackupsResponseBodyDataInstanceSnapshot) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DescribeBackupsResponseBodyDataInstanceSnapshot) GetMinorVersion() *string {
	return s.MinorVersion
}

func (s *DescribeBackupsResponseBodyDataInstanceSnapshot) GetNodeGroups() []*DescribeBackupsResponseBodyDataInstanceSnapshotNodeGroups {
	return s.NodeGroups
}

func (s *DescribeBackupsResponseBodyDataInstanceSnapshot) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeBackupsResponseBodyDataInstanceSnapshot) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeBackupsResponseBodyDataInstanceSnapshot) GetRunMode() *string {
	return s.RunMode
}

func (s *DescribeBackupsResponseBodyDataInstanceSnapshot) GetSpecType() *string {
	return s.SpecType
}

func (s *DescribeBackupsResponseBodyDataInstanceSnapshot) GetTags() []*DescribeBackupsResponseBodyDataInstanceSnapshotTags {
	return s.Tags
}

func (s *DescribeBackupsResponseBodyDataInstanceSnapshot) GetVersion() *string {
	return s.Version
}

func (s *DescribeBackupsResponseBodyDataInstanceSnapshot) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeBackupsResponseBodyDataInstanceSnapshot) SetInstanceName(v string) *DescribeBackupsResponseBodyDataInstanceSnapshot {
	s.InstanceName = &v
	return s
}

func (s *DescribeBackupsResponseBodyDataInstanceSnapshot) SetMinorVersion(v string) *DescribeBackupsResponseBodyDataInstanceSnapshot {
	s.MinorVersion = &v
	return s
}

func (s *DescribeBackupsResponseBodyDataInstanceSnapshot) SetNodeGroups(v []*DescribeBackupsResponseBodyDataInstanceSnapshotNodeGroups) *DescribeBackupsResponseBodyDataInstanceSnapshot {
	s.NodeGroups = v
	return s
}

func (s *DescribeBackupsResponseBodyDataInstanceSnapshot) SetRegionId(v string) *DescribeBackupsResponseBodyDataInstanceSnapshot {
	s.RegionId = &v
	return s
}

func (s *DescribeBackupsResponseBodyDataInstanceSnapshot) SetResourceGroupId(v string) *DescribeBackupsResponseBodyDataInstanceSnapshot {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeBackupsResponseBodyDataInstanceSnapshot) SetRunMode(v string) *DescribeBackupsResponseBodyDataInstanceSnapshot {
	s.RunMode = &v
	return s
}

func (s *DescribeBackupsResponseBodyDataInstanceSnapshot) SetSpecType(v string) *DescribeBackupsResponseBodyDataInstanceSnapshot {
	s.SpecType = &v
	return s
}

func (s *DescribeBackupsResponseBodyDataInstanceSnapshot) SetTags(v []*DescribeBackupsResponseBodyDataInstanceSnapshotTags) *DescribeBackupsResponseBodyDataInstanceSnapshot {
	s.Tags = v
	return s
}

func (s *DescribeBackupsResponseBodyDataInstanceSnapshot) SetVersion(v string) *DescribeBackupsResponseBodyDataInstanceSnapshot {
	s.Version = &v
	return s
}

func (s *DescribeBackupsResponseBodyDataInstanceSnapshot) SetVpcId(v string) *DescribeBackupsResponseBodyDataInstanceSnapshot {
	s.VpcId = &v
	return s
}

func (s *DescribeBackupsResponseBodyDataInstanceSnapshot) Validate() error {
	if s.NodeGroups != nil {
		for _, item := range s.NodeGroups {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeBackupsResponseBodyDataInstanceSnapshotNodeGroups struct {
	// example:
	//
	// FE
	ComponentType *string `json:"ComponentType,omitempty" xml:"ComponentType,omitempty"`
	// example:
	//
	// 4
	Cu *int32 `json:"Cu,omitempty" xml:"Cu,omitempty"`
	// example:
	//
	// 2
	DiskNumber *string `json:"DiskNumber,omitempty" xml:"DiskNumber,omitempty"`
	// example:
	//
	// null
	LocalStorageInstanceType *string `json:"LocalStorageInstanceType,omitempty" xml:"LocalStorageInstanceType,omitempty"`
	// example:
	//
	// 3
	ResidentNodeNumber *string `json:"ResidentNodeNumber,omitempty" xml:"ResidentNodeNumber,omitempty"`
	// example:
	//
	// standard
	SpecType *string `json:"SpecType,omitempty" xml:"SpecType,omitempty"`
	// example:
	//
	// pl1
	StoragePerformanceLevel *string `json:"StoragePerformanceLevel,omitempty" xml:"StoragePerformanceLevel,omitempty"`
	// example:
	//
	// 200
	StorageSize *int32 `json:"StorageSize,omitempty" xml:"StorageSize,omitempty"`
}

func (s DescribeBackupsResponseBodyDataInstanceSnapshotNodeGroups) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupsResponseBodyDataInstanceSnapshotNodeGroups) GoString() string {
	return s.String()
}

func (s *DescribeBackupsResponseBodyDataInstanceSnapshotNodeGroups) GetComponentType() *string {
	return s.ComponentType
}

func (s *DescribeBackupsResponseBodyDataInstanceSnapshotNodeGroups) GetCu() *int32 {
	return s.Cu
}

func (s *DescribeBackupsResponseBodyDataInstanceSnapshotNodeGroups) GetDiskNumber() *string {
	return s.DiskNumber
}

func (s *DescribeBackupsResponseBodyDataInstanceSnapshotNodeGroups) GetLocalStorageInstanceType() *string {
	return s.LocalStorageInstanceType
}

func (s *DescribeBackupsResponseBodyDataInstanceSnapshotNodeGroups) GetResidentNodeNumber() *string {
	return s.ResidentNodeNumber
}

func (s *DescribeBackupsResponseBodyDataInstanceSnapshotNodeGroups) GetSpecType() *string {
	return s.SpecType
}

func (s *DescribeBackupsResponseBodyDataInstanceSnapshotNodeGroups) GetStoragePerformanceLevel() *string {
	return s.StoragePerformanceLevel
}

func (s *DescribeBackupsResponseBodyDataInstanceSnapshotNodeGroups) GetStorageSize() *int32 {
	return s.StorageSize
}

func (s *DescribeBackupsResponseBodyDataInstanceSnapshotNodeGroups) SetComponentType(v string) *DescribeBackupsResponseBodyDataInstanceSnapshotNodeGroups {
	s.ComponentType = &v
	return s
}

func (s *DescribeBackupsResponseBodyDataInstanceSnapshotNodeGroups) SetCu(v int32) *DescribeBackupsResponseBodyDataInstanceSnapshotNodeGroups {
	s.Cu = &v
	return s
}

func (s *DescribeBackupsResponseBodyDataInstanceSnapshotNodeGroups) SetDiskNumber(v string) *DescribeBackupsResponseBodyDataInstanceSnapshotNodeGroups {
	s.DiskNumber = &v
	return s
}

func (s *DescribeBackupsResponseBodyDataInstanceSnapshotNodeGroups) SetLocalStorageInstanceType(v string) *DescribeBackupsResponseBodyDataInstanceSnapshotNodeGroups {
	s.LocalStorageInstanceType = &v
	return s
}

func (s *DescribeBackupsResponseBodyDataInstanceSnapshotNodeGroups) SetResidentNodeNumber(v string) *DescribeBackupsResponseBodyDataInstanceSnapshotNodeGroups {
	s.ResidentNodeNumber = &v
	return s
}

func (s *DescribeBackupsResponseBodyDataInstanceSnapshotNodeGroups) SetSpecType(v string) *DescribeBackupsResponseBodyDataInstanceSnapshotNodeGroups {
	s.SpecType = &v
	return s
}

func (s *DescribeBackupsResponseBodyDataInstanceSnapshotNodeGroups) SetStoragePerformanceLevel(v string) *DescribeBackupsResponseBodyDataInstanceSnapshotNodeGroups {
	s.StoragePerformanceLevel = &v
	return s
}

func (s *DescribeBackupsResponseBodyDataInstanceSnapshotNodeGroups) SetStorageSize(v int32) *DescribeBackupsResponseBodyDataInstanceSnapshotNodeGroups {
	s.StorageSize = &v
	return s
}

func (s *DescribeBackupsResponseBodyDataInstanceSnapshotNodeGroups) Validate() error {
	return dara.Validate(s)
}

type DescribeBackupsResponseBodyDataInstanceSnapshotTags struct {
	// example:
	//
	// bk-time
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// 1747708000
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeBackupsResponseBodyDataInstanceSnapshotTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupsResponseBodyDataInstanceSnapshotTags) GoString() string {
	return s.String()
}

func (s *DescribeBackupsResponseBodyDataInstanceSnapshotTags) GetKey() *string {
	return s.Key
}

func (s *DescribeBackupsResponseBodyDataInstanceSnapshotTags) GetValue() *string {
	return s.Value
}

func (s *DescribeBackupsResponseBodyDataInstanceSnapshotTags) SetKey(v string) *DescribeBackupsResponseBodyDataInstanceSnapshotTags {
	s.Key = &v
	return s
}

func (s *DescribeBackupsResponseBodyDataInstanceSnapshotTags) SetValue(v string) *DescribeBackupsResponseBodyDataInstanceSnapshotTags {
	s.Value = &v
	return s
}

func (s *DescribeBackupsResponseBodyDataInstanceSnapshotTags) Validate() error {
	return dara.Validate(s)
}

type DescribeBackupsResponseBodyDataSubTasks struct {
	// example:
	//
	// test1
	DataBase *string `json:"DataBase,omitempty" xml:"DataBase,omitempty"`
	// example:
	//
	// []
	Detail *string `json:"Detail,omitempty" xml:"Detail,omitempty"`
	// example:
	//
	// 1747718190
	FinishedTime *int64 `json:"FinishedTime,omitempty" xml:"FinishedTime,omitempty"`
	// example:
	//
	// 1
	Size *int64 `json:"Size,omitempty" xml:"Size,omitempty"`
	// example:
	//
	// 2025-02-10_backup
	SnapshotName *string `json:"SnapshotName,omitempty" xml:"SnapshotName,omitempty"`
	// example:
	//
	// 1747708190
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// cdc_ods_t2030_lcpf_api_topic_msg
	Table *string `json:"Table,omitempty" xml:"Table,omitempty"`
}

func (s DescribeBackupsResponseBodyDataSubTasks) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupsResponseBodyDataSubTasks) GoString() string {
	return s.String()
}

func (s *DescribeBackupsResponseBodyDataSubTasks) GetDataBase() *string {
	return s.DataBase
}

func (s *DescribeBackupsResponseBodyDataSubTasks) GetDetail() *string {
	return s.Detail
}

func (s *DescribeBackupsResponseBodyDataSubTasks) GetFinishedTime() *int64 {
	return s.FinishedTime
}

func (s *DescribeBackupsResponseBodyDataSubTasks) GetSize() *int64 {
	return s.Size
}

func (s *DescribeBackupsResponseBodyDataSubTasks) GetSnapshotName() *string {
	return s.SnapshotName
}

func (s *DescribeBackupsResponseBodyDataSubTasks) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeBackupsResponseBodyDataSubTasks) GetStatus() *string {
	return s.Status
}

func (s *DescribeBackupsResponseBodyDataSubTasks) GetTable() *string {
	return s.Table
}

func (s *DescribeBackupsResponseBodyDataSubTasks) SetDataBase(v string) *DescribeBackupsResponseBodyDataSubTasks {
	s.DataBase = &v
	return s
}

func (s *DescribeBackupsResponseBodyDataSubTasks) SetDetail(v string) *DescribeBackupsResponseBodyDataSubTasks {
	s.Detail = &v
	return s
}

func (s *DescribeBackupsResponseBodyDataSubTasks) SetFinishedTime(v int64) *DescribeBackupsResponseBodyDataSubTasks {
	s.FinishedTime = &v
	return s
}

func (s *DescribeBackupsResponseBodyDataSubTasks) SetSize(v int64) *DescribeBackupsResponseBodyDataSubTasks {
	s.Size = &v
	return s
}

func (s *DescribeBackupsResponseBodyDataSubTasks) SetSnapshotName(v string) *DescribeBackupsResponseBodyDataSubTasks {
	s.SnapshotName = &v
	return s
}

func (s *DescribeBackupsResponseBodyDataSubTasks) SetStartTime(v int64) *DescribeBackupsResponseBodyDataSubTasks {
	s.StartTime = &v
	return s
}

func (s *DescribeBackupsResponseBodyDataSubTasks) SetStatus(v string) *DescribeBackupsResponseBodyDataSubTasks {
	s.Status = &v
	return s
}

func (s *DescribeBackupsResponseBodyDataSubTasks) SetTable(v string) *DescribeBackupsResponseBodyDataSubTasks {
	s.Table = &v
	return s
}

func (s *DescribeBackupsResponseBodyDataSubTasks) Validate() error {
	return dara.Validate(s)
}
