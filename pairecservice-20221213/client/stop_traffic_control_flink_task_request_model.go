// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopTrafficControlFlinkTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnvironment(v string) *StopTrafficControlFlinkTaskRequest
	GetEnvironment() *string
	SetInstanceId(v string) *StopTrafficControlFlinkTaskRequest
	GetInstanceId() *string
}

type StopTrafficControlFlinkTaskRequest struct {
	// The environment to which the instance belongs. Valid values:
	//
	// - Daily: daily environment.
	//
	// - Pre: staging environment.
	//
	// - Prod: production environment.
	//
	// example:
	//
	// Daily
	Environment *string `json:"Environment,omitempty" xml:"Environment,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// pairec_123****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s StopTrafficControlFlinkTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s StopTrafficControlFlinkTaskRequest) GoString() string {
	return s.String()
}

func (s *StopTrafficControlFlinkTaskRequest) GetEnvironment() *string {
	return s.Environment
}

func (s *StopTrafficControlFlinkTaskRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *StopTrafficControlFlinkTaskRequest) SetEnvironment(v string) *StopTrafficControlFlinkTaskRequest {
	s.Environment = &v
	return s
}

func (s *StopTrafficControlFlinkTaskRequest) SetInstanceId(v string) *StopTrafficControlFlinkTaskRequest {
	s.InstanceId = &v
	return s
}

func (s *StopTrafficControlFlinkTaskRequest) Validate() error {
	return dara.Validate(s)
}
