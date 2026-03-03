// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConversation interface {
	dara.Model
	String() string
	GoString() string
	SetCreatedAt(v int64) *Conversation
	GetCreatedAt() *int64
	SetId(v string) *Conversation
	GetId() *string
	SetTitle(v string) *Conversation
	GetTitle() *string
	SetUpdatedAt(v int64) *Conversation
	GetUpdatedAt() *int64
}

type Conversation struct {
	CreatedAt *int64  `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	Id        *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Title     *string `json:"Title,omitempty" xml:"Title,omitempty"`
	UpdatedAt *int64  `json:"UpdatedAt,omitempty" xml:"UpdatedAt,omitempty"`
}

func (s Conversation) String() string {
	return dara.Prettify(s)
}

func (s Conversation) GoString() string {
	return s.String()
}

func (s *Conversation) GetCreatedAt() *int64 {
	return s.CreatedAt
}

func (s *Conversation) GetId() *string {
	return s.Id
}

func (s *Conversation) GetTitle() *string {
	return s.Title
}

func (s *Conversation) GetUpdatedAt() *int64 {
	return s.UpdatedAt
}

func (s *Conversation) SetCreatedAt(v int64) *Conversation {
	s.CreatedAt = &v
	return s
}

func (s *Conversation) SetId(v string) *Conversation {
	s.Id = &v
	return s
}

func (s *Conversation) SetTitle(v string) *Conversation {
	s.Title = &v
	return s
}

func (s *Conversation) SetUpdatedAt(v int64) *Conversation {
	s.UpdatedAt = &v
	return s
}

func (s *Conversation) Validate() error {
	return dara.Validate(s)
}
