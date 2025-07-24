// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateServiceInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateServiceInstanceRequest
	GetClientToken() *string
	SetCommodity(v *CreateServiceInstanceRequestCommodity) *CreateServiceInstanceRequest
	GetCommodity() *CreateServiceInstanceRequestCommodity
	SetContactGroup(v string) *CreateServiceInstanceRequest
	GetContactGroup() *string
	SetDryRun(v bool) *CreateServiceInstanceRequest
	GetDryRun() *bool
	SetEnableInstanceOps(v bool) *CreateServiceInstanceRequest
	GetEnableInstanceOps() *bool
	SetEnableUserPrometheus(v bool) *CreateServiceInstanceRequest
	GetEnableUserPrometheus() *bool
	SetName(v string) *CreateServiceInstanceRequest
	GetName() *string
	SetOperationMetadata(v *CreateServiceInstanceRequestOperationMetadata) *CreateServiceInstanceRequest
	GetOperationMetadata() *CreateServiceInstanceRequestOperationMetadata
	SetParameters(v map[string]interface{}) *CreateServiceInstanceRequest
	GetParameters() map[string]interface{}
	SetRegionId(v string) *CreateServiceInstanceRequest
	GetRegionId() *string
	SetResourceAutoPay(v bool) *CreateServiceInstanceRequest
	GetResourceAutoPay() *bool
	SetResourceGroupId(v string) *CreateServiceInstanceRequest
	GetResourceGroupId() *string
	SetServiceId(v string) *CreateServiceInstanceRequest
	GetServiceId() *string
	SetServiceInstanceId(v string) *CreateServiceInstanceRequest
	GetServiceInstanceId() *string
	SetServiceVersion(v string) *CreateServiceInstanceRequest
	GetServiceVersion() *string
	SetSpecificationCode(v string) *CreateServiceInstanceRequest
	GetSpecificationCode() *string
	SetSpecificationName(v string) *CreateServiceInstanceRequest
	GetSpecificationName() *string
	SetTag(v []*CreateServiceInstanceRequestTag) *CreateServiceInstanceRequest
	GetTag() []*CreateServiceInstanceRequestTag
	SetTemplateName(v string) *CreateServiceInstanceRequest
	GetTemplateName() *string
	SetTrialType(v string) *CreateServiceInstanceRequest
	GetTrialType() *string
}

type CreateServiceInstanceRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The information about the order placed in Alibaba Cloud Marketplace. You do not need to specify this parameter if the service is not published in Alibaba Cloud Marketplace or uses the pay-as-you-go billing method.
	Commodity *CreateServiceInstanceRequestCommodity `json:"Commodity,omitempty" xml:"Commodity,omitempty" type:"Struct"`
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
	OperationMetadata *CreateServiceInstanceRequestOperationMetadata `json:"OperationMetadata,omitempty" xml:"OperationMetadata,omitempty" type:"Struct"`
	// The parameters that the customer specifies to deploy the service instance.
	//
	// >  If region information is required to create a service instance, you must specify the region ID in the value of Parameters.
	//
	// example:
	//
	// {"NodeCount": 3, "SystemDiskSize": 40, "InstancePassword": "******"}
	Parameters map[string]interface{} `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
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
	Tag []*CreateServiceInstanceRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
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

func (s CreateServiceInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateServiceInstanceRequest) GoString() string {
	return s.String()
}

func (s *CreateServiceInstanceRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateServiceInstanceRequest) GetCommodity() *CreateServiceInstanceRequestCommodity {
	return s.Commodity
}

func (s *CreateServiceInstanceRequest) GetContactGroup() *string {
	return s.ContactGroup
}

func (s *CreateServiceInstanceRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *CreateServiceInstanceRequest) GetEnableInstanceOps() *bool {
	return s.EnableInstanceOps
}

func (s *CreateServiceInstanceRequest) GetEnableUserPrometheus() *bool {
	return s.EnableUserPrometheus
}

func (s *CreateServiceInstanceRequest) GetName() *string {
	return s.Name
}

func (s *CreateServiceInstanceRequest) GetOperationMetadata() *CreateServiceInstanceRequestOperationMetadata {
	return s.OperationMetadata
}

func (s *CreateServiceInstanceRequest) GetParameters() map[string]interface{} {
	return s.Parameters
}

func (s *CreateServiceInstanceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateServiceInstanceRequest) GetResourceAutoPay() *bool {
	return s.ResourceAutoPay
}

func (s *CreateServiceInstanceRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateServiceInstanceRequest) GetServiceId() *string {
	return s.ServiceId
}

func (s *CreateServiceInstanceRequest) GetServiceInstanceId() *string {
	return s.ServiceInstanceId
}

func (s *CreateServiceInstanceRequest) GetServiceVersion() *string {
	return s.ServiceVersion
}

func (s *CreateServiceInstanceRequest) GetSpecificationCode() *string {
	return s.SpecificationCode
}

func (s *CreateServiceInstanceRequest) GetSpecificationName() *string {
	return s.SpecificationName
}

func (s *CreateServiceInstanceRequest) GetTag() []*CreateServiceInstanceRequestTag {
	return s.Tag
}

func (s *CreateServiceInstanceRequest) GetTemplateName() *string {
	return s.TemplateName
}

func (s *CreateServiceInstanceRequest) GetTrialType() *string {
	return s.TrialType
}

func (s *CreateServiceInstanceRequest) SetClientToken(v string) *CreateServiceInstanceRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateServiceInstanceRequest) SetCommodity(v *CreateServiceInstanceRequestCommodity) *CreateServiceInstanceRequest {
	s.Commodity = v
	return s
}

func (s *CreateServiceInstanceRequest) SetContactGroup(v string) *CreateServiceInstanceRequest {
	s.ContactGroup = &v
	return s
}

func (s *CreateServiceInstanceRequest) SetDryRun(v bool) *CreateServiceInstanceRequest {
	s.DryRun = &v
	return s
}

func (s *CreateServiceInstanceRequest) SetEnableInstanceOps(v bool) *CreateServiceInstanceRequest {
	s.EnableInstanceOps = &v
	return s
}

func (s *CreateServiceInstanceRequest) SetEnableUserPrometheus(v bool) *CreateServiceInstanceRequest {
	s.EnableUserPrometheus = &v
	return s
}

func (s *CreateServiceInstanceRequest) SetName(v string) *CreateServiceInstanceRequest {
	s.Name = &v
	return s
}

func (s *CreateServiceInstanceRequest) SetOperationMetadata(v *CreateServiceInstanceRequestOperationMetadata) *CreateServiceInstanceRequest {
	s.OperationMetadata = v
	return s
}

func (s *CreateServiceInstanceRequest) SetParameters(v map[string]interface{}) *CreateServiceInstanceRequest {
	s.Parameters = v
	return s
}

func (s *CreateServiceInstanceRequest) SetRegionId(v string) *CreateServiceInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *CreateServiceInstanceRequest) SetResourceAutoPay(v bool) *CreateServiceInstanceRequest {
	s.ResourceAutoPay = &v
	return s
}

func (s *CreateServiceInstanceRequest) SetResourceGroupId(v string) *CreateServiceInstanceRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateServiceInstanceRequest) SetServiceId(v string) *CreateServiceInstanceRequest {
	s.ServiceId = &v
	return s
}

func (s *CreateServiceInstanceRequest) SetServiceInstanceId(v string) *CreateServiceInstanceRequest {
	s.ServiceInstanceId = &v
	return s
}

func (s *CreateServiceInstanceRequest) SetServiceVersion(v string) *CreateServiceInstanceRequest {
	s.ServiceVersion = &v
	return s
}

func (s *CreateServiceInstanceRequest) SetSpecificationCode(v string) *CreateServiceInstanceRequest {
	s.SpecificationCode = &v
	return s
}

func (s *CreateServiceInstanceRequest) SetSpecificationName(v string) *CreateServiceInstanceRequest {
	s.SpecificationName = &v
	return s
}

func (s *CreateServiceInstanceRequest) SetTag(v []*CreateServiceInstanceRequestTag) *CreateServiceInstanceRequest {
	s.Tag = v
	return s
}

func (s *CreateServiceInstanceRequest) SetTemplateName(v string) *CreateServiceInstanceRequest {
	s.TemplateName = &v
	return s
}

func (s *CreateServiceInstanceRequest) SetTrialType(v string) *CreateServiceInstanceRequest {
	s.TrialType = &v
	return s
}

func (s *CreateServiceInstanceRequest) Validate() error {
	return dara.Validate(s)
}

type CreateServiceInstanceRequestCommodity struct {
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

func (s CreateServiceInstanceRequestCommodity) String() string {
	return dara.Prettify(s)
}

func (s CreateServiceInstanceRequestCommodity) GoString() string {
	return s.String()
}

func (s *CreateServiceInstanceRequestCommodity) GetAutoPay() *bool {
	return s.AutoPay
}

func (s *CreateServiceInstanceRequestCommodity) GetAutoRenew() *bool {
	return s.AutoRenew
}

func (s *CreateServiceInstanceRequestCommodity) GetCouponId() *string {
	return s.CouponId
}

func (s *CreateServiceInstanceRequestCommodity) GetPayPeriod() *int64 {
	return s.PayPeriod
}

func (s *CreateServiceInstanceRequestCommodity) GetPayPeriodUnit() *string {
	return s.PayPeriodUnit
}

func (s *CreateServiceInstanceRequestCommodity) GetQuotationId() *string {
	return s.QuotationId
}

func (s *CreateServiceInstanceRequestCommodity) SetAutoPay(v bool) *CreateServiceInstanceRequestCommodity {
	s.AutoPay = &v
	return s
}

func (s *CreateServiceInstanceRequestCommodity) SetAutoRenew(v bool) *CreateServiceInstanceRequestCommodity {
	s.AutoRenew = &v
	return s
}

func (s *CreateServiceInstanceRequestCommodity) SetCouponId(v string) *CreateServiceInstanceRequestCommodity {
	s.CouponId = &v
	return s
}

func (s *CreateServiceInstanceRequestCommodity) SetPayPeriod(v int64) *CreateServiceInstanceRequestCommodity {
	s.PayPeriod = &v
	return s
}

func (s *CreateServiceInstanceRequestCommodity) SetPayPeriodUnit(v string) *CreateServiceInstanceRequestCommodity {
	s.PayPeriodUnit = &v
	return s
}

func (s *CreateServiceInstanceRequestCommodity) SetQuotationId(v string) *CreateServiceInstanceRequestCommodity {
	s.QuotationId = &v
	return s
}

func (s *CreateServiceInstanceRequestCommodity) Validate() error {
	return dara.Validate(s)
}

type CreateServiceInstanceRequestOperationMetadata struct {
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

func (s CreateServiceInstanceRequestOperationMetadata) String() string {
	return dara.Prettify(s)
}

func (s CreateServiceInstanceRequestOperationMetadata) GoString() string {
	return s.String()
}

func (s *CreateServiceInstanceRequestOperationMetadata) GetEndTime() *string {
	return s.EndTime
}

func (s *CreateServiceInstanceRequestOperationMetadata) GetExtraInfo() *string {
	return s.ExtraInfo
}

func (s *CreateServiceInstanceRequestOperationMetadata) GetResources() *string {
	return s.Resources
}

func (s *CreateServiceInstanceRequestOperationMetadata) GetServiceInstanceId() *string {
	return s.ServiceInstanceId
}

func (s *CreateServiceInstanceRequestOperationMetadata) GetStartTime() *string {
	return s.StartTime
}

func (s *CreateServiceInstanceRequestOperationMetadata) SetEndTime(v string) *CreateServiceInstanceRequestOperationMetadata {
	s.EndTime = &v
	return s
}

func (s *CreateServiceInstanceRequestOperationMetadata) SetExtraInfo(v string) *CreateServiceInstanceRequestOperationMetadata {
	s.ExtraInfo = &v
	return s
}

func (s *CreateServiceInstanceRequestOperationMetadata) SetResources(v string) *CreateServiceInstanceRequestOperationMetadata {
	s.Resources = &v
	return s
}

func (s *CreateServiceInstanceRequestOperationMetadata) SetServiceInstanceId(v string) *CreateServiceInstanceRequestOperationMetadata {
	s.ServiceInstanceId = &v
	return s
}

func (s *CreateServiceInstanceRequestOperationMetadata) SetStartTime(v string) *CreateServiceInstanceRequestOperationMetadata {
	s.StartTime = &v
	return s
}

func (s *CreateServiceInstanceRequestOperationMetadata) Validate() error {
	return dara.Validate(s)
}

type CreateServiceInstanceRequestTag struct {
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

func (s CreateServiceInstanceRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateServiceInstanceRequestTag) GoString() string {
	return s.String()
}

func (s *CreateServiceInstanceRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateServiceInstanceRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateServiceInstanceRequestTag) SetKey(v string) *CreateServiceInstanceRequestTag {
	s.Key = &v
	return s
}

func (s *CreateServiceInstanceRequestTag) SetValue(v string) *CreateServiceInstanceRequestTag {
	s.Value = &v
	return s
}

func (s *CreateServiceInstanceRequestTag) Validate() error {
	return dara.Validate(s)
}
