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
	// Real name.
	//
	// example:
	//
	// 张三
	CertName *string `json:"CertName,omitempty" xml:"CertName,omitempty"`
	// ID number
	//
	// example:
	//
	// 330103xxxxxxxxxxxx
	CertNo *string `json:"CertNo,omitempty" xml:"CertNo,omitempty"`
	// Type of identification. Currently, only IDENTITY_CARD is supported and must be provided.
	//
	// example:
	//
	// IDENTITY_CARD
	CertType *string `json:"CertType,omitempty" xml:"CertType,omitempty"`
	// The CertifyId of a previously passed real-person authentication, with the photo taken during that authentication used as the comparison photo.
	//
	// > Among the four ways to input images (FaceContrastPicture, FaceContrastPictureUrl, CertifyId, OSS), choose one to provide.
	//
	// example:
	//
	// 0bfa7c493f850e5178b9f8613634c9xx
	CertifyId *string `json:"CertifyId,omitempty" xml:"CertifyId,omitempty"`
	// Allow face image cropping:
	//
	// -  **T*	- – Cropping is allowed.
	//
	// -  **F*	- (default) – Cropping is not allowed.
	//
	// example:
	//
	// T
	Crop *string `json:"Crop,omitempty" xml:"Crop,omitempty"`
	// Risk Identification - Device Token
	//
	// example:
	//
	// McozS1ZWRcRZStlERcZZo_QOytx5jcgZoZJEoRLOxxxxxxx
	DeviceToken *string `json:"DeviceToken,omitempty" xml:"DeviceToken,omitempty"`
	// Encryption type. Leave it empty if no encryption is required.
	//
	// If you enable encrypted transmission, you must specify the encryption algorithm; currently, only the SM2 (Chinese national standard) algorithm is supported.
	//
	// When an encryption algorithm is specified, encrypt both **CertName*	- and **CertNo*	- and submit the resulting ciphertext. For more details on parameter encryption, see the [Parameter Encryption documentation](https://help.aliyun.com/zh/id-verification/financial-grade-id-verification/description-of-parameter-encryption?spm=a2c4g.11186623.0.0.49541a8554cELI#task-2229332).
	//
	// example:
	//
	// SM2
	EncryptType *string `json:"EncryptType,omitempty" xml:"EncryptType,omitempty"`
	// Local video file.
	//
	// example:
	//
	// -
	FaceContrastFile *string `json:"FaceContrastFile,omitempty" xml:"FaceContrastFile,omitempty"`
	// Base64 encoded photo
	//
	// example:
	//
	// /9j/4AAQSkZJRgABAQAASxxxxxxx
	FaceContrastPicture *string `json:"FaceContrastPicture,omitempty" xml:"FaceContrastPicture,omitempty"`
	// OSS photo URL, currently only supports authorized OSS photo URLs.
	//
	// > Among the four ways to input images (FaceContrastPicture, FaceContrastPictureUrl, CertifyId, OSS), choose one to input.
	//
	// example:
	//
	// https://cn-shanghai-aliyun-cloudauth-xxxxxx.oss-cn-shanghai.aliyuncs.com/verify/xxxxx/xxxxx.jpeg
	FaceContrastPictureUrl *string `json:"FaceContrastPictureUrl,omitempty" xml:"FaceContrastPictureUrl,omitempty"`
	// User IP.
	//
	// example:
	//
	// 114.xxx.xxx.xxx
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// User\\"s phone number.
	//
	// example:
	//
	// 130xxxxxxxx
	Mobile *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	// Liveness detection type. Possible values:
	//
	// • **NO_LIVENESS*	- – Liveness detection is disabled.
	//
	// • **FRONT_CAMERA_LIVENESS*	- (default) – Liveness detection on face images captured with the mobile device’s front camera.
	//
	// • **REAR_CAMERA_LIVENESS*	- – Liveness detection on face images captured in other scenarios (e.g., via the rear camera).
	//
	// example:
	//
	// FRONT_CAMERA_LIVENESS
	Model *string `json:"Model,omitempty" xml:"Model,omitempty"`
	// Authorized OSS space Bucket name. In the methods of passing images, including FaceContrastPicture, FaceContrastPictureUrl, CertifyId, and OSS, choose one to pass in.
	//
	// example:
	//
	// cn-shanghai-aliyun-cloudauth-xxxxx
	OssBucketName *string `json:"OssBucketName,omitempty" xml:"OssBucketName,omitempty"`
	// Filename of the authorized OSS space.
	//
	// > Among the four ways to input images (FaceContrastPicture, FaceContrastPictureUrl, CertifyId, OSS), choose one to input.
	//
	// example:
	//
	// verify/xxxxx/xxxxxx.jpeg
	OssObjectName *string `json:"OssObjectName,omitempty" xml:"OssObjectName,omitempty"`
	// A unique identifier for the merchant\\"s request. It is a 32-character alphanumeric combination. The first few characters are a custom abbreviation defined by the merchant, followed by a period, and the latter part can be a random or incrementing sequence.
	//
	// example:
	//
	// e0c34a77f5ac40a5aa5e6ed20c353888
	OuterOrderNo *string `json:"OuterOrderNo,omitempty" xml:"OuterOrderNo,omitempty"`
	// Fixed value: ID_MIN.
	//
	// example:
	//
	// ID_MIN
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// Authentication scenario ID.
	//
	// example:
	//
	// 1000000006
	SceneId *int64 `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	// Custom user ID defined by the customer\\"s business.
	//
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
