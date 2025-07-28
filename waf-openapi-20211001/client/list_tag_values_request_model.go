// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTagValuesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListTagValuesRequest
	GetInstanceId() *string
	SetKey(v string) *ListTagValuesRequest
	GetKey() *string
	SetNextToken(v string) *ListTagValuesRequest
	GetNextToken() *string
	SetRegionId(v string) *ListTagValuesRequest
	GetRegionId() *string
	SetResourceManagerResourceGroupId(v string) *ListTagValuesRequest
	GetResourceManagerResourceGroupId() *string
	SetResourceType(v string) *ListTagValuesRequest
	GetResourceType() *string
}

type ListTagValuesRequest struct {
	// The ID of the WAF instance.
	//
	// >  Obtain the ID of the WAF instance by calling the [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf_v2_public_cn-lbj****x10g
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The tag key.
	//
	// This parameter is required.
	//
	// example:
	//
	// demoKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0*****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The region of the WAF instance. Valid values:
	//
	// 	- **cn-hangzhou**: Chinese mainland.
	//
	// 	- **ap-southeast-1**: Outside the Chinese mainland.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the Alibaba Cloud resource group.
	//
	// example:
	//
	// rg-aekzwwk****cv5i
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
	// The type of the resource. Set the value to ALIYUN::WAF::DEFENSERESOURCE.
	//
	// This parameter is required.
	//
	// example:
	//
	// ALIYUN::WAF::DEFENSERESOURCE
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s ListTagValuesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTagValuesRequest) GoString() string {
	return s.String()
}

func (s *ListTagValuesRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListTagValuesRequest) GetKey() *string {
	return s.Key
}

func (s *ListTagValuesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListTagValuesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListTagValuesRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *ListTagValuesRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *ListTagValuesRequest) SetInstanceId(v string) *ListTagValuesRequest {
	s.InstanceId = &v
	return s
}

func (s *ListTagValuesRequest) SetKey(v string) *ListTagValuesRequest {
	s.Key = &v
	return s
}

func (s *ListTagValuesRequest) SetNextToken(v string) *ListTagValuesRequest {
	s.NextToken = &v
	return s
}

func (s *ListTagValuesRequest) SetRegionId(v string) *ListTagValuesRequest {
	s.RegionId = &v
	return s
}

func (s *ListTagValuesRequest) SetResourceManagerResourceGroupId(v string) *ListTagValuesRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *ListTagValuesRequest) SetResourceType(v string) *ListTagValuesRequest {
	s.ResourceType = &v
	return s
}

func (s *ListTagValuesRequest) Validate() error {
	return dara.Validate(s)
}
