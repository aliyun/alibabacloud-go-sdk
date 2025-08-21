// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainStatusCodeCountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomain(v string) *DescribeDomainStatusCodeCountRequest
	GetDomain() *string
	SetEndTime(v int64) *DescribeDomainStatusCodeCountRequest
	GetEndTime() *int64
	SetResourceGroupId(v string) *DescribeDomainStatusCodeCountRequest
	GetResourceGroupId() *string
	SetStartTime(v int64) *DescribeDomainStatusCodeCountRequest
	GetStartTime() *int64
}

type DescribeDomainStatusCodeCountRequest struct {
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

func (s DescribeDomainStatusCodeCountRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainStatusCodeCountRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainStatusCodeCountRequest) GetDomain() *string {
	return s.Domain
}

func (s *DescribeDomainStatusCodeCountRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeDomainStatusCodeCountRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeDomainStatusCodeCountRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeDomainStatusCodeCountRequest) SetDomain(v string) *DescribeDomainStatusCodeCountRequest {
	s.Domain = &v
	return s
}

func (s *DescribeDomainStatusCodeCountRequest) SetEndTime(v int64) *DescribeDomainStatusCodeCountRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDomainStatusCodeCountRequest) SetResourceGroupId(v string) *DescribeDomainStatusCodeCountRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDomainStatusCodeCountRequest) SetStartTime(v int64) *DescribeDomainStatusCodeCountRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDomainStatusCodeCountRequest) Validate() error {
	return dara.Validate(s)
}
