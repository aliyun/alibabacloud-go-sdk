// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitDeepWriteTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SubmitDeepWriteTaskResponseBody
	GetCode() *string
	SetData(v *SubmitDeepWriteTaskResponseBodyData) *SubmitDeepWriteTaskResponseBody
	GetData() *SubmitDeepWriteTaskResponseBodyData
	SetHttpStatusCode(v int32) *SubmitDeepWriteTaskResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *SubmitDeepWriteTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *SubmitDeepWriteTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SubmitDeepWriteTaskResponseBody
	GetSuccess() *bool
}

type SubmitDeepWriteTaskResponseBody struct {
	// example:
	//
	// DataNotExists
	Code *string                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *SubmitDeepWriteTaskResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 400
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// 错误消息
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// xxxxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SubmitDeepWriteTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitDeepWriteTaskResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitDeepWriteTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *SubmitDeepWriteTaskResponseBody) GetData() *SubmitDeepWriteTaskResponseBodyData {
	return s.Data
}

func (s *SubmitDeepWriteTaskResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *SubmitDeepWriteTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SubmitDeepWriteTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitDeepWriteTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SubmitDeepWriteTaskResponseBody) SetCode(v string) *SubmitDeepWriteTaskResponseBody {
	s.Code = &v
	return s
}

func (s *SubmitDeepWriteTaskResponseBody) SetData(v *SubmitDeepWriteTaskResponseBodyData) *SubmitDeepWriteTaskResponseBody {
	s.Data = v
	return s
}

func (s *SubmitDeepWriteTaskResponseBody) SetHttpStatusCode(v int32) *SubmitDeepWriteTaskResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *SubmitDeepWriteTaskResponseBody) SetMessage(v string) *SubmitDeepWriteTaskResponseBody {
	s.Message = &v
	return s
}

func (s *SubmitDeepWriteTaskResponseBody) SetRequestId(v string) *SubmitDeepWriteTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitDeepWriteTaskResponseBody) SetSuccess(v bool) *SubmitDeepWriteTaskResponseBody {
	s.Success = &v
	return s
}

func (s *SubmitDeepWriteTaskResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SubmitDeepWriteTaskResponseBodyData struct {
	// example:
	//
	// queued
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// xbabac91-fdad-44d6-95ce-******
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s SubmitDeepWriteTaskResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s SubmitDeepWriteTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *SubmitDeepWriteTaskResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *SubmitDeepWriteTaskResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *SubmitDeepWriteTaskResponseBodyData) SetStatus(v string) *SubmitDeepWriteTaskResponseBodyData {
	s.Status = &v
	return s
}

func (s *SubmitDeepWriteTaskResponseBodyData) SetTaskId(v string) *SubmitDeepWriteTaskResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *SubmitDeepWriteTaskResponseBodyData) Validate() error {
	return dara.Validate(s)
}
