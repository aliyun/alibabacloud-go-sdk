// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFaceCompareV2Request interface {
	dara.Model
	String() string
	GoString() string
	SetFacePictureQualityCheck(v string) *FaceCompareV2Request
	GetFacePictureQualityCheck() *string
	SetFaceQualityCheck(v string) *FaceCompareV2Request
	GetFaceQualityCheck() *string
	SetMerchantBizId(v string) *FaceCompareV2Request
	GetMerchantBizId() *string
	SetSourceFacePicture(v string) *FaceCompareV2Request
	GetSourceFacePicture() *string
	SetSourceFacePictureFile(v string) *FaceCompareV2Request
	GetSourceFacePictureFile() *string
	SetSourceFacePictureUrl(v string) *FaceCompareV2Request
	GetSourceFacePictureUrl() *string
	SetTargetFacePicture(v string) *FaceCompareV2Request
	GetTargetFacePicture() *string
	SetTargetFacePictureFile(v string) *FaceCompareV2Request
	GetTargetFacePictureFile() *string
	SetTargetFacePictureUrl(v string) *FaceCompareV2Request
	GetTargetFacePictureUrl() *string
}

type FaceCompareV2Request struct {
	// Specifies whether to enable quality check for the input face images.
	//
	// 	Danger: Deprecated.
	//
	// example:
	//
	// N
	FacePictureQualityCheck *string `json:"FacePictureQualityCheck,omitempty" xml:"FacePictureQualityCheck,omitempty"`
	// Specifies whether to enable face quality check.
	//
	// example:
	//
	// Y
	FaceQualityCheck *string `json:"FaceQualityCheck,omitempty" xml:"FaceQualityCheck,omitempty"`
	// The merchant-defined unique business ID used for subsequent troubleshooting. The value can be a combination of letters and numbers with a maximum length of 32 characters. Ensure that the value is unique.
	//
	// example:
	//
	// e0c34a77f5ac40a5aa5e6ed20c353888
	MerchantBizId *string `json:"MerchantBizId,omitempty" xml:"MerchantBizId,omitempty"`
	// The Base64-encoded source face image.
	//
	// > **Note**
	//
	// - If you use this method to pass in the image, check the image size and do not pass in an excessively large image.
	//
	// - Specify one of the following parameters: SourceFacePicture, SourceFacePictureUrl, or SourceFacePictureFile.
	//
	// example:
	//
	// base64
	SourceFacePicture *string `json:"SourceFacePicture,omitempty" xml:"SourceFacePicture,omitempty"`
	// The file stream of the source face image.
	//
	// example:
	//
	// InputStream
	SourceFacePictureFile *string `json:"SourceFacePictureFile,omitempty" xml:"SourceFacePictureFile,omitempty"`
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
	// - If you use this method to pass in the image, check the image size and do not pass in an excessively large image.
	//
	// - Specify one of the following parameters: TargetFacePicture, TargetFacePictureUrl, or TargetFacePictureFile.
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
	// The HTTPS or HTTP URL of the reference face image.
	//
	// example:
	//
	// https://***face2.jpeg
	TargetFacePictureUrl *string `json:"TargetFacePictureUrl,omitempty" xml:"TargetFacePictureUrl,omitempty"`
}

func (s FaceCompareV2Request) String() string {
	return dara.Prettify(s)
}

func (s FaceCompareV2Request) GoString() string {
	return s.String()
}

func (s *FaceCompareV2Request) GetFacePictureQualityCheck() *string {
	return s.FacePictureQualityCheck
}

func (s *FaceCompareV2Request) GetFaceQualityCheck() *string {
	return s.FaceQualityCheck
}

func (s *FaceCompareV2Request) GetMerchantBizId() *string {
	return s.MerchantBizId
}

func (s *FaceCompareV2Request) GetSourceFacePicture() *string {
	return s.SourceFacePicture
}

func (s *FaceCompareV2Request) GetSourceFacePictureFile() *string {
	return s.SourceFacePictureFile
}

func (s *FaceCompareV2Request) GetSourceFacePictureUrl() *string {
	return s.SourceFacePictureUrl
}

func (s *FaceCompareV2Request) GetTargetFacePicture() *string {
	return s.TargetFacePicture
}

func (s *FaceCompareV2Request) GetTargetFacePictureFile() *string {
	return s.TargetFacePictureFile
}

func (s *FaceCompareV2Request) GetTargetFacePictureUrl() *string {
	return s.TargetFacePictureUrl
}

func (s *FaceCompareV2Request) SetFacePictureQualityCheck(v string) *FaceCompareV2Request {
	s.FacePictureQualityCheck = &v
	return s
}

func (s *FaceCompareV2Request) SetFaceQualityCheck(v string) *FaceCompareV2Request {
	s.FaceQualityCheck = &v
	return s
}

func (s *FaceCompareV2Request) SetMerchantBizId(v string) *FaceCompareV2Request {
	s.MerchantBizId = &v
	return s
}

func (s *FaceCompareV2Request) SetSourceFacePicture(v string) *FaceCompareV2Request {
	s.SourceFacePicture = &v
	return s
}

func (s *FaceCompareV2Request) SetSourceFacePictureFile(v string) *FaceCompareV2Request {
	s.SourceFacePictureFile = &v
	return s
}

func (s *FaceCompareV2Request) SetSourceFacePictureUrl(v string) *FaceCompareV2Request {
	s.SourceFacePictureUrl = &v
	return s
}

func (s *FaceCompareV2Request) SetTargetFacePicture(v string) *FaceCompareV2Request {
	s.TargetFacePicture = &v
	return s
}

func (s *FaceCompareV2Request) SetTargetFacePictureFile(v string) *FaceCompareV2Request {
	s.TargetFacePictureFile = &v
	return s
}

func (s *FaceCompareV2Request) SetTargetFacePictureUrl(v string) *FaceCompareV2Request {
	s.TargetFacePictureUrl = &v
	return s
}

func (s *FaceCompareV2Request) Validate() error {
	return dara.Validate(s)
}
