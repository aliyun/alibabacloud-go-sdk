// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAvailableCrossRegionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRegions(v *DescribeAvailableCrossRegionResponseBodyRegions) *DescribeAvailableCrossRegionResponseBody
	GetRegions() *DescribeAvailableCrossRegionResponseBodyRegions
	SetRequestId(v string) *DescribeAvailableCrossRegionResponseBody
	GetRequestId() *string
}

type DescribeAvailableCrossRegionResponseBody struct {
	// An array that consists of destination regions for cross-region backups.
	Regions *DescribeAvailableCrossRegionResponseBodyRegions `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 39265F46-EC77-4036-8AC4-F035F32F6BE2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeAvailableCrossRegionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAvailableCrossRegionResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAvailableCrossRegionResponseBody) GetRegions() *DescribeAvailableCrossRegionResponseBodyRegions {
	return s.Regions
}

func (s *DescribeAvailableCrossRegionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAvailableCrossRegionResponseBody) SetRegions(v *DescribeAvailableCrossRegionResponseBodyRegions) *DescribeAvailableCrossRegionResponseBody {
	s.Regions = v
	return s
}

func (s *DescribeAvailableCrossRegionResponseBody) SetRequestId(v string) *DescribeAvailableCrossRegionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAvailableCrossRegionResponseBody) Validate() error {
	if s.Regions != nil {
		if err := s.Regions.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeAvailableCrossRegionResponseBodyRegions struct {
	Region []*string `json:"Region,omitempty" xml:"Region,omitempty" type:"Repeated"`
}

func (s DescribeAvailableCrossRegionResponseBodyRegions) String() string {
	return dara.Prettify(s)
}

func (s DescribeAvailableCrossRegionResponseBodyRegions) GoString() string {
	return s.String()
}

func (s *DescribeAvailableCrossRegionResponseBodyRegions) GetRegion() []*string {
	return s.Region
}

func (s *DescribeAvailableCrossRegionResponseBodyRegions) SetRegion(v []*string) *DescribeAvailableCrossRegionResponseBodyRegions {
	s.Region = v
	return s
}

func (s *DescribeAvailableCrossRegionResponseBodyRegions) Validate() error {
	return dara.Validate(s)
}
