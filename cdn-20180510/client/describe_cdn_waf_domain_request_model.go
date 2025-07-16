// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCdnWafDomainRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeCdnWafDomainRequest
	GetDomainName() *string
	SetRegionId(v string) *DescribeCdnWafDomainRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeCdnWafDomainRequest
	GetResourceGroupId() *string
}

type DescribeCdnWafDomainRequest struct {
	// The domain name that you want to query.
	//
	// You can specify only one domain name in each request. You have three options to configure this parameter:
	//
	// 	- Specify an exact domain name. For example, if you set this parameter to example.com, configuration information of example.com is queried.
	//
	// 	- Specify a keyword. For example, if you set this parameter to example, configuration information about all domain names that contain example is queried.
	//
	// 	- Leave this parameter empty. If this parameter is left empty, all accelerated domain names for which WAF is configured are queried.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The region where WAF is enabled. Valid values:
	//
	// 	- **cn-hangzhou**: inside the Chinese mainland
	//
	// 	- **ap-southeast-1**: outside the Chinese mainland
	//
	// > ap-southeast-1 includes Hong Kong (China), Macao (China), Taiwan (China), and other countries and regions.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// 1234
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s DescribeCdnWafDomainRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnWafDomainRequest) GoString() string {
	return s.String()
}

func (s *DescribeCdnWafDomainRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeCdnWafDomainRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeCdnWafDomainRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeCdnWafDomainRequest) SetDomainName(v string) *DescribeCdnWafDomainRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeCdnWafDomainRequest) SetRegionId(v string) *DescribeCdnWafDomainRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeCdnWafDomainRequest) SetResourceGroupId(v string) *DescribeCdnWafDomainRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeCdnWafDomainRequest) Validate() error {
	return dara.Validate(s)
}
