// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeL7RsPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomain(v string) *DescribeL7RsPolicyRequest
	GetDomain() *string
	SetRealServers(v []*string) *DescribeL7RsPolicyRequest
	GetRealServers() []*string
	SetResourceGroupId(v string) *DescribeL7RsPolicyRequest
	GetResourceGroupId() *string
}

type DescribeL7RsPolicyRequest struct {
	// The domain name of the website to query.
	//
	// > A forwarding rule must be configured for the domain name. You can call the [DescribeDomains](https://help.aliyun.com/document_detail/91724.html) operation to query the domain names for which forwarding rules are configured.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// An array that consists of N addresses of origin servers to query. The maximum value of N is 200. You can specify up to 200 addresses.
	//
	// example:
	//
	// 1.***.***.1
	RealServers []*string `json:"RealServers,omitempty" xml:"RealServers,omitempty" type:"Repeated"`
	// The ID of the resource group to which the instance belongs in Resource Management. This parameter is empty by default, which indicates that the instance belongs to the default resource group.
	//
	// For more information about resource groups, see [Create a resource group](https://help.aliyun.com/document_detail/94485.html).
	//
	// example:
	//
	// rg-acfm2pz25js****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s DescribeL7RsPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeL7RsPolicyRequest) GoString() string {
	return s.String()
}

func (s *DescribeL7RsPolicyRequest) GetDomain() *string {
	return s.Domain
}

func (s *DescribeL7RsPolicyRequest) GetRealServers() []*string {
	return s.RealServers
}

func (s *DescribeL7RsPolicyRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeL7RsPolicyRequest) SetDomain(v string) *DescribeL7RsPolicyRequest {
	s.Domain = &v
	return s
}

func (s *DescribeL7RsPolicyRequest) SetRealServers(v []*string) *DescribeL7RsPolicyRequest {
	s.RealServers = v
	return s
}

func (s *DescribeL7RsPolicyRequest) SetResourceGroupId(v string) *DescribeL7RsPolicyRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeL7RsPolicyRequest) Validate() error {
	return dara.Validate(s)
}
