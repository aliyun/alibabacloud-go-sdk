// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCaCertResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCaCert(v string) *GetCaCertResponseBody
	GetCaCert() *string
	SetRequestId(v string) *GetCaCertResponseBody
	GetRequestId() *string
}

type GetCaCertResponseBody struct {
	// The Base64-encoded content of the CA certificate.
	//
	// example:
	//
	// -----BEGIN CERTIFICATE-----\\nMIIFszCCA5ugAwIBAgIDM/1OMA0GCSqGSIb3DQEBCwUAME427zhT4HDLcCEW****-----END CERTIFICATE-----\\n
	CaCert *string `json:"CaCert,omitempty" xml:"CaCert,omitempty"`
	// The request ID.
	//
	// example:
	//
	// E0496204-7586-5B4C-B364-2361CC0ED****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetCaCertResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetCaCertResponseBody) GoString() string {
	return s.String()
}

func (s *GetCaCertResponseBody) GetCaCert() *string {
	return s.CaCert
}

func (s *GetCaCertResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetCaCertResponseBody) SetCaCert(v string) *GetCaCertResponseBody {
	s.CaCert = &v
	return s
}

func (s *GetCaCertResponseBody) SetRequestId(v string) *GetCaCertResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCaCertResponseBody) Validate() error {
	return dara.Validate(s)
}
