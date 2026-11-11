// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteContextDatabaseWorkspaceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCapabilities(v *DeleteContextDatabaseWorkspaceResponseBodyCapabilities) *DeleteContextDatabaseWorkspaceResponseBody
	GetCapabilities() *DeleteContextDatabaseWorkspaceResponseBodyCapabilities
	SetClusterId(v string) *DeleteContextDatabaseWorkspaceResponseBody
	GetClusterId() *string
	SetConsolePermissionMode(v string) *DeleteContextDatabaseWorkspaceResponseBody
	GetConsolePermissionMode() *string
	SetCreatedAt(v string) *DeleteContextDatabaseWorkspaceResponseBody
	GetCreatedAt() *string
	SetPermissionMode(v string) *DeleteContextDatabaseWorkspaceResponseBody
	GetPermissionMode() *string
	SetRequestId(v string) *DeleteContextDatabaseWorkspaceResponseBody
	GetRequestId() *string
	SetStatus(v string) *DeleteContextDatabaseWorkspaceResponseBody
	GetStatus() *string
	SetType(v string) *DeleteContextDatabaseWorkspaceResponseBody
	GetType() *string
	SetViewer(v *DeleteContextDatabaseWorkspaceResponseBodyViewer) *DeleteContextDatabaseWorkspaceResponseBody
	GetViewer() *DeleteContextDatabaseWorkspaceResponseBodyViewer
	SetWorkspaceId(v string) *DeleteContextDatabaseWorkspaceResponseBody
	GetWorkspaceId() *string
	SetWorkspaceName(v string) *DeleteContextDatabaseWorkspaceResponseBody
	GetWorkspaceName() *string
}

