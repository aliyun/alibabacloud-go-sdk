// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveUsersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *RemoveUsersRequest
	GetAppId() *string
	SetChannelId(v string) *RemoveUsersRequest
	GetChannelId() *string
	SetUsers(v []*RemoveUsersRequestUsers) *RemoveUsersRequest
	GetUsers() []*RemoveUsersRequestUsers
}

type RemoveUsersRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// eo85****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// testid
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	// This parameter is required.
	Users []*RemoveUsersRequestUsers `json:"Users,omitempty" xml:"Users,omitempty" type:"Repeated"`
}

func (s RemoveUsersRequest) String() string {
	return dara.Prettify(s)
}

func (s RemoveUsersRequest) GoString() string {
	return s.String()
}

func (s *RemoveUsersRequest) GetAppId() *string {
	return s.AppId
}

func (s *RemoveUsersRequest) GetChannelId() *string {
	return s.ChannelId
}

func (s *RemoveUsersRequest) GetUsers() []*RemoveUsersRequestUsers {
	return s.Users
}

func (s *RemoveUsersRequest) SetAppId(v string) *RemoveUsersRequest {
	s.AppId = &v
	return s
}

func (s *RemoveUsersRequest) SetChannelId(v string) *RemoveUsersRequest {
	s.ChannelId = &v
	return s
}

func (s *RemoveUsersRequest) SetUsers(v []*RemoveUsersRequestUsers) *RemoveUsersRequest {
	s.Users = v
	return s
}

func (s *RemoveUsersRequest) Validate() error {
	return dara.Validate(s)
}

type RemoveUsersRequestUsers struct {
	// This parameter is required.
	//
	// example:
	//
	// 1811****
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s RemoveUsersRequestUsers) String() string {
	return dara.Prettify(s)
}

func (s RemoveUsersRequestUsers) GoString() string {
	return s.String()
}

func (s *RemoveUsersRequestUsers) GetUserId() *string {
	return s.UserId
}

func (s *RemoveUsersRequestUsers) SetUserId(v string) *RemoveUsersRequestUsers {
	s.UserId = &v
	return s
}

func (s *RemoveUsersRequestUsers) Validate() error {
	return dara.Validate(s)
}
