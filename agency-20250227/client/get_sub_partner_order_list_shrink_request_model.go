// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSubPartnerOrderListShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOrderCreateAfter(v int64) *GetSubPartnerOrderListShrinkRequest
	GetOrderCreateAfter() *int64
	SetOrderCreateBefore(v int64) *GetSubPartnerOrderListShrinkRequest
	GetOrderCreateBefore() *int64
	SetOrderId(v int64) *GetSubPartnerOrderListShrinkRequest
	GetOrderId() *int64
	SetOrderPayAfter(v int64) *GetSubPartnerOrderListShrinkRequest
	GetOrderPayAfter() *int64
	SetOrderPayBefore(v int64) *GetSubPartnerOrderListShrinkRequest
	GetOrderPayBefore() *int64
	SetOrderStatus(v int64) *GetSubPartnerOrderListShrinkRequest
	GetOrderStatus() *int64
	SetOrderTypeListShrink(v string) *GetSubPartnerOrderListShrinkRequest
	GetOrderTypeListShrink() *string
	SetPageNo(v int32) *GetSubPartnerOrderListShrinkRequest
	GetPageNo() *int32
	SetPageSize(v int32) *GetSubPartnerOrderListShrinkRequest
	GetPageSize() *int32
	SetPayAmountAfter(v int64) *GetSubPartnerOrderListShrinkRequest
	GetPayAmountAfter() *int64
	SetPayAmountBefore(v int64) *GetSubPartnerOrderListShrinkRequest
	GetPayAmountBefore() *int64
	SetPayType(v int64) *GetSubPartnerOrderListShrinkRequest
	GetPayType() *int64
	SetProductCode(v string) *GetSubPartnerOrderListShrinkRequest
	GetProductCode() *string
	SetProductName(v string) *GetSubPartnerOrderListShrinkRequest
	GetProductName() *string
	SetProjectId(v int64) *GetSubPartnerOrderListShrinkRequest
	GetProjectId() *int64
	SetSubPartnerName(v string) *GetSubPartnerOrderListShrinkRequest
	GetSubPartnerName() *string
	SetSubPartnerUid(v int64) *GetSubPartnerOrderListShrinkRequest
	GetSubPartnerUid() *int64
}

