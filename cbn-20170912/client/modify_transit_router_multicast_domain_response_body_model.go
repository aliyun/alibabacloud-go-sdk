// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyTransitRouterMulticastDomainResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyTransitRouterMulticastDomainResponseBody
	GetRequestId() *string
}

type ModifyTransitRouterMulticastDomainResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 26273D23-5CB0-5EFC-AF5F-78A5448084C9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyTransitRouterMulticastDomainResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyTransitRouterMulticastDomainResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyTransitRouterMulticastDomainResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyTransitRouterMulticastDomainResponseBody) SetRequestId(v string) *ModifyTransitRouterMulticastDomainResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyTransitRouterMulticastDomainResponseBody) Validate() error {
	return dara.Validate(s)
}
