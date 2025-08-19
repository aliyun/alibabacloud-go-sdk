// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLivenessFaceVerifyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCertifyId(v string) *LivenessFaceVerifyRequest
	GetCertifyId() *string
	SetCrop(v string) *LivenessFaceVerifyRequest
	GetCrop() *string
	SetDeviceToken(v string) *LivenessFaceVerifyRequest
	GetDeviceToken() *string
	SetFaceContrastPicture(v string) *LivenessFaceVerifyRequest
	GetFaceContrastPicture() *string
	SetFaceContrastPictureUrl(v string) *LivenessFaceVerifyRequest
	GetFaceContrastPictureUrl() *string
	SetIp(v string) *LivenessFaceVerifyRequest
	GetIp() *string
	SetMobile(v string) *LivenessFaceVerifyRequest
	GetMobile() *string
	SetModel(v string) *LivenessFaceVerifyRequest
	GetModel() *string
	SetOssBucketName(v string) *LivenessFaceVerifyRequest
	GetOssBucketName() *string
	SetOssObjectName(v string) *LivenessFaceVerifyRequest
	GetOssObjectName() *string
	SetOuterOrderNo(v string) *LivenessFaceVerifyRequest
	GetOuterOrderNo() *string
	SetProductCode(v string) *LivenessFaceVerifyRequest
	GetProductCode() *string
	SetSceneId(v int64) *LivenessFaceVerifyRequest
	GetSceneId() *int64
	SetUserId(v string) *LivenessFaceVerifyRequest
	GetUserId() *string
}

