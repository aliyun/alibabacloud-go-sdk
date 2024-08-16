// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type CancelExecutionRequest struct {
	// The ID of the execution.
	//
	// This parameter is required.
	//
	// example:
	//
	// exec-xxx
	ExecutionId *string `json:"ExecutionId,omitempty" xml:"ExecutionId,omitempty"`
	// The ID of the region.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CancelExecutionRequest) String() string {
	return tea.Prettify(s)
}

func (s CancelExecutionRequest) GoString() string {
	return s.String()
}

func (s *CancelExecutionRequest) SetExecutionId(v string) *CancelExecutionRequest {
	s.ExecutionId = &v
	return s
}

func (s *CancelExecutionRequest) SetRegionId(v string) *CancelExecutionRequest {
	s.RegionId = &v
	return s
}

type CancelExecutionResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 14A07460-EBE7-47CA-9757-12CC4761D47A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CancelExecutionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CancelExecutionResponseBody) GoString() string {
	return s.String()
}

func (s *CancelExecutionResponseBody) SetRequestId(v string) *CancelExecutionResponseBody {
	s.RequestId = &v
	return s
}

type CancelExecutionResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CancelExecutionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CancelExecutionResponse) String() string {
	return tea.Prettify(s)
}

func (s CancelExecutionResponse) GoString() string {
	return s.String()
}

func (s *CancelExecutionResponse) SetHeaders(v map[string]*string) *CancelExecutionResponse {
	s.Headers = v
	return s
}

func (s *CancelExecutionResponse) SetStatusCode(v int32) *CancelExecutionResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelExecutionResponse) SetBody(v *CancelExecutionResponseBody) *CancelExecutionResponse {
	s.Body = v
	return s
}

type ChangeResourceGroupRequest struct {
	// The ID of the resource group to which the cloud resource is to be moved. You can use resource groups to manage resources owned by your Alibaba Cloud account. Resource groups simplify the resource and permission management of your Alibaba Cloud account. For more information, see [What is Resource Management?](https://help.aliyun.com/document_detail/94475.html)
	//
	// This parameter is required.
	//
	// example:
	//
	// rg-acfm3peow3k****
	NewResourceGroupId *string `json:"NewResourceGroupId,omitempty" xml:"NewResourceGroupId,omitempty"`
	// The ID of the region.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the cloud resource that you want to move to another resource group.
	//
	// 	- If the ResourceType parameter is set to template, set the ResourceId parameter to the name of the template.
	//
	// 	- If the ResourceType parameter is set to parameter, set the ResourceId parameter to the name of the parameter.
	//
	// 	- If the ResourceType parameter is set to secretparameter, set the ResourceId parameter to the name of the encryption parameter.
	//
	// 	- If the ResourceType parameter is set to stateconfiguration, set the ResourceId parameter to the ID of the desired-state configuration.
	//
	// 	- If the ResourceType parameter is set to application, set the ResourceId parameter to the name of the application.
	//
	// This parameter is required.
	//
	// example:
	//
	// TemplateName
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The type of the cloud resource. Valid values:
	//
	// 	- template: template
	//
	// 	- parameter: parameter
	//
	// 	- secretparameter: encryption parameter
	//
	// 	- stateconfiguration: desired-state configuration
	//
	// 	- application: application
	//
	// This parameter is required.
	//
	// example:
	//
	// template
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
	// The ID of the request.
	//
	// example:
	//
	// 0620E49F-B884-5F98-A834-69D72922E5CF
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

type ContinueDeployApplicationGroupRequest struct {
	// The name of the application.
	//
	// This parameter is required.
	//
	// example:
	//
	// MyApplication
	ApplicationName *string `json:"ApplicationName,omitempty" xml:"ApplicationName,omitempty"`
	// The deployment information about the application group.
	//
	// This parameter is required.
	//
	// example:
	//
	// {       "TemplateURL": "https://ros-template.oss-cn-zhangjiakou.aliyuncs.com/App_Management_Existing_Vpc_Ecs_Instance.json",       "Parameters": {         "ZoneId": "cn-hangzhou-k",         "ProjectName": "test",         "SystemDiskSize": 40,         "InstanceChargeType": "PostPaid",         "SecurityGroupId": "sg-bp1a4374akk63jl8tddy",         "VSwitchId": "vsw-bp1fcvc3zn0jrag86rrlm",         "SystemDiskCategory": "cloud_essd",         "InstancePassword": "******",         "InternetChargeType": "PayByTraffic",         "InstanceCount": 1,         "InternetMaxBandwidthOut": 0,         "VpcId": "vpc-bp1i99boyas8i8m9t3skp",         "EcsImageId": "centos_8_5_x64_20G_alibase_20211228.vhd",         "DataDiskSize": 100,         "EcsInstanceType": "ecs.s6-c1m4.small",         "DataDiskCategory": "cloud_efficiency",         "EnvironmentCommandId": "c-hz028fc3g031gcg"       }
	DeployParameters *string `json:"DeployParameters,omitempty" xml:"DeployParameters,omitempty"`
	// The name of the application group.
	//
	// This parameter is required.
	//
	// example:
	//
	// MyApplicationGroup
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the region.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ContinueDeployApplicationGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s ContinueDeployApplicationGroupRequest) GoString() string {
	return s.String()
}

func (s *ContinueDeployApplicationGroupRequest) SetApplicationName(v string) *ContinueDeployApplicationGroupRequest {
	s.ApplicationName = &v
	return s
}

func (s *ContinueDeployApplicationGroupRequest) SetDeployParameters(v string) *ContinueDeployApplicationGroupRequest {
	s.DeployParameters = &v
	return s
}

func (s *ContinueDeployApplicationGroupRequest) SetName(v string) *ContinueDeployApplicationGroupRequest {
	s.Name = &v
	return s
}

func (s *ContinueDeployApplicationGroupRequest) SetRegionId(v string) *ContinueDeployApplicationGroupRequest {
	s.RegionId = &v
	return s
}

type ContinueDeployApplicationGroupResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 8AF4800A-A316-589A-90C4-313B1FEEB084
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ContinueDeployApplicationGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ContinueDeployApplicationGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ContinueDeployApplicationGroupResponseBody) SetRequestId(v string) *ContinueDeployApplicationGroupResponseBody {
	s.RequestId = &v
	return s
}

type ContinueDeployApplicationGroupResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ContinueDeployApplicationGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ContinueDeployApplicationGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s ContinueDeployApplicationGroupResponse) GoString() string {
	return s.String()
}

func (s *ContinueDeployApplicationGroupResponse) SetHeaders(v map[string]*string) *ContinueDeployApplicationGroupResponse {
	s.Headers = v
	return s
}

func (s *ContinueDeployApplicationGroupResponse) SetStatusCode(v int32) *ContinueDeployApplicationGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *ContinueDeployApplicationGroupResponse) SetBody(v *ContinueDeployApplicationGroupResponseBody) *ContinueDeployApplicationGroupResponse {
	s.Body = v
	return s
}

type CreateApplicationRequest struct {
	// The configurations of application alerts.
	AlarmConfig *CreateApplicationRequestAlarmConfig `json:"AlarmConfig,omitempty" xml:"AlarmConfig,omitempty" type:"Struct"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// example:
	//
	// TF-CreateApplication-1647587475-84104b89-eba5-47a8-b2fd-807b8b7d
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The description of the application.
	//
	// example:
	//
	// application
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The application name.
	//
	// This parameter is required.
	//
	// example:
	//
	// MyApplication
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The region ID. Set the value to cn-hangzhou.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfmxsn4m******
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The ID of the service.
	//
	// example:
	//
	// service-79538e30e44541b699d8
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// The tags.
	//
	// example:
	//
	// {"k1":"v1","k2":"v2"}
	Tags map[string]interface{} `json:"Tags,omitempty" xml:"Tags,omitempty"`
}

func (s CreateApplicationRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateApplicationRequest) GoString() string {
	return s.String()
}

func (s *CreateApplicationRequest) SetAlarmConfig(v *CreateApplicationRequestAlarmConfig) *CreateApplicationRequest {
	s.AlarmConfig = v
	return s
}

func (s *CreateApplicationRequest) SetClientToken(v string) *CreateApplicationRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateApplicationRequest) SetDescription(v string) *CreateApplicationRequest {
	s.Description = &v
	return s
}

func (s *CreateApplicationRequest) SetName(v string) *CreateApplicationRequest {
	s.Name = &v
	return s
}

func (s *CreateApplicationRequest) SetRegionId(v string) *CreateApplicationRequest {
	s.RegionId = &v
	return s
}

func (s *CreateApplicationRequest) SetResourceGroupId(v string) *CreateApplicationRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateApplicationRequest) SetServiceId(v string) *CreateApplicationRequest {
	s.ServiceId = &v
	return s
}

func (s *CreateApplicationRequest) SetTags(v map[string]interface{}) *CreateApplicationRequest {
	s.Tags = v
	return s
}

type CreateApplicationRequestAlarmConfig struct {
	// The alert contact groups.
	ContactGroups []*string `json:"ContactGroups,omitempty" xml:"ContactGroups,omitempty" type:"Repeated"`
	// The health check URL of the application.
	//
	// example:
	//
	// /healthcheck/tcp50122
	HealthCheckUrl *string `json:"HealthCheckUrl,omitempty" xml:"HealthCheckUrl,omitempty"`
	// The alert templates.
	TemplateIds []*string `json:"TemplateIds,omitempty" xml:"TemplateIds,omitempty" type:"Repeated"`
}

func (s CreateApplicationRequestAlarmConfig) String() string {
	return tea.Prettify(s)
}

func (s CreateApplicationRequestAlarmConfig) GoString() string {
	return s.String()
}

func (s *CreateApplicationRequestAlarmConfig) SetContactGroups(v []*string) *CreateApplicationRequestAlarmConfig {
	s.ContactGroups = v
	return s
}

func (s *CreateApplicationRequestAlarmConfig) SetHealthCheckUrl(v string) *CreateApplicationRequestAlarmConfig {
	s.HealthCheckUrl = &v
	return s
}

func (s *CreateApplicationRequestAlarmConfig) SetTemplateIds(v []*string) *CreateApplicationRequestAlarmConfig {
	s.TemplateIds = v
	return s
}

type CreateApplicationShrinkRequest struct {
	// The configurations of application alerts.
	AlarmConfigShrink *string `json:"AlarmConfig,omitempty" xml:"AlarmConfig,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// example:
	//
	// TF-CreateApplication-1647587475-84104b89-eba5-47a8-b2fd-807b8b7d
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The description of the application.
	//
	// example:
	//
	// application
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The application name.
	//
	// This parameter is required.
	//
	// example:
	//
	// MyApplication
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The region ID. Set the value to cn-hangzhou.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfmxsn4m******
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The ID of the service.
	//
	// example:
	//
	// service-79538e30e44541b699d8
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// The tags.
	//
	// example:
	//
	// {"k1":"v1","k2":"v2"}
	TagsShrink *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
}

func (s CreateApplicationShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateApplicationShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateApplicationShrinkRequest) SetAlarmConfigShrink(v string) *CreateApplicationShrinkRequest {
	s.AlarmConfigShrink = &v
	return s
}

func (s *CreateApplicationShrinkRequest) SetClientToken(v string) *CreateApplicationShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateApplicationShrinkRequest) SetDescription(v string) *CreateApplicationShrinkRequest {
	s.Description = &v
	return s
}

func (s *CreateApplicationShrinkRequest) SetName(v string) *CreateApplicationShrinkRequest {
	s.Name = &v
	return s
}

func (s *CreateApplicationShrinkRequest) SetRegionId(v string) *CreateApplicationShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *CreateApplicationShrinkRequest) SetResourceGroupId(v string) *CreateApplicationShrinkRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateApplicationShrinkRequest) SetServiceId(v string) *CreateApplicationShrinkRequest {
	s.ServiceId = &v
	return s
}

func (s *CreateApplicationShrinkRequest) SetTagsShrink(v string) *CreateApplicationShrinkRequest {
	s.TagsShrink = &v
	return s
}

type CreateApplicationResponseBody struct {
	// The information about the application.
	Application *CreateApplicationResponseBodyApplication `json:"Application,omitempty" xml:"Application,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 274917E8-8E74-5928-A82F-4940F52F7ACB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateApplicationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateApplicationResponseBody) GoString() string {
	return s.String()
}

func (s *CreateApplicationResponseBody) SetApplication(v *CreateApplicationResponseBodyApplication) *CreateApplicationResponseBody {
	s.Application = v
	return s
}

func (s *CreateApplicationResponseBody) SetRequestId(v string) *CreateApplicationResponseBody {
	s.RequestId = &v
	return s
}

type CreateApplicationResponseBodyApplication struct {
	// The time when the application was created.
	//
	// example:
	//
	// 2021-09-07T09:17:46Z
	CreateDate *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	// The description of the application.
	//
	// example:
	//
	// application
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The application name.
	//
	// example:
	//
	// Myapplication
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The tags.
	//
	// example:
	//
	// {"k1":"v1","k2":"v2"}
	Tags map[string]*string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The time when the application was updated.
	//
	// example:
	//
	// 2021-09-07T09:17:46Z
	UpdateDate *string `json:"UpdateDate,omitempty" xml:"UpdateDate,omitempty"`
}

func (s CreateApplicationResponseBodyApplication) String() string {
	return tea.Prettify(s)
}

func (s CreateApplicationResponseBodyApplication) GoString() string {
	return s.String()
}

func (s *CreateApplicationResponseBodyApplication) SetCreateDate(v string) *CreateApplicationResponseBodyApplication {
	s.CreateDate = &v
	return s
}

func (s *CreateApplicationResponseBodyApplication) SetDescription(v string) *CreateApplicationResponseBodyApplication {
	s.Description = &v
	return s
}

func (s *CreateApplicationResponseBodyApplication) SetName(v string) *CreateApplicationResponseBodyApplication {
	s.Name = &v
	return s
}

func (s *CreateApplicationResponseBodyApplication) SetTags(v map[string]*string) *CreateApplicationResponseBodyApplication {
	s.Tags = v
	return s
}

func (s *CreateApplicationResponseBodyApplication) SetUpdateDate(v string) *CreateApplicationResponseBodyApplication {
	s.UpdateDate = &v
	return s
}

type CreateApplicationResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateApplicationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateApplicationResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateApplicationResponse) GoString() string {
	return s.String()
}

func (s *CreateApplicationResponse) SetHeaders(v map[string]*string) *CreateApplicationResponse {
	s.Headers = v
	return s
}

func (s *CreateApplicationResponse) SetStatusCode(v int32) *CreateApplicationResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateApplicationResponse) SetBody(v *CreateApplicationResponseBody) *CreateApplicationResponse {
	s.Body = v
	return s
}

type CreateApplicationGroupRequest struct {
	// The application name.
	//
	// This parameter is required.
	//
	// example:
	//
	// MyApplication
	ApplicationName *string `json:"ApplicationName,omitempty" xml:"ApplicationName,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// example:
	//
	// -
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the application group in CloudMonitor.
	//
	// example:
	//
	// 218026174
	CmsGroupId *string `json:"CmsGroupId,omitempty" xml:"CmsGroupId,omitempty"`
	// The ID of the region in which the related sources reside.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	DeployRegionId *string `json:"DeployRegionId,omitempty" xml:"DeployRegionId,omitempty"`
	// The description of the application group.
	//
	// example:
	//
	// ApplicationGroup
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The key of the tag. You must set both the ImportTagKey and the ImportTagValue parameters, or leave both of them empty. If you do not set the ImportTagKey and ImportTagValue parameters, the application name is used for this parameter by default.
	//
	// example:
	//
	// k1
	ImportTagKey *string `json:"ImportTagKey,omitempty" xml:"ImportTagKey,omitempty"`
	// The value of the tag. You must set both the ImportTagKey and the ImportTagValue parameters, or leave both of them empty. If you do not set the ImportTagKey and ImportTagValue parameters, the application group name is used for this parameter by default.
	//
	// example:
	//
	// v1
	ImportTagValue *string `json:"ImportTagValue,omitempty" xml:"ImportTagValue,omitempty"`
	// The name of the application group.
	//
	// This parameter is required.
	//
	// example:
	//
	// MyApplicationGroup
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The region ID. Set the value to cn-hangzhou.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateApplicationGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateApplicationGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateApplicationGroupRequest) SetApplicationName(v string) *CreateApplicationGroupRequest {
	s.ApplicationName = &v
	return s
}

func (s *CreateApplicationGroupRequest) SetClientToken(v string) *CreateApplicationGroupRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateApplicationGroupRequest) SetCmsGroupId(v string) *CreateApplicationGroupRequest {
	s.CmsGroupId = &v
	return s
}

func (s *CreateApplicationGroupRequest) SetDeployRegionId(v string) *CreateApplicationGroupRequest {
	s.DeployRegionId = &v
	return s
}

func (s *CreateApplicationGroupRequest) SetDescription(v string) *CreateApplicationGroupRequest {
	s.Description = &v
	return s
}

func (s *CreateApplicationGroupRequest) SetImportTagKey(v string) *CreateApplicationGroupRequest {
	s.ImportTagKey = &v
	return s
}

func (s *CreateApplicationGroupRequest) SetImportTagValue(v string) *CreateApplicationGroupRequest {
	s.ImportTagValue = &v
	return s
}

func (s *CreateApplicationGroupRequest) SetName(v string) *CreateApplicationGroupRequest {
	s.Name = &v
	return s
}

func (s *CreateApplicationGroupRequest) SetRegionId(v string) *CreateApplicationGroupRequest {
	s.RegionId = &v
	return s
}

type CreateApplicationGroupResponseBody struct {
	// The information about the application group.
	ApplicationGroup *CreateApplicationGroupResponseBodyApplicationGroup `json:"ApplicationGroup,omitempty" xml:"ApplicationGroup,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 0E6BEBD3-7F9E-5878-834B-097633AB5F33
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateApplicationGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateApplicationGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateApplicationGroupResponseBody) SetApplicationGroup(v *CreateApplicationGroupResponseBodyApplicationGroup) *CreateApplicationGroupResponseBody {
	s.ApplicationGroup = v
	return s
}

func (s *CreateApplicationGroupResponseBody) SetRequestId(v string) *CreateApplicationGroupResponseBody {
	s.RequestId = &v
	return s
}

type CreateApplicationGroupResponseBodyApplicationGroup struct {
	// The application name.
	//
	// example:
	//
	// MyApplication
	ApplicationName *string `json:"ApplicationName,omitempty" xml:"ApplicationName,omitempty"`
	// The ID of the application group in CloudMonitor.
	//
	// example:
	//
	// 1245768
	CmsGroupId *string `json:"CmsGroupId,omitempty" xml:"CmsGroupId,omitempty"`
	// The time when the application group was created.
	//
	// example:
	//
	// 2021-09-07T10:28:25Z
	CreateDate *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	// The ID of the region in which the related sources reside.
	//
	// example:
	//
	// cn-hangzhou
	DeployRegionId *string `json:"DeployRegionId,omitempty" xml:"DeployRegionId,omitempty"`
	// The description of the application group.
	//
	// example:
	//
	// ApplicationGroup
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The key of the tag.
	//
	// example:
	//
	// k1
	ImportTagKey *string `json:"ImportTagKey,omitempty" xml:"ImportTagKey,omitempty"`
	// The value of the tag.
	//
	// example:
	//
	// v1
	ImportTagValue *string `json:"ImportTagValue,omitempty" xml:"ImportTagValue,omitempty"`
	// The name of the application group.
	//
	// example:
	//
	// MyApplicationGroup
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The time when the application group was updated.
	//
	// example:
	//
	// 2021-09-07T10:28:25Z
	UpdateDate *string `json:"UpdateDate,omitempty" xml:"UpdateDate,omitempty"`
}

func (s CreateApplicationGroupResponseBodyApplicationGroup) String() string {
	return tea.Prettify(s)
}

func (s CreateApplicationGroupResponseBodyApplicationGroup) GoString() string {
	return s.String()
}

func (s *CreateApplicationGroupResponseBodyApplicationGroup) SetApplicationName(v string) *CreateApplicationGroupResponseBodyApplicationGroup {
	s.ApplicationName = &v
	return s
}

func (s *CreateApplicationGroupResponseBodyApplicationGroup) SetCmsGroupId(v string) *CreateApplicationGroupResponseBodyApplicationGroup {
	s.CmsGroupId = &v
	return s
}

func (s *CreateApplicationGroupResponseBodyApplicationGroup) SetCreateDate(v string) *CreateApplicationGroupResponseBodyApplicationGroup {
	s.CreateDate = &v
	return s
}

func (s *CreateApplicationGroupResponseBodyApplicationGroup) SetDeployRegionId(v string) *CreateApplicationGroupResponseBodyApplicationGroup {
	s.DeployRegionId = &v
	return s
}

func (s *CreateApplicationGroupResponseBodyApplicationGroup) SetDescription(v string) *CreateApplicationGroupResponseBodyApplicationGroup {
	s.Description = &v
	return s
}

func (s *CreateApplicationGroupResponseBodyApplicationGroup) SetImportTagKey(v string) *CreateApplicationGroupResponseBodyApplicationGroup {
	s.ImportTagKey = &v
	return s
}

func (s *CreateApplicationGroupResponseBodyApplicationGroup) SetImportTagValue(v string) *CreateApplicationGroupResponseBodyApplicationGroup {
	s.ImportTagValue = &v
	return s
}

func (s *CreateApplicationGroupResponseBodyApplicationGroup) SetName(v string) *CreateApplicationGroupResponseBodyApplicationGroup {
	s.Name = &v
	return s
}

func (s *CreateApplicationGroupResponseBodyApplicationGroup) SetUpdateDate(v string) *CreateApplicationGroupResponseBodyApplicationGroup {
	s.UpdateDate = &v
	return s
}

type CreateApplicationGroupResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateApplicationGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateApplicationGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateApplicationGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateApplicationGroupResponse) SetHeaders(v map[string]*string) *CreateApplicationGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateApplicationGroupResponse) SetStatusCode(v int32) *CreateApplicationGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateApplicationGroupResponse) SetBody(v *CreateApplicationGroupResponseBody) *CreateApplicationGroupResponse {
	s.Body = v
	return s
}

type CreateOpsItemRequest struct {
	// The category.
	//
	// Valid values:
	//
	// 	- Availability
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- Performance
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- Security
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- Cost
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- Recovery
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// This parameter is required.
	//
	// example:
	//
	// Security
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// example:
	//
	// 123e56767-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The string to be deduplicated.
	//
	// example:
	//
	// ecs_instance_Sys
	DedupString *string `json:"DedupString,omitempty" xml:"DedupString,omitempty"`
	// The description of the operation.
	//
	// example:
	//
	// OpsItem
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The priority. Valid values: 1 to 5. 1 indicates the highest priority.
	//
	// example:
	//
	// 4
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfmxsn4m4******
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The Alibaba Cloud Resource Names (ARNs) of the associated resources.
	//
	// example:
	//
	// [\\"acs:oos:cn-hangzhou:1563457855438522:application/test-33333/applicationgroup/default-cn-hangzhou-1\\"]
	Resources *string `json:"Resources,omitempty" xml:"Resources,omitempty"`
	// The severity level.
	//
	// Valid values:
	//
	// 	- High
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- Low
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- Medium
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- Critical
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// This parameter is required.
	//
	// example:
	//
	// Medium
	Severity *string `json:"Severity,omitempty" xml:"Severity,omitempty"`
	// The solutions.
	//
	// example:
	//
	// [{\\n \\\\"priority\\\\":3,\\n \\\\"type\\\\":\\\\"SingleAZEcs\\\\",\\n \\\\"url\\\\":\\\\"http://ecs.consle.aliyuncs.com\\\\",\\n \\\\"description\\\\":\\\\"Create Elastic Compute Service (ECS) instances in different zones and import them to an application group.\\\\"\\n}]
	Solutions *string `json:"Solutions,omitempty" xml:"Solutions,omitempty"`
	// The source business.
	//
	// This parameter is required.
	//
	// example:
	//
	// /aliyun/ecs
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The tags.
	//
	// example:
	//
	// {
	//
	//       "k1": "v1",
	//
	//       "k2": "v2"
	//
	// }
	Tags map[string]interface{} `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The title of the O\\&M item.
	//
	// This parameter is required.
	//
	// example:
	//
	// ECS reboot is scheduled
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s CreateOpsItemRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateOpsItemRequest) GoString() string {
	return s.String()
}

func (s *CreateOpsItemRequest) SetCategory(v string) *CreateOpsItemRequest {
	s.Category = &v
	return s
}

func (s *CreateOpsItemRequest) SetClientToken(v string) *CreateOpsItemRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateOpsItemRequest) SetDedupString(v string) *CreateOpsItemRequest {
	s.DedupString = &v
	return s
}

func (s *CreateOpsItemRequest) SetDescription(v string) *CreateOpsItemRequest {
	s.Description = &v
	return s
}

func (s *CreateOpsItemRequest) SetPriority(v int32) *CreateOpsItemRequest {
	s.Priority = &v
	return s
}

func (s *CreateOpsItemRequest) SetRegionId(v string) *CreateOpsItemRequest {
	s.RegionId = &v
	return s
}

func (s *CreateOpsItemRequest) SetResourceGroupId(v string) *CreateOpsItemRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateOpsItemRequest) SetResources(v string) *CreateOpsItemRequest {
	s.Resources = &v
	return s
}

func (s *CreateOpsItemRequest) SetSeverity(v string) *CreateOpsItemRequest {
	s.Severity = &v
	return s
}

func (s *CreateOpsItemRequest) SetSolutions(v string) *CreateOpsItemRequest {
	s.Solutions = &v
	return s
}

func (s *CreateOpsItemRequest) SetSource(v string) *CreateOpsItemRequest {
	s.Source = &v
	return s
}

func (s *CreateOpsItemRequest) SetTags(v map[string]interface{}) *CreateOpsItemRequest {
	s.Tags = v
	return s
}

func (s *CreateOpsItemRequest) SetTitle(v string) *CreateOpsItemRequest {
	s.Title = &v
	return s
}

type CreateOpsItemShrinkRequest struct {
	// The category.
	//
	// Valid values:
	//
	// 	- Availability
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- Performance
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- Security
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- Cost
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- Recovery
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// This parameter is required.
	//
	// example:
	//
	// Security
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// example:
	//
	// 123e56767-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The string to be deduplicated.
	//
	// example:
	//
	// ecs_instance_Sys
	DedupString *string `json:"DedupString,omitempty" xml:"DedupString,omitempty"`
	// The description of the operation.
	//
	// example:
	//
	// OpsItem
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The priority. Valid values: 1 to 5. 1 indicates the highest priority.
	//
	// example:
	//
	// 4
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfmxsn4m4******
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The Alibaba Cloud Resource Names (ARNs) of the associated resources.
	//
	// example:
	//
	// [\\"acs:oos:cn-hangzhou:1563457855438522:application/test-33333/applicationgroup/default-cn-hangzhou-1\\"]
	Resources *string `json:"Resources,omitempty" xml:"Resources,omitempty"`
	// The severity level.
	//
	// Valid values:
	//
	// 	- High
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- Low
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- Medium
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- Critical
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// This parameter is required.
	//
	// example:
	//
	// Medium
	Severity *string `json:"Severity,omitempty" xml:"Severity,omitempty"`
	// The solutions.
	//
	// example:
	//
	// [{\\n \\\\"priority\\\\":3,\\n \\\\"type\\\\":\\\\"SingleAZEcs\\\\",\\n \\\\"url\\\\":\\\\"http://ecs.consle.aliyuncs.com\\\\",\\n \\\\"description\\\\":\\\\"Create Elastic Compute Service (ECS) instances in different zones and import them to an application group.\\\\"\\n}]
	Solutions *string `json:"Solutions,omitempty" xml:"Solutions,omitempty"`
	// The source business.
	//
	// This parameter is required.
	//
	// example:
	//
	// /aliyun/ecs
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The tags.
	//
	// example:
	//
	// {
	//
	//       "k1": "v1",
	//
	//       "k2": "v2"
	//
	// }
	TagsShrink *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The title of the O\\&M item.
	//
	// This parameter is required.
	//
	// example:
	//
	// ECS reboot is scheduled
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s CreateOpsItemShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateOpsItemShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateOpsItemShrinkRequest) SetCategory(v string) *CreateOpsItemShrinkRequest {
	s.Category = &v
	return s
}

func (s *CreateOpsItemShrinkRequest) SetClientToken(v string) *CreateOpsItemShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateOpsItemShrinkRequest) SetDedupString(v string) *CreateOpsItemShrinkRequest {
	s.DedupString = &v
	return s
}

func (s *CreateOpsItemShrinkRequest) SetDescription(v string) *CreateOpsItemShrinkRequest {
	s.Description = &v
	return s
}

func (s *CreateOpsItemShrinkRequest) SetPriority(v int32) *CreateOpsItemShrinkRequest {
	s.Priority = &v
	return s
}

func (s *CreateOpsItemShrinkRequest) SetRegionId(v string) *CreateOpsItemShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *CreateOpsItemShrinkRequest) SetResourceGroupId(v string) *CreateOpsItemShrinkRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateOpsItemShrinkRequest) SetResources(v string) *CreateOpsItemShrinkRequest {
	s.Resources = &v
	return s
}

func (s *CreateOpsItemShrinkRequest) SetSeverity(v string) *CreateOpsItemShrinkRequest {
	s.Severity = &v
	return s
}

func (s *CreateOpsItemShrinkRequest) SetSolutions(v string) *CreateOpsItemShrinkRequest {
	s.Solutions = &v
	return s
}

func (s *CreateOpsItemShrinkRequest) SetSource(v string) *CreateOpsItemShrinkRequest {
	s.Source = &v
	return s
}

func (s *CreateOpsItemShrinkRequest) SetTagsShrink(v string) *CreateOpsItemShrinkRequest {
	s.TagsShrink = &v
	return s
}

func (s *CreateOpsItemShrinkRequest) SetTitle(v string) *CreateOpsItemShrinkRequest {
	s.Title = &v
	return s
}

type CreateOpsItemResponseBody struct {
	// The information about the O\\&M item.
	OpsItem *CreateOpsItemResponseBodyOpsItem `json:"OpsItem,omitempty" xml:"OpsItem,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// DA4F08D4-DA54-5407-84B9-108C71D75B17
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateOpsItemResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateOpsItemResponseBody) GoString() string {
	return s.String()
}

func (s *CreateOpsItemResponseBody) SetOpsItem(v *CreateOpsItemResponseBodyOpsItem) *CreateOpsItemResponseBody {
	s.OpsItem = v
	return s
}

func (s *CreateOpsItemResponseBody) SetRequestId(v string) *CreateOpsItemResponseBody {
	s.RequestId = &v
	return s
}

type CreateOpsItemResponseBodyOpsItem struct {
	// The attributes of the O\\&M item.
	//
	// example:
	//
	// {\\"regionId\\":\\"cn-zhangjiakou\\",\\"appId\\":\\"992BKO1X75GRQ4E8\\",\\"instanceId\\":\\"i-8vbcymxagqsyyyjppbr4\\",\\"instance_name\\":\\"cdae\\"}
	Attributes *string `json:"Attributes,omitempty" xml:"Attributes,omitempty"`
	// The category of the O\\&M item.
	//
	// example:
	//
	// Security
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The time when the O\\&M item was created.
	//
	// example:
	//
	// 2023-03-24T03:55Z
	CreateDate *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	// The user who created the O\\&M item.
	//
	// example:
	//
	// root(130900000)
	CreatedBy *string `json:"CreatedBy,omitempty" xml:"CreatedBy,omitempty"`
	// The description of the O\\&M item.
	//
	// example:
	//
	// OpsItem
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The user who last modified the O\\&M item.
	//
	// example:
	//
	// root(130900000)
	LastModifiedBy *string `json:"LastModifiedBy,omitempty" xml:"LastModifiedBy,omitempty"`
	// The ID of the O\\&M item.
	//
	// example:
	//
	// oi-dba9c6eec9254a4d87c1
	OpsItemId *string `json:"OpsItemId,omitempty" xml:"OpsItemId,omitempty"`
	// The priority of the O\\&M item.
	//
	// example:
	//
	// 2
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfmxsn4m4******
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The ARNs of the associated resources.
	//
	// example:
	//
	// [\\"acs:oos:cn-hangzhou:1563457855438522:application/dingTest/applicationgroup/fltest\\"]
	Resources *string `json:"Resources,omitempty" xml:"Resources,omitempty"`
	// The severity level of the O\\&M item.
	//
	// example:
	//
	// Medium
	Severity *string `json:"Severity,omitempty" xml:"Severity,omitempty"`
	// The solutions.
	//
	// example:
	//
	// [{\\n \\\\"priority\\\\":3,\\n \\\\"type\\\\":\\\\"url\\\\",\\n \\\\"url\\\\":\\\\"https://example..com\\\\",\\n \\\\"description\\\\":\\\\"Specify a cross-zone high availability cluster. \\\\"\\n}]
	Solutions *string `json:"Solutions,omitempty" xml:"Solutions,omitempty"`
	// The source business of the O\\&M item.
	//
	// example:
	//
	// /aliyun/ecs
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The state of the O\\&M item.
	//
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tags of the O\\&M item.
	//
	// example:
	//
	// {"k1": "v1"}
	Tags map[string]interface{} `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The title of the O\\&M item.
	//
	// example:
	//
	// ECS reboot is scheduled
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// The time when the O\\&M item was updated.
	//
	// example:
	//
	// 2023-03-24T03:55Z
	UpdateDate *string `json:"UpdateDate,omitempty" xml:"UpdateDate,omitempty"`
}

func (s CreateOpsItemResponseBodyOpsItem) String() string {
	return tea.Prettify(s)
}

func (s CreateOpsItemResponseBodyOpsItem) GoString() string {
	return s.String()
}

func (s *CreateOpsItemResponseBodyOpsItem) SetAttributes(v string) *CreateOpsItemResponseBodyOpsItem {
	s.Attributes = &v
	return s
}

func (s *CreateOpsItemResponseBodyOpsItem) SetCategory(v string) *CreateOpsItemResponseBodyOpsItem {
	s.Category = &v
	return s
}

func (s *CreateOpsItemResponseBodyOpsItem) SetCreateDate(v string) *CreateOpsItemResponseBodyOpsItem {
	s.CreateDate = &v
	return s
}

func (s *CreateOpsItemResponseBodyOpsItem) SetCreatedBy(v string) *CreateOpsItemResponseBodyOpsItem {
	s.CreatedBy = &v
	return s
}

func (s *CreateOpsItemResponseBodyOpsItem) SetDescription(v string) *CreateOpsItemResponseBodyOpsItem {
	s.Description = &v
	return s
}

func (s *CreateOpsItemResponseBodyOpsItem) SetLastModifiedBy(v string) *CreateOpsItemResponseBodyOpsItem {
	s.LastModifiedBy = &v
	return s
}

func (s *CreateOpsItemResponseBodyOpsItem) SetOpsItemId(v string) *CreateOpsItemResponseBodyOpsItem {
	s.OpsItemId = &v
	return s
}

func (s *CreateOpsItemResponseBodyOpsItem) SetPriority(v int32) *CreateOpsItemResponseBodyOpsItem {
	s.Priority = &v
	return s
}

func (s *CreateOpsItemResponseBodyOpsItem) SetResourceGroupId(v string) *CreateOpsItemResponseBodyOpsItem {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateOpsItemResponseBodyOpsItem) SetResources(v string) *CreateOpsItemResponseBodyOpsItem {
	s.Resources = &v
	return s
}

func (s *CreateOpsItemResponseBodyOpsItem) SetSeverity(v string) *CreateOpsItemResponseBodyOpsItem {
	s.Severity = &v
	return s
}

func (s *CreateOpsItemResponseBodyOpsItem) SetSolutions(v string) *CreateOpsItemResponseBodyOpsItem {
	s.Solutions = &v
	return s
}

func (s *CreateOpsItemResponseBodyOpsItem) SetSource(v string) *CreateOpsItemResponseBodyOpsItem {
	s.Source = &v
	return s
}

func (s *CreateOpsItemResponseBodyOpsItem) SetStatus(v string) *CreateOpsItemResponseBodyOpsItem {
	s.Status = &v
	return s
}

func (s *CreateOpsItemResponseBodyOpsItem) SetTags(v map[string]interface{}) *CreateOpsItemResponseBodyOpsItem {
	s.Tags = v
	return s
}

func (s *CreateOpsItemResponseBodyOpsItem) SetTitle(v string) *CreateOpsItemResponseBodyOpsItem {
	s.Title = &v
	return s
}

func (s *CreateOpsItemResponseBodyOpsItem) SetUpdateDate(v string) *CreateOpsItemResponseBodyOpsItem {
	s.UpdateDate = &v
	return s
}

type CreateOpsItemResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateOpsItemResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateOpsItemResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateOpsItemResponse) GoString() string {
	return s.String()
}

func (s *CreateOpsItemResponse) SetHeaders(v map[string]*string) *CreateOpsItemResponse {
	s.Headers = v
	return s
}

func (s *CreateOpsItemResponse) SetStatusCode(v int32) *CreateOpsItemResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateOpsItemResponse) SetBody(v *CreateOpsItemResponseBody) *CreateOpsItemResponse {
	s.Body = v
	return s
}

type CreateParameterRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can be up to 64 characters in length, and can contain letters, digits, hyphens (-), and underscores (_). For more information, see "How to ensure idempotence".
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-42665544****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The constraints of the common parameter. By default, this parameter is null. Valid values:
	//
	// 	- AllowedValues: The value that is allowed for the common parameter. It must be an array string.
	//
	// 	- AllowedPattern: The pattern that is allowed for the common parameter. It must be a regular expression.
	//
	// 	- MinLength: The minimum length of the common parameter.
	//
	// 	- MaxLength: The maximum length of the common parameter.
	//
	// example:
	//
	// {
	//
	//     "AllowedValues": [
	//
	//         "parameter"
	//
	//     ],
	//
	//     "AllowedPattern": "parameter",
	//
	//     "MinLength": 0,
	//
	//     "MaxLength": 20
	//
	// }
	Constraints *string `json:"Constraints,omitempty" xml:"Constraints,omitempty"`
	// The description of the common parameter. The description must be 1 to 200 characters in length.
	//
	// example:
	//
	// parameter
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the parameter. The name must be 1 to 200 characters in length, and can contain letters, digits, hyphens (-), and underscores (_). It cannot start with ALIYUN, ACS, ALIBABA, ALICLOUD, or OOS.
	//
	// This parameter is required.
	//
	// example:
	//
	// MyParameter
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the region.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfmxsn4m4******
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The tags.
	//
	// example:
	//
	// {"k1": "v1", "k2": "v2"}
	Tags map[string]interface{} `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The data type of the parameter. Valid values: String and StringList.
	//
	// This parameter is required.
	//
	// example:
	//
	// String
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The value of the common parameter. The value must be 1 to 4096 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// parameter
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateParameterRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateParameterRequest) GoString() string {
	return s.String()
}

func (s *CreateParameterRequest) SetClientToken(v string) *CreateParameterRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateParameterRequest) SetConstraints(v string) *CreateParameterRequest {
	s.Constraints = &v
	return s
}

func (s *CreateParameterRequest) SetDescription(v string) *CreateParameterRequest {
	s.Description = &v
	return s
}

func (s *CreateParameterRequest) SetName(v string) *CreateParameterRequest {
	s.Name = &v
	return s
}

func (s *CreateParameterRequest) SetRegionId(v string) *CreateParameterRequest {
	s.RegionId = &v
	return s
}

func (s *CreateParameterRequest) SetResourceGroupId(v string) *CreateParameterRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateParameterRequest) SetTags(v map[string]interface{}) *CreateParameterRequest {
	s.Tags = v
	return s
}

func (s *CreateParameterRequest) SetType(v string) *CreateParameterRequest {
	s.Type = &v
	return s
}

func (s *CreateParameterRequest) SetValue(v string) *CreateParameterRequest {
	s.Value = &v
	return s
}

type CreateParameterShrinkRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can be up to 64 characters in length, and can contain letters, digits, hyphens (-), and underscores (_). For more information, see "How to ensure idempotence".
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-42665544****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The constraints of the common parameter. By default, this parameter is null. Valid values:
	//
	// 	- AllowedValues: The value that is allowed for the common parameter. It must be an array string.
	//
	// 	- AllowedPattern: The pattern that is allowed for the common parameter. It must be a regular expression.
	//
	// 	- MinLength: The minimum length of the common parameter.
	//
	// 	- MaxLength: The maximum length of the common parameter.
	//
	// example:
	//
	// {
	//
	//     "AllowedValues": [
	//
	//         "parameter"
	//
	//     ],
	//
	//     "AllowedPattern": "parameter",
	//
	//     "MinLength": 0,
	//
	//     "MaxLength": 20
	//
	// }
	Constraints *string `json:"Constraints,omitempty" xml:"Constraints,omitempty"`
	// The description of the common parameter. The description must be 1 to 200 characters in length.
	//
	// example:
	//
	// parameter
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the parameter. The name must be 1 to 200 characters in length, and can contain letters, digits, hyphens (-), and underscores (_). It cannot start with ALIYUN, ACS, ALIBABA, ALICLOUD, or OOS.
	//
	// This parameter is required.
	//
	// example:
	//
	// MyParameter
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the region.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfmxsn4m4******
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The tags.
	//
	// example:
	//
	// {"k1": "v1", "k2": "v2"}
	TagsShrink *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The data type of the parameter. Valid values: String and StringList.
	//
	// This parameter is required.
	//
	// example:
	//
	// String
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The value of the common parameter. The value must be 1 to 4096 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// parameter
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateParameterShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateParameterShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateParameterShrinkRequest) SetClientToken(v string) *CreateParameterShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateParameterShrinkRequest) SetConstraints(v string) *CreateParameterShrinkRequest {
	s.Constraints = &v
	return s
}

func (s *CreateParameterShrinkRequest) SetDescription(v string) *CreateParameterShrinkRequest {
	s.Description = &v
	return s
}

func (s *CreateParameterShrinkRequest) SetName(v string) *CreateParameterShrinkRequest {
	s.Name = &v
	return s
}

func (s *CreateParameterShrinkRequest) SetRegionId(v string) *CreateParameterShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *CreateParameterShrinkRequest) SetResourceGroupId(v string) *CreateParameterShrinkRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateParameterShrinkRequest) SetTagsShrink(v string) *CreateParameterShrinkRequest {
	s.TagsShrink = &v
	return s
}

func (s *CreateParameterShrinkRequest) SetType(v string) *CreateParameterShrinkRequest {
	s.Type = &v
	return s
}

func (s *CreateParameterShrinkRequest) SetValue(v string) *CreateParameterShrinkRequest {
	s.Value = &v
	return s
}

type CreateParameterResponseBody struct {
	// The information about the common parameter.
	Parameter *CreateParameterResponseBodyParameter `json:"Parameter,omitempty" xml:"Parameter,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 0B419FF3-ABC6-4DF0-95E5-636DC8CBB8AF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateParameterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateParameterResponseBody) GoString() string {
	return s.String()
}

func (s *CreateParameterResponseBody) SetParameter(v *CreateParameterResponseBodyParameter) *CreateParameterResponseBody {
	s.Parameter = v
	return s
}

func (s *CreateParameterResponseBody) SetRequestId(v string) *CreateParameterResponseBody {
	s.RequestId = &v
	return s
}

type CreateParameterResponseBodyParameter struct {
	// The constraints of the common parameter.
	//
	// example:
	//
	// "{\\"AllowedValues\\":[\\"parameter\\"],\\"AllowedPattern\\":\\"parameter\\",\\"MinLength\\":0,\\"MaxLength\\":20}"
	Constraints *string `json:"Constraints,omitempty" xml:"Constraints,omitempty"`
	// The user who created the common parameter.
	//
	// example:
	//
	// root(130900000)
	CreatedBy *string `json:"CreatedBy,omitempty" xml:"CreatedBy,omitempty"`
	// The time when the common parameter was created.
	//
	// example:
	//
	// 2020-09-01T08:01:43Z
	CreatedDate *string `json:"CreatedDate,omitempty" xml:"CreatedDate,omitempty"`
	// The description of the common parameter.
	//
	// example:
	//
	// parameter
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the common parameter.
	//
	// example:
	//
	// p-4c4b401cab6747xxxxxx
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the common parameter.
	//
	// example:
	//
	// MyParameter
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The version number of the common parameter.
	//
	// example:
	//
	// 1
	ParameterVersion *int32 `json:"ParameterVersion,omitempty" xml:"ParameterVersion,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfmxsn4m4******
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The share type of the common parameter.
	//
	// example:
	//
	// Private
	ShareType *string `json:"ShareType,omitempty" xml:"ShareType,omitempty"`
	// The tags.
	//
	// example:
	//
	// {"k1": "v1", "k2": "v2"}
	Tags map[string]interface{} `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The type of the common parameter.
	//
	// example:
	//
	// String
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The user who updated the common parameter.
	//
	// example:
	//
	// root(130900000)
	UpdatedBy *string `json:"UpdatedBy,omitempty" xml:"UpdatedBy,omitempty"`
	// The time when the common parameter was updated.
	//
	// example:
	//
	// 2020-09-01T08:01:43Z
	UpdatedDate *string `json:"UpdatedDate,omitempty" xml:"UpdatedDate,omitempty"`
}

func (s CreateParameterResponseBodyParameter) String() string {
	return tea.Prettify(s)
}

func (s CreateParameterResponseBodyParameter) GoString() string {
	return s.String()
}

func (s *CreateParameterResponseBodyParameter) SetConstraints(v string) *CreateParameterResponseBodyParameter {
	s.Constraints = &v
	return s
}

func (s *CreateParameterResponseBodyParameter) SetCreatedBy(v string) *CreateParameterResponseBodyParameter {
	s.CreatedBy = &v
	return s
}

func (s *CreateParameterResponseBodyParameter) SetCreatedDate(v string) *CreateParameterResponseBodyParameter {
	s.CreatedDate = &v
	return s
}

func (s *CreateParameterResponseBodyParameter) SetDescription(v string) *CreateParameterResponseBodyParameter {
	s.Description = &v
	return s
}

func (s *CreateParameterResponseBodyParameter) SetId(v string) *CreateParameterResponseBodyParameter {
	s.Id = &v
	return s
}

func (s *CreateParameterResponseBodyParameter) SetName(v string) *CreateParameterResponseBodyParameter {
	s.Name = &v
	return s
}

func (s *CreateParameterResponseBodyParameter) SetParameterVersion(v int32) *CreateParameterResponseBodyParameter {
	s.ParameterVersion = &v
	return s
}

func (s *CreateParameterResponseBodyParameter) SetResourceGroupId(v string) *CreateParameterResponseBodyParameter {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateParameterResponseBodyParameter) SetShareType(v string) *CreateParameterResponseBodyParameter {
	s.ShareType = &v
	return s
}

func (s *CreateParameterResponseBodyParameter) SetTags(v map[string]interface{}) *CreateParameterResponseBodyParameter {
	s.Tags = v
	return s
}

func (s *CreateParameterResponseBodyParameter) SetType(v string) *CreateParameterResponseBodyParameter {
	s.Type = &v
	return s
}

func (s *CreateParameterResponseBodyParameter) SetUpdatedBy(v string) *CreateParameterResponseBodyParameter {
	s.UpdatedBy = &v
	return s
}

func (s *CreateParameterResponseBodyParameter) SetUpdatedDate(v string) *CreateParameterResponseBodyParameter {
	s.UpdatedDate = &v
	return s
}

type CreateParameterResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateParameterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateParameterResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateParameterResponse) GoString() string {
	return s.String()
}

func (s *CreateParameterResponse) SetHeaders(v map[string]*string) *CreateParameterResponse {
	s.Headers = v
	return s
}

func (s *CreateParameterResponse) SetStatusCode(v int32) *CreateParameterResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateParameterResponse) SetBody(v *CreateParameterResponseBody) *CreateParameterResponse {
	s.Body = v
	return s
}

type CreatePatchBaselineRequest struct {
	// The rules of scanning and installing patches for the specified operating system.
	//
	// This parameter is required.
	//
	// example:
	//
	// {"PatchRules":[{"PatchFilterGroup":[{"Key":"PatchSet","Values":["OS"]},{"Key":"ProductFamily","Values":["Windows"]},{"Key":"Product","Values":["Windows 10","Windows 7"]},{"Key":"Classification","Values":["Security Updates","Updates","Update Rollups","Critical Updates"]},{"Key":"Severity","Values":["Critical","Important","Moderate"]}],"ApproveAfterDays":7,"ApproveUntilDate":"","EnableNonSecurity":true,"ComplianceLevel":"Medium"}]}
	ApprovalRules *string `json:"ApprovalRules,omitempty" xml:"ApprovalRules,omitempty"`
	// The approved patches.
	ApprovedPatches []*string `json:"ApprovedPatches,omitempty" xml:"ApprovedPatches,omitempty" type:"Repeated"`
	// Specifies whether the approved patch involves updates other than security-related updates.
	//
	// example:
	//
	// true
	ApprovedPatchesEnableNonSecurity *bool `json:"ApprovedPatchesEnableNonSecurity,omitempty" xml:"ApprovedPatchesEnableNonSecurity,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// example:
	//
	// -
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The description of the patch baseline.
	//
	// example:
	//
	// PatchBaseline
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the patch baseline.
	//
	// This parameter is required.
	//
	// example:
	//
	// MyPatchBaseline
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The type of the operating system. Valid values:
	//
	// 	- Windows
	//
	// 	- Ubuntu
	//
	// 	- CentOS
	//
	// 	- Debian
	//
	// 	- AliyunLinux
	//
	// 	- RedhatEnterpriseLinux
	//
	// 	- Anolis
	//
	// 	- AlmaLinux
	//
	// This parameter is required.
	//
	// example:
	//
	// Windows
	OperationSystem *string `json:"OperationSystem,omitempty" xml:"OperationSystem,omitempty"`
	// The ID of the region in which you want to create a patch baseline.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The rejected patches.
	RejectedPatches []*string `json:"RejectedPatches,omitempty" xml:"RejectedPatches,omitempty" type:"Repeated"`
	// The action of the rejected patch.
	//
	// example:
	//
	// ALLOW_AS_DEPENDENCY
	RejectedPatchesAction *string `json:"RejectedPatchesAction,omitempty" xml:"RejectedPatchesAction,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfmxsn4m4******
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The patch source configurations.
	Sources []*string `json:"Sources,omitempty" xml:"Sources,omitempty" type:"Repeated"`
	// The tags.
	Tags []*CreatePatchBaselineRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s CreatePatchBaselineRequest) String() string {
	return tea.Prettify(s)
}

func (s CreatePatchBaselineRequest) GoString() string {
	return s.String()
}

func (s *CreatePatchBaselineRequest) SetApprovalRules(v string) *CreatePatchBaselineRequest {
	s.ApprovalRules = &v
	return s
}

func (s *CreatePatchBaselineRequest) SetApprovedPatches(v []*string) *CreatePatchBaselineRequest {
	s.ApprovedPatches = v
	return s
}

func (s *CreatePatchBaselineRequest) SetApprovedPatchesEnableNonSecurity(v bool) *CreatePatchBaselineRequest {
	s.ApprovedPatchesEnableNonSecurity = &v
	return s
}

func (s *CreatePatchBaselineRequest) SetClientToken(v string) *CreatePatchBaselineRequest {
	s.ClientToken = &v
	return s
}

func (s *CreatePatchBaselineRequest) SetDescription(v string) *CreatePatchBaselineRequest {
	s.Description = &v
	return s
}

func (s *CreatePatchBaselineRequest) SetName(v string) *CreatePatchBaselineRequest {
	s.Name = &v
	return s
}

func (s *CreatePatchBaselineRequest) SetOperationSystem(v string) *CreatePatchBaselineRequest {
	s.OperationSystem = &v
	return s
}

func (s *CreatePatchBaselineRequest) SetRegionId(v string) *CreatePatchBaselineRequest {
	s.RegionId = &v
	return s
}

func (s *CreatePatchBaselineRequest) SetRejectedPatches(v []*string) *CreatePatchBaselineRequest {
	s.RejectedPatches = v
	return s
}

func (s *CreatePatchBaselineRequest) SetRejectedPatchesAction(v string) *CreatePatchBaselineRequest {
	s.RejectedPatchesAction = &v
	return s
}

func (s *CreatePatchBaselineRequest) SetResourceGroupId(v string) *CreatePatchBaselineRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreatePatchBaselineRequest) SetSources(v []*string) *CreatePatchBaselineRequest {
	s.Sources = v
	return s
}

func (s *CreatePatchBaselineRequest) SetTags(v []*CreatePatchBaselineRequestTags) *CreatePatchBaselineRequest {
	s.Tags = v
	return s
}

type CreatePatchBaselineRequestTags struct {
	// The key of the tag.
	//
	// example:
	//
	// key
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag.
	//
	// example:
	//
	// value
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreatePatchBaselineRequestTags) String() string {
	return tea.Prettify(s)
}

func (s CreatePatchBaselineRequestTags) GoString() string {
	return s.String()
}

func (s *CreatePatchBaselineRequestTags) SetKey(v string) *CreatePatchBaselineRequestTags {
	s.Key = &v
	return s
}

func (s *CreatePatchBaselineRequestTags) SetValue(v string) *CreatePatchBaselineRequestTags {
	s.Value = &v
	return s
}

type CreatePatchBaselineShrinkRequest struct {
	// The rules of scanning and installing patches for the specified operating system.
	//
	// This parameter is required.
	//
	// example:
	//
	// {"PatchRules":[{"PatchFilterGroup":[{"Key":"PatchSet","Values":["OS"]},{"Key":"ProductFamily","Values":["Windows"]},{"Key":"Product","Values":["Windows 10","Windows 7"]},{"Key":"Classification","Values":["Security Updates","Updates","Update Rollups","Critical Updates"]},{"Key":"Severity","Values":["Critical","Important","Moderate"]}],"ApproveAfterDays":7,"ApproveUntilDate":"","EnableNonSecurity":true,"ComplianceLevel":"Medium"}]}
	ApprovalRules *string `json:"ApprovalRules,omitempty" xml:"ApprovalRules,omitempty"`
	// The approved patches.
	ApprovedPatchesShrink *string `json:"ApprovedPatches,omitempty" xml:"ApprovedPatches,omitempty"`
	// Specifies whether the approved patch involves updates other than security-related updates.
	//
	// example:
	//
	// true
	ApprovedPatchesEnableNonSecurity *bool `json:"ApprovedPatchesEnableNonSecurity,omitempty" xml:"ApprovedPatchesEnableNonSecurity,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// example:
	//
	// -
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The description of the patch baseline.
	//
	// example:
	//
	// PatchBaseline
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the patch baseline.
	//
	// This parameter is required.
	//
	// example:
	//
	// MyPatchBaseline
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The type of the operating system. Valid values:
	//
	// 	- Windows
	//
	// 	- Ubuntu
	//
	// 	- CentOS
	//
	// 	- Debian
	//
	// 	- AliyunLinux
	//
	// 	- RedhatEnterpriseLinux
	//
	// 	- Anolis
	//
	// 	- AlmaLinux
	//
	// This parameter is required.
	//
	// example:
	//
	// Windows
	OperationSystem *string `json:"OperationSystem,omitempty" xml:"OperationSystem,omitempty"`
	// The ID of the region in which you want to create a patch baseline.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The rejected patches.
	RejectedPatchesShrink *string `json:"RejectedPatches,omitempty" xml:"RejectedPatches,omitempty"`
	// The action of the rejected patch.
	//
	// example:
	//
	// ALLOW_AS_DEPENDENCY
	RejectedPatchesAction *string `json:"RejectedPatchesAction,omitempty" xml:"RejectedPatchesAction,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfmxsn4m4******
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The patch source configurations.
	SourcesShrink *string `json:"Sources,omitempty" xml:"Sources,omitempty"`
	// The tags.
	TagsShrink *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
}

func (s CreatePatchBaselineShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreatePatchBaselineShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreatePatchBaselineShrinkRequest) SetApprovalRules(v string) *CreatePatchBaselineShrinkRequest {
	s.ApprovalRules = &v
	return s
}

func (s *CreatePatchBaselineShrinkRequest) SetApprovedPatchesShrink(v string) *CreatePatchBaselineShrinkRequest {
	s.ApprovedPatchesShrink = &v
	return s
}

func (s *CreatePatchBaselineShrinkRequest) SetApprovedPatchesEnableNonSecurity(v bool) *CreatePatchBaselineShrinkRequest {
	s.ApprovedPatchesEnableNonSecurity = &v
	return s
}

func (s *CreatePatchBaselineShrinkRequest) SetClientToken(v string) *CreatePatchBaselineShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *CreatePatchBaselineShrinkRequest) SetDescription(v string) *CreatePatchBaselineShrinkRequest {
	s.Description = &v
	return s
}

func (s *CreatePatchBaselineShrinkRequest) SetName(v string) *CreatePatchBaselineShrinkRequest {
	s.Name = &v
	return s
}

func (s *CreatePatchBaselineShrinkRequest) SetOperationSystem(v string) *CreatePatchBaselineShrinkRequest {
	s.OperationSystem = &v
	return s
}

func (s *CreatePatchBaselineShrinkRequest) SetRegionId(v string) *CreatePatchBaselineShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *CreatePatchBaselineShrinkRequest) SetRejectedPatchesShrink(v string) *CreatePatchBaselineShrinkRequest {
	s.RejectedPatchesShrink = &v
	return s
}

func (s *CreatePatchBaselineShrinkRequest) SetRejectedPatchesAction(v string) *CreatePatchBaselineShrinkRequest {
	s.RejectedPatchesAction = &v
	return s
}

func (s *CreatePatchBaselineShrinkRequest) SetResourceGroupId(v string) *CreatePatchBaselineShrinkRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreatePatchBaselineShrinkRequest) SetSourcesShrink(v string) *CreatePatchBaselineShrinkRequest {
	s.SourcesShrink = &v
	return s
}

func (s *CreatePatchBaselineShrinkRequest) SetTagsShrink(v string) *CreatePatchBaselineShrinkRequest {
	s.TagsShrink = &v
	return s
}

type CreatePatchBaselineResponseBody struct {
	// The details of the patch baseline.
	PatchBaseline *CreatePatchBaselineResponseBodyPatchBaseline `json:"PatchBaseline,omitempty" xml:"PatchBaseline,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// A5173FF6-D10D-5E8C-8F71-943C2A3E25C0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreatePatchBaselineResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreatePatchBaselineResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePatchBaselineResponseBody) SetPatchBaseline(v *CreatePatchBaselineResponseBodyPatchBaseline) *CreatePatchBaselineResponseBody {
	s.PatchBaseline = v
	return s
}

func (s *CreatePatchBaselineResponseBody) SetRequestId(v string) *CreatePatchBaselineResponseBody {
	s.RequestId = &v
	return s
}

type CreatePatchBaselineResponseBodyPatchBaseline struct {
	// The rules of scanning and installing patches for the specified operating system.
	//
	// example:
	//
	// {"PatchRules":[{"PatchFilterGroup":[{"Key":"PatchSet","Values":["OS"]},{"Key":"ProductFamily","Values":["Windows"]},{"Key":"Product","Values":["Windows 10","Windows 7"]},{"Key":"Classification","Values":["Security Updates","Updates","Update Rollups","Critical Updates"]},{"Key":"Severity","Values":["Critical","Important","Moderate"]}],"ApproveAfterDays":7,"ApproveUntilDate":"","EnableNonSecurity":true,"ComplianceLevel":"Medium"}]}
	ApprovalRules *string `json:"ApprovalRules,omitempty" xml:"ApprovalRules,omitempty"`
	// The approved patches.
	ApprovedPatches []*string `json:"ApprovedPatches,omitempty" xml:"ApprovedPatches,omitempty" type:"Repeated"`
	// Indicates whether the approved patch involves updates other than security-related updates.
	//
	// example:
	//
	// true
	ApprovedPatchesEnableNonSecurity *bool `json:"ApprovedPatchesEnableNonSecurity,omitempty" xml:"ApprovedPatchesEnableNonSecurity,omitempty"`
	// The creator of the patch baseline.
	//
	// example:
	//
	// root(130900000)
	CreatedBy *string `json:"CreatedBy,omitempty" xml:"CreatedBy,omitempty"`
	// The time when the patch baseline was created.
	//
	// example:
	//
	// 2021-09-08T06:25:41Z
	CreatedDate *string `json:"CreatedDate,omitempty" xml:"CreatedDate,omitempty"`
	// The description of the patch baseline.
	//
	// example:
	//
	// PatchBaseline
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the patch baseline.
	//
	// example:
	//
	// pb-0a0aeda72ed147eb97ea
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the patch baseline.
	//
	// example:
	//
	// MyPatchBaseline
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The type of the operating system.
	//
	// example:
	//
	// Windows
	OperationSystem *string `json:"OperationSystem,omitempty" xml:"OperationSystem,omitempty"`
	// The rejected patches.
	RejectedPatches []*string `json:"RejectedPatches,omitempty" xml:"RejectedPatches,omitempty" type:"Repeated"`
	// The action of the rejected patch.
	//
	// example:
	//
	// "ALLOW_AS_DEPENDENCY"
	RejectedPatchesAction *string `json:"RejectedPatchesAction,omitempty" xml:"RejectedPatchesAction,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfm3comlufxpva
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The share type of the patch baseline.
	//
	// example:
	//
	// Private
	ShareType *string `json:"ShareType,omitempty" xml:"ShareType,omitempty"`
	// The patch source configurations.
	Sources []*string `json:"Sources,omitempty" xml:"Sources,omitempty" type:"Repeated"`
	// The tags.
	Tags []*CreatePatchBaselineResponseBodyPatchBaselineTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The Alibaba Cloud account that last modified the information about the patch baseline.
	//
	// example:
	//
	// root(130900000)
	UpdatedBy *string `json:"UpdatedBy,omitempty" xml:"UpdatedBy,omitempty"`
	// The time when the information about the patch baseline was last modified.
	//
	// example:
	//
	// 2021-09-08T06:25:41Z
	UpdatedDate *string `json:"UpdatedDate,omitempty" xml:"UpdatedDate,omitempty"`
}

func (s CreatePatchBaselineResponseBodyPatchBaseline) String() string {
	return tea.Prettify(s)
}

func (s CreatePatchBaselineResponseBodyPatchBaseline) GoString() string {
	return s.String()
}

func (s *CreatePatchBaselineResponseBodyPatchBaseline) SetApprovalRules(v string) *CreatePatchBaselineResponseBodyPatchBaseline {
	s.ApprovalRules = &v
	return s
}

func (s *CreatePatchBaselineResponseBodyPatchBaseline) SetApprovedPatches(v []*string) *CreatePatchBaselineResponseBodyPatchBaseline {
	s.ApprovedPatches = v
	return s
}

func (s *CreatePatchBaselineResponseBodyPatchBaseline) SetApprovedPatchesEnableNonSecurity(v bool) *CreatePatchBaselineResponseBodyPatchBaseline {
	s.ApprovedPatchesEnableNonSecurity = &v
	return s
}

func (s *CreatePatchBaselineResponseBodyPatchBaseline) SetCreatedBy(v string) *CreatePatchBaselineResponseBodyPatchBaseline {
	s.CreatedBy = &v
	return s
}

func (s *CreatePatchBaselineResponseBodyPatchBaseline) SetCreatedDate(v string) *CreatePatchBaselineResponseBodyPatchBaseline {
	s.CreatedDate = &v
	return s
}

func (s *CreatePatchBaselineResponseBodyPatchBaseline) SetDescription(v string) *CreatePatchBaselineResponseBodyPatchBaseline {
	s.Description = &v
	return s
}

func (s *CreatePatchBaselineResponseBodyPatchBaseline) SetId(v string) *CreatePatchBaselineResponseBodyPatchBaseline {
	s.Id = &v
	return s
}

func (s *CreatePatchBaselineResponseBodyPatchBaseline) SetName(v string) *CreatePatchBaselineResponseBodyPatchBaseline {
	s.Name = &v
	return s
}

func (s *CreatePatchBaselineResponseBodyPatchBaseline) SetOperationSystem(v string) *CreatePatchBaselineResponseBodyPatchBaseline {
	s.OperationSystem = &v
	return s
}

func (s *CreatePatchBaselineResponseBodyPatchBaseline) SetRejectedPatches(v []*string) *CreatePatchBaselineResponseBodyPatchBaseline {
	s.RejectedPatches = v
	return s
}

func (s *CreatePatchBaselineResponseBodyPatchBaseline) SetRejectedPatchesAction(v string) *CreatePatchBaselineResponseBodyPatchBaseline {
	s.RejectedPatchesAction = &v
	return s
}

func (s *CreatePatchBaselineResponseBodyPatchBaseline) SetResourceGroupId(v string) *CreatePatchBaselineResponseBodyPatchBaseline {
	s.ResourceGroupId = &v
	return s
}

func (s *CreatePatchBaselineResponseBodyPatchBaseline) SetShareType(v string) *CreatePatchBaselineResponseBodyPatchBaseline {
	s.ShareType = &v
	return s
}

func (s *CreatePatchBaselineResponseBodyPatchBaseline) SetSources(v []*string) *CreatePatchBaselineResponseBodyPatchBaseline {
	s.Sources = v
	return s
}

func (s *CreatePatchBaselineResponseBodyPatchBaseline) SetTags(v []*CreatePatchBaselineResponseBodyPatchBaselineTags) *CreatePatchBaselineResponseBodyPatchBaseline {
	s.Tags = v
	return s
}

func (s *CreatePatchBaselineResponseBodyPatchBaseline) SetUpdatedBy(v string) *CreatePatchBaselineResponseBodyPatchBaseline {
	s.UpdatedBy = &v
	return s
}

func (s *CreatePatchBaselineResponseBodyPatchBaseline) SetUpdatedDate(v string) *CreatePatchBaselineResponseBodyPatchBaseline {
	s.UpdatedDate = &v
	return s
}

type CreatePatchBaselineResponseBodyPatchBaselineTags struct {
	// The key of the tag.
	//
	// example:
	//
	// key
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The value of the tag.
	//
	// example:
	//
	// value
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s CreatePatchBaselineResponseBodyPatchBaselineTags) String() string {
	return tea.Prettify(s)
}

func (s CreatePatchBaselineResponseBodyPatchBaselineTags) GoString() string {
	return s.String()
}

func (s *CreatePatchBaselineResponseBodyPatchBaselineTags) SetTagKey(v string) *CreatePatchBaselineResponseBodyPatchBaselineTags {
	s.TagKey = &v
	return s
}

func (s *CreatePatchBaselineResponseBodyPatchBaselineTags) SetTagValue(v string) *CreatePatchBaselineResponseBodyPatchBaselineTags {
	s.TagValue = &v
	return s
}

type CreatePatchBaselineResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreatePatchBaselineResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreatePatchBaselineResponse) String() string {
	return tea.Prettify(s)
}

func (s CreatePatchBaselineResponse) GoString() string {
	return s.String()
}

func (s *CreatePatchBaselineResponse) SetHeaders(v map[string]*string) *CreatePatchBaselineResponse {
	s.Headers = v
	return s
}

func (s *CreatePatchBaselineResponse) SetStatusCode(v int32) *CreatePatchBaselineResponse {
	s.StatusCode = &v
	return s
}

func (s *CreatePatchBaselineResponse) SetBody(v *CreatePatchBaselineResponseBody) *CreatePatchBaselineResponse {
	s.Body = v
	return s
}

type CreateSecretParameterRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can be up to 64 characters in length, and can contain letters, digits, hyphens (-), and underscores (_). For more information, see "How to ensure idempotence".
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-42665544****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The constraints of the encryption parameter. By default, this parameter is null. Valid values:
	//
	// 	- AllowedValues: The value that is allowed for the encryption parameter. It must be an array string.
	//
	// 	- AllowedPattern: The pattern that is allowed for the encryption parameter. It must be a regular expression.
	//
	// 	- MinLength: The minimum length of the encryption parameter.
	//
	// 	- MaxLength: The maximum length of the encryption parameter.
	//
	// example:
	//
	// \\"{\\"\\"AllowedValues":["secretparameter"],"AllowedPattern":"secretparameter","MinLength":0,"MaxLength":20}\\"
	Constraints *string `json:"Constraints,omitempty" xml:"Constraints,omitempty"`
	// The instance ID of the KMS instance.
	//
	// example:
	//
	// kst-hzz****
	DKMSInstanceId *string `json:"DKMSInstanceId,omitempty" xml:"DKMSInstanceId,omitempty"`
	// The description of the encryption parameter. The description must be 1 to 200 characters in length.
	//
	// example:
	//
	// SecretParameter
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The key ID of Key Management Service (KMS) that is used to encrypt the parameter.
	//
	// example:
	//
	// 80e9409f-78fa-42ab-84bd-83f40c******
	KeyId *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	// The name of the parameter. The name must be 1 to 180 characters in length, and can contain letters, digits, hyphens (-), and underscores (_). It cannot start with ALIYUN, ACS, ALIBABA, ALICLOUD, or OOS.
	//
	// This parameter is required.
	//
	// example:
	//
	// MySecretParameter
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the region.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfmxsn4m4******
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The tags.
	//
	// example:
	//
	// {"k1": "v1", "k2": "v2"}
	Tags map[string]interface{} `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The type of the parameter. Set the value to Secret.
	//
	// example:
	//
	// Secret
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The value of the encryption parameter. The value must be 1 to 4096 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// SecretParameter
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateSecretParameterRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateSecretParameterRequest) GoString() string {
	return s.String()
}

func (s *CreateSecretParameterRequest) SetClientToken(v string) *CreateSecretParameterRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateSecretParameterRequest) SetConstraints(v string) *CreateSecretParameterRequest {
	s.Constraints = &v
	return s
}

func (s *CreateSecretParameterRequest) SetDKMSInstanceId(v string) *CreateSecretParameterRequest {
	s.DKMSInstanceId = &v
	return s
}

func (s *CreateSecretParameterRequest) SetDescription(v string) *CreateSecretParameterRequest {
	s.Description = &v
	return s
}

func (s *CreateSecretParameterRequest) SetKeyId(v string) *CreateSecretParameterRequest {
	s.KeyId = &v
	return s
}

func (s *CreateSecretParameterRequest) SetName(v string) *CreateSecretParameterRequest {
	s.Name = &v
	return s
}

func (s *CreateSecretParameterRequest) SetRegionId(v string) *CreateSecretParameterRequest {
	s.RegionId = &v
	return s
}

func (s *CreateSecretParameterRequest) SetResourceGroupId(v string) *CreateSecretParameterRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateSecretParameterRequest) SetTags(v map[string]interface{}) *CreateSecretParameterRequest {
	s.Tags = v
	return s
}

func (s *CreateSecretParameterRequest) SetType(v string) *CreateSecretParameterRequest {
	s.Type = &v
	return s
}

func (s *CreateSecretParameterRequest) SetValue(v string) *CreateSecretParameterRequest {
	s.Value = &v
	return s
}

type CreateSecretParameterShrinkRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can be up to 64 characters in length, and can contain letters, digits, hyphens (-), and underscores (_). For more information, see "How to ensure idempotence".
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-42665544****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The constraints of the encryption parameter. By default, this parameter is null. Valid values:
	//
	// 	- AllowedValues: The value that is allowed for the encryption parameter. It must be an array string.
	//
	// 	- AllowedPattern: The pattern that is allowed for the encryption parameter. It must be a regular expression.
	//
	// 	- MinLength: The minimum length of the encryption parameter.
	//
	// 	- MaxLength: The maximum length of the encryption parameter.
	//
	// example:
	//
	// \\"{\\"\\"AllowedValues":["secretparameter"],"AllowedPattern":"secretparameter","MinLength":0,"MaxLength":20}\\"
	Constraints *string `json:"Constraints,omitempty" xml:"Constraints,omitempty"`
	// The instance ID of the KMS instance.
	//
	// example:
	//
	// kst-hzz****
	DKMSInstanceId *string `json:"DKMSInstanceId,omitempty" xml:"DKMSInstanceId,omitempty"`
	// The description of the encryption parameter. The description must be 1 to 200 characters in length.
	//
	// example:
	//
	// SecretParameter
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The key ID of Key Management Service (KMS) that is used to encrypt the parameter.
	//
	// example:
	//
	// 80e9409f-78fa-42ab-84bd-83f40c******
	KeyId *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	// The name of the parameter. The name must be 1 to 180 characters in length, and can contain letters, digits, hyphens (-), and underscores (_). It cannot start with ALIYUN, ACS, ALIBABA, ALICLOUD, or OOS.
	//
	// This parameter is required.
	//
	// example:
	//
	// MySecretParameter
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the region.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfmxsn4m4******
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The tags.
	//
	// example:
	//
	// {"k1": "v1", "k2": "v2"}
	TagsShrink *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The type of the parameter. Set the value to Secret.
	//
	// example:
	//
	// Secret
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The value of the encryption parameter. The value must be 1 to 4096 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// SecretParameter
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateSecretParameterShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateSecretParameterShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateSecretParameterShrinkRequest) SetClientToken(v string) *CreateSecretParameterShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateSecretParameterShrinkRequest) SetConstraints(v string) *CreateSecretParameterShrinkRequest {
	s.Constraints = &v
	return s
}

func (s *CreateSecretParameterShrinkRequest) SetDKMSInstanceId(v string) *CreateSecretParameterShrinkRequest {
	s.DKMSInstanceId = &v
	return s
}

func (s *CreateSecretParameterShrinkRequest) SetDescription(v string) *CreateSecretParameterShrinkRequest {
	s.Description = &v
	return s
}

func (s *CreateSecretParameterShrinkRequest) SetKeyId(v string) *CreateSecretParameterShrinkRequest {
	s.KeyId = &v
	return s
}

func (s *CreateSecretParameterShrinkRequest) SetName(v string) *CreateSecretParameterShrinkRequest {
	s.Name = &v
	return s
}

func (s *CreateSecretParameterShrinkRequest) SetRegionId(v string) *CreateSecretParameterShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *CreateSecretParameterShrinkRequest) SetResourceGroupId(v string) *CreateSecretParameterShrinkRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateSecretParameterShrinkRequest) SetTagsShrink(v string) *CreateSecretParameterShrinkRequest {
	s.TagsShrink = &v
	return s
}

func (s *CreateSecretParameterShrinkRequest) SetType(v string) *CreateSecretParameterShrinkRequest {
	s.Type = &v
	return s
}

func (s *CreateSecretParameterShrinkRequest) SetValue(v string) *CreateSecretParameterShrinkRequest {
	s.Value = &v
	return s
}

type CreateSecretParameterResponseBody struct {
	// The information about the encryption parameter.
	Parameter *CreateSecretParameterResponseBodyParameter `json:"Parameter,omitempty" xml:"Parameter,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 0B419FF3-ABC6-4DF0-95E5-636DC8CBB8AF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateSecretParameterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateSecretParameterResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSecretParameterResponseBody) SetParameter(v *CreateSecretParameterResponseBodyParameter) *CreateSecretParameterResponseBody {
	s.Parameter = v
	return s
}

func (s *CreateSecretParameterResponseBody) SetRequestId(v string) *CreateSecretParameterResponseBody {
	s.RequestId = &v
	return s
}

type CreateSecretParameterResponseBodyParameter struct {
	// The constraints of the encryption parameter.
	//
	// example:
	//
	// \\"{ 	"AllowedValues": ["secretparameter"], 	"AllowedPattern": "secretparameter", 	"MinLength": 0, 	"MaxLength": 20 }\\"
	Constraints *string `json:"Constraints,omitempty" xml:"Constraints,omitempty"`
	// The user who created the encryption parameter.
	//
	// example:
	//
	// root(130900000)
	CreatedBy *string `json:"CreatedBy,omitempty" xml:"CreatedBy,omitempty"`
	// The time when the encryption parameter was created.
	//
	// example:
	//
	// 2020-09-01T09:30:36Z
	CreatedDate *string `json:"CreatedDate,omitempty" xml:"CreatedDate,omitempty"`
	// The instance ID of the KMS instance.
	//
	// example:
	//
	// kst-hzz****
	DKMSInstanceId *string `json:"DKMSInstanceId,omitempty" xml:"DKMSInstanceId,omitempty"`
	// The description of the encryption parameter.
	//
	// example:
	//
	// SecretParameter
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the encryption parameter.
	//
	// example:
	//
	// p-0b0fff9919c946xxxxxx
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The key ID of KMS that is used to encrypt the parameter.
	//
	// example:
	//
	// 80e9409f-78fa-42ab-84bd-83f40c******
	KeyId *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	// The name of the encryption parameter.
	//
	// example:
	//
	// MyParameter
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The version number of the encryption parameter.
	//
	// example:
	//
	// 1
	ParameterVersion *int32 `json:"ParameterVersion,omitempty" xml:"ParameterVersion,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfmxsn4m4******
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The share type of the encryption parameter.
	//
	// example:
	//
	// Private
	ShareType *string `json:"ShareType,omitempty" xml:"ShareType,omitempty"`
	// The tags.
	//
	// example:
	//
	// {"k1": "v1", "k2": "v2"}
	Tags map[string]interface{} `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The type of the parameter.
	//
	// example:
	//
	// Secret
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The user who updated the encryption parameter.
	//
	// example:
	//
	// root(130900000)
	UpdatedBy *string `json:"UpdatedBy,omitempty" xml:"UpdatedBy,omitempty"`
	// The time when the encryption parameter was updated.
	//
	// example:
	//
	// 2020-09-01T09:30:36Z
	UpdatedDate *string `json:"UpdatedDate,omitempty" xml:"UpdatedDate,omitempty"`
}

func (s CreateSecretParameterResponseBodyParameter) String() string {
	return tea.Prettify(s)
}

func (s CreateSecretParameterResponseBodyParameter) GoString() string {
	return s.String()
}

func (s *CreateSecretParameterResponseBodyParameter) SetConstraints(v string) *CreateSecretParameterResponseBodyParameter {
	s.Constraints = &v
	return s
}

func (s *CreateSecretParameterResponseBodyParameter) SetCreatedBy(v string) *CreateSecretParameterResponseBodyParameter {
	s.CreatedBy = &v
	return s
}

func (s *CreateSecretParameterResponseBodyParameter) SetCreatedDate(v string) *CreateSecretParameterResponseBodyParameter {
	s.CreatedDate = &v
	return s
}

func (s *CreateSecretParameterResponseBodyParameter) SetDKMSInstanceId(v string) *CreateSecretParameterResponseBodyParameter {
	s.DKMSInstanceId = &v
	return s
}

func (s *CreateSecretParameterResponseBodyParameter) SetDescription(v string) *CreateSecretParameterResponseBodyParameter {
	s.Description = &v
	return s
}

func (s *CreateSecretParameterResponseBodyParameter) SetId(v string) *CreateSecretParameterResponseBodyParameter {
	s.Id = &v
	return s
}

func (s *CreateSecretParameterResponseBodyParameter) SetKeyId(v string) *CreateSecretParameterResponseBodyParameter {
	s.KeyId = &v
	return s
}

func (s *CreateSecretParameterResponseBodyParameter) SetName(v string) *CreateSecretParameterResponseBodyParameter {
	s.Name = &v
	return s
}

func (s *CreateSecretParameterResponseBodyParameter) SetParameterVersion(v int32) *CreateSecretParameterResponseBodyParameter {
	s.ParameterVersion = &v
	return s
}

func (s *CreateSecretParameterResponseBodyParameter) SetResourceGroupId(v string) *CreateSecretParameterResponseBodyParameter {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateSecretParameterResponseBodyParameter) SetShareType(v string) *CreateSecretParameterResponseBodyParameter {
	s.ShareType = &v
	return s
}

func (s *CreateSecretParameterResponseBodyParameter) SetTags(v map[string]interface{}) *CreateSecretParameterResponseBodyParameter {
	s.Tags = v
	return s
}

func (s *CreateSecretParameterResponseBodyParameter) SetType(v string) *CreateSecretParameterResponseBodyParameter {
	s.Type = &v
	return s
}

func (s *CreateSecretParameterResponseBodyParameter) SetUpdatedBy(v string) *CreateSecretParameterResponseBodyParameter {
	s.UpdatedBy = &v
	return s
}

func (s *CreateSecretParameterResponseBodyParameter) SetUpdatedDate(v string) *CreateSecretParameterResponseBodyParameter {
	s.UpdatedDate = &v
	return s
}

type CreateSecretParameterResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateSecretParameterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateSecretParameterResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateSecretParameterResponse) GoString() string {
	return s.String()
}

func (s *CreateSecretParameterResponse) SetHeaders(v map[string]*string) *CreateSecretParameterResponse {
	s.Headers = v
	return s
}

func (s *CreateSecretParameterResponse) SetStatusCode(v int32) *CreateSecretParameterResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSecretParameterResponse) SetBody(v *CreateSecretParameterResponseBody) *CreateSecretParameterResponse {
	s.Body = v
	return s
}

type CreateStateConfigurationRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// example:
	//
	// DASKJJLKADS-AHKLJHJSAKL-AJK
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The configuration mode. Valid values: ApplyOnce: The configuration is applied only once. After a configuration is updated, the new configuration is applied. ApplyAndMonitor: The configuration is applied only once. After the configuration is applied, the system only checks whether the configuration is migrated in the future. ApplyAndAutoCorrect: The configuration is always applied.
	//
	// example:
	//
	// ApplyOnce
	ConfigureMode *string `json:"ConfigureMode,omitempty" xml:"ConfigureMode,omitempty"`
	// The description of the desired-state configuration.
	//
	// example:
	//
	// The region ID.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The parameters.
	//
	// example:
	//
	// {     "policy": {       "ACS:Application": {         "Collection": "Enabled"       },       "ACS:Network": {         "Collection": "Enabled"       }     }   }
	Parameters *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfmxsn4m4******
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The schedule expression. The interval between two schedules must be a minimum of 30 minutes.
	//
	// This parameter is required.
	//
	// example:
	//
	// The ID of the resource group.
	ScheduleExpression *string `json:"ScheduleExpression,omitempty" xml:"ScheduleExpression,omitempty"`
	// The schedule type. Set the value to rate.
	//
	// This parameter is required.
	//
	// example:
	//
	// rate
	ScheduleType *string `json:"ScheduleType,omitempty" xml:"ScheduleType,omitempty"`
	// The tags to be added to the configuration.
	//
	// example:
	//
	// {"Key": "oos", "Value": "inventory"}
	Tags map[string]interface{} `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The resources to be queried.
	//
	// This parameter is required.
	//
	// example:
	//
	// {     "ResourceType": "ALIYUN::ECS::Instance",     "Filters": [       {         "Type": "All",         "RegionId": "cn-hangzhou",         "Parameters": {           "RegionId": "cn-hangzhou",           "Status": "Running"         }       }     ]   }
	Targets *string `json:"Targets,omitempty" xml:"Targets,omitempty"`
	// The name of the template. The name must be 1 to 200 characters in length and can contain letters, digits, hyphens (-), and underscores (_).
	//
	// This parameter is required.
	//
	// example:
	//
	// ACS-ECS-InventoryDataCollection
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// The version number of the template. If you do not specify this parameter, the latest version of the template is used.
	//
	// example:
	//
	// v1
	TemplateVersion *string `json:"TemplateVersion,omitempty" xml:"TemplateVersion,omitempty"`
}

func (s CreateStateConfigurationRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateStateConfigurationRequest) GoString() string {
	return s.String()
}

func (s *CreateStateConfigurationRequest) SetClientToken(v string) *CreateStateConfigurationRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateStateConfigurationRequest) SetConfigureMode(v string) *CreateStateConfigurationRequest {
	s.ConfigureMode = &v
	return s
}

func (s *CreateStateConfigurationRequest) SetDescription(v string) *CreateStateConfigurationRequest {
	s.Description = &v
	return s
}

func (s *CreateStateConfigurationRequest) SetParameters(v string) *CreateStateConfigurationRequest {
	s.Parameters = &v
	return s
}

func (s *CreateStateConfigurationRequest) SetRegionId(v string) *CreateStateConfigurationRequest {
	s.RegionId = &v
	return s
}

func (s *CreateStateConfigurationRequest) SetResourceGroupId(v string) *CreateStateConfigurationRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateStateConfigurationRequest) SetScheduleExpression(v string) *CreateStateConfigurationRequest {
	s.ScheduleExpression = &v
	return s
}

func (s *CreateStateConfigurationRequest) SetScheduleType(v string) *CreateStateConfigurationRequest {
	s.ScheduleType = &v
	return s
}

func (s *CreateStateConfigurationRequest) SetTags(v map[string]interface{}) *CreateStateConfigurationRequest {
	s.Tags = v
	return s
}

func (s *CreateStateConfigurationRequest) SetTargets(v string) *CreateStateConfigurationRequest {
	s.Targets = &v
	return s
}

func (s *CreateStateConfigurationRequest) SetTemplateName(v string) *CreateStateConfigurationRequest {
	s.TemplateName = &v
	return s
}

func (s *CreateStateConfigurationRequest) SetTemplateVersion(v string) *CreateStateConfigurationRequest {
	s.TemplateVersion = &v
	return s
}

type CreateStateConfigurationShrinkRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// example:
	//
	// DASKJJLKADS-AHKLJHJSAKL-AJK
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The configuration mode. Valid values: ApplyOnce: The configuration is applied only once. After a configuration is updated, the new configuration is applied. ApplyAndMonitor: The configuration is applied only once. After the configuration is applied, the system only checks whether the configuration is migrated in the future. ApplyAndAutoCorrect: The configuration is always applied.
	//
	// example:
	//
	// ApplyOnce
	ConfigureMode *string `json:"ConfigureMode,omitempty" xml:"ConfigureMode,omitempty"`
	// The description of the desired-state configuration.
	//
	// example:
	//
	// The region ID.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The parameters.
	//
	// example:
	//
	// {     "policy": {       "ACS:Application": {         "Collection": "Enabled"       },       "ACS:Network": {         "Collection": "Enabled"       }     }   }
	Parameters *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfmxsn4m4******
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The schedule expression. The interval between two schedules must be a minimum of 30 minutes.
	//
	// This parameter is required.
	//
	// example:
	//
	// The ID of the resource group.
	ScheduleExpression *string `json:"ScheduleExpression,omitempty" xml:"ScheduleExpression,omitempty"`
	// The schedule type. Set the value to rate.
	//
	// This parameter is required.
	//
	// example:
	//
	// rate
	ScheduleType *string `json:"ScheduleType,omitempty" xml:"ScheduleType,omitempty"`
	// The tags to be added to the configuration.
	//
	// example:
	//
	// {"Key": "oos", "Value": "inventory"}
	TagsShrink *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The resources to be queried.
	//
	// This parameter is required.
	//
	// example:
	//
	// {     "ResourceType": "ALIYUN::ECS::Instance",     "Filters": [       {         "Type": "All",         "RegionId": "cn-hangzhou",         "Parameters": {           "RegionId": "cn-hangzhou",           "Status": "Running"         }       }     ]   }
	Targets *string `json:"Targets,omitempty" xml:"Targets,omitempty"`
	// The name of the template. The name must be 1 to 200 characters in length and can contain letters, digits, hyphens (-), and underscores (_).
	//
	// This parameter is required.
	//
	// example:
	//
	// ACS-ECS-InventoryDataCollection
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// The version number of the template. If you do not specify this parameter, the latest version of the template is used.
	//
	// example:
	//
	// v1
	TemplateVersion *string `json:"TemplateVersion,omitempty" xml:"TemplateVersion,omitempty"`
}

func (s CreateStateConfigurationShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateStateConfigurationShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateStateConfigurationShrinkRequest) SetClientToken(v string) *CreateStateConfigurationShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateStateConfigurationShrinkRequest) SetConfigureMode(v string) *CreateStateConfigurationShrinkRequest {
	s.ConfigureMode = &v
	return s
}

func (s *CreateStateConfigurationShrinkRequest) SetDescription(v string) *CreateStateConfigurationShrinkRequest {
	s.Description = &v
	return s
}

func (s *CreateStateConfigurationShrinkRequest) SetParameters(v string) *CreateStateConfigurationShrinkRequest {
	s.Parameters = &v
	return s
}

func (s *CreateStateConfigurationShrinkRequest) SetRegionId(v string) *CreateStateConfigurationShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *CreateStateConfigurationShrinkRequest) SetResourceGroupId(v string) *CreateStateConfigurationShrinkRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateStateConfigurationShrinkRequest) SetScheduleExpression(v string) *CreateStateConfigurationShrinkRequest {
	s.ScheduleExpression = &v
	return s
}

func (s *CreateStateConfigurationShrinkRequest) SetScheduleType(v string) *CreateStateConfigurationShrinkRequest {
	s.ScheduleType = &v
	return s
}

func (s *CreateStateConfigurationShrinkRequest) SetTagsShrink(v string) *CreateStateConfigurationShrinkRequest {
	s.TagsShrink = &v
	return s
}

func (s *CreateStateConfigurationShrinkRequest) SetTargets(v string) *CreateStateConfigurationShrinkRequest {
	s.Targets = &v
	return s
}

func (s *CreateStateConfigurationShrinkRequest) SetTemplateName(v string) *CreateStateConfigurationShrinkRequest {
	s.TemplateName = &v
	return s
}

func (s *CreateStateConfigurationShrinkRequest) SetTemplateVersion(v string) *CreateStateConfigurationShrinkRequest {
	s.TemplateVersion = &v
	return s
}

type CreateStateConfigurationResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 1306108F-610C-40FD-AAD5-DA13E8B00BE9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the desired-state configuration.
	StateConfiguration *CreateStateConfigurationResponseBodyStateConfiguration `json:"StateConfiguration,omitempty" xml:"StateConfiguration,omitempty" type:"Struct"`
}

func (s CreateStateConfigurationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateStateConfigurationResponseBody) GoString() string {
	return s.String()
}

func (s *CreateStateConfigurationResponseBody) SetRequestId(v string) *CreateStateConfigurationResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateStateConfigurationResponseBody) SetStateConfiguration(v *CreateStateConfigurationResponseBodyStateConfiguration) *CreateStateConfigurationResponseBody {
	s.StateConfiguration = v
	return s
}

type CreateStateConfigurationResponseBodyStateConfiguration struct {
	// The configuration mode. Valid values:
	//
	// example:
	//
	// ApplyAndAutoCorrect
	ConfigureMode *string `json:"ConfigureMode,omitempty" xml:"ConfigureMode,omitempty"`
	// The time when the desired-state configuration was created.
	//
	// example:
	//
	// 2021-03-22T03:13:32Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description.
	//
	// example:
	//
	// collect inventory data
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The parameters.
	//
	// example:
	//
	// {"policy": {"ACS:Network": {"Collection": "Enabled"}, "ACS:Application": {"Collection": "Enabled"}}}
	Parameters map[string]interface{} `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfmxsn4m4******
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The schedule expression.
	//
	// example:
	//
	// 1 hour
	ScheduleExpression *string `json:"ScheduleExpression,omitempty" xml:"ScheduleExpression,omitempty"`
	// The schedule type.
	//
	// example:
	//
	// rate
	ScheduleType *string `json:"ScheduleType,omitempty" xml:"ScheduleType,omitempty"`
	// The ID of the desired-state configuration.
	//
	// example:
	//
	// sc-a538febe18fabcdef
	StateConfigurationId *string `json:"StateConfigurationId,omitempty" xml:"StateConfigurationId,omitempty"`
	// The tags added to the configuration.
	//
	// example:
	//
	// {"Key": "oos", "Value": "inventory"}
	Tags map[string]interface{} `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The queried resources.
	//
	// example:
	//
	// {     "ResourceType": "ALIYUN::ECS::Instance",     "Filters": [       {         "Type": "All",         "RegionId": "cn-hangzhou",         "Parameters": {           "RegionId": "cn-hangzhou",           "Status": "Running"         }       }     ]   }
	Targets *string `json:"Targets,omitempty" xml:"Targets,omitempty"`
	// The template ID.
	//
	// example:
	//
	// t-1234asadf
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The name of the template.
	//
	// example:
	//
	// ACS-ECS-InventoryDataCollection
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// The name of the template version.
	//
	// example:
	//
	// v1
	TemplateVersion *string `json:"TemplateVersion,omitempty" xml:"TemplateVersion,omitempty"`
}

func (s CreateStateConfigurationResponseBodyStateConfiguration) String() string {
	return tea.Prettify(s)
}

func (s CreateStateConfigurationResponseBodyStateConfiguration) GoString() string {
	return s.String()
}

func (s *CreateStateConfigurationResponseBodyStateConfiguration) SetConfigureMode(v string) *CreateStateConfigurationResponseBodyStateConfiguration {
	s.ConfigureMode = &v
	return s
}

func (s *CreateStateConfigurationResponseBodyStateConfiguration) SetCreateTime(v string) *CreateStateConfigurationResponseBodyStateConfiguration {
	s.CreateTime = &v
	return s
}

func (s *CreateStateConfigurationResponseBodyStateConfiguration) SetDescription(v string) *CreateStateConfigurationResponseBodyStateConfiguration {
	s.Description = &v
	return s
}

func (s *CreateStateConfigurationResponseBodyStateConfiguration) SetParameters(v map[string]interface{}) *CreateStateConfigurationResponseBodyStateConfiguration {
	s.Parameters = v
	return s
}

func (s *CreateStateConfigurationResponseBodyStateConfiguration) SetResourceGroupId(v string) *CreateStateConfigurationResponseBodyStateConfiguration {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateStateConfigurationResponseBodyStateConfiguration) SetScheduleExpression(v string) *CreateStateConfigurationResponseBodyStateConfiguration {
	s.ScheduleExpression = &v
	return s
}

func (s *CreateStateConfigurationResponseBodyStateConfiguration) SetScheduleType(v string) *CreateStateConfigurationResponseBodyStateConfiguration {
	s.ScheduleType = &v
	return s
}

func (s *CreateStateConfigurationResponseBodyStateConfiguration) SetStateConfigurationId(v string) *CreateStateConfigurationResponseBodyStateConfiguration {
	s.StateConfigurationId = &v
	return s
}

func (s *CreateStateConfigurationResponseBodyStateConfiguration) SetTags(v map[string]interface{}) *CreateStateConfigurationResponseBodyStateConfiguration {
	s.Tags = v
	return s
}

func (s *CreateStateConfigurationResponseBodyStateConfiguration) SetTargets(v string) *CreateStateConfigurationResponseBodyStateConfiguration {
	s.Targets = &v
	return s
}

func (s *CreateStateConfigurationResponseBodyStateConfiguration) SetTemplateId(v string) *CreateStateConfigurationResponseBodyStateConfiguration {
	s.TemplateId = &v
	return s
}

func (s *CreateStateConfigurationResponseBodyStateConfiguration) SetTemplateName(v string) *CreateStateConfigurationResponseBodyStateConfiguration {
	s.TemplateName = &v
	return s
}

func (s *CreateStateConfigurationResponseBodyStateConfiguration) SetTemplateVersion(v string) *CreateStateConfigurationResponseBodyStateConfiguration {
	s.TemplateVersion = &v
	return s
}

type CreateStateConfigurationResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateStateConfigurationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateStateConfigurationResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateStateConfigurationResponse) GoString() string {
	return s.String()
}

func (s *CreateStateConfigurationResponse) SetHeaders(v map[string]*string) *CreateStateConfigurationResponse {
	s.Headers = v
	return s
}

func (s *CreateStateConfigurationResponse) SetStatusCode(v int32) *CreateStateConfigurationResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateStateConfigurationResponse) SetBody(v *CreateStateConfigurationResponseBody) *CreateStateConfigurationResponse {
	s.Body = v
	return s
}

type CreateTemplateRequest struct {
	// The content of the template. The content must be in the JSON or YAML format, and its maximum size is 64 KB.
	//
	// This parameter is required.
	//
	// example:
	//
	// {"FormatVersion": "OOS-2019-06-01", "Description": "Describe instances of given status", "Parameters": {"Status": {"Type": "String", "Description": "(Required) The status of the Ecs instance."}}, "Tasks": [{"Properties": {"Parameters": {"Status": "{{ Status }}"}, "API": "DescribeInstances", "Service": "Ecs"}, "Name": "foo", "Action": "ACS::ExecuteApi"}]}
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The ID of the region.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfmxsn4m4******
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The tag keys and tag values. The number of key-value pairs ranges from 1 to 20.
	//
	// example:
	//
	// {"k1":"v1","k2":"v2"}
	Tags map[string]interface{} `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The name of the template. The name can be 1 to 200 characters in length and can contain letters, digits, hyphens (-), and underscores (_). The name cannot start with ALIYUN, ACS, ALIBABA, or ALICLOUD.
	//
	// This parameter is required.
	//
	// example:
	//
	// MyTemplate
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// The name of the version of the template.
	//
	// example:
	//
	// v2
	VersionName *string `json:"VersionName,omitempty" xml:"VersionName,omitempty"`
}

func (s CreateTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateTemplateRequest) GoString() string {
	return s.String()
}

func (s *CreateTemplateRequest) SetContent(v string) *CreateTemplateRequest {
	s.Content = &v
	return s
}

func (s *CreateTemplateRequest) SetRegionId(v string) *CreateTemplateRequest {
	s.RegionId = &v
	return s
}

func (s *CreateTemplateRequest) SetResourceGroupId(v string) *CreateTemplateRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateTemplateRequest) SetTags(v map[string]interface{}) *CreateTemplateRequest {
	s.Tags = v
	return s
}

func (s *CreateTemplateRequest) SetTemplateName(v string) *CreateTemplateRequest {
	s.TemplateName = &v
	return s
}

func (s *CreateTemplateRequest) SetVersionName(v string) *CreateTemplateRequest {
	s.VersionName = &v
	return s
}

type CreateTemplateShrinkRequest struct {
	// The content of the template. The content must be in the JSON or YAML format, and its maximum size is 64 KB.
	//
	// This parameter is required.
	//
	// example:
	//
	// {"FormatVersion": "OOS-2019-06-01", "Description": "Describe instances of given status", "Parameters": {"Status": {"Type": "String", "Description": "(Required) The status of the Ecs instance."}}, "Tasks": [{"Properties": {"Parameters": {"Status": "{{ Status }}"}, "API": "DescribeInstances", "Service": "Ecs"}, "Name": "foo", "Action": "ACS::ExecuteApi"}]}
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The ID of the region.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfmxsn4m4******
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The tag keys and tag values. The number of key-value pairs ranges from 1 to 20.
	//
	// example:
	//
	// {"k1":"v1","k2":"v2"}
	TagsShrink *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The name of the template. The name can be 1 to 200 characters in length and can contain letters, digits, hyphens (-), and underscores (_). The name cannot start with ALIYUN, ACS, ALIBABA, or ALICLOUD.
	//
	// This parameter is required.
	//
	// example:
	//
	// MyTemplate
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// The name of the version of the template.
	//
	// example:
	//
	// v2
	VersionName *string `json:"VersionName,omitempty" xml:"VersionName,omitempty"`
}

func (s CreateTemplateShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateTemplateShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateTemplateShrinkRequest) SetContent(v string) *CreateTemplateShrinkRequest {
	s.Content = &v
	return s
}

func (s *CreateTemplateShrinkRequest) SetRegionId(v string) *CreateTemplateShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *CreateTemplateShrinkRequest) SetResourceGroupId(v string) *CreateTemplateShrinkRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateTemplateShrinkRequest) SetTagsShrink(v string) *CreateTemplateShrinkRequest {
	s.TagsShrink = &v
	return s
}

func (s *CreateTemplateShrinkRequest) SetTemplateName(v string) *CreateTemplateShrinkRequest {
	s.TemplateName = &v
	return s
}

func (s *CreateTemplateShrinkRequest) SetVersionName(v string) *CreateTemplateShrinkRequest {
	s.VersionName = &v
	return s
}

type CreateTemplateResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 20758A-585D-4A41-A9B2-28DA8F4F534F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The metadata of the template.
	Template *CreateTemplateResponseBodyTemplate `json:"Template,omitempty" xml:"Template,omitempty" type:"Struct"`
	// The type of the template.
	//
	// example:
	//
	// Private
	TemplateType *string `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
}

func (s CreateTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTemplateResponseBody) SetRequestId(v string) *CreateTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateTemplateResponseBody) SetTemplate(v *CreateTemplateResponseBodyTemplate) *CreateTemplateResponseBody {
	s.Template = v
	return s
}

func (s *CreateTemplateResponseBody) SetTemplateType(v string) *CreateTemplateResponseBody {
	s.TemplateType = &v
	return s
}

type CreateTemplateResponseBodyTemplate struct {
	// The creator of the template.
	//
	// example:
	//
	// root(13090000)
	CreatedBy *string `json:"CreatedBy,omitempty" xml:"CreatedBy,omitempty"`
	// The time when the template was created.
	//
	// example:
	//
	// 2019-05-16T10:26:14Z
	CreatedDate *string `json:"CreatedDate,omitempty" xml:"CreatedDate,omitempty"`
	// The description of the template.
	//
	// example:
	//
	// Describe instances of given status
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Indicates whether the template was configured with a trigger.
	//
	// example:
	//
	// true
	HasTrigger *bool `json:"HasTrigger,omitempty" xml:"HasTrigger,omitempty"`
	// The SHA-256 value of the template content.
	//
	// example:
	//
	// 4bc7d7a21b3e003434b9c223f6e6d2578b5ebfeb5be28c1fcf8a8a1b11907bb4
	Hash *string `json:"Hash,omitempty" xml:"Hash,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfmxsn4m4******
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The share type of the template. The share type of the template that you create is Private.
	//
	// example:
	//
	// Private
	ShareType *string `json:"ShareType,omitempty" xml:"ShareType,omitempty"`
	// The tags of the resources.
	//
	// example:
	//
	// {     "k1":"v1",     "k2":"v2" }
	Tags map[string]interface{} `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The format of the template. The system automatically determines whether the format is JSON or YAML.
	//
	// example:
	//
	// JSON
	TemplateFormat *string `json:"TemplateFormat,omitempty" xml:"TemplateFormat,omitempty"`
	// The ID of the template.
	//
	// example:
	//
	// t-94753d38
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The name of the template.
	//
	// example:
	//
	// MyTemplate
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// The version of the template. The name of the version consists of the letter v and a number. The number starts from 1.
	//
	// example:
	//
	// v1
	TemplateVersion *string `json:"TemplateVersion,omitempty" xml:"TemplateVersion,omitempty"`
	// The Alibaba Cloud account that last modified the information about the template.
	//
	// example:
	//
	// root(130900000)
	UpdatedBy *string `json:"UpdatedBy,omitempty" xml:"UpdatedBy,omitempty"`
	// The time when the template was last updated.
	//
	// example:
	//
	// 2019-05-16T10:26:14Z
	UpdatedDate *string `json:"UpdatedDate,omitempty" xml:"UpdatedDate,omitempty"`
}

func (s CreateTemplateResponseBodyTemplate) String() string {
	return tea.Prettify(s)
}

func (s CreateTemplateResponseBodyTemplate) GoString() string {
	return s.String()
}

func (s *CreateTemplateResponseBodyTemplate) SetCreatedBy(v string) *CreateTemplateResponseBodyTemplate {
	s.CreatedBy = &v
	return s
}

func (s *CreateTemplateResponseBodyTemplate) SetCreatedDate(v string) *CreateTemplateResponseBodyTemplate {
	s.CreatedDate = &v
	return s
}

func (s *CreateTemplateResponseBodyTemplate) SetDescription(v string) *CreateTemplateResponseBodyTemplate {
	s.Description = &v
	return s
}

func (s *CreateTemplateResponseBodyTemplate) SetHasTrigger(v bool) *CreateTemplateResponseBodyTemplate {
	s.HasTrigger = &v
	return s
}

func (s *CreateTemplateResponseBodyTemplate) SetHash(v string) *CreateTemplateResponseBodyTemplate {
	s.Hash = &v
	return s
}

func (s *CreateTemplateResponseBodyTemplate) SetResourceGroupId(v string) *CreateTemplateResponseBodyTemplate {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateTemplateResponseBodyTemplate) SetShareType(v string) *CreateTemplateResponseBodyTemplate {
	s.ShareType = &v
	return s
}

func (s *CreateTemplateResponseBodyTemplate) SetTags(v map[string]interface{}) *CreateTemplateResponseBodyTemplate {
	s.Tags = v
	return s
}

func (s *CreateTemplateResponseBodyTemplate) SetTemplateFormat(v string) *CreateTemplateResponseBodyTemplate {
	s.TemplateFormat = &v
	return s
}

func (s *CreateTemplateResponseBodyTemplate) SetTemplateId(v string) *CreateTemplateResponseBodyTemplate {
	s.TemplateId = &v
	return s
}

func (s *CreateTemplateResponseBodyTemplate) SetTemplateName(v string) *CreateTemplateResponseBodyTemplate {
	s.TemplateName = &v
	return s
}

func (s *CreateTemplateResponseBodyTemplate) SetTemplateVersion(v string) *CreateTemplateResponseBodyTemplate {
	s.TemplateVersion = &v
	return s
}

func (s *CreateTemplateResponseBodyTemplate) SetUpdatedBy(v string) *CreateTemplateResponseBodyTemplate {
	s.UpdatedBy = &v
	return s
}

func (s *CreateTemplateResponseBodyTemplate) SetUpdatedDate(v string) *CreateTemplateResponseBodyTemplate {
	s.UpdatedDate = &v
	return s
}

type CreateTemplateResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateTemplateResponse) GoString() string {
	return s.String()
}

func (s *CreateTemplateResponse) SetHeaders(v map[string]*string) *CreateTemplateResponse {
	s.Headers = v
	return s
}

func (s *CreateTemplateResponse) SetStatusCode(v int32) *CreateTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateTemplateResponse) SetBody(v *CreateTemplateResponseBody) *CreateTemplateResponse {
	s.Body = v
	return s
}

type DeleteApplicationRequest struct {
	// Specifies whether to forcibly delete the application. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// False
	Force *bool `json:"Force,omitempty" xml:"Force,omitempty"`
	// The application name.
	//
	// This parameter is required.
	//
	// example:
	//
	// MyApplication
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The region ID. Set the value to cn-hangzhou.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Specifies whether to retain resources created by application manager when deleting the application. Valid values:
	//
	// - true
	//
	// - false
	//
	// example:
	//
	// false
	RetainResource *bool `json:"RetainResource,omitempty" xml:"RetainResource,omitempty"`
}

func (s DeleteApplicationRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteApplicationRequest) GoString() string {
	return s.String()
}

func (s *DeleteApplicationRequest) SetForce(v bool) *DeleteApplicationRequest {
	s.Force = &v
	return s
}

func (s *DeleteApplicationRequest) SetName(v string) *DeleteApplicationRequest {
	s.Name = &v
	return s
}

func (s *DeleteApplicationRequest) SetRegionId(v string) *DeleteApplicationRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteApplicationRequest) SetRetainResource(v bool) *DeleteApplicationRequest {
	s.RetainResource = &v
	return s
}

type DeleteApplicationResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 37A9F0FD-51D0-58D5-B91F-DF570281556B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteApplicationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteApplicationResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteApplicationResponseBody) SetRequestId(v string) *DeleteApplicationResponseBody {
	s.RequestId = &v
	return s
}

type DeleteApplicationResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteApplicationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteApplicationResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteApplicationResponse) GoString() string {
	return s.String()
}

func (s *DeleteApplicationResponse) SetHeaders(v map[string]*string) *DeleteApplicationResponse {
	s.Headers = v
	return s
}

func (s *DeleteApplicationResponse) SetStatusCode(v int32) *DeleteApplicationResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteApplicationResponse) SetBody(v *DeleteApplicationResponseBody) *DeleteApplicationResponse {
	s.Body = v
	return s
}

type DeleteApplicationGroupRequest struct {
	// The name of the application.
	//
	// This parameter is required.
	//
	// example:
	//
	// MyApplication
	ApplicationName *string `json:"ApplicationName,omitempty" xml:"ApplicationName,omitempty"`
	// The name of the application group.
	//
	// This parameter is required.
	//
	// example:
	//
	// MyApplicationGroup
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the region. Set the value to cn-hangzhou.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Specifies whether to retain resources created by application manager when deleting the application. Valid values:
	//
	// - true
	//
	// - false
	//
	// example:
	//
	// false
	RetainResource *bool `json:"RetainResource,omitempty" xml:"RetainResource,omitempty"`
}

func (s DeleteApplicationGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteApplicationGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteApplicationGroupRequest) SetApplicationName(v string) *DeleteApplicationGroupRequest {
	s.ApplicationName = &v
	return s
}

func (s *DeleteApplicationGroupRequest) SetName(v string) *DeleteApplicationGroupRequest {
	s.Name = &v
	return s
}

func (s *DeleteApplicationGroupRequest) SetRegionId(v string) *DeleteApplicationGroupRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteApplicationGroupRequest) SetRetainResource(v bool) *DeleteApplicationGroupRequest {
	s.RetainResource = &v
	return s
}

type DeleteApplicationGroupResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 9E011F98-4EE5-5E3A-8FA3-D38F2C531C1F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteApplicationGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteApplicationGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteApplicationGroupResponseBody) SetRequestId(v string) *DeleteApplicationGroupResponseBody {
	s.RequestId = &v
	return s
}

type DeleteApplicationGroupResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteApplicationGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteApplicationGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteApplicationGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteApplicationGroupResponse) SetHeaders(v map[string]*string) *DeleteApplicationGroupResponse {
	s.Headers = v
	return s
}

func (s *DeleteApplicationGroupResponse) SetStatusCode(v int32) *DeleteApplicationGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteApplicationGroupResponse) SetBody(v *DeleteApplicationGroupResponseBody) *DeleteApplicationGroupResponse {
	s.Body = v
	return s
}

type DeleteExecutionsRequest struct {
	// The execution IDs.
	//
	// You can specify multiple execution IDs in a JSON array in the format of `["xxxxxxxxx", "yyyyyyyyy", ... "zzzzzzzzz"]`. You can specify up to 100 execution IDs at a time. Separate multiple IDs with commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// ["exec-xxx"]
	ExecutionIds *string `json:"ExecutionIds,omitempty" xml:"ExecutionIds,omitempty"`
	Force        *bool   `json:"Force,omitempty" xml:"Force,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteExecutionsRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteExecutionsRequest) GoString() string {
	return s.String()
}

func (s *DeleteExecutionsRequest) SetExecutionIds(v string) *DeleteExecutionsRequest {
	s.ExecutionIds = &v
	return s
}

func (s *DeleteExecutionsRequest) SetForce(v bool) *DeleteExecutionsRequest {
	s.Force = &v
	return s
}

func (s *DeleteExecutionsRequest) SetRegionId(v string) *DeleteExecutionsRequest {
	s.RegionId = &v
	return s
}

type DeleteExecutionsResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 491DF8C2-34C9-4679-9DB3-4C0F49B129AC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteExecutionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteExecutionsResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteExecutionsResponseBody) SetRequestId(v string) *DeleteExecutionsResponseBody {
	s.RequestId = &v
	return s
}

type DeleteExecutionsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteExecutionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteExecutionsResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteExecutionsResponse) GoString() string {
	return s.String()
}

func (s *DeleteExecutionsResponse) SetHeaders(v map[string]*string) *DeleteExecutionsResponse {
	s.Headers = v
	return s
}

func (s *DeleteExecutionsResponse) SetStatusCode(v int32) *DeleteExecutionsResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteExecutionsResponse) SetBody(v *DeleteExecutionsResponseBody) *DeleteExecutionsResponse {
	s.Body = v
	return s
}

type DeleteParameterRequest struct {
	// The name of the common parameter. The name can be up to 180 characters in length and can contain only letters, digits, hyphens (-), and underscores (_). It cannot start with aliyun, acs, alibaba, alicloud, or oos.
	//
	// This parameter is required.
	//
	// example:
	//
	// MyParameter
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteParameterRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteParameterRequest) GoString() string {
	return s.String()
}

func (s *DeleteParameterRequest) SetName(v string) *DeleteParameterRequest {
	s.Name = &v
	return s
}

func (s *DeleteParameterRequest) SetRegionId(v string) *DeleteParameterRequest {
	s.RegionId = &v
	return s
}

type DeleteParameterResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// C0D02BDF-77F6-49F2-95C9-8E87121D2979
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteParameterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteParameterResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteParameterResponseBody) SetRequestId(v string) *DeleteParameterResponseBody {
	s.RequestId = &v
	return s
}

type DeleteParameterResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteParameterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteParameterResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteParameterResponse) GoString() string {
	return s.String()
}

func (s *DeleteParameterResponse) SetHeaders(v map[string]*string) *DeleteParameterResponse {
	s.Headers = v
	return s
}

func (s *DeleteParameterResponse) SetStatusCode(v int32) *DeleteParameterResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteParameterResponse) SetBody(v *DeleteParameterResponseBody) *DeleteParameterResponse {
	s.Body = v
	return s
}

type DeletePatchBaselineRequest struct {
	// The name of the patch baseline.
	//
	// This parameter is required.
	//
	// example:
	//
	// MyPatchBaseline
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeletePatchBaselineRequest) String() string {
	return tea.Prettify(s)
}

func (s DeletePatchBaselineRequest) GoString() string {
	return s.String()
}

func (s *DeletePatchBaselineRequest) SetName(v string) *DeletePatchBaselineRequest {
	s.Name = &v
	return s
}

func (s *DeletePatchBaselineRequest) SetRegionId(v string) *DeletePatchBaselineRequest {
	s.RegionId = &v
	return s
}

type DeletePatchBaselineResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 85197066-0209-5775-BBED-99DF9DA44E48
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeletePatchBaselineResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeletePatchBaselineResponseBody) GoString() string {
	return s.String()
}

func (s *DeletePatchBaselineResponseBody) SetRequestId(v string) *DeletePatchBaselineResponseBody {
	s.RequestId = &v
	return s
}

type DeletePatchBaselineResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeletePatchBaselineResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeletePatchBaselineResponse) String() string {
	return tea.Prettify(s)
}

func (s DeletePatchBaselineResponse) GoString() string {
	return s.String()
}

func (s *DeletePatchBaselineResponse) SetHeaders(v map[string]*string) *DeletePatchBaselineResponse {
	s.Headers = v
	return s
}

func (s *DeletePatchBaselineResponse) SetStatusCode(v int32) *DeletePatchBaselineResponse {
	s.StatusCode = &v
	return s
}

func (s *DeletePatchBaselineResponse) SetBody(v *DeletePatchBaselineResponseBody) *DeletePatchBaselineResponse {
	s.Body = v
	return s
}

type DeleteSecretParameterRequest struct {
	// The name of the encryption parameter. The name must be 1 to 180 characters in length and can contain letters, digits, hyphens (-), and underscores (_). It cannot start with ALIYUN, ACS, ALIBABA, ALICLOUD, or OOS.
	//
	// This parameter is required.
	//
	// example:
	//
	// MySecretParameter
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteSecretParameterRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteSecretParameterRequest) GoString() string {
	return s.String()
}

func (s *DeleteSecretParameterRequest) SetName(v string) *DeleteSecretParameterRequest {
	s.Name = &v
	return s
}

func (s *DeleteSecretParameterRequest) SetRegionId(v string) *DeleteSecretParameterRequest {
	s.RegionId = &v
	return s
}

type DeleteSecretParameterResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// C0D02BDF-77F6-49F2-95C9-8E87121D1944
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteSecretParameterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteSecretParameterResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSecretParameterResponseBody) SetRequestId(v string) *DeleteSecretParameterResponseBody {
	s.RequestId = &v
	return s
}

type DeleteSecretParameterResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteSecretParameterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteSecretParameterResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteSecretParameterResponse) GoString() string {
	return s.String()
}

func (s *DeleteSecretParameterResponse) SetHeaders(v map[string]*string) *DeleteSecretParameterResponse {
	s.Headers = v
	return s
}

func (s *DeleteSecretParameterResponse) SetStatusCode(v int32) *DeleteSecretParameterResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteSecretParameterResponse) SetBody(v *DeleteSecretParameterResponseBody) *DeleteSecretParameterResponse {
	s.Body = v
	return s
}

type DeleteStateConfigurationsRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// example:
	//
	// abcde3OARpx77No54nv6
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The IDs of desired-state configurations.
	//
	// This parameter is required.
	//
	// example:
	//
	// ["sc-asfgdhj12345"]
	StateConfigurationIds *string `json:"StateConfigurationIds,omitempty" xml:"StateConfigurationIds,omitempty"`
}

func (s DeleteStateConfigurationsRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteStateConfigurationsRequest) GoString() string {
	return s.String()
}

func (s *DeleteStateConfigurationsRequest) SetClientToken(v string) *DeleteStateConfigurationsRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteStateConfigurationsRequest) SetRegionId(v string) *DeleteStateConfigurationsRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteStateConfigurationsRequest) SetStateConfigurationIds(v string) *DeleteStateConfigurationsRequest {
	s.StateConfigurationIds = &v
	return s
}

type DeleteStateConfigurationsResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 12345B731-0FCE-48BA-8D42-605abcde
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteStateConfigurationsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteStateConfigurationsResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteStateConfigurationsResponseBody) SetRequestId(v string) *DeleteStateConfigurationsResponseBody {
	s.RequestId = &v
	return s
}

type DeleteStateConfigurationsResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteStateConfigurationsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteStateConfigurationsResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteStateConfigurationsResponse) GoString() string {
	return s.String()
}

func (s *DeleteStateConfigurationsResponse) SetHeaders(v map[string]*string) *DeleteStateConfigurationsResponse {
	s.Headers = v
	return s
}

func (s *DeleteStateConfigurationsResponse) SetStatusCode(v int32) *DeleteStateConfigurationsResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteStateConfigurationsResponse) SetBody(v *DeleteStateConfigurationsResponseBody) *DeleteStateConfigurationsResponse {
	s.Body = v
	return s
}

type DeleteTemplateRequest struct {
	// Specifies whether to delete the related executions when a template is deleted.
	//
	// example:
	//
	// false
	AutoDeleteExecutions *bool `json:"AutoDeleteExecutions,omitempty" xml:"AutoDeleteExecutions,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the template. The name can be 1 to 200 characters in length and can contain letters, digits, hyphens (-), and underscores (_). It cannot start with ALIYUN, ACS, ALIBABA, or ALICLOUD.
	//
	// This parameter is required.
	//
	// example:
	//
	// MyTemplate
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
}

func (s DeleteTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteTemplateRequest) GoString() string {
	return s.String()
}

func (s *DeleteTemplateRequest) SetAutoDeleteExecutions(v bool) *DeleteTemplateRequest {
	s.AutoDeleteExecutions = &v
	return s
}

func (s *DeleteTemplateRequest) SetRegionId(v string) *DeleteTemplateRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteTemplateRequest) SetTemplateName(v string) *DeleteTemplateRequest {
	s.TemplateName = &v
	return s
}

type DeleteTemplateResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 2075899A-585D-4A41-A9B2-28DA8534F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteTemplateResponseBody) SetRequestId(v string) *DeleteTemplateResponseBody {
	s.RequestId = &v
	return s
}

type DeleteTemplateResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteTemplateResponse) GoString() string {
	return s.String()
}

func (s *DeleteTemplateResponse) SetHeaders(v map[string]*string) *DeleteTemplateResponse {
	s.Headers = v
	return s
}

func (s *DeleteTemplateResponse) SetStatusCode(v int32) *DeleteTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteTemplateResponse) SetBody(v *DeleteTemplateResponseBody) *DeleteTemplateResponse {
	s.Body = v
	return s
}

type DeleteTemplatesRequest struct {
	// Specifies whether to delete the related executions when a template is deleted.
	//
	// example:
	//
	// false
	AutoDeleteExecutions *bool `json:"AutoDeleteExecutions,omitempty" xml:"AutoDeleteExecutions,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The names of the templates to be deleted.
	//
	// This parameter is required.
	//
	// example:
	//
	// ["t1","t2"]
	TemplateNames *string `json:"TemplateNames,omitempty" xml:"TemplateNames,omitempty"`
}

func (s DeleteTemplatesRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteTemplatesRequest) GoString() string {
	return s.String()
}

func (s *DeleteTemplatesRequest) SetAutoDeleteExecutions(v bool) *DeleteTemplatesRequest {
	s.AutoDeleteExecutions = &v
	return s
}

func (s *DeleteTemplatesRequest) SetRegionId(v string) *DeleteTemplatesRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteTemplatesRequest) SetTemplateNames(v string) *DeleteTemplatesRequest {
	s.TemplateNames = &v
	return s
}

type DeleteTemplatesResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 2075899A-585D-4A41-A9B2-28DA8534
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteTemplatesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteTemplatesResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteTemplatesResponseBody) SetRequestId(v string) *DeleteTemplatesResponseBody {
	s.RequestId = &v
	return s
}

type DeleteTemplatesResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteTemplatesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteTemplatesResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteTemplatesResponse) GoString() string {
	return s.String()
}

func (s *DeleteTemplatesResponse) SetHeaders(v map[string]*string) *DeleteTemplatesResponse {
	s.Headers = v
	return s
}

func (s *DeleteTemplatesResponse) SetStatusCode(v int32) *DeleteTemplatesResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteTemplatesResponse) SetBody(v *DeleteTemplatesResponseBody) *DeleteTemplatesResponse {
	s.Body = v
	return s
}

type DeployApplicationGroupRequest struct {
	// The name of the application.
	//
	// This parameter is required.
	//
	// example:
	//
	// MyApplication
	ApplicationName *string `json:"ApplicationName,omitempty" xml:"ApplicationName,omitempty"`
	// The deployment information about the application group.
	//
	// This parameter is required.
	//
	// example:
	//
	// {       "TemplateURL": "https://ros-template.oss-cn-zhangjiakou.aliyuncs.com/App_Management_Existing_Vpc_Ecs_Instance.json",       "Parameters": {         "ZoneId": "cn-hangzhou-k",         "ProjectName": "test",         "SystemDiskSize": 40,         "InstanceChargeType": "PostPaid",         "SecurityGroupId": "sg-bp1a4374akk63jl8tddy",         "VSwitchId": "vsw-bp1fcvc3zn0jrag86rrlm",         "SystemDiskCategory": "cloud_essd",         "InstancePassword": "******",         "InternetChargeType": "PayByTraffic",         "InstanceCount": 1,         "InternetMaxBandwidthOut": 0,         "VpcId": "vpc-bp1i99boyas8i8m9t3skp",         "EcsImageId": "centos_8_5_x64_20G_alibase_20211228.vhd",         "DataDiskSize": 100,         "EcsInstanceType": "ecs.s6-c1m4.small",         "DataDiskCategory": "cloud_efficiency",         "EnvironmentCommandId": "c-hz028fc3g031gcg"       }
	DeployParameters *string `json:"DeployParameters,omitempty" xml:"DeployParameters,omitempty"`
	// The name of the application group.
	//
	// This parameter is required.
	//
	// example:
	//
	// MyApplicationGroup
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the region in which you want to deploy the application group.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeployApplicationGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s DeployApplicationGroupRequest) GoString() string {
	return s.String()
}

func (s *DeployApplicationGroupRequest) SetApplicationName(v string) *DeployApplicationGroupRequest {
	s.ApplicationName = &v
	return s
}

func (s *DeployApplicationGroupRequest) SetDeployParameters(v string) *DeployApplicationGroupRequest {
	s.DeployParameters = &v
	return s
}

func (s *DeployApplicationGroupRequest) SetName(v string) *DeployApplicationGroupRequest {
	s.Name = &v
	return s
}

func (s *DeployApplicationGroupRequest) SetRegionId(v string) *DeployApplicationGroupRequest {
	s.RegionId = &v
	return s
}

type DeployApplicationGroupResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 8AF4800A-A316-589A-90C4-313B1FEEB084
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeployApplicationGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeployApplicationGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeployApplicationGroupResponseBody) SetRequestId(v string) *DeployApplicationGroupResponseBody {
	s.RequestId = &v
	return s
}

type DeployApplicationGroupResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeployApplicationGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeployApplicationGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s DeployApplicationGroupResponse) GoString() string {
	return s.String()
}

func (s *DeployApplicationGroupResponse) SetHeaders(v map[string]*string) *DeployApplicationGroupResponse {
	s.Headers = v
	return s
}

func (s *DeployApplicationGroupResponse) SetStatusCode(v int32) *DeployApplicationGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DeployApplicationGroupResponse) SetBody(v *DeployApplicationGroupResponseBody) *DeployApplicationGroupResponse {
	s.Body = v
	return s
}

type DescribeApplicationGroupBillRequest struct {
	// The application name.
	//
	// This parameter is required.
	//
	// example:
	//
	// test_application
	ApplicationName *string `json:"ApplicationName,omitempty" xml:"ApplicationName,omitempty"`
	// The billing cycle, in the YYYY-MM format.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2023-06
	BillingCycle *string `json:"BillingCycle,omitempty" xml:"BillingCycle,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The application group name.
	//
	// This parameter is required.
	//
	// example:
	//
	// test_application_group
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The token that is used to retrieve the next page of results.
	//
	// example:
	//
	// -
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the region.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The type of the cloud resource.
	//
	// This parameter is required.
	//
	// example:
	//
	// ALIYUN::ECS::INSTANCE
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s DescribeApplicationGroupBillRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeApplicationGroupBillRequest) GoString() string {
	return s.String()
}

func (s *DescribeApplicationGroupBillRequest) SetApplicationName(v string) *DescribeApplicationGroupBillRequest {
	s.ApplicationName = &v
	return s
}

func (s *DescribeApplicationGroupBillRequest) SetBillingCycle(v string) *DescribeApplicationGroupBillRequest {
	s.BillingCycle = &v
	return s
}

func (s *DescribeApplicationGroupBillRequest) SetMaxResults(v int32) *DescribeApplicationGroupBillRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeApplicationGroupBillRequest) SetName(v string) *DescribeApplicationGroupBillRequest {
	s.Name = &v
	return s
}

func (s *DescribeApplicationGroupBillRequest) SetNextToken(v string) *DescribeApplicationGroupBillRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeApplicationGroupBillRequest) SetRegionId(v string) *DescribeApplicationGroupBillRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeApplicationGroupBillRequest) SetResourceType(v string) *DescribeApplicationGroupBillRequest {
	s.ResourceType = &v
	return s
}

type DescribeApplicationGroupBillResponseBody struct {
	// The consume of application group.
	ApplicationGroupConsume []*DescribeApplicationGroupBillResponseBodyApplicationGroupConsume `json:"ApplicationGroupConsume,omitempty" xml:"ApplicationGroupConsume,omitempty" type:"Repeated"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that is used to retrieve the next page of results.
	//
	// example:
	//
	// ""
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// E897A1AB-4701-5B71-93AD-38FD703763A3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeApplicationGroupBillResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeApplicationGroupBillResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeApplicationGroupBillResponseBody) SetApplicationGroupConsume(v []*DescribeApplicationGroupBillResponseBodyApplicationGroupConsume) *DescribeApplicationGroupBillResponseBody {
	s.ApplicationGroupConsume = v
	return s
}

func (s *DescribeApplicationGroupBillResponseBody) SetMaxResults(v int32) *DescribeApplicationGroupBillResponseBody {
	s.MaxResults = &v
	return s
}

func (s *DescribeApplicationGroupBillResponseBody) SetNextToken(v string) *DescribeApplicationGroupBillResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeApplicationGroupBillResponseBody) SetRequestId(v string) *DescribeApplicationGroupBillResponseBody {
	s.RequestId = &v
	return s
}

type DescribeApplicationGroupBillResponseBodyApplicationGroupConsume struct {
	// The amount consumed by the instance.
	//
	// example:
	//
	// 18.88
	Amount *float32 `json:"Amount,omitempty" xml:"Amount,omitempty"`
	// The time when the instance was created.
	//
	// example:
	//
	// 2023-06-10T06:00Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The currency unit.
	//
	// example:
	//
	// CNY
	Currency *string `json:"Currency,omitempty" xml:"Currency,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// i-0jl781czrhqey0s5zpsj
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the instance.
	//
	// example:
	//
	// test-
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The instance type.
	//
	// example:
	//
	// ALIYUN::ECS::INSTANCE
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// Optimization suggestions.
	//
	// example:
	//
	// 1
	Optimization *string `json:"Optimization,omitempty" xml:"Optimization,omitempty"`
	// The peak type.
	//
	// example:
	//
	// WHITE
	PeakType *string `json:"PeakType,omitempty" xml:"PeakType,omitempty"`
	// The performance of the data synchronization instance.
	//
	// example:
	//
	// "{\\"mem\\":\\"6.82\\",\\"cpu\\":\\"0.55\\"}"
	Performance *string `json:"Performance,omitempty" xml:"Performance,omitempty"`
	// The status of instance.
	//
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeApplicationGroupBillResponseBodyApplicationGroupConsume) String() string {
	return tea.Prettify(s)
}

func (s DescribeApplicationGroupBillResponseBodyApplicationGroupConsume) GoString() string {
	return s.String()
}

func (s *DescribeApplicationGroupBillResponseBodyApplicationGroupConsume) SetAmount(v float32) *DescribeApplicationGroupBillResponseBodyApplicationGroupConsume {
	s.Amount = &v
	return s
}

func (s *DescribeApplicationGroupBillResponseBodyApplicationGroupConsume) SetCreationTime(v string) *DescribeApplicationGroupBillResponseBodyApplicationGroupConsume {
	s.CreationTime = &v
	return s
}

func (s *DescribeApplicationGroupBillResponseBodyApplicationGroupConsume) SetCurrency(v string) *DescribeApplicationGroupBillResponseBodyApplicationGroupConsume {
	s.Currency = &v
	return s
}

func (s *DescribeApplicationGroupBillResponseBodyApplicationGroupConsume) SetInstanceId(v string) *DescribeApplicationGroupBillResponseBodyApplicationGroupConsume {
	s.InstanceId = &v
	return s
}

func (s *DescribeApplicationGroupBillResponseBodyApplicationGroupConsume) SetInstanceName(v string) *DescribeApplicationGroupBillResponseBodyApplicationGroupConsume {
	s.InstanceName = &v
	return s
}

func (s *DescribeApplicationGroupBillResponseBodyApplicationGroupConsume) SetInstanceType(v string) *DescribeApplicationGroupBillResponseBodyApplicationGroupConsume {
	s.InstanceType = &v
	return s
}

func (s *DescribeApplicationGroupBillResponseBodyApplicationGroupConsume) SetOptimization(v string) *DescribeApplicationGroupBillResponseBodyApplicationGroupConsume {
	s.Optimization = &v
	return s
}

func (s *DescribeApplicationGroupBillResponseBodyApplicationGroupConsume) SetPeakType(v string) *DescribeApplicationGroupBillResponseBodyApplicationGroupConsume {
	s.PeakType = &v
	return s
}

func (s *DescribeApplicationGroupBillResponseBodyApplicationGroupConsume) SetPerformance(v string) *DescribeApplicationGroupBillResponseBodyApplicationGroupConsume {
	s.Performance = &v
	return s
}

func (s *DescribeApplicationGroupBillResponseBodyApplicationGroupConsume) SetStatus(v string) *DescribeApplicationGroupBillResponseBodyApplicationGroupConsume {
	s.Status = &v
	return s
}

type DescribeApplicationGroupBillResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeApplicationGroupBillResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeApplicationGroupBillResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeApplicationGroupBillResponse) GoString() string {
	return s.String()
}

func (s *DescribeApplicationGroupBillResponse) SetHeaders(v map[string]*string) *DescribeApplicationGroupBillResponse {
	s.Headers = v
	return s
}

func (s *DescribeApplicationGroupBillResponse) SetStatusCode(v int32) *DescribeApplicationGroupBillResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeApplicationGroupBillResponse) SetBody(v *DescribeApplicationGroupBillResponseBody) *DescribeApplicationGroupBillResponse {
	s.Body = v
	return s
}

type DescribeRegionsRequest struct {
	// The supported natural language. Valid values:
	//
	// 	- zh-CN: Chinese
	//
	// 	- en-US: English
	//
	// example:
	//
	// zh-CN
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeRegionsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsRequest) GoString() string {
	return s.String()
}

func (s *DescribeRegionsRequest) SetAcceptLanguage(v string) *DescribeRegionsRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *DescribeRegionsRequest) SetRegionId(v string) *DescribeRegionsRequest {
	s.RegionId = &v
	return s
}

type DescribeRegionsResponseBody struct {
	// The details of the regions.
	Regions []*DescribeRegionsResponseBodyRegions `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 469E79B1-90A3-4A57-B7C4-2FE0C8B5175E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeRegionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBody) SetRegions(v []*DescribeRegionsResponseBodyRegions) *DescribeRegionsResponseBody {
	s.Regions = v
	return s
}

func (s *DescribeRegionsResponseBody) SetRequestId(v string) *DescribeRegionsResponseBody {
	s.RequestId = &v
	return s
}

type DescribeRegionsResponseBodyRegions struct {
	// The name of the region.
	//
	// example:
	//
	// China (Shenzhen)
	LocalName *string `json:"LocalName,omitempty" xml:"LocalName,omitempty"`
	// The endpoint of the region.
	//
	// example:
	//
	// oos.cn-shenzhen.aliyuncs.com
	RegionEndpoint *string `json:"RegionEndpoint,omitempty" xml:"RegionEndpoint,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-shenzhen
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeRegionsResponseBodyRegions) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBodyRegions) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBodyRegions) SetLocalName(v string) *DescribeRegionsResponseBodyRegions {
	s.LocalName = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegions) SetRegionEndpoint(v string) *DescribeRegionsResponseBodyRegions {
	s.RegionEndpoint = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegions) SetRegionId(v string) *DescribeRegionsResponseBodyRegions {
	s.RegionId = &v
	return s
}

type DescribeRegionsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRegionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRegionsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponse) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponse) SetHeaders(v map[string]*string) *DescribeRegionsResponse {
	s.Headers = v
	return s
}

func (s *DescribeRegionsResponse) SetStatusCode(v int32) *DescribeRegionsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRegionsResponse) SetBody(v *DescribeRegionsResponseBody) *DescribeRegionsResponse {
	s.Body = v
	return s
}

type GenerateExecutionPolicyRequest struct {
	// The RAM role.
	//
	// example:
	//
	// AliyunServiceRoleForOOSBandwidthScheduler
	RamRole *string `json:"RamRole,omitempty" xml:"RamRole,omitempty"`
	// The ID of the region.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The content of the template in the JSON or YAML format. This parameter is the same as the Content parameter that you can specify when you call the CreateTemplate operation. You can use this parameter to specify the tasks that you want to run. This way, you do not need to create a template before you start an execution. If you select an existing template, you do not need to specify this parameter.
	//
	// example:
	//
	// {   "Description": "Example template, describe instances in some status",   "FormatVersion": "OOS-2019-06-01",   "Parameters": {},   "Tasks": [     {       "Name": "describeInstances",       "Action": "ACS::ExecuteAPI",       "Description": "desc-en",       "Properties": {         "Service": "ECS",         "API": "DescribeInstances",         "Parameters": {           "Status": "Running"         }       }     }   ] }
	TemplateContent *string `json:"TemplateContent,omitempty" xml:"TemplateContent,omitempty"`
	// The name of the template.
	//
	// example:
	//
	// vmeixme
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// The version of the template. The default value is the latest version of the template.
	//
	// example:
	//
	// v2
	TemplateVersion *string `json:"TemplateVersion,omitempty" xml:"TemplateVersion,omitempty"`
}

func (s GenerateExecutionPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s GenerateExecutionPolicyRequest) GoString() string {
	return s.String()
}

func (s *GenerateExecutionPolicyRequest) SetRamRole(v string) *GenerateExecutionPolicyRequest {
	s.RamRole = &v
	return s
}

func (s *GenerateExecutionPolicyRequest) SetRegionId(v string) *GenerateExecutionPolicyRequest {
	s.RegionId = &v
	return s
}

func (s *GenerateExecutionPolicyRequest) SetTemplateContent(v string) *GenerateExecutionPolicyRequest {
	s.TemplateContent = &v
	return s
}

func (s *GenerateExecutionPolicyRequest) SetTemplateName(v string) *GenerateExecutionPolicyRequest {
	s.TemplateName = &v
	return s
}

func (s *GenerateExecutionPolicyRequest) SetTemplateVersion(v string) *GenerateExecutionPolicyRequest {
	s.TemplateVersion = &v
	return s
}

type GenerateExecutionPolicyResponseBody struct {
	// The policies that are missing.
	//
	// example:
	//
	// [{\\"Action\\": [\\"ecs:DescribeInvocationResults\\", \\"ecs:DescribeInstances\\", \\"ecs:RunCommand\\", \\"ecs:DescribeInvocations\\"], \\"ServiceName\\": \\"ecs\\", \\"Resources\\": \\"*\\"}]
	MissingPolicy *string `json:"MissingPolicy,omitempty" xml:"MissingPolicy,omitempty"`
	// The RAM policy.
	//
	// example:
	//
	// {}
	Policy *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 14A07460-EBE7-47CA-9757-12CC4761D47A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GenerateExecutionPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GenerateExecutionPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateExecutionPolicyResponseBody) SetMissingPolicy(v string) *GenerateExecutionPolicyResponseBody {
	s.MissingPolicy = &v
	return s
}

func (s *GenerateExecutionPolicyResponseBody) SetPolicy(v string) *GenerateExecutionPolicyResponseBody {
	s.Policy = &v
	return s
}

func (s *GenerateExecutionPolicyResponseBody) SetRequestId(v string) *GenerateExecutionPolicyResponseBody {
	s.RequestId = &v
	return s
}

type GenerateExecutionPolicyResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GenerateExecutionPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GenerateExecutionPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s GenerateExecutionPolicyResponse) GoString() string {
	return s.String()
}

func (s *GenerateExecutionPolicyResponse) SetHeaders(v map[string]*string) *GenerateExecutionPolicyResponse {
	s.Headers = v
	return s
}

func (s *GenerateExecutionPolicyResponse) SetStatusCode(v int32) *GenerateExecutionPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *GenerateExecutionPolicyResponse) SetBody(v *GenerateExecutionPolicyResponseBody) *GenerateExecutionPolicyResponse {
	s.Body = v
	return s
}

type GetApplicationRequest struct {
	// The application name.
	//
	// This parameter is required.
	//
	// example:
	//
	// MyApplication
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The region ID. Set the value to cn-hangzhou.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetApplicationRequest) String() string {
	return tea.Prettify(s)
}

func (s GetApplicationRequest) GoString() string {
	return s.String()
}

func (s *GetApplicationRequest) SetName(v string) *GetApplicationRequest {
	s.Name = &v
	return s
}

func (s *GetApplicationRequest) SetRegionId(v string) *GetApplicationRequest {
	s.RegionId = &v
	return s
}

type GetApplicationResponseBody struct {
	// The information about the application.
	Application *GetApplicationResponseBodyApplication `json:"Application,omitempty" xml:"Application,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 51004B8A-6D9A-5ACB-9158-6C6794496AD0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetApplicationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetApplicationResponseBody) GoString() string {
	return s.String()
}

func (s *GetApplicationResponseBody) SetApplication(v *GetApplicationResponseBodyApplication) *GetApplicationResponseBody {
	s.Application = v
	return s
}

func (s *GetApplicationResponseBody) SetRequestId(v string) *GetApplicationResponseBody {
	s.RequestId = &v
	return s
}

type GetApplicationResponseBodyApplication struct {
	// The configurations of application alerts.
	AlarmConfig *GetApplicationResponseBodyApplicationAlarmConfig `json:"AlarmConfig,omitempty" xml:"AlarmConfig,omitempty" type:"Struct"`
	// The type of the application.
	//
	// Valid values:
	//
	// 	- ComputeNest
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- Custom
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- DingTalk
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// DingTalk
	ApplicationType *string `json:"ApplicationType,omitempty" xml:"ApplicationType,omitempty"`
	// The time when the application was created.
	//
	// example:
	//
	// 2021-09-07T09:17:46Z
	CreateDate *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	// The description of the application.
	//
	// example:
	//
	// Application
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The application name.
	//
	// example:
	//
	// MyApplication
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfmxsn4m******
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The ID of the service.
	//
	// example:
	//
	// service-79538e30e44541b699d8
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// The tags.
	//
	// example:
	//
	// {"k1": "v1", "k2": "v2"}
	Tags map[string]interface{} `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The time when the application was updated.
	//
	// example:
	//
	// 2021-09-07T09:17:46Z
	UpdateDate *string `json:"UpdateDate,omitempty" xml:"UpdateDate,omitempty"`
}

func (s GetApplicationResponseBodyApplication) String() string {
	return tea.Prettify(s)
}

func (s GetApplicationResponseBodyApplication) GoString() string {
	return s.String()
}

func (s *GetApplicationResponseBodyApplication) SetAlarmConfig(v *GetApplicationResponseBodyApplicationAlarmConfig) *GetApplicationResponseBodyApplication {
	s.AlarmConfig = v
	return s
}

func (s *GetApplicationResponseBodyApplication) SetApplicationType(v string) *GetApplicationResponseBodyApplication {
	s.ApplicationType = &v
	return s
}

func (s *GetApplicationResponseBodyApplication) SetCreateDate(v string) *GetApplicationResponseBodyApplication {
	s.CreateDate = &v
	return s
}

func (s *GetApplicationResponseBodyApplication) SetDescription(v string) *GetApplicationResponseBodyApplication {
	s.Description = &v
	return s
}

func (s *GetApplicationResponseBodyApplication) SetName(v string) *GetApplicationResponseBodyApplication {
	s.Name = &v
	return s
}

func (s *GetApplicationResponseBodyApplication) SetResourceGroupId(v string) *GetApplicationResponseBodyApplication {
	s.ResourceGroupId = &v
	return s
}

func (s *GetApplicationResponseBodyApplication) SetServiceId(v string) *GetApplicationResponseBodyApplication {
	s.ServiceId = &v
	return s
}

func (s *GetApplicationResponseBodyApplication) SetTags(v map[string]interface{}) *GetApplicationResponseBodyApplication {
	s.Tags = v
	return s
}

func (s *GetApplicationResponseBodyApplication) SetUpdateDate(v string) *GetApplicationResponseBodyApplication {
	s.UpdateDate = &v
	return s
}

type GetApplicationResponseBodyApplicationAlarmConfig struct {
	// The alert contact list.
	ContactGroups []*string `json:"ContactGroups,omitempty" xml:"ContactGroups,omitempty" type:"Repeated"`
	// The health check URL of the application.
	//
	// example:
	//
	// /api/health/
	HealthCheckUrl *string `json:"HealthCheckUrl,omitempty" xml:"HealthCheckUrl,omitempty"`
	// The ID of the alert template.
	TemplateIds []*string `json:"TemplateIds,omitempty" xml:"TemplateIds,omitempty" type:"Repeated"`
}

func (s GetApplicationResponseBodyApplicationAlarmConfig) String() string {
	return tea.Prettify(s)
}

func (s GetApplicationResponseBodyApplicationAlarmConfig) GoString() string {
	return s.String()
}

func (s *GetApplicationResponseBodyApplicationAlarmConfig) SetContactGroups(v []*string) *GetApplicationResponseBodyApplicationAlarmConfig {
	s.ContactGroups = v
	return s
}

func (s *GetApplicationResponseBodyApplicationAlarmConfig) SetHealthCheckUrl(v string) *GetApplicationResponseBodyApplicationAlarmConfig {
	s.HealthCheckUrl = &v
	return s
}

func (s *GetApplicationResponseBodyApplicationAlarmConfig) SetTemplateIds(v []*string) *GetApplicationResponseBodyApplicationAlarmConfig {
	s.TemplateIds = v
	return s
}

type GetApplicationResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetApplicationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetApplicationResponse) String() string {
	return tea.Prettify(s)
}

func (s GetApplicationResponse) GoString() string {
	return s.String()
}

func (s *GetApplicationResponse) SetHeaders(v map[string]*string) *GetApplicationResponse {
	s.Headers = v
	return s
}

func (s *GetApplicationResponse) SetStatusCode(v int32) *GetApplicationResponse {
	s.StatusCode = &v
	return s
}

func (s *GetApplicationResponse) SetBody(v *GetApplicationResponseBody) *GetApplicationResponse {
	s.Body = v
	return s
}

type GetApplicationGroupRequest struct {
	// The name of the application.
	//
	// This parameter is required.
	//
	// example:
	//
	// MyApplication
	ApplicationName *string `json:"ApplicationName,omitempty" xml:"ApplicationName,omitempty"`
	// The name of the application group.
	//
	// This parameter is required.
	//
	// example:
	//
	// MyApplicationGroup
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the region. Set the value to cn-hangzhou.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetApplicationGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s GetApplicationGroupRequest) GoString() string {
	return s.String()
}

func (s *GetApplicationGroupRequest) SetApplicationName(v string) *GetApplicationGroupRequest {
	s.ApplicationName = &v
	return s
}

func (s *GetApplicationGroupRequest) SetName(v string) *GetApplicationGroupRequest {
	s.Name = &v
	return s
}

func (s *GetApplicationGroupRequest) SetRegionId(v string) *GetApplicationGroupRequest {
	s.RegionId = &v
	return s
}

type GetApplicationGroupResponseBody struct {
	// The details of the application group.
	ApplicationGroup *GetApplicationGroupResponseBodyApplicationGroup `json:"ApplicationGroup,omitempty" xml:"ApplicationGroup,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 92EA60ED-544D-55E9-93D9-237BE9976C0D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetApplicationGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetApplicationGroupResponseBody) GoString() string {
	return s.String()
}

func (s *GetApplicationGroupResponseBody) SetApplicationGroup(v *GetApplicationGroupResponseBodyApplicationGroup) *GetApplicationGroupResponseBody {
	s.ApplicationGroup = v
	return s
}

func (s *GetApplicationGroupResponseBody) SetRequestId(v string) *GetApplicationGroupResponseBody {
	s.RequestId = &v
	return s
}

type GetApplicationGroupResponseBodyApplicationGroup struct {
	// The name of the application.
	//
	// example:
	//
	// MyApplication
	ApplicationName *string `json:"ApplicationName,omitempty" xml:"ApplicationName,omitempty"`
	// The ID of the application group in CloudMonitor.
	//
	// example:
	//
	// 12345678
	CmsGroupId *string `json:"CmsGroupId,omitempty" xml:"CmsGroupId,omitempty"`
	// The time when the application group was created.
	//
	// example:
	//
	// 2021-09-07T10:28:25Z
	CreateDate *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	// The output of the deployment result.
	//
	// example:
	//
	// {       "Outputs": [         {           "Description": "No description given",           "OutputKey": "InstanceIds"         }       ],       "StackId": "6ef4b860-f6e7-4145-8d22-f4a2a32363e1"     }   }
	DeployOutputs *string `json:"DeployOutputs,omitempty" xml:"DeployOutputs,omitempty"`
	// The configuration information of the application group.
	//
	// example:
	//
	// {       "TemplateURL": "https://ros-template.oss-cn-zhangjiakou.aliyuncs.com/App_Management_Existing_Vpc_Ecs_Instance.json",       "Parameters": {         "ZoneId": "cn-hangzhou-k",         "ProjectName": "test",         "SystemDiskSize": 40,         "InstanceChargeType": "PostPaid",         "SecurityGroupId": "sg-bp1a4374akk63jl8tddy",         "VSwitchId": "vsw-bp1fcvc3zn0jrag86rrlm",         "SystemDiskCategory": "cloud_essd",         "InstancePassword": "******",         "InternetChargeType": "PayByTraffic",         "InstanceCount": 1,         "InternetMaxBandwidthOut": 0,         "VpcId": "vpc-bp1i99boyas8i8m9t3skp",         "EcsImageId": "centos_8_5_x64_20G_alibase_20211228.vhd",         "DataDiskSize": 100,         "EcsInstanceType": "ecs.s6-c1m4.small",         "DataDiskCategory": "cloud_efficiency",         "EnvironmentCommandId": "c-hz028fc3g031gcg"       }
	DeployParameters *string `json:"DeployParameters,omitempty" xml:"DeployParameters,omitempty"`
	// The ID of the region in which you deploy the application group.
	//
	// example:
	//
	// cn-hangzhou
	DeployRegionId *string `json:"DeployRegionId,omitempty" xml:"DeployRegionId,omitempty"`
	// The description of the application group.
	//
	// example:
	//
	// ApplicationGroup
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The tag key.
	//
	// example:
	//
	// k1
	ImportTagKey *string `json:"ImportTagKey,omitempty" xml:"ImportTagKey,omitempty"`
	// The tag value.
	//
	// example:
	//
	// v1
	ImportTagValue *string `json:"ImportTagValue,omitempty" xml:"ImportTagValue,omitempty"`
	// The name of the application group.
	//
	// example:
	//
	// MyApplicationGroup
	Name              *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OperationMetadata *string `json:"OperationMetadata,omitempty" xml:"OperationMetadata,omitempty"`
	// The creation progress of the application instance.
	//
	// example:
	//
	// 100
	Progress *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// The state of the application group.
	//
	// example:
	//
	// Created
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The state information of the application group.
	//
	// example:
	//
	// Stack CREATE completed successfully
	StatusReason *string `json:"StatusReason,omitempty" xml:"StatusReason,omitempty"`
	// The time when the application group was last modified.
	//
	// example:
	//
	// 2021-09-07T10:28:25Z
	UpdateDate *string `json:"UpdateDate,omitempty" xml:"UpdateDate,omitempty"`
}

func (s GetApplicationGroupResponseBodyApplicationGroup) String() string {
	return tea.Prettify(s)
}

func (s GetApplicationGroupResponseBodyApplicationGroup) GoString() string {
	return s.String()
}

func (s *GetApplicationGroupResponseBodyApplicationGroup) SetApplicationName(v string) *GetApplicationGroupResponseBodyApplicationGroup {
	s.ApplicationName = &v
	return s
}

func (s *GetApplicationGroupResponseBodyApplicationGroup) SetCmsGroupId(v string) *GetApplicationGroupResponseBodyApplicationGroup {
	s.CmsGroupId = &v
	return s
}

func (s *GetApplicationGroupResponseBodyApplicationGroup) SetCreateDate(v string) *GetApplicationGroupResponseBodyApplicationGroup {
	s.CreateDate = &v
	return s
}

func (s *GetApplicationGroupResponseBodyApplicationGroup) SetDeployOutputs(v string) *GetApplicationGroupResponseBodyApplicationGroup {
	s.DeployOutputs = &v
	return s
}

func (s *GetApplicationGroupResponseBodyApplicationGroup) SetDeployParameters(v string) *GetApplicationGroupResponseBodyApplicationGroup {
	s.DeployParameters = &v
	return s
}

func (s *GetApplicationGroupResponseBodyApplicationGroup) SetDeployRegionId(v string) *GetApplicationGroupResponseBodyApplicationGroup {
	s.DeployRegionId = &v
	return s
}

func (s *GetApplicationGroupResponseBodyApplicationGroup) SetDescription(v string) *GetApplicationGroupResponseBodyApplicationGroup {
	s.Description = &v
	return s
}

func (s *GetApplicationGroupResponseBodyApplicationGroup) SetImportTagKey(v string) *GetApplicationGroupResponseBodyApplicationGroup {
	s.ImportTagKey = &v
	return s
}

func (s *GetApplicationGroupResponseBodyApplicationGroup) SetImportTagValue(v string) *GetApplicationGroupResponseBodyApplicationGroup {
	s.ImportTagValue = &v
	return s
}

func (s *GetApplicationGroupResponseBodyApplicationGroup) SetName(v string) *GetApplicationGroupResponseBodyApplicationGroup {
	s.Name = &v
	return s
}

func (s *GetApplicationGroupResponseBodyApplicationGroup) SetOperationMetadata(v string) *GetApplicationGroupResponseBodyApplicationGroup {
	s.OperationMetadata = &v
	return s
}

func (s *GetApplicationGroupResponseBodyApplicationGroup) SetProgress(v string) *GetApplicationGroupResponseBodyApplicationGroup {
	s.Progress = &v
	return s
}

func (s *GetApplicationGroupResponseBodyApplicationGroup) SetStatus(v string) *GetApplicationGroupResponseBodyApplicationGroup {
	s.Status = &v
	return s
}

func (s *GetApplicationGroupResponseBodyApplicationGroup) SetStatusReason(v string) *GetApplicationGroupResponseBodyApplicationGroup {
	s.StatusReason = &v
	return s
}

func (s *GetApplicationGroupResponseBodyApplicationGroup) SetUpdateDate(v string) *GetApplicationGroupResponseBodyApplicationGroup {
	s.UpdateDate = &v
	return s
}

type GetApplicationGroupResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetApplicationGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetApplicationGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s GetApplicationGroupResponse) GoString() string {
	return s.String()
}

func (s *GetApplicationGroupResponse) SetHeaders(v map[string]*string) *GetApplicationGroupResponse {
	s.Headers = v
	return s
}

func (s *GetApplicationGroupResponse) SetStatusCode(v int32) *GetApplicationGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *GetApplicationGroupResponse) SetBody(v *GetApplicationGroupResponseBody) *GetApplicationGroupResponse {
	s.Body = v
	return s
}

type GetExecutionTemplateRequest struct {
	// The ID of the execution.
	//
	// This parameter is required.
	//
	// example:
	//
	// exec-046490ff88f242
	ExecutionId *string `json:"ExecutionId,omitempty" xml:"ExecutionId,omitempty"`
	// The ID of the region.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetExecutionTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s GetExecutionTemplateRequest) GoString() string {
	return s.String()
}

func (s *GetExecutionTemplateRequest) SetExecutionId(v string) *GetExecutionTemplateRequest {
	s.ExecutionId = &v
	return s
}

func (s *GetExecutionTemplateRequest) SetRegionId(v string) *GetExecutionTemplateRequest {
	s.RegionId = &v
	return s
}

type GetExecutionTemplateResponseBody struct {
	// The content of the template.
	//
	// example:
	//
	// "{\\n \\"FormatVersion\\": \\"OOS-2019-06-01\\",\\n \\"Parameters\\": {\\n \\"Status\\": {\\n \\"Type\\": \\"String\\",\\n \\"Description\\": \\"(Required) The ID of the ECS instance.\\"\\n }\\n },\\n \\"Tasks\\": [\\n {\\n \\"Name\\": \\"bar\\",\\n \\"Properties\\": {\\n \\"Parameters\\": {\\n \\"Status\\": \\"{{ Status }}\\"\\n },\\n \\"API\\": \\"DescribeInstances\\",\\n \\"Service\\": \\"Ecs\\"\\n },\\n \\"Action\\": \\"acs::ExecuteAPI\\",\\n \\"Outputs\\": {\\n \\"InstanceIds\\", {\\n \\"ValueSelector\\": \\".Instances.Instance[].InstanceId\\",\\n \\"Type\\": \\"List\\"\\n }\\n }\\n }\\n ],\\n \\"Outputs\\": {\\n \\"InstanceIds\\": {\\n \\"Value\\": \\" {{ bar.InstanceIds }} \\",\\n \\"Type\\": \\"List\\"\\n }\\n }\\n}\\n"
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 14A60-EBE7-47CA-9757-12C1D47A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The metadata of the template.
	Template *GetExecutionTemplateResponseBodyTemplate `json:"Template,omitempty" xml:"Template,omitempty" type:"Struct"`
}

func (s GetExecutionTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetExecutionTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *GetExecutionTemplateResponseBody) SetContent(v string) *GetExecutionTemplateResponseBody {
	s.Content = &v
	return s
}

func (s *GetExecutionTemplateResponseBody) SetRequestId(v string) *GetExecutionTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetExecutionTemplateResponseBody) SetTemplate(v *GetExecutionTemplateResponseBodyTemplate) *GetExecutionTemplateResponseBody {
	s.Template = v
	return s
}

type GetExecutionTemplateResponseBodyTemplate struct {
	// The creator of the template.
	//
	// example:
	//
	// root(13090000)
	CreatedBy *string `json:"CreatedBy,omitempty" xml:"CreatedBy,omitempty"`
	// The time when the template was created.
	//
	// example:
	//
	// 2019-05-16T10:26:14Z
	CreatedDate *string `json:"CreatedDate,omitempty" xml:"CreatedDate,omitempty"`
	// The description of the template.
	//
	// example:
	//
	// Get status of instances
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The SHA-256 value of the template content.
	//
	// example:
	//
	// 4bc7d7a21b3e003434b9c223f6e6d2578b5ebfeb5be28c1fcf8a8a1b11907bb4
	Hash *string `json:"Hash,omitempty" xml:"Hash,omitempty"`
	// The share type of the template. The share type of a user-created template is **Private**.
	//
	// example:
	//
	// Private
	ShareType *string `json:"ShareType,omitempty" xml:"ShareType,omitempty"`
	// The tag keys and values. The number of key-value pairs ranges from 1 to 20.
	//
	// example:
	//
	// {"k1":"k2","k2":"v2"}
	Tags map[string]interface{} `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The format of the template. The system automatically determines whether the format is JSON or YAML.
	//
	// example:
	//
	// JSON
	TemplateFormat *string `json:"TemplateFormat,omitempty" xml:"TemplateFormat,omitempty"`
	// The ID of the template.
	//
	// example:
	//
	// t-94753d4d828d38
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The name of the template.
	//
	// example:
	//
	// MyTemplate
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// The version of the template. The name of the version consists of the letter v and a number. The number starts from 1.
	//
	// example:
	//
	// v1
	TemplateVersion *string `json:"TemplateVersion,omitempty" xml:"TemplateVersion,omitempty"`
	// The user who last updated the template.
	//
	// example:
	//
	// root(13090000)
	UpdatedBy *string `json:"UpdatedBy,omitempty" xml:"UpdatedBy,omitempty"`
	// The time when the template was last updated.
	//
	// example:
	//
	// 2019-05-16T10:26:14Z
	UpdatedDate *string `json:"UpdatedDate,omitempty" xml:"UpdatedDate,omitempty"`
}

func (s GetExecutionTemplateResponseBodyTemplate) String() string {
	return tea.Prettify(s)
}

func (s GetExecutionTemplateResponseBodyTemplate) GoString() string {
	return s.String()
}

func (s *GetExecutionTemplateResponseBodyTemplate) SetCreatedBy(v string) *GetExecutionTemplateResponseBodyTemplate {
	s.CreatedBy = &v
	return s
}

func (s *GetExecutionTemplateResponseBodyTemplate) SetCreatedDate(v string) *GetExecutionTemplateResponseBodyTemplate {
	s.CreatedDate = &v
	return s
}

func (s *GetExecutionTemplateResponseBodyTemplate) SetDescription(v string) *GetExecutionTemplateResponseBodyTemplate {
	s.Description = &v
	return s
}

func (s *GetExecutionTemplateResponseBodyTemplate) SetHash(v string) *GetExecutionTemplateResponseBodyTemplate {
	s.Hash = &v
	return s
}

func (s *GetExecutionTemplateResponseBodyTemplate) SetShareType(v string) *GetExecutionTemplateResponseBodyTemplate {
	s.ShareType = &v
	return s
}

func (s *GetExecutionTemplateResponseBodyTemplate) SetTags(v map[string]interface{}) *GetExecutionTemplateResponseBodyTemplate {
	s.Tags = v
	return s
}

func (s *GetExecutionTemplateResponseBodyTemplate) SetTemplateFormat(v string) *GetExecutionTemplateResponseBodyTemplate {
	s.TemplateFormat = &v
	return s
}

func (s *GetExecutionTemplateResponseBodyTemplate) SetTemplateId(v string) *GetExecutionTemplateResponseBodyTemplate {
	s.TemplateId = &v
	return s
}

func (s *GetExecutionTemplateResponseBodyTemplate) SetTemplateName(v string) *GetExecutionTemplateResponseBodyTemplate {
	s.TemplateName = &v
	return s
}

func (s *GetExecutionTemplateResponseBodyTemplate) SetTemplateVersion(v string) *GetExecutionTemplateResponseBodyTemplate {
	s.TemplateVersion = &v
	return s
}

func (s *GetExecutionTemplateResponseBodyTemplate) SetUpdatedBy(v string) *GetExecutionTemplateResponseBodyTemplate {
	s.UpdatedBy = &v
	return s
}

func (s *GetExecutionTemplateResponseBodyTemplate) SetUpdatedDate(v string) *GetExecutionTemplateResponseBodyTemplate {
	s.UpdatedDate = &v
	return s
}

type GetExecutionTemplateResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetExecutionTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetExecutionTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s GetExecutionTemplateResponse) GoString() string {
	return s.String()
}

func (s *GetExecutionTemplateResponse) SetHeaders(v map[string]*string) *GetExecutionTemplateResponse {
	s.Headers = v
	return s
}

func (s *GetExecutionTemplateResponse) SetStatusCode(v int32) *GetExecutionTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *GetExecutionTemplateResponse) SetBody(v *GetExecutionTemplateResponseBody) *GetExecutionTemplateResponse {
	s.Body = v
	return s
}

type GetInventorySchemaRequest struct {
	// Specifies whether to return only properties that support the aggregate feature in the configuration list. Valid values:
	//
	// 	- true: only returns properties that support the aggregate feature in the configuration list.
	//
	// 	- false: returns all properties in the configuration list.
	//
	// example:
	//
	// false
	Aggregator *bool `json:"Aggregator,omitempty" xml:"Aggregator,omitempty"`
	// The number of entries per page. Valid values: 1 to 100. Default value: 50.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// A pagination token. It can be used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// gAAAAABfh8MVLQI9AuKGACLgjbsXbWs-Mna47IDM6tr6wK7TZ1
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The configuration list type name. Valid values:
	//
	// 	- ACS:InstanceInformation
	//
	// 	- ACS:Application
	//
	// 	- ACS:File
	//
	// 	- ACS:Network
	//
	// 	- ACS:WindowsRole
	//
	// 	- ACS:Service
	//
	// 	- ACS:WindowsUpdate
	//
	// 	- ACS:WindowsRegistry
	//
	// example:
	//
	// ACS:Application
	TypeName *string `json:"TypeName,omitempty" xml:"TypeName,omitempty"`
}

func (s GetInventorySchemaRequest) String() string {
	return tea.Prettify(s)
}

func (s GetInventorySchemaRequest) GoString() string {
	return s.String()
}

func (s *GetInventorySchemaRequest) SetAggregator(v bool) *GetInventorySchemaRequest {
	s.Aggregator = &v
	return s
}

func (s *GetInventorySchemaRequest) SetMaxResults(v int32) *GetInventorySchemaRequest {
	s.MaxResults = &v
	return s
}

func (s *GetInventorySchemaRequest) SetNextToken(v string) *GetInventorySchemaRequest {
	s.NextToken = &v
	return s
}

func (s *GetInventorySchemaRequest) SetRegionId(v string) *GetInventorySchemaRequest {
	s.RegionId = &v
	return s
}

func (s *GetInventorySchemaRequest) SetTypeName(v string) *GetInventorySchemaRequest {
	s.TypeName = &v
	return s
}

type GetInventorySchemaResponseBody struct {
	// The number of entries per page.
	//
	// example:
	//
	// 1
	MaxResults *string `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that was used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// gAAAAABfh8MVLQI9AuKGACLgjbsXbWs-Mna47IDM6tr6wK7TZ1
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 89117642-7167-4F4D-B7F1-876582279E3E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The detailed configurations of the configuration list.
	Schemas []*GetInventorySchemaResponseBodySchemas `json:"Schemas,omitempty" xml:"Schemas,omitempty" type:"Repeated"`
}

func (s GetInventorySchemaResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetInventorySchemaResponseBody) GoString() string {
	return s.String()
}

func (s *GetInventorySchemaResponseBody) SetMaxResults(v string) *GetInventorySchemaResponseBody {
	s.MaxResults = &v
	return s
}

func (s *GetInventorySchemaResponseBody) SetNextToken(v string) *GetInventorySchemaResponseBody {
	s.NextToken = &v
	return s
}

func (s *GetInventorySchemaResponseBody) SetRequestId(v string) *GetInventorySchemaResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetInventorySchemaResponseBody) SetSchemas(v []*GetInventorySchemaResponseBodySchemas) *GetInventorySchemaResponseBody {
	s.Schemas = v
	return s
}

type GetInventorySchemaResponseBodySchemas struct {
	// The properties of the configuration list.
	Attributes []*GetInventorySchemaResponseBodySchemasAttributes `json:"Attributes,omitempty" xml:"Attributes,omitempty" type:"Repeated"`
	// The name of the configuration list.
	//
	// example:
	//
	// ACS:Application
	TypeName *string `json:"TypeName,omitempty" xml:"TypeName,omitempty"`
	// The version of the configuration list.
	//
	// example:
	//
	// 1.0
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s GetInventorySchemaResponseBodySchemas) String() string {
	return tea.Prettify(s)
}

func (s GetInventorySchemaResponseBodySchemas) GoString() string {
	return s.String()
}

func (s *GetInventorySchemaResponseBodySchemas) SetAttributes(v []*GetInventorySchemaResponseBodySchemasAttributes) *GetInventorySchemaResponseBodySchemas {
	s.Attributes = v
	return s
}

func (s *GetInventorySchemaResponseBodySchemas) SetTypeName(v string) *GetInventorySchemaResponseBodySchemas {
	s.TypeName = &v
	return s
}

func (s *GetInventorySchemaResponseBodySchemas) SetVersion(v string) *GetInventorySchemaResponseBodySchemas {
	s.Version = &v
	return s
}

type GetInventorySchemaResponseBodySchemasAttributes struct {
	// The data type of the property.
	//
	// example:
	//
	// STRING
	DataType *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	// The name of the property.
	//
	// example:
	//
	// ApplicationType
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetInventorySchemaResponseBodySchemasAttributes) String() string {
	return tea.Prettify(s)
}

func (s GetInventorySchemaResponseBodySchemasAttributes) GoString() string {
	return s.String()
}

func (s *GetInventorySchemaResponseBodySchemasAttributes) SetDataType(v string) *GetInventorySchemaResponseBodySchemasAttributes {
	s.DataType = &v
	return s
}

func (s *GetInventorySchemaResponseBodySchemasAttributes) SetName(v string) *GetInventorySchemaResponseBodySchemasAttributes {
	s.Name = &v
	return s
}

type GetInventorySchemaResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetInventorySchemaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetInventorySchemaResponse) String() string {
	return tea.Prettify(s)
}

func (s GetInventorySchemaResponse) GoString() string {
	return s.String()
}

func (s *GetInventorySchemaResponse) SetHeaders(v map[string]*string) *GetInventorySchemaResponse {
	s.Headers = v
	return s
}

func (s *GetInventorySchemaResponse) SetStatusCode(v int32) *GetInventorySchemaResponse {
	s.StatusCode = &v
	return s
}

func (s *GetInventorySchemaResponse) SetBody(v *GetInventorySchemaResponseBody) *GetInventorySchemaResponse {
	s.Body = v
	return s
}

type GetOpsItemRequest struct {
	// The O\\&M item ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// oi-d52b08695e2b46ae8413
	OpsItemId *string `json:"OpsItemId,omitempty" xml:"OpsItemId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetOpsItemRequest) String() string {
	return tea.Prettify(s)
}

func (s GetOpsItemRequest) GoString() string {
	return s.String()
}

func (s *GetOpsItemRequest) SetOpsItemId(v string) *GetOpsItemRequest {
	s.OpsItemId = &v
	return s
}

func (s *GetOpsItemRequest) SetRegionId(v string) *GetOpsItemRequest {
	s.RegionId = &v
	return s
}

type GetOpsItemResponseBody struct {
	// The information about the O\\&M item.
	OpsItem *GetOpsItemResponseBodyOpsItem `json:"OpsItem,omitempty" xml:"OpsItem,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 8BED4C16-BD30-5E27-94D4-7EBCCECF70C1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetOpsItemResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetOpsItemResponseBody) GoString() string {
	return s.String()
}

func (s *GetOpsItemResponseBody) SetOpsItem(v *GetOpsItemResponseBodyOpsItem) *GetOpsItemResponseBody {
	s.OpsItem = v
	return s
}

func (s *GetOpsItemResponseBody) SetRequestId(v string) *GetOpsItemResponseBody {
	s.RequestId = &v
	return s
}

type GetOpsItemResponseBodyOpsItem struct {
	// The information about the attributes of the O\\&M item.
	Attributes map[string]interface{} `json:"Attributes,omitempty" xml:"Attributes,omitempty"`
	// The category of the O\\&M item.
	//
	// example:
	//
	// Security
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The user who created the O\\&M item.
	//
	// example:
	//
	// root(130900000)
	CreateBy *string `json:"CreateBy,omitempty" xml:"CreateBy,omitempty"`
	// The time when the O\\&M item was created.
	//
	// example:
	//
	// 2023-04-10T06:15Z
	CreateDate *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	// The description.
	//
	// example:
	//
	// test-update
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The user who last modified the O\\&M item.
	//
	// example:
	//
	// modifiedBy
	LastModifiedBy *string `json:"LastModifiedBy,omitempty" xml:"LastModifiedBy,omitempty"`
	// The O\\&M item ID.
	//
	// example:
	//
	// oi-d52b08695e2b46ae8413
	OpsItemId *string `json:"OpsItemId,omitempty" xml:"OpsItemId,omitempty"`
	// The priority of the O\\&M item.
	//
	// example:
	//
	// 1
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-aekzxkofnlxtn2i
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The Alibaba Cloud Resource Names (ARNs) of the associated resources.
	Resources []*string `json:"Resources,omitempty" xml:"Resources,omitempty" type:"Repeated"`
	// The severity level of the O\\&M item.
	//
	// example:
	//
	// Medium
	Severity *string `json:"Severity,omitempty" xml:"Severity,omitempty"`
	// The solutions to the O\\&M item.
	Solutions []map[string]interface{} `json:"Solutions,omitempty" xml:"Solutions,omitempty" type:"Repeated"`
	// The source business of the O\\&M item.
	//
	// example:
	//
	// /aliyun/appManager
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The status of the O\\&M item.
	//
	// example:
	//
	// Open
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tags attached to the O\\&M item.
	//
	// example:
	//
	// {"K1":"V1"}
	Tags map[string]interface{} `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The title of the O\\&M item.
	//
	// example:
	//
	// test
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// The time when the O\\&M item was updated.
	//
	// example:
	//
	// 2023-04-10T06:15Z
	UpdateDate *string `json:"UpdateDate,omitempty" xml:"UpdateDate,omitempty"`
}

func (s GetOpsItemResponseBodyOpsItem) String() string {
	return tea.Prettify(s)
}

func (s GetOpsItemResponseBodyOpsItem) GoString() string {
	return s.String()
}

func (s *GetOpsItemResponseBodyOpsItem) SetAttributes(v map[string]interface{}) *GetOpsItemResponseBodyOpsItem {
	s.Attributes = v
	return s
}

func (s *GetOpsItemResponseBodyOpsItem) SetCategory(v string) *GetOpsItemResponseBodyOpsItem {
	s.Category = &v
	return s
}

func (s *GetOpsItemResponseBodyOpsItem) SetCreateBy(v string) *GetOpsItemResponseBodyOpsItem {
	s.CreateBy = &v
	return s
}

func (s *GetOpsItemResponseBodyOpsItem) SetCreateDate(v string) *GetOpsItemResponseBodyOpsItem {
	s.CreateDate = &v
	return s
}

func (s *GetOpsItemResponseBodyOpsItem) SetDescription(v string) *GetOpsItemResponseBodyOpsItem {
	s.Description = &v
	return s
}

func (s *GetOpsItemResponseBodyOpsItem) SetLastModifiedBy(v string) *GetOpsItemResponseBodyOpsItem {
	s.LastModifiedBy = &v
	return s
}

func (s *GetOpsItemResponseBodyOpsItem) SetOpsItemId(v string) *GetOpsItemResponseBodyOpsItem {
	s.OpsItemId = &v
	return s
}

func (s *GetOpsItemResponseBodyOpsItem) SetPriority(v int32) *GetOpsItemResponseBodyOpsItem {
	s.Priority = &v
	return s
}

func (s *GetOpsItemResponseBodyOpsItem) SetResourceGroupId(v string) *GetOpsItemResponseBodyOpsItem {
	s.ResourceGroupId = &v
	return s
}

func (s *GetOpsItemResponseBodyOpsItem) SetResources(v []*string) *GetOpsItemResponseBodyOpsItem {
	s.Resources = v
	return s
}

func (s *GetOpsItemResponseBodyOpsItem) SetSeverity(v string) *GetOpsItemResponseBodyOpsItem {
	s.Severity = &v
	return s
}

func (s *GetOpsItemResponseBodyOpsItem) SetSolutions(v []map[string]interface{}) *GetOpsItemResponseBodyOpsItem {
	s.Solutions = v
	return s
}

func (s *GetOpsItemResponseBodyOpsItem) SetSource(v string) *GetOpsItemResponseBodyOpsItem {
	s.Source = &v
	return s
}

func (s *GetOpsItemResponseBodyOpsItem) SetStatus(v string) *GetOpsItemResponseBodyOpsItem {
	s.Status = &v
	return s
}

func (s *GetOpsItemResponseBodyOpsItem) SetTags(v map[string]interface{}) *GetOpsItemResponseBodyOpsItem {
	s.Tags = v
	return s
}

func (s *GetOpsItemResponseBodyOpsItem) SetTitle(v string) *GetOpsItemResponseBodyOpsItem {
	s.Title = &v
	return s
}

func (s *GetOpsItemResponseBodyOpsItem) SetUpdateDate(v string) *GetOpsItemResponseBodyOpsItem {
	s.UpdateDate = &v
	return s
}

type GetOpsItemResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetOpsItemResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetOpsItemResponse) String() string {
	return tea.Prettify(s)
}

func (s GetOpsItemResponse) GoString() string {
	return s.String()
}

func (s *GetOpsItemResponse) SetHeaders(v map[string]*string) *GetOpsItemResponse {
	s.Headers = v
	return s
}

func (s *GetOpsItemResponse) SetStatusCode(v int32) *GetOpsItemResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOpsItemResponse) SetBody(v *GetOpsItemResponseBody) *GetOpsItemResponse {
	s.Body = v
	return s
}

type GetParameterRequest struct {
	// The name of the common parameter. The name can be up to 200 characters in length and can contain letters, digits, hyphens (-), and underscores (_).
	//
	// This parameter is required.
	//
	// example:
	//
	// MyParameter
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The version number of the common parameter. Valid values: 1 to 100.
	//
	// example:
	//
	// 1
	ParameterVersion *int32 `json:"ParameterVersion,omitempty" xml:"ParameterVersion,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfmxsn4m*****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s GetParameterRequest) String() string {
	return tea.Prettify(s)
}

func (s GetParameterRequest) GoString() string {
	return s.String()
}

func (s *GetParameterRequest) SetName(v string) *GetParameterRequest {
	s.Name = &v
	return s
}

func (s *GetParameterRequest) SetParameterVersion(v int32) *GetParameterRequest {
	s.ParameterVersion = &v
	return s
}

func (s *GetParameterRequest) SetRegionId(v string) *GetParameterRequest {
	s.RegionId = &v
	return s
}

func (s *GetParameterRequest) SetResourceGroupId(v string) *GetParameterRequest {
	s.ResourceGroupId = &v
	return s
}

type GetParameterResponseBody struct {
	// The information about the common parameter.
	Parameter *GetParameterResponseBodyParameter `json:"Parameter,omitempty" xml:"Parameter,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// BA326372-2A10-4C3B-BE3E-6439DB7557CC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetParameterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetParameterResponseBody) GoString() string {
	return s.String()
}

func (s *GetParameterResponseBody) SetParameter(v *GetParameterResponseBodyParameter) *GetParameterResponseBody {
	s.Parameter = v
	return s
}

func (s *GetParameterResponseBody) SetRequestId(v string) *GetParameterResponseBody {
	s.RequestId = &v
	return s
}

type GetParameterResponseBodyParameter struct {
	// The constraints of the common parameter.
	//
	// example:
	//
	// \\"{\\"\\"AllowedValues":["parameter"],"AllowedPattern":"parameter","MinLength":0,"MaxLength":20}\\"
	Constraints *string `json:"Constraints,omitempty" xml:"Constraints,omitempty"`
	// The user who created the common parameter.
	//
	// example:
	//
	// root(130900000)
	CreatedBy *string `json:"CreatedBy,omitempty" xml:"CreatedBy,omitempty"`
	// The time when the common parameter was created.
	//
	// example:
	//
	// 2020-09-01T08:01:43Z
	CreatedDate *string `json:"CreatedDate,omitempty" xml:"CreatedDate,omitempty"`
	// The description of the common parameter.
	//
	// example:
	//
	// parameter
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the common parameter.
	//
	// example:
	//
	// p-4c4b401cab6747xxxxxx
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the common parameter.
	//
	// example:
	//
	// MyParameter
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The version number of the common parameter.
	//
	// example:
	//
	// 1
	ParameterVersion *int32 `json:"ParameterVersion,omitempty" xml:"ParameterVersion,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfmxsn4m*****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The share type of the common parameter.
	//
	// example:
	//
	// Private
	ShareType *string `json:"ShareType,omitempty" xml:"ShareType,omitempty"`
	// The tags added to the common parameter.
	//
	// example:
	//
	// {"k1":"v1","k2":"v2"}
	Tags map[string]interface{} `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The data type of the common parameter.
	//
	// example:
	//
	// String
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The user who updated the common parameter.
	//
	// example:
	//
	// root(130900000)
	UpdatedBy *string `json:"UpdatedBy,omitempty" xml:"UpdatedBy,omitempty"`
	// The time when the common parameter was updated.
	//
	// example:
	//
	// 2020-09-01T08:01:43Z
	UpdatedDate *string `json:"UpdatedDate,omitempty" xml:"UpdatedDate,omitempty"`
	// The value of the common parameter.
	//
	// example:
	//
	// parameter
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetParameterResponseBodyParameter) String() string {
	return tea.Prettify(s)
}

func (s GetParameterResponseBodyParameter) GoString() string {
	return s.String()
}

func (s *GetParameterResponseBodyParameter) SetConstraints(v string) *GetParameterResponseBodyParameter {
	s.Constraints = &v
	return s
}

func (s *GetParameterResponseBodyParameter) SetCreatedBy(v string) *GetParameterResponseBodyParameter {
	s.CreatedBy = &v
	return s
}

func (s *GetParameterResponseBodyParameter) SetCreatedDate(v string) *GetParameterResponseBodyParameter {
	s.CreatedDate = &v
	return s
}

func (s *GetParameterResponseBodyParameter) SetDescription(v string) *GetParameterResponseBodyParameter {
	s.Description = &v
	return s
}

func (s *GetParameterResponseBodyParameter) SetId(v string) *GetParameterResponseBodyParameter {
	s.Id = &v
	return s
}

func (s *GetParameterResponseBodyParameter) SetName(v string) *GetParameterResponseBodyParameter {
	s.Name = &v
	return s
}

func (s *GetParameterResponseBodyParameter) SetParameterVersion(v int32) *GetParameterResponseBodyParameter {
	s.ParameterVersion = &v
	return s
}

func (s *GetParameterResponseBodyParameter) SetResourceGroupId(v string) *GetParameterResponseBodyParameter {
	s.ResourceGroupId = &v
	return s
}

func (s *GetParameterResponseBodyParameter) SetShareType(v string) *GetParameterResponseBodyParameter {
	s.ShareType = &v
	return s
}

func (s *GetParameterResponseBodyParameter) SetTags(v map[string]interface{}) *GetParameterResponseBodyParameter {
	s.Tags = v
	return s
}

func (s *GetParameterResponseBodyParameter) SetType(v string) *GetParameterResponseBodyParameter {
	s.Type = &v
	return s
}

func (s *GetParameterResponseBodyParameter) SetUpdatedBy(v string) *GetParameterResponseBodyParameter {
	s.UpdatedBy = &v
	return s
}

func (s *GetParameterResponseBodyParameter) SetUpdatedDate(v string) *GetParameterResponseBodyParameter {
	s.UpdatedDate = &v
	return s
}

func (s *GetParameterResponseBodyParameter) SetValue(v string) *GetParameterResponseBodyParameter {
	s.Value = &v
	return s
}

type GetParameterResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetParameterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetParameterResponse) String() string {
	return tea.Prettify(s)
}

func (s GetParameterResponse) GoString() string {
	return s.String()
}

func (s *GetParameterResponse) SetHeaders(v map[string]*string) *GetParameterResponse {
	s.Headers = v
	return s
}

func (s *GetParameterResponse) SetStatusCode(v int32) *GetParameterResponse {
	s.StatusCode = &v
	return s
}

func (s *GetParameterResponse) SetBody(v *GetParameterResponseBody) *GetParameterResponse {
	s.Body = v
	return s
}

type GetParametersRequest struct {
	// The names of the common parameters.
	//
	// This parameter is required.
	//
	// example:
	//
	// ["parameter1","parameter2"]
	Names *string `json:"Names,omitempty" xml:"Names,omitempty"`
	// The ID of the region.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetParametersRequest) String() string {
	return tea.Prettify(s)
}

func (s GetParametersRequest) GoString() string {
	return s.String()
}

func (s *GetParametersRequest) SetNames(v string) *GetParametersRequest {
	s.Names = &v
	return s
}

func (s *GetParametersRequest) SetRegionId(v string) *GetParametersRequest {
	s.RegionId = &v
	return s
}

type GetParametersResponseBody struct {
	// Invalid parameters.
	InvalidParameters []*string `json:"InvalidParameters,omitempty" xml:"InvalidParameters,omitempty" type:"Repeated"`
	// The information about the common parameters.
	Parameters []*GetParametersResponseBodyParameters `json:"Parameters,omitempty" xml:"Parameters,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 2597E94B-5346-42D1-BB58-D3333EDD0975
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetParametersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetParametersResponseBody) GoString() string {
	return s.String()
}

func (s *GetParametersResponseBody) SetInvalidParameters(v []*string) *GetParametersResponseBody {
	s.InvalidParameters = v
	return s
}

func (s *GetParametersResponseBody) SetParameters(v []*GetParametersResponseBodyParameters) *GetParametersResponseBody {
	s.Parameters = v
	return s
}

func (s *GetParametersResponseBody) SetRequestId(v string) *GetParametersResponseBody {
	s.RequestId = &v
	return s
}

type GetParametersResponseBodyParameters struct {
	// The constraints of the common parameter.
	//
	// example:
	//
	// {\\"MaxLength\\": 2}
	Constraints *string `json:"Constraints,omitempty" xml:"Constraints,omitempty"`
	// The user who created the common parameter.
	//
	// example:
	//
	// root(130900000)
	CreatedBy *string `json:"CreatedBy,omitempty" xml:"CreatedBy,omitempty"`
	// The time when the common parameter was created.
	//
	// example:
	//
	// 2020-10-22T03:30:45Z
	CreatedDate *string `json:"CreatedDate,omitempty" xml:"CreatedDate,omitempty"`
	// The description of the common parameter.
	//
	// example:
	//
	// parameter
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the common parameter.
	//
	// example:
	//
	// p-7cdc0000000000000000
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the common parameter.
	//
	// example:
	//
	// MyParameter
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The version number of the common parameter.
	//
	// example:
	//
	// 1
	ParameterVersion *int32 `json:"ParameterVersion,omitempty" xml:"ParameterVersion,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfmxsn4m4******
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The share type of the common parameter.
	//
	// example:
	//
	// Private
	ShareType *string `json:"ShareType,omitempty" xml:"ShareType,omitempty"`
	// The tags.
	//
	// example:
	//
	// {"k1": "v1", "k2": "v2"}
	Tags map[string]interface{} `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The type of the parameter.
	//
	// example:
	//
	// StringList
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The user who updated the common parameter.
	//
	// example:
	//
	// root(130900000)
	UpdatedBy *string `json:"UpdatedBy,omitempty" xml:"UpdatedBy,omitempty"`
	// The time when the parameter was updated.
	//
	// example:
	//
	// 2020-10-22T03:30:45Z
	UpdatedDate *string `json:"UpdatedDate,omitempty" xml:"UpdatedDate,omitempty"`
	// The value of the common parameter.
	//
	// example:
	//
	// parameter,parameter1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetParametersResponseBodyParameters) String() string {
	return tea.Prettify(s)
}

func (s GetParametersResponseBodyParameters) GoString() string {
	return s.String()
}

func (s *GetParametersResponseBodyParameters) SetConstraints(v string) *GetParametersResponseBodyParameters {
	s.Constraints = &v
	return s
}

func (s *GetParametersResponseBodyParameters) SetCreatedBy(v string) *GetParametersResponseBodyParameters {
	s.CreatedBy = &v
	return s
}

func (s *GetParametersResponseBodyParameters) SetCreatedDate(v string) *GetParametersResponseBodyParameters {
	s.CreatedDate = &v
	return s
}

func (s *GetParametersResponseBodyParameters) SetDescription(v string) *GetParametersResponseBodyParameters {
	s.Description = &v
	return s
}

func (s *GetParametersResponseBodyParameters) SetId(v string) *GetParametersResponseBodyParameters {
	s.Id = &v
	return s
}

func (s *GetParametersResponseBodyParameters) SetName(v string) *GetParametersResponseBodyParameters {
	s.Name = &v
	return s
}

func (s *GetParametersResponseBodyParameters) SetParameterVersion(v int32) *GetParametersResponseBodyParameters {
	s.ParameterVersion = &v
	return s
}

func (s *GetParametersResponseBodyParameters) SetResourceGroupId(v string) *GetParametersResponseBodyParameters {
	s.ResourceGroupId = &v
	return s
}

func (s *GetParametersResponseBodyParameters) SetShareType(v string) *GetParametersResponseBodyParameters {
	s.ShareType = &v
	return s
}

func (s *GetParametersResponseBodyParameters) SetTags(v map[string]interface{}) *GetParametersResponseBodyParameters {
	s.Tags = v
	return s
}

func (s *GetParametersResponseBodyParameters) SetType(v string) *GetParametersResponseBodyParameters {
	s.Type = &v
	return s
}

func (s *GetParametersResponseBodyParameters) SetUpdatedBy(v string) *GetParametersResponseBodyParameters {
	s.UpdatedBy = &v
	return s
}

func (s *GetParametersResponseBodyParameters) SetUpdatedDate(v string) *GetParametersResponseBodyParameters {
	s.UpdatedDate = &v
	return s
}

func (s *GetParametersResponseBodyParameters) SetValue(v string) *GetParametersResponseBodyParameters {
	s.Value = &v
	return s
}

type GetParametersResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetParametersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetParametersResponse) String() string {
	return tea.Prettify(s)
}

func (s GetParametersResponse) GoString() string {
	return s.String()
}

func (s *GetParametersResponse) SetHeaders(v map[string]*string) *GetParametersResponse {
	s.Headers = v
	return s
}

func (s *GetParametersResponse) SetStatusCode(v int32) *GetParametersResponse {
	s.StatusCode = &v
	return s
}

func (s *GetParametersResponse) SetBody(v *GetParametersResponseBody) *GetParametersResponse {
	s.Body = v
	return s
}

type GetParametersByPathRequest struct {
	// The number of entries per page.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// A pagination token. It can be used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// MTRBMDc0NjAtRUJFNy00N0NBLTk3NTctMTJDQzA
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The path of the parameter. For example, if the name of a parameter is /path/path1/Myparameter, the path of the parameter is /path/path1/.
	//
	// This parameter is required.
	//
	// example:
	//
	// /parameter
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// Specifies whether to recursively query encryption parameters from all levels of directories in the specified path. Valid values: true and false. For example, if you want to query the /secretParameter/mySecretParameter and /secretParameter/secretParameter 1/mySecretParameter parameters, the valid values specify the parameters to be returned.
	//
	// 	- true: returns both of the /secretParameter/mySecretParameter and /secretParameter/secretParameter1/mySecretParameter parameters.
	//
	// 	- false: returns only the /secretParameter/mySecretParameter parameter.
	//
	// example:
	//
	// false
	Recursive *bool `json:"Recursive,omitempty" xml:"Recursive,omitempty"`
	// The ID of the region.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetParametersByPathRequest) String() string {
	return tea.Prettify(s)
}

func (s GetParametersByPathRequest) GoString() string {
	return s.String()
}

func (s *GetParametersByPathRequest) SetMaxResults(v int32) *GetParametersByPathRequest {
	s.MaxResults = &v
	return s
}

func (s *GetParametersByPathRequest) SetNextToken(v string) *GetParametersByPathRequest {
	s.NextToken = &v
	return s
}

func (s *GetParametersByPathRequest) SetPath(v string) *GetParametersByPathRequest {
	s.Path = &v
	return s
}

func (s *GetParametersByPathRequest) SetRecursive(v bool) *GetParametersByPathRequest {
	s.Recursive = &v
	return s
}

func (s *GetParametersByPathRequest) SetRegionId(v string) *GetParametersByPathRequest {
	s.RegionId = &v
	return s
}

type GetParametersByPathResponseBody struct {
	// The number of entries per page.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// A pagination token. It can be used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// gAAAAABfTgv5ewUWmNdJ3g7JVLvX70sPH90GZOVGC
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The information about the common parameters.
	Parameters []*GetParametersByPathResponseBodyParameters `json:"Parameters,omitempty" xml:"Parameters,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 25156E99-7437-4590-AA58-2ACA17DE405C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of returned entries.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s GetParametersByPathResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetParametersByPathResponseBody) GoString() string {
	return s.String()
}

func (s *GetParametersByPathResponseBody) SetMaxResults(v int32) *GetParametersByPathResponseBody {
	s.MaxResults = &v
	return s
}

func (s *GetParametersByPathResponseBody) SetNextToken(v string) *GetParametersByPathResponseBody {
	s.NextToken = &v
	return s
}

func (s *GetParametersByPathResponseBody) SetParameters(v []*GetParametersByPathResponseBodyParameters) *GetParametersByPathResponseBody {
	s.Parameters = v
	return s
}

func (s *GetParametersByPathResponseBody) SetRequestId(v string) *GetParametersByPathResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetParametersByPathResponseBody) SetTotalCount(v int32) *GetParametersByPathResponseBody {
	s.TotalCount = &v
	return s
}

type GetParametersByPathResponseBodyParameters struct {
	// The constraints of the common parameter.
	//
	// example:
	//
	// {\\"MaxLength\\": 2}
	Constraints *string `json:"Constraints,omitempty" xml:"Constraints,omitempty"`
	// The user who created the common parameter.
	//
	// example:
	//
	// root(130900000)
	CreatedBy *string `json:"CreatedBy,omitempty" xml:"CreatedBy,omitempty"`
	// The time when the common parameter was created.
	//
	// example:
	//
	// 2020-10-21T04:03:12Z
	CreatedDate *string `json:"CreatedDate,omitempty" xml:"CreatedDate,omitempty"`
	// The description of the common parameter.
	//
	// example:
	//
	// parameter
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the common parameter.
	//
	// example:
	//
	// p-7cdc0000000000000000
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the common parameter.
	//
	// example:
	//
	// myParameter
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The version number of the common parameter.
	//
	// example:
	//
	// 1
	ParameterVersion *int32 `json:"ParameterVersion,omitempty" xml:"ParameterVersion,omitempty"`
	// The share type of the common parameter.
	//
	// example:
	//
	// Private
	ShareType *string `json:"ShareType,omitempty" xml:"ShareType,omitempty"`
	// The tags added to the common parameters.
	//
	// example:
	//
	// {"k1": "v1", "k2": "v2"}
	Tags map[string]interface{} `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The type of the common parameter.
	//
	// example:
	//
	// StringList
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The user who updated the common parameter.
	//
	// example:
	//
	// root(130900000)
	UpdatedBy *string `json:"UpdatedBy,omitempty" xml:"UpdatedBy,omitempty"`
	// The time when the common parameter was last updated.
	//
	// example:
	//
	// 2020-10-21T04:03:12Z
	UpdatedDate *string `json:"UpdatedDate,omitempty" xml:"UpdatedDate,omitempty"`
	// The value of the common parameter.
	//
	// example:
	//
	// "parameter1,parameter2"
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetParametersByPathResponseBodyParameters) String() string {
	return tea.Prettify(s)
}

func (s GetParametersByPathResponseBodyParameters) GoString() string {
	return s.String()
}

func (s *GetParametersByPathResponseBodyParameters) SetConstraints(v string) *GetParametersByPathResponseBodyParameters {
	s.Constraints = &v
	return s
}

func (s *GetParametersByPathResponseBodyParameters) SetCreatedBy(v string) *GetParametersByPathResponseBodyParameters {
	s.CreatedBy = &v
	return s
}

func (s *GetParametersByPathResponseBodyParameters) SetCreatedDate(v string) *GetParametersByPathResponseBodyParameters {
	s.CreatedDate = &v
	return s
}

func (s *GetParametersByPathResponseBodyParameters) SetDescription(v string) *GetParametersByPathResponseBodyParameters {
	s.Description = &v
	return s
}

func (s *GetParametersByPathResponseBodyParameters) SetId(v string) *GetParametersByPathResponseBodyParameters {
	s.Id = &v
	return s
}

func (s *GetParametersByPathResponseBodyParameters) SetName(v string) *GetParametersByPathResponseBodyParameters {
	s.Name = &v
	return s
}

func (s *GetParametersByPathResponseBodyParameters) SetParameterVersion(v int32) *GetParametersByPathResponseBodyParameters {
	s.ParameterVersion = &v
	return s
}

func (s *GetParametersByPathResponseBodyParameters) SetShareType(v string) *GetParametersByPathResponseBodyParameters {
	s.ShareType = &v
	return s
}

func (s *GetParametersByPathResponseBodyParameters) SetTags(v map[string]interface{}) *GetParametersByPathResponseBodyParameters {
	s.Tags = v
	return s
}

func (s *GetParametersByPathResponseBodyParameters) SetType(v string) *GetParametersByPathResponseBodyParameters {
	s.Type = &v
	return s
}

func (s *GetParametersByPathResponseBodyParameters) SetUpdatedBy(v string) *GetParametersByPathResponseBodyParameters {
	s.UpdatedBy = &v
	return s
}

func (s *GetParametersByPathResponseBodyParameters) SetUpdatedDate(v string) *GetParametersByPathResponseBodyParameters {
	s.UpdatedDate = &v
	return s
}

func (s *GetParametersByPathResponseBodyParameters) SetValue(v string) *GetParametersByPathResponseBodyParameters {
	s.Value = &v
	return s
}

type GetParametersByPathResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetParametersByPathResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetParametersByPathResponse) String() string {
	return tea.Prettify(s)
}

func (s GetParametersByPathResponse) GoString() string {
	return s.String()
}

func (s *GetParametersByPathResponse) SetHeaders(v map[string]*string) *GetParametersByPathResponse {
	s.Headers = v
	return s
}

func (s *GetParametersByPathResponse) SetStatusCode(v int32) *GetParametersByPathResponse {
	s.StatusCode = &v
	return s
}

func (s *GetParametersByPathResponse) SetBody(v *GetParametersByPathResponseBody) *GetParametersByPathResponse {
	s.Body = v
	return s
}

type GetPatchBaselineRequest struct {
	// The name of the patch baseline.
	//
	// This parameter is required.
	//
	// example:
	//
	// MyPatchBaseline
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the region in which the patch baseline whose details you want to query resides.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetPatchBaselineRequest) String() string {
	return tea.Prettify(s)
}

func (s GetPatchBaselineRequest) GoString() string {
	return s.String()
}

func (s *GetPatchBaselineRequest) SetName(v string) *GetPatchBaselineRequest {
	s.Name = &v
	return s
}

func (s *GetPatchBaselineRequest) SetRegionId(v string) *GetPatchBaselineRequest {
	s.RegionId = &v
	return s
}

type GetPatchBaselineResponseBody struct {
	// The details of the patch baseline.
	PatchBaseline *GetPatchBaselineResponseBodyPatchBaseline `json:"PatchBaseline,omitempty" xml:"PatchBaseline,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 2C630E64-7273-57AC-A598-1B2B8B35CEA5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetPatchBaselineResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetPatchBaselineResponseBody) GoString() string {
	return s.String()
}

func (s *GetPatchBaselineResponseBody) SetPatchBaseline(v *GetPatchBaselineResponseBodyPatchBaseline) *GetPatchBaselineResponseBody {
	s.PatchBaseline = v
	return s
}

func (s *GetPatchBaselineResponseBody) SetRequestId(v string) *GetPatchBaselineResponseBody {
	s.RequestId = &v
	return s
}

type GetPatchBaselineResponseBodyPatchBaseline struct {
	// The rules of scanning and installing patches for the specified operating system.
	//
	// example:
	//
	// {"PatchRules":[{"PatchFilterGroup":[{"Key":"PatchSet","Values":["OS"]},{"Key":"ProductFamily","Values":["Windows"]},{"Key":"Product","Values":["Windows 10","Windows 7"]},{"Key":"Classification","Values":["Security Updates","Updates","Update Rollups","Critical Updates"]},{"Key":"Severity","Values":["Critical","Important","Moderate"]}],"ApproveAfterDays":7,"ApproveUntilDate":"","EnableNonSecurity":true,"ComplianceLevel":"Medium"}]}
	ApprovalRules *string `json:"ApprovalRules,omitempty" xml:"ApprovalRules,omitempty"`
	// The approved patches.
	ApprovedPatches []*string `json:"ApprovedPatches,omitempty" xml:"ApprovedPatches,omitempty" type:"Repeated"`
	// Indicates whether the approved patch involves updates other than security-related updates.
	//
	// example:
	//
	// true
	ApprovedPatchesEnableNonSecurity *bool `json:"ApprovedPatchesEnableNonSecurity,omitempty" xml:"ApprovedPatchesEnableNonSecurity,omitempty"`
	// The creator of the patch baseline.
	//
	// example:
	//
	// root(130900000)
	CreatedBy *string `json:"CreatedBy,omitempty" xml:"CreatedBy,omitempty"`
	// The time when the patch baseline was created.
	//
	// example:
	//
	// 2021-09-07T03:42:56Z
	CreatedDate *string `json:"CreatedDate,omitempty" xml:"CreatedDate,omitempty"`
	// The description of the patch baseline.
	//
	// example:
	//
	// UpdatePatchBaseline
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the patch baseline.
	//
	// example:
	//
	// pb-445340b5c6504a85a300
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// Indicates whether the patch baseline is set as the default patch baseline.
	//
	// example:
	//
	// false
	IsDefault *bool `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	// The name of the patch baseline.
	//
	// example:
	//
	// MypatchBaseline
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The type of the operating system.
	//
	// example:
	//
	// Windows
	OperationSystem *string `json:"OperationSystem,omitempty" xml:"OperationSystem,omitempty"`
	// The rejected patches.
	RejectedPatches []*string `json:"RejectedPatches,omitempty" xml:"RejectedPatches,omitempty" type:"Repeated"`
	// The action of the rejected patch.
	//
	// example:
	//
	// ALLOW_AS_DEPENDENCY
	RejectedPatchesAction *string `json:"RejectedPatchesAction,omitempty" xml:"RejectedPatchesAction,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfmzmhzoaad5oq
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The share type of the patch baseline.
	//
	// example:
	//
	// Private
	ShareType *string `json:"ShareType,omitempty" xml:"ShareType,omitempty"`
	// The patch source configurations.
	Sources []*string `json:"Sources,omitempty" xml:"Sources,omitempty" type:"Repeated"`
	// The tags.
	Tags []*GetPatchBaselineResponseBodyPatchBaselineTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The user who last modified the patch baseline.
	//
	// example:
	//
	// root(130900000)
	UpdatedBy *string `json:"UpdatedBy,omitempty" xml:"UpdatedBy,omitempty"`
	// The time when the patch baseline was last modified.
	//
	// example:
	//
	// 2021-09-08T07:26:38Z
	UpdatedDate *string `json:"UpdatedDate,omitempty" xml:"UpdatedDate,omitempty"`
}

func (s GetPatchBaselineResponseBodyPatchBaseline) String() string {
	return tea.Prettify(s)
}

func (s GetPatchBaselineResponseBodyPatchBaseline) GoString() string {
	return s.String()
}

func (s *GetPatchBaselineResponseBodyPatchBaseline) SetApprovalRules(v string) *GetPatchBaselineResponseBodyPatchBaseline {
	s.ApprovalRules = &v
	return s
}

func (s *GetPatchBaselineResponseBodyPatchBaseline) SetApprovedPatches(v []*string) *GetPatchBaselineResponseBodyPatchBaseline {
	s.ApprovedPatches = v
	return s
}

func (s *GetPatchBaselineResponseBodyPatchBaseline) SetApprovedPatchesEnableNonSecurity(v bool) *GetPatchBaselineResponseBodyPatchBaseline {
	s.ApprovedPatchesEnableNonSecurity = &v
	return s
}

func (s *GetPatchBaselineResponseBodyPatchBaseline) SetCreatedBy(v string) *GetPatchBaselineResponseBodyPatchBaseline {
	s.CreatedBy = &v
	return s
}

func (s *GetPatchBaselineResponseBodyPatchBaseline) SetCreatedDate(v string) *GetPatchBaselineResponseBodyPatchBaseline {
	s.CreatedDate = &v
	return s
}

func (s *GetPatchBaselineResponseBodyPatchBaseline) SetDescription(v string) *GetPatchBaselineResponseBodyPatchBaseline {
	s.Description = &v
	return s
}

func (s *GetPatchBaselineResponseBodyPatchBaseline) SetId(v string) *GetPatchBaselineResponseBodyPatchBaseline {
	s.Id = &v
	return s
}

func (s *GetPatchBaselineResponseBodyPatchBaseline) SetIsDefault(v bool) *GetPatchBaselineResponseBodyPatchBaseline {
	s.IsDefault = &v
	return s
}

func (s *GetPatchBaselineResponseBodyPatchBaseline) SetName(v string) *GetPatchBaselineResponseBodyPatchBaseline {
	s.Name = &v
	return s
}

func (s *GetPatchBaselineResponseBodyPatchBaseline) SetOperationSystem(v string) *GetPatchBaselineResponseBodyPatchBaseline {
	s.OperationSystem = &v
	return s
}

func (s *GetPatchBaselineResponseBodyPatchBaseline) SetRejectedPatches(v []*string) *GetPatchBaselineResponseBodyPatchBaseline {
	s.RejectedPatches = v
	return s
}

func (s *GetPatchBaselineResponseBodyPatchBaseline) SetRejectedPatchesAction(v string) *GetPatchBaselineResponseBodyPatchBaseline {
	s.RejectedPatchesAction = &v
	return s
}

func (s *GetPatchBaselineResponseBodyPatchBaseline) SetResourceGroupId(v string) *GetPatchBaselineResponseBodyPatchBaseline {
	s.ResourceGroupId = &v
	return s
}

func (s *GetPatchBaselineResponseBodyPatchBaseline) SetShareType(v string) *GetPatchBaselineResponseBodyPatchBaseline {
	s.ShareType = &v
	return s
}

func (s *GetPatchBaselineResponseBodyPatchBaseline) SetSources(v []*string) *GetPatchBaselineResponseBodyPatchBaseline {
	s.Sources = v
	return s
}

func (s *GetPatchBaselineResponseBodyPatchBaseline) SetTags(v []*GetPatchBaselineResponseBodyPatchBaselineTags) *GetPatchBaselineResponseBodyPatchBaseline {
	s.Tags = v
	return s
}

func (s *GetPatchBaselineResponseBodyPatchBaseline) SetUpdatedBy(v string) *GetPatchBaselineResponseBodyPatchBaseline {
	s.UpdatedBy = &v
	return s
}

func (s *GetPatchBaselineResponseBodyPatchBaseline) SetUpdatedDate(v string) *GetPatchBaselineResponseBodyPatchBaseline {
	s.UpdatedDate = &v
	return s
}

type GetPatchBaselineResponseBodyPatchBaselineTags struct {
	// The key of the tag.
	//
	// example:
	//
	// key
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The value of the tag.
	//
	// example:
	//
	// value
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s GetPatchBaselineResponseBodyPatchBaselineTags) String() string {
	return tea.Prettify(s)
}

func (s GetPatchBaselineResponseBodyPatchBaselineTags) GoString() string {
	return s.String()
}

func (s *GetPatchBaselineResponseBodyPatchBaselineTags) SetTagKey(v string) *GetPatchBaselineResponseBodyPatchBaselineTags {
	s.TagKey = &v
	return s
}

func (s *GetPatchBaselineResponseBodyPatchBaselineTags) SetTagValue(v string) *GetPatchBaselineResponseBodyPatchBaselineTags {
	s.TagValue = &v
	return s
}

type GetPatchBaselineResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPatchBaselineResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPatchBaselineResponse) String() string {
	return tea.Prettify(s)
}

func (s GetPatchBaselineResponse) GoString() string {
	return s.String()
}

func (s *GetPatchBaselineResponse) SetHeaders(v map[string]*string) *GetPatchBaselineResponse {
	s.Headers = v
	return s
}

func (s *GetPatchBaselineResponse) SetStatusCode(v int32) *GetPatchBaselineResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPatchBaselineResponse) SetBody(v *GetPatchBaselineResponseBody) *GetPatchBaselineResponse {
	s.Body = v
	return s
}

type GetSecretParameterRequest struct {
	// The name of the parameter. The name must be 1 to 180 characters in length, and can contain letters, digits, hyphens (-), and underscores (_). It cannot start with ALIYUN, ACS, ALIBABA, ALICLOUD, or OOS.
	//
	// This parameter is required.
	//
	// example:
	//
	// MySecretParameter
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The version number of the common parameter. Valid values: 1 to 100.
	//
	// example:
	//
	// 1
	ParameterVersion *int32 `json:"ParameterVersion,omitempty" xml:"ParameterVersion,omitempty"`
	// The ID of the region.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Specifies whether to decrypt the parameter value. The decrypted parameter value is returned only if this parameter is set to true. Otherwise, null is returned.
	//
	// example:
	//
	// false
	WithDecryption *bool `json:"WithDecryption,omitempty" xml:"WithDecryption,omitempty"`
}

func (s GetSecretParameterRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSecretParameterRequest) GoString() string {
	return s.String()
}

func (s *GetSecretParameterRequest) SetName(v string) *GetSecretParameterRequest {
	s.Name = &v
	return s
}

func (s *GetSecretParameterRequest) SetParameterVersion(v int32) *GetSecretParameterRequest {
	s.ParameterVersion = &v
	return s
}

func (s *GetSecretParameterRequest) SetRegionId(v string) *GetSecretParameterRequest {
	s.RegionId = &v
	return s
}

func (s *GetSecretParameterRequest) SetWithDecryption(v bool) *GetSecretParameterRequest {
	s.WithDecryption = &v
	return s
}

type GetSecretParameterResponseBody struct {
	// The information about the encryption parameter.
	Parameter *GetSecretParameterResponseBodyParameter `json:"Parameter,omitempty" xml:"Parameter,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 7F14FB7C-C9BE-44AE-92ED-21ACC02FBFD2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetSecretParameterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSecretParameterResponseBody) GoString() string {
	return s.String()
}

func (s *GetSecretParameterResponseBody) SetParameter(v *GetSecretParameterResponseBodyParameter) *GetSecretParameterResponseBody {
	s.Parameter = v
	return s
}

func (s *GetSecretParameterResponseBody) SetRequestId(v string) *GetSecretParameterResponseBody {
	s.RequestId = &v
	return s
}

type GetSecretParameterResponseBodyParameter struct {
	// The constraints of the encryption parameter.
	//
	// example:
	//
	// \\"{\\"\\"AllowedValues":["secretparameter"],"AllowedPattern":".*","MinLength":0,"MaxLength":20}\\"
	Constraints *string `json:"Constraints,omitempty" xml:"Constraints,omitempty"`
	// The user who created the encryption parameter.
	//
	// example:
	//
	// root(130900000)
	CreatedBy *string `json:"CreatedBy,omitempty" xml:"CreatedBy,omitempty"`
	// The time when the encryption parameter was created.
	//
	// example:
	//
	// 2020-09-01T09:28:47Z
	CreatedDate *string `json:"CreatedDate,omitempty" xml:"CreatedDate,omitempty"`
	// The instance ID of the KMS instance.
	//
	// example:
	//
	// kst-hzz****
	DKMSInstanceId *string `json:"DKMSInstanceId,omitempty" xml:"DKMSInstanceId,omitempty"`
	// The description of the encryption parameter.
	//
	// example:
	//
	// SecretParameter
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the encryption parameter.
	//
	// example:
	//
	// p-14ed150fdcd048xxxxxx
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The ID of the key of Key Management Service (KMS) that is used for encryption.
	//
	// example:
	//
	// 80e9409f-78fa-42ab-84bd-83f40c******
	KeyId *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	// The name of the encryption parameter.
	//
	// example:
	//
	// MySecretParameter
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The version number of the encryption parameter.
	//
	// example:
	//
	// 1
	ParameterVersion *int32 `json:"ParameterVersion,omitempty" xml:"ParameterVersion,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfmxsn4m4******
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The share type of the encryption parameter.
	//
	// example:
	//
	// Private
	ShareType *string `json:"ShareType,omitempty" xml:"ShareType,omitempty"`
	// The tags of the parameter.
	//
	// example:
	//
	// {"k1": "v1", "k2": "v2"}
	Tags map[string]interface{} `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The type of the parameter.
	//
	// example:
	//
	// Secret
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The user who updated the encryption parameter.
	//
	// example:
	//
	// root(130900000)
	UpdatedBy *string `json:"UpdatedBy,omitempty" xml:"UpdatedBy,omitempty"`
	// The time when the encryption parameter was updated.
	//
	// example:
	//
	// 2020-09-01T09:35:17Z
	UpdatedDate *string `json:"UpdatedDate,omitempty" xml:"UpdatedDate,omitempty"`
	// The value of the encryption parameter.
	//
	// example:
	//
	// SecretParameter
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetSecretParameterResponseBodyParameter) String() string {
	return tea.Prettify(s)
}

func (s GetSecretParameterResponseBodyParameter) GoString() string {
	return s.String()
}

func (s *GetSecretParameterResponseBodyParameter) SetConstraints(v string) *GetSecretParameterResponseBodyParameter {
	s.Constraints = &v
	return s
}

func (s *GetSecretParameterResponseBodyParameter) SetCreatedBy(v string) *GetSecretParameterResponseBodyParameter {
	s.CreatedBy = &v
	return s
}

func (s *GetSecretParameterResponseBodyParameter) SetCreatedDate(v string) *GetSecretParameterResponseBodyParameter {
	s.CreatedDate = &v
	return s
}

func (s *GetSecretParameterResponseBodyParameter) SetDKMSInstanceId(v string) *GetSecretParameterResponseBodyParameter {
	s.DKMSInstanceId = &v
	return s
}

func (s *GetSecretParameterResponseBodyParameter) SetDescription(v string) *GetSecretParameterResponseBodyParameter {
	s.Description = &v
	return s
}

func (s *GetSecretParameterResponseBodyParameter) SetId(v string) *GetSecretParameterResponseBodyParameter {
	s.Id = &v
	return s
}

func (s *GetSecretParameterResponseBodyParameter) SetKeyId(v string) *GetSecretParameterResponseBodyParameter {
	s.KeyId = &v
	return s
}

func (s *GetSecretParameterResponseBodyParameter) SetName(v string) *GetSecretParameterResponseBodyParameter {
	s.Name = &v
	return s
}

func (s *GetSecretParameterResponseBodyParameter) SetParameterVersion(v int32) *GetSecretParameterResponseBodyParameter {
	s.ParameterVersion = &v
	return s
}

func (s *GetSecretParameterResponseBodyParameter) SetResourceGroupId(v string) *GetSecretParameterResponseBodyParameter {
	s.ResourceGroupId = &v
	return s
}

func (s *GetSecretParameterResponseBodyParameter) SetShareType(v string) *GetSecretParameterResponseBodyParameter {
	s.ShareType = &v
	return s
}

func (s *GetSecretParameterResponseBodyParameter) SetTags(v map[string]interface{}) *GetSecretParameterResponseBodyParameter {
	s.Tags = v
	return s
}

func (s *GetSecretParameterResponseBodyParameter) SetType(v string) *GetSecretParameterResponseBodyParameter {
	s.Type = &v
	return s
}

func (s *GetSecretParameterResponseBodyParameter) SetUpdatedBy(v string) *GetSecretParameterResponseBodyParameter {
	s.UpdatedBy = &v
	return s
}

func (s *GetSecretParameterResponseBodyParameter) SetUpdatedDate(v string) *GetSecretParameterResponseBodyParameter {
	s.UpdatedDate = &v
	return s
}

func (s *GetSecretParameterResponseBodyParameter) SetValue(v string) *GetSecretParameterResponseBodyParameter {
	s.Value = &v
	return s
}

type GetSecretParameterResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSecretParameterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSecretParameterResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSecretParameterResponse) GoString() string {
	return s.String()
}

func (s *GetSecretParameterResponse) SetHeaders(v map[string]*string) *GetSecretParameterResponse {
	s.Headers = v
	return s
}

func (s *GetSecretParameterResponse) SetStatusCode(v int32) *GetSecretParameterResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSecretParameterResponse) SetBody(v *GetSecretParameterResponseBody) *GetSecretParameterResponse {
	s.Body = v
	return s
}

type GetSecretParametersRequest struct {
	// The name of the encryption parameter. Multiple encryption parameters can form a JSON array in the format of ["xxxxxxxxx", "yyyyyyyyy", … "zzzzzzzzz"]. Each JSON array can contain a maximum of 10 encryption parameters. Multiple encryption parameters in the array are separated by commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// ["secretParameter","secretParameter1"]
	Names *string `json:"Names,omitempty" xml:"Names,omitempty"`
	// The ID of the region.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Specifies whether to decrypt the parameter value. Default value: false. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// false
	WithDecryption *bool `json:"WithDecryption,omitempty" xml:"WithDecryption,omitempty"`
}

func (s GetSecretParametersRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSecretParametersRequest) GoString() string {
	return s.String()
}

func (s *GetSecretParametersRequest) SetNames(v string) *GetSecretParametersRequest {
	s.Names = &v
	return s
}

func (s *GetSecretParametersRequest) SetRegionId(v string) *GetSecretParametersRequest {
	s.RegionId = &v
	return s
}

func (s *GetSecretParametersRequest) SetWithDecryption(v bool) *GetSecretParametersRequest {
	s.WithDecryption = &v
	return s
}

type GetSecretParametersResponseBody struct {
	// Invalid encryption parameter.
	InvalidParameters []*string `json:"InvalidParameters,omitempty" xml:"InvalidParameters,omitempty" type:"Repeated"`
	// The information about the encryption parameter.
	Parameters []*GetSecretParametersResponseBodyParameters `json:"Parameters,omitempty" xml:"Parameters,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// A5320F1D-92D9-44BB-A416-5FC525ED6D57
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetSecretParametersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSecretParametersResponseBody) GoString() string {
	return s.String()
}

func (s *GetSecretParametersResponseBody) SetInvalidParameters(v []*string) *GetSecretParametersResponseBody {
	s.InvalidParameters = v
	return s
}

func (s *GetSecretParametersResponseBody) SetParameters(v []*GetSecretParametersResponseBodyParameters) *GetSecretParametersResponseBody {
	s.Parameters = v
	return s
}

func (s *GetSecretParametersResponseBody) SetRequestId(v string) *GetSecretParametersResponseBody {
	s.RequestId = &v
	return s
}

type GetSecretParametersResponseBodyParameters struct {
	// The constraints of the encryption parameter.
	//
	// example:
	//
	// {\\"AllowedValues\\": [\\"test\\"]}
	Constraints *string `json:"Constraints,omitempty" xml:"Constraints,omitempty"`
	// The user who created the encryption parameter.
	//
	// example:
	//
	// root(130900000)
	CreatedBy *string `json:"CreatedBy,omitempty" xml:"CreatedBy,omitempty"`
	// The time when the encryption parameter was created.
	//
	// example:
	//
	// 2020-10-22T03:11:13Z
	CreatedDate *string `json:"CreatedDate,omitempty" xml:"CreatedDate,omitempty"`
	// The description of the encryption parameter.
	//
	// example:
	//
	// secretParameter
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the encryption parameter.
	//
	// example:
	//
	// p-7cdc0000000000000000
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The ID of the key.
	//
	// example:
	//
	// ssh-bp67acfmxazb4p****
	KeyId *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	// The name of the encryption parameter.
	//
	// example:
	//
	// MySecretParameter
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The version number of the encryption parameter.
	//
	// example:
	//
	// 1
	ParameterVersion *int32 `json:"ParameterVersion,omitempty" xml:"ParameterVersion,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfmxsn4m4******
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The share type of the encryption parameter.
	//
	// example:
	//
	// Private
	ShareType *string `json:"ShareType,omitempty" xml:"ShareType,omitempty"`
	// The tags.
	//
	// example:
	//
	// {"k1": "v1", "k2": "v2"}
	Tags map[string]interface{} `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The data type of the encryption parameter.
	//
	// example:
	//
	// Secret
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The user who updated the encryption parameter.
	//
	// example:
	//
	// root(130900000)
	UpdatedBy *string `json:"UpdatedBy,omitempty" xml:"UpdatedBy,omitempty"`
	// The time when the encryption parameter was updated.
	//
	// example:
	//
	// 2020-10-22T03:11:13Z
	UpdatedDate *string `json:"UpdatedDate,omitempty" xml:"UpdatedDate,omitempty"`
	// The value of the encryption parameter.
	//
	// example:
	//
	// secretParameter,secretParameter1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetSecretParametersResponseBodyParameters) String() string {
	return tea.Prettify(s)
}

func (s GetSecretParametersResponseBodyParameters) GoString() string {
	return s.String()
}

func (s *GetSecretParametersResponseBodyParameters) SetConstraints(v string) *GetSecretParametersResponseBodyParameters {
	s.Constraints = &v
	return s
}

func (s *GetSecretParametersResponseBodyParameters) SetCreatedBy(v string) *GetSecretParametersResponseBodyParameters {
	s.CreatedBy = &v
	return s
}

func (s *GetSecretParametersResponseBodyParameters) SetCreatedDate(v string) *GetSecretParametersResponseBodyParameters {
	s.CreatedDate = &v
	return s
}

func (s *GetSecretParametersResponseBodyParameters) SetDescription(v string) *GetSecretParametersResponseBodyParameters {
	s.Description = &v
	return s
}

func (s *GetSecretParametersResponseBodyParameters) SetId(v string) *GetSecretParametersResponseBodyParameters {
	s.Id = &v
	return s
}

func (s *GetSecretParametersResponseBodyParameters) SetKeyId(v string) *GetSecretParametersResponseBodyParameters {
	s.KeyId = &v
	return s
}

func (s *GetSecretParametersResponseBodyParameters) SetName(v string) *GetSecretParametersResponseBodyParameters {
	s.Name = &v
	return s
}

func (s *GetSecretParametersResponseBodyParameters) SetParameterVersion(v int32) *GetSecretParametersResponseBodyParameters {
	s.ParameterVersion = &v
	return s
}

func (s *GetSecretParametersResponseBodyParameters) SetResourceGroupId(v string) *GetSecretParametersResponseBodyParameters {
	s.ResourceGroupId = &v
	return s
}

func (s *GetSecretParametersResponseBodyParameters) SetShareType(v string) *GetSecretParametersResponseBodyParameters {
	s.ShareType = &v
	return s
}

func (s *GetSecretParametersResponseBodyParameters) SetTags(v map[string]interface{}) *GetSecretParametersResponseBodyParameters {
	s.Tags = v
	return s
}

func (s *GetSecretParametersResponseBodyParameters) SetType(v string) *GetSecretParametersResponseBodyParameters {
	s.Type = &v
	return s
}

func (s *GetSecretParametersResponseBodyParameters) SetUpdatedBy(v string) *GetSecretParametersResponseBodyParameters {
	s.UpdatedBy = &v
	return s
}

func (s *GetSecretParametersResponseBodyParameters) SetUpdatedDate(v string) *GetSecretParametersResponseBodyParameters {
	s.UpdatedDate = &v
	return s
}

func (s *GetSecretParametersResponseBodyParameters) SetValue(v string) *GetSecretParametersResponseBodyParameters {
	s.Value = &v
	return s
}

type GetSecretParametersResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSecretParametersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSecretParametersResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSecretParametersResponse) GoString() string {
	return s.String()
}

func (s *GetSecretParametersResponse) SetHeaders(v map[string]*string) *GetSecretParametersResponse {
	s.Headers = v
	return s
}

func (s *GetSecretParametersResponse) SetStatusCode(v int32) *GetSecretParametersResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSecretParametersResponse) SetBody(v *GetSecretParametersResponseBody) *GetSecretParametersResponse {
	s.Body = v
	return s
}

type GetSecretParametersByPathRequest struct {
	// The number of entries per page. Valid values: 1 to 10. Default value: 10.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// A pagination token. It can be used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// MTRBMDc0NjAtRUJFNy00N0NBLTk3NTctMTJDQzA
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The path of the encryption parameter. The path must be 1 to 200 characters in length. For example, if the name of an encryption parameter is /secretParameter/mySecretParameter, the path of the encryption parameter is /secretParameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// /secretParameter
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// Specifies whether to recursively query encryption parameters from all levels of directories in the specified path. Valid values: true and false. For example, if you want to query the /secretParameter/mySecretParameter and /secretParameter/secretParameter 1/mySecretParameter parameters, the valid values specify the parameters to be returned.
	//
	// 	- true: returns both of the /secretParameter/mySecretParameter and /secretParameter/secretParameter1/mySecretParameter parameters.
	//
	// 	- false: returns only the /secretParameter/mySecretParameter parameter.
	//
	// example:
	//
	// false
	Recursive *bool `json:"Recursive,omitempty" xml:"Recursive,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Specifies whether to decrypt the parameter value. Default value: false. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	WithDecryption *bool `json:"WithDecryption,omitempty" xml:"WithDecryption,omitempty"`
}

func (s GetSecretParametersByPathRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSecretParametersByPathRequest) GoString() string {
	return s.String()
}

func (s *GetSecretParametersByPathRequest) SetMaxResults(v int32) *GetSecretParametersByPathRequest {
	s.MaxResults = &v
	return s
}

func (s *GetSecretParametersByPathRequest) SetNextToken(v string) *GetSecretParametersByPathRequest {
	s.NextToken = &v
	return s
}

func (s *GetSecretParametersByPathRequest) SetPath(v string) *GetSecretParametersByPathRequest {
	s.Path = &v
	return s
}

func (s *GetSecretParametersByPathRequest) SetRecursive(v bool) *GetSecretParametersByPathRequest {
	s.Recursive = &v
	return s
}

func (s *GetSecretParametersByPathRequest) SetRegionId(v string) *GetSecretParametersByPathRequest {
	s.RegionId = &v
	return s
}

func (s *GetSecretParametersByPathRequest) SetWithDecryption(v bool) *GetSecretParametersByPathRequest {
	s.WithDecryption = &v
	return s
}

type GetSecretParametersByPathResponseBody struct {
	// The number of entries per page.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// A pagination token. It can be used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// gAAAAABfTgv5ewUWmNdJ3g7JVLvX70sPH90GZOVGC
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The information about the encryption parameters.
	Parameters []*GetSecretParametersByPathResponseBodyParameters `json:"Parameters,omitempty" xml:"Parameters,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 25156E99-7437-4590-AA58-2ACA17DE405C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of returned entries.
	//
	// example:
	//
	// 10
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s GetSecretParametersByPathResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSecretParametersByPathResponseBody) GoString() string {
	return s.String()
}

func (s *GetSecretParametersByPathResponseBody) SetMaxResults(v int32) *GetSecretParametersByPathResponseBody {
	s.MaxResults = &v
	return s
}

func (s *GetSecretParametersByPathResponseBody) SetNextToken(v string) *GetSecretParametersByPathResponseBody {
	s.NextToken = &v
	return s
}

func (s *GetSecretParametersByPathResponseBody) SetParameters(v []*GetSecretParametersByPathResponseBodyParameters) *GetSecretParametersByPathResponseBody {
	s.Parameters = v
	return s
}

func (s *GetSecretParametersByPathResponseBody) SetRequestId(v string) *GetSecretParametersByPathResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSecretParametersByPathResponseBody) SetTotalCount(v int32) *GetSecretParametersByPathResponseBody {
	s.TotalCount = &v
	return s
}

type GetSecretParametersByPathResponseBodyParameters struct {
	// The constraints of the encryption parameter.
	//
	// example:
	//
	// {\\"AllowedPattern\\": \\"^[a-g]*$\\"}
	Constraints *string `json:"Constraints,omitempty" xml:"Constraints,omitempty"`
	// The user who created the encryption parameter.
	//
	// example:
	//
	// root(130900000)
	CreatedBy *string `json:"CreatedBy,omitempty" xml:"CreatedBy,omitempty"`
	// The time when the encryption parameter was updated.
	//
	// example:
	//
	// 2020-10-21T06:22:48Z
	CreatedDate *string `json:"CreatedDate,omitempty" xml:"CreatedDate,omitempty"`
	// The description of the encryption parameter.
	//
	// example:
	//
	// secretParameter
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the encryption parameter.
	//
	// example:
	//
	// p-7cdc0000000000000000
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The ID of the key.
	//
	// example:
	//
	// 090xxbex-xexx-xxxx-axfc-ddxxcxxxxcex
	KeyId *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	// The name of the encryption parameter.
	//
	// example:
	//
	// mySecretParameter
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The version number of the encryption parameter.
	//
	// example:
	//
	// 1
	ParameterVersion *int32 `json:"ParameterVersion,omitempty" xml:"ParameterVersion,omitempty"`
	// The share type of the encryption parameter.
	//
	// example:
	//
	// Private
	ShareType *string `json:"ShareType,omitempty" xml:"ShareType,omitempty"`
	// The data type of the encryption parameter.
	//
	// example:
	//
	// Secret
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The user who updated the encryption parameter.
	//
	// example:
	//
	// root(130900000)
	UpdatedBy *string `json:"UpdatedBy,omitempty" xml:"UpdatedBy,omitempty"`
	// The time when the encryption parameter was updated.
	//
	// example:
	//
	// 2020-10-21T06:22:48Z
	UpdatedDate *string `json:"UpdatedDate,omitempty" xml:"UpdatedDate,omitempty"`
	// The value of the encryption parameter.
	//
	// example:
	//
	// secretParameter
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetSecretParametersByPathResponseBodyParameters) String() string {
	return tea.Prettify(s)
}

func (s GetSecretParametersByPathResponseBodyParameters) GoString() string {
	return s.String()
}

func (s *GetSecretParametersByPathResponseBodyParameters) SetConstraints(v string) *GetSecretParametersByPathResponseBodyParameters {
	s.Constraints = &v
	return s
}

func (s *GetSecretParametersByPathResponseBodyParameters) SetCreatedBy(v string) *GetSecretParametersByPathResponseBodyParameters {
	s.CreatedBy = &v
	return s
}

func (s *GetSecretParametersByPathResponseBodyParameters) SetCreatedDate(v string) *GetSecretParametersByPathResponseBodyParameters {
	s.CreatedDate = &v
	return s
}

func (s *GetSecretParametersByPathResponseBodyParameters) SetDescription(v string) *GetSecretParametersByPathResponseBodyParameters {
	s.Description = &v
	return s
}

func (s *GetSecretParametersByPathResponseBodyParameters) SetId(v string) *GetSecretParametersByPathResponseBodyParameters {
	s.Id = &v
	return s
}

func (s *GetSecretParametersByPathResponseBodyParameters) SetKeyId(v string) *GetSecretParametersByPathResponseBodyParameters {
	s.KeyId = &v
	return s
}

func (s *GetSecretParametersByPathResponseBodyParameters) SetName(v string) *GetSecretParametersByPathResponseBodyParameters {
	s.Name = &v
	return s
}

func (s *GetSecretParametersByPathResponseBodyParameters) SetParameterVersion(v int32) *GetSecretParametersByPathResponseBodyParameters {
	s.ParameterVersion = &v
	return s
}

func (s *GetSecretParametersByPathResponseBodyParameters) SetShareType(v string) *GetSecretParametersByPathResponseBodyParameters {
	s.ShareType = &v
	return s
}

func (s *GetSecretParametersByPathResponseBodyParameters) SetType(v string) *GetSecretParametersByPathResponseBodyParameters {
	s.Type = &v
	return s
}

func (s *GetSecretParametersByPathResponseBodyParameters) SetUpdatedBy(v string) *GetSecretParametersByPathResponseBodyParameters {
	s.UpdatedBy = &v
	return s
}

func (s *GetSecretParametersByPathResponseBodyParameters) SetUpdatedDate(v string) *GetSecretParametersByPathResponseBodyParameters {
	s.UpdatedDate = &v
	return s
}

func (s *GetSecretParametersByPathResponseBodyParameters) SetValue(v string) *GetSecretParametersByPathResponseBodyParameters {
	s.Value = &v
	return s
}

type GetSecretParametersByPathResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSecretParametersByPathResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSecretParametersByPathResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSecretParametersByPathResponse) GoString() string {
	return s.String()
}

func (s *GetSecretParametersByPathResponse) SetHeaders(v map[string]*string) *GetSecretParametersByPathResponse {
	s.Headers = v
	return s
}

func (s *GetSecretParametersByPathResponse) SetStatusCode(v int32) *GetSecretParametersByPathResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSecretParametersByPathResponse) SetBody(v *GetSecretParametersByPathResponseBody) *GetSecretParametersByPathResponse {
	s.Body = v
	return s
}

type GetServiceSettingsRequest struct {
	// The ID of the region.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetServiceSettingsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetServiceSettingsRequest) GoString() string {
	return s.String()
}

func (s *GetServiceSettingsRequest) SetRegionId(v string) *GetServiceSettingsRequest {
	s.RegionId = &v
	return s
}

type GetServiceSettingsResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 9F755DC9-C0CF-4598-B2E3-2CC763F18CB2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information of service settings.
	ServiceSettings []*GetServiceSettingsResponseBodyServiceSettings `json:"ServiceSettings,omitempty" xml:"ServiceSettings,omitempty" type:"Repeated"`
}

func (s GetServiceSettingsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetServiceSettingsResponseBody) GoString() string {
	return s.String()
}

func (s *GetServiceSettingsResponseBody) SetRequestId(v string) *GetServiceSettingsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetServiceSettingsResponseBody) SetServiceSettings(v []*GetServiceSettingsResponseBodyServiceSettings) *GetServiceSettingsResponseBody {
	s.ServiceSettings = v
	return s
}

type GetServiceSettingsResponseBodyServiceSettings struct {
	// The name of OSS bucket to deliver.
	//
	// example:
	//
	// OssBucketName
	DeliveryOssBucketName *string `json:"DeliveryOssBucketName,omitempty" xml:"DeliveryOssBucketName,omitempty"`
	// Whether to enable OSS delivery.
	//
	// example:
	//
	// false
	DeliveryOssEnabled *bool `json:"DeliveryOssEnabled,omitempty" xml:"DeliveryOssEnabled,omitempty"`
	// The key prefix of OSS to deliver.
	//
	// example:
	//
	// oos/execution
	DeliveryOssKeyPrefix *string `json:"DeliveryOssKeyPrefix,omitempty" xml:"DeliveryOssKeyPrefix,omitempty"`
	// Whether to enable SLS delivery.
	//
	// example:
	//
	// false
	DeliverySlsEnabled *bool `json:"DeliverySlsEnabled,omitempty" xml:"DeliverySlsEnabled,omitempty"`
	// The name of SLS project to deliver.
	//
	// example:
	//
	// SlsProjectName
	DeliverySlsProjectName *string `json:"DeliverySlsProjectName,omitempty" xml:"DeliverySlsProjectName,omitempty"`
	// The id of RDC Enterprise.
	//
	// example:
	//
	// RdcEnterpriseId
	RdcEnterpriseId *string `json:"RdcEnterpriseId,omitempty" xml:"RdcEnterpriseId,omitempty"`
}

func (s GetServiceSettingsResponseBodyServiceSettings) String() string {
	return tea.Prettify(s)
}

func (s GetServiceSettingsResponseBodyServiceSettings) GoString() string {
	return s.String()
}

func (s *GetServiceSettingsResponseBodyServiceSettings) SetDeliveryOssBucketName(v string) *GetServiceSettingsResponseBodyServiceSettings {
	s.DeliveryOssBucketName = &v
	return s
}

func (s *GetServiceSettingsResponseBodyServiceSettings) SetDeliveryOssEnabled(v bool) *GetServiceSettingsResponseBodyServiceSettings {
	s.DeliveryOssEnabled = &v
	return s
}

func (s *GetServiceSettingsResponseBodyServiceSettings) SetDeliveryOssKeyPrefix(v string) *GetServiceSettingsResponseBodyServiceSettings {
	s.DeliveryOssKeyPrefix = &v
	return s
}

func (s *GetServiceSettingsResponseBodyServiceSettings) SetDeliverySlsEnabled(v bool) *GetServiceSettingsResponseBodyServiceSettings {
	s.DeliverySlsEnabled = &v
	return s
}

func (s *GetServiceSettingsResponseBodyServiceSettings) SetDeliverySlsProjectName(v string) *GetServiceSettingsResponseBodyServiceSettings {
	s.DeliverySlsProjectName = &v
	return s
}

func (s *GetServiceSettingsResponseBodyServiceSettings) SetRdcEnterpriseId(v string) *GetServiceSettingsResponseBodyServiceSettings {
	s.RdcEnterpriseId = &v
	return s
}

type GetServiceSettingsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetServiceSettingsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetServiceSettingsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetServiceSettingsResponse) GoString() string {
	return s.String()
}

func (s *GetServiceSettingsResponse) SetHeaders(v map[string]*string) *GetServiceSettingsResponse {
	s.Headers = v
	return s
}

func (s *GetServiceSettingsResponse) SetStatusCode(v int32) *GetServiceSettingsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetServiceSettingsResponse) SetBody(v *GetServiceSettingsResponseBody) *GetServiceSettingsResponse {
	s.Body = v
	return s
}

type GetTemplateRequest struct {
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the template. The name can be 1 to 200 characters in length and can contain letters, digits, hyphens (-), and underscores (_). The name cannot start with ALIYUN, ACS, ALIBABA, or ALICLOUD.
	//
	// This parameter is required.
	//
	// example:
	//
	// MyTemplate
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// The version of the template. The default value is the latest version of the template.
	//
	// example:
	//
	// v1
	TemplateVersion *string `json:"TemplateVersion,omitempty" xml:"TemplateVersion,omitempty"`
}

func (s GetTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTemplateRequest) GoString() string {
	return s.String()
}

func (s *GetTemplateRequest) SetRegionId(v string) *GetTemplateRequest {
	s.RegionId = &v
	return s
}

func (s *GetTemplateRequest) SetTemplateName(v string) *GetTemplateRequest {
	s.TemplateName = &v
	return s
}

func (s *GetTemplateRequest) SetTemplateVersion(v string) *GetTemplateRequest {
	s.TemplateVersion = &v
	return s
}

type GetTemplateResponseBody struct {
	// The content of the template.
	//
	// example:
	//
	// "FormatVersion: OOS-2019-06-01\\nDescription:\\n  en:  Creates an ECS image\\n  zh-cn: 创建一个ECS镜像\\n  name-en: Create Image\\n  name-zh-cn: 创建镜像\\n  categories:\\n    - image_manage\\n    - application_manage\\nParameters:\\n  regionId:\\n    Type: String\\n    Label:\\n      en: RegionId\\n      zh-cn: 地域ID\\n    AssociationProperty: RegionId\\n    Default: \\"{{ ACS::RegionId }}\\"\\n  instanceId:\\n    Label:\\n      en: InstanceId\\n      zh-cn: ECS实例ID\\n    Type: String\\n    AssociationProperty: ALIYUN::ECS::Instance::InstanceId\\n    AssociationPropertyMetadata:\\n      RegionId: regionId\\n  imageName:\\n    Label:\\n      en: ImageName\\n      zh-cn: 新镜像的名称\\n    Type: String\\n    Description:\\n      en: <p class=\\"p\\"	Note:</p> <ul class=\\"ul\\"> <li class=\\"li\\">Length is 2~128 English or Chinese characters</li> <li class=\\"li\\"><font color=\\"red\\">must start with big or small letters or Chinese, not http:// and https://. </font></li> <li class=\\"li\\">Can contain numbers, colons (:), underscores (_), or dashes (-). </li> </ul>\\n      zh-cn: <p class=\\"p\\">注意：</p> <ul class=\\"ul\\"> <li class=\\"li\\">长度为2~128个英文或中文字符</li> <li class=\\"li\\"><font color=\\"red\\">必须以大小字母或中文开头，不能以http://和https://开头。</font></li> <li class=\\"li\\">可以包含数字、半角冒号（:）、下划线（_）或者短划线（-）。</li> </ul>\\n  tags:\\n    Label:\\n      en: Tags\\n      zh-cn: 镜像标签\\n    Type: Json\\n    AssociationProperty: Tags\\n    AssociationPropertyMetadata:\\n      ShowSystem: false\\n    Default: []\\n  OOSAssumeRole:\\n    Label:\\n      en: OOSAssumeRole\\n      zh-cn: OOS扮演的RAM角色\\n    Type: String\\n    Default: OOSServiceRole\\nRamRole: \\"{{ OOSAssumeRole }}\\"\\nTasks:\\n- Name: createImage\\n  Action: ACS::ECS::CreateImage\\n  Description:\\n    en: Create new image with the specified image name and instance ID\\n    zh-cn: 通过指定实例ID和镜像名称创建新的镜像\\n  Properties:\\n    regionId: \\"{{ regionId }}\\"\\n    imageName: \\"{{ imageName }}__on_{{ ACS::ExecutionId }}_at_{{ Acs::CurrentDate }}\\"\\n    instanceId: \\"{{ instanceId }}\\"\\n    tags: \\"{{tags}}\\"\\n  Outputs:\\n    imageId:\\n      ValueSelector: imageId\\n      Type: String\\nOutputs:\\n  imageId:\\n    Type: String\\n    Value: \\"{{ createImage.imageId }}\\"\\nMetadata:\\n  ALIYUN::OOS::Interface:\\n    ParameterGroups:\\n      - Parameters:\\n          - regionId\\n          - instanceId\\n        Label:\\n          default:\\n            zh-cn: 选择实例\\n            en: Select Ecs Instances\\n      - Parameters:\\n          - imageName\\n          - tags\\n        Label:\\n          default:\\n            zh-cn: 镜像设置\\n            en: Image Configure\\n      - Parameters:\\n          - OOSAssumeRole\\n        Label:\\n          default:\\n            zh-cn: 高级选项\\n            en: Control Options"
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 5BBE2663-A18E-5261-9BBB-F4832F5294D9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The metadata of the template.
	Template *GetTemplateResponseBodyTemplate `json:"Template,omitempty" xml:"Template,omitempty" type:"Struct"`
}

func (s GetTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *GetTemplateResponseBody) SetContent(v string) *GetTemplateResponseBody {
	s.Content = &v
	return s
}

func (s *GetTemplateResponseBody) SetRequestId(v string) *GetTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTemplateResponseBody) SetTemplate(v *GetTemplateResponseBodyTemplate) *GetTemplateResponseBody {
	s.Template = v
	return s
}

type GetTemplateResponseBodyTemplate struct {
	// The creator of the template.
	//
	// example:
	//
	// ACS
	CreatedBy *string `json:"CreatedBy,omitempty" xml:"CreatedBy,omitempty"`
	// The time when the template was created.
	//
	// example:
	//
	// 2019-05-16T10:26:14Z
	CreatedDate *string `json:"CreatedDate,omitempty" xml:"CreatedDate,omitempty"`
	// The description of the template.
	//
	// example:
	//
	// "{\\"en\\": \\"Creates an ECS image\\", \\"zh-cn\\": \\"创建一个ECS镜像\\", \\"name-en\\": \\"Create Image\\", \\"name-zh-cn\\": \\"创建镜像\\", \\"categories\\": [\\"image_manage\\", \\"application_manage\\"]}"
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Indicates whether the template was configured with a trigger.
	//
	// example:
	//
	// false
	HasTrigger *bool `json:"HasTrigger,omitempty" xml:"HasTrigger,omitempty"`
	// The SHA-256 value of the template content.
	//
	// example:
	//
	// 40fb5e3e08ef6c8a499ff7cd8441194f518028ad08338a84cb70c023a64576f1
	Hash *string `json:"Hash,omitempty" xml:"Hash,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfmxsn4m4******
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The share type of the template. The share type of a user-created template is **Private**.
	//
	// example:
	//
	// Public
	ShareType *string `json:"ShareType,omitempty" xml:"ShareType,omitempty"`
	// The tag keys and values. The number of key-value pairs ranges from 1 to 20.
	//
	// example:
	//
	// {"k1":"k2","k2":"v2"}
	Tags map[string]interface{} `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The format of the template. The system automatically determines whether the format is JSON or YAML.
	//
	// example:
	//
	// YAML
	TemplateFormat *string `json:"TemplateFormat,omitempty" xml:"TemplateFormat,omitempty"`
	// The ID of the template.
	//
	// example:
	//
	// t-4bdb1745c171401883a2
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The name of the template.
	//
	// example:
	//
	// ACS-ECS-CreateImage
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// The type of the template.
	//
	// example:
	//
	// Automation
	TemplateType *string `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
	// The version of the template. The name of the version consists of the letter v and a number. The number starts from 1.
	//
	// example:
	//
	// v15
	TemplateVersion *string `json:"TemplateVersion,omitempty" xml:"TemplateVersion,omitempty"`
	// The user who last updated the template.
	//
	// example:
	//
	// ACS
	UpdatedBy *string `json:"UpdatedBy,omitempty" xml:"UpdatedBy,omitempty"`
	// The time when the template was last updated.
	//
	// example:
	//
	// 2022-04-26T08:37:07Z
	UpdatedDate *string `json:"UpdatedDate,omitempty" xml:"UpdatedDate,omitempty"`
	// The name of the version of the template.
	//
	// example:
	//
	// version15
	VersionName *string `json:"VersionName,omitempty" xml:"VersionName,omitempty"`
}

func (s GetTemplateResponseBodyTemplate) String() string {
	return tea.Prettify(s)
}

func (s GetTemplateResponseBodyTemplate) GoString() string {
	return s.String()
}

func (s *GetTemplateResponseBodyTemplate) SetCreatedBy(v string) *GetTemplateResponseBodyTemplate {
	s.CreatedBy = &v
	return s
}

func (s *GetTemplateResponseBodyTemplate) SetCreatedDate(v string) *GetTemplateResponseBodyTemplate {
	s.CreatedDate = &v
	return s
}

func (s *GetTemplateResponseBodyTemplate) SetDescription(v string) *GetTemplateResponseBodyTemplate {
	s.Description = &v
	return s
}

func (s *GetTemplateResponseBodyTemplate) SetHasTrigger(v bool) *GetTemplateResponseBodyTemplate {
	s.HasTrigger = &v
	return s
}

func (s *GetTemplateResponseBodyTemplate) SetHash(v string) *GetTemplateResponseBodyTemplate {
	s.Hash = &v
	return s
}

func (s *GetTemplateResponseBodyTemplate) SetResourceGroupId(v string) *GetTemplateResponseBodyTemplate {
	s.ResourceGroupId = &v
	return s
}

func (s *GetTemplateResponseBodyTemplate) SetShareType(v string) *GetTemplateResponseBodyTemplate {
	s.ShareType = &v
	return s
}

func (s *GetTemplateResponseBodyTemplate) SetTags(v map[string]interface{}) *GetTemplateResponseBodyTemplate {
	s.Tags = v
	return s
}

func (s *GetTemplateResponseBodyTemplate) SetTemplateFormat(v string) *GetTemplateResponseBodyTemplate {
	s.TemplateFormat = &v
	return s
}

func (s *GetTemplateResponseBodyTemplate) SetTemplateId(v string) *GetTemplateResponseBodyTemplate {
	s.TemplateId = &v
	return s
}

func (s *GetTemplateResponseBodyTemplate) SetTemplateName(v string) *GetTemplateResponseBodyTemplate {
	s.TemplateName = &v
	return s
}

func (s *GetTemplateResponseBodyTemplate) SetTemplateType(v string) *GetTemplateResponseBodyTemplate {
	s.TemplateType = &v
	return s
}

func (s *GetTemplateResponseBodyTemplate) SetTemplateVersion(v string) *GetTemplateResponseBodyTemplate {
	s.TemplateVersion = &v
	return s
}

func (s *GetTemplateResponseBodyTemplate) SetUpdatedBy(v string) *GetTemplateResponseBodyTemplate {
	s.UpdatedBy = &v
	return s
}

func (s *GetTemplateResponseBodyTemplate) SetUpdatedDate(v string) *GetTemplateResponseBodyTemplate {
	s.UpdatedDate = &v
	return s
}

func (s *GetTemplateResponseBodyTemplate) SetVersionName(v string) *GetTemplateResponseBodyTemplate {
	s.VersionName = &v
	return s
}

type GetTemplateResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTemplateResponse) GoString() string {
	return s.String()
}

func (s *GetTemplateResponse) SetHeaders(v map[string]*string) *GetTemplateResponse {
	s.Headers = v
	return s
}

func (s *GetTemplateResponse) SetStatusCode(v int32) *GetTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTemplateResponse) SetBody(v *GetTemplateResponseBody) *GetTemplateResponse {
	s.Body = v
	return s
}

type ListActionsRequest struct {
	// The number of entries to return on each page. Valid values: 20 to 100. Default value: 50.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that is used to retrieve the next page of results.
	//
	// example:
	//
	// -
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The name of the action. All actions whose names contain the specified action name are returned.
	//
	// example:
	//
	// MyTemplate
	OOSActionName *string `json:"OOSActionName,omitempty" xml:"OOSActionName,omitempty"`
	// The ID of the region.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListActionsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListActionsRequest) GoString() string {
	return s.String()
}

func (s *ListActionsRequest) SetMaxResults(v int32) *ListActionsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListActionsRequest) SetNextToken(v string) *ListActionsRequest {
	s.NextToken = &v
	return s
}

func (s *ListActionsRequest) SetOOSActionName(v string) *ListActionsRequest {
	s.OOSActionName = &v
	return s
}

func (s *ListActionsRequest) SetRegionId(v string) *ListActionsRequest {
	s.RegionId = &v
	return s
}

type ListActionsResponseBody struct {
	// The details of the actions.
	Actions []*ListActionsResponseBodyActions `json:"Actions,omitempty" xml:"Actions,omitempty" type:"Repeated"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that is used to retrieve the next page of results.
	//
	// example:
	//
	// xxx
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// F9154C02-F847-4563-BB6A-6DD01A4F0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListActionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListActionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListActionsResponseBody) SetActions(v []*ListActionsResponseBodyActions) *ListActionsResponseBody {
	s.Actions = v
	return s
}

func (s *ListActionsResponseBody) SetMaxResults(v int32) *ListActionsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListActionsResponseBody) SetNextToken(v string) *ListActionsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListActionsResponseBody) SetRequestId(v string) *ListActionsResponseBody {
	s.RequestId = &v
	return s
}

type ListActionsResponseBodyActions struct {
	// The type of the action.
	//
	// 1.  Atomic actions
	//
	//     	- Atomic.API
	//
	//     	- Atomic.Trigger
	//
	//     	- Atomic.Control
	//
	//     	- Atomic.Embedded
	//
	// 2.  Cloud product actions
	//
	//     	- Product.ECS
	//
	//     	- Product.RDS
	//
	//     	- Product.VPC
	//
	//     	- Product.FC
	//
	//     	- ...
	//
	// example:
	//
	// ACS::Template
	ActionType *string `json:"ActionType,omitempty" xml:"ActionType,omitempty"`
	// The time when the action was created.
	//
	// example:
	//
	// 2019-05-16T10:26:14Z
	CreatedDate *string `json:"CreatedDate,omitempty" xml:"CreatedDate,omitempty"`
	// The description of the action.
	//
	// example:
	//
	// ReplaceSystemDisk
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the action.
	//
	// example:
	//
	// ACS::ECS::ReplaceSystemDisk
	OOSActionName *string `json:"OOSActionName,omitempty" xml:"OOSActionName,omitempty"`
	// The number of times that the action is used.
	//
	// example:
	//
	// 5
	Popularity *int32 `json:"Popularity,omitempty" xml:"Popularity,omitempty"`
	// The parameters of the action.
	//
	// example:
	//
	// { "ImageId": { "Description": "The mirror ID you will use when resetting the system", "Type": "String" }, "InstanceId": { "Description": "the instance id that you will handle .", "Type": "String" } }
	Properties *string `json:"Properties,omitempty" xml:"Properties,omitempty"`
	// The version of the template that corresponds to the action.
	//
	// >  For atomic actions, this parameter is not returned.
	//
	// example:
	//
	// v1
	TemplateVersion *string `json:"TemplateVersion,omitempty" xml:"TemplateVersion,omitempty"`
}

func (s ListActionsResponseBodyActions) String() string {
	return tea.Prettify(s)
}

func (s ListActionsResponseBodyActions) GoString() string {
	return s.String()
}

func (s *ListActionsResponseBodyActions) SetActionType(v string) *ListActionsResponseBodyActions {
	s.ActionType = &v
	return s
}

func (s *ListActionsResponseBodyActions) SetCreatedDate(v string) *ListActionsResponseBodyActions {
	s.CreatedDate = &v
	return s
}

func (s *ListActionsResponseBodyActions) SetDescription(v string) *ListActionsResponseBodyActions {
	s.Description = &v
	return s
}

func (s *ListActionsResponseBodyActions) SetOOSActionName(v string) *ListActionsResponseBodyActions {
	s.OOSActionName = &v
	return s
}

func (s *ListActionsResponseBodyActions) SetPopularity(v int32) *ListActionsResponseBodyActions {
	s.Popularity = &v
	return s
}

func (s *ListActionsResponseBodyActions) SetProperties(v string) *ListActionsResponseBodyActions {
	s.Properties = &v
	return s
}

func (s *ListActionsResponseBodyActions) SetTemplateVersion(v string) *ListActionsResponseBodyActions {
	s.TemplateVersion = &v
	return s
}

type ListActionsResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListActionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListActionsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListActionsResponse) GoString() string {
	return s.String()
}

func (s *ListActionsResponse) SetHeaders(v map[string]*string) *ListActionsResponse {
	s.Headers = v
	return s
}

func (s *ListActionsResponse) SetStatusCode(v int32) *ListActionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListActionsResponse) SetBody(v *ListActionsResponseBody) *ListActionsResponse {
	s.Body = v
	return s
}

type ListApplicationGroupsRequest struct {
	// The name of the application.
	//
	// example:
	//
	// MyApplication
	ApplicationName *string `json:"ApplicationName,omitempty" xml:"ApplicationName,omitempty"`
	// The ID of the region in which the related resources reside.
	//
	// example:
	//
	// cn-hangzhou
	DeployRegionId *string `json:"DeployRegionId,omitempty" xml:"DeployRegionId,omitempty"`
	// The number of entries to return on each page.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that is used to retrieve the next page of results.
	//
	// example:
	//
	// -
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the region. Set the value to cn-hangzhou.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the cloud resource.
	//
	// example:
	//
	// i-2vcj9raxrhxb48zz3whw
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The code of the product to which the cloud resource belongs.
	//
	// example:
	//
	// ecs
	ResourceProduct *string `json:"ResourceProduct,omitempty" xml:"ResourceProduct,omitempty"`
	// The type of the cloud resource.
	//
	// example:
	//
	// instance
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s ListApplicationGroupsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListApplicationGroupsRequest) GoString() string {
	return s.String()
}

func (s *ListApplicationGroupsRequest) SetApplicationName(v string) *ListApplicationGroupsRequest {
	s.ApplicationName = &v
	return s
}

func (s *ListApplicationGroupsRequest) SetDeployRegionId(v string) *ListApplicationGroupsRequest {
	s.DeployRegionId = &v
	return s
}

func (s *ListApplicationGroupsRequest) SetMaxResults(v int32) *ListApplicationGroupsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListApplicationGroupsRequest) SetNextToken(v string) *ListApplicationGroupsRequest {
	s.NextToken = &v
	return s
}

func (s *ListApplicationGroupsRequest) SetRegionId(v string) *ListApplicationGroupsRequest {
	s.RegionId = &v
	return s
}

func (s *ListApplicationGroupsRequest) SetResourceId(v string) *ListApplicationGroupsRequest {
	s.ResourceId = &v
	return s
}

func (s *ListApplicationGroupsRequest) SetResourceProduct(v string) *ListApplicationGroupsRequest {
	s.ResourceProduct = &v
	return s
}

func (s *ListApplicationGroupsRequest) SetResourceType(v string) *ListApplicationGroupsRequest {
	s.ResourceType = &v
	return s
}

type ListApplicationGroupsResponseBody struct {
	// The details of the application group.
	ApplicationGroups []*ListApplicationGroupsResponseBodyApplicationGroups `json:"ApplicationGroups,omitempty" xml:"ApplicationGroups,omitempty" type:"Repeated"`
	// The number of entries returned on each page.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that is used to retrieve the next page of results.
	//
	// example:
	//
	// -
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 69D97BF2-5DF2-544C-A650-36A474E17BC3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListApplicationGroupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListApplicationGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *ListApplicationGroupsResponseBody) SetApplicationGroups(v []*ListApplicationGroupsResponseBodyApplicationGroups) *ListApplicationGroupsResponseBody {
	s.ApplicationGroups = v
	return s
}

func (s *ListApplicationGroupsResponseBody) SetMaxResults(v int32) *ListApplicationGroupsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListApplicationGroupsResponseBody) SetNextToken(v string) *ListApplicationGroupsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListApplicationGroupsResponseBody) SetRequestId(v string) *ListApplicationGroupsResponseBody {
	s.RequestId = &v
	return s
}

type ListApplicationGroupsResponseBodyApplicationGroups struct {
	// The name of the application.
	//
	// example:
	//
	// MyApplication
	ApplicationName *string `json:"ApplicationName,omitempty" xml:"ApplicationName,omitempty"`
	// The ID of the application group in CloudMonitor.
	//
	// example:
	//
	// 12345678
	CmsGroupId *string `json:"CmsGroupId,omitempty" xml:"CmsGroupId,omitempty"`
	// The time when the application group was created.
	//
	// example:
	//
	// 2021-09-07T10:28:25Z
	CreateDate *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	// The configuration information of the application group.
	//
	// example:
	//
	// {   "TemplateURL": "https://ros-template.oss-cn-zhangjiakou.aliyuncs.com/App_Management_Existing_Vpc_Ecs_Instance.json",   "Parameters": {     "ZoneId": "cn-hangzhou-k",     "ProjectName": "test",     "SystemDiskSize": 40,     "InstanceChargeType": "PostPaid",     "SecurityGroupId": "sg-bp1a4374akk63jl8tddy",     "VSwitchId": "vsw-bp1fcvc3zn0jrag86rrlm",     "SystemDiskCategory": "cloud_essd",     "InstancePassword": "******",     "InternetChargeType": "PayByTraffic",     "InstanceCount": 1,     "InternetMaxBandwidthOut": 0,     "VpcId": "vpc-bp1i99boyas8i8m9t3skp",     "EcsImageId": "centos_8_5_x64_20G_alibase_20211228.vhd",     "DataDiskSize": 100,     "EcsInstanceType": "ecs.s6-c1m4.small",     "DataDiskCategory": "cloud_efficiency",     "EnvironmentCommandId": "c-hz028fc3g031gcg"   },   "RegionId": "cn-hangzhou",   "StackName": "stack-1645688523068-3no_AKhOJ",   "DisableRollback": true }
	DeployParameters *string `json:"DeployParameters,omitempty" xml:"DeployParameters,omitempty"`
	// The ID of the region in which the related resources reside.
	//
	// example:
	//
	// cn-hangzhou
	DeployRegionId *string `json:"DeployRegionId,omitempty" xml:"DeployRegionId,omitempty"`
	// The description of the application group.
	//
	// example:
	//
	// ApplicationGroup
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The tag key.
	//
	// example:
	//
	// k1
	ImportTagKey *string `json:"ImportTagKey,omitempty" xml:"ImportTagKey,omitempty"`
	// The tag value.
	//
	// example:
	//
	// v1
	ImportTagValue *string `json:"ImportTagValue,omitempty" xml:"ImportTagValue,omitempty"`
	// The name of the application group.
	//
	// example:
	//
	// UpdateMyApplicationGroup
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The state of the application group.
	//
	// example:
	//
	// Created
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The state information of the application group.
	//
	// example:
	//
	// ApplicationGroup is Created.
	StatusReason *string `json:"StatusReason,omitempty" xml:"StatusReason,omitempty"`
	// The time when the application group was updated.
	//
	// example:
	//
	// 2021-09-08T03:01:53Z
	UpdateDate *string `json:"UpdateDate,omitempty" xml:"UpdateDate,omitempty"`
}

func (s ListApplicationGroupsResponseBodyApplicationGroups) String() string {
	return tea.Prettify(s)
}

func (s ListApplicationGroupsResponseBodyApplicationGroups) GoString() string {
	return s.String()
}

func (s *ListApplicationGroupsResponseBodyApplicationGroups) SetApplicationName(v string) *ListApplicationGroupsResponseBodyApplicationGroups {
	s.ApplicationName = &v
	return s
}

func (s *ListApplicationGroupsResponseBodyApplicationGroups) SetCmsGroupId(v string) *ListApplicationGroupsResponseBodyApplicationGroups {
	s.CmsGroupId = &v
	return s
}

func (s *ListApplicationGroupsResponseBodyApplicationGroups) SetCreateDate(v string) *ListApplicationGroupsResponseBodyApplicationGroups {
	s.CreateDate = &v
	return s
}

func (s *ListApplicationGroupsResponseBodyApplicationGroups) SetDeployParameters(v string) *ListApplicationGroupsResponseBodyApplicationGroups {
	s.DeployParameters = &v
	return s
}

func (s *ListApplicationGroupsResponseBodyApplicationGroups) SetDeployRegionId(v string) *ListApplicationGroupsResponseBodyApplicationGroups {
	s.DeployRegionId = &v
	return s
}

func (s *ListApplicationGroupsResponseBodyApplicationGroups) SetDescription(v string) *ListApplicationGroupsResponseBodyApplicationGroups {
	s.Description = &v
	return s
}

func (s *ListApplicationGroupsResponseBodyApplicationGroups) SetImportTagKey(v string) *ListApplicationGroupsResponseBodyApplicationGroups {
	s.ImportTagKey = &v
	return s
}

func (s *ListApplicationGroupsResponseBodyApplicationGroups) SetImportTagValue(v string) *ListApplicationGroupsResponseBodyApplicationGroups {
	s.ImportTagValue = &v
	return s
}

func (s *ListApplicationGroupsResponseBodyApplicationGroups) SetName(v string) *ListApplicationGroupsResponseBodyApplicationGroups {
	s.Name = &v
	return s
}

func (s *ListApplicationGroupsResponseBodyApplicationGroups) SetStatus(v string) *ListApplicationGroupsResponseBodyApplicationGroups {
	s.Status = &v
	return s
}

func (s *ListApplicationGroupsResponseBodyApplicationGroups) SetStatusReason(v string) *ListApplicationGroupsResponseBodyApplicationGroups {
	s.StatusReason = &v
	return s
}

func (s *ListApplicationGroupsResponseBodyApplicationGroups) SetUpdateDate(v string) *ListApplicationGroupsResponseBodyApplicationGroups {
	s.UpdateDate = &v
	return s
}

type ListApplicationGroupsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListApplicationGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListApplicationGroupsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListApplicationGroupsResponse) GoString() string {
	return s.String()
}

func (s *ListApplicationGroupsResponse) SetHeaders(v map[string]*string) *ListApplicationGroupsResponse {
	s.Headers = v
	return s
}

func (s *ListApplicationGroupsResponse) SetStatusCode(v int32) *ListApplicationGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListApplicationGroupsResponse) SetBody(v *ListApplicationGroupsResponseBody) *ListApplicationGroupsResponse {
	s.Body = v
	return s
}

type ListApplicationsRequest struct {
	// The type of the application.
	//
	// Valid values:
	//
	// 	- ComputeNest
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- Custom
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- DingTalk
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// DingTalk
	ApplicationType *string `json:"ApplicationType,omitempty" xml:"ApplicationType,omitempty"`
	// The number of entries to return on each page. Valid values: 10 to 100. Default value: 50.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The name of the application.
	//
	// example:
	//
	// "MyApplications"
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The names of the applications.
	//
	// example:
	//
	// ["MyApplication"]
	Names *string `json:"Names,omitempty" xml:"Names,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// -
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The region ID. Set the value to cn-hangzhou.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The tags.
	//
	// example:
	//
	// {"k1": "v1","k2": "v2"}
	Tags map[string]interface{} `json:"Tags,omitempty" xml:"Tags,omitempty"`
}

func (s ListApplicationsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListApplicationsRequest) GoString() string {
	return s.String()
}

func (s *ListApplicationsRequest) SetApplicationType(v string) *ListApplicationsRequest {
	s.ApplicationType = &v
	return s
}

func (s *ListApplicationsRequest) SetMaxResults(v int32) *ListApplicationsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListApplicationsRequest) SetName(v string) *ListApplicationsRequest {
	s.Name = &v
	return s
}

func (s *ListApplicationsRequest) SetNames(v string) *ListApplicationsRequest {
	s.Names = &v
	return s
}

func (s *ListApplicationsRequest) SetNextToken(v string) *ListApplicationsRequest {
	s.NextToken = &v
	return s
}

func (s *ListApplicationsRequest) SetRegionId(v string) *ListApplicationsRequest {
	s.RegionId = &v
	return s
}

func (s *ListApplicationsRequest) SetTags(v map[string]interface{}) *ListApplicationsRequest {
	s.Tags = v
	return s
}

type ListApplicationsShrinkRequest struct {
	// The type of the application.
	//
	// Valid values:
	//
	// 	- ComputeNest
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- Custom
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- DingTalk
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// DingTalk
	ApplicationType *string `json:"ApplicationType,omitempty" xml:"ApplicationType,omitempty"`
	// The number of entries to return on each page. Valid values: 10 to 100. Default value: 50.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The name of the application.
	//
	// example:
	//
	// "MyApplications"
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The names of the applications.
	//
	// example:
	//
	// ["MyApplication"]
	Names *string `json:"Names,omitempty" xml:"Names,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// -
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The region ID. Set the value to cn-hangzhou.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The tags.
	//
	// example:
	//
	// {"k1": "v1","k2": "v2"}
	TagsShrink *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
}

func (s ListApplicationsShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListApplicationsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListApplicationsShrinkRequest) SetApplicationType(v string) *ListApplicationsShrinkRequest {
	s.ApplicationType = &v
	return s
}

func (s *ListApplicationsShrinkRequest) SetMaxResults(v int32) *ListApplicationsShrinkRequest {
	s.MaxResults = &v
	return s
}

func (s *ListApplicationsShrinkRequest) SetName(v string) *ListApplicationsShrinkRequest {
	s.Name = &v
	return s
}

func (s *ListApplicationsShrinkRequest) SetNames(v string) *ListApplicationsShrinkRequest {
	s.Names = &v
	return s
}

func (s *ListApplicationsShrinkRequest) SetNextToken(v string) *ListApplicationsShrinkRequest {
	s.NextToken = &v
	return s
}

func (s *ListApplicationsShrinkRequest) SetRegionId(v string) *ListApplicationsShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *ListApplicationsShrinkRequest) SetTagsShrink(v string) *ListApplicationsShrinkRequest {
	s.TagsShrink = &v
	return s
}

type ListApplicationsResponseBody struct {
	// The details of the application.
	Applications []*ListApplicationsResponseBodyApplications `json:"Applications,omitempty" xml:"Applications,omitempty" type:"Repeated"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// -
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 12067D53-56A9-561B-ADD6-61429D207117
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListApplicationsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListApplicationsResponseBody) GoString() string {
	return s.String()
}

func (s *ListApplicationsResponseBody) SetApplications(v []*ListApplicationsResponseBodyApplications) *ListApplicationsResponseBody {
	s.Applications = v
	return s
}

func (s *ListApplicationsResponseBody) SetMaxResults(v int32) *ListApplicationsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListApplicationsResponseBody) SetNextToken(v string) *ListApplicationsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListApplicationsResponseBody) SetRequestId(v string) *ListApplicationsResponseBody {
	s.RequestId = &v
	return s
}

type ListApplicationsResponseBodyApplications struct {
	// The type of the application.
	//
	// example:
	//
	// DingTalk
	ApplicationType *string `json:"ApplicationType,omitempty" xml:"ApplicationType,omitempty"`
	// The time when the application was created.
	//
	// example:
	//
	// 2021-09-07T09:09:59Z
	CreateDate *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	// The description of the application.
	//
	// example:
	//
	// Application
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the application.
	//
	// example:
	//
	// MyApplication
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfmxsn4m******
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The tags added to the application.
	//
	// example:
	//
	// {"k1": "v1","k2": "v2"}
	Tags map[string]interface{} `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The time when the application was updated.
	//
	// example:
	//
	// 2021-09-07T09:09:59Z
	UpdateDate *string `json:"UpdateDate,omitempty" xml:"UpdateDate,omitempty"`
}

func (s ListApplicationsResponseBodyApplications) String() string {
	return tea.Prettify(s)
}

func (s ListApplicationsResponseBodyApplications) GoString() string {
	return s.String()
}

func (s *ListApplicationsResponseBodyApplications) SetApplicationType(v string) *ListApplicationsResponseBodyApplications {
	s.ApplicationType = &v
	return s
}

func (s *ListApplicationsResponseBodyApplications) SetCreateDate(v string) *ListApplicationsResponseBodyApplications {
	s.CreateDate = &v
	return s
}

func (s *ListApplicationsResponseBodyApplications) SetDescription(v string) *ListApplicationsResponseBodyApplications {
	s.Description = &v
	return s
}

func (s *ListApplicationsResponseBodyApplications) SetName(v string) *ListApplicationsResponseBodyApplications {
	s.Name = &v
	return s
}

func (s *ListApplicationsResponseBodyApplications) SetResourceGroupId(v string) *ListApplicationsResponseBodyApplications {
	s.ResourceGroupId = &v
	return s
}

func (s *ListApplicationsResponseBodyApplications) SetTags(v map[string]interface{}) *ListApplicationsResponseBodyApplications {
	s.Tags = v
	return s
}

func (s *ListApplicationsResponseBodyApplications) SetUpdateDate(v string) *ListApplicationsResponseBodyApplications {
	s.UpdateDate = &v
	return s
}

type ListApplicationsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListApplicationsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListApplicationsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListApplicationsResponse) GoString() string {
	return s.String()
}

func (s *ListApplicationsResponse) SetHeaders(v map[string]*string) *ListApplicationsResponse {
	s.Headers = v
	return s
}

func (s *ListApplicationsResponse) SetStatusCode(v int32) *ListApplicationsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListApplicationsResponse) SetBody(v *ListApplicationsResponseBody) *ListApplicationsResponse {
	s.Body = v
	return s
}

type ListExecutionLogsRequest struct {
	// The ID of the execution.
	//
	// This parameter is required.
	//
	// example:
	//
	// exec-xxx
	ExecutionId *string `json:"ExecutionId,omitempty" xml:"ExecutionId,omitempty"`
	// The type of the log.
	//
	// example:
	//
	// System
	LogType *string `json:"LogType,omitempty" xml:"LogType,omitempty"`
	// The number of entries to return on each page.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that is used to retrieve the next page of results.
	//
	// example:
	//
	// MTRBMDc0NjAtRUJFNy00N0NBLTk3NTctMTJDQzQ3NjFENDdB
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the region in which you want to query the logs of the execution.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The execution ID of the task.
	//
	// example:
	//
	// exec-1234567zxcvb.t0010
	TaskExecutionId *string `json:"TaskExecutionId,omitempty" xml:"TaskExecutionId,omitempty"`
}

func (s ListExecutionLogsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListExecutionLogsRequest) GoString() string {
	return s.String()
}

func (s *ListExecutionLogsRequest) SetExecutionId(v string) *ListExecutionLogsRequest {
	s.ExecutionId = &v
	return s
}

func (s *ListExecutionLogsRequest) SetLogType(v string) *ListExecutionLogsRequest {
	s.LogType = &v
	return s
}

func (s *ListExecutionLogsRequest) SetMaxResults(v int32) *ListExecutionLogsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListExecutionLogsRequest) SetNextToken(v string) *ListExecutionLogsRequest {
	s.NextToken = &v
	return s
}

func (s *ListExecutionLogsRequest) SetRegionId(v string) *ListExecutionLogsRequest {
	s.RegionId = &v
	return s
}

func (s *ListExecutionLogsRequest) SetTaskExecutionId(v string) *ListExecutionLogsRequest {
	s.TaskExecutionId = &v
	return s
}

type ListExecutionLogsResponseBody struct {
	// The logs of the execution.
	ExecutionLogs []*ListExecutionLogsResponseBodyExecutionLogs `json:"ExecutionLogs,omitempty" xml:"ExecutionLogs,omitempty" type:"Repeated"`
	// Indicates whether the log is truncated.
	//
	// example:
	//
	// true
	IsTruncated *bool `json:"IsTruncated,omitempty" xml:"IsTruncated,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// A pagination token. It can be used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// gAAAAABdpsGWjX8dJ-a6dl_pvoS7AFxNHSNJKHLCAJJ0ylgA53nWW5V4HTEZKCYTaEPNOrxFir4z43UTOjE150cFr8AGTifA==
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 14A07460-EBE7-47CA-9757-12CC4761D47A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListExecutionLogsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListExecutionLogsResponseBody) GoString() string {
	return s.String()
}

func (s *ListExecutionLogsResponseBody) SetExecutionLogs(v []*ListExecutionLogsResponseBodyExecutionLogs) *ListExecutionLogsResponseBody {
	s.ExecutionLogs = v
	return s
}

func (s *ListExecutionLogsResponseBody) SetIsTruncated(v bool) *ListExecutionLogsResponseBody {
	s.IsTruncated = &v
	return s
}

func (s *ListExecutionLogsResponseBody) SetMaxResults(v int32) *ListExecutionLogsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListExecutionLogsResponseBody) SetNextToken(v string) *ListExecutionLogsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListExecutionLogsResponseBody) SetRequestId(v string) *ListExecutionLogsResponseBody {
	s.RequestId = &v
	return s
}

type ListExecutionLogsResponseBodyExecutionLogs struct {
	// The log type.
	//
	// example:
	//
	// System
	LogType *string `json:"LogType,omitempty" xml:"LogType,omitempty"`
	// The details of the task execution.
	//
	// example:
	//
	// The task CheckDiskCategory completed.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The task execution ID.
	//
	// example:
	//
	// exec-1234567zxcvb.t0010
	TaskExecutionId *string `json:"TaskExecutionId,omitempty" xml:"TaskExecutionId,omitempty"`
	// The timestamp when the task was run.
	//
	// example:
	//
	// 2019-05-24T:02:29:07Z
	Timestamp *string `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
}

func (s ListExecutionLogsResponseBodyExecutionLogs) String() string {
	return tea.Prettify(s)
}

func (s ListExecutionLogsResponseBodyExecutionLogs) GoString() string {
	return s.String()
}

func (s *ListExecutionLogsResponseBodyExecutionLogs) SetLogType(v string) *ListExecutionLogsResponseBodyExecutionLogs {
	s.LogType = &v
	return s
}

func (s *ListExecutionLogsResponseBodyExecutionLogs) SetMessage(v string) *ListExecutionLogsResponseBodyExecutionLogs {
	s.Message = &v
	return s
}

func (s *ListExecutionLogsResponseBodyExecutionLogs) SetTaskExecutionId(v string) *ListExecutionLogsResponseBodyExecutionLogs {
	s.TaskExecutionId = &v
	return s
}

func (s *ListExecutionLogsResponseBodyExecutionLogs) SetTimestamp(v string) *ListExecutionLogsResponseBodyExecutionLogs {
	s.Timestamp = &v
	return s
}

type ListExecutionLogsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListExecutionLogsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListExecutionLogsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListExecutionLogsResponse) GoString() string {
	return s.String()
}

func (s *ListExecutionLogsResponse) SetHeaders(v map[string]*string) *ListExecutionLogsResponse {
	s.Headers = v
	return s
}

func (s *ListExecutionLogsResponse) SetStatusCode(v int32) *ListExecutionLogsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListExecutionLogsResponse) SetBody(v *ListExecutionLogsResponseBody) *ListExecutionLogsResponse {
	s.Body = v
	return s
}

type ListExecutionRiskyTasksRequest struct {
	// The ID of the region.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the template.
	//
	// This parameter is required.
	//
	// example:
	//
	// myTemplate
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
}

func (s ListExecutionRiskyTasksRequest) String() string {
	return tea.Prettify(s)
}

func (s ListExecutionRiskyTasksRequest) GoString() string {
	return s.String()
}

func (s *ListExecutionRiskyTasksRequest) SetRegionId(v string) *ListExecutionRiskyTasksRequest {
	s.RegionId = &v
	return s
}

func (s *ListExecutionRiskyTasksRequest) SetTemplateName(v string) *ListExecutionRiskyTasksRequest {
	s.TemplateName = &v
	return s
}

type ListExecutionRiskyTasksResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// C04B668D-D2DD-4B40-B6E9-0E3C4F53D5B5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about high-risk tasks.
	RiskyTasks []*ListExecutionRiskyTasksResponseBodyRiskyTasks `json:"RiskyTasks,omitempty" xml:"RiskyTasks,omitempty" type:"Repeated"`
}

func (s ListExecutionRiskyTasksResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListExecutionRiskyTasksResponseBody) GoString() string {
	return s.String()
}

func (s *ListExecutionRiskyTasksResponseBody) SetRequestId(v string) *ListExecutionRiskyTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListExecutionRiskyTasksResponseBody) SetRiskyTasks(v []*ListExecutionRiskyTasksResponseBodyRiskyTasks) *ListExecutionRiskyTasksResponseBody {
	s.RiskyTasks = v
	return s
}

type ListExecutionRiskyTasksResponseBodyRiskyTasks struct {
	// The name of the operation that the high-risk task calls.
	//
	// example:
	//
	// DeleteInstance
	API *string `json:"API,omitempty" xml:"API,omitempty"`
	// The cloud service in which the high-risk task runs.
	//
	// example:
	//
	// ECS
	Service *string `json:"Service,omitempty" xml:"Service,omitempty"`
	// The details of the high-risk task.
	Task []*string `json:"Task,omitempty" xml:"Task,omitempty" type:"Repeated"`
	// The details of templates to which the high-risk task belongs.
	Template []*string `json:"Template,omitempty" xml:"Template,omitempty" type:"Repeated"`
}

func (s ListExecutionRiskyTasksResponseBodyRiskyTasks) String() string {
	return tea.Prettify(s)
}

func (s ListExecutionRiskyTasksResponseBodyRiskyTasks) GoString() string {
	return s.String()
}

func (s *ListExecutionRiskyTasksResponseBodyRiskyTasks) SetAPI(v string) *ListExecutionRiskyTasksResponseBodyRiskyTasks {
	s.API = &v
	return s
}

func (s *ListExecutionRiskyTasksResponseBodyRiskyTasks) SetService(v string) *ListExecutionRiskyTasksResponseBodyRiskyTasks {
	s.Service = &v
	return s
}

func (s *ListExecutionRiskyTasksResponseBodyRiskyTasks) SetTask(v []*string) *ListExecutionRiskyTasksResponseBodyRiskyTasks {
	s.Task = v
	return s
}

func (s *ListExecutionRiskyTasksResponseBodyRiskyTasks) SetTemplate(v []*string) *ListExecutionRiskyTasksResponseBodyRiskyTasks {
	s.Template = v
	return s
}

type ListExecutionRiskyTasksResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListExecutionRiskyTasksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListExecutionRiskyTasksResponse) String() string {
	return tea.Prettify(s)
}

func (s ListExecutionRiskyTasksResponse) GoString() string {
	return s.String()
}

func (s *ListExecutionRiskyTasksResponse) SetHeaders(v map[string]*string) *ListExecutionRiskyTasksResponse {
	s.Headers = v
	return s
}

func (s *ListExecutionRiskyTasksResponse) SetStatusCode(v int32) *ListExecutionRiskyTasksResponse {
	s.StatusCode = &v
	return s
}

func (s *ListExecutionRiskyTasksResponse) SetBody(v *ListExecutionRiskyTasksResponseBody) *ListExecutionRiskyTasksResponse {
	s.Body = v
	return s
}

type ListExecutionsRequest struct {
	// The types of the execution template. Valid values: Other, TimerTrigger, EventTrigger, and AlarmTrigger. You can specify only one of the Categories and Category parameters. We recommend that you specify Categories.
	//
	// example:
	//
	// ["TimerTrigger"、"EventTrigger"]
	Categories *string `json:"Categories,omitempty" xml:"Categories,omitempty"`
	// The type of the execution template. Valid values: Other, TimerTrigger, EventTrigger, and AlarmTrigger.
	//
	// example:
	//
	// Other
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The depth of execution. Valid values: RootDepth and FirstChildDepth. If you set this parameter to RootDepth, only the parent execution is returned. If you set this parameter to FirstChildDepth, only the child executions at the first level are returned. You can specify only one of the Depth and IncludeChildExecution parameters. We recommend that you specify Depth.
	//
	// example:
	//
	// RootDepth
	Depth *string `json:"Depth,omitempty" xml:"Depth,omitempty"`
	// The description of the execution.
	//
	// example:
	//
	// MyDescription
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The earliest end time. The executions that stop running at or later than the specified time are queried.
	//
	// example:
	//
	// 2019-05-16T10:26:14Z
	EndDateAfter *string `json:"EndDateAfter,omitempty" xml:"EndDateAfter,omitempty"`
	// The latest end time. The executions that stop running at or earlier than the specified time are queried.
	//
	// example:
	//
	// 2019-05-16T10:26:14Z
	EndDateBefore *string `json:"EndDateBefore,omitempty" xml:"EndDateBefore,omitempty"`
	// The executor.
	//
	// example:
	//
	// vme
	ExecutedBy *string `json:"ExecutedBy,omitempty" xml:"ExecutedBy,omitempty"`
	// The ID of the execution.
	//
	// example:
	//
	// exec-xxx
	ExecutionId *string `json:"ExecutionId,omitempty" xml:"ExecutionId,omitempty"`
	// Specifies whether to include child executions. Default value: False.
	//
	// example:
	//
	// true
	IncludeChildExecution *bool `json:"IncludeChildExecution,omitempty" xml:"IncludeChildExecution,omitempty"`
	// The number of entries to return on each page. Valid values: 10 to 100. Default value: 50.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The execution mode. Valid values:
	//
	// 	- **Automatic**
	//
	// 	- **Debug**
	//
	// example:
	//
	// Automatic
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// The token that is used to retrieve the next page of results.
	//
	// example:
	//
	// MTRBMDc0NjAtRUJFNy00N0NBLTk3NTctMTJDQzQ
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the parent execution.
	//
	// example:
	//
	// exec-xxx
	ParentExecutionId *string `json:"ParentExecutionId,omitempty" xml:"ParentExecutionId,omitempty"`
	// The RAM role.
	//
	// example:
	//
	// OOSServiceRole
	RamRole *string `json:"RamRole,omitempty" xml:"RamRole,omitempty"`
	// The ID of the region.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the instances you want to query belong.
	//
	// example:
	//
	// rg-acfmxsn4m4******
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The ID of the Elastic Compute Service (ECS) resource.
	//
	// example:
	//
	// i-xxx
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The name of the resource template.
	//
	// example:
	//
	// ACS-ECS-TEST
	ResourceTemplateName *string `json:"ResourceTemplateName,omitempty" xml:"ResourceTemplateName,omitempty"`
	// The field that is used to sort the executions to query. Valid values:
	//
	// 	- **StartDate**: specifies that the executions are sorted based on the time when they are created. This is the default value.
	//
	// 	- **EndDate**: specifies that the executions are sorted based on the time when they stop running.
	//
	// 	- **Status**: specifies that the executions are sorted based on their states.
	//
	// example:
	//
	// StartDate
	SortField *string `json:"SortField,omitempty" xml:"SortField,omitempty"`
	// The order in which you want to sort the results. Valid values:
	//
	// 	- **Ascending**: ascending order.
	//
	// 	- **Descending**: descending order. This is the default value.
	//
	// example:
	//
	// Ascending
	SortOrder *string `json:"SortOrder,omitempty" xml:"SortOrder,omitempty"`
	// The earliest start time. The executions that start to run at or later than the specified time are queried.
	//
	// example:
	//
	// 2019-05-16T10:26:14Z
	StartDateAfter *string `json:"StartDateAfter,omitempty" xml:"StartDateAfter,omitempty"`
	// The latest start time. The executions that start to run at or earlier than the specified point in time are queried.
	//
	// example:
	//
	// 2019-05-16T10:26:14Z
	StartDateBefore *string `json:"StartDateBefore,omitempty" xml:"StartDateBefore,omitempty"`
	// The status of the execution. Valid values: Running, Started, Success, Failed, Waiting, Cancelled, Pending, and Skipped.
	//
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tags for the execution.
	//
	// example:
	//
	// {"k1":"v2","k2":"v2"}
	Tags map[string]interface{} `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The name of the template. All templates whose names contain the specified template name are queried.
	//
	// example:
	//
	// MyTemplate
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
}

func (s ListExecutionsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListExecutionsRequest) GoString() string {
	return s.String()
}

func (s *ListExecutionsRequest) SetCategories(v string) *ListExecutionsRequest {
	s.Categories = &v
	return s
}

func (s *ListExecutionsRequest) SetCategory(v string) *ListExecutionsRequest {
	s.Category = &v
	return s
}

func (s *ListExecutionsRequest) SetDepth(v string) *ListExecutionsRequest {
	s.Depth = &v
	return s
}

func (s *ListExecutionsRequest) SetDescription(v string) *ListExecutionsRequest {
	s.Description = &v
	return s
}

func (s *ListExecutionsRequest) SetEndDateAfter(v string) *ListExecutionsRequest {
	s.EndDateAfter = &v
	return s
}

func (s *ListExecutionsRequest) SetEndDateBefore(v string) *ListExecutionsRequest {
	s.EndDateBefore = &v
	return s
}

func (s *ListExecutionsRequest) SetExecutedBy(v string) *ListExecutionsRequest {
	s.ExecutedBy = &v
	return s
}

func (s *ListExecutionsRequest) SetExecutionId(v string) *ListExecutionsRequest {
	s.ExecutionId = &v
	return s
}

func (s *ListExecutionsRequest) SetIncludeChildExecution(v bool) *ListExecutionsRequest {
	s.IncludeChildExecution = &v
	return s
}

func (s *ListExecutionsRequest) SetMaxResults(v int32) *ListExecutionsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListExecutionsRequest) SetMode(v string) *ListExecutionsRequest {
	s.Mode = &v
	return s
}

func (s *ListExecutionsRequest) SetNextToken(v string) *ListExecutionsRequest {
	s.NextToken = &v
	return s
}

func (s *ListExecutionsRequest) SetParentExecutionId(v string) *ListExecutionsRequest {
	s.ParentExecutionId = &v
	return s
}

func (s *ListExecutionsRequest) SetRamRole(v string) *ListExecutionsRequest {
	s.RamRole = &v
	return s
}

func (s *ListExecutionsRequest) SetRegionId(v string) *ListExecutionsRequest {
	s.RegionId = &v
	return s
}

func (s *ListExecutionsRequest) SetResourceGroupId(v string) *ListExecutionsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListExecutionsRequest) SetResourceId(v string) *ListExecutionsRequest {
	s.ResourceId = &v
	return s
}

func (s *ListExecutionsRequest) SetResourceTemplateName(v string) *ListExecutionsRequest {
	s.ResourceTemplateName = &v
	return s
}

func (s *ListExecutionsRequest) SetSortField(v string) *ListExecutionsRequest {
	s.SortField = &v
	return s
}

func (s *ListExecutionsRequest) SetSortOrder(v string) *ListExecutionsRequest {
	s.SortOrder = &v
	return s
}

func (s *ListExecutionsRequest) SetStartDateAfter(v string) *ListExecutionsRequest {
	s.StartDateAfter = &v
	return s
}

func (s *ListExecutionsRequest) SetStartDateBefore(v string) *ListExecutionsRequest {
	s.StartDateBefore = &v
	return s
}

func (s *ListExecutionsRequest) SetStatus(v string) *ListExecutionsRequest {
	s.Status = &v
	return s
}

func (s *ListExecutionsRequest) SetTags(v map[string]interface{}) *ListExecutionsRequest {
	s.Tags = v
	return s
}

func (s *ListExecutionsRequest) SetTemplateName(v string) *ListExecutionsRequest {
	s.TemplateName = &v
	return s
}

type ListExecutionsShrinkRequest struct {
	// The types of the execution template. Valid values: Other, TimerTrigger, EventTrigger, and AlarmTrigger. You can specify only one of the Categories and Category parameters. We recommend that you specify Categories.
	//
	// example:
	//
	// ["TimerTrigger"、"EventTrigger"]
	Categories *string `json:"Categories,omitempty" xml:"Categories,omitempty"`
	// The type of the execution template. Valid values: Other, TimerTrigger, EventTrigger, and AlarmTrigger.
	//
	// example:
	//
	// Other
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The depth of execution. Valid values: RootDepth and FirstChildDepth. If you set this parameter to RootDepth, only the parent execution is returned. If you set this parameter to FirstChildDepth, only the child executions at the first level are returned. You can specify only one of the Depth and IncludeChildExecution parameters. We recommend that you specify Depth.
	//
	// example:
	//
	// RootDepth
	Depth *string `json:"Depth,omitempty" xml:"Depth,omitempty"`
	// The description of the execution.
	//
	// example:
	//
	// MyDescription
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The earliest end time. The executions that stop running at or later than the specified time are queried.
	//
	// example:
	//
	// 2019-05-16T10:26:14Z
	EndDateAfter *string `json:"EndDateAfter,omitempty" xml:"EndDateAfter,omitempty"`
	// The latest end time. The executions that stop running at or earlier than the specified time are queried.
	//
	// example:
	//
	// 2019-05-16T10:26:14Z
	EndDateBefore *string `json:"EndDateBefore,omitempty" xml:"EndDateBefore,omitempty"`
	// The executor.
	//
	// example:
	//
	// vme
	ExecutedBy *string `json:"ExecutedBy,omitempty" xml:"ExecutedBy,omitempty"`
	// The ID of the execution.
	//
	// example:
	//
	// exec-xxx
	ExecutionId *string `json:"ExecutionId,omitempty" xml:"ExecutionId,omitempty"`
	// Specifies whether to include child executions. Default value: False.
	//
	// example:
	//
	// true
	IncludeChildExecution *bool `json:"IncludeChildExecution,omitempty" xml:"IncludeChildExecution,omitempty"`
	// The number of entries to return on each page. Valid values: 10 to 100. Default value: 50.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The execution mode. Valid values:
	//
	// 	- **Automatic**
	//
	// 	- **Debug**
	//
	// example:
	//
	// Automatic
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// The token that is used to retrieve the next page of results.
	//
	// example:
	//
	// MTRBMDc0NjAtRUJFNy00N0NBLTk3NTctMTJDQzQ
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the parent execution.
	//
	// example:
	//
	// exec-xxx
	ParentExecutionId *string `json:"ParentExecutionId,omitempty" xml:"ParentExecutionId,omitempty"`
	// The RAM role.
	//
	// example:
	//
	// OOSServiceRole
	RamRole *string `json:"RamRole,omitempty" xml:"RamRole,omitempty"`
	// The ID of the region.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the instances you want to query belong.
	//
	// example:
	//
	// rg-acfmxsn4m4******
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The ID of the Elastic Compute Service (ECS) resource.
	//
	// example:
	//
	// i-xxx
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The name of the resource template.
	//
	// example:
	//
	// ACS-ECS-TEST
	ResourceTemplateName *string `json:"ResourceTemplateName,omitempty" xml:"ResourceTemplateName,omitempty"`
	// The field that is used to sort the executions to query. Valid values:
	//
	// 	- **StartDate**: specifies that the executions are sorted based on the time when they are created. This is the default value.
	//
	// 	- **EndDate**: specifies that the executions are sorted based on the time when they stop running.
	//
	// 	- **Status**: specifies that the executions are sorted based on their states.
	//
	// example:
	//
	// StartDate
	SortField *string `json:"SortField,omitempty" xml:"SortField,omitempty"`
	// The order in which you want to sort the results. Valid values:
	//
	// 	- **Ascending**: ascending order.
	//
	// 	- **Descending**: descending order. This is the default value.
	//
	// example:
	//
	// Ascending
	SortOrder *string `json:"SortOrder,omitempty" xml:"SortOrder,omitempty"`
	// The earliest start time. The executions that start to run at or later than the specified time are queried.
	//
	// example:
	//
	// 2019-05-16T10:26:14Z
	StartDateAfter *string `json:"StartDateAfter,omitempty" xml:"StartDateAfter,omitempty"`
	// The latest start time. The executions that start to run at or earlier than the specified point in time are queried.
	//
	// example:
	//
	// 2019-05-16T10:26:14Z
	StartDateBefore *string `json:"StartDateBefore,omitempty" xml:"StartDateBefore,omitempty"`
	// The status of the execution. Valid values: Running, Started, Success, Failed, Waiting, Cancelled, Pending, and Skipped.
	//
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tags for the execution.
	//
	// example:
	//
	// {"k1":"v2","k2":"v2"}
	TagsShrink *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The name of the template. All templates whose names contain the specified template name are queried.
	//
	// example:
	//
	// MyTemplate
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
}

func (s ListExecutionsShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListExecutionsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListExecutionsShrinkRequest) SetCategories(v string) *ListExecutionsShrinkRequest {
	s.Categories = &v
	return s
}

func (s *ListExecutionsShrinkRequest) SetCategory(v string) *ListExecutionsShrinkRequest {
	s.Category = &v
	return s
}

func (s *ListExecutionsShrinkRequest) SetDepth(v string) *ListExecutionsShrinkRequest {
	s.Depth = &v
	return s
}

func (s *ListExecutionsShrinkRequest) SetDescription(v string) *ListExecutionsShrinkRequest {
	s.Description = &v
	return s
}

func (s *ListExecutionsShrinkRequest) SetEndDateAfter(v string) *ListExecutionsShrinkRequest {
	s.EndDateAfter = &v
	return s
}

func (s *ListExecutionsShrinkRequest) SetEndDateBefore(v string) *ListExecutionsShrinkRequest {
	s.EndDateBefore = &v
	return s
}

func (s *ListExecutionsShrinkRequest) SetExecutedBy(v string) *ListExecutionsShrinkRequest {
	s.ExecutedBy = &v
	return s
}

func (s *ListExecutionsShrinkRequest) SetExecutionId(v string) *ListExecutionsShrinkRequest {
	s.ExecutionId = &v
	return s
}

func (s *ListExecutionsShrinkRequest) SetIncludeChildExecution(v bool) *ListExecutionsShrinkRequest {
	s.IncludeChildExecution = &v
	return s
}

func (s *ListExecutionsShrinkRequest) SetMaxResults(v int32) *ListExecutionsShrinkRequest {
	s.MaxResults = &v
	return s
}

func (s *ListExecutionsShrinkRequest) SetMode(v string) *ListExecutionsShrinkRequest {
	s.Mode = &v
	return s
}

func (s *ListExecutionsShrinkRequest) SetNextToken(v string) *ListExecutionsShrinkRequest {
	s.NextToken = &v
	return s
}

func (s *ListExecutionsShrinkRequest) SetParentExecutionId(v string) *ListExecutionsShrinkRequest {
	s.ParentExecutionId = &v
	return s
}

func (s *ListExecutionsShrinkRequest) SetRamRole(v string) *ListExecutionsShrinkRequest {
	s.RamRole = &v
	return s
}

func (s *ListExecutionsShrinkRequest) SetRegionId(v string) *ListExecutionsShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *ListExecutionsShrinkRequest) SetResourceGroupId(v string) *ListExecutionsShrinkRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListExecutionsShrinkRequest) SetResourceId(v string) *ListExecutionsShrinkRequest {
	s.ResourceId = &v
	return s
}

func (s *ListExecutionsShrinkRequest) SetResourceTemplateName(v string) *ListExecutionsShrinkRequest {
	s.ResourceTemplateName = &v
	return s
}

func (s *ListExecutionsShrinkRequest) SetSortField(v string) *ListExecutionsShrinkRequest {
	s.SortField = &v
	return s
}

func (s *ListExecutionsShrinkRequest) SetSortOrder(v string) *ListExecutionsShrinkRequest {
	s.SortOrder = &v
	return s
}

func (s *ListExecutionsShrinkRequest) SetStartDateAfter(v string) *ListExecutionsShrinkRequest {
	s.StartDateAfter = &v
	return s
}

func (s *ListExecutionsShrinkRequest) SetStartDateBefore(v string) *ListExecutionsShrinkRequest {
	s.StartDateBefore = &v
	return s
}

func (s *ListExecutionsShrinkRequest) SetStatus(v string) *ListExecutionsShrinkRequest {
	s.Status = &v
	return s
}

func (s *ListExecutionsShrinkRequest) SetTagsShrink(v string) *ListExecutionsShrinkRequest {
	s.TagsShrink = &v
	return s
}

func (s *ListExecutionsShrinkRequest) SetTemplateName(v string) *ListExecutionsShrinkRequest {
	s.TemplateName = &v
	return s
}

type ListExecutionsResponseBody struct {
	// The details of the task executions.
	Executions []*ListExecutionsResponseBodyExecutions `json:"Executions,omitempty" xml:"Executions,omitempty" type:"Repeated"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that is used to retrieve the next page of results.
	//
	// example:
	//
	// MTRBMDc0NjAtRUJFNy00N0NBLTk3NTctMTJDQzQ
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 14A074-47CA-9757-12CC4761D47A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListExecutionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListExecutionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListExecutionsResponseBody) SetExecutions(v []*ListExecutionsResponseBodyExecutions) *ListExecutionsResponseBody {
	s.Executions = v
	return s
}

func (s *ListExecutionsResponseBody) SetMaxResults(v int32) *ListExecutionsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListExecutionsResponseBody) SetNextToken(v string) *ListExecutionsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListExecutionsResponseBody) SetRequestId(v string) *ListExecutionsResponseBody {
	s.RequestId = &v
	return s
}

type ListExecutionsResponseBodyExecutions struct {
	// The type of the execution template. Valid values: Other, TimerTrigger, EventTrigger, and AlarmTrigger.
	//
	// example:
	//
	// Other
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The number of tasks that are counted by execution status.
	//
	// example:
	//
	// {"Failed": 0,"Success": 1,"Total": 2}
	Counters map[string]interface{} `json:"Counters,omitempty" xml:"Counters,omitempty"`
	// The time when the execution was created.
	//
	// example:
	//
	// 2019-05-16T10:26:14Z
	CreateDate *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	// The information about the tasks that are running.
	CurrentTasks []*ListExecutionsResponseBodyExecutionsCurrentTasks `json:"CurrentTasks,omitempty" xml:"CurrentTasks,omitempty" type:"Repeated"`
	// The description of the execution.
	//
	// example:
	//
	// test execution.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The time when the execution stops running.
	//
	// example:
	//
	// 2019-05-16T10:26:14Z
	EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// The account ID of the user who started the execution of the template.
	//
	// example:
	//
	// 1309252800
	ExecutedBy *string `json:"ExecutedBy,omitempty" xml:"ExecutedBy,omitempty"`
	// The unique ID of the execution.
	//
	// example:
	//
	// exec-44d32b45d2a449e
	ExecutionId *string `json:"ExecutionId,omitempty" xml:"ExecutionId,omitempty"`
	// Indicates whether the execution contains child executions.
	//
	// example:
	//
	// false
	IsParent *bool `json:"IsParent,omitempty" xml:"IsParent,omitempty"`
	// The time when the template was last successfully triggered.
	//
	// example:
	//
	// 2019-05-27T09:29:18Z
	LastSuccessfulTriggerTime *string `json:"LastSuccessfulTriggerTime,omitempty" xml:"LastSuccessfulTriggerTime,omitempty"`
	// The outputs of last trigger.
	//
	// example:
	//
	// {
	//
	//       "InstanceId": "i-xxx"
	//
	// }
	LastTriggerOutputs *string `json:"LastTriggerOutputs,omitempty" xml:"LastTriggerOutputs,omitempty"`
	// The status of the execution after the template was last triggered.
	//
	// example:
	//
	// Success
	LastTriggerStatus *string `json:"LastTriggerStatus,omitempty" xml:"LastTriggerStatus,omitempty"`
	// The status message of last trigger.
	//
	// example:
	//
	// ""
	LastTriggerStatusMessage *string `json:"LastTriggerStatusMessage,omitempty" xml:"LastTriggerStatusMessage,omitempty"`
	// The time when the template was last successfully triggered.
	//
	// example:
	//
	// 2019-05-27T09:29:18Z
	LastTriggerTime *string `json:"LastTriggerTime,omitempty" xml:"LastTriggerTime,omitempty"`
	// The execution mode.
	//
	// example:
	//
	// Automatic
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// The output of the execution.
	//
	// example:
	//
	// { "InstanceId":"i-xxx" }
	Outputs *string `json:"Outputs,omitempty" xml:"Outputs,omitempty"`
	// The input parameters of the execution.
	//
	// example:
	//
	// { "Status":"Running" }
	Parameters map[string]interface{} `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// The ID of the parent execution.
	//
	// example:
	//
	// exec-xxx
	ParentExecutionId *string `json:"ParentExecutionId,omitempty" xml:"ParentExecutionId,omitempty"`
	// The role that started the execution of the template.
	//
	// example:
	//
	// OOSServiceRole
	RamRole *string `json:"RamRole,omitempty" xml:"RamRole,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfmxsn4m4******
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The status of the resource.
	//
	// example:
	//
	// { 			"Success": 1 		}
	ResourceStatus *string `json:"ResourceStatus,omitempty" xml:"ResourceStatus,omitempty"`
	// The security check mode. Valid values: Skip, and ConfirmEveryHighRiskAction.
	//
	// example:
	//
	// Skip
	SafetyCheck *string `json:"SafetyCheck,omitempty" xml:"SafetyCheck,omitempty"`
	// The time when the execution was started.
	//
	// example:
	//
	// 2019-05-16T10:26:14Z
	StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	// The status of the execution. Valid values: Started, Queued, Running, Waiting, Success, Failed, and Cancelled.
	//
	// example:
	//
	// Success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The status of the task execution.
	//
	// example:
	//
	// “”
	StatusMessage *string `json:"StatusMessage,omitempty" xml:"StatusMessage,omitempty"`
	// The reason for which the status occurs.
	//
	// example:
	//
	// ""
	StatusReason *string `json:"StatusReason,omitempty" xml:"StatusReason,omitempty"`
	// The tags of the execution.
	//
	// example:
	//
	// {}
	Tags map[string]interface{} `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The target resource.
	//
	// example:
	//
	// "{"ResourceType": "ALIYUN::ECS::Instance", "Filters": [{"ResourceIds": ["i-bp14z07dg3464980x72o"], "RegionId": "cn-hangzhou", "Type": "ResourceIds"}]}"
	Targets *string `json:"Targets,omitempty" xml:"Targets,omitempty"`
	// The ID of the template.
	//
	// example:
	//
	// 123
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The name of the template.
	//
	// example:
	//
	// MyTemplate
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// The version number of the template.
	//
	// example:
	//
	// v1
	TemplateVersion *string `json:"TemplateVersion,omitempty" xml:"TemplateVersion,omitempty"`
	// The time when the execution was updated.
	//
	// example:
	//
	// 2019-05-16T10:26:14Z
	UpdateDate *string `json:"UpdateDate,omitempty" xml:"UpdateDate,omitempty"`
	// The Waiting state.
	//
	// example:
	//
	// ""
	WaitingStatus *string `json:"WaitingStatus,omitempty" xml:"WaitingStatus,omitempty"`
}

func (s ListExecutionsResponseBodyExecutions) String() string {
	return tea.Prettify(s)
}

func (s ListExecutionsResponseBodyExecutions) GoString() string {
	return s.String()
}

func (s *ListExecutionsResponseBodyExecutions) SetCategory(v string) *ListExecutionsResponseBodyExecutions {
	s.Category = &v
	return s
}

func (s *ListExecutionsResponseBodyExecutions) SetCounters(v map[string]interface{}) *ListExecutionsResponseBodyExecutions {
	s.Counters = v
	return s
}

func (s *ListExecutionsResponseBodyExecutions) SetCreateDate(v string) *ListExecutionsResponseBodyExecutions {
	s.CreateDate = &v
	return s
}

func (s *ListExecutionsResponseBodyExecutions) SetCurrentTasks(v []*ListExecutionsResponseBodyExecutionsCurrentTasks) *ListExecutionsResponseBodyExecutions {
	s.CurrentTasks = v
	return s
}

func (s *ListExecutionsResponseBodyExecutions) SetDescription(v string) *ListExecutionsResponseBodyExecutions {
	s.Description = &v
	return s
}

func (s *ListExecutionsResponseBodyExecutions) SetEndDate(v string) *ListExecutionsResponseBodyExecutions {
	s.EndDate = &v
	return s
}

func (s *ListExecutionsResponseBodyExecutions) SetExecutedBy(v string) *ListExecutionsResponseBodyExecutions {
	s.ExecutedBy = &v
	return s
}

func (s *ListExecutionsResponseBodyExecutions) SetExecutionId(v string) *ListExecutionsResponseBodyExecutions {
	s.ExecutionId = &v
	return s
}

func (s *ListExecutionsResponseBodyExecutions) SetIsParent(v bool) *ListExecutionsResponseBodyExecutions {
	s.IsParent = &v
	return s
}

func (s *ListExecutionsResponseBodyExecutions) SetLastSuccessfulTriggerTime(v string) *ListExecutionsResponseBodyExecutions {
	s.LastSuccessfulTriggerTime = &v
	return s
}

func (s *ListExecutionsResponseBodyExecutions) SetLastTriggerOutputs(v string) *ListExecutionsResponseBodyExecutions {
	s.LastTriggerOutputs = &v
	return s
}

func (s *ListExecutionsResponseBodyExecutions) SetLastTriggerStatus(v string) *ListExecutionsResponseBodyExecutions {
	s.LastTriggerStatus = &v
	return s
}

func (s *ListExecutionsResponseBodyExecutions) SetLastTriggerStatusMessage(v string) *ListExecutionsResponseBodyExecutions {
	s.LastTriggerStatusMessage = &v
	return s
}

func (s *ListExecutionsResponseBodyExecutions) SetLastTriggerTime(v string) *ListExecutionsResponseBodyExecutions {
	s.LastTriggerTime = &v
	return s
}

func (s *ListExecutionsResponseBodyExecutions) SetMode(v string) *ListExecutionsResponseBodyExecutions {
	s.Mode = &v
	return s
}

func (s *ListExecutionsResponseBodyExecutions) SetOutputs(v string) *ListExecutionsResponseBodyExecutions {
	s.Outputs = &v
	return s
}

func (s *ListExecutionsResponseBodyExecutions) SetParameters(v map[string]interface{}) *ListExecutionsResponseBodyExecutions {
	s.Parameters = v
	return s
}

func (s *ListExecutionsResponseBodyExecutions) SetParentExecutionId(v string) *ListExecutionsResponseBodyExecutions {
	s.ParentExecutionId = &v
	return s
}

func (s *ListExecutionsResponseBodyExecutions) SetRamRole(v string) *ListExecutionsResponseBodyExecutions {
	s.RamRole = &v
	return s
}

func (s *ListExecutionsResponseBodyExecutions) SetResourceGroupId(v string) *ListExecutionsResponseBodyExecutions {
	s.ResourceGroupId = &v
	return s
}

func (s *ListExecutionsResponseBodyExecutions) SetResourceStatus(v string) *ListExecutionsResponseBodyExecutions {
	s.ResourceStatus = &v
	return s
}

func (s *ListExecutionsResponseBodyExecutions) SetSafetyCheck(v string) *ListExecutionsResponseBodyExecutions {
	s.SafetyCheck = &v
	return s
}

func (s *ListExecutionsResponseBodyExecutions) SetStartDate(v string) *ListExecutionsResponseBodyExecutions {
	s.StartDate = &v
	return s
}

func (s *ListExecutionsResponseBodyExecutions) SetStatus(v string) *ListExecutionsResponseBodyExecutions {
	s.Status = &v
	return s
}

func (s *ListExecutionsResponseBodyExecutions) SetStatusMessage(v string) *ListExecutionsResponseBodyExecutions {
	s.StatusMessage = &v
	return s
}

func (s *ListExecutionsResponseBodyExecutions) SetStatusReason(v string) *ListExecutionsResponseBodyExecutions {
	s.StatusReason = &v
	return s
}

func (s *ListExecutionsResponseBodyExecutions) SetTags(v map[string]interface{}) *ListExecutionsResponseBodyExecutions {
	s.Tags = v
	return s
}

func (s *ListExecutionsResponseBodyExecutions) SetTargets(v string) *ListExecutionsResponseBodyExecutions {
	s.Targets = &v
	return s
}

func (s *ListExecutionsResponseBodyExecutions) SetTemplateId(v string) *ListExecutionsResponseBodyExecutions {
	s.TemplateId = &v
	return s
}

func (s *ListExecutionsResponseBodyExecutions) SetTemplateName(v string) *ListExecutionsResponseBodyExecutions {
	s.TemplateName = &v
	return s
}

func (s *ListExecutionsResponseBodyExecutions) SetTemplateVersion(v string) *ListExecutionsResponseBodyExecutions {
	s.TemplateVersion = &v
	return s
}

func (s *ListExecutionsResponseBodyExecutions) SetUpdateDate(v string) *ListExecutionsResponseBodyExecutions {
	s.UpdateDate = &v
	return s
}

func (s *ListExecutionsResponseBodyExecutions) SetWaitingStatus(v string) *ListExecutionsResponseBodyExecutions {
	s.WaitingStatus = &v
	return s
}

type ListExecutionsResponseBodyExecutionsCurrentTasks struct {
	// The execution template of the task.
	//
	// example:
	//
	// acs::Template
	TaskAction *string `json:"TaskAction,omitempty" xml:"TaskAction,omitempty"`
	// The ID of the task execution.
	//
	// example:
	//
	// task-exec-44d32b45d2a49899#1
	TaskExecutionId *string `json:"TaskExecutionId,omitempty" xml:"TaskExecutionId,omitempty"`
	// The name of the task.
	//
	// example:
	//
	// installSLSILogtail
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
}

func (s ListExecutionsResponseBodyExecutionsCurrentTasks) String() string {
	return tea.Prettify(s)
}

func (s ListExecutionsResponseBodyExecutionsCurrentTasks) GoString() string {
	return s.String()
}

func (s *ListExecutionsResponseBodyExecutionsCurrentTasks) SetTaskAction(v string) *ListExecutionsResponseBodyExecutionsCurrentTasks {
	s.TaskAction = &v
	return s
}

func (s *ListExecutionsResponseBodyExecutionsCurrentTasks) SetTaskExecutionId(v string) *ListExecutionsResponseBodyExecutionsCurrentTasks {
	s.TaskExecutionId = &v
	return s
}

func (s *ListExecutionsResponseBodyExecutionsCurrentTasks) SetTaskName(v string) *ListExecutionsResponseBodyExecutionsCurrentTasks {
	s.TaskName = &v
	return s
}

type ListExecutionsResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListExecutionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListExecutionsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListExecutionsResponse) GoString() string {
	return s.String()
}

func (s *ListExecutionsResponse) SetHeaders(v map[string]*string) *ListExecutionsResponse {
	s.Headers = v
	return s
}

func (s *ListExecutionsResponse) SetStatusCode(v int32) *ListExecutionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListExecutionsResponse) SetBody(v *ListExecutionsResponseBody) *ListExecutionsResponse {
	s.Body = v
	return s
}

type ListGitRepositoriesRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	OrgName     *string `json:"OrgName,omitempty" xml:"OrgName,omitempty"`
	// This parameter is required.
	Owner      *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	Platform *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListGitRepositoriesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListGitRepositoriesRequest) GoString() string {
	return s.String()
}

func (s *ListGitRepositoriesRequest) SetClientToken(v string) *ListGitRepositoriesRequest {
	s.ClientToken = &v
	return s
}

func (s *ListGitRepositoriesRequest) SetOrgName(v string) *ListGitRepositoriesRequest {
	s.OrgName = &v
	return s
}

func (s *ListGitRepositoriesRequest) SetOwner(v string) *ListGitRepositoriesRequest {
	s.Owner = &v
	return s
}

func (s *ListGitRepositoriesRequest) SetPageNumber(v int32) *ListGitRepositoriesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListGitRepositoriesRequest) SetPageSize(v int32) *ListGitRepositoriesRequest {
	s.PageSize = &v
	return s
}

func (s *ListGitRepositoriesRequest) SetPlatform(v string) *ListGitRepositoriesRequest {
	s.Platform = &v
	return s
}

func (s *ListGitRepositoriesRequest) SetRegionId(v string) *ListGitRepositoriesRequest {
	s.RegionId = &v
	return s
}

type ListGitRepositoriesResponseBody struct {
	Count     *int32                                     `json:"Count,omitempty" xml:"Count,omitempty"`
	GitRepos  []*ListGitRepositoriesResponseBodyGitRepos `json:"GitRepos,omitempty" xml:"GitRepos,omitempty" type:"Repeated"`
	RequestId *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListGitRepositoriesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListGitRepositoriesResponseBody) GoString() string {
	return s.String()
}

func (s *ListGitRepositoriesResponseBody) SetCount(v int32) *ListGitRepositoriesResponseBody {
	s.Count = &v
	return s
}

func (s *ListGitRepositoriesResponseBody) SetGitRepos(v []*ListGitRepositoriesResponseBodyGitRepos) *ListGitRepositoriesResponseBody {
	s.GitRepos = v
	return s
}

func (s *ListGitRepositoriesResponseBody) SetRequestId(v string) *ListGitRepositoriesResponseBody {
	s.RequestId = &v
	return s
}

type ListGitRepositoriesResponseBodyGitRepos struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	FullName    *string `json:"FullName,omitempty" xml:"FullName,omitempty"`
	HtmlUrl     *string `json:"HtmlUrl,omitempty" xml:"HtmlUrl,omitempty"`
	IsPrivate   *bool   `json:"IsPrivate,omitempty" xml:"IsPrivate,omitempty"`
}

func (s ListGitRepositoriesResponseBodyGitRepos) String() string {
	return tea.Prettify(s)
}

func (s ListGitRepositoriesResponseBodyGitRepos) GoString() string {
	return s.String()
}

func (s *ListGitRepositoriesResponseBodyGitRepos) SetDescription(v string) *ListGitRepositoriesResponseBodyGitRepos {
	s.Description = &v
	return s
}

func (s *ListGitRepositoriesResponseBodyGitRepos) SetFullName(v string) *ListGitRepositoriesResponseBodyGitRepos {
	s.FullName = &v
	return s
}

func (s *ListGitRepositoriesResponseBodyGitRepos) SetHtmlUrl(v string) *ListGitRepositoriesResponseBodyGitRepos {
	s.HtmlUrl = &v
	return s
}

func (s *ListGitRepositoriesResponseBodyGitRepos) SetIsPrivate(v bool) *ListGitRepositoriesResponseBodyGitRepos {
	s.IsPrivate = &v
	return s
}

type ListGitRepositoriesResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListGitRepositoriesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListGitRepositoriesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListGitRepositoriesResponse) GoString() string {
	return s.String()
}

func (s *ListGitRepositoriesResponse) SetHeaders(v map[string]*string) *ListGitRepositoriesResponse {
	s.Headers = v
	return s
}

func (s *ListGitRepositoriesResponse) SetStatusCode(v int32) *ListGitRepositoriesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListGitRepositoriesResponse) SetBody(v *ListGitRepositoriesResponseBody) *ListGitRepositoriesResponse {
	s.Body = v
	return s
}

type ListInstancePackageStatesRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// i-bp1cpoxxxwxxxxxxxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// MTRBMDc0NjAtRUJFNy00N0NBLTk3NTctzxxxxxxx
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// ["template1","template2"]
	TemplateNames *string `json:"TemplateNames,omitempty" xml:"TemplateNames,omitempty"`
}

func (s ListInstancePackageStatesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListInstancePackageStatesRequest) GoString() string {
	return s.String()
}

func (s *ListInstancePackageStatesRequest) SetInstanceId(v string) *ListInstancePackageStatesRequest {
	s.InstanceId = &v
	return s
}

func (s *ListInstancePackageStatesRequest) SetMaxResults(v int32) *ListInstancePackageStatesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListInstancePackageStatesRequest) SetNextToken(v string) *ListInstancePackageStatesRequest {
	s.NextToken = &v
	return s
}

func (s *ListInstancePackageStatesRequest) SetRegionId(v string) *ListInstancePackageStatesRequest {
	s.RegionId = &v
	return s
}

func (s *ListInstancePackageStatesRequest) SetTemplateNames(v string) *ListInstancePackageStatesRequest {
	s.TemplateNames = &v
	return s
}

type ListInstancePackageStatesResponseBody struct {
	// example:
	//
	// 50
	MaxResults *string `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// MTRBMDc0NjAtRUJFNy00N0NBLTk3NTctzxxxxxxx
	NextToken     *string                                               `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	PackageStates []*ListInstancePackageStatesResponseBodyPackageStates `json:"PackageStates,omitempty" xml:"PackageStates,omitempty" type:"Repeated"`
	// Id of the request
	//
	// example:
	//
	// 1306108F-610C-40FD-AAD5-XXXXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListInstancePackageStatesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListInstancePackageStatesResponseBody) GoString() string {
	return s.String()
}

func (s *ListInstancePackageStatesResponseBody) SetMaxResults(v string) *ListInstancePackageStatesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListInstancePackageStatesResponseBody) SetNextToken(v string) *ListInstancePackageStatesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListInstancePackageStatesResponseBody) SetPackageStates(v []*ListInstancePackageStatesResponseBodyPackageStates) *ListInstancePackageStatesResponseBody {
	s.PackageStates = v
	return s
}

func (s *ListInstancePackageStatesResponseBody) SetRequestId(v string) *ListInstancePackageStatesResponseBody {
	s.RequestId = &v
	return s
}

type ListInstancePackageStatesResponseBodyPackageStates struct {
	// example:
	//
	// template description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// {}
	Parameters *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// example:
	//
	// Alibaba Cloud
	Publisher *string `json:"Publisher,omitempty" xml:"Publisher,omitempty"`
	// example:
	//
	// Package
	TemplateCategory *string `json:"TemplateCategory,omitempty" xml:"TemplateCategory,omitempty"`
	// example:
	//
	// 087b1e11072a40259f6fxxxxxxxxx
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// example:
	//
	// ACS-ECS-Docker
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// example:
	//
	// v3
	TemplateVersion *string `json:"TemplateVersion,omitempty" xml:"TemplateVersion,omitempty"`
	// example:
	//
	// fix bug
	TemplateVersionName *string `json:"TemplateVersionName,omitempty" xml:"TemplateVersionName,omitempty"`
	// example:
	//
	// 2024-05-04T11:17:28
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListInstancePackageStatesResponseBodyPackageStates) String() string {
	return tea.Prettify(s)
}

func (s ListInstancePackageStatesResponseBodyPackageStates) GoString() string {
	return s.String()
}

func (s *ListInstancePackageStatesResponseBodyPackageStates) SetDescription(v string) *ListInstancePackageStatesResponseBodyPackageStates {
	s.Description = &v
	return s
}

func (s *ListInstancePackageStatesResponseBodyPackageStates) SetParameters(v string) *ListInstancePackageStatesResponseBodyPackageStates {
	s.Parameters = &v
	return s
}

func (s *ListInstancePackageStatesResponseBodyPackageStates) SetPublisher(v string) *ListInstancePackageStatesResponseBodyPackageStates {
	s.Publisher = &v
	return s
}

func (s *ListInstancePackageStatesResponseBodyPackageStates) SetTemplateCategory(v string) *ListInstancePackageStatesResponseBodyPackageStates {
	s.TemplateCategory = &v
	return s
}

func (s *ListInstancePackageStatesResponseBodyPackageStates) SetTemplateId(v string) *ListInstancePackageStatesResponseBodyPackageStates {
	s.TemplateId = &v
	return s
}

func (s *ListInstancePackageStatesResponseBodyPackageStates) SetTemplateName(v string) *ListInstancePackageStatesResponseBodyPackageStates {
	s.TemplateName = &v
	return s
}

func (s *ListInstancePackageStatesResponseBodyPackageStates) SetTemplateVersion(v string) *ListInstancePackageStatesResponseBodyPackageStates {
	s.TemplateVersion = &v
	return s
}

func (s *ListInstancePackageStatesResponseBodyPackageStates) SetTemplateVersionName(v string) *ListInstancePackageStatesResponseBodyPackageStates {
	s.TemplateVersionName = &v
	return s
}

func (s *ListInstancePackageStatesResponseBodyPackageStates) SetUpdateTime(v string) *ListInstancePackageStatesResponseBodyPackageStates {
	s.UpdateTime = &v
	return s
}

type ListInstancePackageStatesResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListInstancePackageStatesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListInstancePackageStatesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListInstancePackageStatesResponse) GoString() string {
	return s.String()
}

func (s *ListInstancePackageStatesResponse) SetHeaders(v map[string]*string) *ListInstancePackageStatesResponse {
	s.Headers = v
	return s
}

func (s *ListInstancePackageStatesResponse) SetStatusCode(v int32) *ListInstancePackageStatesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListInstancePackageStatesResponse) SetBody(v *ListInstancePackageStatesResponseBody) *ListInstancePackageStatesResponse {
	s.Body = v
	return s
}

type ListInstancePatchStatesRequest struct {
	// The ID of the Elastic Compute Service (ECS) instance. The value can be a JSON array that consists of up to 100 instance IDs. Separate the instance IDs with commas (,).
	//
	// example:
	//
	// ["i-bp1jaxa2bs4bps7*****", "i-bp67acfmxazb4p****", … "i-bp67acfmxazb4p****"]
	InstanceIds *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
	// The number of entries to return on each page.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that is used to retrieve the next page of results.
	//
	// example:
	//
	// -
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the region in which the instance whose patches you want to query resides.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListInstancePatchStatesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListInstancePatchStatesRequest) GoString() string {
	return s.String()
}

func (s *ListInstancePatchStatesRequest) SetInstanceIds(v string) *ListInstancePatchStatesRequest {
	s.InstanceIds = &v
	return s
}

func (s *ListInstancePatchStatesRequest) SetMaxResults(v int32) *ListInstancePatchStatesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListInstancePatchStatesRequest) SetNextToken(v string) *ListInstancePatchStatesRequest {
	s.NextToken = &v
	return s
}

func (s *ListInstancePatchStatesRequest) SetRegionId(v string) *ListInstancePatchStatesRequest {
	s.RegionId = &v
	return s
}

type ListInstancePatchStatesResponseBody struct {
	// The details of patches of the instance.
	InstancePatchStates []*ListInstancePatchStatesResponseBodyInstancePatchStates `json:"InstancePatchStates,omitempty" xml:"InstancePatchStates,omitempty" type:"Repeated"`
	// The number of entries returned on each page.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that is used to retrieve the next page of results.
	//
	// example:
	//
	// -
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 9A47C086-E64D-52EE-8B2C-EFD23877C55E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListInstancePatchStatesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListInstancePatchStatesResponseBody) GoString() string {
	return s.String()
}

func (s *ListInstancePatchStatesResponseBody) SetInstancePatchStates(v []*ListInstancePatchStatesResponseBodyInstancePatchStates) *ListInstancePatchStatesResponseBody {
	s.InstancePatchStates = v
	return s
}

func (s *ListInstancePatchStatesResponseBody) SetMaxResults(v int32) *ListInstancePatchStatesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListInstancePatchStatesResponseBody) SetNextToken(v string) *ListInstancePatchStatesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListInstancePatchStatesResponseBody) SetRequestId(v string) *ListInstancePatchStatesResponseBody {
	s.RequestId = &v
	return s
}

type ListInstancePatchStatesResponseBodyInstancePatchStates struct {
	// The ID of the patch baseline.
	//
	// example:
	//
	// pb-f9393021b7a049e1b34e
	BaselineId *string `json:"BaselineId,omitempty" xml:"BaselineId,omitempty"`
	// The number of patches that failed to be installed.
	//
	// example:
	//
	// 0
	FailedCount *string `json:"FailedCount,omitempty" xml:"FailedCount,omitempty"`
	// The number of installed patches.
	//
	// example:
	//
	// 0
	InstalledCount *string `json:"InstalledCount,omitempty" xml:"InstalledCount,omitempty"`
	// The number of patches that do not meet the baseline.
	//
	// example:
	//
	// 0
	InstalledOtherCount *string `json:"InstalledOtherCount,omitempty" xml:"InstalledOtherCount,omitempty"`
	// The number of patches that have been installed but require a restart to take effect.
	//
	// example:
	//
	// 0
	InstalledPendingRebootCount *string `json:"InstalledPendingRebootCount,omitempty" xml:"InstalledPendingRebootCount,omitempty"`
	// The number of patches that are rejected by the user.
	//
	// example:
	//
	// 0
	InstalledRejectedCount *string `json:"InstalledRejectedCount,omitempty" xml:"InstalledRejectedCount,omitempty"`
	// The ID of the ECS instance.
	//
	// example:
	//
	// i-bp1jaxa2bs4bps7*****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The number of patches that are not installed.
	//
	// example:
	//
	// 0
	MissingCount *string `json:"MissingCount,omitempty" xml:"MissingCount,omitempty"`
	// The time when the operation ended.
	//
	// example:
	//
	// 2021-09-10T11:42:22Z
	OperationEndTime *string `json:"OperationEndTime,omitempty" xml:"OperationEndTime,omitempty"`
	// The time when the operation was initiated.
	//
	// example:
	//
	// 2021-09-10T11:42:22Z
	OperationStartTime *string `json:"OperationStartTime,omitempty" xml:"OperationStartTime,omitempty"`
	// The operation type.
	//
	// example:
	//
	// scan
	OperationType *string `json:"OperationType,omitempty" xml:"OperationType,omitempty"`
	// The information about the user.
	//
	// example:
	//
	// “”
	OwnerInformation *string `json:"OwnerInformation,omitempty" xml:"OwnerInformation,omitempty"`
	// The patch group.
	//
	// example:
	//
	// null
	PatchGroup *string `json:"PatchGroup,omitempty" xml:"PatchGroup,omitempty"`
}

func (s ListInstancePatchStatesResponseBodyInstancePatchStates) String() string {
	return tea.Prettify(s)
}

func (s ListInstancePatchStatesResponseBodyInstancePatchStates) GoString() string {
	return s.String()
}

func (s *ListInstancePatchStatesResponseBodyInstancePatchStates) SetBaselineId(v string) *ListInstancePatchStatesResponseBodyInstancePatchStates {
	s.BaselineId = &v
	return s
}

func (s *ListInstancePatchStatesResponseBodyInstancePatchStates) SetFailedCount(v string) *ListInstancePatchStatesResponseBodyInstancePatchStates {
	s.FailedCount = &v
	return s
}

func (s *ListInstancePatchStatesResponseBodyInstancePatchStates) SetInstalledCount(v string) *ListInstancePatchStatesResponseBodyInstancePatchStates {
	s.InstalledCount = &v
	return s
}

func (s *ListInstancePatchStatesResponseBodyInstancePatchStates) SetInstalledOtherCount(v string) *ListInstancePatchStatesResponseBodyInstancePatchStates {
	s.InstalledOtherCount = &v
	return s
}

func (s *ListInstancePatchStatesResponseBodyInstancePatchStates) SetInstalledPendingRebootCount(v string) *ListInstancePatchStatesResponseBodyInstancePatchStates {
	s.InstalledPendingRebootCount = &v
	return s
}

func (s *ListInstancePatchStatesResponseBodyInstancePatchStates) SetInstalledRejectedCount(v string) *ListInstancePatchStatesResponseBodyInstancePatchStates {
	s.InstalledRejectedCount = &v
	return s
}

func (s *ListInstancePatchStatesResponseBodyInstancePatchStates) SetInstanceId(v string) *ListInstancePatchStatesResponseBodyInstancePatchStates {
	s.InstanceId = &v
	return s
}

func (s *ListInstancePatchStatesResponseBodyInstancePatchStates) SetMissingCount(v string) *ListInstancePatchStatesResponseBodyInstancePatchStates {
	s.MissingCount = &v
	return s
}

func (s *ListInstancePatchStatesResponseBodyInstancePatchStates) SetOperationEndTime(v string) *ListInstancePatchStatesResponseBodyInstancePatchStates {
	s.OperationEndTime = &v
	return s
}

func (s *ListInstancePatchStatesResponseBodyInstancePatchStates) SetOperationStartTime(v string) *ListInstancePatchStatesResponseBodyInstancePatchStates {
	s.OperationStartTime = &v
	return s
}

func (s *ListInstancePatchStatesResponseBodyInstancePatchStates) SetOperationType(v string) *ListInstancePatchStatesResponseBodyInstancePatchStates {
	s.OperationType = &v
	return s
}

func (s *ListInstancePatchStatesResponseBodyInstancePatchStates) SetOwnerInformation(v string) *ListInstancePatchStatesResponseBodyInstancePatchStates {
	s.OwnerInformation = &v
	return s
}

func (s *ListInstancePatchStatesResponseBodyInstancePatchStates) SetPatchGroup(v string) *ListInstancePatchStatesResponseBodyInstancePatchStates {
	s.PatchGroup = &v
	return s
}

type ListInstancePatchStatesResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListInstancePatchStatesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListInstancePatchStatesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListInstancePatchStatesResponse) GoString() string {
	return s.String()
}

func (s *ListInstancePatchStatesResponse) SetHeaders(v map[string]*string) *ListInstancePatchStatesResponse {
	s.Headers = v
	return s
}

func (s *ListInstancePatchStatesResponse) SetStatusCode(v int32) *ListInstancePatchStatesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListInstancePatchStatesResponse) SetBody(v *ListInstancePatchStatesResponseBody) *ListInstancePatchStatesResponse {
	s.Body = v
	return s
}

type ListInstancePatchesRequest struct {
	// The ID of the instance.
	//
	// example:
	//
	// i-bp1jaxa2bs4bps7*****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The number of entries to return on each page.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that is used to retrieve the next page of results.
	//
	// example:
	//
	// gAAAAABfTgv5ewUWmNdJ3g7JVLvX70sPH90GZOVGC6KPDUL0FIIb
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The status of the patches that you want to query. If you do not set this parameter, patches are not filtered.
	//
	// example:
	//
	// Installed
	PatchStatuses *string `json:"PatchStatuses,omitempty" xml:"PatchStatuses,omitempty"`
	// The ID of the region in which the instance whose patches you want to query resides.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListInstancePatchesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListInstancePatchesRequest) GoString() string {
	return s.String()
}

func (s *ListInstancePatchesRequest) SetInstanceId(v string) *ListInstancePatchesRequest {
	s.InstanceId = &v
	return s
}

func (s *ListInstancePatchesRequest) SetMaxResults(v int32) *ListInstancePatchesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListInstancePatchesRequest) SetNextToken(v string) *ListInstancePatchesRequest {
	s.NextToken = &v
	return s
}

func (s *ListInstancePatchesRequest) SetPatchStatuses(v string) *ListInstancePatchesRequest {
	s.PatchStatuses = &v
	return s
}

func (s *ListInstancePatchesRequest) SetRegionId(v string) *ListInstancePatchesRequest {
	s.RegionId = &v
	return s
}

type ListInstancePatchesResponseBody struct {
	// The number of entries returned on each page.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that is used to retrieve the next page of results.
	//
	// example:
	//
	// -
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The information about the patch.
	Patches []*ListInstancePatchesResponseBodyPatches `json:"Patches,omitempty" xml:"Patches,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 0A615755-9C86-5EA6-BF9E-6E8F1AFF9403
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListInstancePatchesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListInstancePatchesResponseBody) GoString() string {
	return s.String()
}

func (s *ListInstancePatchesResponseBody) SetMaxResults(v int32) *ListInstancePatchesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListInstancePatchesResponseBody) SetNextToken(v string) *ListInstancePatchesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListInstancePatchesResponseBody) SetPatches(v []*ListInstancePatchesResponseBodyPatches) *ListInstancePatchesResponseBody {
	s.Patches = v
	return s
}

func (s *ListInstancePatchesResponseBody) SetRequestId(v string) *ListInstancePatchesResponseBody {
	s.RequestId = &v
	return s
}

type ListInstancePatchesResponseBodyPatches struct {
	// The classification of the patch.
	//
	// example:
	//
	// “”
	Classification *string `json:"Classification,omitempty" xml:"Classification,omitempty"`
	// The time when the patch was installed.
	//
	// example:
	//
	// 2021-01-28T07:07:20Z
	InstalledTime *string `json:"InstalledTime,omitempty" xml:"InstalledTime,omitempty"`
	// The Id of KBId.
	//
	// example:
	//
	// apt-utils.amd64
	KBId *string `json:"KBId,omitempty" xml:"KBId,omitempty"`
	// The level of the severity.
	//
	// example:
	//
	// important
	Severity *string `json:"Severity,omitempty" xml:"Severity,omitempty"`
	// The status of the installation.
	//
	// example:
	//
	// Installed
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The name of the patch.
	//
	// example:
	//
	// isc-dhcp-common.amd64:4.3.5-3ubuntu7.3
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s ListInstancePatchesResponseBodyPatches) String() string {
	return tea.Prettify(s)
}

func (s ListInstancePatchesResponseBodyPatches) GoString() string {
	return s.String()
}

func (s *ListInstancePatchesResponseBodyPatches) SetClassification(v string) *ListInstancePatchesResponseBodyPatches {
	s.Classification = &v
	return s
}

func (s *ListInstancePatchesResponseBodyPatches) SetInstalledTime(v string) *ListInstancePatchesResponseBodyPatches {
	s.InstalledTime = &v
	return s
}

func (s *ListInstancePatchesResponseBodyPatches) SetKBId(v string) *ListInstancePatchesResponseBodyPatches {
	s.KBId = &v
	return s
}

func (s *ListInstancePatchesResponseBodyPatches) SetSeverity(v string) *ListInstancePatchesResponseBodyPatches {
	s.Severity = &v
	return s
}

func (s *ListInstancePatchesResponseBodyPatches) SetStatus(v string) *ListInstancePatchesResponseBodyPatches {
	s.Status = &v
	return s
}

func (s *ListInstancePatchesResponseBodyPatches) SetTitle(v string) *ListInstancePatchesResponseBodyPatches {
	s.Title = &v
	return s
}

type ListInstancePatchesResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListInstancePatchesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListInstancePatchesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListInstancePatchesResponse) GoString() string {
	return s.String()
}

func (s *ListInstancePatchesResponse) SetHeaders(v map[string]*string) *ListInstancePatchesResponse {
	s.Headers = v
	return s
}

func (s *ListInstancePatchesResponse) SetStatusCode(v int32) *ListInstancePatchesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListInstancePatchesResponse) SetBody(v *ListInstancePatchesResponseBody) *ListInstancePatchesResponse {
	s.Body = v
	return s
}

type ListInventoryEntriesRequest struct {
	// The filter rules for the component.
	Filter []*ListInventoryEntriesRequestFilter `json:"Filter,omitempty" xml:"Filter,omitempty" type:"Repeated"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// i-bp1cpoxxxwxxxxxxxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The number of entries per page. Valid values: 1 to 100. Default value: 50.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. You do not need to specify this parameter for the first request.
	//
	// example:
	//
	// MTRBMDc0NjAtRUJFNy00N0NBLTk3NTctMTJDQzA
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the region in which the instance resides.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the component. Valid values:
	//
	// 	- ACS:InstanceInformation
	//
	// 	- ACS:Application
	//
	// 	- ACS:File
	//
	// 	- ACS:Network
	//
	// 	- ACS:WindowsRole
	//
	// 	- ACS:Service
	//
	// 	- ACS:WindowsRegistry
	//
	// 	- ACS:WindowsUpdate
	//
	// This parameter is required.
	//
	// example:
	//
	// ACS:InstanceInformation
	TypeName *string `json:"TypeName,omitempty" xml:"TypeName,omitempty"`
}

func (s ListInventoryEntriesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListInventoryEntriesRequest) GoString() string {
	return s.String()
}

func (s *ListInventoryEntriesRequest) SetFilter(v []*ListInventoryEntriesRequestFilter) *ListInventoryEntriesRequest {
	s.Filter = v
	return s
}

func (s *ListInventoryEntriesRequest) SetInstanceId(v string) *ListInventoryEntriesRequest {
	s.InstanceId = &v
	return s
}

func (s *ListInventoryEntriesRequest) SetMaxResults(v int32) *ListInventoryEntriesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListInventoryEntriesRequest) SetNextToken(v string) *ListInventoryEntriesRequest {
	s.NextToken = &v
	return s
}

func (s *ListInventoryEntriesRequest) SetRegionId(v string) *ListInventoryEntriesRequest {
	s.RegionId = &v
	return s
}

func (s *ListInventoryEntriesRequest) SetTypeName(v string) *ListInventoryEntriesRequest {
	s.TypeName = &v
	return s
}

type ListInventoryEntriesRequestFilter struct {
	// The name of the component property. Valid values of N: 1 to 5.
	//
	// example:
	//
	// PlatformName
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The comparison operator that is used to filter property values. Valid values of N: 1 to 5. Valid values:
	//
	// 	- Equal
	//
	// 	- NotEqual
	//
	// 	- BeginWith
	//
	// 	- LessThan
	//
	// 	- GreaterThan
	//
	// example:
	//
	// Equal
	Operator *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	// The values of properties. Valid values of the first N: 1 to 5. Valid values of the second N: 1 to 20.
	//
	// example:
	//
	// test
	Value []*string `json:"Value,omitempty" xml:"Value,omitempty" type:"Repeated"`
}

func (s ListInventoryEntriesRequestFilter) String() string {
	return tea.Prettify(s)
}

func (s ListInventoryEntriesRequestFilter) GoString() string {
	return s.String()
}

func (s *ListInventoryEntriesRequestFilter) SetName(v string) *ListInventoryEntriesRequestFilter {
	s.Name = &v
	return s
}

func (s *ListInventoryEntriesRequestFilter) SetOperator(v string) *ListInventoryEntriesRequestFilter {
	s.Operator = &v
	return s
}

func (s *ListInventoryEntriesRequestFilter) SetValue(v []*string) *ListInventoryEntriesRequestFilter {
	s.Value = v
	return s
}

type ListInventoryEntriesResponseBody struct {
	// The time when the request was sent.
	//
	// example:
	//
	// 2020-09-17T12:28:13Z
	CaptureTime *string `json:"CaptureTime,omitempty" xml:"CaptureTime,omitempty"`
	// The configurations of the component.
	Entries []map[string]interface{} `json:"Entries,omitempty" xml:"Entries,omitempty" type:"Repeated"`
	// The ID of the ECS instance.
	//
	// example:
	//
	// i-bp1cpoxxxwxxxxxxxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The returned value of NextToken is a pagination token, which can be used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// gAAAAABfTgv5ewUWmNdJ3g7JVLvX70sPH90GZOVGC
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// A81E4B2E-6B33-4BAE-9856-55DB7C893E01
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The version number of the component.
	//
	// example:
	//
	// 1.0
	SchemaVersion *string `json:"SchemaVersion,omitempty" xml:"SchemaVersion,omitempty"`
	// The name of the component.
	//
	// example:
	//
	// ACS:InstanceInformation
	TypeName *string `json:"TypeName,omitempty" xml:"TypeName,omitempty"`
}

func (s ListInventoryEntriesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListInventoryEntriesResponseBody) GoString() string {
	return s.String()
}

func (s *ListInventoryEntriesResponseBody) SetCaptureTime(v string) *ListInventoryEntriesResponseBody {
	s.CaptureTime = &v
	return s
}

func (s *ListInventoryEntriesResponseBody) SetEntries(v []map[string]interface{}) *ListInventoryEntriesResponseBody {
	s.Entries = v
	return s
}

func (s *ListInventoryEntriesResponseBody) SetInstanceId(v string) *ListInventoryEntriesResponseBody {
	s.InstanceId = &v
	return s
}

func (s *ListInventoryEntriesResponseBody) SetMaxResults(v int32) *ListInventoryEntriesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListInventoryEntriesResponseBody) SetNextToken(v string) *ListInventoryEntriesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListInventoryEntriesResponseBody) SetRequestId(v string) *ListInventoryEntriesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListInventoryEntriesResponseBody) SetSchemaVersion(v string) *ListInventoryEntriesResponseBody {
	s.SchemaVersion = &v
	return s
}

func (s *ListInventoryEntriesResponseBody) SetTypeName(v string) *ListInventoryEntriesResponseBody {
	s.TypeName = &v
	return s
}

type ListInventoryEntriesResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListInventoryEntriesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListInventoryEntriesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListInventoryEntriesResponse) GoString() string {
	return s.String()
}

func (s *ListInventoryEntriesResponse) SetHeaders(v map[string]*string) *ListInventoryEntriesResponse {
	s.Headers = v
	return s
}

func (s *ListInventoryEntriesResponse) SetStatusCode(v int32) *ListInventoryEntriesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListInventoryEntriesResponse) SetBody(v *ListInventoryEntriesResponseBody) *ListInventoryEntriesResponse {
	s.Body = v
	return s
}

type ListOpsItemsRequest struct {
	// The filter rules for the component.
	Filter []*ListOpsItemsRequestFilter `json:"Filter,omitempty" xml:"Filter,omitempty" type:"Repeated"`
	// The number of entries to return on each page. Valid values: 10 to 100. Default value: 50.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that is used to retrieve the next page of results.
	//
	// example:
	//
	// MTRBMDc0NjAtRUJFNy00N0NBLTk3NTctMTJDQzQ3NjFENDdB
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The information about resource tags.
	//
	// example:
	//
	// {
	//
	//       "k1": "v1",
	//
	//       "k2": "v2"
	//
	// }
	ResourceTags map[string]interface{} `json:"ResourceTags,omitempty" xml:"ResourceTags,omitempty"`
	// The tags.
	//
	// example:
	//
	// {"k1": "v1", "k2": "v2"}
	Tags map[string]interface{} `json:"Tags,omitempty" xml:"Tags,omitempty"`
}

func (s ListOpsItemsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListOpsItemsRequest) GoString() string {
	return s.String()
}

func (s *ListOpsItemsRequest) SetFilter(v []*ListOpsItemsRequestFilter) *ListOpsItemsRequest {
	s.Filter = v
	return s
}

func (s *ListOpsItemsRequest) SetMaxResults(v int32) *ListOpsItemsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListOpsItemsRequest) SetNextToken(v string) *ListOpsItemsRequest {
	s.NextToken = &v
	return s
}

func (s *ListOpsItemsRequest) SetRegionId(v string) *ListOpsItemsRequest {
	s.RegionId = &v
	return s
}

func (s *ListOpsItemsRequest) SetResourceTags(v map[string]interface{}) *ListOpsItemsRequest {
	s.ResourceTags = v
	return s
}

func (s *ListOpsItemsRequest) SetTags(v map[string]interface{}) *ListOpsItemsRequest {
	s.Tags = v
	return s
}

type ListOpsItemsRequestFilter struct {
	// The parameter name of the filter.
	//
	// example:
	//
	// Status
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The comparison operator that is used to filter property values.
	//
	// example:
	//
	// Equal
	Operator *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	// The parameter values of the filter.
	Value []*string `json:"Value,omitempty" xml:"Value,omitempty" type:"Repeated"`
}

func (s ListOpsItemsRequestFilter) String() string {
	return tea.Prettify(s)
}

func (s ListOpsItemsRequestFilter) GoString() string {
	return s.String()
}

func (s *ListOpsItemsRequestFilter) SetName(v string) *ListOpsItemsRequestFilter {
	s.Name = &v
	return s
}

func (s *ListOpsItemsRequestFilter) SetOperator(v string) *ListOpsItemsRequestFilter {
	s.Operator = &v
	return s
}

func (s *ListOpsItemsRequestFilter) SetValue(v []*string) *ListOpsItemsRequestFilter {
	s.Value = v
	return s
}

type ListOpsItemsShrinkRequest struct {
	// The filter rules for the component.
	Filter []*ListOpsItemsShrinkRequestFilter `json:"Filter,omitempty" xml:"Filter,omitempty" type:"Repeated"`
	// The number of entries to return on each page. Valid values: 10 to 100. Default value: 50.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that is used to retrieve the next page of results.
	//
	// example:
	//
	// MTRBMDc0NjAtRUJFNy00N0NBLTk3NTctMTJDQzQ3NjFENDdB
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The information about resource tags.
	//
	// example:
	//
	// {
	//
	//       "k1": "v1",
	//
	//       "k2": "v2"
	//
	// }
	ResourceTagsShrink *string `json:"ResourceTags,omitempty" xml:"ResourceTags,omitempty"`
	// The tags.
	//
	// example:
	//
	// {"k1": "v1", "k2": "v2"}
	TagsShrink *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
}

func (s ListOpsItemsShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListOpsItemsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListOpsItemsShrinkRequest) SetFilter(v []*ListOpsItemsShrinkRequestFilter) *ListOpsItemsShrinkRequest {
	s.Filter = v
	return s
}

func (s *ListOpsItemsShrinkRequest) SetMaxResults(v int32) *ListOpsItemsShrinkRequest {
	s.MaxResults = &v
	return s
}

func (s *ListOpsItemsShrinkRequest) SetNextToken(v string) *ListOpsItemsShrinkRequest {
	s.NextToken = &v
	return s
}

func (s *ListOpsItemsShrinkRequest) SetRegionId(v string) *ListOpsItemsShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *ListOpsItemsShrinkRequest) SetResourceTagsShrink(v string) *ListOpsItemsShrinkRequest {
	s.ResourceTagsShrink = &v
	return s
}

func (s *ListOpsItemsShrinkRequest) SetTagsShrink(v string) *ListOpsItemsShrinkRequest {
	s.TagsShrink = &v
	return s
}

type ListOpsItemsShrinkRequestFilter struct {
	// The parameter name of the filter.
	//
	// example:
	//
	// Status
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The comparison operator that is used to filter property values.
	//
	// example:
	//
	// Equal
	Operator *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	// The parameter values of the filter.
	Value []*string `json:"Value,omitempty" xml:"Value,omitempty" type:"Repeated"`
}

func (s ListOpsItemsShrinkRequestFilter) String() string {
	return tea.Prettify(s)
}

func (s ListOpsItemsShrinkRequestFilter) GoString() string {
	return s.String()
}

func (s *ListOpsItemsShrinkRequestFilter) SetName(v string) *ListOpsItemsShrinkRequestFilter {
	s.Name = &v
	return s
}

func (s *ListOpsItemsShrinkRequestFilter) SetOperator(v string) *ListOpsItemsShrinkRequestFilter {
	s.Operator = &v
	return s
}

func (s *ListOpsItemsShrinkRequestFilter) SetValue(v []*string) *ListOpsItemsShrinkRequestFilter {
	s.Value = v
	return s
}

type ListOpsItemsResponseBody struct {
	// The number of entries returned on each page.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that can be used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// gAAAAABfTgv5ewUWmNdJ3g7JVLvX70sPH90GZOVGC6KPDUL0FIIb
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The list of O\\&M items.
	OpsItems []*ListOpsItemsResponseBodyOpsItems `json:"OpsItems,omitempty" xml:"OpsItems,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 113DD533-389C-5F83-9C69-F64D5BAB10B2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListOpsItemsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListOpsItemsResponseBody) GoString() string {
	return s.String()
}

func (s *ListOpsItemsResponseBody) SetMaxResults(v int32) *ListOpsItemsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListOpsItemsResponseBody) SetNextToken(v string) *ListOpsItemsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListOpsItemsResponseBody) SetOpsItems(v []*ListOpsItemsResponseBodyOpsItems) *ListOpsItemsResponseBody {
	s.OpsItems = v
	return s
}

func (s *ListOpsItemsResponseBody) SetRequestId(v string) *ListOpsItemsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListOpsItemsResponseBody) SetTotalCount(v int32) *ListOpsItemsResponseBody {
	s.TotalCount = &v
	return s
}

type ListOpsItemsResponseBodyOpsItems struct {
	// The category.
	//
	// example:
	//
	// Security
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The time when the O\\&M item was created.
	//
	// example:
	//
	// 2023-07-09T10:01Z
	CreateDate *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	// The ID of the O\\&M item.
	//
	// example:
	//
	// oi-d52b08695e2b46ae8413
	OpsItemId *string `json:"OpsItemId,omitempty" xml:"OpsItemId,omitempty"`
	// The priority.
	//
	// example:
	//
	// 1
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The Alibaba Resource Names (ARNs) of the associated resources.
	Resources []*string `json:"Resources,omitempty" xml:"Resources,omitempty" type:"Repeated"`
	// The severity level.
	//
	// example:
	//
	// Medium
	Severity *string `json:"Severity,omitempty" xml:"Severity,omitempty"`
	// The source business.
	//
	// example:
	//
	// /aliyun/ecs
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The status of the O\\&M item.
	//
	// example:
	//
	// Open
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tags.
	//
	// example:
	//
	// {"k1":"v1"}
	Tags map[string]interface{} `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The title of the O\\&M item.
	//
	// example:
	//
	// Test
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// The time when the O\\&M item was updated.
	//
	// example:
	//
	// 2023-07-09T10:01Z
	UpdateDate *string `json:"UpdateDate,omitempty" xml:"UpdateDate,omitempty"`
}

func (s ListOpsItemsResponseBodyOpsItems) String() string {
	return tea.Prettify(s)
}

func (s ListOpsItemsResponseBodyOpsItems) GoString() string {
	return s.String()
}

func (s *ListOpsItemsResponseBodyOpsItems) SetCategory(v string) *ListOpsItemsResponseBodyOpsItems {
	s.Category = &v
	return s
}

func (s *ListOpsItemsResponseBodyOpsItems) SetCreateDate(v string) *ListOpsItemsResponseBodyOpsItems {
	s.CreateDate = &v
	return s
}

func (s *ListOpsItemsResponseBodyOpsItems) SetOpsItemId(v string) *ListOpsItemsResponseBodyOpsItems {
	s.OpsItemId = &v
	return s
}

func (s *ListOpsItemsResponseBodyOpsItems) SetPriority(v int32) *ListOpsItemsResponseBodyOpsItems {
	s.Priority = &v
	return s
}

func (s *ListOpsItemsResponseBodyOpsItems) SetResources(v []*string) *ListOpsItemsResponseBodyOpsItems {
	s.Resources = v
	return s
}

func (s *ListOpsItemsResponseBodyOpsItems) SetSeverity(v string) *ListOpsItemsResponseBodyOpsItems {
	s.Severity = &v
	return s
}

func (s *ListOpsItemsResponseBodyOpsItems) SetSource(v string) *ListOpsItemsResponseBodyOpsItems {
	s.Source = &v
	return s
}

func (s *ListOpsItemsResponseBodyOpsItems) SetStatus(v string) *ListOpsItemsResponseBodyOpsItems {
	s.Status = &v
	return s
}

func (s *ListOpsItemsResponseBodyOpsItems) SetTags(v map[string]interface{}) *ListOpsItemsResponseBodyOpsItems {
	s.Tags = v
	return s
}

func (s *ListOpsItemsResponseBodyOpsItems) SetTitle(v string) *ListOpsItemsResponseBodyOpsItems {
	s.Title = &v
	return s
}

func (s *ListOpsItemsResponseBodyOpsItems) SetUpdateDate(v string) *ListOpsItemsResponseBodyOpsItems {
	s.UpdateDate = &v
	return s
}

type ListOpsItemsResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListOpsItemsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListOpsItemsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListOpsItemsResponse) GoString() string {
	return s.String()
}

func (s *ListOpsItemsResponse) SetHeaders(v map[string]*string) *ListOpsItemsResponse {
	s.Headers = v
	return s
}

func (s *ListOpsItemsResponse) SetStatusCode(v int32) *ListOpsItemsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListOpsItemsResponse) SetBody(v *ListOpsItemsResponseBody) *ListOpsItemsResponse {
	s.Body = v
	return s
}

type ListParameterVersionsRequest struct {
	// The number of entries per page. Valid values: 10 to 100. Default value: 50.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The name of the common parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// MyParameter
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// A pagination token. It can be used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// MTRBMDc0NjAtRUJFNy00N0NBLTk3NTctMTJDQzQ3NjFENDdB
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The share type of the common parameter.
	//
	// example:
	//
	// Private
	ShareType *string `json:"ShareType,omitempty" xml:"ShareType,omitempty"`
}

func (s ListParameterVersionsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListParameterVersionsRequest) GoString() string {
	return s.String()
}

func (s *ListParameterVersionsRequest) SetMaxResults(v int32) *ListParameterVersionsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListParameterVersionsRequest) SetName(v string) *ListParameterVersionsRequest {
	s.Name = &v
	return s
}

func (s *ListParameterVersionsRequest) SetNextToken(v string) *ListParameterVersionsRequest {
	s.NextToken = &v
	return s
}

func (s *ListParameterVersionsRequest) SetRegionId(v string) *ListParameterVersionsRequest {
	s.RegionId = &v
	return s
}

func (s *ListParameterVersionsRequest) SetShareType(v string) *ListParameterVersionsRequest {
	s.ShareType = &v
	return s
}

type ListParameterVersionsResponseBody struct {
	// The user who created the common parameter.
	//
	// example:
	//
	// root(130900000)
	CreatedBy *string `json:"CreatedBy,omitempty" xml:"CreatedBy,omitempty"`
	// The time when the common parameter was created.
	//
	// example:
	//
	// 2020-09-07T11:37:29Z
	CreatedDate *string `json:"CreatedDate,omitempty" xml:"CreatedDate,omitempty"`
	// The description of the common parameter.
	//
	// example:
	//
	// parameter-description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the common parameter.
	//
	// example:
	//
	// p-a483b520e0axxxxxxxxx
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The name of the common parameter.
	//
	// example:
	//
	// MyParameter
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The pagination token that was used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// MTRBMDc0NjAtRUJFNy00N0NBLTk3NTctMTJDQzQ3NjFENDdB
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The information about the version of the common parameter.
	ParameterVersions []*ListParameterVersionsResponseBodyParameterVersions `json:"ParameterVersions,omitempty" xml:"ParameterVersions,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// FD08D89D-B6C8-4AA2-A2B4-521D3F4A39FA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 50
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The data type of the common parameter.
	//
	// example:
	//
	// String
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListParameterVersionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListParameterVersionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListParameterVersionsResponseBody) SetCreatedBy(v string) *ListParameterVersionsResponseBody {
	s.CreatedBy = &v
	return s
}

func (s *ListParameterVersionsResponseBody) SetCreatedDate(v string) *ListParameterVersionsResponseBody {
	s.CreatedDate = &v
	return s
}

func (s *ListParameterVersionsResponseBody) SetDescription(v string) *ListParameterVersionsResponseBody {
	s.Description = &v
	return s
}

func (s *ListParameterVersionsResponseBody) SetId(v string) *ListParameterVersionsResponseBody {
	s.Id = &v
	return s
}

func (s *ListParameterVersionsResponseBody) SetMaxResults(v int32) *ListParameterVersionsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListParameterVersionsResponseBody) SetName(v string) *ListParameterVersionsResponseBody {
	s.Name = &v
	return s
}

func (s *ListParameterVersionsResponseBody) SetNextToken(v string) *ListParameterVersionsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListParameterVersionsResponseBody) SetParameterVersions(v []*ListParameterVersionsResponseBodyParameterVersions) *ListParameterVersionsResponseBody {
	s.ParameterVersions = v
	return s
}

func (s *ListParameterVersionsResponseBody) SetRequestId(v string) *ListParameterVersionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListParameterVersionsResponseBody) SetTotalCount(v int32) *ListParameterVersionsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListParameterVersionsResponseBody) SetType(v string) *ListParameterVersionsResponseBody {
	s.Type = &v
	return s
}

type ListParameterVersionsResponseBodyParameterVersions struct {
	// The version number of the common parameter.
	//
	// example:
	//
	// 1
	ParameterVersion *int32 `json:"ParameterVersion,omitempty" xml:"ParameterVersion,omitempty"`
	// The user who updated the common parameter.
	//
	// example:
	//
	// root(130900000)
	UpdatedBy *string `json:"UpdatedBy,omitempty" xml:"UpdatedBy,omitempty"`
	// The time when the common parameter was last updated.
	//
	// example:
	//
	// 2020-09-07T11:37:29Z
	UpdatedDate *string `json:"UpdatedDate,omitempty" xml:"UpdatedDate,omitempty"`
	// The value of the common parameter.
	//
	// example:
	//
	// MyParameter
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListParameterVersionsResponseBodyParameterVersions) String() string {
	return tea.Prettify(s)
}

func (s ListParameterVersionsResponseBodyParameterVersions) GoString() string {
	return s.String()
}

func (s *ListParameterVersionsResponseBodyParameterVersions) SetParameterVersion(v int32) *ListParameterVersionsResponseBodyParameterVersions {
	s.ParameterVersion = &v
	return s
}

func (s *ListParameterVersionsResponseBodyParameterVersions) SetUpdatedBy(v string) *ListParameterVersionsResponseBodyParameterVersions {
	s.UpdatedBy = &v
	return s
}

func (s *ListParameterVersionsResponseBodyParameterVersions) SetUpdatedDate(v string) *ListParameterVersionsResponseBodyParameterVersions {
	s.UpdatedDate = &v
	return s
}

func (s *ListParameterVersionsResponseBodyParameterVersions) SetValue(v string) *ListParameterVersionsResponseBodyParameterVersions {
	s.Value = &v
	return s
}

type ListParameterVersionsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListParameterVersionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListParameterVersionsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListParameterVersionsResponse) GoString() string {
	return s.String()
}

func (s *ListParameterVersionsResponse) SetHeaders(v map[string]*string) *ListParameterVersionsResponse {
	s.Headers = v
	return s
}

func (s *ListParameterVersionsResponse) SetStatusCode(v int32) *ListParameterVersionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListParameterVersionsResponse) SetBody(v *ListParameterVersionsResponseBody) *ListParameterVersionsResponse {
	s.Body = v
	return s
}

type ListParametersRequest struct {
	// The number of entries per page. Valid values: 10 to 100. Default value: 50.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The name of the common parameter.
	//
	// example:
	//
	// MyParameter
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The pagination token that can be used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// MTRBMDc0NjAtRUJFNy00N0NBLTk3NTctMTJDQzA
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The path of the parameter. For example, if the name of a parameter is /path/path1/Myparameter, the path of the parameter is /path/path1/.
	//
	// example:
	//
	// /path1/path2/
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// Specifies whether to query parameters from all levels of directories in the specified path. Default value: false.
	//
	// example:
	//
	// false
	Recursive *bool `json:"Recursive,omitempty" xml:"Recursive,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfmxsn4m4******
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The share type of the common parameter. Valid values:
	//
	// 	- Public
	//
	// 	- Private
	//
	// Default value: Private.
	//
	// example:
	//
	// ‘Private’
	ShareType *string `json:"ShareType,omitempty" xml:"ShareType,omitempty"`
	// The field used to sort the query results. Valid values:
	//
	// 	- Name
	//
	// 	- CreatedDate
	//
	// example:
	//
	// Name
	SortField *string `json:"SortField,omitempty" xml:"SortField,omitempty"`
	// The order in which the entries are sorted. Valid values:
	//
	// 	- Ascending
	//
	// 	- Descending (Default)
	//
	// example:
	//
	// Descending
	SortOrder *string `json:"SortOrder,omitempty" xml:"SortOrder,omitempty"`
	// The tags.
	//
	// example:
	//
	// {"k1": "v1", "k2": "v2"}
	Tags map[string]interface{} `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The data type of the common parameter.
	//
	// example:
	//
	// String
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListParametersRequest) String() string {
	return tea.Prettify(s)
}

func (s ListParametersRequest) GoString() string {
	return s.String()
}

func (s *ListParametersRequest) SetMaxResults(v int32) *ListParametersRequest {
	s.MaxResults = &v
	return s
}

func (s *ListParametersRequest) SetName(v string) *ListParametersRequest {
	s.Name = &v
	return s
}

func (s *ListParametersRequest) SetNextToken(v string) *ListParametersRequest {
	s.NextToken = &v
	return s
}

func (s *ListParametersRequest) SetPath(v string) *ListParametersRequest {
	s.Path = &v
	return s
}

func (s *ListParametersRequest) SetRecursive(v bool) *ListParametersRequest {
	s.Recursive = &v
	return s
}

func (s *ListParametersRequest) SetRegionId(v string) *ListParametersRequest {
	s.RegionId = &v
	return s
}

func (s *ListParametersRequest) SetResourceGroupId(v string) *ListParametersRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListParametersRequest) SetShareType(v string) *ListParametersRequest {
	s.ShareType = &v
	return s
}

func (s *ListParametersRequest) SetSortField(v string) *ListParametersRequest {
	s.SortField = &v
	return s
}

func (s *ListParametersRequest) SetSortOrder(v string) *ListParametersRequest {
	s.SortOrder = &v
	return s
}

func (s *ListParametersRequest) SetTags(v map[string]interface{}) *ListParametersRequest {
	s.Tags = v
	return s
}

func (s *ListParametersRequest) SetType(v string) *ListParametersRequest {
	s.Type = &v
	return s
}

type ListParametersShrinkRequest struct {
	// The number of entries per page. Valid values: 10 to 100. Default value: 50.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The name of the common parameter.
	//
	// example:
	//
	// MyParameter
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The pagination token that can be used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// MTRBMDc0NjAtRUJFNy00N0NBLTk3NTctMTJDQzA
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The path of the parameter. For example, if the name of a parameter is /path/path1/Myparameter, the path of the parameter is /path/path1/.
	//
	// example:
	//
	// /path1/path2/
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// Specifies whether to query parameters from all levels of directories in the specified path. Default value: false.
	//
	// example:
	//
	// false
	Recursive *bool `json:"Recursive,omitempty" xml:"Recursive,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfmxsn4m4******
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The share type of the common parameter. Valid values:
	//
	// 	- Public
	//
	// 	- Private
	//
	// Default value: Private.
	//
	// example:
	//
	// ‘Private’
	ShareType *string `json:"ShareType,omitempty" xml:"ShareType,omitempty"`
	// The field used to sort the query results. Valid values:
	//
	// 	- Name
	//
	// 	- CreatedDate
	//
	// example:
	//
	// Name
	SortField *string `json:"SortField,omitempty" xml:"SortField,omitempty"`
	// The order in which the entries are sorted. Valid values:
	//
	// 	- Ascending
	//
	// 	- Descending (Default)
	//
	// example:
	//
	// Descending
	SortOrder *string `json:"SortOrder,omitempty" xml:"SortOrder,omitempty"`
	// The tags.
	//
	// example:
	//
	// {"k1": "v1", "k2": "v2"}
	TagsShrink *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The data type of the common parameter.
	//
	// example:
	//
	// String
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListParametersShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListParametersShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListParametersShrinkRequest) SetMaxResults(v int32) *ListParametersShrinkRequest {
	s.MaxResults = &v
	return s
}

func (s *ListParametersShrinkRequest) SetName(v string) *ListParametersShrinkRequest {
	s.Name = &v
	return s
}

func (s *ListParametersShrinkRequest) SetNextToken(v string) *ListParametersShrinkRequest {
	s.NextToken = &v
	return s
}

func (s *ListParametersShrinkRequest) SetPath(v string) *ListParametersShrinkRequest {
	s.Path = &v
	return s
}

func (s *ListParametersShrinkRequest) SetRecursive(v bool) *ListParametersShrinkRequest {
	s.Recursive = &v
	return s
}

func (s *ListParametersShrinkRequest) SetRegionId(v string) *ListParametersShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *ListParametersShrinkRequest) SetResourceGroupId(v string) *ListParametersShrinkRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListParametersShrinkRequest) SetShareType(v string) *ListParametersShrinkRequest {
	s.ShareType = &v
	return s
}

func (s *ListParametersShrinkRequest) SetSortField(v string) *ListParametersShrinkRequest {
	s.SortField = &v
	return s
}

func (s *ListParametersShrinkRequest) SetSortOrder(v string) *ListParametersShrinkRequest {
	s.SortOrder = &v
	return s
}

func (s *ListParametersShrinkRequest) SetTagsShrink(v string) *ListParametersShrinkRequest {
	s.TagsShrink = &v
	return s
}

func (s *ListParametersShrinkRequest) SetType(v string) *ListParametersShrinkRequest {
	s.Type = &v
	return s
}

type ListParametersResponseBody struct {
	// The number of entries per page. Valid values: 10 to 100. Default value: 50.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// A pagination token. It can be used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// gAAAAABfTgv5ewUWmNdJ3g7JVLvX70sPH90GZOVGC6KPDUL0FIIb
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The information about the common parameter.
	Parameters []*ListParametersResponseBodyParameters `json:"Parameters,omitempty" xml:"Parameters,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// A81E4B2E-6B33-4BAE-9856-55DB7C893E01
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 5
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListParametersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListParametersResponseBody) GoString() string {
	return s.String()
}

func (s *ListParametersResponseBody) SetMaxResults(v int32) *ListParametersResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListParametersResponseBody) SetNextToken(v string) *ListParametersResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListParametersResponseBody) SetParameters(v []*ListParametersResponseBodyParameters) *ListParametersResponseBody {
	s.Parameters = v
	return s
}

func (s *ListParametersResponseBody) SetRequestId(v string) *ListParametersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListParametersResponseBody) SetTotalCount(v int32) *ListParametersResponseBody {
	s.TotalCount = &v
	return s
}

type ListParametersResponseBodyParameters struct {
	// The user who created the common parameter.
	//
	// example:
	//
	// root(130900000)
	CreatedBy *string `json:"CreatedBy,omitempty" xml:"CreatedBy,omitempty"`
	// The time when the common parameter was created.
	//
	// example:
	//
	// 2020-09-01T08:01:43Z
	CreatedDate *string `json:"CreatedDate,omitempty" xml:"CreatedDate,omitempty"`
	// The description of the common parameter.
	//
	// example:
	//
	// parameter
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The common parameter ID.
	//
	// example:
	//
	// p-4c4b401cab6747xxxxxx
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the common parameter.
	//
	// example:
	//
	// MyParameter
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The version number of the common parameter.
	//
	// example:
	//
	// 1
	ParameterVersion *string `json:"ParameterVersion,omitempty" xml:"ParameterVersion,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfmxsn4m4******
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The share type of the common parameter.
	//
	// example:
	//
	// Private
	ShareType *string `json:"ShareType,omitempty" xml:"ShareType,omitempty"`
	// The tags added to the common parameter.
	//
	// example:
	//
	// {"k1": "v1", "k2": "v2"}
	Tags map[string]interface{} `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The data type of the common parameter.
	//
	// example:
	//
	// String
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The user who updated the common parameter.
	//
	// example:
	//
	// root(130900000)
	UpdatedBy *string `json:"UpdatedBy,omitempty" xml:"UpdatedBy,omitempty"`
	// The time when the common parameter was updated.
	//
	// example:
	//
	// 2020-09-01T08:01:43Z
	UpdatedDate *string `json:"UpdatedDate,omitempty" xml:"UpdatedDate,omitempty"`
}

func (s ListParametersResponseBodyParameters) String() string {
	return tea.Prettify(s)
}

func (s ListParametersResponseBodyParameters) GoString() string {
	return s.String()
}

func (s *ListParametersResponseBodyParameters) SetCreatedBy(v string) *ListParametersResponseBodyParameters {
	s.CreatedBy = &v
	return s
}

func (s *ListParametersResponseBodyParameters) SetCreatedDate(v string) *ListParametersResponseBodyParameters {
	s.CreatedDate = &v
	return s
}

func (s *ListParametersResponseBodyParameters) SetDescription(v string) *ListParametersResponseBodyParameters {
	s.Description = &v
	return s
}

func (s *ListParametersResponseBodyParameters) SetId(v string) *ListParametersResponseBodyParameters {
	s.Id = &v
	return s
}

func (s *ListParametersResponseBodyParameters) SetName(v string) *ListParametersResponseBodyParameters {
	s.Name = &v
	return s
}

func (s *ListParametersResponseBodyParameters) SetParameterVersion(v string) *ListParametersResponseBodyParameters {
	s.ParameterVersion = &v
	return s
}

func (s *ListParametersResponseBodyParameters) SetResourceGroupId(v string) *ListParametersResponseBodyParameters {
	s.ResourceGroupId = &v
	return s
}

func (s *ListParametersResponseBodyParameters) SetShareType(v string) *ListParametersResponseBodyParameters {
	s.ShareType = &v
	return s
}

func (s *ListParametersResponseBodyParameters) SetTags(v map[string]interface{}) *ListParametersResponseBodyParameters {
	s.Tags = v
	return s
}

func (s *ListParametersResponseBodyParameters) SetType(v string) *ListParametersResponseBodyParameters {
	s.Type = &v
	return s
}

func (s *ListParametersResponseBodyParameters) SetUpdatedBy(v string) *ListParametersResponseBodyParameters {
	s.UpdatedBy = &v
	return s
}

func (s *ListParametersResponseBodyParameters) SetUpdatedDate(v string) *ListParametersResponseBodyParameters {
	s.UpdatedDate = &v
	return s
}

type ListParametersResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListParametersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListParametersResponse) String() string {
	return tea.Prettify(s)
}

func (s ListParametersResponse) GoString() string {
	return s.String()
}

func (s *ListParametersResponse) SetHeaders(v map[string]*string) *ListParametersResponse {
	s.Headers = v
	return s
}

func (s *ListParametersResponse) SetStatusCode(v int32) *ListParametersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListParametersResponse) SetBody(v *ListParametersResponseBody) *ListParametersResponse {
	s.Body = v
	return s
}

type ListPatchBaselinesRequest struct {
	// The approved patches.
	ApprovedPatches []*string `json:"ApprovedPatches,omitempty" xml:"ApprovedPatches,omitempty" type:"Repeated"`
	// Specifies whether the approved patch involves updates other than security-related updates.
	//
	// example:
	//
	// true
	ApprovedPatchesEnableNonSecurity *bool `json:"ApprovedPatchesEnableNonSecurity,omitempty" xml:"ApprovedPatchesEnableNonSecurity,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The name of the patch baseline.
	//
	// example:
	//
	// MyPatchBaseline
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The token that is used to retrieve the next page of results.
	//
	// example:
	//
	// -
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The type of the operating system. Valid values:
	//
	// 	- Windows
	//
	// 	- Ubuntu
	//
	// 	- CentOS
	//
	// 	- Debian
	//
	// 	- AliyunLinux
	//
	// 	- RedhatEnterpriseLinux
	//
	// 	- Anolis
	//
	// 	- AlmaLinux
	//
	// example:
	//
	// AliyunLinux
	OperationSystem *string `json:"OperationSystem,omitempty" xml:"OperationSystem,omitempty"`
	// The ID of the region.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfmxsn4m4******
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The share type of the template. Valid values:
	//
	// 	- **Public**
	//
	// 	- **Private**
	//
	// example:
	//
	// Private
	ShareType *string `json:"ShareType,omitempty" xml:"ShareType,omitempty"`
	// The patch source configurations.
	Sources []*string `json:"Sources,omitempty" xml:"Sources,omitempty" type:"Repeated"`
	// The tags.
	Tags []*ListPatchBaselinesRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s ListPatchBaselinesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListPatchBaselinesRequest) GoString() string {
	return s.String()
}

func (s *ListPatchBaselinesRequest) SetApprovedPatches(v []*string) *ListPatchBaselinesRequest {
	s.ApprovedPatches = v
	return s
}

func (s *ListPatchBaselinesRequest) SetApprovedPatchesEnableNonSecurity(v bool) *ListPatchBaselinesRequest {
	s.ApprovedPatchesEnableNonSecurity = &v
	return s
}

func (s *ListPatchBaselinesRequest) SetMaxResults(v int32) *ListPatchBaselinesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListPatchBaselinesRequest) SetName(v string) *ListPatchBaselinesRequest {
	s.Name = &v
	return s
}

func (s *ListPatchBaselinesRequest) SetNextToken(v string) *ListPatchBaselinesRequest {
	s.NextToken = &v
	return s
}

func (s *ListPatchBaselinesRequest) SetOperationSystem(v string) *ListPatchBaselinesRequest {
	s.OperationSystem = &v
	return s
}

func (s *ListPatchBaselinesRequest) SetRegionId(v string) *ListPatchBaselinesRequest {
	s.RegionId = &v
	return s
}

func (s *ListPatchBaselinesRequest) SetResourceGroupId(v string) *ListPatchBaselinesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListPatchBaselinesRequest) SetShareType(v string) *ListPatchBaselinesRequest {
	s.ShareType = &v
	return s
}

func (s *ListPatchBaselinesRequest) SetSources(v []*string) *ListPatchBaselinesRequest {
	s.Sources = v
	return s
}

func (s *ListPatchBaselinesRequest) SetTags(v []*ListPatchBaselinesRequestTags) *ListPatchBaselinesRequest {
	s.Tags = v
	return s
}

type ListPatchBaselinesRequestTags struct {
	// The key of the tag.
	//
	// example:
	//
	// key
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag.
	//
	// example:
	//
	// value
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListPatchBaselinesRequestTags) String() string {
	return tea.Prettify(s)
}

func (s ListPatchBaselinesRequestTags) GoString() string {
	return s.String()
}

func (s *ListPatchBaselinesRequestTags) SetKey(v string) *ListPatchBaselinesRequestTags {
	s.Key = &v
	return s
}

func (s *ListPatchBaselinesRequestTags) SetValue(v string) *ListPatchBaselinesRequestTags {
	s.Value = &v
	return s
}

type ListPatchBaselinesShrinkRequest struct {
	// The approved patches.
	ApprovedPatchesShrink *string `json:"ApprovedPatches,omitempty" xml:"ApprovedPatches,omitempty"`
	// Specifies whether the approved patch involves updates other than security-related updates.
	//
	// example:
	//
	// true
	ApprovedPatchesEnableNonSecurity *bool `json:"ApprovedPatchesEnableNonSecurity,omitempty" xml:"ApprovedPatchesEnableNonSecurity,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The name of the patch baseline.
	//
	// example:
	//
	// MyPatchBaseline
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The token that is used to retrieve the next page of results.
	//
	// example:
	//
	// -
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The type of the operating system. Valid values:
	//
	// 	- Windows
	//
	// 	- Ubuntu
	//
	// 	- CentOS
	//
	// 	- Debian
	//
	// 	- AliyunLinux
	//
	// 	- RedhatEnterpriseLinux
	//
	// 	- Anolis
	//
	// 	- AlmaLinux
	//
	// example:
	//
	// AliyunLinux
	OperationSystem *string `json:"OperationSystem,omitempty" xml:"OperationSystem,omitempty"`
	// The ID of the region.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfmxsn4m4******
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The share type of the template. Valid values:
	//
	// 	- **Public**
	//
	// 	- **Private**
	//
	// example:
	//
	// Private
	ShareType *string `json:"ShareType,omitempty" xml:"ShareType,omitempty"`
	// The patch source configurations.
	SourcesShrink *string `json:"Sources,omitempty" xml:"Sources,omitempty"`
	// The tags.
	TagsShrink *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
}

func (s ListPatchBaselinesShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListPatchBaselinesShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListPatchBaselinesShrinkRequest) SetApprovedPatchesShrink(v string) *ListPatchBaselinesShrinkRequest {
	s.ApprovedPatchesShrink = &v
	return s
}

func (s *ListPatchBaselinesShrinkRequest) SetApprovedPatchesEnableNonSecurity(v bool) *ListPatchBaselinesShrinkRequest {
	s.ApprovedPatchesEnableNonSecurity = &v
	return s
}

func (s *ListPatchBaselinesShrinkRequest) SetMaxResults(v int32) *ListPatchBaselinesShrinkRequest {
	s.MaxResults = &v
	return s
}

func (s *ListPatchBaselinesShrinkRequest) SetName(v string) *ListPatchBaselinesShrinkRequest {
	s.Name = &v
	return s
}

func (s *ListPatchBaselinesShrinkRequest) SetNextToken(v string) *ListPatchBaselinesShrinkRequest {
	s.NextToken = &v
	return s
}

func (s *ListPatchBaselinesShrinkRequest) SetOperationSystem(v string) *ListPatchBaselinesShrinkRequest {
	s.OperationSystem = &v
	return s
}

func (s *ListPatchBaselinesShrinkRequest) SetRegionId(v string) *ListPatchBaselinesShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *ListPatchBaselinesShrinkRequest) SetResourceGroupId(v string) *ListPatchBaselinesShrinkRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListPatchBaselinesShrinkRequest) SetShareType(v string) *ListPatchBaselinesShrinkRequest {
	s.ShareType = &v
	return s
}

func (s *ListPatchBaselinesShrinkRequest) SetSourcesShrink(v string) *ListPatchBaselinesShrinkRequest {
	s.SourcesShrink = &v
	return s
}

func (s *ListPatchBaselinesShrinkRequest) SetTagsShrink(v string) *ListPatchBaselinesShrinkRequest {
	s.TagsShrink = &v
	return s
}

type ListPatchBaselinesResponseBody struct {
	// The number of entries per page.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that is used to retrieve the next page of results.
	//
	// example:
	//
	// The number of entries returned on each page.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The patch baselines.
	PatchBaselines []*ListPatchBaselinesResponseBodyPatchBaselines `json:"PatchBaselines,omitempty" xml:"PatchBaselines,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 432996A1-03C0-5C4C-A8E6-66C4110765B8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListPatchBaselinesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListPatchBaselinesResponseBody) GoString() string {
	return s.String()
}

func (s *ListPatchBaselinesResponseBody) SetMaxResults(v int32) *ListPatchBaselinesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListPatchBaselinesResponseBody) SetNextToken(v string) *ListPatchBaselinesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListPatchBaselinesResponseBody) SetPatchBaselines(v []*ListPatchBaselinesResponseBodyPatchBaselines) *ListPatchBaselinesResponseBody {
	s.PatchBaselines = v
	return s
}

func (s *ListPatchBaselinesResponseBody) SetRequestId(v string) *ListPatchBaselinesResponseBody {
	s.RequestId = &v
	return s
}

type ListPatchBaselinesResponseBodyPatchBaselines struct {
	// The approved patches.
	ApprovedPatches []*string `json:"ApprovedPatches,omitempty" xml:"ApprovedPatches,omitempty" type:"Repeated"`
	// Indicates whether the approved patch involves updates other than security-related updates.
	//
	// example:
	//
	// true
	ApprovedPatchesEnableNonSecurity *bool `json:"ApprovedPatchesEnableNonSecurity,omitempty" xml:"ApprovedPatchesEnableNonSecurity,omitempty"`
	// The user who created the patch baseline.
	//
	// example:
	//
	// root(130900000)
	CreatedBy *string `json:"CreatedBy,omitempty" xml:"CreatedBy,omitempty"`
	// The time when the patch baseline was created.
	//
	// example:
	//
	// 2021-09-08T03:41:23Z
	CreatedDate *string `json:"CreatedDate,omitempty" xml:"CreatedDate,omitempty"`
	// The description of the patch baseline.
	//
	// example:
	//
	// ListPatchBaseline
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the patch baseline.
	//
	// example:
	//
	// pb-c2838b5d89b540e19ee6
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// Indicates whether the patch baseline is set as the default patch baseline.
	//
	// example:
	//
	// false
	IsDefault *bool `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	// The name of the patch baseline.
	//
	// example:
	//
	// MyPatchBaseline
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The type of the operating system.
	//
	// example:
	//
	// AliyunLinux
	OperationSystem *string `json:"OperationSystem,omitempty" xml:"OperationSystem,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-aek256ia6vhsndy
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The share type of the patch baseline.
	//
	// example:
	//
	// Private
	ShareType *string `json:"ShareType,omitempty" xml:"ShareType,omitempty"`
	// The configurations of patch sources.
	Sources []*string `json:"Sources,omitempty" xml:"Sources,omitempty" type:"Repeated"`
	// The tags of the patch baseline.
	Tags []*ListPatchBaselinesResponseBodyPatchBaselinesTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The user who last updated the patch baseline.
	//
	// example:
	//
	// root(130900000)
	UpdatedBy *string `json:"UpdatedBy,omitempty" xml:"UpdatedBy,omitempty"`
	// The time when the patch baseline was updated.
	//
	// example:
	//
	// 2021-09-08T03:44:34Z
	UpdatedDate *string `json:"UpdatedDate,omitempty" xml:"UpdatedDate,omitempty"`
}

func (s ListPatchBaselinesResponseBodyPatchBaselines) String() string {
	return tea.Prettify(s)
}

func (s ListPatchBaselinesResponseBodyPatchBaselines) GoString() string {
	return s.String()
}

func (s *ListPatchBaselinesResponseBodyPatchBaselines) SetApprovedPatches(v []*string) *ListPatchBaselinesResponseBodyPatchBaselines {
	s.ApprovedPatches = v
	return s
}

func (s *ListPatchBaselinesResponseBodyPatchBaselines) SetApprovedPatchesEnableNonSecurity(v bool) *ListPatchBaselinesResponseBodyPatchBaselines {
	s.ApprovedPatchesEnableNonSecurity = &v
	return s
}

func (s *ListPatchBaselinesResponseBodyPatchBaselines) SetCreatedBy(v string) *ListPatchBaselinesResponseBodyPatchBaselines {
	s.CreatedBy = &v
	return s
}

func (s *ListPatchBaselinesResponseBodyPatchBaselines) SetCreatedDate(v string) *ListPatchBaselinesResponseBodyPatchBaselines {
	s.CreatedDate = &v
	return s
}

func (s *ListPatchBaselinesResponseBodyPatchBaselines) SetDescription(v string) *ListPatchBaselinesResponseBodyPatchBaselines {
	s.Description = &v
	return s
}

func (s *ListPatchBaselinesResponseBodyPatchBaselines) SetId(v string) *ListPatchBaselinesResponseBodyPatchBaselines {
	s.Id = &v
	return s
}

func (s *ListPatchBaselinesResponseBodyPatchBaselines) SetIsDefault(v bool) *ListPatchBaselinesResponseBodyPatchBaselines {
	s.IsDefault = &v
	return s
}

func (s *ListPatchBaselinesResponseBodyPatchBaselines) SetName(v string) *ListPatchBaselinesResponseBodyPatchBaselines {
	s.Name = &v
	return s
}

func (s *ListPatchBaselinesResponseBodyPatchBaselines) SetOperationSystem(v string) *ListPatchBaselinesResponseBodyPatchBaselines {
	s.OperationSystem = &v
	return s
}

func (s *ListPatchBaselinesResponseBodyPatchBaselines) SetResourceGroupId(v string) *ListPatchBaselinesResponseBodyPatchBaselines {
	s.ResourceGroupId = &v
	return s
}

func (s *ListPatchBaselinesResponseBodyPatchBaselines) SetShareType(v string) *ListPatchBaselinesResponseBodyPatchBaselines {
	s.ShareType = &v
	return s
}

func (s *ListPatchBaselinesResponseBodyPatchBaselines) SetSources(v []*string) *ListPatchBaselinesResponseBodyPatchBaselines {
	s.Sources = v
	return s
}

func (s *ListPatchBaselinesResponseBodyPatchBaselines) SetTags(v []*ListPatchBaselinesResponseBodyPatchBaselinesTags) *ListPatchBaselinesResponseBodyPatchBaselines {
	s.Tags = v
	return s
}

func (s *ListPatchBaselinesResponseBodyPatchBaselines) SetUpdatedBy(v string) *ListPatchBaselinesResponseBodyPatchBaselines {
	s.UpdatedBy = &v
	return s
}

func (s *ListPatchBaselinesResponseBodyPatchBaselines) SetUpdatedDate(v string) *ListPatchBaselinesResponseBodyPatchBaselines {
	s.UpdatedDate = &v
	return s
}

type ListPatchBaselinesResponseBodyPatchBaselinesTags struct {
	// The key of the tag.
	//
	// example:
	//
	// key
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The value of the tag.
	//
	// example:
	//
	// value
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s ListPatchBaselinesResponseBodyPatchBaselinesTags) String() string {
	return tea.Prettify(s)
}

func (s ListPatchBaselinesResponseBodyPatchBaselinesTags) GoString() string {
	return s.String()
}

func (s *ListPatchBaselinesResponseBodyPatchBaselinesTags) SetTagKey(v string) *ListPatchBaselinesResponseBodyPatchBaselinesTags {
	s.TagKey = &v
	return s
}

func (s *ListPatchBaselinesResponseBodyPatchBaselinesTags) SetTagValue(v string) *ListPatchBaselinesResponseBodyPatchBaselinesTags {
	s.TagValue = &v
	return s
}

type ListPatchBaselinesResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPatchBaselinesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPatchBaselinesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListPatchBaselinesResponse) GoString() string {
	return s.String()
}

func (s *ListPatchBaselinesResponse) SetHeaders(v map[string]*string) *ListPatchBaselinesResponse {
	s.Headers = v
	return s
}

func (s *ListPatchBaselinesResponse) SetStatusCode(v int32) *ListPatchBaselinesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPatchBaselinesResponse) SetBody(v *ListPatchBaselinesResponseBody) *ListPatchBaselinesResponse {
	s.Body = v
	return s
}

type ListResourceExecutionStatusRequest struct {
	// The ID of the execution.
	//
	// This parameter is required.
	//
	// example:
	//
	// exec-xxxxxxxxxxxx
	ExecutionId *string `json:"ExecutionId,omitempty" xml:"ExecutionId,omitempty"`
	// The number of entries to return on each page. Valid values: 10 to 100. Default value: 50.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// MTRBMDc0NjAtRUJFNy00N0NBLTk3NTctMTJDQzQ
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the region.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListResourceExecutionStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s ListResourceExecutionStatusRequest) GoString() string {
	return s.String()
}

func (s *ListResourceExecutionStatusRequest) SetExecutionId(v string) *ListResourceExecutionStatusRequest {
	s.ExecutionId = &v
	return s
}

func (s *ListResourceExecutionStatusRequest) SetMaxResults(v int32) *ListResourceExecutionStatusRequest {
	s.MaxResults = &v
	return s
}

func (s *ListResourceExecutionStatusRequest) SetNextToken(v string) *ListResourceExecutionStatusRequest {
	s.NextToken = &v
	return s
}

func (s *ListResourceExecutionStatusRequest) SetRegionId(v string) *ListResourceExecutionStatusRequest {
	s.RegionId = &v
	return s
}

type ListResourceExecutionStatusResponseBody struct {
	// The number of entries returned on each page.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// MTRBMDc0NjAtRUJFNy00N0NBLTk3NTctMTJDQzQ
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// ED571CBD-9F96-419D-B919-CF340BBCA157
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The execution information of the resource.
	ResourceExecutionStatus []*ListResourceExecutionStatusResponseBodyResourceExecutionStatus `json:"ResourceExecutionStatus,omitempty" xml:"ResourceExecutionStatus,omitempty" type:"Repeated"`
}

func (s ListResourceExecutionStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListResourceExecutionStatusResponseBody) GoString() string {
	return s.String()
}

func (s *ListResourceExecutionStatusResponseBody) SetMaxResults(v int32) *ListResourceExecutionStatusResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListResourceExecutionStatusResponseBody) SetNextToken(v string) *ListResourceExecutionStatusResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListResourceExecutionStatusResponseBody) SetRequestId(v string) *ListResourceExecutionStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListResourceExecutionStatusResponseBody) SetResourceExecutionStatus(v []*ListResourceExecutionStatusResponseBodyResourceExecutionStatus) *ListResourceExecutionStatusResponseBody {
	s.ResourceExecutionStatus = v
	return s
}

type ListResourceExecutionStatusResponseBodyResourceExecutionStatus struct {
	// The ID of the execution.
	//
	// example:
	//
	// exec-6be6d6ff805349d9ac13
	ExecutionId *string `json:"ExecutionId,omitempty" xml:"ExecutionId,omitempty"`
	// The time when the execution started running.
	//
	// example:
	//
	// 2020-11-13T08:48:34Z
	ExecutionTime *string `json:"ExecutionTime,omitempty" xml:"ExecutionTime,omitempty"`
	// The output of the template.
	//
	// example:
	//
	// { 				"commandOutput": "hello\\n" 			}
	Outputs *string `json:"Outputs,omitempty" xml:"Outputs,omitempty"`
	// The ID of the resource.
	//
	// example:
	//
	// i-bp1e1bxxxxxxxxxxxxxx
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The status of the execution.
	//
	// example:
	//
	// Success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListResourceExecutionStatusResponseBodyResourceExecutionStatus) String() string {
	return tea.Prettify(s)
}

func (s ListResourceExecutionStatusResponseBodyResourceExecutionStatus) GoString() string {
	return s.String()
}

func (s *ListResourceExecutionStatusResponseBodyResourceExecutionStatus) SetExecutionId(v string) *ListResourceExecutionStatusResponseBodyResourceExecutionStatus {
	s.ExecutionId = &v
	return s
}

func (s *ListResourceExecutionStatusResponseBodyResourceExecutionStatus) SetExecutionTime(v string) *ListResourceExecutionStatusResponseBodyResourceExecutionStatus {
	s.ExecutionTime = &v
	return s
}

func (s *ListResourceExecutionStatusResponseBodyResourceExecutionStatus) SetOutputs(v string) *ListResourceExecutionStatusResponseBodyResourceExecutionStatus {
	s.Outputs = &v
	return s
}

func (s *ListResourceExecutionStatusResponseBodyResourceExecutionStatus) SetResourceId(v string) *ListResourceExecutionStatusResponseBodyResourceExecutionStatus {
	s.ResourceId = &v
	return s
}

func (s *ListResourceExecutionStatusResponseBodyResourceExecutionStatus) SetStatus(v string) *ListResourceExecutionStatusResponseBodyResourceExecutionStatus {
	s.Status = &v
	return s
}

type ListResourceExecutionStatusResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListResourceExecutionStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListResourceExecutionStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s ListResourceExecutionStatusResponse) GoString() string {
	return s.String()
}

func (s *ListResourceExecutionStatusResponse) SetHeaders(v map[string]*string) *ListResourceExecutionStatusResponse {
	s.Headers = v
	return s
}

func (s *ListResourceExecutionStatusResponse) SetStatusCode(v int32) *ListResourceExecutionStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *ListResourceExecutionStatusResponse) SetBody(v *ListResourceExecutionStatusResponseBody) *ListResourceExecutionStatusResponse {
	s.Body = v
	return s
}

type ListSecretParameterVersionsRequest struct {
	// The number of entries per page. Valid values: 10 to 100. Default value: 50.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The name of the encryption parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// MySecretParameter
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// MTRBMDc0NjAtRUJFNy00N0NBLTk3NTctMTJDQzQ3NjFENDdB
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the region.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The share type of the encryption parameter.
	//
	// example:
	//
	// Private
	ShareType *string `json:"ShareType,omitempty" xml:"ShareType,omitempty"`
	// Specifies whether to decrypt the parameter value. The decrypted parameter value is returned only if this parameter is set to true. Otherwise, null is returned.
	//
	// example:
	//
	// false
	WithDecryption *bool `json:"WithDecryption,omitempty" xml:"WithDecryption,omitempty"`
}

func (s ListSecretParameterVersionsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListSecretParameterVersionsRequest) GoString() string {
	return s.String()
}

func (s *ListSecretParameterVersionsRequest) SetMaxResults(v int32) *ListSecretParameterVersionsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListSecretParameterVersionsRequest) SetName(v string) *ListSecretParameterVersionsRequest {
	s.Name = &v
	return s
}

func (s *ListSecretParameterVersionsRequest) SetNextToken(v string) *ListSecretParameterVersionsRequest {
	s.NextToken = &v
	return s
}

func (s *ListSecretParameterVersionsRequest) SetRegionId(v string) *ListSecretParameterVersionsRequest {
	s.RegionId = &v
	return s
}

func (s *ListSecretParameterVersionsRequest) SetShareType(v string) *ListSecretParameterVersionsRequest {
	s.ShareType = &v
	return s
}

func (s *ListSecretParameterVersionsRequest) SetWithDecryption(v bool) *ListSecretParameterVersionsRequest {
	s.WithDecryption = &v
	return s
}

type ListSecretParameterVersionsResponseBody struct {
	// The user who created the encryption parameter.
	//
	// example:
	//
	// root(130900000)
	CreatedBy *string `json:"CreatedBy,omitempty" xml:"CreatedBy,omitempty"`
	// The time when the encryption parameter was created.
	//
	// example:
	//
	// 2020-09-01T08:01:43Z
	CreatedDate *string `json:"CreatedDate,omitempty" xml:"CreatedDate,omitempty"`
	// The description of the encryption parameter.
	//
	// example:
	//
	// SecretParameter
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the encryption parameter.
	//
	// example:
	//
	// p-4c4b401cab6747xxxxxx
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The name of the encryption parameter.
	//
	// example:
	//
	// MySecretParameter
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// MTRBMDc0NjAtRUJFNy00N0NBLTk3NTctMTJDQzQ3NjFENDdB
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The information about the version of the encryption parameter.
	ParameterVersions []*ListSecretParameterVersionsResponseBodyParameterVersions `json:"ParameterVersions,omitempty" xml:"ParameterVersions,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// DBA6E6C8-F75D-41DE-AFF5-1FA03F551CA3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The type of the encryption parameter.
	//
	// example:
	//
	// Secret
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListSecretParameterVersionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListSecretParameterVersionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListSecretParameterVersionsResponseBody) SetCreatedBy(v string) *ListSecretParameterVersionsResponseBody {
	s.CreatedBy = &v
	return s
}

func (s *ListSecretParameterVersionsResponseBody) SetCreatedDate(v string) *ListSecretParameterVersionsResponseBody {
	s.CreatedDate = &v
	return s
}

func (s *ListSecretParameterVersionsResponseBody) SetDescription(v string) *ListSecretParameterVersionsResponseBody {
	s.Description = &v
	return s
}

func (s *ListSecretParameterVersionsResponseBody) SetId(v string) *ListSecretParameterVersionsResponseBody {
	s.Id = &v
	return s
}

func (s *ListSecretParameterVersionsResponseBody) SetMaxResults(v int32) *ListSecretParameterVersionsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListSecretParameterVersionsResponseBody) SetName(v string) *ListSecretParameterVersionsResponseBody {
	s.Name = &v
	return s
}

func (s *ListSecretParameterVersionsResponseBody) SetNextToken(v string) *ListSecretParameterVersionsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListSecretParameterVersionsResponseBody) SetParameterVersions(v []*ListSecretParameterVersionsResponseBodyParameterVersions) *ListSecretParameterVersionsResponseBody {
	s.ParameterVersions = v
	return s
}

func (s *ListSecretParameterVersionsResponseBody) SetRequestId(v string) *ListSecretParameterVersionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSecretParameterVersionsResponseBody) SetTotalCount(v int32) *ListSecretParameterVersionsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListSecretParameterVersionsResponseBody) SetType(v string) *ListSecretParameterVersionsResponseBody {
	s.Type = &v
	return s
}

type ListSecretParameterVersionsResponseBodyParameterVersions struct {
	// The version number of the encryption parameter.
	//
	// example:
	//
	// 1
	ParameterVersion *int32 `json:"ParameterVersion,omitempty" xml:"ParameterVersion,omitempty"`
	// The user who updated the encryption parameter.
	//
	// example:
	//
	// root(130900000)
	UpdatedBy *string `json:"UpdatedBy,omitempty" xml:"UpdatedBy,omitempty"`
	// The time when the encryption parameter was updated.
	//
	// example:
	//
	// 2020-09-01T08:01:43Z
	UpdatedDate *string `json:"UpdatedDate,omitempty" xml:"UpdatedDate,omitempty"`
	// The value of the encryption parameter.
	//
	// example:
	//
	// SecretParameter
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListSecretParameterVersionsResponseBodyParameterVersions) String() string {
	return tea.Prettify(s)
}

func (s ListSecretParameterVersionsResponseBodyParameterVersions) GoString() string {
	return s.String()
}

func (s *ListSecretParameterVersionsResponseBodyParameterVersions) SetParameterVersion(v int32) *ListSecretParameterVersionsResponseBodyParameterVersions {
	s.ParameterVersion = &v
	return s
}

func (s *ListSecretParameterVersionsResponseBodyParameterVersions) SetUpdatedBy(v string) *ListSecretParameterVersionsResponseBodyParameterVersions {
	s.UpdatedBy = &v
	return s
}

func (s *ListSecretParameterVersionsResponseBodyParameterVersions) SetUpdatedDate(v string) *ListSecretParameterVersionsResponseBodyParameterVersions {
	s.UpdatedDate = &v
	return s
}

func (s *ListSecretParameterVersionsResponseBodyParameterVersions) SetValue(v string) *ListSecretParameterVersionsResponseBodyParameterVersions {
	s.Value = &v
	return s
}

type ListSecretParameterVersionsResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSecretParameterVersionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSecretParameterVersionsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListSecretParameterVersionsResponse) GoString() string {
	return s.String()
}

func (s *ListSecretParameterVersionsResponse) SetHeaders(v map[string]*string) *ListSecretParameterVersionsResponse {
	s.Headers = v
	return s
}

func (s *ListSecretParameterVersionsResponse) SetStatusCode(v int32) *ListSecretParameterVersionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSecretParameterVersionsResponse) SetBody(v *ListSecretParameterVersionsResponseBody) *ListSecretParameterVersionsResponse {
	s.Body = v
	return s
}

type ListSecretParametersRequest struct {
	// The number of entries per page. Valid values: 10 to 100. Default value: 50.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The name of the parameter. **You can enter a keyword to query parameter names in fuzzy match mode.
	//
	// example:
	//
	// MySecretParameter
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The token that is used to retrieve the next page of results.
	//
	// example:
	//
	// sPH90GZOVGC6KPDUL0FIIbEtMQHq_19S6_4O_XqA
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The path of the parameter. For example, if the name of a parameter is /path/path1/Myparameter, the path of the parameter is /path/path1/.
	//
	// example:
	//
	// /path1/path2/
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// Specifies whether to query parameters from all levels of directories in the specified path. Default value: false.
	//
	// example:
	//
	// false
	Recursive *bool `json:"Recursive,omitempty" xml:"Recursive,omitempty"`
	// The ID of the region.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfmxsn4m4******
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The field used to sort the query results. Valid values:
	//
	// 	- Name
	//
	// 	- CreatedDate
	//
	// example:
	//
	// Name
	SortField *string `json:"SortField,omitempty" xml:"SortField,omitempty"`
	// The order in which the entries are sorted. Valid values:
	//
	// 	- Ascending
	//
	// 	- Descending (Default)
	//
	// example:
	//
	// Descending
	SortOrder *string `json:"SortOrder,omitempty" xml:"SortOrder,omitempty"`
	// The tags of the parameter.
	//
	// example:
	//
	// {"k1": "v1", "k2": "v2"}
	Tags map[string]interface{} `json:"Tags,omitempty" xml:"Tags,omitempty"`
}

func (s ListSecretParametersRequest) String() string {
	return tea.Prettify(s)
}

func (s ListSecretParametersRequest) GoString() string {
	return s.String()
}

func (s *ListSecretParametersRequest) SetMaxResults(v int32) *ListSecretParametersRequest {
	s.MaxResults = &v
	return s
}

func (s *ListSecretParametersRequest) SetName(v string) *ListSecretParametersRequest {
	s.Name = &v
	return s
}

func (s *ListSecretParametersRequest) SetNextToken(v string) *ListSecretParametersRequest {
	s.NextToken = &v
	return s
}

func (s *ListSecretParametersRequest) SetPath(v string) *ListSecretParametersRequest {
	s.Path = &v
	return s
}

func (s *ListSecretParametersRequest) SetRecursive(v bool) *ListSecretParametersRequest {
	s.Recursive = &v
	return s
}

func (s *ListSecretParametersRequest) SetRegionId(v string) *ListSecretParametersRequest {
	s.RegionId = &v
	return s
}

func (s *ListSecretParametersRequest) SetResourceGroupId(v string) *ListSecretParametersRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListSecretParametersRequest) SetSortField(v string) *ListSecretParametersRequest {
	s.SortField = &v
	return s
}

func (s *ListSecretParametersRequest) SetSortOrder(v string) *ListSecretParametersRequest {
	s.SortOrder = &v
	return s
}

func (s *ListSecretParametersRequest) SetTags(v map[string]interface{}) *ListSecretParametersRequest {
	s.Tags = v
	return s
}

type ListSecretParametersShrinkRequest struct {
	// The number of entries per page. Valid values: 10 to 100. Default value: 50.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The name of the parameter. **You can enter a keyword to query parameter names in fuzzy match mode.
	//
	// example:
	//
	// MySecretParameter
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The token that is used to retrieve the next page of results.
	//
	// example:
	//
	// sPH90GZOVGC6KPDUL0FIIbEtMQHq_19S6_4O_XqA
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The path of the parameter. For example, if the name of a parameter is /path/path1/Myparameter, the path of the parameter is /path/path1/.
	//
	// example:
	//
	// /path1/path2/
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// Specifies whether to query parameters from all levels of directories in the specified path. Default value: false.
	//
	// example:
	//
	// false
	Recursive *bool `json:"Recursive,omitempty" xml:"Recursive,omitempty"`
	// The ID of the region.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfmxsn4m4******
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The field used to sort the query results. Valid values:
	//
	// 	- Name
	//
	// 	- CreatedDate
	//
	// example:
	//
	// Name
	SortField *string `json:"SortField,omitempty" xml:"SortField,omitempty"`
	// The order in which the entries are sorted. Valid values:
	//
	// 	- Ascending
	//
	// 	- Descending (Default)
	//
	// example:
	//
	// Descending
	SortOrder *string `json:"SortOrder,omitempty" xml:"SortOrder,omitempty"`
	// The tags of the parameter.
	//
	// example:
	//
	// {"k1": "v1", "k2": "v2"}
	TagsShrink *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
}

func (s ListSecretParametersShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListSecretParametersShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListSecretParametersShrinkRequest) SetMaxResults(v int32) *ListSecretParametersShrinkRequest {
	s.MaxResults = &v
	return s
}

func (s *ListSecretParametersShrinkRequest) SetName(v string) *ListSecretParametersShrinkRequest {
	s.Name = &v
	return s
}

func (s *ListSecretParametersShrinkRequest) SetNextToken(v string) *ListSecretParametersShrinkRequest {
	s.NextToken = &v
	return s
}

func (s *ListSecretParametersShrinkRequest) SetPath(v string) *ListSecretParametersShrinkRequest {
	s.Path = &v
	return s
}

func (s *ListSecretParametersShrinkRequest) SetRecursive(v bool) *ListSecretParametersShrinkRequest {
	s.Recursive = &v
	return s
}

func (s *ListSecretParametersShrinkRequest) SetRegionId(v string) *ListSecretParametersShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *ListSecretParametersShrinkRequest) SetResourceGroupId(v string) *ListSecretParametersShrinkRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListSecretParametersShrinkRequest) SetSortField(v string) *ListSecretParametersShrinkRequest {
	s.SortField = &v
	return s
}

func (s *ListSecretParametersShrinkRequest) SetSortOrder(v string) *ListSecretParametersShrinkRequest {
	s.SortOrder = &v
	return s
}

func (s *ListSecretParametersShrinkRequest) SetTagsShrink(v string) *ListSecretParametersShrinkRequest {
	s.TagsShrink = &v
	return s
}

type ListSecretParametersResponseBody struct {
	// The number of entries returned per page.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that is used to retrieve the next page of results.
	//
	// example:
	//
	// sPH90GZOVGC6KPDUL0FIIbEtMQHq_19S6_4O_XqA
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The information about the parameters.
	Parameters []*ListSecretParametersResponseBodyParameters `json:"Parameters,omitempty" xml:"Parameters,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// CA9C6248-AF2A-4AE9-A166-88FD901BBB90
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListSecretParametersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListSecretParametersResponseBody) GoString() string {
	return s.String()
}

func (s *ListSecretParametersResponseBody) SetMaxResults(v int32) *ListSecretParametersResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListSecretParametersResponseBody) SetNextToken(v string) *ListSecretParametersResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListSecretParametersResponseBody) SetParameters(v []*ListSecretParametersResponseBodyParameters) *ListSecretParametersResponseBody {
	s.Parameters = v
	return s
}

func (s *ListSecretParametersResponseBody) SetRequestId(v string) *ListSecretParametersResponseBody {
	s.RequestId = &v
	return s
}

type ListSecretParametersResponseBodyParameters struct {
	// The user who created the parameter.
	//
	// example:
	//
	// root(130900000)
	CreatedBy *string `json:"CreatedBy,omitempty" xml:"CreatedBy,omitempty"`
	// The time when the parameter was created.
	//
	// example:
	//
	// 2020-09-01T09:28:47Z
	CreatedDate *string `json:"CreatedDate,omitempty" xml:"CreatedDate,omitempty"`
	// The description of the parameter.
	//
	// example:
	//
	// SecretParameter
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the parameter.
	//
	// example:
	//
	// p-14ed150fdcd048xxxxxx
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The ID of the KMS customer master key (CMK) that is used for encryption.
	//
	// example:
	//
	// 80e9409f-78fa-42ab-84bd-83f40c******
	KeyId *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	// The name of the parameter.
	//
	// example:
	//
	// MySecretParameter
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The version number of the parameter.
	//
	// example:
	//
	// 1
	ParameterVersion *string `json:"ParameterVersion,omitempty" xml:"ParameterVersion,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfmxsn4m4******
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The share type of the parameter.
	//
	// example:
	//
	// Private
	ShareType *string `json:"ShareType,omitempty" xml:"ShareType,omitempty"`
	// The tags of the parameter.
	//
	// example:
	//
	// {"k1": "v1", "k2": "v2"}
	Tags map[string]interface{} `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The type of the parameter.
	//
	// example:
	//
	// Secret
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The user who updated the parameter.
	//
	// example:
	//
	// root(130900000)
	UpdatedBy *string `json:"UpdatedBy,omitempty" xml:"UpdatedBy,omitempty"`
	// The time when the parameter was updated.
	//
	// example:
	//
	// 2020-09-01T09:35:17Z
	UpdatedDate *string `json:"UpdatedDate,omitempty" xml:"UpdatedDate,omitempty"`
}

func (s ListSecretParametersResponseBodyParameters) String() string {
	return tea.Prettify(s)
}

func (s ListSecretParametersResponseBodyParameters) GoString() string {
	return s.String()
}

func (s *ListSecretParametersResponseBodyParameters) SetCreatedBy(v string) *ListSecretParametersResponseBodyParameters {
	s.CreatedBy = &v
	return s
}

func (s *ListSecretParametersResponseBodyParameters) SetCreatedDate(v string) *ListSecretParametersResponseBodyParameters {
	s.CreatedDate = &v
	return s
}

func (s *ListSecretParametersResponseBodyParameters) SetDescription(v string) *ListSecretParametersResponseBodyParameters {
	s.Description = &v
	return s
}

func (s *ListSecretParametersResponseBodyParameters) SetId(v string) *ListSecretParametersResponseBodyParameters {
	s.Id = &v
	return s
}

func (s *ListSecretParametersResponseBodyParameters) SetKeyId(v string) *ListSecretParametersResponseBodyParameters {
	s.KeyId = &v
	return s
}

func (s *ListSecretParametersResponseBodyParameters) SetName(v string) *ListSecretParametersResponseBodyParameters {
	s.Name = &v
	return s
}

func (s *ListSecretParametersResponseBodyParameters) SetParameterVersion(v string) *ListSecretParametersResponseBodyParameters {
	s.ParameterVersion = &v
	return s
}

func (s *ListSecretParametersResponseBodyParameters) SetResourceGroupId(v string) *ListSecretParametersResponseBodyParameters {
	s.ResourceGroupId = &v
	return s
}

func (s *ListSecretParametersResponseBodyParameters) SetShareType(v string) *ListSecretParametersResponseBodyParameters {
	s.ShareType = &v
	return s
}

func (s *ListSecretParametersResponseBodyParameters) SetTags(v map[string]interface{}) *ListSecretParametersResponseBodyParameters {
	s.Tags = v
	return s
}

func (s *ListSecretParametersResponseBodyParameters) SetType(v string) *ListSecretParametersResponseBodyParameters {
	s.Type = &v
	return s
}

func (s *ListSecretParametersResponseBodyParameters) SetUpdatedBy(v string) *ListSecretParametersResponseBodyParameters {
	s.UpdatedBy = &v
	return s
}

func (s *ListSecretParametersResponseBodyParameters) SetUpdatedDate(v string) *ListSecretParametersResponseBodyParameters {
	s.UpdatedDate = &v
	return s
}

type ListSecretParametersResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSecretParametersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSecretParametersResponse) String() string {
	return tea.Prettify(s)
}

func (s ListSecretParametersResponse) GoString() string {
	return s.String()
}

func (s *ListSecretParametersResponse) SetHeaders(v map[string]*string) *ListSecretParametersResponse {
	s.Headers = v
	return s
}

func (s *ListSecretParametersResponse) SetStatusCode(v int32) *ListSecretParametersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSecretParametersResponse) SetBody(v *ListSecretParametersResponseBody) *ListSecretParametersResponse {
	s.Body = v
	return s
}

type ListStateConfigurationsRequest struct {
	// The maximum number of entries per page.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// A pagination token. It can be used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// AHJKH-AHKJHDJK-AKHDIOWJL
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfmxsn4m4******
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The ID of the desired-state configuration.
	//
	// example:
	//
	// ["sc-asfgdhj12345"]
	StateConfigurationIds *string `json:"StateConfigurationIds,omitempty" xml:"StateConfigurationIds,omitempty"`
	// The tags to be added to the configuration.
	//
	// example:
	//
	// {"Key": "oos", "Value": "inventory"}
	Tags map[string]interface{} `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The name of the template. The name must be 1 to 200 characters in length and can contain letters, digits, hyphens (-), and underscores (_).
	//
	// example:
	//
	// ACS-ECS-InventoryDataCollection
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// The version number of the template. If you do not specify this parameter, the latest version of the template is used.
	//
	// example:
	//
	// v1
	TemplateVersion *string `json:"TemplateVersion,omitempty" xml:"TemplateVersion,omitempty"`
}

func (s ListStateConfigurationsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListStateConfigurationsRequest) GoString() string {
	return s.String()
}

func (s *ListStateConfigurationsRequest) SetMaxResults(v int32) *ListStateConfigurationsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListStateConfigurationsRequest) SetNextToken(v string) *ListStateConfigurationsRequest {
	s.NextToken = &v
	return s
}

func (s *ListStateConfigurationsRequest) SetRegionId(v string) *ListStateConfigurationsRequest {
	s.RegionId = &v
	return s
}

func (s *ListStateConfigurationsRequest) SetResourceGroupId(v string) *ListStateConfigurationsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListStateConfigurationsRequest) SetStateConfigurationIds(v string) *ListStateConfigurationsRequest {
	s.StateConfigurationIds = &v
	return s
}

func (s *ListStateConfigurationsRequest) SetTags(v map[string]interface{}) *ListStateConfigurationsRequest {
	s.Tags = v
	return s
}

func (s *ListStateConfigurationsRequest) SetTemplateName(v string) *ListStateConfigurationsRequest {
	s.TemplateName = &v
	return s
}

func (s *ListStateConfigurationsRequest) SetTemplateVersion(v string) *ListStateConfigurationsRequest {
	s.TemplateVersion = &v
	return s
}

type ListStateConfigurationsShrinkRequest struct {
	// The maximum number of entries per page.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// A pagination token. It can be used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// AHJKH-AHKJHDJK-AKHDIOWJL
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfmxsn4m4******
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The ID of the desired-state configuration.
	//
	// example:
	//
	// ["sc-asfgdhj12345"]
	StateConfigurationIds *string `json:"StateConfigurationIds,omitempty" xml:"StateConfigurationIds,omitempty"`
	// The tags to be added to the configuration.
	//
	// example:
	//
	// {"Key": "oos", "Value": "inventory"}
	TagsShrink *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The name of the template. The name must be 1 to 200 characters in length and can contain letters, digits, hyphens (-), and underscores (_).
	//
	// example:
	//
	// ACS-ECS-InventoryDataCollection
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// The version number of the template. If you do not specify this parameter, the latest version of the template is used.
	//
	// example:
	//
	// v1
	TemplateVersion *string `json:"TemplateVersion,omitempty" xml:"TemplateVersion,omitempty"`
}

func (s ListStateConfigurationsShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListStateConfigurationsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListStateConfigurationsShrinkRequest) SetMaxResults(v int32) *ListStateConfigurationsShrinkRequest {
	s.MaxResults = &v
	return s
}

func (s *ListStateConfigurationsShrinkRequest) SetNextToken(v string) *ListStateConfigurationsShrinkRequest {
	s.NextToken = &v
	return s
}

func (s *ListStateConfigurationsShrinkRequest) SetRegionId(v string) *ListStateConfigurationsShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *ListStateConfigurationsShrinkRequest) SetResourceGroupId(v string) *ListStateConfigurationsShrinkRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListStateConfigurationsShrinkRequest) SetStateConfigurationIds(v string) *ListStateConfigurationsShrinkRequest {
	s.StateConfigurationIds = &v
	return s
}

func (s *ListStateConfigurationsShrinkRequest) SetTagsShrink(v string) *ListStateConfigurationsShrinkRequest {
	s.TagsShrink = &v
	return s
}

func (s *ListStateConfigurationsShrinkRequest) SetTemplateName(v string) *ListStateConfigurationsShrinkRequest {
	s.TemplateName = &v
	return s
}

func (s *ListStateConfigurationsShrinkRequest) SetTemplateVersion(v string) *ListStateConfigurationsShrinkRequest {
	s.TemplateVersion = &v
	return s
}

type ListStateConfigurationsResponseBody struct {
	// The pagination token that was used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// AAAAASO3cL82+b5iji7bfPNpMh8=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1306108F-610C-40FD-AAD5-DA13E8B00BE9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the desired-state configurations.
	StateConfigurations []*ListStateConfigurationsResponseBodyStateConfigurations `json:"StateConfigurations,omitempty" xml:"StateConfigurations,omitempty" type:"Repeated"`
}

func (s ListStateConfigurationsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListStateConfigurationsResponseBody) GoString() string {
	return s.String()
}

func (s *ListStateConfigurationsResponseBody) SetNextToken(v string) *ListStateConfigurationsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListStateConfigurationsResponseBody) SetRequestId(v string) *ListStateConfigurationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListStateConfigurationsResponseBody) SetStateConfigurations(v []*ListStateConfigurationsResponseBodyStateConfigurations) *ListStateConfigurationsResponseBody {
	s.StateConfigurations = v
	return s
}

type ListStateConfigurationsResponseBodyStateConfigurations struct {
	// The configuration mode. Valid values:
	//
	// example:
	//
	// ApplyAndAutoCorrect
	ConfigureMode *string `json:"ConfigureMode,omitempty" xml:"ConfigureMode,omitempty"`
	// The time when the desired-state configuration was created.
	//
	// example:
	//
	// 2021-03-22T03:13:32Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description.
	//
	// example:
	//
	// Collect inventory data
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The parameters.
	//
	// example:
	//
	// {"policy": {"ACS:Network": {"Collection": "Enabled"}, "ACS:Application": {"Collection": "Enabled"}}}
	Parameters *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfmxsn4m4******
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The schedule expression.
	//
	// example:
	//
	// 1 hour
	ScheduleExpression *string `json:"ScheduleExpression,omitempty" xml:"ScheduleExpression,omitempty"`
	// The schedule type.
	//
	// example:
	//
	// rate
	ScheduleType *string `json:"ScheduleType,omitempty" xml:"ScheduleType,omitempty"`
	// The ID of the desired-state configuration.
	//
	// example:
	//
	// sc-a538febe18fabcdef
	StateConfigurationId *string `json:"StateConfigurationId,omitempty" xml:"StateConfigurationId,omitempty"`
	// The tags added to the configuration.
	//
	// example:
	//
	// {"Key": "oos", "Value": "inventory"}
	Tags map[string]interface{} `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The queried resources.
	//
	// example:
	//
	// { "ResourceType": "ALIYUN::ECS::Instance", "Filters": [ { "Type": "All", "RegionId": "cn-hangzhou", "Parameters": { "RegionId": "cn-hangzhou", "Status": "Running" } } ] }
	Targets *string `json:"Targets,omitempty" xml:"Targets,omitempty"`
	// The template ID.
	//
	// example:
	//
	// t-ajshjalscfhjk2214
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The name of the template.
	//
	// example:
	//
	// ACS-ECS-InventoryDataCollection
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// The version of the template.
	//
	// example:
	//
	// v2
	TemplateVersion *string `json:"TemplateVersion,omitempty" xml:"TemplateVersion,omitempty"`
	// The time when the configuration was updated.
	//
	// example:
	//
	// 2021-04-22T03:13:32Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListStateConfigurationsResponseBodyStateConfigurations) String() string {
	return tea.Prettify(s)
}

func (s ListStateConfigurationsResponseBodyStateConfigurations) GoString() string {
	return s.String()
}

func (s *ListStateConfigurationsResponseBodyStateConfigurations) SetConfigureMode(v string) *ListStateConfigurationsResponseBodyStateConfigurations {
	s.ConfigureMode = &v
	return s
}

func (s *ListStateConfigurationsResponseBodyStateConfigurations) SetCreateTime(v string) *ListStateConfigurationsResponseBodyStateConfigurations {
	s.CreateTime = &v
	return s
}

func (s *ListStateConfigurationsResponseBodyStateConfigurations) SetDescription(v string) *ListStateConfigurationsResponseBodyStateConfigurations {
	s.Description = &v
	return s
}

func (s *ListStateConfigurationsResponseBodyStateConfigurations) SetParameters(v string) *ListStateConfigurationsResponseBodyStateConfigurations {
	s.Parameters = &v
	return s
}

func (s *ListStateConfigurationsResponseBodyStateConfigurations) SetResourceGroupId(v string) *ListStateConfigurationsResponseBodyStateConfigurations {
	s.ResourceGroupId = &v
	return s
}

func (s *ListStateConfigurationsResponseBodyStateConfigurations) SetScheduleExpression(v string) *ListStateConfigurationsResponseBodyStateConfigurations {
	s.ScheduleExpression = &v
	return s
}

func (s *ListStateConfigurationsResponseBodyStateConfigurations) SetScheduleType(v string) *ListStateConfigurationsResponseBodyStateConfigurations {
	s.ScheduleType = &v
	return s
}

func (s *ListStateConfigurationsResponseBodyStateConfigurations) SetStateConfigurationId(v string) *ListStateConfigurationsResponseBodyStateConfigurations {
	s.StateConfigurationId = &v
	return s
}

func (s *ListStateConfigurationsResponseBodyStateConfigurations) SetTags(v map[string]interface{}) *ListStateConfigurationsResponseBodyStateConfigurations {
	s.Tags = v
	return s
}

func (s *ListStateConfigurationsResponseBodyStateConfigurations) SetTargets(v string) *ListStateConfigurationsResponseBodyStateConfigurations {
	s.Targets = &v
	return s
}

func (s *ListStateConfigurationsResponseBodyStateConfigurations) SetTemplateId(v string) *ListStateConfigurationsResponseBodyStateConfigurations {
	s.TemplateId = &v
	return s
}

func (s *ListStateConfigurationsResponseBodyStateConfigurations) SetTemplateName(v string) *ListStateConfigurationsResponseBodyStateConfigurations {
	s.TemplateName = &v
	return s
}

func (s *ListStateConfigurationsResponseBodyStateConfigurations) SetTemplateVersion(v string) *ListStateConfigurationsResponseBodyStateConfigurations {
	s.TemplateVersion = &v
	return s
}

func (s *ListStateConfigurationsResponseBodyStateConfigurations) SetUpdateTime(v string) *ListStateConfigurationsResponseBodyStateConfigurations {
	s.UpdateTime = &v
	return s
}

type ListStateConfigurationsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListStateConfigurationsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListStateConfigurationsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListStateConfigurationsResponse) GoString() string {
	return s.String()
}

func (s *ListStateConfigurationsResponse) SetHeaders(v map[string]*string) *ListStateConfigurationsResponse {
	s.Headers = v
	return s
}

func (s *ListStateConfigurationsResponse) SetStatusCode(v int32) *ListStateConfigurationsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListStateConfigurationsResponse) SetBody(v *ListStateConfigurationsResponseBody) *ListStateConfigurationsResponse {
	s.Body = v
	return s
}

type ListTagKeysRequest struct {
	// The maximum number of entries to return on each page. Valid value: 10 to 100. Default value: 50.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that is used to retrieve the next page.
	//
	// example:
	//
	// djsdlkasd
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the region.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The type of the resource to which the tag is added.
	//
	// example:
	//
	// template
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s ListTagKeysRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTagKeysRequest) GoString() string {
	return s.String()
}

func (s *ListTagKeysRequest) SetMaxResults(v int32) *ListTagKeysRequest {
	s.MaxResults = &v
	return s
}

func (s *ListTagKeysRequest) SetNextToken(v string) *ListTagKeysRequest {
	s.NextToken = &v
	return s
}

func (s *ListTagKeysRequest) SetRegionId(v string) *ListTagKeysRequest {
	s.RegionId = &v
	return s
}

func (s *ListTagKeysRequest) SetResourceType(v string) *ListTagKeysRequest {
	s.ResourceType = &v
	return s
}

type ListTagKeysResponseBody struct {
	// The tag keys.
	Keys []*string `json:"Keys,omitempty" xml:"Keys,omitempty" type:"Repeated"`
	// The maximum number of entries to return on each page.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that is used to retrieve the next page.
	//
	// example:
	//
	// 87y29h80h20h3f2
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 36210B73-8262-4D08-9D3A-7F96789733C8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListTagKeysResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTagKeysResponseBody) GoString() string {
	return s.String()
}

func (s *ListTagKeysResponseBody) SetKeys(v []*string) *ListTagKeysResponseBody {
	s.Keys = v
	return s
}

func (s *ListTagKeysResponseBody) SetMaxResults(v int32) *ListTagKeysResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListTagKeysResponseBody) SetNextToken(v string) *ListTagKeysResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListTagKeysResponseBody) SetRequestId(v string) *ListTagKeysResponseBody {
	s.RequestId = &v
	return s
}

type ListTagKeysResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTagKeysResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTagKeysResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTagKeysResponse) GoString() string {
	return s.String()
}

func (s *ListTagKeysResponse) SetHeaders(v map[string]*string) *ListTagKeysResponse {
	s.Headers = v
	return s
}

func (s *ListTagKeysResponse) SetStatusCode(v int32) *ListTagKeysResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTagKeysResponse) SetBody(v *ListTagKeysResponseBody) *ListTagKeysResponse {
	s.Body = v
	return s
}

type ListTagResourcesRequest struct {
	// The pagination token that is used in the next request to retrieve a new page of results. You do not need to specify this parameter for the first request.
	//
	// example:
	//
	// MTRBMDc0NjAtRUJFNy00N0NBLTk3NTctMTJDQzQ
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the region.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The IDs of resources. The number of resource IDs ranges from 1 to 50.
	//
	// example:
	//
	// ["templateNam1","templateName2"]
	ResourceIds map[string]interface{} `json:"ResourceIds,omitempty" xml:"ResourceIds,omitempty"`
	// The type of the resource. Valid values: template execution
	//
	// This parameter is required.
	//
	// example:
	//
	// template
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The tag keys and values. The number of key-value pairs ranges from 1 to 20.
	//
	// example:
	//
	// {"k1":"v2","k2":"v2"}
	Tags map[string]interface{} `json:"Tags,omitempty" xml:"Tags,omitempty"`
}

func (s ListTagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesRequest) GoString() string {
	return s.String()
}

func (s *ListTagResourcesRequest) SetNextToken(v string) *ListTagResourcesRequest {
	s.NextToken = &v
	return s
}

func (s *ListTagResourcesRequest) SetRegionId(v string) *ListTagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *ListTagResourcesRequest) SetResourceIds(v map[string]interface{}) *ListTagResourcesRequest {
	s.ResourceIds = v
	return s
}

func (s *ListTagResourcesRequest) SetResourceType(v string) *ListTagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *ListTagResourcesRequest) SetTags(v map[string]interface{}) *ListTagResourcesRequest {
	s.Tags = v
	return s
}

type ListTagResourcesShrinkRequest struct {
	// The pagination token that is used in the next request to retrieve a new page of results. You do not need to specify this parameter for the first request.
	//
	// example:
	//
	// MTRBMDc0NjAtRUJFNy00N0NBLTk3NTctMTJDQzQ
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the region.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The IDs of resources. The number of resource IDs ranges from 1 to 50.
	//
	// example:
	//
	// ["templateNam1","templateName2"]
	ResourceIdsShrink *string `json:"ResourceIds,omitempty" xml:"ResourceIds,omitempty"`
	// The type of the resource. Valid values: template execution
	//
	// This parameter is required.
	//
	// example:
	//
	// template
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The tag keys and values. The number of key-value pairs ranges from 1 to 20.
	//
	// example:
	//
	// {"k1":"v2","k2":"v2"}
	TagsShrink *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
}

func (s ListTagResourcesShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListTagResourcesShrinkRequest) SetNextToken(v string) *ListTagResourcesShrinkRequest {
	s.NextToken = &v
	return s
}

func (s *ListTagResourcesShrinkRequest) SetRegionId(v string) *ListTagResourcesShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *ListTagResourcesShrinkRequest) SetResourceIdsShrink(v string) *ListTagResourcesShrinkRequest {
	s.ResourceIdsShrink = &v
	return s
}

func (s *ListTagResourcesShrinkRequest) SetResourceType(v string) *ListTagResourcesShrinkRequest {
	s.ResourceType = &v
	return s
}

func (s *ListTagResourcesShrinkRequest) SetTagsShrink(v string) *ListTagResourcesShrinkRequest {
	s.TagsShrink = &v
	return s
}

type ListTagResourcesResponseBody struct {
	// The returned value of NextToken is a pagination token, which can be used in the next request to retrieve a new page of results. If the return value of the NextToken parameter is empty, the next page does not exist.
	//
	// example:
	//
	// MTRBMDc0NjAtRUJFNy00N0NBLTk3NTctMTJDQzQ
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// A5EF78C1-67FC-4E36-A6A8-7DF9C51726DF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The set of resources and the tags that are added to the resources.
	TagResources *ListTagResourcesResponseBodyTagResources `json:"TagResources,omitempty" xml:"TagResources,omitempty" type:"Struct"`
}

func (s ListTagResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBody) SetNextToken(v string) *ListTagResourcesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListTagResourcesResponseBody) SetRequestId(v string) *ListTagResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTagResourcesResponseBody) SetTagResources(v *ListTagResourcesResponseBodyTagResources) *ListTagResourcesResponseBody {
	s.TagResources = v
	return s
}

type ListTagResourcesResponseBodyTagResources struct {
	TagResource []*ListTagResourcesResponseBodyTagResourcesTagResource `json:"TagResource,omitempty" xml:"TagResource,omitempty" type:"Repeated"`
}

func (s ListTagResourcesResponseBodyTagResources) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponseBodyTagResources) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBodyTagResources) SetTagResource(v []*ListTagResourcesResponseBodyTagResourcesTagResource) *ListTagResourcesResponseBodyTagResources {
	s.TagResource = v
	return s
}

type ListTagResourcesResponseBodyTagResourcesTagResource struct {
	// The ID of the resource.
	//
	// example:
	//
	// TagTest2
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The type of the resource.
	//
	// example:
	//
	// template
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The tag key.
	//
	// example:
	//
	// k1
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The tag value.
	//
	// example:
	//
	// v1
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s ListTagResourcesResponseBodyTagResourcesTagResource) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponseBodyTagResourcesTagResource) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBodyTagResourcesTagResource) SetResourceId(v string) *ListTagResourcesResponseBodyTagResourcesTagResource {
	s.ResourceId = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResourcesTagResource) SetResourceType(v string) *ListTagResourcesResponseBodyTagResourcesTagResource {
	s.ResourceType = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResourcesTagResource) SetTagKey(v string) *ListTagResourcesResponseBodyTagResourcesTagResource {
	s.TagKey = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResourcesTagResource) SetTagValue(v string) *ListTagResourcesResponseBodyTagResourcesTagResource {
	s.TagValue = &v
	return s
}

type ListTagResourcesResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTagResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponse) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponse) SetHeaders(v map[string]*string) *ListTagResourcesResponse {
	s.Headers = v
	return s
}

func (s *ListTagResourcesResponse) SetStatusCode(v int32) *ListTagResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTagResourcesResponse) SetBody(v *ListTagResourcesResponseBody) *ListTagResourcesResponse {
	s.Body = v
	return s
}

type ListTagValuesRequest struct {
	// The tag key to query.
	//
	// example:
	//
	// k1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The maximum number of results on each page.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. You do not need to specify this parameter for the first request.
	//
	// example:
	//
	// 3272h923879hsaksad
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the region.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The type of the tagged resource.
	//
	// example:
	//
	// template
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s ListTagValuesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTagValuesRequest) GoString() string {
	return s.String()
}

func (s *ListTagValuesRequest) SetKey(v string) *ListTagValuesRequest {
	s.Key = &v
	return s
}

func (s *ListTagValuesRequest) SetMaxResults(v int32) *ListTagValuesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListTagValuesRequest) SetNextToken(v string) *ListTagValuesRequest {
	s.NextToken = &v
	return s
}

func (s *ListTagValuesRequest) SetRegionId(v string) *ListTagValuesRequest {
	s.RegionId = &v
	return s
}

func (s *ListTagValuesRequest) SetResourceType(v string) *ListTagValuesRequest {
	s.ResourceType = &v
	return s
}

type ListTagValuesResponseBody struct {
	// The maximum number of results on each page.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The returned value of NextToken is a pagination token, which can be used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// 83u29j2dj3dskds
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 65591133-1188-4935-B78F-20F72
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The tag values returned.
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s ListTagValuesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTagValuesResponseBody) GoString() string {
	return s.String()
}

func (s *ListTagValuesResponseBody) SetMaxResults(v int32) *ListTagValuesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListTagValuesResponseBody) SetNextToken(v string) *ListTagValuesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListTagValuesResponseBody) SetRequestId(v string) *ListTagValuesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTagValuesResponseBody) SetValues(v []*string) *ListTagValuesResponseBody {
	s.Values = v
	return s
}

type ListTagValuesResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTagValuesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTagValuesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTagValuesResponse) GoString() string {
	return s.String()
}

func (s *ListTagValuesResponse) SetHeaders(v map[string]*string) *ListTagValuesResponse {
	s.Headers = v
	return s
}

func (s *ListTagValuesResponse) SetStatusCode(v int32) *ListTagValuesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTagValuesResponse) SetBody(v *ListTagValuesResponseBody) *ListTagValuesResponse {
	s.Body = v
	return s
}

type ListTaskExecutionsRequest struct {
	// The execution ID of the task.
	//
	// example:
	//
	// 2019-05-16T10:26:14Z
	EndDateAfter *string `json:"EndDateAfter,omitempty" xml:"EndDateAfter,omitempty"`
	// Specifies to query task executions that stop running at or later than the specified time.
	//
	// example:
	//
	// 2019-05-16T10:26:14Z
	EndDateBefore *string `json:"EndDateBefore,omitempty" xml:"EndDateBefore,omitempty"`
	// The status of the execution. Valid values: Running, Started, Success, Failed, Waiting, Cancelled, Pending, and Skipped.
	//
	// example:
	//
	// exec-xxx
	ExecutionId *string `json:"ExecutionId,omitempty" xml:"ExecutionId,omitempty"`
	// The number of entries to return on each page. Valid values: 20 to 100. Default value: 50.
	//
	// example:
	//
	// false
	IncludeChildTaskExecution *bool `json:"IncludeChildTaskExecution,omitempty" xml:"IncludeChildTaskExecution,omitempty"`
	// The token that is used to retrieve the next page of results.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// Sorts the task executions to query. Valid values:
	//
	// 	- **StartDate**: specifies that the task executions are sorted based on the time when they are created. This is the default value.
	//
	// 	- **EndDate**: specifies that the task executions are sorted based on the time when the time when they stop running.
	//
	// 	- **Status**: specifies that the task executions are sorted based on their statuses.
	//
	// example:
	//
	// MTRBMDc0NjAtRUJFNy00N0NBLTk3NTctMTJDQzQ3NjFENDdB
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// Specifies whether to show the child nodes in the loop task. Default value: False.
	//
	// example:
	//
	// task-exec-xxx
	ParentTaskExecutionId *string `json:"ParentTaskExecutionId,omitempty" xml:"ParentTaskExecutionId,omitempty"`
	// The ID of the execution.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The order in which you want to sort the task executions to query. Valid values:
	//
	// 	- **Ascending**: ascending order.
	//
	// 	- **Descending**: descending order. This is the default value.
	//
	// example:
	//
	// StartDate
	SortField *string `json:"SortField,omitempty" xml:"SortField,omitempty"`
	// The token that is used to retrieve the next page of results.
	//
	// example:
	//
	// Ascending
	SortOrder *string `json:"SortOrder,omitempty" xml:"SortOrder,omitempty"`
	// Specifies to query task executions that stop running at or before the specified time.
	//
	// example:
	//
	// 2019-05-16T10:26:14Z
	StartDateAfter *string `json:"StartDateAfter,omitempty" xml:"StartDateAfter,omitempty"`
	// Specifies to query task executions that start to run at or later than the specified time.
	//
	// example:
	//
	// 2019-05-16T10:26:14Z
	StartDateBefore *string `json:"StartDateBefore,omitempty" xml:"StartDateBefore,omitempty"`
	// Specifies to query task executions that start to run at or before the specified time.
	//
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The execution ID of the parent node. In a loop task, set this parameter to the execution ID of the parent node.
	//
	// example:
	//
	// ACS::Sleep
	TaskAction *string `json:"TaskAction,omitempty" xml:"TaskAction,omitempty"`
	// The name of the task.
	//
	// example:
	//
	// task-exec-xxx
	TaskExecutionId *string `json:"TaskExecutionId,omitempty" xml:"TaskExecutionId,omitempty"`
	// The action of the task.
	//
	// example:
	//
	// describeInstance
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
}

func (s ListTaskExecutionsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTaskExecutionsRequest) GoString() string {
	return s.String()
}

func (s *ListTaskExecutionsRequest) SetEndDateAfter(v string) *ListTaskExecutionsRequest {
	s.EndDateAfter = &v
	return s
}

func (s *ListTaskExecutionsRequest) SetEndDateBefore(v string) *ListTaskExecutionsRequest {
	s.EndDateBefore = &v
	return s
}

func (s *ListTaskExecutionsRequest) SetExecutionId(v string) *ListTaskExecutionsRequest {
	s.ExecutionId = &v
	return s
}

func (s *ListTaskExecutionsRequest) SetIncludeChildTaskExecution(v bool) *ListTaskExecutionsRequest {
	s.IncludeChildTaskExecution = &v
	return s
}

func (s *ListTaskExecutionsRequest) SetMaxResults(v int32) *ListTaskExecutionsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListTaskExecutionsRequest) SetNextToken(v string) *ListTaskExecutionsRequest {
	s.NextToken = &v
	return s
}

func (s *ListTaskExecutionsRequest) SetParentTaskExecutionId(v string) *ListTaskExecutionsRequest {
	s.ParentTaskExecutionId = &v
	return s
}

func (s *ListTaskExecutionsRequest) SetRegionId(v string) *ListTaskExecutionsRequest {
	s.RegionId = &v
	return s
}

func (s *ListTaskExecutionsRequest) SetSortField(v string) *ListTaskExecutionsRequest {
	s.SortField = &v
	return s
}

func (s *ListTaskExecutionsRequest) SetSortOrder(v string) *ListTaskExecutionsRequest {
	s.SortOrder = &v
	return s
}

func (s *ListTaskExecutionsRequest) SetStartDateAfter(v string) *ListTaskExecutionsRequest {
	s.StartDateAfter = &v
	return s
}

func (s *ListTaskExecutionsRequest) SetStartDateBefore(v string) *ListTaskExecutionsRequest {
	s.StartDateBefore = &v
	return s
}

func (s *ListTaskExecutionsRequest) SetStatus(v string) *ListTaskExecutionsRequest {
	s.Status = &v
	return s
}

func (s *ListTaskExecutionsRequest) SetTaskAction(v string) *ListTaskExecutionsRequest {
	s.TaskAction = &v
	return s
}

func (s *ListTaskExecutionsRequest) SetTaskExecutionId(v string) *ListTaskExecutionsRequest {
	s.TaskExecutionId = &v
	return s
}

func (s *ListTaskExecutionsRequest) SetTaskName(v string) *ListTaskExecutionsRequest {
	s.TaskName = &v
	return s
}

type ListTaskExecutionsResponseBody struct {
	// The details of the task executions.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// MTRBMDc0NjAtRUJFNy00N0NBLTk3NTctMTJDQzQ3NjFENDdB
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The number of entries returned on each page.
	//
	// example:
	//
	// CDABABABAB-FC28-4D9C-8FB5-68DC6F0486FC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The execution ID of the child node.
	TaskExecutions []*ListTaskExecutionsResponseBodyTaskExecutions `json:"TaskExecutions,omitempty" xml:"TaskExecutions,omitempty" type:"Repeated"`
}

func (s ListTaskExecutionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTaskExecutionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListTaskExecutionsResponseBody) SetMaxResults(v int32) *ListTaskExecutionsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListTaskExecutionsResponseBody) SetNextToken(v string) *ListTaskExecutionsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListTaskExecutionsResponseBody) SetRequestId(v string) *ListTaskExecutionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTaskExecutionsResponseBody) SetTaskExecutions(v []*ListTaskExecutionsResponseBodyTaskExecutions) *ListTaskExecutionsResponseBody {
	s.TaskExecutions = v
	return s
}

type ListTaskExecutionsResponseBodyTaskExecutions struct {
	// The output of the execution.
	//
	// example:
	//
	// exec-xxx
	ChildExecutionId *string `json:"ChildExecutionId,omitempty" xml:"ChildExecutionId,omitempty"`
	// The ID of the execution.
	//
	// example:
	//
	// 2019-05-16T10:26:14Z
	CreateDate *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	// The execution ID of the parent node.
	//
	// example:
	//
	// 2019-05-16T10:26:14Z
	EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// The action of the task.
	//
	// example:
	//
	// exec-44d32b45d2a449e49899
	ExecutionId *string `json:"ExecutionId,omitempty" xml:"ExecutionId,omitempty"`
	// The Input parameters of the task execution.
	//
	// example:
	//
	// {                     "NotifyNote":""                 }
	ExtraData map[string]interface{} `json:"ExtraData,omitempty" xml:"ExtraData,omitempty"`
	// The ID of the template.
	//
	// example:
	//
	// {}
	Loop map[string]interface{} `json:"Loop,omitempty" xml:"Loop,omitempty"`
	// The status information of the task execution.
	//
	// example:
	//
	// 2
	LoopBatchNumber *int32 `json:"LoopBatchNumber,omitempty" xml:"LoopBatchNumber,omitempty"`
	// The time when the execution was created.
	//
	// example:
	//
	// i-1234566zxcvvb
	LoopItem *string `json:"LoopItem,omitempty" xml:"LoopItem,omitempty"`
	// The status of the task.
	//
	// example:
	//
	// { "InstanceId":"i-xxx" }
	Outputs *string `json:"Outputs,omitempty" xml:"Outputs,omitempty"`
	// The name of the task.
	//
	// example:
	//
	// task-exec-xxx
	ParentTaskExecutionId *string `json:"ParentTaskExecutionId,omitempty" xml:"ParentTaskExecutionId,omitempty"`
	// Queries task executions. Multiple methods are supported to filter task executions.
	//
	// example:
	//
	// { "Status":"Running" }
	Properties *string `json:"Properties,omitempty" xml:"Properties,omitempty"`
	// The elements in the loop task.
	//
	// example:
	//
	// 2019-05-16T10:26:14Z
	StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	// The time when the task execution stopped running.
	//
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The additional information.
	//
	// example:
	//
	// ""
	StatusMessage *string `json:"StatusMessage,omitempty" xml:"StatusMessage,omitempty"`
	// The execution ID of the task.
	//
	// example:
	//
	// ACS::Sleep
	TaskAction *string `json:"TaskAction,omitempty" xml:"TaskAction,omitempty"`
	// The time when the execution was last updated.
	//
	// example:
	//
	// task-exec-xxx
	TaskExecutionId *string `json:"TaskExecutionId,omitempty" xml:"TaskExecutionId,omitempty"`
	// The time when the execution started.
	//
	// example:
	//
	// describeInstance
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	// The number of times for which the loop task is run.
	//
	// example:
	//
	// xxx
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The configuration and statistics information of the loop task. This parameter is returned only for the parent node of the loop task.
	//
	// example:
	//
	// 2019-05-16T10:26:14Z
	UpdateDate *string `json:"UpdateDate,omitempty" xml:"UpdateDate,omitempty"`
}

func (s ListTaskExecutionsResponseBodyTaskExecutions) String() string {
	return tea.Prettify(s)
}

func (s ListTaskExecutionsResponseBodyTaskExecutions) GoString() string {
	return s.String()
}

func (s *ListTaskExecutionsResponseBodyTaskExecutions) SetChildExecutionId(v string) *ListTaskExecutionsResponseBodyTaskExecutions {
	s.ChildExecutionId = &v
	return s
}

func (s *ListTaskExecutionsResponseBodyTaskExecutions) SetCreateDate(v string) *ListTaskExecutionsResponseBodyTaskExecutions {
	s.CreateDate = &v
	return s
}

func (s *ListTaskExecutionsResponseBodyTaskExecutions) SetEndDate(v string) *ListTaskExecutionsResponseBodyTaskExecutions {
	s.EndDate = &v
	return s
}

func (s *ListTaskExecutionsResponseBodyTaskExecutions) SetExecutionId(v string) *ListTaskExecutionsResponseBodyTaskExecutions {
	s.ExecutionId = &v
	return s
}

func (s *ListTaskExecutionsResponseBodyTaskExecutions) SetExtraData(v map[string]interface{}) *ListTaskExecutionsResponseBodyTaskExecutions {
	s.ExtraData = v
	return s
}

func (s *ListTaskExecutionsResponseBodyTaskExecutions) SetLoop(v map[string]interface{}) *ListTaskExecutionsResponseBodyTaskExecutions {
	s.Loop = v
	return s
}

func (s *ListTaskExecutionsResponseBodyTaskExecutions) SetLoopBatchNumber(v int32) *ListTaskExecutionsResponseBodyTaskExecutions {
	s.LoopBatchNumber = &v
	return s
}

func (s *ListTaskExecutionsResponseBodyTaskExecutions) SetLoopItem(v string) *ListTaskExecutionsResponseBodyTaskExecutions {
	s.LoopItem = &v
	return s
}

func (s *ListTaskExecutionsResponseBodyTaskExecutions) SetOutputs(v string) *ListTaskExecutionsResponseBodyTaskExecutions {
	s.Outputs = &v
	return s
}

func (s *ListTaskExecutionsResponseBodyTaskExecutions) SetParentTaskExecutionId(v string) *ListTaskExecutionsResponseBodyTaskExecutions {
	s.ParentTaskExecutionId = &v
	return s
}

func (s *ListTaskExecutionsResponseBodyTaskExecutions) SetProperties(v string) *ListTaskExecutionsResponseBodyTaskExecutions {
	s.Properties = &v
	return s
}

func (s *ListTaskExecutionsResponseBodyTaskExecutions) SetStartDate(v string) *ListTaskExecutionsResponseBodyTaskExecutions {
	s.StartDate = &v
	return s
}

func (s *ListTaskExecutionsResponseBodyTaskExecutions) SetStatus(v string) *ListTaskExecutionsResponseBodyTaskExecutions {
	s.Status = &v
	return s
}

func (s *ListTaskExecutionsResponseBodyTaskExecutions) SetStatusMessage(v string) *ListTaskExecutionsResponseBodyTaskExecutions {
	s.StatusMessage = &v
	return s
}

func (s *ListTaskExecutionsResponseBodyTaskExecutions) SetTaskAction(v string) *ListTaskExecutionsResponseBodyTaskExecutions {
	s.TaskAction = &v
	return s
}

func (s *ListTaskExecutionsResponseBodyTaskExecutions) SetTaskExecutionId(v string) *ListTaskExecutionsResponseBodyTaskExecutions {
	s.TaskExecutionId = &v
	return s
}

func (s *ListTaskExecutionsResponseBodyTaskExecutions) SetTaskName(v string) *ListTaskExecutionsResponseBodyTaskExecutions {
	s.TaskName = &v
	return s
}

func (s *ListTaskExecutionsResponseBodyTaskExecutions) SetTemplateId(v string) *ListTaskExecutionsResponseBodyTaskExecutions {
	s.TemplateId = &v
	return s
}

func (s *ListTaskExecutionsResponseBodyTaskExecutions) SetUpdateDate(v string) *ListTaskExecutionsResponseBodyTaskExecutions {
	s.UpdateDate = &v
	return s
}

type ListTaskExecutionsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTaskExecutionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTaskExecutionsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTaskExecutionsResponse) GoString() string {
	return s.String()
}

func (s *ListTaskExecutionsResponse) SetHeaders(v map[string]*string) *ListTaskExecutionsResponse {
	s.Headers = v
	return s
}

func (s *ListTaskExecutionsResponse) SetStatusCode(v int32) *ListTaskExecutionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTaskExecutionsResponse) SetBody(v *ListTaskExecutionsResponseBody) *ListTaskExecutionsResponse {
	s.Body = v
	return s
}

type ListTemplateVersionsRequest struct {
	// The number of entries per page. Valid values: 10 to 100
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// A pagination token. It can be used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// H8xj9c-398djs9-39ajd9asdjjJ
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The type of the template. Valid values: Private and Public.
	//
	// example:
	//
	// Private
	ShareType *string `json:"ShareType,omitempty" xml:"ShareType,omitempty"`
	// The name of the template.
	//
	// This parameter is required.
	//
	// example:
	//
	// describeinstances
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
}

func (s ListTemplateVersionsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTemplateVersionsRequest) GoString() string {
	return s.String()
}

func (s *ListTemplateVersionsRequest) SetMaxResults(v int32) *ListTemplateVersionsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListTemplateVersionsRequest) SetNextToken(v string) *ListTemplateVersionsRequest {
	s.NextToken = &v
	return s
}

func (s *ListTemplateVersionsRequest) SetRegionId(v string) *ListTemplateVersionsRequest {
	s.RegionId = &v
	return s
}

func (s *ListTemplateVersionsRequest) SetShareType(v string) *ListTemplateVersionsRequest {
	s.ShareType = &v
	return s
}

func (s *ListTemplateVersionsRequest) SetTemplateName(v string) *ListTemplateVersionsRequest {
	s.TemplateName = &v
	return s
}

type ListTemplateVersionsResponseBody struct {
	// The number of entries per page.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// A pagination token. It can be used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// NJSNDKLJS-SJKJDO090k30-JSDK232
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// E6CD612B-5889-4F1A-823F-8A4029E46
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The versions of the template.
	TemplateVersions []*ListTemplateVersionsResponseBodyTemplateVersions `json:"TemplateVersions,omitempty" xml:"TemplateVersions,omitempty" type:"Repeated"`
}

func (s ListTemplateVersionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTemplateVersionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListTemplateVersionsResponseBody) SetMaxResults(v int32) *ListTemplateVersionsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListTemplateVersionsResponseBody) SetNextToken(v string) *ListTemplateVersionsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListTemplateVersionsResponseBody) SetRequestId(v string) *ListTemplateVersionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTemplateVersionsResponseBody) SetTemplateVersions(v []*ListTemplateVersionsResponseBodyTemplateVersions) *ListTemplateVersionsResponseBody {
	s.TemplateVersions = v
	return s
}

type ListTemplateVersionsResponseBodyTemplateVersions struct {
	// The description of the version.
	//
	// example:
	//
	// Detach the eip from the special instance.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The format of the template content. Valid values: YAML and JSON.
	//
	// example:
	//
	// YAML
	TemplateFormat *string `json:"TemplateFormat,omitempty" xml:"TemplateFormat,omitempty"`
	// The number of the version.
	//
	// example:
	//
	// v2
	TemplateVersion *string `json:"TemplateVersion,omitempty" xml:"TemplateVersion,omitempty"`
	// The user who last updated the version.
	//
	// example:
	//
	// foo
	UpdatedBy *string `json:"UpdatedBy,omitempty" xml:"UpdatedBy,omitempty"`
	// The time when the version was last updated.
	//
	// example:
	//
	// 2020-05-19T06:05:41Z
	UpdatedDate *string `json:"UpdatedDate,omitempty" xml:"UpdatedDate,omitempty"`
	// The name of the version.
	//
	// example:
	//
	// baz
	VersionName *string `json:"VersionName,omitempty" xml:"VersionName,omitempty"`
}

func (s ListTemplateVersionsResponseBodyTemplateVersions) String() string {
	return tea.Prettify(s)
}

func (s ListTemplateVersionsResponseBodyTemplateVersions) GoString() string {
	return s.String()
}

func (s *ListTemplateVersionsResponseBodyTemplateVersions) SetDescription(v string) *ListTemplateVersionsResponseBodyTemplateVersions {
	s.Description = &v
	return s
}

func (s *ListTemplateVersionsResponseBodyTemplateVersions) SetTemplateFormat(v string) *ListTemplateVersionsResponseBodyTemplateVersions {
	s.TemplateFormat = &v
	return s
}

func (s *ListTemplateVersionsResponseBodyTemplateVersions) SetTemplateVersion(v string) *ListTemplateVersionsResponseBodyTemplateVersions {
	s.TemplateVersion = &v
	return s
}

func (s *ListTemplateVersionsResponseBodyTemplateVersions) SetUpdatedBy(v string) *ListTemplateVersionsResponseBodyTemplateVersions {
	s.UpdatedBy = &v
	return s
}

func (s *ListTemplateVersionsResponseBodyTemplateVersions) SetUpdatedDate(v string) *ListTemplateVersionsResponseBodyTemplateVersions {
	s.UpdatedDate = &v
	return s
}

func (s *ListTemplateVersionsResponseBodyTemplateVersions) SetVersionName(v string) *ListTemplateVersionsResponseBodyTemplateVersions {
	s.VersionName = &v
	return s
}

type ListTemplateVersionsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTemplateVersionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTemplateVersionsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTemplateVersionsResponse) GoString() string {
	return s.String()
}

func (s *ListTemplateVersionsResponse) SetHeaders(v map[string]*string) *ListTemplateVersionsResponse {
	s.Headers = v
	return s
}

func (s *ListTemplateVersionsResponse) SetStatusCode(v int32) *ListTemplateVersionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTemplateVersionsResponse) SetBody(v *ListTemplateVersionsResponseBody) *ListTemplateVersionsResponse {
	s.Body = v
	return s
}

type ListTemplatesRequest struct {
	// The type of the template. Valid values include TimerTrigger, EventTrigger, AlarmTrigger, and Other.
	//
	// example:
	//
	// TimerTrigger
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The creator of the template.
	//
	// 	- To query the template provided by Alibaba Cloud, set this parameter to **ACS**.
	//
	// 	- To query the template created by a user, set this parameter to the **ID*	- of the template or the **name of the user*	- who creates the template.
	//
	// example:
	//
	// ACS
	CreatedBy *string `json:"CreatedBy,omitempty" xml:"CreatedBy,omitempty"`
	// Specifies to query the template that is created at or later than the specified time.
	//
	// The value must be in the YYYY-MM-DDThh:mm:ssZ format.
	//
	// example:
	//
	// 2019-05-16T10:26:14Z
	CreatedDateAfter *string `json:"CreatedDateAfter,omitempty" xml:"CreatedDateAfter,omitempty"`
	// Specifies to query the template that is created at or before the specified time.
	//
	// The value must be in the YYYY-MM-DDThh:mm::ssZ format.
	//
	// example:
	//
	// 2019-05-16T10:26:14Z
	CreatedDateBefore *string `json:"CreatedDateBefore,omitempty" xml:"CreatedDateBefore,omitempty"`
	// Specifies whether to query the template that is configured with a trigger.
	//
	// example:
	//
	// true
	HasTrigger *bool `json:"HasTrigger,omitempty" xml:"HasTrigger,omitempty"`
	// Specifies whether the template is an example template
	//
	// example:
	//
	// false
	IsExample *bool `json:"IsExample,omitempty" xml:"IsExample,omitempty"`
	// Specifies whether the template is added to favorites.
	//
	// example:
	//
	// true
	IsFavorite *bool `json:"IsFavorite,omitempty" xml:"IsFavorite,omitempty"`
	// The number of entries per page. Valid values: 10 to 100. Default value: 50.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that is used to retrieve the next page of results.
	//
	// example:
	//
	// xxx
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the region in which you want to query templates.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfmxsn4m4******
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The share type of the template. Valid values:
	//
	// 	- **Public**
	//
	// 	- **Private**
	//
	// example:
	//
	// Private
	ShareType *string `json:"ShareType,omitempty" xml:"ShareType,omitempty"`
	// The field that is used to sort the templates to be queried. Valid values:
	//
	// 	- **TotalExecutionCount*	- (default): The system sorts the returned templates based on the total number of times that the templates are used.
	//
	// 	- **Popularity**: The system sorts the returned templates based on the popularity of the templates.
	//
	// 	- **TemplateName**: The system sorts the returned templates based on the names of the templates.
	//
	// 	- **CreatedDate**: The system sorts the returned templates based on the points in time when the templates are created.
	//
	// 	- **UpdatedDate**: The system sorts the returned templates based on the points in time when the templates are updated.
	//
	// example:
	//
	// Popularity
	SortField *string `json:"SortField,omitempty" xml:"SortField,omitempty"`
	// The order in which you want to sort the results. Valid values:
	//
	// 	- **Ascending**: ascending order.
	//
	// 	- **Descending**: descending order. This is the default value.
	//
	// example:
	//
	// Descending
	SortOrder *string `json:"SortOrder,omitempty" xml:"SortOrder,omitempty"`
	// The tag keys and values. The number of key-value pairs ranges from 1 to 20.
	//
	// example:
	//
	// {"k1":"k2","k2":"v2"}
	Tags map[string]interface{} `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The format of the template. Valid values:
	//
	// 	- **JSON**
	//
	// 	- **YAML**
	//
	// example:
	//
	// YAML
	TemplateFormat *string `json:"TemplateFormat,omitempty" xml:"TemplateFormat,omitempty"`
	// The name of the template. All templates whose names contain the specified template name are to be returned.
	//
	// example:
	//
	// MyTemplate
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// The type of the template. Valid values:
	//
	// 	- Automation: the template for automated tasks.
	//
	// 	- State: the template for configuration inventories.
	//
	// 	- Package: the template for software packages.
	//
	// example:
	//
	// private
	TemplateType *string `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
}

func (s ListTemplatesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTemplatesRequest) GoString() string {
	return s.String()
}

func (s *ListTemplatesRequest) SetCategory(v string) *ListTemplatesRequest {
	s.Category = &v
	return s
}

func (s *ListTemplatesRequest) SetCreatedBy(v string) *ListTemplatesRequest {
	s.CreatedBy = &v
	return s
}

func (s *ListTemplatesRequest) SetCreatedDateAfter(v string) *ListTemplatesRequest {
	s.CreatedDateAfter = &v
	return s
}

func (s *ListTemplatesRequest) SetCreatedDateBefore(v string) *ListTemplatesRequest {
	s.CreatedDateBefore = &v
	return s
}

func (s *ListTemplatesRequest) SetHasTrigger(v bool) *ListTemplatesRequest {
	s.HasTrigger = &v
	return s
}

func (s *ListTemplatesRequest) SetIsExample(v bool) *ListTemplatesRequest {
	s.IsExample = &v
	return s
}

func (s *ListTemplatesRequest) SetIsFavorite(v bool) *ListTemplatesRequest {
	s.IsFavorite = &v
	return s
}

func (s *ListTemplatesRequest) SetMaxResults(v int32) *ListTemplatesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListTemplatesRequest) SetNextToken(v string) *ListTemplatesRequest {
	s.NextToken = &v
	return s
}

func (s *ListTemplatesRequest) SetRegionId(v string) *ListTemplatesRequest {
	s.RegionId = &v
	return s
}

func (s *ListTemplatesRequest) SetResourceGroupId(v string) *ListTemplatesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListTemplatesRequest) SetShareType(v string) *ListTemplatesRequest {
	s.ShareType = &v
	return s
}

func (s *ListTemplatesRequest) SetSortField(v string) *ListTemplatesRequest {
	s.SortField = &v
	return s
}

func (s *ListTemplatesRequest) SetSortOrder(v string) *ListTemplatesRequest {
	s.SortOrder = &v
	return s
}

func (s *ListTemplatesRequest) SetTags(v map[string]interface{}) *ListTemplatesRequest {
	s.Tags = v
	return s
}

func (s *ListTemplatesRequest) SetTemplateFormat(v string) *ListTemplatesRequest {
	s.TemplateFormat = &v
	return s
}

func (s *ListTemplatesRequest) SetTemplateName(v string) *ListTemplatesRequest {
	s.TemplateName = &v
	return s
}

func (s *ListTemplatesRequest) SetTemplateType(v string) *ListTemplatesRequest {
	s.TemplateType = &v
	return s
}

type ListTemplatesShrinkRequest struct {
	// The type of the template. Valid values include TimerTrigger, EventTrigger, AlarmTrigger, and Other.
	//
	// example:
	//
	// TimerTrigger
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The creator of the template.
	//
	// 	- To query the template provided by Alibaba Cloud, set this parameter to **ACS**.
	//
	// 	- To query the template created by a user, set this parameter to the **ID*	- of the template or the **name of the user*	- who creates the template.
	//
	// example:
	//
	// ACS
	CreatedBy *string `json:"CreatedBy,omitempty" xml:"CreatedBy,omitempty"`
	// Specifies to query the template that is created at or later than the specified time.
	//
	// The value must be in the YYYY-MM-DDThh:mm:ssZ format.
	//
	// example:
	//
	// 2019-05-16T10:26:14Z
	CreatedDateAfter *string `json:"CreatedDateAfter,omitempty" xml:"CreatedDateAfter,omitempty"`
	// Specifies to query the template that is created at or before the specified time.
	//
	// The value must be in the YYYY-MM-DDThh:mm::ssZ format.
	//
	// example:
	//
	// 2019-05-16T10:26:14Z
	CreatedDateBefore *string `json:"CreatedDateBefore,omitempty" xml:"CreatedDateBefore,omitempty"`
	// Specifies whether to query the template that is configured with a trigger.
	//
	// example:
	//
	// true
	HasTrigger *bool `json:"HasTrigger,omitempty" xml:"HasTrigger,omitempty"`
	// Specifies whether the template is an example template
	//
	// example:
	//
	// false
	IsExample *bool `json:"IsExample,omitempty" xml:"IsExample,omitempty"`
	// Specifies whether the template is added to favorites.
	//
	// example:
	//
	// true
	IsFavorite *bool `json:"IsFavorite,omitempty" xml:"IsFavorite,omitempty"`
	// The number of entries per page. Valid values: 10 to 100. Default value: 50.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that is used to retrieve the next page of results.
	//
	// example:
	//
	// xxx
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the region in which you want to query templates.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfmxsn4m4******
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The share type of the template. Valid values:
	//
	// 	- **Public**
	//
	// 	- **Private**
	//
	// example:
	//
	// Private
	ShareType *string `json:"ShareType,omitempty" xml:"ShareType,omitempty"`
	// The field that is used to sort the templates to be queried. Valid values:
	//
	// 	- **TotalExecutionCount*	- (default): The system sorts the returned templates based on the total number of times that the templates are used.
	//
	// 	- **Popularity**: The system sorts the returned templates based on the popularity of the templates.
	//
	// 	- **TemplateName**: The system sorts the returned templates based on the names of the templates.
	//
	// 	- **CreatedDate**: The system sorts the returned templates based on the points in time when the templates are created.
	//
	// 	- **UpdatedDate**: The system sorts the returned templates based on the points in time when the templates are updated.
	//
	// example:
	//
	// Popularity
	SortField *string `json:"SortField,omitempty" xml:"SortField,omitempty"`
	// The order in which you want to sort the results. Valid values:
	//
	// 	- **Ascending**: ascending order.
	//
	// 	- **Descending**: descending order. This is the default value.
	//
	// example:
	//
	// Descending
	SortOrder *string `json:"SortOrder,omitempty" xml:"SortOrder,omitempty"`
	// The tag keys and values. The number of key-value pairs ranges from 1 to 20.
	//
	// example:
	//
	// {"k1":"k2","k2":"v2"}
	TagsShrink *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The format of the template. Valid values:
	//
	// 	- **JSON**
	//
	// 	- **YAML**
	//
	// example:
	//
	// YAML
	TemplateFormat *string `json:"TemplateFormat,omitempty" xml:"TemplateFormat,omitempty"`
	// The name of the template. All templates whose names contain the specified template name are to be returned.
	//
	// example:
	//
	// MyTemplate
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// The type of the template. Valid values:
	//
	// 	- Automation: the template for automated tasks.
	//
	// 	- State: the template for configuration inventories.
	//
	// 	- Package: the template for software packages.
	//
	// example:
	//
	// private
	TemplateType *string `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
}

func (s ListTemplatesShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTemplatesShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListTemplatesShrinkRequest) SetCategory(v string) *ListTemplatesShrinkRequest {
	s.Category = &v
	return s
}

func (s *ListTemplatesShrinkRequest) SetCreatedBy(v string) *ListTemplatesShrinkRequest {
	s.CreatedBy = &v
	return s
}

func (s *ListTemplatesShrinkRequest) SetCreatedDateAfter(v string) *ListTemplatesShrinkRequest {
	s.CreatedDateAfter = &v
	return s
}

func (s *ListTemplatesShrinkRequest) SetCreatedDateBefore(v string) *ListTemplatesShrinkRequest {
	s.CreatedDateBefore = &v
	return s
}

func (s *ListTemplatesShrinkRequest) SetHasTrigger(v bool) *ListTemplatesShrinkRequest {
	s.HasTrigger = &v
	return s
}

func (s *ListTemplatesShrinkRequest) SetIsExample(v bool) *ListTemplatesShrinkRequest {
	s.IsExample = &v
	return s
}

func (s *ListTemplatesShrinkRequest) SetIsFavorite(v bool) *ListTemplatesShrinkRequest {
	s.IsFavorite = &v
	return s
}

func (s *ListTemplatesShrinkRequest) SetMaxResults(v int32) *ListTemplatesShrinkRequest {
	s.MaxResults = &v
	return s
}

func (s *ListTemplatesShrinkRequest) SetNextToken(v string) *ListTemplatesShrinkRequest {
	s.NextToken = &v
	return s
}

func (s *ListTemplatesShrinkRequest) SetRegionId(v string) *ListTemplatesShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *ListTemplatesShrinkRequest) SetResourceGroupId(v string) *ListTemplatesShrinkRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListTemplatesShrinkRequest) SetShareType(v string) *ListTemplatesShrinkRequest {
	s.ShareType = &v
	return s
}

func (s *ListTemplatesShrinkRequest) SetSortField(v string) *ListTemplatesShrinkRequest {
	s.SortField = &v
	return s
}

func (s *ListTemplatesShrinkRequest) SetSortOrder(v string) *ListTemplatesShrinkRequest {
	s.SortOrder = &v
	return s
}

func (s *ListTemplatesShrinkRequest) SetTagsShrink(v string) *ListTemplatesShrinkRequest {
	s.TagsShrink = &v
	return s
}

func (s *ListTemplatesShrinkRequest) SetTemplateFormat(v string) *ListTemplatesShrinkRequest {
	s.TemplateFormat = &v
	return s
}

func (s *ListTemplatesShrinkRequest) SetTemplateName(v string) *ListTemplatesShrinkRequest {
	s.TemplateName = &v
	return s
}

func (s *ListTemplatesShrinkRequest) SetTemplateType(v string) *ListTemplatesShrinkRequest {
	s.TemplateType = &v
	return s
}

type ListTemplatesResponseBody struct {
	// The number of entries returned on each page.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that is used to retrieve the next page of results.
	//
	// example:
	//
	// xxx
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// BEF54BA-17B6-449F-A219-49ACB157E3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The template metadata.
	Templates []*ListTemplatesResponseBodyTemplates `json:"Templates,omitempty" xml:"Templates,omitempty" type:"Repeated"`
}

func (s ListTemplatesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTemplatesResponseBody) GoString() string {
	return s.String()
}

func (s *ListTemplatesResponseBody) SetMaxResults(v int32) *ListTemplatesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListTemplatesResponseBody) SetNextToken(v string) *ListTemplatesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListTemplatesResponseBody) SetRequestId(v string) *ListTemplatesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTemplatesResponseBody) SetTemplates(v []*ListTemplatesResponseBodyTemplates) *ListTemplatesResponseBody {
	s.Templates = v
	return s
}

type ListTemplatesResponseBodyTemplates struct {
	// The template type.
	//
	// example:
	//
	// TimerTrigger
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The template constraints.
	//
	// example:
	//
	// {
	//
	//   "InstanceTypeFamilies": ["ecs.g8y", "ecs.c8y"],
	//
	//   "ImageTypes": ["system"],
	//
	//   "OSPlatforms": ["CentOS", "Ubuntu"],
	//
	//   "OSVersions": ["CentOS7.9 64bit"]
	//
	// }
	Constraints *string `json:"Constraints,omitempty" xml:"Constraints,omitempty"`
	// The user who created the template.
	//
	// example:
	//
	// root(1309200)
	CreatedBy *string `json:"CreatedBy,omitempty" xml:"CreatedBy,omitempty"`
	// The creation time of the template.
	//
	// example:
	//
	// 2019-05-16T10:26:14Z
	CreatedDate *string `json:"CreatedDate,omitempty" xml:"CreatedDate,omitempty"`
	// The template description.
	//
	// example:
	//
	// Describe instances of given status
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Indicates whether the template was configured with a trigger.
	//
	// example:
	//
	// true
	HasTrigger *bool `json:"HasTrigger,omitempty" xml:"HasTrigger,omitempty"`
	// The SHA256 value of the template content.
	//
	// example:
	//
	// 4bc7d7a21b3e003434b9c223f6e6d2578b5ebfeb5be28c1fcf8a8a1b11907bb4
	Hash *string `json:"Hash,omitempty" xml:"Hash,omitempty"`
	// Indicates whether the template is added to favorites.
	//
	// example:
	//
	// true
	IsFavorite *bool `json:"IsFavorite,omitempty" xml:"IsFavorite,omitempty"`
	// The popularity of the public template. Valid values: **1-10**. A greater value indicates higher popularity. If **ShareType*	- is set to **Private**, the value of this parameter is `-1`.
	//
	// >  This parameter is valid only if **ShareType*	- is set to **Public**.
	//
	// example:
	//
	// 8
	Popularity *int32 `json:"Popularity,omitempty" xml:"Popularity,omitempty"`
	// The user who published the template.
	//
	// example:
	//
	// aliyun
	Publisher *string `json:"Publisher,omitempty" xml:"Publisher,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfmxsn4m4******
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The share type of the template. The share type of a template created by a user is **Private**. Valid values:
	//
	// 	- **Public**
	//
	// 	- **Private**
	//
	// example:
	//
	// Public
	ShareType *string `json:"ShareType,omitempty" xml:"ShareType,omitempty"`
	// The tag keys and values. The number of key-value pairs ranges from 1 to 20.
	//
	// example:
	//
	// {"k1":"v1","k2":"v2"}
	Tags map[string]interface{} `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The template format. The system automatically determines whether the format of the template is JSON or YAML.
	//
	// example:
	//
	// JSON
	TemplateFormat *string `json:"TemplateFormat,omitempty" xml:"TemplateFormat,omitempty"`
	// The template ID.
	//
	// example:
	//
	// t-94753deed38
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The template name.
	//
	// example:
	//
	// MyTemplate
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// The template type.
	//
	// example:
	//
	// private
	TemplateType *string `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
	// The template version. The version contains the letter v and a number. The number starts from 1.
	//
	// example:
	//
	// v1
	TemplateVersion *string `json:"TemplateVersion,omitempty" xml:"TemplateVersion,omitempty"`
	// The number of times for which the private template is executed. If **ShareType*	- is set to **Public**, the value of this parameter is `-1`.
	//
	// >  This parameter is valid only if **ShareType*	- is set to **Private**.
	//
	// example:
	//
	// 5
	TotalExecutionCount *int32 `json:"TotalExecutionCount,omitempty" xml:"TotalExecutionCount,omitempty"`
	// The user who last updated the template.
	//
	// example:
	//
	// root(13092000)
	UpdatedBy *string `json:"UpdatedBy,omitempty" xml:"UpdatedBy,omitempty"`
	// The time when the template was last updated.
	//
	// example:
	//
	// 2019-05-16T10:26:14Z
	UpdatedDate *string `json:"UpdatedDate,omitempty" xml:"UpdatedDate,omitempty"`
}

func (s ListTemplatesResponseBodyTemplates) String() string {
	return tea.Prettify(s)
}

func (s ListTemplatesResponseBodyTemplates) GoString() string {
	return s.String()
}

func (s *ListTemplatesResponseBodyTemplates) SetCategory(v string) *ListTemplatesResponseBodyTemplates {
	s.Category = &v
	return s
}

func (s *ListTemplatesResponseBodyTemplates) SetConstraints(v string) *ListTemplatesResponseBodyTemplates {
	s.Constraints = &v
	return s
}

func (s *ListTemplatesResponseBodyTemplates) SetCreatedBy(v string) *ListTemplatesResponseBodyTemplates {
	s.CreatedBy = &v
	return s
}

func (s *ListTemplatesResponseBodyTemplates) SetCreatedDate(v string) *ListTemplatesResponseBodyTemplates {
	s.CreatedDate = &v
	return s
}

func (s *ListTemplatesResponseBodyTemplates) SetDescription(v string) *ListTemplatesResponseBodyTemplates {
	s.Description = &v
	return s
}

func (s *ListTemplatesResponseBodyTemplates) SetHasTrigger(v bool) *ListTemplatesResponseBodyTemplates {
	s.HasTrigger = &v
	return s
}

func (s *ListTemplatesResponseBodyTemplates) SetHash(v string) *ListTemplatesResponseBodyTemplates {
	s.Hash = &v
	return s
}

func (s *ListTemplatesResponseBodyTemplates) SetIsFavorite(v bool) *ListTemplatesResponseBodyTemplates {
	s.IsFavorite = &v
	return s
}

func (s *ListTemplatesResponseBodyTemplates) SetPopularity(v int32) *ListTemplatesResponseBodyTemplates {
	s.Popularity = &v
	return s
}

func (s *ListTemplatesResponseBodyTemplates) SetPublisher(v string) *ListTemplatesResponseBodyTemplates {
	s.Publisher = &v
	return s
}

func (s *ListTemplatesResponseBodyTemplates) SetResourceGroupId(v string) *ListTemplatesResponseBodyTemplates {
	s.ResourceGroupId = &v
	return s
}

func (s *ListTemplatesResponseBodyTemplates) SetShareType(v string) *ListTemplatesResponseBodyTemplates {
	s.ShareType = &v
	return s
}

func (s *ListTemplatesResponseBodyTemplates) SetTags(v map[string]interface{}) *ListTemplatesResponseBodyTemplates {
	s.Tags = v
	return s
}

func (s *ListTemplatesResponseBodyTemplates) SetTemplateFormat(v string) *ListTemplatesResponseBodyTemplates {
	s.TemplateFormat = &v
	return s
}

func (s *ListTemplatesResponseBodyTemplates) SetTemplateId(v string) *ListTemplatesResponseBodyTemplates {
	s.TemplateId = &v
	return s
}

func (s *ListTemplatesResponseBodyTemplates) SetTemplateName(v string) *ListTemplatesResponseBodyTemplates {
	s.TemplateName = &v
	return s
}

func (s *ListTemplatesResponseBodyTemplates) SetTemplateType(v string) *ListTemplatesResponseBodyTemplates {
	s.TemplateType = &v
	return s
}

func (s *ListTemplatesResponseBodyTemplates) SetTemplateVersion(v string) *ListTemplatesResponseBodyTemplates {
	s.TemplateVersion = &v
	return s
}

func (s *ListTemplatesResponseBodyTemplates) SetTotalExecutionCount(v int32) *ListTemplatesResponseBodyTemplates {
	s.TotalExecutionCount = &v
	return s
}

func (s *ListTemplatesResponseBodyTemplates) SetUpdatedBy(v string) *ListTemplatesResponseBodyTemplates {
	s.UpdatedBy = &v
	return s
}

func (s *ListTemplatesResponseBodyTemplates) SetUpdatedDate(v string) *ListTemplatesResponseBodyTemplates {
	s.UpdatedDate = &v
	return s
}

type ListTemplatesResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTemplatesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTemplatesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTemplatesResponse) GoString() string {
	return s.String()
}

func (s *ListTemplatesResponse) SetHeaders(v map[string]*string) *ListTemplatesResponse {
	s.Headers = v
	return s
}

func (s *ListTemplatesResponse) SetStatusCode(v int32) *ListTemplatesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTemplatesResponse) SetBody(v *ListTemplatesResponseBody) *ListTemplatesResponse {
	s.Body = v
	return s
}

type NotifyExecutionRequest struct {
	// The ID of the execution.
	//
	// This parameter is required.
	//
	// example:
	//
	// exec-xxx
	ExecutionId *string `json:"ExecutionId,omitempty" xml:"ExecutionId,omitempty"`
	// The state of the terminated execution. This parameter is valid if you set the NotifyType parameter to CompleteExecution.
	//
	// example:
	//
	// Success
	ExecutionStatus *string `json:"ExecutionStatus,omitempty" xml:"ExecutionStatus,omitempty"`
	// The items of the child node in the loop task.
	//
	// example:
	//
	// i-xxx
	LoopItem *string `json:"LoopItem,omitempty" xml:"LoopItem,omitempty"`
	// The description for the notification.
	//
	// example:
	//
	// Note
	NotifyNote *string `json:"NotifyNote,omitempty" xml:"NotifyNote,omitempty"`
	// The type of the notification. Valid values:
	//
	// 	- **ExecuteTask**: starts to run a specific task. This value is used if you perform debugging in the Debug mode. If you set this parameter to ExecuteTask, you also need to set the Parameters parameter.
	//
	// 	- **CancelTask**: cancels a current task. This value is used if you perform debugging in the Debug mode.
	//
	// 	- **CompleteExecution**: manually terminates an execution if you perform debugging in the Debug mode. You can specify the state of the terminated execution by using the **ExecutionStatus*	- parameter.
	//
	// 	- **Approve**: approves an execution. For example, you are aware of the risks of an operation task and agree to approve the execution.
	//
	// 	- **Reject**: rejects an execution. For example, you want to reject the execution of a high-risk operation task.
	//
	// 	- **RetryTask**: retries a failed task whose execution mode is Suspend upon Failure.
	//
	// 	- **SkipTask**: skips a failed task whose execution mode is Suspend upon Failure.
	//
	// This parameter is required.
	//
	// example:
	//
	// Approve
	NotifyType *string `json:"NotifyType,omitempty" xml:"NotifyType,omitempty"`
	// The parameters of the subsequent task. This parameter is valid if you set the NotifyType parameter to ExecuteTask.
	//
	// example:
	//
	// {}
	Parameters *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// The ID of the region in which the execution resides.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The execution ID of the task.
	//
	// example:
	//
	// task-exec-xxx
	TaskExecutionId *string `json:"TaskExecutionId,omitempty" xml:"TaskExecutionId,omitempty"`
	// The execution IDs of the tasks.
	//
	// example:
	//
	// ["exec-79c321c11003a97c","exec-79c321c11003aqw97cz"]
	TaskExecutionIds *string `json:"TaskExecutionIds,omitempty" xml:"TaskExecutionIds,omitempty"`
	// The name of the subsequent task.
	//
	// example:
	//
	// describeInstance
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
}

func (s NotifyExecutionRequest) String() string {
	return tea.Prettify(s)
}

func (s NotifyExecutionRequest) GoString() string {
	return s.String()
}

func (s *NotifyExecutionRequest) SetExecutionId(v string) *NotifyExecutionRequest {
	s.ExecutionId = &v
	return s
}

func (s *NotifyExecutionRequest) SetExecutionStatus(v string) *NotifyExecutionRequest {
	s.ExecutionStatus = &v
	return s
}

func (s *NotifyExecutionRequest) SetLoopItem(v string) *NotifyExecutionRequest {
	s.LoopItem = &v
	return s
}

func (s *NotifyExecutionRequest) SetNotifyNote(v string) *NotifyExecutionRequest {
	s.NotifyNote = &v
	return s
}

func (s *NotifyExecutionRequest) SetNotifyType(v string) *NotifyExecutionRequest {
	s.NotifyType = &v
	return s
}

func (s *NotifyExecutionRequest) SetParameters(v string) *NotifyExecutionRequest {
	s.Parameters = &v
	return s
}

func (s *NotifyExecutionRequest) SetRegionId(v string) *NotifyExecutionRequest {
	s.RegionId = &v
	return s
}

func (s *NotifyExecutionRequest) SetTaskExecutionId(v string) *NotifyExecutionRequest {
	s.TaskExecutionId = &v
	return s
}

func (s *NotifyExecutionRequest) SetTaskExecutionIds(v string) *NotifyExecutionRequest {
	s.TaskExecutionIds = &v
	return s
}

func (s *NotifyExecutionRequest) SetTaskName(v string) *NotifyExecutionRequest {
	s.TaskName = &v
	return s
}

type NotifyExecutionResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 491DF8C2-34C9-4679-9DB3-4C0F49B129AC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s NotifyExecutionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s NotifyExecutionResponseBody) GoString() string {
	return s.String()
}

func (s *NotifyExecutionResponseBody) SetRequestId(v string) *NotifyExecutionResponseBody {
	s.RequestId = &v
	return s
}

type NotifyExecutionResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *NotifyExecutionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s NotifyExecutionResponse) String() string {
	return tea.Prettify(s)
}

func (s NotifyExecutionResponse) GoString() string {
	return s.String()
}

func (s *NotifyExecutionResponse) SetHeaders(v map[string]*string) *NotifyExecutionResponse {
	s.Headers = v
	return s
}

func (s *NotifyExecutionResponse) SetStatusCode(v int32) *NotifyExecutionResponse {
	s.StatusCode = &v
	return s
}

func (s *NotifyExecutionResponse) SetBody(v *NotifyExecutionResponseBody) *NotifyExecutionResponse {
	s.Body = v
	return s
}

type RegisterDefaultPatchBaselineRequest struct {
	// The name of the patch baseline.
	//
	// This parameter is required.
	//
	// example:
	//
	// MyPatchBaseline
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the region.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s RegisterDefaultPatchBaselineRequest) String() string {
	return tea.Prettify(s)
}

func (s RegisterDefaultPatchBaselineRequest) GoString() string {
	return s.String()
}

func (s *RegisterDefaultPatchBaselineRequest) SetName(v string) *RegisterDefaultPatchBaselineRequest {
	s.Name = &v
	return s
}

func (s *RegisterDefaultPatchBaselineRequest) SetRegionId(v string) *RegisterDefaultPatchBaselineRequest {
	s.RegionId = &v
	return s
}

type RegisterDefaultPatchBaselineResponseBody struct {
	// The details of the patch baseline.
	PatchBaseline *RegisterDefaultPatchBaselineResponseBodyPatchBaseline `json:"PatchBaseline,omitempty" xml:"PatchBaseline,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// D6850689-348D-59FC-AE13-BB0EDB7C4BE8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RegisterDefaultPatchBaselineResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RegisterDefaultPatchBaselineResponseBody) GoString() string {
	return s.String()
}

func (s *RegisterDefaultPatchBaselineResponseBody) SetPatchBaseline(v *RegisterDefaultPatchBaselineResponseBodyPatchBaseline) *RegisterDefaultPatchBaselineResponseBody {
	s.PatchBaseline = v
	return s
}

func (s *RegisterDefaultPatchBaselineResponseBody) SetRequestId(v string) *RegisterDefaultPatchBaselineResponseBody {
	s.RequestId = &v
	return s
}

type RegisterDefaultPatchBaselineResponseBodyPatchBaseline struct {
	// The rules of scanning and installing patches for the specified operating system.
	//
	// example:
	//
	// {"PatchRules":[{"PatchFilterGroup":[{"Key":"PatchSet","Values":["OS"]},{"Key":"ProductFamily","Values":["Windows"]},{"Key":"Product","Values":["Windows 10","Windows 7"]},{"Key":"Classification","Values":["Security Updates","Updates","Update Rollups","Critical Updates"]},{"Key":"Severity","Values":["Critical","Important","Moderate"]}],"ApproveAfterDays":7,"ApproveUntilDate":"","EnableNonSecurity":true,"ComplianceLevel":"Medium"}]}
	ApprovalRules *string `json:"ApprovalRules,omitempty" xml:"ApprovalRules,omitempty"`
	// The user who created the patch baseline.
	//
	// example:
	//
	// root(130900000)
	CreatedBy *string `json:"CreatedBy,omitempty" xml:"CreatedBy,omitempty"`
	// The time when the patch baseline was created.
	//
	// example:
	//
	// 2021-09-07T03:42:56Z
	CreatedDate *string `json:"CreatedDate,omitempty" xml:"CreatedDate,omitempty"`
	// The description of the patch baseline.
	//
	// example:
	//
	// RegisterPatchBaseline
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the patch baseline.
	//
	// example:
	//
	// pb-445340b5c6504a85a300
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the patch baseline.
	//
	// example:
	//
	// MyPatchBaseline
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The operating system.
	//
	// example:
	//
	// Windows
	OperationSystem *string `json:"OperationSystem,omitempty" xml:"OperationSystem,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfm4dpaq2yox6q
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The share type of the patch baseline.
	//
	// example:
	//
	// Private
	ShareType *string `json:"ShareType,omitempty" xml:"ShareType,omitempty"`
	// The user who last updated the patch baseline.
	//
	// example:
	//
	// root(130900000)
	UpdatedBy *string `json:"UpdatedBy,omitempty" xml:"UpdatedBy,omitempty"`
	// The time when the patch baseline was last updated.
	//
	// example:
	//
	// 2021-09-07T03:42:56Z
	UpdatedDate *string `json:"UpdatedDate,omitempty" xml:"UpdatedDate,omitempty"`
}

func (s RegisterDefaultPatchBaselineResponseBodyPatchBaseline) String() string {
	return tea.Prettify(s)
}

func (s RegisterDefaultPatchBaselineResponseBodyPatchBaseline) GoString() string {
	return s.String()
}

func (s *RegisterDefaultPatchBaselineResponseBodyPatchBaseline) SetApprovalRules(v string) *RegisterDefaultPatchBaselineResponseBodyPatchBaseline {
	s.ApprovalRules = &v
	return s
}

func (s *RegisterDefaultPatchBaselineResponseBodyPatchBaseline) SetCreatedBy(v string) *RegisterDefaultPatchBaselineResponseBodyPatchBaseline {
	s.CreatedBy = &v
	return s
}

func (s *RegisterDefaultPatchBaselineResponseBodyPatchBaseline) SetCreatedDate(v string) *RegisterDefaultPatchBaselineResponseBodyPatchBaseline {
	s.CreatedDate = &v
	return s
}

func (s *RegisterDefaultPatchBaselineResponseBodyPatchBaseline) SetDescription(v string) *RegisterDefaultPatchBaselineResponseBodyPatchBaseline {
	s.Description = &v
	return s
}

func (s *RegisterDefaultPatchBaselineResponseBodyPatchBaseline) SetId(v string) *RegisterDefaultPatchBaselineResponseBodyPatchBaseline {
	s.Id = &v
	return s
}

func (s *RegisterDefaultPatchBaselineResponseBodyPatchBaseline) SetName(v string) *RegisterDefaultPatchBaselineResponseBodyPatchBaseline {
	s.Name = &v
	return s
}

func (s *RegisterDefaultPatchBaselineResponseBodyPatchBaseline) SetOperationSystem(v string) *RegisterDefaultPatchBaselineResponseBodyPatchBaseline {
	s.OperationSystem = &v
	return s
}

func (s *RegisterDefaultPatchBaselineResponseBodyPatchBaseline) SetResourceGroupId(v string) *RegisterDefaultPatchBaselineResponseBodyPatchBaseline {
	s.ResourceGroupId = &v
	return s
}

func (s *RegisterDefaultPatchBaselineResponseBodyPatchBaseline) SetShareType(v string) *RegisterDefaultPatchBaselineResponseBodyPatchBaseline {
	s.ShareType = &v
	return s
}

func (s *RegisterDefaultPatchBaselineResponseBodyPatchBaseline) SetUpdatedBy(v string) *RegisterDefaultPatchBaselineResponseBodyPatchBaseline {
	s.UpdatedBy = &v
	return s
}

func (s *RegisterDefaultPatchBaselineResponseBodyPatchBaseline) SetUpdatedDate(v string) *RegisterDefaultPatchBaselineResponseBodyPatchBaseline {
	s.UpdatedDate = &v
	return s
}

type RegisterDefaultPatchBaselineResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RegisterDefaultPatchBaselineResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RegisterDefaultPatchBaselineResponse) String() string {
	return tea.Prettify(s)
}

func (s RegisterDefaultPatchBaselineResponse) GoString() string {
	return s.String()
}

func (s *RegisterDefaultPatchBaselineResponse) SetHeaders(v map[string]*string) *RegisterDefaultPatchBaselineResponse {
	s.Headers = v
	return s
}

func (s *RegisterDefaultPatchBaselineResponse) SetStatusCode(v int32) *RegisterDefaultPatchBaselineResponse {
	s.StatusCode = &v
	return s
}

func (s *RegisterDefaultPatchBaselineResponse) SetBody(v *RegisterDefaultPatchBaselineResponseBody) *RegisterDefaultPatchBaselineResponse {
	s.Body = v
	return s
}

type SearchInventoryRequest struct {
	// The information about aggregators. You can use one or more aggregators to query the aggregate information of an instance. Valid values:
	//
	// 	- ACS:Application.Name
	//
	// 	- ACS:Application.Version
	//
	// example:
	//
	// ACS:Application.Name
	Aggregator []*string `json:"Aggregator,omitempty" xml:"Aggregator,omitempty" type:"Repeated"`
	// The filter rules for the component.
	Filter []*SearchInventoryRequestFilter `json:"Filter,omitempty" xml:"Filter,omitempty" type:"Repeated"`
	// The number of entries per page. Valid values: 1 to 100. Default value: 50.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that is used to retrieve the next page of results.
	//
	// example:
	//
	// gAAAAABfTgv5ewUWmNdJ3g7JVLvX70sPH90GZOVGC
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the region.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s SearchInventoryRequest) String() string {
	return tea.Prettify(s)
}

func (s SearchInventoryRequest) GoString() string {
	return s.String()
}

func (s *SearchInventoryRequest) SetAggregator(v []*string) *SearchInventoryRequest {
	s.Aggregator = v
	return s
}

func (s *SearchInventoryRequest) SetFilter(v []*SearchInventoryRequestFilter) *SearchInventoryRequest {
	s.Filter = v
	return s
}

func (s *SearchInventoryRequest) SetMaxResults(v int32) *SearchInventoryRequest {
	s.MaxResults = &v
	return s
}

func (s *SearchInventoryRequest) SetNextToken(v string) *SearchInventoryRequest {
	s.NextToken = &v
	return s
}

func (s *SearchInventoryRequest) SetRegionId(v string) *SearchInventoryRequest {
	s.RegionId = &v
	return s
}

type SearchInventoryRequestFilter struct {
	// The name of the component property. Valid values of N: 1 to 5. Different components have different property names. You can call the [GetInventorySchema](https://api.aliyun.com/#/?product=oos\\&version=2019-06-01\\&api=GetInventorySchema) operation to query the property names of different components. For example, the ACS:InstanceInformation component has the InstanceId property. Therefore, you can set this parameter to ACS:InstanceInformation.InstanceId.
	//
	// example:
	//
	// ACS:InstanceInformation.InstanceId
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The comparison operator that is used to filter property values. Valid values of N: 1 to 5. Valid values:
	//
	// 	- Equal
	//
	// 	- NotEqual
	//
	// 	- BeginWith
	//
	// 	- LessThan
	//
	// 	- GreaterThan
	//
	// example:
	//
	// Equal
	Operator *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	// The property values to query.
	//
	// example:
	//
	// i-bp1cpoxxxxxxxxxxxxxx
	Value []*string `json:"Value,omitempty" xml:"Value,omitempty" type:"Repeated"`
}

func (s SearchInventoryRequestFilter) String() string {
	return tea.Prettify(s)
}

func (s SearchInventoryRequestFilter) GoString() string {
	return s.String()
}

func (s *SearchInventoryRequestFilter) SetName(v string) *SearchInventoryRequestFilter {
	s.Name = &v
	return s
}

func (s *SearchInventoryRequestFilter) SetOperator(v string) *SearchInventoryRequestFilter {
	s.Operator = &v
	return s
}

func (s *SearchInventoryRequestFilter) SetValue(v []*string) *SearchInventoryRequestFilter {
	s.Value = v
	return s
}

type SearchInventoryResponseBody struct {
	Entities []map[string]interface{} `json:"Entities,omitempty" xml:"Entities,omitempty" type:"Repeated"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that is used to retrieve the next page of results.
	//
	// example:
	//
	// gAAAAABfTgv5ewUWmNdJ3g7JVLvX70sPH90GZOVGC
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// A81E4B2E-6B33-4BAE-9856-55DB7C893E01
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SearchInventoryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SearchInventoryResponseBody) GoString() string {
	return s.String()
}

func (s *SearchInventoryResponseBody) SetEntities(v []map[string]interface{}) *SearchInventoryResponseBody {
	s.Entities = v
	return s
}

func (s *SearchInventoryResponseBody) SetMaxResults(v int32) *SearchInventoryResponseBody {
	s.MaxResults = &v
	return s
}

func (s *SearchInventoryResponseBody) SetNextToken(v string) *SearchInventoryResponseBody {
	s.NextToken = &v
	return s
}

func (s *SearchInventoryResponseBody) SetRequestId(v string) *SearchInventoryResponseBody {
	s.RequestId = &v
	return s
}

type SearchInventoryResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SearchInventoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SearchInventoryResponse) String() string {
	return tea.Prettify(s)
}

func (s SearchInventoryResponse) GoString() string {
	return s.String()
}

func (s *SearchInventoryResponse) SetHeaders(v map[string]*string) *SearchInventoryResponse {
	s.Headers = v
	return s
}

func (s *SearchInventoryResponse) SetStatusCode(v int32) *SearchInventoryResponse {
	s.StatusCode = &v
	return s
}

func (s *SearchInventoryResponse) SetBody(v *SearchInventoryResponseBody) *SearchInventoryResponse {
	s.Body = v
	return s
}

type SetServiceSettingsRequest struct {
	// The name of OSS bucket to deliver.
	//
	// example:
	//
	// OssBucketName
	DeliveryOssBucketName *string `json:"DeliveryOssBucketName,omitempty" xml:"DeliveryOssBucketName,omitempty"`
	// Whether to enable OSS delivery.
	//
	// example:
	//
	// false
	DeliveryOssEnabled *bool `json:"DeliveryOssEnabled,omitempty" xml:"DeliveryOssEnabled,omitempty"`
	// The key prefix of OSS to deliver.
	//
	// example:
	//
	// oos/execution
	DeliveryOssKeyPrefix *string `json:"DeliveryOssKeyPrefix,omitempty" xml:"DeliveryOssKeyPrefix,omitempty"`
	// Whether to enable SLS delivery.
	//
	// example:
	//
	// false
	DeliverySlsEnabled *bool `json:"DeliverySlsEnabled,omitempty" xml:"DeliverySlsEnabled,omitempty"`
	// The name of SLS project to deliver.
	//
	// example:
	//
	// SlsProjectName
	DeliverySlsProjectName *string `json:"DeliverySlsProjectName,omitempty" xml:"DeliverySlsProjectName,omitempty"`
	// The id of RDC Enterprise.
	//
	// example:
	//
	// RdcEnterpriseId
	RdcEnterpriseId *string `json:"RdcEnterpriseId,omitempty" xml:"RdcEnterpriseId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s SetServiceSettingsRequest) String() string {
	return tea.Prettify(s)
}

func (s SetServiceSettingsRequest) GoString() string {
	return s.String()
}

func (s *SetServiceSettingsRequest) SetDeliveryOssBucketName(v string) *SetServiceSettingsRequest {
	s.DeliveryOssBucketName = &v
	return s
}

func (s *SetServiceSettingsRequest) SetDeliveryOssEnabled(v bool) *SetServiceSettingsRequest {
	s.DeliveryOssEnabled = &v
	return s
}

func (s *SetServiceSettingsRequest) SetDeliveryOssKeyPrefix(v string) *SetServiceSettingsRequest {
	s.DeliveryOssKeyPrefix = &v
	return s
}

func (s *SetServiceSettingsRequest) SetDeliverySlsEnabled(v bool) *SetServiceSettingsRequest {
	s.DeliverySlsEnabled = &v
	return s
}

func (s *SetServiceSettingsRequest) SetDeliverySlsProjectName(v string) *SetServiceSettingsRequest {
	s.DeliverySlsProjectName = &v
	return s
}

func (s *SetServiceSettingsRequest) SetRdcEnterpriseId(v string) *SetServiceSettingsRequest {
	s.RdcEnterpriseId = &v
	return s
}

func (s *SetServiceSettingsRequest) SetRegionId(v string) *SetServiceSettingsRequest {
	s.RegionId = &v
	return s
}

type SetServiceSettingsResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// CBEC8072-BEC2-478E-8EAE-E723BA79CF19
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information of service settings.
	ServiceSettings []*SetServiceSettingsResponseBodyServiceSettings `json:"ServiceSettings,omitempty" xml:"ServiceSettings,omitempty" type:"Repeated"`
}

func (s SetServiceSettingsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetServiceSettingsResponseBody) GoString() string {
	return s.String()
}

func (s *SetServiceSettingsResponseBody) SetRequestId(v string) *SetServiceSettingsResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetServiceSettingsResponseBody) SetServiceSettings(v []*SetServiceSettingsResponseBodyServiceSettings) *SetServiceSettingsResponseBody {
	s.ServiceSettings = v
	return s
}

type SetServiceSettingsResponseBodyServiceSettings struct {
	// The name of OSS bucket to deliver.
	//
	// example:
	//
	// OssBucketName
	DeliveryOssBucketName *string `json:"DeliveryOssBucketName,omitempty" xml:"DeliveryOssBucketName,omitempty"`
	// Whether to enable OSS delivery.
	//
	// example:
	//
	// true
	DeliveryOssEnabled *bool `json:"DeliveryOssEnabled,omitempty" xml:"DeliveryOssEnabled,omitempty"`
	// The key prefix of OSS to deliver.
	//
	// example:
	//
	// oos/execution
	DeliveryOssKeyPrefix *string `json:"DeliveryOssKeyPrefix,omitempty" xml:"DeliveryOssKeyPrefix,omitempty"`
	// Whether to enable SLS delivery.
	//
	// example:
	//
	// false
	DeliverySlsEnabled *bool `json:"DeliverySlsEnabled,omitempty" xml:"DeliverySlsEnabled,omitempty"`
	// The name of SLS project to deliver.
	//
	// example:
	//
	// SlsProjectName
	DeliverySlsProjectName *string `json:"DeliverySlsProjectName,omitempty" xml:"DeliverySlsProjectName,omitempty"`
	// The id of RDC Enterprise.
	//
	// example:
	//
	// RdcEnterpriseId
	RdcEnterpriseId *string `json:"RdcEnterpriseId,omitempty" xml:"RdcEnterpriseId,omitempty"`
}

func (s SetServiceSettingsResponseBodyServiceSettings) String() string {
	return tea.Prettify(s)
}

func (s SetServiceSettingsResponseBodyServiceSettings) GoString() string {
	return s.String()
}

func (s *SetServiceSettingsResponseBodyServiceSettings) SetDeliveryOssBucketName(v string) *SetServiceSettingsResponseBodyServiceSettings {
	s.DeliveryOssBucketName = &v
	return s
}

func (s *SetServiceSettingsResponseBodyServiceSettings) SetDeliveryOssEnabled(v bool) *SetServiceSettingsResponseBodyServiceSettings {
	s.DeliveryOssEnabled = &v
	return s
}

func (s *SetServiceSettingsResponseBodyServiceSettings) SetDeliveryOssKeyPrefix(v string) *SetServiceSettingsResponseBodyServiceSettings {
	s.DeliveryOssKeyPrefix = &v
	return s
}

func (s *SetServiceSettingsResponseBodyServiceSettings) SetDeliverySlsEnabled(v bool) *SetServiceSettingsResponseBodyServiceSettings {
	s.DeliverySlsEnabled = &v
	return s
}

func (s *SetServiceSettingsResponseBodyServiceSettings) SetDeliverySlsProjectName(v string) *SetServiceSettingsResponseBodyServiceSettings {
	s.DeliverySlsProjectName = &v
	return s
}

func (s *SetServiceSettingsResponseBodyServiceSettings) SetRdcEnterpriseId(v string) *SetServiceSettingsResponseBodyServiceSettings {
	s.RdcEnterpriseId = &v
	return s
}

type SetServiceSettingsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetServiceSettingsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetServiceSettingsResponse) String() string {
	return tea.Prettify(s)
}

func (s SetServiceSettingsResponse) GoString() string {
	return s.String()
}

func (s *SetServiceSettingsResponse) SetHeaders(v map[string]*string) *SetServiceSettingsResponse {
	s.Headers = v
	return s
}

func (s *SetServiceSettingsResponse) SetStatusCode(v int32) *SetServiceSettingsResponse {
	s.StatusCode = &v
	return s
}

func (s *SetServiceSettingsResponse) SetBody(v *SetServiceSettingsResponseBody) *SetServiceSettingsResponse {
	s.Body = v
	return s
}

type StartExecutionRequest struct {
	// The access token.
	//
	// example:
	//
	// 123e56767-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The description of the execution.
	//
	// example:
	//
	// test execution.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The loop mode. Valid values:
	//
	// 	- Automatic: does not suspend the execution of the template. This is the default value.
	//
	// 	- FirstBatchPause: suspends the execution of the template after the first batch is complete.
	//
	// 	- EveryBatchPause: suspends the execution of the template after each batch is complete.
	//
	// example:
	//
	// Automatic
	LoopMode *string `json:"LoopMode,omitempty" xml:"LoopMode,omitempty"`
	// The execution mode. Valid values:
	//
	// 	- Automatic: automatically starts the execution of the template. This is the default value.
	//
	// 	- FailurePause: suspends the execution of the template upon a failure.
	//
	// 	- Debug: manually starts the execution of the template.
	//
	// example:
	//
	// Automatic
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// The JSON string that consists of a set of parameters. Default value: {}.
	//
	// example:
	//
	// {"Status":"Running"}
	Parameters *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// The ID of the parent execution.
	//
	// example:
	//
	// exec-xxx
	ParentExecutionId *string `json:"ParentExecutionId,omitempty" xml:"ParentExecutionId,omitempty"`
	// The ID of the region in which the execution resides.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfmxsn4m4******
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The security check mode. Valid values:
	//
	// 	- Skip: specifies that you are aware of the risks. The system performs all actions in the execution without manual confirmation, regardless of the risk level. This parameter is valid only if the `Mode` parameter is set to Automatic.
	//
	// 	- ConfirmEveryHighRiskAction: requires you to confirm each high-risk action. This is the default value. You can call the **NotifyExecution*	- operation to confirm or cancel an action.
	//
	// example:
	//
	// Skip
	SafetyCheck *string `json:"SafetyCheck,omitempty" xml:"SafetyCheck,omitempty"`
	// The tags for the execution.
	//
	// example:
	//
	// {"k1":"v2","k2":"v2"}
	Tags map[string]interface{} `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The content of the template in the JSON or YAML format. This parameter is the same as the Content parameter that you can specify when you call the CreateTemplate operation. You can use this parameter to specify the tasks that you want to run. This way, you do not need to create a template before you start an execution. If you select an existing template, you do not need to specify this parameter.
	//
	// example:
	//
	// {   "Description": "Example template, describe instances in some status",   "FormatVersion": "OOS-2019-06-01",   "Parameters": {},   "Tasks": [     {       "Name": "describeInstances",       "Action": "ACS::ExecuteAPI",       "Description": "desc-en",       "Properties": {         "Service": "ECS",         "API": "DescribeInstances",         "Parameters": {           "Status": "Running"         }       }     }   ] }
	TemplateContent *string `json:"TemplateContent,omitempty" xml:"TemplateContent,omitempty"`
	// The name of the template. The name must be 1 to 200 characters in length, and can contain letters, digits, hyphens (-), and underscores (_).
	//
	// example:
	//
	// vmeixme
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// The Object Storage Service (OSS) URL of the object that stores the content of the Operation Orchestration Service (OOS) template. The access control list (ACL) of the object must be public-read. You can use this parameter to specify the tasks that you want to run. This way, you do not need to create a template before you start an execution. If you select an existing template, you do not need to specify this parameter.
	//
	// example:
	//
	// http://oos-template.cn-hangzhou.oss.aliyun-inc.com/oos-test-template.json
	TemplateURL *string `json:"TemplateURL,omitempty" xml:"TemplateURL,omitempty"`
	// The version number of the template. If you do not specify this parameter, the system uses the latest version.
	//
	// example:
	//
	// v1
	TemplateVersion *string `json:"TemplateVersion,omitempty" xml:"TemplateVersion,omitempty"`
}

func (s StartExecutionRequest) String() string {
	return tea.Prettify(s)
}

func (s StartExecutionRequest) GoString() string {
	return s.String()
}

func (s *StartExecutionRequest) SetClientToken(v string) *StartExecutionRequest {
	s.ClientToken = &v
	return s
}

func (s *StartExecutionRequest) SetDescription(v string) *StartExecutionRequest {
	s.Description = &v
	return s
}

func (s *StartExecutionRequest) SetLoopMode(v string) *StartExecutionRequest {
	s.LoopMode = &v
	return s
}

func (s *StartExecutionRequest) SetMode(v string) *StartExecutionRequest {
	s.Mode = &v
	return s
}

func (s *StartExecutionRequest) SetParameters(v string) *StartExecutionRequest {
	s.Parameters = &v
	return s
}

func (s *StartExecutionRequest) SetParentExecutionId(v string) *StartExecutionRequest {
	s.ParentExecutionId = &v
	return s
}

func (s *StartExecutionRequest) SetRegionId(v string) *StartExecutionRequest {
	s.RegionId = &v
	return s
}

func (s *StartExecutionRequest) SetResourceGroupId(v string) *StartExecutionRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *StartExecutionRequest) SetSafetyCheck(v string) *StartExecutionRequest {
	s.SafetyCheck = &v
	return s
}

func (s *StartExecutionRequest) SetTags(v map[string]interface{}) *StartExecutionRequest {
	s.Tags = v
	return s
}

func (s *StartExecutionRequest) SetTemplateContent(v string) *StartExecutionRequest {
	s.TemplateContent = &v
	return s
}

func (s *StartExecutionRequest) SetTemplateName(v string) *StartExecutionRequest {
	s.TemplateName = &v
	return s
}

func (s *StartExecutionRequest) SetTemplateURL(v string) *StartExecutionRequest {
	s.TemplateURL = &v
	return s
}

func (s *StartExecutionRequest) SetTemplateVersion(v string) *StartExecutionRequest {
	s.TemplateVersion = &v
	return s
}

type StartExecutionShrinkRequest struct {
	// The access token.
	//
	// example:
	//
	// 123e56767-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The description of the execution.
	//
	// example:
	//
	// test execution.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The loop mode. Valid values:
	//
	// 	- Automatic: does not suspend the execution of the template. This is the default value.
	//
	// 	- FirstBatchPause: suspends the execution of the template after the first batch is complete.
	//
	// 	- EveryBatchPause: suspends the execution of the template after each batch is complete.
	//
	// example:
	//
	// Automatic
	LoopMode *string `json:"LoopMode,omitempty" xml:"LoopMode,omitempty"`
	// The execution mode. Valid values:
	//
	// 	- Automatic: automatically starts the execution of the template. This is the default value.
	//
	// 	- FailurePause: suspends the execution of the template upon a failure.
	//
	// 	- Debug: manually starts the execution of the template.
	//
	// example:
	//
	// Automatic
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// The JSON string that consists of a set of parameters. Default value: {}.
	//
	// example:
	//
	// {"Status":"Running"}
	Parameters *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// The ID of the parent execution.
	//
	// example:
	//
	// exec-xxx
	ParentExecutionId *string `json:"ParentExecutionId,omitempty" xml:"ParentExecutionId,omitempty"`
	// The ID of the region in which the execution resides.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfmxsn4m4******
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The security check mode. Valid values:
	//
	// 	- Skip: specifies that you are aware of the risks. The system performs all actions in the execution without manual confirmation, regardless of the risk level. This parameter is valid only if the `Mode` parameter is set to Automatic.
	//
	// 	- ConfirmEveryHighRiskAction: requires you to confirm each high-risk action. This is the default value. You can call the **NotifyExecution*	- operation to confirm or cancel an action.
	//
	// example:
	//
	// Skip
	SafetyCheck *string `json:"SafetyCheck,omitempty" xml:"SafetyCheck,omitempty"`
	// The tags for the execution.
	//
	// example:
	//
	// {"k1":"v2","k2":"v2"}
	TagsShrink *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The content of the template in the JSON or YAML format. This parameter is the same as the Content parameter that you can specify when you call the CreateTemplate operation. You can use this parameter to specify the tasks that you want to run. This way, you do not need to create a template before you start an execution. If you select an existing template, you do not need to specify this parameter.
	//
	// example:
	//
	// {   "Description": "Example template, describe instances in some status",   "FormatVersion": "OOS-2019-06-01",   "Parameters": {},   "Tasks": [     {       "Name": "describeInstances",       "Action": "ACS::ExecuteAPI",       "Description": "desc-en",       "Properties": {         "Service": "ECS",         "API": "DescribeInstances",         "Parameters": {           "Status": "Running"         }       }     }   ] }
	TemplateContent *string `json:"TemplateContent,omitempty" xml:"TemplateContent,omitempty"`
	// The name of the template. The name must be 1 to 200 characters in length, and can contain letters, digits, hyphens (-), and underscores (_).
	//
	// example:
	//
	// vmeixme
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// The Object Storage Service (OSS) URL of the object that stores the content of the Operation Orchestration Service (OOS) template. The access control list (ACL) of the object must be public-read. You can use this parameter to specify the tasks that you want to run. This way, you do not need to create a template before you start an execution. If you select an existing template, you do not need to specify this parameter.
	//
	// example:
	//
	// http://oos-template.cn-hangzhou.oss.aliyun-inc.com/oos-test-template.json
	TemplateURL *string `json:"TemplateURL,omitempty" xml:"TemplateURL,omitempty"`
	// The version number of the template. If you do not specify this parameter, the system uses the latest version.
	//
	// example:
	//
	// v1
	TemplateVersion *string `json:"TemplateVersion,omitempty" xml:"TemplateVersion,omitempty"`
}

func (s StartExecutionShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s StartExecutionShrinkRequest) GoString() string {
	return s.String()
}

func (s *StartExecutionShrinkRequest) SetClientToken(v string) *StartExecutionShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *StartExecutionShrinkRequest) SetDescription(v string) *StartExecutionShrinkRequest {
	s.Description = &v
	return s
}

func (s *StartExecutionShrinkRequest) SetLoopMode(v string) *StartExecutionShrinkRequest {
	s.LoopMode = &v
	return s
}

func (s *StartExecutionShrinkRequest) SetMode(v string) *StartExecutionShrinkRequest {
	s.Mode = &v
	return s
}

func (s *StartExecutionShrinkRequest) SetParameters(v string) *StartExecutionShrinkRequest {
	s.Parameters = &v
	return s
}

func (s *StartExecutionShrinkRequest) SetParentExecutionId(v string) *StartExecutionShrinkRequest {
	s.ParentExecutionId = &v
	return s
}

func (s *StartExecutionShrinkRequest) SetRegionId(v string) *StartExecutionShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *StartExecutionShrinkRequest) SetResourceGroupId(v string) *StartExecutionShrinkRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *StartExecutionShrinkRequest) SetSafetyCheck(v string) *StartExecutionShrinkRequest {
	s.SafetyCheck = &v
	return s
}

func (s *StartExecutionShrinkRequest) SetTagsShrink(v string) *StartExecutionShrinkRequest {
	s.TagsShrink = &v
	return s
}

func (s *StartExecutionShrinkRequest) SetTemplateContent(v string) *StartExecutionShrinkRequest {
	s.TemplateContent = &v
	return s
}

func (s *StartExecutionShrinkRequest) SetTemplateName(v string) *StartExecutionShrinkRequest {
	s.TemplateName = &v
	return s
}

func (s *StartExecutionShrinkRequest) SetTemplateURL(v string) *StartExecutionShrinkRequest {
	s.TemplateURL = &v
	return s
}

func (s *StartExecutionShrinkRequest) SetTemplateVersion(v string) *StartExecutionShrinkRequest {
	s.TemplateVersion = &v
	return s
}

type StartExecutionResponseBody struct {
	// The details of the execution.
	Execution *StartExecutionResponseBodyExecution `json:"Execution,omitempty" xml:"Execution,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 14A07460-EBE7-47CA-9757-12CC4761D47A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StartExecutionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StartExecutionResponseBody) GoString() string {
	return s.String()
}

func (s *StartExecutionResponseBody) SetExecution(v *StartExecutionResponseBodyExecution) *StartExecutionResponseBody {
	s.Execution = v
	return s
}

func (s *StartExecutionResponseBody) SetRequestId(v string) *StartExecutionResponseBody {
	s.RequestId = &v
	return s
}

type StartExecutionResponseBodyExecution struct {
	// The number of executions.
	//
	// example:
	//
	// 1
	Counters map[string]interface{} `json:"Counters,omitempty" xml:"Counters,omitempty"`
	// The time when the execution was created.
	//
	// example:
	//
	// 2019-05-16T10:26:14Z
	CreateDate *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	// The information about in-progress tasks.
	CurrentTasks []*StartExecutionResponseBodyExecutionCurrentTasks `json:"CurrentTasks,omitempty" xml:"CurrentTasks,omitempty" type:"Repeated"`
	// The description of the execution.
	//
	// example:
	//
	// test execution.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The time when the execution stopped.
	//
	// example:
	//
	// 2019-05-16T10:26:14Z
	EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// The account ID of the user who started the execution of the template.
	//
	// example:
	//
	// root(13092080xx12344)
	ExecutedBy *string `json:"ExecutedBy,omitempty" xml:"ExecutedBy,omitempty"`
	// The GUID of the execution.
	//
	// example:
	//
	// exec-xxxyyy
	ExecutionId *string `json:"ExecutionId,omitempty" xml:"ExecutionId,omitempty"`
	// Indicates whether the execution is a parent execution.
	//
	// example:
	//
	// false
	IsParent *bool `json:"IsParent,omitempty" xml:"IsParent,omitempty"`
	// The loop mode.
	//
	// example:
	//
	// Automatic
	LoopMode *string `json:"LoopMode,omitempty" xml:"LoopMode,omitempty"`
	// The execution mode.
	//
	// example:
	//
	// Automatic
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// The output of the execution.
	//
	// example:
	//
	// { "InstanceId":"i-xxx" }
	Outputs *string `json:"Outputs,omitempty" xml:"Outputs,omitempty"`
	// The input parameters of the execution.
	//
	// example:
	//
	// { "Status":"Running" }
	Parameters *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// The ID of the parent execution.
	//
	// example:
	//
	// exec-xxxx
	ParentExecutionId *string `json:"ParentExecutionId,omitempty" xml:"ParentExecutionId,omitempty"`
	// The role that started the execution of the template.
	//
	// example:
	//
	// OOSServiceRole
	RamRole *string `json:"RamRole,omitempty" xml:"RamRole,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfmxsn4m4******
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The security check mode.
	//
	// example:
	//
	// Skip
	SafetyCheck *string `json:"SafetyCheck,omitempty" xml:"SafetyCheck,omitempty"`
	// The time when the execution was started.
	//
	// example:
	//
	// 2019-05-16T10:26:14Z
	StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	// The status of the execution.
	//
	// example:
	//
	// Success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The status information of the execution.
	//
	// example:
	//
	// ""
	StatusMessage *string `json:"StatusMessage,omitempty" xml:"StatusMessage,omitempty"`
	// The tags of the execution.
	//
	// example:
	//
	// {"k1":"v2","k2":"v2"}
	Tags map[string]interface{} `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The ID of the template.
	//
	// example:
	//
	// t-1bd341007f
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The name of the template.
	//
	// example:
	//
	// MyTemplate
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// The version number of the template.
	//
	// example:
	//
	// v1
	TemplateVersion *string `json:"TemplateVersion,omitempty" xml:"TemplateVersion,omitempty"`
	// The time when the execution was last updated.
	//
	// example:
	//
	// 2019-05-16T10:26:14Z
	UpdateDate *string `json:"UpdateDate,omitempty" xml:"UpdateDate,omitempty"`
}

func (s StartExecutionResponseBodyExecution) String() string {
	return tea.Prettify(s)
}

func (s StartExecutionResponseBodyExecution) GoString() string {
	return s.String()
}

func (s *StartExecutionResponseBodyExecution) SetCounters(v map[string]interface{}) *StartExecutionResponseBodyExecution {
	s.Counters = v
	return s
}

func (s *StartExecutionResponseBodyExecution) SetCreateDate(v string) *StartExecutionResponseBodyExecution {
	s.CreateDate = &v
	return s
}

func (s *StartExecutionResponseBodyExecution) SetCurrentTasks(v []*StartExecutionResponseBodyExecutionCurrentTasks) *StartExecutionResponseBodyExecution {
	s.CurrentTasks = v
	return s
}

func (s *StartExecutionResponseBodyExecution) SetDescription(v string) *StartExecutionResponseBodyExecution {
	s.Description = &v
	return s
}

func (s *StartExecutionResponseBodyExecution) SetEndDate(v string) *StartExecutionResponseBodyExecution {
	s.EndDate = &v
	return s
}

func (s *StartExecutionResponseBodyExecution) SetExecutedBy(v string) *StartExecutionResponseBodyExecution {
	s.ExecutedBy = &v
	return s
}

func (s *StartExecutionResponseBodyExecution) SetExecutionId(v string) *StartExecutionResponseBodyExecution {
	s.ExecutionId = &v
	return s
}

func (s *StartExecutionResponseBodyExecution) SetIsParent(v bool) *StartExecutionResponseBodyExecution {
	s.IsParent = &v
	return s
}

func (s *StartExecutionResponseBodyExecution) SetLoopMode(v string) *StartExecutionResponseBodyExecution {
	s.LoopMode = &v
	return s
}

func (s *StartExecutionResponseBodyExecution) SetMode(v string) *StartExecutionResponseBodyExecution {
	s.Mode = &v
	return s
}

func (s *StartExecutionResponseBodyExecution) SetOutputs(v string) *StartExecutionResponseBodyExecution {
	s.Outputs = &v
	return s
}

func (s *StartExecutionResponseBodyExecution) SetParameters(v string) *StartExecutionResponseBodyExecution {
	s.Parameters = &v
	return s
}

func (s *StartExecutionResponseBodyExecution) SetParentExecutionId(v string) *StartExecutionResponseBodyExecution {
	s.ParentExecutionId = &v
	return s
}

func (s *StartExecutionResponseBodyExecution) SetRamRole(v string) *StartExecutionResponseBodyExecution {
	s.RamRole = &v
	return s
}

func (s *StartExecutionResponseBodyExecution) SetResourceGroupId(v string) *StartExecutionResponseBodyExecution {
	s.ResourceGroupId = &v
	return s
}

func (s *StartExecutionResponseBodyExecution) SetSafetyCheck(v string) *StartExecutionResponseBodyExecution {
	s.SafetyCheck = &v
	return s
}

func (s *StartExecutionResponseBodyExecution) SetStartDate(v string) *StartExecutionResponseBodyExecution {
	s.StartDate = &v
	return s
}

func (s *StartExecutionResponseBodyExecution) SetStatus(v string) *StartExecutionResponseBodyExecution {
	s.Status = &v
	return s
}

func (s *StartExecutionResponseBodyExecution) SetStatusMessage(v string) *StartExecutionResponseBodyExecution {
	s.StatusMessage = &v
	return s
}

func (s *StartExecutionResponseBodyExecution) SetTags(v map[string]interface{}) *StartExecutionResponseBodyExecution {
	s.Tags = v
	return s
}

func (s *StartExecutionResponseBodyExecution) SetTemplateId(v string) *StartExecutionResponseBodyExecution {
	s.TemplateId = &v
	return s
}

func (s *StartExecutionResponseBodyExecution) SetTemplateName(v string) *StartExecutionResponseBodyExecution {
	s.TemplateName = &v
	return s
}

func (s *StartExecutionResponseBodyExecution) SetTemplateVersion(v string) *StartExecutionResponseBodyExecution {
	s.TemplateVersion = &v
	return s
}

func (s *StartExecutionResponseBodyExecution) SetUpdateDate(v string) *StartExecutionResponseBodyExecution {
	s.UpdateDate = &v
	return s
}

type StartExecutionResponseBodyExecutionCurrentTasks struct {
	// The action of the task.
	//
	// example:
	//
	// ACS::TimerTrigger
	TaskAction *string `json:"TaskAction,omitempty" xml:"TaskAction,omitempty"`
	// The execution ID of the task.
	//
	// example:
	//
	// exec-dsadasdawq
	TaskExecutionId *string `json:"TaskExecutionId,omitempty" xml:"TaskExecutionId,omitempty"`
	// The name of the task.
	//
	// example:
	//
	// testTask
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
}

func (s StartExecutionResponseBodyExecutionCurrentTasks) String() string {
	return tea.Prettify(s)
}

func (s StartExecutionResponseBodyExecutionCurrentTasks) GoString() string {
	return s.String()
}

func (s *StartExecutionResponseBodyExecutionCurrentTasks) SetTaskAction(v string) *StartExecutionResponseBodyExecutionCurrentTasks {
	s.TaskAction = &v
	return s
}

func (s *StartExecutionResponseBodyExecutionCurrentTasks) SetTaskExecutionId(v string) *StartExecutionResponseBodyExecutionCurrentTasks {
	s.TaskExecutionId = &v
	return s
}

func (s *StartExecutionResponseBodyExecutionCurrentTasks) SetTaskName(v string) *StartExecutionResponseBodyExecutionCurrentTasks {
	s.TaskName = &v
	return s
}

type StartExecutionResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartExecutionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartExecutionResponse) String() string {
	return tea.Prettify(s)
}

func (s StartExecutionResponse) GoString() string {
	return s.String()
}

func (s *StartExecutionResponse) SetHeaders(v map[string]*string) *StartExecutionResponse {
	s.Headers = v
	return s
}

func (s *StartExecutionResponse) SetStatusCode(v int32) *StartExecutionResponse {
	s.StatusCode = &v
	return s
}

func (s *StartExecutionResponse) SetBody(v *StartExecutionResponseBody) *StartExecutionResponse {
	s.Body = v
	return s
}

type TagResourcesRequest struct {
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The IDs of the resources for which you want to modify the resource group. The number of resource IDs is 1 to 50.
	//
	// 	- If you set ResourceType to template, specify ResourceIds in the ["TemplateName1","TemplateName2"] format.
	//
	// 	- If you set ResourceType to parameter, specify ResourceIds in the ["Name1","Name2"] format.
	//
	// 	- If you set ResourceType to secretparameter, specify ResourceIds in the ["Name1","Name2"] format.
	//
	// 	- If you set ResourceType to stateconfiguration, specify ResourceIds in the ["StateConfigurationId 1","StateConfigurationId 2"] format.
	//
	// 	- If you set ResourceType to application, specify ResourceIds in the ["Name1","Name2"] format.
	//
	// This parameter is required.
	//
	// example:
	//
	// ["templateName1","templateName2"]
	ResourceIds map[string]interface{} `json:"ResourceIds,omitempty" xml:"ResourceIds,omitempty"`
	// The type of the resource for which you want to modify the resource group. Valid values:
	//
	// 	- template: template.
	//
	// 	- parameter: parameter.
	//
	// 	- secretparameter: encryption parameter.
	//
	// 	- stateconfiguration: desired-state configuration.
	//
	// 	- application: application.
	//
	// This parameter is required.
	//
	// example:
	//
	// template
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The tag keys and values. The number of key-value pairs ranges from 1 to 20.
	//
	// This parameter is required.
	//
	// example:
	//
	// {"k1":"v1","k2":"v2"}
	Tags map[string]interface{} `json:"Tags,omitempty" xml:"Tags,omitempty"`
}

func (s TagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesRequest) GoString() string {
	return s.String()
}

func (s *TagResourcesRequest) SetRegionId(v string) *TagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *TagResourcesRequest) SetResourceIds(v map[string]interface{}) *TagResourcesRequest {
	s.ResourceIds = v
	return s
}

func (s *TagResourcesRequest) SetResourceType(v string) *TagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *TagResourcesRequest) SetTags(v map[string]interface{}) *TagResourcesRequest {
	s.Tags = v
	return s
}

type TagResourcesShrinkRequest struct {
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The IDs of the resources for which you want to modify the resource group. The number of resource IDs is 1 to 50.
	//
	// 	- If you set ResourceType to template, specify ResourceIds in the ["TemplateName1","TemplateName2"] format.
	//
	// 	- If you set ResourceType to parameter, specify ResourceIds in the ["Name1","Name2"] format.
	//
	// 	- If you set ResourceType to secretparameter, specify ResourceIds in the ["Name1","Name2"] format.
	//
	// 	- If you set ResourceType to stateconfiguration, specify ResourceIds in the ["StateConfigurationId 1","StateConfigurationId 2"] format.
	//
	// 	- If you set ResourceType to application, specify ResourceIds in the ["Name1","Name2"] format.
	//
	// This parameter is required.
	//
	// example:
	//
	// ["templateName1","templateName2"]
	ResourceIdsShrink *string `json:"ResourceIds,omitempty" xml:"ResourceIds,omitempty"`
	// The type of the resource for which you want to modify the resource group. Valid values:
	//
	// 	- template: template.
	//
	// 	- parameter: parameter.
	//
	// 	- secretparameter: encryption parameter.
	//
	// 	- stateconfiguration: desired-state configuration.
	//
	// 	- application: application.
	//
	// This parameter is required.
	//
	// example:
	//
	// template
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The tag keys and values. The number of key-value pairs ranges from 1 to 20.
	//
	// This parameter is required.
	//
	// example:
	//
	// {"k1":"v1","k2":"v2"}
	TagsShrink *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
}

func (s TagResourcesShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesShrinkRequest) GoString() string {
	return s.String()
}

func (s *TagResourcesShrinkRequest) SetRegionId(v string) *TagResourcesShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *TagResourcesShrinkRequest) SetResourceIdsShrink(v string) *TagResourcesShrinkRequest {
	s.ResourceIdsShrink = &v
	return s
}

func (s *TagResourcesShrinkRequest) SetResourceType(v string) *TagResourcesShrinkRequest {
	s.ResourceType = &v
	return s
}

func (s *TagResourcesShrinkRequest) SetTagsShrink(v string) *TagResourcesShrinkRequest {
	s.TagsShrink = &v
	return s
}

type TagResourcesResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// B19AE203-FD99-49C7-9253-FAAACAD46F4A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s TagResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *TagResourcesResponseBody) SetRequestId(v string) *TagResourcesResponseBody {
	s.RequestId = &v
	return s
}

type TagResourcesResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TagResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesResponse) GoString() string {
	return s.String()
}

func (s *TagResourcesResponse) SetHeaders(v map[string]*string) *TagResourcesResponse {
	s.Headers = v
	return s
}

func (s *TagResourcesResponse) SetStatusCode(v int32) *TagResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *TagResourcesResponse) SetBody(v *TagResourcesResponseBody) *TagResourcesResponse {
	s.Body = v
	return s
}

type TriggerExecutionRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// example:
	//
	// dswe2-3i0-029
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The message body to be sent to the trigger task.
	//
	// example:
	//
	// {"eventTime": "20181226T220114.058+0800", "id": "9435EAD6-3CF6-4494-8F7A-3A********77","level": "INFO","name": "Instance:StateChange","product": "ECS","regionId":"cn-hangzhou","resourceId": "acs:ecs:cn-hangzhou:169070********30:instance/i-bp1ecr********5go2go","userId": "169070********30","ver": "1.0","content": {"resourceId": "i-bp1ecr********5go2go", "resourceType": "ALIYUN::ECS::Instance","state": "Stopping"} }
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The ID of the event-, alert-, or timer-triggered execution.
	//
	// This parameter is required.
	//
	// example:
	//
	// exec-sadw3f23rsad
	ExecutionId *string `json:"ExecutionId,omitempty" xml:"ExecutionId,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The type of the trigger. Valid values:
	//
	// 	- Event
	//
	// 	- Alarm
	//
	// 	- Timer
	//
	// This parameter is required.
	//
	// example:
	//
	// Event
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s TriggerExecutionRequest) String() string {
	return tea.Prettify(s)
}

func (s TriggerExecutionRequest) GoString() string {
	return s.String()
}

func (s *TriggerExecutionRequest) SetClientToken(v string) *TriggerExecutionRequest {
	s.ClientToken = &v
	return s
}

func (s *TriggerExecutionRequest) SetContent(v string) *TriggerExecutionRequest {
	s.Content = &v
	return s
}

func (s *TriggerExecutionRequest) SetExecutionId(v string) *TriggerExecutionRequest {
	s.ExecutionId = &v
	return s
}

func (s *TriggerExecutionRequest) SetRegionId(v string) *TriggerExecutionRequest {
	s.RegionId = &v
	return s
}

func (s *TriggerExecutionRequest) SetType(v string) *TriggerExecutionRequest {
	s.Type = &v
	return s
}

type TriggerExecutionResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 14A07460-EBE7-47CA-9757-12CC4761D47A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s TriggerExecutionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TriggerExecutionResponseBody) GoString() string {
	return s.String()
}

func (s *TriggerExecutionResponseBody) SetRequestId(v string) *TriggerExecutionResponseBody {
	s.RequestId = &v
	return s
}

type TriggerExecutionResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TriggerExecutionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TriggerExecutionResponse) String() string {
	return tea.Prettify(s)
}

func (s TriggerExecutionResponse) GoString() string {
	return s.String()
}

func (s *TriggerExecutionResponse) SetHeaders(v map[string]*string) *TriggerExecutionResponse {
	s.Headers = v
	return s
}

func (s *TriggerExecutionResponse) SetStatusCode(v int32) *TriggerExecutionResponse {
	s.StatusCode = &v
	return s
}

func (s *TriggerExecutionResponse) SetBody(v *TriggerExecutionResponseBody) *TriggerExecutionResponse {
	s.Body = v
	return s
}

type UntagResourcesRequest struct {
	// Specifies whether to remove all tags. This parameter takes effect only if TagKeys is left empty. Valid values: true and false. Default value: false. TagKeys is required if this parameter is set to false.
	//
	// example:
	//
	// false
	All *bool `json:"All,omitempty" xml:"All,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The IDs of the resources for which you want to modify the resource group. The number of resource IDs is 1 to 50.
	//
	// 	- If you set ResourceType to template, specify ResourceIds in the ["TemplateName1","TemplateName2"] format.
	//
	// 	- If you set ResourceType to parameter, specify ResourceIds in the ["Name1","Name2"] format.
	//
	// 	- If you set ResourceType to secretparameter, specify ResourceIds in the ["Name1","Name2"] format.
	//
	// 	- If you set ResourceType to stateconfiguration, specify ResourceIds in the ["StateConfigurationId 1","StateConfigurationId 2"] format.
	//
	// 	- If you set ResourceType to application, specify ResourceIds in the ["Name1","Name2"] format.
	//
	// This parameter is required.
	//
	// example:
	//
	// ["templateName1","templateName2"]
	ResourceIds map[string]interface{} `json:"ResourceIds,omitempty" xml:"ResourceIds,omitempty"`
	// The type of the resource for which you want to modify the resource group. Valid values:
	//
	// 	- template: template.
	//
	// 	- parameter: parameter.
	//
	// 	- secretparameter: encryption parameter.
	//
	// 	- stateconfiguration: desired-state configuration.
	//
	// 	- application: application.
	//
	// This parameter is required.
	//
	// example:
	//
	// template
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The tag keys. The number of keys ranges from 1 to 20.
	//
	// example:
	//
	// ["k1","k2"]
	TagKeys map[string]interface{} `json:"TagKeys,omitempty" xml:"TagKeys,omitempty"`
}

func (s UntagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesRequest) GoString() string {
	return s.String()
}

func (s *UntagResourcesRequest) SetAll(v bool) *UntagResourcesRequest {
	s.All = &v
	return s
}

func (s *UntagResourcesRequest) SetRegionId(v string) *UntagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *UntagResourcesRequest) SetResourceIds(v map[string]interface{}) *UntagResourcesRequest {
	s.ResourceIds = v
	return s
}

func (s *UntagResourcesRequest) SetResourceType(v string) *UntagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *UntagResourcesRequest) SetTagKeys(v map[string]interface{}) *UntagResourcesRequest {
	s.TagKeys = v
	return s
}

type UntagResourcesShrinkRequest struct {
	// Specifies whether to remove all tags. This parameter takes effect only if TagKeys is left empty. Valid values: true and false. Default value: false. TagKeys is required if this parameter is set to false.
	//
	// example:
	//
	// false
	All *bool `json:"All,omitempty" xml:"All,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The IDs of the resources for which you want to modify the resource group. The number of resource IDs is 1 to 50.
	//
	// 	- If you set ResourceType to template, specify ResourceIds in the ["TemplateName1","TemplateName2"] format.
	//
	// 	- If you set ResourceType to parameter, specify ResourceIds in the ["Name1","Name2"] format.
	//
	// 	- If you set ResourceType to secretparameter, specify ResourceIds in the ["Name1","Name2"] format.
	//
	// 	- If you set ResourceType to stateconfiguration, specify ResourceIds in the ["StateConfigurationId 1","StateConfigurationId 2"] format.
	//
	// 	- If you set ResourceType to application, specify ResourceIds in the ["Name1","Name2"] format.
	//
	// This parameter is required.
	//
	// example:
	//
	// ["templateName1","templateName2"]
	ResourceIdsShrink *string `json:"ResourceIds,omitempty" xml:"ResourceIds,omitempty"`
	// The type of the resource for which you want to modify the resource group. Valid values:
	//
	// 	- template: template.
	//
	// 	- parameter: parameter.
	//
	// 	- secretparameter: encryption parameter.
	//
	// 	- stateconfiguration: desired-state configuration.
	//
	// 	- application: application.
	//
	// This parameter is required.
	//
	// example:
	//
	// template
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The tag keys. The number of keys ranges from 1 to 20.
	//
	// example:
	//
	// ["k1","k2"]
	TagKeysShrink *string `json:"TagKeys,omitempty" xml:"TagKeys,omitempty"`
}

func (s UntagResourcesShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesShrinkRequest) GoString() string {
	return s.String()
}

func (s *UntagResourcesShrinkRequest) SetAll(v bool) *UntagResourcesShrinkRequest {
	s.All = &v
	return s
}

func (s *UntagResourcesShrinkRequest) SetRegionId(v string) *UntagResourcesShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *UntagResourcesShrinkRequest) SetResourceIdsShrink(v string) *UntagResourcesShrinkRequest {
	s.ResourceIdsShrink = &v
	return s
}

func (s *UntagResourcesShrinkRequest) SetResourceType(v string) *UntagResourcesShrinkRequest {
	s.ResourceType = &v
	return s
}

func (s *UntagResourcesShrinkRequest) SetTagKeysShrink(v string) *UntagResourcesShrinkRequest {
	s.TagKeysShrink = &v
	return s
}

type UntagResourcesResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 491DF8C2-34C9-4679-9DB3-4C0F49B129AC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UntagResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *UntagResourcesResponseBody) SetRequestId(v string) *UntagResourcesResponseBody {
	s.RequestId = &v
	return s
}

type UntagResourcesResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UntagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UntagResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesResponse) GoString() string {
	return s.String()
}

func (s *UntagResourcesResponse) SetHeaders(v map[string]*string) *UntagResourcesResponse {
	s.Headers = v
	return s
}

func (s *UntagResourcesResponse) SetStatusCode(v int32) *UntagResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *UntagResourcesResponse) SetBody(v *UntagResourcesResponseBody) *UntagResourcesResponse {
	s.Body = v
	return s
}

type UpdateApplicationRequest struct {
	// The configurations of application alerts.
	AlarmConfig *UpdateApplicationRequestAlarmConfig `json:"AlarmConfig,omitempty" xml:"AlarmConfig,omitempty" type:"Struct"`
	// Specifies whether to delete existing alert rules before applying the alert template. Default value: false.
	//
	// example:
	//
	// false
	DeleteAlarmRulesBeforeUpdate *bool `json:"DeleteAlarmRulesBeforeUpdate,omitempty" xml:"DeleteAlarmRulesBeforeUpdate,omitempty"`
	// The description to be updated for the application.
	//
	// example:
	//
	// test application
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The application name.
	//
	// This parameter is required.
	//
	// example:
	//
	// My-Application
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The region ID. Set the value to cn-hangzhou.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The tags.
	//
	// example:
	//
	// {"k1":"v1","k2":"v2"}
	Tags map[string]interface{} `json:"Tags,omitempty" xml:"Tags,omitempty"`
}

func (s UpdateApplicationRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateApplicationRequest) GoString() string {
	return s.String()
}

func (s *UpdateApplicationRequest) SetAlarmConfig(v *UpdateApplicationRequestAlarmConfig) *UpdateApplicationRequest {
	s.AlarmConfig = v
	return s
}

func (s *UpdateApplicationRequest) SetDeleteAlarmRulesBeforeUpdate(v bool) *UpdateApplicationRequest {
	s.DeleteAlarmRulesBeforeUpdate = &v
	return s
}

func (s *UpdateApplicationRequest) SetDescription(v string) *UpdateApplicationRequest {
	s.Description = &v
	return s
}

func (s *UpdateApplicationRequest) SetName(v string) *UpdateApplicationRequest {
	s.Name = &v
	return s
}

func (s *UpdateApplicationRequest) SetRegionId(v string) *UpdateApplicationRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateApplicationRequest) SetTags(v map[string]interface{}) *UpdateApplicationRequest {
	s.Tags = v
	return s
}

type UpdateApplicationRequestAlarmConfig struct {
	// The alert contact groups.
	ContactGroups []*string `json:"ContactGroups,omitempty" xml:"ContactGroups,omitempty" type:"Repeated"`
	// The health check URL of the application.
	//
	// example:
	//
	// /healthcheck/tcp50122
	HealthCheckUrl *string `json:"HealthCheckUrl,omitempty" xml:"HealthCheckUrl,omitempty"`
	// The alert templates.
	TemplateIds []*string `json:"TemplateIds,omitempty" xml:"TemplateIds,omitempty" type:"Repeated"`
}

func (s UpdateApplicationRequestAlarmConfig) String() string {
	return tea.Prettify(s)
}

func (s UpdateApplicationRequestAlarmConfig) GoString() string {
	return s.String()
}

func (s *UpdateApplicationRequestAlarmConfig) SetContactGroups(v []*string) *UpdateApplicationRequestAlarmConfig {
	s.ContactGroups = v
	return s
}

func (s *UpdateApplicationRequestAlarmConfig) SetHealthCheckUrl(v string) *UpdateApplicationRequestAlarmConfig {
	s.HealthCheckUrl = &v
	return s
}

func (s *UpdateApplicationRequestAlarmConfig) SetTemplateIds(v []*string) *UpdateApplicationRequestAlarmConfig {
	s.TemplateIds = v
	return s
}

type UpdateApplicationShrinkRequest struct {
	// The configurations of application alerts.
	AlarmConfigShrink *string `json:"AlarmConfig,omitempty" xml:"AlarmConfig,omitempty"`
	// Specifies whether to delete existing alert rules before applying the alert template. Default value: false.
	//
	// example:
	//
	// false
	DeleteAlarmRulesBeforeUpdate *bool `json:"DeleteAlarmRulesBeforeUpdate,omitempty" xml:"DeleteAlarmRulesBeforeUpdate,omitempty"`
	// The description to be updated for the application.
	//
	// example:
	//
	// test application
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The application name.
	//
	// This parameter is required.
	//
	// example:
	//
	// My-Application
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The region ID. Set the value to cn-hangzhou.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The tags.
	//
	// example:
	//
	// {"k1":"v1","k2":"v2"}
	TagsShrink *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
}

func (s UpdateApplicationShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateApplicationShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateApplicationShrinkRequest) SetAlarmConfigShrink(v string) *UpdateApplicationShrinkRequest {
	s.AlarmConfigShrink = &v
	return s
}

func (s *UpdateApplicationShrinkRequest) SetDeleteAlarmRulesBeforeUpdate(v bool) *UpdateApplicationShrinkRequest {
	s.DeleteAlarmRulesBeforeUpdate = &v
	return s
}

func (s *UpdateApplicationShrinkRequest) SetDescription(v string) *UpdateApplicationShrinkRequest {
	s.Description = &v
	return s
}

func (s *UpdateApplicationShrinkRequest) SetName(v string) *UpdateApplicationShrinkRequest {
	s.Name = &v
	return s
}

func (s *UpdateApplicationShrinkRequest) SetRegionId(v string) *UpdateApplicationShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateApplicationShrinkRequest) SetTagsShrink(v string) *UpdateApplicationShrinkRequest {
	s.TagsShrink = &v
	return s
}

type UpdateApplicationResponseBody struct {
	// The information about the application.
	Application *UpdateApplicationResponseBodyApplication `json:"Application,omitempty" xml:"Application,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// F1F00F41-D24C-5377-831B-C97F739CE1AB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateApplicationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateApplicationResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateApplicationResponseBody) SetApplication(v *UpdateApplicationResponseBodyApplication) *UpdateApplicationResponseBody {
	s.Application = v
	return s
}

func (s *UpdateApplicationResponseBody) SetRequestId(v string) *UpdateApplicationResponseBody {
	s.RequestId = &v
	return s
}

type UpdateApplicationResponseBodyApplication struct {
	// The time when the application was created.
	//
	// example:
	//
	// 2021-09-07T09:17:46Z
	CreatedDate *string `json:"CreatedDate,omitempty" xml:"CreatedDate,omitempty"`
	// The description of the application.
	//
	// example:
	//
	// test application
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The application name.
	//
	// example:
	//
	// My-Application
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfmxsn4m*****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The tags.
	//
	// example:
	//
	// {"k1":"v1","k2":"v2"}
	Tags map[string]interface{} `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The time when the application was updated.
	//
	// example:
	//
	// 2021-09-07T10:17:46Z
	UpdatedDate *string `json:"UpdatedDate,omitempty" xml:"UpdatedDate,omitempty"`
}

func (s UpdateApplicationResponseBodyApplication) String() string {
	return tea.Prettify(s)
}

func (s UpdateApplicationResponseBodyApplication) GoString() string {
	return s.String()
}

func (s *UpdateApplicationResponseBodyApplication) SetCreatedDate(v string) *UpdateApplicationResponseBodyApplication {
	s.CreatedDate = &v
	return s
}

func (s *UpdateApplicationResponseBodyApplication) SetDescription(v string) *UpdateApplicationResponseBodyApplication {
	s.Description = &v
	return s
}

func (s *UpdateApplicationResponseBodyApplication) SetName(v string) *UpdateApplicationResponseBodyApplication {
	s.Name = &v
	return s
}

func (s *UpdateApplicationResponseBodyApplication) SetResourceGroupId(v string) *UpdateApplicationResponseBodyApplication {
	s.ResourceGroupId = &v
	return s
}

func (s *UpdateApplicationResponseBodyApplication) SetTags(v map[string]interface{}) *UpdateApplicationResponseBodyApplication {
	s.Tags = v
	return s
}

func (s *UpdateApplicationResponseBodyApplication) SetUpdatedDate(v string) *UpdateApplicationResponseBodyApplication {
	s.UpdatedDate = &v
	return s
}

type UpdateApplicationResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateApplicationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateApplicationResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateApplicationResponse) GoString() string {
	return s.String()
}

func (s *UpdateApplicationResponse) SetHeaders(v map[string]*string) *UpdateApplicationResponse {
	s.Headers = v
	return s
}

func (s *UpdateApplicationResponse) SetStatusCode(v int32) *UpdateApplicationResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateApplicationResponse) SetBody(v *UpdateApplicationResponseBody) *UpdateApplicationResponse {
	s.Body = v
	return s
}

type UpdateApplicationGroupRequest struct {
	// The application name.
	//
	// This parameter is required.
	//
	// example:
	//
	// MyApplication
	ApplicationName *string `json:"ApplicationName,omitempty" xml:"ApplicationName,omitempty"`
	// The name of the application group.
	//
	// This parameter is required.
	//
	// example:
	//
	// MyApplicationGroup
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The new name of the application group.
	//
	// example:
	//
	// UpdateMyApplicationGroup
	NewName       *string                `json:"NewName,omitempty" xml:"NewName,omitempty"`
	OperationName *string                `json:"OperationName,omitempty" xml:"OperationName,omitempty"`
	Parameters    map[string]interface{} `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// The region ID. Set the value to cn-hangzhou.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s UpdateApplicationGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateApplicationGroupRequest) GoString() string {
	return s.String()
}

func (s *UpdateApplicationGroupRequest) SetApplicationName(v string) *UpdateApplicationGroupRequest {
	s.ApplicationName = &v
	return s
}

func (s *UpdateApplicationGroupRequest) SetName(v string) *UpdateApplicationGroupRequest {
	s.Name = &v
	return s
}

func (s *UpdateApplicationGroupRequest) SetNewName(v string) *UpdateApplicationGroupRequest {
	s.NewName = &v
	return s
}

func (s *UpdateApplicationGroupRequest) SetOperationName(v string) *UpdateApplicationGroupRequest {
	s.OperationName = &v
	return s
}

func (s *UpdateApplicationGroupRequest) SetParameters(v map[string]interface{}) *UpdateApplicationGroupRequest {
	s.Parameters = v
	return s
}

func (s *UpdateApplicationGroupRequest) SetRegionId(v string) *UpdateApplicationGroupRequest {
	s.RegionId = &v
	return s
}

type UpdateApplicationGroupShrinkRequest struct {
	// The application name.
	//
	// This parameter is required.
	//
	// example:
	//
	// MyApplication
	ApplicationName *string `json:"ApplicationName,omitempty" xml:"ApplicationName,omitempty"`
	// The name of the application group.
	//
	// This parameter is required.
	//
	// example:
	//
	// MyApplicationGroup
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The new name of the application group.
	//
	// example:
	//
	// UpdateMyApplicationGroup
	NewName          *string `json:"NewName,omitempty" xml:"NewName,omitempty"`
	OperationName    *string `json:"OperationName,omitempty" xml:"OperationName,omitempty"`
	ParametersShrink *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// The region ID. Set the value to cn-hangzhou.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s UpdateApplicationGroupShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateApplicationGroupShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateApplicationGroupShrinkRequest) SetApplicationName(v string) *UpdateApplicationGroupShrinkRequest {
	s.ApplicationName = &v
	return s
}

func (s *UpdateApplicationGroupShrinkRequest) SetName(v string) *UpdateApplicationGroupShrinkRequest {
	s.Name = &v
	return s
}

func (s *UpdateApplicationGroupShrinkRequest) SetNewName(v string) *UpdateApplicationGroupShrinkRequest {
	s.NewName = &v
	return s
}

func (s *UpdateApplicationGroupShrinkRequest) SetOperationName(v string) *UpdateApplicationGroupShrinkRequest {
	s.OperationName = &v
	return s
}

func (s *UpdateApplicationGroupShrinkRequest) SetParametersShrink(v string) *UpdateApplicationGroupShrinkRequest {
	s.ParametersShrink = &v
	return s
}

func (s *UpdateApplicationGroupShrinkRequest) SetRegionId(v string) *UpdateApplicationGroupShrinkRequest {
	s.RegionId = &v
	return s
}

type UpdateApplicationGroupResponseBody struct {
	// The information about the application group.
	ApplicationGroup *UpdateApplicationGroupResponseBodyApplicationGroup `json:"ApplicationGroup,omitempty" xml:"ApplicationGroup,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// AA9FA778-AE4B-55EC-81CC-C46BAF08A166
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateApplicationGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateApplicationGroupResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateApplicationGroupResponseBody) SetApplicationGroup(v *UpdateApplicationGroupResponseBodyApplicationGroup) *UpdateApplicationGroupResponseBody {
	s.ApplicationGroup = v
	return s
}

func (s *UpdateApplicationGroupResponseBody) SetRequestId(v string) *UpdateApplicationGroupResponseBody {
	s.RequestId = &v
	return s
}

type UpdateApplicationGroupResponseBodyApplicationGroup struct {
	// The application name.
	//
	// example:
	//
	// MyApplication
	ApplicationName *string `json:"ApplicationName,omitempty" xml:"ApplicationName,omitempty"`
	// The time when the application group was created.
	//
	// example:
	//
	// 2021-09-07T10:28:25Z
	CreatedDate *string `json:"CreatedDate,omitempty" xml:"CreatedDate,omitempty"`
	// The ID of the region in which the related resources reside.
	//
	// example:
	//
	// cn-hangzhou
	DeployRegionId *string `json:"DeployRegionId,omitempty" xml:"DeployRegionId,omitempty"`
	// The description of the application group.
	//
	// example:
	//
	// ApplicationGroup
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The key of the tag.
	//
	// example:
	//
	// k1
	ImportTagKey *string `json:"ImportTagKey,omitempty" xml:"ImportTagKey,omitempty"`
	// The value of the tag.
	//
	// example:
	//
	// v1
	ImportTagValue *string `json:"ImportTagValue,omitempty" xml:"ImportTagValue,omitempty"`
	// The name of the application group.
	//
	// example:
	//
	// UpdateMyApplicationGroup
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The time when the application group was updated.
	//
	// example:
	//
	// 2021-09-08T03:01:53Z
	UpdatedDate *string `json:"UpdatedDate,omitempty" xml:"UpdatedDate,omitempty"`
}

func (s UpdateApplicationGroupResponseBodyApplicationGroup) String() string {
	return tea.Prettify(s)
}

func (s UpdateApplicationGroupResponseBodyApplicationGroup) GoString() string {
	return s.String()
}

func (s *UpdateApplicationGroupResponseBodyApplicationGroup) SetApplicationName(v string) *UpdateApplicationGroupResponseBodyApplicationGroup {
	s.ApplicationName = &v
	return s
}

func (s *UpdateApplicationGroupResponseBodyApplicationGroup) SetCreatedDate(v string) *UpdateApplicationGroupResponseBodyApplicationGroup {
	s.CreatedDate = &v
	return s
}

func (s *UpdateApplicationGroupResponseBodyApplicationGroup) SetDeployRegionId(v string) *UpdateApplicationGroupResponseBodyApplicationGroup {
	s.DeployRegionId = &v
	return s
}

func (s *UpdateApplicationGroupResponseBodyApplicationGroup) SetDescription(v string) *UpdateApplicationGroupResponseBodyApplicationGroup {
	s.Description = &v
	return s
}

func (s *UpdateApplicationGroupResponseBodyApplicationGroup) SetImportTagKey(v string) *UpdateApplicationGroupResponseBodyApplicationGroup {
	s.ImportTagKey = &v
	return s
}

func (s *UpdateApplicationGroupResponseBodyApplicationGroup) SetImportTagValue(v string) *UpdateApplicationGroupResponseBodyApplicationGroup {
	s.ImportTagValue = &v
	return s
}

func (s *UpdateApplicationGroupResponseBodyApplicationGroup) SetName(v string) *UpdateApplicationGroupResponseBodyApplicationGroup {
	s.Name = &v
	return s
}

func (s *UpdateApplicationGroupResponseBodyApplicationGroup) SetUpdatedDate(v string) *UpdateApplicationGroupResponseBodyApplicationGroup {
	s.UpdatedDate = &v
	return s
}

type UpdateApplicationGroupResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateApplicationGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateApplicationGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateApplicationGroupResponse) GoString() string {
	return s.String()
}

func (s *UpdateApplicationGroupResponse) SetHeaders(v map[string]*string) *UpdateApplicationGroupResponse {
	s.Headers = v
	return s
}

func (s *UpdateApplicationGroupResponse) SetStatusCode(v int32) *UpdateApplicationGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateApplicationGroupResponse) SetBody(v *UpdateApplicationGroupResponseBody) *UpdateApplicationGroupResponse {
	s.Body = v
	return s
}

type UpdateExecutionRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-42665544****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The description of the execution.
	//
	// example:
	//
	// Execution description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the execution.
	//
	// This parameter is required.
	//
	// example:
	//
	// exec-c223xxxxxxxxxxxxxxxx
	ExecutionId *string `json:"ExecutionId,omitempty" xml:"ExecutionId,omitempty"`
	// The information about the parameters.
	//
	// example:
	//
	// {"Status":"Running"}
	Parameters *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// The ID of the region.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfmxsn4m*****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The tags of the execution.
	//
	// example:
	//
	// {"k1": "v1", "k2": "v2"}
	Tags *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
}

func (s UpdateExecutionRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateExecutionRequest) GoString() string {
	return s.String()
}

func (s *UpdateExecutionRequest) SetClientToken(v string) *UpdateExecutionRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateExecutionRequest) SetDescription(v string) *UpdateExecutionRequest {
	s.Description = &v
	return s
}

func (s *UpdateExecutionRequest) SetExecutionId(v string) *UpdateExecutionRequest {
	s.ExecutionId = &v
	return s
}

func (s *UpdateExecutionRequest) SetParameters(v string) *UpdateExecutionRequest {
	s.Parameters = &v
	return s
}

func (s *UpdateExecutionRequest) SetRegionId(v string) *UpdateExecutionRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateExecutionRequest) SetResourceGroupId(v string) *UpdateExecutionRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *UpdateExecutionRequest) SetTags(v string) *UpdateExecutionRequest {
	s.Tags = &v
	return s
}

type UpdateExecutionResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// C8345E88-5334-469E-901D-F912C8CB9C55
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateExecutionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateExecutionResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateExecutionResponseBody) SetRequestId(v string) *UpdateExecutionResponseBody {
	s.RequestId = &v
	return s
}

type UpdateExecutionResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateExecutionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateExecutionResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateExecutionResponse) GoString() string {
	return s.String()
}

func (s *UpdateExecutionResponse) SetHeaders(v map[string]*string) *UpdateExecutionResponse {
	s.Headers = v
	return s
}

func (s *UpdateExecutionResponse) SetStatusCode(v int32) *UpdateExecutionResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateExecutionResponse) SetBody(v *UpdateExecutionResponseBody) *UpdateExecutionResponse {
	s.Body = v
	return s
}

type UpdateInstancePackageStateRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// install
	ConfigureAction *string `json:"ConfigureAction,omitempty" xml:"ConfigureAction,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// i-bp1jaxa2bs4bps7*****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// {"username": "xx"}
	Parameters map[string]interface{} `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ACS-ECS-InventoryDataCollection
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// example:
	//
	// v1
	TemplateVersion *string `json:"TemplateVersion,omitempty" xml:"TemplateVersion,omitempty"`
}

func (s UpdateInstancePackageStateRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateInstancePackageStateRequest) GoString() string {
	return s.String()
}

func (s *UpdateInstancePackageStateRequest) SetConfigureAction(v string) *UpdateInstancePackageStateRequest {
	s.ConfigureAction = &v
	return s
}

func (s *UpdateInstancePackageStateRequest) SetInstanceId(v string) *UpdateInstancePackageStateRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateInstancePackageStateRequest) SetParameters(v map[string]interface{}) *UpdateInstancePackageStateRequest {
	s.Parameters = v
	return s
}

func (s *UpdateInstancePackageStateRequest) SetRegionId(v string) *UpdateInstancePackageStateRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateInstancePackageStateRequest) SetTemplateName(v string) *UpdateInstancePackageStateRequest {
	s.TemplateName = &v
	return s
}

func (s *UpdateInstancePackageStateRequest) SetTemplateVersion(v string) *UpdateInstancePackageStateRequest {
	s.TemplateVersion = &v
	return s
}

type UpdateInstancePackageStateShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// install
	ConfigureAction *string `json:"ConfigureAction,omitempty" xml:"ConfigureAction,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// i-bp1jaxa2bs4bps7*****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// {"username": "xx"}
	ParametersShrink *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ACS-ECS-InventoryDataCollection
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// example:
	//
	// v1
	TemplateVersion *string `json:"TemplateVersion,omitempty" xml:"TemplateVersion,omitempty"`
}

func (s UpdateInstancePackageStateShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateInstancePackageStateShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateInstancePackageStateShrinkRequest) SetConfigureAction(v string) *UpdateInstancePackageStateShrinkRequest {
	s.ConfigureAction = &v
	return s
}

func (s *UpdateInstancePackageStateShrinkRequest) SetInstanceId(v string) *UpdateInstancePackageStateShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateInstancePackageStateShrinkRequest) SetParametersShrink(v string) *UpdateInstancePackageStateShrinkRequest {
	s.ParametersShrink = &v
	return s
}

func (s *UpdateInstancePackageStateShrinkRequest) SetRegionId(v string) *UpdateInstancePackageStateShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateInstancePackageStateShrinkRequest) SetTemplateName(v string) *UpdateInstancePackageStateShrinkRequest {
	s.TemplateName = &v
	return s
}

func (s *UpdateInstancePackageStateShrinkRequest) SetTemplateVersion(v string) *UpdateInstancePackageStateShrinkRequest {
	s.TemplateVersion = &v
	return s
}

type UpdateInstancePackageStateResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 2597E94B-5346-42D1-BB58-XXXXXXXXXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateInstancePackageStateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateInstancePackageStateResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateInstancePackageStateResponseBody) SetRequestId(v string) *UpdateInstancePackageStateResponseBody {
	s.RequestId = &v
	return s
}

type UpdateInstancePackageStateResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateInstancePackageStateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateInstancePackageStateResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateInstancePackageStateResponse) GoString() string {
	return s.String()
}

func (s *UpdateInstancePackageStateResponse) SetHeaders(v map[string]*string) *UpdateInstancePackageStateResponse {
	s.Headers = v
	return s
}

func (s *UpdateInstancePackageStateResponse) SetStatusCode(v int32) *UpdateInstancePackageStateResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateInstancePackageStateResponse) SetBody(v *UpdateInstancePackageStateResponseBody) *UpdateInstancePackageStateResponse {
	s.Body = v
	return s
}

type UpdateOpsItemRequest struct {
	// The category.
	//
	// example:
	//
	// Security
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// example:
	//
	// DASKJJLKADS-AHKLJHJSAKL-AJK
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The string to be deduplicated.
	//
	// example:
	//
	// ecs_instance_SystemMaintenance.Reboot
	DedupString *string `json:"DedupString,omitempty" xml:"DedupString,omitempty"`
	// The description of the O\\&M item.
	//
	// example:
	//
	// test-update
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the O\\&M item.
	//
	// example:
	//
	// oi-e2264dcf040c472598e9
	OpsItemId *string `json:"OpsItemId,omitempty" xml:"OpsItemId,omitempty"`
	// The priority.
	//
	// example:
	//
	// 2
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfmxsn4m4******
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The Alibaba Resource Names (ARNs) of the associated resources.
	//
	// example:
	//
	// [\\"arn:acs:ecs:cn-heyuan:1139354755361920:instance/i-f8z928h7aqotd3o65032\\"]
	Resources *string `json:"Resources,omitempty" xml:"Resources,omitempty"`
	// The severity level.
	//
	// example:
	//
	// Medium
	Severity *string `json:"Severity,omitempty" xml:"Severity,omitempty"`
	// The solutions.
	//
	// example:
	//
	// [{\\n \\\\"priority\\\\":3,\\n \\\\"type\\\\":\\\\"url\\\\",\\n \\\\"url\\\\":\\\\"https://example.com\\\\",\\n \\\\"description\\\\":\\\\"Specify a cross-zone high availability cluster. \\\\"\\n}]
	Solutions *string `json:"Solutions,omitempty" xml:"Solutions,omitempty"`
	// The source business.
	//
	// example:
	//
	// /aliyun/ecs
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The status.
	//
	// example:
	//
	// Open
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tags.
	//
	// example:
	//
	// {
	//
	//       "k1": "v1",
	//
	//       "k2": "v2"
	//
	// }
	Tags map[string]interface{} `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The title of the O\\&M item.
	//
	// example:
	//
	// Test
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s UpdateOpsItemRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateOpsItemRequest) GoString() string {
	return s.String()
}

func (s *UpdateOpsItemRequest) SetCategory(v string) *UpdateOpsItemRequest {
	s.Category = &v
	return s
}

func (s *UpdateOpsItemRequest) SetClientToken(v string) *UpdateOpsItemRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateOpsItemRequest) SetDedupString(v string) *UpdateOpsItemRequest {
	s.DedupString = &v
	return s
}

func (s *UpdateOpsItemRequest) SetDescription(v string) *UpdateOpsItemRequest {
	s.Description = &v
	return s
}

func (s *UpdateOpsItemRequest) SetOpsItemId(v string) *UpdateOpsItemRequest {
	s.OpsItemId = &v
	return s
}

func (s *UpdateOpsItemRequest) SetPriority(v int32) *UpdateOpsItemRequest {
	s.Priority = &v
	return s
}

func (s *UpdateOpsItemRequest) SetRegionId(v string) *UpdateOpsItemRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateOpsItemRequest) SetResourceGroupId(v string) *UpdateOpsItemRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *UpdateOpsItemRequest) SetResources(v string) *UpdateOpsItemRequest {
	s.Resources = &v
	return s
}

func (s *UpdateOpsItemRequest) SetSeverity(v string) *UpdateOpsItemRequest {
	s.Severity = &v
	return s
}

func (s *UpdateOpsItemRequest) SetSolutions(v string) *UpdateOpsItemRequest {
	s.Solutions = &v
	return s
}

func (s *UpdateOpsItemRequest) SetSource(v string) *UpdateOpsItemRequest {
	s.Source = &v
	return s
}

func (s *UpdateOpsItemRequest) SetStatus(v string) *UpdateOpsItemRequest {
	s.Status = &v
	return s
}

func (s *UpdateOpsItemRequest) SetTags(v map[string]interface{}) *UpdateOpsItemRequest {
	s.Tags = v
	return s
}

func (s *UpdateOpsItemRequest) SetTitle(v string) *UpdateOpsItemRequest {
	s.Title = &v
	return s
}

type UpdateOpsItemShrinkRequest struct {
	// The category.
	//
	// example:
	//
	// Security
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// example:
	//
	// DASKJJLKADS-AHKLJHJSAKL-AJK
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The string to be deduplicated.
	//
	// example:
	//
	// ecs_instance_SystemMaintenance.Reboot
	DedupString *string `json:"DedupString,omitempty" xml:"DedupString,omitempty"`
	// The description of the O\\&M item.
	//
	// example:
	//
	// test-update
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the O\\&M item.
	//
	// example:
	//
	// oi-e2264dcf040c472598e9
	OpsItemId *string `json:"OpsItemId,omitempty" xml:"OpsItemId,omitempty"`
	// The priority.
	//
	// example:
	//
	// 2
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfmxsn4m4******
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The Alibaba Resource Names (ARNs) of the associated resources.
	//
	// example:
	//
	// [\\"arn:acs:ecs:cn-heyuan:1139354755361920:instance/i-f8z928h7aqotd3o65032\\"]
	Resources *string `json:"Resources,omitempty" xml:"Resources,omitempty"`
	// The severity level.
	//
	// example:
	//
	// Medium
	Severity *string `json:"Severity,omitempty" xml:"Severity,omitempty"`
	// The solutions.
	//
	// example:
	//
	// [{\\n \\\\"priority\\\\":3,\\n \\\\"type\\\\":\\\\"url\\\\",\\n \\\\"url\\\\":\\\\"https://example.com\\\\",\\n \\\\"description\\\\":\\\\"Specify a cross-zone high availability cluster. \\\\"\\n}]
	Solutions *string `json:"Solutions,omitempty" xml:"Solutions,omitempty"`
	// The source business.
	//
	// example:
	//
	// /aliyun/ecs
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The status.
	//
	// example:
	//
	// Open
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tags.
	//
	// example:
	//
	// {
	//
	//       "k1": "v1",
	//
	//       "k2": "v2"
	//
	// }
	TagsShrink *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The title of the O\\&M item.
	//
	// example:
	//
	// Test
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s UpdateOpsItemShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateOpsItemShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateOpsItemShrinkRequest) SetCategory(v string) *UpdateOpsItemShrinkRequest {
	s.Category = &v
	return s
}

func (s *UpdateOpsItemShrinkRequest) SetClientToken(v string) *UpdateOpsItemShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateOpsItemShrinkRequest) SetDedupString(v string) *UpdateOpsItemShrinkRequest {
	s.DedupString = &v
	return s
}

func (s *UpdateOpsItemShrinkRequest) SetDescription(v string) *UpdateOpsItemShrinkRequest {
	s.Description = &v
	return s
}

func (s *UpdateOpsItemShrinkRequest) SetOpsItemId(v string) *UpdateOpsItemShrinkRequest {
	s.OpsItemId = &v
	return s
}

func (s *UpdateOpsItemShrinkRequest) SetPriority(v int32) *UpdateOpsItemShrinkRequest {
	s.Priority = &v
	return s
}

func (s *UpdateOpsItemShrinkRequest) SetRegionId(v string) *UpdateOpsItemShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateOpsItemShrinkRequest) SetResourceGroupId(v string) *UpdateOpsItemShrinkRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *UpdateOpsItemShrinkRequest) SetResources(v string) *UpdateOpsItemShrinkRequest {
	s.Resources = &v
	return s
}

func (s *UpdateOpsItemShrinkRequest) SetSeverity(v string) *UpdateOpsItemShrinkRequest {
	s.Severity = &v
	return s
}

func (s *UpdateOpsItemShrinkRequest) SetSolutions(v string) *UpdateOpsItemShrinkRequest {
	s.Solutions = &v
	return s
}

func (s *UpdateOpsItemShrinkRequest) SetSource(v string) *UpdateOpsItemShrinkRequest {
	s.Source = &v
	return s
}

func (s *UpdateOpsItemShrinkRequest) SetStatus(v string) *UpdateOpsItemShrinkRequest {
	s.Status = &v
	return s
}

func (s *UpdateOpsItemShrinkRequest) SetTagsShrink(v string) *UpdateOpsItemShrinkRequest {
	s.TagsShrink = &v
	return s
}

func (s *UpdateOpsItemShrinkRequest) SetTitle(v string) *UpdateOpsItemShrinkRequest {
	s.Title = &v
	return s
}

type UpdateOpsItemResponseBody struct {
	// The information about the O\\&M item.
	OpsItem *UpdateOpsItemResponseBodyOpsItem `json:"OpsItem,omitempty" xml:"OpsItem,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// C996DECB-3D2B-5321-B359-BE7031B6399E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateOpsItemResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateOpsItemResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateOpsItemResponseBody) SetOpsItem(v *UpdateOpsItemResponseBodyOpsItem) *UpdateOpsItemResponseBody {
	s.OpsItem = v
	return s
}

func (s *UpdateOpsItemResponseBody) SetRequestId(v string) *UpdateOpsItemResponseBody {
	s.RequestId = &v
	return s
}

type UpdateOpsItemResponseBodyOpsItem struct {
	// The attributes of the O\\&M item.
	//
	// example:
	//
	// [{\\"Attribute\\": {\\"Weight\\": 100}, \\"RealServer\\": \\"uaejc8hnqzqz5valyh8dibolpvza48ik.yundunwaf5.com\\"}]
	Attributes *string `json:"Attributes,omitempty" xml:"Attributes,omitempty"`
	// The category.
	//
	// example:
	//
	// Security
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The time when the O\\&M item was created.
	//
	// example:
	//
	// 2023-03-16T07:04Z
	CreateDate *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	// The user who created the patch baseline.
	//
	// example:
	//
	// root(130900000)
	CreatedBy *string `json:"CreatedBy,omitempty" xml:"CreatedBy,omitempty"`
	// The description.
	//
	// example:
	//
	// test-update
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The user who modified the O\\&M item.
	//
	// example:
	//
	// root(130900000)
	LastModifiedBy *string `json:"LastModifiedBy,omitempty" xml:"LastModifiedBy,omitempty"`
	// The ID of the O\\&M item.
	//
	// example:
	//
	// oi-e2264dcf040c472598e9
	OpsItemId *string `json:"OpsItemId,omitempty" xml:"OpsItemId,omitempty"`
	// The priority.
	//
	// example:
	//
	// 2
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfmxsn4m4******
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The ARNs of the associated resources.
	Resources []*string `json:"Resources,omitempty" xml:"Resources,omitempty" type:"Repeated"`
	// The severity level.
	//
	// example:
	//
	// Medium
	Severity *string `json:"Severity,omitempty" xml:"Severity,omitempty"`
	// The solutions.
	Solutions []*string `json:"Solutions,omitempty" xml:"Solutions,omitempty" type:"Repeated"`
	// The source business.
	//
	// example:
	//
	// /aliyun/ecs
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The status.
	//
	// example:
	//
	// Open
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tags.
	//
	// example:
	//
	// {
	//
	//       "k1": "v1",
	//
	//       "k2": "v2"
	//
	// }
	Tags map[string]interface{} `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The title of the O\\&M item.
	//
	// example:
	//
	// Test
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// The time when the O\\&M item was updated.
	//
	// example:
	//
	// 2023-03-16T08:04Z
	UpdateDate *string `json:"UpdateDate,omitempty" xml:"UpdateDate,omitempty"`
}

func (s UpdateOpsItemResponseBodyOpsItem) String() string {
	return tea.Prettify(s)
}

func (s UpdateOpsItemResponseBodyOpsItem) GoString() string {
	return s.String()
}

func (s *UpdateOpsItemResponseBodyOpsItem) SetAttributes(v string) *UpdateOpsItemResponseBodyOpsItem {
	s.Attributes = &v
	return s
}

func (s *UpdateOpsItemResponseBodyOpsItem) SetCategory(v string) *UpdateOpsItemResponseBodyOpsItem {
	s.Category = &v
	return s
}

func (s *UpdateOpsItemResponseBodyOpsItem) SetCreateDate(v string) *UpdateOpsItemResponseBodyOpsItem {
	s.CreateDate = &v
	return s
}

func (s *UpdateOpsItemResponseBodyOpsItem) SetCreatedBy(v string) *UpdateOpsItemResponseBodyOpsItem {
	s.CreatedBy = &v
	return s
}

func (s *UpdateOpsItemResponseBodyOpsItem) SetDescription(v string) *UpdateOpsItemResponseBodyOpsItem {
	s.Description = &v
	return s
}

func (s *UpdateOpsItemResponseBodyOpsItem) SetLastModifiedBy(v string) *UpdateOpsItemResponseBodyOpsItem {
	s.LastModifiedBy = &v
	return s
}

func (s *UpdateOpsItemResponseBodyOpsItem) SetOpsItemId(v string) *UpdateOpsItemResponseBodyOpsItem {
	s.OpsItemId = &v
	return s
}

func (s *UpdateOpsItemResponseBodyOpsItem) SetPriority(v int32) *UpdateOpsItemResponseBodyOpsItem {
	s.Priority = &v
	return s
}

func (s *UpdateOpsItemResponseBodyOpsItem) SetResourceGroupId(v string) *UpdateOpsItemResponseBodyOpsItem {
	s.ResourceGroupId = &v
	return s
}

func (s *UpdateOpsItemResponseBodyOpsItem) SetResources(v []*string) *UpdateOpsItemResponseBodyOpsItem {
	s.Resources = v
	return s
}

func (s *UpdateOpsItemResponseBodyOpsItem) SetSeverity(v string) *UpdateOpsItemResponseBodyOpsItem {
	s.Severity = &v
	return s
}

func (s *UpdateOpsItemResponseBodyOpsItem) SetSolutions(v []*string) *UpdateOpsItemResponseBodyOpsItem {
	s.Solutions = v
	return s
}

func (s *UpdateOpsItemResponseBodyOpsItem) SetSource(v string) *UpdateOpsItemResponseBodyOpsItem {
	s.Source = &v
	return s
}

func (s *UpdateOpsItemResponseBodyOpsItem) SetStatus(v string) *UpdateOpsItemResponseBodyOpsItem {
	s.Status = &v
	return s
}

func (s *UpdateOpsItemResponseBodyOpsItem) SetTags(v map[string]interface{}) *UpdateOpsItemResponseBodyOpsItem {
	s.Tags = v
	return s
}

func (s *UpdateOpsItemResponseBodyOpsItem) SetTitle(v string) *UpdateOpsItemResponseBodyOpsItem {
	s.Title = &v
	return s
}

func (s *UpdateOpsItemResponseBodyOpsItem) SetUpdateDate(v string) *UpdateOpsItemResponseBodyOpsItem {
	s.UpdateDate = &v
	return s
}

type UpdateOpsItemResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateOpsItemResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateOpsItemResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateOpsItemResponse) GoString() string {
	return s.String()
}

func (s *UpdateOpsItemResponse) SetHeaders(v map[string]*string) *UpdateOpsItemResponse {
	s.Headers = v
	return s
}

func (s *UpdateOpsItemResponse) SetStatusCode(v int32) *UpdateOpsItemResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateOpsItemResponse) SetBody(v *UpdateOpsItemResponseBody) *UpdateOpsItemResponse {
	s.Body = v
	return s
}

type UpdateParameterRequest struct {
	// The description of the common parameter. The description must be 1 to 200 characters in length.
	//
	// example:
	//
	// update
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the common parameter. The name must be 1 to 200 characters in length, and can contain letters, digits, hyphens (-), and underscores (_). It cannot start with ALIYUN, ACS, ALIBABA, ALICLOUD, or OOS.
	//
	// This parameter is required.
	//
	// example:
	//
	// MyParameter
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfmxsn4m4******
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The tags to be added to common parameter.
	//
	// example:
	//
	// {"k1": "v1", "k2": "v2"}
	Tags *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The value of the common parameter. The value must be 1 to 4,096 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// update
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateParameterRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateParameterRequest) GoString() string {
	return s.String()
}

func (s *UpdateParameterRequest) SetDescription(v string) *UpdateParameterRequest {
	s.Description = &v
	return s
}

func (s *UpdateParameterRequest) SetName(v string) *UpdateParameterRequest {
	s.Name = &v
	return s
}

func (s *UpdateParameterRequest) SetRegionId(v string) *UpdateParameterRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateParameterRequest) SetResourceGroupId(v string) *UpdateParameterRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *UpdateParameterRequest) SetTags(v string) *UpdateParameterRequest {
	s.Tags = &v
	return s
}

func (s *UpdateParameterRequest) SetValue(v string) *UpdateParameterRequest {
	s.Value = &v
	return s
}

type UpdateParameterResponseBody struct {
	// The information about the common parameter.
	Parameter *UpdateParameterResponseBodyParameter `json:"Parameter,omitempty" xml:"Parameter,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// AF1AE6DE-61C4-435E-8687-072CFACCCEC7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateParameterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateParameterResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateParameterResponseBody) SetParameter(v *UpdateParameterResponseBodyParameter) *UpdateParameterResponseBody {
	s.Parameter = v
	return s
}

func (s *UpdateParameterResponseBody) SetRequestId(v string) *UpdateParameterResponseBody {
	s.RequestId = &v
	return s
}

type UpdateParameterResponseBodyParameter struct {
	// The constraints of the common parameter.
	//
	// example:
	//
	// "{\\"AllowedValues\\":[\\"parameter\\"],\\"AllowedPattern\\":\\"parameter\\",\\"MinLength\\":0,\\"MaxLength\\":20}"
	Constraints *string `json:"Constraints,omitempty" xml:"Constraints,omitempty"`
	// The user who created the common parameter.
	//
	// example:
	//
	// root(130900000)
	CreatedBy *string `json:"CreatedBy,omitempty" xml:"CreatedBy,omitempty"`
	// The time when the common parameter was created.
	//
	// example:
	//
	// 2020-09-01T08:01:43Z
	CreatedDate *string `json:"CreatedDate,omitempty" xml:"CreatedDate,omitempty"`
	// The description of the common parameter.
	//
	// example:
	//
	// update
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The parameter ID.
	//
	// example:
	//
	// p-4c4b401cab6747xxxxxx
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the common parameter.
	//
	// example:
	//
	// MyParameter
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The version number of the common parameter.
	//
	// example:
	//
	// 2
	ParameterVersion *int32 `json:"ParameterVersion,omitempty" xml:"ParameterVersion,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfmxsn4m4******
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The share type of the common parameter.
	//
	// example:
	//
	// Private
	ShareType *string `json:"ShareType,omitempty" xml:"ShareType,omitempty"`
	// The tag added to the common parameter.
	//
	// example:
	//
	// {"k1": "v1", "k2": "v2"}
	Tags *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The data type of the common parameter.
	//
	// example:
	//
	// String
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The user who updated the common parameter.
	//
	// example:
	//
	// root(130900000)
	UpdatedBy *string `json:"UpdatedBy,omitempty" xml:"UpdatedBy,omitempty"`
	// The time when the common parameter was updated.
	//
	// example:
	//
	// 2020-09-01T08:04:23Z
	UpdatedDate *string `json:"UpdatedDate,omitempty" xml:"UpdatedDate,omitempty"`
}

func (s UpdateParameterResponseBodyParameter) String() string {
	return tea.Prettify(s)
}

func (s UpdateParameterResponseBodyParameter) GoString() string {
	return s.String()
}

func (s *UpdateParameterResponseBodyParameter) SetConstraints(v string) *UpdateParameterResponseBodyParameter {
	s.Constraints = &v
	return s
}

func (s *UpdateParameterResponseBodyParameter) SetCreatedBy(v string) *UpdateParameterResponseBodyParameter {
	s.CreatedBy = &v
	return s
}

func (s *UpdateParameterResponseBodyParameter) SetCreatedDate(v string) *UpdateParameterResponseBodyParameter {
	s.CreatedDate = &v
	return s
}

func (s *UpdateParameterResponseBodyParameter) SetDescription(v string) *UpdateParameterResponseBodyParameter {
	s.Description = &v
	return s
}

func (s *UpdateParameterResponseBodyParameter) SetId(v string) *UpdateParameterResponseBodyParameter {
	s.Id = &v
	return s
}

func (s *UpdateParameterResponseBodyParameter) SetName(v string) *UpdateParameterResponseBodyParameter {
	s.Name = &v
	return s
}

func (s *UpdateParameterResponseBodyParameter) SetParameterVersion(v int32) *UpdateParameterResponseBodyParameter {
	s.ParameterVersion = &v
	return s
}

func (s *UpdateParameterResponseBodyParameter) SetResourceGroupId(v string) *UpdateParameterResponseBodyParameter {
	s.ResourceGroupId = &v
	return s
}

func (s *UpdateParameterResponseBodyParameter) SetShareType(v string) *UpdateParameterResponseBodyParameter {
	s.ShareType = &v
	return s
}

func (s *UpdateParameterResponseBodyParameter) SetTags(v string) *UpdateParameterResponseBodyParameter {
	s.Tags = &v
	return s
}

func (s *UpdateParameterResponseBodyParameter) SetType(v string) *UpdateParameterResponseBodyParameter {
	s.Type = &v
	return s
}

func (s *UpdateParameterResponseBodyParameter) SetUpdatedBy(v string) *UpdateParameterResponseBodyParameter {
	s.UpdatedBy = &v
	return s
}

func (s *UpdateParameterResponseBodyParameter) SetUpdatedDate(v string) *UpdateParameterResponseBodyParameter {
	s.UpdatedDate = &v
	return s
}

type UpdateParameterResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateParameterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateParameterResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateParameterResponse) GoString() string {
	return s.String()
}

func (s *UpdateParameterResponse) SetHeaders(v map[string]*string) *UpdateParameterResponse {
	s.Headers = v
	return s
}

func (s *UpdateParameterResponse) SetStatusCode(v int32) *UpdateParameterResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateParameterResponse) SetBody(v *UpdateParameterResponseBody) *UpdateParameterResponse {
	s.Body = v
	return s
}

type UpdatePatchBaselineRequest struct {
	// The rules of scanning and installing patches for the specified operating system.
	//
	// example:
	//
	// {"PatchRules":[{"PatchFilterGroup":[{"Key":"PatchSet","Values":["OS"]},{"Key":"ProductFamily","Values":["Windows"]},{"Key":"Product","Values":["Windows 10","Windows 7"]},{"Key":"Classification","Values":["Security Updates","Updates","Update Rollups","Critical Updates"]},{"Key":"Severity","Values":["Critical","Important","Moderate"]}],"ApproveAfterDays":7,"ApproveUntilDate":"","EnableNonSecurity":true,"ComplianceLevel":"Medium"}]}
	ApprovalRules *string `json:"ApprovalRules,omitempty" xml:"ApprovalRules,omitempty"`
	// The approved patches.
	ApprovedPatches []*string `json:"ApprovedPatches,omitempty" xml:"ApprovedPatches,omitempty" type:"Repeated"`
	// Indicates whether the approved patch involves updates other than security-related updates.
	//
	// example:
	//
	// true
	ApprovedPatchesEnableNonSecurity *bool `json:"ApprovedPatchesEnableNonSecurity,omitempty" xml:"ApprovedPatchesEnableNonSecurity,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// example:
	//
	// -
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The description of the patch baseline.
	//
	// example:
	//
	// UpdatePatchBaseline
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the patch baseline.
	//
	// This parameter is required.
	//
	// example:
	//
	// MyPatchBaseline
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the region.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The rejected patches.
	RejectedPatches []*string `json:"RejectedPatches,omitempty" xml:"RejectedPatches,omitempty" type:"Repeated"`
	// The action of the rejected patch.
	//
	// example:
	//
	// ALLOW_AS_DEPENDENCY
	RejectedPatchesAction *string `json:"RejectedPatchesAction,omitempty" xml:"RejectedPatchesAction,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfmxsn4m4******
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The patch source configurations.
	Sources []*string `json:"Sources,omitempty" xml:"Sources,omitempty" type:"Repeated"`
	// The tags.
	Tags []*UpdatePatchBaselineRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s UpdatePatchBaselineRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdatePatchBaselineRequest) GoString() string {
	return s.String()
}

func (s *UpdatePatchBaselineRequest) SetApprovalRules(v string) *UpdatePatchBaselineRequest {
	s.ApprovalRules = &v
	return s
}

func (s *UpdatePatchBaselineRequest) SetApprovedPatches(v []*string) *UpdatePatchBaselineRequest {
	s.ApprovedPatches = v
	return s
}

func (s *UpdatePatchBaselineRequest) SetApprovedPatchesEnableNonSecurity(v bool) *UpdatePatchBaselineRequest {
	s.ApprovedPatchesEnableNonSecurity = &v
	return s
}

func (s *UpdatePatchBaselineRequest) SetClientToken(v string) *UpdatePatchBaselineRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdatePatchBaselineRequest) SetDescription(v string) *UpdatePatchBaselineRequest {
	s.Description = &v
	return s
}

func (s *UpdatePatchBaselineRequest) SetName(v string) *UpdatePatchBaselineRequest {
	s.Name = &v
	return s
}

func (s *UpdatePatchBaselineRequest) SetRegionId(v string) *UpdatePatchBaselineRequest {
	s.RegionId = &v
	return s
}

func (s *UpdatePatchBaselineRequest) SetRejectedPatches(v []*string) *UpdatePatchBaselineRequest {
	s.RejectedPatches = v
	return s
}

func (s *UpdatePatchBaselineRequest) SetRejectedPatchesAction(v string) *UpdatePatchBaselineRequest {
	s.RejectedPatchesAction = &v
	return s
}

func (s *UpdatePatchBaselineRequest) SetResourceGroupId(v string) *UpdatePatchBaselineRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *UpdatePatchBaselineRequest) SetSources(v []*string) *UpdatePatchBaselineRequest {
	s.Sources = v
	return s
}

func (s *UpdatePatchBaselineRequest) SetTags(v []*UpdatePatchBaselineRequestTags) *UpdatePatchBaselineRequest {
	s.Tags = v
	return s
}

type UpdatePatchBaselineRequestTags struct {
	// The key of the tag.
	//
	// example:
	//
	// key
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag.
	//
	// example:
	//
	// value
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdatePatchBaselineRequestTags) String() string {
	return tea.Prettify(s)
}

func (s UpdatePatchBaselineRequestTags) GoString() string {
	return s.String()
}

func (s *UpdatePatchBaselineRequestTags) SetKey(v string) *UpdatePatchBaselineRequestTags {
	s.Key = &v
	return s
}

func (s *UpdatePatchBaselineRequestTags) SetValue(v string) *UpdatePatchBaselineRequestTags {
	s.Value = &v
	return s
}

type UpdatePatchBaselineShrinkRequest struct {
	// The rules of scanning and installing patches for the specified operating system.
	//
	// example:
	//
	// {"PatchRules":[{"PatchFilterGroup":[{"Key":"PatchSet","Values":["OS"]},{"Key":"ProductFamily","Values":["Windows"]},{"Key":"Product","Values":["Windows 10","Windows 7"]},{"Key":"Classification","Values":["Security Updates","Updates","Update Rollups","Critical Updates"]},{"Key":"Severity","Values":["Critical","Important","Moderate"]}],"ApproveAfterDays":7,"ApproveUntilDate":"","EnableNonSecurity":true,"ComplianceLevel":"Medium"}]}
	ApprovalRules *string `json:"ApprovalRules,omitempty" xml:"ApprovalRules,omitempty"`
	// The approved patches.
	ApprovedPatchesShrink *string `json:"ApprovedPatches,omitempty" xml:"ApprovedPatches,omitempty"`
	// Indicates whether the approved patch involves updates other than security-related updates.
	//
	// example:
	//
	// true
	ApprovedPatchesEnableNonSecurity *bool `json:"ApprovedPatchesEnableNonSecurity,omitempty" xml:"ApprovedPatchesEnableNonSecurity,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// example:
	//
	// -
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The description of the patch baseline.
	//
	// example:
	//
	// UpdatePatchBaseline
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the patch baseline.
	//
	// This parameter is required.
	//
	// example:
	//
	// MyPatchBaseline
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the region.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The rejected patches.
	RejectedPatchesShrink *string `json:"RejectedPatches,omitempty" xml:"RejectedPatches,omitempty"`
	// The action of the rejected patch.
	//
	// example:
	//
	// ALLOW_AS_DEPENDENCY
	RejectedPatchesAction *string `json:"RejectedPatchesAction,omitempty" xml:"RejectedPatchesAction,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfmxsn4m4******
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The patch source configurations.
	SourcesShrink *string `json:"Sources,omitempty" xml:"Sources,omitempty"`
	// The tags.
	TagsShrink *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
}

func (s UpdatePatchBaselineShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdatePatchBaselineShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdatePatchBaselineShrinkRequest) SetApprovalRules(v string) *UpdatePatchBaselineShrinkRequest {
	s.ApprovalRules = &v
	return s
}

func (s *UpdatePatchBaselineShrinkRequest) SetApprovedPatchesShrink(v string) *UpdatePatchBaselineShrinkRequest {
	s.ApprovedPatchesShrink = &v
	return s
}

func (s *UpdatePatchBaselineShrinkRequest) SetApprovedPatchesEnableNonSecurity(v bool) *UpdatePatchBaselineShrinkRequest {
	s.ApprovedPatchesEnableNonSecurity = &v
	return s
}

func (s *UpdatePatchBaselineShrinkRequest) SetClientToken(v string) *UpdatePatchBaselineShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdatePatchBaselineShrinkRequest) SetDescription(v string) *UpdatePatchBaselineShrinkRequest {
	s.Description = &v
	return s
}

func (s *UpdatePatchBaselineShrinkRequest) SetName(v string) *UpdatePatchBaselineShrinkRequest {
	s.Name = &v
	return s
}

func (s *UpdatePatchBaselineShrinkRequest) SetRegionId(v string) *UpdatePatchBaselineShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *UpdatePatchBaselineShrinkRequest) SetRejectedPatchesShrink(v string) *UpdatePatchBaselineShrinkRequest {
	s.RejectedPatchesShrink = &v
	return s
}

func (s *UpdatePatchBaselineShrinkRequest) SetRejectedPatchesAction(v string) *UpdatePatchBaselineShrinkRequest {
	s.RejectedPatchesAction = &v
	return s
}

func (s *UpdatePatchBaselineShrinkRequest) SetResourceGroupId(v string) *UpdatePatchBaselineShrinkRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *UpdatePatchBaselineShrinkRequest) SetSourcesShrink(v string) *UpdatePatchBaselineShrinkRequest {
	s.SourcesShrink = &v
	return s
}

func (s *UpdatePatchBaselineShrinkRequest) SetTagsShrink(v string) *UpdatePatchBaselineShrinkRequest {
	s.TagsShrink = &v
	return s
}

type UpdatePatchBaselineResponseBody struct {
	// The details of the patch baseline.
	PatchBaseline *UpdatePatchBaselineResponseBodyPatchBaseline `json:"PatchBaseline,omitempty" xml:"PatchBaseline,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 1457F46C-7AAE-59FA-BD12-0BDB3751E6F8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdatePatchBaselineResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdatePatchBaselineResponseBody) GoString() string {
	return s.String()
}

func (s *UpdatePatchBaselineResponseBody) SetPatchBaseline(v *UpdatePatchBaselineResponseBodyPatchBaseline) *UpdatePatchBaselineResponseBody {
	s.PatchBaseline = v
	return s
}

func (s *UpdatePatchBaselineResponseBody) SetRequestId(v string) *UpdatePatchBaselineResponseBody {
	s.RequestId = &v
	return s
}

type UpdatePatchBaselineResponseBodyPatchBaseline struct {
	// The rules of scanning and installing patches for the specified operating system.
	//
	// example:
	//
	// {"PatchRules":[{"PatchFilterGroup":[{"Key":"PatchSet","Values":["OS"]},{"Key":"ProductFamily","Values":["Windows"]},{"Key":"Product","Values":["Windows 10","Windows 7"]},{"Key":"Classification","Values":["Security Updates","Updates","Update Rollups","Critical Updates"]},{"Key":"Severity","Values":["Critical","Important","Moderate"]}],"ApproveAfterDays":7,"ApproveUntilDate":"","EnableNonSecurity":true,"ComplianceLevel":"Medium"}]}
	ApprovalRules *string `json:"ApprovalRules,omitempty" xml:"ApprovalRules,omitempty"`
	// The approved patches.
	ApprovedPatches []*string `json:"ApprovedPatches,omitempty" xml:"ApprovedPatches,omitempty" type:"Repeated"`
	// Indicates whether the approved patch involves updates other than security-related updates.
	//
	// example:
	//
	// true
	ApprovedPatchesEnableNonSecurity *bool `json:"ApprovedPatchesEnableNonSecurity,omitempty" xml:"ApprovedPatchesEnableNonSecurity,omitempty"`
	// The creator of the patch baseline.
	//
	// example:
	//
	// root(130900000)
	CreatedBy *string `json:"CreatedBy,omitempty" xml:"CreatedBy,omitempty"`
	// The time when the patch baseline was created.
	//
	// example:
	//
	// 2021-09-07T03:42:56Z
	CreatedDate *string `json:"CreatedDate,omitempty" xml:"CreatedDate,omitempty"`
	// The description of the patch baseline.
	//
	// example:
	//
	// UpdatePatchBaseline
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the patch baseline.
	//
	// example:
	//
	// pb-445340b5c6504a85a300
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the patch baseline.
	//
	// example:
	//
	// MyPatchBaseline
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The operating system.
	//
	// example:
	//
	// Windows
	OperationSystem *string `json:"OperationSystem,omitempty" xml:"OperationSystem,omitempty"`
	// The rejected patches.
	RejectedPatches []*string `json:"RejectedPatches,omitempty" xml:"RejectedPatches,omitempty" type:"Repeated"`
	// The action of the rejected patch.
	//
	// example:
	//
	// ALLOW_AS_DEPENDENCY
	RejectedPatchesAction *string `json:"RejectedPatchesAction,omitempty" xml:"RejectedPatchesAction,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfmy2zdbbjplii
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The share type of the patch baseline.
	//
	// example:
	//
	// Private
	ShareType *string `json:"ShareType,omitempty" xml:"ShareType,omitempty"`
	// The patch source configurations.
	Sources []*string `json:"Sources,omitempty" xml:"Sources,omitempty" type:"Repeated"`
	// The tags.
	Tags []*UpdatePatchBaselineResponseBodyPatchBaselineTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The user who updated the patch baseline.
	//
	// example:
	//
	// root(130900000)
	UpdatedBy *string `json:"UpdatedBy,omitempty" xml:"UpdatedBy,omitempty"`
	// The time when the patch baseline was updated.
	//
	// example:
	//
	// 2021-09-08T07:26:37Z
	UpdatedDate *string `json:"UpdatedDate,omitempty" xml:"UpdatedDate,omitempty"`
}

func (s UpdatePatchBaselineResponseBodyPatchBaseline) String() string {
	return tea.Prettify(s)
}

func (s UpdatePatchBaselineResponseBodyPatchBaseline) GoString() string {
	return s.String()
}

func (s *UpdatePatchBaselineResponseBodyPatchBaseline) SetApprovalRules(v string) *UpdatePatchBaselineResponseBodyPatchBaseline {
	s.ApprovalRules = &v
	return s
}

func (s *UpdatePatchBaselineResponseBodyPatchBaseline) SetApprovedPatches(v []*string) *UpdatePatchBaselineResponseBodyPatchBaseline {
	s.ApprovedPatches = v
	return s
}

func (s *UpdatePatchBaselineResponseBodyPatchBaseline) SetApprovedPatchesEnableNonSecurity(v bool) *UpdatePatchBaselineResponseBodyPatchBaseline {
	s.ApprovedPatchesEnableNonSecurity = &v
	return s
}

func (s *UpdatePatchBaselineResponseBodyPatchBaseline) SetCreatedBy(v string) *UpdatePatchBaselineResponseBodyPatchBaseline {
	s.CreatedBy = &v
	return s
}

func (s *UpdatePatchBaselineResponseBodyPatchBaseline) SetCreatedDate(v string) *UpdatePatchBaselineResponseBodyPatchBaseline {
	s.CreatedDate = &v
	return s
}

func (s *UpdatePatchBaselineResponseBodyPatchBaseline) SetDescription(v string) *UpdatePatchBaselineResponseBodyPatchBaseline {
	s.Description = &v
	return s
}

func (s *UpdatePatchBaselineResponseBodyPatchBaseline) SetId(v string) *UpdatePatchBaselineResponseBodyPatchBaseline {
	s.Id = &v
	return s
}

func (s *UpdatePatchBaselineResponseBodyPatchBaseline) SetName(v string) *UpdatePatchBaselineResponseBodyPatchBaseline {
	s.Name = &v
	return s
}

func (s *UpdatePatchBaselineResponseBodyPatchBaseline) SetOperationSystem(v string) *UpdatePatchBaselineResponseBodyPatchBaseline {
	s.OperationSystem = &v
	return s
}

func (s *UpdatePatchBaselineResponseBodyPatchBaseline) SetRejectedPatches(v []*string) *UpdatePatchBaselineResponseBodyPatchBaseline {
	s.RejectedPatches = v
	return s
}

func (s *UpdatePatchBaselineResponseBodyPatchBaseline) SetRejectedPatchesAction(v string) *UpdatePatchBaselineResponseBodyPatchBaseline {
	s.RejectedPatchesAction = &v
	return s
}

func (s *UpdatePatchBaselineResponseBodyPatchBaseline) SetResourceGroupId(v string) *UpdatePatchBaselineResponseBodyPatchBaseline {
	s.ResourceGroupId = &v
	return s
}

func (s *UpdatePatchBaselineResponseBodyPatchBaseline) SetShareType(v string) *UpdatePatchBaselineResponseBodyPatchBaseline {
	s.ShareType = &v
	return s
}

func (s *UpdatePatchBaselineResponseBodyPatchBaseline) SetSources(v []*string) *UpdatePatchBaselineResponseBodyPatchBaseline {
	s.Sources = v
	return s
}

func (s *UpdatePatchBaselineResponseBodyPatchBaseline) SetTags(v []*UpdatePatchBaselineResponseBodyPatchBaselineTags) *UpdatePatchBaselineResponseBodyPatchBaseline {
	s.Tags = v
	return s
}

func (s *UpdatePatchBaselineResponseBodyPatchBaseline) SetUpdatedBy(v string) *UpdatePatchBaselineResponseBodyPatchBaseline {
	s.UpdatedBy = &v
	return s
}

func (s *UpdatePatchBaselineResponseBodyPatchBaseline) SetUpdatedDate(v string) *UpdatePatchBaselineResponseBodyPatchBaseline {
	s.UpdatedDate = &v
	return s
}

type UpdatePatchBaselineResponseBodyPatchBaselineTags struct {
	// The key of the tag.
	//
	// example:
	//
	// key
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The value of the tag.
	//
	// example:
	//
	// value
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s UpdatePatchBaselineResponseBodyPatchBaselineTags) String() string {
	return tea.Prettify(s)
}

func (s UpdatePatchBaselineResponseBodyPatchBaselineTags) GoString() string {
	return s.String()
}

func (s *UpdatePatchBaselineResponseBodyPatchBaselineTags) SetTagKey(v string) *UpdatePatchBaselineResponseBodyPatchBaselineTags {
	s.TagKey = &v
	return s
}

func (s *UpdatePatchBaselineResponseBodyPatchBaselineTags) SetTagValue(v string) *UpdatePatchBaselineResponseBodyPatchBaselineTags {
	s.TagValue = &v
	return s
}

type UpdatePatchBaselineResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdatePatchBaselineResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdatePatchBaselineResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdatePatchBaselineResponse) GoString() string {
	return s.String()
}

func (s *UpdatePatchBaselineResponse) SetHeaders(v map[string]*string) *UpdatePatchBaselineResponse {
	s.Headers = v
	return s
}

func (s *UpdatePatchBaselineResponse) SetStatusCode(v int32) *UpdatePatchBaselineResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdatePatchBaselineResponse) SetBody(v *UpdatePatchBaselineResponseBody) *UpdatePatchBaselineResponse {
	s.Body = v
	return s
}

type UpdateSecretParameterRequest struct {
	// The description of the parameter. The description must be 1 to 200 characters in length.
	//
	// example:
	//
	// SecretParameter
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the parameter. The name must be 1 to 180 characters in length, and can contain letters, digits, hyphens (-), and underscores (_). It cannot start with ALIYUN, ACS, ALIBABA, ALICLOUD, or OOS.
	//
	// This parameter is required.
	//
	// example:
	//
	// MySecretParameter
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the region.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfmxsn4m4******
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The tags of the parameter.
	//
	// example:
	//
	// {"k1": "v1", "k2": "v2"}
	Tags map[string]interface{} `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The value of the parameter. The value must be 1 to 4096 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// update
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateSecretParameterRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateSecretParameterRequest) GoString() string {
	return s.String()
}

func (s *UpdateSecretParameterRequest) SetDescription(v string) *UpdateSecretParameterRequest {
	s.Description = &v
	return s
}

func (s *UpdateSecretParameterRequest) SetName(v string) *UpdateSecretParameterRequest {
	s.Name = &v
	return s
}

func (s *UpdateSecretParameterRequest) SetRegionId(v string) *UpdateSecretParameterRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateSecretParameterRequest) SetResourceGroupId(v string) *UpdateSecretParameterRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *UpdateSecretParameterRequest) SetTags(v map[string]interface{}) *UpdateSecretParameterRequest {
	s.Tags = v
	return s
}

func (s *UpdateSecretParameterRequest) SetValue(v string) *UpdateSecretParameterRequest {
	s.Value = &v
	return s
}

type UpdateSecretParameterShrinkRequest struct {
	// The description of the parameter. The description must be 1 to 200 characters in length.
	//
	// example:
	//
	// SecretParameter
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the parameter. The name must be 1 to 180 characters in length, and can contain letters, digits, hyphens (-), and underscores (_). It cannot start with ALIYUN, ACS, ALIBABA, ALICLOUD, or OOS.
	//
	// This parameter is required.
	//
	// example:
	//
	// MySecretParameter
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the region.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfmxsn4m4******
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The tags of the parameter.
	//
	// example:
	//
	// {"k1": "v1", "k2": "v2"}
	TagsShrink *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The value of the parameter. The value must be 1 to 4096 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// update
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateSecretParameterShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateSecretParameterShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateSecretParameterShrinkRequest) SetDescription(v string) *UpdateSecretParameterShrinkRequest {
	s.Description = &v
	return s
}

func (s *UpdateSecretParameterShrinkRequest) SetName(v string) *UpdateSecretParameterShrinkRequest {
	s.Name = &v
	return s
}

func (s *UpdateSecretParameterShrinkRequest) SetRegionId(v string) *UpdateSecretParameterShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateSecretParameterShrinkRequest) SetResourceGroupId(v string) *UpdateSecretParameterShrinkRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *UpdateSecretParameterShrinkRequest) SetTagsShrink(v string) *UpdateSecretParameterShrinkRequest {
	s.TagsShrink = &v
	return s
}

func (s *UpdateSecretParameterShrinkRequest) SetValue(v string) *UpdateSecretParameterShrinkRequest {
	s.Value = &v
	return s
}

type UpdateSecretParameterResponseBody struct {
	// The information about the parameter.
	Parameter *UpdateSecretParameterResponseBodyParameter `json:"Parameter,omitempty" xml:"Parameter,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 0B419FF3-ABC6-4DF0-95E5-636DC8CBB8AF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateSecretParameterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateSecretParameterResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateSecretParameterResponseBody) SetParameter(v *UpdateSecretParameterResponseBodyParameter) *UpdateSecretParameterResponseBody {
	s.Parameter = v
	return s
}

func (s *UpdateSecretParameterResponseBody) SetRequestId(v string) *UpdateSecretParameterResponseBody {
	s.RequestId = &v
	return s
}

type UpdateSecretParameterResponseBodyParameter struct {
	// The constraints of the parameter.
	//
	// example:
	//
	// \\"{\\"\\"AllowedValues":["secretparameter"],"AllowedPattern":".*","MinLength":0,"MaxLength":20}\\"
	Constraints *string `json:"Constraints,omitempty" xml:"Constraints,omitempty"`
	// The user who created the parameter.
	//
	// example:
	//
	// root(130900000)
	CreatedBy *string `json:"CreatedBy,omitempty" xml:"CreatedBy,omitempty"`
	// The time when the parameter was created.
	//
	// example:
	//
	// 2020-09-01T09:30:36Z
	CreatedDate *string `json:"CreatedDate,omitempty" xml:"CreatedDate,omitempty"`
	// The description of the parameter.
	//
	// example:
	//
	// SecretParameter
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the parameter.
	//
	// example:
	//
	// p-0b0fff9919c946xxxxxx
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The ID of customer master key (CMK) of Key Management Service (KMS) that is used for encryption.
	//
	// example:
	//
	// 80e9409f-78fa-42ab-84bd-83f40c******
	KeyId *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	// The name of the parameter.
	//
	// example:
	//
	// MyParameter
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The version number of the parameter.
	//
	// example:
	//
	// 2
	ParameterVersion *int32 `json:"ParameterVersion,omitempty" xml:"ParameterVersion,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfmxsn4m4******
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The share type of the parameter.
	//
	// example:
	//
	// Private
	ShareType *string `json:"ShareType,omitempty" xml:"ShareType,omitempty"`
	// The tags of the parameter.
	//
	// example:
	//
	// {"k1": "v1", "k2": "v2"}
	Tags *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The type of the parameter.
	//
	// example:
	//
	// Secret
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The user who updated the parameter.
	//
	// example:
	//
	// root(130900000)
	UpdatedBy *string `json:"UpdatedBy,omitempty" xml:"UpdatedBy,omitempty"`
	// The time when the parameter was updated.
	//
	// example:
	//
	// 2020-09-01T09:33:11Z
	UpdatedDate *string `json:"UpdatedDate,omitempty" xml:"UpdatedDate,omitempty"`
}

func (s UpdateSecretParameterResponseBodyParameter) String() string {
	return tea.Prettify(s)
}

func (s UpdateSecretParameterResponseBodyParameter) GoString() string {
	return s.String()
}

func (s *UpdateSecretParameterResponseBodyParameter) SetConstraints(v string) *UpdateSecretParameterResponseBodyParameter {
	s.Constraints = &v
	return s
}

func (s *UpdateSecretParameterResponseBodyParameter) SetCreatedBy(v string) *UpdateSecretParameterResponseBodyParameter {
	s.CreatedBy = &v
	return s
}

func (s *UpdateSecretParameterResponseBodyParameter) SetCreatedDate(v string) *UpdateSecretParameterResponseBodyParameter {
	s.CreatedDate = &v
	return s
}

func (s *UpdateSecretParameterResponseBodyParameter) SetDescription(v string) *UpdateSecretParameterResponseBodyParameter {
	s.Description = &v
	return s
}

func (s *UpdateSecretParameterResponseBodyParameter) SetId(v string) *UpdateSecretParameterResponseBodyParameter {
	s.Id = &v
	return s
}

func (s *UpdateSecretParameterResponseBodyParameter) SetKeyId(v string) *UpdateSecretParameterResponseBodyParameter {
	s.KeyId = &v
	return s
}

func (s *UpdateSecretParameterResponseBodyParameter) SetName(v string) *UpdateSecretParameterResponseBodyParameter {
	s.Name = &v
	return s
}

func (s *UpdateSecretParameterResponseBodyParameter) SetParameterVersion(v int32) *UpdateSecretParameterResponseBodyParameter {
	s.ParameterVersion = &v
	return s
}

func (s *UpdateSecretParameterResponseBodyParameter) SetResourceGroupId(v string) *UpdateSecretParameterResponseBodyParameter {
	s.ResourceGroupId = &v
	return s
}

func (s *UpdateSecretParameterResponseBodyParameter) SetShareType(v string) *UpdateSecretParameterResponseBodyParameter {
	s.ShareType = &v
	return s
}

func (s *UpdateSecretParameterResponseBodyParameter) SetTags(v string) *UpdateSecretParameterResponseBodyParameter {
	s.Tags = &v
	return s
}

func (s *UpdateSecretParameterResponseBodyParameter) SetType(v string) *UpdateSecretParameterResponseBodyParameter {
	s.Type = &v
	return s
}

func (s *UpdateSecretParameterResponseBodyParameter) SetUpdatedBy(v string) *UpdateSecretParameterResponseBodyParameter {
	s.UpdatedBy = &v
	return s
}

func (s *UpdateSecretParameterResponseBodyParameter) SetUpdatedDate(v string) *UpdateSecretParameterResponseBodyParameter {
	s.UpdatedDate = &v
	return s
}

type UpdateSecretParameterResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateSecretParameterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateSecretParameterResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateSecretParameterResponse) GoString() string {
	return s.String()
}

func (s *UpdateSecretParameterResponse) SetHeaders(v map[string]*string) *UpdateSecretParameterResponse {
	s.Headers = v
	return s
}

func (s *UpdateSecretParameterResponse) SetStatusCode(v int32) *UpdateSecretParameterResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateSecretParameterResponse) SetBody(v *UpdateSecretParameterResponseBody) *UpdateSecretParameterResponse {
	s.Body = v
	return s
}

type UpdateStateConfigurationRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// example:
	//
	// DASKJJLKADS-AHKLJHJSAKL-AJK
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The configuration mode. Valid values: ApplyOnce: The configuration is applied only once. After a configuration is updated, the new configuration is applied. ApplyAndMonitor: The configuration is applied only once. After the configuration is applied, the system only checks whether the configuration is migrated in the future. ApplyAndAutoCorrect: The configuration is always applied.
	//
	// example:
	//
	// ApplyOnce
	ConfigureMode *string `json:"ConfigureMode,omitempty" xml:"ConfigureMode,omitempty"`
	// The description.
	//
	// example:
	//
	// ACS-ECS-InventoryDataCollection
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The parameters.
	//
	// example:
	//
	// { "policy": { "ACS:Application": { "Collection": "Enabled" }, "ACS:Network": { "Collection": "Enabled" } } }
	Parameters map[string]interface{} `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfmxsn4m4******
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The schedule expression.
	//
	// example:
	//
	// 1 hour
	ScheduleExpression *string `json:"ScheduleExpression,omitempty" xml:"ScheduleExpression,omitempty"`
	// The schedule type.
	//
	// example:
	//
	// rate
	ScheduleType *string `json:"ScheduleType,omitempty" xml:"ScheduleType,omitempty"`
	// The ID of the desired-state configuration.
	//
	// This parameter is required.
	//
	// example:
	//
	// sc-asfgdhj12345
	StateConfigurationId *string `json:"StateConfigurationId,omitempty" xml:"StateConfigurationId,omitempty"`
	// The tags to be added to the configuration.
	//
	// example:
	//
	// {"Key": "oos", "Value": "sc"}
	Tags map[string]interface{} `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The resources to be queried.
	//
	// example:
	//
	// { "ResourceType": "ALIYUN::ECS::Instance", "Filters": [ { "Type": "All", "RegionId": "cn-hangzhou", "Parameters": { "RegionId": "cn-hangzhou", "Status": "Running" } } ] }
	Targets *string `json:"Targets,omitempty" xml:"Targets,omitempty"`
}

func (s UpdateStateConfigurationRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateStateConfigurationRequest) GoString() string {
	return s.String()
}

func (s *UpdateStateConfigurationRequest) SetClientToken(v string) *UpdateStateConfigurationRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateStateConfigurationRequest) SetConfigureMode(v string) *UpdateStateConfigurationRequest {
	s.ConfigureMode = &v
	return s
}

func (s *UpdateStateConfigurationRequest) SetDescription(v string) *UpdateStateConfigurationRequest {
	s.Description = &v
	return s
}

func (s *UpdateStateConfigurationRequest) SetParameters(v map[string]interface{}) *UpdateStateConfigurationRequest {
	s.Parameters = v
	return s
}

func (s *UpdateStateConfigurationRequest) SetRegionId(v string) *UpdateStateConfigurationRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateStateConfigurationRequest) SetResourceGroupId(v string) *UpdateStateConfigurationRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *UpdateStateConfigurationRequest) SetScheduleExpression(v string) *UpdateStateConfigurationRequest {
	s.ScheduleExpression = &v
	return s
}

func (s *UpdateStateConfigurationRequest) SetScheduleType(v string) *UpdateStateConfigurationRequest {
	s.ScheduleType = &v
	return s
}

func (s *UpdateStateConfigurationRequest) SetStateConfigurationId(v string) *UpdateStateConfigurationRequest {
	s.StateConfigurationId = &v
	return s
}

func (s *UpdateStateConfigurationRequest) SetTags(v map[string]interface{}) *UpdateStateConfigurationRequest {
	s.Tags = v
	return s
}

func (s *UpdateStateConfigurationRequest) SetTargets(v string) *UpdateStateConfigurationRequest {
	s.Targets = &v
	return s
}

type UpdateStateConfigurationShrinkRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// example:
	//
	// DASKJJLKADS-AHKLJHJSAKL-AJK
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The configuration mode. Valid values: ApplyOnce: The configuration is applied only once. After a configuration is updated, the new configuration is applied. ApplyAndMonitor: The configuration is applied only once. After the configuration is applied, the system only checks whether the configuration is migrated in the future. ApplyAndAutoCorrect: The configuration is always applied.
	//
	// example:
	//
	// ApplyOnce
	ConfigureMode *string `json:"ConfigureMode,omitempty" xml:"ConfigureMode,omitempty"`
	// The description.
	//
	// example:
	//
	// ACS-ECS-InventoryDataCollection
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The parameters.
	//
	// example:
	//
	// { "policy": { "ACS:Application": { "Collection": "Enabled" }, "ACS:Network": { "Collection": "Enabled" } } }
	ParametersShrink *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfmxsn4m4******
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The schedule expression.
	//
	// example:
	//
	// 1 hour
	ScheduleExpression *string `json:"ScheduleExpression,omitempty" xml:"ScheduleExpression,omitempty"`
	// The schedule type.
	//
	// example:
	//
	// rate
	ScheduleType *string `json:"ScheduleType,omitempty" xml:"ScheduleType,omitempty"`
	// The ID of the desired-state configuration.
	//
	// This parameter is required.
	//
	// example:
	//
	// sc-asfgdhj12345
	StateConfigurationId *string `json:"StateConfigurationId,omitempty" xml:"StateConfigurationId,omitempty"`
	// The tags to be added to the configuration.
	//
	// example:
	//
	// {"Key": "oos", "Value": "sc"}
	TagsShrink *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The resources to be queried.
	//
	// example:
	//
	// { "ResourceType": "ALIYUN::ECS::Instance", "Filters": [ { "Type": "All", "RegionId": "cn-hangzhou", "Parameters": { "RegionId": "cn-hangzhou", "Status": "Running" } } ] }
	Targets *string `json:"Targets,omitempty" xml:"Targets,omitempty"`
}

func (s UpdateStateConfigurationShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateStateConfigurationShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateStateConfigurationShrinkRequest) SetClientToken(v string) *UpdateStateConfigurationShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateStateConfigurationShrinkRequest) SetConfigureMode(v string) *UpdateStateConfigurationShrinkRequest {
	s.ConfigureMode = &v
	return s
}

func (s *UpdateStateConfigurationShrinkRequest) SetDescription(v string) *UpdateStateConfigurationShrinkRequest {
	s.Description = &v
	return s
}

func (s *UpdateStateConfigurationShrinkRequest) SetParametersShrink(v string) *UpdateStateConfigurationShrinkRequest {
	s.ParametersShrink = &v
	return s
}

func (s *UpdateStateConfigurationShrinkRequest) SetRegionId(v string) *UpdateStateConfigurationShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateStateConfigurationShrinkRequest) SetResourceGroupId(v string) *UpdateStateConfigurationShrinkRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *UpdateStateConfigurationShrinkRequest) SetScheduleExpression(v string) *UpdateStateConfigurationShrinkRequest {
	s.ScheduleExpression = &v
	return s
}

func (s *UpdateStateConfigurationShrinkRequest) SetScheduleType(v string) *UpdateStateConfigurationShrinkRequest {
	s.ScheduleType = &v
	return s
}

func (s *UpdateStateConfigurationShrinkRequest) SetStateConfigurationId(v string) *UpdateStateConfigurationShrinkRequest {
	s.StateConfigurationId = &v
	return s
}

func (s *UpdateStateConfigurationShrinkRequest) SetTagsShrink(v string) *UpdateStateConfigurationShrinkRequest {
	s.TagsShrink = &v
	return s
}

func (s *UpdateStateConfigurationShrinkRequest) SetTargets(v string) *UpdateStateConfigurationShrinkRequest {
	s.Targets = &v
	return s
}

type UpdateStateConfigurationResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 1306108F-610C-40FD-AAD5-DA13E8B00BE9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The details of the configuration.
	StateConfiguration []*UpdateStateConfigurationResponseBodyStateConfiguration `json:"StateConfiguration,omitempty" xml:"StateConfiguration,omitempty" type:"Repeated"`
}

func (s UpdateStateConfigurationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateStateConfigurationResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateStateConfigurationResponseBody) SetRequestId(v string) *UpdateStateConfigurationResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateStateConfigurationResponseBody) SetStateConfiguration(v []*UpdateStateConfigurationResponseBodyStateConfiguration) *UpdateStateConfigurationResponseBody {
	s.StateConfiguration = v
	return s
}

type UpdateStateConfigurationResponseBodyStateConfiguration struct {
	// The configuration mode. Valid values:
	//
	// example:
	//
	// ApplyAndAutoCorrect
	ConfigureMode *string `json:"ConfigureMode,omitempty" xml:"ConfigureMode,omitempty"`
	// The time when the configuration was created.
	//
	// example:
	//
	// 2021-03-22T03:13:32Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the desired-state configuration.
	//
	// example:
	//
	// collect inventory data
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The parameters.
	//
	// example:
	//
	// {"policy": {"ACS:Network": {"Collection": "Enabled"}, "ACS:Application": {"Collection": "Enabled"}}}
	Parameters *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfmxsn4m4******
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The CRON expression.
	//
	// example:
	//
	// 1 hour
	ScheduleExpression *string `json:"ScheduleExpression,omitempty" xml:"ScheduleExpression,omitempty"`
	// The schedule type.
	//
	// example:
	//
	// rate
	ScheduleType *string `json:"ScheduleType,omitempty" xml:"ScheduleType,omitempty"`
	// The ID of the desired-state configuration.
	//
	// example:
	//
	// StateConfigurationId
	StateConfigurationId *string `json:"StateConfigurationId,omitempty" xml:"StateConfigurationId,omitempty"`
	// The tags added to the configuration.
	//
	// example:
	//
	// {"Key": "oos", "Value": "inventory"}
	Tags map[string]interface{} `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The queried resources.
	//
	// example:
	//
	// { "ResourceType": "ALIYUN::ECS::Instance", "Filters": [ { "Type": "All", "RegionId": "cn-hangzhou", "Parameters": { "RegionId": "cn-hangzhou", "Status": "Running" } } ] }
	Targets *string `json:"Targets,omitempty" xml:"Targets,omitempty"`
	// The template ID.
	//
	// example:
	//
	// t-1234asadf
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The name of the template.
	//
	// example:
	//
	// ACS-ECS-InventoryDataCollection
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// The name of the template version.
	//
	// example:
	//
	// v1
	TemplateVersion *string `json:"TemplateVersion,omitempty" xml:"TemplateVersion,omitempty"`
	// The time when the configuration was updated.
	//
	// example:
	//
	// 2021-03-22T03:13:32Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s UpdateStateConfigurationResponseBodyStateConfiguration) String() string {
	return tea.Prettify(s)
}

func (s UpdateStateConfigurationResponseBodyStateConfiguration) GoString() string {
	return s.String()
}

func (s *UpdateStateConfigurationResponseBodyStateConfiguration) SetConfigureMode(v string) *UpdateStateConfigurationResponseBodyStateConfiguration {
	s.ConfigureMode = &v
	return s
}

func (s *UpdateStateConfigurationResponseBodyStateConfiguration) SetCreateTime(v string) *UpdateStateConfigurationResponseBodyStateConfiguration {
	s.CreateTime = &v
	return s
}

func (s *UpdateStateConfigurationResponseBodyStateConfiguration) SetDescription(v string) *UpdateStateConfigurationResponseBodyStateConfiguration {
	s.Description = &v
	return s
}

func (s *UpdateStateConfigurationResponseBodyStateConfiguration) SetParameters(v string) *UpdateStateConfigurationResponseBodyStateConfiguration {
	s.Parameters = &v
	return s
}

func (s *UpdateStateConfigurationResponseBodyStateConfiguration) SetResourceGroupId(v string) *UpdateStateConfigurationResponseBodyStateConfiguration {
	s.ResourceGroupId = &v
	return s
}

func (s *UpdateStateConfigurationResponseBodyStateConfiguration) SetScheduleExpression(v string) *UpdateStateConfigurationResponseBodyStateConfiguration {
	s.ScheduleExpression = &v
	return s
}

func (s *UpdateStateConfigurationResponseBodyStateConfiguration) SetScheduleType(v string) *UpdateStateConfigurationResponseBodyStateConfiguration {
	s.ScheduleType = &v
	return s
}

func (s *UpdateStateConfigurationResponseBodyStateConfiguration) SetStateConfigurationId(v string) *UpdateStateConfigurationResponseBodyStateConfiguration {
	s.StateConfigurationId = &v
	return s
}

func (s *UpdateStateConfigurationResponseBodyStateConfiguration) SetTags(v map[string]interface{}) *UpdateStateConfigurationResponseBodyStateConfiguration {
	s.Tags = v
	return s
}

func (s *UpdateStateConfigurationResponseBodyStateConfiguration) SetTargets(v string) *UpdateStateConfigurationResponseBodyStateConfiguration {
	s.Targets = &v
	return s
}

func (s *UpdateStateConfigurationResponseBodyStateConfiguration) SetTemplateId(v string) *UpdateStateConfigurationResponseBodyStateConfiguration {
	s.TemplateId = &v
	return s
}

func (s *UpdateStateConfigurationResponseBodyStateConfiguration) SetTemplateName(v string) *UpdateStateConfigurationResponseBodyStateConfiguration {
	s.TemplateName = &v
	return s
}

func (s *UpdateStateConfigurationResponseBodyStateConfiguration) SetTemplateVersion(v string) *UpdateStateConfigurationResponseBodyStateConfiguration {
	s.TemplateVersion = &v
	return s
}

func (s *UpdateStateConfigurationResponseBodyStateConfiguration) SetUpdateTime(v string) *UpdateStateConfigurationResponseBodyStateConfiguration {
	s.UpdateTime = &v
	return s
}

type UpdateStateConfigurationResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateStateConfigurationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateStateConfigurationResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateStateConfigurationResponse) GoString() string {
	return s.String()
}

func (s *UpdateStateConfigurationResponse) SetHeaders(v map[string]*string) *UpdateStateConfigurationResponse {
	s.Headers = v
	return s
}

func (s *UpdateStateConfigurationResponse) SetStatusCode(v int32) *UpdateStateConfigurationResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateStateConfigurationResponse) SetBody(v *UpdateStateConfigurationResponseBody) *UpdateStateConfigurationResponse {
	s.Body = v
	return s
}

type UpdateTemplateRequest struct {
	// The content of the template. The content must be in the JSON or YAML format, and its maximum size is 64 KB.
	//
	// This parameter is required.
	//
	// example:
	//
	// { "FormatVersion": "OOS-2019-06-01", "Description": { "en": "Bulky starts the ECS instances", "name-en": "Bulky Start Instances", }, "Parameters": { "regionId": { "Type": "String", "Label": { "en": "RegionId", }, "AssociationProperty": "RegionId", "Default": "{{ ACS::RegionId }}" }, "targets": { "Type": "Json", "Label": { "en": "TargetInstance", }, "AssociationProperty": "Targets", "AssociationPropertyMetadata": { "ResourceType": "ALIYUN::ECS::Instance", "RegionId": "regionId" } }, "rateControl": { "Label": { "en": "RateControl", }, "Type": "Json", "AssociationProperty": "RateControl", "Default": { "Mode": "Concurrency", "MaxErrors": 0, "Concurrency": 10 } }, "OOSAssumeRole": { "Label": { "en": "OOSAssumeRole", }, "Type": "String", "Default": "OOSServiceRole" } }, "RamRole": "{{ OOSAssumeRole }}", "Tasks": [ { "Name": "getInstance", "Description": { "en": "Views the ECS instances", }, "Action": "ACS::SelectTargets", "Properties": { "ResourceType": "ALIYUN::ECS::Instance", "RegionId": "{{ regionId }}", "Filters": [ "{{ targets }}" ] }, "Outputs": { "instanceIds": { "Type": "List", "ValueSelector": "Instances.Instance[].InstanceId" } } }, { "Name": "startInstance", "Action": "ACS::ECS::StartInstance", "Description": { "en": "Starts the ECS instances", }, "Properties": { "regionId": "{{ regionId }}", "instanceId": "{{ ACS::TaskLoopItem }}" }, "Loop": { "RateControl": "{{ rateControl }}", "Items": "{{ getInstance.instanceIds }}" } } ], "Outputs": { "instanceIds": { "Type": "List", "Value": "{{ getInstance.instanceIds }}" } } }
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The ID of the region.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfmxsn4m4******
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The tag keys and values. The number of key-value pairs ranges from 1 to 20.
	//
	// example:
	//
	// {"k1":"k2","k2":"v2"}
	Tags map[string]interface{} `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The name of the template. The name can be up to 200 characters in length and can contain letters, digits, hyphens (-), and underscores (_). The name cannot start with ALIYUN, ACS, ALIBABA, or ALICLOUD.
	//
	// This parameter is required.
	//
	// example:
	//
	// MyTemplate
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// The name of the template version.
	//
	// example:
	//
	// v2
	VersionName *string `json:"VersionName,omitempty" xml:"VersionName,omitempty"`
}

func (s UpdateTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateTemplateRequest) GoString() string {
	return s.String()
}

func (s *UpdateTemplateRequest) SetContent(v string) *UpdateTemplateRequest {
	s.Content = &v
	return s
}

func (s *UpdateTemplateRequest) SetRegionId(v string) *UpdateTemplateRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateTemplateRequest) SetResourceGroupId(v string) *UpdateTemplateRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *UpdateTemplateRequest) SetTags(v map[string]interface{}) *UpdateTemplateRequest {
	s.Tags = v
	return s
}

func (s *UpdateTemplateRequest) SetTemplateName(v string) *UpdateTemplateRequest {
	s.TemplateName = &v
	return s
}

func (s *UpdateTemplateRequest) SetVersionName(v string) *UpdateTemplateRequest {
	s.VersionName = &v
	return s
}

type UpdateTemplateShrinkRequest struct {
	// The content of the template. The content must be in the JSON or YAML format, and its maximum size is 64 KB.
	//
	// This parameter is required.
	//
	// example:
	//
	// { "FormatVersion": "OOS-2019-06-01", "Description": { "en": "Bulky starts the ECS instances", "name-en": "Bulky Start Instances", }, "Parameters": { "regionId": { "Type": "String", "Label": { "en": "RegionId", }, "AssociationProperty": "RegionId", "Default": "{{ ACS::RegionId }}" }, "targets": { "Type": "Json", "Label": { "en": "TargetInstance", }, "AssociationProperty": "Targets", "AssociationPropertyMetadata": { "ResourceType": "ALIYUN::ECS::Instance", "RegionId": "regionId" } }, "rateControl": { "Label": { "en": "RateControl", }, "Type": "Json", "AssociationProperty": "RateControl", "Default": { "Mode": "Concurrency", "MaxErrors": 0, "Concurrency": 10 } }, "OOSAssumeRole": { "Label": { "en": "OOSAssumeRole", }, "Type": "String", "Default": "OOSServiceRole" } }, "RamRole": "{{ OOSAssumeRole }}", "Tasks": [ { "Name": "getInstance", "Description": { "en": "Views the ECS instances", }, "Action": "ACS::SelectTargets", "Properties": { "ResourceType": "ALIYUN::ECS::Instance", "RegionId": "{{ regionId }}", "Filters": [ "{{ targets }}" ] }, "Outputs": { "instanceIds": { "Type": "List", "ValueSelector": "Instances.Instance[].InstanceId" } } }, { "Name": "startInstance", "Action": "ACS::ECS::StartInstance", "Description": { "en": "Starts the ECS instances", }, "Properties": { "regionId": "{{ regionId }}", "instanceId": "{{ ACS::TaskLoopItem }}" }, "Loop": { "RateControl": "{{ rateControl }}", "Items": "{{ getInstance.instanceIds }}" } } ], "Outputs": { "instanceIds": { "Type": "List", "Value": "{{ getInstance.instanceIds }}" } } }
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The ID of the region.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfmxsn4m4******
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The tag keys and values. The number of key-value pairs ranges from 1 to 20.
	//
	// example:
	//
	// {"k1":"k2","k2":"v2"}
	TagsShrink *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The name of the template. The name can be up to 200 characters in length and can contain letters, digits, hyphens (-), and underscores (_). The name cannot start with ALIYUN, ACS, ALIBABA, or ALICLOUD.
	//
	// This parameter is required.
	//
	// example:
	//
	// MyTemplate
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// The name of the template version.
	//
	// example:
	//
	// v2
	VersionName *string `json:"VersionName,omitempty" xml:"VersionName,omitempty"`
}

func (s UpdateTemplateShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateTemplateShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateTemplateShrinkRequest) SetContent(v string) *UpdateTemplateShrinkRequest {
	s.Content = &v
	return s
}

func (s *UpdateTemplateShrinkRequest) SetRegionId(v string) *UpdateTemplateShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateTemplateShrinkRequest) SetResourceGroupId(v string) *UpdateTemplateShrinkRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *UpdateTemplateShrinkRequest) SetTagsShrink(v string) *UpdateTemplateShrinkRequest {
	s.TagsShrink = &v
	return s
}

func (s *UpdateTemplateShrinkRequest) SetTemplateName(v string) *UpdateTemplateShrinkRequest {
	s.TemplateName = &v
	return s
}

func (s *UpdateTemplateShrinkRequest) SetVersionName(v string) *UpdateTemplateShrinkRequest {
	s.VersionName = &v
	return s
}

type UpdateTemplateResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 2075899A-585D-4A41-A9B2-28DF4F534F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The metadata of the template.
	Template *UpdateTemplateResponseBodyTemplate `json:"Template,omitempty" xml:"Template,omitempty" type:"Struct"`
}

func (s UpdateTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateTemplateResponseBody) SetRequestId(v string) *UpdateTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateTemplateResponseBody) SetTemplate(v *UpdateTemplateResponseBodyTemplate) *UpdateTemplateResponseBody {
	s.Template = v
	return s
}

type UpdateTemplateResponseBodyTemplate struct {
	// The user who created the template.
	//
	// example:
	//
	// root(130920000)
	CreatedBy *string `json:"CreatedBy,omitempty" xml:"CreatedBy,omitempty"`
	// The time when the template was created.
	//
	// example:
	//
	// 2019-05-16T10:26:14Z
	CreatedDate *string `json:"CreatedDate,omitempty" xml:"CreatedDate,omitempty"`
	// The description of the template.
	//
	// example:
	//
	// Describe instances of given status
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Indicates whether the template is configured with a trigger.
	//
	// example:
	//
	// true
	HasTrigger *bool `json:"HasTrigger,omitempty" xml:"HasTrigger,omitempty"`
	// The SHA-256 value of the template content.
	//
	// example:
	//
	// 4bc7d7a21b3e003434b9c223f6e6d2578b5ebfeb5be28c1fcf8a8a1b11907bb4
	Hash *string `json:"Hash,omitempty" xml:"Hash,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfmxsn4m4******
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The share type of the template. The share type of a user-created template is **Private**.
	//
	// example:
	//
	// Private
	ShareType *string `json:"ShareType,omitempty" xml:"ShareType,omitempty"`
	// The tag keys and values. The number of key-value pairs ranges from 1 to 20.
	//
	// example:
	//
	// {"k1":"k2","k2":"v2"}
	Tags map[string]interface{} `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The format of the template. The system automatically determines whether the format is JSON or YAML.
	//
	// example:
	//
	// JSON
	TemplateFormat *string `json:"TemplateFormat,omitempty" xml:"TemplateFormat,omitempty"`
	// The ID of the template.
	//
	// example:
	//
	// t-94753deed38
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The name of the template.
	//
	// example:
	//
	// MyTemplate
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// The version of the template. The name of the version consists of the letter v and a number. The number starts from 1.
	//
	// example:
	//
	// v2
	TemplateVersion *string `json:"TemplateVersion,omitempty" xml:"TemplateVersion,omitempty"`
	// The user who last modified the information about the template.
	//
	// example:
	//
	// root(1309000)
	UpdatedBy *string `json:"UpdatedBy,omitempty" xml:"UpdatedBy,omitempty"`
	// The time when the information about the template was last modified.
	//
	// example:
	//
	// 2019-05-16T10:26:14Z
	UpdatedDate *string `json:"UpdatedDate,omitempty" xml:"UpdatedDate,omitempty"`
}

func (s UpdateTemplateResponseBodyTemplate) String() string {
	return tea.Prettify(s)
}

func (s UpdateTemplateResponseBodyTemplate) GoString() string {
	return s.String()
}

func (s *UpdateTemplateResponseBodyTemplate) SetCreatedBy(v string) *UpdateTemplateResponseBodyTemplate {
	s.CreatedBy = &v
	return s
}

func (s *UpdateTemplateResponseBodyTemplate) SetCreatedDate(v string) *UpdateTemplateResponseBodyTemplate {
	s.CreatedDate = &v
	return s
}

func (s *UpdateTemplateResponseBodyTemplate) SetDescription(v string) *UpdateTemplateResponseBodyTemplate {
	s.Description = &v
	return s
}

func (s *UpdateTemplateResponseBodyTemplate) SetHasTrigger(v bool) *UpdateTemplateResponseBodyTemplate {
	s.HasTrigger = &v
	return s
}

func (s *UpdateTemplateResponseBodyTemplate) SetHash(v string) *UpdateTemplateResponseBodyTemplate {
	s.Hash = &v
	return s
}

func (s *UpdateTemplateResponseBodyTemplate) SetResourceGroupId(v string) *UpdateTemplateResponseBodyTemplate {
	s.ResourceGroupId = &v
	return s
}

func (s *UpdateTemplateResponseBodyTemplate) SetShareType(v string) *UpdateTemplateResponseBodyTemplate {
	s.ShareType = &v
	return s
}

func (s *UpdateTemplateResponseBodyTemplate) SetTags(v map[string]interface{}) *UpdateTemplateResponseBodyTemplate {
	s.Tags = v
	return s
}

func (s *UpdateTemplateResponseBodyTemplate) SetTemplateFormat(v string) *UpdateTemplateResponseBodyTemplate {
	s.TemplateFormat = &v
	return s
}

func (s *UpdateTemplateResponseBodyTemplate) SetTemplateId(v string) *UpdateTemplateResponseBodyTemplate {
	s.TemplateId = &v
	return s
}

func (s *UpdateTemplateResponseBodyTemplate) SetTemplateName(v string) *UpdateTemplateResponseBodyTemplate {
	s.TemplateName = &v
	return s
}

func (s *UpdateTemplateResponseBodyTemplate) SetTemplateVersion(v string) *UpdateTemplateResponseBodyTemplate {
	s.TemplateVersion = &v
	return s
}

func (s *UpdateTemplateResponseBodyTemplate) SetUpdatedBy(v string) *UpdateTemplateResponseBodyTemplate {
	s.UpdatedBy = &v
	return s
}

func (s *UpdateTemplateResponseBodyTemplate) SetUpdatedDate(v string) *UpdateTemplateResponseBodyTemplate {
	s.UpdatedDate = &v
	return s
}

type UpdateTemplateResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateTemplateResponse) GoString() string {
	return s.String()
}

func (s *UpdateTemplateResponse) SetHeaders(v map[string]*string) *UpdateTemplateResponse {
	s.Headers = v
	return s
}

func (s *UpdateTemplateResponse) SetStatusCode(v int32) *UpdateTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateTemplateResponse) SetBody(v *UpdateTemplateResponseBody) *UpdateTemplateResponse {
	s.Body = v
	return s
}

type ValidateTemplateContentRequest struct {
	// The content of the template.
	//
	// example:
	//
	// {"FormatVersion": "OOS-2019-06-01", "Description": "Describe instances of given status", "Parameters": {"Status": {"Type": "String", "Description": "(Required) The status of the Ecs instance."}}, "Tasks": [{"Properties": {"Parameters": {"Status": "{{ Status }}"}, "API": "DescribeInstances", "Service": "Ecs"}, "Name": "foo", "Action": "ACS::ExecuteApi"}]}
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The ID of the region.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The URL that is used to store the content of the Operation Orchestration Service (OOS) template in the Alibaba Cloud Object Storage Service (OSS). Only the public-read URL is supported. You can use this parameter to specify the tasks that you want to run. This way, you do not need to create a template before you start an execution. If you select an existing template, you do not need to specify this parameter.
	//
	// example:
	//
	// http:/oos-template.cn-hangzhou.oss.aliyun-inc.com/oos-template.json
	TemplateURL *string `json:"TemplateURL,omitempty" xml:"TemplateURL,omitempty"`
}

func (s ValidateTemplateContentRequest) String() string {
	return tea.Prettify(s)
}

func (s ValidateTemplateContentRequest) GoString() string {
	return s.String()
}

func (s *ValidateTemplateContentRequest) SetContent(v string) *ValidateTemplateContentRequest {
	s.Content = &v
	return s
}

func (s *ValidateTemplateContentRequest) SetRegionId(v string) *ValidateTemplateContentRequest {
	s.RegionId = &v
	return s
}

func (s *ValidateTemplateContentRequest) SetTemplateURL(v string) *ValidateTemplateContentRequest {
	s.TemplateURL = &v
	return s
}

type ValidateTemplateContentResponseBody struct {
	// The outputs of the template.
	//
	// example:
	//
	// {}
	Outputs *string `json:"Outputs,omitempty" xml:"Outputs,omitempty"`
	// The parameters of the template.
	//
	// example:
	//
	// { "Status": { "Description": "(Required) The status of the Ecs instance.", "Type": "String" } }
	Parameters *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// The RAM role.
	//
	// example:
	//
	// OOSServiceRole
	RamRole *string `json:"RamRole,omitempty" xml:"RamRole,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// D5EE9591-1F2D-573E-8751-7F08BBB388D4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The task defined in the template.
	Tasks []*ValidateTemplateContentResponseBodyTasks `json:"Tasks,omitempty" xml:"Tasks,omitempty" type:"Repeated"`
}

func (s ValidateTemplateContentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ValidateTemplateContentResponseBody) GoString() string {
	return s.String()
}

func (s *ValidateTemplateContentResponseBody) SetOutputs(v string) *ValidateTemplateContentResponseBody {
	s.Outputs = &v
	return s
}

func (s *ValidateTemplateContentResponseBody) SetParameters(v string) *ValidateTemplateContentResponseBody {
	s.Parameters = &v
	return s
}

func (s *ValidateTemplateContentResponseBody) SetRamRole(v string) *ValidateTemplateContentResponseBody {
	s.RamRole = &v
	return s
}

func (s *ValidateTemplateContentResponseBody) SetRequestId(v string) *ValidateTemplateContentResponseBody {
	s.RequestId = &v
	return s
}

func (s *ValidateTemplateContentResponseBody) SetTasks(v []*ValidateTemplateContentResponseBodyTasks) *ValidateTemplateContentResponseBody {
	s.Tasks = v
	return s
}

type ValidateTemplateContentResponseBodyTasks struct {
	// The description of the task.
	//
	// example:
	//
	// (Required) The status of the Ecs instance.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the task.
	//
	// example:
	//
	// foo
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The outputs of the task.
	//
	// example:
	//
	// .instanceId
	Outputs *string `json:"Outputs,omitempty" xml:"Outputs,omitempty"`
	// The properties of the task.
	//
	// example:
	//
	// {"API": "DescribeInstances","Parameters": {"Status": "{{ Status }}"},"Service": "Ecs"}
	Properties *string `json:"Properties,omitempty" xml:"Properties,omitempty"`
	// The type of the task.
	//
	// example:
	//
	// ACS::ExecuteAPI
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ValidateTemplateContentResponseBodyTasks) String() string {
	return tea.Prettify(s)
}

func (s ValidateTemplateContentResponseBodyTasks) GoString() string {
	return s.String()
}

func (s *ValidateTemplateContentResponseBodyTasks) SetDescription(v string) *ValidateTemplateContentResponseBodyTasks {
	s.Description = &v
	return s
}

func (s *ValidateTemplateContentResponseBodyTasks) SetName(v string) *ValidateTemplateContentResponseBodyTasks {
	s.Name = &v
	return s
}

func (s *ValidateTemplateContentResponseBodyTasks) SetOutputs(v string) *ValidateTemplateContentResponseBodyTasks {
	s.Outputs = &v
	return s
}

func (s *ValidateTemplateContentResponseBodyTasks) SetProperties(v string) *ValidateTemplateContentResponseBodyTasks {
	s.Properties = &v
	return s
}

func (s *ValidateTemplateContentResponseBodyTasks) SetType(v string) *ValidateTemplateContentResponseBodyTasks {
	s.Type = &v
	return s
}

type ValidateTemplateContentResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ValidateTemplateContentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ValidateTemplateContentResponse) String() string {
	return tea.Prettify(s)
}

func (s ValidateTemplateContentResponse) GoString() string {
	return s.String()
}

func (s *ValidateTemplateContentResponse) SetHeaders(v map[string]*string) *ValidateTemplateContentResponse {
	s.Headers = v
	return s
}

func (s *ValidateTemplateContentResponse) SetStatusCode(v int32) *ValidateTemplateContentResponse {
	s.StatusCode = &v
	return s
}

func (s *ValidateTemplateContentResponse) SetBody(v *ValidateTemplateContentResponseBody) *ValidateTemplateContentResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("oos"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// Cancels an execution.
//
// @param request - CancelExecutionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelExecutionResponse
func (client *Client) CancelExecutionWithOptions(request *CancelExecutionRequest, runtime *util.RuntimeOptions) (_result *CancelExecutionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ExecutionId)) {
		query["ExecutionId"] = request.ExecutionId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CancelExecution"),
		Version:     tea.String("2019-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CancelExecutionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Cancels an execution.
//
// @param request - CancelExecutionRequest
//
// @return CancelExecutionResponse
func (client *Client) CancelExecution(request *CancelExecutionRequest) (_result *CancelExecutionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CancelExecutionResponse{}
	_body, _err := client.CancelExecutionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the resource group to which a cloud resource belongs.
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
		Version:     tea.String("2019-06-01"),
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
// Modifies the resource group to which a cloud resource belongs.
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
// Continues deploying an application group when an error occurs for calling the DeployApplicationGroup operation. You can call this operation only for the applications which reside in the China (Hangzhou) region. Use an endpoint of the China (Hangzhou) region.
//
// @param request - ContinueDeployApplicationGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ContinueDeployApplicationGroupResponse
func (client *Client) ContinueDeployApplicationGroupWithOptions(request *ContinueDeployApplicationGroupRequest, runtime *util.RuntimeOptions) (_result *ContinueDeployApplicationGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApplicationName)) {
		query["ApplicationName"] = request.ApplicationName
	}

	if !tea.BoolValue(util.IsUnset(request.DeployParameters)) {
		query["DeployParameters"] = request.DeployParameters
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ContinueDeployApplicationGroup"),
		Version:     tea.String("2019-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ContinueDeployApplicationGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Continues deploying an application group when an error occurs for calling the DeployApplicationGroup operation. You can call this operation only for the applications which reside in the China (Hangzhou) region. Use an endpoint of the China (Hangzhou) region.
//
// @param request - ContinueDeployApplicationGroupRequest
//
// @return ContinueDeployApplicationGroupResponse
func (client *Client) ContinueDeployApplicationGroup(request *ContinueDeployApplicationGroupRequest) (_result *ContinueDeployApplicationGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ContinueDeployApplicationGroupResponse{}
	_body, _err := client.ContinueDeployApplicationGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an application. You can call this operation only for the applications that reside in the China (Hangzhou) region. Use an endpoint of the China (Hangzhou) region.
//
// @param tmpReq - CreateApplicationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateApplicationResponse
func (client *Client) CreateApplicationWithOptions(tmpReq *CreateApplicationRequest, runtime *util.RuntimeOptions) (_result *CreateApplicationResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateApplicationShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.AlarmConfig)) {
		request.AlarmConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AlarmConfig, tea.String("AlarmConfig"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Tags)) {
		request.TagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tags, tea.String("Tags"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AlarmConfigShrink)) {
		query["AlarmConfig"] = request.AlarmConfigShrink
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
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

	if !tea.BoolValue(util.IsUnset(request.TagsShrink)) {
		query["Tags"] = request.TagsShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateApplication"),
		Version:     tea.String("2019-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateApplicationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an application. You can call this operation only for the applications that reside in the China (Hangzhou) region. Use an endpoint of the China (Hangzhou) region.
//
// @param request - CreateApplicationRequest
//
// @return CreateApplicationResponse
func (client *Client) CreateApplication(request *CreateApplicationRequest) (_result *CreateApplicationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateApplicationResponse{}
	_body, _err := client.CreateApplicationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an application group. You can call this operation only for the application groups that reside in the China (Hangzhou) region. Use an endpoint of the China (Hangzhou) region.
//
// @param request - CreateApplicationGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateApplicationGroupResponse
func (client *Client) CreateApplicationGroupWithOptions(request *CreateApplicationGroupRequest, runtime *util.RuntimeOptions) (_result *CreateApplicationGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApplicationName)) {
		query["ApplicationName"] = request.ApplicationName
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.CmsGroupId)) {
		query["CmsGroupId"] = request.CmsGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.DeployRegionId)) {
		query["DeployRegionId"] = request.DeployRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.ImportTagKey)) {
		query["ImportTagKey"] = request.ImportTagKey
	}

	if !tea.BoolValue(util.IsUnset(request.ImportTagValue)) {
		query["ImportTagValue"] = request.ImportTagValue
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateApplicationGroup"),
		Version:     tea.String("2019-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateApplicationGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an application group. You can call this operation only for the application groups that reside in the China (Hangzhou) region. Use an endpoint of the China (Hangzhou) region.
//
// @param request - CreateApplicationGroupRequest
//
// @return CreateApplicationGroupResponse
func (client *Client) CreateApplicationGroup(request *CreateApplicationGroupRequest) (_result *CreateApplicationGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateApplicationGroupResponse{}
	_body, _err := client.CreateApplicationGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an O\\\\\\\\\\\\&M Item.
//
// @param tmpReq - CreateOpsItemRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateOpsItemResponse
func (client *Client) CreateOpsItemWithOptions(tmpReq *CreateOpsItemRequest, runtime *util.RuntimeOptions) (_result *CreateOpsItemResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateOpsItemShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Tags)) {
		request.TagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tags, tea.String("Tags"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Category)) {
		query["Category"] = request.Category
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DedupString)) {
		query["DedupString"] = request.DedupString
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Priority)) {
		query["Priority"] = request.Priority
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.Resources)) {
		query["Resources"] = request.Resources
	}

	if !tea.BoolValue(util.IsUnset(request.Severity)) {
		query["Severity"] = request.Severity
	}

	if !tea.BoolValue(util.IsUnset(request.Solutions)) {
		query["Solutions"] = request.Solutions
	}

	if !tea.BoolValue(util.IsUnset(request.Source)) {
		query["Source"] = request.Source
	}

	if !tea.BoolValue(util.IsUnset(request.TagsShrink)) {
		query["Tags"] = request.TagsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.Title)) {
		query["Title"] = request.Title
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateOpsItem"),
		Version:     tea.String("2019-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateOpsItemResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an O\\\\\\\\\\\\&M Item.
//
// @param request - CreateOpsItemRequest
//
// @return CreateOpsItemResponse
func (client *Client) CreateOpsItem(request *CreateOpsItemRequest) (_result *CreateOpsItemResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateOpsItemResponse{}
	_body, _err := client.CreateOpsItemWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a common parameter.
//
// @param tmpReq - CreateParameterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateParameterResponse
func (client *Client) CreateParameterWithOptions(tmpReq *CreateParameterRequest, runtime *util.RuntimeOptions) (_result *CreateParameterResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateParameterShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Tags)) {
		request.TagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tags, tea.String("Tags"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Constraints)) {
		query["Constraints"] = request.Constraints
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.TagsShrink)) {
		query["Tags"] = request.TagsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["Type"] = request.Type
	}

	if !tea.BoolValue(util.IsUnset(request.Value)) {
		query["Value"] = request.Value
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateParameter"),
		Version:     tea.String("2019-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateParameterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a common parameter.
//
// @param request - CreateParameterRequest
//
// @return CreateParameterResponse
func (client *Client) CreateParameter(request *CreateParameterRequest) (_result *CreateParameterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateParameterResponse{}
	_body, _err := client.CreateParameterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a patch baseline.
//
// @param tmpReq - CreatePatchBaselineRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatePatchBaselineResponse
func (client *Client) CreatePatchBaselineWithOptions(tmpReq *CreatePatchBaselineRequest, runtime *util.RuntimeOptions) (_result *CreatePatchBaselineResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreatePatchBaselineShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.ApprovedPatches)) {
		request.ApprovedPatchesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ApprovedPatches, tea.String("ApprovedPatches"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.RejectedPatches)) {
		request.RejectedPatchesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RejectedPatches, tea.String("RejectedPatches"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Sources)) {
		request.SourcesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Sources, tea.String("Sources"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Tags)) {
		request.TagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tags, tea.String("Tags"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApprovalRules)) {
		query["ApprovalRules"] = request.ApprovalRules
	}

	if !tea.BoolValue(util.IsUnset(request.ApprovedPatchesShrink)) {
		query["ApprovedPatches"] = request.ApprovedPatchesShrink
	}

	if !tea.BoolValue(util.IsUnset(request.ApprovedPatchesEnableNonSecurity)) {
		query["ApprovedPatchesEnableNonSecurity"] = request.ApprovedPatchesEnableNonSecurity
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.OperationSystem)) {
		query["OperationSystem"] = request.OperationSystem
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RejectedPatchesShrink)) {
		query["RejectedPatches"] = request.RejectedPatchesShrink
	}

	if !tea.BoolValue(util.IsUnset(request.RejectedPatchesAction)) {
		query["RejectedPatchesAction"] = request.RejectedPatchesAction
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.SourcesShrink)) {
		query["Sources"] = request.SourcesShrink
	}

	if !tea.BoolValue(util.IsUnset(request.TagsShrink)) {
		query["Tags"] = request.TagsShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreatePatchBaseline"),
		Version:     tea.String("2019-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreatePatchBaselineResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a patch baseline.
//
// @param request - CreatePatchBaselineRequest
//
// @return CreatePatchBaselineResponse
func (client *Client) CreatePatchBaseline(request *CreatePatchBaselineRequest) (_result *CreatePatchBaselineResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreatePatchBaselineResponse{}
	_body, _err := client.CreatePatchBaselineWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an encryption parameter. Make sure that you have the permissions to call this operation.
//
// @param tmpReq - CreateSecretParameterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSecretParameterResponse
func (client *Client) CreateSecretParameterWithOptions(tmpReq *CreateSecretParameterRequest, runtime *util.RuntimeOptions) (_result *CreateSecretParameterResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateSecretParameterShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Tags)) {
		request.TagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tags, tea.String("Tags"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Constraints)) {
		query["Constraints"] = request.Constraints
	}

	if !tea.BoolValue(util.IsUnset(request.DKMSInstanceId)) {
		query["DKMSInstanceId"] = request.DKMSInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.KeyId)) {
		query["KeyId"] = request.KeyId
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.TagsShrink)) {
		query["Tags"] = request.TagsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["Type"] = request.Type
	}

	if !tea.BoolValue(util.IsUnset(request.Value)) {
		query["Value"] = request.Value
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateSecretParameter"),
		Version:     tea.String("2019-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateSecretParameterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an encryption parameter. Make sure that you have the permissions to call this operation.
//
// @param request - CreateSecretParameterRequest
//
// @return CreateSecretParameterResponse
func (client *Client) CreateSecretParameter(request *CreateSecretParameterRequest) (_result *CreateSecretParameterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateSecretParameterResponse{}
	_body, _err := client.CreateSecretParameterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a desired-state configuration.
//
// @param tmpReq - CreateStateConfigurationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateStateConfigurationResponse
func (client *Client) CreateStateConfigurationWithOptions(tmpReq *CreateStateConfigurationRequest, runtime *util.RuntimeOptions) (_result *CreateStateConfigurationResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateStateConfigurationShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Tags)) {
		request.TagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tags, tea.String("Tags"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.ConfigureMode)) {
		query["ConfigureMode"] = request.ConfigureMode
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Parameters)) {
		query["Parameters"] = request.Parameters
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ScheduleExpression)) {
		query["ScheduleExpression"] = request.ScheduleExpression
	}

	if !tea.BoolValue(util.IsUnset(request.ScheduleType)) {
		query["ScheduleType"] = request.ScheduleType
	}

	if !tea.BoolValue(util.IsUnset(request.TagsShrink)) {
		query["Tags"] = request.TagsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.Targets)) {
		query["Targets"] = request.Targets
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateName)) {
		query["TemplateName"] = request.TemplateName
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateVersion)) {
		query["TemplateVersion"] = request.TemplateVersion
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateStateConfiguration"),
		Version:     tea.String("2019-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateStateConfigurationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a desired-state configuration.
//
// @param request - CreateStateConfigurationRequest
//
// @return CreateStateConfigurationResponse
func (client *Client) CreateStateConfiguration(request *CreateStateConfigurationRequest) (_result *CreateStateConfigurationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateStateConfigurationResponse{}
	_body, _err := client.CreateStateConfigurationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a template.
//
// @param tmpReq - CreateTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateTemplateResponse
func (client *Client) CreateTemplateWithOptions(tmpReq *CreateTemplateRequest, runtime *util.RuntimeOptions) (_result *CreateTemplateResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateTemplateShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Tags)) {
		request.TagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tags, tea.String("Tags"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Content)) {
		query["Content"] = request.Content
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.TagsShrink)) {
		query["Tags"] = request.TagsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateName)) {
		query["TemplateName"] = request.TemplateName
	}

	if !tea.BoolValue(util.IsUnset(request.VersionName)) {
		query["VersionName"] = request.VersionName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateTemplate"),
		Version:     tea.String("2019-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a template.
//
// @param request - CreateTemplateRequest
//
// @return CreateTemplateResponse
func (client *Client) CreateTemplate(request *CreateTemplateRequest) (_result *CreateTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateTemplateResponse{}
	_body, _err := client.CreateTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes an application. You can call this operation only for the applications that reside in the China (Hangzhou) region. Use an endpoint of the China (Hangzhou) region.
//
// @param request - DeleteApplicationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteApplicationResponse
func (client *Client) DeleteApplicationWithOptions(request *DeleteApplicationRequest, runtime *util.RuntimeOptions) (_result *DeleteApplicationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Force)) {
		query["Force"] = request.Force
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RetainResource)) {
		query["RetainResource"] = request.RetainResource
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteApplication"),
		Version:     tea.String("2019-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteApplicationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an application. You can call this operation only for the applications that reside in the China (Hangzhou) region. Use an endpoint of the China (Hangzhou) region.
//
// @param request - DeleteApplicationRequest
//
// @return DeleteApplicationResponse
func (client *Client) DeleteApplication(request *DeleteApplicationRequest) (_result *DeleteApplicationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteApplicationResponse{}
	_body, _err := client.DeleteApplicationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes an application group. You can call this operation only for the application groups which reside in the China (Hangzhou) region. Use an endpoint of the China (Hangzhou) region.
//
// @param request - DeleteApplicationGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteApplicationGroupResponse
func (client *Client) DeleteApplicationGroupWithOptions(request *DeleteApplicationGroupRequest, runtime *util.RuntimeOptions) (_result *DeleteApplicationGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApplicationName)) {
		query["ApplicationName"] = request.ApplicationName
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RetainResource)) {
		query["RetainResource"] = request.RetainResource
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteApplicationGroup"),
		Version:     tea.String("2019-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteApplicationGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an application group. You can call this operation only for the application groups which reside in the China (Hangzhou) region. Use an endpoint of the China (Hangzhou) region.
//
// @param request - DeleteApplicationGroupRequest
//
// @return DeleteApplicationGroupResponse
func (client *Client) DeleteApplicationGroup(request *DeleteApplicationGroupRequest) (_result *DeleteApplicationGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteApplicationGroupResponse{}
	_body, _err := client.DeleteApplicationGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes multiple executions.
//
// @param request - DeleteExecutionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteExecutionsResponse
func (client *Client) DeleteExecutionsWithOptions(request *DeleteExecutionsRequest, runtime *util.RuntimeOptions) (_result *DeleteExecutionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ExecutionIds)) {
		query["ExecutionIds"] = request.ExecutionIds
	}

	if !tea.BoolValue(util.IsUnset(request.Force)) {
		query["Force"] = request.Force
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteExecutions"),
		Version:     tea.String("2019-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteExecutionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes multiple executions.
//
// @param request - DeleteExecutionsRequest
//
// @return DeleteExecutionsResponse
func (client *Client) DeleteExecutions(request *DeleteExecutionsRequest) (_result *DeleteExecutionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteExecutionsResponse{}
	_body, _err := client.DeleteExecutionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a common parameter.
//
// @param request - DeleteParameterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteParameterResponse
func (client *Client) DeleteParameterWithOptions(request *DeleteParameterRequest, runtime *util.RuntimeOptions) (_result *DeleteParameterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteParameter"),
		Version:     tea.String("2019-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteParameterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a common parameter.
//
// @param request - DeleteParameterRequest
//
// @return DeleteParameterResponse
func (client *Client) DeleteParameter(request *DeleteParameterRequest) (_result *DeleteParameterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteParameterResponse{}
	_body, _err := client.DeleteParameterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a patch baseline.
//
// @param request - DeletePatchBaselineRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeletePatchBaselineResponse
func (client *Client) DeletePatchBaselineWithOptions(request *DeletePatchBaselineRequest, runtime *util.RuntimeOptions) (_result *DeletePatchBaselineResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeletePatchBaseline"),
		Version:     tea.String("2019-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeletePatchBaselineResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a patch baseline.
//
// @param request - DeletePatchBaselineRequest
//
// @return DeletePatchBaselineResponse
func (client *Client) DeletePatchBaseline(request *DeletePatchBaselineRequest) (_result *DeletePatchBaselineResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeletePatchBaselineResponse{}
	_body, _err := client.DeletePatchBaselineWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes an encryption parameter. Make sure that you have the permissions to call the DeleteSecret operation before you call this operation.
//
// @param request - DeleteSecretParameterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteSecretParameterResponse
func (client *Client) DeleteSecretParameterWithOptions(request *DeleteSecretParameterRequest, runtime *util.RuntimeOptions) (_result *DeleteSecretParameterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteSecretParameter"),
		Version:     tea.String("2019-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteSecretParameterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an encryption parameter. Make sure that you have the permissions to call the DeleteSecret operation before you call this operation.
//
// @param request - DeleteSecretParameterRequest
//
// @return DeleteSecretParameterResponse
func (client *Client) DeleteSecretParameter(request *DeleteSecretParameterRequest) (_result *DeleteSecretParameterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteSecretParameterResponse{}
	_body, _err := client.DeleteSecretParameterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes multiple desired-state configurations at a time.
//
// @param request - DeleteStateConfigurationsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteStateConfigurationsResponse
func (client *Client) DeleteStateConfigurationsWithOptions(request *DeleteStateConfigurationsRequest, runtime *util.RuntimeOptions) (_result *DeleteStateConfigurationsResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.StateConfigurationIds)) {
		query["StateConfigurationIds"] = request.StateConfigurationIds
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteStateConfigurations"),
		Version:     tea.String("2019-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteStateConfigurationsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes multiple desired-state configurations at a time.
//
// @param request - DeleteStateConfigurationsRequest
//
// @return DeleteStateConfigurationsResponse
func (client *Client) DeleteStateConfigurations(request *DeleteStateConfigurationsRequest) (_result *DeleteStateConfigurationsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteStateConfigurationsResponse{}
	_body, _err := client.DeleteStateConfigurationsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a template.
//
// @param request - DeleteTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteTemplateResponse
func (client *Client) DeleteTemplateWithOptions(request *DeleteTemplateRequest, runtime *util.RuntimeOptions) (_result *DeleteTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AutoDeleteExecutions)) {
		query["AutoDeleteExecutions"] = request.AutoDeleteExecutions
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateName)) {
		query["TemplateName"] = request.TemplateName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteTemplate"),
		Version:     tea.String("2019-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a template.
//
// @param request - DeleteTemplateRequest
//
// @return DeleteTemplateResponse
func (client *Client) DeleteTemplate(request *DeleteTemplateRequest) (_result *DeleteTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteTemplateResponse{}
	_body, _err := client.DeleteTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes multiple templates.
//
// @param request - DeleteTemplatesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteTemplatesResponse
func (client *Client) DeleteTemplatesWithOptions(request *DeleteTemplatesRequest, runtime *util.RuntimeOptions) (_result *DeleteTemplatesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AutoDeleteExecutions)) {
		query["AutoDeleteExecutions"] = request.AutoDeleteExecutions
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateNames)) {
		query["TemplateNames"] = request.TemplateNames
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteTemplates"),
		Version:     tea.String("2019-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteTemplatesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes multiple templates.
//
// @param request - DeleteTemplatesRequest
//
// @return DeleteTemplatesResponse
func (client *Client) DeleteTemplates(request *DeleteTemplatesRequest) (_result *DeleteTemplatesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteTemplatesResponse{}
	_body, _err := client.DeleteTemplatesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deploys an application group. You can call this operation only for the applications which reside in the China (Hangzhou) region. Use an endpoint of the China (Hangzhou) region.
//
// @param request - DeployApplicationGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeployApplicationGroupResponse
func (client *Client) DeployApplicationGroupWithOptions(request *DeployApplicationGroupRequest, runtime *util.RuntimeOptions) (_result *DeployApplicationGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApplicationName)) {
		query["ApplicationName"] = request.ApplicationName
	}

	if !tea.BoolValue(util.IsUnset(request.DeployParameters)) {
		query["DeployParameters"] = request.DeployParameters
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeployApplicationGroup"),
		Version:     tea.String("2019-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeployApplicationGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deploys an application group. You can call this operation only for the applications which reside in the China (Hangzhou) region. Use an endpoint of the China (Hangzhou) region.
//
// @param request - DeployApplicationGroupRequest
//
// @return DeployApplicationGroupResponse
func (client *Client) DeployApplicationGroup(request *DeployApplicationGroupRequest) (_result *DeployApplicationGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeployApplicationGroupResponse{}
	_body, _err := client.DeployApplicationGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询应用分组资源成本
//
// @param request - DescribeApplicationGroupBillRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeApplicationGroupBillResponse
func (client *Client) DescribeApplicationGroupBillWithOptions(request *DescribeApplicationGroupBillRequest, runtime *util.RuntimeOptions) (_result *DescribeApplicationGroupBillResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApplicationName)) {
		query["ApplicationName"] = request.ApplicationName
	}

	if !tea.BoolValue(util.IsUnset(request.BillingCycle)) {
		query["BillingCycle"] = request.BillingCycle
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeApplicationGroupBill"),
		Version:     tea.String("2019-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeApplicationGroupBillResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询应用分组资源成本
//
// @param request - DescribeApplicationGroupBillRequest
//
// @return DescribeApplicationGroupBillResponse
func (client *Client) DescribeApplicationGroupBill(request *DescribeApplicationGroupBillRequest) (_result *DescribeApplicationGroupBillResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeApplicationGroupBillResponse{}
	_body, _err := client.DescribeApplicationGroupBillWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries supported regions.
//
// @param request - DescribeRegionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRegionsResponse
func (client *Client) DescribeRegionsWithOptions(request *DescribeRegionsRequest, runtime *util.RuntimeOptions) (_result *DescribeRegionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AcceptLanguage)) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeRegions"),
		Version:     tea.String("2019-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeRegionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries supported regions.
//
// @param request - DescribeRegionsRequest
//
// @return DescribeRegionsResponse
func (client *Client) DescribeRegions(request *DescribeRegionsRequest) (_result *DescribeRegionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeRegionsResponse{}
	_body, _err := client.DescribeRegionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the Resource Access Management (RAM) policy required for template execution.
//
// @param request - GenerateExecutionPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GenerateExecutionPolicyResponse
func (client *Client) GenerateExecutionPolicyWithOptions(request *GenerateExecutionPolicyRequest, runtime *util.RuntimeOptions) (_result *GenerateExecutionPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RamRole)) {
		query["RamRole"] = request.RamRole
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateContent)) {
		query["TemplateContent"] = request.TemplateContent
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateName)) {
		query["TemplateName"] = request.TemplateName
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateVersion)) {
		query["TemplateVersion"] = request.TemplateVersion
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GenerateExecutionPolicy"),
		Version:     tea.String("2019-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GenerateExecutionPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the Resource Access Management (RAM) policy required for template execution.
//
// @param request - GenerateExecutionPolicyRequest
//
// @return GenerateExecutionPolicyResponse
func (client *Client) GenerateExecutionPolicy(request *GenerateExecutionPolicyRequest) (_result *GenerateExecutionPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GenerateExecutionPolicyResponse{}
	_body, _err := client.GenerateExecutionPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information of an application. You can call this operation only for the applications that reside in the China (Hangzhou) region. Use an endpoint of the China (Hangzhou) region.
//
// @param request - GetApplicationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetApplicationResponse
func (client *Client) GetApplicationWithOptions(request *GetApplicationRequest, runtime *util.RuntimeOptions) (_result *GetApplicationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetApplication"),
		Version:     tea.String("2019-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetApplicationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information of an application. You can call this operation only for the applications that reside in the China (Hangzhou) region. Use an endpoint of the China (Hangzhou) region.
//
// @param request - GetApplicationRequest
//
// @return GetApplicationResponse
func (client *Client) GetApplication(request *GetApplicationRequest) (_result *GetApplicationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetApplicationResponse{}
	_body, _err := client.GetApplicationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about an application group. You can call this operation only for the application groups that reside in the China (Hangzhou) region. Use an endpoint of the China (Hangzhou) region.
//
// @param request - GetApplicationGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetApplicationGroupResponse
func (client *Client) GetApplicationGroupWithOptions(request *GetApplicationGroupRequest, runtime *util.RuntimeOptions) (_result *GetApplicationGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApplicationName)) {
		query["ApplicationName"] = request.ApplicationName
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetApplicationGroup"),
		Version:     tea.String("2019-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetApplicationGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about an application group. You can call this operation only for the application groups that reside in the China (Hangzhou) region. Use an endpoint of the China (Hangzhou) region.
//
// @param request - GetApplicationGroupRequest
//
// @return GetApplicationGroupResponse
func (client *Client) GetApplicationGroup(request *GetApplicationGroupRequest) (_result *GetApplicationGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetApplicationGroupResponse{}
	_body, _err := client.GetApplicationGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about the template of an execution, including the content of the template.
//
// @param request - GetExecutionTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetExecutionTemplateResponse
func (client *Client) GetExecutionTemplateWithOptions(request *GetExecutionTemplateRequest, runtime *util.RuntimeOptions) (_result *GetExecutionTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ExecutionId)) {
		query["ExecutionId"] = request.ExecutionId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetExecutionTemplate"),
		Version:     tea.String("2019-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetExecutionTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about the template of an execution, including the content of the template.
//
// @param request - GetExecutionTemplateRequest
//
// @return GetExecutionTemplateResponse
func (client *Client) GetExecutionTemplate(request *GetExecutionTemplateRequest) (_result *GetExecutionTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetExecutionTemplateResponse{}
	_body, _err := client.GetExecutionTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the properties of a configuration list.
//
// @param request - GetInventorySchemaRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetInventorySchemaResponse
func (client *Client) GetInventorySchemaWithOptions(request *GetInventorySchemaRequest, runtime *util.RuntimeOptions) (_result *GetInventorySchemaResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Aggregator)) {
		query["Aggregator"] = request.Aggregator
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

	if !tea.BoolValue(util.IsUnset(request.TypeName)) {
		query["TypeName"] = request.TypeName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetInventorySchema"),
		Version:     tea.String("2019-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetInventorySchemaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the properties of a configuration list.
//
// @param request - GetInventorySchemaRequest
//
// @return GetInventorySchemaResponse
func (client *Client) GetInventorySchema(request *GetInventorySchemaRequest) (_result *GetInventorySchemaResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetInventorySchemaResponse{}
	_body, _err := client.GetInventorySchemaWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of an O\\\\\\\\\\\\&M item.
//
// @param request - GetOpsItemRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetOpsItemResponse
func (client *Client) GetOpsItemWithOptions(request *GetOpsItemRequest, runtime *util.RuntimeOptions) (_result *GetOpsItemResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OpsItemId)) {
		query["OpsItemId"] = request.OpsItemId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetOpsItem"),
		Version:     tea.String("2019-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetOpsItemResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of an O\\\\\\\\\\\\&M item.
//
// @param request - GetOpsItemRequest
//
// @return GetOpsItemResponse
func (client *Client) GetOpsItem(request *GetOpsItemRequest) (_result *GetOpsItemResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetOpsItemResponse{}
	_body, _err := client.GetOpsItemWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a common parameter and its value.
//
// @param request - GetParameterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetParameterResponse
func (client *Client) GetParameterWithOptions(request *GetParameterRequest, runtime *util.RuntimeOptions) (_result *GetParameterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.ParameterVersion)) {
		query["ParameterVersion"] = request.ParameterVersion
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetParameter"),
		Version:     tea.String("2019-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetParameterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a common parameter and its value.
//
// @param request - GetParameterRequest
//
// @return GetParameterResponse
func (client *Client) GetParameter(request *GetParameterRequest) (_result *GetParameterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetParameterResponse{}
	_body, _err := client.GetParameterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about one or more parameters.
//
// @param request - GetParametersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetParametersResponse
func (client *Client) GetParametersWithOptions(request *GetParametersRequest, runtime *util.RuntimeOptions) (_result *GetParametersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Names)) {
		query["Names"] = request.Names
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetParameters"),
		Version:     tea.String("2019-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetParametersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about one or more parameters.
//
// @param request - GetParametersRequest
//
// @return GetParametersResponse
func (client *Client) GetParameters(request *GetParametersRequest) (_result *GetParametersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetParametersResponse{}
	_body, _err := client.GetParametersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries one or more parameters by path.
//
// @param request - GetParametersByPathRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetParametersByPathResponse
func (client *Client) GetParametersByPathWithOptions(request *GetParametersByPathRequest, runtime *util.RuntimeOptions) (_result *GetParametersByPathResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.Path)) {
		query["Path"] = request.Path
	}

	if !tea.BoolValue(util.IsUnset(request.Recursive)) {
		query["Recursive"] = request.Recursive
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetParametersByPath"),
		Version:     tea.String("2019-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetParametersByPathResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries one or more parameters by path.
//
// @param request - GetParametersByPathRequest
//
// @return GetParametersByPathResponse
func (client *Client) GetParametersByPath(request *GetParametersByPathRequest) (_result *GetParametersByPathResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetParametersByPathResponse{}
	_body, _err := client.GetParametersByPathWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information of a patch baseline.
//
// @param request - GetPatchBaselineRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPatchBaselineResponse
func (client *Client) GetPatchBaselineWithOptions(request *GetPatchBaselineRequest, runtime *util.RuntimeOptions) (_result *GetPatchBaselineResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetPatchBaseline"),
		Version:     tea.String("2019-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetPatchBaselineResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information of a patch baseline.
//
// @param request - GetPatchBaselineRequest
//
// @return GetPatchBaselineResponse
func (client *Client) GetPatchBaseline(request *GetPatchBaselineRequest) (_result *GetPatchBaselineResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetPatchBaselineResponse{}
	_body, _err := client.GetPatchBaselineWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about an encryption parameter, including the parameter value. Make sure that you have the permissions to call the GetSecretValue operation before you call this operation.
//
// @param request - GetSecretParameterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSecretParameterResponse
func (client *Client) GetSecretParameterWithOptions(request *GetSecretParameterRequest, runtime *util.RuntimeOptions) (_result *GetSecretParameterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.ParameterVersion)) {
		query["ParameterVersion"] = request.ParameterVersion
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.WithDecryption)) {
		query["WithDecryption"] = request.WithDecryption
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetSecretParameter"),
		Version:     tea.String("2019-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetSecretParameterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about an encryption parameter, including the parameter value. Make sure that you have the permissions to call the GetSecretValue operation before you call this operation.
//
// @param request - GetSecretParameterRequest
//
// @return GetSecretParameterResponse
func (client *Client) GetSecretParameter(request *GetSecretParameterRequest) (_result *GetSecretParameterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetSecretParameterResponse{}
	_body, _err := client.GetSecretParameterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about one or more encryption parameters. Make sure that you have the permissions to call the GetSecretValue operation before you call this operation.
//
// @param request - GetSecretParametersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSecretParametersResponse
func (client *Client) GetSecretParametersWithOptions(request *GetSecretParametersRequest, runtime *util.RuntimeOptions) (_result *GetSecretParametersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Names)) {
		query["Names"] = request.Names
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.WithDecryption)) {
		query["WithDecryption"] = request.WithDecryption
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetSecretParameters"),
		Version:     tea.String("2019-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetSecretParametersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about one or more encryption parameters. Make sure that you have the permissions to call the GetSecretValue operation before you call this operation.
//
// @param request - GetSecretParametersRequest
//
// @return GetSecretParametersResponse
func (client *Client) GetSecretParameters(request *GetSecretParametersRequest) (_result *GetSecretParametersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetSecretParametersResponse{}
	_body, _err := client.GetSecretParametersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries encryption parameters by path. Make sure that you have the permissions to call the GetSecretValue operation before you call this operation.
//
// @param request - GetSecretParametersByPathRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSecretParametersByPathResponse
func (client *Client) GetSecretParametersByPathWithOptions(request *GetSecretParametersByPathRequest, runtime *util.RuntimeOptions) (_result *GetSecretParametersByPathResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.Path)) {
		query["Path"] = request.Path
	}

	if !tea.BoolValue(util.IsUnset(request.Recursive)) {
		query["Recursive"] = request.Recursive
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.WithDecryption)) {
		query["WithDecryption"] = request.WithDecryption
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetSecretParametersByPath"),
		Version:     tea.String("2019-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetSecretParametersByPathResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries encryption parameters by path. Make sure that you have the permissions to call the GetSecretValue operation before you call this operation.
//
// @param request - GetSecretParametersByPathRequest
//
// @return GetSecretParametersByPathResponse
func (client *Client) GetSecretParametersByPath(request *GetSecretParametersByPathRequest) (_result *GetSecretParametersByPathResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetSecretParametersByPathResponse{}
	_body, _err := client.GetSecretParametersByPathWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the settings of the delivery feature.
//
// @param request - GetServiceSettingsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetServiceSettingsResponse
func (client *Client) GetServiceSettingsWithOptions(request *GetServiceSettingsRequest, runtime *util.RuntimeOptions) (_result *GetServiceSettingsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetServiceSettings"),
		Version:     tea.String("2019-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetServiceSettingsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the settings of the delivery feature.
//
// @param request - GetServiceSettingsRequest
//
// @return GetServiceSettingsResponse
func (client *Client) GetServiceSettings(request *GetServiceSettingsRequest) (_result *GetServiceSettingsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetServiceSettingsResponse{}
	_body, _err := client.GetServiceSettingsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about a template, including the content of the template.
//
// @param request - GetTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTemplateResponse
func (client *Client) GetTemplateWithOptions(request *GetTemplateRequest, runtime *util.RuntimeOptions) (_result *GetTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateName)) {
		query["TemplateName"] = request.TemplateName
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateVersion)) {
		query["TemplateVersion"] = request.TemplateVersion
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetTemplate"),
		Version:     tea.String("2019-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a template, including the content of the template.
//
// @param request - GetTemplateRequest
//
// @return GetTemplateResponse
func (client *Client) GetTemplate(request *GetTemplateRequest) (_result *GetTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetTemplateResponse{}
	_body, _err := client.GetTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the available actions, including atomic actions and cloud product actions.
//
// @param request - ListActionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListActionsResponse
func (client *Client) ListActionsWithOptions(request *ListActionsRequest, runtime *util.RuntimeOptions) (_result *ListActionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.OOSActionName)) {
		query["OOSActionName"] = request.OOSActionName
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListActions"),
		Version:     tea.String("2019-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListActionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the available actions, including atomic actions and cloud product actions.
//
// @param request - ListActionsRequest
//
// @return ListActionsResponse
func (client *Client) ListActions(request *ListActionsRequest) (_result *ListActionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListActionsResponse{}
	_body, _err := client.ListActionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of application groups. You can call this operation only for the application groups that reside in the China (Hangzhou) region. Use an endpoint of the China (Hangzhou) region.
//
// @param request - ListApplicationGroupsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListApplicationGroupsResponse
func (client *Client) ListApplicationGroupsWithOptions(request *ListApplicationGroupsRequest, runtime *util.RuntimeOptions) (_result *ListApplicationGroupsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApplicationName)) {
		query["ApplicationName"] = request.ApplicationName
	}

	if !tea.BoolValue(util.IsUnset(request.DeployRegionId)) {
		query["DeployRegionId"] = request.DeployRegionId
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

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceProduct)) {
		query["ResourceProduct"] = request.ResourceProduct
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListApplicationGroups"),
		Version:     tea.String("2019-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListApplicationGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of application groups. You can call this operation only for the application groups that reside in the China (Hangzhou) region. Use an endpoint of the China (Hangzhou) region.
//
// @param request - ListApplicationGroupsRequest
//
// @return ListApplicationGroupsResponse
func (client *Client) ListApplicationGroups(request *ListApplicationGroupsRequest) (_result *ListApplicationGroupsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListApplicationGroupsResponse{}
	_body, _err := client.ListApplicationGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of applications. You can call this operation only for the applications that reside in the China (Hangzhou) region. Use an endpoint of the China (Hangzhou) region.
//
// @param tmpReq - ListApplicationsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListApplicationsResponse
func (client *Client) ListApplicationsWithOptions(tmpReq *ListApplicationsRequest, runtime *util.RuntimeOptions) (_result *ListApplicationsResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ListApplicationsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Tags)) {
		request.TagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tags, tea.String("Tags"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApplicationType)) {
		query["ApplicationType"] = request.ApplicationType
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Names)) {
		query["Names"] = request.Names
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.TagsShrink)) {
		query["Tags"] = request.TagsShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListApplications"),
		Version:     tea.String("2019-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListApplicationsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of applications. You can call this operation only for the applications that reside in the China (Hangzhou) region. Use an endpoint of the China (Hangzhou) region.
//
// @param request - ListApplicationsRequest
//
// @return ListApplicationsResponse
func (client *Client) ListApplications(request *ListApplicationsRequest) (_result *ListApplicationsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListApplicationsResponse{}
	_body, _err := client.ListApplicationsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the logs of an execution.
//
// Description:
//
// ***
//
// @param request - ListExecutionLogsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListExecutionLogsResponse
func (client *Client) ListExecutionLogsWithOptions(request *ListExecutionLogsRequest, runtime *util.RuntimeOptions) (_result *ListExecutionLogsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ExecutionId)) {
		query["ExecutionId"] = request.ExecutionId
	}

	if !tea.BoolValue(util.IsUnset(request.LogType)) {
		query["LogType"] = request.LogType
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

	if !tea.BoolValue(util.IsUnset(request.TaskExecutionId)) {
		query["TaskExecutionId"] = request.TaskExecutionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListExecutionLogs"),
		Version:     tea.String("2019-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListExecutionLogsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the logs of an execution.
//
// Description:
//
// ***
//
// @param request - ListExecutionLogsRequest
//
// @return ListExecutionLogsResponse
func (client *Client) ListExecutionLogs(request *ListExecutionLogsRequest) (_result *ListExecutionLogsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListExecutionLogsResponse{}
	_body, _err := client.ListExecutionLogsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries high-risk tasks in the execution of a template.
//
// @param request - ListExecutionRiskyTasksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListExecutionRiskyTasksResponse
func (client *Client) ListExecutionRiskyTasksWithOptions(request *ListExecutionRiskyTasksRequest, runtime *util.RuntimeOptions) (_result *ListExecutionRiskyTasksResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateName)) {
		query["TemplateName"] = request.TemplateName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListExecutionRiskyTasks"),
		Version:     tea.String("2019-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListExecutionRiskyTasksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries high-risk tasks in the execution of a template.
//
// @param request - ListExecutionRiskyTasksRequest
//
// @return ListExecutionRiskyTasksResponse
func (client *Client) ListExecutionRiskyTasks(request *ListExecutionRiskyTasksRequest) (_result *ListExecutionRiskyTasksResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListExecutionRiskyTasksResponse{}
	_body, _err := client.ListExecutionRiskyTasksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries executions. Multiple methods are supported to filter executions.
//
// @param tmpReq - ListExecutionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListExecutionsResponse
func (client *Client) ListExecutionsWithOptions(tmpReq *ListExecutionsRequest, runtime *util.RuntimeOptions) (_result *ListExecutionsResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ListExecutionsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Tags)) {
		request.TagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tags, tea.String("Tags"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Categories)) {
		query["Categories"] = request.Categories
	}

	if !tea.BoolValue(util.IsUnset(request.Category)) {
		query["Category"] = request.Category
	}

	if !tea.BoolValue(util.IsUnset(request.Depth)) {
		query["Depth"] = request.Depth
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.EndDateAfter)) {
		query["EndDateAfter"] = request.EndDateAfter
	}

	if !tea.BoolValue(util.IsUnset(request.EndDateBefore)) {
		query["EndDateBefore"] = request.EndDateBefore
	}

	if !tea.BoolValue(util.IsUnset(request.ExecutedBy)) {
		query["ExecutedBy"] = request.ExecutedBy
	}

	if !tea.BoolValue(util.IsUnset(request.ExecutionId)) {
		query["ExecutionId"] = request.ExecutionId
	}

	if !tea.BoolValue(util.IsUnset(request.IncludeChildExecution)) {
		query["IncludeChildExecution"] = request.IncludeChildExecution
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.Mode)) {
		query["Mode"] = request.Mode
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.ParentExecutionId)) {
		query["ParentExecutionId"] = request.ParentExecutionId
	}

	if !tea.BoolValue(util.IsUnset(request.RamRole)) {
		query["RamRole"] = request.RamRole
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceTemplateName)) {
		query["ResourceTemplateName"] = request.ResourceTemplateName
	}

	if !tea.BoolValue(util.IsUnset(request.SortField)) {
		query["SortField"] = request.SortField
	}

	if !tea.BoolValue(util.IsUnset(request.SortOrder)) {
		query["SortOrder"] = request.SortOrder
	}

	if !tea.BoolValue(util.IsUnset(request.StartDateAfter)) {
		query["StartDateAfter"] = request.StartDateAfter
	}

	if !tea.BoolValue(util.IsUnset(request.StartDateBefore)) {
		query["StartDateBefore"] = request.StartDateBefore
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.TagsShrink)) {
		query["Tags"] = request.TagsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateName)) {
		query["TemplateName"] = request.TemplateName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListExecutions"),
		Version:     tea.String("2019-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListExecutionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries executions. Multiple methods are supported to filter executions.
//
// @param request - ListExecutionsRequest
//
// @return ListExecutionsResponse
func (client *Client) ListExecutions(request *ListExecutionsRequest) (_result *ListExecutionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListExecutionsResponse{}
	_body, _err := client.ListExecutionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取仓库信息
//
// @param request - ListGitRepositoriesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListGitRepositoriesResponse
func (client *Client) ListGitRepositoriesWithOptions(request *ListGitRepositoriesRequest, runtime *util.RuntimeOptions) (_result *ListGitRepositoriesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.OrgName)) {
		query["OrgName"] = request.OrgName
	}

	if !tea.BoolValue(util.IsUnset(request.Owner)) {
		query["Owner"] = request.Owner
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Platform)) {
		query["Platform"] = request.Platform
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListGitRepositories"),
		Version:     tea.String("2019-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListGitRepositoriesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取仓库信息
//
// @param request - ListGitRepositoriesRequest
//
// @return ListGitRepositoriesResponse
func (client *Client) ListGitRepositories(request *ListGitRepositoriesRequest) (_result *ListGitRepositoriesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListGitRepositoriesResponse{}
	_body, _err := client.ListGitRepositoriesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 列出实例软件包状态
//
// @param request - ListInstancePackageStatesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListInstancePackageStatesResponse
func (client *Client) ListInstancePackageStatesWithOptions(request *ListInstancePackageStatesRequest, runtime *util.RuntimeOptions) (_result *ListInstancePackageStatesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
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

	if !tea.BoolValue(util.IsUnset(request.TemplateNames)) {
		query["TemplateNames"] = request.TemplateNames
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListInstancePackageStates"),
		Version:     tea.String("2019-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListInstancePackageStatesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 列出实例软件包状态
//
// @param request - ListInstancePackageStatesRequest
//
// @return ListInstancePackageStatesResponse
func (client *Client) ListInstancePackageStates(request *ListInstancePackageStatesRequest) (_result *ListInstancePackageStatesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListInstancePackageStatesResponse{}
	_body, _err := client.ListInstancePackageStatesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about the patches of an instance.
//
// @param request - ListInstancePatchStatesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListInstancePatchStatesResponse
func (client *Client) ListInstancePatchStatesWithOptions(request *ListInstancePatchStatesRequest, runtime *util.RuntimeOptions) (_result *ListInstancePatchStatesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceIds)) {
		query["InstanceIds"] = request.InstanceIds
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

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListInstancePatchStates"),
		Version:     tea.String("2019-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListInstancePatchStatesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about the patches of an instance.
//
// @param request - ListInstancePatchStatesRequest
//
// @return ListInstancePatchStatesResponse
func (client *Client) ListInstancePatchStates(request *ListInstancePatchStatesRequest) (_result *ListInstancePatchStatesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListInstancePatchStatesResponse{}
	_body, _err := client.ListInstancePatchStatesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about the patches of an instance.
//
// @param request - ListInstancePatchesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListInstancePatchesResponse
func (client *Client) ListInstancePatchesWithOptions(request *ListInstancePatchesRequest, runtime *util.RuntimeOptions) (_result *ListInstancePatchesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.PatchStatuses)) {
		query["PatchStatuses"] = request.PatchStatuses
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListInstancePatches"),
		Version:     tea.String("2019-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListInstancePatchesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about the patches of an instance.
//
// @param request - ListInstancePatchesRequest
//
// @return ListInstancePatchesResponse
func (client *Client) ListInstancePatches(request *ListInstancePatchesRequest) (_result *ListInstancePatchesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListInstancePatchesResponse{}
	_body, _err := client.ListInstancePatchesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the configurations of an Elastic Compute Service (ECS) instance.
//
// @param request - ListInventoryEntriesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListInventoryEntriesResponse
func (client *Client) ListInventoryEntriesWithOptions(request *ListInventoryEntriesRequest, runtime *util.RuntimeOptions) (_result *ListInventoryEntriesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Filter)) {
		query["Filter"] = request.Filter
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
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

	if !tea.BoolValue(util.IsUnset(request.TypeName)) {
		query["TypeName"] = request.TypeName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListInventoryEntries"),
		Version:     tea.String("2019-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListInventoryEntriesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the configurations of an Elastic Compute Service (ECS) instance.
//
// @param request - ListInventoryEntriesRequest
//
// @return ListInventoryEntriesResponse
func (client *Client) ListInventoryEntries(request *ListInventoryEntriesRequest) (_result *ListInventoryEntriesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListInventoryEntriesResponse{}
	_body, _err := client.ListInventoryEntriesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries O\\&M items.
//
// @param tmpReq - ListOpsItemsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListOpsItemsResponse
func (client *Client) ListOpsItemsWithOptions(tmpReq *ListOpsItemsRequest, runtime *util.RuntimeOptions) (_result *ListOpsItemsResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ListOpsItemsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.ResourceTags)) {
		request.ResourceTagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ResourceTags, tea.String("ResourceTags"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Tags)) {
		request.TagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tags, tea.String("Tags"), tea.String("json"))
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

	if !tea.BoolValue(util.IsUnset(request.ResourceTagsShrink)) {
		query["ResourceTags"] = request.ResourceTagsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.TagsShrink)) {
		query["Tags"] = request.TagsShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListOpsItems"),
		Version:     tea.String("2019-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListOpsItemsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries O\\&M items.
//
// @param request - ListOpsItemsRequest
//
// @return ListOpsItemsResponse
func (client *Client) ListOpsItems(request *ListOpsItemsRequest) (_result *ListOpsItemsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListOpsItemsResponse{}
	_body, _err := client.ListOpsItemsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the versions of a common parameter.
//
// @param request - ListParameterVersionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListParameterVersionsResponse
func (client *Client) ListParameterVersionsWithOptions(request *ListParameterVersionsRequest, runtime *util.RuntimeOptions) (_result *ListParameterVersionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ShareType)) {
		query["ShareType"] = request.ShareType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListParameterVersions"),
		Version:     tea.String("2019-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListParameterVersionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the versions of a common parameter.
//
// @param request - ListParameterVersionsRequest
//
// @return ListParameterVersionsResponse
func (client *Client) ListParameterVersions(request *ListParameterVersionsRequest) (_result *ListParameterVersionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListParameterVersionsResponse{}
	_body, _err := client.ListParameterVersionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries common parameters. Multiple methods are supported to filter common parameters.
//
// @param tmpReq - ListParametersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListParametersResponse
func (client *Client) ListParametersWithOptions(tmpReq *ListParametersRequest, runtime *util.RuntimeOptions) (_result *ListParametersResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ListParametersShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Tags)) {
		request.TagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tags, tea.String("Tags"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.Path)) {
		query["Path"] = request.Path
	}

	if !tea.BoolValue(util.IsUnset(request.Recursive)) {
		query["Recursive"] = request.Recursive
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ShareType)) {
		query["ShareType"] = request.ShareType
	}

	if !tea.BoolValue(util.IsUnset(request.SortField)) {
		query["SortField"] = request.SortField
	}

	if !tea.BoolValue(util.IsUnset(request.SortOrder)) {
		query["SortOrder"] = request.SortOrder
	}

	if !tea.BoolValue(util.IsUnset(request.TagsShrink)) {
		query["Tags"] = request.TagsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["Type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListParameters"),
		Version:     tea.String("2019-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListParametersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries common parameters. Multiple methods are supported to filter common parameters.
//
// @param request - ListParametersRequest
//
// @return ListParametersResponse
func (client *Client) ListParameters(request *ListParametersRequest) (_result *ListParametersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListParametersResponse{}
	_body, _err := client.ListParametersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of patch baselines.
//
// @param tmpReq - ListPatchBaselinesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPatchBaselinesResponse
func (client *Client) ListPatchBaselinesWithOptions(tmpReq *ListPatchBaselinesRequest, runtime *util.RuntimeOptions) (_result *ListPatchBaselinesResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ListPatchBaselinesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.ApprovedPatches)) {
		request.ApprovedPatchesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ApprovedPatches, tea.String("ApprovedPatches"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Sources)) {
		request.SourcesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Sources, tea.String("Sources"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Tags)) {
		request.TagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tags, tea.String("Tags"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApprovedPatchesShrink)) {
		query["ApprovedPatches"] = request.ApprovedPatchesShrink
	}

	if !tea.BoolValue(util.IsUnset(request.ApprovedPatchesEnableNonSecurity)) {
		query["ApprovedPatchesEnableNonSecurity"] = request.ApprovedPatchesEnableNonSecurity
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.OperationSystem)) {
		query["OperationSystem"] = request.OperationSystem
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ShareType)) {
		query["ShareType"] = request.ShareType
	}

	if !tea.BoolValue(util.IsUnset(request.SourcesShrink)) {
		query["Sources"] = request.SourcesShrink
	}

	if !tea.BoolValue(util.IsUnset(request.TagsShrink)) {
		query["Tags"] = request.TagsShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListPatchBaselines"),
		Version:     tea.String("2019-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListPatchBaselinesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of patch baselines.
//
// @param request - ListPatchBaselinesRequest
//
// @return ListPatchBaselinesResponse
func (client *Client) ListPatchBaselines(request *ListPatchBaselinesRequest) (_result *ListPatchBaselinesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListPatchBaselinesResponse{}
	_body, _err := client.ListPatchBaselinesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about a scheduled execution that involves O&M operations on Elastic Compute Service (ECS) instances.
//
// @param request - ListResourceExecutionStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListResourceExecutionStatusResponse
func (client *Client) ListResourceExecutionStatusWithOptions(request *ListResourceExecutionStatusRequest, runtime *util.RuntimeOptions) (_result *ListResourceExecutionStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ExecutionId)) {
		query["ExecutionId"] = request.ExecutionId
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

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListResourceExecutionStatus"),
		Version:     tea.String("2019-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListResourceExecutionStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a scheduled execution that involves O&M operations on Elastic Compute Service (ECS) instances.
//
// @param request - ListResourceExecutionStatusRequest
//
// @return ListResourceExecutionStatusResponse
func (client *Client) ListResourceExecutionStatus(request *ListResourceExecutionStatusRequest) (_result *ListResourceExecutionStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListResourceExecutionStatusResponse{}
	_body, _err := client.ListResourceExecutionStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries versions of an encryption parameter.
//
// @param request - ListSecretParameterVersionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSecretParameterVersionsResponse
func (client *Client) ListSecretParameterVersionsWithOptions(request *ListSecretParameterVersionsRequest, runtime *util.RuntimeOptions) (_result *ListSecretParameterVersionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ShareType)) {
		query["ShareType"] = request.ShareType
	}

	if !tea.BoolValue(util.IsUnset(request.WithDecryption)) {
		query["WithDecryption"] = request.WithDecryption
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListSecretParameterVersions"),
		Version:     tea.String("2019-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListSecretParameterVersionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries versions of an encryption parameter.
//
// @param request - ListSecretParameterVersionsRequest
//
// @return ListSecretParameterVersionsResponse
func (client *Client) ListSecretParameterVersions(request *ListSecretParameterVersionsRequest) (_result *ListSecretParameterVersionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListSecretParameterVersionsResponse{}
	_body, _err := client.ListSecretParameterVersionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries common parameters. Multiple types of queries are supported.
//
// Description:
//
// Before you call this operation, make sure that you have the permission to manage Key Management Service (KMS) secrets.
//
// @param tmpReq - ListSecretParametersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSecretParametersResponse
func (client *Client) ListSecretParametersWithOptions(tmpReq *ListSecretParametersRequest, runtime *util.RuntimeOptions) (_result *ListSecretParametersResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ListSecretParametersShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Tags)) {
		request.TagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tags, tea.String("Tags"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.Path)) {
		query["Path"] = request.Path
	}

	if !tea.BoolValue(util.IsUnset(request.Recursive)) {
		query["Recursive"] = request.Recursive
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.SortField)) {
		query["SortField"] = request.SortField
	}

	if !tea.BoolValue(util.IsUnset(request.SortOrder)) {
		query["SortOrder"] = request.SortOrder
	}

	if !tea.BoolValue(util.IsUnset(request.TagsShrink)) {
		query["Tags"] = request.TagsShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListSecretParameters"),
		Version:     tea.String("2019-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListSecretParametersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries common parameters. Multiple types of queries are supported.
//
// Description:
//
// Before you call this operation, make sure that you have the permission to manage Key Management Service (KMS) secrets.
//
// @param request - ListSecretParametersRequest
//
// @return ListSecretParametersResponse
func (client *Client) ListSecretParameters(request *ListSecretParametersRequest) (_result *ListSecretParametersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListSecretParametersResponse{}
	_body, _err := client.ListSecretParametersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries desired-state configurations.
//
// @param tmpReq - ListStateConfigurationsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListStateConfigurationsResponse
func (client *Client) ListStateConfigurationsWithOptions(tmpReq *ListStateConfigurationsRequest, runtime *util.RuntimeOptions) (_result *ListStateConfigurationsResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ListStateConfigurationsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Tags)) {
		request.TagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tags, tea.String("Tags"), tea.String("json"))
	}

	query := map[string]interface{}{}
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

	if !tea.BoolValue(util.IsUnset(request.StateConfigurationIds)) {
		query["StateConfigurationIds"] = request.StateConfigurationIds
	}

	if !tea.BoolValue(util.IsUnset(request.TagsShrink)) {
		query["Tags"] = request.TagsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateName)) {
		query["TemplateName"] = request.TemplateName
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateVersion)) {
		query["TemplateVersion"] = request.TemplateVersion
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListStateConfigurations"),
		Version:     tea.String("2019-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListStateConfigurationsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries desired-state configurations.
//
// @param request - ListStateConfigurationsRequest
//
// @return ListStateConfigurationsResponse
func (client *Client) ListStateConfigurations(request *ListStateConfigurationsRequest) (_result *ListStateConfigurationsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListStateConfigurationsResponse{}
	_body, _err := client.ListStateConfigurationsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the tags.
//
// @param request - ListTagKeysRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTagKeysResponse
func (client *Client) ListTagKeysWithOptions(request *ListTagKeysRequest, runtime *util.RuntimeOptions) (_result *ListTagKeysResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTagKeys"),
		Version:     tea.String("2019-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListTagKeysResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the tags.
//
// @param request - ListTagKeysRequest
//
// @return ListTagKeysResponse
func (client *Client) ListTagKeys(request *ListTagKeysRequest) (_result *ListTagKeysResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListTagKeysResponse{}
	_body, _err := client.ListTagKeysWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the tags that are added to one or more resources.
//
// @param tmpReq - ListTagResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTagResourcesResponse
func (client *Client) ListTagResourcesWithOptions(tmpReq *ListTagResourcesRequest, runtime *util.RuntimeOptions) (_result *ListTagResourcesResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ListTagResourcesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.ResourceIds)) {
		request.ResourceIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ResourceIds, tea.String("ResourceIds"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Tags)) {
		request.TagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tags, tea.String("Tags"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceIdsShrink)) {
		query["ResourceIds"] = request.ResourceIdsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.TagsShrink)) {
		query["Tags"] = request.TagsShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTagResources"),
		Version:     tea.String("2019-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListTagResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the tags that are added to one or more resources.
//
// @param request - ListTagResourcesRequest
//
// @return ListTagResourcesResponse
func (client *Client) ListTagResources(request *ListTagResourcesRequest) (_result *ListTagResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListTagResourcesResponse{}
	_body, _err := client.ListTagResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the values of created tags.
//
// @param request - ListTagValuesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTagValuesResponse
func (client *Client) ListTagValuesWithOptions(request *ListTagValuesRequest, runtime *util.RuntimeOptions) (_result *ListTagValuesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Key)) {
		query["Key"] = request.Key
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

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTagValues"),
		Version:     tea.String("2019-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListTagValuesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the values of created tags.
//
// @param request - ListTagValuesRequest
//
// @return ListTagValuesResponse
func (client *Client) ListTagValues(request *ListTagValuesRequest) (_result *ListTagValuesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListTagValuesResponse{}
	_body, _err := client.ListTagValuesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries task executions. Multiple methods are supported to filter task executions.
//
// @param request - ListTaskExecutionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTaskExecutionsResponse
func (client *Client) ListTaskExecutionsWithOptions(request *ListTaskExecutionsRequest, runtime *util.RuntimeOptions) (_result *ListTaskExecutionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndDateAfter)) {
		query["EndDateAfter"] = request.EndDateAfter
	}

	if !tea.BoolValue(util.IsUnset(request.EndDateBefore)) {
		query["EndDateBefore"] = request.EndDateBefore
	}

	if !tea.BoolValue(util.IsUnset(request.ExecutionId)) {
		query["ExecutionId"] = request.ExecutionId
	}

	if !tea.BoolValue(util.IsUnset(request.IncludeChildTaskExecution)) {
		query["IncludeChildTaskExecution"] = request.IncludeChildTaskExecution
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.ParentTaskExecutionId)) {
		query["ParentTaskExecutionId"] = request.ParentTaskExecutionId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SortField)) {
		query["SortField"] = request.SortField
	}

	if !tea.BoolValue(util.IsUnset(request.SortOrder)) {
		query["SortOrder"] = request.SortOrder
	}

	if !tea.BoolValue(util.IsUnset(request.StartDateAfter)) {
		query["StartDateAfter"] = request.StartDateAfter
	}

	if !tea.BoolValue(util.IsUnset(request.StartDateBefore)) {
		query["StartDateBefore"] = request.StartDateBefore
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.TaskAction)) {
		query["TaskAction"] = request.TaskAction
	}

	if !tea.BoolValue(util.IsUnset(request.TaskExecutionId)) {
		query["TaskExecutionId"] = request.TaskExecutionId
	}

	if !tea.BoolValue(util.IsUnset(request.TaskName)) {
		query["TaskName"] = request.TaskName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTaskExecutions"),
		Version:     tea.String("2019-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListTaskExecutionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries task executions. Multiple methods are supported to filter task executions.
//
// @param request - ListTaskExecutionsRequest
//
// @return ListTaskExecutionsResponse
func (client *Client) ListTaskExecutions(request *ListTaskExecutionsRequest) (_result *ListTaskExecutionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListTaskExecutionsResponse{}
	_body, _err := client.ListTaskExecutionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of versions of a template.
//
// @param request - ListTemplateVersionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTemplateVersionsResponse
func (client *Client) ListTemplateVersionsWithOptions(request *ListTemplateVersionsRequest, runtime *util.RuntimeOptions) (_result *ListTemplateVersionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ShareType)) {
		query["ShareType"] = request.ShareType
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateName)) {
		query["TemplateName"] = request.TemplateName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTemplateVersions"),
		Version:     tea.String("2019-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListTemplateVersionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of versions of a template.
//
// @param request - ListTemplateVersionsRequest
//
// @return ListTemplateVersionsResponse
func (client *Client) ListTemplateVersions(request *ListTemplateVersionsRequest) (_result *ListTemplateVersionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListTemplateVersionsResponse{}
	_body, _err := client.ListTemplateVersionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries templates. Multiple methods are supported to filter templates.
//
// @param tmpReq - ListTemplatesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTemplatesResponse
func (client *Client) ListTemplatesWithOptions(tmpReq *ListTemplatesRequest, runtime *util.RuntimeOptions) (_result *ListTemplatesResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ListTemplatesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Tags)) {
		request.TagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tags, tea.String("Tags"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Category)) {
		query["Category"] = request.Category
	}

	if !tea.BoolValue(util.IsUnset(request.CreatedBy)) {
		query["CreatedBy"] = request.CreatedBy
	}

	if !tea.BoolValue(util.IsUnset(request.CreatedDateAfter)) {
		query["CreatedDateAfter"] = request.CreatedDateAfter
	}

	if !tea.BoolValue(util.IsUnset(request.CreatedDateBefore)) {
		query["CreatedDateBefore"] = request.CreatedDateBefore
	}

	if !tea.BoolValue(util.IsUnset(request.HasTrigger)) {
		query["HasTrigger"] = request.HasTrigger
	}

	if !tea.BoolValue(util.IsUnset(request.IsExample)) {
		query["IsExample"] = request.IsExample
	}

	if !tea.BoolValue(util.IsUnset(request.IsFavorite)) {
		query["IsFavorite"] = request.IsFavorite
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

	if !tea.BoolValue(util.IsUnset(request.ShareType)) {
		query["ShareType"] = request.ShareType
	}

	if !tea.BoolValue(util.IsUnset(request.SortField)) {
		query["SortField"] = request.SortField
	}

	if !tea.BoolValue(util.IsUnset(request.SortOrder)) {
		query["SortOrder"] = request.SortOrder
	}

	if !tea.BoolValue(util.IsUnset(request.TagsShrink)) {
		query["Tags"] = request.TagsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateFormat)) {
		query["TemplateFormat"] = request.TemplateFormat
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateName)) {
		query["TemplateName"] = request.TemplateName
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateType)) {
		query["TemplateType"] = request.TemplateType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTemplates"),
		Version:     tea.String("2019-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListTemplatesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries templates. Multiple methods are supported to filter templates.
//
// @param request - ListTemplatesRequest
//
// @return ListTemplatesResponse
func (client *Client) ListTemplates(request *ListTemplatesRequest) (_result *ListTemplatesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListTemplatesResponse{}
	_body, _err := client.ListTemplatesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Notifies an execution in the Waiting state of the subsequent operations.
//
// Description:
//
// You can call this operation to notify an execution in the following scenarios:
//
// 	- If a template contains a special task, such as an approval task, the Operation Orchestration Service (OOS) execution engine sets the execution state to Waiting when the approval task is being run. You can call this operation to specify whether to continue the execution.
//
// 	- If you perform debugging in the debug mode, you can call this operation to notify the execution of the subsequent operations after the execution is created or a task is complete.
//
// 	- If a high-risk operation task waits for approval, you can call this operation to specify whether to continue the execution.
//
// @param request - NotifyExecutionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return NotifyExecutionResponse
func (client *Client) NotifyExecutionWithOptions(request *NotifyExecutionRequest, runtime *util.RuntimeOptions) (_result *NotifyExecutionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ExecutionId)) {
		query["ExecutionId"] = request.ExecutionId
	}

	if !tea.BoolValue(util.IsUnset(request.ExecutionStatus)) {
		query["ExecutionStatus"] = request.ExecutionStatus
	}

	if !tea.BoolValue(util.IsUnset(request.LoopItem)) {
		query["LoopItem"] = request.LoopItem
	}

	if !tea.BoolValue(util.IsUnset(request.NotifyNote)) {
		query["NotifyNote"] = request.NotifyNote
	}

	if !tea.BoolValue(util.IsUnset(request.NotifyType)) {
		query["NotifyType"] = request.NotifyType
	}

	if !tea.BoolValue(util.IsUnset(request.Parameters)) {
		query["Parameters"] = request.Parameters
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.TaskExecutionId)) {
		query["TaskExecutionId"] = request.TaskExecutionId
	}

	if !tea.BoolValue(util.IsUnset(request.TaskExecutionIds)) {
		query["TaskExecutionIds"] = request.TaskExecutionIds
	}

	if !tea.BoolValue(util.IsUnset(request.TaskName)) {
		query["TaskName"] = request.TaskName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("NotifyExecution"),
		Version:     tea.String("2019-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &NotifyExecutionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Notifies an execution in the Waiting state of the subsequent operations.
//
// Description:
//
// You can call this operation to notify an execution in the following scenarios:
//
// 	- If a template contains a special task, such as an approval task, the Operation Orchestration Service (OOS) execution engine sets the execution state to Waiting when the approval task is being run. You can call this operation to specify whether to continue the execution.
//
// 	- If you perform debugging in the debug mode, you can call this operation to notify the execution of the subsequent operations after the execution is created or a task is complete.
//
// 	- If a high-risk operation task waits for approval, you can call this operation to specify whether to continue the execution.
//
// @param request - NotifyExecutionRequest
//
// @return NotifyExecutionResponse
func (client *Client) NotifyExecution(request *NotifyExecutionRequest) (_result *NotifyExecutionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &NotifyExecutionResponse{}
	_body, _err := client.NotifyExecutionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Registers the default patch baseline.
//
// @param request - RegisterDefaultPatchBaselineRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RegisterDefaultPatchBaselineResponse
func (client *Client) RegisterDefaultPatchBaselineWithOptions(request *RegisterDefaultPatchBaselineRequest, runtime *util.RuntimeOptions) (_result *RegisterDefaultPatchBaselineResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RegisterDefaultPatchBaseline"),
		Version:     tea.String("2019-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RegisterDefaultPatchBaselineResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Registers the default patch baseline.
//
// @param request - RegisterDefaultPatchBaselineRequest
//
// @return RegisterDefaultPatchBaselineResponse
func (client *Client) RegisterDefaultPatchBaseline(request *RegisterDefaultPatchBaselineRequest) (_result *RegisterDefaultPatchBaselineResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RegisterDefaultPatchBaselineResponse{}
	_body, _err := client.RegisterDefaultPatchBaselineWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details or aggregate information of a configuration inventory.
//
// @param request - SearchInventoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SearchInventoryResponse
func (client *Client) SearchInventoryWithOptions(request *SearchInventoryRequest, runtime *util.RuntimeOptions) (_result *SearchInventoryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Aggregator)) {
		query["Aggregator"] = request.Aggregator
	}

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

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SearchInventory"),
		Version:     tea.String("2019-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SearchInventoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details or aggregate information of a configuration inventory.
//
// @param request - SearchInventoryRequest
//
// @return SearchInventoryResponse
func (client *Client) SearchInventory(request *SearchInventoryRequest) (_result *SearchInventoryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SearchInventoryResponse{}
	_body, _err := client.SearchInventoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Enables or disables the feature of delivering template execution records and sets the storage location.
//
// @param request - SetServiceSettingsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetServiceSettingsResponse
func (client *Client) SetServiceSettingsWithOptions(request *SetServiceSettingsRequest, runtime *util.RuntimeOptions) (_result *SetServiceSettingsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeliveryOssBucketName)) {
		query["DeliveryOssBucketName"] = request.DeliveryOssBucketName
	}

	if !tea.BoolValue(util.IsUnset(request.DeliveryOssEnabled)) {
		query["DeliveryOssEnabled"] = request.DeliveryOssEnabled
	}

	if !tea.BoolValue(util.IsUnset(request.DeliveryOssKeyPrefix)) {
		query["DeliveryOssKeyPrefix"] = request.DeliveryOssKeyPrefix
	}

	if !tea.BoolValue(util.IsUnset(request.DeliverySlsEnabled)) {
		query["DeliverySlsEnabled"] = request.DeliverySlsEnabled
	}

	if !tea.BoolValue(util.IsUnset(request.DeliverySlsProjectName)) {
		query["DeliverySlsProjectName"] = request.DeliverySlsProjectName
	}

	if !tea.BoolValue(util.IsUnset(request.RdcEnterpriseId)) {
		query["RdcEnterpriseId"] = request.RdcEnterpriseId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SetServiceSettings"),
		Version:     tea.String("2019-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SetServiceSettingsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Enables or disables the feature of delivering template execution records and sets the storage location.
//
// @param request - SetServiceSettingsRequest
//
// @return SetServiceSettingsResponse
func (client *Client) SetServiceSettings(request *SetServiceSettingsRequest) (_result *SetServiceSettingsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetServiceSettingsResponse{}
	_body, _err := client.SetServiceSettingsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Starts an execution.
//
// @param tmpReq - StartExecutionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartExecutionResponse
func (client *Client) StartExecutionWithOptions(tmpReq *StartExecutionRequest, runtime *util.RuntimeOptions) (_result *StartExecutionResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &StartExecutionShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Tags)) {
		request.TagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tags, tea.String("Tags"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.LoopMode)) {
		query["LoopMode"] = request.LoopMode
	}

	if !tea.BoolValue(util.IsUnset(request.Mode)) {
		query["Mode"] = request.Mode
	}

	if !tea.BoolValue(util.IsUnset(request.Parameters)) {
		query["Parameters"] = request.Parameters
	}

	if !tea.BoolValue(util.IsUnset(request.ParentExecutionId)) {
		query["ParentExecutionId"] = request.ParentExecutionId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.SafetyCheck)) {
		query["SafetyCheck"] = request.SafetyCheck
	}

	if !tea.BoolValue(util.IsUnset(request.TagsShrink)) {
		query["Tags"] = request.TagsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateContent)) {
		query["TemplateContent"] = request.TemplateContent
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateName)) {
		query["TemplateName"] = request.TemplateName
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateURL)) {
		query["TemplateURL"] = request.TemplateURL
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateVersion)) {
		query["TemplateVersion"] = request.TemplateVersion
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("StartExecution"),
		Version:     tea.String("2019-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &StartExecutionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Starts an execution.
//
// @param request - StartExecutionRequest
//
// @return StartExecutionResponse
func (client *Client) StartExecution(request *StartExecutionRequest) (_result *StartExecutionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StartExecutionResponse{}
	_body, _err := client.StartExecutionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds tags to one or more resources.
//
// @param tmpReq - TagResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TagResourcesResponse
func (client *Client) TagResourcesWithOptions(tmpReq *TagResourcesRequest, runtime *util.RuntimeOptions) (_result *TagResourcesResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &TagResourcesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.ResourceIds)) {
		request.ResourceIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ResourceIds, tea.String("ResourceIds"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Tags)) {
		request.TagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tags, tea.String("Tags"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceIdsShrink)) {
		query["ResourceIds"] = request.ResourceIdsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.TagsShrink)) {
		query["Tags"] = request.TagsShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("TagResources"),
		Version:     tea.String("2019-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &TagResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds tags to one or more resources.
//
// @param request - TagResourcesRequest
//
// @return TagResourcesResponse
func (client *Client) TagResources(request *TagResourcesRequest) (_result *TagResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &TagResourcesResponse{}
	_body, _err := client.TagResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Debugs a started execution that contains an event trigger task or alert trigger task. If the operation is called, a message body is sent to the event trigger task or alert trigger task. After the trigger task receives the message body, the trigger task generates a new child execution.
//
// @param request - TriggerExecutionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TriggerExecutionResponse
func (client *Client) TriggerExecutionWithOptions(request *TriggerExecutionRequest, runtime *util.RuntimeOptions) (_result *TriggerExecutionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Content)) {
		query["Content"] = request.Content
	}

	if !tea.BoolValue(util.IsUnset(request.ExecutionId)) {
		query["ExecutionId"] = request.ExecutionId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["Type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("TriggerExecution"),
		Version:     tea.String("2019-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &TriggerExecutionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Debugs a started execution that contains an event trigger task or alert trigger task. If the operation is called, a message body is sent to the event trigger task or alert trigger task. After the trigger task receives the message body, the trigger task generates a new child execution.
//
// @param request - TriggerExecutionRequest
//
// @return TriggerExecutionResponse
func (client *Client) TriggerExecution(request *TriggerExecutionRequest) (_result *TriggerExecutionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &TriggerExecutionResponse{}
	_body, _err := client.TriggerExecutionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Removes tags from one or more resources.
//
// @param tmpReq - UntagResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UntagResourcesResponse
func (client *Client) UntagResourcesWithOptions(tmpReq *UntagResourcesRequest, runtime *util.RuntimeOptions) (_result *UntagResourcesResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &UntagResourcesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.ResourceIds)) {
		request.ResourceIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ResourceIds, tea.String("ResourceIds"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.TagKeys)) {
		request.TagKeysShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TagKeys, tea.String("TagKeys"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.All)) {
		query["All"] = request.All
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceIdsShrink)) {
		query["ResourceIds"] = request.ResourceIdsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.TagKeysShrink)) {
		query["TagKeys"] = request.TagKeysShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UntagResources"),
		Version:     tea.String("2019-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UntagResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Removes tags from one or more resources.
//
// @param request - UntagResourcesRequest
//
// @return UntagResourcesResponse
func (client *Client) UntagResources(request *UntagResourcesRequest) (_result *UntagResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UntagResourcesResponse{}
	_body, _err := client.UntagResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates an application. You can call this operation only for the applications that reside in the China (Hangzhou) region. Use an endpoint of the China (Hangzhou) region.
//
// @param tmpReq - UpdateApplicationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateApplicationResponse
func (client *Client) UpdateApplicationWithOptions(tmpReq *UpdateApplicationRequest, runtime *util.RuntimeOptions) (_result *UpdateApplicationResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &UpdateApplicationShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.AlarmConfig)) {
		request.AlarmConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AlarmConfig, tea.String("AlarmConfig"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Tags)) {
		request.TagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tags, tea.String("Tags"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AlarmConfigShrink)) {
		query["AlarmConfig"] = request.AlarmConfigShrink
	}

	if !tea.BoolValue(util.IsUnset(request.DeleteAlarmRulesBeforeUpdate)) {
		query["DeleteAlarmRulesBeforeUpdate"] = request.DeleteAlarmRulesBeforeUpdate
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.TagsShrink)) {
		query["Tags"] = request.TagsShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateApplication"),
		Version:     tea.String("2019-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateApplicationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates an application. You can call this operation only for the applications that reside in the China (Hangzhou) region. Use an endpoint of the China (Hangzhou) region.
//
// @param request - UpdateApplicationRequest
//
// @return UpdateApplicationResponse
func (client *Client) UpdateApplication(request *UpdateApplicationRequest) (_result *UpdateApplicationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateApplicationResponse{}
	_body, _err := client.UpdateApplicationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the information of an application group. You can call this operation only for the application groups that reside in the China (Hangzhou) region. Use an endpoint of the China (Hangzhou) region.
//
// @param tmpReq - UpdateApplicationGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateApplicationGroupResponse
func (client *Client) UpdateApplicationGroupWithOptions(tmpReq *UpdateApplicationGroupRequest, runtime *util.RuntimeOptions) (_result *UpdateApplicationGroupResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &UpdateApplicationGroupShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Parameters)) {
		request.ParametersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Parameters, tea.String("Parameters"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApplicationName)) {
		query["ApplicationName"] = request.ApplicationName
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.NewName)) {
		query["NewName"] = request.NewName
	}

	if !tea.BoolValue(util.IsUnset(request.OperationName)) {
		query["OperationName"] = request.OperationName
	}

	if !tea.BoolValue(util.IsUnset(request.ParametersShrink)) {
		query["Parameters"] = request.ParametersShrink
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateApplicationGroup"),
		Version:     tea.String("2019-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateApplicationGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the information of an application group. You can call this operation only for the application groups that reside in the China (Hangzhou) region. Use an endpoint of the China (Hangzhou) region.
//
// @param request - UpdateApplicationGroupRequest
//
// @return UpdateApplicationGroupResponse
func (client *Client) UpdateApplicationGroup(request *UpdateApplicationGroupRequest) (_result *UpdateApplicationGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateApplicationGroupResponse{}
	_body, _err := client.UpdateApplicationGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates an execution.
//
// @param request - UpdateExecutionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateExecutionResponse
func (client *Client) UpdateExecutionWithOptions(request *UpdateExecutionRequest, runtime *util.RuntimeOptions) (_result *UpdateExecutionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.ExecutionId)) {
		query["ExecutionId"] = request.ExecutionId
	}

	if !tea.BoolValue(util.IsUnset(request.Parameters)) {
		query["Parameters"] = request.Parameters
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.Tags)) {
		query["Tags"] = request.Tags
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateExecution"),
		Version:     tea.String("2019-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateExecutionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates an execution.
//
// @param request - UpdateExecutionRequest
//
// @return UpdateExecutionResponse
func (client *Client) UpdateExecution(request *UpdateExecutionRequest) (_result *UpdateExecutionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateExecutionResponse{}
	_body, _err := client.UpdateExecutionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新实例软件包状态
//
// @param tmpReq - UpdateInstancePackageStateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateInstancePackageStateResponse
func (client *Client) UpdateInstancePackageStateWithOptions(tmpReq *UpdateInstancePackageStateRequest, runtime *util.RuntimeOptions) (_result *UpdateInstancePackageStateResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &UpdateInstancePackageStateShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Parameters)) {
		request.ParametersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Parameters, tea.String("Parameters"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConfigureAction)) {
		query["ConfigureAction"] = request.ConfigureAction
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.ParametersShrink)) {
		query["Parameters"] = request.ParametersShrink
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateName)) {
		query["TemplateName"] = request.TemplateName
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateVersion)) {
		query["TemplateVersion"] = request.TemplateVersion
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateInstancePackageState"),
		Version:     tea.String("2019-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateInstancePackageStateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新实例软件包状态
//
// @param request - UpdateInstancePackageStateRequest
//
// @return UpdateInstancePackageStateResponse
func (client *Client) UpdateInstancePackageState(request *UpdateInstancePackageStateRequest) (_result *UpdateInstancePackageStateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateInstancePackageStateResponse{}
	_body, _err := client.UpdateInstancePackageStateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies an O\\\\\\\\\\\\&M item.
//
// @param tmpReq - UpdateOpsItemRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateOpsItemResponse
func (client *Client) UpdateOpsItemWithOptions(tmpReq *UpdateOpsItemRequest, runtime *util.RuntimeOptions) (_result *UpdateOpsItemResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &UpdateOpsItemShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Tags)) {
		request.TagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tags, tea.String("Tags"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Category)) {
		query["Category"] = request.Category
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DedupString)) {
		query["DedupString"] = request.DedupString
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.OpsItemId)) {
		query["OpsItemId"] = request.OpsItemId
	}

	if !tea.BoolValue(util.IsUnset(request.Priority)) {
		query["Priority"] = request.Priority
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.Resources)) {
		query["Resources"] = request.Resources
	}

	if !tea.BoolValue(util.IsUnset(request.Severity)) {
		query["Severity"] = request.Severity
	}

	if !tea.BoolValue(util.IsUnset(request.Solutions)) {
		query["Solutions"] = request.Solutions
	}

	if !tea.BoolValue(util.IsUnset(request.Source)) {
		query["Source"] = request.Source
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.TagsShrink)) {
		query["Tags"] = request.TagsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.Title)) {
		query["Title"] = request.Title
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateOpsItem"),
		Version:     tea.String("2019-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateOpsItemResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies an O\\\\\\\\\\\\&M item.
//
// @param request - UpdateOpsItemRequest
//
// @return UpdateOpsItemResponse
func (client *Client) UpdateOpsItem(request *UpdateOpsItemRequest) (_result *UpdateOpsItemResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateOpsItemResponse{}
	_body, _err := client.UpdateOpsItemWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates a common parameter.
//
// @param request - UpdateParameterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateParameterResponse
func (client *Client) UpdateParameterWithOptions(request *UpdateParameterRequest, runtime *util.RuntimeOptions) (_result *UpdateParameterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.Tags)) {
		query["Tags"] = request.Tags
	}

	if !tea.BoolValue(util.IsUnset(request.Value)) {
		query["Value"] = request.Value
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateParameter"),
		Version:     tea.String("2019-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateParameterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a common parameter.
//
// @param request - UpdateParameterRequest
//
// @return UpdateParameterResponse
func (client *Client) UpdateParameter(request *UpdateParameterRequest) (_result *UpdateParameterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateParameterResponse{}
	_body, _err := client.UpdateParameterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates a patch baseline.
//
// @param tmpReq - UpdatePatchBaselineRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdatePatchBaselineResponse
func (client *Client) UpdatePatchBaselineWithOptions(tmpReq *UpdatePatchBaselineRequest, runtime *util.RuntimeOptions) (_result *UpdatePatchBaselineResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &UpdatePatchBaselineShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.ApprovedPatches)) {
		request.ApprovedPatchesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ApprovedPatches, tea.String("ApprovedPatches"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.RejectedPatches)) {
		request.RejectedPatchesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RejectedPatches, tea.String("RejectedPatches"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Sources)) {
		request.SourcesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Sources, tea.String("Sources"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Tags)) {
		request.TagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tags, tea.String("Tags"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApprovalRules)) {
		query["ApprovalRules"] = request.ApprovalRules
	}

	if !tea.BoolValue(util.IsUnset(request.ApprovedPatchesShrink)) {
		query["ApprovedPatches"] = request.ApprovedPatchesShrink
	}

	if !tea.BoolValue(util.IsUnset(request.ApprovedPatchesEnableNonSecurity)) {
		query["ApprovedPatchesEnableNonSecurity"] = request.ApprovedPatchesEnableNonSecurity
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RejectedPatchesShrink)) {
		query["RejectedPatches"] = request.RejectedPatchesShrink
	}

	if !tea.BoolValue(util.IsUnset(request.RejectedPatchesAction)) {
		query["RejectedPatchesAction"] = request.RejectedPatchesAction
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.SourcesShrink)) {
		query["Sources"] = request.SourcesShrink
	}

	if !tea.BoolValue(util.IsUnset(request.TagsShrink)) {
		query["Tags"] = request.TagsShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdatePatchBaseline"),
		Version:     tea.String("2019-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdatePatchBaselineResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a patch baseline.
//
// @param request - UpdatePatchBaselineRequest
//
// @return UpdatePatchBaselineResponse
func (client *Client) UpdatePatchBaseline(request *UpdatePatchBaselineRequest) (_result *UpdatePatchBaselineResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdatePatchBaselineResponse{}
	_body, _err := client.UpdatePatchBaselineWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates an encryption parameter.
//
// @param tmpReq - UpdateSecretParameterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateSecretParameterResponse
func (client *Client) UpdateSecretParameterWithOptions(tmpReq *UpdateSecretParameterRequest, runtime *util.RuntimeOptions) (_result *UpdateSecretParameterResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &UpdateSecretParameterShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Tags)) {
		request.TagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tags, tea.String("Tags"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.TagsShrink)) {
		query["Tags"] = request.TagsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.Value)) {
		query["Value"] = request.Value
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateSecretParameter"),
		Version:     tea.String("2019-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateSecretParameterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates an encryption parameter.
//
// @param request - UpdateSecretParameterRequest
//
// @return UpdateSecretParameterResponse
func (client *Client) UpdateSecretParameter(request *UpdateSecretParameterRequest) (_result *UpdateSecretParameterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateSecretParameterResponse{}
	_body, _err := client.UpdateSecretParameterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates a desired-state configuration.
//
// @param tmpReq - UpdateStateConfigurationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateStateConfigurationResponse
func (client *Client) UpdateStateConfigurationWithOptions(tmpReq *UpdateStateConfigurationRequest, runtime *util.RuntimeOptions) (_result *UpdateStateConfigurationResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &UpdateStateConfigurationShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Parameters)) {
		request.ParametersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Parameters, tea.String("Parameters"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Tags)) {
		request.TagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tags, tea.String("Tags"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.ConfigureMode)) {
		query["ConfigureMode"] = request.ConfigureMode
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
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

	if !tea.BoolValue(util.IsUnset(request.ScheduleExpression)) {
		query["ScheduleExpression"] = request.ScheduleExpression
	}

	if !tea.BoolValue(util.IsUnset(request.ScheduleType)) {
		query["ScheduleType"] = request.ScheduleType
	}

	if !tea.BoolValue(util.IsUnset(request.StateConfigurationId)) {
		query["StateConfigurationId"] = request.StateConfigurationId
	}

	if !tea.BoolValue(util.IsUnset(request.TagsShrink)) {
		query["Tags"] = request.TagsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.Targets)) {
		query["Targets"] = request.Targets
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateStateConfiguration"),
		Version:     tea.String("2019-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateStateConfigurationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a desired-state configuration.
//
// @param request - UpdateStateConfigurationRequest
//
// @return UpdateStateConfigurationResponse
func (client *Client) UpdateStateConfiguration(request *UpdateStateConfigurationRequest) (_result *UpdateStateConfigurationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateStateConfigurationResponse{}
	_body, _err := client.UpdateStateConfigurationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the information about an existing template.
//
// @param tmpReq - UpdateTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateTemplateResponse
func (client *Client) UpdateTemplateWithOptions(tmpReq *UpdateTemplateRequest, runtime *util.RuntimeOptions) (_result *UpdateTemplateResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &UpdateTemplateShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Tags)) {
		request.TagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tags, tea.String("Tags"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Content)) {
		query["Content"] = request.Content
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.TagsShrink)) {
		query["Tags"] = request.TagsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateName)) {
		query["TemplateName"] = request.TemplateName
	}

	if !tea.BoolValue(util.IsUnset(request.VersionName)) {
		query["VersionName"] = request.VersionName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateTemplate"),
		Version:     tea.String("2019-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the information about an existing template.
//
// @param request - UpdateTemplateRequest
//
// @return UpdateTemplateResponse
func (client *Client) UpdateTemplate(request *UpdateTemplateRequest) (_result *UpdateTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateTemplateResponse{}
	_body, _err := client.UpdateTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Check whether a template is valid.
//
// @param request - ValidateTemplateContentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ValidateTemplateContentResponse
func (client *Client) ValidateTemplateContentWithOptions(request *ValidateTemplateContentRequest, runtime *util.RuntimeOptions) (_result *ValidateTemplateContentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Content)) {
		query["Content"] = request.Content
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateURL)) {
		query["TemplateURL"] = request.TemplateURL
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ValidateTemplateContent"),
		Version:     tea.String("2019-06-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ValidateTemplateContentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Check whether a template is valid.
//
// @param request - ValidateTemplateContentRequest
//
// @return ValidateTemplateContentResponse
func (client *Client) ValidateTemplateContent(request *ValidateTemplateContentRequest) (_result *ValidateTemplateContentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ValidateTemplateContentResponse{}
	_body, _err := client.ValidateTemplateContentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
