// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListResourceGroupsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPagingInfo(v *ListResourceGroupsResponseBodyPagingInfo) *ListResourceGroupsResponseBody
	GetPagingInfo() *ListResourceGroupsResponseBodyPagingInfo
	SetRequestId(v string) *ListResourceGroupsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListResourceGroupsResponseBody
	GetSuccess() *bool
}

type ListResourceGroupsResponseBody struct {
	// The pagination information.
	PagingInfo *ListResourceGroupsResponseBodyPagingInfo `json:"PagingInfo,omitempty" xml:"PagingInfo,omitempty" type:"Struct"`
	// The request ID. You can use this ID to locate logs and troubleshoot issues.
	//
	// example:
	//
	// 6A6CBE87-9F91-1323-B680-E7A7065XXXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListResourceGroupsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListResourceGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *ListResourceGroupsResponseBody) GetPagingInfo() *ListResourceGroupsResponseBodyPagingInfo {
	return s.PagingInfo
}

func (s *ListResourceGroupsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListResourceGroupsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListResourceGroupsResponseBody) SetPagingInfo(v *ListResourceGroupsResponseBodyPagingInfo) *ListResourceGroupsResponseBody {
	s.PagingInfo = v
	return s
}

func (s *ListResourceGroupsResponseBody) SetRequestId(v string) *ListResourceGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListResourceGroupsResponseBody) SetSuccess(v bool) *ListResourceGroupsResponseBody {
	s.Success = &v
	return s
}

