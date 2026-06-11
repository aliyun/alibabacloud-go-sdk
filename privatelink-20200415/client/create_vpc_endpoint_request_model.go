// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVpcEndpointRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddressIpVersion(v string) *CreateVpcEndpointRequest
	GetAddressIpVersion() *string
	SetClientToken(v string) *CreateVpcEndpointRequest
	GetClientToken() *string
	SetCrossRegionBandwidth(v int32) *CreateVpcEndpointRequest
	GetCrossRegionBandwidth() *int32
	SetDryRun(v bool) *CreateVpcEndpointRequest
	GetDryRun() *bool
	SetEndpointDescription(v string) *CreateVpcEndpointRequest
	GetEndpointDescription() *string
	SetEndpointName(v string) *CreateVpcEndpointRequest
	GetEndpointName() *string
	SetEndpointType(v string) *CreateVpcEndpointRequest
	GetEndpointType() *string
	SetPolicyDocument(v string) *CreateVpcEndpointRequest
	GetPolicyDocument() *string
	SetProtectedEnabled(v bool) *CreateVpcEndpointRequest
	GetProtectedEnabled() *bool
	SetRegionId(v string) *CreateVpcEndpointRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CreateVpcEndpointRequest
	GetResourceGroupId() *string
	SetSecurityGroupId(v []*string) *CreateVpcEndpointRequest
	GetSecurityGroupId() []*string
	SetServiceId(v string) *CreateVpcEndpointRequest
	GetServiceId() *string
	SetServiceName(v string) *CreateVpcEndpointRequest
	GetServiceName() *string
	SetServiceRegionId(v string) *CreateVpcEndpointRequest
	GetServiceRegionId() *string
	SetTag(v []*CreateVpcEndpointRequestTag) *CreateVpcEndpointRequest
	GetTag() []*CreateVpcEndpointRequestTag
	SetVpcId(v string) *CreateVpcEndpointRequest
	GetVpcId() *string
	SetZone(v []*CreateVpcEndpointRequestZone) *CreateVpcEndpointRequest
	GetZone() []*CreateVpcEndpointRequestZone
	SetZoneAffinityEnabled(v bool) *CreateVpcEndpointRequest
	GetZoneAffinityEnabled() *bool
	SetZonePrivateIpAddressCount(v int64) *CreateVpcEndpointRequest
	GetZonePrivateIpAddressCount() *int64
}

