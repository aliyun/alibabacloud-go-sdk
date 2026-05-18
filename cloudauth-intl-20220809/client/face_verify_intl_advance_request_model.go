// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
	"io"
)

type iFaceVerifyIntlAdvanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoRegistration(v string) *FaceVerifyIntlAdvanceRequest
	GetAutoRegistration() *string
	SetFaceGroupCodes(v string) *FaceVerifyIntlAdvanceRequest
	GetFaceGroupCodes() *string
	SetFaceQualityCheck(v string) *FaceVerifyIntlAdvanceRequest
	GetFaceQualityCheck() *string
	SetFaceRegisterGroupCode(v string) *FaceVerifyIntlAdvanceRequest
	GetFaceRegisterGroupCode() *string
	SetMerchantBizId(v string) *FaceVerifyIntlAdvanceRequest
	GetMerchantBizId() *string
	SetMerchantUserId(v string) *FaceVerifyIntlAdvanceRequest
	GetMerchantUserId() *string
	SetProductCode(v string) *FaceVerifyIntlAdvanceRequest
	GetProductCode() *string
	SetReturnFaces(v string) *FaceVerifyIntlAdvanceRequest
	GetReturnFaces() *string
	SetSourceFacePicture(v string) *FaceVerifyIntlAdvanceRequest
	GetSourceFacePicture() *string
	SetSourceFacePictureFileObject(v io.Reader) *FaceVerifyIntlAdvanceRequest
	GetSourceFacePictureFileObject() io.Reader
	SetSourceFacePictureUrl(v string) *FaceVerifyIntlAdvanceRequest
	GetSourceFacePictureUrl() *string
	SetTargetFacePicture(v string) *FaceVerifyIntlAdvanceRequest
	GetTargetFacePicture() *string
	SetTargetFacePictureFileObject(v io.Reader) *FaceVerifyIntlAdvanceRequest
	GetTargetFacePictureFileObject() io.Reader
	SetTargetFacePictureUrl(v string) *FaceVerifyIntlAdvanceRequest
	GetTargetFacePictureUrl() *string
	SetVerifyModel(v string) *FaceVerifyIntlAdvanceRequest
	GetVerifyModel() *string
}

type FaceVerifyIntlAdvanceRequest struct {
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
	SourceFacePictureFileObject io.Reader `json:"SourceFacePictureFile,omitempty" xml:"SourceFacePictureFile,omitempty"`
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
	TargetFacePictureFileObject io.Reader `json:"TargetFacePictureFile,omitempty" xml:"TargetFacePictureFile,omitempty"`
	// example:
	//
	// https://***face2.jpeg
	TargetFacePictureUrl *string `json:"TargetFacePictureUrl,omitempty" xml:"TargetFacePictureUrl,omitempty"`
	// example:
	//
	// 0
	VerifyModel *string `json:"VerifyModel,omitempty" xml:"VerifyModel,omitempty"`
}

func (s FaceVerifyIntlAdvanceRequest) String() string {
	return dara.Prettify(s)
}

func (s FaceVerifyIntlAdvanceRequest) GoString() string {
	return s.String()
}

func (s *FaceVerifyIntlAdvanceRequest) GetAutoRegistration() *string {
	return s.AutoRegistration
}

func (s *FaceVerifyIntlAdvanceRequest) GetFaceGroupCodes() *string {
	return s.FaceGroupCodes
}

func (s *FaceVerifyIntlAdvanceRequest) GetFaceQualityCheck() *string {
	return s.FaceQualityCheck
}

func (s *FaceVerifyIntlAdvanceRequest) GetFaceRegisterGroupCode() *string {
	return s.FaceRegisterGroupCode
}

func (s *FaceVerifyIntlAdvanceRequest) GetMerchantBizId() *string {
	return s.MerchantBizId
}

func (s *FaceVerifyIntlAdvanceRequest) GetMerchantUserId() *string {
	return s.MerchantUserId
}

func (s *FaceVerifyIntlAdvanceRequest) GetProductCode() *string {
	return s.ProductCode
}

func (s *FaceVerifyIntlAdvanceRequest) GetReturnFaces() *string {
	return s.ReturnFaces
}

