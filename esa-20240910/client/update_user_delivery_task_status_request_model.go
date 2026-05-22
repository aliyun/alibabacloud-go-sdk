// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateUserDeliveryTaskStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMethod(v string) *UpdateUserDeliveryTaskStatusRequest
	GetMethod() *string
	SetTaskName(v string) *UpdateUserDeliveryTaskStatusRequest
	GetTaskName() *string
}

type UpdateUserDeliveryTaskStatusRequest struct {
	// Enables or disables the delivery task. Valid values: online and offline.
	//
	// This parameter is required.
	//
	// example:
	//
	// online
	Method *string `json:"Method,omitempty" xml:"Method,omitempty"`
	// The name of the delivery task.
	//
	// This parameter is required.
	//
	// example:
	//
	// test_project
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
}

func (s UpdateUserDeliveryTaskStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateUserDeliveryTaskStatusRequest) GoString() string {
	return s.String()
}

func (s *UpdateUserDeliveryTaskStatusRequest) GetMethod() *string {
	return s.Method
}

func (s *UpdateUserDeliveryTaskStatusRequest) GetTaskName() *string {
	return s.TaskName
}

func (s *UpdateUserDeliveryTaskStatusRequest) SetMethod(v string) *UpdateUserDeliveryTaskStatusRequest {
	s.Method = &v
	return s
}

func (s *UpdateUserDeliveryTaskStatusRequest) SetTaskName(v string) *UpdateUserDeliveryTaskStatusRequest {
	s.TaskName = &v
	return s
}

func (s *UpdateUserDeliveryTaskStatusRequest) Validate() error {
	return dara.Validate(s)
}
