// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePolarClawAgentShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentId(v string) *UpdatePolarClawAgentShrinkRequest
	GetAgentId() *string
	SetApplicationId(v string) *UpdatePolarClawAgentShrinkRequest
	GetApplicationId() *string
	SetAvatar(v string) *UpdatePolarClawAgentShrinkRequest
	GetAvatar() *string
	SetFilesShrink(v string) *UpdatePolarClawAgentShrinkRequest
	GetFilesShrink() *string
	SetIsDefault(v bool) *UpdatePolarClawAgentShrinkRequest
	GetIsDefault() *bool
	SetKeepWorkspaceFiles(v bool) *UpdatePolarClawAgentShrinkRequest
	GetKeepWorkspaceFiles() *bool
	SetModel(v string) *UpdatePolarClawAgentShrinkRequest
	GetModel() *string
	SetName(v string) *UpdatePolarClawAgentShrinkRequest
	GetName() *string
	SetRestart(v bool) *UpdatePolarClawAgentShrinkRequest
	GetRestart() *bool
	SetWorkspace(v string) *UpdatePolarClawAgentShrinkRequest
	GetWorkspace() *string
}

type UpdatePolarClawAgentShrinkRequest struct {
	// Agent ID to update
	//
	// This parameter is required.
	//
	// example:
	//
	// work
	AgentId *string `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	// Application ID
	//
	// This parameter is required.
	//
	// example:
	//
	// pa-**************
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// New avatar
	//
	// example:
	//
	// test
	Avatar *string `json:"Avatar,omitempty" xml:"Avatar,omitempty"`
	// List of files to update
	//
	// example:
	//
	// [{"FileName":"SOUL.md","FileContent":"You are a helpful assistant."}]
	FilesShrink *string `json:"Files,omitempty" xml:"Files,omitempty"`
	// example:
	//
	// true
	IsDefault *bool `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	// example:
	//
	// false
	KeepWorkspaceFiles *bool `json:"KeepWorkspaceFiles,omitempty" xml:"KeepWorkspaceFiles,omitempty"`
	// Model override
	//
	// example:
	//
	// claude-sonnet-4-5
	Model *string `json:"Model,omitempty" xml:"Model,omitempty"`
	// New display name
	//
	// example:
	//
	// Work Bot
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Whether to restart the gateway after creation, default is true
	//
	// example:
	//
	// true
	Restart *bool `json:"Restart,omitempty" xml:"Restart,omitempty"`
	// New workspace directory path
	//
	// example:
	//
	// /home/node/.openclaw/workspace-work-v2
	Workspace *string `json:"Workspace,omitempty" xml:"Workspace,omitempty"`
}

func (s UpdatePolarClawAgentShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdatePolarClawAgentShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdatePolarClawAgentShrinkRequest) GetAgentId() *string {
	return s.AgentId
}

func (s *UpdatePolarClawAgentShrinkRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *UpdatePolarClawAgentShrinkRequest) GetAvatar() *string {
	return s.Avatar
}

func (s *UpdatePolarClawAgentShrinkRequest) GetFilesShrink() *string {
	return s.FilesShrink
}

func (s *UpdatePolarClawAgentShrinkRequest) GetIsDefault() *bool {
	return s.IsDefault
}

func (s *UpdatePolarClawAgentShrinkRequest) GetKeepWorkspaceFiles() *bool {
	return s.KeepWorkspaceFiles
}

func (s *UpdatePolarClawAgentShrinkRequest) GetModel() *string {
	return s.Model
}

func (s *UpdatePolarClawAgentShrinkRequest) GetName() *string {
	return s.Name
}

func (s *UpdatePolarClawAgentShrinkRequest) GetRestart() *bool {
	return s.Restart
}

func (s *UpdatePolarClawAgentShrinkRequest) GetWorkspace() *string {
	return s.Workspace
}

func (s *UpdatePolarClawAgentShrinkRequest) SetAgentId(v string) *UpdatePolarClawAgentShrinkRequest {
	s.AgentId = &v
	return s
}

func (s *UpdatePolarClawAgentShrinkRequest) SetApplicationId(v string) *UpdatePolarClawAgentShrinkRequest {
	s.ApplicationId = &v
	return s
}

func (s *UpdatePolarClawAgentShrinkRequest) SetAvatar(v string) *UpdatePolarClawAgentShrinkRequest {
	s.Avatar = &v
	return s
}

func (s *UpdatePolarClawAgentShrinkRequest) SetFilesShrink(v string) *UpdatePolarClawAgentShrinkRequest {
	s.FilesShrink = &v
	return s
}

func (s *UpdatePolarClawAgentShrinkRequest) SetIsDefault(v bool) *UpdatePolarClawAgentShrinkRequest {
	s.IsDefault = &v
	return s
}

func (s *UpdatePolarClawAgentShrinkRequest) SetKeepWorkspaceFiles(v bool) *UpdatePolarClawAgentShrinkRequest {
	s.KeepWorkspaceFiles = &v
	return s
}

func (s *UpdatePolarClawAgentShrinkRequest) SetModel(v string) *UpdatePolarClawAgentShrinkRequest {
	s.Model = &v
	return s
}

func (s *UpdatePolarClawAgentShrinkRequest) SetName(v string) *UpdatePolarClawAgentShrinkRequest {
	s.Name = &v
	return s
}

func (s *UpdatePolarClawAgentShrinkRequest) SetRestart(v bool) *UpdatePolarClawAgentShrinkRequest {
	s.Restart = &v
	return s
}

func (s *UpdatePolarClawAgentShrinkRequest) SetWorkspace(v string) *UpdatePolarClawAgentShrinkRequest {
	s.Workspace = &v
	return s
}

func (s *UpdatePolarClawAgentShrinkRequest) Validate() error {
	return dara.Validate(s)
}
