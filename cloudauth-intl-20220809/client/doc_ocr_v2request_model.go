// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDocOcrV2Request interface {
	dara.Model
	String() string
	GoString() string
	SetCardSide(v string) *DocOcrV2Request
	GetCardSide() *string
	SetDocType(v string) *DocOcrV2Request
	GetDocType() *string
	SetIdFaceQuality(v string) *DocOcrV2Request
	GetIdFaceQuality() *string
	SetIdOcrPictureBase64(v string) *DocOcrV2Request
	GetIdOcrPictureBase64() *string
	SetIdOcrPictureFile(v string) *DocOcrV2Request
	GetIdOcrPictureFile() *string
	SetIdOcrPictureUrl(v string) *DocOcrV2Request
	GetIdOcrPictureUrl() *string
	SetIdThreshold(v string) *DocOcrV2Request
	GetIdThreshold() *string
	SetMerchantBizId(v string) *DocOcrV2Request
	GetMerchantBizId() *string
	SetMerchantUserId(v string) *DocOcrV2Request
	GetMerchantUserId() *string
	SetOcr(v string) *DocOcrV2Request
	GetOcr() *string
	SetProductCode(v string) *DocOcrV2Request
	GetProductCode() *string
	SetSpoof(v string) *DocOcrV2Request
	GetSpoof() *string
}

type DocOcrV2Request struct {
	// Specifies the side of the ID card, which can be the national emblem side or the portrait side. If you do not specify this parameter, the portrait side is used by default. Valid values:
	//
	// - OCR_ID_FACE: the portrait side. Default value.
	//
	// - OCR_ID_NATIONAL_EMBLEM: the national emblem side.
	//
	// example:
	//
	// OCR_ID_FACE
	CardSide *string `json:"CardSide,omitempty" xml:"CardSide,omitempty"`
	// The type of the ID card.
	//
	// example:
	//
	// 00000006
	DocType *string `json:"DocType,omitempty" xml:"DocType,omitempty"`
	// Specifies whether to perform face quality check for the ID card. Valid values:
	//
	// - T: Face quality check is performed.
	//
	// - F: Face quality check is not performed. Default value: F.
	//
	// example:
	//
	// F
	IdFaceQuality *string `json:"IdFaceQuality,omitempty" xml:"IdFaceQuality,omitempty"`
	// The Base64-encoded image of the ID card.
	//
	// If you choose to pass the ID card photo by using the Base64-encoded image (IdOcrPictureBase64), check the photo size and do not pass oversized photos.
	//
	// example:
	//
	// base64
	IdOcrPictureBase64 *string `json:"IdOcrPictureBase64,omitempty" xml:"IdOcrPictureBase64,omitempty"`
	// The file stream of the front image of the ID card.
	//
	// example:
	//
	// InputStream
	IdOcrPictureFile *string `json:"IdOcrPictureFile,omitempty" xml:"IdOcrPictureFile,omitempty"`
	// The URL of the front image of the ID card.
	//
	// example:
	//
	// https://digital-cardocr-prod8.oss-cn-hangzhou.aliyuncs.com/1669520556530-expo/default/face/20221127114236530_w3kx2e6t.jpg
	IdOcrPictureUrl *string `json:"IdOcrPictureUrl,omitempty" xml:"IdOcrPictureUrl,omitempty"`
	// The custom threshold mode for OCR quality check. Valid values:
	//
	// - 0: system default.
	//
	// - 1: strict mode.
	//
	// - 2: loose mode.
	//
	// - 3: disables quality check. Default value: 3.
	//
	// example:
	//
	// 0
	IdThreshold *string `json:"IdThreshold,omitempty" xml:"IdThreshold,omitempty"`
	// The custom unique business identifier on the merchant side, which is used for subsequent troubleshooting. It supports a combination of letters and digits, with a maximum length of 32 characters. Make sure that the identifier is unique.
	//
	// example:
	//
	// dso9322***dsjsd22
	MerchantBizId *string `json:"MerchantBizId,omitempty" xml:"MerchantBizId,omitempty"`
	// The custom user ID in the business. Make sure that the ID is unique.
	//
	// example:
	//
	// 123456789
	MerchantUserId *string `json:"MerchantUserId,omitempty" xml:"MerchantUserId,omitempty"`
	// Specifies whether to perform OCR on the ID card. Valid values:
	//
	// - T: OCR is performed on the ID card.
	//
	// - F: OCR is not performed on the ID card.
	//
	// example:
	//
	// T
	Ocr *string `json:"Ocr,omitempty" xml:"Ocr,omitempty"`
	// The product code.
	//
	// example:
	//
	// Product solution type ID_OCR_MIN
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// Specifies whether to enable anti-spoofing check. Valid values:
	//
	// - T: Anti-spoofing check is enabled.
	//
	// - F: Anti-spoofing check is disabled.
	//
	// example:
	//
	// F
	Spoof *string `json:"Spoof,omitempty" xml:"Spoof,omitempty"`
}

