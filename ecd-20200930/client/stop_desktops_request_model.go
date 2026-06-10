// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopDesktopsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCreateSnapshot(v string) *StopDesktopsRequest
	GetCreateSnapshot() *string
	SetDesktopId(v []*string) *StopDesktopsRequest
	GetDesktopId() []*string
	SetOsUpdate(v bool) *StopDesktopsRequest
	GetOsUpdate() *bool
	SetPatchId(v string) *StopDesktopsRequest
	GetPatchId() *string
	SetRegionId(v string) *StopDesktopsRequest
	GetRegionId() *string
	SetStoppedMode(v string) *StopDesktopsRequest
	GetStoppedMode() *string
}

type StopDesktopsRequest struct {
	// example:
	//
	// false
	CreateSnapshot *string `json:"CreateSnapshot,omitempty" xml:"CreateSnapshot,omitempty"`
	// An array of 1 to 100 cloud desktop IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// ecd-7w78ozhjcwa3u****
	DesktopId []*string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty" type:"Repeated"`
	// Specifies whether to apply pending patch updates.
	//
	// example:
	//
	// false
	OsUpdate *bool `json:"OsUpdate,omitempty" xml:"OsUpdate,omitempty"`
	// example:
	//
	// KB5082063
	PatchId *string `json:"PatchId,omitempty" xml:"PatchId,omitempty"`
	// The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/196646.html) operation to get a list of regions that Elastic Desktop Service supports.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Specifies the billing mode for the cloud desktops after they are stopped.
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

func (s *StopDesktopsRequest) GetCreateSnapshot() *string {
	return s.CreateSnapshot
}

func (s *StopDesktopsRequest) GetDesktopId() []*string {
	return s.DesktopId
}

func (s *StopDesktopsRequest) GetOsUpdate() *bool {
	return s.OsUpdate
}

func (s *StopDesktopsRequest) GetPatchId() *string {
	return s.PatchId
}

func (s *StopDesktopsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *StopDesktopsRequest) GetStoppedMode() *string {
	return s.StoppedMode
}

func (s *StopDesktopsRequest) SetCreateSnapshot(v string) *StopDesktopsRequest {
	s.CreateSnapshot = &v
	return s
}

func (s *StopDesktopsRequest) SetDesktopId(v []*string) *StopDesktopsRequest {
	s.DesktopId = v
	return s
}

func (s *StopDesktopsRequest) SetOsUpdate(v bool) *StopDesktopsRequest {
	s.OsUpdate = &v
	return s
}

func (s *StopDesktopsRequest) SetPatchId(v string) *StopDesktopsRequest {
	s.PatchId = &v
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
