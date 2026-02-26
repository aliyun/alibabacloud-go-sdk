// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPurchaseOrderRenderQuery interface {
	dara.Model
	String() string
	GoString() string
	SetBuyerId(v string) *PurchaseOrderRenderQuery
	GetBuyerId() *string
	SetDeliveryAddress(v *AddressInfo) *PurchaseOrderRenderQuery
	GetDeliveryAddress() *AddressInfo
	SetExtInfo(v map[string]interface{}) *PurchaseOrderRenderQuery
	GetExtInfo() map[string]interface{}
	SetProductList(v []*OrderRenderProductDTO) *PurchaseOrderRenderQuery
	GetProductList() []*OrderRenderProductDTO
}

type PurchaseOrderRenderQuery struct {
	// This parameter is required.
	//
	// example:
	//
	// test1234567
	BuyerId *string `json:"buyerId,omitempty" xml:"buyerId,omitempty"`
	// This parameter is required.
	DeliveryAddress *AddressInfo `json:"deliveryAddress,omitempty" xml:"deliveryAddress,omitempty"`
	// example:
	//
	// {}
	ExtInfo map[string]interface{} `json:"extInfo,omitempty" xml:"extInfo,omitempty"`
	// This parameter is required.
	ProductList []*OrderRenderProductDTO `json:"productList,omitempty" xml:"productList,omitempty" type:"Repeated"`
}

func (s PurchaseOrderRenderQuery) String() string {
	return dara.Prettify(s)
}

func (s PurchaseOrderRenderQuery) GoString() string {
	return s.String()
}

func (s *PurchaseOrderRenderQuery) GetBuyerId() *string {
	return s.BuyerId
}

func (s *PurchaseOrderRenderQuery) GetDeliveryAddress() *AddressInfo {
	return s.DeliveryAddress
}

func (s *PurchaseOrderRenderQuery) GetExtInfo() map[string]interface{} {
	return s.ExtInfo
}

func (s *PurchaseOrderRenderQuery) GetProductList() []*OrderRenderProductDTO {
	return s.ProductList
}

func (s *PurchaseOrderRenderQuery) SetBuyerId(v string) *PurchaseOrderRenderQuery {
	s.BuyerId = &v
	return s
}

func (s *PurchaseOrderRenderQuery) SetDeliveryAddress(v *AddressInfo) *PurchaseOrderRenderQuery {
	s.DeliveryAddress = v
	return s
}

func (s *PurchaseOrderRenderQuery) SetExtInfo(v map[string]interface{}) *PurchaseOrderRenderQuery {
	s.ExtInfo = v
	return s
}

func (s *PurchaseOrderRenderQuery) SetProductList(v []*OrderRenderProductDTO) *PurchaseOrderRenderQuery {
	s.ProductList = v
	return s
}

func (s *PurchaseOrderRenderQuery) Validate() error {
	if s.DeliveryAddress != nil {
		if err := s.DeliveryAddress.Validate(); err != nil {
			return err
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
