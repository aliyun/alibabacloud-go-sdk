// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnWafRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRuleId(v int64) *DescribeDcdnWafRuleRequest
	GetRuleId() *int64
}

type DescribeDcdnWafRuleRequest struct {
	// The ID of the protection rule. You can specify only one ID in each request.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1000001
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s DescribeDcdnWafRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnWafRuleRequest) GoString() string {
	return s.String()
}

func (s *DescribeDcdnWafRuleRequest) GetRuleId() *int64 {
	return s.RuleId
}

func (s *DescribeDcdnWafRuleRequest) SetRuleId(v int64) *DescribeDcdnWafRuleRequest {
	s.RuleId = &v
	return s
}

func (s *DescribeDcdnWafRuleRequest) Validate() error {
	return dara.Validate(s)
}
