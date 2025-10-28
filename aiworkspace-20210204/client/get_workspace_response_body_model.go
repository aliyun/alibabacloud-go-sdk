// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWorkspaceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAdminNames(v []*string) *GetWorkspaceResponseBody
	GetAdminNames() []*string
	SetCreator(v string) *GetWorkspaceResponseBody
	GetCreator() *string
	SetDescription(v string) *GetWorkspaceResponseBody
	GetDescription() *string
	SetDisplayName(v string) *GetWorkspaceResponseBody
	GetDisplayName() *string
	SetEnvTypes(v []*string) *GetWorkspaceResponseBody
	GetEnvTypes() []*string
	SetExtraInfos(v map[string]interface{}) *GetWorkspaceResponseBody
	GetExtraInfos() map[string]interface{}
	SetGmtCreateTime(v string) *GetWorkspaceResponseBody
	GetGmtCreateTime() *string
	SetGmtModifiedTime(v string) *GetWorkspaceResponseBody
	GetGmtModifiedTime() *string
	SetIsDefault(v bool) *GetWorkspaceResponseBody
	GetIsDefault() *bool
	SetOwner(v *GetWorkspaceResponseBodyOwner) *GetWorkspaceResponseBody
	GetOwner() *GetWorkspaceResponseBodyOwner
	SetRequestId(v string) *GetWorkspaceResponseBody
	GetRequestId() *string
	SetResourceGroupId(v string) *GetWorkspaceResponseBody
	GetResourceGroupId() *string
	SetStatus(v string) *GetWorkspaceResponseBody
	GetStatus() *string
	SetWorkspaceId(v string) *GetWorkspaceResponseBody
	GetWorkspaceId() *string
	SetWorkspaceName(v string) *GetWorkspaceResponseBody
	GetWorkspaceName() *string
}

type GetWorkspaceResponseBody struct {
	// The names of the administrator accounts.
	AdminNames []*string `json:"AdminNames,omitempty" xml:"AdminNames,omitempty" type:"Repeated"`
	// The ID of the user who creates the workspace.
	//
	// example:
	//
	// 1157******94123
	Creator *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// The description of the workspace.
	//
	// example:
	//
	// workspace description example
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The display name of the workspace.
	//
	// example:
	//
	// workspace-example
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The environment information of the workspace.
	//
	// 	- Workspaces in basic mode can run only in the production environment.
	//
	// 	- Workspaces in standard mode can run in both the development and production environments.
	EnvTypes []*string `json:"EnvTypes,omitempty" xml:"EnvTypes,omitempty" type:"Repeated"`
	// The additional information, which only contains the TenantId field.
	//
	// example:
	//
	// {"TenantId": "4286******98"}
	ExtraInfos map[string]interface{} `json:"ExtraInfos,omitempty" xml:"ExtraInfos,omitempty"`
	// The time when the workspace is created, in UTC. The time follows the ISO 8601 standard.
	//
	// example:
	//
	// 2021-01-21T17:12:35.232Z
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// The time when the workspace is modified, in UTC. The time follows the ISO 8601 standard.
	//
	// example:
	//
	// 2021-01-21T17:12:35.232Z
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	// Indicates whether the workspace is the default workspace. Valid values:
	//
	// 	- false
	//
	// 	- true
	//
	// example:
	//
	// true
	IsDefault *bool `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	// The information about the workspace owner. This parameter is valid only when Verbose is set to true.
	Owner *GetWorkspaceResponseBodyOwner `json:"Owner,omitempty" xml:"Owner,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// A0F049F0-8D69-5BAC-8F10-B4DED1B5A34C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfmwp7rkyq****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The workspace state. Valid values:
	//
	// 	- ENABLED
	//
	// 	- INITIALIZING
	//
	// 	- FAILURE:
	//
	// 	- DISABLED
	//
	// 	- FROZEN
	//
	// 	- UPDATING
	//
	// example:
	//
	// ENABLED
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The workspace ID.
	//
	// example:
	//
	// 1234
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
	// The name of the workspace.
	//
	// example:
	//
	// workspace-example
	WorkspaceName *string `json:"WorkspaceName,omitempty" xml:"WorkspaceName,omitempty"`
}

func (s GetWorkspaceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetWorkspaceResponseBody) GoString() string {
	return s.String()
}

func (s *GetWorkspaceResponseBody) GetAdminNames() []*string {
	return s.AdminNames
}

func (s *GetWorkspaceResponseBody) GetCreator() *string {
	return s.Creator
}

func (s *GetWorkspaceResponseBody) GetDescription() *string {
	return s.Description
}

func (s *GetWorkspaceResponseBody) GetDisplayName() *string {
	return s.DisplayName
}

func (s *GetWorkspaceResponseBody) GetEnvTypes() []*string {
	return s.EnvTypes
}

func (s *GetWorkspaceResponseBody) GetExtraInfos() map[string]interface{} {
	return s.ExtraInfos
}

func (s *GetWorkspaceResponseBody) GetGmtCreateTime() *string {
	return s.GmtCreateTime
}

