// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDingtalkMeetingMemberEventShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTenantContextShrink(v string) *GetDingtalkMeetingMemberEventShrinkRequest
	GetTenantContextShrink() *string
	SetBeginTime(v int64) *GetDingtalkMeetingMemberEventShrinkRequest
	GetBeginTime() *int64
	SetConferenceId(v string) *GetDingtalkMeetingMemberEventShrinkRequest
	GetConferenceId() *string
	SetEndTime(v int64) *GetDingtalkMeetingMemberEventShrinkRequest
	GetEndTime() *int64
	SetOrgId(v string) *GetDingtalkMeetingMemberEventShrinkRequest
	GetOrgId() *string
	SetWorkNo(v string) *GetDingtalkMeetingMemberEventShrinkRequest
	GetWorkNo() *string
}

type GetDingtalkMeetingMemberEventShrinkRequest struct {
	TenantContextShrink *string `json:"TenantContext,omitempty" xml:"TenantContext,omitempty"`
	// example:
	//
	// 1638360000000
	BeginTime *int64 `json:"beginTime,omitempty" xml:"beginTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 61289fxxx
	ConferenceId *string `json:"conferenceId,omitempty" xml:"conferenceId,omitempty"`
	// example:
	//
	// 1638363600000
	EndTime *int64 `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// example:
	//
	// 21001
	OrgId *string `json:"orgId,omitempty" xml:"orgId,omitempty"`
	// example:
	//
	// 34242
	WorkNo *string `json:"workNo,omitempty" xml:"workNo,omitempty"`
}

func (s GetDingtalkMeetingMemberEventShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDingtalkMeetingMemberEventShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetDingtalkMeetingMemberEventShrinkRequest) GetTenantContextShrink() *string {
	return s.TenantContextShrink
}

func (s *GetDingtalkMeetingMemberEventShrinkRequest) GetBeginTime() *int64 {
	return s.BeginTime
}

func (s *GetDingtalkMeetingMemberEventShrinkRequest) GetConferenceId() *string {
	return s.ConferenceId
}

func (s *GetDingtalkMeetingMemberEventShrinkRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *GetDingtalkMeetingMemberEventShrinkRequest) GetOrgId() *string {
	return s.OrgId
}

func (s *GetDingtalkMeetingMemberEventShrinkRequest) GetWorkNo() *string {
	return s.WorkNo
}

func (s *GetDingtalkMeetingMemberEventShrinkRequest) SetTenantContextShrink(v string) *GetDingtalkMeetingMemberEventShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *GetDingtalkMeetingMemberEventShrinkRequest) SetBeginTime(v int64) *GetDingtalkMeetingMemberEventShrinkRequest {
	s.BeginTime = &v
	return s
}

func (s *GetDingtalkMeetingMemberEventShrinkRequest) SetConferenceId(v string) *GetDingtalkMeetingMemberEventShrinkRequest {
	s.ConferenceId = &v
	return s
}

func (s *GetDingtalkMeetingMemberEventShrinkRequest) SetEndTime(v int64) *GetDingtalkMeetingMemberEventShrinkRequest {
	s.EndTime = &v
	return s
}

func (s *GetDingtalkMeetingMemberEventShrinkRequest) SetOrgId(v string) *GetDingtalkMeetingMemberEventShrinkRequest {
	s.OrgId = &v
	return s
}

func (s *GetDingtalkMeetingMemberEventShrinkRequest) SetWorkNo(v string) *GetDingtalkMeetingMemberEventShrinkRequest {
	s.WorkNo = &v
	return s
}

func (s *GetDingtalkMeetingMemberEventShrinkRequest) Validate() error {
	return dara.Validate(s)
}
