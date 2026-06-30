// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMiningTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateMiningTaskResponseBody
	GetCode() *string
	SetData(v *CreateMiningTaskResponseBodyData) *CreateMiningTaskResponseBody
	GetData() *CreateMiningTaskResponseBodyData
	SetHttpStatusCode(v string) *CreateMiningTaskResponseBody
	GetHttpStatusCode() *string
	SetMessage(v string) *CreateMiningTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateMiningTaskResponseBody
	GetRequestId() *string
	SetSuccess(v string) *CreateMiningTaskResponseBody
	GetSuccess() *string
}

type CreateMiningTaskResponseBody struct {
	// Result code. A value of **200*	- indicates success. Any other value indicates failure. Use this field to identify the cause of failure.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Response data
	Data *CreateMiningTaskResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// HTTP status code
	//
	// example:
	//
	// 200
	HttpStatusCode *string `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Error details if the call failed. Returns successful if the call succeeded.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID
	//
	// example:
	//
	// 106C6CA0-282D-4AF7-85F0-D2D24F4CE647
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call succeeded
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateMiningTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateMiningTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateMiningTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateMiningTaskResponseBody) GetData() *CreateMiningTaskResponseBodyData {
	return s.Data
}

func (s *CreateMiningTaskResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *CreateMiningTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateMiningTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateMiningTaskResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *CreateMiningTaskResponseBody) SetCode(v string) *CreateMiningTaskResponseBody {
	s.Code = &v
	return s
}

func (s *CreateMiningTaskResponseBody) SetData(v *CreateMiningTaskResponseBodyData) *CreateMiningTaskResponseBody {
	s.Data = v
	return s
}

func (s *CreateMiningTaskResponseBody) SetHttpStatusCode(v string) *CreateMiningTaskResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateMiningTaskResponseBody) SetMessage(v string) *CreateMiningTaskResponseBody {
	s.Message = &v
	return s
}

func (s *CreateMiningTaskResponseBody) SetRequestId(v string) *CreateMiningTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateMiningTaskResponseBody) SetSuccess(v string) *CreateMiningTaskResponseBody {
	s.Success = &v
	return s
}

func (s *CreateMiningTaskResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateMiningTaskResponseBodyData struct {
	// Task ID
	//
	// example:
	//
	// 7C1DEF5F-2C18-4D36-99C6-8C27*****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CreateMiningTaskResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateMiningTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateMiningTaskResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *CreateMiningTaskResponseBodyData) SetTaskId(v string) *CreateMiningTaskResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *CreateMiningTaskResponseBodyData) Validate() error {
	return dara.Validate(s)
}
