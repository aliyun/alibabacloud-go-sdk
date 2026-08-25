// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUserProvisioningEventsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetIsTruncated(v bool) *ListUserProvisioningEventsResponseBody
	GetIsTruncated() *bool
	SetMaxResults(v int32) *ListUserProvisioningEventsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListUserProvisioningEventsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListUserProvisioningEventsResponseBody
	GetRequestId() *string
	SetTotalCounts(v int32) *ListUserProvisioningEventsResponseBody
	GetTotalCounts() *int32
	SetUserProvisioningEvents(v []*ListUserProvisioningEventsResponseBodyUserProvisioningEvents) *ListUserProvisioningEventsResponseBody
	GetUserProvisioningEvents() []*ListUserProvisioningEventsResponseBodyUserProvisioningEvents
}

type ListUserProvisioningEventsResponseBody struct {
	// Indicates whether the queried entries are truncated. Valid values:
	//
	// - true
	//
	// - false
	//
	// example:
	//
	// false
	IsTruncated *bool `json:"IsTruncated,omitempty" xml:"IsTruncated,omitempty"`
	// The maximum number of entries per page.
	//
	// Valid values: 1 to 100.
	//
	// Default value: 10.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that is used to initiate the next request.
	//
	// > This parameter is returned only when the `IsTruncated` parameter is set to `true`.
	//
	// example:
	//
	// 2eEMmhmLa1b7Bbj9UzCgZUGj8DpDeG5TbNknuNKNP2h84KjJRnAb7vzzSDkYNmsidnAybyJYBfnPPB6xfgw54B1Wub2KQmC8LofzqBW2Y****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0D85B43D-EF98-396D-B426-837E428D2D39
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 110
	TotalCounts *int32 `json:"TotalCounts,omitempty" xml:"TotalCounts,omitempty"`
	// The RAM user provisioning events.
	UserProvisioningEvents []*ListUserProvisioningEventsResponseBodyUserProvisioningEvents `json:"UserProvisioningEvents,omitempty" xml:"UserProvisioningEvents,omitempty" type:"Repeated"`
}

func (s ListUserProvisioningEventsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListUserProvisioningEventsResponseBody) GoString() string {
	return s.String()
}

func (s *ListUserProvisioningEventsResponseBody) GetIsTruncated() *bool {
	return s.IsTruncated
}

func (s *ListUserProvisioningEventsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListUserProvisioningEventsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListUserProvisioningEventsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListUserProvisioningEventsResponseBody) GetTotalCounts() *int32 {
	return s.TotalCounts
}

func (s *ListUserProvisioningEventsResponseBody) GetUserProvisioningEvents() []*ListUserProvisioningEventsResponseBodyUserProvisioningEvents {
	return s.UserProvisioningEvents
}

func (s *ListUserProvisioningEventsResponseBody) SetIsTruncated(v bool) *ListUserProvisioningEventsResponseBody {
	s.IsTruncated = &v
	return s
}

func (s *ListUserProvisioningEventsResponseBody) SetMaxResults(v int32) *ListUserProvisioningEventsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListUserProvisioningEventsResponseBody) SetNextToken(v string) *ListUserProvisioningEventsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListUserProvisioningEventsResponseBody) SetRequestId(v string) *ListUserProvisioningEventsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListUserProvisioningEventsResponseBody) SetTotalCounts(v int32) *ListUserProvisioningEventsResponseBody {
	s.TotalCounts = &v
	return s
}

func (s *ListUserProvisioningEventsResponseBody) SetUserProvisioningEvents(v []*ListUserProvisioningEventsResponseBodyUserProvisioningEvents) *ListUserProvisioningEventsResponseBody {
	s.UserProvisioningEvents = v
	return s
}

