// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUserAssetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *DescribeUserAssetRequest
	GetClusterId() *string
	SetDataType(v string) *DescribeUserAssetRequest
	GetDataType() *string
	SetDays(v string) *DescribeUserAssetRequest
	GetDays() *string
	SetInstanceId(v string) *DescribeUserAssetRequest
	GetInstanceId() *string
	SetRegionId(v string) *DescribeUserAssetRequest
	GetRegionId() *string
	SetResourceManagerResourceGroupId(v string) *DescribeUserAssetRequest
	GetResourceManagerResourceGroupId() *string
}

type DescribeUserAssetRequest struct {
	// The ID of the hybrid cloud cluster.
	//
	// >For hybrid cloud scenarios only, you can call the [DescribeHybridCloudClusters](https://help.aliyun.com/document_detail/2849376.html) operation to query the hybrid cloud clusters.
	//
	// example:
	//
	// 428
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The type of the statistics. Valid values:
	//
	// 	- **asset_num**: total number of APIs
	//
	// 	- **asset_active**: number of active APIs
	//
	// 	- **asset_newborn**: number of new APIs
	//
	// 	- **asset_offline**: number of deactivated APIs
	//
	// 	- **asset_bot**: number of APIs that are called by bots
	//
	// 	- **asset_cross_border**: number of APIs that are called for cross-border data transfer
	//
	// 	- **sensitive_api**: number of response-sensitive APIs
	//
	// 	- **sensitive_domain**: number of response-sensitive domain names
	//
	// This parameter is required.
	//
	// example:
	//
	// asset_num
	DataType *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	// Deprecated
	//
	// The time at which the API was called. The value is a UNIX timestamp displayed in UTC. Unit: seconds.
	//
	// 	Notice: The parameter has been deprecated.
	//
	// example:
	//
	// 1723435200
	Days *string `json:"Days,omitempty" xml:"Days,omitempty"`
	// The ID of the Web Application Firewall (WAF) instance.
	//
	// >  You can call the [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) operation to query the ID of the WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf-cn-uax37ijm***
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
}

func (s DescribeUserAssetRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserAssetRequest) GoString() string {
	return s.String()
}

func (s *DescribeUserAssetRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeUserAssetRequest) GetDataType() *string {
	return s.DataType
}

func (s *DescribeUserAssetRequest) GetDays() *string {
	return s.Days
}

func (s *DescribeUserAssetRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeUserAssetRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeUserAssetRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *DescribeUserAssetRequest) SetClusterId(v string) *DescribeUserAssetRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeUserAssetRequest) SetDataType(v string) *DescribeUserAssetRequest {
	s.DataType = &v
	return s
}

func (s *DescribeUserAssetRequest) SetDays(v string) *DescribeUserAssetRequest {
	s.Days = &v
	return s
}

func (s *DescribeUserAssetRequest) SetInstanceId(v string) *DescribeUserAssetRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeUserAssetRequest) SetRegionId(v string) *DescribeUserAssetRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeUserAssetRequest) SetResourceManagerResourceGroupId(v string) *DescribeUserAssetRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribeUserAssetRequest) Validate() error {
	return dara.Validate(s)
}
