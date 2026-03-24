// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApisecMatchedHostsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *DescribeApisecMatchedHostsRequest
	GetClusterId() *string
	SetInstanceId(v string) *DescribeApisecMatchedHostsRequest
	GetInstanceId() *string
	SetMatchedHost(v string) *DescribeApisecMatchedHostsRequest
	GetMatchedHost() *string
	SetPageNumber(v string) *DescribeApisecMatchedHostsRequest
	GetPageNumber() *string
	SetPageSize(v string) *DescribeApisecMatchedHostsRequest
	GetPageSize() *string
	SetRegionId(v string) *DescribeApisecMatchedHostsRequest
	GetRegionId() *string
	SetResourceManagerResourceGroupId(v string) *DescribeApisecMatchedHostsRequest
	GetResourceManagerResourceGroupId() *string
	SetType(v string) *DescribeApisecMatchedHostsRequest
	GetType() *string
}

type DescribeApisecMatchedHostsRequest struct {
	// The ID of the hybrid cloud cluster.
	//
	// > This parameter is used only in hybrid cloud scenarios. Call the [DescribeHybridCloudClusters](https://help.aliyun.com/document_detail/2849376.html) operation to obtain information about hybrid cloud clusters.
	//
	// example:
	//
	// 433
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The ID of the WAF instance.
	//
	// > Call the [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) operation to view the ID of the current WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf_v2_public_cn-x0r37plpl0g
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The domain name or IP address.
	//
	// example:
	//
	// bc.aliyun.com
	MatchedHost *string `json:"MatchedHost,omitempty" xml:"MatchedHost,omitempty"`
	// The number of the page to return for a paged query. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page for a paged query. Default value: **10**.
	//
	// example:
	//
	// 8
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region of the WAF instance. Valid values:
	//
	// - **cn-hangzhou**: the Chinese mainland.
	//
	// - **ap-southeast-1**: outside the Chinese mainland.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the Alibaba Cloud resource group.
	//
	// example:
	//
	// rg-aekz5qqo7jthcsa
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
	// The detection type. Valid values:
	//
	// - **api**: the list of domain names for API assets.
	//
	// - **abnormal**: the list of domain names that have threats.
	//
	// - **event**: the list of domain names that have management events.
	//
	// example:
	//
	// event
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeApisecMatchedHostsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeApisecMatchedHostsRequest) GoString() string {
	return s.String()
}

func (s *DescribeApisecMatchedHostsRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeApisecMatchedHostsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeApisecMatchedHostsRequest) GetMatchedHost() *string {
	return s.MatchedHost
}

func (s *DescribeApisecMatchedHostsRequest) GetPageNumber() *string {
	return s.PageNumber
}

func (s *DescribeApisecMatchedHostsRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *DescribeApisecMatchedHostsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeApisecMatchedHostsRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *DescribeApisecMatchedHostsRequest) GetType() *string {
	return s.Type
}

func (s *DescribeApisecMatchedHostsRequest) SetClusterId(v string) *DescribeApisecMatchedHostsRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeApisecMatchedHostsRequest) SetInstanceId(v string) *DescribeApisecMatchedHostsRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeApisecMatchedHostsRequest) SetMatchedHost(v string) *DescribeApisecMatchedHostsRequest {
	s.MatchedHost = &v
	return s
}

func (s *DescribeApisecMatchedHostsRequest) SetPageNumber(v string) *DescribeApisecMatchedHostsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeApisecMatchedHostsRequest) SetPageSize(v string) *DescribeApisecMatchedHostsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeApisecMatchedHostsRequest) SetRegionId(v string) *DescribeApisecMatchedHostsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeApisecMatchedHostsRequest) SetResourceManagerResourceGroupId(v string) *DescribeApisecMatchedHostsRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribeApisecMatchedHostsRequest) SetType(v string) *DescribeApisecMatchedHostsRequest {
	s.Type = &v
	return s
}

func (s *DescribeApisecMatchedHostsRequest) Validate() error {
	return dara.Validate(s)
}
