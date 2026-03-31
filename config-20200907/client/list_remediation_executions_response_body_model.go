// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRemediationExecutionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRemediationExecutionData(v *ListRemediationExecutionsResponseBodyRemediationExecutionData) *ListRemediationExecutionsResponseBody
	GetRemediationExecutionData() *ListRemediationExecutionsResponseBodyRemediationExecutionData
	SetRequestId(v string) *ListRemediationExecutionsResponseBody
	GetRequestId() *string
}

type ListRemediationExecutionsResponseBody struct {
	// The queried remediation records.
	RemediationExecutionData *ListRemediationExecutionsResponseBodyRemediationExecutionData `json:"RemediationExecutionData,omitempty" xml:"RemediationExecutionData,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 13E67493-3165-529A-A961-BE9E4B11BA11
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListRemediationExecutionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListRemediationExecutionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListRemediationExecutionsResponseBody) GetRemediationExecutionData() *ListRemediationExecutionsResponseBodyRemediationExecutionData {
	return s.RemediationExecutionData
}

func (s *ListRemediationExecutionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListRemediationExecutionsResponseBody) SetRemediationExecutionData(v *ListRemediationExecutionsResponseBodyRemediationExecutionData) *ListRemediationExecutionsResponseBody {
	s.RemediationExecutionData = v
	return s
}

func (s *ListRemediationExecutionsResponseBody) SetRequestId(v string) *ListRemediationExecutionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRemediationExecutionsResponseBody) Validate() error {
	if s.RemediationExecutionData != nil {
		if err := s.RemediationExecutionData.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListRemediationExecutionsResponseBodyRemediationExecutionData struct {
	// The maximum number of entries to return for a single request.
	//
	// example:
	//
	// 10
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// A pagination token. It can be used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// cMbjqNaYs0Ps7zSNiu37****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The queried remediation records.
	RemediationExecutions []*ListRemediationExecutionsResponseBodyRemediationExecutionDataRemediationExecutions `json:"RemediationExecutions,omitempty" xml:"RemediationExecutions,omitempty" type:"Repeated"`
}

func (s ListRemediationExecutionsResponseBodyRemediationExecutionData) String() string {
	return dara.Prettify(s)
}

func (s ListRemediationExecutionsResponseBodyRemediationExecutionData) GoString() string {
	return s.String()
}

func (s *ListRemediationExecutionsResponseBodyRemediationExecutionData) GetMaxResults() *int64 {
	return s.MaxResults
}

func (s *ListRemediationExecutionsResponseBodyRemediationExecutionData) GetNextToken() *string {
	return s.NextToken
}

func (s *ListRemediationExecutionsResponseBodyRemediationExecutionData) GetRemediationExecutions() []*ListRemediationExecutionsResponseBodyRemediationExecutionDataRemediationExecutions {
	return s.RemediationExecutions
}

func (s *ListRemediationExecutionsResponseBodyRemediationExecutionData) SetMaxResults(v int64) *ListRemediationExecutionsResponseBodyRemediationExecutionData {
	s.MaxResults = &v
	return s
}

func (s *ListRemediationExecutionsResponseBodyRemediationExecutionData) SetNextToken(v string) *ListRemediationExecutionsResponseBodyRemediationExecutionData {
	s.NextToken = &v
	return s
}

func (s *ListRemediationExecutionsResponseBodyRemediationExecutionData) SetRemediationExecutions(v []*ListRemediationExecutionsResponseBodyRemediationExecutionDataRemediationExecutions) *ListRemediationExecutionsResponseBodyRemediationExecutionData {
	s.RemediationExecutions = v
	return s
}

func (s *ListRemediationExecutionsResponseBodyRemediationExecutionData) Validate() error {
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

type ListRemediationExecutionsResponseBodyRemediationExecutionDataRemediationExecutions struct {
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
	// The IDs of the resources to which the remediation belongs. Separate multiple resource IDs with commas (,).
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
	// The status of the remediation record. Valid values:
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

func (s ListRemediationExecutionsResponseBodyRemediationExecutionDataRemediationExecutions) String() string {
	return dara.Prettify(s)
}

func (s ListRemediationExecutionsResponseBodyRemediationExecutionDataRemediationExecutions) GoString() string {
	return s.String()
}

func (s *ListRemediationExecutionsResponseBodyRemediationExecutionDataRemediationExecutions) GetExecutionCreateDate() *string {
	return s.ExecutionCreateDate
}

func (s *ListRemediationExecutionsResponseBodyRemediationExecutionDataRemediationExecutions) GetExecutionInvocationId() *string {
	return s.ExecutionInvocationId
}

func (s *ListRemediationExecutionsResponseBodyRemediationExecutionDataRemediationExecutions) GetExecutionResourceIds() *string {
	return s.ExecutionResourceIds
}

func (s *ListRemediationExecutionsResponseBodyRemediationExecutionDataRemediationExecutions) GetExecutionResourceType() *string {
	return s.ExecutionResourceType
}

func (s *ListRemediationExecutionsResponseBodyRemediationExecutionDataRemediationExecutions) GetExecutionStatus() *string {
	return s.ExecutionStatus
}

func (s *ListRemediationExecutionsResponseBodyRemediationExecutionDataRemediationExecutions) GetExecutionStatusMessage() *string {
	return s.ExecutionStatusMessage
}

func (s *ListRemediationExecutionsResponseBodyRemediationExecutionDataRemediationExecutions) SetExecutionCreateDate(v string) *ListRemediationExecutionsResponseBodyRemediationExecutionDataRemediationExecutions {
	s.ExecutionCreateDate = &v
	return s
}

func (s *ListRemediationExecutionsResponseBodyRemediationExecutionDataRemediationExecutions) SetExecutionInvocationId(v string) *ListRemediationExecutionsResponseBodyRemediationExecutionDataRemediationExecutions {
	s.ExecutionInvocationId = &v
	return s
}

func (s *ListRemediationExecutionsResponseBodyRemediationExecutionDataRemediationExecutions) SetExecutionResourceIds(v string) *ListRemediationExecutionsResponseBodyRemediationExecutionDataRemediationExecutions {
	s.ExecutionResourceIds = &v
	return s
}

func (s *ListRemediationExecutionsResponseBodyRemediationExecutionDataRemediationExecutions) SetExecutionResourceType(v string) *ListRemediationExecutionsResponseBodyRemediationExecutionDataRemediationExecutions {
	s.ExecutionResourceType = &v
	return s
}

func (s *ListRemediationExecutionsResponseBodyRemediationExecutionDataRemediationExecutions) SetExecutionStatus(v string) *ListRemediationExecutionsResponseBodyRemediationExecutionDataRemediationExecutions {
	s.ExecutionStatus = &v
	return s
}

func (s *ListRemediationExecutionsResponseBodyRemediationExecutionDataRemediationExecutions) SetExecutionStatusMessage(v string) *ListRemediationExecutionsResponseBodyRemediationExecutionDataRemediationExecutions {
	s.ExecutionStatusMessage = &v
	return s
}

func (s *ListRemediationExecutionsResponseBodyRemediationExecutionDataRemediationExecutions) Validate() error {
	return dara.Validate(s)
}
