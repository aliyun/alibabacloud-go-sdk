// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckLiveMessageUsersInGroupShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *CheckLiveMessageUsersInGroupShrinkRequest
	GetAppId() *string
	SetDataCenter(v string) *CheckLiveMessageUsersInGroupShrinkRequest
	GetDataCenter() *string
	SetGroupId(v string) *CheckLiveMessageUsersInGroupShrinkRequest
	GetGroupId() *string
	SetUserIdsShrink(v string) *CheckLiveMessageUsersInGroupShrinkRequest
	GetUserIdsShrink() *string
}

type CheckLiveMessageUsersInGroupShrinkRequest struct {
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
	UserIdsShrink *string `json:"UserIds,omitempty" xml:"UserIds,omitempty"`
}

func (s CheckLiveMessageUsersInGroupShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CheckLiveMessageUsersInGroupShrinkRequest) GoString() string {
	return s.String()
}

func (s *CheckLiveMessageUsersInGroupShrinkRequest) GetAppId() *string {
	return s.AppId
}

func (s *CheckLiveMessageUsersInGroupShrinkRequest) GetDataCenter() *string {
	return s.DataCenter
}

func (s *CheckLiveMessageUsersInGroupShrinkRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *CheckLiveMessageUsersInGroupShrinkRequest) GetUserIdsShrink() *string {
	return s.UserIdsShrink
}

func (s *CheckLiveMessageUsersInGroupShrinkRequest) SetAppId(v string) *CheckLiveMessageUsersInGroupShrinkRequest {
	s.AppId = &v
	return s
}

func (s *CheckLiveMessageUsersInGroupShrinkRequest) SetDataCenter(v string) *CheckLiveMessageUsersInGroupShrinkRequest {
	s.DataCenter = &v
	return s
}

func (s *CheckLiveMessageUsersInGroupShrinkRequest) SetGroupId(v string) *CheckLiveMessageUsersInGroupShrinkRequest {
	s.GroupId = &v
	return s
}

func (s *CheckLiveMessageUsersInGroupShrinkRequest) SetUserIdsShrink(v string) *CheckLiveMessageUsersInGroupShrinkRequest {
	s.UserIdsShrink = &v
	return s
}

func (s *CheckLiveMessageUsersInGroupShrinkRequest) Validate() error {
	return dara.Validate(s)
}
