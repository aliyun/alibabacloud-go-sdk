// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserDeliveryTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTaskName(v string) *GetUserDeliveryTaskRequest
	GetTaskName() *string
}

type GetUserDeliveryTaskRequest struct {
	// The name of the delivery task.
	//
	// This parameter is required.
	//
	// example:
	//
	// test_project
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
}

func (s GetUserDeliveryTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s GetUserDeliveryTaskRequest) GoString() string {
	return s.String()
}

func (s *GetUserDeliveryTaskRequest) GetTaskName() *string {
	return s.TaskName
}

func (s *GetUserDeliveryTaskRequest) SetTaskName(v string) *GetUserDeliveryTaskRequest {
	s.TaskName = &v
	return s
}

func (s *GetUserDeliveryTaskRequest) Validate() error {
	return dara.Validate(s)
}
