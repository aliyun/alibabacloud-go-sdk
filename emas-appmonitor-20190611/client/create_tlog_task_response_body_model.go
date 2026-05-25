// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTlogTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v int64) *CreateTlogTaskResponseBody
	GetErrorCode() *int64
	SetMessage(v string) *CreateTlogTaskResponseBody
	GetMessage() *string
	SetModel(v string) *CreateTlogTaskResponseBody
	GetModel() *string
	SetRequestId(v string) *CreateTlogTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateTlogTaskResponseBody
	GetSuccess() *bool
}

type CreateTlogTaskResponseBody struct {
	// example:
	//
	// 200
	ErrorCode *int64 `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// Successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 1779434100158kg2B0tZvQOmR1Yn_p_8usw
	Model *string `json:"Model,omitempty" xml:"Model,omitempty"`
	// example:
	//
	// A8313212-EB4E-4E15-A7F9-D9C8F3FE8E94
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateTlogTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateTlogTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTlogTaskResponseBody) GetErrorCode() *int64 {
	return s.ErrorCode
}

func (s *CreateTlogTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateTlogTaskResponseBody) GetModel() *string {
	return s.Model
}

func (s *CreateTlogTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateTlogTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateTlogTaskResponseBody) SetErrorCode(v int64) *CreateTlogTaskResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateTlogTaskResponseBody) SetMessage(v string) *CreateTlogTaskResponseBody {
	s.Message = &v
	return s
}

func (s *CreateTlogTaskResponseBody) SetModel(v string) *CreateTlogTaskResponseBody {
	s.Model = &v
	return s
}

func (s *CreateTlogTaskResponseBody) SetRequestId(v string) *CreateTlogTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateTlogTaskResponseBody) SetSuccess(v bool) *CreateTlogTaskResponseBody {
	s.Success = &v
	return s
}

func (s *CreateTlogTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
