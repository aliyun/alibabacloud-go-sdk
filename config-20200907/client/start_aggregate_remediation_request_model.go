// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartAggregateRemediationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAggregatorId(v string) *StartAggregateRemediationRequest
	GetAggregatorId() *string
	SetConfigRuleId(v string) *StartAggregateRemediationRequest
	GetConfigRuleId() *string
	SetResourceAccountId(v int64) *StartAggregateRemediationRequest
	GetResourceAccountId() *int64
}

type StartAggregateRemediationRequest struct {
	// This parameter is required.
	AggregatorId *string `json:"AggregatorId,omitempty" xml:"AggregatorId,omitempty"`
	// This parameter is required.
	ConfigRuleId      *string `json:"ConfigRuleId,omitempty" xml:"ConfigRuleId,omitempty"`
	ResourceAccountId *int64  `json:"ResourceAccountId,omitempty" xml:"ResourceAccountId,omitempty"`
}

func (s StartAggregateRemediationRequest) String() string {
	return dara.Prettify(s)
}

func (s StartAggregateRemediationRequest) GoString() string {
	return s.String()
}

func (s *StartAggregateRemediationRequest) GetAggregatorId() *string {
	return s.AggregatorId
}

func (s *StartAggregateRemediationRequest) GetConfigRuleId() *string {
	return s.ConfigRuleId
}

func (s *StartAggregateRemediationRequest) GetResourceAccountId() *int64 {
	return s.ResourceAccountId
}

func (s *StartAggregateRemediationRequest) SetAggregatorId(v string) *StartAggregateRemediationRequest {
	s.AggregatorId = &v
	return s
}

func (s *StartAggregateRemediationRequest) SetConfigRuleId(v string) *StartAggregateRemediationRequest {
	s.ConfigRuleId = &v
	return s
}

func (s *StartAggregateRemediationRequest) SetResourceAccountId(v int64) *StartAggregateRemediationRequest {
	s.ResourceAccountId = &v
	return s
}

func (s *StartAggregateRemediationRequest) Validate() error {
	return dara.Validate(s)
}
