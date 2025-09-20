// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMessage interface {
	dara.Model
	String() string
	GoString() string
	SetAssistantType(v string) *Message
	GetAssistantType() *string
	SetContent(v string) *Message
	GetContent() *string
	SetCreateTime(v string) *Message
	GetCreateTime() *string
	SetDatasetName(v string) *Message
	GetDatasetName() *string
	SetLanguage(v string) *Message
	GetLanguage() *string
	SetRegenerate(v bool) *Message
	GetRegenerate() *bool
	SetReply(v string) *Message
	GetReply() *string
	SetScore(v float64) *Message
	GetScore() *float64
	SetSourceURI(v string) *Message
	GetSourceURI() *string
	SetSuggestion(v string) *Message
	GetSuggestion() *string
	SetTone(v string) *Message
	GetTone() *string
	SetTopic(v string) *Message
	GetTopic() *string
}

type Message struct {
	AssistantType *string  `json:"AssistantType,omitempty" xml:"AssistantType,omitempty"`
	Content       *string  `json:"Content,omitempty" xml:"Content,omitempty"`
	CreateTime    *string  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	DatasetName   *string  `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	Language      *string  `json:"Language,omitempty" xml:"Language,omitempty"`
	Regenerate    *bool    `json:"Regenerate,omitempty" xml:"Regenerate,omitempty"`
	Reply         *string  `json:"Reply,omitempty" xml:"Reply,omitempty"`
	Score         *float64 `json:"Score,omitempty" xml:"Score,omitempty"`
	SourceURI     *string  `json:"SourceURI,omitempty" xml:"SourceURI,omitempty"`
	Suggestion    *string  `json:"Suggestion,omitempty" xml:"Suggestion,omitempty"`
	Tone          *string  `json:"Tone,omitempty" xml:"Tone,omitempty"`
	Topic         *string  `json:"Topic,omitempty" xml:"Topic,omitempty"`
}

func (s Message) String() string {
	return dara.Prettify(s)
}

func (s Message) GoString() string {
	return s.String()
}

func (s *Message) GetAssistantType() *string {
	return s.AssistantType
}

func (s *Message) GetContent() *string {
	return s.Content
}

func (s *Message) GetCreateTime() *string {
	return s.CreateTime
}

func (s *Message) GetDatasetName() *string {
	return s.DatasetName
}

func (s *Message) GetLanguage() *string {
	return s.Language
}

func (s *Message) GetRegenerate() *bool {
	return s.Regenerate
}

func (s *Message) GetReply() *string {
	return s.Reply
}

func (s *Message) GetScore() *float64 {
	return s.Score
}

func (s *Message) GetSourceURI() *string {
	return s.SourceURI
}

func (s *Message) GetSuggestion() *string {
	return s.Suggestion
}

func (s *Message) GetTone() *string {
	return s.Tone
}

func (s *Message) GetTopic() *string {
	return s.Topic
}

func (s *Message) SetAssistantType(v string) *Message {
	s.AssistantType = &v
	return s
}

func (s *Message) SetContent(v string) *Message {
	s.Content = &v
	return s
}

func (s *Message) SetCreateTime(v string) *Message {
	s.CreateTime = &v
	return s
}

func (s *Message) SetDatasetName(v string) *Message {
	s.DatasetName = &v
	return s
}

func (s *Message) SetLanguage(v string) *Message {
	s.Language = &v
	return s
}

func (s *Message) SetRegenerate(v bool) *Message {
	s.Regenerate = &v
	return s
}

func (s *Message) SetReply(v string) *Message {
	s.Reply = &v
	return s
}

func (s *Message) SetScore(v float64) *Message {
	s.Score = &v
	return s
}

func (s *Message) SetSourceURI(v string) *Message {
	s.SourceURI = &v
	return s
}

func (s *Message) SetSuggestion(v string) *Message {
	s.Suggestion = &v
	return s
}

func (s *Message) SetTone(v string) *Message {
	s.Tone = &v
	return s
}

func (s *Message) SetTopic(v string) *Message {
	s.Topic = &v
	return s
}

func (s *Message) Validate() error {
	return dara.Validate(s)
}
