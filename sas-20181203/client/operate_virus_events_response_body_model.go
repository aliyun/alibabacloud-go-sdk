// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOperateVirusEventsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *OperateVirusEventsResponseBody
	GetRequestId() *string
	SetTaskId(v int64) *OperateVirusEventsResponseBody
	GetTaskId() *int64
}

type OperateVirusEventsResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// B72BEC03-001C-5C77-A4BB-1E6XXXXXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the task to handle multiple alert events at a time.
	//
	// example:
	//
	// 1
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s OperateVirusEventsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s OperateVirusEventsResponseBody) GoString() string {
	return s.String()
}

func (s *OperateVirusEventsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *OperateVirusEventsResponseBody) GetTaskId() *int64 {
	return s.TaskId
}

func (s *OperateVirusEventsResponseBody) SetRequestId(v string) *OperateVirusEventsResponseBody {
	s.RequestId = &v
	return s
}

func (s *OperateVirusEventsResponseBody) SetTaskId(v int64) *OperateVirusEventsResponseBody {
	s.TaskId = &v
	return s
}

func (s *OperateVirusEventsResponseBody) Validate() error {
	return dara.Validate(s)
}
