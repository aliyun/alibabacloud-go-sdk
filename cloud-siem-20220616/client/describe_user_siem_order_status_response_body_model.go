// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUserSiemOrderStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DescribeUserSiemOrderStatusResponseBodyData) *DescribeUserSiemOrderStatusResponseBody
	GetData() *DescribeUserSiemOrderStatusResponseBodyData
	SetRequestId(v string) *DescribeUserSiemOrderStatusResponseBody
	GetRequestId() *string
}

type DescribeUserSiemOrderStatusResponseBody struct {
	// The response data.
	Data *DescribeUserSiemOrderStatusResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 6276D891-*****-55B2-87B9-74D413F7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeUserSiemOrderStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserSiemOrderStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeUserSiemOrderStatusResponseBody) GetData() *DescribeUserSiemOrderStatusResponseBodyData {
	return s.Data
}

func (s *DescribeUserSiemOrderStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeUserSiemOrderStatusResponseBody) SetData(v *DescribeUserSiemOrderStatusResponseBodyData) *DescribeUserSiemOrderStatusResponseBody {
	s.Data = v
	return s
}

func (s *DescribeUserSiemOrderStatusResponseBody) SetRequestId(v string) *DescribeUserSiemOrderStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeUserSiemOrderStatusResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeUserSiemOrderStatusResponseBodyData struct {
	// The ID of the Agentic SOC Credits instance. If SiemOrderFrom is CREDITS_PRE_PAY, the ID of the Credits subscription instance is returned. If SiemOrderFrom is CREDITS_POST_PAY, the ID of the Credits pay-as-you-go instance is returned. If no Credits instance is found, this field is empty. The Security Center instance ID for legacy orders is returned by SasInstanceId.
	//
	// example:
	//
	// asoc-instance-xxxxx
	AsocInstanceId *string `json:"AsocInstanceId,omitempty" xml:"AsocInstanceId,omitempty"`
	// The end time of the Agentic SOC Credits subscription. The value is a UNIX timestamp in milliseconds (13 digits). This field is returned only when SiemOrderFrom is CREDITS_PRE_PAY. Otherwise, this field is empty.
	//
	// example:
	//
	// 1785542400456
	AsocSubscriptionInstanceEndTime *int64 `json:"AsocSubscriptionInstanceEndTime,omitempty" xml:"AsocSubscriptionInstanceEndTime,omitempty"`
	// The start time of the Agentic SOC Credits subscription. The value is a UNIX timestamp in milliseconds (13 digits). This field is returned only when SiemOrderFrom is CREDITS_PRE_PAY. Otherwise, this field is empty.
	//
	// example:
	//
	// 1754006400123
	AsocSubscriptionInstanceStartTime *int64 `json:"AsocSubscriptionInstanceStartTime,omitempty" xml:"AsocSubscriptionInstanceStartTime,omitempty"`
	// Specifies whether the current account can perform operations on threat analysis orders. Valid values:
	//
	// - true: The account can purchase, upgrade, or change the configuration of threat analysis.
	//
	// - false: The account cannot perform operations on threat analysis orders.
	//
	// example:
	//
	// true
	CanBuy *bool `json:"CanBuy,omitempty" xml:"CanBuy,omitempty"`
	// The purchased Simple Log Service (SLS) capacity for threat analysis. Unit: GB.
	//
	// example:
	//
	// 1024
	Capacity *int32 `json:"Capacity,omitempty" xml:"Capacity,omitempty"`
	// The source of the log storage capacity order. Valid values:
	//
	// - PRE_PAY_CAPACITY: Prepaid capacity order.
	//
	// - POST_PAY_CAPACITY: Pay-as-you-go capacity order.
	//
	// The source of the capacity order is independent of the traffic order source represented by SiemOrderFrom.
	//
	// example:
	//
	// PRE_PAY_CAPACITY
	CapacityOrderFrom *string `json:"CapacityOrderFrom,omitempty" xml:"CapacityOrderFrom,omitempty"`
	// The information about the Credits pay-as-you-go instance. This field is returned when a Credits pay-as-you-go instance is found, and can be returned simultaneously with a valid legacy traffic order.
	CreditsPostpaidInfo *DescribeUserSiemOrderStatusResponseBodyDataCreditsPostpaidInfo `json:"CreditsPostpaidInfo,omitempty" xml:"CreditsPostpaidInfo,omitempty" type:"Struct"`
	// The purchased SLS capacity for Threat Analysis 1.0. Unit: GB.
	//
	// example:
	//
	// 1024
	DeliveryCapacity *int32 `json:"DeliveryCapacity,omitempty" xml:"DeliveryCapacity,omitempty"`
	// The number of days until the threat analysis service expires.
	//
	// example:
	//
	// 3
	DurationDays *int64 `json:"DurationDays,omitempty" xml:"DurationDays,omitempty"`
	// The expiration time of threat analysis. The value is a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1669823999000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The purchased traffic for threat analysis. Unit: GB.
	//
	// example:
	//
	// 1024
	FlowCapacity *int32 `json:"FlowCapacity,omitempty" xml:"FlowCapacity,omitempty"`
	// The Alibaba Cloud account ID that purchased threat analysis.
	//
	// example:
	//
	// 123XXXXXX
	MainUserId *int64 `json:"MainUserId,omitempty" xml:"MainUserId,omitempty"`
	// The ID of the master account in the resource directory.
	//
	// example:
	//
	// 123XXXXXX
	MasterUserId *int64 `json:"MasterUserId,omitempty" xml:"MasterUserId,omitempty"`
	// The ID of the resource directory.
	//
	// example:
	//
	// rd-xxxxxx
	RdId *string `json:"RdId,omitempty" xml:"RdId,omitempty"`
	// Specifies whether the order is a SIEM public preview order.
	//
	// example:
	//
	// 1
	RdOrder *int32 `json:"RdOrder,omitempty" xml:"RdOrder,omitempty"`
	// The Security Center instance ID.
	//
	// example:
	//
	// sas-instance-xxxxx
	SasInstanceId *string `json:"SasInstanceId,omitempty" xml:"SasInstanceId,omitempty"`
	// The source of the traffic order. Valid values:
	//
	// - PRE_PAY_FLOW: Prepaid traffic order for threat analysis.
	//
	// - POST_PAY_FLOW: Pay-as-you-go traffic order for threat analysis.
	//
	// - CREDITS_PRE_PAY: Agentic SOC Credits subscription.
	//
	// - CREDITS_POST_PAY: Agentic SOC Credits pay-as-you-go instance.
	//
	// This field describes the source of the traffic order. The source of the log storage capacity order is independently represented by CapacityOrderFrom.
	//
	// example:
	//
	// CREDITS_PRE_PAY
	SiemOrderFrom *string `json:"SiemOrderFrom,omitempty" xml:"SiemOrderFrom,omitempty"`
	// Specifies whether there is a valid SIEM order. Valid values:
	//
	// - 1: The SIEM order is valid.
	//
	// - 0: The SIEM order is invalid.
	//
	// example:
	//
	// 1
	SiemOrderStatus *int32 `json:"SiemOrderStatus,omitempty" xml:"SiemOrderStatus,omitempty"`
	// The Alibaba Cloud account ID of the current logon.
	//
	// example:
	//
	// 123XXXXXX
	SubUserId *int64 `json:"SubUserId,omitempty" xml:"SubUserId,omitempty"`
	// The user type.
	//
	// example:
	//
	// normal
	UserType *string `json:"UserType,omitempty" xml:"UserType,omitempty"`
}

func (s DescribeUserSiemOrderStatusResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserSiemOrderStatusResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeUserSiemOrderStatusResponseBodyData) GetAsocInstanceId() *string {
	return s.AsocInstanceId
}

func (s *DescribeUserSiemOrderStatusResponseBodyData) GetAsocSubscriptionInstanceEndTime() *int64 {
	return s.AsocSubscriptionInstanceEndTime
}

func (s *DescribeUserSiemOrderStatusResponseBodyData) GetAsocSubscriptionInstanceStartTime() *int64 {
	return s.AsocSubscriptionInstanceStartTime
}

func (s *DescribeUserSiemOrderStatusResponseBodyData) GetCanBuy() *bool {
	return s.CanBuy
}

func (s *DescribeUserSiemOrderStatusResponseBodyData) GetCapacity() *int32 {
	return s.Capacity
}

func (s *DescribeUserSiemOrderStatusResponseBodyData) GetCapacityOrderFrom() *string {
	return s.CapacityOrderFrom
}

func (s *DescribeUserSiemOrderStatusResponseBodyData) GetCreditsPostpaidInfo() *DescribeUserSiemOrderStatusResponseBodyDataCreditsPostpaidInfo {
	return s.CreditsPostpaidInfo
}

func (s *DescribeUserSiemOrderStatusResponseBodyData) GetDeliveryCapacity() *int32 {
	return s.DeliveryCapacity
}

