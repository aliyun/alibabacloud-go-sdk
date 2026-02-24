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
	RemediationExecutionData *ListRemediationExecutionsResponseBodyRemediationExecutionData `json:"RemediationExecutionData,omitempty" xml:"RemediationExecutionData,omitempty" type:"Struct"`
	RequestId                *string                                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	MaxResults            *int64                                                                                `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken             *string                                                                               `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
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
	ExecutionCreateDate    *string `json:"ExecutionCreateDate,omitempty" xml:"ExecutionCreateDate,omitempty"`
	ExecutionInvocationId  *string `json:"ExecutionInvocationId,omitempty" xml:"ExecutionInvocationId,omitempty"`
	ExecutionResourceIds   *string `json:"ExecutionResourceIds,omitempty" xml:"ExecutionResourceIds,omitempty"`
	ExecutionResourceType  *string `json:"ExecutionResourceType,omitempty" xml:"ExecutionResourceType,omitempty"`
	ExecutionStatus        *string `json:"ExecutionStatus,omitempty" xml:"ExecutionStatus,omitempty"`
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
