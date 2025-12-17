// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDocOcrMaxRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthorize(v string) *DocOcrMaxRequest
	GetAuthorize() *string
	SetDocPage(v string) *DocOcrMaxRequest
	GetDocPage() *string
	SetDocType(v string) *DocOcrMaxRequest
	GetDocType() *string
	SetIdOcrPictureBase64(v string) *DocOcrMaxRequest
	GetIdOcrPictureBase64() *string
	SetIdOcrPictureUrl(v string) *DocOcrMaxRequest
	GetIdOcrPictureUrl() *string
	SetIdSpoof(v string) *DocOcrMaxRequest
	GetIdSpoof() *string
	SetIdThreshold(v string) *DocOcrMaxRequest
	GetIdThreshold() *string
	SetMerchantBizId(v string) *DocOcrMaxRequest
	GetMerchantBizId() *string
	SetMerchantUserId(v string) *DocOcrMaxRequest
	GetMerchantUserId() *string
	SetOcrModel(v string) *DocOcrMaxRequest
	GetOcrModel() *string
	SetOcrValueStandard(v string) *DocOcrMaxRequest
	GetOcrValueStandard() *string
	SetProductCode(v string) *DocOcrMaxRequest
	GetProductCode() *string
	SetPrompt(v string) *DocOcrMaxRequest
	GetPrompt() *string
	SetSceneCode(v string) *DocOcrMaxRequest
	GetSceneCode() *string
	SetSpoof(v string) *DocOcrMaxRequest
	GetSpoof() *string
}

type DocOcrMaxRequest struct {
	// Specifies whether to enable verification with an authoritative data source to enhance document anti-spoofing capabilities.
	//
	// - **T**: Enable
	//
	// - **F**: Disable (default)
	//
	// >
	//
	// > - **Applicable document types**: Chinese resident ID cards (CHN01001) and Chinese mainland driver\\"s licenses (CHN02001).
	//
	// > - **Data transfer declaration**: If you enable this parameter, you agree to transfer the user\\"s name and certificate number to an authoritative data source in the Chinese mainland for consistency verification.
	//
	// > - **Performance impact:*	- After you enable this feature, the response time of the API operation increases by 1 to 2 seconds. Adjust the timeout setting.
	//
	// example:
	//
	// T
	Authorize *string `json:"Authorize,omitempty" xml:"Authorize,omitempty"`
	// Page expected to be recognized
	//
	// - 01 (default): ID portrait.
	//
	// - 02: Back of the certificate
	//
	// example:
	//
	// 01
	DocPage *string `json:"DocPage,omitempty" xml:"DocPage,omitempty"`
	// Document type.
	//
	// Format: Country (region) code + document type abbreviation + page (optional)
	//
	// Note: If provided, it will automatically check if it matches the model recognition result; if empty, the document type will be returned after model recognition.
	//
	// example:
	//
	// CNSSC01
	DocType *string `json:"DocType,omitempty" xml:"DocType,omitempty"`
	// Document image, base64 encoded binary stream
	//
	// example:
	//
	// base64
	IdOcrPictureBase64 *string `json:"IdOcrPictureBase64,omitempty" xml:"IdOcrPictureBase64,omitempty"`
	// Document image URL
	//
	// example:
	//
	// https://***********.oss-cn-hangzhou.aliyuncs.com/1669520556530-expo/default/face/20221127114236530_w3kx2e6t.jpg
	IdOcrPictureUrl *string `json:"IdOcrPictureUrl,omitempty" xml:"IdOcrPictureUrl,omitempty"`
	// Whether to turn on the certificate anti-counterfeiting function:
	//
	// - T: open
	//
	// - F (default): not turned on.
	//
	// example:
	//
	// F
	IdSpoof *string `json:"IdSpoof,omitempty" xml:"IdSpoof,omitempty"`
	// Custom OCR quality detection threshold mode:
	//
	// - 0: System default
	//
	// - 1: Strict mode
	//
	// - 2: Lenient mode
	//
	// - 3 (default): Disable quality detection
	//
	// example:
	//
	// 0
	IdThreshold *string `json:"IdThreshold,omitempty" xml:"IdThreshold,omitempty"`
	// A unique business identifier defined by the merchant, used for subsequent problem localization and troubleshooting. It supports a combination of letters and numbers, with a maximum length of 32 characters. Please ensure its uniqueness.
	//
	// example:
	//
	// e0c34a77f5ac40a5aa5e6ed20c******
	MerchantBizId *string `json:"MerchantBizId,omitempty" xml:"MerchantBizId,omitempty"`
	// Your custom user ID or other identifiers that can uniquely identify a specific user, such as a phone number or email address. It is strongly recommended to pre-desensitize the value of this field, for example, by hashing it.
	//
	// example:
	//
	// 123456789
	MerchantUserId *string `json:"MerchantUserId,omitempty" xml:"MerchantUserId,omitempty"`
	// OCR recognition mode.
	//
	// 0: General document mode.
	//
	// 1: Custom mode.
	//
	// example:
	//
	// 0
	OcrModel *string `json:"OcrModel,omitempty" xml:"OcrModel,omitempty"`
	// Specifies whether to return additional OCR fields in a standardized format:
	//
	// - **0**: No (default)
	//
	// - **1**: Yes
	//
	// example:
	//
	// 0
	OcrValueStandard *string `json:"OcrValueStandard,omitempty" xml:"OcrValueStandard,omitempty"`
	// The product solution to be integrated.
	//
	// Value: ID_OCR_MAX
	//
	// example:
	//
	// ID_OCR_MAX
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// Prompt (for custom mode)
	//
	// example:
	//
	// xxxocr识别
	Prompt *string `json:"Prompt,omitempty" xml:"Prompt,omitempty"`
	// Custom scene code, used to distinguish business scenarios, a 10-digit number.
	//
	// example:
	//
	// 1234567890
	SceneCode *string `json:"SceneCode,omitempty" xml:"SceneCode,omitempty"`
	// Whether to enable document anti-counterfeiting function, default is not enabled.
	//
	// - T: Enable document anti-counterfeiting function.
	//
	// - F: Do not enable.
	//
	// example:
	//
	// F
	Spoof *string `json:"Spoof,omitempty" xml:"Spoof,omitempty"`
}

