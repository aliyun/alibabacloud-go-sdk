// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSslVpnClientCertResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *CreateSslVpnClientCertResponseBody
	GetName() *string
	SetRequestId(v string) *CreateSslVpnClientCertResponseBody
	GetRequestId() *string
	SetSslVpnClientCertId(v string) *CreateSslVpnClientCertResponseBody
	GetSslVpnClientCertId() *string
}

type CreateSslVpnClientCertResponseBody struct {
	// The name of the SSL client certificate.
	//
	// example:
	//
	// SslVpnClientCert
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 079874CD-AEC1-43E6-AC03-ADD96B6E4907
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the SSL client certificate.
	//
	// example:
	//
	// vsc-m5euof6s5jy8vs5kd****
	SslVpnClientCertId *string `json:"SslVpnClientCertId,omitempty" xml:"SslVpnClientCertId,omitempty"`
}

func (s CreateSslVpnClientCertResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateSslVpnClientCertResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSslVpnClientCertResponseBody) GetName() *string {
	return s.Name
}

func (s *CreateSslVpnClientCertResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateSslVpnClientCertResponseBody) GetSslVpnClientCertId() *string {
	return s.SslVpnClientCertId
}

func (s *CreateSslVpnClientCertResponseBody) SetName(v string) *CreateSslVpnClientCertResponseBody {
	s.Name = &v
	return s
}

func (s *CreateSslVpnClientCertResponseBody) SetRequestId(v string) *CreateSslVpnClientCertResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSslVpnClientCertResponseBody) SetSslVpnClientCertId(v string) *CreateSslVpnClientCertResponseBody {
	s.SslVpnClientCertId = &v
	return s
}

func (s *CreateSslVpnClientCertResponseBody) Validate() error {
	return dara.Validate(s)
}
