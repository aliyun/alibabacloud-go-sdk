// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePolarClawAgentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentId(v string) *UpdatePolarClawAgentRequest
	GetAgentId() *string
	SetApplicationId(v string) *UpdatePolarClawAgentRequest
	GetApplicationId() *string
	SetAvatar(v string) *UpdatePolarClawAgentRequest
	GetAvatar() *string
	SetFiles(v []*UpdatePolarClawAgentRequestFiles) *UpdatePolarClawAgentRequest
	GetFiles() []*UpdatePolarClawAgentRequestFiles
	SetModel(v string) *UpdatePolarClawAgentRequest
	GetModel() *string
	SetName(v string) *UpdatePolarClawAgentRequest
	GetName() *string
	SetRestart(v bool) *UpdatePolarClawAgentRequest
	GetRestart() *bool
	SetWorkspace(v string) *UpdatePolarClawAgentRequest
	GetWorkspace() *string
}

type UpdatePolarClawAgentRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// work
	AgentId *string `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// pa-**************
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// example:
	//
	// test
	Avatar *string `json:"Avatar,omitempty" xml:"Avatar,omitempty"`
	// example:
	//
	// [{"FileName":"SOUL.md","FileContent":"You are a helpful assistant."}]
	Files []*UpdatePolarClawAgentRequestFiles `json:"Files,omitempty" xml:"Files,omitempty" type:"Repeated"`
	// example:
	//
	// claude-sonnet-4-5
	Model *string `json:"Model,omitempty" xml:"Model,omitempty"`
	// example:
	//
	// Work Bot
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// true
	Restart *bool `json:"Restart,omitempty" xml:"Restart,omitempty"`
	// example:
	//
	// /home/node/.openclaw/workspace-work-v2
	Workspace *string `json:"Workspace,omitempty" xml:"Workspace,omitempty"`
}

func (s UpdatePolarClawAgentRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdatePolarClawAgentRequest) GoString() string {
	return s.String()
}

func (s *UpdatePolarClawAgentRequest) GetAgentId() *string {
	return s.AgentId
}

func (s *UpdatePolarClawAgentRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *UpdatePolarClawAgentRequest) GetAvatar() *string {
	return s.Avatar
}

func (s *UpdatePolarClawAgentRequest) GetFiles() []*UpdatePolarClawAgentRequestFiles {
	return s.Files
}

func (s *UpdatePolarClawAgentRequest) GetModel() *string {
	return s.Model
}

func (s *UpdatePolarClawAgentRequest) GetName() *string {
	return s.Name
}

func (s *UpdatePolarClawAgentRequest) GetRestart() *bool {
	return s.Restart
}

func (s *UpdatePolarClawAgentRequest) GetWorkspace() *string {
	return s.Workspace
}

func (s *UpdatePolarClawAgentRequest) SetAgentId(v string) *UpdatePolarClawAgentRequest {
	s.AgentId = &v
	return s
}

func (s *UpdatePolarClawAgentRequest) SetApplicationId(v string) *UpdatePolarClawAgentRequest {
	s.ApplicationId = &v
	return s
}

func (s *UpdatePolarClawAgentRequest) SetAvatar(v string) *UpdatePolarClawAgentRequest {
	s.Avatar = &v
	return s
}

func (s *UpdatePolarClawAgentRequest) SetFiles(v []*UpdatePolarClawAgentRequestFiles) *UpdatePolarClawAgentRequest {
	s.Files = v
	return s
}

func (s *UpdatePolarClawAgentRequest) SetModel(v string) *UpdatePolarClawAgentRequest {
	s.Model = &v
	return s
}

func (s *UpdatePolarClawAgentRequest) SetName(v string) *UpdatePolarClawAgentRequest {
	s.Name = &v
	return s
}

func (s *UpdatePolarClawAgentRequest) SetRestart(v bool) *UpdatePolarClawAgentRequest {
	s.Restart = &v
	return s
}

func (s *UpdatePolarClawAgentRequest) SetWorkspace(v string) *UpdatePolarClawAgentRequest {
	s.Workspace = &v
	return s
}

func (s *UpdatePolarClawAgentRequest) Validate() error {
	if s.Files != nil {
		for _, item := range s.Files {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdatePolarClawAgentRequestFiles struct {
	// example:
	//
	// You are a helpful assistant.
	FileContent *string `json:"FileContent,omitempty" xml:"FileContent,omitempty"`
	// example:
	//
	// SOUL.md
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
}

func (s UpdatePolarClawAgentRequestFiles) String() string {
	return dara.Prettify(s)
}

func (s UpdatePolarClawAgentRequestFiles) GoString() string {
	return s.String()
}

func (s *UpdatePolarClawAgentRequestFiles) GetFileContent() *string {
	return s.FileContent
}

func (s *UpdatePolarClawAgentRequestFiles) GetFileName() *string {
	return s.FileName
}

func (s *UpdatePolarClawAgentRequestFiles) SetFileContent(v string) *UpdatePolarClawAgentRequestFiles {
	s.FileContent = &v
	return s
}

func (s *UpdatePolarClawAgentRequestFiles) SetFileName(v string) *UpdatePolarClawAgentRequestFiles {
	s.FileName = &v
	return s
}

func (s *UpdatePolarClawAgentRequestFiles) Validate() error {
	return dara.Validate(s)
}
