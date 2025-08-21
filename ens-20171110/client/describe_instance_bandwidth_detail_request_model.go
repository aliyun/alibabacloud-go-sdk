// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceBandwidthDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *DescribeInstanceBandwidthDetailRequest
	GetEndTime() *string
	SetEnsRegionId(v string) *DescribeInstanceBandwidthDetailRequest
	GetEnsRegionId() *string
	SetInstanceId(v string) *DescribeInstanceBandwidthDetailRequest
	GetInstanceId() *string
	SetInstanceType(v string) *DescribeInstanceBandwidthDetailRequest
	GetInstanceType() *string
	SetPageNumber(v int32) *DescribeInstanceBandwidthDetailRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeInstanceBandwidthDetailRequest
	GetPageSize() *int32
	SetServiceType(v string) *DescribeInstanceBandwidthDetailRequest
	GetServiceType() *string
	SetStartTime(v string) *DescribeInstanceBandwidthDetailRequest
	GetStartTime() *string
}

type DescribeInstanceBandwidthDetailRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 2024-01-11 30:00:00
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// cn-beijing-cmcc
	EnsRegionId *string `json:"EnsRegionId,omitempty" xml:"EnsRegionId,omitempty"`
	// example:
	//
	// i-6ecpqvkicnchxccozrpxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// vm
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: 200.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// vm
	ServiceType *string `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2024-01-11 00:00:00
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeInstanceBandwidthDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceBandwidthDetailRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstanceBandwidthDetailRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeInstanceBandwidthDetailRequest) GetEnsRegionId() *string {
	return s.EnsRegionId
}

func (s *DescribeInstanceBandwidthDetailRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeInstanceBandwidthDetailRequest) GetInstanceType() *string {
	return s.InstanceType
}

func (s *DescribeInstanceBandwidthDetailRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeInstanceBandwidthDetailRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeInstanceBandwidthDetailRequest) GetServiceType() *string {
	return s.ServiceType
}

func (s *DescribeInstanceBandwidthDetailRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeInstanceBandwidthDetailRequest) SetEndTime(v string) *DescribeInstanceBandwidthDetailRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeInstanceBandwidthDetailRequest) SetEnsRegionId(v string) *DescribeInstanceBandwidthDetailRequest {
	s.EnsRegionId = &v
	return s
}

func (s *DescribeInstanceBandwidthDetailRequest) SetInstanceId(v string) *DescribeInstanceBandwidthDetailRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstanceBandwidthDetailRequest) SetInstanceType(v string) *DescribeInstanceBandwidthDetailRequest {
	s.InstanceType = &v
	return s
}

func (s *DescribeInstanceBandwidthDetailRequest) SetPageNumber(v int32) *DescribeInstanceBandwidthDetailRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeInstanceBandwidthDetailRequest) SetPageSize(v int32) *DescribeInstanceBandwidthDetailRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeInstanceBandwidthDetailRequest) SetServiceType(v string) *DescribeInstanceBandwidthDetailRequest {
	s.ServiceType = &v
	return s
}

func (s *DescribeInstanceBandwidthDetailRequest) SetStartTime(v string) *DescribeInstanceBandwidthDetailRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeInstanceBandwidthDetailRequest) Validate() error {
	return dara.Validate(s)
}
