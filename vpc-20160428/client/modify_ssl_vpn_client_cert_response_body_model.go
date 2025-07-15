// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySslVpnClientCertResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *ModifySslVpnClientCertResponseBody
	GetName() *string
	SetRequestId(v string) *ModifySslVpnClientCertResponseBody
	GetRequestId() *string
	SetSslVpnClientCertId(v string) *ModifySslVpnClientCertResponseBody
	GetSslVpnClientCertId() *string
}

type ModifySslVpnClientCertResponseBody struct {
	// The name of the SSL client certificate.
	//
	// example:
	//
	// cert2
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 606998F0-B94D-48FE-8316-ACA81BB230DA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the SSL client certificate.
	//
	// example:
	//
	// vsc-bp1n8wcf134yl0osr****
	SslVpnClientCertId *string `json:"SslVpnClientCertId,omitempty" xml:"SslVpnClientCertId,omitempty"`
}

func (s ModifySslVpnClientCertResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifySslVpnClientCertResponseBody) GoString() string {
	return s.String()
}

func (s *ModifySslVpnClientCertResponseBody) GetName() *string {
	return s.Name
}

func (s *ModifySslVpnClientCertResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifySslVpnClientCertResponseBody) GetSslVpnClientCertId() *string {
	return s.SslVpnClientCertId
}

func (s *ModifySslVpnClientCertResponseBody) SetName(v string) *ModifySslVpnClientCertResponseBody {
	s.Name = &v
	return s
}

func (s *ModifySslVpnClientCertResponseBody) SetRequestId(v string) *ModifySslVpnClientCertResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifySslVpnClientCertResponseBody) SetSslVpnClientCertId(v string) *ModifySslVpnClientCertResponseBody {
	s.SslVpnClientCertId = &v
	return s
}

func (s *ModifySslVpnClientCertResponseBody) Validate() error {
	return dara.Validate(s)
}
