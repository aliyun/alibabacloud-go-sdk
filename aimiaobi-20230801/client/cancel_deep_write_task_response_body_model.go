// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelDeepWriteTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CancelDeepWriteTaskResponseBody
	GetCode() *string
	SetData(v *CancelDeepWriteTaskResponseBodyData) *CancelDeepWriteTaskResponseBody
	GetData() *CancelDeepWriteTaskResponseBodyData
	SetHttpStatusCode(v int32) *CancelDeepWriteTaskResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CancelDeepWriteTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *CancelDeepWriteTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CancelDeepWriteTaskResponseBody
	GetSuccess() *bool
}

type CancelDeepWriteTaskResponseBody struct {
	// example:
	//
	// DataNotExists
	Code *string                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *CancelDeepWriteTaskResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s CancelDeepWriteTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CancelDeepWriteTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CancelDeepWriteTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *CancelDeepWriteTaskResponseBody) GetData() *CancelDeepWriteTaskResponseBodyData {
	return s.Data
}

func (s *CancelDeepWriteTaskResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CancelDeepWriteTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CancelDeepWriteTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CancelDeepWriteTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CancelDeepWriteTaskResponseBody) SetCode(v string) *CancelDeepWriteTaskResponseBody {
	s.Code = &v
	return s
}

func (s *CancelDeepWriteTaskResponseBody) SetData(v *CancelDeepWriteTaskResponseBodyData) *CancelDeepWriteTaskResponseBody {
	s.Data = v
	return s
}

func (s *CancelDeepWriteTaskResponseBody) SetHttpStatusCode(v int32) *CancelDeepWriteTaskResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CancelDeepWriteTaskResponseBody) SetMessage(v string) *CancelDeepWriteTaskResponseBody {
	s.Message = &v
	return s
}

func (s *CancelDeepWriteTaskResponseBody) SetRequestId(v string) *CancelDeepWriteTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CancelDeepWriteTaskResponseBody) SetSuccess(v bool) *CancelDeepWriteTaskResponseBody {
	s.Success = &v
	return s
}

func (s *CancelDeepWriteTaskResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CancelDeepWriteTaskResponseBodyData struct {
	Input *string `json:"Input,omitempty" xml:"Input,omitempty"`
	// example:
	//
	// cancelled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// xbabac91-fdad-44d6-95ce-******
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CancelDeepWriteTaskResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CancelDeepWriteTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *CancelDeepWriteTaskResponseBodyData) GetInput() *string {
	return s.Input
}

func (s *CancelDeepWriteTaskResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *CancelDeepWriteTaskResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *CancelDeepWriteTaskResponseBodyData) SetInput(v string) *CancelDeepWriteTaskResponseBodyData {
	s.Input = &v
	return s
}

func (s *CancelDeepWriteTaskResponseBodyData) SetStatus(v string) *CancelDeepWriteTaskResponseBodyData {
	s.Status = &v
	return s
}

func (s *CancelDeepWriteTaskResponseBodyData) SetTaskId(v string) *CancelDeepWriteTaskResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *CancelDeepWriteTaskResponseBodyData) Validate() error {
	return dara.Validate(s)
}
