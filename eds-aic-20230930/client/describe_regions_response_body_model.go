// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRegionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRegionModels(v []*DescribeRegionsResponseBodyRegionModels) *DescribeRegionsResponseBody
	GetRegionModels() []*DescribeRegionsResponseBodyRegionModels
	SetRequestId(v string) *DescribeRegionsResponseBody
	GetRequestId() *string
}

type DescribeRegionsResponseBody struct {
	// Available regions.
	RegionModels []*DescribeRegionsResponseBodyRegionModels `json:"RegionModels,omitempty" xml:"RegionModels,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// A87B3769-0D05-5383-B236-5798B455****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeRegionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRegionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBody) GetRegionModels() []*DescribeRegionsResponseBodyRegionModels {
	return s.RegionModels
}

func (s *DescribeRegionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRegionsResponseBody) SetRegionModels(v []*DescribeRegionsResponseBodyRegionModels) *DescribeRegionsResponseBody {
	s.RegionModels = v
	return s
}

func (s *DescribeRegionsResponseBody) SetRequestId(v string) *DescribeRegionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRegionsResponseBody) Validate() error {
	if s.RegionModels != nil {
		for _, item := range s.RegionModels {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeRegionsResponseBodyRegionModels struct {
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The region name.
	//
	// example:
	//
	// China (Hangzhou)
	RegionName *string `json:"RegionName,omitempty" xml:"RegionName,omitempty"`
}

func (s DescribeRegionsResponseBodyRegionModels) String() string {
	return dara.Prettify(s)
}

func (s DescribeRegionsResponseBodyRegionModels) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBodyRegionModels) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeRegionsResponseBodyRegionModels) GetRegionName() *string {
	return s.RegionName
}

func (s *DescribeRegionsResponseBodyRegionModels) SetRegionId(v string) *DescribeRegionsResponseBodyRegionModels {
	s.RegionId = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegionModels) SetRegionName(v string) *DescribeRegionsResponseBodyRegionModels {
	s.RegionName = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegionModels) Validate() error {
	return dara.Validate(s)
}
