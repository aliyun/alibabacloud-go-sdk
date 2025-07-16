// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateScheduleConfSettingsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetScheduleConfSettingModelShrink(v string) *UpdateScheduleConfSettingsShrinkRequest
	GetScheduleConfSettingModelShrink() *string
	SetScheduleConferenceId(v string) *UpdateScheduleConfSettingsShrinkRequest
	GetScheduleConferenceId() *string
	SetTenantContextShrink(v string) *UpdateScheduleConfSettingsShrinkRequest
	GetTenantContextShrink() *string
}

type UpdateScheduleConfSettingsShrinkRequest struct {
	ScheduleConfSettingModelShrink *string `json:"ScheduleConfSettingModel,omitempty" xml:"ScheduleConfSettingModel,omitempty"`
	// example:
	//
	// f6fb627e-a7e8-403e-b1f8-26e85450f4a9
	ScheduleConferenceId *string `json:"ScheduleConferenceId,omitempty" xml:"ScheduleConferenceId,omitempty"`
	TenantContextShrink  *string `json:"TenantContext,omitempty" xml:"TenantContext,omitempty"`
}

func (s UpdateScheduleConfSettingsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateScheduleConfSettingsShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateScheduleConfSettingsShrinkRequest) GetScheduleConfSettingModelShrink() *string {
	return s.ScheduleConfSettingModelShrink
}

func (s *UpdateScheduleConfSettingsShrinkRequest) GetScheduleConferenceId() *string {
	return s.ScheduleConferenceId
}

func (s *UpdateScheduleConfSettingsShrinkRequest) GetTenantContextShrink() *string {
	return s.TenantContextShrink
}

func (s *UpdateScheduleConfSettingsShrinkRequest) SetScheduleConfSettingModelShrink(v string) *UpdateScheduleConfSettingsShrinkRequest {
	s.ScheduleConfSettingModelShrink = &v
	return s
}

func (s *UpdateScheduleConfSettingsShrinkRequest) SetScheduleConferenceId(v string) *UpdateScheduleConfSettingsShrinkRequest {
	s.ScheduleConferenceId = &v
	return s
}

func (s *UpdateScheduleConfSettingsShrinkRequest) SetTenantContextShrink(v string) *UpdateScheduleConfSettingsShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *UpdateScheduleConfSettingsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
