// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFaceDuplicationCheckIntlRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoRegistration(v string) *FaceDuplicationCheckIntlRequest
	GetAutoRegistration() *string
	SetFaceAttributeCheck(v string) *FaceDuplicationCheckIntlRequest
	GetFaceAttributeCheck() *string
	SetFaceGroupCodes(v string) *FaceDuplicationCheckIntlRequest
	GetFaceGroupCodes() *string
	SetFaceQualityCheck(v string) *FaceDuplicationCheckIntlRequest
	GetFaceQualityCheck() *string
	SetFaceRegisterGroupCode(v string) *FaceDuplicationCheckIntlRequest
	GetFaceRegisterGroupCode() *string
	SetFaceVerifyThreshold(v string) *FaceDuplicationCheckIntlRequest
	GetFaceVerifyThreshold() *string
	SetLiveness(v string) *FaceDuplicationCheckIntlRequest
	GetLiveness() *string
	SetMerchantBizId(v string) *FaceDuplicationCheckIntlRequest
	GetMerchantBizId() *string
	SetMerchantUserId(v string) *FaceDuplicationCheckIntlRequest
	GetMerchantUserId() *string
	SetProductCode(v string) *FaceDuplicationCheckIntlRequest
	GetProductCode() *string
	SetReturnFaces(v string) *FaceDuplicationCheckIntlRequest
	GetReturnFaces() *string
	SetSaveFacePicture(v string) *FaceDuplicationCheckIntlRequest
	GetSaveFacePicture() *string
	SetSceneCode(v string) *FaceDuplicationCheckIntlRequest
	GetSceneCode() *string
	SetSourceFacePicture(v string) *FaceDuplicationCheckIntlRequest
	GetSourceFacePicture() *string
	SetSourceFacePictureUrl(v string) *FaceDuplicationCheckIntlRequest
	GetSourceFacePictureUrl() *string
	SetTargetFacePicture(v string) *FaceDuplicationCheckIntlRequest
	GetTargetFacePicture() *string
	SetTargetFacePictureUrl(v string) *FaceDuplicationCheckIntlRequest
	GetTargetFacePictureUrl() *string
	SetUpdateFaceIfUserExists(v string) *FaceDuplicationCheckIntlRequest
	GetUpdateFaceIfUserExists() *string
	SetVerifyModel(v string) *FaceDuplicationCheckIntlRequest
	GetVerifyModel() *string
}

