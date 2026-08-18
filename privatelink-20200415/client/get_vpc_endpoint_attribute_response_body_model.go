// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVpcEndpointAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAddressIpVersion(v string) *GetVpcEndpointAttributeResponseBody
	GetAddressIpVersion() *string
	SetBandwidth(v int32) *GetVpcEndpointAttributeResponseBody
	GetBandwidth() *int32
	SetConnectionStatus(v string) *GetVpcEndpointAttributeResponseBody
	GetConnectionStatus() *string
	SetCreateTime(v string) *GetVpcEndpointAttributeResponseBody
	GetCreateTime() *string
	SetCrossRegionBandwidth(v int32) *GetVpcEndpointAttributeResponseBody
	GetCrossRegionBandwidth() *int32
	SetEndpointBusinessStatus(v string) *GetVpcEndpointAttributeResponseBody
	GetEndpointBusinessStatus() *string
	SetEndpointDescription(v string) *GetVpcEndpointAttributeResponseBody
	GetEndpointDescription() *string
	SetEndpointDomain(v string) *GetVpcEndpointAttributeResponseBody
	GetEndpointDomain() *string
	SetEndpointId(v string) *GetVpcEndpointAttributeResponseBody
	GetEndpointId() *string
	SetEndpointName(v string) *GetVpcEndpointAttributeResponseBody
	GetEndpointName() *string
	SetEndpointStatus(v string) *GetVpcEndpointAttributeResponseBody
	GetEndpointStatus() *string
	SetEndpointType(v string) *GetVpcEndpointAttributeResponseBody
	GetEndpointType() *string
	SetPayer(v string) *GetVpcEndpointAttributeResponseBody
	GetPayer() *string
	SetPolicyDocument(v string) *GetVpcEndpointAttributeResponseBody
	GetPolicyDocument() *string
	SetProtectedEnabled(v bool) *GetVpcEndpointAttributeResponseBody
	GetProtectedEnabled() *bool
	SetRegionId(v string) *GetVpcEndpointAttributeResponseBody
	GetRegionId() *string
	SetRequestId(v string) *GetVpcEndpointAttributeResponseBody
	GetRequestId() *string
	SetResourceGroupId(v string) *GetVpcEndpointAttributeResponseBody
	GetResourceGroupId() *string
	SetResourceOwner(v bool) *GetVpcEndpointAttributeResponseBody
	GetResourceOwner() *bool
	SetServiceId(v string) *GetVpcEndpointAttributeResponseBody
	GetServiceId() *string
	SetServiceName(v string) *GetVpcEndpointAttributeResponseBody
	GetServiceName() *string
	SetServiceRegionId(v string) *GetVpcEndpointAttributeResponseBody
	GetServiceRegionId() *string
	SetVpcId(v string) *GetVpcEndpointAttributeResponseBody
	GetVpcId() *string
	SetZoneAffinityEnabled(v bool) *GetVpcEndpointAttributeResponseBody
	GetZoneAffinityEnabled() *bool
	SetZonePrivateIpAddressCount(v int64) *GetVpcEndpointAttributeResponseBody
	GetZonePrivateIpAddressCount() *int64
}

