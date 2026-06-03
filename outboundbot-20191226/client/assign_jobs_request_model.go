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
	CallingNumber []*string `json:"CallingNumber,omitempty" xml:"CallingNumber,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 12f3dd08-0c55-44ce-9b64-e69d35ed3a76
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// false
	IsAsynchrony *bool `json:"IsAsynchrony,omitempty" xml:"IsAsynchrony,omitempty"`
	// example:
	//
	// d004cfd2-6a81-491c-83c6-cbe186620c95
	JobDataParsingTaskId *string `json:"JobDataParsingTaskId,omitempty" xml:"JobDataParsingTaskId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// d004cfd2-6a81-491c-83c6-cbe186620c95
	JobGroupId *string   `json:"JobGroupId,omitempty" xml:"JobGroupId,omitempty"`
	JobsJson   []*string `json:"JobsJson,omitempty" xml:"JobsJson,omitempty" type:"Repeated"`
	// example:
	//
	// json
	RosterType   *string `json:"RosterType,omitempty" xml:"RosterType,omitempty"`
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
