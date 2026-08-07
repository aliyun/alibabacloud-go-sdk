// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRegionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DescribeRegionsResponseBodyData) *DescribeRegionsResponseBody
	GetData() *DescribeRegionsResponseBodyData
	SetRequestId(v string) *DescribeRegionsResponseBody
	GetRequestId() *string
}

type DescribeRegionsResponseBody struct {
	// Response parameters
	Data *DescribeRegionsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// AE43C4CB-8074-5EBD-9806-8CA6D12800B1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeRegionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRegionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBody) GetData() *DescribeRegionsResponseBodyData {
	return s.Data
}

func (s *DescribeRegionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRegionsResponseBody) SetData(v *DescribeRegionsResponseBodyData) *DescribeRegionsResponseBody {
	s.Data = v
	return s
}

func (s *DescribeRegionsResponseBody) SetRequestId(v string) *DescribeRegionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRegionsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeRegionsResponseBodyData struct {
	// Regions
	Regions []*DescribeRegionsResponseBodyDataRegions `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Repeated"`
}

func (s DescribeRegionsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeRegionsResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBodyData) GetRegions() []*DescribeRegionsResponseBodyDataRegions {
	return s.Regions
}

func (s *DescribeRegionsResponseBodyData) SetRegions(v []*DescribeRegionsResponseBodyDataRegions) *DescribeRegionsResponseBodyData {
	s.Regions = v
	return s
}

func (s *DescribeRegionsResponseBodyData) Validate() error {
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

type DescribeRegionsResponseBodyDataRegions struct {
	// example:
	//
	// China (Hangzhou)
	LocalName *string `json:"LocalName,omitempty" xml:"LocalName,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeRegionsResponseBodyDataRegions) String() string {
	return dara.Prettify(s)
}

func (s DescribeRegionsResponseBodyDataRegions) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBodyDataRegions) GetLocalName() *string {
	return s.LocalName
}

func (s *DescribeRegionsResponseBodyDataRegions) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeRegionsResponseBodyDataRegions) SetLocalName(v string) *DescribeRegionsResponseBodyDataRegions {
	s.LocalName = &v
	return s
}

func (s *DescribeRegionsResponseBodyDataRegions) SetRegionId(v string) *DescribeRegionsResponseBodyDataRegions {
	s.RegionId = &v
	return s
}

func (s *DescribeRegionsResponseBodyDataRegions) Validate() error {
	return dara.Validate(s)
}
