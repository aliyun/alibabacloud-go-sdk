// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateContextDatabaseWorkspaceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCapabilities(v *UpdateContextDatabaseWorkspaceResponseBodyCapabilities) *UpdateContextDatabaseWorkspaceResponseBody
	GetCapabilities() *UpdateContextDatabaseWorkspaceResponseBodyCapabilities
	SetClusterId(v string) *UpdateContextDatabaseWorkspaceResponseBody
	GetClusterId() *string
	SetConsolePermissionMode(v string) *UpdateContextDatabaseWorkspaceResponseBody
	GetConsolePermissionMode() *string
	SetCreatedAt(v string) *UpdateContextDatabaseWorkspaceResponseBody
	GetCreatedAt() *string
	SetPermissionMode(v string) *UpdateContextDatabaseWorkspaceResponseBody
	GetPermissionMode() *string
	SetRequestId(v string) *UpdateContextDatabaseWorkspaceResponseBody
	GetRequestId() *string
	SetStatus(v string) *UpdateContextDatabaseWorkspaceResponseBody
	GetStatus() *string
	SetType(v string) *UpdateContextDatabaseWorkspaceResponseBody
	GetType() *string
	SetViewer(v *UpdateContextDatabaseWorkspaceResponseBodyViewer) *UpdateContextDatabaseWorkspaceResponseBody
	GetViewer() *UpdateContextDatabaseWorkspaceResponseBodyViewer
	SetWorkspaceId(v string) *UpdateContextDatabaseWorkspaceResponseBody
	GetWorkspaceId() *string
	SetWorkspaceName(v string) *UpdateContextDatabaseWorkspaceResponseBody
	GetWorkspaceName() *string
}

type UpdateContextDatabaseWorkspaceResponseBody struct {
	// The account-level capability information of the current request. It only indicates the operations that the account compatibility adapter can perform and does not represent resource-level authorization. This field is optional and is not returned if no account-level executable operations are provided.
	Capabilities *UpdateContextDatabaseWorkspaceResponseBodyCapabilities `json:"Capabilities,omitempty" xml:"Capabilities,omitempty" type:"Struct"`
	// The ID of the cluster to which the workspace belongs. For shared clusters, `default` is returned. For dedicated clusters, the corresponding cluster ID is returned.
	//
	// example:
	//
	// default
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The permission mode used by the server to parse the identity of the current request. Valid values:
	//
	// - `personal_compatibility`: compatibility mode.
	//
	// - `advanced`: advanced permission mode.
	//
	// This field describes the identity context of the current request. For the actual permission mode of the workspace, refer to PermissionMode.
	//
	// example:
	//
	// advanced
	ConsolePermissionMode *string `json:"ConsolePermissionMode,omitempty" xml:"ConsolePermissionMode,omitempty"`
	// The creation time in ISO 8601 format.
	//
	// example:
	//
	// 2026-05-28T17:59:55Z
	CreatedAt *string `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	// The permission mode of the workspace. Valid values:
	//
	// - `personal_compatibility`: compatibility mode.
	//
	// - `advanced`: advanced permission mode.
	//
	// example:
	//
	// advanced
	PermissionMode *string `json:"PermissionMode,omitempty" xml:"PermissionMode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// FE9C65D7-930F-57A5-A207-8C396329****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The workspace status.
	//
	// example:
	//
	// Active
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The workspace type.
	//
	// example:
	//
	// personal
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The viewer identity information of the current request, including the identity resolution mode, caller account type, and member information. Member-related fields are optional.
	Viewer *UpdateContextDatabaseWorkspaceResponseBodyViewer `json:"Viewer,omitempty" xml:"Viewer,omitempty" type:"Struct"`
	// The workspace ID.
	//
	// example:
	//
	// ws-as1llqmkol****
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
	// The updated name of the workspace.
	//
	// example:
	//
	// my-workspace
	WorkspaceName *string `json:"WorkspaceName,omitempty" xml:"WorkspaceName,omitempty"`
}

func (s UpdateContextDatabaseWorkspaceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateContextDatabaseWorkspaceResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateContextDatabaseWorkspaceResponseBody) GetCapabilities() *UpdateContextDatabaseWorkspaceResponseBodyCapabilities {
	return s.Capabilities
}

func (s *UpdateContextDatabaseWorkspaceResponseBody) GetClusterId() *string {
	return s.ClusterId
}

func (s *UpdateContextDatabaseWorkspaceResponseBody) GetConsolePermissionMode() *string {
	return s.ConsolePermissionMode
}

