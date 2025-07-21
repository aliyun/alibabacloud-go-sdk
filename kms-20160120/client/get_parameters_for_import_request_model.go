// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetParametersForImportRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKeyId(v string) *GetParametersForImportRequest
	GetKeyId() *string
	SetWrappingAlgorithm(v string) *GetParametersForImportRequest
	GetWrappingAlgorithm() *string
	SetWrappingKeySpec(v string) *GetParametersForImportRequest
	GetWrappingKeySpec() *string
}

type GetParametersForImportRequest struct {
	// The globally unique ID of the CMK.
	//
	// >  You can import key material only for CMKs whose Origin parameter is set to EXTERNAL.
	//
	// This parameter is required.
	//
	// example:
	//
	// 202b9877-5a25-46e3-a763-e20791b5****
	KeyId *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	// The algorithm that is used to encrypt key material.
	//
	// This parameter is required.
	//
	// example:
	//
	// RSAES_PKCS1_V1_5
	WrappingAlgorithm *string `json:"WrappingAlgorithm,omitempty" xml:"WrappingAlgorithm,omitempty"`
	// The type of the public key that is used to encrypt key material.
	//
	// This parameter is required.
	//
	// example:
	//
	// RSA_2048
	WrappingKeySpec *string `json:"WrappingKeySpec,omitempty" xml:"WrappingKeySpec,omitempty"`
}

func (s GetParametersForImportRequest) String() string {
	return dara.Prettify(s)
}

func (s GetParametersForImportRequest) GoString() string {
	return s.String()
}

func (s *GetParametersForImportRequest) GetKeyId() *string {
	return s.KeyId
}

func (s *GetParametersForImportRequest) GetWrappingAlgorithm() *string {
	return s.WrappingAlgorithm
}

func (s *GetParametersForImportRequest) GetWrappingKeySpec() *string {
	return s.WrappingKeySpec
}

func (s *GetParametersForImportRequest) SetKeyId(v string) *GetParametersForImportRequest {
	s.KeyId = &v
	return s
}

func (s *GetParametersForImportRequest) SetWrappingAlgorithm(v string) *GetParametersForImportRequest {
	s.WrappingAlgorithm = &v
	return s
}

func (s *GetParametersForImportRequest) SetWrappingKeySpec(v string) *GetParametersForImportRequest {
	s.WrappingKeySpec = &v
	return s
}

func (s *GetParametersForImportRequest) Validate() error {
	return dara.Validate(s)
}
