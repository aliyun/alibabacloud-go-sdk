// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRouterInterfaceAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessPointId(v string) *DescribeRouterInterfaceAttributeResponseBody
	GetAccessPointId() *string
	SetBandwidth(v int32) *DescribeRouterInterfaceAttributeResponseBody
	GetBandwidth() *int32
	SetBusinessStatus(v string) *DescribeRouterInterfaceAttributeResponseBody
	GetBusinessStatus() *string
	SetChargeType(v string) *DescribeRouterInterfaceAttributeResponseBody
	GetChargeType() *string
	SetCode(v string) *DescribeRouterInterfaceAttributeResponseBody
	GetCode() *string
	SetConnectedTime(v string) *DescribeRouterInterfaceAttributeResponseBody
	GetConnectedTime() *string
	SetCreationTime(v string) *DescribeRouterInterfaceAttributeResponseBody
	GetCreationTime() *string
	SetCrossBorder(v bool) *DescribeRouterInterfaceAttributeResponseBody
	GetCrossBorder() *bool
	SetDescription(v string) *DescribeRouterInterfaceAttributeResponseBody
	GetDescription() *string
	SetEndTime(v string) *DescribeRouterInterfaceAttributeResponseBody
	GetEndTime() *string
	SetFastLinkMode(v string) *DescribeRouterInterfaceAttributeResponseBody
	GetFastLinkMode() *string
	SetGmtModified(v string) *DescribeRouterInterfaceAttributeResponseBody
	GetGmtModified() *string
	SetHasReservationData(v string) *DescribeRouterInterfaceAttributeResponseBody
	GetHasReservationData() *string
	SetHcRate(v int32) *DescribeRouterInterfaceAttributeResponseBody
	GetHcRate() *int32
	SetHcThreshold(v int32) *DescribeRouterInterfaceAttributeResponseBody
	GetHcThreshold() *int32
	SetHealthCheckSourceIp(v string) *DescribeRouterInterfaceAttributeResponseBody
	GetHealthCheckSourceIp() *string
	SetHealthCheckStatus(v string) *DescribeRouterInterfaceAttributeResponseBody
	GetHealthCheckStatus() *string
	SetHealthCheckTargetIp(v string) *DescribeRouterInterfaceAttributeResponseBody
	GetHealthCheckTargetIp() *string
	SetMessage(v string) *DescribeRouterInterfaceAttributeResponseBody
	GetMessage() *string
	SetName(v string) *DescribeRouterInterfaceAttributeResponseBody
	GetName() *string
	SetOppositeAccessPointId(v string) *DescribeRouterInterfaceAttributeResponseBody
	GetOppositeAccessPointId() *string
	SetOppositeBandwidth(v int32) *DescribeRouterInterfaceAttributeResponseBody
	GetOppositeBandwidth() *int32
	SetOppositeInterfaceBusinessStatus(v string) *DescribeRouterInterfaceAttributeResponseBody
	GetOppositeInterfaceBusinessStatus() *string
	SetOppositeInterfaceId(v string) *DescribeRouterInterfaceAttributeResponseBody
	GetOppositeInterfaceId() *string
	SetOppositeInterfaceOwnerId(v string) *DescribeRouterInterfaceAttributeResponseBody
	GetOppositeInterfaceOwnerId() *string
	SetOppositeInterfaceSpec(v string) *DescribeRouterInterfaceAttributeResponseBody
	GetOppositeInterfaceSpec() *string
	SetOppositeInterfaceStatus(v string) *DescribeRouterInterfaceAttributeResponseBody
	GetOppositeInterfaceStatus() *string
	SetOppositeRegionId(v string) *DescribeRouterInterfaceAttributeResponseBody
	GetOppositeRegionId() *string
	SetOppositeRouterId(v string) *DescribeRouterInterfaceAttributeResponseBody
	GetOppositeRouterId() *string
	SetOppositeRouterType(v string) *DescribeRouterInterfaceAttributeResponseBody
	GetOppositeRouterType() *string
	SetOppositeVpcInstanceId(v string) *DescribeRouterInterfaceAttributeResponseBody
	GetOppositeVpcInstanceId() *string
	SetRequestId(v string) *DescribeRouterInterfaceAttributeResponseBody
	GetRequestId() *string
	SetReservationActiveTime(v string) *DescribeRouterInterfaceAttributeResponseBody
	GetReservationActiveTime() *string
	SetReservationBandwidth(v string) *DescribeRouterInterfaceAttributeResponseBody
	GetReservationBandwidth() *string
	SetReservationInternetChargeType(v string) *DescribeRouterInterfaceAttributeResponseBody
	GetReservationInternetChargeType() *string
	SetReservationOrderType(v string) *DescribeRouterInterfaceAttributeResponseBody
	GetReservationOrderType() *string
	SetResourceGroupId(v string) *DescribeRouterInterfaceAttributeResponseBody
	GetResourceGroupId() *string
	SetRole(v string) *DescribeRouterInterfaceAttributeResponseBody
	GetRole() *string
	SetRouterId(v string) *DescribeRouterInterfaceAttributeResponseBody
	GetRouterId() *string
	SetRouterInterfaceId(v string) *DescribeRouterInterfaceAttributeResponseBody
	GetRouterInterfaceId() *string
	SetRouterType(v string) *DescribeRouterInterfaceAttributeResponseBody
	GetRouterType() *string
	SetSpec(v string) *DescribeRouterInterfaceAttributeResponseBody
	GetSpec() *string
	SetStatus(v string) *DescribeRouterInterfaceAttributeResponseBody
	GetStatus() *string
	SetSuccess(v bool) *DescribeRouterInterfaceAttributeResponseBody
	GetSuccess() *bool
	SetTags(v *DescribeRouterInterfaceAttributeResponseBodyTags) *DescribeRouterInterfaceAttributeResponseBody
	GetTags() *DescribeRouterInterfaceAttributeResponseBodyTags
	SetVpcInstanceId(v string) *DescribeRouterInterfaceAttributeResponseBody
	GetVpcInstanceId() *string
}

