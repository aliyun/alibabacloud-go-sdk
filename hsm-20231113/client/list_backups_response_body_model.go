// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBackupsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBackups(v []*ListBackupsResponseBodyBackups) *ListBackupsResponseBody
	GetBackups() []*ListBackupsResponseBodyBackups
	SetCurrentPage(v int32) *ListBackupsResponseBody
	GetCurrentPage() *int32
	SetPageSize(v int32) *ListBackupsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListBackupsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListBackupsResponseBody
	GetTotalCount() *int32
}

type ListBackupsResponseBody struct {
	// The backups returned.
	Backups []*ListBackupsResponseBodyBackups `json:"Backups,omitempty" xml:"Backups,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 4C467B38-3910-447D-87BC-AC049166F216
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 30
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListBackupsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListBackupsResponseBody) GoString() string {
	return s.String()
}

func (s *ListBackupsResponseBody) GetBackups() []*ListBackupsResponseBodyBackups {
	return s.Backups
}

func (s *ListBackupsResponseBody) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListBackupsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListBackupsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListBackupsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListBackupsResponseBody) SetBackups(v []*ListBackupsResponseBodyBackups) *ListBackupsResponseBody {
	s.Backups = v
	return s
}

func (s *ListBackupsResponseBody) SetCurrentPage(v int32) *ListBackupsResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *ListBackupsResponseBody) SetPageSize(v int32) *ListBackupsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListBackupsResponseBody) SetRequestId(v string) *ListBackupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListBackupsResponseBody) SetTotalCount(v int32) *ListBackupsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListBackupsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListBackupsResponseBodyBackups struct {
	// The number of images that are automatically backed up.
	//
	// example:
	//
	// 1
	AutoImageCount *int64 `json:"AutoImageCount,omitempty" xml:"AutoImageCount,omitempty"`
	// The backup time on the hour that is in the 24-hour clock.
	//
	// example:
	//
	// 13
	BackupHourInDay *string `json:"BackupHourInDay,omitempty" xml:"BackupHourInDay,omitempty"`
	// The ID of the backup.
	//
	// example:
	//
	// backup-1648438****
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
	// 1637229596000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The expiration time of the backup. The value is accurate to the millisecond. The value is a UNIX timestamp.
	//
	// example:
	//
	// 1682417553781
	ExpireTime *int64 `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The ID of the HSM that is associated with the backup.
	//
	// example:
	//
	// hsm-cn-vj30bil8****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The number of images.
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
	// The time when the image is created next time. The value is accurate to the millisecond. The value is a UNIX timestamp.
	//
	// example:
	//
	// 1682417553781
	NextImageCreateTime *int64 `json:"NextImageCreateTime,omitempty" xml:"NextImageCreateTime,omitempty"`
	// The ID of the HSM to which the backup belongs. This parameter is available only for HSM backups outside the Chinese mainland and the value of this parameter is consistent with the value of InstanceId.
	//
	// example:
	//
	// hsm-cn-vj30bil8****
	OwnerInstanceId *string `json:"OwnerInstanceId,omitempty" xml:"OwnerInstanceId,omitempty"`
	// The ID of the region.
	//
	// example:
	//
	// ap-southeast-1
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
	// normal backup
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The ID of the backup. This parameter is available only for HSM backups in the Chinese mainland.
	//
	// example:
	//
	// backup-fdb897sdfg534-****
	SpInstanceId *string `json:"SpInstanceId,omitempty" xml:"SpInstanceId,omitempty"`
	// The status of the backup. Valid values:
	//
	// 	- NEW: The backup is disabled.
	//
	// 	- EXPIRED: The backup expired.
	//
	// 	- ENABLED: The backup is enabled.
	//
	// example:
	//
	// ENABLED
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The type of the backup. Valid values:
	//
	// 	- DEFAULT
	//
	// 	- NORMAL
	//
	// example:
	//
	// NORMAL
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListBackupsResponseBodyBackups) String() string {
	return dara.Prettify(s)
}

func (s ListBackupsResponseBodyBackups) GoString() string {
	return s.String()
}

func (s *ListBackupsResponseBodyBackups) GetAutoImageCount() *int64 {
	return s.AutoImageCount
}

func (s *ListBackupsResponseBodyBackups) GetBackupHourInDay() *string {
	return s.BackupHourInDay
}

func (s *ListBackupsResponseBodyBackups) GetBackupId() *string {
	return s.BackupId
}

