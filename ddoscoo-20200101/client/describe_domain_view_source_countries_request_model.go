// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainViewSourceCountriesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomain(v string) *DescribeDomainViewSourceCountriesRequest
	GetDomain() *string
	SetEndTime(v int64) *DescribeDomainViewSourceCountriesRequest
	GetEndTime() *int64
	SetResourceGroupId(v string) *DescribeDomainViewSourceCountriesRequest
	GetResourceGroupId() *string
	SetStartTime(v int64) *DescribeDomainViewSourceCountriesRequest
	GetStartTime() *int64
}

type DescribeDomainViewSourceCountriesRequest struct {
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
}

func (s DescribeDomainViewSourceCountriesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainViewSourceCountriesRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainViewSourceCountriesRequest) GetDomain() *string {
	return s.Domain
}

func (s *DescribeDomainViewSourceCountriesRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeDomainViewSourceCountriesRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeDomainViewSourceCountriesRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeDomainViewSourceCountriesRequest) SetDomain(v string) *DescribeDomainViewSourceCountriesRequest {
	s.Domain = &v
	return s
}

func (s *DescribeDomainViewSourceCountriesRequest) SetEndTime(v int64) *DescribeDomainViewSourceCountriesRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDomainViewSourceCountriesRequest) SetResourceGroupId(v string) *DescribeDomainViewSourceCountriesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDomainViewSourceCountriesRequest) SetStartTime(v int64) *DescribeDomainViewSourceCountriesRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDomainViewSourceCountriesRequest) Validate() error {
	return dara.Validate(s)
}
