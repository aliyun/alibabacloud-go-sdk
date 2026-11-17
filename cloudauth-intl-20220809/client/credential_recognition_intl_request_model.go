// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCredentialRecognitionIntlRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCheckRuleConfig(v string) *CredentialRecognitionIntlRequest
	GetCheckRuleConfig() *string
	SetCredentialOcrPictureBase64(v string) *CredentialRecognitionIntlRequest
	GetCredentialOcrPictureBase64() *string
	SetCredentialOcrPictureUrl(v string) *CredentialRecognitionIntlRequest
	GetCredentialOcrPictureUrl() *string
	SetDocType(v string) *CredentialRecognitionIntlRequest
	GetDocType() *string
	SetFileInputType(v string) *CredentialRecognitionIntlRequest
	GetFileInputType() *string
	SetFraudCheck(v string) *CredentialRecognitionIntlRequest
	GetFraudCheck() *string
	SetIdQuality(v string) *CredentialRecognitionIntlRequest
	GetIdQuality() *string
	SetOcrArea(v string) *CredentialRecognitionIntlRequest
	GetOcrArea() *string
	SetOcrTranslation(v string) *CredentialRecognitionIntlRequest
	GetOcrTranslation() *string
	SetOcrValueStandard(v string) *CredentialRecognitionIntlRequest
	GetOcrValueStandard() *string
	SetProductCode(v string) *CredentialRecognitionIntlRequest
	GetProductCode() *string
}

type CredentialRecognitionIntlRequest struct {
	// The field validation rule configuration. The value is a JSON string.
	//
	// example:
	//
	// {
	//
	// 	"address_rule": "Includes Adrress Hangzhou***",
	//
	// 	"name_rule": "Includes Name Zhang*",
	//
	// 	"date_of_issue_rule": "Whthin 2026.05.20"
	//
	// }
	CheckRuleConfig *string `json:"CheckRuleConfig,omitempty" xml:"CheckRuleConfig,omitempty"`
	// The Base64-encoded image. If you use the IdOcrPictureBase64 parameter to pass the Base64-encoded photo, check the photo size and do not pass oversized photos.
	//
	// example:
	//
	// base64
	CredentialOcrPictureBase64 *string `json:"CredentialOcrPictureBase64,omitempty" xml:"CredentialOcrPictureBase64,omitempty"`
	// The URL of the image. The URL must be a publicly accessible HTTP or HTTPS link.
	//
	// example:
	//
	// https://***
	CredentialOcrPictureUrl *string `json:"CredentialOcrPictureUrl,omitempty" xml:"CredentialOcrPictureUrl,omitempty"`
	// The credential type. Valid values:
	//
	// - 01: transaction credentials, including electronic bills for utilities such as water, electricity, and gas, as well as credit card statements.
	//
	// This parameter is required.
	//
	// example:
	//
	// 01
	DocType *string `json:"DocType,omitempty" xml:"DocType,omitempty"`
	// The material input type. Valid values:
	//
	// - IMAGE: image. This is the default value.
	//
	// - PDF: PDF format.
	//
	// example:
	//
	// IMAGE
	FileInputType *string `json:"FileInputType,omitempty" xml:"FileInputType,omitempty"`
	// Specifies whether to enable tampering detection. Valid values:
	//
	// - true: enabled.
	//
	// - false: disabled.
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	FraudCheck *string `json:"FraudCheck,omitempty" xml:"FraudCheck,omitempty"`
	// Specifies whether to enable the quality check. Valid values:
	//
	// - Y: enabled.
	//
	// - N: disabled.
	//
	// example:
	//
	// Y
	IdQuality *string `json:"IdQuality,omitempty" xml:"IdQuality,omitempty"`
	// The extraction type. Valid values:
	//
	// - 0101: electronic bill address and name module. The address and name are extracted through intelligent analysis.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0101
	OcrArea *string `json:"OcrArea,omitempty" xml:"OcrArea,omitempty"`
	// Specifies whether to enable translation. Valid values:
	//
	// - 0: disabled.
	//
	// - 1: enabled.
	//
	// example:
	//
	// 1
	OcrTranslation *string `json:"OcrTranslation,omitempty" xml:"OcrTranslation,omitempty"`
	// Specifies whether to enable OCR result standardization. Valid values:
	//
	// - 0: disabled.
	//
	// - 1: enabled.
	//
	// example:
	//
	// 1
	OcrValueStandard *string `json:"OcrValueStandard,omitempty" xml:"OcrValueStandard,omitempty"`
	// The product solution to be integrated. Valid value: CREDENTIAL_RECOGNITION.
	//
	// This parameter is required.
	//
	// example:
	//
	// CREDENTIAL_RECOGNITION
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
}

