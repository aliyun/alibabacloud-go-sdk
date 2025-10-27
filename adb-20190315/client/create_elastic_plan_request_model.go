// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateElasticPlanRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *CreateElasticPlanRequest
	GetDBClusterId() *string
	SetElasticPlanEnable(v bool) *CreateElasticPlanRequest
	GetElasticPlanEnable() *bool
	SetElasticPlanEndDay(v string) *CreateElasticPlanRequest
	GetElasticPlanEndDay() *string
	SetElasticPlanMonthlyRepeat(v string) *CreateElasticPlanRequest
	GetElasticPlanMonthlyRepeat() *string
	SetElasticPlanName(v string) *CreateElasticPlanRequest
	GetElasticPlanName() *string
	SetElasticPlanNodeNum(v int32) *CreateElasticPlanRequest
	GetElasticPlanNodeNum() *int32
	SetElasticPlanStartDay(v string) *CreateElasticPlanRequest
	GetElasticPlanStartDay() *string
	SetElasticPlanTimeEnd(v string) *CreateElasticPlanRequest
	GetElasticPlanTimeEnd() *string
	SetElasticPlanTimeStart(v string) *CreateElasticPlanRequest
	GetElasticPlanTimeStart() *string
	SetElasticPlanType(v string) *CreateElasticPlanRequest
	GetElasticPlanType() *string
	SetElasticPlanWeeklyRepeat(v string) *CreateElasticPlanRequest
	GetElasticPlanWeeklyRepeat() *string
	SetElasticPlanWorkerSpec(v string) *CreateElasticPlanRequest
	GetElasticPlanWorkerSpec() *string
	SetOwnerAccount(v string) *CreateElasticPlanRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateElasticPlanRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *CreateElasticPlanRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateElasticPlanRequest
	GetResourceOwnerId() *int64
	SetResourcePoolName(v string) *CreateElasticPlanRequest
	GetResourcePoolName() *string
}

type CreateElasticPlanRequest struct {
	// The ID of the AnalyticDB for MySQL Data Warehouse Edition (V3.0) cluster.
	//
	// > You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/129857.html) operation to query the IDs of all AnalyticDB for MySQL Data Warehouse Edition (V3.0) clusters within a region.
	//
	// This parameter is required.
	//
	// example:
	//
	// am-bp1d8lbdj22rx****
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
	// This parameter is required.
	//
	// example:
	//
	// test
	ElasticPlanName *string `json:"ElasticPlanName,omitempty" xml:"ElasticPlanName,omitempty"`
	// The number of nodes that are involved in the scaling plan.
	//
	// 	- If ElasticPlanType is set to **worker**, you can set this parameter to 0 or leave this parameter empty.
	//
	// 	- If ElasticPlanType is set to **executorcombineworker*	- or **executor**, you must set this parameter to a value that is greater than 0.
	//
	// if can be null:
	// false
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
	// This parameter is required.
	//
	// example:
	//
	// 10:00:00
	ElasticPlanTimeEnd *string `json:"ElasticPlanTimeEnd,omitempty" xml:"ElasticPlanTimeEnd,omitempty"`
	// The scale-up time of the scaling plan. Specify the time on the hour in the HH:mm:ss format.
	//
	// This parameter is required.
	//
	// example:
	//
	// 08:00:00
	ElasticPlanTimeStart *string `json:"ElasticPlanTimeStart,omitempty" xml:"ElasticPlanTimeStart,omitempty"`
	// The type of the scaling plan. Valid values:
	//
	// 	- **worker**: scales only elastic I/O resources.
	//
	// 	- **executor**: scales only computing resources.
	//
	// 	- **executorcombineworker*	- (default): scales both elastic I/O resources and computing resources by proportion.
	//
	// > - If you want to set this parameter to **executorcombineworker**, make sure that the cluster runs a minor version of 3.1.3.2 or later.
	//
	// > - If you want to set this parameter to **worker*	- or **executor**, make sure that the cluster runs a minor version of 3.1.6.1 or later and a ticket is submitted. After your request is approved, you can set this parameter to **worker*	- or **executor**.
	//
	// example:
	//
	// worker
	ElasticPlanType *string `json:"ElasticPlanType,omitempty" xml:"ElasticPlanType,omitempty"`
	// The days of the week when you want to execute the scaling plan. Valid values: 0 to 6 (Sunday to Saturday). Separate multiple values with commas (,).
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
	// 32 Core 64 GB
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
	// realtime
	ResourcePoolName *string `json:"ResourcePoolName,omitempty" xml:"ResourcePoolName,omitempty"`
}

func (s CreateElasticPlanRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateElasticPlanRequest) GoString() string {
	return s.String()
}

