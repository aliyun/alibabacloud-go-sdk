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
	SetFaceAttributeCheck(v string) *FaceVerifyIntlRequest
	GetFaceAttributeCheck() *string
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
	SetUpdateFaceIfUserExists(v string) *FaceVerifyIntlRequest
	GetUpdateFaceIfUserExists() *string
	SetVerifyModel(v string) *FaceVerifyIntlRequest
	GetVerifyModel() *string
}

type FaceVerifyIntlRequest struct {
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
	SourceFacePictureFile *string `json:"SourceFacePictureFile,omitempty" xml:"SourceFacePictureFile,omitempty"`
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
	TargetFacePictureFile *string `json:"TargetFacePictureFile,omitempty" xml:"TargetFacePictureFile,omitempty"`
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

func (s FaceVerifyIntlRequest) String() string {
	return dara.Prettify(s)
}

func (s FaceVerifyIntlRequest) GoString() string {
	return s.String()
}

func (s *FaceVerifyIntlRequest) GetAutoRegistration() *string {
	return s.AutoRegistration
}

func (s *FaceVerifyIntlRequest) GetFaceAttributeCheck() *string {
	return s.FaceAttributeCheck
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

func (s *FaceVerifyIntlRequest) GetUpdateFaceIfUserExists() *string {
	return s.UpdateFaceIfUserExists
}

func (s *FaceVerifyIntlRequest) GetVerifyModel() *string {
	return s.VerifyModel
}

func (s *FaceVerifyIntlRequest) SetAutoRegistration(v string) *FaceVerifyIntlRequest {
	s.AutoRegistration = &v
	return s
}

func (s *FaceVerifyIntlRequest) SetFaceAttributeCheck(v string) *FaceVerifyIntlRequest {
	s.FaceAttributeCheck = &v
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

func (s *FaceVerifyIntlRequest) SetUpdateFaceIfUserExists(v string) *FaceVerifyIntlRequest {
	s.UpdateFaceIfUserExists = &v
	return s
}

func (s *FaceVerifyIntlRequest) SetVerifyModel(v string) *FaceVerifyIntlRequest {
	s.VerifyModel = &v
	return s
}

func (s *FaceVerifyIntlRequest) Validate() error {
	return dara.Validate(s)
}
