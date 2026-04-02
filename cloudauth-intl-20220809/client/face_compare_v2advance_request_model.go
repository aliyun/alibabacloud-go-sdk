// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
	"io"
)

type iFaceCompareV2AdvanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFacePictureQualityCheck(v string) *FaceCompareV2AdvanceRequest
	GetFacePictureQualityCheck() *string
	SetMerchantBizId(v string) *FaceCompareV2AdvanceRequest
	GetMerchantBizId() *string
	SetSourceFacePicture(v string) *FaceCompareV2AdvanceRequest
	GetSourceFacePicture() *string
	SetSourceFacePictureFileObject(v io.Reader) *FaceCompareV2AdvanceRequest
	GetSourceFacePictureFileObject() io.Reader
	SetSourceFacePictureUrl(v string) *FaceCompareV2AdvanceRequest
	GetSourceFacePictureUrl() *string
	SetTargetFacePicture(v string) *FaceCompareV2AdvanceRequest
	GetTargetFacePicture() *string
	SetTargetFacePictureFileObject(v io.Reader) *FaceCompareV2AdvanceRequest
	GetTargetFacePictureFileObject() io.Reader
	SetTargetFacePictureUrl(v string) *FaceCompareV2AdvanceRequest
	GetTargetFacePictureUrl() *string
}

type FaceCompareV2AdvanceRequest struct {
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
	SourceFacePictureFileObject io.Reader `json:"SourceFacePictureFile,omitempty" xml:"SourceFacePictureFile,omitempty"`
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
	TargetFacePictureFileObject io.Reader `json:"TargetFacePictureFile,omitempty" xml:"TargetFacePictureFile,omitempty"`
	// example:
	//
	// https://***face2.jpeg
	TargetFacePictureUrl *string `json:"TargetFacePictureUrl,omitempty" xml:"TargetFacePictureUrl,omitempty"`
}

func (s FaceCompareV2AdvanceRequest) String() string {
	return dara.Prettify(s)
}

func (s FaceCompareV2AdvanceRequest) GoString() string {
	return s.String()
}

func (s *FaceCompareV2AdvanceRequest) GetFacePictureQualityCheck() *string {
	return s.FacePictureQualityCheck
}

func (s *FaceCompareV2AdvanceRequest) GetMerchantBizId() *string {
	return s.MerchantBizId
}

func (s *FaceCompareV2AdvanceRequest) GetSourceFacePicture() *string {
	return s.SourceFacePicture
}

func (s *FaceCompareV2AdvanceRequest) GetSourceFacePictureFileObject() io.Reader {
	return s.SourceFacePictureFileObject
}

func (s *FaceCompareV2AdvanceRequest) GetSourceFacePictureUrl() *string {
	return s.SourceFacePictureUrl
}

func (s *FaceCompareV2AdvanceRequest) GetTargetFacePicture() *string {
	return s.TargetFacePicture
}

func (s *FaceCompareV2AdvanceRequest) GetTargetFacePictureFileObject() io.Reader {
	return s.TargetFacePictureFileObject
}

func (s *FaceCompareV2AdvanceRequest) GetTargetFacePictureUrl() *string {
	return s.TargetFacePictureUrl
}

func (s *FaceCompareV2AdvanceRequest) SetFacePictureQualityCheck(v string) *FaceCompareV2AdvanceRequest {
	s.FacePictureQualityCheck = &v
	return s
}

func (s *FaceCompareV2AdvanceRequest) SetMerchantBizId(v string) *FaceCompareV2AdvanceRequest {
	s.MerchantBizId = &v
	return s
}

func (s *FaceCompareV2AdvanceRequest) SetSourceFacePicture(v string) *FaceCompareV2AdvanceRequest {
	s.SourceFacePicture = &v
	return s
}

func (s *FaceCompareV2AdvanceRequest) SetSourceFacePictureFileObject(v io.Reader) *FaceCompareV2AdvanceRequest {
	s.SourceFacePictureFileObject = v
	return s
}

func (s *FaceCompareV2AdvanceRequest) SetSourceFacePictureUrl(v string) *FaceCompareV2AdvanceRequest {
	s.SourceFacePictureUrl = &v
	return s
}

func (s *FaceCompareV2AdvanceRequest) SetTargetFacePicture(v string) *FaceCompareV2AdvanceRequest {
	s.TargetFacePicture = &v
	return s
}

func (s *FaceCompareV2AdvanceRequest) SetTargetFacePictureFileObject(v io.Reader) *FaceCompareV2AdvanceRequest {
	s.TargetFacePictureFileObject = v
	return s
}

func (s *FaceCompareV2AdvanceRequest) SetTargetFacePictureUrl(v string) *FaceCompareV2AdvanceRequest {
	s.TargetFacePictureUrl = &v
	return s
}

func (s *FaceCompareV2AdvanceRequest) Validate() error {
	return dara.Validate(s)
}
