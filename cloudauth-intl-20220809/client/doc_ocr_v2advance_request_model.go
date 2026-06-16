// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
	"io"
)

type iDocOcrV2AdvanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCardSide(v string) *DocOcrV2AdvanceRequest
	GetCardSide() *string
	SetDocType(v string) *DocOcrV2AdvanceRequest
	GetDocType() *string
	SetIdFaceQuality(v string) *DocOcrV2AdvanceRequest
	GetIdFaceQuality() *string
	SetIdOcrPictureBase64(v string) *DocOcrV2AdvanceRequest
	GetIdOcrPictureBase64() *string
	SetIdOcrPictureFileObject(v io.Reader) *DocOcrV2AdvanceRequest
	GetIdOcrPictureFileObject() io.Reader
	SetIdOcrPictureUrl(v string) *DocOcrV2AdvanceRequest
	GetIdOcrPictureUrl() *string
	SetIdThreshold(v string) *DocOcrV2AdvanceRequest
	GetIdThreshold() *string
	SetMerchantBizId(v string) *DocOcrV2AdvanceRequest
	GetMerchantBizId() *string
	SetMerchantUserId(v string) *DocOcrV2AdvanceRequest
	GetMerchantUserId() *string
	SetOcr(v string) *DocOcrV2AdvanceRequest
	GetOcr() *string
	SetProductCode(v string) *DocOcrV2AdvanceRequest
	GetProductCode() *string
	SetSpoof(v string) *DocOcrV2AdvanceRequest
	GetSpoof() *string
}

type DocOcrV2AdvanceRequest struct {
	// Specifies the side of the certificate. If this parameter is not specified, the portrait side is used by default.
	//
	// - OCR_ID_FACE (default): the portrait side
	//
	// - OCR_ID_NATIONAL_EMBLEM: the national emblem side.
	//
	// example:
	//
	// OCR_ID_FACE
	CardSide *string `json:"CardSide,omitempty" xml:"CardSide,omitempty"`
	// The certificate type.
	//
	// example:
	//
	// 00000006
	DocType *string `json:"DocType,omitempty" xml:"DocType,omitempty"`
	// Specifies whether to perform face quality detection on the certificate.
	//
	// - T: Detection is required.
	//
	// - F: Detection is not required. (Default: F).
	//
	// example:
	//
	// F
	IdFaceQuality *string `json:"IdFaceQuality,omitempty" xml:"IdFaceQuality,omitempty"`
	// The Base64-encoded card or certificate image.
	//
	// If you use IdOcrPictureBase64 to pass in the certificate image, check the image size and do not pass in an excessively large image.
	//
	// example:
	//
	// base64
	IdOcrPictureBase64 *string `json:"IdOcrPictureBase64,omitempty" xml:"IdOcrPictureBase64,omitempty"`
	// The file stream of the front side of the certificate image.
	//
	// example:
	//
	// InputStream
	IdOcrPictureFileObject io.Reader `json:"IdOcrPictureFile,omitempty" xml:"IdOcrPictureFile,omitempty"`
	// The URL of the front side of the certificate image.
	//
	// example:
	//
	// https://digital-cardocr-prod8.oss-cn-hangzhou.aliyuncs.com/1669520556530-expo/default/face/20221127114236530_w3kx2e6t.jpg
	IdOcrPictureUrl *string `json:"IdOcrPictureUrl,omitempty" xml:"IdOcrPictureUrl,omitempty"`
	// The custom OCR quality detection threshold mode. Valid values:
	//
	// - 0: system default
	//
	// - 1: strict mode
	//
	// - 2: loose mode
	//
	// - 3 (default): quality detection is disabled.
	//
	// example:
	//
	// 0
	IdThreshold *string `json:"IdThreshold,omitempty" xml:"IdThreshold,omitempty"`
	// The merchant-defined unique business ID used for subsequent troubleshooting. The value can contain letters and digits with a maximum length of 32 characters. Ensure that the value is unique.
	//
	// example:
	//
	// dso9322***dsjsd22
	MerchantBizId *string `json:"MerchantBizId,omitempty" xml:"MerchantBizId,omitempty"`
	// The custom user ID defined in your business. Ensure that the value is unique.
	//
	// example:
	//
	// 123456789
	MerchantUserId *string `json:"MerchantUserId,omitempty" xml:"MerchantUserId,omitempty"`
	// Specifies whether to perform certificate OCR.
	//
	// - T: OCR is required.
	//
	// - F: OCR is not required.
	//
	// example:
	//
	// T
	Ocr *string `json:"Ocr,omitempty" xml:"Ocr,omitempty"`
	// The product code.
	//
	// example:
	//
	// 产品方案类型ID_OCR_MIN
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// Specifies whether to enable anti-spoofing detection.
	//
	// - T: Anti-spoofing is enabled.
	//
	// - F: Anti-spoofing is disabled.
	//
	// example:
	//
	// F
	Spoof *string `json:"Spoof,omitempty" xml:"Spoof,omitempty"`
}

