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
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// 3
	DataKeepQuantity *int32 `json:"dataKeepQuantity,omitempty" xml:"dataKeepQuantity,omitempty"`
	// example:
	//
	// cn-beijing
	DstRegion *string `json:"dstRegion,omitempty" xml:"dstRegion,omitempty"`
	// example:
	//
	// true
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// example:
	//
	// 20
	Hour *int32 `json:"hour,omitempty" xml:"hour,omitempty"`
	// example:
	//
	// hgprecn-cn-zvp25ysxxxx
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// example:
	//
	// 3
	ManualDataKeepQuantity *int32 `json:"manualDataKeepQuantity,omitempty" xml:"manualDataKeepQuantity,omitempty"`
	// example:
	//
	// remote
	ScheduleType *string `json:"scheduleType,omitempty" xml:"scheduleType,omitempty"`
	// example:
	//
	// 1,3,5
	Week *string `json:"week,omitempty" xml:"week,omitempty"`
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
