// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPipelineJobHistorysResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *ListPipelineJobHistorysResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ListPipelineJobHistorysResponseBody
	GetErrorMessage() *string
	SetJobs(v []*ListPipelineJobHistorysResponseBodyJobs) *ListPipelineJobHistorysResponseBody
	GetJobs() []*ListPipelineJobHistorysResponseBodyJobs
	SetNextToken(v string) *ListPipelineJobHistorysResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListPipelineJobHistorysResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListPipelineJobHistorysResponseBody
	GetSuccess() *bool
	SetTotalCount(v int32) *ListPipelineJobHistorysResponseBody
	GetTotalCount() *int32
}

type ListPipelineJobHistorysResponseBody struct {
	// example:
	//
	// ”“
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// ”“
	ErrorMessage *string                                    `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	Jobs         []*ListPipelineJobHistorysResponseBodyJobs `json:"jobs,omitempty" xml:"jobs,omitempty" type:"Repeated"`
	// example:
	//
	// xsxxs
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// example:
	//
	// ASSDS-ASSASX-XSAXSA-XSAXSAXS
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true 接口调用成功，false 接口调用失败
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 20
	TotalCount *int32 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListPipelineJobHistorysResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListPipelineJobHistorysResponseBody) GoString() string {
	return s.String()
}

func (s *ListPipelineJobHistorysResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListPipelineJobHistorysResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListPipelineJobHistorysResponseBody) GetJobs() []*ListPipelineJobHistorysResponseBodyJobs {
	return s.Jobs
}

func (s *ListPipelineJobHistorysResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListPipelineJobHistorysResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListPipelineJobHistorysResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListPipelineJobHistorysResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListPipelineJobHistorysResponseBody) SetErrorCode(v string) *ListPipelineJobHistorysResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListPipelineJobHistorysResponseBody) SetErrorMessage(v string) *ListPipelineJobHistorysResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListPipelineJobHistorysResponseBody) SetJobs(v []*ListPipelineJobHistorysResponseBodyJobs) *ListPipelineJobHistorysResponseBody {
	s.Jobs = v
	return s
}

func (s *ListPipelineJobHistorysResponseBody) SetNextToken(v string) *ListPipelineJobHistorysResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListPipelineJobHistorysResponseBody) SetRequestId(v string) *ListPipelineJobHistorysResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPipelineJobHistorysResponseBody) SetSuccess(v bool) *ListPipelineJobHistorysResponseBody {
	s.Success = &v
	return s
}

func (s *ListPipelineJobHistorysResponseBody) SetTotalCount(v int32) *ListPipelineJobHistorysResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListPipelineJobHistorysResponseBody) Validate() error {
	if s.Jobs != nil {
		for _, item := range s.Jobs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListPipelineJobHistorysResponseBodyJobs struct {
	// example:
	//
	// 1
	ExecuteNumber *int32 `json:"executeNumber,omitempty" xml:"executeNumber,omitempty"`
	// example:
	//
	// 10_xaxxsxa
	Identifier *string `json:"identifier,omitempty" xml:"identifier,omitempty"`
	// example:
	//
	// 123
	JobId   *int64  `json:"jobId,omitempty" xml:"jobId,omitempty"`
	JobName *string `json:"jobName,omitempty" xml:"jobName,omitempty"`
	// example:
	//
	// ssaasssa
	OperatorAccountId *string `json:"operatorAccountId,omitempty" xml:"operatorAccountId,omitempty"`
	// example:
	//
	// 123
	PipelineId *int64 `json:"pipelineId,omitempty" xml:"pipelineId,omitempty"`
	// example:
	//
	// 123
	PipelineRunId *int64 `json:"pipelineRunId,omitempty" xml:"pipelineRunId,omitempty"`
	// example:
	//
	// {}
	Sources *string `json:"sources,omitempty" xml:"sources,omitempty"`
	// example:
	//
	// SUCCESS
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s ListPipelineJobHistorysResponseBodyJobs) String() string {
	return dara.Prettify(s)
}

func (s ListPipelineJobHistorysResponseBodyJobs) GoString() string {
	return s.String()
}

func (s *ListPipelineJobHistorysResponseBodyJobs) GetExecuteNumber() *int32 {
	return s.ExecuteNumber
}

func (s *ListPipelineJobHistorysResponseBodyJobs) GetIdentifier() *string {
	return s.Identifier
}

func (s *ListPipelineJobHistorysResponseBodyJobs) GetJobId() *int64 {
	return s.JobId
}

func (s *ListPipelineJobHistorysResponseBodyJobs) GetJobName() *string {
	return s.JobName
}

func (s *ListPipelineJobHistorysResponseBodyJobs) GetOperatorAccountId() *string {
	return s.OperatorAccountId
}

func (s *ListPipelineJobHistorysResponseBodyJobs) GetPipelineId() *int64 {
	return s.PipelineId
}

func (s *ListPipelineJobHistorysResponseBodyJobs) GetPipelineRunId() *int64 {
	return s.PipelineRunId
}

func (s *ListPipelineJobHistorysResponseBodyJobs) GetSources() *string {
	return s.Sources
}

func (s *ListPipelineJobHistorysResponseBodyJobs) GetStatus() *string {
	return s.Status
}

func (s *ListPipelineJobHistorysResponseBodyJobs) SetExecuteNumber(v int32) *ListPipelineJobHistorysResponseBodyJobs {
	s.ExecuteNumber = &v
	return s
}

func (s *ListPipelineJobHistorysResponseBodyJobs) SetIdentifier(v string) *ListPipelineJobHistorysResponseBodyJobs {
	s.Identifier = &v
	return s
}

func (s *ListPipelineJobHistorysResponseBodyJobs) SetJobId(v int64) *ListPipelineJobHistorysResponseBodyJobs {
	s.JobId = &v
	return s
}

func (s *ListPipelineJobHistorysResponseBodyJobs) SetJobName(v string) *ListPipelineJobHistorysResponseBodyJobs {
	s.JobName = &v
	return s
}

func (s *ListPipelineJobHistorysResponseBodyJobs) SetOperatorAccountId(v string) *ListPipelineJobHistorysResponseBodyJobs {
	s.OperatorAccountId = &v
	return s
}

func (s *ListPipelineJobHistorysResponseBodyJobs) SetPipelineId(v int64) *ListPipelineJobHistorysResponseBodyJobs {
	s.PipelineId = &v
	return s
}

func (s *ListPipelineJobHistorysResponseBodyJobs) SetPipelineRunId(v int64) *ListPipelineJobHistorysResponseBodyJobs {
	s.PipelineRunId = &v
	return s
}

func (s *ListPipelineJobHistorysResponseBodyJobs) SetSources(v string) *ListPipelineJobHistorysResponseBodyJobs {
	s.Sources = &v
	return s
}

func (s *ListPipelineJobHistorysResponseBodyJobs) SetStatus(v string) *ListPipelineJobHistorysResponseBodyJobs {
	s.Status = &v
	return s
}

func (s *ListPipelineJobHistorysResponseBodyJobs) Validate() error {
	return dara.Validate(s)
}
