// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSslVpnClientCertResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteSslVpnClientCertResponseBody
	GetRequestId() *string
}

type DeleteSslVpnClientCertResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 606998F0-B94D-48FE-8316-ACA81BB230DA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteSslVpnClientCertResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteSslVpnClientCertResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSslVpnClientCertResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteSslVpnClientCertResponseBody) SetRequestId(v string) *DeleteSslVpnClientCertResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteSslVpnClientCertResponseBody) Validate() error {
	return dara.Validate(s)
}
