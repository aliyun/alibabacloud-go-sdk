// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateKMSDataKeyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCiphertextBlob(v string) *GenerateKMSDataKeyResponseBody
	GetCiphertextBlob() *string
	SetKeyId(v string) *GenerateKMSDataKeyResponseBody
	GetKeyId() *string
	SetPlaintext(v string) *GenerateKMSDataKeyResponseBody
	GetPlaintext() *string
	SetRequestId(v string) *GenerateKMSDataKeyResponseBody
	GetRequestId() *string
}

type GenerateKMSDataKeyResponseBody struct {
	// The ciphertext of the encrypted data key. This is used as CipherText when you create a transcoding job.
	//
	// example:
	//
	// ODZhOWVmZDktM2QxNi00ODk0LWJkNGYtMWZjNDNmM2YyYWJmS7FmDBBQ0BkKsQrtRnidtPwirmDcS0ZuJCU41xxAAWk4Z8qsADfbV0b+i6kQmlvj79dJdGOvtX69Uycs901qOjop4bTS****
	CiphertextBlob *string `json:"CiphertextBlob,omitempty" xml:"CiphertextBlob,omitempty"`
	// The ID of the customer master key (CMK). The ID must be globally unique.
	//
	// example:
	//
	// 7906979c-8e06-46a2-be2d-68e3ccbc****
	KeyId *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	// The Base64-encoded plaintext of the data key.
	//
	// example:
	//
	// QmFzZTY0IGVuY29kZWQgcGxhaW50****
	Plaintext *string `json:"Plaintext,omitempty" xml:"Plaintext,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 25818875-5F78-4A*****F6-D7393642CA58
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GenerateKMSDataKeyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GenerateKMSDataKeyResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateKMSDataKeyResponseBody) GetCiphertextBlob() *string {
	return s.CiphertextBlob
}

func (s *GenerateKMSDataKeyResponseBody) GetKeyId() *string {
	return s.KeyId
}

func (s *GenerateKMSDataKeyResponseBody) GetPlaintext() *string {
	return s.Plaintext
}

func (s *GenerateKMSDataKeyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GenerateKMSDataKeyResponseBody) SetCiphertextBlob(v string) *GenerateKMSDataKeyResponseBody {
	s.CiphertextBlob = &v
	return s
}

func (s *GenerateKMSDataKeyResponseBody) SetKeyId(v string) *GenerateKMSDataKeyResponseBody {
	s.KeyId = &v
	return s
}

func (s *GenerateKMSDataKeyResponseBody) SetPlaintext(v string) *GenerateKMSDataKeyResponseBody {
	s.Plaintext = &v
	return s
}

func (s *GenerateKMSDataKeyResponseBody) SetRequestId(v string) *GenerateKMSDataKeyResponseBody {
	s.RequestId = &v
	return s
}

func (s *GenerateKMSDataKeyResponseBody) Validate() error {
	return dara.Validate(s)
}
