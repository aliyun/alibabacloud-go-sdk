// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyOfficeSiteCrossDesktopAccessRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnableCrossDesktopAccess(v bool) *ModifyOfficeSiteCrossDesktopAccessRequest
	GetEnableCrossDesktopAccess() *bool
	SetOfficeSiteId(v string) *ModifyOfficeSiteCrossDesktopAccessRequest
	GetOfficeSiteId() *string
	SetRegionId(v string) *ModifyOfficeSiteCrossDesktopAccessRequest
	GetRegionId() *string
}

type ModifyOfficeSiteCrossDesktopAccessRequest struct {
	// Specifies whether to enable the communication between cloud computers in an office network. If you enable the communication between cloud computers in an office network, the cloud computers can access each other.
	//
	// This parameter is required.
	//
	// example:
	//
	// false
	EnableCrossDesktopAccess *bool `json:"EnableCrossDesktopAccess,omitempty" xml:"EnableCrossDesktopAccess,omitempty"`
	// The office network ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou+dir-068266****
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	// The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/196646.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ModifyOfficeSiteCrossDesktopAccessRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyOfficeSiteCrossDesktopAccessRequest) GoString() string {
	return s.String()
}

func (s *ModifyOfficeSiteCrossDesktopAccessRequest) GetEnableCrossDesktopAccess() *bool {
	return s.EnableCrossDesktopAccess
}

func (s *ModifyOfficeSiteCrossDesktopAccessRequest) GetOfficeSiteId() *string {
	return s.OfficeSiteId
}

func (s *ModifyOfficeSiteCrossDesktopAccessRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyOfficeSiteCrossDesktopAccessRequest) SetEnableCrossDesktopAccess(v bool) *ModifyOfficeSiteCrossDesktopAccessRequest {
	s.EnableCrossDesktopAccess = &v
	return s
}

func (s *ModifyOfficeSiteCrossDesktopAccessRequest) SetOfficeSiteId(v string) *ModifyOfficeSiteCrossDesktopAccessRequest {
	s.OfficeSiteId = &v
	return s
}

func (s *ModifyOfficeSiteCrossDesktopAccessRequest) SetRegionId(v string) *ModifyOfficeSiteCrossDesktopAccessRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyOfficeSiteCrossDesktopAccessRequest) Validate() error {
	return dara.Validate(s)
}
