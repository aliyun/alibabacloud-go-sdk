// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLoginPolarClawChannelRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *LoginPolarClawChannelRequest
	GetApplicationId() *string
	SetChannelId(v string) *LoginPolarClawChannelRequest
	GetChannelId() *string
}

type LoginPolarClawChannelRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pa-**************
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// openclaw-weixin
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
}

func (s LoginPolarClawChannelRequest) String() string {
	return dara.Prettify(s)
}

func (s LoginPolarClawChannelRequest) GoString() string {
	return s.String()
}

func (s *LoginPolarClawChannelRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *LoginPolarClawChannelRequest) GetChannelId() *string {
	return s.ChannelId
}

func (s *LoginPolarClawChannelRequest) SetApplicationId(v string) *LoginPolarClawChannelRequest {
	s.ApplicationId = &v
	return s
}

func (s *LoginPolarClawChannelRequest) SetChannelId(v string) *LoginPolarClawChannelRequest {
	s.ChannelId = &v
	return s
}

func (s *LoginPolarClawChannelRequest) Validate() error {
	return dara.Validate(s)
}
