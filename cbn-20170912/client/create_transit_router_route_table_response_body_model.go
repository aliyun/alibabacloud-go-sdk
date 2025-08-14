// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTransitRouterRouteTableResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateTransitRouterRouteTableResponseBody
	GetRequestId() *string
	SetTransitRouterRouteTableId(v string) *CreateTransitRouterRouteTableResponseBody
	GetTransitRouterRouteTableId() *string
}

type CreateTransitRouterRouteTableResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 599904C8-A3DA-4E5F-83B6-D5364E664247
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the custom route table.
	//
	// example:
	//
	// vtb-bp1xbcgpgcz9axl9m****
	TransitRouterRouteTableId *string `json:"TransitRouterRouteTableId,omitempty" xml:"TransitRouterRouteTableId,omitempty"`
}

func (s CreateTransitRouterRouteTableResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateTransitRouterRouteTableResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTransitRouterRouteTableResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateTransitRouterRouteTableResponseBody) GetTransitRouterRouteTableId() *string {
	return s.TransitRouterRouteTableId
}

func (s *CreateTransitRouterRouteTableResponseBody) SetRequestId(v string) *CreateTransitRouterRouteTableResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateTransitRouterRouteTableResponseBody) SetTransitRouterRouteTableId(v string) *CreateTransitRouterRouteTableResponseBody {
	s.TransitRouterRouteTableId = &v
	return s
}

func (s *CreateTransitRouterRouteTableResponseBody) Validate() error {
	return dara.Validate(s)
}
