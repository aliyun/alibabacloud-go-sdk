// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBusiRegionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRegions(v []*ListBusiRegionsResponseBodyRegions) *ListBusiRegionsResponseBody
	GetRegions() []*ListBusiRegionsResponseBodyRegions
	SetRequestId(v string) *ListBusiRegionsResponseBody
	GetRequestId() *string
}

type ListBusiRegionsResponseBody struct {
	// The information about the acceleration regions that are supported by GA.
	Regions []*ListBusiRegionsResponseBodyRegions `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 6FEA0CF3-D3B9-43E5-A304-D217037876A8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListBusiRegionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListBusiRegionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListBusiRegionsResponseBody) GetRegions() []*ListBusiRegionsResponseBodyRegions {
	return s.Regions
}

func (s *ListBusiRegionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListBusiRegionsResponseBody) SetRegions(v []*ListBusiRegionsResponseBodyRegions) *ListBusiRegionsResponseBody {
	s.Regions = v
	return s
}

func (s *ListBusiRegionsResponseBody) SetRequestId(v string) *ListBusiRegionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListBusiRegionsResponseBody) Validate() error {
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

type ListBusiRegionsResponseBodyRegions struct {
	// The name of the region.
	//
	// example:
	//
	// China (Hangzhou)
	LocalName *string `json:"LocalName,omitempty" xml:"LocalName,omitempty"`
	// The ID of the region.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListBusiRegionsResponseBodyRegions) String() string {
	return dara.Prettify(s)
}

func (s ListBusiRegionsResponseBodyRegions) GoString() string {
	return s.String()
}

func (s *ListBusiRegionsResponseBodyRegions) GetLocalName() *string {
	return s.LocalName
}

func (s *ListBusiRegionsResponseBodyRegions) GetRegionId() *string {
	return s.RegionId
}

func (s *ListBusiRegionsResponseBodyRegions) SetLocalName(v string) *ListBusiRegionsResponseBodyRegions {
	s.LocalName = &v
	return s
}

func (s *ListBusiRegionsResponseBodyRegions) SetRegionId(v string) *ListBusiRegionsResponseBodyRegions {
	s.RegionId = &v
	return s
}

func (s *ListBusiRegionsResponseBodyRegions) Validate() error {
	return dara.Validate(s)
}
