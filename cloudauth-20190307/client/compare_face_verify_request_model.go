// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCompareFaceVerifyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCrop(v string) *CompareFaceVerifyRequest
	GetCrop() *string
	SetOuterOrderNo(v string) *CompareFaceVerifyRequest
	GetOuterOrderNo() *string
	SetProductCode(v string) *CompareFaceVerifyRequest
	GetProductCode() *string
	SetSceneId(v int64) *CompareFaceVerifyRequest
	GetSceneId() *int64
	SetSourceCertifyId(v string) *CompareFaceVerifyRequest
	GetSourceCertifyId() *string
	SetSourceFaceContrastPicture(v string) *CompareFaceVerifyRequest
	GetSourceFaceContrastPicture() *string
	SetSourceFaceContrastPictureUrl(v string) *CompareFaceVerifyRequest
	GetSourceFaceContrastPictureUrl() *string
	SetSourceOssBucketName(v string) *CompareFaceVerifyRequest
	GetSourceOssBucketName() *string
	SetSourceOssObjectName(v string) *CompareFaceVerifyRequest
	GetSourceOssObjectName() *string
	SetTargetCertifyId(v string) *CompareFaceVerifyRequest
	GetTargetCertifyId() *string
	SetTargetFaceContrastPicture(v string) *CompareFaceVerifyRequest
	GetTargetFaceContrastPicture() *string
	SetTargetFaceContrastPictureUrl(v string) *CompareFaceVerifyRequest
	GetTargetFaceContrastPictureUrl() *string
	SetTargetOssBucketName(v string) *CompareFaceVerifyRequest
	GetTargetOssBucketName() *string
	SetTargetOssObjectName(v string) *CompareFaceVerifyRequest
	GetTargetOssObjectName() *string
}

