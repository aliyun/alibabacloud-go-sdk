// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAsymmetricDecryptResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetKeyId(v string) *AsymmetricDecryptResponseBody
	GetKeyId() *string
	SetKeyVersionId(v string) *AsymmetricDecryptResponseBody
	GetKeyVersionId() *string
	SetPlaintext(v string) *AsymmetricDecryptResponseBody
	GetPlaintext() *string
	SetRequestId(v string) *AsymmetricDecryptResponseBody
	GetRequestId() *string
}

type AsymmetricDecryptResponseBody struct {
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
	// The Base64-encoded plaintext that is generated after decryption.
	//
	// example:
	//
	// SGVsbG8gd29ybGQ=
	Plaintext *string `json:"Plaintext,omitempty" xml:"Plaintext,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 475f1620-b9d3-4d35-b5c6-3fbdd941423d
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AsymmetricDecryptResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AsymmetricDecryptResponseBody) GoString() string {
	return s.String()
}

func (s *AsymmetricDecryptResponseBody) GetKeyId() *string {
	return s.KeyId
}

func (s *AsymmetricDecryptResponseBody) GetKeyVersionId() *string {
	return s.KeyVersionId
}

func (s *AsymmetricDecryptResponseBody) GetPlaintext() *string {
	return s.Plaintext
}

func (s *AsymmetricDecryptResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AsymmetricDecryptResponseBody) SetKeyId(v string) *AsymmetricDecryptResponseBody {
	s.KeyId = &v
	return s
}

func (s *AsymmetricDecryptResponseBody) SetKeyVersionId(v string) *AsymmetricDecryptResponseBody {
	s.KeyVersionId = &v
	return s
}

func (s *AsymmetricDecryptResponseBody) SetPlaintext(v string) *AsymmetricDecryptResponseBody {
	s.Plaintext = &v
	return s
}

func (s *AsymmetricDecryptResponseBody) SetRequestId(v string) *AsymmetricDecryptResponseBody {
	s.RequestId = &v
	return s
}

func (s *AsymmetricDecryptResponseBody) Validate() error {
	return dara.Validate(s)
}