func (s *DescribeUserSiemOrderStatusResponseBodyData) GetDurationDays() *int64 {
	return s.DurationDays
}

func (s *DescribeUserSiemOrderStatusResponseBodyData) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeUserSiemOrderStatusResponseBodyData) GetFlowCapacity() *int32 {
	return s.FlowCapacity
}

func (s *DescribeUserSiemOrderStatusResponseBodyData) GetMainUserId() *int64 {
	return s.MainUserId
}

func (s *DescribeUserSiemOrderStatusResponseBodyData) GetMasterUserId() *int64 {
	return s.MasterUserId
}

func (s *DescribeUserSiemOrderStatusResponseBodyData) GetRdId() *string {
	return s.RdId
}

func (s *DescribeUserSiemOrderStatusResponseBodyData) GetRdOrder() *int32 {
	return s.RdOrder
}

func (s *DescribeUserSiemOrderStatusResponseBodyData) GetSasInstanceId() *string {
	return s.SasInstanceId
}

func (s *DescribeUserSiemOrderStatusResponseBodyData) GetSiemOrderFrom() *string {
	return s.SiemOrderFrom
}

func (s *DescribeUserSiemOrderStatusResponseBodyData) GetSiemOrderStatus() *int32 {
	return s.SiemOrderStatus
}

func (s *DescribeUserSiemOrderStatusResponseBodyData) GetSubUserId() *int64 {
	return s.SubUserId
}

func (s *DescribeUserSiemOrderStatusResponseBodyData) GetUserType() *string {
	return s.UserType
}

func (s *DescribeUserSiemOrderStatusResponseBodyData) SetAsocInstanceId(v string) *DescribeUserSiemOrderStatusResponseBodyData {
	s.AsocInstanceId = &v
	return s
}

func (s *DescribeUserSiemOrderStatusResponseBodyData) SetAsocSubscriptionInstanceEndTime(v int64) *DescribeUserSiemOrderStatusResponseBodyData {
	s.AsocSubscriptionInstanceEndTime = &v
	return s
}

func (s *DescribeUserSiemOrderStatusResponseBodyData) SetAsocSubscriptionInstanceStartTime(v int64) *DescribeUserSiemOrderStatusResponseBodyData {
	s.AsocSubscriptionInstanceStartTime = &v
	return s
}

func (s *DescribeUserSiemOrderStatusResponseBodyData) SetCanBuy(v bool) *DescribeUserSiemOrderStatusResponseBodyData {
	s.CanBuy = &v
	return s
}

func (s *DescribeUserSiemOrderStatusResponseBodyData) SetCapacity(v int32) *DescribeUserSiemOrderStatusResponseBodyData {
	s.Capacity = &v
	return s
}

func (s *DescribeUserSiemOrderStatusResponseBodyData) SetCapacityOrderFrom(v string) *DescribeUserSiemOrderStatusResponseBodyData {
	s.CapacityOrderFrom = &v
	return s
}

func (s *DescribeUserSiemOrderStatusResponseBodyData) SetCreditsPostpaidInfo(v *DescribeUserSiemOrderStatusResponseBodyDataCreditsPostpaidInfo) *DescribeUserSiemOrderStatusResponseBodyData {
	s.CreditsPostpaidInfo = v
	return s
}

func (s *DescribeUserSiemOrderStatusResponseBodyData) SetDeliveryCapacity(v int32) *DescribeUserSiemOrderStatusResponseBodyData {
	s.DeliveryCapacity = &v
	return s
}

func (s *DescribeUserSiemOrderStatusResponseBodyData) SetDurationDays(v int64) *DescribeUserSiemOrderStatusResponseBodyData {
	s.DurationDays = &v
	return s
}

func (s *DescribeUserSiemOrderStatusResponseBodyData) SetEndTime(v int64) *DescribeUserSiemOrderStatusResponseBodyData {
	s.EndTime = &v
	return s
}

func (s *DescribeUserSiemOrderStatusResponseBodyData) SetFlowCapacity(v int32) *DescribeUserSiemOrderStatusResponseBodyData {
	s.FlowCapacity = &v
	return s
}

func (s *DescribeUserSiemOrderStatusResponseBodyData) SetMainUserId(v int64) *DescribeUserSiemOrderStatusResponseBodyData {
	s.MainUserId = &v
	return s
}

func (s *DescribeUserSiemOrderStatusResponseBodyData) SetMasterUserId(v int64) *DescribeUserSiemOrderStatusResponseBodyData {
	s.MasterUserId = &v
	return s
}

