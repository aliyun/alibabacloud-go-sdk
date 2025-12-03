// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPipelineRunsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *ListPipelineRunsResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ListPipelineRunsResponseBody
	GetErrorMessage() *string
	SetNextToken(v string) *ListPipelineRunsResponseBody
	GetNextToken() *string
	SetPipelineRuns(v []*ListPipelineRunsResponseBodyPipelineRuns) *ListPipelineRunsResponseBody
	GetPipelineRuns() []*ListPipelineRunsResponseBodyPipelineRuns
	SetRequestId(v string) *ListPipelineRunsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListPipelineRunsResponseBody
	GetSuccess() *bool
	SetTotalCount(v int64) *ListPipelineRunsResponseBody
	GetTotalCount() *int64
}

type ListPipelineRunsResponseBody struct {
	// example:
	//
	// ”“
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// ”“
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// example:
	//
	// xzxsasasaas
	NextToken    *string                                     `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	PipelineRuns []*ListPipelineRunsResponseBodyPipelineRuns `json:"pipelineRuns,omitempty" xml:"pipelineRuns,omitempty" type:"Repeated"`
	// example:
	//
	// ASSDS-ASSASX-XSAXSA-XSAXSAXS
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 50
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListPipelineRunsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListPipelineRunsResponseBody) GoString() string {
	return s.String()
}

func (s *ListPipelineRunsResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListPipelineRunsResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListPipelineRunsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListPipelineRunsResponseBody) GetPipelineRuns() []*ListPipelineRunsResponseBodyPipelineRuns {
	return s.PipelineRuns
}

func (s *ListPipelineRunsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListPipelineRunsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListPipelineRunsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListPipelineRunsResponseBody) SetErrorCode(v string) *ListPipelineRunsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListPipelineRunsResponseBody) SetErrorMessage(v string) *ListPipelineRunsResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListPipelineRunsResponseBody) SetNextToken(v string) *ListPipelineRunsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListPipelineRunsResponseBody) SetPipelineRuns(v []*ListPipelineRunsResponseBodyPipelineRuns) *ListPipelineRunsResponseBody {
	s.PipelineRuns = v
	return s
}

func (s *ListPipelineRunsResponseBody) SetRequestId(v string) *ListPipelineRunsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPipelineRunsResponseBody) SetSuccess(v bool) *ListPipelineRunsResponseBody {
	s.Success = &v
	return s
}

func (s *ListPipelineRunsResponseBody) SetTotalCount(v int64) *ListPipelineRunsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListPipelineRunsResponseBody) Validate() error {
	if s.PipelineRuns != nil {
		for _, item := range s.PipelineRuns {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListPipelineRunsResponseBodyPipelineRuns struct {
	// example:
	//
	// 1222222222
	CreatorAccountId *string `json:"creatorAccountId,omitempty" xml:"creatorAccountId,omitempty"`
	// example:
	//
	// 1586863220000
	EndTime *int64 `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// example:
	//
	// 1234
	PipelineId *int64 `json:"pipelineId,omitempty" xml:"pipelineId,omitempty"`
	// example:
	//
	// 11
	PipelineRunId *int64 `json:"pipelineRunId,omitempty" xml:"pipelineRunId,omitempty"`
	// example:
	//
	// 1586863220000
	StartTime *int64 `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// example:
	//
	// 状态 FAIL 运行失败 SUCCESS 运行成功 RUNNING 运行中
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// 1
	TriggerMode *int64 `json:"triggerMode,omitempty" xml:"triggerMode,omitempty"`
}

func (s ListPipelineRunsResponseBodyPipelineRuns) String() string {
	return dara.Prettify(s)
}

func (s ListPipelineRunsResponseBodyPipelineRuns) GoString() string {
	return s.String()
}

func (s *ListPipelineRunsResponseBodyPipelineRuns) GetCreatorAccountId() *string {
	return s.CreatorAccountId
}

func (s *ListPipelineRunsResponseBodyPipelineRuns) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ListPipelineRunsResponseBodyPipelineRuns) GetPipelineId() *int64 {
	return s.PipelineId
}

func (s *ListPipelineRunsResponseBodyPipelineRuns) GetPipelineRunId() *int64 {
	return s.PipelineRunId
}

func (s *ListPipelineRunsResponseBodyPipelineRuns) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ListPipelineRunsResponseBodyPipelineRuns) GetStatus() *string {
	return s.Status
}

func (s *ListPipelineRunsResponseBodyPipelineRuns) GetTriggerMode() *int64 {
	return s.TriggerMode
}

func (s *ListPipelineRunsResponseBodyPipelineRuns) SetCreatorAccountId(v string) *ListPipelineRunsResponseBodyPipelineRuns {
	s.CreatorAccountId = &v
	return s
}

func (s *ListPipelineRunsResponseBodyPipelineRuns) SetEndTime(v int64) *ListPipelineRunsResponseBodyPipelineRuns {
	s.EndTime = &v
	return s
}

func (s *ListPipelineRunsResponseBodyPipelineRuns) SetPipelineId(v int64) *ListPipelineRunsResponseBodyPipelineRuns {
	s.PipelineId = &v
	return s
}

func (s *ListPipelineRunsResponseBodyPipelineRuns) SetPipelineRunId(v int64) *ListPipelineRunsResponseBodyPipelineRuns {
	s.PipelineRunId = &v
	return s
}

func (s *ListPipelineRunsResponseBodyPipelineRuns) SetStartTime(v int64) *ListPipelineRunsResponseBodyPipelineRuns {
	s.StartTime = &v
	return s
}

func (s *ListPipelineRunsResponseBodyPipelineRuns) SetStatus(v string) *ListPipelineRunsResponseBodyPipelineRuns {
	s.Status = &v
	return s
}

func (s *ListPipelineRunsResponseBodyPipelineRuns) SetTriggerMode(v int64) *ListPipelineRunsResponseBodyPipelineRuns {
	s.TriggerMode = &v
	return s
}

func (s *ListPipelineRunsResponseBodyPipelineRuns) Validate() error {
	return dara.Validate(s)
}
