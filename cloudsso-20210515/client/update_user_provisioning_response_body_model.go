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
	// The Resource Access Management (RAM) user synchronization information.
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
	// The creation time. The time is displayed in UTC and follows the RFC 3339 standard in the YYYY-MM-DDTHH:mm:ssZ format.
	//
	// example:
	//
	// 2022-11-28T03:55:42Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The deletion policy. This policy specifies how to handle synchronized Resource Access Management (RAM) users when the Resource Access Management (RAM) user synchronization is deleted. Valid values:
	//
	// - Delete: Delete the users. When the Resource Access Management (RAM) user synchronization is deleted, the system deletes the Resource Access Management (RAM) users that have been synchronized from CloudSSO to Resource Access Management (RAM).
	//
	// - Keep: Keep the users. When the Resource Access Management (RAM) user synchronization is deleted, the system retains the Resource Access Management (RAM) users that have been synchronized from CloudSSO to Resource Access Management (RAM).
	//
	// example:
	//
	// Delete
	DeletionStrategy *string `json:"DeletionStrategy,omitempty" xml:"DeletionStrategy,omitempty"`
	// The description of the Resource Access Management (RAM) user synchronization.
	//
	// example:
	//
	// This is a user provisioning.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The directory ID.
	//
	// example:
	//
	// d-003qew84****
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	// The duplication strategy. This strategy specifies how to handle name conflicts when a CloudSSO user is synchronized to Resource Access Management (RAM) user with the same name already exists. Valid values:
	//
	// - KeepBoth: Keep both users. When a CloudSSO user is synchronized to Resource Access Management (RAM) user with the same name already exists, the system appends the `_sso` suffix to the CloudSSO username and attempts to create a RAM user with the new username.
	//
	// - TakeOver: Replace the existing user. When a CloudSSO user is synchronized to Resource Access Management (RAM) user with the same name already exists, the system directly replaces the existing RAM user with the synchronized CloudSSO user.
	//
	// example:
	//
	// KeepBoth
	DuplicationStrategy *string `json:"DuplicationStrategy,omitempty" xml:"DuplicationStrategy,omitempty"`
	// The ID of the Alibaba Cloud account to which the directory belongs.
	//
	// example:
	//
	// 164987310*****
	OwnerPk *string `json:"OwnerPk,omitempty" xml:"OwnerPk,omitempty"`
	// The ID of the principal for the Resource Access Management (RAM) user synchronization. Valid values:
	//
	// - If `PrincipalType` is set to `Group`, the value is the ID of the CloudSSO group (g-********).
	//
	// - If `PrincipalType` is set to `User`, the value is the ID of the CloudSSO user (u-********).
	//
	// example:
	//
	// g-02ha881d*****
	PrincipalId *string `json:"PrincipalId,omitempty" xml:"PrincipalId,omitempty"`
	// The name of the principal for the Resource Access Management (RAM) user synchronization. Valid values:
	//
	// - If `PrincipalType` is set to `Group`, the value is the name of the CloudSSO group.
	//
	// - If `PrincipalType` is set to `User`, the value is the CloudSSO username.
	//
	// example:
	//
	// testUserName
	PrincipalName *string `json:"PrincipalName,omitempty" xml:"PrincipalName,omitempty"`
	// The type of the principal for the Resource Access Management (RAM) user synchronization. Valid values:
	//
	// - User: The principal is a CloudSSO user.
	//
	// - Group: The principal is a CloudSSO group.
	//
	// example:
	//
	// User
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
	// The ID of the target for the Resource Access Management (RAM) user synchronization. Currently, the value is the ID of the account in a resource directory.
	//
	// example:
	//
	// u-02ha881d*****
	TargetId *string `json:"TargetId,omitempty" xml:"TargetId,omitempty"`
	// The name of the target for the Resource Access Management (RAM) user synchronization. Currently, the value is the name of the account in a resource directory.
	//
	// example:
	//
	// testMemberName
	TargetName *string `json:"TargetName,omitempty" xml:"TargetName,omitempty"`
	// The resource directory path of the target for the Resource Access Management (RAM) user synchronization.
	TargetPath *string `json:"TargetPath,omitempty" xml:"TargetPath,omitempty"`
	// The type of the target for the Resource Access Management (RAM) user synchronization. Currently, the value is `RD-Account`.
	//
	// example:
	//
	// RD-Account
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
	// The update time. The time is displayed in UTC and follows the RFC 3339 standard in the YYYY-MM-DDTHH:mm:ssZ format.
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
