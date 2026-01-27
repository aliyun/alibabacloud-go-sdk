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
	// The batch operation mode. Set the value to **AllTogether**. In this mode, if all instances are stopped, a success message is returned. If an instance fails the verification, none of the instances can be stopped and an error message is returned.
	//
	// example:
	//
	// AllTogether
	BatchOptimization *string `json:"BatchOptimization,omitempty" xml:"BatchOptimization,omitempty"`
	// Specifies whether to forcefully stop the instance. Valid values:
	//
	// 	- **true**: forcefully stops the instance. If an instance fails to stop due to system or network issues, a forced stop can be triggered, **though it may result in data loss.**
	//
	// 	- **false**: does not forcefully stop the instance. This is the default value.
	//
	// example:
	//
	// false
	ForceStop *bool `json:"ForceStop,omitempty" xml:"ForceStop,omitempty"`
	// The node IDs.
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	// The region ID of the instance. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/26243.html) operation to query the most recent region list.
	//
	// example:
	//
	// cn-hangzhou
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	StoppedMode *string `json:"StoppedMode,omitempty" xml:"StoppedMode,omitempty"`
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