type FaceDuplicationCheckIntlRequest struct {
	// Specifies whether to automatically register the face to the specified face library if no duplicate face is found during the search. Valid values:
	//
	// - 0: Register.
	//
	// - 1: Do not register. This is the default value.
	//
	// example:
	//
	// 1
	AutoRegistration *string `json:"AutoRegistration,omitempty" xml:"AutoRegistration,omitempty"`
	// Specifies whether to return additional face attribute information. Valid values:
	//
	// - N: Do not return. This is the default value.
	//
	// - Y: Return. The result is located in Result.ExtFaceInfo.FaceAttributeInfo.
	//
	// example:
	//
	// N
	FaceAttributeCheck *string `json:"FaceAttributeCheck,omitempty" xml:"FaceAttributeCheck,omitempty"`
	// The codes of the face libraries created by the customer in the console. A maximum of 10 face libraries can be queried at the same time. Separate multiple face library codes with commas.
	//
	// example:
	//
	// 1232344，23444
	FaceGroupCodes *string `json:"FaceGroupCodes,omitempty" xml:"FaceGroupCodes,omitempty"`
	// The face quality check.
	//
	// example:
	//
	// Y
	FaceQualityCheck *string `json:"FaceQualityCheck,omitempty" xml:"FaceQualityCheck,omitempty"`
	// The code of the face library for face registration.
	//
	// example:
	//
	// 0e0c34a77f
	FaceRegisterGroupCode *string `json:"FaceRegisterGroupCode,omitempty" xml:"FaceRegisterGroupCode,omitempty"`
	// The face matching threshold. <warning>Reserved field. Not enabled yet.</warning>
	//
	// example:
	//
	// 0.5
	FaceVerifyThreshold *string `json:"FaceVerifyThreshold,omitempty" xml:"FaceVerifyThreshold,omitempty"`
	// Specifies whether to enable passive liveness detection. Valid values:
	//
	// - 0: Disable.
	//
	// - 1: Enable.
	//
	// example:
	//
	// 0
	Liveness *string `json:"Liveness,omitempty" xml:"Liveness,omitempty"`
	// The custom unique business identifier used for subsequent troubleshooting. The value supports a combination of up to 32 letters and digits. Make sure that the value is unique.
	//
	// This parameter is required.
	//
	// example:
	//
	// e0c34a77f5ac40a5aa5e6ed20c35****
	MerchantBizId *string `json:"MerchantBizId,omitempty" xml:"MerchantBizId,omitempty"`
	// The custom user ID or other identifier that can identify a specific user, such as a mobile phone number or email address. Desensitize the value of this field in advance, such as by hashing the value.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234567890
	MerchantUserId *string `json:"MerchantUserId,omitempty" xml:"MerchantUserId,omitempty"`
	// The product code.
	//
	// This parameter is required.
	//
	// example:
	//
	// FACE_IDU_MIN
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// Specifies the number of faces to return when multiple faces exceed the matching threshold. Default value: 1. Maximum value: 5.
	//
	// example:
	//
	// 1
	ReturnFaces *string `json:"ReturnFaces,omitempty" xml:"ReturnFaces,omitempty"`
	// Specifies the type of face data to save. Valid values:
	//
	// - 0: Face image. This is the default value.
	//
	// - 1: Face feature.
	//
	// <warning>Reserved field. Not enabled yet.</warning>
	//
	// example:
	//
	// 0
	SaveFacePicture *string `json:"SaveFacePicture,omitempty" xml:"SaveFacePicture,omitempty"`
	// The custom authentication scene ID.
	//
	// example:
	//
	// 1234567890
	SceneCode *string `json:"SceneCode,omitempty" xml:"SceneCode,omitempty"`
	// The Base64-encoded portrait photo.
	//
	// example:
	//
	// base64
	SourceFacePicture *string `json:"SourceFacePicture,omitempty" xml:"SourceFacePicture,omitempty"`
	// The URL of the portrait image. The URL must be a publicly accessible HTTP or HTTPS link.
	//
	// example:
	//
	// https://***face1.jpeg
	SourceFacePictureUrl *string `json:"SourceFacePictureUrl,omitempty" xml:"SourceFacePictureUrl,omitempty"`
	// The Base64-encoded portrait photo.
	//
	// example:
	//
	// base64
	TargetFacePicture *string `json:"TargetFacePicture,omitempty" xml:"TargetFacePicture,omitempty"`
	// The URL of the portrait image. The URL must be a publicly accessible HTTP or HTTPS link.
	//
	// example:
	//
	// https://***face2.jpeg
	TargetFacePictureUrl *string `json:"TargetFacePictureUrl,omitempty" xml:"TargetFacePictureUrl,omitempty"`
	// Specifies whether to overwrite the original face with the current face if the MerchantUserId already exists during automatic registration. Valid values:
	//
	// - Y: Overwrite.
	//
	// - N: Do not overwrite and return that the UserId already exists.
	//
	// example:
	//
	// N
	UpdateFaceIfUserExists *string `json:"UpdateFaceIfUserExists,omitempty" xml:"UpdateFaceIfUserExists,omitempty"`
	// The authentication type. Valid values:
	//
	// - 0: Retrieve pattern
	//
	//   - Feature: The system receives a face library and a user face image (sourceFacePicture), and automatically retrieves from the face library to check whether the specified face image (sourceFacePicture) already exists. Passive liveness detection can be enabled for the face image (sourceFacePicture).
	//
	//   - Recommended scenario: Scenarios where a real person needs to create an account and duplicate registration is not allowed.
	//
	// - 1 (default): Authenticate pattern
	//
	//   - Feature: The system receives a specified face image (sourceFacePicture) and a retained face image (TargetFacePicture), and automatically authenticates whether the faces match. Passive liveness detection can be enabled for the specified face image (sourceFacePicture).
	//
	//   - Recommended scenario: Scenarios where it is required to authenticate whether the operation is performed by the user when modifying logon or account information.
	//
	// - 2: Comprehensive pattern
	//
	//   - Feature: The system receives a face library, a specified face image (sourceFacePicture), and a retained face image (TargetFacePicture) at the same time, and automatically retrieves from the face library to check whether the specified face image (sourceFacePicture) exists and authenticates whether it matches the retained face. Passive liveness detection can be enabled for the specified face image (sourceFacePicture).
	//
	//   - Recommended scenario: Scenarios where it is required to authenticate that the user is a new user and the operation is performed by the user.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	VerifyModel *string `json:"VerifyModel,omitempty" xml:"VerifyModel,omitempty"`
}

