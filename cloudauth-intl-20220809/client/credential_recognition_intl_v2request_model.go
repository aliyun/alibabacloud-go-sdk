// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCredentialRecognitionIntlV2Request interface {
	dara.Model
	String() string
	GoString() string
	SetCheckRuleConfig(v string) *CredentialRecognitionIntlV2Request
	GetCheckRuleConfig() *string
	SetCredentialOcrPictureBase64(v string) *CredentialRecognitionIntlV2Request
	GetCredentialOcrPictureBase64() *string
	SetCredentialOcrPictureFile(v string) *CredentialRecognitionIntlV2Request
	GetCredentialOcrPictureFile() *string
	SetCredentialOcrPictureUrl(v string) *CredentialRecognitionIntlV2Request
	GetCredentialOcrPictureUrl() *string
	SetDocType(v string) *CredentialRecognitionIntlV2Request
	GetDocType() *string
	SetFileInputType(v string) *CredentialRecognitionIntlV2Request
	GetFileInputType() *string
	SetFraudCheck(v string) *CredentialRecognitionIntlV2Request
	GetFraudCheck() *string
	SetIdQuality(v string) *CredentialRecognitionIntlV2Request
	GetIdQuality() *string
	SetOcrArea(v string) *CredentialRecognitionIntlV2Request
	GetOcrArea() *string
	SetOcrTranslation(v string) *CredentialRecognitionIntlV2Request
	GetOcrTranslation() *string
	SetOcrValueStandard(v string) *CredentialRecognitionIntlV2Request
	GetOcrValueStandard() *string
	SetProductCode(v string) *CredentialRecognitionIntlV2Request
	GetProductCode() *string
}

