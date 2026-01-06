// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDingtalkMeetingMetricDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTenantContext(v *GetDingtalkMeetingMetricDataRequestTenantContext) *GetDingtalkMeetingMetricDataRequest
	GetTenantContext() *GetDingtalkMeetingMetricDataRequestTenantContext
	SetBeginTime(v int64) *GetDingtalkMeetingMetricDataRequest
	GetBeginTime() *int64
	SetConferenceId(v string) *GetDingtalkMeetingMetricDataRequest
	GetConferenceId() *string
	SetEndTime(v int64) *GetDingtalkMeetingMetricDataRequest
	GetEndTime() *int64
	SetOrgId(v string) *GetDingtalkMeetingMetricDataRequest
	GetOrgId() *string
	SetTypeName(v string) *GetDingtalkMeetingMetricDataRequest
	GetTypeName() *string
	SetWorkNo(v string) *GetDingtalkMeetingMetricDataRequest
	GetWorkNo() *string
}

type GetDingtalkMeetingMetricDataRequest struct {
	TenantContext *GetDingtalkMeetingMetricDataRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
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

func (s GetDingtalkMeetingMetricDataRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDingtalkMeetingMetricDataRequest) GoString() string {
	return s.String()
}

func (s *GetDingtalkMeetingMetricDataRequest) GetTenantContext() *GetDingtalkMeetingMetricDataRequestTenantContext {
	return s.TenantContext
}

func (s *GetDingtalkMeetingMetricDataRequest) GetBeginTime() *int64 {
	return s.BeginTime
}

func (s *GetDingtalkMeetingMetricDataRequest) GetConferenceId() *string {
	return s.ConferenceId
}

func (s *GetDingtalkMeetingMetricDataRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *GetDingtalkMeetingMetricDataRequest) GetOrgId() *string {
	return s.OrgId
}

func (s *GetDingtalkMeetingMetricDataRequest) GetTypeName() *string {
	return s.TypeName
}

func (s *GetDingtalkMeetingMetricDataRequest) GetWorkNo() *string {
	return s.WorkNo
}

func (s *GetDingtalkMeetingMetricDataRequest) SetTenantContext(v *GetDingtalkMeetingMetricDataRequestTenantContext) *GetDingtalkMeetingMetricDataRequest {
	s.TenantContext = v
	return s
}

func (s *GetDingtalkMeetingMetricDataRequest) SetBeginTime(v int64) *GetDingtalkMeetingMetricDataRequest {
	s.BeginTime = &v
	return s
}

func (s *GetDingtalkMeetingMetricDataRequest) SetConferenceId(v string) *GetDingtalkMeetingMetricDataRequest {
	s.ConferenceId = &v
	return s
}

func (s *GetDingtalkMeetingMetricDataRequest) SetEndTime(v int64) *GetDingtalkMeetingMetricDataRequest {
	s.EndTime = &v
	return s
}

func (s *GetDingtalkMeetingMetricDataRequest) SetOrgId(v string) *GetDingtalkMeetingMetricDataRequest {
	s.OrgId = &v
	return s
}

func (s *GetDingtalkMeetingMetricDataRequest) SetTypeName(v string) *GetDingtalkMeetingMetricDataRequest {
	s.TypeName = &v
	return s
}

func (s *GetDingtalkMeetingMetricDataRequest) SetWorkNo(v string) *GetDingtalkMeetingMetricDataRequest {
	s.WorkNo = &v
	return s
}

func (s *GetDingtalkMeetingMetricDataRequest) Validate() error {
	if s.TenantContext != nil {
		if err := s.TenantContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetDingtalkMeetingMetricDataRequestTenantContext struct {
	// example:
	//
	// xxxxxx
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s GetDingtalkMeetingMetricDataRequestTenantContext) String() string {
	return dara.Prettify(s)
}

func (s GetDingtalkMeetingMetricDataRequestTenantContext) GoString() string {
	return s.String()
}

func (s *GetDingtalkMeetingMetricDataRequestTenantContext) GetTenantId() *string {
	return s.TenantId
}

func (s *GetDingtalkMeetingMetricDataRequestTenantContext) SetTenantId(v string) *GetDingtalkMeetingMetricDataRequestTenantContext {
	s.TenantId = &v
	return s
}

func (s *GetDingtalkMeetingMetricDataRequestTenantContext) Validate() error {
	return dara.Validate(s)
}
