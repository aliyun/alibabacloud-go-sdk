// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRCInstanceAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribeRCInstanceAttributeRequest
	GetInstanceId() *string
	SetInstanceName(v string) *DescribeRCInstanceAttributeRequest
	GetInstanceName() *string
	SetMaxDisksResults(v int64) *DescribeRCInstanceAttributeRequest
	GetMaxDisksResults() *int64
	SetPrivateIpAddress(v string) *DescribeRCInstanceAttributeRequest
	GetPrivateIpAddress() *string
	SetRegionId(v string) *DescribeRCInstanceAttributeRequest
	GetRegionId() *string
}

type DescribeRCInstanceAttributeRequest struct {
	// The instance ID.
	//
	// example:
	//
	// rc-dh2jf9n6j4s14926****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The instance name
	//
	// example:
	//
	// k8s-node
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// Set the upper limit for the number of instance disks in the return result. The valid range is 10 to 500.
	//
	// - If no value is set, the default value is 20.
	//
	// - If the set value is less than 10, it is fixed to 10.
	//
	// - If the set value is greater than or equal to 10 and less than or equal to 500, the set value is used.
	//
	// example:
	//
	// 20
	MaxDisksResults *int64 `json:"MaxDisksResults,omitempty" xml:"MaxDisksResults,omitempty"`
	// The VPC network IP address of the instance, that is, the private IP address.
	//
	// example:
	//
	// 192.168.XXX.XXX
	PrivateIpAddress *string `json:"PrivateIpAddress,omitempty" xml:"PrivateIpAddress,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeRCInstanceAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRCInstanceAttributeRequest) GoString() string {
	return s.String()
}

func (s *DescribeRCInstanceAttributeRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeRCInstanceAttributeRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DescribeRCInstanceAttributeRequest) GetMaxDisksResults() *int64 {
	return s.MaxDisksResults
}

func (s *DescribeRCInstanceAttributeRequest) GetPrivateIpAddress() *string {
	return s.PrivateIpAddress
}

func (s *DescribeRCInstanceAttributeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeRCInstanceAttributeRequest) SetInstanceId(v string) *DescribeRCInstanceAttributeRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeRCInstanceAttributeRequest) SetInstanceName(v string) *DescribeRCInstanceAttributeRequest {
	s.InstanceName = &v
	return s
}

func (s *DescribeRCInstanceAttributeRequest) SetMaxDisksResults(v int64) *DescribeRCInstanceAttributeRequest {
	s.MaxDisksResults = &v
	return s
}

func (s *DescribeRCInstanceAttributeRequest) SetPrivateIpAddress(v string) *DescribeRCInstanceAttributeRequest {
	s.PrivateIpAddress = &v
	return s
}

func (s *DescribeRCInstanceAttributeRequest) SetRegionId(v string) *DescribeRCInstanceAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeRCInstanceAttributeRequest) Validate() error {
	return dara.Validate(s)
}
