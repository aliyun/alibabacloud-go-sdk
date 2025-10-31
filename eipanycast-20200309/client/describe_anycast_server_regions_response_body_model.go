// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAnycastServerRegionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAnycastServerRegionList(v []*DescribeAnycastServerRegionsResponseBodyAnycastServerRegionList) *DescribeAnycastServerRegionsResponseBody
	GetAnycastServerRegionList() []*DescribeAnycastServerRegionsResponseBodyAnycastServerRegionList
	SetCount(v string) *DescribeAnycastServerRegionsResponseBody
	GetCount() *string
	SetRequestId(v string) *DescribeAnycastServerRegionsResponseBody
	GetRequestId() *string
}

type DescribeAnycastServerRegionsResponseBody struct {
	// The list of regions where you can associate Anycast EIPs with endpoints.
	AnycastServerRegionList []*DescribeAnycastServerRegionsResponseBodyAnycastServerRegionList `json:"AnycastServerRegionList,omitempty" xml:"AnycastServerRegionList,omitempty" type:"Repeated"`
	// The number of returned entries.
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

func (s DescribeAnycastServerRegionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAnycastServerRegionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAnycastServerRegionsResponseBody) GetAnycastServerRegionList() []*DescribeAnycastServerRegionsResponseBodyAnycastServerRegionList {
	return s.AnycastServerRegionList
}

func (s *DescribeAnycastServerRegionsResponseBody) GetCount() *string {
	return s.Count
}

func (s *DescribeAnycastServerRegionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAnycastServerRegionsResponseBody) SetAnycastServerRegionList(v []*DescribeAnycastServerRegionsResponseBodyAnycastServerRegionList) *DescribeAnycastServerRegionsResponseBody {
	s.AnycastServerRegionList = v
	return s
}

func (s *DescribeAnycastServerRegionsResponseBody) SetCount(v string) *DescribeAnycastServerRegionsResponseBody {
	s.Count = &v
	return s
}

func (s *DescribeAnycastServerRegionsResponseBody) SetRequestId(v string) *DescribeAnycastServerRegionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAnycastServerRegionsResponseBody) Validate() error {
	if s.AnycastServerRegionList != nil {
		for _, item := range s.AnycastServerRegionList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeAnycastServerRegionsResponseBodyAnycastServerRegionList struct {
	// The ID of the region.
	//
	// example:
	//
	// eu-west-1
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the region.
	//
	// example:
	//
	// eu-west-1-gb33-a01
	RegionName *string `json:"RegionName,omitempty" xml:"RegionName,omitempty"`
}

func (s DescribeAnycastServerRegionsResponseBodyAnycastServerRegionList) String() string {
	return dara.Prettify(s)
}

func (s DescribeAnycastServerRegionsResponseBodyAnycastServerRegionList) GoString() string {
	return s.String()
}

func (s *DescribeAnycastServerRegionsResponseBodyAnycastServerRegionList) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeAnycastServerRegionsResponseBodyAnycastServerRegionList) GetRegionName() *string {
	return s.RegionName
}

func (s *DescribeAnycastServerRegionsResponseBodyAnycastServerRegionList) SetRegionId(v string) *DescribeAnycastServerRegionsResponseBodyAnycastServerRegionList {
	s.RegionId = &v
	return s
}

func (s *DescribeAnycastServerRegionsResponseBodyAnycastServerRegionList) SetRegionName(v string) *DescribeAnycastServerRegionsResponseBodyAnycastServerRegionList {
	s.RegionName = &v
	return s
}

func (s *DescribeAnycastServerRegionsResponseBodyAnycastServerRegionList) Validate() error {
	return dara.Validate(s)
}
