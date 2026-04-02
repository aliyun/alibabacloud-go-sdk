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
	// example:
	//
	// N
	FacePictureQualityCheck *string `json:"FacePictureQualityCheck,omitempty" xml:"FacePictureQualityCheck,omitempty"`
	// example:
	//
	// e0c34a77f5ac40a5aa5e6ed20c353888
	MerchantBizId *string `json:"MerchantBizId,omitempty" xml:"MerchantBizId,omitempty"`
	// example:
	//
	// base64
	SourceFacePicture *string `json:"SourceFacePicture,omitempty" xml:"SourceFacePicture,omitempty"`
	// example:
	//
	// InputStream
	SourceFacePictureFile *string `json:"SourceFacePictureFile,omitempty" xml:"SourceFacePictureFile,omitempty"`
	// example:
	//
	// https://***face1.jpeg
	SourceFacePictureUrl *string `json:"SourceFacePictureUrl,omitempty" xml:"SourceFacePictureUrl,omitempty"`
	// example:
	//
	// base64
	TargetFacePicture *string `json:"TargetFacePicture,omitempty" xml:"TargetFacePicture,omitempty"`
	// example:
	//
	// InputStream
	TargetFacePictureFile *string `json:"TargetFacePictureFile,omitempty" xml:"TargetFacePictureFile,omitempty"`
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
