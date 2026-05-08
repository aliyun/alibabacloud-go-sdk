// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryImageToVideoTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessage(v string) *QueryImageToVideoTaskResponseBody
	GetMessage() *string
	SetOriginUrl(v string) *QueryImageToVideoTaskResponseBody
	GetOriginUrl() *string
	SetRequestId(v string) *QueryImageToVideoTaskResponseBody
	GetRequestId() *string
	SetStatus(v int32) *QueryImageToVideoTaskResponseBody
	GetStatus() *int32
	SetSuccess(v bool) *QueryImageToVideoTaskResponseBody
	GetSuccess() *bool
	SetTaskId(v string) *QueryImageToVideoTaskResponseBody
	GetTaskId() *string
}

type QueryImageToVideoTaskResponseBody struct {
	// example:
	//
	// None
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// https://xxx/xxx.mp4
	OriginUrl *string `json:"originUrl,omitempty" xml:"originUrl,omitempty"`
	// example:
	//
	// CC2967CA-0114-57E0-A0CF-7DEEEDAB953D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 1
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 868125994191405056
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s QueryImageToVideoTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryImageToVideoTaskResponseBody) GoString() string {
	return s.String()
}

func (s *QueryImageToVideoTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryImageToVideoTaskResponseBody) GetOriginUrl() *string {
	return s.OriginUrl
}

func (s *QueryImageToVideoTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryImageToVideoTaskResponseBody) GetStatus() *int32 {
	return s.Status
}

func (s *QueryImageToVideoTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryImageToVideoTaskResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *QueryImageToVideoTaskResponseBody) SetMessage(v string) *QueryImageToVideoTaskResponseBody {
	s.Message = &v
	return s
}

func (s *QueryImageToVideoTaskResponseBody) SetOriginUrl(v string) *QueryImageToVideoTaskResponseBody {
	s.OriginUrl = &v
	return s
}

func (s *QueryImageToVideoTaskResponseBody) SetRequestId(v string) *QueryImageToVideoTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryImageToVideoTaskResponseBody) SetStatus(v int32) *QueryImageToVideoTaskResponseBody {
	s.Status = &v
	return s
}

func (s *QueryImageToVideoTaskResponseBody) SetSuccess(v bool) *QueryImageToVideoTaskResponseBody {
	s.Success = &v
	return s
}

func (s *QueryImageToVideoTaskResponseBody) SetTaskId(v string) *QueryImageToVideoTaskResponseBody {
	s.TaskId = &v
	return s
}

func (s *QueryImageToVideoTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
