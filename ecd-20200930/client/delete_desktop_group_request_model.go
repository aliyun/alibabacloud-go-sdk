// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDesktopGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDesktopGroupId(v string) *DeleteDesktopGroupRequest
	GetDesktopGroupId() *string
	SetRegionId(v string) *DeleteDesktopGroupRequest
	GetRegionId() *string
	SetResellerOwnerUid(v int64) *DeleteDesktopGroupRequest
	GetResellerOwnerUid() *int64
}

type DeleteDesktopGroupRequest struct {
	// The ID of the shared Cloud Desktop.
	//
	// This parameter is required.
	//
	// example:
	//
	// dg-2i8qxpv6t1a03****
	DesktopGroupId *string `json:"DesktopGroupId,omitempty" xml:"DesktopGroupId,omitempty"`
	// The region ID. You can call [DescribeRegions](https://help.aliyun.com/document_detail/196646.html) to query the regions supported by WUYING Workspace.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource ownership user in the reseller pattern. This parameter is not required in the non-reseller pattern.
	//
	// example:
	//
	// 1422724566551XXX
	ResellerOwnerUid *int64 `json:"ResellerOwnerUid,omitempty" xml:"ResellerOwnerUid,omitempty"`
}

func (s DeleteDesktopGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDesktopGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteDesktopGroupRequest) GetDesktopGroupId() *string {
	return s.DesktopGroupId
}

func (s *DeleteDesktopGroupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteDesktopGroupRequest) GetResellerOwnerUid() *int64 {
	return s.ResellerOwnerUid
}

func (s *DeleteDesktopGroupRequest) SetDesktopGroupId(v string) *DeleteDesktopGroupRequest {
	s.DesktopGroupId = &v
	return s
}

func (s *DeleteDesktopGroupRequest) SetRegionId(v string) *DeleteDesktopGroupRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteDesktopGroupRequest) SetResellerOwnerUid(v int64) *DeleteDesktopGroupRequest {
	s.ResellerOwnerUid = &v
	return s
}

func (s *DeleteDesktopGroupRequest) Validate() error {
	return dara.Validate(s)
}
