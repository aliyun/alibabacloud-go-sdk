// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRouterInterfacesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *DescribeRouterInterfacesResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeRouterInterfacesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeRouterInterfacesResponseBody
	GetRequestId() *string
	SetRouterInterfaceSet(v *DescribeRouterInterfacesResponseBodyRouterInterfaceSet) *DescribeRouterInterfacesResponseBody
	GetRouterInterfaceSet() *DescribeRouterInterfacesResponseBodyRouterInterfaceSet
	SetTotalCount(v int32) *DescribeRouterInterfacesResponseBody
	GetTotalCount() *int32
}

type DescribeRouterInterfacesResponseBody struct {
	// The page number. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Maximum value: **50**. Default value: **10**.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// C7F6FCBD-F9CC-4501-8EF3-CDC9577CAE45
	RequestId          *string                                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RouterInterfaceSet *DescribeRouterInterfacesResponseBodyRouterInterfaceSet `json:"RouterInterfaceSet,omitempty" xml:"RouterInterfaceSet,omitempty" type:"Struct"`
	// The number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeRouterInterfacesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRouterInterfacesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRouterInterfacesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeRouterInterfacesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeRouterInterfacesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRouterInterfacesResponseBody) GetRouterInterfaceSet() *DescribeRouterInterfacesResponseBodyRouterInterfaceSet {
	return s.RouterInterfaceSet
}

func (s *DescribeRouterInterfacesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeRouterInterfacesResponseBody) SetPageNumber(v int32) *DescribeRouterInterfacesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeRouterInterfacesResponseBody) SetPageSize(v int32) *DescribeRouterInterfacesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeRouterInterfacesResponseBody) SetRequestId(v string) *DescribeRouterInterfacesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRouterInterfacesResponseBody) SetRouterInterfaceSet(v *DescribeRouterInterfacesResponseBodyRouterInterfaceSet) *DescribeRouterInterfacesResponseBody {
	s.RouterInterfaceSet = v
	return s
}

