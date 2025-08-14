// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyTransitRouterCidrResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyTransitRouterCidrResponseBody
	GetRequestId() *string
}

type ModifyTransitRouterCidrResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 0876E54E-3E36-5C31-89F0-9EE8A9266F9A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyTransitRouterCidrResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyTransitRouterCidrResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyTransitRouterCidrResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyTransitRouterCidrResponseBody) SetRequestId(v string) *ModifyTransitRouterCidrResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyTransitRouterCidrResponseBody) Validate() error {
	return dara.Validate(s)
}
