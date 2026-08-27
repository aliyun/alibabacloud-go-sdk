// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePersonalAliDingChatResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetChatId(v string) *CreatePersonalAliDingChatResponseBody
	GetChatId() *string
	SetCode(v string) *CreatePersonalAliDingChatResponseBody
	GetCode() *string
	SetDirectoryId(v string) *CreatePersonalAliDingChatResponseBody
	GetDirectoryId() *string
	SetGmtCreate(v string) *CreatePersonalAliDingChatResponseBody
	GetGmtCreate() *string
	SetMessage(v string) *CreatePersonalAliDingChatResponseBody
	GetMessage() *string
	SetName(v string) *CreatePersonalAliDingChatResponseBody
	GetName() *string
	SetRequestId(v string) *CreatePersonalAliDingChatResponseBody
	GetRequestId() *string
	SetScope(v string) *CreatePersonalAliDingChatResponseBody
	GetScope() *string
	SetSourceId(v string) *CreatePersonalAliDingChatResponseBody
	GetSourceId() *string
	SetStatus(v string) *CreatePersonalAliDingChatResponseBody
	GetStatus() *string
}

type CreatePersonalAliDingChatResponseBody struct {
	// The DingTalk group chat session ID.
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
	// The folder ID.
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
	// The skill name.
	//
	// example:
	//
	// CustomerProjectGroup
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The request trace ID.
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
	// The original project ID.
	//
	// example:
	//
	// exampleSourceId
	SourceId *string `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
	// The call status. Valid values:
	//
	// - **PENDING**: Waiting for acknowledgment.
	//
	// - **SUCCESS**: Succeeded.
	//
	// - **FAILED**: Failed.
	//
	// - **TIMEOUT**: Timed out.
	//
	// example:
	//
	// PENDING
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s CreatePersonalAliDingChatResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreatePersonalAliDingChatResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePersonalAliDingChatResponseBody) GetChatId() *string {
	return s.ChatId
}

func (s *CreatePersonalAliDingChatResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreatePersonalAliDingChatResponseBody) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *CreatePersonalAliDingChatResponseBody) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *CreatePersonalAliDingChatResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreatePersonalAliDingChatResponseBody) GetName() *string {
	return s.Name
}

func (s *CreatePersonalAliDingChatResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreatePersonalAliDingChatResponseBody) GetScope() *string {
	return s.Scope
}

func (s *CreatePersonalAliDingChatResponseBody) GetSourceId() *string {
	return s.SourceId
}

func (s *CreatePersonalAliDingChatResponseBody) GetStatus() *string {
	return s.Status
}

func (s *CreatePersonalAliDingChatResponseBody) SetChatId(v string) *CreatePersonalAliDingChatResponseBody {
	s.ChatId = &v
	return s
}

func (s *CreatePersonalAliDingChatResponseBody) SetCode(v string) *CreatePersonalAliDingChatResponseBody {
	s.Code = &v
	return s
}

func (s *CreatePersonalAliDingChatResponseBody) SetDirectoryId(v string) *CreatePersonalAliDingChatResponseBody {
	s.DirectoryId = &v
	return s
}

func (s *CreatePersonalAliDingChatResponseBody) SetGmtCreate(v string) *CreatePersonalAliDingChatResponseBody {
	s.GmtCreate = &v
	return s
}

func (s *CreatePersonalAliDingChatResponseBody) SetMessage(v string) *CreatePersonalAliDingChatResponseBody {
	s.Message = &v
	return s
}

func (s *CreatePersonalAliDingChatResponseBody) SetName(v string) *CreatePersonalAliDingChatResponseBody {
	s.Name = &v
	return s
}

func (s *CreatePersonalAliDingChatResponseBody) SetRequestId(v string) *CreatePersonalAliDingChatResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreatePersonalAliDingChatResponseBody) SetScope(v string) *CreatePersonalAliDingChatResponseBody {
	s.Scope = &v
	return s
}

func (s *CreatePersonalAliDingChatResponseBody) SetSourceId(v string) *CreatePersonalAliDingChatResponseBody {
	s.SourceId = &v
	return s
}

func (s *CreatePersonalAliDingChatResponseBody) SetStatus(v string) *CreatePersonalAliDingChatResponseBody {
	s.Status = &v
	return s
}

func (s *CreatePersonalAliDingChatResponseBody) Validate() error {
	return dara.Validate(s)
}
