// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iClusterScript interface {
	dara.Model
	String() string
	GoString() string
	SetExecutionFailStrategy(v string) *ClusterScript
	GetExecutionFailStrategy() *string
	SetExecutionMoment(v string) *ClusterScript
	GetExecutionMoment() *string
	SetNodeSelect(v *NodeSelector) *ClusterScript
	GetNodeSelect() *NodeSelector
	SetPriority(v int32) *ClusterScript
	GetPriority() *int32
	SetScriptArgs(v string) *ClusterScript
	GetScriptArgs() *string
	SetScriptName(v string) *ClusterScript
	GetScriptName() *string
	SetScriptPath(v string) *ClusterScript
	GetScriptPath() *string
}

type ClusterScript struct {
	// example:
	//
	// 取值:FAILED_CONTINUE, FAILED_BLOCKED
	ExecutionFailStrategy *string `json:"ExecutionFailStrategy,omitempty" xml:"ExecutionFailStrategy,omitempty"`
	// example:
	//
	// 取值:BEFORE_INSTALL, AFTER_STARTED
	ExecutionMoment *string       `json:"ExecutionMoment,omitempty" xml:"ExecutionMoment,omitempty"`
	NodeSelect      *NodeSelector `json:"NodeSelect,omitempty" xml:"NodeSelect,omitempty"`
	Priority        *int32        `json:"Priority,omitempty" xml:"Priority,omitempty"`
	ScriptArgs      *string       `json:"ScriptArgs,omitempty" xml:"ScriptArgs,omitempty"`
	ScriptName      *string       `json:"ScriptName,omitempty" xml:"ScriptName,omitempty"`
	ScriptPath      *string       `json:"ScriptPath,omitempty" xml:"ScriptPath,omitempty"`
}

func (s ClusterScript) String() string {
	return dara.Prettify(s)
}

func (s ClusterScript) GoString() string {
	return s.String()
}

func (s *ClusterScript) GetExecutionFailStrategy() *string {
	return s.ExecutionFailStrategy
}

func (s *ClusterScript) GetExecutionMoment() *string {
	return s.ExecutionMoment
}

func (s *ClusterScript) GetNodeSelect() *NodeSelector {
	return s.NodeSelect
}

func (s *ClusterScript) GetPriority() *int32 {
	return s.Priority
}

func (s *ClusterScript) GetScriptArgs() *string {
	return s.ScriptArgs
}

func (s *ClusterScript) GetScriptName() *string {
	return s.ScriptName
}

func (s *ClusterScript) GetScriptPath() *string {
	return s.ScriptPath
}

func (s *ClusterScript) SetExecutionFailStrategy(v string) *ClusterScript {
	s.ExecutionFailStrategy = &v
	return s
}

func (s *ClusterScript) SetExecutionMoment(v string) *ClusterScript {
	s.ExecutionMoment = &v
	return s
}

func (s *ClusterScript) SetNodeSelect(v *NodeSelector) *ClusterScript {
	s.NodeSelect = v
	return s
}

func (s *ClusterScript) SetPriority(v int32) *ClusterScript {
	s.Priority = &v
	return s
}

func (s *ClusterScript) SetScriptArgs(v string) *ClusterScript {
	s.ScriptArgs = &v
	return s
}

func (s *ClusterScript) SetScriptName(v string) *ClusterScript {
	s.ScriptName = &v
	return s
}

func (s *ClusterScript) SetScriptPath(v string) *ClusterScript {
	s.ScriptPath = &v
	return s
}

func (s *ClusterScript) Validate() error {
	if s.NodeSelect != nil {
		if err := s.NodeSelect.Validate(); err != nil {
			return err
		}
	}
	return nil
}
