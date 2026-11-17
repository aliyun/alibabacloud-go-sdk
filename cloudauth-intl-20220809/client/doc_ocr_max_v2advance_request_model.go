// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
	"io"
)

type iDocOcrMaxV2AdvanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthorize(v string) *DocOcrMaxV2AdvanceRequest
	GetAuthorize() *string
	SetDocPage(v string) *DocOcrMaxV2AdvanceRequest
	GetDocPage() *string
	SetDocType(v string) *DocOcrMaxV2AdvanceRequest
	GetDocType() *string
	SetIdOcrPictureBase64(v string) *DocOcrMaxV2AdvanceRequest
	GetIdOcrPictureBase64() *string
	SetIdOcrPictureFileObject(v io.Reader) *DocOcrMaxV2AdvanceRequest
	GetIdOcrPictureFileObject() io.Reader
	SetIdOcrPictureUrl(v string) *DocOcrMaxV2AdvanceRequest
	GetIdOcrPictureUrl() *string
	SetIdSpoof(v string) *DocOcrMaxV2AdvanceRequest
	GetIdSpoof() *string
	SetIdThreshold(v string) *DocOcrMaxV2AdvanceRequest
	GetIdThreshold() *string
	SetMerchantBizId(v string) *DocOcrMaxV2AdvanceRequest
	GetMerchantBizId() *string
	SetMerchantUserId(v string) *DocOcrMaxV2AdvanceRequest
	GetMerchantUserId() *string
	SetOcrModel(v string) *DocOcrMaxV2AdvanceRequest
	GetOcrModel() *string
	SetOcrValueStandard(v string) *DocOcrMaxV2AdvanceRequest
	GetOcrValueStandard() *string
	SetProductCode(v string) *DocOcrMaxV2AdvanceRequest
	GetProductCode() *string
	SetSceneCode(v string) *DocOcrMaxV2AdvanceRequest
	GetSceneCode() *string
}

