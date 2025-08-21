// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUdpReflectRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribeUdpReflectRequest
	GetInstanceId() *string
	SetRegionId(v string) *DescribeUdpReflectRequest
	GetRegionId() *string
}

type DescribeUdpReflectRequest struct {
	// The ID of the instance to query.
	//
	// > You can call the [DescribeInstanceIds](https://help.aliyun.com/document_detail/157459.html) operation to query the IDs of all instances.
	//
	// This parameter is required.
	//
	// example:
	//
	// ddoscoo-cn-i7m25564****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the instance. Valid values:
	//
	// 	- **cn-hangzhou**: indicates an Anti-DDoS Proxy (Chinese Mainland) instance. This is the default value.
	//
	// 	- **ap-southeast-1**: indicates an Anti-DDoS Proxy (Outside Chinese Mainland) instance.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeUdpReflectRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeUdpReflectRequest) GoString() string {
	return s.String()
}

func (s *DescribeUdpReflectRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeUdpReflectRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeUdpReflectRequest) SetInstanceId(v string) *DescribeUdpReflectRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeUdpReflectRequest) SetRegionId(v string) *DescribeUdpReflectRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeUdpReflectRequest) Validate() error {
	return dara.Validate(s)
}
