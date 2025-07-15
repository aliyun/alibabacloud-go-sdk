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
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The details of the router interface.
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
	return dara.Validate(s)
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
	return dara.Validate(s)
}

type DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType struct {
	// The ID of the access point.
	//
	// example:
	//
	// ap-cn-shanghaiSZ-****
	AccessPointId *string `json:"AccessPointId,omitempty" xml:"AccessPointId,omitempty"`
	// The bandwidth of the router interface. Unit: Mbit/s.
	//
	// example:
	//
	// 10
	Bandwidth *int32 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The service status of the router interface. Valid values:
	//
	// 	- **Normal**
	//
	// 	- **FinancialLocked**
	//
	// 	- **SecurityLocked**
	//
	// example:
	//
	// Normal
	BusinessStatus *string `json:"BusinessStatus,omitempty" xml:"BusinessStatus,omitempty"`
	// The billing method.
	//
	// example:
	//
	// PayByTraffic
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// The time when the connection was established.
	//
	// The time follows the ISO8601 standard in the `YYYY-MM-DDThh:mmZ` format. The time is displayed in UTC.
	//
	// example:
	//
	// 2021-06-08T12:20:55
	ConnectedTime *string `json:"ConnectedTime,omitempty" xml:"ConnectedTime,omitempty"`
	// The time when the route table was created.
	//
	// The time follows the ISO8601 standard in the `YYYY-MM-DDThh:mmZ` format. The time is displayed in UTC.
	//
	// example:
	//
	// 2021-06-08T12:20:55
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// Indicates whether the connection is a cross-border connection.
	//
	// example:
	//
	// false
	CrossBorder *bool `json:"CrossBorder,omitempty" xml:"CrossBorder,omitempty"`
	// The description of the router interface.
	//
	// example:
	//
	// The description of the router interface.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The end of the time range during which data was queried.
	//
	// The time follows the ISO8601 standard in the `YYYY-MM-DDThh:mmZ` format. The time is displayed in UTC.
	//
	// example:
	//
	// 2021-06-08T12:20:55
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// Indicates whether the VBR that is created in the Fast Link mode is uplinked to the router interface. The Fast Link mode helps automatically connect router interfaces that are created for the VBR and its peer VPC. Valid values:
	//
	// 	- **true**
	//
	// 	- **false*	- (default)
	//
	// example:
	//
	// false
	FastLinkMode *bool `json:"FastLinkMode,omitempty" xml:"FastLinkMode,omitempty"`
	// Indicates whether renewal data is included.
	//
	// example:
	//
	// false
	HasReservationData *string `json:"HasReservationData,omitempty" xml:"HasReservationData,omitempty"`
	// The rate of heath checks.
	//
	// example:
	//
	// 1
	HcRate *int32 `json:"HcRate,omitempty" xml:"HcRate,omitempty"`
	// The health check threshold.
	//
	// example:
	//
	// 2
	HcThreshold *int32 `json:"HcThreshold,omitempty" xml:"HcThreshold,omitempty"`
	// The source IP address that is used for the health check.
	//
	// example:
	//
	// 116.62.XX.XX
	HealthCheckSourceIp *string `json:"HealthCheckSourceIp,omitempty" xml:"HealthCheckSourceIp,omitempty"`
	// The destination IP address that is used for the health check.
	//
	// example:
	//
	// 116.62.XX.XX
	HealthCheckTargetIp *string `json:"HealthCheckTargetIp,omitempty" xml:"HealthCheckTargetIp,omitempty"`
	// Indicates whether protection against malicious IPv6 traffic is enabled. Valid values:
	//
	// 	- **on**
	//
	// 	- **off**
	//
	// 	- **unsupport**
	//
	// example:
	//
	// on
	Ipv6Status *string `json:"Ipv6Status,omitempty" xml:"Ipv6Status,omitempty"`
	// The custom name.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the peer access point.
	//
	// example:
	//
	// ap-cn-shanghaiSZ-****
	OppositeAccessPointId *string `json:"OppositeAccessPointId,omitempty" xml:"OppositeAccessPointId,omitempty"`
	// The maximum bandwidth of the peer router interface. Unit: Mbit/s.
	//
	// example:
	//
	// 12
	OppositeBandwidth *int32 `json:"OppositeBandwidth,omitempty" xml:"OppositeBandwidth,omitempty"`
	// The service status of the peer router interface.
	//
	// example:
	//
	// Normal
	OppositeInterfaceBusinessStatus *string `json:"OppositeInterfaceBusinessStatus,omitempty" xml:"OppositeInterfaceBusinessStatus,omitempty"`
	// The ID of the peer router interface.
	//
	// example:
	//
	// ri-bp1itx13bwe6f2wfh****
	OppositeInterfaceId *string `json:"OppositeInterfaceId,omitempty" xml:"OppositeInterfaceId,omitempty"`
	// The ID of the Alibaba Cloud account to which the peer router interface belongs.
	//
	// example:
	//
	// 271598332402530847
	OppositeInterfaceOwnerId *string `json:"OppositeInterfaceOwnerId,omitempty" xml:"OppositeInterfaceOwnerId,omitempty"`
	// The specification of the peer router interface.
	//
	// example:
	//
	// Large
	OppositeInterfaceSpec *string `json:"OppositeInterfaceSpec,omitempty" xml:"OppositeInterfaceSpec,omitempty"`
	// The status of the peer router interface.
	//
	// example:
	//
	// Normal
	OppositeInterfaceStatus *string `json:"OppositeInterfaceStatus,omitempty" xml:"OppositeInterfaceStatus,omitempty"`
	// The region ID of the peer router interface.
	//
	// example:
	//
	// cn-shanghai
	OppositeRegionId *string `json:"OppositeRegionId,omitempty" xml:"OppositeRegionId,omitempty"`
	// The ID of the router to which the peer router interface belongs.
	//
	// example:
	//
	// vrt-bp1d3bxtdv68tfd7g****
	OppositeRouterId *string `json:"OppositeRouterId,omitempty" xml:"OppositeRouterId,omitempty"`
	// The type of the router to which the peer router interface belongs.
	//
	// example:
	//
	// VRouter
	OppositeRouterType *string `json:"OppositeRouterType,omitempty" xml:"OppositeRouterType,omitempty"`
	// The ID of the peer VPC.
	//
	// example:
	//
	// vpc-bp1qpo0kug3a20qqe****
	OppositeVpcInstanceId *string `json:"OppositeVpcInstanceId,omitempty" xml:"OppositeVpcInstanceId,omitempty"`
	// The time when the renewal takes effect.
	//
	// The time follows the ISO8601 standard in the `YYYY-MM-DDThh:mmZ` format. The time is displayed in UTC.
	//
	// example:
	//
	// 2021-03-11T16:00:00Z
	ReservationActiveTime *string `json:"ReservationActiveTime,omitempty" xml:"ReservationActiveTime,omitempty"`
	// The maximum bandwidth after the renewal takes effect. Unit: Mbit/s.
	//
	// example:
	//
	// 10
	ReservationBandwidth *string `json:"ReservationBandwidth,omitempty" xml:"ReservationBandwidth,omitempty"`
	// The metering method that is used after the renewal takes effect. Valid values:
	//
	// example:
	//
	// PayByBandwidth
	ReservationInternetChargeType *string `json:"ReservationInternetChargeType,omitempty" xml:"ReservationInternetChargeType,omitempty"`
	// The type of the renewal order. Valid values:
	//
	// example:
	//
	// RENEWCHANGE
	ReservationOrderType *string `json:"ReservationOrderType,omitempty" xml:"ReservationOrderType,omitempty"`
	// Resource Group ID.
	//
	// For more information about resource groups, please refer to [What is a Resource Group?](https://help.aliyun.com/document_detail/94475.html)
	//
	// example:
	//
	// rg-acfmxazb4ph6aiy****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// Indicates whether the router interface is the initiator or acceptor of the peering connection.
	//
	// example:
	//
	// InitiatingSide
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
	// The ID of the router to which the route entry belongs.
	//
	// example:
	//
	// vrt-bp1d3bxtdv68tfd7g****
	RouterId *string `json:"RouterId,omitempty" xml:"RouterId,omitempty"`
	// The ID of the router interface.
	//
	// example:
	//
	// ri-2zenfgfpyu3v93koa****
	RouterInterfaceId *string `json:"RouterInterfaceId,omitempty" xml:"RouterInterfaceId,omitempty"`
	// The type of the router to which the route table belongs. Valid values:
	//
	// 	- **VRouter**
	//
	// 	- **VBR**
	//
	// example:
	//
	// VRouter
	RouterType *string `json:"RouterType,omitempty" xml:"RouterType,omitempty"`
	// The specification of the router interface.
	//
	// example:
	//
	// Large
	Spec *string `json:"Spec,omitempty" xml:"Spec,omitempty"`
	// The status of the router interface.
	//
	// example:
	//
	// active
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tags of the resource.
	Tags *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceTypeTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	// The ID of the local virtual private cloud (VPC) in the peering connection.
	//
	// example:
	//
	// vpc-2ze3tq4uxhysg717x****
	VpcInstanceId *string `json:"VpcInstanceId,omitempty" xml:"VpcInstanceId,omitempty"`
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
	return dara.Validate(s)
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
	return dara.Validate(s)
}

type DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceTypeTagsTags struct {
	// The key of the resource tag. At least one tag key must be entered, and a maximum of 20 tag keys are supported. If this value needs to be passed in, it cannot be an empty string.
	//
	// A tag key can support up to 128 characters, cannot start with \\"aliyun\\" or \\"acs:\\", and cannot contain \\"http://\\" or \\"https://\\".
	//
	// example:
	//
	// FinanceDept
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the resource tag. A maximum of 20 tag values can be entered. If this value needs to be passed in, an empty string can be entered.
	//
	// A maximum of 128 characters are supported, it cannot start with \\"aliyun\\" or \\"acs:\\", and it cannot contain \\"http://\\" or \\"https://\\".
	//
	// example:
	//
	// FinanceJoshua
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
