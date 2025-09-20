// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateImageSplicingTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEventId(v string) *CreateImageSplicingTaskResponseBody
	GetEventId() *string
	SetRequestId(v string) *CreateImageSplicingTaskResponseBody
	GetRequestId() *string
	SetTaskId(v string) *CreateImageSplicingTaskResponseBody
	GetTaskId() *string
}

type CreateImageSplicingTaskResponseBody struct {
	// The event ID.
	//
	// example:
	//
	// 10C-1R6p7Km0H5Ieg38LKXTIvw*****
	EventId *string `json:"EventId,omitempty" xml:"EventId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 94D6F994-E298-037E-8E8B-0090F27*****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The task ID.
	//
	// example:
	//
	// ImageSplicing-537cc157-7645-444a-a631-c8db4d02*****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CreateImageSplicingTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateImageSplicingTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateImageSplicingTaskResponseBody) GetEventId() *string {
	return s.EventId
}

func (s *CreateImageSplicingTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateImageSplicingTaskResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *CreateImageSplicingTaskResponseBody) SetEventId(v string) *CreateImageSplicingTaskResponseBody {
	s.EventId = &v
	return s
}

func (s *CreateImageSplicingTaskResponseBody) SetRequestId(v string) *CreateImageSplicingTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateImageSplicingTaskResponseBody) SetTaskId(v string) *CreateImageSplicingTaskResponseBody {
	s.TaskId = &v
	return s
}

func (s *CreateImageSplicingTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