type GetVpcEndpointAttributeResponseBody struct {
	// The protocol version. Valid values:
	//
	// - **IPv4**: IPv4.
	//
	// - **DualStack**: dual-stack.
	//
	// example:
	//
	// IPv4
	AddressIpVersion *string `json:"AddressIpVersion,omitempty" xml:"AddressIpVersion,omitempty"`
	// The bandwidth of the endpoint connection. Unit: Mbps.
	//
	// example:
	//
	// 1024
	Bandwidth *int32 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The endpoint connection status. Valid values:
	//
	// - **Pending**: being modified.
	//
	// - **Connecting**: connecting.
	//
	// - **Connected**: connected.
	//
	// - **Disconnecting**: disconnecting.
	//
	// - **Disconnected**: disconnected.
	//
	// - **Deleting**: being deleted.
	//
	// - **ServiceDeleted**: the corresponding endpoint service has been deleted.
	//
	// example:
	//
	// Connected
	ConnectionStatus *string `json:"ConnectionStatus,omitempty" xml:"ConnectionStatus,omitempty"`
	// The time when the endpoint was created.
	//
	// example:
	//
	// 2021-09-24T18:00:07Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The cross-region bandwidth of the endpoint. Unit: Mbps.
	//
	// example:
	//
	// 1000
	CrossRegionBandwidth *int32 `json:"CrossRegionBandwidth,omitempty" xml:"CrossRegionBandwidth,omitempty"`
	// The business status of the endpoint. Valid values:
	//
	// - **Normal**: Normal.
	//
	// - **FinancialLocked**: Financial lock.
	//
	// example:
	//
	// Normal
	EndpointBusinessStatus *string `json:"EndpointBusinessStatus,omitempty" xml:"EndpointBusinessStatus,omitempty"`
	// The description of the endpoint.
	//
	// example:
	//
	// This is my Endpoint.
	EndpointDescription *string `json:"EndpointDescription,omitempty" xml:"EndpointDescription,omitempty"`
	// The endpoint domain name.
	//
	// example:
	//
	// ep-hp33b2e43fays7s8****.epsrv-hp3xdsq46ael67lo****.cn-huhehaote.privatelink.aliyuncs.com
	EndpointDomain *string `json:"EndpointDomain,omitempty" xml:"EndpointDomain,omitempty"`
	// The endpoint ID.
	//
	// example:
	//
	// ep-hp33b2e43fays7s8****
	EndpointId *string `json:"EndpointId,omitempty" xml:"EndpointId,omitempty"`
	// The name of the endpoint.
	//
	// example:
	//
	// test
	EndpointName *string `json:"EndpointName,omitempty" xml:"EndpointName,omitempty"`
	// The status of the endpoint. Valid values:
	//
	// - **Creating**: being created.
	//
	// - **Active**: available.
	//
	// - **Pending**: being modified.
	//
	// - **Deleting**: being deleted.
	//
	// example:
	//
	// Active
	EndpointStatus *string `json:"EndpointStatus,omitempty" xml:"EndpointStatus,omitempty"`
	// The endpoint type. Valid values:
	//
	// - **Interface**: interface endpoint.
	//
	// - **Reverse**: reverse endpoint.
	//
	// - **GatewayLoadBalancer**: Gateway Load Balancer endpoint (GWLBe).
	//
	// example:
	//
	// Interface
	EndpointType *string `json:"EndpointType,omitempty" xml:"EndpointType,omitempty"`
	// The payer. Valid values:
	//
	// - **Endpoint**: the service consumer.
	//
	// - **EndpointService**: the service provider.
	//
	// example:
	//
	// Endpoint
	Payer *string `json:"Payer,omitempty" xml:"Payer,omitempty"`
	// The RAM access policy. For more information about the policy definition, see [Policy elements](https://help.aliyun.com/document_detail/93738.html).
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
	// Indicates whether managed protection is enabled. This parameter takes effect only when the STS calling method is used. Valid values:
	//
	// **true**: enabled. After managed protection is enabled, only the same user who created the endpoint can modify or delete the endpoint by using STS.
	//
	// **false**: disabled.
	//
	// This parameter is required.
	ProtectedEnabled *bool `json:"ProtectedEnabled,omitempty" xml:"ProtectedEnabled,omitempty"`
	// The region ID of the endpoint.
	//
	// example:
	//
	// cn-huhehaote
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 8D8992C1-6712-423C-BAC5-E5E817484C6B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfmz7nocpei***
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// Indicates whether the endpoint and the endpoint service belong to the same Alibaba Cloud account. Valid values:
	//
	// - **true**: The endpoint and the endpoint service belong to the same account.
	//
	// - **false**: The endpoint and the endpoint service belong to different accounts.
	//
	// example:
	//
	// true
	ResourceOwner *bool `json:"ResourceOwner,omitempty" xml:"ResourceOwner,omitempty"`
	// The ID of the endpoint service with which the endpoint is associated.
	//
	// example:
	//
	// epsrv-hp3vpx8yqxblby3i****
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// The name of the endpoint service with which the endpoint is associated.
	//
	// example:
	//
	// com.aliyuncs.privatelink.cn-huhehaote.epsrv-hp3xdsq46ael67lo****
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	// The region ID of the endpoint service with which the endpoint is associated.
	//
	// example:
	//
	// cn-hangzhou
	ServiceRegionId *string `json:"ServiceRegionId,omitempty" xml:"ServiceRegionId,omitempty"`
	// The ID of the virtual private cloud (VPC) to which the endpoint belongs.
	//
	// example:
	//
	// vpc-fdfhkjafhjvcvdjf****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// Indicates whether zone affinity is enabled for the endpoint domain name to resolve to the connected service. Valid values:
	//
	// - **true**: enabled.
	//
	// - **false**: disabled.
	//
	// example:
	//
	// true
	ZoneAffinityEnabled *bool `json:"ZoneAffinityEnabled,omitempty" xml:"ZoneAffinityEnabled,omitempty"`
	// The number of private IP addresses assigned to the endpoint elastic network interface (ENI) in each zone. The value is fixed to **1**.
	//
	// example:
	//
	// 1
	ZonePrivateIpAddressCount *int64 `json:"ZonePrivateIpAddressCount,omitempty" xml:"ZonePrivateIpAddressCount,omitempty"`
}

