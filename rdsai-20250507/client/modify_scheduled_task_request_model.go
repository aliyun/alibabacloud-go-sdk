// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyScheduledTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *ModifyScheduledTaskRequest
	GetDescription() *string
	SetFrequency(v string) *ModifyScheduledTaskRequest
	GetFrequency() *string
	SetInstanceIds(v string) *ModifyScheduledTaskRequest
	GetInstanceIds() *string
	SetName(v string) *ModifyScheduledTaskRequest
	GetName() *string
	SetScheduledId(v string) *ModifyScheduledTaskRequest
	GetScheduledId() *string
	SetStartTime(v string) *ModifyScheduledTaskRequest
	GetStartTime() *string
	SetTimeRange(v string) *ModifyScheduledTaskRequest
	GetTimeRange() *string
}

type ModifyScheduledTaskRequest struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// Monday
	Frequency *string `json:"Frequency,omitempty" xml:"Frequency,omitempty"`
	// example:
	//
	// rm-2ze6mk259v322****,rm-2zef3b65430j0****
	InstanceIds *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 847268a4-196f-416b-aa12-bfe0c115****
	ScheduledId *string `json:"ScheduledId,omitempty" xml:"ScheduledId,omitempty"`
	// example:
	//
	// 02:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// 24
	TimeRange *string `json:"TimeRange,omitempty" xml:"TimeRange,omitempty"`
}

func (s ModifyScheduledTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyScheduledTaskRequest) GoString() string {
	return s.String()
}

func (s *ModifyScheduledTaskRequest) GetDescription() *string {
	return s.Description
}

func (s *ModifyScheduledTaskRequest) GetFrequency() *string {
	return s.Frequency
}

func (s *ModifyScheduledTaskRequest) GetInstanceIds() *string {
	return s.InstanceIds
}

func (s *ModifyScheduledTaskRequest) GetName() *string {
	return s.Name
}

func (s *ModifyScheduledTaskRequest) GetScheduledId() *string {
	return s.ScheduledId
}

func (s *ModifyScheduledTaskRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *ModifyScheduledTaskRequest) GetTimeRange() *string {
	return s.TimeRange
}

func (s *ModifyScheduledTaskRequest) SetDescription(v string) *ModifyScheduledTaskRequest {
	s.Description = &v
	return s
}

func (s *ModifyScheduledTaskRequest) SetFrequency(v string) *ModifyScheduledTaskRequest {
	s.Frequency = &v
	return s
}

func (s *ModifyScheduledTaskRequest) SetInstanceIds(v string) *ModifyScheduledTaskRequest {
	s.InstanceIds = &v
	return s
}

func (s *ModifyScheduledTaskRequest) SetName(v string) *ModifyScheduledTaskRequest {
	s.Name = &v
	return s
}

func (s *ModifyScheduledTaskRequest) SetScheduledId(v string) *ModifyScheduledTaskRequest {
	s.ScheduledId = &v
	return s
}

func (s *ModifyScheduledTaskRequest) SetStartTime(v string) *ModifyScheduledTaskRequest {
	s.StartTime = &v
	return s
}

func (s *ModifyScheduledTaskRequest) SetTimeRange(v string) *ModifyScheduledTaskRequest {
	s.TimeRange = &v
	return s
}

func (s *ModifyScheduledTaskRequest) Validate() error {
	return dara.Validate(s)
}
