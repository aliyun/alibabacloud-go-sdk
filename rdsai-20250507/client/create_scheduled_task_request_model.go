// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateScheduledTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateScheduledTaskRequest
	GetDescription() *string
	SetFrequency(v string) *CreateScheduledTaskRequest
	GetFrequency() *string
	SetInstanceIds(v string) *CreateScheduledTaskRequest
	GetInstanceIds() *string
	SetName(v string) *CreateScheduledTaskRequest
	GetName() *string
	SetStartTime(v string) *CreateScheduledTaskRequest
	GetStartTime() *string
	SetTimeRange(v string) *CreateScheduledTaskRequest
	GetTimeRange() *string
}

type CreateScheduledTaskRequest struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// Monday
	Frequency *string `json:"Frequency,omitempty" xml:"Frequency,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// rm-2ze6mk259v322****,rm-2zef3b65430j0****
	InstanceIds *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
	// This parameter is required.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 02:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// 24
	TimeRange *string `json:"TimeRange,omitempty" xml:"TimeRange,omitempty"`
}

func (s CreateScheduledTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateScheduledTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateScheduledTaskRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateScheduledTaskRequest) GetFrequency() *string {
	return s.Frequency
}

func (s *CreateScheduledTaskRequest) GetInstanceIds() *string {
	return s.InstanceIds
}

func (s *CreateScheduledTaskRequest) GetName() *string {
	return s.Name
}

func (s *CreateScheduledTaskRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *CreateScheduledTaskRequest) GetTimeRange() *string {
	return s.TimeRange
}

func (s *CreateScheduledTaskRequest) SetDescription(v string) *CreateScheduledTaskRequest {
	s.Description = &v
	return s
}

func (s *CreateScheduledTaskRequest) SetFrequency(v string) *CreateScheduledTaskRequest {
	s.Frequency = &v
	return s
}

func (s *CreateScheduledTaskRequest) SetInstanceIds(v string) *CreateScheduledTaskRequest {
	s.InstanceIds = &v
	return s
}

func (s *CreateScheduledTaskRequest) SetName(v string) *CreateScheduledTaskRequest {
	s.Name = &v
	return s
}

func (s *CreateScheduledTaskRequest) SetStartTime(v string) *CreateScheduledTaskRequest {
	s.StartTime = &v
	return s
}

func (s *CreateScheduledTaskRequest) SetTimeRange(v string) *CreateScheduledTaskRequest {
	s.TimeRange = &v
	return s
}

func (s *CreateScheduledTaskRequest) Validate() error {
	return dara.Validate(s)
}
