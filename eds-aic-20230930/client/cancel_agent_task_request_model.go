// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelAgentTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTaskIds(v []*string) *CancelAgentTaskRequest
	GetTaskIds() []*string
}

type CancelAgentTaskRequest struct {
	// This parameter is required.
	TaskIds []*string `json:"TaskIds,omitempty" xml:"TaskIds,omitempty" type:"Repeated"`
}

func (s CancelAgentTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CancelAgentTaskRequest) GoString() string {
	return s.String()
}

func (s *CancelAgentTaskRequest) GetTaskIds() []*string {
	return s.TaskIds
}

func (s *CancelAgentTaskRequest) SetTaskIds(v []*string) *CancelAgentTaskRequest {
	s.TaskIds = v
	return s
}

func (s *CancelAgentTaskRequest) Validate() error {
	return dara.Validate(s)
}
