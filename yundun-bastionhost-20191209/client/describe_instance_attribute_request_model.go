// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribeInstanceAttributeRequest
	GetInstanceId() *string
	SetRegionId(v string) *DescribeInstanceAttributeRequest
	GetRegionId() *string
}

type DescribeInstanceAttributeRequest struct {
	// The ID of the Bastionhost instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// bastionhost-cn-78v1ghxxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the region.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeInstanceAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceAttributeRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstanceAttributeRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeInstanceAttributeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeInstanceAttributeRequest) SetInstanceId(v string) *DescribeInstanceAttributeRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstanceAttributeRequest) SetRegionId(v string) *DescribeInstanceAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeInstanceAttributeRequest) Validate() error {
	return dara.Validate(s)
}
