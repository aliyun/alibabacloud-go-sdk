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
	// - If paramType is set to normal, enter the plaintext bank card number.
	//
	// - If paramType is set to md5, enter the plaintext of the bank card number except the last 6 digits, followed by the 32-character lowercase MD5 hash of the last 6 digits.
	//
	// This parameter is required.
	//
	// example:
	//
	// ● Plaintext: 12345678******
	//
	// ● Ciphertext:
	//
	// 12345678f8ee21920e37807b43e7e912ab829b6e
	BankCard *string `json:"BankCard,omitempty" xml:"BankCard,omitempty"`
	// The identification document number.
	//
	// - If paramType is set to normal, enter the plaintext identification document number.
	//
	// - If paramType is set to md5:
	//
	//     - For an ID card, enter the first 6 digits in plaintext, followed by the ciphertext of the date of birth, and then the last 4 digits in plaintext.
	//
	//     - For other identification documents, encrypt the last two characters by using MD5.
	//
	// Important:
	//
	// This parameter is required if ProductType is set to one of the following values:
	//
	// - BANK_CARD_3_META
	//
	// - BANK_CARD_4_META
	//
	// example:
	//
	// ● Plaintext: 429001********8211
	//
	// ● Ciphertext:
	//
	// 42900132fa7bcd874161bea8ec8fd98f390ec98211
	IdentifyNum *string `json:"IdentifyNum,omitempty" xml:"IdentifyNum,omitempty"`
	// The type of the identification document. If you leave this parameter empty, the default value is ID card. For other types of identification documents, see the following table.
	//
	// example:
	//
	// 01
	IdentityType *string `json:"IdentityType,omitempty" xml:"IdentityType,omitempty"`
	// The mobile phone number.
	//
	// - If paramType is set to normal, enter the plaintext mobile phone number.
	//
	// - If paramType is set to md5, enter the ciphertext mobile phone number.
	//
	// Important:
	//
	// - This parameter is required if ProductType is set to BANK_CARD_4_META.
	//
	// example:
	//
	// ● Plaintext: 186****2055
	//
	// ● Ciphertext:
	//
	// 849169cd3b20621c1c78bd61a11a4fc2
	Mobile *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	// The encryption method. Valid values:
	//
	// - normal: No encryption.
	//
	// - md5: MD5 encryption.
	//
	// Important:
	//
	// - The ciphertext of all parameters must be a 32-character lowercase MD5 string.
	//
	// - Different MD5 tools may generate different ciphertexts. If the API call succeeds before encryption but fails after encryption, try using a different MD5 tool.
	//
	// This parameter is required.
	//
	// example:
	//
	// normal
	ParamType *string `json:"ParamType,omitempty" xml:"ParamType,omitempty"`
	// The fixed value. Valid value: BANK_CARD_N_META.
	//
	// This parameter is required.
	//
	// example:
	//
	// BANK_CARD_N_META
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The product type. Valid values:
	//
	// - BANK_CARD_2_META: Verifies the bank card number and name.
	//
	// - BANK_CARD_3_META: Verifies the bank card number, name, and ID card number.
	//
	// - BANK_CARD_4_META: Verifies the bank card number, name, ID card number, and mobile phone number.
	//
	// This parameter is required.
	//
	// example:
	//
	// BANK_CARD_4_META
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	// The name.
	//
	// - If paramType is set to normal, enter the plaintext name.
	//
	// - If paramType is set to md5, enter the 32-character lowercase MD5 hash of the first character of the name, followed by the plaintext of the remaining characters.
	//
	// This parameter is required.
	//
	// example:
	//
	// Zhang*
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	// The verification mode. Valid value: VERIFY_BANK_CARD. This value indicates the bank card authentication mode, which verifies whether the provided bank card number matches the real name, ID card number, and mobile phone number of the user.
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
