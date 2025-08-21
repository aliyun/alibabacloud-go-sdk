// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUnreadMessageCountShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetUserInfoShrink(v string) *GetUnreadMessageCountShrinkRequest
	GetUserInfoShrink() *string
}

type GetUnreadMessageCountShrinkRequest struct {
	UserInfoShrink *string `json:"UserInfo,omitempty" xml:"UserInfo,omitempty"`
}

func (s GetUnreadMessageCountShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetUnreadMessageCountShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetUnreadMessageCountShrinkRequest) GetUserInfoShrink() *string {
	return s.UserInfoShrink
}

func (s *GetUnreadMessageCountShrinkRequest) SetUserInfoShrink(v string) *GetUnreadMessageCountShrinkRequest {
	s.UserInfoShrink = &v
	return s
}

func (s *GetUnreadMessageCountShrinkRequest) Validate() error {
	return dara.Validate(s)
}
