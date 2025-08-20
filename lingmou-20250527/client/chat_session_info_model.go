// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChatSessionInfo interface {
	dara.Model
	String() string
	GoString() string
	SetCreatedAt(v int64) *ChatSessionInfo
	GetCreatedAt() *int64
	SetMainAccountId(v int64) *ChatSessionInfo
	GetMainAccountId() *int64
	SetSessionId(v string) *ChatSessionInfo
	GetSessionId() *string
}

type ChatSessionInfo struct {
	CreatedAt     *int64  `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	MainAccountId *int64  `json:"mainAccountId,omitempty" xml:"mainAccountId,omitempty"`
	SessionId     *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
}

func (s ChatSessionInfo) String() string {
	return dara.Prettify(s)
}

func (s ChatSessionInfo) GoString() string {
	return s.String()
}

func (s *ChatSessionInfo) GetCreatedAt() *int64 {
	return s.CreatedAt
}

func (s *ChatSessionInfo) GetMainAccountId() *int64 {
	return s.MainAccountId
}

func (s *ChatSessionInfo) GetSessionId() *string {
	return s.SessionId
}

func (s *ChatSessionInfo) SetCreatedAt(v int64) *ChatSessionInfo {
	s.CreatedAt = &v
	return s
}

func (s *ChatSessionInfo) SetMainAccountId(v int64) *ChatSessionInfo {
	s.MainAccountId = &v
	return s
}

func (s *ChatSessionInfo) SetSessionId(v string) *ChatSessionInfo {
	s.SessionId = &v
	return s
}

func (s *ChatSessionInfo) Validate() error {
	return dara.Validate(s)
}
