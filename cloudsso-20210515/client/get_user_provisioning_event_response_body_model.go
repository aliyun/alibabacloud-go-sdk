// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserProvisioningEventResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetUserProvisioningEventResponseBody
	GetRequestId() *string
	SetUserProvisioningEvent(v *GetUserProvisioningEventResponseBodyUserProvisioningEvent) *GetUserProvisioningEventResponseBody
	GetUserProvisioningEvent() *GetUserProvisioningEventResponseBodyUserProvisioningEvent
}

type GetUserProvisioningEventResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// B801715C-97EA-3067-AC97-EF1EBECBB39C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The RAM user provisioning event.
	UserProvisioningEvent *GetUserProvisioningEventResponseBodyUserProvisioningEvent `json:"UserProvisioningEvent,omitempty" xml:"UserProvisioningEvent,omitempty" type:"Struct"`
}

func (s GetUserProvisioningEventResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetUserProvisioningEventResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserProvisioningEventResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetUserProvisioningEventResponseBody) GetUserProvisioningEvent() *GetUserProvisioningEventResponseBodyUserProvisioningEvent {
	return s.UserProvisioningEvent
}

func (s *GetUserProvisioningEventResponseBody) SetRequestId(v string) *GetUserProvisioningEventResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetUserProvisioningEventResponseBody) SetUserProvisioningEvent(v *GetUserProvisioningEventResponseBodyUserProvisioningEvent) *GetUserProvisioningEventResponseBody {
	s.UserProvisioningEvent = v
	return s
}

func (s *GetUserProvisioningEventResponseBody) Validate() error {
	if s.UserProvisioningEvent != nil {
		if err := s.UserProvisioningEvent.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetUserProvisioningEventResponseBodyUserProvisioningEvent struct {
	// The creation time. The time is displayed in UTC.
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
	// The number of execution failures.
	//
	// example:
	//
	// 3
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
	// 2022-11-28T03:55:42Z
	LatestAsyncTime *string `json:"LatestAsyncTime,omitempty" xml:"LatestAsyncTime,omitempty"`
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
	// exampleGroupName
	PrincipalName *string `json:"PrincipalName,omitempty" xml:"PrincipalName,omitempty"`
	// The identity type of the RAM user provisioning. Valid values:
	//
	// 	- User: The identity of the RAM user provisioning is a CloudSSO user.
	//
	// 	- Group: The identity of the RAM user provisioning is a CloudSSO user group.
	//
	// example:
	//
	// Group
	PrincipalType *string `json:"PrincipalType,omitempty" xml:"PrincipalType,omitempty"`
	// The type of the source operation. Valid values:
	//
	// 	- StartProvisioning: enables the RAM user provisioning.
	//
	// 	- DeleteProvisioning: deletes the RAM user provisioning.
	//
	// 	- AddUserToGroup: adds a user to a user group.
	//
	// 	- RemoveUserFromGroup: removes a user from a user group.
	//
	// 	- UserProvisioningDeletionClearing: deletes the RAM user provisioning and clears resources in the background.
	//
	// example:
	//
	// AddUserToGroup
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// The ID of the object for which you create the RAM user provisioning.
	//
	// The value is fixed as the ID of the member in the resource directory.````
	//
	// example:
	//
	// 1743382******
	TargetId *string `json:"TargetId,omitempty" xml:"TargetId,omitempty"`
	// The name of the object for which you create the RAM user provisioning.
	//
	// The value is fixed as the name of the member in the resource directory.````
	//
	// example:
	//
	// exampleRdMember
	TargetName *string `json:"TargetName,omitempty" xml:"TargetName,omitempty"`
	// The path of the resource directory in which you create the RAM user provisioning for the object.
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
	// The ID of the RAM user provisioning event.
	//
	// example:
	//
	// up-002axzhapcbz6e63****
	UserProvisioningId *string `json:"UserProvisioningId,omitempty" xml:"UserProvisioningId,omitempty"`
}

func (s GetUserProvisioningEventResponseBodyUserProvisioningEvent) String() string {
	return dara.Prettify(s)
}

func (s GetUserProvisioningEventResponseBodyUserProvisioningEvent) GoString() string {
	return s.String()
}

func (s *GetUserProvisioningEventResponseBodyUserProvisioningEvent) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetUserProvisioningEventResponseBodyUserProvisioningEvent) GetDeletionStrategy() *string {
	return s.DeletionStrategy
}

func (s *GetUserProvisioningEventResponseBodyUserProvisioningEvent) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *GetUserProvisioningEventResponseBodyUserProvisioningEvent) GetDuplicationStrategy() *string {
	return s.DuplicationStrategy
}

func (s *GetUserProvisioningEventResponseBodyUserProvisioningEvent) GetErrorCount() *int64 {
	return s.ErrorCount
}

func (s *GetUserProvisioningEventResponseBodyUserProvisioningEvent) GetErrorInfo() *string {
	return s.ErrorInfo
}

func (s *GetUserProvisioningEventResponseBodyUserProvisioningEvent) GetEventId() *string {
	return s.EventId
}

func (s *GetUserProvisioningEventResponseBodyUserProvisioningEvent) GetLatestAsyncTime() *string {
	return s.LatestAsyncTime
}

func (s *GetUserProvisioningEventResponseBodyUserProvisioningEvent) GetPrincipalId() *string {
	return s.PrincipalId
}

