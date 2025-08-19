// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iContrastFaceVerifyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCertName(v string) *ContrastFaceVerifyRequest
	GetCertName() *string
	SetCertNo(v string) *ContrastFaceVerifyRequest
	GetCertNo() *string
	SetCertType(v string) *ContrastFaceVerifyRequest
	GetCertType() *string
	SetCertifyId(v string) *ContrastFaceVerifyRequest
	GetCertifyId() *string
	SetCrop(v string) *ContrastFaceVerifyRequest
	GetCrop() *string
	SetDeviceToken(v string) *ContrastFaceVerifyRequest
	GetDeviceToken() *string
	SetEncryptType(v string) *ContrastFaceVerifyRequest
	GetEncryptType() *string
	SetFaceContrastFile(v string) *ContrastFaceVerifyRequest
	GetFaceContrastFile() *string
	SetFaceContrastPicture(v string) *ContrastFaceVerifyRequest
	GetFaceContrastPicture() *string
	SetFaceContrastPictureUrl(v string) *ContrastFaceVerifyRequest
	GetFaceContrastPictureUrl() *string
	SetIp(v string) *ContrastFaceVerifyRequest
	GetIp() *string
	SetMobile(v string) *ContrastFaceVerifyRequest
	GetMobile() *string
	SetModel(v string) *ContrastFaceVerifyRequest
	GetModel() *string
	SetOssBucketName(v string) *ContrastFaceVerifyRequest
	GetOssBucketName() *string
	SetOssObjectName(v string) *ContrastFaceVerifyRequest
	GetOssObjectName() *string
	SetOuterOrderNo(v string) *ContrastFaceVerifyRequest
	GetOuterOrderNo() *string
	SetProductCode(v string) *ContrastFaceVerifyRequest
	GetProductCode() *string
	SetSceneId(v int64) *ContrastFaceVerifyRequest
	GetSceneId() *int64
	SetUserId(v string) *ContrastFaceVerifyRequest
	GetUserId() *string
}

