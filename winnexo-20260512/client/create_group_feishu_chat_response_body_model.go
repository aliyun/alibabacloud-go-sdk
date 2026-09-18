// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateGroupFeishuChatResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetChatId(v string) *CreateGroupFeishuChatResponseBody
	GetChatId() *string
	SetCode(v string) *CreateGroupFeishuChatResponseBody
	GetCode() *string
	SetDirectoryId(v string) *CreateGroupFeishuChatResponseBody
	GetDirectoryId() *string
	SetGmtCreate(v string) *CreateGroupFeishuChatResponseBody
	GetGmtCreate() *string
	SetGroupId(v string) *CreateGroupFeishuChatResponseBody
	GetGroupId() *string
	SetMessage(v string) *CreateGroupFeishuChatResponseBody
	GetMessage() *string
	SetName(v string) *CreateGroupFeishuChatResponseBody
	GetName() *string
	SetRequestId(v string) *CreateGroupFeishuChatResponseBody
	GetRequestId() *string
	SetScope(v string) *CreateGroupFeishuChatResponseBody
	GetScope() *string
	SetSourceId(v string) *CreateGroupFeishuChatResponseBody
	GetSourceId() *string
	SetStatus(v string) *CreateGroupFeishuChatResponseBody
	GetStatus() *string
}

type CreateGroupFeishuChatResponseBody struct {
	// 飞书群聊ID
	//
	// example:
	//
	// cidxxxxxxxx
	ChatId *string `json:"chatId,omitempty" xml:"chatId,omitempty"`
	// 业务状态码
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 解析并绑定的真实目录ID
	//
	// example:
	//
	// exampleDirectoryId
	DirectoryId *string `json:"directoryId,omitempty" xml:"directoryId,omitempty"`
	// 创建时间，ISO8601格式
	//
	// example:
	//
	// 2026-08-26T10:00:00+08:00
	GmtCreate *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// 协作空间ID
	//
	// example:
	//
	// exampleGroupId
	GroupId *string `json:"groupId,omitempty" xml:"groupId,omitempty"`
	// 错误描述
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Provider处理后的实际资料名称
	//
	// example:
	//
	// oklabs_tongyici
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 请求追踪ID
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// 资料范围，固定GROUP
	//
	// example:
	//
	// PERSONAL
	Scope *string `json:"scope,omitempty" xml:"scope,omitempty"`
	// 新建资料ID
	//
	// example:
	//
	// exampleSourceId
	SourceId *string `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
	// 实际资料状态；RUNNING表示处理中，FAILED表示创建处理失败
	//
	// example:
	//
	// PENDING
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s CreateGroupFeishuChatResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateGroupFeishuChatResponseBody) GoString() string {
	return s.String()
}

func (s *CreateGroupFeishuChatResponseBody) GetChatId() *string {
	return s.ChatId
}

func (s *CreateGroupFeishuChatResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateGroupFeishuChatResponseBody) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *CreateGroupFeishuChatResponseBody) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *CreateGroupFeishuChatResponseBody) GetGroupId() *string {
	return s.GroupId
}

func (s *CreateGroupFeishuChatResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateGroupFeishuChatResponseBody) GetName() *string {
	return s.Name
}

func (s *CreateGroupFeishuChatResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateGroupFeishuChatResponseBody) GetScope() *string {
	return s.Scope
}

func (s *CreateGroupFeishuChatResponseBody) GetSourceId() *string {
	return s.SourceId
}

func (s *CreateGroupFeishuChatResponseBody) GetStatus() *string {
	return s.Status
}

func (s *CreateGroupFeishuChatResponseBody) SetChatId(v string) *CreateGroupFeishuChatResponseBody {
	s.ChatId = &v
	return s
}

func (s *CreateGroupFeishuChatResponseBody) SetCode(v string) *CreateGroupFeishuChatResponseBody {
	s.Code = &v
	return s
}

func (s *CreateGroupFeishuChatResponseBody) SetDirectoryId(v string) *CreateGroupFeishuChatResponseBody {
	s.DirectoryId = &v
	return s
}

func (s *CreateGroupFeishuChatResponseBody) SetGmtCreate(v string) *CreateGroupFeishuChatResponseBody {
	s.GmtCreate = &v
	return s
}

func (s *CreateGroupFeishuChatResponseBody) SetGroupId(v string) *CreateGroupFeishuChatResponseBody {
	s.GroupId = &v
	return s
}

func (s *CreateGroupFeishuChatResponseBody) SetMessage(v string) *CreateGroupFeishuChatResponseBody {
	s.Message = &v
	return s
}

func (s *CreateGroupFeishuChatResponseBody) SetName(v string) *CreateGroupFeishuChatResponseBody {
	s.Name = &v
	return s
}

func (s *CreateGroupFeishuChatResponseBody) SetRequestId(v string) *CreateGroupFeishuChatResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateGroupFeishuChatResponseBody) SetScope(v string) *CreateGroupFeishuChatResponseBody {
	s.Scope = &v
	return s
}

func (s *CreateGroupFeishuChatResponseBody) SetSourceId(v string) *CreateGroupFeishuChatResponseBody {
	s.SourceId = &v
	return s
}

func (s *CreateGroupFeishuChatResponseBody) SetStatus(v string) *CreateGroupFeishuChatResponseBody {
	s.Status = &v
	return s
}

func (s *CreateGroupFeishuChatResponseBody) Validate() error {
	return dara.Validate(s)
}
