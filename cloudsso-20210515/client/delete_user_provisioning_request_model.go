// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteUserProvisioningRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeletionStrategy(v string) *DeleteUserProvisioningRequest
	GetDeletionStrategy() *string
	SetDirectoryId(v string) *DeleteUserProvisioningRequest
	GetDirectoryId() *string
	SetUserProvisioningId(v string) *DeleteUserProvisioningRequest
	GetUserProvisioningId() *string
}

type DeleteUserProvisioningRequest struct {
	// The deletion strategy. The processing strategy for synchronized RAM users when a RAM user synchronization is deleted. Valid values:
	//
	// - Delete: Deletes the RAM users that have been synchronized from CloudSSO to Resource Access Management (RAM) user synchronization is deleted.
	//
	// - Keep: Retains the RAM users that have been synchronized from CloudSSO to Resource Access Management (RAM) user synchronization is deleted.
	//
	// > If this parameter is left empty, the deletion strategy specified when the RAM user synchronization was created is used.
	//
	// example:
	//
	// Delete
	DeletionStrategy *string `json:"DeletionStrategy,omitempty" xml:"DeletionStrategy,omitempty"`
	// The directory ID.
	//
	// example:
	//
	// d-003qew84****
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	// The ID of the Resource Access Management (RAM) user synchronization.
	//
	// example:
	//
	// up-002axzhapcbz6e63****
	UserProvisioningId *string `json:"UserProvisioningId,omitempty" xml:"UserProvisioningId,omitempty"`
}

func (s DeleteUserProvisioningRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteUserProvisioningRequest) GoString() string {
	return s.String()
}

func (s *DeleteUserProvisioningRequest) GetDeletionStrategy() *string {
	return s.DeletionStrategy
}

func (s *DeleteUserProvisioningRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *DeleteUserProvisioningRequest) GetUserProvisioningId() *string {
	return s.UserProvisioningId
}

func (s *DeleteUserProvisioningRequest) SetDeletionStrategy(v string) *DeleteUserProvisioningRequest {
	s.DeletionStrategy = &v
	return s
}

func (s *DeleteUserProvisioningRequest) SetDirectoryId(v string) *DeleteUserProvisioningRequest {
	s.DirectoryId = &v
	return s
}

func (s *DeleteUserProvisioningRequest) SetUserProvisioningId(v string) *DeleteUserProvisioningRequest {
	s.UserProvisioningId = &v
	return s
}

func (s *DeleteUserProvisioningRequest) Validate() error {
	return dara.Validate(s)
}
