// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeOfficeSiteBridgeInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBridgeId(v string) *DescribeOfficeSiteBridgeInfoRequest
	GetBridgeId() *string
	SetOfficeSiteId(v string) *DescribeOfficeSiteBridgeInfoRequest
	GetOfficeSiteId() *string
	SetRegionId(v string) *DescribeOfficeSiteBridgeInfoRequest
	GetRegionId() *string
}

type DescribeOfficeSiteBridgeInfoRequest struct {
	// The virtual bridge ID.
	//
	// example:
	//
	// vb-sdfsifhisdhf****
	BridgeId *string `json:"BridgeId,omitempty" xml:"BridgeId,omitempty"`
	// The office network ID.
	//
	// example:
	//
	// cn-hangzhou+dir-363353****
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	// The region ID. You can call [DescribeRegions](https://help.aliyun.com/document_detail/196646.html) to query the regions supported by Elastic Desktop Service.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeOfficeSiteBridgeInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeOfficeSiteBridgeInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeOfficeSiteBridgeInfoRequest) GetBridgeId() *string {
	return s.BridgeId
}

func (s *DescribeOfficeSiteBridgeInfoRequest) GetOfficeSiteId() *string {
	return s.OfficeSiteId
}

func (s *DescribeOfficeSiteBridgeInfoRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeOfficeSiteBridgeInfoRequest) SetBridgeId(v string) *DescribeOfficeSiteBridgeInfoRequest {
	s.BridgeId = &v
	return s
}

func (s *DescribeOfficeSiteBridgeInfoRequest) SetOfficeSiteId(v string) *DescribeOfficeSiteBridgeInfoRequest {
	s.OfficeSiteId = &v
	return s
}

func (s *DescribeOfficeSiteBridgeInfoRequest) SetRegionId(v string) *DescribeOfficeSiteBridgeInfoRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeOfficeSiteBridgeInfoRequest) Validate() error {
	return dara.Validate(s)
}
