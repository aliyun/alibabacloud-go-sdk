// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTaskAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetTaskAttributeResponseBody
	GetRequestId() *string
	SetTaskMessage(v string) *GetTaskAttributeResponseBody
	GetTaskMessage() *string
	SetTaskStatus(v string) *GetTaskAttributeResponseBody
	GetTaskStatus() *string
}

type GetTaskAttributeResponseBody struct {
	// example:
	//
	// CBF1E9B7-D6A0-4E9E-AD3E-2B47E6C2837D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// error
	TaskMessage *string `json:"TaskMessage,omitempty" xml:"TaskMessage,omitempty"`
	// example:
	//
	// success
	TaskStatus *string `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
}

func (s GetTaskAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTaskAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *GetTaskAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTaskAttributeResponseBody) GetTaskMessage() *string {
	return s.TaskMessage
}

func (s *GetTaskAttributeResponseBody) GetTaskStatus() *string {
	return s.TaskStatus
}

func (s *GetTaskAttributeResponseBody) SetRequestId(v string) *GetTaskAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTaskAttributeResponseBody) SetTaskMessage(v string) *GetTaskAttributeResponseBody {
	s.TaskMessage = &v
	return s
}

func (s *GetTaskAttributeResponseBody) SetTaskStatus(v string) *GetTaskAttributeResponseBody {
	s.TaskStatus = &v
	return s
}

func (s *GetTaskAttributeResponseBody) Validate() error {
	return dara.Validate(s)
}
