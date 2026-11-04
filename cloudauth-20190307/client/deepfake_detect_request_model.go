// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeepfakeDetectRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFaceAttributeCheck(v string) *DeepfakeDetectRequest
	GetFaceAttributeCheck() *string
	SetFaceBase64(v string) *DeepfakeDetectRequest
	GetFaceBase64() *string
	SetFaceInputType(v string) *DeepfakeDetectRequest
	GetFaceInputType() *string
	SetFaceUrl(v string) *DeepfakeDetectRequest
	GetFaceUrl() *string
	SetOuterOrderNo(v string) *DeepfakeDetectRequest
	GetOuterOrderNo() *string
}

type DeepfakeDetectRequest struct {
	// Specifies whether to enable facial attribute detection.
	//
	// example:
	//
	// Y
	FaceAttributeCheck *string `json:"FaceAttributeCheck,omitempty" xml:"FaceAttributeCheck,omitempty"`
	// The Base64-encoded string of the facial image.
	//
	// > You must specify either FaceUrl or FaceBase64.
	//
	// example:
	//
	// /9j/4AAQSkZJRgABAQAASxxxxxxx
	FaceBase64 *string `json:"FaceBase64,omitempty" xml:"FaceBase64,omitempty"`
	// The input type of the facial material. Valid values:
	//
	// - IMAGE (default): facial image.
	//
	// - VIDEO: facial video.
	//
	// > Video processing takes a long time. Set the timeout period to greater than 3 seconds.
	//
	// example:
	//
	// IMAGE
	FaceInputType *string `json:"FaceInputType,omitempty" xml:"FaceInputType,omitempty"`
	// The URL of the facial image.
	//
	// > You must specify either FaceUrl or FaceBase64.
	//
	// example:
	//
	// https://cn-shanghai-aliyun-cloudauth-xxxxxx.oss-cn-shanghai.aliyuncs.com/verify/xxxxx/xxxxx.jpeg
	FaceUrl *string `json:"FaceUrl,omitempty" xml:"FaceUrl,omitempty"`
	// The unique identifier of the merchant request. The value is a 32-character alphanumeric string. The first few characters consist of a custom abbreviation defined by the merchant, the middle part can contain a time period, and the latter part can be a random or incremental sequence.
	//
	// example:
	//
	// e0c34a77f5ac40a5aa5e6ed20c******
	OuterOrderNo *string `json:"OuterOrderNo,omitempty" xml:"OuterOrderNo,omitempty"`
}

func (s DeepfakeDetectRequest) String() string {
	return dara.Prettify(s)
}

func (s DeepfakeDetectRequest) GoString() string {
	return s.String()
}

func (s *DeepfakeDetectRequest) GetFaceAttributeCheck() *string {
	return s.FaceAttributeCheck
}

func (s *DeepfakeDetectRequest) GetFaceBase64() *string {
	return s.FaceBase64
}

func (s *DeepfakeDetectRequest) GetFaceInputType() *string {
	return s.FaceInputType
}

func (s *DeepfakeDetectRequest) GetFaceUrl() *string {
	return s.FaceUrl
}

func (s *DeepfakeDetectRequest) GetOuterOrderNo() *string {
	return s.OuterOrderNo
}

func (s *DeepfakeDetectRequest) SetFaceAttributeCheck(v string) *DeepfakeDetectRequest {
	s.FaceAttributeCheck = &v
	return s
}

func (s *DeepfakeDetectRequest) SetFaceBase64(v string) *DeepfakeDetectRequest {
	s.FaceBase64 = &v
	return s
}

func (s *DeepfakeDetectRequest) SetFaceInputType(v string) *DeepfakeDetectRequest {
	s.FaceInputType = &v
	return s
}

func (s *DeepfakeDetectRequest) SetFaceUrl(v string) *DeepfakeDetectRequest {
	s.FaceUrl = &v
	return s
}

func (s *DeepfakeDetectRequest) SetOuterOrderNo(v string) *DeepfakeDetectRequest {
	s.OuterOrderNo = &v
	return s
}

func (s *DeepfakeDetectRequest) Validate() error {
	return dara.Validate(s)
}
