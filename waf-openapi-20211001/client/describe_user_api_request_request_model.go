// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUserApiRequestRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiFormat(v string) *DescribeUserApiRequestRequest
	GetApiFormat() *string
	SetApiId(v string) *DescribeUserApiRequestRequest
	GetApiId() *string
	SetClusterId(v string) *DescribeUserApiRequestRequest
	GetClusterId() *string
	SetDomain(v string) *DescribeUserApiRequestRequest
	GetDomain() *string
	SetInstanceId(v string) *DescribeUserApiRequestRequest
	GetInstanceId() *string
	SetRegionId(v string) *DescribeUserApiRequestRequest
	GetRegionId() *string
	SetResourceManagerResourceGroupId(v string) *DescribeUserApiRequestRequest
	GetResourceManagerResourceGroupId() *string
	SetType(v string) *DescribeUserApiRequestRequest
	GetType() *string
}

type DescribeUserApiRequestRequest struct {
	// The API.
	//
	// 	Notice: This parameter is deprecated, please use ApiId to query.
	//
	// example:
	//
	// /api/v1/know
	ApiFormat *string `json:"ApiFormat,omitempty" xml:"ApiFormat,omitempty"`
	// The ID of the API.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3799f0695c0d687f3295d132fe49bc14
	ApiId *string `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	// The ID of the hybrid cloud cluster.
	//
	// >For hybrid cloud scenarios only, you can call the [DescribeHybridCloudClusters](https://help.aliyun.com/document_detail/2849376.html) operation to query the hybrid cloud clusters.
	//
	// example:
	//
	// 428
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The domain name or IP address of the API.
	//
	// 	Notice: This parameter is deprecated, please use ApiId to query.
	//
	// example:
	//
	// c.***.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The ID of the Web Application Firewall (WAF) instance.
	//
	// >  You can call the [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) operation to query the ID of the WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf_v3prepaid_public_cn-zxu38***
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region in which the WAF instance is deployed. Valid values:
	//
	// 	- **cn-hangzhou**: Chinese mainland
	//
	// 	- **ap-southeast-1**: outside the Chinese mainland
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
	// The type of the statistics. Valid values:
	//
	// 	- **api_ip**: total traffic
	//
	// 	- **api_cross_border_ip**: cross-border traffic
	//
	// 	- **api_bot_ip**: bot traffic
	//
	// 	- **remote_region**: geographical location
	//
	// 	- **client_id**: client type
	//
	// 	- **http_referer**: Referer
	//
	// 	- **api_cnt**: total number of calls
	//
	// 	- **bot_cnt**: number of bot calls
	//
	// 	- **cross_border_cnt**: number of cross-border requests
	//
	// 	- **api_freq**: call frequency
	//
	// example:
	//
	// api_ip
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeUserApiRequestRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserApiRequestRequest) GoString() string {
	return s.String()
}

func (s *DescribeUserApiRequestRequest) GetApiFormat() *string {
	return s.ApiFormat
}

func (s *DescribeUserApiRequestRequest) GetApiId() *string {
	return s.ApiId
}

func (s *DescribeUserApiRequestRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeUserApiRequestRequest) GetDomain() *string {
	return s.Domain
}

func (s *DescribeUserApiRequestRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeUserApiRequestRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeUserApiRequestRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *DescribeUserApiRequestRequest) GetType() *string {
	return s.Type
}

func (s *DescribeUserApiRequestRequest) SetApiFormat(v string) *DescribeUserApiRequestRequest {
	s.ApiFormat = &v
	return s
}

func (s *DescribeUserApiRequestRequest) SetApiId(v string) *DescribeUserApiRequestRequest {
	s.ApiId = &v
	return s
}

func (s *DescribeUserApiRequestRequest) SetClusterId(v string) *DescribeUserApiRequestRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeUserApiRequestRequest) SetDomain(v string) *DescribeUserApiRequestRequest {
	s.Domain = &v
	return s
}

func (s *DescribeUserApiRequestRequest) SetInstanceId(v string) *DescribeUserApiRequestRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeUserApiRequestRequest) SetRegionId(v string) *DescribeUserApiRequestRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeUserApiRequestRequest) SetResourceManagerResourceGroupId(v string) *DescribeUserApiRequestRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribeUserApiRequestRequest) SetType(v string) *DescribeUserApiRequestRequest {
	s.Type = &v
	return s
}

func (s *DescribeUserApiRequestRequest) Validate() error {
	return dara.Validate(s)
}
