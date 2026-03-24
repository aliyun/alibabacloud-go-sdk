// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyApisecApiResourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiId(v string) *ModifyApisecApiResourceRequest
	GetApiId() *string
	SetClusterId(v string) *ModifyApisecApiResourceRequest
	GetClusterId() *string
	SetFollow(v int64) *ModifyApisecApiResourceRequest
	GetFollow() *int64
	SetInstanceId(v string) *ModifyApisecApiResourceRequest
	GetInstanceId() *string
	SetNote(v string) *ModifyApisecApiResourceRequest
	GetNote() *string
	SetRegionId(v string) *ModifyApisecApiResourceRequest
	GetRegionId() *string
	SetResourceManagerResourceGroupId(v string) *ModifyApisecApiResourceRequest
	GetResourceManagerResourceGroupId() *string
}

type ModifyApisecApiResourceRequest struct {
	// The ID of the API asset that you want to modify.
	//
	// This parameter is required.
	//
	// example:
	//
	// c68995b89069595c5c0399676f3ca64f
	ApiId *string `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	// The ID of the hybrid cloud cluster.
	//
	// > This parameter is required only for hybrid cloud scenarios. You can call the [DescribeHybridCloudClusters](https://help.aliyun.com/document_detail/2849376.html) operation to query the ID of the hybrid cloud cluster.
	//
	// example:
	//
	// 428
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// Indicates whether to follow the API asset. Valid values:
	//
	// - **1**: follows the API asset.
	//
	// - **0*	- (default): does not follow the API asset.
	//
	// example:
	//
	// 0
	Follow *int64 `json:"Follow,omitempty" xml:"Follow,omitempty"`
	// The ID of the WAF instance.
	//
	// > You can call the [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) operation to query the ID of the WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf_elasticity-cn-0xldbqt****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The remarks of the API asset. You can use this parameter to add a custom annotation to the API asset for easier identification.
	//
	// example:
	//
	// know
	Note *string `json:"Note,omitempty" xml:"Note,omitempty"`
	// The region where the WAF instance resides. Valid values:
	//
	// - **cn-hangzhou**: the Chinese mainland.
	//
	// - **ap-southeast-1**: outside the Chinese mainland.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the Alibaba Cloud resource group to which the WAF instance belongs.
	//
	// example:
	//
	// rg-acfm***q
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
}

func (s ModifyApisecApiResourceRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyApisecApiResourceRequest) GoString() string {
	return s.String()
}

func (s *ModifyApisecApiResourceRequest) GetApiId() *string {
	return s.ApiId
}

func (s *ModifyApisecApiResourceRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *ModifyApisecApiResourceRequest) GetFollow() *int64 {
	return s.Follow
}

func (s *ModifyApisecApiResourceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyApisecApiResourceRequest) GetNote() *string {
	return s.Note
}

func (s *ModifyApisecApiResourceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyApisecApiResourceRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *ModifyApisecApiResourceRequest) SetApiId(v string) *ModifyApisecApiResourceRequest {
	s.ApiId = &v
	return s
}

func (s *ModifyApisecApiResourceRequest) SetClusterId(v string) *ModifyApisecApiResourceRequest {
	s.ClusterId = &v
	return s
}

func (s *ModifyApisecApiResourceRequest) SetFollow(v int64) *ModifyApisecApiResourceRequest {
	s.Follow = &v
	return s
}

func (s *ModifyApisecApiResourceRequest) SetInstanceId(v string) *ModifyApisecApiResourceRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyApisecApiResourceRequest) SetNote(v string) *ModifyApisecApiResourceRequest {
	s.Note = &v
	return s
}

func (s *ModifyApisecApiResourceRequest) SetRegionId(v string) *ModifyApisecApiResourceRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyApisecApiResourceRequest) SetResourceManagerResourceGroupId(v string) *ModifyApisecApiResourceRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *ModifyApisecApiResourceRequest) Validate() error {
	return dara.Validate(s)
}
