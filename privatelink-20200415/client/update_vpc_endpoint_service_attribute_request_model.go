// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateVpcEndpointServiceAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddSupportedRegionSet(v []*string) *UpdateVpcEndpointServiceAttributeRequest
	GetAddSupportedRegionSet() []*string
	SetAddressIpVersion(v string) *UpdateVpcEndpointServiceAttributeRequest
	GetAddressIpVersion() *string
	SetAutoAcceptEnabled(v bool) *UpdateVpcEndpointServiceAttributeRequest
	GetAutoAcceptEnabled() *bool
	SetClientToken(v string) *UpdateVpcEndpointServiceAttributeRequest
	GetClientToken() *string
	SetConnectBandwidth(v int32) *UpdateVpcEndpointServiceAttributeRequest
	GetConnectBandwidth() *int32
	SetDeleteSupportedRegionSet(v []*string) *UpdateVpcEndpointServiceAttributeRequest
	GetDeleteSupportedRegionSet() []*string
	SetDryRun(v bool) *UpdateVpcEndpointServiceAttributeRequest
	GetDryRun() *bool
	SetRegionId(v string) *UpdateVpcEndpointServiceAttributeRequest
	GetRegionId() *string
	SetServiceDescription(v string) *UpdateVpcEndpointServiceAttributeRequest
	GetServiceDescription() *string
	SetServiceId(v string) *UpdateVpcEndpointServiceAttributeRequest
	GetServiceId() *string
	SetServiceSupportIPv6(v bool) *UpdateVpcEndpointServiceAttributeRequest
	GetServiceSupportIPv6() *bool
	SetZoneAffinityEnabled(v bool) *UpdateVpcEndpointServiceAttributeRequest
	GetZoneAffinityEnabled() *bool
}

