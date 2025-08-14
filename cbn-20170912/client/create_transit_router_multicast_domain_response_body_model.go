// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTransitRouterMulticastDomainResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateTransitRouterMulticastDomainResponseBody
	GetRequestId() *string
	SetTransitRouterMulticastDomainId(v string) *CreateTransitRouterMulticastDomainResponseBody
	GetTransitRouterMulticastDomainId() *string
}

type CreateTransitRouterMulticastDomainResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 94E19C6F-206F-5223-9A63-64B85851BC04
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the multicast domain.
	//
	// example:
	//
	// tr-mcast-domain-40cwj0rgzgdtam****
	TransitRouterMulticastDomainId *string `json:"TransitRouterMulticastDomainId,omitempty" xml:"TransitRouterMulticastDomainId,omitempty"`
}

func (s CreateTransitRouterMulticastDomainResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateTransitRouterMulticastDomainResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTransitRouterMulticastDomainResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateTransitRouterMulticastDomainResponseBody) GetTransitRouterMulticastDomainId() *string {
	return s.TransitRouterMulticastDomainId
}

func (s *CreateTransitRouterMulticastDomainResponseBody) SetRequestId(v string) *CreateTransitRouterMulticastDomainResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateTransitRouterMulticastDomainResponseBody) SetTransitRouterMulticastDomainId(v string) *CreateTransitRouterMulticastDomainResponseBody {
	s.TransitRouterMulticastDomainId = &v
	return s
}

func (s *CreateTransitRouterMulticastDomainResponseBody) Validate() error {
	return dara.Validate(s)
}
