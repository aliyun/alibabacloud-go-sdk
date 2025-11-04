// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateKMSDataKeyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataKey(v *GenerateKMSDataKeyResponseBodyDataKey) *GenerateKMSDataKeyResponseBody
	GetDataKey() *GenerateKMSDataKeyResponseBodyDataKey
	SetRequestId(v string) *GenerateKMSDataKeyResponseBody
	GetRequestId() *string
}

type GenerateKMSDataKeyResponseBody struct {
	// The information about the data key.
	DataKey *GenerateKMSDataKeyResponseBodyDataKey `json:"DataKey,omitempty" xml:"DataKey,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// ******11-DB8D-4A9A-875B-275798******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GenerateKMSDataKeyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GenerateKMSDataKeyResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateKMSDataKeyResponseBody) GetDataKey() *GenerateKMSDataKeyResponseBodyDataKey {
	return s.DataKey
}

func (s *GenerateKMSDataKeyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GenerateKMSDataKeyResponseBody) SetDataKey(v *GenerateKMSDataKeyResponseBodyDataKey) *GenerateKMSDataKeyResponseBody {
	s.DataKey = v
	return s
}

func (s *GenerateKMSDataKeyResponseBody) SetRequestId(v string) *GenerateKMSDataKeyResponseBody {
	s.RequestId = &v
	return s
}

func (s *GenerateKMSDataKeyResponseBody) Validate() error {
	if s.DataKey != nil {
		if err := s.DataKey.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GenerateKMSDataKeyResponseBodyDataKey struct {
	// The ciphertext of the encrypted data key. This parameter is used as CipherText when you create a transcoding job.
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
}

func (s GenerateKMSDataKeyResponseBodyDataKey) String() string {
	return dara.Prettify(s)
}

func (s GenerateKMSDataKeyResponseBodyDataKey) GoString() string {
	return s.String()
}

func (s *GenerateKMSDataKeyResponseBodyDataKey) GetCiphertextBlob() *string {
	return s.CiphertextBlob
}

func (s *GenerateKMSDataKeyResponseBodyDataKey) GetKeyId() *string {
	return s.KeyId
}

func (s *GenerateKMSDataKeyResponseBodyDataKey) GetPlaintext() *string {
	return s.Plaintext
}

func (s *GenerateKMSDataKeyResponseBodyDataKey) SetCiphertextBlob(v string) *GenerateKMSDataKeyResponseBodyDataKey {
	s.CiphertextBlob = &v
	return s
}

func (s *GenerateKMSDataKeyResponseBodyDataKey) SetKeyId(v string) *GenerateKMSDataKeyResponseBodyDataKey {
	s.KeyId = &v
	return s
}

func (s *GenerateKMSDataKeyResponseBodyDataKey) SetPlaintext(v string) *GenerateKMSDataKeyResponseBodyDataKey {
	s.Plaintext = &v
	return s
}

func (s *GenerateKMSDataKeyResponseBodyDataKey) Validate() error {
	return dara.Validate(s)
}