func (s *ListUserProvisioningEventsResponseBody) Validate() error {
	if s.UserProvisioningEvents != nil {
		for _, item := range s.UserProvisioningEvents {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListUserProvisioningEventsResponseBodyUserProvisioningEvents struct {
	// The creation time. The time is displayed in UTC.
	//
	// example:
	//
	// 2022-11-28T03:55:55Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The deletion policy. The policy is used to manage synchronized users when you delete the RAM user provisioning. Valid values:
	//
	// - Delete: When you delete the RAM user provisioning, the system deletes the synchronized users.
	//
	// - Keep: When you delete the RAM user provisioning, the system retains the synchronized users.
	//
	// example:
	//
	// Keep
	DeletionStrategy *string `json:"DeletionStrategy,omitempty" xml:"DeletionStrategy,omitempty"`
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
	// The number of execution failures.
	//
	// example:
	//
	// 1
	ErrorCount *int64 `json:"ErrorCount,omitempty" xml:"ErrorCount,omitempty"`
	// The error message that is displayed when the last execution of the RAM user provisioning event failed.
	//
	// example:
	//
	// OperationConflict.UserProvisioning.Process.fail.ImsUserExists
	ErrorInfo *string `json:"ErrorInfo,omitempty" xml:"ErrorInfo,omitempty"`
	// The ID of the RAM user provisioning event.
	//
	// example:
	//
	// upe-wjKyNDmZvyZOiRcJ****
	EventId *string `json:"EventId,omitempty" xml:"EventId,omitempty"`
	// The time at which the RAM user provisioning event was last executed. The time is displayed in UTC.
	//
	// example:
	//
	// 2022-11-28T03:55:55Z
	LatestAsyncTime *string `json:"LatestAsyncTime,omitempty" xml:"LatestAsyncTime,omitempty"`
	// The identity ID of the RAM user provisioning. Valid values:
	//
	// - If you set the `PrincipalType` parameter to `Group`, the value of this parameter is the ID of a CloudSSO user group (g-\\*\\*\\*\\*\\*\\*\\*\\*).
	//
	// - If you set the `PrincipalType` parameter to `User`, the value of this parameter is the ID of a CloudSSO user (u-\\*\\*\\*\\*\\*\\*\\*\\*).
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
	// exampleGroupName
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
	// The type of the source operation. Valid values:
	//
	// - StartProvisioning: enables the RAM user provisioning.
	//
	// - DeleteProvisioning: deletes the RAM user provisioning.
	//
	// - AddUserToGroup: adds a user to a user group.
	//
	// - RemoveUserFromGroup: removes a user from a user group.
	//
	// - UserProvisioningDeletionClearing: deletes the RAM user provisioning and clears resources in the background.
	//
	// example:
	//
	// StartProvisioning
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// The ID of the object for which you create the RAM user provisioning.
	//
	// The value is fixed as the ID of the account in the resource directory.\\`\\`\\`\\`
	//
	// example:
	//
	// 153218*******
	TargetId *string `json:"TargetId,omitempty" xml:"TargetId,omitempty"`
	// The name of the object for which you create the RAM user provisioning.
	//
	// If `RD-Account` is returned, the value of this parameter is the name of the account that is used to access the instance.\\`\\`
	//
	// example:
	//
	// exampleRdMember
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
	// 2022-11-28T03:55:55Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// The ID of the RAM user provisioning.
	//
	// example:
	//
	// up-002axzhapcbz6e63****
	UserProvisioningId *string `json:"UserProvisioningId,omitempty" xml:"UserProvisioningId,omitempty"`
}

func (s ListUserProvisioningEventsResponseBodyUserProvisioningEvents) String() string {
	return dara.Prettify(s)
}

func (s ListUserProvisioningEventsResponseBodyUserProvisioningEvents) GoString() string {
	return s.String()
}

func (s *ListUserProvisioningEventsResponseBodyUserProvisioningEvents) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListUserProvisioningEventsResponseBodyUserProvisioningEvents) GetDeletionStrategy() *string {
	return s.DeletionStrategy
}

func (s *ListUserProvisioningEventsResponseBodyUserProvisioningEvents) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *ListUserProvisioningEventsResponseBodyUserProvisioningEvents) GetDuplicationStrategy() *string {
	return s.DuplicationStrategy
}

func (s *ListUserProvisioningEventsResponseBodyUserProvisioningEvents) GetErrorCount() *int64 {
	return s.ErrorCount
}

func (s *ListUserProvisioningEventsResponseBodyUserProvisioningEvents) GetErrorInfo() *string {
	return s.ErrorInfo
}

func (s *ListUserProvisioningEventsResponseBodyUserProvisioningEvents) GetEventId() *string {
	return s.EventId
}

func (s *ListUserProvisioningEventsResponseBodyUserProvisioningEvents) GetLatestAsyncTime() *string {
	return s.LatestAsyncTime
}

func (s *ListUserProvisioningEventsResponseBodyUserProvisioningEvents) GetPrincipalId() *string {
	return s.PrincipalId
}

func (s *ListUserProvisioningEventsResponseBodyUserProvisioningEvents) GetPrincipalName() *string {
	return s.PrincipalName
}

func (s *ListUserProvisioningEventsResponseBodyUserProvisioningEvents) GetPrincipalType() *string {
	return s.PrincipalType
}

func (s *ListUserProvisioningEventsResponseBodyUserProvisioningEvents) GetSourceType() *string {
	return s.SourceType
}

func (s *ListUserProvisioningEventsResponseBodyUserProvisioningEvents) GetTargetId() *string {
	return s.TargetId
}

func (s *ListUserProvisioningEventsResponseBodyUserProvisioningEvents) GetTargetName() *string {
	return s.TargetName
}

func (s *ListUserProvisioningEventsResponseBodyUserProvisioningEvents) GetTargetPath() *string {
	return s.TargetPath
}

func (s *ListUserProvisioningEventsResponseBodyUserProvisioningEvents) GetTargetType() *string {
	return s.TargetType
}

func (s *ListUserProvisioningEventsResponseBodyUserProvisioningEvents) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *ListUserProvisioningEventsResponseBodyUserProvisioningEvents) GetUserProvisioningId() *string {
	return s.UserProvisioningId
}

