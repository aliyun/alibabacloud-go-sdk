// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListImageRegistryRegionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRegions(v []*ListImageRegistryRegionResponseBodyRegions) *ListImageRegistryRegionResponseBody
	GetRegions() []*ListImageRegistryRegionResponseBodyRegions
	SetRequestId(v string) *ListImageRegistryRegionResponseBody
	GetRequestId() *string
}

type ListImageRegistryRegionResponseBody struct {
	// An array that consists of regions.
	Regions []*ListImageRegistryRegionResponseBodyRegions `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Repeated"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 291B49F9-1685-4005-9D34-606B6F78****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListImageRegistryRegionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListImageRegistryRegionResponseBody) GoString() string {
	return s.String()
}

func (s *ListImageRegistryRegionResponseBody) GetRegions() []*ListImageRegistryRegionResponseBodyRegions {
	return s.Regions
}

func (s *ListImageRegistryRegionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListImageRegistryRegionResponseBody) SetRegions(v []*ListImageRegistryRegionResponseBodyRegions) *ListImageRegistryRegionResponseBody {
	s.Regions = v
	return s
}

func (s *ListImageRegistryRegionResponseBody) SetRequestId(v string) *ListImageRegistryRegionResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListImageRegistryRegionResponseBody) Validate() error {
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

type ListImageRegistryRegionResponseBodyRegions struct {
	// The region ID of the image.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the region.
	//
	// example:
	//
	// cn-hangzhou
	RegionName *string `json:"RegionName,omitempty" xml:"RegionName,omitempty"`
}

func (s ListImageRegistryRegionResponseBodyRegions) String() string {
	return dara.Prettify(s)
}

func (s ListImageRegistryRegionResponseBodyRegions) GoString() string {
	return s.String()
}

func (s *ListImageRegistryRegionResponseBodyRegions) GetRegionId() *string {
	return s.RegionId
}

func (s *ListImageRegistryRegionResponseBodyRegions) GetRegionName() *string {
	return s.RegionName
}

func (s *ListImageRegistryRegionResponseBodyRegions) SetRegionId(v string) *ListImageRegistryRegionResponseBodyRegions {
	s.RegionId = &v
	return s
}

func (s *ListImageRegistryRegionResponseBodyRegions) SetRegionName(v string) *ListImageRegistryRegionResponseBodyRegions {
	s.RegionName = &v
	return s
}

func (s *ListImageRegistryRegionResponseBodyRegions) Validate() error {
	return dara.Validate(s)
}
