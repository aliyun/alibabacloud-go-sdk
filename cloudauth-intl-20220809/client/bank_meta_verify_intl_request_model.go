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
	// This parameter is required.
	BankCard    *string `json:"BankCard,omitempty" xml:"BankCard,omitempty"`
	IdentifyNum *string `json:"IdentifyNum,omitempty" xml:"IdentifyNum,omitempty"`
	// example:
	//
	// 01
	IdentityType *string `json:"IdentityType,omitempty" xml:"IdentityType,omitempty"`
	Mobile       *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// normal
	ParamType *string `json:"ParamType,omitempty" xml:"ParamType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// BANK_CARD_N_META
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// BANK_CARD_4_META
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	// This parameter is required.
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
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
