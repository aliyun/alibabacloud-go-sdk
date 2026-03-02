// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteOfficeSiteAcceleratorRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOfficeSiteId(v string) *DeleteOfficeSiteAcceleratorRequest
	GetOfficeSiteId() *string
	SetRegionId(v string) *DeleteOfficeSiteAcceleratorRequest
	GetRegionId() *string
}

type DeleteOfficeSiteAcceleratorRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou+dir-885351****
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteOfficeSiteAcceleratorRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteOfficeSiteAcceleratorRequest) GoString() string {
	return s.String()
}

func (s *DeleteOfficeSiteAcceleratorRequest) GetOfficeSiteId() *string {
	return s.OfficeSiteId
}

func (s *DeleteOfficeSiteAcceleratorRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteOfficeSiteAcceleratorRequest) SetOfficeSiteId(v string) *DeleteOfficeSiteAcceleratorRequest {
	s.OfficeSiteId = &v
	return s
}

func (s *DeleteOfficeSiteAcceleratorRequest) SetRegionId(v string) *DeleteOfficeSiteAcceleratorRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteOfficeSiteAcceleratorRequest) Validate() error {
	return dara.Validate(s)
}
