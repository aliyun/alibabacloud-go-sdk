// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateUserResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateUserResponseBody
	GetCode() *string
	SetMessage(v string) *UpdateUserResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateUserResponseBody
	GetRequestId() *string
	SetWnUserId(v string) *UpdateUserResponseBody
	GetWnUserId() *string
}

type UpdateUserResponseBody struct {
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

func (s UpdateUserResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateUserResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateUserResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateUserResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateUserResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateUserResponseBody) GetWnUserId() *string {
	return s.WnUserId
}

func (s *UpdateUserResponseBody) SetCode(v string) *UpdateUserResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateUserResponseBody) SetMessage(v string) *UpdateUserResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateUserResponseBody) SetRequestId(v string) *UpdateUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateUserResponseBody) SetWnUserId(v string) *UpdateUserResponseBody {
	s.WnUserId = &v
	return s
}

func (s *UpdateUserResponseBody) Validate() error {
	return dara.Validate(s)
}
