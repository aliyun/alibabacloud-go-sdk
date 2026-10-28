// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeElasticPlanSpecificationsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeElasticPlanSpecificationsRequest
	GetDBClusterId() *string
	SetResourceGroupName(v string) *DescribeElasticPlanSpecificationsRequest
	GetResourceGroupName() *string
	SetType(v string) *DescribeElasticPlanSpecificationsRequest
	GetType() *string
}

type DescribeElasticPlanSpecificationsRequest struct {
	// The ID of the cluster.
	//
	// > You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/454250.html) operation to query the cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// amv-wz9509beptiz****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The name of the resource group.
	//
	// >- To query the specifications supported by the elastic plan of an Interactive resource group (compute resource type), you must specify the resource group name. To query the specifications supported by the elastic plan of the storage acceleration service (EIU), you do not need to specify the resource group name.
	//
	// - You can call the [DescribeDBResourceGroup](https://help.aliyun.com/document_detail/459446.html) operation to query the resource group name of a specific cluster.
	//
	// example:
	//
	// test
	ResourceGroupName *string `json:"ResourceGroupName,omitempty" xml:"ResourceGroupName,omitempty"`
	// The type of the elastic plan. Valid values:
	//
	// - EXECUTOR: Interactive resource group, which is the compute resource type.
	//
	// - WORKER: storage acceleration service (EIU).
	//
	// This parameter is required.
	//
	// example:
	//
	// EXECUTOR
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeElasticPlanSpecificationsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeElasticPlanSpecificationsRequest) GoString() string {
	return s.String()
}

func (s *DescribeElasticPlanSpecificationsRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeElasticPlanSpecificationsRequest) GetResourceGroupName() *string {
	return s.ResourceGroupName
}

func (s *DescribeElasticPlanSpecificationsRequest) GetType() *string {
	return s.Type
}

func (s *DescribeElasticPlanSpecificationsRequest) SetDBClusterId(v string) *DescribeElasticPlanSpecificationsRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeElasticPlanSpecificationsRequest) SetResourceGroupName(v string) *DescribeElasticPlanSpecificationsRequest {
	s.ResourceGroupName = &v
	return s
}

func (s *DescribeElasticPlanSpecificationsRequest) SetType(v string) *DescribeElasticPlanSpecificationsRequest {
	s.Type = &v
	return s
}

func (s *DescribeElasticPlanSpecificationsRequest) Validate() error {
	return dara.Validate(s)
}
