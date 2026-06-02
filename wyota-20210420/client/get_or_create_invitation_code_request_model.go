// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOrCreateInvitationCodeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetExpireDays(v int32) *GetOrCreateInvitationCodeRequest
	GetExpireDays() *int32
	SetExpireMinutes(v int32) *GetOrCreateInvitationCodeRequest
	GetExpireMinutes() *int32
	SetTerminalGroupId(v string) *GetOrCreateInvitationCodeRequest
	GetTerminalGroupId() *string
	SetType(v int32) *GetOrCreateInvitationCodeRequest
	GetType() *int32
}

type GetOrCreateInvitationCodeRequest struct {
	// example:
	//
	// 1
	ExpireDays *int32 `json:"ExpireDays,omitempty" xml:"ExpireDays,omitempty"`
	// example:
	//
	// 10
	ExpireMinutes *int32 `json:"ExpireMinutes,omitempty" xml:"ExpireMinutes,omitempty"`
	// example:
	//
	// tg-XXX
	TerminalGroupId *string `json:"TerminalGroupId,omitempty" xml:"TerminalGroupId,omitempty"`
	// example:
	//
	// cron
	Type *int32 `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetOrCreateInvitationCodeRequest) String() string {
	return dara.Prettify(s)
}

func (s GetOrCreateInvitationCodeRequest) GoString() string {
	return s.String()
}

func (s *GetOrCreateInvitationCodeRequest) GetExpireDays() *int32 {
	return s.ExpireDays
}

func (s *GetOrCreateInvitationCodeRequest) GetExpireMinutes() *int32 {
	return s.ExpireMinutes
}

func (s *GetOrCreateInvitationCodeRequest) GetTerminalGroupId() *string {
	return s.TerminalGroupId
}

func (s *GetOrCreateInvitationCodeRequest) GetType() *int32 {
	return s.Type
}

func (s *GetOrCreateInvitationCodeRequest) SetExpireDays(v int32) *GetOrCreateInvitationCodeRequest {
	s.ExpireDays = &v
	return s
}

func (s *GetOrCreateInvitationCodeRequest) SetExpireMinutes(v int32) *GetOrCreateInvitationCodeRequest {
	s.ExpireMinutes = &v
	return s
}

func (s *GetOrCreateInvitationCodeRequest) SetTerminalGroupId(v string) *GetOrCreateInvitationCodeRequest {
	s.TerminalGroupId = &v
	return s
}

func (s *GetOrCreateInvitationCodeRequest) SetType(v int32) *GetOrCreateInvitationCodeRequest {
	s.Type = &v
	return s
}

func (s *GetOrCreateInvitationCodeRequest) Validate() error {
	return dara.Validate(s)
}
