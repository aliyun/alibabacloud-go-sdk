// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssignJobsAsyncRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCallingNumber(v []*string) *AssignJobsAsyncRequest
	GetCallingNumber() []*string
	SetInstanceId(v string) *AssignJobsAsyncRequest
	GetInstanceId() *string
	SetJobGroupId(v string) *AssignJobsAsyncRequest
	GetJobGroupId() *string
	SetJobsJson(v []*string) *AssignJobsAsyncRequest
	GetJobsJson() []*string
	SetStrategyJson(v string) *AssignJobsAsyncRequest
	GetStrategyJson() *string
}

type AssignJobsAsyncRequest struct {
	CallingNumber []*string `json:"CallingNumber,omitempty" xml:"CallingNumber,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 12f3dd08-0c55-44ce-9b64-e69d35ed3a76
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// d004cfd2-6a81-491c-83c6-cbe186620c95
	JobGroupId *string   `json:"JobGroupId,omitempty" xml:"JobGroupId,omitempty"`
	JobsJson   []*string `json:"JobsJson,omitempty" xml:"JobsJson,omitempty" type:"Repeated"`
	// example:
	//
	// {"maxAttemptsPerDay":"3","minAttemptInterval":"10","routingStrategy":"LocalProvinceFirst","repeatDays":["1","2","3"],"workingTime":[{"beginTime":"10:00:00","endTime":"11:00:00"},{"beginTime":"14:00:00","endTime":"15:00:00"}],"repeatable":true,"endTime":1707494400000,"startTime":1706976000000,"repeatBy":"Week"}
	StrategyJson *string `json:"StrategyJson,omitempty" xml:"StrategyJson,omitempty"`
}

func (s AssignJobsAsyncRequest) String() string {
	return dara.Prettify(s)
}

func (s AssignJobsAsyncRequest) GoString() string {
	return s.String()
}

func (s *AssignJobsAsyncRequest) GetCallingNumber() []*string {
	return s.CallingNumber
}

func (s *AssignJobsAsyncRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *AssignJobsAsyncRequest) GetJobGroupId() *string {
	return s.JobGroupId
}

func (s *AssignJobsAsyncRequest) GetJobsJson() []*string {
	return s.JobsJson
}

func (s *AssignJobsAsyncRequest) GetStrategyJson() *string {
	return s.StrategyJson
}

func (s *AssignJobsAsyncRequest) SetCallingNumber(v []*string) *AssignJobsAsyncRequest {
	s.CallingNumber = v
	return s
}

func (s *AssignJobsAsyncRequest) SetInstanceId(v string) *AssignJobsAsyncRequest {
	s.InstanceId = &v
	return s
}

func (s *AssignJobsAsyncRequest) SetJobGroupId(v string) *AssignJobsAsyncRequest {
	s.JobGroupId = &v
	return s
}

func (s *AssignJobsAsyncRequest) SetJobsJson(v []*string) *AssignJobsAsyncRequest {
	s.JobsJson = v
	return s
}

func (s *AssignJobsAsyncRequest) SetStrategyJson(v string) *AssignJobsAsyncRequest {
	s.StrategyJson = &v
	return s
}

func (s *AssignJobsAsyncRequest) Validate() error {
	return dara.Validate(s)
}
