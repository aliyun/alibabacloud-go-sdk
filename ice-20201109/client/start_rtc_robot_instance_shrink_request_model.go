// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartRtcRobotInstanceShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthToken(v string) *StartRtcRobotInstanceShrinkRequest
	GetAuthToken() *string
	SetChannelId(v string) *StartRtcRobotInstanceShrinkRequest
	GetChannelId() *string
	SetConfigShrink(v string) *StartRtcRobotInstanceShrinkRequest
	GetConfigShrink() *string
	SetRobotId(v string) *StartRtcRobotInstanceShrinkRequest
	GetRobotId() *string
	SetUserData(v string) *StartRtcRobotInstanceShrinkRequest
	GetUserData() *string
	SetUserId(v string) *StartRtcRobotInstanceShrinkRequest
	GetUserId() *string
}

type StartRtcRobotInstanceShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// **********
	AuthToken *string `json:"AuthToken,omitempty" xml:"AuthToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// testId
	ChannelId    *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	ConfigShrink *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ca28b08ad3464ebcb42e5c0f7c6d2e89
	RobotId *string `json:"RobotId,omitempty" xml:"RobotId,omitempty"`
	// example:
	//
	// {}
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// my-robot
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s StartRtcRobotInstanceShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s StartRtcRobotInstanceShrinkRequest) GoString() string {
	return s.String()
}

func (s *StartRtcRobotInstanceShrinkRequest) GetAuthToken() *string {
	return s.AuthToken
}

func (s *StartRtcRobotInstanceShrinkRequest) GetChannelId() *string {
	return s.ChannelId
}

func (s *StartRtcRobotInstanceShrinkRequest) GetConfigShrink() *string {
	return s.ConfigShrink
}

func (s *StartRtcRobotInstanceShrinkRequest) GetRobotId() *string {
	return s.RobotId
}

func (s *StartRtcRobotInstanceShrinkRequest) GetUserData() *string {
	return s.UserData
}

func (s *StartRtcRobotInstanceShrinkRequest) GetUserId() *string {
	return s.UserId
}

func (s *StartRtcRobotInstanceShrinkRequest) SetAuthToken(v string) *StartRtcRobotInstanceShrinkRequest {
	s.AuthToken = &v
	return s
}

func (s *StartRtcRobotInstanceShrinkRequest) SetChannelId(v string) *StartRtcRobotInstanceShrinkRequest {
	s.ChannelId = &v
	return s
}

func (s *StartRtcRobotInstanceShrinkRequest) SetConfigShrink(v string) *StartRtcRobotInstanceShrinkRequest {
	s.ConfigShrink = &v
	return s
}

func (s *StartRtcRobotInstanceShrinkRequest) SetRobotId(v string) *StartRtcRobotInstanceShrinkRequest {
	s.RobotId = &v
	return s
}

func (s *StartRtcRobotInstanceShrinkRequest) SetUserData(v string) *StartRtcRobotInstanceShrinkRequest {
	s.UserData = &v
	return s
}

func (s *StartRtcRobotInstanceShrinkRequest) SetUserId(v string) *StartRtcRobotInstanceShrinkRequest {
	s.UserId = &v
	return s
}

func (s *StartRtcRobotInstanceShrinkRequest) Validate() error {
	return dara.Validate(s)
}