func (s *UpdateContextDatabaseWorkspaceResponseBody) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *UpdateContextDatabaseWorkspaceResponseBody) GetPermissionMode() *string {
	return s.PermissionMode
}

func (s *UpdateContextDatabaseWorkspaceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateContextDatabaseWorkspaceResponseBody) GetStatus() *string {
	return s.Status
}

func (s *UpdateContextDatabaseWorkspaceResponseBody) GetType() *string {
	return s.Type
}

func (s *UpdateContextDatabaseWorkspaceResponseBody) GetViewer() *UpdateContextDatabaseWorkspaceResponseBodyViewer {
	return s.Viewer
}

func (s *UpdateContextDatabaseWorkspaceResponseBody) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *UpdateContextDatabaseWorkspaceResponseBody) GetWorkspaceName() *string {
	return s.WorkspaceName
}

func (s *UpdateContextDatabaseWorkspaceResponseBody) SetCapabilities(v *UpdateContextDatabaseWorkspaceResponseBodyCapabilities) *UpdateContextDatabaseWorkspaceResponseBody {
	s.Capabilities = v
	return s
}

func (s *UpdateContextDatabaseWorkspaceResponseBody) SetClusterId(v string) *UpdateContextDatabaseWorkspaceResponseBody {
	s.ClusterId = &v
	return s
}

func (s *UpdateContextDatabaseWorkspaceResponseBody) SetConsolePermissionMode(v string) *UpdateContextDatabaseWorkspaceResponseBody {
	s.ConsolePermissionMode = &v
	return s
}

func (s *UpdateContextDatabaseWorkspaceResponseBody) SetCreatedAt(v string) *UpdateContextDatabaseWorkspaceResponseBody {
	s.CreatedAt = &v
	return s
}

func (s *UpdateContextDatabaseWorkspaceResponseBody) SetPermissionMode(v string) *UpdateContextDatabaseWorkspaceResponseBody {
	s.PermissionMode = &v
	return s
}

func (s *UpdateContextDatabaseWorkspaceResponseBody) SetRequestId(v string) *UpdateContextDatabaseWorkspaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateContextDatabaseWorkspaceResponseBody) SetStatus(v string) *UpdateContextDatabaseWorkspaceResponseBody {
	s.Status = &v
	return s
}

func (s *UpdateContextDatabaseWorkspaceResponseBody) SetType(v string) *UpdateContextDatabaseWorkspaceResponseBody {
	s.Type = &v
	return s
}

func (s *UpdateContextDatabaseWorkspaceResponseBody) SetViewer(v *UpdateContextDatabaseWorkspaceResponseBodyViewer) *UpdateContextDatabaseWorkspaceResponseBody {
	s.Viewer = v
	return s
}

func (s *UpdateContextDatabaseWorkspaceResponseBody) SetWorkspaceId(v string) *UpdateContextDatabaseWorkspaceResponseBody {
	s.WorkspaceId = &v
	return s
}

func (s *UpdateContextDatabaseWorkspaceResponseBody) SetWorkspaceName(v string) *UpdateContextDatabaseWorkspaceResponseBody {
	s.WorkspaceName = &v
	return s
}

