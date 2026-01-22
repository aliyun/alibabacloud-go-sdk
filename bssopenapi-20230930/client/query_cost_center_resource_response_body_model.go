// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryCostCenterResourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCostCenterResourceDtoList(v []*QueryCostCenterResourceResponseBodyCostCenterResourceDtoList) *QueryCostCenterResourceResponseBody
	GetCostCenterResourceDtoList() []*QueryCostCenterResourceResponseBodyCostCenterResourceDtoList
	SetMaxResults(v int32) *QueryCostCenterResourceResponseBody
	GetMaxResults() *int32
	SetMetadata(v interface{}) *QueryCostCenterResourceResponseBody
	GetMetadata() interface{}
	SetNextToken(v string) *QueryCostCenterResourceResponseBody
	GetNextToken() *string
	SetRequestId(v string) *QueryCostCenterResourceResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *QueryCostCenterResourceResponseBody
	GetTotalCount() *int32
}

type QueryCostCenterResourceResponseBody struct {
	CostCenterResourceDtoList []*QueryCostCenterResourceResponseBodyCostCenterResourceDtoList `json:"CostCenterResourceDtoList,omitempty" xml:"CostCenterResourceDtoList,omitempty" type:"Repeated"`
	// example:
	//
	// 200
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// {}
	Metadata interface{} `json:"Metadata,omitempty" xml:"Metadata,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// eyJwYWdlTnVtIjoyLCJwYWdlU2l6ZSI6MTB9
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// 79EE7556-0CFD-44EB-9CD6-B3B526E3A85F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 2000
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s QueryCostCenterResourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryCostCenterResourceResponseBody) GoString() string {
	return s.String()
}

func (s *QueryCostCenterResourceResponseBody) GetCostCenterResourceDtoList() []*QueryCostCenterResourceResponseBodyCostCenterResourceDtoList {
	return s.CostCenterResourceDtoList
}

func (s *QueryCostCenterResourceResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *QueryCostCenterResourceResponseBody) GetMetadata() interface{} {
	return s.Metadata
}

func (s *QueryCostCenterResourceResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *QueryCostCenterResourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryCostCenterResourceResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *QueryCostCenterResourceResponseBody) SetCostCenterResourceDtoList(v []*QueryCostCenterResourceResponseBodyCostCenterResourceDtoList) *QueryCostCenterResourceResponseBody {
	s.CostCenterResourceDtoList = v
	return s
}

func (s *QueryCostCenterResourceResponseBody) SetMaxResults(v int32) *QueryCostCenterResourceResponseBody {
	s.MaxResults = &v
	return s
}

func (s *QueryCostCenterResourceResponseBody) SetMetadata(v interface{}) *QueryCostCenterResourceResponseBody {
	s.Metadata = v
	return s
}

func (s *QueryCostCenterResourceResponseBody) SetNextToken(v string) *QueryCostCenterResourceResponseBody {
	s.NextToken = &v
	return s
}

func (s *QueryCostCenterResourceResponseBody) SetRequestId(v string) *QueryCostCenterResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryCostCenterResourceResponseBody) SetTotalCount(v int32) *QueryCostCenterResourceResponseBody {
	s.TotalCount = &v
	return s
}

func (s *QueryCostCenterResourceResponseBody) Validate() error {
	if s.CostCenterResourceDtoList != nil {
		for _, item := range s.CostCenterResourceDtoList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryCostCenterResourceResponseBodyCostCenterResourceDtoList struct {
	// example:
	//
	// AUTO_ALLOCATE
	AddStrategy *string `json:"AddStrategy,omitempty" xml:"AddStrategy,omitempty"`
	// example:
	//
	// 自动分配
	AddStrategyName *string `json:"AddStrategyName,omitempty" xml:"AddStrategyName,omitempty"`
	// example:
	//
	// 3
	ApplicablePeriodNum *int64 `json:"ApplicablePeriodNum,omitempty" xml:"ApplicablePeriodNum,omitempty"`
	// example:
	//
	// test
	ApportionItemCode *string `json:"ApportionItemCode,omitempty" xml:"ApportionItemCode,omitempty"`
	// example:
	//
	// test
	ApportionItemName *string `json:"ApportionItemName,omitempty" xml:"ApportionItemName,omitempty"`
	// example:
	//
	// otsbag
	CommodityCode *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	CommodityName *string `json:"CommodityName,omitempty" xml:"CommodityName,omitempty"`
	// example:
	//
	// code
	CostCenterCode *string `json:"CostCenterCode,omitempty" xml:"CostCenterCode,omitempty"`
	// example:
	//
	// 2025-05-18 12:12:25
	CostCenterCreateTime *string `json:"CostCenterCreateTime,omitempty" xml:"CostCenterCreateTime,omitempty"`
	// example:
	//
	// 123456
	CostCenterId *int64 `json:"CostCenterId,omitempty" xml:"CostCenterId,omitempty"`
	// example:
	//
	// test
	CostCenterName *string `json:"CostCenterName,omitempty" xml:"CostCenterName,omitempty"`
	// example:
	//
	// 2025-05-18 16:12:25
	CostCenterUpdateTime *string `json:"CostCenterUpdateTime,omitempty" xml:"CostCenterUpdateTime,omitempty"`
	// example:
	//
	// 3
	FinanceUnitRuleVersion *int64  `json:"FinanceUnitRuleVersion,omitempty" xml:"FinanceUnitRuleVersion,omitempty"`
	InstanceId             *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// ecs
	MasterCommodityCode *string `json:"MasterCommodityCode,omitempty" xml:"MasterCommodityCode,omitempty"`
	// example:
	//
	// i-xxxxx
	MasterInstanceId *string `json:"MasterInstanceId,omitempty" xml:"MasterInstanceId,omitempty"`
	// example:
	//
	// 1234567812345678
	OwnerAccountId *int64 `json:"OwnerAccountId,omitempty" xml:"OwnerAccountId,omitempty"`
	// example:
	//
	// XXX公司
	OwnerAccountName *string `json:"OwnerAccountName,omitempty" xml:"OwnerAccountName,omitempty"`
	// example:
	//
	// 123456
	ParentCostCenterId *int64 `json:"ParentCostCenterId,omitempty" xml:"ParentCostCenterId,omitempty"`
	// example:
	//
	// rds
	PipCode *string `json:"PipCode,omitempty" xml:"PipCode,omitempty"`
	PipName *string `json:"PipName,omitempty" xml:"PipName,omitempty"`
	// example:
	//
	// 202509
	RecentBillingMonth *int64 `json:"RecentBillingMonth,omitempty" xml:"RecentBillingMonth,omitempty"`
	// example:
	//
	// 上海
	RegionName *string `json:"RegionName,omitempty" xml:"RegionName,omitempty"`
	// example:
	//
	// cn-shanghai
	RegionNo      *string `json:"RegionNo,omitempty" xml:"RegionNo,omitempty"`
	ResourceGroup *string `json:"ResourceGroup,omitempty" xml:"ResourceGroup,omitempty"`
	// example:
	//
	// OSSBAG-cn-v0h1s4hma018
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// example:
	//
	// testResource
	ResourceNick *string `json:"ResourceNick,omitempty" xml:"ResourceNick,omitempty"`
	// example:
	//
	// MANUAL_ALLOCATE
	ResourceSource *string `json:"ResourceSource,omitempty" xml:"ResourceSource,omitempty"`
	// example:
	//
	// tag
	ResourceTag *string `json:"ResourceTag,omitempty" xml:"ResourceTag,omitempty"`
	// example:
	//
	// FPT_ossbag_absolute_Storage_bj
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// example:
	//
	// 2025-05-18 16:12:25
	ResourceUpdateTime *string `json:"ResourceUpdateTime,omitempty" xml:"ResourceUpdateTime,omitempty"`
	// example:
	//
	// 1234567812345678
	ResourceUserId *int64 `json:"ResourceUserId,omitempty" xml:"ResourceUserId,omitempty"`
	// example:
	//
	// test@test.aliyun.com
	ResourceUserName *string `json:"ResourceUserName,omitempty" xml:"ResourceUserName,omitempty"`
	// example:
	//
	// -1
	RootCostCenterId *int64 `json:"RootCostCenterId,omitempty" xml:"RootCostCenterId,omitempty"`
	// example:
	//
	// 202509
	StartBillingMonth *int64 `json:"StartBillingMonth,omitempty" xml:"StartBillingMonth,omitempty"`
}

func (s QueryCostCenterResourceResponseBodyCostCenterResourceDtoList) String() string {
	return dara.Prettify(s)
}

func (s QueryCostCenterResourceResponseBodyCostCenterResourceDtoList) GoString() string {
	return s.String()
}

func (s *QueryCostCenterResourceResponseBodyCostCenterResourceDtoList) GetAddStrategy() *string {
	return s.AddStrategy
}

func (s *QueryCostCenterResourceResponseBodyCostCenterResourceDtoList) GetAddStrategyName() *string {
	return s.AddStrategyName
}

func (s *QueryCostCenterResourceResponseBodyCostCenterResourceDtoList) GetApplicablePeriodNum() *int64 {
	return s.ApplicablePeriodNum
}

func (s *QueryCostCenterResourceResponseBodyCostCenterResourceDtoList) GetApportionItemCode() *string {
	return s.ApportionItemCode
}

func (s *QueryCostCenterResourceResponseBodyCostCenterResourceDtoList) GetApportionItemName() *string {
	return s.ApportionItemName
}

func (s *QueryCostCenterResourceResponseBodyCostCenterResourceDtoList) GetCommodityCode() *string {
	return s.CommodityCode
}

func (s *QueryCostCenterResourceResponseBodyCostCenterResourceDtoList) GetCommodityName() *string {
	return s.CommodityName
}

func (s *QueryCostCenterResourceResponseBodyCostCenterResourceDtoList) GetCostCenterCode() *string {
	return s.CostCenterCode
}

func (s *QueryCostCenterResourceResponseBodyCostCenterResourceDtoList) GetCostCenterCreateTime() *string {
	return s.CostCenterCreateTime
}

func (s *QueryCostCenterResourceResponseBodyCostCenterResourceDtoList) GetCostCenterId() *int64 {
	return s.CostCenterId
}

func (s *QueryCostCenterResourceResponseBodyCostCenterResourceDtoList) GetCostCenterName() *string {
	return s.CostCenterName
}

func (s *QueryCostCenterResourceResponseBodyCostCenterResourceDtoList) GetCostCenterUpdateTime() *string {
	return s.CostCenterUpdateTime
}

func (s *QueryCostCenterResourceResponseBodyCostCenterResourceDtoList) GetFinanceUnitRuleVersion() *int64 {
	return s.FinanceUnitRuleVersion
}

func (s *QueryCostCenterResourceResponseBodyCostCenterResourceDtoList) GetInstanceId() *string {
	return s.InstanceId
}

func (s *QueryCostCenterResourceResponseBodyCostCenterResourceDtoList) GetMasterCommodityCode() *string {
	return s.MasterCommodityCode
}

func (s *QueryCostCenterResourceResponseBodyCostCenterResourceDtoList) GetMasterInstanceId() *string {
	return s.MasterInstanceId
}

func (s *QueryCostCenterResourceResponseBodyCostCenterResourceDtoList) GetOwnerAccountId() *int64 {
	return s.OwnerAccountId
}

func (s *QueryCostCenterResourceResponseBodyCostCenterResourceDtoList) GetOwnerAccountName() *string {
	return s.OwnerAccountName
}

func (s *QueryCostCenterResourceResponseBodyCostCenterResourceDtoList) GetParentCostCenterId() *int64 {
	return s.ParentCostCenterId
}

func (s *QueryCostCenterResourceResponseBodyCostCenterResourceDtoList) GetPipCode() *string {
	return s.PipCode
}

func (s *QueryCostCenterResourceResponseBodyCostCenterResourceDtoList) GetPipName() *string {
	return s.PipName
}

func (s *QueryCostCenterResourceResponseBodyCostCenterResourceDtoList) GetRecentBillingMonth() *int64 {
	return s.RecentBillingMonth
}

func (s *QueryCostCenterResourceResponseBodyCostCenterResourceDtoList) GetRegionName() *string {
	return s.RegionName
}

func (s *QueryCostCenterResourceResponseBodyCostCenterResourceDtoList) GetRegionNo() *string {
	return s.RegionNo
}

func (s *QueryCostCenterResourceResponseBodyCostCenterResourceDtoList) GetResourceGroup() *string {
	return s.ResourceGroup
}

func (s *QueryCostCenterResourceResponseBodyCostCenterResourceDtoList) GetResourceId() *string {
	return s.ResourceId
}

func (s *QueryCostCenterResourceResponseBodyCostCenterResourceDtoList) GetResourceNick() *string {
	return s.ResourceNick
}

func (s *QueryCostCenterResourceResponseBodyCostCenterResourceDtoList) GetResourceSource() *string {
	return s.ResourceSource
}

func (s *QueryCostCenterResourceResponseBodyCostCenterResourceDtoList) GetResourceTag() *string {
	return s.ResourceTag
}

func (s *QueryCostCenterResourceResponseBodyCostCenterResourceDtoList) GetResourceType() *string {
	return s.ResourceType
}

func (s *QueryCostCenterResourceResponseBodyCostCenterResourceDtoList) GetResourceUpdateTime() *string {
	return s.ResourceUpdateTime
}

func (s *QueryCostCenterResourceResponseBodyCostCenterResourceDtoList) GetResourceUserId() *int64 {
	return s.ResourceUserId
}

func (s *QueryCostCenterResourceResponseBodyCostCenterResourceDtoList) GetResourceUserName() *string {
	return s.ResourceUserName
}

func (s *QueryCostCenterResourceResponseBodyCostCenterResourceDtoList) GetRootCostCenterId() *int64 {
	return s.RootCostCenterId
}

func (s *QueryCostCenterResourceResponseBodyCostCenterResourceDtoList) GetStartBillingMonth() *int64 {
	return s.StartBillingMonth
}

func (s *QueryCostCenterResourceResponseBodyCostCenterResourceDtoList) SetAddStrategy(v string) *QueryCostCenterResourceResponseBodyCostCenterResourceDtoList {
	s.AddStrategy = &v
	return s
}

func (s *QueryCostCenterResourceResponseBodyCostCenterResourceDtoList) SetAddStrategyName(v string) *QueryCostCenterResourceResponseBodyCostCenterResourceDtoList {
	s.AddStrategyName = &v
	return s
}

func (s *QueryCostCenterResourceResponseBodyCostCenterResourceDtoList) SetApplicablePeriodNum(v int64) *QueryCostCenterResourceResponseBodyCostCenterResourceDtoList {
	s.ApplicablePeriodNum = &v
	return s
}

func (s *QueryCostCenterResourceResponseBodyCostCenterResourceDtoList) SetApportionItemCode(v string) *QueryCostCenterResourceResponseBodyCostCenterResourceDtoList {
	s.ApportionItemCode = &v
	return s
}

func (s *QueryCostCenterResourceResponseBodyCostCenterResourceDtoList) SetApportionItemName(v string) *QueryCostCenterResourceResponseBodyCostCenterResourceDtoList {
	s.ApportionItemName = &v
	return s
}

func (s *QueryCostCenterResourceResponseBodyCostCenterResourceDtoList) SetCommodityCode(v string) *QueryCostCenterResourceResponseBodyCostCenterResourceDtoList {
	s.CommodityCode = &v
	return s
}

func (s *QueryCostCenterResourceResponseBodyCostCenterResourceDtoList) SetCommodityName(v string) *QueryCostCenterResourceResponseBodyCostCenterResourceDtoList {
	s.CommodityName = &v
	return s
}

func (s *QueryCostCenterResourceResponseBodyCostCenterResourceDtoList) SetCostCenterCode(v string) *QueryCostCenterResourceResponseBodyCostCenterResourceDtoList {
	s.CostCenterCode = &v
	return s
}

func (s *QueryCostCenterResourceResponseBodyCostCenterResourceDtoList) SetCostCenterCreateTime(v string) *QueryCostCenterResourceResponseBodyCostCenterResourceDtoList {
	s.CostCenterCreateTime = &v
	return s
}

func (s *QueryCostCenterResourceResponseBodyCostCenterResourceDtoList) SetCostCenterId(v int64) *QueryCostCenterResourceResponseBodyCostCenterResourceDtoList {
	s.CostCenterId = &v
	return s
}

func (s *QueryCostCenterResourceResponseBodyCostCenterResourceDtoList) SetCostCenterName(v string) *QueryCostCenterResourceResponseBodyCostCenterResourceDtoList {
	s.CostCenterName = &v
	return s
}

func (s *QueryCostCenterResourceResponseBodyCostCenterResourceDtoList) SetCostCenterUpdateTime(v string) *QueryCostCenterResourceResponseBodyCostCenterResourceDtoList {
	s.CostCenterUpdateTime = &v
	return s
}

func (s *QueryCostCenterResourceResponseBodyCostCenterResourceDtoList) SetFinanceUnitRuleVersion(v int64) *QueryCostCenterResourceResponseBodyCostCenterResourceDtoList {
	s.FinanceUnitRuleVersion = &v
	return s
}

func (s *QueryCostCenterResourceResponseBodyCostCenterResourceDtoList) SetInstanceId(v string) *QueryCostCenterResourceResponseBodyCostCenterResourceDtoList {
	s.InstanceId = &v
	return s
}

func (s *QueryCostCenterResourceResponseBodyCostCenterResourceDtoList) SetMasterCommodityCode(v string) *QueryCostCenterResourceResponseBodyCostCenterResourceDtoList {
	s.MasterCommodityCode = &v
	return s
}

func (s *QueryCostCenterResourceResponseBodyCostCenterResourceDtoList) SetMasterInstanceId(v string) *QueryCostCenterResourceResponseBodyCostCenterResourceDtoList {
	s.MasterInstanceId = &v
	return s
}

func (s *QueryCostCenterResourceResponseBodyCostCenterResourceDtoList) SetOwnerAccountId(v int64) *QueryCostCenterResourceResponseBodyCostCenterResourceDtoList {
	s.OwnerAccountId = &v
	return s
}

func (s *QueryCostCenterResourceResponseBodyCostCenterResourceDtoList) SetOwnerAccountName(v string) *QueryCostCenterResourceResponseBodyCostCenterResourceDtoList {
	s.OwnerAccountName = &v
	return s
}

func (s *QueryCostCenterResourceResponseBodyCostCenterResourceDtoList) SetParentCostCenterId(v int64) *QueryCostCenterResourceResponseBodyCostCenterResourceDtoList {
	s.ParentCostCenterId = &v
	return s
}

func (s *QueryCostCenterResourceResponseBodyCostCenterResourceDtoList) SetPipCode(v string) *QueryCostCenterResourceResponseBodyCostCenterResourceDtoList {
	s.PipCode = &v
	return s
}

func (s *QueryCostCenterResourceResponseBodyCostCenterResourceDtoList) SetPipName(v string) *QueryCostCenterResourceResponseBodyCostCenterResourceDtoList {
	s.PipName = &v
	return s
}

func (s *QueryCostCenterResourceResponseBodyCostCenterResourceDtoList) SetRecentBillingMonth(v int64) *QueryCostCenterResourceResponseBodyCostCenterResourceDtoList {
	s.RecentBillingMonth = &v
	return s
}

func (s *QueryCostCenterResourceResponseBodyCostCenterResourceDtoList) SetRegionName(v string) *QueryCostCenterResourceResponseBodyCostCenterResourceDtoList {
	s.RegionName = &v
	return s
}

func (s *QueryCostCenterResourceResponseBodyCostCenterResourceDtoList) SetRegionNo(v string) *QueryCostCenterResourceResponseBodyCostCenterResourceDtoList {
	s.RegionNo = &v
	return s
}

func (s *QueryCostCenterResourceResponseBodyCostCenterResourceDtoList) SetResourceGroup(v string) *QueryCostCenterResourceResponseBodyCostCenterResourceDtoList {
	s.ResourceGroup = &v
	return s
}

func (s *QueryCostCenterResourceResponseBodyCostCenterResourceDtoList) SetResourceId(v string) *QueryCostCenterResourceResponseBodyCostCenterResourceDtoList {
	s.ResourceId = &v
	return s
}

func (s *QueryCostCenterResourceResponseBodyCostCenterResourceDtoList) SetResourceNick(v string) *QueryCostCenterResourceResponseBodyCostCenterResourceDtoList {
	s.ResourceNick = &v
	return s
}

func (s *QueryCostCenterResourceResponseBodyCostCenterResourceDtoList) SetResourceSource(v string) *QueryCostCenterResourceResponseBodyCostCenterResourceDtoList {
	s.ResourceSource = &v
	return s
}

func (s *QueryCostCenterResourceResponseBodyCostCenterResourceDtoList) SetResourceTag(v string) *QueryCostCenterResourceResponseBodyCostCenterResourceDtoList {
	s.ResourceTag = &v
	return s
}

func (s *QueryCostCenterResourceResponseBodyCostCenterResourceDtoList) SetResourceType(v string) *QueryCostCenterResourceResponseBodyCostCenterResourceDtoList {
	s.ResourceType = &v
	return s
}

func (s *QueryCostCenterResourceResponseBodyCostCenterResourceDtoList) SetResourceUpdateTime(v string) *QueryCostCenterResourceResponseBodyCostCenterResourceDtoList {
	s.ResourceUpdateTime = &v
	return s
}

func (s *QueryCostCenterResourceResponseBodyCostCenterResourceDtoList) SetResourceUserId(v int64) *QueryCostCenterResourceResponseBodyCostCenterResourceDtoList {
	s.ResourceUserId = &v
	return s
}

func (s *QueryCostCenterResourceResponseBodyCostCenterResourceDtoList) SetResourceUserName(v string) *QueryCostCenterResourceResponseBodyCostCenterResourceDtoList {
	s.ResourceUserName = &v
	return s
}

func (s *QueryCostCenterResourceResponseBodyCostCenterResourceDtoList) SetRootCostCenterId(v int64) *QueryCostCenterResourceResponseBodyCostCenterResourceDtoList {
	s.RootCostCenterId = &v
	return s
}

func (s *QueryCostCenterResourceResponseBodyCostCenterResourceDtoList) SetStartBillingMonth(v int64) *QueryCostCenterResourceResponseBodyCostCenterResourceDtoList {
	s.StartBillingMonth = &v
	return s
}

func (s *QueryCostCenterResourceResponseBodyCostCenterResourceDtoList) Validate() error {
	return dara.Validate(s)
}
