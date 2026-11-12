// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHook interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *Hook
	GetDescription() *string
	SetEnable(v bool) *Hook
	GetEnable() *bool
	SetEvent(v string) *Hook
	GetEvent() *string
	SetExecution(v *HookExecution) *Hook
	GetExecution() *HookExecution
	SetMatch(v []*string) *Hook
	GetMatch() []*string
	SetName(v string) *Hook
	GetName() *string
	SetSchemaVersion(v string) *Hook
	GetSchemaVersion() *string
	SetScope(v string) *Hook
	GetScope() *string
}

type Hook struct {
	// The description.
	//
	// example:
	//
	// 检查工具调用中的命令是否安全。
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// Specifies whether the Hook is enabled. The Hook is executed only when this parameter is set to true. If not configured, the Hook is disabled.
	//
	// example:
	//
	// true
	Enable *bool `json:"enable,omitempty" xml:"enable,omitempty"`
	// The event processed by the Hook. Each Hook processes only one event.
	//
	// example:
	//
	// ToolStart
	Event *string `json:"event,omitempty" xml:"event,omitempty"`
	// The execution configuration that specifies the synchronous execution, total timeout, number of retries, and invalid return value handling policy for the Hook.
	Execution *HookExecution `json:"execution,omitempty" xml:"execution,omitempty" type:"Struct"`
	// The list of tool names. Only ToolStart and ToolEnd events perform exact matching based on the actual tool name. Tool names such as Read, Write, Bash, and Edit are supported, but wildcards are not supported. An empty array does not restrict tools. This field is ignored for other events.
	//
	// example:
	//
	// ["Read","Write","Bash","Edit"]
	Match []*string `json:"match,omitempty" xml:"match,omitempty" type:"Repeated"`
	// The Hook name. Only digits, letters, underscores, and hyphens are allowed. The name cannot be empty. Hooks with the same name are still executed independently in the order of the list.
	//
	// example:
	//
	// command-guard
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The configuration protocol version.
	//
	// example:
	//
	// v1
	SchemaVersion *string `json:"schemaVersion,omitempty" xml:"schemaVersion,omitempty"`
	// The effective scope. Valid values: all, chat, and mission. A value of all indicates that the scope includes Chat and Mission. A value of chat or mission indicates that the scope applies only to the corresponding scenario.
	//
	// example:
	//
	// all
	Scope *string `json:"scope,omitempty" xml:"scope,omitempty"`
}

func (s Hook) String() string {
	return dara.Prettify(s)
}

func (s Hook) GoString() string {
	return s.String()
}

func (s *Hook) GetDescription() *string {
	return s.Description
}

func (s *Hook) GetEnable() *bool {
	return s.Enable
}

func (s *Hook) GetEvent() *string {
	return s.Event
}

func (s *Hook) GetExecution() *HookExecution {
	return s.Execution
}

func (s *Hook) GetMatch() []*string {
	return s.Match
}

func (s *Hook) GetName() *string {
	return s.Name
}

func (s *Hook) GetSchemaVersion() *string {
	return s.SchemaVersion
}

func (s *Hook) GetScope() *string {
	return s.Scope
}

func (s *Hook) SetDescription(v string) *Hook {
	s.Description = &v
	return s
}

func (s *Hook) SetEnable(v bool) *Hook {
	s.Enable = &v
	return s
}

func (s *Hook) SetEvent(v string) *Hook {
	s.Event = &v
	return s
}

func (s *Hook) SetExecution(v *HookExecution) *Hook {
	s.Execution = v
	return s
}

func (s *Hook) SetMatch(v []*string) *Hook {
	s.Match = v
	return s
}

func (s *Hook) SetName(v string) *Hook {
	s.Name = &v
	return s
}

func (s *Hook) SetSchemaVersion(v string) *Hook {
	s.SchemaVersion = &v
	return s
}

func (s *Hook) SetScope(v string) *Hook {
	s.Scope = &v
	return s
}

func (s *Hook) Validate() error {
	if s.Execution != nil {
		if err := s.Execution.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type HookExecution struct {
	// The handling policy for invalid return values. If the handler return value remains invalid after retries, it is processed as allow or reject.
	//
	// example:
	//
	// allow
	FailResult *string `json:"failResult,omitempty" xml:"failResult,omitempty"`
	// The execution instruction. The prompt type indicates a prompt, the bash type indicates a Bash script, the python3 type indicates a Python 3 script, and the hint type indicates contextual text directly appended. A non-empty UTF-8 text must be provided for all four types.
	//
	// example:
	//
	// 检查当前工具调用是否安全。
	Instruction *string `json:"instruction,omitempty" xml:"instruction,omitempty"`
	// The maximum number of additional attempts after a failure. Valid values: 0 to 3.
	//
	// example:
	//
	// 0
	MaxRetry *int64 `json:"maxRetry,omitempty" xml:"maxRetry,omitempty"`
	// The timeout period in seconds. Valid values: 1 to 30. All attempts and waits for the Hook share this timeout period.
	//
	// example:
	//
	// 10
	TimeoutSeconds *int64 `json:"timeoutSeconds,omitempty" xml:"timeoutSeconds,omitempty"`
	// The handler type. Valid values: prompt, bash, python3, and hint.
	//
	// example:
	//
	// prompt
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s HookExecution) String() string {
	return dara.Prettify(s)
}

func (s HookExecution) GoString() string {
	return s.String()
}

func (s *HookExecution) GetFailResult() *string {
	return s.FailResult
}

func (s *HookExecution) GetInstruction() *string {
	return s.Instruction
}

func (s *HookExecution) GetMaxRetry() *int64 {
	return s.MaxRetry
}

func (s *HookExecution) GetTimeoutSeconds() *int64 {
	return s.TimeoutSeconds
}

func (s *HookExecution) GetType() *string {
	return s.Type
}

func (s *HookExecution) SetFailResult(v string) *HookExecution {
	s.FailResult = &v
	return s
}

func (s *HookExecution) SetInstruction(v string) *HookExecution {
	s.Instruction = &v
	return s
}

func (s *HookExecution) SetMaxRetry(v int64) *HookExecution {
	s.MaxRetry = &v
	return s
}

func (s *HookExecution) SetTimeoutSeconds(v int64) *HookExecution {
	s.TimeoutSeconds = &v
	return s
}

func (s *HookExecution) SetType(v string) *HookExecution {
	s.Type = &v
	return s
}

func (s *HookExecution) Validate() error {
	return dara.Validate(s)
}
