// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainOverviewRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomain(v string) *DescribeDomainOverviewRequest
	GetDomain() *string
	SetEndTime(v int64) *DescribeDomainOverviewRequest
	GetEndTime() *int64
	SetResourceGroupId(v string) *DescribeDomainOverviewRequest
	GetResourceGroupId() *string
	SetStartTime(v int64) *DescribeDomainOverviewRequest
	GetStartTime() *int64
}

type DescribeDomainOverviewRequest struct {
	// The domain name of the website that you want to query. If you leave this parameter unspecified, the statistics on all domain names are queried.
	//
	// > The domain name must be added to Anti-DDoS Pro or Anti-DDoS Premium. You can call the [DescribeDomains](https://help.aliyun.com/document_detail/91724.html) operation to query all the domain names that are added to Anti-DDoS Pro or Anti-DDoS Premium.
	//
	// example:
	//
	// example.aliyundoc.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The end of the time range to query. The value is a UNIX timestamp. Unit: seconds. If you leave this parameter unspecified, the current system time is used as the end time.
	//
	// > This UNIX timestamp must indicate a point in time that is accurate to the minute.
	//
	// example:
	//
	// 1623427200
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The ID of the resource group to which the instance belongs in Resource Management. This parameter is empty by default, which indicates that the instance belongs to the default resource group.
	//
	// For more information about resource groups, see [Create a resource group](https://help.aliyun.com/document_detail/94485.html).
	//
	// example:
	//
	// rg-acfm2pz25js****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The beginning of the time range to query. The value is a UNIX timestamp. Unit: seconds.
	//
	// > This UNIX timestamp must indicate a point in time that is accurate to the minute.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1619798400
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDomainOverviewRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainOverviewRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainOverviewRequest) GetDomain() *string {
	return s.Domain
}

func (s *DescribeDomainOverviewRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeDomainOverviewRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeDomainOverviewRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeDomainOverviewRequest) SetDomain(v string) *DescribeDomainOverviewRequest {
	s.Domain = &v
	return s
}

func (s *DescribeDomainOverviewRequest) SetEndTime(v int64) *DescribeDomainOverviewRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDomainOverviewRequest) SetResourceGroupId(v string) *DescribeDomainOverviewRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDomainOverviewRequest) SetStartTime(v int64) *DescribeDomainOverviewRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDomainOverviewRequest) Validate() error {
	return dara.Validate(s)
}
