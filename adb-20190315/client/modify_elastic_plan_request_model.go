// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyElasticPlanRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *ModifyElasticPlanRequest
	GetDBClusterId() *string
	SetElasticPlanEnable(v bool) *ModifyElasticPlanRequest
	GetElasticPlanEnable() *bool
	SetElasticPlanEndDay(v string) *ModifyElasticPlanRequest
	GetElasticPlanEndDay() *string
	SetElasticPlanMonthlyRepeat(v string) *ModifyElasticPlanRequest
	GetElasticPlanMonthlyRepeat() *string
	SetElasticPlanName(v string) *ModifyElasticPlanRequest
	GetElasticPlanName() *string
	SetElasticPlanNodeNum(v int32) *ModifyElasticPlanRequest
	GetElasticPlanNodeNum() *int32
	SetElasticPlanStartDay(v string) *ModifyElasticPlanRequest
	GetElasticPlanStartDay() *string
	SetElasticPlanTimeEnd(v string) *ModifyElasticPlanRequest
	GetElasticPlanTimeEnd() *string
	SetElasticPlanTimeStart(v string) *ModifyElasticPlanRequest
	GetElasticPlanTimeStart() *string
	SetElasticPlanType(v string) *ModifyElasticPlanRequest
	GetElasticPlanType() *string
	SetElasticPlanWeeklyRepeat(v string) *ModifyElasticPlanRequest
	GetElasticPlanWeeklyRepeat() *string
	SetElasticPlanWorkerSpec(v string) *ModifyElasticPlanRequest
	GetElasticPlanWorkerSpec() *string
	SetOwnerAccount(v string) *ModifyElasticPlanRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyElasticPlanRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ModifyElasticPlanRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyElasticPlanRequest
	GetResourceOwnerId() *int64
	SetResourcePoolName(v string) *ModifyElasticPlanRequest
	GetResourcePoolName() *string
}

