// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeElbAvailableResourceInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetElbAvailableResourceInfo(v []*DescribeElbAvailableResourceInfoResponseBodyElbAvailableResourceInfo) *DescribeElbAvailableResourceInfoResponseBody
	GetElbAvailableResourceInfo() []*DescribeElbAvailableResourceInfoResponseBodyElbAvailableResourceInfo
	SetRequestId(v string) *DescribeElbAvailableResourceInfoResponseBody
	GetRequestId() *string
}

type DescribeElbAvailableResourceInfoResponseBody struct {
	// The information about resources.
	ElbAvailableResourceInfo []*DescribeElbAvailableResourceInfoResponseBodyElbAvailableResourceInfo `json:"ElbAvailableResourceInfo,omitempty" xml:"ElbAvailableResourceInfo,omitempty" type:"Repeated"`
	// The ID of the request. This parameter is a common parameter. Each request has a unique ID. You can use the ID to troubleshoot issues.
	//
	// example:
	//
	// 25AAD194-4A37-51CF-B1CA-1E86FDAC23A6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeElbAvailableResourceInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeElbAvailableResourceInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeElbAvailableResourceInfoResponseBody) GetElbAvailableResourceInfo() []*DescribeElbAvailableResourceInfoResponseBodyElbAvailableResourceInfo {
	return s.ElbAvailableResourceInfo
}

func (s *DescribeElbAvailableResourceInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeElbAvailableResourceInfoResponseBody) SetElbAvailableResourceInfo(v []*DescribeElbAvailableResourceInfoResponseBodyElbAvailableResourceInfo) *DescribeElbAvailableResourceInfoResponseBody {
	s.ElbAvailableResourceInfo = v
	return s
}

func (s *DescribeElbAvailableResourceInfoResponseBody) SetRequestId(v string) *DescribeElbAvailableResourceInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeElbAvailableResourceInfoResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeElbAvailableResourceInfoResponseBodyElbAvailableResourceInfo struct {
	Ability []*string `json:"Ability,omitempty" xml:"Ability,omitempty" type:"Repeated"`
	// The ID of the region.
	//
	// example:
	//
	// SouthEast
	Area *string `json:"Area,omitempty" xml:"Area,omitempty"`
	// The number of resources that you can purchase.
	//
	// example:
	//
	// 1
	CanBuyCount *string `json:"CanBuyCount,omitempty" xml:"CanBuyCount,omitempty"`
	// The name of the node.
	//
	// example:
	//
	// cn-guangdong-10
	EnName *string `json:"EnName,omitempty" xml:"EnName,omitempty"`
	// The ID of the Edge Node Service (ENS) node.
	//
	// example:
	//
	// cn-guangdong-10
	EnsRegionId *string `json:"EnsRegionId,omitempty" xml:"EnsRegionId,omitempty"`
	// The specifications of the ELB instances.
	LoadBalancerSpec []*string `json:"LoadBalancerSpec,omitempty" xml:"LoadBalancerSpec,omitempty" type:"Repeated"`
	// The Chinese name of the node.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The province where the node is deployed.
	//
	// example:
	//
	// Shanghai
	Province *string `json:"Province,omitempty" xml:"Province,omitempty"`
}

func (s DescribeElbAvailableResourceInfoResponseBodyElbAvailableResourceInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeElbAvailableResourceInfoResponseBodyElbAvailableResourceInfo) GoString() string {
	return s.String()
}

func (s *DescribeElbAvailableResourceInfoResponseBodyElbAvailableResourceInfo) GetAbility() []*string {
	return s.Ability
}

func (s *DescribeElbAvailableResourceInfoResponseBodyElbAvailableResourceInfo) GetArea() *string {
	return s.Area
}

func (s *DescribeElbAvailableResourceInfoResponseBodyElbAvailableResourceInfo) GetCanBuyCount() *string {
	return s.CanBuyCount
}

func (s *DescribeElbAvailableResourceInfoResponseBodyElbAvailableResourceInfo) GetEnName() *string {
	return s.EnName
}

func (s *DescribeElbAvailableResourceInfoResponseBodyElbAvailableResourceInfo) GetEnsRegionId() *string {
	return s.EnsRegionId
}

func (s *DescribeElbAvailableResourceInfoResponseBodyElbAvailableResourceInfo) GetLoadBalancerSpec() []*string {
	return s.LoadBalancerSpec
}

func (s *DescribeElbAvailableResourceInfoResponseBodyElbAvailableResourceInfo) GetName() *string {
	return s.Name
}

func (s *DescribeElbAvailableResourceInfoResponseBodyElbAvailableResourceInfo) GetProvince() *string {
	return s.Province
}

func (s *DescribeElbAvailableResourceInfoResponseBodyElbAvailableResourceInfo) SetAbility(v []*string) *DescribeElbAvailableResourceInfoResponseBodyElbAvailableResourceInfo {
	s.Ability = v
	return s
}

func (s *DescribeElbAvailableResourceInfoResponseBodyElbAvailableResourceInfo) SetArea(v string) *DescribeElbAvailableResourceInfoResponseBodyElbAvailableResourceInfo {
	s.Area = &v
	return s
}

func (s *DescribeElbAvailableResourceInfoResponseBodyElbAvailableResourceInfo) SetCanBuyCount(v string) *DescribeElbAvailableResourceInfoResponseBodyElbAvailableResourceInfo {
	s.CanBuyCount = &v
	return s
}

func (s *DescribeElbAvailableResourceInfoResponseBodyElbAvailableResourceInfo) SetEnName(v string) *DescribeElbAvailableResourceInfoResponseBodyElbAvailableResourceInfo {
	s.EnName = &v
	return s
}

func (s *DescribeElbAvailableResourceInfoResponseBodyElbAvailableResourceInfo) SetEnsRegionId(v string) *DescribeElbAvailableResourceInfoResponseBodyElbAvailableResourceInfo {
	s.EnsRegionId = &v
	return s
}

func (s *DescribeElbAvailableResourceInfoResponseBodyElbAvailableResourceInfo) SetLoadBalancerSpec(v []*string) *DescribeElbAvailableResourceInfoResponseBodyElbAvailableResourceInfo {
	s.LoadBalancerSpec = v
	return s
}

func (s *DescribeElbAvailableResourceInfoResponseBodyElbAvailableResourceInfo) SetName(v string) *DescribeElbAvailableResourceInfoResponseBodyElbAvailableResourceInfo {
	s.Name = &v
	return s
}

func (s *DescribeElbAvailableResourceInfoResponseBodyElbAvailableResourceInfo) SetProvince(v string) *DescribeElbAvailableResourceInfoResponseBodyElbAvailableResourceInfo {
	s.Province = &v
	return s
}

func (s *DescribeElbAvailableResourceInfoResponseBodyElbAvailableResourceInfo) Validate() error {
	return dara.Validate(s)
}
