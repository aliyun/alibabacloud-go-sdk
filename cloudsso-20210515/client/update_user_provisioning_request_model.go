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
	// The directory ID.
	//
	// example:
	//
	// d-003qew84****
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	// The new deletion policy. This policy specifies how to handle synchronized Resource Access Management (RAM) users when the Resource Access Management (RAM) user synchronization is deleted. Valid values:
	//
	// - Delete: Delete the users. When the Resource Access Management (RAM) user synchronization is deleted, the system deletes the Resource Access Management (RAM) users that have been synchronized from CloudSSO to Resource Access Management (RAM).
	//
	// - Keep: Keep the users. When the Resource Access Management (RAM) user synchronization is deleted, the system retains the Resource Access Management (RAM) users that have been synchronized from CloudSSO to Resource Access Management (RAM).
	//
	// example:
	//
	// Delete
	NewDeletionStrategy *string `json:"NewDeletionStrategy,omitempty" xml:"NewDeletionStrategy,omitempty"`
	// The new description of the Resource Access Management (RAM) user synchronization.
	//
	// example:
	//
	// description*****
	NewDescription *string `json:"NewDescription,omitempty" xml:"NewDescription,omitempty"`
	// The new duplication strategy. This strategy specifies how to handle name conflicts when a CloudSSO user is synchronized to Resource Access Management (RAM) user with the same name already exists. Valid values:
	//
	// - KeepBoth: Keep both users. When a CloudSSO user is synchronized to Resource Access Management (RAM) user with the same name already exists, the system appends the `_sso` suffix to the CloudSSO username and attempts to create a RAM user with the new username.
	//
	// - TakeOver: Replace the existing user. When a CloudSSO user is synchronized to Resource Access Management (RAM) user with the same name already exists, the system directly replaces the existing RAM user with the synchronized CloudSSO user.
	//
	// example:
	//
	// KeepBoth
	NewDuplicationStrategy *string `json:"NewDuplicationStrategy,omitempty" xml:"NewDuplicationStrategy,omitempty"`
	// The ID of the Resource Access Management (RAM) user synchronization.
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
