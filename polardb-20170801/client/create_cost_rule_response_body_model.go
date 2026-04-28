// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCostRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCostRuleId(v string) *CreateCostRuleResponseBody
	GetCostRuleId() *string
	SetRequestId(v string) *CreateCostRuleResponseBody
	GetRequestId() *string
}

type CreateCostRuleResponseBody struct {
	// example:
	//
	// 924d450014e64e88ac6e8486f8e990**
	CostRuleId *string `json:"CostRuleId,omitempty" xml:"CostRuleId,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 580EF224-9647-59E7-9950-D9EBFD6A2921
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateCostRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateCostRuleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCostRuleResponseBody) GetCostRuleId() *string {
	return s.CostRuleId
}

func (s *CreateCostRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateCostRuleResponseBody) SetCostRuleId(v string) *CreateCostRuleResponseBody {
	s.CostRuleId = &v
	return s
}

func (s *CreateCostRuleResponseBody) SetRequestId(v string) *CreateCostRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateCostRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
