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
	DynamicRule *Rule `json:"DynamicRule,omitempty" xml:"DynamicRule,omitempty"`
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