func (s *GetUserProvisioningEventResponseBodyUserProvisioningEvent) GetPrincipalName() *string {
	return s.PrincipalName
}

func (s *GetUserProvisioningEventResponseBodyUserProvisioningEvent) GetPrincipalType() *string {
	return s.PrincipalType
}

func (s *GetUserProvisioningEventResponseBodyUserProvisioningEvent) GetSourceType() *string {
	return s.SourceType
}

func (s *GetUserProvisioningEventResponseBodyUserProvisioningEvent) GetTargetId() *string {
	return s.TargetId
}

func (s *GetUserProvisioningEventResponseBodyUserProvisioningEvent) GetTargetName() *string {
	return s.TargetName
}

func (s *GetUserProvisioningEventResponseBodyUserProvisioningEvent) GetTargetPath() *string {
	return s.TargetPath
}

func (s *GetUserProvisioningEventResponseBodyUserProvisioningEvent) GetTargetType() *string {
	return s.TargetType
}

func (s *GetUserProvisioningEventResponseBodyUserProvisioningEvent) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *GetUserProvisioningEventResponseBodyUserProvisioningEvent) GetUserProvisioningId() *string {
	return s.UserProvisioningId
}

func (s *GetUserProvisioningEventResponseBodyUserProvisioningEvent) SetCreateTime(v string) *GetUserProvisioningEventResponseBodyUserProvisioningEvent {
	s.CreateTime = &v
	return s
}

func (s *GetUserProvisioningEventResponseBodyUserProvisioningEvent) SetDeletionStrategy(v string) *GetUserProvisioningEventResponseBodyUserProvisioningEvent {
	s.DeletionStrategy = &v
	return s
}

func (s *GetUserProvisioningEventResponseBodyUserProvisioningEvent) SetDirectoryId(v string) *GetUserProvisioningEventResponseBodyUserProvisioningEvent {
	s.DirectoryId = &v
	return s
}

func (s *GetUserProvisioningEventResponseBodyUserProvisioningEvent) SetDuplicationStrategy(v string) *GetUserProvisioningEventResponseBodyUserProvisioningEvent {
	s.DuplicationStrategy = &v
	return s
}

func (s *GetUserProvisioningEventResponseBodyUserProvisioningEvent) SetErrorCount(v int64) *GetUserProvisioningEventResponseBodyUserProvisioningEvent {
	s.ErrorCount = &v
	return s
}

func (s *GetUserProvisioningEventResponseBodyUserProvisioningEvent) SetErrorInfo(v string) *GetUserProvisioningEventResponseBodyUserProvisioningEvent {
	s.ErrorInfo = &v
	return s
}

func (s *GetUserProvisioningEventResponseBodyUserProvisioningEvent) SetEventId(v string) *GetUserProvisioningEventResponseBodyUserProvisioningEvent {
	s.EventId = &v
	return s
}

func (s *GetUserProvisioningEventResponseBodyUserProvisioningEvent) SetLatestAsyncTime(v string) *GetUserProvisioningEventResponseBodyUserProvisioningEvent {
	s.LatestAsyncTime = &v
	return s
}

func (s *GetUserProvisioningEventResponseBodyUserProvisioningEvent) SetPrincipalId(v string) *GetUserProvisioningEventResponseBodyUserProvisioningEvent {
	s.PrincipalId = &v
	return s
}

func (s *GetUserProvisioningEventResponseBodyUserProvisioningEvent) SetPrincipalName(v string) *GetUserProvisioningEventResponseBodyUserProvisioningEvent {
	s.PrincipalName = &v
	return s
}

func (s *GetUserProvisioningEventResponseBodyUserProvisioningEvent) SetPrincipalType(v string) *GetUserProvisioningEventResponseBodyUserProvisioningEvent {
	s.PrincipalType = &v
	return s
}

func (s *GetUserProvisioningEventResponseBodyUserProvisioningEvent) SetSourceType(v string) *GetUserProvisioningEventResponseBodyUserProvisioningEvent {
	s.SourceType = &v
	return s
}

func (s *GetUserProvisioningEventResponseBodyUserProvisioningEvent) SetTargetId(v string) *GetUserProvisioningEventResponseBodyUserProvisioningEvent {
	s.TargetId = &v
	return s
}

func (s *GetUserProvisioningEventResponseBodyUserProvisioningEvent) SetTargetName(v string) *GetUserProvisioningEventResponseBodyUserProvisioningEvent {
	s.TargetName = &v
	return s
}

func (s *GetUserProvisioningEventResponseBodyUserProvisioningEvent) SetTargetPath(v string) *GetUserProvisioningEventResponseBodyUserProvisioningEvent {
	s.TargetPath = &v
	return s
}

func (s *GetUserProvisioningEventResponseBodyUserProvisioningEvent) SetTargetType(v string) *GetUserProvisioningEventResponseBodyUserProvisioningEvent {
	s.TargetType = &v
	return s
}

func (s *GetUserProvisioningEventResponseBodyUserProvisioningEvent) SetUpdateTime(v string) *GetUserProvisioningEventResponseBodyUserProvisioningEvent {
	s.UpdateTime = &v
	return s
}

func (s *GetUserProvisioningEventResponseBodyUserProvisioningEvent) SetUserProvisioningId(v string) *GetUserProvisioningEventResponseBodyUserProvisioningEvent {
	s.UserProvisioningId = &v
	return s
}

func (s *GetUserProvisioningEventResponseBodyUserProvisioningEvent) Validate() error {
	return dara.Validate(s)
}
