// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateThreadRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAssistantId(v string) *CreateThreadRequest
	GetAssistantId() *string
	SetExtLoginUser(v *CreateThreadRequestExtLoginUser) *CreateThreadRequest
	GetExtLoginUser() *CreateThreadRequestExtLoginUser
	SetOriginalAssistantId(v string) *CreateThreadRequest
	GetOriginalAssistantId() *string
	SetSourceIdOfOriginalAssistantId(v string) *CreateThreadRequest
	GetSourceIdOfOriginalAssistantId() *string
	SetSourceTypeOfOriginalAssistantId(v int32) *CreateThreadRequest
	GetSourceTypeOfOriginalAssistantId() *int32
}

type CreateThreadRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// assistantId1
	AssistantId  *string                          `json:"assistantId,omitempty" xml:"assistantId,omitempty"`
	ExtLoginUser *CreateThreadRequestExtLoginUser `json:"extLoginUser,omitempty" xml:"extLoginUser,omitempty" type:"Struct"`
	// example:
	//
	// assistantId
	OriginalAssistantId *string `json:"originalAssistantId,omitempty" xml:"originalAssistantId,omitempty"`
	// example:
	//
	// agentKey1
	SourceIdOfOriginalAssistantId   *string `json:"sourceIdOfOriginalAssistantId,omitempty" xml:"sourceIdOfOriginalAssistantId,omitempty"`
	SourceTypeOfOriginalAssistantId *int32  `json:"sourceTypeOfOriginalAssistantId,omitempty" xml:"sourceTypeOfOriginalAssistantId,omitempty"`
}

func (s CreateThreadRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateThreadRequest) GoString() string {
	return s.String()
}

func (s *CreateThreadRequest) GetAssistantId() *string {
	return s.AssistantId
}

func (s *CreateThreadRequest) GetExtLoginUser() *CreateThreadRequestExtLoginUser {
	return s.ExtLoginUser
}

func (s *CreateThreadRequest) GetOriginalAssistantId() *string {
	return s.OriginalAssistantId
}

func (s *CreateThreadRequest) GetSourceIdOfOriginalAssistantId() *string {
	return s.SourceIdOfOriginalAssistantId
}

func (s *CreateThreadRequest) GetSourceTypeOfOriginalAssistantId() *int32 {
	return s.SourceTypeOfOriginalAssistantId
}

func (s *CreateThreadRequest) SetAssistantId(v string) *CreateThreadRequest {
	s.AssistantId = &v
	return s
}

func (s *CreateThreadRequest) SetExtLoginUser(v *CreateThreadRequestExtLoginUser) *CreateThreadRequest {
	s.ExtLoginUser = v
	return s
}

func (s *CreateThreadRequest) SetOriginalAssistantId(v string) *CreateThreadRequest {
	s.OriginalAssistantId = &v
	return s
}

func (s *CreateThreadRequest) SetSourceIdOfOriginalAssistantId(v string) *CreateThreadRequest {
	s.SourceIdOfOriginalAssistantId = &v
	return s
}

func (s *CreateThreadRequest) SetSourceTypeOfOriginalAssistantId(v int32) *CreateThreadRequest {
	s.SourceTypeOfOriginalAssistantId = &v
	return s
}

func (s *CreateThreadRequest) Validate() error {
	if s.ExtLoginUser != nil {
		if err := s.ExtLoginUser.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateThreadRequestExtLoginUser struct {
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

func (s CreateThreadRequestExtLoginUser) String() string {
	return dara.Prettify(s)
}

func (s CreateThreadRequestExtLoginUser) GoString() string {
	return s.String()
}

func (s *CreateThreadRequestExtLoginUser) GetExtLoginUserDomain() *string {
	return s.ExtLoginUserDomain
}

func (s *CreateThreadRequestExtLoginUser) GetExtLoginUserId() *string {
	return s.ExtLoginUserId
}

func (s *CreateThreadRequestExtLoginUser) GetExtLoginUserName() *string {
	return s.ExtLoginUserName
}

func (s *CreateThreadRequestExtLoginUser) SetExtLoginUserDomain(v string) *CreateThreadRequestExtLoginUser {
	s.ExtLoginUserDomain = &v
	return s
}

func (s *CreateThreadRequestExtLoginUser) SetExtLoginUserId(v string) *CreateThreadRequestExtLoginUser {
	s.ExtLoginUserId = &v
	return s
}

func (s *CreateThreadRequestExtLoginUser) SetExtLoginUserName(v string) *CreateThreadRequestExtLoginUser {
	s.ExtLoginUserName = &v
	return s
}

func (s *CreateThreadRequestExtLoginUser) Validate() error {
	return dara.Validate(s)
}
