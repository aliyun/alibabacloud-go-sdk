// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
	"io"
)

type iCredentialVerifyV2AdvanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCertNum(v string) *CredentialVerifyV2AdvanceRequest
	GetCertNum() *string
	SetCredName(v string) *CredentialVerifyV2AdvanceRequest
	GetCredName() *string
	SetCredType(v string) *CredentialVerifyV2AdvanceRequest
	GetCredType() *string
	SetIdentifyNum(v string) *CredentialVerifyV2AdvanceRequest
	GetIdentifyNum() *string
	SetImageContext(v string) *CredentialVerifyV2AdvanceRequest
	GetImageContext() *string
	SetImageFileObject(v io.Reader) *CredentialVerifyV2AdvanceRequest
	GetImageFileObject() io.Reader
	SetImageUrl(v string) *CredentialVerifyV2AdvanceRequest
	GetImageUrl() *string
	SetIsCheck(v string) *CredentialVerifyV2AdvanceRequest
	GetIsCheck() *string
	SetIsOcr(v string) *CredentialVerifyV2AdvanceRequest
	GetIsOcr() *string
	SetMerchantDetail(v []*CredentialVerifyV2AdvanceRequestMerchantDetail) *CredentialVerifyV2AdvanceRequest
	GetMerchantDetail() []*CredentialVerifyV2AdvanceRequestMerchantDetail
	SetMerchantId(v string) *CredentialVerifyV2AdvanceRequest
	GetMerchantId() *string
	SetProductCode(v string) *CredentialVerifyV2AdvanceRequest
	GetProductCode() *string
	SetPrompt(v string) *CredentialVerifyV2AdvanceRequest
	GetPrompt() *string
	SetPromptModel(v string) *CredentialVerifyV2AdvanceRequest
	GetPromptModel() *string
	SetUserName(v string) *CredentialVerifyV2AdvanceRequest
	GetUserName() *string
}

