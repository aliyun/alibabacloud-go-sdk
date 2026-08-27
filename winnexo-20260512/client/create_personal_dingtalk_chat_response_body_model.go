// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePersonalDingtalkChatResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetChatId(v string) *CreatePersonalDingtalkChatResponseBody
	GetChatId() *string
	SetCode(v string) *CreatePersonalDingtalkChatResponseBody
	GetCode() *string
	SetDirectoryId(v string) *CreatePersonalDingtalkChatResponseBody
	GetDirectoryId() *string
	SetGmtCreate(v string) *CreatePersonalDingtalkChatResponseBody
	GetGmtCreate() *string
	SetMessage(v string) *CreatePersonalDingtalkChatResponseBody
	GetMessage() *string
	SetName(v string) *CreatePersonalDingtalkChatResponseBody
	GetName() *string
	SetRequestId(v string) *CreatePersonalDingtalkChatResponseBody
	GetRequestId() *string
	SetScope(v string) *CreatePersonalDingtalkChatResponseBody
	GetScope() *string
	SetSourceId(v string) *CreatePersonalDingtalkChatResponseBody
	GetSourceId() *string
	SetStatus(v string) *CreatePersonalDingtalkChatResponseBody
	GetStatus() *string
}

type CreatePersonalDingtalkChatResponseBody struct {
	// The DingTalk group chat session ID.
	//
	// example:
	//
	// cidxxxxxxxx
	ChatId *string `json:"chatId,omitempty" xml:"chatId,omitempty"`
	// The error code.
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
	// The response message.
	//
	// example:
	//
	// ok
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The name of the AI assistant.
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
	// PERSONAL
	Scope *string `json:"scope,omitempty" xml:"scope,omitempty"`
	// The source ID.
	//
	// example:
	//
	// exampleSourceId
	SourceId *string `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
	// The task running status.
	//
	// example:
	//
	// PENDING
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s CreatePersonalDingtalkChatResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreatePersonalDingtalkChatResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePersonalDingtalkChatResponseBody) GetChatId() *string {
	return s.ChatId
}

func (s *CreatePersonalDingtalkChatResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreatePersonalDingtalkChatResponseBody) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *CreatePersonalDingtalkChatResponseBody) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *CreatePersonalDingtalkChatResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreatePersonalDingtalkChatResponseBody) GetName() *string {
	return s.Name
}

func (s *CreatePersonalDingtalkChatResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreatePersonalDingtalkChatResponseBody) GetScope() *string {
	return s.Scope
}

func (s *CreatePersonalDingtalkChatResponseBody) GetSourceId() *string {
	return s.SourceId
}

func (s *CreatePersonalDingtalkChatResponseBody) GetStatus() *string {
	return s.Status
}

func (s *CreatePersonalDingtalkChatResponseBody) SetChatId(v string) *CreatePersonalDingtalkChatResponseBody {
	s.ChatId = &v
	return s
}

func (s *CreatePersonalDingtalkChatResponseBody) SetCode(v string) *CreatePersonalDingtalkChatResponseBody {
	s.Code = &v
	return s
}

func (s *CreatePersonalDingtalkChatResponseBody) SetDirectoryId(v string) *CreatePersonalDingtalkChatResponseBody {
	s.DirectoryId = &v
	return s
}

func (s *CreatePersonalDingtalkChatResponseBody) SetGmtCreate(v string) *CreatePersonalDingtalkChatResponseBody {
	s.GmtCreate = &v
	return s
}

func (s *CreatePersonalDingtalkChatResponseBody) SetMessage(v string) *CreatePersonalDingtalkChatResponseBody {
	s.Message = &v
	return s
}

func (s *CreatePersonalDingtalkChatResponseBody) SetName(v string) *CreatePersonalDingtalkChatResponseBody {
	s.Name = &v
	return s
}

func (s *CreatePersonalDingtalkChatResponseBody) SetRequestId(v string) *CreatePersonalDingtalkChatResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreatePersonalDingtalkChatResponseBody) SetScope(v string) *CreatePersonalDingtalkChatResponseBody {
	s.Scope = &v
	return s
}

func (s *CreatePersonalDingtalkChatResponseBody) SetSourceId(v string) *CreatePersonalDingtalkChatResponseBody {
	s.SourceId = &v
	return s
}

func (s *CreatePersonalDingtalkChatResponseBody) SetStatus(v string) *CreatePersonalDingtalkChatResponseBody {
	s.Status = &v
	return s
}

func (s *CreatePersonalDingtalkChatResponseBody) Validate() error {
	return dara.Validate(s)
}
