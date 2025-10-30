// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateVpcEndpointZoneConnectionResourceAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *UpdateVpcEndpointZoneConnectionResourceAttributeRequest
	GetClientToken() *string
	SetDryRun(v bool) *UpdateVpcEndpointZoneConnectionResourceAttributeRequest
	GetDryRun() *bool
	SetEndpointId(v string) *UpdateVpcEndpointZoneConnectionResourceAttributeRequest
	GetEndpointId() *string
	SetRegionId(v string) *UpdateVpcEndpointZoneConnectionResourceAttributeRequest
	GetRegionId() *string
	SetResourceAllocateMode(v string) *UpdateVpcEndpointZoneConnectionResourceAttributeRequest
	GetResourceAllocateMode() *string
	SetResourceId(v string) *UpdateVpcEndpointZoneConnectionResourceAttributeRequest
	GetResourceId() *string
	SetResourceReplaceMode(v string) *UpdateVpcEndpointZoneConnectionResourceAttributeRequest
	GetResourceReplaceMode() *string
	SetResourceType(v string) *UpdateVpcEndpointZoneConnectionResourceAttributeRequest
	GetResourceType() *string
	SetServiceId(v string) *UpdateVpcEndpointZoneConnectionResourceAttributeRequest
	GetServiceId() *string
	SetZoneId(v string) *UpdateVpcEndpointZoneConnectionResourceAttributeRequest
	GetZoneId() *string
}

type UpdateVpcEndpointZoneConnectionResourceAttributeRequest struct {
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
	// The region ID of the endpoint connection whose bandwidth you want to modify.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/120468.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource allocation mode. You can change the resource allocation mode only if the endpoint connection is in the **Disconnected*	- state. Valid values:
	//
	// 	- **Auto**: automatically and randomly allocates service resources. In this mode, the specified service resource is deleted.
	//
	// 	- **Manual**: manually allocates service resources. If you set the value to Manual, you must also specify the **ResourceId*	- and **ResourceType*	- parameters.
	//
	// example:
	//
	// Auto
	ResourceAllocateMode *string `json:"ResourceAllocateMode,omitempty" xml:"ResourceAllocateMode,omitempty"`
	// The ID of the service resource that you want to manually allocate or migrate in the zone where the endpoint connection is deployed.
	//
	// > If **ResourceAllocateMode*	- is set to **Mannual**, or **ResourceReplaceMode*	- is set, this parameter is required.
	//
	// example:
	//
	// lb-hp32z1wp5peaoox2q****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The migration mode of the service resource. Valid values:
	//
	// 	- **Graceful**: smooth migration. Service resources in the zone are smoothly migrated.
	//
	// 	- **Force**: forced migration. Service resources in the zone are forcefully migrated.
	//
	// >  You need to specify this parameter only if you want to migrate service resources and the endpoint connection is in the **Connected*	- state. If you specify this parameter, you must also specify the **ResourceId*	- and **ResourceType*	- parameters.
	//
	// example:
	//
	// Graceful
	ResourceReplaceMode *string `json:"ResourceReplaceMode,omitempty" xml:"ResourceReplaceMode,omitempty"`
	// The type of the service resource. Valid values:
	//
	// 	- **slb**: a CLB instance that supports PrivateLink. In addition, the CLB instance is deployed in a VPC.
	//
	// 	- **alb**: an Application Load Balancer (ALB) instance that supports PrivateLink. In addition, the ALB instance is deployed in a VPC.
	//
	// > If **ResourceAllocateMode*	- is set to **Mannual**, or **ResourceReplaceMode*	- is set, this parameter is required.
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
	// The zone ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou-b
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s UpdateVpcEndpointZoneConnectionResourceAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateVpcEndpointZoneConnectionResourceAttributeRequest) GoString() string {
	return s.String()
}

func (s *UpdateVpcEndpointZoneConnectionResourceAttributeRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateVpcEndpointZoneConnectionResourceAttributeRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *UpdateVpcEndpointZoneConnectionResourceAttributeRequest) GetEndpointId() *string {
	return s.EndpointId
}

func (s *UpdateVpcEndpointZoneConnectionResourceAttributeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateVpcEndpointZoneConnectionResourceAttributeRequest) GetResourceAllocateMode() *string {
	return s.ResourceAllocateMode
}

func (s *UpdateVpcEndpointZoneConnectionResourceAttributeRequest) GetResourceId() *string {
	return s.ResourceId
}

func (s *UpdateVpcEndpointZoneConnectionResourceAttributeRequest) GetResourceReplaceMode() *string {
	return s.ResourceReplaceMode
}

func (s *UpdateVpcEndpointZoneConnectionResourceAttributeRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *UpdateVpcEndpointZoneConnectionResourceAttributeRequest) GetServiceId() *string {
	return s.ServiceId
}

func (s *UpdateVpcEndpointZoneConnectionResourceAttributeRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *UpdateVpcEndpointZoneConnectionResourceAttributeRequest) SetClientToken(v string) *UpdateVpcEndpointZoneConnectionResourceAttributeRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateVpcEndpointZoneConnectionResourceAttributeRequest) SetDryRun(v bool) *UpdateVpcEndpointZoneConnectionResourceAttributeRequest {
	s.DryRun = &v
	return s
}

func (s *UpdateVpcEndpointZoneConnectionResourceAttributeRequest) SetEndpointId(v string) *UpdateVpcEndpointZoneConnectionResourceAttributeRequest {
	s.EndpointId = &v
	return s
}

func (s *UpdateVpcEndpointZoneConnectionResourceAttributeRequest) SetRegionId(v string) *UpdateVpcEndpointZoneConnectionResourceAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateVpcEndpointZoneConnectionResourceAttributeRequest) SetResourceAllocateMode(v string) *UpdateVpcEndpointZoneConnectionResourceAttributeRequest {
	s.ResourceAllocateMode = &v
	return s
}

func (s *UpdateVpcEndpointZoneConnectionResourceAttributeRequest) SetResourceId(v string) *UpdateVpcEndpointZoneConnectionResourceAttributeRequest {
	s.ResourceId = &v
	return s
}

func (s *UpdateVpcEndpointZoneConnectionResourceAttributeRequest) SetResourceReplaceMode(v string) *UpdateVpcEndpointZoneConnectionResourceAttributeRequest {
	s.ResourceReplaceMode = &v
	return s
}

func (s *UpdateVpcEndpointZoneConnectionResourceAttributeRequest) SetResourceType(v string) *UpdateVpcEndpointZoneConnectionResourceAttributeRequest {
	s.ResourceType = &v
	return s
}

func (s *UpdateVpcEndpointZoneConnectionResourceAttributeRequest) SetServiceId(v string) *UpdateVpcEndpointZoneConnectionResourceAttributeRequest {
	s.ServiceId = &v
	return s
}

func (s *UpdateVpcEndpointZoneConnectionResourceAttributeRequest) SetZoneId(v string) *UpdateVpcEndpointZoneConnectionResourceAttributeRequest {
	s.ZoneId = &v
	return s
}

func (s *UpdateVpcEndpointZoneConnectionResourceAttributeRequest) Validate() error {
	return dara.Validate(s)
}