func (s *DescribeUserSiemOrderStatusResponseBodyData) SetRdId(v string) *DescribeUserSiemOrderStatusResponseBodyData {
	s.RdId = &v
	return s
}

func (s *DescribeUserSiemOrderStatusResponseBodyData) SetRdOrder(v int32) *DescribeUserSiemOrderStatusResponseBodyData {
	s.RdOrder = &v
	return s
}

func (s *DescribeUserSiemOrderStatusResponseBodyData) SetSasInstanceId(v string) *DescribeUserSiemOrderStatusResponseBodyData {
	s.SasInstanceId = &v
	return s
}

func (s *DescribeUserSiemOrderStatusResponseBodyData) SetSiemOrderFrom(v string) *DescribeUserSiemOrderStatusResponseBodyData {
	s.SiemOrderFrom = &v
	return s
}

func (s *DescribeUserSiemOrderStatusResponseBodyData) SetSiemOrderStatus(v int32) *DescribeUserSiemOrderStatusResponseBodyData {
	s.SiemOrderStatus = &v
	return s
}

func (s *DescribeUserSiemOrderStatusResponseBodyData) SetSubUserId(v int64) *DescribeUserSiemOrderStatusResponseBodyData {
	s.SubUserId = &v
	return s
}

func (s *DescribeUserSiemOrderStatusResponseBodyData) SetUserType(v string) *DescribeUserSiemOrderStatusResponseBodyData {
	s.UserType = &v
	return s
}

func (s *DescribeUserSiemOrderStatusResponseBodyData) Validate() error {
	if s.CreditsPostpaidInfo != nil {
		if err := s.CreditsPostpaidInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeUserSiemOrderStatusResponseBodyDataCreditsPostpaidInfo struct {
	// The billing status of the instance. Valid values:
	//
	// - 1: Normal.
	//
	// - 2: Suspended due to overdue payment.
	//
	// example:
	//
	// 2
	BillingStatus *int32 `json:"BillingStatus,omitempty" xml:"BillingStatus,omitempty"`
	// The source of the instance. Valid values:
	//
	// - POSTPAID: Independent pay-as-you-go instance.
	//
	// - SUBSCRIPTION_DEFAULT_PAYG: Default pay-as-you-go instance for subscription.
	//
	// example:
	//
	// POSTPAID
	InstanceSource *string `json:"InstanceSource,omitempty" xml:"InstanceSource,omitempty"`
	// The user switch status. Valid values:
	//
	// - 1: Enabled.
	//
	// - 0: Disabled.
	//
	// example:
	//
	// 1
	UserSwitchStatus *int32 `json:"UserSwitchStatus,omitempty" xml:"UserSwitchStatus,omitempty"`
}

func (s DescribeUserSiemOrderStatusResponseBodyDataCreditsPostpaidInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserSiemOrderStatusResponseBodyDataCreditsPostpaidInfo) GoString() string {
	return s.String()
}

func (s *DescribeUserSiemOrderStatusResponseBodyDataCreditsPostpaidInfo) GetBillingStatus() *int32 {
	return s.BillingStatus
}

func (s *DescribeUserSiemOrderStatusResponseBodyDataCreditsPostpaidInfo) GetInstanceSource() *string {
	return s.InstanceSource
}

func (s *DescribeUserSiemOrderStatusResponseBodyDataCreditsPostpaidInfo) GetUserSwitchStatus() *int32 {
	return s.UserSwitchStatus
}

func (s *DescribeUserSiemOrderStatusResponseBodyDataCreditsPostpaidInfo) SetBillingStatus(v int32) *DescribeUserSiemOrderStatusResponseBodyDataCreditsPostpaidInfo {
	s.BillingStatus = &v
	return s
}

func (s *DescribeUserSiemOrderStatusResponseBodyDataCreditsPostpaidInfo) SetInstanceSource(v string) *DescribeUserSiemOrderStatusResponseBodyDataCreditsPostpaidInfo {
	s.InstanceSource = &v
	return s
}

func (s *DescribeUserSiemOrderStatusResponseBodyDataCreditsPostpaidInfo) SetUserSwitchStatus(v int32) *DescribeUserSiemOrderStatusResponseBodyDataCreditsPostpaidInfo {
	s.UserSwitchStatus = &v
	return s
}

func (s *DescribeUserSiemOrderStatusResponseBodyDataCreditsPostpaidInfo) Validate() error {
	return dara.Validate(s)
}
