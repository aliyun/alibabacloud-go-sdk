// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCustomRoutingEndpointTrafficPoliciesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteCustomRoutingEndpointTrafficPoliciesResponseBody
	GetRequestId() *string
}

type DeleteCustomRoutingEndpointTrafficPoliciesResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteCustomRoutingEndpointTrafficPoliciesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteCustomRoutingEndpointTrafficPoliciesResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCustomRoutingEndpointTrafficPoliciesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteCustomRoutingEndpointTrafficPoliciesResponseBody) SetRequestId(v string) *DeleteCustomRoutingEndpointTrafficPoliciesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteCustomRoutingEndpointTrafficPoliciesResponseBody) Validate() error {
	return dara.Validate(s)
}
