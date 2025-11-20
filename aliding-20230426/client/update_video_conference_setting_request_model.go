// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateVideoConferenceSettingRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAllowUnmuteSelf(v bool) *UpdateVideoConferenceSettingRequest
	GetAllowUnmuteSelf() *bool
	SetAutoTransferHost(v bool) *UpdateVideoConferenceSettingRequest
	GetAutoTransferHost() *bool
	SetForbiddenShareScreen(v bool) *UpdateVideoConferenceSettingRequest
	GetForbiddenShareScreen() *bool
	SetLockConference(v bool) *UpdateVideoConferenceSettingRequest
	GetLockConference() *bool
	SetMuteAll(v bool) *UpdateVideoConferenceSettingRequest
	GetMuteAll() *bool
	SetOnlyInternalEmployeesJoin(v bool) *UpdateVideoConferenceSettingRequest
	GetOnlyInternalEmployeesJoin() *bool
	SetTenantContext(v *UpdateVideoConferenceSettingRequestTenantContext) *UpdateVideoConferenceSettingRequest
	GetTenantContext() *UpdateVideoConferenceSettingRequestTenantContext
	SetConferenceId(v string) *UpdateVideoConferenceSettingRequest
	GetConferenceId() *string
}

type UpdateVideoConferenceSettingRequest struct {
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
	OnlyInternalEmployeesJoin *bool                                             `json:"OnlyInternalEmployeesJoin,omitempty" xml:"OnlyInternalEmployeesJoin,omitempty"`
	TenantContext             *UpdateVideoConferenceSettingRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 61289fxxx
	ConferenceId *string `json:"conferenceId,omitempty" xml:"conferenceId,omitempty"`
}

func (s UpdateVideoConferenceSettingRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateVideoConferenceSettingRequest) GoString() string {
	return s.String()
}

func (s *UpdateVideoConferenceSettingRequest) GetAllowUnmuteSelf() *bool {
	return s.AllowUnmuteSelf
}

func (s *UpdateVideoConferenceSettingRequest) GetAutoTransferHost() *bool {
	return s.AutoTransferHost
}

func (s *UpdateVideoConferenceSettingRequest) GetForbiddenShareScreen() *bool {
	return s.ForbiddenShareScreen
}

func (s *UpdateVideoConferenceSettingRequest) GetLockConference() *bool {
	return s.LockConference
}

func (s *UpdateVideoConferenceSettingRequest) GetMuteAll() *bool {
	return s.MuteAll
}

func (s *UpdateVideoConferenceSettingRequest) GetOnlyInternalEmployeesJoin() *bool {
	return s.OnlyInternalEmployeesJoin
}

func (s *UpdateVideoConferenceSettingRequest) GetTenantContext() *UpdateVideoConferenceSettingRequestTenantContext {
	return s.TenantContext
}

func (s *UpdateVideoConferenceSettingRequest) GetConferenceId() *string {
	return s.ConferenceId
}

func (s *UpdateVideoConferenceSettingRequest) SetAllowUnmuteSelf(v bool) *UpdateVideoConferenceSettingRequest {
	s.AllowUnmuteSelf = &v
	return s
}

func (s *UpdateVideoConferenceSettingRequest) SetAutoTransferHost(v bool) *UpdateVideoConferenceSettingRequest {
	s.AutoTransferHost = &v
	return s
}

func (s *UpdateVideoConferenceSettingRequest) SetForbiddenShareScreen(v bool) *UpdateVideoConferenceSettingRequest {
	s.ForbiddenShareScreen = &v
	return s
}

func (s *UpdateVideoConferenceSettingRequest) SetLockConference(v bool) *UpdateVideoConferenceSettingRequest {
	s.LockConference = &v
	return s
}

func (s *UpdateVideoConferenceSettingRequest) SetMuteAll(v bool) *UpdateVideoConferenceSettingRequest {
	s.MuteAll = &v
	return s
}

func (s *UpdateVideoConferenceSettingRequest) SetOnlyInternalEmployeesJoin(v bool) *UpdateVideoConferenceSettingRequest {
	s.OnlyInternalEmployeesJoin = &v
	return s
}

func (s *UpdateVideoConferenceSettingRequest) SetTenantContext(v *UpdateVideoConferenceSettingRequestTenantContext) *UpdateVideoConferenceSettingRequest {
	s.TenantContext = v
	return s
}

func (s *UpdateVideoConferenceSettingRequest) SetConferenceId(v string) *UpdateVideoConferenceSettingRequest {
	s.ConferenceId = &v
	return s
}

func (s *UpdateVideoConferenceSettingRequest) Validate() error {
	if s.TenantContext != nil {
		if err := s.TenantContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateVideoConferenceSettingRequestTenantContext struct {
	// example:
	//
	// xxxxxx
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s UpdateVideoConferenceSettingRequestTenantContext) String() string {
	return dara.Prettify(s)
}

func (s UpdateVideoConferenceSettingRequestTenantContext) GoString() string {
	return s.String()
}

func (s *UpdateVideoConferenceSettingRequestTenantContext) GetTenantId() *string {
	return s.TenantId
}

func (s *UpdateVideoConferenceSettingRequestTenantContext) SetTenantId(v string) *UpdateVideoConferenceSettingRequestTenantContext {
	s.TenantId = &v
	return s
}

func (s *UpdateVideoConferenceSettingRequestTenantContext) Validate() error {
	return dara.Validate(s)
}