func (s *GetWorkspaceResponseBody) GetGmtModifiedTime() *string {
	return s.GmtModifiedTime
}

func (s *GetWorkspaceResponseBody) GetIsDefault() *bool {
	return s.IsDefault
}

func (s *GetWorkspaceResponseBody) GetOwner() *GetWorkspaceResponseBodyOwner {
	return s.Owner
}

func (s *GetWorkspaceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetWorkspaceResponseBody) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *GetWorkspaceResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetWorkspaceResponseBody) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *GetWorkspaceResponseBody) GetWorkspaceName() *string {
	return s.WorkspaceName
}

func (s *GetWorkspaceResponseBody) SetAdminNames(v []*string) *GetWorkspaceResponseBody {
	s.AdminNames = v
	return s
}

func (s *GetWorkspaceResponseBody) SetCreator(v string) *GetWorkspaceResponseBody {
	s.Creator = &v
	return s
}

func (s *GetWorkspaceResponseBody) SetDescription(v string) *GetWorkspaceResponseBody {
	s.Description = &v
	return s
}

func (s *GetWorkspaceResponseBody) SetDisplayName(v string) *GetWorkspaceResponseBody {
	s.DisplayName = &v
	return s
}

func (s *GetWorkspaceResponseBody) SetEnvTypes(v []*string) *GetWorkspaceResponseBody {
	s.EnvTypes = v
	return s
}

func (s *GetWorkspaceResponseBody) SetExtraInfos(v map[string]interface{}) *GetWorkspaceResponseBody {
	s.ExtraInfos = v
	return s
}

func (s *GetWorkspaceResponseBody) SetGmtCreateTime(v string) *GetWorkspaceResponseBody {
	s.GmtCreateTime = &v
	return s
}

func (s *GetWorkspaceResponseBody) SetGmtModifiedTime(v string) *GetWorkspaceResponseBody {
	s.GmtModifiedTime = &v
	return s
}

func (s *GetWorkspaceResponseBody) SetIsDefault(v bool) *GetWorkspaceResponseBody {
	s.IsDefault = &v
	return s
}

func (s *GetWorkspaceResponseBody) SetOwner(v *GetWorkspaceResponseBodyOwner) *GetWorkspaceResponseBody {
	s.Owner = v
	return s
}

func (s *GetWorkspaceResponseBody) SetRequestId(v string) *GetWorkspaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetWorkspaceResponseBody) SetResourceGroupId(v string) *GetWorkspaceResponseBody {
	s.ResourceGroupId = &v
	return s
}

func (s *GetWorkspaceResponseBody) SetStatus(v string) *GetWorkspaceResponseBody {
	s.Status = &v
	return s
}

func (s *GetWorkspaceResponseBody) SetWorkspaceId(v string) *GetWorkspaceResponseBody {
	s.WorkspaceId = &v
	return s
}

func (s *GetWorkspaceResponseBody) SetWorkspaceName(v string) *GetWorkspaceResponseBody {
	s.WorkspaceName = &v
	return s
}

func (s *GetWorkspaceResponseBody) Validate() error {
	if s.Owner != nil {
		if err := s.Owner.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetWorkspaceResponseBodyOwner struct {
	// The display name.
	//
	// example:
	//
	// mings****t
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The user ID.
	//
	// example:
	//
	// 1157******94123
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The user ID.
	//
	// example:
	//
	// 1157******94123
	UserKp *string `json:"UserKp,omitempty" xml:"UserKp,omitempty"`
	// The username.
	//
	// example:
	//
	// mings****t
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s GetWorkspaceResponseBodyOwner) String() string {
	return dara.Prettify(s)
}

func (s GetWorkspaceResponseBodyOwner) GoString() string {
	return s.String()
}

func (s *GetWorkspaceResponseBodyOwner) GetDisplayName() *string {
	return s.DisplayName
}

func (s *GetWorkspaceResponseBodyOwner) GetUserId() *string {
	return s.UserId
}

func (s *GetWorkspaceResponseBodyOwner) GetUserKp() *string {
	return s.UserKp
}

func (s *GetWorkspaceResponseBodyOwner) GetUserName() *string {
	return s.UserName
}

func (s *GetWorkspaceResponseBodyOwner) SetDisplayName(v string) *GetWorkspaceResponseBodyOwner {
	s.DisplayName = &v
	return s
}

func (s *GetWorkspaceResponseBodyOwner) SetUserId(v string) *GetWorkspaceResponseBodyOwner {
	s.UserId = &v
	return s
}

func (s *GetWorkspaceResponseBodyOwner) SetUserKp(v string) *GetWorkspaceResponseBodyOwner {
	s.UserKp = &v
	return s
}

func (s *GetWorkspaceResponseBodyOwner) SetUserName(v string) *GetWorkspaceResponseBodyOwner {
	s.UserName = &v
	return s
}

func (s *GetWorkspaceResponseBodyOwner) Validate() error {
	return dara.Validate(s)
}
