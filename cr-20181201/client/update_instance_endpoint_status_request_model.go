// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateInstanceEndpointStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnable(v bool) *UpdateInstanceEndpointStatusRequest
	GetEnable() *bool
	SetEndpointType(v string) *UpdateInstanceEndpointStatusRequest
	GetEndpointType() *string
	SetInstanceId(v string) *UpdateInstanceEndpointStatusRequest
	GetInstanceId() *string
	SetModuleName(v string) *UpdateInstanceEndpointStatusRequest
	GetModuleName() *string
}

type UpdateInstanceEndpointStatusRequest struct {
	// Specifies whether to enable the instance endpoint. Valid values:
	//
	// 	- `true`: enables the instance endpoint.
	//
	// 	- `false`: disables the instance endpoint
	//
	// This parameter is required.
	//
	// example:
	//
	// false
	Enable *bool `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// The type of the endpoint. Set the value to Internet.
	//
	// This parameter is required.
	//
	// example:
	//
	// internet
	EndpointType *string `json:"EndpointType,omitempty" xml:"EndpointType,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cri-kmsiwlxxdcva****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the module that you want to access. Valid values:
	//
	// 	- `Registry`: the image repository.
	//
	// 	- `Chart`: a Helm chart.
	//
	// example:
	//
	// Chart
	ModuleName *string `json:"ModuleName,omitempty" xml:"ModuleName,omitempty"`
}

func (s UpdateInstanceEndpointStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateInstanceEndpointStatusRequest) GoString() string {
	return s.String()
}

func (s *UpdateInstanceEndpointStatusRequest) GetEnable() *bool {
	return s.Enable
}

func (s *UpdateInstanceEndpointStatusRequest) GetEndpointType() *string {
	return s.EndpointType
}

func (s *UpdateInstanceEndpointStatusRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateInstanceEndpointStatusRequest) GetModuleName() *string {
	return s.ModuleName
}

func (s *UpdateInstanceEndpointStatusRequest) SetEnable(v bool) *UpdateInstanceEndpointStatusRequest {
	s.Enable = &v
	return s
}

func (s *UpdateInstanceEndpointStatusRequest) SetEndpointType(v string) *UpdateInstanceEndpointStatusRequest {
	s.EndpointType = &v
	return s
}

func (s *UpdateInstanceEndpointStatusRequest) SetInstanceId(v string) *UpdateInstanceEndpointStatusRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateInstanceEndpointStatusRequest) SetModuleName(v string) *UpdateInstanceEndpointStatusRequest {
	s.ModuleName = &v
	return s
}

func (s *UpdateInstanceEndpointStatusRequest) Validate() error {
	return dara.Validate(s)
}
