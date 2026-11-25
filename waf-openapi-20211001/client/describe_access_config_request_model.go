// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAccessConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigKey(v string) *DescribeAccessConfigRequest
	GetConfigKey() *string
	SetInstanceId(v string) *DescribeAccessConfigRequest
	GetInstanceId() *string
	SetRegionId(v string) *DescribeAccessConfigRequest
	GetRegionId() *string
	SetResourceManagerResourceGroupId(v string) *DescribeAccessConfigRequest
	GetResourceManagerResourceGroupId() *string
}

type DescribeAccessConfigRequest struct {
	// The name of the configuration item to query. Valid values:
	//
	// - **auto_access**: Specifies whether to enable the automatic access feature.
	//
	// This parameter is required.
	//
	// example:
	//
	// auto_access
	ConfigKey *string `json:"ConfigKey,omitempty" xml:"ConfigKey,omitempty"`
	// Instance ID of the WAF instance.
	//
	// > You can call the [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) operation to view instance ID of the current WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf_elasticity-cn-0xldbqtm005
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
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
	// rg-acfm***q
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
}

func (s DescribeAccessConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAccessConfigRequest) GoString() string {
	return s.String()
}

func (s *DescribeAccessConfigRequest) GetConfigKey() *string {
	return s.ConfigKey
}

func (s *DescribeAccessConfigRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeAccessConfigRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeAccessConfigRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *DescribeAccessConfigRequest) SetConfigKey(v string) *DescribeAccessConfigRequest {
	s.ConfigKey = &v
	return s
}

func (s *DescribeAccessConfigRequest) SetInstanceId(v string) *DescribeAccessConfigRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeAccessConfigRequest) SetRegionId(v string) *DescribeAccessConfigRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeAccessConfigRequest) SetResourceManagerResourceGroupId(v string) *DescribeAccessConfigRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribeAccessConfigRequest) Validate() error {
	return dara.Validate(s)
}
