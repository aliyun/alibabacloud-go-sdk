// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIntlFlightOtaItemDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBtripUserId(v string) *IntlFlightOtaItemDetailRequest
	GetBtripUserId() *string
	SetBuyerName(v string) *IntlFlightOtaItemDetailRequest
	GetBuyerName() *string
	SetIsvName(v string) *IntlFlightOtaItemDetailRequest
	GetIsvName() *string
	SetLanguage(v string) *IntlFlightOtaItemDetailRequest
	GetLanguage() *string
	SetSupplierCode(v string) *IntlFlightOtaItemDetailRequest
	GetSupplierCode() *string
}

type IntlFlightOtaItemDetailRequest struct {
	// example:
	//
	// 10001
	BtripUserId *string `json:"btrip_user_id,omitempty" xml:"btrip_user_id,omitempty"`
	BuyerName   *string `json:"buyer_name,omitempty" xml:"buyer_name,omitempty"`
	// example:
	//
	// open12igetbis4o07v10B1TlOWcM00
	IsvName  *string `json:"isv_name,omitempty" xml:"isv_name,omitempty"`
	Language *string `json:"language,omitempty" xml:"language,omitempty"`
	// example:
	//
	// JIANHANG
	SupplierCode *string `json:"supplier_code,omitempty" xml:"supplier_code,omitempty"`
}

func (s IntlFlightOtaItemDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightOtaItemDetailRequest) GoString() string {
	return s.String()
}

func (s *IntlFlightOtaItemDetailRequest) GetBtripUserId() *string {
	return s.BtripUserId
}

func (s *IntlFlightOtaItemDetailRequest) GetBuyerName() *string {
	return s.BuyerName
}

func (s *IntlFlightOtaItemDetailRequest) GetIsvName() *string {
	return s.IsvName
}

func (s *IntlFlightOtaItemDetailRequest) GetLanguage() *string {
	return s.Language
}

func (s *IntlFlightOtaItemDetailRequest) GetSupplierCode() *string {
	return s.SupplierCode
}

func (s *IntlFlightOtaItemDetailRequest) SetBtripUserId(v string) *IntlFlightOtaItemDetailRequest {
	s.BtripUserId = &v
	return s
}

func (s *IntlFlightOtaItemDetailRequest) SetBuyerName(v string) *IntlFlightOtaItemDetailRequest {
	s.BuyerName = &v
	return s
}

func (s *IntlFlightOtaItemDetailRequest) SetIsvName(v string) *IntlFlightOtaItemDetailRequest {
	s.IsvName = &v
	return s
}

func (s *IntlFlightOtaItemDetailRequest) SetLanguage(v string) *IntlFlightOtaItemDetailRequest {
	s.Language = &v
	return s
}

func (s *IntlFlightOtaItemDetailRequest) SetSupplierCode(v string) *IntlFlightOtaItemDetailRequest {
	s.SupplierCode = &v
	return s
}

func (s *IntlFlightOtaItemDetailRequest) Validate() error {
	return dara.Validate(s)
}
