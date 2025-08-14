// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPocCreateTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *PocCreateTaskResponseBody
	GetCode() *string
	SetData(v string) *PocCreateTaskResponseBody
	GetData() *string
	SetHttpStatusCode(v string) *PocCreateTaskResponseBody
	GetHttpStatusCode() *string
	SetMessage(v string) *PocCreateTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *PocCreateTaskResponseBody
	GetRequestId() *string
	SetSuccess(v string) *PocCreateTaskResponseBody
	GetSuccess() *string
}

type PocCreateTaskResponseBody struct {
	// Status code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Task creation result
	//
	// example:
	//
	// {\\"Values\\": {\\"status\\": {\\"value\\": 4, \\"label\\": u\\"\\u5904\\u7406\\u4e2d\\"}, \\"bbs_ticket\\": True, \\"description\\": u\\"[LV-ERROR
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// HTTP status code
	//
	// example:
	//
	// 200
	HttpStatusCode *string `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Response message.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID.
	//
	// example:
	//
	// A32FE941-35F2-5378-B37C-4B8FDB16F094
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Whether the task was successful. **true*	- indicates success, **false*	- indicates failure.
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s PocCreateTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PocCreateTaskResponseBody) GoString() string {
	return s.String()
}

func (s *PocCreateTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *PocCreateTaskResponseBody) GetData() *string {
	return s.Data
}

func (s *PocCreateTaskResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *PocCreateTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *PocCreateTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PocCreateTaskResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *PocCreateTaskResponseBody) SetCode(v string) *PocCreateTaskResponseBody {
	s.Code = &v
	return s
}

func (s *PocCreateTaskResponseBody) SetData(v string) *PocCreateTaskResponseBody {
	s.Data = &v
	return s
}

func (s *PocCreateTaskResponseBody) SetHttpStatusCode(v string) *PocCreateTaskResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *PocCreateTaskResponseBody) SetMessage(v string) *PocCreateTaskResponseBody {
	s.Message = &v
	return s
}

func (s *PocCreateTaskResponseBody) SetRequestId(v string) *PocCreateTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *PocCreateTaskResponseBody) SetSuccess(v string) *PocCreateTaskResponseBody {
	s.Success = &v
	return s
}

func (s *PocCreateTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