func (s *UpdateContextDatabaseWorkspaceResponseBody) Validate() error {
	if s.Capabilities != nil {
		if err := s.Capabilities.Validate(); err != nil {
			return err
		}
	}
	if s.Viewer != nil {
		if err := s.Viewer.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateContextDatabaseWorkspaceResponseBodyCapabilities struct {
	// The list of operation names that the account compatibility adapter can perform. The list content does not represent read/write, member group, or access control authorization for specific resources.
	ExecutableActions []*string `json:"ExecutableActions,omitempty" xml:"ExecutableActions,omitempty" type:"Repeated"`
}

func (s UpdateContextDatabaseWorkspaceResponseBodyCapabilities) String() string {
	return dara.Prettify(s)
}

func (s UpdateContextDatabaseWorkspaceResponseBodyCapabilities) GoString() string {
	return s.String()
}

func (s *UpdateContextDatabaseWorkspaceResponseBodyCapabilities) GetExecutableActions() []*string {
	return s.ExecutableActions
}

func (s *UpdateContextDatabaseWorkspaceResponseBodyCapabilities) SetExecutableActions(v []*string) *UpdateContextDatabaseWorkspaceResponseBodyCapabilities {
	s.ExecutableActions = v
	return s
}

func (s *UpdateContextDatabaseWorkspaceResponseBodyCapabilities) Validate() error {
	return dara.Validate(s)
}

type UpdateContextDatabaseWorkspaceResponseBodyViewer struct {
	// The ID of the bound workspace member in the identity information of the current request. Optional. This field is empty if no bound member information is provided.
	//
	// example:
	//
	// member-example
	BoundMemberId *string `json:"BoundMemberId,omitempty" xml:"BoundMemberId,omitempty"`
	// The role of the bound member in the workspace, such as `Owner`. Optional. This field is empty if no bound member role information is provided.
	//
	// example:
	//
	// Owner
	BoundMemberRole *string `json:"BoundMemberRole,omitempty" xml:"BoundMemberRole,omitempty"`
	// The status of the bound member, such as `active`. Optional. This field is empty if no bound member status information is provided.
	//
	// example:
	//
	// active
	BoundMemberStatus *string `json:"BoundMemberStatus,omitempty" xml:"BoundMemberStatus,omitempty"`
	// The account type of the caller for the current request, such as `customer`. This field describes the caller account and does not indicate the role of a workspace member.
	//
	// example:
	//
	// customer
	CallerType *string `json:"CallerType,omitempty" xml:"CallerType,omitempty"`
	// The identity resolution mode of the current request. Valid values:
	//
	// - `account_compatibility`: account compatibility identity mode.
	//
	// - `bound_member`: bound member identity mode.
	//
	// The identity resolution mode and the workspace permission mode are different concepts. You cannot infer whether advanced permissions are enabled for the workspace based on this field.
	//
	// example:
	//
	// account_compatibility
	IdentityMode *string `json:"IdentityMode,omitempty" xml:"IdentityMode,omitempty"`
	// The ID of the selected workspace member in the account compatibility identity mode. Optional. This field is empty if no member is selected for the current request. This field and the bound member ID represent the member selection result and the binding relationship, respectively.
	//
	// example:
	//
	// member-example
	SelectedMemberId *string `json:"SelectedMemberId,omitempty" xml:"SelectedMemberId,omitempty"`
}

func (s UpdateContextDatabaseWorkspaceResponseBodyViewer) String() string {
	return dara.Prettify(s)
}

func (s UpdateContextDatabaseWorkspaceResponseBodyViewer) GoString() string {
	return s.String()
}

func (s *UpdateContextDatabaseWorkspaceResponseBodyViewer) GetBoundMemberId() *string {
	return s.BoundMemberId
}

func (s *UpdateContextDatabaseWorkspaceResponseBodyViewer) GetBoundMemberRole() *string {
	return s.BoundMemberRole
}

func (s *UpdateContextDatabaseWorkspaceResponseBodyViewer) GetBoundMemberStatus() *string {
	return s.BoundMemberStatus
}

func (s *UpdateContextDatabaseWorkspaceResponseBodyViewer) GetCallerType() *string {
	return s.CallerType
}

func (s *UpdateContextDatabaseWorkspaceResponseBodyViewer) GetIdentityMode() *string {
	return s.IdentityMode
}

func (s *UpdateContextDatabaseWorkspaceResponseBodyViewer) GetSelectedMemberId() *string {
	return s.SelectedMemberId
}

func (s *UpdateContextDatabaseWorkspaceResponseBodyViewer) SetBoundMemberId(v string) *UpdateContextDatabaseWorkspaceResponseBodyViewer {
	s.BoundMemberId = &v
	return s
}

func (s *UpdateContextDatabaseWorkspaceResponseBodyViewer) SetBoundMemberRole(v string) *UpdateContextDatabaseWorkspaceResponseBodyViewer {
	s.BoundMemberRole = &v
	return s
}

func (s *UpdateContextDatabaseWorkspaceResponseBodyViewer) SetBoundMemberStatus(v string) *UpdateContextDatabaseWorkspaceResponseBodyViewer {
	s.BoundMemberStatus = &v
	return s
}

func (s *UpdateContextDatabaseWorkspaceResponseBodyViewer) SetCallerType(v string) *UpdateContextDatabaseWorkspaceResponseBodyViewer {
	s.CallerType = &v
	return s
}

func (s *UpdateContextDatabaseWorkspaceResponseBodyViewer) SetIdentityMode(v string) *UpdateContextDatabaseWorkspaceResponseBodyViewer {
	s.IdentityMode = &v
	return s
}

func (s *UpdateContextDatabaseWorkspaceResponseBodyViewer) SetSelectedMemberId(v string) *UpdateContextDatabaseWorkspaceResponseBodyViewer {
	s.SelectedMemberId = &v
	return s
}

func (s *UpdateContextDatabaseWorkspaceResponseBodyViewer) Validate() error {
	return dara.Validate(s)
}
