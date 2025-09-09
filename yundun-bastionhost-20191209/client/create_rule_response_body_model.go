// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateRuleResponseBody
	GetRequestId() *string
	SetRuleId(v string) *CreateRuleResponseBody
	GetRuleId() *string
}

type CreateRuleResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// BFA818E3-0A53-51F4-8DB5-AF2A62A6D042
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The authorization rule ID.
	//
	// example:
	//
	// 1
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s CreateRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateRuleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateRuleResponseBody) GetRuleId() *string {
	return s.RuleId
}

func (s *CreateRuleResponseBody) SetRequestId(v string) *CreateRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateRuleResponseBody) SetRuleId(v string) *CreateRuleResponseBody {
	s.RuleId = &v
	return s
}

func (s *CreateRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
