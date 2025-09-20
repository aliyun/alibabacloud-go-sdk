// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCompressPointCloudTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEventId(v string) *CreateCompressPointCloudTaskResponseBody
	GetEventId() *string
	SetRequestId(v string) *CreateCompressPointCloudTaskResponseBody
	GetRequestId() *string
	SetTaskId(v string) *CreateCompressPointCloudTaskResponseBody
	GetTaskId() *string
}

type CreateCompressPointCloudTaskResponseBody struct {
	// The event ID.
	//
	// example:
	//
	// 0B7-1LR4Wcue1aBhk2xT85MfL*****
	EventId *string `json:"EventId,omitempty" xml:"EventId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// CA995EFD-083D-4F40-BE8A-BDF75FFF****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The task ID.
	//
	// example:
	//
	// PointCloudCompress-badda57d-a3ab-4e6d-938f-49b77ce****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CreateCompressPointCloudTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateCompressPointCloudTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCompressPointCloudTaskResponseBody) GetEventId() *string {
	return s.EventId
}

func (s *CreateCompressPointCloudTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateCompressPointCloudTaskResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *CreateCompressPointCloudTaskResponseBody) SetEventId(v string) *CreateCompressPointCloudTaskResponseBody {
	s.EventId = &v
	return s
}

func (s *CreateCompressPointCloudTaskResponseBody) SetRequestId(v string) *CreateCompressPointCloudTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateCompressPointCloudTaskResponseBody) SetTaskId(v string) *CreateCompressPointCloudTaskResponseBody {
	s.TaskId = &v
	return s
}

func (s *CreateCompressPointCloudTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
