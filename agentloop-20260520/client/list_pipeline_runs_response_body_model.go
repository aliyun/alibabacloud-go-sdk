// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPipelineRunsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListPipelineRunsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListPipelineRunsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListPipelineRunsResponseBody
	GetRequestId() *string
	SetRuns(v []*ListPipelineRunsResponseBodyRuns) *ListPipelineRunsResponseBody
	GetRuns() []*ListPipelineRunsResponseBodyRuns
	SetTotalCount(v int32) *ListPipelineRunsResponseBody
	GetTotalCount() *int32
}

type ListPipelineRunsResponseBody struct {
	// The maximum number of entries per page specified in the request. This value is echoed back.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The token for the next page. An empty string indicates that the current page is the last page.
	//
	// example:
	//
	// MTIzNDU2Nzg5MA==
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The request ID, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 9ACFB10A-1B2C-3D4E-5F6G-7H8I9J0K1L2M
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The list of pipeline runs.
	Runs []*ListPipelineRunsResponseBodyRuns `json:"runs,omitempty" xml:"runs,omitempty" type:"Repeated"`
	// The total number of pipeline runs that match the filter conditions.
	//
	// example:
	//
	// 100
	TotalCount *int32 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListPipelineRunsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListPipelineRunsResponseBody) GoString() string {
	return s.String()
}

func (s *ListPipelineRunsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListPipelineRunsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListPipelineRunsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListPipelineRunsResponseBody) GetRuns() []*ListPipelineRunsResponseBodyRuns {
	return s.Runs
}

func (s *ListPipelineRunsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListPipelineRunsResponseBody) SetMaxResults(v int32) *ListPipelineRunsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListPipelineRunsResponseBody) SetNextToken(v string) *ListPipelineRunsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListPipelineRunsResponseBody) SetRequestId(v string) *ListPipelineRunsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPipelineRunsResponseBody) SetRuns(v []*ListPipelineRunsResponseBodyRuns) *ListPipelineRunsResponseBody {
	s.Runs = v
	return s
}

func (s *ListPipelineRunsResponseBody) SetTotalCount(v int32) *ListPipelineRunsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListPipelineRunsResponseBody) Validate() error {
	if s.Runs != nil {
		for _, item := range s.Runs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListPipelineRunsResponseBodyRuns struct {
	// The name of the AgentSpace to which the run belongs.
	//
	// example:
	//
	// my-agent-space
	AgentSpaceName *string `json:"agentSpaceName,omitempty" xml:"agentSpaceName,omitempty"`
	// The current retry count.
	//
	// example:
	//
	// 0
	Attempt *int32 `json:"attempt,omitempty" xml:"attempt,omitempty"`
	// The semantic error code.
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// The error message.
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// The finish time in ISO 8601 UTC format.
	//
	// Use the UTC time format: yyyy-MM-ddTHH:mm:ss.SSSZ
	//
	// example:
	//
	// 2026-01-01T00:00:10.000Z
	FinishTime *string `json:"finishTime,omitempty" xml:"finishTime,omitempty"`
	// The start of the data window, in UNIX seconds.
	//
	// example:
	//
	// 1735660800
	FromTime *int64 `json:"fromTime,omitempty" xml:"fromTime,omitempty"`
	// The maximum number of retries.
	//
	// example:
	//
	// 3
	MaxAttempts *int32 `json:"maxAttempts,omitempty" xml:"maxAttempts,omitempty"`
	// The next retry time in ISO 8601 UTC format.
	//
	// Use the UTC time format: yyyy-MM-ddTHH:mm:ss.SSSZ
	//
	// example:
	//
	// 2026-01-01T00:01:00.000Z
	NextRetryTime *string `json:"nextRetryTime,omitempty" xml:"nextRetryTime,omitempty"`
	// The name of the pipeline.
	//
	// example:
	//
	// my-pipeline
	PipelineName *string `json:"pipelineName,omitempty" xml:"pipelineName,omitempty"`
	// Run Id
	//
	// example:
	//
	// run-20260101-0001
	RunId *string `json:"runId,omitempty" xml:"runId,omitempty"`
	// The execution start time in ISO 8601 UTC format.
	//
	// Use the UTC time format: yyyy-MM-ddTHH:mm:ss.SSSZ
	//
	// example:
	//
	// 2026-01-01T00:00:01.000Z
	StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// The run statistics, including processedRows (number of processed rows), processedBytes (number of processed bytes), outputRows (number of output rows), outputBytes (number of output bytes), elapsedMs (elapsed time in milliseconds), cpuSec (CPU seconds), cpuCores (number of CPU cores), and tokenCount (number of tokens consumed).
	Stats map[string]interface{} `json:"stats,omitempty" xml:"stats,omitempty"`
	// The run status. Valid values: Pending, Running, Succeeded, Failed, and Cancelled.
	//
	// example:
	//
	// Succeeded
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The end of the data window, in UNIX seconds.
	//
	// example:
	//
	// 1735747200
	ToTime *int64 `json:"toTime,omitempty" xml:"toTime,omitempty"`
	// The trigger time in ISO 8601 UTC format.
	//
	// Use the UTC time format: yyyy-MM-ddTHH:mm:ss.SSSZ
	//
	// example:
	//
	// 2026-01-01T00:00:00.000Z
	TriggerTime *string `json:"triggerTime,omitempty" xml:"triggerTime,omitempty"`
	// The trigger type. Valid values: Scheduled, Manual, and RunOnce.
	//
	// example:
	//
	// Scheduled
	TriggerType *string `json:"triggerType,omitempty" xml:"triggerType,omitempty"`
}

func (s ListPipelineRunsResponseBodyRuns) String() string {
	return dara.Prettify(s)
}

func (s ListPipelineRunsResponseBodyRuns) GoString() string {
	return s.String()
}

func (s *ListPipelineRunsResponseBodyRuns) GetAgentSpaceName() *string {
	return s.AgentSpaceName
}

func (s *ListPipelineRunsResponseBodyRuns) GetAttempt() *int32 {
	return s.Attempt
}

func (s *ListPipelineRunsResponseBodyRuns) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListPipelineRunsResponseBodyRuns) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListPipelineRunsResponseBodyRuns) GetFinishTime() *string {
	return s.FinishTime
}

