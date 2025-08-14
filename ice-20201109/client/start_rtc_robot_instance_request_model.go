// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartRtcRobotInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthToken(v string) *StartRtcRobotInstanceRequest
	GetAuthToken() *string
	SetChannelId(v string) *StartRtcRobotInstanceRequest
	GetChannelId() *string
	SetConfig(v *StartRtcRobotInstanceRequestConfig) *StartRtcRobotInstanceRequest
	GetConfig() *StartRtcRobotInstanceRequestConfig
	SetRobotId(v string) *StartRtcRobotInstanceRequest
	GetRobotId() *string
	SetUserData(v string) *StartRtcRobotInstanceRequest
	GetUserData() *string
	SetUserId(v string) *StartRtcRobotInstanceRequest
	GetUserId() *string
}

type StartRtcRobotInstanceRequest struct {
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
	ChannelId *string                             `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	Config    *StartRtcRobotInstanceRequestConfig `json:"Config,omitempty" xml:"Config,omitempty" type:"Struct"`
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

func (s StartRtcRobotInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s StartRtcRobotInstanceRequest) GoString() string {
	return s.String()
}

func (s *StartRtcRobotInstanceRequest) GetAuthToken() *string {
	return s.AuthToken
}

func (s *StartRtcRobotInstanceRequest) GetChannelId() *string {
	return s.ChannelId
}

func (s *StartRtcRobotInstanceRequest) GetConfig() *StartRtcRobotInstanceRequestConfig {
	return s.Config
}

func (s *StartRtcRobotInstanceRequest) GetRobotId() *string {
	return s.RobotId
}

func (s *StartRtcRobotInstanceRequest) GetUserData() *string {
	return s.UserData
}

func (s *StartRtcRobotInstanceRequest) GetUserId() *string {
	return s.UserId
}

func (s *StartRtcRobotInstanceRequest) SetAuthToken(v string) *StartRtcRobotInstanceRequest {
	s.AuthToken = &v
	return s
}

func (s *StartRtcRobotInstanceRequest) SetChannelId(v string) *StartRtcRobotInstanceRequest {
	s.ChannelId = &v
	return s
}

func (s *StartRtcRobotInstanceRequest) SetConfig(v *StartRtcRobotInstanceRequestConfig) *StartRtcRobotInstanceRequest {
	s.Config = v
	return s
}

func (s *StartRtcRobotInstanceRequest) SetRobotId(v string) *StartRtcRobotInstanceRequest {
	s.RobotId = &v
	return s
}

func (s *StartRtcRobotInstanceRequest) SetUserData(v string) *StartRtcRobotInstanceRequest {
	s.UserData = &v
	return s
}

func (s *StartRtcRobotInstanceRequest) SetUserId(v string) *StartRtcRobotInstanceRequest {
	s.UserId = &v
	return s
}

func (s *StartRtcRobotInstanceRequest) Validate() error {
	return dara.Validate(s)
}

type StartRtcRobotInstanceRequestConfig struct {
	AsrMaxSilence *int32 `json:"AsrMaxSilence,omitempty" xml:"AsrMaxSilence,omitempty"`
	// example:
	//
	// true
	EnableVoiceInterrupt *bool   `json:"EnableVoiceInterrupt,omitempty" xml:"EnableVoiceInterrupt,omitempty"`
	Greeting             *string `json:"Greeting,omitempty" xml:"Greeting,omitempty"`
	UseVoiceprint        *bool   `json:"UseVoiceprint,omitempty" xml:"UseVoiceprint,omitempty"`
	UserOfflineTimeout   *int32  `json:"UserOfflineTimeout,omitempty" xml:"UserOfflineTimeout,omitempty"`
	UserOnlineTimeout    *int32  `json:"UserOnlineTimeout,omitempty" xml:"UserOnlineTimeout,omitempty"`
	// example:
	//
	// zhixiaoxia
	VoiceId      *string `json:"VoiceId,omitempty" xml:"VoiceId,omitempty"`
	VoiceprintId *string `json:"VoiceprintId,omitempty" xml:"VoiceprintId,omitempty"`
	Volume       *int64  `json:"Volume,omitempty" xml:"Volume,omitempty"`
}

func (s StartRtcRobotInstanceRequestConfig) String() string {
	return dara.Prettify(s)
}

func (s StartRtcRobotInstanceRequestConfig) GoString() string {
	return s.String()
}

func (s *StartRtcRobotInstanceRequestConfig) GetAsrMaxSilence() *int32 {
	return s.AsrMaxSilence
}

func (s *StartRtcRobotInstanceRequestConfig) GetEnableVoiceInterrupt() *bool {
	return s.EnableVoiceInterrupt
}

func (s *StartRtcRobotInstanceRequestConfig) GetGreeting() *string {
	return s.Greeting
}

func (s *StartRtcRobotInstanceRequestConfig) GetUseVoiceprint() *bool {
	return s.UseVoiceprint
}

func (s *StartRtcRobotInstanceRequestConfig) GetUserOfflineTimeout() *int32 {
	return s.UserOfflineTimeout
}

func (s *StartRtcRobotInstanceRequestConfig) GetUserOnlineTimeout() *int32 {
	return s.UserOnlineTimeout
}

func (s *StartRtcRobotInstanceRequestConfig) GetVoiceId() *string {
	return s.VoiceId
}

func (s *StartRtcRobotInstanceRequestConfig) GetVoiceprintId() *string {
	return s.VoiceprintId
}

func (s *StartRtcRobotInstanceRequestConfig) GetVolume() *int64 {
	return s.Volume
}

func (s *StartRtcRobotInstanceRequestConfig) SetAsrMaxSilence(v int32) *StartRtcRobotInstanceRequestConfig {
	s.AsrMaxSilence = &v
	return s
}

func (s *StartRtcRobotInstanceRequestConfig) SetEnableVoiceInterrupt(v bool) *StartRtcRobotInstanceRequestConfig {
	s.EnableVoiceInterrupt = &v
	return s
}

func (s *StartRtcRobotInstanceRequestConfig) SetGreeting(v string) *StartRtcRobotInstanceRequestConfig {
	s.Greeting = &v
	return s
}

func (s *StartRtcRobotInstanceRequestConfig) SetUseVoiceprint(v bool) *StartRtcRobotInstanceRequestConfig {
	s.UseVoiceprint = &v
	return s
}

func (s *StartRtcRobotInstanceRequestConfig) SetUserOfflineTimeout(v int32) *StartRtcRobotInstanceRequestConfig {
	s.UserOfflineTimeout = &v
	return s
}

func (s *StartRtcRobotInstanceRequestConfig) SetUserOnlineTimeout(v int32) *StartRtcRobotInstanceRequestConfig {
	s.UserOnlineTimeout = &v
	return s
}

func (s *StartRtcRobotInstanceRequestConfig) SetVoiceId(v string) *StartRtcRobotInstanceRequestConfig {
	s.VoiceId = &v
	return s
}

func (s *StartRtcRobotInstanceRequestConfig) SetVoiceprintId(v string) *StartRtcRobotInstanceRequestConfig {
	s.VoiceprintId = &v
	return s
}

func (s *StartRtcRobotInstanceRequestConfig) SetVolume(v int64) *StartRtcRobotInstanceRequestConfig {
	s.Volume = &v
	return s
}

func (s *StartRtcRobotInstanceRequestConfig) Validate() error {
	return dara.Validate(s)
}
