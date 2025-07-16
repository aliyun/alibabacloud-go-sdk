// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDocTranslateTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateDocTranslateTaskResponseBody
	GetRequestId() *string
	SetStatus(v string) *CreateDocTranslateTaskResponseBody
	GetStatus() *string
	SetTaskId(v string) *CreateDocTranslateTaskResponseBody
	GetTaskId() *string
}

type CreateDocTranslateTaskResponseBody struct {
	// example:
	//
	// D3920BC3-A395-4CAD-9E84-7C39EB07507B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// ready
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 0586df512c8b4bb382d7d9a6a01b5854
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CreateDocTranslateTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDocTranslateTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDocTranslateTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDocTranslateTaskResponseBody) GetStatus() *string {
	return s.Status
}

func (s *CreateDocTranslateTaskResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *CreateDocTranslateTaskResponseBody) SetRequestId(v string) *CreateDocTranslateTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDocTranslateTaskResponseBody) SetStatus(v string) *CreateDocTranslateTaskResponseBody {
	s.Status = &v
	return s
}

func (s *CreateDocTranslateTaskResponseBody) SetTaskId(v string) *CreateDocTranslateTaskResponseBody {
	s.TaskId = &v
	return s
}

func (s *CreateDocTranslateTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
