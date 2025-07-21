// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAsymmetricSignRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlgorithm(v string) *AsymmetricSignRequest
	GetAlgorithm() *string
	SetDigest(v string) *AsymmetricSignRequest
	GetDigest() *string
	SetDryRun(v string) *AsymmetricSignRequest
	GetDryRun() *string
	SetKeyId(v string) *AsymmetricSignRequest
	GetKeyId() *string
	SetKeyVersionId(v string) *AsymmetricSignRequest
	GetKeyVersionId() *string
}

type AsymmetricSignRequest struct {
	// The version ID of the CMK. The ID must be globally unique.
	//
	// This parameter is required.
	//
	// example:
	//
	// RSA_PSS_SHA_256
	Algorithm *string `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	// The signature algorithm.
	//
	// This parameter is required.
	//
	// example:
	//
	// ZOyIygCyaOW6GjVnihtTFtIS9PNmskdyMlNKiu****=
	Digest *string `json:"Digest,omitempty" xml:"Digest,omitempty"`
	DryRun *string `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The operation that you want to perform. Set the value to **AsymmetricSign**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 5c438b18-05be-40ad-b6c2-3be6752c****
	KeyId *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	// The ID of the customer master key (CMK). The ID must be globally unique.
	//
	// >  You can also set this parameter to an alias that is bound to the CMK. For more information, see [Alias overview](https://help.aliyun.com/document_detail/68522.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// 2ab1a983-7072-4bbc-a582-584b5bd8****
	KeyVersionId *string `json:"KeyVersionId,omitempty" xml:"KeyVersionId,omitempty"`
}

func (s AsymmetricSignRequest) String() string {
	return dara.Prettify(s)
}

func (s AsymmetricSignRequest) GoString() string {
	return s.String()
}

func (s *AsymmetricSignRequest) GetAlgorithm() *string {
	return s.Algorithm
}

func (s *AsymmetricSignRequest) GetDigest() *string {
	return s.Digest
}

func (s *AsymmetricSignRequest) GetDryRun() *string {
	return s.DryRun
}

func (s *AsymmetricSignRequest) GetKeyId() *string {
	return s.KeyId
}

func (s *AsymmetricSignRequest) GetKeyVersionId() *string {
	return s.KeyVersionId
}

func (s *AsymmetricSignRequest) SetAlgorithm(v string) *AsymmetricSignRequest {
	s.Algorithm = &v
	return s
}

func (s *AsymmetricSignRequest) SetDigest(v string) *AsymmetricSignRequest {
	s.Digest = &v
	return s
}

func (s *AsymmetricSignRequest) SetDryRun(v string) *AsymmetricSignRequest {
	s.DryRun = &v
	return s
}

func (s *AsymmetricSignRequest) SetKeyId(v string) *AsymmetricSignRequest {
	s.KeyId = &v
	return s
}

func (s *AsymmetricSignRequest) SetKeyVersionId(v string) *AsymmetricSignRequest {
	s.KeyVersionId = &v
	return s
}

func (s *AsymmetricSignRequest) Validate() error {
	return dara.Validate(s)
}