func (s DocOcrV2AdvanceRequest) String() string {
	return dara.Prettify(s)
}

func (s DocOcrV2AdvanceRequest) GoString() string {
	return s.String()
}

func (s *DocOcrV2AdvanceRequest) GetCardSide() *string {
	return s.CardSide
}

func (s *DocOcrV2AdvanceRequest) GetDocType() *string {
	return s.DocType
}

func (s *DocOcrV2AdvanceRequest) GetIdFaceQuality() *string {
	return s.IdFaceQuality
}

func (s *DocOcrV2AdvanceRequest) GetIdOcrPictureBase64() *string {
	return s.IdOcrPictureBase64
}

func (s *DocOcrV2AdvanceRequest) GetIdOcrPictureFileObject() io.Reader {
	return s.IdOcrPictureFileObject
}

func (s *DocOcrV2AdvanceRequest) GetIdOcrPictureUrl() *string {
	return s.IdOcrPictureUrl
}

func (s *DocOcrV2AdvanceRequest) GetIdThreshold() *string {
	return s.IdThreshold
}

func (s *DocOcrV2AdvanceRequest) GetMerchantBizId() *string {
	return s.MerchantBizId
}

func (s *DocOcrV2AdvanceRequest) GetMerchantUserId() *string {
	return s.MerchantUserId
}

func (s *DocOcrV2AdvanceRequest) GetOcr() *string {
	return s.Ocr
}

func (s *DocOcrV2AdvanceRequest) GetProductCode() *string {
	return s.ProductCode
}

func (s *DocOcrV2AdvanceRequest) GetSpoof() *string {
	return s.Spoof
}

func (s *DocOcrV2AdvanceRequest) SetCardSide(v string) *DocOcrV2AdvanceRequest {
	s.CardSide = &v
	return s
}

func (s *DocOcrV2AdvanceRequest) SetDocType(v string) *DocOcrV2AdvanceRequest {
	s.DocType = &v
	return s
}

func (s *DocOcrV2AdvanceRequest) SetIdFaceQuality(v string) *DocOcrV2AdvanceRequest {
	s.IdFaceQuality = &v
	return s
}

func (s *DocOcrV2AdvanceRequest) SetIdOcrPictureBase64(v string) *DocOcrV2AdvanceRequest {
	s.IdOcrPictureBase64 = &v
	return s
}

func (s *DocOcrV2AdvanceRequest) SetIdOcrPictureFileObject(v io.Reader) *DocOcrV2AdvanceRequest {
	s.IdOcrPictureFileObject = v
	return s
}

func (s *DocOcrV2AdvanceRequest) SetIdOcrPictureUrl(v string) *DocOcrV2AdvanceRequest {
	s.IdOcrPictureUrl = &v
	return s
}

func (s *DocOcrV2AdvanceRequest) SetIdThreshold(v string) *DocOcrV2AdvanceRequest {
	s.IdThreshold = &v
	return s
}

func (s *DocOcrV2AdvanceRequest) SetMerchantBizId(v string) *DocOcrV2AdvanceRequest {
	s.MerchantBizId = &v
	return s
}

func (s *DocOcrV2AdvanceRequest) SetMerchantUserId(v string) *DocOcrV2AdvanceRequest {
	s.MerchantUserId = &v
	return s
}

func (s *DocOcrV2AdvanceRequest) SetOcr(v string) *DocOcrV2AdvanceRequest {
	s.Ocr = &v
	return s
}

func (s *DocOcrV2AdvanceRequest) SetProductCode(v string) *DocOcrV2AdvanceRequest {
	s.ProductCode = &v
	return s
}

func (s *DocOcrV2AdvanceRequest) SetSpoof(v string) *DocOcrV2AdvanceRequest {
	s.Spoof = &v
	return s
}

func (s *DocOcrV2AdvanceRequest) Validate() error {
	return dara.Validate(s)
}
