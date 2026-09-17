// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCredentialSubmitIntlV2Request interface {
	dara.Model
	String() string
	GoString() string
	SetCheckRuleConfig(v string) *CredentialSubmitIntlV2Request
	GetCheckRuleConfig() *string
	SetCredentialOcrPictureBase64(v string) *CredentialSubmitIntlV2Request
	GetCredentialOcrPictureBase64() *string
	SetCredentialOcrPictureFile(v string) *CredentialSubmitIntlV2Request
	GetCredentialOcrPictureFile() *string
	SetCredentialOcrPictureUrl(v string) *CredentialSubmitIntlV2Request
	GetCredentialOcrPictureUrl() *string
	SetDocType(v string) *CredentialSubmitIntlV2Request
	GetDocType() *string
	SetFileInputType(v string) *CredentialSubmitIntlV2Request
	GetFileInputType() *string
	SetFraudCheck(v string) *CredentialSubmitIntlV2Request
	GetFraudCheck() *string
	SetIdQuality(v string) *CredentialSubmitIntlV2Request
	GetIdQuality() *string
	SetMerchantBizId(v string) *CredentialSubmitIntlV2Request
	GetMerchantBizId() *string
	SetOcrArea(v string) *CredentialSubmitIntlV2Request
	GetOcrArea() *string
	SetOcrTranslation(v string) *CredentialSubmitIntlV2Request
	GetOcrTranslation() *string
	SetOcrValueStandard(v string) *CredentialSubmitIntlV2Request
	GetOcrValueStandard() *string
	SetProductCode(v string) *CredentialSubmitIntlV2Request
	GetProductCode() *string
	SetSceneCode(v string) *CredentialSubmitIntlV2Request
	GetSceneCode() *string
}

type CredentialSubmitIntlV2Request struct {
	// The field validation rule configuration. The value is a JSON string.
	//
	// example:
	//
	// {
	//
	// 	"address_rule": "Includes Adrress Hangzhou ***",
	//
	// 	"name_rule": "Includes Name  Zhang*",
	//
	// 	"date_of_issue_rule": "Whthin 2026.05.20"
	//
	// }
	CheckRuleConfig *string `json:"CheckRuleConfig,omitempty" xml:"CheckRuleConfig,omitempty"`
	// The Base64-encoded image. If you choose this method to submit a photo, check the photo size and do not submit an excessively large photo.
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
	// The URL of the image. The URL must be a publicly accessible HTTP or HTTPS link.
	//
	// example:
	//
	// https://***
	CredentialOcrPictureUrl *string `json:"CredentialOcrPictureUrl,omitempty" xml:"CredentialOcrPictureUrl,omitempty"`
	// The credential type. Valid values:
	//
	// - 01: Transaction voucher, which includes electronic bill images for utilities such as water, electricity, gas, and credit cards.
	//
	// - 02: Vehicle registration certificate.
	//
	// - 03: Transfer transaction record.
	//
	// - 04: Proof of address (POA).
	//
	// This parameter is required.
	//
	// example:
	//
	// 02
	DocType *string `json:"DocType,omitempty" xml:"DocType,omitempty"`
	// The type of the input material. Valid values:
	//
	// - IMAGE (default): image.
	//
	// - PDF: PDF format.
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
	// true
	FraudCheck *string `json:"FraudCheck,omitempty" xml:"FraudCheck,omitempty"`
	// Specifies whether to enable quality detection. Valid values: Y (enabled) and N (disabled).
	//
	// example:
	//
	// Y
	IdQuality *string `json:"IdQuality,omitempty" xml:"IdQuality,omitempty"`
	// The unique identifier of the merchant request. The value is a 32-character alphanumeric string.
	//
	// The first few characters consist of a custom abbreviation defined by the merchant, the middle part can contain a time segment, and the last part can use a random or incremental sequence.
	//
	// This parameter is required.
	//
	// example:
	//
	// dso932dsjsd22
	MerchantBizId *string `json:"MerchantBizId,omitempty" xml:"MerchantBizId,omitempty"`
	// The extraction type. Valid values:
	//
	// 0101: electronic bill address and name module (extracts address and name modules through intelligent analysis).
	//
	// 0201:
	//
	// 0301: transfer transaction amount information.
	//
	// 0401: POA credential extraction information.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0201
	OcrArea *string `json:"OcrArea,omitempty" xml:"OcrArea,omitempty"`
	// Specifies whether to enable translation. Valid values: 0 (disabled) and 1 (enabled).
	//
	// example:
	//
	// 1
	OcrTranslation *string `json:"OcrTranslation,omitempty" xml:"OcrTranslation,omitempty"`
	// Specifies whether to enable standardization of key fields recognized by OCR. Valid values:
	//
	// - 0: Disabled (default).
	//
	// - 1: Enabled.
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
	// The custom verification scenario ID. You can use this scenario ID to query related records in the console. The value is a combination of up to 10 letters, digits, or underscores.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234567890
	SceneCode *string `json:"SceneCode,omitempty" xml:"SceneCode,omitempty"`
}

