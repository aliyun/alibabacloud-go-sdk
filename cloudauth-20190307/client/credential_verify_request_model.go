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
	IsOCR          *string                                  `json:"IsOCR,omitempty" xml:"IsOCR,omitempty"`
	MerchantDetail []*CredentialVerifyRequestMerchantDetail `json:"MerchantDetail,omitempty" xml:"MerchantDetail,omitempty" type:"Repeated"`
	MerchantId     *string                                  `json:"MerchantId,omitempty" xml:"MerchantId,omitempty"`
	ProductCode    *string                                  `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	Prompt         *string                                  `json:"Prompt,omitempty" xml:"Prompt,omitempty"`
	PromptModel    *string                                  `json:"PromptModel,omitempty" xml:"PromptModel,omitempty"`
	UserName       *string                                  `json:"UserName,omitempty" xml:"UserName,omitempty"`
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
	return dara.Validate(s)
}

type CredentialVerifyRequestMerchantDetail struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
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
