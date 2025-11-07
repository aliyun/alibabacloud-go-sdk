// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCredentialVerifyV2Request interface {
	dara.Model
	String() string
	GoString() string
	SetCertNum(v string) *CredentialVerifyV2Request
	GetCertNum() *string
	SetCredName(v string) *CredentialVerifyV2Request
	GetCredName() *string
	SetCredType(v string) *CredentialVerifyV2Request
	GetCredType() *string
	SetIdentifyNum(v string) *CredentialVerifyV2Request
	GetIdentifyNum() *string
	SetImageContext(v string) *CredentialVerifyV2Request
	GetImageContext() *string
	SetImageFile(v string) *CredentialVerifyV2Request
	GetImageFile() *string
	SetImageUrl(v string) *CredentialVerifyV2Request
	GetImageUrl() *string
	SetIsCheck(v string) *CredentialVerifyV2Request
	GetIsCheck() *string
	SetIsOcr(v string) *CredentialVerifyV2Request
	GetIsOcr() *string
	SetMerchantDetail(v []*CredentialVerifyV2RequestMerchantDetail) *CredentialVerifyV2Request
	GetMerchantDetail() []*CredentialVerifyV2RequestMerchantDetail
	SetMerchantId(v string) *CredentialVerifyV2Request
	GetMerchantId() *string
	SetProductCode(v string) *CredentialVerifyV2Request
	GetProductCode() *string
	SetPrompt(v string) *CredentialVerifyV2Request
	GetPrompt() *string
	SetPromptModel(v string) *CredentialVerifyV2Request
	GetPromptModel() *string
	SetUserName(v string) *CredentialVerifyV2Request
	GetUserName() *string
}

type CredentialVerifyV2Request struct {
	// Relevant certificate number.
	//
	// example:
	//
	// 4601*****
	CertNum *string `json:"CertNum,omitempty" xml:"CertNum,omitempty"`
	// - 01: Personal ID cards
	//
	//   - 0101: ID card
	//
	//   - 0102: Bank card
	//
	//   - 0104: Teacher qualification certificate
	//
	//   - 0107: Student ID card
	//
	// - 02: Business scenario
	//
	//   - 0201: Storefront photo
	//
	//   - 0202: Counter photo
	//
	//   - 0203: Scene photo
	//
	// - 03: Corporate qualifications
	//
	//   - 0301: Business license
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
	// ID number.
	//
	// example:
	//
	// 4****************1
	IdentifyNum *string `json:"IdentifyNum,omitempty" xml:"IdentifyNum,omitempty"`
	// Base64 encoded image, choose one from `imageUrl`, `imageFile`, or `imageContext`.
	//
	// example:
	//
	// base64
	ImageContext *string `json:"ImageContext,omitempty" xml:"ImageContext,omitempty"`
	// Image input stream, choose one from `imageUrl`, `imageFile`, or `imageContext`.
	//
	// example:
	//
	// none
	ImageFile *string `json:"ImageFile,omitempty" xml:"ImageFile,omitempty"`
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
	// Whether to use OCR
	//
	// example:
	//
	// 0
	IsOcr *string `json:"IsOcr,omitempty" xml:"IsOcr,omitempty"`
	// Merchant details:
	//
	// MerchantName: Merchant name
	//
	// BusinessType: Industry information
	//
	// BusinessContent: Business content
	//
	// This field is required when PromptModel is set to DEFAULT.
	MerchantDetail []*CredentialVerifyV2RequestMerchantDetail `json:"MerchantDetail,omitempty" xml:"MerchantDetail,omitempty" type:"Repeated"`
	// Merchant ID. This field is required when ****CredName***	- is set to **02**.
	//
	// example:
	//
	// none
	MerchantId *string `json:"MerchantId,omitempty" xml:"MerchantId,omitempty"`
	// Invocation mode:
	//
	// - ANTI_FAKE_CHECK: Image anti-forgery check
	//
	// - ANTI_FAKE_VL: Image anti-forgery check and semantic understanding
	//
	// - IMAGE_VL_COG: Image semantic understanding
	//
	// Default value: ANTI_FAKE_CHECK
	//
	// When CredType is set to 02, ProductCode can only be ANTI_FAKE_VL or IMAGE_VL_COG.
	//
	// example:
	//
	// ANTI_FAKE_CHECK
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// Customer-defined prompt content for image semantic understanding.
	//
	// This field is required when PromptModel is set to CUSTOM.
	//
	// example:
	//
	// none
	Prompt *string `json:"Prompt,omitempty" xml:"Prompt,omitempty"`
	// Prompt acquisition method for image semantic understanding:
	//
	// - DEFAULT: System default
	//
	// - CUSTOM: Customer-defined
	//
	// Note: When ProductCode is set to ANTI_FAKE_VL or IMAGE_VL_COG, this parameter must be provided.
	//
	// example:
	//
	// DEFAULT
	PromptModel *string `json:"PromptModel,omitempty" xml:"PromptModel,omitempty"`
	// Name.
	//
	// example:
	//
	// 张三
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s CredentialVerifyV2Request) String() string {
	return dara.Prettify(s)
}

func (s CredentialVerifyV2Request) GoString() string {
	return s.String()
}

