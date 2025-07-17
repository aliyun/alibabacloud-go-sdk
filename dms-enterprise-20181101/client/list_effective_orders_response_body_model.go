// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEffectiveOrdersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *ListEffectiveOrdersResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ListEffectiveOrdersResponseBody
	GetErrorMessage() *string
	SetOrderSummary(v []*ListEffectiveOrdersResponseBodyOrderSummary) *ListEffectiveOrdersResponseBody
	GetOrderSummary() []*ListEffectiveOrdersResponseBodyOrderSummary
	SetRequestId(v string) *ListEffectiveOrdersResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListEffectiveOrdersResponseBody
	GetSuccess() *bool
}

type ListEffectiveOrdersResponseBody struct {
	// The error code returned if the request failed.
	//
	// example:
	//
	// UnknownError
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned if the request failed.
	//
	// example:
	//
	// UnknownError
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The information about orders.
	OrderSummary []*ListEffectiveOrdersResponseBodyOrderSummary `json:"OrderSummary,omitempty" xml:"OrderSummary,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// A8FE12AA-300D-5FDF-806F-C2CB99161F32
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**: The request was successful.
	//
	// 	- **false**: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListEffectiveOrdersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListEffectiveOrdersResponseBody) GoString() string {
	return s.String()
}

func (s *ListEffectiveOrdersResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListEffectiveOrdersResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListEffectiveOrdersResponseBody) GetOrderSummary() []*ListEffectiveOrdersResponseBodyOrderSummary {
	return s.OrderSummary
}

func (s *ListEffectiveOrdersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListEffectiveOrdersResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListEffectiveOrdersResponseBody) SetErrorCode(v string) *ListEffectiveOrdersResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListEffectiveOrdersResponseBody) SetErrorMessage(v string) *ListEffectiveOrdersResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListEffectiveOrdersResponseBody) SetOrderSummary(v []*ListEffectiveOrdersResponseBodyOrderSummary) *ListEffectiveOrdersResponseBody {
	s.OrderSummary = v
	return s
}

func (s *ListEffectiveOrdersResponseBody) SetRequestId(v string) *ListEffectiveOrdersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListEffectiveOrdersResponseBody) SetSuccess(v bool) *ListEffectiveOrdersResponseBody {
	s.Success = &v
	return s
}

func (s *ListEffectiveOrdersResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListEffectiveOrdersResponseBodyOrderSummary struct {
	// The commodity code of DMS.
	//
	// 	- dms_pre_public_cn: DMS that uses the subscription billing method
	//
	// 	- dms_post_public_cn: DMS that uses the pay-as-you-go billing method
	//
	// example:
	//
	// dms_pre_public_cn
	CommodityCode *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	// The type of the service.
	//
	// 	- **VersionType**: DMS that supports control modes
	//
	// 	- **SensitiveDataProtection**: DMS that supports sensitive data protection
	//
	// example:
	//
	// VersionType
	CommodityType *string `json:"CommodityType,omitempty" xml:"CommodityType,omitempty"`
	// Details about the orders.
	OrderList []*ListEffectiveOrdersResponseBodyOrderSummaryOrderList `json:"OrderList,omitempty" xml:"OrderList,omitempty" type:"Repeated"`
	// The sum of the number of instances that you can use DMS to manage in all orders.
	//
	// example:
	//
	// 12
	TotalQuota *int32 `json:"TotalQuota,omitempty" xml:"TotalQuota,omitempty"`
	// The control mode of DMS. Valid values:
	//
	// 	- **stand**: Stable Change
	//
	// 	- **safety**: Security Collaboration
	//
	// example:
	//
	// safety
	VersionType *string `json:"VersionType,omitempty" xml:"VersionType,omitempty"`
}

func (s ListEffectiveOrdersResponseBodyOrderSummary) String() string {
	return dara.Prettify(s)
}

func (s ListEffectiveOrdersResponseBodyOrderSummary) GoString() string {
	return s.String()
}

func (s *ListEffectiveOrdersResponseBodyOrderSummary) GetCommodityCode() *string {
	return s.CommodityCode
}

func (s *ListEffectiveOrdersResponseBodyOrderSummary) GetCommodityType() *string {
	return s.CommodityType
}

func (s *ListEffectiveOrdersResponseBodyOrderSummary) GetOrderList() []*ListEffectiveOrdersResponseBodyOrderSummaryOrderList {
	return s.OrderList
}

func (s *ListEffectiveOrdersResponseBodyOrderSummary) GetTotalQuota() *int32 {
	return s.TotalQuota
}

func (s *ListEffectiveOrdersResponseBodyOrderSummary) GetVersionType() *string {
	return s.VersionType
}

func (s *ListEffectiveOrdersResponseBodyOrderSummary) SetCommodityCode(v string) *ListEffectiveOrdersResponseBodyOrderSummary {
	s.CommodityCode = &v
	return s
}

func (s *ListEffectiveOrdersResponseBodyOrderSummary) SetCommodityType(v string) *ListEffectiveOrdersResponseBodyOrderSummary {
	s.CommodityType = &v
	return s
}

func (s *ListEffectiveOrdersResponseBodyOrderSummary) SetOrderList(v []*ListEffectiveOrdersResponseBodyOrderSummaryOrderList) *ListEffectiveOrdersResponseBodyOrderSummary {
	s.OrderList = v
	return s
}

func (s *ListEffectiveOrdersResponseBodyOrderSummary) SetTotalQuota(v int32) *ListEffectiveOrdersResponseBodyOrderSummary {
	s.TotalQuota = &v
	return s
}

func (s *ListEffectiveOrdersResponseBodyOrderSummary) SetVersionType(v string) *ListEffectiveOrdersResponseBodyOrderSummary {
	s.VersionType = &v
	return s
}

func (s *ListEffectiveOrdersResponseBodyOrderSummary) Validate() error {
	return dara.Validate(s)
}

type ListEffectiveOrdersResponseBodyOrderSummaryOrderList struct {
	// The UID of the user who placed the order.
	//
	// example:
	//
	// 2698420314****
	BuyerId *string `json:"BuyerId,omitempty" xml:"BuyerId,omitempty"`
	// The time when the instance expires.
	//
	// example:
	//
	// 2022-11-24 00:00:00
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The maximum number of database instances that you can use DMS to manage.
	//
	// example:
	//
	// 7
	InsNum *string `json:"InsNum,omitempty" xml:"InsNum,omitempty"`
	// The ID of the instance for the purchased service.
	//
	// example:
	//
	// rm-bp1xd1v866****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the order.
	//
	// example:
	//
	// 2190037****
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The time when the instance is started.
	//
	// example:
	//
	// 2022-10-24 00:00:00
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ListEffectiveOrdersResponseBodyOrderSummaryOrderList) String() string {
	return dara.Prettify(s)
}

func (s ListEffectiveOrdersResponseBodyOrderSummaryOrderList) GoString() string {
	return s.String()
}

func (s *ListEffectiveOrdersResponseBodyOrderSummaryOrderList) GetBuyerId() *string {
	return s.BuyerId
}

func (s *ListEffectiveOrdersResponseBodyOrderSummaryOrderList) GetEndTime() *string {
	return s.EndTime
}

func (s *ListEffectiveOrdersResponseBodyOrderSummaryOrderList) GetInsNum() *string {
	return s.InsNum
}

func (s *ListEffectiveOrdersResponseBodyOrderSummaryOrderList) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListEffectiveOrdersResponseBodyOrderSummaryOrderList) GetOrderId() *string {
	return s.OrderId
}

func (s *ListEffectiveOrdersResponseBodyOrderSummaryOrderList) GetStartTime() *string {
	return s.StartTime
}

func (s *ListEffectiveOrdersResponseBodyOrderSummaryOrderList) SetBuyerId(v string) *ListEffectiveOrdersResponseBodyOrderSummaryOrderList {
	s.BuyerId = &v
	return s
}

func (s *ListEffectiveOrdersResponseBodyOrderSummaryOrderList) SetEndTime(v string) *ListEffectiveOrdersResponseBodyOrderSummaryOrderList {
	s.EndTime = &v
	return s
}

func (s *ListEffectiveOrdersResponseBodyOrderSummaryOrderList) SetInsNum(v string) *ListEffectiveOrdersResponseBodyOrderSummaryOrderList {
	s.InsNum = &v
	return s
}

func (s *ListEffectiveOrdersResponseBodyOrderSummaryOrderList) SetInstanceId(v string) *ListEffectiveOrdersResponseBodyOrderSummaryOrderList {
	s.InstanceId = &v
	return s
}

func (s *ListEffectiveOrdersResponseBodyOrderSummaryOrderList) SetOrderId(v string) *ListEffectiveOrdersResponseBodyOrderSummaryOrderList {
	s.OrderId = &v
	return s
}

func (s *ListEffectiveOrdersResponseBodyOrderSummaryOrderList) SetStartTime(v string) *ListEffectiveOrdersResponseBodyOrderSummaryOrderList {
	s.StartTime = &v
	return s
}

func (s *ListEffectiveOrdersResponseBodyOrderSummaryOrderList) Validate() error {
	return dara.Validate(s)
}
