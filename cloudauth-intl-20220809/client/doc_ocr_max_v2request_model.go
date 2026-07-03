// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDocOcrMaxV2Request interface {
	dara.Model
	String() string
	GoString() string
	SetAuthorize(v string) *DocOcrMaxV2Request
	GetAuthorize() *string
	SetDocPage(v string) *DocOcrMaxV2Request
	GetDocPage() *string
	SetDocType(v string) *DocOcrMaxV2Request
	GetDocType() *string
	SetIdOcrPictureBase64(v string) *DocOcrMaxV2Request
	GetIdOcrPictureBase64() *string
	SetIdOcrPictureFile(v string) *DocOcrMaxV2Request
	GetIdOcrPictureFile() *string
	SetIdOcrPictureUrl(v string) *DocOcrMaxV2Request
	GetIdOcrPictureUrl() *string
	SetIdSpoof(v string) *DocOcrMaxV2Request
	GetIdSpoof() *string
	SetIdThreshold(v string) *DocOcrMaxV2Request
	GetIdThreshold() *string
	SetMerchantBizId(v string) *DocOcrMaxV2Request
	GetMerchantBizId() *string
	SetMerchantUserId(v string) *DocOcrMaxV2Request
	GetMerchantUserId() *string
	SetOcrModel(v string) *DocOcrMaxV2Request
	GetOcrModel() *string
	SetOcrValueStandard(v string) *DocOcrMaxV2Request
	GetOcrValueStandard() *string
	SetProductCode(v string) *DocOcrMaxV2Request
	GetProductCode() *string
	SetSceneCode(v string) *DocOcrMaxV2Request
	GetSceneCode() *string
}

type DocOcrMaxV2Request struct {
	// Specifies whether to enable authoritative data source verification to enhance document anti-forgery capabilities. Valid values:
	//
	// - **T**: enabled.
	//
	// - **F*	- (default): disabled.
	//
	// >
	//
	// > - **Applicable document types**: China resident identity card (CHN01001) and China mainland driver\\"s license (CHN02001).
	//
	// > - **Data transmission statement**: Enabling this parameter indicates consent to transmit the user\\"s name and document number to an authoritative data source in the Chinese mainland for consistency verification.
	//
	// > - **Performance impact**: Enabling this parameter increases the API response time by approximately 1 to 2 seconds. Adjust the timeout settings accordingly.
	//
	// example:
	//
	// T
	Authorize *string `json:"Authorize,omitempty" xml:"Authorize,omitempty"`
	// The expected page to recognize. Valid values:
	//
	// - 01 (default): the portrait side of the document.
	//
	// - 02: the back side of the document.
	//
	// example:
	//
	// 01
	DocPage *string `json:"DocPage,omitempty" xml:"DocPage,omitempty"`
	// The document type.
	//
	// - Format: country code + document type abbreviation + page (optional).
	//
	// Note:
	//
	// - OcrModel = 0: DocType is required. Specify the document type. The existing logic remains unchanged.
	//
	// - OcrModel = 1 or 2: DocType must be left empty.
	//
	// example:
	//
	// CHN01001
	DocType *string `json:"DocType,omitempty" xml:"DocType,omitempty"`
	// The Base64-encoded image of the identity document.
	//
	// If you use IdOcrPictureBase64 to pass in the document image, check the image size and do not pass in an excessively large image.
	//
	// example:
	//
	// base64
	IdOcrPictureBase64 *string `json:"IdOcrPictureBase64,omitempty" xml:"IdOcrPictureBase64,omitempty"`
	// The file stream of the document image.
	//
	// example:
	//
	// InputStream
	IdOcrPictureFile *string `json:"IdOcrPictureFile,omitempty" xml:"IdOcrPictureFile,omitempty"`
	// The URL of the identity document image. The URL must be a publicly accessible HTTP or HTTPS link.
	//
	// example:
	//
	// https://***********.oss-cn-hangzhou.aliyuncs.com/1669520556530-expo/default/face/20221127114236530_w3kx2e6t.jpg
	IdOcrPictureUrl *string `json:"IdOcrPictureUrl,omitempty" xml:"IdOcrPictureUrl,omitempty"`
	// Specifies whether to enable the document anti-forgery feature. Valid values:
	//
	// - T: enabled.
	//
	// - F (default): disabled.
	//
	// example:
	//
	// F
	IdSpoof *string `json:"IdSpoof,omitempty" xml:"IdSpoof,omitempty"`
	// The custom OCR quality detection threshold mode. Valid values:
	//
	// - 0: system default.
	//
	// - 1: strict mode.
	//
	// - 2: loose mode.
	//
	// - 3 (default): quality detection disabled.
	//
	// example:
	//
	// 0
	IdThreshold *string `json:"IdThreshold,omitempty" xml:"IdThreshold,omitempty"`
	// The merchant-defined unique business identifier, used for subsequent troubleshooting. The value can contain letters and digits, with a maximum length of 32 characters. Make sure the value is unique.
	//
	// example:
	//
	// e0c34a77f5ac40a5aa5e6ed20c353888
	MerchantBizId *string `json:"MerchantBizId,omitempty" xml:"MerchantBizId,omitempty"`
	// The custom user ID, or another identifier that can identify a specific user, such as a phone number or email address.
	//
	// Hash or otherwise desensitize this field value before passing it in.
	//
	// example:
	//
	// 123456789
	MerchantUserId *string `json:"MerchantUserId,omitempty" xml:"MerchantUserId,omitempty"`
	// The OCR recognition mode. Valid values:
	//
	// - 0: general document recognition mode (default).
	//
	// - 1: automatic document classification mode.
	//
	// - 2: automatic document classification and general recognition mode.
	//
	// example:
	//
	// 0
	OcrModel *string `json:"OcrModel,omitempty" xml:"OcrModel,omitempty"`
	// Specifies whether to enable OCR key field standardization. Valid values:
	//
	// - 0 (default): disabled.
	//
	// - 1: enabled.
	//
	// example:
	//
	// 0
	OcrValueStandard *string `json:"OcrValueStandard,omitempty" xml:"OcrValueStandard,omitempty"`
	// The product solution to use.
	//
	// Set this parameter to ID_OCR_MAX.
	//
	// example:
	//
	// ID_OCR_MAX
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The custom verification scenario ID. You can use this scenario ID to query related records in the console.
	//
	// The value can contain letters, digits, and underscores, with a maximum length of 10 characters.
	//
	// example:
	//
	// 1234567890
	SceneCode *string `json:"SceneCode,omitempty" xml:"SceneCode,omitempty"`
}

