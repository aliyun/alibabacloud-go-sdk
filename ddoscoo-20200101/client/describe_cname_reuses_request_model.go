// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCnameReusesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomains(v []*string) *DescribeCnameReusesRequest
	GetDomains() []*string
	SetResourceGroupId(v string) *DescribeCnameReusesRequest
	GetResourceGroupId() *string
}

type DescribeCnameReusesRequest struct {
	// The domain names of the websites. You can specify the domain names of up to 200 websites.
	//
	// > A forwarding rule must be configured for the domain name. You can call [DescribeDomains](https://help.aliyun.com/document_detail/91724.html) to query all domain names.
	//
	// This parameter is required.
	//
	// example:
	//
	// www.aliyun.com
	Domains []*string `json:"Domains,omitempty" xml:"Domains,omitempty" type:"Repeated"`
	// The ID of the resource group to which the Anti-DDoS Proxy instance belongs in Resource Management. This parameter is empty by default, which indicates that the instance belongs to the default resource group.
	//
	// example:
	//
	// default
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s DescribeCnameReusesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCnameReusesRequest) GoString() string {
	return s.String()
}

func (s *DescribeCnameReusesRequest) GetDomains() []*string {
	return s.Domains
}

func (s *DescribeCnameReusesRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeCnameReusesRequest) SetDomains(v []*string) *DescribeCnameReusesRequest {
	s.Domains = v
	return s
}

func (s *DescribeCnameReusesRequest) SetResourceGroupId(v string) *DescribeCnameReusesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeCnameReusesRequest) Validate() error {
	return dara.Validate(s)
}
