// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyBatchJobsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ModifyBatchJobsResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ModifyBatchJobsResponseBody
	GetHttpStatusCode() *int32
	SetJobGroup(v *ModifyBatchJobsResponseBodyJobGroup) *ModifyBatchJobsResponseBody
	GetJobGroup() *ModifyBatchJobsResponseBodyJobGroup
	SetMessage(v string) *ModifyBatchJobsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ModifyBatchJobsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModifyBatchJobsResponseBody
	GetSuccess() *bool
}

type ModifyBatchJobsResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32                               `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	JobGroup       *ModifyBatchJobsResponseBodyJobGroup `json:"JobGroup,omitempty" xml:"JobGroup,omitempty" type:"Struct"`
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

func (s ModifyBatchJobsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyBatchJobsResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyBatchJobsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ModifyBatchJobsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ModifyBatchJobsResponseBody) GetJobGroup() *ModifyBatchJobsResponseBodyJobGroup {
	return s.JobGroup
}

func (s *ModifyBatchJobsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ModifyBatchJobsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyBatchJobsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModifyBatchJobsResponseBody) SetCode(v string) *ModifyBatchJobsResponseBody {
	s.Code = &v
	return s
}

func (s *ModifyBatchJobsResponseBody) SetHttpStatusCode(v int32) *ModifyBatchJobsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModifyBatchJobsResponseBody) SetJobGroup(v *ModifyBatchJobsResponseBodyJobGroup) *ModifyBatchJobsResponseBody {
	s.JobGroup = v
	return s
}

func (s *ModifyBatchJobsResponseBody) SetMessage(v string) *ModifyBatchJobsResponseBody {
	s.Message = &v
	return s
}

func (s *ModifyBatchJobsResponseBody) SetRequestId(v string) *ModifyBatchJobsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyBatchJobsResponseBody) SetSuccess(v bool) *ModifyBatchJobsResponseBody {
	s.Success = &v
	return s
}

func (s *ModifyBatchJobsResponseBody) Validate() error {
	if s.JobGroup != nil {
		if err := s.JobGroup.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModifyBatchJobsResponseBodyJobGroup struct {
	CallingNumbers []*string `json:"CallingNumbers,omitempty" xml:"CallingNumbers,omitempty" type:"Repeated"`
	// example:
	//
	// 1579068424000
	CreationTime *int64 `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// example:
	//
	// 52e80b02-0126-4556-a1e6-ef5b3747ed53/a9a3ddc7-d7d7-48cd-82b5-b31bb5510e71_2a66f8ad-dfbb-4980-9b84-439171295a11.xlsx
	JobFilePath         *string `json:"JobFilePath,omitempty" xml:"JobFilePath,omitempty"`
	JobGroupDescription *string `json:"JobGroupDescription,omitempty" xml:"JobGroupDescription,omitempty"`
	// example:
	//
	// 5a7e8b09-baf9-4cab-b540-c971f47a7146
	JobGroupId   *string `json:"JobGroupId,omitempty" xml:"JobGroupId,omitempty"`
	JobGroupName *string `json:"JobGroupName,omitempty" xml:"JobGroupName,omitempty"`
	// example:
	//
	// 6cea9bed-63e6-439e-ae4c-b3333efff53d
	ScenarioId *string                                      `json:"ScenarioId,omitempty" xml:"ScenarioId,omitempty"`
	Strategy   *ModifyBatchJobsResponseBodyJobGroupStrategy `json:"Strategy,omitempty" xml:"Strategy,omitempty" type:"Struct"`
}

func (s ModifyBatchJobsResponseBodyJobGroup) String() string {
	return dara.Prettify(s)
}

func (s ModifyBatchJobsResponseBodyJobGroup) GoString() string {
	return s.String()
}

func (s *ModifyBatchJobsResponseBodyJobGroup) GetCallingNumbers() []*string {
	return s.CallingNumbers
}

func (s *ModifyBatchJobsResponseBodyJobGroup) GetCreationTime() *int64 {
	return s.CreationTime
}

func (s *ModifyBatchJobsResponseBodyJobGroup) GetJobFilePath() *string {
	return s.JobFilePath
}

func (s *ModifyBatchJobsResponseBodyJobGroup) GetJobGroupDescription() *string {
	return s.JobGroupDescription
}

func (s *ModifyBatchJobsResponseBodyJobGroup) GetJobGroupId() *string {
	return s.JobGroupId
}

func (s *ModifyBatchJobsResponseBodyJobGroup) GetJobGroupName() *string {
	return s.JobGroupName
}

func (s *ModifyBatchJobsResponseBodyJobGroup) GetScenarioId() *string {
	return s.ScenarioId
}

func (s *ModifyBatchJobsResponseBodyJobGroup) GetStrategy() *ModifyBatchJobsResponseBodyJobGroupStrategy {
	return s.Strategy
}

