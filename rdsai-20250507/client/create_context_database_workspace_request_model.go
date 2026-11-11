// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateContextDatabaseWorkspaceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *CreateContextDatabaseWorkspaceRequest
	GetClusterId() *string
	SetGenerateInitialKey(v bool) *CreateContextDatabaseWorkspaceRequest
	GetGenerateInitialKey() *bool
	SetMemberName(v string) *CreateContextDatabaseWorkspaceRequest
	GetMemberName() *string
	SetPermissionMode(v string) *CreateContextDatabaseWorkspaceRequest
	GetPermissionMode() *string
	SetWorkspaceName(v string) *CreateContextDatabaseWorkspaceRequest
	GetWorkspaceName() *string
}

type CreateContextDatabaseWorkspaceRequest struct {
	// The ID of the cluster to which the workspace belongs. This parameter is optional.
	//
	// - If you do not specify this parameter, the cluster selection logic that prioritizes account-exclusive storage is used.
	//
	// - `default`: The shared cluster is used.
	//
	// - Cluster IDs starting with `acc-`: The associated dedicated cluster is used.
	//
	// example:
	//
	// default
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// Specifies whether to generate an initial API key. This parameter is optional. Default value: `true`.
	//
	// - `true`: Generates the first API key and returns the key name and plaintext key.
	//
	// - `false`: Does not generate an initial API key. The ApiKeyName and ApiKey fields in the response are empty.
	//
	// example:
	//
	// true
	GenerateInitialKey *bool `json:"GenerateInitialKey,omitempty" xml:"GenerateInitialKey,omitempty"`
	// The name of the first member when the workspace is created. This parameter is optional.
	//
	// example:
	//
	// my-member
	MemberName *string `json:"MemberName,omitempty" xml:"MemberName,omitempty"`
	// The permission mode selected when the workspace is created. Valid values:
	//
	// - `compatibility`: The compatibility mode. This mode is used if you do not specify this parameter.
	//
	// - `advanced`: The advanced permission mode. Only the Alibaba Cloud account can select this mode. It takes effect immediately after the workspace is created and cannot be disabled.
	//
	// This parameter is a string, not a Boolean switch. Valid values are case-insensitive and ignore leading and trailing whitespace. If you explicitly pass an empty string, pure whitespace, or other values, `InvalidParameter` (HTTP 400) is returned. If a RAM user selects `advanced`, `NoPermission` (HTTP 403) is returned.
	//
	// example:
	//
	// advanced
	PermissionMode *string `json:"PermissionMode,omitempty" xml:"PermissionMode,omitempty"`
	// The name of the workspace.
	//
	// This parameter is required.
	//
	// example:
	//
	// my-workspace
	WorkspaceName *string `json:"WorkspaceName,omitempty" xml:"WorkspaceName,omitempty"`
}

func (s CreateContextDatabaseWorkspaceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateContextDatabaseWorkspaceRequest) GoString() string {
	return s.String()
}

func (s *CreateContextDatabaseWorkspaceRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *CreateContextDatabaseWorkspaceRequest) GetGenerateInitialKey() *bool {
	return s.GenerateInitialKey
}

func (s *CreateContextDatabaseWorkspaceRequest) GetMemberName() *string {
	return s.MemberName
}

func (s *CreateContextDatabaseWorkspaceRequest) GetPermissionMode() *string {
	return s.PermissionMode
}

func (s *CreateContextDatabaseWorkspaceRequest) GetWorkspaceName() *string {
	return s.WorkspaceName
}

func (s *CreateContextDatabaseWorkspaceRequest) SetClusterId(v string) *CreateContextDatabaseWorkspaceRequest {
	s.ClusterId = &v
	return s
}

func (s *CreateContextDatabaseWorkspaceRequest) SetGenerateInitialKey(v bool) *CreateContextDatabaseWorkspaceRequest {
	s.GenerateInitialKey = &v
	return s
}

func (s *CreateContextDatabaseWorkspaceRequest) SetMemberName(v string) *CreateContextDatabaseWorkspaceRequest {
	s.MemberName = &v
	return s
}

func (s *CreateContextDatabaseWorkspaceRequest) SetPermissionMode(v string) *CreateContextDatabaseWorkspaceRequest {
	s.PermissionMode = &v
	return s
}

func (s *CreateContextDatabaseWorkspaceRequest) SetWorkspaceName(v string) *CreateContextDatabaseWorkspaceRequest {
	s.WorkspaceName = &v
	return s
}

func (s *CreateContextDatabaseWorkspaceRequest) Validate() error {
	return dara.Validate(s)
}
