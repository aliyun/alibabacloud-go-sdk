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
	// The ID of the account group.
	//
	// For more information about how to obtain the ID of an account group, see [ListAggregators](https://help.aliyun.com/document_detail/255797.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// ca-edd3626622af00b3****
	AggregatorId *string `json:"AggregatorId,omitempty" xml:"AggregatorId,omitempty"`
	// The rule ID.
	//
	// For more information about how to obtain the ID of a rule, see [ListAggregateConfigRules](https://help.aliyun.com/document_detail/264148.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// cr-2792626622af0c21****
	ConfigRuleId *string `json:"ConfigRuleId,omitempty" xml:"ConfigRuleId,omitempty"`
	// The status of the remediation. Valid values:
	//
	// 	- Success
	//
	// 	- Failed
	//
	// example:
	//
	// Success
	ExecutionStatus *string `json:"ExecutionStatus,omitempty" xml:"ExecutionStatus,omitempty"`
	// The maximum number of entries to return for a single request. Valid values: 10 to 100.
	//
	// example:
	//
	// 10
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. You do not need to specify this parameter for the first request. You must specify the token that is obtained from the previous query as the value of NextToken.
	//
	// example:
	//
	// aVCjqFaSy0Ps6zSMGw09****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the member account in the account group. When left unfilled, this rule queries the remediation result for the account that created the rule. If the account is not in the account group, the result will be empty.
	//
	// example:
	//
	// 126672004088****
	ResourceAccountId *int64 `json:"ResourceAccountId,omitempty" xml:"ResourceAccountId,omitempty"`
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
