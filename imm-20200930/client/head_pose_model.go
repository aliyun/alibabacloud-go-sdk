// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHeadPose interface {
	dara.Model
	String() string
	GoString() string
	SetPitch(v float32) *HeadPose
	GetPitch() *float32
	SetRoll(v float32) *HeadPose
	GetRoll() *float32
	SetYaw(v float32) *HeadPose
	GetYaw() *float32
}

type HeadPose struct {
	// The angel of elevation or depression of the head. Unit: degree. Valid values: -180 to 180. A recommended range for reliable results is from -30 to 30.
	//
	// example:
	//
	// 18.385589599609375
	Pitch *float32 `json:"Pitch,omitempty" xml:"Pitch,omitempty"`
	// The angle of the tilt to the side. Unit: degree. Valid values: -180 to 180. A recommended range for reliable results is from -45 to 45.
	//
	// example:
	//
	// 4.204030513763428
	Roll *float32 `json:"Roll,omitempty" xml:"Roll,omitempty"`
	// The angle of leftward or rightward rotation of the head. Unit: degree. Valid values: -180 to 180. A recommended range for reliable results is from -80 to 80.
	//
	// example:
	//
	// 2.4945924282073975
	Yaw *float32 `json:"Yaw,omitempty" xml:"Yaw,omitempty"`
}

func (s HeadPose) String() string {
	return dara.Prettify(s)
}

func (s HeadPose) GoString() string {
	return s.String()
}

func (s *HeadPose) GetPitch() *float32 {
	return s.Pitch
}

func (s *HeadPose) GetRoll() *float32 {
	return s.Roll
}

func (s *HeadPose) GetYaw() *float32 {
	return s.Yaw
}

func (s *HeadPose) SetPitch(v float32) *HeadPose {
	s.Pitch = &v
	return s
}

func (s *HeadPose) SetRoll(v float32) *HeadPose {
	s.Roll = &v
	return s
}

func (s *HeadPose) SetYaw(v float32) *HeadPose {
	s.Yaw = &v
	return s
}

func (s *HeadPose) Validate() error {
	return dara.Validate(s)
}
