// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVpcEndpointServiceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddressIpVersion(v string) *CreateVpcEndpointServiceRequest
	GetAddressIpVersion() *string
	SetAutoAcceptEnabled(v bool) *CreateVpcEndpointServiceRequest
	GetAutoAcceptEnabled() *bool
	SetClientToken(v string) *CreateVpcEndpointServiceRequest
	GetClientToken() *string
	SetDryRun(v bool) *CreateVpcEndpointServiceRequest
	GetDryRun() *bool
	SetPayer(v string) *CreateVpcEndpointServiceRequest
	GetPayer() *string
	SetRegionId(v string) *CreateVpcEndpointServiceRequest
	GetRegionId() *string
	SetResource(v []*CreateVpcEndpointServiceRequestResource) *CreateVpcEndpointServiceRequest
	GetResource() []*CreateVpcEndpointServiceRequestResource
	SetResourceGroupId(v string) *CreateVpcEndpointServiceRequest
	GetResourceGroupId() *string
	SetServiceDescription(v string) *CreateVpcEndpointServiceRequest
	GetServiceDescription() *string
	SetServiceResourceType(v string) *CreateVpcEndpointServiceRequest
	GetServiceResourceType() *string
	SetServiceSupportIPv6(v bool) *CreateVpcEndpointServiceRequest
	GetServiceSupportIPv6() *bool
	SetTag(v []*CreateVpcEndpointServiceRequestTag) *CreateVpcEndpointServiceRequest
	GetTag() []*CreateVpcEndpointServiceRequestTag
	SetZoneAffinityEnabled(v bool) *CreateVpcEndpointServiceRequest
	GetZoneAffinityEnabled() *bool
}

type CreateVpcEndpointServiceRequest struct {
	// The protocol. Valid values:
	//
	// 	- **IPv4*	- (default)
	//
	// 	- **DualStack**
	//
	// >  You can set the protocol to DualStack only for endpoint services whose backend resource type is NLB. An endpoint service supports dual-stack only if its backend resources support dual-stack.
	//
	// example:
	//
	// IPv4
	AddressIpVersion *string `json:"AddressIpVersion,omitempty" xml:"AddressIpVersion,omitempty"`
	// Specifies whether to automatically accept endpoint connection requests. Valid values:
	//
	// 	- **true**
	//
	// 	- **false*	- (default)
	//
	// example:
	//
	// false
	AutoAcceptEnabled *bool `json:"AutoAcceptEnabled,omitempty" xml:"AutoAcceptEnabled,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// example:
	//
	// 0c593ea1-3bea-11e9-b96b-88e9fe637760
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform only a dry run, without performing the actual request.
	//
	// 	- **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// 	- **false*	- (default): performs a dry run and performs the actual request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The payer. Valid values:
	//
	// 	- **Endpoint**: service consumer
	//
	// 	- **EndpointService**: service provider
	//
	// example:
	//
	// Endpoint
	Payer *string `json:"Payer,omitempty" xml:"Payer,omitempty"`
	// The region ID of the endpoint service.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/120468.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-huhehaote
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The service resources of the endpoint service. You can create at most 10 resources. After the resource is created, you can continue to add service resources to the endpoint.
	Resource []*CreateVpcEndpointServiceRequestResource `json:"Resource,omitempty" xml:"Resource,omitempty" type:"Repeated"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfmy*****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The description of the endpoint service.
	//
	// example:
	//
	// This is my EndpointService.
	ServiceDescription *string `json:"ServiceDescription,omitempty" xml:"ServiceDescription,omitempty"`
	// The type of the service resource. Valid values:
	//
	// 	- **slb**: Classic Load Balancer (CLB) instance
	//
	// 	- **alb**: Application Load Balancer (ALB) instance
	//
	// 	- **nlb**: Network Load Balancer (NLB) instance
	//
	// >  You cannot access TCP/SSL listeners configured for NLB instances.
	//
	// example:
	//
	// slb
	ServiceResourceType *string `json:"ServiceResourceType,omitempty" xml:"ServiceResourceType,omitempty"`
	// Deprecated
	//
	// Specifies whether to enable IPv6 for the endpoint service. Valid values:
	//
	// 	- **true**
	//
	// 	- **false*	- (default)
	//
	// example:
	//
	// false
	ServiceSupportIPv6 *bool `json:"ServiceSupportIPv6,omitempty" xml:"ServiceSupportIPv6,omitempty"`
	// The tags to add to the resource.
	Tag []*CreateVpcEndpointServiceRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// Specifies whether to first resolve the domain name of the nearest endpoint that is associated with the endpoint service. Valid values:
	//
	// 	- **true**
	//
	// 	- **false*	- (default)
	//
	// example:
	//
	// false
	ZoneAffinityEnabled *bool `json:"ZoneAffinityEnabled,omitempty" xml:"ZoneAffinityEnabled,omitempty"`
}

func (s CreateVpcEndpointServiceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateVpcEndpointServiceRequest) GoString() string {
	return s.String()
}

func (s *CreateVpcEndpointServiceRequest) GetAddressIpVersion() *string {
	return s.AddressIpVersion
}

func (s *CreateVpcEndpointServiceRequest) GetAutoAcceptEnabled() *bool {
	return s.AutoAcceptEnabled
}

func (s *CreateVpcEndpointServiceRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateVpcEndpointServiceRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *CreateVpcEndpointServiceRequest) GetPayer() *string {
	return s.Payer
}

func (s *CreateVpcEndpointServiceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateVpcEndpointServiceRequest) GetResource() []*CreateVpcEndpointServiceRequestResource {
	return s.Resource
}

func (s *CreateVpcEndpointServiceRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateVpcEndpointServiceRequest) GetServiceDescription() *string {
	return s.ServiceDescription
}

func (s *CreateVpcEndpointServiceRequest) GetServiceResourceType() *string {
	return s.ServiceResourceType
}

func (s *CreateVpcEndpointServiceRequest) GetServiceSupportIPv6() *bool {
	return s.ServiceSupportIPv6
}

func (s *CreateVpcEndpointServiceRequest) GetTag() []*CreateVpcEndpointServiceRequestTag {
	return s.Tag
}

func (s *CreateVpcEndpointServiceRequest) GetZoneAffinityEnabled() *bool {
	return s.ZoneAffinityEnabled
}

func (s *CreateVpcEndpointServiceRequest) SetAddressIpVersion(v string) *CreateVpcEndpointServiceRequest {
	s.AddressIpVersion = &v
	return s
}

func (s *CreateVpcEndpointServiceRequest) SetAutoAcceptEnabled(v bool) *CreateVpcEndpointServiceRequest {
	s.AutoAcceptEnabled = &v
	return s
}

func (s *CreateVpcEndpointServiceRequest) SetClientToken(v string) *CreateVpcEndpointServiceRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateVpcEndpointServiceRequest) SetDryRun(v bool) *CreateVpcEndpointServiceRequest {
	s.DryRun = &v
	return s
}

func (s *CreateVpcEndpointServiceRequest) SetPayer(v string) *CreateVpcEndpointServiceRequest {
	s.Payer = &v
	return s
}

func (s *CreateVpcEndpointServiceRequest) SetRegionId(v string) *CreateVpcEndpointServiceRequest {
	s.RegionId = &v
	return s
}

func (s *CreateVpcEndpointServiceRequest) SetResource(v []*CreateVpcEndpointServiceRequestResource) *CreateVpcEndpointServiceRequest {
	s.Resource = v
	return s
}

func (s *CreateVpcEndpointServiceRequest) SetResourceGroupId(v string) *CreateVpcEndpointServiceRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateVpcEndpointServiceRequest) SetServiceDescription(v string) *CreateVpcEndpointServiceRequest {
	s.ServiceDescription = &v
	return s
}

func (s *CreateVpcEndpointServiceRequest) SetServiceResourceType(v string) *CreateVpcEndpointServiceRequest {
	s.ServiceResourceType = &v
	return s
}

