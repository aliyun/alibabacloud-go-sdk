// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFailoverTestRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigs(v string) *FailoverTestRequest
	GetConfigs() *string
	SetDescription(v string) *FailoverTestRequest
	GetDescription() *string
	SetDuration(v int64) *FailoverTestRequest
	GetDuration() *int64
	SetExecuteTime(v int64) *FailoverTestRequest
	GetExecuteTime() *int64
	SetInstanceId(v string) *FailoverTestRequest
	GetInstanceId() *string
	SetName(v string) *FailoverTestRequest
	GetName() *string
	SetRegionId(v string) *FailoverTestRequest
	GetRegionId() *string
	SetType(v string) *FailoverTestRequest
	GetType() *string
}

type FailoverTestRequest struct {
	Configs     *string `json:"Configs,omitempty" xml:"Configs,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Duration    *int64  `json:"Duration,omitempty" xml:"Duration,omitempty"`
	ExecuteTime *int64  `json:"ExecuteTime,omitempty" xml:"ExecuteTime,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Type        *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s FailoverTestRequest) String() string {
	return dara.Prettify(s)
}

func (s FailoverTestRequest) GoString() string {
	return s.String()
}

func (s *FailoverTestRequest) GetConfigs() *string {
	return s.Configs
}

func (s *FailoverTestRequest) GetDescription() *string {
	return s.Description
}

func (s *FailoverTestRequest) GetDuration() *int64 {
	return s.Duration
}

func (s *FailoverTestRequest) GetExecuteTime() *int64 {
	return s.ExecuteTime
}

func (s *FailoverTestRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *FailoverTestRequest) GetName() *string {
	return s.Name
}

func (s *FailoverTestRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *FailoverTestRequest) GetType() *string {
	return s.Type
}

func (s *FailoverTestRequest) SetConfigs(v string) *FailoverTestRequest {
	s.Configs = &v
	return s
}

func (s *FailoverTestRequest) SetDescription(v string) *FailoverTestRequest {
	s.Description = &v
	return s
}

func (s *FailoverTestRequest) SetDuration(v int64) *FailoverTestRequest {
	s.Duration = &v
	return s
}

func (s *FailoverTestRequest) SetExecuteTime(v int64) *FailoverTestRequest {
	s.ExecuteTime = &v
	return s
}

func (s *FailoverTestRequest) SetInstanceId(v string) *FailoverTestRequest {
	s.InstanceId = &v
	return s
}

func (s *FailoverTestRequest) SetName(v string) *FailoverTestRequest {
	s.Name = &v
	return s
}

func (s *FailoverTestRequest) SetRegionId(v string) *FailoverTestRequest {
	s.RegionId = &v
	return s
}

func (s *FailoverTestRequest) SetType(v string) *FailoverTestRequest {
	s.Type = &v
	return s
}

func (s *FailoverTestRequest) Validate() error {
	return dara.Validate(s)
}
