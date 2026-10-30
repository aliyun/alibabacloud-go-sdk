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
	// The collection of resource information.
	ElbAvailableResourceInfo []*DescribeElbAvailableResourceInfoResponseBodyElbAvailableResourceInfo `json:"ElbAvailableResourceInfo,omitempty" xml:"ElbAvailableResourceInfo,omitempty" type:"Repeated"`
	// The common parameter. The ID of each request is unique and active for troubleshooting and locating issues.
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
	if s.ElbAvailableResourceInfo != nil {
		for _, item := range s.ElbAvailableResourceInfo {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeElbAvailableResourceInfoResponseBodyElbAvailableResourceInfo struct {
	Ability []*string `json:"Ability,omitempty" xml:"Ability,omitempty" type:"Repeated"`
	// The area.
	//
	// example:
	//
	// SouthEast
	Area *string `json:"Area,omitempty" xml:"Area,omitempty"`
	// The Quantity available for purchase.
	//
	// example:
	//
	// 1
	CanBuyCount *string `json:"CanBuyCount,omitempty" xml:"CanBuyCount,omitempty"`
	// The English name of the edge zone.
	//
	// example:
	//
	// cn-guangdong-10
	EnName *string `json:"EnName,omitempty" xml:"EnName,omitempty"`
	// The edge zone ID.
	//
	// example:
	//
	// cn-guangdong-10
	EnsRegionId *string `json:"EnsRegionId,omitempty" xml:"EnsRegionId,omitempty"`
	// The specification of the load balancing instance.
	LoadBalancerSpec []*string `json:"LoadBalancerSpec,omitempty" xml:"LoadBalancerSpec,omitempty" type:"Repeated"`
	// The name of the edge zone.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The province.
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