type DeleteContextDatabaseWorkspaceResponseBody struct {
	// The account-level capability information for the current request. This information only indicates the operations that the account compatibility adapter can perform and does not represent resource-level authorization. This field is optional and is not returned when no account-level executable operations are provided.
	Capabilities *DeleteContextDatabaseWorkspaceResponseBodyCapabilities `json:"Capabilities,omitempty" xml:"Capabilities,omitempty" type:"Struct"`
	// The ID of the cluster to which the workspace belongs. For shared clusters, default is returned. For dedicated clusters, the corresponding cluster ID is returned.
	//
	// example:
	//
	// default
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The permission mode used by the server to parse the identity of the current request. Valid values:
	//
	// - personal_compatibility: Compatibility mode.
	//
	// - advanced: Advanced permission mode.
	//
	// This field describes the identity context of the current request. For the actual permission mode of the workspace, refer to PermissionMode.
	//
	// example:
	//
	// advanced
	ConsolePermissionMode *string `json:"ConsolePermissionMode,omitempty" xml:"ConsolePermissionMode,omitempty"`
	// The creation time in ISO-8601 format.
	//
	// example:
	//
	// 2026-05-28T17:59:55Z
	CreatedAt *string `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	// The permission mode of the workspace. Valid values:
	//
	// - personal_compatibility: Compatibility mode.
	//
	// - advanced: Advanced permission mode.
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
	// The status of the workspace. This value is fixed to Deleted.
	//
	// example:
	//
	// Deleted
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The type of the workspace.
	//
	// example:
	//
	// personal
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The viewer identity information for the current request, including the identity resolution mode, caller account type, and member information. Member-related fields are optional.
	Viewer *DeleteContextDatabaseWorkspaceResponseBodyViewer `json:"Viewer,omitempty" xml:"Viewer,omitempty" type:"Struct"`
	// The ID of the deleted workspace.
	//
	// example:
	//
	// ws-as1llqmkol****
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
	// The name of the deleted workspace.
	//
	// example:
	//
	// my-workspace
	WorkspaceName *string `json:"WorkspaceName,omitempty" xml:"WorkspaceName,omitempty"`
}

func (s DeleteContextDatabaseWorkspaceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteContextDatabaseWorkspaceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteContextDatabaseWorkspaceResponseBody) GetCapabilities() *DeleteContextDatabaseWorkspaceResponseBodyCapabilities {
	return s.Capabilities
}

func (s *DeleteContextDatabaseWorkspaceResponseBody) GetClusterId() *string {
	return s.ClusterId
}

func (s *DeleteContextDatabaseWorkspaceResponseBody) GetConsolePermissionMode() *string {
	return s.ConsolePermissionMode
}

func (s *DeleteContextDatabaseWorkspaceResponseBody) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *DeleteContextDatabaseWorkspaceResponseBody) GetPermissionMode() *string {
	return s.PermissionMode
}

func (s *DeleteContextDatabaseWorkspaceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteContextDatabaseWorkspaceResponseBody) GetStatus() *string {
	return s.Status
}

func (s *DeleteContextDatabaseWorkspaceResponseBody) GetType() *string {
	return s.Type
}

func (s *DeleteContextDatabaseWorkspaceResponseBody) GetViewer() *DeleteContextDatabaseWorkspaceResponseBodyViewer {
	return s.Viewer
}

func (s *DeleteContextDatabaseWorkspaceResponseBody) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *DeleteContextDatabaseWorkspaceResponseBody) GetWorkspaceName() *string {
	return s.WorkspaceName
}

func (s *DeleteContextDatabaseWorkspaceResponseBody) SetCapabilities(v *DeleteContextDatabaseWorkspaceResponseBodyCapabilities) *DeleteContextDatabaseWorkspaceResponseBody {
	s.Capabilities = v
	return s
}

func (s *DeleteContextDatabaseWorkspaceResponseBody) SetClusterId(v string) *DeleteContextDatabaseWorkspaceResponseBody {
	s.ClusterId = &v
	return s
}

func (s *DeleteContextDatabaseWorkspaceResponseBody) SetConsolePermissionMode(v string) *DeleteContextDatabaseWorkspaceResponseBody {
	s.ConsolePermissionMode = &v
	return s
}

func (s *DeleteContextDatabaseWorkspaceResponseBody) SetCreatedAt(v string) *DeleteContextDatabaseWorkspaceResponseBody {
	s.CreatedAt = &v
	return s
}

func (s *DeleteContextDatabaseWorkspaceResponseBody) SetPermissionMode(v string) *DeleteContextDatabaseWorkspaceResponseBody {
	s.PermissionMode = &v
	return s
}

func (s *DeleteContextDatabaseWorkspaceResponseBody) SetRequestId(v string) *DeleteContextDatabaseWorkspaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteContextDatabaseWorkspaceResponseBody) SetStatus(v string) *DeleteContextDatabaseWorkspaceResponseBody {
	s.Status = &v
	return s
}

func (s *DeleteContextDatabaseWorkspaceResponseBody) SetType(v string) *DeleteContextDatabaseWorkspaceResponseBody {
	s.Type = &v
	return s
}

func (s *DeleteContextDatabaseWorkspaceResponseBody) SetViewer(v *DeleteContextDatabaseWorkspaceResponseBodyViewer) *DeleteContextDatabaseWorkspaceResponseBody {
	s.Viewer = v
	return s
}

func (s *DeleteContextDatabaseWorkspaceResponseBody) SetWorkspaceId(v string) *DeleteContextDatabaseWorkspaceResponseBody {
	s.WorkspaceId = &v
	return s
}

func (s *DeleteContextDatabaseWorkspaceResponseBody) SetWorkspaceName(v string) *DeleteContextDatabaseWorkspaceResponseBody {
	s.WorkspaceName = &v
	return s
}

func (s *DeleteContextDatabaseWorkspaceResponseBody) Validate() error {
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

type DeleteContextDatabaseWorkspaceResponseBodyCapabilities struct {
	// The list of operation names that the account compatibility adapter can perform. The list content does not represent read/write, member group, or access control authorization for specific resources.
	ExecutableActions []*string `json:"ExecutableActions,omitempty" xml:"ExecutableActions,omitempty" type:"Repeated"`
}

func (s DeleteContextDatabaseWorkspaceResponseBodyCapabilities) String() string {
	return dara.Prettify(s)
}

func (s DeleteContextDatabaseWorkspaceResponseBodyCapabilities) GoString() string {
	return s.String()
}

func (s *DeleteContextDatabaseWorkspaceResponseBodyCapabilities) GetExecutableActions() []*string {
	return s.ExecutableActions
}

func (s *DeleteContextDatabaseWorkspaceResponseBodyCapabilities) SetExecutableActions(v []*string) *DeleteContextDatabaseWorkspaceResponseBodyCapabilities {
	s.ExecutableActions = v
	return s
}

func (s *DeleteContextDatabaseWorkspaceResponseBodyCapabilities) Validate() error {
	return dara.Validate(s)
}

type DeleteContextDatabaseWorkspaceResponseBodyViewer struct {
	// The ID of the bound workspace member in the identity information for the current request. This field is optional and has no value when bound member information is not provided.
	//
	// example:
	//
	// member-example
	BoundMemberId *string `json:"BoundMemberId,omitempty" xml:"BoundMemberId,omitempty"`
	// The role of the bound member in the workspace, such as Owner. This field is optional and has no value when bound member role information is not provided.
	//
	// example:
	//
	// Owner
	BoundMemberRole *string `json:"BoundMemberRole,omitempty" xml:"BoundMemberRole,omitempty"`
	// The status of the bound member, such as active. This field is optional and has no value when bound member status information is not provided.
	//
	// example:
	//
	// active
	BoundMemberStatus *string `json:"BoundMemberStatus,omitempty" xml:"BoundMemberStatus,omitempty"`
	// The account type of the caller for the current request, such as customer. This field describes the caller account and does not represent the role of a workspace member.
	//
	// example:
	//
	// customer
	CallerType *string `json:"CallerType,omitempty" xml:"CallerType,omitempty"`
	// The identity resolution mode for the current request. Valid values:
	//
	// - account_compatibility: Account compatibility identity mode.
	//
	// - bound_member: Bound member identity mode.
	//
	// The identity resolution mode and the workspace permission mode are different concepts. You cannot infer whether advanced permissions are enabled for the workspace based on this field.
	//
	// example:
	//
	// account_compatibility
	IdentityMode *string `json:"IdentityMode,omitempty" xml:"IdentityMode,omitempty"`
	// The ID of the selected workspace member in the account compatibility identity mode. This field is optional and has no value when no member is selected for the current request. This field and the bound member ID represent the member selection result and the binding relationship, respectively.
	//
	// example:
	//
	// member-example
	SelectedMemberId *string `json:"SelectedMemberId,omitempty" xml:"SelectedMemberId,omitempty"`
}

func (s DeleteContextDatabaseWorkspaceResponseBodyViewer) String() string {
	return dara.Prettify(s)
}

func (s DeleteContextDatabaseWorkspaceResponseBodyViewer) GoString() string {
	return s.String()
}

func (s *DeleteContextDatabaseWorkspaceResponseBodyViewer) GetBoundMemberId() *string {
	return s.BoundMemberId
}

func (s *DeleteContextDatabaseWorkspaceResponseBodyViewer) GetBoundMemberRole() *string {
	return s.BoundMemberRole
}

func (s *DeleteContextDatabaseWorkspaceResponseBodyViewer) GetBoundMemberStatus() *string {
	return s.BoundMemberStatus
}

func (s *DeleteContextDatabaseWorkspaceResponseBodyViewer) GetCallerType() *string {
	return s.CallerType
}

func (s *DeleteContextDatabaseWorkspaceResponseBodyViewer) GetIdentityMode() *string {
	return s.IdentityMode
}

func (s *DeleteContextDatabaseWorkspaceResponseBodyViewer) GetSelectedMemberId() *string {
	return s.SelectedMemberId
}

func (s *DeleteContextDatabaseWorkspaceResponseBodyViewer) SetBoundMemberId(v string) *DeleteContextDatabaseWorkspaceResponseBodyViewer {
	s.BoundMemberId = &v
	return s
}

func (s *DeleteContextDatabaseWorkspaceResponseBodyViewer) SetBoundMemberRole(v string) *DeleteContextDatabaseWorkspaceResponseBodyViewer {
	s.BoundMemberRole = &v
	return s
}

func (s *DeleteContextDatabaseWorkspaceResponseBodyViewer) SetBoundMemberStatus(v string) *DeleteContextDatabaseWorkspaceResponseBodyViewer {
	s.BoundMemberStatus = &v
	return s
}

func (s *DeleteContextDatabaseWorkspaceResponseBodyViewer) SetCallerType(v string) *DeleteContextDatabaseWorkspaceResponseBodyViewer {
	s.CallerType = &v
	return s
}

func (s *DeleteContextDatabaseWorkspaceResponseBodyViewer) SetIdentityMode(v string) *DeleteContextDatabaseWorkspaceResponseBodyViewer {
	s.IdentityMode = &v
	return s
}

func (s *DeleteContextDatabaseWorkspaceResponseBodyViewer) SetSelectedMemberId(v string) *DeleteContextDatabaseWorkspaceResponseBodyViewer {
	s.SelectedMemberId = &v
	return s
}

func (s *DeleteContextDatabaseWorkspaceResponseBodyViewer) Validate() error {
	return dara.Validate(s)
}
