// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLoadBalancersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddress(v string) *DescribeLoadBalancersRequest
	GetAddress() *string
	SetEnsRegionId(v string) *DescribeLoadBalancersRequest
	GetEnsRegionId() *string
	SetEnsRegionIds(v []*string) *DescribeLoadBalancersRequest
	GetEnsRegionIds() []*string
	SetLoadBalancerId(v string) *DescribeLoadBalancersRequest
	GetLoadBalancerId() *string
	SetLoadBalancerName(v string) *DescribeLoadBalancersRequest
	GetLoadBalancerName() *string
	SetLoadBalancerStatus(v string) *DescribeLoadBalancersRequest
	GetLoadBalancerStatus() *string
	SetLoadBalancerType(v string) *DescribeLoadBalancersRequest
	GetLoadBalancerType() *string
	SetNetworkId(v string) *DescribeLoadBalancersRequest
	GetNetworkId() *string
	SetPageNumber(v int32) *DescribeLoadBalancersRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeLoadBalancersRequest
	GetPageSize() *int32
	SetServerId(v string) *DescribeLoadBalancersRequest
	GetServerId() *string
	SetVSwitchId(v string) *DescribeLoadBalancersRequest
	GetVSwitchId() *string
}

type DescribeLoadBalancersRequest struct {
	// The IP address that the ELB instance uses to provide services.
	//
	// example:
	//
	// 10.0.XX.XX
	Address *string `json:"Address,omitempty" xml:"Address,omitempty"`
	// The ID of the Edge Node Service (ENS) node.
	//
	// example:
	//
	// cn-guangzhou-10
	EnsRegionId *string `json:"EnsRegionId,omitempty" xml:"EnsRegionId,omitempty"`
	// The IDs of the Edge Node Service (ENS) nodes.
	EnsRegionIds []*string `json:"EnsRegionIds,omitempty" xml:"EnsRegionIds,omitempty" type:"Repeated"`
	// The ID of the ELB instance.
	//
	// example:
	//
	// lb-5q73cv04zeyh43lh74lp4****
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
	// n-5s9ayrxsd9hszrlt5fgv2****
	NetworkId *string `json:"NetworkId,omitempty" xml:"NetworkId,omitempty"`
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Default value: 10. Valid values: **10*	- to **100**.
	//
	// example:
	//
	// 100
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the backend server.
	//
	// example:
	//
	// i-5f67ffjc004wwz0t****
	ServerId *string `json:"ServerId,omitempty" xml:"ServerId,omitempty"`
	// The ID of the vSwitch.
	//
	// example:
	//
	// vsw-5sy773iy25rulsmgskmba****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
}

func (s DescribeLoadBalancersRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeLoadBalancersRequest) GoString() string {
	return s.String()
}

func (s *DescribeLoadBalancersRequest) GetAddress() *string {
	return s.Address
}

func (s *DescribeLoadBalancersRequest) GetEnsRegionId() *string {
	return s.EnsRegionId
}

func (s *DescribeLoadBalancersRequest) GetEnsRegionIds() []*string {
	return s.EnsRegionIds
}

func (s *DescribeLoadBalancersRequest) GetLoadBalancerId() *string {
	return s.LoadBalancerId
}

func (s *DescribeLoadBalancersRequest) GetLoadBalancerName() *string {
	return s.LoadBalancerName
}

func (s *DescribeLoadBalancersRequest) GetLoadBalancerStatus() *string {
	return s.LoadBalancerStatus
}

func (s *DescribeLoadBalancersRequest) GetLoadBalancerType() *string {
	return s.LoadBalancerType
}

func (s *DescribeLoadBalancersRequest) GetNetworkId() *string {
	return s.NetworkId
}

func (s *DescribeLoadBalancersRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeLoadBalancersRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeLoadBalancersRequest) GetServerId() *string {
	return s.ServerId
}

func (s *DescribeLoadBalancersRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *DescribeLoadBalancersRequest) SetAddress(v string) *DescribeLoadBalancersRequest {
	s.Address = &v
	return s
}

func (s *DescribeLoadBalancersRequest) SetEnsRegionId(v string) *DescribeLoadBalancersRequest {
	s.EnsRegionId = &v
	return s
}

func (s *DescribeLoadBalancersRequest) SetEnsRegionIds(v []*string) *DescribeLoadBalancersRequest {
	s.EnsRegionIds = v
	return s
}

func (s *DescribeLoadBalancersRequest) SetLoadBalancerId(v string) *DescribeLoadBalancersRequest {
	s.LoadBalancerId = &v
	return s
}

func (s *DescribeLoadBalancersRequest) SetLoadBalancerName(v string) *DescribeLoadBalancersRequest {
	s.LoadBalancerName = &v
	return s
}

func (s *DescribeLoadBalancersRequest) SetLoadBalancerStatus(v string) *DescribeLoadBalancersRequest {
	s.LoadBalancerStatus = &v
	return s
}

func (s *DescribeLoadBalancersRequest) SetLoadBalancerType(v string) *DescribeLoadBalancersRequest {
	s.LoadBalancerType = &v
	return s
}

func (s *DescribeLoadBalancersRequest) SetNetworkId(v string) *DescribeLoadBalancersRequest {
	s.NetworkId = &v
	return s
}

func (s *DescribeLoadBalancersRequest) SetPageNumber(v int32) *DescribeLoadBalancersRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeLoadBalancersRequest) SetPageSize(v int32) *DescribeLoadBalancersRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeLoadBalancersRequest) SetServerId(v string) *DescribeLoadBalancersRequest {
	s.ServerId = &v
	return s
}

func (s *DescribeLoadBalancersRequest) SetVSwitchId(v string) *DescribeLoadBalancersRequest {
	s.VSwitchId = &v
	return s
}

func (s *DescribeLoadBalancersRequest) Validate() error {
	return dara.Validate(s)
}
