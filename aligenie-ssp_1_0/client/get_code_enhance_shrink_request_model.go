// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCodeEnhanceShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChannelInfoShrink(v string) *GetCodeEnhanceShrinkRequest
	GetChannelInfoShrink() *string
	SetUserInfoShrink(v string) *GetCodeEnhanceShrinkRequest
	GetUserInfoShrink() *string
}

type GetCodeEnhanceShrinkRequest struct {
	// This parameter is required.
	ChannelInfoShrink *string `json:"ChannelInfo,omitempty" xml:"ChannelInfo,omitempty"`
	// This parameter is required.
	UserInfoShrink *string `json:"UserInfo,omitempty" xml:"UserInfo,omitempty"`
}

func (s GetCodeEnhanceShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetCodeEnhanceShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetCodeEnhanceShrinkRequest) GetChannelInfoShrink() *string {
	return s.ChannelInfoShrink
}

func (s *GetCodeEnhanceShrinkRequest) GetUserInfoShrink() *string {
	return s.UserInfoShrink
}

func (s *GetCodeEnhanceShrinkRequest) SetChannelInfoShrink(v string) *GetCodeEnhanceShrinkRequest {
	s.ChannelInfoShrink = &v
	return s
}

func (s *GetCodeEnhanceShrinkRequest) SetUserInfoShrink(v string) *GetCodeEnhanceShrinkRequest {
	s.UserInfoShrink = &v
	return s
}

func (s *GetCodeEnhanceShrinkRequest) Validate() error {
	return dara.Validate(s)
}
