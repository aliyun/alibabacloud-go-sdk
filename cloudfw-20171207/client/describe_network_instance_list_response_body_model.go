// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNetworkInstanceListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNetworkInstanceList(v []*DescribeNetworkInstanceListResponseBodyNetworkInstanceList) *DescribeNetworkInstanceListResponseBody
	GetNetworkInstanceList() []*DescribeNetworkInstanceListResponseBodyNetworkInstanceList
	SetRequestId(v string) *DescribeNetworkInstanceListResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeNetworkInstanceListResponseBody
	GetTotalCount() *int32
}

type DescribeNetworkInstanceListResponseBody struct {
	NetworkInstanceList []*DescribeNetworkInstanceListResponseBodyNetworkInstanceList `json:"NetworkInstanceList,omitempty" xml:"NetworkInstanceList,omitempty" type:"Repeated"`
	// example:
	//
	// D2373503-3921-59F2-93A6-3DA7FB7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 21
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeNetworkInstanceListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeNetworkInstanceListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeNetworkInstanceListResponseBody) GetNetworkInstanceList() []*DescribeNetworkInstanceListResponseBodyNetworkInstanceList {
	return s.NetworkInstanceList
}

func (s *DescribeNetworkInstanceListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeNetworkInstanceListResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeNetworkInstanceListResponseBody) SetNetworkInstanceList(v []*DescribeNetworkInstanceListResponseBodyNetworkInstanceList) *DescribeNetworkInstanceListResponseBody {
	s.NetworkInstanceList = v
	return s
}

func (s *DescribeNetworkInstanceListResponseBody) SetRequestId(v string) *DescribeNetworkInstanceListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeNetworkInstanceListResponseBody) SetTotalCount(v int32) *DescribeNetworkInstanceListResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeNetworkInstanceListResponseBody) Validate() error {
	if s.NetworkInstanceList != nil {
		for _, item := range s.NetworkInstanceList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeNetworkInstanceListResponseBodyNetworkInstanceList struct {
	// example:
	//
	// vpc-m5ewlqkuf7or****
	NetworkInstanceId *string `json:"NetworkInstanceId,omitempty" xml:"NetworkInstanceId,omitempty"`
	// example:
	//
	// vpc-test
	NetworkInstanceName *string `json:"NetworkInstanceName,omitempty" xml:"NetworkInstanceName,omitempty"`
	// example:
	//
	// vpc
	NetworkInstanceType *string `json:"NetworkInstanceType,omitempty" xml:"NetworkInstanceType,omitempty"`
	// example:
	//
	// cn-shanghai
	RegionNo *string `json:"RegionNo,omitempty" xml:"RegionNo,omitempty"`
}

func (s DescribeNetworkInstanceListResponseBodyNetworkInstanceList) String() string {
	return dara.Prettify(s)
}

func (s DescribeNetworkInstanceListResponseBodyNetworkInstanceList) GoString() string {
	return s.String()
}

func (s *DescribeNetworkInstanceListResponseBodyNetworkInstanceList) GetNetworkInstanceId() *string {
	return s.NetworkInstanceId
}

func (s *DescribeNetworkInstanceListResponseBodyNetworkInstanceList) GetNetworkInstanceName() *string {
	return s.NetworkInstanceName
}

func (s *DescribeNetworkInstanceListResponseBodyNetworkInstanceList) GetNetworkInstanceType() *string {
	return s.NetworkInstanceType
}

func (s *DescribeNetworkInstanceListResponseBodyNetworkInstanceList) GetRegionNo() *string {
	return s.RegionNo
}

func (s *DescribeNetworkInstanceListResponseBodyNetworkInstanceList) SetNetworkInstanceId(v string) *DescribeNetworkInstanceListResponseBodyNetworkInstanceList {
	s.NetworkInstanceId = &v
	return s
}

func (s *DescribeNetworkInstanceListResponseBodyNetworkInstanceList) SetNetworkInstanceName(v string) *DescribeNetworkInstanceListResponseBodyNetworkInstanceList {
	s.NetworkInstanceName = &v
	return s
}

func (s *DescribeNetworkInstanceListResponseBodyNetworkInstanceList) SetNetworkInstanceType(v string) *DescribeNetworkInstanceListResponseBodyNetworkInstanceList {
	s.NetworkInstanceType = &v
	return s
}

func (s *DescribeNetworkInstanceListResponseBodyNetworkInstanceList) SetRegionNo(v string) *DescribeNetworkInstanceListResponseBodyNetworkInstanceList {
	s.RegionNo = &v
	return s
}

func (s *DescribeNetworkInstanceListResponseBodyNetworkInstanceList) Validate() error {
	return dara.Validate(s)
}
