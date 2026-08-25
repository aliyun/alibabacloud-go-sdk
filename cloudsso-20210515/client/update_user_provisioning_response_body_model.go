// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateUserProvisioningResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateUserProvisioningResponseBody
	GetRequestId() *string
	SetUserProvisioning(v *UpdateUserProvisioningResponseBodyUserProvisioning) *UpdateUserProvisioningResponseBody
	GetUserProvisioning() *UpdateUserProvisioningResponseBodyUserProvisioning
}

type UpdateUserProvisioningResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// F6F90F3D-4502-5877-B80B-97476F6AE2CC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the RAM user provisioning.
	UserProvisioning *UpdateUserProvisioningResponseBodyUserProvisioning `json:"UserProvisioning,omitempty" xml:"UserProvisioning,omitempty" type:"Struct"`
}

func (s UpdateUserProvisioningResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateUserProvisioningResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateUserProvisioningResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateUserProvisioningResponseBody) GetUserProvisioning() *UpdateUserProvisioningResponseBodyUserProvisioning {
	return s.UserProvisioning
}

func (s *UpdateUserProvisioningResponseBody) SetRequestId(v string) *UpdateUserProvisioningResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateUserProvisioningResponseBody) SetUserProvisioning(v *UpdateUserProvisioningResponseBodyUserProvisioning) *UpdateUserProvisioningResponseBody {
	s.UserProvisioning = v
	return s
}

func (s *UpdateUserProvisioningResponseBody) Validate() error {
	if s.UserProvisioning != nil {
		if err := s.UserProvisioning.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateUserProvisioningResponseBodyUserProvisioning struct {
	// The creation time.
	//
	// example:
	//
	// 2022-11-28T03:55:42Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
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
	// The description for the RAM user provisioning.
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
	// The ID of the Alibaba Cloud account to which the resource directory belongs.
	//
	// example:
	//
	// 164987310*****
	OwnerPk *string `json:"OwnerPk,omitempty" xml:"OwnerPk,omitempty"`
	// The identity ID of the RAM user provisioning. Valid values:
	//
	// - If `Group` is returned for the `PrincipalType` parameter, the value of this parameter is the ID of a CloudSSO user group (g-\\*\\*\\*\\*\\*\\*\\*\\*).
	//
	// - If `User` is returned for the `PrincipalType` parameter, the value of this parameter is the ID of a CloudSSO user (u-\\*\\*\\*\\*\\*\\*\\*\\*).
	//
	// example:
	//
	// g-02ha881d*****
	PrincipalId *string `json:"PrincipalId,omitempty" xml:"PrincipalId,omitempty"`
	// The identity name of the RAM user provisioning. Valid values:
	//
	// - If `Group` is returned for the `PrincipalType` parameter, the value of this parameter is the name of a CloudSSO user group.
	//
	// - If `User` is returned for the `PrincipalType` parameter, the value of this parameter is the name of a CloudSSO user.
	//
	// example:
	//
	// testUserName
	PrincipalName *string `json:"PrincipalName,omitempty" xml:"PrincipalName,omitempty"`
	// The identity type of the RAM user provisioning. Valid values:
	//
	// - User: indicates that the identity of the RAM user provisioning is a CloudSSO user.
	//
	// - Group: indicates that the identity of the RAM user provisioning is a CloudSSO user group.
	//
	// example:
	//
	// User
	PrincipalType *string `json:"PrincipalType,omitempty" xml:"PrincipalType,omitempty"`
	// The status of the RAM user provisioning. Valid values:
	//
	// - Enabled
	//
	// - Disabled
	//
	// example:
	//
	// Enabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the object for which you create the RAM user provisioning. The value is fixed as the ID of the account in the resource directory.
	//
	// example:
	//
	// u-02ha881d*****
	TargetId *string `json:"TargetId,omitempty" xml:"TargetId,omitempty"`
	// The name of the object for which you create the RAM user provisioning. The value is fixed as the name of the resource directory.
	//
	// example:
	//
	// testMemberName
	TargetName *string `json:"TargetName,omitempty" xml:"TargetName,omitempty"`
	// The path of the resource directory in which you create the RAM user provisioning for the object.
	TargetPath *string `json:"TargetPath,omitempty" xml:"TargetPath,omitempty"`
	// The object for which you create the RAM user provisioning. The value is fixed as `RD-Account`.
	//
	// example:
	//
	// RD-Account
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
	// The modification time.
	//
	// example:
	//
	// 2022-11-28T03:55:42Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// The ID of the RAM user provisioning.
	//
	// example:
	//
	// up-002axzhapcbz6e63****
	UserProvisioningId *string `json:"UserProvisioningId,omitempty" xml:"UserProvisioningId,omitempty"`
}

func (s UpdateUserProvisioningResponseBodyUserProvisioning) String() string {
	return dara.Prettify(s)
}

func (s UpdateUserProvisioningResponseBodyUserProvisioning) GoString() string {
	return s.String()
}

func (s *UpdateUserProvisioningResponseBodyUserProvisioning) GetCreateTime() *string {
	return s.CreateTime
}

func (s *UpdateUserProvisioningResponseBodyUserProvisioning) GetDeletionStrategy() *string {
	return s.DeletionStrategy
}

func (s *UpdateUserProvisioningResponseBodyUserProvisioning) GetDescription() *string {
	return s.Description
}

func (s *UpdateUserProvisioningResponseBodyUserProvisioning) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *UpdateUserProvisioningResponseBodyUserProvisioning) GetDuplicationStrategy() *string {
	return s.DuplicationStrategy
}

func (s *UpdateUserProvisioningResponseBodyUserProvisioning) GetOwnerPk() *string {
	return s.OwnerPk
}

func (s *UpdateUserProvisioningResponseBodyUserProvisioning) GetPrincipalId() *string {
	return s.PrincipalId
}

func (s *UpdateUserProvisioningResponseBodyUserProvisioning) GetPrincipalName() *string {
	return s.PrincipalName
}

func (s *UpdateUserProvisioningResponseBodyUserProvisioning) GetPrincipalType() *string {
	return s.PrincipalType
}

func (s *UpdateUserProvisioningResponseBodyUserProvisioning) GetStatus() *string {
	return s.Status
}

func (s *UpdateUserProvisioningResponseBodyUserProvisioning) GetTargetId() *string {
	return s.TargetId
}

func (s *UpdateUserProvisioningResponseBodyUserProvisioning) GetTargetName() *string {
	return s.TargetName
}

func (s *UpdateUserProvisioningResponseBodyUserProvisioning) GetTargetPath() *string {
	return s.TargetPath
}

func (s *UpdateUserProvisioningResponseBodyUserProvisioning) GetTargetType() *string {
	return s.TargetType
}

func (s *UpdateUserProvisioningResponseBodyUserProvisioning) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *UpdateUserProvisioningResponseBodyUserProvisioning) GetUserProvisioningId() *string {
	return s.UserProvisioningId
}

