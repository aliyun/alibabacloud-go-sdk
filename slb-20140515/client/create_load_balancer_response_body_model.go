// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLoadBalancerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAddress(v string) *CreateLoadBalancerResponseBody
	GetAddress() *string
	SetAddressIPVersion(v string) *CreateLoadBalancerResponseBody
	GetAddressIPVersion() *string
	SetLoadBalancerId(v string) *CreateLoadBalancerResponseBody
	GetLoadBalancerId() *string
	SetLoadBalancerName(v string) *CreateLoadBalancerResponseBody
	GetLoadBalancerName() *string
	SetNetworkType(v string) *CreateLoadBalancerResponseBody
	GetNetworkType() *string
	SetOrderId(v int64) *CreateLoadBalancerResponseBody
	GetOrderId() *int64
	SetRequestId(v string) *CreateLoadBalancerResponseBody
	GetRequestId() *string
	SetResourceGroupId(v string) *CreateLoadBalancerResponseBody
	GetResourceGroupId() *string
	SetVSwitchId(v string) *CreateLoadBalancerResponseBody
	GetVSwitchId() *string
	SetVpcId(v string) *CreateLoadBalancerResponseBody
	GetVpcId() *string
}

type CreateLoadBalancerResponseBody struct {
	// The IP address that is allocated to the CLB instance.
	//
	// example:
	//
	// 42.XX.XX.6
	Address *string `json:"Address,omitempty" xml:"Address,omitempty"`
	// The IP version that is used by the CLB instance.
	//
	// example:
	//
	// ipv4
	AddressIPVersion *string `json:"AddressIPVersion,omitempty" xml:"AddressIPVersion,omitempty"`
	// The CLB instance ID.
	//
	// example:
	//
	// lb-hddhfjg****
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	// The CLB instance name.
	//
	// example:
	//
	// lb-bp1o94dp5i6ea****
	LoadBalancerName *string `json:"LoadBalancerName,omitempty" xml:"LoadBalancerName,omitempty"`
	// The network type of the CLB instance. Valid values:
	//
	// 	- **vpc**: VPC
	//
	// 	- **Classic**: classic network
	//
	// example:
	//
	// classic
	NetworkType *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	// The order ID of the subscription CLB instance.
	//
	// example:
	//
	// 20212961978****
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 365F4154-92F6-4AE4-92F8-7FF34B540710
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the resource group to which the CLB instance belongs.
	//
	// example:
	//
	// rg-atstuj3rto****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The ID of the vSwitch to which the CLB instance belongs.
	//
	// example:
	//
	// vsw-255ecr****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The ID of the VPC to which the CLB instance belongs.
	//
	// example:
	//
	// vpc-25dvzy9****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s CreateLoadBalancerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateLoadBalancerResponseBody) GoString() string {
	return s.String()
}

func (s *CreateLoadBalancerResponseBody) GetAddress() *string {
	return s.Address
}

func (s *CreateLoadBalancerResponseBody) GetAddressIPVersion() *string {
	return s.AddressIPVersion
}

func (s *CreateLoadBalancerResponseBody) GetLoadBalancerId() *string {
	return s.LoadBalancerId
}

func (s *CreateLoadBalancerResponseBody) GetLoadBalancerName() *string {
	return s.LoadBalancerName
}

func (s *CreateLoadBalancerResponseBody) GetNetworkType() *string {
	return s.NetworkType
}

func (s *CreateLoadBalancerResponseBody) GetOrderId() *int64 {
	return s.OrderId
}

func (s *CreateLoadBalancerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateLoadBalancerResponseBody) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateLoadBalancerResponseBody) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *CreateLoadBalancerResponseBody) GetVpcId() *string {
	return s.VpcId
}

func (s *CreateLoadBalancerResponseBody) SetAddress(v string) *CreateLoadBalancerResponseBody {
	s.Address = &v
	return s
}

func (s *CreateLoadBalancerResponseBody) SetAddressIPVersion(v string) *CreateLoadBalancerResponseBody {
	s.AddressIPVersion = &v
	return s
}

func (s *CreateLoadBalancerResponseBody) SetLoadBalancerId(v string) *CreateLoadBalancerResponseBody {
	s.LoadBalancerId = &v
	return s
}

func (s *CreateLoadBalancerResponseBody) SetLoadBalancerName(v string) *CreateLoadBalancerResponseBody {
	s.LoadBalancerName = &v
	return s
}

func (s *CreateLoadBalancerResponseBody) SetNetworkType(v string) *CreateLoadBalancerResponseBody {
	s.NetworkType = &v
	return s
}

func (s *CreateLoadBalancerResponseBody) SetOrderId(v int64) *CreateLoadBalancerResponseBody {
	s.OrderId = &v
	return s
}

func (s *CreateLoadBalancerResponseBody) SetRequestId(v string) *CreateLoadBalancerResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateLoadBalancerResponseBody) SetResourceGroupId(v string) *CreateLoadBalancerResponseBody {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateLoadBalancerResponseBody) SetVSwitchId(v string) *CreateLoadBalancerResponseBody {
	s.VSwitchId = &v
	return s
}

func (s *CreateLoadBalancerResponseBody) SetVpcId(v string) *CreateLoadBalancerResponseBody {
	s.VpcId = &v
	return s
}

func (s *CreateLoadBalancerResponseBody) Validate() error {
	return dara.Validate(s)
}
