// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFaceLivenessV2Request interface {
	dara.Model
	String() string
	GoString() string
	SetFacePictureBase64(v string) *FaceLivenessV2Request
	GetFacePictureBase64() *string
	SetFacePictureFile(v string) *FaceLivenessV2Request
	GetFacePictureFile() *string
	SetFacePictureUrl(v string) *FaceLivenessV2Request
	GetFacePictureUrl() *string
	SetFaceQualityCheck(v string) *FaceLivenessV2Request
	GetFaceQualityCheck() *string
	SetMerchantBizId(v string) *FaceLivenessV2Request
	GetMerchantBizId() *string
	SetMerchantUserId(v string) *FaceLivenessV2Request
	GetMerchantUserId() *string
	SetProductCode(v string) *FaceLivenessV2Request
	GetProductCode() *string
}

type FaceLivenessV2Request struct {
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
	FacePictureFile *string `json:"FacePictureFile,omitempty" xml:"FacePictureFile,omitempty"`
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

func (s FaceLivenessV2Request) String() string {
	return dara.Prettify(s)
}

func (s FaceLivenessV2Request) GoString() string {
	return s.String()
}

func (s *FaceLivenessV2Request) GetFacePictureBase64() *string {
	return s.FacePictureBase64
}

func (s *FaceLivenessV2Request) GetFacePictureFile() *string {
	return s.FacePictureFile
}

func (s *FaceLivenessV2Request) GetFacePictureUrl() *string {
	return s.FacePictureUrl
}

func (s *FaceLivenessV2Request) GetFaceQualityCheck() *string {
	return s.FaceQualityCheck
}

func (s *FaceLivenessV2Request) GetMerchantBizId() *string {
	return s.MerchantBizId
}

func (s *FaceLivenessV2Request) GetMerchantUserId() *string {
	return s.MerchantUserId
}

func (s *FaceLivenessV2Request) GetProductCode() *string {
	return s.ProductCode
}

func (s *FaceLivenessV2Request) SetFacePictureBase64(v string) *FaceLivenessV2Request {
	s.FacePictureBase64 = &v
	return s
}

func (s *FaceLivenessV2Request) SetFacePictureFile(v string) *FaceLivenessV2Request {
	s.FacePictureFile = &v
	return s
}

func (s *FaceLivenessV2Request) SetFacePictureUrl(v string) *FaceLivenessV2Request {
	s.FacePictureUrl = &v
	return s
}

func (s *FaceLivenessV2Request) SetFaceQualityCheck(v string) *FaceLivenessV2Request {
	s.FaceQualityCheck = &v
	return s
}

func (s *FaceLivenessV2Request) SetMerchantBizId(v string) *FaceLivenessV2Request {
	s.MerchantBizId = &v
	return s
}

func (s *FaceLivenessV2Request) SetMerchantUserId(v string) *FaceLivenessV2Request {
	s.MerchantUserId = &v
	return s
}

func (s *FaceLivenessV2Request) SetProductCode(v string) *FaceLivenessV2Request {
	s.ProductCode = &v
	return s
}

func (s *FaceLivenessV2Request) Validate() error {
	return dara.Validate(s)
}
