// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeResourcePortRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribeResourcePortRequest
	GetInstanceId() *string
	SetRegionId(v string) *DescribeResourcePortRequest
	GetRegionId() *string
	SetResourceInstanceId(v string) *DescribeResourcePortRequest
	GetResourceInstanceId() *string
	SetResourceManagerResourceGroupId(v string) *DescribeResourcePortRequest
	GetResourceManagerResourceGroupId() *string
}

type DescribeResourcePortRequest struct {
	// The ID of the WAF instance.
	//
	// > You can call [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) to query the ID of the current WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf_cdnsdf3****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of WAF instance. Valid values:
	//
	// - **cn-hangzhou**: The Chinese mainland.
	//
	// - **ap-southeast-1**: Outside the Chinese mainland.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the cloud service instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// 8vb40vj87znu3ai7l5lv4-80****
	ResourceInstanceId *string `json:"ResourceInstanceId,omitempty" xml:"ResourceInstanceId,omitempty"`
	// The ID of the Alibaba Cloud resource group.
	//
	// example:
	//
	// rg-acfm***q
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
}

func (s DescribeResourcePortRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeResourcePortRequest) GoString() string {
	return s.String()
}

func (s *DescribeResourcePortRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeResourcePortRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeResourcePortRequest) GetResourceInstanceId() *string {
	return s.ResourceInstanceId
}

func (s *DescribeResourcePortRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *DescribeResourcePortRequest) SetInstanceId(v string) *DescribeResourcePortRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeResourcePortRequest) SetRegionId(v string) *DescribeResourcePortRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeResourcePortRequest) SetResourceInstanceId(v string) *DescribeResourcePortRequest {
	s.ResourceInstanceId = &v
	return s
}

func (s *DescribeResourcePortRequest) SetResourceManagerResourceGroupId(v string) *DescribeResourcePortRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribeResourcePortRequest) Validate() error {
	return dara.Validate(s)
}
