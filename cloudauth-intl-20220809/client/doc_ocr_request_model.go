// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDocOcrRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCardSide(v string) *DocOcrRequest
	GetCardSide() *string
	SetDocType(v string) *DocOcrRequest
	GetDocType() *string
	SetIdFaceQuality(v string) *DocOcrRequest
	GetIdFaceQuality() *string
	SetIdOcrPictureBase64(v string) *DocOcrRequest
	GetIdOcrPictureBase64() *string
	SetIdOcrPictureUrl(v string) *DocOcrRequest
	GetIdOcrPictureUrl() *string
	SetIdThreshold(v string) *DocOcrRequest
	GetIdThreshold() *string
	SetMerchantBizId(v string) *DocOcrRequest
	GetMerchantBizId() *string
	SetMerchantUserId(v string) *DocOcrRequest
	GetMerchantUserId() *string
	SetOcr(v string) *DocOcrRequest
	GetOcr() *string
	SetProductCode(v string) *DocOcrRequest
	GetProductCode() *string
	SetSpoof(v string) *DocOcrRequest
	GetSpoof() *string
}

type DocOcrRequest struct {
	// Specifies the side of the certificate to distinguish between the national emblem side and the portrait side. If this parameter is not specified, the portrait side is used by default. Valid values:
	//
	// - OCR_ID_FACE (default): portrait side
	//
	// - OCR_ID_NATIONAL_EMBLEM: national emblem side.
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
	// Specifies whether to perform certificate face quality detection. Valid values:
	//
	// - T: Detection is required.
	//
	// - F: Detection is not required. (Default value: F).
	//
	// example:
	//
	// F
	IdFaceQuality *string `json:"IdFaceQuality,omitempty" xml:"IdFaceQuality,omitempty"`
	// The Base64-encoded card or certificate image.
	//
	// If you use IdOcrPictureBase64 (Base64-encoded photo) to submit the certificate photo, check the photo size and do not submit an excessively large photo.
	//
	// example:
	//
	// base64
	IdOcrPictureBase64 *string `json:"IdOcrPictureBase64,omitempty" xml:"IdOcrPictureBase64,omitempty"`
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
	// - 3 (default): quality detection disabled.
	//
	// example:
	//
	// 0
	IdThreshold *string `json:"IdThreshold,omitempty" xml:"IdThreshold,omitempty"`
	// The custom business unique identifier on the merchant side, used for subsequent issue tracking and troubleshooting. The value can be a combination of letters and digits with a maximum length of 32 characters. Ensure that the value is unique.
	//
	// example:
	//
	// dso9322***dsjsd22
	MerchantBizId *string `json:"MerchantBizId,omitempty" xml:"MerchantBizId,omitempty"`
	// The custom user ID in the business. Ensure that the value is unique.
	//
	// example:
	//
	// 123456789
	MerchantUserId *string `json:"MerchantUserId,omitempty" xml:"MerchantUserId,omitempty"`
	// Specifies whether to perform certificate OCR. Valid values:
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
	// Specifies whether to enable anti-spoofing detection. Valid values:
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

func (s DocOcrRequest) String() string {
	return dara.Prettify(s)
}

func (s DocOcrRequest) GoString() string {
	return s.String()
}

func (s *DocOcrRequest) GetCardSide() *string {
	return s.CardSide
}

func (s *DocOcrRequest) GetDocType() *string {
	return s.DocType
}

func (s *DocOcrRequest) GetIdFaceQuality() *string {
	return s.IdFaceQuality
}

func (s *DocOcrRequest) GetIdOcrPictureBase64() *string {
	return s.IdOcrPictureBase64
}

func (s *DocOcrRequest) GetIdOcrPictureUrl() *string {
	return s.IdOcrPictureUrl
}

func (s *DocOcrRequest) GetIdThreshold() *string {
	return s.IdThreshold
}

func (s *DocOcrRequest) GetMerchantBizId() *string {
	return s.MerchantBizId
}

func (s *DocOcrRequest) GetMerchantUserId() *string {
	return s.MerchantUserId
}

func (s *DocOcrRequest) GetOcr() *string {
	return s.Ocr
}

func (s *DocOcrRequest) GetProductCode() *string {
	return s.ProductCode
}

func (s *DocOcrRequest) GetSpoof() *string {
	return s.Spoof
}

func (s *DocOcrRequest) SetCardSide(v string) *DocOcrRequest {
	s.CardSide = &v
	return s
}

func (s *DocOcrRequest) SetDocType(v string) *DocOcrRequest {
	s.DocType = &v
	return s
}

func (s *DocOcrRequest) SetIdFaceQuality(v string) *DocOcrRequest {
	s.IdFaceQuality = &v
	return s
}

func (s *DocOcrRequest) SetIdOcrPictureBase64(v string) *DocOcrRequest {
	s.IdOcrPictureBase64 = &v
	return s
}

func (s *DocOcrRequest) SetIdOcrPictureUrl(v string) *DocOcrRequest {
	s.IdOcrPictureUrl = &v
	return s
}

func (s *DocOcrRequest) SetIdThreshold(v string) *DocOcrRequest {
	s.IdThreshold = &v
	return s
}

func (s *DocOcrRequest) SetMerchantBizId(v string) *DocOcrRequest {
	s.MerchantBizId = &v
	return s
}

func (s *DocOcrRequest) SetMerchantUserId(v string) *DocOcrRequest {
	s.MerchantUserId = &v
	return s
}

func (s *DocOcrRequest) SetOcr(v string) *DocOcrRequest {
	s.Ocr = &v
	return s
}

func (s *DocOcrRequest) SetProductCode(v string) *DocOcrRequest {
	s.ProductCode = &v
	return s
}

func (s *DocOcrRequest) SetSpoof(v string) *DocOcrRequest {
	s.Spoof = &v
	return s
}

func (s *DocOcrRequest) Validate() error {
	return dara.Validate(s)
}
