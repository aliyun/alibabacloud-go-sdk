// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeIpv6AddressesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetIpv6Addresses(v *DescribeIpv6AddressesResponseBodyIpv6Addresses) *DescribeIpv6AddressesResponseBody
	GetIpv6Addresses() *DescribeIpv6AddressesResponseBodyIpv6Addresses
	SetPageNumber(v int32) *DescribeIpv6AddressesResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeIpv6AddressesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeIpv6AddressesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeIpv6AddressesResponseBody
	GetTotalCount() *int32
}

type DescribeIpv6AddressesResponseBody struct {
	Ipv6Addresses *DescribeIpv6AddressesResponseBodyIpv6Addresses `json:"Ipv6Addresses,omitempty" xml:"Ipv6Addresses,omitempty" type:"Struct"`
	// The page number of the returned page. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page. Maximum value: **50**. Default value: **10**.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// AA4486A8-B6AE-469E-AB09-820EF8ECFA2B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of returned entries.
	//
	// example:
	//
	// 2
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeIpv6AddressesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeIpv6AddressesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeIpv6AddressesResponseBody) GetIpv6Addresses() *DescribeIpv6AddressesResponseBodyIpv6Addresses {
	return s.Ipv6Addresses
}

func (s *DescribeIpv6AddressesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeIpv6AddressesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeIpv6AddressesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeIpv6AddressesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeIpv6AddressesResponseBody) SetIpv6Addresses(v *DescribeIpv6AddressesResponseBodyIpv6Addresses) *DescribeIpv6AddressesResponseBody {
	s.Ipv6Addresses = v
	return s
}

