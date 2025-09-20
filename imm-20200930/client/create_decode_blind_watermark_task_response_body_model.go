// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDecodeBlindWatermarkTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEventId(v string) *CreateDecodeBlindWatermarkTaskResponseBody
	GetEventId() *string
	SetRequestId(v string) *CreateDecodeBlindWatermarkTaskResponseBody
	GetRequestId() *string
	SetTaskId(v string) *CreateDecodeBlindWatermarkTaskResponseBody
	GetTaskId() *string
}

type CreateDecodeBlindWatermarkTaskResponseBody struct {
	// The event ID.
	//
	// example:
	//
	// 27C-1jyAP5qQI7RoI8lFFwvMrWtl0ft
	EventId *string `json:"EventId,omitempty" xml:"EventId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 4A7A2D0E-D8B8-4DA0-8127-EB32C6600ADE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The task ID.
	//
	// example:
	//
	// DecodeBlindWatermark-78ac8f3b-59e0-45a6-9b67-32168c3f22b9
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CreateDecodeBlindWatermarkTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDecodeBlindWatermarkTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDecodeBlindWatermarkTaskResponseBody) GetEventId() *string {
	return s.EventId
}

func (s *CreateDecodeBlindWatermarkTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDecodeBlindWatermarkTaskResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *CreateDecodeBlindWatermarkTaskResponseBody) SetEventId(v string) *CreateDecodeBlindWatermarkTaskResponseBody {
	s.EventId = &v
	return s
}

func (s *CreateDecodeBlindWatermarkTaskResponseBody) SetRequestId(v string) *CreateDecodeBlindWatermarkTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDecodeBlindWatermarkTaskResponseBody) SetTaskId(v string) *CreateDecodeBlindWatermarkTaskResponseBody {
	s.TaskId = &v
	return s
}

func (s *CreateDecodeBlindWatermarkTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
