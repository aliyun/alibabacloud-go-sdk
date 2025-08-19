// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBankMetaVerifyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBankCard(v string) *BankMetaVerifyRequest
	GetBankCard() *string
	SetIdentifyNum(v string) *BankMetaVerifyRequest
	GetIdentifyNum() *string
	SetIdentityType(v string) *BankMetaVerifyRequest
	GetIdentityType() *string
	SetMobile(v string) *BankMetaVerifyRequest
	GetMobile() *string
	SetParamType(v string) *BankMetaVerifyRequest
	GetParamType() *string
	SetProductType(v string) *BankMetaVerifyRequest
	GetProductType() *string
	SetUserName(v string) *BankMetaVerifyRequest
	GetUserName() *string
	SetVerifyMode(v string) *BankMetaVerifyRequest
	GetVerifyMode() *string
}

type BankMetaVerifyRequest struct {
	// Bank card number.
	//
	// - When `paramType` is `normal`, enter the plain text bank card number.
	//
	// - When `paramType` is `md5`, enter the part of the card number except the last 6 digits in plain text + the last 6 digits encrypted with MD5 (32 lowercase).
	//
	// example:
	//
	// 610*************1181
	BankCard *string `json:"BankCard,omitempty" xml:"BankCard,omitempty"`
	// ID number.
	//
	// - When `ProductType` is `BANK_CARD_3_META`, this field is required.
	//
	// - When `paramType` is `normal`, enter the plain text ID number.
	//
	// - When `paramType` is `md5`, enter the first 6 digits of the ID number in plain text + the birth date encrypted with MD5 (32 lowercase MD5) + the last 4 digits of the ID number.
	//
	// example:
	//
	// 429001********8211
	IdentifyNum *string `json:"IdentifyNum,omitempty" xml:"IdentifyNum,omitempty"`
	// Identity type.
	//
	// example:
	//
	// 01
	IdentityType *string `json:"IdentityType,omitempty" xml:"IdentityType,omitempty"`
	// Mobile phone number.
	//
	// - When `ProductType` is `BANK_CARD_4_META`, this field is required.
	//
	// - When `paramType` is `normal`, enter the plain text mobile phone number.
	//
	// - When `paramType` is `md5`, enter the mobile phone number (32 lowercase MD5).
	//
	// example:
	//
	// 138******11
	Mobile *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	// Parameter type:
	//
	// - normal: Unencrypted.
	//
	// - md5: MD5 encrypted.
	//
	// example:
	//
	// normal
	ParamType *string `json:"ParamType,omitempty" xml:"ParamType,omitempty"`
	// Product type to call:
	//
	// - BANK_CARD_2_META: Bank card number + name verification.
	//
	// - BANK_CARD_3_META: Bank card number + name + ID number verification.
	//
	// - BANK_CARD_4_META: Bank card number + name + ID number + mobile phone number verification.
	//
	// example:
	//
	// BANK_CARD_2_META
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	// Name.
	//
	// - When `paramType` is `normal`, enter the plain text name.
	//
	// - When `paramType` is `md5`, encrypt the first character of the name with MD5 (32 lowercase MD5) + the rest of the name in plain text.
	//
	// example:
	//
	// 张*
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	// VERIFY_BANK_CARD: Bank card authentication mode. It indicates whether the provided bank card number matches the user\\"s real name, ID number, and mobile phone number.
	//
	// example:
	//
	// VERIFY_BANK_CARD
	VerifyMode *string `json:"VerifyMode,omitempty" xml:"VerifyMode,omitempty"`
}

func (s BankMetaVerifyRequest) String() string {
	return dara.Prettify(s)
}

func (s BankMetaVerifyRequest) GoString() string {
	return s.String()
}

func (s *BankMetaVerifyRequest) GetBankCard() *string {
	return s.BankCard
}

func (s *BankMetaVerifyRequest) GetIdentifyNum() *string {
	return s.IdentifyNum
}

func (s *BankMetaVerifyRequest) GetIdentityType() *string {
	return s.IdentityType
}

func (s *BankMetaVerifyRequest) GetMobile() *string {
	return s.Mobile
}

func (s *BankMetaVerifyRequest) GetParamType() *string {
	return s.ParamType
}

func (s *BankMetaVerifyRequest) GetProductType() *string {
	return s.ProductType
}

func (s *BankMetaVerifyRequest) GetUserName() *string {
	return s.UserName
}

func (s *BankMetaVerifyRequest) GetVerifyMode() *string {
	return s.VerifyMode
}

func (s *BankMetaVerifyRequest) SetBankCard(v string) *BankMetaVerifyRequest {
	s.BankCard = &v
	return s
}

func (s *BankMetaVerifyRequest) SetIdentifyNum(v string) *BankMetaVerifyRequest {
	s.IdentifyNum = &v
	return s
}

func (s *BankMetaVerifyRequest) SetIdentityType(v string) *BankMetaVerifyRequest {
	s.IdentityType = &v
	return s
}

func (s *BankMetaVerifyRequest) SetMobile(v string) *BankMetaVerifyRequest {
	s.Mobile = &v
	return s
}

func (s *BankMetaVerifyRequest) SetParamType(v string) *BankMetaVerifyRequest {
	s.ParamType = &v
	return s
}

func (s *BankMetaVerifyRequest) SetProductType(v string) *BankMetaVerifyRequest {
	s.ProductType = &v
	return s
}

func (s *BankMetaVerifyRequest) SetUserName(v string) *BankMetaVerifyRequest {
	s.UserName = &v
	return s
}

func (s *BankMetaVerifyRequest) SetVerifyMode(v string) *BankMetaVerifyRequest {
	s.VerifyMode = &v
	return s
}

func (s *BankMetaVerifyRequest) Validate() error {
	return dara.Validate(s)
}