type DescribeRouterInterfaceAttributeResponseBody struct {
	// The ID of the access point.
	//
	// example:
	//
	// ap-cn-qingdao-ls-A
	AccessPointId *string `json:"AccessPointId,omitempty" xml:"AccessPointId,omitempty"`
	// The bandwidth of the router interface. Unit: Mbit/s.
	//
	// example:
	//
	// 2
	Bandwidth *int32 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The status of the router interface. Valid values:
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
	// The billing method. Valid values:
	//
	// 	- **AfterPay**: pay-as-you-go
	//
	// 	- **PrePaid**: subscription
	//
	// example:
	//
	// AfterPay
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The time when the connection was established.
	//
	// example:
	//
	// 2022-04-14T08:58:04Z
	ConnectedTime *string `json:"ConnectedTime,omitempty" xml:"ConnectedTime,omitempty"`
	// The time when the router interface was created.
	//
	// example:
	//
	// 2022-04-14T08:57:24Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// Indicates whether the connection is a cross-border connection. Valid values:
	//
	// 	- **false**
	//
	// 	- **true**
	//
	// example:
	//
	// false
	CrossBorder *bool `json:"CrossBorder,omitempty" xml:"CrossBorder,omitempty"`
	// The description of the router interface.
	//
	// example:
	//
	// Peer interface.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The end of the time range during which data was queried.
	//
	// example:
	//
	// 2999-09-08T16:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// Indicates whether the VBR that is created in the Fast Link mode is uplinked to the router interface. The Fast Link mode helps automatically connect router interfaces that are created for the VBR and its peer VPC. Valid values:
	//
	// 	- **true**
	//
	// 	- **false*	- (default)
	//
	// >
	//
	// 	- This parameter takes effect only when **RouterType*	- is set to **VBR*	- and **OppositeRouterType*	- is set to **VRouter**.
	//
	// 	- When **FastLinkMode*	- is set to **true**, **Role*	- must be set to **InitiatingSide**. **AccessPointId**, **OppositeRouterType**, **OpppsiteRouterId**, and **OppositeInterfaceOwnerId*	- are required.
	//
	// example:
	//
	// false
	FastLinkMode *string `json:"FastLinkMode,omitempty" xml:"FastLinkMode,omitempty"`
	// The time when the router interface was modified.
	//
	// example:
	//
	// 2022-04-28T10:02:12Z
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// Indicates whether renewal data is included. Valid values:
	//
	// 	- **false**
	//
	// 	- **true**
	//
	// example:
	//
	// false
	HasReservationData *string `json:"HasReservationData,omitempty" xml:"HasReservationData,omitempty"`
	// The rate of health checks. Unit: seconds. The value indicates the interval at which probe packets are sent during a health check.
	//
	// example:
	//
	// 2
	HcRate *int32 `json:"HcRate,omitempty" xml:"HcRate,omitempty"`
	// The healthy threshold. This value indicates the number of probe packets that are sent during a health check. Unit: packets.
	//
	// example:
	//
	// 8
	HcThreshold *int32 `json:"HcThreshold,omitempty" xml:"HcThreshold,omitempty"`
	// The source IP address that is used for the health check.
	//
	// example:
	//
	// 1.1.XX.XX
	HealthCheckSourceIp *string `json:"HealthCheckSourceIp,omitempty" xml:"HealthCheckSourceIp,omitempty"`
	// The status of the health check. Valid values:
	//
	// 	- **Abnormal**
	//
	// 	- **Normal**
	//
	// 	- **NoRedundantRoute**
	//
	// 	- **NoHealthCheckConfig**
	//
	// example:
	//
	// normal
	HealthCheckStatus *string `json:"HealthCheckStatus,omitempty" xml:"HealthCheckStatus,omitempty"`
	// The destination IP address that is used for the health check.
	//
	// example:
	//
	// 2.2.XX.XX
	HealthCheckTargetIp *string `json:"HealthCheckTargetIp,omitempty" xml:"HealthCheckTargetIp,omitempty"`
	// The response parameters.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The name of the router interface.
	//
	// example:
	//
	// RouterInterface1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the peer access point.
	//
	// example:
	//
	// ap-cn-qingdao-ls-B
	OppositeAccessPointId *string `json:"OppositeAccessPointId,omitempty" xml:"OppositeAccessPointId,omitempty"`
	// The maximum bandwidth of the peer router interface. Unit: Mbit/s.
	//
	// example:
	//
	// 0
	OppositeBandwidth *int32 `json:"OppositeBandwidth,omitempty" xml:"OppositeBandwidth,omitempty"`
	// The service status of the peer router interface. Valid values:
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
	OppositeInterfaceBusinessStatus *string `json:"OppositeInterfaceBusinessStatus,omitempty" xml:"OppositeInterfaceBusinessStatus,omitempty"`
	// The ID of the peer router interface.
	//
	// example:
	//
	// ri-bp1xkrzttximaoxbl****
	OppositeInterfaceId *string `json:"OppositeInterfaceId,omitempty" xml:"OppositeInterfaceId,omitempty"`
	// The ID of the Alibaba Cloud account to which the peer router interface belongs.
	//
	// example:
	//
	// 1321932713****
	OppositeInterfaceOwnerId *string `json:"OppositeInterfaceOwnerId,omitempty" xml:"OppositeInterfaceOwnerId,omitempty"`
	// The specification of the peer router interface. Valid values:
	//
	// 	- **Mini.2**: 2 Mbit/s
	//
	// 	- **Mini.5**: 5 Mbit/s
	//
	// 	- **Small.1**: 10 Mbit/s
	//
	// 	- **Small.2**: 20 Mbit/s
	//
	// 	- **Small.5**: 50 Mbit/s
	//
	// 	- **Middle.1**: 100 Mbit/s
	//
	// 	- **Middle.2**: 200 Mbit/s
	//
	// 	- **Middle.5**: 500 Mbit/s
	//
	// 	- **Large.1**: 1,000 Mbit/s
	//
	// 	- **Large.2**: 2,000 Mbit/s
	//
	// 	- **Large.5**: 5,000 Mbit/s
	//
	// 	- **Xlarge.1**: 10,000 Mbit/s
	//
	// 	- **Negative**: not applicable
	//
	// example:
	//
	// Negative
	OppositeInterfaceSpec *string `json:"OppositeInterfaceSpec,omitempty" xml:"OppositeInterfaceSpec,omitempty"`
	// The status of the peer router interface. Valid values:
	//
	// 	- **Idle**
	//
	// 	- **AcceptingConnecting**
	//
	// 	- **Connecting**
	//
	// 	- **Activating**
	//
	// 	- **Active**
	//
	// 	- **Modifying**
	//
	// 	- **Deactivating**
	//
	// 	- **Inactive**
	//
	// 	- **Deleting**
	//
	// 	- **Deleted**
	//
	// example:
	//
	// Active
	OppositeInterfaceStatus *string `json:"OppositeInterfaceStatus,omitempty" xml:"OppositeInterfaceStatus,omitempty"`
	// The region ID of the peer router interface.
	//
	// example:
	//
	// cn-hangzhou
	OppositeRegionId *string `json:"OppositeRegionId,omitempty" xml:"OppositeRegionId,omitempty"`
	// The ID of the router to which the peer router interface belongs.
	//
	// example:
	//
	// vrt-bp11xvy6lb9photuu****
	OppositeRouterId *string `json:"OppositeRouterId,omitempty" xml:"OppositeRouterId,omitempty"`
	// The type of the router to which the peer router interface belongs. Valid values:
	//
	// 	- **VRouter**
	//
	// 	- **VBR**
	//
	// example:
	//
	// VRouter
	OppositeRouterType *string `json:"OppositeRouterType,omitempty" xml:"OppositeRouterType,omitempty"`
	// The ID of the peer VPC.
	//
	// example:
	//
	// vpc-bp1b49rqrybk45nio****
	OppositeVpcInstanceId *string `json:"OppositeVpcInstanceId,omitempty" xml:"OppositeVpcInstanceId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 01818199-04F6-47F4-9ADF-7CC824CF57A4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The time when the renewal takes effect.
	//
	// example:
	//
	// 2022-06-11T16:00:00Z
	ReservationActiveTime *string `json:"ReservationActiveTime,omitempty" xml:"ReservationActiveTime,omitempty"`
	// The maximum bandwidth after the renewal takes effect. Unit: Mbit/s.
	//
	// example:
	//
	// 100
	ReservationBandwidth *string `json:"ReservationBandwidth,omitempty" xml:"ReservationBandwidth,omitempty"`
	// The metering method that is used after the renewal takes effect. Valid values: If **PayByBandwidth*	- is returned, it indicates that the Express Connect circuit is billed on a pay-by-bandwidth basis.
	//
	// example:
	//
	// PayByBandwidth
	ReservationInternetChargeType *string `json:"ReservationInternetChargeType,omitempty" xml:"ReservationInternetChargeType,omitempty"`
	// The type of the renewal order. Only **RENEW*	- may be returned, which indicates that the order is placed for service renewal.
	//
	// example:
	//
	// RENEW
	ReservationOrderType *string `json:"ReservationOrderType,omitempty" xml:"ReservationOrderType,omitempty"`
	// The resource group ID.
	//
	// For more information about resource groups, see [What is a resource group?](https://help.aliyun.com/document_detail/94475.html)
	//
	// example:
	//
	// rg-acfmxazb4ph6aiy****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The role of the router interface in the peering connection.
	//
	// example:
	//
	// InitiatingSide
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
	// The ID of the router to which the router interface belongs.
	//
	// example:
	//
	// vbr-m5ex0xf63xk8s5bob****
	RouterId *string `json:"RouterId,omitempty" xml:"RouterId,omitempty"`
	// The ID of the router interface.
	//
	// example:
	//
	// ri-m5egfc10sednwk2yt****
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
	// The specification of the router interface. Valid values:
	//
	// 	- **Mini.2**: 2 Mbit/s
	//
	// 	- **Mini.5**: 5 Mbit/s
	//
	// 	- **Small.1**: 10 Mbit/s
	//
	// 	- **Small.2**: 20 Mbit/s
	//
	// 	- **Small.5**: 50 Mbit/s
	//
	// 	- **Middle.1**: 100 Mbit/s
	//
	// 	- **Middle.2**: 200 Mbit/s
	//
	// 	- **Middle.5**: 500 Mbit/s
	//
	// 	- **Large.1**: 1,000 Mbit/s
	//
	// 	- **Large.2**: 2,000 Mbit/s
	//
	// 	- **Large.5**: 5,000 Mbit/s
	//
	// 	- **Xlarge.1**: 10,000 Mbit/s
	//
	// example:
	//
	// Mini.2
	Spec *string `json:"Spec,omitempty" xml:"Spec,omitempty"`
	// The status of the router interface. Valid values:
	//
	// 	- **Idle**
	//
	// 	- **AcceptingConnecting**
	//
	// 	- **Connecting**
	//
	// 	- **Activating**
	//
	// 	- **Active**
	//
	// 	- **Modifying**
	//
	// 	- **Deactivating**
	//
	// 	- **Inactive**
	//
	// 	- **Deleting**
	//
	// example:
	//
	// Active
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Indicates whether the request is successful. Valid values: true and false.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The tag of the resource.
	Tags *DescribeRouterInterfaceAttributeResponseBodyTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	// The ID of the virtual private cloud (VPC) to which the router interface belongs.
	//
	// example:
	//
	// vpc-bp1b49rqrybk45nio****
	VpcInstanceId *string `json:"VpcInstanceId,omitempty" xml:"VpcInstanceId,omitempty"`
}

func (s DescribeRouterInterfaceAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRouterInterfaceAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRouterInterfaceAttributeResponseBody) GetAccessPointId() *string {
	return s.AccessPointId
}

func (s *DescribeRouterInterfaceAttributeResponseBody) GetBandwidth() *int32 {
	return s.Bandwidth
}

func (s *DescribeRouterInterfaceAttributeResponseBody) GetBusinessStatus() *string {
	return s.BusinessStatus
}

func (s *DescribeRouterInterfaceAttributeResponseBody) GetChargeType() *string {
	return s.ChargeType
}

func (s *DescribeRouterInterfaceAttributeResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeRouterInterfaceAttributeResponseBody) GetConnectedTime() *string {
	return s.ConnectedTime
}

func (s *DescribeRouterInterfaceAttributeResponseBody) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DescribeRouterInterfaceAttributeResponseBody) GetCrossBorder() *bool {
	return s.CrossBorder
}

func (s *DescribeRouterInterfaceAttributeResponseBody) GetDescription() *string {
	return s.Description
}

func (s *DescribeRouterInterfaceAttributeResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeRouterInterfaceAttributeResponseBody) GetFastLinkMode() *string {
	return s.FastLinkMode
}

func (s *DescribeRouterInterfaceAttributeResponseBody) GetGmtModified() *string {
	return s.GmtModified
}

func (s *DescribeRouterInterfaceAttributeResponseBody) GetHasReservationData() *string {
	return s.HasReservationData
}

func (s *DescribeRouterInterfaceAttributeResponseBody) GetHcRate() *int32 {
	return s.HcRate
}

func (s *DescribeRouterInterfaceAttributeResponseBody) GetHcThreshold() *int32 {
	return s.HcThreshold
}

func (s *DescribeRouterInterfaceAttributeResponseBody) GetHealthCheckSourceIp() *string {
	return s.HealthCheckSourceIp
}

func (s *DescribeRouterInterfaceAttributeResponseBody) GetHealthCheckStatus() *string {
	return s.HealthCheckStatus
}

func (s *DescribeRouterInterfaceAttributeResponseBody) GetHealthCheckTargetIp() *string {
	return s.HealthCheckTargetIp
}

func (s *DescribeRouterInterfaceAttributeResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeRouterInterfaceAttributeResponseBody) GetName() *string {
	return s.Name
}

func (s *DescribeRouterInterfaceAttributeResponseBody) GetOppositeAccessPointId() *string {
	return s.OppositeAccessPointId
}

func (s *DescribeRouterInterfaceAttributeResponseBody) GetOppositeBandwidth() *int32 {
	return s.OppositeBandwidth
}

func (s *DescribeRouterInterfaceAttributeResponseBody) GetOppositeInterfaceBusinessStatus() *string {
	return s.OppositeInterfaceBusinessStatus
}

func (s *DescribeRouterInterfaceAttributeResponseBody) GetOppositeInterfaceId() *string {
	return s.OppositeInterfaceId
}

func (s *DescribeRouterInterfaceAttributeResponseBody) GetOppositeInterfaceOwnerId() *string {
	return s.OppositeInterfaceOwnerId
}

func (s *DescribeRouterInterfaceAttributeResponseBody) GetOppositeInterfaceSpec() *string {
	return s.OppositeInterfaceSpec
}

func (s *DescribeRouterInterfaceAttributeResponseBody) GetOppositeInterfaceStatus() *string {
	return s.OppositeInterfaceStatus
}

func (s *DescribeRouterInterfaceAttributeResponseBody) GetOppositeRegionId() *string {
	return s.OppositeRegionId
}

func (s *DescribeRouterInterfaceAttributeResponseBody) GetOppositeRouterId() *string {
	return s.OppositeRouterId
}

func (s *DescribeRouterInterfaceAttributeResponseBody) GetOppositeRouterType() *string {
	return s.OppositeRouterType
}

func (s *DescribeRouterInterfaceAttributeResponseBody) GetOppositeVpcInstanceId() *string {
	return s.OppositeVpcInstanceId
}

func (s *DescribeRouterInterfaceAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRouterInterfaceAttributeResponseBody) GetReservationActiveTime() *string {
	return s.ReservationActiveTime
}

func (s *DescribeRouterInterfaceAttributeResponseBody) GetReservationBandwidth() *string {
	return s.ReservationBandwidth
}

func (s *DescribeRouterInterfaceAttributeResponseBody) GetReservationInternetChargeType() *string {
	return s.ReservationInternetChargeType
}

func (s *DescribeRouterInterfaceAttributeResponseBody) GetReservationOrderType() *string {
	return s.ReservationOrderType
}

func (s *DescribeRouterInterfaceAttributeResponseBody) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeRouterInterfaceAttributeResponseBody) GetRole() *string {
	return s.Role
}

func (s *DescribeRouterInterfaceAttributeResponseBody) GetRouterId() *string {
	return s.RouterId
}

func (s *DescribeRouterInterfaceAttributeResponseBody) GetRouterInterfaceId() *string {
	return s.RouterInterfaceId
}

func (s *DescribeRouterInterfaceAttributeResponseBody) GetRouterType() *string {
	return s.RouterType
}

func (s *DescribeRouterInterfaceAttributeResponseBody) GetSpec() *string {
	return s.Spec
}

func (s *DescribeRouterInterfaceAttributeResponseBody) GetStatus() *string {
	return s.Status
}

func (s *DescribeRouterInterfaceAttributeResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeRouterInterfaceAttributeResponseBody) GetTags() *DescribeRouterInterfaceAttributeResponseBodyTags {
	return s.Tags
}

func (s *DescribeRouterInterfaceAttributeResponseBody) GetVpcInstanceId() *string {
	return s.VpcInstanceId
}

func (s *DescribeRouterInterfaceAttributeResponseBody) SetAccessPointId(v string) *DescribeRouterInterfaceAttributeResponseBody {
	s.AccessPointId = &v
	return s
}

func (s *DescribeRouterInterfaceAttributeResponseBody) SetBandwidth(v int32) *DescribeRouterInterfaceAttributeResponseBody {
	s.Bandwidth = &v
	return s
}

func (s *DescribeRouterInterfaceAttributeResponseBody) SetBusinessStatus(v string) *DescribeRouterInterfaceAttributeResponseBody {
	s.BusinessStatus = &v
	return s
}

func (s *DescribeRouterInterfaceAttributeResponseBody) SetChargeType(v string) *DescribeRouterInterfaceAttributeResponseBody {
	s.ChargeType = &v
	return s
}

func (s *DescribeRouterInterfaceAttributeResponseBody) SetCode(v string) *DescribeRouterInterfaceAttributeResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeRouterInterfaceAttributeResponseBody) SetConnectedTime(v string) *DescribeRouterInterfaceAttributeResponseBody {
	s.ConnectedTime = &v
	return s
}

func (s *DescribeRouterInterfaceAttributeResponseBody) SetCreationTime(v string) *DescribeRouterInterfaceAttributeResponseBody {
	s.CreationTime = &v
	return s
}

func (s *DescribeRouterInterfaceAttributeResponseBody) SetCrossBorder(v bool) *DescribeRouterInterfaceAttributeResponseBody {
	s.CrossBorder = &v
	return s
}

func (s *DescribeRouterInterfaceAttributeResponseBody) SetDescription(v string) *DescribeRouterInterfaceAttributeResponseBody {
	s.Description = &v
	return s
}

func (s *DescribeRouterInterfaceAttributeResponseBody) SetEndTime(v string) *DescribeRouterInterfaceAttributeResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeRouterInterfaceAttributeResponseBody) SetFastLinkMode(v string) *DescribeRouterInterfaceAttributeResponseBody {
	s.FastLinkMode = &v
	return s
}

func (s *DescribeRouterInterfaceAttributeResponseBody) SetGmtModified(v string) *DescribeRouterInterfaceAttributeResponseBody {
	s.GmtModified = &v
	return s
}

func (s *DescribeRouterInterfaceAttributeResponseBody) SetHasReservationData(v string) *DescribeRouterInterfaceAttributeResponseBody {
	s.HasReservationData = &v
	return s
}

func (s *DescribeRouterInterfaceAttributeResponseBody) SetHcRate(v int32) *DescribeRouterInterfaceAttributeResponseBody {
	s.HcRate = &v
	return s
}

func (s *DescribeRouterInterfaceAttributeResponseBody) SetHcThreshold(v int32) *DescribeRouterInterfaceAttributeResponseBody {
	s.HcThreshold = &v
	return s
}

func (s *DescribeRouterInterfaceAttributeResponseBody) SetHealthCheckSourceIp(v string) *DescribeRouterInterfaceAttributeResponseBody {
	s.HealthCheckSourceIp = &v
	return s
}

func (s *DescribeRouterInterfaceAttributeResponseBody) SetHealthCheckStatus(v string) *DescribeRouterInterfaceAttributeResponseBody {
	s.HealthCheckStatus = &v
	return s
}

func (s *DescribeRouterInterfaceAttributeResponseBody) SetHealthCheckTargetIp(v string) *DescribeRouterInterfaceAttributeResponseBody {
	s.HealthCheckTargetIp = &v
	return s
}

func (s *DescribeRouterInterfaceAttributeResponseBody) SetMessage(v string) *DescribeRouterInterfaceAttributeResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeRouterInterfaceAttributeResponseBody) SetName(v string) *DescribeRouterInterfaceAttributeResponseBody {
	s.Name = &v
	return s
}

func (s *DescribeRouterInterfaceAttributeResponseBody) SetOppositeAccessPointId(v string) *DescribeRouterInterfaceAttributeResponseBody {
	s.OppositeAccessPointId = &v
	return s
}

func (s *DescribeRouterInterfaceAttributeResponseBody) SetOppositeBandwidth(v int32) *DescribeRouterInterfaceAttributeResponseBody {
	s.OppositeBandwidth = &v
	return s
}

func (s *DescribeRouterInterfaceAttributeResponseBody) SetOppositeInterfaceBusinessStatus(v string) *DescribeRouterInterfaceAttributeResponseBody {
	s.OppositeInterfaceBusinessStatus = &v
	return s
}

func (s *DescribeRouterInterfaceAttributeResponseBody) SetOppositeInterfaceId(v string) *DescribeRouterInterfaceAttributeResponseBody {
	s.OppositeInterfaceId = &v
	return s
}

func (s *DescribeRouterInterfaceAttributeResponseBody) SetOppositeInterfaceOwnerId(v string) *DescribeRouterInterfaceAttributeResponseBody {
	s.OppositeInterfaceOwnerId = &v
	return s
}

func (s *DescribeRouterInterfaceAttributeResponseBody) SetOppositeInterfaceSpec(v string) *DescribeRouterInterfaceAttributeResponseBody {
	s.OppositeInterfaceSpec = &v
	return s
}

func (s *DescribeRouterInterfaceAttributeResponseBody) SetOppositeInterfaceStatus(v string) *DescribeRouterInterfaceAttributeResponseBody {
	s.OppositeInterfaceStatus = &v
	return s
}

func (s *DescribeRouterInterfaceAttributeResponseBody) SetOppositeRegionId(v string) *DescribeRouterInterfaceAttributeResponseBody {
	s.OppositeRegionId = &v
	return s
}

func (s *DescribeRouterInterfaceAttributeResponseBody) SetOppositeRouterId(v string) *DescribeRouterInterfaceAttributeResponseBody {
	s.OppositeRouterId = &v
	return s
}

func (s *DescribeRouterInterfaceAttributeResponseBody) SetOppositeRouterType(v string) *DescribeRouterInterfaceAttributeResponseBody {
	s.OppositeRouterType = &v
	return s
}

func (s *DescribeRouterInterfaceAttributeResponseBody) SetOppositeVpcInstanceId(v string) *DescribeRouterInterfaceAttributeResponseBody {
	s.OppositeVpcInstanceId = &v
	return s
}

func (s *DescribeRouterInterfaceAttributeResponseBody) SetRequestId(v string) *DescribeRouterInterfaceAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRouterInterfaceAttributeResponseBody) SetReservationActiveTime(v string) *DescribeRouterInterfaceAttributeResponseBody {
	s.ReservationActiveTime = &v
	return s
}

func (s *DescribeRouterInterfaceAttributeResponseBody) SetReservationBandwidth(v string) *DescribeRouterInterfaceAttributeResponseBody {
	s.ReservationBandwidth = &v
	return s
}

func (s *DescribeRouterInterfaceAttributeResponseBody) SetReservationInternetChargeType(v string) *DescribeRouterInterfaceAttributeResponseBody {
	s.ReservationInternetChargeType = &v
	return s
}

func (s *DescribeRouterInterfaceAttributeResponseBody) SetReservationOrderType(v string) *DescribeRouterInterfaceAttributeResponseBody {
	s.ReservationOrderType = &v
	return s
}

func (s *DescribeRouterInterfaceAttributeResponseBody) SetResourceGroupId(v string) *DescribeRouterInterfaceAttributeResponseBody {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeRouterInterfaceAttributeResponseBody) SetRole(v string) *DescribeRouterInterfaceAttributeResponseBody {
	s.Role = &v
	return s
}

func (s *DescribeRouterInterfaceAttributeResponseBody) SetRouterId(v string) *DescribeRouterInterfaceAttributeResponseBody {
	s.RouterId = &v
	return s
}

func (s *DescribeRouterInterfaceAttributeResponseBody) SetRouterInterfaceId(v string) *DescribeRouterInterfaceAttributeResponseBody {
	s.RouterInterfaceId = &v
	return s
}

func (s *DescribeRouterInterfaceAttributeResponseBody) SetRouterType(v string) *DescribeRouterInterfaceAttributeResponseBody {
	s.RouterType = &v
	return s
}

func (s *DescribeRouterInterfaceAttributeResponseBody) SetSpec(v string) *DescribeRouterInterfaceAttributeResponseBody {
	s.Spec = &v
	return s
}

func (s *DescribeRouterInterfaceAttributeResponseBody) SetStatus(v string) *DescribeRouterInterfaceAttributeResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeRouterInterfaceAttributeResponseBody) SetSuccess(v bool) *DescribeRouterInterfaceAttributeResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeRouterInterfaceAttributeResponseBody) SetTags(v *DescribeRouterInterfaceAttributeResponseBodyTags) *DescribeRouterInterfaceAttributeResponseBody {
	s.Tags = v
	return s
}

func (s *DescribeRouterInterfaceAttributeResponseBody) SetVpcInstanceId(v string) *DescribeRouterInterfaceAttributeResponseBody {
	s.VpcInstanceId = &v
	return s
}

func (s *DescribeRouterInterfaceAttributeResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeRouterInterfaceAttributeResponseBodyTags struct {
	Tags []*DescribeRouterInterfaceAttributeResponseBodyTagsTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s DescribeRouterInterfaceAttributeResponseBodyTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeRouterInterfaceAttributeResponseBodyTags) GoString() string {
	return s.String()
}

func (s *DescribeRouterInterfaceAttributeResponseBodyTags) GetTags() []*DescribeRouterInterfaceAttributeResponseBodyTagsTags {
	return s.Tags
}

func (s *DescribeRouterInterfaceAttributeResponseBodyTags) SetTags(v []*DescribeRouterInterfaceAttributeResponseBodyTagsTags) *DescribeRouterInterfaceAttributeResponseBodyTags {
	s.Tags = v
	return s
}

func (s *DescribeRouterInterfaceAttributeResponseBodyTags) Validate() error {
	return dara.Validate(s)
}

type DescribeRouterInterfaceAttributeResponseBodyTagsTags struct {
	// The key of tag N added to the resource. You must enter at least one tag key and at most 20 tag keys. The tag key cannot be an empty string.
	//
	// The tag key can be up to 64 characters in length and can contain digits, periods (.), underscores (_), and hyphens (-). It cannot start with `aliyun` or `acs:`, and cannot contain `http://` or `https://`.
	//
	// example:
	//
	// FinanceDept
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of tag N added to the resource. You can specify at most 20 tag values. The tag value can be an empty string.
	//
	// It can be up to 128 characters in length and can contain digits, periods (.), underscores (_), and hyphens (-). It cannot start with `aliyun` or `acs:`, and cannot contain `http://` or `https://`.
	//
	// example:
	//
	// FinanceJoshua
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeRouterInterfaceAttributeResponseBodyTagsTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeRouterInterfaceAttributeResponseBodyTagsTags) GoString() string {
	return s.String()
}

func (s *DescribeRouterInterfaceAttributeResponseBodyTagsTags) GetKey() *string {
	return s.Key
}

func (s *DescribeRouterInterfaceAttributeResponseBodyTagsTags) GetValue() *string {
	return s.Value
}

func (s *DescribeRouterInterfaceAttributeResponseBodyTagsTags) SetKey(v string) *DescribeRouterInterfaceAttributeResponseBodyTagsTags {
	s.Key = &v
	return s
}

func (s *DescribeRouterInterfaceAttributeResponseBodyTagsTags) SetValue(v string) *DescribeRouterInterfaceAttributeResponseBodyTagsTags {
	s.Value = &v
	return s
}

func (s *DescribeRouterInterfaceAttributeResponseBodyTagsTags) Validate() error {
	return dara.Validate(s)
}
