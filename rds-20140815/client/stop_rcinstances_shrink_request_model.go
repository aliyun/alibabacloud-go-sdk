// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopRCInstancesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBatchOptimization(v string) *StopRCInstancesShrinkRequest
	GetBatchOptimization() *string
	SetForceStop(v bool) *StopRCInstancesShrinkRequest
	GetForceStop() *bool
	SetInstanceIdsShrink(v string) *StopRCInstancesShrinkRequest
	GetInstanceIdsShrink() *string
	SetRegionId(v string) *StopRCInstancesShrinkRequest
	GetRegionId() *string
	SetStoppedMode(v string) *StopRCInstancesShrinkRequest
	GetStoppedMode() *string
}

type StopRCInstancesShrinkRequest struct {
	BatchOptimization *string `json:"BatchOptimization,omitempty" xml:"BatchOptimization,omitempty"`
	ForceStop         *bool   `json:"ForceStop,omitempty" xml:"ForceStop,omitempty"`
	InstanceIdsShrink *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
	RegionId          *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	StoppedMode       *string `json:"StoppedMode,omitempty" xml:"StoppedMode,omitempty"`
}

func (s StopRCInstancesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s StopRCInstancesShrinkRequest) GoString() string {
	return s.String()
}

func (s *StopRCInstancesShrinkRequest) GetBatchOptimization() *string {
	return s.BatchOptimization
}

func (s *StopRCInstancesShrinkRequest) GetForceStop() *bool {
	return s.ForceStop
}

func (s *StopRCInstancesShrinkRequest) GetInstanceIdsShrink() *string {
	return s.InstanceIdsShrink
}

func (s *StopRCInstancesShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *StopRCInstancesShrinkRequest) GetStoppedMode() *string {
	return s.StoppedMode
}

func (s *StopRCInstancesShrinkRequest) SetBatchOptimization(v string) *StopRCInstancesShrinkRequest {
	s.BatchOptimization = &v
	return s
}

func (s *StopRCInstancesShrinkRequest) SetForceStop(v bool) *StopRCInstancesShrinkRequest {
	s.ForceStop = &v
	return s
}

func (s *StopRCInstancesShrinkRequest) SetInstanceIdsShrink(v string) *StopRCInstancesShrinkRequest {
	s.InstanceIdsShrink = &v
	return s
}

func (s *StopRCInstancesShrinkRequest) SetRegionId(v string) *StopRCInstancesShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *StopRCInstancesShrinkRequest) SetStoppedMode(v string) *StopRCInstancesShrinkRequest {
	s.StoppedMode = &v
	return s
}

func (s *StopRCInstancesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
