// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateJobGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateJobGroupResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *CreateJobGroupResponseBody
	GetHttpStatusCode() *int32
	SetJobGroup(v *CreateJobGroupResponseBodyJobGroup) *CreateJobGroupResponseBody
	GetJobGroup() *CreateJobGroupResponseBodyJobGroup
	SetMessage(v string) *CreateJobGroupResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateJobGroupResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateJobGroupResponseBody
	GetSuccess() *bool
}

type CreateJobGroupResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32                              `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	JobGroup       *CreateJobGroupResponseBodyJobGroup `json:"JobGroup,omitempty" xml:"JobGroup,omitempty" type:"Struct"`
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 254EB995-DEDF-48A4-9101-9CA5B72FFBCC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateJobGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateJobGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateJobGroupResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateJobGroupResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateJobGroupResponseBody) GetJobGroup() *CreateJobGroupResponseBodyJobGroup {
	return s.JobGroup
}

func (s *CreateJobGroupResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateJobGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateJobGroupResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateJobGroupResponseBody) SetCode(v string) *CreateJobGroupResponseBody {
	s.Code = &v
	return s
}

func (s *CreateJobGroupResponseBody) SetHttpStatusCode(v int32) *CreateJobGroupResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateJobGroupResponseBody) SetJobGroup(v *CreateJobGroupResponseBodyJobGroup) *CreateJobGroupResponseBody {
	s.JobGroup = v
	return s
}

func (s *CreateJobGroupResponseBody) SetMessage(v string) *CreateJobGroupResponseBody {
	s.Message = &v
	return s
}

func (s *CreateJobGroupResponseBody) SetRequestId(v string) *CreateJobGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateJobGroupResponseBody) SetSuccess(v bool) *CreateJobGroupResponseBody {
	s.Success = &v
	return s
}

func (s *CreateJobGroupResponseBody) Validate() error {
	if s.JobGroup != nil {
		if err := s.JobGroup.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateJobGroupResponseBodyJobGroup struct {
	CallingNumbers []*string `json:"CallingNumbers,omitempty" xml:"CallingNumbers,omitempty" type:"Repeated"`
	// example:
	//
	// 1578550074361
	CreationTime   *int64                                            `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	ExportProgress *CreateJobGroupResponseBodyJobGroupExportProgress `json:"ExportProgress,omitempty" xml:"ExportProgress,omitempty" type:"Struct"`
	// example:
	//
	// 744ff448-2b4c-40d4-94ca-51f246905b0f
	JobDataParsingTaskId *string `json:"JobDataParsingTaskId,omitempty" xml:"JobDataParsingTaskId,omitempty"`
	// example:
	//
	// UPLOADED/JOB/b3865dc3-40fa-4afd-9fe4-dc7cda305a24/229eac13-379d-4abe-96e0-8cf026b56c0b_template (1).xlsx
	JobFilePath         *string `json:"JobFilePath,omitempty" xml:"JobFilePath,omitempty"`
	JobGroupDescription *string `json:"JobGroupDescription,omitempty" xml:"JobGroupDescription,omitempty"`
	// example:
	//
	// c62e6789-28a8-41db-941e-171a01d3b3b9
	JobGroupId   *string `json:"JobGroupId,omitempty" xml:"JobGroupId,omitempty"`
	JobGroupName *string `json:"JobGroupName,omitempty" xml:"JobGroupName,omitempty"`
	// example:
	//
	// 1
	MinConcurrency *int64 `json:"MinConcurrency,omitempty" xml:"MinConcurrency,omitempty"`
	// example:
	//
	// 1628425608429
	ModifyTime           *string                                           `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	Priority             *string                                           `json:"Priority,omitempty" xml:"Priority,omitempty"`
	RecallCallingNumbers []*string                                         `json:"RecallCallingNumbers,omitempty" xml:"RecallCallingNumbers,omitempty" type:"Repeated"`
	RecallStrategy       *CreateJobGroupResponseBodyJobGroupRecallStrategy `json:"RecallStrategy,omitempty" xml:"RecallStrategy,omitempty" type:"Struct"`
	// example:
	//
	// 30
	RingingDuration *int64 `json:"RingingDuration,omitempty" xml:"RingingDuration,omitempty"`
	// example:
	//
	// 6cea9bed-63e6-439e-ae4c-b3333efff53d
	ScenarioId *string `json:"ScenarioId,omitempty" xml:"ScenarioId,omitempty"`
	ScriptName *string `json:"ScriptName,omitempty" xml:"ScriptName,omitempty"`
	// example:
	//
	// 1628425608429
	ScriptVersion *string `json:"ScriptVersion,omitempty" xml:"ScriptVersion,omitempty"`
	// example:
	//
	// Scheduling
	Status   *string                                     `json:"Status,omitempty" xml:"Status,omitempty"`
	Strategy *CreateJobGroupResponseBodyJobGroupStrategy `json:"Strategy,omitempty" xml:"Strategy,omitempty" type:"Struct"`
}

func (s CreateJobGroupResponseBodyJobGroup) String() string {
	return dara.Prettify(s)
}

func (s CreateJobGroupResponseBodyJobGroup) GoString() string {
	return s.String()
}

func (s *CreateJobGroupResponseBodyJobGroup) GetCallingNumbers() []*string {
	return s.CallingNumbers
}

func (s *CreateJobGroupResponseBodyJobGroup) GetCreationTime() *int64 {
	return s.CreationTime
}

func (s *CreateJobGroupResponseBodyJobGroup) GetExportProgress() *CreateJobGroupResponseBodyJobGroupExportProgress {
	return s.ExportProgress
}

func (s *CreateJobGroupResponseBodyJobGroup) GetJobDataParsingTaskId() *string {
	return s.JobDataParsingTaskId
}

func (s *CreateJobGroupResponseBodyJobGroup) GetJobFilePath() *string {
	return s.JobFilePath
}

func (s *CreateJobGroupResponseBodyJobGroup) GetJobGroupDescription() *string {
	return s.JobGroupDescription
}

func (s *CreateJobGroupResponseBodyJobGroup) GetJobGroupId() *string {
	return s.JobGroupId
}

func (s *CreateJobGroupResponseBodyJobGroup) GetJobGroupName() *string {
	return s.JobGroupName
}

func (s *CreateJobGroupResponseBodyJobGroup) GetMinConcurrency() *int64 {
	return s.MinConcurrency
}

func (s *CreateJobGroupResponseBodyJobGroup) GetModifyTime() *string {
	return s.ModifyTime
}

func (s *CreateJobGroupResponseBodyJobGroup) GetPriority() *string {
	return s.Priority
}

func (s *CreateJobGroupResponseBodyJobGroup) GetRecallCallingNumbers() []*string {
	return s.RecallCallingNumbers
}

func (s *CreateJobGroupResponseBodyJobGroup) GetRecallStrategy() *CreateJobGroupResponseBodyJobGroupRecallStrategy {
	return s.RecallStrategy
}

func (s *CreateJobGroupResponseBodyJobGroup) GetRingingDuration() *int64 {
	return s.RingingDuration
}

func (s *CreateJobGroupResponseBodyJobGroup) GetScenarioId() *string {
	return s.ScenarioId
}

func (s *CreateJobGroupResponseBodyJobGroup) GetScriptName() *string {
	return s.ScriptName
}

func (s *CreateJobGroupResponseBodyJobGroup) GetScriptVersion() *string {
	return s.ScriptVersion
}

func (s *CreateJobGroupResponseBodyJobGroup) GetStatus() *string {
	return s.Status
}

func (s *CreateJobGroupResponseBodyJobGroup) GetStrategy() *CreateJobGroupResponseBodyJobGroupStrategy {
	return s.Strategy
}

func (s *CreateJobGroupResponseBodyJobGroup) SetCallingNumbers(v []*string) *CreateJobGroupResponseBodyJobGroup {
	s.CallingNumbers = v
	return s
}

func (s *CreateJobGroupResponseBodyJobGroup) SetCreationTime(v int64) *CreateJobGroupResponseBodyJobGroup {
	s.CreationTime = &v
	return s
}

func (s *CreateJobGroupResponseBodyJobGroup) SetExportProgress(v *CreateJobGroupResponseBodyJobGroupExportProgress) *CreateJobGroupResponseBodyJobGroup {
	s.ExportProgress = v
	return s
}

func (s *CreateJobGroupResponseBodyJobGroup) SetJobDataParsingTaskId(v string) *CreateJobGroupResponseBodyJobGroup {
	s.JobDataParsingTaskId = &v
	return s
}

func (s *CreateJobGroupResponseBodyJobGroup) SetJobFilePath(v string) *CreateJobGroupResponseBodyJobGroup {
	s.JobFilePath = &v
	return s
}

func (s *CreateJobGroupResponseBodyJobGroup) SetJobGroupDescription(v string) *CreateJobGroupResponseBodyJobGroup {
	s.JobGroupDescription = &v
	return s
}

func (s *CreateJobGroupResponseBodyJobGroup) SetJobGroupId(v string) *CreateJobGroupResponseBodyJobGroup {
	s.JobGroupId = &v
	return s
}

func (s *CreateJobGroupResponseBodyJobGroup) SetJobGroupName(v string) *CreateJobGroupResponseBodyJobGroup {
	s.JobGroupName = &v
	return s
}

func (s *CreateJobGroupResponseBodyJobGroup) SetMinConcurrency(v int64) *CreateJobGroupResponseBodyJobGroup {
	s.MinConcurrency = &v
	return s
}

func (s *CreateJobGroupResponseBodyJobGroup) SetModifyTime(v string) *CreateJobGroupResponseBodyJobGroup {
	s.ModifyTime = &v
	return s
}

func (s *CreateJobGroupResponseBodyJobGroup) SetPriority(v string) *CreateJobGroupResponseBodyJobGroup {
	s.Priority = &v
	return s
}

func (s *CreateJobGroupResponseBodyJobGroup) SetRecallCallingNumbers(v []*string) *CreateJobGroupResponseBodyJobGroup {
	s.RecallCallingNumbers = v
	return s
}

func (s *CreateJobGroupResponseBodyJobGroup) SetRecallStrategy(v *CreateJobGroupResponseBodyJobGroupRecallStrategy) *CreateJobGroupResponseBodyJobGroup {
	s.RecallStrategy = v
	return s
}

func (s *CreateJobGroupResponseBodyJobGroup) SetRingingDuration(v int64) *CreateJobGroupResponseBodyJobGroup {
	s.RingingDuration = &v
	return s
}

func (s *CreateJobGroupResponseBodyJobGroup) SetScenarioId(v string) *CreateJobGroupResponseBodyJobGroup {
	s.ScenarioId = &v
	return s
}

func (s *CreateJobGroupResponseBodyJobGroup) SetScriptName(v string) *CreateJobGroupResponseBodyJobGroup {
	s.ScriptName = &v
	return s
}

func (s *CreateJobGroupResponseBodyJobGroup) SetScriptVersion(v string) *CreateJobGroupResponseBodyJobGroup {
	s.ScriptVersion = &v
	return s
}

func (s *CreateJobGroupResponseBodyJobGroup) SetStatus(v string) *CreateJobGroupResponseBodyJobGroup {
	s.Status = &v
	return s
}

func (s *CreateJobGroupResponseBodyJobGroup) SetStrategy(v *CreateJobGroupResponseBodyJobGroupStrategy) *CreateJobGroupResponseBodyJobGroup {
	s.Strategy = v
	return s
}

func (s *CreateJobGroupResponseBodyJobGroup) Validate() error {
	if s.ExportProgress != nil {
		if err := s.ExportProgress.Validate(); err != nil {
			return err
		}
	}
	if s.RecallStrategy != nil {
		if err := s.RecallStrategy.Validate(); err != nil {
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

type CreateJobGroupResponseBodyJobGroupExportProgress struct {
	// example:
	//
	// https://***.oss-cn-shanghai.aliyuncs.com/sample
	FileHttpUrl *string `json:"FileHttpUrl,omitempty" xml:"FileHttpUrl,omitempty"`
	// example:
	//
	// 1
	Progress *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// example:
	//
	// FINISHED
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s CreateJobGroupResponseBodyJobGroupExportProgress) String() string {
	return dara.Prettify(s)
}

func (s CreateJobGroupResponseBodyJobGroupExportProgress) GoString() string {
	return s.String()
}

func (s *CreateJobGroupResponseBodyJobGroupExportProgress) GetFileHttpUrl() *string {
	return s.FileHttpUrl
}

func (s *CreateJobGroupResponseBodyJobGroupExportProgress) GetProgress() *string {
	return s.Progress
}

func (s *CreateJobGroupResponseBodyJobGroupExportProgress) GetStatus() *string {
	return s.Status
}

func (s *CreateJobGroupResponseBodyJobGroupExportProgress) SetFileHttpUrl(v string) *CreateJobGroupResponseBodyJobGroupExportProgress {
	s.FileHttpUrl = &v
	return s
}

func (s *CreateJobGroupResponseBodyJobGroupExportProgress) SetProgress(v string) *CreateJobGroupResponseBodyJobGroupExportProgress {
	s.Progress = &v
	return s
}

func (s *CreateJobGroupResponseBodyJobGroupExportProgress) SetStatus(v string) *CreateJobGroupResponseBodyJobGroupExportProgress {
	s.Status = &v
	return s
}

func (s *CreateJobGroupResponseBodyJobGroupExportProgress) Validate() error {
	return dara.Validate(s)
}

type CreateJobGroupResponseBodyJobGroupRecallStrategy struct {
	EmptyNumberIgnore  *bool `json:"EmptyNumberIgnore,omitempty" xml:"EmptyNumberIgnore,omitempty"`
	InArrearsIgnore    *bool `json:"InArrearsIgnore,omitempty" xml:"InArrearsIgnore,omitempty"`
	OutOfServiceIgnore *bool `json:"OutOfServiceIgnore,omitempty" xml:"OutOfServiceIgnore,omitempty"`
}

func (s CreateJobGroupResponseBodyJobGroupRecallStrategy) String() string {
	return dara.Prettify(s)
}

func (s CreateJobGroupResponseBodyJobGroupRecallStrategy) GoString() string {
	return s.String()
}

func (s *CreateJobGroupResponseBodyJobGroupRecallStrategy) GetEmptyNumberIgnore() *bool {
	return s.EmptyNumberIgnore
}

func (s *CreateJobGroupResponseBodyJobGroupRecallStrategy) GetInArrearsIgnore() *bool {
	return s.InArrearsIgnore
}

func (s *CreateJobGroupResponseBodyJobGroupRecallStrategy) GetOutOfServiceIgnore() *bool {
	return s.OutOfServiceIgnore
}

func (s *CreateJobGroupResponseBodyJobGroupRecallStrategy) SetEmptyNumberIgnore(v bool) *CreateJobGroupResponseBodyJobGroupRecallStrategy {
	s.EmptyNumberIgnore = &v
	return s
}

func (s *CreateJobGroupResponseBodyJobGroupRecallStrategy) SetInArrearsIgnore(v bool) *CreateJobGroupResponseBodyJobGroupRecallStrategy {
	s.InArrearsIgnore = &v
	return s
}

func (s *CreateJobGroupResponseBodyJobGroupRecallStrategy) SetOutOfServiceIgnore(v bool) *CreateJobGroupResponseBodyJobGroupRecallStrategy {
	s.OutOfServiceIgnore = &v
	return s
}

func (s *CreateJobGroupResponseBodyJobGroupRecallStrategy) Validate() error {
	return dara.Validate(s)
}

type CreateJobGroupResponseBodyJobGroupStrategy struct {
	// example:
	//
	// {}
	Customized *string `json:"Customized,omitempty" xml:"Customized,omitempty"`
	// example:
	//
	// 2209702074000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// CONTINUE
	FollowUpStrategy *string `json:"FollowUpStrategy,omitempty" xml:"FollowUpStrategy,omitempty"`
	// example:
	//
	// false
	IsTemplate *bool `json:"IsTemplate,omitempty" xml:"IsTemplate,omitempty"`
	// example:
	//
	// 3
	MaxAttemptsPerDay *int32 `json:"MaxAttemptsPerDay,omitempty" xml:"MaxAttemptsPerDay,omitempty"`
	// example:
	//
	// 3
	MinAttemptInterval *int32 `json:"MinAttemptInterval,omitempty" xml:"MinAttemptInterval,omitempty"`
	// example:
	//
	// Once
	RepeatBy   *string   `json:"RepeatBy,omitempty" xml:"RepeatBy,omitempty"`
	RepeatDays []*string `json:"RepeatDays,omitempty" xml:"RepeatDays,omitempty" type:"Repeated"`
	// example:
	//
	// LocalFirst
	RoutingStrategy *string `json:"RoutingStrategy,omitempty" xml:"RoutingStrategy,omitempty"`
	// example:
	//
	// 1578550074000
	StartTime           *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	StrategyDescription *string `json:"StrategyDescription,omitempty" xml:"StrategyDescription,omitempty"`
	// example:
	//
	// cc9a436e-03b0-4ada-8364-77ec2290aa39
	StrategyId   *string `json:"StrategyId,omitempty" xml:"StrategyId,omitempty"`
	StrategyName *string `json:"StrategyName,omitempty" xml:"StrategyName,omitempty"`
	// example:
	//
	// Repeatable
	Type        *string                                                  `json:"Type,omitempty" xml:"Type,omitempty"`
	WorkingTime []*CreateJobGroupResponseBodyJobGroupStrategyWorkingTime `json:"WorkingTime,omitempty" xml:"WorkingTime,omitempty" type:"Repeated"`
}

func (s CreateJobGroupResponseBodyJobGroupStrategy) String() string {
	return dara.Prettify(s)
}

func (s CreateJobGroupResponseBodyJobGroupStrategy) GoString() string {
	return s.String()
}

func (s *CreateJobGroupResponseBodyJobGroupStrategy) GetCustomized() *string {
	return s.Customized
}

func (s *CreateJobGroupResponseBodyJobGroupStrategy) GetEndTime() *int64 {
	return s.EndTime
}

func (s *CreateJobGroupResponseBodyJobGroupStrategy) GetFollowUpStrategy() *string {
	return s.FollowUpStrategy
}

func (s *CreateJobGroupResponseBodyJobGroupStrategy) GetIsTemplate() *bool {
	return s.IsTemplate
}

func (s *CreateJobGroupResponseBodyJobGroupStrategy) GetMaxAttemptsPerDay() *int32 {
	return s.MaxAttemptsPerDay
}

func (s *CreateJobGroupResponseBodyJobGroupStrategy) GetMinAttemptInterval() *int32 {
	return s.MinAttemptInterval
}

func (s *CreateJobGroupResponseBodyJobGroupStrategy) GetRepeatBy() *string {
	return s.RepeatBy
}

func (s *CreateJobGroupResponseBodyJobGroupStrategy) GetRepeatDays() []*string {
	return s.RepeatDays
}

func (s *CreateJobGroupResponseBodyJobGroupStrategy) GetRoutingStrategy() *string {
	return s.RoutingStrategy
}

func (s *CreateJobGroupResponseBodyJobGroupStrategy) GetStartTime() *int64 {
	return s.StartTime
}

func (s *CreateJobGroupResponseBodyJobGroupStrategy) GetStrategyDescription() *string {
	return s.StrategyDescription
}

func (s *CreateJobGroupResponseBodyJobGroupStrategy) GetStrategyId() *string {
	return s.StrategyId
}

func (s *CreateJobGroupResponseBodyJobGroupStrategy) GetStrategyName() *string {
	return s.StrategyName
}

func (s *CreateJobGroupResponseBodyJobGroupStrategy) GetType() *string {
	return s.Type
}

func (s *CreateJobGroupResponseBodyJobGroupStrategy) GetWorkingTime() []*CreateJobGroupResponseBodyJobGroupStrategyWorkingTime {
	return s.WorkingTime
}

func (s *CreateJobGroupResponseBodyJobGroupStrategy) SetCustomized(v string) *CreateJobGroupResponseBodyJobGroupStrategy {
	s.Customized = &v
	return s
}

func (s *CreateJobGroupResponseBodyJobGroupStrategy) SetEndTime(v int64) *CreateJobGroupResponseBodyJobGroupStrategy {
	s.EndTime = &v
	return s
}

func (s *CreateJobGroupResponseBodyJobGroupStrategy) SetFollowUpStrategy(v string) *CreateJobGroupResponseBodyJobGroupStrategy {
	s.FollowUpStrategy = &v
	return s
}

func (s *CreateJobGroupResponseBodyJobGroupStrategy) SetIsTemplate(v bool) *CreateJobGroupResponseBodyJobGroupStrategy {
	s.IsTemplate = &v
	return s
}

func (s *CreateJobGroupResponseBodyJobGroupStrategy) SetMaxAttemptsPerDay(v int32) *CreateJobGroupResponseBodyJobGroupStrategy {
	s.MaxAttemptsPerDay = &v
	return s
}

func (s *CreateJobGroupResponseBodyJobGroupStrategy) SetMinAttemptInterval(v int32) *CreateJobGroupResponseBodyJobGroupStrategy {
	s.MinAttemptInterval = &v
	return s
}

func (s *CreateJobGroupResponseBodyJobGroupStrategy) SetRepeatBy(v string) *CreateJobGroupResponseBodyJobGroupStrategy {
	s.RepeatBy = &v
	return s
}

func (s *CreateJobGroupResponseBodyJobGroupStrategy) SetRepeatDays(v []*string) *CreateJobGroupResponseBodyJobGroupStrategy {
	s.RepeatDays = v
	return s
}

func (s *CreateJobGroupResponseBodyJobGroupStrategy) SetRoutingStrategy(v string) *CreateJobGroupResponseBodyJobGroupStrategy {
	s.RoutingStrategy = &v
	return s
}

func (s *CreateJobGroupResponseBodyJobGroupStrategy) SetStartTime(v int64) *CreateJobGroupResponseBodyJobGroupStrategy {
	s.StartTime = &v
	return s
}

func (s *CreateJobGroupResponseBodyJobGroupStrategy) SetStrategyDescription(v string) *CreateJobGroupResponseBodyJobGroupStrategy {
	s.StrategyDescription = &v
	return s
}

func (s *CreateJobGroupResponseBodyJobGroupStrategy) SetStrategyId(v string) *CreateJobGroupResponseBodyJobGroupStrategy {
	s.StrategyId = &v
	return s
}

func (s *CreateJobGroupResponseBodyJobGroupStrategy) SetStrategyName(v string) *CreateJobGroupResponseBodyJobGroupStrategy {
	s.StrategyName = &v
	return s
}

func (s *CreateJobGroupResponseBodyJobGroupStrategy) SetType(v string) *CreateJobGroupResponseBodyJobGroupStrategy {
	s.Type = &v
	return s
}

func (s *CreateJobGroupResponseBodyJobGroupStrategy) SetWorkingTime(v []*CreateJobGroupResponseBodyJobGroupStrategyWorkingTime) *CreateJobGroupResponseBodyJobGroupStrategy {
	s.WorkingTime = v
	return s
}

func (s *CreateJobGroupResponseBodyJobGroupStrategy) Validate() error {
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

type CreateJobGroupResponseBodyJobGroupStrategyWorkingTime struct {
	// example:
	//
	// 09:00:00
	BeginTime *string `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	// example:
	//
	// 12:00:00
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
}

func (s CreateJobGroupResponseBodyJobGroupStrategyWorkingTime) String() string {
	return dara.Prettify(s)
}

func (s CreateJobGroupResponseBodyJobGroupStrategyWorkingTime) GoString() string {
	return s.String()
}

func (s *CreateJobGroupResponseBodyJobGroupStrategyWorkingTime) GetBeginTime() *string {
	return s.BeginTime
}

func (s *CreateJobGroupResponseBodyJobGroupStrategyWorkingTime) GetEndTime() *string {
	return s.EndTime
}

func (s *CreateJobGroupResponseBodyJobGroupStrategyWorkingTime) SetBeginTime(v string) *CreateJobGroupResponseBodyJobGroupStrategyWorkingTime {
	s.BeginTime = &v
	return s
}

func (s *CreateJobGroupResponseBodyJobGroupStrategyWorkingTime) SetEndTime(v string) *CreateJobGroupResponseBodyJobGroupStrategyWorkingTime {
	s.EndTime = &v
	return s
}

func (s *CreateJobGroupResponseBodyJobGroupStrategyWorkingTime) Validate() error {
	return dara.Validate(s)
}
