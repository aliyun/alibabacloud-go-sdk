// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
	"io"
)

type iCredentialRecognitionIntlV2AdvanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCheckRuleConfig(v string) *CredentialRecognitionIntlV2AdvanceRequest
	GetCheckRuleConfig() *string
	SetCredentialOcrPictureBase64(v string) *CredentialRecognitionIntlV2AdvanceRequest
	GetCredentialOcrPictureBase64() *string
	SetCredentialOcrPictureFileObject(v io.Reader) *CredentialRecognitionIntlV2AdvanceRequest
	GetCredentialOcrPictureFileObject() io.Reader
	SetCredentialOcrPictureUrl(v string) *CredentialRecognitionIntlV2AdvanceRequest
	GetCredentialOcrPictureUrl() *string
	SetDocType(v string) *CredentialRecognitionIntlV2AdvanceRequest
	GetDocType() *string
	SetFileInputType(v string) *CredentialRecognitionIntlV2AdvanceRequest
	GetFileInputType() *string
	SetFraudCheck(v string) *CredentialRecognitionIntlV2AdvanceRequest
	GetFraudCheck() *string
	SetIdQuality(v string) *CredentialRecognitionIntlV2AdvanceRequest
	GetIdQuality() *string
	SetOcrArea(v string) *CredentialRecognitionIntlV2AdvanceRequest
	GetOcrArea() *string
	SetOcrTranslation(v string) *CredentialRecognitionIntlV2AdvanceRequest
	GetOcrTranslation() *string
	SetOcrValueStandard(v string) *CredentialRecognitionIntlV2AdvanceRequest
	GetOcrValueStandard() *string
	SetProductCode(v string) *CredentialRecognitionIntlV2AdvanceRequest
	GetProductCode() *string
}

type CredentialRecognitionIntlV2AdvanceRequest struct {
	// The field validation rule configuration, in JSON string format.
	//
	// example:
	//
	// {
	//
	// 	"address_rule": "Includes Adrress Hangzhou***",
	//
	// 	"name_rule": "Includes Name  Zhang*",
	//
	// 	"date_of_issue_rule": "Whthin 2026.05.20"
	//
	// }
	CheckRuleConfig *string `json:"CheckRuleConfig,omitempty" xml:"CheckRuleConfig,omitempty"`
	// The Base64-encoded image. If you choose to pass in the image by using IdOcrPictureBase64 (Base64-encoded photo), check the photo size and do not pass in an excessively large photo.
	//
	// example:
	//
	// base64
	CredentialOcrPictureBase64 *string `json:"CredentialOcrPictureBase64,omitempty" xml:"CredentialOcrPictureBase64,omitempty"`
	// The image file stream.
	//
	// example:
	//
	// InputStream
	CredentialOcrPictureFileObject io.Reader `json:"CredentialOcrPictureFile,omitempty" xml:"CredentialOcrPictureFile,omitempty"`
	// The URL of the image. The URL must be a publicly accessible HTTP or HTTPS link.
	//
	// example:
	//
	// https://***
	CredentialOcrPictureUrl *string `json:"CredentialOcrPictureUrl,omitempty" xml:"CredentialOcrPictureUrl,omitempty"`
	// The credential type. Valid values:
	//
	// - 01: transaction credential (including electronic bill images for water, electricity, gas, credit cards, and other types)
	//
	// - 02: vehicle registration certificate
	//
	// - 03: transfer transaction record
	//
	// - 04: POA address proof
	//
	// This parameter is required.
	//
	// example:
	//
	// 01
	DocType *string `json:"DocType,omitempty" xml:"DocType,omitempty"`
	// The input material type. Valid values:
	//
	// - IMAGE (default): image
	//
	// - PDF: PDF format
	//
	// example:
	//
	// IMAGE
	FileInputType *string `json:"FileInputType,omitempty" xml:"FileInputType,omitempty"`
	// Specifies whether to enable tampering detection. Valid values:
	//
	// - true: Enabled.
	//
	// - false: Disabled.
	//
	// This parameter is required.
	//
	// example:
	//
	// false
	FraudCheck *string `json:"FraudCheck,omitempty" xml:"FraudCheck,omitempty"`
	// Specifies whether to enable quality detection. Valid values: Y (enabled) and N (disabled).
	//
	// example:
	//
	// Y
	IdQuality *string `json:"IdQuality,omitempty" xml:"IdQuality,omitempty"`
	// The extraction type. Valid values:
	//
	// - 0101: electronic bill address and name module (extracts address and name modules through intelligent analysis)
	//
	// - 0201: Thailand vehicle registration certificate
	//
	// - 0301: transfer transaction amount information
	//
	// - 0401: POA credential extraction information
	//
	// This parameter is required.
	//
	// example:
	//
	// 0101
	OcrArea *string `json:"OcrArea,omitempty" xml:"OcrArea,omitempty"`
	// Specifies whether to enable translation. Valid values: 0 (disabled) and 1 (enabled).
	//
	// example:
	//
	// 1
	OcrTranslation *string `json:"OcrTranslation,omitempty" xml:"OcrTranslation,omitempty"`
	// Specifies whether to enable OCR result normalization. Valid values: 0 (disabled) and 1 (enabled).
	//
	// example:
	//
	// 0
	OcrValueStandard *string `json:"OcrValueStandard,omitempty" xml:"OcrValueStandard,omitempty"`
	// The product solution to use. Set the value to CREDENTIAL_RECOGNITION.
	//
	// This parameter is required.
	//
	// example:
	//
	// CREDENTIAL_RECOGNITION
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
}