type LivenessFaceVerifyRequest struct {
	// example:
	//
	// 91707dc296d469ad38e4c5efa6a0f24b
	CertifyId *string `json:"CertifyId,omitempty" xml:"CertifyId,omitempty"`
	// example:
	//
	// T
	Crop *string `json:"Crop,omitempty" xml:"Crop,omitempty"`
	// example:
	//
	// McozS1ZWRcRZStlERcZZo_QOytx5jcgZoZJEoRLOxxxxxxx
	DeviceToken *string `json:"DeviceToken,omitempty" xml:"DeviceToken,omitempty"`
	// example:
	//
	// /9j/4AAQSkZJRgABAQAASxxxxxxx
	FaceContrastPicture *string `json:"FaceContrastPicture,omitempty" xml:"FaceContrastPicture,omitempty"`
	// example:
	//
	// https://ware.cdeledu.com/upfile/uploadPic/2025/03/21/dd62fbb9c966433ab0ba9a7252816b30.jpg
	FaceContrastPictureUrl *string `json:"FaceContrastPictureUrl,omitempty" xml:"FaceContrastPictureUrl,omitempty"`
	// example:
	//
	// 114.xxx.xxx.xxx
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// example:
	//
	// 130xxxxxxxx
	Mobile *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	// example:
	//
	// FRONT_CAMERA_LIVENESS
	Model *string `json:"Model,omitempty" xml:"Model,omitempty"`
	// example:
	//
	// cn-shanghai-aliyun-cloudauth-1494517779820665
	OssBucketName *string `json:"OssBucketName,omitempty" xml:"OssBucketName,omitempty"`
	// example:
	//
	// facedetect/17483113370916034.jpg
	OssObjectName *string `json:"OssObjectName,omitempty" xml:"OssObjectName,omitempty"`
	// example:
	//
	// e0c34a77f5ac40a5aa5e6ed20c353888
	OuterOrderNo *string `json:"OuterOrderNo,omitempty" xml:"OuterOrderNo,omitempty"`
	// example:
	//
	// LR_FR_MIN
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// example:
	//
	// 100000****
	SceneId *int64 `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	// example:
	//
	// 123456789
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s LivenessFaceVerifyRequest) String() string {
	return dara.Prettify(s)
}

func (s LivenessFaceVerifyRequest) GoString() string {
	return s.String()
}

func (s *LivenessFaceVerifyRequest) GetCertifyId() *string {
	return s.CertifyId
}

func (s *LivenessFaceVerifyRequest) GetCrop() *string {
	return s.Crop
}

func (s *LivenessFaceVerifyRequest) GetDeviceToken() *string {
	return s.DeviceToken
}

func (s *LivenessFaceVerifyRequest) GetFaceContrastPicture() *string {
	return s.FaceContrastPicture
}

func (s *LivenessFaceVerifyRequest) GetFaceContrastPictureUrl() *string {
	return s.FaceContrastPictureUrl
}

func (s *LivenessFaceVerifyRequest) GetIp() *string {
	return s.Ip
}

func (s *LivenessFaceVerifyRequest) GetMobile() *string {
	return s.Mobile
}

func (s *LivenessFaceVerifyRequest) GetModel() *string {
	return s.Model
}

func (s *LivenessFaceVerifyRequest) GetOssBucketName() *string {
	return s.OssBucketName
}

func (s *LivenessFaceVerifyRequest) GetOssObjectName() *string {
	return s.OssObjectName
}

func (s *LivenessFaceVerifyRequest) GetOuterOrderNo() *string {
	return s.OuterOrderNo
}

func (s *LivenessFaceVerifyRequest) GetProductCode() *string {
	return s.ProductCode
}

func (s *LivenessFaceVerifyRequest) GetSceneId() *int64 {
	return s.SceneId
}

func (s *LivenessFaceVerifyRequest) GetUserId() *string {
	return s.UserId
}

func (s *LivenessFaceVerifyRequest) SetCertifyId(v string) *LivenessFaceVerifyRequest {
	s.CertifyId = &v
	return s
}

func (s *LivenessFaceVerifyRequest) SetCrop(v string) *LivenessFaceVerifyRequest {
	s.Crop = &v
	return s
}

func (s *LivenessFaceVerifyRequest) SetDeviceToken(v string) *LivenessFaceVerifyRequest {
	s.DeviceToken = &v
	return s
}

func (s *LivenessFaceVerifyRequest) SetFaceContrastPicture(v string) *LivenessFaceVerifyRequest {
	s.FaceContrastPicture = &v
	return s
}

func (s *LivenessFaceVerifyRequest) SetFaceContrastPictureUrl(v string) *LivenessFaceVerifyRequest {
	s.FaceContrastPictureUrl = &v
	return s
}

func (s *LivenessFaceVerifyRequest) SetIp(v string) *LivenessFaceVerifyRequest {
	s.Ip = &v
	return s
}

func (s *LivenessFaceVerifyRequest) SetMobile(v string) *LivenessFaceVerifyRequest {
	s.Mobile = &v
	return s
}

func (s *LivenessFaceVerifyRequest) SetModel(v string) *LivenessFaceVerifyRequest {
	s.Model = &v
	return s
}

func (s *LivenessFaceVerifyRequest) SetOssBucketName(v string) *LivenessFaceVerifyRequest {
	s.OssBucketName = &v
	return s
}

func (s *LivenessFaceVerifyRequest) SetOssObjectName(v string) *LivenessFaceVerifyRequest {
	s.OssObjectName = &v
	return s
}

func (s *LivenessFaceVerifyRequest) SetOuterOrderNo(v string) *LivenessFaceVerifyRequest {
	s.OuterOrderNo = &v
	return s
}

func (s *LivenessFaceVerifyRequest) SetProductCode(v string) *LivenessFaceVerifyRequest {
	s.ProductCode = &v
	return s
}

func (s *LivenessFaceVerifyRequest) SetSceneId(v int64) *LivenessFaceVerifyRequest {
	s.SceneId = &v
	return s
}

func (s *LivenessFaceVerifyRequest) SetUserId(v string) *LivenessFaceVerifyRequest {
	s.UserId = &v
	return s
}

func (s *LivenessFaceVerifyRequest) Validate() error {
	return dara.Validate(s)
}
