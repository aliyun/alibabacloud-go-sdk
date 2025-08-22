// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnWafDomainRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeDcdnWafDomainRequest
	GetDomainName() *string
	SetRegionId(v string) *DescribeDcdnWafDomainRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeDcdnWafDomainRequest
	GetResourceGroupId() *string
}

type DescribeDcdnWafDomainRequest struct {
	// The accelerated domain name. If you do not specify an accelerated domain name, all accelerated domain names are queried.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The region where WAF is enabled.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group. If you leave this parameter empty, the default resource group is used.
	//
	// example:
	//
	// /
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s DescribeDcdnWafDomainRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnWafDomainRequest) GoString() string {
	return s.String()
}

func (s *DescribeDcdnWafDomainRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDcdnWafDomainRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDcdnWafDomainRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeDcdnWafDomainRequest) SetDomainName(v string) *DescribeDcdnWafDomainRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDcdnWafDomainRequest) SetRegionId(v string) *DescribeDcdnWafDomainRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDcdnWafDomainRequest) SetResourceGroupId(v string) *DescribeDcdnWafDomainRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDcdnWafDomainRequest) Validate() error {
	return dara.Validate(s)
}
