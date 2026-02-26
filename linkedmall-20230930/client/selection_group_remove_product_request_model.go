// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSelectionGroupRemoveProductRequest interface {
	dara.Model
	String() string
	GoString() string
	SetProductIds(v []*string) *SelectionGroupRemoveProductRequest
	GetProductIds() []*string
	SetPurchaserId(v string) *SelectionGroupRemoveProductRequest
	GetPurchaserId() *string
}

type SelectionGroupRemoveProductRequest struct {
	// This parameter is required.
	ProductIds []*string `json:"productIds,omitempty" xml:"productIds,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// PIDxxxxx
	PurchaserId *string `json:"purchaserId,omitempty" xml:"purchaserId,omitempty"`
}

func (s SelectionGroupRemoveProductRequest) String() string {
	return dara.Prettify(s)
}

func (s SelectionGroupRemoveProductRequest) GoString() string {
	return s.String()
}

func (s *SelectionGroupRemoveProductRequest) GetProductIds() []*string {
	return s.ProductIds
}

func (s *SelectionGroupRemoveProductRequest) GetPurchaserId() *string {
	return s.PurchaserId
}

func (s *SelectionGroupRemoveProductRequest) SetProductIds(v []*string) *SelectionGroupRemoveProductRequest {
	s.ProductIds = v
	return s
}

func (s *SelectionGroupRemoveProductRequest) SetPurchaserId(v string) *SelectionGroupRemoveProductRequest {
	s.PurchaserId = &v
	return s
}

func (s *SelectionGroupRemoveProductRequest) Validate() error {
	return dara.Validate(s)
}
