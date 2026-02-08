// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateBatchJobsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBatchJob(v *CreateBatchJobsResponseBodyBatchJob) *CreateBatchJobsResponseBody
	GetBatchJob() *CreateBatchJobsResponseBodyBatchJob
	SetCode(v string) *CreateBatchJobsResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *CreateBatchJobsResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreateBatchJobsResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateBatchJobsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateBatchJobsResponseBody
	GetSuccess() *bool
}

type CreateBatchJobsResponseBody struct {
	// Job group information
	BatchJob *CreateBatchJobsResponseBodyBatchJob `json:"BatchJob,omitempty" xml:"BatchJob,omitempty" type:"Struct"`
	// API status code
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// HTTP status code
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Message
	//
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID
	//
	// example:
	//
	// 254EB995-DEDF-48A4-9101-9CA5B72FFBCC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the operation succeeded
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateBatchJobsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateBatchJobsResponseBody) GoString() string {
	return s.String()
}

func (s *CreateBatchJobsResponseBody) GetBatchJob() *CreateBatchJobsResponseBodyBatchJob {
	return s.BatchJob
}

func (s *CreateBatchJobsResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateBatchJobsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateBatchJobsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateBatchJobsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateBatchJobsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateBatchJobsResponseBody) SetBatchJob(v *CreateBatchJobsResponseBodyBatchJob) *CreateBatchJobsResponseBody {
	s.BatchJob = v
	return s
}

func (s *CreateBatchJobsResponseBody) SetCode(v string) *CreateBatchJobsResponseBody {
	s.Code = &v
	return s
}

func (s *CreateBatchJobsResponseBody) SetHttpStatusCode(v int32) *CreateBatchJobsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateBatchJobsResponseBody) SetMessage(v string) *CreateBatchJobsResponseBody {
	s.Message = &v
	return s
}

func (s *CreateBatchJobsResponseBody) SetRequestId(v string) *CreateBatchJobsResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateBatchJobsResponseBody) SetSuccess(v bool) *CreateBatchJobsResponseBody {
	s.Success = &v
	return s
}

