// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopAIAgentInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *StopAIAgentInstanceRequest
	GetInstanceId() *string
}

type StopAIAgentInstanceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 39f8e0bc005e4f309379701645f4****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s StopAIAgentInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s StopAIAgentInstanceRequest) GoString() string {
	return s.String()
}

func (s *StopAIAgentInstanceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *StopAIAgentInstanceRequest) SetInstanceId(v string) *StopAIAgentInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *StopAIAgentInstanceRequest) Validate() error {
	return dara.Validate(s)
}
