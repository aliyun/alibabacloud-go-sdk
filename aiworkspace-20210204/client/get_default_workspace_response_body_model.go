// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDefaultWorkspaceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConditions(v []*GetDefaultWorkspaceResponseBodyConditions) *GetDefaultWorkspaceResponseBody
	GetConditions() []*GetDefaultWorkspaceResponseBodyConditions
	SetCreator(v string) *GetDefaultWorkspaceResponseBody
	GetCreator() *string
	SetDescription(v string) *GetDefaultWorkspaceResponseBody
	GetDescription() *string
	SetDisplayName(v string) *GetDefaultWorkspaceResponseBody
	GetDisplayName() *string
	SetEnvTypes(v []*string) *GetDefaultWorkspaceResponseBody
	GetEnvTypes() []*string
	SetGmtCreateTime(v string) *GetDefaultWorkspaceResponseBody
	GetGmtCreateTime() *string
	SetGmtModifiedTime(v string) *GetDefaultWorkspaceResponseBody
	GetGmtModifiedTime() *string
	SetOwner(v *GetDefaultWorkspaceResponseBodyOwner) *GetDefaultWorkspaceResponseBody
	GetOwner() *GetDefaultWorkspaceResponseBodyOwner
	SetRequestId(v string) *GetDefaultWorkspaceResponseBody
	GetRequestId() *string
	SetStatus(v string) *GetDefaultWorkspaceResponseBody
	GetStatus() *string
	SetWorkspaceId(v string) *GetDefaultWorkspaceResponseBody
	GetWorkspaceId() *string
	SetWorkspaceName(v string) *GetDefaultWorkspaceResponseBody
	GetWorkspaceName() *string
}

