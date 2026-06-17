// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOrdersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListOrdersResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListOrdersResponseBody
	GetNextToken() *string
	SetOrderList(v []*ListOrdersResponseBodyOrderList) *ListOrdersResponseBody
	GetOrderList() []*ListOrdersResponseBodyOrderList
	SetPageNumber(v int32) *ListOrdersResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListOrdersResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListOrdersResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListOrdersResponseBody
	GetTotalCount() *int32
}

type ListOrdersResponseBody struct {
	// The maximum number of entries returned for the current request. Default value: 10.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// A pagination token. If the query results are not returned in a single call, this token is returned. Use this token in a subsequent call to retrieve the remaining results.
	//
	// example:
	//
	// 212db86sca4384811e0b5e8707e******
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The list of orders.
	//
	// This parameter is required.
	OrderList []*ListOrdersResponseBodyOrderList `json:"OrderList,omitempty" xml:"OrderList,omitempty" type:"Repeated"`
	// The page number of the returned page. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page. Valid values:
	//
	// - **30**
	//
	// - **50**
	//
	// - **100**
	//
	// Default value: 30.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 3AA69096-757C-4647-B36C-29EBC2******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 7
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListOrdersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListOrdersResponseBody) GoString() string {
	return s.String()
}

func (s *ListOrdersResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListOrdersResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListOrdersResponseBody) GetOrderList() []*ListOrdersResponseBodyOrderList {
	return s.OrderList
}

func (s *ListOrdersResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListOrdersResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListOrdersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListOrdersResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListOrdersResponseBody) SetMaxResults(v int32) *ListOrdersResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListOrdersResponseBody) SetNextToken(v string) *ListOrdersResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListOrdersResponseBody) SetOrderList(v []*ListOrdersResponseBodyOrderList) *ListOrdersResponseBody {
	s.OrderList = v
	return s
}

