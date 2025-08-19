// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCredentialVerifyShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCertNum(v string) *CredentialVerifyShrinkRequest
	GetCertNum() *string
	SetCredName(v string) *CredentialVerifyShrinkRequest
	GetCredName() *string
	SetCredType(v string) *CredentialVerifyShrinkRequest
	GetCredType() *string
	SetIdentifyNum(v string) *CredentialVerifyShrinkRequest
	GetIdentifyNum() *string
	SetImageContext(v string) *CredentialVerifyShrinkRequest
	GetImageContext() *string
	SetImageUrl(v string) *CredentialVerifyShrinkRequest
	GetImageUrl() *string
	SetIsCheck(v string) *CredentialVerifyShrinkRequest
	GetIsCheck() *string
	SetIsOCR(v string) *CredentialVerifyShrinkRequest
	GetIsOCR() *string
	SetMerchantDetailShrink(v string) *CredentialVerifyShrinkRequest
	GetMerchantDetailShrink() *string
	SetMerchantId(v string) *CredentialVerifyShrinkRequest
	GetMerchantId() *string
	SetProductCode(v string) *CredentialVerifyShrinkRequest
	GetProductCode() *string
	SetPrompt(v string) *CredentialVerifyShrinkRequest
	GetPrompt() *string
	SetPromptModel(v string) *CredentialVerifyShrinkRequest
	GetPromptModel() *string
	SetUserName(v string) *CredentialVerifyShrinkRequest
	GetUserName() *string
}

type CredentialVerifyShrinkRequest struct {
	// example:
	//
	// 4601*****
	CertNum *string `json:"CertNum,omitempty" xml:"CertNum,omitempty"`
	// example:
	//
	// 0104
	CredName *string `json:"CredName,omitempty" xml:"CredName,omitempty"`
	// example:
	//
	// 01
	CredType *string `json:"CredType,omitempty" xml:"CredType,omitempty"`
	// example:
	//
	// 429001********8211
	IdentifyNum *string `json:"IdentifyNum,omitempty" xml:"IdentifyNum,omitempty"`
	// example:
	//
	// base64
	ImageContext *string `json:"ImageContext,omitempty" xml:"ImageContext,omitempty"`
	// example:
	//
	// http://marry.momocdn.com/avatar/3B/B6/3BB6527E-7467-926E-1048-B43614F20CC420230803_L.jpg
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	// example:
	//
	// 0
	IsCheck *string `json:"IsCheck,omitempty" xml:"IsCheck,omitempty"`
	// example:
	//
	// 1
	IsOCR                *string `json:"IsOCR,omitempty" xml:"IsOCR,omitempty"`
	MerchantDetailShrink *string `json:"MerchantDetail,omitempty" xml:"MerchantDetail,omitempty"`
	MerchantId           *string `json:"MerchantId,omitempty" xml:"MerchantId,omitempty"`
	ProductCode          *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	Prompt               *string `json:"Prompt,omitempty" xml:"Prompt,omitempty"`
	PromptModel          *string `json:"PromptModel,omitempty" xml:"PromptModel,omitempty"`
	UserName             *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s CredentialVerifyShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CredentialVerifyShrinkRequest) GoString() string {
	return s.String()
}

func (s *CredentialVerifyShrinkRequest) GetCertNum() *string {
	return s.CertNum
}

func (s *CredentialVerifyShrinkRequest) GetCredName() *string {
	return s.CredName
}

func (s *CredentialVerifyShrinkRequest) GetCredType() *string {
	return s.CredType
}

func (s *CredentialVerifyShrinkRequest) GetIdentifyNum() *string {
	return s.IdentifyNum
}

func (s *CredentialVerifyShrinkRequest) GetImageContext() *string {
	return s.ImageContext
}

func (s *CredentialVerifyShrinkRequest) GetImageUrl() *string {
	return s.ImageUrl
}

func (s *CredentialVerifyShrinkRequest) GetIsCheck() *string {
	return s.IsCheck
}

func (s *CredentialVerifyShrinkRequest) GetIsOCR() *string {
	return s.IsOCR
}

func (s *CredentialVerifyShrinkRequest) GetMerchantDetailShrink() *string {
	return s.MerchantDetailShrink
}

func (s *CredentialVerifyShrinkRequest) GetMerchantId() *string {
	return s.MerchantId
}

func (s *CredentialVerifyShrinkRequest) GetProductCode() *string {
	return s.ProductCode
}

func (s *CredentialVerifyShrinkRequest) GetPrompt() *string {
	return s.Prompt
}

func (s *CredentialVerifyShrinkRequest) GetPromptModel() *string {
	return s.PromptModel
}

func (s *CredentialVerifyShrinkRequest) GetUserName() *string {
	return s.UserName
}

func (s *CredentialVerifyShrinkRequest) SetCertNum(v string) *CredentialVerifyShrinkRequest {
	s.CertNum = &v
	return s
}

func (s *CredentialVerifyShrinkRequest) SetCredName(v string) *CredentialVerifyShrinkRequest {
	s.CredName = &v
	return s
}

func (s *CredentialVerifyShrinkRequest) SetCredType(v string) *CredentialVerifyShrinkRequest {
	s.CredType = &v
	return s
}

func (s *CredentialVerifyShrinkRequest) SetIdentifyNum(v string) *CredentialVerifyShrinkRequest {
	s.IdentifyNum = &v
	return s
}

func (s *CredentialVerifyShrinkRequest) SetImageContext(v string) *CredentialVerifyShrinkRequest {
	s.ImageContext = &v
	return s
}

func (s *CredentialVerifyShrinkRequest) SetImageUrl(v string) *CredentialVerifyShrinkRequest {
	s.ImageUrl = &v
	return s
}

func (s *CredentialVerifyShrinkRequest) SetIsCheck(v string) *CredentialVerifyShrinkRequest {
	s.IsCheck = &v
	return s
}

func (s *CredentialVerifyShrinkRequest) SetIsOCR(v string) *CredentialVerifyShrinkRequest {
	s.IsOCR = &v
	return s
}

func (s *CredentialVerifyShrinkRequest) SetMerchantDetailShrink(v string) *CredentialVerifyShrinkRequest {
	s.MerchantDetailShrink = &v
	return s
}

func (s *CredentialVerifyShrinkRequest) SetMerchantId(v string) *CredentialVerifyShrinkRequest {
	s.MerchantId = &v
	return s
}

func (s *CredentialVerifyShrinkRequest) SetProductCode(v string) *CredentialVerifyShrinkRequest {
	s.ProductCode = &v
	return s
}

func (s *CredentialVerifyShrinkRequest) SetPrompt(v string) *CredentialVerifyShrinkRequest {
	s.Prompt = &v
	return s
}

func (s *CredentialVerifyShrinkRequest) SetPromptModel(v string) *CredentialVerifyShrinkRequest {
	s.PromptModel = &v
	return s
}

func (s *CredentialVerifyShrinkRequest) SetUserName(v string) *CredentialVerifyShrinkRequest {
	s.UserName = &v
	return s
}

func (s *CredentialVerifyShrinkRequest) Validate() error {
	return dara.Validate(s)
}
