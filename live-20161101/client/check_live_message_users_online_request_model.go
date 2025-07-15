// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckLiveMessageUsersOnlineRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *CheckLiveMessageUsersOnlineRequest
	GetAppId() *string
	SetDataCenter(v string) *CheckLiveMessageUsersOnlineRequest
	GetDataCenter() *string
	SetUserIds(v []*string) *CheckLiveMessageUsersOnlineRequest
	GetUserIds() []*string
}

type CheckLiveMessageUsersOnlineRequest struct {
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
	UserIds []*string `json:"UserIds,omitempty" xml:"UserIds,omitempty" type:"Repeated"`
}

func (s CheckLiveMessageUsersOnlineRequest) String() string {
	return dara.Prettify(s)
}

func (s CheckLiveMessageUsersOnlineRequest) GoString() string {
	return s.String()
}

func (s *CheckLiveMessageUsersOnlineRequest) GetAppId() *string {
	return s.AppId
}

func (s *CheckLiveMessageUsersOnlineRequest) GetDataCenter() *string {
	return s.DataCenter
}

func (s *CheckLiveMessageUsersOnlineRequest) GetUserIds() []*string {
	return s.UserIds
}

func (s *CheckLiveMessageUsersOnlineRequest) SetAppId(v string) *CheckLiveMessageUsersOnlineRequest {
	s.AppId = &v
	return s
}

func (s *CheckLiveMessageUsersOnlineRequest) SetDataCenter(v string) *CheckLiveMessageUsersOnlineRequest {
	s.DataCenter = &v
	return s
}

func (s *CheckLiveMessageUsersOnlineRequest) SetUserIds(v []*string) *CheckLiveMessageUsersOnlineRequest {
	s.UserIds = v
	return s
}

func (s *CheckLiveMessageUsersOnlineRequest) Validate() error {
	return dara.Validate(s)
}
