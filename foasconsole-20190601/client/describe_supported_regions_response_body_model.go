// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSupportedRegionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRegions(v []*DescribeSupportedRegionsResponseBodyRegions) *DescribeSupportedRegionsResponseBody
	GetRegions() []*DescribeSupportedRegionsResponseBodyRegions
	SetRequestId(v string) *DescribeSupportedRegionsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeSupportedRegionsResponseBody
	GetSuccess() *bool
}

type DescribeSupportedRegionsResponseBody struct {
	Regions []*DescribeSupportedRegionsResponseBodyRegions `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Repeated"`
	// example:
	//
	// B21DC47E-8928-199A-9F32-36D45E4693B4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeSupportedRegionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSupportedRegionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSupportedRegionsResponseBody) GetRegions() []*DescribeSupportedRegionsResponseBodyRegions {
	return s.Regions
}

func (s *DescribeSupportedRegionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSupportedRegionsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeSupportedRegionsResponseBody) SetRegions(v []*DescribeSupportedRegionsResponseBodyRegions) *DescribeSupportedRegionsResponseBody {
	s.Regions = v
	return s
}

func (s *DescribeSupportedRegionsResponseBody) SetRequestId(v string) *DescribeSupportedRegionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSupportedRegionsResponseBody) SetSuccess(v bool) *DescribeSupportedRegionsResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeSupportedRegionsResponseBody) Validate() error {
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

type DescribeSupportedRegionsResponseBodyRegions struct {
	// example:
	//
	// cn-beijing
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// example:
	//
	// 华北2 (北京)
	RegionName *string `json:"RegionName,omitempty" xml:"RegionName,omitempty"`
}

func (s DescribeSupportedRegionsResponseBodyRegions) String() string {
	return dara.Prettify(s)
}

func (s DescribeSupportedRegionsResponseBodyRegions) GoString() string {
	return s.String()
}

func (s *DescribeSupportedRegionsResponseBodyRegions) GetRegion() *string {
	return s.Region
}

func (s *DescribeSupportedRegionsResponseBodyRegions) GetRegionName() *string {
	return s.RegionName
}

func (s *DescribeSupportedRegionsResponseBodyRegions) SetRegion(v string) *DescribeSupportedRegionsResponseBodyRegions {
	s.Region = &v
	return s
}

func (s *DescribeSupportedRegionsResponseBodyRegions) SetRegionName(v string) *DescribeSupportedRegionsResponseBodyRegions {
	s.RegionName = &v
	return s
}

func (s *DescribeSupportedRegionsResponseBodyRegions) Validate() error {
	return dara.Validate(s)
}
