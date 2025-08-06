// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDecryptKMSDataKeyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetKeyId(v string) *DecryptKMSDataKeyResponseBody
	GetKeyId() *string
	SetPlaintext(v string) *DecryptKMSDataKeyResponseBody
	GetPlaintext() *string
	SetRequestId(v string) *DecryptKMSDataKeyResponseBody
	GetRequestId() *string
}

type DecryptKMSDataKeyResponseBody struct {
	// The ID of the customer master key (CMK) that was used to decrypt the ciphertext.
	//
	// example:
	//
	// 202b9877-5a25-46e3-a763-e20791b5****
	KeyId *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	// The plaintext that is generated after decryption.
	//
	// example:
	//
	// tRYXuCwgja12xxO1N/gZERDDCLw9doZEQiPDk/Bv****
	Plaintext *string `json:"Plaintext,omitempty" xml:"Plaintext,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 25818875-5F78-4A*****F6-D7393642CA58
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DecryptKMSDataKeyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DecryptKMSDataKeyResponseBody) GoString() string {
	return s.String()
}

func (s *DecryptKMSDataKeyResponseBody) GetKeyId() *string {
	return s.KeyId
}

func (s *DecryptKMSDataKeyResponseBody) GetPlaintext() *string {
	return s.Plaintext
}

func (s *DecryptKMSDataKeyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DecryptKMSDataKeyResponseBody) SetKeyId(v string) *DecryptKMSDataKeyResponseBody {
	s.KeyId = &v
	return s
}

func (s *DecryptKMSDataKeyResponseBody) SetPlaintext(v string) *DecryptKMSDataKeyResponseBody {
	s.Plaintext = &v
	return s
}

func (s *DecryptKMSDataKeyResponseBody) SetRequestId(v string) *DecryptKMSDataKeyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DecryptKMSDataKeyResponseBody) Validate() error {
	return dara.Validate(s)
}
