// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGrantAgentUsersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GrantAgentUsersResponseBody
	GetCode() *string
	SetGrantedCount(v int64) *GrantAgentUsersResponseBody
	GetGrantedCount() *int64
	SetMessage(v string) *GrantAgentUsersResponseBody
	GetMessage() *string
	SetRequestId(v string) *GrantAgentUsersResponseBody
	GetRequestId() *string
}

type GrantAgentUsersResponseBody struct {
	// 业务状态码：成功为 200，失败为后端错误码（ERR.	- / InvalidParameter.*）
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 本次处理的授权记录数（含新增与更新）
	//
	// example:
	//
	// 1
	GrantedCount *int64 `json:"grantedCount,omitempty" xml:"grantedCount,omitempty"`
	// 错误描述，成功时为空
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// 请求追踪 ID
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GrantAgentUsersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GrantAgentUsersResponseBody) GoString() string {
	return s.String()
}

func (s *GrantAgentUsersResponseBody) GetCode() *string {
	return s.Code
}

func (s *GrantAgentUsersResponseBody) GetGrantedCount() *int64 {
	return s.GrantedCount
}

func (s *GrantAgentUsersResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GrantAgentUsersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GrantAgentUsersResponseBody) SetCode(v string) *GrantAgentUsersResponseBody {
	s.Code = &v
	return s
}

func (s *GrantAgentUsersResponseBody) SetGrantedCount(v int64) *GrantAgentUsersResponseBody {
	s.GrantedCount = &v
	return s
}

func (s *GrantAgentUsersResponseBody) SetMessage(v string) *GrantAgentUsersResponseBody {
	s.Message = &v
	return s
}

func (s *GrantAgentUsersResponseBody) SetRequestId(v string) *GrantAgentUsersResponseBody {
	s.RequestId = &v
	return s
}

func (s *GrantAgentUsersResponseBody) Validate() error {
	return dara.Validate(s)
}
