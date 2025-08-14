// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopRtcRobotInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *StopRtcRobotInstanceRequest
	GetInstanceId() *string
}

type StopRtcRobotInstanceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 727dc0e296014bb58670940a3da95592
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s StopRtcRobotInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s StopRtcRobotInstanceRequest) GoString() string {
	return s.String()
}

func (s *StopRtcRobotInstanceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *StopRtcRobotInstanceRequest) SetInstanceId(v string) *StopRtcRobotInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *StopRtcRobotInstanceRequest) Validate() error {
	return dara.Validate(s)
}
