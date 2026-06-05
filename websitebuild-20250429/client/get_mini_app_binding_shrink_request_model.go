// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMiniAppBindingShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizId(v string) *GetMiniAppBindingShrinkRequest
	GetBizId() *string
	SetChannel(v string) *GetMiniAppBindingShrinkRequest
	GetChannel() *string
	SetSettingKeysShrink(v string) *GetMiniAppBindingShrinkRequest
	GetSettingKeysShrink() *string
}

type GetMiniAppBindingShrinkRequest struct {
	// example:
	//
	// WS20250814102215000001
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	// example:
	//
	// WECHAT
	Channel           *string `json:"Channel,omitempty" xml:"Channel,omitempty"`
	SettingKeysShrink *string `json:"SettingKeys,omitempty" xml:"SettingKeys,omitempty"`
}

func (s GetMiniAppBindingShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetMiniAppBindingShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetMiniAppBindingShrinkRequest) GetBizId() *string {
	return s.BizId
}

func (s *GetMiniAppBindingShrinkRequest) GetChannel() *string {
	return s.Channel
}

func (s *GetMiniAppBindingShrinkRequest) GetSettingKeysShrink() *string {
	return s.SettingKeysShrink
}

func (s *GetMiniAppBindingShrinkRequest) SetBizId(v string) *GetMiniAppBindingShrinkRequest {
	s.BizId = &v
	return s
}

func (s *GetMiniAppBindingShrinkRequest) SetChannel(v string) *GetMiniAppBindingShrinkRequest {
	s.Channel = &v
	return s
}

func (s *GetMiniAppBindingShrinkRequest) SetSettingKeysShrink(v string) *GetMiniAppBindingShrinkRequest {
	s.SettingKeysShrink = &v
	return s
}

func (s *GetMiniAppBindingShrinkRequest) Validate() error {
	return dara.Validate(s)
}
