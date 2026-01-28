// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDataAgentSessionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDMSUnit(v string) *CreateDataAgentSessionRequest
	GetDMSUnit() *string
	SetFile(v string) *CreateDataAgentSessionRequest
	GetFile() *string
	SetSessionConfig(v *CreateDataAgentSessionRequestSessionConfig) *CreateDataAgentSessionRequest
	GetSessionConfig() *CreateDataAgentSessionRequestSessionConfig
	SetTitle(v string) *CreateDataAgentSessionRequest
	GetTitle() *string
	SetWorkspaceId(v string) *CreateDataAgentSessionRequest
	GetWorkspaceId() *string
}

type CreateDataAgentSessionRequest struct {
	// example:
	//
	// cn-hangzhou
	DMSUnit *string `json:"DMSUnit,omitempty" xml:"DMSUnit,omitempty"`
	// example:
	//
	// f-8*******01m
	File          *string                                     `json:"File,omitempty" xml:"File,omitempty"`
	SessionConfig *CreateDataAgentSessionRequestSessionConfig `json:"SessionConfig,omitempty" xml:"SessionConfig,omitempty" type:"Struct"`
	Title         *string                                     `json:"Title,omitempty" xml:"Title,omitempty"`
	// example:
	//
	// 12****
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s CreateDataAgentSessionRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDataAgentSessionRequest) GoString() string {
	return s.String()
}

func (s *CreateDataAgentSessionRequest) GetDMSUnit() *string {
	return s.DMSUnit
}

func (s *CreateDataAgentSessionRequest) GetFile() *string {
	return s.File
}

func (s *CreateDataAgentSessionRequest) GetSessionConfig() *CreateDataAgentSessionRequestSessionConfig {
	return s.SessionConfig
}

func (s *CreateDataAgentSessionRequest) GetTitle() *string {
	return s.Title
}

func (s *CreateDataAgentSessionRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *CreateDataAgentSessionRequest) SetDMSUnit(v string) *CreateDataAgentSessionRequest {
	s.DMSUnit = &v
	return s
}

func (s *CreateDataAgentSessionRequest) SetFile(v string) *CreateDataAgentSessionRequest {
	s.File = &v
	return s
}

func (s *CreateDataAgentSessionRequest) SetSessionConfig(v *CreateDataAgentSessionRequestSessionConfig) *CreateDataAgentSessionRequest {
	s.SessionConfig = v
	return s
}

func (s *CreateDataAgentSessionRequest) SetTitle(v string) *CreateDataAgentSessionRequest {
	s.Title = &v
	return s
}

func (s *CreateDataAgentSessionRequest) SetWorkspaceId(v string) *CreateDataAgentSessionRequest {
	s.WorkspaceId = &v
	return s
}

func (s *CreateDataAgentSessionRequest) Validate() error {
	if s.SessionConfig != nil {
		if err := s.SessionConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateDataAgentSessionRequestSessionConfig struct {
	// example:
	//
	// ca-e*******ckd
	CustomAgentId *string `json:"CustomAgentId,omitempty" xml:"CustomAgentId,omitempty"`
	// example:
	//
	// debug
	CustomAgentStage *string `json:"CustomAgentStage,omitempty" xml:"CustomAgentStage,omitempty"`
	// example:
	//
	// false
	EnableSearch *bool `json:"EnableSearch,omitempty" xml:"EnableSearch,omitempty"`
	// example:
	//
	// CHINESE
	Language     *string   `json:"Language,omitempty" xml:"Language,omitempty"`
	McpServerIds []*string `json:"McpServerIds,omitempty" xml:"McpServerIds,omitempty" type:"Repeated"`
	// example:
	//
	// ANALYSIS
	Mode          *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	UserOssBucket *string `json:"UserOssBucket,omitempty" xml:"UserOssBucket,omitempty"`
}

func (s CreateDataAgentSessionRequestSessionConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateDataAgentSessionRequestSessionConfig) GoString() string {
	return s.String()
}

func (s *CreateDataAgentSessionRequestSessionConfig) GetCustomAgentId() *string {
	return s.CustomAgentId
}

func (s *CreateDataAgentSessionRequestSessionConfig) GetCustomAgentStage() *string {
	return s.CustomAgentStage
}

func (s *CreateDataAgentSessionRequestSessionConfig) GetEnableSearch() *bool {
	return s.EnableSearch
}

func (s *CreateDataAgentSessionRequestSessionConfig) GetLanguage() *string {
	return s.Language
}

func (s *CreateDataAgentSessionRequestSessionConfig) GetMcpServerIds() []*string {
	return s.McpServerIds
}

func (s *CreateDataAgentSessionRequestSessionConfig) GetMode() *string {
	return s.Mode
}

func (s *CreateDataAgentSessionRequestSessionConfig) GetUserOssBucket() *string {
	return s.UserOssBucket
}

func (s *CreateDataAgentSessionRequestSessionConfig) SetCustomAgentId(v string) *CreateDataAgentSessionRequestSessionConfig {
	s.CustomAgentId = &v
	return s
}

func (s *CreateDataAgentSessionRequestSessionConfig) SetCustomAgentStage(v string) *CreateDataAgentSessionRequestSessionConfig {
	s.CustomAgentStage = &v
	return s
}

func (s *CreateDataAgentSessionRequestSessionConfig) SetEnableSearch(v bool) *CreateDataAgentSessionRequestSessionConfig {
	s.EnableSearch = &v
	return s
}

func (s *CreateDataAgentSessionRequestSessionConfig) SetLanguage(v string) *CreateDataAgentSessionRequestSessionConfig {
	s.Language = &v
	return s
}

func (s *CreateDataAgentSessionRequestSessionConfig) SetMcpServerIds(v []*string) *CreateDataAgentSessionRequestSessionConfig {
	s.McpServerIds = v
	return s
}

func (s *CreateDataAgentSessionRequestSessionConfig) SetMode(v string) *CreateDataAgentSessionRequestSessionConfig {
	s.Mode = &v
	return s
}

func (s *CreateDataAgentSessionRequestSessionConfig) SetUserOssBucket(v string) *CreateDataAgentSessionRequestSessionConfig {
	s.UserOssBucket = &v
	return s
}

func (s *CreateDataAgentSessionRequestSessionConfig) Validate() error {
	return dara.Validate(s)
}