func (s CredentialRecognitionIntlV2AdvanceRequest) String() string {
	return dara.Prettify(s)
}

func (s CredentialRecognitionIntlV2AdvanceRequest) GoString() string {
	return s.String()
}

func (s *CredentialRecognitionIntlV2AdvanceRequest) GetCheckRuleConfig() *string {
	return s.CheckRuleConfig
}

func (s *CredentialRecognitionIntlV2AdvanceRequest) GetCredentialOcrPictureBase64() *string {
	return s.CredentialOcrPictureBase64
}

func (s *CredentialRecognitionIntlV2AdvanceRequest) GetCredentialOcrPictureFileObject() io.Reader {
	return s.CredentialOcrPictureFileObject
}

func (s *CredentialRecognitionIntlV2AdvanceRequest) GetCredentialOcrPictureUrl() *string {
	return s.CredentialOcrPictureUrl
}

func (s *CredentialRecognitionIntlV2AdvanceRequest) GetDocType() *string {
	return s.DocType
}

func (s *CredentialRecognitionIntlV2AdvanceRequest) GetFileInputType() *string {
	return s.FileInputType
}

func (s *CredentialRecognitionIntlV2AdvanceRequest) GetFraudCheck() *string {
	return s.FraudCheck
}

func (s *CredentialRecognitionIntlV2AdvanceRequest) GetIdQuality() *string {
	return s.IdQuality
}

func (s *CredentialRecognitionIntlV2AdvanceRequest) GetOcrArea() *string {
	return s.OcrArea
}

func (s *CredentialRecognitionIntlV2AdvanceRequest) GetOcrTranslation() *string {
	return s.OcrTranslation
}

func (s *CredentialRecognitionIntlV2AdvanceRequest) GetOcrValueStandard() *string {
	return s.OcrValueStandard
}

func (s *CredentialRecognitionIntlV2AdvanceRequest) GetProductCode() *string {
	return s.ProductCode
}

func (s *CredentialRecognitionIntlV2AdvanceRequest) SetCheckRuleConfig(v string) *CredentialRecognitionIntlV2AdvanceRequest {
	s.CheckRuleConfig = &v
	return s
}

func (s *CredentialRecognitionIntlV2AdvanceRequest) SetCredentialOcrPictureBase64(v string) *CredentialRecognitionIntlV2AdvanceRequest {
	s.CredentialOcrPictureBase64 = &v
	return s
}

func (s *CredentialRecognitionIntlV2AdvanceRequest) SetCredentialOcrPictureFileObject(v io.Reader) *CredentialRecognitionIntlV2AdvanceRequest {
	s.CredentialOcrPictureFileObject = v
	return s
}

func (s *CredentialRecognitionIntlV2AdvanceRequest) SetCredentialOcrPictureUrl(v string) *CredentialRecognitionIntlV2AdvanceRequest {
	s.CredentialOcrPictureUrl = &v
	return s
}

func (s *CredentialRecognitionIntlV2AdvanceRequest) SetDocType(v string) *CredentialRecognitionIntlV2AdvanceRequest {
	s.DocType = &v
	return s
}

func (s *CredentialRecognitionIntlV2AdvanceRequest) SetFileInputType(v string) *CredentialRecognitionIntlV2AdvanceRequest {
	s.FileInputType = &v
	return s
}

func (s *CredentialRecognitionIntlV2AdvanceRequest) SetFraudCheck(v string) *CredentialRecognitionIntlV2AdvanceRequest {
	s.FraudCheck = &v
	return s
}

func (s *CredentialRecognitionIntlV2AdvanceRequest) SetIdQuality(v string) *CredentialRecognitionIntlV2AdvanceRequest {
	s.IdQuality = &v
	return s
}

func (s *CredentialRecognitionIntlV2AdvanceRequest) SetOcrArea(v string) *CredentialRecognitionIntlV2AdvanceRequest {
	s.OcrArea = &v
	return s
}

func (s *CredentialRecognitionIntlV2AdvanceRequest) SetOcrTranslation(v string) *CredentialRecognitionIntlV2AdvanceRequest {
	s.OcrTranslation = &v
	return s
}

func (s *CredentialRecognitionIntlV2AdvanceRequest) SetOcrValueStandard(v string) *CredentialRecognitionIntlV2AdvanceRequest {
	s.OcrValueStandard = &v
	return s
}

func (s *CredentialRecognitionIntlV2AdvanceRequest) SetProductCode(v string) *CredentialRecognitionIntlV2AdvanceRequest {
	s.ProductCode = &v
	return s
}

func (s *CredentialRecognitionIntlV2AdvanceRequest) Validate() error {
	return dara.Validate(s)
}
