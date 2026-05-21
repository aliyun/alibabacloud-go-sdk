// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTempFileTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteTempFileTaskResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *DeleteTempFileTaskResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DeleteTempFileTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteTempFileTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteTempFileTaskResponseBody
	GetSuccess() *bool
	SetTempFileTaskId(v string) *DeleteTempFileTaskResponseBody
	GetTempFileTaskId() *string
}

type DeleteTempFileTaskResponseBody struct {
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

func (s DeleteTempFileTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteTempFileTaskResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteTempFileTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteTempFileTaskResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteTempFileTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteTempFileTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteTempFileTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteTempFileTaskResponseBody) GetTempFileTaskId() *string {
	return s.TempFileTaskId
}

func (s *DeleteTempFileTaskResponseBody) SetCode(v string) *DeleteTempFileTaskResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteTempFileTaskResponseBody) SetHttpStatusCode(v int32) *DeleteTempFileTaskResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteTempFileTaskResponseBody) SetMessage(v string) *DeleteTempFileTaskResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteTempFileTaskResponseBody) SetRequestId(v string) *DeleteTempFileTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteTempFileTaskResponseBody) SetSuccess(v bool) *DeleteTempFileTaskResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteTempFileTaskResponseBody) SetTempFileTaskId(v string) *DeleteTempFileTaskResponseBody {
	s.TempFileTaskId = &v
	return s
}

func (s *DeleteTempFileTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
