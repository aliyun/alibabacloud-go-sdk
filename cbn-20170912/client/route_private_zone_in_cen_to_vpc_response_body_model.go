// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRoutePrivateZoneInCenToVpcResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RoutePrivateZoneInCenToVpcResponseBody
	GetRequestId() *string
}

type RoutePrivateZoneInCenToVpcResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// C0245BEF-52AC-44A8-A776-EF96FD26A5CA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RoutePrivateZoneInCenToVpcResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RoutePrivateZoneInCenToVpcResponseBody) GoString() string {
	return s.String()
}

func (s *RoutePrivateZoneInCenToVpcResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RoutePrivateZoneInCenToVpcResponseBody) SetRequestId(v string) *RoutePrivateZoneInCenToVpcResponseBody {
	s.RequestId = &v
	return s
}

func (s *RoutePrivateZoneInCenToVpcResponseBody) Validate() error {
	return dara.Validate(s)
}
