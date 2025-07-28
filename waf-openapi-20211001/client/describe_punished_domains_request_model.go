// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePunishedDomainsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomains(v []*string) *DescribePunishedDomainsRequest
	GetDomains() []*string
	SetInstanceId(v string) *DescribePunishedDomainsRequest
	GetInstanceId() *string
	SetPunishType(v string) *DescribePunishedDomainsRequest
	GetPunishType() *string
	SetRegionId(v string) *DescribePunishedDomainsRequest
	GetRegionId() *string
	SetResourceManagerResourceGroupId(v string) *DescribePunishedDomainsRequest
	GetResourceManagerResourceGroupId() *string
}

type DescribePunishedDomainsRequest struct {
	// The domain names that are added to WAF.
	Domains []*string `json:"Domains,omitempty" xml:"Domains,omitempty" type:"Repeated"`
	// The ID of the WAF instance.
	//
	// >  You can call the [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) operation to query the ID of the WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf_v3prepaid_public_cn-uqm****qa07
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The type of punishment. Valid values:
	//
	// 	- **beian*	- (default): the filing center.
	//
	// 	- **punishCenter**: the punishment center.
	//
	// example:
	//
	// beian
	PunishType *string `json:"PunishType,omitempty" xml:"PunishType,omitempty"`
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
	// rg-aekz7nc****aata
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
}

func (s DescribePunishedDomainsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePunishedDomainsRequest) GoString() string {
	return s.String()
}

func (s *DescribePunishedDomainsRequest) GetDomains() []*string {
	return s.Domains
}

func (s *DescribePunishedDomainsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribePunishedDomainsRequest) GetPunishType() *string {
	return s.PunishType
}

func (s *DescribePunishedDomainsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribePunishedDomainsRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *DescribePunishedDomainsRequest) SetDomains(v []*string) *DescribePunishedDomainsRequest {
	s.Domains = v
	return s
}

func (s *DescribePunishedDomainsRequest) SetInstanceId(v string) *DescribePunishedDomainsRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribePunishedDomainsRequest) SetPunishType(v string) *DescribePunishedDomainsRequest {
	s.PunishType = &v
	return s
}

func (s *DescribePunishedDomainsRequest) SetRegionId(v string) *DescribePunishedDomainsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribePunishedDomainsRequest) SetResourceManagerResourceGroupId(v string) *DescribePunishedDomainsRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribePunishedDomainsRequest) Validate() error {
	return dara.Validate(s)
}
