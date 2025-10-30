// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachResourceFromVpcEndpointServiceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DetachResourceFromVpcEndpointServiceRequest
	GetClientToken() *string
	SetDryRun(v bool) *DetachResourceFromVpcEndpointServiceRequest
	GetDryRun() *bool
	SetRegionId(v string) *DetachResourceFromVpcEndpointServiceRequest
	GetRegionId() *string
	SetResourceId(v string) *DetachResourceFromVpcEndpointServiceRequest
	GetResourceId() *string
	SetResourceType(v string) *DetachResourceFromVpcEndpointServiceRequest
	GetResourceType() *string
	SetServiceId(v string) *DetachResourceFromVpcEndpointServiceRequest
	GetServiceId() *string
	SetZoneId(v string) *DetachResourceFromVpcEndpointServiceRequest
	GetZoneId() *string
}

type DetachResourceFromVpcEndpointServiceRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate a value, but you must make sure that the value is unique among different requests. The token can contain only ASCII characters.
	//
	// example:
	//
	// 0c593ea1-3bea-11e9-b96b-88e9fe637760
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// 	- **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error code is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// 	- **false*	- (default): performs a dry run and performs the actual request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The region ID of the endpoint.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/120468.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The service resource ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// lb-hp32z1wp5peaoox2q****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The type of the service resource. Valid values:
	//
	// 	- **slb**: a Classic Load Balancer (CLB) instance that supports PrivateLink. In addition, the CLB instance is deployed in a virtual private cloud (VPC).
	//
	// 	- **alb**: an Application Load Balancer (ALB) instance that supports PrivateLink. In addition, the ALB instance is deployed in a VPC.
	//
	// example:
	//
	// slb
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The endpoint service ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// epsrv-hp3vpx8yqxblby3i****
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// The ID of the zone that you want to remove.
	//
	// example:
	//
	// cn-hangzhou-c
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DetachResourceFromVpcEndpointServiceRequest) String() string {
	return dara.Prettify(s)
}

func (s DetachResourceFromVpcEndpointServiceRequest) GoString() string {
	return s.String()
}

func (s *DetachResourceFromVpcEndpointServiceRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DetachResourceFromVpcEndpointServiceRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *DetachResourceFromVpcEndpointServiceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DetachResourceFromVpcEndpointServiceRequest) GetResourceId() *string {
	return s.ResourceId
}

func (s *DetachResourceFromVpcEndpointServiceRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *DetachResourceFromVpcEndpointServiceRequest) GetServiceId() *string {
	return s.ServiceId
}

func (s *DetachResourceFromVpcEndpointServiceRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *DetachResourceFromVpcEndpointServiceRequest) SetClientToken(v string) *DetachResourceFromVpcEndpointServiceRequest {
	s.ClientToken = &v
	return s
}

func (s *DetachResourceFromVpcEndpointServiceRequest) SetDryRun(v bool) *DetachResourceFromVpcEndpointServiceRequest {
	s.DryRun = &v
	return s
}

func (s *DetachResourceFromVpcEndpointServiceRequest) SetRegionId(v string) *DetachResourceFromVpcEndpointServiceRequest {
	s.RegionId = &v
	return s
}

func (s *DetachResourceFromVpcEndpointServiceRequest) SetResourceId(v string) *DetachResourceFromVpcEndpointServiceRequest {
	s.ResourceId = &v
	return s
}

func (s *DetachResourceFromVpcEndpointServiceRequest) SetResourceType(v string) *DetachResourceFromVpcEndpointServiceRequest {
	s.ResourceType = &v
	return s
}

func (s *DetachResourceFromVpcEndpointServiceRequest) SetServiceId(v string) *DetachResourceFromVpcEndpointServiceRequest {
	s.ServiceId = &v
	return s
}

func (s *DetachResourceFromVpcEndpointServiceRequest) SetZoneId(v string) *DetachResourceFromVpcEndpointServiceRequest {
	s.ZoneId = &v
	return s
}

func (s *DetachResourceFromVpcEndpointServiceRequest) Validate() error {
	return dara.Validate(s)
}