type CreateVpcEndpointRequest struct {
	// The IP version of the endpoint. Valid values:
	//
	// - **IPv4**: IPv4 (default).
	//
	// - **DualStack**: dual-stack.
	//
	// > To use the dual-stack feature, make sure that the associated endpoint service and the VPC in which the endpoint is created support the dual-stack feature.
	//
	// example:
	//
	// IPv4
	AddressIpVersion *string `json:"AddressIpVersion,omitempty" xml:"AddressIpVersion,omitempty"`
	// A client-generated token to ensure the idempotence of the request.
	//
	// You must generate a unique value for this token. The token can contain only ASCII characters.
	//
	// example:
	//
	// 0c593ea1-3bea-11e9-b96b-88e9fe637760
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The bandwidth for a cross-region connection, in Mbps. This parameter applies only when the endpoint and endpoint service are in different regions. Valid values:
	//
	// - **Default**: 1000 for cross-region connections within the Chinese mainland. In all other cases, the value is 100.
	//
	// - **Minimum value**: 100.
	//
	// - **Maximum value**: subject to your account\\"s quota. For more information, see [Quotas and limits](https://help.aliyun.com/zh/privatelink/quotas-and-limits?spm=a2c4g.11174283.help-menu-search-120462.d_0).
	//
	// > To use this parameter, make sure that you are creating a cross-region endpoint.
	//
	// example:
	//
	// 1000
	CrossRegionBandwidth *int32 `json:"CrossRegionBandwidth,omitempty" xml:"CrossRegionBandwidth,omitempty"`
	// Specifies whether to perform a dry run. Valid values:
	//
	// - **true**: Performs a dry run to check the request\\"s validity without committing the action. The system checks for required parameters, request format, and service limits. If the check passes, the `DryRunOperation` error code is returned. If it fails, an error message is returned.
	//
	// - **false*	- (default): Sends the request. If the request is valid, the operation is performed and a 2xx HTTP status code is returned.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The description of the endpoint.
	//
	// The description must be 2 to 256 characters in length and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// This is my Endpoint.
	EndpointDescription *string `json:"EndpointDescription,omitempty" xml:"EndpointDescription,omitempty"`
	// The name of the endpoint.
	//
	// The name must be 2 to 128 characters long, start with a letter or a Chinese character, and can contain digits, hyphens (-), and underscores (_).
	//
	// example:
	//
	// test
	EndpointName *string `json:"EndpointName,omitempty" xml:"EndpointName,omitempty"`
	// The type of the endpoint. Valid values:
	//
	// - **Interface**: an interface endpoint. You can add Application Load Balancer (ALB), Classic Load Balancer (CLB), and Network Load Balancer (NLB) instances as service resources.
	//
	// - **Reverse**: a reverse endpoint. You can add a VPC NAT Gateway as a service resource.
	//
	// - **GatewayLoadBalancer**: a Gateway Load Balancer endpoint. You can add a Gateway Load Balancer (GWLB) as a service resource.
	//
	// > Services that support reverse endpoints are provided exclusively by Alibaba Cloud and its partners. You cannot create them by default. To request access, contact your account manager.
	//
	// example:
	//
	// Interface
	EndpointType *string `json:"EndpointType,omitempty" xml:"EndpointType,omitempty"`
	// The Resource Access Management (RAM) policy. For more information about the policy syntax, see [Basic elements of a policy](https://help.aliyun.com/document_detail/93738.html).
	//
	// example:
	//
	// {
	//
	//   "Version": "1",
	//
	//   "Statement": [
	//
	//     {
	//
	//       "Effect": "Allow",
	//
	//       "Action": [
	//
	//         "oss:List*",
	//
	//         "oss:PutObject",
	//
	//         "oss:GetObject"
	//
	//       ],
	//
	//       "Resource": [
	//
	//         "acs:oss:oss-*:*:pvl-policy-test/policy-test.txt"
	//
	//       ],
	//
	//       "Principal": {
	//
	//         "RAM": [
	//
	//           "acs:ram::14199xxxxxx:*"
	//
	//         ]
	//
	//       }
	//
	//     }
	//
	//   ]
	//
	// }
	PolicyDocument *string `json:"PolicyDocument,omitempty" xml:"PolicyDocument,omitempty"`
	// Specifies whether to enable managed protection. This parameter is effective only for requests made with a Security Token Service (STS) token. Valid values:
	//
	// - **true**: enables managed protection. After you enable managed protection, only the user who creates the endpoint can modify or delete it by using an STS token.
	//
	// - **false*	- (default): disables managed protection.
	//
	// example:
	//
	// false
	ProtectedEnabled *bool `json:"ProtectedEnabled,omitempty" xml:"ProtectedEnabled,omitempty"`
	// The ID of the region in which to create the endpoint.
	//
	// You can obtain the region ID by calling the [DescribeRegions](https://help.aliyun.com/document_detail/120468.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-huhehaote
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfmy*****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The IDs of security groups to associate with the endpoint ENI.
	//
	// example:
	//
	// sg-hp33bw6ynvm2yb0e****
	SecurityGroupId []*string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty" type:"Repeated"`
	// The ID of the associated endpoint service.
	//
	// example:
	//
	// epsrv-hp3xdsq46ael67lo****
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// The name of the associated endpoint service.
	//
	// example:
	//
	// com.aliyuncs.privatelink.cn-huhehaote.epsrv-hp3vpx8yqxblby3i****
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	// The ID of the region where the endpoint service is deployed. Defaults to the endpoint\\"s region.
	//
	// example:
	//
	// cn-huhehaote
	ServiceRegionId *string `json:"ServiceRegionId,omitempty" xml:"ServiceRegionId,omitempty"`
	// The list of tags.
	Tag []*CreateVpcEndpointRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The ID of the Virtual Private Cloud (VPC) where the endpoint will be created.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpc-hp356stwkxg3fn2xe****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The list of zones for the endpoint.
	Zone []*CreateVpcEndpointRequestZone `json:"Zone,omitempty" xml:"Zone,omitempty" type:"Repeated"`
	// Specifies whether to enable zone affinity. If enabled, requests are routed to the endpoint in the same zone as the client. Valid values:
	//
	// - **true**: enables zone affinity.
	//
	// - **false*	- (default): disables zone affinity.
	//
	// example:
	//
	// false
	ZoneAffinityEnabled *bool `json:"ZoneAffinityEnabled,omitempty" xml:"ZoneAffinityEnabled,omitempty"`
	// The number of private IP addresses for the endpoint\\"s elastic network interface (ENI) in each zone. The value must be **1**.
	//
	// example:
	//
	// 1
	ZonePrivateIpAddressCount *int64 `json:"ZonePrivateIpAddressCount,omitempty" xml:"ZonePrivateIpAddressCount,omitempty"`
}

