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
	// CardSide
	//
	// example:
	//
	// 0
	CardSide *string `json:"CardSide,omitempty" xml:"CardSide,omitempty"`
	// Document type
	//
	// example:
	//
	// 00000006
	DocType *string `json:"DocType,omitempty" xml:"DocType,omitempty"`
	// Whether to perform ID face quality detection
	//
	// - T: Indicates that detection is required
	//
	// - F: Indicates that detection is not required (default F)
	//
	// example:
	//
	// F
	IdFaceQuality *string `json:"IdFaceQuality,omitempty" xml:"IdFaceQuality,omitempty"`
	// Base64 of the front side of the document image
	//
	// example:
	//
	// base64
	IdOcrPictureBase64 *string `json:"IdOcrPictureBase64,omitempty" xml:"IdOcrPictureBase64,omitempty"`
	// URL of the front side of the document image
	//
	// example:
	//
	// https://digital-cardocr-prod8.oss-cn-hangzhou.aliyuncs.com/1669520556530-expo/default/face/20221127114236530_w3kx2e6t.jpg
	IdOcrPictureUrl *string `json:"IdOcrPictureUrl,omitempty" xml:"IdOcrPictureUrl,omitempty"`
	// IdThreshold
	//
	// example:
	//
	// 0
	IdThreshold *string `json:"IdThreshold,omitempty" xml:"IdThreshold,omitempty"`
	// A unique business identifier defined by the merchant, used for subsequent troubleshooting. It supports a combination of letters and numbers, with a maximum length of 32 characters. Please ensure uniqueness.
	//
	// example:
	//
	// dso9322***dsjsd22
	MerchantBizId *string `json:"MerchantBizId,omitempty" xml:"MerchantBizId,omitempty"`
	// A custom user ID in the business, please keep it unique.
	//
	// example:
	//
	// 123456789
	MerchantUserId *string `json:"MerchantUserId,omitempty" xml:"MerchantUserId,omitempty"`
	// Whether to perform document OCR
	//
	// - T: Indicates that document OCR is required
	//
	// - F: Indicates that document OCR is not required
	//
	// example:
	//
	// T
	Ocr *string `json:"Ocr,omitempty" xml:"Ocr,omitempty"`
	// Product code
	//
	// example:
	//
	// 产品方案类型ID_OCR_MIN
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// Whether to enable anti-counterfeiting detection
	//
	// - T: Indicates that anti-counterfeiting is enabled
	//
	// - F: Indicates that anti-counterfeiting is disabled
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
