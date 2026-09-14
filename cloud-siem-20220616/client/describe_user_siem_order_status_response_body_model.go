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
	// The request ID.
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
	// The Agentic SOC Credits instance ID. If SiemOrderFrom is CREDITS_PRE_PAY, this field returns the Credits subscription instance ID for prepaid orders. If SiemOrderFrom is CREDITS_POST_PAY, this field returns the Credits pay-as-you-go instance ID. This field is empty if no Credits instance is found. For legacy orders, the Security Center instance ID is returned by SasInstanceId.
	//
	// example:
	//
	// asoc-instance-xxxxx
	AsocInstanceId *string `json:"AsocInstanceId,omitempty" xml:"AsocInstanceId,omitempty"`
	// The end time of the Agentic SOC Credits prepaid subscription, expressed as a 13-digit Unix timestamp in milliseconds. This field is returned only when SiemOrderFrom is CREDITS_PRE_PAY. In other cases, this field is empty.
	//
	// example:
	//
	// 1785542400456
	AsocSubscriptionInstanceEndTime *int64 `json:"AsocSubscriptionInstanceEndTime,omitempty" xml:"AsocSubscriptionInstanceEndTime,omitempty"`
	// The start time of the Agentic SOC Credits prepaid subscription, expressed as a 13-digit Unix timestamp in milliseconds. This field is returned only when SiemOrderFrom is CREDITS_PRE_PAY. In other cases, this field is empty.
	//
	// example:
	//
	// 1754006400123
	AsocSubscriptionInstanceStartTime *int64 `json:"AsocSubscriptionInstanceStartTime,omitempty" xml:"AsocSubscriptionInstanceStartTime,omitempty"`
	// Indicates whether the current account can perform order operations for threat detection and response. Valid values:
	//
	// - true: The account can purchase, upgrade, or change specifications.
	//
	// - false: The account cannot perform order operations for threat detection and response.
	//
	// example:
	//
	// true
	CanBuy *bool `json:"CanBuy,omitempty" xml:"CanBuy,omitempty"`
	// The SLS log storage capacity purchased for threat detection and response, in GB.
	//
	// example:
	//
	// 1024
	Capacity *int32 `json:"Capacity,omitempty" xml:"Capacity,omitempty"`
	// The source of the log storage capacity order. Valid values:
	//
	// - PRE_PAY_CAPACITY: a prepaid capacity order.
	//
	// - POST_PAY_CAPACITY: a pay-as-you-go capacity order.
	//
	// The capacity order source is independent of the traffic order source indicated by SiemOrderFrom.
	//
	// example:
	//
	// PRE_PAY_CAPACITY
	CapacityOrderFrom *string `json:"CapacityOrderFrom,omitempty" xml:"CapacityOrderFrom,omitempty"`
	// The SLS log storage capacity purchased for threat detection and response 1.0, in GB.
	//
	// example:
	//
	// 1024
	DeliveryCapacity *int32 `json:"DeliveryCapacity,omitempty" xml:"DeliveryCapacity,omitempty"`
	// The number of days until the threat detection and response service expires.
	//
	// example:
	//
	// 3
	DurationDays *int64 `json:"DurationDays,omitempty" xml:"DurationDays,omitempty"`
	// The expiration time of threat detection and response, expressed as a millisecond-level timestamp.
	//
	// example:
	//
	// 1669823999000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The traffic capacity purchased for threat detection and response, in GB.
	//
	// example:
	//
	// 1024
	FlowCapacity *int32 `json:"FlowCapacity,omitempty" xml:"FlowCapacity,omitempty"`
	// The Alibaba Cloud account ID that purchased threat detection and response.
	//
	// example:
	//
	// 123XXXXXX
	MainUserId *int64 `json:"MainUserId,omitempty" xml:"MainUserId,omitempty"`
	// The master account ID of the resource directory.
	//
	// example:
	//
	// 123XXXXXX
	MasterUserId *int64 `json:"MasterUserId,omitempty" xml:"MasterUserId,omitempty"`
	// The resource directory ID.
	//
	// example:
	//
	// rd-xxxxxx
	RdId *string `json:"RdId,omitempty" xml:"RdId,omitempty"`
	// Indicates whether the order is a SIEM public preview order.
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
	// - PRE_PAY_FLOW: a prepaid traffic order for threat detection and response.
	//
	// - POST_PAY_FLOW: a pay-as-you-go traffic order for threat detection and response.
	//
	// - CREDITS_PRE_PAY: an Agentic SOC Credits prepaid subscription.
	//
	// - CREDITS_POST_PAY: an Agentic SOC Credits pay-as-you-go instance.
	//
	// This field describes the traffic order source. The log storage capacity order source is independently indicated by CapacityOrderFrom.
	//
	// example:
	//
	// CREDITS_PRE_PAY
	SiemOrderFrom *string `json:"SiemOrderFrom,omitempty" xml:"SiemOrderFrom,omitempty"`
	// Indicates whether a valid SIEM order exists. Valid values:
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
	return dara.Validate(s)
}
