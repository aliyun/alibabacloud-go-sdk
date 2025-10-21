// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchContentAsyncDetectResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *BatchContentAsyncDetectResponseBody
	GetCode() *string
	SetData(v *BatchContentAsyncDetectResponseBodyData) *BatchContentAsyncDetectResponseBody
	GetData() *BatchContentAsyncDetectResponseBodyData
	SetHttpStatusCode(v int32) *BatchContentAsyncDetectResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *BatchContentAsyncDetectResponseBody
	GetMessage() *string
	SetRequestId(v string) *BatchContentAsyncDetectResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *BatchContentAsyncDetectResponseBody
	GetSuccess() *bool
}

type BatchContentAsyncDetectResponseBody struct {
	// example:
	//
	// 00000
	Code *string                                  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *BatchContentAsyncDetectResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// ""
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 9736C44E-F718-566B-821F-710216aAAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****F68692
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s BatchContentAsyncDetectResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BatchContentAsyncDetectResponseBody) GoString() string {
	return s.String()
}

func (s *BatchContentAsyncDetectResponseBody) GetCode() *string {
	return s.Code
}

func (s *BatchContentAsyncDetectResponseBody) GetData() *BatchContentAsyncDetectResponseBodyData {
	return s.Data
}

func (s *BatchContentAsyncDetectResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *BatchContentAsyncDetectResponseBody) GetMessage() *string {
	return s.Message
}

func (s *BatchContentAsyncDetectResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BatchContentAsyncDetectResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *BatchContentAsyncDetectResponseBody) SetCode(v string) *BatchContentAsyncDetectResponseBody {
	s.Code = &v
	return s
}

func (s *BatchContentAsyncDetectResponseBody) SetData(v *BatchContentAsyncDetectResponseBodyData) *BatchContentAsyncDetectResponseBody {
	s.Data = v
	return s
}

func (s *BatchContentAsyncDetectResponseBody) SetHttpStatusCode(v int32) *BatchContentAsyncDetectResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *BatchContentAsyncDetectResponseBody) SetMessage(v string) *BatchContentAsyncDetectResponseBody {
	s.Message = &v
	return s
}

func (s *BatchContentAsyncDetectResponseBody) SetRequestId(v string) *BatchContentAsyncDetectResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchContentAsyncDetectResponseBody) SetSuccess(v bool) *BatchContentAsyncDetectResponseBody {
	s.Success = &v
	return s
}

func (s *BatchContentAsyncDetectResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type BatchContentAsyncDetectResponseBodyData struct {
	// example:
	//
	// 19657954336
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s BatchContentAsyncDetectResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s BatchContentAsyncDetectResponseBodyData) GoString() string {
	return s.String()
}

func (s *BatchContentAsyncDetectResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *BatchContentAsyncDetectResponseBodyData) SetTaskId(v string) *BatchContentAsyncDetectResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *BatchContentAsyncDetectResponseBodyData) Validate() error {
	return dara.Validate(s)
}
