// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDesktopOversoldGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConcurrenceCount(v int32) *ModifyDesktopOversoldGroupRequest
	GetConcurrenceCount() *int32
	SetDescription(v string) *ModifyDesktopOversoldGroupRequest
	GetDescription() *string
	SetIdleDisconnectDuration(v int64) *ModifyDesktopOversoldGroupRequest
	GetIdleDisconnectDuration() *int64
	SetImageId(v string) *ModifyDesktopOversoldGroupRequest
	GetImageId() *string
	SetKeepDuration(v int32) *ModifyDesktopOversoldGroupRequest
	GetKeepDuration() *int32
	SetName(v string) *ModifyDesktopOversoldGroupRequest
	GetName() *string
	SetOversoldGroupId(v string) *ModifyDesktopOversoldGroupRequest
	GetOversoldGroupId() *string
	SetOversoldUserCount(v int32) *ModifyDesktopOversoldGroupRequest
	GetOversoldUserCount() *int32
	SetOversoldWarn(v int32) *ModifyDesktopOversoldGroupRequest
	GetOversoldWarn() *int32
	SetPolicyGroupId(v string) *ModifyDesktopOversoldGroupRequest
	GetPolicyGroupId() *string
	SetStopDuration(v int32) *ModifyDesktopOversoldGroupRequest
	GetStopDuration() *int32
}

type ModifyDesktopOversoldGroupRequest struct {
	ConcurrenceCount       *int32  `json:"ConcurrenceCount,omitempty" xml:"ConcurrenceCount,omitempty"`
	Description            *string `json:"Description,omitempty" xml:"Description,omitempty"`
	IdleDisconnectDuration *int64  `json:"IdleDisconnectDuration,omitempty" xml:"IdleDisconnectDuration,omitempty"`
	ImageId                *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	KeepDuration           *int32  `json:"KeepDuration,omitempty" xml:"KeepDuration,omitempty"`
	Name                   *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OversoldGroupId        *string `json:"OversoldGroupId,omitempty" xml:"OversoldGroupId,omitempty"`
	OversoldUserCount      *int32  `json:"OversoldUserCount,omitempty" xml:"OversoldUserCount,omitempty"`
	OversoldWarn           *int32  `json:"OversoldWarn,omitempty" xml:"OversoldWarn,omitempty"`
	PolicyGroupId          *string `json:"PolicyGroupId,omitempty" xml:"PolicyGroupId,omitempty"`
	StopDuration           *int32  `json:"StopDuration,omitempty" xml:"StopDuration,omitempty"`
}

func (s ModifyDesktopOversoldGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDesktopOversoldGroupRequest) GoString() string {
	return s.String()
}

func (s *ModifyDesktopOversoldGroupRequest) GetConcurrenceCount() *int32 {
	return s.ConcurrenceCount
}

func (s *ModifyDesktopOversoldGroupRequest) GetDescription() *string {
	return s.Description
}

func (s *ModifyDesktopOversoldGroupRequest) GetIdleDisconnectDuration() *int64 {
	return s.IdleDisconnectDuration
}

func (s *ModifyDesktopOversoldGroupRequest) GetImageId() *string {
	return s.ImageId
}

func (s *ModifyDesktopOversoldGroupRequest) GetKeepDuration() *int32 {
	return s.KeepDuration
}

func (s *ModifyDesktopOversoldGroupRequest) GetName() *string {
	return s.Name
}

func (s *ModifyDesktopOversoldGroupRequest) GetOversoldGroupId() *string {
	return s.OversoldGroupId
}

func (s *ModifyDesktopOversoldGroupRequest) GetOversoldUserCount() *int32 {
	return s.OversoldUserCount
}

func (s *ModifyDesktopOversoldGroupRequest) GetOversoldWarn() *int32 {
	return s.OversoldWarn
}

func (s *ModifyDesktopOversoldGroupRequest) GetPolicyGroupId() *string {
	return s.PolicyGroupId
}

func (s *ModifyDesktopOversoldGroupRequest) GetStopDuration() *int32 {
	return s.StopDuration
}

func (s *ModifyDesktopOversoldGroupRequest) SetConcurrenceCount(v int32) *ModifyDesktopOversoldGroupRequest {
	s.ConcurrenceCount = &v
	return s
}

func (s *ModifyDesktopOversoldGroupRequest) SetDescription(v string) *ModifyDesktopOversoldGroupRequest {
	s.Description = &v
	return s
}

func (s *ModifyDesktopOversoldGroupRequest) SetIdleDisconnectDuration(v int64) *ModifyDesktopOversoldGroupRequest {
	s.IdleDisconnectDuration = &v
	return s
}

func (s *ModifyDesktopOversoldGroupRequest) SetImageId(v string) *ModifyDesktopOversoldGroupRequest {
	s.ImageId = &v
	return s
}

func (s *ModifyDesktopOversoldGroupRequest) SetKeepDuration(v int32) *ModifyDesktopOversoldGroupRequest {
	s.KeepDuration = &v
	return s
}

func (s *ModifyDesktopOversoldGroupRequest) SetName(v string) *ModifyDesktopOversoldGroupRequest {
	s.Name = &v
	return s
}

func (s *ModifyDesktopOversoldGroupRequest) SetOversoldGroupId(v string) *ModifyDesktopOversoldGroupRequest {
	s.OversoldGroupId = &v
	return s
}

func (s *ModifyDesktopOversoldGroupRequest) SetOversoldUserCount(v int32) *ModifyDesktopOversoldGroupRequest {
	s.OversoldUserCount = &v
	return s
}

func (s *ModifyDesktopOversoldGroupRequest) SetOversoldWarn(v int32) *ModifyDesktopOversoldGroupRequest {
	s.OversoldWarn = &v
	return s
}

func (s *ModifyDesktopOversoldGroupRequest) SetPolicyGroupId(v string) *ModifyDesktopOversoldGroupRequest {
	s.PolicyGroupId = &v
	return s
}

func (s *ModifyDesktopOversoldGroupRequest) SetStopDuration(v int32) *ModifyDesktopOversoldGroupRequest {
	s.StopDuration = &v
	return s
}

func (s *ModifyDesktopOversoldGroupRequest) Validate() error {
	return dara.Validate(s)
}
