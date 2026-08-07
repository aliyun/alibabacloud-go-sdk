// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOnlineTestResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateOnlineTestResponseBody
	GetRequestId() *string
	SetServiceCode(v string) *CreateOnlineTestResponseBody
	GetServiceCode() *string
	SetTaskId(v string) *CreateOnlineTestResponseBody
	GetTaskId() *string
	SetTaskStatus(v string) *CreateOnlineTestResponseBody
	GetTaskStatus() *string
	SetUrl(v string) *CreateOnlineTestResponseBody
	GetUrl() *string
}

type CreateOnlineTestResponseBody struct {
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
	// VideoModeration
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
	// The URL to be detected.
	//
	// example:
	//
	// https://xxxxxxxxxx.com/data/data.png
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s CreateOnlineTestResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateOnlineTestResponseBody) GoString() string {
	return s.String()
}

func (s *CreateOnlineTestResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateOnlineTestResponseBody) GetServiceCode() *string {
	return s.ServiceCode
}

func (s *CreateOnlineTestResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *CreateOnlineTestResponseBody) GetTaskStatus() *string {
	return s.TaskStatus
}

func (s *CreateOnlineTestResponseBody) GetUrl() *string {
	return s.Url
}

func (s *CreateOnlineTestResponseBody) SetRequestId(v string) *CreateOnlineTestResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateOnlineTestResponseBody) SetServiceCode(v string) *CreateOnlineTestResponseBody {
	s.ServiceCode = &v
	return s
}

func (s *CreateOnlineTestResponseBody) SetTaskId(v string) *CreateOnlineTestResponseBody {
	s.TaskId = &v
	return s
}

func (s *CreateOnlineTestResponseBody) SetTaskStatus(v string) *CreateOnlineTestResponseBody {
	s.TaskStatus = &v
	return s
}

func (s *CreateOnlineTestResponseBody) SetUrl(v string) *CreateOnlineTestResponseBody {
	s.Url = &v
	return s
}

func (s *CreateOnlineTestResponseBody) Validate() error {
	return dara.Validate(s)
}
