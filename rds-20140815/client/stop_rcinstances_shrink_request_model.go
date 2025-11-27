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
}

type StopRCInstancesShrinkRequest struct {
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
	InstanceIdsShrink *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
	// The region ID of the instance. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/26243.html) operation to query the most recent region list.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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

func (s *StopRCInstancesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
