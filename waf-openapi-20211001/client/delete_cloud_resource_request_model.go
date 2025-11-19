// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCloudResourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCloudResourceId(v string) *DeleteCloudResourceRequest
	GetCloudResourceId() *string
	SetInstanceId(v string) *DeleteCloudResourceRequest
	GetInstanceId() *string
	SetPort(v int32) *DeleteCloudResourceRequest
	GetPort() *int32
	SetRegionId(v string) *DeleteCloudResourceRequest
	GetRegionId() *string
	SetResourceInstanceId(v string) *DeleteCloudResourceRequest
	GetResourceInstanceId() *string
	SetResourceManagerResourceGroupId(v string) *DeleteCloudResourceRequest
	GetResourceManagerResourceGroupId() *string
	SetResourceProduct(v string) *DeleteCloudResourceRequest
	GetResourceProduct() *string
}

type DeleteCloudResourceRequest struct {
	CloudResourceId *string `json:"CloudResourceId,omitempty" xml:"CloudResourceId,omitempty"`
	// The ID of the WAF instance.
	//
	// >  You can call the [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) operation to query the ID of the WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf_v2_public_cn-***
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Deprecated
	//
	// The port of the resource that is added to WAF.
	//
	// example:
	//
	// 443
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The region in which the WAF instance is deployed. Valid values:
	//
	// 	- **cn-hangzhou**: the Chinese mainland.
	//
	// 	- **ap-southeast-1**: outside the Chinese mainland.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Deprecated
	//
	// The ID of the instance.
	//
	// example:
	//
	// lb-bp1*****jqnnqk5uj2p
	ResourceInstanceId *string `json:"ResourceInstanceId,omitempty" xml:"ResourceInstanceId,omitempty"`
	// The ID of the Alibaba Cloud resource group.
	//
	// example:
	//
	// rg-acfm***q
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
	// Deprecated
	//
	// The cloud service. Valid values:
	//
	// 	- **clb4**: Layer 4 CLB.
	//
	// 	- **clb7**: Layer 7 CLB.
	//
	// 	- **ecs**: ECS.
	//
	// example:
	//
	// clb7
	ResourceProduct *string `json:"ResourceProduct,omitempty" xml:"ResourceProduct,omitempty"`
}

func (s DeleteCloudResourceRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteCloudResourceRequest) GoString() string {
	return s.String()
}

func (s *DeleteCloudResourceRequest) GetCloudResourceId() *string {
	return s.CloudResourceId
}

func (s *DeleteCloudResourceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteCloudResourceRequest) GetPort() *int32 {
	return s.Port
}

func (s *DeleteCloudResourceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteCloudResourceRequest) GetResourceInstanceId() *string {
	return s.ResourceInstanceId
}

func (s *DeleteCloudResourceRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *DeleteCloudResourceRequest) GetResourceProduct() *string {
	return s.ResourceProduct
}

func (s *DeleteCloudResourceRequest) SetCloudResourceId(v string) *DeleteCloudResourceRequest {
	s.CloudResourceId = &v
	return s
}

func (s *DeleteCloudResourceRequest) SetInstanceId(v string) *DeleteCloudResourceRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteCloudResourceRequest) SetPort(v int32) *DeleteCloudResourceRequest {
	s.Port = &v
	return s
}

func (s *DeleteCloudResourceRequest) SetRegionId(v string) *DeleteCloudResourceRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteCloudResourceRequest) SetResourceInstanceId(v string) *DeleteCloudResourceRequest {
	s.ResourceInstanceId = &v
	return s
}

func (s *DeleteCloudResourceRequest) SetResourceManagerResourceGroupId(v string) *DeleteCloudResourceRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DeleteCloudResourceRequest) SetResourceProduct(v string) *DeleteCloudResourceRequest {
	s.ResourceProduct = &v
	return s
}

func (s *DeleteCloudResourceRequest) Validate() error {
	return dara.Validate(s)
}