func (s *ModifyBatchJobsResponseBodyJobGroup) SetCallingNumbers(v []*string) *ModifyBatchJobsResponseBodyJobGroup {
	s.CallingNumbers = v
	return s
}

func (s *ModifyBatchJobsResponseBodyJobGroup) SetCreationTime(v int64) *ModifyBatchJobsResponseBodyJobGroup {
	s.CreationTime = &v
	return s
}

func (s *ModifyBatchJobsResponseBodyJobGroup) SetJobFilePath(v string) *ModifyBatchJobsResponseBodyJobGroup {
	s.JobFilePath = &v
	return s
}

func (s *ModifyBatchJobsResponseBodyJobGroup) SetJobGroupDescription(v string) *ModifyBatchJobsResponseBodyJobGroup {
	s.JobGroupDescription = &v
	return s
}

func (s *ModifyBatchJobsResponseBodyJobGroup) SetJobGroupId(v string) *ModifyBatchJobsResponseBodyJobGroup {
	s.JobGroupId = &v
	return s
}

func (s *ModifyBatchJobsResponseBodyJobGroup) SetJobGroupName(v string) *ModifyBatchJobsResponseBodyJobGroup {
	s.JobGroupName = &v
	return s
}

func (s *ModifyBatchJobsResponseBodyJobGroup) SetScenarioId(v string) *ModifyBatchJobsResponseBodyJobGroup {
	s.ScenarioId = &v
	return s
}

func (s *ModifyBatchJobsResponseBodyJobGroup) SetStrategy(v *ModifyBatchJobsResponseBodyJobGroupStrategy) *ModifyBatchJobsResponseBodyJobGroup {
	s.Strategy = v
	return s
}

func (s *ModifyBatchJobsResponseBodyJobGroup) Validate() error {
	if s.Strategy != nil {
		if err := s.Strategy.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModifyBatchJobsResponseBodyJobGroupStrategy struct {
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
	Type        *string                                                   `json:"Type,omitempty" xml:"Type,omitempty"`
	WorkingTime []*ModifyBatchJobsResponseBodyJobGroupStrategyWorkingTime `json:"WorkingTime,omitempty" xml:"WorkingTime,omitempty" type:"Repeated"`
}

func (s ModifyBatchJobsResponseBodyJobGroupStrategy) String() string {
	return dara.Prettify(s)
}

func (s ModifyBatchJobsResponseBodyJobGroupStrategy) GoString() string {
	return s.String()
}

func (s *ModifyBatchJobsResponseBodyJobGroupStrategy) GetCustomized() *string {
	return s.Customized
}

func (s *ModifyBatchJobsResponseBodyJobGroupStrategy) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ModifyBatchJobsResponseBodyJobGroupStrategy) GetFollowUpStrategy() *string {
	return s.FollowUpStrategy
}

func (s *ModifyBatchJobsResponseBodyJobGroupStrategy) GetIsTemplate() *bool {
	return s.IsTemplate
}

func (s *ModifyBatchJobsResponseBodyJobGroupStrategy) GetMaxAttemptsPerDay() *int32 {
	return s.MaxAttemptsPerDay
}

func (s *ModifyBatchJobsResponseBodyJobGroupStrategy) GetMinAttemptInterval() *int32 {
	return s.MinAttemptInterval
}

func (s *ModifyBatchJobsResponseBodyJobGroupStrategy) GetRepeatBy() *string {
	return s.RepeatBy
}

func (s *ModifyBatchJobsResponseBodyJobGroupStrategy) GetRepeatDays() []*string {
	return s.RepeatDays
}

func (s *ModifyBatchJobsResponseBodyJobGroupStrategy) GetRoutingStrategy() *string {
	return s.RoutingStrategy
}

func (s *ModifyBatchJobsResponseBodyJobGroupStrategy) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ModifyBatchJobsResponseBodyJobGroupStrategy) GetStrategyDescription() *string {
	return s.StrategyDescription
}

func (s *ModifyBatchJobsResponseBodyJobGroupStrategy) GetStrategyId() *string {
	return s.StrategyId
}

func (s *ModifyBatchJobsResponseBodyJobGroupStrategy) GetStrategyName() *string {
	return s.StrategyName
}

func (s *ModifyBatchJobsResponseBodyJobGroupStrategy) GetType() *string {
	return s.Type
}

func (s *ModifyBatchJobsResponseBodyJobGroupStrategy) GetWorkingTime() []*ModifyBatchJobsResponseBodyJobGroupStrategyWorkingTime {
	return s.WorkingTime
}

func (s *ModifyBatchJobsResponseBodyJobGroupStrategy) SetCustomized(v string) *ModifyBatchJobsResponseBodyJobGroupStrategy {
	s.Customized = &v
	return s
}

