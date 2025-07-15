// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateApplicationShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlarmConfigShrink(v string) *UpdateApplicationShrinkRequest
	GetAlarmConfigShrink() *string
	SetDeleteAlarmRulesBeforeUpdate(v bool) *UpdateApplicationShrinkRequest
	GetDeleteAlarmRulesBeforeUpdate() *bool
	SetDescription(v string) *UpdateApplicationShrinkRequest
	GetDescription() *string
	SetName(v string) *UpdateApplicationShrinkRequest
	GetName() *string
	SetRegionId(v string) *UpdateApplicationShrinkRequest
	GetRegionId() *string
	SetTagsShrink(v string) *UpdateApplicationShrinkRequest
	GetTagsShrink() *string
}

type UpdateApplicationShrinkRequest struct {
	// The configurations of application alerts.
	AlarmConfigShrink *string `json:"AlarmConfig,omitempty" xml:"AlarmConfig,omitempty"`
	// Specifies whether to delete existing alert rules before applying the alert template. Default value: false.
	//
	// example:
	//
	// false
	DeleteAlarmRulesBeforeUpdate *bool `json:"DeleteAlarmRulesBeforeUpdate,omitempty" xml:"DeleteAlarmRulesBeforeUpdate,omitempty"`
	// The description to be updated for the application.
	//
	// example:
	//
	// test application
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The application name.
	//
	// This parameter is required.
	//
	// example:
	//
	// My-Application
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The region ID. Set the value to cn-hangzhou.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The tags.
	//
	// example:
	//
	// {"k1":"v1","k2":"v2"}
	TagsShrink *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
}

func (s UpdateApplicationShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateApplicationShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateApplicationShrinkRequest) GetAlarmConfigShrink() *string {
	return s.AlarmConfigShrink
}

func (s *UpdateApplicationShrinkRequest) GetDeleteAlarmRulesBeforeUpdate() *bool {
	return s.DeleteAlarmRulesBeforeUpdate
}

func (s *UpdateApplicationShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateApplicationShrinkRequest) GetName() *string {
	return s.Name
}

func (s *UpdateApplicationShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateApplicationShrinkRequest) GetTagsShrink() *string {
	return s.TagsShrink
}

func (s *UpdateApplicationShrinkRequest) SetAlarmConfigShrink(v string) *UpdateApplicationShrinkRequest {
	s.AlarmConfigShrink = &v
	return s
}

func (s *UpdateApplicationShrinkRequest) SetDeleteAlarmRulesBeforeUpdate(v bool) *UpdateApplicationShrinkRequest {
	s.DeleteAlarmRulesBeforeUpdate = &v
	return s
}

func (s *UpdateApplicationShrinkRequest) SetDescription(v string) *UpdateApplicationShrinkRequest {
	s.Description = &v
	return s
}

func (s *UpdateApplicationShrinkRequest) SetName(v string) *UpdateApplicationShrinkRequest {
	s.Name = &v
	return s
}

func (s *UpdateApplicationShrinkRequest) SetRegionId(v string) *UpdateApplicationShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateApplicationShrinkRequest) SetTagsShrink(v string) *UpdateApplicationShrinkRequest {
	s.TagsShrink = &v
	return s
}

func (s *UpdateApplicationShrinkRequest) Validate() error {
	return dara.Validate(s)
}
