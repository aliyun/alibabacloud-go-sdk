// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWmEmbedTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTaskId(v string) *GetWmEmbedTaskRequest
	GetTaskId() *string
}

type GetWmEmbedTaskRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// job:5GfrJYsoaffmCE7Z5bZtjU********
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s GetWmEmbedTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s GetWmEmbedTaskRequest) GoString() string {
	return s.String()
}

func (s *GetWmEmbedTaskRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *GetWmEmbedTaskRequest) SetTaskId(v string) *GetWmEmbedTaskRequest {
	s.TaskId = &v
	return s
}

func (s *GetWmEmbedTaskRequest) Validate() error {
	return dara.Validate(s)
}
