// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteUserDeliveryTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTaskName(v string) *DeleteUserDeliveryTaskRequest
	GetTaskName() *string
}

type DeleteUserDeliveryTaskRequest struct {
	// The name of the delivery task.
	//
	// This parameter is required.
	//
	// example:
	//
	// test-project
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
}

func (s DeleteUserDeliveryTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteUserDeliveryTaskRequest) GoString() string {
	return s.String()
}

func (s *DeleteUserDeliveryTaskRequest) GetTaskName() *string {
	return s.TaskName
}

func (s *DeleteUserDeliveryTaskRequest) SetTaskName(v string) *DeleteUserDeliveryTaskRequest {
	s.TaskName = &v
	return s
}

func (s *DeleteUserDeliveryTaskRequest) Validate() error {
	return dara.Validate(s)
}
