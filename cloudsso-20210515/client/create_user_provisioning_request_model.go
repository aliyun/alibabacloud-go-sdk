// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateUserProvisioningRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeletionStrategy(v string) *CreateUserProvisioningRequest
	GetDeletionStrategy() *string
	SetDescription(v string) *CreateUserProvisioningRequest
	GetDescription() *string
	SetDirectoryId(v string) *CreateUserProvisioningRequest
	GetDirectoryId() *string
	SetDuplicationStrategy(v string) *CreateUserProvisioningRequest
	GetDuplicationStrategy() *string
	SetPrincipalId(v string) *CreateUserProvisioningRequest
	GetPrincipalId() *string
	SetPrincipalType(v string) *CreateUserProvisioningRequest
	GetPrincipalType() *string
	SetTargetId(v string) *CreateUserProvisioningRequest
	GetTargetId() *string
	SetTargetType(v string) *CreateUserProvisioningRequest
	GetTargetType() *string
}

type CreateUserProvisioningRequest struct {
	// The deletion policy. Specifies the action to take for synchronized Resource Access Management (RAM) users when a Resource Access Management (RAM) user synchronization is deleted. Valid values:
	//
	// - Delete: Delete. When a Resource Access Management (RAM) user synchronization is deleted, the system deletes the Resource Access Management (RAM) users that have been synchronized from CloudSSO to RAM.
	//
	// - Keep: Keep. When a Resource Access Management (RAM) user synchronization is deleted, the system retains the Resource Access Management (RAM) users that have been synchronized from CloudSSO to RAM.
	//
	// example:
	//
	// Delete
	DeletionStrategy *string `json:"DeletionStrategy,omitempty" xml:"DeletionStrategy,omitempty"`
	// The description.
	//
	// example:
	//
	// This is a user provisioning.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the directory.
	//
	// example:
	//
	// d-003qew84****
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	// The conflict resolution policy. Specifies the action to take when a CloudSSO user is synchronized to Resource Access Management (RAM) user with the same name already exists. Valid values:
	//
	// - KeepBoth: Keep both. When a CloudSSO user is synchronized to Resource Access Management (RAM) user with the same name already exists, the system adds the `_sso` suffix to the username of the CloudSSO user and attempts to create a RAM user with the suffixed username.
	//
	// - TakeOver: Take over. When a CloudSSO user is synchronized to Resource Access Management (RAM) user with the same name already exists, the system replaces the existing RAM user with the synchronized CloudSSO user.
	//
	// example:
	//
	// KeepBoth
	DuplicationStrategy *string `json:"DuplicationStrategy,omitempty" xml:"DuplicationStrategy,omitempty"`
	// The ID of the identity for the Resource Access Management (RAM) user synchronization. Valid values:
	//
	// - When `PrincipalType` is set to `Group`, the value is the ID of the CloudSSO group (g-********).
	//
	// - When `PrincipalType` is set to `User`, the value is the ID of the CloudSSO user (u-********).
	//
	// example:
	//
	// g-02ha881d*****
	PrincipalId *string `json:"PrincipalId,omitempty" xml:"PrincipalId,omitempty"`
	// The type of the Resource Access Management (RAM) user synchronization. Valid values:
	//
	// - User: The Resource Access Management (RAM) user synchronization is a CloudSSO user.
	//
	// - Group: The Resource Access Management (RAM) user synchronization is a CloudSSO group.
	//
	// example:
	//
	// Group
	PrincipalType *string `json:"PrincipalType,omitempty" xml:"PrincipalType,omitempty"`
	// The ID of the destination for the Resource Access Management (RAM) user synchronization. Currently, the value is the ID of the account in a resource directory.
	//
	// example:
	//
	// 1743382******
	TargetId *string `json:"TargetId,omitempty" xml:"TargetId,omitempty"`
	// The type of the destination for the Resource Access Management (RAM) user synchronization. Currently, the only valid value is `RD-Account`.
	//
	// example:
	//
	// RD-Account
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
}

func (s CreateUserProvisioningRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateUserProvisioningRequest) GoString() string {
	return s.String()
}

func (s *CreateUserProvisioningRequest) GetDeletionStrategy() *string {
	return s.DeletionStrategy
}

func (s *CreateUserProvisioningRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateUserProvisioningRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *CreateUserProvisioningRequest) GetDuplicationStrategy() *string {
	return s.DuplicationStrategy
}

func (s *CreateUserProvisioningRequest) GetPrincipalId() *string {
	return s.PrincipalId
}

func (s *CreateUserProvisioningRequest) GetPrincipalType() *string {
	return s.PrincipalType
}

func (s *CreateUserProvisioningRequest) GetTargetId() *string {
	return s.TargetId
}

func (s *CreateUserProvisioningRequest) GetTargetType() *string {
	return s.TargetType
}

func (s *CreateUserProvisioningRequest) SetDeletionStrategy(v string) *CreateUserProvisioningRequest {
	s.DeletionStrategy = &v
	return s
}

func (s *CreateUserProvisioningRequest) SetDescription(v string) *CreateUserProvisioningRequest {
	s.Description = &v
	return s
}

func (s *CreateUserProvisioningRequest) SetDirectoryId(v string) *CreateUserProvisioningRequest {
	s.DirectoryId = &v
	return s
}

func (s *CreateUserProvisioningRequest) SetDuplicationStrategy(v string) *CreateUserProvisioningRequest {
	s.DuplicationStrategy = &v
	return s
}

func (s *CreateUserProvisioningRequest) SetPrincipalId(v string) *CreateUserProvisioningRequest {
	s.PrincipalId = &v
	return s
}

func (s *CreateUserProvisioningRequest) SetPrincipalType(v string) *CreateUserProvisioningRequest {
	s.PrincipalType = &v
	return s
}

func (s *CreateUserProvisioningRequest) SetTargetId(v string) *CreateUserProvisioningRequest {
	s.TargetId = &v
	return s
}

func (s *CreateUserProvisioningRequest) SetTargetType(v string) *CreateUserProvisioningRequest {
	s.TargetType = &v
	return s
}

func (s *CreateUserProvisioningRequest) Validate() error {
	return dara.Validate(s)
}
