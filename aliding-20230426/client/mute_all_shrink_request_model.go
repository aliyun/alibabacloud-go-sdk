// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMuteAllShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetForceMute(v bool) *MuteAllShrinkRequest
	GetForceMute() *bool
	SetTenantContextShrink(v string) *MuteAllShrinkRequest
	GetTenantContextShrink() *string
	SetConferenceId(v string) *MuteAllShrinkRequest
	GetConferenceId() *string
	SetMuteAction(v string) *MuteAllShrinkRequest
	GetMuteAction() *string
}

type MuteAllShrinkRequest struct {
	// example:
	//
	// false
	ForceMute           *bool   `json:"ForceMute,omitempty" xml:"ForceMute,omitempty"`
	TenantContextShrink *string `json:"TenantContext,omitempty" xml:"TenantContext,omitempty"`
	// This parameter is required.
	ConferenceId *string `json:"conferenceId,omitempty" xml:"conferenceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// mute
	MuteAction *string `json:"muteAction,omitempty" xml:"muteAction,omitempty"`
}

func (s MuteAllShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s MuteAllShrinkRequest) GoString() string {
	return s.String()
}

func (s *MuteAllShrinkRequest) GetForceMute() *bool {
	return s.ForceMute
}

func (s *MuteAllShrinkRequest) GetTenantContextShrink() *string {
	return s.TenantContextShrink
}

func (s *MuteAllShrinkRequest) GetConferenceId() *string {
	return s.ConferenceId
}

func (s *MuteAllShrinkRequest) GetMuteAction() *string {
	return s.MuteAction
}

func (s *MuteAllShrinkRequest) SetForceMute(v bool) *MuteAllShrinkRequest {
	s.ForceMute = &v
	return s
}

func (s *MuteAllShrinkRequest) SetTenantContextShrink(v string) *MuteAllShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *MuteAllShrinkRequest) SetConferenceId(v string) *MuteAllShrinkRequest {
	s.ConferenceId = &v
	return s
}

func (s *MuteAllShrinkRequest) SetMuteAction(v string) *MuteAllShrinkRequest {
	s.MuteAction = &v
	return s
}

func (s *MuteAllShrinkRequest) Validate() error {
	return dara.Validate(s)
}