func (s *ListBackupsResponseBodyBackups) GetBackupPeriod() *int64 {
	return s.BackupPeriod
}

func (s *ListBackupsResponseBodyBackups) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListBackupsResponseBodyBackups) GetExpireTime() *int64 {
	return s.ExpireTime
}

func (s *ListBackupsResponseBodyBackups) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListBackupsResponseBodyBackups) GetMaxImageCount() *string {
	return s.MaxImageCount
}

func (s *ListBackupsResponseBodyBackups) GetName() *string {
	return s.Name
}

func (s *ListBackupsResponseBodyBackups) GetNextImageCreateTime() *int64 {
	return s.NextImageCreateTime
}

func (s *ListBackupsResponseBodyBackups) GetOwnerInstanceId() *string {
	return s.OwnerInstanceId
}

func (s *ListBackupsResponseBodyBackups) GetRegionId() *string {
	return s.RegionId
}

func (s *ListBackupsResponseBodyBackups) GetReleaseTime() *int64 {
	return s.ReleaseTime
}

func (s *ListBackupsResponseBodyBackups) GetRemark() *string {
	return s.Remark
}

func (s *ListBackupsResponseBodyBackups) GetSpInstanceId() *string {
	return s.SpInstanceId
}

func (s *ListBackupsResponseBodyBackups) GetStatus() *string {
	return s.Status
}

func (s *ListBackupsResponseBodyBackups) GetType() *string {
	return s.Type
}

func (s *ListBackupsResponseBodyBackups) SetAutoImageCount(v int64) *ListBackupsResponseBodyBackups {
	s.AutoImageCount = &v
	return s
}

func (s *ListBackupsResponseBodyBackups) SetBackupHourInDay(v string) *ListBackupsResponseBodyBackups {
	s.BackupHourInDay = &v
	return s
}

func (s *ListBackupsResponseBodyBackups) SetBackupId(v string) *ListBackupsResponseBodyBackups {
	s.BackupId = &v
	return s
}

func (s *ListBackupsResponseBodyBackups) SetBackupPeriod(v int64) *ListBackupsResponseBodyBackups {
	s.BackupPeriod = &v
	return s
}

func (s *ListBackupsResponseBodyBackups) SetCreateTime(v int64) *ListBackupsResponseBodyBackups {
	s.CreateTime = &v
	return s
}

func (s *ListBackupsResponseBodyBackups) SetExpireTime(v int64) *ListBackupsResponseBodyBackups {
	s.ExpireTime = &v
	return s
}

func (s *ListBackupsResponseBodyBackups) SetInstanceId(v string) *ListBackupsResponseBodyBackups {
	s.InstanceId = &v
	return s
}

func (s *ListBackupsResponseBodyBackups) SetMaxImageCount(v string) *ListBackupsResponseBodyBackups {
	s.MaxImageCount = &v
	return s
}

func (s *ListBackupsResponseBodyBackups) SetName(v string) *ListBackupsResponseBodyBackups {
	s.Name = &v
	return s
}

func (s *ListBackupsResponseBodyBackups) SetNextImageCreateTime(v int64) *ListBackupsResponseBodyBackups {
	s.NextImageCreateTime = &v
	return s
}

func (s *ListBackupsResponseBodyBackups) SetOwnerInstanceId(v string) *ListBackupsResponseBodyBackups {
	s.OwnerInstanceId = &v
	return s
}

func (s *ListBackupsResponseBodyBackups) SetRegionId(v string) *ListBackupsResponseBodyBackups {
	s.RegionId = &v
	return s
}

func (s *ListBackupsResponseBodyBackups) SetReleaseTime(v int64) *ListBackupsResponseBodyBackups {
	s.ReleaseTime = &v
	return s
}

func (s *ListBackupsResponseBodyBackups) SetRemark(v string) *ListBackupsResponseBodyBackups {
	s.Remark = &v
	return s
}

func (s *ListBackupsResponseBodyBackups) SetSpInstanceId(v string) *ListBackupsResponseBodyBackups {
	s.SpInstanceId = &v
	return s
}

func (s *ListBackupsResponseBodyBackups) SetStatus(v string) *ListBackupsResponseBodyBackups {
	s.Status = &v
	return s
}

func (s *ListBackupsResponseBodyBackups) SetType(v string) *ListBackupsResponseBodyBackups {
	s.Type = &v
	return s
}

func (s *ListBackupsResponseBodyBackups) Validate() error {
	return dara.Validate(s)
}
