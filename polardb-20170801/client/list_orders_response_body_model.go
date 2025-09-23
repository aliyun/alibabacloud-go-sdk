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
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// 212db86sca4384811e0b5e8707e******
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// This parameter is required.
	OrderList []*ListOrdersResponseBodyOrderList `json:"OrderList,omitempty" xml:"OrderList,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 3AA69096-757C-4647-B36C-29EBC2******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	return dara.Validate(s)
}

type ListOrdersResponseBodyOrderList struct {
	// example:
	//
	// 1910740440664****
	AliUid *string `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	// example:
	//
	// Prepaid
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// example:
	//
	// polardb_payg_intl
	CommodityCode *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	// example:
	//
	// 2021-03-31T16:09:13
	CreatedTime *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// example:
	//
	// pc-uf6k532gav*******
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 25808743077*****
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// example:
	//
	// completed
	OrderStatus *string `json:"OrderStatus,omitempty" xml:"OrderStatus,omitempty"`
	// example:
	//
	// BUY
	OrderType *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	// example:
	//
	// polardb
	ProduceCode *string `json:"ProduceCode,omitempty" xml:"ProduceCode,omitempty"`
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
