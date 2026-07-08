// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDecryptResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCertIdentifier(v string) *DecryptResponseBody
	GetCertIdentifier() *string
	SetPlaintext(v string) *DecryptResponseBody
	GetPlaintext() *string
	SetRequestId(v string) *DecryptResponseBody
	GetRequestId() *string
}

type DecryptResponseBody struct {
	// The unique identifier of the certificate.
	//
	// example:
	//
	// 1ef1da5f-38ed-69b3-****-037781890265
	CertIdentifier *string `json:"CertIdentifier,omitempty" xml:"CertIdentifier,omitempty"`
	// The decrypted data.
	//
	// example:
	//
	// VGhlIHF1aWNrIGJyb3duIGZveCBqdW1wcyBvdmVyIHRoZSBsYXp5IGRvZy4
	Plaintext *string `json:"Plaintext,omitempty" xml:"Plaintext,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 5979d897-d69f-4fc9-87dd-f3bb73c40b80
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DecryptResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DecryptResponseBody) GoString() string {
	return s.String()
}

func (s *DecryptResponseBody) GetCertIdentifier() *string {
	return s.CertIdentifier
}

func (s *DecryptResponseBody) GetPlaintext() *string {
	return s.Plaintext
}

func (s *DecryptResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DecryptResponseBody) SetCertIdentifier(v string) *DecryptResponseBody {
	s.CertIdentifier = &v
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
