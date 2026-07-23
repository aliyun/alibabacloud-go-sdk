// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReleaseTrafficControlTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnvironment(v string) *ReleaseTrafficControlTaskRequest
	GetEnvironment() *string
	SetInstanceId(v string) *ReleaseTrafficControlTaskRequest
	GetInstanceId() *string
}

type ReleaseTrafficControlTaskRequest struct {
	// The execution environment for the traffic control task. Valid values: `Daily`, `Pre`, and `Prod`.
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

func (s ReleaseTrafficControlTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s ReleaseTrafficControlTaskRequest) GoString() string {
	return s.String()
}

func (s *ReleaseTrafficControlTaskRequest) GetEnvironment() *string {
	return s.Environment
}

func (s *ReleaseTrafficControlTaskRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ReleaseTrafficControlTaskRequest) SetEnvironment(v string) *ReleaseTrafficControlTaskRequest {
	s.Environment = &v
	return s
}

func (s *ReleaseTrafficControlTaskRequest) SetInstanceId(v string) *ReleaseTrafficControlTaskRequest {
	s.InstanceId = &v
	return s
}

func (s *ReleaseTrafficControlTaskRequest) Validate() error {
	return dara.Validate(s)
}
