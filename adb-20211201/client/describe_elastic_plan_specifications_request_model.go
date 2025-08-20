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
	// The cluster ID.
	//
	// >  You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/454250.html) operation to query the IDs of all AnalyticDB for MySQL clusters within a region.
	//
	// This parameter is required.
	//
	// example:
	//
	// amv-wz9509beptiz****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The name of the resource group.
	//
	// >
	//
	// 	- This parameter must be specified only when you query the resource specifications that are supported by an interactive resource group.
	//
	// 	- You can call the [DescribeDBResourceGroup](https://help.aliyun.com/document_detail/459446.html) operation to query the name of a resource group within a cluster.
	//
	// example:
	//
	// test
	ResourceGroupName *string `json:"ResourceGroupName,omitempty" xml:"ResourceGroupName,omitempty"`
	// The type of the scaling plan. Valid values:
	//
	// 	- EXECUTOR: the interactive resource group type, which specifies the computing resource type.
	//
	// 	- WORKER: the elastic I/O unit (EIU) type.
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
