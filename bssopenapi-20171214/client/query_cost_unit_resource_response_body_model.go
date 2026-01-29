// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryCostUnitResourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QueryCostUnitResourceResponseBody
	GetCode() *string
	SetData(v *QueryCostUnitResourceResponseBodyData) *QueryCostUnitResourceResponseBody
	GetData() *QueryCostUnitResourceResponseBodyData
	SetMessage(v string) *QueryCostUnitResourceResponseBody
	GetMessage() *string
	SetRequestId(v string) *QueryCostUnitResourceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryCostUnitResourceResponseBody
	GetSuccess() *bool
}

type QueryCostUnitResourceResponseBody struct {
	// The status code.
	//
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	Data *QueryCostUnitResourceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message.
	//
	// example:
	//
	// Successful！
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 04332CB7-9A57-4461-97E0-02821D044414
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryCostUnitResourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryCostUnitResourceResponseBody) GoString() string {
	return s.String()
}

func (s *QueryCostUnitResourceResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryCostUnitResourceResponseBody) GetData() *QueryCostUnitResourceResponseBodyData {
	return s.Data
}

func (s *QueryCostUnitResourceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryCostUnitResourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryCostUnitResourceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryCostUnitResourceResponseBody) SetCode(v string) *QueryCostUnitResourceResponseBody {
	s.Code = &v
	return s
}

func (s *QueryCostUnitResourceResponseBody) SetData(v *QueryCostUnitResourceResponseBodyData) *QueryCostUnitResourceResponseBody {
	s.Data = v
	return s
}

func (s *QueryCostUnitResourceResponseBody) SetMessage(v string) *QueryCostUnitResourceResponseBody {
	s.Message = &v
	return s
}

func (s *QueryCostUnitResourceResponseBody) SetRequestId(v string) *QueryCostUnitResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryCostUnitResourceResponseBody) SetSuccess(v bool) *QueryCostUnitResourceResponseBody {
	s.Success = &v
	return s
}

