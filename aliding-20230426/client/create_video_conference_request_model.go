// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVideoConferenceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfTitle(v string) *CreateVideoConferenceRequest
	GetConfTitle() *string
	SetInviteCaller(v bool) *CreateVideoConferenceRequest
	GetInviteCaller() *bool
	SetInviteUserIds(v []*string) *CreateVideoConferenceRequest
	GetInviteUserIds() []*string
}

type CreateVideoConferenceRequest struct {
	// This parameter is required.
	ConfTitle *string `json:"ConfTitle,omitempty" xml:"ConfTitle,omitempty"`
	// example:
	//
	// true
	InviteCaller  *bool     `json:"InviteCaller,omitempty" xml:"InviteCaller,omitempty"`
	InviteUserIds []*string `json:"InviteUserIds,omitempty" xml:"InviteUserIds,omitempty" type:"Repeated"`
}

func (s CreateVideoConferenceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateVideoConferenceRequest) GoString() string {
	return s.String()
}

func (s *CreateVideoConferenceRequest) GetConfTitle() *string {
	return s.ConfTitle
}

func (s *CreateVideoConferenceRequest) GetInviteCaller() *bool {
	return s.InviteCaller
}

func (s *CreateVideoConferenceRequest) GetInviteUserIds() []*string {
	return s.InviteUserIds
}

func (s *CreateVideoConferenceRequest) SetConfTitle(v string) *CreateVideoConferenceRequest {
	s.ConfTitle = &v
	return s
}

func (s *CreateVideoConferenceRequest) SetInviteCaller(v bool) *CreateVideoConferenceRequest {
	s.InviteCaller = &v
	return s
}

func (s *CreateVideoConferenceRequest) SetInviteUserIds(v []*string) *CreateVideoConferenceRequest {
	s.InviteUserIds = v
	return s
}

func (s *CreateVideoConferenceRequest) Validate() error {
	return dara.Validate(s)
}
