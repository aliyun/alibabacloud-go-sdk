// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRemediationExecutionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigRuleId(v string) *ListRemediationExecutionsRequest
	GetConfigRuleId() *string
	SetExecutionStatus(v string) *ListRemediationExecutionsRequest
	GetExecutionStatus() *string
	SetMaxResults(v int64) *ListRemediationExecutionsRequest
	GetMaxResults() *int64
	SetNextToken(v string) *ListRemediationExecutionsRequest
	GetNextToken() *string
}

type ListRemediationExecutionsRequest struct {
	// The rule ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cr-5392626622af0000****
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
	// aVCjqNaSy0Ps7zSMGu25****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListRemediationExecutionsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListRemediationExecutionsRequest) GoString() string {
	return s.String()
}

func (s *ListRemediationExecutionsRequest) GetConfigRuleId() *string {
	return s.ConfigRuleId
}

func (s *ListRemediationExecutionsRequest) GetExecutionStatus() *string {
	return s.ExecutionStatus
}

func (s *ListRemediationExecutionsRequest) GetMaxResults() *int64 {
	return s.MaxResults
}

func (s *ListRemediationExecutionsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListRemediationExecutionsRequest) SetConfigRuleId(v string) *ListRemediationExecutionsRequest {
	s.ConfigRuleId = &v
	return s
}

func (s *ListRemediationExecutionsRequest) SetExecutionStatus(v string) *ListRemediationExecutionsRequest {
	s.ExecutionStatus = &v
	return s
}

func (s *ListRemediationExecutionsRequest) SetMaxResults(v int64) *ListRemediationExecutionsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListRemediationExecutionsRequest) SetNextToken(v string) *ListRemediationExecutionsRequest {
	s.NextToken = &v
	return s
}

func (s *ListRemediationExecutionsRequest) Validate() error {
	return dara.Validate(s)
}
