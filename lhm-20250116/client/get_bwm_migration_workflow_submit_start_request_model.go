// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBwmMigrationWorkflowSubmitStartRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTaskId(v string) *GetBwmMigrationWorkflowSubmitStartRequest
	GetTaskId() *string
}

type GetBwmMigrationWorkflowSubmitStartRequest struct {
	// The scheduling migration task ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12345
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s GetBwmMigrationWorkflowSubmitStartRequest) String() string {
	return dara.Prettify(s)
}

func (s GetBwmMigrationWorkflowSubmitStartRequest) GoString() string {
	return s.String()
}

func (s *GetBwmMigrationWorkflowSubmitStartRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *GetBwmMigrationWorkflowSubmitStartRequest) SetTaskId(v string) *GetBwmMigrationWorkflowSubmitStartRequest {
	s.TaskId = &v
	return s
}

func (s *GetBwmMigrationWorkflowSubmitStartRequest) Validate() error {
	return dara.Validate(s)
}
