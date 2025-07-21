// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDecryptResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetKeyId(v string) *DecryptResponseBody
	GetKeyId() *string
	SetKeyVersionId(v string) *DecryptResponseBody
	GetKeyVersionId() *string
	SetPlaintext(v string) *DecryptResponseBody
	GetPlaintext() *string
	SetRequestId(v string) *DecryptResponseBody
	GetRequestId() *string
}

type DecryptResponseBody struct {
	// The ID of the customer master key (CMK) that is used to decrypt the ciphertext.
	//
	// It is the GUID of the CMK.
	//
	// example:
	//
	// 202b9877-5a25-46e3-a763-e20791b5****
	KeyId *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	// The ID of the CMK version that is used to decrypt the ciphertext.
	//
	// example:
	//
	// 2ab1a983-7072-4bbc-a582-584b5bd8****
	KeyVersionId *string `json:"KeyVersionId,omitempty" xml:"KeyVersionId,omitempty"`
	// The plaintext that is generated after decryption.
	//
	// example:
	//
	// tRYXuCwgja12xxO1N/gZERDDCLw9doZEQiPDk/Bv****
	Plaintext *string `json:"Plaintext,omitempty" xml:"Plaintext,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 207596a2-36d3-4840-b1bd-f87044699bd7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DecryptResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DecryptResponseBody) GoString() string {
	return s.String()
}

func (s *DecryptResponseBody) GetKeyId() *string {
	return s.KeyId
}

func (s *DecryptResponseBody) GetKeyVersionId() *string {
	return s.KeyVersionId
}

func (s *DecryptResponseBody) GetPlaintext() *string {
	return s.Plaintext
}

func (s *DecryptResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DecryptResponseBody) SetKeyId(v string) *DecryptResponseBody {
	s.KeyId = &v
	return s
}

func (s *DecryptResponseBody) SetKeyVersionId(v string) *DecryptResponseBody {
	s.KeyVersionId = &v
	return s
}

func (s *DecryptResponseBody) SetPlaintext(v string) *DecryptResponseBody {
	s.Plaintext = &v
	return s
}

func (s *DecryptResponseBody) SetRequestId(v string) *DecryptResponseBody {
	s.RequestId = &v
	return s
}

func (s *DecryptResponseBody) Validate() error {
	return dara.Validate(s)
}