func (s GetVpcEndpointAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetVpcEndpointAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *GetVpcEndpointAttributeResponseBody) GetAddressIpVersion() *string {
	return s.AddressIpVersion
}

func (s *GetVpcEndpointAttributeResponseBody) GetBandwidth() *int32 {
	return s.Bandwidth
}

func (s *GetVpcEndpointAttributeResponseBody) GetConnectionStatus() *string {
	return s.ConnectionStatus
}

func (s *GetVpcEndpointAttributeResponseBody) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetVpcEndpointAttributeResponseBody) GetCrossRegionBandwidth() *int32 {
	return s.CrossRegionBandwidth
}

func (s *GetVpcEndpointAttributeResponseBody) GetEndpointBusinessStatus() *string {
	return s.EndpointBusinessStatus
}

func (s *GetVpcEndpointAttributeResponseBody) GetEndpointDescription() *string {
	return s.EndpointDescription
}

func (s *GetVpcEndpointAttributeResponseBody) GetEndpointDomain() *string {
	return s.EndpointDomain
}

func (s *GetVpcEndpointAttributeResponseBody) GetEndpointId() *string {
	return s.EndpointId
}

func (s *GetVpcEndpointAttributeResponseBody) GetEndpointName() *string {
	return s.EndpointName
}

func (s *GetVpcEndpointAttributeResponseBody) GetEndpointStatus() *string {
	return s.EndpointStatus
}

func (s *GetVpcEndpointAttributeResponseBody) GetEndpointType() *string {
	return s.EndpointType
}

func (s *GetVpcEndpointAttributeResponseBody) GetPayer() *string {
	return s.Payer
}

func (s *GetVpcEndpointAttributeResponseBody) GetPolicyDocument() *string {
	return s.PolicyDocument
}

func (s *GetVpcEndpointAttributeResponseBody) GetProtectedEnabled() *bool {
	return s.ProtectedEnabled
}

func (s *GetVpcEndpointAttributeResponseBody) GetRegionId() *string {
	return s.RegionId
}

func (s *GetVpcEndpointAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetVpcEndpointAttributeResponseBody) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *GetVpcEndpointAttributeResponseBody) GetResourceOwner() *bool {
	return s.ResourceOwner
}

func (s *GetVpcEndpointAttributeResponseBody) GetServiceId() *string {
	return s.ServiceId
}

func (s *GetVpcEndpointAttributeResponseBody) GetServiceName() *string {
	return s.ServiceName
}

func (s *GetVpcEndpointAttributeResponseBody) GetServiceRegionId() *string {
	return s.ServiceRegionId
}

func (s *GetVpcEndpointAttributeResponseBody) GetVpcId() *string {
	return s.VpcId
}

func (s *GetVpcEndpointAttributeResponseBody) GetZoneAffinityEnabled() *bool {
	return s.ZoneAffinityEnabled
}

func (s *GetVpcEndpointAttributeResponseBody) GetZonePrivateIpAddressCount() *int64 {
	return s.ZonePrivateIpAddressCount
}

func (s *GetVpcEndpointAttributeResponseBody) SetAddressIpVersion(v string) *GetVpcEndpointAttributeResponseBody {
	s.AddressIpVersion = &v
	return s
}

func (s *GetVpcEndpointAttributeResponseBody) SetBandwidth(v int32) *GetVpcEndpointAttributeResponseBody {
	s.Bandwidth = &v
	return s
}

func (s *GetVpcEndpointAttributeResponseBody) SetConnectionStatus(v string) *GetVpcEndpointAttributeResponseBody {
	s.ConnectionStatus = &v
	return s
}