type GetSubPartnerOrderListShrinkRequest struct {
	// The UNIX timestamp of the start time for order creation. The time range must not exceed six months.
	//
	// The time range for order creation and the time range for order payment cannot both be empty.
	//
	// example:
	//
	// 1727789348000
	OrderCreateAfter *int64 `json:"OrderCreateAfter,omitempty" xml:"OrderCreateAfter,omitempty"`
	// The UNIX timestamp of the end time for order creation. The time range must not exceed six months.
	//
	// The time range for order creation and the time range for order payment cannot both be empty.
	//
	// example:
	//
	// 1741008566000
	OrderCreateBefore *int64 `json:"OrderCreateBefore,omitempty" xml:"OrderCreateBefore,omitempty"`
	// Order ID
	//
	// example:
	//
	// 209335720330622
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The UNIX timestamp of the start time for order payment. The time range must not exceed six months.
	//
	// The time range for order creation and the time range for order payment cannot both be empty.
	//
	// example:
	//
	// 1727789348000
	OrderPayAfter *int64 `json:"OrderPayAfter,omitempty" xml:"OrderPayAfter,omitempty"`
	// UNIX timestamp of the order payment deadline. The time range cannot exceed six months.
	//
	// The time range for order creation and the time range for order payment cannot both be empty.
	//
	// example:
	//
	// 1741008566000
	OrderPayBefore *int64 `json:"OrderPayBefore,omitempty" xml:"OrderPayBefore,omitempty"`
	// Order status. Valid values:
	//
	// 1: Unpaid
	//
	// 2: Paid
	//
	// 3: Canceled
	//
	// example:
	//
	// 3
	OrderStatus *int64 `json:"OrderStatus,omitempty" xml:"OrderStatus,omitempty"`
	// List of order types.
	//
	// Valid order types include: BUY, UPGRADE, DOWNGRADE, RENEW, REFUND, OTHERS
	OrderTypeListShrink *string `json:"OrderTypeList,omitempty" xml:"OrderTypeList,omitempty"`
	// Page number
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// Page size
	//
	// This parameter is required.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Minimum paid amount
	//
	// example:
	//
	// 1
	PayAmountAfter *int64 `json:"PayAmountAfter,omitempty" xml:"PayAmountAfter,omitempty"`
	// Maximum paid amount
	//
	// example:
	//
	// 100
	PayAmountBefore *int64 `json:"PayAmountBefore,omitempty" xml:"PayAmountBefore,omitempty"`
	// Payment Type:
	//
	// 1: Non-agent payment
	//
	// 2: Agent payment
	//
	// example:
	//
	// 1
	PayType *int64 `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// Product code
	//
	// example:
	//
	// ecs
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// Product Name
	//
	// example:
	//
	// 弹性计算
	ProductName *string `json:"ProductName,omitempty" xml:"ProductName,omitempty"`
	// Opportunity ID
	//
	// example:
	//
	// 202501101023
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// Sub-partner name
	//
	// example:
	//
	// XXX有限公司
	SubPartnerName *string `json:"SubPartnerName,omitempty" xml:"SubPartnerName,omitempty"`
	// Sub-partner UID
	//
	// example:
	//
	// 123432311
	SubPartnerUid *int64 `json:"SubPartnerUid,omitempty" xml:"SubPartnerUid,omitempty"`
}

func (s GetSubPartnerOrderListShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSubPartnerOrderListShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetSubPartnerOrderListShrinkRequest) GetOrderCreateAfter() *int64 {
	return s.OrderCreateAfter
}

func (s *GetSubPartnerOrderListShrinkRequest) GetOrderCreateBefore() *int64 {
	return s.OrderCreateBefore
}

func (s *GetSubPartnerOrderListShrinkRequest) GetOrderId() *int64 {
	return s.OrderId
}

func (s *GetSubPartnerOrderListShrinkRequest) GetOrderPayAfter() *int64 {
	return s.OrderPayAfter
}

func (s *GetSubPartnerOrderListShrinkRequest) GetOrderPayBefore() *int64 {
	return s.OrderPayBefore
}

func (s *GetSubPartnerOrderListShrinkRequest) GetOrderStatus() *int64 {
	return s.OrderStatus
}

func (s *GetSubPartnerOrderListShrinkRequest) GetOrderTypeListShrink() *string {
	return s.OrderTypeListShrink
}

func (s *GetSubPartnerOrderListShrinkRequest) GetPageNo() *int32 {
	return s.PageNo
}

func (s *GetSubPartnerOrderListShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetSubPartnerOrderListShrinkRequest) GetPayAmountAfter() *int64 {
	return s.PayAmountAfter
}

func (s *GetSubPartnerOrderListShrinkRequest) GetPayAmountBefore() *int64 {
	return s.PayAmountBefore
}

func (s *GetSubPartnerOrderListShrinkRequest) GetPayType() *int64 {
	return s.PayType
}

func (s *GetSubPartnerOrderListShrinkRequest) GetProductCode() *string {
	return s.ProductCode
}

func (s *GetSubPartnerOrderListShrinkRequest) GetProductName() *string {
	return s.ProductName
}

func (s *GetSubPartnerOrderListShrinkRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *GetSubPartnerOrderListShrinkRequest) GetSubPartnerName() *string {
	return s.SubPartnerName
}

func (s *GetSubPartnerOrderListShrinkRequest) GetSubPartnerUid() *int64 {
	return s.SubPartnerUid
}

func (s *GetSubPartnerOrderListShrinkRequest) SetOrderCreateAfter(v int64) *GetSubPartnerOrderListShrinkRequest {
	s.OrderCreateAfter = &v
	return s
}

func (s *GetSubPartnerOrderListShrinkRequest) SetOrderCreateBefore(v int64) *GetSubPartnerOrderListShrinkRequest {
	s.OrderCreateBefore = &v
	return s
}

func (s *GetSubPartnerOrderListShrinkRequest) SetOrderId(v int64) *GetSubPartnerOrderListShrinkRequest {
	s.OrderId = &v
	return s
}

func (s *GetSubPartnerOrderListShrinkRequest) SetOrderPayAfter(v int64) *GetSubPartnerOrderListShrinkRequest {
	s.OrderPayAfter = &v
	return s
}

func (s *GetSubPartnerOrderListShrinkRequest) SetOrderPayBefore(v int64) *GetSubPartnerOrderListShrinkRequest {
	s.OrderPayBefore = &v
	return s
}

func (s *GetSubPartnerOrderListShrinkRequest) SetOrderStatus(v int64) *GetSubPartnerOrderListShrinkRequest {
	s.OrderStatus = &v
	return s
}

func (s *GetSubPartnerOrderListShrinkRequest) SetOrderTypeListShrink(v string) *GetSubPartnerOrderListShrinkRequest {
	s.OrderTypeListShrink = &v
	return s
}

func (s *GetSubPartnerOrderListShrinkRequest) SetPageNo(v int32) *GetSubPartnerOrderListShrinkRequest {
	s.PageNo = &v
	return s
}

func (s *GetSubPartnerOrderListShrinkRequest) SetPageSize(v int32) *GetSubPartnerOrderListShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *GetSubPartnerOrderListShrinkRequest) SetPayAmountAfter(v int64) *GetSubPartnerOrderListShrinkRequest {
	s.PayAmountAfter = &v
	return s
}

func (s *GetSubPartnerOrderListShrinkRequest) SetPayAmountBefore(v int64) *GetSubPartnerOrderListShrinkRequest {
	s.PayAmountBefore = &v
	return s
}

func (s *GetSubPartnerOrderListShrinkRequest) SetPayType(v int64) *GetSubPartnerOrderListShrinkRequest {
	s.PayType = &v
	return s
}

func (s *GetSubPartnerOrderListShrinkRequest) SetProductCode(v string) *GetSubPartnerOrderListShrinkRequest {
	s.ProductCode = &v
	return s
}

func (s *GetSubPartnerOrderListShrinkRequest) SetProductName(v string) *GetSubPartnerOrderListShrinkRequest {
	s.ProductName = &v
	return s
}

func (s *GetSubPartnerOrderListShrinkRequest) SetProjectId(v int64) *GetSubPartnerOrderListShrinkRequest {
	s.ProjectId = &v
	return s
}

func (s *GetSubPartnerOrderListShrinkRequest) SetSubPartnerName(v string) *GetSubPartnerOrderListShrinkRequest {
	s.SubPartnerName = &v
	return s
}

func (s *GetSubPartnerOrderListShrinkRequest) SetSubPartnerUid(v int64) *GetSubPartnerOrderListShrinkRequest {
	s.SubPartnerUid = &v
	return s
}

func (s *GetSubPartnerOrderListShrinkRequest) Validate() error {
	return dara.Validate(s)
}
