// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRtcRobotInstanceShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigShrink(v string) *UpdateRtcRobotInstanceShrinkRequest
	GetConfigShrink() *string
	SetInstanceId(v string) *UpdateRtcRobotInstanceShrinkRequest
	GetInstanceId() *string
}

type UpdateRtcRobotInstanceShrinkRequest struct {
	ConfigShrink *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 727dc0e296014bb58670940a3da95592
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s UpdateRtcRobotInstanceShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateRtcRobotInstanceShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateRtcRobotInstanceShrinkRequest) GetConfigShrink() *string {
	return s.ConfigShrink
}

func (s *UpdateRtcRobotInstanceShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateRtcRobotInstanceShrinkRequest) SetConfigShrink(v string) *UpdateRtcRobotInstanceShrinkRequest {
	s.ConfigShrink = &v
	return s
}

func (s *UpdateRtcRobotInstanceShrinkRequest) SetInstanceId(v string) *UpdateRtcRobotInstanceShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateRtcRobotInstanceShrinkRequest) Validate() error {
	return dara.Validate(s)
}
