// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFaceCompareRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFacePictureQualityCheck(v string) *FaceCompareRequest
	GetFacePictureQualityCheck() *string
	SetFaceQualityCheck(v string) *FaceCompareRequest
	GetFaceQualityCheck() *string
	SetMerchantBizId(v string) *FaceCompareRequest
	GetMerchantBizId() *string
	SetSourceFacePicture(v string) *FaceCompareRequest
	GetSourceFacePicture() *string
	SetSourceFacePictureUrl(v string) *FaceCompareRequest
	GetSourceFacePictureUrl() *string
	SetTargetFacePicture(v string) *FaceCompareRequest
	GetTargetFacePicture() *string
	SetTargetFacePictureUrl(v string) *FaceCompareRequest
	GetTargetFacePictureUrl() *string
}

type FaceCompareRequest struct {
	// Specifies whether to enable face image quality check.<danger>Deprecated.</danger>.
	//
	// example:
	//
	// N
	FacePictureQualityCheck *string `json:"FacePictureQualityCheck,omitempty" xml:"FacePictureQualityCheck,omitempty"`
	// The face quality check.
	//
	// example:
	//
	// Y
	FaceQualityCheck *string `json:"FaceQualityCheck,omitempty" xml:"FaceQualityCheck,omitempty"`
	// The merchant-defined unique business ID used for subsequent troubleshooting. The value can be a combination of letters and digits with a maximum length of 32 characters. Ensure that the value is unique.
	//
	// example:
	//
	// e0c34a77f5ac40a5aa5e6ed20c353888
	MerchantBizId *string `json:"MerchantBizId,omitempty" xml:"MerchantBizId,omitempty"`
	// The Base64-encoded source face image.
	//
	// > **Note**
	//
	// > - If you use this method to pass in the image, check the image size and do not pass in an excessively large image.
	//
	// > - Specify either SourceFacePicture or SourceFacePictureUrl.
	//
	// example:
	//
	// base64
	SourceFacePicture *string `json:"SourceFacePicture,omitempty" xml:"SourceFacePicture,omitempty"`
	// The HTTPS or HTTP URL of the source face image.
	//
	// example:
	//
	// https://***face1.jpeg
	SourceFacePictureUrl *string `json:"SourceFacePictureUrl,omitempty" xml:"SourceFacePictureUrl,omitempty"`
	// The Base64-encoded reference face image.
	//
	// > **Note**
	//
	// > - If you use this method to pass in the image, check the image size and do not pass in an excessively large image.
	//
	// > - Specify either TargetFacePicture or TargetFacePictureUrl.
	//
	// example:
	//
	// base64
	TargetFacePicture *string `json:"TargetFacePicture,omitempty" xml:"TargetFacePicture,omitempty"`
	// The HTTPS or HTTP URL of the reference face image.
	//
	// example:
	//
	// https://***face2.jpeg
	TargetFacePictureUrl *string `json:"TargetFacePictureUrl,omitempty" xml:"TargetFacePictureUrl,omitempty"`
}

func (s FaceCompareRequest) String() string {
	return dara.Prettify(s)
}

func (s FaceCompareRequest) GoString() string {
	return s.String()
}

func (s *FaceCompareRequest) GetFacePictureQualityCheck() *string {
	return s.FacePictureQualityCheck
}

func (s *FaceCompareRequest) GetFaceQualityCheck() *string {
	return s.FaceQualityCheck
}

func (s *FaceCompareRequest) GetMerchantBizId() *string {
	return s.MerchantBizId
}

func (s *FaceCompareRequest) GetSourceFacePicture() *string {
	return s.SourceFacePicture
}

func (s *FaceCompareRequest) GetSourceFacePictureUrl() *string {
	return s.SourceFacePictureUrl
}

func (s *FaceCompareRequest) GetTargetFacePicture() *string {
	return s.TargetFacePicture
}

func (s *FaceCompareRequest) GetTargetFacePictureUrl() *string {
	return s.TargetFacePictureUrl
}

func (s *FaceCompareRequest) SetFacePictureQualityCheck(v string) *FaceCompareRequest {
	s.FacePictureQualityCheck = &v
	return s
}

func (s *FaceCompareRequest) SetFaceQualityCheck(v string) *FaceCompareRequest {
	s.FaceQualityCheck = &v
	return s
}

func (s *FaceCompareRequest) SetMerchantBizId(v string) *FaceCompareRequest {
	s.MerchantBizId = &v
	return s
}

func (s *FaceCompareRequest) SetSourceFacePicture(v string) *FaceCompareRequest {
	s.SourceFacePicture = &v
	return s
}

func (s *FaceCompareRequest) SetSourceFacePictureUrl(v string) *FaceCompareRequest {
	s.SourceFacePictureUrl = &v
	return s
}

func (s *FaceCompareRequest) SetTargetFacePicture(v string) *FaceCompareRequest {
	s.TargetFacePicture = &v
	return s
}

func (s *FaceCompareRequest) SetTargetFacePictureUrl(v string) *FaceCompareRequest {
	s.TargetFacePictureUrl = &v
	return s
}

func (s *FaceCompareRequest) Validate() error {
	return dara.Validate(s)
}
