// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateVpcEndpointServiceAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddressIpVersion(v string) *UpdateVpcEndpointServiceAttributeRequest
	GetAddressIpVersion() *string
	SetAutoAcceptEnabled(v bool) *UpdateVpcEndpointServiceAttributeRequest
	GetAutoAcceptEnabled() *bool
	SetClientToken(v string) *UpdateVpcEndpointServiceAttributeRequest
	GetClientToken() *string
	SetConnectBandwidth(v int32) *UpdateVpcEndpointServiceAttributeRequest
	GetConnectBandwidth() *int32
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
	// The protocol. Valid values:
	//
	// 	- **IPv4**
	//
	// 	- **DualStack**
	//
	// >  You can set the protocol to DualStack only for endpoint services whose backend resource type is NLB.
	//
	// example:
	//
	// IPv4
	AddressIpVersion *string `json:"AddressIpVersion,omitempty" xml:"AddressIpVersion,omitempty"`
	// Specifies whether to automatically accept endpoint connection requests. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
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
	// The default maximum bandwidth of the endpoint connection. Unit: Mbit/s. Default value: **3072**.
	//
	// Valid values: **100*	- to **10240**.
	//
	// >  You can specify this parameter only if you specify Classic Load Balancer (CLB) instances or Application Load Balancer (ALB) instances as service resources.
	//
	// example:
	//
	// 200
	ConnectBandwidth *int32 `json:"ConnectBandwidth,omitempty" xml:"ConnectBandwidth,omitempty"`
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
	// The endpoint service ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// epsrv-hp3vpx8yqxblby3i****
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// Deprecated
	//
	// Specifies whether to enable IPv6. Valid values:
	//
	// 	- **true**
	//
	// 	- **false*	- (default)
	//
	// example:
	//
	// false
	ServiceSupportIPv6 *bool `json:"ServiceSupportIPv6,omitempty" xml:"ServiceSupportIPv6,omitempty"`
	// Specifies whether to first resolve the domain name of the nearest endpoint that is associated with the endpoint service. Valid values:
	//
	// 	- **true*	- (default)
	//
	// 	- **false**
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
