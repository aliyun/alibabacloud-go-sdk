// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDesktopOversoldGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConcurrenceCount(v int32) *CreateDesktopOversoldGroupRequest
	GetConcurrenceCount() *int32
	SetDataDiskSize(v int32) *CreateDesktopOversoldGroupRequest
	GetDataDiskSize() *int32
	SetDescription(v string) *CreateDesktopOversoldGroupRequest
	GetDescription() *string
	SetDesktopType(v string) *CreateDesktopOversoldGroupRequest
	GetDesktopType() *string
	SetDirectoryId(v string) *CreateDesktopOversoldGroupRequest
	GetDirectoryId() *string
	SetIdleDisconnectDuration(v int64) *CreateDesktopOversoldGroupRequest
	GetIdleDisconnectDuration() *int64
	SetImageId(v string) *CreateDesktopOversoldGroupRequest
	GetImageId() *string
	SetKeepDuration(v int32) *CreateDesktopOversoldGroupRequest
	GetKeepDuration() *int32
	SetName(v string) *CreateDesktopOversoldGroupRequest
	GetName() *string
	SetOversoldUserCount(v int32) *CreateDesktopOversoldGroupRequest
	GetOversoldUserCount() *int32
	SetOversoldWarn(v int32) *CreateDesktopOversoldGroupRequest
	GetOversoldWarn() *int32
	SetPeriod(v int32) *CreateDesktopOversoldGroupRequest
	GetPeriod() *int32
	SetPeriodUnit(v string) *CreateDesktopOversoldGroupRequest
	GetPeriodUnit() *string
	SetPolicyGroupId(v string) *CreateDesktopOversoldGroupRequest
	GetPolicyGroupId() *string
	SetStopDuration(v int32) *CreateDesktopOversoldGroupRequest
	GetStopDuration() *int32
	SetSystemDiskSize(v int32) *CreateDesktopOversoldGroupRequest
	GetSystemDiskSize() *int32
}

type CreateDesktopOversoldGroupRequest struct {
	ConcurrenceCount       *int32  `json:"ConcurrenceCount,omitempty" xml:"ConcurrenceCount,omitempty"`
	DataDiskSize           *int32  `json:"DataDiskSize,omitempty" xml:"DataDiskSize,omitempty"`
	Description            *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DesktopType            *string `json:"DesktopType,omitempty" xml:"DesktopType,omitempty"`
	DirectoryId            *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	IdleDisconnectDuration *int64  `json:"IdleDisconnectDuration,omitempty" xml:"IdleDisconnectDuration,omitempty"`
	ImageId                *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	KeepDuration           *int32  `json:"KeepDuration,omitempty" xml:"KeepDuration,omitempty"`
	Name                   *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OversoldUserCount      *int32  `json:"OversoldUserCount,omitempty" xml:"OversoldUserCount,omitempty"`
	OversoldWarn           *int32  `json:"OversoldWarn,omitempty" xml:"OversoldWarn,omitempty"`
	Period                 *int32  `json:"Period,omitempty" xml:"Period,omitempty"`
	PeriodUnit             *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	PolicyGroupId          *string `json:"PolicyGroupId,omitempty" xml:"PolicyGroupId,omitempty"`
	StopDuration           *int32  `json:"StopDuration,omitempty" xml:"StopDuration,omitempty"`
	SystemDiskSize         *int32  `json:"SystemDiskSize,omitempty" xml:"SystemDiskSize,omitempty"`
}

func (s CreateDesktopOversoldGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDesktopOversoldGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateDesktopOversoldGroupRequest) GetConcurrenceCount() *int32 {
	return s.ConcurrenceCount
}

func (s *CreateDesktopOversoldGroupRequest) GetDataDiskSize() *int32 {
	return s.DataDiskSize
}

func (s *CreateDesktopOversoldGroupRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateDesktopOversoldGroupRequest) GetDesktopType() *string {
	return s.DesktopType
}

func (s *CreateDesktopOversoldGroupRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *CreateDesktopOversoldGroupRequest) GetIdleDisconnectDuration() *int64 {
	return s.IdleDisconnectDuration
}

func (s *CreateDesktopOversoldGroupRequest) GetImageId() *string {
	return s.ImageId
}

