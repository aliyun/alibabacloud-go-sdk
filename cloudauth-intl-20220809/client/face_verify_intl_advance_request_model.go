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
	// Required when ProductCode is set to FACE_IDU_MIN.
	//
	// Specifies whether to automatically register the face to the specified face library when no duplicate face is found during retrieval. Valid values:
	//
	// - 0: automatic registration.
	//
	// - 1: no registration. This is the default value.
	//
	// example:
	//
	// 1
	AutoRegistration *string `json:"AutoRegistration,omitempty" xml:"AutoRegistration,omitempty"`
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
	// The face library for registration.
	//
	// example:
	//
	// 0e0c34a77f
	FaceRegisterGroupCode *string `json:"FaceRegisterGroupCode,omitempty" xml:"FaceRegisterGroupCode,omitempty"`
	// A custom unique business identifier used for subsequent troubleshooting. The value supports a combination of letters and numbers up to 32 characters in length. Make sure the value is unique.
	//
	// example:
	//
	// e0c34a77f5ac40a5aa5e6ed20c35****
	MerchantBizId *string `json:"MerchantBizId,omitempty" xml:"MerchantBizId,omitempty"`
	// A custom user ID or other identifier that can identify a specific user, such as a phone number or email address. We strongly recommend that you desensitize the value of this field in advance, for example, by hashing the value.
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
	// Specifies the number of faces to return when multiple faces exist above the matching threshold. You can use this parameter to customize the number of returned faces.
	//
	// - Default value: 1.
	//
	// - Maximum value: 5.
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
	SourceFacePictureFileObject io.Reader `json:"SourceFacePictureFile,omitempty" xml:"SourceFacePictureFile,omitempty"`
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
	TargetFacePictureFileObject io.Reader `json:"TargetFacePictureFile,omitempty" xml:"TargetFacePictureFile,omitempty"`
	// The HTTPS URL of the reference face image.
	//
	// example:
	//
	// https://***face2.jpeg
	TargetFacePictureUrl *string `json:"TargetFacePictureUrl,omitempty" xml:"TargetFacePictureUrl,omitempty"`
	// Required when ProductCode is set to FACE_IDU_MIN.
	//
	// The verification type. Valid values:
	//
	// - 0: retrieval pattern.
	//
	// > - Feature: Pass in a face library and a user face image (sourceFacePicture). The system automatically retrieves whether the specified face image (sourceFacePicture) already exists in the face library. Passive liveness detection can be enabled for the face image (sourceFacePicture).
	//
	// > - Recommended scenario: real-person account creation where duplicate registration is not allowed.
	//
	// - 1 (default): authentication pattern.
	//
	// > - Feature: Pass in a specified face image (sourceFacePicture) and a reference face image (TargetFacePicture). The system automatically authenticates whether the faces match. Passive liveness detection can be enabled for the specified face image (sourceFacePicture).
	//
	// > - Recommended scenario: authenticating the identity of the user when modifying logon credentials or account information.
	//
	// - 2: comprehensive pattern.
	//
	// > - Feature: Pass in a face library, a specified face image (sourceFacePicture), and a reference face image (TargetFacePicture). The system automatically retrieves whether the specified face image (sourceFacePicture) exists in the face library, authenticates whether it matches the reference face, and supports enabling passive liveness detection for the specified face image (sourceFacePicture).
	//
	// > - Recommended scenario: verifying that the user is new and creating an account in person.
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
