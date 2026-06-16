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
	IdOcrPictureFile *string `json:"IdOcrPictureFile,omitempty" xml:"IdOcrPictureFile,omitempty"`
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
