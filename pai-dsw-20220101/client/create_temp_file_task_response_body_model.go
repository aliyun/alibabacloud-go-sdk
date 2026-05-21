// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTempFileTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateTempFileTaskResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *CreateTempFileTaskResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreateTempFileTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateTempFileTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateTempFileTaskResponseBody
	GetSuccess() *bool
	SetTempFileTaskId(v string) *CreateTempFileTaskResponseBody
	GetTempFileTaskId() *string
}

type CreateTempFileTaskResponseBody struct {
	// example:
	//
	// "200"
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// "XXX"
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// E7D55162-4489-1619-AAF5-3F97D5FCA948
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// task-05cexxxxxxxxx
	TempFileTaskId *string `json:"TempFileTaskId,omitempty" xml:"TempFileTaskId,omitempty"`
}

func (s CreateTempFileTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateTempFileTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTempFileTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateTempFileTaskResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateTempFileTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateTempFileTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateTempFileTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateTempFileTaskResponseBody) GetTempFileTaskId() *string {
	return s.TempFileTaskId
}

func (s *CreateTempFileTaskResponseBody) SetCode(v string) *CreateTempFileTaskResponseBody {
	s.Code = &v
	return s
}

func (s *CreateTempFileTaskResponseBody) SetHttpStatusCode(v int32) *CreateTempFileTaskResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateTempFileTaskResponseBody) SetMessage(v string) *CreateTempFileTaskResponseBody {
	s.Message = &v
	return s
}

func (s *CreateTempFileTaskResponseBody) SetRequestId(v string) *CreateTempFileTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateTempFileTaskResponseBody) SetSuccess(v bool) *CreateTempFileTaskResponseBody {
	s.Success = &v
	return s
}

func (s *CreateTempFileTaskResponseBody) SetTempFileTaskId(v string) *CreateTempFileTaskResponseBody {
	s.TempFileTaskId = &v
	return s
}

func (s *CreateTempFileTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
