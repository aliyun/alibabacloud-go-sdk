// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCustomRoutingEndpointTrafficPoliciesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPolicyIds(v []*string) *UpdateCustomRoutingEndpointTrafficPoliciesResponseBody
	GetPolicyIds() []*string
	SetRequestId(v string) *UpdateCustomRoutingEndpointTrafficPoliciesResponseBody
	GetRequestId() *string
}

type UpdateCustomRoutingEndpointTrafficPoliciesResponseBody struct {
	// The IDs of the traffic destinations.
	PolicyIds []*string `json:"PolicyIds,omitempty" xml:"PolicyIds,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateCustomRoutingEndpointTrafficPoliciesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateCustomRoutingEndpointTrafficPoliciesResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateCustomRoutingEndpointTrafficPoliciesResponseBody) GetPolicyIds() []*string {
	return s.PolicyIds
}

func (s *UpdateCustomRoutingEndpointTrafficPoliciesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateCustomRoutingEndpointTrafficPoliciesResponseBody) SetPolicyIds(v []*string) *UpdateCustomRoutingEndpointTrafficPoliciesResponseBody {
	s.PolicyIds = v
	return s
}

func (s *UpdateCustomRoutingEndpointTrafficPoliciesResponseBody) SetRequestId(v string) *UpdateCustomRoutingEndpointTrafficPoliciesResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateCustomRoutingEndpointTrafficPoliciesResponseBody) Validate() error {
	return dara.Validate(s)
}
