// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRouteTableResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateRouteTableResponseBody
	GetRequestId() *string
	SetRouteTableId(v string) *CreateRouteTableResponseBody
	GetRouteTableId() *string
}

type CreateRouteTableResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 62172DD5-6BAC-45DF-8D44-E11ED28DCAC0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The IDs of the route tables.
	//
	// example:
	//
	// vtb-bp145q7glnuzdvzu2****
	RouteTableId *string `json:"RouteTableId,omitempty" xml:"RouteTableId,omitempty"`
}

func (s CreateRouteTableResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateRouteTableResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRouteTableResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateRouteTableResponseBody) GetRouteTableId() *string {
	return s.RouteTableId
}

func (s *CreateRouteTableResponseBody) SetRequestId(v string) *CreateRouteTableResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateRouteTableResponseBody) SetRouteTableId(v string) *CreateRouteTableResponseBody {
	s.RouteTableId = &v
	return s
}

func (s *CreateRouteTableResponseBody) Validate() error {
	return dara.Validate(s)
}
