// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFaceVerifyIntlRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoRegistration(v string) *FaceVerifyIntlRequest
	GetAutoRegistration() *string
	SetFaceGroupCodes(v string) *FaceVerifyIntlRequest
	GetFaceGroupCodes() *string
	SetFaceQualityCheck(v string) *FaceVerifyIntlRequest
	GetFaceQualityCheck() *string
	SetFaceRegisterGroupCode(v string) *FaceVerifyIntlRequest
	GetFaceRegisterGroupCode() *string
	SetMerchantBizId(v string) *FaceVerifyIntlRequest
	GetMerchantBizId() *string
	SetMerchantUserId(v string) *FaceVerifyIntlRequest
	GetMerchantUserId() *string
	SetProductCode(v string) *FaceVerifyIntlRequest
	GetProductCode() *string
	SetReturnFaces(v string) *FaceVerifyIntlRequest
	GetReturnFaces() *string
	SetSourceFacePicture(v string) *FaceVerifyIntlRequest
	GetSourceFacePicture() *string
	SetSourceFacePictureFile(v string) *FaceVerifyIntlRequest
	GetSourceFacePictureFile() *string
	SetSourceFacePictureUrl(v string) *FaceVerifyIntlRequest
	GetSourceFacePictureUrl() *string
	SetTargetFacePicture(v string) *FaceVerifyIntlRequest
	GetTargetFacePicture() *string
	SetTargetFacePictureFile(v string) *FaceVerifyIntlRequest
	GetTargetFacePictureFile() *string
	SetTargetFacePictureUrl(v string) *FaceVerifyIntlRequest
	GetTargetFacePictureUrl() *string
	SetVerifyModel(v string) *FaceVerifyIntlRequest
	GetVerifyModel() *string
}

