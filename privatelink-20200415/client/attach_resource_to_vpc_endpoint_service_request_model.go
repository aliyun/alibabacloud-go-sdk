// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachResourceToVpcEndpointServiceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *AttachResourceToVpcEndpointServiceRequest
	GetClientToken() *string
	SetDryRun(v bool) *AttachResourceToVpcEndpointServiceRequest
	GetDryRun() *bool
	SetRegionId(v string) *AttachResourceToVpcEndpointServiceRequest
	GetRegionId() *string
	SetResourceId(v string) *AttachResourceToVpcEndpointServiceRequest
	GetResourceId() *string
	SetResourceType(v string) *AttachResourceToVpcEndpointServiceRequest
	GetResourceType() *string
	SetServiceId(v string) *AttachResourceToVpcEndpointServiceRequest
	GetServiceId() *string
	SetZoneId(v string) *AttachResourceToVpcEndpointServiceRequest
	GetZoneId() *string
}

type AttachResourceToVpcEndpointServiceRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The **ClientToken*	- can contain only ASCII characters.
	//
	// example:
	//
	// 0c593ea1-3bea-11e9-b96b-88e9fe637760
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform a dry run. Valid values:
	//
	// - **true**: Performs a dry run to check the request for potential issues, including required parameters, request format, and service limits. The system does not add the service resource. If the check fails, an error message is returned. If the check passes, the `DryRunOperation` error code is returned.
	//
	// - **false*	- (default): Sends a normal request. After the request passes the check, an HTTP 2xx status code is returned and the service resource is added.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ID of the region where the endpoint service is deployed.
	//
	// Call the [DescribeRegions](https://help.aliyun.com/document_detail/448570.html) operation to obtain the region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the service resource.
	//
	// This parameter is required.
	//
	// example:
	//
	// lb-hp32z1wp5peaoox2q****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The type of the service resource. Valid values:
	//
	// - **slb**: Classic Load Balancer (CLB).
	//
	// - **alb**: Application Load Balancer (ALB).
	//
	// - **nlb**: Network Load Balancer (NLB).
	//
	// - **gwlb**: Gateway Load Balancer (GWLB).
	//
	// This parameter is required.
	//
	// example:
	//
	// slb
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The endpoint service to which you want to add the service resource.
	//
	// This parameter is required.
	//
	// example:
	//
	// epsrv-hp3vpx8yqxblby3i****
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// The zone where the service resource is located. This parameter is required if the service resource is an ALB, NLB, or GWLB instance.
	//
	// example:
	//
	// cn-hangzhou-j
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s AttachResourceToVpcEndpointServiceRequest) String() string {
	return dara.Prettify(s)
}

func (s AttachResourceToVpcEndpointServiceRequest) GoString() string {
	return s.String()
}

func (s *AttachResourceToVpcEndpointServiceRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *AttachResourceToVpcEndpointServiceRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *AttachResourceToVpcEndpointServiceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AttachResourceToVpcEndpointServiceRequest) GetResourceId() *string {
	return s.ResourceId
}

func (s *AttachResourceToVpcEndpointServiceRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *AttachResourceToVpcEndpointServiceRequest) GetServiceId() *string {
	return s.ServiceId
}

func (s *AttachResourceToVpcEndpointServiceRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *AttachResourceToVpcEndpointServiceRequest) SetClientToken(v string) *AttachResourceToVpcEndpointServiceRequest {
	s.ClientToken = &v
	return s
}

func (s *AttachResourceToVpcEndpointServiceRequest) SetDryRun(v bool) *AttachResourceToVpcEndpointServiceRequest {
	s.DryRun = &v
	return s
}

func (s *AttachResourceToVpcEndpointServiceRequest) SetRegionId(v string) *AttachResourceToVpcEndpointServiceRequest {
	s.RegionId = &v
	return s
}

func (s *AttachResourceToVpcEndpointServiceRequest) SetResourceId(v string) *AttachResourceToVpcEndpointServiceRequest {
	s.ResourceId = &v
	return s
}

func (s *AttachResourceToVpcEndpointServiceRequest) SetResourceType(v string) *AttachResourceToVpcEndpointServiceRequest {
	s.ResourceType = &v
	return s
}

func (s *AttachResourceToVpcEndpointServiceRequest) SetServiceId(v string) *AttachResourceToVpcEndpointServiceRequest {
	s.ServiceId = &v
	return s
}

func (s *AttachResourceToVpcEndpointServiceRequest) SetZoneId(v string) *AttachResourceToVpcEndpointServiceRequest {
	s.ZoneId = &v
	return s
}

func (s *AttachResourceToVpcEndpointServiceRequest) Validate() error {
	return dara.Validate(s)
}
