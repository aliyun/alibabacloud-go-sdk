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
	// The digest of the original message. The digest is generated using the hash algorithm that corresponds to the value of the Algorithm parameter.
	//
	// > - The value is Base64-encoded.
	//
	// - For information about how to calculate a message digest, see the "Pre-signing: calculate a message digest" section of the [Asymmetric digital signature](https://help.aliyun.com/document_detail/148146.html) topic.
	//
	// This parameter is required.
	//
	// example:
	//
	// ZOyIygCyaOW6GjVnihtTFtIS9PNmskdyMlNKiu****=
	Digest *string `json:"Digest,omitempty" xml:"Digest,omitempty"`
	// Specifies whether to enable the dry-run feature.
	//
	// - true: enables the feature.
	//
	// - false (default): disables the feature.
	//
	// The dry-run feature is used to test API calls and verify the permissions on the resources that you have and the validity of the request parameters. If you enable the dry-run feature, KMS always returns a failure response and a failure reason. The failure reasons include the following:
	//
	// - DryRunOperationError: The request would have succeeded if the DryRun parameter is not configured.
	//
	// - ValidationError: The specified parameters in the request are invalid.
	//
	// - AccessDeniedError: You are not authorized to perform the operation on the KMS resource.
	//
	// example:
	//
	// false
	DryRun *string `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The globally unique identifier (GUID) of the customer master key (CMK).
	//
	// > You can also specify the alias that is bound to the CMK. For more information, see [Overview of aliases](https://help.aliyun.com/document_detail/68522.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// 5c438b18-05be-40ad-b6c2-3be6752c****
	KeyId *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	// The ID of the key version. The ID must be the GUID of the key version.
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
