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
	// The execution failure strategy. Valid values:
	//
	// - `FAILED_CONTINUE`: If the script fails, cluster creation or scaling continues.
	//
	// - `FAILED_BLOCK`: If the script fails, cluster creation or scaling is blocked.
	//
	// example:
	//
	// FAILED_CONTINUE
	ExecutionFailStrategy *string `json:"ExecutionFailStrategy,omitempty" xml:"ExecutionFailStrategy,omitempty"`
	// The execution timing for the script. Valid values:
	//
	// - `BEFORE_INSTALL`: The script runs before applications are installed.
	//
	// - `AFTER_STARTED`: The script runs after applications start.
	//
	// example:
	//
	// BEFORE_INSTALL
	ExecutionMoment *string `json:"ExecutionMoment,omitempty" xml:"ExecutionMoment,omitempty"`
	// Specifies the nodes on which the script runs.
	//
	// This parameter is required.
	NodeSelector *NodeSelector `json:"NodeSelector,omitempty" xml:"NodeSelector,omitempty"`
	// Deprecated
	//
	// > This parameter is deprecated. Scripts run in the order they are defined.
	//
	// example:
	//
	// 1
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The optional script execution arguments.
	//
	// example:
	//
	// -host 10.0.10.5 -m 30
	ScriptArgs *string `json:"ScriptArgs,omitempty" xml:"ScriptArgs,omitempty"`
	// The required script name. The name must be 1 to 64 characters long and start with a letter or a Chinese character. It cannot start with `http://` or `https://`. It can contain Chinese characters, letters, numbers, underscores (`_`), or hyphens (`-`).
	//
	// This parameter is required.
	//
	// example:
	//
	// 脚本名-1
	ScriptName *string `json:"ScriptName,omitempty" xml:"ScriptName,omitempty"`
	// The required path to the script in Object Storage Service (OSS). The path must start with `oss://`.
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
