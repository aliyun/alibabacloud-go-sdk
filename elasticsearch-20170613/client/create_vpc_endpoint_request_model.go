// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVpcEndpointRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateVpcEndpointRequest
	GetClientToken() *string
	SetServiceId(v string) *CreateVpcEndpointRequest
	GetServiceId() *string
	SetZoneId(v string) *CreateVpcEndpointRequest
	GetZoneId() *string
	SetDryRun(v bool) *CreateVpcEndpointRequest
	GetDryRun() *bool
}

type CreateVpcEndpointRequest struct {
	// The returned result details.
	//
	// example:
	//
	// 5FFD9ED4-C2EC-4E89-B22B-1ACB6FE1D***
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// example:
	//
	// epsrv-hp3xdsq46ael67lo****
	ServiceId *string `json:"serviceId,omitempty" xml:"serviceId,omitempty"`
	// example:
	//
	// cn-hangzhou-i
	ZoneId *string `json:"zoneId,omitempty" xml:"zoneId,omitempty"`
	// The ID of the user endpoint service associated with the endpoint.
	//
	// example:
	//
	// false
	DryRun *bool `json:"dryRun,omitempty" xml:"dryRun,omitempty"`
}

func (s CreateVpcEndpointRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateVpcEndpointRequest) GoString() string {
	return s.String()
}

func (s *CreateVpcEndpointRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateVpcEndpointRequest) GetServiceId() *string {
	return s.ServiceId
}

func (s *CreateVpcEndpointRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *CreateVpcEndpointRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *CreateVpcEndpointRequest) SetClientToken(v string) *CreateVpcEndpointRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateVpcEndpointRequest) SetServiceId(v string) *CreateVpcEndpointRequest {
	s.ServiceId = &v
	return s
}

func (s *CreateVpcEndpointRequest) SetZoneId(v string) *CreateVpcEndpointRequest {
	s.ZoneId = &v
	return s
}

func (s *CreateVpcEndpointRequest) SetDryRun(v bool) *CreateVpcEndpointRequest {
	s.DryRun = &v
	return s
}

func (s *CreateVpcEndpointRequest) Validate() error {
	return dara.Validate(s)
}
