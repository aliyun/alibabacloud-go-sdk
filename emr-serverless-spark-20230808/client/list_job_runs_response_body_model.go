// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListJobRunsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetJobRuns(v []*ListJobRunsResponseBodyJobRuns) *ListJobRunsResponseBody
	GetJobRuns() []*ListJobRunsResponseBodyJobRuns
	SetMaxResults(v int32) *ListJobRunsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListJobRunsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListJobRunsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListJobRunsResponseBody
	GetTotalCount() *int32
}

type ListJobRunsResponseBody struct {
	// The Spark jobs.
	JobRuns []*ListJobRunsResponseBodyJobRuns `json:"jobRuns,omitempty" xml:"jobRuns,omitempty" type:"Repeated"`
	// The maximum number of entries returned.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// A pagination token.
	//
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C89568980
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C8944****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 200
	TotalCount *int32 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListJobRunsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListJobRunsResponseBody) GoString() string {
	return s.String()
}

func (s *ListJobRunsResponseBody) GetJobRuns() []*ListJobRunsResponseBodyJobRuns {
	return s.JobRuns
}

func (s *ListJobRunsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListJobRunsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListJobRunsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListJobRunsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListJobRunsResponseBody) SetJobRuns(v []*ListJobRunsResponseBodyJobRuns) *ListJobRunsResponseBody {
	s.JobRuns = v
	return s
}

func (s *ListJobRunsResponseBody) SetMaxResults(v int32) *ListJobRunsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListJobRunsResponseBody) SetNextToken(v string) *ListJobRunsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListJobRunsResponseBody) SetRequestId(v string) *ListJobRunsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListJobRunsResponseBody) SetTotalCount(v int32) *ListJobRunsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListJobRunsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListJobRunsResponseBodyJobRuns struct {
	// The code type of the job. Valid values:
	//
	// SQL
	//
	// JAR
	//
	// PYTHON
	//
	// example:
	//
	// SQL
	CodeType *string `json:"codeType,omitempty" xml:"codeType,omitempty"`
	// The advanced configurations of Spark.
	ConfigurationOverrides *ListJobRunsResponseBodyJobRunsConfigurationOverrides `json:"configurationOverrides,omitempty" xml:"configurationOverrides,omitempty" type:"Struct"`
	// The ID of the user who created the job.
	//
	// example:
	//
	// 1509789347011222
	Creator *string `json:"creator,omitempty" xml:"creator,omitempty"`
	// The number of CUs consumed during a specified cycle of a task. The value is an estimated value. Refer to your Alibaba Cloud bill for the actual number of consumed CUs.
	//
	// example:
	//
	// 2.059
	CuHours *float64 `json:"cuHours,omitempty" xml:"cuHours,omitempty"`
	// The version of Spark on which the jobs run.
	//
	// example:
	//
	// esr-3.0.0 (Spark 3.4.3, Scala 2.12)
	DisplayReleaseVersion *string `json:"displayReleaseVersion,omitempty" xml:"displayReleaseVersion,omitempty"`
	// The end time of the job.
	//
	// example:
	//
	// 1684119314000
	EndTime *int64 `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// The timeout period of the job.
	//
	// example:
	//
	// 3600
	ExecutionTimeoutSeconds *int32 `json:"executionTimeoutSeconds,omitempty" xml:"executionTimeoutSeconds,omitempty"`
	// Indicates whether the Fusion engine is used for acceleration.
	//
	// example:
	//
	// true
	Fusion *bool `json:"fusion,omitempty" xml:"fusion,omitempty"`
	// The information about Spark Driver.
	JobDriver *JobDriver `json:"jobDriver,omitempty" xml:"jobDriver,omitempty"`
	// The job ID.
	//
	// example:
	//
	// jr-231231
	JobRunId *string `json:"jobRunId,omitempty" xml:"jobRunId,omitempty"`
	// The path where the operational logs are stored.
	Log *RunLog `json:"log,omitempty" xml:"log,omitempty"`
	// The total amount of memory allocated to the job multiplied by the running duration (seconds).
	//
	// example:
	//
	// 33030784
	MbSeconds *int64 `json:"mbSeconds,omitempty" xml:"mbSeconds,omitempty"`
	// The job name.
	//
	// example:
	//
	// jobName
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The version of Spark on which the jobs run.
	//
	// example:
	//
	// esr-native-3.4.0
	ReleaseVersion *string `json:"releaseVersion,omitempty" xml:"releaseVersion,omitempty"`
	// The job state.
	//
	// example:
	//
	// Running
	State *string `json:"state,omitempty" xml:"state,omitempty"`
	// The reason of the job status change.
	StateChangeReason *ListJobRunsResponseBodyJobRunsStateChangeReason `json:"stateChangeReason,omitempty" xml:"stateChangeReason,omitempty" type:"Struct"`
	// The time when the job was submitted.
	//
	// example:
	//
	// 1684119314000
	SubmitTime *int64 `json:"submitTime,omitempty" xml:"submitTime,omitempty"`
	// The tags of the job.
	Tags []*Tag `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
	// The total number of CPU cores allocated to the job multiplied by the running duration (seconds).
	//
	// example:
	//
	// 8236
	VcoreSeconds *int64 `json:"vcoreSeconds,omitempty" xml:"vcoreSeconds,omitempty"`
	// The web UI of the job.
	//
	// example:
	//
	// http://spark-ui
	WebUI *string `json:"webUI,omitempty" xml:"webUI,omitempty"`
	// The workspace ID.
	//
	// example:
	//
	// w-1234abcd
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s ListJobRunsResponseBodyJobRuns) String() string {
	return dara.Prettify(s)
}

