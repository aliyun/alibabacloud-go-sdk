// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateDataKeyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCiphertextBlob(v string) *GenerateDataKeyResponseBody
	GetCiphertextBlob() *string
	SetKeyId(v string) *GenerateDataKeyResponseBody
	GetKeyId() *string
	SetKeyVersionId(v string) *GenerateDataKeyResponseBody
	GetKeyVersionId() *string
	SetPlaintext(v string) *GenerateDataKeyResponseBody
	GetPlaintext() *string
	SetRequestId(v string) *GenerateDataKeyResponseBody
	GetRequestId() *string
}

type GenerateDataKeyResponseBody struct {
	// The ciphertext of the data key that is encrypted by using the primary version of the specified CMK.
	//
	// example:
	//
	// ODZhOWVmZDktM2QxNi00ODk0LWJkNGYtMWZjNDNmM2YyYWJmS7FmDBBQ0BkKsQrtRnidtPwirmDcS0ZuJCU41xxAAWk4Z8qsADfbV0b+i6kQmlvj79dJdGOvtX69Uycs901qOjop4bTS****
	CiphertextBlob *string `json:"CiphertextBlob,omitempty" xml:"CiphertextBlob,omitempty"`
	// The ID of the CMK. The ID must be globally unique.
	//
	// >  If you set the KeyId parameter in the request to an alias of the CMK, the ID of the CMK to which the alias is bound is returned.
	//
	// example:
	//
	// 7906979c-8e06-46a2-be2d-68e3ccbc****
	KeyId *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	// The ID of the CMK version. The ID must be globally unique.
	//
	// example:
	//
	// 2ab1a983-7072-4bbc-a582-584b5bd8****
	KeyVersionId *string `json:"KeyVersionId,omitempty" xml:"KeyVersionId,omitempty"`
	// The Base64 encoded plaintext of the data key.
	//
	// example:
	//
	// QmFzZTY0IGVuY29kZWQgcGxhaW50****
	Plaintext *string `json:"Plaintext,omitempty" xml:"Plaintext,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 7021b6ec-4be7-4d3c-8a68-1e85d4d515a0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GenerateDataKeyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GenerateDataKeyResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateDataKeyResponseBody) GetCiphertextBlob() *string {
	return s.CiphertextBlob
}

func (s *GenerateDataKeyResponseBody) GetKeyId() *string {
	return s.KeyId
}

func (s *GenerateDataKeyResponseBody) GetKeyVersionId() *string {
	return s.KeyVersionId
}

func (s *GenerateDataKeyResponseBody) GetPlaintext() *string {
	return s.Plaintext
}

func (s *GenerateDataKeyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GenerateDataKeyResponseBody) SetCiphertextBlob(v string) *GenerateDataKeyResponseBody {
	s.CiphertextBlob = &v
	return s
}

func (s *GenerateDataKeyResponseBody) SetKeyId(v string) *GenerateDataKeyResponseBody {
	s.KeyId = &v
	return s
}

func (s *GenerateDataKeyResponseBody) SetKeyVersionId(v string) *GenerateDataKeyResponseBody {
	s.KeyVersionId = &v
	return s
}

func (s *GenerateDataKeyResponseBody) SetPlaintext(v string) *GenerateDataKeyResponseBody {
	s.Plaintext = &v
	return s
}

func (s *GenerateDataKeyResponseBody) SetRequestId(v string) *GenerateDataKeyResponseBody {
	s.RequestId = &v
	return s
}

func (s *GenerateDataKeyResponseBody) Validate() error {
	return dara.Validate(s)
}
