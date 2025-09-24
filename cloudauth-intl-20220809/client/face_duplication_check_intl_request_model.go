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
	// Indicates whether to automatically register the face to the specified face library if no duplicate face is found.
	//
	// - 0- Auto-register (default)
	//
	// - 1- Do not register
	//
	// example:
	//
	// 0
	AutoRegistration *string `json:"AutoRegistration,omitempty" xml:"AutoRegistration,omitempty"`
	// The face library code created through the console, supporting up to 10 face libraries simultaneously. When multiple face library codes are passed, they should be separated by commas.
	//
	// example:
	//
	// 1232344，23444
	FaceGroupCodes *string `json:"FaceGroupCodes,omitempty" xml:"FaceGroupCodes,omitempty"`
	// Face registration library.
	//
	// example:
	//
	// 0e0c34a77f
	FaceRegisterGroupCode *string `json:"FaceRegisterGroupCode,omitempty" xml:"FaceRegisterGroupCode,omitempty"`
	// Face matching threshold.
	//
	// example:
	//
	// 0.5
	FaceVerifyThreshold *string `json:"FaceVerifyThreshold,omitempty" xml:"FaceVerifyThreshold,omitempty"`
	// Whether to enable silent liveness detection
	//
	// - 0- Disabled
	//
	// - 1- Enabled
	//
	// example:
	//
	// 0
	Liveness *string `json:"Liveness,omitempty" xml:"Liveness,omitempty"`
	// A unique business identifier for troubleshooting purposes. It supports a combination of 32 alphanumeric characters, please ensure its uniqueness.
	//
	// This parameter is required.
	//
	// example:
	//
	// e0c34a77f5ac40a5aa5e6ed20c35****
	MerchantBizId *string `json:"MerchantBizId,omitempty" xml:"MerchantBizId,omitempty"`
	// Your custom user ID or other identifiers that can uniquely identify a specific user, such as a phone number or email address. It is strongly recommended to pre-desensitize the value of this field, for example, by hashing it.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234567890
	MerchantUserId *string `json:"MerchantUserId,omitempty" xml:"MerchantUserId,omitempty"`
	// Product code
	//
	// This parameter is required.
	//
	// example:
	//
	// FACE_IDU_MIN
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// When there are multiple faces above the matching threshold, you can use this parameter to customize the number of returned faces
	//
	// - Default returns 1
	//
	// - Maximum support 5
	//
	// example:
	//
	// 1
	ReturnFaces *string `json:"ReturnFaces,omitempty" xml:"ReturnFaces,omitempty"`
	// Distinguishes between saving the face image and features
	//
	// - 0- Face (default)
	//
	// - 1- Features
	//
	// example:
	//
	// 0
	SaveFacePicture *string `json:"SaveFacePicture,omitempty" xml:"SaveFacePicture,omitempty"`
	// Your custom authentication scenario ID.
	//
	// example:
	//
	// 1234567890
	SceneCode *string `json:"SceneCode,omitempty" xml:"SceneCode,omitempty"`
	// Base64 encoded portrait photo.
	//
	// example:
	//
	// base64
	SourceFacePicture *string `json:"SourceFacePicture,omitempty" xml:"SourceFacePicture,omitempty"`
	// Portrait image URL, accessible via public HTTP or HTTPS link.
	//
	// example:
	//
	// https://***face1.jpeg
	SourceFacePictureUrl *string `json:"SourceFacePictureUrl,omitempty" xml:"SourceFacePictureUrl,omitempty"`
	// Base64 encoded portrait photo.
	//
	// example:
	//
	// base64
	TargetFacePicture *string `json:"TargetFacePicture,omitempty" xml:"TargetFacePicture,omitempty"`
	// Portrait image URL, accessible via public HTTP or HTTPS link.
	//
	// example:
	//
	// https://***face2.jpeg
	TargetFacePictureUrl *string `json:"TargetFacePictureUrl,omitempty" xml:"TargetFacePictureUrl,omitempty"`
	// Verification type
	//
	// - 0- 1:N (default)
	//
	// - 1- 1:1
	//
	// - 2- 1:N + 1:1
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
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
