// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainQPSListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomain(v string) *DescribeDomainQPSListRequest
	GetDomain() *string
	SetEndTime(v int64) *DescribeDomainQPSListRequest
	GetEndTime() *int64
	SetInterval(v int64) *DescribeDomainQPSListRequest
	GetInterval() *int64
	SetResourceGroupId(v string) *DescribeDomainQPSListRequest
	GetResourceGroupId() *string
	SetStartTime(v int64) *DescribeDomainQPSListRequest
	GetStartTime() *int64
}

type DescribeDomainQPSListRequest struct {
	// The domain name of the website. If you do not specify this parameter, the statistics on the QPS of all domain names are queried.
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
	// The interval for returning data. Unit: seconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1000
	Interval *int64 `json:"Interval,omitempty" xml:"Interval,omitempty"`
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

func (s DescribeDomainQPSListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainQPSListRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainQPSListRequest) GetDomain() *string {
	return s.Domain
}

func (s *DescribeDomainQPSListRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeDomainQPSListRequest) GetInterval() *int64 {
	return s.Interval
}

func (s *DescribeDomainQPSListRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeDomainQPSListRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeDomainQPSListRequest) SetDomain(v string) *DescribeDomainQPSListRequest {
	s.Domain = &v
	return s
}

func (s *DescribeDomainQPSListRequest) SetEndTime(v int64) *DescribeDomainQPSListRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDomainQPSListRequest) SetInterval(v int64) *DescribeDomainQPSListRequest {
	s.Interval = &v
	return s
}

func (s *DescribeDomainQPSListRequest) SetResourceGroupId(v string) *DescribeDomainQPSListRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDomainQPSListRequest) SetStartTime(v int64) *DescribeDomainQPSListRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDomainQPSListRequest) Validate() error {
	return dara.Validate(s)
}
