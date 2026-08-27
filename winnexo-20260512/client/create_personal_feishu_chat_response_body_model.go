// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePersonalFeishuChatResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetChatId(v string) *CreatePersonalFeishuChatResponseBody
	GetChatId() *string
	SetCode(v string) *CreatePersonalFeishuChatResponseBody
	GetCode() *string
	SetDirectoryId(v string) *CreatePersonalFeishuChatResponseBody
	GetDirectoryId() *string
	SetGmtCreate(v string) *CreatePersonalFeishuChatResponseBody
	GetGmtCreate() *string
	SetMessage(v string) *CreatePersonalFeishuChatResponseBody
	GetMessage() *string
	SetName(v string) *CreatePersonalFeishuChatResponseBody
	GetName() *string
	SetRequestId(v string) *CreatePersonalFeishuChatResponseBody
	GetRequestId() *string
	SetScope(v string) *CreatePersonalFeishuChatResponseBody
	GetScope() *string
	SetSourceId(v string) *CreatePersonalFeishuChatResponseBody
	GetSourceId() *string
	SetStatus(v string) *CreatePersonalFeishuChatResponseBody
	GetStatus() *string
}

type CreatePersonalFeishuChatResponseBody struct {
	// The group chat session ID.
	//
	// example:
	//
	// oc_abc123
	ChatId *string `json:"chatId,omitempty" xml:"chatId,omitempty"`
	// The status code.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The directory ID.
	//
	// example:
	//
	// dir_personal_1
	DirectoryId *string `json:"directoryId,omitempty" xml:"directoryId,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 2026-08-18T10:30:00+08:00
	GmtCreate *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// The status code description.
	//
	// example:
	//
	// successful
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The skill name.
	//
	// example:
	//
	// Product R&D Group
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The permission scope.
	//
	// example:
	//
	// PERSONAL
	Scope *string `json:"scope,omitempty" xml:"scope,omitempty"`
	// The source ID.
	//
	// example:
	//
	// src_feishu_1
	SourceId *string `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
	// The signing status. Valid values:
	//
	// - CREATED: Created but not signed.
	//
	// - SUCCESS: Signed successfully.
	//
	// - STOP: Terminated.
	//
	// example:
	//
	// RUNNING
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s CreatePersonalFeishuChatResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreatePersonalFeishuChatResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePersonalFeishuChatResponseBody) GetChatId() *string {
	return s.ChatId
}

func (s *CreatePersonalFeishuChatResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreatePersonalFeishuChatResponseBody) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *CreatePersonalFeishuChatResponseBody) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *CreatePersonalFeishuChatResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreatePersonalFeishuChatResponseBody) GetName() *string {
	return s.Name
}

func (s *CreatePersonalFeishuChatResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreatePersonalFeishuChatResponseBody) GetScope() *string {
	return s.Scope
}

func (s *CreatePersonalFeishuChatResponseBody) GetSourceId() *string {
	return s.SourceId
}

func (s *CreatePersonalFeishuChatResponseBody) GetStatus() *string {
	return s.Status
}

func (s *CreatePersonalFeishuChatResponseBody) SetChatId(v string) *CreatePersonalFeishuChatResponseBody {
	s.ChatId = &v
	return s
}

func (s *CreatePersonalFeishuChatResponseBody) SetCode(v string) *CreatePersonalFeishuChatResponseBody {
	s.Code = &v
	return s
}

func (s *CreatePersonalFeishuChatResponseBody) SetDirectoryId(v string) *CreatePersonalFeishuChatResponseBody {
	s.DirectoryId = &v
	return s
}

func (s *CreatePersonalFeishuChatResponseBody) SetGmtCreate(v string) *CreatePersonalFeishuChatResponseBody {
	s.GmtCreate = &v
	return s
}

func (s *CreatePersonalFeishuChatResponseBody) SetMessage(v string) *CreatePersonalFeishuChatResponseBody {
	s.Message = &v
	return s
}

func (s *CreatePersonalFeishuChatResponseBody) SetName(v string) *CreatePersonalFeishuChatResponseBody {
	s.Name = &v
	return s
}

func (s *CreatePersonalFeishuChatResponseBody) SetRequestId(v string) *CreatePersonalFeishuChatResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreatePersonalFeishuChatResponseBody) SetScope(v string) *CreatePersonalFeishuChatResponseBody {
	s.Scope = &v
	return s
}

func (s *CreatePersonalFeishuChatResponseBody) SetSourceId(v string) *CreatePersonalFeishuChatResponseBody {
	s.SourceId = &v
	return s
}

func (s *CreatePersonalFeishuChatResponseBody) SetStatus(v string) *CreatePersonalFeishuChatResponseBody {
	s.Status = &v
	return s
}

func (s *CreatePersonalFeishuChatResponseBody) Validate() error {
	return dara.Validate(s)
}
