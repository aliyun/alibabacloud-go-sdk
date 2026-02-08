// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssignJobsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCallingNumber(v []*string) *AssignJobsRequest
	GetCallingNumber() []*string
	SetInstanceId(v string) *AssignJobsRequest
	GetInstanceId() *string
	SetIsAsynchrony(v bool) *AssignJobsRequest
	GetIsAsynchrony() *bool
	SetJobDataParsingTaskId(v string) *AssignJobsRequest
	GetJobDataParsingTaskId() *string
	SetJobGroupId(v string) *AssignJobsRequest
	GetJobGroupId() *string
	SetJobsJson(v []*string) *AssignJobsRequest
	GetJobsJson() []*string
	SetRosterType(v string) *AssignJobsRequest
	GetRosterType() *string
	SetStrategyJson(v string) *AssignJobsRequest
	GetStrategyJson() *string
}

type AssignJobsRequest struct {
	// List of calling numbers (displayed to customers when making outbound calls).
	//
	// > If this parameter is not specified, all available numbers are selected by default.
	CallingNumber []*string `json:"CallingNumber,omitempty" xml:"CallingNumber,omitempty" type:"Repeated"`
	// Instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12f3dd08-0c55-44ce-9b64-e69d35ed3a76
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Indicates whether the operation is asynchronous.
	//
	// This parameter can be left blank if the outbound call job is delivered via API integration. The default value is false.
	//
	// example:
	//
	// false
	IsAsynchrony *bool `json:"IsAsynchrony,omitempty" xml:"IsAsynchrony,omitempty"`
	// The task ID for parsing outbound call job data.
	//
	// This parameter can be left blank if the outbound call job is delivered via API integration.
	//
	// > If you use file upload, you must invoke CreateJobDataParsingTask to obtain the corresponding parameter value.
	//
	// example:
	//
	// d004cfd2-6a81-491c-83c6-cbe186620c95
	JobDataParsingTaskId *string `json:"JobDataParsingTaskId,omitempty" xml:"JobDataParsingTaskId,omitempty"`
	// Task group ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// d004cfd2-6a81-491c-83c6-cbe186620c95
	JobGroupId *string `json:"JobGroupId,omitempty" xml:"JobGroupId,omitempty"`
	// Job data in the format of a JSON array. For the specific format, refer to the Example Value.
	//
	// JobsJson.N is a List, where each JobJson corresponds to one contact.
	JobsJson []*string `json:"JobsJson,omitempty" xml:"JobsJson,omitempty" type:"Repeated"`
	// List type.
	//
	// For API integration, the default value is JSON and this parameter can be omitted.
	//
	// If uploading a list, specify "excel".
	//
	// example:
	//
	// json
	RosterType *string `json:"RosterType,omitempty" xml:"RosterType,omitempty"`
	// Execution policy for the job group:
	//
	// - repeatBy: Recurrence type. Valid values: Once (no recurrence), Day (daily), Week (weekly), Month (monthly).
	//
	// - startTime: Policy start time.
	//
	// - endTime: Policy end time.
	//
	// - workingTime: Allowed outbound calling time window.
	//
	// - maxAttemptsPerDay: Maximum number of call attempts per day for each number under this job.
	//
	// - minAttemptInterval: Minimum interval between retry calls for a number, in minutes.
	//
	// - routingStrategy: Number routing strategy. Valid values: None (not specified), LocalFirst (local city numbers prioritized), LocalProvinceFirst (local province numbers prioritized).
	//
	// - repeatDays: Execution dates corresponding to the recurrence type. If RepeatBy is Week, 0 represents Sunday and 1–6 represent Monday to Saturday. If RepeatBy is Month, values 1–31 represent the 1st to the 31st day of the month; months without the specified date will skip execution (for example, if day 30 is selected, February will be skipped).
	//
	// - repeatable: Indicates whether loop tasks are enabled. Valid values: true or false.
	//
	// example:
	//
	// {"maxAttemptsPerDay":"3","minAttemptInterval":"10","routingStrategy":"LocalProvinceFirst","repeatDays":["1","2","3"],"workingTime":[{"beginTime":"10:00:00","endTime":"11:00:00"},{"beginTime":"14:00:00","endTime":"15:00:00"}],"repeatable":true,"endTime":1707494400000,"startTime":1706976000000,"repeatBy":"Week"}
	StrategyJson *string `json:"StrategyJson,omitempty" xml:"StrategyJson,omitempty"`
}

func (s AssignJobsRequest) String() string {
	return dara.Prettify(s)
}

func (s AssignJobsRequest) GoString() string {
	return s.String()
}

func (s *AssignJobsRequest) GetCallingNumber() []*string {
	return s.CallingNumber
}

func (s *AssignJobsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *AssignJobsRequest) GetIsAsynchrony() *bool {
	return s.IsAsynchrony
}

func (s *AssignJobsRequest) GetJobDataParsingTaskId() *string {
	return s.JobDataParsingTaskId
}

func (s *AssignJobsRequest) GetJobGroupId() *string {
	return s.JobGroupId
}

func (s *AssignJobsRequest) GetJobsJson() []*string {
	return s.JobsJson
}

func (s *AssignJobsRequest) GetRosterType() *string {
	return s.RosterType
}

func (s *AssignJobsRequest) GetStrategyJson() *string {
	return s.StrategyJson
}

func (s *AssignJobsRequest) SetCallingNumber(v []*string) *AssignJobsRequest {
	s.CallingNumber = v
	return s
}

func (s *AssignJobsRequest) SetInstanceId(v string) *AssignJobsRequest {
	s.InstanceId = &v
	return s
}

func (s *AssignJobsRequest) SetIsAsynchrony(v bool) *AssignJobsRequest {
	s.IsAsynchrony = &v
	return s
}

func (s *AssignJobsRequest) SetJobDataParsingTaskId(v string) *AssignJobsRequest {
	s.JobDataParsingTaskId = &v
	return s
}

func (s *AssignJobsRequest) SetJobGroupId(v string) *AssignJobsRequest {
	s.JobGroupId = &v
	return s
}

func (s *AssignJobsRequest) SetJobsJson(v []*string) *AssignJobsRequest {
	s.JobsJson = v
	return s
}

func (s *AssignJobsRequest) SetRosterType(v string) *AssignJobsRequest {
	s.RosterType = &v
	return s
}

func (s *AssignJobsRequest) SetStrategyJson(v string) *AssignJobsRequest {
	s.StrategyJson = &v
	return s
}

func (s *AssignJobsRequest) Validate() error {
	return dara.Validate(s)
}
