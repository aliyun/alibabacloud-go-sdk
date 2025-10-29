// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetResourceGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetResourceGroupResponseBody
	GetRequestId() *string
	SetResourceGroup(v *GetResourceGroupResponseBodyResourceGroup) *GetResourceGroupResponseBody
	GetResourceGroup() *GetResourceGroupResponseBodyResourceGroup
	SetSuccess(v bool) *GetResourceGroupResponseBody
	GetSuccess() *bool
}

type GetResourceGroupResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 6A6CBE87-9F91-1323-B680-E7A7065XXXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The details about the resource group.
	ResourceGroup *GetResourceGroupResponseBodyResourceGroup `json:"ResourceGroup,omitempty" xml:"ResourceGroup,omitempty" type:"Struct"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetResourceGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetResourceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *GetResourceGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetResourceGroupResponseBody) GetResourceGroup() *GetResourceGroupResponseBodyResourceGroup {
	return s.ResourceGroup
}

func (s *GetResourceGroupResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetResourceGroupResponseBody) SetRequestId(v string) *GetResourceGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetResourceGroupResponseBody) SetResourceGroup(v *GetResourceGroupResponseBodyResourceGroup) *GetResourceGroupResponseBody {
	s.ResourceGroup = v
	return s
}

func (s *GetResourceGroupResponseBody) SetSuccess(v bool) *GetResourceGroupResponseBody {
	s.Success = &v
	return s
}

func (s *GetResourceGroupResponseBody) Validate() error {
	if s.ResourceGroup != nil {
		if err := s.ResourceGroup.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetResourceGroupResponseBodyResourceGroup struct {
	// The ID of the Alibaba Cloud resource group.
	//
	// example:
	//
	// rg-aek2kqofrgXXXXX
	AliyunResourceGroupId *string `json:"AliyunResourceGroupId,omitempty" xml:"AliyunResourceGroupId,omitempty"`
	// The tags.
	AliyunResourceTags []*GetResourceGroupResponseBodyResourceGroupAliyunResourceTags `json:"AliyunResourceTags,omitempty" xml:"AliyunResourceTags,omitempty" type:"Repeated"`
	// The time when the resource group was created. The value is a 64-bit timestamp.
	//
	// example:
	//
	// 1727055811000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The ID of the account that is used to create the resource group.
	//
	// example:
	//
	// 11075500042XXXXX
	CreateUser *string `json:"CreateUser,omitempty" xml:"CreateUser,omitempty"`
	// The ID of the virtual private cloud (VPC) with which the resource group is associated by default.
	//
	// example:
	//
	// vpc-m2et4f3oc8msfbccXXXXX
	DefaultVpcId *string `json:"DefaultVpcId,omitempty" xml:"DefaultVpcId,omitempty"`
	// The ID of the vSwitch with which the resource group is associated by default.
	//
	// example:
	//
	// vsw-uf8usrhs7hjd9amsXXXXX
	DefaultVswitchId *string `json:"DefaultVswitchId,omitempty" xml:"DefaultVswitchId,omitempty"`
	// The ID of the resource group.
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
	// The instance ID of the order that is used to create the resource group.
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
	// The description of the resource group.
	//
	// example:
	//
	// Create a common resource group for common tasks
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The type of the resource group. Valid values:
	//
	// 	- CommonV2: Serverless resource group.
	//
	// 	- ExclusiveDataIntegration: Exclusive resource group for Data Integration.
	//
	// 	- ExclusiveScheduler: Exclusive resource group for scheduling.
	//
	// 	- ExclusiveDataService: Exclusive resource group for DataService Studio.
	//
	// example:
	//
	// CommonV2
	ResourceGroupType *string `json:"ResourceGroupType,omitempty" xml:"ResourceGroupType,omitempty"`
	// The specifications of the resource group.
	Spec *GetResourceGroupResponseBodyResourceGroupSpec `json:"Spec,omitempty" xml:"Spec,omitempty" type:"Struct"`
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

func (s GetResourceGroupResponseBodyResourceGroup) String() string {
	return dara.Prettify(s)
}

func (s GetResourceGroupResponseBodyResourceGroup) GoString() string {
	return s.String()
}

func (s *GetResourceGroupResponseBodyResourceGroup) GetAliyunResourceGroupId() *string {
	return s.AliyunResourceGroupId
}

func (s *GetResourceGroupResponseBodyResourceGroup) GetAliyunResourceTags() []*GetResourceGroupResponseBodyResourceGroupAliyunResourceTags {
	return s.AliyunResourceTags
}

func (s *GetResourceGroupResponseBodyResourceGroup) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *GetResourceGroupResponseBodyResourceGroup) GetCreateUser() *string {
	return s.CreateUser
}

func (s *GetResourceGroupResponseBodyResourceGroup) GetDefaultVpcId() *string {
	return s.DefaultVpcId
}

func (s *GetResourceGroupResponseBodyResourceGroup) GetDefaultVswitchId() *string {
	return s.DefaultVswitchId
}

func (s *GetResourceGroupResponseBodyResourceGroup) GetId() *string {
	return s.Id
}

func (s *GetResourceGroupResponseBodyResourceGroup) GetName() *string {
	return s.Name
}

func (s *GetResourceGroupResponseBodyResourceGroup) GetOrderInstanceId() *string {
	return s.OrderInstanceId
}

func (s *GetResourceGroupResponseBodyResourceGroup) GetPaymentType() *string {
	return s.PaymentType
}

func (s *GetResourceGroupResponseBodyResourceGroup) GetRemark() *string {
	return s.Remark
}

func (s *GetResourceGroupResponseBodyResourceGroup) GetResourceGroupType() *string {
	return s.ResourceGroupType
}

func (s *GetResourceGroupResponseBodyResourceGroup) GetSpec() *GetResourceGroupResponseBodyResourceGroupSpec {
	return s.Spec
}

func (s *GetResourceGroupResponseBodyResourceGroup) GetStatus() *string {
	return s.Status
}

func (s *GetResourceGroupResponseBodyResourceGroup) SetAliyunResourceGroupId(v string) *GetResourceGroupResponseBodyResourceGroup {
	s.AliyunResourceGroupId = &v
	return s
}

func (s *GetResourceGroupResponseBodyResourceGroup) SetAliyunResourceTags(v []*GetResourceGroupResponseBodyResourceGroupAliyunResourceTags) *GetResourceGroupResponseBodyResourceGroup {
	s.AliyunResourceTags = v
	return s
}

func (s *GetResourceGroupResponseBodyResourceGroup) SetCreateTime(v int64) *GetResourceGroupResponseBodyResourceGroup {
	s.CreateTime = &v
	return s
}

func (s *GetResourceGroupResponseBodyResourceGroup) SetCreateUser(v string) *GetResourceGroupResponseBodyResourceGroup {
	s.CreateUser = &v
	return s
}

func (s *GetResourceGroupResponseBodyResourceGroup) SetDefaultVpcId(v string) *GetResourceGroupResponseBodyResourceGroup {
	s.DefaultVpcId = &v
	return s
}

func (s *GetResourceGroupResponseBodyResourceGroup) SetDefaultVswitchId(v string) *GetResourceGroupResponseBodyResourceGroup {
	s.DefaultVswitchId = &v
	return s
}

func (s *GetResourceGroupResponseBodyResourceGroup) SetId(v string) *GetResourceGroupResponseBodyResourceGroup {
	s.Id = &v
	return s
}

func (s *GetResourceGroupResponseBodyResourceGroup) SetName(v string) *GetResourceGroupResponseBodyResourceGroup {
	s.Name = &v
	return s
}

func (s *GetResourceGroupResponseBodyResourceGroup) SetOrderInstanceId(v string) *GetResourceGroupResponseBodyResourceGroup {
	s.OrderInstanceId = &v
	return s
}

func (s *GetResourceGroupResponseBodyResourceGroup) SetPaymentType(v string) *GetResourceGroupResponseBodyResourceGroup {
	s.PaymentType = &v
	return s
}

func (s *GetResourceGroupResponseBodyResourceGroup) SetRemark(v string) *GetResourceGroupResponseBodyResourceGroup {
	s.Remark = &v
	return s
}

func (s *GetResourceGroupResponseBodyResourceGroup) SetResourceGroupType(v string) *GetResourceGroupResponseBodyResourceGroup {
	s.ResourceGroupType = &v
	return s
}

func (s *GetResourceGroupResponseBodyResourceGroup) SetSpec(v *GetResourceGroupResponseBodyResourceGroupSpec) *GetResourceGroupResponseBodyResourceGroup {
	s.Spec = v
	return s
}

func (s *GetResourceGroupResponseBodyResourceGroup) SetStatus(v string) *GetResourceGroupResponseBodyResourceGroup {
	s.Status = &v
	return s
}

func (s *GetResourceGroupResponseBodyResourceGroup) Validate() error {
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

type GetResourceGroupResponseBodyResourceGroupAliyunResourceTags struct {
	// The tag key.
	//
	// example:
	//
	// key
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// value
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetResourceGroupResponseBodyResourceGroupAliyunResourceTags) String() string {
	return dara.Prettify(s)
}

func (s GetResourceGroupResponseBodyResourceGroupAliyunResourceTags) GoString() string {
	return s.String()
}

func (s *GetResourceGroupResponseBodyResourceGroupAliyunResourceTags) GetKey() *string {
	return s.Key
}

func (s *GetResourceGroupResponseBodyResourceGroupAliyunResourceTags) GetValue() *string {
	return s.Value
}

func (s *GetResourceGroupResponseBodyResourceGroupAliyunResourceTags) SetKey(v string) *GetResourceGroupResponseBodyResourceGroupAliyunResourceTags {
	s.Key = &v
	return s
}

func (s *GetResourceGroupResponseBodyResourceGroupAliyunResourceTags) SetValue(v string) *GetResourceGroupResponseBodyResourceGroupAliyunResourceTags {
	s.Value = &v
	return s
}

func (s *GetResourceGroupResponseBodyResourceGroupAliyunResourceTags) Validate() error {
	return dara.Validate(s)
}

type GetResourceGroupResponseBodyResourceGroupSpec struct {
	// The number of resources in the resource group.
	//
	// example:
	//
	// 1
	Amount *int32 `json:"Amount,omitempty" xml:"Amount,omitempty"`
	// The number of compute units (CUs) in the resource group.
	//
	// example:
	//
	// 2CU
	Standard *string `json:"Standard,omitempty" xml:"Standard,omitempty"`
}

func (s GetResourceGroupResponseBodyResourceGroupSpec) String() string {
	return dara.Prettify(s)
}

func (s GetResourceGroupResponseBodyResourceGroupSpec) GoString() string {
	return s.String()
}

func (s *GetResourceGroupResponseBodyResourceGroupSpec) GetAmount() *int32 {
	return s.Amount
}

func (s *GetResourceGroupResponseBodyResourceGroupSpec) GetStandard() *string {
	return s.Standard
}

func (s *GetResourceGroupResponseBodyResourceGroupSpec) SetAmount(v int32) *GetResourceGroupResponseBodyResourceGroupSpec {
	s.Amount = &v
	return s
}

func (s *GetResourceGroupResponseBodyResourceGroupSpec) SetStandard(v string) *GetResourceGroupResponseBodyResourceGroupSpec {
	s.Standard = &v
	return s
}

func (s *GetResourceGroupResponseBodyResourceGroupSpec) Validate() error {
	return dara.Validate(s)
}
