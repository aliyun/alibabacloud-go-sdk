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
	// This parameter is required.
	DesktopIds []*string `json:"DesktopIds,omitempty" xml:"DesktopIds,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// enter
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
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
