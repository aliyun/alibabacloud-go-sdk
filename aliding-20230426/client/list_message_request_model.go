// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMessageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAssistantId(v string) *ListMessageRequest
	GetAssistantId() *string
	SetExtLoginUser(v *ListMessageRequestExtLoginUser) *ListMessageRequest
	GetExtLoginUser() *ListMessageRequestExtLoginUser
	SetLimit(v int32) *ListMessageRequest
	GetLimit() *int32
	SetOrder(v string) *ListMessageRequest
	GetOrder() *string
	SetOriginalAssistantId(v string) *ListMessageRequest
	GetOriginalAssistantId() *string
	SetRunId(v string) *ListMessageRequest
	GetRunId() *string
	SetSourceIdOfOriginalAssistantId(v string) *ListMessageRequest
	GetSourceIdOfOriginalAssistantId() *string
	SetSourceTypeOfOriginalAssistantId(v string) *ListMessageRequest
	GetSourceTypeOfOriginalAssistantId() *string
	SetThreadId(v string) *ListMessageRequest
	GetThreadId() *string
}

type ListMessageRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// assistantId1
	AssistantId  *string                         `json:"assistantId,omitempty" xml:"assistantId,omitempty"`
	ExtLoginUser *ListMessageRequestExtLoginUser `json:"extLoginUser,omitempty" xml:"extLoginUser,omitempty" type:"Struct"`
	// example:
	//
	// 20
	Limit *int32 `json:"limit,omitempty" xml:"limit,omitempty"`
	// example:
	//
	// desc
	Order *string `json:"order,omitempty" xml:"order,omitempty"`
	// example:
	//
	// assistantId2
	OriginalAssistantId *string `json:"originalAssistantId,omitempty" xml:"originalAssistantId,omitempty"`
	// example:
	//
	// runId123
	RunId *string `json:"runId,omitempty" xml:"runId,omitempty"`
	// example:
	//
	// agentKey1
	SourceIdOfOriginalAssistantId *string `json:"sourceIdOfOriginalAssistantId,omitempty" xml:"sourceIdOfOriginalAssistantId,omitempty"`
	// example:
	//
	// 1
	SourceTypeOfOriginalAssistantId *string `json:"sourceTypeOfOriginalAssistantId,omitempty" xml:"sourceTypeOfOriginalAssistantId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// threadId123
	ThreadId *string `json:"threadId,omitempty" xml:"threadId,omitempty"`
}

func (s ListMessageRequest) String() string {
	return dara.Prettify(s)
}

func (s ListMessageRequest) GoString() string {
	return s.String()
}

func (s *ListMessageRequest) GetAssistantId() *string {
	return s.AssistantId
}

func (s *ListMessageRequest) GetExtLoginUser() *ListMessageRequestExtLoginUser {
	return s.ExtLoginUser
}

func (s *ListMessageRequest) GetLimit() *int32 {
	return s.Limit
}

func (s *ListMessageRequest) GetOrder() *string {
	return s.Order
}

func (s *ListMessageRequest) GetOriginalAssistantId() *string {
	return s.OriginalAssistantId
}

func (s *ListMessageRequest) GetRunId() *string {
	return s.RunId
}

func (s *ListMessageRequest) GetSourceIdOfOriginalAssistantId() *string {
	return s.SourceIdOfOriginalAssistantId
}

func (s *ListMessageRequest) GetSourceTypeOfOriginalAssistantId() *string {
	return s.SourceTypeOfOriginalAssistantId
}

func (s *ListMessageRequest) GetThreadId() *string {
	return s.ThreadId
}

func (s *ListMessageRequest) SetAssistantId(v string) *ListMessageRequest {
	s.AssistantId = &v
	return s
}

func (s *ListMessageRequest) SetExtLoginUser(v *ListMessageRequestExtLoginUser) *ListMessageRequest {
	s.ExtLoginUser = v
	return s
}

func (s *ListMessageRequest) SetLimit(v int32) *ListMessageRequest {
	s.Limit = &v
	return s
}

func (s *ListMessageRequest) SetOrder(v string) *ListMessageRequest {
	s.Order = &v
	return s
}

func (s *ListMessageRequest) SetOriginalAssistantId(v string) *ListMessageRequest {
	s.OriginalAssistantId = &v
	return s
}

func (s *ListMessageRequest) SetRunId(v string) *ListMessageRequest {
	s.RunId = &v
	return s
}

func (s *ListMessageRequest) SetSourceIdOfOriginalAssistantId(v string) *ListMessageRequest {
	s.SourceIdOfOriginalAssistantId = &v
	return s
}

func (s *ListMessageRequest) SetSourceTypeOfOriginalAssistantId(v string) *ListMessageRequest {
	s.SourceTypeOfOriginalAssistantId = &v
	return s
}

func (s *ListMessageRequest) SetThreadId(v string) *ListMessageRequest {
	s.ThreadId = &v
	return s
}

func (s *ListMessageRequest) Validate() error {
	if s.ExtLoginUser != nil {
		if err := s.ExtLoginUser.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListMessageRequestExtLoginUser struct {
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

func (s ListMessageRequestExtLoginUser) String() string {
	return dara.Prettify(s)
}

func (s ListMessageRequestExtLoginUser) GoString() string {
	return s.String()
}

func (s *ListMessageRequestExtLoginUser) GetExtLoginUserDomain() *string {
	return s.ExtLoginUserDomain
}

func (s *ListMessageRequestExtLoginUser) GetExtLoginUserId() *string {
	return s.ExtLoginUserId
}

func (s *ListMessageRequestExtLoginUser) GetExtLoginUserName() *string {
	return s.ExtLoginUserName
}

func (s *ListMessageRequestExtLoginUser) SetExtLoginUserDomain(v string) *ListMessageRequestExtLoginUser {
	s.ExtLoginUserDomain = &v
	return s
}

func (s *ListMessageRequestExtLoginUser) SetExtLoginUserId(v string) *ListMessageRequestExtLoginUser {
	s.ExtLoginUserId = &v
	return s
}

func (s *ListMessageRequestExtLoginUser) SetExtLoginUserName(v string) *ListMessageRequestExtLoginUser {
	s.ExtLoginUserName = &v
	return s
}

func (s *ListMessageRequestExtLoginUser) Validate() error {
	return dara.Validate(s)
}
