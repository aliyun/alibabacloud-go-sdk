// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCsrDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCsr(v string) *GetCsrDetailResponseBody
	GetCsr() *string
	SetPrivateKey(v string) *GetCsrDetailResponseBody
	GetPrivateKey() *string
	SetRequestId(v string) *GetCsrDetailResponseBody
	GetRequestId() *string
}

type GetCsrDetailResponseBody struct {
	// The content of the CSR.
	//
	// example:
	//
	// -----BEGIN CERTIFICATE REQUEST-----   ...... -----END CERTIFICATE REQUEST-----
	Csr *string `json:"Csr,omitempty" xml:"Csr,omitempty"`
	// The private key. Specify a Base64-encoded string.
	//
	// example:
	//
	// -----BEGIN RSA PRIVATE KEY-----…… -----END RSA PRIVATE KEY-----
	PrivateKey *string `json:"PrivateKey,omitempty" xml:"PrivateKey,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 08F45EA0-66A7-4504-9B31-3589F5CE308D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetCsrDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetCsrDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GetCsrDetailResponseBody) GetCsr() *string {
	return s.Csr
}

func (s *GetCsrDetailResponseBody) GetPrivateKey() *string {
	return s.PrivateKey
}

func (s *GetCsrDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetCsrDetailResponseBody) SetCsr(v string) *GetCsrDetailResponseBody {
	s.Csr = &v
	return s
}

func (s *GetCsrDetailResponseBody) SetPrivateKey(v string) *GetCsrDetailResponseBody {
	s.PrivateKey = &v
	return s
}

func (s *GetCsrDetailResponseBody) SetRequestId(v string) *GetCsrDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCsrDetailResponseBody) Validate() error {
	return dara.Validate(s)
}
