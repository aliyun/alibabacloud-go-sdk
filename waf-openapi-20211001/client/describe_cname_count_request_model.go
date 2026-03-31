// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCnameCountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribeCnameCountRequest
	GetInstanceId() *string
	SetRegionId(v string) *DescribeCnameCountRequest
	GetRegionId() *string
	SetResourceManagerResourceGroupId(v string) *DescribeCnameCountRequest
	GetResourceManagerResourceGroupId() *string
}

type DescribeCnameCountRequest struct {
	// The ID of the WAF instance.
	//
	// >  You can call the [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) operation to query the ID of the WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf_v3prepaid_public_cn-***
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region in which the WAF instance is deployed. Valid values:
	//
	// 	- **cn-hangzhou**: the Chinese mainland.
	//
	// 	- **ap-southeast-1**: outside the Chinese mainland.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// 阿里云资源组ID。
	//
	// example:
	//
	// rg-acfm***q
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
}

func (s DescribeCnameCountRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCnameCountRequest) GoString() string {
	return s.String()
}

func (s *DescribeCnameCountRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeCnameCountRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeCnameCountRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *DescribeCnameCountRequest) SetInstanceId(v string) *DescribeCnameCountRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeCnameCountRequest) SetRegionId(v string) *DescribeCnameCountRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeCnameCountRequest) SetResourceManagerResourceGroupId(v string) *DescribeCnameCountRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribeCnameCountRequest) Validate() error {
	return dara.Validate(s)
}
