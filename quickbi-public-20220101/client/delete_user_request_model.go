// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteUserRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTransferUserId(v string) *DeleteUserRequest
	GetTransferUserId() *string
	SetUserId(v string) *DeleteUserRequest
	GetUserId() *string
}

type DeleteUserRequest struct {
	// The ID of the successor. If not empty, the report resources in the workspace of the deleted user will be transferred to the successor; otherwise, they will be transferred to the space owner.
	//
	// - The successor cannot be an organization visitor
	//
	// - The permissions of the successor in the workspace must not be less than those of the deleted user, with management permissions > development permissions > sharing permissions > viewing permissions
	//
	// - If the successor is not in the workspace, they will be automatically added to the workspace
	//
	// example:
	//
	// f5****afccd9e434a274
	TransferUserId *string `json:"TransferUserId,omitempty" xml:"TransferUserId,omitempty"`
	// The ID of the user to be deleted. This user ID is the Quick BI UserID, not the Alibaba Cloud UID.
	//
	// This parameter is required.
	//
	// example:
	//
	// fe67f61a35a94b7da1a34ba174a7****
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s DeleteUserRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteUserRequest) GoString() string {
	return s.String()
}

func (s *DeleteUserRequest) GetTransferUserId() *string {
	return s.TransferUserId
}

func (s *DeleteUserRequest) GetUserId() *string {
	return s.UserId
}

func (s *DeleteUserRequest) SetTransferUserId(v string) *DeleteUserRequest {
	s.TransferUserId = &v
	return s
}

func (s *DeleteUserRequest) SetUserId(v string) *DeleteUserRequest {
	s.UserId = &v
	return s
}

func (s *DeleteUserRequest) Validate() error {
	return dara.Validate(s)
}
