// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeepfakeDetectIntlStreamRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFaceBase64(v string) *DeepfakeDetectIntlStreamRequest
	GetFaceBase64() *string
	SetFaceFile(v string) *DeepfakeDetectIntlStreamRequest
	GetFaceFile() *string
	SetFaceInputType(v string) *DeepfakeDetectIntlStreamRequest
	GetFaceInputType() *string
	SetFaceUrl(v string) *DeepfakeDetectIntlStreamRequest
	GetFaceUrl() *string
	SetMerchantBizId(v string) *DeepfakeDetectIntlStreamRequest
	GetMerchantBizId() *string
	SetProductCode(v string) *DeepfakeDetectIntlStreamRequest
	GetProductCode() *string
	SetSceneCode(v string) *DeepfakeDetectIntlStreamRequest
	GetSceneCode() *string
}

type DeepfakeDetectIntlStreamRequest struct {
	// Enter the Base64 encoded format of the face image; for video formats, it is recommended to input via stream.
	//
	// example:
	//
	// /9j/4AAQSkZJRgABAQAASxxxxxxx
	FaceBase64 *string `json:"FaceBase64,omitempty" xml:"FaceBase64,omitempty"`
	// Image input stream.
	//
	// example:
	//
	// -
	FaceFile *string `json:"FaceFile,omitempty" xml:"FaceFile,omitempty"`
	// Face material input type:
	//
	// - IMAGE (default): Face image
	//
	// - VIDEO: Face video
	//
	// Note: Video processing takes longer, it is recommended to set the timeout > 3S.
	//
	// example:
	//
	// IMAGE
	FaceInputType *string `json:"FaceInputType,omitempty" xml:"FaceInputType,omitempty"`
	// Enter the URL address of the face image.
	//
	// example:
	//
	// https://cn-shanghai-aliyun-cloudauth-xxxxxx.oss-cn-shanghai.aliyuncs.com/verify/xxxxx/xxxxx.jpeg
	FaceUrl *string `json:"FaceUrl,omitempty" xml:"FaceUrl,omitempty"`
	// A unique identifier for the merchant\\"s request, consisting of a 32-character alphanumeric combination.
	//
	// The first few characters are composed of a custom abbreviation defined by the merchant, the middle part can include a period of time, and the latter part can use a random or incremental sequence.
	//
	// example:
	//
	// e0c34a***353888
	MerchantBizId *string `json:"MerchantBizId,omitempty" xml:"MerchantBizId,omitempty"`
	// The product solution to be integrated.
	//
	// Value: FACE_DEEPFAKE
	//
	// example:
	//
	// FACE_DEEPFAKE
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// Your custom authentication scenario ID, used for querying related records by entering this scenario ID in the console later.
	//
	// Supports a combination of 10 characters, including letters, numbers, or underscores.
	//
	// example:
	//
	// 123****123
	SceneCode *string `json:"SceneCode,omitempty" xml:"SceneCode,omitempty"`
}

func (s DeepfakeDetectIntlStreamRequest) String() string {
	return dara.Prettify(s)
}

func (s DeepfakeDetectIntlStreamRequest) GoString() string {
	return s.String()
}

func (s *DeepfakeDetectIntlStreamRequest) GetFaceBase64() *string {
	return s.FaceBase64
}

func (s *DeepfakeDetectIntlStreamRequest) GetFaceFile() *string {
	return s.FaceFile
}

func (s *DeepfakeDetectIntlStreamRequest) GetFaceInputType() *string {
	return s.FaceInputType
}

func (s *DeepfakeDetectIntlStreamRequest) GetFaceUrl() *string {
	return s.FaceUrl
}

func (s *DeepfakeDetectIntlStreamRequest) GetMerchantBizId() *string {
	return s.MerchantBizId
}

func (s *DeepfakeDetectIntlStreamRequest) GetProductCode() *string {
	return s.ProductCode
}

func (s *DeepfakeDetectIntlStreamRequest) GetSceneCode() *string {
	return s.SceneCode
}

func (s *DeepfakeDetectIntlStreamRequest) SetFaceBase64(v string) *DeepfakeDetectIntlStreamRequest {
	s.FaceBase64 = &v
	return s
}

func (s *DeepfakeDetectIntlStreamRequest) SetFaceFile(v string) *DeepfakeDetectIntlStreamRequest {
	s.FaceFile = &v
	return s
}

func (s *DeepfakeDetectIntlStreamRequest) SetFaceInputType(v string) *DeepfakeDetectIntlStreamRequest {
	s.FaceInputType = &v
	return s
}

func (s *DeepfakeDetectIntlStreamRequest) SetFaceUrl(v string) *DeepfakeDetectIntlStreamRequest {
	s.FaceUrl = &v
	return s
}

func (s *DeepfakeDetectIntlStreamRequest) SetMerchantBizId(v string) *DeepfakeDetectIntlStreamRequest {
	s.MerchantBizId = &v
	return s
}

func (s *DeepfakeDetectIntlStreamRequest) SetProductCode(v string) *DeepfakeDetectIntlStreamRequest {
	s.ProductCode = &v
	return s
}

func (s *DeepfakeDetectIntlStreamRequest) SetSceneCode(v string) *DeepfakeDetectIntlStreamRequest {
	s.SceneCode = &v
	return s
}

func (s *DeepfakeDetectIntlStreamRequest) Validate() error {
	return dara.Validate(s)
}