func (s ListJobRunsResponseBodyJobRuns) GoString() string {
	return s.String()
}

func (s *ListJobRunsResponseBodyJobRuns) GetCodeType() *string {
	return s.CodeType
}

func (s *ListJobRunsResponseBodyJobRuns) GetConfigurationOverrides() *ListJobRunsResponseBodyJobRunsConfigurationOverrides {
	return s.ConfigurationOverrides
}

func (s *ListJobRunsResponseBodyJobRuns) GetCreator() *string {
	return s.Creator
}

func (s *ListJobRunsResponseBodyJobRuns) GetCuHours() *float64 {
	return s.CuHours
}

func (s *ListJobRunsResponseBodyJobRuns) GetDisplayReleaseVersion() *string {
	return s.DisplayReleaseVersion
}

func (s *ListJobRunsResponseBodyJobRuns) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ListJobRunsResponseBodyJobRuns) GetExecutionTimeoutSeconds() *int32 {
	return s.ExecutionTimeoutSeconds
}

func (s *ListJobRunsResponseBodyJobRuns) GetFusion() *bool {
	return s.Fusion
}

func (s *ListJobRunsResponseBodyJobRuns) GetJobDriver() *JobDriver {
	return s.JobDriver
}

func (s *ListJobRunsResponseBodyJobRuns) GetJobRunId() *string {
	return s.JobRunId
}

func (s *ListJobRunsResponseBodyJobRuns) GetLog() *RunLog {
	return s.Log
}

func (s *ListJobRunsResponseBodyJobRuns) GetMbSeconds() *int64 {
	return s.MbSeconds
}

func (s *ListJobRunsResponseBodyJobRuns) GetName() *string {
	return s.Name
}

func (s *ListJobRunsResponseBodyJobRuns) GetReleaseVersion() *string {
	return s.ReleaseVersion
}

func (s *ListJobRunsResponseBodyJobRuns) GetState() *string {
	return s.State
}

func (s *ListJobRunsResponseBodyJobRuns) GetStateChangeReason() *ListJobRunsResponseBodyJobRunsStateChangeReason {
	return s.StateChangeReason
}

func (s *ListJobRunsResponseBodyJobRuns) GetSubmitTime() *int64 {
	return s.SubmitTime
}

func (s *ListJobRunsResponseBodyJobRuns) GetTags() []*Tag {
	return s.Tags
}

func (s *ListJobRunsResponseBodyJobRuns) GetVcoreSeconds() *int64 {
	return s.VcoreSeconds
}

func (s *ListJobRunsResponseBodyJobRuns) GetWebUI() *string {
	return s.WebUI
}

func (s *ListJobRunsResponseBodyJobRuns) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListJobRunsResponseBodyJobRuns) SetCodeType(v string) *ListJobRunsResponseBodyJobRuns {
	s.CodeType = &v
	return s
}

func (s *ListJobRunsResponseBodyJobRuns) SetConfigurationOverrides(v *ListJobRunsResponseBodyJobRunsConfigurationOverrides) *ListJobRunsResponseBodyJobRuns {
	s.ConfigurationOverrides = v
	return s
}

func (s *ListJobRunsResponseBodyJobRuns) SetCreator(v string) *ListJobRunsResponseBodyJobRuns {
	s.Creator = &v
	return s
}

func (s *ListJobRunsResponseBodyJobRuns) SetCuHours(v float64) *ListJobRunsResponseBodyJobRuns {
	s.CuHours = &v
	return s
}

func (s *ListJobRunsResponseBodyJobRuns) SetDisplayReleaseVersion(v string) *ListJobRunsResponseBodyJobRuns {
	s.DisplayReleaseVersion = &v
	return s
}

func (s *ListJobRunsResponseBodyJobRuns) SetEndTime(v int64) *ListJobRunsResponseBodyJobRuns {
	s.EndTime = &v
	return s
}

func (s *ListJobRunsResponseBodyJobRuns) SetExecutionTimeoutSeconds(v int32) *ListJobRunsResponseBodyJobRuns {
	s.ExecutionTimeoutSeconds = &v
	return s
}

