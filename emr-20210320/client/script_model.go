// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iScript interface {
	dara.Model
	String() string
	GoString() string
	SetExecutionFailStrategy(v string) *Script
	GetExecutionFailStrategy() *string
	SetExecutionMoment(v string) *Script
	GetExecutionMoment() *string
	SetNodeSelector(v *NodeSelector) *Script
	GetNodeSelector() *NodeSelector
	SetPriority(v int32) *Script
	GetPriority() *int32
	SetScriptArgs(v string) *Script
	GetScriptArgs() *string
	SetScriptName(v string) *Script
	GetScriptName() *string
	SetScriptPath(v string) *Script
	GetScriptPath() *string
}

type Script struct {
	// 执行失败策略。
	//
	// example:
	//
	// FAILED_CONTINUE
	ExecutionFailStrategy *string `json:"ExecutionFailStrategy,omitempty" xml:"ExecutionFailStrategy,omitempty"`
	// 脚本的执行时机。
	//
	// example:
	//
	// BEFORE_INSTALL
	ExecutionMoment *string `json:"ExecutionMoment,omitempty" xml:"ExecutionMoment,omitempty"`
	// 节点选择器。
	//
	// This parameter is required.
	NodeSelector *NodeSelector `json:"NodeSelector,omitempty" xml:"NodeSelector,omitempty"`
	// Deprecated
	//
	// 脚本执行优先级。取值范围：1~100。
	//
	// example:
	//
	// 1
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// 脚本执行参数。
	//
	// example:
	//
	// -host 10.0.10.5 -m 30
	ScriptArgs *string `json:"ScriptArgs,omitempty" xml:"ScriptArgs,omitempty"`
	// 脚本名称。长度为1~64个字符，必须以大小字母或中文开头，不能以http://和https://开头。可以包含中文、英文、数字、下划线（_）、或者短划线（-）
	//
	// This parameter is required.
	//
	// example:
	//
	// 脚本名-1
	ScriptName *string `json:"ScriptName,omitempty" xml:"ScriptName,omitempty"`
	// 脚本所在OSS路径。
	//
	// This parameter is required.
	//
	// example:
	//
	// oss://bucket1/update_hosts.sh
	ScriptPath *string `json:"ScriptPath,omitempty" xml:"ScriptPath,omitempty"`
}

func (s Script) String() string {
	return dara.Prettify(s)
}

func (s Script) GoString() string {
	return s.String()
}

func (s *Script) GetExecutionFailStrategy() *string {
	return s.ExecutionFailStrategy
}

func (s *Script) GetExecutionMoment() *string {
	return s.ExecutionMoment
}

func (s *Script) GetNodeSelector() *NodeSelector {
	return s.NodeSelector
}

func (s *Script) GetPriority() *int32 {
	return s.Priority
}

func (s *Script) GetScriptArgs() *string {
	return s.ScriptArgs
}

func (s *Script) GetScriptName() *string {
	return s.ScriptName
}

func (s *Script) GetScriptPath() *string {
	return s.ScriptPath
}

func (s *Script) SetExecutionFailStrategy(v string) *Script {
	s.ExecutionFailStrategy = &v
	return s
}

func (s *Script) SetExecutionMoment(v string) *Script {
	s.ExecutionMoment = &v
	return s
}

func (s *Script) SetNodeSelector(v *NodeSelector) *Script {
	s.NodeSelector = v
	return s
}

func (s *Script) SetPriority(v int32) *Script {
	s.Priority = &v
	return s
}

func (s *Script) SetScriptArgs(v string) *Script {
	s.ScriptArgs = &v
	return s
}

func (s *Script) SetScriptName(v string) *Script {
	s.ScriptName = &v
	return s
}

func (s *Script) SetScriptPath(v string) *Script {
	s.ScriptPath = &v
	return s
}

func (s *Script) Validate() error {
	if s.NodeSelector != nil {
		if err := s.NodeSelector.Validate(); err != nil {
			return err
		}
	}
	return nil
}
