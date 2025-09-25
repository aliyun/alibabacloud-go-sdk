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
	FaceContrastFileObject io.Reader `json:"FaceContrastFile,omitempty" xml:"FaceContrastFile,omitempty"`
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
