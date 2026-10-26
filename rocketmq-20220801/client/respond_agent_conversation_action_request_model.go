// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRespondAgentConversationActionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAnswers(v map[string]*string) *RespondAgentConversationActionRequest
	GetAnswers() map[string]*string
	SetDecision(v string) *RespondAgentConversationActionRequest
	GetDecision() *string
	SetNote(v string) *RespondAgentConversationActionRequest
	GetNote() *string
}

type RespondAgentConversationActionRequest struct {
	// example:
	//
	// {"请选择实例地域":"华东 1（杭州）"}
	Answers map[string]*string `json:"answers,omitempty" xml:"answers,omitempty"`
	// example:
	//
	// allow
	Decision *string `json:"decision,omitempty" xml:"decision,omitempty"`
	// example:
	//
	// 允许执行只读查询
	Note *string `json:"note,omitempty" xml:"note,omitempty"`
}

func (s RespondAgentConversationActionRequest) String() string {
	return dara.Prettify(s)
}

func (s RespondAgentConversationActionRequest) GoString() string {
	return s.String()
}

func (s *RespondAgentConversationActionRequest) GetAnswers() map[string]*string {
	return s.Answers
}

func (s *RespondAgentConversationActionRequest) GetDecision() *string {
	return s.Decision
}

func (s *RespondAgentConversationActionRequest) GetNote() *string {
	return s.Note
}

func (s *RespondAgentConversationActionRequest) SetAnswers(v map[string]*string) *RespondAgentConversationActionRequest {
	s.Answers = v
	return s
}

func (s *RespondAgentConversationActionRequest) SetDecision(v string) *RespondAgentConversationActionRequest {
	s.Decision = &v
	return s
}

func (s *RespondAgentConversationActionRequest) SetNote(v string) *RespondAgentConversationActionRequest {
	s.Note = &v
	return s
}

func (s *RespondAgentConversationActionRequest) Validate() error {
	return dara.Validate(s)
}
