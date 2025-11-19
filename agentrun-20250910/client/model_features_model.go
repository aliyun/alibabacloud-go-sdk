// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelFeatures interface {
	dara.Model
	String() string
	GoString() string
	SetAgentThought(v bool) *ModelFeatures
	GetAgentThought() *bool
	SetMultiToolCall(v bool) *ModelFeatures
	GetMultiToolCall() *bool
	SetStreamToolCall(v bool) *ModelFeatures
	GetStreamToolCall() *bool
	SetToolCall(v bool) *ModelFeatures
	GetToolCall() *bool
	SetVision(v bool) *ModelFeatures
	GetVision() *bool
}

type ModelFeatures struct {
	AgentThought   *bool `json:"agentThought,omitempty" xml:"agentThought,omitempty"`
	MultiToolCall  *bool `json:"multiToolCall,omitempty" xml:"multiToolCall,omitempty"`
	StreamToolCall *bool `json:"streamToolCall,omitempty" xml:"streamToolCall,omitempty"`
	ToolCall       *bool `json:"toolCall,omitempty" xml:"toolCall,omitempty"`
	Vision         *bool `json:"vision,omitempty" xml:"vision,omitempty"`
}

func (s ModelFeatures) String() string {
	return dara.Prettify(s)
}

func (s ModelFeatures) GoString() string {
	return s.String()
}

func (s *ModelFeatures) GetAgentThought() *bool {
	return s.AgentThought
}

func (s *ModelFeatures) GetMultiToolCall() *bool {
	return s.MultiToolCall
}

func (s *ModelFeatures) GetStreamToolCall() *bool {
	return s.StreamToolCall
}

func (s *ModelFeatures) GetToolCall() *bool {
	return s.ToolCall
}

func (s *ModelFeatures) GetVision() *bool {
	return s.Vision
}

func (s *ModelFeatures) SetAgentThought(v bool) *ModelFeatures {
	s.AgentThought = &v
	return s
}

func (s *ModelFeatures) SetMultiToolCall(v bool) *ModelFeatures {
	s.MultiToolCall = &v
	return s
}

func (s *ModelFeatures) SetStreamToolCall(v bool) *ModelFeatures {
	s.StreamToolCall = &v
	return s
}

func (s *ModelFeatures) SetToolCall(v bool) *ModelFeatures {
	s.ToolCall = &v
	return s
}

func (s *ModelFeatures) SetVision(v bool) *ModelFeatures {
	s.Vision = &v
	return s
}

func (s *ModelFeatures) Validate() error {
	return dara.Validate(s)
}
