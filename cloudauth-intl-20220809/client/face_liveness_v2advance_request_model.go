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
	// example:
	//
	// Base64
	FacePictureBase64 *string `json:"FacePictureBase64,omitempty" xml:"FacePictureBase64,omitempty"`
	// example:
	//
	// InputStream
	FacePictureFileObject io.Reader `json:"FacePictureFile,omitempty" xml:"FacePictureFile,omitempty"`
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

func (s FaceLivenessV2AdvanceRequest) String() string {
	return dara.Prettify(s)
}

func (s FaceLivenessV2AdvanceRequest) GoString() string {
	return s.String()
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