func (s DocOcrV2Request) String() string {
	return dara.Prettify(s)
}

func (s DocOcrV2Request) GoString() string {
	return s.String()
}

func (s *DocOcrV2Request) GetCardSide() *string {
	return s.CardSide
}

func (s *DocOcrV2Request) GetDocType() *string {
	return s.DocType
}

func (s *DocOcrV2Request) GetIdFaceQuality() *string {
	return s.IdFaceQuality
}

func (s *DocOcrV2Request) GetIdOcrPictureBase64() *string {
	return s.IdOcrPictureBase64
}

func (s *DocOcrV2Request) GetIdOcrPictureFile() *string {
	return s.IdOcrPictureFile
}

func (s *DocOcrV2Request) GetIdOcrPictureUrl() *string {
	return s.IdOcrPictureUrl
}

func (s *DocOcrV2Request) GetIdThreshold() *string {
	return s.IdThreshold
}

func (s *DocOcrV2Request) GetMerchantBizId() *string {
	return s.MerchantBizId
}

func (s *DocOcrV2Request) GetMerchantUserId() *string {
	return s.MerchantUserId
}

func (s *DocOcrV2Request) GetOcr() *string {
	return s.Ocr
}

func (s *DocOcrV2Request) GetProductCode() *string {
	return s.ProductCode
}

func (s *DocOcrV2Request) GetSpoof() *string {
	return s.Spoof
}

func (s *DocOcrV2Request) SetCardSide(v string) *DocOcrV2Request {
	s.CardSide = &v
	return s
}

func (s *DocOcrV2Request) SetDocType(v string) *DocOcrV2Request {
	s.DocType = &v
	return s
}

func (s *DocOcrV2Request) SetIdFaceQuality(v string) *DocOcrV2Request {
	s.IdFaceQuality = &v
	return s
}

func (s *DocOcrV2Request) SetIdOcrPictureBase64(v string) *DocOcrV2Request {
	s.IdOcrPictureBase64 = &v
	return s
}

func (s *DocOcrV2Request) SetIdOcrPictureFile(v string) *DocOcrV2Request {
	s.IdOcrPictureFile = &v
	return s
}

func (s *DocOcrV2Request) SetIdOcrPictureUrl(v string) *DocOcrV2Request {
	s.IdOcrPictureUrl = &v
	return s
}

func (s *DocOcrV2Request) SetIdThreshold(v string) *DocOcrV2Request {
	s.IdThreshold = &v
	return s
}

func (s *DocOcrV2Request) SetMerchantBizId(v string) *DocOcrV2Request {
	s.MerchantBizId = &v
	return s
}

func (s *DocOcrV2Request) SetMerchantUserId(v string) *DocOcrV2Request {
	s.MerchantUserId = &v
	return s
}

func (s *DocOcrV2Request) SetOcr(v string) *DocOcrV2Request {
	s.Ocr = &v
	return s
}

func (s *DocOcrV2Request) SetProductCode(v string) *DocOcrV2Request {
	s.ProductCode = &v
	return s
}

func (s *DocOcrV2Request) SetSpoof(v string) *DocOcrV2Request {
	s.Spoof = &v
	return s
}

func (s *DocOcrV2Request) Validate() error {
	return dara.Validate(s)
}
