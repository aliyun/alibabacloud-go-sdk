// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckLiveMessageUsersInGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *CheckLiveMessageUsersInGroupRequest
	GetAppId() *string
	SetDataCenter(v string) *CheckLiveMessageUsersInGroupRequest
	GetDataCenter() *string
	SetGroupId(v string) *CheckLiveMessageUsersInGroupRequest
	GetGroupId() *string
	SetUserIds(v []*string) *CheckLiveMessageUsersInGroupRequest
	GetUserIds() []*string
}

type CheckLiveMessageUsersInGroupRequest struct {
	// The application ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// demo
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The data center. It must be the same as the data center that was specified when you called the [CreateLiveMessageApp](https://help.aliyun.com/document_detail/2848162.html) operation to create the interactive messaging application. Valid values: cn-shanghai and ap-southeast-1 (Singapore).
	//
	// example:
	//
	// cn-shanghai
	DataCenter *string `json:"DataCenter,omitempty" xml:"DataCenter,omitempty"`
	// The group ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// grouptest
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The list of users that you want to query.
	//
	// This parameter is required.
	UserIds []*string `json:"UserIds,omitempty" xml:"UserIds,omitempty" type:"Repeated"`
}

func (s CheckLiveMessageUsersInGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s CheckLiveMessageUsersInGroupRequest) GoString() string {
	return s.String()
}

func (s *CheckLiveMessageUsersInGroupRequest) GetAppId() *string {
	return s.AppId
}

func (s *CheckLiveMessageUsersInGroupRequest) GetDataCenter() *string {
	return s.DataCenter
}

func (s *CheckLiveMessageUsersInGroupRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *CheckLiveMessageUsersInGroupRequest) GetUserIds() []*string {
	return s.UserIds
}

func (s *CheckLiveMessageUsersInGroupRequest) SetAppId(v string) *CheckLiveMessageUsersInGroupRequest {
	s.AppId = &v
	return s
}

func (s *CheckLiveMessageUsersInGroupRequest) SetDataCenter(v string) *CheckLiveMessageUsersInGroupRequest {
	s.DataCenter = &v
	return s
}

func (s *CheckLiveMessageUsersInGroupRequest) SetGroupId(v string) *CheckLiveMessageUsersInGroupRequest {
	s.GroupId = &v
	return s
}

func (s *CheckLiveMessageUsersInGroupRequest) SetUserIds(v []*string) *CheckLiveMessageUsersInGroupRequest {
	s.UserIds = v
	return s
}

func (s *CheckLiveMessageUsersInGroupRequest) Validate() error {
	return dara.Validate(s)
}
