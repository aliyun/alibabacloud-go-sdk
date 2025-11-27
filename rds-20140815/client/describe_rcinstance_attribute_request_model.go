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
	// example:
	//
	// k8s-node
	InstanceName     *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	MaxDisksResults  *int64  `json:"MaxDisksResults,omitempty" xml:"MaxDisksResults,omitempty"`
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
