// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPurchaseOrderCreateCmd interface {
	dara.Model
	String() string
	GoString() string
	SetBuyerId(v string) *PurchaseOrderCreateCmd
	GetBuyerId() *string
	SetDeliveryAddress(v *AddressInfo) *PurchaseOrderCreateCmd
	GetDeliveryAddress() *AddressInfo
	SetExtInfo(v map[string]interface{}) *PurchaseOrderCreateCmd
	GetExtInfo() map[string]interface{}
	SetOuterPurchaseOrderId(v string) *PurchaseOrderCreateCmd
	GetOuterPurchaseOrderId() *string
	SetProductList(v []*ProductDTO) *PurchaseOrderCreateCmd
	GetProductList() []*ProductDTO
}

type PurchaseOrderCreateCmd struct {
	// This parameter is required.
	BuyerId *string `json:"buyerId,omitempty" xml:"buyerId,omitempty"`
	// This parameter is required.
	DeliveryAddress *AddressInfo           `json:"deliveryAddress,omitempty" xml:"deliveryAddress,omitempty"`
	ExtInfo         map[string]interface{} `json:"extInfo,omitempty" xml:"extInfo,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// outer123456
	OuterPurchaseOrderId *string `json:"outerPurchaseOrderId,omitempty" xml:"outerPurchaseOrderId,omitempty"`
	// This parameter is required.
	ProductList []*ProductDTO `json:"productList,omitempty" xml:"productList,omitempty" type:"Repeated"`
}

func (s PurchaseOrderCreateCmd) String() string {
	return dara.Prettify(s)
}

func (s PurchaseOrderCreateCmd) GoString() string {
	return s.String()
}

func (s *PurchaseOrderCreateCmd) GetBuyerId() *string {
	return s.BuyerId
}

func (s *PurchaseOrderCreateCmd) GetDeliveryAddress() *AddressInfo {
	return s.DeliveryAddress
}

func (s *PurchaseOrderCreateCmd) GetExtInfo() map[string]interface{} {
	return s.ExtInfo
}

func (s *PurchaseOrderCreateCmd) GetOuterPurchaseOrderId() *string {
	return s.OuterPurchaseOrderId
}

func (s *PurchaseOrderCreateCmd) GetProductList() []*ProductDTO {
	return s.ProductList
}

func (s *PurchaseOrderCreateCmd) SetBuyerId(v string) *PurchaseOrderCreateCmd {
	s.BuyerId = &v
	return s
}

func (s *PurchaseOrderCreateCmd) SetDeliveryAddress(v *AddressInfo) *PurchaseOrderCreateCmd {
	s.DeliveryAddress = v
	return s
}

func (s *PurchaseOrderCreateCmd) SetExtInfo(v map[string]interface{}) *PurchaseOrderCreateCmd {
	s.ExtInfo = v
	return s
}

func (s *PurchaseOrderCreateCmd) SetOuterPurchaseOrderId(v string) *PurchaseOrderCreateCmd {
	s.OuterPurchaseOrderId = &v
	return s
}

func (s *PurchaseOrderCreateCmd) SetProductList(v []*ProductDTO) *PurchaseOrderCreateCmd {
	s.ProductList = v
	return s
}

func (s *PurchaseOrderCreateCmd) Validate() error {
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
