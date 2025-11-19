// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePackageOrdersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribePackageOrdersResponseBody
	GetCode() *string
	SetMessage(v string) *DescribePackageOrdersResponseBody
	GetMessage() *string
	SetPage(v *DescribePackageOrdersResponseBodyPage) *DescribePackageOrdersResponseBody
	GetPage() *DescribePackageOrdersResponseBodyPage
	SetRequestId(v string) *DescribePackageOrdersResponseBody
	GetRequestId() *string
	SetTraceId(v string) *DescribePackageOrdersResponseBody
	GetTraceId() *string
}

type DescribePackageOrdersResponseBody struct {
	Code      *string                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string                                `json:"Message,omitempty" xml:"Message,omitempty"`
	Page      *DescribePackageOrdersResponseBodyPage `json:"Page,omitempty" xml:"Page,omitempty" type:"Struct"`
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TraceId   *string                                `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s DescribePackageOrdersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePackageOrdersResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePackageOrdersResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribePackageOrdersResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribePackageOrdersResponseBody) GetPage() *DescribePackageOrdersResponseBodyPage {
	return s.Page
}

func (s *DescribePackageOrdersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePackageOrdersResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *DescribePackageOrdersResponseBody) SetCode(v string) *DescribePackageOrdersResponseBody {
	s.Code = &v
	return s
}

func (s *DescribePackageOrdersResponseBody) SetMessage(v string) *DescribePackageOrdersResponseBody {
	s.Message = &v
	return s
}

func (s *DescribePackageOrdersResponseBody) SetPage(v *DescribePackageOrdersResponseBodyPage) *DescribePackageOrdersResponseBody {
	s.Page = v
	return s
}

func (s *DescribePackageOrdersResponseBody) SetRequestId(v string) *DescribePackageOrdersResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePackageOrdersResponseBody) SetTraceId(v string) *DescribePackageOrdersResponseBody {
	s.TraceId = &v
	return s
}

func (s *DescribePackageOrdersResponseBody) Validate() error {
	if s.Page != nil {
		if err := s.Page.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribePackageOrdersResponseBodyPage struct {
	CurrentPage *int64                                            `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	OrderList   []*DescribePackageOrdersResponseBodyPageOrderList `json:"OrderList,omitempty" xml:"OrderList,omitempty" type:"Repeated"`
	PageSize    *int64                                            `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount  *int64                                            `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribePackageOrdersResponseBodyPage) String() string {
	return dara.Prettify(s)
}

func (s DescribePackageOrdersResponseBodyPage) GoString() string {
	return s.String()
}

func (s *DescribePackageOrdersResponseBodyPage) GetCurrentPage() *int64 {
	return s.CurrentPage
}

func (s *DescribePackageOrdersResponseBodyPage) GetOrderList() []*DescribePackageOrdersResponseBodyPageOrderList {
	return s.OrderList
}

func (s *DescribePackageOrdersResponseBodyPage) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribePackageOrdersResponseBodyPage) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribePackageOrdersResponseBodyPage) SetCurrentPage(v int64) *DescribePackageOrdersResponseBodyPage {
	s.CurrentPage = &v
	return s
}

func (s *DescribePackageOrdersResponseBodyPage) SetOrderList(v []*DescribePackageOrdersResponseBodyPageOrderList) *DescribePackageOrdersResponseBodyPage {
	s.OrderList = v
	return s
}

func (s *DescribePackageOrdersResponseBodyPage) SetPageSize(v int64) *DescribePackageOrdersResponseBodyPage {
	s.PageSize = &v
	return s
}

func (s *DescribePackageOrdersResponseBodyPage) SetTotalCount(v int64) *DescribePackageOrdersResponseBodyPage {
	s.TotalCount = &v
	return s
}

func (s *DescribePackageOrdersResponseBodyPage) Validate() error {
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

type DescribePackageOrdersResponseBodyPageOrderList struct {
	Amount          *int32  `json:"Amount,omitempty" xml:"Amount,omitempty"`
	CreatorAliUid   *int64  `json:"CreatorAliUid,omitempty" xml:"CreatorAliUid,omitempty"`
	DeliveryAddress *string `json:"DeliveryAddress,omitempty" xml:"DeliveryAddress,omitempty"`
	DesktopId       *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty"`
	// Use the UTC time format: yyyy-MM-ddTHH:mm:ssZ
	GmtCanceledTime *string `json:"GmtCanceledTime,omitempty" xml:"GmtCanceledTime,omitempty"`
	// Use the UTC time format: yyyy-MM-ddTHH:mm:ssZ
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// Use the UTC time format: yyyy-MM-ddTHH:mm:ssZ
	GmtPaidTime    *string `json:"GmtPaidTime,omitempty" xml:"GmtPaidTime,omitempty"`
	InstanceId     *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OrderId        *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	OrderStatus    *string `json:"OrderStatus,omitempty" xml:"OrderStatus,omitempty"`
	OrderType      *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	ProductCode    *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	ProductSkuCode *string `json:"ProductSkuCode,omitempty" xml:"ProductSkuCode,omitempty"`
	ProductType    *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	TradePrice     *string `json:"TradePrice,omitempty" xml:"TradePrice,omitempty"`
}

func (s DescribePackageOrdersResponseBodyPageOrderList) String() string {
	return dara.Prettify(s)
}

func (s DescribePackageOrdersResponseBodyPageOrderList) GoString() string {
	return s.String()
}

func (s *DescribePackageOrdersResponseBodyPageOrderList) GetAmount() *int32 {
	return s.Amount
}

func (s *DescribePackageOrdersResponseBodyPageOrderList) GetCreatorAliUid() *int64 {
	return s.CreatorAliUid
}

func (s *DescribePackageOrdersResponseBodyPageOrderList) GetDeliveryAddress() *string {
	return s.DeliveryAddress
}

func (s *DescribePackageOrdersResponseBodyPageOrderList) GetDesktopId() *string {
	return s.DesktopId
}

func (s *DescribePackageOrdersResponseBodyPageOrderList) GetGmtCanceledTime() *string {
	return s.GmtCanceledTime
}

func (s *DescribePackageOrdersResponseBodyPageOrderList) GetGmtCreateTime() *string {
	return s.GmtCreateTime
}

func (s *DescribePackageOrdersResponseBodyPageOrderList) GetGmtPaidTime() *string {
	return s.GmtPaidTime
}

func (s *DescribePackageOrdersResponseBodyPageOrderList) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribePackageOrdersResponseBodyPageOrderList) GetOrderId() *string {
	return s.OrderId
}

