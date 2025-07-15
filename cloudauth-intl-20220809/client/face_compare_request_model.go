// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFaceCompareRequest interface {
	dara.Model
	String() string
	GoString() string
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
	// example:
	//
	// e0c34a77f5ac40a5aa5e6ed20c353888
	MerchantBizId     *string `json:"MerchantBizId,omitempty" xml:"MerchantBizId,omitempty"`
	SourceFacePicture *string `json:"SourceFacePicture,omitempty" xml:"SourceFacePicture,omitempty"`
	// example:
	//
	// https://***face1.jpeg
	SourceFacePictureUrl *string `json:"SourceFacePictureUrl,omitempty" xml:"SourceFacePictureUrl,omitempty"`
	TargetFacePicture    *string `json:"TargetFacePicture,omitempty" xml:"TargetFacePicture,omitempty"`
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
