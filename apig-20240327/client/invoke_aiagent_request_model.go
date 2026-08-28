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
	// This parameter is required.
	//
	// example:
	//
	// slsLogDiagnose
	AgentName *string                        `json:"agentName,omitempty" xml:"agentName,omitempty"`
	BizParams map[string]*string             `json:"bizParams,omitempty" xml:"bizParams,omitempty"`
	History   []*InvokeAIAgentRequestHistory `json:"history,omitempty" xml:"history,omitempty" type:"Repeated"`
	// example:
	//
	// zh / en
	OutputLanguage *string `json:"outputLanguage,omitempty" xml:"outputLanguage,omitempty"`
	// example:
	//
	// 帮我诊断这个错误日志
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
	// example:
	//
	// 上一轮问题
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
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
