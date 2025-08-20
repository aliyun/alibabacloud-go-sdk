// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeElasticPlansRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeElasticPlansRequest
	GetDBClusterId() *string
	SetElasticPlanName(v string) *DescribeElasticPlansRequest
	GetElasticPlanName() *string
	SetEnabled(v bool) *DescribeElasticPlansRequest
	GetEnabled() *bool
	SetPageNumber(v int32) *DescribeElasticPlansRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeElasticPlansRequest
	GetPageSize() *int32
	SetResourceGroupName(v string) *DescribeElasticPlansRequest
	GetResourceGroupName() *string
	SetType(v string) *DescribeElasticPlansRequest
	GetType() *string
}

type DescribeElasticPlansRequest struct {
	// The cluster ID.
	//
	// >  You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/612397.html) operation to query the IDs of all AnalyticDB for MySQL clusters within a region.
	//
	// This parameter is required.
	//
	// example:
	//
	// amv-wz9509beptiz****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The name of the scaling plan.
	//
	// > If you do not specify this parameter, all scaling plans are queried.
	//
	// example:
	//
	// test
	ElasticPlanName *string `json:"ElasticPlanName,omitempty" xml:"ElasticPlanName,omitempty"`
	// Specifies whether to query the scaling plans that are immediately enabled after the plans are created. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The page number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The name of the resource group.
	//
	// > 	- If you do not specify this parameter, the scaling plans of all resource groups are queried, covering the interactive resource group type and the elastic I/O unit (EIU) type.
	//
	// >	- You can call the [DescribeDBResourceGroup](https://help.aliyun.com/document_detail/459446.html) operation to query the name of a resource group within a cluster.
	//
	// example:
	//
	// test
	ResourceGroupName *string `json:"ResourceGroupName,omitempty" xml:"ResourceGroupName,omitempty"`
	// The type of the scaling plan. Valid values:
	//
	// 	- **EXECUTOR**: the interactive resource group type, which specifies the computing resource type.
	//
	// 	- **WORKER**: the EIU type.
	//
	// example:
	//
	// EXECUTOR
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeElasticPlansRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeElasticPlansRequest) GoString() string {
	return s.String()
}

func (s *DescribeElasticPlansRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeElasticPlansRequest) GetElasticPlanName() *string {
	return s.ElasticPlanName
}

func (s *DescribeElasticPlansRequest) GetEnabled() *bool {
	return s.Enabled
}

func (s *DescribeElasticPlansRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeElasticPlansRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeElasticPlansRequest) GetResourceGroupName() *string {
	return s.ResourceGroupName
}

func (s *DescribeElasticPlansRequest) GetType() *string {
	return s.Type
}

func (s *DescribeElasticPlansRequest) SetDBClusterId(v string) *DescribeElasticPlansRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeElasticPlansRequest) SetElasticPlanName(v string) *DescribeElasticPlansRequest {
	s.ElasticPlanName = &v
	return s
}

func (s *DescribeElasticPlansRequest) SetEnabled(v bool) *DescribeElasticPlansRequest {
	s.Enabled = &v
	return s
}

func (s *DescribeElasticPlansRequest) SetPageNumber(v int32) *DescribeElasticPlansRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeElasticPlansRequest) SetPageSize(v int32) *DescribeElasticPlansRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeElasticPlansRequest) SetResourceGroupName(v string) *DescribeElasticPlansRequest {
	s.ResourceGroupName = &v
	return s
}

func (s *DescribeElasticPlansRequest) SetType(v string) *DescribeElasticPlansRequest {
	s.Type = &v
	return s
}

func (s *DescribeElasticPlansRequest) Validate() error {
	return dara.Validate(s)
}