type DocOcrMaxV2AdvanceRequest struct {
	// Specifies whether to enable authoritative data source verification to enhance the ID document anti-counterfeiting capability. Valid values:
	//
	// - **T**: enabled.
	//
	// - **F**: disabled (default).
	//
	// >
	//
	// > - **Applicable ID document types**: Chinese resident identity card (CHN01001) and Chinese mainland driver\\"s license (CHN02001).
	//
	// > - **Data transmission statement**: Enabling this parameter indicates that you agree to transmit the user\\"s name and ID document number to an authoritative data source in the Chinese mainland for consistency verification.
	//
	// > - **Performance impact**: After this parameter is enabled, the API response time increases by about 1 to 2 seconds. Adjust the timeout settings accordingly.
	//
	// example:
	//
	// T
	Authorize *string `json:"Authorize,omitempty" xml:"Authorize,omitempty"`
	// The page to be recognized. Valid values:
	//
	// - 01: the portrait side of the ID document (default).
	//
	// - 02: the back side of the ID document.
	//
	// example:
	//
	// 01
	DocPage *string `json:"DocPage,omitempty" xml:"DocPage,omitempty"`
	// The type of the ID document.
	//
	// - Format: country code + ID document type abbreviation + page (optional).
	//
	// Note
	//
	// - If OcrModel is set to 0, DocType is required and specifies the ID document type. The existing logic remains unchanged.
	//
	// - If OcrModel is set to 1 or 2, DocType must be left empty.
	//
	// example:
	//
	// CHN01001
	DocType *string `json:"DocType,omitempty" xml:"DocType,omitempty"`
	// The Base64-encoded image of the ID document. If you use the IdOcrPictureBase64 parameter to pass the Base64-encoded photo of the ID document, check the photo size and do not upload excessively large photos.
	//
	// example:
	//
	// base64
	IdOcrPictureBase64 *string `json:"IdOcrPictureBase64,omitempty" xml:"IdOcrPictureBase64,omitempty"`
	// The file stream of the ID document image.
	//
	// example:
	//
	// InputStream
	IdOcrPictureFileObject io.Reader `json:"IdOcrPictureFile,omitempty" xml:"IdOcrPictureFile,omitempty"`
	// The URL of the ID document image. The URL must be a publicly accessible HTTP or HTTPS link.
	//
	// example:
	//
	// https://***********.oss-cn-hangzhou.aliyuncs.com/1669520556530-expo/default/face/20221127114236530_w3kx2e6t.jpg
	IdOcrPictureUrl *string `json:"IdOcrPictureUrl,omitempty" xml:"IdOcrPictureUrl,omitempty"`
	// Specifies whether to enable the ID document anti-counterfeiting feature. Valid values:
	//
	// - T: enabled.
	//
	// - F: disabled (default).
	//
	// example:
	//
	// F
	IdSpoof *string `json:"IdSpoof,omitempty" xml:"IdSpoof,omitempty"`
	// The custom threshold mode for OCR quality detection. Valid values:
	//
	// - 0: system default.
	//
	// - 1: strict mode.
	//
	// - 2: loose mode.
	//
	// - 3: disable quality detection (default).
	//
	// example:
	//
	// 0
	IdThreshold *string `json:"IdThreshold,omitempty" xml:"IdThreshold,omitempty"`
	// The custom unique business identifier on the merchant side, which is used for subsequent troubleshooting. The value can be up to 32 characters in length and can contain letters and digits. Make sure that the value is unique.
	//
	// example:
	//
	// e0c34a77f5ac40a5aa5e6ed20c353888
	MerchantBizId *string `json:"MerchantBizId,omitempty" xml:"MerchantBizId,omitempty"`
	// The custom user ID or other identifier that can identify a specific user, such as a mobile phone number or an email address. Desensitize the value of this field in advance, for example, by hashing the value.
	//
	// example:
	//
	// 123456789
	MerchantUserId *string `json:"MerchantUserId,omitempty" xml:"MerchantUserId,omitempty"`
	// The OCR recognition mode. Valid values:
	//
	// - 0: general ID document recognition mode (default).
	//
	// - 1: automatic ID document classification mode.
	//
	// - 2: automatic ID document classification and general recognition mode.
	//
	// example:
	//
	// 0
	OcrModel *string `json:"OcrModel,omitempty" xml:"OcrModel,omitempty"`
	// Specifies whether to enable standardization for key OCR fields. Valid values:
	//
	// - 0: disabled (default).
	//
	// - 1: enabled.
	//
	// example:
	//
	// 0
	OcrValueStandard *string `json:"OcrValueStandard,omitempty" xml:"OcrValueStandard,omitempty"`
	// The product solution to access. Valid value: ID_OCR_MAX.
	//
	// example:
	//
	// ID_OCR_MAX
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The custom authentication scene ID, which is used to query related records by entering this scene ID in the console later. The value can be up to 10 characters in length and can contain letters, digits, and underscores.
	//
	// example:
	//
	// 1234567890
	SceneCode *string `json:"SceneCode,omitempty" xml:"SceneCode,omitempty"`
}

func (s DocOcrMaxV2AdvanceRequest) String() string {
	return dara.Prettify(s)
}

func (s DocOcrMaxV2AdvanceRequest) GoString() string {
	return s.String()
}

func (s *DocOcrMaxV2AdvanceRequest) GetAuthorize() *string {
	return s.Authorize
}

func (s *DocOcrMaxV2AdvanceRequest) GetDocPage() *string {
	return s.DocPage
}

func (s *DocOcrMaxV2AdvanceRequest) GetDocType() *string {
	return s.DocType
}

func (s *DocOcrMaxV2AdvanceRequest) GetIdOcrPictureBase64() *string {
	return s.IdOcrPictureBase64
}

func (s *DocOcrMaxV2AdvanceRequest) GetIdOcrPictureFileObject() io.Reader {
	return s.IdOcrPictureFileObject
}

