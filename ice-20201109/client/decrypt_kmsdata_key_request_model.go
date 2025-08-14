// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDecryptKMSDataKeyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCiphertextBlob(v string) *DecryptKMSDataKeyRequest
	GetCiphertextBlob() *string
}

type DecryptKMSDataKeyRequest struct {
	// The ciphertext that you want to decrypt.
	//
	// This parameter is required.
	//
	// example:
	//
	// ODZhOWVmZDktM2QxNi00ODk0LWJkNGYtMWZjNDNmM2YyYWJmS7FmDBBQ0BkKsQrtRnidtPwirmDcS0ZuJCU41xxAAWk4Z8qsADfbV0b+i6kQmlvj79dJdGOvtX69Uycs901qOjop4bTS****
	CiphertextBlob *string `json:"CiphertextBlob,omitempty" xml:"CiphertextBlob,omitempty"`
}

func (s DecryptKMSDataKeyRequest) String() string {
	return dara.Prettify(s)
}

func (s DecryptKMSDataKeyRequest) GoString() string {
	return s.String()
}

func (s *DecryptKMSDataKeyRequest) GetCiphertextBlob() *string {
	return s.CiphertextBlob
}

func (s *DecryptKMSDataKeyRequest) SetCiphertextBlob(v string) *DecryptKMSDataKeyRequest {
	s.CiphertextBlob = &v
	return s
}

func (s *DecryptKMSDataKeyRequest) Validate() error {
	return dara.Validate(s)
}