func (s FaceDuplicationCheckIntlRequest) String() string {
	return dara.Prettify(s)
}

func (s FaceDuplicationCheckIntlRequest) GoString() string {
	return s.String()
}

func (s *FaceDuplicationCheckIntlRequest) GetAutoRegistration() *string {
	return s.AutoRegistration
}

func (s *FaceDuplicationCheckIntlRequest) GetFaceAttributeCheck() *string {
	return s.FaceAttributeCheck
}

func (s *FaceDuplicationCheckIntlRequest) GetFaceGroupCodes() *string {
	return s.FaceGroupCodes
}

func (s *FaceDuplicationCheckIntlRequest) GetFaceQualityCheck() *string {
	return s.FaceQualityCheck
}

func (s *FaceDuplicationCheckIntlRequest) GetFaceRegisterGroupCode() *string {
	return s.FaceRegisterGroupCode
}

func (s *FaceDuplicationCheckIntlRequest) GetFaceVerifyThreshold() *string {
	return s.FaceVerifyThreshold
}

func (s *FaceDuplicationCheckIntlRequest) GetLiveness() *string {
	return s.Liveness
}

func (s *FaceDuplicationCheckIntlRequest) GetMerchantBizId() *string {
	return s.MerchantBizId
}

func (s *FaceDuplicationCheckIntlRequest) GetMerchantUserId() *string {
	return s.MerchantUserId
}

func (s *FaceDuplicationCheckIntlRequest) GetProductCode() *string {
	return s.ProductCode
}

func (s *FaceDuplicationCheckIntlRequest) GetReturnFaces() *string {
	return s.ReturnFaces
}

func (s *FaceDuplicationCheckIntlRequest) GetSaveFacePicture() *string {
	return s.SaveFacePicture
}

func (s *FaceDuplicationCheckIntlRequest) GetSceneCode() *string {
	return s.SceneCode
}

func (s *FaceDuplicationCheckIntlRequest) GetSourceFacePicture() *string {
	return s.SourceFacePicture
}

func (s *FaceDuplicationCheckIntlRequest) GetSourceFacePictureUrl() *string {
	return s.SourceFacePictureUrl
}

func (s *FaceDuplicationCheckIntlRequest) GetTargetFacePicture() *string {
	return s.TargetFacePicture
}

func (s *FaceDuplicationCheckIntlRequest) GetTargetFacePictureUrl() *string {
	return s.TargetFacePictureUrl
}

func (s *FaceDuplicationCheckIntlRequest) GetUpdateFaceIfUserExists() *string {
	return s.UpdateFaceIfUserExists
}

