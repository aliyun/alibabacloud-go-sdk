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
	// Specifies whether to enable authoritative data source verification to enhance the anti-spoofing capability of identity documents. Valid values:
	//
	// - **T**: Enabled.
	//
	// - **F**: Not enabled (default).
	//
	// >
	//
	// > - **Applicable identity document types**: Chinese resident identity card (CHN01001) and Chinese mainland driver\\"s license (CHN02001).
	//
	// > - **Data transmission statement**: Enabling this parameter indicates that you agree to transmit the user\\"s name and identity document number to an authoritative data source in the Chinese mainland for consistency verification.
	//
	// > - **Performance impact**: After this parameter is enabled, the API response time increases by about 1 to 2 seconds. Adjust the timeout settings accordingly.
	//
	// example:
	//
	// T
	Authorize *string `json:"Authorize,omitempty" xml:"Authorize,omitempty"`
	// The page to be recognized. Valid values:
	//
	// - 01: portrait side of the identity document (default)
	//
	// - 02: back side of the identity document
	//
	// example:
	//
	// 01
	DocPage *string `json:"DocPage,omitempty" xml:"DocPage,omitempty"`
	// The type of the identity document.
	//
	// - Format: country code + identity document type abbreviation + page (optional).
	//
	// Note
	//
	// - OcrModel = 0: DocType is required and specifies the identity document type. The existing logic remains unchanged.
	//
	// - OcrModel = 1 or 2: DocType must be left empty.
	//
	// example:
	//
	// CNSSC01
	DocType *string `json:"DocType,omitempty" xml:"DocType,omitempty"`
	// The Base64-encoded image of the identity document. If you choose to pass the identity document image by using the IdOcrPictureBase64 (Base64-encoded image) method, check the image size and do not pass an excessively large image.
	//
	// example:
	//
	// base64
	IdOcrPictureBase64 *string `json:"IdOcrPictureBase64,omitempty" xml:"IdOcrPictureBase64,omitempty"`
	// The URL of the identity document image. The URL must be an HTTP or HTTPS link that is accessible over the Internet.
	//
	// example:
	//
	// https://***********.oss-cn-hangzhou.aliyuncs.com/1669520556530-expo/default/face/20221127114236530_w3kx2e6t.jpg
	IdOcrPictureUrl *string `json:"IdOcrPictureUrl,omitempty" xml:"IdOcrPictureUrl,omitempty"`
	// Specifies whether to enable the identity document anti-spoofing feature. Valid values:
	//
	// - T: Enabled.
	//
	// - F: Not enabled (default).
	//
	// example:
	//
	// F
	IdSpoof *string `json:"IdSpoof,omitempty" xml:"IdSpoof,omitempty"`
	// The custom OCR quality check threshold mode. Valid values:
	//
	// - 0: system default
	//
	// - 1: strict mode
	//
	// - 2: loose mode
	//
	// - 3: disable quality check (default)
	//
	// example:
	//
	// 0
	IdThreshold *string `json:"IdThreshold,omitempty" xml:"IdThreshold,omitempty"`
	// The custom unique business identifier, which is used to locate and troubleshoot issues later. The value can be up to 32 characters in length and can contain letters and digits. Make sure that the value is unique.
	//
	// example:
	//
	// e0c34a77f5ac40a5aa5e6ed20c******
	MerchantBizId *string `json:"MerchantBizId,omitempty" xml:"MerchantBizId,omitempty"`
	// The custom user ID or other identifiers that can identify a specific user, such as a mobile phone number or an email address. Desensitize the value of this field in advance, such as by hashing the value.
	//
	// example:
	//
	// 123456789
	MerchantUserId *string `json:"MerchantUserId,omitempty" xml:"MerchantUserId,omitempty"`
	// The OCR recognition mode. Valid values:
	//
	// - 0: general identity document recognition mode (default)
	//
	// - 1: automatic identity document classification mode
	//
	// - 2: automatic identity document classification and general recognition mode
	//
	// example:
	//
	// 0
	OcrModel *string `json:"OcrModel,omitempty" xml:"OcrModel,omitempty"`
	// Specifies whether to enable the standardization of key fields in OCR recognition. Valid values:
	//
	// - 0: no (default)
	//
	// - 1: yes
	//
	// example:
	//
	// 0
	OcrValueStandard *string `json:"OcrValueStandard,omitempty" xml:"OcrValueStandard,omitempty"`
	// The product solution to be integrated. Valid value: ID_OCR_MAX.
	//
	// example:
	//
	// ID_OCR_MAX
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// **[Deprecated]**
	//
	// example:
	//
	// Deprecated
	Prompt *string `json:"Prompt,omitempty" xml:"Prompt,omitempty"`
	// The custom authentication scene ID, which is used to query related records by entering this scene ID in the console later. The value can be up to 10 characters in length and can contain letters, digits, and underscores.
	//
	// example:
	//
	// 1234567890
	SceneCode *string `json:"SceneCode,omitempty" xml:"SceneCode,omitempty"`
	// **[Deprecated]**
	//
	// example:
	//
	// Deprecated
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
