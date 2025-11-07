// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCredentialVerifyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCertNum(v string) *CredentialVerifyRequest
	GetCertNum() *string
	SetCredName(v string) *CredentialVerifyRequest
	GetCredName() *string
	SetCredType(v string) *CredentialVerifyRequest
	GetCredType() *string
	SetIdentifyNum(v string) *CredentialVerifyRequest
	GetIdentifyNum() *string
	SetImageContext(v string) *CredentialVerifyRequest
	GetImageContext() *string
	SetImageUrl(v string) *CredentialVerifyRequest
	GetImageUrl() *string
	SetIsCheck(v string) *CredentialVerifyRequest
	GetIsCheck() *string
	SetIsOCR(v string) *CredentialVerifyRequest
	GetIsOCR() *string
	SetMerchantDetail(v []*CredentialVerifyRequestMerchantDetail) *CredentialVerifyRequest
	GetMerchantDetail() []*CredentialVerifyRequestMerchantDetail
	SetMerchantId(v string) *CredentialVerifyRequest
	GetMerchantId() *string
	SetProductCode(v string) *CredentialVerifyRequest
	GetProductCode() *string
	SetPrompt(v string) *CredentialVerifyRequest
	GetPrompt() *string
	SetPromptModel(v string) *CredentialVerifyRequest
	GetPromptModel() *string
	SetUserName(v string) *CredentialVerifyRequest
	GetUserName() *string
}

type CredentialVerifyRequest struct {
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
	MerchantDetail []*CredentialVerifyRequestMerchantDetail `json:"MerchantDetail,omitempty" xml:"MerchantDetail,omitempty" type:"Repeated"`
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

func (s CredentialVerifyRequest) String() string {
	return dara.Prettify(s)
}

func (s CredentialVerifyRequest) GoString() string {
	return s.String()
}

func (s *CredentialVerifyRequest) GetCertNum() *string {
	return s.CertNum
}

func (s *CredentialVerifyRequest) GetCredName() *string {
	return s.CredName
}

func (s *CredentialVerifyRequest) GetCredType() *string {
	return s.CredType
}

func (s *CredentialVerifyRequest) GetIdentifyNum() *string {
	return s.IdentifyNum
}

func (s *CredentialVerifyRequest) GetImageContext() *string {
	return s.ImageContext
}

func (s *CredentialVerifyRequest) GetImageUrl() *string {
	return s.ImageUrl
}

func (s *CredentialVerifyRequest) GetIsCheck() *string {
	return s.IsCheck
}

func (s *CredentialVerifyRequest) GetIsOCR() *string {
	return s.IsOCR
}

func (s *CredentialVerifyRequest) GetMerchantDetail() []*CredentialVerifyRequestMerchantDetail {
	return s.MerchantDetail
}

func (s *CredentialVerifyRequest) GetMerchantId() *string {
	return s.MerchantId
}

func (s *CredentialVerifyRequest) GetProductCode() *string {
	return s.ProductCode
}

func (s *CredentialVerifyRequest) GetPrompt() *string {
	return s.Prompt
}

func (s *CredentialVerifyRequest) GetPromptModel() *string {
	return s.PromptModel
}

func (s *CredentialVerifyRequest) GetUserName() *string {
	return s.UserName
}

func (s *CredentialVerifyRequest) SetCertNum(v string) *CredentialVerifyRequest {
	s.CertNum = &v
	return s
}

func (s *CredentialVerifyRequest) SetCredName(v string) *CredentialVerifyRequest {
	s.CredName = &v
	return s
}

func (s *CredentialVerifyRequest) SetCredType(v string) *CredentialVerifyRequest {
	s.CredType = &v
	return s
}

func (s *CredentialVerifyRequest) SetIdentifyNum(v string) *CredentialVerifyRequest {
	s.IdentifyNum = &v
	return s
}

func (s *CredentialVerifyRequest) SetImageContext(v string) *CredentialVerifyRequest {
	s.ImageContext = &v
	return s
}

func (s *CredentialVerifyRequest) SetImageUrl(v string) *CredentialVerifyRequest {
	s.ImageUrl = &v
	return s
}

func (s *CredentialVerifyRequest) SetIsCheck(v string) *CredentialVerifyRequest {
	s.IsCheck = &v
	return s
}

func (s *CredentialVerifyRequest) SetIsOCR(v string) *CredentialVerifyRequest {
	s.IsOCR = &v
	return s
}

func (s *CredentialVerifyRequest) SetMerchantDetail(v []*CredentialVerifyRequestMerchantDetail) *CredentialVerifyRequest {
	s.MerchantDetail = v
	return s
}

func (s *CredentialVerifyRequest) SetMerchantId(v string) *CredentialVerifyRequest {
	s.MerchantId = &v
	return s
}

func (s *CredentialVerifyRequest) SetProductCode(v string) *CredentialVerifyRequest {
	s.ProductCode = &v
	return s
}

func (s *CredentialVerifyRequest) SetPrompt(v string) *CredentialVerifyRequest {
	s.Prompt = &v
	return s
}

func (s *CredentialVerifyRequest) SetPromptModel(v string) *CredentialVerifyRequest {
	s.PromptModel = &v
	return s
}

func (s *CredentialVerifyRequest) SetUserName(v string) *CredentialVerifyRequest {
	s.UserName = &v
	return s
}

func (s *CredentialVerifyRequest) Validate() error {
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

type CredentialVerifyRequestMerchantDetail struct {
	// The private key of the certificate.
	//
	// >  If this parameter is specified, you must also specify **CertName*	- and **Cert**. If **CertName**, **Cert**, and **Key*	- are specified, you do not need to specify **CertId**.
	//
	// example:
	//
	// keyword
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// Keyword value.
	//
	// example:
	//
	// -
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CredentialVerifyRequestMerchantDetail) String() string {
	return dara.Prettify(s)
}

func (s CredentialVerifyRequestMerchantDetail) GoString() string {
	return s.String()
}

func (s *CredentialVerifyRequestMerchantDetail) GetKey() *string {
	return s.Key
}

func (s *CredentialVerifyRequestMerchantDetail) GetValue() *string {
	return s.Value
}

func (s *CredentialVerifyRequestMerchantDetail) SetKey(v string) *CredentialVerifyRequestMerchantDetail {
	s.Key = &v
	return s
}

func (s *CredentialVerifyRequestMerchantDetail) SetValue(v string) *CredentialVerifyRequestMerchantDetail {
	s.Value = &v
	return s
}

func (s *CredentialVerifyRequestMerchantDetail) Validate() error {
	return dara.Validate(s)
}
