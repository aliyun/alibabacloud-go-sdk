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
	// The details of the job.
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
	// The code type of the job. Valid values:
	//
	// 	- SQL
	//
	// 	- JAR
	//
	// 	- PYTHON
	//
	// example:
	//
	// SQL
	CodeType *string `json:"codeType,omitempty" xml:"codeType,omitempty"`
	// The configurations of the Spark jobs.
	ConfigurationOverrides *GetJobRunResponseBodyJobRunConfigurationOverrides `json:"configurationOverrides,omitempty" xml:"configurationOverrides,omitempty" type:"Struct"`
	// The version of the Spark engine.
	//
	// example:
	//
	// esr-4.0.0 (Spark 3.5.2, Scala 2.12)
	DisplayReleaseVersion *string `json:"displayReleaseVersion,omitempty" xml:"displayReleaseVersion,omitempty"`
	// The end time of the job.
	//
	// example:
	//
	// 1684119314000
	EndTime *int64 `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// The environment ID.
	//
	// example:
	//
	// env-cpv569tlhtgndjl8****
	EnvironmentId *string `json:"environmentId,omitempty" xml:"environmentId,omitempty"`
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
	// false
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
	// The job name.
	//
	// example:
	//
	// jobName
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// http://workflow-ide-cn-hangzhou.oss-cn-hangzhou.aliyuncs.com/spark-notebook-output/w-xxxxxxxxx/xxxxxxx
	NotebookAccessUrl *string `json:"notebookAccessUrl,omitempty" xml:"notebookAccessUrl,omitempty"`
	// The version of the Spark engine on which the job runs.
	//
	// example:
	//
	// esr-3.3.1
	ReleaseVersion *string `json:"releaseVersion,omitempty" xml:"releaseVersion,omitempty"`
	// The ID of the user who created the job.
	//
	// example:
	//
	// 1509789347011222
	ResourceOwnerId *string `json:"resourceOwnerId,omitempty" xml:"resourceOwnerId,omitempty"`
	// The name of the queue on which the job runs.
	//
	// example:
	//
	// root_queue
	ResourceQueueId *string `json:"resourceQueueId,omitempty" xml:"resourceQueueId,omitempty"`
	// The job state.
	//
	// example:
	//
	// Running
	State *string `json:"state,omitempty" xml:"state,omitempty"`
	// The reason of the job status change.
	StateChangeReason *GetJobRunResponseBodyJobRunStateChangeReason `json:"stateChangeReason,omitempty" xml:"stateChangeReason,omitempty" type:"Struct"`
	// The time when the job was submitted.
	//
	// example:
	//
	// 1684119314000
	SubmitTime *int64 `json:"submitTime,omitempty" xml:"submitTime,omitempty"`
	// The tags of the job.
	Tags []*Tag `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
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
	return nil
}

type GetJobRunResponseBodyJobRunConfigurationOverrides struct {
	// The configurations.
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
