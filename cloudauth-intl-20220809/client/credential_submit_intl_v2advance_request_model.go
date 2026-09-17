// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
	"io"
)

type iCredentialSubmitIntlV2AdvanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCheckRuleConfig(v string) *CredentialSubmitIntlV2AdvanceRequest
	GetCheckRuleConfig() *string
	SetCredentialOcrPictureBase64(v string) *CredentialSubmitIntlV2AdvanceRequest
	GetCredentialOcrPictureBase64() *string
	SetCredentialOcrPictureFileObject(v io.Reader) *CredentialSubmitIntlV2AdvanceRequest
	GetCredentialOcrPictureFileObject() io.Reader
	SetCredentialOcrPictureUrl(v string) *CredentialSubmitIntlV2AdvanceRequest
	GetCredentialOcrPictureUrl() *string
	SetDocType(v string) *CredentialSubmitIntlV2AdvanceRequest
	GetDocType() *string
	SetFileInputType(v string) *CredentialSubmitIntlV2AdvanceRequest
	GetFileInputType() *string
	SetFraudCheck(v string) *CredentialSubmitIntlV2AdvanceRequest
	GetFraudCheck() *string
	SetIdQuality(v string) *CredentialSubmitIntlV2AdvanceRequest
	GetIdQuality() *string
	SetMerchantBizId(v string) *CredentialSubmitIntlV2AdvanceRequest
	GetMerchantBizId() *string
	SetOcrArea(v string) *CredentialSubmitIntlV2AdvanceRequest
	GetOcrArea() *string
	SetOcrTranslation(v string) *CredentialSubmitIntlV2AdvanceRequest
	GetOcrTranslation() *string
	SetOcrValueStandard(v string) *CredentialSubmitIntlV2AdvanceRequest
	GetOcrValueStandard() *string
	SetProductCode(v string) *CredentialSubmitIntlV2AdvanceRequest
	GetProductCode() *string
	SetSceneCode(v string) *CredentialSubmitIntlV2AdvanceRequest
	GetSceneCode() *string
}

type CredentialSubmitIntlV2AdvanceRequest struct {
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
	CredentialOcrPictureFileObject io.Reader `json:"CredentialOcrPictureFile,omitempty" xml:"CredentialOcrPictureFile,omitempty"`
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

func (s CredentialSubmitIntlV2AdvanceRequest) String() string {
	return dara.Prettify(s)
}

func (s CredentialSubmitIntlV2AdvanceRequest) GoString() string {
	return s.String()
}

func (s *CredentialSubmitIntlV2AdvanceRequest) GetCheckRuleConfig() *string {
	return s.CheckRuleConfig
}

func (s *CredentialSubmitIntlV2AdvanceRequest) GetCredentialOcrPictureBase64() *string {
	return s.CredentialOcrPictureBase64
}

func (s *CredentialSubmitIntlV2AdvanceRequest) GetCredentialOcrPictureFileObject() io.Reader {
	return s.CredentialOcrPictureFileObject
}

func (s *CredentialSubmitIntlV2AdvanceRequest) GetCredentialOcrPictureUrl() *string {
	return s.CredentialOcrPictureUrl
}

func (s *CredentialSubmitIntlV2AdvanceRequest) GetDocType() *string {
	return s.DocType
}

func (s *CredentialSubmitIntlV2AdvanceRequest) GetFileInputType() *string {
	return s.FileInputType
}

func (s *CredentialSubmitIntlV2AdvanceRequest) GetFraudCheck() *string {
	return s.FraudCheck
}

func (s *CredentialSubmitIntlV2AdvanceRequest) GetIdQuality() *string {
	return s.IdQuality
}

func (s *CredentialSubmitIntlV2AdvanceRequest) GetMerchantBizId() *string {
	return s.MerchantBizId
}

func (s *CredentialSubmitIntlV2AdvanceRequest) GetOcrArea() *string {
	return s.OcrArea
}

func (s *CredentialSubmitIntlV2AdvanceRequest) GetOcrTranslation() *string {
	return s.OcrTranslation
}

func (s *CredentialSubmitIntlV2AdvanceRequest) GetOcrValueStandard() *string {
	return s.OcrValueStandard
}

func (s *CredentialSubmitIntlV2AdvanceRequest) GetProductCode() *string {
	return s.ProductCode
}

func (s *CredentialSubmitIntlV2AdvanceRequest) GetSceneCode() *string {
	return s.SceneCode
}

func (s *CredentialSubmitIntlV2AdvanceRequest) SetCheckRuleConfig(v string) *CredentialSubmitIntlV2AdvanceRequest {
	s.CheckRuleConfig = &v
	return s
}

func (s *CredentialSubmitIntlV2AdvanceRequest) SetCredentialOcrPictureBase64(v string) *CredentialSubmitIntlV2AdvanceRequest {
	s.CredentialOcrPictureBase64 = &v
	return s
}

func (s *CredentialSubmitIntlV2AdvanceRequest) SetCredentialOcrPictureFileObject(v io.Reader) *CredentialSubmitIntlV2AdvanceRequest {
	s.CredentialOcrPictureFileObject = v
	return s
}

func (s *CredentialSubmitIntlV2AdvanceRequest) SetCredentialOcrPictureUrl(v string) *CredentialSubmitIntlV2AdvanceRequest {
	s.CredentialOcrPictureUrl = &v
	return s
}

func (s *CredentialSubmitIntlV2AdvanceRequest) SetDocType(v string) *CredentialSubmitIntlV2AdvanceRequest {
	s.DocType = &v
	return s
}

func (s *CredentialSubmitIntlV2AdvanceRequest) SetFileInputType(v string) *CredentialSubmitIntlV2AdvanceRequest {
	s.FileInputType = &v
	return s
}

func (s *CredentialSubmitIntlV2AdvanceRequest) SetFraudCheck(v string) *CredentialSubmitIntlV2AdvanceRequest {
	s.FraudCheck = &v
	return s
}

func (s *CredentialSubmitIntlV2AdvanceRequest) SetIdQuality(v string) *CredentialSubmitIntlV2AdvanceRequest {
	s.IdQuality = &v
	return s
}

func (s *CredentialSubmitIntlV2AdvanceRequest) SetMerchantBizId(v string) *CredentialSubmitIntlV2AdvanceRequest {
	s.MerchantBizId = &v
	return s
}

func (s *CredentialSubmitIntlV2AdvanceRequest) SetOcrArea(v string) *CredentialSubmitIntlV2AdvanceRequest {
	s.OcrArea = &v
	return s
}

func (s *CredentialSubmitIntlV2AdvanceRequest) SetOcrTranslation(v string) *CredentialSubmitIntlV2AdvanceRequest {
	s.OcrTranslation = &v
	return s
}

func (s *CredentialSubmitIntlV2AdvanceRequest) SetOcrValueStandard(v string) *CredentialSubmitIntlV2AdvanceRequest {
	s.OcrValueStandard = &v
	return s
}

func (s *CredentialSubmitIntlV2AdvanceRequest) SetProductCode(v string) *CredentialSubmitIntlV2AdvanceRequest {
	s.ProductCode = &v
	return s
}

func (s *CredentialSubmitIntlV2AdvanceRequest) SetSceneCode(v string) *CredentialSubmitIntlV2AdvanceRequest {
	s.SceneCode = &v
	return s
}

func (s *CredentialSubmitIntlV2AdvanceRequest) Validate() error {
	return dara.Validate(s)
}
