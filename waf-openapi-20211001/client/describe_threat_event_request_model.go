// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeThreatEventRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeThreatEventRequest
	GetDomainName() *string
	SetEndTime(v int64) *DescribeThreatEventRequest
	GetEndTime() *int64
	SetInstanceId(v string) *DescribeThreatEventRequest
	GetInstanceId() *string
	SetPageNumber(v int32) *DescribeThreatEventRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeThreatEventRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeThreatEventRequest
	GetRegionId() *string
	SetResourceManagerResourceGroupId(v string) *DescribeThreatEventRequest
	GetResourceManagerResourceGroupId() *string
	SetStartTime(v int64) *DescribeThreatEventRequest
	GetStartTime() *int64
}

type DescribeThreatEventRequest struct {
	// example:
	//
	// www.abc.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1749916800000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// waf_v2_public_cn-2bl4****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// rg-aeky65ka*****
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1668496310000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeThreatEventRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeThreatEventRequest) GoString() string {
	return s.String()
}

func (s *DescribeThreatEventRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeThreatEventRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeThreatEventRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeThreatEventRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeThreatEventRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeThreatEventRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeThreatEventRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *DescribeThreatEventRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeThreatEventRequest) SetDomainName(v string) *DescribeThreatEventRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeThreatEventRequest) SetEndTime(v int64) *DescribeThreatEventRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeThreatEventRequest) SetInstanceId(v string) *DescribeThreatEventRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeThreatEventRequest) SetPageNumber(v int32) *DescribeThreatEventRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeThreatEventRequest) SetPageSize(v int32) *DescribeThreatEventRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeThreatEventRequest) SetRegionId(v string) *DescribeThreatEventRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeThreatEventRequest) SetResourceManagerResourceGroupId(v string) *DescribeThreatEventRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribeThreatEventRequest) SetStartTime(v int64) *DescribeThreatEventRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeThreatEventRequest) Validate() error {
	return dara.Validate(s)
}
