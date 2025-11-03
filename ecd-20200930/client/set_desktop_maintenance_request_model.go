// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetDesktopMaintenanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDesktopIds(v []*string) *SetDesktopMaintenanceRequest
	GetDesktopIds() []*string
	SetMode(v string) *SetDesktopMaintenanceRequest
	GetMode() *string
	SetRegionId(v string) *SetDesktopMaintenanceRequest
	GetRegionId() *string
}

type SetDesktopMaintenanceRequest struct {
	// A list of cloud computer IDs for which you want to set maintenance mode. A maximum of 100 cloud computer IDs are supported.
	//
	// This parameter is required.
	DesktopIds []*string `json:"DesktopIds,omitempty" xml:"DesktopIds,omitempty" type:"Repeated"`
	// Enter or exit cloud computer maintenance mode.
	//
	// Enumerated values:
	//
	// 	- ENTER: The enters the maintenance mode.
	//
	// 	- EXIT: The exits the maintenance mode.
	//
	// This parameter is required.
	//
	// example:
	//
	// enter
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// The region ID. You can call the [DescribeRegions](~~DescribeRegions~~) operation to query the list of regions where Elastic Desktop Service (EDS) Enterprise is available.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s SetDesktopMaintenanceRequest) String() string {
	return dara.Prettify(s)
}

func (s SetDesktopMaintenanceRequest) GoString() string {
	return s.String()
}

func (s *SetDesktopMaintenanceRequest) GetDesktopIds() []*string {
	return s.DesktopIds
}

func (s *SetDesktopMaintenanceRequest) GetMode() *string {
	return s.Mode
}

func (s *SetDesktopMaintenanceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *SetDesktopMaintenanceRequest) SetDesktopIds(v []*string) *SetDesktopMaintenanceRequest {
	s.DesktopIds = v
	return s
}

func (s *SetDesktopMaintenanceRequest) SetMode(v string) *SetDesktopMaintenanceRequest {
	s.Mode = &v
	return s
}

func (s *SetDesktopMaintenanceRequest) SetRegionId(v string) *SetDesktopMaintenanceRequest {
	s.RegionId = &v
	return s
}

func (s *SetDesktopMaintenanceRequest) Validate() error {
	return dara.Validate(s)
}
