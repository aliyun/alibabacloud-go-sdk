// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
	"io"
)

type iAddFaceRecordAdvanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFaceGroupCode(v string) *AddFaceRecordAdvanceRequest
	GetFaceGroupCode() *string
	SetFacePicture(v string) *AddFaceRecordAdvanceRequest
	GetFacePicture() *string
	SetFacePictureFileObject(v io.Reader) *AddFaceRecordAdvanceRequest
	GetFacePictureFileObject() io.Reader
	SetFacePictureUrl(v string) *AddFaceRecordAdvanceRequest
	GetFacePictureUrl() *string
	SetFaceQualityCheck(v string) *AddFaceRecordAdvanceRequest
	GetFaceQualityCheck() *string
	SetMerchantUserId(v string) *AddFaceRecordAdvanceRequest
	GetMerchantUserId() *string
	SetProductCode(v string) *AddFaceRecordAdvanceRequest
	GetProductCode() *string
}

type AddFaceRecordAdvanceRequest struct {
	// The face library code.
	//
	// This parameter is required.
	//
	// example:
	//
	// sgl****7uc
	FaceGroupCode *string `json:"FaceGroupCode,omitempty" xml:"FaceGroupCode,omitempty"`
	// The Base64-encoded face image to register.
	//
	// example:
	//
	// base64
	FacePicture *string `json:"FacePicture,omitempty" xml:"FacePicture,omitempty"`
	// The file stream of the face image to register.
	//
	// example:
	//
	// InputStream
	FacePictureFileObject io.Reader `json:"FacePictureFile,omitempty" xml:"FacePictureFile,omitempty"`
	// The URL of the face image to register.
	//
	// example:
	//
	// https://www.xxxxx.com/test.jpg
	FacePictureUrl *string `json:"FacePictureUrl,omitempty" xml:"FacePictureUrl,omitempty"`
	// Specifies whether to check the quality of the face image. Valid values:
	//
	// - Y: enabled.
	//
	// - N: disabled (default).
	//
	// example:
	//
	// N
	FaceQualityCheck *string `json:"FaceQualityCheck,omitempty" xml:"FaceQualityCheck,omitempty"`
	// The custom unique user ID. The value cannot exceed 32 characters.
	//
	// - If this parameter is specified, the system registers the user with the specified MerchantUserId.
	//
	// - If this parameter is not specified, the image name is used by default.
	//
	// example:
	//
	// 130A2C10B9EE4D8488E35384FF03hst
	MerchantUserId *string `json:"MerchantUserId,omitempty" xml:"MerchantUserId,omitempty"`
	// The product code.
	//
	// This parameter is required.
	//
	// example:
	//
	// FACE_ENROLL
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
}

func (s AddFaceRecordAdvanceRequest) String() string {
	return dara.Prettify(s)
}

func (s AddFaceRecordAdvanceRequest) GoString() string {
	return s.String()
}

func (s *AddFaceRecordAdvanceRequest) GetFaceGroupCode() *string {
	return s.FaceGroupCode
}

func (s *AddFaceRecordAdvanceRequest) GetFacePicture() *string {
	return s.FacePicture
}

func (s *AddFaceRecordAdvanceRequest) GetFacePictureFileObject() io.Reader {
	return s.FacePictureFileObject
}

func (s *AddFaceRecordAdvanceRequest) GetFacePictureUrl() *string {
	return s.FacePictureUrl
}

func (s *AddFaceRecordAdvanceRequest) GetFaceQualityCheck() *string {
	return s.FaceQualityCheck
}

func (s *AddFaceRecordAdvanceRequest) GetMerchantUserId() *string {
	return s.MerchantUserId
}

func (s *AddFaceRecordAdvanceRequest) GetProductCode() *string {
	return s.ProductCode
}

func (s *AddFaceRecordAdvanceRequest) SetFaceGroupCode(v string) *AddFaceRecordAdvanceRequest {
	s.FaceGroupCode = &v
	return s
}

func (s *AddFaceRecordAdvanceRequest) SetFacePicture(v string) *AddFaceRecordAdvanceRequest {
	s.FacePicture = &v
	return s
}

func (s *AddFaceRecordAdvanceRequest) SetFacePictureFileObject(v io.Reader) *AddFaceRecordAdvanceRequest {
	s.FacePictureFileObject = v
	return s
}

func (s *AddFaceRecordAdvanceRequest) SetFacePictureUrl(v string) *AddFaceRecordAdvanceRequest {
	s.FacePictureUrl = &v
	return s
}

func (s *AddFaceRecordAdvanceRequest) SetFaceQualityCheck(v string) *AddFaceRecordAdvanceRequest {
	s.FaceQualityCheck = &v
	return s
}

func (s *AddFaceRecordAdvanceRequest) SetMerchantUserId(v string) *AddFaceRecordAdvanceRequest {
	s.MerchantUserId = &v
	return s
}

func (s *AddFaceRecordAdvanceRequest) SetProductCode(v string) *AddFaceRecordAdvanceRequest {
	s.ProductCode = &v
	return s
}

func (s *AddFaceRecordAdvanceRequest) Validate() error {
	return dara.Validate(s)
}