type ModifyElasticPlanRequest struct {
	// The ID of the AnalyticDB for MySQL Data Warehouse Edition (V3.0) cluster.
	//
	// > You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/129857.html) operation to query the IDs of all AnalyticDB for MySQL Data Warehouse Edition (V3.0) clusters within a region.
	//
	// This parameter is required.
	//
	// example:
	//
	// am-bp1rqvm70uh2v****
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
	// The end date of the scaling plan. Specify the date in the yyyy-MM-dd format.
	//
	// example:
	//
	// 2022-12-09
	ElasticPlanEndDay *string `json:"ElasticPlanEndDay,omitempty" xml:"ElasticPlanEndDay,omitempty"`
	// The dates of the month when you want to execute the scaling plan. A number specifies a date of the month. Separate multiple values with commas (,).
	//
	// example:
	//
	// 1,15,25
	ElasticPlanMonthlyRepeat *string `json:"ElasticPlanMonthlyRepeat,omitempty" xml:"ElasticPlanMonthlyRepeat,omitempty"`
	// The name of the scaling plan.
	//
	// 	- The name must be 2 to 30 characters in length.
	//
	// 	- The name can contain letters, digits, and underscores (_).
	//
	// > You can call the [DescribeElasticPlan](https://help.aliyun.com/document_detail/190596.html) operation to query the information about all scaling plans of a cluster, including the scaling plan names.
	//
	// This parameter is required.
	//
	// example:
	//
	// realtime
	ElasticPlanName *string `json:"ElasticPlanName,omitempty" xml:"ElasticPlanName,omitempty"`
	// The number of nodes that are involved in the scaling plan.
	//
	// 	- If ElasticPlanType is set to **worker**, you can set this parameter to 0 or leave this parameter empty.
	//
	// 	- If ElasticPlanType is set to **executorcombineworker*	- or **executor**, you must set this parameter to a value that is greater than 0.
	//
	// example:
	//
	// 0
	ElasticPlanNodeNum *int32 `json:"ElasticPlanNodeNum,omitempty" xml:"ElasticPlanNodeNum,omitempty"`
	// The start date of the scaling plan. Specify the date in the yyyy-MM-dd format.
	//
	// example:
	//
	// 2022-12-02
	ElasticPlanStartDay *string `json:"ElasticPlanStartDay,omitempty" xml:"ElasticPlanStartDay,omitempty"`
	// The restoration time of the scaling plan. Specify the time on the hour in the HH:mm:ss format. The interval between the scale-up time and the restoration time cannot be more than 24 hours.
	//
	// example:
	//
	// 10:00:00
	ElasticPlanTimeEnd *string `json:"ElasticPlanTimeEnd,omitempty" xml:"ElasticPlanTimeEnd,omitempty"`
	// The scale-up time of the scaling plan. Specify the time on the hour in the HH:mm:ss format.
	//
	// example:
	//
	// 8:00:00
	ElasticPlanTimeStart *string `json:"ElasticPlanTimeStart,omitempty" xml:"ElasticPlanTimeStart,omitempty"`
	// The type of the scaling plan. Valid values:
	//
	// 	- **worker**: scales only elastic I/O resources.
	//
	// 	- **executor**: scales only computing resources.
	//
	// 	- **executorcombineworker*	- (default): scales both elastic I/O resources and computing resources by proportion.
	//
	// >
	//
	// 	- If you want to set this parameter to **executorcombineworker**, make sure that the cluster runs a minor version of 3.1.3.2 or later.
	//
	// 	- If you want to set this parameter to **worker*	- or **executor**, make sure that the cluster runs a minor version of 3.1.6.1 or later and a ticket is submitted. After your request is approved, you can set this parameter to **worker*	- or **executor**.
	//
	// example:
	//
	// worker
	ElasticPlanType *string `json:"ElasticPlanType,omitempty" xml:"ElasticPlanType,omitempty"`
	// The days of the week when you want to execute the scaling plan. Valid values: 0 to 6, which indicate Sunday to Saturday. Separate multiple values with commas (,).
	//
	// example:
	//
	// 1,2,3,4,5
	ElasticPlanWeeklyRepeat *string `json:"ElasticPlanWeeklyRepeat,omitempty" xml:"ElasticPlanWeeklyRepeat,omitempty"`
	// The resource specifications that can be scaled up by the scaling plan. Valid values:
	//
	// 	- 8 Core 64 GB (default)
	//
	// 	- 16 Core 64 GB
	//
	// 	- 32 Core 64 GB
	//
	// 	- 64 Core 128 GB
	//
	// 	- 12 Core 96 GB
	//
	// 	- 24 Core 96 GB
	//
	// 	- 52 Core 86 GB
	//
	// example:
	//
	// 16 Core 64 GB
	ElasticPlanWorkerSpec *string `json:"ElasticPlanWorkerSpec,omitempty" xml:"ElasticPlanWorkerSpec,omitempty"`
	OwnerAccount          *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId               *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount  *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId       *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The name of the resource group.
	//
	// >  You can call the [DescribeDBResourceGroup](https://help.aliyun.com/document_detail/466685.html) operation to query the resource group name.
	//
	// example:
	//
	// test
	ResourcePoolName *string `json:"ResourcePoolName,omitempty" xml:"ResourcePoolName,omitempty"`
}

func (s ModifyElasticPlanRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyElasticPlanRequest) GoString() string {
	return s.String()
}

func (s *ModifyElasticPlanRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *ModifyElasticPlanRequest) GetElasticPlanEnable() *bool {
	return s.ElasticPlanEnable
}

func (s *ModifyElasticPlanRequest) GetElasticPlanEndDay() *string {
	return s.ElasticPlanEndDay
}

func (s *ModifyElasticPlanRequest) GetElasticPlanMonthlyRepeat() *string {
	return s.ElasticPlanMonthlyRepeat
}

func (s *ModifyElasticPlanRequest) GetElasticPlanName() *string {
	return s.ElasticPlanName
}