func (s *CredentialVerifyV2Request) GetCertNum() *string {
	return s.CertNum
}

func (s *CredentialVerifyV2Request) GetCredName() *string {
	return s.CredName
}

func (s *CredentialVerifyV2Request) GetCredType() *string {
	return s.CredType
}

func (s *CredentialVerifyV2Request) GetIdentifyNum() *string {
	return s.IdentifyNum
}

func (s *CredentialVerifyV2Request) GetImageContext() *string {
	return s.ImageContext
}

func (s *CredentialVerifyV2Request) GetImageFile() *string {
	return s.ImageFile
}

func (s *CredentialVerifyV2Request) GetImageUrl() *string {
	return s.ImageUrl
}

func (s *CredentialVerifyV2Request) GetIsCheck() *string {
	return s.IsCheck
}

func (s *CredentialVerifyV2Request) GetIsOcr() *string {
	return s.IsOcr
}

func (s *CredentialVerifyV2Request) GetMerchantDetail() []*CredentialVerifyV2RequestMerchantDetail {
	return s.MerchantDetail
}

func (s *CredentialVerifyV2Request) GetMerchantId() *string {
	return s.MerchantId
}

func (s *CredentialVerifyV2Request) GetProductCode() *string {
	return s.ProductCode
}

func (s *CredentialVerifyV2Request) GetPrompt() *string {
	return s.Prompt
}

func (s *CredentialVerifyV2Request) GetPromptModel() *string {
	return s.PromptModel
}

func (s *CredentialVerifyV2Request) GetUserName() *string {
	return s.UserName
}

func (s *CredentialVerifyV2Request) SetCertNum(v string) *CredentialVerifyV2Request {
	s.CertNum = &v
	return s
}

func (s *CredentialVerifyV2Request) SetCredName(v string) *CredentialVerifyV2Request {
	s.CredName = &v
	return s
}

func (s *CredentialVerifyV2Request) SetCredType(v string) *CredentialVerifyV2Request {
	s.CredType = &v
	return s
}

func (s *CredentialVerifyV2Request) SetIdentifyNum(v string) *CredentialVerifyV2Request {
	s.IdentifyNum = &v
	return s
}

func (s *CredentialVerifyV2Request) SetImageContext(v string) *CredentialVerifyV2Request {
	s.ImageContext = &v
	return s
}

func (s *CredentialVerifyV2Request) SetImageFile(v string) *CredentialVerifyV2Request {
	s.ImageFile = &v
	return s
}

func (s *CredentialVerifyV2Request) SetImageUrl(v string) *CredentialVerifyV2Request {
	s.ImageUrl = &v
	return s
}

func (s *CredentialVerifyV2Request) SetIsCheck(v string) *CredentialVerifyV2Request {
	s.IsCheck = &v
	return s
}

func (s *CredentialVerifyV2Request) SetIsOcr(v string) *CredentialVerifyV2Request {
	s.IsOcr = &v
	return s
}

func (s *CredentialVerifyV2Request) SetMerchantDetail(v []*CredentialVerifyV2RequestMerchantDetail) *CredentialVerifyV2Request {
	s.MerchantDetail = v
	return s
}

func (s *CredentialVerifyV2Request) SetMerchantId(v string) *CredentialVerifyV2Request {
	s.MerchantId = &v
	return s
}

func (s *CredentialVerifyV2Request) SetProductCode(v string) *CredentialVerifyV2Request {
	s.ProductCode = &v
	return s
}

func (s *CredentialVerifyV2Request) SetPrompt(v string) *CredentialVerifyV2Request {
	s.Prompt = &v
	return s
}

func (s *CredentialVerifyV2Request) SetPromptModel(v string) *CredentialVerifyV2Request {
	s.PromptModel = &v
	return s
}

func (s *CredentialVerifyV2Request) SetUserName(v string) *CredentialVerifyV2Request {
	s.UserName = &v
	return s
}

func (s *CredentialVerifyV2Request) Validate() error {
	if s.MerchantDetail != nil {
		for _, item := range s.MerchantDetail {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CredentialVerifyV2RequestMerchantDetail struct {
	// Keyword key.
	//
	// example:
	//
	// MerchantName
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// Keyword value.
	//
	// example:
	//
	// ***
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CredentialVerifyV2RequestMerchantDetail) String() string {
	return dara.Prettify(s)
}

func (s CredentialVerifyV2RequestMerchantDetail) GoString() string {
	return s.String()
}

func (s *CredentialVerifyV2RequestMerchantDetail) GetKey() *string {
	return s.Key
}

func (s *CredentialVerifyV2RequestMerchantDetail) GetValue() *string {
	return s.Value
}

func (s *CredentialVerifyV2RequestMerchantDetail) SetKey(v string) *CredentialVerifyV2RequestMerchantDetail {
	s.Key = &v
	return s
}

func (s *CredentialVerifyV2RequestMerchantDetail) SetValue(v string) *CredentialVerifyV2RequestMerchantDetail {
	s.Value = &v
	return s
}

func (s *CredentialVerifyV2RequestMerchantDetail) Validate() error {
	return dara.Validate(s)
}
