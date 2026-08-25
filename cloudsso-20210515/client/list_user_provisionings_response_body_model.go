// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUserProvisioningsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetIsTruncated(v bool) *ListUserProvisioningsResponseBody
	GetIsTruncated() *bool
	SetMaxResults(v int32) *ListUserProvisioningsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListUserProvisioningsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListUserProvisioningsResponseBody
	GetRequestId() *string
	SetTotalCounts(v int32) *ListUserProvisioningsResponseBody
	GetTotalCounts() *int32
	SetUserProvisionings(v []*ListUserProvisioningsResponseBodyUserProvisionings) *ListUserProvisioningsResponseBody
	GetUserProvisionings() []*ListUserProvisioningsResponseBodyUserProvisionings
}

type ListUserProvisioningsResponseBody struct {
	// Indicates whether the queried entries are truncated. Valid values:
	//
	// - true
	//
	// - false
	//
	// example:
	//
	// true
	IsTruncated *bool `json:"IsTruncated,omitempty" xml:"IsTruncated,omitempty"`
	// The maximum number of entries per page.
	//
	// Valid values: 1 to 100.
	//
	// Default value: 10.
	//
	// example:
	//
	// 100
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that is used to initiate the next request.
	//
	// > This parameter is returned only when the `IsTruncated` parameter is set to `true`.
	//
	// example:
	//
	// 27EbL9j4ZgZjsMZFqbZFgbwQ1VXFU1Khcpx9e2vrW1zwzTBmTGWaM7ixHhRin8SCsxaJdazYVCzeKc2UF2QkyGb83cPhr8ZxrzoaiTd****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// F76AF4FC-****-****-B7CB-74F3********
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 110
	TotalCounts *int32 `json:"TotalCounts,omitempty" xml:"TotalCounts,omitempty"`
	// The RAM user provisionings.
	UserProvisionings []*ListUserProvisioningsResponseBodyUserProvisionings `json:"UserProvisionings,omitempty" xml:"UserProvisionings,omitempty" type:"Repeated"`
}

func (s ListUserProvisioningsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListUserProvisioningsResponseBody) GoString() string {
	return s.String()
}

func (s *ListUserProvisioningsResponseBody) GetIsTruncated() *bool {
	return s.IsTruncated
}

func (s *ListUserProvisioningsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListUserProvisioningsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListUserProvisioningsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListUserProvisioningsResponseBody) GetTotalCounts() *int32 {
	return s.TotalCounts
}

func (s *ListUserProvisioningsResponseBody) GetUserProvisionings() []*ListUserProvisioningsResponseBodyUserProvisionings {
	return s.UserProvisionings
}

func (s *ListUserProvisioningsResponseBody) SetIsTruncated(v bool) *ListUserProvisioningsResponseBody {
	s.IsTruncated = &v
	return s
}

func (s *ListUserProvisioningsResponseBody) SetMaxResults(v int32) *ListUserProvisioningsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListUserProvisioningsResponseBody) SetNextToken(v string) *ListUserProvisioningsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListUserProvisioningsResponseBody) SetRequestId(v string) *ListUserProvisioningsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListUserProvisioningsResponseBody) SetTotalCounts(v int32) *ListUserProvisioningsResponseBody {
	s.TotalCounts = &v
	return s
}

func (s *ListUserProvisioningsResponseBody) SetUserProvisionings(v []*ListUserProvisioningsResponseBodyUserProvisionings) *ListUserProvisioningsResponseBody {
	s.UserProvisionings = v
	return s
}