func (s CredentialRecognitionIntlRequest) String() string {
	return dara.Prettify(s)
}

func (s CredentialRecognitionIntlRequest) GoString() string {
	return s.String()
}

func (s *CredentialRecognitionIntlRequest) GetCheckRuleConfig() *string {
	return s.CheckRuleConfig
}

func (s *CredentialRecognitionIntlRequest) GetCredentialOcrPictureBase64() *string {
	return s.CredentialOcrPictureBase64
}

func (s *CredentialRecognitionIntlRequest) GetCredentialOcrPictureUrl() *string {
	return s.CredentialOcrPictureUrl
}

func (s *CredentialRecognitionIntlRequest) GetDocType() *string {
	return s.DocType
}

func (s *CredentialRecognitionIntlRequest) GetFileInputType() *string {
	return s.FileInputType
}

func (s *CredentialRecognitionIntlRequest) GetFraudCheck() *string {
	return s.FraudCheck
}

func (s *CredentialRecognitionIntlRequest) GetIdQuality() *string {
	return s.IdQuality
}

func (s *CredentialRecognitionIntlRequest) GetOcrArea() *string {
	return s.OcrArea
}

func (s *CredentialRecognitionIntlRequest) GetOcrTranslation() *string {
	return s.OcrTranslation
}

func (s *CredentialRecognitionIntlRequest) GetOcrValueStandard() *string {
	return s.OcrValueStandard
}

func (s *CredentialRecognitionIntlRequest) GetProductCode() *string {
	return s.ProductCode
}

func (s *CredentialRecognitionIntlRequest) SetCheckRuleConfig(v string) *CredentialRecognitionIntlRequest {
	s.CheckRuleConfig = &v
	return s
}

func (s *CredentialRecognitionIntlRequest) SetCredentialOcrPictureBase64(v string) *CredentialRecognitionIntlRequest {
	s.CredentialOcrPictureBase64 = &v
	return s
}

func (s *CredentialRecognitionIntlRequest) SetCredentialOcrPictureUrl(v string) *CredentialRecognitionIntlRequest {
	s.CredentialOcrPictureUrl = &v
	return s
}

func (s *CredentialRecognitionIntlRequest) SetDocType(v string) *CredentialRecognitionIntlRequest {
	s.DocType = &v
	return s
}

func (s *CredentialRecognitionIntlRequest) SetFileInputType(v string) *CredentialRecognitionIntlRequest {
	s.FileInputType = &v
	return s
}

func (s *CredentialRecognitionIntlRequest) SetFraudCheck(v string) *CredentialRecognitionIntlRequest {
	s.FraudCheck = &v
	return s
}

func (s *CredentialRecognitionIntlRequest) SetIdQuality(v string) *CredentialRecognitionIntlRequest {
	s.IdQuality = &v
	return s
}

func (s *CredentialRecognitionIntlRequest) SetOcrArea(v string) *CredentialRecognitionIntlRequest {
	s.OcrArea = &v
	return s
}

func (s *CredentialRecognitionIntlRequest) SetOcrTranslation(v string) *CredentialRecognitionIntlRequest {
	s.OcrTranslation = &v
	return s
}

func (s *CredentialRecognitionIntlRequest) SetOcrValueStandard(v string) *CredentialRecognitionIntlRequest {
	s.OcrValueStandard = &v
	return s
}

func (s *CredentialRecognitionIntlRequest) SetProductCode(v string) *CredentialRecognitionIntlRequest {
	s.ProductCode = &v
	return s
}

func (s *CredentialRecognitionIntlRequest) Validate() error {
	return dara.Validate(s)
}
