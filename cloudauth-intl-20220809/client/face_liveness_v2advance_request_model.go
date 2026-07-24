// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
	"io"
)

type iFaceLivenessV2AdvanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFaceAttributeCheck(v string) *FaceLivenessV2AdvanceRequest
	GetFaceAttributeCheck() *string
	SetFacePictureBase64(v string) *FaceLivenessV2AdvanceRequest
	GetFacePictureBase64() *string
	SetFacePictureFileObject(v io.Reader) *FaceLivenessV2AdvanceRequest
	GetFacePictureFileObject() io.Reader
	SetFacePictureUrl(v string) *FaceLivenessV2AdvanceRequest
	GetFacePictureUrl() *string
	SetFaceQualityCheck(v string) *FaceLivenessV2AdvanceRequest
	GetFaceQualityCheck() *string
	SetMerchantBizId(v string) *FaceLivenessV2AdvanceRequest
	GetMerchantBizId() *string
	SetMerchantUserId(v string) *FaceLivenessV2AdvanceRequest
	GetMerchantUserId() *string
	SetProductCode(v string) *FaceLivenessV2AdvanceRequest
	GetProductCode() *string
}

type FaceLivenessV2AdvanceRequest struct {
	FaceAttributeCheck *string `json:"FaceAttributeCheck,omitempty" xml:"FaceAttributeCheck,omitempty"`
	// The Base64-encoded face image.
	//
	// > **Note**
	//
	// - If you use this method to pass in the image, check the image size and do not pass in an excessively large image.
	//
	// - Specify one of the following parameters: FacePictureBase64, FacePictureUrl, or FacePictureFile.
	//
	// example:
	//
	// Base64
	FacePictureBase64 *string `json:"FacePictureBase64,omitempty" xml:"FacePictureBase64,omitempty"`
	// The file stream of the face image.
	//
	// example:
	//
	// InputStream
	FacePictureFileObject io.Reader `json:"FacePictureFile,omitempty" xml:"FacePictureFile,omitempty"`
	// The URL of the face image. The URL must be a publicly accessible HTTPS URL.
	//
	// example:
	//
	// https://***
	FacePictureUrl *string `json:"FacePictureUrl,omitempty" xml:"FacePictureUrl,omitempty"`
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
	// The merchant-defined unique business ID for subsequent troubleshooting. The value can be a combination of letters and digits with a maximum length of 32 characters. Make sure the value is unique.
	//
	// example:
	//
	// e0c34a***353888
	MerchantBizId *string `json:"MerchantBizId,omitempty" xml:"MerchantBizId,omitempty"`
	// The custom user ID or another identifier that can identify a specific user, such as a phone number or email address. We strongly recommend that you desensitize the value of this field in advance, for example, by hashing the value.
	//
	// example:
	//
	// 123456789
	MerchantUserId *string `json:"MerchantUserId,omitempty" xml:"MerchantUserId,omitempty"`
	// The product plan to use. Valid values: FACE_LIVENESS_MIN_PRO and FACE_LIVENESS_MIN.
	//
	// This parameter is required.
	//
	// example:
	//
	// FACE_LIVENESS_MIN_PRO
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
}

func (s FaceLivenessV2AdvanceRequest) String() string {
	return dara.Prettify(s)
}

func (s FaceLivenessV2AdvanceRequest) GoString() string {
	return s.String()
}

func (s *FaceLivenessV2AdvanceRequest) GetFaceAttributeCheck() *string {
	return s.FaceAttributeCheck
}

func (s *FaceLivenessV2AdvanceRequest) GetFacePictureBase64() *string {
	return s.FacePictureBase64
}

func (s *FaceLivenessV2AdvanceRequest) GetFacePictureFileObject() io.Reader {
	return s.FacePictureFileObject
}

func (s *FaceLivenessV2AdvanceRequest) GetFacePictureUrl() *string {
	return s.FacePictureUrl
}

func (s *FaceLivenessV2AdvanceRequest) GetFaceQualityCheck() *string {
	return s.FaceQualityCheck
}

func (s *FaceLivenessV2AdvanceRequest) GetMerchantBizId() *string {
	return s.MerchantBizId
}

func (s *FaceLivenessV2AdvanceRequest) GetMerchantUserId() *string {
	return s.MerchantUserId
}

func (s *FaceLivenessV2AdvanceRequest) GetProductCode() *string {
	return s.ProductCode
}

func (s *FaceLivenessV2AdvanceRequest) SetFaceAttributeCheck(v string) *FaceLivenessV2AdvanceRequest {
	s.FaceAttributeCheck = &v
	return s
}

func (s *FaceLivenessV2AdvanceRequest) SetFacePictureBase64(v string) *FaceLivenessV2AdvanceRequest {
	s.FacePictureBase64 = &v
	return s
}

func (s *FaceLivenessV2AdvanceRequest) SetFacePictureFileObject(v io.Reader) *FaceLivenessV2AdvanceRequest {
	s.FacePictureFileObject = v
	return s
}

func (s *FaceLivenessV2AdvanceRequest) SetFacePictureUrl(v string) *FaceLivenessV2AdvanceRequest {
	s.FacePictureUrl = &v
	return s
}

func (s *FaceLivenessV2AdvanceRequest) SetFaceQualityCheck(v string) *FaceLivenessV2AdvanceRequest {
	s.FaceQualityCheck = &v
	return s
}

func (s *FaceLivenessV2AdvanceRequest) SetMerchantBizId(v string) *FaceLivenessV2AdvanceRequest {
	s.MerchantBizId = &v
	return s
}

func (s *FaceLivenessV2AdvanceRequest) SetMerchantUserId(v string) *FaceLivenessV2AdvanceRequest {
	s.MerchantUserId = &v
	return s
}

func (s *FaceLivenessV2AdvanceRequest) SetProductCode(v string) *FaceLivenessV2AdvanceRequest {
	s.ProductCode = &v
	return s
}

func (s *FaceLivenessV2AdvanceRequest) Validate() error {
	return dara.Validate(s)
}
