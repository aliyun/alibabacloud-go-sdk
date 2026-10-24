// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetClientRuleStrategyStatisticsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPolicyType(v int32) *GetClientRuleStrategyStatisticsRequest
	GetPolicyType() *int32
}

type GetClientRuleStrategyStatisticsRequest struct {
	// The policy type. Set this parameter to 2 to query the statistics of proactive defense policies.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2
	PolicyType *int32 `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
}

func (s GetClientRuleStrategyStatisticsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetClientRuleStrategyStatisticsRequest) GoString() string {
	return s.String()
}

func (s *GetClientRuleStrategyStatisticsRequest) GetPolicyType() *int32 {
	return s.PolicyType
}

func (s *GetClientRuleStrategyStatisticsRequest) SetPolicyType(v int32) *GetClientRuleStrategyStatisticsRequest {
	s.PolicyType = &v
	return s
}

func (s *GetClientRuleStrategyStatisticsRequest) Validate() error {
	return dara.Validate(s)
}
