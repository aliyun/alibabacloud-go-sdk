// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateUserProvisioningResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateUserProvisioningResponseBody
	GetRequestId() *string
	SetUserProvisioning(v *CreateUserProvisioningResponseBodyUserProvisioning) *CreateUserProvisioningResponseBody
	GetUserProvisioning() *CreateUserProvisioningResponseBodyUserProvisioning
}

type CreateUserProvisioningResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// F6F90F3D-4502-5877-B80B-97476F6AE2CC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the RAM user provisioning.
	UserProvisioning *CreateUserProvisioningResponseBodyUserProvisioning `json:"UserProvisioning,omitempty" xml:"UserProvisioning,omitempty" type:"Struct"`
}

func (s CreateUserProvisioningResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateUserProvisioningResponseBody) GoString() string {
	return s.String()
}

func (s *CreateUserProvisioningResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateUserProvisioningResponseBody) GetUserProvisioning() *CreateUserProvisioningResponseBodyUserProvisioning {
	return s.UserProvisioning
}

func (s *CreateUserProvisioningResponseBody) SetRequestId(v string) *CreateUserProvisioningResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateUserProvisioningResponseBody) SetUserProvisioning(v *CreateUserProvisioningResponseBodyUserProvisioning) *CreateUserProvisioningResponseBody {
	s.UserProvisioning = v
	return s
}

