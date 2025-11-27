// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRCInvocationResultsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInvocationResults(v []*DescribeRCInvocationResultsResponseBodyInvocationResults) *DescribeRCInvocationResultsResponseBody
	GetInvocationResults() []*DescribeRCInvocationResultsResponseBodyInvocationResults
	SetNextToken(v string) *DescribeRCInvocationResultsResponseBody
	GetNextToken() *string
	SetPageNumber(v string) *DescribeRCInvocationResultsResponseBody
	GetPageNumber() *string
	SetPageSize(v string) *DescribeRCInvocationResultsResponseBody
	GetPageSize() *string
	SetRequestId(v string) *DescribeRCInvocationResultsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeRCInvocationResultsResponseBody
	GetTotalCount() *int32
}

type DescribeRCInvocationResultsResponseBody struct {
	InvocationResults []*DescribeRCInvocationResultsResponseBodyInvocationResults `json:"InvocationResults,omitempty" xml:"InvocationResults,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// AAAAAdDWBF2
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 1
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 49BC2500-8078-5AC4-A545-20AA5945B0E6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeRCInvocationResultsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRCInvocationResultsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRCInvocationResultsResponseBody) GetInvocationResults() []*DescribeRCInvocationResultsResponseBodyInvocationResults {
	return s.InvocationResults
}

func (s *DescribeRCInvocationResultsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeRCInvocationResultsResponseBody) GetPageNumber() *string {
	return s.PageNumber
}

func (s *DescribeRCInvocationResultsResponseBody) GetPageSize() *string {
	return s.PageSize
}

func (s *DescribeRCInvocationResultsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRCInvocationResultsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeRCInvocationResultsResponseBody) SetInvocationResults(v []*DescribeRCInvocationResultsResponseBodyInvocationResults) *DescribeRCInvocationResultsResponseBody {
	s.InvocationResults = v
	return s
}

func (s *DescribeRCInvocationResultsResponseBody) SetNextToken(v string) *DescribeRCInvocationResultsResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeRCInvocationResultsResponseBody) SetPageNumber(v string) *DescribeRCInvocationResultsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeRCInvocationResultsResponseBody) SetPageSize(v string) *DescribeRCInvocationResultsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeRCInvocationResultsResponseBody) SetRequestId(v string) *DescribeRCInvocationResultsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRCInvocationResultsResponseBody) SetTotalCount(v int32) *DescribeRCInvocationResultsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeRCInvocationResultsResponseBody) Validate() error {
	if s.InvocationResults != nil {
		for _, item := range s.InvocationResults {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeRCInvocationResultsResponseBodyInvocationResults struct {
	// example:
	//
	// c-7d2a745b412b4601b2d47f6a768d****
	CommandId *string `json:"CommandId,omitempty" xml:"CommandId,omitempty"`
	// example:
	//
	// ab141ddfbacfe02d9dbc25966ed971536124527097398d419a6746873fea****
	ContainerId *string `json:"ContainerId,omitempty" xml:"ContainerId,omitempty"`
	// example:
	//
	// test-container
	ContainerName *string `json:"ContainerName,omitempty" xml:"ContainerName,omitempty"`
	// example:
	//
	// 0
	Dropped *int32 `json:"Dropped,omitempty" xml:"Dropped,omitempty"`
	// example:
	//
	// InstanceNotExists
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// the specified instance does not exists
	ErrorInfo *string `json:"ErrorInfo,omitempty" xml:"ErrorInfo,omitempty"`
	// example:
	//
	// 0
	ExitCode *int32 `json:"ExitCode,omitempty" xml:"ExitCode,omitempty"`
	// example:
	//
	// 2024-12-20T06:15:56Z
	FinishedTime *string `json:"FinishedTime,omitempty" xml:"FinishedTime,omitempty"`
	// example:
	//
	// rc-i322y2t562oh7o******
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// Success
	InvocationStatus *string `json:"InvocationStatus,omitempty" xml:"InvocationStatus,omitempty"`
	// example:
	//
	// t-7d2a745b412b4601b2d47f6a768d****
	InvokeId *string `json:"InvokeId,omitempty" xml:"InvokeId,omitempty"`
	// example:
	//
	// Running
	InvokeRecordStatus *string `json:"InvokeRecordStatus,omitempty" xml:"InvokeRecordStatus,omitempty"`
	// example:
	//
	// MTU6MzA6MDEK
	Output *string `json:"Output,omitempty" xml:"Output,omitempty"`
	// example:
	//
	// 0
	Repeats *string `json:"Repeats,omitempty" xml:"Repeats,omitempty"`
	// example:
	//
	// 2024-12-20T06:15:55Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// 2025-01-19T09:15:47Z
	StopTime *string                                                         `json:"StopTime,omitempty" xml:"StopTime,omitempty"`
	Tags     []*DescribeRCInvocationResultsResponseBodyInvocationResultsTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// example:
	//
	// testuser
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
}

func (s DescribeRCInvocationResultsResponseBodyInvocationResults) String() string {
	return dara.Prettify(s)
}

func (s DescribeRCInvocationResultsResponseBodyInvocationResults) GoString() string {
	return s.String()
}

func (s *DescribeRCInvocationResultsResponseBodyInvocationResults) GetCommandId() *string {
	return s.CommandId
}

func (s *DescribeRCInvocationResultsResponseBodyInvocationResults) GetContainerId() *string {
	return s.ContainerId
}

func (s *DescribeRCInvocationResultsResponseBodyInvocationResults) GetContainerName() *string {
	return s.ContainerName
}

func (s *DescribeRCInvocationResultsResponseBodyInvocationResults) GetDropped() *int32 {
	return s.Dropped
}

func (s *DescribeRCInvocationResultsResponseBodyInvocationResults) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DescribeRCInvocationResultsResponseBodyInvocationResults) GetErrorInfo() *string {
	return s.ErrorInfo
}

func (s *DescribeRCInvocationResultsResponseBodyInvocationResults) GetExitCode() *int32 {
	return s.ExitCode
}

func (s *DescribeRCInvocationResultsResponseBodyInvocationResults) GetFinishedTime() *string {
	return s.FinishedTime
}

func (s *DescribeRCInvocationResultsResponseBodyInvocationResults) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeRCInvocationResultsResponseBodyInvocationResults) GetInvocationStatus() *string {
	return s.InvocationStatus
}

func (s *DescribeRCInvocationResultsResponseBodyInvocationResults) GetInvokeId() *string {
	return s.InvokeId
}

func (s *DescribeRCInvocationResultsResponseBodyInvocationResults) GetInvokeRecordStatus() *string {
	return s.InvokeRecordStatus
}

func (s *DescribeRCInvocationResultsResponseBodyInvocationResults) GetOutput() *string {
	return s.Output
}

func (s *DescribeRCInvocationResultsResponseBodyInvocationResults) GetRepeats() *string {
	return s.Repeats
}

func (s *DescribeRCInvocationResultsResponseBodyInvocationResults) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeRCInvocationResultsResponseBodyInvocationResults) GetStopTime() *string {
	return s.StopTime
}

func (s *DescribeRCInvocationResultsResponseBodyInvocationResults) GetTags() []*DescribeRCInvocationResultsResponseBodyInvocationResultsTags {
	return s.Tags
}

func (s *DescribeRCInvocationResultsResponseBodyInvocationResults) GetUsername() *string {
	return s.Username
}

func (s *DescribeRCInvocationResultsResponseBodyInvocationResults) SetCommandId(v string) *DescribeRCInvocationResultsResponseBodyInvocationResults {
	s.CommandId = &v
	return s
}

func (s *DescribeRCInvocationResultsResponseBodyInvocationResults) SetContainerId(v string) *DescribeRCInvocationResultsResponseBodyInvocationResults {
	s.ContainerId = &v
	return s
}

func (s *DescribeRCInvocationResultsResponseBodyInvocationResults) SetContainerName(v string) *DescribeRCInvocationResultsResponseBodyInvocationResults {
	s.ContainerName = &v
	return s
}

func (s *DescribeRCInvocationResultsResponseBodyInvocationResults) SetDropped(v int32) *DescribeRCInvocationResultsResponseBodyInvocationResults {
	s.Dropped = &v
	return s
}

func (s *DescribeRCInvocationResultsResponseBodyInvocationResults) SetErrorCode(v string) *DescribeRCInvocationResultsResponseBodyInvocationResults {
	s.ErrorCode = &v
	return s
}

func (s *DescribeRCInvocationResultsResponseBodyInvocationResults) SetErrorInfo(v string) *DescribeRCInvocationResultsResponseBodyInvocationResults {
	s.ErrorInfo = &v
	return s
}

func (s *DescribeRCInvocationResultsResponseBodyInvocationResults) SetExitCode(v int32) *DescribeRCInvocationResultsResponseBodyInvocationResults {
	s.ExitCode = &v
	return s
}

func (s *DescribeRCInvocationResultsResponseBodyInvocationResults) SetFinishedTime(v string) *DescribeRCInvocationResultsResponseBodyInvocationResults {
	s.FinishedTime = &v
	return s
}

func (s *DescribeRCInvocationResultsResponseBodyInvocationResults) SetInstanceId(v string) *DescribeRCInvocationResultsResponseBodyInvocationResults {
	s.InstanceId = &v
	return s
}

func (s *DescribeRCInvocationResultsResponseBodyInvocationResults) SetInvocationStatus(v string) *DescribeRCInvocationResultsResponseBodyInvocationResults {
	s.InvocationStatus = &v
	return s
}

func (s *DescribeRCInvocationResultsResponseBodyInvocationResults) SetInvokeId(v string) *DescribeRCInvocationResultsResponseBodyInvocationResults {
	s.InvokeId = &v
	return s
}

func (s *DescribeRCInvocationResultsResponseBodyInvocationResults) SetInvokeRecordStatus(v string) *DescribeRCInvocationResultsResponseBodyInvocationResults {
	s.InvokeRecordStatus = &v
	return s
}

func (s *DescribeRCInvocationResultsResponseBodyInvocationResults) SetOutput(v string) *DescribeRCInvocationResultsResponseBodyInvocationResults {
	s.Output = &v
	return s
}

func (s *DescribeRCInvocationResultsResponseBodyInvocationResults) SetRepeats(v string) *DescribeRCInvocationResultsResponseBodyInvocationResults {
	s.Repeats = &v
	return s
}

func (s *DescribeRCInvocationResultsResponseBodyInvocationResults) SetStartTime(v string) *DescribeRCInvocationResultsResponseBodyInvocationResults {
	s.StartTime = &v
	return s
}

func (s *DescribeRCInvocationResultsResponseBodyInvocationResults) SetStopTime(v string) *DescribeRCInvocationResultsResponseBodyInvocationResults {
	s.StopTime = &v
	return s
}

func (s *DescribeRCInvocationResultsResponseBodyInvocationResults) SetTags(v []*DescribeRCInvocationResultsResponseBodyInvocationResultsTags) *DescribeRCInvocationResultsResponseBodyInvocationResults {
	s.Tags = v
	return s
}

func (s *DescribeRCInvocationResultsResponseBodyInvocationResults) SetUsername(v string) *DescribeRCInvocationResultsResponseBodyInvocationResults {
	s.Username = &v
	return s
}

func (s *DescribeRCInvocationResultsResponseBodyInvocationResults) Validate() error {
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeRCInvocationResultsResponseBodyInvocationResultsTags struct {
	// example:
	//
	// testKey
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// example:
	//
	// testValue
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s DescribeRCInvocationResultsResponseBodyInvocationResultsTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeRCInvocationResultsResponseBodyInvocationResultsTags) GoString() string {
	return s.String()
}

func (s *DescribeRCInvocationResultsResponseBodyInvocationResultsTags) GetTagKey() *string {
	return s.TagKey
}

func (s *DescribeRCInvocationResultsResponseBodyInvocationResultsTags) GetTagValue() *string {
	return s.TagValue
}

func (s *DescribeRCInvocationResultsResponseBodyInvocationResultsTags) SetTagKey(v string) *DescribeRCInvocationResultsResponseBodyInvocationResultsTags {
	s.TagKey = &v
	return s
}

func (s *DescribeRCInvocationResultsResponseBodyInvocationResultsTags) SetTagValue(v string) *DescribeRCInvocationResultsResponseBodyInvocationResultsTags {
	s.TagValue = &v
	return s
}

func (s *DescribeRCInvocationResultsResponseBodyInvocationResultsTags) Validate() error {
	return dara.Validate(s)
}