type CompareFaceVerifyRequest struct {
	// Whether cropping is allowed. Default is not allowed, T/F.
	//
	// - T: Indicates that cropping is required
	//
	// - F: Indicates that cropping is not required (default F)
	//
	// example:
	//
	// T
	Crop *string `json:"Crop,omitempty" xml:"Crop,omitempty"`
	// A unique identifier for the merchant\\"s request. The value is a 32-character alphanumeric combination, where the first few characters are a custom abbreviation defined by the merchant, followed by a period, and the latter part can be a random or incrementing sequence.
	//
	// example:
	//
	// e0c34a77f5ac40a5aa5e6ed20c353888
	OuterOrderNo *string `json:"OuterOrderNo,omitempty" xml:"OuterOrderNo,omitempty"`
	// Fixed value: PV_FC.
	//
	// example:
	//
	// PV_FC
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// Authentication scenario ID.
	//
	// example:
	//
	// 1000000006
	SceneId *int64 `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	// The CertifyId of a previously successful real-person verification, where the photo taken during that verification is used as the face comparison photo.
	//
	// > Among the four ways to input facial photos (FaceContrastPicture, FaceContrastPictureUrl, CertifyId, OSS), choose one to provide.
	//
	// example:
	//
	// 0bfa7c493f850e5178b9f8613634c9xx
	SourceCertifyId *string `json:"SourceCertifyId,omitempty" xml:"SourceCertifyId,omitempty"`
	// Base64 encoding of the photo.
	//
	// > Choose one of the four ways to input a face photo: FaceContrastPicture, FaceContrastPictureUrl, CertifyId, or OSS.
	//
	// example:
	//
	// /9j/4AAQSkZJRgABAQAASxxxxxxx
	SourceFaceContrastPicture *string `json:"SourceFaceContrastPicture,omitempty" xml:"SourceFaceContrastPicture,omitempty"`
	// OSS photo URL, currently only supports authorized OSS photo URLs.
	//
	// > Four ways to input face photos: FaceContrastPicture, FaceContrastPictureUrl, CertifyId, and OSS. Choose one of them to input.
	//
	// example:
	//
	// https://cn-shanghai-aliyun-cloudauth-xxxxxx.oss-cn-shanghai.aliyuncs.com/verify/xxxxx/xxxxx.jpeg
	SourceFaceContrastPictureUrl *string `json:"SourceFaceContrastPictureUrl,omitempty" xml:"SourceFaceContrastPictureUrl,omitempty"`
	// Name of the authorized OSS bucket.
	//
	// > Choose one of the four ways to input face photos: FaceContrastPicture, FaceContrastPictureUrl, CertifyId, or OSS.
	//
	// example:
	//
	// cn-shanghai-aliyun-cloudauth-xxxxx
	SourceOssBucketName *string `json:"SourceOssBucketName,omitempty" xml:"SourceOssBucketName,omitempty"`
	// Filename of the authorized OSS space.
	//
	// > Choose one of the four ways to input face photos: FaceContrastPicture, FaceContrastPictureUrl, CertifyId, or OSS.
	//
	// example:
	//
	// verify/xxxxx/xxxxxx.jpeg
	SourceOssObjectName *string `json:"SourceOssObjectName,omitempty" xml:"SourceOssObjectName,omitempty"`
	// CertifyId from a previously successful real-person authentication, where the photo taken during the authentication is used for face comparison.
	//
	// > Choose one of the four methods to provide the reference face photo: FaceContrastPicture, FaceContrastPictureUrl, CertifyId, or OSS.
	//
	// example:
	//
	// 0bfa7c493f850e5178b9f8613634c9xx
	TargetCertifyId *string `json:"TargetCertifyId,omitempty" xml:"TargetCertifyId,omitempty"`
	// Base64 encoding of the reference photo.
	//
	// > Choose one of the four methods to provide the reference face photo: FaceContrastPicture, FaceContrastPictureUrl, CertifyId, or OSS.
	//
	// example:
	//
	// /9j/4AAQSkZJRgABAQAASxxxxxxx
	TargetFaceContrastPicture *string `json:"TargetFaceContrastPicture,omitempty" xml:"TargetFaceContrastPicture,omitempty"`
	// OSS address of the reference photo. Currently, only authorized OSS addresses are supported.
	//
	// > Choose one of the four methods to provide the reference face photo: FaceContrastPicture, FaceContrastPictureUrl, CertifyId, or OSS.
	//
	// example:
	//
	// https://cn-shanghai-aliyun-cloudauth-xxxxxx.oss-cn-shanghai.aliyuncs.com/verify/xxxxx/xxxxx.jpeg
	TargetFaceContrastPictureUrl *string `json:"TargetFaceContrastPictureUrl,omitempty" xml:"TargetFaceContrastPictureUrl,omitempty"`
	// Name of the authorized OSS bucket.
	//
	// > Choose one of the four methods to provide the reference face photo: FaceContrastPicture, FaceContrastPictureUrl, CertifyId, or OSS.
	//
	// example:
	//
	// cn-shanghai-aliyun-cloudauth-xxxxx
	TargetOssBucketName *string `json:"TargetOssBucketName,omitempty" xml:"TargetOssBucketName,omitempty"`
	// File name in the authorized OSS space.
	//
	// > Choose one of the four methods to provide the reference face photo: FaceContrastPicture, FaceContrastPictureUrl, CertifyId, or OSS.
	//
	// example:
	//
	// verify/xxxxx/xxxxxx.jpeg
	TargetOssObjectName *string `json:"TargetOssObjectName,omitempty" xml:"TargetOssObjectName,omitempty"`
}

func (s CompareFaceVerifyRequest) String() string {
	return dara.Prettify(s)
}

func (s CompareFaceVerifyRequest) GoString() string {
	return s.String()
}

func (s *CompareFaceVerifyRequest) GetCrop() *string {
	return s.Crop
}

func (s *CompareFaceVerifyRequest) GetOuterOrderNo() *string {
	return s.OuterOrderNo
}

func (s *CompareFaceVerifyRequest) GetProductCode() *string {
	return s.ProductCode
}

func (s *CompareFaceVerifyRequest) GetSceneId() *int64 {
	return s.SceneId
}

func (s *CompareFaceVerifyRequest) GetSourceCertifyId() *string {
	return s.SourceCertifyId
}

func (s *CompareFaceVerifyRequest) GetSourceFaceContrastPicture() *string {
	return s.SourceFaceContrastPicture
}

func (s *CompareFaceVerifyRequest) GetSourceFaceContrastPictureUrl() *string {
	return s.SourceFaceContrastPictureUrl
}

func (s *CompareFaceVerifyRequest) GetSourceOssBucketName() *string {
	return s.SourceOssBucketName
}

func (s *CompareFaceVerifyRequest) GetSourceOssObjectName() *string {
	return s.SourceOssObjectName
}

func (s *CompareFaceVerifyRequest) GetTargetCertifyId() *string {
	return s.TargetCertifyId
}

func (s *CompareFaceVerifyRequest) GetTargetFaceContrastPicture() *string {
	return s.TargetFaceContrastPicture
}

func (s *CompareFaceVerifyRequest) GetTargetFaceContrastPictureUrl() *string {
	return s.TargetFaceContrastPictureUrl
}

func (s *CompareFaceVerifyRequest) GetTargetOssBucketName() *string {
	return s.TargetOssBucketName
}

func (s *CompareFaceVerifyRequest) GetTargetOssObjectName() *string {
	return s.TargetOssObjectName
}

func (s *CompareFaceVerifyRequest) SetCrop(v string) *CompareFaceVerifyRequest {
	s.Crop = &v
	return s
}

func (s *CompareFaceVerifyRequest) SetOuterOrderNo(v string) *CompareFaceVerifyRequest {
	s.OuterOrderNo = &v
	return s
}

func (s *CompareFaceVerifyRequest) SetProductCode(v string) *CompareFaceVerifyRequest {
	s.ProductCode = &v
	return s
}

func (s *CompareFaceVerifyRequest) SetSceneId(v int64) *CompareFaceVerifyRequest {
	s.SceneId = &v
	return s
}

func (s *CompareFaceVerifyRequest) SetSourceCertifyId(v string) *CompareFaceVerifyRequest {
	s.SourceCertifyId = &v
	return s
}

func (s *CompareFaceVerifyRequest) SetSourceFaceContrastPicture(v string) *CompareFaceVerifyRequest {
	s.SourceFaceContrastPicture = &v
	return s
}

func (s *CompareFaceVerifyRequest) SetSourceFaceContrastPictureUrl(v string) *CompareFaceVerifyRequest {
	s.SourceFaceContrastPictureUrl = &v
	return s
}

func (s *CompareFaceVerifyRequest) SetSourceOssBucketName(v string) *CompareFaceVerifyRequest {
	s.SourceOssBucketName = &v
	return s
}

func (s *CompareFaceVerifyRequest) SetSourceOssObjectName(v string) *CompareFaceVerifyRequest {
	s.SourceOssObjectName = &v
	return s
}

func (s *CompareFaceVerifyRequest) SetTargetCertifyId(v string) *CompareFaceVerifyRequest {
	s.TargetCertifyId = &v
	return s
}

func (s *CompareFaceVerifyRequest) SetTargetFaceContrastPicture(v string) *CompareFaceVerifyRequest {
	s.TargetFaceContrastPicture = &v
	return s
}

func (s *CompareFaceVerifyRequest) SetTargetFaceContrastPictureUrl(v string) *CompareFaceVerifyRequest {
	s.TargetFaceContrastPictureUrl = &v
	return s
}

func (s *CompareFaceVerifyRequest) SetTargetOssBucketName(v string) *CompareFaceVerifyRequest {
	s.TargetOssBucketName = &v
	return s
}

func (s *CompareFaceVerifyRequest) SetTargetOssObjectName(v string) *CompareFaceVerifyRequest {
	s.TargetOssObjectName = &v
	return s
}

func (s *CompareFaceVerifyRequest) Validate() error {
	return dara.Validate(s)
}
