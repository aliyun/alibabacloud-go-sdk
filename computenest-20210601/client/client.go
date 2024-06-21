// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type ChangeResourceGroupRequest struct {
	// example:
	//
	// rg-acfmzmhzo******
	NewResourceGroupId *string `json:"NewResourceGroupId,omitempty" xml:"NewResourceGroupId,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// si-5dc794a7fd254e******
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// example:
	//
	// service
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s ChangeResourceGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s ChangeResourceGroupRequest) GoString() string {
	return s.String()
}

func (s *ChangeResourceGroupRequest) SetNewResourceGroupId(v string) *ChangeResourceGroupRequest {
	s.NewResourceGroupId = &v
	return s
}

func (s *ChangeResourceGroupRequest) SetRegionId(v string) *ChangeResourceGroupRequest {
	s.RegionId = &v
	return s
}

func (s *ChangeResourceGroupRequest) SetResourceId(v string) *ChangeResourceGroupRequest {
	s.ResourceId = &v
	return s
}

func (s *ChangeResourceGroupRequest) SetResourceType(v string) *ChangeResourceGroupRequest {
	s.ResourceType = &v
	return s
}

type ChangeResourceGroupResponseBody struct {
	// example:
	//
	// 464C8CB6-A548-5206-B83C-D32A8E43EC21
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ChangeResourceGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ChangeResourceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ChangeResourceGroupResponseBody) SetRequestId(v string) *ChangeResourceGroupResponseBody {
	s.RequestId = &v
	return s
}

type ChangeResourceGroupResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ChangeResourceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ChangeResourceGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s ChangeResourceGroupResponse) GoString() string {
	return s.String()
}

func (s *ChangeResourceGroupResponse) SetHeaders(v map[string]*string) *ChangeResourceGroupResponse {
	s.Headers = v
	return s
}

func (s *ChangeResourceGroupResponse) SetStatusCode(v int32) *ChangeResourceGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *ChangeResourceGroupResponse) SetBody(v *ChangeResourceGroupResponseBody) *ChangeResourceGroupResponse {
	s.Body = v
	return s
}

type ContinueDeployServiceInstanceRequest struct {
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// example:
	//
	// false
	DryRun *bool     `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	Option []*string `json:"Option,omitempty" xml:"Option,omitempty" type:"Repeated"`
	// example:
	//
	// {"NodeCount": 3, "SystemDiskSize": 40, "InstancePassword": "******"}
	Parameters *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// si-0e6fca6a51a54420****
	ServiceInstanceId *string `json:"ServiceInstanceId,omitempty" xml:"ServiceInstanceId,omitempty"`
}

func (s ContinueDeployServiceInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s ContinueDeployServiceInstanceRequest) GoString() string {
	return s.String()
}

func (s *ContinueDeployServiceInstanceRequest) SetClientToken(v string) *ContinueDeployServiceInstanceRequest {
	s.ClientToken = &v
	return s
}

func (s *ContinueDeployServiceInstanceRequest) SetDryRun(v bool) *ContinueDeployServiceInstanceRequest {
	s.DryRun = &v
	return s
}

func (s *ContinueDeployServiceInstanceRequest) SetOption(v []*string) *ContinueDeployServiceInstanceRequest {
	s.Option = v
	return s
}

func (s *ContinueDeployServiceInstanceRequest) SetParameters(v string) *ContinueDeployServiceInstanceRequest {
	s.Parameters = &v
	return s
}

func (s *ContinueDeployServiceInstanceRequest) SetRegionId(v string) *ContinueDeployServiceInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *ContinueDeployServiceInstanceRequest) SetServiceInstanceId(v string) *ContinueDeployServiceInstanceRequest {
	s.ServiceInstanceId = &v
	return s
}

type ContinueDeployServiceInstanceResponseBody struct {
	DryRunResult *ContinueDeployServiceInstanceResponseBodyDryRunResult `json:"DryRunResult,omitempty" xml:"DryRunResult,omitempty" type:"Struct"`
	// example:
	//
	// 4DB0F536-B3BE-4F0D-BD29-E83FB56D550C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// si-d6ab3a63ccbb4b17****
	ServiceInstanceId *string `json:"ServiceInstanceId,omitempty" xml:"ServiceInstanceId,omitempty"`
}

func (s ContinueDeployServiceInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ContinueDeployServiceInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *ContinueDeployServiceInstanceResponseBody) SetDryRunResult(v *ContinueDeployServiceInstanceResponseBodyDryRunResult) *ContinueDeployServiceInstanceResponseBody {
	s.DryRunResult = v
	return s
}

func (s *ContinueDeployServiceInstanceResponseBody) SetRequestId(v string) *ContinueDeployServiceInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ContinueDeployServiceInstanceResponseBody) SetServiceInstanceId(v string) *ContinueDeployServiceInstanceResponseBody {
	s.ServiceInstanceId = &v
	return s
}

type ContinueDeployServiceInstanceResponseBodyDryRunResult struct {
	ParametersAllowedToBeModified              []*string `json:"ParametersAllowedToBeModified,omitempty" xml:"ParametersAllowedToBeModified,omitempty" type:"Repeated"`
	ParametersConditionallyAllowedToBeModified []*string `json:"ParametersConditionallyAllowedToBeModified,omitempty" xml:"ParametersConditionallyAllowedToBeModified,omitempty" type:"Repeated"`
	ParametersNotAllowedToBeModified           []*string `json:"ParametersNotAllowedToBeModified,omitempty" xml:"ParametersNotAllowedToBeModified,omitempty" type:"Repeated"`
}

func (s ContinueDeployServiceInstanceResponseBodyDryRunResult) String() string {
	return tea.Prettify(s)
}

func (s ContinueDeployServiceInstanceResponseBodyDryRunResult) GoString() string {
	return s.String()
}

func (s *ContinueDeployServiceInstanceResponseBodyDryRunResult) SetParametersAllowedToBeModified(v []*string) *ContinueDeployServiceInstanceResponseBodyDryRunResult {
	s.ParametersAllowedToBeModified = v
	return s
}

func (s *ContinueDeployServiceInstanceResponseBodyDryRunResult) SetParametersConditionallyAllowedToBeModified(v []*string) *ContinueDeployServiceInstanceResponseBodyDryRunResult {
	s.ParametersConditionallyAllowedToBeModified = v
	return s
}

func (s *ContinueDeployServiceInstanceResponseBodyDryRunResult) SetParametersNotAllowedToBeModified(v []*string) *ContinueDeployServiceInstanceResponseBodyDryRunResult {
	s.ParametersNotAllowedToBeModified = v
	return s
}

type ContinueDeployServiceInstanceResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ContinueDeployServiceInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ContinueDeployServiceInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s ContinueDeployServiceInstanceResponse) GoString() string {
	return s.String()
}

func (s *ContinueDeployServiceInstanceResponse) SetHeaders(v map[string]*string) *ContinueDeployServiceInstanceResponse {
	s.Headers = v
	return s
}

func (s *ContinueDeployServiceInstanceResponse) SetStatusCode(v int32) *ContinueDeployServiceInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *ContinueDeployServiceInstanceResponse) SetBody(v *ContinueDeployServiceInstanceResponseBody) *ContinueDeployServiceInstanceResponse {
	s.Body = v
	return s
}

type CreateServiceInstanceRequest struct {
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string                                `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Commodity   *CreateServiceInstanceRequestCommodity `json:"Commodity,omitempty" xml:"Commodity,omitempty" type:"Struct"`
	// 接收告警的云监控联系人组。
	//
	// example:
	//
	// 云账号报警联系人
	ContactGroup *string `json:"ContactGroup,omitempty" xml:"ContactGroup,omitempty"`
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// example:
	//
	// true
	EnableInstanceOps *bool `json:"EnableInstanceOps,omitempty" xml:"EnableInstanceOps,omitempty"`
	// example:
	//
	// true
	EnableUserPrometheus *bool `json:"EnableUserPrometheus,omitempty" xml:"EnableUserPrometheus,omitempty"`
	// 服务实例名称。格式要求如下：
	//
	// - 长度不超过64个字符。
	//
	// - 必须以数字或英文字母开头，可包含数字、英文字母、短划线（-）和下划线（_）。
	//
	// example:
	//
	// TestName
	Name              *string                                        `json:"Name,omitempty" xml:"Name,omitempty"`
	OperationMetadata *CreateServiceInstanceRequestOperationMetadata `json:"OperationMetadata,omitempty" xml:"OperationMetadata,omitempty" type:"Struct"`
	// example:
	//
	// {"NodeCount": 3, "SystemDiskSize": 40, "InstancePassword": "******"}
	Parameters map[string]interface{} `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// rg-bp67acfmxazb4p****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// service-0e6fca6a51a54420****
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// example:
	//
	// 1
	ServiceVersion *string `json:"ServiceVersion,omitempty" xml:"ServiceVersion,omitempty"`
	// example:
	//
	// yuncode5425200001
	SpecificationCode *string `json:"SpecificationCode,omitempty" xml:"SpecificationCode,omitempty"`
	// 套餐规格名称。
	//
	// example:
	//
	// 套餐一
	SpecificationName *string `json:"SpecificationName,omitempty" xml:"SpecificationName,omitempty"`
	// 用户自定义标签。
	Tag          []*CreateServiceInstanceRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	TemplateName *string                            `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// 使用类型。可选值：
	//
	// - Trial：支持试用。
	//
	// - NotTrial：不支持试用。
	//
	// example:
	//
	// Trial
	TrialType *string `json:"TrialType,omitempty" xml:"TrialType,omitempty"`
}

func (s CreateServiceInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceInstanceRequest) GoString() string {
	return s.String()
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

func (s *CreateServiceInstanceRequest) SetResourceGroupId(v string) *CreateServiceInstanceRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateServiceInstanceRequest) SetServiceId(v string) *CreateServiceInstanceRequest {
	s.ServiceId = &v
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

type CreateServiceInstanceRequestCommodity struct {
	AutoPay   *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	AutoRenew *bool `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	// example:
	//
	// 1
	PayPeriod *int64 `json:"PayPeriod,omitempty" xml:"PayPeriod,omitempty"`
	// example:
	//
	// Year
	PayPeriodUnit *string `json:"PayPeriodUnit,omitempty" xml:"PayPeriodUnit,omitempty"`
}

func (s CreateServiceInstanceRequestCommodity) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceInstanceRequestCommodity) GoString() string {
	return s.String()
}

func (s *CreateServiceInstanceRequestCommodity) SetAutoPay(v bool) *CreateServiceInstanceRequestCommodity {
	s.AutoPay = &v
	return s
}

