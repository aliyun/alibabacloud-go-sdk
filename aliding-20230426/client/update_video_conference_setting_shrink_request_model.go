// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateVideoConferenceSettingShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAllowUnmuteSelf(v bool) *UpdateVideoConferenceSettingShrinkRequest
	GetAllowUnmuteSelf() *bool
	SetAutoTransferHost(v bool) *UpdateVideoConferenceSettingShrinkRequest
	GetAutoTransferHost() *bool
	SetForbiddenShareScreen(v bool) *UpdateVideoConferenceSettingShrinkRequest
	GetForbiddenShareScreen() *bool
	SetLockConference(v bool) *UpdateVideoConferenceSettingShrinkRequest
	GetLockConference() *bool
	SetMuteAll(v bool) *UpdateVideoConferenceSettingShrinkRequest
	GetMuteAll() *bool
	SetOnlyInternalEmployeesJoin(v bool) *UpdateVideoConferenceSettingShrinkRequest
	GetOnlyInternalEmployeesJoin() *bool
	SetTenantContextShrink(v string) *UpdateVideoConferenceSettingShrinkRequest
	GetTenantContextShrink() *string
	SetConferenceId(v string) *UpdateVideoConferenceSettingShrinkRequest
	GetConferenceId() *string
}

type UpdateVideoConferenceSettingShrinkRequest struct {
	// example:
	//
	// true
	AllowUnmuteSelf *bool `json:"AllowUnmuteSelf,omitempty" xml:"AllowUnmuteSelf,omitempty"`
	// example:
	//
	// true
	AutoTransferHost *bool `json:"AutoTransferHost,omitempty" xml:"AutoTransferHost,omitempty"`
	// example:
	//
	// true
	ForbiddenShareScreen *bool `json:"ForbiddenShareScreen,omitempty" xml:"ForbiddenShareScreen,omitempty"`
	// example:
	//
	// true
	LockConference *bool `json:"LockConference,omitempty" xml:"LockConference,omitempty"`
	// example:
	//
	// true
	MuteAll *bool `json:"MuteAll,omitempty" xml:"MuteAll,omitempty"`
	// example:
	//
	// true
	OnlyInternalEmployeesJoin *bool   `json:"OnlyInternalEmployeesJoin,omitempty" xml:"OnlyInternalEmployeesJoin,omitempty"`
	TenantContextShrink       *string `json:"TenantContext,omitempty" xml:"TenantContext,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 61289fxxx
	ConferenceId *string `json:"conferenceId,omitempty" xml:"conferenceId,omitempty"`
}

func (s UpdateVideoConferenceSettingShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateVideoConferenceSettingShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateVideoConferenceSettingShrinkRequest) GetAllowUnmuteSelf() *bool {
	return s.AllowUnmuteSelf
}

func (s *UpdateVideoConferenceSettingShrinkRequest) GetAutoTransferHost() *bool {
	return s.AutoTransferHost
}

func (s *UpdateVideoConferenceSettingShrinkRequest) GetForbiddenShareScreen() *bool {
	return s.ForbiddenShareScreen
}

func (s *UpdateVideoConferenceSettingShrinkRequest) GetLockConference() *bool {
	return s.LockConference
}

func (s *UpdateVideoConferenceSettingShrinkRequest) GetMuteAll() *bool {
	return s.MuteAll
}

func (s *UpdateVideoConferenceSettingShrinkRequest) GetOnlyInternalEmployeesJoin() *bool {
	return s.OnlyInternalEmployeesJoin
}

func (s *UpdateVideoConferenceSettingShrinkRequest) GetTenantContextShrink() *string {
	return s.TenantContextShrink
}

func (s *UpdateVideoConferenceSettingShrinkRequest) GetConferenceId() *string {
	return s.ConferenceId
}

func (s *UpdateVideoConferenceSettingShrinkRequest) SetAllowUnmuteSelf(v bool) *UpdateVideoConferenceSettingShrinkRequest {
	s.AllowUnmuteSelf = &v
	return s
}

func (s *UpdateVideoConferenceSettingShrinkRequest) SetAutoTransferHost(v bool) *UpdateVideoConferenceSettingShrinkRequest {
	s.AutoTransferHost = &v
	return s
}

func (s *UpdateVideoConferenceSettingShrinkRequest) SetForbiddenShareScreen(v bool) *UpdateVideoConferenceSettingShrinkRequest {
	s.ForbiddenShareScreen = &v
	return s
}

func (s *UpdateVideoConferenceSettingShrinkRequest) SetLockConference(v bool) *UpdateVideoConferenceSettingShrinkRequest {
	s.LockConference = &v
	return s
}

func (s *UpdateVideoConferenceSettingShrinkRequest) SetMuteAll(v bool) *UpdateVideoConferenceSettingShrinkRequest {
	s.MuteAll = &v
	return s
}

func (s *UpdateVideoConferenceSettingShrinkRequest) SetOnlyInternalEmployeesJoin(v bool) *UpdateVideoConferenceSettingShrinkRequest {
	s.OnlyInternalEmployeesJoin = &v
	return s
}

func (s *UpdateVideoConferenceSettingShrinkRequest) SetTenantContextShrink(v string) *UpdateVideoConferenceSettingShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *UpdateVideoConferenceSettingShrinkRequest) SetConferenceId(v string) *UpdateVideoConferenceSettingShrinkRequest {
	s.ConferenceId = &v
	return s
}

func (s *UpdateVideoConferenceSettingShrinkRequest) Validate() error {
	return dara.Validate(s)
}
