// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopTrafficControlTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *StopTrafficControlTaskRequest
	GetRegionId() *string
	SetEnvironment(v string) *StopTrafficControlTaskRequest
	GetEnvironment() *string
	SetInstanceId(v string) *StopTrafficControlTaskRequest
	GetInstanceId() *string
}

type StopTrafficControlTaskRequest struct {
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// Daily
	Environment *string `json:"Environment,omitempty" xml:"Environment,omitempty"`
	// example:
	//
	// pairec_123****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s StopTrafficControlTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s StopTrafficControlTaskRequest) GoString() string {
	return s.String()
}

func (s *StopTrafficControlTaskRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *StopTrafficControlTaskRequest) GetEnvironment() *string {
	return s.Environment
}

func (s *StopTrafficControlTaskRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *StopTrafficControlTaskRequest) SetRegionId(v string) *StopTrafficControlTaskRequest {
	s.RegionId = &v
	return s
}

func (s *StopTrafficControlTaskRequest) SetEnvironment(v string) *StopTrafficControlTaskRequest {
	s.Environment = &v
	return s
}

func (s *StopTrafficControlTaskRequest) SetInstanceId(v string) *StopTrafficControlTaskRequest {
	s.InstanceId = &v
	return s
}

func (s *StopTrafficControlTaskRequest) Validate() error {
	return dara.Validate(s)
}