type CredentialVerifyV2AdvanceRequest struct {
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
	ImageFileObject io.Reader `json:"ImageFile,omitempty" xml:"ImageFile,omitempty"`
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
	MerchantDetail []*CredentialVerifyV2AdvanceRequestMerchantDetail `json:"MerchantDetail,omitempty" xml:"MerchantDetail,omitempty" type:"Repeated"`
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

func (s CredentialVerifyV2AdvanceRequest) String() string {
	return dara.Prettify(s)
}

func (s CredentialVerifyV2AdvanceRequest) GoString() string {
	return s.String()
}

func (s *CredentialVerifyV2AdvanceRequest) GetCertNum() *string {
	return s.CertNum
}

func (s *CredentialVerifyV2AdvanceRequest) GetCredName() *string {
	return s.CredName
}

func (s *CredentialVerifyV2AdvanceRequest) GetCredType() *string {
	return s.CredType
}

func (s *CredentialVerifyV2AdvanceRequest) GetIdentifyNum() *string {
	return s.IdentifyNum
}

func (s *CredentialVerifyV2AdvanceRequest) GetImageContext() *string {
	return s.ImageContext
}

func (s *CredentialVerifyV2AdvanceRequest) GetImageFileObject() io.Reader {
	return s.ImageFileObject
}

func (s *CredentialVerifyV2AdvanceRequest) GetImageUrl() *string {
	return s.ImageUrl
}

func (s *CredentialVerifyV2AdvanceRequest) GetIsCheck() *string {
	return s.IsCheck
}

func (s *CredentialVerifyV2AdvanceRequest) GetIsOcr() *string {
	return s.IsOcr
}

func (s *CredentialVerifyV2AdvanceRequest) GetMerchantDetail() []*CredentialVerifyV2AdvanceRequestMerchantDetail {
	return s.MerchantDetail
}

func (s *CredentialVerifyV2AdvanceRequest) GetMerchantId() *string {
	return s.MerchantId
}

func (s *CredentialVerifyV2AdvanceRequest) GetProductCode() *string {
	return s.ProductCode
}

func (s *CredentialVerifyV2AdvanceRequest) GetPrompt() *string {
	return s.Prompt
}

func (s *CredentialVerifyV2AdvanceRequest) GetPromptModel() *string {
	return s.PromptModel
}

func (s *CredentialVerifyV2AdvanceRequest) GetUserName() *string {
	return s.UserName
}

func (s *CredentialVerifyV2AdvanceRequest) SetCertNum(v string) *CredentialVerifyV2AdvanceRequest {
	s.CertNum = &v
	return s
}

func (s *CredentialVerifyV2AdvanceRequest) SetCredName(v string) *CredentialVerifyV2AdvanceRequest {
	s.CredName = &v
	return s
}

func (s *CredentialVerifyV2AdvanceRequest) SetCredType(v string) *CredentialVerifyV2AdvanceRequest {
	s.CredType = &v
	return s
}

func (s *CredentialVerifyV2AdvanceRequest) SetIdentifyNum(v string) *CredentialVerifyV2AdvanceRequest {
	s.IdentifyNum = &v
	return s
}

func (s *CredentialVerifyV2AdvanceRequest) SetImageContext(v string) *CredentialVerifyV2AdvanceRequest {
	s.ImageContext = &v
	return s
}

func (s *CredentialVerifyV2AdvanceRequest) SetImageFileObject(v io.Reader) *CredentialVerifyV2AdvanceRequest {
	s.ImageFileObject = v
	return s
}

func (s *CredentialVerifyV2AdvanceRequest) SetImageUrl(v string) *CredentialVerifyV2AdvanceRequest {
	s.ImageUrl = &v
	return s
}

func (s *CredentialVerifyV2AdvanceRequest) SetIsCheck(v string) *CredentialVerifyV2AdvanceRequest {
	s.IsCheck = &v
	return s
}

func (s *CredentialVerifyV2AdvanceRequest) SetIsOcr(v string) *CredentialVerifyV2AdvanceRequest {
	s.IsOcr = &v
	return s
}

func (s *CredentialVerifyV2AdvanceRequest) SetMerchantDetail(v []*CredentialVerifyV2AdvanceRequestMerchantDetail) *CredentialVerifyV2AdvanceRequest {
	s.MerchantDetail = v
	return s
}

func (s *CredentialVerifyV2AdvanceRequest) SetMerchantId(v string) *CredentialVerifyV2AdvanceRequest {
	s.MerchantId = &v
	return s
}

func (s *CredentialVerifyV2AdvanceRequest) SetProductCode(v string) *CredentialVerifyV2AdvanceRequest {
	s.ProductCode = &v
	return s
}

func (s *CredentialVerifyV2AdvanceRequest) SetPrompt(v string) *CredentialVerifyV2AdvanceRequest {
	s.Prompt = &v
	return s
}

func (s *CredentialVerifyV2AdvanceRequest) SetPromptModel(v string) *CredentialVerifyV2AdvanceRequest {
	s.PromptModel = &v
	return s
}

func (s *CredentialVerifyV2AdvanceRequest) SetUserName(v string) *CredentialVerifyV2AdvanceRequest {
	s.UserName = &v
	return s
}

func (s *CredentialVerifyV2AdvanceRequest) Validate() error {
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

type CredentialVerifyV2AdvanceRequestMerchantDetail struct {
	// This feature is offline. This parameter no longer takes effect.
	//
	// example:
	//
	// -
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// This feature is offline. This parameter no longer takes effect.
	//
	// example:
	//
	// -
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CredentialVerifyV2AdvanceRequestMerchantDetail) String() string {
	return dara.Prettify(s)
}

func (s CredentialVerifyV2AdvanceRequestMerchantDetail) GoString() string {
	return s.String()
}

func (s *CredentialVerifyV2AdvanceRequestMerchantDetail) GetKey() *string {
	return s.Key
}

func (s *CredentialVerifyV2AdvanceRequestMerchantDetail) GetValue() *string {
	return s.Value
}

func (s *CredentialVerifyV2AdvanceRequestMerchantDetail) SetKey(v string) *CredentialVerifyV2AdvanceRequestMerchantDetail {
	s.Key = &v
	return s
}

func (s *CredentialVerifyV2AdvanceRequestMerchantDetail) SetValue(v string) *CredentialVerifyV2AdvanceRequestMerchantDetail {
	s.Value = &v
	return s
}

func (s *CredentialVerifyV2AdvanceRequestMerchantDetail) Validate() error {
	return dara.Validate(s)
}
