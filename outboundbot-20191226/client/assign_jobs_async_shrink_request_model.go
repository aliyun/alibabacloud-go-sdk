// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssignJobsAsyncShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCallingNumberShrink(v string) *AssignJobsAsyncShrinkRequest
	GetCallingNumberShrink() *string
	SetInstanceId(v string) *AssignJobsAsyncShrinkRequest
	GetInstanceId() *string
	SetJobGroupId(v string) *AssignJobsAsyncShrinkRequest
	GetJobGroupId() *string
	SetJobsJsonShrink(v string) *AssignJobsAsyncShrinkRequest
	GetJobsJsonShrink() *string
	SetStrategyJson(v string) *AssignJobsAsyncShrinkRequest
	GetStrategyJson() *string
}

type AssignJobsAsyncShrinkRequest struct {
	// List of calling numbers (the caller ID displayed when dialing the customer).
	//
	// > If left unspecified, all available numbers are selected by default.
	CallingNumberShrink *string `json:"CallingNumber,omitempty" xml:"CallingNumber,omitempty"`
	// Instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12f3dd08-0c55-44ce-9b64-e69d35ed3a76
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Task group ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// d004cfd2-6a81-491c-83c6-cbe186620c95
	JobGroupId *string `json:"JobGroupId,omitempty" xml:"JobGroupId,omitempty"`
	// Job data, formatted as a JSON array. For the specific format, see Example Value.
	//
	// JobsJson.N is a list where each JobJson corresponds to one contact.
	JobsJsonShrink *string `json:"JobsJson,omitempty" xml:"JobsJson,omitempty"`
	// Job group execution policy.
	//
	// - repeatBy
	//
	//  Recurrence type. Valid values: Once (no recurrence), Day (daily recurrence), Week (weekly recurrence), Month (monthly recurrence).
	//
	// - startTime
	//
	//  Policy start time.
	//
	// - endTime
	//
	//  Policy end time.
	//
	// - workingTime
	//
	//  Allowed outbound calling time segment.
	//
	// - maxAttemptsPerDay
	//
	//  Maximum number of call attempts per contact per day for this job.
	//
	// - minAttemptInterval
	//
	//  Minimum interval between retry calls for a contact, in minutes.
	//
	// - routingStrategy
	//
	//  Number routing strategy. Valid values: None (not specified), LocalFirst (local city numbers preferred), LocalProvinceFirst (numbers from the same province preferred).
	//
	// - repeatDays
	//
	//  Execution dates corresponding to the recurrence type. If RepeatBy is set to Week, 0 represents Sunday and 1–6 represent Monday through Saturday, respectively. If RepeatBy is set to Month, values 1–31 represent the 1st through the 31st of each month. Months that do not have the specified date will skip execution (for example, if the 30th is selected, February will be skipped).
	//
	// - repeatable
	//
	//  Indicates whether loop tasks are enabled. Valid values: true or false.
	//
	// example:
	//
	// {"maxAttemptsPerDay":"3","minAttemptInterval":"10","routingStrategy":"LocalProvinceFirst","repeatDays":["1","2","3"],"workingTime":[{"beginTime":"10:00:00","endTime":"11:00:00"},{"beginTime":"14:00:00","endTime":"15:00:00"}],"repeatable":true,"endTime":1707494400000,"startTime":1706976000000,"repeatBy":"Week"}
	StrategyJson *string `json:"StrategyJson,omitempty" xml:"StrategyJson,omitempty"`
}

func (s AssignJobsAsyncShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s AssignJobsAsyncShrinkRequest) GoString() string {
	return s.String()
}

func (s *AssignJobsAsyncShrinkRequest) GetCallingNumberShrink() *string {
	return s.CallingNumberShrink
}

func (s *AssignJobsAsyncShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *AssignJobsAsyncShrinkRequest) GetJobGroupId() *string {
	return s.JobGroupId
}

func (s *AssignJobsAsyncShrinkRequest) GetJobsJsonShrink() *string {
	return s.JobsJsonShrink
}

func (s *AssignJobsAsyncShrinkRequest) GetStrategyJson() *string {
	return s.StrategyJson
}

func (s *AssignJobsAsyncShrinkRequest) SetCallingNumberShrink(v string) *AssignJobsAsyncShrinkRequest {
	s.CallingNumberShrink = &v
	return s
}

func (s *AssignJobsAsyncShrinkRequest) SetInstanceId(v string) *AssignJobsAsyncShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *AssignJobsAsyncShrinkRequest) SetJobGroupId(v string) *AssignJobsAsyncShrinkRequest {
	s.JobGroupId = &v
	return s
}

func (s *AssignJobsAsyncShrinkRequest) SetJobsJsonShrink(v string) *AssignJobsAsyncShrinkRequest {
	s.JobsJsonShrink = &v
	return s
}

func (s *AssignJobsAsyncShrinkRequest) SetStrategyJson(v string) *AssignJobsAsyncShrinkRequest {
	s.StrategyJson = &v
	return s
}

func (s *AssignJobsAsyncShrinkRequest) Validate() error {
	return dara.Validate(s)
}
