// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeJobGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeJobGroupResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *DescribeJobGroupResponseBody
	GetHttpStatusCode() *int32
	SetJobGroup(v *DescribeJobGroupResponseBodyJobGroup) *DescribeJobGroupResponseBody
	GetJobGroup() *DescribeJobGroupResponseBodyJobGroup
	SetMessage(v string) *DescribeJobGroupResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeJobGroupResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeJobGroupResponseBody
	GetSuccess() *bool
}

type DescribeJobGroupResponseBody struct {
	// The API status code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The job group information.
	JobGroup *DescribeJobGroupResponseBodyJobGroup `json:"JobGroup,omitempty" xml:"JobGroup,omitempty" type:"Struct"`
	// The API message.
	//
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 254EB995-DEDF-48A4-9101-9CA5B72FFBCC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeJobGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeJobGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeJobGroupResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeJobGroupResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DescribeJobGroupResponseBody) GetJobGroup() *DescribeJobGroupResponseBodyJobGroup {
	return s.JobGroup
}

func (s *DescribeJobGroupResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeJobGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeJobGroupResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeJobGroupResponseBody) SetCode(v string) *DescribeJobGroupResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeJobGroupResponseBody) SetHttpStatusCode(v int32) *DescribeJobGroupResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeJobGroupResponseBody) SetJobGroup(v *DescribeJobGroupResponseBodyJobGroup) *DescribeJobGroupResponseBody {
	s.JobGroup = v
	return s
}

