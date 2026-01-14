// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeIpSetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIpSetId(v string) *DescribeIpSetRequest
	GetIpSetId() *string
	SetRegionId(v string) *DescribeIpSetRequest
	GetRegionId() *string
}

type DescribeIpSetRequest struct {
	// The ID of the acceleration region.
	//
	// You can call the [ListIpSets](https://help.aliyun.com/document_detail/2253273.html) operation to query the IDs of acceleration regions of a specific GA instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// ips-bp11ilwqjdkjeg9r7****
	IpSetId *string `json:"IpSetId,omitempty" xml:"IpSetId,omitempty"`
	// The ID of the region where the Global Accelerator (GA) instance is deployed. Set the value to **cn-hangzhou**.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeIpSetRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeIpSetRequest) GoString() string {
	return s.String()
}

func (s *DescribeIpSetRequest) GetIpSetId() *string {
	return s.IpSetId
}

func (s *DescribeIpSetRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeIpSetRequest) SetIpSetId(v string) *DescribeIpSetRequest {
	s.IpSetId = &v
	return s
}

func (s *DescribeIpSetRequest) SetRegionId(v string) *DescribeIpSetRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeIpSetRequest) Validate() error {
	return dara.Validate(s)
}
