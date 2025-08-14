// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTransitRouterMulticastDomainResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteTransitRouterMulticastDomainResponseBody
	GetRequestId() *string
}

type DeleteTransitRouterMulticastDomainResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 40194E53-2484-5831-BB53-E11D123C1A32
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteTransitRouterMulticastDomainResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteTransitRouterMulticastDomainResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteTransitRouterMulticastDomainResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteTransitRouterMulticastDomainResponseBody) SetRequestId(v string) *DeleteTransitRouterMulticastDomainResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteTransitRouterMulticastDomainResponseBody) Validate() error {
	return dara.Validate(s)
}