func (s CredentialSubmitIntlV2Request) String() string {
	return dara.Prettify(s)
}

func (s CredentialSubmitIntlV2Request) GoString() string {
	return s.String()
}

func (s *CredentialSubmitIntlV2Request) GetCheckRuleConfig() *string {
	return s.CheckRuleConfig
}

func (s *CredentialSubmitIntlV2Request) GetCredentialOcrPictureBase64() *string {
	return s.CredentialOcrPictureBase64
}

func (s *CredentialSubmitIntlV2Request) GetCredentialOcrPictureFile() *string {
	return s.CredentialOcrPictureFile
}

func (s *CredentialSubmitIntlV2Request) GetCredentialOcrPictureUrl() *string {
	return s.CredentialOcrPictureUrl
}

func (s *CredentialSubmitIntlV2Request) GetDocType() *string {
	return s.DocType
}

func (s *CredentialSubmitIntlV2Request) GetFileInputType() *string {
	return s.FileInputType
}

func (s *CredentialSubmitIntlV2Request) GetFraudCheck() *string {
	return s.FraudCheck
}

func (s *CredentialSubmitIntlV2Request) GetIdQuality() *string {
	return s.IdQuality
}

func (s *CredentialSubmitIntlV2Request) GetMerchantBizId() *string {
	return s.MerchantBizId
}

func (s *CredentialSubmitIntlV2Request) GetOcrArea() *string {
	return s.OcrArea
}

func (s *CredentialSubmitIntlV2Request) GetOcrTranslation() *string {
	return s.OcrTranslation
}

func (s *CredentialSubmitIntlV2Request) GetOcrValueStandard() *string {
	return s.OcrValueStandard
}

func (s *CredentialSubmitIntlV2Request) GetProductCode() *string {
	return s.ProductCode
}

func (s *CredentialSubmitIntlV2Request) GetSceneCode() *string {
	return s.SceneCode
}

func (s *CredentialSubmitIntlV2Request) SetCheckRuleConfig(v string) *CredentialSubmitIntlV2Request {
	s.CheckRuleConfig = &v
	return s
}

func (s *CredentialSubmitIntlV2Request) SetCredentialOcrPictureBase64(v string) *CredentialSubmitIntlV2Request {
	s.CredentialOcrPictureBase64 = &v
	return s
}

func (s *CredentialSubmitIntlV2Request) SetCredentialOcrPictureFile(v string) *CredentialSubmitIntlV2Request {
	s.CredentialOcrPictureFile = &v
	return s
}

func (s *CredentialSubmitIntlV2Request) SetCredentialOcrPictureUrl(v string) *CredentialSubmitIntlV2Request {
	s.CredentialOcrPictureUrl = &v
	return s
}

func (s *CredentialSubmitIntlV2Request) SetDocType(v string) *CredentialSubmitIntlV2Request {
	s.DocType = &v
	return s
}

func (s *CredentialSubmitIntlV2Request) SetFileInputType(v string) *CredentialSubmitIntlV2Request {
	s.FileInputType = &v
	return s
}

func (s *CredentialSubmitIntlV2Request) SetFraudCheck(v string) *CredentialSubmitIntlV2Request {
	s.FraudCheck = &v
	return s
}

func (s *CredentialSubmitIntlV2Request) SetIdQuality(v string) *CredentialSubmitIntlV2Request {
	s.IdQuality = &v
	return s
}

func (s *CredentialSubmitIntlV2Request) SetMerchantBizId(v string) *CredentialSubmitIntlV2Request {
	s.MerchantBizId = &v
	return s
}

func (s *CredentialSubmitIntlV2Request) SetOcrArea(v string) *CredentialSubmitIntlV2Request {
	s.OcrArea = &v
	return s
}

func (s *CredentialSubmitIntlV2Request) SetOcrTranslation(v string) *CredentialSubmitIntlV2Request {
	s.OcrTranslation = &v
	return s
}

func (s *CredentialSubmitIntlV2Request) SetOcrValueStandard(v string) *CredentialSubmitIntlV2Request {
	s.OcrValueStandard = &v
	return s
}

func (s *CredentialSubmitIntlV2Request) SetProductCode(v string) *CredentialSubmitIntlV2Request {
	s.ProductCode = &v
	return s
}

func (s *CredentialSubmitIntlV2Request) SetSceneCode(v string) *CredentialSubmitIntlV2Request {
	s.SceneCode = &v
	return s
}

func (s *CredentialSubmitIntlV2Request) Validate() error {
	return dara.Validate(s)
}
