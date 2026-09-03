// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableDesktopsInGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDesktopGroupId(v string) *DisableDesktopsInGroupRequest
	GetDesktopGroupId() *string
	SetDesktopIds(v []*string) *DisableDesktopsInGroupRequest
	GetDesktopIds() []*string
	SetRegionId(v string) *DisableDesktopsInGroupRequest
	GetRegionId() *string
}

type DisableDesktopsInGroupRequest struct {
	// The ID of the shared cloud computer group.
	//
	// This parameter is required.
	//
	// example:
	//
	// dg-2i8qxpv6t1a03****
	DesktopGroupId *string `json:"DesktopGroupId,omitempty" xml:"DesktopGroupId,omitempty"`
	// The list of cloud computer IDs.
	//
	// This parameter is required.
	DesktopIds []*string `json:"DesktopIds,omitempty" xml:"DesktopIds,omitempty" type:"Repeated"`
	// The region ID. You can call [DescribeRegions](https://help.aliyun.com/document_detail/196646.html) to query the regions supported by Elastic Desktop Service.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DisableDesktopsInGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s DisableDesktopsInGroupRequest) GoString() string {
	return s.String()
}

func (s *DisableDesktopsInGroupRequest) GetDesktopGroupId() *string {
	return s.DesktopGroupId
}

func (s *DisableDesktopsInGroupRequest) GetDesktopIds() []*string {
	return s.DesktopIds
}

func (s *DisableDesktopsInGroupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DisableDesktopsInGroupRequest) SetDesktopGroupId(v string) *DisableDesktopsInGroupRequest {
	s.DesktopGroupId = &v
	return s
}

func (s *DisableDesktopsInGroupRequest) SetDesktopIds(v []*string) *DisableDesktopsInGroupRequest {
	s.DesktopIds = v
	return s
}

func (s *DisableDesktopsInGroupRequest) SetRegionId(v string) *DisableDesktopsInGroupRequest {
	s.RegionId = &v
	return s
}

func (s *DisableDesktopsInGroupRequest) Validate() error {
	return dara.Validate(s)
}
