// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBusinessRegionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCount(v int64) *ListBusinessRegionsResponseBody
	GetCount() *int64
	SetGeographicSubRegions(v []*ListBusinessRegionsResponseBodyGeographicSubRegions) *ListBusinessRegionsResponseBody
	GetGeographicSubRegions() []*ListBusinessRegionsResponseBodyGeographicSubRegions
	SetRequestId(v string) *ListBusinessRegionsResponseBody
	GetRequestId() *string
}

type ListBusinessRegionsResponseBody struct {
	// The total number of entries returned.
	//
	// example:
	//
	// 10
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The list of regions available for Express Connect circuits.
	GeographicSubRegions []*ListBusinessRegionsResponseBodyGeographicSubRegions `json:"GeographicSubRegions,omitempty" xml:"GeographicSubRegions,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 304FE68E-16D8-5B90-B2B3-FE5C5C08C24B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListBusinessRegionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListBusinessRegionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListBusinessRegionsResponseBody) GetCount() *int64 {
	return s.Count
}

func (s *ListBusinessRegionsResponseBody) GetGeographicSubRegions() []*ListBusinessRegionsResponseBodyGeographicSubRegions {
	return s.GeographicSubRegions
}

func (s *ListBusinessRegionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListBusinessRegionsResponseBody) SetCount(v int64) *ListBusinessRegionsResponseBody {
	s.Count = &v
	return s
}

func (s *ListBusinessRegionsResponseBody) SetGeographicSubRegions(v []*ListBusinessRegionsResponseBodyGeographicSubRegions) *ListBusinessRegionsResponseBody {
	s.GeographicSubRegions = v
	return s
}

func (s *ListBusinessRegionsResponseBody) SetRequestId(v string) *ListBusinessRegionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListBusinessRegionsResponseBody) Validate() error {
	if s.GeographicSubRegions != nil {
		for _, item := range s.GeographicSubRegions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListBusinessRegionsResponseBodyGeographicSubRegions struct {
	// The name of the region where circuits are available.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the region where circuits are available.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListBusinessRegionsResponseBodyGeographicSubRegions) String() string {
	return dara.Prettify(s)
}

func (s ListBusinessRegionsResponseBodyGeographicSubRegions) GoString() string {
	return s.String()
}

func (s *ListBusinessRegionsResponseBodyGeographicSubRegions) GetName() *string {
	return s.Name
}

func (s *ListBusinessRegionsResponseBodyGeographicSubRegions) GetRegionId() *string {
	return s.RegionId
}

func (s *ListBusinessRegionsResponseBodyGeographicSubRegions) SetName(v string) *ListBusinessRegionsResponseBodyGeographicSubRegions {
	s.Name = &v
	return s
}

func (s *ListBusinessRegionsResponseBodyGeographicSubRegions) SetRegionId(v string) *ListBusinessRegionsResponseBodyGeographicSubRegions {
	s.RegionId = &v
	return s
}

func (s *ListBusinessRegionsResponseBodyGeographicSubRegions) Validate() error {
	return dara.Validate(s)
}
