// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iId3MetaVerifyPRORequest interface {
	dara.Model
	String() string
	GoString() string
	SetCrop(v string) *Id3MetaVerifyPRORequest
	GetCrop() *string
	SetEnableFallback(v string) *Id3MetaVerifyPRORequest
	GetEnableFallback() *string
	SetFaceFile(v string) *Id3MetaVerifyPRORequest
	GetFaceFile() *string
	SetFacePicture(v string) *Id3MetaVerifyPRORequest
	GetFacePicture() *string
	SetFaceUrl(v string) *Id3MetaVerifyPRORequest
	GetFaceUrl() *string
	SetIdentifyNum(v string) *Id3MetaVerifyPRORequest
	GetIdentifyNum() *string
	SetLivenessCheck(v string) *Id3MetaVerifyPRORequest
	GetLivenessCheck() *string
	SetParamType(v string) *Id3MetaVerifyPRORequest
	GetParamType() *string
	SetUserName(v string) *Id3MetaVerifyPRORequest
	GetUserName() *string
}

type Id3MetaVerifyPRORequest struct {
	// Specifies whether to allow cropping of the facial photo. By default, cropping is not allowed. Valid values:
	//
	// - **T**: Cropping is allowed.
	//
	// - **F**: Cropping is not allowed.
	//
	// > If the requested image is not captured by a standard liveness detection SDK, allow cropping of the facial photo.
	//
	// After this feature is enabled, the requested image is first cropped and corrected for the face, and then the request is sent to the service.
	//
	// example:
	//
	// T
	Crop *string `json:"Crop,omitempty" xml:"Crop,omitempty"`
	// Specifies whether to allow fallback to a non-public security source. Valid values:
	//
	// - **N*	- (default): Disabled.
	//
	// - **Y**: Enabled.
	//
	// example:
	//
	// Y
	EnableFallback *string `json:"EnableFallback,omitempty" xml:"EnableFallback,omitempty"`
	// The input stream of the facial photo.
	//
	// example:
	//
	// For the specific integration method, refer to the file upload Advance API configuration
	FaceFile *string `json:"FaceFile,omitempty" xml:"FaceFile,omitempty"`
	// The Base64-encoded photo. If you use this method to submit the facial photo, check the photo size and do not submit an excessively large photo.
	//
	// example:
	//
	// /9j/4AAQSkZJRgABAQAASxxxxxxx
	FacePicture *string `json:"FacePicture,omitempty" xml:"FacePicture,omitempty"`
	// The URL of the facial photo. The URL must be a publicly accessible HTTP or HTTPS link.
	//
	// example:
	//
	// https://cn-shanghai-aliyun-cloudauth-xxxxxx.oss-cn-shanghai.aliyuncs.com/verify/xxxxx/xxxxx.jpeg
	FaceUrl *string `json:"FaceUrl,omitempty" xml:"FaceUrl,omitempty"`
	// The ID card number.
	//
	// - If **paramType*	- is set to normal: Enter the ID card number in plaintext.
	//
	// - If **paramType*	- is set to sm2: Enter the encrypted ID card number.
	//
	//
	// > Due to authoritative source limitations, only second-generation resident ID card numbers are supported.
	//
	// example:
	//
	// Plaintext: 429001********8211
	//
	// Ciphertext: MHsCIEPDt1ycBNSVgA2yKsMnvWGheiI+STWqJLmYWlZnklhnAiEAwq1bk5YVepHwqfMsd9ErlK71OrdXx8E+wfqbzyFCwoMEIM1QdrFeekvQh6fwK7sVXAsNePiNm7Eulqm/zahRosbKBBKgtyhm3SopJ3tO/wALKXQQW+g=
	IdentifyNum *string `json:"IdentifyNum,omitempty" xml:"IdentifyNum,omitempty"`
	// Specifies whether to enable liveness detection. Valid values:
	//
	// - **N*	- (default): Liveness detection is disabled.
	//
	// - **Y**: Liveness detection is enabled.
	//
	// example:
	//
	// Y
	LivenessCheck *string `json:"LivenessCheck,omitempty" xml:"LivenessCheck,omitempty"`
	// The encryption method. Valid values:
	//
	// - **normal**: Plaintext without encryption.
	//
	// - **sm2**: SM2 encryption.
	//
	// example:
	//
	// normal
	ParamType *string `json:"ParamType,omitempty" xml:"ParamType,omitempty"`
	// The name.
	//
	// - If **paramType*	- is set to normal: Enter the name in plaintext.
	//
	// - If **paramType*	- is set to sm2: Enter the encrypted name.
	//
	// example:
	//
	// Plaintext: Zhang San
	//
	// Ciphertext: MG8CIQCxI0wNYbc0c2BRL+7+tSethTXfQC391ZFnszRRcvRZ9AIgekQYhgDtxaDuoRD4bde/5fnFdlUp4YoxlEnIFLLm2mQEILkctuy6Rw6lfAUxBtkpPoPVCJAeD5al/RX8JFUvTACEBAYAgLJjBe0=
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s Id3MetaVerifyPRORequest) String() string {
	return dara.Prettify(s)
}

