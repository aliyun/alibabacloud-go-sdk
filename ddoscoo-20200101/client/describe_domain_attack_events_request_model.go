// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainAttackEventsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomain(v string) *DescribeDomainAttackEventsRequest
	GetDomain() *string
	SetEndTime(v int64) *DescribeDomainAttackEventsRequest
	GetEndTime() *int64
	SetPageNumber(v int32) *DescribeDomainAttackEventsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeDomainAttackEventsRequest
	GetPageSize() *int32
	SetResourceGroupId(v string) *DescribeDomainAttackEventsRequest
	GetResourceGroupId() *string
	SetStartTime(v int64) *DescribeDomainAttackEventsRequest
	GetStartTime() *int64
}

type DescribeDomainAttackEventsRequest struct {
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
	// The number of the page to return. For example, to query the returned results on the first page, set the value to **1**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
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
}

func (s DescribeDomainAttackEventsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainAttackEventsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainAttackEventsRequest) GetDomain() *string {
	return s.Domain
}

func (s *DescribeDomainAttackEventsRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeDomainAttackEventsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeDomainAttackEventsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeDomainAttackEventsRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeDomainAttackEventsRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeDomainAttackEventsRequest) SetDomain(v string) *DescribeDomainAttackEventsRequest {
	s.Domain = &v
	return s
}

func (s *DescribeDomainAttackEventsRequest) SetEndTime(v int64) *DescribeDomainAttackEventsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDomainAttackEventsRequest) SetPageNumber(v int32) *DescribeDomainAttackEventsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDomainAttackEventsRequest) SetPageSize(v int32) *DescribeDomainAttackEventsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDomainAttackEventsRequest) SetResourceGroupId(v string) *DescribeDomainAttackEventsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDomainAttackEventsRequest) SetStartTime(v int64) *DescribeDomainAttackEventsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDomainAttackEventsRequest) Validate() error {
	return dara.Validate(s)
}
