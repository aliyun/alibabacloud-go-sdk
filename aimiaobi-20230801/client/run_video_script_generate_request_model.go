// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunVideoScriptGenerateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLanguage(v string) *RunVideoScriptGenerateRequest
	GetLanguage() *string
	SetPrompt(v string) *RunVideoScriptGenerateRequest
	GetPrompt() *string
	SetScriptLength(v string) *RunVideoScriptGenerateRequest
	GetScriptLength() *string
	SetScriptNumber(v int32) *RunVideoScriptGenerateRequest
	GetScriptNumber() *int32
	SetUseSearch(v bool) *RunVideoScriptGenerateRequest
	GetUseSearch() *bool
	SetWorkspaceId(v string) *RunVideoScriptGenerateRequest
	GetWorkspaceId() *string
}

type RunVideoScriptGenerateRequest struct {
	// example:
	//
	// en-US
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// This parameter is required.
	Prompt *string `json:"Prompt,omitempty" xml:"Prompt,omitempty"`
	// example:
	//
	// >=300
	ScriptLength *string `json:"ScriptLength,omitempty" xml:"ScriptLength,omitempty"`
	// example:
	//
	// 2
	ScriptNumber *int32 `json:"ScriptNumber,omitempty" xml:"ScriptNumber,omitempty"`
	// example:
	//
	// true
	UseSearch *bool `json:"UseSearch,omitempty" xml:"UseSearch,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// llm-xx
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s RunVideoScriptGenerateRequest) String() string {
	return dara.Prettify(s)
}

func (s RunVideoScriptGenerateRequest) GoString() string {
	return s.String()
}

func (s *RunVideoScriptGenerateRequest) GetLanguage() *string {
	return s.Language
}

func (s *RunVideoScriptGenerateRequest) GetPrompt() *string {
	return s.Prompt
}

func (s *RunVideoScriptGenerateRequest) GetScriptLength() *string {
	return s.ScriptLength
}

func (s *RunVideoScriptGenerateRequest) GetScriptNumber() *int32 {
	return s.ScriptNumber
}

func (s *RunVideoScriptGenerateRequest) GetUseSearch() *bool {
	return s.UseSearch
}

func (s *RunVideoScriptGenerateRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *RunVideoScriptGenerateRequest) SetLanguage(v string) *RunVideoScriptGenerateRequest {
	s.Language = &v
	return s
}

func (s *RunVideoScriptGenerateRequest) SetPrompt(v string) *RunVideoScriptGenerateRequest {
	s.Prompt = &v
	return s
}

func (s *RunVideoScriptGenerateRequest) SetScriptLength(v string) *RunVideoScriptGenerateRequest {
	s.ScriptLength = &v
	return s
}

func (s *RunVideoScriptGenerateRequest) SetScriptNumber(v int32) *RunVideoScriptGenerateRequest {
	s.ScriptNumber = &v
	return s
}

func (s *RunVideoScriptGenerateRequest) SetUseSearch(v bool) *RunVideoScriptGenerateRequest {
	s.UseSearch = &v
	return s
}

func (s *RunVideoScriptGenerateRequest) SetWorkspaceId(v string) *RunVideoScriptGenerateRequest {
	s.WorkspaceId = &v
	return s
}

func (s *RunVideoScriptGenerateRequest) Validate() error {
	return dara.Validate(s)
}
