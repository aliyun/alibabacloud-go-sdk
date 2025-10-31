// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBackupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBackup(v *GetBackupResponseBodyBackup) *GetBackupResponseBody
	GetBackup() *GetBackupResponseBodyBackup
	SetRequestId(v string) *GetBackupResponseBody
	GetRequestId() *string
}

type GetBackupResponseBody struct {
	// The information about the backup.
	Backup *GetBackupResponseBodyBackup `json:"Backup,omitempty" xml:"Backup,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 4C467B38-3910-447D-87BC-AC049166F216
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetBackupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetBackupResponseBody) GoString() string {
	return s.String()
}

func (s *GetBackupResponseBody) GetBackup() *GetBackupResponseBodyBackup {
	return s.Backup
}

func (s *GetBackupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetBackupResponseBody) SetBackup(v *GetBackupResponseBodyBackup) *GetBackupResponseBody {
	s.Backup = v
	return s
}

func (s *GetBackupResponseBody) SetRequestId(v string) *GetBackupResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetBackupResponseBody) Validate() error {
	if s.Backup != nil {
		if err := s.Backup.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetBackupResponseBodyBackup struct {
	// The number of images that are automatically backed up.
	//
	// example:
	//
	// 1
	AutoImageCount *int64 `json:"AutoImageCount,omitempty" xml:"AutoImageCount,omitempty"`
	// The backup time in the 24-hour format.
	//
	// example:
	//
	// 10
	BackupHourInDay *string `json:"BackupHourInDay,omitempty" xml:"BackupHourInDay,omitempty"`
	// The ID of the backup.
	//
	// example:
	//
	// backup-fdb897sdf****
	BackupId *string `json:"BackupId,omitempty" xml:"BackupId,omitempty"`
	// The automatic backup cycle. Unit: days.
	//
	// example:
	//
	// 3
	BackupPeriod *int64 `json:"BackupPeriod,omitempty" xml:"BackupPeriod,omitempty"`
	// The time when the backup is created. The value is accurate to the millisecond. The value is a UNIX timestamp.
	//
	// example:
	//
	// 1682417553781
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The expiration time of the backup. The value is accurate to the millisecond. The value is a UNIX timestamp.
	//
	// example:
	//
	// 1682417553781
	ExpireTime *int64 `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The ID of the hardware security module (HSM) that is associated with the backup.
	//
	// example:
	//
	// hsm-cn-5yd35431****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The maximum number of images.
	//
	// example:
	//
	// 3
	MaxImageCount *string `json:"MaxImageCount,omitempty" xml:"MaxImageCount,omitempty"`
	// The name of the backup.
	//
	// example:
	//
	// backup-te****
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The next time when the image is created. The value is accurate to the millisecond. The value is a UNIX timestamp.
	//
	// example:
	//
	// 1682417553781
	NextImageCreateTime *int64 `json:"NextImageCreateTime,omitempty" xml:"NextImageCreateTime,omitempty"`
	// The ID of the HSM to which the backup belongs. This parameter is available only for HSM backups outside the Chinese mainland and the value of this parameter is consistent with the value of InstanceId.
	//
	// example:
	//
	// hsm-cn-huoahd****
	OwnerInstanceId *string `json:"OwnerInstanceId,omitempty" xml:"OwnerInstanceId,omitempty"`
	// The ID of the region.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The time when the backup is released. The value is accurate to the millisecond. The value is a UNIX timestamp.
	//
	// example:
	//
	// 1641275680000
	ReleaseTime *int64 `json:"ReleaseTime,omitempty" xml:"ReleaseTime,omitempty"`
	// The description of the backup.
	//
	// example:
	//
	// test
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The ID of the backup. This parameter is available only for HSM backups in the Chinese mainland.
	//
	// example:
	//
	// backup-fdb897sdfg53****
	SpInstanceId *string `json:"SpInstanceId,omitempty" xml:"SpInstanceId,omitempty"`
	// The status of the backup. Valid values:
	//
	// 	- NEW
	//
	// 	- EXPIRED
	//
	// 	- ENABLED
	//
	// example:
	//
	// NEW
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The type of the backup. Valid values:
	//
	// 	- DEFAULT
	//
	// 	- NORMAL
	//
	// example:
	//
	// DEFAULT
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetBackupResponseBodyBackup) String() string {
	return dara.Prettify(s)
}

func (s GetBackupResponseBodyBackup) GoString() string {
	return s.String()
}

func (s *GetBackupResponseBodyBackup) GetAutoImageCount() *int64 {
	return s.AutoImageCount
}

