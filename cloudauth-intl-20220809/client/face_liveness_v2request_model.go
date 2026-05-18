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
	// example:
	//
	// Base64
	FacePictureBase64 *string `json:"FacePictureBase64,omitempty" xml:"FacePictureBase64,omitempty"`
	// example:
	//
	// InputStream
	FacePictureFile *string `json:"FacePictureFile,omitempty" xml:"FacePictureFile,omitempty"`
	// example:
	//
	// https://***
	FacePictureUrl *string `json:"FacePictureUrl,omitempty" xml:"FacePictureUrl,omitempty"`
	// example:
	//
	// Y
	FaceQualityCheck *string `json:"FaceQualityCheck,omitempty" xml:"FaceQualityCheck,omitempty"`
	// example:
	//
	// e0c34a***353888
	MerchantBizId *string `json:"MerchantBizId,omitempty" xml:"MerchantBizId,omitempty"`
	// example:
	//
	// 123456789
	MerchantUserId *string `json:"MerchantUserId,omitempty" xml:"MerchantUserId,omitempty"`
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
