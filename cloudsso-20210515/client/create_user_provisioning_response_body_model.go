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
	// The Resource Access Management (RAM) user synchronization information.
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
	// The ID of the Alibaba Cloud account to which the directory belongs.
	//
	// example:
	//
	// 1639738******
	OwnerPk *string `json:"OwnerPk,omitempty" xml:"OwnerPk,omitempty"`
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
	// The name of the identity for the Resource Access Management (RAM) user synchronization. Valid values:
	//
	// - When `PrincipalType` is set to `Group`, the value is the name of the CloudSSO group.
	//
	// - When `PrincipalType` is set to `User`, the value is the CloudSSO username.
	//
	// example:
	//
	// testGroupName
	PrincipalName *string `json:"PrincipalName,omitempty" xml:"PrincipalName,omitempty"`
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
	// The status of the Resource Access Management (RAM) user synchronization. Valid values:
	//
	// - Enabled: The Resource Access Management (RAM) user synchronization is enabled.
	//
	// - Disabled: The Resource Access Management (RAM) user synchronization is disabled.
	//
	// example:
	//
	// Enabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the destination for the Resource Access Management (RAM) user synchronization. Currently, the value is the ID of the account in a resource directory.
	//
	// example:
	//
	// 1743382******
	TargetId *string `json:"TargetId,omitempty" xml:"TargetId,omitempty"`
	// The name of the destination for the Resource Access Management (RAM) user synchronization. Currently, the value is the name of the account in a resource directory.
	//
	// example:
	//
	// testTargetName
	TargetName *string `json:"TargetName,omitempty" xml:"TargetName,omitempty"`
	// The path of the destination account in a resource directory for the Resource Access Management (RAM) user synchronization.
	TargetPath *string `json:"TargetPath,omitempty" xml:"TargetPath,omitempty"`
	// The type of the destination for the Resource Access Management (RAM) user synchronization. Currently, the only valid value is `RD-Account`.
	//
	// example:
	//
	// RD-Account
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
	// The update time. The time is displayed in UTC.
	//
	// example:
	//
	// 2022-11-28T03:55:42Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// The ID of the Resource Access Management (RAM) user synchronization.
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
