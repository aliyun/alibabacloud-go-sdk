// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteReminderResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v int32) *DeleteReminderResponseBody
	GetErrorCode() *int32
	SetErrorMsg(v string) *DeleteReminderResponseBody
	GetErrorMsg() *string
	SetSuccess(v bool) *DeleteReminderResponseBody
	GetSuccess() *bool
}

type DeleteReminderResponseBody struct {
	// example:
	//
	// 400
	ErrorCode *int32 `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// 参数错误。
	ErrorMsg *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteReminderResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteReminderResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteReminderResponseBody) GetErrorCode() *int32 {
	return s.ErrorCode
}

func (s *DeleteReminderResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *DeleteReminderResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteReminderResponseBody) SetErrorCode(v int32) *DeleteReminderResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteReminderResponseBody) SetErrorMsg(v string) *DeleteReminderResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *DeleteReminderResponseBody) SetSuccess(v bool) *DeleteReminderResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteReminderResponseBody) Validate() error {
	return dara.Validate(s)
}
