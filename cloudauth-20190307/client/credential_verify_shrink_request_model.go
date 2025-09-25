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
	// Relevant certificate number.
	//
	// example:
	//
	// 4601*****
	CertNum *string `json:"CertNum,omitempty" xml:"CertNum,omitempty"`
	// - 01: Personal ID cards
	//
	//   - **0101**: ID card
	//
	//   - **0102**: Bank card
	//
	//   - **0104**: Teacher qualification certificate
	//
	//   - **0107**: Student ID card
	//
	// - 02: Business scenario
	//
	//   - **0201**: Storefront photo
	//
	//   - **0202**: Counter photo
	//
	//   - **0203**: Scene photo
	//
	// - 03: Corporate qualifications
	//
	//   - **0301**: Business license
	//
	// example:
	//
	// 0104
	CredName *string `json:"CredName,omitempty" xml:"CredName,omitempty"`
	// Credential type:
	//
	// - 01: Personal ID cards
	//
	// - 02: Business scenario
	//
	// - 03: Corporate qualifications
	//
	// example:
	//
	// 01
	CredType *string `json:"CredType,omitempty" xml:"CredType,omitempty"`
	// ID number:
	//
	// Note
	//
	// Only supports the ID numbers of second-generation resident IDs and Hong Kong, Macao, and Taiwan residence permits.
	//
	// - When paramType is normal: enter the plaintext ID number.
	//
	// - When paramType is md5: first 6 digits of the ID number (plaintext) + date of birth (ciphertext) + last 4 digits of the ID number (plaintext).
	//
	// example:
	//
	// 429001********8211
	IdentifyNum *string `json:"IdentifyNum,omitempty" xml:"IdentifyNum,omitempty"`
	// Base64 encoded image, choose one from `imageUrl`, `imageFile`, or `imageContext`.
	//
	// example:
	//
	// base64
	ImageContext *string `json:"ImageContext,omitempty" xml:"ImageContext,omitempty"`
	// Image URL, choose one from `imageUrl`, `imageFile`, or `imageContext`.
	//
	// example:
	//
	// http://marry.momocdn.com/avatar/3B/B6/3BB6527E-7467-926E-1048-B43614F20CC420230803_L.jpg
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	// Whether to enable authoritative authentication
	//
	// - ****0****: No
	//
	// - **1**: Yes
	//
	// example:
	//
	// 0
	IsCheck *string `json:"IsCheck,omitempty" xml:"IsCheck,omitempty"`
	// Whether to enable OCR recognition.
	//
	// - **0**: No
	//
	// - **1**: Yes
	//
	// > IsOCR can be set to 1 only when **CredType*	- is 01.
	//
	// example:
	//
	// 1
	IsOCR *string `json:"IsOCR,omitempty" xml:"IsOCR,omitempty"`
	// Merchant details:
	//
	//
	// > This field is required when PromptModel is set to DEFAULT.
	MerchantDetailShrink *string `json:"MerchantDetail,omitempty" xml:"MerchantDetail,omitempty"`
	// Merchant ID.
	//
	// > This field is required when ****CredName***	- is set to **02**.
	//
	// example:
	//
	// 913100********KW8P
	MerchantId *string `json:"MerchantId,omitempty" xml:"MerchantId,omitempty"`
	// Invocation mode:
	//
	// - **ANTI_FAKE_CHECK**: Image anti-forgery check
	//
	// - **ANTI_FAKE_VL**: Image anti-forgery check and semantic understanding
	//
	// - **IMAGE_VL_COG**: Image semantic understanding
	//
	// Default value: ANTI_FAKE_CHECK
	//
	// > When **CredType*	- is set to 02, **ProductCode*	- can only be ANTI_FAKE_VL or IMAGE_VL_COG.
	//
	// example:
	//
	// ANTI_FAKE_CHECK
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// Customer-defined prompt content for image semantic understanding.
	//
	//
	// > This field is required when PromptModel is set to CUSTOM.
	//
	// example:
	//
	// -
	Prompt *string `json:"Prompt,omitempty" xml:"Prompt,omitempty"`
	// Prompt acquisition method for image semantic understanding:
	//
	// - **DEFAULT**: System default
	//
	// - **CUSTOM**: Customer-defined
	//
	// > When **ProductCode*	- is set to **ANTI_FAKE_VL*	- or **IMAGE_VL_COG**, this parameter must be provided.
	//
	// example:
	//
	// DEFAULT
	PromptModel *string `json:"PromptModel,omitempty" xml:"PromptModel,omitempty"`
	// UserName
	//
	// example:
	//
	// 张*
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
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
