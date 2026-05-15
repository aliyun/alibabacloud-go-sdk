// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMuteMembersShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTenantContextShrink(v string) *MuteMembersShrinkRequest
	GetTenantContextShrink() *string
	SetUserIdsShrink(v string) *MuteMembersShrinkRequest
	GetUserIdsShrink() *string
	SetConferenceId(v string) *MuteMembersShrinkRequest
	GetConferenceId() *string
	SetMuteAction(v string) *MuteMembersShrinkRequest
	GetMuteAction() *string
}

type MuteMembersShrinkRequest struct {
	TenantContextShrink *string `json:"TenantContext,omitempty" xml:"TenantContext,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ["012345"]
	UserIdsShrink *string `json:"UserIds,omitempty" xml:"UserIds,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 607452e01401526ee39609e1
	ConferenceId *string `json:"conferenceId,omitempty" xml:"conferenceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// mute
	MuteAction *string `json:"muteAction,omitempty" xml:"muteAction,omitempty"`
}

func (s MuteMembersShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s MuteMembersShrinkRequest) GoString() string {
	return s.String()
}

func (s *MuteMembersShrinkRequest) GetTenantContextShrink() *string {
	return s.TenantContextShrink
}

func (s *MuteMembersShrinkRequest) GetUserIdsShrink() *string {
	return s.UserIdsShrink
}

func (s *MuteMembersShrinkRequest) GetConferenceId() *string {
	return s.ConferenceId
}

func (s *MuteMembersShrinkRequest) GetMuteAction() *string {
	return s.MuteAction
}

func (s *MuteMembersShrinkRequest) SetTenantContextShrink(v string) *MuteMembersShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *MuteMembersShrinkRequest) SetUserIdsShrink(v string) *MuteMembersShrinkRequest {
	s.UserIdsShrink = &v
	return s
}

func (s *MuteMembersShrinkRequest) SetConferenceId(v string) *MuteMembersShrinkRequest {
	s.ConferenceId = &v
	return s
}

func (s *MuteMembersShrinkRequest) SetMuteAction(v string) *MuteMembersShrinkRequest {
	s.MuteAction = &v
	return s
}

func (s *MuteMembersShrinkRequest) Validate() error {
	return dara.Validate(s)
}
