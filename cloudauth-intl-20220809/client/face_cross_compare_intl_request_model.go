// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFaceCrossCompareIntlRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCompareModel(v string) *FaceCrossCompareIntlRequest
	GetCompareModel() *string
	SetFaceVerifyThreshold(v string) *FaceCrossCompareIntlRequest
	GetFaceVerifyThreshold() *string
	SetMerchantBizId(v string) *FaceCrossCompareIntlRequest
	GetMerchantBizId() *string
	SetProductCode(v string) *FaceCrossCompareIntlRequest
	GetProductCode() *string
	SetSceneCode(v string) *FaceCrossCompareIntlRequest
	GetSceneCode() *string
	SetSourceAFacePicture(v string) *FaceCrossCompareIntlRequest
	GetSourceAFacePicture() *string
	SetSourceAFacePictureUrl(v string) *FaceCrossCompareIntlRequest
	GetSourceAFacePictureUrl() *string
	SetSourceBFacePicture(v string) *FaceCrossCompareIntlRequest
	GetSourceBFacePicture() *string
	SetSourceBFacePictureUrl(v string) *FaceCrossCompareIntlRequest
	GetSourceBFacePictureUrl() *string
	SetSourceCFacePicture(v string) *FaceCrossCompareIntlRequest
	GetSourceCFacePicture() *string
	SetSourceCFacePictureUrl(v string) *FaceCrossCompareIntlRequest
	GetSourceCFacePictureUrl() *string
}

type FaceCrossCompareIntlRequest struct {
	// example:
	//
	// 0
	CompareModel *string `json:"CompareModel,omitempty" xml:"CompareModel,omitempty"`
	// example:
	//
	// 0.5
	FaceVerifyThreshold *string `json:"FaceVerifyThreshold,omitempty" xml:"FaceVerifyThreshold,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// e0c34a77f5ac40a5aa5e6ed20c35****
	MerchantBizId *string `json:"MerchantBizId,omitempty" xml:"MerchantBizId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// FACE_CROSS_COMPARE
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// example:
	//
	// 1234567890
	SceneCode *string `json:"SceneCode,omitempty" xml:"SceneCode,omitempty"`
	// example:
	//
	// base64
	SourceAFacePicture *string `json:"SourceAFacePicture,omitempty" xml:"SourceAFacePicture,omitempty"`
	// example:
	//
	// https://www.xxx.com/1.jpg
	SourceAFacePictureUrl *string `json:"SourceAFacePictureUrl,omitempty" xml:"SourceAFacePictureUrl,omitempty"`
	// example:
	//
	// base64
	SourceBFacePicture *string `json:"SourceBFacePicture,omitempty" xml:"SourceBFacePicture,omitempty"`
	// example:
	//
	// https://www.xxx.com/1.jpg
	SourceBFacePictureUrl *string `json:"SourceBFacePictureUrl,omitempty" xml:"SourceBFacePictureUrl,omitempty"`
	// example:
	//
	// base64
	SourceCFacePicture *string `json:"SourceCFacePicture,omitempty" xml:"SourceCFacePicture,omitempty"`
	// example:
	//
	// https://www.xxx.com/1.jpg
	SourceCFacePictureUrl *string `json:"SourceCFacePictureUrl,omitempty" xml:"SourceCFacePictureUrl,omitempty"`
}

func (s FaceCrossCompareIntlRequest) String() string {
	return dara.Prettify(s)
}

func (s FaceCrossCompareIntlRequest) GoString() string {
	return s.String()
}

func (s *FaceCrossCompareIntlRequest) GetCompareModel() *string {
	return s.CompareModel
}

func (s *FaceCrossCompareIntlRequest) GetFaceVerifyThreshold() *string {
	return s.FaceVerifyThreshold
}

func (s *FaceCrossCompareIntlRequest) GetMerchantBizId() *string {
	return s.MerchantBizId
}

func (s *FaceCrossCompareIntlRequest) GetProductCode() *string {
	return s.ProductCode
}

func (s *FaceCrossCompareIntlRequest) GetSceneCode() *string {
	return s.SceneCode
}

func (s *FaceCrossCompareIntlRequest) GetSourceAFacePicture() *string {
	return s.SourceAFacePicture
}

func (s *FaceCrossCompareIntlRequest) GetSourceAFacePictureUrl() *string {
	return s.SourceAFacePictureUrl
}

func (s *FaceCrossCompareIntlRequest) GetSourceBFacePicture() *string {
	return s.SourceBFacePicture
}

func (s *FaceCrossCompareIntlRequest) GetSourceBFacePictureUrl() *string {
	return s.SourceBFacePictureUrl
}

func (s *FaceCrossCompareIntlRequest) GetSourceCFacePicture() *string {
	return s.SourceCFacePicture
}

func (s *FaceCrossCompareIntlRequest) GetSourceCFacePictureUrl() *string {
	return s.SourceCFacePictureUrl
}

func (s *FaceCrossCompareIntlRequest) SetCompareModel(v string) *FaceCrossCompareIntlRequest {
	s.CompareModel = &v
	return s
}

func (s *FaceCrossCompareIntlRequest) SetFaceVerifyThreshold(v string) *FaceCrossCompareIntlRequest {
	s.FaceVerifyThreshold = &v
	return s
}

func (s *FaceCrossCompareIntlRequest) SetMerchantBizId(v string) *FaceCrossCompareIntlRequest {
	s.MerchantBizId = &v
	return s
}

func (s *FaceCrossCompareIntlRequest) SetProductCode(v string) *FaceCrossCompareIntlRequest {
	s.ProductCode = &v
	return s
}

func (s *FaceCrossCompareIntlRequest) SetSceneCode(v string) *FaceCrossCompareIntlRequest {
	s.SceneCode = &v
	return s
}

func (s *FaceCrossCompareIntlRequest) SetSourceAFacePicture(v string) *FaceCrossCompareIntlRequest {
	s.SourceAFacePicture = &v
	return s
}

func (s *FaceCrossCompareIntlRequest) SetSourceAFacePictureUrl(v string) *FaceCrossCompareIntlRequest {
	s.SourceAFacePictureUrl = &v
	return s
}

func (s *FaceCrossCompareIntlRequest) SetSourceBFacePicture(v string) *FaceCrossCompareIntlRequest {
	s.SourceBFacePicture = &v
	return s
}

func (s *FaceCrossCompareIntlRequest) SetSourceBFacePictureUrl(v string) *FaceCrossCompareIntlRequest {
	s.SourceBFacePictureUrl = &v
	return s
}

func (s *FaceCrossCompareIntlRequest) SetSourceCFacePicture(v string) *FaceCrossCompareIntlRequest {
	s.SourceCFacePicture = &v
	return s
}

func (s *FaceCrossCompareIntlRequest) SetSourceCFacePictureUrl(v string) *FaceCrossCompareIntlRequest {
	s.SourceCFacePictureUrl = &v
	return s
}

func (s *FaceCrossCompareIntlRequest) Validate() error {
	return dara.Validate(s)
}
