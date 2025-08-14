// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisassociateTransitRouterMulticastDomainResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DisassociateTransitRouterMulticastDomainResponseBody
	GetRequestId() *string
}

type DisassociateTransitRouterMulticastDomainResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 6F6B3FF0-45D1-5416-B189-C45A42A0222B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DisassociateTransitRouterMulticastDomainResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DisassociateTransitRouterMulticastDomainResponseBody) GoString() string {
	return s.String()
}

func (s *DisassociateTransitRouterMulticastDomainResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DisassociateTransitRouterMulticastDomainResponseBody) SetRequestId(v string) *DisassociateTransitRouterMulticastDomainResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisassociateTransitRouterMulticastDomainResponseBody) Validate() error {
	return dara.Validate(s)
}
