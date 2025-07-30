// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRegionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRegions(v []*ListRegionsResponseBodyRegions) *ListRegionsResponseBody
	GetRegions() []*ListRegionsResponseBodyRegions
	SetRequestId(v string) *ListRegionsResponseBody
	GetRequestId() *string
}

type ListRegionsResponseBody struct {
	// The supported regions.
	Regions []*ListRegionsResponseBodyRegions `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListRegionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListRegionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListRegionsResponseBody) GetRegions() []*ListRegionsResponseBodyRegions {
	return s.Regions
}

func (s *ListRegionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListRegionsResponseBody) SetRegions(v []*ListRegionsResponseBodyRegions) *ListRegionsResponseBody {
	s.Regions = v
	return s
}

func (s *ListRegionsResponseBody) SetRequestId(v string) *ListRegionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRegionsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListRegionsResponseBodyRegions struct {
	// The name of the region.
	//
	// example:
	//
	// China (Hangzhou)
	LocalName *string `json:"LocalName,omitempty" xml:"LocalName,omitempty"`
	// The endpoint of the region.
	//
	// example:
	//
	// eiam.cn-hangzhou.aliyuncs.com
	RegionEndpoint *string `json:"RegionEndpoint,omitempty" xml:"RegionEndpoint,omitempty"`
	// The ID of the region.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListRegionsResponseBodyRegions) String() string {
	return dara.Prettify(s)
}

func (s ListRegionsResponseBodyRegions) GoString() string {
	return s.String()
}

func (s *ListRegionsResponseBodyRegions) GetLocalName() *string {
	return s.LocalName
}

func (s *ListRegionsResponseBodyRegions) GetRegionEndpoint() *string {
	return s.RegionEndpoint
}

func (s *ListRegionsResponseBodyRegions) GetRegionId() *string {
	return s.RegionId
}

func (s *ListRegionsResponseBodyRegions) SetLocalName(v string) *ListRegionsResponseBodyRegions {
	s.LocalName = &v
	return s
}

func (s *ListRegionsResponseBodyRegions) SetRegionEndpoint(v string) *ListRegionsResponseBodyRegions {
	s.RegionEndpoint = &v
	return s
}

func (s *ListRegionsResponseBodyRegions) SetRegionId(v string) *ListRegionsResponseBodyRegions {
	s.RegionId = &v
	return s
}

func (s *ListRegionsResponseBodyRegions) Validate() error {
	return dara.Validate(s)
}
