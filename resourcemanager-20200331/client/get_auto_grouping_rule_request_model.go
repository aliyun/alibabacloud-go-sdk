// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAutoGroupingRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRuleId(v string) *GetAutoGroupingRuleRequest
	GetRuleId() *string
}

type GetAutoGroupingRuleRequest struct {
	// The ID of the rule.
	//
	// This parameter is required.
	//
	// example:
	//
	// gr-acfo******hy6a
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s GetAutoGroupingRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAutoGroupingRuleRequest) GoString() string {
	return s.String()
}

func (s *GetAutoGroupingRuleRequest) GetRuleId() *string {
	return s.RuleId
}

func (s *GetAutoGroupingRuleRequest) SetRuleId(v string) *GetAutoGroupingRuleRequest {
	s.RuleId = &v
	return s
}

func (s *GetAutoGroupingRuleRequest) Validate() error {
	return dara.Validate(s)
}
