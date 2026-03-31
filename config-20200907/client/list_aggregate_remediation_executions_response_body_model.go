// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAggregateRemediationExecutionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRemediationExecutionData(v *ListAggregateRemediationExecutionsResponseBodyRemediationExecutionData) *ListAggregateRemediationExecutionsResponseBody
	GetRemediationExecutionData() *ListAggregateRemediationExecutionsResponseBodyRemediationExecutionData
	SetRequestId(v string) *ListAggregateRemediationExecutionsResponseBody
	GetRequestId() *string
}

type ListAggregateRemediationExecutionsResponseBody struct {
	// The queried remediation records.
	RemediationExecutionData *ListAggregateRemediationExecutionsResponseBodyRemediationExecutionData `json:"RemediationExecutionData,omitempty" xml:"RemediationExecutionData,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 13E67493-3165-529A-A961-BE9E4B11BA11
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListAggregateRemediationExecutionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAggregateRemediationExecutionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAggregateRemediationExecutionsResponseBody) GetRemediationExecutionData() *ListAggregateRemediationExecutionsResponseBodyRemediationExecutionData {
	return s.RemediationExecutionData
}

func (s *ListAggregateRemediationExecutionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAggregateRemediationExecutionsResponseBody) SetRemediationExecutionData(v *ListAggregateRemediationExecutionsResponseBodyRemediationExecutionData) *ListAggregateRemediationExecutionsResponseBody {
	s.RemediationExecutionData = v
	return s
}

func (s *ListAggregateRemediationExecutionsResponseBody) SetRequestId(v string) *ListAggregateRemediationExecutionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAggregateRemediationExecutionsResponseBody) Validate() error {
	if s.RemediationExecutionData != nil {
		if err := s.RemediationExecutionData.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListAggregateRemediationExecutionsResponseBodyRemediationExecutionData struct {
	// The maximum number of entries returned for a single request.
	//
	// example:
	//
	// 10
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// A pagination token. It can be used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// cNclqNaKs0Ds7zSNip0a****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The queried remediation records.
	RemediationExecutions []*ListAggregateRemediationExecutionsResponseBodyRemediationExecutionDataRemediationExecutions `json:"RemediationExecutions,omitempty" xml:"RemediationExecutions,omitempty" type:"Repeated"`
}

func (s ListAggregateRemediationExecutionsResponseBodyRemediationExecutionData) String() string {
	return dara.Prettify(s)
}

func (s ListAggregateRemediationExecutionsResponseBodyRemediationExecutionData) GoString() string {
	return s.String()
}

func (s *ListAggregateRemediationExecutionsResponseBodyRemediationExecutionData) GetMaxResults() *int64 {
	return s.MaxResults
}

func (s *ListAggregateRemediationExecutionsResponseBodyRemediationExecutionData) GetNextToken() *string {
	return s.NextToken
}

func (s *ListAggregateRemediationExecutionsResponseBodyRemediationExecutionData) GetRemediationExecutions() []*ListAggregateRemediationExecutionsResponseBodyRemediationExecutionDataRemediationExecutions {
	return s.RemediationExecutions
}

func (s *ListAggregateRemediationExecutionsResponseBodyRemediationExecutionData) SetMaxResults(v int64) *ListAggregateRemediationExecutionsResponseBodyRemediationExecutionData {
	s.MaxResults = &v
	return s
}

func (s *ListAggregateRemediationExecutionsResponseBodyRemediationExecutionData) SetNextToken(v string) *ListAggregateRemediationExecutionsResponseBodyRemediationExecutionData {
	s.NextToken = &v
	return s
}

func (s *ListAggregateRemediationExecutionsResponseBodyRemediationExecutionData) SetRemediationExecutions(v []*ListAggregateRemediationExecutionsResponseBodyRemediationExecutionDataRemediationExecutions) *ListAggregateRemediationExecutionsResponseBodyRemediationExecutionData {
	s.RemediationExecutions = v
	return s
}

func (s *ListAggregateRemediationExecutionsResponseBodyRemediationExecutionData) Validate() error {
	if s.RemediationExecutions != nil {
		for _, item := range s.RemediationExecutions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAggregateRemediationExecutionsResponseBodyRemediationExecutionDataRemediationExecutions struct {
	// The time when the remediation record was created.
	//
	// example:
	//
	// 2023-06-25T11:48:15Z
	ExecutionCreateDate *string `json:"ExecutionCreateDate,omitempty" xml:"ExecutionCreateDate,omitempty"`
	// The invocation ID of the remediation record.
	//
	// example:
	//
	// exec-befded3781994ccf****
	ExecutionInvocationId *string `json:"ExecutionInvocationId,omitempty" xml:"ExecutionInvocationId,omitempty"`
	// The IDs of the remediated resources. Multiple resource IDs are separated with commas (,).
	//
	// example:
	//
	// rm-0jlk629z240l8****
	ExecutionResourceIds *string `json:"ExecutionResourceIds,omitempty" xml:"ExecutionResourceIds,omitempty"`
	// The resource type.
	//
	// example:
	//
	// ACS::RDS::DBInstance
	ExecutionResourceType *string `json:"ExecutionResourceType,omitempty" xml:"ExecutionResourceType,omitempty"`
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
	// The error message returned when the remediation fails.
	//
	// example:
	//
	// Invocation time out.
	ExecutionStatusMessage *string `json:"ExecutionStatusMessage,omitempty" xml:"ExecutionStatusMessage,omitempty"`
}

func (s ListAggregateRemediationExecutionsResponseBodyRemediationExecutionDataRemediationExecutions) String() string {
	return dara.Prettify(s)
}

func (s ListAggregateRemediationExecutionsResponseBodyRemediationExecutionDataRemediationExecutions) GoString() string {
	return s.String()
}

func (s *ListAggregateRemediationExecutionsResponseBodyRemediationExecutionDataRemediationExecutions) GetExecutionCreateDate() *string {
	return s.ExecutionCreateDate
}

func (s *ListAggregateRemediationExecutionsResponseBodyRemediationExecutionDataRemediationExecutions) GetExecutionInvocationId() *string {
	return s.ExecutionInvocationId
}

func (s *ListAggregateRemediationExecutionsResponseBodyRemediationExecutionDataRemediationExecutions) GetExecutionResourceIds() *string {
	return s.ExecutionResourceIds
}

func (s *ListAggregateRemediationExecutionsResponseBodyRemediationExecutionDataRemediationExecutions) GetExecutionResourceType() *string {
	return s.ExecutionResourceType
}

func (s *ListAggregateRemediationExecutionsResponseBodyRemediationExecutionDataRemediationExecutions) GetExecutionStatus() *string {
	return s.ExecutionStatus
}

func (s *ListAggregateRemediationExecutionsResponseBodyRemediationExecutionDataRemediationExecutions) GetExecutionStatusMessage() *string {
	return s.ExecutionStatusMessage
}

func (s *ListAggregateRemediationExecutionsResponseBodyRemediationExecutionDataRemediationExecutions) SetExecutionCreateDate(v string) *ListAggregateRemediationExecutionsResponseBodyRemediationExecutionDataRemediationExecutions {
	s.ExecutionCreateDate = &v
	return s
}

func (s *ListAggregateRemediationExecutionsResponseBodyRemediationExecutionDataRemediationExecutions) SetExecutionInvocationId(v string) *ListAggregateRemediationExecutionsResponseBodyRemediationExecutionDataRemediationExecutions {
	s.ExecutionInvocationId = &v
	return s
}

func (s *ListAggregateRemediationExecutionsResponseBodyRemediationExecutionDataRemediationExecutions) SetExecutionResourceIds(v string) *ListAggregateRemediationExecutionsResponseBodyRemediationExecutionDataRemediationExecutions {
	s.ExecutionResourceIds = &v
	return s
}

func (s *ListAggregateRemediationExecutionsResponseBodyRemediationExecutionDataRemediationExecutions) SetExecutionResourceType(v string) *ListAggregateRemediationExecutionsResponseBodyRemediationExecutionDataRemediationExecutions {
	s.ExecutionResourceType = &v
	return s
}

func (s *ListAggregateRemediationExecutionsResponseBodyRemediationExecutionDataRemediationExecutions) SetExecutionStatus(v string) *ListAggregateRemediationExecutionsResponseBodyRemediationExecutionDataRemediationExecutions {
	s.ExecutionStatus = &v
	return s
}

func (s *ListAggregateRemediationExecutionsResponseBodyRemediationExecutionDataRemediationExecutions) SetExecutionStatusMessage(v string) *ListAggregateRemediationExecutionsResponseBodyRemediationExecutionDataRemediationExecutions {
	s.ExecutionStatusMessage = &v
	return s
}

func (s *ListAggregateRemediationExecutionsResponseBodyRemediationExecutionDataRemediationExecutions) Validate() error {
	return dara.Validate(s)
}