func (s *FaceVerifyIntlAdvanceRequest) GetSourceFacePicture() *string {
	return s.SourceFacePicture
}

func (s *FaceVerifyIntlAdvanceRequest) GetSourceFacePictureFileObject() io.Reader {
	return s.SourceFacePictureFileObject
}

func (s *FaceVerifyIntlAdvanceRequest) GetSourceFacePictureUrl() *string {
	return s.SourceFacePictureUrl
}

func (s *FaceVerifyIntlAdvanceRequest) GetTargetFacePicture() *string {
	return s.TargetFacePicture
}

func (s *FaceVerifyIntlAdvanceRequest) GetTargetFacePictureFileObject() io.Reader {
	return s.TargetFacePictureFileObject
}

func (s *FaceVerifyIntlAdvanceRequest) GetTargetFacePictureUrl() *string {
	return s.TargetFacePictureUrl
}

func (s *FaceVerifyIntlAdvanceRequest) GetVerifyModel() *string {
	return s.VerifyModel
}

func (s *FaceVerifyIntlAdvanceRequest) SetAutoRegistration(v string) *FaceVerifyIntlAdvanceRequest {
	s.AutoRegistration = &v
	return s
}

func (s *FaceVerifyIntlAdvanceRequest) SetFaceGroupCodes(v string) *FaceVerifyIntlAdvanceRequest {
	s.FaceGroupCodes = &v
	return s
}

func (s *FaceVerifyIntlAdvanceRequest) SetFaceQualityCheck(v string) *FaceVerifyIntlAdvanceRequest {
	s.FaceQualityCheck = &v
	return s
}

func (s *FaceVerifyIntlAdvanceRequest) SetFaceRegisterGroupCode(v string) *FaceVerifyIntlAdvanceRequest {
	s.FaceRegisterGroupCode = &v
	return s
}

func (s *FaceVerifyIntlAdvanceRequest) SetMerchantBizId(v string) *FaceVerifyIntlAdvanceRequest {
	s.MerchantBizId = &v
	return s
}

func (s *FaceVerifyIntlAdvanceRequest) SetMerchantUserId(v string) *FaceVerifyIntlAdvanceRequest {
	s.MerchantUserId = &v
	return s
}

func (s *FaceVerifyIntlAdvanceRequest) SetProductCode(v string) *FaceVerifyIntlAdvanceRequest {
	s.ProductCode = &v
	return s
}

func (s *FaceVerifyIntlAdvanceRequest) SetReturnFaces(v string) *FaceVerifyIntlAdvanceRequest {
	s.ReturnFaces = &v
	return s
}

func (s *FaceVerifyIntlAdvanceRequest) SetSourceFacePicture(v string) *FaceVerifyIntlAdvanceRequest {
	s.SourceFacePicture = &v
	return s
}

func (s *FaceVerifyIntlAdvanceRequest) SetSourceFacePictureFileObject(v io.Reader) *FaceVerifyIntlAdvanceRequest {
	s.SourceFacePictureFileObject = v
	return s
}

func (s *FaceVerifyIntlAdvanceRequest) SetSourceFacePictureUrl(v string) *FaceVerifyIntlAdvanceRequest {
	s.SourceFacePictureUrl = &v
	return s
}

func (s *FaceVerifyIntlAdvanceRequest) SetTargetFacePicture(v string) *FaceVerifyIntlAdvanceRequest {
	s.TargetFacePicture = &v
	return s
}

func (s *FaceVerifyIntlAdvanceRequest) SetTargetFacePictureFileObject(v io.Reader) *FaceVerifyIntlAdvanceRequest {
	s.TargetFacePictureFileObject = v
	return s
}

func (s *FaceVerifyIntlAdvanceRequest) SetTargetFacePictureUrl(v string) *FaceVerifyIntlAdvanceRequest {
	s.TargetFacePictureUrl = &v
	return s
}

func (s *FaceVerifyIntlAdvanceRequest) SetVerifyModel(v string) *FaceVerifyIntlAdvanceRequest {
	s.VerifyModel = &v
	return s
}

func (s *FaceVerifyIntlAdvanceRequest) Validate() error {
	return dara.Validate(s)
}
