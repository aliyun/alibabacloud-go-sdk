// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateVpcEndpointServiceResourceAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoAllocatedEnabled(v bool) *UpdateVpcEndpointServiceResourceAttributeRequest
	GetAutoAllocatedEnabled() *bool
	SetClientToken(v string) *UpdateVpcEndpointServiceResourceAttributeRequest
	GetClientToken() *string
	SetDryRun(v bool) *UpdateVpcEndpointServiceResourceAttributeRequest
	GetDryRun() *bool
	SetRegionId(v string) *UpdateVpcEndpointServiceResourceAttributeRequest
	GetRegionId() *string
	SetResourceId(v string) *UpdateVpcEndpointServiceResourceAttributeRequest
	GetResourceId() *string
	SetServiceId(v string) *UpdateVpcEndpointServiceResourceAttributeRequest
	GetServiceId() *string
	SetZoneId(v string) *UpdateVpcEndpointServiceResourceAttributeRequest
	GetZoneId() *string
}

type UpdateVpcEndpointServiceResourceAttributeRequest struct {
	// Specifies whether to enable automatic resource allocation. Valid values:
	//
	// - **true**: Enables automatic resource allocation.
	//
	// - **false**: Disables automatic resource allocation.
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	AutoAllocatedEnabled *bool `json:"AutoAllocatedEnabled,omitempty" xml:"AutoAllocatedEnabled,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// Generate a parameter value from your client to ensure that the value is unique among different requests. **ClientToken*	- can contain only ASCII characters.
	//
	// example:
	//
	// 0c593ea1-3bea-11e9-b96b-88e9fe637760
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform a dry run. Valid values:
	//
	// - **true**: sends a check request without modifying the properties of the service resource. The system checks the required parameters, request format, and service limits. If the request fails the check, an error is returned. If the request passes the check, the `DryRunOperation` error code is returned.
	//
	// - **false*	- (default): sends a normal request. After the request passes the check, an HTTP 2xx status code is returned and the operation is performed.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ID of the region where you want to modify the service resource.
	//
	// Call the [DescribeRegions](https://help.aliyun.com/document_detail/120468.html) operation to obtain the region ID.
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
	// The ID of the endpoint service.
	//
	// This parameter is required.
	//
	// example:
	//
	// epsrv-hp3vpx8yqxblby3i****
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// The zone where the service resource is located. This parameter is required if the service resource is an Application Load Balancer (ALB), a Network Load Balancer (NLB), or a Gateway Load Balancer (GWLB).
	//
	// example:
	//
	// cn-hangzhou-i
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s UpdateVpcEndpointServiceResourceAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateVpcEndpointServiceResourceAttributeRequest) GoString() string {
	return s.String()
}

func (s *UpdateVpcEndpointServiceResourceAttributeRequest) GetAutoAllocatedEnabled() *bool {
	return s.AutoAllocatedEnabled
}

func (s *UpdateVpcEndpointServiceResourceAttributeRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateVpcEndpointServiceResourceAttributeRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *UpdateVpcEndpointServiceResourceAttributeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateVpcEndpointServiceResourceAttributeRequest) GetResourceId() *string {
	return s.ResourceId
}

func (s *UpdateVpcEndpointServiceResourceAttributeRequest) GetServiceId() *string {
	return s.ServiceId
}

func (s *UpdateVpcEndpointServiceResourceAttributeRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *UpdateVpcEndpointServiceResourceAttributeRequest) SetAutoAllocatedEnabled(v bool) *UpdateVpcEndpointServiceResourceAttributeRequest {
	s.AutoAllocatedEnabled = &v
	return s
}

func (s *UpdateVpcEndpointServiceResourceAttributeRequest) SetClientToken(v string) *UpdateVpcEndpointServiceResourceAttributeRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateVpcEndpointServiceResourceAttributeRequest) SetDryRun(v bool) *UpdateVpcEndpointServiceResourceAttributeRequest {
	s.DryRun = &v
	return s
}

func (s *UpdateVpcEndpointServiceResourceAttributeRequest) SetRegionId(v string) *UpdateVpcEndpointServiceResourceAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateVpcEndpointServiceResourceAttributeRequest) SetResourceId(v string) *UpdateVpcEndpointServiceResourceAttributeRequest {
	s.ResourceId = &v
	return s
}

func (s *UpdateVpcEndpointServiceResourceAttributeRequest) SetServiceId(v string) *UpdateVpcEndpointServiceResourceAttributeRequest {
	s.ServiceId = &v
	return s
}

func (s *UpdateVpcEndpointServiceResourceAttributeRequest) SetZoneId(v string) *UpdateVpcEndpointServiceResourceAttributeRequest {
	s.ZoneId = &v
	return s
}

func (s *UpdateVpcEndpointServiceResourceAttributeRequest) Validate() error {
	return dara.Validate(s)
}
