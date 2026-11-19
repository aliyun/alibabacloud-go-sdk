// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSelectionProductSaleInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDivisionCode(v string) *GetSelectionProductSaleInfoRequest
	GetDivisionCode() *string
	SetPurchaserId(v string) *GetSelectionProductSaleInfoRequest
	GetPurchaserId() *string
}

type GetSelectionProductSaleInfoRequest struct {
	// The division code. A level-5 address code for a town or sub-district is recommended.
	//
	// example:
	//
	// 330106109
	DivisionCode *string `json:"divisionCode,omitempty" xml:"divisionCode,omitempty"`
	// The purchaser ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// PID22000009
	PurchaserId *string `json:"purchaserId,omitempty" xml:"purchaserId,omitempty"`
}

func (s GetSelectionProductSaleInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSelectionProductSaleInfoRequest) GoString() string {
	return s.String()
}

func (s *GetSelectionProductSaleInfoRequest) GetDivisionCode() *string {
	return s.DivisionCode
}

func (s *GetSelectionProductSaleInfoRequest) GetPurchaserId() *string {
	return s.PurchaserId
}

func (s *GetSelectionProductSaleInfoRequest) SetDivisionCode(v string) *GetSelectionProductSaleInfoRequest {
	s.DivisionCode = &v
	return s
}

func (s *GetSelectionProductSaleInfoRequest) SetPurchaserId(v string) *GetSelectionProductSaleInfoRequest {
	s.PurchaserId = &v
	return s
}

func (s *GetSelectionProductSaleInfoRequest) Validate() error {
	return dara.Validate(s)
}
