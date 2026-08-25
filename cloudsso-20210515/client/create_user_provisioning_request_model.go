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
	// The deletion policy. The policy is used to manage synchronized users when you delete the RAM user provisioning. Valid values:
	//
	// - Delete: When you delete the RAM user provisioning, the system deletes the synchronized users.
	//
	// - Keep: When you delete the RAM user provisioning, the system retains the synchronized users.
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
	// The ID of the resource directory.
	//
	// example:
	//
	// d-003qew84****
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	// The conflict handling policy. The policy is used when a RAM user has the same username as the CloudSSO user who is synchronized to RAM. Valid values:
	//
	// - KeepBoth: When a CloudSSO user is synchronized to RAM, if a RAM user who has the same username as the CloudSSO user exists, the system creates a RAM user whose username is the username of the CloudSSO user plus the suffix `_sso`.
	//
	// - TakeOver: When a CloudSSO user is synchronized to RAM, if a RAM user who has the same username as the CloudSSO user exists, the system replaces the RAM user with the CloudSSO user.
	//
	// example:
	//
	// KeepBoth
	DuplicationStrategy *string `json:"DuplicationStrategy,omitempty" xml:"DuplicationStrategy,omitempty"`
	// The identity ID of the RAM user provisioning. Valid values:
	//
	// - If you set the `PrincipalType` parameter to `Group`, the value of this parameter is the ID of a CloudSSO user group (g-\\*\\*\\*\\*\\*\\*\\*\\*).
	//
	// - If you set the `PrincipalType` parameter to `User`, the value of this parameter is the ID of a CloudSSO user (u-\\*\\*\\*\\*\\*\\*\\*\\*).
	//
	// example:
	//
	// g-02ha881d*****
	PrincipalId *string `json:"PrincipalId,omitempty" xml:"PrincipalId,omitempty"`
	// The identity type of the RAM user provisioning. Valid values:
	//
	// - User: The identity of the RAM user provisioning is a CloudSSO user.
	//
	// - Group: The identity of the RAM user provisioning is a CloudSSO user group.
	//
	// example:
	//
	// Group
	PrincipalType *string `json:"PrincipalType,omitempty" xml:"PrincipalType,omitempty"`
	// The ID of the object for which you create the RAM user provisioning. The value is fixed as the ID of the member in the resource directory.
	//
	// example:
	//
	// 1743382******
	TargetId *string `json:"TargetId,omitempty" xml:"TargetId,omitempty"`
	// The object for which you create the RAM user provisioning. The value is fixed as `RD-Account`.
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
