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
	// example:
	//
	// T
	Authorize *string `json:"Authorize,omitempty" xml:"Authorize,omitempty"`
	// example:
	//
	// 01
	DocPage *string `json:"DocPage,omitempty" xml:"DocPage,omitempty"`
	// example:
	//
	// CHN01001
	DocType *string `json:"DocType,omitempty" xml:"DocType,omitempty"`
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
	// https://***********.oss-cn-hangzhou.aliyuncs.com/1669520556530-expo/default/face/20221127114236530_w3kx2e6t.jpg
	IdOcrPictureUrl *string `json:"IdOcrPictureUrl,omitempty" xml:"IdOcrPictureUrl,omitempty"`
	// example:
	//
	// F
	IdSpoof *string `json:"IdSpoof,omitempty" xml:"IdSpoof,omitempty"`
	// example:
	//
	// 0
	IdThreshold *string `json:"IdThreshold,omitempty" xml:"IdThreshold,omitempty"`
	// example:
	//
	// e0c34a77f5ac40a5aa5e6ed20c353888
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
	// 0
	OcrValueStandard *string `json:"OcrValueStandard,omitempty" xml:"OcrValueStandard,omitempty"`
	// example:
	//
	// ID_OCR_MAX
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
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
