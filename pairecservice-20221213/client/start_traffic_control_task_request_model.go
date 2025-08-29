// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartTrafficControlTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnvironment(v string) *StartTrafficControlTaskRequest
	GetEnvironment() *string
	SetInstanceId(v string) *StartTrafficControlTaskRequest
	GetInstanceId() *string
}

type StartTrafficControlTaskRequest struct {
	Environment *string `json:"Environment,omitempty" xml:"Environment,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s StartTrafficControlTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s StartTrafficControlTaskRequest) GoString() string {
	return s.String()
}

func (s *StartTrafficControlTaskRequest) GetEnvironment() *string {
	return s.Environment
}

func (s *StartTrafficControlTaskRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *StartTrafficControlTaskRequest) SetEnvironment(v string) *StartTrafficControlTaskRequest {
	s.Environment = &v
	return s
}

func (s *StartTrafficControlTaskRequest) SetInstanceId(v string) *StartTrafficControlTaskRequest {
	s.InstanceId = &v
	return s
}

func (s *StartTrafficControlTaskRequest) Validate() error {
	return dara.Validate(s)
}
