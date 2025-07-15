// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVpnRouteEntryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteVpnRouteEntryResponseBody
	GetRequestId() *string
}

type DeleteVpnRouteEntryResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 5BE01CD7-5A50-472D-AC14-CA181C5C03BE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteVpnRouteEntryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteVpnRouteEntryResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteVpnRouteEntryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteVpnRouteEntryResponseBody) SetRequestId(v string) *DeleteVpnRouteEntryResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteVpnRouteEntryResponseBody) Validate() error {
	return dara.Validate(s)
}
