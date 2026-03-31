// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePeakTrendRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTimestamp(v string) *DescribePeakTrendRequest
	GetEndTimestamp() *string
	SetInstanceId(v string) *DescribePeakTrendRequest
	GetInstanceId() *string
	SetInterval(v string) *DescribePeakTrendRequest
	GetInterval() *string
	SetRegionId(v string) *DescribePeakTrendRequest
	GetRegionId() *string
	SetResource(v string) *DescribePeakTrendRequest
	GetResource() *string
	SetResourceManagerResourceGroupId(v string) *DescribePeakTrendRequest
	GetResourceManagerResourceGroupId() *string
	SetStartTimestamp(v string) *DescribePeakTrendRequest
	GetStartTimestamp() *string
}

type DescribePeakTrendRequest struct {
	// The end of the time range to query. Unit: seconds. If you do not specify this parameter, the current time is used.
	//
	// example:
	//
	// 1665386340
	EndTimestamp *string `json:"EndTimestamp,omitempty" xml:"EndTimestamp,omitempty"`
	// The ID of the WAF instance.
	//
	// >  You can call the [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) operation to obtain the ID of the WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf_cdnsdf3****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The time interval. Unit: seconds. The value must be an integral multiple of 60.
	//
	// This parameter is required.
	//
	// example:
	//
	// 300
	Interval *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// The region where the WAF instance resides. Valid values:
	//
	// 	- **cn-hangzhou:*	- the Chinese mainland.
	//
	// 	- **ap-southeast-1:*	- outside the Chinese mainland.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The protected object.
	//
	// example:
	//
	// www.aliyundoc.com
	Resource *string `json:"Resource,omitempty" xml:"Resource,omitempty"`
	// The ID of the Alibaba Cloud resource group.
	//
	// example:
	//
	// rg-acfm***q
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
	// The beginning of the time range to query. Unit: seconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1665331200
	StartTimestamp *string `json:"StartTimestamp,omitempty" xml:"StartTimestamp,omitempty"`
}

func (s DescribePeakTrendRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePeakTrendRequest) GoString() string {
	return s.String()
}

func (s *DescribePeakTrendRequest) GetEndTimestamp() *string {
	return s.EndTimestamp
}

func (s *DescribePeakTrendRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribePeakTrendRequest) GetInterval() *string {
	return s.Interval
}

func (s *DescribePeakTrendRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribePeakTrendRequest) GetResource() *string {
	return s.Resource
}

func (s *DescribePeakTrendRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *DescribePeakTrendRequest) GetStartTimestamp() *string {
	return s.StartTimestamp
}

func (s *DescribePeakTrendRequest) SetEndTimestamp(v string) *DescribePeakTrendRequest {
	s.EndTimestamp = &v
	return s
}

func (s *DescribePeakTrendRequest) SetInstanceId(v string) *DescribePeakTrendRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribePeakTrendRequest) SetInterval(v string) *DescribePeakTrendRequest {
	s.Interval = &v
	return s
}

func (s *DescribePeakTrendRequest) SetRegionId(v string) *DescribePeakTrendRequest {
	s.RegionId = &v
	return s
}

func (s *DescribePeakTrendRequest) SetResource(v string) *DescribePeakTrendRequest {
	s.Resource = &v
	return s
}

func (s *DescribePeakTrendRequest) SetResourceManagerResourceGroupId(v string) *DescribePeakTrendRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribePeakTrendRequest) SetStartTimestamp(v string) *DescribePeakTrendRequest {
	s.StartTimestamp = &v
	return s
}

func (s *DescribePeakTrendRequest) Validate() error {
	return dara.Validate(s)
}
