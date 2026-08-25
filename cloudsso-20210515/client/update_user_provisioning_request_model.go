// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateUserProvisioningRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDirectoryId(v string) *UpdateUserProvisioningRequest
	GetDirectoryId() *string
	SetNewDeletionStrategy(v string) *UpdateUserProvisioningRequest
	GetNewDeletionStrategy() *string
	SetNewDescription(v string) *UpdateUserProvisioningRequest
	GetNewDescription() *string
	SetNewDuplicationStrategy(v string) *UpdateUserProvisioningRequest
	GetNewDuplicationStrategy() *string
	SetUserProvisioningId(v string) *UpdateUserProvisioningRequest
	GetUserProvisioningId() *string
}

type UpdateUserProvisioningRequest struct {
	// The ID of the resource directory.
	//
	// example:
	//
	// d-003qew84****
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	// The new deletion policy. The policy is used to manage synchronized users when you delete the RAM user provisioning. Valid values:
	//
	// - Delete: When you delete the RAM user provisioning, the system deletes the synchronized users.
	//
	// - Keep: When you delete the RAM user provisioning, the system retains the synchronized users.
	//
	// example:
	//
	// Delete
	NewDeletionStrategy *string `json:"NewDeletionStrategy,omitempty" xml:"NewDeletionStrategy,omitempty"`
	// The new description of the RAM user provisioning.
	//
	// example:
	//
	// description*****
	NewDescription *string `json:"NewDescription,omitempty" xml:"NewDescription,omitempty"`
	// The new conflict handling policy. The policy is used when a RAM user has the same username as the CloudSSO user who is synchronized to RAM. Valid values:
	//
	// - KeepBoth: When a CloudSSO user is synchronized to RAM, if a RAM user who has the same username as the CloudSSO user exists, the system creates a RAM user whose username is the username of the CloudSSO user plus the suffix `_sso`.
	//
	// - TakeOver: When a CloudSSO user is synchronized to RAM, if a RAM user who has the same username as the CloudSSO user exists, the system replaces the RAM user with the CloudSSO user.
	//
	// example:
	//
	// KeepBoth
	NewDuplicationStrategy *string `json:"NewDuplicationStrategy,omitempty" xml:"NewDuplicationStrategy,omitempty"`
	// The ID of the RAM user provisioning.
	//
	// example:
	//
	// up-002axzhapcbz6e63****
	UserProvisioningId *string `json:"UserProvisioningId,omitempty" xml:"UserProvisioningId,omitempty"`
}

func (s UpdateUserProvisioningRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateUserProvisioningRequest) GoString() string {
	return s.String()
}

func (s *UpdateUserProvisioningRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *UpdateUserProvisioningRequest) GetNewDeletionStrategy() *string {
	return s.NewDeletionStrategy
}

func (s *UpdateUserProvisioningRequest) GetNewDescription() *string {
	return s.NewDescription
}

func (s *UpdateUserProvisioningRequest) GetNewDuplicationStrategy() *string {
	return s.NewDuplicationStrategy
}

func (s *UpdateUserProvisioningRequest) GetUserProvisioningId() *string {
	return s.UserProvisioningId
}

func (s *UpdateUserProvisioningRequest) SetDirectoryId(v string) *UpdateUserProvisioningRequest {
	s.DirectoryId = &v
	return s
}

func (s *UpdateUserProvisioningRequest) SetNewDeletionStrategy(v string) *UpdateUserProvisioningRequest {
	s.NewDeletionStrategy = &v
	return s
}

func (s *UpdateUserProvisioningRequest) SetNewDescription(v string) *UpdateUserProvisioningRequest {
	s.NewDescription = &v
	return s
}

func (s *UpdateUserProvisioningRequest) SetNewDuplicationStrategy(v string) *UpdateUserProvisioningRequest {
	s.NewDuplicationStrategy = &v
	return s
}

func (s *UpdateUserProvisioningRequest) SetUserProvisioningId(v string) *UpdateUserProvisioningRequest {
	s.UserProvisioningId = &v
	return s
}

func (s *UpdateUserProvisioningRequest) Validate() error {
	return dara.Validate(s)
}