func (s *DocOcrMaxV2AdvanceRequest) GetIdOcrPictureUrl() *string {
	return s.IdOcrPictureUrl
}

func (s *DocOcrMaxV2AdvanceRequest) GetIdSpoof() *string {
	return s.IdSpoof
}

func (s *DocOcrMaxV2AdvanceRequest) GetIdThreshold() *string {
	return s.IdThreshold
}

func (s *DocOcrMaxV2AdvanceRequest) GetMerchantBizId() *string {
	return s.MerchantBizId
}

func (s *DocOcrMaxV2AdvanceRequest) GetMerchantUserId() *string {
	return s.MerchantUserId
}

func (s *DocOcrMaxV2AdvanceRequest) GetOcrModel() *string {
	return s.OcrModel
}

func (s *DocOcrMaxV2AdvanceRequest) GetOcrValueStandard() *string {
	return s.OcrValueStandard
}

func (s *DocOcrMaxV2AdvanceRequest) GetProductCode() *string {
	return s.ProductCode
}

func (s *DocOcrMaxV2AdvanceRequest) GetSceneCode() *string {
	return s.SceneCode
}

func (s *DocOcrMaxV2AdvanceRequest) SetAuthorize(v string) *DocOcrMaxV2AdvanceRequest {
	s.Authorize = &v
	return s
}

func (s *DocOcrMaxV2AdvanceRequest) SetDocPage(v string) *DocOcrMaxV2AdvanceRequest {
	s.DocPage = &v
	return s
}

func (s *DocOcrMaxV2AdvanceRequest) SetDocType(v string) *DocOcrMaxV2AdvanceRequest {
	s.DocType = &v
	return s
}

func (s *DocOcrMaxV2AdvanceRequest) SetIdOcrPictureBase64(v string) *DocOcrMaxV2AdvanceRequest {
	s.IdOcrPictureBase64 = &v
	return s
}

func (s *DocOcrMaxV2AdvanceRequest) SetIdOcrPictureFileObject(v io.Reader) *DocOcrMaxV2AdvanceRequest {
	s.IdOcrPictureFileObject = v
	return s
}

func (s *DocOcrMaxV2AdvanceRequest) SetIdOcrPictureUrl(v string) *DocOcrMaxV2AdvanceRequest {
	s.IdOcrPictureUrl = &v
	return s
}

func (s *DocOcrMaxV2AdvanceRequest) SetIdSpoof(v string) *DocOcrMaxV2AdvanceRequest {
	s.IdSpoof = &v
	return s
}

func (s *DocOcrMaxV2AdvanceRequest) SetIdThreshold(v string) *DocOcrMaxV2AdvanceRequest {
	s.IdThreshold = &v
	return s
}

func (s *DocOcrMaxV2AdvanceRequest) SetMerchantBizId(v string) *DocOcrMaxV2AdvanceRequest {
	s.MerchantBizId = &v
	return s
}

func (s *DocOcrMaxV2AdvanceRequest) SetMerchantUserId(v string) *DocOcrMaxV2AdvanceRequest {
	s.MerchantUserId = &v
	return s
}

func (s *DocOcrMaxV2AdvanceRequest) SetOcrModel(v string) *DocOcrMaxV2AdvanceRequest {
	s.OcrModel = &v
	return s
}

func (s *DocOcrMaxV2AdvanceRequest) SetOcrValueStandard(v string) *DocOcrMaxV2AdvanceRequest {
	s.OcrValueStandard = &v
	return s
}

func (s *DocOcrMaxV2AdvanceRequest) SetProductCode(v string) *DocOcrMaxV2AdvanceRequest {
	s.ProductCode = &v
	return s
}

func (s *DocOcrMaxV2AdvanceRequest) SetSceneCode(v string) *DocOcrMaxV2AdvanceRequest {
	s.SceneCode = &v
	return s
}

func (s *DocOcrMaxV2AdvanceRequest) Validate() error {
	return dara.Validate(s)
}
