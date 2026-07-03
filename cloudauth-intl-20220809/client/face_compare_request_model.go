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
	// Whether to enable quality detection for the input face image	Danger: Deprecated
	//
	// example:
	//
	// N
	FacePictureQualityCheck *string `json:"FacePictureQualityCheck,omitempty" xml:"FacePictureQualityCheck,omitempty"`
	// Face quality check
	//
	// example:
	//
	// Y
	FaceQualityCheck *string `json:"FaceQualityCheck,omitempty" xml:"FaceQualityCheck,omitempty"`
	// A unique business identifier customized by the merchant, used for subsequent troubleshooting. Supports a combination of letters and numbers with a maximum length of 32 characters. Ensure it is unique.
	//
	// example:
	//
	// e0c34a77f5ac40a5aa5e6ed20c353888
	MerchantBizId *string `json:"MerchantBizId,omitempty" xml:"MerchantBizId,omitempty"`
	// Base64-encoded face photo.
	//
	// Note
	//
	// - If you choose this method to pass in the photo, check the photo size and do not pass in an oversized photo.
	//
	// - Either SourceFacePicture or SourceFacePictureUrl must be specified.
	//
	// example:
	//
	// base64
	SourceFacePicture *string `json:"SourceFacePicture,omitempty" xml:"SourceFacePicture,omitempty"`
	// The HTTPS or HTTP URL of the face image.
	//
	// example:
	//
	// https://***face1.jpeg
	SourceFacePictureUrl *string `json:"SourceFacePictureUrl,omitempty" xml:"SourceFacePictureUrl,omitempty"`
	// Base64-encoded reference photo.
	//
	// Note
	//
	// - If you choose this method to pass in the photo, check the photo size and do not pass in an oversized photo.
	//
	// - Either TargetFacePicture or TargetFacePictureUrl must be specified.
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
