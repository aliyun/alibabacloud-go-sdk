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
	// Base64 encoded photo.
	//
	// > Choose one of the three ways to input images: FaceContrastPicture, FaceContrastPictureUrl, or OSS.
	//
	// example:
	//
	// /9j/4AAQSkZJRgABAQAASxxxxxxx
	FaceContrastPicture *string `json:"FaceContrastPicture,omitempty" xml:"FaceContrastPicture,omitempty"`
	// Portrait address, accessible via public HTTP or HTTPS link.
	//
	// > Choose one of the three ways to input images: FaceContrastPicture, FaceContrastPictureUrl, or OSS.
	//
	// example:
	//
	// https://cn-shanghai-aliyun-cloudauth-xxxxxx.oss-cn-shanghai.aliyuncs.com/verify/xxxxx/xxxxx.jpeg
	FaceContrastPictureUrl *string `json:"FaceContrastPictureUrl,omitempty" xml:"FaceContrastPictureUrl,omitempty"`
	// Authorized OSS bucket name.
	//
	// > Choose one of the three ways to input images: FaceContrastPicture, FaceContrastPictureUrl, or OSS.
	//
	// example:
	//
	// cn-shanghai-aliyun-cloudauth-xxxxx
	OssBucketName *string `json:"OssBucketName,omitempty" xml:"OssBucketName,omitempty"`
	// Authorized OSS file name.
	//
	// > Choose one of the three ways to input images: FaceContrastPicture, FaceContrastPictureUrl, or OSS.
	//
	// example:
	//
	// verify/xxxxx/xxxxxx.jpeg
	OssObjectName *string `json:"OssObjectName,omitempty" xml:"OssObjectName,omitempty"`
	// A unique business identifier defined by the client side, used for subsequent troubleshooting. The value should be a combination of letters and numbers with a maximum length of 32 characters, please ensure its uniqueness.
	//
	// example:
	//
	// e0c34a77f5ac40a5aa5e6ed20c353888
	OuterOrderNo *string `json:"OuterOrderNo,omitempty" xml:"OuterOrderNo,omitempty"`
	// Product solution
	//
	// example:
	//
	// LR_FR_AIGC
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// Authentication scene ID. This ID is automatically generated after creating an authentication scene in the console. For how to create an authentication scene, see Adding an Authentication Scene.
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
