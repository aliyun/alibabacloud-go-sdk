// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainViewTopCostTimeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomain(v string) *DescribeDomainViewTopCostTimeRequest
	GetDomain() *string
	SetEndTime(v int64) *DescribeDomainViewTopCostTimeRequest
	GetEndTime() *int64
	SetResourceGroupId(v string) *DescribeDomainViewTopCostTimeRequest
	GetResourceGroupId() *string
	SetStartTime(v int64) *DescribeDomainViewTopCostTimeRequest
	GetStartTime() *int64
	SetTop(v int32) *DescribeDomainViewTopCostTimeRequest
	GetTop() *int32
}

type DescribeDomainViewTopCostTimeRequest struct {
	// The domain name of the website.
	//
	// > A forwarding rule must be configured for the domain name. You can call the [DescribeDomains](https://help.aliyun.com/document_detail/91724.html) operation to query all domain names.
	//
	// example:
	//
	// www.aliyun.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The end of the time range to query. The value is a UNIX timestamp. Unit: seconds.
	//
	// > This UNIX timestamp must indicate a point in time that is accurate to the minute.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1583683200
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The ID of the resource group to which the instance belongs in Resource Management. This parameter is empty by default, which indicates that the instance belongs to the default resource group.
	//
	// example:
	//
	// default
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The beginning of the time range to query. The value is a UNIX timestamp. Unit: seconds.
	//
	// > This UNIX timestamp must indicate a point in time that is accurate to the minute.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1582992000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The number of URLs to query. Valid values: **1*	- to **100**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 5
	Top *int32 `json:"Top,omitempty" xml:"Top,omitempty"`
}

func (s DescribeDomainViewTopCostTimeRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainViewTopCostTimeRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainViewTopCostTimeRequest) GetDomain() *string {
	return s.Domain
}

func (s *DescribeDomainViewTopCostTimeRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeDomainViewTopCostTimeRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeDomainViewTopCostTimeRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeDomainViewTopCostTimeRequest) GetTop() *int32 {
	return s.Top
}

func (s *DescribeDomainViewTopCostTimeRequest) SetDomain(v string) *DescribeDomainViewTopCostTimeRequest {
	s.Domain = &v
	return s
}

func (s *DescribeDomainViewTopCostTimeRequest) SetEndTime(v int64) *DescribeDomainViewTopCostTimeRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDomainViewTopCostTimeRequest) SetResourceGroupId(v string) *DescribeDomainViewTopCostTimeRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDomainViewTopCostTimeRequest) SetStartTime(v int64) *DescribeDomainViewTopCostTimeRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDomainViewTopCostTimeRequest) SetTop(v int32) *DescribeDomainViewTopCostTimeRequest {
	s.Top = &v
	return s
}

func (s *DescribeDomainViewTopCostTimeRequest) Validate() error {
	return dara.Validate(s)
}
