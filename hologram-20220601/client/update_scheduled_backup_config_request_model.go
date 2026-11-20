// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateScheduledBackupConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *UpdateScheduledBackupConfigRequest
	GetRegionId() *string
	SetDataKeepQuantity(v int32) *UpdateScheduledBackupConfigRequest
	GetDataKeepQuantity() *int32
	SetDstRegion(v string) *UpdateScheduledBackupConfigRequest
	GetDstRegion() *string
	SetEnabled(v bool) *UpdateScheduledBackupConfigRequest
	GetEnabled() *bool
	SetHour(v int32) *UpdateScheduledBackupConfigRequest
	GetHour() *int32
	SetInstanceId(v string) *UpdateScheduledBackupConfigRequest
	GetInstanceId() *string
	SetManualDataKeepQuantity(v int32) *UpdateScheduledBackupConfigRequest
	GetManualDataKeepQuantity() *int32
	SetScheduleType(v string) *UpdateScheduledBackupConfigRequest
	GetScheduleType() *string
	SetWeek(v string) *UpdateScheduledBackupConfigRequest
	GetWeek() *string
	SetZoneId(v string) *UpdateScheduledBackupConfigRequest
	GetZoneId() *string
}

type UpdateScheduledBackupConfigRequest struct {
	// The region ID.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The number of periodic backups to retain. For example, if you set this parameter to 3, the system automatically retains the latest three backups and deletes other expired backups.
	//
	// example:
	//
	// 3
	DataKeepQuantity *int32 `json:"dataKeepQuantity,omitempty" xml:"dataKeepQuantity,omitempty"`
	// The destination region for backups. This parameter takes effect only for remote backups. You can back up data from the local region to another region to implement cross-region disaster recovery. For example, you can back up an instance in Hangzhou to Beijing.
	//
	// example:
	//
	// cn-beijing
	DstRegion *string `json:"dstRegion,omitempty" xml:"dstRegion,omitempty"`
	// Specifies whether to enable periodic backup. Valid values:
	//
	// - true: Periodic backup is enabled.
	//
	// - false: Periodic backup is disabled.
	//
	// example:
	//
	// true
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// The start time for backup dates. The backup task starts at this time. The time applies to all valid dates. For example, if week is set to 1,3,5 and hour is set to 20, backups start at 20:00:00 every Monday, Wednesday, and Friday.
	//
	// example:
	//
	// 20
	Hour *int32 `json:"hour,omitempty" xml:"hour,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// hgprecn-cn-zvp25ysxxxx
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// The number of manual backups to retain. Manual backups and automatic backups are stored separately, and each has a completely independent data retention policy.
	//
	// example:
	//
	// 3
	ManualDataKeepQuantity *int32 `json:"manualDataKeepQuantity,omitempty" xml:"manualDataKeepQuantity,omitempty"`
	// The backup type. You need to configure this field only for remote backups. You can ignore this field for other backup types.
	//
	// example:
	//
	// remote
	ScheduleType *string `json:"scheduleType,omitempty" xml:"scheduleType,omitempty"`
	// The weekly configuration for backup dates. You can specify multiple days separated by commas. For example, to configure Monday, Wednesday, and Friday, set the value to "1,3,5".
	//
	// example:
	//
	// 1,3,5
	Week *string `json:"week,omitempty" xml:"week,omitempty"`
	// The destination zone for backups. This parameter takes effect only for remote backups.
	//
	// example:
	//
	// cn-beijing-i
	ZoneId *string `json:"zoneId,omitempty" xml:"zoneId,omitempty"`
}

func (s UpdateScheduledBackupConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateScheduledBackupConfigRequest) GoString() string {
	return s.String()
}

func (s *UpdateScheduledBackupConfigRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateScheduledBackupConfigRequest) GetDataKeepQuantity() *int32 {
	return s.DataKeepQuantity
}

func (s *UpdateScheduledBackupConfigRequest) GetDstRegion() *string {
	return s.DstRegion
}

func (s *UpdateScheduledBackupConfigRequest) GetEnabled() *bool {
	return s.Enabled
}

func (s *UpdateScheduledBackupConfigRequest) GetHour() *int32 {
	return s.Hour
}

func (s *UpdateScheduledBackupConfigRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateScheduledBackupConfigRequest) GetManualDataKeepQuantity() *int32 {
	return s.ManualDataKeepQuantity
}

func (s *UpdateScheduledBackupConfigRequest) GetScheduleType() *string {
	return s.ScheduleType
}

func (s *UpdateScheduledBackupConfigRequest) GetWeek() *string {
	return s.Week
}

func (s *UpdateScheduledBackupConfigRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *UpdateScheduledBackupConfigRequest) SetRegionId(v string) *UpdateScheduledBackupConfigRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateScheduledBackupConfigRequest) SetDataKeepQuantity(v int32) *UpdateScheduledBackupConfigRequest {
	s.DataKeepQuantity = &v
	return s
}

func (s *UpdateScheduledBackupConfigRequest) SetDstRegion(v string) *UpdateScheduledBackupConfigRequest {
	s.DstRegion = &v
	return s
}

func (s *UpdateScheduledBackupConfigRequest) SetEnabled(v bool) *UpdateScheduledBackupConfigRequest {
	s.Enabled = &v
	return s
}

func (s *UpdateScheduledBackupConfigRequest) SetHour(v int32) *UpdateScheduledBackupConfigRequest {
	s.Hour = &v
	return s
}

func (s *UpdateScheduledBackupConfigRequest) SetInstanceId(v string) *UpdateScheduledBackupConfigRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateScheduledBackupConfigRequest) SetManualDataKeepQuantity(v int32) *UpdateScheduledBackupConfigRequest {
	s.ManualDataKeepQuantity = &v
	return s
}

func (s *UpdateScheduledBackupConfigRequest) SetScheduleType(v string) *UpdateScheduledBackupConfigRequest {
	s.ScheduleType = &v
	return s
}

func (s *UpdateScheduledBackupConfigRequest) SetWeek(v string) *UpdateScheduledBackupConfigRequest {
	s.Week = &v
	return s
}

func (s *UpdateScheduledBackupConfigRequest) SetZoneId(v string) *UpdateScheduledBackupConfigRequest {
	s.ZoneId = &v
	return s
}

func (s *UpdateScheduledBackupConfigRequest) Validate() error {
	return dara.Validate(s)
}
