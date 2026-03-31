// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteApisecAbnormalsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAbnormalIds(v []*string) *DeleteApisecAbnormalsRequest
	GetAbnormalIds() []*string
	SetClusterId(v string) *DeleteApisecAbnormalsRequest
	GetClusterId() *string
	SetInstanceId(v string) *DeleteApisecAbnormalsRequest
	GetInstanceId() *string
	SetRegionId(v string) *DeleteApisecAbnormalsRequest
	GetRegionId() *string
	SetResourceManagerResourceGroupId(v string) *DeleteApisecAbnormalsRequest
	GetResourceManagerResourceGroupId() *string
}

type DeleteApisecAbnormalsRequest struct {
	// The risk IDs.
	//
	// This parameter is required.
	AbnormalIds []*string `json:"AbnormalIds,omitempty" xml:"AbnormalIds,omitempty" type:"Repeated"`
	// The ID of the hybrid cloud cluster.
	//
	// >For hybrid cloud scenarios only, you can call the [DescribeHybridCloudClusters](https://help.aliyun.com/document_detail/2849376.html) operation to query the hybrid cloud clusters.
	//
	// example:
	//
	// 428
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The ID of the Web Application Firewall (WAF) instance.
	//
	// >  You can call the [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) operation to query the ID of the WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf_v3prepaid_public_cn-nwy*****
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

func (s DeleteApisecAbnormalsRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteApisecAbnormalsRequest) GoString() string {
	return s.String()
}

func (s *DeleteApisecAbnormalsRequest) GetAbnormalIds() []*string {
	return s.AbnormalIds
}

func (s *DeleteApisecAbnormalsRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *DeleteApisecAbnormalsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteApisecAbnormalsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteApisecAbnormalsRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *DeleteApisecAbnormalsRequest) SetAbnormalIds(v []*string) *DeleteApisecAbnormalsRequest {
	s.AbnormalIds = v
	return s
}

func (s *DeleteApisecAbnormalsRequest) SetClusterId(v string) *DeleteApisecAbnormalsRequest {
	s.ClusterId = &v
	return s
}

func (s *DeleteApisecAbnormalsRequest) SetInstanceId(v string) *DeleteApisecAbnormalsRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteApisecAbnormalsRequest) SetRegionId(v string) *DeleteApisecAbnormalsRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteApisecAbnormalsRequest) SetResourceManagerResourceGroupId(v string) *DeleteApisecAbnormalsRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DeleteApisecAbnormalsRequest) Validate() error {
	return dara.Validate(s)
}