func (s *CreateVpcEndpointServiceRequest) SetServiceSupportIPv6(v bool) *CreateVpcEndpointServiceRequest {
	s.ServiceSupportIPv6 = &v
	return s
}

func (s *CreateVpcEndpointServiceRequest) SetTag(v []*CreateVpcEndpointServiceRequestTag) *CreateVpcEndpointServiceRequest {
	s.Tag = v
	return s
}

func (s *CreateVpcEndpointServiceRequest) SetZoneAffinityEnabled(v bool) *CreateVpcEndpointServiceRequest {
	s.ZoneAffinityEnabled = &v
	return s
}

func (s *CreateVpcEndpointServiceRequest) Validate() error {
	if s.Resource != nil {
		for _, item := range s.Resource {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateVpcEndpointServiceRequestResource struct {
	// The ID of the service resource that is added to the endpoint service.
	//
	// example:
	//
	// lb-hp32z1wp5peaoox2q****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The type of the service resource that is added to the endpoint service. You can add up to 20 service resources to the endpoint service. Valid values:
	//
	// 	- **slb**: CLB instance
	//
	// 	- **alb**: ALB instance
	//
	// 	- **nlb**: NLB instance
	//
	// >  In regions where PrivateLink is supported, CLB instances deployed in virtual private clouds (VPCs) can serve as the service resources of the endpoint service. You cannot access TCP/SSL listeners configured for NLB instances.
	//
	// example:
	//
	// slb
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The zone ID of the cluster.
	//
	// example:
	//
	// cn-huhehaote-a
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s CreateVpcEndpointServiceRequestResource) String() string {
	return dara.Prettify(s)
}

func (s CreateVpcEndpointServiceRequestResource) GoString() string {
	return s.String()
}

func (s *CreateVpcEndpointServiceRequestResource) GetResourceId() *string {
	return s.ResourceId
}

func (s *CreateVpcEndpointServiceRequestResource) GetResourceType() *string {
	return s.ResourceType
}

func (s *CreateVpcEndpointServiceRequestResource) GetZoneId() *string {
	return s.ZoneId
}

func (s *CreateVpcEndpointServiceRequestResource) SetResourceId(v string) *CreateVpcEndpointServiceRequestResource {
	s.ResourceId = &v
	return s
}

func (s *CreateVpcEndpointServiceRequestResource) SetResourceType(v string) *CreateVpcEndpointServiceRequestResource {
	s.ResourceType = &v
	return s
}

func (s *CreateVpcEndpointServiceRequestResource) SetZoneId(v string) *CreateVpcEndpointServiceRequestResource {
	s.ZoneId = &v
	return s
}

func (s *CreateVpcEndpointServiceRequestResource) Validate() error {
	return dara.Validate(s)
}

type CreateVpcEndpointServiceRequestTag struct {
	// The key of the tag to add to the resource. You can specify up to 20 tag keys. The tag key cannot be an empty string.
	//
	// The tag key can be up to 64 characters in length and cannot contain `http://` or `https://`. The tag key cannot start with `aliyun` or `acs:`.
	//
	// example:
	//
	// env
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag to add to the resource. You can specify up to 20 tag values. The tag value can be an empty string.
	//
	// The tag value can be up to 128 characters in length and cannot contain `http://` or `https://`. The tag value cannot start with `aliyun` or `acs:`.
	//
	// example:
	//
	// prod
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateVpcEndpointServiceRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateVpcEndpointServiceRequestTag) GoString() string {
	return s.String()
}

func (s *CreateVpcEndpointServiceRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateVpcEndpointServiceRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateVpcEndpointServiceRequestTag) SetKey(v string) *CreateVpcEndpointServiceRequestTag {
	s.Key = &v
	return s
}

func (s *CreateVpcEndpointServiceRequestTag) SetValue(v string) *CreateVpcEndpointServiceRequestTag {
	s.Value = &v
	return s
}

func (s *CreateVpcEndpointServiceRequestTag) Validate() error {
	return dara.Validate(s)
}
