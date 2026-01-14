// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateForwardingRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetForwardingRules(v []*CreateForwardingRulesResponseBodyForwardingRules) *CreateForwardingRulesResponseBody
	GetForwardingRules() []*CreateForwardingRulesResponseBodyForwardingRules
	SetRequestId(v string) *CreateForwardingRulesResponseBody
	GetRequestId() *string
}

type CreateForwardingRulesResponseBody struct {
	// The forwarding rules.
	ForwardingRules []*CreateForwardingRulesResponseBodyForwardingRules `json:"ForwardingRules,omitempty" xml:"ForwardingRules,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 64ADAB1E-0B7F-4FD8-A404-3BECC0E9CCFF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateForwardingRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateForwardingRulesResponseBody) GoString() string {
	return s.String()
}

func (s *CreateForwardingRulesResponseBody) GetForwardingRules() []*CreateForwardingRulesResponseBodyForwardingRules {
	return s.ForwardingRules
}

func (s *CreateForwardingRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateForwardingRulesResponseBody) SetForwardingRules(v []*CreateForwardingRulesResponseBodyForwardingRules) *CreateForwardingRulesResponseBody {
	s.ForwardingRules = v
	return s
}

func (s *CreateForwardingRulesResponseBody) SetRequestId(v string) *CreateForwardingRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateForwardingRulesResponseBody) Validate() error {
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

type CreateForwardingRulesResponseBodyForwardingRules struct {
	// The ID of the forwarding rule.
	//
	// example:
	//
	// frule-bp1dii16gu9qdvb34****
	ForwardingRuleId *string `json:"ForwardingRuleId,omitempty" xml:"ForwardingRuleId,omitempty"`
}

func (s CreateForwardingRulesResponseBodyForwardingRules) String() string {
	return dara.Prettify(s)
}

func (s CreateForwardingRulesResponseBodyForwardingRules) GoString() string {
	return s.String()
}

func (s *CreateForwardingRulesResponseBodyForwardingRules) GetForwardingRuleId() *string {
	return s.ForwardingRuleId
}

func (s *CreateForwardingRulesResponseBodyForwardingRules) SetForwardingRuleId(v string) *CreateForwardingRulesResponseBodyForwardingRules {
	s.ForwardingRuleId = &v
	return s
}

func (s *CreateForwardingRulesResponseBodyForwardingRules) Validate() error {
	return dara.Validate(s)
}
