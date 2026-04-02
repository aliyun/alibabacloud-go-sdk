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
	// example:
	//
	// OCR_ID_FACE
	CardSide *string `json:"CardSide,omitempty" xml:"CardSide,omitempty"`
	// example:
	//
	// 00000006
	DocType *string `json:"DocType,omitempty" xml:"DocType,omitempty"`
	// example:
	//
	// F
	IdFaceQuality *string `json:"IdFaceQuality,omitempty" xml:"IdFaceQuality,omitempty"`
	// example:
	//
	// base64
	IdOcrPictureBase64 *string `json:"IdOcrPictureBase64,omitempty" xml:"IdOcrPictureBase64,omitempty"`
	// example:
	//
	// InputStream
	IdOcrPictureFileObject io.Reader `json:"IdOcrPictureFile,omitempty" xml:"IdOcrPictureFile,omitempty"`
	// example:
	//
	// https://digital-cardocr-prod8.oss-cn-hangzhou.aliyuncs.com/1669520556530-expo/default/face/20221127114236530_w3kx2e6t.jpg
	IdOcrPictureUrl *string `json:"IdOcrPictureUrl,omitempty" xml:"IdOcrPictureUrl,omitempty"`
	// example:
	//
	// 0
	IdThreshold *string `json:"IdThreshold,omitempty" xml:"IdThreshold,omitempty"`
	// example:
	//
	// dso9322***dsjsd22
	MerchantBizId *string `json:"MerchantBizId,omitempty" xml:"MerchantBizId,omitempty"`
	// example:
	//
	// 123456789
	MerchantUserId *string `json:"MerchantUserId,omitempty" xml:"MerchantUserId,omitempty"`
	// example:
	//
	// T
	Ocr         *string `json:"Ocr,omitempty" xml:"Ocr,omitempty"`
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
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
