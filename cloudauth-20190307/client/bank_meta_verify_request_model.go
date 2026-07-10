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
	// The bank card number.
	//
	// - If paramType is set to normal, enter the bank card number in plaintext.
	//
	// - If paramType is set to md5, enter the card number excluding the last 6 digits in plaintext + the MD5 hash (32-bit lowercase) of the last 6 digits.
	//
	// example:
	//
	// 610*************1181
	BankCard *string `json:"BankCard,omitempty" xml:"BankCard,omitempty"`
	// The ID card number.
	//
	// - This parameter is required if ProductType is set to BANK_CARD_3_META.
	//
	// - If paramType is set to normal, enter the ID card number in plaintext.
	//
	// - If paramType is set to md5, enter the first 6 digits of the ID card number in plaintext + the MD5 hash (32-bit lowercase) of the date of birth + the last 4 digits of the ID card number.
	//
	// example:
	//
	// 429001********8211
	IdentifyNum *string `json:"IdentifyNum,omitempty" xml:"IdentifyNum,omitempty"`
	// The identity document type.
	//
	// example:
	//
	// 01
	IdentityType *string `json:"IdentityType,omitempty" xml:"IdentityType,omitempty"`
	// The phone number.
	//
	// - This parameter is required if ProductType is set to BANK_CARD_4_META.
	//
	// - If paramType is set to normal, enter the phone number in plaintext.
	//
	// - If paramType is set to md5, enter the MD5 hash (32-bit lowercase) of the phone number.
	//
	// example:
	//
	// 138******11
	Mobile *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	// The parameter type. Valid values:
	//
	// - normal: not encrypted.
	//
	// - md5: MD5-encrypted.
	//
	// example:
	//
	// normal
	ParamType *string `json:"ParamType,omitempty" xml:"ParamType,omitempty"`
	// The product type. Valid values:
	//
	// - BANK_CARD_2_META: bank card number + name verification.
	//
	// - BANK_CARD_3_META: bank card number + name + ID card number verification.
	//
	// - BANK_CARD_4_META: bank card number + name + ID card number + phone number verification.
	//
	// example:
	//
	// BANK_CARD_2_META
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	// The name.
	//
	// - If paramType is set to normal, enter the name in plaintext.
	//
	// - If paramType is set to md5, enter the MD5 hash (32-bit lowercase) of the first character of the name + the remaining characters of the name in plaintext.
	//
	// example:
	//
	// 张*
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	// VERIFY_BANK_CARD: bank card verification mode. Specifies whether the provided bank card number matches the real name, ID card number, and phone number of the user.
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
