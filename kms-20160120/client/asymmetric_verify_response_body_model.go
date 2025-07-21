// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAsymmetricVerifyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetKeyId(v string) *AsymmetricVerifyResponseBody
	GetKeyId() *string
	SetKeyVersionId(v string) *AsymmetricVerifyResponseBody
	GetKeyVersionId() *string
	SetRequestId(v string) *AsymmetricVerifyResponseBody
	GetRequestId() *string
	SetValue(v bool) *AsymmetricVerifyResponseBody
	GetValue() *bool
}

type AsymmetricVerifyResponseBody struct {
	// The ID of the CMK. The ID must be globally unique.
	//
	// >  If you set the KeyId parameter in the request to an alias, the ID of the CMK to which the alias is bound is returned.
	//
	// example:
	//
	// 5c438b18-05be-40ad-b6c2-3be6752c****
	KeyId *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	// The version ID of the CMK that is used to encrypt the plaintext.
	//
	// example:
	//
	// 2ab1a983-7072-4bbc-a582-584b5bd8****
	KeyVersionId *string `json:"KeyVersionId,omitempty" xml:"KeyVersionId,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 475f1620-b9d3-4d35-b5c6-3fbdd941423d
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the signature passed the verification.
	//
	// example:
	//
	// true
	Value *bool `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s AsymmetricVerifyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AsymmetricVerifyResponseBody) GoString() string {
	return s.String()
}

func (s *AsymmetricVerifyResponseBody) GetKeyId() *string {
	return s.KeyId
}

func (s *AsymmetricVerifyResponseBody) GetKeyVersionId() *string {
	return s.KeyVersionId
}

func (s *AsymmetricVerifyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AsymmetricVerifyResponseBody) GetValue() *bool {
	return s.Value
}

func (s *AsymmetricVerifyResponseBody) SetKeyId(v string) *AsymmetricVerifyResponseBody {
	s.KeyId = &v
	return s
}

func (s *AsymmetricVerifyResponseBody) SetKeyVersionId(v string) *AsymmetricVerifyResponseBody {
	s.KeyVersionId = &v
	return s
}

func (s *AsymmetricVerifyResponseBody) SetRequestId(v string) *AsymmetricVerifyResponseBody {
	s.RequestId = &v
	return s
}

func (s *AsymmetricVerifyResponseBody) SetValue(v bool) *AsymmetricVerifyResponseBody {
	s.Value = &v
	return s
}

func (s *AsymmetricVerifyResponseBody) Validate() error {
	return dara.Validate(s)
}
