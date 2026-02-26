// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSelectionProductRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDivisionCode(v string) *GetSelectionProductRequest
	GetDivisionCode() *string
	SetPurchaserId(v string) *GetSelectionProductRequest
	GetPurchaserId() *string
}

type GetSelectionProductRequest struct {
	// example:
	//
	// 110000
	DivisionCode *string `json:"divisionCode,omitempty" xml:"divisionCode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 56****2304
	PurchaserId *string `json:"purchaserId,omitempty" xml:"purchaserId,omitempty"`
}

func (s GetSelectionProductRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSelectionProductRequest) GoString() string {
	return s.String()
}

func (s *GetSelectionProductRequest) GetDivisionCode() *string {
	return s.DivisionCode
}

func (s *GetSelectionProductRequest) GetPurchaserId() *string {
	return s.PurchaserId
}

func (s *GetSelectionProductRequest) SetDivisionCode(v string) *GetSelectionProductRequest {
	s.DivisionCode = &v
	return s
}

func (s *GetSelectionProductRequest) SetPurchaserId(v string) *GetSelectionProductRequest {
	s.PurchaserId = &v
	return s
}

func (s *GetSelectionProductRequest) Validate() error {
	return dara.Validate(s)
}
