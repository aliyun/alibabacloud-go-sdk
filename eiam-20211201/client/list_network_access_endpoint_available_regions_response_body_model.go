// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNetworkAccessEndpointAvailableRegionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRegions(v []*ListNetworkAccessEndpointAvailableRegionsResponseBodyRegions) *ListNetworkAccessEndpointAvailableRegionsResponseBody
	GetRegions() []*ListNetworkAccessEndpointAvailableRegionsResponseBodyRegions
	SetRequestId(v string) *ListNetworkAccessEndpointAvailableRegionsResponseBody
	GetRequestId() *string
}

type ListNetworkAccessEndpointAvailableRegionsResponseBody struct {
	// The information of region.
	Regions []*ListNetworkAccessEndpointAvailableRegionsResponseBodyRegions `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListNetworkAccessEndpointAvailableRegionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListNetworkAccessEndpointAvailableRegionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListNetworkAccessEndpointAvailableRegionsResponseBody) GetRegions() []*ListNetworkAccessEndpointAvailableRegionsResponseBodyRegions {
	return s.Regions
}

func (s *ListNetworkAccessEndpointAvailableRegionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListNetworkAccessEndpointAvailableRegionsResponseBody) SetRegions(v []*ListNetworkAccessEndpointAvailableRegionsResponseBodyRegions) *ListNetworkAccessEndpointAvailableRegionsResponseBody {
	s.Regions = v
	return s
}

func (s *ListNetworkAccessEndpointAvailableRegionsResponseBody) SetRequestId(v string) *ListNetworkAccessEndpointAvailableRegionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListNetworkAccessEndpointAvailableRegionsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListNetworkAccessEndpointAvailableRegionsResponseBodyRegions struct {
	// The name of the region.
	//
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

func (s ListNetworkAccessEndpointAvailableRegionsResponseBodyRegions) String() string {
	return dara.Prettify(s)
}

func (s ListNetworkAccessEndpointAvailableRegionsResponseBodyRegions) GoString() string {
	return s.String()
}

func (s *ListNetworkAccessEndpointAvailableRegionsResponseBodyRegions) GetLocalName() *string {
	return s.LocalName
}

func (s *ListNetworkAccessEndpointAvailableRegionsResponseBodyRegions) GetRegionId() *string {
	return s.RegionId
}

func (s *ListNetworkAccessEndpointAvailableRegionsResponseBodyRegions) SetLocalName(v string) *ListNetworkAccessEndpointAvailableRegionsResponseBodyRegions {
	s.LocalName = &v
	return s
}

func (s *ListNetworkAccessEndpointAvailableRegionsResponseBodyRegions) SetRegionId(v string) *ListNetworkAccessEndpointAvailableRegionsResponseBodyRegions {
	s.RegionId = &v
	return s
}

func (s *ListNetworkAccessEndpointAvailableRegionsResponseBodyRegions) Validate() error {
	return dara.Validate(s)
}