func (s CreateVpcEndpointRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateVpcEndpointRequest) GoString() string {
	return s.String()
}

func (s *CreateVpcEndpointRequest) GetAddressIpVersion() *string {
	return s.AddressIpVersion
}

func (s *CreateVpcEndpointRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateVpcEndpointRequest) GetCrossRegionBandwidth() *int32 {
	return s.CrossRegionBandwidth
}

func (s *CreateVpcEndpointRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *CreateVpcEndpointRequest) GetEndpointDescription() *string {
	return s.EndpointDescription
}

func (s *CreateVpcEndpointRequest) GetEndpointName() *string {
	return s.EndpointName
}

func (s *CreateVpcEndpointRequest) GetEndpointType() *string {
	return s.EndpointType
}

func (s *CreateVpcEndpointRequest) GetPolicyDocument() *string {
	return s.PolicyDocument
}

func (s *CreateVpcEndpointRequest) GetProtectedEnabled() *bool {
	return s.ProtectedEnabled
}

func (s *CreateVpcEndpointRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateVpcEndpointRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateVpcEndpointRequest) GetSecurityGroupId() []*string {
	return s.SecurityGroupId
}

func (s *CreateVpcEndpointRequest) GetServiceId() *string {
	return s.ServiceId
}

func (s *CreateVpcEndpointRequest) GetServiceName() *string {
	return s.ServiceName
}

func (s *CreateVpcEndpointRequest) GetServiceRegionId() *string {
	return s.ServiceRegionId
}

func (s *CreateVpcEndpointRequest) GetTag() []*CreateVpcEndpointRequestTag {
	return s.Tag
}

func (s *CreateVpcEndpointRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *CreateVpcEndpointRequest) GetZone() []*CreateVpcEndpointRequestZone {
	return s.Zone
}

func (s *CreateVpcEndpointRequest) GetZoneAffinityEnabled() *bool {
	return s.ZoneAffinityEnabled
}

func (s *CreateVpcEndpointRequest) GetZonePrivateIpAddressCount() *int64 {
	return s.ZonePrivateIpAddressCount
}

func (s *CreateVpcEndpointRequest) SetAddressIpVersion(v string) *CreateVpcEndpointRequest {
	s.AddressIpVersion = &v
	return s
}

func (s *CreateVpcEndpointRequest) SetClientToken(v string) *CreateVpcEndpointRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateVpcEndpointRequest) SetCrossRegionBandwidth(v int32) *CreateVpcEndpointRequest {
	s.CrossRegionBandwidth = &v
	return s
}

func (s *CreateVpcEndpointRequest) SetDryRun(v bool) *CreateVpcEndpointRequest {
	s.DryRun = &v
	return s
}

func (s *CreateVpcEndpointRequest) SetEndpointDescription(v string) *CreateVpcEndpointRequest {
	s.EndpointDescription = &v
	return s
}

func (s *CreateVpcEndpointRequest) SetEndpointName(v string) *CreateVpcEndpointRequest {
	s.EndpointName = &v
	return s
}

func (s *CreateVpcEndpointRequest) SetEndpointType(v string) *CreateVpcEndpointRequest {
	s.EndpointType = &v
	return s
}

func (s *CreateVpcEndpointRequest) SetPolicyDocument(v string) *CreateVpcEndpointRequest {
	s.PolicyDocument = &v
	return s
}

func (s *CreateVpcEndpointRequest) SetProtectedEnabled(v bool) *CreateVpcEndpointRequest {
	s.ProtectedEnabled = &v
	return s
}

func (s *CreateVpcEndpointRequest) SetRegionId(v string) *CreateVpcEndpointRequest {
	s.RegionId = &v
	return s
}

func (s *CreateVpcEndpointRequest) SetResourceGroupId(v string) *CreateVpcEndpointRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateVpcEndpointRequest) SetSecurityGroupId(v []*string) *CreateVpcEndpointRequest {
	s.SecurityGroupId = v
	return s
}

func (s *CreateVpcEndpointRequest) SetServiceId(v string) *CreateVpcEndpointRequest {
	s.ServiceId = &v
	return s
}

