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
	// This parameter is required.
	ConfigRuleId    *string `json:"ConfigRuleId,omitempty" xml:"ConfigRuleId,omitempty"`
	ExecutionStatus *string `json:"ExecutionStatus,omitempty" xml:"ExecutionStatus,omitempty"`
	MaxResults      *int64  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken       *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
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