func (s *ListOrdersResponseBody) SetPageNumber(v int32) *ListOrdersResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListOrdersResponseBody) SetPageSize(v int32) *ListOrdersResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListOrdersResponseBody) SetRequestId(v string) *ListOrdersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListOrdersResponseBody) SetTotalCount(v int32) *ListOrdersResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListOrdersResponseBody) Validate() error {
	if s.OrderList != nil {
		for _, item := range s.OrderList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListOrdersResponseBodyOrderList struct {
	// The ID of the Alibaba Cloud account.
	//
	// example:
	//
	// 1910740440664****
	AliUid *string `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	// The billing method of the instance. Valid values:
	//
	// - **Postpaid**: pay-as-you-go
	//
	// - **Prepaid**: subscription
	//
	// example:
	//
	// Prepaid
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// The commodity code. Valid values:
	//
	// - polardb_sub: subscription in the Chinese mainland.
	//
	// - polardb_sub_intl: subscription in Hong Kong (China) and regions outside China.
	//
	// - polardb_payg: pay-as-you-go in the Chinese mainland.
	//
	// - polardb_payg_intl: pay-as-you-go in Hong Kong (China) and regions outside China.
	//
	// - polardb_sub_jushita: Jushita subscription.
	//
	// - polardb_payg_jushita: Jushita pay-as-you-go.
	//
	// - polardb_sub_cainiao: Cainiao subscription.
	//
	// - polardb_payg_cainiao: Cainiao pay-as-you-go.
	//
	// > 	- If you use an Alibaba Cloud account for the China site, you can view only the commodity codes for the Chinese mainland.
	//
	// >
	//
	// > 	- If you use an Alibaba Cloud international site account, you can view only the commodity codes for regions outside the Chinese mainland.
	//
	// >
	//
	// > 	- If you use a Jushita account, you can view only the commodity codes for Jushita.
	//
	// >
	//
	// > 	- If you use a Cainiao account, you can view only the commodity codes for Cainiao.
	//
	// example:
	//
	// polardb_payg_intl
	CommodityCode *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	// The time when the order was created.
	//
	// example:
	//
	// 2021-03-31T16:09:13
	CreatedTime *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// The cluster ID.
	//
	// example:
	//
	// pc-uf6k532gav*******
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The order ID.
	//
	// example:
	//
	// 25808743077*****
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The status of the order.
	//
	// - **pending**: The task is waiting to start.
	//
	// - **create**: The order is placed and is being processed.
	//
	// - **fail**: The instance failed to be created.
	//
	// - **cancel**: The order is canceled.
	//
	// - **success**: The instance is created.
	//
	// example:
	//
	// success
	OrderStatus *string `json:"OrderStatus,omitempty" xml:"OrderStatus,omitempty"`
	// The type of the order. Valid values:
	//
	// - BUY: The instance is purchased.
	//
	// - UPGRADE: The instance configuration is changed.
	//
	// - RENEW: The subscription is renewed.
	//
	// - CONVERT: The billing method is changed.
	//
	// example:
	//
	// BUY
	OrderType *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	// The product code.
	//
	// example:
	//
	// polardb
	ProduceCode *string `json:"ProduceCode,omitempty" xml:"ProduceCode,omitempty"`
	// The region information
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
}

func (s ListOrdersResponseBodyOrderList) String() string {
	return dara.Prettify(s)
}

func (s ListOrdersResponseBodyOrderList) GoString() string {
	return s.String()
}

func (s *ListOrdersResponseBodyOrderList) GetAliUid() *string {
	return s.AliUid
}

func (s *ListOrdersResponseBodyOrderList) GetChargeType() *string {
	return s.ChargeType
}

func (s *ListOrdersResponseBodyOrderList) GetCommodityCode() *string {
	return s.CommodityCode
}

func (s *ListOrdersResponseBodyOrderList) GetCreatedTime() *string {
	return s.CreatedTime
}

func (s *ListOrdersResponseBodyOrderList) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListOrdersResponseBodyOrderList) GetOrderId() *string {
	return s.OrderId
}

func (s *ListOrdersResponseBodyOrderList) GetOrderStatus() *string {
	return s.OrderStatus
}

func (s *ListOrdersResponseBodyOrderList) GetOrderType() *string {
	return s.OrderType
}

func (s *ListOrdersResponseBodyOrderList) GetProduceCode() *string {
	return s.ProduceCode
}

func (s *ListOrdersResponseBodyOrderList) GetRegion() *string {
	return s.Region
}

func (s *ListOrdersResponseBodyOrderList) SetAliUid(v string) *ListOrdersResponseBodyOrderList {
	s.AliUid = &v
	return s
}

func (s *ListOrdersResponseBodyOrderList) SetChargeType(v string) *ListOrdersResponseBodyOrderList {
	s.ChargeType = &v
	return s
}

func (s *ListOrdersResponseBodyOrderList) SetCommodityCode(v string) *ListOrdersResponseBodyOrderList {
	s.CommodityCode = &v
	return s
}

func (s *ListOrdersResponseBodyOrderList) SetCreatedTime(v string) *ListOrdersResponseBodyOrderList {
	s.CreatedTime = &v
	return s
}

func (s *ListOrdersResponseBodyOrderList) SetInstanceId(v string) *ListOrdersResponseBodyOrderList {
	s.InstanceId = &v
	return s
}

func (s *ListOrdersResponseBodyOrderList) SetOrderId(v string) *ListOrdersResponseBodyOrderList {
	s.OrderId = &v
	return s
}

func (s *ListOrdersResponseBodyOrderList) SetOrderStatus(v string) *ListOrdersResponseBodyOrderList {
	s.OrderStatus = &v
	return s
}

func (s *ListOrdersResponseBodyOrderList) SetOrderType(v string) *ListOrdersResponseBodyOrderList {
	s.OrderType = &v
	return s
}

func (s *ListOrdersResponseBodyOrderList) SetProduceCode(v string) *ListOrdersResponseBodyOrderList {
	s.ProduceCode = &v
	return s
}

func (s *ListOrdersResponseBodyOrderList) SetRegion(v string) *ListOrdersResponseBodyOrderList {
	s.Region = &v
	return s
}

func (s *ListOrdersResponseBodyOrderList) Validate() error {
	return dara.Validate(s)
}
