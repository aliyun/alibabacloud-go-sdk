// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateServiceInstanceShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateServiceInstanceShrinkRequest
	GetClientToken() *string
	SetCommodity(v *CreateServiceInstanceShrinkRequestCommodity) *CreateServiceInstanceShrinkRequest
	GetCommodity() *CreateServiceInstanceShrinkRequestCommodity
	SetContactGroup(v string) *CreateServiceInstanceShrinkRequest
	GetContactGroup() *string
	SetDryRun(v bool) *CreateServiceInstanceShrinkRequest
	GetDryRun() *bool
	SetEnableInstanceOps(v bool) *CreateServiceInstanceShrinkRequest
	GetEnableInstanceOps() *bool
	SetEnableUserPrometheus(v bool) *CreateServiceInstanceShrinkRequest
	GetEnableUserPrometheus() *bool
	SetName(v string) *CreateServiceInstanceShrinkRequest
	GetName() *string
	SetOperationMetadata(v *CreateServiceInstanceShrinkRequestOperationMetadata) *CreateServiceInstanceShrinkRequest
	GetOperationMetadata() *CreateServiceInstanceShrinkRequestOperationMetadata
	SetParametersShrink(v string) *CreateServiceInstanceShrinkRequest
	GetParametersShrink() *string
	SetRegionId(v string) *CreateServiceInstanceShrinkRequest
	GetRegionId() *string
	SetResourceAutoPay(v bool) *CreateServiceInstanceShrinkRequest
	GetResourceAutoPay() *bool
	SetResourceGroupId(v string) *CreateServiceInstanceShrinkRequest
	GetResourceGroupId() *string
	SetServiceId(v string) *CreateServiceInstanceShrinkRequest
	GetServiceId() *string
	SetServiceInstanceId(v string) *CreateServiceInstanceShrinkRequest
	GetServiceInstanceId() *string
	SetServiceVersion(v string) *CreateServiceInstanceShrinkRequest
	GetServiceVersion() *string
	SetSpecificationCode(v string) *CreateServiceInstanceShrinkRequest
	GetSpecificationCode() *string
	SetSpecificationName(v string) *CreateServiceInstanceShrinkRequest
	GetSpecificationName() *string
	SetTag(v []*CreateServiceInstanceShrinkRequestTag) *CreateServiceInstanceShrinkRequest
	GetTag() []*CreateServiceInstanceShrinkRequestTag
	SetTemplateName(v string) *CreateServiceInstanceShrinkRequest
	GetTemplateName() *string
	SetTrialType(v string) *CreateServiceInstanceShrinkRequest
	GetTrialType() *string
}

type CreateServiceInstanceShrinkRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The information about the order placed in Alibaba Cloud Marketplace. You do not need to specify this parameter if the service is not published in Alibaba Cloud Marketplace or uses the pay-as-you-go billing method.
	Commodity *CreateServiceInstanceShrinkRequestCommodity `json:"Commodity,omitempty" xml:"Commodity,omitempty" type:"Struct"`
	// The alert contact group.
	//
	// example:
	//
	// Default Group
	ContactGroup *string `json:"ContactGroup,omitempty" xml:"ContactGroup,omitempty"`
	// Specifies whether to perform only a dry run for the request to check information such as the permissions and instance status. Valid values:
	//
	// 	- **true**: performs a dry run for the request, but does not create a service instance.
	//
	// 	- **false**: performs a dry run for the request, and creates a service instance if the request passes the dry run.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// Specifies whether the service instance supports the hosted O\\&M feature. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	EnableInstanceOps *bool `json:"EnableInstanceOps,omitempty" xml:"EnableInstanceOps,omitempty"`
	// Specifies whether to enable the Prometheus monitoring feature. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	EnableUserPrometheus *bool `json:"EnableUserPrometheus,omitempty" xml:"EnableUserPrometheus,omitempty"`
	// The serviceInstance name.
	//
	// example:
	//
	// TestName
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The operation metadata.
	OperationMetadata *CreateServiceInstanceShrinkRequestOperationMetadata `json:"OperationMetadata,omitempty" xml:"OperationMetadata,omitempty" type:"Struct"`
	// The parameters that the customer specifies to deploy the service instance.
	//
	// >  If region information is required to create a service instance, you must specify the region ID in the value of Parameters.
	//
	// example:
	//
	// {"NodeCount": 3, "SystemDiskSize": 40, "InstancePassword": "******"}
	ParametersShrink *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// The region ID. Valid values:
	//
	// 	- cn-hangzhou: China (Hangzhou).
	//
	// 	- ap-southeast-1: Singapore.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Specifies whether to automatically deduct the resource fees from the account balance. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	ResourceAutoPay *bool `json:"ResourceAutoPay,omitempty" xml:"ResourceAutoPay,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-bp67acfmxazb4p****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The service ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// service-0e6fca6a51a54420****
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// The trial service instance id witch you want to convert to formal
	//
	// example:
	//
	// si-d32fbcef30664721b785
	ServiceInstanceId *string `json:"ServiceInstanceId,omitempty" xml:"ServiceInstanceId,omitempty"`
	// The service version.
	//
	// example:
	//
	// 1
	ServiceVersion *string `json:"ServiceVersion,omitempty" xml:"ServiceVersion,omitempty"`
	// Specification code.
	//
	// example:
	//
	// yuncode5425200001
	SpecificationCode *string `json:"SpecificationCode,omitempty" xml:"SpecificationCode,omitempty"`
	// The package name.
	//
	// example:
	//
	// Default Ppackage
	SpecificationName *string `json:"SpecificationName,omitempty" xml:"SpecificationName,omitempty"`
	// The tags.
	Tag []*CreateServiceInstanceShrinkRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The name of the template.
	//
	// example:
	//
	// ECS Template
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// The trial type of the service instance. Valid values:
	//
	// 	- **Trial**: Trials are supported.
	//
	// 	- **NotTrial**: Trials are not supported.
	//
	// example:
	//
	// Trial
	TrialType *string `json:"TrialType,omitempty" xml:"TrialType,omitempty"`
}

func (s CreateServiceInstanceShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateServiceInstanceShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateServiceInstanceShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateServiceInstanceShrinkRequest) GetCommodity() *CreateServiceInstanceShrinkRequestCommodity {
	return s.Commodity
}

func (s *CreateServiceInstanceShrinkRequest) GetContactGroup() *string {
	return s.ContactGroup
}

func (s *CreateServiceInstanceShrinkRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *CreateServiceInstanceShrinkRequest) GetEnableInstanceOps() *bool {
	return s.EnableInstanceOps
}

func (s *CreateServiceInstanceShrinkRequest) GetEnableUserPrometheus() *bool {
	return s.EnableUserPrometheus
}

func (s *CreateServiceInstanceShrinkRequest) GetName() *string {
	return s.Name
}

func (s *CreateServiceInstanceShrinkRequest) GetOperationMetadata() *CreateServiceInstanceShrinkRequestOperationMetadata {
	return s.OperationMetadata
}

func (s *CreateServiceInstanceShrinkRequest) GetParametersShrink() *string {
	return s.ParametersShrink
}

func (s *CreateServiceInstanceShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateServiceInstanceShrinkRequest) GetResourceAutoPay() *bool {
	return s.ResourceAutoPay
}

func (s *CreateServiceInstanceShrinkRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateServiceInstanceShrinkRequest) GetServiceId() *string {
	return s.ServiceId
}

func (s *CreateServiceInstanceShrinkRequest) GetServiceInstanceId() *string {
	return s.ServiceInstanceId
}

func (s *CreateServiceInstanceShrinkRequest) GetServiceVersion() *string {
	return s.ServiceVersion
}

func (s *CreateServiceInstanceShrinkRequest) GetSpecificationCode() *string {
	return s.SpecificationCode
}

func (s *CreateServiceInstanceShrinkRequest) GetSpecificationName() *string {
	return s.SpecificationName
}

