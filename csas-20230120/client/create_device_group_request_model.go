// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDeviceGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateDeviceGroupRequest
	GetDescription() *string
	SetDynamicOperator(v string) *CreateDeviceGroupRequest
	GetDynamicOperator() *string
	SetDynamicRule(v *Rule) *CreateDeviceGroupRequest
	GetDynamicRule() *Rule
	SetGroupType(v string) *CreateDeviceGroupRequest
	GetGroupType() *string
	SetName(v string) *CreateDeviceGroupRequest
	GetName() *string
}

type CreateDeviceGroupRequest struct {
	// The description of the device tag. The description can contain Chinese characters, uppercase and lowercase letters, digits, spaces, periods (.), underscores (_), and hyphens (-). This parameter can be left empty.
	//
	// example:
	//
	// Test device group description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Deprecated
	//
	// The operator for dynamic device group rules.
	//
	// example:
	//
	// AND
	DynamicOperator *string `json:"DynamicOperator,omitempty" xml:"DynamicOperator,omitempty"`
	// The matching rule for the dynamic device tag.
	DynamicRule *Rule `json:"DynamicRule,omitempty" xml:"DynamicRule,omitempty"`
	// The type of the device tag. Valid values:
	//
	// - **static**: static device tag. After the tag is created, you can manually add terminal devices by calling the [AddDeviceGroupMatchDevices](~~AddDeviceGroupMatchDevices~~) operation.
	//
	// - **dynamic**: dynamic device tag. Members are automatically matched based on the DynamicRule matching rule.
	//
	// example:
	//
	// static
	GroupType *string `json:"GroupType,omitempty" xml:"GroupType,omitempty"`
	// The name of the device tag. The name must be 1 to 128 characters in length and can contain Chinese characters, uppercase and lowercase letters, digits, periods (.), underscores (_), and hyphens (-). Spaces are not allowed.
	//
	// This parameter is required.
	//
	// example:
	//
	// example
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s CreateDeviceGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDeviceGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateDeviceGroupRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateDeviceGroupRequest) GetDynamicOperator() *string {
	return s.DynamicOperator
}

func (s *CreateDeviceGroupRequest) GetDynamicRule() *Rule {
	return s.DynamicRule
}

func (s *CreateDeviceGroupRequest) GetGroupType() *string {
	return s.GroupType
}

func (s *CreateDeviceGroupRequest) GetName() *string {
	return s.Name
}

func (s *CreateDeviceGroupRequest) SetDescription(v string) *CreateDeviceGroupRequest {
	s.Description = &v
	return s
}

func (s *CreateDeviceGroupRequest) SetDynamicOperator(v string) *CreateDeviceGroupRequest {
	s.DynamicOperator = &v
	return s
}

func (s *CreateDeviceGroupRequest) SetDynamicRule(v *Rule) *CreateDeviceGroupRequest {
	s.DynamicRule = v
	return s
}

func (s *CreateDeviceGroupRequest) SetGroupType(v string) *CreateDeviceGroupRequest {
	s.GroupType = &v
	return s
}

func (s *CreateDeviceGroupRequest) SetName(v string) *CreateDeviceGroupRequest {
	s.Name = &v
	return s
}

func (s *CreateDeviceGroupRequest) Validate() error {
	if s.DynamicRule != nil {
		if err := s.DynamicRule.Validate(); err != nil {
			return err
		}
	}
	return nil
}
