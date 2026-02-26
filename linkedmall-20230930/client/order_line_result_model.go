// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOrderLineResult interface {
	dara.Model
	String() string
	GoString() string
	SetEticketInfos(v []*EticketInfo) *OrderLineResult
	GetEticketInfos() []*EticketInfo
	SetLogisticsStatus(v string) *OrderLineResult
	GetLogisticsStatus() *string
	SetNumber(v string) *OrderLineResult
	GetNumber() *string
	SetOrderId(v string) *OrderLineResult
	GetOrderId() *string
	SetOrderLineId(v string) *OrderLineResult
	GetOrderLineId() *string
	SetOrderLineStatus(v string) *OrderLineResult
	GetOrderLineStatus() *string
	SetPayFee(v int64) *OrderLineResult
	GetPayFee() *int64
	SetProductId(v string) *OrderLineResult
	GetProductId() *string
	SetProductPic(v string) *OrderLineResult
	GetProductPic() *string
	SetProductTitle(v string) *OrderLineResult
	GetProductTitle() *string
	SetSkuId(v string) *OrderLineResult
	GetSkuId() *string
	SetSkuTitle(v string) *OrderLineResult
	GetSkuTitle() *string
}

type OrderLineResult struct {
	EticketInfos []*EticketInfo `json:"eticketInfos,omitempty" xml:"eticketInfos,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	LogisticsStatus *string `json:"logisticsStatus,omitempty" xml:"logisticsStatus,omitempty"`
	// example:
	//
	// 1
	Number *string `json:"number,omitempty" xml:"number,omitempty"`
	// example:
	//
	// 6692****5457
	OrderId *string `json:"orderId,omitempty" xml:"orderId,omitempty"`
	// example:
	//
	// 6692****5458
	OrderLineId *string `json:"orderLineId,omitempty" xml:"orderLineId,omitempty"`
	// example:
	//
	// 1
	OrderLineStatus *string `json:"orderLineStatus,omitempty" xml:"orderLineStatus,omitempty"`
	// example:
	//
	// 100
	PayFee *int64 `json:"payFee,omitempty" xml:"payFee,omitempty"`
	// example:
	//
	// 6600****6736
	ProductId *string `json:"productId,omitempty" xml:"productId,omitempty"`
	// example:
	//
	// //img.alicdn.com/imgextra/i4/2216003305543/O1CN01bip3Un1qokG0
	ProductPic   *string `json:"productPic,omitempty" xml:"productPic,omitempty"`
	ProductTitle *string `json:"productTitle,omitempty" xml:"productTitle,omitempty"`
	// skuId
	//
	// example:
	//
	// 6600****6737
	SkuId    *string `json:"skuId,omitempty" xml:"skuId,omitempty"`
	SkuTitle *string `json:"skuTitle,omitempty" xml:"skuTitle,omitempty"`
}

func (s OrderLineResult) String() string {
	return dara.Prettify(s)
}

func (s OrderLineResult) GoString() string {
	return s.String()
}

func (s *OrderLineResult) GetEticketInfos() []*EticketInfo {
	return s.EticketInfos
}

func (s *OrderLineResult) GetLogisticsStatus() *string {
	return s.LogisticsStatus
}

func (s *OrderLineResult) GetNumber() *string {
	return s.Number
}

func (s *OrderLineResult) GetOrderId() *string {
	return s.OrderId
}

func (s *OrderLineResult) GetOrderLineId() *string {
	return s.OrderLineId
}

func (s *OrderLineResult) GetOrderLineStatus() *string {
	return s.OrderLineStatus
}

func (s *OrderLineResult) GetPayFee() *int64 {
	return s.PayFee
}

func (s *OrderLineResult) GetProductId() *string {
	return s.ProductId
}

func (s *OrderLineResult) GetProductPic() *string {
	return s.ProductPic
}

func (s *OrderLineResult) GetProductTitle() *string {
	return s.ProductTitle
}

func (s *OrderLineResult) GetSkuId() *string {
	return s.SkuId
}

func (s *OrderLineResult) GetSkuTitle() *string {
	return s.SkuTitle
}

func (s *OrderLineResult) SetEticketInfos(v []*EticketInfo) *OrderLineResult {
	s.EticketInfos = v
	return s
}

func (s *OrderLineResult) SetLogisticsStatus(v string) *OrderLineResult {
	s.LogisticsStatus = &v
	return s
}

func (s *OrderLineResult) SetNumber(v string) *OrderLineResult {
	s.Number = &v
	return s
}

func (s *OrderLineResult) SetOrderId(v string) *OrderLineResult {
	s.OrderId = &v
	return s
}

func (s *OrderLineResult) SetOrderLineId(v string) *OrderLineResult {
	s.OrderLineId = &v
	return s
}

func (s *OrderLineResult) SetOrderLineStatus(v string) *OrderLineResult {
	s.OrderLineStatus = &v
	return s
}

func (s *OrderLineResult) SetPayFee(v int64) *OrderLineResult {
	s.PayFee = &v
	return s
}

func (s *OrderLineResult) SetProductId(v string) *OrderLineResult {
	s.ProductId = &v
	return s
}

func (s *OrderLineResult) SetProductPic(v string) *OrderLineResult {
	s.ProductPic = &v
	return s
}

func (s *OrderLineResult) SetProductTitle(v string) *OrderLineResult {
	s.ProductTitle = &v
	return s
}

func (s *OrderLineResult) SetSkuId(v string) *OrderLineResult {
	s.SkuId = &v
	return s
}

func (s *OrderLineResult) SetSkuTitle(v string) *OrderLineResult {
	s.SkuTitle = &v
	return s
}

func (s *OrderLineResult) Validate() error {
	if s.EticketInfos != nil {
		for _, item := range s.EticketInfos {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
