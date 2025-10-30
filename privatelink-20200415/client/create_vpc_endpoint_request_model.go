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
	// The protocol. Valid values:
	//
	// 	- **IPv4*	- (default)
	//
	// 	- **DualStack**
	//
	// >  An endpoint supports dual-stack if its associated endpoint service and VPC both support dual-stack.
	//
	// example:
	//
	// IPv4
	AddressIpVersion *string `json:"AddressIpVersion,omitempty" xml:"AddressIpVersion,omitempty"`
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
	// The description of the endpoint.
	//
	// The description must be 2 to 256 characters in length, and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// This is my Endpoint.
	EndpointDescription *string `json:"EndpointDescription,omitempty" xml:"EndpointDescription,omitempty"`
	// The name of the endpoint.
	//
	// The name must be 2 to 128 characters in length, and can contain digits, underscores (_), and hyphens (-). The name must start with a letter.
	//
	// example:
	//
	// test
	EndpointName *string `json:"EndpointName,omitempty" xml:"EndpointName,omitempty"`
	// The endpoint type. Valid values:
	//
	// 	- **Interface*	- You can specify an Application Load Balancer (ALB) instance, a Classic Load Balancer (CLB) instance, or a Network Load Balancer (NLB) instance.
	//
	// 	- **Reverse*	- You can specify a Virtual Private Cloud (VPC) NAT gateway.
	//
	// >  Services that support reverse endpoints are provided by Alibaba Cloud or Alibaba Cloud partners. To create such a service on your own, contact your account manager.
	//
	// example:
	//
	// Interface
	EndpointType *string `json:"EndpointType,omitempty" xml:"EndpointType,omitempty"`
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
	// Specifies whether to enable user authentication. This parameter is available in Security Token Service (STS) mode. Valid values:
	//
	// 	- **true**: enables user authentication. After user authentication is enabled, only the user who creates the endpoint can modify or delete the endpoint in STS mode.
	//
	// 	- **false*	- (default): disables user authentication.
	//
	// example:
	//
	// false
	ProtectedEnabled *bool `json:"ProtectedEnabled,omitempty" xml:"ProtectedEnabled,omitempty"`
	// The region ID of the endpoint.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/120468.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-huhehaote
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfmy*****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The IDs of security groups that are associated with the endpoint elastic network interface (ENI).
	//
	// example:
	//
	// sg-hp33bw6ynvm2yb0e****
	SecurityGroupId []*string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty" type:"Repeated"`
	// The ID of the endpoint service with which the endpoint is associated.
	//
	// example:
	//
	// epsrv-hp3xdsq46ael67lo****
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// The name of the endpoint service with which the endpoint is associated.
	//
	// example:
	//
	// com.aliyuncs.privatelink.cn-huhehaote.epsrv-hp3vpx8yqxblby3i****
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	// The tags to add to the resource.
	Tag []*CreateVpcEndpointRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The ID of the virtual private cloud (VPC) to which the endpoint belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpc-hp356stwkxg3fn2xe****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The zones where the endpoint is deployed.
	Zone                []*CreateVpcEndpointRequestZone `json:"Zone,omitempty" xml:"Zone,omitempty" type:"Repeated"`
	ZoneAffinityEnabled *bool                           `json:"ZoneAffinityEnabled,omitempty" xml:"ZoneAffinityEnabled,omitempty"`
	// The number of private IP addresses that are assigned to an elastic network interface (ENI) in each zone. Set the value to **1**.
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
	// The key of the tag to add to the resource.
	//
	// example:
	//
	// env
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag to add to the resource.
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
	// The IPv6 address of the zone where the endpoint is deployed.
	//
	// >  You can specify this parameter only if AddressIpVersion is set to DualStack.
	//
	// example:
	//
	// 2408:4005:34d:****:a58b:62a3:6b55:****
	Ipv6Address *string `json:"Ipv6Address,omitempty" xml:"Ipv6Address,omitempty"`
	// The ID of the vSwitch for which you want to create the endpoint elastic network interface (ENI) in the zone. You can specify up to 10 vSwitches.
	//
	// example:
	//
	// vsw-hp3uf6045ljdhd5zr****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The ID of the zone where the endpoint service is deployed.
	//
	// You can specify up to 10 zones.
	//
	// example:
	//
	// cn-huhehaote-b
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	// The IP address of the zone where the endpoint is deployed.
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
