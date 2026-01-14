// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCustomRoutingEndpointTrafficPoliciesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPolicyIds(v []*string) *CreateCustomRoutingEndpointTrafficPoliciesResponseBody
	GetPolicyIds() []*string
	SetRequestId(v string) *CreateCustomRoutingEndpointTrafficPoliciesResponseBody
	GetRequestId() *string
}

type CreateCustomRoutingEndpointTrafficPoliciesResponseBody struct {
	// The IDs of the traffic destinations.
	PolicyIds []*string `json:"PolicyIds,omitempty" xml:"PolicyIds,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateCustomRoutingEndpointTrafficPoliciesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateCustomRoutingEndpointTrafficPoliciesResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCustomRoutingEndpointTrafficPoliciesResponseBody) GetPolicyIds() []*string {
	return s.PolicyIds
}

func (s *CreateCustomRoutingEndpointTrafficPoliciesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateCustomRoutingEndpointTrafficPoliciesResponseBody) SetPolicyIds(v []*string) *CreateCustomRoutingEndpointTrafficPoliciesResponseBody {
	s.PolicyIds = v
	return s
}

func (s *CreateCustomRoutingEndpointTrafficPoliciesResponseBody) SetRequestId(v string) *CreateCustomRoutingEndpointTrafficPoliciesResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateCustomRoutingEndpointTrafficPoliciesResponseBody) Validate() error {
	return dara.Validate(s)
}
