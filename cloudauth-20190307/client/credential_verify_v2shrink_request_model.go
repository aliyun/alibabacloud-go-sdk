// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCredentialVerifyV2ShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCertNum(v string) *CredentialVerifyV2ShrinkRequest
	GetCertNum() *string
	SetCredName(v string) *CredentialVerifyV2ShrinkRequest
	GetCredName() *string
	SetCredType(v string) *CredentialVerifyV2ShrinkRequest
	GetCredType() *string
	SetIdentifyNum(v string) *CredentialVerifyV2ShrinkRequest
	GetIdentifyNum() *string
	SetImageContext(v string) *CredentialVerifyV2ShrinkRequest
	GetImageContext() *string
	SetImageFile(v string) *CredentialVerifyV2ShrinkRequest
	GetImageFile() *string
	SetImageUrl(v string) *CredentialVerifyV2ShrinkRequest
	GetImageUrl() *string
	SetIsCheck(v string) *CredentialVerifyV2ShrinkRequest
	GetIsCheck() *string
	SetIsOcr(v string) *CredentialVerifyV2ShrinkRequest
	GetIsOcr() *string
	SetMerchantDetailShrink(v string) *CredentialVerifyV2ShrinkRequest
	GetMerchantDetailShrink() *string
	SetMerchantId(v string) *CredentialVerifyV2ShrinkRequest
	GetMerchantId() *string
	SetProductCode(v string) *CredentialVerifyV2ShrinkRequest
	GetProductCode() *string
	SetPrompt(v string) *CredentialVerifyV2ShrinkRequest
	GetPrompt() *string
	SetPromptModel(v string) *CredentialVerifyV2ShrinkRequest
	GetPromptModel() *string
	SetUserName(v string) *CredentialVerifyV2ShrinkRequest
	GetUserName() *string
}

type CredentialVerifyV2ShrinkRequest struct {
	// The certificate number.
	//
	// example:
	//
	// 4601*****
	CertNum *string `json:"CertNum,omitempty" xml:"CertNum,omitempty"`
	// The credential name. Valid values:
	//
	// - 01: personal card or certificate
	//
	//   - 0101: ID card
	//
	//   - 0102: bank card
	//
	//   - 0104: teacher qualification certificate
	//
	//   - 0107: student ID card
	//
	// - 02: business scenario
	//
	//   - 0201: storefront photo
	//
	//   - 0202: counter photo
	//
	//   - 0203: scene photo
	//
	// - 03: enterprise qualification
	//
	//   - 0301: business license.
	//
	// example:
	//
	// 0104
	CredName *string `json:"CredName,omitempty" xml:"CredName,omitempty"`
	// The credential type. Valid values:
	//
	// - 01: personal card or certificate
	//
	// - 02: business scenario
	//
	// - 03: enterprise qualification.
	//
	// example:
	//
	// 01
	CredType *string `json:"CredType,omitempty" xml:"CredType,omitempty"`
	// The ID card number.
	//
	// example:
	//
	// 4****************1
	IdentifyNum *string `json:"IdentifyNum,omitempty" xml:"IdentifyNum,omitempty"`
	// The Base64-encoded image. Specify one of imageUrl, imageFile, or imageContext.
	//
	// example:
	//
	// 无
	ImageContext *string `json:"ImageContext,omitempty" xml:"ImageContext,omitempty"`
	// The input stream of the image. Specify one of imageUrl, imageFile, or imageContext.
	//
	// example:
	//
	// 无
	ImageFile *string `json:"ImageFile,omitempty" xml:"ImageFile,omitempty"`
	// The URL of the image. Specify one of imageUrl, imageFile, or imageContext.
	//
	// example:
	//
	// http://marry.momocdn.com/avatar/3B/B6/3BB6527E-7467-926E-1048-B43614F20CC420230803_L.jpg
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	// Specifies whether to enable authoritative verification. Valid values:
	//
	// - **0**: Disabled.
	//
	// - **1**: Enabled.
	//
	// 	Danger: Deprecated.
	//
	// example:
	//
	// 0
	IsCheck *string `json:"IsCheck,omitempty" xml:"IsCheck,omitempty"`
	// Specifies whether to enable OCR.
	//
	// 	Danger: Deprecated.
	//
	// example:
	//
	// 0
	IsOcr *string `json:"IsOcr,omitempty" xml:"IsOcr,omitempty"`
	// This feature is offline. This parameter no longer takes effect.
	MerchantDetailShrink *string `json:"MerchantDetail,omitempty" xml:"MerchantDetail,omitempty"`
	// The merchant ID. This parameter is required when CredName is set to 02.
	//
	// example:
	//
	// 无。
	MerchantId *string `json:"MerchantId,omitempty" xml:"MerchantId,omitempty"`
	// The call mode. Valid values:
	//
	// - ANTI_FAKE_CHECK (default): image anti-forgery detection.
	//
	// example:
	//
	// ANTI_FAKE_CHECK
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// This feature is offline. This parameter no longer takes effect.
	//
	// example:
	//
	// -
	Prompt *string `json:"Prompt,omitempty" xml:"Prompt,omitempty"`
	// This feature is offline. This parameter no longer takes effect.
	//
	// example:
	//
	// -
	PromptModel *string `json:"PromptModel,omitempty" xml:"PromptModel,omitempty"`
	// The name.
	//
	// example:
	//
	// 张三
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s CredentialVerifyV2ShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CredentialVerifyV2ShrinkRequest) GoString() string {
	return s.String()
}

