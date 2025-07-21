// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAsymmetricEncryptRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlgorithm(v string) *AsymmetricEncryptRequest
	GetAlgorithm() *string
	SetDryRun(v string) *AsymmetricEncryptRequest
	GetDryRun() *string
	SetKeyId(v string) *AsymmetricEncryptRequest
	GetKeyId() *string
	SetKeyVersionId(v string) *AsymmetricEncryptRequest
	GetKeyVersionId() *string
	SetPlaintext(v string) *AsymmetricEncryptRequest
	GetPlaintext() *string
}

type AsymmetricEncryptRequest struct {
	// The encryption algorithm.
	//
	// This parameter is required.
	//
	// example:
	//
	// RSAES_OAEP_SHA_1
	Algorithm *string `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	DryRun    *string `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ID of the CMK. The ID must be globally unique.
	//
	// >  You can also set this parameter to an alias that is bound to the CMK. For more information, see [Overview of aliases](https://help.aliyun.com/document_detail/68522.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// 5c438b18-05be-40ad-b6c2-3be6752c****
	KeyId *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	// The version ID of the CMK. The ID must be globally unique.
	//
	// >  You can call the [ListKeyVersions](https://help.aliyun.com/document_detail/133966.html) operation to query the versions of a CMK. The ID of a version is specified by the KeyVersionId parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2ab1a983-7072-4bbc-a582-584b5bd8****
	KeyVersionId *string `json:"KeyVersionId,omitempty" xml:"KeyVersionId,omitempty"`
	// The plaintext that you want to encrypt. The plaintext must be Base64-encoded.
	//
	// This parameter is required.
	//
	// example:
	//
	// SGVsbG8gd29ybGQ=
	Plaintext *string `json:"Plaintext,omitempty" xml:"Plaintext,omitempty"`
}

func (s AsymmetricEncryptRequest) String() string {
	return dara.Prettify(s)
}

func (s AsymmetricEncryptRequest) GoString() string {
	return s.String()
}

func (s *AsymmetricEncryptRequest) GetAlgorithm() *string {
	return s.Algorithm
}

func (s *AsymmetricEncryptRequest) GetDryRun() *string {
	return s.DryRun
}

func (s *AsymmetricEncryptRequest) GetKeyId() *string {
	return s.KeyId
}

func (s *AsymmetricEncryptRequest) GetKeyVersionId() *string {
	return s.KeyVersionId
}

func (s *AsymmetricEncryptRequest) GetPlaintext() *string {
	return s.Plaintext
}

func (s *AsymmetricEncryptRequest) SetAlgorithm(v string) *AsymmetricEncryptRequest {
	s.Algorithm = &v
	return s
}

func (s *AsymmetricEncryptRequest) SetDryRun(v string) *AsymmetricEncryptRequest {
	s.DryRun = &v
	return s
}

func (s *AsymmetricEncryptRequest) SetKeyId(v string) *AsymmetricEncryptRequest {
	s.KeyId = &v
	return s
}

func (s *AsymmetricEncryptRequest) SetKeyVersionId(v string) *AsymmetricEncryptRequest {
	s.KeyVersionId = &v
	return s
}

func (s *AsymmetricEncryptRequest) SetPlaintext(v string) *AsymmetricEncryptRequest {
	s.Plaintext = &v
	return s
}

func (s *AsymmetricEncryptRequest) Validate() error {
	return dara.Validate(s)
}
