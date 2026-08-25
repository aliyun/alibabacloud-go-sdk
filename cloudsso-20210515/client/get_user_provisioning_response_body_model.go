// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserProvisioningResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetUserProvisioningResponseBody
	GetRequestId() *string
	SetUserProvisioning(v *GetUserProvisioningResponseBodyUserProvisioning) *GetUserProvisioningResponseBody
	GetUserProvisioning() *GetUserProvisioningResponseBodyUserProvisioning
}

type GetUserProvisioningResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// F6F90F3D-4502-5877-B80B-97476F6AE2CC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the RAM user provisioning.
	UserProvisioning *GetUserProvisioningResponseBodyUserProvisioning `json:"UserProvisioning,omitempty" xml:"UserProvisioning,omitempty" type:"Struct"`
}

func (s GetUserProvisioningResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetUserProvisioningResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserProvisioningResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetUserProvisioningResponseBody) GetUserProvisioning() *GetUserProvisioningResponseBodyUserProvisioning {
	return s.UserProvisioning
}

func (s *GetUserProvisioningResponseBody) SetRequestId(v string) *GetUserProvisioningResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetUserProvisioningResponseBody) SetUserProvisioning(v *GetUserProvisioningResponseBodyUserProvisioning) *GetUserProvisioningResponseBody {
	s.UserProvisioning = v
	return s
}

func (s *GetUserProvisioningResponseBody) Validate() error {
	if s.UserProvisioning != nil {
		if err := s.UserProvisioning.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetUserProvisioningResponseBodyUserProvisioning struct {
	// The creation time.
	//
	// example:
	//
	// 2022-11-28T03:55:42Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The deletion policy. The policy is used to manage synchronized users when you delete the RAM user provisioning. Valid values:
	//
	// 	- Delete: When you delete the RAM user provisioning, the system deletes the synchronized users.
	//
	// 	- Keep: When you delete the RAM user provisioning, the system retains the synchronized users.
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
	// 	- KeepBoth: When a CloudSSO user is synchronized to RAM, if a RAM user who has the same username as the CloudSSO user exists, the system creates a RAM user whose username is the username of the CloudSSO user plus the suffix `_sso`.
	//
	// 	- TakeOver: When a CloudSSO user is synchronized to RAM, if a RAM user who has the same username as the CloudSSO user exists, the system replaces the RAM user with the CloudSSO user.
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
	// 	- If `Group` is returned for the `PrincipalType` parameter, the value of this parameter is the ID of a CloudSSO user group (g-\\*\\*\\*\\*\\*\\*\\*\\*).
	//
	// 	- If `User` is returned for the `PrincipalType` parameter, the value of this parameter is the ID of a CloudSSO user (u-\\*\\*\\*\\*\\*\\*\\*\\*).
	//
	// example:
	//
	// g-02ha881d*****
	PrincipalId *string `json:"PrincipalId,omitempty" xml:"PrincipalId,omitempty"`
	// The identity name of the RAM user provisioning. Valid values:
	//
	// 	- If `Group` is returned for the `PrincipalType` parameter, the value of this parameter is the name of a CloudSSO user group.
	//
	// 	- If `User` is returned for the `PrincipalType` parameter, the value of this parameter is the name of a CloudSSO user.
	//
	// example:
	//
	// testGroupName
	PrincipalName *string `json:"PrincipalName,omitempty" xml:"PrincipalName,omitempty"`
	// The identity type of the RAM user provisioning. Valid values:
	//
	// 	- User: indicates that the identity of the RAM user provisioning is a CloudSSO user.
	//
	// 	- Group: indicates that the identity of the RAM user provisioning is a CloudSSO user group.
	//
	// example:
	//
	// Group
	PrincipalType *string `json:"PrincipalType,omitempty" xml:"PrincipalType,omitempty"`
	// The status of the RAM user provisioning. Valid values:
	//
	// 	- Enabled
	//
	// 	- Disabled
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
	// testRdMember
	TargetName *string `json:"TargetName,omitempty" xml:"TargetName,omitempty"`
	// The path of the resource directory in which you create the RAM user provisioning for the member.
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

func (s GetUserProvisioningResponseBodyUserProvisioning) String() string {
	return dara.Prettify(s)
}

func (s GetUserProvisioningResponseBodyUserProvisioning) GoString() string {
	return s.String()
}

func (s *GetUserProvisioningResponseBodyUserProvisioning) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetUserProvisioningResponseBodyUserProvisioning) GetDeletionStrategy() *string {
	return s.DeletionStrategy
}

func (s *GetUserProvisioningResponseBodyUserProvisioning) GetDescription() *string {
	return s.Description
}

func (s *GetUserProvisioningResponseBodyUserProvisioning) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *GetUserProvisioningResponseBodyUserProvisioning) GetDuplicationStrategy() *string {
	return s.DuplicationStrategy
}

func (s *GetUserProvisioningResponseBodyUserProvisioning) GetOwnerPk() *string {
	return s.OwnerPk
}

func (s *GetUserProvisioningResponseBodyUserProvisioning) GetPrincipalId() *string {
	return s.PrincipalId
}

func (s *GetUserProvisioningResponseBodyUserProvisioning) GetPrincipalName() *string {
	return s.PrincipalName
}

func (s *GetUserProvisioningResponseBodyUserProvisioning) GetPrincipalType() *string {
	return s.PrincipalType
}

func (s *GetUserProvisioningResponseBodyUserProvisioning) GetStatus() *string {
	return s.Status
}

func (s *GetUserProvisioningResponseBodyUserProvisioning) GetTargetId() *string {
	return s.TargetId
}

func (s *GetUserProvisioningResponseBodyUserProvisioning) GetTargetName() *string {
	return s.TargetName
}

func (s *GetUserProvisioningResponseBodyUserProvisioning) GetTargetPath() *string {
	return s.TargetPath
}

func (s *GetUserProvisioningResponseBodyUserProvisioning) GetTargetType() *string {
	return s.TargetType
}

func (s *GetUserProvisioningResponseBodyUserProvisioning) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *GetUserProvisioningResponseBodyUserProvisioning) GetUserProvisioningId() *string {
	return s.UserProvisioningId
}