func (s *ModifyBatchJobsResponseBodyJobGroupStrategy) SetEndTime(v int64) *ModifyBatchJobsResponseBodyJobGroupStrategy {
	s.EndTime = &v
	return s
}

func (s *ModifyBatchJobsResponseBodyJobGroupStrategy) SetFollowUpStrategy(v string) *ModifyBatchJobsResponseBodyJobGroupStrategy {
	s.FollowUpStrategy = &v
	return s
}

func (s *ModifyBatchJobsResponseBodyJobGroupStrategy) SetIsTemplate(v bool) *ModifyBatchJobsResponseBodyJobGroupStrategy {
	s.IsTemplate = &v
	return s
}

func (s *ModifyBatchJobsResponseBodyJobGroupStrategy) SetMaxAttemptsPerDay(v int32) *ModifyBatchJobsResponseBodyJobGroupStrategy {
	s.MaxAttemptsPerDay = &v
	return s
}

func (s *ModifyBatchJobsResponseBodyJobGroupStrategy) SetMinAttemptInterval(v int32) *ModifyBatchJobsResponseBodyJobGroupStrategy {
	s.MinAttemptInterval = &v
	return s
}

func (s *ModifyBatchJobsResponseBodyJobGroupStrategy) SetRepeatBy(v string) *ModifyBatchJobsResponseBodyJobGroupStrategy {
	s.RepeatBy = &v
	return s
}

func (s *ModifyBatchJobsResponseBodyJobGroupStrategy) SetRepeatDays(v []*string) *ModifyBatchJobsResponseBodyJobGroupStrategy {
	s.RepeatDays = v
	return s
}

func (s *ModifyBatchJobsResponseBodyJobGroupStrategy) SetRoutingStrategy(v string) *ModifyBatchJobsResponseBodyJobGroupStrategy {
	s.RoutingStrategy = &v
	return s
}

func (s *ModifyBatchJobsResponseBodyJobGroupStrategy) SetStartTime(v int64) *ModifyBatchJobsResponseBodyJobGroupStrategy {
	s.StartTime = &v
	return s
}

func (s *ModifyBatchJobsResponseBodyJobGroupStrategy) SetStrategyDescription(v string) *ModifyBatchJobsResponseBodyJobGroupStrategy {
	s.StrategyDescription = &v
	return s
}

func (s *ModifyBatchJobsResponseBodyJobGroupStrategy) SetStrategyId(v string) *ModifyBatchJobsResponseBodyJobGroupStrategy {
	s.StrategyId = &v
	return s
}

func (s *ModifyBatchJobsResponseBodyJobGroupStrategy) SetStrategyName(v string) *ModifyBatchJobsResponseBodyJobGroupStrategy {
	s.StrategyName = &v
	return s
}

func (s *ModifyBatchJobsResponseBodyJobGroupStrategy) SetType(v string) *ModifyBatchJobsResponseBodyJobGroupStrategy {
	s.Type = &v
	return s
}

func (s *ModifyBatchJobsResponseBodyJobGroupStrategy) SetWorkingTime(v []*ModifyBatchJobsResponseBodyJobGroupStrategyWorkingTime) *ModifyBatchJobsResponseBodyJobGroupStrategy {
	s.WorkingTime = v
	return s
}

func (s *ModifyBatchJobsResponseBodyJobGroupStrategy) Validate() error {
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

type ModifyBatchJobsResponseBodyJobGroupStrategyWorkingTime struct {
	// example:
	//
	// 1581937093000
	BeginTime *string `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	// example:
	//
	// 1581997093000
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
}

func (s ModifyBatchJobsResponseBodyJobGroupStrategyWorkingTime) String() string {
	return dara.Prettify(s)
}

func (s ModifyBatchJobsResponseBodyJobGroupStrategyWorkingTime) GoString() string {
	return s.String()
}

func (s *ModifyBatchJobsResponseBodyJobGroupStrategyWorkingTime) GetBeginTime() *string {
	return s.BeginTime
}

func (s *ModifyBatchJobsResponseBodyJobGroupStrategyWorkingTime) GetEndTime() *string {
	return s.EndTime
}

func (s *ModifyBatchJobsResponseBodyJobGroupStrategyWorkingTime) SetBeginTime(v string) *ModifyBatchJobsResponseBodyJobGroupStrategyWorkingTime {
	s.BeginTime = &v
	return s
}

func (s *ModifyBatchJobsResponseBodyJobGroupStrategyWorkingTime) SetEndTime(v string) *ModifyBatchJobsResponseBodyJobGroupStrategyWorkingTime {
	s.EndTime = &v
	return s
}

func (s *ModifyBatchJobsResponseBodyJobGroupStrategyWorkingTime) Validate() error {
	return dara.Validate(s)
}
