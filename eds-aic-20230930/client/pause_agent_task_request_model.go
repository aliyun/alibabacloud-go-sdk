// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPauseAgentTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTaskIds(v []*string) *PauseAgentTaskRequest
	GetTaskIds() []*string
}

type PauseAgentTaskRequest struct {
	// A list of task IDs.
	//
	// This parameter is required.
	TaskIds []*string `json:"TaskIds,omitempty" xml:"TaskIds,omitempty" type:"Repeated"`
}

func (s PauseAgentTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s PauseAgentTaskRequest) GoString() string {
	return s.String()
}

func (s *PauseAgentTaskRequest) GetTaskIds() []*string {
	return s.TaskIds
}

func (s *PauseAgentTaskRequest) SetTaskIds(v []*string) *PauseAgentTaskRequest {
	s.TaskIds = v
	return s
}

func (s *PauseAgentTaskRequest) Validate() error {
	return dara.Validate(s)
}
