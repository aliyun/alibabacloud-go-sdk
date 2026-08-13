// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevokeAgentUsersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *RevokeAgentUsersResponseBody
	GetCode() *string
	SetMessage(v string) *RevokeAgentUsersResponseBody
	GetMessage() *string
	SetRequestId(v string) *RevokeAgentUsersResponseBody
	GetRequestId() *string
	SetRevokedCount(v int64) *RevokeAgentUsersResponseBody
	GetRevokedCount() *int64
}

type RevokeAgentUsersResponseBody struct {
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
	// 本次成功撤销的记录数
	//
	// example:
	//
	// 1
	RevokedCount *int64 `json:"revokedCount,omitempty" xml:"revokedCount,omitempty"`
}

func (s RevokeAgentUsersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RevokeAgentUsersResponseBody) GoString() string {
	return s.String()
}

func (s *RevokeAgentUsersResponseBody) GetCode() *string {
	return s.Code
}

func (s *RevokeAgentUsersResponseBody) GetMessage() *string {
	return s.Message
}

func (s *RevokeAgentUsersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RevokeAgentUsersResponseBody) GetRevokedCount() *int64 {
	return s.RevokedCount
}

func (s *RevokeAgentUsersResponseBody) SetCode(v string) *RevokeAgentUsersResponseBody {
	s.Code = &v
	return s
}

func (s *RevokeAgentUsersResponseBody) SetMessage(v string) *RevokeAgentUsersResponseBody {
	s.Message = &v
	return s
}

func (s *RevokeAgentUsersResponseBody) SetRequestId(v string) *RevokeAgentUsersResponseBody {
	s.RequestId = &v
	return s
}

func (s *RevokeAgentUsersResponseBody) SetRevokedCount(v int64) *RevokeAgentUsersResponseBody {
	s.RevokedCount = &v
	return s
}

func (s *RevokeAgentUsersResponseBody) Validate() error {
	return dara.Validate(s)
}