type GetDefaultWorkspaceResponseBody struct {
	// The details about the creation stages of the default workspace.
	Conditions []*GetDefaultWorkspaceResponseBodyConditions `json:"Conditions,omitempty" xml:"Conditions,omitempty" type:"Repeated"`
	// The Alibaba Cloud account that created the workspace.
	//
	// example:
	//
	// 17915******4216
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
	// The environments in the workspace.
	//
	// - A workspace in basic mode contains only the production (prod) environment.
	//
	// - A workspace in standard mode contains the development (dev) and production (prod) environments.
	EnvTypes []*string `json:"EnvTypes,omitempty" xml:"EnvTypes,omitempty" type:"Repeated"`
	// The time when the workspace was created. The time is in Coordinated Universal Time (UTC) and is formatted in ISO 8601.
	//
	// example:
	//
	// 2021-01-21T17:12:35.232Z
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// The time when the workspace was last modified. The time is in UTC and is formatted in ISO 8601.
	//
	// example:
	//
	// 2021-01-21T17:12:35.232Z
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	// The Alibaba Cloud account that created the workspace.
	Owner *GetDefaultWorkspaceResponseBodyOwner `json:"Owner,omitempty" xml:"Owner,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 5A14FA81-DD4E-******-6343FE44B941
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The status of the workspace. Valid values:
	//
	// - ENABLED: Normal.
	//
	// - INITIALIZING: Initializing.
	//
	// - FAILURE: Failed.
	//
	// - DISABLED: Disabled.
	//
	// - FROZEN: Frozen due to an overdue payment.
	//
	// - UPDATING: The workspace is being updated.
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
	// The name of the workspace. The name must be unique within the same region.
	//
	// example:
	//
	// workspace-example
	WorkspaceName *string `json:"WorkspaceName,omitempty" xml:"WorkspaceName,omitempty"`
}

func (s GetDefaultWorkspaceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDefaultWorkspaceResponseBody) GoString() string {
	return s.String()
}

func (s *GetDefaultWorkspaceResponseBody) GetConditions() []*GetDefaultWorkspaceResponseBodyConditions {
	return s.Conditions
}

func (s *GetDefaultWorkspaceResponseBody) GetCreator() *string {
	return s.Creator
}

func (s *GetDefaultWorkspaceResponseBody) GetDescription() *string {
	return s.Description
}

func (s *GetDefaultWorkspaceResponseBody) GetDisplayName() *string {
	return s.DisplayName
}

func (s *GetDefaultWorkspaceResponseBody) GetEnvTypes() []*string {
	return s.EnvTypes
}

func (s *GetDefaultWorkspaceResponseBody) GetGmtCreateTime() *string {
	return s.GmtCreateTime
}

func (s *GetDefaultWorkspaceResponseBody) GetGmtModifiedTime() *string {
	return s.GmtModifiedTime
}

func (s *GetDefaultWorkspaceResponseBody) GetOwner() *GetDefaultWorkspaceResponseBodyOwner {
	return s.Owner
}

func (s *GetDefaultWorkspaceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDefaultWorkspaceResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetDefaultWorkspaceResponseBody) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *GetDefaultWorkspaceResponseBody) GetWorkspaceName() *string {
	return s.WorkspaceName
}

func (s *GetDefaultWorkspaceResponseBody) SetConditions(v []*GetDefaultWorkspaceResponseBodyConditions) *GetDefaultWorkspaceResponseBody {
	s.Conditions = v
	return s
}

func (s *GetDefaultWorkspaceResponseBody) SetCreator(v string) *GetDefaultWorkspaceResponseBody {
	s.Creator = &v
	return s
}

func (s *GetDefaultWorkspaceResponseBody) SetDescription(v string) *GetDefaultWorkspaceResponseBody {
	s.Description = &v
	return s
}

func (s *GetDefaultWorkspaceResponseBody) SetDisplayName(v string) *GetDefaultWorkspaceResponseBody {
	s.DisplayName = &v
	return s
}

func (s *GetDefaultWorkspaceResponseBody) SetEnvTypes(v []*string) *GetDefaultWorkspaceResponseBody {
	s.EnvTypes = v
	return s
}

func (s *GetDefaultWorkspaceResponseBody) SetGmtCreateTime(v string) *GetDefaultWorkspaceResponseBody {
	s.GmtCreateTime = &v
	return s
}

func (s *GetDefaultWorkspaceResponseBody) SetGmtModifiedTime(v string) *GetDefaultWorkspaceResponseBody {
	s.GmtModifiedTime = &v
	return s
}

func (s *GetDefaultWorkspaceResponseBody) SetOwner(v *GetDefaultWorkspaceResponseBodyOwner) *GetDefaultWorkspaceResponseBody {
	s.Owner = v
	return s
}

func (s *GetDefaultWorkspaceResponseBody) SetRequestId(v string) *GetDefaultWorkspaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDefaultWorkspaceResponseBody) SetStatus(v string) *GetDefaultWorkspaceResponseBody {
	s.Status = &v
	return s
}

func (s *GetDefaultWorkspaceResponseBody) SetWorkspaceId(v string) *GetDefaultWorkspaceResponseBody {
	s.WorkspaceId = &v
	return s
}

func (s *GetDefaultWorkspaceResponseBody) SetWorkspaceName(v string) *GetDefaultWorkspaceResponseBody {
	s.WorkspaceName = &v
	return s
}

func (s *GetDefaultWorkspaceResponseBody) Validate() error {
	if s.Conditions != nil {
		for _, item := range s.Conditions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Owner != nil {
		if err := s.Owner.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetDefaultWorkspaceResponseBodyConditions struct {
	// The status code. A value of 200 indicates that the request was successful. Other values indicate that the request failed.
	//
	// example:
	//
	// 200
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error message. This parameter is empty if the status code is 200.
	//
	// example:
	//
	// Create Failed
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The task type. Valid values:
	//
	// - CREATING: Creating.
	//
	// - WORKSPACE_CREATED: The workspace is created.
	//
	// - MEMBERS_ADDED: Members are added.
	//
	// - ENABLED: The process is complete.
	//
	// example:
	//
	// CREATING
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetDefaultWorkspaceResponseBodyConditions) String() string {
	return dara.Prettify(s)
}

func (s GetDefaultWorkspaceResponseBodyConditions) GoString() string {
	return s.String()
}

func (s *GetDefaultWorkspaceResponseBodyConditions) GetCode() *int64 {
	return s.Code
}

func (s *GetDefaultWorkspaceResponseBodyConditions) GetMessage() *string {
	return s.Message
}

func (s *GetDefaultWorkspaceResponseBodyConditions) GetType() *string {
	return s.Type
}

func (s *GetDefaultWorkspaceResponseBodyConditions) SetCode(v int64) *GetDefaultWorkspaceResponseBodyConditions {
	s.Code = &v
	return s
}

func (s *GetDefaultWorkspaceResponseBodyConditions) SetMessage(v string) *GetDefaultWorkspaceResponseBodyConditions {
	s.Message = &v
	return s
}

func (s *GetDefaultWorkspaceResponseBodyConditions) SetType(v string) *GetDefaultWorkspaceResponseBodyConditions {
	s.Type = &v
	return s
}

func (s *GetDefaultWorkspaceResponseBodyConditions) Validate() error {
	return dara.Validate(s)
}

type GetDefaultWorkspaceResponseBodyOwner struct {
	// The UID of the user.
	//
	// example:
	//
	// 17915******4216
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The UID of the user.
	//
	// example:
	//
	// 17915******4216
	UserKp *string `json:"UserKp,omitempty" xml:"UserKp,omitempty"`
	// The username.
	//
	// example:
	//
	// username
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s GetDefaultWorkspaceResponseBodyOwner) String() string {
	return dara.Prettify(s)
}

func (s GetDefaultWorkspaceResponseBodyOwner) GoString() string {
	return s.String()
}

func (s *GetDefaultWorkspaceResponseBodyOwner) GetUserId() *string {
	return s.UserId
}

func (s *GetDefaultWorkspaceResponseBodyOwner) GetUserKp() *string {
	return s.UserKp
}

func (s *GetDefaultWorkspaceResponseBodyOwner) GetUserName() *string {
	return s.UserName
}

func (s *GetDefaultWorkspaceResponseBodyOwner) SetUserId(v string) *GetDefaultWorkspaceResponseBodyOwner {
	s.UserId = &v
	return s
}

func (s *GetDefaultWorkspaceResponseBodyOwner) SetUserKp(v string) *GetDefaultWorkspaceResponseBodyOwner {
	s.UserKp = &v
	return s
}

func (s *GetDefaultWorkspaceResponseBodyOwner) SetUserName(v string) *GetDefaultWorkspaceResponseBodyOwner {
	s.UserName = &v
	return s
}

func (s *GetDefaultWorkspaceResponseBodyOwner) Validate() error {
	return dara.Validate(s)
}
