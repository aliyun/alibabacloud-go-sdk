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
	SetFaceAttributeCheck(v string) *FaceVerifyIntlAdvanceRequest
	GetFaceAttributeCheck() *string
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
	SetUpdateFaceIfUserExists(v string) *FaceVerifyIntlAdvanceRequest
	GetUpdateFaceIfUserExists() *string
	SetVerifyModel(v string) *FaceVerifyIntlAdvanceRequest
	GetVerifyModel() *string
}

type FaceVerifyIntlAdvanceRequest struct {
	// Required when ProductCode is set to FACE_IDU_MIN.
	//
	// Specifies whether to automatically register the face to the specified face library when no duplicate face is retrieved. Valid values:
	//
	// - 0: automatically register.
	//
	// - 1: do not register. This is the default value.
	//
	// example:
	//
	// 1
	AutoRegistration *string `json:"AutoRegistration,omitempty" xml:"AutoRegistration,omitempty"`
	// Specifies whether to return additional face attribute information. Valid values:
	//
	// - N: do not return. This is the default value.
	//
	// - Y: return. The result is included in Result.ExtFaceInfo.FaceAttributeInfo.
	//
	// example:
	//
	// N
	FaceAttributeCheck *string `json:"FaceAttributeCheck,omitempty" xml:"FaceAttributeCheck,omitempty"`
	// Required when ProductCode is set to FACE_IDU_MIN.
	//
	// The codes of the face libraries created by the customer in the console. You can query up to 10 face libraries at the same time. Separate multiple face library codes with commas.
	//
	// example:
	//
	// 1232344，23444
	FaceGroupCodes *string `json:"FaceGroupCodes,omitempty" xml:"FaceGroupCodes,omitempty"`
	// Specifies whether to check the quality of the face image. Valid values:
	//
	// - Y: enabled.
	//
	// - N: disabled. This is the default value.
	//
	// example:
	//
	// Y
	FaceQualityCheck *string `json:"FaceQualityCheck,omitempty" xml:"FaceQualityCheck,omitempty"`
	// Required when ProductCode is set to FACE_IDU_MIN.
	//
	// The code of the face library for face registration.
	//
	// example:
	//
	// 0e0c34a77f
	FaceRegisterGroupCode *string `json:"FaceRegisterGroupCode,omitempty" xml:"FaceRegisterGroupCode,omitempty"`
	// The custom unique business identifier used for subsequent troubleshooting. The value must be a unique combination of up to 32 letters and digits.
	//
	// example:
	//
	// e0c34a77f5ac40a5aa5e6ed20c35****
	MerchantBizId *string `json:"MerchantBizId,omitempty" xml:"MerchantBizId,omitempty"`
	// The custom user ID or another identifier that can identify a specific user, such as a mobile phone number or email address. Desensitize the value of this field in advance, such as by hashing the value.
	//
	// example:
	//
	// 123456789
	MerchantUserId *string `json:"MerchantUserId,omitempty" xml:"MerchantUserId,omitempty"`
	// The product code. Valid values: FACE_VERIFY_MIN and FACE_IDU_MIN.
	//
	// This parameter is required.
	//
	// example:
	//
	// FACE_VERIFY_MIN
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// Required when ProductCode is set to FACE_IDU_MIN.
	//
	// The custom number of faces to return when multiple faces exceed the matching threshold. Default value: 1. Maximum value: 5.
	//
	// example:
	//
	// 1
	ReturnFaces *string `json:"ReturnFaces,omitempty" xml:"ReturnFaces,omitempty"`
	// The Base64-encoded portrait photo.
	//
	// Note
	//
	// - If you use this method to pass the portrait photo, check the photo size and do not pass oversized photos.
	//
	// - Specify only one of the following parameters: SourceFacePicture, SourceFacePictureUrl, and SourceFacePictureFile.
	//
	// example:
	//
	// base64
	SourceFacePicture *string `json:"SourceFacePicture,omitempty" xml:"SourceFacePicture,omitempty"`
	// The file stream of the face image.
	//
	// example:
	//
	// InputStream
	SourceFacePictureFileObject io.Reader `json:"SourceFacePictureFile,omitempty" xml:"SourceFacePictureFile,omitempty"`
	// The URL of the portrait image. The URL must be a publicly accessible HTTPS link.
	//
	// example:
	//
	// https://***face1.jpeg
	SourceFacePictureUrl *string `json:"SourceFacePictureUrl,omitempty" xml:"SourceFacePictureUrl,omitempty"`
	// The Base64-encoded reference photo.
	//
	// Note
	//
	// - If you use this method to pass the reference photo, check the photo size and do not pass oversized photos.
	//
	// - Specify only one of the following parameters: TargetFacePicture, TargetFacePictureUrl, and TargetFacePictureFile.
	//
	// example:
	//
	// base64
	TargetFacePicture *string `json:"TargetFacePicture,omitempty" xml:"TargetFacePicture,omitempty"`
	// The file stream of the reference face image.
	//
	// example:
	//
	// InputStream
	TargetFacePictureFileObject io.Reader `json:"TargetFacePictureFile,omitempty" xml:"TargetFacePictureFile,omitempty"`
	// The HTTPS URL of the reference face image.
	//
	// example:
	//
	// https://***face2.jpeg
	TargetFacePictureUrl *string `json:"TargetFacePictureUrl,omitempty" xml:"TargetFacePictureUrl,omitempty"`
	// Specifies whether to overwrite the existing face with the current face when the MerchantUserId already exists during automatic registration. Valid values:
	//
	// - Y: overwrite.
	//
	// - N: do not overwrite and return that the UserId already exists.
	//
	// example:
	//
	// N
	UpdateFaceIfUserExists *string `json:"UpdateFaceIfUserExists,omitempty" xml:"UpdateFaceIfUserExists,omitempty"`
	// Required when ProductCode is set to FACE_IDU_MIN.
	//
	// The verification type. Valid values:
	//
	// - 0: retrieval mode.
	//
	// > - Feature: Passes the face library and the user face image (SourceFacePicture). The system automatically retrieves whether the specified face image (SourceFacePicture) already exists in the face library. Passive liveness detection can be enabled for the face image (SourceFacePicture).
	//
	// > - Recommended scenario: Scenarios where real persons create an account and duplicate registrations are not allowed.
	//
	// - 1: verification mode. This is the default value.
	//
	// > - Feature: Passes the specified face image (SourceFacePicture) and the reference face image (TargetFacePicture). The system automatically verifies whether the faces match. Passive liveness detection can be enabled for the specified face image (SourceFacePicture).
	//
	// > - Recommended scenario: Scenarios where you need to verify whether the operation is performed by the user in person when modifying logon or account information.
	//
	// - 2: comprehensive mode.
	//
	// > - Feature: Passes the face library, the specified face image (SourceFacePicture), and the reference face image (TargetFacePicture) at the same time. The system automatically retrieves whether the specified face image (SourceFacePicture) exists in the face library and whether it matches the reference face. Passive liveness detection can be enabled for the specified face image (SourceFacePicture).
	//
	// > - Recommended scenario: Scenarios where you need to verify that the user is a new user and the operation is performed in person.
	//
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

func (s *FaceVerifyIntlAdvanceRequest) GetFaceAttributeCheck() *string {
	return s.FaceAttributeCheck
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

func (s *FaceVerifyIntlAdvanceRequest) GetUpdateFaceIfUserExists() *string {
	return s.UpdateFaceIfUserExists
}

func (s *FaceVerifyIntlAdvanceRequest) GetVerifyModel() *string {
	return s.VerifyModel
}

func (s *FaceVerifyIntlAdvanceRequest) SetAutoRegistration(v string) *FaceVerifyIntlAdvanceRequest {
	s.AutoRegistration = &v
	return s
}

func (s *FaceVerifyIntlAdvanceRequest) SetFaceAttributeCheck(v string) *FaceVerifyIntlAdvanceRequest {
	s.FaceAttributeCheck = &v
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

func (s *FaceVerifyIntlAdvanceRequest) SetUpdateFaceIfUserExists(v string) *FaceVerifyIntlAdvanceRequest {
	s.UpdateFaceIfUserExists = &v
	return s
}

func (s *FaceVerifyIntlAdvanceRequest) SetVerifyModel(v string) *FaceVerifyIntlAdvanceRequest {
	s.VerifyModel = &v
	return s
}

func (s *FaceVerifyIntlAdvanceRequest) Validate() error {
	return dara.Validate(s)
}