type ContrastFaceVerifyRequest struct {
	CertName *string `json:"CertName,omitempty" xml:"CertName,omitempty"`
	// example:
	//
	// 330103xxxxxxxxxxxx
	CertNo *string `json:"CertNo,omitempty" xml:"CertNo,omitempty"`
	// example:
	//
	// IDENTITY_CARD
	CertType *string `json:"CertType,omitempty" xml:"CertType,omitempty"`
	// example:
	//
	// 0bfa7c493f850e5178b9f8613634c9xx
	CertifyId *string `json:"CertifyId,omitempty" xml:"CertifyId,omitempty"`
	Crop      *string `json:"Crop,omitempty" xml:"Crop,omitempty"`
	// example:
	//
	// McozS1ZWRcRZStlERcZZo_QOytx5jcgZoZJEoRLOxxxxxxx
	DeviceToken      *string `json:"DeviceToken,omitempty" xml:"DeviceToken,omitempty"`
	EncryptType      *string `json:"EncryptType,omitempty" xml:"EncryptType,omitempty"`
	FaceContrastFile *string `json:"FaceContrastFile,omitempty" xml:"FaceContrastFile,omitempty"`
	// example:
	//
	// /9j/4AAQSkZJRgABAQAASxxxxxxx
	FaceContrastPicture *string `json:"FaceContrastPicture,omitempty" xml:"FaceContrastPicture,omitempty"`
	// example:
	//
	// https://cn-shanghai-aliyun-cloudauth-xxxxxx.oss-cn-shanghai.aliyuncs.com/verify/xxxxx/xxxxx.jpeg
	FaceContrastPictureUrl *string `json:"FaceContrastPictureUrl,omitempty" xml:"FaceContrastPictureUrl,omitempty"`
	// example:
	//
	// 114.xxx.xxx.xxx
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// example:
	//
	// 130xxxxxxxx
	Mobile *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	Model  *string `json:"Model,omitempty" xml:"Model,omitempty"`
	// example:
	//
	// cn-shanghai-aliyun-cloudauth-xxxxx
	OssBucketName *string `json:"OssBucketName,omitempty" xml:"OssBucketName,omitempty"`
	// example:
	//
	// verify/xxxxx/xxxxxx.jpeg
	OssObjectName *string `json:"OssObjectName,omitempty" xml:"OssObjectName,omitempty"`
	// example:
	//
	// e0c34a77f5ac40a5aa5e6ed20c353888
	OuterOrderNo *string `json:"OuterOrderNo,omitempty" xml:"OuterOrderNo,omitempty"`
	// example:
	//
	// ID_MIN
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// example:
	//
	// 1000000006
	SceneId *int64 `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	// example:
	//
	// 123456789
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ContrastFaceVerifyRequest) String() string {
	return dara.Prettify(s)
}

func (s ContrastFaceVerifyRequest) GoString() string {
	return s.String()
}

func (s *ContrastFaceVerifyRequest) GetCertName() *string {
	return s.CertName
}

func (s *ContrastFaceVerifyRequest) GetCertNo() *string {
	return s.CertNo
}

func (s *ContrastFaceVerifyRequest) GetCertType() *string {
	return s.CertType
}

func (s *ContrastFaceVerifyRequest) GetCertifyId() *string {
	return s.CertifyId
}

func (s *ContrastFaceVerifyRequest) GetCrop() *string {
	return s.Crop
}

func (s *ContrastFaceVerifyRequest) GetDeviceToken() *string {
	return s.DeviceToken
}

func (s *ContrastFaceVerifyRequest) GetEncryptType() *string {
	return s.EncryptType
}

func (s *ContrastFaceVerifyRequest) GetFaceContrastFile() *string {
	return s.FaceContrastFile
}

func (s *ContrastFaceVerifyRequest) GetFaceContrastPicture() *string {
	return s.FaceContrastPicture
}

func (s *ContrastFaceVerifyRequest) GetFaceContrastPictureUrl() *string {
	return s.FaceContrastPictureUrl
}

func (s *ContrastFaceVerifyRequest) GetIp() *string {
	return s.Ip
}

func (s *ContrastFaceVerifyRequest) GetMobile() *string {
	return s.Mobile
}

func (s *ContrastFaceVerifyRequest) GetModel() *string {
	return s.Model
}

func (s *ContrastFaceVerifyRequest) GetOssBucketName() *string {
	return s.OssBucketName
}

func (s *ContrastFaceVerifyRequest) GetOssObjectName() *string {
	return s.OssObjectName
}

func (s *ContrastFaceVerifyRequest) GetOuterOrderNo() *string {
	return s.OuterOrderNo
}

func (s *ContrastFaceVerifyRequest) GetProductCode() *string {
	return s.ProductCode
}

func (s *ContrastFaceVerifyRequest) GetSceneId() *int64 {
	return s.SceneId
}

func (s *ContrastFaceVerifyRequest) GetUserId() *string {
	return s.UserId
}

func (s *ContrastFaceVerifyRequest) SetCertName(v string) *ContrastFaceVerifyRequest {
	s.CertName = &v
	return s
}

func (s *ContrastFaceVerifyRequest) SetCertNo(v string) *ContrastFaceVerifyRequest {
	s.CertNo = &v
	return s
}

func (s *ContrastFaceVerifyRequest) SetCertType(v string) *ContrastFaceVerifyRequest {
	s.CertType = &v
	return s
}

func (s *ContrastFaceVerifyRequest) SetCertifyId(v string) *ContrastFaceVerifyRequest {
	s.CertifyId = &v
	return s
}

func (s *ContrastFaceVerifyRequest) SetCrop(v string) *ContrastFaceVerifyRequest {
	s.Crop = &v
	return s
}

func (s *ContrastFaceVerifyRequest) SetDeviceToken(v string) *ContrastFaceVerifyRequest {
	s.DeviceToken = &v
	return s
}

func (s *ContrastFaceVerifyRequest) SetEncryptType(v string) *ContrastFaceVerifyRequest {
	s.EncryptType = &v
	return s
}

func (s *ContrastFaceVerifyRequest) SetFaceContrastFile(v string) *ContrastFaceVerifyRequest {
	s.FaceContrastFile = &v
	return s
}

func (s *ContrastFaceVerifyRequest) SetFaceContrastPicture(v string) *ContrastFaceVerifyRequest {
	s.FaceContrastPicture = &v
	return s
}

func (s *ContrastFaceVerifyRequest) SetFaceContrastPictureUrl(v string) *ContrastFaceVerifyRequest {
	s.FaceContrastPictureUrl = &v
	return s
}

func (s *ContrastFaceVerifyRequest) SetIp(v string) *ContrastFaceVerifyRequest {
	s.Ip = &v
	return s
}

func (s *ContrastFaceVerifyRequest) SetMobile(v string) *ContrastFaceVerifyRequest {
	s.Mobile = &v
	return s
}

func (s *ContrastFaceVerifyRequest) SetModel(v string) *ContrastFaceVerifyRequest {
	s.Model = &v
	return s
}

func (s *ContrastFaceVerifyRequest) SetOssBucketName(v string) *ContrastFaceVerifyRequest {
	s.OssBucketName = &v
	return s
}

func (s *ContrastFaceVerifyRequest) SetOssObjectName(v string) *ContrastFaceVerifyRequest {
	s.OssObjectName = &v
	return s
}

func (s *ContrastFaceVerifyRequest) SetOuterOrderNo(v string) *ContrastFaceVerifyRequest {
	s.OuterOrderNo = &v
	return s
}

func (s *ContrastFaceVerifyRequest) SetProductCode(v string) *ContrastFaceVerifyRequest {
	s.ProductCode = &v
	return s
}

func (s *ContrastFaceVerifyRequest) SetSceneId(v int64) *ContrastFaceVerifyRequest {
	s.SceneId = &v
	return s
}

func (s *ContrastFaceVerifyRequest) SetUserId(v string) *ContrastFaceVerifyRequest {
	s.UserId = &v
	return s
}

func (s *ContrastFaceVerifyRequest) Validate() error {
	return dara.Validate(s)
}
