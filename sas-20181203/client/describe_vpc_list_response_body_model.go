// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVpcListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCount(v int32) *DescribeVpcListResponseBody
	GetCount() *int32
	SetRequestId(v string) *DescribeVpcListResponseBody
	GetRequestId() *string
	SetVpcList(v []*DescribeVpcListResponseBodyVpcList) *DescribeVpcListResponseBody
	GetVpcList() []*DescribeVpcListResponseBodyVpcList
}

type DescribeVpcListResponseBody struct {
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 7E0618A9-D5EF-4220-9471-C42B5E92719F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// An array that consists of VPCs.
	VpcList []*DescribeVpcListResponseBodyVpcList `json:"VpcList,omitempty" xml:"VpcList,omitempty" type:"Repeated"`
}

func (s DescribeVpcListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpcListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVpcListResponseBody) GetCount() *int32 {
	return s.Count
}

func (s *DescribeVpcListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVpcListResponseBody) GetVpcList() []*DescribeVpcListResponseBodyVpcList {
	return s.VpcList
}

func (s *DescribeVpcListResponseBody) SetCount(v int32) *DescribeVpcListResponseBody {
	s.Count = &v
	return s
}

func (s *DescribeVpcListResponseBody) SetRequestId(v string) *DescribeVpcListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVpcListResponseBody) SetVpcList(v []*DescribeVpcListResponseBodyVpcList) *DescribeVpcListResponseBody {
	s.VpcList = v
	return s
}

func (s *DescribeVpcListResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeVpcListResponseBodyVpcList struct {
	// The number of Elastic Compute Service (ECS) instances.
	//
	// example:
	//
	// 9
	EcsCount *int32 `json:"EcsCount,omitempty" xml:"EcsCount,omitempty"`
	// The information about the virtual private cloud (VPC).
	//
	// example:
	//
	// TestVpcNote
	InstanceDesc *string `json:"InstanceDesc,omitempty" xml:"InstanceDesc,omitempty"`
	// The ID of the ECS instance.
	//
	// example:
	//
	// ins_1321_asedb_ada
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the VPC.
	//
	// example:
	//
	// test
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The region in which the server resides.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeVpcListResponseBodyVpcList) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpcListResponseBodyVpcList) GoString() string {
	return s.String()
}

func (s *DescribeVpcListResponseBodyVpcList) GetEcsCount() *int32 {
	return s.EcsCount
}

func (s *DescribeVpcListResponseBodyVpcList) GetInstanceDesc() *string {
	return s.InstanceDesc
}

func (s *DescribeVpcListResponseBodyVpcList) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeVpcListResponseBodyVpcList) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DescribeVpcListResponseBodyVpcList) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeVpcListResponseBodyVpcList) SetEcsCount(v int32) *DescribeVpcListResponseBodyVpcList {
	s.EcsCount = &v
	return s
}

func (s *DescribeVpcListResponseBodyVpcList) SetInstanceDesc(v string) *DescribeVpcListResponseBodyVpcList {
	s.InstanceDesc = &v
	return s
}

func (s *DescribeVpcListResponseBodyVpcList) SetInstanceId(v string) *DescribeVpcListResponseBodyVpcList {
	s.InstanceId = &v
	return s
}

func (s *DescribeVpcListResponseBodyVpcList) SetInstanceName(v string) *DescribeVpcListResponseBodyVpcList {
	s.InstanceName = &v
	return s
}

func (s *DescribeVpcListResponseBodyVpcList) SetRegionId(v string) *DescribeVpcListResponseBodyVpcList {
	s.RegionId = &v
	return s
}

func (s *DescribeVpcListResponseBodyVpcList) Validate() error {
	return dara.Validate(s)
}
