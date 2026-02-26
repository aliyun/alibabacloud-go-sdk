// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSelectionGroupAddProductRequest interface {
	dara.Model
	String() string
	GoString() string
	SetProductIds(v []*string) *SelectionGroupAddProductRequest
	GetProductIds() []*string
	SetPurchaserId(v string) *SelectionGroupAddProductRequest
	GetPurchaserId() *string
}

type SelectionGroupAddProductRequest struct {
	// This parameter is required.
	ProductIds []*string `json:"productIds,omitempty" xml:"productIds,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// PIDxxxxx
	PurchaserId *string `json:"purchaserId,omitempty" xml:"purchaserId,omitempty"`
}

func (s SelectionGroupAddProductRequest) String() string {
	return dara.Prettify(s)
}

func (s SelectionGroupAddProductRequest) GoString() string {
	return s.String()
}

func (s *SelectionGroupAddProductRequest) GetProductIds() []*string {
	return s.ProductIds
}

func (s *SelectionGroupAddProductRequest) GetPurchaserId() *string {
	return s.PurchaserId
}

func (s *SelectionGroupAddProductRequest) SetProductIds(v []*string) *SelectionGroupAddProductRequest {
	s.ProductIds = v
	return s
}

func (s *SelectionGroupAddProductRequest) SetPurchaserId(v string) *SelectionGroupAddProductRequest {
	s.PurchaserId = &v
	return s
}

func (s *SelectionGroupAddProductRequest) Validate() error {
	return dara.Validate(s)
}
