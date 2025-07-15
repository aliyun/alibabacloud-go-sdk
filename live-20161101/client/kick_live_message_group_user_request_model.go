// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iKickLiveMessageGroupUserRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *KickLiveMessageGroupUserRequest
	GetAppId() *string
	SetDataCenter(v string) *KickLiveMessageGroupUserRequest
	GetDataCenter() *string
	SetGroupId(v string) *KickLiveMessageGroupUserRequest
	GetGroupId() *string
	SetKickoffUser(v string) *KickLiveMessageGroupUserRequest
	GetKickoffUser() *string
}

type KickLiveMessageGroupUserRequest struct {
	// The ID of the interactive messaging application to which the interactive messaging group belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// coims-****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The data center. It must be the same as the data center that was specified when you called the [CreateLiveMessageApp](https://help.aliyun.com/document_detail/2848162.html) operation to create the interactive messaging application. Valid values: cn-shanghai and ap-southeast-1 (Singapore).
	//
	// example:
	//
	// cn-shanghai
	DataCenter *string `json:"DataCenter,omitempty" xml:"DataCenter,omitempty"`
	// The ID of the interactive messaging group from which you want to remove the user.
	//
	// This parameter is required.
	//
	// example:
	//
	// grouptest
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The user that you want to remove.
	//
	// This parameter is required.
	//
	// example:
	//
	// uid1
	KickoffUser *string `json:"KickoffUser,omitempty" xml:"KickoffUser,omitempty"`
}

func (s KickLiveMessageGroupUserRequest) String() string {
	return dara.Prettify(s)
}

func (s KickLiveMessageGroupUserRequest) GoString() string {
	return s.String()
}

func (s *KickLiveMessageGroupUserRequest) GetAppId() *string {
	return s.AppId
}

func (s *KickLiveMessageGroupUserRequest) GetDataCenter() *string {
	return s.DataCenter
}

func (s *KickLiveMessageGroupUserRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *KickLiveMessageGroupUserRequest) GetKickoffUser() *string {
	return s.KickoffUser
}

func (s *KickLiveMessageGroupUserRequest) SetAppId(v string) *KickLiveMessageGroupUserRequest {
	s.AppId = &v
	return s
}

func (s *KickLiveMessageGroupUserRequest) SetDataCenter(v string) *KickLiveMessageGroupUserRequest {
	s.DataCenter = &v
	return s
}

func (s *KickLiveMessageGroupUserRequest) SetGroupId(v string) *KickLiveMessageGroupUserRequest {
	s.GroupId = &v
	return s
}

func (s *KickLiveMessageGroupUserRequest) SetKickoffUser(v string) *KickLiveMessageGroupUserRequest {
	s.KickoffUser = &v
	return s
}

func (s *KickLiveMessageGroupUserRequest) Validate() error {
	return dara.Validate(s)
}
