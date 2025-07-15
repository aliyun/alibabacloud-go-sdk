// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartChatShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessChannelId(v string) *StartChatShrinkRequest
	GetAccessChannelId() *string
	SetInstanceId(v string) *StartChatShrinkRequest
	GetInstanceId() *string
	SetToken(v string) *StartChatShrinkRequest
	GetToken() *string
	SetUserListShrink(v string) *StartChatShrinkRequest
	GetUserListShrink() *string
}

type StartChatShrinkRequest struct {
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
	UserListShrink *string `json:"UserList,omitempty" xml:"UserList,omitempty"`
}

func (s StartChatShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s StartChatShrinkRequest) GoString() string {
	return s.String()
}

func (s *StartChatShrinkRequest) GetAccessChannelId() *string {
	return s.AccessChannelId
}

func (s *StartChatShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *StartChatShrinkRequest) GetToken() *string {
	return s.Token
}

func (s *StartChatShrinkRequest) GetUserListShrink() *string {
	return s.UserListShrink
}

func (s *StartChatShrinkRequest) SetAccessChannelId(v string) *StartChatShrinkRequest {
	s.AccessChannelId = &v
	return s
}

func (s *StartChatShrinkRequest) SetInstanceId(v string) *StartChatShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *StartChatShrinkRequest) SetToken(v string) *StartChatShrinkRequest {
	s.Token = &v
	return s
}

func (s *StartChatShrinkRequest) SetUserListShrink(v string) *StartChatShrinkRequest {
	s.UserListShrink = &v
	return s
}

func (s *StartChatShrinkRequest) Validate() error {
	return dara.Validate(s)
}
