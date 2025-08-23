// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddressVerifyV2IntlRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceToken(v string) *AddressVerifyV2IntlRequest
	GetDeviceToken() *string
	SetMobile(v string) *AddressVerifyV2IntlRequest
	GetMobile() *string
	SetProductCode(v string) *AddressVerifyV2IntlRequest
	GetProductCode() *string
	SetRegCountry(v string) *AddressVerifyV2IntlRequest
	GetRegCountry() *string
	SetText(v string) *AddressVerifyV2IntlRequest
	GetText() *string
	SetVerifyType(v string) *AddressVerifyV2IntlRequest
	GetVerifyType() *string
}

type AddressVerifyV2IntlRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// Tk9SSUQuMS*****************ZDNmNWY5NzQxOW1o
	DeviceToken *string `json:"DeviceToken,omitempty" xml:"DeviceToken,omitempty"`
	// example:
	//
	// 1872334****
	Mobile *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ADD_VERIFY_PRO
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// This parameter is required.
	RegCountry *string `json:"RegCountry,omitempty" xml:"RegCountry,omitempty"`
	Text       *string `json:"Text,omitempty" xml:"Text,omitempty"`
	// example:
	//
	// HOME
	VerifyType *string `json:"VerifyType,omitempty" xml:"VerifyType,omitempty"`
}

func (s AddressVerifyV2IntlRequest) String() string {
	return dara.Prettify(s)
}

func (s AddressVerifyV2IntlRequest) GoString() string {
	return s.String()
}

func (s *AddressVerifyV2IntlRequest) GetDeviceToken() *string {
	return s.DeviceToken
}

func (s *AddressVerifyV2IntlRequest) GetMobile() *string {
	return s.Mobile
}

func (s *AddressVerifyV2IntlRequest) GetProductCode() *string {
	return s.ProductCode
}

func (s *AddressVerifyV2IntlRequest) GetRegCountry() *string {
	return s.RegCountry
}

func (s *AddressVerifyV2IntlRequest) GetText() *string {
	return s.Text
}

func (s *AddressVerifyV2IntlRequest) GetVerifyType() *string {
	return s.VerifyType
}

func (s *AddressVerifyV2IntlRequest) SetDeviceToken(v string) *AddressVerifyV2IntlRequest {
	s.DeviceToken = &v
	return s
}

func (s *AddressVerifyV2IntlRequest) SetMobile(v string) *AddressVerifyV2IntlRequest {
	s.Mobile = &v
	return s
}

func (s *AddressVerifyV2IntlRequest) SetProductCode(v string) *AddressVerifyV2IntlRequest {
	s.ProductCode = &v
	return s
}

func (s *AddressVerifyV2IntlRequest) SetRegCountry(v string) *AddressVerifyV2IntlRequest {
	s.RegCountry = &v
	return s
}

func (s *AddressVerifyV2IntlRequest) SetText(v string) *AddressVerifyV2IntlRequest {
	s.Text = &v
	return s
}

func (s *AddressVerifyV2IntlRequest) SetVerifyType(v string) *AddressVerifyV2IntlRequest {
	s.VerifyType = &v
	return s
}

func (s *AddressVerifyV2IntlRequest) Validate() error {
	return dara.Validate(s)
}
