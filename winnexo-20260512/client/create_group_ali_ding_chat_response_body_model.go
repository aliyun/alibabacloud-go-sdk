// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateGroupAliDingChatResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetChatId(v string) *CreateGroupAliDingChatResponseBody
	GetChatId() *string
	SetCode(v string) *CreateGroupAliDingChatResponseBody
	GetCode() *string
	SetDirectoryId(v string) *CreateGroupAliDingChatResponseBody
	GetDirectoryId() *string
	SetGmtCreate(v string) *CreateGroupAliDingChatResponseBody
	GetGmtCreate() *string
	SetGroupId(v string) *CreateGroupAliDingChatResponseBody
	GetGroupId() *string
	SetMessage(v string) *CreateGroupAliDingChatResponseBody
	GetMessage() *string
	SetName(v string) *CreateGroupAliDingChatResponseBody
	GetName() *string
	SetRequestId(v string) *CreateGroupAliDingChatResponseBody
	GetRequestId() *string
	SetScope(v string) *CreateGroupAliDingChatResponseBody
	GetScope() *string
	SetSourceId(v string) *CreateGroupAliDingChatResponseBody
	GetSourceId() *string
	SetStatus(v string) *CreateGroupAliDingChatResponseBody
	GetStatus() *string
}

type CreateGroupAliDingChatResponseBody struct {
	// The session ID, typically used for JSSDK.
	//
	// example:
	//
	// cidxxxxxxxx
	ChatId *string `json:"chatId,omitempty" xml:"chatId,omitempty"`
	// The response code.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The directory ID.
	//
	// example:
	//
	// exampleDirectoryId
	DirectoryId *string `json:"directoryId,omitempty" xml:"directoryId,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 2026-08-18T10:00:00Z
	GmtCreate *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// The project group ID.
	//
	// example:
	//
	// exampleGroupId
	GroupId *string `json:"groupId,omitempty" xml:"groupId,omitempty"`
	// The error details.
	//
	// example:
	//
	// ok
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The bot name.
	//
	// example:
	//
	// CustomerProjectGroup
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
	// GROUP
	Scope *string `json:"scope,omitempty" xml:"scope,omitempty"`
	// The unique identifier on the business system side, that is, the business ID.
	//
	// example:
	//
	// exampleSourceId
	SourceId *string `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
	// The signing status. Valid values:
	//
	// - CREATED: Created but not signed.
	//
	// - SUCCESS: Signed.
	//
	// - STOP: Terminated.
	//
	// example:
	//
	// PENDING
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s CreateGroupAliDingChatResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateGroupAliDingChatResponseBody) GoString() string {
	return s.String()
}

func (s *CreateGroupAliDingChatResponseBody) GetChatId() *string {
	return s.ChatId
}

func (s *CreateGroupAliDingChatResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateGroupAliDingChatResponseBody) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *CreateGroupAliDingChatResponseBody) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *CreateGroupAliDingChatResponseBody) GetGroupId() *string {
	return s.GroupId
}

func (s *CreateGroupAliDingChatResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateGroupAliDingChatResponseBody) GetName() *string {
	return s.Name
}

func (s *CreateGroupAliDingChatResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateGroupAliDingChatResponseBody) GetScope() *string {
	return s.Scope
}

func (s *CreateGroupAliDingChatResponseBody) GetSourceId() *string {
	return s.SourceId
}

func (s *CreateGroupAliDingChatResponseBody) GetStatus() *string {
	return s.Status
}

func (s *CreateGroupAliDingChatResponseBody) SetChatId(v string) *CreateGroupAliDingChatResponseBody {
	s.ChatId = &v
	return s
}

func (s *CreateGroupAliDingChatResponseBody) SetCode(v string) *CreateGroupAliDingChatResponseBody {
	s.Code = &v
	return s
}

func (s *CreateGroupAliDingChatResponseBody) SetDirectoryId(v string) *CreateGroupAliDingChatResponseBody {
	s.DirectoryId = &v
	return s
}

func (s *CreateGroupAliDingChatResponseBody) SetGmtCreate(v string) *CreateGroupAliDingChatResponseBody {
	s.GmtCreate = &v
	return s
}

func (s *CreateGroupAliDingChatResponseBody) SetGroupId(v string) *CreateGroupAliDingChatResponseBody {
	s.GroupId = &v
	return s
}

func (s *CreateGroupAliDingChatResponseBody) SetMessage(v string) *CreateGroupAliDingChatResponseBody {
	s.Message = &v
	return s
}

func (s *CreateGroupAliDingChatResponseBody) SetName(v string) *CreateGroupAliDingChatResponseBody {
	s.Name = &v
	return s
}

func (s *CreateGroupAliDingChatResponseBody) SetRequestId(v string) *CreateGroupAliDingChatResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateGroupAliDingChatResponseBody) SetScope(v string) *CreateGroupAliDingChatResponseBody {
	s.Scope = &v
	return s
}

func (s *CreateGroupAliDingChatResponseBody) SetSourceId(v string) *CreateGroupAliDingChatResponseBody {
	s.SourceId = &v
	return s
}

func (s *CreateGroupAliDingChatResponseBody) SetStatus(v string) *CreateGroupAliDingChatResponseBody {
	s.Status = &v
	return s
}

func (s *CreateGroupAliDingChatResponseBody) Validate() error {
	return dara.Validate(s)
}
