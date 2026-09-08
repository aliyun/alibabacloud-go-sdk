// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAgentDataSemanticsStageProgress interface {
	dara.Model
	String() string
	GoString() string
	SetStage(v string) *AgentDataSemanticsStageProgress
	GetStage() *string
	SetStatus(v string) *AgentDataSemanticsStageProgress
	GetStatus() *string
}

type AgentDataSemanticsStageProgress struct {
	// The stage name.
	//
	// This parameter is required.
	//
	// example:
	//
	// PROFILE
	Stage *string `json:"Stage,omitempty" xml:"Stage,omitempty"`
	// The stage status.
	//
	// This parameter is required.
	//
	// example:
	//
	// SUCCEEDED
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s AgentDataSemanticsStageProgress) String() string {
	return dara.Prettify(s)
}

func (s AgentDataSemanticsStageProgress) GoString() string {
	return s.String()
}

func (s *AgentDataSemanticsStageProgress) GetStage() *string {
	return s.Stage
}

func (s *AgentDataSemanticsStageProgress) GetStatus() *string {
	return s.Status
}

func (s *AgentDataSemanticsStageProgress) SetStage(v string) *AgentDataSemanticsStageProgress {
	s.Stage = &v
	return s
}

func (s *AgentDataSemanticsStageProgress) SetStatus(v string) *AgentDataSemanticsStageProgress {
	s.Status = &v
	return s
}

func (s *AgentDataSemanticsStageProgress) Validate() error {
	return dara.Validate(s)
}
