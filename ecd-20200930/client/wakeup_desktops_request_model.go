// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iWakeupDesktopsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDesktopId(v []*string) *WakeupDesktopsRequest
	GetDesktopId() []*string
	SetRegionId(v string) *WakeupDesktopsRequest
	GetRegionId() *string
}

type WakeupDesktopsRequest struct {
	// The IDs of the cloud computers. You can specify the IDs of 1 to 100 cloud computers.
	//
	// This parameter is required.
	DesktopId []*string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty" type:"Repeated"`
	// The region ID. You can call the [DescribeRegions](~~DescribeRegions~~) operation to query the regions supported by Elastic Desktop Service (EDS).
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s WakeupDesktopsRequest) String() string {
	return dara.Prettify(s)
}

func (s WakeupDesktopsRequest) GoString() string {
	return s.String()
}

func (s *WakeupDesktopsRequest) GetDesktopId() []*string {
	return s.DesktopId
}

func (s *WakeupDesktopsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *WakeupDesktopsRequest) SetDesktopId(v []*string) *WakeupDesktopsRequest {
	s.DesktopId = v
	return s
}

func (s *WakeupDesktopsRequest) SetRegionId(v string) *WakeupDesktopsRequest {
	s.RegionId = &v
	return s
}

func (s *WakeupDesktopsRequest) Validate() error {
	return dara.Validate(s)
}