func (s *ModifyElasticPlanRequest) GetElasticPlanNodeNum() *int32 {
	return s.ElasticPlanNodeNum
}

func (s *ModifyElasticPlanRequest) GetElasticPlanStartDay() *string {
	return s.ElasticPlanStartDay
}

func (s *ModifyElasticPlanRequest) GetElasticPlanTimeEnd() *string {
	return s.ElasticPlanTimeEnd
}

func (s *ModifyElasticPlanRequest) GetElasticPlanTimeStart() *string {
	return s.ElasticPlanTimeStart
}

func (s *ModifyElasticPlanRequest) GetElasticPlanType() *string {
	return s.ElasticPlanType
}

func (s *ModifyElasticPlanRequest) GetElasticPlanWeeklyRepeat() *string {
	return s.ElasticPlanWeeklyRepeat
}

func (s *ModifyElasticPlanRequest) GetElasticPlanWorkerSpec() *string {
	return s.ElasticPlanWorkerSpec
}

func (s *ModifyElasticPlanRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyElasticPlanRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyElasticPlanRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyElasticPlanRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyElasticPlanRequest) GetResourcePoolName() *string {
	return s.ResourcePoolName
}

func (s *ModifyElasticPlanRequest) SetDBClusterId(v string) *ModifyElasticPlanRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyElasticPlanRequest) SetElasticPlanEnable(v bool) *ModifyElasticPlanRequest {
	s.ElasticPlanEnable = &v
	return s
}

func (s *ModifyElasticPlanRequest) SetElasticPlanEndDay(v string) *ModifyElasticPlanRequest {
	s.ElasticPlanEndDay = &v
	return s
}

func (s *ModifyElasticPlanRequest) SetElasticPlanMonthlyRepeat(v string) *ModifyElasticPlanRequest {
	s.ElasticPlanMonthlyRepeat = &v
	return s
}

func (s *ModifyElasticPlanRequest) SetElasticPlanName(v string) *ModifyElasticPlanRequest {
	s.ElasticPlanName = &v
	return s
}

func (s *ModifyElasticPlanRequest) SetElasticPlanNodeNum(v int32) *ModifyElasticPlanRequest {
	s.ElasticPlanNodeNum = &v
	return s
}

func (s *ModifyElasticPlanRequest) SetElasticPlanStartDay(v string) *ModifyElasticPlanRequest {
	s.ElasticPlanStartDay = &v
	return s
}

func (s *ModifyElasticPlanRequest) SetElasticPlanTimeEnd(v string) *ModifyElasticPlanRequest {
	s.ElasticPlanTimeEnd = &v
	return s
}

func (s *ModifyElasticPlanRequest) SetElasticPlanTimeStart(v string) *ModifyElasticPlanRequest {
	s.ElasticPlanTimeStart = &v
	return s
}

func (s *ModifyElasticPlanRequest) SetElasticPlanType(v string) *ModifyElasticPlanRequest {
	s.ElasticPlanType = &v
	return s
}

func (s *ModifyElasticPlanRequest) SetElasticPlanWeeklyRepeat(v string) *ModifyElasticPlanRequest {
	s.ElasticPlanWeeklyRepeat = &v
	return s
}

func (s *ModifyElasticPlanRequest) SetElasticPlanWorkerSpec(v string) *ModifyElasticPlanRequest {
	s.ElasticPlanWorkerSpec = &v
	return s
}

func (s *ModifyElasticPlanRequest) SetOwnerAccount(v string) *ModifyElasticPlanRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyElasticPlanRequest) SetOwnerId(v int64) *ModifyElasticPlanRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyElasticPlanRequest) SetResourceOwnerAccount(v string) *ModifyElasticPlanRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyElasticPlanRequest) SetResourceOwnerId(v int64) *ModifyElasticPlanRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyElasticPlanRequest) SetResourcePoolName(v string) *ModifyElasticPlanRequest {
	s.ResourcePoolName = &v
	return s
}

func (s *ModifyElasticPlanRequest) Validate() error {
	return dara.Validate(s)
}
