// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetAIAgentVoiceprintShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInputShrink(v string) *SetAIAgentVoiceprintShrinkRequest
	GetInputShrink() *string
	SetVoiceprintId(v string) *SetAIAgentVoiceprintShrinkRequest
	GetVoiceprintId() *string
}

type SetAIAgentVoiceprintShrinkRequest struct {
	// The input file.
	InputShrink *string `json:"Input,omitempty" xml:"Input,omitempty"`
	// A unique identifier for the voiceprint. Generate this ID based on your own business rules. Requirement: 1 to 127 characters in length.
	//
	// example:
	//
	// vp_1699123456_8527
	VoiceprintId *string `json:"VoiceprintId,omitempty" xml:"VoiceprintId,omitempty"`
}

func (s SetAIAgentVoiceprintShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s SetAIAgentVoiceprintShrinkRequest) GoString() string {
	return s.String()
}

func (s *SetAIAgentVoiceprintShrinkRequest) GetInputShrink() *string {
	return s.InputShrink
}

func (s *SetAIAgentVoiceprintShrinkRequest) GetVoiceprintId() *string {
	return s.VoiceprintId
}

func (s *SetAIAgentVoiceprintShrinkRequest) SetInputShrink(v string) *SetAIAgentVoiceprintShrinkRequest {
	s.InputShrink = &v
	return s
}

func (s *SetAIAgentVoiceprintShrinkRequest) SetVoiceprintId(v string) *SetAIAgentVoiceprintShrinkRequest {
	s.VoiceprintId = &v
	return s
}

func (s *SetAIAgentVoiceprintShrinkRequest) Validate() error {
	return dara.Validate(s)
}