func (s DocOcrMaxV2Request) String() string {
	return dara.Prettify(s)
}

func (s DocOcrMaxV2Request) GoString() string {
	return s.String()
}

func (s *DocOcrMaxV2Request) GetAuthorize() *string {
	return s.Authorize
}

func (s *DocOcrMaxV2Request) GetDocPage() *string {
	return s.DocPage
}

func (s *DocOcrMaxV2Request) GetDocType() *string {
	return s.DocType
}

func (s *DocOcrMaxV2Request) GetIdOcrPictureBase64() *string {
	return s.IdOcrPictureBase64
}

func (s *DocOcrMaxV2Request) GetIdOcrPictureFile() *string {
	return s.IdOcrPictureFile
}

func (s *DocOcrMaxV2Request) GetIdOcrPictureUrl() *string {
	return s.IdOcrPictureUrl
}

func (s *DocOcrMaxV2Request) GetIdSpoof() *string {
	return s.IdSpoof
}

func (s *DocOcrMaxV2Request) GetIdThreshold() *string {
	return s.IdThreshold
}

func (s *DocOcrMaxV2Request) GetMerchantBizId() *string {
	return s.MerchantBizId
}

func (s *DocOcrMaxV2Request) GetMerchantUserId() *string {
	return s.MerchantUserId
}

func (s *DocOcrMaxV2Request) GetOcrModel() *string {
	return s.OcrModel
}

func (s *DocOcrMaxV2Request) GetOcrValueStandard() *string {
	return s.OcrValueStandard
}

func (s *DocOcrMaxV2Request) GetProductCode() *string {
	return s.ProductCode
}

func (s *DocOcrMaxV2Request) GetSceneCode() *string {
	return s.SceneCode
}

func (s *DocOcrMaxV2Request) SetAuthorize(v string) *DocOcrMaxV2Request {
	s.Authorize = &v
	return s
}

func (s *DocOcrMaxV2Request) SetDocPage(v string) *DocOcrMaxV2Request {
	s.DocPage = &v
	return s
}

func (s *DocOcrMaxV2Request) SetDocType(v string) *DocOcrMaxV2Request {
	s.DocType = &v
	return s
}

func (s *DocOcrMaxV2Request) SetIdOcrPictureBase64(v string) *DocOcrMaxV2Request {
	s.IdOcrPictureBase64 = &v
	return s
}

func (s *DocOcrMaxV2Request) SetIdOcrPictureFile(v string) *DocOcrMaxV2Request {
	s.IdOcrPictureFile = &v
	return s
}

func (s *DocOcrMaxV2Request) SetIdOcrPictureUrl(v string) *DocOcrMaxV2Request {
	s.IdOcrPictureUrl = &v
	return s
}

func (s *DocOcrMaxV2Request) SetIdSpoof(v string) *DocOcrMaxV2Request {
	s.IdSpoof = &v
	return s
}

func (s *DocOcrMaxV2Request) SetIdThreshold(v string) *DocOcrMaxV2Request {
	s.IdThreshold = &v
	return s
}

func (s *DocOcrMaxV2Request) SetMerchantBizId(v string) *DocOcrMaxV2Request {
	s.MerchantBizId = &v
	return s
}

func (s *DocOcrMaxV2Request) SetMerchantUserId(v string) *DocOcrMaxV2Request {
	s.MerchantUserId = &v
	return s
}

func (s *DocOcrMaxV2Request) SetOcrModel(v string) *DocOcrMaxV2Request {
	s.OcrModel = &v
	return s
}

func (s *DocOcrMaxV2Request) SetOcrValueStandard(v string) *DocOcrMaxV2Request {
	s.OcrValueStandard = &v
	return s
}

func (s *DocOcrMaxV2Request) SetProductCode(v string) *DocOcrMaxV2Request {
	s.ProductCode = &v
	return s
}

func (s *DocOcrMaxV2Request) SetSceneCode(v string) *DocOcrMaxV2Request {
	s.SceneCode = &v
	return s
}

func (s *DocOcrMaxV2Request) Validate() error {
	return dara.Validate(s)
}
