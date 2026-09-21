// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDataAgentApplication interface {
	dara.Model
	String() string
	GoString() string
	SetAgentId(v string) *DataAgentApplication
	GetAgentId() *string
	SetAppId(v string) *DataAgentApplication
	GetAppId() *string
	SetAppName(v string) *DataAgentApplication
	GetAppName() *string
	SetApplicationExtraInfo(v string) *DataAgentApplication
	GetApplicationExtraInfo() *string
	SetCreatorName(v string) *DataAgentApplication
	GetCreatorName() *string
	SetCreatorUid(v string) *DataAgentApplication
	GetCreatorUid() *string
	SetDescription(v string) *DataAgentApplication
	GetDescription() *string
	SetGmtCreated(v string) *DataAgentApplication
	GetGmtCreated() *string
	SetGmtModified(v string) *DataAgentApplication
	GetGmtModified() *string
	SetMainUid(v string) *DataAgentApplication
	GetMainUid() *string
	SetRegion(v string) *DataAgentApplication
	GetRegion() *string
	SetSessionId(v string) *DataAgentApplication
	GetSessionId() *string
	SetStatus(v string) *DataAgentApplication
	GetStatus() *string
	SetWorkspaceId(v string) *DataAgentApplication
	GetWorkspaceId() *string
}

type DataAgentApplication struct {
	// The ID of the currently associated Data Agent.
	//
	// example:
	//
	// avgwuxxxxxxxxhldt7el9
	AgentId *string `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	// The stable identifier of the application.
	//
	// example:
	//
	// ac6izw6xxxxxxxxxxx3ulya0d
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The name of the application.
	//
	// example:
	//
	// test-app-name
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The extension information of the application.
	//
	// example:
	//
	// {}
	ApplicationExtraInfo *string `json:"ApplicationExtraInfo,omitempty" xml:"ApplicationExtraInfo,omitempty"`
	// The name of the application creator.
	//
	// example:
	//
	// test-name
	CreatorName *string `json:"CreatorName,omitempty" xml:"CreatorName,omitempty"`
	// The UID of the application owner.
	//
	// example:
	//
	// 30031588888885853
	CreatorUid *string `json:"CreatorUid,omitempty" xml:"CreatorUid,omitempty"`
	// The description of the application. The description can be up to 250 characters in length.
	//
	// example:
	//
	// this is a test application
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The time when the application was created.
	//
	// example:
	//
	// 2026-09-12T10:40:12.000+00:00
	GmtCreated *string `json:"GmtCreated,omitempty" xml:"GmtCreated,omitempty"`
	// The time when the application was last modified.
	//
	// example:
	//
	// 2026-09-12T10:40:12.000+00:00
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The UID of the Alibaba Cloud account.
	//
	// example:
	//
	// 1673828888852166
	MainUid *string `json:"MainUid,omitempty" xml:"MainUid,omitempty"`
	// The region.
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The ID of the current or most recently associated session.
	//
	// example:
	//
	// axc3lsxxxxxxxxxdapwe
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// The status of the application. Valid values:
	//
	// - REGISTERED
	//
	// - DEPLOYING
	//
	// - DEPLOYED
	//
	// example:
	//
	// REGISTERED
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the workspace.
	//
	// example:
	//
	// 3700inkdc2y7zs0r37m5pika6
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s DataAgentApplication) String() string {
	return dara.Prettify(s)
}

func (s DataAgentApplication) GoString() string {
	return s.String()
}

func (s *DataAgentApplication) GetAgentId() *string {
	return s.AgentId
}

func (s *DataAgentApplication) GetAppId() *string {
	return s.AppId
}

func (s *DataAgentApplication) GetAppName() *string {
	return s.AppName
}

func (s *DataAgentApplication) GetApplicationExtraInfo() *string {
	return s.ApplicationExtraInfo
}

func (s *DataAgentApplication) GetCreatorName() *string {
	return s.CreatorName
}

func (s *DataAgentApplication) GetCreatorUid() *string {
	return s.CreatorUid
}

func (s *DataAgentApplication) GetDescription() *string {
	return s.Description
}

func (s *DataAgentApplication) GetGmtCreated() *string {
	return s.GmtCreated
}

func (s *DataAgentApplication) GetGmtModified() *string {
	return s.GmtModified
}

func (s *DataAgentApplication) GetMainUid() *string {
	return s.MainUid
}

func (s *DataAgentApplication) GetRegion() *string {
	return s.Region
}

func (s *DataAgentApplication) GetSessionId() *string {
	return s.SessionId
}

func (s *DataAgentApplication) GetStatus() *string {
	return s.Status
}

func (s *DataAgentApplication) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *DataAgentApplication) SetAgentId(v string) *DataAgentApplication {
	s.AgentId = &v
	return s
}

func (s *DataAgentApplication) SetAppId(v string) *DataAgentApplication {
	s.AppId = &v
	return s
}

func (s *DataAgentApplication) SetAppName(v string) *DataAgentApplication {
	s.AppName = &v
	return s
}

func (s *DataAgentApplication) SetApplicationExtraInfo(v string) *DataAgentApplication {
	s.ApplicationExtraInfo = &v
	return s
}

func (s *DataAgentApplication) SetCreatorName(v string) *DataAgentApplication {
	s.CreatorName = &v
	return s
}

func (s *DataAgentApplication) SetCreatorUid(v string) *DataAgentApplication {
	s.CreatorUid = &v
	return s
}

func (s *DataAgentApplication) SetDescription(v string) *DataAgentApplication {
	s.Description = &v
	return s
}

func (s *DataAgentApplication) SetGmtCreated(v string) *DataAgentApplication {
	s.GmtCreated = &v
	return s
}

func (s *DataAgentApplication) SetGmtModified(v string) *DataAgentApplication {
	s.GmtModified = &v
	return s
}

func (s *DataAgentApplication) SetMainUid(v string) *DataAgentApplication {
	s.MainUid = &v
	return s
}

func (s *DataAgentApplication) SetRegion(v string) *DataAgentApplication {
	s.Region = &v
	return s
}

func (s *DataAgentApplication) SetSessionId(v string) *DataAgentApplication {
	s.SessionId = &v
	return s
}

func (s *DataAgentApplication) SetStatus(v string) *DataAgentApplication {
	s.Status = &v
	return s
}

func (s *DataAgentApplication) SetWorkspaceId(v string) *DataAgentApplication {
	s.WorkspaceId = &v
	return s
}

func (s *DataAgentApplication) Validate() error {
	return dara.Validate(s)
}
