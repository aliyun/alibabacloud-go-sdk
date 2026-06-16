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
	// The device token, which is used for risk identification.
	//
	// This parameter is required.
	//
	// example:
	//
	// Tk9SSUQuMS*****************ZDNmNWY5NzQxOW1o
	DeviceToken *string `json:"DeviceToken,omitempty" xml:"DeviceToken,omitempty"`
	// The China mobile phone number.
	//
	// example:
	//
	// 1872334****
	Mobile *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	// The product code. Set this parameter to ADD_VERIFY_PRO.
	//
	// This parameter is required.
	//
	// example:
	//
	// ADD_VERIFY_PRO
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The list of prohibited countries.
	//
	// This parameter is required.
	//
	// example:
	//
	// 目前仅支持：USA
	RegCountry *string `json:"RegCountry,omitempty" xml:"RegCountry,omitempty"`
	// The detailed address text.
	//
	// example:
	//
	// 江苏省常州市*******小区
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
	// The address verification method. Valid values:
	//
	// - HOME: home address verification.
	//
	// - WORK: work address verification.
	//
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
