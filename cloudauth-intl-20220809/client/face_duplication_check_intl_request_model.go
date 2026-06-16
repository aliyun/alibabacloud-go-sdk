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
	SetVerifyModel(v string) *FaceDuplicationCheckIntlRequest
	GetVerifyModel() *string
}

type FaceDuplicationCheckIntlRequest struct {
	// Specifies whether to automatically register the face in the specified face database when no duplicate face is found during the search. Valid values:
	//
	// - 0: automatic registration
	//
	// - 1: no registration (default).
	//
	// example:
	//
	// 1
	AutoRegistration *string `json:"AutoRegistration,omitempty" xml:"AutoRegistration,omitempty"`
	// The face database codes created in the console. A maximum of 10 face databases can be queried at a time. Separate multiple face database codes with commas (,).
	//
	// example:
	//
	// 1232344，23444
	FaceGroupCodes *string `json:"FaceGroupCodes,omitempty" xml:"FaceGroupCodes,omitempty"`
	// Specifies whether to enable face quality check.
	//
	// example:
	//
	// Y
	FaceQualityCheck *string `json:"FaceQualityCheck,omitempty" xml:"FaceQualityCheck,omitempty"`
	// The face database for registration.
	//
	// example:
	//
	// 0e0c34a77f
	FaceRegisterGroupCode *string `json:"FaceRegisterGroupCode,omitempty" xml:"FaceRegisterGroupCode,omitempty"`
	// The face matching threshold. 	Warning: This is a reserved field and is not currently enabled.</warning>.
	//
	// example:
	//
	// 0.5
	FaceVerifyThreshold *string `json:"FaceVerifyThreshold,omitempty" xml:"FaceVerifyThreshold,omitempty"`
	// Specifies whether to enable passive liveness detection. Valid values:
	//
	// - 0: disabled
	//
	// - 1: enabled.
	//
	// example:
	//
	// 0
	Liveness *string `json:"Liveness,omitempty" xml:"Liveness,omitempty"`
	// The custom unique business identifier, which is used for subsequent troubleshooting. The value is a combination of letters and digits up to 32 characters in length. Ensure that the value is unique.
	//
	// This parameter is required.
	//
	// example:
	//
	// e0c34a77f5ac40a5aa5e6ed20c35****
	MerchantBizId *string `json:"MerchantBizId,omitempty" xml:"MerchantBizId,omitempty"`
	// The custom user ID or another identifier that can identify a specific user, such as a phone number or email address. We strongly recommend that you desensitize the value of this field in advance, for example, by hashing the value.
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
	// The number of faces to return when multiple faces above the matching threshold are found.
	//
	// - Default value: 1.
	//
	// - Maximum value: 5.
	//
	// example:
	//
	// 1
	ReturnFaces *string `json:"ReturnFaces,omitempty" xml:"ReturnFaces,omitempty"`
	// Specifies the type of face data to save. Valid values:
	//
	// - 0: face image (default)
	//
	// - 1: feature
	//
	// 	Warning: This is a reserved field and is not currently enabled.</warning>.
	//
	// example:
	//
	// 0
	SaveFacePicture *string `json:"SaveFacePicture,omitempty" xml:"SaveFacePicture,omitempty"`
	// The custom verification scenario ID.
	//
	// example:
	//
	// 1234567890
	SceneCode *string `json:"SceneCode,omitempty" xml:"SceneCode,omitempty"`
	// The Base64-encoded facial image.
	//
	// example:
	//
	// base64
	SourceFacePicture *string `json:"SourceFacePicture,omitempty" xml:"SourceFacePicture,omitempty"`
	// The URL of the facial image. The URL must be a publicly accessible HTTP or HTTPS link.
	//
	// example:
	//
	// https://***face1.jpeg
	SourceFacePictureUrl *string `json:"SourceFacePictureUrl,omitempty" xml:"SourceFacePictureUrl,omitempty"`
	// The Base64-encoded facial image.
	//
	// example:
	//
	// base64
	TargetFacePicture *string `json:"TargetFacePicture,omitempty" xml:"TargetFacePicture,omitempty"`
	// The URL of the facial image. The URL must be a publicly accessible HTTP or HTTPS link.
	//
	// example:
	//
	// https://***face2.jpeg
	TargetFacePictureUrl *string `json:"TargetFacePictureUrl,omitempty" xml:"TargetFacePictureUrl,omitempty"`
	// The verification type. Valid values:
	//
	// - 0: retrieve pattern
	//
	// > - Feature: Submits a face database and a user facial image (sourceFacePicture). The system automatically retrieves the face database to check whether the specified facial image (sourceFacePicture) already exists. Passive liveness detection can be enabled for the facial image (sourceFacePicture).
	//
	// > - Recommended scenario: Real-person create an account where duplicate registration is not allowed.
	//
	// - 1 (default): authenticate pattern
	//
	// > - Feature: Submits a specified facial image (sourceFacePicture) and a stored facial image (TargetFacePicture). The system automatically authenticates whether the two faces match. Passive liveness detection can be enabled for the specified facial image (sourceFacePicture).
	//
	// > - Recommended scenario: Authenticating whether the operation is performed by the account owner when logon credentials or account information is modified.
	//
	// - 2: comprehensive pattern
	//
	// > - Feature: Submits a face database, a specified facial image (sourceFacePicture), and a stored facial image (TargetFacePicture). The system automatically retrieves the face database to check whether the specified facial image (sourceFacePicture) exists and whether it matches the stored face. Passive liveness detection can be enabled for the specified facial image (sourceFacePicture).
	//
	// > - Recommended scenario: Authenticating that the user is new and the operation is performed by the user.
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

func (s *FaceDuplicationCheckIntlRequest) GetVerifyModel() *string {
	return s.VerifyModel
}

func (s *FaceDuplicationCheckIntlRequest) SetAutoRegistration(v string) *FaceDuplicationCheckIntlRequest {
	s.AutoRegistration = &v
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

func (s *FaceDuplicationCheckIntlRequest) SetVerifyModel(v string) *FaceDuplicationCheckIntlRequest {
	s.VerifyModel = &v
	return s
}

func (s *FaceDuplicationCheckIntlRequest) Validate() error {
	return dara.Validate(s)
}
