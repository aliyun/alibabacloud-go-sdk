// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendMessageShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetUrl(v string) *SendMessageShrinkRequest
	GetUrl() *string
	SetUserInfoShrink(v string) *SendMessageShrinkRequest
	GetUserInfoShrink() *string
}

type SendMessageShrinkRequest struct {
	// example:
	//
	// http://xx
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
	// This parameter is required.
	UserInfoShrink *string `json:"UserInfo,omitempty" xml:"UserInfo,omitempty"`
}

func (s SendMessageShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s SendMessageShrinkRequest) GoString() string {
	return s.String()
}

func (s *SendMessageShrinkRequest) GetUrl() *string {
	return s.Url
}

func (s *SendMessageShrinkRequest) GetUserInfoShrink() *string {
	return s.UserInfoShrink
}

func (s *SendMessageShrinkRequest) SetUrl(v string) *SendMessageShrinkRequest {
	s.Url = &v
	return s
}

func (s *SendMessageShrinkRequest) SetUserInfoShrink(v string) *SendMessageShrinkRequest {
	s.UserInfoShrink = &v
	return s
}

func (s *SendMessageShrinkRequest) Validate() error {
	return dara.Validate(s)
}
