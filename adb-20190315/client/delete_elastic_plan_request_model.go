// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteElasticPlanRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DeleteElasticPlanRequest
	GetDBClusterId() *string
	SetElasticPlanName(v string) *DeleteElasticPlanRequest
	GetElasticPlanName() *string
	SetOwnerAccount(v string) *DeleteElasticPlanRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeleteElasticPlanRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DeleteElasticPlanRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteElasticPlanRequest
	GetResourceOwnerId() *int64
}

type DeleteElasticPlanRequest struct {
	// The ID of the AnalyticDB for MySQL Data Warehouse Edition cluster.
	//
	// >  You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/612241.html) operation to query the IDs of all AnalyticDB for MySQL Data Warehouse Edition clusters within a region.
	//
	// This parameter is required.
	//
	// example:
	//
	// am-bp1xxxxxxxx47
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The name of the scaling plan.
	//
	// > You can call the [DescribeElasticPlans](https://help.aliyun.com/document_detail/601334.html) operation to query the names of scaling plans.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	ElasticPlanName      *string `json:"ElasticPlanName,omitempty" xml:"ElasticPlanName,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DeleteElasticPlanRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteElasticPlanRequest) GoString() string {
	return s.String()
}

func (s *DeleteElasticPlanRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DeleteElasticPlanRequest) GetElasticPlanName() *string {
	return s.ElasticPlanName
}

func (s *DeleteElasticPlanRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeleteElasticPlanRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteElasticPlanRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteElasticPlanRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteElasticPlanRequest) SetDBClusterId(v string) *DeleteElasticPlanRequest {
	s.DBClusterId = &v
	return s
}

func (s *DeleteElasticPlanRequest) SetElasticPlanName(v string) *DeleteElasticPlanRequest {
	s.ElasticPlanName = &v
	return s
}

func (s *DeleteElasticPlanRequest) SetOwnerAccount(v string) *DeleteElasticPlanRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteElasticPlanRequest) SetOwnerId(v int64) *DeleteElasticPlanRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteElasticPlanRequest) SetResourceOwnerAccount(v string) *DeleteElasticPlanRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteElasticPlanRequest) SetResourceOwnerId(v int64) *DeleteElasticPlanRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteElasticPlanRequest) Validate() error {
	return dara.Validate(s)
}
