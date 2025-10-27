// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeElasticPlanRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeElasticPlanRequest
	GetDBClusterId() *string
	SetElasticPlanEnable(v bool) *DescribeElasticPlanRequest
	GetElasticPlanEnable() *bool
	SetElasticPlanName(v string) *DescribeElasticPlanRequest
	GetElasticPlanName() *string
	SetOwnerAccount(v string) *DescribeElasticPlanRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeElasticPlanRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DescribeElasticPlanRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeElasticPlanRequest
	GetResourceOwnerId() *int64
	SetResourcePoolName(v string) *DescribeElasticPlanRequest
	GetResourcePoolName() *string
}

type DescribeElasticPlanRequest struct {
	// The ID of the AnalyticDB for MySQL Data Warehouse Edition (V3.0) cluster.
	//
	// > You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/129857.html) operation to query the IDs of all AnalyticDB for MySQL Data Warehouse Edition (V3.0) clusters within a region.
	//
	// This parameter is required.
	//
	// example:
	//
	// am-bp278jg9****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// Specifies whether the scaling plan takes effect. Valid values:
	//
	// 	- **true*	- (default)
	//
	// 	- **false**
	//
	// example:
	//
	// true
	ElasticPlanEnable *bool `json:"ElasticPlanEnable,omitempty" xml:"ElasticPlanEnable,omitempty"`
	// The name of the scaling plan.
	//
	// 	- The name must be 2 to 30 characters in length.
	//
	// 	- The name can contain letters, digits, and underscores (_).
	//
	// > If you do not specify this parameter, the information about all scaling plans for the specified cluster is returned.
	//
	// example:
	//
	// realtime
	ElasticPlanName      *string `json:"ElasticPlanName,omitempty" xml:"ElasticPlanName,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The name of the resource group.
	//
	// > You can call the [DescribeDBResourceGroup](https://help.aliyun.com/document_detail/466685.html) operation to query the resource group name.
	//
	// example:
	//
	// USER_DEFAULT
	ResourcePoolName *string `json:"ResourcePoolName,omitempty" xml:"ResourcePoolName,omitempty"`
}

func (s DescribeElasticPlanRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeElasticPlanRequest) GoString() string {
	return s.String()
}

func (s *DescribeElasticPlanRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeElasticPlanRequest) GetElasticPlanEnable() *bool {
	return s.ElasticPlanEnable
}

func (s *DescribeElasticPlanRequest) GetElasticPlanName() *string {
	return s.ElasticPlanName
}

func (s *DescribeElasticPlanRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeElasticPlanRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeElasticPlanRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeElasticPlanRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeElasticPlanRequest) GetResourcePoolName() *string {
	return s.ResourcePoolName
}

func (s *DescribeElasticPlanRequest) SetDBClusterId(v string) *DescribeElasticPlanRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeElasticPlanRequest) SetElasticPlanEnable(v bool) *DescribeElasticPlanRequest {
	s.ElasticPlanEnable = &v
	return s
}

func (s *DescribeElasticPlanRequest) SetElasticPlanName(v string) *DescribeElasticPlanRequest {
	s.ElasticPlanName = &v
	return s
}

func (s *DescribeElasticPlanRequest) SetOwnerAccount(v string) *DescribeElasticPlanRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeElasticPlanRequest) SetOwnerId(v int64) *DescribeElasticPlanRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeElasticPlanRequest) SetResourceOwnerAccount(v string) *DescribeElasticPlanRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeElasticPlanRequest) SetResourceOwnerId(v int64) *DescribeElasticPlanRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeElasticPlanRequest) SetResourcePoolName(v string) *DescribeElasticPlanRequest {
	s.ResourcePoolName = &v
	return s
}

func (s *DescribeElasticPlanRequest) Validate() error {
	return dara.Validate(s)
}
