// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVpnPbrRouteEntryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteVpnPbrRouteEntryResponseBody
	GetRequestId() *string
}

type DeleteVpnPbrRouteEntryResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 5BE01CD7-5A50-472D-AC14-CA181C5C03BE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteVpnPbrRouteEntryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteVpnPbrRouteEntryResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteVpnPbrRouteEntryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteVpnPbrRouteEntryResponseBody) SetRequestId(v string) *DeleteVpnPbrRouteEntryResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteVpnPbrRouteEntryResponseBody) Validate() error {
	return dara.Validate(s)
}
