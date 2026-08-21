// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDeviceGroupShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateDeviceGroupShrinkRequest
	GetDescription() *string
	SetDynamicOperator(v string) *CreateDeviceGroupShrinkRequest
	GetDynamicOperator() *string
	SetDynamicRuleShrink(v string) *CreateDeviceGroupShrinkRequest
	GetDynamicRuleShrink() *string
	SetGroupType(v string) *CreateDeviceGroupShrinkRequest
	GetGroupType() *string
	SetName(v string) *CreateDeviceGroupShrinkRequest
	GetName() *string
}

type CreateDeviceGroupShrinkRequest struct {
	// The description of the device label. The description can contain letters, digits, Chinese characters, spaces, periods (.), underscores (_), and hyphens (-). This parameter can be left empty.
	//
	// example:
	//
	// Test device group description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
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
	// The type of the device label. Valid values:
	//
	// - **static**: static device label. After creation, manually add terminal devices by calling [AddDeviceGroupMatchDevices](~~AddDeviceGroupMatchDevices~~).
	//
	// - **dynamic**: dynamic device label. Members are automatically matched by the DynamicRule matching rule.
	//
	// example:
	//
	// static
	GroupType *string `json:"GroupType,omitempty" xml:"GroupType,omitempty"`
	// The name of the device label. The name must be 1 to 128 characters in length and can contain letters, digits, Chinese characters, periods (.), underscores (_), and hyphens (-). Spaces are not supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// example
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s CreateDeviceGroupShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDeviceGroupShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateDeviceGroupShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateDeviceGroupShrinkRequest) GetDynamicOperator() *string {
	return s.DynamicOperator
}

func (s *CreateDeviceGroupShrinkRequest) GetDynamicRuleShrink() *string {
	return s.DynamicRuleShrink
}

func (s *CreateDeviceGroupShrinkRequest) GetGroupType() *string {
	return s.GroupType
}

func (s *CreateDeviceGroupShrinkRequest) GetName() *string {
	return s.Name
}

func (s *CreateDeviceGroupShrinkRequest) SetDescription(v string) *CreateDeviceGroupShrinkRequest {
	s.Description = &v
	return s
}

func (s *CreateDeviceGroupShrinkRequest) SetDynamicOperator(v string) *CreateDeviceGroupShrinkRequest {
	s.DynamicOperator = &v
	return s
}

func (s *CreateDeviceGroupShrinkRequest) SetDynamicRuleShrink(v string) *CreateDeviceGroupShrinkRequest {
	s.DynamicRuleShrink = &v
	return s
}

func (s *CreateDeviceGroupShrinkRequest) SetGroupType(v string) *CreateDeviceGroupShrinkRequest {
	s.GroupType = &v
	return s
}

func (s *CreateDeviceGroupShrinkRequest) SetName(v string) *CreateDeviceGroupShrinkRequest {
	s.Name = &v
	return s
}

func (s *CreateDeviceGroupShrinkRequest) Validate() error {
	return dara.Validate(s)
}