func (s *DescribeRouterInterfacesResponseBody) SetTotalCount(v int32) *DescribeRouterInterfacesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeRouterInterfacesResponseBody) Validate() error {
	if s.RouterInterfaceSet != nil {
		if err := s.RouterInterfaceSet.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeRouterInterfacesResponseBodyRouterInterfaceSet struct {
	RouterInterfaceType []*DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType `json:"RouterInterfaceType,omitempty" xml:"RouterInterfaceType,omitempty" type:"Repeated"`
}

func (s DescribeRouterInterfacesResponseBodyRouterInterfaceSet) String() string {
	return dara.Prettify(s)
}

func (s DescribeRouterInterfacesResponseBodyRouterInterfaceSet) GoString() string {
	return s.String()
}

func (s *DescribeRouterInterfacesResponseBodyRouterInterfaceSet) GetRouterInterfaceType() []*DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType {
	return s.RouterInterfaceType
}

func (s *DescribeRouterInterfacesResponseBodyRouterInterfaceSet) SetRouterInterfaceType(v []*DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType) *DescribeRouterInterfacesResponseBodyRouterInterfaceSet {
	s.RouterInterfaceType = v
	return s
}

func (s *DescribeRouterInterfacesResponseBodyRouterInterfaceSet) Validate() error {
	if s.RouterInterfaceType != nil {
		for _, item := range s.RouterInterfaceType {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType struct {
	AccessPointId                   *string                                                                        `json:"AccessPointId,omitempty" xml:"AccessPointId,omitempty"`
	Bandwidth                       *int32                                                                         `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	BusinessStatus                  *string                                                                        `json:"BusinessStatus,omitempty" xml:"BusinessStatus,omitempty"`
	ChargeType                      *string                                                                        `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	ConnectedTime                   *string                                                                        `json:"ConnectedTime,omitempty" xml:"ConnectedTime,omitempty"`
	CreationTime                    *string                                                                        `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	CrossBorder                     *bool                                                                          `json:"CrossBorder,omitempty" xml:"CrossBorder,omitempty"`
	Description                     *string                                                                        `json:"Description,omitempty" xml:"Description,omitempty"`
	EndTime                         *string                                                                        `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	FastLinkMode                    *bool                                                                          `json:"FastLinkMode,omitempty" xml:"FastLinkMode,omitempty"`
	HasReservationData              *string                                                                        `json:"HasReservationData,omitempty" xml:"HasReservationData,omitempty"`
	HcRate                          *int32                                                                         `json:"HcRate,omitempty" xml:"HcRate,omitempty"`
	HcThreshold                     *int32                                                                         `json:"HcThreshold,omitempty" xml:"HcThreshold,omitempty"`
	HealthCheckSourceIp             *string                                                                        `json:"HealthCheckSourceIp,omitempty" xml:"HealthCheckSourceIp,omitempty"`
	HealthCheckTargetIp             *string                                                                        `json:"HealthCheckTargetIp,omitempty" xml:"HealthCheckTargetIp,omitempty"`
	Ipv6Status                      *string                                                                        `json:"Ipv6Status,omitempty" xml:"Ipv6Status,omitempty"`
	Name                            *string                                                                        `json:"Name,omitempty" xml:"Name,omitempty"`
	OppositeAccessPointId           *string                                                                        `json:"OppositeAccessPointId,omitempty" xml:"OppositeAccessPointId,omitempty"`
	OppositeBandwidth               *int32                                                                         `json:"OppositeBandwidth,omitempty" xml:"OppositeBandwidth,omitempty"`
	OppositeInterfaceBusinessStatus *string                                                                        `json:"OppositeInterfaceBusinessStatus,omitempty" xml:"OppositeInterfaceBusinessStatus,omitempty"`
	OppositeInterfaceId             *string                                                                        `json:"OppositeInterfaceId,omitempty" xml:"OppositeInterfaceId,omitempty"`
	OppositeInterfaceOwnerId        *string                                                                        `json:"OppositeInterfaceOwnerId,omitempty" xml:"OppositeInterfaceOwnerId,omitempty"`
	OppositeInterfaceSpec           *string                                                                        `json:"OppositeInterfaceSpec,omitempty" xml:"OppositeInterfaceSpec,omitempty"`
	OppositeInterfaceStatus         *string                                                                        `json:"OppositeInterfaceStatus,omitempty" xml:"OppositeInterfaceStatus,omitempty"`
	OppositeRegionId                *string                                                                        `json:"OppositeRegionId,omitempty" xml:"OppositeRegionId,omitempty"`
	OppositeRouterId                *string                                                                        `json:"OppositeRouterId,omitempty" xml:"OppositeRouterId,omitempty"`
	OppositeRouterType              *string                                                                        `json:"OppositeRouterType,omitempty" xml:"OppositeRouterType,omitempty"`
	OppositeVpcInstanceId           *string                                                                        `json:"OppositeVpcInstanceId,omitempty" xml:"OppositeVpcInstanceId,omitempty"`
	ReservationActiveTime           *string                                                                        `json:"ReservationActiveTime,omitempty" xml:"ReservationActiveTime,omitempty"`
	ReservationBandwidth            *string                                                                        `json:"ReservationBandwidth,omitempty" xml:"ReservationBandwidth,omitempty"`
	ReservationInternetChargeType   *string                                                                        `json:"ReservationInternetChargeType,omitempty" xml:"ReservationInternetChargeType,omitempty"`
	ReservationOrderType            *string                                                                        `json:"ReservationOrderType,omitempty" xml:"ReservationOrderType,omitempty"`
	ResourceGroupId                 *string                                                                        `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Role                            *string                                                                        `json:"Role,omitempty" xml:"Role,omitempty"`
	RouterId                        *string                                                                        `json:"RouterId,omitempty" xml:"RouterId,omitempty"`
	RouterInterfaceId               *string                                                                        `json:"RouterInterfaceId,omitempty" xml:"RouterInterfaceId,omitempty"`
	RouterType                      *string                                                                        `json:"RouterType,omitempty" xml:"RouterType,omitempty"`
	Spec                            *string                                                                        `json:"Spec,omitempty" xml:"Spec,omitempty"`
	Status                          *string                                                                        `json:"Status,omitempty" xml:"Status,omitempty"`
	Tags                            *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceTypeTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	VpcInstanceId                   *string                                                                        `json:"VpcInstanceId,omitempty" xml:"VpcInstanceId,omitempty"`
}

func (s DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType) String() string {
	return dara.Prettify(s)
}

func (s DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType) GoString() string {
	return s.String()
}

func (s *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType) GetAccessPointId() *string {
	return s.AccessPointId
}

func (s *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType) GetBandwidth() *int32 {
	return s.Bandwidth
}

func (s *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType) GetBusinessStatus() *string {
	return s.BusinessStatus
}

func (s *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType) GetChargeType() *string {
	return s.ChargeType
}

func (s *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType) GetConnectedTime() *string {
	return s.ConnectedTime
}

func (s *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType) GetCrossBorder() *bool {
	return s.CrossBorder
}

func (s *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType) GetDescription() *string {
	return s.Description
}

func (s *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType) GetFastLinkMode() *bool {
	return s.FastLinkMode
}

func (s *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType) GetHasReservationData() *string {
	return s.HasReservationData
}

func (s *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType) GetHcRate() *int32 {
	return s.HcRate
}

func (s *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType) GetHcThreshold() *int32 {
	return s.HcThreshold
}

func (s *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType) GetHealthCheckSourceIp() *string {
	return s.HealthCheckSourceIp
}

func (s *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType) GetHealthCheckTargetIp() *string {
	return s.HealthCheckTargetIp
}

func (s *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType) GetIpv6Status() *string {
	return s.Ipv6Status
}

func (s *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType) GetName() *string {
	return s.Name
}

func (s *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType) GetOppositeAccessPointId() *string {
	return s.OppositeAccessPointId
}

func (s *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType) GetOppositeBandwidth() *int32 {
	return s.OppositeBandwidth
}

func (s *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType) GetOppositeInterfaceBusinessStatus() *string {
	return s.OppositeInterfaceBusinessStatus
}

func (s *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType) GetOppositeInterfaceId() *string {
	return s.OppositeInterfaceId
}

func (s *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType) GetOppositeInterfaceOwnerId() *string {
	return s.OppositeInterfaceOwnerId
}

func (s *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType) GetOppositeInterfaceSpec() *string {
	return s.OppositeInterfaceSpec
}

func (s *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType) GetOppositeInterfaceStatus() *string {
	return s.OppositeInterfaceStatus
}

func (s *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType) GetOppositeRegionId() *string {
	return s.OppositeRegionId
}

func (s *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType) GetOppositeRouterId() *string {
	return s.OppositeRouterId
}

func (s *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType) GetOppositeRouterType() *string {
	return s.OppositeRouterType
}

func (s *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType) GetOppositeVpcInstanceId() *string {
	return s.OppositeVpcInstanceId
}

func (s *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType) GetReservationActiveTime() *string {
	return s.ReservationActiveTime
}

func (s *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType) GetReservationBandwidth() *string {
	return s.ReservationBandwidth
}

func (s *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType) GetReservationInternetChargeType() *string {
	return s.ReservationInternetChargeType
}

func (s *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType) GetReservationOrderType() *string {
	return s.ReservationOrderType
}

func (s *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType) GetRole() *string {
	return s.Role
}

func (s *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType) GetRouterId() *string {
	return s.RouterId
}

func (s *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType) GetRouterInterfaceId() *string {
	return s.RouterInterfaceId
}

func (s *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType) GetRouterType() *string {
	return s.RouterType
}

func (s *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType) GetSpec() *string {
	return s.Spec
}

func (s *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType) GetStatus() *string {
	return s.Status
}

func (s *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType) GetTags() *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceTypeTags {
	return s.Tags
}

func (s *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType) GetVpcInstanceId() *string {
	return s.VpcInstanceId
}

func (s *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType) SetAccessPointId(v string) *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType {
	s.AccessPointId = &v
	return s
}

func (s *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType) SetBandwidth(v int32) *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType {
	s.Bandwidth = &v
	return s
}

func (s *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType) SetBusinessStatus(v string) *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType {
	s.BusinessStatus = &v
	return s
}

func (s *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType) SetChargeType(v string) *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType {
	s.ChargeType = &v
	return s
}

func (s *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType) SetConnectedTime(v string) *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType {
	s.ConnectedTime = &v
	return s
}

func (s *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType) SetCreationTime(v string) *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType {
	s.CreationTime = &v
	return s
}

func (s *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType) SetCrossBorder(v bool) *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType {
	s.CrossBorder = &v
	return s
}

func (s *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType) SetDescription(v string) *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType {
	s.Description = &v
	return s
}

func (s *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType) SetEndTime(v string) *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType {
	s.EndTime = &v
	return s
}

func (s *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType) SetFastLinkMode(v bool) *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType {
	s.FastLinkMode = &v
	return s
}

func (s *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType) SetHasReservationData(v string) *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType {
	s.HasReservationData = &v
	return s
}

func (s *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType) SetHcRate(v int32) *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType {
	s.HcRate = &v
	return s
}

func (s *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType) SetHcThreshold(v int32) *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType {
	s.HcThreshold = &v
	return s
}

func (s *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType) SetHealthCheckSourceIp(v string) *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType {
	s.HealthCheckSourceIp = &v
	return s
}

func (s *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType) SetHealthCheckTargetIp(v string) *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType {
	s.HealthCheckTargetIp = &v
	return s
}

func (s *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType) SetIpv6Status(v string) *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType {
	s.Ipv6Status = &v
	return s
}

func (s *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType) SetName(v string) *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType {
	s.Name = &v
	return s
}

func (s *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType) SetOppositeAccessPointId(v string) *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType {
	s.OppositeAccessPointId = &v
	return s
}

func (s *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType) SetOppositeBandwidth(v int32) *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType {
	s.OppositeBandwidth = &v
	return s
}

func (s *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType) SetOppositeInterfaceBusinessStatus(v string) *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType {
	s.OppositeInterfaceBusinessStatus = &v
	return s
}

func (s *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType) SetOppositeInterfaceId(v string) *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType {
	s.OppositeInterfaceId = &v
	return s
}

func (s *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType) SetOppositeInterfaceOwnerId(v string) *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType {
	s.OppositeInterfaceOwnerId = &v
	return s
}

func (s *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType) SetOppositeInterfaceSpec(v string) *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType {
	s.OppositeInterfaceSpec = &v
	return s
}

func (s *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType) SetOppositeInterfaceStatus(v string) *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType {
	s.OppositeInterfaceStatus = &v
	return s
}

func (s *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType) SetOppositeRegionId(v string) *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType {
	s.OppositeRegionId = &v
	return s
}

func (s *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType) SetOppositeRouterId(v string) *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType {
	s.OppositeRouterId = &v
	return s
}

func (s *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType) SetOppositeRouterType(v string) *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType {
	s.OppositeRouterType = &v
	return s
}

func (s *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType) SetOppositeVpcInstanceId(v string) *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType {
	s.OppositeVpcInstanceId = &v
	return s
}

func (s *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType) SetReservationActiveTime(v string) *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType {
	s.ReservationActiveTime = &v
	return s
}

func (s *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType) SetReservationBandwidth(v string) *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType {
	s.ReservationBandwidth = &v
	return s
}

func (s *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType) SetReservationInternetChargeType(v string) *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType {
	s.ReservationInternetChargeType = &v
	return s
}

func (s *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType) SetReservationOrderType(v string) *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType {
	s.ReservationOrderType = &v
	return s
}

func (s *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType) SetResourceGroupId(v string) *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType) SetRole(v string) *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType {
	s.Role = &v
	return s
}

func (s *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType) SetRouterId(v string) *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType {
	s.RouterId = &v
	return s
}

func (s *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType) SetRouterInterfaceId(v string) *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType {
	s.RouterInterfaceId = &v
	return s
}

func (s *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType) SetRouterType(v string) *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType {
	s.RouterType = &v
	return s
}

func (s *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType) SetSpec(v string) *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType {
	s.Spec = &v
	return s
}

func (s *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType) SetStatus(v string) *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType {
	s.Status = &v
	return s
}

func (s *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType) SetTags(v *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceTypeTags) *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType {
	s.Tags = v
	return s
}

func (s *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType) SetVpcInstanceId(v string) *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType {
	s.VpcInstanceId = &v
	return s
}

func (s *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType) Validate() error {
	if s.Tags != nil {
		if err := s.Tags.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceTypeTags struct {
	Tags []*DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceTypeTagsTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceTypeTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceTypeTags) GoString() string {
	return s.String()
}

func (s *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceTypeTags) GetTags() []*DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceTypeTagsTags {
	return s.Tags
}

func (s *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceTypeTags) SetTags(v []*DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceTypeTagsTags) *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceTypeTags {
	s.Tags = v
	return s
}

func (s *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceTypeTags) Validate() error {
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceTypeTagsTags struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceTypeTagsTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceTypeTagsTags) GoString() string {
	return s.String()
}

func (s *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceTypeTagsTags) GetKey() *string {
	return s.Key
}

func (s *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceTypeTagsTags) GetValue() *string {
	return s.Value
}

func (s *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceTypeTagsTags) SetKey(v string) *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceTypeTagsTags {
	s.Key = &v
	return s
}

func (s *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceTypeTagsTags) SetValue(v string) *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceTypeTagsTags {
	s.Value = &v
	return s
}

func (s *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceTypeTagsTags) Validate() error {
	return dara.Validate(s)
}
