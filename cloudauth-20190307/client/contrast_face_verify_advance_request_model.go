// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
	"io"
)

type iContrastFaceVerifyAdvanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCertName(v string) *ContrastFaceVerifyAdvanceRequest
	GetCertName() *string
	SetCertNo(v string) *ContrastFaceVerifyAdvanceRequest
	GetCertNo() *string
	SetCertType(v string) *ContrastFaceVerifyAdvanceRequest
	GetCertType() *string
	SetCertifyId(v string) *ContrastFaceVerifyAdvanceRequest
	GetCertifyId() *string
	SetCrop(v string) *ContrastFaceVerifyAdvanceRequest
	GetCrop() *string
	SetDeviceToken(v string) *ContrastFaceVerifyAdvanceRequest
	GetDeviceToken() *string
	SetEncryptType(v string) *ContrastFaceVerifyAdvanceRequest
	GetEncryptType() *string
	SetFaceContrastFileObject(v io.Reader) *ContrastFaceVerifyAdvanceRequest
	GetFaceContrastFileObject() io.Reader
	SetFaceContrastPicture(v string) *ContrastFaceVerifyAdvanceRequest
	GetFaceContrastPicture() *string
	SetFaceContrastPictureUrl(v string) *ContrastFaceVerifyAdvanceRequest
	GetFaceContrastPictureUrl() *string
	SetIp(v string) *ContrastFaceVerifyAdvanceRequest
	GetIp() *string
	SetMobile(v string) *ContrastFaceVerifyAdvanceRequest
	GetMobile() *string
	SetModel(v string) *ContrastFaceVerifyAdvanceRequest
	GetModel() *string
	SetOssBucketName(v string) *ContrastFaceVerifyAdvanceRequest
	GetOssBucketName() *string
	SetOssObjectName(v string) *ContrastFaceVerifyAdvanceRequest
	GetOssObjectName() *string
	SetOuterOrderNo(v string) *ContrastFaceVerifyAdvanceRequest
	GetOuterOrderNo() *string
	SetProductCode(v string) *ContrastFaceVerifyAdvanceRequest
	GetProductCode() *string
	SetSceneId(v int64) *ContrastFaceVerifyAdvanceRequest
	GetSceneId() *int64
	SetUserId(v string) *ContrastFaceVerifyAdvanceRequest
	GetUserId() *string
}