func (s *CreateServiceInstanceShrinkRequest) GetTag() []*CreateServiceInstanceShrinkRequestTag {
	return s.Tag
}

func (s *CreateServiceInstanceShrinkRequest) GetTemplateName() *string {
	return s.TemplateName
}

func (s *CreateServiceInstanceShrinkRequest) GetTrialType() *string {
	return s.TrialType
}

func (s *CreateServiceInstanceShrinkRequest) SetClientToken(v string) *CreateServiceInstanceShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateServiceInstanceShrinkRequest) SetCommodity(v *CreateServiceInstanceShrinkRequestCommodity) *CreateServiceInstanceShrinkRequest {
	s.Commodity = v
	return s
}

func (s *CreateServiceInstanceShrinkRequest) SetContactGroup(v string) *CreateServiceInstanceShrinkRequest {
	s.ContactGroup = &v
	return s
}

func (s *CreateServiceInstanceShrinkRequest) SetDryRun(v bool) *CreateServiceInstanceShrinkRequest {
	s.DryRun = &v
	return s
}

func (s *CreateServiceInstanceShrinkRequest) SetEnableInstanceOps(v bool) *CreateServiceInstanceShrinkRequest {
	s.EnableInstanceOps = &v
	return s
}

func (s *CreateServiceInstanceShrinkRequest) SetEnableUserPrometheus(v bool) *CreateServiceInstanceShrinkRequest {
	s.EnableUserPrometheus = &v
	return s
}

func (s *CreateServiceInstanceShrinkRequest) SetName(v string) *CreateServiceInstanceShrinkRequest {
	s.Name = &v
	return s
}

func (s *CreateServiceInstanceShrinkRequest) SetOperationMetadata(v *CreateServiceInstanceShrinkRequestOperationMetadata) *CreateServiceInstanceShrinkRequest {
	s.OperationMetadata = v
	return s
}

func (s *CreateServiceInstanceShrinkRequest) SetParametersShrink(v string) *CreateServiceInstanceShrinkRequest {
	s.ParametersShrink = &v
	return s
}

func (s *CreateServiceInstanceShrinkRequest) SetRegionId(v string) *CreateServiceInstanceShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *CreateServiceInstanceShrinkRequest) SetResourceAutoPay(v bool) *CreateServiceInstanceShrinkRequest {
	s.ResourceAutoPay = &v
	return s
}

func (s *CreateServiceInstanceShrinkRequest) SetResourceGroupId(v string) *CreateServiceInstanceShrinkRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateServiceInstanceShrinkRequest) SetServiceId(v string) *CreateServiceInstanceShrinkRequest {
	s.ServiceId = &v
	return s
}

func (s *CreateServiceInstanceShrinkRequest) SetServiceInstanceId(v string) *CreateServiceInstanceShrinkRequest {
	s.ServiceInstanceId = &v
	return s
}

func (s *CreateServiceInstanceShrinkRequest) SetServiceVersion(v string) *CreateServiceInstanceShrinkRequest {
	s.ServiceVersion = &v
	return s
}

func (s *CreateServiceInstanceShrinkRequest) SetSpecificationCode(v string) *CreateServiceInstanceShrinkRequest {
	s.SpecificationCode = &v
	return s
}

func (s *CreateServiceInstanceShrinkRequest) SetSpecificationName(v string) *CreateServiceInstanceShrinkRequest {
	s.SpecificationName = &v
	return s
}

func (s *CreateServiceInstanceShrinkRequest) SetTag(v []*CreateServiceInstanceShrinkRequestTag) *CreateServiceInstanceShrinkRequest {
	s.Tag = v
	return s
}

func (s *CreateServiceInstanceShrinkRequest) SetTemplateName(v string) *CreateServiceInstanceShrinkRequest {
	s.TemplateName = &v
	return s
}

func (s *CreateServiceInstanceShrinkRequest) SetTrialType(v string) *CreateServiceInstanceShrinkRequest {
	s.TrialType = &v
	return s
}

