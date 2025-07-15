// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDesktopGroupDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDesktopGroupId(v string) *GetDesktopGroupDetailRequest
	GetDesktopGroupId() *string
	SetRegionId(v string) *GetDesktopGroupDetailRequest
	GetRegionId() *string
}

type GetDesktopGroupDetailRequest struct {
	// The ID of the cloud computer share.
	//
	// This parameter is required.
	//
	// example:
	//
	// dg-2i8qxpv6t1a03****
	DesktopGroupId *string `json:"DesktopGroupId,omitempty" xml:"DesktopGroupId,omitempty"`
	// The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/196646.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetDesktopGroupDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDesktopGroupDetailRequest) GoString() string {
	return s.String()
}

func (s *GetDesktopGroupDetailRequest) GetDesktopGroupId() *string {
	return s.DesktopGroupId
}

func (s *GetDesktopGroupDetailRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetDesktopGroupDetailRequest) SetDesktopGroupId(v string) *GetDesktopGroupDetailRequest {
	s.DesktopGroupId = &v
	return s
}

func (s *GetDesktopGroupDetailRequest) SetRegionId(v string) *GetDesktopGroupDetailRequest {
	s.RegionId = &v
	return s
}

func (s *GetDesktopGroupDetailRequest) Validate() error {
	return dara.Validate(s)
}
