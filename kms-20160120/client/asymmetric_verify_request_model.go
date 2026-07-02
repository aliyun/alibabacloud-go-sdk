// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAsymmetricVerifyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlgorithm(v string) *AsymmetricVerifyRequest
	GetAlgorithm() *string
	SetDigest(v string) *AsymmetricVerifyRequest
	GetDigest() *string
	SetDryRun(v string) *AsymmetricVerifyRequest
	GetDryRun() *string
	SetKeyId(v string) *AsymmetricVerifyRequest
	GetKeyId() *string
	SetKeyVersionId(v string) *AsymmetricVerifyRequest
	GetKeyVersionId() *string
	SetValue(v string) *AsymmetricVerifyRequest
	GetValue() *string
}

type AsymmetricVerifyRequest struct {
	// The signature algorithm.
	//
	// This parameter is required.
	//
	// example:
	//
	// RSA_PSS_SHA_256
	Algorithm *string `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	// The digest that is generated using the hash algorithm that corresponds to the value of **Algorithm*	- to hash the original message.
	//
	// > The value is Base64-encoded.
	//
	// This parameter is required.
	//
	// example:
	//
	// ZOyIygCyaOW6GjVnihtTFtIS9PNmskdyMlNKiuy****=
	Digest *string `json:"Digest,omitempty" xml:"Digest,omitempty"`
	// Specifies whether to perform a dry run.
	//
	// - true: performs a dry run.
	//
	// - false (default): does not perform a dry run.
	//
	// A dry run is used to test API calls and verify whether you have the permissions to access the specified resources and whether the request parameters are valid. If you perform a dry run, KMS always returns a failure response that indicates the cause of the failure. The following failure causes are included:
	//
	// - DryRunOperationError: The request would have succeeded if the DryRun parameter is not specified.
	//
	// - ValidationError: The specified parameters in the request are invalid.
	//
	// - AccessDeniedError: You are not authorized to perform this operation on the KMS resource.
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
	// The signature value to be verified.
	//
	// > The value is Base64-encoded.
	//
	// This parameter is required.
	//
	// example:
	//
	// M2CceNZH00ZgL9ED/ZHFp21YRAvYeZHknJUc207OCZ0N9wNn9As4z2bON3FF3je+1Nu+2+/8Zj50HpMTpzYpMp2R93cYmACCmhaYoKydxylbyGzJR8y9likZRCrkD38lRoS40aBBvv/6iRKzQuo9EGYVcel36cMNg00VmYNBy3pa1rwg3gA4l3cy6kjayZja1WGPkVhrVKsrJMdbpl0ApLjXKuD8rw1n1XLCwCUEL5eLPljTZaAveqdOFQOiZnZEGI27qIiZe7I1fN8tcz6anS/gTM7xRKE++5egEvRWlTQQTJeApnPSiUPA+8ZykNdelQsOQh5SrGoyI4A5pq****==
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s AsymmetricVerifyRequest) String() string {
	return dara.Prettify(s)
}

func (s AsymmetricVerifyRequest) GoString() string {
	return s.String()
}

func (s *AsymmetricVerifyRequest) GetAlgorithm() *string {
	return s.Algorithm
}

func (s *AsymmetricVerifyRequest) GetDigest() *string {
	return s.Digest
}

func (s *AsymmetricVerifyRequest) GetDryRun() *string {
	return s.DryRun
}

func (s *AsymmetricVerifyRequest) GetKeyId() *string {
	return s.KeyId
}

func (s *AsymmetricVerifyRequest) GetKeyVersionId() *string {
	return s.KeyVersionId
}

func (s *AsymmetricVerifyRequest) GetValue() *string {
	return s.Value
}

func (s *AsymmetricVerifyRequest) SetAlgorithm(v string) *AsymmetricVerifyRequest {
	s.Algorithm = &v
	return s
}

func (s *AsymmetricVerifyRequest) SetDigest(v string) *AsymmetricVerifyRequest {
	s.Digest = &v
	return s
}

func (s *AsymmetricVerifyRequest) SetDryRun(v string) *AsymmetricVerifyRequest {
	s.DryRun = &v
	return s
}

func (s *AsymmetricVerifyRequest) SetKeyId(v string) *AsymmetricVerifyRequest {
	s.KeyId = &v
	return s
}

func (s *AsymmetricVerifyRequest) SetKeyVersionId(v string) *AsymmetricVerifyRequest {
	s.KeyVersionId = &v
	return s
}

func (s *AsymmetricVerifyRequest) SetValue(v string) *AsymmetricVerifyRequest {
	s.Value = &v
	return s
}

func (s *AsymmetricVerifyRequest) Validate() error {
	return dara.Validate(s)
}
