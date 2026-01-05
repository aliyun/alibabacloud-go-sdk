// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateEndpointRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *CreateEndpointRequest
	GetInstanceId() *string
	SetName(v string) *CreateEndpointRequest
	GetName() *string
	SetType(v string) *CreateEndpointRequest
	GetType() *string
	SetVpcId(v string) *CreateEndpointRequest
	GetVpcId() *string
	SetVswitchId(v string) *CreateEndpointRequest
	GetVswitchId() *string
}

type CreateEndpointRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// inst-ivrq92qhbyrg4jctih
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// endpoint-1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// VPC
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// vpc-j6co2fcdsl1q0gnuc3ym3
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// example:
	//
	// vsw-j6cmr00qjyrft6jo2mg7g
	VswitchId *string `json:"VswitchId,omitempty" xml:"VswitchId,omitempty"`
}

func (s CreateEndpointRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateEndpointRequest) GoString() string {
	return s.String()
}

func (s *CreateEndpointRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateEndpointRequest) GetName() *string {
	return s.Name
}

func (s *CreateEndpointRequest) GetType() *string {
	return s.Type
}

func (s *CreateEndpointRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *CreateEndpointRequest) GetVswitchId() *string {
	return s.VswitchId
}

func (s *CreateEndpointRequest) SetInstanceId(v string) *CreateEndpointRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateEndpointRequest) SetName(v string) *CreateEndpointRequest {
	s.Name = &v
	return s
}

func (s *CreateEndpointRequest) SetType(v string) *CreateEndpointRequest {
	s.Type = &v
	return s
}

func (s *CreateEndpointRequest) SetVpcId(v string) *CreateEndpointRequest {
	s.VpcId = &v
	return s
}

func (s *CreateEndpointRequest) SetVswitchId(v string) *CreateEndpointRequest {
	s.VswitchId = &v
	return s
}

func (s *CreateEndpointRequest) Validate() error {
	return dara.Validate(s)
}
