// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBackupDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBackupDataList(v []*ListBackupDataResponseBodyBackupDataList) *ListBackupDataResponseBody
	GetBackupDataList() []*ListBackupDataResponseBodyBackupDataList
	SetRequestId(v string) *ListBackupDataResponseBody
	GetRequestId() *string
}

type ListBackupDataResponseBody struct {
	// The backups.
	BackupDataList []*ListBackupDataResponseBodyBackupDataList `json:"BackupDataList,omitempty" xml:"BackupDataList,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 4AA0C48F-B5BB-5FF9-A43B-6B91E0715D46
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListBackupDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListBackupDataResponseBody) GoString() string {
	return s.String()
}

func (s *ListBackupDataResponseBody) GetBackupDataList() []*ListBackupDataResponseBodyBackupDataList {
	return s.BackupDataList
}

func (s *ListBackupDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListBackupDataResponseBody) SetBackupDataList(v []*ListBackupDataResponseBodyBackupDataList) *ListBackupDataResponseBody {
	s.BackupDataList = v
	return s
}

func (s *ListBackupDataResponseBody) SetRequestId(v string) *ListBackupDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListBackupDataResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListBackupDataResponseBodyBackupDataList struct {
	// The backup type. In general, the following two types are supported: local backup and remote backup. In the local backup type, snapshots reside in the same region as your instance. The following two sub-types are available: full (single backup, single replica) and redundant (zone-redundant storage, multiple replicas). In the remote backup type, snapshots and your instance reside in different regions. Remote backups are the replicas of the backups of the full or redundant type in another region. The values local and remote do not represent specific types, but are used only for data filtering. The value local indicates all local backups, and the value remote indicates all remote backups.
	//
	// example:
	//
	// redundant
	BackupType *string `json:"BackupType,omitempty" xml:"BackupType,omitempty"`
	// The size of cold data. Unit: bytes.
	//
	// example:
	//
	// 32413521
	ColdDataSize *int64 `json:"ColdDataSize,omitempty" xml:"ColdDataSize,omitempty"`
	// The description of the backup data.
	//
	// example:
	//
	// demo
	DataDesc *string `json:"DataDesc,omitempty" xml:"DataDesc,omitempty"`
	// The backup granularity.
	//
	// Valid values:
	//
	// 	- instance
	//
	// example:
	//
	// instance
	DataGran *string `json:"DataGran,omitempty" xml:"DataGran,omitempty"`
	// The size of the backup data. Unit: bytes.
	//
	// example:
	//
	// 76085723136
	DataSize *int64 `json:"DataSize,omitempty" xml:"DataSize,omitempty"`
	// The snapshot time. The value format of this parameter follows the same standard as that of the StartTime parameter.
	//
	// example:
	//
	// 2024-10-28T12:23:37.000+00:00
	DataTime *string `json:"DataTime,omitempty" xml:"DataTime,omitempty"`
	// The end time of the backup task. The value format of this parameter follows the same standard as that of the StartTime parameter.
	//
	// example:
	//
	// 2024-10-28T12:27:34.000+00:00
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The unique ID of the backup.
	//
	// example:
	//
	// 1780805690994479105
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// hgpostcn-cn-pe33jdxxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the instance.
	//
	// example:
	//
	// my-hologres-dw
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The region in which the instance resides.
	//
	// example:
	//
	// cn-hangzhou
	InstanceRegion *string `json:"InstanceRegion,omitempty" xml:"InstanceRegion,omitempty"`
	// The type of the instance.
	//
	// Valid values:
	//
	// 	- Warehouse: virtual warehouse instance
	//
	// 	- Standard: general-purpose instance
	//
	// example:
	//
	// Warehouse
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The zone in which the instance resides.
	//
	// example:
	//
	// cn-hangzhou-j
	InstanceZoneId *string `json:"InstanceZoneId,omitempty" xml:"InstanceZoneId,omitempty"`
	// The region in which the backup data resides.
	//
	// example:
	//
	// cn-hangzhou
	SnapshotRegion *string `json:"SnapshotRegion,omitempty" xml:"SnapshotRegion,omitempty"`
	// The zone in which the backup data resides. In zone-redundant storage mode, backup data is stored in different zones, including the current zone.
	//
	// example:
	//
	// cn-hangzhou-j
	SnapshotZoneId *string `json:"SnapshotZoneId,omitempty" xml:"SnapshotZoneId,omitempty"`
	// The start time of the backup task. The time follows the ISO 8601 standard in the YYYY-MM-DDTHH:mm:ss.SSSTZ format. The time is displayed in UTC (the same below).
	//
	// example:
	//
	// 2024-10-28T11:19:56.000+00:00
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The status of the backup task.
	//
	// Valid values:
	//
	// 	- processing
	//
	// 	- completed
	//
	// 	- failed
	//
	// example:
	//
	// completed
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The mode in which the backup task is triggered.
	//
	// Valid values:
	//
	// 	- scheduled: periodic backup
	//
	// 	- manual: manual backup
	//
	// example:
	//
	// scheduled
	TriggerType *string `json:"TriggerType,omitempty" xml:"TriggerType,omitempty"`
}

func (s ListBackupDataResponseBodyBackupDataList) String() string {
	return dara.Prettify(s)
}

func (s ListBackupDataResponseBodyBackupDataList) GoString() string {
	return s.String()
}

func (s *ListBackupDataResponseBodyBackupDataList) GetBackupType() *string {
	return s.BackupType
}

func (s *ListBackupDataResponseBodyBackupDataList) GetColdDataSize() *int64 {
	return s.ColdDataSize
}

func (s *ListBackupDataResponseBodyBackupDataList) GetDataDesc() *string {
	return s.DataDesc
}

func (s *ListBackupDataResponseBodyBackupDataList) GetDataGran() *string {
	return s.DataGran
}

func (s *ListBackupDataResponseBodyBackupDataList) GetDataSize() *int64 {
	return s.DataSize
}

func (s *ListBackupDataResponseBodyBackupDataList) GetDataTime() *string {
	return s.DataTime
}

func (s *ListBackupDataResponseBodyBackupDataList) GetEndTime() *string {
	return s.EndTime
}

func (s *ListBackupDataResponseBodyBackupDataList) GetId() *int64 {
	return s.Id
}

func (s *ListBackupDataResponseBodyBackupDataList) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListBackupDataResponseBodyBackupDataList) GetInstanceName() *string {
	return s.InstanceName
}

func (s *ListBackupDataResponseBodyBackupDataList) GetInstanceRegion() *string {
	return s.InstanceRegion
}

func (s *ListBackupDataResponseBodyBackupDataList) GetInstanceType() *string {
	return s.InstanceType
}

func (s *ListBackupDataResponseBodyBackupDataList) GetInstanceZoneId() *string {
	return s.InstanceZoneId
}

func (s *ListBackupDataResponseBodyBackupDataList) GetSnapshotRegion() *string {
	return s.SnapshotRegion
}

func (s *ListBackupDataResponseBodyBackupDataList) GetSnapshotZoneId() *string {
	return s.SnapshotZoneId
}

func (s *ListBackupDataResponseBodyBackupDataList) GetStartTime() *string {
	return s.StartTime
}

func (s *ListBackupDataResponseBodyBackupDataList) GetStatus() *string {
	return s.Status
}

func (s *ListBackupDataResponseBodyBackupDataList) GetTriggerType() *string {
	return s.TriggerType
}

func (s *ListBackupDataResponseBodyBackupDataList) SetBackupType(v string) *ListBackupDataResponseBodyBackupDataList {
	s.BackupType = &v
	return s
}

func (s *ListBackupDataResponseBodyBackupDataList) SetColdDataSize(v int64) *ListBackupDataResponseBodyBackupDataList {
	s.ColdDataSize = &v
	return s
}

func (s *ListBackupDataResponseBodyBackupDataList) SetDataDesc(v string) *ListBackupDataResponseBodyBackupDataList {
	s.DataDesc = &v
	return s
}

func (s *ListBackupDataResponseBodyBackupDataList) SetDataGran(v string) *ListBackupDataResponseBodyBackupDataList {
	s.DataGran = &v
	return s
}

func (s *ListBackupDataResponseBodyBackupDataList) SetDataSize(v int64) *ListBackupDataResponseBodyBackupDataList {
	s.DataSize = &v
	return s
}

func (s *ListBackupDataResponseBodyBackupDataList) SetDataTime(v string) *ListBackupDataResponseBodyBackupDataList {
	s.DataTime = &v
	return s
}

func (s *ListBackupDataResponseBodyBackupDataList) SetEndTime(v string) *ListBackupDataResponseBodyBackupDataList {
	s.EndTime = &v
	return s
}

func (s *ListBackupDataResponseBodyBackupDataList) SetId(v int64) *ListBackupDataResponseBodyBackupDataList {
	s.Id = &v
	return s
}

func (s *ListBackupDataResponseBodyBackupDataList) SetInstanceId(v string) *ListBackupDataResponseBodyBackupDataList {
	s.InstanceId = &v
	return s
}

func (s *ListBackupDataResponseBodyBackupDataList) SetInstanceName(v string) *ListBackupDataResponseBodyBackupDataList {
	s.InstanceName = &v
	return s
}

func (s *ListBackupDataResponseBodyBackupDataList) SetInstanceRegion(v string) *ListBackupDataResponseBodyBackupDataList {
	s.InstanceRegion = &v
	return s
}

func (s *ListBackupDataResponseBodyBackupDataList) SetInstanceType(v string) *ListBackupDataResponseBodyBackupDataList {
	s.InstanceType = &v
	return s
}

func (s *ListBackupDataResponseBodyBackupDataList) SetInstanceZoneId(v string) *ListBackupDataResponseBodyBackupDataList {
	s.InstanceZoneId = &v
	return s
}

func (s *ListBackupDataResponseBodyBackupDataList) SetSnapshotRegion(v string) *ListBackupDataResponseBodyBackupDataList {
	s.SnapshotRegion = &v
	return s
}

func (s *ListBackupDataResponseBodyBackupDataList) SetSnapshotZoneId(v string) *ListBackupDataResponseBodyBackupDataList {
	s.SnapshotZoneId = &v
	return s
}

func (s *ListBackupDataResponseBodyBackupDataList) SetStartTime(v string) *ListBackupDataResponseBodyBackupDataList {
	s.StartTime = &v
	return s
}

func (s *ListBackupDataResponseBodyBackupDataList) SetStatus(v string) *ListBackupDataResponseBodyBackupDataList {
	s.Status = &v
	return s
}

func (s *ListBackupDataResponseBodyBackupDataList) SetTriggerType(v string) *ListBackupDataResponseBodyBackupDataList {
	s.TriggerType = &v
	return s
}

func (s *ListBackupDataResponseBodyBackupDataList) Validate() error {
	return dara.Validate(s)
}
