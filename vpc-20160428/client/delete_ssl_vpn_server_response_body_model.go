// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSslVpnServerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteSslVpnServerResponseBody
	GetRequestId() *string
}

type DeleteSslVpnServerResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 606998F0-B94D-48FE-8316-ACA81BB230DA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteSslVpnServerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteSslVpnServerResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSslVpnServerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteSslVpnServerResponseBody) SetRequestId(v string) *DeleteSslVpnServerResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteSslVpnServerResponseBody) Validate() error {
	return dara.Validate(s)
}
