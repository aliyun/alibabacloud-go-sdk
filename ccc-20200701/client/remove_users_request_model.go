// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveUsersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFilePath(v string) *RemoveUsersRequest
	GetFilePath() *string
	SetForce(v bool) *RemoveUsersRequest
	GetForce() *bool
	SetInstanceId(v string) *RemoveUsersRequest
	GetInstanceId() *string
	SetNotificationEmail(v string) *RemoveUsersRequest
	GetNotificationEmail() *string
	SetUserIdList(v string) *RemoveUsersRequest
	GetUserIdList() *string
}

type RemoveUsersRequest struct {
	// The ObjectKey of the OSS object containing the file with the list of users to delete.
	//
	// example:
	//
	// temp/removeUsers.xlsx
	FilePath *string `json:"FilePath,omitempty" xml:"FilePath,omitempty"`
	// Specifies whether to force delete.
	//
	// example:
	//
	// true
	Force *bool `json:"Force,omitempty" xml:"Force,omitempty"`
	// Instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Notification email address.
	//
	// example:
	//
	// 123@abc.com
	NotificationEmail *string `json:"NotificationEmail,omitempty" xml:"NotificationEmail,omitempty"`
	// A JSON-formatted string representing an array of agent IDs. Each array element is an agent ID.
	//
	// example:
	//
	// ["agent1@ccc-test","agent2@ccc-test"]
	UserIdList *string `json:"UserIdList,omitempty" xml:"UserIdList,omitempty"`
}

func (s RemoveUsersRequest) String() string {
	return dara.Prettify(s)
}

func (s RemoveUsersRequest) GoString() string {
	return s.String()
}

func (s *RemoveUsersRequest) GetFilePath() *string {
	return s.FilePath
}

func (s *RemoveUsersRequest) GetForce() *bool {
	return s.Force
}

func (s *RemoveUsersRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *RemoveUsersRequest) GetNotificationEmail() *string {
	return s.NotificationEmail
}

func (s *RemoveUsersRequest) GetUserIdList() *string {
	return s.UserIdList
}

func (s *RemoveUsersRequest) SetFilePath(v string) *RemoveUsersRequest {
	s.FilePath = &v
	return s
}

func (s *RemoveUsersRequest) SetForce(v bool) *RemoveUsersRequest {
	s.Force = &v
	return s
}

func (s *RemoveUsersRequest) SetInstanceId(v string) *RemoveUsersRequest {
	s.InstanceId = &v
	return s
}

func (s *RemoveUsersRequest) SetNotificationEmail(v string) *RemoveUsersRequest {
	s.NotificationEmail = &v
	return s
}

func (s *RemoveUsersRequest) SetUserIdList(v string) *RemoveUsersRequest {
	s.UserIdList = &v
	return s
}

func (s *RemoveUsersRequest) Validate() error {
	return dara.Validate(s)
}
