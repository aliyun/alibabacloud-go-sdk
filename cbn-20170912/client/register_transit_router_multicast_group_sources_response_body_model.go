// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRegisterTransitRouterMulticastGroupSourcesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RegisterTransitRouterMulticastGroupSourcesResponseBody
	GetRequestId() *string
}

type RegisterTransitRouterMulticastGroupSourcesResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 9C5D5D70-0AFF-5E5C-8D8A-E92C90C8FB08
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RegisterTransitRouterMulticastGroupSourcesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RegisterTransitRouterMulticastGroupSourcesResponseBody) GoString() string {
	return s.String()
}

func (s *RegisterTransitRouterMulticastGroupSourcesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RegisterTransitRouterMulticastGroupSourcesResponseBody) SetRequestId(v string) *RegisterTransitRouterMulticastGroupSourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *RegisterTransitRouterMulticastGroupSourcesResponseBody) Validate() error {
	return dara.Validate(s)
}