func (s *ListPipelineRunsResponseBodyRuns) GetFromTime() *int64 {
	return s.FromTime
}

func (s *ListPipelineRunsResponseBodyRuns) GetMaxAttempts() *int32 {
	return s.MaxAttempts
}

func (s *ListPipelineRunsResponseBodyRuns) GetNextRetryTime() *string {
	return s.NextRetryTime
}

func (s *ListPipelineRunsResponseBodyRuns) GetPipelineName() *string {
	return s.PipelineName
}

func (s *ListPipelineRunsResponseBodyRuns) GetRunId() *string {
	return s.RunId
}

func (s *ListPipelineRunsResponseBodyRuns) GetStartTime() *string {
	return s.StartTime
}

func (s *ListPipelineRunsResponseBodyRuns) GetStats() map[string]interface{} {
	return s.Stats
}

func (s *ListPipelineRunsResponseBodyRuns) GetStatus() *string {
	return s.Status
}

func (s *ListPipelineRunsResponseBodyRuns) GetToTime() *int64 {
	return s.ToTime
}

func (s *ListPipelineRunsResponseBodyRuns) GetTriggerTime() *string {
	return s.TriggerTime
}

func (s *ListPipelineRunsResponseBodyRuns) GetTriggerType() *string {
	return s.TriggerType
}

func (s *ListPipelineRunsResponseBodyRuns) SetAgentSpaceName(v string) *ListPipelineRunsResponseBodyRuns {
	s.AgentSpaceName = &v
	return s
}

func (s *ListPipelineRunsResponseBodyRuns) SetAttempt(v int32) *ListPipelineRunsResponseBodyRuns {
	s.Attempt = &v
	return s
}

func (s *ListPipelineRunsResponseBodyRuns) SetErrorCode(v string) *ListPipelineRunsResponseBodyRuns {
	s.ErrorCode = &v
	return s
}

func (s *ListPipelineRunsResponseBodyRuns) SetErrorMessage(v string) *ListPipelineRunsResponseBodyRuns {
	s.ErrorMessage = &v
	return s
}

func (s *ListPipelineRunsResponseBodyRuns) SetFinishTime(v string) *ListPipelineRunsResponseBodyRuns {
	s.FinishTime = &v
	return s
}

func (s *ListPipelineRunsResponseBodyRuns) SetFromTime(v int64) *ListPipelineRunsResponseBodyRuns {
	s.FromTime = &v
	return s
}

func (s *ListPipelineRunsResponseBodyRuns) SetMaxAttempts(v int32) *ListPipelineRunsResponseBodyRuns {
	s.MaxAttempts = &v
	return s
}

func (s *ListPipelineRunsResponseBodyRuns) SetNextRetryTime(v string) *ListPipelineRunsResponseBodyRuns {
	s.NextRetryTime = &v
	return s
}

func (s *ListPipelineRunsResponseBodyRuns) SetPipelineName(v string) *ListPipelineRunsResponseBodyRuns {
	s.PipelineName = &v
	return s
}

func (s *ListPipelineRunsResponseBodyRuns) SetRunId(v string) *ListPipelineRunsResponseBodyRuns {
	s.RunId = &v
	return s
}

func (s *ListPipelineRunsResponseBodyRuns) SetStartTime(v string) *ListPipelineRunsResponseBodyRuns {
	s.StartTime = &v
	return s
}

func (s *ListPipelineRunsResponseBodyRuns) SetStats(v map[string]interface{}) *ListPipelineRunsResponseBodyRuns {
	s.Stats = v
	return s
}

func (s *ListPipelineRunsResponseBodyRuns) SetStatus(v string) *ListPipelineRunsResponseBodyRuns {
	s.Status = &v
	return s
}

func (s *ListPipelineRunsResponseBodyRuns) SetToTime(v int64) *ListPipelineRunsResponseBodyRuns {
	s.ToTime = &v
	return s
}

func (s *ListPipelineRunsResponseBodyRuns) SetTriggerTime(v string) *ListPipelineRunsResponseBodyRuns {
	s.TriggerTime = &v
	return s
}

func (s *ListPipelineRunsResponseBodyRuns) SetTriggerType(v string) *ListPipelineRunsResponseBodyRuns {
	s.TriggerType = &v
	return s
}

func (s *ListPipelineRunsResponseBodyRuns) Validate() error {
	return dara.Validate(s)
}
