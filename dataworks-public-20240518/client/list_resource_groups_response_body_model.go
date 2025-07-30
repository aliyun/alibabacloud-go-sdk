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
	// The request ID. You can use the ID to query logs and troubleshoot issues.
	//
	// example:
	//
	// 6A6CBE87-9F91-1323-B680-E7A7065XXXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values: true and false.
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
	return dara.Validate(s)
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
	// The resource groups returned.
	ResourceGroupList []*ListResourceGroupsResponseBodyPagingInfoResourceGroupList `json:"ResourceGroupList,omitempty" xml:"ResourceGroupList,omitempty" type:"Repeated"`
	// All data entries
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
	return dara.Validate(s)
}

type ListResourceGroupsResponseBodyPagingInfoResourceGroupList struct {
	// Alibaba Cloud Resource Group ID
	//
	// example:
	//
	// rg-aek2kqofrgXXXXX
	AliyunResourceGroupId *string `json:"AliyunResourceGroupId,omitempty" xml:"AliyunResourceGroupId,omitempty"`
	// Alibaba Cloud tag list
	AliyunResourceTags []*ListResourceGroupsResponseBodyPagingInfoResourceGroupListAliyunResourceTags `json:"AliyunResourceTags,omitempty" xml:"AliyunResourceTags,omitempty" type:"Repeated"`
	// The creation time, which is a 64-bit timestamp.
	//
	// example:
	//
	// 1727055811000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The ID of the user who created the resource group.
	//
	// example:
	//
	// 11075500042XXXXX
	CreateUser *string `json:"CreateUser,omitempty" xml:"CreateUser,omitempty"`
	// Default VPC ID bound to a common resource group
	//
	// example:
	//
	// vpc-m2et4f3oc8msfbccXXXXX
	DefaultVpcId *string `json:"DefaultVpcId,omitempty" xml:"DefaultVpcId,omitempty"`
	// The default switch ID bound to the common resource group.
	//
	// example:
	//
	// vsw-uf8usrhs7hjd9amsXXXXX
	DefaultVswicthId *string `json:"DefaultVswicthId,omitempty" xml:"DefaultVswicthId,omitempty"`
	// Unique identifier of a resource group
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
	// The order instance ID of the resource group.
	//
	// example:
	//
	// c442b330-3b10-4584-959e-736e4edXXXXX
	OrderInstanceId *string `json:"OrderInstanceId,omitempty" xml:"OrderInstanceId,omitempty"`
	// The billing method of the resource group. Valid values: PrePaid and PostPaid. The value PrePaid indicates the subscription billing method, and the value PostPaid indicates the pay-as-you-go billing method.
	//
	// example:
	//
	// PrePaid
	PaymentType *string `json:"PaymentType,omitempty" xml:"PaymentType,omitempty"`
	// Remarks for resource groups
	//
	// example:
	//
	// Create a common resource group for common tasks
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// Resource group types:
	//
	// 	- CommonV2: Serverless resource group
	//
	// 	- ExclusiveDataIntegration: Exclusive resource group for Data Integration
	//
	// 	- ExclusiveScheduler: Exclusive resource group for scheduling
	//
	// 	- ExclusiveDataService: Exclusive resource group for DataService Studio
	//
	// example:
	//
	// CommonV2
	ResourceGroupType *string `json:"ResourceGroupType,omitempty" xml:"ResourceGroupType,omitempty"`
	// Resource Group specifications
	Spec *ListResourceGroupsResponseBodyPagingInfoResourceGroupListSpec `json:"Spec,omitempty" xml:"Spec,omitempty" type:"Struct"`
	// The status of the resource group. Valid values:
	//
	// 	- Normal: The resource group is running or in use.
	//
	// 	- Stop: The resource group is expired.
	//
	// 	- Deleted: The resource group is released or destroyed.
	//
	// 	- Creating: The resource group is being created.
	//
	// 	- CreateFailed: The resource group fails to be created.
	//
	// 	- Updating: The resource group is being scaled in or out, or the configurations of the resource group are being changed.
	//
	// 	- UpdateFailed: The resource group fails to be scaled out or upgraded.
	//
	// 	- Deleting: The resource group is being released or destroyed.
	//
	// 	- DeleteFailed: The resource group fails to be released or destroyed.
	//
	// 	- Timeout: The operations that are performed on the resource group time out.
	//
	// 	- Freezed: The resource group is frozen.
	//
	// 	- Starting: The resource group is being started.
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
	return dara.Validate(s)
}

type ListResourceGroupsResponseBodyPagingInfoResourceGroupListAliyunResourceTags struct {
	// Tag Key
	//
	// example:
	//
	// Key
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// Tag Value
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
	// Quantity
	//
	// example:
	//
	// 1
	Amount *int32 `json:"Amount,omitempty" xml:"Amount,omitempty"`
	// Specification details
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
