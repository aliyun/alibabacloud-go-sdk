// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartTrafficControlTargetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnvironment(v string) *StartTrafficControlTargetRequest
	GetEnvironment() *string
	SetInstanceId(v string) *StartTrafficControlTargetRequest
	GetInstanceId() *string
}

type StartTrafficControlTargetRequest struct {
	// The environment.
	//
	// example:
	//
	// Prod
	Environment *string `json:"Environment,omitempty" xml:"Environment,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// pairec-1324***
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s StartTrafficControlTargetRequest) String() string {
	return dara.Prettify(s)
}

func (s StartTrafficControlTargetRequest) GoString() string {
	return s.String()
}

func (s *StartTrafficControlTargetRequest) GetEnvironment() *string {
	return s.Environment
}

func (s *StartTrafficControlTargetRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *StartTrafficControlTargetRequest) SetEnvironment(v string) *StartTrafficControlTargetRequest {
	s.Environment = &v
	return s
}

func (s *StartTrafficControlTargetRequest) SetInstanceId(v string) *StartTrafficControlTargetRequest {
	s.InstanceId = &v
	return s
}

func (s *StartTrafficControlTargetRequest) Validate() error {
	return dara.Validate(s)
}
