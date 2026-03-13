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
	LoadBalancers *DescribeLoadBalancersResponseBodyLoadBalancers `json:"LoadBalancers,omitempty" xml:"LoadBalancers,omitempty" type:"Struct"`
	PageNumber    *int32                                          `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize      *int32                                          `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId     *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount    *int32                                          `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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
	Address            *string `json:"Address,omitempty" xml:"Address,omitempty"`
	AddressType        *string `json:"AddressType,omitempty" xml:"AddressType,omitempty"`
	CreateTime         *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	CreateTimeStamp    *int64  `json:"CreateTimeStamp,omitempty" xml:"CreateTimeStamp,omitempty"`
	InternetChargeType *string `json:"InternetChargeType,omitempty" xml:"InternetChargeType,omitempty"`
	LoadBalancerId     *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	LoadBalancerName   *string `json:"LoadBalancerName,omitempty" xml:"LoadBalancerName,omitempty"`
	LoadBalancerStatus *string `json:"LoadBalancerStatus,omitempty" xml:"LoadBalancerStatus,omitempty"`
	MasterZoneId       *string `json:"MasterZoneId,omitempty" xml:"MasterZoneId,omitempty"`
	NetworkType        *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	PayType            *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	RegionId           *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RegionIdAlias      *string `json:"RegionIdAlias,omitempty" xml:"RegionIdAlias,omitempty"`
	ResourceGroupId    *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	SlaveZoneId        *string `json:"SlaveZoneId,omitempty" xml:"SlaveZoneId,omitempty"`
	VSwitchId          *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	VpcId              *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
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

func (s *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer) GetAddressType() *string {
	return s.AddressType
}

func (s *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer) GetCreateTimeStamp() *int64 {
	return s.CreateTimeStamp
}

func (s *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer) GetInternetChargeType() *string {
	return s.InternetChargeType
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

func (s *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer) GetMasterZoneId() *string {
	return s.MasterZoneId
}

func (s *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer) GetNetworkType() *string {
	return s.NetworkType
}

func (s *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer) GetPayType() *string {
	return s.PayType
}

func (s *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer) GetRegionIdAlias() *string {
	return s.RegionIdAlias
}

func (s *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer) GetSlaveZoneId() *string {
	return s.SlaveZoneId
}

func (s *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer) SetAddress(v string) *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer {
	s.Address = &v
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

func (s *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer) SetCreateTimeStamp(v int64) *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer {
	s.CreateTimeStamp = &v
	return s
}

func (s *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer) SetInternetChargeType(v string) *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer {
	s.InternetChargeType = &v
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

func (s *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer) SetMasterZoneId(v string) *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer {
	s.MasterZoneId = &v
	return s
}

func (s *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer) SetNetworkType(v string) *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer {
	s.NetworkType = &v
	return s
}

func (s *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer) SetPayType(v string) *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer {
	s.PayType = &v
	return s
}

func (s *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer) SetRegionId(v string) *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer {
	s.RegionId = &v
	return s
}

func (s *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer) SetRegionIdAlias(v string) *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer {
	s.RegionIdAlias = &v
	return s
}

func (s *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer) SetResourceGroupId(v string) *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer) SetSlaveZoneId(v string) *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer {
	s.SlaveZoneId = &v
	return s
}

func (s *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer) SetVSwitchId(v string) *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer {
	s.VSwitchId = &v
	return s
}

func (s *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer) SetVpcId(v string) *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer {
	s.VpcId = &v
	return s
}

func (s *DescribeLoadBalancersResponseBodyLoadBalancersLoadBalancer) Validate() error {
	return dara.Validate(s)
}
