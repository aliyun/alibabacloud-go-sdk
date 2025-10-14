// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLoadBalancersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLoadBalancers(v *DescribeLoadBalancersResponseBodyLoadBalancers) *DescribeLoadBalancersResponseBody
	GetLoadBalancers() *DescribeLoadBalancersResponseBodyLoadBalancers
	SetPageNumber(v int32) *DescribeLoadBalancersResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeLoadBalancersResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeLoadBalancersResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeLoadBalancersResponseBody
	GetTotalCount() *int32
}

type DescribeLoadBalancersResponseBody struct {
	// An array of ELB instances.
	LoadBalancers *DescribeLoadBalancersResponseBodyLoadBalancers `json:"LoadBalancers,omitempty" xml:"LoadBalancers,omitempty" type:"Struct"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: 10. Valid values: **10*	- to **100**.
	//
	// example:
	//
	// 100
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of entries returned.
	//
	// example:
	//
	// 6
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeLoadBalancersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeLoadBalancersResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLoadBalancersResponseBody) GetLoadBalancers() *DescribeLoadBalancersResponseBodyLoadBalancers {
	return s.LoadBalancers
}

func (s *DescribeLoadBalancersResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeLoadBalancersResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeLoadBalancersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeLoadBalancersResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeLoadBalancersResponseBody) SetLoadBalancers(v *DescribeLoadBalancersResponseBodyLoadBalancers) *DescribeLoadBalancersResponseBody {
	s.LoadBalancers = v
	return s
}

func (s *DescribeLoadBalancersResponseBody) SetPageNumber(v int32) *DescribeLoadBalancersResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeLoadBalancersResponseBody) SetPageSize(v int32) *DescribeLoadBalancersResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeLoadBalancersResponseBody) SetRequestId(v string) *DescribeLoadBalancersResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLoadBalancersResponseBody) SetTotalCount(v int32) *DescribeLoadBalancersResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeLoadBalancersResponseBody) Validate() error {
	if s.LoadBalancers != nil {
		if err := s.LoadBalancers.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeLoadBalancersResponseBodyLoadBalancers struct {
	LoadBalancer []*DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer `json:"LoadBalancer,omitempty" xml:"LoadBalancer,omitempty" type:"Repeated"`
}

func (s DescribeLoadBalancersResponseBodyLoadBalancers) String() string {
	return dara.Prettify(s)
}

func (s DescribeLoadBalancersResponseBodyLoadBalancers) GoString() string {
	return s.String()
}

func (s *DescribeLoadBalancersResponseBodyLoadBalancers) GetLoadBalancer() []*DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer {
	return s.LoadBalancer
}

func (s *DescribeLoadBalancersResponseBodyLoadBalancers) SetLoadBalancer(v []*DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer) *DescribeLoadBalancersResponseBodyLoadBalancers {
	s.LoadBalancer = v
	return s
}

func (s *DescribeLoadBalancersResponseBodyLoadBalancers) Validate() error {
	if s.LoadBalancer != nil {
		for _, item := range s.LoadBalancer {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer struct {
	// The IP address that the ELB instance uses to provide services.
	//
	// example:
	//
	// 10.10.XX.XX
	Address *string `json:"Address,omitempty" xml:"Address,omitempty"`
	// The IP version. Valid values: ipv4 and ipv6.
	//
	// example:
	//
	// ipv4
	AddressIPVersion *string `json:"AddressIPVersion,omitempty" xml:"AddressIPVersion,omitempty"`
	AddressType      *string `json:"AddressType,omitempty" xml:"AddressType,omitempty"`
	// The time when the ELB instance was created. The time is displayed in UTC.
	//
	// example:
	//
	// 2021-05-06T11:13:41Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The ID of the ENS node.
	//
	// example:
	//
	// cn-wuhan-telecom
	EnsRegionId *string `json:"EnsRegionId,omitempty" xml:"EnsRegionId,omitempty"`
	// The ID of the ELB instance.
	//
	// example:
	//
	// lb-5snthcyu1x10g7tywj7iu****
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	// The name of the ELB instance.
	//
	// example:
	//
	// example
	LoadBalancerName *string `json:"LoadBalancerName,omitempty" xml:"LoadBalancerName,omitempty"`
	// The status of the listener for the ELB instance. Valid values:
	//
	// 	- **Active**: The listener for the instance can forward the received traffic based on forwarding rules.
	//
	// 	- **InActive**: The listener for the instance does not forward the received traffic.
	//
	// example:
	//
	// InActive
	LoadBalancerStatus *string `json:"LoadBalancerStatus,omitempty" xml:"LoadBalancerStatus,omitempty"`
	LoadBalancerType   *string `json:"LoadBalancerType,omitempty" xml:"LoadBalancerType,omitempty"`
	// The ID of the network.
	//
	// example:
	//
	// n-5rz0rj1caexauilpsjx0w****
	NetworkId *string `json:"NetworkId,omitempty" xml:"NetworkId,omitempty"`
	// The billing method. Valid values:
	//
	// 	- **PrePaid**: subscription.
	//
	// 	- **PostPaid**: pay-as-you-go. Only this billing method is supported.
	//
	// example:
	//
	// PostPaid
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// The ID of the vSwitch.
	//
	// example:
	//
	// vsw-5rllcjb3ol6duzjdnbm1om****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
}

func (s DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer) String() string {
	return dara.Prettify(s)
}

func (s DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer) GoString() string {
	return s.String()
}

func (s *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer) GetAddress() *string {
	return s.Address
}

func (s *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer) GetAddressIPVersion() *string {
	return s.AddressIPVersion
}

func (s *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer) GetAddressType() *string {
	return s.AddressType
}

func (s *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer) GetEnsRegionId() *string {
	return s.EnsRegionId
}

func (s *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer) GetLoadBalancerId() *string {
	return s.LoadBalancerId
}

func (s *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer) GetLoadBalancerName() *string {
	return s.LoadBalancerName
}

func (s *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer) GetLoadBalancerStatus() *string {
	return s.LoadBalancerStatus
}

func (s *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer) GetLoadBalancerType() *string {
	return s.LoadBalancerType
}

func (s *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer) GetNetworkId() *string {
	return s.NetworkId
}

func (s *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer) GetPayType() *string {
	return s.PayType
}

func (s *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer) SetAddress(v string) *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer {
	s.Address = &v
	return s
}

func (s *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer) SetAddressIPVersion(v string) *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer {
	s.AddressIPVersion = &v
	return s
}

func (s *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer) SetAddressType(v string) *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer {
	s.AddressType = &v
	return s
}

func (s *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer) SetCreateTime(v string) *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer {
	s.CreateTime = &v
	return s
}

func (s *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer) SetEnsRegionId(v string) *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer {
	s.EnsRegionId = &v
	return s
}

func (s *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer) SetLoadBalancerId(v string) *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer {
	s.LoadBalancerId = &v
	return s
}

func (s *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer) SetLoadBalancerName(v string) *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer {
	s.LoadBalancerName = &v
	return s
}

func (s *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer) SetLoadBalancerStatus(v string) *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer {
	s.LoadBalancerStatus = &v
	return s
}

func (s *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer) SetLoadBalancerType(v string) *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer {
	s.LoadBalancerType = &v
	return s
}

func (s *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer) SetNetworkId(v string) *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer {
	s.NetworkId = &v
	return s
}

func (s *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer) SetPayType(v string) *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer {
	s.PayType = &v
	return s
}

func (s *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer) SetVSwitchId(v string) *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer {
	s.VSwitchId = &v
	return s
}

func (s *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer) Validate() error {
	return dara.Validate(s)
}
