// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartMinutesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTenantContextShrink(v string) *StartMinutesShrinkRequest
	GetTenantContextShrink() *string
	SetConferenceId(v string) *StartMinutesShrinkRequest
	GetConferenceId() *string
	SetOwnerUserId(v string) *StartMinutesShrinkRequest
	GetOwnerUserId() *string
	SetRecordAudio(v bool) *StartMinutesShrinkRequest
	GetRecordAudio() *bool
}

type StartMinutesShrinkRequest struct {
	TenantContextShrink *string `json:"TenantContext,omitempty" xml:"TenantContext,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 61289fxxx
	ConferenceId *string `json:"conferenceId,omitempty" xml:"conferenceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 012345
	OwnerUserId *string `json:"ownerUserId,omitempty" xml:"ownerUserId,omitempty"`
	// example:
	//
	// false
	RecordAudio *bool `json:"recordAudio,omitempty" xml:"recordAudio,omitempty"`
}

func (s StartMinutesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s StartMinutesShrinkRequest) GoString() string {
	return s.String()
}

func (s *StartMinutesShrinkRequest) GetTenantContextShrink() *string {
	return s.TenantContextShrink
}

func (s *StartMinutesShrinkRequest) GetConferenceId() *string {
	return s.ConferenceId
}

func (s *StartMinutesShrinkRequest) GetOwnerUserId() *string {
	return s.OwnerUserId
}

func (s *StartMinutesShrinkRequest) GetRecordAudio() *bool {
	return s.RecordAudio
}

func (s *StartMinutesShrinkRequest) SetTenantContextShrink(v string) *StartMinutesShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *StartMinutesShrinkRequest) SetConferenceId(v string) *StartMinutesShrinkRequest {
	s.ConferenceId = &v
	return s
}

func (s *StartMinutesShrinkRequest) SetOwnerUserId(v string) *StartMinutesShrinkRequest {
	s.OwnerUserId = &v
	return s
}

func (s *StartMinutesShrinkRequest) SetRecordAudio(v bool) *StartMinutesShrinkRequest {
	s.RecordAudio = &v
	return s
}

func (s *StartMinutesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
