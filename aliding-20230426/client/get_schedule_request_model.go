// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetScheduleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *GetScheduleRequest
	GetEndTime() *string
	SetStartTime(v string) *GetScheduleRequest
	GetStartTime() *string
	SetTenantContext(v *GetScheduleRequestTenantContext) *GetScheduleRequest
	GetTenantContext() *GetScheduleRequestTenantContext
	SetUserIds(v []*string) *GetScheduleRequest
	GetUserIds() []*string
}

type GetScheduleRequest struct {
	// example:
	//
	// 2020-01-01T10:15:30+08:00
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 2020-01-01T10:15:30+08:00
	StartTime     *string                          `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	TenantContext *GetScheduleRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
	UserIds       []*string                        `json:"UserIds,omitempty" xml:"UserIds,omitempty" type:"Repeated"`
}

func (s GetScheduleRequest) String() string {
	return dara.Prettify(s)
}

func (s GetScheduleRequest) GoString() string {
	return s.String()
}

func (s *GetScheduleRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *GetScheduleRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *GetScheduleRequest) GetTenantContext() *GetScheduleRequestTenantContext {
	return s.TenantContext
}

func (s *GetScheduleRequest) GetUserIds() []*string {
	return s.UserIds
}

func (s *GetScheduleRequest) SetEndTime(v string) *GetScheduleRequest {
	s.EndTime = &v
	return s
}

func (s *GetScheduleRequest) SetStartTime(v string) *GetScheduleRequest {
	s.StartTime = &v
	return s
}

func (s *GetScheduleRequest) SetTenantContext(v *GetScheduleRequestTenantContext) *GetScheduleRequest {
	s.TenantContext = v
	return s
}

func (s *GetScheduleRequest) SetUserIds(v []*string) *GetScheduleRequest {
	s.UserIds = v
	return s
}

func (s *GetScheduleRequest) Validate() error {
	if s.TenantContext != nil {
		if err := s.TenantContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetScheduleRequestTenantContext struct {
	// example:
	//
	// xxxxxx
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s GetScheduleRequestTenantContext) String() string {
	return dara.Prettify(s)
}

func (s GetScheduleRequestTenantContext) GoString() string {
	return s.String()
}

func (s *GetScheduleRequestTenantContext) GetTenantId() *string {
	return s.TenantId
}

func (s *GetScheduleRequestTenantContext) SetTenantId(v string) *GetScheduleRequestTenantContext {
	s.TenantId = &v
	return s
}

func (s *GetScheduleRequestTenantContext) Validate() error {
	return dara.Validate(s)
}
