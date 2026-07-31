// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunSemanticJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *RunSemanticJobResponseBodyData) *RunSemanticJobResponseBody
	GetData() *RunSemanticJobResponseBodyData
	SetRequestId(v string) *RunSemanticJobResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *RunSemanticJobResponseBody
	GetSuccess() *bool
}

type RunSemanticJobResponseBody struct {
	// The run information for this submission. A successful submission does not mean that the semantic model output has been generated. Use the detail operation to confirm the status before downloading results.
	Data *RunSemanticJobResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID. Used for locating logs and troubleshooting issues.
	//
	// example:
	//
	// 676271D6-53B4-57BE-89FA-72F7AE1418DF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RunSemanticJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RunSemanticJobResponseBody) GoString() string {
	return s.String()
}

func (s *RunSemanticJobResponseBody) GetData() *RunSemanticJobResponseBodyData {
	return s.Data
}

func (s *RunSemanticJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RunSemanticJobResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *RunSemanticJobResponseBody) SetData(v *RunSemanticJobResponseBodyData) *RunSemanticJobResponseBody {
	s.Data = v
	return s
}

func (s *RunSemanticJobResponseBody) SetRequestId(v string) *RunSemanticJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *RunSemanticJobResponseBody) SetSuccess(v bool) *RunSemanticJobResponseBody {
	s.Success = &v
	return s
}

func (s *RunSemanticJobResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type RunSemanticJobResponseBodyData struct {
	// The current SQL fragment index returned by the executor in the submission response.
	//
	// example:
	//
	// 0
	CurrentSqlIndex *int32 `json:"CurrentSqlIndex,omitempty" xml:"CurrentSqlIndex,omitempty"`
	// The runtime environment identifier returned by the executor in the submission response.
	//
	// example:
	//
	// PROD
	Env *string `json:"Env,omitempty" xml:"Env,omitempty"`
	// The list of execution type codes returned by the executor in the submission response.
	ExecTypes []*int32 `json:"ExecTypes,omitempty" xml:"ExecTypes,omitempty" type:"Repeated"`
	// The executor job identifier. Pass this value to the ExecutorJobId parameter of GetSemanticJobDetail, GetSemanticJobLog, or KillSemanticJob.
	//
	// example:
	//
	// exec-job-demo
	ExecutorJobId *string `json:"ExecutorJobId,omitempty" xml:"ExecutorJobId,omitempty"`
	// The unique identifier of this run. Pass this value to the JobRunId parameter of DownloadSemanticResults to obtain the output of this run.
	//
	// example:
	//
	// 01H00000000000000000000000
	JobRunId *string `json:"JobRunId,omitempty" xml:"JobRunId,omitempty"`
	// The list of status codes returned by the executor in the submission response. The status at the submission stage does not indicate that the results are complete.
	Statuses []*int32 `json:"Statuses,omitempty" xml:"Statuses,omitempty" type:"Repeated"`
}

func (s RunSemanticJobResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s RunSemanticJobResponseBodyData) GoString() string {
	return s.String()
}

func (s *RunSemanticJobResponseBodyData) GetCurrentSqlIndex() *int32 {
	return s.CurrentSqlIndex
}

func (s *RunSemanticJobResponseBodyData) GetEnv() *string {
	return s.Env
}

func (s *RunSemanticJobResponseBodyData) GetExecTypes() []*int32 {
	return s.ExecTypes
}

func (s *RunSemanticJobResponseBodyData) GetExecutorJobId() *string {
	return s.ExecutorJobId
}

func (s *RunSemanticJobResponseBodyData) GetJobRunId() *string {
	return s.JobRunId
}

func (s *RunSemanticJobResponseBodyData) GetStatuses() []*int32 {
	return s.Statuses
}

func (s *RunSemanticJobResponseBodyData) SetCurrentSqlIndex(v int32) *RunSemanticJobResponseBodyData {
	s.CurrentSqlIndex = &v
	return s
}

func (s *RunSemanticJobResponseBodyData) SetEnv(v string) *RunSemanticJobResponseBodyData {
	s.Env = &v
	return s
}

func (s *RunSemanticJobResponseBodyData) SetExecTypes(v []*int32) *RunSemanticJobResponseBodyData {
	s.ExecTypes = v
	return s
}

func (s *RunSemanticJobResponseBodyData) SetExecutorJobId(v string) *RunSemanticJobResponseBodyData {
	s.ExecutorJobId = &v
	return s
}

func (s *RunSemanticJobResponseBodyData) SetJobRunId(v string) *RunSemanticJobResponseBodyData {
	s.JobRunId = &v
	return s
}

func (s *RunSemanticJobResponseBodyData) SetStatuses(v []*int32) *RunSemanticJobResponseBodyData {
	s.Statuses = v
	return s
}

func (s *RunSemanticJobResponseBodyData) Validate() error {
	return dara.Validate(s)
}
