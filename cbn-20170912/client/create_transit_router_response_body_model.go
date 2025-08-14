// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTransitRouterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateTransitRouterResponseBody
	GetRequestId() *string
	SetTransitRouterId(v string) *CreateTransitRouterResponseBody
	GetTransitRouterId() *string
}

type CreateTransitRouterResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 404DA7EC-F495-44B5-B543-6EDCDF90F3D1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the Enterprise Edition transit router.
	//
	// example:
	//
	// tr-uf6llz2286805i44g****
	TransitRouterId *string `json:"TransitRouterId,omitempty" xml:"TransitRouterId,omitempty"`
}

func (s CreateTransitRouterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateTransitRouterResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTransitRouterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateTransitRouterResponseBody) GetTransitRouterId() *string {
	return s.TransitRouterId
}

func (s *CreateTransitRouterResponseBody) SetRequestId(v string) *CreateTransitRouterResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateTransitRouterResponseBody) SetTransitRouterId(v string) *CreateTransitRouterResponseBody {
	s.TransitRouterId = &v
	return s
}

func (s *CreateTransitRouterResponseBody) Validate() error {
	return dara.Validate(s)
}
