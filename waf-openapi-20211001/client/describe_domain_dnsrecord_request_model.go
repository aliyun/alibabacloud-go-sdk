// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainDNSRecordRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomain(v string) *DescribeDomainDNSRecordRequest
	GetDomain() *string
	SetInstanceId(v string) *DescribeDomainDNSRecordRequest
	GetInstanceId() *string
	SetRegionId(v string) *DescribeDomainDNSRecordRequest
	GetRegionId() *string
	SetResourceManagerResourceGroupId(v string) *DescribeDomainDNSRecordRequest
	GetResourceManagerResourceGroupId() *string
}

type DescribeDomainDNSRecordRequest struct {
	// The domain name whose DNS settings you want to check.
	//
	// This parameter is required.
	//
	// example:
	//
	// xxx.aliyundoc.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The ID of the Web Application Firewall (WAF) instance.
	//
	// >  You can call the [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) operation to query the ID of the WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf_v3prepaid_public_cn-nwy****is02
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region in which the WAF instance is deployed. Valid values:
	//
	// 	- **cn-hangzhou**: Chinese mainland.
	//
	// 	- **ap-southeast-1**: outside the Chinese mainland.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the Alibaba Cloud resource group.
	//
	// example:
	//
	// rg-acfm***q
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
}

func (s DescribeDomainDNSRecordRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainDNSRecordRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainDNSRecordRequest) GetDomain() *string {
	return s.Domain
}

func (s *DescribeDomainDNSRecordRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeDomainDNSRecordRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDomainDNSRecordRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *DescribeDomainDNSRecordRequest) SetDomain(v string) *DescribeDomainDNSRecordRequest {
	s.Domain = &v
	return s
}

func (s *DescribeDomainDNSRecordRequest) SetInstanceId(v string) *DescribeDomainDNSRecordRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeDomainDNSRecordRequest) SetRegionId(v string) *DescribeDomainDNSRecordRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDomainDNSRecordRequest) SetResourceManagerResourceGroupId(v string) *DescribeDomainDNSRecordRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribeDomainDNSRecordRequest) Validate() error {
	return dara.Validate(s)
}