func (s *GetBackupResponseBodyBackup) GetBackupHourInDay() *string {
	return s.BackupHourInDay
}

func (s *GetBackupResponseBodyBackup) GetBackupId() *string {
	return s.BackupId
}

func (s *GetBackupResponseBodyBackup) GetBackupPeriod() *int64 {
	return s.BackupPeriod
}

func (s *GetBackupResponseBodyBackup) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *GetBackupResponseBodyBackup) GetExpireTime() *int64 {
	return s.ExpireTime
}

func (s *GetBackupResponseBodyBackup) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetBackupResponseBodyBackup) GetMaxImageCount() *string {
	return s.MaxImageCount
}

func (s *GetBackupResponseBodyBackup) GetName() *string {
	return s.Name
}

func (s *GetBackupResponseBodyBackup) GetNextImageCreateTime() *int64 {
	return s.NextImageCreateTime
}

func (s *GetBackupResponseBodyBackup) GetOwnerInstanceId() *string {
	return s.OwnerInstanceId
}

func (s *GetBackupResponseBodyBackup) GetRegionId() *string {
	return s.RegionId
}

func (s *GetBackupResponseBodyBackup) GetReleaseTime() *int64 {
	return s.ReleaseTime
}

func (s *GetBackupResponseBodyBackup) GetRemark() *string {
	return s.Remark
}

func (s *GetBackupResponseBodyBackup) GetSpInstanceId() *string {
	return s.SpInstanceId
}

func (s *GetBackupResponseBodyBackup) GetStatus() *string {
	return s.Status
}

func (s *GetBackupResponseBodyBackup) GetType() *string {
	return s.Type
}

func (s *GetBackupResponseBodyBackup) SetAutoImageCount(v int64) *GetBackupResponseBodyBackup {
	s.AutoImageCount = &v
	return s
}

func (s *GetBackupResponseBodyBackup) SetBackupHourInDay(v string) *GetBackupResponseBodyBackup {
	s.BackupHourInDay = &v
	return s
}

func (s *GetBackupResponseBodyBackup) SetBackupId(v string) *GetBackupResponseBodyBackup {
	s.BackupId = &v
	return s
}

func (s *GetBackupResponseBodyBackup) SetBackupPeriod(v int64) *GetBackupResponseBodyBackup {
	s.BackupPeriod = &v
	return s
}

func (s *GetBackupResponseBodyBackup) SetCreateTime(v int64) *GetBackupResponseBodyBackup {
	s.CreateTime = &v
	return s
}

func (s *GetBackupResponseBodyBackup) SetExpireTime(v int64) *GetBackupResponseBodyBackup {
	s.ExpireTime = &v
	return s
}

func (s *GetBackupResponseBodyBackup) SetInstanceId(v string) *GetBackupResponseBodyBackup {
	s.InstanceId = &v
	return s
}

func (s *GetBackupResponseBodyBackup) SetMaxImageCount(v string) *GetBackupResponseBodyBackup {
	s.MaxImageCount = &v
	return s
}

func (s *GetBackupResponseBodyBackup) SetName(v string) *GetBackupResponseBodyBackup {
	s.Name = &v
	return s
}

func (s *GetBackupResponseBodyBackup) SetNextImageCreateTime(v int64) *GetBackupResponseBodyBackup {
	s.NextImageCreateTime = &v
	return s
}

func (s *GetBackupResponseBodyBackup) SetOwnerInstanceId(v string) *GetBackupResponseBodyBackup {
	s.OwnerInstanceId = &v
	return s
}

func (s *GetBackupResponseBodyBackup) SetRegionId(v string) *GetBackupResponseBodyBackup {
	s.RegionId = &v
	return s
}

func (s *GetBackupResponseBodyBackup) SetReleaseTime(v int64) *GetBackupResponseBodyBackup {
	s.ReleaseTime = &v
	return s
}

func (s *GetBackupResponseBodyBackup) SetRemark(v string) *GetBackupResponseBodyBackup {
	s.Remark = &v
	return s
}

func (s *GetBackupResponseBodyBackup) SetSpInstanceId(v string) *GetBackupResponseBodyBackup {
	s.SpInstanceId = &v
	return s
}

func (s *GetBackupResponseBodyBackup) SetStatus(v string) *GetBackupResponseBodyBackup {
	s.Status = &v
	return s
}

func (s *GetBackupResponseBodyBackup) SetType(v string) *GetBackupResponseBodyBackup {
	s.Type = &v
	return s
}

func (s *GetBackupResponseBodyBackup) Validate() error {
	return dara.Validate(s)
}