func (s *GetVpcEndpointAttributeResponseBody) SetCreateTime(v string) *GetVpcEndpointAttributeResponseBody {
	s.CreateTime = &v
	return s
}

func (s *GetVpcEndpointAttributeResponseBody) SetCrossRegionBandwidth(v int32) *GetVpcEndpointAttributeResponseBody {
	s.CrossRegionBandwidth = &v
	return s
}

func (s *GetVpcEndpointAttributeResponseBody) SetEndpointBusinessStatus(v string) *GetVpcEndpointAttributeResponseBody {
	s.EndpointBusinessStatus = &v
	return s
}

func (s *GetVpcEndpointAttributeResponseBody) SetEndpointDescription(v string) *GetVpcEndpointAttributeResponseBody {
	s.EndpointDescription = &v
	return s
}

func (s *GetVpcEndpointAttributeResponseBody) SetEndpointDomain(v string) *GetVpcEndpointAttributeResponseBody {
	s.EndpointDomain = &v
	return s
}

func (s *GetVpcEndpointAttributeResponseBody) SetEndpointId(v string) *GetVpcEndpointAttributeResponseBody {
	s.EndpointId = &v
	return s
}

func (s *GetVpcEndpointAttributeResponseBody) SetEndpointName(v string) *GetVpcEndpointAttributeResponseBody {
	s.EndpointName = &v
	return s
}

func (s *GetVpcEndpointAttributeResponseBody) SetEndpointStatus(v string) *GetVpcEndpointAttributeResponseBody {
	s.EndpointStatus = &v
	return s
}

func (s *GetVpcEndpointAttributeResponseBody) SetEndpointType(v string) *GetVpcEndpointAttributeResponseBody {
	s.EndpointType = &v
	return s
}

func (s *GetVpcEndpointAttributeResponseBody) SetPayer(v string) *GetVpcEndpointAttributeResponseBody {
	s.Payer = &v
	return s
}

func (s *GetVpcEndpointAttributeResponseBody) SetPolicyDocument(v string) *GetVpcEndpointAttributeResponseBody {
	s.PolicyDocument = &v
	return s
}

func (s *GetVpcEndpointAttributeResponseBody) SetProtectedEnabled(v bool) *GetVpcEndpointAttributeResponseBody {
	s.ProtectedEnabled = &v
	return s
}

func (s *GetVpcEndpointAttributeResponseBody) SetRegionId(v string) *GetVpcEndpointAttributeResponseBody {
	s.RegionId = &v
	return s
}

func (s *GetVpcEndpointAttributeResponseBody) SetRequestId(v string) *GetVpcEndpointAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetVpcEndpointAttributeResponseBody) SetResourceGroupId(v string) *GetVpcEndpointAttributeResponseBody {
	s.ResourceGroupId = &v
	return s
}

func (s *GetVpcEndpointAttributeResponseBody) SetResourceOwner(v bool) *GetVpcEndpointAttributeResponseBody {
	s.ResourceOwner = &v
	return s
}

func (s *GetVpcEndpointAttributeResponseBody) SetServiceId(v string) *GetVpcEndpointAttributeResponseBody {
	s.ServiceId = &v
	return s
}

func (s *GetVpcEndpointAttributeResponseBody) SetServiceName(v string) *GetVpcEndpointAttributeResponseBody {
	s.ServiceName = &v
	return s
}

func (s *GetVpcEndpointAttributeResponseBody) SetServiceRegionId(v string) *GetVpcEndpointAttributeResponseBody {
	s.ServiceRegionId = &v
	return s
}

func (s *GetVpcEndpointAttributeResponseBody) SetVpcId(v string) *GetVpcEndpointAttributeResponseBody {
	s.VpcId = &v
	return s
}

func (s *GetVpcEndpointAttributeResponseBody) SetZoneAffinityEnabled(v bool) *GetVpcEndpointAttributeResponseBody {
	s.ZoneAffinityEnabled = &v
	return s
}

func (s *GetVpcEndpointAttributeResponseBody) SetZonePrivateIpAddressCount(v int64) *GetVpcEndpointAttributeResponseBody {
	s.ZonePrivateIpAddressCount = &v
	return s
}

func (s *GetVpcEndpointAttributeResponseBody) Validate() error {
	return dara.Validate(s)
}