func (s *CreateUserProvisioningResponseBody) Validate() error {
	if s.UserProvisioning != nil {
		if err := s.UserProvisioning.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateUserProvisioningResponseBodyUserProvisioning struct {
	// The creation time. The time is displayed in UTC.
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
	// The ID of the Alibaba Cloud account to which the resource directory belongs.
	//
	// example:
	//
	// 1639738******
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
	// testGroupName
	PrincipalName *string `json:"PrincipalName,omitempty" xml:"PrincipalName,omitempty"`
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
	// The ID of the object for which you create the RAM user provisioning. The value is fixed as the ID of the member in the resource directory.
	//
	// example:
	//
	// 1743382******
	TargetId *string `json:"TargetId,omitempty" xml:"TargetId,omitempty"`
	// The name of the object for which you create the RAM user provisioning. The value is fixed as the name of the member in the resource directory.
	//
	// example:
	//
	// testTargetName
	TargetName *string `json:"TargetName,omitempty" xml:"TargetName,omitempty"`
	// The path of the resource directory in which you create the RAM user provisioning for the member.
	TargetPath *string `json:"TargetPath,omitempty" xml:"TargetPath,omitempty"`
	// The object for which you create the RAM user provisioning. The value is fixed as `RD-Account`.
	//
	// example:
	//
	// RD-Account
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
	// The modification time. The time is displayed in UTC.
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

func (s CreateUserProvisioningResponseBodyUserProvisioning) String() string {
	return dara.Prettify(s)
}

func (s CreateUserProvisioningResponseBodyUserProvisioning) GoString() string {
	return s.String()
}

func (s *CreateUserProvisioningResponseBodyUserProvisioning) GetCreateTime() *string {
	return s.CreateTime
}

func (s *CreateUserProvisioningResponseBodyUserProvisioning) GetDeletionStrategy() *string {
	return s.DeletionStrategy
}

func (s *CreateUserProvisioningResponseBodyUserProvisioning) GetDescription() *string {
	return s.Description
}

func (s *CreateUserProvisioningResponseBodyUserProvisioning) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *CreateUserProvisioningResponseBodyUserProvisioning) GetDuplicationStrategy() *string {
	return s.DuplicationStrategy
}

func (s *CreateUserProvisioningResponseBodyUserProvisioning) GetOwnerPk() *string {
	return s.OwnerPk
}

func (s *CreateUserProvisioningResponseBodyUserProvisioning) GetPrincipalId() *string {
	return s.PrincipalId
}

func (s *CreateUserProvisioningResponseBodyUserProvisioning) GetPrincipalName() *string {
	return s.PrincipalName
}

func (s *CreateUserProvisioningResponseBodyUserProvisioning) GetPrincipalType() *string {
	return s.PrincipalType
}

func (s *CreateUserProvisioningResponseBodyUserProvisioning) GetStatus() *string {
	return s.Status
}

func (s *CreateUserProvisioningResponseBodyUserProvisioning) GetTargetId() *string {
	return s.TargetId
}

func (s *CreateUserProvisioningResponseBodyUserProvisioning) GetTargetName() *string {
	return s.TargetName
}

func (s *CreateUserProvisioningResponseBodyUserProvisioning) GetTargetPath() *string {
	return s.TargetPath
}

func (s *CreateUserProvisioningResponseBodyUserProvisioning) GetTargetType() *string {
	return s.TargetType
}

func (s *CreateUserProvisioningResponseBodyUserProvisioning) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *CreateUserProvisioningResponseBodyUserProvisioning) GetUserProvisioningId() *string {
	return s.UserProvisioningId
}

func (s *CreateUserProvisioningResponseBodyUserProvisioning) SetCreateTime(v string) *CreateUserProvisioningResponseBodyUserProvisioning {
	s.CreateTime = &v
	return s
}

func (s *CreateUserProvisioningResponseBodyUserProvisioning) SetDeletionStrategy(v string) *CreateUserProvisioningResponseBodyUserProvisioning {
	s.DeletionStrategy = &v
	return s
}

func (s *CreateUserProvisioningResponseBodyUserProvisioning) SetDescription(v string) *CreateUserProvisioningResponseBodyUserProvisioning {
	s.Description = &v
	return s
}

func (s *CreateUserProvisioningResponseBodyUserProvisioning) SetDirectoryId(v string) *CreateUserProvisioningResponseBodyUserProvisioning {
	s.DirectoryId = &v
	return s
}

func (s *CreateUserProvisioningResponseBodyUserProvisioning) SetDuplicationStrategy(v string) *CreateUserProvisioningResponseBodyUserProvisioning {
	s.DuplicationStrategy = &v
	return s
}

func (s *CreateUserProvisioningResponseBodyUserProvisioning) SetOwnerPk(v string) *CreateUserProvisioningResponseBodyUserProvisioning {
	s.OwnerPk = &v
	return s
}

func (s *CreateUserProvisioningResponseBodyUserProvisioning) SetPrincipalId(v string) *CreateUserProvisioningResponseBodyUserProvisioning {
	s.PrincipalId = &v
	return s
}

func (s *CreateUserProvisioningResponseBodyUserProvisioning) SetPrincipalName(v string) *CreateUserProvisioningResponseBodyUserProvisioning {
	s.PrincipalName = &v
	return s
}

func (s *CreateUserProvisioningResponseBodyUserProvisioning) SetPrincipalType(v string) *CreateUserProvisioningResponseBodyUserProvisioning {
	s.PrincipalType = &v
	return s
}

func (s *CreateUserProvisioningResponseBodyUserProvisioning) SetStatus(v string) *CreateUserProvisioningResponseBodyUserProvisioning {
	s.Status = &v
	return s
}

func (s *CreateUserProvisioningResponseBodyUserProvisioning) SetTargetId(v string) *CreateUserProvisioningResponseBodyUserProvisioning {
	s.TargetId = &v
	return s
}

func (s *CreateUserProvisioningResponseBodyUserProvisioning) SetTargetName(v string) *CreateUserProvisioningResponseBodyUserProvisioning {
	s.TargetName = &v
	return s
}

func (s *CreateUserProvisioningResponseBodyUserProvisioning) SetTargetPath(v string) *CreateUserProvisioningResponseBodyUserProvisioning {
	s.TargetPath = &v
	return s
}

func (s *CreateUserProvisioningResponseBodyUserProvisioning) SetTargetType(v string) *CreateUserProvisioningResponseBodyUserProvisioning {
	s.TargetType = &v
	return s
}

func (s *CreateUserProvisioningResponseBodyUserProvisioning) SetUpdateTime(v string) *CreateUserProvisioningResponseBodyUserProvisioning {
	s.UpdateTime = &v
	return s
}

func (s *CreateUserProvisioningResponseBodyUserProvisioning) SetUserProvisioningId(v string) *CreateUserProvisioningResponseBodyUserProvisioning {
	s.UserProvisioningId = &v
	return s
}

func (s *CreateUserProvisioningResponseBodyUserProvisioning) Validate() error {
	return dara.Validate(s)
}
