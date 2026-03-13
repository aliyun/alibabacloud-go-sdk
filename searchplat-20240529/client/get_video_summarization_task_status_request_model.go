// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVideoSummarizationTaskStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTaskId(v string) *GetVideoSummarizationTaskStatusRequest
	GetTaskId() *string
}

type GetVideoSummarizationTaskStatusRequest struct {
	// This parameter is required.
	TaskId *string `json:"task_id,omitempty" xml:"task_id,omitempty"`
}

func (s GetVideoSummarizationTaskStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s GetVideoSummarizationTaskStatusRequest) GoString() string {
	return s.String()
}

func (s *GetVideoSummarizationTaskStatusRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *GetVideoSummarizationTaskStatusRequest) SetTaskId(v string) *GetVideoSummarizationTaskStatusRequest {
	s.TaskId = &v
	return s
}

func (s *GetVideoSummarizationTaskStatusRequest) Validate() error {
	return dara.Validate(s)
}
