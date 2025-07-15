// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeepfakeDetectIntlRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFaceBase64(v string) *DeepfakeDetectIntlRequest
	GetFaceBase64() *string
	SetFaceInputType(v string) *DeepfakeDetectIntlRequest
	GetFaceInputType() *string
	SetFaceUrl(v string) *DeepfakeDetectIntlRequest
	GetFaceUrl() *string
	SetMerchantBizId(v string) *DeepfakeDetectIntlRequest
	GetMerchantBizId() *string
	SetProductCode(v string) *DeepfakeDetectIntlRequest
	GetProductCode() *string
	SetSceneCode(v string) *DeepfakeDetectIntlRequest
	GetSceneCode() *string
}

type DeepfakeDetectIntlRequest struct {
	// example:
	//
	// /9j/4AAQSkZJRgABAQAASxxxxxxx
	FaceBase64 *string `json:"FaceBase64,omitempty" xml:"FaceBase64,omitempty"`
	// example:
	//
	// IMAGE
	FaceInputType *string `json:"FaceInputType,omitempty" xml:"FaceInputType,omitempty"`
	// example:
	//
	// https://cn-shanghai-aliyun-cloudauth-xxxxxx.oss-cn-shanghai.aliyuncs.com/verify/xxxxx/xxxxx.jpeg
	FaceUrl *string `json:"FaceUrl,omitempty" xml:"FaceUrl,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// e0c34a77f5ac40a5aa5e6ed20c******
	MerchantBizId *string `json:"MerchantBizId,omitempty" xml:"MerchantBizId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// FACE_DEEPFAKE
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// example:
	//
	// 1234567890
	SceneCode *string `json:"SceneCode,omitempty" xml:"SceneCode,omitempty"`
}

func (s DeepfakeDetectIntlRequest) String() string {
	return dara.Prettify(s)
}

func (s DeepfakeDetectIntlRequest) GoString() string {
	return s.String()
}

func (s *DeepfakeDetectIntlRequest) GetFaceBase64() *string {
	return s.FaceBase64
}

func (s *DeepfakeDetectIntlRequest) GetFaceInputType() *string {
	return s.FaceInputType
}

func (s *DeepfakeDetectIntlRequest) GetFaceUrl() *string {
	return s.FaceUrl
}

func (s *DeepfakeDetectIntlRequest) GetMerchantBizId() *string {
	return s.MerchantBizId
}

func (s *DeepfakeDetectIntlRequest) GetProductCode() *string {
	return s.ProductCode
}

func (s *DeepfakeDetectIntlRequest) GetSceneCode() *string {
	return s.SceneCode
}

func (s *DeepfakeDetectIntlRequest) SetFaceBase64(v string) *DeepfakeDetectIntlRequest {
	s.FaceBase64 = &v
	return s
}

func (s *DeepfakeDetectIntlRequest) SetFaceInputType(v string) *DeepfakeDetectIntlRequest {
	s.FaceInputType = &v
	return s
}

func (s *DeepfakeDetectIntlRequest) SetFaceUrl(v string) *DeepfakeDetectIntlRequest {
	s.FaceUrl = &v
	return s
}

func (s *DeepfakeDetectIntlRequest) SetMerchantBizId(v string) *DeepfakeDetectIntlRequest {
	s.MerchantBizId = &v
	return s
}

func (s *DeepfakeDetectIntlRequest) SetProductCode(v string) *DeepfakeDetectIntlRequest {
	s.ProductCode = &v
	return s
}

func (s *DeepfakeDetectIntlRequest) SetSceneCode(v string) *DeepfakeDetectIntlRequest {
	s.SceneCode = &v
	return s
}

func (s *DeepfakeDetectIntlRequest) Validate() error {
	return dara.Validate(s)
}
