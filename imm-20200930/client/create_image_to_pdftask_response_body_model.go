// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateImageToPDFTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEventId(v string) *CreateImageToPDFTaskResponseBody
	GetEventId() *string
	SetRequestId(v string) *CreateImageToPDFTaskResponseBody
	GetRequestId() *string
	SetTaskId(v string) *CreateImageToPDFTaskResponseBody
	GetTaskId() *string
}

type CreateImageToPDFTaskResponseBody struct {
	// The event ID.
	//
	// example:
	//
	// 0ED-1Bz8z71k5TtsUejT4UJ16Es*****
	EventId *string `json:"EventId,omitempty" xml:"EventId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// EC564A9A-BA5C-4499-A087-D9B9E76E*****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The task ID.
	//
	// example:
	//
	// ImageToPDF-cbe6ae3e-f8dc-4566-9da7-535d5d*****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CreateImageToPDFTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateImageToPDFTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateImageToPDFTaskResponseBody) GetEventId() *string {
	return s.EventId
}

func (s *CreateImageToPDFTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateImageToPDFTaskResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *CreateImageToPDFTaskResponseBody) SetEventId(v string) *CreateImageToPDFTaskResponseBody {
	s.EventId = &v
	return s
}

func (s *CreateImageToPDFTaskResponseBody) SetRequestId(v string) *CreateImageToPDFTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateImageToPDFTaskResponseBody) SetTaskId(v string) *CreateImageToPDFTaskResponseBody {
	s.TaskId = &v
	return s
}

func (s *CreateImageToPDFTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
