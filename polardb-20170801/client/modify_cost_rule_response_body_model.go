// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCostRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCostRuleId(v string) *ModifyCostRuleResponseBody
	GetCostRuleId() *string
	SetRequestId(v string) *ModifyCostRuleResponseBody
	GetRequestId() *string
}

type ModifyCostRuleResponseBody struct {
	// example:
	//
	// 924d450014e64e88ac6e8486f8e990**
	CostRuleId *string `json:"CostRuleId,omitempty" xml:"CostRuleId,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 6BD9CDE4-5E7B-4BF3-9BB8-83C73E******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyCostRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyCostRuleResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyCostRuleResponseBody) GetCostRuleId() *string {
	return s.CostRuleId
}

func (s *ModifyCostRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyCostRuleResponseBody) SetCostRuleId(v string) *ModifyCostRuleResponseBody {
	s.CostRuleId = &v
	return s
}

func (s *ModifyCostRuleResponseBody) SetRequestId(v string) *ModifyCostRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyCostRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
