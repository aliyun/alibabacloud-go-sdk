// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetPasswordResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ResetPasswordResponseBody
	GetCode() *string
	SetMessage(v string) *ResetPasswordResponseBody
	GetMessage() *string
	SetRequestId(v string) *ResetPasswordResponseBody
	GetRequestId() *string
	SetWnUserId(v string) *ResetPasswordResponseBody
	GetWnUserId() *string
}

type ResetPasswordResponseBody struct {
	// 业务状态码：成功为 200，失败为后端错误码（ERR.	- / InvalidParameter.*）
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 错误描述，成功时为空
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// 请求追踪 ID
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// WINNEXO 平台用户ID
	//
	// example:
	//
	// 1
	WnUserId *string `json:"wnUserId,omitempty" xml:"wnUserId,omitempty"`
}

func (s ResetPasswordResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ResetPasswordResponseBody) GoString() string {
	return s.String()
}

func (s *ResetPasswordResponseBody) GetCode() *string {
	return s.Code
}

func (s *ResetPasswordResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ResetPasswordResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ResetPasswordResponseBody) GetWnUserId() *string {
	return s.WnUserId
}

func (s *ResetPasswordResponseBody) SetCode(v string) *ResetPasswordResponseBody {
	s.Code = &v
	return s
}

func (s *ResetPasswordResponseBody) SetMessage(v string) *ResetPasswordResponseBody {
	s.Message = &v
	return s
}

func (s *ResetPasswordResponseBody) SetRequestId(v string) *ResetPasswordResponseBody {
	s.RequestId = &v
	return s
}

func (s *ResetPasswordResponseBody) SetWnUserId(v string) *ResetPasswordResponseBody {
	s.WnUserId = &v
	return s
}

func (s *ResetPasswordResponseBody) Validate() error {
	return dara.Validate(s)
}
