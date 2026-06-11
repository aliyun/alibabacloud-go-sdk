// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDataAgentSessionShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDMSUnit(v string) *CreateDataAgentSessionShrinkRequest
	GetDMSUnit() *string
	SetFile(v string) *CreateDataAgentSessionShrinkRequest
	GetFile() *string
	SetSessionConfigShrink(v string) *CreateDataAgentSessionShrinkRequest
	GetSessionConfigShrink() *string
	SetTitle(v string) *CreateDataAgentSessionShrinkRequest
	GetTitle() *string
	SetWorkspaceId(v string) *CreateDataAgentSessionShrinkRequest
	GetWorkspaceId() *string
}

type CreateDataAgentSessionShrinkRequest struct {
	// The current DMS unit.
	//
	// example:
	//
	// cn-hangzhou
	DMSUnit *string `json:"DMSUnit,omitempty" xml:"DMSUnit,omitempty"`
	// The file ID.
	//
	// example:
	//
	// f-8*******01m
	File *string `json:"File,omitempty" xml:"File,omitempty"`
	// The session configuration.
	SessionConfigShrink *string `json:"SessionConfig,omitempty" xml:"SessionConfig,omitempty"`
	// The session title. Supports Chinese and English. Maximum length: 128 characters.
	//
	// example:
	//
	// 帮我分析一下这份数据，给出报告。
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// The workspace ID.
	//
	// example:
	//
	// 12****
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s CreateDataAgentSessionShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDataAgentSessionShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateDataAgentSessionShrinkRequest) GetDMSUnit() *string {
	return s.DMSUnit
}

func (s *CreateDataAgentSessionShrinkRequest) GetFile() *string {
	return s.File
}

func (s *CreateDataAgentSessionShrinkRequest) GetSessionConfigShrink() *string {
	return s.SessionConfigShrink
}

func (s *CreateDataAgentSessionShrinkRequest) GetTitle() *string {
	return s.Title
}

func (s *CreateDataAgentSessionShrinkRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *CreateDataAgentSessionShrinkRequest) SetDMSUnit(v string) *CreateDataAgentSessionShrinkRequest {
	s.DMSUnit = &v
	return s
}

func (s *CreateDataAgentSessionShrinkRequest) SetFile(v string) *CreateDataAgentSessionShrinkRequest {
	s.File = &v
	return s
}

func (s *CreateDataAgentSessionShrinkRequest) SetSessionConfigShrink(v string) *CreateDataAgentSessionShrinkRequest {
	s.SessionConfigShrink = &v
	return s
}

func (s *CreateDataAgentSessionShrinkRequest) SetTitle(v string) *CreateDataAgentSessionShrinkRequest {
	s.Title = &v
	return s
}

func (s *CreateDataAgentSessionShrinkRequest) SetWorkspaceId(v string) *CreateDataAgentSessionShrinkRequest {
	s.WorkspaceId = &v
	return s
}

func (s *CreateDataAgentSessionShrinkRequest) Validate() error {
	return dara.Validate(s)
}
