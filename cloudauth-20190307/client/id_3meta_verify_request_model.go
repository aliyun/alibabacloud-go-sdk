// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iId3MetaVerifyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCrop(v string) *Id3MetaVerifyRequest
	GetCrop() *string
	SetFaceFile(v string) *Id3MetaVerifyRequest
	GetFaceFile() *string
	SetFacePicture(v string) *Id3MetaVerifyRequest
	GetFacePicture() *string
	SetFaceUrl(v string) *Id3MetaVerifyRequest
	GetFaceUrl() *string
	SetIdentifyNum(v string) *Id3MetaVerifyRequest
	GetIdentifyNum() *string
	SetParamType(v string) *Id3MetaVerifyRequest
	GetParamType() *string
	SetUserName(v string) *Id3MetaVerifyRequest
	GetUserName() *string
}

type Id3MetaVerifyRequest struct {
	// Specifies whether to allow cropping of the facial photo. By default, cropping is not allowed. Valid values:
	//
	// - T: allows cropping.
	//
	// - F: does not allow cropping.
	//
	// **Note**
	//
	// If the requested image is not captured by a standard liveness detection SDK, allow cropping of the facial photo. After this feature is enabled, the requested image is first cropped and corrected for face alignment before the service request is initiated.
	//
	// example:
	//
	// T
	Crop *string `json:"Crop,omitempty" xml:"Crop,omitempty"`
	// The input stream of the ID card portrait photo. Specify either CertUrl or CertFile.
	//
	// example:
	//
	// None
	FaceFile *string `json:"FaceFile,omitempty" xml:"FaceFile,omitempty"`
	// The Base64-encoded facial photo.
	//
	// example:
	//
	// /9j/4AAQSkZJRgABAQAASxxxxxxx
	FacePicture *string `json:"FacePicture,omitempty" xml:"FacePicture,omitempty"`
	// The URL of the ID card portrait photo. The URL must be a publicly accessible HTTP or HTTPS link. Specify either CertUrl or CertFile.
	//
	// example:
	//
	// https://www.aliyun.com/cert.jpeg
	FaceUrl *string `json:"FaceUrl,omitempty" xml:"FaceUrl,omitempty"`
	// The ID card number:
	//
	// - If paramType is set to normal: enter the ID card number in plaintext.
	//
	// - If paramType is set to md5: first 6 digits of the ID card number (plaintext) + date of birth (ciphertext) + last 4 digits of the ID card number (plaintext).
	//
	// example:
	//
	// ● Plaintext: 429001********8211
	//
	// ● Ciphertext:
	//
	// 42900132fa7bcd874161bea8ec8fd98f39****8211
	IdentifyNum *string `json:"IdentifyNum,omitempty" xml:"IdentifyNum,omitempty"`
	// The encryption method. Valid values:
	//
	// - normal: plaintext without encryption.
	//
	// - md5: MD5 encryption.
	//
	// **Important**
	//
	// - All encrypted parameter values use 32-character lowercase MD5 strings.
	//
	// - Different MD5 tools may produce different ciphertext. If the API call succeeds with plaintext but fails after encryption, try a different MD5 tool.
	//
	// example:
	//
	// normal
	ParamType *string `json:"ParamType,omitempty" xml:"ParamType,omitempty"`
	// The name:
	//
	// - If paramType is set to normal: enter the name in plaintext.
	//
	// - If paramType is set to md5: ciphertext of the first character of the name + plaintext of the remaining characters.
	//
	// example:
	//
	// ● Plaintext: Zhang San
	//
	// ● Ciphertext:
	//
	// 6499fc7409049355527ef6a2ba5706b8San​
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s Id3MetaVerifyRequest) String() string {
	return dara.Prettify(s)
}

func (s Id3MetaVerifyRequest) GoString() string {
	return s.String()
}

func (s *Id3MetaVerifyRequest) GetCrop() *string {
	return s.Crop
}

func (s *Id3MetaVerifyRequest) GetFaceFile() *string {
	return s.FaceFile
}

func (s *Id3MetaVerifyRequest) GetFacePicture() *string {
	return s.FacePicture
}

func (s *Id3MetaVerifyRequest) GetFaceUrl() *string {
	return s.FaceUrl
}

func (s *Id3MetaVerifyRequest) GetIdentifyNum() *string {
	return s.IdentifyNum
}

func (s *Id3MetaVerifyRequest) GetParamType() *string {
	return s.ParamType
}

func (s *Id3MetaVerifyRequest) GetUserName() *string {
	return s.UserName
}

func (s *Id3MetaVerifyRequest) SetCrop(v string) *Id3MetaVerifyRequest {
	s.Crop = &v
	return s
}

func (s *Id3MetaVerifyRequest) SetFaceFile(v string) *Id3MetaVerifyRequest {
	s.FaceFile = &v
	return s
}

func (s *Id3MetaVerifyRequest) SetFacePicture(v string) *Id3MetaVerifyRequest {
	s.FacePicture = &v
	return s
}

func (s *Id3MetaVerifyRequest) SetFaceUrl(v string) *Id3MetaVerifyRequest {
	s.FaceUrl = &v
	return s
}

func (s *Id3MetaVerifyRequest) SetIdentifyNum(v string) *Id3MetaVerifyRequest {
	s.IdentifyNum = &v
	return s
}

func (s *Id3MetaVerifyRequest) SetParamType(v string) *Id3MetaVerifyRequest {
	s.ParamType = &v
	return s
}

func (s *Id3MetaVerifyRequest) SetUserName(v string) *Id3MetaVerifyRequest {
	s.UserName = &v
	return s
}

func (s *Id3MetaVerifyRequest) Validate() error {
	return dara.Validate(s)
}
