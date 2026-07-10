// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAIGCFaceVerifyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFaceContrastPicture(v string) *AIGCFaceVerifyRequest
	GetFaceContrastPicture() *string
	SetFaceContrastPictureUrl(v string) *AIGCFaceVerifyRequest
	GetFaceContrastPictureUrl() *string
	SetOssBucketName(v string) *AIGCFaceVerifyRequest
	GetOssBucketName() *string
	SetOssObjectName(v string) *AIGCFaceVerifyRequest
	GetOssObjectName() *string
	SetOuterOrderNo(v string) *AIGCFaceVerifyRequest
	GetOuterOrderNo() *string
	SetProductCode(v string) *AIGCFaceVerifyRequest
	GetProductCode() *string
	SetSceneId(v int64) *AIGCFaceVerifyRequest
	GetSceneId() *int64
}

type AIGCFaceVerifyRequest struct {
	// The Base64-encoded photo.
	//
	// > You can use one of the following methods to pass in the image: FaceContrastPicture, FaceContrastPictureUrl, or OSS.
	//
	// example:
	//
	// /9j/4AAQSkZJRgABAQAASxxxxxxx
	FaceContrastPicture *string `json:"FaceContrastPicture,omitempty" xml:"FaceContrastPicture,omitempty"`
	// The URL of the face image. The URL must be a publicly accessible HTTP or HTTPS link.
	//
	// > You can use one of the following methods to pass in the image: FaceContrastPicture, FaceContrastPictureUrl, or OSS.
	//
	// example:
	//
	// https://cn-shanghai-aliyun-cloudauth-xxxxxx.oss-cn-shanghai.aliyuncs.com/verify/xxxxx/xxxxx.jpeg
	FaceContrastPictureUrl *string `json:"FaceContrastPictureUrl,omitempty" xml:"FaceContrastPictureUrl,omitempty"`
	// The name of the authorized OSS bucket.
	//
	// > You can use one of the following methods to pass in the image: FaceContrastPicture, FaceContrastPictureUrl, or OSS.
	//
	// example:
	//
	// cn-shanghai-aliyun-cloudauth-xxxxx
	OssBucketName *string `json:"OssBucketName,omitempty" xml:"OssBucketName,omitempty"`
	// The file name in the authorized OSS bucket.
	//
	// > You can use one of the following methods to pass in the image: FaceContrastPicture, FaceContrastPictureUrl, or OSS.
	//
	// example:
	//
	// verify/xxxxx/xxxxxx.jpeg
	OssObjectName *string `json:"OssObjectName,omitempty" xml:"OssObjectName,omitempty"`
	// The custom business unique identifier on the client side, used for subsequent troubleshooting. The value can contain up to 32 characters, including letters and digits. Make sure the value is unique.
	//
	// example:
	//
	// e0c34a77f5ac40a5aa5e6ed20c353888
	OuterOrderNo *string `json:"OuterOrderNo,omitempty" xml:"OuterOrderNo,omitempty"`
	// The product plan.
	//
	// example:
	//
	// LR_FR_AIGC
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The ID of the verification scenario. This ID is automatically generated after you create a verification scenario in the console. For more information about how to create a verification scenario, refer to Add a verification scenario.
	//
	// example:
	//
	// 100000xxxx
	SceneId *int64 `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
}

func (s AIGCFaceVerifyRequest) String() string {
	return dara.Prettify(s)
}

func (s AIGCFaceVerifyRequest) GoString() string {
	return s.String()
}

func (s *AIGCFaceVerifyRequest) GetFaceContrastPicture() *string {
	return s.FaceContrastPicture
}

func (s *AIGCFaceVerifyRequest) GetFaceContrastPictureUrl() *string {
	return s.FaceContrastPictureUrl
}

func (s *AIGCFaceVerifyRequest) GetOssBucketName() *string {
	return s.OssBucketName
}

func (s *AIGCFaceVerifyRequest) GetOssObjectName() *string {
	return s.OssObjectName
}

func (s *AIGCFaceVerifyRequest) GetOuterOrderNo() *string {
	return s.OuterOrderNo
}

func (s *AIGCFaceVerifyRequest) GetProductCode() *string {
	return s.ProductCode
}

func (s *AIGCFaceVerifyRequest) GetSceneId() *int64 {
	return s.SceneId
}

func (s *AIGCFaceVerifyRequest) SetFaceContrastPicture(v string) *AIGCFaceVerifyRequest {
	s.FaceContrastPicture = &v
	return s
}

func (s *AIGCFaceVerifyRequest) SetFaceContrastPictureUrl(v string) *AIGCFaceVerifyRequest {
	s.FaceContrastPictureUrl = &v
	return s
}

func (s *AIGCFaceVerifyRequest) SetOssBucketName(v string) *AIGCFaceVerifyRequest {
	s.OssBucketName = &v
	return s
}

func (s *AIGCFaceVerifyRequest) SetOssObjectName(v string) *AIGCFaceVerifyRequest {
	s.OssObjectName = &v
	return s
}

func (s *AIGCFaceVerifyRequest) SetOuterOrderNo(v string) *AIGCFaceVerifyRequest {
	s.OuterOrderNo = &v
	return s
}

func (s *AIGCFaceVerifyRequest) SetProductCode(v string) *AIGCFaceVerifyRequest {
	s.ProductCode = &v
	return s
}

func (s *AIGCFaceVerifyRequest) SetSceneId(v int64) *AIGCFaceVerifyRequest {
	s.SceneId = &v
	return s
}

func (s *AIGCFaceVerifyRequest) Validate() error {
	return dara.Validate(s)
}
