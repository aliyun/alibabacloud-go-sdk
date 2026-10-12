// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInvokeAIAgentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentName(v string) *InvokeAIAgentRequest
	GetAgentName() *string
	SetBizParams(v map[string]*string) *InvokeAIAgentRequest
	GetBizParams() map[string]*string
	SetHistory(v []*InvokeAIAgentRequestHistory) *InvokeAIAgentRequest
	GetHistory() []*InvokeAIAgentRequestHistory
	SetOutputLanguage(v string) *InvokeAIAgentRequest
	GetOutputLanguage() *string
	SetPrompt(v string) *InvokeAIAgentRequest
	GetPrompt() *string
}

type InvokeAIAgentRequest struct {
	// The agent name.
	//
	// This parameter is required.
	//
	// example:
	//
	// slsLogDiagnose
	AgentName *string `json:"agentName,omitempty" xml:"agentName,omitempty"`
	// The business parameters. Each agent performs custom validation on these parameters, such as the SLS context passed for log diagnostics.
	BizParams map[string]*string `json:"bizParams,omitempty" xml:"bizParams,omitempty"`
	// The conversation history. Each element contains role (user/assistant) and content. The role and content fields in each element cannot be empty.
	History []*InvokeAIAgentRequestHistory `json:"history,omitempty" xml:"history,omitempty" type:"Repeated"`
	// The output language. When set to en, fallback error messages are displayed in English.
	//
	// example:
	//
	// zh / en
	OutputLanguage *string `json:"outputLanguage,omitempty" xml:"outputLanguage,omitempty"`
	// The user input for the current turn. If left empty, a default prompt is used based on the scenario, such as "Start".
	//
	// example:
	//
	// Help me diagnose this error log
	Prompt *string `json:"prompt,omitempty" xml:"prompt,omitempty"`
}

func (s InvokeAIAgentRequest) String() string {
	return dara.Prettify(s)
}

func (s InvokeAIAgentRequest) GoString() string {
	return s.String()
}

func (s *InvokeAIAgentRequest) GetAgentName() *string {
	return s.AgentName
}

func (s *InvokeAIAgentRequest) GetBizParams() map[string]*string {
	return s.BizParams
}

func (s *InvokeAIAgentRequest) GetHistory() []*InvokeAIAgentRequestHistory {
	return s.History
}

func (s *InvokeAIAgentRequest) GetOutputLanguage() *string {
	return s.OutputLanguage
}

func (s *InvokeAIAgentRequest) GetPrompt() *string {
	return s.Prompt
}

func (s *InvokeAIAgentRequest) SetAgentName(v string) *InvokeAIAgentRequest {
	s.AgentName = &v
	return s
}

func (s *InvokeAIAgentRequest) SetBizParams(v map[string]*string) *InvokeAIAgentRequest {
	s.BizParams = v
	return s
}

func (s *InvokeAIAgentRequest) SetHistory(v []*InvokeAIAgentRequestHistory) *InvokeAIAgentRequest {
	s.History = v
	return s
}

func (s *InvokeAIAgentRequest) SetOutputLanguage(v string) *InvokeAIAgentRequest {
	s.OutputLanguage = &v
	return s
}

func (s *InvokeAIAgentRequest) SetPrompt(v string) *InvokeAIAgentRequest {
	s.Prompt = &v
	return s
}

func (s *InvokeAIAgentRequest) Validate() error {
	if s.History != nil {
		for _, item := range s.History {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type InvokeAIAgentRequestHistory struct {
	// The content of the message. The returned data is in rich text format.
	//
	// example:
	//
	// Previous question
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// The role.
	//
	// example:
	//
	// user
	Role *string `json:"role,omitempty" xml:"role,omitempty"`
}

func (s InvokeAIAgentRequestHistory) String() string {
	return dara.Prettify(s)
}

func (s InvokeAIAgentRequestHistory) GoString() string {
	return s.String()
}

func (s *InvokeAIAgentRequestHistory) GetContent() *string {
	return s.Content
}

func (s *InvokeAIAgentRequestHistory) GetRole() *string {
	return s.Role
}

func (s *InvokeAIAgentRequestHistory) SetContent(v string) *InvokeAIAgentRequestHistory {
	s.Content = &v
	return s
}

func (s *InvokeAIAgentRequestHistory) SetRole(v string) *InvokeAIAgentRequestHistory {
	s.Role = &v
	return s
}

func (s *InvokeAIAgentRequestHistory) Validate() error {
	return dara.Validate(s)
}
