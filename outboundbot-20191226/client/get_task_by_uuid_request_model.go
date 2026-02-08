// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTaskByUuidRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *GetTaskByUuidRequest
	GetInstanceId() *string
	SetTaskId(v string) *GetTaskByUuidRequest
	GetTaskId() *string
	SetWithConversations(v bool) *GetTaskByUuidRequest
	GetWithConversations() *bool
}

type GetTaskByUuidRequest struct {
	// Instance ID
	//
	// This parameter is required.
	//
	// example:
	//
	// 024f8cf0-c842-4c01-b74b-c8667e4579c7
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Job ID
	//
	// This parameter is required.
	//
	// example:
	//
	// 15160071061
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// Specifies whether to include conversation information.
	//
	// > If this parameter is not specified, the default value is false.
	//
	// example:
	//
	// false
	WithConversations *bool `json:"WithConversations,omitempty" xml:"WithConversations,omitempty"`
}

func (s GetTaskByUuidRequest) String() string {
	return dara.Prettify(s)
}

func (s GetTaskByUuidRequest) GoString() string {
	return s.String()
}

func (s *GetTaskByUuidRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetTaskByUuidRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *GetTaskByUuidRequest) GetWithConversations() *bool {
	return s.WithConversations
}

func (s *GetTaskByUuidRequest) SetInstanceId(v string) *GetTaskByUuidRequest {
	s.InstanceId = &v
	return s
}

func (s *GetTaskByUuidRequest) SetTaskId(v string) *GetTaskByUuidRequest {
	s.TaskId = &v
	return s
}

func (s *GetTaskByUuidRequest) SetWithConversations(v bool) *GetTaskByUuidRequest {
	s.WithConversations = &v
	return s
}

func (s *GetTaskByUuidRequest) Validate() error {
	return dara.Validate(s)
}
