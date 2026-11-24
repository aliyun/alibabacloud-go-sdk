// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeResourceStatisticByRegionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIncludeGroupDesktops(v bool) *DescribeResourceStatisticByRegionRequest
	GetIncludeGroupDesktops() *bool
	SetRegionId(v string) *DescribeResourceStatisticByRegionRequest
	GetRegionId() *string
	SetSearchRegionId(v string) *DescribeResourceStatisticByRegionRequest
	GetSearchRegionId() *string
}

type DescribeResourceStatisticByRegionRequest struct {
	// Specifies whether the statistics include the number of shared cloud desktops.
	//
	// example:
	//
	// true
	IncludeGroupDesktops *bool `json:"IncludeGroupDesktops,omitempty" xml:"IncludeGroupDesktops,omitempty"`
	// The ID of the service region. Valid values: China (Shanghai) and Singapore.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the region to search. This parameter is used to filter cloud desktops in the specified region.
	//
	// example:
	//
	// cn-shanghai
	SearchRegionId *string `json:"SearchRegionId,omitempty" xml:"SearchRegionId,omitempty"`
}

func (s DescribeResourceStatisticByRegionRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeResourceStatisticByRegionRequest) GoString() string {
	return s.String()
}

func (s *DescribeResourceStatisticByRegionRequest) GetIncludeGroupDesktops() *bool {
	return s.IncludeGroupDesktops
}

func (s *DescribeResourceStatisticByRegionRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeResourceStatisticByRegionRequest) GetSearchRegionId() *string {
	return s.SearchRegionId
}

func (s *DescribeResourceStatisticByRegionRequest) SetIncludeGroupDesktops(v bool) *DescribeResourceStatisticByRegionRequest {
	s.IncludeGroupDesktops = &v
	return s
}

func (s *DescribeResourceStatisticByRegionRequest) SetRegionId(v string) *DescribeResourceStatisticByRegionRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeResourceStatisticByRegionRequest) SetSearchRegionId(v string) *DescribeResourceStatisticByRegionRequest {
	s.SearchRegionId = &v
	return s
}

func (s *DescribeResourceStatisticByRegionRequest) Validate() error {
	return dara.Validate(s)
}
