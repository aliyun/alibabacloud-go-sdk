// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryScheduleConferenceShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTenantContextShrink(v string) *QueryScheduleConferenceShrinkRequest
	GetTenantContextShrink() *string
	SetScheduleConferenceId(v string) *QueryScheduleConferenceShrinkRequest
	GetScheduleConferenceId() *string
}

type QueryScheduleConferenceShrinkRequest struct {
	TenantContextShrink *string `json:"TenantContext,omitempty" xml:"TenantContext,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2a489c68-xxxx-xxxx-xxxx-xxxxxxxxxxxx
	ScheduleConferenceId *string `json:"scheduleConferenceId,omitempty" xml:"scheduleConferenceId,omitempty"`
}

func (s QueryScheduleConferenceShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryScheduleConferenceShrinkRequest) GoString() string {
	return s.String()
}

func (s *QueryScheduleConferenceShrinkRequest) GetTenantContextShrink() *string {
	return s.TenantContextShrink
}

func (s *QueryScheduleConferenceShrinkRequest) GetScheduleConferenceId() *string {
	return s.ScheduleConferenceId
}

func (s *QueryScheduleConferenceShrinkRequest) SetTenantContextShrink(v string) *QueryScheduleConferenceShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *QueryScheduleConferenceShrinkRequest) SetScheduleConferenceId(v string) *QueryScheduleConferenceShrinkRequest {
	s.ScheduleConferenceId = &v
	return s
}

func (s *QueryScheduleConferenceShrinkRequest) Validate() error {
	return dara.Validate(s)
}
