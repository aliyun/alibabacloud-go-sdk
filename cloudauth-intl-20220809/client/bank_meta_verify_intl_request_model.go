// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBankMetaVerifyIntlRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBankCard(v string) *BankMetaVerifyIntlRequest
	GetBankCard() *string
	SetIdentifyNum(v string) *BankMetaVerifyIntlRequest
	GetIdentifyNum() *string
	SetIdentityType(v string) *BankMetaVerifyIntlRequest
	GetIdentityType() *string
	SetMobile(v string) *BankMetaVerifyIntlRequest
	GetMobile() *string
	SetParamType(v string) *BankMetaVerifyIntlRequest
	GetParamType() *string
	SetProductCode(v string) *BankMetaVerifyIntlRequest
	GetProductCode() *string
	SetProductType(v string) *BankMetaVerifyIntlRequest
	GetProductType() *string
	SetUserName(v string) *BankMetaVerifyIntlRequest
	GetUserName() *string
	SetVerifyMode(v string) *BankMetaVerifyIntlRequest
	GetVerifyMode() *string
}

type BankMetaVerifyIntlRequest struct {
	// The bank card number.
	//
	// - If paramType is set to normal, enter the bank card number in plaintext.
	//
	// - If paramType is set to md5, provide the plaintext of all digits except the last 6 digits + the MD5 value (32-character lowercase) of the last 6 digits.
	//
	// This parameter is required.
	//
	// example:
	//
	// ● 明文：12345678******
	//
	// ● 密文：
	//
	// 12345678f8ee21920e37807b43e7e912ab829b6e
	BankCard *string `json:"BankCard,omitempty" xml:"BankCard,omitempty"`
	// The ID document number.
	//
	// - If paramType is set to normal, enter the ID document number in plaintext.
	//
	// - If paramType is set to md5:
	//
	//     - For ID cards: the first 6 digits (plaintext) + date of birth (ciphertext) + the last 4 digits (plaintext).
	//
	//     - For other documents: the last 2 characters are MD5-encrypted.
	//
	// Important:
	//
	// This parameter is required when ProductType is set to one of the following values:
	//
	// - BANK_CARD_3_META
	//
	// - BANK_CARD_4_META.
	//
	// example:
	//
	// ● 明文：429001********8211
	//
	// ● 密文：
	//
	// 42900132fa7bcd874161bea8ec8fd98f390ec98211
	IdentifyNum *string `json:"IdentifyNum,omitempty" xml:"IdentifyNum,omitempty"`
	// The ID document type. If left empty, the default value is ID card. For other document types, refer to the table below.
	//
	// example:
	//
	// 01
	IdentityType *string `json:"IdentityType,omitempty" xml:"IdentityType,omitempty"`
	// The phone number.
	//
	// - If paramType is set to normal, enter the phone number in plaintext.
	//
	// - If paramType is set to md5, enter the phone number in ciphertext.
	//
	// Important:
	//
	// - This parameter is required when ProductType is set to BANK_CARD_4_META.
	//
	// example:
	//
	// ● 明文：186****2055
	//
	// ● 密文：
	//
	// 849169cd3b20621c1c78bd61a11a4fc2
	Mobile *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	// The encryption method. Valid values:
	//
	// - normal: no encryption
	//
	// - md5: MD5 encryption
	//
	// Important:
	//
	// - All encrypted parameter values must be 32-character lowercase MD5 strings.
	//
	// - Different MD5 tools may produce different ciphertext. If the API call succeeds before encryption but fails after encryption, try a different MD5 tool.
	//
	// This parameter is required.
	//
	// example:
	//
	// normal
	ParamType *string `json:"ParamType,omitempty" xml:"ParamType,omitempty"`
	// Fixed value: BANK_CARD_N_META.
	//
	// This parameter is required.
	//
	// example:
	//
	// BANK_CARD_N_META
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The product type to call. Valid values:
	//
	// - BANK_CARD_2_META: Bank card number + name verification.
	//
	// - BANK_CARD_3_META: Bank card number + name + ID card number verification.
	//
	// - BANK_CARD_4_META: Bank card number + name + ID card number + phone number verification.
	//
	// This parameter is required.
	//
	// example:
	//
	// BANK_CARD_4_META
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	// The name.
	//
	// - If paramType is set to normal, enter the name in plaintext.
	//
	// - If paramType is set to md5, provide the MD5-encrypted first character of the name (32-character lowercase MD5) + the plaintext of the remaining characters.
	//
	// This parameter is required.
	//
	// example:
	//
	// 张*
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	// VERIFY_BANK_CARD: bank card authentication mode. Indicates whether the provided bank card number matches the user\\"s real name, ID card number, and phone number.
	//
	// This parameter is required.
	//
	// example:
	//
	// VERIFY_BANK_CARD
	VerifyMode *string `json:"VerifyMode,omitempty" xml:"VerifyMode,omitempty"`
}

func (s BankMetaVerifyIntlRequest) String() string {
	return dara.Prettify(s)
}

func (s BankMetaVerifyIntlRequest) GoString() string {
	return s.String()
}

func (s *BankMetaVerifyIntlRequest) GetBankCard() *string {
	return s.BankCard
}

func (s *BankMetaVerifyIntlRequest) GetIdentifyNum() *string {
	return s.IdentifyNum
}

func (s *BankMetaVerifyIntlRequest) GetIdentityType() *string {
	return s.IdentityType
}

func (s *BankMetaVerifyIntlRequest) GetMobile() *string {
	return s.Mobile
}

func (s *BankMetaVerifyIntlRequest) GetParamType() *string {
	return s.ParamType
}

func (s *BankMetaVerifyIntlRequest) GetProductCode() *string {
	return s.ProductCode
}

func (s *BankMetaVerifyIntlRequest) GetProductType() *string {
	return s.ProductType
}

func (s *BankMetaVerifyIntlRequest) GetUserName() *string {
	return s.UserName
}

func (s *BankMetaVerifyIntlRequest) GetVerifyMode() *string {
	return s.VerifyMode
}

func (s *BankMetaVerifyIntlRequest) SetBankCard(v string) *BankMetaVerifyIntlRequest {
	s.BankCard = &v
	return s
}

func (s *BankMetaVerifyIntlRequest) SetIdentifyNum(v string) *BankMetaVerifyIntlRequest {
	s.IdentifyNum = &v
	return s
}

func (s *BankMetaVerifyIntlRequest) SetIdentityType(v string) *BankMetaVerifyIntlRequest {
	s.IdentityType = &v
	return s
}

func (s *BankMetaVerifyIntlRequest) SetMobile(v string) *BankMetaVerifyIntlRequest {
	s.Mobile = &v
	return s
}

func (s *BankMetaVerifyIntlRequest) SetParamType(v string) *BankMetaVerifyIntlRequest {
	s.ParamType = &v
	return s
}

func (s *BankMetaVerifyIntlRequest) SetProductCode(v string) *BankMetaVerifyIntlRequest {
	s.ProductCode = &v
	return s
}

func (s *BankMetaVerifyIntlRequest) SetProductType(v string) *BankMetaVerifyIntlRequest {
	s.ProductType = &v
	return s
}

func (s *BankMetaVerifyIntlRequest) SetUserName(v string) *BankMetaVerifyIntlRequest {
	s.UserName = &v
	return s
}

func (s *BankMetaVerifyIntlRequest) SetVerifyMode(v string) *BankMetaVerifyIntlRequest {
	s.VerifyMode = &v
	return s
}

func (s *BankMetaVerifyIntlRequest) Validate() error {
	return dara.Validate(s)
}
