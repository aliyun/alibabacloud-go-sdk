// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTransitRouterPolicyTableResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateTransitRouterPolicyTableResponseBody
	GetRequestId() *string
}

type UpdateTransitRouterPolicyTableResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// F7DDDC17-FA06-4AC2-8F35-59D2470FCFC1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateTransitRouterPolicyTableResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateTransitRouterPolicyTableResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateTransitRouterPolicyTableResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateTransitRouterPolicyTableResponseBody) SetRequestId(v string) *UpdateTransitRouterPolicyTableResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateTransitRouterPolicyTableResponseBody) Validate() error {
	return dara.Validate(s)
}
