// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApplyTripTaskExecuteRequest interface {
	dara.Model
	String() string
	GoString() string
	SetActionFrom(v string) *ApplyTripTaskExecuteRequest
	GetActionFrom() *string
	SetComment(v string) *ApplyTripTaskExecuteRequest
	GetComment() *string
	SetTaskAction(v string) *ApplyTripTaskExecuteRequest
	GetTaskAction() *string
	SetTaskId(v int64) *ApplyTripTaskExecuteRequest
	GetTaskId() *int64
	SetUserId(v string) *ApplyTripTaskExecuteRequest
	GetUserId() *string
	SetUserName(v string) *ApplyTripTaskExecuteRequest
	GetUserName() *string
}

type ApplyTripTaskExecuteRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// open
	ActionFrom *string `json:"action_from,omitempty" xml:"action_from,omitempty"`
	Comment    *string `json:"comment,omitempty" xml:"comment,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// agree
	TaskAction *string `json:"task_action,omitempty" xml:"task_action,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 12345
	TaskId *int64 `json:"task_id,omitempty" xml:"task_id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// thirdpart12138
	UserId   *string `json:"user_id,omitempty" xml:"user_id,omitempty"`
	UserName *string `json:"user_name,omitempty" xml:"user_name,omitempty"`
}

func (s ApplyTripTaskExecuteRequest) String() string {
	return dara.Prettify(s)
}

func (s ApplyTripTaskExecuteRequest) GoString() string {
	return s.String()
}

func (s *ApplyTripTaskExecuteRequest) GetActionFrom() *string {
	return s.ActionFrom
}

func (s *ApplyTripTaskExecuteRequest) GetComment() *string {
	return s.Comment
}

func (s *ApplyTripTaskExecuteRequest) GetTaskAction() *string {
	return s.TaskAction
}

func (s *ApplyTripTaskExecuteRequest) GetTaskId() *int64 {
	return s.TaskId
}

func (s *ApplyTripTaskExecuteRequest) GetUserId() *string {
	return s.UserId
}

func (s *ApplyTripTaskExecuteRequest) GetUserName() *string {
	return s.UserName
}

func (s *ApplyTripTaskExecuteRequest) SetActionFrom(v string) *ApplyTripTaskExecuteRequest {
	s.ActionFrom = &v
	return s
}

func (s *ApplyTripTaskExecuteRequest) SetComment(v string) *ApplyTripTaskExecuteRequest {
	s.Comment = &v
	return s
}

func (s *ApplyTripTaskExecuteRequest) SetTaskAction(v string) *ApplyTripTaskExecuteRequest {
	s.TaskAction = &v
	return s
}

func (s *ApplyTripTaskExecuteRequest) SetTaskId(v int64) *ApplyTripTaskExecuteRequest {
	s.TaskId = &v
	return s
}

func (s *ApplyTripTaskExecuteRequest) SetUserId(v string) *ApplyTripTaskExecuteRequest {
	s.UserId = &v
	return s
}

func (s *ApplyTripTaskExecuteRequest) SetUserName(v string) *ApplyTripTaskExecuteRequest {
	s.UserName = &v
	return s
}

func (s *ApplyTripTaskExecuteRequest) Validate() error {
	return dara.Validate(s)
}
