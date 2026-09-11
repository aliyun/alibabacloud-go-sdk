// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDeviceGroupShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *UpdateDeviceGroupShrinkRequest
	GetDescription() *string
	SetDeviceGroupId(v string) *UpdateDeviceGroupShrinkRequest
	GetDeviceGroupId() *string
	SetDynamicOperator(v string) *UpdateDeviceGroupShrinkRequest
	GetDynamicOperator() *string
	SetDynamicRuleShrink(v string) *UpdateDeviceGroupShrinkRequest
	GetDynamicRuleShrink() *string
	SetName(v string) *UpdateDeviceGroupShrinkRequest
	GetName() *string
}

type UpdateDeviceGroupShrinkRequest struct {
	// The description of the device label. Set this parameter to an empty string to clear the description. The description can contain letters, digits, Chinese characters, spaces, periods (.), underscores (_), and hyphens (-).
	//
	// example:
	//
	// Test device group description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the device label.
	//
	// example:
	//
	// device-group-5191cf830a5e****
	DeviceGroupId *string `json:"DeviceGroupId,omitempty" xml:"DeviceGroupId,omitempty"`
	// Deprecated
	//
	// The operator of the dynamic device group rule.
	//
	// example:
	//
	// AND
	DynamicOperator *string `json:"DynamicOperator,omitempty" xml:"DynamicOperator,omitempty"`
	// The matching rule of the dynamic device label.
	DynamicRuleShrink *string `json:"DynamicRule,omitempty" xml:"DynamicRule,omitempty"`
	// The name of the device label. The name must be 1 to 128 characters in length and can contain letters, digits, Chinese characters, periods (.), underscores (_), and hyphens (-). Spaces are not supported.
	//
	// example:
	//
	// example
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s UpdateDeviceGroupShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateDeviceGroupShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateDeviceGroupShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateDeviceGroupShrinkRequest) GetDeviceGroupId() *string {
	return s.DeviceGroupId
}

func (s *UpdateDeviceGroupShrinkRequest) GetDynamicOperator() *string {
	return s.DynamicOperator
}

func (s *UpdateDeviceGroupShrinkRequest) GetDynamicRuleShrink() *string {
	return s.DynamicRuleShrink
}

func (s *UpdateDeviceGroupShrinkRequest) GetName() *string {
	return s.Name
}

func (s *UpdateDeviceGroupShrinkRequest) SetDescription(v string) *UpdateDeviceGroupShrinkRequest {
	s.Description = &v
	return s
}

func (s *UpdateDeviceGroupShrinkRequest) SetDeviceGroupId(v string) *UpdateDeviceGroupShrinkRequest {
	s.DeviceGroupId = &v
	return s
}

func (s *UpdateDeviceGroupShrinkRequest) SetDynamicOperator(v string) *UpdateDeviceGroupShrinkRequest {
	s.DynamicOperator = &v
	return s
}

func (s *UpdateDeviceGroupShrinkRequest) SetDynamicRuleShrink(v string) *UpdateDeviceGroupShrinkRequest {
	s.DynamicRuleShrink = &v
	return s
}

func (s *UpdateDeviceGroupShrinkRequest) SetName(v string) *UpdateDeviceGroupShrinkRequest {
	s.Name = &v
	return s
}

func (s *UpdateDeviceGroupShrinkRequest) Validate() error {
	return dara.Validate(s)
}
