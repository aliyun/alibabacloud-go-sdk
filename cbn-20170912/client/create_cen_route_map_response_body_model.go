// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCenRouteMapResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateCenRouteMapResponseBody
	GetRequestId() *string
	SetRouteMapId(v string) *CreateCenRouteMapResponseBody
	GetRouteMapId() *string
}

type CreateCenRouteMapResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 62172DD5-6BAC-45DF-8D44-56SDF467BAC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the routing policy.
	//
	// example:
	//
	// cenrmap-w4yf7toozfol3q****
	RouteMapId *string `json:"RouteMapId,omitempty" xml:"RouteMapId,omitempty"`
}

func (s CreateCenRouteMapResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateCenRouteMapResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCenRouteMapResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateCenRouteMapResponseBody) GetRouteMapId() *string {
	return s.RouteMapId
}

func (s *CreateCenRouteMapResponseBody) SetRequestId(v string) *CreateCenRouteMapResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateCenRouteMapResponseBody) SetRouteMapId(v string) *CreateCenRouteMapResponseBody {
	s.RouteMapId = &v
	return s
}

func (s *CreateCenRouteMapResponseBody) Validate() error {
	return dara.Validate(s)
}
