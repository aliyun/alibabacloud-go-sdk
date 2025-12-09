// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetChatContentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCategory(v string) *GetChatContentResponseBody
	GetCategory() *string
	SetCheckpoint(v int64) *GetChatContentResponseBody
	GetCheckpoint() *int64
	SetContent(v string) *GetChatContentResponseBody
	GetContent() *string
	SetContentType(v string) *GetChatContentResponseBody
	GetContentType() *string
	SetEventType(v string) *GetChatContentResponseBody
	GetEventType() *string
	SetLevel(v int64) *GetChatContentResponseBody
	GetLevel() *int64
}

type GetChatContentResponseBody struct {
	// example:
	//
	// PLAN
	Category *string `json:"category,omitempty" xml:"category,omitempty"`
	// example:
	//
	// 0
	Checkpoint *int64 `json:"checkpoint,omitempty" xml:"checkpoint,omitempty"`
	// example:
	//
	// Data understanding completed.
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// example:
	//
	// str
	ContentType *string `json:"content_type,omitempty" xml:"content_type,omitempty"`
	// example:
	//
	// DATA
	EventType *string `json:"event_type,omitempty" xml:"event_type,omitempty"`
	// example:
	//
	// 20
	Level *int64 `json:"level,omitempty" xml:"level,omitempty"`
}

func (s GetChatContentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetChatContentResponseBody) GoString() string {
	return s.String()
}

func (s *GetChatContentResponseBody) GetCategory() *string {
	return s.Category
}

func (s *GetChatContentResponseBody) GetCheckpoint() *int64 {
	return s.Checkpoint
}

func (s *GetChatContentResponseBody) GetContent() *string {
	return s.Content
}

func (s *GetChatContentResponseBody) GetContentType() *string {
	return s.ContentType
}

func (s *GetChatContentResponseBody) GetEventType() *string {
	return s.EventType
}

func (s *GetChatContentResponseBody) GetLevel() *int64 {
	return s.Level
}

func (s *GetChatContentResponseBody) SetCategory(v string) *GetChatContentResponseBody {
	s.Category = &v
	return s
}

func (s *GetChatContentResponseBody) SetCheckpoint(v int64) *GetChatContentResponseBody {
	s.Checkpoint = &v
	return s
}

func (s *GetChatContentResponseBody) SetContent(v string) *GetChatContentResponseBody {
	s.Content = &v
	return s
}

func (s *GetChatContentResponseBody) SetContentType(v string) *GetChatContentResponseBody {
	s.ContentType = &v
	return s
}

func (s *GetChatContentResponseBody) SetEventType(v string) *GetChatContentResponseBody {
	s.EventType = &v
	return s
}

func (s *GetChatContentResponseBody) SetLevel(v int64) *GetChatContentResponseBody {
	s.Level = &v
	return s
}

func (s *GetChatContentResponseBody) Validate() error {
	return dara.Validate(s)
}