func (s *DescribeJobGroupResponseBody) SetMessage(v string) *DescribeJobGroupResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeJobGroupResponseBody) SetRequestId(v string) *DescribeJobGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeJobGroupResponseBody) SetSuccess(v bool) *DescribeJobGroupResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeJobGroupResponseBody) Validate() error {
	if s.JobGroup != nil {
		if err := s.JobGroup.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeJobGroupResponseBodyJobGroup struct {
	// The list of calling numbers.
	CallingNumbers []*string `json:"CallingNumbers,omitempty" xml:"CallingNumbers,omitempty" type:"Repeated"`
	// The creation time. The value is a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1578881227404
	CreationTime *int64 `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// **[Deprecated]*	- The export progress.
	//
	// > This information is obtained from the ListJobGroups operation.
	ExportProgress *DescribeJobGroupResponseBodyJobGroupExportProgress `json:"ExportProgress,omitempty" xml:"ExportProgress,omitempty" type:"Struct"`
	// The configuration parameters for flash SMS push in JSON format, including the configuration information of third-party flash SMS.
	FlashSmsExtras *DescribeJobGroupResponseBodyJobGroupFlashSmsExtras `json:"FlashSmsExtras,omitempty" xml:"FlashSmsExtras,omitempty" type:"Struct"`
	// **[Deprecated]*	- The ID of the job list parsing operation.
	//
	// example:
	//
	// da37319b-6c83-4268-9f19-814aed62e401/a5a9a310-b902-4674-a6e1-29975cbaa312_100.xlsx
	JobDataParsingTaskId *string `json:"JobDataParsingTaskId,omitempty" xml:"JobDataParsingTaskId,omitempty"`
	// **[Deprecated]*	- The OSS file key of the uploaded job list.
	//
	// example:
	//
	// da37319b-6c83-4268-9f19-814aed62e401/a5a9a310-b902-4674-a6e1-29975cbaa312_100.xlsx
	JobFilePath *string `json:"JobFilePath,omitempty" xml:"JobFilePath,omitempty"`
	// The description of the job group.
	//
	// example:
	//
	// Collection job group
	JobGroupDescription *string `json:"JobGroupDescription,omitempty" xml:"JobGroupDescription,omitempty"`
	// The ID of the job group.
	//
	// example:
	//
	// 46a9ad0c-3e11-44da-a9a7-2c21bf5ce185
	JobGroupId *string `json:"JobGroupId,omitempty" xml:"JobGroupId,omitempty"`
	// The name of the job group.
	//
	// example:
	//
	// Collection job group
	JobGroupName *string `json:"JobGroupName,omitempty" xml:"JobGroupName,omitempty"`
	// The minimum concurrency.
	//
	// example:
	//
	// 1
	MinConcurrency *int64 `json:"MinConcurrency,omitempty" xml:"MinConcurrency,omitempty"`
	// The modification time. The value is a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1578881227404
	ModifyTime *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// The priority.
	//
	// example:
	//
	// 1
	Priority *string `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The job progress.
	Progress *DescribeJobGroupResponseBodyJobGroupProgress `json:"Progress,omitempty" xml:"Progress,omitempty" type:"Struct"`
	// The list of calling numbers used for recall calls when the initial call fails. If not specified, it defaults to the primary calling number list.
	RecallCallingNumbers []*string `json:"RecallCallingNumbers,omitempty" xml:"RecallCallingNumbers,omitempty" type:"Repeated"`
	// The redial strategy.
	RecallStrategy *DescribeJobGroupResponseBodyJobGroupRecallStrategy `json:"RecallStrategy,omitempty" xml:"RecallStrategy,omitempty" type:"Struct"`
	// **[Deprecated]*	- The statistics information. This is a legacy parameter and is no longer provided.
	Result *DescribeJobGroupResponseBodyJobGroupResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// The ringing duration.
	//
	// example:
	//
	// 30
	RingingDuration *int64 `json:"RingingDuration,omitempty" xml:"RingingDuration,omitempty"`
	// The scenario ID. This is a legacy canvas parameter. For the new canvas, the scenario ID is the same as the ScriptId.
	//
	// example:
	//
	// fce6c599-8ede-40e3-9f78-0928eda7b4e8
	ScenarioId *string `json:"ScenarioId,omitempty" xml:"ScenarioId,omitempty"`
	// The ID of the script scenario.
	//
	// example:
	//
	// 49f00b0d-78ac-4d51-91de-a9e8e92b8470
	ScriptId *string `json:"ScriptId,omitempty" xml:"ScriptId,omitempty"`
	// The name of the script scenario.
	//
	// example:
	//
	// Collection script
	ScriptName *string `json:"ScriptName,omitempty" xml:"ScriptName,omitempty"`
	// The natural language understanding (NLU) engine used by the script.
	//
	// example:
	//
	// Prompts
	ScriptNluEngine *string `json:"ScriptNluEngine,omitempty" xml:"ScriptNluEngine,omitempty"`
	// The version number of the script scenario.
	//
	// example:
	//
	// 49f00b0d-78ac-4d51-91de-a9e8e92b8470
	ScriptVersion *string `json:"ScriptVersion,omitempty" xml:"ScriptVersion,omitempty"`
	// The job status.
	//
	// example:
	//
	// Completed
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The execution strategy of the job.
	Strategy *DescribeJobGroupResponseBodyJobGroupStrategy `json:"Strategy,omitempty" xml:"Strategy,omitempty" type:"Struct"`
}

func (s DescribeJobGroupResponseBodyJobGroup) String() string {
	return dara.Prettify(s)
}

func (s DescribeJobGroupResponseBodyJobGroup) GoString() string {
	return s.String()
}

func (s *DescribeJobGroupResponseBodyJobGroup) GetCallingNumbers() []*string {
	return s.CallingNumbers
}

func (s *DescribeJobGroupResponseBodyJobGroup) GetCreationTime() *int64 {
	return s.CreationTime
}

func (s *DescribeJobGroupResponseBodyJobGroup) GetExportProgress() *DescribeJobGroupResponseBodyJobGroupExportProgress {
	return s.ExportProgress
}

func (s *DescribeJobGroupResponseBodyJobGroup) GetFlashSmsExtras() *DescribeJobGroupResponseBodyJobGroupFlashSmsExtras {
	return s.FlashSmsExtras
}

func (s *DescribeJobGroupResponseBodyJobGroup) GetJobDataParsingTaskId() *string {
	return s.JobDataParsingTaskId
}

func (s *DescribeJobGroupResponseBodyJobGroup) GetJobFilePath() *string {
	return s.JobFilePath
}

func (s *DescribeJobGroupResponseBodyJobGroup) GetJobGroupDescription() *string {
	return s.JobGroupDescription
}

func (s *DescribeJobGroupResponseBodyJobGroup) GetJobGroupId() *string {
	return s.JobGroupId
}

func (s *DescribeJobGroupResponseBodyJobGroup) GetJobGroupName() *string {
	return s.JobGroupName
}

func (s *DescribeJobGroupResponseBodyJobGroup) GetMinConcurrency() *int64 {
	return s.MinConcurrency
}

func (s *DescribeJobGroupResponseBodyJobGroup) GetModifyTime() *string {
	return s.ModifyTime
}

func (s *DescribeJobGroupResponseBodyJobGroup) GetPriority() *string {
	return s.Priority
}

func (s *DescribeJobGroupResponseBodyJobGroup) GetProgress() *DescribeJobGroupResponseBodyJobGroupProgress {
	return s.Progress
}

func (s *DescribeJobGroupResponseBodyJobGroup) GetRecallCallingNumbers() []*string {
	return s.RecallCallingNumbers
}

func (s *DescribeJobGroupResponseBodyJobGroup) GetRecallStrategy() *DescribeJobGroupResponseBodyJobGroupRecallStrategy {
	return s.RecallStrategy
}

func (s *DescribeJobGroupResponseBodyJobGroup) GetResult() *DescribeJobGroupResponseBodyJobGroupResult {
	return s.Result
}

func (s *DescribeJobGroupResponseBodyJobGroup) GetRingingDuration() *int64 {
	return s.RingingDuration
}

func (s *DescribeJobGroupResponseBodyJobGroup) GetScenarioId() *string {
	return s.ScenarioId
}

func (s *DescribeJobGroupResponseBodyJobGroup) GetScriptId() *string {
	return s.ScriptId
}

func (s *DescribeJobGroupResponseBodyJobGroup) GetScriptName() *string {
	return s.ScriptName
}

func (s *DescribeJobGroupResponseBodyJobGroup) GetScriptNluEngine() *string {
	return s.ScriptNluEngine
}

func (s *DescribeJobGroupResponseBodyJobGroup) GetScriptVersion() *string {
	return s.ScriptVersion
}

func (s *DescribeJobGroupResponseBodyJobGroup) GetStatus() *string {
	return s.Status
}

func (s *DescribeJobGroupResponseBodyJobGroup) GetStrategy() *DescribeJobGroupResponseBodyJobGroupStrategy {
	return s.Strategy
}

func (s *DescribeJobGroupResponseBodyJobGroup) SetCallingNumbers(v []*string) *DescribeJobGroupResponseBodyJobGroup {
	s.CallingNumbers = v
	return s
}

func (s *DescribeJobGroupResponseBodyJobGroup) SetCreationTime(v int64) *DescribeJobGroupResponseBodyJobGroup {
	s.CreationTime = &v
	return s
}

func (s *DescribeJobGroupResponseBodyJobGroup) SetExportProgress(v *DescribeJobGroupResponseBodyJobGroupExportProgress) *DescribeJobGroupResponseBodyJobGroup {
	s.ExportProgress = v
	return s
}

func (s *DescribeJobGroupResponseBodyJobGroup) SetFlashSmsExtras(v *DescribeJobGroupResponseBodyJobGroupFlashSmsExtras) *DescribeJobGroupResponseBodyJobGroup {
	s.FlashSmsExtras = v
	return s
}

func (s *DescribeJobGroupResponseBodyJobGroup) SetJobDataParsingTaskId(v string) *DescribeJobGroupResponseBodyJobGroup {
	s.JobDataParsingTaskId = &v
	return s
}

func (s *DescribeJobGroupResponseBodyJobGroup) SetJobFilePath(v string) *DescribeJobGroupResponseBodyJobGroup {
	s.JobFilePath = &v
	return s
}

func (s *DescribeJobGroupResponseBodyJobGroup) SetJobGroupDescription(v string) *DescribeJobGroupResponseBodyJobGroup {
	s.JobGroupDescription = &v
	return s
}

func (s *DescribeJobGroupResponseBodyJobGroup) SetJobGroupId(v string) *DescribeJobGroupResponseBodyJobGroup {
	s.JobGroupId = &v
	return s
}

func (s *DescribeJobGroupResponseBodyJobGroup) SetJobGroupName(v string) *DescribeJobGroupResponseBodyJobGroup {
	s.JobGroupName = &v
	return s
}

func (s *DescribeJobGroupResponseBodyJobGroup) SetMinConcurrency(v int64) *DescribeJobGroupResponseBodyJobGroup {
	s.MinConcurrency = &v
	return s
}

func (s *DescribeJobGroupResponseBodyJobGroup) SetModifyTime(v string) *DescribeJobGroupResponseBodyJobGroup {
	s.ModifyTime = &v
	return s
}

func (s *DescribeJobGroupResponseBodyJobGroup) SetPriority(v string) *DescribeJobGroupResponseBodyJobGroup {
	s.Priority = &v
	return s
}

func (s *DescribeJobGroupResponseBodyJobGroup) SetProgress(v *DescribeJobGroupResponseBodyJobGroupProgress) *DescribeJobGroupResponseBodyJobGroup {
	s.Progress = v
	return s
}

func (s *DescribeJobGroupResponseBodyJobGroup) SetRecallCallingNumbers(v []*string) *DescribeJobGroupResponseBodyJobGroup {
	s.RecallCallingNumbers = v
	return s
}

func (s *DescribeJobGroupResponseBodyJobGroup) SetRecallStrategy(v *DescribeJobGroupResponseBodyJobGroupRecallStrategy) *DescribeJobGroupResponseBodyJobGroup {
	s.RecallStrategy = v
	return s
}

func (s *DescribeJobGroupResponseBodyJobGroup) SetResult(v *DescribeJobGroupResponseBodyJobGroupResult) *DescribeJobGroupResponseBodyJobGroup {
	s.Result = v
	return s
}

func (s *DescribeJobGroupResponseBodyJobGroup) SetRingingDuration(v int64) *DescribeJobGroupResponseBodyJobGroup {
	s.RingingDuration = &v
	return s
}

func (s *DescribeJobGroupResponseBodyJobGroup) SetScenarioId(v string) *DescribeJobGroupResponseBodyJobGroup {
	s.ScenarioId = &v
	return s
}

func (s *DescribeJobGroupResponseBodyJobGroup) SetScriptId(v string) *DescribeJobGroupResponseBodyJobGroup {
	s.ScriptId = &v
	return s
}

func (s *DescribeJobGroupResponseBodyJobGroup) SetScriptName(v string) *DescribeJobGroupResponseBodyJobGroup {
	s.ScriptName = &v
	return s
}

func (s *DescribeJobGroupResponseBodyJobGroup) SetScriptNluEngine(v string) *DescribeJobGroupResponseBodyJobGroup {
	s.ScriptNluEngine = &v
	return s
}

func (s *DescribeJobGroupResponseBodyJobGroup) SetScriptVersion(v string) *DescribeJobGroupResponseBodyJobGroup {
	s.ScriptVersion = &v
	return s
}

func (s *DescribeJobGroupResponseBodyJobGroup) SetStatus(v string) *DescribeJobGroupResponseBodyJobGroup {
	s.Status = &v
	return s
}

func (s *DescribeJobGroupResponseBodyJobGroup) SetStrategy(v *DescribeJobGroupResponseBodyJobGroupStrategy) *DescribeJobGroupResponseBodyJobGroup {
	s.Strategy = v
	return s
}

func (s *DescribeJobGroupResponseBodyJobGroup) Validate() error {
	if s.ExportProgress != nil {
		if err := s.ExportProgress.Validate(); err != nil {
			return err
		}
	}
	if s.FlashSmsExtras != nil {
		if err := s.FlashSmsExtras.Validate(); err != nil {
			return err
		}
	}
	if s.Progress != nil {
		if err := s.Progress.Validate(); err != nil {
			return err
		}
	}
	if s.RecallStrategy != nil {
		if err := s.RecallStrategy.Validate(); err != nil {
			return err
		}
	}
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	if s.Strategy != nil {
		if err := s.Strategy.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeJobGroupResponseBodyJobGroupExportProgress struct {
	// **[Deprecated]*	- The file URL.
	//
	// example:
	//
	// https://oss-cn-shanghai.aliyuncs.com/xx.zip
	FileHttpUrl *string `json:"FileHttpUrl,omitempty" xml:"FileHttpUrl,omitempty"`
	// **[Deprecated]*	- The export progress.
	//
	// example:
	//
	// 1
	Progress *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// **[Deprecated]*	- The export status.
	//
	// example:
	//
	// FINISHED
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeJobGroupResponseBodyJobGroupExportProgress) String() string {
	return dara.Prettify(s)
}

func (s DescribeJobGroupResponseBodyJobGroupExportProgress) GoString() string {
	return s.String()
}

func (s *DescribeJobGroupResponseBodyJobGroupExportProgress) GetFileHttpUrl() *string {
	return s.FileHttpUrl
}

func (s *DescribeJobGroupResponseBodyJobGroupExportProgress) GetProgress() *string {
	return s.Progress
}

func (s *DescribeJobGroupResponseBodyJobGroupExportProgress) GetStatus() *string {
	return s.Status
}

func (s *DescribeJobGroupResponseBodyJobGroupExportProgress) SetFileHttpUrl(v string) *DescribeJobGroupResponseBodyJobGroupExportProgress {
	s.FileHttpUrl = &v
	return s
}

func (s *DescribeJobGroupResponseBodyJobGroupExportProgress) SetProgress(v string) *DescribeJobGroupResponseBodyJobGroupExportProgress {
	s.Progress = &v
	return s
}

func (s *DescribeJobGroupResponseBodyJobGroupExportProgress) SetStatus(v string) *DescribeJobGroupResponseBodyJobGroupExportProgress {
	s.Status = &v
	return s
}

func (s *DescribeJobGroupResponseBodyJobGroupExportProgress) Validate() error {
	return dara.Validate(s)
}

type DescribeJobGroupResponseBodyJobGroupFlashSmsExtras struct {
	// The ID of the flash SMS configuration.
	//
	// example:
	//
	// bebe23e6e2604f50bbeb12900c07703d
	ConfigId *string `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// The ID of the flash SMS template.
	//
	// example:
	//
	// e080d737-5c0b-4350-b7d4-1162530579e7
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s DescribeJobGroupResponseBodyJobGroupFlashSmsExtras) String() string {
	return dara.Prettify(s)
}

func (s DescribeJobGroupResponseBodyJobGroupFlashSmsExtras) GoString() string {
	return s.String()
}

func (s *DescribeJobGroupResponseBodyJobGroupFlashSmsExtras) GetConfigId() *string {
	return s.ConfigId
}

func (s *DescribeJobGroupResponseBodyJobGroupFlashSmsExtras) GetTemplateId() *string {
	return s.TemplateId
}

func (s *DescribeJobGroupResponseBodyJobGroupFlashSmsExtras) SetConfigId(v string) *DescribeJobGroupResponseBodyJobGroupFlashSmsExtras {
	s.ConfigId = &v
	return s
}

func (s *DescribeJobGroupResponseBodyJobGroupFlashSmsExtras) SetTemplateId(v string) *DescribeJobGroupResponseBodyJobGroupFlashSmsExtras {
	s.TemplateId = &v
	return s
}

func (s *DescribeJobGroupResponseBodyJobGroupFlashSmsExtras) Validate() error {
	return dara.Validate(s)
}

type DescribeJobGroupResponseBodyJobGroupProgress struct {
	// **[Deprecated]*	- The business result information.
	Briefs []*DescribeJobGroupResponseBodyJobGroupProgressBriefs `json:"Briefs,omitempty" xml:"Briefs,omitempty" type:"Repeated"`
	// The number of canceled jobs.
	//
	// example:
	//
	// 0
	Cancelled *int32 `json:"Cancelled,omitempty" xml:"Cancelled,omitempty"`
	// **[Deprecated]*	- The statistics information in key-value format.
	Categories []*DescribeJobGroupResponseBodyJobGroupProgressCategories `json:"Categories,omitempty" xml:"Categories,omitempty" type:"Repeated"`
	// **[Deprecated]*	- The execution duration of the job.
	//
	// example:
	//
	// 1000
	Duration *int32 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The number of executing jobs.
	//
	// example:
	//
	// 2
	Executing *int32 `json:"Executing,omitempty" xml:"Executing,omitempty"`
	// The number of failed jobs.
	//
	// example:
	//
	// 0
	Failed *int32 `json:"Failed,omitempty" xml:"Failed,omitempty"`
	// The number of paused jobs.
	//
	// example:
	//
	// 0
	Paused *int32 `json:"Paused,omitempty" xml:"Paused,omitempty"`
	// The number of jobs being scheduled.
	//
	// example:
	//
	// 5
	Scheduling *int32 `json:"Scheduling,omitempty" xml:"Scheduling,omitempty"`
	// **[Deprecated]*	- The start time of the execution. The value is a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1578881227404
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// **[Deprecated]*	- The overall execution status of the job.
	//
	// example:
	//
	// Executing
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The total number of completed jobs.
	//
	// example:
	//
	// 2
	TotalCompleted *int32 `json:"TotalCompleted,omitempty" xml:"TotalCompleted,omitempty"`
	// The total number of jobs.
	//
	// example:
	//
	// 10
	TotalJobs *int32 `json:"TotalJobs,omitempty" xml:"TotalJobs,omitempty"`
	// **[Deprecated]*	- The total number of unanswered jobs.
	//
	// example:
	//
	// 1
	TotalNotAnswered *int32 `json:"TotalNotAnswered,omitempty" xml:"TotalNotAnswered,omitempty"`
}

func (s DescribeJobGroupResponseBodyJobGroupProgress) String() string {
	return dara.Prettify(s)
}

func (s DescribeJobGroupResponseBodyJobGroupProgress) GoString() string {
	return s.String()
}

func (s *DescribeJobGroupResponseBodyJobGroupProgress) GetBriefs() []*DescribeJobGroupResponseBodyJobGroupProgressBriefs {
	return s.Briefs
}

func (s *DescribeJobGroupResponseBodyJobGroupProgress) GetCancelled() *int32 {
	return s.Cancelled
}

func (s *DescribeJobGroupResponseBodyJobGroupProgress) GetCategories() []*DescribeJobGroupResponseBodyJobGroupProgressCategories {
	return s.Categories
}

func (s *DescribeJobGroupResponseBodyJobGroupProgress) GetDuration() *int32 {
	return s.Duration
}

func (s *DescribeJobGroupResponseBodyJobGroupProgress) GetExecuting() *int32 {
	return s.Executing
}

func (s *DescribeJobGroupResponseBodyJobGroupProgress) GetFailed() *int32 {
	return s.Failed
}

func (s *DescribeJobGroupResponseBodyJobGroupProgress) GetPaused() *int32 {
	return s.Paused
}

func (s *DescribeJobGroupResponseBodyJobGroupProgress) GetScheduling() *int32 {
	return s.Scheduling
}

func (s *DescribeJobGroupResponseBodyJobGroupProgress) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeJobGroupResponseBodyJobGroupProgress) GetStatus() *string {
	return s.Status
}

func (s *DescribeJobGroupResponseBodyJobGroupProgress) GetTotalCompleted() *int32 {
	return s.TotalCompleted
}

func (s *DescribeJobGroupResponseBodyJobGroupProgress) GetTotalJobs() *int32 {
	return s.TotalJobs
}

func (s *DescribeJobGroupResponseBodyJobGroupProgress) GetTotalNotAnswered() *int32 {
	return s.TotalNotAnswered
}

func (s *DescribeJobGroupResponseBodyJobGroupProgress) SetBriefs(v []*DescribeJobGroupResponseBodyJobGroupProgressBriefs) *DescribeJobGroupResponseBodyJobGroupProgress {
	s.Briefs = v
	return s
}

func (s *DescribeJobGroupResponseBodyJobGroupProgress) SetCancelled(v int32) *DescribeJobGroupResponseBodyJobGroupProgress {
	s.Cancelled = &v
	return s
}

func (s *DescribeJobGroupResponseBodyJobGroupProgress) SetCategories(v []*DescribeJobGroupResponseBodyJobGroupProgressCategories) *DescribeJobGroupResponseBodyJobGroupProgress {
	s.Categories = v
	return s
}

func (s *DescribeJobGroupResponseBodyJobGroupProgress) SetDuration(v int32) *DescribeJobGroupResponseBodyJobGroupProgress {
	s.Duration = &v
	return s
}

func (s *DescribeJobGroupResponseBodyJobGroupProgress) SetExecuting(v int32) *DescribeJobGroupResponseBodyJobGroupProgress {
	s.Executing = &v
	return s
}

func (s *DescribeJobGroupResponseBodyJobGroupProgress) SetFailed(v int32) *DescribeJobGroupResponseBodyJobGroupProgress {
	s.Failed = &v
	return s
}

func (s *DescribeJobGroupResponseBodyJobGroupProgress) SetPaused(v int32) *DescribeJobGroupResponseBodyJobGroupProgress {
	s.Paused = &v
	return s
}

func (s *DescribeJobGroupResponseBodyJobGroupProgress) SetScheduling(v int32) *DescribeJobGroupResponseBodyJobGroupProgress {
	s.Scheduling = &v
	return s
}

func (s *DescribeJobGroupResponseBodyJobGroupProgress) SetStartTime(v int64) *DescribeJobGroupResponseBodyJobGroupProgress {
	s.StartTime = &v
	return s
}

func (s *DescribeJobGroupResponseBodyJobGroupProgress) SetStatus(v string) *DescribeJobGroupResponseBodyJobGroupProgress {
	s.Status = &v
	return s
}

func (s *DescribeJobGroupResponseBodyJobGroupProgress) SetTotalCompleted(v int32) *DescribeJobGroupResponseBodyJobGroupProgress {
	s.TotalCompleted = &v
	return s
}

func (s *DescribeJobGroupResponseBodyJobGroupProgress) SetTotalJobs(v int32) *DescribeJobGroupResponseBodyJobGroupProgress {
	s.TotalJobs = &v
	return s
}

func (s *DescribeJobGroupResponseBodyJobGroupProgress) SetTotalNotAnswered(v int32) *DescribeJobGroupResponseBodyJobGroupProgress {
	s.TotalNotAnswered = &v
	return s
}

func (s *DescribeJobGroupResponseBodyJobGroupProgress) Validate() error {
	if s.Briefs != nil {
		for _, item := range s.Briefs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Categories != nil {
		for _, item := range s.Categories {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeJobGroupResponseBodyJobGroupProgressBriefs struct {
	// The name of the business result.
	//
	// example:
	//
	// score
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the business result.
	//
	// example:
	//
	// 5
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeJobGroupResponseBodyJobGroupProgressBriefs) String() string {
	return dara.Prettify(s)
}

func (s DescribeJobGroupResponseBodyJobGroupProgressBriefs) GoString() string {
	return s.String()
}

func (s *DescribeJobGroupResponseBodyJobGroupProgressBriefs) GetKey() *string {
	return s.Key
}

func (s *DescribeJobGroupResponseBodyJobGroupProgressBriefs) GetValue() *string {
	return s.Value
}

func (s *DescribeJobGroupResponseBodyJobGroupProgressBriefs) SetKey(v string) *DescribeJobGroupResponseBodyJobGroupProgressBriefs {
	s.Key = &v
	return s
}

func (s *DescribeJobGroupResponseBodyJobGroupProgressBriefs) SetValue(v string) *DescribeJobGroupResponseBodyJobGroupProgressBriefs {
	s.Value = &v
	return s
}

func (s *DescribeJobGroupResponseBodyJobGroupProgressBriefs) Validate() error {
	return dara.Validate(s)
}

type DescribeJobGroupResponseBodyJobGroupProgressCategories struct {
	// The name of the custom statistics information field.
	//
	// example:
	//
	// success
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The custom statistics information value (IV).
	//
	// example:
	//
	// 10
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeJobGroupResponseBodyJobGroupProgressCategories) String() string {
	return dara.Prettify(s)
}

func (s DescribeJobGroupResponseBodyJobGroupProgressCategories) GoString() string {
	return s.String()
}

func (s *DescribeJobGroupResponseBodyJobGroupProgressCategories) GetKey() *string {
	return s.Key
}

func (s *DescribeJobGroupResponseBodyJobGroupProgressCategories) GetValue() *string {
	return s.Value
}

func (s *DescribeJobGroupResponseBodyJobGroupProgressCategories) SetKey(v string) *DescribeJobGroupResponseBodyJobGroupProgressCategories {
	s.Key = &v
	return s
}

func (s *DescribeJobGroupResponseBodyJobGroupProgressCategories) SetValue(v string) *DescribeJobGroupResponseBodyJobGroupProgressCategories {
	s.Value = &v
	return s
}

func (s *DescribeJobGroupResponseBodyJobGroupProgressCategories) Validate() error {
	return dara.Validate(s)
}

type DescribeJobGroupResponseBodyJobGroupRecallStrategy struct {
	// Specifies whether to stop calling nonexistent numbers.
	//
	// example:
	//
	// false
	EmptyNumberIgnore *bool `json:"EmptyNumberIgnore,omitempty" xml:"EmptyNumberIgnore,omitempty"`
	// Specifies whether to stop calling numbers with overdue payments.
	//
	// example:
	//
	// false
	InArrearsIgnore *bool `json:"InArrearsIgnore,omitempty" xml:"InArrearsIgnore,omitempty"`
	// Specifies whether to stop calling out-of-service numbers.
	//
	// example:
	//
	// false
	OutOfServiceIgnore *bool `json:"OutOfServiceIgnore,omitempty" xml:"OutOfServiceIgnore,omitempty"`
}

func (s DescribeJobGroupResponseBodyJobGroupRecallStrategy) String() string {
	return dara.Prettify(s)
}

func (s DescribeJobGroupResponseBodyJobGroupRecallStrategy) GoString() string {
	return s.String()
}

func (s *DescribeJobGroupResponseBodyJobGroupRecallStrategy) GetEmptyNumberIgnore() *bool {
	return s.EmptyNumberIgnore
}

func (s *DescribeJobGroupResponseBodyJobGroupRecallStrategy) GetInArrearsIgnore() *bool {
	return s.InArrearsIgnore
}

func (s *DescribeJobGroupResponseBodyJobGroupRecallStrategy) GetOutOfServiceIgnore() *bool {
	return s.OutOfServiceIgnore
}

func (s *DescribeJobGroupResponseBodyJobGroupRecallStrategy) SetEmptyNumberIgnore(v bool) *DescribeJobGroupResponseBodyJobGroupRecallStrategy {
	s.EmptyNumberIgnore = &v
	return s
}

func (s *DescribeJobGroupResponseBodyJobGroupRecallStrategy) SetInArrearsIgnore(v bool) *DescribeJobGroupResponseBodyJobGroupRecallStrategy {
	s.InArrearsIgnore = &v
	return s
}

func (s *DescribeJobGroupResponseBodyJobGroupRecallStrategy) SetOutOfServiceIgnore(v bool) *DescribeJobGroupResponseBodyJobGroupRecallStrategy {
	s.OutOfServiceIgnore = &v
	return s
}

func (s *DescribeJobGroupResponseBodyJobGroupRecallStrategy) Validate() error {
	return dara.Validate(s)
}

type DescribeJobGroupResponseBodyJobGroupResult struct {
	// **[Deprecated]*	- The number of hang-ups by the customer.
	//
	// example:
	//
	// 1
	ClientHangupNum *int32 `json:"ClientHangupNum,omitempty" xml:"ClientHangupNum,omitempty"`
	// **[Deprecated]*	- The number of completed jobs.
	//
	// example:
	//
	// 1
	FinishedNum *int32 `json:"FinishedNum,omitempty" xml:"FinishedNum,omitempty"`
	// **[Deprecated]*	- The number of jobs without interaction.
	//
	// example:
	//
	// 1
	NoInteractNum *int32 `json:"NoInteractNum,omitempty" xml:"NoInteractNum,omitempty"`
	// **[Deprecated]*	- The number of hang-ups due to timeout.
	//
	// example:
	//
	// 1
	TimeoutHangupNum *int32 `json:"TimeoutHangupNum,omitempty" xml:"TimeoutHangupNum,omitempty"`
	// **[Deprecated]*	- The number of unrecognized utterances.
	//
	// example:
	//
	// 1
	UnrecognizedNum *int32 `json:"UnrecognizedNum,omitempty" xml:"UnrecognizedNum,omitempty"`
}

func (s DescribeJobGroupResponseBodyJobGroupResult) String() string {
	return dara.Prettify(s)
}

func (s DescribeJobGroupResponseBodyJobGroupResult) GoString() string {
	return s.String()
}

func (s *DescribeJobGroupResponseBodyJobGroupResult) GetClientHangupNum() *int32 {
	return s.ClientHangupNum
}

func (s *DescribeJobGroupResponseBodyJobGroupResult) GetFinishedNum() *int32 {
	return s.FinishedNum
}

func (s *DescribeJobGroupResponseBodyJobGroupResult) GetNoInteractNum() *int32 {
	return s.NoInteractNum
}

func (s *DescribeJobGroupResponseBodyJobGroupResult) GetTimeoutHangupNum() *int32 {
	return s.TimeoutHangupNum
}

func (s *DescribeJobGroupResponseBodyJobGroupResult) GetUnrecognizedNum() *int32 {
	return s.UnrecognizedNum
}

func (s *DescribeJobGroupResponseBodyJobGroupResult) SetClientHangupNum(v int32) *DescribeJobGroupResponseBodyJobGroupResult {
	s.ClientHangupNum = &v
	return s
}

func (s *DescribeJobGroupResponseBodyJobGroupResult) SetFinishedNum(v int32) *DescribeJobGroupResponseBodyJobGroupResult {
	s.FinishedNum = &v
	return s
}

func (s *DescribeJobGroupResponseBodyJobGroupResult) SetNoInteractNum(v int32) *DescribeJobGroupResponseBodyJobGroupResult {
	s.NoInteractNum = &v
	return s
}

func (s *DescribeJobGroupResponseBodyJobGroupResult) SetTimeoutHangupNum(v int32) *DescribeJobGroupResponseBodyJobGroupResult {
	s.TimeoutHangupNum = &v
	return s
}

func (s *DescribeJobGroupResponseBodyJobGroupResult) SetUnrecognizedNum(v int32) *DescribeJobGroupResponseBodyJobGroupResult {
	s.UnrecognizedNum = &v
	return s
}

func (s *DescribeJobGroupResponseBodyJobGroupResult) Validate() error {
	return dara.Validate(s)
}

type DescribeJobGroupResponseBodyJobGroupStrategy struct {
	// The custom data of the strategy. This is a legacy parameter and is no longer used.
	//
	// example:
	//
	// {}
	Customized *string `json:"Customized,omitempty" xml:"Customized,omitempty"`
	// The overall end time of the scheduling strategy. The value is a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1579881227404
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The processing method after the execution cycle ends. This is a legacy parameter and is no longer used.
	//
	// example:
	//
	// NONE
	FollowUpStrategy *string `json:"FollowUpStrategy,omitempty" xml:"FollowUpStrategy,omitempty"`
	// Indicates whether the strategy is a template.
	//
	// example:
	//
	// false
	IsTemplate *bool `json:"IsTemplate,omitempty" xml:"IsTemplate,omitempty"`
	// The maximum number of calls per day when the call is not answered.
	//
	// example:
	//
	// 2
	MaxAttemptsPerDay *int32 `json:"MaxAttemptsPerDay,omitempty" xml:"MaxAttemptsPerDay,omitempty"`
	// The interval for redialing.
	//
	// example:
	//
	// 10
	MinAttemptInterval *int32 `json:"MinAttemptInterval,omitempty" xml:"MinAttemptInterval,omitempty"`
	// The repetition method of the job.
	//
	// example:
	//
	// Once
	RepeatBy *string `json:"RepeatBy,omitempty" xml:"RepeatBy,omitempty"`
	// The days for repeated execution. This parameter is used in combination with RepeatBy.
	//
	// - If the repetition cycle is weekly, 0 indicates Sunday, and 1 to 6 indicate Monday to Saturday.
	//
	// - If the repetition cycle is monthly, 1 to 31 indicate the 1st to the 31st day of the month. The job is not executed in months that do not have the specified day. For example, if the 30th day is selected, the job is not executed in February.
	RepeatDays []*string `json:"RepeatDays,omitempty" xml:"RepeatDays,omitempty" type:"Repeated"`
	// Indicates whether the job is a recurring task.
	//
	// example:
	//
	// false
	Repeatable *bool `json:"Repeatable,omitempty" xml:"Repeatable,omitempty"`
	// The number routing strategy.
	//
	// example:
	//
	// LocalFirst
	RoutingStrategy *string `json:"RoutingStrategy,omitempty" xml:"RoutingStrategy,omitempty"`
	// The overall start time of the strategy. The value is a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1578881227404
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The description of the strategy.
	//
	// example:
	//
	// Collection strategy
	StrategyDescription *string `json:"StrategyDescription,omitempty" xml:"StrategyDescription,omitempty"`
	// The ID of the strategy.
	//
	// example:
	//
	// a2bff22c-2604-4df2-83d6-5952e2438c5a
	StrategyId *string `json:"StrategyId,omitempty" xml:"StrategyId,omitempty"`
	// The name of the strategy.
	//
	// example:
	//
	// Collection strategy name
	StrategyName *string `json:"StrategyName,omitempty" xml:"StrategyName,omitempty"`
	// The strategy type.
	//
	// example:
	//
	// Repeatable
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The working hours.
	WorkingTime []*DescribeJobGroupResponseBodyJobGroupStrategyWorkingTime `json:"WorkingTime,omitempty" xml:"WorkingTime,omitempty" type:"Repeated"`
}

func (s DescribeJobGroupResponseBodyJobGroupStrategy) String() string {
	return dara.Prettify(s)
}

func (s DescribeJobGroupResponseBodyJobGroupStrategy) GoString() string {
	return s.String()
}

func (s *DescribeJobGroupResponseBodyJobGroupStrategy) GetCustomized() *string {
	return s.Customized
}

func (s *DescribeJobGroupResponseBodyJobGroupStrategy) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeJobGroupResponseBodyJobGroupStrategy) GetFollowUpStrategy() *string {
	return s.FollowUpStrategy
}

func (s *DescribeJobGroupResponseBodyJobGroupStrategy) GetIsTemplate() *bool {
	return s.IsTemplate
}

func (s *DescribeJobGroupResponseBodyJobGroupStrategy) GetMaxAttemptsPerDay() *int32 {
	return s.MaxAttemptsPerDay
}

func (s *DescribeJobGroupResponseBodyJobGroupStrategy) GetMinAttemptInterval() *int32 {
	return s.MinAttemptInterval
}

func (s *DescribeJobGroupResponseBodyJobGroupStrategy) GetRepeatBy() *string {
	return s.RepeatBy
}

func (s *DescribeJobGroupResponseBodyJobGroupStrategy) GetRepeatDays() []*string {
	return s.RepeatDays
}

func (s *DescribeJobGroupResponseBodyJobGroupStrategy) GetRepeatable() *bool {
	return s.Repeatable
}

func (s *DescribeJobGroupResponseBodyJobGroupStrategy) GetRoutingStrategy() *string {
	return s.RoutingStrategy
}

func (s *DescribeJobGroupResponseBodyJobGroupStrategy) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeJobGroupResponseBodyJobGroupStrategy) GetStrategyDescription() *string {
	return s.StrategyDescription
}

func (s *DescribeJobGroupResponseBodyJobGroupStrategy) GetStrategyId() *string {
	return s.StrategyId
}

func (s *DescribeJobGroupResponseBodyJobGroupStrategy) GetStrategyName() *string {
	return s.StrategyName
}

func (s *DescribeJobGroupResponseBodyJobGroupStrategy) GetType() *string {
	return s.Type
}

func (s *DescribeJobGroupResponseBodyJobGroupStrategy) GetWorkingTime() []*DescribeJobGroupResponseBodyJobGroupStrategyWorkingTime {
	return s.WorkingTime
}

func (s *DescribeJobGroupResponseBodyJobGroupStrategy) SetCustomized(v string) *DescribeJobGroupResponseBodyJobGroupStrategy {
	s.Customized = &v
	return s
}

func (s *DescribeJobGroupResponseBodyJobGroupStrategy) SetEndTime(v int64) *DescribeJobGroupResponseBodyJobGroupStrategy {
	s.EndTime = &v
	return s
}

func (s *DescribeJobGroupResponseBodyJobGroupStrategy) SetFollowUpStrategy(v string) *DescribeJobGroupResponseBodyJobGroupStrategy {
	s.FollowUpStrategy = &v
	return s
}

func (s *DescribeJobGroupResponseBodyJobGroupStrategy) SetIsTemplate(v bool) *DescribeJobGroupResponseBodyJobGroupStrategy {
	s.IsTemplate = &v
	return s
}

func (s *DescribeJobGroupResponseBodyJobGroupStrategy) SetMaxAttemptsPerDay(v int32) *DescribeJobGroupResponseBodyJobGroupStrategy {
	s.MaxAttemptsPerDay = &v
	return s
}

func (s *DescribeJobGroupResponseBodyJobGroupStrategy) SetMinAttemptInterval(v int32) *DescribeJobGroupResponseBodyJobGroupStrategy {
	s.MinAttemptInterval = &v
	return s
}

func (s *DescribeJobGroupResponseBodyJobGroupStrategy) SetRepeatBy(v string) *DescribeJobGroupResponseBodyJobGroupStrategy {
	s.RepeatBy = &v
	return s
}

func (s *DescribeJobGroupResponseBodyJobGroupStrategy) SetRepeatDays(v []*string) *DescribeJobGroupResponseBodyJobGroupStrategy {
	s.RepeatDays = v
	return s
}

func (s *DescribeJobGroupResponseBodyJobGroupStrategy) SetRepeatable(v bool) *DescribeJobGroupResponseBodyJobGroupStrategy {
	s.Repeatable = &v
	return s
}

func (s *DescribeJobGroupResponseBodyJobGroupStrategy) SetRoutingStrategy(v string) *DescribeJobGroupResponseBodyJobGroupStrategy {
	s.RoutingStrategy = &v
	return s
}

func (s *DescribeJobGroupResponseBodyJobGroupStrategy) SetStartTime(v int64) *DescribeJobGroupResponseBodyJobGroupStrategy {
	s.StartTime = &v
	return s
}

func (s *DescribeJobGroupResponseBodyJobGroupStrategy) SetStrategyDescription(v string) *DescribeJobGroupResponseBodyJobGroupStrategy {
	s.StrategyDescription = &v
	return s
}

func (s *DescribeJobGroupResponseBodyJobGroupStrategy) SetStrategyId(v string) *DescribeJobGroupResponseBodyJobGroupStrategy {
	s.StrategyId = &v
	return s
}

func (s *DescribeJobGroupResponseBodyJobGroupStrategy) SetStrategyName(v string) *DescribeJobGroupResponseBodyJobGroupStrategy {
	s.StrategyName = &v
	return s
}

func (s *DescribeJobGroupResponseBodyJobGroupStrategy) SetType(v string) *DescribeJobGroupResponseBodyJobGroupStrategy {
	s.Type = &v
	return s
}

func (s *DescribeJobGroupResponseBodyJobGroupStrategy) SetWorkingTime(v []*DescribeJobGroupResponseBodyJobGroupStrategyWorkingTime) *DescribeJobGroupResponseBodyJobGroupStrategy {
	s.WorkingTime = v
	return s
}

func (s *DescribeJobGroupResponseBodyJobGroupStrategy) Validate() error {
	if s.WorkingTime != nil {
		for _, item := range s.WorkingTime {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeJobGroupResponseBodyJobGroupStrategyWorkingTime struct {
	// The start time. The format is HH:mm:ss.
	//
	// example:
	//
	// 09:00:00
	BeginTime *string `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	// The end time. The format is HH:mm:ss.
	//
	// example:
	//
	// 12:00:00
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
}

func (s DescribeJobGroupResponseBodyJobGroupStrategyWorkingTime) String() string {
	return dara.Prettify(s)
}

func (s DescribeJobGroupResponseBodyJobGroupStrategyWorkingTime) GoString() string {
	return s.String()
}

func (s *DescribeJobGroupResponseBodyJobGroupStrategyWorkingTime) GetBeginTime() *string {
	return s.BeginTime
}

func (s *DescribeJobGroupResponseBodyJobGroupStrategyWorkingTime) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeJobGroupResponseBodyJobGroupStrategyWorkingTime) SetBeginTime(v string) *DescribeJobGroupResponseBodyJobGroupStrategyWorkingTime {
	s.BeginTime = &v
	return s
}

func (s *DescribeJobGroupResponseBodyJobGroupStrategyWorkingTime) SetEndTime(v string) *DescribeJobGroupResponseBodyJobGroupStrategyWorkingTime {
	s.EndTime = &v
	return s
}

func (s *DescribeJobGroupResponseBodyJobGroupStrategyWorkingTime) Validate() error {
	return dara.Validate(s)
}
