// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRegionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeRegionsResponseBody
	GetCode() *string
	SetRegions(v []*DescribeRegionsResponseBodyRegions) *DescribeRegionsResponseBody
	GetRegions() []*DescribeRegionsResponseBodyRegions
	SetRequestId(v string) *DescribeRegionsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeRegionsResponseBody
	GetSuccess() *bool
}

type DescribeRegionsResponseBody struct {
	// The HTTP status code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The information about the regions of cloud assets that can be protected by Anti-DDoS Origin. The information includes region IDs and names.
	Regions []*DescribeRegionsResponseBodyRegions `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// F7CA8B4E-FB15-4336-A351-8DC29D66EA82
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeRegionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRegionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeRegionsResponseBody) GetRegions() []*DescribeRegionsResponseBodyRegions {
	return s.Regions
}

func (s *DescribeRegionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRegionsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeRegionsResponseBody) SetCode(v string) *DescribeRegionsResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeRegionsResponseBody) SetRegions(v []*DescribeRegionsResponseBodyRegions) *DescribeRegionsResponseBody {
	s.Regions = v
	return s
}

func (s *DescribeRegionsResponseBody) SetRequestId(v string) *DescribeRegionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRegionsResponseBody) SetSuccess(v bool) *DescribeRegionsResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeRegionsResponseBody) Validate() error {
	if s.Regions != nil {
		for _, item := range s.Regions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeRegionsResponseBodyRegions struct {
	// The English name of the region.
	//
	// example:
	//
	// China (Hangzhou)
	RegionEnName *string `json:"RegionEnName,omitempty" xml:"RegionEnName,omitempty"`
	// The ID of the region.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The Chinese name of the region.
	//
	// example:
	//
	// 华东1（杭州）
	RegionName *string `json:"RegionName,omitempty" xml:"RegionName,omitempty"`
}

func (s DescribeRegionsResponseBodyRegions) String() string {
	return dara.Prettify(s)
}

func (s DescribeRegionsResponseBodyRegions) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBodyRegions) GetRegionEnName() *string {
	return s.RegionEnName
}

func (s *DescribeRegionsResponseBodyRegions) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeRegionsResponseBodyRegions) GetRegionName() *string {
	return s.RegionName
}

func (s *DescribeRegionsResponseBodyRegions) SetRegionEnName(v string) *DescribeRegionsResponseBodyRegions {
	s.RegionEnName = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegions) SetRegionId(v string) *DescribeRegionsResponseBodyRegions {
	s.RegionId = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegions) SetRegionName(v string) *DescribeRegionsResponseBodyRegions {
	s.RegionName = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegions) Validate() error {
	return dara.Validate(s)
}
