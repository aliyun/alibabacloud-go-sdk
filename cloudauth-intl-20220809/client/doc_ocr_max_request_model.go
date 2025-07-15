// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDocOcrMaxRequest interface {
	dara.Model
	String() string
	GoString() string
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
	DocPage *string `json:"DocPage,omitempty" xml:"DocPage,omitempty"`
	// example:
	//
	// CNSSC01
	DocType *string `json:"DocType,omitempty" xml:"DocType,omitempty"`
	// example:
	//
	// base64
	IdOcrPictureBase64 *string `json:"IdOcrPictureBase64,omitempty" xml:"IdOcrPictureBase64,omitempty"`
	// example:
	//
	// https://***********.oss-cn-hangzhou.aliyuncs.com/1669520556530-expo/default/face/20221127114236530_w3kx2e6t.jpg
	IdOcrPictureUrl *string `json:"IdOcrPictureUrl,omitempty" xml:"IdOcrPictureUrl,omitempty"`
	IdSpoof         *string `json:"IdSpoof,omitempty" xml:"IdSpoof,omitempty"`
	// example:
	//
	// 0
	IdThreshold *string `json:"IdThreshold,omitempty" xml:"IdThreshold,omitempty"`
	// example:
	//
	// e0c34a77f5ac40a5aa5e6ed20c******
	MerchantBizId *string `json:"MerchantBizId,omitempty" xml:"MerchantBizId,omitempty"`
	// example:
	//
	// 123456789
	MerchantUserId *string `json:"MerchantUserId,omitempty" xml:"MerchantUserId,omitempty"`
	// example:
	//
	// 0
	OcrModel *string `json:"OcrModel,omitempty" xml:"OcrModel,omitempty"`
	// example:
	//
	// ID_OCR_MAX
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	Prompt      *string `json:"Prompt,omitempty" xml:"Prompt,omitempty"`
	// example:
	//
	// 1234567890
	SceneCode *string `json:"SceneCode,omitempty" xml:"SceneCode,omitempty"`
	// example:
	//
	// F
	Spoof *string `json:"Spoof,omitempty" xml:"Spoof,omitempty"`
}

func (s DocOcrMaxRequest) String() string {
	return dara.Prettify(s)
}

func (s DocOcrMaxRequest) GoString() string {
	return s.String()
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