func (s *GetUserProvisioningResponseBodyUserProvisioning) SetCreateTime(v string) *GetUserProvisioningResponseBodyUserProvisioning {
	s.CreateTime = &v
	return s
}

func (s *GetUserProvisioningResponseBodyUserProvisioning) SetDeletionStrategy(v string) *GetUserProvisioningResponseBodyUserProvisioning {
	s.DeletionStrategy = &v
	return s
}

func (s *GetUserProvisioningResponseBodyUserProvisioning) SetDescription(v string) *GetUserProvisioningResponseBodyUserProvisioning {
	s.Description = &v
	return s
}

func (s *GetUserProvisioningResponseBodyUserProvisioning) SetDirectoryId(v string) *GetUserProvisioningResponseBodyUserProvisioning {
	s.DirectoryId = &v
	return s
}

func (s *GetUserProvisioningResponseBodyUserProvisioning) SetDuplicationStrategy(v string) *GetUserProvisioningResponseBodyUserProvisioning {
	s.DuplicationStrategy = &v
	return s
}

func (s *GetUserProvisioningResponseBodyUserProvisioning) SetOwnerPk(v string) *GetUserProvisioningResponseBodyUserProvisioning {
	s.OwnerPk = &v
	return s
}

func (s *GetUserProvisioningResponseBodyUserProvisioning) SetPrincipalId(v string) *GetUserProvisioningResponseBodyUserProvisioning {
	s.PrincipalId = &v
	return s
}

func (s *GetUserProvisioningResponseBodyUserProvisioning) SetPrincipalName(v string) *GetUserProvisioningResponseBodyUserProvisioning {
	s.PrincipalName = &v
	return s
}

func (s *GetUserProvisioningResponseBodyUserProvisioning) SetPrincipalType(v string) *GetUserProvisioningResponseBodyUserProvisioning {
	s.PrincipalType = &v
	return s
}

func (s *GetUserProvisioningResponseBodyUserProvisioning) SetStatus(v string) *GetUserProvisioningResponseBodyUserProvisioning {
	s.Status = &v
	return s
}

func (s *GetUserProvisioningResponseBodyUserProvisioning) SetTargetId(v string) *GetUserProvisioningResponseBodyUserProvisioning {
	s.TargetId = &v
	return s
}

func (s *GetUserProvisioningResponseBodyUserProvisioning) SetTargetName(v string) *GetUserProvisioningResponseBodyUserProvisioning {
	s.TargetName = &v
	return s
}

func (s *GetUserProvisioningResponseBodyUserProvisioning) SetTargetPath(v string) *GetUserProvisioningResponseBodyUserProvisioning {
	s.TargetPath = &v
	return s
}

func (s *GetUserProvisioningResponseBodyUserProvisioning) SetTargetType(v string) *GetUserProvisioningResponseBodyUserProvisioning {
	s.TargetType = &v
	return s
}

func (s *GetUserProvisioningResponseBodyUserProvisioning) SetUpdateTime(v string) *GetUserProvisioningResponseBodyUserProvisioning {
	s.UpdateTime = &v
	return s
}

func (s *GetUserProvisioningResponseBodyUserProvisioning) SetUserProvisioningId(v string) *GetUserProvisioningResponseBodyUserProvisioning {
	s.UserProvisioningId = &v
	return s
}

func (s *GetUserProvisioningResponseBodyUserProvisioning) Validate() error {
	return dara.Validate(s)
}
