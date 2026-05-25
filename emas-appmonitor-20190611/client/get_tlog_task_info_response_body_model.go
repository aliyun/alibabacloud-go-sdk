// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTlogTaskInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v int64) *GetTlogTaskInfoResponseBody
	GetErrorCode() *int64
	SetMessage(v string) *GetTlogTaskInfoResponseBody
	GetMessage() *string
	SetModel(v interface{}) *GetTlogTaskInfoResponseBody
	GetModel() interface{}
	SetRequestId(v string) *GetTlogTaskInfoResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetTlogTaskInfoResponseBody
	GetSuccess() *bool
}

type GetTlogTaskInfoResponseBody struct {
	// example:
	//
	// 200
	ErrorCode *int64 `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// Successful
	Message *string     `json:"Message,omitempty" xml:"Message,omitempty"`
	Model   interface{} `json:"Model,omitempty" xml:"Model,omitempty"`
	// example:
	//
	// A8313212-EB4E-4E15-A7F9-D9C8F3FE8E94
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetTlogTaskInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTlogTaskInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetTlogTaskInfoResponseBody) GetErrorCode() *int64 {
	return s.ErrorCode
}

func (s *GetTlogTaskInfoResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetTlogTaskInfoResponseBody) GetModel() interface{} {
	return s.Model
}

func (s *GetTlogTaskInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTlogTaskInfoResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetTlogTaskInfoResponseBody) SetErrorCode(v int64) *GetTlogTaskInfoResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetTlogTaskInfoResponseBody) SetMessage(v string) *GetTlogTaskInfoResponseBody {
	s.Message = &v
	return s
}

func (s *GetTlogTaskInfoResponseBody) SetModel(v interface{}) *GetTlogTaskInfoResponseBody {
	s.Model = v
	return s
}

func (s *GetTlogTaskInfoResponseBody) SetRequestId(v string) *GetTlogTaskInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTlogTaskInfoResponseBody) SetSuccess(v bool) *GetTlogTaskInfoResponseBody {
	s.Success = &v
	return s
}

func (s *GetTlogTaskInfoResponseBody) Validate() error {
	return dara.Validate(s)
}