func (s *UpdateUserProvisioningResponseBodyUserProvisioning) SetCreateTime(v string) *UpdateUserProvisioningResponseBodyUserProvisioning {
	s.CreateTime = &v
	return s
}

func (s *UpdateUserProvisioningResponseBodyUserProvisioning) SetDeletionStrategy(v string) *UpdateUserProvisioningResponseBodyUserProvisioning {
	s.DeletionStrategy = &v
	return s
}

func (s *UpdateUserProvisioningResponseBodyUserProvisioning) SetDescription(v string) *UpdateUserProvisioningResponseBodyUserProvisioning {
	s.Description = &v
	return s
}

func (s *UpdateUserProvisioningResponseBodyUserProvisioning) SetDirectoryId(v string) *UpdateUserProvisioningResponseBodyUserProvisioning {
	s.DirectoryId = &v
	return s
}

func (s *UpdateUserProvisioningResponseBodyUserProvisioning) SetDuplicationStrategy(v string) *UpdateUserProvisioningResponseBodyUserProvisioning {
	s.DuplicationStrategy = &v
	return s
}

func (s *UpdateUserProvisioningResponseBodyUserProvisioning) SetOwnerPk(v string) *UpdateUserProvisioningResponseBodyUserProvisioning {
	s.OwnerPk = &v
	return s
}

func (s *UpdateUserProvisioningResponseBodyUserProvisioning) SetPrincipalId(v string) *UpdateUserProvisioningResponseBodyUserProvisioning {
	s.PrincipalId = &v
	return s
}

func (s *UpdateUserProvisioningResponseBodyUserProvisioning) SetPrincipalName(v string) *UpdateUserProvisioningResponseBodyUserProvisioning {
	s.PrincipalName = &v
	return s
}

func (s *UpdateUserProvisioningResponseBodyUserProvisioning) SetPrincipalType(v string) *UpdateUserProvisioningResponseBodyUserProvisioning {
	s.PrincipalType = &v
	return s
}

func (s *UpdateUserProvisioningResponseBodyUserProvisioning) SetStatus(v string) *UpdateUserProvisioningResponseBodyUserProvisioning {
	s.Status = &v
	return s
}

func (s *UpdateUserProvisioningResponseBodyUserProvisioning) SetTargetId(v string) *UpdateUserProvisioningResponseBodyUserProvisioning {
	s.TargetId = &v
	return s
}

func (s *UpdateUserProvisioningResponseBodyUserProvisioning) SetTargetName(v string) *UpdateUserProvisioningResponseBodyUserProvisioning {
	s.TargetName = &v
	return s
}

func (s *UpdateUserProvisioningResponseBodyUserProvisioning) SetTargetPath(v string) *UpdateUserProvisioningResponseBodyUserProvisioning {
	s.TargetPath = &v
	return s
}

func (s *UpdateUserProvisioningResponseBodyUserProvisioning) SetTargetType(v string) *UpdateUserProvisioningResponseBodyUserProvisioning {
	s.TargetType = &v
	return s
}

func (s *UpdateUserProvisioningResponseBodyUserProvisioning) SetUpdateTime(v string) *UpdateUserProvisioningResponseBodyUserProvisioning {
	s.UpdateTime = &v
	return s
}

func (s *UpdateUserProvisioningResponseBodyUserProvisioning) SetUserProvisioningId(v string) *UpdateUserProvisioningResponseBodyUserProvisioning {
	s.UserProvisioningId = &v
	return s
}

func (s *UpdateUserProvisioningResponseBodyUserProvisioning) Validate() error {
	return dara.Validate(s)
}
