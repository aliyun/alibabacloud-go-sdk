// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeWebAreaBlockConfigsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomains(v []*string) *DescribeWebAreaBlockConfigsRequest
	GetDomains() []*string
	SetResourceGroupId(v string) *DescribeWebAreaBlockConfigsRequest
	GetResourceGroupId() *string
}

type DescribeWebAreaBlockConfigsRequest struct {
	// The domain names of the website services.
	//
	// > The domain names must have forwarding rules configured for website services. You can call [DescribeDomains](https://help.aliyun.com/document_detail/91724.html) to query all domain names.
	//
	// This parameter is required.
	//
	// example:
	//
	// www.aliyun.com
	Domains []*string `json:"Domains,omitempty" xml:"Domains,omitempty" type:"Repeated"`
	// The resource group ID of the Anti-DDoS Proxy instance in Resource Management. The default value is empty, which indicates that the instance belongs to the default resource group.
	//
	// example:
	//
	// default
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s DescribeWebAreaBlockConfigsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeWebAreaBlockConfigsRequest) GoString() string {
	return s.String()
}

func (s *DescribeWebAreaBlockConfigsRequest) GetDomains() []*string {
	return s.Domains
}

func (s *DescribeWebAreaBlockConfigsRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeWebAreaBlockConfigsRequest) SetDomains(v []*string) *DescribeWebAreaBlockConfigsRequest {
	s.Domains = v
	return s
}

func (s *DescribeWebAreaBlockConfigsRequest) SetResourceGroupId(v string) *DescribeWebAreaBlockConfigsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeWebAreaBlockConfigsRequest) Validate() error {
	return dara.Validate(s)
}
