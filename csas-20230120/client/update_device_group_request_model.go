// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDeviceGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *UpdateDeviceGroupRequest
	GetDescription() *string
	SetDeviceGroupId(v string) *UpdateDeviceGroupRequest
	GetDeviceGroupId() *string
	SetDynamicOperator(v string) *UpdateDeviceGroupRequest
	GetDynamicOperator() *string
	SetName(v string) *UpdateDeviceGroupRequest
	GetName() *string
}

type UpdateDeviceGroupRequest struct {
	// The description of the device label. If you pass in an empty string, the description is cleared. The description can contain letters, digits, spaces, periods (.), underscores (_), and hyphens (-). Chinese characters are supported.
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
	// The rule operator of the dynamic device group.
	//
	// example:
	//
	// AND
	DynamicOperator *string `json:"DynamicOperator,omitempty" xml:"DynamicOperator,omitempty"`
	// The name of the device label. The name must be 1 to 128 characters in length and can contain letters, digits, periods (.), underscores (_), and hyphens (-). Chinese characters are supported. Spaces are not supported.
	//
	// example:
	//
	// example
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s UpdateDeviceGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateDeviceGroupRequest) GoString() string {
	return s.String()
}

func (s *UpdateDeviceGroupRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateDeviceGroupRequest) GetDeviceGroupId() *string {
	return s.DeviceGroupId
}

func (s *UpdateDeviceGroupRequest) GetDynamicOperator() *string {
	return s.DynamicOperator
}

func (s *UpdateDeviceGroupRequest) GetName() *string {
	return s.Name
}

func (s *UpdateDeviceGroupRequest) SetDescription(v string) *UpdateDeviceGroupRequest {
	s.Description = &v
	return s
}

func (s *UpdateDeviceGroupRequest) SetDeviceGroupId(v string) *UpdateDeviceGroupRequest {
	s.DeviceGroupId = &v
	return s
}

func (s *UpdateDeviceGroupRequest) SetDynamicOperator(v string) *UpdateDeviceGroupRequest {
	s.DynamicOperator = &v
	return s
}

func (s *UpdateDeviceGroupRequest) SetName(v string) *UpdateDeviceGroupRequest {
	s.Name = &v
	return s
}

func (s *UpdateDeviceGroupRequest) Validate() error {
	return dara.Validate(s)
}
