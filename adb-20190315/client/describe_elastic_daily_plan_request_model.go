// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeElasticDailyPlanRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeElasticDailyPlanRequest
	GetDBClusterId() *string
	SetElasticDailyPlanDay(v string) *DescribeElasticDailyPlanRequest
	GetElasticDailyPlanDay() *string
	SetElasticDailyPlanStatusList(v string) *DescribeElasticDailyPlanRequest
	GetElasticDailyPlanStatusList() *string
	SetElasticPlanName(v string) *DescribeElasticDailyPlanRequest
	GetElasticPlanName() *string
	SetOwnerAccount(v string) *DescribeElasticDailyPlanRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeElasticDailyPlanRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DescribeElasticDailyPlanRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeElasticDailyPlanRequest
	GetResourceOwnerId() *int64
	SetResourcePoolName(v string) *DescribeElasticDailyPlanRequest
	GetResourcePoolName() *string
}

type DescribeElasticDailyPlanRequest struct {
	// The ID of the AnalyticDB for MySQL Data Warehouse Edition (V3.0) cluster.
	//
	// >  You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/129857.html) operation to query the cluster IDs of all AnalyticDB for MySQL Data Warehouse Edition (V3.0) clusters within a specific region.
	//
	// This parameter is required.
	//
	// example:
	//
	// am-bp11q28kvl688****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The start date of the current-day scaling plan. Specify the date in the yyyy-MM-dd format.
	//
	// example:
	//
	// 2022-12-02
	ElasticDailyPlanDay *string `json:"ElasticDailyPlanDay,omitempty" xml:"ElasticDailyPlanDay,omitempty"`
	// The execution state of the current-day scaling plan. Separate multiple values with commas (,). Valid values:
	//
	// 	- **1**: The scaling plan is not executed.
	//
	// 	- **2**: The scaling plan is being executed.
	//
	// 	- **3**: The scaling plan is executed.
	//
	// 	- **4**: The scaling plan fails to be executed.
	//
	// example:
	//
	// 3
	ElasticDailyPlanStatusList *string `json:"ElasticDailyPlanStatusList,omitempty" xml:"ElasticDailyPlanStatusList,omitempty"`
	// The name of the scaling plan. Valid values:
	//
	// 	- The name must be 2 to 30 characters in length.
	//
	// 	- The name can contain letters, digits, and underscores (_).
	//
	// example:
	//
	// realtimeplan
	ElasticPlanName      *string `json:"ElasticPlanName,omitempty" xml:"ElasticPlanName,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The name of the resource group.
	//
	// >  You can call the [DescribeDBResourceGroup](https://help.aliyun.com/document_detail/466685.html) operation to query the resource group name.
	//
	// example:
	//
	// test
	ResourcePoolName *string `json:"ResourcePoolName,omitempty" xml:"ResourcePoolName,omitempty"`
}

func (s DescribeElasticDailyPlanRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeElasticDailyPlanRequest) GoString() string {
	return s.String()
}

func (s *DescribeElasticDailyPlanRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeElasticDailyPlanRequest) GetElasticDailyPlanDay() *string {
	return s.ElasticDailyPlanDay
}

func (s *DescribeElasticDailyPlanRequest) GetElasticDailyPlanStatusList() *string {
	return s.ElasticDailyPlanStatusList
}

func (s *DescribeElasticDailyPlanRequest) GetElasticPlanName() *string {
	return s.ElasticPlanName
}

func (s *DescribeElasticDailyPlanRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeElasticDailyPlanRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeElasticDailyPlanRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeElasticDailyPlanRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeElasticDailyPlanRequest) GetResourcePoolName() *string {
	return s.ResourcePoolName
}

func (s *DescribeElasticDailyPlanRequest) SetDBClusterId(v string) *DescribeElasticDailyPlanRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeElasticDailyPlanRequest) SetElasticDailyPlanDay(v string) *DescribeElasticDailyPlanRequest {
	s.ElasticDailyPlanDay = &v
	return s
}

func (s *DescribeElasticDailyPlanRequest) SetElasticDailyPlanStatusList(v string) *DescribeElasticDailyPlanRequest {
	s.ElasticDailyPlanStatusList = &v
	return s
}

func (s *DescribeElasticDailyPlanRequest) SetElasticPlanName(v string) *DescribeElasticDailyPlanRequest {
	s.ElasticPlanName = &v
	return s
}

func (s *DescribeElasticDailyPlanRequest) SetOwnerAccount(v string) *DescribeElasticDailyPlanRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeElasticDailyPlanRequest) SetOwnerId(v int64) *DescribeElasticDailyPlanRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeElasticDailyPlanRequest) SetResourceOwnerAccount(v string) *DescribeElasticDailyPlanRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeElasticDailyPlanRequest) SetResourceOwnerId(v int64) *DescribeElasticDailyPlanRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeElasticDailyPlanRequest) SetResourcePoolName(v string) *DescribeElasticDailyPlanRequest {
	s.ResourcePoolName = &v
	return s
}

func (s *DescribeElasticDailyPlanRequest) Validate() error {
	return dara.Validate(s)
}
