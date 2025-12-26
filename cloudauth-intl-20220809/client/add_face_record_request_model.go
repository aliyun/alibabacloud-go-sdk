// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddFaceRecordRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFaceGroupCode(v string) *AddFaceRecordRequest
	GetFaceGroupCode() *string
	SetFacePicture(v string) *AddFaceRecordRequest
	GetFacePicture() *string
	SetFacePictureFile(v string) *AddFaceRecordRequest
	GetFacePictureFile() *string
	SetFacePictureUrl(v string) *AddFaceRecordRequest
	GetFacePictureUrl() *string
	SetFaceQualityCheck(v string) *AddFaceRecordRequest
	GetFaceQualityCheck() *string
	SetMerchantUserId(v string) *AddFaceRecordRequest
	GetMerchantUserId() *string
	SetProductCode(v string) *AddFaceRecordRequest
	GetProductCode() *string
}

type AddFaceRecordRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// sgl****7uc
	FaceGroupCode *string `json:"FaceGroupCode,omitempty" xml:"FaceGroupCode,omitempty"`
	// example:
	//
	// base64
	FacePicture *string `json:"FacePicture,omitempty" xml:"FacePicture,omitempty"`
	// example:
	//
	// InputStream
	FacePictureFile *string `json:"FacePictureFile,omitempty" xml:"FacePictureFile,omitempty"`
	// example:
	//
	// https://www.xxxxx.com/test.jpg
	FacePictureUrl *string `json:"FacePictureUrl,omitempty" xml:"FacePictureUrl,omitempty"`
	// example:
	//
	// N
	FaceQualityCheck *string `json:"FaceQualityCheck,omitempty" xml:"FaceQualityCheck,omitempty"`
	// example:
	//
	// 130A2C10B9EE4D8488E35384FF03hst
	MerchantUserId *string `json:"MerchantUserId,omitempty" xml:"MerchantUserId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// FACE_ENROLL
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
}

func (s AddFaceRecordRequest) String() string {
	return dara.Prettify(s)
}

func (s AddFaceRecordRequest) GoString() string {
	return s.String()
}

func (s *AddFaceRecordRequest) GetFaceGroupCode() *string {
	return s.FaceGroupCode
}

func (s *AddFaceRecordRequest) GetFacePicture() *string {
	return s.FacePicture
}

func (s *AddFaceRecordRequest) GetFacePictureFile() *string {
	return s.FacePictureFile
}

func (s *AddFaceRecordRequest) GetFacePictureUrl() *string {
	return s.FacePictureUrl
}

func (s *AddFaceRecordRequest) GetFaceQualityCheck() *string {
	return s.FaceQualityCheck
}

func (s *AddFaceRecordRequest) GetMerchantUserId() *string {
	return s.MerchantUserId
}

func (s *AddFaceRecordRequest) GetProductCode() *string {
	return s.ProductCode
}

func (s *AddFaceRecordRequest) SetFaceGroupCode(v string) *AddFaceRecordRequest {
	s.FaceGroupCode = &v
	return s
}

func (s *AddFaceRecordRequest) SetFacePicture(v string) *AddFaceRecordRequest {
	s.FacePicture = &v
	return s
}

func (s *AddFaceRecordRequest) SetFacePictureFile(v string) *AddFaceRecordRequest {
	s.FacePictureFile = &v
	return s
}

func (s *AddFaceRecordRequest) SetFacePictureUrl(v string) *AddFaceRecordRequest {
	s.FacePictureUrl = &v
	return s
}

func (s *AddFaceRecordRequest) SetFaceQualityCheck(v string) *AddFaceRecordRequest {
	s.FaceQualityCheck = &v
	return s
}

func (s *AddFaceRecordRequest) SetMerchantUserId(v string) *AddFaceRecordRequest {
	s.MerchantUserId = &v
	return s
}

func (s *AddFaceRecordRequest) SetProductCode(v string) *AddFaceRecordRequest {
	s.ProductCode = &v
	return s
}

func (s *AddFaceRecordRequest) Validate() error {
	return dara.Validate(s)
}
