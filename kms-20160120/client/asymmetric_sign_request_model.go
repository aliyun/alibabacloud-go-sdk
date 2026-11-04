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
	// The signature algorithm.
	//
	// This parameter is required.
	//
	// example:
	//
	// RSA_PSS_SHA_256
	Algorithm *string `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	// The digest generated from the original message by using the hash algorithm that corresponds to the value of the Algorithm parameter.
	//
	// > - The value is Base64-encoded.
	//
	// - For more information about how to calculate a message digest, see the **Signature preprocessing: Calculate the message digest*	- section in [Asymmetric digital signatures](https://help.aliyun.com/document_detail/148146.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// ZOyIygCyaOW6GjVnihtTFtIS9PNmskdyMlNKiu****=
	Digest *string `json:"Digest,omitempty" xml:"Digest,omitempty"`
	// Specifies whether to enable the dry run mode.
	//
	// - true: Enables the dry run mode.
	//
	// - false: Disables the dry run mode. This is the default value.
	//
	// The dry run mode is used to test API calls, verify whether you have the required permissions on the resources, and check whether the request parameters are correctly configured. When the dry run mode is enabled, KMS always returns a failure and provides the cause of the failure. The causes include:
	//
	// - DryRunOperationError: The request would succeed if the DryRun parameter is not configured.
	//
	// - ValidationError: The specified parameters in the request are invalid.
	//
	// - AccessDeniedError: You are not authorized to perform the operation on the KMS resources.
	//
	// example:
	//
	// false
	DryRun *string `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The globally unique identifier of the customer master key (CMK).
	//
	// > You can also set this parameter to the alias that is bound to the CMK. For more information, see [Usage notes for aliases](https://help.aliyun.com/document_detail/68522.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// 5c438b18-05be-40ad-b6c2-3be6752c****
	KeyId *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	// The key version ID, which is the globally unique identifier of the key version.
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