type CredentialRecognitionIntlV2Request struct {
	// The field validation rule configuration. The value is a JSON string.
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
	// The Base64-encoded image. If you choose to pass the photo by using IdOcrPictureBase64, check the photo size and do not pass excessively large photos.
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
	CredentialOcrPictureFile *string `json:"CredentialOcrPictureFile,omitempty" xml:"CredentialOcrPictureFile,omitempty"`
	// The image URL. The URL must be a publicly accessible HTTP or HTTPS link.
	//
	// example:
	//
	// https://***
	CredentialOcrPictureUrl *string `json:"CredentialOcrPictureUrl,omitempty" xml:"CredentialOcrPictureUrl,omitempty"`
	// The credential type. Valid values:
	//
	// - 01: transaction receipt, including electronic bills for utilities and credit cards
	//
	// - 02: vehicle registration certificate
	//
	// - 03: transfer transaction record
	//
	// - 04: proof of address (POA)
	//
	// This parameter is required.
	//
	// example:
	//
	// 01
	DocType *string `json:"DocType,omitempty" xml:"DocType,omitempty"`
	// The material input type. Valid values:
	//
	// - IMAGE: image
	//
	// - PDF: PDF format
	//
	// Default value: IMAGE.
	//
	// example:
	//
	// IMAGE
	FileInputType *string `json:"FileInputType,omitempty" xml:"FileInputType,omitempty"`
	// Specifies whether to enable tampering detection. Valid values:
	//
	// - true: enables tampering detection.
	//
	// - false: disables tampering detection.
	//
	// This parameter is required.
	//
	// example:
	//
	// false
	FraudCheck *string `json:"FraudCheck,omitempty" xml:"FraudCheck,omitempty"`
	// Specifies whether to enable quality detection. Valid values:
	//
	// - Y: enables quality detection.
	//
	// - N: disables quality detection.
	//
	// example:
	//
	// Y
	IdQuality *string `json:"IdQuality,omitempty" xml:"IdQuality,omitempty"`
	// The extraction type. Valid values:
	//
	// - 0101: electronic bill address and name module, which extracts the address and name module through intelligent analysis
	//
	// - 0201: Thailand vehicle registration certificate
	//
	// - 0301: transfer transaction record amount information
	//
	// - 0401: POA credential extraction information
	//
	// This parameter is required.
	//
	// example:
	//
	// 0101
	OcrArea *string `json:"OcrArea,omitempty" xml:"OcrArea,omitempty"`
	// Specifies whether to enable translation. Valid values:
	//
	// - 0: disables translation.
	//
	// - 1: enables translation.
	//
	// example:
	//
	// 1
	OcrTranslation *string `json:"OcrTranslation,omitempty" xml:"OcrTranslation,omitempty"`
	// Specifies whether to enable OCR result standardization. Valid values:
	//
	// - 0: disables OCR result standardization.
	//
	// - 1: enables OCR result standardization.
	//
	// example:
	//
	// 0
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

func (s CredentialRecognitionIntlV2Request) String() string {
	return dara.Prettify(s)
}

func (s CredentialRecognitionIntlV2Request) GoString() string {
	return s.String()
}

func (s *CredentialRecognitionIntlV2Request) GetCheckRuleConfig() *string {
	return s.CheckRuleConfig
}

func (s *CredentialRecognitionIntlV2Request) GetCredentialOcrPictureBase64() *string {
	return s.CredentialOcrPictureBase64
}

func (s *CredentialRecognitionIntlV2Request) GetCredentialOcrPictureFile() *string {
	return s.CredentialOcrPictureFile
}

func (s *CredentialRecognitionIntlV2Request) GetCredentialOcrPictureUrl() *string {
	return s.CredentialOcrPictureUrl
}

func (s *CredentialRecognitionIntlV2Request) GetDocType() *string {
	return s.DocType
}

func (s *CredentialRecognitionIntlV2Request) GetFileInputType() *string {
	return s.FileInputType
}

func (s *CredentialRecognitionIntlV2Request) GetFraudCheck() *string {
	return s.FraudCheck
}

func (s *CredentialRecognitionIntlV2Request) GetIdQuality() *string {
	return s.IdQuality
}

func (s *CredentialRecognitionIntlV2Request) GetOcrArea() *string {
	return s.OcrArea
}

func (s *CredentialRecognitionIntlV2Request) GetOcrTranslation() *string {
	return s.OcrTranslation
}

func (s *CredentialRecognitionIntlV2Request) GetOcrValueStandard() *string {
	return s.OcrValueStandard
}

func (s *CredentialRecognitionIntlV2Request) GetProductCode() *string {
	return s.ProductCode
}

func (s *CredentialRecognitionIntlV2Request) SetCheckRuleConfig(v string) *CredentialRecognitionIntlV2Request {
	s.CheckRuleConfig = &v
	return s
}

func (s *CredentialRecognitionIntlV2Request) SetCredentialOcrPictureBase64(v string) *CredentialRecognitionIntlV2Request {
	s.CredentialOcrPictureBase64 = &v
	return s
}

func (s *CredentialRecognitionIntlV2Request) SetCredentialOcrPictureFile(v string) *CredentialRecognitionIntlV2Request {
	s.CredentialOcrPictureFile = &v
	return s
}

func (s *CredentialRecognitionIntlV2Request) SetCredentialOcrPictureUrl(v string) *CredentialRecognitionIntlV2Request {
	s.CredentialOcrPictureUrl = &v
	return s
}

func (s *CredentialRecognitionIntlV2Request) SetDocType(v string) *CredentialRecognitionIntlV2Request {
	s.DocType = &v
	return s
}

func (s *CredentialRecognitionIntlV2Request) SetFileInputType(v string) *CredentialRecognitionIntlV2Request {
	s.FileInputType = &v
	return s
}

func (s *CredentialRecognitionIntlV2Request) SetFraudCheck(v string) *CredentialRecognitionIntlV2Request {
	s.FraudCheck = &v
	return s
}

func (s *CredentialRecognitionIntlV2Request) SetIdQuality(v string) *CredentialRecognitionIntlV2Request {
	s.IdQuality = &v
	return s
}

func (s *CredentialRecognitionIntlV2Request) SetOcrArea(v string) *CredentialRecognitionIntlV2Request {
	s.OcrArea = &v
	return s
}

func (s *CredentialRecognitionIntlV2Request) SetOcrTranslation(v string) *CredentialRecognitionIntlV2Request {
	s.OcrTranslation = &v
	return s
}

func (s *CredentialRecognitionIntlV2Request) SetOcrValueStandard(v string) *CredentialRecognitionIntlV2Request {
	s.OcrValueStandard = &v
	return s
}

func (s *CredentialRecognitionIntlV2Request) SetProductCode(v string) *CredentialRecognitionIntlV2Request {
	s.ProductCode = &v
	return s
}

func (s *CredentialRecognitionIntlV2Request) Validate() error {
	return dara.Validate(s)
}
