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
	SetStoppedMode(v string) *StopRCInstanceRequest
	GetStoppedMode() *string
}

type StopRCInstanceRequest struct {
	// Specifies whether to forcefully stop the instance. Valid values:
	//
	// -   **true**: Forcefully stops the instance.
	//
	// -   **false*	- (default): Gracefully stops the instance.
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
	// The stop mode of the instance. Valid values:
	//
	//   - StopCharging: economical mode. After economical mode is enabled:
	//
	//     - Billing for compute resources is suspended.
	//
	//     - Billing for system cloud disks and data cloud disks continues.
	//
	//     - Because compute resources are released, the instance may fail to start due to insufficient resources. Try again later or change the instance type.
	//
	//   - KeepCharging: standard mode. Billing continues after the instance is stopped.
	//
	// example:
	//
	// KeepCharging
	StoppedMode *string `json:"StoppedMode,omitempty" xml:"StoppedMode,omitempty"`
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

func (s *StopRCInstanceRequest) GetStoppedMode() *string {
	return s.StoppedMode
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

func (s *StopRCInstanceRequest) SetStoppedMode(v string) *StopRCInstanceRequest {
	s.StoppedMode = &v
	return s
}

func (s *StopRCInstanceRequest) Validate() error {
	return dara.Validate(s)
}
