// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateContextDatabaseWorkspaceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApiKey(v string) *CreateContextDatabaseWorkspaceResponseBody
	GetApiKey() *string
	SetApiKeyName(v string) *CreateContextDatabaseWorkspaceResponseBody
	GetApiKeyName() *string
	SetCapabilities(v *CreateContextDatabaseWorkspaceResponseBodyCapabilities) *CreateContextDatabaseWorkspaceResponseBody
	GetCapabilities() *CreateContextDatabaseWorkspaceResponseBodyCapabilities
	SetClusterId(v string) *CreateContextDatabaseWorkspaceResponseBody
	GetClusterId() *string
	SetConsolePermissionMode(v string) *CreateContextDatabaseWorkspaceResponseBody
	GetConsolePermissionMode() *string
	SetCreatedAt(v string) *CreateContextDatabaseWorkspaceResponseBody
	GetCreatedAt() *string
	SetMemberId(v string) *CreateContextDatabaseWorkspaceResponseBody
	GetMemberId() *string
	SetMemberName(v string) *CreateContextDatabaseWorkspaceResponseBody
	GetMemberName() *string
	SetPermissionMode(v string) *CreateContextDatabaseWorkspaceResponseBody
	GetPermissionMode() *string
	SetRequestId(v string) *CreateContextDatabaseWorkspaceResponseBody
	GetRequestId() *string
	SetRole(v string) *CreateContextDatabaseWorkspaceResponseBody
	GetRole() *string
	SetStatus(v string) *CreateContextDatabaseWorkspaceResponseBody
	GetStatus() *string
	SetType(v string) *CreateContextDatabaseWorkspaceResponseBody
	GetType() *string
	SetViewer(v *CreateContextDatabaseWorkspaceResponseBodyViewer) *CreateContextDatabaseWorkspaceResponseBody
	GetViewer() *CreateContextDatabaseWorkspaceResponseBodyViewer
	SetWorkspaceId(v string) *CreateContextDatabaseWorkspaceResponseBody
	GetWorkspaceId() *string
	SetWorkspaceName(v string) *CreateContextDatabaseWorkspaceResponseBody
	GetWorkspaceName() *string
}

