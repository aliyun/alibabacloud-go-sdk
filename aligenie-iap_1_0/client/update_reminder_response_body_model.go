// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateReminderResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v int32) *UpdateReminderResponseBody
	GetErrorCode() *int32
	SetErrorMsg(v string) *UpdateReminderResponseBody
	GetErrorMsg() *string
	SetModel(v int64) *UpdateReminderResponseBody
	GetModel() *int64
	SetSuccess(v bool) *UpdateReminderResponseBody
	GetSuccess() *bool
}

type UpdateReminderResponseBody struct {
	// example:
	//
	// 400
	ErrorCode *int32 `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
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

func (s UpdateReminderResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateReminderResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateReminderResponseBody) GetErrorCode() *int32 {
	return s.ErrorCode
}

func (s *UpdateReminderResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *UpdateReminderResponseBody) GetModel() *int64 {
	return s.Model
}

func (s *UpdateReminderResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateReminderResponseBody) SetErrorCode(v int32) *UpdateReminderResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateReminderResponseBody) SetErrorMsg(v string) *UpdateReminderResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *UpdateReminderResponseBody) SetModel(v int64) *UpdateReminderResponseBody {
	s.Model = &v
	return s
}

func (s *UpdateReminderResponseBody) SetSuccess(v bool) *UpdateReminderResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateReminderResponseBody) Validate() error {
	return dara.Validate(s)
}
