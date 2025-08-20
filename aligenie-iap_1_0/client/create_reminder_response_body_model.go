// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateReminderResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *CreateReminderResponseBody
	GetErrorCode() *string
	SetErrorMsg(v string) *CreateReminderResponseBody
	GetErrorMsg() *string
	SetModel(v int64) *CreateReminderResponseBody
	GetModel() *int64
	SetSuccess(v bool) *CreateReminderResponseBody
	GetSuccess() *bool
}

type CreateReminderResponseBody struct {
	// example:
	//
	// 400
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// 不能设置过去的时间。
	ErrorMsg *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	// example:
	//
	// 20****1
	Model *int64 `json:"Model,omitempty" xml:"Model,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateReminderResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateReminderResponseBody) GoString() string {
	return s.String()
}

func (s *CreateReminderResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *CreateReminderResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *CreateReminderResponseBody) GetModel() *int64 {
	return s.Model
}

func (s *CreateReminderResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateReminderResponseBody) SetErrorCode(v string) *CreateReminderResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateReminderResponseBody) SetErrorMsg(v string) *CreateReminderResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *CreateReminderResponseBody) SetModel(v int64) *CreateReminderResponseBody {
	s.Model = &v
	return s
}

func (s *CreateReminderResponseBody) SetSuccess(v bool) *CreateReminderResponseBody {
	s.Success = &v
	return s
}

func (s *CreateReminderResponseBody) Validate() error {
	return dara.Validate(s)
}
