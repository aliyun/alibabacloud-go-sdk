// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopDesktopsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDesktopId(v []*string) *StopDesktopsRequest
	GetDesktopId() []*string
	SetOsUpdate(v bool) *StopDesktopsRequest
	GetOsUpdate() *bool
	SetRegionId(v string) *StopDesktopsRequest
	GetRegionId() *string
	SetStoppedMode(v string) *StopDesktopsRequest
	GetStoppedMode() *string
}

type StopDesktopsRequest struct {
	// The cloud computer IDs. You can specify the IDs of 1 to 100 cloud computers.
	//
	// This parameter is required.
	//
	// example:
	//
	// ecd-7w78ozhjcwa3u****
	DesktopId []*string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty" type:"Repeated"`
	// Whether to perform a patch update when the update is ready. A value of true indicates that a patch update is performed.
	//
	// example:
	//
	// false
	OsUpdate *bool `json:"OsUpdate,omitempty" xml:"OsUpdate,omitempty"`
	// The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/196646.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The billing mode after you stop the cloud computer.
	//
	// Default value: StopCharging. Valid values:
	//
	// 	- StopCharging: After the cloud computer is stopped, the system automatically reclaims computing resources. You are no longer charged for computing resources. However, you are still charged for storage resources.
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- KeepCharging: After the cloud computer is stopped, the system does not reclaim resources to prevent insufficient resources and startup failures. You are still charged for the resources.
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// StopCharging
	StoppedMode *string `json:"StoppedMode,omitempty" xml:"StoppedMode,omitempty"`
}

func (s StopDesktopsRequest) String() string {
	return dara.Prettify(s)
}

func (s StopDesktopsRequest) GoString() string {
	return s.String()
}

func (s *StopDesktopsRequest) GetDesktopId() []*string {
	return s.DesktopId
}

func (s *StopDesktopsRequest) GetOsUpdate() *bool {
	return s.OsUpdate
}

func (s *StopDesktopsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *StopDesktopsRequest) GetStoppedMode() *string {
	return s.StoppedMode
}

func (s *StopDesktopsRequest) SetDesktopId(v []*string) *StopDesktopsRequest {
	s.DesktopId = v
	return s
}

func (s *StopDesktopsRequest) SetOsUpdate(v bool) *StopDesktopsRequest {
	s.OsUpdate = &v
	return s
}

func (s *StopDesktopsRequest) SetRegionId(v string) *StopDesktopsRequest {
	s.RegionId = &v
	return s
}

func (s *StopDesktopsRequest) SetStoppedMode(v string) *StopDesktopsRequest {
	s.StoppedMode = &v
	return s
}

func (s *StopDesktopsRequest) Validate() error {
	return dara.Validate(s)
}
