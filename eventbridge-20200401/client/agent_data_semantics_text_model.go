// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAgentDataSemanticsText interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v string) *AgentDataSemanticsText
	GetContent() *string
}

type AgentDataSemanticsText struct {
	// The content of the Markdown-formatted text knowledge.
	//
	// This parameter is required.
	//
	// example:
	//
	// ## Data description
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
}

func (s AgentDataSemanticsText) String() string {
	return dara.Prettify(s)
}

func (s AgentDataSemanticsText) GoString() string {
	return s.String()
}

func (s *AgentDataSemanticsText) GetContent() *string {
	return s.Content
}

func (s *AgentDataSemanticsText) SetContent(v string) *AgentDataSemanticsText {
	s.Content = &v
	return s
}

func (s *AgentDataSemanticsText) Validate() error {
	return dara.Validate(s)
}
