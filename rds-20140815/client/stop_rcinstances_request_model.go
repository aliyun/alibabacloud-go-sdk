// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopRCInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBatchOptimization(v string) *StopRCInstancesRequest
	GetBatchOptimization() *string
	SetForceStop(v bool) *StopRCInstancesRequest
	GetForceStop() *bool
	SetInstanceIds(v []*string) *StopRCInstancesRequest
	GetInstanceIds() []*string
	SetRegionId(v string) *StopRCInstancesRequest
	GetRegionId() *string
	SetStoppedMode(v string) *StopRCInstancesRequest
	GetStoppedMode() *string
}

type StopRCInstancesRequest struct {
	BatchOptimization *string   `json:"BatchOptimization,omitempty" xml:"BatchOptimization,omitempty"`
	ForceStop         *bool     `json:"ForceStop,omitempty" xml:"ForceStop,omitempty"`
	InstanceIds       []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	RegionId          *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	StoppedMode       *string   `json:"StoppedMode,omitempty" xml:"StoppedMode,omitempty"`
}

func (s StopRCInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s StopRCInstancesRequest) GoString() string {
	return s.String()
}

func (s *StopRCInstancesRequest) GetBatchOptimization() *string {
	return s.BatchOptimization
}

func (s *StopRCInstancesRequest) GetForceStop() *bool {
	return s.ForceStop
}

func (s *StopRCInstancesRequest) GetInstanceIds() []*string {
	return s.InstanceIds
}

func (s *StopRCInstancesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *StopRCInstancesRequest) GetStoppedMode() *string {
	return s.StoppedMode
}

func (s *StopRCInstancesRequest) SetBatchOptimization(v string) *StopRCInstancesRequest {
	s.BatchOptimization = &v
	return s
}

func (s *StopRCInstancesRequest) SetForceStop(v bool) *StopRCInstancesRequest {
	s.ForceStop = &v
	return s
}

func (s *StopRCInstancesRequest) SetInstanceIds(v []*string) *StopRCInstancesRequest {
	s.InstanceIds = v
	return s
}

func (s *StopRCInstancesRequest) SetRegionId(v string) *StopRCInstancesRequest {
	s.RegionId = &v
	return s
}

func (s *StopRCInstancesRequest) SetStoppedMode(v string) *StopRCInstancesRequest {
	s.StoppedMode = &v
	return s
}

func (s *StopRCInstancesRequest) Validate() error {
	return dara.Validate(s)
}