func (s *ListResourceGroupsResponseBody) Validate() error {
	if s.PagingInfo != nil {
		if err := s.PagingInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListResourceGroupsResponseBodyPagingInfo struct {
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 100
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The list of resource groups.
	ResourceGroupList []*ListResourceGroupsResponseBodyPagingInfoResourceGroupList `json:"ResourceGroupList,omitempty" xml:"ResourceGroupList,omitempty" type:"Repeated"`
	// The total number of entries.
	//
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListResourceGroupsResponseBodyPagingInfo) String() string {
	return dara.Prettify(s)
}

func (s ListResourceGroupsResponseBodyPagingInfo) GoString() string {
	return s.String()
}

func (s *ListResourceGroupsResponseBodyPagingInfo) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListResourceGroupsResponseBodyPagingInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListResourceGroupsResponseBodyPagingInfo) GetResourceGroupList() []*ListResourceGroupsResponseBodyPagingInfoResourceGroupList {
	return s.ResourceGroupList
}

func (s *ListResourceGroupsResponseBodyPagingInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListResourceGroupsResponseBodyPagingInfo) SetPageNumber(v int32) *ListResourceGroupsResponseBodyPagingInfo {
	s.PageNumber = &v
	return s
}

func (s *ListResourceGroupsResponseBodyPagingInfo) SetPageSize(v int32) *ListResourceGroupsResponseBodyPagingInfo {
	s.PageSize = &v
	return s
}

func (s *ListResourceGroupsResponseBodyPagingInfo) SetResourceGroupList(v []*ListResourceGroupsResponseBodyPagingInfoResourceGroupList) *ListResourceGroupsResponseBodyPagingInfo {
	s.ResourceGroupList = v
	return s
}

func (s *ListResourceGroupsResponseBodyPagingInfo) SetTotalCount(v int32) *ListResourceGroupsResponseBodyPagingInfo {
	s.TotalCount = &v
	return s
}

func (s *ListResourceGroupsResponseBodyPagingInfo) Validate() error {
	if s.ResourceGroupList != nil {
		for _, item := range s.ResourceGroupList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListResourceGroupsResponseBodyPagingInfoResourceGroupList struct {
	// The Alibaba Cloud resource group ID.
	//
	// example:
	//
	// rg-aek2kqofrgXXXXX
	AliyunResourceGroupId *string `json:"AliyunResourceGroupId,omitempty" xml:"AliyunResourceGroupId,omitempty"`
	// A list of Alibaba Cloud tags.
	AliyunResourceTags []*ListResourceGroupsResponseBodyPagingInfoResourceGroupListAliyunResourceTags `json:"AliyunResourceTags,omitempty" xml:"AliyunResourceTags,omitempty" type:"Repeated"`
	// The creation time of the resource group, as a 64-bit timestamp.
	//
	// example:
	//
	// 1727055811000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The user ID of the creator.
	//
	// example:
	//
	// 11075500042XXXXX
	CreateUser *string `json:"CreateUser,omitempty" xml:"CreateUser,omitempty"`
	// The ID of the default Virtual Private Cloud (VPC) bound to the general-purpose resource group.
	//
	// example:
	//
	// vpc-m2et4f3oc8msfbccXXXXX
	DefaultVpcId *string `json:"DefaultVpcId,omitempty" xml:"DefaultVpcId,omitempty"`
	// The ID of the default vSwitch bound to the general-purpose resource group.
	//
	// example:
	//
	// vsw-uf8usrhs7hjd9amsXXXXX
	DefaultVswicthId *string `json:"DefaultVswicthId,omitempty" xml:"DefaultVswicthId,omitempty"`
	// The unique identifier of the resource group.
	//
	// example:
	//
	// Serverless_res_group_524257424564736_6831777003XXXXX
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the resource group.
	//
	// example:
	//
	// common_resource_group
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the order for the resource group.
	//
	// example:
	//
	// c442b330-3b10-4584-959e-736e4edXXXXX
	OrderInstanceId *string `json:"OrderInstanceId,omitempty" xml:"OrderInstanceId,omitempty"`
	// The billing method of the resource group. `PrePaid` indicates subscription and `PostPaid` indicates pay-as-you-go.
	//
	// example:
	//
	// PrePaid
	PaymentType *string `json:"PaymentType,omitempty" xml:"PaymentType,omitempty"`
	// The description of the resource group.
	//
	// example:
	//
	// Create a general-purpose resource group for common tasks.
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The type of the resource group. Valid values:
	//
	// - `CommonV2`: The new general-purpose resource group.
	//
	// - `ExclusiveDataIntegration`: The exclusive resource group for data integration.
	//
	// - `ExclusiveScheduler`: The exclusive resource group for scheduling.
	//
	// - `ExclusiveDataService`: The exclusive resource group for data services.
	//
	// example:
	//
	// CommonV2
	ResourceGroupType *string `json:"ResourceGroupType,omitempty" xml:"ResourceGroupType,omitempty"`
	// The specifications of the resource group.
	Spec *ListResourceGroupsResponseBodyPagingInfoResourceGroupListSpec `json:"Spec,omitempty" xml:"Spec,omitempty" type:"Struct"`
	// The status of the resource group. Valid values:
	//
	// - `Normal`: Running.
	//
	// - `Stop`: Frozen due to expiration.
	//
	// - `Deleted`: Released.
	//
	// - `Creating`: Creation in progress.
	//
	// - `CreateFailed`: Creation failed.
	//
	// - `Updating`: Update in progress.
	//
	// - `UpdateFailed`: Update failed.
	//
	// - `Deleting`: Release in progress.
	//
	// - `DeleteFailed`: Release failed.
	//
	// - `Timeout`: The operation timed out.
	//
	// - `Freezed`: Frozen.
	//
	// - `Starting`: Starting.
	//
	// example:
	//
	// Normal
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListResourceGroupsResponseBodyPagingInfoResourceGroupList) String() string {
	return dara.Prettify(s)
}

func (s ListResourceGroupsResponseBodyPagingInfoResourceGroupList) GoString() string {
	return s.String()
}

func (s *ListResourceGroupsResponseBodyPagingInfoResourceGroupList) GetAliyunResourceGroupId() *string {
	return s.AliyunResourceGroupId
}

func (s *ListResourceGroupsResponseBodyPagingInfoResourceGroupList) GetAliyunResourceTags() []*ListResourceGroupsResponseBodyPagingInfoResourceGroupListAliyunResourceTags {
	return s.AliyunResourceTags
}

func (s *ListResourceGroupsResponseBodyPagingInfoResourceGroupList) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListResourceGroupsResponseBodyPagingInfoResourceGroupList) GetCreateUser() *string {
	return s.CreateUser
}

func (s *ListResourceGroupsResponseBodyPagingInfoResourceGroupList) GetDefaultVpcId() *string {
	return s.DefaultVpcId
}

func (s *ListResourceGroupsResponseBodyPagingInfoResourceGroupList) GetDefaultVswicthId() *string {
	return s.DefaultVswicthId
}

func (s *ListResourceGroupsResponseBodyPagingInfoResourceGroupList) GetId() *string {
	return s.Id
}

func (s *ListResourceGroupsResponseBodyPagingInfoResourceGroupList) GetName() *string {
	return s.Name
}

func (s *ListResourceGroupsResponseBodyPagingInfoResourceGroupList) GetOrderInstanceId() *string {
	return s.OrderInstanceId
}

func (s *ListResourceGroupsResponseBodyPagingInfoResourceGroupList) GetPaymentType() *string {
	return s.PaymentType
}

func (s *ListResourceGroupsResponseBodyPagingInfoResourceGroupList) GetRemark() *string {
	return s.Remark
}

func (s *ListResourceGroupsResponseBodyPagingInfoResourceGroupList) GetResourceGroupType() *string {
	return s.ResourceGroupType
}

func (s *ListResourceGroupsResponseBodyPagingInfoResourceGroupList) GetSpec() *ListResourceGroupsResponseBodyPagingInfoResourceGroupListSpec {
	return s.Spec
}

func (s *ListResourceGroupsResponseBodyPagingInfoResourceGroupList) GetStatus() *string {
	return s.Status
}

func (s *ListResourceGroupsResponseBodyPagingInfoResourceGroupList) SetAliyunResourceGroupId(v string) *ListResourceGroupsResponseBodyPagingInfoResourceGroupList {
	s.AliyunResourceGroupId = &v
	return s
}

func (s *ListResourceGroupsResponseBodyPagingInfoResourceGroupList) SetAliyunResourceTags(v []*ListResourceGroupsResponseBodyPagingInfoResourceGroupListAliyunResourceTags) *ListResourceGroupsResponseBodyPagingInfoResourceGroupList {
	s.AliyunResourceTags = v
	return s
}

func (s *ListResourceGroupsResponseBodyPagingInfoResourceGroupList) SetCreateTime(v int64) *ListResourceGroupsResponseBodyPagingInfoResourceGroupList {
	s.CreateTime = &v
	return s
}

func (s *ListResourceGroupsResponseBodyPagingInfoResourceGroupList) SetCreateUser(v string) *ListResourceGroupsResponseBodyPagingInfoResourceGroupList {
	s.CreateUser = &v
	return s
}

func (s *ListResourceGroupsResponseBodyPagingInfoResourceGroupList) SetDefaultVpcId(v string) *ListResourceGroupsResponseBodyPagingInfoResourceGroupList {
	s.DefaultVpcId = &v
	return s
}

func (s *ListResourceGroupsResponseBodyPagingInfoResourceGroupList) SetDefaultVswicthId(v string) *ListResourceGroupsResponseBodyPagingInfoResourceGroupList {
	s.DefaultVswicthId = &v
	return s
}

func (s *ListResourceGroupsResponseBodyPagingInfoResourceGroupList) SetId(v string) *ListResourceGroupsResponseBodyPagingInfoResourceGroupList {
	s.Id = &v
	return s
}

func (s *ListResourceGroupsResponseBodyPagingInfoResourceGroupList) SetName(v string) *ListResourceGroupsResponseBodyPagingInfoResourceGroupList {
	s.Name = &v
	return s
}

func (s *ListResourceGroupsResponseBodyPagingInfoResourceGroupList) SetOrderInstanceId(v string) *ListResourceGroupsResponseBodyPagingInfoResourceGroupList {
	s.OrderInstanceId = &v
	return s
}

func (s *ListResourceGroupsResponseBodyPagingInfoResourceGroupList) SetPaymentType(v string) *ListResourceGroupsResponseBodyPagingInfoResourceGroupList {
	s.PaymentType = &v
	return s
}

func (s *ListResourceGroupsResponseBodyPagingInfoResourceGroupList) SetRemark(v string) *ListResourceGroupsResponseBodyPagingInfoResourceGroupList {
	s.Remark = &v
	return s
}

func (s *ListResourceGroupsResponseBodyPagingInfoResourceGroupList) SetResourceGroupType(v string) *ListResourceGroupsResponseBodyPagingInfoResourceGroupList {
	s.ResourceGroupType = &v
	return s
}

func (s *ListResourceGroupsResponseBodyPagingInfoResourceGroupList) SetSpec(v *ListResourceGroupsResponseBodyPagingInfoResourceGroupListSpec) *ListResourceGroupsResponseBodyPagingInfoResourceGroupList {
	s.Spec = v
	return s
}

func (s *ListResourceGroupsResponseBodyPagingInfoResourceGroupList) SetStatus(v string) *ListResourceGroupsResponseBodyPagingInfoResourceGroupList {
	s.Status = &v
	return s
}

func (s *ListResourceGroupsResponseBodyPagingInfoResourceGroupList) Validate() error {
	if s.AliyunResourceTags != nil {
		for _, item := range s.AliyunResourceTags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Spec != nil {
		if err := s.Spec.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListResourceGroupsResponseBodyPagingInfoResourceGroupListAliyunResourceTags struct {
	// The tag key.
	//
	// example:
	//
	// Key
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// Value
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListResourceGroupsResponseBodyPagingInfoResourceGroupListAliyunResourceTags) String() string {
	return dara.Prettify(s)
}

func (s ListResourceGroupsResponseBodyPagingInfoResourceGroupListAliyunResourceTags) GoString() string {
	return s.String()
}

func (s *ListResourceGroupsResponseBodyPagingInfoResourceGroupListAliyunResourceTags) GetKey() *string {
	return s.Key
}

func (s *ListResourceGroupsResponseBodyPagingInfoResourceGroupListAliyunResourceTags) GetValue() *string {
	return s.Value
}

func (s *ListResourceGroupsResponseBodyPagingInfoResourceGroupListAliyunResourceTags) SetKey(v string) *ListResourceGroupsResponseBodyPagingInfoResourceGroupListAliyunResourceTags {
	s.Key = &v
	return s
}

func (s *ListResourceGroupsResponseBodyPagingInfoResourceGroupListAliyunResourceTags) SetValue(v string) *ListResourceGroupsResponseBodyPagingInfoResourceGroupListAliyunResourceTags {
	s.Value = &v
	return s
}

func (s *ListResourceGroupsResponseBodyPagingInfoResourceGroupListAliyunResourceTags) Validate() error {
	return dara.Validate(s)
}

type ListResourceGroupsResponseBodyPagingInfoResourceGroupListSpec struct {
	// The number of resource units.
	//
	// example:
	//
	// 1
	Amount *int32 `json:"Amount,omitempty" xml:"Amount,omitempty"`
	// The resource specifications.
	//
	// example:
	//
	// 2CU
	Standard *string `json:"Standard,omitempty" xml:"Standard,omitempty"`
}

func (s ListResourceGroupsResponseBodyPagingInfoResourceGroupListSpec) String() string {
	return dara.Prettify(s)
}

func (s ListResourceGroupsResponseBodyPagingInfoResourceGroupListSpec) GoString() string {
	return s.String()
}

func (s *ListResourceGroupsResponseBodyPagingInfoResourceGroupListSpec) GetAmount() *int32 {
	return s.Amount
}

func (s *ListResourceGroupsResponseBodyPagingInfoResourceGroupListSpec) GetStandard() *string {
	return s.Standard
}

func (s *ListResourceGroupsResponseBodyPagingInfoResourceGroupListSpec) SetAmount(v int32) *ListResourceGroupsResponseBodyPagingInfoResourceGroupListSpec {
	s.Amount = &v
	return s
}

func (s *ListResourceGroupsResponseBodyPagingInfoResourceGroupListSpec) SetStandard(v string) *ListResourceGroupsResponseBodyPagingInfoResourceGroupListSpec {
	s.Standard = &v
	return s
}

func (s *ListResourceGroupsResponseBodyPagingInfoResourceGroupListSpec) Validate() error {
	return dara.Validate(s)
}
