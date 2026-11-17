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
	// The Base64-encoded string of the face image. For videos, pass the data as a stream.
	//
	// example:
	//
	// /9j/4AAQSkZJRgABAQAASxxxxxxx
	FaceBase64 *string `json:"FaceBase64,omitempty" xml:"FaceBase64,omitempty"`
	// The input stream of the image.
	//
	// example:
	//
	// -
	FaceFile *string `json:"FaceFile,omitempty" xml:"FaceFile,omitempty"`
	// The input type of the face material. Valid values:
	//
	// - IMAGE: face image. This is the default value.
	//
	// - VIDEO: face video.
	//
	// Note: Video processing takes a long time. Set the timeout period to a value greater than 3 seconds.
	//
	// example:
	//
	// IMAGE
	FaceInputType *string `json:"FaceInputType,omitempty" xml:"FaceInputType,omitempty"`
	// The URL of the face image.
	//
	// example:
	//
	// https://cn-shanghai-aliyun-cloudauth-xxxxxx.oss-cn-shanghai.aliyuncs.com/verify/xxxxx/xxxxx.jpeg
	FaceUrl *string `json:"FaceUrl,omitempty" xml:"FaceUrl,omitempty"`
	// The unique identifier of the merchant request. The value is a 32-character alphanumeric string. The first few characters consist of a custom merchant abbreviation, the middle part can contain a time period, and the latter part can be a random or incremental sequence.
	//
	// example:
	//
	// e0c34a***353888
	MerchantBizId *string `json:"MerchantBizId,omitempty" xml:"MerchantBizId,omitempty"`
	// The product solution to be integrated. Valid values: FACE_DEEPFAKE.
	//
	// example:
	//
	// FACE_DEEPFAKE
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The custom authentication scenario ID, which is used to query related records by entering this scenario ID in the console. The value must be 10 characters in length and can contain letters, digits, and underscores.
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