type CreateContextDatabaseWorkspaceResponseBody struct {
	// The plaintext of the first API key. It is returned only once during creation. Save the key properly. If GenerateInitialKey is set to `false`, this field is empty.
	//
	// example:
	//
	// ctxdb-*****
	ApiKey *string `json:"ApiKey,omitempty" xml:"ApiKey,omitempty"`
	// The name of the first API key. The name is `default` when the key is generated. If GenerateInitialKey is set to `false`, this field is empty.
	//
	// example:
	//
	// default
	ApiKeyName *string `json:"ApiKeyName,omitempty" xml:"ApiKeyName,omitempty"`
	// The account-level capability information for the current call. It only indicates the operations that the account compatibility adapter can perform and does not represent resource-level authorization. This field is optional and is not returned if no account-level executable operations are provided.
	Capabilities *CreateContextDatabaseWorkspaceResponseBodyCapabilities `json:"Capabilities,omitempty" xml:"Capabilities,omitempty" type:"Struct"`
	// The ID of the cluster to which the workspace belongs. `default` is returned for a shared cluster, and the corresponding cluster ID is returned for a dedicated cluster.
	//
	// example:
	//
	// default
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The permission mode used by the server to parse the identity of the current call. Valid values:
	//
	// - `personal_compatibility`: The compatibility mode.
	//
	// - `advanced`: The advanced permission mode.
	//
	// This field describes the identity context of the current call. The actual permission mode of the workspace is subject to PermissionMode. Creating a workspace uses the account-level identity context, so this field may be empty.
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
	// The ID of the first member.
	//
	// example:
	//
	// mb-cz51tnnp8****
	MemberId *string `json:"MemberId,omitempty" xml:"MemberId,omitempty"`
	// The name of the first member.
	//
	// example:
	//
	// my-member
	MemberName *string `json:"MemberName,omitempty" xml:"MemberName,omitempty"`
	// The permission mode that actually takes effect after the workspace is created. A lowercase string is returned.
	//
	// - `personal_compatibility`: The compatibility mode, which corresponds to the request value `compatibility`.
	//
	// - `advanced`: The advanced permission mode.
	//
	// This field is optional. If it is not returned in the response, call GetContextDatabaseWorkspace to query the workspace details and confirm the actual mode. Do not recreate the workspace simply because this field is missing.
	//
	// example:
	//
	// advanced
	PermissionMode *string `json:"PermissionMode,omitempty" xml:"PermissionMode,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// FE9C65D7-930F-57A5-A207-8C396329****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The role of the first member, which is fixed to owner.
	//
	// example:
	//
	// owner
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
	// The status of the workspace (Active: running normally; Locked: stopped due to overdue payments or locked due to expiration).
	//
	// example:
	//
	// Active
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The type of the workspace (personal: individual, enterprise: enterprise).
	//
	// example:
	//
	// personal
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The viewer identity information for the current call, including the identity parsing mode, caller account type, and member information. Member-related fields are optional.
	Viewer *CreateContextDatabaseWorkspaceResponseBodyViewer `json:"Viewer,omitempty" xml:"Viewer,omitempty" type:"Struct"`
	// The ID of the new workspace.
	//
	// example:
	//
	// ws-as1llqmkol****
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
	// The name of the workspace.
	//
	// example:
	//
	// my-workspace
	WorkspaceName *string `json:"WorkspaceName,omitempty" xml:"WorkspaceName,omitempty"`
}

func (s CreateContextDatabaseWorkspaceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateContextDatabaseWorkspaceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateContextDatabaseWorkspaceResponseBody) GetApiKey() *string {
	return s.ApiKey
}

func (s *CreateContextDatabaseWorkspaceResponseBody) GetApiKeyName() *string {
	return s.ApiKeyName
}

func (s *CreateContextDatabaseWorkspaceResponseBody) GetCapabilities() *CreateContextDatabaseWorkspaceResponseBodyCapabilities {
	return s.Capabilities
}

func (s *CreateContextDatabaseWorkspaceResponseBody) GetClusterId() *string {
	return s.ClusterId
}

func (s *CreateContextDatabaseWorkspaceResponseBody) GetConsolePermissionMode() *string {
	return s.ConsolePermissionMode
}

func (s *CreateContextDatabaseWorkspaceResponseBody) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *CreateContextDatabaseWorkspaceResponseBody) GetMemberId() *string {
	return s.MemberId
}

func (s *CreateContextDatabaseWorkspaceResponseBody) GetMemberName() *string {
	return s.MemberName
}

func (s *CreateContextDatabaseWorkspaceResponseBody) GetPermissionMode() *string {
	return s.PermissionMode
}

func (s *CreateContextDatabaseWorkspaceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateContextDatabaseWorkspaceResponseBody) GetRole() *string {
	return s.Role
}

func (s *CreateContextDatabaseWorkspaceResponseBody) GetStatus() *string {
	return s.Status
}

func (s *CreateContextDatabaseWorkspaceResponseBody) GetType() *string {
	return s.Type
}

func (s *CreateContextDatabaseWorkspaceResponseBody) GetViewer() *CreateContextDatabaseWorkspaceResponseBodyViewer {
	return s.Viewer
}

func (s *CreateContextDatabaseWorkspaceResponseBody) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *CreateContextDatabaseWorkspaceResponseBody) GetWorkspaceName() *string {
	return s.WorkspaceName
}

func (s *CreateContextDatabaseWorkspaceResponseBody) SetApiKey(v string) *CreateContextDatabaseWorkspaceResponseBody {
	s.ApiKey = &v
	return s
}

func (s *CreateContextDatabaseWorkspaceResponseBody) SetApiKeyName(v string) *CreateContextDatabaseWorkspaceResponseBody {
	s.ApiKeyName = &v
	return s
}

func (s *CreateContextDatabaseWorkspaceResponseBody) SetCapabilities(v *CreateContextDatabaseWorkspaceResponseBodyCapabilities) *CreateContextDatabaseWorkspaceResponseBody {
	s.Capabilities = v
	return s
}

func (s *CreateContextDatabaseWorkspaceResponseBody) SetClusterId(v string) *CreateContextDatabaseWorkspaceResponseBody {
	s.ClusterId = &v
	return s
}

func (s *CreateContextDatabaseWorkspaceResponseBody) SetConsolePermissionMode(v string) *CreateContextDatabaseWorkspaceResponseBody {
	s.ConsolePermissionMode = &v
	return s
}

func (s *CreateContextDatabaseWorkspaceResponseBody) SetCreatedAt(v string) *CreateContextDatabaseWorkspaceResponseBody {
	s.CreatedAt = &v
	return s
}

func (s *CreateContextDatabaseWorkspaceResponseBody) SetMemberId(v string) *CreateContextDatabaseWorkspaceResponseBody {
	s.MemberId = &v
	return s
}

func (s *CreateContextDatabaseWorkspaceResponseBody) SetMemberName(v string) *CreateContextDatabaseWorkspaceResponseBody {
	s.MemberName = &v
	return s
}

func (s *CreateContextDatabaseWorkspaceResponseBody) SetPermissionMode(v string) *CreateContextDatabaseWorkspaceResponseBody {
	s.PermissionMode = &v
	return s
}

func (s *CreateContextDatabaseWorkspaceResponseBody) SetRequestId(v string) *CreateContextDatabaseWorkspaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateContextDatabaseWorkspaceResponseBody) SetRole(v string) *CreateContextDatabaseWorkspaceResponseBody {
	s.Role = &v
	return s
}

func (s *CreateContextDatabaseWorkspaceResponseBody) SetStatus(v string) *CreateContextDatabaseWorkspaceResponseBody {
	s.Status = &v
	return s
}

func (s *CreateContextDatabaseWorkspaceResponseBody) SetType(v string) *CreateContextDatabaseWorkspaceResponseBody {
	s.Type = &v
	return s
}

func (s *CreateContextDatabaseWorkspaceResponseBody) SetViewer(v *CreateContextDatabaseWorkspaceResponseBodyViewer) *CreateContextDatabaseWorkspaceResponseBody {
	s.Viewer = v
	return s
}

func (s *CreateContextDatabaseWorkspaceResponseBody) SetWorkspaceId(v string) *CreateContextDatabaseWorkspaceResponseBody {
	s.WorkspaceId = &v
	return s
}

func (s *CreateContextDatabaseWorkspaceResponseBody) SetWorkspaceName(v string) *CreateContextDatabaseWorkspaceResponseBody {
	s.WorkspaceName = &v
	return s
}

func (s *CreateContextDatabaseWorkspaceResponseBody) Validate() error {
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

type CreateContextDatabaseWorkspaceResponseBodyCapabilities struct {
	// The list of operation names that the account compatibility adapter can perform. The list content does not represent read/write, member group, or access control authorization for specific resources.
	ExecutableActions []*string `json:"ExecutableActions,omitempty" xml:"ExecutableActions,omitempty" type:"Repeated"`
}

func (s CreateContextDatabaseWorkspaceResponseBodyCapabilities) String() string {
	return dara.Prettify(s)
}

func (s CreateContextDatabaseWorkspaceResponseBodyCapabilities) GoString() string {
	return s.String()
}

func (s *CreateContextDatabaseWorkspaceResponseBodyCapabilities) GetExecutableActions() []*string {
	return s.ExecutableActions
}

func (s *CreateContextDatabaseWorkspaceResponseBodyCapabilities) SetExecutableActions(v []*string) *CreateContextDatabaseWorkspaceResponseBodyCapabilities {
	s.ExecutableActions = v
	return s
}

func (s *CreateContextDatabaseWorkspaceResponseBodyCapabilities) Validate() error {
	return dara.Validate(s)
}

type CreateContextDatabaseWorkspaceResponseBodyViewer struct {
	// The ID of the workspace member bound in the identity information of the current call. This field is optional and is empty if no bound member information is provided.
	//
	// example:
	//
	// member-example
	BoundMemberId *string `json:"BoundMemberId,omitempty" xml:"BoundMemberId,omitempty"`
	// The role of the bound member in the workspace, such as `Owner`. This field is optional and is empty if no bound member role information is provided.
	//
	// example:
	//
	// Owner
	BoundMemberRole *string `json:"BoundMemberRole,omitempty" xml:"BoundMemberRole,omitempty"`
	// The status of the bound member, such as `active`. This field is optional and is empty if no bound member status information is provided.
	//
	// example:
	//
	// active
	BoundMemberStatus *string `json:"BoundMemberStatus,omitempty" xml:"BoundMemberStatus,omitempty"`
	// The account type of the caller for the current call, such as `customer`. This field describes the caller account and does not represent the role of the workspace member.
	//
	// example:
	//
	// customer
	CallerType *string `json:"CallerType,omitempty" xml:"CallerType,omitempty"`
	// The identity parsing mode for the current call. Valid values:
	//
	// - `account_compatibility`: The account compatibility identity mode.
	//
	// - `bound_member`: The bound member identity mode.
	//
	// The identity parsing mode and the permission mode of the workspace are different concepts. You cannot infer whether advanced permissions are enabled for the workspace based on this field.
	//
	// example:
	//
	// account_compatibility
	IdentityMode *string `json:"IdentityMode,omitempty" xml:"IdentityMode,omitempty"`
	// The ID of the workspace member selected in the account compatibility identity mode. This field is optional and is empty if no member is selected for the current call. This field and the bound member ID represent the member selection result and the binding relationship, respectively.
	//
	// example:
	//
	// member-example
	SelectedMemberId *string `json:"SelectedMemberId,omitempty" xml:"SelectedMemberId,omitempty"`
}

func (s CreateContextDatabaseWorkspaceResponseBodyViewer) String() string {
	return dara.Prettify(s)
}

func (s CreateContextDatabaseWorkspaceResponseBodyViewer) GoString() string {
	return s.String()
}

func (s *CreateContextDatabaseWorkspaceResponseBodyViewer) GetBoundMemberId() *string {
	return s.BoundMemberId
}

func (s *CreateContextDatabaseWorkspaceResponseBodyViewer) GetBoundMemberRole() *string {
	return s.BoundMemberRole
}

func (s *CreateContextDatabaseWorkspaceResponseBodyViewer) GetBoundMemberStatus() *string {
	return s.BoundMemberStatus
}

func (s *CreateContextDatabaseWorkspaceResponseBodyViewer) GetCallerType() *string {
	return s.CallerType
}

func (s *CreateContextDatabaseWorkspaceResponseBodyViewer) GetIdentityMode() *string {
	return s.IdentityMode
}

func (s *CreateContextDatabaseWorkspaceResponseBodyViewer) GetSelectedMemberId() *string {
	return s.SelectedMemberId
}

func (s *CreateContextDatabaseWorkspaceResponseBodyViewer) SetBoundMemberId(v string) *CreateContextDatabaseWorkspaceResponseBodyViewer {
	s.BoundMemberId = &v
	return s
}

func (s *CreateContextDatabaseWorkspaceResponseBodyViewer) SetBoundMemberRole(v string) *CreateContextDatabaseWorkspaceResponseBodyViewer {
	s.BoundMemberRole = &v
	return s
}

func (s *CreateContextDatabaseWorkspaceResponseBodyViewer) SetBoundMemberStatus(v string) *CreateContextDatabaseWorkspaceResponseBodyViewer {
	s.BoundMemberStatus = &v
	return s
}

func (s *CreateContextDatabaseWorkspaceResponseBodyViewer) SetCallerType(v string) *CreateContextDatabaseWorkspaceResponseBodyViewer {
	s.CallerType = &v
	return s
}

func (s *CreateContextDatabaseWorkspaceResponseBodyViewer) SetIdentityMode(v string) *CreateContextDatabaseWorkspaceResponseBodyViewer {
	s.IdentityMode = &v
	return s
}

func (s *CreateContextDatabaseWorkspaceResponseBodyViewer) SetSelectedMemberId(v string) *CreateContextDatabaseWorkspaceResponseBodyViewer {
	s.SelectedMemberId = &v
	return s
}

func (s *CreateContextDatabaseWorkspaceResponseBodyViewer) Validate() error {
	return dara.Validate(s)
}
