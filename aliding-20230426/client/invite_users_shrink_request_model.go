// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInviteUsersShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInviteeListShrink(v string) *InviteUsersShrinkRequest
	GetInviteeListShrink() *string
	SetTenantContextShrink(v string) *InviteUsersShrinkRequest
	GetTenantContextShrink() *string
	SetConferenceId(v string) *InviteUsersShrinkRequest
	GetConferenceId() *string
	SetPhoneInviteeListShrink(v string) *InviteUsersShrinkRequest
	GetPhoneInviteeListShrink() *string
}

type InviteUsersShrinkRequest struct {
	InviteeListShrink   *string `json:"InviteeList,omitempty" xml:"InviteeList,omitempty"`
	TenantContextShrink *string `json:"TenantContext,omitempty" xml:"TenantContext,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 61289fxxx
	ConferenceId           *string `json:"conferenceId,omitempty" xml:"conferenceId,omitempty"`
	PhoneInviteeListShrink *string `json:"phoneInviteeList,omitempty" xml:"phoneInviteeList,omitempty"`
}

func (s InviteUsersShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s InviteUsersShrinkRequest) GoString() string {
	return s.String()
}

func (s *InviteUsersShrinkRequest) GetInviteeListShrink() *string {
	return s.InviteeListShrink
}

func (s *InviteUsersShrinkRequest) GetTenantContextShrink() *string {
	return s.TenantContextShrink
}

func (s *InviteUsersShrinkRequest) GetConferenceId() *string {
	return s.ConferenceId
}

func (s *InviteUsersShrinkRequest) GetPhoneInviteeListShrink() *string {
	return s.PhoneInviteeListShrink
}

func (s *InviteUsersShrinkRequest) SetInviteeListShrink(v string) *InviteUsersShrinkRequest {
	s.InviteeListShrink = &v
	return s
}

func (s *InviteUsersShrinkRequest) SetTenantContextShrink(v string) *InviteUsersShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *InviteUsersShrinkRequest) SetConferenceId(v string) *InviteUsersShrinkRequest {
	s.ConferenceId = &v
	return s
}

func (s *InviteUsersShrinkRequest) SetPhoneInviteeListShrink(v string) *InviteUsersShrinkRequest {
	s.PhoneInviteeListShrink = &v
	return s
}

func (s *InviteUsersShrinkRequest) Validate() error {
	return dara.Validate(s)
}
