// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDesktopInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDesktopId(v []*string) *DescribeDesktopInfoRequest
	GetDesktopId() []*string
	SetNeedExtraInfo(v bool) *DescribeDesktopInfoRequest
	GetNeedExtraInfo() *bool
	SetRegionId(v string) *DescribeDesktopInfoRequest
	GetRegionId() *string
}

type DescribeDesktopInfoRequest struct {
	// The IDs of the cloud computers. You can specify 1 to 100 IDs.
	DesktopId     []*string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty" type:"Repeated"`
	NeedExtraInfo *bool     `json:"NeedExtraInfo,omitempty" xml:"NeedExtraInfo,omitempty"`
	// The ID of the region. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/196646.html) operation to query the regions supported by Elastic Desktop Service.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeDesktopInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDesktopInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeDesktopInfoRequest) GetDesktopId() []*string {
	return s.DesktopId
}

func (s *DescribeDesktopInfoRequest) GetNeedExtraInfo() *bool {
	return s.NeedExtraInfo
}

func (s *DescribeDesktopInfoRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDesktopInfoRequest) SetDesktopId(v []*string) *DescribeDesktopInfoRequest {
	s.DesktopId = v
	return s
}

func (s *DescribeDesktopInfoRequest) SetNeedExtraInfo(v bool) *DescribeDesktopInfoRequest {
	s.NeedExtraInfo = &v
	return s
}

func (s *DescribeDesktopInfoRequest) SetRegionId(v string) *DescribeDesktopInfoRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDesktopInfoRequest) Validate() error {
	return dara.Validate(s)
}