func (s *QueryCostUnitResourceResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryCostUnitResourceResponseBodyData struct {
	// The information about the cost center.
	CostUnit *QueryCostUnitResourceResponseBodyDataCostUnit `json:"CostUnit,omitempty" xml:"CostUnit,omitempty" type:"Struct"`
	// The statistical information about the cost center.
	CostUnitStatisInfo *QueryCostUnitResourceResponseBodyDataCostUnitStatisInfo `json:"CostUnitStatisInfo,omitempty" xml:"CostUnitStatisInfo,omitempty" type:"Struct"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The resource instances.
	ResourceInstanceDtoList []*QueryCostUnitResourceResponseBodyDataResourceInstanceDtoList `json:"ResourceInstanceDtoList,omitempty" xml:"ResourceInstanceDtoList,omitempty" type:"Repeated"`
	// The total number of returned entries.
	//
	// example:
	//
	// 12
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s QueryCostUnitResourceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryCostUnitResourceResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryCostUnitResourceResponseBodyData) GetCostUnit() *QueryCostUnitResourceResponseBodyDataCostUnit {
	return s.CostUnit
}

func (s *QueryCostUnitResourceResponseBodyData) GetCostUnitStatisInfo() *QueryCostUnitResourceResponseBodyDataCostUnitStatisInfo {
	return s.CostUnitStatisInfo
}

func (s *QueryCostUnitResourceResponseBodyData) GetPageNum() *int32 {
	return s.PageNum
}

func (s *QueryCostUnitResourceResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryCostUnitResourceResponseBodyData) GetResourceInstanceDtoList() []*QueryCostUnitResourceResponseBodyDataResourceInstanceDtoList {
	return s.ResourceInstanceDtoList
}

func (s *QueryCostUnitResourceResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *QueryCostUnitResourceResponseBodyData) SetCostUnit(v *QueryCostUnitResourceResponseBodyDataCostUnit) *QueryCostUnitResourceResponseBodyData {
	s.CostUnit = v
	return s
}

func (s *QueryCostUnitResourceResponseBodyData) SetCostUnitStatisInfo(v *QueryCostUnitResourceResponseBodyDataCostUnitStatisInfo) *QueryCostUnitResourceResponseBodyData {
	s.CostUnitStatisInfo = v
	return s
}

func (s *QueryCostUnitResourceResponseBodyData) SetPageNum(v int32) *QueryCostUnitResourceResponseBodyData {
	s.PageNum = &v
	return s
}

func (s *QueryCostUnitResourceResponseBodyData) SetPageSize(v int32) *QueryCostUnitResourceResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *QueryCostUnitResourceResponseBodyData) SetResourceInstanceDtoList(v []*QueryCostUnitResourceResponseBodyDataResourceInstanceDtoList) *QueryCostUnitResourceResponseBodyData {
	s.ResourceInstanceDtoList = v
	return s
}

func (s *QueryCostUnitResourceResponseBodyData) SetTotalCount(v int32) *QueryCostUnitResourceResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *QueryCostUnitResourceResponseBodyData) Validate() error {
	if s.CostUnit != nil {
		if err := s.CostUnit.Validate(); err != nil {
			return err
		}
	}
	if s.CostUnitStatisInfo != nil {
		if err := s.CostUnitStatisInfo.Validate(); err != nil {
			return err
		}
	}
	if s.ResourceInstanceDtoList != nil {
		for _, item := range s.ResourceInstanceDtoList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryCostUnitResourceResponseBodyDataCostUnit struct {
	// The user ID of the cost center owner.
	//
	// example:
	//
	// 321432
	OwnerUid *int64 `json:"OwnerUid,omitempty" xml:"OwnerUid,omitempty"`
	// The ID of the parent cost center. A value of -1 indicates the root cost center.
	//
	// example:
	//
	// 23421
	ParentUnitId *int64 `json:"ParentUnitId,omitempty" xml:"ParentUnitId,omitempty"`
	// The ID of the cost center.
	//
	// example:
	//
	// 123412
	UnitId *int64 `json:"UnitId,omitempty" xml:"UnitId,omitempty"`
	// The name of the cost center.
	//
	// example:
	//
	// test
	UnitName *string `json:"UnitName,omitempty" xml:"UnitName,omitempty"`
}

func (s QueryCostUnitResourceResponseBodyDataCostUnit) String() string {
	return dara.Prettify(s)
}

func (s QueryCostUnitResourceResponseBodyDataCostUnit) GoString() string {
	return s.String()
}

func (s *QueryCostUnitResourceResponseBodyDataCostUnit) GetOwnerUid() *int64 {
	return s.OwnerUid
}

func (s *QueryCostUnitResourceResponseBodyDataCostUnit) GetParentUnitId() *int64 {
	return s.ParentUnitId
}

func (s *QueryCostUnitResourceResponseBodyDataCostUnit) GetUnitId() *int64 {
	return s.UnitId
}

func (s *QueryCostUnitResourceResponseBodyDataCostUnit) GetUnitName() *string {
	return s.UnitName
}

func (s *QueryCostUnitResourceResponseBodyDataCostUnit) SetOwnerUid(v int64) *QueryCostUnitResourceResponseBodyDataCostUnit {
	s.OwnerUid = &v
	return s
}

func (s *QueryCostUnitResourceResponseBodyDataCostUnit) SetParentUnitId(v int64) *QueryCostUnitResourceResponseBodyDataCostUnit {
	s.ParentUnitId = &v
	return s
}

func (s *QueryCostUnitResourceResponseBodyDataCostUnit) SetUnitId(v int64) *QueryCostUnitResourceResponseBodyDataCostUnit {
	s.UnitId = &v
	return s
}

func (s *QueryCostUnitResourceResponseBodyDataCostUnit) SetUnitName(v string) *QueryCostUnitResourceResponseBodyDataCostUnit {
	s.UnitName = &v
	return s
}

func (s *QueryCostUnitResourceResponseBodyDataCostUnit) Validate() error {
	return dara.Validate(s)
}

type QueryCostUnitResourceResponseBodyDataCostUnitStatisInfo struct {
	// The number of resource instances in the cost center.
	//
	// example:
	//
	// 1
	ResourceCount *int64 `json:"ResourceCount,omitempty" xml:"ResourceCount,omitempty"`
	// The number of resource groups in the cost center.
	//
	// example:
	//
	// 1
	ResourceGroupCount *int64 `json:"ResourceGroupCount,omitempty" xml:"ResourceGroupCount,omitempty"`
	// The number of sub-cost centers in the cost center.
	//
	// example:
	//
	// 2
	SubUnitCount *int64 `json:"SubUnitCount,omitempty" xml:"SubUnitCount,omitempty"`
	// The total number of resource instances, including resource instances of sub-cost centers, in the cost center.
	//
	// example:
	//
	// 3
	TotalResourceCount *int64 `json:"TotalResourceCount,omitempty" xml:"TotalResourceCount,omitempty"`
	// The total number of resource groups, including resource groups of sub-cost centers, in the cost center.
	//
	// example:
	//
	// 2
	TotalResourceGroupCount *int64 `json:"TotalResourceGroupCount,omitempty" xml:"TotalResourceGroupCount,omitempty"`
	// The total number of the associated accounts, including associated accounts of sub-cost centers, in the cost center.
	//
	// example:
	//
	// 2
	TotalUserCount *int64 `json:"TotalUserCount,omitempty" xml:"TotalUserCount,omitempty"`
	// The number of sub-cost centers in the cost center.
	//
	// example:
	//
	// 0
	UserCount *int64 `json:"UserCount,omitempty" xml:"UserCount,omitempty"`
}

func (s QueryCostUnitResourceResponseBodyDataCostUnitStatisInfo) String() string {
	return dara.Prettify(s)
}

func (s QueryCostUnitResourceResponseBodyDataCostUnitStatisInfo) GoString() string {
	return s.String()
}

func (s *QueryCostUnitResourceResponseBodyDataCostUnitStatisInfo) GetResourceCount() *int64 {
	return s.ResourceCount
}

func (s *QueryCostUnitResourceResponseBodyDataCostUnitStatisInfo) GetResourceGroupCount() *int64 {
	return s.ResourceGroupCount
}

func (s *QueryCostUnitResourceResponseBodyDataCostUnitStatisInfo) GetSubUnitCount() *int64 {
	return s.SubUnitCount
}

func (s *QueryCostUnitResourceResponseBodyDataCostUnitStatisInfo) GetTotalResourceCount() *int64 {
	return s.TotalResourceCount
}

func (s *QueryCostUnitResourceResponseBodyDataCostUnitStatisInfo) GetTotalResourceGroupCount() *int64 {
	return s.TotalResourceGroupCount
}

func (s *QueryCostUnitResourceResponseBodyDataCostUnitStatisInfo) GetTotalUserCount() *int64 {
	return s.TotalUserCount
}

func (s *QueryCostUnitResourceResponseBodyDataCostUnitStatisInfo) GetUserCount() *int64 {
	return s.UserCount
}

func (s *QueryCostUnitResourceResponseBodyDataCostUnitStatisInfo) SetResourceCount(v int64) *QueryCostUnitResourceResponseBodyDataCostUnitStatisInfo {
	s.ResourceCount = &v
	return s
}

func (s *QueryCostUnitResourceResponseBodyDataCostUnitStatisInfo) SetResourceGroupCount(v int64) *QueryCostUnitResourceResponseBodyDataCostUnitStatisInfo {
	s.ResourceGroupCount = &v
	return s
}

func (s *QueryCostUnitResourceResponseBodyDataCostUnitStatisInfo) SetSubUnitCount(v int64) *QueryCostUnitResourceResponseBodyDataCostUnitStatisInfo {
	s.SubUnitCount = &v
	return s
}

func (s *QueryCostUnitResourceResponseBodyDataCostUnitStatisInfo) SetTotalResourceCount(v int64) *QueryCostUnitResourceResponseBodyDataCostUnitStatisInfo {
	s.TotalResourceCount = &v
	return s
}

func (s *QueryCostUnitResourceResponseBodyDataCostUnitStatisInfo) SetTotalResourceGroupCount(v int64) *QueryCostUnitResourceResponseBodyDataCostUnitStatisInfo {
	s.TotalResourceGroupCount = &v
	return s
}

func (s *QueryCostUnitResourceResponseBodyDataCostUnitStatisInfo) SetTotalUserCount(v int64) *QueryCostUnitResourceResponseBodyDataCostUnitStatisInfo {
	s.TotalUserCount = &v
	return s
}

func (s *QueryCostUnitResourceResponseBodyDataCostUnitStatisInfo) SetUserCount(v int64) *QueryCostUnitResourceResponseBodyDataCostUnitStatisInfo {
	s.UserCount = &v
	return s
}

func (s *QueryCostUnitResourceResponseBodyDataCostUnitStatisInfo) Validate() error {
	return dara.Validate(s)
}

type QueryCostUnitResourceResponseBodyDataResourceInstanceDtoList struct {
	// The split code of the resource.
	//
	// example:
	//
	// test
	ApportionCode *string `json:"ApportionCode,omitempty" xml:"ApportionCode,omitempty"`
	// The split name of the resource.
	//
	// example:
	//
	// test
	ApportionName *string `json:"ApportionName,omitempty" xml:"ApportionName,omitempty"`
	// The product code of the resource.
	//
	// example:
	//
	// rds
	CommodityCode *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	// The commodity name of the resource.
	//
	// example:
	//
	// ApsaraDB
	CommodityName *string `json:"CommodityName,omitempty" xml:"CommodityName,omitempty"`
	// The code of the service. The code is the same as that in Cost Center.
	//
	// example:
	//
	// rds
	PipCode *string `json:"PipCode,omitempty" xml:"PipCode,omitempty"`
	// The resources related to the resource instance.
	//
	// example:
	//
	// oss
	RelatedResources *string `json:"RelatedResources,omitempty" xml:"RelatedResources,omitempty"`
	// The resource group to which the resource belongs.
	//
	// example:
	//
	// Default Resource Group
	ResourceGroup *string `json:"ResourceGroup,omitempty" xml:"ResourceGroup,omitempty"`
	// The instance ID of the resource.
	//
	// example:
	//
	// OSSBAG-cn-v0h1s4hma018
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The custom name of the resource.
	//
	// example:
	//
	// testResource
	ResourceNick *string `json:"ResourceNick,omitempty" xml:"ResourceNick,omitempty"`
	// The source of the resource. Value:
	//
	// - AUTO_ALLOCATE
	//
	// - MANUAL_ALLOCATE
	//
	// example:
	//
	// MANUAL_ALLOCATE
	ResourceSource *string `json:"ResourceSource,omitempty" xml:"ResourceSource,omitempty"`
	// The status of the resource.
	//
	// example:
	//
	// Available
	ResourceStatus *string `json:"ResourceStatus,omitempty" xml:"ResourceStatus,omitempty"`
	// The tags of the resource.
	//
	// example:
	//
	// testResource
	ResourceTag *string `json:"ResourceTag,omitempty" xml:"ResourceTag,omitempty"`
	// The type of the resource.
	//
	// example:
	//
	// FPT_ossbag_absolute_Storage_bj
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The user ID of the resource owner.
	//
	// example:
	//
	// 2424242134
	ResourceUserId *int64 `json:"ResourceUserId,omitempty" xml:"ResourceUserId,omitempty"`
	// The username of the resource owner.
	//
	// example:
	//
	// test@test.aliyun.com
	ResourceUserName *string `json:"ResourceUserName,omitempty" xml:"ResourceUserName,omitempty"`
}

func (s QueryCostUnitResourceResponseBodyDataResourceInstanceDtoList) String() string {
	return dara.Prettify(s)
}

func (s QueryCostUnitResourceResponseBodyDataResourceInstanceDtoList) GoString() string {
	return s.String()
}

func (s *QueryCostUnitResourceResponseBodyDataResourceInstanceDtoList) GetApportionCode() *string {
	return s.ApportionCode
}

func (s *QueryCostUnitResourceResponseBodyDataResourceInstanceDtoList) GetApportionName() *string {
	return s.ApportionName
}

func (s *QueryCostUnitResourceResponseBodyDataResourceInstanceDtoList) GetCommodityCode() *string {
	return s.CommodityCode
}

func (s *QueryCostUnitResourceResponseBodyDataResourceInstanceDtoList) GetCommodityName() *string {
	return s.CommodityName
}

func (s *QueryCostUnitResourceResponseBodyDataResourceInstanceDtoList) GetPipCode() *string {
	return s.PipCode
}

func (s *QueryCostUnitResourceResponseBodyDataResourceInstanceDtoList) GetRelatedResources() *string {
	return s.RelatedResources
}

func (s *QueryCostUnitResourceResponseBodyDataResourceInstanceDtoList) GetResourceGroup() *string {
	return s.ResourceGroup
}

func (s *QueryCostUnitResourceResponseBodyDataResourceInstanceDtoList) GetResourceId() *string {
	return s.ResourceId
}

func (s *QueryCostUnitResourceResponseBodyDataResourceInstanceDtoList) GetResourceNick() *string {
	return s.ResourceNick
}

func (s *QueryCostUnitResourceResponseBodyDataResourceInstanceDtoList) GetResourceSource() *string {
	return s.ResourceSource
}

func (s *QueryCostUnitResourceResponseBodyDataResourceInstanceDtoList) GetResourceStatus() *string {
	return s.ResourceStatus
}

func (s *QueryCostUnitResourceResponseBodyDataResourceInstanceDtoList) GetResourceTag() *string {
	return s.ResourceTag
}

func (s *QueryCostUnitResourceResponseBodyDataResourceInstanceDtoList) GetResourceType() *string {
	return s.ResourceType
}

func (s *QueryCostUnitResourceResponseBodyDataResourceInstanceDtoList) GetResourceUserId() *int64 {
	return s.ResourceUserId
}

func (s *QueryCostUnitResourceResponseBodyDataResourceInstanceDtoList) GetResourceUserName() *string {
	return s.ResourceUserName
}

func (s *QueryCostUnitResourceResponseBodyDataResourceInstanceDtoList) SetApportionCode(v string) *QueryCostUnitResourceResponseBodyDataResourceInstanceDtoList {
	s.ApportionCode = &v
	return s
}

func (s *QueryCostUnitResourceResponseBodyDataResourceInstanceDtoList) SetApportionName(v string) *QueryCostUnitResourceResponseBodyDataResourceInstanceDtoList {
	s.ApportionName = &v
	return s
}

func (s *QueryCostUnitResourceResponseBodyDataResourceInstanceDtoList) SetCommodityCode(v string) *QueryCostUnitResourceResponseBodyDataResourceInstanceDtoList {
	s.CommodityCode = &v
	return s
}

func (s *QueryCostUnitResourceResponseBodyDataResourceInstanceDtoList) SetCommodityName(v string) *QueryCostUnitResourceResponseBodyDataResourceInstanceDtoList {
	s.CommodityName = &v
	return s
}

func (s *QueryCostUnitResourceResponseBodyDataResourceInstanceDtoList) SetPipCode(v string) *QueryCostUnitResourceResponseBodyDataResourceInstanceDtoList {
	s.PipCode = &v
	return s
}

func (s *QueryCostUnitResourceResponseBodyDataResourceInstanceDtoList) SetRelatedResources(v string) *QueryCostUnitResourceResponseBodyDataResourceInstanceDtoList {
	s.RelatedResources = &v
	return s
}

func (s *QueryCostUnitResourceResponseBodyDataResourceInstanceDtoList) SetResourceGroup(v string) *QueryCostUnitResourceResponseBodyDataResourceInstanceDtoList {
	s.ResourceGroup = &v
	return s
}

func (s *QueryCostUnitResourceResponseBodyDataResourceInstanceDtoList) SetResourceId(v string) *QueryCostUnitResourceResponseBodyDataResourceInstanceDtoList {
	s.ResourceId = &v
	return s
}

func (s *QueryCostUnitResourceResponseBodyDataResourceInstanceDtoList) SetResourceNick(v string) *QueryCostUnitResourceResponseBodyDataResourceInstanceDtoList {
	s.ResourceNick = &v
	return s
}

func (s *QueryCostUnitResourceResponseBodyDataResourceInstanceDtoList) SetResourceSource(v string) *QueryCostUnitResourceResponseBodyDataResourceInstanceDtoList {
	s.ResourceSource = &v
	return s
}

func (s *QueryCostUnitResourceResponseBodyDataResourceInstanceDtoList) SetResourceStatus(v string) *QueryCostUnitResourceResponseBodyDataResourceInstanceDtoList {
	s.ResourceStatus = &v
	return s
}

func (s *QueryCostUnitResourceResponseBodyDataResourceInstanceDtoList) SetResourceTag(v string) *QueryCostUnitResourceResponseBodyDataResourceInstanceDtoList {
	s.ResourceTag = &v
	return s
}

func (s *QueryCostUnitResourceResponseBodyDataResourceInstanceDtoList) SetResourceType(v string) *QueryCostUnitResourceResponseBodyDataResourceInstanceDtoList {
	s.ResourceType = &v
	return s
}

func (s *QueryCostUnitResourceResponseBodyDataResourceInstanceDtoList) SetResourceUserId(v int64) *QueryCostUnitResourceResponseBodyDataResourceInstanceDtoList {
	s.ResourceUserId = &v
	return s
}

func (s *QueryCostUnitResourceResponseBodyDataResourceInstanceDtoList) SetResourceUserName(v string) *QueryCostUnitResourceResponseBodyDataResourceInstanceDtoList {
	s.ResourceUserName = &v
	return s
}

func (s *QueryCostUnitResourceResponseBodyDataResourceInstanceDtoList) Validate() error {
	return dara.Validate(s)
}