type UpdateVpcEndpointServiceAttributeRequest struct {
	// The list of remote regions to add for the endpoint service.
	AddSupportedRegionSet []*string `json:"AddSupportedRegionSet,omitempty" xml:"AddSupportedRegionSet,omitempty" type:"Repeated"`
	// The protocol version. Valid values:
	//
	// - **IPv4**: IPv4.
	//
	// - **DualStack**: dual-stack.
	//
	// > Currently, only endpoint services whose backend resource type is NLB or GWLB support setting the IP address protocol to DualStack.
	//
	// example:
	//
	// IPv4
	AddressIpVersion *string `json:"AddressIpVersion,omitempty" xml:"AddressIpVersion,omitempty"`
	// Specifies whether to automatically accept endpoint connections. Valid values:
	//
	// - **true**: automatically accepts endpoint connections.
	//
	// - **false**: does not automatically accept endpoint connections.
	//
	// example:
	//
	// false
	AutoAcceptEnabled *bool `json:"AutoAcceptEnabled,omitempty" xml:"AutoAcceptEnabled,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The **ClientToken*	- parameter supports only ASCII characters.
	//
	// example:
	//
	// 0c593ea1-3bea-11e9-b96b-88e9fe637760
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The default bandwidth limit. Default value: **3072**. Unit: Mbit/s.
	//
	// Valid values: **100*	- to **10240**.
	//
	// > Settings for the default bandwidth limit are supported when the service resource is a Classic Load Balancer (CLB) instance or an Application Load Balancer (ALB) instance. When the service resource is a Network Load Balancer (NLB) instance, the connection bandwidth cannot be configured.
	//
	// example:
	//
	// 3072
	ConnectBandwidth *int32 `json:"ConnectBandwidth,omitempty" xml:"ConnectBandwidth,omitempty"`
	// The list of remote regions to remove from the endpoint service.
	DeleteSupportedRegionSet []*string `json:"DeleteSupportedRegionSet,omitempty" xml:"DeleteSupportedRegionSet,omitempty" type:"Repeated"`
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// - **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error code is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// - **false*	- (default): performs a dry run and performs the actual request. If the request passes the dry run, an HTTP 2xx status code is returned and the resource attributes are modified.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The region ID of the endpoint service.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/120468.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The description of the endpoint service.
	//
	// example:
	//
	// This is my EndpointService.
	ServiceDescription *string `json:"ServiceDescription,omitempty" xml:"ServiceDescription,omitempty"`
	// The ID of the endpoint service.
	//
	// This parameter is required.
	//
	// example:
	//
	// epsrv-hp3vpx8yqxblby3i****
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// Deprecated
	//
	// Specifies whether the endpoint service supports IPv6. Valid values:
	//
	// - **true**: yes.
	//
	// - **false*	- (default): no.
	//
	// example:
	//
	// false
	ServiceSupportIPv6 *bool `json:"ServiceSupportIPv6,omitempty" xml:"ServiceSupportIPv6,omitempty"`
	// Specifies whether to resolve the domain name of the endpoint that is connected to the endpoint service to the nearest access point. Valid values:
	//
	// - **true*	- (default): yes.
	//
	// - **false**: no.
	//
	// example:
	//
	// true
	ZoneAffinityEnabled *bool `json:"ZoneAffinityEnabled,omitempty" xml:"ZoneAffinityEnabled,omitempty"`
}

func (s UpdateVpcEndpointServiceAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateVpcEndpointServiceAttributeRequest) GoString() string {
	return s.String()
}

func (s *UpdateVpcEndpointServiceAttributeRequest) GetAddSupportedRegionSet() []*string {
	return s.AddSupportedRegionSet
}

func (s *UpdateVpcEndpointServiceAttributeRequest) GetAddressIpVersion() *string {
	return s.AddressIpVersion
}

func (s *UpdateVpcEndpointServiceAttributeRequest) GetAutoAcceptEnabled() *bool {
	return s.AutoAcceptEnabled
}

func (s *UpdateVpcEndpointServiceAttributeRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateVpcEndpointServiceAttributeRequest) GetConnectBandwidth() *int32 {
	return s.ConnectBandwidth
}

func (s *UpdateVpcEndpointServiceAttributeRequest) GetDeleteSupportedRegionSet() []*string {
	return s.DeleteSupportedRegionSet
}

func (s *UpdateVpcEndpointServiceAttributeRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *UpdateVpcEndpointServiceAttributeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateVpcEndpointServiceAttributeRequest) GetServiceDescription() *string {
	return s.ServiceDescription
}

func (s *UpdateVpcEndpointServiceAttributeRequest) GetServiceId() *string {
	return s.ServiceId
}

func (s *UpdateVpcEndpointServiceAttributeRequest) GetServiceSupportIPv6() *bool {
	return s.ServiceSupportIPv6
}

func (s *UpdateVpcEndpointServiceAttributeRequest) GetZoneAffinityEnabled() *bool {
	return s.ZoneAffinityEnabled
}

func (s *UpdateVpcEndpointServiceAttributeRequest) SetAddSupportedRegionSet(v []*string) *UpdateVpcEndpointServiceAttributeRequest {
	s.AddSupportedRegionSet = v
	return s
}

func (s *UpdateVpcEndpointServiceAttributeRequest) SetAddressIpVersion(v string) *UpdateVpcEndpointServiceAttributeRequest {
	s.AddressIpVersion = &v
	return s
}

func (s *UpdateVpcEndpointServiceAttributeRequest) SetAutoAcceptEnabled(v bool) *UpdateVpcEndpointServiceAttributeRequest {
	s.AutoAcceptEnabled = &v
	return s
}

func (s *UpdateVpcEndpointServiceAttributeRequest) SetClientToken(v string) *UpdateVpcEndpointServiceAttributeRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateVpcEndpointServiceAttributeRequest) SetConnectBandwidth(v int32) *UpdateVpcEndpointServiceAttributeRequest {
	s.ConnectBandwidth = &v
	return s
}

func (s *UpdateVpcEndpointServiceAttributeRequest) SetDeleteSupportedRegionSet(v []*string) *UpdateVpcEndpointServiceAttributeRequest {
	s.DeleteSupportedRegionSet = v
	return s
}

func (s *UpdateVpcEndpointServiceAttributeRequest) SetDryRun(v bool) *UpdateVpcEndpointServiceAttributeRequest {
	s.DryRun = &v
	return s
}

func (s *UpdateVpcEndpointServiceAttributeRequest) SetRegionId(v string) *UpdateVpcEndpointServiceAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateVpcEndpointServiceAttributeRequest) SetServiceDescription(v string) *UpdateVpcEndpointServiceAttributeRequest {
	s.ServiceDescription = &v
	return s
}

func (s *UpdateVpcEndpointServiceAttributeRequest) SetServiceId(v string) *UpdateVpcEndpointServiceAttributeRequest {
	s.ServiceId = &v
	return s
}

func (s *UpdateVpcEndpointServiceAttributeRequest) SetServiceSupportIPv6(v bool) *UpdateVpcEndpointServiceAttributeRequest {
	s.ServiceSupportIPv6 = &v
	return s
}

func (s *UpdateVpcEndpointServiceAttributeRequest) SetZoneAffinityEnabled(v bool) *UpdateVpcEndpointServiceAttributeRequest {
	s.ZoneAffinityEnabled = &v
	return s
}

func (s *UpdateVpcEndpointServiceAttributeRequest) Validate() error {
	return dara.Validate(s)
}
