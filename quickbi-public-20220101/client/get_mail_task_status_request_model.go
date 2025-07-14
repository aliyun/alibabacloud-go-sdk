// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMailTaskStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMailId(v string) *GetMailTaskStatusRequest
	GetMailId() *string
	SetTaskId(v int64) *GetMailTaskStatusRequest
	GetTaskId() *int64
}

type GetMailTaskStatusRequest struct {
	// Mail ID
	//
	// This parameter is required.
	//
	// example:
	//
	// d5a5****8b634d****5584f8dc159c62
	MailId *string `json:"MailId,omitempty" xml:"MailId,omitempty"`
	// Task ID
	//
	// > - If the task ID is not provided, the latest task status will be returned by default;
	//
	// > - If the task ID is provided, the status of the specified task will be returned.
	//
	// example:
	//
	// 7218****0392****212
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s GetMailTaskStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s GetMailTaskStatusRequest) GoString() string {
	return s.String()
}

func (s *GetMailTaskStatusRequest) GetMailId() *string {
	return s.MailId
}

func (s *GetMailTaskStatusRequest) GetTaskId() *int64 {
	return s.TaskId
}

func (s *GetMailTaskStatusRequest) SetMailId(v string) *GetMailTaskStatusRequest {
	s.MailId = &v
	return s
}

func (s *GetMailTaskStatusRequest) SetTaskId(v int64) *GetMailTaskStatusRequest {
	s.TaskId = &v
	return s
}

func (s *GetMailTaskStatusRequest) Validate() error {
	return dara.Validate(s)
}