func (s *ListJobRunsResponseBodyJobRuns) SetFusion(v bool) *ListJobRunsResponseBodyJobRuns {
	s.Fusion = &v
	return s
}

func (s *ListJobRunsResponseBodyJobRuns) SetJobDriver(v *JobDriver) *ListJobRunsResponseBodyJobRuns {
	s.JobDriver = v
	return s
}

func (s *ListJobRunsResponseBodyJobRuns) SetJobRunId(v string) *ListJobRunsResponseBodyJobRuns {
	s.JobRunId = &v
	return s
}

func (s *ListJobRunsResponseBodyJobRuns) SetLog(v *RunLog) *ListJobRunsResponseBodyJobRuns {
	s.Log = v
	return s
}

func (s *ListJobRunsResponseBodyJobRuns) SetMbSeconds(v int64) *ListJobRunsResponseBodyJobRuns {
	s.MbSeconds = &v
	return s
}

func (s *ListJobRunsResponseBodyJobRuns) SetName(v string) *ListJobRunsResponseBodyJobRuns {
	s.Name = &v
	return s
}

func (s *ListJobRunsResponseBodyJobRuns) SetReleaseVersion(v string) *ListJobRunsResponseBodyJobRuns {
	s.ReleaseVersion = &v
	return s
}

func (s *ListJobRunsResponseBodyJobRuns) SetState(v string) *ListJobRunsResponseBodyJobRuns {
	s.State = &v
	return s
}

func (s *ListJobRunsResponseBodyJobRuns) SetStateChangeReason(v *ListJobRunsResponseBodyJobRunsStateChangeReason) *ListJobRunsResponseBodyJobRuns {
	s.StateChangeReason = v
	return s
}

func (s *ListJobRunsResponseBodyJobRuns) SetSubmitTime(v int64) *ListJobRunsResponseBodyJobRuns {
	s.SubmitTime = &v
	return s
}

func (s *ListJobRunsResponseBodyJobRuns) SetTags(v []*Tag) *ListJobRunsResponseBodyJobRuns {
	s.Tags = v
	return s
}

func (s *ListJobRunsResponseBodyJobRuns) SetVcoreSeconds(v int64) *ListJobRunsResponseBodyJobRuns {
	s.VcoreSeconds = &v
	return s
}

func (s *ListJobRunsResponseBodyJobRuns) SetWebUI(v string) *ListJobRunsResponseBodyJobRuns {
	s.WebUI = &v
	return s
}

func (s *ListJobRunsResponseBodyJobRuns) SetWorkspaceId(v string) *ListJobRunsResponseBodyJobRuns {
	s.WorkspaceId = &v
	return s
}

func (s *ListJobRunsResponseBodyJobRuns) Validate() error {
	return dara.Validate(s)
}

type ListJobRunsResponseBodyJobRunsConfigurationOverrides struct {
	// The SparkConf objects.
	Configurations []*Configuration `json:"configurations,omitempty" xml:"configurations,omitempty" type:"Repeated"`
}

func (s ListJobRunsResponseBodyJobRunsConfigurationOverrides) String() string {
	return dara.Prettify(s)
}

func (s ListJobRunsResponseBodyJobRunsConfigurationOverrides) GoString() string {
	return s.String()
}

func (s *ListJobRunsResponseBodyJobRunsConfigurationOverrides) GetConfigurations() []*Configuration {
	return s.Configurations
}

func (s *ListJobRunsResponseBodyJobRunsConfigurationOverrides) SetConfigurations(v []*Configuration) *ListJobRunsResponseBodyJobRunsConfigurationOverrides {
	s.Configurations = v
	return s
}

func (s *ListJobRunsResponseBodyJobRunsConfigurationOverrides) Validate() error {
	return dara.Validate(s)
}

type ListJobRunsResponseBodyJobRunsStateChangeReason struct {
	// The error code.
	//
	// example:
	//
	// 0
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The error message.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
}

func (s ListJobRunsResponseBodyJobRunsStateChangeReason) String() string {
	return dara.Prettify(s)
}

func (s ListJobRunsResponseBodyJobRunsStateChangeReason) GoString() string {
	return s.String()
}

func (s *ListJobRunsResponseBodyJobRunsStateChangeReason) GetCode() *string {
	return s.Code
}

func (s *ListJobRunsResponseBodyJobRunsStateChangeReason) GetMessage() *string {
	return s.Message
}

func (s *ListJobRunsResponseBodyJobRunsStateChangeReason) SetCode(v string) *ListJobRunsResponseBodyJobRunsStateChangeReason {
	s.Code = &v
	return s
}

func (s *ListJobRunsResponseBodyJobRunsStateChangeReason) SetMessage(v string) *ListJobRunsResponseBodyJobRunsStateChangeReason {
	s.Message = &v
	return s
}

func (s *ListJobRunsResponseBodyJobRunsStateChangeReason) Validate() error {
	return dara.Validate(s)
}
