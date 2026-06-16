// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
	"io"
)

type iDeepfakeDetectIntlStreamAdvanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFaceBase64(v string) *DeepfakeDetectIntlStreamAdvanceRequest
	GetFaceBase64() *string
	SetFaceFileObject(v io.Reader) *DeepfakeDetectIntlStreamAdvanceRequest
	GetFaceFileObject() io.Reader
	SetFaceInputType(v string) *DeepfakeDetectIntlStreamAdvanceRequest
	GetFaceInputType() *string
	SetFaceUrl(v string) *DeepfakeDetectIntlStreamAdvanceRequest
	GetFaceUrl() *string
	SetMerchantBizId(v string) *DeepfakeDetectIntlStreamAdvanceRequest
	GetMerchantBizId() *string
	SetProductCode(v string) *DeepfakeDetectIntlStreamAdvanceRequest
	GetProductCode() *string
	SetSceneCode(v string) *DeepfakeDetectIntlStreamAdvanceRequest
	GetSceneCode() *string
}

type DeepfakeDetectIntlStreamAdvanceRequest struct {
	// The Base64-encoded face image. For videos, we recommend that you use the stream method for transmission.
	//
	// example:
	//
	// /9j/4AAQSkZJRgABAQAASxxxxxxx
	FaceBase64 *string `json:"FaceBase64,omitempty" xml:"FaceBase64,omitempty"`
	// The image input stream.
	//
	// example:
	//
	// -
	FaceFileObject io.Reader `json:"FaceFile,omitempty" xml:"FaceFile,omitempty"`
	// The type of facial material input:
	//
	// - IMAGE (default): Face image
	//
	// - VIDEO: Face video
	//
	// Note
	//
	// Video processing takes a long time. We recommend that you set the timeout period to more than 3 seconds.
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
	// The unique identifier of the merchant request. The value is an alphanumeric string with a length of 32 characters.
	//
	// The first few characters consist of a custom abbreviation defined by the merchant, the middle part can contain a timestamp, and the last part can use a random or incremental sequence.
	//
	// example:
	//
	// e0c34a***353888
	MerchantBizId *string `json:"MerchantBizId,omitempty" xml:"MerchantBizId,omitempty"`
	// The product solution to integrate.
	//
	// Valid value: FACE_DEEPFAKE
	//
	// example:
	//
	// FACE_DEEPFAKE
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// A custom verification scenario ID that you define. This ID is used to query related records in the console.
	//
	// The value is a combination of letters, digits, or underscores (_) with a maximum length of 10 characters.
	//
	// example:
	//
	// 123****123
	SceneCode *string `json:"SceneCode,omitempty" xml:"SceneCode,omitempty"`
}

func (s DeepfakeDetectIntlStreamAdvanceRequest) String() string {
	return dara.Prettify(s)
}

func (s DeepfakeDetectIntlStreamAdvanceRequest) GoString() string {
	return s.String()
}

func (s *DeepfakeDetectIntlStreamAdvanceRequest) GetFaceBase64() *string {
	return s.FaceBase64
}

func (s *DeepfakeDetectIntlStreamAdvanceRequest) GetFaceFileObject() io.Reader {
	return s.FaceFileObject
}

func (s *DeepfakeDetectIntlStreamAdvanceRequest) GetFaceInputType() *string {
	return s.FaceInputType
}

func (s *DeepfakeDetectIntlStreamAdvanceRequest) GetFaceUrl() *string {
	return s.FaceUrl
}

func (s *DeepfakeDetectIntlStreamAdvanceRequest) GetMerchantBizId() *string {
	return s.MerchantBizId
}

func (s *DeepfakeDetectIntlStreamAdvanceRequest) GetProductCode() *string {
	return s.ProductCode
}

func (s *DeepfakeDetectIntlStreamAdvanceRequest) GetSceneCode() *string {
	return s.SceneCode
}

func (s *DeepfakeDetectIntlStreamAdvanceRequest) SetFaceBase64(v string) *DeepfakeDetectIntlStreamAdvanceRequest {
	s.FaceBase64 = &v
	return s
}

func (s *DeepfakeDetectIntlStreamAdvanceRequest) SetFaceFileObject(v io.Reader) *DeepfakeDetectIntlStreamAdvanceRequest {
	s.FaceFileObject = v
	return s
}

func (s *DeepfakeDetectIntlStreamAdvanceRequest) SetFaceInputType(v string) *DeepfakeDetectIntlStreamAdvanceRequest {
	s.FaceInputType = &v
	return s
}

func (s *DeepfakeDetectIntlStreamAdvanceRequest) SetFaceUrl(v string) *DeepfakeDetectIntlStreamAdvanceRequest {
	s.FaceUrl = &v
	return s
}

func (s *DeepfakeDetectIntlStreamAdvanceRequest) SetMerchantBizId(v string) *DeepfakeDetectIntlStreamAdvanceRequest {
	s.MerchantBizId = &v
	return s
}

func (s *DeepfakeDetectIntlStreamAdvanceRequest) SetProductCode(v string) *DeepfakeDetectIntlStreamAdvanceRequest {
	s.ProductCode = &v
	return s
}

func (s *DeepfakeDetectIntlStreamAdvanceRequest) SetSceneCode(v string) *DeepfakeDetectIntlStreamAdvanceRequest {
	s.SceneCode = &v
	return s
}

func (s *DeepfakeDetectIntlStreamAdvanceRequest) Validate() error {
	return dara.Validate(s)
}