func (s *CredentialVerifyV2ShrinkRequest) GetCertNum() *string {
	return s.CertNum
}

func (s *CredentialVerifyV2ShrinkRequest) GetCredName() *string {
	return s.CredName
}

func (s *CredentialVerifyV2ShrinkRequest) GetCredType() *string {
	return s.CredType
}

func (s *CredentialVerifyV2ShrinkRequest) GetIdentifyNum() *string {
	return s.IdentifyNum
}

func (s *CredentialVerifyV2ShrinkRequest) GetImageContext() *string {
	return s.ImageContext
}

func (s *CredentialVerifyV2ShrinkRequest) GetImageFile() *string {
	return s.ImageFile
}

func (s *CredentialVerifyV2ShrinkRequest) GetImageUrl() *string {
	return s.ImageUrl
}

func (s *CredentialVerifyV2ShrinkRequest) GetIsCheck() *string {
	return s.IsCheck
}

func (s *CredentialVerifyV2ShrinkRequest) GetIsOcr() *string {
	return s.IsOcr
}

func (s *CredentialVerifyV2ShrinkRequest) GetMerchantDetailShrink() *string {
	return s.MerchantDetailShrink
}

func (s *CredentialVerifyV2ShrinkRequest) GetMerchantId() *string {
	return s.MerchantId
}

func (s *CredentialVerifyV2ShrinkRequest) GetProductCode() *string {
	return s.ProductCode
}

func (s *CredentialVerifyV2ShrinkRequest) GetPrompt() *string {
	return s.Prompt
}

func (s *CredentialVerifyV2ShrinkRequest) GetPromptModel() *string {
	return s.PromptModel
}

func (s *CredentialVerifyV2ShrinkRequest) GetUserName() *string {
	return s.UserName
}

func (s *CredentialVerifyV2ShrinkRequest) SetCertNum(v string) *CredentialVerifyV2ShrinkRequest {
	s.CertNum = &v
	return s
}

func (s *CredentialVerifyV2ShrinkRequest) SetCredName(v string) *CredentialVerifyV2ShrinkRequest {
	s.CredName = &v
	return s
}

func (s *CredentialVerifyV2ShrinkRequest) SetCredType(v string) *CredentialVerifyV2ShrinkRequest {
	s.CredType = &v
	return s
}

func (s *CredentialVerifyV2ShrinkRequest) SetIdentifyNum(v string) *CredentialVerifyV2ShrinkRequest {
	s.IdentifyNum = &v
	return s
}

func (s *CredentialVerifyV2ShrinkRequest) SetImageContext(v string) *CredentialVerifyV2ShrinkRequest {
	s.ImageContext = &v
	return s
}

func (s *CredentialVerifyV2ShrinkRequest) SetImageFile(v string) *CredentialVerifyV2ShrinkRequest {
	s.ImageFile = &v
	return s
}

func (s *CredentialVerifyV2ShrinkRequest) SetImageUrl(v string) *CredentialVerifyV2ShrinkRequest {
	s.ImageUrl = &v
	return s
}

func (s *CredentialVerifyV2ShrinkRequest) SetIsCheck(v string) *CredentialVerifyV2ShrinkRequest {
	s.IsCheck = &v
	return s
}

func (s *CredentialVerifyV2ShrinkRequest) SetIsOcr(v string) *CredentialVerifyV2ShrinkRequest {
	s.IsOcr = &v
	return s
}

func (s *CredentialVerifyV2ShrinkRequest) SetMerchantDetailShrink(v string) *CredentialVerifyV2ShrinkRequest {
	s.MerchantDetailShrink = &v
	return s
}

func (s *CredentialVerifyV2ShrinkRequest) SetMerchantId(v string) *CredentialVerifyV2ShrinkRequest {
	s.MerchantId = &v
	return s
}

func (s *CredentialVerifyV2ShrinkRequest) SetProductCode(v string) *CredentialVerifyV2ShrinkRequest {
	s.ProductCode = &v
	return s
}

func (s *CredentialVerifyV2ShrinkRequest) SetPrompt(v string) *CredentialVerifyV2ShrinkRequest {
	s.Prompt = &v
	return s
}

func (s *CredentialVerifyV2ShrinkRequest) SetPromptModel(v string) *CredentialVerifyV2ShrinkRequest {
	s.PromptModel = &v
	return s
}

func (s *CredentialVerifyV2ShrinkRequest) SetUserName(v string) *CredentialVerifyV2ShrinkRequest {
	s.UserName = &v
	return s
}

func (s *CredentialVerifyV2ShrinkRequest) Validate() error {
	return dara.Validate(s)
}
