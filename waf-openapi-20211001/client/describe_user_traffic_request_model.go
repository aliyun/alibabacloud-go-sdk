// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUserTrafficRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTimestamp(v int64) *DescribeUserTrafficRequest
	GetEndTimestamp() *int64
	SetInstanceId(v string) *DescribeUserTrafficRequest
	GetInstanceId() *string
	SetInterval(v int64) *DescribeUserTrafficRequest
	GetInterval() *int64
	SetRegionId(v string) *DescribeUserTrafficRequest
	GetRegionId() *string
	SetResourceManagerResourceGroupId(v string) *DescribeUserTrafficRequest
	GetResourceManagerResourceGroupId() *string
	SetStartTimestamp(v int64) *DescribeUserTrafficRequest
	GetStartTimestamp() *int64
	SetType(v string) *DescribeUserTrafficRequest
	GetType() *string
}

type DescribeUserTrafficRequest struct {
	// example:
	//
	// 1665386280
	EndTimestamp *int64 `json:"EndTimestamp,omitempty" xml:"EndTimestamp,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// waf_v3prepaid_public_cn-***
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 3600
	Interval *int64 `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// rg-ae*******i
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
	// example:
	//
	// 1665331200
	StartTimestamp *int64 `json:"StartTimestamp,omitempty" xml:"StartTimestamp,omitempty"`
	// example:
	//
	// qps
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeUserTrafficRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserTrafficRequest) GoString() string {
	return s.String()
}

func (s *DescribeUserTrafficRequest) GetEndTimestamp() *int64 {
	return s.EndTimestamp
}

func (s *DescribeUserTrafficRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeUserTrafficRequest) GetInterval() *int64 {
	return s.Interval
}

func (s *DescribeUserTrafficRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeUserTrafficRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *DescribeUserTrafficRequest) GetStartTimestamp() *int64 {
	return s.StartTimestamp
}

func (s *DescribeUserTrafficRequest) GetType() *string {
	return s.Type
}

func (s *DescribeUserTrafficRequest) SetEndTimestamp(v int64) *DescribeUserTrafficRequest {
	s.EndTimestamp = &v
	return s
}

func (s *DescribeUserTrafficRequest) SetInstanceId(v string) *DescribeUserTrafficRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeUserTrafficRequest) SetInterval(v int64) *DescribeUserTrafficRequest {
	s.Interval = &v
	return s
}

func (s *DescribeUserTrafficRequest) SetRegionId(v string) *DescribeUserTrafficRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeUserTrafficRequest) SetResourceManagerResourceGroupId(v string) *DescribeUserTrafficRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribeUserTrafficRequest) SetStartTimestamp(v int64) *DescribeUserTrafficRequest {
	s.StartTimestamp = &v
	return s
}

func (s *DescribeUserTrafficRequest) SetType(v string) *DescribeUserTrafficRequest {
	s.Type = &v
	return s
}

func (s *DescribeUserTrafficRequest) Validate() error {
	return dara.Validate(s)
}