func (s *CreateDesktopOversoldGroupRequest) GetKeepDuration() *int32 {
	return s.KeepDuration
}

func (s *CreateDesktopOversoldGroupRequest) GetName() *string {
	return s.Name
}

func (s *CreateDesktopOversoldGroupRequest) GetOversoldUserCount() *int32 {
	return s.OversoldUserCount
}

func (s *CreateDesktopOversoldGroupRequest) GetOversoldWarn() *int32 {
	return s.OversoldWarn
}

func (s *CreateDesktopOversoldGroupRequest) GetPeriod() *int32 {
	return s.Period
}

func (s *CreateDesktopOversoldGroupRequest) GetPeriodUnit() *string {
	return s.PeriodUnit
}

func (s *CreateDesktopOversoldGroupRequest) GetPolicyGroupId() *string {
	return s.PolicyGroupId
}

func (s *CreateDesktopOversoldGroupRequest) GetStopDuration() *int32 {
	return s.StopDuration
}

func (s *CreateDesktopOversoldGroupRequest) GetSystemDiskSize() *int32 {
	return s.SystemDiskSize
}

func (s *CreateDesktopOversoldGroupRequest) SetConcurrenceCount(v int32) *CreateDesktopOversoldGroupRequest {
	s.ConcurrenceCount = &v
	return s
}

func (s *CreateDesktopOversoldGroupRequest) SetDataDiskSize(v int32) *CreateDesktopOversoldGroupRequest {
	s.DataDiskSize = &v
	return s
}

func (s *CreateDesktopOversoldGroupRequest) SetDescription(v string) *CreateDesktopOversoldGroupRequest {
	s.Description = &v
	return s
}

func (s *CreateDesktopOversoldGroupRequest) SetDesktopType(v string) *CreateDesktopOversoldGroupRequest {
	s.DesktopType = &v
	return s
}

func (s *CreateDesktopOversoldGroupRequest) SetDirectoryId(v string) *CreateDesktopOversoldGroupRequest {
	s.DirectoryId = &v
	return s
}

func (s *CreateDesktopOversoldGroupRequest) SetIdleDisconnectDuration(v int64) *CreateDesktopOversoldGroupRequest {
	s.IdleDisconnectDuration = &v
	return s
}

func (s *CreateDesktopOversoldGroupRequest) SetImageId(v string) *CreateDesktopOversoldGroupRequest {
	s.ImageId = &v
	return s
}

func (s *CreateDesktopOversoldGroupRequest) SetKeepDuration(v int32) *CreateDesktopOversoldGroupRequest {
	s.KeepDuration = &v
	return s
}

func (s *CreateDesktopOversoldGroupRequest) SetName(v string) *CreateDesktopOversoldGroupRequest {
	s.Name = &v
	return s
}

func (s *CreateDesktopOversoldGroupRequest) SetOversoldUserCount(v int32) *CreateDesktopOversoldGroupRequest {
	s.OversoldUserCount = &v
	return s
}

func (s *CreateDesktopOversoldGroupRequest) SetOversoldWarn(v int32) *CreateDesktopOversoldGroupRequest {
	s.OversoldWarn = &v
	return s
}

func (s *CreateDesktopOversoldGroupRequest) SetPeriod(v int32) *CreateDesktopOversoldGroupRequest {
	s.Period = &v
	return s
}

func (s *CreateDesktopOversoldGroupRequest) SetPeriodUnit(v string) *CreateDesktopOversoldGroupRequest {
	s.PeriodUnit = &v
	return s
}

func (s *CreateDesktopOversoldGroupRequest) SetPolicyGroupId(v string) *CreateDesktopOversoldGroupRequest {
	s.PolicyGroupId = &v
	return s
}

func (s *CreateDesktopOversoldGroupRequest) SetStopDuration(v int32) *CreateDesktopOversoldGroupRequest {
	s.StopDuration = &v
	return s
}

func (s *CreateDesktopOversoldGroupRequest) SetSystemDiskSize(v int32) *CreateDesktopOversoldGroupRequest {
	s.SystemDiskSize = &v
	return s
}

func (s *CreateDesktopOversoldGroupRequest) Validate() error {
	return dara.Validate(s)
}
