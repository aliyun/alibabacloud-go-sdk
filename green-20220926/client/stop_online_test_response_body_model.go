// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopOnlineTestResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *StopOnlineTestResponseBody
	GetRequestId() *string
	SetServiceCode(v string) *StopOnlineTestResponseBody
	GetServiceCode() *string
	SetTaskId(v string) *StopOnlineTestResponseBody
	GetTaskId() *string
	SetTaskStatus(v string) *StopOnlineTestResponseBody
	GetTaskStatus() *string
	SetUrl(v string) *StopOnlineTestResponseBody
	GetUrl() *string
}

type StopOnlineTestResponseBody struct {
	// The ID assigned by the backend to uniquely identify a request. You can use this ID to troubleshoot issues.
	//
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The service code.
	//
	// example:
	//
	// baselineCheck
	ServiceCode *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty"`
	// The ID of the detection task.
	//
	// example:
	//
	// xxxxx-xxxxx
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The detection status.
	//
	// example:
	//
	// SUCCESS
	TaskStatus *string `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	// The detection URL.
	//
	// example:
	//
	// https://xxxxxxxx.com/data/data.png
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s StopOnlineTestResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StopOnlineTestResponseBody) GoString() string {
	return s.String()
}

func (s *StopOnlineTestResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StopOnlineTestResponseBody) GetServiceCode() *string {
	return s.ServiceCode
}

func (s *StopOnlineTestResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *StopOnlineTestResponseBody) GetTaskStatus() *string {
	return s.TaskStatus
}

func (s *StopOnlineTestResponseBody) GetUrl() *string {
	return s.Url
}

func (s *StopOnlineTestResponseBody) SetRequestId(v string) *StopOnlineTestResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopOnlineTestResponseBody) SetServiceCode(v string) *StopOnlineTestResponseBody {
	s.ServiceCode = &v
	return s
}

func (s *StopOnlineTestResponseBody) SetTaskId(v string) *StopOnlineTestResponseBody {
	s.TaskId = &v
	return s
}

func (s *StopOnlineTestResponseBody) SetTaskStatus(v string) *StopOnlineTestResponseBody {
	s.TaskStatus = &v
	return s
}

func (s *StopOnlineTestResponseBody) SetUrl(v string) *StopOnlineTestResponseBody {
	s.Url = &v
	return s
}

func (s *StopOnlineTestResponseBody) Validate() error {
	return dara.Validate(s)
}