func (s *DescribeIpv6AddressesResponseBody) SetPageNumber(v int32) *DescribeIpv6AddressesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeIpv6AddressesResponseBody) SetPageSize(v int32) *DescribeIpv6AddressesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeIpv6AddressesResponseBody) SetRequestId(v string) *DescribeIpv6AddressesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeIpv6AddressesResponseBody) SetTotalCount(v int32) *DescribeIpv6AddressesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeIpv6AddressesResponseBody) Validate() error {
	if s.Ipv6Addresses != nil {
		if err := s.Ipv6Addresses.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeIpv6AddressesResponseBodyIpv6Addresses struct {
	Ipv6Address []*DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6Address `json:"Ipv6Address,omitempty" xml:"Ipv6Address,omitempty" type:"Repeated"`
}

func (s DescribeIpv6AddressesResponseBodyIpv6Addresses) String() string {
	return dara.Prettify(s)
}

func (s DescribeIpv6AddressesResponseBodyIpv6Addresses) GoString() string {
	return s.String()
}

func (s *DescribeIpv6AddressesResponseBodyIpv6Addresses) GetIpv6Address() []*DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6Address {
	return s.Ipv6Address
}

func (s *DescribeIpv6AddressesResponseBodyIpv6Addresses) SetIpv6Address(v []*DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6Address) *DescribeIpv6AddressesResponseBodyIpv6Addresses {
	s.Ipv6Address = v
	return s
}

func (s *DescribeIpv6AddressesResponseBodyIpv6Addresses) Validate() error {
	if s.Ipv6Address != nil {
		for _, item := range s.Ipv6Address {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6Address struct {
	AddressType            *string                                                                         `json:"AddressType,omitempty" xml:"AddressType,omitempty"`
	AllocationTime         *string                                                                         `json:"AllocationTime,omitempty" xml:"AllocationTime,omitempty"`
	AssociatedInstanceId   *string                                                                         `json:"AssociatedInstanceId,omitempty" xml:"AssociatedInstanceId,omitempty"`
	AssociatedInstanceType *string                                                                         `json:"AssociatedInstanceType,omitempty" xml:"AssociatedInstanceType,omitempty"`
	Ipv6Address            *string                                                                         `json:"Ipv6Address,omitempty" xml:"Ipv6Address,omitempty"`
	Ipv6AddressDescription *string                                                                         `json:"Ipv6AddressDescription,omitempty" xml:"Ipv6AddressDescription,omitempty"`
	Ipv6AddressId          *string                                                                         `json:"Ipv6AddressId,omitempty" xml:"Ipv6AddressId,omitempty"`
	Ipv6AddressName        *string                                                                         `json:"Ipv6AddressName,omitempty" xml:"Ipv6AddressName,omitempty"`
	Ipv6GatewayId          *string                                                                         `json:"Ipv6GatewayId,omitempty" xml:"Ipv6GatewayId,omitempty"`
	Ipv6InternetBandwidth  *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6AddressIpv6InternetBandwidth `json:"Ipv6InternetBandwidth,omitempty" xml:"Ipv6InternetBandwidth,omitempty" type:"Struct"`
	Ipv6Isp                *string                                                                         `json:"Ipv6Isp,omitempty" xml:"Ipv6Isp,omitempty"`
	NetworkType            *string                                                                         `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	RealBandwidth          *int32                                                                          `json:"RealBandwidth,omitempty" xml:"RealBandwidth,omitempty"`
	ResourceGroupId        *string                                                                         `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ServiceManaged         *int32                                                                          `json:"ServiceManaged,omitempty" xml:"ServiceManaged,omitempty"`
	Status                 *string                                                                         `json:"Status,omitempty" xml:"Status,omitempty"`
	Tags                   *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6AddressTags                  `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	VSwitchId              *string                                                                         `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	VpcId                  *string                                                                         `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6Address) String() string {
	return dara.Prettify(s)
}

func (s DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6Address) GoString() string {
	return s.String()
}

func (s *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6Address) GetAddressType() *string {
	return s.AddressType
}

func (s *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6Address) GetAllocationTime() *string {
	return s.AllocationTime
}

func (s *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6Address) GetAssociatedInstanceId() *string {
	return s.AssociatedInstanceId
}

func (s *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6Address) GetAssociatedInstanceType() *string {
	return s.AssociatedInstanceType
}

func (s *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6Address) GetIpv6Address() *string {
	return s.Ipv6Address
}

func (s *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6Address) GetIpv6AddressDescription() *string {
	return s.Ipv6AddressDescription
}

func (s *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6Address) GetIpv6AddressId() *string {
	return s.Ipv6AddressId
}

func (s *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6Address) GetIpv6AddressName() *string {
	return s.Ipv6AddressName
}

func (s *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6Address) GetIpv6GatewayId() *string {
	return s.Ipv6GatewayId
}

func (s *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6Address) GetIpv6InternetBandwidth() *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6AddressIpv6InternetBandwidth {
	return s.Ipv6InternetBandwidth
}

func (s *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6Address) GetIpv6Isp() *string {
	return s.Ipv6Isp
}

func (s *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6Address) GetNetworkType() *string {
	return s.NetworkType
}

func (s *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6Address) GetRealBandwidth() *int32 {
	return s.RealBandwidth
}

func (s *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6Address) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6Address) GetServiceManaged() *int32 {
	return s.ServiceManaged
}

func (s *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6Address) GetStatus() *string {
	return s.Status
}

func (s *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6Address) GetTags() *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6AddressTags {
	return s.Tags
}

func (s *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6Address) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6Address) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6Address) SetAddressType(v string) *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6Address {
	s.AddressType = &v
	return s
}

func (s *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6Address) SetAllocationTime(v string) *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6Address {
	s.AllocationTime = &v
	return s
}

func (s *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6Address) SetAssociatedInstanceId(v string) *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6Address {
	s.AssociatedInstanceId = &v
	return s
}

func (s *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6Address) SetAssociatedInstanceType(v string) *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6Address {
	s.AssociatedInstanceType = &v
	return s
}

func (s *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6Address) SetIpv6Address(v string) *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6Address {
	s.Ipv6Address = &v
	return s
}

func (s *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6Address) SetIpv6AddressDescription(v string) *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6Address {
	s.Ipv6AddressDescription = &v
	return s
}

func (s *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6Address) SetIpv6AddressId(v string) *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6Address {
	s.Ipv6AddressId = &v
	return s
}

func (s *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6Address) SetIpv6AddressName(v string) *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6Address {
	s.Ipv6AddressName = &v
	return s
}

func (s *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6Address) SetIpv6GatewayId(v string) *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6Address {
	s.Ipv6GatewayId = &v
	return s
}

func (s *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6Address) SetIpv6InternetBandwidth(v *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6AddressIpv6InternetBandwidth) *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6Address {
	s.Ipv6InternetBandwidth = v
	return s
}

func (s *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6Address) SetIpv6Isp(v string) *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6Address {
	s.Ipv6Isp = &v
	return s
}

func (s *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6Address) SetNetworkType(v string) *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6Address {
	s.NetworkType = &v
	return s
}

func (s *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6Address) SetRealBandwidth(v int32) *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6Address {
	s.RealBandwidth = &v
	return s
}

func (s *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6Address) SetResourceGroupId(v string) *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6Address {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6Address) SetServiceManaged(v int32) *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6Address {
	s.ServiceManaged = &v
	return s
}

func (s *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6Address) SetStatus(v string) *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6Address {
	s.Status = &v
	return s
}

func (s *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6Address) SetTags(v *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6AddressTags) *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6Address {
	s.Tags = v
	return s
}

func (s *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6Address) SetVSwitchId(v string) *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6Address {
	s.VSwitchId = &v
	return s
}

func (s *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6Address) SetVpcId(v string) *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6Address {
	s.VpcId = &v
	return s
}

func (s *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6Address) Validate() error {
	if s.Ipv6InternetBandwidth != nil {
		if err := s.Ipv6InternetBandwidth.Validate(); err != nil {
			return err
		}
	}
	if s.Tags != nil {
		if err := s.Tags.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6AddressIpv6InternetBandwidth struct {
	Bandwidth                     *int32  `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	BusinessStatus                *string `json:"BusinessStatus,omitempty" xml:"BusinessStatus,omitempty"`
	HasReservationData            *bool   `json:"HasReservationData,omitempty" xml:"HasReservationData,omitempty"`
	InstanceChargeType            *string `json:"InstanceChargeType,omitempty" xml:"InstanceChargeType,omitempty"`
	InternetChargeType            *string `json:"InternetChargeType,omitempty" xml:"InternetChargeType,omitempty"`
	Ipv6InternetBandwidthId       *string `json:"Ipv6InternetBandwidthId,omitempty" xml:"Ipv6InternetBandwidthId,omitempty"`
	ReservationActiveTime         *string `json:"ReservationActiveTime,omitempty" xml:"ReservationActiveTime,omitempty"`
	ReservationBandwidth          *int64  `json:"ReservationBandwidth,omitempty" xml:"ReservationBandwidth,omitempty"`
	ReservationInternetChargeType *string `json:"ReservationInternetChargeType,omitempty" xml:"ReservationInternetChargeType,omitempty"`
	ReservationOrderType          *string `json:"ReservationOrderType,omitempty" xml:"ReservationOrderType,omitempty"`
}

func (s DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6AddressIpv6InternetBandwidth) String() string {
	return dara.Prettify(s)
}

func (s DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6AddressIpv6InternetBandwidth) GoString() string {
	return s.String()
}

func (s *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6AddressIpv6InternetBandwidth) GetBandwidth() *int32 {
	return s.Bandwidth
}

func (s *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6AddressIpv6InternetBandwidth) GetBusinessStatus() *string {
	return s.BusinessStatus
}

func (s *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6AddressIpv6InternetBandwidth) GetHasReservationData() *bool {
	return s.HasReservationData
}

func (s *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6AddressIpv6InternetBandwidth) GetInstanceChargeType() *string {
	return s.InstanceChargeType
}

func (s *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6AddressIpv6InternetBandwidth) GetInternetChargeType() *string {
	return s.InternetChargeType
}

func (s *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6AddressIpv6InternetBandwidth) GetIpv6InternetBandwidthId() *string {
	return s.Ipv6InternetBandwidthId
}

func (s *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6AddressIpv6InternetBandwidth) GetReservationActiveTime() *string {
	return s.ReservationActiveTime
}

func (s *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6AddressIpv6InternetBandwidth) GetReservationBandwidth() *int64 {
	return s.ReservationBandwidth
}

func (s *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6AddressIpv6InternetBandwidth) GetReservationInternetChargeType() *string {
	return s.ReservationInternetChargeType
}

func (s *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6AddressIpv6InternetBandwidth) GetReservationOrderType() *string {
	return s.ReservationOrderType
}

func (s *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6AddressIpv6InternetBandwidth) SetBandwidth(v int32) *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6AddressIpv6InternetBandwidth {
	s.Bandwidth = &v
	return s
}

func (s *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6AddressIpv6InternetBandwidth) SetBusinessStatus(v string) *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6AddressIpv6InternetBandwidth {
	s.BusinessStatus = &v
	return s
}

func (s *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6AddressIpv6InternetBandwidth) SetHasReservationData(v bool) *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6AddressIpv6InternetBandwidth {
	s.HasReservationData = &v
	return s
}

func (s *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6AddressIpv6InternetBandwidth) SetInstanceChargeType(v string) *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6AddressIpv6InternetBandwidth {
	s.InstanceChargeType = &v
	return s
}

func (s *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6AddressIpv6InternetBandwidth) SetInternetChargeType(v string) *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6AddressIpv6InternetBandwidth {
	s.InternetChargeType = &v
	return s
}

func (s *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6AddressIpv6InternetBandwidth) SetIpv6InternetBandwidthId(v string) *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6AddressIpv6InternetBandwidth {
	s.Ipv6InternetBandwidthId = &v
	return s
}

func (s *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6AddressIpv6InternetBandwidth) SetReservationActiveTime(v string) *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6AddressIpv6InternetBandwidth {
	s.ReservationActiveTime = &v
	return s
}

func (s *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6AddressIpv6InternetBandwidth) SetReservationBandwidth(v int64) *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6AddressIpv6InternetBandwidth {
	s.ReservationBandwidth = &v
	return s
}

func (s *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6AddressIpv6InternetBandwidth) SetReservationInternetChargeType(v string) *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6AddressIpv6InternetBandwidth {
	s.ReservationInternetChargeType = &v
	return s
}

func (s *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6AddressIpv6InternetBandwidth) SetReservationOrderType(v string) *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6AddressIpv6InternetBandwidth {
	s.ReservationOrderType = &v
	return s
}

func (s *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6AddressIpv6InternetBandwidth) Validate() error {
	return dara.Validate(s)
}

type DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6AddressTags struct {
	Tag []*DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6AddressTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6AddressTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6AddressTags) GoString() string {
	return s.String()
}

func (s *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6AddressTags) GetTag() []*DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6AddressTagsTag {
	return s.Tag
}

func (s *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6AddressTags) SetTag(v []*DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6AddressTagsTag) *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6AddressTags {
	s.Tag = v
	return s
}

func (s *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6AddressTags) Validate() error {
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6AddressTagsTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6AddressTagsTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6AddressTagsTag) GoString() string {
	return s.String()
}

func (s *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6AddressTagsTag) GetKey() *string {
	return s.Key
}

func (s *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6AddressTagsTag) GetValue() *string {
	return s.Value
}

func (s *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6AddressTagsTag) SetKey(v string) *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6AddressTagsTag {
	s.Key = &v
	return s
}

func (s *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6AddressTagsTag) SetValue(v string) *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6AddressTagsTag {
	s.Value = &v
	return s
}

func (s *DescribeIpv6AddressesResponseBodyIpv6AddressesIpv6AddressTagsTag) Validate() error {
	return dara.Validate(s)
}
