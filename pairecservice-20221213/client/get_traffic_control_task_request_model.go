// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTrafficControlTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetControlTargetFilter(v string) *GetTrafficControlTaskRequest
	GetControlTargetFilter() *string
	SetEnvironment(v string) *GetTrafficControlTaskRequest
	GetEnvironment() *string
	SetInstanceId(v string) *GetTrafficControlTaskRequest
	GetInstanceId() *string
	SetRegionId(v string) *GetTrafficControlTaskRequest
	GetRegionId() *string
	SetVersion(v string) *GetTrafficControlTaskRequest
	GetVersion() *string
}

type GetTrafficControlTaskRequest struct {
	ControlTargetFilter *string `json:"ControlTargetFilter,omitempty" xml:"ControlTargetFilter,omitempty"`
	Environment         *string `json:"Environment,omitempty" xml:"Environment,omitempty"`
	InstanceId          *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId            *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Version             *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s GetTrafficControlTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s GetTrafficControlTaskRequest) GoString() string {
	return s.String()
}

func (s *GetTrafficControlTaskRequest) GetControlTargetFilter() *string {
	return s.ControlTargetFilter
}

func (s *GetTrafficControlTaskRequest) GetEnvironment() *string {
	return s.Environment
}

func (s *GetTrafficControlTaskRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetTrafficControlTaskRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetTrafficControlTaskRequest) GetVersion() *string {
	return s.Version
}

func (s *GetTrafficControlTaskRequest) SetControlTargetFilter(v string) *GetTrafficControlTaskRequest {
	s.ControlTargetFilter = &v
	return s
}

func (s *GetTrafficControlTaskRequest) SetEnvironment(v string) *GetTrafficControlTaskRequest {
	s.Environment = &v
	return s
}

func (s *GetTrafficControlTaskRequest) SetInstanceId(v string) *GetTrafficControlTaskRequest {
	s.InstanceId = &v
	return s
}

func (s *GetTrafficControlTaskRequest) SetRegionId(v string) *GetTrafficControlTaskRequest {
	s.RegionId = &v
	return s
}

func (s *GetTrafficControlTaskRequest) SetVersion(v string) *GetTrafficControlTaskRequest {
	s.Version = &v
	return s
}

func (s *GetTrafficControlTaskRequest) Validate() error {
	return dara.Validate(s)
}
