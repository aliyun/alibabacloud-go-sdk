// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRumAppRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppGroup(v string) *DeleteRumAppRequest
	GetAppGroup() *string
	SetAppId(v string) *DeleteRumAppRequest
	GetAppId() *string
	SetRealRegionId(v string) *DeleteRumAppRequest
	GetRealRegionId() *string
	SetRegionId(v string) *DeleteRumAppRequest
	GetRegionId() *string
}

type DeleteRumAppRequest struct {
	// The group where the application resides.
	//
	// example:
	//
	// default
	AppGroup *string `json:"AppGroup,omitempty" xml:"AppGroup,omitempty"`
	// The ID of the application.
	//
	// This parameter is required.
	//
	// example:
	//
	// b590lhguqs@28f515462******
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The region where the application resides. You can leave this parameter empty or set it to China East 2 Finance.
	//
	// example:
	//
	// cn-shanghai-finance-1
	RealRegionId *string `json:"RealRegionId,omitempty" xml:"RealRegionId,omitempty"`
	// The ID of the region.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteRumAppRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteRumAppRequest) GoString() string {
	return s.String()
}

func (s *DeleteRumAppRequest) GetAppGroup() *string {
	return s.AppGroup
}

func (s *DeleteRumAppRequest) GetAppId() *string {
	return s.AppId
}

func (s *DeleteRumAppRequest) GetRealRegionId() *string {
	return s.RealRegionId
}

func (s *DeleteRumAppRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteRumAppRequest) SetAppGroup(v string) *DeleteRumAppRequest {
	s.AppGroup = &v
	return s
}

func (s *DeleteRumAppRequest) SetAppId(v string) *DeleteRumAppRequest {
	s.AppId = &v
	return s
}

func (s *DeleteRumAppRequest) SetRealRegionId(v string) *DeleteRumAppRequest {
	s.RealRegionId = &v
	return s
}

func (s *DeleteRumAppRequest) SetRegionId(v string) *DeleteRumAppRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteRumAppRequest) Validate() error {
	return dara.Validate(s)
}
