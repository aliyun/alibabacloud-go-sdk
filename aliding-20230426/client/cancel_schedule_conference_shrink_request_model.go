// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelScheduleConferenceShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetScheduleConferenceId(v string) *CancelScheduleConferenceShrinkRequest
	GetScheduleConferenceId() *string
	SetTenantContextShrink(v string) *CancelScheduleConferenceShrinkRequest
	GetTenantContextShrink() *string
}

type CancelScheduleConferenceShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 2a489xxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
	ScheduleConferenceId *string `json:"ScheduleConferenceId,omitempty" xml:"ScheduleConferenceId,omitempty"`
	TenantContextShrink  *string `json:"TenantContext,omitempty" xml:"TenantContext,omitempty"`
}

func (s CancelScheduleConferenceShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CancelScheduleConferenceShrinkRequest) GoString() string {
	return s.String()
}

func (s *CancelScheduleConferenceShrinkRequest) GetScheduleConferenceId() *string {
	return s.ScheduleConferenceId
}

func (s *CancelScheduleConferenceShrinkRequest) GetTenantContextShrink() *string {
	return s.TenantContextShrink
}

func (s *CancelScheduleConferenceShrinkRequest) SetScheduleConferenceId(v string) *CancelScheduleConferenceShrinkRequest {
	s.ScheduleConferenceId = &v
	return s
}

func (s *CancelScheduleConferenceShrinkRequest) SetTenantContextShrink(v string) *CancelScheduleConferenceShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *CancelScheduleConferenceShrinkRequest) Validate() error {
	return dara.Validate(s)
}
