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
	// 	- **true**
	//
	// 	- **false**
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	AutoAllocatedEnabled *bool `json:"AutoAllocatedEnabled,omitempty" xml:"AutoAllocatedEnabled,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// example:
	//
	// 0c593ea1-3bea-11e9-b96b-88e9fe637760
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform a dry run. Valid values:
	//
	// 	- **true**: performs a dry run. The system checks the required parameters, request syntax, and limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// 	- **false**: performs a dry run and sends the request. If the request passes the dry run, an HTTP 2xx status code is returned and the operation is performed. This is the default value.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ID of the region where the service resource is deployed.
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
	// The endpoint service ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// epsrv-hp3vpx8yqxblby3i****
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// The zone ID of the service resource.
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
