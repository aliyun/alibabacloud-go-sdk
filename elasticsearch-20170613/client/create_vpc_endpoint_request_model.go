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
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must ensure that it is unique among different requests. The token can contain a maximum of 64 ASCII characters.
	//
	// example:
	//
	// 5FFD9ED4-C2EC-4E89-B22B-1ACB6FE1D***
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the endpoint service that you created. The endpoint on the service VPC side is associated with this service ID.
	//
	// example:
	//
	// epsrv-hp3xdsq46ael67lo****
	ServiceId *string `json:"serviceId,omitempty" xml:"serviceId,omitempty"`
	// The zone ID of the endpoint service that you created.
	//
	// example:
	//
	// cn-hangzhou-i
	ZoneId *string `json:"zoneId,omitempty" xml:"zoneId,omitempty"`
	// Specifies whether to only perform a dry run on the request. Valid values:
	//
	// - true: sends a check request and does not create an endpoint node. The check items include whether the AccessKey pair is valid, the authorization of the Resource Access Management (RAM) user, and whether the required parameters are specified.
	//
	// - false (default): immediately triggers the creation of an endpoint node.
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
