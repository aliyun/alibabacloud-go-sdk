// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRunRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAllowStructViewContent(v bool) *CreateRunRequest
	GetAllowStructViewContent() *bool
	SetAssistantId(v string) *CreateRunRequest
	GetAssistantId() *string
	SetExtLoginUser(v *CreateRunRequestExtLoginUser) *CreateRunRequest
	GetExtLoginUser() *CreateRunRequestExtLoginUser
	SetOriginalAssistantId(v string) *CreateRunRequest
	GetOriginalAssistantId() *string
	SetSourceIdOfOriginalAssistantId(v string) *CreateRunRequest
	GetSourceIdOfOriginalAssistantId() *string
	SetSourceTypeOfOriginalAssistantId(v string) *CreateRunRequest
	GetSourceTypeOfOriginalAssistantId() *string
	SetStream(v bool) *CreateRunRequest
	GetStream() *bool
	SetThreadId(v string) *CreateRunRequest
	GetThreadId() *string
}

type CreateRunRequest struct {
	AllowStructViewContent *bool `json:"allowStructViewContent,omitempty" xml:"allowStructViewContent,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// assistantId1
	AssistantId  *string                       `json:"assistantId,omitempty" xml:"assistantId,omitempty"`
	ExtLoginUser *CreateRunRequestExtLoginUser `json:"extLoginUser,omitempty" xml:"extLoginUser,omitempty" type:"Struct"`
	// example:
	//
	// assistantId
	OriginalAssistantId *string `json:"originalAssistantId,omitempty" xml:"originalAssistantId,omitempty"`
	// example:
	//
	// agentKey1
	SourceIdOfOriginalAssistantId *string `json:"sourceIdOfOriginalAssistantId,omitempty" xml:"sourceIdOfOriginalAssistantId,omitempty"`
	// example:
	//
	// 1
	SourceTypeOfOriginalAssistantId *string `json:"sourceTypeOfOriginalAssistantId,omitempty" xml:"sourceTypeOfOriginalAssistantId,omitempty"`
	// example:
	//
	// false
	Stream *bool `json:"stream,omitempty" xml:"stream,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// threadId123
	ThreadId *string `json:"threadId,omitempty" xml:"threadId,omitempty"`
}

func (s CreateRunRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateRunRequest) GoString() string {
	return s.String()
}

func (s *CreateRunRequest) GetAllowStructViewContent() *bool {
	return s.AllowStructViewContent
}

func (s *CreateRunRequest) GetAssistantId() *string {
	return s.AssistantId
}

func (s *CreateRunRequest) GetExtLoginUser() *CreateRunRequestExtLoginUser {
	return s.ExtLoginUser
}

func (s *CreateRunRequest) GetOriginalAssistantId() *string {
	return s.OriginalAssistantId
}

func (s *CreateRunRequest) GetSourceIdOfOriginalAssistantId() *string {
	return s.SourceIdOfOriginalAssistantId
}

func (s *CreateRunRequest) GetSourceTypeOfOriginalAssistantId() *string {
	return s.SourceTypeOfOriginalAssistantId
}

func (s *CreateRunRequest) GetStream() *bool {
	return s.Stream
}

func (s *CreateRunRequest) GetThreadId() *string {
	return s.ThreadId
}

func (s *CreateRunRequest) SetAllowStructViewContent(v bool) *CreateRunRequest {
	s.AllowStructViewContent = &v
	return s
}

func (s *CreateRunRequest) SetAssistantId(v string) *CreateRunRequest {
	s.AssistantId = &v
	return s
}

func (s *CreateRunRequest) SetExtLoginUser(v *CreateRunRequestExtLoginUser) *CreateRunRequest {
	s.ExtLoginUser = v
	return s
}

func (s *CreateRunRequest) SetOriginalAssistantId(v string) *CreateRunRequest {
	s.OriginalAssistantId = &v
	return s
}

func (s *CreateRunRequest) SetSourceIdOfOriginalAssistantId(v string) *CreateRunRequest {
	s.SourceIdOfOriginalAssistantId = &v
	return s
}

func (s *CreateRunRequest) SetSourceTypeOfOriginalAssistantId(v string) *CreateRunRequest {
	s.SourceTypeOfOriginalAssistantId = &v
	return s
}

func (s *CreateRunRequest) SetStream(v bool) *CreateRunRequest {
	s.Stream = &v
	return s
}

func (s *CreateRunRequest) SetThreadId(v string) *CreateRunRequest {
	s.ThreadId = &v
	return s
}

func (s *CreateRunRequest) Validate() error {
	if s.ExtLoginUser != nil {
		if err := s.ExtLoginUser.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateRunRequestExtLoginUser struct {
	// example:
	//
	// mozi
	ExtLoginUserDomain *string `json:"extLoginUserDomain,omitempty" xml:"extLoginUserDomain,omitempty"`
	// example:
	//
	// outeruserId123
	ExtLoginUserId *string `json:"extLoginUserId,omitempty" xml:"extLoginUserId,omitempty"`
	// example:
	//
	// 外部游客1
	ExtLoginUserName *string `json:"extLoginUserName,omitempty" xml:"extLoginUserName,omitempty"`
}

func (s CreateRunRequestExtLoginUser) String() string {
	return dara.Prettify(s)
}

func (s CreateRunRequestExtLoginUser) GoString() string {
	return s.String()
}

func (s *CreateRunRequestExtLoginUser) GetExtLoginUserDomain() *string {
	return s.ExtLoginUserDomain
}

func (s *CreateRunRequestExtLoginUser) GetExtLoginUserId() *string {
	return s.ExtLoginUserId
}

func (s *CreateRunRequestExtLoginUser) GetExtLoginUserName() *string {
	return s.ExtLoginUserName
}

func (s *CreateRunRequestExtLoginUser) SetExtLoginUserDomain(v string) *CreateRunRequestExtLoginUser {
	s.ExtLoginUserDomain = &v
	return s
}

func (s *CreateRunRequestExtLoginUser) SetExtLoginUserId(v string) *CreateRunRequestExtLoginUser {
	s.ExtLoginUserId = &v
	return s
}

func (s *CreateRunRequestExtLoginUser) SetExtLoginUserName(v string) *CreateRunRequestExtLoginUser {
	s.ExtLoginUserName = &v
	return s
}

func (s *CreateRunRequestExtLoginUser) Validate() error {
	return dara.Validate(s)
}
