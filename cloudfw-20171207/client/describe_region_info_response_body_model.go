// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRegionInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataList(v []*DescribeRegionInfoResponseBodyDataList) *DescribeRegionInfoResponseBody
	GetDataList() []*DescribeRegionInfoResponseBodyDataList
	SetRequestId(v string) *DescribeRegionInfoResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeRegionInfoResponseBody
	GetTotalCount() *int32
}

type DescribeRegionInfoResponseBody struct {
	DataList []*DescribeRegionInfoResponseBodyDataList `json:"DataList,omitempty" xml:"DataList,omitempty" type:"Repeated"`
	// example:
	//
	// 25E655B0-CAED-53D4-8054-F983126****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 8
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeRegionInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRegionInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRegionInfoResponseBody) GetDataList() []*DescribeRegionInfoResponseBodyDataList {
	return s.DataList
}

func (s *DescribeRegionInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRegionInfoResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeRegionInfoResponseBody) SetDataList(v []*DescribeRegionInfoResponseBodyDataList) *DescribeRegionInfoResponseBody {
	s.DataList = v
	return s
}

func (s *DescribeRegionInfoResponseBody) SetRequestId(v string) *DescribeRegionInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRegionInfoResponseBody) SetTotalCount(v int32) *DescribeRegionInfoResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeRegionInfoResponseBody) Validate() error {
	if s.DataList != nil {
		for _, item := range s.DataList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeRegionInfoResponseBodyDataList struct {
	// example:
	//
	// cn-shanghai
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RegionName *string `json:"RegionName,omitempty" xml:"RegionName,omitempty"`
}

func (s DescribeRegionInfoResponseBodyDataList) String() string {
	return dara.Prettify(s)
}

func (s DescribeRegionInfoResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *DescribeRegionInfoResponseBodyDataList) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeRegionInfoResponseBodyDataList) GetRegionName() *string {
	return s.RegionName
}

func (s *DescribeRegionInfoResponseBodyDataList) SetRegionId(v string) *DescribeRegionInfoResponseBodyDataList {
	s.RegionId = &v
	return s
}

func (s *DescribeRegionInfoResponseBodyDataList) SetRegionName(v string) *DescribeRegionInfoResponseBodyDataList {
	s.RegionName = &v
	return s
}

func (s *DescribeRegionInfoResponseBodyDataList) Validate() error {
	return dara.Validate(s)
}
