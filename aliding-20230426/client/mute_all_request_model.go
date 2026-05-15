// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMuteAllRequest interface {
	dara.Model
	String() string
	GoString() string
	SetForceMute(v bool) *MuteAllRequest
	GetForceMute() *bool
	SetTenantContext(v *MuteAllRequestTenantContext) *MuteAllRequest
	GetTenantContext() *MuteAllRequestTenantContext
	SetConferenceId(v string) *MuteAllRequest
	GetConferenceId() *string
	SetMuteAction(v string) *MuteAllRequest
	GetMuteAction() *string
}

type MuteAllRequest struct {
	// example:
	//
	// false
	ForceMute     *bool                        `json:"ForceMute,omitempty" xml:"ForceMute,omitempty"`
	TenantContext *MuteAllRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
	// This parameter is required.
	ConferenceId *string `json:"conferenceId,omitempty" xml:"conferenceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// mute
	MuteAction *string `json:"muteAction,omitempty" xml:"muteAction,omitempty"`
}

func (s MuteAllRequest) String() string {
	return dara.Prettify(s)
}

func (s MuteAllRequest) GoString() string {
	return s.String()
}

func (s *MuteAllRequest) GetForceMute() *bool {
	return s.ForceMute
}

func (s *MuteAllRequest) GetTenantContext() *MuteAllRequestTenantContext {
	return s.TenantContext
}

func (s *MuteAllRequest) GetConferenceId() *string {
	return s.ConferenceId
}

func (s *MuteAllRequest) GetMuteAction() *string {
	return s.MuteAction
}

func (s *MuteAllRequest) SetForceMute(v bool) *MuteAllRequest {
	s.ForceMute = &v
	return s
}

func (s *MuteAllRequest) SetTenantContext(v *MuteAllRequestTenantContext) *MuteAllRequest {
	s.TenantContext = v
	return s
}

func (s *MuteAllRequest) SetConferenceId(v string) *MuteAllRequest {
	s.ConferenceId = &v
	return s
}

func (s *MuteAllRequest) SetMuteAction(v string) *MuteAllRequest {
	s.MuteAction = &v
	return s
}

func (s *MuteAllRequest) Validate() error {
	if s.TenantContext != nil {
		if err := s.TenantContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type MuteAllRequestTenantContext struct {
	// example:
	//
	// xxxxxx
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s MuteAllRequestTenantContext) String() string {
	return dara.Prettify(s)
}

func (s MuteAllRequestTenantContext) GoString() string {
	return s.String()
}

func (s *MuteAllRequestTenantContext) GetTenantId() *string {
	return s.TenantId
}

func (s *MuteAllRequestTenantContext) SetTenantId(v string) *MuteAllRequestTenantContext {
	s.TenantId = &v
	return s
}

func (s *MuteAllRequestTenantContext) Validate() error {
	return dara.Validate(s)
}
