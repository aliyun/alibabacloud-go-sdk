// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMeetingFlashMinutesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTenantContextShrink(v string) *MeetingFlashMinutesShrinkRequest
	GetTenantContextShrink() *string
	SetConferenceId(v string) *MeetingFlashMinutesShrinkRequest
	GetConferenceId() *string
}

type MeetingFlashMinutesShrinkRequest struct {
	TenantContextShrink *string `json:"TenantContext,omitempty" xml:"TenantContext,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 61289fxxx
	ConferenceId *string `json:"conferenceId,omitempty" xml:"conferenceId,omitempty"`
}

func (s MeetingFlashMinutesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s MeetingFlashMinutesShrinkRequest) GoString() string {
	return s.String()
}

func (s *MeetingFlashMinutesShrinkRequest) GetTenantContextShrink() *string {
	return s.TenantContextShrink
}

func (s *MeetingFlashMinutesShrinkRequest) GetConferenceId() *string {
	return s.ConferenceId
}

func (s *MeetingFlashMinutesShrinkRequest) SetTenantContextShrink(v string) *MeetingFlashMinutesShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *MeetingFlashMinutesShrinkRequest) SetConferenceId(v string) *MeetingFlashMinutesShrinkRequest {
	s.ConferenceId = &v
	return s
}

func (s *MeetingFlashMinutesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
