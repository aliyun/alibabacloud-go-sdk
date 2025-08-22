// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTaskType(v string) *GetJobRequest
	GetTaskType() *string
}

type GetJobRequest struct {
	TaskType *string `json:"taskType,omitempty" xml:"taskType,omitempty"`
}

func (s GetJobRequest) String() string {
	return dara.Prettify(s)
}

func (s GetJobRequest) GoString() string {
	return s.String()
}

func (s *GetJobRequest) GetTaskType() *string {
	return s.TaskType
}

func (s *GetJobRequest) SetTaskType(v string) *GetJobRequest {
	s.TaskType = &v
	return s
}

func (s *GetJobRequest) Validate() error {
	return dara.Validate(s)
}
