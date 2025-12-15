// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetImageAnalyzeTaskStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTaskId(v string) *GetImageAnalyzeTaskStatusRequest
	GetTaskId() *string
}

type GetImageAnalyzeTaskStatusRequest struct {
	// This parameter is required.
	TaskId *string `json:"task_id,omitempty" xml:"task_id,omitempty"`
}

func (s GetImageAnalyzeTaskStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s GetImageAnalyzeTaskStatusRequest) GoString() string {
	return s.String()
}

func (s *GetImageAnalyzeTaskStatusRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *GetImageAnalyzeTaskStatusRequest) SetTaskId(v string) *GetImageAnalyzeTaskStatusRequest {
	s.TaskId = &v
	return s
}

func (s *GetImageAnalyzeTaskStatusRequest) Validate() error {
	return dara.Validate(s)
}