func (s *CreateServiceInstanceShrinkRequest) Validate() error {
	return dara.Validate(s)
}

type CreateServiceInstanceShrinkRequestCommodity struct {
	// Specifies whether to automatically complete the payment. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// false
	AutoPay *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	// Specifies whether to enable auto-renewal for the service instance. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// false
	AutoRenew *bool `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	// The coupon ID.
	//
	// example:
	//
	// 302070970220
	CouponId *string `json:"CouponId,omitempty" xml:"CouponId,omitempty"`
	// The subscription duration.
	//
	// example:
	//
	// 1
	PayPeriod *int64 `json:"PayPeriod,omitempty" xml:"PayPeriod,omitempty"`
	// The unit of the subscription duration. Valid values:
	//
	// 	- **Year**
	//
	// 	- **Month**
	//
	// 	- **Day**
	//
	// example:
	//
	// Year
	PayPeriodUnit *string `json:"PayPeriodUnit,omitempty" xml:"PayPeriodUnit,omitempty"`
	QuotationId   *string `json:"QuotationId,omitempty" xml:"QuotationId,omitempty"`
}

func (s CreateServiceInstanceShrinkRequestCommodity) String() string {
	return dara.Prettify(s)
}

func (s CreateServiceInstanceShrinkRequestCommodity) GoString() string {
	return s.String()
}

func (s *CreateServiceInstanceShrinkRequestCommodity) GetAutoPay() *bool {
	return s.AutoPay
}

func (s *CreateServiceInstanceShrinkRequestCommodity) GetAutoRenew() *bool {
	return s.AutoRenew
}

func (s *CreateServiceInstanceShrinkRequestCommodity) GetCouponId() *string {
	return s.CouponId
}

func (s *CreateServiceInstanceShrinkRequestCommodity) GetPayPeriod() *int64 {
	return s.PayPeriod
}

func (s *CreateServiceInstanceShrinkRequestCommodity) GetPayPeriodUnit() *string {
	return s.PayPeriodUnit
}

func (s *CreateServiceInstanceShrinkRequestCommodity) GetQuotationId() *string {
	return s.QuotationId
}

func (s *CreateServiceInstanceShrinkRequestCommodity) SetAutoPay(v bool) *CreateServiceInstanceShrinkRequestCommodity {
	s.AutoPay = &v
	return s
}

func (s *CreateServiceInstanceShrinkRequestCommodity) SetAutoRenew(v bool) *CreateServiceInstanceShrinkRequestCommodity {
	s.AutoRenew = &v
	return s
}

func (s *CreateServiceInstanceShrinkRequestCommodity) SetCouponId(v string) *CreateServiceInstanceShrinkRequestCommodity {
	s.CouponId = &v
	return s
}

func (s *CreateServiceInstanceShrinkRequestCommodity) SetPayPeriod(v int64) *CreateServiceInstanceShrinkRequestCommodity {
	s.PayPeriod = &v
	return s
}

func (s *CreateServiceInstanceShrinkRequestCommodity) SetPayPeriodUnit(v string) *CreateServiceInstanceShrinkRequestCommodity {
	s.PayPeriodUnit = &v
	return s
}

func (s *CreateServiceInstanceShrinkRequestCommodity) SetQuotationId(v string) *CreateServiceInstanceShrinkRequestCommodity {
	s.QuotationId = &v
	return s
}

func (s *CreateServiceInstanceShrinkRequestCommodity) Validate() error {
	return dara.Validate(s)
}

type CreateServiceInstanceShrinkRequestOperationMetadata struct {
	// The operation end time.
	//
	// example:
	//
	// 2022-01-28T06:48:56Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The additional information.
	//
	// example:
	//
	// ```json
	//
	//   {
	//
	//     "vncInfo": [
	//
	//       {
	//
	//         "instanceId": "i-001",
	//
	//         "username": "admin",
	//
	//         "password": "******",
	//
	//         "vncPassword": "******"
	//
	//       }
	//
	//     ]
	//
	//   }
	//
	//   ```
	ExtraInfo *string `json:"ExtraInfo,omitempty" xml:"ExtraInfo,omitempty"`
	// Imported resource.
	//
	// example:
	//
	// {   "RegionId": "cn-hangzhou",   "Type": "ResourceIds",   "ResourceIds": {     "ALIYUN::ECS::INSTANCE": ["i-xxx", "i-yyy"],     "ALIYUN::RDS::INSTANCE": ["rm-xxx", "rm-yyy"],     "ALIYUN::VPC::VPC": ["vpc-xxx", "vpc-yyy"],     "ALIYUN::SLB::INSTANCE": ["lb-xxx", "lb-yyy"]   } }
	Resources *string `json:"Resources,omitempty" xml:"Resources,omitempty"`
	// The ID of the service instance.
	//
	// example:
	//
	// si-d6ab3a63ccbb4b17****
	ServiceInstanceId *string `json:"ServiceInstanceId,omitempty" xml:"ServiceInstanceId,omitempty"`
	// The operation start time.
	//
	// example:
	//
	// 2021-12-29T06:48:56Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s CreateServiceInstanceShrinkRequestOperationMetadata) String() string {
	return dara.Prettify(s)
}