func (s *FaceDuplicationCheckIntlRequest) GetVerifyModel() *string {
	return s.VerifyModel
}

func (s *FaceDuplicationCheckIntlRequest) SetAutoRegistration(v string) *FaceDuplicationCheckIntlRequest {
	s.AutoRegistration = &v
	return s
}

func (s *FaceDuplicationCheckIntlRequest) SetFaceAttributeCheck(v string) *FaceDuplicationCheckIntlRequest {
	s.FaceAttributeCheck = &v
	return s
}

func (s *FaceDuplicationCheckIntlRequest) SetFaceGroupCodes(v string) *FaceDuplicationCheckIntlRequest {
	s.FaceGroupCodes = &v
	return s
}

func (s *FaceDuplicationCheckIntlRequest) SetFaceQualityCheck(v string) *FaceDuplicationCheckIntlRequest {
	s.FaceQualityCheck = &v
	return s
}

func (s *FaceDuplicationCheckIntlRequest) SetFaceRegisterGroupCode(v string) *FaceDuplicationCheckIntlRequest {
	s.FaceRegisterGroupCode = &v
	return s
}

func (s *FaceDuplicationCheckIntlRequest) SetFaceVerifyThreshold(v string) *FaceDuplicationCheckIntlRequest {
	s.FaceVerifyThreshold = &v
	return s
}

func (s *FaceDuplicationCheckIntlRequest) SetLiveness(v string) *FaceDuplicationCheckIntlRequest {
	s.Liveness = &v
	return s
}

func (s *FaceDuplicationCheckIntlRequest) SetMerchantBizId(v string) *FaceDuplicationCheckIntlRequest {
	s.MerchantBizId = &v
	return s
}

func (s *FaceDuplicationCheckIntlRequest) SetMerchantUserId(v string) *FaceDuplicationCheckIntlRequest {
	s.MerchantUserId = &v
	return s
}

func (s *FaceDuplicationCheckIntlRequest) SetProductCode(v string) *FaceDuplicationCheckIntlRequest {
	s.ProductCode = &v
	return s
}

func (s *FaceDuplicationCheckIntlRequest) SetReturnFaces(v string) *FaceDuplicationCheckIntlRequest {
	s.ReturnFaces = &v
	return s
}

func (s *FaceDuplicationCheckIntlRequest) SetSaveFacePicture(v string) *FaceDuplicationCheckIntlRequest {
	s.SaveFacePicture = &v
	return s
}

func (s *FaceDuplicationCheckIntlRequest) SetSceneCode(v string) *FaceDuplicationCheckIntlRequest {
	s.SceneCode = &v
	return s
}

func (s *FaceDuplicationCheckIntlRequest) SetSourceFacePicture(v string) *FaceDuplicationCheckIntlRequest {
	s.SourceFacePicture = &v
	return s
}

func (s *FaceDuplicationCheckIntlRequest) SetSourceFacePictureUrl(v string) *FaceDuplicationCheckIntlRequest {
	s.SourceFacePictureUrl = &v
	return s
}

func (s *FaceDuplicationCheckIntlRequest) SetTargetFacePicture(v string) *FaceDuplicationCheckIntlRequest {
	s.TargetFacePicture = &v
	return s
}

func (s *FaceDuplicationCheckIntlRequest) SetTargetFacePictureUrl(v string) *FaceDuplicationCheckIntlRequest {
	s.TargetFacePictureUrl = &v
	return s
}

func (s *FaceDuplicationCheckIntlRequest) SetUpdateFaceIfUserExists(v string) *FaceDuplicationCheckIntlRequest {
	s.UpdateFaceIfUserExists = &v
	return s
}

func (s *FaceDuplicationCheckIntlRequest) SetVerifyModel(v string) *FaceDuplicationCheckIntlRequest {
	s.VerifyModel = &v
	return s
}

func (s *FaceDuplicationCheckIntlRequest) Validate() error {
	return dara.Validate(s)
}
