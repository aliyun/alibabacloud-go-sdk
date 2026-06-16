// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCardOcrRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDocType(v string) *CardOcrRequest
	GetDocType() *string
	SetIdFaceQuality(v string) *CardOcrRequest
	GetIdFaceQuality() *string
	SetIdOcrPictureBase64(v string) *CardOcrRequest
	GetIdOcrPictureBase64() *string
	SetIdOcrPictureUrl(v string) *CardOcrRequest
	GetIdOcrPictureUrl() *string
	SetMerchantBizId(v string) *CardOcrRequest
	GetMerchantBizId() *string
	SetMerchantUserId(v string) *CardOcrRequest
	GetMerchantUserId() *string
	SetOcr(v string) *CardOcrRequest
	GetOcr() *string
	SetProductCode(v string) *CardOcrRequest
	GetProductCode() *string
	SetSpoof(v string) *CardOcrRequest
	GetSpoof() *string
}

type CardOcrRequest struct {
	// The document type.
	//
	// example:
	//
	// 00000006
	DocType *string `json:"DocType,omitempty" xml:"DocType,omitempty"`
	// Specifies whether to perform face quality detection on the ID document. Valid values:
	//
	// - T: Perform face quality detection.
	//
	// - F: Do not perform face quality detection. This is the default value.
	//
	// example:
	//
	// F
	IdFaceQuality *string `json:"IdFaceQuality,omitempty" xml:"IdFaceQuality,omitempty"`
	// The Base64-encoded image of the front side of the ID document. The value is a Base64 encoding of the image.
	//
	// example:
	//
	// base64
	IdOcrPictureBase64 *string `json:"IdOcrPictureBase64,omitempty" xml:"IdOcrPictureBase64,omitempty"`
	// The URL of the front-side image of the ID document.
	//
	// example:
	//
	// https://digital-cardocr-prod8.oss-cn-hangzhou.aliyuncs.com/1669520556530-expo/default/face/20221127114236530_w3kx2e6t.jpg
	IdOcrPictureUrl *string `json:"IdOcrPictureUrl,omitempty" xml:"IdOcrPictureUrl,omitempty"`
	// The merchant-defined unique business ID used for subsequent troubleshooting. The value can be a combination of letters and digits with a maximum length of 32 characters. Ensure that the value is unique.
	//
	// example:
	//
	// dso9322***dsjsd22
	MerchantBizId *string `json:"MerchantBizId,omitempty" xml:"MerchantBizId,omitempty"`
	// The merchant user ID or another identifier that can be used to identify a specific user, such as a phone number or email address. We strongly recommend that you pre-desensitize the value of the userId field, for example, by hashing the value.
	//
	// example:
	//
	// 123456789
	MerchantUserId *string `json:"MerchantUserId,omitempty" xml:"MerchantUserId,omitempty"`
	// Specifies whether to perform OCR on the ID document. Valid values:
	//
	// - T: Perform OCR on the ID document. This is the default value.
	//
	// - F: Do not perform OCR.
	//
	// example:
	//
	// T
	Ocr *string `json:"Ocr,omitempty" xml:"Ocr,omitempty"`
	// The product code.
	//
	// example:
	//
	// ID_OCR_MIN
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// Specifies whether to enable anti-spoofing detection. Valid values:
	//
	// - T: Enable anti-spoofing detection.
	//
	// - F: Disable anti-spoofing detection. This is the default value.
	//
	// example:
	//
	// F
	Spoof *string `json:"Spoof,omitempty" xml:"Spoof,omitempty"`
}

func (s CardOcrRequest) String() string {
	return dara.Prettify(s)
}

func (s CardOcrRequest) GoString() string {
	return s.String()
}

func (s *CardOcrRequest) GetDocType() *string {
	return s.DocType
}

func (s *CardOcrRequest) GetIdFaceQuality() *string {
	return s.IdFaceQuality
}

func (s *CardOcrRequest) GetIdOcrPictureBase64() *string {
	return s.IdOcrPictureBase64
}

func (s *CardOcrRequest) GetIdOcrPictureUrl() *string {
	return s.IdOcrPictureUrl
}

func (s *CardOcrRequest) GetMerchantBizId() *string {
	return s.MerchantBizId
}

func (s *CardOcrRequest) GetMerchantUserId() *string {
	return s.MerchantUserId
}

func (s *CardOcrRequest) GetOcr() *string {
	return s.Ocr
}

func (s *CardOcrRequest) GetProductCode() *string {
	return s.ProductCode
}

func (s *CardOcrRequest) GetSpoof() *string {
	return s.Spoof
}

func (s *CardOcrRequest) SetDocType(v string) *CardOcrRequest {
	s.DocType = &v
	return s
}

func (s *CardOcrRequest) SetIdFaceQuality(v string) *CardOcrRequest {
	s.IdFaceQuality = &v
	return s
}

func (s *CardOcrRequest) SetIdOcrPictureBase64(v string) *CardOcrRequest {
	s.IdOcrPictureBase64 = &v
	return s
}

func (s *CardOcrRequest) SetIdOcrPictureUrl(v string) *CardOcrRequest {
	s.IdOcrPictureUrl = &v
	return s
}

func (s *CardOcrRequest) SetMerchantBizId(v string) *CardOcrRequest {
	s.MerchantBizId = &v
	return s
}

func (s *CardOcrRequest) SetMerchantUserId(v string) *CardOcrRequest {
	s.MerchantUserId = &v
	return s
}

func (s *CardOcrRequest) SetOcr(v string) *CardOcrRequest {
	s.Ocr = &v
	return s
}

func (s *CardOcrRequest) SetProductCode(v string) *CardOcrRequest {
	s.ProductCode = &v
	return s
}

func (s *CardOcrRequest) SetSpoof(v string) *CardOcrRequest {
	s.Spoof = &v
	return s
}

func (s *CardOcrRequest) Validate() error {
	return dara.Validate(s)
}
