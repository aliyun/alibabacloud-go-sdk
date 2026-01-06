// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDingtalkMeetingMetricDataShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTenantContextShrink(v string) *GetDingtalkMeetingMetricDataShrinkRequest
	GetTenantContextShrink() *string
	SetBeginTime(v int64) *GetDingtalkMeetingMetricDataShrinkRequest
	GetBeginTime() *int64
	SetConferenceId(v string) *GetDingtalkMeetingMetricDataShrinkRequest
	GetConferenceId() *string
	SetEndTime(v int64) *GetDingtalkMeetingMetricDataShrinkRequest
	GetEndTime() *int64
	SetOrgId(v string) *GetDingtalkMeetingMetricDataShrinkRequest
	GetOrgId() *string
	SetTypeName(v string) *GetDingtalkMeetingMetricDataShrinkRequest
	GetTypeName() *string
	SetWorkNo(v string) *GetDingtalkMeetingMetricDataShrinkRequest
	GetWorkNo() *string
}

type GetDingtalkMeetingMetricDataShrinkRequest struct {
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
	// audio
	TypeName *string `json:"typeName,omitempty" xml:"typeName,omitempty"`
	// example:
	//
	// 123456
	WorkNo *string `json:"workNo,omitempty" xml:"workNo,omitempty"`
}

func (s GetDingtalkMeetingMetricDataShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDingtalkMeetingMetricDataShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetDingtalkMeetingMetricDataShrinkRequest) GetTenantContextShrink() *string {
	return s.TenantContextShrink
}

func (s *GetDingtalkMeetingMetricDataShrinkRequest) GetBeginTime() *int64 {
	return s.BeginTime
}

func (s *GetDingtalkMeetingMetricDataShrinkRequest) GetConferenceId() *string {
	return s.ConferenceId
}

func (s *GetDingtalkMeetingMetricDataShrinkRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *GetDingtalkMeetingMetricDataShrinkRequest) GetOrgId() *string {
	return s.OrgId
}

func (s *GetDingtalkMeetingMetricDataShrinkRequest) GetTypeName() *string {
	return s.TypeName
}

func (s *GetDingtalkMeetingMetricDataShrinkRequest) GetWorkNo() *string {
	return s.WorkNo
}

func (s *GetDingtalkMeetingMetricDataShrinkRequest) SetTenantContextShrink(v string) *GetDingtalkMeetingMetricDataShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *GetDingtalkMeetingMetricDataShrinkRequest) SetBeginTime(v int64) *GetDingtalkMeetingMetricDataShrinkRequest {
	s.BeginTime = &v
	return s
}

func (s *GetDingtalkMeetingMetricDataShrinkRequest) SetConferenceId(v string) *GetDingtalkMeetingMetricDataShrinkRequest {
	s.ConferenceId = &v
	return s
}

func (s *GetDingtalkMeetingMetricDataShrinkRequest) SetEndTime(v int64) *GetDingtalkMeetingMetricDataShrinkRequest {
	s.EndTime = &v
	return s
}

func (s *GetDingtalkMeetingMetricDataShrinkRequest) SetOrgId(v string) *GetDingtalkMeetingMetricDataShrinkRequest {
	s.OrgId = &v
	return s
}

func (s *GetDingtalkMeetingMetricDataShrinkRequest) SetTypeName(v string) *GetDingtalkMeetingMetricDataShrinkRequest {
	s.TypeName = &v
	return s
}

func (s *GetDingtalkMeetingMetricDataShrinkRequest) SetWorkNo(v string) *GetDingtalkMeetingMetricDataShrinkRequest {
	s.WorkNo = &v
	return s
}

func (s *GetDingtalkMeetingMetricDataShrinkRequest) Validate() error {
	return dara.Validate(s)
}
