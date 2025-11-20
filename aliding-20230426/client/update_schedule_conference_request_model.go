// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateScheduleConferenceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v int64) *UpdateScheduleConferenceRequest
	GetEndTime() *int64
	SetScheduleConferenceId(v string) *UpdateScheduleConferenceRequest
	GetScheduleConferenceId() *string
	SetStartTime(v int64) *UpdateScheduleConferenceRequest
	GetStartTime() *int64
	SetTenantContext(v *UpdateScheduleConferenceRequestTenantContext) *UpdateScheduleConferenceRequest
	GetTenantContext() *UpdateScheduleConferenceRequestTenantContext
	SetTitle(v string) *UpdateScheduleConferenceRequest
	GetTitle() *string
}

type UpdateScheduleConferenceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1687928400000L
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2a489xxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
	ScheduleConferenceId *string `json:"ScheduleConferenceId,omitempty" xml:"ScheduleConferenceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1687924800000L
	StartTime     *int64                                        `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	TenantContext *UpdateScheduleConferenceRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 预约会议标题
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s UpdateScheduleConferenceRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateScheduleConferenceRequest) GoString() string {
	return s.String()
}

func (s *UpdateScheduleConferenceRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *UpdateScheduleConferenceRequest) GetScheduleConferenceId() *string {
	return s.ScheduleConferenceId
}

func (s *UpdateScheduleConferenceRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *UpdateScheduleConferenceRequest) GetTenantContext() *UpdateScheduleConferenceRequestTenantContext {
	return s.TenantContext
}

func (s *UpdateScheduleConferenceRequest) GetTitle() *string {
	return s.Title
}

func (s *UpdateScheduleConferenceRequest) SetEndTime(v int64) *UpdateScheduleConferenceRequest {
	s.EndTime = &v
	return s
}

func (s *UpdateScheduleConferenceRequest) SetScheduleConferenceId(v string) *UpdateScheduleConferenceRequest {
	s.ScheduleConferenceId = &v
	return s
}

func (s *UpdateScheduleConferenceRequest) SetStartTime(v int64) *UpdateScheduleConferenceRequest {
	s.StartTime = &v
	return s
}

func (s *UpdateScheduleConferenceRequest) SetTenantContext(v *UpdateScheduleConferenceRequestTenantContext) *UpdateScheduleConferenceRequest {
	s.TenantContext = v
	return s
}

func (s *UpdateScheduleConferenceRequest) SetTitle(v string) *UpdateScheduleConferenceRequest {
	s.Title = &v
	return s
}

func (s *UpdateScheduleConferenceRequest) Validate() error {
	if s.TenantContext != nil {
		if err := s.TenantContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateScheduleConferenceRequestTenantContext struct {
	// example:
	//
	// xxxxxx
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s UpdateScheduleConferenceRequestTenantContext) String() string {
	return dara.Prettify(s)
}

func (s UpdateScheduleConferenceRequestTenantContext) GoString() string {
	return s.String()
}

func (s *UpdateScheduleConferenceRequestTenantContext) GetTenantId() *string {
	return s.TenantId
}

func (s *UpdateScheduleConferenceRequestTenantContext) SetTenantId(v string) *UpdateScheduleConferenceRequestTenantContext {
	s.TenantId = &v
	return s
}

func (s *UpdateScheduleConferenceRequestTenantContext) Validate() error {
	return dara.Validate(s)
}
