// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopTrafficControlTargetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnvironment(v string) *StopTrafficControlTargetRequest
	GetEnvironment() *string
	SetInstanceId(v string) *StopTrafficControlTargetRequest
	GetInstanceId() *string
}

type StopTrafficControlTargetRequest struct {
	// The environment.
	//
	// example:
	//
	// Pre
	Environment *string `json:"Environment,omitempty" xml:"Environment,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// pairec-1324***
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s StopTrafficControlTargetRequest) String() string {
	return dara.Prettify(s)
}

func (s StopTrafficControlTargetRequest) GoString() string {
	return s.String()
}

func (s *StopTrafficControlTargetRequest) GetEnvironment() *string {
	return s.Environment
}

func (s *StopTrafficControlTargetRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *StopTrafficControlTargetRequest) SetEnvironment(v string) *StopTrafficControlTargetRequest {
	s.Environment = &v
	return s
}

func (s *StopTrafficControlTargetRequest) SetInstanceId(v string) *StopTrafficControlTargetRequest {
	s.InstanceId = &v
	return s
}

func (s *StopTrafficControlTargetRequest) Validate() error {
	return dara.Validate(s)
}
