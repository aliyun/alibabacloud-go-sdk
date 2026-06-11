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
	// These numbers are displayed as the caller ID to the contact.
	//
	// > If unspecified, all available calling numbers are used.
	CallingNumber []*string `json:"CallingNumber,omitempty" xml:"CallingNumber,omitempty" type:"Repeated"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12f3dd08-0c55-44ce-9b64-e69d35ed3a76
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Specifies whether the task is asynchronous.
	//
	// You can omit this parameter if you create jobs by calling this API. The default value is `false`.
	//
	// example:
	//
	// false
	IsAsynchrony *bool `json:"IsAsynchrony,omitempty" xml:"IsAsynchrony,omitempty"`
	// The ID of the data parsing task for the outbound call job.
	//
	// You can omit this parameter if you create jobs by calling this API.
	//
	// > If you create jobs by uploading a file, you must call the `CreateJobDataParsingTask` operation to obtain a value for this parameter.
	//
	// example:
	//
	// d004cfd2-6a81-491c-83c6-cbe186620c95
	JobDataParsingTaskId *string `json:"JobDataParsingTaskId,omitempty" xml:"JobDataParsingTaskId,omitempty"`
	// The job group ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// d004cfd2-6a81-491c-83c6-cbe186620c95
	JobGroupId *string `json:"JobGroupId,omitempty" xml:"JobGroupId,omitempty"`
	// 	Notice:
	//
	// This parameter is required.
	//
	//
	//
	// The job data, a JSON array where each object represents a contact. For formatting details, see the example.
	//
	// The array can contain up to 99 elements.
	JobsJson []*string `json:"JobsJson,omitempty" xml:"JobsJson,omitempty" type:"Repeated"`
	// The roster type.
	//
	// If you create jobs by calling this API, you can omit this parameter. The default value is `json`. If you upload a contact list file, set this parameter to `excel`.
	//
	// example:
	//
	// json
	RosterType *string `json:"RosterType,omitempty" xml:"RosterType,omitempty"`
	// The execution strategy for the job group.
	//
	// - `repeatBy`: The recurrence type. Valid values: `Once` (runs once), `Day` (repeats daily), `Week` (repeats weekly), and `Month` (repeats monthly).
	//
	// - `startTime`: The start time of the strategy.
	//
	// - `endTime`: The end time of the strategy.
	//
	// - `workingTime`: The time windows during which calls can be made.
	//
	// - `maxAttemptsPerDay`: The maximum daily call attempts per phone number.
	//
	// - `minAttemptInterval`: The minimum interval between call retries, in minutes.
	//
	// - `routingStrategy`: The number routing strategy. Valid values: `None` (no preference), `LocalFirst` (prioritizes numbers in the same city), and `LocalProvinceFirst` (prioritizes numbers in the same province).
	//
	// - `repeatDays`: The specific days on which the job runs, based on the `repeatBy` type. If `repeatBy` is set to `Week`, `0` represents Sunday, and `1` through `6` represent Monday through Saturday. If `repeatBy` is set to `Month`, values from `1` to `31` represent the days of the month. If a month does not have a specified day (for example, February 30), the job is skipped for that month.
	//
	// - `repeatable`: Specifies whether the job is recurring. Valid values are `true` and `false`.
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