func (s *CreateVpcEndpointRequest) SetServiceName(v string) *CreateVpcEndpointRequest {
	s.ServiceName = &v
	return s
}

func (s *CreateVpcEndpointRequest) SetServiceRegionId(v string) *CreateVpcEndpointRequest {
	s.ServiceRegionId = &v
	return s
}

func (s *CreateVpcEndpointRequest) SetTag(v []*CreateVpcEndpointRequestTag) *CreateVpcEndpointRequest {
	s.Tag = v
	return s
}

func (s *CreateVpcEndpointRequest) SetVpcId(v string) *CreateVpcEndpointRequest {
	s.VpcId = &v
	return s
}

func (s *CreateVpcEndpointRequest) SetZone(v []*CreateVpcEndpointRequestZone) *CreateVpcEndpointRequest {
	s.Zone = v
	return s
}

func (s *CreateVpcEndpointRequest) SetZoneAffinityEnabled(v bool) *CreateVpcEndpointRequest {
	s.ZoneAffinityEnabled = &v
	return s
}

func (s *CreateVpcEndpointRequest) SetZonePrivateIpAddressCount(v int64) *CreateVpcEndpointRequest {
	s.ZonePrivateIpAddressCount = &v
	return s
}

func (s *CreateVpcEndpointRequest) Validate() error {
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Zone != nil {
		for _, item := range s.Zone {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateVpcEndpointRequestTag struct {
	// The tag key.
	//
	// example:
	//
	// env
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// prod
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateVpcEndpointRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateVpcEndpointRequestTag) GoString() string {
	return s.String()
}

func (s *CreateVpcEndpointRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateVpcEndpointRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateVpcEndpointRequestTag) SetKey(v string) *CreateVpcEndpointRequestTag {
	s.Key = &v
	return s
}

func (s *CreateVpcEndpointRequestTag) SetValue(v string) *CreateVpcEndpointRequestTag {
	s.Value = &v
	return s
}

func (s *CreateVpcEndpointRequestTag) Validate() error {
	return dara.Validate(s)
}

type CreateVpcEndpointRequestZone struct {
	// The IPv6 address of the endpoint ENI in the specified zone.
	//
	// > This parameter is valid only when `AddressIpVersion` is set to `DualStack`.
	//
	// example:
	//
	// 2408:4005:34d:****:a58b:62a3:6b55:****
	Ipv6Address *string `json:"Ipv6Address,omitempty" xml:"Ipv6Address,omitempty"`
	// The ID of the vSwitch in the zone where the endpoint ENI will be created.
	//
	// You can specify up to 10 vSwitch IDs.
	//
	// example:
	//
	// vsw-hp3uf6045ljdhd5zr****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The ID of the zone for the endpoint.
	//
	// You can specify up to 10 zone IDs.
	//
	// example:
	//
	// cn-huhehaote-b
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	// The IPv4 address of the endpoint ENI in the specified zone.
	//
	// example:
	//
	// 192.168.XX.XX
	Ip *string `json:"ip,omitempty" xml:"ip,omitempty"`
}

func (s CreateVpcEndpointRequestZone) String() string {
	return dara.Prettify(s)
}

func (s CreateVpcEndpointRequestZone) GoString() string {
	return s.String()
}

func (s *CreateVpcEndpointRequestZone) GetIpv6Address() *string {
	return s.Ipv6Address
}

func (s *CreateVpcEndpointRequestZone) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *CreateVpcEndpointRequestZone) GetZoneId() *string {
	return s.ZoneId
}

func (s *CreateVpcEndpointRequestZone) GetIp() *string {
	return s.Ip
}

func (s *CreateVpcEndpointRequestZone) SetIpv6Address(v string) *CreateVpcEndpointRequestZone {
	s.Ipv6Address = &v
	return s
}

func (s *CreateVpcEndpointRequestZone) SetVSwitchId(v string) *CreateVpcEndpointRequestZone {
	s.VSwitchId = &v
	return s
}

func (s *CreateVpcEndpointRequestZone) SetZoneId(v string) *CreateVpcEndpointRequestZone {
	s.ZoneId = &v
	return s
}

func (s *CreateVpcEndpointRequestZone) SetIp(v string) *CreateVpcEndpointRequestZone {
	s.Ip = &v
	return s
}

func (s *CreateVpcEndpointRequestZone) Validate() error {
	return dara.Validate(s)
}
