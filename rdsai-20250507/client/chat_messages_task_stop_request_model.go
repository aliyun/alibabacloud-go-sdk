// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChatMessagesTaskStopRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTaskId(v string) *ChatMessagesTaskStopRequest
	GetTaskId() *string
	SetWorkspaceId(v string) *ChatMessagesTaskStopRequest
	GetWorkspaceId() *string
}

type ChatMessagesTaskStopRequest struct {
	// The unique ID of the task.
	//
	// example:
	//
	// 09a81048-0528-4de5-9dbd-12c8a12b****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The ContextDB workspace ID.
	//
	// example:
	//
	// 00000000-0000-4000-8000-000000000001
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ChatMessagesTaskStopRequest) String() string {
	return dara.Prettify(s)
}

func (s ChatMessagesTaskStopRequest) GoString() string {
	return s.String()
}

func (s *ChatMessagesTaskStopRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *ChatMessagesTaskStopRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ChatMessagesTaskStopRequest) SetTaskId(v string) *ChatMessagesTaskStopRequest {
	s.TaskId = &v
	return s
}

func (s *ChatMessagesTaskStopRequest) SetWorkspaceId(v string) *ChatMessagesTaskStopRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ChatMessagesTaskStopRequest) Validate() error {
	return dara.Validate(s)
}
