// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOfficeConversionTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEventId(v string) *CreateOfficeConversionTaskResponseBody
	GetEventId() *string
	SetRequestId(v string) *CreateOfficeConversionTaskResponseBody
	GetRequestId() *string
	SetTaskId(v string) *CreateOfficeConversionTaskResponseBody
	GetTaskId() *string
}

type CreateOfficeConversionTaskResponseBody struct {
	// The event ID.
	//
	// example:
	//
	// 2C2-1I0EG57VR37J4rQ8oKG6C9*****
	EventId *string `json:"EventId,omitempty" xml:"EventId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// FF3B7D81-66AE-47E0-BF69-157DCF18*****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The task ID.
	//
	// example:
	//
	// formatconvert-00bec802-073a-4b61-ba3b-39bc2fdd*****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CreateOfficeConversionTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateOfficeConversionTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateOfficeConversionTaskResponseBody) GetEventId() *string {
	return s.EventId
}

func (s *CreateOfficeConversionTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateOfficeConversionTaskResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *CreateOfficeConversionTaskResponseBody) SetEventId(v string) *CreateOfficeConversionTaskResponseBody {
	s.EventId = &v
	return s
}

func (s *CreateOfficeConversionTaskResponseBody) SetRequestId(v string) *CreateOfficeConversionTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateOfficeConversionTaskResponseBody) SetTaskId(v string) *CreateOfficeConversionTaskResponseBody {
	s.TaskId = &v
	return s
}

func (s *CreateOfficeConversionTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
