// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetDesktopGroupTimerStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDesktopGroupId(v string) *SetDesktopGroupTimerStatusRequest
	GetDesktopGroupId() *string
	SetRegionId(v string) *SetDesktopGroupTimerStatusRequest
	GetRegionId() *string
	SetStatus(v int32) *SetDesktopGroupTimerStatusRequest
	GetStatus() *int32
	SetTimerType(v int32) *SetDesktopGroupTimerStatusRequest
	GetTimerType() *int32
}

type SetDesktopGroupTimerStatusRequest struct {
	// The ID of the cloud computer share.
	//
	// This parameter is required.
	//
	// example:
	//
	// dg-fgxsniu6at****
	DesktopGroupId *string `json:"DesktopGroupId,omitempty" xml:"DesktopGroupId,omitempty"`
	// The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/196646.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The status of the scheduled task.
	//
	// Valid values:
	//
	// 	- 1: enabled
	//
	// 	- 2: disabled
	//
	// 	- 3: deleted
	//
	// 	- 100: unknown
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The type of the scheduled task.
	//
	// Valid values:
	//
	// 	- 1: scheduled reset
	//
	// 	- 2: scheduled startup
	//
	// 	- 3: scheduled stop
	//
	// 	- 4: scheduled restart
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	TimerType *int32 `json:"TimerType,omitempty" xml:"TimerType,omitempty"`
}

func (s SetDesktopGroupTimerStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s SetDesktopGroupTimerStatusRequest) GoString() string {
	return s.String()
}

func (s *SetDesktopGroupTimerStatusRequest) GetDesktopGroupId() *string {
	return s.DesktopGroupId
}

func (s *SetDesktopGroupTimerStatusRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *SetDesktopGroupTimerStatusRequest) GetStatus() *int32 {
	return s.Status
}

func (s *SetDesktopGroupTimerStatusRequest) GetTimerType() *int32 {
	return s.TimerType
}

func (s *SetDesktopGroupTimerStatusRequest) SetDesktopGroupId(v string) *SetDesktopGroupTimerStatusRequest {
	s.DesktopGroupId = &v
	return s
}

func (s *SetDesktopGroupTimerStatusRequest) SetRegionId(v string) *SetDesktopGroupTimerStatusRequest {
	s.RegionId = &v
	return s
}

func (s *SetDesktopGroupTimerStatusRequest) SetStatus(v int32) *SetDesktopGroupTimerStatusRequest {
	s.Status = &v
	return s
}

func (s *SetDesktopGroupTimerStatusRequest) SetTimerType(v int32) *SetDesktopGroupTimerStatusRequest {
	s.TimerType = &v
	return s
}

func (s *SetDesktopGroupTimerStatusRequest) Validate() error {
	return dara.Validate(s)
}