func (s *CreateBatchJobsResponseBody) Validate() error {
	if s.BatchJob != nil {
		if err := s.BatchJob.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateBatchJobsResponseBodyBatchJob struct {
	// Job group ID
	//
	// example:
	//
	// 5a7e8b09-baf9-4cab-b540-c971f47a7146
	BatchJobId *string `json:"BatchJobId,omitempty" xml:"BatchJobId,omitempty"`
	// List of calling numbers
	CallingNumbers []*string `json:"CallingNumbers,omitempty" xml:"CallingNumbers,omitempty" type:"Repeated"`
	// Creation Time of the job group
	//
	// example:
	//
	// 1579068424000
	CreationTime *int64 `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// Key of the uploaded Excel file
	//
	// example:
	//
	// 52e80b02-0126-4556-a1e6-ef5b3747ed53/a9a3ddc7-d7d7-48cd-82b5-b31bb5510e71_2a66f8ad-dfbb-4980-9b84-439171295a11.xlsx
	JobFilePath *string `json:"JobFilePath,omitempty" xml:"JobFilePath,omitempty"`
	// Job group description
	//
	// example:
	//
	// 第一个批量作业
	JobGroupDescription *string `json:"JobGroupDescription,omitempty" xml:"JobGroupDescription,omitempty"`
	// Job group name
	//
	// example:
	//
	// 批量作业01
	JobGroupName *string `json:"JobGroupName,omitempty" xml:"JobGroupName,omitempty"`
	// Scenario ID of the job group
	//
	// example:
	//
	// 6cea9bed-63e6-439e-ae4c-b3333efff53d
	ScenarioId *string `json:"ScenarioId,omitempty" xml:"ScenarioId,omitempty"`
	// Job execution policy
	Strategy *CreateBatchJobsResponseBodyBatchJobStrategy `json:"Strategy,omitempty" xml:"Strategy,omitempty" type:"Struct"`
}

func (s CreateBatchJobsResponseBodyBatchJob) String() string {
	return dara.Prettify(s)
}

func (s CreateBatchJobsResponseBodyBatchJob) GoString() string {
	return s.String()
}

func (s *CreateBatchJobsResponseBodyBatchJob) GetBatchJobId() *string {
	return s.BatchJobId
}

func (s *CreateBatchJobsResponseBodyBatchJob) GetCallingNumbers() []*string {
	return s.CallingNumbers
}

func (s *CreateBatchJobsResponseBodyBatchJob) GetCreationTime() *int64 {
	return s.CreationTime
}

func (s *CreateBatchJobsResponseBodyBatchJob) GetJobFilePath() *string {
	return s.JobFilePath
}

func (s *CreateBatchJobsResponseBodyBatchJob) GetJobGroupDescription() *string {
	return s.JobGroupDescription
}

func (s *CreateBatchJobsResponseBodyBatchJob) GetJobGroupName() *string {
	return s.JobGroupName
}

func (s *CreateBatchJobsResponseBodyBatchJob) GetScenarioId() *string {
	return s.ScenarioId
}

func (s *CreateBatchJobsResponseBodyBatchJob) GetStrategy() *CreateBatchJobsResponseBodyBatchJobStrategy {
	return s.Strategy
}

func (s *CreateBatchJobsResponseBodyBatchJob) SetBatchJobId(v string) *CreateBatchJobsResponseBodyBatchJob {
	s.BatchJobId = &v
	return s
}

func (s *CreateBatchJobsResponseBodyBatchJob) SetCallingNumbers(v []*string) *CreateBatchJobsResponseBodyBatchJob {
	s.CallingNumbers = v
	return s
}

func (s *CreateBatchJobsResponseBodyBatchJob) SetCreationTime(v int64) *CreateBatchJobsResponseBodyBatchJob {
	s.CreationTime = &v
	return s
}

func (s *CreateBatchJobsResponseBodyBatchJob) SetJobFilePath(v string) *CreateBatchJobsResponseBodyBatchJob {
	s.JobFilePath = &v
	return s
}

func (s *CreateBatchJobsResponseBodyBatchJob) SetJobGroupDescription(v string) *CreateBatchJobsResponseBodyBatchJob {
	s.JobGroupDescription = &v
	return s
}

func (s *CreateBatchJobsResponseBodyBatchJob) SetJobGroupName(v string) *CreateBatchJobsResponseBodyBatchJob {
	s.JobGroupName = &v
	return s
}

func (s *CreateBatchJobsResponseBodyBatchJob) SetScenarioId(v string) *CreateBatchJobsResponseBodyBatchJob {
	s.ScenarioId = &v
	return s
}

func (s *CreateBatchJobsResponseBodyBatchJob) SetStrategy(v *CreateBatchJobsResponseBodyBatchJobStrategy) *CreateBatchJobsResponseBodyBatchJob {
	s.Strategy = v
	return s
}

func (s *CreateBatchJobsResponseBodyBatchJob) Validate() error {
	if s.Strategy != nil {
		if err := s.Strategy.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateBatchJobsResponseBodyBatchJobStrategy struct {
	// Custom policy information
	//
	// example:
	//
	// {}
	Customized *string `json:"Customized,omitempty" xml:"Customized,omitempty"`
	// Overall end time of the schedule policy
	//
	// example:
	//
	// 2209702074000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// Follow-up handling policy for outbound calls not completed in the previous epoch.
	//
	// example:
	//
	// CONTINUE
	FollowUpStrategy *string `json:"FollowUpStrategy,omitempty" xml:"FollowUpStrategy,omitempty"`
	// Indicates whether it is a template
	//
	// example:
	//
	// false
	IsTemplate *bool `json:"IsTemplate,omitempty" xml:"IsTemplate,omitempty"`
	// Maximum number of redial attempts after a failed call per day
	//
	// example:
	//
	// 3
	MaxAttemptsPerDay *int32 `json:"MaxAttemptsPerDay,omitempty" xml:"MaxAttemptsPerDay,omitempty"`
	// Minimum retry interval
	//
	// example:
	//
	// 10
	MinAttemptInterval *int32 `json:"MinAttemptInterval,omitempty" xml:"MinAttemptInterval,omitempty"`
	// Repetition mode: Once for no repetition, Day for daily repetition, Week for weekly repetition, Month for monthly repetition
	//
	// example:
	//
	// Once
	RepeatBy *string `json:"RepeatBy,omitempty" xml:"RepeatBy,omitempty"`
	// List of days for repeated execution.
	RepeatDays []*string `json:"RepeatDays,omitempty" xml:"RepeatDays,omitempty" type:"Repeated"`
	// Calling number recording policy
	//
	// example:
	//
	// LocalFirst
	RoutingStrategy *string `json:"RoutingStrategy,omitempty" xml:"RoutingStrategy,omitempty"`
	// Start time for job group execution
	//
	// example:
	//
	// 1578550074000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// Policy description
	//
	// example:
	//
	// 作业执行策略
	StrategyDescription *string `json:"StrategyDescription,omitempty" xml:"StrategyDescription,omitempty"`
	// Policy ID
	//
	// example:
	//
	// f718798d-96be-40e4-bef6-317b54855708
	StrategyId *string `json:"StrategyId,omitempty" xml:"StrategyId,omitempty"`
	// Policy name
	//
	// example:
	//
	// 策略
	StrategyName *string `json:"StrategyName,omitempty" xml:"StrategyName,omitempty"`
	// Policy Type
	//
	// example:
	//
	// Repeatable
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// Execution time segments within a day for the job
	WorkingTime []*CreateBatchJobsResponseBodyBatchJobStrategyWorkingTime `json:"WorkingTime,omitempty" xml:"WorkingTime,omitempty" type:"Repeated"`
}

func (s CreateBatchJobsResponseBodyBatchJobStrategy) String() string {
	return dara.Prettify(s)
}

func (s CreateBatchJobsResponseBodyBatchJobStrategy) GoString() string {
	return s.String()
}

func (s *CreateBatchJobsResponseBodyBatchJobStrategy) GetCustomized() *string {
	return s.Customized
}

func (s *CreateBatchJobsResponseBodyBatchJobStrategy) GetEndTime() *int64 {
	return s.EndTime
}

func (s *CreateBatchJobsResponseBodyBatchJobStrategy) GetFollowUpStrategy() *string {
	return s.FollowUpStrategy
}

func (s *CreateBatchJobsResponseBodyBatchJobStrategy) GetIsTemplate() *bool {
	return s.IsTemplate
}

func (s *CreateBatchJobsResponseBodyBatchJobStrategy) GetMaxAttemptsPerDay() *int32 {
	return s.MaxAttemptsPerDay
}

func (s *CreateBatchJobsResponseBodyBatchJobStrategy) GetMinAttemptInterval() *int32 {
	return s.MinAttemptInterval
}

func (s *CreateBatchJobsResponseBodyBatchJobStrategy) GetRepeatBy() *string {
	return s.RepeatBy
}

func (s *CreateBatchJobsResponseBodyBatchJobStrategy) GetRepeatDays() []*string {
	return s.RepeatDays
}

func (s *CreateBatchJobsResponseBodyBatchJobStrategy) GetRoutingStrategy() *string {
	return s.RoutingStrategy
}

func (s *CreateBatchJobsResponseBodyBatchJobStrategy) GetStartTime() *int64 {
	return s.StartTime
}

func (s *CreateBatchJobsResponseBodyBatchJobStrategy) GetStrategyDescription() *string {
	return s.StrategyDescription
}

func (s *CreateBatchJobsResponseBodyBatchJobStrategy) GetStrategyId() *string {
	return s.StrategyId
}

func (s *CreateBatchJobsResponseBodyBatchJobStrategy) GetStrategyName() *string {
	return s.StrategyName
}

func (s *CreateBatchJobsResponseBodyBatchJobStrategy) GetType() *string {
	return s.Type
}

func (s *CreateBatchJobsResponseBodyBatchJobStrategy) GetWorkingTime() []*CreateBatchJobsResponseBodyBatchJobStrategyWorkingTime {
	return s.WorkingTime
}

func (s *CreateBatchJobsResponseBodyBatchJobStrategy) SetCustomized(v string) *CreateBatchJobsResponseBodyBatchJobStrategy {
	s.Customized = &v
	return s
}

func (s *CreateBatchJobsResponseBodyBatchJobStrategy) SetEndTime(v int64) *CreateBatchJobsResponseBodyBatchJobStrategy {
	s.EndTime = &v
	return s
}

func (s *CreateBatchJobsResponseBodyBatchJobStrategy) SetFollowUpStrategy(v string) *CreateBatchJobsResponseBodyBatchJobStrategy {
	s.FollowUpStrategy = &v
	return s
}

func (s *CreateBatchJobsResponseBodyBatchJobStrategy) SetIsTemplate(v bool) *CreateBatchJobsResponseBodyBatchJobStrategy {
	s.IsTemplate = &v
	return s
}

func (s *CreateBatchJobsResponseBodyBatchJobStrategy) SetMaxAttemptsPerDay(v int32) *CreateBatchJobsResponseBodyBatchJobStrategy {
	s.MaxAttemptsPerDay = &v
	return s
}

func (s *CreateBatchJobsResponseBodyBatchJobStrategy) SetMinAttemptInterval(v int32) *CreateBatchJobsResponseBodyBatchJobStrategy {
	s.MinAttemptInterval = &v
	return s
}

func (s *CreateBatchJobsResponseBodyBatchJobStrategy) SetRepeatBy(v string) *CreateBatchJobsResponseBodyBatchJobStrategy {
	s.RepeatBy = &v
	return s
}

func (s *CreateBatchJobsResponseBodyBatchJobStrategy) SetRepeatDays(v []*string) *CreateBatchJobsResponseBodyBatchJobStrategy {
	s.RepeatDays = v
	return s
}

func (s *CreateBatchJobsResponseBodyBatchJobStrategy) SetRoutingStrategy(v string) *CreateBatchJobsResponseBodyBatchJobStrategy {
	s.RoutingStrategy = &v
	return s
}

func (s *CreateBatchJobsResponseBodyBatchJobStrategy) SetStartTime(v int64) *CreateBatchJobsResponseBodyBatchJobStrategy {
	s.StartTime = &v
	return s
}

func (s *CreateBatchJobsResponseBodyBatchJobStrategy) SetStrategyDescription(v string) *CreateBatchJobsResponseBodyBatchJobStrategy {
	s.StrategyDescription = &v
	return s
}

func (s *CreateBatchJobsResponseBodyBatchJobStrategy) SetStrategyId(v string) *CreateBatchJobsResponseBodyBatchJobStrategy {
	s.StrategyId = &v
	return s
}

func (s *CreateBatchJobsResponseBodyBatchJobStrategy) SetStrategyName(v string) *CreateBatchJobsResponseBodyBatchJobStrategy {
	s.StrategyName = &v
	return s
}

func (s *CreateBatchJobsResponseBodyBatchJobStrategy) SetType(v string) *CreateBatchJobsResponseBodyBatchJobStrategy {
	s.Type = &v
	return s
}

func (s *CreateBatchJobsResponseBodyBatchJobStrategy) SetWorkingTime(v []*CreateBatchJobsResponseBodyBatchJobStrategyWorkingTime) *CreateBatchJobsResponseBodyBatchJobStrategy {
	s.WorkingTime = v
	return s
}

func (s *CreateBatchJobsResponseBodyBatchJobStrategy) Validate() error {
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

type CreateBatchJobsResponseBodyBatchJobStrategyWorkingTime struct {
	// Start Time
	//
	// example:
	//
	// 09:00:00
	BeginTime *string `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	// End Time
	//
	// example:
	//
	// 12:00:00
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
}

func (s CreateBatchJobsResponseBodyBatchJobStrategyWorkingTime) String() string {
	return dara.Prettify(s)
}

func (s CreateBatchJobsResponseBodyBatchJobStrategyWorkingTime) GoString() string {
	return s.String()
}

func (s *CreateBatchJobsResponseBodyBatchJobStrategyWorkingTime) GetBeginTime() *string {
	return s.BeginTime
}

func (s *CreateBatchJobsResponseBodyBatchJobStrategyWorkingTime) GetEndTime() *string {
	return s.EndTime
}

func (s *CreateBatchJobsResponseBodyBatchJobStrategyWorkingTime) SetBeginTime(v string) *CreateBatchJobsResponseBodyBatchJobStrategyWorkingTime {
	s.BeginTime = &v
	return s
}

func (s *CreateBatchJobsResponseBodyBatchJobStrategyWorkingTime) SetEndTime(v string) *CreateBatchJobsResponseBodyBatchJobStrategyWorkingTime {
	s.EndTime = &v
	return s
}

func (s *CreateBatchJobsResponseBodyBatchJobStrategyWorkingTime) Validate() error {
	return dara.Validate(s)
}
