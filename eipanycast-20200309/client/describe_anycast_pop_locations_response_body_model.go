// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAnycastPopLocationsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAnycastPopLocationList(v []*DescribeAnycastPopLocationsResponseBodyAnycastPopLocationList) *DescribeAnycastPopLocationsResponseBody
	GetAnycastPopLocationList() []*DescribeAnycastPopLocationsResponseBodyAnycastPopLocationList
	SetCount(v string) *DescribeAnycastPopLocationsResponseBody
	GetCount() *string
	SetRequestId(v string) *DescribeAnycastPopLocationsResponseBody
	GetRequestId() *string
}

type DescribeAnycastPopLocationsResponseBody struct {
	// The list of access points in the specified access area.
	AnycastPopLocationList []*DescribeAnycastPopLocationsResponseBodyAnycastPopLocationList `json:"AnycastPopLocationList,omitempty" xml:"AnycastPopLocationList,omitempty" type:"Repeated"`
	// The number of access points.
	//
	// example:
	//
	// 1
	Count *string `json:"Count,omitempty" xml:"Count,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 4EC47282-1B74-4534-BD0E-403F3EE64CAF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeAnycastPopLocationsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAnycastPopLocationsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAnycastPopLocationsResponseBody) GetAnycastPopLocationList() []*DescribeAnycastPopLocationsResponseBodyAnycastPopLocationList {
	return s.AnycastPopLocationList
}

func (s *DescribeAnycastPopLocationsResponseBody) GetCount() *string {
	return s.Count
}

func (s *DescribeAnycastPopLocationsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAnycastPopLocationsResponseBody) SetAnycastPopLocationList(v []*DescribeAnycastPopLocationsResponseBodyAnycastPopLocationList) *DescribeAnycastPopLocationsResponseBody {
	s.AnycastPopLocationList = v
	return s
}

func (s *DescribeAnycastPopLocationsResponseBody) SetCount(v string) *DescribeAnycastPopLocationsResponseBody {
	s.Count = &v
	return s
}

func (s *DescribeAnycastPopLocationsResponseBody) SetRequestId(v string) *DescribeAnycastPopLocationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAnycastPopLocationsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeAnycastPopLocationsResponseBodyAnycastPopLocationList struct {
	// The ID of the region where the access point is deployed.
	//
	// example:
	//
	// us-west-1-pop
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the region where the access point is deployed.
	//
	// example:
	//
	// us-west-1-pop
	RegionName *string `json:"RegionName,omitempty" xml:"RegionName,omitempty"`
}

func (s DescribeAnycastPopLocationsResponseBodyAnycastPopLocationList) String() string {
	return dara.Prettify(s)
}

func (s DescribeAnycastPopLocationsResponseBodyAnycastPopLocationList) GoString() string {
	return s.String()
}

func (s *DescribeAnycastPopLocationsResponseBodyAnycastPopLocationList) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeAnycastPopLocationsResponseBodyAnycastPopLocationList) GetRegionName() *string {
	return s.RegionName
}

func (s *DescribeAnycastPopLocationsResponseBodyAnycastPopLocationList) SetRegionId(v string) *DescribeAnycastPopLocationsResponseBodyAnycastPopLocationList {
	s.RegionId = &v
	return s
}

func (s *DescribeAnycastPopLocationsResponseBodyAnycastPopLocationList) SetRegionName(v string) *DescribeAnycastPopLocationsResponseBodyAnycastPopLocationList {
	s.RegionName = &v
	return s
}

func (s *DescribeAnycastPopLocationsResponseBodyAnycastPopLocationList) Validate() error {
	return dara.Validate(s)
}
