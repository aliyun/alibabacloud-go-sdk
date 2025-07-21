// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateDataKeyWithoutPlaintextResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCiphertextBlob(v string) *GenerateDataKeyWithoutPlaintextResponseBody
	GetCiphertextBlob() *string
	SetKeyId(v string) *GenerateDataKeyWithoutPlaintextResponseBody
	GetKeyId() *string
	SetKeyVersionId(v string) *GenerateDataKeyWithoutPlaintextResponseBody
	GetKeyVersionId() *string
	SetRequestId(v string) *GenerateDataKeyWithoutPlaintextResponseBody
	GetRequestId() *string
}

type GenerateDataKeyWithoutPlaintextResponseBody struct {
	// The ciphertext of the data that is encrypted by using the primary CMK version.
	//
	// example:
	//
	// ODZhOWVmZDktM2QxNi00ODk0LWJkNGYtMWZjNDNmM2YyYWJmS7FmDBBQ0BkKsQrtRnidtPwirmDcS0ZuJCU41xxAAWk4Z8qsADfbV0b+i6kQmlvj79dJdGOvtX69Uycs901qOjop4bTS****
	CiphertextBlob *string `json:"CiphertextBlob,omitempty" xml:"CiphertextBlob,omitempty"`
	// The globally unique ID of the CMK.
	//
	// >  If you set the KeyId parameter to an alias, the ID of the CMK to which the alias is bound is returned.
	//
	// example:
	//
	// 599fa825-17de-417e-9554-bb032cc6****
	KeyId *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	// The ID of the key version that is used to encrypt the plaintext. It is the primary version of the CMK.
	//
	// example:
	//
	// 2ab1a983-7072-4bbc-a582-584b5bd8****
	KeyVersionId *string `json:"KeyVersionId,omitempty" xml:"KeyVersionId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 7021b6ec-4be7-4d3c-8a68-1e85d4d515a0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GenerateDataKeyWithoutPlaintextResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GenerateDataKeyWithoutPlaintextResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateDataKeyWithoutPlaintextResponseBody) GetCiphertextBlob() *string {
	return s.CiphertextBlob
}

func (s *GenerateDataKeyWithoutPlaintextResponseBody) GetKeyId() *string {
	return s.KeyId
}

func (s *GenerateDataKeyWithoutPlaintextResponseBody) GetKeyVersionId() *string {
	return s.KeyVersionId
}

func (s *GenerateDataKeyWithoutPlaintextResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GenerateDataKeyWithoutPlaintextResponseBody) SetCiphertextBlob(v string) *GenerateDataKeyWithoutPlaintextResponseBody {
	s.CiphertextBlob = &v
	return s
}

func (s *GenerateDataKeyWithoutPlaintextResponseBody) SetKeyId(v string) *GenerateDataKeyWithoutPlaintextResponseBody {
	s.KeyId = &v
	return s
}

func (s *GenerateDataKeyWithoutPlaintextResponseBody) SetKeyVersionId(v string) *GenerateDataKeyWithoutPlaintextResponseBody {
	s.KeyVersionId = &v
	return s
}

func (s *GenerateDataKeyWithoutPlaintextResponseBody) SetRequestId(v string) *GenerateDataKeyWithoutPlaintextResponseBody {
	s.RequestId = &v
	return s
}

func (s *GenerateDataKeyWithoutPlaintextResponseBody) Validate() error {
	return dara.Validate(s)
}
