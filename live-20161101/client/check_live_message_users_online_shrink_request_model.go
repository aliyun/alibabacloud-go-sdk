// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckLiveMessageUsersOnlineShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *CheckLiveMessageUsersOnlineShrinkRequest
	GetAppId() *string
	SetDataCenter(v string) *CheckLiveMessageUsersOnlineShrinkRequest
	GetDataCenter() *string
	SetUserIdsShrink(v string) *CheckLiveMessageUsersOnlineShrinkRequest
	GetUserIdsShrink() *string
}

type CheckLiveMessageUsersOnlineShrinkRequest struct {
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
	// The list of users that you want to query.
	//
	// This parameter is required.
	UserIdsShrink *string `json:"UserIds,omitempty" xml:"UserIds,omitempty"`
}

func (s CheckLiveMessageUsersOnlineShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CheckLiveMessageUsersOnlineShrinkRequest) GoString() string {
	return s.String()
}

func (s *CheckLiveMessageUsersOnlineShrinkRequest) GetAppId() *string {
	return s.AppId
}

func (s *CheckLiveMessageUsersOnlineShrinkRequest) GetDataCenter() *string {
	return s.DataCenter
}

func (s *CheckLiveMessageUsersOnlineShrinkRequest) GetUserIdsShrink() *string {
	return s.UserIdsShrink
}

func (s *CheckLiveMessageUsersOnlineShrinkRequest) SetAppId(v string) *CheckLiveMessageUsersOnlineShrinkRequest {
	s.AppId = &v
	return s
}

func (s *CheckLiveMessageUsersOnlineShrinkRequest) SetDataCenter(v string) *CheckLiveMessageUsersOnlineShrinkRequest {
	s.DataCenter = &v
	return s
}

func (s *CheckLiveMessageUsersOnlineShrinkRequest) SetUserIdsShrink(v string) *CheckLiveMessageUsersOnlineShrinkRequest {
	s.UserIdsShrink = &v
	return s
}

func (s *CheckLiveMessageUsersOnlineShrinkRequest) Validate() error {
	return dara.Validate(s)
}