type FaceVerifyIntlRequest struct {
	// example:
	//
	// 1
	AutoRegistration *string `json:"AutoRegistration,omitempty" xml:"AutoRegistration,omitempty"`
	// example:
	//
	// 1232344，23444
	FaceGroupCodes *string `json:"FaceGroupCodes,omitempty" xml:"FaceGroupCodes,omitempty"`
	// example:
	//
	// Y
	FaceQualityCheck *string `json:"FaceQualityCheck,omitempty" xml:"FaceQualityCheck,omitempty"`
	// example:
	//
	// 0e0c34a77f
	FaceRegisterGroupCode *string `json:"FaceRegisterGroupCode,omitempty" xml:"FaceRegisterGroupCode,omitempty"`
	// example:
	//
	// e0c34a77f5ac40a5aa5e6ed20c35****
	MerchantBizId *string `json:"MerchantBizId,omitempty" xml:"MerchantBizId,omitempty"`
	// example:
	//
	// 123456789
	MerchantUserId *string `json:"MerchantUserId,omitempty" xml:"MerchantUserId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// FACE_VERIFY_MIN
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// example:
	//
	// 1
	ReturnFaces *string `json:"ReturnFaces,omitempty" xml:"ReturnFaces,omitempty"`
	// example:
	//
	// base64
	SourceFacePicture *string `json:"SourceFacePicture,omitempty" xml:"SourceFacePicture,omitempty"`
	// example:
	//
	// InputStream
	SourceFacePictureFile *string `json:"SourceFacePictureFile,omitempty" xml:"SourceFacePictureFile,omitempty"`
	// example:
	//
	// https://***face1.jpeg
	SourceFacePictureUrl *string `json:"SourceFacePictureUrl,omitempty" xml:"SourceFacePictureUrl,omitempty"`
	// example:
	//
	// base64
	TargetFacePicture *string `json:"TargetFacePicture,omitempty" xml:"TargetFacePicture,omitempty"`
	// example:
	//
	// InputStream
	TargetFacePictureFile *string `json:"TargetFacePictureFile,omitempty" xml:"TargetFacePictureFile,omitempty"`
	// example:
	//
	// https://***face2.jpeg
	TargetFacePictureUrl *string `json:"TargetFacePictureUrl,omitempty" xml:"TargetFacePictureUrl,omitempty"`
	// example:
	//
	// 0
	VerifyModel *string `json:"VerifyModel,omitempty" xml:"VerifyModel,omitempty"`
}

func (s FaceVerifyIntlRequest) String() string {
	return dara.Prettify(s)
}

func (s FaceVerifyIntlRequest) GoString() string {
	return s.String()
}

func (s *FaceVerifyIntlRequest) GetAutoRegistration() *string {
	return s.AutoRegistration
}

func (s *FaceVerifyIntlRequest) GetFaceGroupCodes() *string {
	return s.FaceGroupCodes
}

func (s *FaceVerifyIntlRequest) GetFaceQualityCheck() *string {
	return s.FaceQualityCheck
}

func (s *FaceVerifyIntlRequest) GetFaceRegisterGroupCode() *string {
	return s.FaceRegisterGroupCode
}

func (s *FaceVerifyIntlRequest) GetMerchantBizId() *string {
	return s.MerchantBizId
}

func (s *FaceVerifyIntlRequest) GetMerchantUserId() *string {
	return s.MerchantUserId
}

func (s *FaceVerifyIntlRequest) GetProductCode() *string {
	return s.ProductCode
}

func (s *FaceVerifyIntlRequest) GetReturnFaces() *string {
	return s.ReturnFaces
}

func (s *FaceVerifyIntlRequest) GetSourceFacePicture() *string {
	return s.SourceFacePicture
}

func (s *FaceVerifyIntlRequest) GetSourceFacePictureFile() *string {
	return s.SourceFacePictureFile
}

func (s *FaceVerifyIntlRequest) GetSourceFacePictureUrl() *string {
	return s.SourceFacePictureUrl
}

func (s *FaceVerifyIntlRequest) GetTargetFacePicture() *string {
	return s.TargetFacePicture
}

func (s *FaceVerifyIntlRequest) GetTargetFacePictureFile() *string {
	return s.TargetFacePictureFile
}

func (s *FaceVerifyIntlRequest) GetTargetFacePictureUrl() *string {
	return s.TargetFacePictureUrl
}

func (s *FaceVerifyIntlRequest) GetVerifyModel() *string {
	return s.VerifyModel
}

func (s *FaceVerifyIntlRequest) SetAutoRegistration(v string) *FaceVerifyIntlRequest {
	s.AutoRegistration = &v
	return s
}

func (s *FaceVerifyIntlRequest) SetFaceGroupCodes(v string) *FaceVerifyIntlRequest {
	s.FaceGroupCodes = &v
	return s
}

func (s *FaceVerifyIntlRequest) SetFaceQualityCheck(v string) *FaceVerifyIntlRequest {
	s.FaceQualityCheck = &v
	return s
}

func (s *FaceVerifyIntlRequest) SetFaceRegisterGroupCode(v string) *FaceVerifyIntlRequest {
	s.FaceRegisterGroupCode = &v
	return s
}

func (s *FaceVerifyIntlRequest) SetMerchantBizId(v string) *FaceVerifyIntlRequest {
	s.MerchantBizId = &v
	return s
}

func (s *FaceVerifyIntlRequest) SetMerchantUserId(v string) *FaceVerifyIntlRequest {
	s.MerchantUserId = &v
	return s
}

func (s *FaceVerifyIntlRequest) SetProductCode(v string) *FaceVerifyIntlRequest {
	s.ProductCode = &v
	return s
}

func (s *FaceVerifyIntlRequest) SetReturnFaces(v string) *FaceVerifyIntlRequest {
	s.ReturnFaces = &v
	return s
}

func (s *FaceVerifyIntlRequest) SetSourceFacePicture(v string) *FaceVerifyIntlRequest {
	s.SourceFacePicture = &v
	return s
}

func (s *FaceVerifyIntlRequest) SetSourceFacePictureFile(v string) *FaceVerifyIntlRequest {
	s.SourceFacePictureFile = &v
	return s
}

func (s *FaceVerifyIntlRequest) SetSourceFacePictureUrl(v string) *FaceVerifyIntlRequest {
	s.SourceFacePictureUrl = &v
	return s
}

func (s *FaceVerifyIntlRequest) SetTargetFacePicture(v string) *FaceVerifyIntlRequest {
	s.TargetFacePicture = &v
	return s
}

func (s *FaceVerifyIntlRequest) SetTargetFacePictureFile(v string) *FaceVerifyIntlRequest {
	s.TargetFacePictureFile = &v
	return s
}

func (s *FaceVerifyIntlRequest) SetTargetFacePictureUrl(v string) *FaceVerifyIntlRequest {
	s.TargetFacePictureUrl = &v
	return s
}

func (s *FaceVerifyIntlRequest) SetVerifyModel(v string) *FaceVerifyIntlRequest {
	s.VerifyModel = &v
	return s
}

func (s *FaceVerifyIntlRequest) Validate() error {
	return dara.Validate(s)
}
