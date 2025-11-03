// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeActiveOperationTaskRegionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRegionList(v []*DescribeActiveOperationTaskRegionResponseBodyRegionList) *DescribeActiveOperationTaskRegionResponseBody
	GetRegionList() []*DescribeActiveOperationTaskRegionResponseBodyRegionList
	SetRequestId(v string) *DescribeActiveOperationTaskRegionResponseBody
	GetRequestId() *string
}

type DescribeActiveOperationTaskRegionResponseBody struct {
	// The region ID.
	RegionList []*DescribeActiveOperationTaskRegionResponseBodyRegionList `json:"RegionList,omitempty" xml:"RegionList,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 3C4A2494-XXXX-XXXX-93CF-548DB3375193
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeActiveOperationTaskRegionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeActiveOperationTaskRegionResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeActiveOperationTaskRegionResponseBody) GetRegionList() []*DescribeActiveOperationTaskRegionResponseBodyRegionList {
	return s.RegionList
}

func (s *DescribeActiveOperationTaskRegionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeActiveOperationTaskRegionResponseBody) SetRegionList(v []*DescribeActiveOperationTaskRegionResponseBodyRegionList) *DescribeActiveOperationTaskRegionResponseBody {
	s.RegionList = v
	return s
}

func (s *DescribeActiveOperationTaskRegionResponseBody) SetRequestId(v string) *DescribeActiveOperationTaskRegionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeActiveOperationTaskRegionResponseBody) Validate() error {
	if s.RegionList != nil {
		for _, item := range s.RegionList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeActiveOperationTaskRegionResponseBodyRegionList struct {
	// The total number of tasks.
	//
	// example:
	//
	// 1
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The region ID of the instance.
	//
	// example:
	//
	// cn-beijing
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
}

func (s DescribeActiveOperationTaskRegionResponseBodyRegionList) String() string {
	return dara.Prettify(s)
}

func (s DescribeActiveOperationTaskRegionResponseBodyRegionList) GoString() string {
	return s.String()
}

func (s *DescribeActiveOperationTaskRegionResponseBodyRegionList) GetCount() *int32 {
	return s.Count
}

func (s *DescribeActiveOperationTaskRegionResponseBodyRegionList) GetRegion() *string {
	return s.Region
}

func (s *DescribeActiveOperationTaskRegionResponseBodyRegionList) SetCount(v int32) *DescribeActiveOperationTaskRegionResponseBodyRegionList {
	s.Count = &v
	return s
}

func (s *DescribeActiveOperationTaskRegionResponseBodyRegionList) SetRegion(v string) *DescribeActiveOperationTaskRegionResponseBodyRegionList {
	s.Region = &v
	return s
}

func (s *DescribeActiveOperationTaskRegionResponseBodyRegionList) Validate() error {
	return dara.Validate(s)
}
