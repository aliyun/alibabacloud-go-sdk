// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainCcProtectSwitchRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomains(v []*string) *DescribeDomainCcProtectSwitchRequest
	GetDomains() []*string
	SetResourceGroupId(v string) *DescribeDomainCcProtectSwitchRequest
	GetResourceGroupId() *string
}

type DescribeDomainCcProtectSwitchRequest struct {
	// The domain names of the websites. You can specify up to five domain names.
	//
	// > A website forwarding rule must be configured for the domain name. Call [DescribeDomains](https://help.aliyun.com/document_detail/91724.html) to query all domain names.
	//
	// This parameter is required.
	Domains []*string `json:"Domains,omitempty" xml:"Domains,omitempty" type:"Repeated"`
	// The ID of the resource group to which the Anti-DDoS Pro instance belongs in Resource Management. If you do not specify this parameter, the instance belongs to the default resource group.
	//
	// example:
	//
	// rg-acfm2pz25js****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s DescribeDomainCcProtectSwitchRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainCcProtectSwitchRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainCcProtectSwitchRequest) GetDomains() []*string {
	return s.Domains
}

func (s *DescribeDomainCcProtectSwitchRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeDomainCcProtectSwitchRequest) SetDomains(v []*string) *DescribeDomainCcProtectSwitchRequest {
	s.Domains = v
	return s
}

func (s *DescribeDomainCcProtectSwitchRequest) SetResourceGroupId(v string) *DescribeDomainCcProtectSwitchRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDomainCcProtectSwitchRequest) Validate() error {
	return dara.Validate(s)
}