type ContrastFaceVerifyAdvanceRequest struct {
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
	DeviceToken            *string   `json:"DeviceToken,omitempty" xml:"DeviceToken,omitempty"`
	EncryptType            *string   `json:"EncryptType,omitempty" xml:"EncryptType,omitempty"`
	FaceContrastFileObject io.Reader `json:"FaceContrastFile,omitempty" xml:"FaceContrastFile,omitempty"`
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

func (s ContrastFaceVerifyAdvanceRequest) String() string {
	return dara.Prettify(s)
}

func (s ContrastFaceVerifyAdvanceRequest) GoString() string {
	return s.String()
}

func (s *ContrastFaceVerifyAdvanceRequest) GetCertName() *string {
	return s.CertName
}

func (s *ContrastFaceVerifyAdvanceRequest) GetCertNo() *string {
	return s.CertNo
}

func (s *ContrastFaceVerifyAdvanceRequest) GetCertType() *string {
	return s.CertType
}

func (s *ContrastFaceVerifyAdvanceRequest) GetCertifyId() *string {
	return s.CertifyId
}

func (s *ContrastFaceVerifyAdvanceRequest) GetCrop() *string {
	return s.Crop
}

func (s *ContrastFaceVerifyAdvanceRequest) GetDeviceToken() *string {
	return s.DeviceToken
}

func (s *ContrastFaceVerifyAdvanceRequest) GetEncryptType() *string {
	return s.EncryptType
}

func (s *ContrastFaceVerifyAdvanceRequest) GetFaceContrastFileObject() io.Reader {
	return s.FaceContrastFileObject
}

func (s *ContrastFaceVerifyAdvanceRequest) GetFaceContrastPicture() *string {
	return s.FaceContrastPicture
}

func (s *ContrastFaceVerifyAdvanceRequest) GetFaceContrastPictureUrl() *string {
	return s.FaceContrastPictureUrl
}

func (s *ContrastFaceVerifyAdvanceRequest) GetIp() *string {
	return s.Ip
}

func (s *ContrastFaceVerifyAdvanceRequest) GetMobile() *string {
	return s.Mobile
}

func (s *ContrastFaceVerifyAdvanceRequest) GetModel() *string {
	return s.Model
}

func (s *ContrastFaceVerifyAdvanceRequest) GetOssBucketName() *string {
	return s.OssBucketName
}

func (s *ContrastFaceVerifyAdvanceRequest) GetOssObjectName() *string {
	return s.OssObjectName
}

func (s *ContrastFaceVerifyAdvanceRequest) GetOuterOrderNo() *string {
	return s.OuterOrderNo
}

func (s *ContrastFaceVerifyAdvanceRequest) GetProductCode() *string {
	return s.ProductCode
}

func (s *ContrastFaceVerifyAdvanceRequest) GetSceneId() *int64 {
	return s.SceneId
}

func (s *ContrastFaceVerifyAdvanceRequest) GetUserId() *string {
	return s.UserId
}

func (s *ContrastFaceVerifyAdvanceRequest) SetCertName(v string) *ContrastFaceVerifyAdvanceRequest {
	s.CertName = &v
	return s
}

func (s *ContrastFaceVerifyAdvanceRequest) SetCertNo(v string) *ContrastFaceVerifyAdvanceRequest {
	s.CertNo = &v
	return s
}

func (s *ContrastFaceVerifyAdvanceRequest) SetCertType(v string) *ContrastFaceVerifyAdvanceRequest {
	s.CertType = &v
	return s
}

func (s *ContrastFaceVerifyAdvanceRequest) SetCertifyId(v string) *ContrastFaceVerifyAdvanceRequest {
	s.CertifyId = &v
	return s
}

func (s *ContrastFaceVerifyAdvanceRequest) SetCrop(v string) *ContrastFaceVerifyAdvanceRequest {
	s.Crop = &v
	return s
}

func (s *ContrastFaceVerifyAdvanceRequest) SetDeviceToken(v string) *ContrastFaceVerifyAdvanceRequest {
	s.DeviceToken = &v
	return s
}

func (s *ContrastFaceVerifyAdvanceRequest) SetEncryptType(v string) *ContrastFaceVerifyAdvanceRequest {
	s.EncryptType = &v
	return s
}

func (s *ContrastFaceVerifyAdvanceRequest) SetFaceContrastFileObject(v io.Reader) *ContrastFaceVerifyAdvanceRequest {
	s.FaceContrastFileObject = v
	return s
}

func (s *ContrastFaceVerifyAdvanceRequest) SetFaceContrastPicture(v string) *ContrastFaceVerifyAdvanceRequest {
	s.FaceContrastPicture = &v
	return s
}

func (s *ContrastFaceVerifyAdvanceRequest) SetFaceContrastPictureUrl(v string) *ContrastFaceVerifyAdvanceRequest {
	s.FaceContrastPictureUrl = &v
	return s
}

func (s *ContrastFaceVerifyAdvanceRequest) SetIp(v string) *ContrastFaceVerifyAdvanceRequest {
	s.Ip = &v
	return s
}

func (s *ContrastFaceVerifyAdvanceRequest) SetMobile(v string) *ContrastFaceVerifyAdvanceRequest {
	s.Mobile = &v
	return s
}

func (s *ContrastFaceVerifyAdvanceRequest) SetModel(v string) *ContrastFaceVerifyAdvanceRequest {
	s.Model = &v
	return s
}

func (s *ContrastFaceVerifyAdvanceRequest) SetOssBucketName(v string) *ContrastFaceVerifyAdvanceRequest {
	s.OssBucketName = &v
	return s
}

func (s *ContrastFaceVerifyAdvanceRequest) SetOssObjectName(v string) *ContrastFaceVerifyAdvanceRequest {
	s.OssObjectName = &v
	return s
}

func (s *ContrastFaceVerifyAdvanceRequest) SetOuterOrderNo(v string) *ContrastFaceVerifyAdvanceRequest {
	s.OuterOrderNo = &v
	return s
}

func (s *ContrastFaceVerifyAdvanceRequest) SetProductCode(v string) *ContrastFaceVerifyAdvanceRequest {
	s.ProductCode = &v
	return s
}

func (s *ContrastFaceVerifyAdvanceRequest) SetSceneId(v int64) *ContrastFaceVerifyAdvanceRequest {
	s.SceneId = &v
	return s
}

func (s *ContrastFaceVerifyAdvanceRequest) SetUserId(v string) *ContrastFaceVerifyAdvanceRequest {
	s.UserId = &v
	return s
}

func (s *ContrastFaceVerifyAdvanceRequest) Validate() error {
	return dara.Validate(s)
}
