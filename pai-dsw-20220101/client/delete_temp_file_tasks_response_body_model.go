// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTempFileTasksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteTempFileTasksResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *DeleteTempFileTasksResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DeleteTempFileTasksResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteTempFileTasksResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteTempFileTasksResponseBody
	GetSuccess() *bool
}

type DeleteTempFileTasksResponseBody struct {
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
}

func (s DeleteTempFileTasksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteTempFileTasksResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteTempFileTasksResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteTempFileTasksResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteTempFileTasksResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteTempFileTasksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteTempFileTasksResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteTempFileTasksResponseBody) SetCode(v string) *DeleteTempFileTasksResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteTempFileTasksResponseBody) SetHttpStatusCode(v int32) *DeleteTempFileTasksResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteTempFileTasksResponseBody) SetMessage(v string) *DeleteTempFileTasksResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteTempFileTasksResponseBody) SetRequestId(v string) *DeleteTempFileTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteTempFileTasksResponseBody) SetSuccess(v bool) *DeleteTempFileTasksResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteTempFileTasksResponseBody) Validate() error {
	return dara.Validate(s)
}