func (s *DescribePackageOrdersResponseBodyPageOrderList) GetOrderStatus() *string {
	return s.OrderStatus
}

func (s *DescribePackageOrdersResponseBodyPageOrderList) GetOrderType() *string {
	return s.OrderType
}

func (s *DescribePackageOrdersResponseBodyPageOrderList) GetProductCode() *string {
	return s.ProductCode
}

func (s *DescribePackageOrdersResponseBodyPageOrderList) GetProductSkuCode() *string {
	return s.ProductSkuCode
}

func (s *DescribePackageOrdersResponseBodyPageOrderList) GetProductType() *string {
	return s.ProductType
}

func (s *DescribePackageOrdersResponseBodyPageOrderList) GetTradePrice() *string {
	return s.TradePrice
}

func (s *DescribePackageOrdersResponseBodyPageOrderList) SetAmount(v int32) *DescribePackageOrdersResponseBodyPageOrderList {
	s.Amount = &v
	return s
}

func (s *DescribePackageOrdersResponseBodyPageOrderList) SetCreatorAliUid(v int64) *DescribePackageOrdersResponseBodyPageOrderList {
	s.CreatorAliUid = &v
	return s
}

func (s *DescribePackageOrdersResponseBodyPageOrderList) SetDeliveryAddress(v string) *DescribePackageOrdersResponseBodyPageOrderList {
	s.DeliveryAddress = &v
	return s
}

func (s *DescribePackageOrdersResponseBodyPageOrderList) SetDesktopId(v string) *DescribePackageOrdersResponseBodyPageOrderList {
	s.DesktopId = &v
	return s
}

func (s *DescribePackageOrdersResponseBodyPageOrderList) SetGmtCanceledTime(v string) *DescribePackageOrdersResponseBodyPageOrderList {
	s.GmtCanceledTime = &v
	return s
}

func (s *DescribePackageOrdersResponseBodyPageOrderList) SetGmtCreateTime(v string) *DescribePackageOrdersResponseBodyPageOrderList {
	s.GmtCreateTime = &v
	return s
}

func (s *DescribePackageOrdersResponseBodyPageOrderList) SetGmtPaidTime(v string) *DescribePackageOrdersResponseBodyPageOrderList {
	s.GmtPaidTime = &v
	return s
}

func (s *DescribePackageOrdersResponseBodyPageOrderList) SetInstanceId(v string) *DescribePackageOrdersResponseBodyPageOrderList {
	s.InstanceId = &v
	return s
}

func (s *DescribePackageOrdersResponseBodyPageOrderList) SetOrderId(v string) *DescribePackageOrdersResponseBodyPageOrderList {
	s.OrderId = &v
	return s
}

func (s *DescribePackageOrdersResponseBodyPageOrderList) SetOrderStatus(v string) *DescribePackageOrdersResponseBodyPageOrderList {
	s.OrderStatus = &v
	return s
}

func (s *DescribePackageOrdersResponseBodyPageOrderList) SetOrderType(v string) *DescribePackageOrdersResponseBodyPageOrderList {
	s.OrderType = &v
	return s
}

func (s *DescribePackageOrdersResponseBodyPageOrderList) SetProductCode(v string) *DescribePackageOrdersResponseBodyPageOrderList {
	s.ProductCode = &v
	return s
}

func (s *DescribePackageOrdersResponseBodyPageOrderList) SetProductSkuCode(v string) *DescribePackageOrdersResponseBodyPageOrderList {
	s.ProductSkuCode = &v
	return s
}

func (s *DescribePackageOrdersResponseBodyPageOrderList) SetProductType(v string) *DescribePackageOrdersResponseBodyPageOrderList {
	s.ProductType = &v
	return s
}

func (s *DescribePackageOrdersResponseBodyPageOrderList) SetTradePrice(v string) *DescribePackageOrdersResponseBodyPageOrderList {
	s.TradePrice = &v
	return s
}

func (s *DescribePackageOrdersResponseBodyPageOrderList) Validate() error {
	return dara.Validate(s)
}
