// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartChatRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessChannelId(v string) *StartChatRequest
	GetAccessChannelId() *string
	SetInstanceId(v string) *StartChatRequest
	GetInstanceId() *string
	SetToken(v string) *StartChatRequest
	GetToken() *string
	SetUserList(v []*StartChatRequestUserList) *StartChatRequest
	GetUserList() []*StartChatRequestUserList
}

type StartChatRequest struct {
	// example:
	//
	// cf584733-***-***-9699-cb77aa3b7aa6
	AccessChannelId *string `json:"AccessChannelId,omitempty" xml:"AccessChannelId,omitempty"`
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 9XYGTGWtq2wFi_Bpg7aUnIoYi_vG_rO3bjEn0YtsxbHRHrYHlz1LDBLJAyZcLxieRQR4h_6AnWvTjJeNU5jg************Hwej7WgWrmA=
	Token *string `json:"Token,omitempty" xml:"Token,omitempty"`
	// This parameter is required.
	UserList []*StartChatRequestUserList `json:"UserList,omitempty" xml:"UserList,omitempty" type:"Repeated"`
}

func (s StartChatRequest) String() string {
	return dara.Prettify(s)
}

func (s StartChatRequest) GoString() string {
	return s.String()
}

func (s *StartChatRequest) GetAccessChannelId() *string {
	return s.AccessChannelId
}

func (s *StartChatRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *StartChatRequest) GetToken() *string {
	return s.Token
}

func (s *StartChatRequest) GetUserList() []*StartChatRequestUserList {
	return s.UserList
}

func (s *StartChatRequest) SetAccessChannelId(v string) *StartChatRequest {
	s.AccessChannelId = &v
	return s
}

func (s *StartChatRequest) SetInstanceId(v string) *StartChatRequest {
	s.InstanceId = &v
	return s
}

func (s *StartChatRequest) SetToken(v string) *StartChatRequest {
	s.Token = &v
	return s
}

func (s *StartChatRequest) SetUserList(v []*StartChatRequestUserList) *StartChatRequest {
	s.UserList = v
	return s
}

func (s *StartChatRequest) Validate() error {
	return dara.Validate(s)
}

type StartChatRequestUserList struct {
	// example:
	//
	// http://xxx.com/image
	AvatarUrl *string `json:"AvatarUrl,omitempty" xml:"AvatarUrl,omitempty"`
	Nickname  *string `json:"Nickname,omitempty" xml:"Nickname,omitempty"`
	// example:
	//
	// fcd020fe-d8e4-40e5-8c77-1a272a174a7d
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// example:
	//
	// CUSTOMER
	UserType *string `json:"UserType,omitempty" xml:"UserType,omitempty"`
}

func (s StartChatRequestUserList) String() string {
	return dara.Prettify(s)
}

func (s StartChatRequestUserList) GoString() string {
	return s.String()
}

func (s *StartChatRequestUserList) GetAvatarUrl() *string {
	return s.AvatarUrl
}

func (s *StartChatRequestUserList) GetNickname() *string {
	return s.Nickname
}

func (s *StartChatRequestUserList) GetUserId() *string {
	return s.UserId
}

func (s *StartChatRequestUserList) GetUserType() *string {
	return s.UserType
}

func (s *StartChatRequestUserList) SetAvatarUrl(v string) *StartChatRequestUserList {
	s.AvatarUrl = &v
	return s
}

func (s *StartChatRequestUserList) SetNickname(v string) *StartChatRequestUserList {
	s.Nickname = &v
	return s
}

func (s *StartChatRequestUserList) SetUserId(v string) *StartChatRequestUserList {
	s.UserId = &v
	return s
}

func (s *StartChatRequestUserList) SetUserType(v string) *StartChatRequestUserList {
	s.UserType = &v
	return s
}

func (s *StartChatRequestUserList) Validate() error {
	return dara.Validate(s)
}
