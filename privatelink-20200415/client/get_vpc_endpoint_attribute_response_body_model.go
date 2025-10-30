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
	SetVpcId(v string) *GetVpcEndpointAttributeResponseBody
	GetVpcId() *string
	SetZoneAffinityEnabled(v bool) *GetVpcEndpointAttributeResponseBody
	GetZoneAffinityEnabled() *bool
	SetZonePrivateIpAddressCount(v int64) *GetVpcEndpointAttributeResponseBody
	GetZonePrivateIpAddressCount() *int64
}

type GetVpcEndpointAttributeResponseBody struct {
	// The protocol. Valid values:
	//
	// 	- **IPv4**
	//
	// 	- **DualStack**
	//
	// example:
	//
	// IPv4
	AddressIpVersion *string `json:"AddressIpVersion,omitempty" xml:"AddressIpVersion,omitempty"`
	// The bandwidth of the endpoint connection. Unit: Mbit/s.
	//
	// example:
	//
	// 1024
	Bandwidth *int32 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The state of the endpoint connection. Valid values:
	//
	// 	- **Pending**: The connection is being modified.
	//
	// 	- **Connecting**: The connection is being established.
	//
	// 	- **Connected**: The connection is established.
	//
	// 	- **Disconnecting**: The endpoint is being disconnected from the endpoint service.
	//
	// 	- **Disconnected**: The endpoint is disconnected from the endpoint service.
	//
	// 	- **Deleting**: The connection is being deleted.
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
	// The service state of the endpoint. Valid values:
	//
	// 	- **Normal**: The endpoint runs as expected.
	//
	// 	- **FinancialLocked**: The endpoint is locked due to overdue payments.
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
	// The domain name of the endpoint.
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
	// The state of the endpoint. Valid values:
	//
	// 	- **Creating**: The endpoint is being created.
	//
	// 	- **Active**: The endpoint is available.
	//
	// 	- **Pending**: The endpoint is being modified.
	//
	// 	- **Deleting**: The endpoint is being deleted.
	//
	// example:
	//
	// Active
	EndpointStatus *string `json:"EndpointStatus,omitempty" xml:"EndpointStatus,omitempty"`
	// The type of the endpoint.
	//
	// **Interface*	- is returned. The value indicates the interface endpoint with which the Classic Load Balancer (CLB) instances are associated.
	//
	// example:
	//
	// Interface
	EndpointType *string `json:"EndpointType,omitempty" xml:"EndpointType,omitempty"`
	// The payer. Valid values:
	//
	// 	- **Endpoint**: the service consumer.
	//
	// 	- **EndpointService**: the service provider.
	//
	// example:
	//
	// Endpoint
	Payer *string `json:"Payer,omitempty" xml:"Payer,omitempty"`
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
	// 	- **true**: The endpoint and the endpoint service belong to the same Alibaba Cloud account.
	//
	// 	- **false**: The endpoint and the endpoint service do not belong to the same Alibaba Cloud account.
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
	// The ID of the virtual private cloud (VPC) to which the endpoint belongs.
	//
	// example:
	//
	// vpc-fdfhkjafhjvcvdjf****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// Indicates whether the domain name of the nearest endpoint that is associated with the endpoint service is resolved first. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	ZoneAffinityEnabled *bool `json:"ZoneAffinityEnabled,omitempty" xml:"ZoneAffinityEnabled,omitempty"`
	// The number of private IP addresses that are assigned to an elastic network interface (ENI) in each zone. Only **1*	- is returned.
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
