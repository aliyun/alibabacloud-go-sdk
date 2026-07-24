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
	SetVerifyModel(v string) *FaceVerifyIntlRequest
	GetVerifyModel() *string
}

type FaceVerifyIntlRequest struct {
	// Required when ProductCode is set to FACE_IDU_MIN.
	//
	// Specifies whether to automatically register the face to the specified face library when no duplicate face is found during retrieval. Valid values:
	//
	// - 0: Automatic registration.
	//
	// - 1: No registration. This is the default value.
	//
	// example:
	//
	// 1
	AutoRegistration   *string `json:"AutoRegistration,omitempty" xml:"AutoRegistration,omitempty"`
	FaceAttributeCheck *string `json:"FaceAttributeCheck,omitempty" xml:"FaceAttributeCheck,omitempty"`
	// Required when ProductCode is set to FACE_IDU_MIN.
	//
	// The face library codes created by the customer through the console. A maximum of 10 face libraries can be queried at the same time. Separate multiple face library codes with commas (,).
	//
	// example:
	//
	// 1232344，23444
	FaceGroupCodes *string `json:"FaceGroupCodes,omitempty" xml:"FaceGroupCodes,omitempty"`
	// Specifies whether to check the quality of the face image. Valid values:
	//
	// - Y: Enabled.
	//
	// - N: Disabled. This is the default value.
	//
	// example:
	//
	// Y
	FaceQualityCheck *string `json:"FaceQualityCheck,omitempty" xml:"FaceQualityCheck,omitempty"`
	// Required when ProductCode is set to FACE_IDU_MIN.
	//
	// The code of the face library for registration.
	//
	// example:
	//
	// 0e0c34a77f
	FaceRegisterGroupCode *string `json:"FaceRegisterGroupCode,omitempty" xml:"FaceRegisterGroupCode,omitempty"`
	// A custom unique business identifier used for subsequent troubleshooting. The value supports a combination of letters and digits up to 32 characters in length. Make sure the value is unique.
	//
	// example:
	//
	// e0c34a77f5ac40a5aa5e6ed20c35****
	MerchantBizId *string `json:"MerchantBizId,omitempty" xml:"MerchantBizId,omitempty"`
	// A custom user ID or other identifier that can identify a specific user, such as a phone number or email address. We strongly recommend that you desensitize the value of this field in advance, such as by hashing the value.
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
	// Specifies the number of faces to return when multiple faces exist above the matching threshold. Default value: 1. Maximum value: 5.
	//
	// example:
	//
	// 1
	ReturnFaces *string `json:"ReturnFaces,omitempty" xml:"ReturnFaces,omitempty"`
	// The Base64-encoded portrait image.
	//
	// > **Note**
	//
	// >
	//
	// > - If you use this method to pass in the image, check the image size and do not pass in an excessively large image.
	//
	// > - Specify one of the following parameters: SourceFacePicture, SourceFacePictureUrl, or SourceFacePictureFile.
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
	// The publicly accessible HTTPS URL of the portrait image.
	//
	// example:
	//
	// https://***face1.jpeg
	SourceFacePictureUrl *string `json:"SourceFacePictureUrl,omitempty" xml:"SourceFacePictureUrl,omitempty"`
	// The Base64-encoded reference face image.
	//
	// > **Note**
	//
	// >
	//
	// > - If you use this method to pass in the image, check the image size and do not pass in an excessively large image.
	//
	// > - Specify one of the following parameters: TargetFacePicture, TargetFacePictureUrl, or TargetFacePictureFile.
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
	// Required when ProductCode is set to FACE_IDU_MIN. The verification type. Valid values:
	//
	// - 0: retrieve pattern
	//
	// > - Feature: Pass in a face library and a user face image (sourceFacePicture). The system automatically retrieves the face library to check whether the specified face image (sourceFacePicture) already exists. Passive liveness detection can be enabled for the face image (sourceFacePicture).
	//
	// > - Recommended scenario: Real-person create an account where duplicate registration is not allowed.
	//
	// - 1 (default): authenticate pattern
	//
	// > - Feature: Pass in a specified face image (sourceFacePicture) and a reference face image (TargetFacePicture). The system automatically authenticates whether the faces match. Passive liveness detection can be enabled for the specified face image (sourceFacePicture).
	//
	// > - Recommended scenario: Authenticating the identity of the user when modifying logon credentials or account information.
	//
	// - 2: comprehensive pattern
	//
	// > - Feature: Pass in a face library, a specified face image (sourceFacePicture), and a reference face image (TargetFacePicture). The system automatically retrieves the face library to check whether the specified face image (sourceFacePicture) exists, authenticates whether it matches the reference face, and supports enabling passive liveness detection for the specified face image (sourceFacePicture).
	//
	// > - Recommended scenario: Authenticating that the user is a new user and the operation is performed by the user in person.
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

func (s *FaceVerifyIntlRequest) SetVerifyModel(v string) *FaceVerifyIntlRequest {
	s.VerifyModel = &v
	return s
}

func (s *FaceVerifyIntlRequest) Validate() error {
	return dara.Validate(s)
}
