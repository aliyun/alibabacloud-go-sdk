// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRtcRobotInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfig(v *UpdateRtcRobotInstanceRequestConfig) *UpdateRtcRobotInstanceRequest
	GetConfig() *UpdateRtcRobotInstanceRequestConfig
	SetInstanceId(v string) *UpdateRtcRobotInstanceRequest
	GetInstanceId() *string
}

type UpdateRtcRobotInstanceRequest struct {
	Config *UpdateRtcRobotInstanceRequestConfig `json:"Config,omitempty" xml:"Config,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 727dc0e296014bb58670940a3da95592
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s UpdateRtcRobotInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateRtcRobotInstanceRequest) GoString() string {
	return s.String()
}

func (s *UpdateRtcRobotInstanceRequest) GetConfig() *UpdateRtcRobotInstanceRequestConfig {
	return s.Config
}

func (s *UpdateRtcRobotInstanceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateRtcRobotInstanceRequest) SetConfig(v *UpdateRtcRobotInstanceRequestConfig) *UpdateRtcRobotInstanceRequest {
	s.Config = v
	return s
}

func (s *UpdateRtcRobotInstanceRequest) SetInstanceId(v string) *UpdateRtcRobotInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateRtcRobotInstanceRequest) Validate() error {
	return dara.Validate(s)
}

type UpdateRtcRobotInstanceRequestConfig struct {
	// example:
	//
	// false
	EnableVoiceInterrupt *bool   `json:"EnableVoiceInterrupt,omitempty" xml:"EnableVoiceInterrupt,omitempty"`
	Greeting             *string `json:"Greeting,omitempty" xml:"Greeting,omitempty"`
	// example:
	//
	// zhixiaoxia
	VoiceId *string `json:"VoiceId,omitempty" xml:"VoiceId,omitempty"`
}

func (s UpdateRtcRobotInstanceRequestConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateRtcRobotInstanceRequestConfig) GoString() string {
	return s.String()
}

func (s *UpdateRtcRobotInstanceRequestConfig) GetEnableVoiceInterrupt() *bool {
	return s.EnableVoiceInterrupt
}

func (s *UpdateRtcRobotInstanceRequestConfig) GetGreeting() *string {
	return s.Greeting
}

func (s *UpdateRtcRobotInstanceRequestConfig) GetVoiceId() *string {
	return s.VoiceId
}

func (s *UpdateRtcRobotInstanceRequestConfig) SetEnableVoiceInterrupt(v bool) *UpdateRtcRobotInstanceRequestConfig {
	s.EnableVoiceInterrupt = &v
	return s
}

func (s *UpdateRtcRobotInstanceRequestConfig) SetGreeting(v string) *UpdateRtcRobotInstanceRequestConfig {
	s.Greeting = &v
	return s
}

func (s *UpdateRtcRobotInstanceRequestConfig) SetVoiceId(v string) *UpdateRtcRobotInstanceRequestConfig {
	s.VoiceId = &v
	return s
}

func (s *UpdateRtcRobotInstanceRequestConfig) Validate() error {
	return dara.Validate(s)
}