func (s *CreateServiceInstanceRequestCommodity) SetAutoRenew(v bool) *CreateServiceInstanceRequestCommodity {
	s.AutoRenew = &v
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

type CreateServiceInstanceRequestOperationMetadata struct {
	// example:
	//
	// 2022-01-28T06:48:56Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
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
	// example:
	//
	// {   "RegionId": "cn-hangzhou",   "Type": "ResourceIds",   "ResourceIds": {     "ALIYUN::ECS::INSTANCE": ["i-xxx", "i-yyy"],     "ALIYUN::RDS::INSTANCE": ["rm-xxx", "rm-yyy"],     "ALIYUN::VPC::VPC": ["vpc-xxx", "vpc-yyy"],     "ALIYUN::SLB::INSTANCE": ["lb-xxx", "lb-yyy"]   } }
	Resources *string `json:"Resources,omitempty" xml:"Resources,omitempty"`
	// example:
	//
	// si-d6ab3a63ccbb4b17****
	ServiceInstanceId *string `json:"ServiceInstanceId,omitempty" xml:"ServiceInstanceId,omitempty"`
	// example:
	//
	// 2021-12-29T06:48:56Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s CreateServiceInstanceRequestOperationMetadata) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceInstanceRequestOperationMetadata) GoString() string {
	return s.String()
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

type CreateServiceInstanceRequestTag struct {
	// 标签键。
	//
	// example:
	//
	// key1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// 标签值。
	//
	// example:
	//
	// value1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateServiceInstanceRequestTag) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceInstanceRequestTag) GoString() string {
	return s.String()
}

func (s *CreateServiceInstanceRequestTag) SetKey(v string) *CreateServiceInstanceRequestTag {
	s.Key = &v
	return s
}

func (s *CreateServiceInstanceRequestTag) SetValue(v string) *CreateServiceInstanceRequestTag {
	s.Value = &v
	return s
}

type CreateServiceInstanceShrinkRequest struct {
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string                                      `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Commodity   *CreateServiceInstanceShrinkRequestCommodity `json:"Commodity,omitempty" xml:"Commodity,omitempty" type:"Struct"`
	// 接收告警的云监控联系人组。
	//
	// example:
	//
	// 云账号报警联系人
	ContactGroup *string `json:"ContactGroup,omitempty" xml:"ContactGroup,omitempty"`
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// example:
	//
	// true
	EnableInstanceOps *bool `json:"EnableInstanceOps,omitempty" xml:"EnableInstanceOps,omitempty"`
	// example:
	//
	// true
	EnableUserPrometheus *bool `json:"EnableUserPrometheus,omitempty" xml:"EnableUserPrometheus,omitempty"`
	// 服务实例名称。格式要求如下：
	//
	// - 长度不超过64个字符。
	//
	// - 必须以数字或英文字母开头，可包含数字、英文字母、短划线（-）和下划线（_）。
	//
	// example:
	//
	// TestName
	Name              *string                                              `json:"Name,omitempty" xml:"Name,omitempty"`
	OperationMetadata *CreateServiceInstanceShrinkRequestOperationMetadata `json:"OperationMetadata,omitempty" xml:"OperationMetadata,omitempty" type:"Struct"`
	// example:
	//
	// {"NodeCount": 3, "SystemDiskSize": 40, "InstancePassword": "******"}
	ParametersShrink *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// rg-bp67acfmxazb4p****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// service-0e6fca6a51a54420****
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// example:
	//
	// 1
	ServiceVersion *string `json:"ServiceVersion,omitempty" xml:"ServiceVersion,omitempty"`
	// example:
	//
	// yuncode5425200001
	SpecificationCode *string `json:"SpecificationCode,omitempty" xml:"SpecificationCode,omitempty"`
	// 套餐规格名称。
	//
	// example:
	//
	// 套餐一
	SpecificationName *string `json:"SpecificationName,omitempty" xml:"SpecificationName,omitempty"`
	// 用户自定义标签。
	Tag          []*CreateServiceInstanceShrinkRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	TemplateName *string                                  `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// 使用类型。可选值：
	//
	// - Trial：支持试用。
	//
	// - NotTrial：不支持试用。
	//
	// example:
	//
	// Trial
	TrialType *string `json:"TrialType,omitempty" xml:"TrialType,omitempty"`
}

func (s CreateServiceInstanceShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceInstanceShrinkRequest) GoString() string {
	return s.String()
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

func (s *CreateServiceInstanceShrinkRequest) SetResourceGroupId(v string) *CreateServiceInstanceShrinkRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateServiceInstanceShrinkRequest) SetServiceId(v string) *CreateServiceInstanceShrinkRequest {
	s.ServiceId = &v
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

type CreateServiceInstanceShrinkRequestCommodity struct {
	AutoPay   *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	AutoRenew *bool `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	// example:
	//
	// 1
	PayPeriod *int64 `json:"PayPeriod,omitempty" xml:"PayPeriod,omitempty"`
	// example:
	//
	// Year
	PayPeriodUnit *string `json:"PayPeriodUnit,omitempty" xml:"PayPeriodUnit,omitempty"`
}

func (s CreateServiceInstanceShrinkRequestCommodity) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceInstanceShrinkRequestCommodity) GoString() string {
	return s.String()
}

func (s *CreateServiceInstanceShrinkRequestCommodity) SetAutoPay(v bool) *CreateServiceInstanceShrinkRequestCommodity {
	s.AutoPay = &v
	return s
}

func (s *CreateServiceInstanceShrinkRequestCommodity) SetAutoRenew(v bool) *CreateServiceInstanceShrinkRequestCommodity {
	s.AutoRenew = &v
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

type CreateServiceInstanceShrinkRequestOperationMetadata struct {
	// example:
	//
	// 2022-01-28T06:48:56Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
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
	// example:
	//
	// {   "RegionId": "cn-hangzhou",   "Type": "ResourceIds",   "ResourceIds": {     "ALIYUN::ECS::INSTANCE": ["i-xxx", "i-yyy"],     "ALIYUN::RDS::INSTANCE": ["rm-xxx", "rm-yyy"],     "ALIYUN::VPC::VPC": ["vpc-xxx", "vpc-yyy"],     "ALIYUN::SLB::INSTANCE": ["lb-xxx", "lb-yyy"]   } }
	Resources *string `json:"Resources,omitempty" xml:"Resources,omitempty"`
	// example:
	//
	// si-d6ab3a63ccbb4b17****
	ServiceInstanceId *string `json:"ServiceInstanceId,omitempty" xml:"ServiceInstanceId,omitempty"`
	// example:
	//
	// 2021-12-29T06:48:56Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s CreateServiceInstanceShrinkRequestOperationMetadata) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceInstanceShrinkRequestOperationMetadata) GoString() string {
	return s.String()
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

type CreateServiceInstanceShrinkRequestTag struct {
	// 标签键。
	//
	// example:
	//
	// key1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// 标签值。
	//
	// example:
	//
	// value1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateServiceInstanceShrinkRequestTag) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceInstanceShrinkRequestTag) GoString() string {
	return s.String()
}

func (s *CreateServiceInstanceShrinkRequestTag) SetKey(v string) *CreateServiceInstanceShrinkRequestTag {
	s.Key = &v
	return s
}

func (s *CreateServiceInstanceShrinkRequestTag) SetValue(v string) *CreateServiceInstanceShrinkRequestTag {
	s.Value = &v
	return s
}

type CreateServiceInstanceResponseBody struct {
	// example:
	//
	// 786***45
	MarketInstanceId *string `json:"MarketInstanceId,omitempty" xml:"MarketInstanceId,omitempty"`
	OrderId          *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// example:
	//
	// 4DB0F536-B3BE-4F0D-BD29-E83FB56D****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// si-d6ab3a63ccbb4b17****
	ServiceInstanceId *string `json:"ServiceInstanceId,omitempty" xml:"ServiceInstanceId,omitempty"`
	// example:
	//
	// Created
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s CreateServiceInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateServiceInstanceResponseBody) SetMarketInstanceId(v string) *CreateServiceInstanceResponseBody {
	s.MarketInstanceId = &v
	return s
}

func (s *CreateServiceInstanceResponseBody) SetOrderId(v string) *CreateServiceInstanceResponseBody {
	s.OrderId = &v
	return s
}

func (s *CreateServiceInstanceResponseBody) SetRequestId(v string) *CreateServiceInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateServiceInstanceResponseBody) SetServiceInstanceId(v string) *CreateServiceInstanceResponseBody {
	s.ServiceInstanceId = &v
	return s
}

func (s *CreateServiceInstanceResponseBody) SetStatus(v string) *CreateServiceInstanceResponseBody {
	s.Status = &v
	return s
}

type CreateServiceInstanceResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateServiceInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateServiceInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceInstanceResponse) GoString() string {
	return s.String()
}

func (s *CreateServiceInstanceResponse) SetHeaders(v map[string]*string) *CreateServiceInstanceResponse {
	s.Headers = v
	return s
}

func (s *CreateServiceInstanceResponse) SetStatusCode(v int32) *CreateServiceInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateServiceInstanceResponse) SetBody(v *CreateServiceInstanceResponseBody) *CreateServiceInstanceResponse {
	s.Body = v
	return s
}

type DeleteServiceInstancesRequest struct {
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	ServiceInstanceId []*string `json:"ServiceInstanceId,omitempty" xml:"ServiceInstanceId,omitempty" type:"Repeated"`
}

func (s DeleteServiceInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteServiceInstancesRequest) GoString() string {
	return s.String()
}

func (s *DeleteServiceInstancesRequest) SetClientToken(v string) *DeleteServiceInstancesRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteServiceInstancesRequest) SetRegionId(v string) *DeleteServiceInstancesRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteServiceInstancesRequest) SetServiceInstanceId(v []*string) *DeleteServiceInstancesRequest {
	s.ServiceInstanceId = v
	return s
}

type DeleteServiceInstancesResponseBody struct {
	// example:
	//
	// 4DB0F536-B3BE-4F0D-BD29-E83FB56D550C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteServiceInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteServiceInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteServiceInstancesResponseBody) SetRequestId(v string) *DeleteServiceInstancesResponseBody {
	s.RequestId = &v
	return s
}

type DeleteServiceInstancesResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteServiceInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteServiceInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteServiceInstancesResponse) GoString() string {
	return s.String()
}

func (s *DeleteServiceInstancesResponse) SetHeaders(v map[string]*string) *DeleteServiceInstancesResponse {
	s.Headers = v
	return s
}

func (s *DeleteServiceInstancesResponse) SetStatusCode(v int32) *DeleteServiceInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteServiceInstancesResponse) SetBody(v *DeleteServiceInstancesResponseBody) *DeleteServiceInstancesResponse {
	s.Body = v
	return s
}

type GetServiceInstanceRequest struct {
	// example:
	//
	// 704***59
	MarketInstanceId *string `json:"MarketInstanceId,omitempty" xml:"MarketInstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// si-d6ab3a63ccbb4b17****
	ServiceInstanceId *string `json:"ServiceInstanceId,omitempty" xml:"ServiceInstanceId,omitempty"`
}

func (s GetServiceInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s GetServiceInstanceRequest) GoString() string {
	return s.String()
}

func (s *GetServiceInstanceRequest) SetMarketInstanceId(v string) *GetServiceInstanceRequest {
	s.MarketInstanceId = &v
	return s
}

func (s *GetServiceInstanceRequest) SetRegionId(v string) *GetServiceInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *GetServiceInstanceRequest) SetServiceInstanceId(v string) *GetServiceInstanceRequest {
	s.ServiceInstanceId = &v
	return s
}

type GetServiceInstanceResponseBody struct {
	BizStatus *string `json:"BizStatus,omitempty" xml:"BizStatus,omitempty"`
	// 云市场额外计费项。
	//
	// example:
	//
	// {"TiKVServerCount":"3","package_version":"yuncode5398300001","PDServerCount":"3","TiDBServerCount":"2"}
	Components *string `json:"Components,omitempty" xml:"Components,omitempty"`
	// example:
	//
	// 2021-05-20T00:00:00Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// true
	EnableInstanceOps *bool `json:"EnableInstanceOps,omitempty" xml:"EnableInstanceOps,omitempty"`
	// example:
	//
	// true
	EnableUserPrometheus *bool `json:"EnableUserPrometheus,omitempty" xml:"EnableUserPrometheus,omitempty"`
	// example:
	//
	// 2022-01-01T12:00:00
	EndTime             *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	GrafanaDashBoardUrl *string `json:"GrafanaDashBoardUrl,omitempty" xml:"GrafanaDashBoardUrl,omitempty"`
	// example:
	//
	// true
	IsOperated *bool `json:"IsOperated,omitempty" xml:"IsOperated,omitempty"`
	// example:
	//
	// 2022-01-01T12:00:00
	LicenseEndTime *string `json:"LicenseEndTime,omitempty" xml:"LicenseEndTime,omitempty"`
	// example:
	//
	// 704***59
	MarketInstanceId *string `json:"MarketInstanceId,omitempty" xml:"MarketInstanceId,omitempty"`
	// example:
	//
	// TestName
	Name          *string                                      `json:"Name,omitempty" xml:"Name,omitempty"`
	NetworkConfig *GetServiceInstanceResponseBodyNetworkConfig `json:"NetworkConfig,omitempty" xml:"NetworkConfig,omitempty" type:"Struct"`
	// example:
	//
	// si-d6ab3a63ccbb4b17****
	OperatedServiceInstanceId *string `json:"OperatedServiceInstanceId,omitempty" xml:"OperatedServiceInstanceId,omitempty"`
	// example:
	//
	// 2022-01-28T06:48:56Z
	OperationEndTime *string `json:"OperationEndTime,omitempty" xml:"OperationEndTime,omitempty"`
	// example:
	//
	// 2021-12-29T06:48:56Z
	OperationStartTime *string `json:"OperationStartTime,omitempty" xml:"OperationStartTime,omitempty"`
	// example:
	//
	// {"InstanceIds":["i-hp38ofxl0dsyfi7z****"]}
	Outputs *string `json:"Outputs,omitempty" xml:"Outputs,omitempty"`
	// example:
	//
	// {"param":"value"}
	Parameters *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// example:
	//
	// Subscription
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// 套餐名称。
	//
	// example:
	//
	// 套餐一
	PredefinedParameterName *string `json:"PredefinedParameterName,omitempty" xml:"PredefinedParameterName,omitempty"`
	// example:
	//
	// 90
	Progress *int64 `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// example:
	//
	// 4DB0F536-B3BE-4F0D-BD29-E83FB56D550C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// rg-bp67acfmxazb4p****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// example:
	//
	// [{"StackId": "stack-xxx"}]
	Resources *string                                `json:"Resources,omitempty" xml:"Resources,omitempty"`
	Service   *GetServiceInstanceResponseBodyService `json:"Service,omitempty" xml:"Service,omitempty" type:"Struct"`
	// example:
	//
	// si-d6ab3a63ccbb4b17****
	ServiceInstanceId *string `json:"ServiceInstanceId,omitempty" xml:"ServiceInstanceId,omitempty"`
	// example:
	//
	// private
	ServiceType *string `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
	// example:
	//
	// User
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// example:
	//
	// Deployed
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// deploy successfully
	StatusDetail *string `json:"StatusDetail,omitempty" xml:"StatusDetail,omitempty"`
	// example:
	//
	// 158927391332*****
	SupplierUid  *int64                                `json:"SupplierUid,omitempty" xml:"SupplierUid,omitempty"`
	Tags         []*GetServiceInstanceResponseBodyTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	TemplateName *string                               `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// example:
	//
	// 2021-05-20T00:00:00Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// example:
	//
	// 130920852836****
	UserId *int64 `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s GetServiceInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetServiceInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *GetServiceInstanceResponseBody) SetBizStatus(v string) *GetServiceInstanceResponseBody {
	s.BizStatus = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetComponents(v string) *GetServiceInstanceResponseBody {
	s.Components = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetCreateTime(v string) *GetServiceInstanceResponseBody {
	s.CreateTime = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetEnableInstanceOps(v bool) *GetServiceInstanceResponseBody {
	s.EnableInstanceOps = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetEnableUserPrometheus(v bool) *GetServiceInstanceResponseBody {
	s.EnableUserPrometheus = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetEndTime(v string) *GetServiceInstanceResponseBody {
	s.EndTime = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetGrafanaDashBoardUrl(v string) *GetServiceInstanceResponseBody {
	s.GrafanaDashBoardUrl = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetIsOperated(v bool) *GetServiceInstanceResponseBody {
	s.IsOperated = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetLicenseEndTime(v string) *GetServiceInstanceResponseBody {
	s.LicenseEndTime = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetMarketInstanceId(v string) *GetServiceInstanceResponseBody {
	s.MarketInstanceId = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetName(v string) *GetServiceInstanceResponseBody {
	s.Name = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetNetworkConfig(v *GetServiceInstanceResponseBodyNetworkConfig) *GetServiceInstanceResponseBody {
	s.NetworkConfig = v
	return s
}

func (s *GetServiceInstanceResponseBody) SetOperatedServiceInstanceId(v string) *GetServiceInstanceResponseBody {
	s.OperatedServiceInstanceId = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetOperationEndTime(v string) *GetServiceInstanceResponseBody {
	s.OperationEndTime = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetOperationStartTime(v string) *GetServiceInstanceResponseBody {
	s.OperationStartTime = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetOutputs(v string) *GetServiceInstanceResponseBody {
	s.Outputs = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetParameters(v string) *GetServiceInstanceResponseBody {
	s.Parameters = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetPayType(v string) *GetServiceInstanceResponseBody {
	s.PayType = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetPredefinedParameterName(v string) *GetServiceInstanceResponseBody {
	s.PredefinedParameterName = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetProgress(v int64) *GetServiceInstanceResponseBody {
	s.Progress = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetRequestId(v string) *GetServiceInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetResourceGroupId(v string) *GetServiceInstanceResponseBody {
	s.ResourceGroupId = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetResources(v string) *GetServiceInstanceResponseBody {
	s.Resources = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetService(v *GetServiceInstanceResponseBodyService) *GetServiceInstanceResponseBody {
	s.Service = v
	return s
}

func (s *GetServiceInstanceResponseBody) SetServiceInstanceId(v string) *GetServiceInstanceResponseBody {
	s.ServiceInstanceId = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetServiceType(v string) *GetServiceInstanceResponseBody {
	s.ServiceType = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetSource(v string) *GetServiceInstanceResponseBody {
	s.Source = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetStatus(v string) *GetServiceInstanceResponseBody {
	s.Status = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetStatusDetail(v string) *GetServiceInstanceResponseBody {
	s.StatusDetail = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetSupplierUid(v int64) *GetServiceInstanceResponseBody {
	s.SupplierUid = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetTags(v []*GetServiceInstanceResponseBodyTags) *GetServiceInstanceResponseBody {
	s.Tags = v
	return s
}

func (s *GetServiceInstanceResponseBody) SetTemplateName(v string) *GetServiceInstanceResponseBody {
	s.TemplateName = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetUpdateTime(v string) *GetServiceInstanceResponseBody {
	s.UpdateTime = &v
	return s
}

func (s *GetServiceInstanceResponseBody) SetUserId(v int64) *GetServiceInstanceResponseBody {
	s.UserId = &v
	return s
}

type GetServiceInstanceResponseBodyNetworkConfig struct {
	// example:
	//
	// ep-m5ei37240541816b****
	EndpointId            *string                                                             `json:"EndpointId,omitempty" xml:"EndpointId,omitempty"`
	PrivateVpcConnections []*GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnections `json:"PrivateVpcConnections,omitempty" xml:"PrivateVpcConnections,omitempty" type:"Repeated"`
	// example:
	//
	// cb7f214f80ac348d87daaeac1f35****
	PrivateZoneId                *string                                                                    `json:"PrivateZoneId,omitempty" xml:"PrivateZoneId,omitempty"`
	ReversePrivateVpcConnections []*GetServiceInstanceResponseBodyNetworkConfigReversePrivateVpcConnections `json:"ReversePrivateVpcConnections,omitempty" xml:"ReversePrivateVpcConnections,omitempty" type:"Repeated"`
}

func (s GetServiceInstanceResponseBodyNetworkConfig) String() string {
	return tea.Prettify(s)
}

func (s GetServiceInstanceResponseBodyNetworkConfig) GoString() string {
	return s.String()
}

func (s *GetServiceInstanceResponseBodyNetworkConfig) SetEndpointId(v string) *GetServiceInstanceResponseBodyNetworkConfig {
	s.EndpointId = &v
	return s
}

func (s *GetServiceInstanceResponseBodyNetworkConfig) SetPrivateVpcConnections(v []*GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnections) *GetServiceInstanceResponseBodyNetworkConfig {
	s.PrivateVpcConnections = v
	return s
}

func (s *GetServiceInstanceResponseBodyNetworkConfig) SetPrivateZoneId(v string) *GetServiceInstanceResponseBodyNetworkConfig {
	s.PrivateZoneId = &v
	return s
}

func (s *GetServiceInstanceResponseBodyNetworkConfig) SetReversePrivateVpcConnections(v []*GetServiceInstanceResponseBodyNetworkConfigReversePrivateVpcConnections) *GetServiceInstanceResponseBodyNetworkConfig {
	s.ReversePrivateVpcConnections = v
	return s
}

type GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnections struct {
	ConnectionConfigs []*GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnectionsConnectionConfigs `json:"ConnectionConfigs,omitempty" xml:"ConnectionConfigs,omitempty" type:"Repeated"`
	// example:
	//
	// ep-m5ei37240541816b****
	EndpointId *string `json:"EndpointId,omitempty" xml:"EndpointId,omitempty"`
	// example:
	//
	// cb7f214f80ac348d87daaeac1f35****
	PrivateZoneId *string `json:"PrivateZoneId,omitempty" xml:"PrivateZoneId,omitempty"`
	// example:
	//
	// test.computenest.aliyuncs.com
	PrivateZoneName *string `json:"PrivateZoneName,omitempty" xml:"PrivateZoneName,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnections) String() string {
	return tea.Prettify(s)
}

func (s GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnections) GoString() string {
	return s.String()
}

func (s *GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnections) SetConnectionConfigs(v []*GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnectionsConnectionConfigs) *GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnections {
	s.ConnectionConfigs = v
	return s
}

func (s *GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnections) SetEndpointId(v string) *GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnections {
	s.EndpointId = &v
	return s
}

func (s *GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnections) SetPrivateZoneId(v string) *GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnections {
	s.PrivateZoneId = &v
	return s
}

func (s *GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnections) SetPrivateZoneName(v string) *GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnections {
	s.PrivateZoneName = &v
	return s
}

func (s *GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnections) SetRegionId(v string) *GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnections {
	s.RegionId = &v
	return s
}

type GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnectionsConnectionConfigs struct {
	ConnectBandwidth *int32 `json:"ConnectBandwidth,omitempty" xml:"ConnectBandwidth,omitempty"`
	// 域名名称。
	//
	// example:
	//
	// ie-569a9be34f5534f6bc6559b5c1xxxxxx.service-51f80502802e48xxxxxx.cn-hangzhou.computenest.aliyuncs.com
	DomainName  *string   `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EndpointIps []*string `json:"EndpointIps,omitempty" xml:"EndpointIps,omitempty" type:"Repeated"`
	// example:
	//
	// Ready
	IngressEndpointStatus *string `json:"IngressEndpointStatus,omitempty" xml:"IngressEndpointStatus,omitempty"`
	// example:
	//
	// Ready
	NetworkServiceStatus *string `json:"NetworkServiceStatus,omitempty" xml:"NetworkServiceStatus,omitempty"`
	// example:
	//
	// cn-beijing
	RegionId       *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SecurityGroups []*string `json:"SecurityGroups,omitempty" xml:"SecurityGroups,omitempty" type:"Repeated"`
	VSwitches      []*string `json:"VSwitches,omitempty" xml:"VSwitches,omitempty" type:"Repeated"`
	// example:
	//
	// vpc-bp1vwnn14rqpyiczj****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnectionsConnectionConfigs) String() string {
	return tea.Prettify(s)
}

func (s GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnectionsConnectionConfigs) GoString() string {
	return s.String()
}

func (s *GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnectionsConnectionConfigs) SetConnectBandwidth(v int32) *GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnectionsConnectionConfigs {
	s.ConnectBandwidth = &v
	return s
}

func (s *GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnectionsConnectionConfigs) SetDomainName(v string) *GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnectionsConnectionConfigs {
	s.DomainName = &v
	return s
}

func (s *GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnectionsConnectionConfigs) SetEndpointIps(v []*string) *GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnectionsConnectionConfigs {
	s.EndpointIps = v
	return s
}

func (s *GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnectionsConnectionConfigs) SetIngressEndpointStatus(v string) *GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnectionsConnectionConfigs {
	s.IngressEndpointStatus = &v
	return s
}

func (s *GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnectionsConnectionConfigs) SetNetworkServiceStatus(v string) *GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnectionsConnectionConfigs {
	s.NetworkServiceStatus = &v
	return s
}

func (s *GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnectionsConnectionConfigs) SetRegionId(v string) *GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnectionsConnectionConfigs {
	s.RegionId = &v
	return s
}

func (s *GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnectionsConnectionConfigs) SetSecurityGroups(v []*string) *GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnectionsConnectionConfigs {
	s.SecurityGroups = v
	return s
}

func (s *GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnectionsConnectionConfigs) SetVSwitches(v []*string) *GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnectionsConnectionConfigs {
	s.VSwitches = v
	return s
}

func (s *GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnectionsConnectionConfigs) SetVpcId(v string) *GetServiceInstanceResponseBodyNetworkConfigPrivateVpcConnectionsConnectionConfigs {
	s.VpcId = &v
	return s
}

type GetServiceInstanceResponseBodyNetworkConfigReversePrivateVpcConnections struct {
	// example:
	//
	// ep-m5ei42370541816b****
	EndpointId *string `json:"EndpointId,omitempty" xml:"EndpointId,omitempty"`
}

func (s GetServiceInstanceResponseBodyNetworkConfigReversePrivateVpcConnections) String() string {
	return tea.Prettify(s)
}

func (s GetServiceInstanceResponseBodyNetworkConfigReversePrivateVpcConnections) GoString() string {
	return s.String()
}

func (s *GetServiceInstanceResponseBodyNetworkConfigReversePrivateVpcConnections) SetEndpointId(v string) *GetServiceInstanceResponseBodyNetworkConfigReversePrivateVpcConnections {
	s.EndpointId = &v
	return s
}

type GetServiceInstanceResponseBodyService struct {
	DeployMetadata *string `json:"DeployMetadata,omitempty" xml:"DeployMetadata,omitempty"`
	// example:
	//
	// ros
	DeployType *string `json:"DeployType,omitempty" xml:"DeployType,omitempty"`
	// example:
	//
	// 2021-05-21T00:00:00Z
	PublishTime *string `json:"PublishTime,omitempty" xml:"PublishTime,omitempty"`
	// example:
	//
	// http://example.com
	ServiceDocUrl *string `json:"ServiceDocUrl,omitempty" xml:"ServiceDocUrl,omitempty"`
	// example:
	//
	// service-9c8a3522528b4fe8****
	ServiceId    *string                                              `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	ServiceInfos []*GetServiceInstanceResponseBodyServiceServiceInfos `json:"ServiceInfos,omitempty" xml:"ServiceInfos,omitempty" type:"Repeated"`
	// example:
	//
	// https://service-info-private.oss-cn-hangzhou.aliyuncs.com
	ServiceProductUrl *string `json:"ServiceProductUrl,omitempty" xml:"ServiceProductUrl,omitempty"`
	// example:
	//
	// private
	ServiceType *string `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
	// example:
	//
	// Online
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
	SupplierName *string `json:"SupplierName,omitempty" xml:"SupplierName,omitempty"`
	// example:
	//
	// http://example.com
	SupplierUrl               *string   `json:"SupplierUrl,omitempty" xml:"SupplierUrl,omitempty"`
	UpgradableServiceVersions []*string `json:"UpgradableServiceVersions,omitempty" xml:"UpgradableServiceVersions,omitempty" type:"Repeated"`
	UpgradeMetadata           *string   `json:"UpgradeMetadata,omitempty" xml:"UpgradeMetadata,omitempty"`
	// example:
	//
	// 1
	Version     *string `json:"Version,omitempty" xml:"Version,omitempty"`
	VersionName *string `json:"VersionName,omitempty" xml:"VersionName,omitempty"`
}

func (s GetServiceInstanceResponseBodyService) String() string {
	return tea.Prettify(s)
}

func (s GetServiceInstanceResponseBodyService) GoString() string {
	return s.String()
}

func (s *GetServiceInstanceResponseBodyService) SetDeployMetadata(v string) *GetServiceInstanceResponseBodyService {
	s.DeployMetadata = &v
	return s
}

func (s *GetServiceInstanceResponseBodyService) SetDeployType(v string) *GetServiceInstanceResponseBodyService {
	s.DeployType = &v
	return s
}

func (s *GetServiceInstanceResponseBodyService) SetPublishTime(v string) *GetServiceInstanceResponseBodyService {
	s.PublishTime = &v
	return s
}

func (s *GetServiceInstanceResponseBodyService) SetServiceDocUrl(v string) *GetServiceInstanceResponseBodyService {
	s.ServiceDocUrl = &v
	return s
}

func (s *GetServiceInstanceResponseBodyService) SetServiceId(v string) *GetServiceInstanceResponseBodyService {
	s.ServiceId = &v
	return s
}

func (s *GetServiceInstanceResponseBodyService) SetServiceInfos(v []*GetServiceInstanceResponseBodyServiceServiceInfos) *GetServiceInstanceResponseBodyService {
	s.ServiceInfos = v
	return s
}

func (s *GetServiceInstanceResponseBodyService) SetServiceProductUrl(v string) *GetServiceInstanceResponseBodyService {
	s.ServiceProductUrl = &v
	return s
}

func (s *GetServiceInstanceResponseBodyService) SetServiceType(v string) *GetServiceInstanceResponseBodyService {
	s.ServiceType = &v
	return s
}

func (s *GetServiceInstanceResponseBodyService) SetStatus(v string) *GetServiceInstanceResponseBodyService {
	s.Status = &v
	return s
}

func (s *GetServiceInstanceResponseBodyService) SetSupplierName(v string) *GetServiceInstanceResponseBodyService {
	s.SupplierName = &v
	return s
}

func (s *GetServiceInstanceResponseBodyService) SetSupplierUrl(v string) *GetServiceInstanceResponseBodyService {
	s.SupplierUrl = &v
	return s
}

func (s *GetServiceInstanceResponseBodyService) SetUpgradableServiceVersions(v []*string) *GetServiceInstanceResponseBodyService {
	s.UpgradableServiceVersions = v
	return s
}

func (s *GetServiceInstanceResponseBodyService) SetUpgradeMetadata(v string) *GetServiceInstanceResponseBodyService {
	s.UpgradeMetadata = &v
	return s
}

func (s *GetServiceInstanceResponseBodyService) SetVersion(v string) *GetServiceInstanceResponseBodyService {
	s.Version = &v
	return s
}

func (s *GetServiceInstanceResponseBodyService) SetVersionName(v string) *GetServiceInstanceResponseBodyService {
	s.VersionName = &v
	return s
}

type GetServiceInstanceResponseBodyServiceServiceInfos struct {
	// example:
	//
	// https://example.com/service-image/c1c4a559-cc60-4af1-b976-98f356602462.png
	Image *string `json:"Image,omitempty" xml:"Image,omitempty"`
	// example:
	//
	// zh-CN
	Locale           *string `json:"Locale,omitempty" xml:"Locale,omitempty"`
	Name             *string `json:"Name,omitempty" xml:"Name,omitempty"`
	ShortDescription *string `json:"ShortDescription,omitempty" xml:"ShortDescription,omitempty"`
}

func (s GetServiceInstanceResponseBodyServiceServiceInfos) String() string {
	return tea.Prettify(s)
}

func (s GetServiceInstanceResponseBodyServiceServiceInfos) GoString() string {
	return s.String()
}

func (s *GetServiceInstanceResponseBodyServiceServiceInfos) SetImage(v string) *GetServiceInstanceResponseBodyServiceServiceInfos {
	s.Image = &v
	return s
}

func (s *GetServiceInstanceResponseBodyServiceServiceInfos) SetLocale(v string) *GetServiceInstanceResponseBodyServiceServiceInfos {
	s.Locale = &v
	return s
}

func (s *GetServiceInstanceResponseBodyServiceServiceInfos) SetName(v string) *GetServiceInstanceResponseBodyServiceServiceInfos {
	s.Name = &v
	return s
}

func (s *GetServiceInstanceResponseBodyServiceServiceInfos) SetShortDescription(v string) *GetServiceInstanceResponseBodyServiceServiceInfos {
	s.ShortDescription = &v
	return s
}

type GetServiceInstanceResponseBodyTags struct {
	// example:
	//
	// key1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// value1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetServiceInstanceResponseBodyTags) String() string {
	return tea.Prettify(s)
}

func (s GetServiceInstanceResponseBodyTags) GoString() string {
	return s.String()
}

func (s *GetServiceInstanceResponseBodyTags) SetKey(v string) *GetServiceInstanceResponseBodyTags {
	s.Key = &v
	return s
}

func (s *GetServiceInstanceResponseBodyTags) SetValue(v string) *GetServiceInstanceResponseBodyTags {
	s.Value = &v
	return s
}

type GetServiceInstanceResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetServiceInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetServiceInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s GetServiceInstanceResponse) GoString() string {
	return s.String()
}

func (s *GetServiceInstanceResponse) SetHeaders(v map[string]*string) *GetServiceInstanceResponse {
	s.Headers = v
	return s
}

func (s *GetServiceInstanceResponse) SetStatusCode(v int32) *GetServiceInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetServiceInstanceResponse) SetBody(v *GetServiceInstanceResponseBody) *GetServiceInstanceResponse {
	s.Body = v
	return s
}

type GetServiceTemplateParameterConstraintsRequest struct {
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	DeployRegionId *string `json:"DeployRegionId,omitempty" xml:"DeployRegionId,omitempty"`
	// example:
	//
	// true
	EnablePrivateVpcConnection *bool                                                      `json:"EnablePrivateVpcConnection,omitempty" xml:"EnablePrivateVpcConnection,omitempty"`
	Parameters                 []*GetServiceTemplateParameterConstraintsRequestParameters `json:"Parameters,omitempty" xml:"Parameters,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// service-731f788406024axxxxxx
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// example:
	//
	// si-461ee95f46ca46xxxxxx
	ServiceInstanceId *string `json:"ServiceInstanceId,omitempty" xml:"ServiceInstanceId,omitempty"`
	// example:
	//
	// 1
	ServiceVersion    *string `json:"ServiceVersion,omitempty" xml:"ServiceVersion,omitempty"`
	SpecificationName *string `json:"SpecificationName,omitempty" xml:"SpecificationName,omitempty"`
	// This parameter is required.
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// example:
	//
	// NotTrial
	TrialType *string `json:"TrialType,omitempty" xml:"TrialType,omitempty"`
}

func (s GetServiceTemplateParameterConstraintsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetServiceTemplateParameterConstraintsRequest) GoString() string {
	return s.String()
}

func (s *GetServiceTemplateParameterConstraintsRequest) SetClientToken(v string) *GetServiceTemplateParameterConstraintsRequest {
	s.ClientToken = &v
	return s
}

func (s *GetServiceTemplateParameterConstraintsRequest) SetDeployRegionId(v string) *GetServiceTemplateParameterConstraintsRequest {
	s.DeployRegionId = &v
	return s
}

func (s *GetServiceTemplateParameterConstraintsRequest) SetEnablePrivateVpcConnection(v bool) *GetServiceTemplateParameterConstraintsRequest {
	s.EnablePrivateVpcConnection = &v
	return s
}

func (s *GetServiceTemplateParameterConstraintsRequest) SetParameters(v []*GetServiceTemplateParameterConstraintsRequestParameters) *GetServiceTemplateParameterConstraintsRequest {
	s.Parameters = v
	return s
}

func (s *GetServiceTemplateParameterConstraintsRequest) SetRegionId(v string) *GetServiceTemplateParameterConstraintsRequest {
	s.RegionId = &v
	return s
}

func (s *GetServiceTemplateParameterConstraintsRequest) SetServiceId(v string) *GetServiceTemplateParameterConstraintsRequest {
	s.ServiceId = &v
	return s
}

func (s *GetServiceTemplateParameterConstraintsRequest) SetServiceInstanceId(v string) *GetServiceTemplateParameterConstraintsRequest {
	s.ServiceInstanceId = &v
	return s
}

func (s *GetServiceTemplateParameterConstraintsRequest) SetServiceVersion(v string) *GetServiceTemplateParameterConstraintsRequest {
	s.ServiceVersion = &v
	return s
}

func (s *GetServiceTemplateParameterConstraintsRequest) SetSpecificationName(v string) *GetServiceTemplateParameterConstraintsRequest {
	s.SpecificationName = &v
	return s
}

func (s *GetServiceTemplateParameterConstraintsRequest) SetTemplateName(v string) *GetServiceTemplateParameterConstraintsRequest {
	s.TemplateName = &v
	return s
}

func (s *GetServiceTemplateParameterConstraintsRequest) SetTrialType(v string) *GetServiceTemplateParameterConstraintsRequest {
	s.TrialType = &v
	return s
}

type GetServiceTemplateParameterConstraintsRequestParameters struct {
	// example:
	//
	// InstanceType
	ParameterKey *string `json:"ParameterKey,omitempty" xml:"ParameterKey,omitempty"`
	// example:
	//
	// cn-hangzhou-j
	ParameterValue *string `json:"ParameterValue,omitempty" xml:"ParameterValue,omitempty"`
}

func (s GetServiceTemplateParameterConstraintsRequestParameters) String() string {
	return tea.Prettify(s)
}

func (s GetServiceTemplateParameterConstraintsRequestParameters) GoString() string {
	return s.String()
}

func (s *GetServiceTemplateParameterConstraintsRequestParameters) SetParameterKey(v string) *GetServiceTemplateParameterConstraintsRequestParameters {
	s.ParameterKey = &v
	return s
}

func (s *GetServiceTemplateParameterConstraintsRequestParameters) SetParameterValue(v string) *GetServiceTemplateParameterConstraintsRequestParameters {
	s.ParameterValue = &v
	return s
}

type GetServiceTemplateParameterConstraintsResponseBody struct {
	FamilyConstraints    []*string                                                                 `json:"FamilyConstraints,omitempty" xml:"FamilyConstraints,omitempty" type:"Repeated"`
	ParameterConstraints []*GetServiceTemplateParameterConstraintsResponseBodyParameterConstraints `json:"ParameterConstraints,omitempty" xml:"ParameterConstraints,omitempty" type:"Repeated"`
	// example:
	//
	// 449DC03D-A039-56A6-8D6F-6EBCCCE0EE2C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetServiceTemplateParameterConstraintsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetServiceTemplateParameterConstraintsResponseBody) GoString() string {
	return s.String()
}

func (s *GetServiceTemplateParameterConstraintsResponseBody) SetFamilyConstraints(v []*string) *GetServiceTemplateParameterConstraintsResponseBody {
	s.FamilyConstraints = v
	return s
}

func (s *GetServiceTemplateParameterConstraintsResponseBody) SetParameterConstraints(v []*GetServiceTemplateParameterConstraintsResponseBodyParameterConstraints) *GetServiceTemplateParameterConstraintsResponseBody {
	s.ParameterConstraints = v
	return s
}

func (s *GetServiceTemplateParameterConstraintsResponseBody) SetRequestId(v string) *GetServiceTemplateParameterConstraintsResponseBody {
	s.RequestId = &v
	return s
}

type GetServiceTemplateParameterConstraintsResponseBodyParameterConstraints struct {
	AllowedValues             []*string `json:"AllowedValues,omitempty" xml:"AllowedValues,omitempty" type:"Repeated"`
	AssociationParameterNames []*string `json:"AssociationParameterNames,omitempty" xml:"AssociationParameterNames,omitempty" type:"Repeated"`
	// example:
	//
	// NoLimit
	Behavior *string `json:"Behavior,omitempty" xml:"Behavior,omitempty"`
	// example:
	//
	// No resource property refer to the parameter
	BehaviorReason      *string                                                                                      `json:"BehaviorReason,omitempty" xml:"BehaviorReason,omitempty"`
	OriginalConstraints []*GetServiceTemplateParameterConstraintsResponseBodyParameterConstraintsOriginalConstraints `json:"OriginalConstraints,omitempty" xml:"OriginalConstraints,omitempty" type:"Repeated"`
	// example:
	//
	// ZoneInfo
	ParameterKey *string                                                                              `json:"ParameterKey,omitempty" xml:"ParameterKey,omitempty"`
	QueryErrors  []*GetServiceTemplateParameterConstraintsResponseBodyParameterConstraintsQueryErrors `json:"QueryErrors,omitempty" xml:"QueryErrors,omitempty" type:"Repeated"`
	// example:
	//
	// String
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetServiceTemplateParameterConstraintsResponseBodyParameterConstraints) String() string {
	return tea.Prettify(s)
}

func (s GetServiceTemplateParameterConstraintsResponseBodyParameterConstraints) GoString() string {
	return s.String()
}

func (s *GetServiceTemplateParameterConstraintsResponseBodyParameterConstraints) SetAllowedValues(v []*string) *GetServiceTemplateParameterConstraintsResponseBodyParameterConstraints {
	s.AllowedValues = v
	return s
}

func (s *GetServiceTemplateParameterConstraintsResponseBodyParameterConstraints) SetAssociationParameterNames(v []*string) *GetServiceTemplateParameterConstraintsResponseBodyParameterConstraints {
	s.AssociationParameterNames = v
	return s
}

func (s *GetServiceTemplateParameterConstraintsResponseBodyParameterConstraints) SetBehavior(v string) *GetServiceTemplateParameterConstraintsResponseBodyParameterConstraints {
	s.Behavior = &v
	return s
}

func (s *GetServiceTemplateParameterConstraintsResponseBodyParameterConstraints) SetBehaviorReason(v string) *GetServiceTemplateParameterConstraintsResponseBodyParameterConstraints {
	s.BehaviorReason = &v
	return s
}

func (s *GetServiceTemplateParameterConstraintsResponseBodyParameterConstraints) SetOriginalConstraints(v []*GetServiceTemplateParameterConstraintsResponseBodyParameterConstraintsOriginalConstraints) *GetServiceTemplateParameterConstraintsResponseBodyParameterConstraints {
	s.OriginalConstraints = v
	return s
}

func (s *GetServiceTemplateParameterConstraintsResponseBodyParameterConstraints) SetParameterKey(v string) *GetServiceTemplateParameterConstraintsResponseBodyParameterConstraints {
	s.ParameterKey = &v
	return s
}

func (s *GetServiceTemplateParameterConstraintsResponseBodyParameterConstraints) SetQueryErrors(v []*GetServiceTemplateParameterConstraintsResponseBodyParameterConstraintsQueryErrors) *GetServiceTemplateParameterConstraintsResponseBodyParameterConstraints {
	s.QueryErrors = v
	return s
}

func (s *GetServiceTemplateParameterConstraintsResponseBodyParameterConstraints) SetType(v string) *GetServiceTemplateParameterConstraintsResponseBodyParameterConstraints {
	s.Type = &v
	return s
}

type GetServiceTemplateParameterConstraintsResponseBodyParameterConstraintsOriginalConstraints struct {
	AllowedValues []*string `json:"AllowedValues,omitempty" xml:"AllowedValues,omitempty" type:"Repeated"`
	// example:
	//
	// ZoneId
	PropertyName *string `json:"PropertyName,omitempty" xml:"PropertyName,omitempty"`
	// example:
	//
	// MyECS
	ResourceName *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	// example:
	//
	// ALIYUN::ECS::InstanceGroup
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s GetServiceTemplateParameterConstraintsResponseBodyParameterConstraintsOriginalConstraints) String() string {
	return tea.Prettify(s)
}

func (s GetServiceTemplateParameterConstraintsResponseBodyParameterConstraintsOriginalConstraints) GoString() string {
	return s.String()
}

func (s *GetServiceTemplateParameterConstraintsResponseBodyParameterConstraintsOriginalConstraints) SetAllowedValues(v []*string) *GetServiceTemplateParameterConstraintsResponseBodyParameterConstraintsOriginalConstraints {
	s.AllowedValues = v
	return s
}

func (s *GetServiceTemplateParameterConstraintsResponseBodyParameterConstraintsOriginalConstraints) SetPropertyName(v string) *GetServiceTemplateParameterConstraintsResponseBodyParameterConstraintsOriginalConstraints {
	s.PropertyName = &v
	return s
}

func (s *GetServiceTemplateParameterConstraintsResponseBodyParameterConstraintsOriginalConstraints) SetResourceName(v string) *GetServiceTemplateParameterConstraintsResponseBodyParameterConstraintsOriginalConstraints {
	s.ResourceName = &v
	return s
}

func (s *GetServiceTemplateParameterConstraintsResponseBodyParameterConstraintsOriginalConstraints) SetResourceType(v string) *GetServiceTemplateParameterConstraintsResponseBodyParameterConstraintsOriginalConstraints {
	s.ResourceType = &v
	return s
}

type GetServiceTemplateParameterConstraintsResponseBodyParameterConstraintsQueryErrors struct {
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	ResourceName *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s GetServiceTemplateParameterConstraintsResponseBodyParameterConstraintsQueryErrors) String() string {
	return tea.Prettify(s)
}

func (s GetServiceTemplateParameterConstraintsResponseBodyParameterConstraintsQueryErrors) GoString() string {
	return s.String()
}

func (s *GetServiceTemplateParameterConstraintsResponseBodyParameterConstraintsQueryErrors) SetErrorMessage(v string) *GetServiceTemplateParameterConstraintsResponseBodyParameterConstraintsQueryErrors {
	s.ErrorMessage = &v
	return s
}

func (s *GetServiceTemplateParameterConstraintsResponseBodyParameterConstraintsQueryErrors) SetResourceName(v string) *GetServiceTemplateParameterConstraintsResponseBodyParameterConstraintsQueryErrors {
	s.ResourceName = &v
	return s
}

func (s *GetServiceTemplateParameterConstraintsResponseBodyParameterConstraintsQueryErrors) SetResourceType(v string) *GetServiceTemplateParameterConstraintsResponseBodyParameterConstraintsQueryErrors {
	s.ResourceType = &v
	return s
}

type GetServiceTemplateParameterConstraintsResponse struct {
	Headers    map[string]*string                                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetServiceTemplateParameterConstraintsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetServiceTemplateParameterConstraintsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetServiceTemplateParameterConstraintsResponse) GoString() string {
	return s.String()
}

func (s *GetServiceTemplateParameterConstraintsResponse) SetHeaders(v map[string]*string) *GetServiceTemplateParameterConstraintsResponse {
	s.Headers = v
	return s
}

func (s *GetServiceTemplateParameterConstraintsResponse) SetStatusCode(v int32) *GetServiceTemplateParameterConstraintsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetServiceTemplateParameterConstraintsResponse) SetBody(v *GetServiceTemplateParameterConstraintsResponseBody) *GetServiceTemplateParameterConstraintsResponse {
	s.Body = v
	return s
}

type ListServiceInstanceLogsRequest struct {
	LogSource *string `json:"LogSource,omitempty" xml:"LogSource,omitempty"`
	Logstore  *string `json:"Logstore,omitempty" xml:"Logstore,omitempty"`
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// BBBAAfu+XtuBE55iRLHEYYuojI4=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// si-70a3b15bb626435b****
	ServiceInstanceId *string `json:"ServiceInstanceId,omitempty" xml:"ServiceInstanceId,omitempty"`
}

func (s ListServiceInstanceLogsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListServiceInstanceLogsRequest) GoString() string {
	return s.String()
}

func (s *ListServiceInstanceLogsRequest) SetLogSource(v string) *ListServiceInstanceLogsRequest {
	s.LogSource = &v
	return s
}

func (s *ListServiceInstanceLogsRequest) SetLogstore(v string) *ListServiceInstanceLogsRequest {
	s.Logstore = &v
	return s
}

func (s *ListServiceInstanceLogsRequest) SetMaxResults(v int32) *ListServiceInstanceLogsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListServiceInstanceLogsRequest) SetNextToken(v string) *ListServiceInstanceLogsRequest {
	s.NextToken = &v
	return s
}

func (s *ListServiceInstanceLogsRequest) SetRegionId(v string) *ListServiceInstanceLogsRequest {
	s.RegionId = &v
	return s
}

func (s *ListServiceInstanceLogsRequest) SetServiceInstanceId(v string) *ListServiceInstanceLogsRequest {
	s.ServiceInstanceId = &v
	return s
}

type ListServiceInstanceLogsResponseBody struct {
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// AAAAAfu+XtuBE55iRLHEYYuojI4=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// 51945B04-6AA6-410D-93BA-236E0248B104
	RequestId            *string                                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ServiceInstancesLogs []*ListServiceInstanceLogsResponseBodyServiceInstancesLogs `json:"ServiceInstancesLogs,omitempty" xml:"ServiceInstancesLogs,omitempty" type:"Repeated"`
}

func (s ListServiceInstanceLogsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListServiceInstanceLogsResponseBody) GoString() string {
	return s.String()
}

func (s *ListServiceInstanceLogsResponseBody) SetMaxResults(v int32) *ListServiceInstanceLogsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListServiceInstanceLogsResponseBody) SetNextToken(v string) *ListServiceInstanceLogsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListServiceInstanceLogsResponseBody) SetRequestId(v string) *ListServiceInstanceLogsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListServiceInstanceLogsResponseBody) SetServiceInstancesLogs(v []*ListServiceInstanceLogsResponseBodyServiceInstancesLogs) *ListServiceInstanceLogsResponseBody {
	s.ServiceInstancesLogs = v
	return s
}

type ListServiceInstanceLogsResponseBodyServiceInstancesLogs struct {
	// example:
	//
	// Start creating service instance
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// example:
	//
	// serviceInstance
	LogType *string `json:"LogType,omitempty" xml:"LogType,omitempty"`
	// example:
	//
	// si-5c6525c0589545c3****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// example:
	//
	// ROS.Stack
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// example:
	//
	// ros
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// example:
	//
	// Creating
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 2021-05-21T00:00:00Z
	Timestamp *string `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
}

func (s ListServiceInstanceLogsResponseBodyServiceInstancesLogs) String() string {
	return tea.Prettify(s)
}

func (s ListServiceInstanceLogsResponseBodyServiceInstancesLogs) GoString() string {
	return s.String()
}

func (s *ListServiceInstanceLogsResponseBodyServiceInstancesLogs) SetContent(v string) *ListServiceInstanceLogsResponseBodyServiceInstancesLogs {
	s.Content = &v
	return s
}

func (s *ListServiceInstanceLogsResponseBodyServiceInstancesLogs) SetLogType(v string) *ListServiceInstanceLogsResponseBodyServiceInstancesLogs {
	s.LogType = &v
	return s
}

func (s *ListServiceInstanceLogsResponseBodyServiceInstancesLogs) SetResourceId(v string) *ListServiceInstanceLogsResponseBodyServiceInstancesLogs {
	s.ResourceId = &v
	return s
}

func (s *ListServiceInstanceLogsResponseBodyServiceInstancesLogs) SetResourceType(v string) *ListServiceInstanceLogsResponseBodyServiceInstancesLogs {
	s.ResourceType = &v
	return s
}

func (s *ListServiceInstanceLogsResponseBodyServiceInstancesLogs) SetSource(v string) *ListServiceInstanceLogsResponseBodyServiceInstancesLogs {
	s.Source = &v
	return s
}

func (s *ListServiceInstanceLogsResponseBodyServiceInstancesLogs) SetStatus(v string) *ListServiceInstanceLogsResponseBodyServiceInstancesLogs {
	s.Status = &v
	return s
}

func (s *ListServiceInstanceLogsResponseBodyServiceInstancesLogs) SetTimestamp(v string) *ListServiceInstanceLogsResponseBodyServiceInstancesLogs {
	s.Timestamp = &v
	return s
}

type ListServiceInstanceLogsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListServiceInstanceLogsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListServiceInstanceLogsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListServiceInstanceLogsResponse) GoString() string {
	return s.String()
}

func (s *ListServiceInstanceLogsResponse) SetHeaders(v map[string]*string) *ListServiceInstanceLogsResponse {
	s.Headers = v
	return s
}

func (s *ListServiceInstanceLogsResponse) SetStatusCode(v int32) *ListServiceInstanceLogsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListServiceInstanceLogsResponse) SetBody(v *ListServiceInstanceLogsResponseBody) *ListServiceInstanceLogsResponse {
	s.Body = v
	return s
}

type ListServiceInstanceResourcesRequest struct {
	// example:
	//
	// 2022-03-01T12:00:00
	ExpireTimeEnd *string `json:"ExpireTimeEnd,omitempty" xml:"ExpireTimeEnd,omitempty"`
	// example:
	//
	// 2022-01-01T12:00:00
	ExpireTimeStart *string `json:"ExpireTimeStart,omitempty" xml:"ExpireTimeStart,omitempty"`
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// AAAAAc3HCuYhJi/wvpk4xOr0VLbAx7BkQzyYC+ONO+WudHGKEdB0uWSY7AGnM3qCgm/Ynge7zU6NWdbj0Tegyajyqyc=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// Subscription
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId    *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceARN []*string `json:"ResourceARN,omitempty" xml:"ResourceARN,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// si-d8a0cc2a1ee04dce****
	ServiceInstanceId           *string                                   `json:"ServiceInstanceId,omitempty" xml:"ServiceInstanceId,omitempty"`
	ServiceInstanceResourceType *string                                   `json:"ServiceInstanceResourceType,omitempty" xml:"ServiceInstanceResourceType,omitempty"`
	Tag                         []*ListServiceInstanceResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ListServiceInstanceResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListServiceInstanceResourcesRequest) GoString() string {
	return s.String()
}

func (s *ListServiceInstanceResourcesRequest) SetExpireTimeEnd(v string) *ListServiceInstanceResourcesRequest {
	s.ExpireTimeEnd = &v
	return s
}

func (s *ListServiceInstanceResourcesRequest) SetExpireTimeStart(v string) *ListServiceInstanceResourcesRequest {
	s.ExpireTimeStart = &v
	return s
}

func (s *ListServiceInstanceResourcesRequest) SetMaxResults(v int32) *ListServiceInstanceResourcesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListServiceInstanceResourcesRequest) SetNextToken(v string) *ListServiceInstanceResourcesRequest {
	s.NextToken = &v
	return s
}

func (s *ListServiceInstanceResourcesRequest) SetPayType(v string) *ListServiceInstanceResourcesRequest {
	s.PayType = &v
	return s
}

func (s *ListServiceInstanceResourcesRequest) SetRegionId(v string) *ListServiceInstanceResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *ListServiceInstanceResourcesRequest) SetResourceARN(v []*string) *ListServiceInstanceResourcesRequest {
	s.ResourceARN = v
	return s
}

func (s *ListServiceInstanceResourcesRequest) SetServiceInstanceId(v string) *ListServiceInstanceResourcesRequest {
	s.ServiceInstanceId = &v
	return s
}

func (s *ListServiceInstanceResourcesRequest) SetServiceInstanceResourceType(v string) *ListServiceInstanceResourcesRequest {
	s.ServiceInstanceResourceType = &v
	return s
}

func (s *ListServiceInstanceResourcesRequest) SetTag(v []*ListServiceInstanceResourcesRequestTag) *ListServiceInstanceResourcesRequest {
	s.Tag = v
	return s
}

type ListServiceInstanceResourcesRequestTag struct {
	// example:
	//
	// key1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// value1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListServiceInstanceResourcesRequestTag) String() string {
	return tea.Prettify(s)
}

func (s ListServiceInstanceResourcesRequestTag) GoString() string {
	return s.String()
}

func (s *ListServiceInstanceResourcesRequestTag) SetKey(v string) *ListServiceInstanceResourcesRequestTag {
	s.Key = &v
	return s
}

func (s *ListServiceInstanceResourcesRequestTag) SetValue(v string) *ListServiceInstanceResourcesRequestTag {
	s.Value = &v
	return s
}

type ListServiceInstanceResourcesResponseBody struct {
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// AAAAAc3HCuYhJi/wvpk4xOr0VLbAx7BkQzyYC+ONO+WudHGKEdB0uWSY7AGnM3qCgm/Ynge7zU6NWdbj0Tegyajyqyc=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// B288A0BE-D927-4888-B0F7-B35EF84B6E6F
	RequestId *string                                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Resources []*ListServiceInstanceResourcesResponseBodyResources `json:"Resources,omitempty" xml:"Resources,omitempty" type:"Repeated"`
}

func (s ListServiceInstanceResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListServiceInstanceResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *ListServiceInstanceResourcesResponseBody) SetMaxResults(v int32) *ListServiceInstanceResourcesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListServiceInstanceResourcesResponseBody) SetNextToken(v string) *ListServiceInstanceResourcesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListServiceInstanceResourcesResponseBody) SetRequestId(v string) *ListServiceInstanceResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListServiceInstanceResourcesResponseBody) SetResources(v []*ListServiceInstanceResourcesResponseBodyResources) *ListServiceInstanceResourcesResponseBody {
	s.Resources = v
	return s
}

type ListServiceInstanceResourcesResponseBodyResources struct {
	// example:
	//
	// 2022-01-01T12:00:00
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 2022-03-01T12:00:00
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// example:
	//
	// Subscription
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// example:
	//
	// rds
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// example:
	//
	// RDS
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	// example:
	//
	// AutoRenewal
	RenewStatus *string `json:"RenewStatus,omitempty" xml:"RenewStatus,omitempty"`
	// example:
	//
	// 1
	RenewalPeriod *int32 `json:"RenewalPeriod,omitempty" xml:"RenewalPeriod,omitempty"`
	// example:
	//
	// Month
	RenewalPeriodUnit *string `json:"RenewalPeriodUnit,omitempty" xml:"RenewalPeriodUnit,omitempty"`
	// example:
	//
	// arn:acs:sag:cn-hangzhou:130920852836****:ccn/ccn-b3qf0q439sq2de****
	ResourceARN *string `json:"ResourceARN,omitempty" xml:"ResourceARN,omitempty"`
	Status      *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListServiceInstanceResourcesResponseBodyResources) String() string {
	return tea.Prettify(s)
}

func (s ListServiceInstanceResourcesResponseBodyResources) GoString() string {
	return s.String()
}

func (s *ListServiceInstanceResourcesResponseBodyResources) SetCreateTime(v string) *ListServiceInstanceResourcesResponseBodyResources {
	s.CreateTime = &v
	return s
}

func (s *ListServiceInstanceResourcesResponseBodyResources) SetExpireTime(v string) *ListServiceInstanceResourcesResponseBodyResources {
	s.ExpireTime = &v
	return s
}

func (s *ListServiceInstanceResourcesResponseBodyResources) SetPayType(v string) *ListServiceInstanceResourcesResponseBodyResources {
	s.PayType = &v
	return s
}

func (s *ListServiceInstanceResourcesResponseBodyResources) SetProductCode(v string) *ListServiceInstanceResourcesResponseBodyResources {
	s.ProductCode = &v
	return s
}

func (s *ListServiceInstanceResourcesResponseBodyResources) SetProductType(v string) *ListServiceInstanceResourcesResponseBodyResources {
	s.ProductType = &v
	return s
}

func (s *ListServiceInstanceResourcesResponseBodyResources) SetRenewStatus(v string) *ListServiceInstanceResourcesResponseBodyResources {
	s.RenewStatus = &v
	return s
}

func (s *ListServiceInstanceResourcesResponseBodyResources) SetRenewalPeriod(v int32) *ListServiceInstanceResourcesResponseBodyResources {
	s.RenewalPeriod = &v
	return s
}

func (s *ListServiceInstanceResourcesResponseBodyResources) SetRenewalPeriodUnit(v string) *ListServiceInstanceResourcesResponseBodyResources {
	s.RenewalPeriodUnit = &v
	return s
}

func (s *ListServiceInstanceResourcesResponseBodyResources) SetResourceARN(v string) *ListServiceInstanceResourcesResponseBodyResources {
	s.ResourceARN = &v
	return s
}

func (s *ListServiceInstanceResourcesResponseBodyResources) SetStatus(v string) *ListServiceInstanceResourcesResponseBodyResources {
	s.Status = &v
	return s
}

type ListServiceInstanceResourcesResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListServiceInstanceResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListServiceInstanceResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListServiceInstanceResourcesResponse) GoString() string {
	return s.String()
}

func (s *ListServiceInstanceResourcesResponse) SetHeaders(v map[string]*string) *ListServiceInstanceResourcesResponse {
	s.Headers = v
	return s
}

func (s *ListServiceInstanceResourcesResponse) SetStatusCode(v int32) *ListServiceInstanceResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListServiceInstanceResourcesResponse) SetBody(v *ListServiceInstanceResourcesResponseBody) *ListServiceInstanceResourcesResponse {
	s.Body = v
	return s
}

type ListServiceInstancesRequest struct {
	Filter []*ListServiceInstancesRequestFilter `json:"Filter,omitempty" xml:"Filter,omitempty" type:"Repeated"`
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// BBBAAfu+XtuBE55iRLHEYYuojI4=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// rg-bp67acfmxazb4p****
	ResourceGroupId *string                           `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Tag             []*ListServiceInstancesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ListServiceInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListServiceInstancesRequest) GoString() string {
	return s.String()
}

func (s *ListServiceInstancesRequest) SetFilter(v []*ListServiceInstancesRequestFilter) *ListServiceInstancesRequest {
	s.Filter = v
	return s
}

func (s *ListServiceInstancesRequest) SetMaxResults(v int32) *ListServiceInstancesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListServiceInstancesRequest) SetNextToken(v string) *ListServiceInstancesRequest {
	s.NextToken = &v
	return s
}

func (s *ListServiceInstancesRequest) SetRegionId(v string) *ListServiceInstancesRequest {
	s.RegionId = &v
	return s
}

func (s *ListServiceInstancesRequest) SetResourceGroupId(v string) *ListServiceInstancesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListServiceInstancesRequest) SetTag(v []*ListServiceInstancesRequestTag) *ListServiceInstancesRequest {
	s.Tag = v
	return s
}

type ListServiceInstancesRequestFilter struct {
	// example:
	//
	// ServiceInstanceId
	Name  *string   `json:"Name,omitempty" xml:"Name,omitempty"`
	Value []*string `json:"Value,omitempty" xml:"Value,omitempty" type:"Repeated"`
}

func (s ListServiceInstancesRequestFilter) String() string {
	return tea.Prettify(s)
}

func (s ListServiceInstancesRequestFilter) GoString() string {
	return s.String()
}

func (s *ListServiceInstancesRequestFilter) SetName(v string) *ListServiceInstancesRequestFilter {
	s.Name = &v
	return s
}

func (s *ListServiceInstancesRequestFilter) SetValue(v []*string) *ListServiceInstancesRequestFilter {
	s.Value = v
	return s
}

type ListServiceInstancesRequestTag struct {
	// example:
	//
	// key1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// value1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListServiceInstancesRequestTag) String() string {
	return tea.Prettify(s)
}

func (s ListServiceInstancesRequestTag) GoString() string {
	return s.String()
}

func (s *ListServiceInstancesRequestTag) SetKey(v string) *ListServiceInstancesRequestTag {
	s.Key = &v
	return s
}

func (s *ListServiceInstancesRequestTag) SetValue(v string) *ListServiceInstancesRequestTag {
	s.Value = &v
	return s
}

type ListServiceInstancesResponseBody struct {
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// AAAAAfu+XtuBE55iRLHEYYuojI4=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// E50287CB-AABF-4877-92C0-289B339A1546
	RequestId        *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ServiceInstances []*ListServiceInstancesResponseBodyServiceInstances `json:"ServiceInstances,omitempty" xml:"ServiceInstances,omitempty" type:"Repeated"`
	// example:
	//
	// 100
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListServiceInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListServiceInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *ListServiceInstancesResponseBody) SetMaxResults(v int32) *ListServiceInstancesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListServiceInstancesResponseBody) SetNextToken(v string) *ListServiceInstancesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListServiceInstancesResponseBody) SetRequestId(v string) *ListServiceInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListServiceInstancesResponseBody) SetServiceInstances(v []*ListServiceInstancesResponseBodyServiceInstances) *ListServiceInstancesResponseBody {
	s.ServiceInstances = v
	return s
}

func (s *ListServiceInstancesResponseBody) SetTotalCount(v int64) *ListServiceInstancesResponseBody {
	s.TotalCount = &v
	return s
}

type ListServiceInstancesResponseBodyServiceInstances struct {
	BizStatus *string `json:"BizStatus,omitempty" xml:"BizStatus,omitempty"`
	// example:
	//
	// 2021-05-20T00:00:00Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// true
	EnableInstanceOps *bool `json:"EnableInstanceOps,omitempty" xml:"EnableInstanceOps,omitempty"`
	// example:
	//
	// 2022-01-01T12:00:00
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 5827****
	MarketInstanceId *string `json:"MarketInstanceId,omitempty" xml:"MarketInstanceId,omitempty"`
	// example:
	//
	// TestName
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// si-d6ab3a63ccbb4b17****
	OperatedServiceInstanceId *string `json:"OperatedServiceInstanceId,omitempty" xml:"OperatedServiceInstanceId,omitempty"`
	// example:
	//
	// 2022-01-28T06:48:56Z
	OperationEndTime *string `json:"OperationEndTime,omitempty" xml:"OperationEndTime,omitempty"`
	// example:
	//
	// 2021-12-29T06:48:56Z
	OperationStartTime *string `json:"OperationStartTime,omitempty" xml:"OperationStartTime,omitempty"`
	OrderId            *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// example:
	//
	// {"managementUrl": "http://xx.xx"}
	Outputs *string `json:"Outputs,omitempty" xml:"Outputs,omitempty"`
	// example:
	//
	// {"param":"value"}
	Parameters *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// example:
	//
	// Subscription
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// example:
	//
	// 90
	Progress *int64 `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// example:
	//
	// rg-aekz6scpcxxxxxx
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// example:
	//
	// [{"StackId": "stack-xxx"}]
	Resources *string                                                  `json:"Resources,omitempty" xml:"Resources,omitempty"`
	Service   *ListServiceInstancesResponseBodyServiceInstancesService `json:"Service,omitempty" xml:"Service,omitempty" type:"Struct"`
	// example:
	//
	// si-d6ab3a63ccbb4b17****
	ServiceInstanceId *string `json:"ServiceInstanceId,omitempty" xml:"ServiceInstanceId,omitempty"`
	// example:
	//
	// private
	ServiceType *string `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
	// example:
	//
	// Supplier
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// example:
	//
	// Deployed
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// deploy successfully
	StatusDetail *string                                                 `json:"StatusDetail,omitempty" xml:"StatusDetail,omitempty"`
	Tags         []*ListServiceInstancesResponseBodyServiceInstancesTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	TemplateName *string                                                 `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// example:
	//
	// 2021-05-20T00:00:00Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListServiceInstancesResponseBodyServiceInstances) String() string {
	return tea.Prettify(s)
}

func (s ListServiceInstancesResponseBodyServiceInstances) GoString() string {
	return s.String()
}

func (s *ListServiceInstancesResponseBodyServiceInstances) SetBizStatus(v string) *ListServiceInstancesResponseBodyServiceInstances {
	s.BizStatus = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstances) SetCreateTime(v string) *ListServiceInstancesResponseBodyServiceInstances {
	s.CreateTime = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstances) SetEnableInstanceOps(v bool) *ListServiceInstancesResponseBodyServiceInstances {
	s.EnableInstanceOps = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstances) SetEndTime(v string) *ListServiceInstancesResponseBodyServiceInstances {
	s.EndTime = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstances) SetMarketInstanceId(v string) *ListServiceInstancesResponseBodyServiceInstances {
	s.MarketInstanceId = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstances) SetName(v string) *ListServiceInstancesResponseBodyServiceInstances {
	s.Name = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstances) SetOperatedServiceInstanceId(v string) *ListServiceInstancesResponseBodyServiceInstances {
	s.OperatedServiceInstanceId = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstances) SetOperationEndTime(v string) *ListServiceInstancesResponseBodyServiceInstances {
	s.OperationEndTime = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstances) SetOperationStartTime(v string) *ListServiceInstancesResponseBodyServiceInstances {
	s.OperationStartTime = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstances) SetOrderId(v string) *ListServiceInstancesResponseBodyServiceInstances {
	s.OrderId = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstances) SetOutputs(v string) *ListServiceInstancesResponseBodyServiceInstances {
	s.Outputs = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstances) SetParameters(v string) *ListServiceInstancesResponseBodyServiceInstances {
	s.Parameters = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstances) SetPayType(v string) *ListServiceInstancesResponseBodyServiceInstances {
	s.PayType = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstances) SetProgress(v int64) *ListServiceInstancesResponseBodyServiceInstances {
	s.Progress = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstances) SetResourceGroupId(v string) *ListServiceInstancesResponseBodyServiceInstances {
	s.ResourceGroupId = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstances) SetResources(v string) *ListServiceInstancesResponseBodyServiceInstances {
	s.Resources = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstances) SetService(v *ListServiceInstancesResponseBodyServiceInstancesService) *ListServiceInstancesResponseBodyServiceInstances {
	s.Service = v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstances) SetServiceInstanceId(v string) *ListServiceInstancesResponseBodyServiceInstances {
	s.ServiceInstanceId = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstances) SetServiceType(v string) *ListServiceInstancesResponseBodyServiceInstances {
	s.ServiceType = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstances) SetSource(v string) *ListServiceInstancesResponseBodyServiceInstances {
	s.Source = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstances) SetStatus(v string) *ListServiceInstancesResponseBodyServiceInstances {
	s.Status = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstances) SetStatusDetail(v string) *ListServiceInstancesResponseBodyServiceInstances {
	s.StatusDetail = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstances) SetTags(v []*ListServiceInstancesResponseBodyServiceInstancesTags) *ListServiceInstancesResponseBodyServiceInstances {
	s.Tags = v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstances) SetTemplateName(v string) *ListServiceInstancesResponseBodyServiceInstances {
	s.TemplateName = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstances) SetUpdateTime(v string) *ListServiceInstancesResponseBodyServiceInstances {
	s.UpdateTime = &v
	return s
}

type ListServiceInstancesResponseBodyServiceInstancesService struct {
	Commodity *ListServiceInstancesResponseBodyServiceInstancesServiceCommodity `json:"Commodity,omitempty" xml:"Commodity,omitempty" type:"Struct"`
	// example:
	//
	// ros
	DeployType *string `json:"DeployType,omitempty" xml:"DeployType,omitempty"`
	// example:
	//
	// 2021-05-21T00:00:00Z
	PublishTime *string `json:"PublishTime,omitempty" xml:"PublishTime,omitempty"`
	// example:
	//
	// service-0e6fca6a51a54420****
	ServiceId    *string                                                                `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	ServiceInfos []*ListServiceInstancesResponseBodyServiceInstancesServiceServiceInfos `json:"ServiceInfos,omitempty" xml:"ServiceInfos,omitempty" type:"Repeated"`
	// example:
	//
	// private
	ServiceType *string `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
	// example:
	//
	// Online
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
	SupplierName *string `json:"SupplierName,omitempty" xml:"SupplierName,omitempty"`
	// example:
	//
	// http://example.com
	SupplierUrl *string `json:"SupplierUrl,omitempty" xml:"SupplierUrl,omitempty"`
	// example:
	//
	// 1.0
	Version     *string `json:"Version,omitempty" xml:"Version,omitempty"`
	VersionName *string `json:"VersionName,omitempty" xml:"VersionName,omitempty"`
}

func (s ListServiceInstancesResponseBodyServiceInstancesService) String() string {
	return tea.Prettify(s)
}

func (s ListServiceInstancesResponseBodyServiceInstancesService) GoString() string {
	return s.String()
}

func (s *ListServiceInstancesResponseBodyServiceInstancesService) SetCommodity(v *ListServiceInstancesResponseBodyServiceInstancesServiceCommodity) *ListServiceInstancesResponseBodyServiceInstancesService {
	s.Commodity = v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstancesService) SetDeployType(v string) *ListServiceInstancesResponseBodyServiceInstancesService {
	s.DeployType = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstancesService) SetPublishTime(v string) *ListServiceInstancesResponseBodyServiceInstancesService {
	s.PublishTime = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstancesService) SetServiceId(v string) *ListServiceInstancesResponseBodyServiceInstancesService {
	s.ServiceId = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstancesService) SetServiceInfos(v []*ListServiceInstancesResponseBodyServiceInstancesServiceServiceInfos) *ListServiceInstancesResponseBodyServiceInstancesService {
	s.ServiceInfos = v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstancesService) SetServiceType(v string) *ListServiceInstancesResponseBodyServiceInstancesService {
	s.ServiceType = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstancesService) SetStatus(v string) *ListServiceInstancesResponseBodyServiceInstancesService {
	s.Status = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstancesService) SetSupplierName(v string) *ListServiceInstancesResponseBodyServiceInstancesService {
	s.SupplierName = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstancesService) SetSupplierUrl(v string) *ListServiceInstancesResponseBodyServiceInstancesService {
	s.SupplierUrl = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstancesService) SetVersion(v string) *ListServiceInstancesResponseBodyServiceInstancesService {
	s.Version = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstancesService) SetVersionName(v string) *ListServiceInstancesResponseBodyServiceInstancesService {
	s.VersionName = &v
	return s
}

type ListServiceInstancesResponseBodyServiceInstancesServiceCommodity struct {
	SaasBoostMetadata *string `json:"SaasBoostMetadata,omitempty" xml:"SaasBoostMetadata,omitempty"`
	Type              *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListServiceInstancesResponseBodyServiceInstancesServiceCommodity) String() string {
	return tea.Prettify(s)
}

func (s ListServiceInstancesResponseBodyServiceInstancesServiceCommodity) GoString() string {
	return s.String()
}

func (s *ListServiceInstancesResponseBodyServiceInstancesServiceCommodity) SetSaasBoostMetadata(v string) *ListServiceInstancesResponseBodyServiceInstancesServiceCommodity {
	s.SaasBoostMetadata = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstancesServiceCommodity) SetType(v string) *ListServiceInstancesResponseBodyServiceInstancesServiceCommodity {
	s.Type = &v
	return s
}

type ListServiceInstancesResponseBodyServiceInstancesServiceServiceInfos struct {
	// example:
	//
	// https://example.com/service-image/c1c4a559-cc60-4af1-b976-98f356602462.png
	Image *string `json:"Image,omitempty" xml:"Image,omitempty"`
	// example:
	//
	// zh-CN
	Locale           *string `json:"Locale,omitempty" xml:"Locale,omitempty"`
	Name             *string `json:"Name,omitempty" xml:"Name,omitempty"`
	ShortDescription *string `json:"ShortDescription,omitempty" xml:"ShortDescription,omitempty"`
}

func (s ListServiceInstancesResponseBodyServiceInstancesServiceServiceInfos) String() string {
	return tea.Prettify(s)
}

func (s ListServiceInstancesResponseBodyServiceInstancesServiceServiceInfos) GoString() string {
	return s.String()
}

func (s *ListServiceInstancesResponseBodyServiceInstancesServiceServiceInfos) SetImage(v string) *ListServiceInstancesResponseBodyServiceInstancesServiceServiceInfos {
	s.Image = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstancesServiceServiceInfos) SetLocale(v string) *ListServiceInstancesResponseBodyServiceInstancesServiceServiceInfos {
	s.Locale = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstancesServiceServiceInfos) SetName(v string) *ListServiceInstancesResponseBodyServiceInstancesServiceServiceInfos {
	s.Name = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstancesServiceServiceInfos) SetShortDescription(v string) *ListServiceInstancesResponseBodyServiceInstancesServiceServiceInfos {
	s.ShortDescription = &v
	return s
}

type ListServiceInstancesResponseBodyServiceInstancesTags struct {
	// example:
	//
	// key1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// value1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListServiceInstancesResponseBodyServiceInstancesTags) String() string {
	return tea.Prettify(s)
}

func (s ListServiceInstancesResponseBodyServiceInstancesTags) GoString() string {
	return s.String()
}

func (s *ListServiceInstancesResponseBodyServiceInstancesTags) SetKey(v string) *ListServiceInstancesResponseBodyServiceInstancesTags {
	s.Key = &v
	return s
}

func (s *ListServiceInstancesResponseBodyServiceInstancesTags) SetValue(v string) *ListServiceInstancesResponseBodyServiceInstancesTags {
	s.Value = &v
	return s
}

type ListServiceInstancesResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListServiceInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListServiceInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListServiceInstancesResponse) GoString() string {
	return s.String()
}

func (s *ListServiceInstancesResponse) SetHeaders(v map[string]*string) *ListServiceInstancesResponse {
	s.Headers = v
	return s
}

func (s *ListServiceInstancesResponse) SetStatusCode(v int32) *ListServiceInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListServiceInstancesResponse) SetBody(v *ListServiceInstancesResponseBody) *ListServiceInstancesResponse {
	s.Body = v
	return s
}

type RestartServiceInstanceRequest struct {
	// example:
	//
	// 123e4567-e89b-12d3-a456-42665544****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// si-d6ab3a63ccbb4b17****
	ServiceInstanceId *string `json:"ServiceInstanceId,omitempty" xml:"ServiceInstanceId,omitempty"`
}

func (s RestartServiceInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s RestartServiceInstanceRequest) GoString() string {
	return s.String()
}

func (s *RestartServiceInstanceRequest) SetClientToken(v string) *RestartServiceInstanceRequest {
	s.ClientToken = &v
	return s
}

func (s *RestartServiceInstanceRequest) SetRegionId(v string) *RestartServiceInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *RestartServiceInstanceRequest) SetServiceInstanceId(v string) *RestartServiceInstanceRequest {
	s.ServiceInstanceId = &v
	return s
}

type RestartServiceInstanceResponseBody struct {
	// example:
	//
	// 13FE89A5-C036-56BF-A0FF-A31C59819FD7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RestartServiceInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RestartServiceInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *RestartServiceInstanceResponseBody) SetRequestId(v string) *RestartServiceInstanceResponseBody {
	s.RequestId = &v
	return s
}

type RestartServiceInstanceResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RestartServiceInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RestartServiceInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s RestartServiceInstanceResponse) GoString() string {
	return s.String()
}

func (s *RestartServiceInstanceResponse) SetHeaders(v map[string]*string) *RestartServiceInstanceResponse {
	s.Headers = v
	return s
}

func (s *RestartServiceInstanceResponse) SetStatusCode(v int32) *RestartServiceInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *RestartServiceInstanceResponse) SetBody(v *RestartServiceInstanceResponseBody) *RestartServiceInstanceResponse {
	s.Body = v
	return s
}

type StartServiceInstanceRequest struct {
	// example:
	//
	// 123e4567-e89b-12d3-a456-42665544****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// si-d6ab3a63ccbb4b17****
	ServiceInstanceId *string `json:"ServiceInstanceId,omitempty" xml:"ServiceInstanceId,omitempty"`
}

func (s StartServiceInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s StartServiceInstanceRequest) GoString() string {
	return s.String()
}

func (s *StartServiceInstanceRequest) SetClientToken(v string) *StartServiceInstanceRequest {
	s.ClientToken = &v
	return s
}

func (s *StartServiceInstanceRequest) SetRegionId(v string) *StartServiceInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *StartServiceInstanceRequest) SetServiceInstanceId(v string) *StartServiceInstanceRequest {
	s.ServiceInstanceId = &v
	return s
}

type StartServiceInstanceResponseBody struct {
	// example:
	//
	// 464C8CB6-A548-5206-B83C-D32A8E43EC21
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StartServiceInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StartServiceInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *StartServiceInstanceResponseBody) SetRequestId(v string) *StartServiceInstanceResponseBody {
	s.RequestId = &v
	return s
}

type StartServiceInstanceResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartServiceInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartServiceInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s StartServiceInstanceResponse) GoString() string {
	return s.String()
}

func (s *StartServiceInstanceResponse) SetHeaders(v map[string]*string) *StartServiceInstanceResponse {
	s.Headers = v
	return s
}

func (s *StartServiceInstanceResponse) SetStatusCode(v int32) *StartServiceInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *StartServiceInstanceResponse) SetBody(v *StartServiceInstanceResponseBody) *StartServiceInstanceResponse {
	s.Body = v
	return s
}

type StopServiceInstanceRequest struct {
	// example:
	//
	// 123e4567-e89b-12d3-a456-42665544****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// si-b58c874912fc4294****
	ServiceInstanceId *string `json:"ServiceInstanceId,omitempty" xml:"ServiceInstanceId,omitempty"`
}

func (s StopServiceInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s StopServiceInstanceRequest) GoString() string {
	return s.String()
}

func (s *StopServiceInstanceRequest) SetClientToken(v string) *StopServiceInstanceRequest {
	s.ClientToken = &v
	return s
}

func (s *StopServiceInstanceRequest) SetRegionId(v string) *StopServiceInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *StopServiceInstanceRequest) SetServiceInstanceId(v string) *StopServiceInstanceRequest {
	s.ServiceInstanceId = &v
	return s
}

type StopServiceInstanceResponseBody struct {
	// example:
	//
	// 4DB0F536-B3BE-4F0D-BD29-E83FB56D550C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StopServiceInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StopServiceInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *StopServiceInstanceResponseBody) SetRequestId(v string) *StopServiceInstanceResponseBody {
	s.RequestId = &v
	return s
}

type StopServiceInstanceResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopServiceInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopServiceInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s StopServiceInstanceResponse) GoString() string {
	return s.String()
}

func (s *StopServiceInstanceResponse) SetHeaders(v map[string]*string) *StopServiceInstanceResponse {
	s.Headers = v
	return s
}

func (s *StopServiceInstanceResponse) SetStatusCode(v int32) *StopServiceInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *StopServiceInstanceResponse) SetBody(v *StopServiceInstanceResponseBody) *StopServiceInstanceResponse {
	s.Body = v
	return s
}

type UpdateServiceInstanceSpecRequest struct {
	// example:
	//
	// 123e4567-e89b-12d3-a456-42665544****
	ClientToken *string                                    `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Commodity   *UpdateServiceInstanceSpecRequestCommodity `json:"Commodity,omitempty" xml:"Commodity,omitempty" type:"Struct"`
	// example:
	//
	// true
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// example:
	//
	// true
	EnableUserPrometheus *bool   `json:"EnableUserPrometheus,omitempty" xml:"EnableUserPrometheus,omitempty"`
	OperationName        *string `json:"OperationName,omitempty" xml:"OperationName,omitempty"`
	// example:
	//
	// {
	//
	//   "InstanceType": "ecs.g8ise.2xlarge"
	//
	// }
	Parameters               map[string]interface{} `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	PredefinedParametersName *string                `json:"PredefinedParametersName,omitempty" xml:"PredefinedParametersName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// si-d6ab3a63ccbb4b17****
	ServiceInstanceId *string `json:"ServiceInstanceId,omitempty" xml:"ServiceInstanceId,omitempty"`
}

func (s UpdateServiceInstanceSpecRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateServiceInstanceSpecRequest) GoString() string {
	return s.String()
}

func (s *UpdateServiceInstanceSpecRequest) SetClientToken(v string) *UpdateServiceInstanceSpecRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateServiceInstanceSpecRequest) SetCommodity(v *UpdateServiceInstanceSpecRequestCommodity) *UpdateServiceInstanceSpecRequest {
	s.Commodity = v
	return s
}

func (s *UpdateServiceInstanceSpecRequest) SetDryRun(v bool) *UpdateServiceInstanceSpecRequest {
	s.DryRun = &v
	return s
}

func (s *UpdateServiceInstanceSpecRequest) SetEnableUserPrometheus(v bool) *UpdateServiceInstanceSpecRequest {
	s.EnableUserPrometheus = &v
	return s
}

func (s *UpdateServiceInstanceSpecRequest) SetOperationName(v string) *UpdateServiceInstanceSpecRequest {
	s.OperationName = &v
	return s
}

func (s *UpdateServiceInstanceSpecRequest) SetParameters(v map[string]interface{}) *UpdateServiceInstanceSpecRequest {
	s.Parameters = v
	return s
}

func (s *UpdateServiceInstanceSpecRequest) SetPredefinedParametersName(v string) *UpdateServiceInstanceSpecRequest {
	s.PredefinedParametersName = &v
	return s
}

func (s *UpdateServiceInstanceSpecRequest) SetServiceInstanceId(v string) *UpdateServiceInstanceSpecRequest {
	s.ServiceInstanceId = &v
	return s
}

type UpdateServiceInstanceSpecRequestCommodity struct {
	// example:
	//
	// true
	AutoPay *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
}

func (s UpdateServiceInstanceSpecRequestCommodity) String() string {
	return tea.Prettify(s)
}

func (s UpdateServiceInstanceSpecRequestCommodity) GoString() string {
	return s.String()
}

func (s *UpdateServiceInstanceSpecRequestCommodity) SetAutoPay(v bool) *UpdateServiceInstanceSpecRequestCommodity {
	s.AutoPay = &v
	return s
}

type UpdateServiceInstanceSpecShrinkRequest struct {
	// example:
	//
	// 123e4567-e89b-12d3-a456-42665544****
	ClientToken *string                                          `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Commodity   *UpdateServiceInstanceSpecShrinkRequestCommodity `json:"Commodity,omitempty" xml:"Commodity,omitempty" type:"Struct"`
	// example:
	//
	// true
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// example:
	//
	// true
	EnableUserPrometheus *bool   `json:"EnableUserPrometheus,omitempty" xml:"EnableUserPrometheus,omitempty"`
	OperationName        *string `json:"OperationName,omitempty" xml:"OperationName,omitempty"`
	// example:
	//
	// {
	//
	//   "InstanceType": "ecs.g8ise.2xlarge"
	//
	// }
	ParametersShrink         *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	PredefinedParametersName *string `json:"PredefinedParametersName,omitempty" xml:"PredefinedParametersName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// si-d6ab3a63ccbb4b17****
	ServiceInstanceId *string `json:"ServiceInstanceId,omitempty" xml:"ServiceInstanceId,omitempty"`
}

func (s UpdateServiceInstanceSpecShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateServiceInstanceSpecShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateServiceInstanceSpecShrinkRequest) SetClientToken(v string) *UpdateServiceInstanceSpecShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateServiceInstanceSpecShrinkRequest) SetCommodity(v *UpdateServiceInstanceSpecShrinkRequestCommodity) *UpdateServiceInstanceSpecShrinkRequest {
	s.Commodity = v
	return s
}

func (s *UpdateServiceInstanceSpecShrinkRequest) SetDryRun(v bool) *UpdateServiceInstanceSpecShrinkRequest {
	s.DryRun = &v
	return s
}

func (s *UpdateServiceInstanceSpecShrinkRequest) SetEnableUserPrometheus(v bool) *UpdateServiceInstanceSpecShrinkRequest {
	s.EnableUserPrometheus = &v
	return s
}

func (s *UpdateServiceInstanceSpecShrinkRequest) SetOperationName(v string) *UpdateServiceInstanceSpecShrinkRequest {
	s.OperationName = &v
	return s
}

func (s *UpdateServiceInstanceSpecShrinkRequest) SetParametersShrink(v string) *UpdateServiceInstanceSpecShrinkRequest {
	s.ParametersShrink = &v
	return s
}

func (s *UpdateServiceInstanceSpecShrinkRequest) SetPredefinedParametersName(v string) *UpdateServiceInstanceSpecShrinkRequest {
	s.PredefinedParametersName = &v
	return s
}

func (s *UpdateServiceInstanceSpecShrinkRequest) SetServiceInstanceId(v string) *UpdateServiceInstanceSpecShrinkRequest {
	s.ServiceInstanceId = &v
	return s
}

type UpdateServiceInstanceSpecShrinkRequestCommodity struct {
	// example:
	//
	// true
	AutoPay *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
}

func (s UpdateServiceInstanceSpecShrinkRequestCommodity) String() string {
	return tea.Prettify(s)
}

func (s UpdateServiceInstanceSpecShrinkRequestCommodity) GoString() string {
	return s.String()
}

func (s *UpdateServiceInstanceSpecShrinkRequestCommodity) SetAutoPay(v bool) *UpdateServiceInstanceSpecShrinkRequestCommodity {
	s.AutoPay = &v
	return s
}

type UpdateServiceInstanceSpecResponseBody struct {
	// example:
	//
	// 23396265896****
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// example:
	//
	// 06BF8F22-02DC-4750-83DF-3FFC11C065EA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateServiceInstanceSpecResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateServiceInstanceSpecResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateServiceInstanceSpecResponseBody) SetOrderId(v string) *UpdateServiceInstanceSpecResponseBody {
	s.OrderId = &v
	return s
}

func (s *UpdateServiceInstanceSpecResponseBody) SetRequestId(v string) *UpdateServiceInstanceSpecResponseBody {
	s.RequestId = &v
	return s
}

type UpdateServiceInstanceSpecResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateServiceInstanceSpecResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateServiceInstanceSpecResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateServiceInstanceSpecResponse) GoString() string {
	return s.String()
}

func (s *UpdateServiceInstanceSpecResponse) SetHeaders(v map[string]*string) *UpdateServiceInstanceSpecResponse {
	s.Headers = v
	return s
}

func (s *UpdateServiceInstanceSpecResponse) SetStatusCode(v int32) *UpdateServiceInstanceSpecResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateServiceInstanceSpecResponse) SetBody(v *UpdateServiceInstanceSpecResponseBody) *UpdateServiceInstanceSpecResponse {
	s.Body = v
	return s
}

type Client struct {
	openapi.Client
}

func NewClient(config *openapi.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapi.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = tea.String("regional")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("computenest"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !tea.BoolValue(util.Empty(endpoint)) {
		_result = endpoint
		return _result, _err
	}

	if !tea.BoolValue(util.IsUnset(endpointMap)) && !tea.BoolValue(util.Empty(endpointMap[tea.StringValue(regionId)])) {
		_result = endpointMap[tea.StringValue(regionId)]
		return _result, _err
	}

	_body, _err := endpointutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 资源组转组
//
// @param request - ChangeResourceGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChangeResourceGroupResponse
func (client *Client) ChangeResourceGroupWithOptions(request *ChangeResourceGroupRequest, runtime *util.RuntimeOptions) (_result *ChangeResourceGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.NewResourceGroupId)) {
		query["NewResourceGroupId"] = request.NewResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ChangeResourceGroup"),
		Version:     tea.String("2021-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ChangeResourceGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 资源组转组
//
// @param request - ChangeResourceGroupRequest
//
// @return ChangeResourceGroupResponse
func (client *Client) ChangeResourceGroup(request *ChangeResourceGroupRequest) (_result *ChangeResourceGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ChangeResourceGroupResponse{}
	_body, _err := client.ChangeResourceGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 继续部署服务实例
//
// @param request - ContinueDeployServiceInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ContinueDeployServiceInstanceResponse
func (client *Client) ContinueDeployServiceInstanceWithOptions(request *ContinueDeployServiceInstanceRequest, runtime *util.RuntimeOptions) (_result *ContinueDeployServiceInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.Option)) {
		query["Option"] = request.Option
	}

	if !tea.BoolValue(util.IsUnset(request.Parameters)) {
		query["Parameters"] = request.Parameters
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceInstanceId)) {
		query["ServiceInstanceId"] = request.ServiceInstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ContinueDeployServiceInstance"),
		Version:     tea.String("2021-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ContinueDeployServiceInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 继续部署服务实例
//
// @param request - ContinueDeployServiceInstanceRequest
//
// @return ContinueDeployServiceInstanceResponse
func (client *Client) ContinueDeployServiceInstance(request *ContinueDeployServiceInstanceRequest) (_result *ContinueDeployServiceInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ContinueDeployServiceInstanceResponse{}
	_body, _err := client.ContinueDeployServiceInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param tmpReq - CreateServiceInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateServiceInstanceResponse
func (client *Client) CreateServiceInstanceWithOptions(tmpReq *CreateServiceInstanceRequest, runtime *util.RuntimeOptions) (_result *CreateServiceInstanceResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateServiceInstanceShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Parameters)) {
		request.ParametersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Parameters, tea.String("Parameters"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Commodity)) {
		query["Commodity"] = request.Commodity
	}

	if !tea.BoolValue(util.IsUnset(request.ContactGroup)) {
		query["ContactGroup"] = request.ContactGroup
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.EnableInstanceOps)) {
		query["EnableInstanceOps"] = request.EnableInstanceOps
	}

	if !tea.BoolValue(util.IsUnset(request.EnableUserPrometheus)) {
		query["EnableUserPrometheus"] = request.EnableUserPrometheus
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.OperationMetadata)) {
		query["OperationMetadata"] = request.OperationMetadata
	}

	if !tea.BoolValue(util.IsUnset(request.ParametersShrink)) {
		query["Parameters"] = request.ParametersShrink
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceId)) {
		query["ServiceId"] = request.ServiceId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceVersion)) {
		query["ServiceVersion"] = request.ServiceVersion
	}

	if !tea.BoolValue(util.IsUnset(request.SpecificationCode)) {
		query["SpecificationCode"] = request.SpecificationCode
	}

	if !tea.BoolValue(util.IsUnset(request.SpecificationName)) {
		query["SpecificationName"] = request.SpecificationName
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateName)) {
		query["TemplateName"] = request.TemplateName
	}

	if !tea.BoolValue(util.IsUnset(request.TrialType)) {
		query["TrialType"] = request.TrialType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateServiceInstance"),
		Version:     tea.String("2021-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateServiceInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - CreateServiceInstanceRequest
//
// @return CreateServiceInstanceResponse
func (client *Client) CreateServiceInstance(request *CreateServiceInstanceRequest) (_result *CreateServiceInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateServiceInstanceResponse{}
	_body, _err := client.CreateServiceInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除服务实例
//
// @param request - DeleteServiceInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteServiceInstancesResponse
func (client *Client) DeleteServiceInstancesWithOptions(request *DeleteServiceInstancesRequest, runtime *util.RuntimeOptions) (_result *DeleteServiceInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceInstanceId)) {
		query["ServiceInstanceId"] = request.ServiceInstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteServiceInstances"),
		Version:     tea.String("2021-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteServiceInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除服务实例
//
// @param request - DeleteServiceInstancesRequest
//
// @return DeleteServiceInstancesResponse
func (client *Client) DeleteServiceInstances(request *DeleteServiceInstancesRequest) (_result *DeleteServiceInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteServiceInstancesResponse{}
	_body, _err := client.DeleteServiceInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询服务实例详情
//
// @param request - GetServiceInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetServiceInstanceResponse
func (client *Client) GetServiceInstanceWithOptions(request *GetServiceInstanceRequest, runtime *util.RuntimeOptions) (_result *GetServiceInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MarketInstanceId)) {
		query["MarketInstanceId"] = request.MarketInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceInstanceId)) {
		query["ServiceInstanceId"] = request.ServiceInstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetServiceInstance"),
		Version:     tea.String("2021-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetServiceInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询服务实例详情
//
// @param request - GetServiceInstanceRequest
//
// @return GetServiceInstanceResponse
func (client *Client) GetServiceInstance(request *GetServiceInstanceRequest) (_result *GetServiceInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetServiceInstanceResponse{}
	_body, _err := client.GetServiceInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取ROS参数限制
//
// @param request - GetServiceTemplateParameterConstraintsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetServiceTemplateParameterConstraintsResponse
func (client *Client) GetServiceTemplateParameterConstraintsWithOptions(request *GetServiceTemplateParameterConstraintsRequest, runtime *util.RuntimeOptions) (_result *GetServiceTemplateParameterConstraintsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DeployRegionId)) {
		query["DeployRegionId"] = request.DeployRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.EnablePrivateVpcConnection)) {
		query["EnablePrivateVpcConnection"] = request.EnablePrivateVpcConnection
	}

	if !tea.BoolValue(util.IsUnset(request.Parameters)) {
		query["Parameters"] = request.Parameters
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceId)) {
		query["ServiceId"] = request.ServiceId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceInstanceId)) {
		query["ServiceInstanceId"] = request.ServiceInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceVersion)) {
		query["ServiceVersion"] = request.ServiceVersion
	}

	if !tea.BoolValue(util.IsUnset(request.SpecificationName)) {
		query["SpecificationName"] = request.SpecificationName
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateName)) {
		query["TemplateName"] = request.TemplateName
	}

	if !tea.BoolValue(util.IsUnset(request.TrialType)) {
		query["TrialType"] = request.TrialType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetServiceTemplateParameterConstraints"),
		Version:     tea.String("2021-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetServiceTemplateParameterConstraintsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取ROS参数限制
//
// @param request - GetServiceTemplateParameterConstraintsRequest
//
// @return GetServiceTemplateParameterConstraintsResponse
func (client *Client) GetServiceTemplateParameterConstraints(request *GetServiceTemplateParameterConstraintsRequest) (_result *GetServiceTemplateParameterConstraintsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetServiceTemplateParameterConstraintsResponse{}
	_body, _err := client.GetServiceTemplateParameterConstraintsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询服务实例部署升级以及应用日志
//
// @param request - ListServiceInstanceLogsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListServiceInstanceLogsResponse
func (client *Client) ListServiceInstanceLogsWithOptions(request *ListServiceInstanceLogsRequest, runtime *util.RuntimeOptions) (_result *ListServiceInstanceLogsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.LogSource)) {
		query["LogSource"] = request.LogSource
	}

	if !tea.BoolValue(util.IsUnset(request.Logstore)) {
		query["Logstore"] = request.Logstore
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceInstanceId)) {
		query["ServiceInstanceId"] = request.ServiceInstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListServiceInstanceLogs"),
		Version:     tea.String("2021-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListServiceInstanceLogsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询服务实例部署升级以及应用日志
//
// @param request - ListServiceInstanceLogsRequest
//
// @return ListServiceInstanceLogsResponse
func (client *Client) ListServiceInstanceLogs(request *ListServiceInstanceLogsRequest) (_result *ListServiceInstanceLogsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListServiceInstanceLogsResponse{}
	_body, _err := client.ListServiceInstanceLogsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 展示实例资源
//
// @param request - ListServiceInstanceResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListServiceInstanceResourcesResponse
func (client *Client) ListServiceInstanceResourcesWithOptions(request *ListServiceInstanceResourcesRequest, runtime *util.RuntimeOptions) (_result *ListServiceInstanceResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ExpireTimeEnd)) {
		query["ExpireTimeEnd"] = request.ExpireTimeEnd
	}

	if !tea.BoolValue(util.IsUnset(request.ExpireTimeStart)) {
		query["ExpireTimeStart"] = request.ExpireTimeStart
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.PayType)) {
		query["PayType"] = request.PayType
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceARN)) {
		query["ResourceARN"] = request.ResourceARN
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceInstanceId)) {
		query["ServiceInstanceId"] = request.ServiceInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceInstanceResourceType)) {
		query["ServiceInstanceResourceType"] = request.ServiceInstanceResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListServiceInstanceResources"),
		Version:     tea.String("2021-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListServiceInstanceResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 展示实例资源
//
// @param request - ListServiceInstanceResourcesRequest
//
// @return ListServiceInstanceResourcesResponse
func (client *Client) ListServiceInstanceResources(request *ListServiceInstanceResourcesRequest) (_result *ListServiceInstanceResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListServiceInstanceResourcesResponse{}
	_body, _err := client.ListServiceInstanceResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ListServiceInstancesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListServiceInstancesResponse
func (client *Client) ListServiceInstancesWithOptions(request *ListServiceInstancesRequest, runtime *util.RuntimeOptions) (_result *ListServiceInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Filter)) {
		query["Filter"] = request.Filter
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListServiceInstances"),
		Version:     tea.String("2021-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListServiceInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - ListServiceInstancesRequest
//
// @return ListServiceInstancesResponse
func (client *Client) ListServiceInstances(request *ListServiceInstancesRequest) (_result *ListServiceInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListServiceInstancesResponse{}
	_body, _err := client.ListServiceInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 重启服务实例
//
// @param request - RestartServiceInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RestartServiceInstanceResponse
func (client *Client) RestartServiceInstanceWithOptions(request *RestartServiceInstanceRequest, runtime *util.RuntimeOptions) (_result *RestartServiceInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceInstanceId)) {
		query["ServiceInstanceId"] = request.ServiceInstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RestartServiceInstance"),
		Version:     tea.String("2021-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RestartServiceInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 重启服务实例
//
// @param request - RestartServiceInstanceRequest
//
// @return RestartServiceInstanceResponse
func (client *Client) RestartServiceInstance(request *RestartServiceInstanceRequest) (_result *RestartServiceInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RestartServiceInstanceResponse{}
	_body, _err := client.RestartServiceInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 启动服务实例
//
// @param request - StartServiceInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartServiceInstanceResponse
func (client *Client) StartServiceInstanceWithOptions(request *StartServiceInstanceRequest, runtime *util.RuntimeOptions) (_result *StartServiceInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceInstanceId)) {
		query["ServiceInstanceId"] = request.ServiceInstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("StartServiceInstance"),
		Version:     tea.String("2021-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &StartServiceInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 启动服务实例
//
// @param request - StartServiceInstanceRequest
//
// @return StartServiceInstanceResponse
func (client *Client) StartServiceInstance(request *StartServiceInstanceRequest) (_result *StartServiceInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StartServiceInstanceResponse{}
	_body, _err := client.StartServiceInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 停止服务实例
//
// @param request - StopServiceInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopServiceInstanceResponse
func (client *Client) StopServiceInstanceWithOptions(request *StopServiceInstanceRequest, runtime *util.RuntimeOptions) (_result *StopServiceInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceInstanceId)) {
		query["ServiceInstanceId"] = request.ServiceInstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("StopServiceInstance"),
		Version:     tea.String("2021-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &StopServiceInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 停止服务实例
//
// @param request - StopServiceInstanceRequest
//
// @return StopServiceInstanceResponse
func (client *Client) StopServiceInstance(request *StopServiceInstanceRequest) (_result *StopServiceInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StopServiceInstanceResponse{}
	_body, _err := client.StopServiceInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param tmpReq - UpdateServiceInstanceSpecRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateServiceInstanceSpecResponse
func (client *Client) UpdateServiceInstanceSpecWithOptions(tmpReq *UpdateServiceInstanceSpecRequest, runtime *util.RuntimeOptions) (_result *UpdateServiceInstanceSpecResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &UpdateServiceInstanceSpecShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Parameters)) {
		request.ParametersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Parameters, tea.String("Parameters"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Commodity)) {
		query["Commodity"] = request.Commodity
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.EnableUserPrometheus)) {
		query["EnableUserPrometheus"] = request.EnableUserPrometheus
	}

	if !tea.BoolValue(util.IsUnset(request.OperationName)) {
		query["OperationName"] = request.OperationName
	}

	if !tea.BoolValue(util.IsUnset(request.ParametersShrink)) {
		query["Parameters"] = request.ParametersShrink
	}

	if !tea.BoolValue(util.IsUnset(request.PredefinedParametersName)) {
		query["PredefinedParametersName"] = request.PredefinedParametersName
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceInstanceId)) {
		query["ServiceInstanceId"] = request.ServiceInstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateServiceInstanceSpec"),
		Version:     tea.String("2021-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateServiceInstanceSpecResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - UpdateServiceInstanceSpecRequest
//
// @return UpdateServiceInstanceSpecResponse
func (client *Client) UpdateServiceInstanceSpec(request *UpdateServiceInstanceSpecRequest) (_result *UpdateServiceInstanceSpecResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateServiceInstanceSpecResponse{}
	_body, _err := client.UpdateServiceInstanceSpecWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
