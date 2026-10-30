// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRouteMapResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateRouteMapResponseBody
	GetRequestId() *string
	SetRouteMapId(v string) *CreateRouteMapResponseBody
	GetRouteMapId() *string
}

type CreateRouteMapResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 0C2EE7A8-74D4-4081-8236-CEBDE3BBCF50
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the route map.
	//
	// example:
	//
	// rm-abcdedfghij****
	RouteMapId *string `json:"RouteMapId,omitempty" xml:"RouteMapId,omitempty"`
}

func (s CreateRouteMapResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateRouteMapResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRouteMapResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateRouteMapResponseBody) GetRouteMapId() *string {
	return s.RouteMapId
}

func (s *CreateRouteMapResponseBody) SetRequestId(v string) *CreateRouteMapResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateRouteMapResponseBody) SetRouteMapId(v string) *CreateRouteMapResponseBody {
	s.RouteMapId = &v
	return s
}

func (s *CreateRouteMapResponseBody) Validate() error {
	return dara.Validate(s)
}
