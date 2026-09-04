// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateGroupDingtalkChatResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetChatId(v string) *CreateGroupDingtalkChatResponseBody
	GetChatId() *string
	SetCode(v string) *CreateGroupDingtalkChatResponseBody
	GetCode() *string
	SetDirectoryId(v string) *CreateGroupDingtalkChatResponseBody
	GetDirectoryId() *string
	SetGmtCreate(v string) *CreateGroupDingtalkChatResponseBody
	GetGmtCreate() *string
	SetGroupId(v string) *CreateGroupDingtalkChatResponseBody
	GetGroupId() *string
	SetMessage(v string) *CreateGroupDingtalkChatResponseBody
	GetMessage() *string
	SetName(v string) *CreateGroupDingtalkChatResponseBody
	GetName() *string
	SetRequestId(v string) *CreateGroupDingtalkChatResponseBody
	GetRequestId() *string
	SetScope(v string) *CreateGroupDingtalkChatResponseBody
	GetScope() *string
	SetSourceId(v string) *CreateGroupDingtalkChatResponseBody
	GetSourceId() *string
	SetStatus(v string) *CreateGroupDingtalkChatResponseBody
	GetStatus() *string
}

type CreateGroupDingtalkChatResponseBody struct {
	// The session ID, typically used for JSSDK.
	//
	// example:
	//
	// cidxxxxxxxx
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
	// The description of the status code.
	//
	// example:
	//
	// The current zone list is illegal.
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The image name.
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
	// The data source ID (unique within the tenant).
	//
	// example:
	//
	// exampleSourceId
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
	// PENDING
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s CreateGroupDingtalkChatResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateGroupDingtalkChatResponseBody) GoString() string {
	return s.String()
}

func (s *CreateGroupDingtalkChatResponseBody) GetChatId() *string {
	return s.ChatId
}

func (s *CreateGroupDingtalkChatResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateGroupDingtalkChatResponseBody) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *CreateGroupDingtalkChatResponseBody) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *CreateGroupDingtalkChatResponseBody) GetGroupId() *string {
	return s.GroupId
}

func (s *CreateGroupDingtalkChatResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateGroupDingtalkChatResponseBody) GetName() *string {
	return s.Name
}

func (s *CreateGroupDingtalkChatResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateGroupDingtalkChatResponseBody) GetScope() *string {
	return s.Scope
}

func (s *CreateGroupDingtalkChatResponseBody) GetSourceId() *string {
	return s.SourceId
}

func (s *CreateGroupDingtalkChatResponseBody) GetStatus() *string {
	return s.Status
}

func (s *CreateGroupDingtalkChatResponseBody) SetChatId(v string) *CreateGroupDingtalkChatResponseBody {
	s.ChatId = &v
	return s
}

func (s *CreateGroupDingtalkChatResponseBody) SetCode(v string) *CreateGroupDingtalkChatResponseBody {
	s.Code = &v
	return s
}

func (s *CreateGroupDingtalkChatResponseBody) SetDirectoryId(v string) *CreateGroupDingtalkChatResponseBody {
	s.DirectoryId = &v
	return s
}

func (s *CreateGroupDingtalkChatResponseBody) SetGmtCreate(v string) *CreateGroupDingtalkChatResponseBody {
	s.GmtCreate = &v
	return s
}

func (s *CreateGroupDingtalkChatResponseBody) SetGroupId(v string) *CreateGroupDingtalkChatResponseBody {
	s.GroupId = &v
	return s
}

func (s *CreateGroupDingtalkChatResponseBody) SetMessage(v string) *CreateGroupDingtalkChatResponseBody {
	s.Message = &v
	return s
}

func (s *CreateGroupDingtalkChatResponseBody) SetName(v string) *CreateGroupDingtalkChatResponseBody {
	s.Name = &v
	return s
}

func (s *CreateGroupDingtalkChatResponseBody) SetRequestId(v string) *CreateGroupDingtalkChatResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateGroupDingtalkChatResponseBody) SetScope(v string) *CreateGroupDingtalkChatResponseBody {
	s.Scope = &v
	return s
}

func (s *CreateGroupDingtalkChatResponseBody) SetSourceId(v string) *CreateGroupDingtalkChatResponseBody {
	s.SourceId = &v
	return s
}

func (s *CreateGroupDingtalkChatResponseBody) SetStatus(v string) *CreateGroupDingtalkChatResponseBody {
	s.Status = &v
	return s
}

func (s *CreateGroupDingtalkChatResponseBody) Validate() error {
	return dara.Validate(s)
}
