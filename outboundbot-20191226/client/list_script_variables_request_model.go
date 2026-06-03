// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListScriptVariablesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListScriptVariablesRequest
	GetInstanceId() *string
	SetSandbox(v bool) *ListScriptVariablesRequest
	GetSandbox() *bool
	SetScriptId(v string) *ListScriptVariablesRequest
	GetScriptId() *string
}

type ListScriptVariablesRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 2e55ad0f-7a07-420e-a9b0-62d66ba7278a
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Sandbox    *bool   `json:"Sandbox,omitempty" xml:"Sandbox,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 62318d71-1128-4fd7-af15-c852cbe26b93
	ScriptId *string `json:"ScriptId,omitempty" xml:"ScriptId,omitempty"`
}

func (s ListScriptVariablesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListScriptVariablesRequest) GoString() string {
	return s.String()
}

func (s *ListScriptVariablesRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListScriptVariablesRequest) GetSandbox() *bool {
	return s.Sandbox
}

func (s *ListScriptVariablesRequest) GetScriptId() *string {
	return s.ScriptId
}

func (s *ListScriptVariablesRequest) SetInstanceId(v string) *ListScriptVariablesRequest {
	s.InstanceId = &v
	return s
}

func (s *ListScriptVariablesRequest) SetSandbox(v bool) *ListScriptVariablesRequest {
	s.Sandbox = &v
	return s
}

func (s *ListScriptVariablesRequest) SetScriptId(v string) *ListScriptVariablesRequest {
	s.ScriptId = &v
	return s
}

func (s *ListScriptVariablesRequest) Validate() error {
	return dara.Validate(s)
}
