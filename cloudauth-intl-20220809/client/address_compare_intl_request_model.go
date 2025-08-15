// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddressCompareIntlRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDefaultCountry(v string) *AddressCompareIntlRequest
	GetDefaultCountry() *string
	SetProductCode(v string) *AddressCompareIntlRequest
	GetProductCode() *string
	SetText1(v string) *AddressCompareIntlRequest
	GetText1() *string
	SetText2(v string) *AddressCompareIntlRequest
	GetText2() *string
}

type AddressCompareIntlRequest struct {
	// This parameter is required.
	DefaultCountry *string `json:"DefaultCountry,omitempty" xml:"DefaultCountry,omitempty"`
	// ADD_VERIFY
	//
	// This parameter is required.
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// This parameter is required.
	Text1 *string `json:"Text1,omitempty" xml:"Text1,omitempty"`
	// This parameter is required.
	Text2 *string `json:"Text2,omitempty" xml:"Text2,omitempty"`
}

func (s AddressCompareIntlRequest) String() string {
	return dara.Prettify(s)
}

func (s AddressCompareIntlRequest) GoString() string {
	return s.String()
}

func (s *AddressCompareIntlRequest) GetDefaultCountry() *string {
	return s.DefaultCountry
}

func (s *AddressCompareIntlRequest) GetProductCode() *string {
	return s.ProductCode
}

func (s *AddressCompareIntlRequest) GetText1() *string {
	return s.Text1
}

func (s *AddressCompareIntlRequest) GetText2() *string {
	return s.Text2
}

func (s *AddressCompareIntlRequest) SetDefaultCountry(v string) *AddressCompareIntlRequest {
	s.DefaultCountry = &v
	return s
}

func (s *AddressCompareIntlRequest) SetProductCode(v string) *AddressCompareIntlRequest {
	s.ProductCode = &v
	return s
}

func (s *AddressCompareIntlRequest) SetText1(v string) *AddressCompareIntlRequest {
	s.Text1 = &v
	return s
}

func (s *AddressCompareIntlRequest) SetText2(v string) *AddressCompareIntlRequest {
	s.Text2 = &v
	return s
}

func (s *AddressCompareIntlRequest) Validate() error {
	return dara.Validate(s)
}
