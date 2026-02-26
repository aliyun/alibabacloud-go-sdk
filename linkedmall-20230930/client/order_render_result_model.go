// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOrderRenderResult interface {
	dara.Model
	String() string
	GoString() string
	SetCanSell(v bool) *OrderRenderResult
	GetCanSell() *bool
	SetDeliveryInfoList(v []*DeliveryInfo) *OrderRenderResult
	GetDeliveryInfoList() []*DeliveryInfo
	SetExtInfo(v map[string]interface{}) *OrderRenderResult
	GetExtInfo() map[string]interface{}
	SetMessage(v string) *OrderRenderResult
	GetMessage() *string
	SetProductList(v []*OrderProductResult) *OrderRenderResult
	GetProductList() []*OrderProductResult
}

type OrderRenderResult struct {
	// example:
	//
	// true
	CanSell          *bool                  `json:"canSell,omitempty" xml:"canSell,omitempty"`
	DeliveryInfoList []*DeliveryInfo        `json:"deliveryInfoList,omitempty" xml:"deliveryInfoList,omitempty" type:"Repeated"`
	ExtInfo          map[string]interface{} `json:"extInfo,omitempty" xml:"extInfo,omitempty"`
	Message          *string                `json:"message,omitempty" xml:"message,omitempty"`
	ProductList      []*OrderProductResult  `json:"productList,omitempty" xml:"productList,omitempty" type:"Repeated"`
}

func (s OrderRenderResult) String() string {
	return dara.Prettify(s)
}

func (s OrderRenderResult) GoString() string {
	return s.String()
}

func (s *OrderRenderResult) GetCanSell() *bool {
	return s.CanSell
}

func (s *OrderRenderResult) GetDeliveryInfoList() []*DeliveryInfo {
	return s.DeliveryInfoList
}

func (s *OrderRenderResult) GetExtInfo() map[string]interface{} {
	return s.ExtInfo
}

func (s *OrderRenderResult) GetMessage() *string {
	return s.Message
}

func (s *OrderRenderResult) GetProductList() []*OrderProductResult {
	return s.ProductList
}

func (s *OrderRenderResult) SetCanSell(v bool) *OrderRenderResult {
	s.CanSell = &v
	return s
}

func (s *OrderRenderResult) SetDeliveryInfoList(v []*DeliveryInfo) *OrderRenderResult {
	s.DeliveryInfoList = v
	return s
}

func (s *OrderRenderResult) SetExtInfo(v map[string]interface{}) *OrderRenderResult {
	s.ExtInfo = v
	return s
}

func (s *OrderRenderResult) SetMessage(v string) *OrderRenderResult {
	s.Message = &v
	return s
}

func (s *OrderRenderResult) SetProductList(v []*OrderProductResult) *OrderRenderResult {
	s.ProductList = v
	return s
}

func (s *OrderRenderResult) Validate() error {
	if s.DeliveryInfoList != nil {
		for _, item := range s.DeliveryInfoList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ProductList != nil {
		for _, item := range s.ProductList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
