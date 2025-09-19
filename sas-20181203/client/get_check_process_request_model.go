// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCheckProcessRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTaskId(v string) *GetCheckProcessRequest
	GetTaskId() *string
}

type GetCheckProcessRequest struct {
	// The ID of the task.
	//
	// > You can call the [SubmitCheck](~~SubmitCheck~~) operation to query the ID.
	//
	// example:
	//
	// 5347c7b6-c85c-4070-846a-3029e08e****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s GetCheckProcessRequest) String() string {
	return dara.Prettify(s)
}

func (s GetCheckProcessRequest) GoString() string {
	return s.String()
}

func (s *GetCheckProcessRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *GetCheckProcessRequest) SetTaskId(v string) *GetCheckProcessRequest {
	s.TaskId = &v
	return s
}

func (s *GetCheckProcessRequest) Validate() error {
	return dara.Validate(s)
}
