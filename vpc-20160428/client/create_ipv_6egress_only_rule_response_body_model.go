// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateIpv6EgressOnlyRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetIpv6EgressRuleId(v string) *CreateIpv6EgressOnlyRuleResponseBody
	GetIpv6EgressRuleId() *string
	SetRequestId(v string) *CreateIpv6EgressOnlyRuleResponseBody
	GetRequestId() *string
}

type CreateIpv6EgressOnlyRuleResponseBody struct {
	// The ID of the egress-only rule.
	//
	// example:
	//
	// ipv6py-hp3w98rmlbqp01245****
	Ipv6EgressRuleId *string `json:"Ipv6EgressRuleId,omitempty" xml:"Ipv6EgressRuleId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 9DFEDBEE-E5AB-49E8-A2DC-CC114C67AF75
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateIpv6EgressOnlyRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateIpv6EgressOnlyRuleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateIpv6EgressOnlyRuleResponseBody) GetIpv6EgressRuleId() *string {
	return s.Ipv6EgressRuleId
}

func (s *CreateIpv6EgressOnlyRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateIpv6EgressOnlyRuleResponseBody) SetIpv6EgressRuleId(v string) *CreateIpv6EgressOnlyRuleResponseBody {
	s.Ipv6EgressRuleId = &v
	return s
}

func (s *CreateIpv6EgressOnlyRuleResponseBody) SetRequestId(v string) *CreateIpv6EgressOnlyRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateIpv6EgressOnlyRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
