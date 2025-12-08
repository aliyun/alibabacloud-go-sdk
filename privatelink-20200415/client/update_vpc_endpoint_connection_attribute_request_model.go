// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateVpcEndpointConnectionAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBandwidth(v int32) *UpdateVpcEndpointConnectionAttributeRequest
	GetBandwidth() *int32
	SetClientToken(v string) *UpdateVpcEndpointConnectionAttributeRequest
	GetClientToken() *string
	SetDryRun(v bool) *UpdateVpcEndpointConnectionAttributeRequest
	GetDryRun() *bool
	SetEndpointId(v string) *UpdateVpcEndpointConnectionAttributeRequest
	GetEndpointId() *string
	SetRegionId(v string) *UpdateVpcEndpointConnectionAttributeRequest
	GetRegionId() *string
	SetServiceId(v string) *UpdateVpcEndpointConnectionAttributeRequest
	GetServiceId() *string
	SetTrafficControlMode(v string) *UpdateVpcEndpointConnectionAttributeRequest
	GetTrafficControlMode() *string
}

type UpdateVpcEndpointConnectionAttributeRequest struct {
	// The bandwidth of the endpoint connection that you want to modify. Unit: Mbit/s. Valid values: **3072*	- to **10240**.
	//
	// >  The bandwidth of an endpoint connection is in the range of **100*	- to **10,240*	- Mbit/s. The default bandwidth is **3,072*	- Mbit/s. When the endpoint is connected to the endpoint service, the default bandwidth is the minimum bandwidth. In this case, the connection bandwidth range is **3,072*	- to **10,240*	- Mbit/s. If Classic Load Balancer (CLB) instances or Application Load Balancer (ALB) instances are specified as service resources, you can modify the endpoint connection bandwidth based on your business requirements. This parameter is invalid if Network Load Balancer (NLB) instances are specified as service resources.
	//
	// example:
	//
	// 1000
	Bandwidth *int32 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// example:
	//
	// 0c593ea1-3bea-11e9-b96b-88e9fe637760
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// 	- **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// 	- **false*	- (default): performs a dry run and performs the actual request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The endpoint ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ep-hp33b2e43fays7s8****
	EndpointId *string `json:"EndpointId,omitempty" xml:"EndpointId,omitempty"`
	// The region ID of the endpoint connection whose bandwidth you want to modify. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/120468.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// eu-west-1
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The endpoint service ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// epsrv-hp3vpx8yqxblby3i****
	ServiceId          *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	TrafficControlMode *string `json:"TrafficControlMode,omitempty" xml:"TrafficControlMode,omitempty"`
}

func (s UpdateVpcEndpointConnectionAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateVpcEndpointConnectionAttributeRequest) GoString() string {
	return s.String()
}

func (s *UpdateVpcEndpointConnectionAttributeRequest) GetBandwidth() *int32 {
	return s.Bandwidth
}

func (s *UpdateVpcEndpointConnectionAttributeRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateVpcEndpointConnectionAttributeRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *UpdateVpcEndpointConnectionAttributeRequest) GetEndpointId() *string {
	return s.EndpointId
}

func (s *UpdateVpcEndpointConnectionAttributeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateVpcEndpointConnectionAttributeRequest) GetServiceId() *string {
	return s.ServiceId
}

func (s *UpdateVpcEndpointConnectionAttributeRequest) GetTrafficControlMode() *string {
	return s.TrafficControlMode
}

func (s *UpdateVpcEndpointConnectionAttributeRequest) SetBandwidth(v int32) *UpdateVpcEndpointConnectionAttributeRequest {
	s.Bandwidth = &v
	return s
}

func (s *UpdateVpcEndpointConnectionAttributeRequest) SetClientToken(v string) *UpdateVpcEndpointConnectionAttributeRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateVpcEndpointConnectionAttributeRequest) SetDryRun(v bool) *UpdateVpcEndpointConnectionAttributeRequest {
	s.DryRun = &v
	return s
}

func (s *UpdateVpcEndpointConnectionAttributeRequest) SetEndpointId(v string) *UpdateVpcEndpointConnectionAttributeRequest {
	s.EndpointId = &v
	return s
}

func (s *UpdateVpcEndpointConnectionAttributeRequest) SetRegionId(v string) *UpdateVpcEndpointConnectionAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateVpcEndpointConnectionAttributeRequest) SetServiceId(v string) *UpdateVpcEndpointConnectionAttributeRequest {
	s.ServiceId = &v
	return s
}

func (s *UpdateVpcEndpointConnectionAttributeRequest) SetTrafficControlMode(v string) *UpdateVpcEndpointConnectionAttributeRequest {
	s.TrafficControlMode = &v
	return s
}

func (s *UpdateVpcEndpointConnectionAttributeRequest) Validate() error {
	return dara.Validate(s)
}
