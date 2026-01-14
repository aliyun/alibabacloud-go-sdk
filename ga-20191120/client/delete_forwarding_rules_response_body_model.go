// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteForwardingRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetForwardingRules(v []*DeleteForwardingRulesResponseBodyForwardingRules) *DeleteForwardingRulesResponseBody
	GetForwardingRules() []*DeleteForwardingRulesResponseBodyForwardingRules
	SetRequestId(v string) *DeleteForwardingRulesResponseBody
	GetRequestId() *string
}

type DeleteForwardingRulesResponseBody struct {
	// The forwarding rules.
	ForwardingRules []*DeleteForwardingRulesResponseBodyForwardingRules `json:"ForwardingRules,omitempty" xml:"ForwardingRules,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// CFC67ED9-4AB1-431F-B6E3-A752B7B8CCD4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteForwardingRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteForwardingRulesResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteForwardingRulesResponseBody) GetForwardingRules() []*DeleteForwardingRulesResponseBodyForwardingRules {
	return s.ForwardingRules
}

func (s *DeleteForwardingRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteForwardingRulesResponseBody) SetForwardingRules(v []*DeleteForwardingRulesResponseBodyForwardingRules) *DeleteForwardingRulesResponseBody {
	s.ForwardingRules = v
	return s
}

func (s *DeleteForwardingRulesResponseBody) SetRequestId(v string) *DeleteForwardingRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteForwardingRulesResponseBody) Validate() error {
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

type DeleteForwardingRulesResponseBodyForwardingRules struct {
	// The forwarding rule ID.
	//
	// example:
	//
	// frule-bp19a3t3yzr21q3****
	ForwardingRuleId *string `json:"ForwardingRuleId,omitempty" xml:"ForwardingRuleId,omitempty"`
}

func (s DeleteForwardingRulesResponseBodyForwardingRules) String() string {
	return dara.Prettify(s)
}

func (s DeleteForwardingRulesResponseBodyForwardingRules) GoString() string {
	return s.String()
}

func (s *DeleteForwardingRulesResponseBodyForwardingRules) GetForwardingRuleId() *string {
	return s.ForwardingRuleId
}

func (s *DeleteForwardingRulesResponseBodyForwardingRules) SetForwardingRuleId(v string) *DeleteForwardingRulesResponseBodyForwardingRules {
	s.ForwardingRuleId = &v
	return s
}

func (s *DeleteForwardingRulesResponseBodyForwardingRules) Validate() error {
	return dara.Validate(s)
}