func (s *ListUserProvisioningsResponseBody) Validate() error {
	if s.UserProvisionings != nil {
		for _, item := range s.UserProvisionings {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListUserProvisioningsResponseBodyUserProvisionings struct {
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
	// The description.
	//
	// example:
	//
	// this is a user provisioning.
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
	// The object for which you create the RAM user provisioning. The value is fixed as `RD-Account`.
	//
	// example:
	//
	// testRdMember
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

func (s ListUserProvisioningsResponseBodyUserProvisionings) String() string {
	return dara.Prettify(s)
}

func (s ListUserProvisioningsResponseBodyUserProvisionings) GoString() string {
	return s.String()
}

func (s *ListUserProvisioningsResponseBodyUserProvisionings) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListUserProvisioningsResponseBodyUserProvisionings) GetDeletionStrategy() *string {
	return s.DeletionStrategy
}

func (s *ListUserProvisioningsResponseBodyUserProvisionings) GetDescription() *string {
	return s.Description
}

func (s *ListUserProvisioningsResponseBodyUserProvisionings) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *ListUserProvisioningsResponseBodyUserProvisionings) GetDuplicationStrategy() *string {
	return s.DuplicationStrategy
}

func (s *ListUserProvisioningsResponseBodyUserProvisionings) GetOwnerPk() *string {
	return s.OwnerPk
}

func (s *ListUserProvisioningsResponseBodyUserProvisionings) GetPrincipalId() *string {
	return s.PrincipalId
}

func (s *ListUserProvisioningsResponseBodyUserProvisionings) GetPrincipalName() *string {
	return s.PrincipalName
}

func (s *ListUserProvisioningsResponseBodyUserProvisionings) GetPrincipalType() *string {
	return s.PrincipalType
}

func (s *ListUserProvisioningsResponseBodyUserProvisionings) GetStatus() *string {
	return s.Status
}

func (s *ListUserProvisioningsResponseBodyUserProvisionings) GetTargetId() *string {
	return s.TargetId
}

func (s *ListUserProvisioningsResponseBodyUserProvisionings) GetTargetName() *string {
	return s.TargetName
}

func (s *ListUserProvisioningsResponseBodyUserProvisionings) GetTargetPath() *string {
	return s.TargetPath
}

func (s *ListUserProvisioningsResponseBodyUserProvisionings) GetTargetType() *string {
	return s.TargetType
}

func (s *ListUserProvisioningsResponseBodyUserProvisionings) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *ListUserProvisioningsResponseBodyUserProvisionings) GetUserProvisioningId() *string {
	return s.UserProvisioningId
}

func (s *ListUserProvisioningsResponseBodyUserProvisionings) SetCreateTime(v string) *ListUserProvisioningsResponseBodyUserProvisionings {
	s.CreateTime = &v
	return s
}

func (s *ListUserProvisioningsResponseBodyUserProvisionings) SetDeletionStrategy(v string) *ListUserProvisioningsResponseBodyUserProvisionings {
	s.DeletionStrategy = &v
	return s
}

func (s *ListUserProvisioningsResponseBodyUserProvisionings) SetDescription(v string) *ListUserProvisioningsResponseBodyUserProvisionings {
	s.Description = &v
	return s
}

func (s *ListUserProvisioningsResponseBodyUserProvisionings) SetDirectoryId(v string) *ListUserProvisioningsResponseBodyUserProvisionings {
	s.DirectoryId = &v
	return s
}

func (s *ListUserProvisioningsResponseBodyUserProvisionings) SetDuplicationStrategy(v string) *ListUserProvisioningsResponseBodyUserProvisionings {
	s.DuplicationStrategy = &v
	return s
}

func (s *ListUserProvisioningsResponseBodyUserProvisionings) SetOwnerPk(v string) *ListUserProvisioningsResponseBodyUserProvisionings {
	s.OwnerPk = &v
	return s
}

func (s *ListUserProvisioningsResponseBodyUserProvisionings) SetPrincipalId(v string) *ListUserProvisioningsResponseBodyUserProvisionings {
	s.PrincipalId = &v
	return s
}

func (s *ListUserProvisioningsResponseBodyUserProvisionings) SetPrincipalName(v string) *ListUserProvisioningsResponseBodyUserProvisionings {
	s.PrincipalName = &v
	return s
}

func (s *ListUserProvisioningsResponseBodyUserProvisionings) SetPrincipalType(v string) *ListUserProvisioningsResponseBodyUserProvisionings {
	s.PrincipalType = &v
	return s
}

func (s *ListUserProvisioningsResponseBodyUserProvisionings) SetStatus(v string) *ListUserProvisioningsResponseBodyUserProvisionings {
	s.Status = &v
	return s
}

func (s *ListUserProvisioningsResponseBodyUserProvisionings) SetTargetId(v string) *ListUserProvisioningsResponseBodyUserProvisionings {
	s.TargetId = &v
	return s
}

func (s *ListUserProvisioningsResponseBodyUserProvisionings) SetTargetName(v string) *ListUserProvisioningsResponseBodyUserProvisionings {
	s.TargetName = &v
	return s
}

func (s *ListUserProvisioningsResponseBodyUserProvisionings) SetTargetPath(v string) *ListUserProvisioningsResponseBodyUserProvisionings {
	s.TargetPath = &v
	return s
}

func (s *ListUserProvisioningsResponseBodyUserProvisionings) SetTargetType(v string) *ListUserProvisioningsResponseBodyUserProvisionings {
	s.TargetType = &v
	return s
}

func (s *ListUserProvisioningsResponseBodyUserProvisionings) SetUpdateTime(v string) *ListUserProvisioningsResponseBodyUserProvisionings {
	s.UpdateTime = &v
	return s
}

func (s *ListUserProvisioningsResponseBodyUserProvisionings) SetUserProvisioningId(v string) *ListUserProvisioningsResponseBodyUserProvisionings {
	s.UserProvisioningId = &v
	return s
}

func (s *ListUserProvisioningsResponseBodyUserProvisionings) Validate() error {
	return dara.Validate(s)
}
