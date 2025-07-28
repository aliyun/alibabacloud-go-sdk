// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddResolverRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AddResolverRuleResponseBody
	GetRequestId() *string
	SetRuleId(v string) *AddResolverRuleResponseBody
	GetRuleId() *string
}

type AddResolverRuleResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 725B8BED-901F-480C-BBAC-FA59A18580C1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the forwarding rule.
	//
	// example:
	//
	// hr****
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s AddResolverRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddResolverRuleResponseBody) GoString() string {
	return s.String()
}

func (s *AddResolverRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddResolverRuleResponseBody) GetRuleId() *string {
	return s.RuleId
}

func (s *AddResolverRuleResponseBody) SetRequestId(v string) *AddResolverRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddResolverRuleResponseBody) SetRuleId(v string) *AddResolverRuleResponseBody {
	s.RuleId = &v
	return s
}

func (s *AddResolverRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
