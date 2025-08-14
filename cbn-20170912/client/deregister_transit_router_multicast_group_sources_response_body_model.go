// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeregisterTransitRouterMulticastGroupSourcesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeregisterTransitRouterMulticastGroupSourcesResponseBody
	GetRequestId() *string
}

type DeregisterTransitRouterMulticastGroupSourcesResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 53E7E8BE-7F4E-5458-ACCA-9B5C1D6A642D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeregisterTransitRouterMulticastGroupSourcesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeregisterTransitRouterMulticastGroupSourcesResponseBody) GoString() string {
	return s.String()
}

func (s *DeregisterTransitRouterMulticastGroupSourcesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeregisterTransitRouterMulticastGroupSourcesResponseBody) SetRequestId(v string) *DeregisterTransitRouterMulticastGroupSourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeregisterTransitRouterMulticastGroupSourcesResponseBody) Validate() error {
	return dara.Validate(s)
}
