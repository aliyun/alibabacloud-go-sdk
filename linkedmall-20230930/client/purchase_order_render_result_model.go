// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPurchaseOrderRenderResult interface {
	dara.Model
	String() string
	GoString() string
	SetAddressList(v []*AddressInfo) *PurchaseOrderRenderResult
	GetAddressList() []*AddressInfo
	SetCanSell(v bool) *PurchaseOrderRenderResult
	GetCanSell() *bool
	SetExtInfo(v map[string]interface{}) *PurchaseOrderRenderResult
	GetExtInfo() map[string]interface{}
	SetMessage(v string) *PurchaseOrderRenderResult
	GetMessage() *string
	SetOrderList(v []*OrderRenderResult) *PurchaseOrderRenderResult
	GetOrderList() []*OrderRenderResult
	SetRequestId(v string) *PurchaseOrderRenderResult
	GetRequestId() *string
	SetUnsellableOrderList(v []*OrderRenderResult) *PurchaseOrderRenderResult
	GetUnsellableOrderList() []*OrderRenderResult
}

type PurchaseOrderRenderResult struct {
	AddressList []*AddressInfo `json:"addressList,omitempty" xml:"addressList,omitempty" type:"Repeated"`
	// example:
	//
	// true
	CanSell   *bool                  `json:"canSell,omitempty" xml:"canSell,omitempty"`
	ExtInfo   map[string]interface{} `json:"extInfo,omitempty" xml:"extInfo,omitempty"`
	Message   *string                `json:"message,omitempty" xml:"message,omitempty"`
	OrderList []*OrderRenderResult   `json:"orderList,omitempty" xml:"orderList,omitempty" type:"Repeated"`
	// example:
	//
	// 3239281273464326823
	RequestId           *string              `json:"requestId,omitempty" xml:"requestId,omitempty"`
	UnsellableOrderList []*OrderRenderResult `json:"unsellableOrderList,omitempty" xml:"unsellableOrderList,omitempty" type:"Repeated"`
}

func (s PurchaseOrderRenderResult) String() string {
	return dara.Prettify(s)
}

func (s PurchaseOrderRenderResult) GoString() string {
	return s.String()
}

func (s *PurchaseOrderRenderResult) GetAddressList() []*AddressInfo {
	return s.AddressList
}

func (s *PurchaseOrderRenderResult) GetCanSell() *bool {
	return s.CanSell
}

func (s *PurchaseOrderRenderResult) GetExtInfo() map[string]interface{} {
	return s.ExtInfo
}

func (s *PurchaseOrderRenderResult) GetMessage() *string {
	return s.Message
}

func (s *PurchaseOrderRenderResult) GetOrderList() []*OrderRenderResult {
	return s.OrderList
}

func (s *PurchaseOrderRenderResult) GetRequestId() *string {
	return s.RequestId
}

func (s *PurchaseOrderRenderResult) GetUnsellableOrderList() []*OrderRenderResult {
	return s.UnsellableOrderList
}

func (s *PurchaseOrderRenderResult) SetAddressList(v []*AddressInfo) *PurchaseOrderRenderResult {
	s.AddressList = v
	return s
}

func (s *PurchaseOrderRenderResult) SetCanSell(v bool) *PurchaseOrderRenderResult {
	s.CanSell = &v
	return s
}

func (s *PurchaseOrderRenderResult) SetExtInfo(v map[string]interface{}) *PurchaseOrderRenderResult {
	s.ExtInfo = v
	return s
}

func (s *PurchaseOrderRenderResult) SetMessage(v string) *PurchaseOrderRenderResult {
	s.Message = &v
	return s
}

func (s *PurchaseOrderRenderResult) SetOrderList(v []*OrderRenderResult) *PurchaseOrderRenderResult {
	s.OrderList = v
	return s
}

func (s *PurchaseOrderRenderResult) SetRequestId(v string) *PurchaseOrderRenderResult {
	s.RequestId = &v
	return s
}

func (s *PurchaseOrderRenderResult) SetUnsellableOrderList(v []*OrderRenderResult) *PurchaseOrderRenderResult {
	s.UnsellableOrderList = v
	return s
}

func (s *PurchaseOrderRenderResult) Validate() error {
	if s.AddressList != nil {
		for _, item := range s.AddressList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.OrderList != nil {
		for _, item := range s.OrderList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.UnsellableOrderList != nil {
		for _, item := range s.UnsellableOrderList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
