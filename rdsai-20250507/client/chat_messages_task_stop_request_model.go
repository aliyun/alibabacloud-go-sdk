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
}

type ChatMessagesTaskStopRequest struct {
	// The operation that you want to perform. Set the value to **ChatMessagesTaskStop**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 09a81048-0528-4de5-9dbd-12c8a12b****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
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

func (s *ChatMessagesTaskStopRequest) SetTaskId(v string) *ChatMessagesTaskStopRequest {
	s.TaskId = &v
	return s
}

func (s *ChatMessagesTaskStopRequest) Validate() error {
	return dara.Validate(s)
}
