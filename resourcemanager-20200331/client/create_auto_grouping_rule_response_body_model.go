// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAutoGroupingRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateAutoGroupingRuleResponseBody
	GetRequestId() *string
	SetRuleId(v string) *CreateAutoGroupingRuleResponseBody
	GetRuleId() *string
}

type CreateAutoGroupingRuleResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// F7701451-340B-5CB3-AEA7-7D831F7F38C0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the rule.
	//
	// example:
	//
	// gr-acfo******hy6a
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s CreateAutoGroupingRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateAutoGroupingRuleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAutoGroupingRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateAutoGroupingRuleResponseBody) GetRuleId() *string {
	return s.RuleId
}

func (s *CreateAutoGroupingRuleResponseBody) SetRequestId(v string) *CreateAutoGroupingRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAutoGroupingRuleResponseBody) SetRuleId(v string) *CreateAutoGroupingRuleResponseBody {
	s.RuleId = &v
	return s
}

func (s *CreateAutoGroupingRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
