// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyJobGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ModifyJobGroupResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ModifyJobGroupResponseBody
	GetHttpStatusCode() *int32
	SetJobGroup(v *ModifyJobGroupResponseBodyJobGroup) *ModifyJobGroupResponseBody
	GetJobGroup() *ModifyJobGroupResponseBodyJobGroup
	SetMessage(v string) *ModifyJobGroupResponseBody
	GetMessage() *string
	SetRequestId(v string) *ModifyJobGroupResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModifyJobGroupResponseBody
	GetSuccess() *bool
}

type ModifyJobGroupResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32                              `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	JobGroup       *ModifyJobGroupResponseBodyJobGroup `json:"JobGroup,omitempty" xml:"JobGroup,omitempty" type:"Struct"`
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

func (s ModifyJobGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyJobGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyJobGroupResponseBody) GetCode() *string {
	return s.Code
}

func (s *ModifyJobGroupResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ModifyJobGroupResponseBody) GetJobGroup() *ModifyJobGroupResponseBodyJobGroup {
	return s.JobGroup
}

func (s *ModifyJobGroupResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ModifyJobGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyJobGroupResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModifyJobGroupResponseBody) SetCode(v string) *ModifyJobGroupResponseBody {
	s.Code = &v
	return s
}

func (s *ModifyJobGroupResponseBody) SetHttpStatusCode(v int32) *ModifyJobGroupResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModifyJobGroupResponseBody) SetJobGroup(v *ModifyJobGroupResponseBodyJobGroup) *ModifyJobGroupResponseBody {
	s.JobGroup = v
	return s
}

func (s *ModifyJobGroupResponseBody) SetMessage(v string) *ModifyJobGroupResponseBody {
	s.Message = &v
	return s
}

func (s *ModifyJobGroupResponseBody) SetRequestId(v string) *ModifyJobGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyJobGroupResponseBody) SetSuccess(v bool) *ModifyJobGroupResponseBody {
	s.Success = &v
	return s
}

func (s *ModifyJobGroupResponseBody) Validate() error {
	if s.JobGroup != nil {
		if err := s.JobGroup.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModifyJobGroupResponseBodyJobGroup struct {
	CallingNumbers []*string `json:"CallingNumbers,omitempty" xml:"CallingNumbers,omitempty" type:"Repeated"`
	// example:
	//
	// 1578550074361
	CreationTime   *int64                                            `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	ExportProgress *ModifyJobGroupResponseBodyJobGroupExportProgress `json:"ExportProgress,omitempty" xml:"ExportProgress,omitempty" type:"Struct"`
	FlashSmsExtras *string                                           `json:"FlashSmsExtras,omitempty" xml:"FlashSmsExtras,omitempty"`
	// example:
	//
	// e37d28cb-0413-4816-85ed-fd354d025ea3
	JobDataParsingTaskId *string `json:"JobDataParsingTaskId,omitempty" xml:"JobDataParsingTaskId,omitempty"`
	JobFilePath          *string `json:"JobFilePath,omitempty" xml:"JobFilePath,omitempty"`
	JobGroupDescription  *string `json:"JobGroupDescription,omitempty" xml:"JobGroupDescription,omitempty"`
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
	ModifyTime      *string                                           `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	Priority        *string                                           `json:"Priority,omitempty" xml:"Priority,omitempty"`
	RecallStrategy  *ModifyJobGroupResponseBodyJobGroupRecallStrategy `json:"RecallStrategy,omitempty" xml:"RecallStrategy,omitempty" type:"Struct"`
	RingingDuration *int64                                            `json:"RingingDuration,omitempty" xml:"RingingDuration,omitempty"`
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
	Strategy *ModifyJobGroupResponseBodyJobGroupStrategy `json:"Strategy,omitempty" xml:"Strategy,omitempty" type:"Struct"`
}

func (s ModifyJobGroupResponseBodyJobGroup) String() string {
	return dara.Prettify(s)
}

func (s ModifyJobGroupResponseBodyJobGroup) GoString() string {
	return s.String()
}

func (s *ModifyJobGroupResponseBodyJobGroup) GetCallingNumbers() []*string {
	return s.CallingNumbers
}

func (s *ModifyJobGroupResponseBodyJobGroup) GetCreationTime() *int64 {
	return s.CreationTime
}

func (s *ModifyJobGroupResponseBodyJobGroup) GetExportProgress() *ModifyJobGroupResponseBodyJobGroupExportProgress {
	return s.ExportProgress
}

func (s *ModifyJobGroupResponseBodyJobGroup) GetFlashSmsExtras() *string {
	return s.FlashSmsExtras
}

func (s *ModifyJobGroupResponseBodyJobGroup) GetJobDataParsingTaskId() *string {
	return s.JobDataParsingTaskId
}

func (s *ModifyJobGroupResponseBodyJobGroup) GetJobFilePath() *string {
	return s.JobFilePath
}

func (s *ModifyJobGroupResponseBodyJobGroup) GetJobGroupDescription() *string {
	return s.JobGroupDescription
}

func (s *ModifyJobGroupResponseBodyJobGroup) GetJobGroupId() *string {
	return s.JobGroupId
}

func (s *ModifyJobGroupResponseBodyJobGroup) GetJobGroupName() *string {
	return s.JobGroupName
}

func (s *ModifyJobGroupResponseBodyJobGroup) GetMinConcurrency() *int64 {
	return s.MinConcurrency
}

func (s *ModifyJobGroupResponseBodyJobGroup) GetModifyTime() *string {
	return s.ModifyTime
}

func (s *ModifyJobGroupResponseBodyJobGroup) GetPriority() *string {
	return s.Priority
}

func (s *ModifyJobGroupResponseBodyJobGroup) GetRecallStrategy() *ModifyJobGroupResponseBodyJobGroupRecallStrategy {
	return s.RecallStrategy
}

func (s *ModifyJobGroupResponseBodyJobGroup) GetRingingDuration() *int64 {
	return s.RingingDuration
}

func (s *ModifyJobGroupResponseBodyJobGroup) GetScenarioId() *string {
	return s.ScenarioId
}

func (s *ModifyJobGroupResponseBodyJobGroup) GetScriptName() *string {
	return s.ScriptName
}

func (s *ModifyJobGroupResponseBodyJobGroup) GetScriptVersion() *string {
	return s.ScriptVersion
}

func (s *ModifyJobGroupResponseBodyJobGroup) GetStatus() *string {
	return s.Status
}

func (s *ModifyJobGroupResponseBodyJobGroup) GetStrategy() *ModifyJobGroupResponseBodyJobGroupStrategy {
	return s.Strategy
}

func (s *ModifyJobGroupResponseBodyJobGroup) SetCallingNumbers(v []*string) *ModifyJobGroupResponseBodyJobGroup {
	s.CallingNumbers = v
	return s
}

func (s *ModifyJobGroupResponseBodyJobGroup) SetCreationTime(v int64) *ModifyJobGroupResponseBodyJobGroup {
	s.CreationTime = &v
	return s
}

func (s *ModifyJobGroupResponseBodyJobGroup) SetExportProgress(v *ModifyJobGroupResponseBodyJobGroupExportProgress) *ModifyJobGroupResponseBodyJobGroup {
	s.ExportProgress = v
	return s
}

func (s *ModifyJobGroupResponseBodyJobGroup) SetFlashSmsExtras(v string) *ModifyJobGroupResponseBodyJobGroup {
	s.FlashSmsExtras = &v
	return s
}

func (s *ModifyJobGroupResponseBodyJobGroup) SetJobDataParsingTaskId(v string) *ModifyJobGroupResponseBodyJobGroup {
	s.JobDataParsingTaskId = &v
	return s
}

func (s *ModifyJobGroupResponseBodyJobGroup) SetJobFilePath(v string) *ModifyJobGroupResponseBodyJobGroup {
	s.JobFilePath = &v
	return s
}

func (s *ModifyJobGroupResponseBodyJobGroup) SetJobGroupDescription(v string) *ModifyJobGroupResponseBodyJobGroup {
	s.JobGroupDescription = &v
	return s
}

func (s *ModifyJobGroupResponseBodyJobGroup) SetJobGroupId(v string) *ModifyJobGroupResponseBodyJobGroup {
	s.JobGroupId = &v
	return s
}

func (s *ModifyJobGroupResponseBodyJobGroup) SetJobGroupName(v string) *ModifyJobGroupResponseBodyJobGroup {
	s.JobGroupName = &v
	return s
}

func (s *ModifyJobGroupResponseBodyJobGroup) SetMinConcurrency(v int64) *ModifyJobGroupResponseBodyJobGroup {
	s.MinConcurrency = &v
	return s
}

func (s *ModifyJobGroupResponseBodyJobGroup) SetModifyTime(v string) *ModifyJobGroupResponseBodyJobGroup {
	s.ModifyTime = &v
	return s
}

func (s *ModifyJobGroupResponseBodyJobGroup) SetPriority(v string) *ModifyJobGroupResponseBodyJobGroup {
	s.Priority = &v
	return s
}

func (s *ModifyJobGroupResponseBodyJobGroup) SetRecallStrategy(v *ModifyJobGroupResponseBodyJobGroupRecallStrategy) *ModifyJobGroupResponseBodyJobGroup {
	s.RecallStrategy = v
	return s
}

func (s *ModifyJobGroupResponseBodyJobGroup) SetRingingDuration(v int64) *ModifyJobGroupResponseBodyJobGroup {
	s.RingingDuration = &v
	return s
}

func (s *ModifyJobGroupResponseBodyJobGroup) SetScenarioId(v string) *ModifyJobGroupResponseBodyJobGroup {
	s.ScenarioId = &v
	return s
}

func (s *ModifyJobGroupResponseBodyJobGroup) SetScriptName(v string) *ModifyJobGroupResponseBodyJobGroup {
	s.ScriptName = &v
	return s
}

func (s *ModifyJobGroupResponseBodyJobGroup) SetScriptVersion(v string) *ModifyJobGroupResponseBodyJobGroup {
	s.ScriptVersion = &v
	return s
}

func (s *ModifyJobGroupResponseBodyJobGroup) SetStatus(v string) *ModifyJobGroupResponseBodyJobGroup {
	s.Status = &v
	return s
}

func (s *ModifyJobGroupResponseBodyJobGroup) SetStrategy(v *ModifyJobGroupResponseBodyJobGroupStrategy) *ModifyJobGroupResponseBodyJobGroup {
	s.Strategy = v
	return s
}

func (s *ModifyJobGroupResponseBodyJobGroup) Validate() error {
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

type ModifyJobGroupResponseBodyJobGroupExportProgress struct {
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

func (s ModifyJobGroupResponseBodyJobGroupExportProgress) String() string {
	return dara.Prettify(s)
}

func (s ModifyJobGroupResponseBodyJobGroupExportProgress) GoString() string {
	return s.String()
}

func (s *ModifyJobGroupResponseBodyJobGroupExportProgress) GetFileHttpUrl() *string {
	return s.FileHttpUrl
}

func (s *ModifyJobGroupResponseBodyJobGroupExportProgress) GetProgress() *string {
	return s.Progress
}

func (s *ModifyJobGroupResponseBodyJobGroupExportProgress) GetStatus() *string {
	return s.Status
}

func (s *ModifyJobGroupResponseBodyJobGroupExportProgress) SetFileHttpUrl(v string) *ModifyJobGroupResponseBodyJobGroupExportProgress {
	s.FileHttpUrl = &v
	return s
}

func (s *ModifyJobGroupResponseBodyJobGroupExportProgress) SetProgress(v string) *ModifyJobGroupResponseBodyJobGroupExportProgress {
	s.Progress = &v
	return s
}

func (s *ModifyJobGroupResponseBodyJobGroupExportProgress) SetStatus(v string) *ModifyJobGroupResponseBodyJobGroupExportProgress {
	s.Status = &v
	return s
}

func (s *ModifyJobGroupResponseBodyJobGroupExportProgress) Validate() error {
	return dara.Validate(s)
}

type ModifyJobGroupResponseBodyJobGroupRecallStrategy struct {
	// example:
	//
	// false
	EmptyNumberIgnore *bool `json:"EmptyNumberIgnore,omitempty" xml:"EmptyNumberIgnore,omitempty"`
	// example:
	//
	// false
	InArrearsIgnore *bool `json:"InArrearsIgnore,omitempty" xml:"InArrearsIgnore,omitempty"`
	// example:
	//
	// false
	OutOfServiceIgnore *bool `json:"OutOfServiceIgnore,omitempty" xml:"OutOfServiceIgnore,omitempty"`
}

func (s ModifyJobGroupResponseBodyJobGroupRecallStrategy) String() string {
	return dara.Prettify(s)
}

func (s ModifyJobGroupResponseBodyJobGroupRecallStrategy) GoString() string {
	return s.String()
}

func (s *ModifyJobGroupResponseBodyJobGroupRecallStrategy) GetEmptyNumberIgnore() *bool {
	return s.EmptyNumberIgnore
}

func (s *ModifyJobGroupResponseBodyJobGroupRecallStrategy) GetInArrearsIgnore() *bool {
	return s.InArrearsIgnore
}

func (s *ModifyJobGroupResponseBodyJobGroupRecallStrategy) GetOutOfServiceIgnore() *bool {
	return s.OutOfServiceIgnore
}

func (s *ModifyJobGroupResponseBodyJobGroupRecallStrategy) SetEmptyNumberIgnore(v bool) *ModifyJobGroupResponseBodyJobGroupRecallStrategy {
	s.EmptyNumberIgnore = &v
	return s
}

func (s *ModifyJobGroupResponseBodyJobGroupRecallStrategy) SetInArrearsIgnore(v bool) *ModifyJobGroupResponseBodyJobGroupRecallStrategy {
	s.InArrearsIgnore = &v
	return s
}

func (s *ModifyJobGroupResponseBodyJobGroupRecallStrategy) SetOutOfServiceIgnore(v bool) *ModifyJobGroupResponseBodyJobGroupRecallStrategy {
	s.OutOfServiceIgnore = &v
	return s
}

func (s *ModifyJobGroupResponseBodyJobGroupRecallStrategy) Validate() error {
	return dara.Validate(s)
}

type ModifyJobGroupResponseBodyJobGroupStrategy struct {
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
	// 10
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
	// f718798d-96be-40e4-bef6-317b54855708
	StrategyId   *string `json:"StrategyId,omitempty" xml:"StrategyId,omitempty"`
	StrategyName *string `json:"StrategyName,omitempty" xml:"StrategyName,omitempty"`
	// example:
	//
	// Repeatable
	Type        *string                                                  `json:"Type,omitempty" xml:"Type,omitempty"`
	WorkingTime []*ModifyJobGroupResponseBodyJobGroupStrategyWorkingTime `json:"WorkingTime,omitempty" xml:"WorkingTime,omitempty" type:"Repeated"`
}

func (s ModifyJobGroupResponseBodyJobGroupStrategy) String() string {
	return dara.Prettify(s)
}

func (s ModifyJobGroupResponseBodyJobGroupStrategy) GoString() string {
	return s.String()
}

func (s *ModifyJobGroupResponseBodyJobGroupStrategy) GetCustomized() *string {
	return s.Customized
}

func (s *ModifyJobGroupResponseBodyJobGroupStrategy) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ModifyJobGroupResponseBodyJobGroupStrategy) GetFollowUpStrategy() *string {
	return s.FollowUpStrategy
}

func (s *ModifyJobGroupResponseBodyJobGroupStrategy) GetIsTemplate() *bool {
	return s.IsTemplate
}

func (s *ModifyJobGroupResponseBodyJobGroupStrategy) GetMaxAttemptsPerDay() *int32 {
	return s.MaxAttemptsPerDay
}

func (s *ModifyJobGroupResponseBodyJobGroupStrategy) GetMinAttemptInterval() *int32 {
	return s.MinAttemptInterval
}

func (s *ModifyJobGroupResponseBodyJobGroupStrategy) GetRepeatBy() *string {
	return s.RepeatBy
}

func (s *ModifyJobGroupResponseBodyJobGroupStrategy) GetRepeatDays() []*string {
	return s.RepeatDays
}

func (s *ModifyJobGroupResponseBodyJobGroupStrategy) GetRoutingStrategy() *string {
	return s.RoutingStrategy
}

func (s *ModifyJobGroupResponseBodyJobGroupStrategy) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ModifyJobGroupResponseBodyJobGroupStrategy) GetStrategyDescription() *string {
	return s.StrategyDescription
}

func (s *ModifyJobGroupResponseBodyJobGroupStrategy) GetStrategyId() *string {
	return s.StrategyId
}

func (s *ModifyJobGroupResponseBodyJobGroupStrategy) GetStrategyName() *string {
	return s.StrategyName
}

func (s *ModifyJobGroupResponseBodyJobGroupStrategy) GetType() *string {
	return s.Type
}

func (s *ModifyJobGroupResponseBodyJobGroupStrategy) GetWorkingTime() []*ModifyJobGroupResponseBodyJobGroupStrategyWorkingTime {
	return s.WorkingTime
}

func (s *ModifyJobGroupResponseBodyJobGroupStrategy) SetCustomized(v string) *ModifyJobGroupResponseBodyJobGroupStrategy {
	s.Customized = &v
	return s
}

func (s *ModifyJobGroupResponseBodyJobGroupStrategy) SetEndTime(v int64) *ModifyJobGroupResponseBodyJobGroupStrategy {
	s.EndTime = &v
	return s
}

func (s *ModifyJobGroupResponseBodyJobGroupStrategy) SetFollowUpStrategy(v string) *ModifyJobGroupResponseBodyJobGroupStrategy {
	s.FollowUpStrategy = &v
	return s
}

func (s *ModifyJobGroupResponseBodyJobGroupStrategy) SetIsTemplate(v bool) *ModifyJobGroupResponseBodyJobGroupStrategy {
	s.IsTemplate = &v
	return s
}

func (s *ModifyJobGroupResponseBodyJobGroupStrategy) SetMaxAttemptsPerDay(v int32) *ModifyJobGroupResponseBodyJobGroupStrategy {
	s.MaxAttemptsPerDay = &v
	return s
}

func (s *ModifyJobGroupResponseBodyJobGroupStrategy) SetMinAttemptInterval(v int32) *ModifyJobGroupResponseBodyJobGroupStrategy {
	s.MinAttemptInterval = &v
	return s
}

func (s *ModifyJobGroupResponseBodyJobGroupStrategy) SetRepeatBy(v string) *ModifyJobGroupResponseBodyJobGroupStrategy {
	s.RepeatBy = &v
	return s
}

func (s *ModifyJobGroupResponseBodyJobGroupStrategy) SetRepeatDays(v []*string) *ModifyJobGroupResponseBodyJobGroupStrategy {
	s.RepeatDays = v
	return s
}

func (s *ModifyJobGroupResponseBodyJobGroupStrategy) SetRoutingStrategy(v string) *ModifyJobGroupResponseBodyJobGroupStrategy {
	s.RoutingStrategy = &v
	return s
}

func (s *ModifyJobGroupResponseBodyJobGroupStrategy) SetStartTime(v int64) *ModifyJobGroupResponseBodyJobGroupStrategy {
	s.StartTime = &v
	return s
}

func (s *ModifyJobGroupResponseBodyJobGroupStrategy) SetStrategyDescription(v string) *ModifyJobGroupResponseBodyJobGroupStrategy {
	s.StrategyDescription = &v
	return s
}

func (s *ModifyJobGroupResponseBodyJobGroupStrategy) SetStrategyId(v string) *ModifyJobGroupResponseBodyJobGroupStrategy {
	s.StrategyId = &v
	return s
}

func (s *ModifyJobGroupResponseBodyJobGroupStrategy) SetStrategyName(v string) *ModifyJobGroupResponseBodyJobGroupStrategy {
	s.StrategyName = &v
	return s
}

func (s *ModifyJobGroupResponseBodyJobGroupStrategy) SetType(v string) *ModifyJobGroupResponseBodyJobGroupStrategy {
	s.Type = &v
	return s
}

func (s *ModifyJobGroupResponseBodyJobGroupStrategy) SetWorkingTime(v []*ModifyJobGroupResponseBodyJobGroupStrategyWorkingTime) *ModifyJobGroupResponseBodyJobGroupStrategy {
	s.WorkingTime = v
	return s
}

func (s *ModifyJobGroupResponseBodyJobGroupStrategy) Validate() error {
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

type ModifyJobGroupResponseBodyJobGroupStrategyWorkingTime struct {
	// example:
	//
	// 09:00:00
	BeginTime *string `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	// example:
	//
	// 12:00:00
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
}

func (s ModifyJobGroupResponseBodyJobGroupStrategyWorkingTime) String() string {
	return dara.Prettify(s)
}

func (s ModifyJobGroupResponseBodyJobGroupStrategyWorkingTime) GoString() string {
	return s.String()
}

func (s *ModifyJobGroupResponseBodyJobGroupStrategyWorkingTime) GetBeginTime() *string {
	return s.BeginTime
}

func (s *ModifyJobGroupResponseBodyJobGroupStrategyWorkingTime) GetEndTime() *string {
	return s.EndTime
}

func (s *ModifyJobGroupResponseBodyJobGroupStrategyWorkingTime) SetBeginTime(v string) *ModifyJobGroupResponseBodyJobGroupStrategyWorkingTime {
	s.BeginTime = &v
	return s
}

func (s *ModifyJobGroupResponseBodyJobGroupStrategyWorkingTime) SetEndTime(v string) *ModifyJobGroupResponseBodyJobGroupStrategyWorkingTime {
	s.EndTime = &v
	return s
}

func (s *ModifyJobGroupResponseBodyJobGroupStrategyWorkingTime) Validate() error {
	return dara.Validate(s)
}
