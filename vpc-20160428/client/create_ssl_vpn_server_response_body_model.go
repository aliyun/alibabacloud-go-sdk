// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSslVpnServerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *CreateSslVpnServerResponseBody
	GetName() *string
	SetRequestId(v string) *CreateSslVpnServerResponseBody
	GetRequestId() *string
	SetSslVpnServerId(v string) *CreateSslVpnServerResponseBody
	GetSslVpnServerId() *string
}

type CreateSslVpnServerResponseBody struct {
	// The SSL server name.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The request ID.
	//
	// example:
	//
	// E98A9651-7098-40C7-8F85-C818D1EBBA85
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the SSL server.
	//
	// example:
	//
	// vss-bp18q7hzj6largv4v****
	SslVpnServerId *string `json:"SslVpnServerId,omitempty" xml:"SslVpnServerId,omitempty"`
}

func (s CreateSslVpnServerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateSslVpnServerResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSslVpnServerResponseBody) GetName() *string {
	return s.Name
}

func (s *CreateSslVpnServerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateSslVpnServerResponseBody) GetSslVpnServerId() *string {
	return s.SslVpnServerId
}

func (s *CreateSslVpnServerResponseBody) SetName(v string) *CreateSslVpnServerResponseBody {
	s.Name = &v
	return s
}

func (s *CreateSslVpnServerResponseBody) SetRequestId(v string) *CreateSslVpnServerResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSslVpnServerResponseBody) SetSslVpnServerId(v string) *CreateSslVpnServerResponseBody {
	s.SslVpnServerId = &v
	return s
}

func (s *CreateSslVpnServerResponseBody) Validate() error {
	return dara.Validate(s)
}
