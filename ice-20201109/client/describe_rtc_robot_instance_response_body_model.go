// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRtcRobotInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAuthToken(v string) *DescribeRtcRobotInstanceResponseBody
	GetAuthToken() *string
	SetChannelId(v string) *DescribeRtcRobotInstanceResponseBody
	GetChannelId() *string
	SetConfig(v *DescribeRtcRobotInstanceResponseBodyConfig) *DescribeRtcRobotInstanceResponseBody
	GetConfig() *DescribeRtcRobotInstanceResponseBodyConfig
	SetRequestId(v string) *DescribeRtcRobotInstanceResponseBody
	GetRequestId() *string
	SetStatus(v string) *DescribeRtcRobotInstanceResponseBody
	GetStatus() *string
	SetUserData(v string) *DescribeRtcRobotInstanceResponseBody
	GetUserData() *string
	SetUserId(v string) *DescribeRtcRobotInstanceResponseBody
	GetUserId() *string
}

type DescribeRtcRobotInstanceResponseBody struct {
	// example:
	//
	// **********
	AuthToken *string `json:"AuthToken,omitempty" xml:"AuthToken,omitempty"`
	// example:
	//
	// testId
	ChannelId *string                                     `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	Config    *DescribeRtcRobotInstanceResponseBodyConfig `json:"Config,omitempty" xml:"Config,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 20B3A1B6-4BD2-5DE6-BCBC-098C9B4F4E91
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// Executing
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// {}
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
	// example:
	//
	// my-robot
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s DescribeRtcRobotInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRtcRobotInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRtcRobotInstanceResponseBody) GetAuthToken() *string {
	return s.AuthToken
}

func (s *DescribeRtcRobotInstanceResponseBody) GetChannelId() *string {
	return s.ChannelId
}

func (s *DescribeRtcRobotInstanceResponseBody) GetConfig() *DescribeRtcRobotInstanceResponseBodyConfig {
	return s.Config
}

func (s *DescribeRtcRobotInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRtcRobotInstanceResponseBody) GetStatus() *string {
	return s.Status
}

func (s *DescribeRtcRobotInstanceResponseBody) GetUserData() *string {
	return s.UserData
}

func (s *DescribeRtcRobotInstanceResponseBody) GetUserId() *string {
	return s.UserId
}

func (s *DescribeRtcRobotInstanceResponseBody) SetAuthToken(v string) *DescribeRtcRobotInstanceResponseBody {
	s.AuthToken = &v
	return s
}

func (s *DescribeRtcRobotInstanceResponseBody) SetChannelId(v string) *DescribeRtcRobotInstanceResponseBody {
	s.ChannelId = &v
	return s
}

func (s *DescribeRtcRobotInstanceResponseBody) SetConfig(v *DescribeRtcRobotInstanceResponseBodyConfig) *DescribeRtcRobotInstanceResponseBody {
	s.Config = v
	return s
}

func (s *DescribeRtcRobotInstanceResponseBody) SetRequestId(v string) *DescribeRtcRobotInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRtcRobotInstanceResponseBody) SetStatus(v string) *DescribeRtcRobotInstanceResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeRtcRobotInstanceResponseBody) SetUserData(v string) *DescribeRtcRobotInstanceResponseBody {
	s.UserData = &v
	return s
}

func (s *DescribeRtcRobotInstanceResponseBody) SetUserId(v string) *DescribeRtcRobotInstanceResponseBody {
	s.UserId = &v
	return s
}

func (s *DescribeRtcRobotInstanceResponseBody) Validate() error {
	if s.Config != nil {
		if err := s.Config.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeRtcRobotInstanceResponseBodyConfig struct {
	// example:
	//
	// true
	EnableVoiceInterrupt *bool   `json:"EnableVoiceInterrupt,omitempty" xml:"EnableVoiceInterrupt,omitempty"`
	Greeting             *string `json:"Greeting,omitempty" xml:"Greeting,omitempty"`
	// example:
	//
	// zhixiaoxia
	VoiceId *string `json:"VoiceId,omitempty" xml:"VoiceId,omitempty"`
}

func (s DescribeRtcRobotInstanceResponseBodyConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeRtcRobotInstanceResponseBodyConfig) GoString() string {
	return s.String()
}

func (s *DescribeRtcRobotInstanceResponseBodyConfig) GetEnableVoiceInterrupt() *bool {
	return s.EnableVoiceInterrupt
}

func (s *DescribeRtcRobotInstanceResponseBodyConfig) GetGreeting() *string {
	return s.Greeting
}

func (s *DescribeRtcRobotInstanceResponseBodyConfig) GetVoiceId() *string {
	return s.VoiceId
}

func (s *DescribeRtcRobotInstanceResponseBodyConfig) SetEnableVoiceInterrupt(v bool) *DescribeRtcRobotInstanceResponseBodyConfig {
	s.EnableVoiceInterrupt = &v
	return s
}

func (s *DescribeRtcRobotInstanceResponseBodyConfig) SetGreeting(v string) *DescribeRtcRobotInstanceResponseBodyConfig {
	s.Greeting = &v
	return s
}

func (s *DescribeRtcRobotInstanceResponseBodyConfig) SetVoiceId(v string) *DescribeRtcRobotInstanceResponseBodyConfig {
	s.VoiceId = &v
	return s
}

func (s *DescribeRtcRobotInstanceResponseBodyConfig) Validate() error {
	return dara.Validate(s)
}
