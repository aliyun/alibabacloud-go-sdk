// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSubPartnerOrderListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOrderCreateAfter(v int64) *GetSubPartnerOrderListRequest
	GetOrderCreateAfter() *int64
	SetOrderCreateBefore(v int64) *GetSubPartnerOrderListRequest
	GetOrderCreateBefore() *int64
	SetOrderId(v int64) *GetSubPartnerOrderListRequest
	GetOrderId() *int64
	SetOrderPayAfter(v int64) *GetSubPartnerOrderListRequest
	GetOrderPayAfter() *int64
	SetOrderPayBefore(v int64) *GetSubPartnerOrderListRequest
	GetOrderPayBefore() *int64
	SetOrderStatus(v int64) *GetSubPartnerOrderListRequest
	GetOrderStatus() *int64
	SetOrderTypeList(v []*string) *GetSubPartnerOrderListRequest
	GetOrderTypeList() []*string
	SetPageNo(v int32) *GetSubPartnerOrderListRequest
	GetPageNo() *int32
	SetPageSize(v int32) *GetSubPartnerOrderListRequest
	GetPageSize() *int32
	SetPayAmountAfter(v int64) *GetSubPartnerOrderListRequest
	GetPayAmountAfter() *int64
	SetPayAmountBefore(v int64) *GetSubPartnerOrderListRequest
	GetPayAmountBefore() *int64
	SetPayType(v int64) *GetSubPartnerOrderListRequest
	GetPayType() *int64
	SetProductCode(v string) *GetSubPartnerOrderListRequest
	GetProductCode() *string
	SetProductName(v string) *GetSubPartnerOrderListRequest
	GetProductName() *string
	SetProjectId(v int64) *GetSubPartnerOrderListRequest
	GetProjectId() *int64
	SetSubPartnerName(v string) *GetSubPartnerOrderListRequest
	GetSubPartnerName() *string
	SetSubPartnerUid(v int64) *GetSubPartnerOrderListRequest
	GetSubPartnerUid() *int64
}