func (s *CreateElasticPlanRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *CreateElasticPlanRequest) GetElasticPlanEnable() *bool {
	return s.ElasticPlanEnable
}

func (s *CreateElasticPlanRequest) GetElasticPlanEndDay() *string {
	return s.ElasticPlanEndDay
}

func (s *CreateElasticPlanRequest) GetElasticPlanMonthlyRepeat() *string {
	return s.ElasticPlanMonthlyRepeat
}

func (s *CreateElasticPlanRequest) GetElasticPlanName() *string {
	return s.ElasticPlanName
}

func (s *CreateElasticPlanRequest) GetElasticPlanNodeNum() *int32 {
	return s.ElasticPlanNodeNum
}

func (s *CreateElasticPlanRequest) GetElasticPlanStartDay() *string {
	return s.ElasticPlanStartDay
}

func (s *CreateElasticPlanRequest) GetElasticPlanTimeEnd() *string {
	return s.ElasticPlanTimeEnd
}

func (s *CreateElasticPlanRequest) GetElasticPlanTimeStart() *string {
	return s.ElasticPlanTimeStart
}

func (s *CreateElasticPlanRequest) GetElasticPlanType() *string {
	return s.ElasticPlanType
}

func (s *CreateElasticPlanRequest) GetElasticPlanWeeklyRepeat() *string {
	return s.ElasticPlanWeeklyRepeat
}

func (s *CreateElasticPlanRequest) GetElasticPlanWorkerSpec() *string {
	return s.ElasticPlanWorkerSpec
}

func (s *CreateElasticPlanRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateElasticPlanRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateElasticPlanRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateElasticPlanRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateElasticPlanRequest) GetResourcePoolName() *string {
	return s.ResourcePoolName
}

func (s *CreateElasticPlanRequest) SetDBClusterId(v string) *CreateElasticPlanRequest {
	s.DBClusterId = &v
	return s
}

func (s *CreateElasticPlanRequest) SetElasticPlanEnable(v bool) *CreateElasticPlanRequest {
	s.ElasticPlanEnable = &v
	return s
}

func (s *CreateElasticPlanRequest) SetElasticPlanEndDay(v string) *CreateElasticPlanRequest {
	s.ElasticPlanEndDay = &v
	return s
}

func (s *CreateElasticPlanRequest) SetElasticPlanMonthlyRepeat(v string) *CreateElasticPlanRequest {
	s.ElasticPlanMonthlyRepeat = &v
	return s
}

func (s *CreateElasticPlanRequest) SetElasticPlanName(v string) *CreateElasticPlanRequest {
	s.ElasticPlanName = &v
	return s
}

func (s *CreateElasticPlanRequest) SetElasticPlanNodeNum(v int32) *CreateElasticPlanRequest {
	s.ElasticPlanNodeNum = &v
	return s
}

func (s *CreateElasticPlanRequest) SetElasticPlanStartDay(v string) *CreateElasticPlanRequest {
	s.ElasticPlanStartDay = &v
	return s
}

func (s *CreateElasticPlanRequest) SetElasticPlanTimeEnd(v string) *CreateElasticPlanRequest {
	s.ElasticPlanTimeEnd = &v
	return s
}

func (s *CreateElasticPlanRequest) SetElasticPlanTimeStart(v string) *CreateElasticPlanRequest {
	s.ElasticPlanTimeStart = &v
	return s
}

func (s *CreateElasticPlanRequest) SetElasticPlanType(v string) *CreateElasticPlanRequest {
	s.ElasticPlanType = &v
	return s
}

func (s *CreateElasticPlanRequest) SetElasticPlanWeeklyRepeat(v string) *CreateElasticPlanRequest {
	s.ElasticPlanWeeklyRepeat = &v
	return s
}

func (s *CreateElasticPlanRequest) SetElasticPlanWorkerSpec(v string) *CreateElasticPlanRequest {
	s.ElasticPlanWorkerSpec = &v
	return s
}

func (s *CreateElasticPlanRequest) SetOwnerAccount(v string) *CreateElasticPlanRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateElasticPlanRequest) SetOwnerId(v int64) *CreateElasticPlanRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateElasticPlanRequest) SetResourceOwnerAccount(v string) *CreateElasticPlanRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateElasticPlanRequest) SetResourceOwnerId(v int64) *CreateElasticPlanRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateElasticPlanRequest) SetResourcePoolName(v string) *CreateElasticPlanRequest {
	s.ResourcePoolName = &v
	return s
}

func (s *CreateElasticPlanRequest) Validate() error {
	return dara.Validate(s)
}
