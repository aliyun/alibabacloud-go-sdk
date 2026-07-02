// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateMacResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAlgorithm(v string) *GenerateMacResponseBody
	GetAlgorithm() *string
	SetKeyId(v string) *GenerateMacResponseBody
	GetKeyId() *string
	SetMac(v string) *GenerateMacResponseBody
	GetMac() *string
	SetRequestId(v string) *GenerateMacResponseBody
	GetRequestId() *string
}

type GenerateMacResponseBody struct {
	// The algorithm that is used to generate the message authentication code. Valid values vary based on the key specification:
	//
	// - HMAC_SM3
	//
	// - HMAC_SHA_224
	//
	// - HMAC_SHA_256
	//
	// - HMAC_SHA_384
	//
	// - HMAC_SHA_512
	//
	// example:
	//
	// HMAC_SHA_256
	Algorithm *string `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	// The globally unique identifier of the customer master key (CMK).
	//
	// > If the KeyId parameter in the request uses a CMK alias, the response returns the CMK identifier that corresponds to the alias.
	//
	// example:
	//
	// key-hzz630494463ejqjx****
	KeyId *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	// The Base64-encoded message authenticate code.
	//
	// example:
	//
	// vz1Snp+jGJbgydCFRWVWxAwIMdyfKCSp+jnMWQ==
	Mac *string `json:"Mac,omitempty" xml:"Mac,omitempty"`
	// The ID of the request. Alibaba Cloud generates a unique identifier for each request. You can use this ID to troubleshoot issues.
	//
	// example:
	//
	// 4c8ae23f-3a42-6791-a4ba-1faa77831c28
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GenerateMacResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GenerateMacResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateMacResponseBody) GetAlgorithm() *string {
	return s.Algorithm
}

func (s *GenerateMacResponseBody) GetKeyId() *string {
	return s.KeyId
}

func (s *GenerateMacResponseBody) GetMac() *string {
	return s.Mac
}

func (s *GenerateMacResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GenerateMacResponseBody) SetAlgorithm(v string) *GenerateMacResponseBody {
	s.Algorithm = &v
	return s
}

func (s *GenerateMacResponseBody) SetKeyId(v string) *GenerateMacResponseBody {
	s.KeyId = &v
	return s
}

func (s *GenerateMacResponseBody) SetMac(v string) *GenerateMacResponseBody {
	s.Mac = &v
	return s
}

func (s *GenerateMacResponseBody) SetRequestId(v string) *GenerateMacResponseBody {
	s.RequestId = &v
	return s
}

func (s *GenerateMacResponseBody) Validate() error {
	return dara.Validate(s)
}
