// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMinutesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTenantContextShrink(v string) *QueryMinutesShrinkRequest
	GetTenantContextShrink() *string
	SetConferenceId(v string) *QueryMinutesShrinkRequest
	GetConferenceId() *string
}

type QueryMinutesShrinkRequest struct {
	TenantContextShrink *string `json:"TenantContext,omitempty" xml:"TenantContext,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 61289fxxx
	ConferenceId *string `json:"conferenceId,omitempty" xml:"conferenceId,omitempty"`
}

func (s QueryMinutesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryMinutesShrinkRequest) GoString() string {
	return s.String()
}

func (s *QueryMinutesShrinkRequest) GetTenantContextShrink() *string {
	return s.TenantContextShrink
}

func (s *QueryMinutesShrinkRequest) GetConferenceId() *string {
	return s.ConferenceId
}

func (s *QueryMinutesShrinkRequest) SetTenantContextShrink(v string) *QueryMinutesShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *QueryMinutesShrinkRequest) SetConferenceId(v string) *QueryMinutesShrinkRequest {
	s.ConferenceId = &v
	return s
}

func (s *QueryMinutesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
