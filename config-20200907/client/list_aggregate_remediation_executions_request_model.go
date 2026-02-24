// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAggregateRemediationExecutionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAggregatorId(v string) *ListAggregateRemediationExecutionsRequest
	GetAggregatorId() *string
	SetConfigRuleId(v string) *ListAggregateRemediationExecutionsRequest
	GetConfigRuleId() *string
	SetExecutionStatus(v string) *ListAggregateRemediationExecutionsRequest
	GetExecutionStatus() *string
	SetMaxResults(v int64) *ListAggregateRemediationExecutionsRequest
	GetMaxResults() *int64
	SetNextToken(v string) *ListAggregateRemediationExecutionsRequest
	GetNextToken() *string
	SetResourceAccountId(v int64) *ListAggregateRemediationExecutionsRequest
	GetResourceAccountId() *int64
}

type ListAggregateRemediationExecutionsRequest struct {
	// This parameter is required.
	AggregatorId *string `json:"AggregatorId,omitempty" xml:"AggregatorId,omitempty"`
	// This parameter is required.
	ConfigRuleId      *string `json:"ConfigRuleId,omitempty" xml:"ConfigRuleId,omitempty"`
	ExecutionStatus   *string `json:"ExecutionStatus,omitempty" xml:"ExecutionStatus,omitempty"`
	MaxResults        *int64  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken         *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	ResourceAccountId *int64  `json:"ResourceAccountId,omitempty" xml:"ResourceAccountId,omitempty"`
}

func (s ListAggregateRemediationExecutionsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAggregateRemediationExecutionsRequest) GoString() string {
	return s.String()
}

func (s *ListAggregateRemediationExecutionsRequest) GetAggregatorId() *string {
	return s.AggregatorId
}

func (s *ListAggregateRemediationExecutionsRequest) GetConfigRuleId() *string {
	return s.ConfigRuleId
}

func (s *ListAggregateRemediationExecutionsRequest) GetExecutionStatus() *string {
	return s.ExecutionStatus
}

func (s *ListAggregateRemediationExecutionsRequest) GetMaxResults() *int64 {
	return s.MaxResults
}

func (s *ListAggregateRemediationExecutionsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListAggregateRemediationExecutionsRequest) GetResourceAccountId() *int64 {
	return s.ResourceAccountId
}

func (s *ListAggregateRemediationExecutionsRequest) SetAggregatorId(v string) *ListAggregateRemediationExecutionsRequest {
	s.AggregatorId = &v
	return s
}

func (s *ListAggregateRemediationExecutionsRequest) SetConfigRuleId(v string) *ListAggregateRemediationExecutionsRequest {
	s.ConfigRuleId = &v
	return s
}

func (s *ListAggregateRemediationExecutionsRequest) SetExecutionStatus(v string) *ListAggregateRemediationExecutionsRequest {
	s.ExecutionStatus = &v
	return s
}

func (s *ListAggregateRemediationExecutionsRequest) SetMaxResults(v int64) *ListAggregateRemediationExecutionsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListAggregateRemediationExecutionsRequest) SetNextToken(v string) *ListAggregateRemediationExecutionsRequest {
	s.NextToken = &v
	return s
}

func (s *ListAggregateRemediationExecutionsRequest) SetResourceAccountId(v int64) *ListAggregateRemediationExecutionsRequest {
	s.ResourceAccountId = &v
	return s
}

func (s *ListAggregateRemediationExecutionsRequest) Validate() error {
	return dara.Validate(s)
}
