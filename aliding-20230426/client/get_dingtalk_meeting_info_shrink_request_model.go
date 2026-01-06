// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDingtalkMeetingInfoShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTenantContextShrink(v string) *GetDingtalkMeetingInfoShrinkRequest
	GetTenantContextShrink() *string
	SetConferenceId(v string) *GetDingtalkMeetingInfoShrinkRequest
	GetConferenceId() *string
	SetOrgId(v string) *GetDingtalkMeetingInfoShrinkRequest
	GetOrgId() *string
}

type GetDingtalkMeetingInfoShrinkRequest struct {
	TenantContextShrink *string `json:"TenantContext,omitempty" xml:"TenantContext,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 61289fxxx
	ConferenceId *string `json:"conferenceId,omitempty" xml:"conferenceId,omitempty"`
	// example:
	//
	// 21001
	OrgId *string `json:"orgId,omitempty" xml:"orgId,omitempty"`
}

func (s GetDingtalkMeetingInfoShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDingtalkMeetingInfoShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetDingtalkMeetingInfoShrinkRequest) GetTenantContextShrink() *string {
	return s.TenantContextShrink
}

func (s *GetDingtalkMeetingInfoShrinkRequest) GetConferenceId() *string {
	return s.ConferenceId
}

func (s *GetDingtalkMeetingInfoShrinkRequest) GetOrgId() *string {
	return s.OrgId
}

func (s *GetDingtalkMeetingInfoShrinkRequest) SetTenantContextShrink(v string) *GetDingtalkMeetingInfoShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *GetDingtalkMeetingInfoShrinkRequest) SetConferenceId(v string) *GetDingtalkMeetingInfoShrinkRequest {
	s.ConferenceId = &v
	return s
}

func (s *GetDingtalkMeetingInfoShrinkRequest) SetOrgId(v string) *GetDingtalkMeetingInfoShrinkRequest {
	s.OrgId = &v
	return s
}

func (s *GetDingtalkMeetingInfoShrinkRequest) Validate() error {
	return dara.Validate(s)
}
