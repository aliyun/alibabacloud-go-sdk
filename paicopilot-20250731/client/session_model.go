// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSession interface {
	dara.Model
	String() string
	GoString() string
	SetGmtCreateTime(v string) *Session
	GetGmtCreateTime() *string
	SetGmtModified(v string) *Session
	GetGmtModified() *string
	SetOwnerId(v string) *Session
	GetOwnerId() *string
	SetSessionId(v string) *Session
	GetSessionId() *string
	SetSessionTitle(v string) *Session
	GetSessionTitle() *string
	SetUserId(v string) *Session
	GetUserId() *string
}

type Session struct {
	// Use the UTC time format: yyyy-MM-ddTHH:mm:ss.SSSZ
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// Use the UTC time format: yyyy-MM-ddTHH:mm:ss.SSSZ
	GmtModified  *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	OwnerId      *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SessionId    *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	SessionTitle *string `json:"SessionTitle,omitempty" xml:"SessionTitle,omitempty"`
	UserId       *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s Session) String() string {
	return dara.Prettify(s)
}

func (s Session) GoString() string {
	return s.String()
}

func (s *Session) GetGmtCreateTime() *string {
	return s.GmtCreateTime
}

func (s *Session) GetGmtModified() *string {
	return s.GmtModified
}

func (s *Session) GetOwnerId() *string {
	return s.OwnerId
}

func (s *Session) GetSessionId() *string {
	return s.SessionId
}

func (s *Session) GetSessionTitle() *string {
	return s.SessionTitle
}

func (s *Session) GetUserId() *string {
	return s.UserId
}

func (s *Session) SetGmtCreateTime(v string) *Session {
	s.GmtCreateTime = &v
	return s
}

func (s *Session) SetGmtModified(v string) *Session {
	s.GmtModified = &v
	return s
}

func (s *Session) SetOwnerId(v string) *Session {
	s.OwnerId = &v
	return s
}

func (s *Session) SetSessionId(v string) *Session {
	s.SessionId = &v
	return s
}

func (s *Session) SetSessionTitle(v string) *Session {
	s.SessionTitle = &v
	return s
}

func (s *Session) SetUserId(v string) *Session {
	s.UserId = &v
	return s
}

func (s *Session) Validate() error {
	return dara.Validate(s)
}
