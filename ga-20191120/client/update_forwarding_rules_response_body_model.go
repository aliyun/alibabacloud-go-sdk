// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateForwardingRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetForwardingRules(v []*UpdateForwardingRulesResponseBodyForwardingRules) *UpdateForwardingRulesResponseBody
	GetForwardingRules() []*UpdateForwardingRulesResponseBodyForwardingRules
	SetRequestId(v string) *UpdateForwardingRulesResponseBody
	GetRequestId() *string
}

type UpdateForwardingRulesResponseBody struct {
	// Details about the forwarding rules.
	ForwardingRules []*UpdateForwardingRulesResponseBodyForwardingRules `json:"ForwardingRules,omitempty" xml:"ForwardingRules,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 64ADAB1E-0B7F-4FD8-A404-3BECC0E9CCFF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateForwardingRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateForwardingRulesResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateForwardingRulesResponseBody) GetForwardingRules() []*UpdateForwardingRulesResponseBodyForwardingRules {
	return s.ForwardingRules
}

func (s *UpdateForwardingRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateForwardingRulesResponseBody) SetForwardingRules(v []*UpdateForwardingRulesResponseBodyForwardingRules) *UpdateForwardingRulesResponseBody {
	s.ForwardingRules = v
	return s
}

func (s *UpdateForwardingRulesResponseBody) SetRequestId(v string) *UpdateForwardingRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateForwardingRulesResponseBody) Validate() error {
	if s.ForwardingRules != nil {
		for _, item := range s.ForwardingRules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateForwardingRulesResponseBodyForwardingRules struct {
	// The forwarding rule ID.
	//
	// example:
	//
	// frule-bp1dii16gu9qdvb34****
	ForwardingRuleId *string `json:"ForwardingRuleId,omitempty" xml:"ForwardingRuleId,omitempty"`
}

func (s UpdateForwardingRulesResponseBodyForwardingRules) String() string {
	return dara.Prettify(s)
}

func (s UpdateForwardingRulesResponseBodyForwardingRules) GoString() string {
	return s.String()
}

func (s *UpdateForwardingRulesResponseBodyForwardingRules) GetForwardingRuleId() *string {
	return s.ForwardingRuleId
}

func (s *UpdateForwardingRulesResponseBodyForwardingRules) SetForwardingRuleId(v string) *UpdateForwardingRulesResponseBodyForwardingRules {
	s.ForwardingRuleId = &v
	return s
}

func (s *UpdateForwardingRulesResponseBodyForwardingRules) Validate() error {
	return dara.Validate(s)
}
