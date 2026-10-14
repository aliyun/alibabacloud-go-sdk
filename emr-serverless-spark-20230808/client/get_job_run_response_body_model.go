// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetJobRunResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetJobRun(v *GetJobRunResponseBodyJobRun) *GetJobRunResponseBody
	GetJobRun() *GetJobRunResponseBodyJobRun
	SetRequestId(v string) *GetJobRunResponseBody
	GetRequestId() *string
}

type GetJobRunResponseBody struct {
	// The job run details.
	JobRun *GetJobRunResponseBodyJobRun `json:"jobRun,omitempty" xml:"jobRun,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C8944****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetJobRunResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetJobRunResponseBody) GoString() string {
	return s.String()
}

func (s *GetJobRunResponseBody) GetJobRun() *GetJobRunResponseBodyJobRun {
	return s.JobRun
}

func (s *GetJobRunResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetJobRunResponseBody) SetJobRun(v *GetJobRunResponseBodyJobRun) *GetJobRunResponseBody {
	s.JobRun = v
	return s
}

func (s *GetJobRunResponseBody) SetRequestId(v string) *GetJobRunResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetJobRunResponseBody) Validate() error {
	if s.JobRun != nil {
		if err := s.JobRun.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetJobRunResponseBodyJobRun struct {
	// The job code type. Valid values:
	//
	// - SQL
	//
	// - JAR
	//
	// - PYTHON
	//
	// example:
	//
	// SQL
	CodeType *string `json:"codeType,omitempty" xml:"codeType,omitempty"`
	// The Spark job configuration.
	ConfigurationOverrides *GetJobRunResponseBodyJobRunConfigurationOverrides `json:"configurationOverrides,omitempty" xml:"configurationOverrides,omitempty" type:"Struct"`
	// The version displayed in the console.
	//
	// example:
	//
	// esr-4.0.0 (Spark 3.5.2, Scala 2.12)
	DisplayReleaseVersion *string `json:"displayReleaseVersion,omitempty" xml:"displayReleaseVersion,omitempty"`
	// The job end time. The value is a Unix timestamp in milliseconds.
	//
	// example:
	//
	// 1684122914000
	EndTime *int64 `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// The environment ID.
	//
	// example:
	//
	// env-cpv569tlhtgndjl8****
	EnvironmentId *string `json:"environmentId,omitempty" xml:"environmentId,omitempty"`
	// The execution timeout period.
	//
	// example:
	//
	// 3600
	ExecutionTimeoutSeconds *int32 `json:"executionTimeoutSeconds,omitempty" xml:"executionTimeoutSeconds,omitempty"`
	// Specifies whether the Fusion engine acceleration is enabled.
	//
	// example:
	//
	// false
	Fusion *bool `json:"fusion,omitempty" xml:"fusion,omitempty"`
	// The Spark Driver information.
	JobDriver *JobDriver `json:"jobDriver,omitempty" xml:"jobDriver,omitempty"`
	// The job run ID.
	//
	// example:
	//
	// jr-93d98d2f7061****
	JobRunId *string `json:"jobRunId,omitempty" xml:"jobRunId,omitempty"`
	// The path of the run log.
	Log *RunLog `json:"log,omitempty" xml:"log,omitempty"`
	// The job name.
	//
	// example:
	//
	// jobName
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The download URL of the NOTEBOOK file. This parameter is returned only if the job type is NOTEBOOK.
	//
	// example:
	//
	// http://workflow-ide-cn-hangzhou.oss-cn-hangzhou.aliyuncs.com/spark-notebook-output/w-xxxxxxxxx/xxxxxxx
	NotebookAccessUrl *string `json:"notebookAccessUrl,omitempty" xml:"notebookAccessUrl,omitempty"`
	// The job priority.
	//
	// example:
	//
	// 5
	Priority *string `json:"priority,omitempty" xml:"priority,omitempty"`
	// The Spark DPI engine version used to run the node.
	//
	// example:
	//
	// esr-3.3.1
	ReleaseVersion *string `json:"releaseVersion,omitempty" xml:"releaseVersion,omitempty"`
	// The UID of the user who created the job.
	//
	// example:
	//
	// 150978934701****
	ResourceOwnerId *string `json:"resourceOwnerId,omitempty" xml:"resourceOwnerId,omitempty"`
	// The name of the queue used to run the job.
	//
	// example:
	//
	// root_queue
	ResourceQueueId *string `json:"resourceQueueId,omitempty" xml:"resourceQueueId,omitempty"`
	// The job status.
	//
	// example:
	//
	// Running
	State *string `json:"state,omitempty" xml:"state,omitempty"`
	// The reason for the status change.
	StateChangeReason *GetJobRunResponseBodyJobRunStateChangeReason `json:"stateChangeReason,omitempty" xml:"stateChangeReason,omitempty" type:"Struct"`
	// The job submission time. The value is a Unix timestamp in milliseconds.
	//
	// example:
	//
	// 1684119314000
	SubmitTime *int64 `json:"submitTime,omitempty" xml:"submitTime,omitempty"`
	// The tag list.
	Tags []*Tag `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
	// The details of token consumption.
	TokenDetails []*GetJobRunResponseBodyJobRunTokenDetails `json:"tokenDetails,omitempty" xml:"tokenDetails,omitempty" type:"Repeated"`
	// The total number of tokens consumed.
	//
	// example:
	//
	// 10000
	TotalTokens *int64 `json:"totalTokens,omitempty" xml:"totalTokens,omitempty"`
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
	// w-d2d82aa09155****
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s GetJobRunResponseBodyJobRun) String() string {
	return dara.Prettify(s)
}

func (s GetJobRunResponseBodyJobRun) GoString() string {
	return s.String()
}

func (s *GetJobRunResponseBodyJobRun) GetCodeType() *string {
	return s.CodeType
}

func (s *GetJobRunResponseBodyJobRun) GetConfigurationOverrides() *GetJobRunResponseBodyJobRunConfigurationOverrides {
	return s.ConfigurationOverrides
}

func (s *GetJobRunResponseBodyJobRun) GetDisplayReleaseVersion() *string {
	return s.DisplayReleaseVersion
}

func (s *GetJobRunResponseBodyJobRun) GetEndTime() *int64 {
	return s.EndTime
}

func (s *GetJobRunResponseBodyJobRun) GetEnvironmentId() *string {
	return s.EnvironmentId
}

func (s *GetJobRunResponseBodyJobRun) GetExecutionTimeoutSeconds() *int32 {
	return s.ExecutionTimeoutSeconds
}

func (s *GetJobRunResponseBodyJobRun) GetFusion() *bool {
	return s.Fusion
}

func (s *GetJobRunResponseBodyJobRun) GetJobDriver() *JobDriver {
	return s.JobDriver
}

func (s *GetJobRunResponseBodyJobRun) GetJobRunId() *string {
	return s.JobRunId
}

func (s *GetJobRunResponseBodyJobRun) GetLog() *RunLog {
	return s.Log
}

func (s *GetJobRunResponseBodyJobRun) GetName() *string {
	return s.Name
}

func (s *GetJobRunResponseBodyJobRun) GetNotebookAccessUrl() *string {
	return s.NotebookAccessUrl
}

func (s *GetJobRunResponseBodyJobRun) GetPriority() *string {
	return s.Priority
}

func (s *GetJobRunResponseBodyJobRun) GetReleaseVersion() *string {
	return s.ReleaseVersion
}

func (s *GetJobRunResponseBodyJobRun) GetResourceOwnerId() *string {
	return s.ResourceOwnerId
}

func (s *GetJobRunResponseBodyJobRun) GetResourceQueueId() *string {
	return s.ResourceQueueId
}

func (s *GetJobRunResponseBodyJobRun) GetState() *string {
	return s.State
}

func (s *GetJobRunResponseBodyJobRun) GetStateChangeReason() *GetJobRunResponseBodyJobRunStateChangeReason {
	return s.StateChangeReason
}

func (s *GetJobRunResponseBodyJobRun) GetSubmitTime() *int64 {
	return s.SubmitTime
}

func (s *GetJobRunResponseBodyJobRun) GetTags() []*Tag {
	return s.Tags
}

func (s *GetJobRunResponseBodyJobRun) GetTokenDetails() []*GetJobRunResponseBodyJobRunTokenDetails {
	return s.TokenDetails
}

func (s *GetJobRunResponseBodyJobRun) GetTotalTokens() *int64 {
	return s.TotalTokens
}

func (s *GetJobRunResponseBodyJobRun) GetWebUI() *string {
	return s.WebUI
}

func (s *GetJobRunResponseBodyJobRun) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *GetJobRunResponseBodyJobRun) SetCodeType(v string) *GetJobRunResponseBodyJobRun {
	s.CodeType = &v
	return s
}

func (s *GetJobRunResponseBodyJobRun) SetConfigurationOverrides(v *GetJobRunResponseBodyJobRunConfigurationOverrides) *GetJobRunResponseBodyJobRun {
	s.ConfigurationOverrides = v
	return s
}

func (s *GetJobRunResponseBodyJobRun) SetDisplayReleaseVersion(v string) *GetJobRunResponseBodyJobRun {
	s.DisplayReleaseVersion = &v
	return s
}

func (s *GetJobRunResponseBodyJobRun) SetEndTime(v int64) *GetJobRunResponseBodyJobRun {
	s.EndTime = &v
	return s
}

func (s *GetJobRunResponseBodyJobRun) SetEnvironmentId(v string) *GetJobRunResponseBodyJobRun {
	s.EnvironmentId = &v
	return s
}

func (s *GetJobRunResponseBodyJobRun) SetExecutionTimeoutSeconds(v int32) *GetJobRunResponseBodyJobRun {
	s.ExecutionTimeoutSeconds = &v
	return s
}

func (s *GetJobRunResponseBodyJobRun) SetFusion(v bool) *GetJobRunResponseBodyJobRun {
	s.Fusion = &v
	return s
}

func (s *GetJobRunResponseBodyJobRun) SetJobDriver(v *JobDriver) *GetJobRunResponseBodyJobRun {
	s.JobDriver = v
	return s
}

func (s *GetJobRunResponseBodyJobRun) SetJobRunId(v string) *GetJobRunResponseBodyJobRun {
	s.JobRunId = &v
	return s
}

func (s *GetJobRunResponseBodyJobRun) SetLog(v *RunLog) *GetJobRunResponseBodyJobRun {
	s.Log = v
	return s
}

func (s *GetJobRunResponseBodyJobRun) SetName(v string) *GetJobRunResponseBodyJobRun {
	s.Name = &v
	return s
}

func (s *GetJobRunResponseBodyJobRun) SetNotebookAccessUrl(v string) *GetJobRunResponseBodyJobRun {
	s.NotebookAccessUrl = &v
	return s
}

func (s *GetJobRunResponseBodyJobRun) SetPriority(v string) *GetJobRunResponseBodyJobRun {
	s.Priority = &v
	return s
}

func (s *GetJobRunResponseBodyJobRun) SetReleaseVersion(v string) *GetJobRunResponseBodyJobRun {
	s.ReleaseVersion = &v
	return s
}

func (s *GetJobRunResponseBodyJobRun) SetResourceOwnerId(v string) *GetJobRunResponseBodyJobRun {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetJobRunResponseBodyJobRun) SetResourceQueueId(v string) *GetJobRunResponseBodyJobRun {
	s.ResourceQueueId = &v
	return s
}

func (s *GetJobRunResponseBodyJobRun) SetState(v string) *GetJobRunResponseBodyJobRun {
	s.State = &v
	return s
}

func (s *GetJobRunResponseBodyJobRun) SetStateChangeReason(v *GetJobRunResponseBodyJobRunStateChangeReason) *GetJobRunResponseBodyJobRun {
	s.StateChangeReason = v
	return s
}

func (s *GetJobRunResponseBodyJobRun) SetSubmitTime(v int64) *GetJobRunResponseBodyJobRun {
	s.SubmitTime = &v
	return s
}

func (s *GetJobRunResponseBodyJobRun) SetTags(v []*Tag) *GetJobRunResponseBodyJobRun {
	s.Tags = v
	return s
}

func (s *GetJobRunResponseBodyJobRun) SetTokenDetails(v []*GetJobRunResponseBodyJobRunTokenDetails) *GetJobRunResponseBodyJobRun {
	s.TokenDetails = v
	return s
}

func (s *GetJobRunResponseBodyJobRun) SetTotalTokens(v int64) *GetJobRunResponseBodyJobRun {
	s.TotalTokens = &v
	return s
}

func (s *GetJobRunResponseBodyJobRun) SetWebUI(v string) *GetJobRunResponseBodyJobRun {
	s.WebUI = &v
	return s
}

func (s *GetJobRunResponseBodyJobRun) SetWorkspaceId(v string) *GetJobRunResponseBodyJobRun {
	s.WorkspaceId = &v
	return s
}

func (s *GetJobRunResponseBodyJobRun) Validate() error {
	if s.ConfigurationOverrides != nil {
		if err := s.ConfigurationOverrides.Validate(); err != nil {
			return err
		}
	}
	if s.JobDriver != nil {
		if err := s.JobDriver.Validate(); err != nil {
			return err
		}
	}
	if s.Log != nil {
		if err := s.Log.Validate(); err != nil {
			return err
		}
	}
	if s.StateChangeReason != nil {
		if err := s.StateChangeReason.Validate(); err != nil {
			return err
		}
	}
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.TokenDetails != nil {
		for _, item := range s.TokenDetails {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetJobRunResponseBodyJobRunConfigurationOverrides struct {
	// The configuration list.
	Configurations []*Configuration `json:"configurations,omitempty" xml:"configurations,omitempty" type:"Repeated"`
}

func (s GetJobRunResponseBodyJobRunConfigurationOverrides) String() string {
	return dara.Prettify(s)
}

func (s GetJobRunResponseBodyJobRunConfigurationOverrides) GoString() string {
	return s.String()
}

func (s *GetJobRunResponseBodyJobRunConfigurationOverrides) GetConfigurations() []*Configuration {
	return s.Configurations
}

func (s *GetJobRunResponseBodyJobRunConfigurationOverrides) SetConfigurations(v []*Configuration) *GetJobRunResponseBodyJobRunConfigurationOverrides {
	s.Configurations = v
	return s
}

func (s *GetJobRunResponseBodyJobRunConfigurationOverrides) Validate() error {
	if s.Configurations != nil {
		for _, item := range s.Configurations {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetJobRunResponseBodyJobRunStateChangeReason struct {
	// The error code.
	//
	// example:
	//
	// ERR-100000
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The error message.
	//
	// example:
	//
	// connection refused
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
}

func (s GetJobRunResponseBodyJobRunStateChangeReason) String() string {
	return dara.Prettify(s)
}

func (s GetJobRunResponseBodyJobRunStateChangeReason) GoString() string {
	return s.String()
}

func (s *GetJobRunResponseBodyJobRunStateChangeReason) GetCode() *string {
	return s.Code
}

func (s *GetJobRunResponseBodyJobRunStateChangeReason) GetMessage() *string {
	return s.Message
}

func (s *GetJobRunResponseBodyJobRunStateChangeReason) SetCode(v string) *GetJobRunResponseBodyJobRunStateChangeReason {
	s.Code = &v
	return s
}

func (s *GetJobRunResponseBodyJobRunStateChangeReason) SetMessage(v string) *GetJobRunResponseBodyJobRunStateChangeReason {
	s.Message = &v
	return s
}

func (s *GetJobRunResponseBodyJobRunStateChangeReason) Validate() error {
	return dara.Validate(s)
}

type GetJobRunResponseBodyJobRunTokenDetails struct {
	// Specifies whether thinking mode is enabled.
	//
	// example:
	//
	// true
	EnableThinking *string `json:"enableThinking,omitempty" xml:"enableThinking,omitempty"`
	// The input modality.
	//
	// example:
	//
	// -
	InputModality *string `json:"inputModality,omitempty" xml:"inputModality,omitempty"`
	// The model name.
	//
	// example:
	//
	// qwen-plus
	ModelName *string `json:"modelName,omitempty" xml:"modelName,omitempty"`
	// The input token range.
	//
	// example:
	//
	// 0<Token≤256K
	TokenRange *string `json:"tokenRange,omitempty" xml:"tokenRange,omitempty"`
	// The token source. Valid values: input and output.
	//
	// example:
	//
	// input
	TokenSource *string `json:"tokenSource,omitempty" xml:"tokenSource,omitempty"`
	// The number of tokens consumed.
	//
	// example:
	//
	// 10000
	TokenUsed *int64 `json:"tokenUsed,omitempty" xml:"tokenUsed,omitempty"`
}

func (s GetJobRunResponseBodyJobRunTokenDetails) String() string {
	return dara.Prettify(s)
}

func (s GetJobRunResponseBodyJobRunTokenDetails) GoString() string {
	return s.String()
}

func (s *GetJobRunResponseBodyJobRunTokenDetails) GetEnableThinking() *string {
	return s.EnableThinking
}

func (s *GetJobRunResponseBodyJobRunTokenDetails) GetInputModality() *string {
	return s.InputModality
}

func (s *GetJobRunResponseBodyJobRunTokenDetails) GetModelName() *string {
	return s.ModelName
}

func (s *GetJobRunResponseBodyJobRunTokenDetails) GetTokenRange() *string {
	return s.TokenRange
}

func (s *GetJobRunResponseBodyJobRunTokenDetails) GetTokenSource() *string {
	return s.TokenSource
}

func (s *GetJobRunResponseBodyJobRunTokenDetails) GetTokenUsed() *int64 {
	return s.TokenUsed
}

func (s *GetJobRunResponseBodyJobRunTokenDetails) SetEnableThinking(v string) *GetJobRunResponseBodyJobRunTokenDetails {
	s.EnableThinking = &v
	return s
}

func (s *GetJobRunResponseBodyJobRunTokenDetails) SetInputModality(v string) *GetJobRunResponseBodyJobRunTokenDetails {
	s.InputModality = &v
	return s
}

func (s *GetJobRunResponseBodyJobRunTokenDetails) SetModelName(v string) *GetJobRunResponseBodyJobRunTokenDetails {
	s.ModelName = &v
	return s
}

func (s *GetJobRunResponseBodyJobRunTokenDetails) SetTokenRange(v string) *GetJobRunResponseBodyJobRunTokenDetails {
	s.TokenRange = &v
	return s
}

func (s *GetJobRunResponseBodyJobRunTokenDetails) SetTokenSource(v string) *GetJobRunResponseBodyJobRunTokenDetails {
	s.TokenSource = &v
	return s
}

func (s *GetJobRunResponseBodyJobRunTokenDetails) SetTokenUsed(v int64) *GetJobRunResponseBodyJobRunTokenDetails {
	s.TokenUsed = &v
	return s
}

func (s *GetJobRunResponseBodyJobRunTokenDetails) Validate() error {
	return dara.Validate(s)
}
