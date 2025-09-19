// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFaceLivenessRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCrop(v string) *FaceLivenessRequest
	GetCrop() *string
	SetFacePictureBase64(v string) *FaceLivenessRequest
	GetFacePictureBase64() *string
	SetFacePictureUrl(v string) *FaceLivenessRequest
	GetFacePictureUrl() *string
	SetFaceQuality(v string) *FaceLivenessRequest
	GetFaceQuality() *string
	SetMerchantBizId(v string) *FaceLivenessRequest
	GetMerchantBizId() *string
	SetMerchantUserId(v string) *FaceLivenessRequest
	GetMerchantUserId() *string
	SetOcclusion(v string) *FaceLivenessRequest
	GetOcclusion() *string
	SetProductCode(v string) *FaceLivenessRequest
	GetProductCode() *string
}

type FaceLivenessRequest struct {
	// Specifies whether to crop the facial image. The default value is F.
	//
	// - **T**: allows cropping.
	//
	// - **F**: Forbidden
	//
	// example:
	//
	// T
	Crop              *string `json:"Crop,omitempty" xml:"Crop,omitempty"`
	FacePictureBase64 *string `json:"FacePictureBase64,omitempty" xml:"FacePictureBase64,omitempty"`
	// The URL of the portrait image. The URL must be an HTTP or HTTPS link accessible over the Internet.
	//
	// example:
	//
	// https://digital-face-prod8.oss-cn-hangzhou.aliyuncs.com/1669520556530-expo/default/face/20221127114236530_w3kx2e6t.jpg
	FacePictureUrl *string `json:"FacePictureUrl,omitempty" xml:"FacePictureUrl,omitempty"`
	// Specifies whether to return the facial image quality score. The default value is F.
	//
	// - **T**: returns the score.
	//
	// - **F**: does not return the score.
	//
	// example:
	//
	// T
	FaceQuality *string `json:"FaceQuality,omitempty" xml:"FaceQuality,omitempty"`
	// A custom unique business identifier. You can use this identifier to track and troubleshoot issues. The identifier can be up to 32 characters in length and can contain letters and digits. Make sure the identifier is unique.
	//
	// > Alibaba Cloud servers do not check the uniqueness of this value. For better tracking, ensure this value is unique.
	//
	// example:
	//
	// e0c34a77f5ac40a5aa5e6ed20c353888
	MerchantBizId *string `json:"MerchantBizId,omitempty" xml:"MerchantBizId,omitempty"`
	// A  custom user ID or another identifier for a specific user, such as a mobile number or email address. For security, desensitize this value in advance, for example, by hashing it.
	//
	// example:
	//
	// 123456789
	MerchantUserId *string `json:"MerchantUserId,omitempty" xml:"MerchantUserId,omitempty"`
	// Specifies whether to enable occlusion detection. The default value is F.
	//
	// - **T**: enables the feature.
	//
	// - **F**: disables the feature.
	//
	// example:
	//
	// T
	Occlusion *string `json:"Occlusion,omitempty" xml:"Occlusion,omitempty"`
	// The product solution to use. Set the value to **FACE_LIVENESS_MIN*	- to use the passive liveness detection API.
	//
	// example:
	//
	// FACE_LIVENESS_MIN
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
}

func (s FaceLivenessRequest) String() string {
	return dara.Prettify(s)
}

func (s FaceLivenessRequest) GoString() string {
	return s.String()
}

func (s *FaceLivenessRequest) GetCrop() *string {
	return s.Crop
}

func (s *FaceLivenessRequest) GetFacePictureBase64() *string {
	return s.FacePictureBase64
}

func (s *FaceLivenessRequest) GetFacePictureUrl() *string {
	return s.FacePictureUrl
}

func (s *FaceLivenessRequest) GetFaceQuality() *string {
	return s.FaceQuality
}

func (s *FaceLivenessRequest) GetMerchantBizId() *string {
	return s.MerchantBizId
}

func (s *FaceLivenessRequest) GetMerchantUserId() *string {
	return s.MerchantUserId
}

func (s *FaceLivenessRequest) GetOcclusion() *string {
	return s.Occlusion
}

func (s *FaceLivenessRequest) GetProductCode() *string {
	return s.ProductCode
}

func (s *FaceLivenessRequest) SetCrop(v string) *FaceLivenessRequest {
	s.Crop = &v
	return s
}

func (s *FaceLivenessRequest) SetFacePictureBase64(v string) *FaceLivenessRequest {
	s.FacePictureBase64 = &v
	return s
}

func (s *FaceLivenessRequest) SetFacePictureUrl(v string) *FaceLivenessRequest {
	s.FacePictureUrl = &v
	return s
}

func (s *FaceLivenessRequest) SetFaceQuality(v string) *FaceLivenessRequest {
	s.FaceQuality = &v
	return s
}

func (s *FaceLivenessRequest) SetMerchantBizId(v string) *FaceLivenessRequest {
	s.MerchantBizId = &v
	return s
}

func (s *FaceLivenessRequest) SetMerchantUserId(v string) *FaceLivenessRequest {
	s.MerchantUserId = &v
	return s
}

func (s *FaceLivenessRequest) SetOcclusion(v string) *FaceLivenessRequest {
	s.Occlusion = &v
	return s
}

func (s *FaceLivenessRequest) SetProductCode(v string) *FaceLivenessRequest {
	s.ProductCode = &v
	return s
}

func (s *FaceLivenessRequest) Validate() error {
	return dara.Validate(s)
}
