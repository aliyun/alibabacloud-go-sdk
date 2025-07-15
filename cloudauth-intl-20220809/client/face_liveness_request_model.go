// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFaceLivenessRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCrop(v string) *FaceLivenessRequest
	GetCrop() *string
	SetFacePictureBase64(v string) *FaceLivenessRequest
	GetFacePictureBase64() *string
	SetFacePictureUrl(v string) *FaceLivenessRequest
	GetFacePictureUrl() *string
	SetFaceQuality(v string) *FaceLivenessRequest
	GetFaceQuality() *string
	SetMerchantBizId(v string) *FaceLivenessRequest
	GetMerchantBizId() *string
	SetMerchantUserId(v string) *FaceLivenessRequest
	GetMerchantUserId() *string
	SetOcclusion(v string) *FaceLivenessRequest
	GetOcclusion() *string
	SetProductCode(v string) *FaceLivenessRequest
	GetProductCode() *string
}

type FaceLivenessRequest struct {
	// example:
	//
	// T
	Crop              *string `json:"Crop,omitempty" xml:"Crop,omitempty"`
	FacePictureBase64 *string `json:"FacePictureBase64,omitempty" xml:"FacePictureBase64,omitempty"`
	// example:
	//
	// https://digital-face-prod8.oss-cn-hangzhou.aliyuncs.com/1669520556530-expo/default/face/20221127114236530_w3kx2e6t.jpg
	FacePictureUrl *string `json:"FacePictureUrl,omitempty" xml:"FacePictureUrl,omitempty"`
	// example:
	//
	// T
	FaceQuality *string `json:"FaceQuality,omitempty" xml:"FaceQuality,omitempty"`
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
	// T
	Occlusion *string `json:"Occlusion,omitempty" xml:"Occlusion,omitempty"`
	// example:
	//
	// FACE_LIVENESS_MIN
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
}

func (s FaceLivenessRequest) String() string {
	return dara.Prettify(s)
}

func (s FaceLivenessRequest) GoString() string {
	return s.String()
}

func (s *FaceLivenessRequest) GetCrop() *string {
	return s.Crop
}

func (s *FaceLivenessRequest) GetFacePictureBase64() *string {
	return s.FacePictureBase64
}

func (s *FaceLivenessRequest) GetFacePictureUrl() *string {
	return s.FacePictureUrl
}

func (s *FaceLivenessRequest) GetFaceQuality() *string {
	return s.FaceQuality
}

func (s *FaceLivenessRequest) GetMerchantBizId() *string {
	return s.MerchantBizId
}

func (s *FaceLivenessRequest) GetMerchantUserId() *string {
	return s.MerchantUserId
}

func (s *FaceLivenessRequest) GetOcclusion() *string {
	return s.Occlusion
}

func (s *FaceLivenessRequest) GetProductCode() *string {
	return s.ProductCode
}

func (s *FaceLivenessRequest) SetCrop(v string) *FaceLivenessRequest {
	s.Crop = &v
	return s
}

func (s *FaceLivenessRequest) SetFacePictureBase64(v string) *FaceLivenessRequest {
	s.FacePictureBase64 = &v
	return s
}

func (s *FaceLivenessRequest) SetFacePictureUrl(v string) *FaceLivenessRequest {
	s.FacePictureUrl = &v
	return s
}

func (s *FaceLivenessRequest) SetFaceQuality(v string) *FaceLivenessRequest {
	s.FaceQuality = &v
	return s
}

func (s *FaceLivenessRequest) SetMerchantBizId(v string) *FaceLivenessRequest {
	s.MerchantBizId = &v
	return s
}

func (s *FaceLivenessRequest) SetMerchantUserId(v string) *FaceLivenessRequest {
	s.MerchantUserId = &v
	return s
}

func (s *FaceLivenessRequest) SetOcclusion(v string) *FaceLivenessRequest {
	s.Occlusion = &v
	return s
}

func (s *FaceLivenessRequest) SetProductCode(v string) *FaceLivenessRequest {
	s.ProductCode = &v
	return s
}

func (s *FaceLivenessRequest) Validate() error {
	return dara.Validate(s)
}
