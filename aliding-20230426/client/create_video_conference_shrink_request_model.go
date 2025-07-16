// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVideoConferenceShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfTitle(v string) *CreateVideoConferenceShrinkRequest
	GetConfTitle() *string
	SetInviteCaller(v bool) *CreateVideoConferenceShrinkRequest
	GetInviteCaller() *bool
	SetInviteUserIdsShrink(v string) *CreateVideoConferenceShrinkRequest
	GetInviteUserIdsShrink() *string
}

type CreateVideoConferenceShrinkRequest struct {
	// This parameter is required.
	ConfTitle *string `json:"ConfTitle,omitempty" xml:"ConfTitle,omitempty"`
	// example:
	//
	// true
	InviteCaller        *bool   `json:"InviteCaller,omitempty" xml:"InviteCaller,omitempty"`
	InviteUserIdsShrink *string `json:"InviteUserIds,omitempty" xml:"InviteUserIds,omitempty"`
}

func (s CreateVideoConferenceShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateVideoConferenceShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateVideoConferenceShrinkRequest) GetConfTitle() *string {
	return s.ConfTitle
}

func (s *CreateVideoConferenceShrinkRequest) GetInviteCaller() *bool {
	return s.InviteCaller
}

func (s *CreateVideoConferenceShrinkRequest) GetInviteUserIdsShrink() *string {
	return s.InviteUserIdsShrink
}

func (s *CreateVideoConferenceShrinkRequest) SetConfTitle(v string) *CreateVideoConferenceShrinkRequest {
	s.ConfTitle = &v
	return s
}

func (s *CreateVideoConferenceShrinkRequest) SetInviteCaller(v bool) *CreateVideoConferenceShrinkRequest {
	s.InviteCaller = &v
	return s
}

func (s *CreateVideoConferenceShrinkRequest) SetInviteUserIdsShrink(v string) *CreateVideoConferenceShrinkRequest {
	s.InviteUserIdsShrink = &v
	return s
}

func (s *CreateVideoConferenceShrinkRequest) Validate() error {
	return dara.Validate(s)
}