func (s Id3MetaVerifyPRORequest) GoString() string {
	return s.String()
}

func (s *Id3MetaVerifyPRORequest) GetCrop() *string {
	return s.Crop
}

func (s *Id3MetaVerifyPRORequest) GetEnableFallback() *string {
	return s.EnableFallback
}

func (s *Id3MetaVerifyPRORequest) GetFaceFile() *string {
	return s.FaceFile
}

func (s *Id3MetaVerifyPRORequest) GetFacePicture() *string {
	return s.FacePicture
}

func (s *Id3MetaVerifyPRORequest) GetFaceUrl() *string {
	return s.FaceUrl
}

func (s *Id3MetaVerifyPRORequest) GetIdentifyNum() *string {
	return s.IdentifyNum
}

func (s *Id3MetaVerifyPRORequest) GetLivenessCheck() *string {
	return s.LivenessCheck
}

func (s *Id3MetaVerifyPRORequest) GetParamType() *string {
	return s.ParamType
}

func (s *Id3MetaVerifyPRORequest) GetUserName() *string {
	return s.UserName
}

func (s *Id3MetaVerifyPRORequest) SetCrop(v string) *Id3MetaVerifyPRORequest {
	s.Crop = &v
	return s
}

func (s *Id3MetaVerifyPRORequest) SetEnableFallback(v string) *Id3MetaVerifyPRORequest {
	s.EnableFallback = &v
	return s
}

func (s *Id3MetaVerifyPRORequest) SetFaceFile(v string) *Id3MetaVerifyPRORequest {
	s.FaceFile = &v
	return s
}

func (s *Id3MetaVerifyPRORequest) SetFacePicture(v string) *Id3MetaVerifyPRORequest {
	s.FacePicture = &v
	return s
}

func (s *Id3MetaVerifyPRORequest) SetFaceUrl(v string) *Id3MetaVerifyPRORequest {
	s.FaceUrl = &v
	return s
}

func (s *Id3MetaVerifyPRORequest) SetIdentifyNum(v string) *Id3MetaVerifyPRORequest {
	s.IdentifyNum = &v
	return s
}

func (s *Id3MetaVerifyPRORequest) SetLivenessCheck(v string) *Id3MetaVerifyPRORequest {
	s.LivenessCheck = &v
	return s
}

func (s *Id3MetaVerifyPRORequest) SetParamType(v string) *Id3MetaVerifyPRORequest {
	s.ParamType = &v
	return s
}

func (s *Id3MetaVerifyPRORequest) SetUserName(v string) *Id3MetaVerifyPRORequest {
	s.UserName = &v
	return s
}

func (s *Id3MetaVerifyPRORequest) Validate() error {
	return dara.Validate(s)
}
