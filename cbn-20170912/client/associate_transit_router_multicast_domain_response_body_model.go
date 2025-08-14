// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssociateTransitRouterMulticastDomainResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AssociateTransitRouterMulticastDomainResponseBody
	GetRequestId() *string
}

type AssociateTransitRouterMulticastDomainResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// F88AC12C-943B-50E9-A344-4F8820BB07A7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AssociateTransitRouterMulticastDomainResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AssociateTransitRouterMulticastDomainResponseBody) GoString() string {
	return s.String()
}

func (s *AssociateTransitRouterMulticastDomainResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AssociateTransitRouterMulticastDomainResponseBody) SetRequestId(v string) *AssociateTransitRouterMulticastDomainResponseBody {
	s.RequestId = &v
	return s
}

func (s *AssociateTransitRouterMulticastDomainResponseBody) Validate() error {
	return dara.Validate(s)
}
