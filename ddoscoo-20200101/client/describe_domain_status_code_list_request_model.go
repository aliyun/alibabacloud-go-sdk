// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainStatusCodeListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomain(v string) *DescribeDomainStatusCodeListRequest
	GetDomain() *string
	SetEndTime(v int64) *DescribeDomainStatusCodeListRequest
	GetEndTime() *int64
	SetInterval(v int64) *DescribeDomainStatusCodeListRequest
	GetInterval() *int64
	SetQueryType(v string) *DescribeDomainStatusCodeListRequest
	GetQueryType() *string
	SetResourceGroupId(v string) *DescribeDomainStatusCodeListRequest
	GetResourceGroupId() *string
	SetStartTime(v int64) *DescribeDomainStatusCodeListRequest
	GetStartTime() *int64
}

type DescribeDomainStatusCodeListRequest struct {
	// The domain name of the website service.
	//
	// > You can invoke [DescribeDomains](https://help.aliyun.com/document_detail/91724.html) to query the domain names of all website services that are added to Anti-DDoS Pro or Anti-DDoS Premium.
	//
	// example:
	//
	// www.example.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The end of the time range to query. Specify a UNIX timestamp. Unit: seconds.
	//
	// > The value must be on the exact minute.
	//
	// example:
	//
	// 1583683200
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The time interval between consecutive data entries returned. Unit: seconds. This parameter specifies how frequently a query result is returned.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1000
	Interval *int64 `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// The source of the query data. Valid values:
	//
	// - **gf**: Anti-DDoS Pro or Anti-DDoS Premium response data.
	//
	// - **upstream**: origin server response data.
	//
	// This parameter is required.
	//
	// example:
	//
	// upstream
	QueryType *string `json:"QueryType,omitempty" xml:"QueryType,omitempty"`
	// The ID of the resource group to which the Anti-DDoS Pro or Anti-DDoS Premium instance belongs in Resource Management. By default, this parameter is empty, which indicates that the instance belongs to the default resource group.
	//
	// For more information about resource groups, refer to [Create a resource group](https://help.aliyun.com/document_detail/94485.html).
	//
	// example:
	//
	// rg-acfm2pz25js****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The beginning of the time range to query. Specify a UNIX timestamp. Unit: seconds.
	//
	// > The value must be on the exact minute.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1582992000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDomainStatusCodeListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainStatusCodeListRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainStatusCodeListRequest) GetDomain() *string {
	return s.Domain
}

func (s *DescribeDomainStatusCodeListRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeDomainStatusCodeListRequest) GetInterval() *int64 {
	return s.Interval
}

func (s *DescribeDomainStatusCodeListRequest) GetQueryType() *string {
	return s.QueryType
}

func (s *DescribeDomainStatusCodeListRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeDomainStatusCodeListRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeDomainStatusCodeListRequest) SetDomain(v string) *DescribeDomainStatusCodeListRequest {
	s.Domain = &v
	return s
}

func (s *DescribeDomainStatusCodeListRequest) SetEndTime(v int64) *DescribeDomainStatusCodeListRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDomainStatusCodeListRequest) SetInterval(v int64) *DescribeDomainStatusCodeListRequest {
	s.Interval = &v
	return s
}

func (s *DescribeDomainStatusCodeListRequest) SetQueryType(v string) *DescribeDomainStatusCodeListRequest {
	s.QueryType = &v
	return s
}

func (s *DescribeDomainStatusCodeListRequest) SetResourceGroupId(v string) *DescribeDomainStatusCodeListRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDomainStatusCodeListRequest) SetStartTime(v int64) *DescribeDomainStatusCodeListRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDomainStatusCodeListRequest) Validate() error {
	return dara.Validate(s)
}
