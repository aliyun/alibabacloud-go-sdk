// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeWebCacheConfigsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomains(v []*string) *DescribeWebCacheConfigsRequest
	GetDomains() []*string
	SetResourceGroupId(v string) *DescribeWebCacheConfigsRequest
	GetResourceGroupId() *string
}

type DescribeWebCacheConfigsRequest struct {
	// An array consisting of domain names for which you want to query the Static Page Caching configurations.
	//
	// This parameter is required.
	//
	// example:
	//
	// www.aliyun.com
	Domains []*string `json:"Domains,omitempty" xml:"Domains,omitempty" type:"Repeated"`
	// The ID of the resource group to which the instance belongs in Resource Management.
	//
	// If you do not configure this parameter, the instance belongs to the default resource group.
	//
	// example:
	//
	// rg-acfm2pz25js****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s DescribeWebCacheConfigsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeWebCacheConfigsRequest) GoString() string {
	return s.String()
}

func (s *DescribeWebCacheConfigsRequest) GetDomains() []*string {
	return s.Domains
}

func (s *DescribeWebCacheConfigsRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeWebCacheConfigsRequest) SetDomains(v []*string) *DescribeWebCacheConfigsRequest {
	s.Domains = v
	return s
}

func (s *DescribeWebCacheConfigsRequest) SetResourceGroupId(v string) *DescribeWebCacheConfigsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeWebCacheConfigsRequest) Validate() error {
	return dara.Validate(s)
}
