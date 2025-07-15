// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVpnConnectionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteVpnConnectionResponseBody
	GetRequestId() *string
}

type DeleteVpnConnectionResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 606998F0-B94D-48FE-8316-ACA81BB230DA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteVpnConnectionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteVpnConnectionResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteVpnConnectionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteVpnConnectionResponseBody) SetRequestId(v string) *DeleteVpnConnectionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteVpnConnectionResponseBody) Validate() error {
	return dara.Validate(s)
}