func (s *ListUserProvisioningEventsResponseBodyUserProvisioningEvents) SetCreateTime(v string) *ListUserProvisioningEventsResponseBodyUserProvisioningEvents {
	s.CreateTime = &v
	return s
}

func (s *ListUserProvisioningEventsResponseBodyUserProvisioningEvents) SetDeletionStrategy(v string) *ListUserProvisioningEventsResponseBodyUserProvisioningEvents {
	s.DeletionStrategy = &v
	return s
}

func (s *ListUserProvisioningEventsResponseBodyUserProvisioningEvents) SetDirectoryId(v string) *ListUserProvisioningEventsResponseBodyUserProvisioningEvents {
	s.DirectoryId = &v
	return s
}

func (s *ListUserProvisioningEventsResponseBodyUserProvisioningEvents) SetDuplicationStrategy(v string) *ListUserProvisioningEventsResponseBodyUserProvisioningEvents {
	s.DuplicationStrategy = &v
	return s
}

func (s *ListUserProvisioningEventsResponseBodyUserProvisioningEvents) SetErrorCount(v int64) *ListUserProvisioningEventsResponseBodyUserProvisioningEvents {
	s.ErrorCount = &v
	return s
}

func (s *ListUserProvisioningEventsResponseBodyUserProvisioningEvents) SetErrorInfo(v string) *ListUserProvisioningEventsResponseBodyUserProvisioningEvents {
	s.ErrorInfo = &v
	return s
}

func (s *ListUserProvisioningEventsResponseBodyUserProvisioningEvents) SetEventId(v string) *ListUserProvisioningEventsResponseBodyUserProvisioningEvents {
	s.EventId = &v
	return s
}

func (s *ListUserProvisioningEventsResponseBodyUserProvisioningEvents) SetLatestAsyncTime(v string) *ListUserProvisioningEventsResponseBodyUserProvisioningEvents {
	s.LatestAsyncTime = &v
	return s
}

func (s *ListUserProvisioningEventsResponseBodyUserProvisioningEvents) SetPrincipalId(v string) *ListUserProvisioningEventsResponseBodyUserProvisioningEvents {
	s.PrincipalId = &v
	return s
}

func (s *ListUserProvisioningEventsResponseBodyUserProvisioningEvents) SetPrincipalName(v string) *ListUserProvisioningEventsResponseBodyUserProvisioningEvents {
	s.PrincipalName = &v
	return s
}

func (s *ListUserProvisioningEventsResponseBodyUserProvisioningEvents) SetPrincipalType(v string) *ListUserProvisioningEventsResponseBodyUserProvisioningEvents {
	s.PrincipalType = &v
	return s
}

func (s *ListUserProvisioningEventsResponseBodyUserProvisioningEvents) SetSourceType(v string) *ListUserProvisioningEventsResponseBodyUserProvisioningEvents {
	s.SourceType = &v
	return s
}

func (s *ListUserProvisioningEventsResponseBodyUserProvisioningEvents) SetTargetId(v string) *ListUserProvisioningEventsResponseBodyUserProvisioningEvents {
	s.TargetId = &v
	return s
}

func (s *ListUserProvisioningEventsResponseBodyUserProvisioningEvents) SetTargetName(v string) *ListUserProvisioningEventsResponseBodyUserProvisioningEvents {
	s.TargetName = &v
	return s
}

func (s *ListUserProvisioningEventsResponseBodyUserProvisioningEvents) SetTargetPath(v string) *ListUserProvisioningEventsResponseBodyUserProvisioningEvents {
	s.TargetPath = &v
	return s
}

func (s *ListUserProvisioningEventsResponseBodyUserProvisioningEvents) SetTargetType(v string) *ListUserProvisioningEventsResponseBodyUserProvisioningEvents {
	s.TargetType = &v
	return s
}

func (s *ListUserProvisioningEventsResponseBodyUserProvisioningEvents) SetUpdateTime(v string) *ListUserProvisioningEventsResponseBodyUserProvisioningEvents {
	s.UpdateTime = &v
	return s
}

func (s *ListUserProvisioningEventsResponseBodyUserProvisioningEvents) SetUserProvisioningId(v string) *ListUserProvisioningEventsResponseBodyUserProvisioningEvents {
	s.UserProvisioningId = &v
	return s
}

func (s *ListUserProvisioningEventsResponseBodyUserProvisioningEvents) Validate() error {
	return dara.Validate(s)
}