func (s CreateServiceInstanceShrinkRequestOperationMetadata) GoString() string {
	return s.String()
}

func (s *CreateServiceInstanceShrinkRequestOperationMetadata) GetEndTime() *string {
	return s.EndTime
}

func (s *CreateServiceInstanceShrinkRequestOperationMetadata) GetExtraInfo() *string {
	return s.ExtraInfo
}

func (s *CreateServiceInstanceShrinkRequestOperationMetadata) GetResources() *string {
	return s.Resources
}

func (s *CreateServiceInstanceShrinkRequestOperationMetadata) GetServiceInstanceId() *string {
	return s.ServiceInstanceId
}

func (s *CreateServiceInstanceShrinkRequestOperationMetadata) GetStartTime() *string {
	return s.StartTime
}

func (s *CreateServiceInstanceShrinkRequestOperationMetadata) SetEndTime(v string) *CreateServiceInstanceShrinkRequestOperationMetadata {
	s.EndTime = &v
	return s
}

func (s *CreateServiceInstanceShrinkRequestOperationMetadata) SetExtraInfo(v string) *CreateServiceInstanceShrinkRequestOperationMetadata {
	s.ExtraInfo = &v
	return s
}

func (s *CreateServiceInstanceShrinkRequestOperationMetadata) SetResources(v string) *CreateServiceInstanceShrinkRequestOperationMetadata {
	s.Resources = &v
	return s
}

func (s *CreateServiceInstanceShrinkRequestOperationMetadata) SetServiceInstanceId(v string) *CreateServiceInstanceShrinkRequestOperationMetadata {
	s.ServiceInstanceId = &v
	return s
}

func (s *CreateServiceInstanceShrinkRequestOperationMetadata) SetStartTime(v string) *CreateServiceInstanceShrinkRequestOperationMetadata {
	s.StartTime = &v
	return s
}

func (s *CreateServiceInstanceShrinkRequestOperationMetadata) Validate() error {
	return dara.Validate(s)
}

type CreateServiceInstanceShrinkRequestTag struct {
	// The tag key.
	//
	// example:
	//
	// key1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// value1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateServiceInstanceShrinkRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateServiceInstanceShrinkRequestTag) GoString() string {
	return s.String()
}

func (s *CreateServiceInstanceShrinkRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateServiceInstanceShrinkRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateServiceInstanceShrinkRequestTag) SetKey(v string) *CreateServiceInstanceShrinkRequestTag {
	s.Key = &v
	return s
}

func (s *CreateServiceInstanceShrinkRequestTag) SetValue(v string) *CreateServiceInstanceShrinkRequestTag {
	s.Value = &v
	return s
}

func (s *CreateServiceInstanceShrinkRequestTag) Validate() error {
	return dara.Validate(s)
}
