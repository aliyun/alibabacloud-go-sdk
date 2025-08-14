// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeregisterTransitRouterMulticastGroupMembersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeregisterTransitRouterMulticastGroupMembersResponseBody
	GetRequestId() *string
}

type DeregisterTransitRouterMulticastGroupMembersResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 61D0A553-5E4E-53B5-9DA3-01CBA076A286
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeregisterTransitRouterMulticastGroupMembersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeregisterTransitRouterMulticastGroupMembersResponseBody) GoString() string {
	return s.String()
}

func (s *DeregisterTransitRouterMulticastGroupMembersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeregisterTransitRouterMulticastGroupMembersResponseBody) SetRequestId(v string) *DeregisterTransitRouterMulticastGroupMembersResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeregisterTransitRouterMulticastGroupMembersResponseBody) Validate() error {
	return dara.Validate(s)
}
