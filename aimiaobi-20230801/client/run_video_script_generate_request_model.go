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
	// The language of the generated script.
	//
	// Recommended values:
	//
	// zh-CN: Chinese
	//
	// en-US: English
	//
	// The default is Chinese.
	//
	// example:
	//
	// en-US
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// The prompt for the video script.
	//
	// This parameter is required.
	//
	// example:
	//
	// 写一篇关于黄山旅游的脚本
	Prompt *string `json:"Prompt,omitempty" xml:"Prompt,omitempty"`
	// The length of the script. Valid values:
	//
	// 20\\~75: 10 to 15 seconds of normal speaking time.
	//
	// 75\\~150: 15 to 30 seconds of normal speaking time.
	//
	// 150\\~300: Approximately 30 to 60 seconds of normal speaking time.
	//
	// \\>=300: 60 seconds or more of normal speaking time.
	//
	// example:
	//
	// >=300
	ScriptLength *string `json:"ScriptLength,omitempty" xml:"ScriptLength,omitempty"`
	// The number of scripts to generate. The default is 1. You can generate a maximum of three scripts at a time.
	//
	// If you specify multiple scripts, the results are returned in parallel streams. The client distinguishes between the streams using different session IDs.
	//
	// example:
	//
	// 2
	ScriptNumber *int32 `json:"ScriptNumber,omitempty" xml:"ScriptNumber,omitempty"`
	// Specifies whether to use an internet search. If you set this to true, the system performs intention recognition and then searches the internet for relevant reference materials.
	//
	// example:
	//
	// true
	UseSearch *bool `json:"UseSearch,omitempty" xml:"UseSearch,omitempty"`
	// The unique ID of the Alibaba Cloud Model Studio workspace. For more information, see [Get a Workspace ID](https://help.aliyun.com/document_detail/2782167.html).
	//
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
