// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRegisterTransitRouterMulticastGroupMembersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RegisterTransitRouterMulticastGroupMembersResponseBody
	GetRequestId() *string
}

type RegisterTransitRouterMulticastGroupMembersResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// EB985B7E-2CF8-5EC9-A7DB-F7C82ABD3ACE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RegisterTransitRouterMulticastGroupMembersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RegisterTransitRouterMulticastGroupMembersResponseBody) GoString() string {
	return s.String()
}

func (s *RegisterTransitRouterMulticastGroupMembersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RegisterTransitRouterMulticastGroupMembersResponseBody) SetRequestId(v string) *RegisterTransitRouterMulticastGroupMembersResponseBody {
	s.RequestId = &v
	return s
}

func (s *RegisterTransitRouterMulticastGroupMembersResponseBody) Validate() error {
	return dara.Validate(s)
}
