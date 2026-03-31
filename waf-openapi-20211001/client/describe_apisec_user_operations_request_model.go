// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApisecUserOperationsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *DescribeApisecUserOperationsRequest
	GetClusterId() *string
	SetInstanceId(v string) *DescribeApisecUserOperationsRequest
	GetInstanceId() *string
	SetObjectId(v string) *DescribeApisecUserOperationsRequest
	GetObjectId() *string
	SetRegionId(v string) *DescribeApisecUserOperationsRequest
	GetRegionId() *string
	SetResourceManagerResourceGroupId(v string) *DescribeApisecUserOperationsRequest
	GetResourceManagerResourceGroupId() *string
	SetType(v string) *DescribeApisecUserOperationsRequest
	GetType() *string
}

type DescribeApisecUserOperationsRequest struct {
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
	// waf_v2_public_cn-wwo36ksck1e
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The object ID of the operation record.
	//
	// This parameter is required.
	//
	// example:
	//
	// fe8723e92e2037245014ab62161bbec8
	ObjectId *string `json:"ObjectId,omitempty" xml:"ObjectId,omitempty"`
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
	// rg-aek2ax2y5****pi
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
	// The type of the operation record. Valid values:
	//
	// 	- **abnormal**: risk detection
	//
	// 	- **event**: security event
	//
	// example:
	//
	// event
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeApisecUserOperationsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeApisecUserOperationsRequest) GoString() string {
	return s.String()
}

func (s *DescribeApisecUserOperationsRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeApisecUserOperationsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeApisecUserOperationsRequest) GetObjectId() *string {
	return s.ObjectId
}

func (s *DescribeApisecUserOperationsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeApisecUserOperationsRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *DescribeApisecUserOperationsRequest) GetType() *string {
	return s.Type
}

func (s *DescribeApisecUserOperationsRequest) SetClusterId(v string) *DescribeApisecUserOperationsRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeApisecUserOperationsRequest) SetInstanceId(v string) *DescribeApisecUserOperationsRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeApisecUserOperationsRequest) SetObjectId(v string) *DescribeApisecUserOperationsRequest {
	s.ObjectId = &v
	return s
}

func (s *DescribeApisecUserOperationsRequest) SetRegionId(v string) *DescribeApisecUserOperationsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeApisecUserOperationsRequest) SetResourceManagerResourceGroupId(v string) *DescribeApisecUserOperationsRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribeApisecUserOperationsRequest) SetType(v string) *DescribeApisecUserOperationsRequest {
	s.Type = &v
	return s
}

func (s *DescribeApisecUserOperationsRequest) Validate() error {
	return dara.Validate(s)
}
