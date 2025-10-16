// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAccessInstanceZoneListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessInstanceType(v string) *DescribeAccessInstanceZoneListRequest
	GetAccessInstanceType() *string
	SetLang(v string) *DescribeAccessInstanceZoneListRequest
	GetLang() *string
	SetRegionNo(v string) *DescribeAccessInstanceZoneListRequest
	GetRegionNo() *string
}

type DescribeAccessInstanceZoneListRequest struct {
	// example:
	//
	// AckClusterConnector
	AccessInstanceType *string `json:"AccessInstanceType,omitempty" xml:"AccessInstanceType,omitempty"`
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionNo *string `json:"RegionNo,omitempty" xml:"RegionNo,omitempty"`
}

func (s DescribeAccessInstanceZoneListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAccessInstanceZoneListRequest) GoString() string {
	return s.String()
}

func (s *DescribeAccessInstanceZoneListRequest) GetAccessInstanceType() *string {
	return s.AccessInstanceType
}

func (s *DescribeAccessInstanceZoneListRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeAccessInstanceZoneListRequest) GetRegionNo() *string {
	return s.RegionNo
}

func (s *DescribeAccessInstanceZoneListRequest) SetAccessInstanceType(v string) *DescribeAccessInstanceZoneListRequest {
	s.AccessInstanceType = &v
	return s
}

func (s *DescribeAccessInstanceZoneListRequest) SetLang(v string) *DescribeAccessInstanceZoneListRequest {
	s.Lang = &v
	return s
}

func (s *DescribeAccessInstanceZoneListRequest) SetRegionNo(v string) *DescribeAccessInstanceZoneListRequest {
	s.RegionNo = &v
	return s
}

func (s *DescribeAccessInstanceZoneListRequest) Validate() error {
	return dara.Validate(s)
}