func (s DocOcrMaxRequest) String() string {
	return dara.Prettify(s)
}

func (s DocOcrMaxRequest) GoString() string {
	return s.String()
}

func (s *DocOcrMaxRequest) GetAuthorize() *string {
	return s.Authorize
}

func (s *DocOcrMaxRequest) GetDocPage() *string {
	return s.DocPage
}

func (s *DocOcrMaxRequest) GetDocType() *string {
	return s.DocType
}

func (s *DocOcrMaxRequest) GetIdOcrPictureBase64() *string {
	return s.IdOcrPictureBase64
}

func (s *DocOcrMaxRequest) GetIdOcrPictureUrl() *string {
	return s.IdOcrPictureUrl
}

func (s *DocOcrMaxRequest) GetIdSpoof() *string {
	return s.IdSpoof
}

func (s *DocOcrMaxRequest) GetIdThreshold() *string {
	return s.IdThreshold
}

func (s *DocOcrMaxRequest) GetMerchantBizId() *string {
	return s.MerchantBizId
}

func (s *DocOcrMaxRequest) GetMerchantUserId() *string {
	return s.MerchantUserId
}

func (s *DocOcrMaxRequest) GetOcrModel() *string {
	return s.OcrModel
}

func (s *DocOcrMaxRequest) GetOcrValueStandard() *string {
	return s.OcrValueStandard
}

func (s *DocOcrMaxRequest) GetProductCode() *string {
	return s.ProductCode
}

func (s *DocOcrMaxRequest) GetPrompt() *string {
	return s.Prompt
}

func (s *DocOcrMaxRequest) GetSceneCode() *string {
	return s.SceneCode
}

func (s *DocOcrMaxRequest) GetSpoof() *string {
	return s.Spoof
}

func (s *DocOcrMaxRequest) SetAuthorize(v string) *DocOcrMaxRequest {
	s.Authorize = &v
	return s
}

func (s *DocOcrMaxRequest) SetDocPage(v string) *DocOcrMaxRequest {
	s.DocPage = &v
	return s
}

func (s *DocOcrMaxRequest) SetDocType(v string) *DocOcrMaxRequest {
	s.DocType = &v
	return s
}

func (s *DocOcrMaxRequest) SetIdOcrPictureBase64(v string) *DocOcrMaxRequest {
	s.IdOcrPictureBase64 = &v
	return s
}

func (s *DocOcrMaxRequest) SetIdOcrPictureUrl(v string) *DocOcrMaxRequest {
	s.IdOcrPictureUrl = &v
	return s
}

func (s *DocOcrMaxRequest) SetIdSpoof(v string) *DocOcrMaxRequest {
	s.IdSpoof = &v
	return s
}

func (s *DocOcrMaxRequest) SetIdThreshold(v string) *DocOcrMaxRequest {
	s.IdThreshold = &v
	return s
}

func (s *DocOcrMaxRequest) SetMerchantBizId(v string) *DocOcrMaxRequest {
	s.MerchantBizId = &v
	return s
}

func (s *DocOcrMaxRequest) SetMerchantUserId(v string) *DocOcrMaxRequest {
	s.MerchantUserId = &v
	return s
}

func (s *DocOcrMaxRequest) SetOcrModel(v string) *DocOcrMaxRequest {
	s.OcrModel = &v
	return s
}

func (s *DocOcrMaxRequest) SetOcrValueStandard(v string) *DocOcrMaxRequest {
	s.OcrValueStandard = &v
	return s
}

func (s *DocOcrMaxRequest) SetProductCode(v string) *DocOcrMaxRequest {
	s.ProductCode = &v
	return s
}

func (s *DocOcrMaxRequest) SetPrompt(v string) *DocOcrMaxRequest {
	s.Prompt = &v
	return s
}

func (s *DocOcrMaxRequest) SetSceneCode(v string) *DocOcrMaxRequest {
	s.SceneCode = &v
	return s
}

func (s *DocOcrMaxRequest) SetSpoof(v string) *DocOcrMaxRequest {
	s.Spoof = &v
	return s
}

func (s *DocOcrMaxRequest) Validate() error {
	return dara.Validate(s)
}
