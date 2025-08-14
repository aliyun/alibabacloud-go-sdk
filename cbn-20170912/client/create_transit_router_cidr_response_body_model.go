// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTransitRouterCidrResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateTransitRouterCidrResponseBody
	GetRequestId() *string
	SetTransitRouterCidrId(v string) *CreateTransitRouterCidrResponseBody
	GetTransitRouterCidrId() *string
}

type CreateTransitRouterCidrResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 0876E54E-3E36-5C31-89F0-9EE8A9266F9A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the CIDR block.
	//
	// example:
	//
	// cidr-0zv0q9crqpntzz****
	TransitRouterCidrId *string `json:"TransitRouterCidrId,omitempty" xml:"TransitRouterCidrId,omitempty"`
}

func (s CreateTransitRouterCidrResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateTransitRouterCidrResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTransitRouterCidrResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateTransitRouterCidrResponseBody) GetTransitRouterCidrId() *string {
	return s.TransitRouterCidrId
}

func (s *CreateTransitRouterCidrResponseBody) SetRequestId(v string) *CreateTransitRouterCidrResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateTransitRouterCidrResponseBody) SetTransitRouterCidrId(v string) *CreateTransitRouterCidrResponseBody {
	s.TransitRouterCidrId = &v
	return s
}

func (s *CreateTransitRouterCidrResponseBody) Validate() error {
	return dara.Validate(s)
}
