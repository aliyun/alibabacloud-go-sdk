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
	// The list of backups.
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
	if s.BackupDataList != nil {
		for _, item := range s.BackupDataList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListBackupDataResponseBodyBackupDataList struct {
	// The backup type. Backup types are generally divided into two categories: 1) Local backup, where the snapshot and the instance belong to the same region. 2) Remote backup, where the snapshot and the instance belong to different regions. Valid values:
	//
	// - redundant: Local backup with cross-zone multi-copy storage in the same region.
	//
	// - redundant_remote: Remote backup with cross-zone multi-copy storage in the same region, which is a cross-region copy of the redundant type.
	//
	// - full: Local backup with single-copy storage.
	//
	// - full_remote: Remote backup with single-copy storage, which is a cross-region copy of the full type.
	//
	// - local: Does not represent a specific backup type and is only used to filter local backup data.
	//
	// - remote: Does not represent a specific backup type and is only used to filter remote backup data.
	//
	// example:
	//
	// redundant
	BackupType *string `json:"BackupType,omitempty" xml:"BackupType,omitempty"`
	// The size of the cold storage data. Unit: bytes.
	//
	// example:
	//
	// 32413521
	ColdDataSize *int64 `json:"ColdDataSize,omitempty" xml:"ColdDataSize,omitempty"`
	// The description of the backup data. This field can be updated.
	//
	// example:
	//
	// demo
	DataDesc *string `json:"DataDesc,omitempty" xml:"DataDesc,omitempty"`
	// The backup granularity.
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
	// The data snapshot time. For the format, refer to StartTime.
	//
	// example:
	//
	// 2024-10-28T12:23:37.000+00:00
	DataTime *string `json:"DataTime,omitempty" xml:"DataTime,omitempty"`
	// The end time of the backup. For the format, refer to StartTime.
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
	// The instance name.
	//
	// example:
	//
	// my-hologres-dw
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The region where the instance resides.
	//
	// example:
	//
	// cn-hangzhou
	InstanceRegion *string `json:"InstanceRegion,omitempty" xml:"InstanceRegion,omitempty"`
	// The instance type.
	//
	// example:
	//
	// Warehouse
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The zone where the instance resides.
	//
	// example:
	//
	// cn-hangzhou-j
	InstanceZoneId *string `json:"InstanceZoneId,omitempty" xml:"InstanceZoneId,omitempty"`
	// The region where the backup data resides.
	//
	// example:
	//
	// cn-hangzhou
	SnapshotRegion *string `json:"SnapshotRegion,omitempty" xml:"SnapshotRegion,omitempty"`
	// The zone where the backup data resides. For zone-redundant storage, the data is stored across multiple zones, and this field indicates one of the zones.
	//
	// example:
	//
	// cn-hangzhou-j
	SnapshotZoneId *string `json:"SnapshotZoneId,omitempty" xml:"SnapshotZoneId,omitempty"`
	// The start time of the backup. The time is in the UTC+8 time zone and follows the YYYY-MM-DDTHH:mm:ss.SSSTZ format. The same applies to other time fields.
	//
	// example:
	//
	// 2024-10-28T11:19:56.000+00:00
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The backup status.
	//
	// example:
	//
	// completed
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The trigger type.
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
