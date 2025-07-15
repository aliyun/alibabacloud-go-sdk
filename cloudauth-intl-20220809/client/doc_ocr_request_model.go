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
	CardSide *string `json:"CardSide,omitempty" xml:"CardSide,omitempty"`
	// example:
	//
	// 00000006
	DocType *string `json:"DocType,omitempty" xml:"DocType,omitempty"`
	// example:
	//
	// F
	IdFaceQuality      *string `json:"IdFaceQuality,omitempty" xml:"IdFaceQuality,omitempty"`
	IdOcrPictureBase64 *string `json:"IdOcrPictureBase64,omitempty" xml:"IdOcrPictureBase64,omitempty"`
	// example:
	//
	// https://digital-cardocr-prod8.oss-cn-hangzhou.aliyuncs.com/1669520556530-expo/default/face/20221127114236530_w3kx2e6t.jpg
	IdOcrPictureUrl *string `json:"IdOcrPictureUrl,omitempty" xml:"IdOcrPictureUrl,omitempty"`
	IdThreshold     *string `json:"IdThreshold,omitempty" xml:"IdThreshold,omitempty"`
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
