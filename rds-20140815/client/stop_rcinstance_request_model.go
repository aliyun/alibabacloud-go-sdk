// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopRCInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetForceStop(v bool) *StopRCInstanceRequest
	GetForceStop() *bool
	SetInstanceId(v string) *StopRCInstanceRequest
	GetInstanceId() *string
	SetRegionId(v string) *StopRCInstanceRequest
	GetRegionId() *string
}

type StopRCInstanceRequest struct {
	// Specifies whether to forcefully stop the instance. Valid values:
	//
	// 	- **true**
	//
	// 	- **false*	- (default)
	//
	// example:
	//
	// false
	ForceStop *bool `json:"ForceStop,omitempty" xml:"ForceStop,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rc-m5sc1271fv344a1r****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s StopRCInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s StopRCInstanceRequest) GoString() string {
	return s.String()
}

func (s *StopRCInstanceRequest) GetForceStop() *bool {
	return s.ForceStop
}

func (s *StopRCInstanceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *StopRCInstanceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *StopRCInstanceRequest) SetForceStop(v bool) *StopRCInstanceRequest {
	s.ForceStop = &v
	return s
}

func (s *StopRCInstanceRequest) SetInstanceId(v string) *StopRCInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *StopRCInstanceRequest) SetRegionId(v string) *StopRCInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *StopRCInstanceRequest) Validate() error {
	return dara.Validate(s)
}