type GetSubPartnerOrderListRequest struct {
	// The start timestamp of the order creation time range. The time range cannot exceed 6 months. The order creation time range and the order payment time range cannot both be empty.
	//
	// example:
	//
	// 1727789348000
	OrderCreateAfter *int64 `json:"OrderCreateAfter,omitempty" xml:"OrderCreateAfter,omitempty"`
	// The end timestamp of the order creation time range. The time range cannot exceed 6 months. The order creation time range and the order payment time range cannot both be empty.
	//
	// example:
	//
	// 1741008566000
	OrderCreateBefore *int64 `json:"OrderCreateBefore,omitempty" xml:"OrderCreateBefore,omitempty"`
	// The order ID.
	//
	// example:
	//
	// 209335720330622
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The start timestamp of the order payment time range. The time range cannot exceed 6 months. The order creation time range and the order payment time range cannot both be empty.
	//
	// example:
	//
	// 1727789348000
	OrderPayAfter *int64 `json:"OrderPayAfter,omitempty" xml:"OrderPayAfter,omitempty"`
	// The end timestamp of the order payment time range. The time range cannot exceed 6 months. The order creation time range and the order payment time range cannot both be empty.
	//
	// example:
	//
	// 1741008566000
	OrderPayBefore *int64 `json:"OrderPayBefore,omitempty" xml:"OrderPayBefore,omitempty"`
	// The order status. Valid values:
	//
	// - 1: unpaid
	//
	// - 2: paid
	//
	// - 3: voided.
	//
	// example:
	//
	// 3
	OrderStatus *int64 `json:"OrderStatus,omitempty" xml:"OrderStatus,omitempty"`
	// The list of order types.
	//
	// Valid values: BUY, UPGRADE, DOWNGRADE, RENEW, REFUND, and OTHERS.
	OrderTypeList []*string `json:"OrderTypeList,omitempty" xml:"OrderTypeList,omitempty" type:"Repeated"`
	// The page number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page.
	//
	// This parameter is required.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The minimum actual payment amount.
	//
	// example:
	//
	// 1
	PayAmountAfter *int64 `json:"PayAmountAfter,omitempty" xml:"PayAmountAfter,omitempty"`
	// The maximum actual payment amount.
	//
	// example:
	//
	// 100
	PayAmountBefore *int64 `json:"PayAmountBefore,omitempty" xml:"PayAmountBefore,omitempty"`
	// The payment type. Valid values:
	//
	// - 1: non-delegated payment
	//
	// - 2: delegated payment.
	//
	// example:
	//
	// 1
	PayType *int64 `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// The product code.
	//
	// example:
	//
	// ecs
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The product name.
	//
	// example:
	//
	// 弹性计算
	ProductName *string `json:"ProductName,omitempty" xml:"ProductName,omitempty"`
	// The opportunity ID.
	//
	// example:
	//
	// 202501101023
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The name of the secondary partner.
	//
	// example:
	//
	// XXX有限公司
	SubPartnerName *string `json:"SubPartnerName,omitempty" xml:"SubPartnerName,omitempty"`
	// The UID of the secondary partner.
	//
	// example:
	//
	// 123432311
	SubPartnerUid *int64 `json:"SubPartnerUid,omitempty" xml:"SubPartnerUid,omitempty"`
}

func (s GetSubPartnerOrderListRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSubPartnerOrderListRequest) GoString() string {
	return s.String()
}

func (s *GetSubPartnerOrderListRequest) GetOrderCreateAfter() *int64 {
	return s.OrderCreateAfter
}

func (s *GetSubPartnerOrderListRequest) GetOrderCreateBefore() *int64 {
	return s.OrderCreateBefore
}

func (s *GetSubPartnerOrderListRequest) GetOrderId() *int64 {
	return s.OrderId
}

func (s *GetSubPartnerOrderListRequest) GetOrderPayAfter() *int64 {
	return s.OrderPayAfter
}

func (s *GetSubPartnerOrderListRequest) GetOrderPayBefore() *int64 {
	return s.OrderPayBefore
}

func (s *GetSubPartnerOrderListRequest) GetOrderStatus() *int64 {
	return s.OrderStatus
}

func (s *GetSubPartnerOrderListRequest) GetOrderTypeList() []*string {
	return s.OrderTypeList
}

func (s *GetSubPartnerOrderListRequest) GetPageNo() *int32 {
	return s.PageNo
}

func (s *GetSubPartnerOrderListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetSubPartnerOrderListRequest) GetPayAmountAfter() *int64 {
	return s.PayAmountAfter
}

func (s *GetSubPartnerOrderListRequest) GetPayAmountBefore() *int64 {
	return s.PayAmountBefore
}

func (s *GetSubPartnerOrderListRequest) GetPayType() *int64 {
	return s.PayType
}

func (s *GetSubPartnerOrderListRequest) GetProductCode() *string {
	return s.ProductCode
}

func (s *GetSubPartnerOrderListRequest) GetProductName() *string {
	return s.ProductName
}

func (s *GetSubPartnerOrderListRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *GetSubPartnerOrderListRequest) GetSubPartnerName() *string {
	return s.SubPartnerName
}

func (s *GetSubPartnerOrderListRequest) GetSubPartnerUid() *int64 {
	return s.SubPartnerUid
}

func (s *GetSubPartnerOrderListRequest) SetOrderCreateAfter(v int64) *GetSubPartnerOrderListRequest {
	s.OrderCreateAfter = &v
	return s
}

func (s *GetSubPartnerOrderListRequest) SetOrderCreateBefore(v int64) *GetSubPartnerOrderListRequest {
	s.OrderCreateBefore = &v
	return s
}

func (s *GetSubPartnerOrderListRequest) SetOrderId(v int64) *GetSubPartnerOrderListRequest {
	s.OrderId = &v
	return s
}

func (s *GetSubPartnerOrderListRequest) SetOrderPayAfter(v int64) *GetSubPartnerOrderListRequest {
	s.OrderPayAfter = &v
	return s
}

func (s *GetSubPartnerOrderListRequest) SetOrderPayBefore(v int64) *GetSubPartnerOrderListRequest {
	s.OrderPayBefore = &v
	return s
}

func (s *GetSubPartnerOrderListRequest) SetOrderStatus(v int64) *GetSubPartnerOrderListRequest {
	s.OrderStatus = &v
	return s
}

func (s *GetSubPartnerOrderListRequest) SetOrderTypeList(v []*string) *GetSubPartnerOrderListRequest {
	s.OrderTypeList = v
	return s
}

func (s *GetSubPartnerOrderListRequest) SetPageNo(v int32) *GetSubPartnerOrderListRequest {
	s.PageNo = &v
	return s
}

func (s *GetSubPartnerOrderListRequest) SetPageSize(v int32) *GetSubPartnerOrderListRequest {
	s.PageSize = &v
	return s
}

func (s *GetSubPartnerOrderListRequest) SetPayAmountAfter(v int64) *GetSubPartnerOrderListRequest {
	s.PayAmountAfter = &v
	return s
}

func (s *GetSubPartnerOrderListRequest) SetPayAmountBefore(v int64) *GetSubPartnerOrderListRequest {
	s.PayAmountBefore = &v
	return s
}

func (s *GetSubPartnerOrderListRequest) SetPayType(v int64) *GetSubPartnerOrderListRequest {
	s.PayType = &v
	return s
}

func (s *GetSubPartnerOrderListRequest) SetProductCode(v string) *GetSubPartnerOrderListRequest {
	s.ProductCode = &v
	return s
}

func (s *GetSubPartnerOrderListRequest) SetProductName(v string) *GetSubPartnerOrderListRequest {
	s.ProductName = &v
	return s
}

func (s *GetSubPartnerOrderListRequest) SetProjectId(v int64) *GetSubPartnerOrderListRequest {
	s.ProjectId = &v
	return s
}

func (s *GetSubPartnerOrderListRequest) SetSubPartnerName(v string) *GetSubPartnerOrderListRequest {
	s.SubPartnerName = &v
	return s
}

func (s *GetSubPartnerOrderListRequest) SetSubPartnerUid(v int64) *GetSubPartnerOrderListRequest {
	s.SubPartnerUid = &v
	return s
}

func (s *GetSubPartnerOrderListRequest) Validate() error {
	return dara.Validate(s)
}
