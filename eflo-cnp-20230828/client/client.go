// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type DataResultsTaskIndividualResultMapValue struct {
	// Experiment ID
	//
	// example:
	//
	// 54
	ExperimentId *int64 `json:"ExperimentId,omitempty" xml:"ExperimentId,omitempty"`
	// Host IP
	//
	// example:
	//
	// p-jt-waf-app1
	Hostname *string `json:"Hostname,omitempty" xml:"Hostname,omitempty"`
	// Pod name
	//
	// example:
	//
	// fluxserv-6fc89b45cf-w8wq6
	PodName *string `json:"PodName,omitempty" xml:"PodName,omitempty"`
	// GPU数量
	//
	// example:
	//
	// 8
	GpuNum *int32 `json:"GpuNum,omitempty" xml:"GpuNum,omitempty"`
	// GPU名称
	//
	// example:
	//
	// 8x OAM 810 GPU
	GpuName *string `json:"GpuName,omitempty" xml:"GpuName,omitempty"`
	// Whether there is a warning
	//
	// example:
	//
	// false
	WarningFlag *bool `json:"WarningFlag,omitempty" xml:"WarningFlag,omitempty"`
	// Warning message
	//
	// example:
	//
	// warning message
	WarningMsg *string `json:"WarningMsg,omitempty" xml:"WarningMsg,omitempty"`
	// Whether there is an error
	//
	// example:
	//
	// false
	ErrorFlag *bool `json:"ErrorFlag,omitempty" xml:"ErrorFlag,omitempty"`
	// Error message
	//
	// example:
	//
	// error message
	ErrorMsg *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	// TFLOPS value
	//
	// example:
	//
	// 45
	Tflops *float64 `json:"Tflops,omitempty" xml:"Tflops,omitempty"`
	// Throughput
	//
	// example:
	//
	// 23
	SamplesPerSecond *float64 `json:"SamplesPerSecond,omitempty" xml:"SamplesPerSecond,omitempty"`
}

func (s DataResultsTaskIndividualResultMapValue) String() string {
	return tea.Prettify(s)
}

func (s DataResultsTaskIndividualResultMapValue) GoString() string {
	return s.String()
}

func (s *DataResultsTaskIndividualResultMapValue) SetExperimentId(v int64) *DataResultsTaskIndividualResultMapValue {
	s.ExperimentId = &v
	return s
}

func (s *DataResultsTaskIndividualResultMapValue) SetHostname(v string) *DataResultsTaskIndividualResultMapValue {
	s.Hostname = &v
	return s
}

func (s *DataResultsTaskIndividualResultMapValue) SetPodName(v string) *DataResultsTaskIndividualResultMapValue {
	s.PodName = &v
	return s
}

func (s *DataResultsTaskIndividualResultMapValue) SetGpuNum(v int32) *DataResultsTaskIndividualResultMapValue {
	s.GpuNum = &v
	return s
}

func (s *DataResultsTaskIndividualResultMapValue) SetGpuName(v string) *DataResultsTaskIndividualResultMapValue {
	s.GpuName = &v
	return s
}

func (s *DataResultsTaskIndividualResultMapValue) SetWarningFlag(v bool) *DataResultsTaskIndividualResultMapValue {
	s.WarningFlag = &v
	return s
}

func (s *DataResultsTaskIndividualResultMapValue) SetWarningMsg(v string) *DataResultsTaskIndividualResultMapValue {
	s.WarningMsg = &v
	return s
}

func (s *DataResultsTaskIndividualResultMapValue) SetErrorFlag(v bool) *DataResultsTaskIndividualResultMapValue {
	s.ErrorFlag = &v
	return s
}

func (s *DataResultsTaskIndividualResultMapValue) SetErrorMsg(v string) *DataResultsTaskIndividualResultMapValue {
	s.ErrorMsg = &v
	return s
}

func (s *DataResultsTaskIndividualResultMapValue) SetTflops(v float64) *DataResultsTaskIndividualResultMapValue {
	s.Tflops = &v
	return s
}

func (s *DataResultsTaskIndividualResultMapValue) SetSamplesPerSecond(v float64) *DataResultsTaskIndividualResultMapValue {
	s.SamplesPerSecond = &v
	return s
}

type ChangeResourceGroupRequest struct {
	// Region Id
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group id.
	//
	// This parameter is required.
	//
	// example:
	//
	// rg-aek25k3b4pbhc4a
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The resource id.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The resource type.
	//
	// example:
	//
	// ExperimentPlan
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s ChangeResourceGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s ChangeResourceGroupRequest) GoString() string {
	return s.String()
}

func (s *ChangeResourceGroupRequest) SetRegionId(v string) *ChangeResourceGroupRequest {
	s.RegionId = &v
	return s
}

func (s *ChangeResourceGroupRequest) SetResourceGroupId(v string) *ChangeResourceGroupRequest {
	s.ResourceGroupId = &v
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
	// Request Id
	//
	// example:
	//
	// 5514CB39-B7C0-5B89-8534-2DE1E0F2B7AB
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

type CheckServiceLinkedRoleEfloCnpForDeletingRequest struct {
	// The ID of the cloud account.
	//
	// example:
	//
	// 345678901234
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The ID of the deletion task.
	//
	// example:
	//
	// task-003
	DeletionTaskId *string `json:"DeletionTaskId,omitempty" xml:"DeletionTaskId,omitempty"`
	// The Alibaba Cloud Resource Name (ARN) of the RAM role.
	//
	// The trusted entity of the RAM role is an Alibaba Cloud account. For more information, see [Create a RAM role for a trusted Alibaba Cloud account](https://help.aliyun.com/document_detail/93691.html) or [CreateRole](https://help.aliyun.com/document_detail/28710.html).
	//
	// Format: `acs:ram::<account_id>:role/<role_name>`.
	//
	// You can view the ARN in the RAM console or by calling operations. The following items describe the validity periods of storage addresses:
	//
	// 	- For more information about how to view the ARN in the RAM console, see [How do I find the ARN of the RAM role?](https://help.aliyun.com/document_detail/39744.html)
	//
	// 	- For more information about how to view the ARN by calling operations, see [ListRoles](https://help.aliyun.com/document_detail/28713.html) or [GetRole](https://help.aliyun.com/document_detail/28711.html).
	//
	// example:
	//
	// arn:aws:iam::345678901234:role/eflo-cnp-role
	RoleArn *string `json:"RoleArn,omitempty" xml:"RoleArn,omitempty"`
	// The ID of the region.
	//
	// example:
	//
	// cn-hangzhou
	SPIRegionId *string `json:"SPIRegionId,omitempty" xml:"SPIRegionId,omitempty"`
	// The Service Name.
	//
	// example:
	//
	// eflo-cnp
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
}

func (s CheckServiceLinkedRoleEfloCnpForDeletingRequest) String() string {
	return tea.Prettify(s)
}

func (s CheckServiceLinkedRoleEfloCnpForDeletingRequest) GoString() string {
	return s.String()
}

func (s *CheckServiceLinkedRoleEfloCnpForDeletingRequest) SetAccountId(v string) *CheckServiceLinkedRoleEfloCnpForDeletingRequest {
	s.AccountId = &v
	return s
}

func (s *CheckServiceLinkedRoleEfloCnpForDeletingRequest) SetDeletionTaskId(v string) *CheckServiceLinkedRoleEfloCnpForDeletingRequest {
	s.DeletionTaskId = &v
	return s
}

func (s *CheckServiceLinkedRoleEfloCnpForDeletingRequest) SetRoleArn(v string) *CheckServiceLinkedRoleEfloCnpForDeletingRequest {
	s.RoleArn = &v
	return s
}

func (s *CheckServiceLinkedRoleEfloCnpForDeletingRequest) SetSPIRegionId(v string) *CheckServiceLinkedRoleEfloCnpForDeletingRequest {
	s.SPIRegionId = &v
	return s
}

func (s *CheckServiceLinkedRoleEfloCnpForDeletingRequest) SetServiceName(v string) *CheckServiceLinkedRoleEfloCnpForDeletingRequest {
	s.ServiceName = &v
	return s
}

type CheckServiceLinkedRoleEfloCnpForDeletingResponseBody struct {
	// Indicates whether the SLR can be deleted. Valid values:
	//
	// 	- `true`: The SLR can be deleted.
	//
	// 	- `false`: The SLR cannot be deleted.
	//
	// example:
	//
	// True
	Deletable *bool `json:"Deletable,omitempty" xml:"Deletable,omitempty"`
	// Request ID
	//
	// example:
	//
	// 6C212C4A-2CB3-56E6-BA2F-1CE2B03C5C94
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The resources.
	Resources []*CheckServiceLinkedRoleEfloCnpForDeletingResponseBodyResources `json:"Resources,omitempty" xml:"Resources,omitempty" type:"Repeated"`
}

func (s CheckServiceLinkedRoleEfloCnpForDeletingResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CheckServiceLinkedRoleEfloCnpForDeletingResponseBody) GoString() string {
	return s.String()
}

func (s *CheckServiceLinkedRoleEfloCnpForDeletingResponseBody) SetDeletable(v bool) *CheckServiceLinkedRoleEfloCnpForDeletingResponseBody {
	s.Deletable = &v
	return s
}

func (s *CheckServiceLinkedRoleEfloCnpForDeletingResponseBody) SetRequestId(v string) *CheckServiceLinkedRoleEfloCnpForDeletingResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckServiceLinkedRoleEfloCnpForDeletingResponseBody) SetResources(v []*CheckServiceLinkedRoleEfloCnpForDeletingResponseBodyResources) *CheckServiceLinkedRoleEfloCnpForDeletingResponseBody {
	s.Resources = v
	return s
}

type CheckServiceLinkedRoleEfloCnpForDeletingResponseBodyResources struct {
	// The region.
	//
	// example:
	//
	// cn-beijing
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The resources.
	Resources []*string `json:"Resources,omitempty" xml:"Resources,omitempty" type:"Repeated"`
}

func (s CheckServiceLinkedRoleEfloCnpForDeletingResponseBodyResources) String() string {
	return tea.Prettify(s)
}

func (s CheckServiceLinkedRoleEfloCnpForDeletingResponseBodyResources) GoString() string {
	return s.String()
}

func (s *CheckServiceLinkedRoleEfloCnpForDeletingResponseBodyResources) SetRegion(v string) *CheckServiceLinkedRoleEfloCnpForDeletingResponseBodyResources {
	s.Region = &v
	return s
}

func (s *CheckServiceLinkedRoleEfloCnpForDeletingResponseBodyResources) SetResources(v []*string) *CheckServiceLinkedRoleEfloCnpForDeletingResponseBodyResources {
	s.Resources = v
	return s
}

type CheckServiceLinkedRoleEfloCnpForDeletingResponse struct {
	Headers    map[string]*string                                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckServiceLinkedRoleEfloCnpForDeletingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckServiceLinkedRoleEfloCnpForDeletingResponse) String() string {
	return tea.Prettify(s)
}

func (s CheckServiceLinkedRoleEfloCnpForDeletingResponse) GoString() string {
	return s.String()
}

func (s *CheckServiceLinkedRoleEfloCnpForDeletingResponse) SetHeaders(v map[string]*string) *CheckServiceLinkedRoleEfloCnpForDeletingResponse {
	s.Headers = v
	return s
}

func (s *CheckServiceLinkedRoleEfloCnpForDeletingResponse) SetStatusCode(v int32) *CheckServiceLinkedRoleEfloCnpForDeletingResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckServiceLinkedRoleEfloCnpForDeletingResponse) SetBody(v *CheckServiceLinkedRoleEfloCnpForDeletingResponseBody) *CheckServiceLinkedRoleEfloCnpForDeletingResponse {
	s.Body = v
	return s
}

type CreateExperimentPlanRequest struct {
	// Additional parameters
	//
	// example:
	//
	// {}
	ExternalParams map[string]interface{} `json:"ExternalParams,omitempty" xml:"ExternalParams,omitempty"`
	// Plan Template Name
	//
	// example:
	//
	// test
	PlanTemplateName *string `json:"PlanTemplateName,omitempty" xml:"PlanTemplateName,omitempty"`
	// Resource group ID
	//
	// example:
	//
	// rg-aekzij65sf2rr5i
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// Resource ID
	//
	// This parameter is required.
	//
	// example:
	//
	// 189
	ResourceId *int64 `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// Resource tags
	Tag []*CreateExperimentPlanRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// Template ID
	//
	// This parameter is required.
	//
	// example:
	//
	// 349623
	TemplateId *int64 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s CreateExperimentPlanRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateExperimentPlanRequest) GoString() string {
	return s.String()
}

func (s *CreateExperimentPlanRequest) SetExternalParams(v map[string]interface{}) *CreateExperimentPlanRequest {
	s.ExternalParams = v
	return s
}

func (s *CreateExperimentPlanRequest) SetPlanTemplateName(v string) *CreateExperimentPlanRequest {
	s.PlanTemplateName = &v
	return s
}

func (s *CreateExperimentPlanRequest) SetResourceGroupId(v string) *CreateExperimentPlanRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateExperimentPlanRequest) SetResourceId(v int64) *CreateExperimentPlanRequest {
	s.ResourceId = &v
	return s
}

func (s *CreateExperimentPlanRequest) SetTag(v []*CreateExperimentPlanRequestTag) *CreateExperimentPlanRequest {
	s.Tag = v
	return s
}

func (s *CreateExperimentPlanRequest) SetTemplateId(v int64) *CreateExperimentPlanRequest {
	s.TemplateId = &v
	return s
}

type CreateExperimentPlanRequestTag struct {
	// Key
	//
	// example:
	//
	// owner
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// Value
	//
	// example:
	//
	// test
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateExperimentPlanRequestTag) String() string {
	return tea.Prettify(s)
}

func (s CreateExperimentPlanRequestTag) GoString() string {
	return s.String()
}

func (s *CreateExperimentPlanRequestTag) SetKey(v string) *CreateExperimentPlanRequestTag {
	s.Key = &v
	return s
}

func (s *CreateExperimentPlanRequestTag) SetValue(v string) *CreateExperimentPlanRequestTag {
	s.Value = &v
	return s
}

type CreateExperimentPlanShrinkRequest struct {
	// Additional parameters
	//
	// example:
	//
	// {}
	ExternalParamsShrink *string `json:"ExternalParams,omitempty" xml:"ExternalParams,omitempty"`
	// Plan Template Name
	//
	// example:
	//
	// test
	PlanTemplateName *string `json:"PlanTemplateName,omitempty" xml:"PlanTemplateName,omitempty"`
	// Resource group ID
	//
	// example:
	//
	// rg-aekzij65sf2rr5i
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// Resource ID
	//
	// This parameter is required.
	//
	// example:
	//
	// 189
	ResourceId *int64 `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// Resource tags
	Tag []*CreateExperimentPlanShrinkRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// Template ID
	//
	// This parameter is required.
	//
	// example:
	//
	// 349623
	TemplateId *int64 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s CreateExperimentPlanShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateExperimentPlanShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateExperimentPlanShrinkRequest) SetExternalParamsShrink(v string) *CreateExperimentPlanShrinkRequest {
	s.ExternalParamsShrink = &v
	return s
}

func (s *CreateExperimentPlanShrinkRequest) SetPlanTemplateName(v string) *CreateExperimentPlanShrinkRequest {
	s.PlanTemplateName = &v
	return s
}

func (s *CreateExperimentPlanShrinkRequest) SetResourceGroupId(v string) *CreateExperimentPlanShrinkRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateExperimentPlanShrinkRequest) SetResourceId(v int64) *CreateExperimentPlanShrinkRequest {
	s.ResourceId = &v
	return s
}

func (s *CreateExperimentPlanShrinkRequest) SetTag(v []*CreateExperimentPlanShrinkRequestTag) *CreateExperimentPlanShrinkRequest {
	s.Tag = v
	return s
}

func (s *CreateExperimentPlanShrinkRequest) SetTemplateId(v int64) *CreateExperimentPlanShrinkRequest {
	s.TemplateId = &v
	return s
}

type CreateExperimentPlanShrinkRequestTag struct {
	// Key
	//
	// example:
	//
	// owner
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// Value
	//
	// example:
	//
	// test
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateExperimentPlanShrinkRequestTag) String() string {
	return tea.Prettify(s)
}

func (s CreateExperimentPlanShrinkRequestTag) GoString() string {
	return s.String()
}

func (s *CreateExperimentPlanShrinkRequestTag) SetKey(v string) *CreateExperimentPlanShrinkRequestTag {
	s.Key = &v
	return s
}

func (s *CreateExperimentPlanShrinkRequestTag) SetValue(v string) *CreateExperimentPlanShrinkRequestTag {
	s.Value = &v
	return s
}

type CreateExperimentPlanResponseBody struct {
	// Access denied information
	//
	// example:
	//
	// {}
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// Data
	//
	// example:
	//
	// []
	Data *int64 `json:"Data,omitempty" xml:"Data,omitempty"`
	// Request ID
	//
	// example:
	//
	// 5514CB39-B7C0-5B89-8534-2DE1E0F2B7AB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Total count
	//
	// example:
	//
	// 0
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s CreateExperimentPlanResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateExperimentPlanResponseBody) GoString() string {
	return s.String()
}

func (s *CreateExperimentPlanResponseBody) SetAccessDeniedDetail(v string) *CreateExperimentPlanResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CreateExperimentPlanResponseBody) SetData(v int64) *CreateExperimentPlanResponseBody {
	s.Data = &v
	return s
}

func (s *CreateExperimentPlanResponseBody) SetRequestId(v string) *CreateExperimentPlanResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateExperimentPlanResponseBody) SetTotalCount(v int64) *CreateExperimentPlanResponseBody {
	s.TotalCount = &v
	return s
}

type CreateExperimentPlanResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateExperimentPlanResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateExperimentPlanResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateExperimentPlanResponse) GoString() string {
	return s.String()
}

func (s *CreateExperimentPlanResponse) SetHeaders(v map[string]*string) *CreateExperimentPlanResponse {
	s.Headers = v
	return s
}

func (s *CreateExperimentPlanResponse) SetStatusCode(v int32) *CreateExperimentPlanResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateExperimentPlanResponse) SetBody(v *CreateExperimentPlanResponseBody) *CreateExperimentPlanResponse {
	s.Body = v
	return s
}

type CreateExperimentPlanTemplateRequest struct {
	// Privacy Level
	//
	// This parameter is required.
	//
	// example:
	//
	// private
	PrivacyLevel *string `json:"PrivacyLevel,omitempty" xml:"PrivacyLevel,omitempty"`
	// Template Description
	//
	// example:
	//
	// The template installs jdk and tomcat on a new ECS instance.
	TemplateDescription *string `json:"TemplateDescription,omitempty" xml:"TemplateDescription,omitempty"`
	// Template ID
	//
	// example:
	//
	// 4724
	TemplateId *int64 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// Template Name
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// Template Pipeline
	//
	// This parameter is required.
	TemplatePipeline []*CreateExperimentPlanTemplateRequestTemplatePipeline `json:"TemplatePipeline,omitempty" xml:"TemplatePipeline,omitempty" type:"Repeated"`
}

func (s CreateExperimentPlanTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateExperimentPlanTemplateRequest) GoString() string {
	return s.String()
}

func (s *CreateExperimentPlanTemplateRequest) SetPrivacyLevel(v string) *CreateExperimentPlanTemplateRequest {
	s.PrivacyLevel = &v
	return s
}

func (s *CreateExperimentPlanTemplateRequest) SetTemplateDescription(v string) *CreateExperimentPlanTemplateRequest {
	s.TemplateDescription = &v
	return s
}

func (s *CreateExperimentPlanTemplateRequest) SetTemplateId(v int64) *CreateExperimentPlanTemplateRequest {
	s.TemplateId = &v
	return s
}

func (s *CreateExperimentPlanTemplateRequest) SetTemplateName(v string) *CreateExperimentPlanTemplateRequest {
	s.TemplateName = &v
	return s
}

func (s *CreateExperimentPlanTemplateRequest) SetTemplatePipeline(v []*CreateExperimentPlanTemplateRequestTemplatePipeline) *CreateExperimentPlanTemplateRequest {
	s.TemplatePipeline = v
	return s
}

type CreateExperimentPlanTemplateRequestTemplatePipeline struct {
	// Configured Environment Parameters
	//
	// This parameter is required.
	EnvParams *CreateExperimentPlanTemplateRequestTemplatePipelineEnvParams `json:"EnvParams,omitempty" xml:"EnvParams,omitempty" type:"Struct"`
	// Node Order Number
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PipelineOrder *int32 `json:"PipelineOrder,omitempty" xml:"PipelineOrder,omitempty"`
	// Usage Scenario, e.g., "baseline"
	//
	// This parameter is required.
	//
	// example:
	//
	// baseline
	Scene *string `json:"Scene,omitempty" xml:"Scene,omitempty"`
	// Configured Workload Parameters
	SettingParams map[string]*string `json:"SettingParams,omitempty" xml:"SettingParams,omitempty"`
	// Workload ID
	//
	// This parameter is required.
	//
	// example:
	//
	// 14
	WorkloadId *int64 `json:"WorkloadId,omitempty" xml:"WorkloadId,omitempty"`
	// Workload Name
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	WorkloadName *string `json:"WorkloadName,omitempty" xml:"WorkloadName,omitempty"`
}

func (s CreateExperimentPlanTemplateRequestTemplatePipeline) String() string {
	return tea.Prettify(s)
}

func (s CreateExperimentPlanTemplateRequestTemplatePipeline) GoString() string {
	return s.String()
}

func (s *CreateExperimentPlanTemplateRequestTemplatePipeline) SetEnvParams(v *CreateExperimentPlanTemplateRequestTemplatePipelineEnvParams) *CreateExperimentPlanTemplateRequestTemplatePipeline {
	s.EnvParams = v
	return s
}

func (s *CreateExperimentPlanTemplateRequestTemplatePipeline) SetPipelineOrder(v int32) *CreateExperimentPlanTemplateRequestTemplatePipeline {
	s.PipelineOrder = &v
	return s
}

func (s *CreateExperimentPlanTemplateRequestTemplatePipeline) SetScene(v string) *CreateExperimentPlanTemplateRequestTemplatePipeline {
	s.Scene = &v
	return s
}

func (s *CreateExperimentPlanTemplateRequestTemplatePipeline) SetSettingParams(v map[string]*string) *CreateExperimentPlanTemplateRequestTemplatePipeline {
	s.SettingParams = v
	return s
}

func (s *CreateExperimentPlanTemplateRequestTemplatePipeline) SetWorkloadId(v int64) *CreateExperimentPlanTemplateRequestTemplatePipeline {
	s.WorkloadId = &v
	return s
}

func (s *CreateExperimentPlanTemplateRequestTemplatePipeline) SetWorkloadName(v string) *CreateExperimentPlanTemplateRequestTemplatePipeline {
	s.WorkloadName = &v
	return s
}

type CreateExperimentPlanTemplateRequestTemplatePipelineEnvParams struct {
	// CPU allocation count
	//
	// This parameter is required.
	//
	// example:
	//
	// 90
	CpuPerWorker *int32 `json:"CpuPerWorker,omitempty" xml:"CpuPerWorker,omitempty"`
	// CUDA Version
	//
	// example:
	//
	// 1.0.0
	CudaVersion *string `json:"CudaVersion,omitempty" xml:"CudaVersion,omitempty"`
	// GPU Driver Version
	//
	// example:
	//
	// 1.0.0
	GpuDriverVersion *string `json:"GpuDriverVersion,omitempty" xml:"GpuDriverVersion,omitempty"`
	// GPU allocation count
	//
	// This parameter is required.
	//
	// example:
	//
	// 8
	GpuPerWorker *int32 `json:"GpuPerWorker,omitempty" xml:"GpuPerWorker,omitempty"`
	// Memory (GB) allocation count
	//
	// This parameter is required.
	//
	// example:
	//
	// 500
	MemoryPerWorker *int32 `json:"MemoryPerWorker,omitempty" xml:"MemoryPerWorker,omitempty"`
	// NCCL Version
	//
	// example:
	//
	// 1.0.0
	NCCLVersion *string `json:"NCCLVersion,omitempty" xml:"NCCLVersion,omitempty"`
	// PyTorch Version
	//
	// example:
	//
	// 1.0.0
	PyTorchVersion *string `json:"PyTorchVersion,omitempty" xml:"PyTorchVersion,omitempty"`
	// Shared Memory (GB) allocation count
	//
	// This parameter is required.
	//
	// example:
	//
	// 500
	ShareMemory *int32 `json:"ShareMemory,omitempty" xml:"ShareMemory,omitempty"`
	// Number of nodes
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	WorkerNum *int32 `json:"WorkerNum,omitempty" xml:"WorkerNum,omitempty"`
}

func (s CreateExperimentPlanTemplateRequestTemplatePipelineEnvParams) String() string {
	return tea.Prettify(s)
}

func (s CreateExperimentPlanTemplateRequestTemplatePipelineEnvParams) GoString() string {
	return s.String()
}

func (s *CreateExperimentPlanTemplateRequestTemplatePipelineEnvParams) SetCpuPerWorker(v int32) *CreateExperimentPlanTemplateRequestTemplatePipelineEnvParams {
	s.CpuPerWorker = &v
	return s
}

func (s *CreateExperimentPlanTemplateRequestTemplatePipelineEnvParams) SetCudaVersion(v string) *CreateExperimentPlanTemplateRequestTemplatePipelineEnvParams {
	s.CudaVersion = &v
	return s
}

func (s *CreateExperimentPlanTemplateRequestTemplatePipelineEnvParams) SetGpuDriverVersion(v string) *CreateExperimentPlanTemplateRequestTemplatePipelineEnvParams {
	s.GpuDriverVersion = &v
	return s
}

func (s *CreateExperimentPlanTemplateRequestTemplatePipelineEnvParams) SetGpuPerWorker(v int32) *CreateExperimentPlanTemplateRequestTemplatePipelineEnvParams {
	s.GpuPerWorker = &v
	return s
}

func (s *CreateExperimentPlanTemplateRequestTemplatePipelineEnvParams) SetMemoryPerWorker(v int32) *CreateExperimentPlanTemplateRequestTemplatePipelineEnvParams {
	s.MemoryPerWorker = &v
	return s
}

func (s *CreateExperimentPlanTemplateRequestTemplatePipelineEnvParams) SetNCCLVersion(v string) *CreateExperimentPlanTemplateRequestTemplatePipelineEnvParams {
	s.NCCLVersion = &v
	return s
}

func (s *CreateExperimentPlanTemplateRequestTemplatePipelineEnvParams) SetPyTorchVersion(v string) *CreateExperimentPlanTemplateRequestTemplatePipelineEnvParams {
	s.PyTorchVersion = &v
	return s
}

func (s *CreateExperimentPlanTemplateRequestTemplatePipelineEnvParams) SetShareMemory(v int32) *CreateExperimentPlanTemplateRequestTemplatePipelineEnvParams {
	s.ShareMemory = &v
	return s
}

func (s *CreateExperimentPlanTemplateRequestTemplatePipelineEnvParams) SetWorkerNum(v int32) *CreateExperimentPlanTemplateRequestTemplatePipelineEnvParams {
	s.WorkerNum = &v
	return s
}

type CreateExperimentPlanTemplateShrinkRequest struct {
	// Privacy Level
	//
	// This parameter is required.
	//
	// example:
	//
	// private
	PrivacyLevel *string `json:"PrivacyLevel,omitempty" xml:"PrivacyLevel,omitempty"`
	// Template Description
	//
	// example:
	//
	// The template installs jdk and tomcat on a new ECS instance.
	TemplateDescription *string `json:"TemplateDescription,omitempty" xml:"TemplateDescription,omitempty"`
	// Template ID
	//
	// example:
	//
	// 4724
	TemplateId *int64 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// Template Name
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// Template Pipeline
	//
	// This parameter is required.
	TemplatePipelineShrink *string `json:"TemplatePipeline,omitempty" xml:"TemplatePipeline,omitempty"`
}

func (s CreateExperimentPlanTemplateShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateExperimentPlanTemplateShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateExperimentPlanTemplateShrinkRequest) SetPrivacyLevel(v string) *CreateExperimentPlanTemplateShrinkRequest {
	s.PrivacyLevel = &v
	return s
}

func (s *CreateExperimentPlanTemplateShrinkRequest) SetTemplateDescription(v string) *CreateExperimentPlanTemplateShrinkRequest {
	s.TemplateDescription = &v
	return s
}

func (s *CreateExperimentPlanTemplateShrinkRequest) SetTemplateId(v int64) *CreateExperimentPlanTemplateShrinkRequest {
	s.TemplateId = &v
	return s
}

func (s *CreateExperimentPlanTemplateShrinkRequest) SetTemplateName(v string) *CreateExperimentPlanTemplateShrinkRequest {
	s.TemplateName = &v
	return s
}

func (s *CreateExperimentPlanTemplateShrinkRequest) SetTemplatePipelineShrink(v string) *CreateExperimentPlanTemplateShrinkRequest {
	s.TemplatePipelineShrink = &v
	return s
}

type CreateExperimentPlanTemplateResponseBody struct {
	// Access Denied Detail
	//
	// example:
	//
	// {}
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// Data
	Data *CreateExperimentPlanTemplateResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// Request ID
	//
	// example:
	//
	// 5514CB39-B7C0-5B89-8534-2DE1E0F2B7AB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// total
	//
	// example:
	//
	// 0
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s CreateExperimentPlanTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateExperimentPlanTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *CreateExperimentPlanTemplateResponseBody) SetAccessDeniedDetail(v string) *CreateExperimentPlanTemplateResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CreateExperimentPlanTemplateResponseBody) SetData(v *CreateExperimentPlanTemplateResponseBodyData) *CreateExperimentPlanTemplateResponseBody {
	s.Data = v
	return s
}

func (s *CreateExperimentPlanTemplateResponseBody) SetRequestId(v string) *CreateExperimentPlanTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateExperimentPlanTemplateResponseBody) SetTotalCount(v int64) *CreateExperimentPlanTemplateResponseBody {
	s.TotalCount = &v
	return s
}

type CreateExperimentPlanTemplateResponseBodyData struct {
	// Creation Time
	//
	// example:
	//
	// 2024-11-19T02:01:05Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// Primary Account UID
	//
	// example:
	//
	// 12312312312312
	CreatorUid *int64 `json:"CreatorUid,omitempty" xml:"CreatorUid,omitempty"`
	// Is Deleted
	//
	// example:
	//
	// 0
	IsDelete *int32 `json:"IsDelete,omitempty" xml:"IsDelete,omitempty"`
	// Privacy Level
	//
	// example:
	//
	// private
	PrivacyLevel *string `json:"PrivacyLevel,omitempty" xml:"PrivacyLevel,omitempty"`
	// Template Code
	//
	// example:
	//
	// 1
	TemplateCode *int64 `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
	// Template Description
	//
	// example:
	//
	// test
	TemplateDescription *string `json:"TemplateDescription,omitempty" xml:"TemplateDescription,omitempty"`
	// Template ID
	//
	// example:
	//
	// 17615126
	TemplateId *int64 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// Template Name
	//
	// example:
	//
	// test
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// Template Pipeline
	TemplatePipelineParam []*CreateExperimentPlanTemplateResponseBodyDataTemplatePipelineParam `json:"TemplatePipelineParam,omitempty" xml:"TemplatePipelineParam,omitempty" type:"Repeated"`
	// Update Time
	//
	// example:
	//
	// 2023-10-16T01:58Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// Version ID
	//
	// example:
	//
	// 1
	VersionId *int64 `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
}

func (s CreateExperimentPlanTemplateResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateExperimentPlanTemplateResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateExperimentPlanTemplateResponseBodyData) SetCreateTime(v string) *CreateExperimentPlanTemplateResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *CreateExperimentPlanTemplateResponseBodyData) SetCreatorUid(v int64) *CreateExperimentPlanTemplateResponseBodyData {
	s.CreatorUid = &v
	return s
}

func (s *CreateExperimentPlanTemplateResponseBodyData) SetIsDelete(v int32) *CreateExperimentPlanTemplateResponseBodyData {
	s.IsDelete = &v
	return s
}

func (s *CreateExperimentPlanTemplateResponseBodyData) SetPrivacyLevel(v string) *CreateExperimentPlanTemplateResponseBodyData {
	s.PrivacyLevel = &v
	return s
}

func (s *CreateExperimentPlanTemplateResponseBodyData) SetTemplateCode(v int64) *CreateExperimentPlanTemplateResponseBodyData {
	s.TemplateCode = &v
	return s
}

func (s *CreateExperimentPlanTemplateResponseBodyData) SetTemplateDescription(v string) *CreateExperimentPlanTemplateResponseBodyData {
	s.TemplateDescription = &v
	return s
}

func (s *CreateExperimentPlanTemplateResponseBodyData) SetTemplateId(v int64) *CreateExperimentPlanTemplateResponseBodyData {
	s.TemplateId = &v
	return s
}

func (s *CreateExperimentPlanTemplateResponseBodyData) SetTemplateName(v string) *CreateExperimentPlanTemplateResponseBodyData {
	s.TemplateName = &v
	return s
}

func (s *CreateExperimentPlanTemplateResponseBodyData) SetTemplatePipelineParam(v []*CreateExperimentPlanTemplateResponseBodyDataTemplatePipelineParam) *CreateExperimentPlanTemplateResponseBodyData {
	s.TemplatePipelineParam = v
	return s
}

func (s *CreateExperimentPlanTemplateResponseBodyData) SetUpdateTime(v string) *CreateExperimentPlanTemplateResponseBodyData {
	s.UpdateTime = &v
	return s
}

func (s *CreateExperimentPlanTemplateResponseBodyData) SetVersionId(v int64) *CreateExperimentPlanTemplateResponseBodyData {
	s.VersionId = &v
	return s
}

type CreateExperimentPlanTemplateResponseBodyDataTemplatePipelineParam struct {
	// Configured Environment Parameters
	EnvParams *CreateExperimentPlanTemplateResponseBodyDataTemplatePipelineParamEnvParams `json:"EnvParams,omitempty" xml:"EnvParams,omitempty" type:"Struct"`
	// Pipeline Order
	//
	// example:
	//
	// 1
	PipelineOrder *int32 `json:"PipelineOrder,omitempty" xml:"PipelineOrder,omitempty"`
	// Usage Scenario, e.g., "baseline"
	//
	// example:
	//
	// baseline
	Scene *string `json:"Scene,omitempty" xml:"Scene,omitempty"`
	// Configured Workload Parameters
	SettingParams map[string]*string `json:"SettingParams,omitempty" xml:"SettingParams,omitempty"`
	// Workload ID
	//
	// example:
	//
	// 13
	WorkloadId *int64 `json:"WorkloadId,omitempty" xml:"WorkloadId,omitempty"`
	// Workload Name
	//
	// example:
	//
	// test
	WorkloadName *string `json:"WorkloadName,omitempty" xml:"WorkloadName,omitempty"`
}

func (s CreateExperimentPlanTemplateResponseBodyDataTemplatePipelineParam) String() string {
	return tea.Prettify(s)
}

func (s CreateExperimentPlanTemplateResponseBodyDataTemplatePipelineParam) GoString() string {
	return s.String()
}

func (s *CreateExperimentPlanTemplateResponseBodyDataTemplatePipelineParam) SetEnvParams(v *CreateExperimentPlanTemplateResponseBodyDataTemplatePipelineParamEnvParams) *CreateExperimentPlanTemplateResponseBodyDataTemplatePipelineParam {
	s.EnvParams = v
	return s
}

func (s *CreateExperimentPlanTemplateResponseBodyDataTemplatePipelineParam) SetPipelineOrder(v int32) *CreateExperimentPlanTemplateResponseBodyDataTemplatePipelineParam {
	s.PipelineOrder = &v
	return s
}

func (s *CreateExperimentPlanTemplateResponseBodyDataTemplatePipelineParam) SetScene(v string) *CreateExperimentPlanTemplateResponseBodyDataTemplatePipelineParam {
	s.Scene = &v
	return s
}

func (s *CreateExperimentPlanTemplateResponseBodyDataTemplatePipelineParam) SetSettingParams(v map[string]*string) *CreateExperimentPlanTemplateResponseBodyDataTemplatePipelineParam {
	s.SettingParams = v
	return s
}

func (s *CreateExperimentPlanTemplateResponseBodyDataTemplatePipelineParam) SetWorkloadId(v int64) *CreateExperimentPlanTemplateResponseBodyDataTemplatePipelineParam {
	s.WorkloadId = &v
	return s
}

func (s *CreateExperimentPlanTemplateResponseBodyDataTemplatePipelineParam) SetWorkloadName(v string) *CreateExperimentPlanTemplateResponseBodyDataTemplatePipelineParam {
	s.WorkloadName = &v
	return s
}

type CreateExperimentPlanTemplateResponseBodyDataTemplatePipelineParamEnvParams struct {
	// CPU Allocation
	//
	// example:
	//
	// 90
	CpuPerWorker *int32 `json:"CpuPerWorker,omitempty" xml:"CpuPerWorker,omitempty"`
	// cudaVersion
	//
	// example:
	//
	// 1.0.0
	CudaVersion *string `json:"CudaVersion,omitempty" xml:"CudaVersion,omitempty"`
	// GPU Driver Version
	//
	// example:
	//
	// 1.0.0
	GpuDriverVersion *string `json:"GpuDriverVersion,omitempty" xml:"GpuDriverVersion,omitempty"`
	// GPU Allocation
	//
	// example:
	//
	// 8
	GpuPerWorker *int32 `json:"GpuPerWorker,omitempty" xml:"GpuPerWorker,omitempty"`
	// Memory (GB) Allocation
	//
	// example:
	//
	// 500
	MemoryPerWorker *int32 `json:"MemoryPerWorker,omitempty" xml:"MemoryPerWorker,omitempty"`
	// NCCL Version
	//
	// example:
	//
	// 1.0.0
	NCCLVersion *string `json:"NCCLVersion,omitempty" xml:"NCCLVersion,omitempty"`
	// PyTorch Version
	//
	// example:
	//
	// 1.0.0
	PyTorchVersion *string `json:"PyTorchVersion,omitempty" xml:"PyTorchVersion,omitempty"`
	// Shared Memory (GB) Allocation
	//
	// example:
	//
	// 500
	ShareMemory *int32 `json:"ShareMemory,omitempty" xml:"ShareMemory,omitempty"`
	// Number of Nodes
	//
	// example:
	//
	// 1
	WorkerNum *int32 `json:"WorkerNum,omitempty" xml:"WorkerNum,omitempty"`
}

func (s CreateExperimentPlanTemplateResponseBodyDataTemplatePipelineParamEnvParams) String() string {
	return tea.Prettify(s)
}

func (s CreateExperimentPlanTemplateResponseBodyDataTemplatePipelineParamEnvParams) GoString() string {
	return s.String()
}

func (s *CreateExperimentPlanTemplateResponseBodyDataTemplatePipelineParamEnvParams) SetCpuPerWorker(v int32) *CreateExperimentPlanTemplateResponseBodyDataTemplatePipelineParamEnvParams {
	s.CpuPerWorker = &v
	return s
}

func (s *CreateExperimentPlanTemplateResponseBodyDataTemplatePipelineParamEnvParams) SetCudaVersion(v string) *CreateExperimentPlanTemplateResponseBodyDataTemplatePipelineParamEnvParams {
	s.CudaVersion = &v
	return s
}

func (s *CreateExperimentPlanTemplateResponseBodyDataTemplatePipelineParamEnvParams) SetGpuDriverVersion(v string) *CreateExperimentPlanTemplateResponseBodyDataTemplatePipelineParamEnvParams {
	s.GpuDriverVersion = &v
	return s
}

func (s *CreateExperimentPlanTemplateResponseBodyDataTemplatePipelineParamEnvParams) SetGpuPerWorker(v int32) *CreateExperimentPlanTemplateResponseBodyDataTemplatePipelineParamEnvParams {
	s.GpuPerWorker = &v
	return s
}

func (s *CreateExperimentPlanTemplateResponseBodyDataTemplatePipelineParamEnvParams) SetMemoryPerWorker(v int32) *CreateExperimentPlanTemplateResponseBodyDataTemplatePipelineParamEnvParams {
	s.MemoryPerWorker = &v
	return s
}

func (s *CreateExperimentPlanTemplateResponseBodyDataTemplatePipelineParamEnvParams) SetNCCLVersion(v string) *CreateExperimentPlanTemplateResponseBodyDataTemplatePipelineParamEnvParams {
	s.NCCLVersion = &v
	return s
}

func (s *CreateExperimentPlanTemplateResponseBodyDataTemplatePipelineParamEnvParams) SetPyTorchVersion(v string) *CreateExperimentPlanTemplateResponseBodyDataTemplatePipelineParamEnvParams {
	s.PyTorchVersion = &v
	return s
}

func (s *CreateExperimentPlanTemplateResponseBodyDataTemplatePipelineParamEnvParams) SetShareMemory(v int32) *CreateExperimentPlanTemplateResponseBodyDataTemplatePipelineParamEnvParams {
	s.ShareMemory = &v
	return s
}

func (s *CreateExperimentPlanTemplateResponseBodyDataTemplatePipelineParamEnvParams) SetWorkerNum(v int32) *CreateExperimentPlanTemplateResponseBodyDataTemplatePipelineParamEnvParams {
	s.WorkerNum = &v
	return s
}

type CreateExperimentPlanTemplateResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateExperimentPlanTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateExperimentPlanTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateExperimentPlanTemplateResponse) GoString() string {
	return s.String()
}

func (s *CreateExperimentPlanTemplateResponse) SetHeaders(v map[string]*string) *CreateExperimentPlanTemplateResponse {
	s.Headers = v
	return s
}

func (s *CreateExperimentPlanTemplateResponse) SetStatusCode(v int32) *CreateExperimentPlanTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateExperimentPlanTemplateResponse) SetBody(v *CreateExperimentPlanTemplateResponseBody) *CreateExperimentPlanTemplateResponse {
	s.Body = v
	return s
}

type CreateResourceRequest struct {
	// Cluster Description
	//
	// example:
	//
	// ppu集群
	ClusterDesc *string `json:"ClusterDesc,omitempty" xml:"ClusterDesc,omitempty"`
	// Cluster ID
	//
	// This parameter is required.
	//
	// example:
	//
	// ehpc-sh-fj71c0ycfw
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// Cluster Name
	//
	// This parameter is required.
	//
	// example:
	//
	// tre-1-ppu
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// Machine Types
	//
	// This parameter is required.
	MachineTypes *CreateResourceRequestMachineTypes `json:"MachineTypes,omitempty" xml:"MachineTypes,omitempty" type:"Struct"`
	// User Access Parameters
	//
	// This parameter is required.
	UserAccessParam *CreateResourceRequestUserAccessParam `json:"UserAccessParam,omitempty" xml:"UserAccessParam,omitempty" type:"Struct"`
}

func (s CreateResourceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateResourceRequest) GoString() string {
	return s.String()
}

func (s *CreateResourceRequest) SetClusterDesc(v string) *CreateResourceRequest {
	s.ClusterDesc = &v
	return s
}

func (s *CreateResourceRequest) SetClusterId(v string) *CreateResourceRequest {
	s.ClusterId = &v
	return s
}

func (s *CreateResourceRequest) SetClusterName(v string) *CreateResourceRequest {
	s.ClusterName = &v
	return s
}

func (s *CreateResourceRequest) SetMachineTypes(v *CreateResourceRequestMachineTypes) *CreateResourceRequest {
	s.MachineTypes = v
	return s
}

func (s *CreateResourceRequest) SetUserAccessParam(v *CreateResourceRequestUserAccessParam) *CreateResourceRequest {
	s.UserAccessParam = v
	return s
}

type CreateResourceRequestMachineTypes struct {
	// Number of Network Bonds
	//
	// example:
	//
	// 5
	BondNum *int32 `json:"BondNum,omitempty" xml:"BondNum,omitempty"`
	// CPU Information
	//
	// This parameter is required.
	//
	// example:
	//
	// 2x Intel Saphhire Rapid 8469C 48C CPU
	CpuInfo *string `json:"CpuInfo,omitempty" xml:"CpuInfo,omitempty"`
	// Disk Information
	//
	// example:
	//
	// 2x 480GB SATA SSD \\n 4x 3.84TB NVMe SSD
	DiskInfo *string `json:"DiskInfo,omitempty" xml:"DiskInfo,omitempty"`
	// GPU Information
	//
	// This parameter is required.
	//
	// example:
	//
	// 8x NVIDIA SXM4 80GB A100 GPU
	GpuInfo *string `json:"GpuInfo,omitempty" xml:"GpuInfo,omitempty"`
	// Memory Information
	//
	// example:
	//
	// 32x 64GB DDR4 4800 Memory
	MemoryInfo *string `json:"MemoryInfo,omitempty" xml:"MemoryInfo,omitempty"`
	// Specification Name
	//
	// example:
	//
	// efg1.nvga1n
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Network Information
	//
	// example:
	//
	// 1x 200Gbps Dual Port BF3 DPU for VPC\\\\n4x 200Gbps Dual Port EIC
	NetworkInfo *string `json:"NetworkInfo,omitempty" xml:"NetworkInfo,omitempty"`
	// Network Mode
	//
	// example:
	//
	// 2
	NetworkMode *string `json:"NetworkMode,omitempty" xml:"NetworkMode,omitempty"`
	// Number of Nodes
	//
	// example:
	//
	// 1
	NodeCount *int32 `json:"NodeCount,omitempty" xml:"NodeCount,omitempty"`
	// Type
	//
	// example:
	//
	// Private
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateResourceRequestMachineTypes) String() string {
	return tea.Prettify(s)
}

func (s CreateResourceRequestMachineTypes) GoString() string {
	return s.String()
}

func (s *CreateResourceRequestMachineTypes) SetBondNum(v int32) *CreateResourceRequestMachineTypes {
	s.BondNum = &v
	return s
}

func (s *CreateResourceRequestMachineTypes) SetCpuInfo(v string) *CreateResourceRequestMachineTypes {
	s.CpuInfo = &v
	return s
}

func (s *CreateResourceRequestMachineTypes) SetDiskInfo(v string) *CreateResourceRequestMachineTypes {
	s.DiskInfo = &v
	return s
}

func (s *CreateResourceRequestMachineTypes) SetGpuInfo(v string) *CreateResourceRequestMachineTypes {
	s.GpuInfo = &v
	return s
}

func (s *CreateResourceRequestMachineTypes) SetMemoryInfo(v string) *CreateResourceRequestMachineTypes {
	s.MemoryInfo = &v
	return s
}

func (s *CreateResourceRequestMachineTypes) SetName(v string) *CreateResourceRequestMachineTypes {
	s.Name = &v
	return s
}

func (s *CreateResourceRequestMachineTypes) SetNetworkInfo(v string) *CreateResourceRequestMachineTypes {
	s.NetworkInfo = &v
	return s
}

func (s *CreateResourceRequestMachineTypes) SetNetworkMode(v string) *CreateResourceRequestMachineTypes {
	s.NetworkMode = &v
	return s
}

func (s *CreateResourceRequestMachineTypes) SetNodeCount(v int32) *CreateResourceRequestMachineTypes {
	s.NodeCount = &v
	return s
}

func (s *CreateResourceRequestMachineTypes) SetType(v string) *CreateResourceRequestMachineTypes {
	s.Type = &v
	return s
}

type CreateResourceRequestUserAccessParam struct {
	// User ID
	//
	// This parameter is required.
	//
	// example:
	//
	// dev
	AccessId *string `json:"AccessId,omitempty" xml:"AccessId,omitempty"`
	// User Key
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	AccessKey *string `json:"AccessKey,omitempty" xml:"AccessKey,omitempty"`
	// Endpoint
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	Endpoint *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	// Workspace ID
	//
	// This parameter is required.
	//
	// example:
	//
	// 1245688643
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s CreateResourceRequestUserAccessParam) String() string {
	return tea.Prettify(s)
}

func (s CreateResourceRequestUserAccessParam) GoString() string {
	return s.String()
}

func (s *CreateResourceRequestUserAccessParam) SetAccessId(v string) *CreateResourceRequestUserAccessParam {
	s.AccessId = &v
	return s
}

func (s *CreateResourceRequestUserAccessParam) SetAccessKey(v string) *CreateResourceRequestUserAccessParam {
	s.AccessKey = &v
	return s
}

func (s *CreateResourceRequestUserAccessParam) SetEndpoint(v string) *CreateResourceRequestUserAccessParam {
	s.Endpoint = &v
	return s
}

func (s *CreateResourceRequestUserAccessParam) SetWorkspaceId(v string) *CreateResourceRequestUserAccessParam {
	s.WorkspaceId = &v
	return s
}

type CreateResourceShrinkRequest struct {
	// Cluster Description
	//
	// example:
	//
	// ppu集群
	ClusterDesc *string `json:"ClusterDesc,omitempty" xml:"ClusterDesc,omitempty"`
	// Cluster ID
	//
	// This parameter is required.
	//
	// example:
	//
	// ehpc-sh-fj71c0ycfw
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// Cluster Name
	//
	// This parameter is required.
	//
	// example:
	//
	// tre-1-ppu
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// Machine Types
	//
	// This parameter is required.
	MachineTypesShrink *string `json:"MachineTypes,omitempty" xml:"MachineTypes,omitempty"`
	// User Access Parameters
	//
	// This parameter is required.
	UserAccessParamShrink *string `json:"UserAccessParam,omitempty" xml:"UserAccessParam,omitempty"`
}

func (s CreateResourceShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateResourceShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateResourceShrinkRequest) SetClusterDesc(v string) *CreateResourceShrinkRequest {
	s.ClusterDesc = &v
	return s
}

func (s *CreateResourceShrinkRequest) SetClusterId(v string) *CreateResourceShrinkRequest {
	s.ClusterId = &v
	return s
}

func (s *CreateResourceShrinkRequest) SetClusterName(v string) *CreateResourceShrinkRequest {
	s.ClusterName = &v
	return s
}

func (s *CreateResourceShrinkRequest) SetMachineTypesShrink(v string) *CreateResourceShrinkRequest {
	s.MachineTypesShrink = &v
	return s
}

func (s *CreateResourceShrinkRequest) SetUserAccessParamShrink(v string) *CreateResourceShrinkRequest {
	s.UserAccessParamShrink = &v
	return s
}

type CreateResourceResponseBody struct {
	// Data
	//
	// example:
	//
	// []
	Data *int64 `json:"Data,omitempty" xml:"Data,omitempty"`
	// Request ID
	//
	// example:
	//
	// 5514CB39-B7C0-5B89-8534-2DE1E0F2B7AB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Total Count
	//
	// example:
	//
	// 0
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s CreateResourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateResourceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateResourceResponseBody) SetData(v int64) *CreateResourceResponseBody {
	s.Data = &v
	return s
}

func (s *CreateResourceResponseBody) SetRequestId(v string) *CreateResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateResourceResponseBody) SetTotalCount(v int64) *CreateResourceResponseBody {
	s.TotalCount = &v
	return s
}

type CreateResourceResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateResourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateResourceResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateResourceResponse) GoString() string {
	return s.String()
}

func (s *CreateResourceResponse) SetHeaders(v map[string]*string) *CreateResourceResponse {
	s.Headers = v
	return s
}

func (s *CreateResourceResponse) SetStatusCode(v int32) *CreateResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateResourceResponse) SetBody(v *CreateResourceResponseBody) *CreateResourceResponse {
	s.Body = v
	return s
}

type DeleteExperimentRequest struct {
	// Plan ID
	//
	// This parameter is required.
	//
	// example:
	//
	// 234
	ExperimentId *int64 `json:"ExperimentId,omitempty" xml:"ExperimentId,omitempty"`
	// Resource Group Id
	//
	// example:
	//
	// rg-sdkfjgnvd24
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s DeleteExperimentRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteExperimentRequest) GoString() string {
	return s.String()
}

func (s *DeleteExperimentRequest) SetExperimentId(v int64) *DeleteExperimentRequest {
	s.ExperimentId = &v
	return s
}

func (s *DeleteExperimentRequest) SetResourceGroupId(v string) *DeleteExperimentRequest {
	s.ResourceGroupId = &v
	return s
}

type DeleteExperimentResponseBody struct {
	// Access denied information
	//
	// example:
	//
	// {}
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// Data
	//
	// example:
	//
	// []
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// Request ID
	//
	// example:
	//
	// 5514CB39-B7C0-5B89-8534-2DE1E0F2B7AB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Total count of queries
	//
	// example:
	//
	// 0
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DeleteExperimentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteExperimentResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteExperimentResponseBody) SetAccessDeniedDetail(v string) *DeleteExperimentResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *DeleteExperimentResponseBody) SetData(v bool) *DeleteExperimentResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteExperimentResponseBody) SetRequestId(v string) *DeleteExperimentResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteExperimentResponseBody) SetTotalCount(v int64) *DeleteExperimentResponseBody {
	s.TotalCount = &v
	return s
}

type DeleteExperimentResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteExperimentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteExperimentResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteExperimentResponse) GoString() string {
	return s.String()
}

func (s *DeleteExperimentResponse) SetHeaders(v map[string]*string) *DeleteExperimentResponse {
	s.Headers = v
	return s
}

func (s *DeleteExperimentResponse) SetStatusCode(v int32) *DeleteExperimentResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteExperimentResponse) SetBody(v *DeleteExperimentResponseBody) *DeleteExperimentResponse {
	s.Body = v
	return s
}

type DeleteExperimentPlanRequest struct {
	// Plan ID
	//
	// This parameter is required.
	//
	// example:
	//
	// 189
	PlanId *int64 `json:"PlanId,omitempty" xml:"PlanId,omitempty"`
}

func (s DeleteExperimentPlanRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteExperimentPlanRequest) GoString() string {
	return s.String()
}

func (s *DeleteExperimentPlanRequest) SetPlanId(v int64) *DeleteExperimentPlanRequest {
	s.PlanId = &v
	return s
}

type DeleteExperimentPlanResponseBody struct {
	// Access denied information
	//
	// example:
	//
	// {}
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// Data
	//
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// Request ID
	//
	// example:
	//
	// E67E2E4C-2B47-5C55-AA17-1D771E070AEF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Total
	//
	// example:
	//
	// 0
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DeleteExperimentPlanResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteExperimentPlanResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteExperimentPlanResponseBody) SetAccessDeniedDetail(v string) *DeleteExperimentPlanResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *DeleteExperimentPlanResponseBody) SetData(v bool) *DeleteExperimentPlanResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteExperimentPlanResponseBody) SetRequestId(v string) *DeleteExperimentPlanResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteExperimentPlanResponseBody) SetTotalCount(v int64) *DeleteExperimentPlanResponseBody {
	s.TotalCount = &v
	return s
}

type DeleteExperimentPlanResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteExperimentPlanResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteExperimentPlanResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteExperimentPlanResponse) GoString() string {
	return s.String()
}

func (s *DeleteExperimentPlanResponse) SetHeaders(v map[string]*string) *DeleteExperimentPlanResponse {
	s.Headers = v
	return s
}

func (s *DeleteExperimentPlanResponse) SetStatusCode(v int32) *DeleteExperimentPlanResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteExperimentPlanResponse) SetBody(v *DeleteExperimentPlanResponseBody) *DeleteExperimentPlanResponse {
	s.Body = v
	return s
}

type DeleteExperimentPlanTemplateRequest struct {
	// Template ID
	//
	// This parameter is required.
	//
	// example:
	//
	// 346527
	TemplateId *int64 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s DeleteExperimentPlanTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteExperimentPlanTemplateRequest) GoString() string {
	return s.String()
}

func (s *DeleteExperimentPlanTemplateRequest) SetTemplateId(v int64) *DeleteExperimentPlanTemplateRequest {
	s.TemplateId = &v
	return s
}

type DeleteExperimentPlanTemplateResponseBody struct {
	// Data
	//
	// example:
	//
	// []
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// Request ID
	//
	// example:
	//
	// 4D3FD55F-3BCD-5914-9B74-A1F4961327E7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Total Count
	//
	// example:
	//
	// 0
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DeleteExperimentPlanTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteExperimentPlanTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteExperimentPlanTemplateResponseBody) SetData(v bool) *DeleteExperimentPlanTemplateResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteExperimentPlanTemplateResponseBody) SetRequestId(v string) *DeleteExperimentPlanTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteExperimentPlanTemplateResponseBody) SetTotalCount(v int64) *DeleteExperimentPlanTemplateResponseBody {
	s.TotalCount = &v
	return s
}

type DeleteExperimentPlanTemplateResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteExperimentPlanTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteExperimentPlanTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteExperimentPlanTemplateResponse) GoString() string {
	return s.String()
}

func (s *DeleteExperimentPlanTemplateResponse) SetHeaders(v map[string]*string) *DeleteExperimentPlanTemplateResponse {
	s.Headers = v
	return s
}

func (s *DeleteExperimentPlanTemplateResponse) SetStatusCode(v int32) *DeleteExperimentPlanTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteExperimentPlanTemplateResponse) SetBody(v *DeleteExperimentPlanTemplateResponseBody) *DeleteExperimentPlanTemplateResponse {
	s.Body = v
	return s
}

type GetExperimentRequest struct {
	// Experiment ID
	//
	// This parameter is required.
	//
	// example:
	//
	// 234
	ExperimentId *int64 `json:"ExperimentId,omitempty" xml:"ExperimentId,omitempty"`
	// Resource Group Id
	//
	// example:
	//
	// rg-sdsmfg23
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s GetExperimentRequest) String() string {
	return tea.Prettify(s)
}

func (s GetExperimentRequest) GoString() string {
	return s.String()
}

func (s *GetExperimentRequest) SetExperimentId(v int64) *GetExperimentRequest {
	s.ExperimentId = &v
	return s
}

func (s *GetExperimentRequest) SetResourceGroupId(v string) *GetExperimentRequest {
	s.ResourceGroupId = &v
	return s
}

type GetExperimentResponseBody struct {
	// Access denied detail
	//
	// example:
	//
	// {}
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// Data
	Data *GetExperimentResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// RequestId
	//
	// example:
	//
	// E67E2E4C-2B47-5C55-AA17-1D771E070AEF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// total
	//
	// example:
	//
	// 0
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s GetExperimentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetExperimentResponseBody) GoString() string {
	return s.String()
}

func (s *GetExperimentResponseBody) SetAccessDeniedDetail(v string) *GetExperimentResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *GetExperimentResponseBody) SetData(v *GetExperimentResponseBodyData) *GetExperimentResponseBody {
	s.Data = v
	return s
}

func (s *GetExperimentResponseBody) SetRequestId(v string) *GetExperimentResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetExperimentResponseBody) SetTotalCount(v int64) *GetExperimentResponseBody {
	s.TotalCount = &v
	return s
}

type GetExperimentResponseBodyData struct {
	// Creation time
	//
	// example:
	//
	// 2024-11-29 02:16:35
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// Task end time
	//
	// example:
	//
	// 2024-11-29 02:26:35
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// Running environment parameters
	EnvParams *GetExperimentResponseBodyDataEnvParams `json:"EnvParams,omitempty" xml:"EnvParams,omitempty" type:"Struct"`
	// Experiment ID
	//
	// example:
	//
	// 1726882991828688898
	ExperimentId *int64 `json:"ExperimentId,omitempty" xml:"ExperimentId,omitempty"`
	// Experiment name
	//
	// example:
	//
	// test
	ExperimentName *string `json:"ExperimentName,omitempty" xml:"ExperimentName,omitempty"`
	// Experiment type
	//
	// example:
	//
	// AI
	ExperimentType *string `json:"ExperimentType,omitempty" xml:"ExperimentType,omitempty"`
	// Parsed workload parameters
	GetParams map[string]*string `json:"GetParams,omitempty" xml:"GetParams,omitempty"`
	// cluster info
	Resource *GetExperimentResponseBodyDataResource `json:"Resource,omitempty" xml:"Resource,omitempty" type:"Struct"`
	// Resource name
	//
	// example:
	//
	// cifnews-guoyuan
	ResourceName *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	// Task results
	Results *GetExperimentResponseBodyDataResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Struct"`
	// Running workload parameters
	SetParams map[string]*string `json:"SetParams,omitempty" xml:"SetParams,omitempty"`
	// Task start time
	//
	// example:
	//
	// 2024-11-29 02:16:35
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// Status
	//
	// example:
	//
	// RUNNING
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Experiment task
	Task *GetExperimentResponseBodyDataTask `json:"Task,omitempty" xml:"Task,omitempty" type:"Struct"`
	// Update time
	//
	// example:
	//
	// 2024-11-29 02:16:35
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// Workload information
	Workload *GetExperimentResponseBodyDataWorkload `json:"Workload,omitempty" xml:"Workload,omitempty" type:"Struct"`
	// Workload name
	//
	// example:
	//
	// test
	WorkloadName *string `json:"WorkloadName,omitempty" xml:"WorkloadName,omitempty"`
}

func (s GetExperimentResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetExperimentResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetExperimentResponseBodyData) SetCreateTime(v int64) *GetExperimentResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *GetExperimentResponseBodyData) SetEndTime(v string) *GetExperimentResponseBodyData {
	s.EndTime = &v
	return s
}

func (s *GetExperimentResponseBodyData) SetEnvParams(v *GetExperimentResponseBodyDataEnvParams) *GetExperimentResponseBodyData {
	s.EnvParams = v
	return s
}

func (s *GetExperimentResponseBodyData) SetExperimentId(v int64) *GetExperimentResponseBodyData {
	s.ExperimentId = &v
	return s
}

func (s *GetExperimentResponseBodyData) SetExperimentName(v string) *GetExperimentResponseBodyData {
	s.ExperimentName = &v
	return s
}

func (s *GetExperimentResponseBodyData) SetExperimentType(v string) *GetExperimentResponseBodyData {
	s.ExperimentType = &v
	return s
}

func (s *GetExperimentResponseBodyData) SetGetParams(v map[string]*string) *GetExperimentResponseBodyData {
	s.GetParams = v
	return s
}

func (s *GetExperimentResponseBodyData) SetResource(v *GetExperimentResponseBodyDataResource) *GetExperimentResponseBodyData {
	s.Resource = v
	return s
}

func (s *GetExperimentResponseBodyData) SetResourceName(v string) *GetExperimentResponseBodyData {
	s.ResourceName = &v
	return s
}

func (s *GetExperimentResponseBodyData) SetResults(v *GetExperimentResponseBodyDataResults) *GetExperimentResponseBodyData {
	s.Results = v
	return s
}

func (s *GetExperimentResponseBodyData) SetSetParams(v map[string]*string) *GetExperimentResponseBodyData {
	s.SetParams = v
	return s
}

func (s *GetExperimentResponseBodyData) SetStartTime(v string) *GetExperimentResponseBodyData {
	s.StartTime = &v
	return s
}

func (s *GetExperimentResponseBodyData) SetStatus(v string) *GetExperimentResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetExperimentResponseBodyData) SetTask(v *GetExperimentResponseBodyDataTask) *GetExperimentResponseBodyData {
	s.Task = v
	return s
}

func (s *GetExperimentResponseBodyData) SetUpdateTime(v int64) *GetExperimentResponseBodyData {
	s.UpdateTime = &v
	return s
}

func (s *GetExperimentResponseBodyData) SetWorkload(v *GetExperimentResponseBodyDataWorkload) *GetExperimentResponseBodyData {
	s.Workload = v
	return s
}

func (s *GetExperimentResponseBodyData) SetWorkloadName(v string) *GetExperimentResponseBodyData {
	s.WorkloadName = &v
	return s
}

type GetExperimentResponseBodyDataEnvParams struct {
	// CPU allocation number
	//
	// example:
	//
	// 90
	CpuPerWorker *int32 `json:"CpuPerWorker,omitempty" xml:"CpuPerWorker,omitempty"`
	// cudaVersion
	//
	// example:
	//
	// 1.0.0
	CudaVersion *string `json:"CudaVersion,omitempty" xml:"CudaVersion,omitempty"`
	// Additional parameters
	ExtendParam map[string]*string `json:"ExtendParam,omitempty" xml:"ExtendParam,omitempty"`
	// GPU driver version
	//
	// example:
	//
	// 1.0.0
	GpuDriverVersion *string `json:"GpuDriverVersion,omitempty" xml:"GpuDriverVersion,omitempty"`
	// GPU allocation number
	//
	// example:
	//
	// 8
	GpuPerWorker *int32 `json:"GpuPerWorker,omitempty" xml:"GpuPerWorker,omitempty"`
	// Memory Per Worker
	//
	// example:
	//
	// 500
	MemoryPerWorker *int32 `json:"MemoryPerWorker,omitempty" xml:"MemoryPerWorker,omitempty"`
	// NCCL version
	//
	// example:
	//
	// 1.0.0
	NCCLVersion *string `json:"NCCLVersion,omitempty" xml:"NCCLVersion,omitempty"`
	// PyTorch version
	//
	// example:
	//
	// 1.0.0
	PyTorchVersion *string `json:"PyTorchVersion,omitempty" xml:"PyTorchVersion,omitempty"`
	// Specified nodes
	ResourceNodes []*GetExperimentResponseBodyDataEnvParamsResourceNodes `json:"ResourceNodes,omitempty" xml:"ResourceNodes,omitempty" type:"Repeated"`
	// Share Memory
	//
	// example:
	//
	// 500
	ShareMemory *int32 `json:"ShareMemory,omitempty" xml:"ShareMemory,omitempty"`
	// Worker number
	//
	// example:
	//
	// 1
	WorkerNum *int32 `json:"WorkerNum,omitempty" xml:"WorkerNum,omitempty"`
}

func (s GetExperimentResponseBodyDataEnvParams) String() string {
	return tea.Prettify(s)
}

func (s GetExperimentResponseBodyDataEnvParams) GoString() string {
	return s.String()
}

func (s *GetExperimentResponseBodyDataEnvParams) SetCpuPerWorker(v int32) *GetExperimentResponseBodyDataEnvParams {
	s.CpuPerWorker = &v
	return s
}

func (s *GetExperimentResponseBodyDataEnvParams) SetCudaVersion(v string) *GetExperimentResponseBodyDataEnvParams {
	s.CudaVersion = &v
	return s
}

func (s *GetExperimentResponseBodyDataEnvParams) SetExtendParam(v map[string]*string) *GetExperimentResponseBodyDataEnvParams {
	s.ExtendParam = v
	return s
}

func (s *GetExperimentResponseBodyDataEnvParams) SetGpuDriverVersion(v string) *GetExperimentResponseBodyDataEnvParams {
	s.GpuDriverVersion = &v
	return s
}

func (s *GetExperimentResponseBodyDataEnvParams) SetGpuPerWorker(v int32) *GetExperimentResponseBodyDataEnvParams {
	s.GpuPerWorker = &v
	return s
}

func (s *GetExperimentResponseBodyDataEnvParams) SetMemoryPerWorker(v int32) *GetExperimentResponseBodyDataEnvParams {
	s.MemoryPerWorker = &v
	return s
}

func (s *GetExperimentResponseBodyDataEnvParams) SetNCCLVersion(v string) *GetExperimentResponseBodyDataEnvParams {
	s.NCCLVersion = &v
	return s
}

func (s *GetExperimentResponseBodyDataEnvParams) SetPyTorchVersion(v string) *GetExperimentResponseBodyDataEnvParams {
	s.PyTorchVersion = &v
	return s
}

func (s *GetExperimentResponseBodyDataEnvParams) SetResourceNodes(v []*GetExperimentResponseBodyDataEnvParamsResourceNodes) *GetExperimentResponseBodyDataEnvParams {
	s.ResourceNodes = v
	return s
}

func (s *GetExperimentResponseBodyDataEnvParams) SetShareMemory(v int32) *GetExperimentResponseBodyDataEnvParams {
	s.ShareMemory = &v
	return s
}

func (s *GetExperimentResponseBodyDataEnvParams) SetWorkerNum(v int32) *GetExperimentResponseBodyDataEnvParams {
	s.WorkerNum = &v
	return s
}

type GetExperimentResponseBodyDataEnvParamsResourceNodes struct {
	// Node name
	//
	// example:
	//
	// p-jt-waf-app1
	NodeName *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	// Requested CPU
	//
	// example:
	//
	// 90
	RequestCPU *int32 `json:"RequestCPU,omitempty" xml:"RequestCPU,omitempty"`
	// Requested GPU
	//
	// example:
	//
	// 8
	RequestGPU *int32 `json:"RequestGPU,omitempty" xml:"RequestGPU,omitempty"`
	// Requested memory
	//
	// example:
	//
	// 500
	RequestMemory *int32 `json:"RequestMemory,omitempty" xml:"RequestMemory,omitempty"`
	// Total CPU
	//
	// example:
	//
	// 90
	TotalCPU *int32 `json:"TotalCPU,omitempty" xml:"TotalCPU,omitempty"`
	// Total GPU
	//
	// example:
	//
	// 8
	TotalGPU *int32 `json:"TotalGPU,omitempty" xml:"TotalGPU,omitempty"`
	// Total memory
	//
	// example:
	//
	// 500
	TotalMemory *int64 `json:"TotalMemory,omitempty" xml:"TotalMemory,omitempty"`
}

func (s GetExperimentResponseBodyDataEnvParamsResourceNodes) String() string {
	return tea.Prettify(s)
}

func (s GetExperimentResponseBodyDataEnvParamsResourceNodes) GoString() string {
	return s.String()
}

func (s *GetExperimentResponseBodyDataEnvParamsResourceNodes) SetNodeName(v string) *GetExperimentResponseBodyDataEnvParamsResourceNodes {
	s.NodeName = &v
	return s
}

func (s *GetExperimentResponseBodyDataEnvParamsResourceNodes) SetRequestCPU(v int32) *GetExperimentResponseBodyDataEnvParamsResourceNodes {
	s.RequestCPU = &v
	return s
}

func (s *GetExperimentResponseBodyDataEnvParamsResourceNodes) SetRequestGPU(v int32) *GetExperimentResponseBodyDataEnvParamsResourceNodes {
	s.RequestGPU = &v
	return s
}

func (s *GetExperimentResponseBodyDataEnvParamsResourceNodes) SetRequestMemory(v int32) *GetExperimentResponseBodyDataEnvParamsResourceNodes {
	s.RequestMemory = &v
	return s
}

func (s *GetExperimentResponseBodyDataEnvParamsResourceNodes) SetTotalCPU(v int32) *GetExperimentResponseBodyDataEnvParamsResourceNodes {
	s.TotalCPU = &v
	return s
}

func (s *GetExperimentResponseBodyDataEnvParamsResourceNodes) SetTotalGPU(v int32) *GetExperimentResponseBodyDataEnvParamsResourceNodes {
	s.TotalGPU = &v
	return s
}

func (s *GetExperimentResponseBodyDataEnvParamsResourceNodes) SetTotalMemory(v int64) *GetExperimentResponseBodyDataEnvParamsResourceNodes {
	s.TotalMemory = &v
	return s
}

type GetExperimentResponseBodyDataResource struct {
	// Used CPU
	//
	// example:
	//
	// 90
	CpuCoreLimit *int32 `json:"CpuCoreLimit,omitempty" xml:"CpuCoreLimit,omitempty"`
	// Used GPU
	//
	// example:
	//
	// 8
	GpuLimit *int32 `json:"GpuLimit,omitempty" xml:"GpuLimit,omitempty"`
	// Instance type
	MachineType *GetExperimentResponseBodyDataResourceMachineType `json:"MachineType,omitempty" xml:"MachineType,omitempty" type:"Struct"`
	// Used memory
	//
	// example:
	//
	// 90
	MaxCpuCore *int32 `json:"MaxCpuCore,omitempty" xml:"MaxCpuCore,omitempty"`
	// Used memory
	//
	// example:
	//
	// 8
	MaxGpu *int32 `json:"MaxGpu,omitempty" xml:"MaxGpu,omitempty"`
	// Used memory
	//
	// example:
	//
	// 500
	MaxMemory *int64 `json:"MaxMemory,omitempty" xml:"MaxMemory,omitempty"`
	// Used memory
	//
	// example:
	//
	// 500
	MemoryLimit *int64 `json:"MemoryLimit,omitempty" xml:"MemoryLimit,omitempty"`
	// Cluster ID
	//
	// example:
	//
	// 189
	ResourceId *int64 `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// Cluster name
	//
	// example:
	//
	// ecs.r8y.4xlarge
	ResourceName *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	// Resource node list
	ResourceNodes []*GetExperimentResponseBodyDataResourceResourceNodes `json:"ResourceNodes,omitempty" xml:"ResourceNodes,omitempty" type:"Repeated"`
	// User authorization parameters
	UserAccessParam *GetExperimentResponseBodyDataResourceUserAccessParam `json:"UserAccessParam,omitempty" xml:"UserAccessParam,omitempty" type:"Struct"`
}

func (s GetExperimentResponseBodyDataResource) String() string {
	return tea.Prettify(s)
}

func (s GetExperimentResponseBodyDataResource) GoString() string {
	return s.String()
}

func (s *GetExperimentResponseBodyDataResource) SetCpuCoreLimit(v int32) *GetExperimentResponseBodyDataResource {
	s.CpuCoreLimit = &v
	return s
}

func (s *GetExperimentResponseBodyDataResource) SetGpuLimit(v int32) *GetExperimentResponseBodyDataResource {
	s.GpuLimit = &v
	return s
}

func (s *GetExperimentResponseBodyDataResource) SetMachineType(v *GetExperimentResponseBodyDataResourceMachineType) *GetExperimentResponseBodyDataResource {
	s.MachineType = v
	return s
}

func (s *GetExperimentResponseBodyDataResource) SetMaxCpuCore(v int32) *GetExperimentResponseBodyDataResource {
	s.MaxCpuCore = &v
	return s
}

func (s *GetExperimentResponseBodyDataResource) SetMaxGpu(v int32) *GetExperimentResponseBodyDataResource {
	s.MaxGpu = &v
	return s
}

func (s *GetExperimentResponseBodyDataResource) SetMaxMemory(v int64) *GetExperimentResponseBodyDataResource {
	s.MaxMemory = &v
	return s
}

func (s *GetExperimentResponseBodyDataResource) SetMemoryLimit(v int64) *GetExperimentResponseBodyDataResource {
	s.MemoryLimit = &v
	return s
}

func (s *GetExperimentResponseBodyDataResource) SetResourceId(v int64) *GetExperimentResponseBodyDataResource {
	s.ResourceId = &v
	return s
}

func (s *GetExperimentResponseBodyDataResource) SetResourceName(v string) *GetExperimentResponseBodyDataResource {
	s.ResourceName = &v
	return s
}

func (s *GetExperimentResponseBodyDataResource) SetResourceNodes(v []*GetExperimentResponseBodyDataResourceResourceNodes) *GetExperimentResponseBodyDataResource {
	s.ResourceNodes = v
	return s
}

func (s *GetExperimentResponseBodyDataResource) SetUserAccessParam(v *GetExperimentResponseBodyDataResourceUserAccessParam) *GetExperimentResponseBodyDataResource {
	s.UserAccessParam = v
	return s
}

type GetExperimentResponseBodyDataResourceMachineType struct {
	// Number of network bonds
	//
	// example:
	//
	// 5
	BondNum *int32 `json:"BondNum,omitempty" xml:"BondNum,omitempty"`
	// CPU information
	//
	// example:
	//
	// 2x Intel Icelake 8369B 32C CPU
	CpuInfo *string `json:"CpuInfo,omitempty" xml:"CpuInfo,omitempty"`
	// Disk information
	//
	// example:
	//
	// 2x 480GB SATA SSD \\n 4x 3.84TB NVMe SSD
	DiskInfo *string `json:"DiskInfo,omitempty" xml:"DiskInfo,omitempty"`
	// GPU information
	//
	// example:
	//
	// 8x NVIDIA SXM4 80GB A100 GPU
	GpuInfo *string `json:"GpuInfo,omitempty" xml:"GpuInfo,omitempty"`
	// Memory information
	//
	// example:
	//
	// 32x 64GB DDR4 3200 Memory
	MemoryInfo *string `json:"MemoryInfo,omitempty" xml:"MemoryInfo,omitempty"`
	// Specification name
	//
	// example:
	//
	// efg1.nvga1n
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Network information
	//
	// example:
	//
	// 1x 100Gbps DP NIC for VPC \\n 4x 100Gbps DP RoCE NIC
	NetworkInfo *string `json:"NetworkInfo,omitempty" xml:"NetworkInfo,omitempty"`
	// Network mode
	//
	// example:
	//
	// 2
	NetworkMode *string `json:"NetworkMode,omitempty" xml:"NetworkMode,omitempty"`
	// Number of nodes
	//
	// example:
	//
	// 1
	NodeCount *int32 `json:"NodeCount,omitempty" xml:"NodeCount,omitempty"`
	// Type
	//
	// example:
	//
	// Public
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetExperimentResponseBodyDataResourceMachineType) String() string {
	return tea.Prettify(s)
}

func (s GetExperimentResponseBodyDataResourceMachineType) GoString() string {
	return s.String()
}

func (s *GetExperimentResponseBodyDataResourceMachineType) SetBondNum(v int32) *GetExperimentResponseBodyDataResourceMachineType {
	s.BondNum = &v
	return s
}

func (s *GetExperimentResponseBodyDataResourceMachineType) SetCpuInfo(v string) *GetExperimentResponseBodyDataResourceMachineType {
	s.CpuInfo = &v
	return s
}

func (s *GetExperimentResponseBodyDataResourceMachineType) SetDiskInfo(v string) *GetExperimentResponseBodyDataResourceMachineType {
	s.DiskInfo = &v
	return s
}

func (s *GetExperimentResponseBodyDataResourceMachineType) SetGpuInfo(v string) *GetExperimentResponseBodyDataResourceMachineType {
	s.GpuInfo = &v
	return s
}

func (s *GetExperimentResponseBodyDataResourceMachineType) SetMemoryInfo(v string) *GetExperimentResponseBodyDataResourceMachineType {
	s.MemoryInfo = &v
	return s
}

func (s *GetExperimentResponseBodyDataResourceMachineType) SetName(v string) *GetExperimentResponseBodyDataResourceMachineType {
	s.Name = &v
	return s
}

func (s *GetExperimentResponseBodyDataResourceMachineType) SetNetworkInfo(v string) *GetExperimentResponseBodyDataResourceMachineType {
	s.NetworkInfo = &v
	return s
}

func (s *GetExperimentResponseBodyDataResourceMachineType) SetNetworkMode(v string) *GetExperimentResponseBodyDataResourceMachineType {
	s.NetworkMode = &v
	return s
}

func (s *GetExperimentResponseBodyDataResourceMachineType) SetNodeCount(v int32) *GetExperimentResponseBodyDataResourceMachineType {
	s.NodeCount = &v
	return s
}

func (s *GetExperimentResponseBodyDataResourceMachineType) SetType(v string) *GetExperimentResponseBodyDataResourceMachineType {
	s.Type = &v
	return s
}

type GetExperimentResponseBodyDataResourceResourceNodes struct {
	// Node name
	//
	// example:
	//
	// InputCheck
	NodeName *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
}

func (s GetExperimentResponseBodyDataResourceResourceNodes) String() string {
	return tea.Prettify(s)
}

func (s GetExperimentResponseBodyDataResourceResourceNodes) GoString() string {
	return s.String()
}

func (s *GetExperimentResponseBodyDataResourceResourceNodes) SetNodeName(v string) *GetExperimentResponseBodyDataResourceResourceNodes {
	s.NodeName = &v
	return s
}

type GetExperimentResponseBodyDataResourceUserAccessParam struct {
	// User ID
	//
	// example:
	//
	// dev
	AccessId *string `json:"AccessId,omitempty" xml:"AccessId,omitempty"`
	// User key
	//
	// example:
	//
	// test
	AccessKey *string `json:"AccessKey,omitempty" xml:"AccessKey,omitempty"`
	// Endpoint
	//
	// example:
	//
	// test
	Endpoint *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	// Workspace ID
	//
	// example:
	//
	// 123434542498
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s GetExperimentResponseBodyDataResourceUserAccessParam) String() string {
	return tea.Prettify(s)
}

func (s GetExperimentResponseBodyDataResourceUserAccessParam) GoString() string {
	return s.String()
}

func (s *GetExperimentResponseBodyDataResourceUserAccessParam) SetAccessId(v string) *GetExperimentResponseBodyDataResourceUserAccessParam {
	s.AccessId = &v
	return s
}

func (s *GetExperimentResponseBodyDataResourceUserAccessParam) SetAccessKey(v string) *GetExperimentResponseBodyDataResourceUserAccessParam {
	s.AccessKey = &v
	return s
}

func (s *GetExperimentResponseBodyDataResourceUserAccessParam) SetEndpoint(v string) *GetExperimentResponseBodyDataResourceUserAccessParam {
	s.Endpoint = &v
	return s
}

func (s *GetExperimentResponseBodyDataResourceUserAccessParam) SetWorkspaceId(v string) *GetExperimentResponseBodyDataResourceUserAccessParam {
	s.WorkspaceId = &v
	return s
}

type GetExperimentResponseBodyDataResults struct {
	// Duration
	//
	// example:
	//
	// 764
	Duration *float64 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// Error node
	ErrorWorker []*GetExperimentResponseBodyDataResultsErrorWorker `json:"ErrorWorker,omitempty" xml:"ErrorWorker,omitempty" type:"Repeated"`
	// Parameter name
	//
	// example:
	//
	// 1748274952976261121
	ExperimentId *int64 `json:"ExperimentId,omitempty" xml:"ExperimentId,omitempty"`
	// MFU
	//
	// example:
	//
	// 54.2
	Mfu *float64 `json:"Mfu,omitempty" xml:"Mfu,omitempty"`
	// Samples Per Second
	//
	// example:
	//
	// 10
	SamplesPerSecond *float64 `json:"SamplesPerSecond,omitempty" xml:"SamplesPerSecond,omitempty"`
	// Seconds per iteration
	//
	// example:
	//
	// 1000
	SecondsPerIteration *float64 `json:"SecondsPerIteration,omitempty" xml:"SecondsPerIteration,omitempty"`
	// Task individual result list
	TaskIndividualResultList []*GetExperimentResponseBodyDataResultsTaskIndividualResultList `json:"TaskIndividualResultList,omitempty" xml:"TaskIndividualResultList,omitempty" type:"Repeated"`
	// Invalid task results
	TaskIndividualResultMap map[string][]*DataResultsTaskIndividualResultMapValue `json:"TaskIndividualResultMap,omitempty" xml:"TaskIndividualResultMap,omitempty"`
	// Warning bound list
	WarningBoundList []*GetExperimentResponseBodyDataResultsWarningBoundList `json:"WarningBoundList,omitempty" xml:"WarningBoundList,omitempty" type:"Repeated"`
	// Warning worker
	WarningWorker []*GetExperimentResponseBodyDataResultsWarningWorker `json:"WarningWorker,omitempty" xml:"WarningWorker,omitempty" type:"Repeated"`
}

func (s GetExperimentResponseBodyDataResults) String() string {
	return tea.Prettify(s)
}

func (s GetExperimentResponseBodyDataResults) GoString() string {
	return s.String()
}

func (s *GetExperimentResponseBodyDataResults) SetDuration(v float64) *GetExperimentResponseBodyDataResults {
	s.Duration = &v
	return s
}

func (s *GetExperimentResponseBodyDataResults) SetErrorWorker(v []*GetExperimentResponseBodyDataResultsErrorWorker) *GetExperimentResponseBodyDataResults {
	s.ErrorWorker = v
	return s
}

func (s *GetExperimentResponseBodyDataResults) SetExperimentId(v int64) *GetExperimentResponseBodyDataResults {
	s.ExperimentId = &v
	return s
}

func (s *GetExperimentResponseBodyDataResults) SetMfu(v float64) *GetExperimentResponseBodyDataResults {
	s.Mfu = &v
	return s
}

func (s *GetExperimentResponseBodyDataResults) SetSamplesPerSecond(v float64) *GetExperimentResponseBodyDataResults {
	s.SamplesPerSecond = &v
	return s
}

func (s *GetExperimentResponseBodyDataResults) SetSecondsPerIteration(v float64) *GetExperimentResponseBodyDataResults {
	s.SecondsPerIteration = &v
	return s
}

func (s *GetExperimentResponseBodyDataResults) SetTaskIndividualResultList(v []*GetExperimentResponseBodyDataResultsTaskIndividualResultList) *GetExperimentResponseBodyDataResults {
	s.TaskIndividualResultList = v
	return s
}

func (s *GetExperimentResponseBodyDataResults) SetTaskIndividualResultMap(v map[string][]*DataResultsTaskIndividualResultMapValue) *GetExperimentResponseBodyDataResults {
	s.TaskIndividualResultMap = v
	return s
}

func (s *GetExperimentResponseBodyDataResults) SetWarningBoundList(v []*GetExperimentResponseBodyDataResultsWarningBoundList) *GetExperimentResponseBodyDataResults {
	s.WarningBoundList = v
	return s
}

func (s *GetExperimentResponseBodyDataResults) SetWarningWorker(v []*GetExperimentResponseBodyDataResultsWarningWorker) *GetExperimentResponseBodyDataResults {
	s.WarningWorker = v
	return s
}

type GetExperimentResponseBodyDataResultsErrorWorker struct {
	// error flag
	//
	// example:
	//
	// true
	ErrorFlag *bool `json:"ErrorFlag,omitempty" xml:"ErrorFlag,omitempty"`
	// error message
	//
	// example:
	//
	// Connection reset
	ErrorMsg *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	// Experiment ID
	//
	// example:
	//
	// 97
	ExperimentId *int64 `json:"ExperimentId,omitempty" xml:"ExperimentId,omitempty"`
	// GPU name
	//
	// example:
	//
	// 8x OAM 810 GPU
	GpuName *string `json:"GpuName,omitempty" xml:"GpuName,omitempty"`
	// Number of GPUs
	//
	// example:
	//
	// 8
	GpuNum *int32 `json:"GpuNum,omitempty" xml:"GpuNum,omitempty"`
	// Service address
	//
	// example:
	//
	// 60.188.98.209
	Hostname *string `json:"Hostname,omitempty" xml:"Hostname,omitempty"`
	// Pod name.
	//
	// example:
	//
	// hzs-forge-sdxl-online-7ff4d86444-pc95h
	PodName *string `json:"PodName,omitempty" xml:"PodName,omitempty"`
	// Samples Per Second
	//
	// example:
	//
	// 23
	SamplesPerSecond *float64 `json:"SamplesPerSecond,omitempty" xml:"SamplesPerSecond,omitempty"`
	// TFLOPS
	//
	// example:
	//
	// 12
	Tflops *float64 `json:"Tflops,omitempty" xml:"Tflops,omitempty"`
	// Whether there is a warning
	//
	// example:
	//
	// false
	WarningFlag *bool `json:"WarningFlag,omitempty" xml:"WarningFlag,omitempty"`
	// Warning message
	//
	// example:
	//
	// warning message
	WarningMsg *string `json:"WarningMsg,omitempty" xml:"WarningMsg,omitempty"`
}

func (s GetExperimentResponseBodyDataResultsErrorWorker) String() string {
	return tea.Prettify(s)
}

func (s GetExperimentResponseBodyDataResultsErrorWorker) GoString() string {
	return s.String()
}

func (s *GetExperimentResponseBodyDataResultsErrorWorker) SetErrorFlag(v bool) *GetExperimentResponseBodyDataResultsErrorWorker {
	s.ErrorFlag = &v
	return s
}

func (s *GetExperimentResponseBodyDataResultsErrorWorker) SetErrorMsg(v string) *GetExperimentResponseBodyDataResultsErrorWorker {
	s.ErrorMsg = &v
	return s
}

func (s *GetExperimentResponseBodyDataResultsErrorWorker) SetExperimentId(v int64) *GetExperimentResponseBodyDataResultsErrorWorker {
	s.ExperimentId = &v
	return s
}

func (s *GetExperimentResponseBodyDataResultsErrorWorker) SetGpuName(v string) *GetExperimentResponseBodyDataResultsErrorWorker {
	s.GpuName = &v
	return s
}

func (s *GetExperimentResponseBodyDataResultsErrorWorker) SetGpuNum(v int32) *GetExperimentResponseBodyDataResultsErrorWorker {
	s.GpuNum = &v
	return s
}

func (s *GetExperimentResponseBodyDataResultsErrorWorker) SetHostname(v string) *GetExperimentResponseBodyDataResultsErrorWorker {
	s.Hostname = &v
	return s
}

func (s *GetExperimentResponseBodyDataResultsErrorWorker) SetPodName(v string) *GetExperimentResponseBodyDataResultsErrorWorker {
	s.PodName = &v
	return s
}

func (s *GetExperimentResponseBodyDataResultsErrorWorker) SetSamplesPerSecond(v float64) *GetExperimentResponseBodyDataResultsErrorWorker {
	s.SamplesPerSecond = &v
	return s
}

func (s *GetExperimentResponseBodyDataResultsErrorWorker) SetTflops(v float64) *GetExperimentResponseBodyDataResultsErrorWorker {
	s.Tflops = &v
	return s
}

func (s *GetExperimentResponseBodyDataResultsErrorWorker) SetWarningFlag(v bool) *GetExperimentResponseBodyDataResultsErrorWorker {
	s.WarningFlag = &v
	return s
}

func (s *GetExperimentResponseBodyDataResultsErrorWorker) SetWarningMsg(v string) *GetExperimentResponseBodyDataResultsErrorWorker {
	s.WarningMsg = &v
	return s
}

type GetExperimentResponseBodyDataResultsTaskIndividualResultList struct {
	// Whether there is an error
	//
	// example:
	//
	// false
	ErrorFlag *bool `json:"ErrorFlag,omitempty" xml:"ErrorFlag,omitempty"`
	// Error message
	//
	// example:
	//
	// error message
	ErrorMsg *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	// 实验ID。
	//
	// example:
	//
	// 48
	ExperimentId *int64 `json:"ExperimentId,omitempty" xml:"ExperimentId,omitempty"`
	// GPU name
	//
	// example:
	//
	// 8x OAM 810 GPU
	GpuName *string `json:"GpuName,omitempty" xml:"GpuName,omitempty"`
	// Number of GPUs
	//
	// example:
	//
	// 8
	GpuNum *int32 `json:"GpuNum,omitempty" xml:"GpuNum,omitempty"`
	// 节点主机名称。
	//
	// example:
	//
	// p-jt-waf-app1
	Hostname *string `json:"Hostname,omitempty" xml:"Hostname,omitempty"`
	// Pod名称。
	//
	// example:
	//
	// fluxserv-6fc89b45cf-w8wq6
	PodName *string `json:"PodName,omitempty" xml:"PodName,omitempty"`
	// Throughput
	//
	// example:
	//
	// 28
	SamplesPerSecond *float64 `json:"SamplesPerSecond,omitempty" xml:"SamplesPerSecond,omitempty"`
	// TFLOPS value
	//
	// example:
	//
	// 16
	Tflops *float64 `json:"Tflops,omitempty" xml:"Tflops,omitempty"`
	// Whether there is a warning
	//
	// example:
	//
	// false
	WarningFlag *bool `json:"WarningFlag,omitempty" xml:"WarningFlag,omitempty"`
	// Warning message
	//
	// example:
	//
	// warning message
	WarningMsg *string `json:"WarningMsg,omitempty" xml:"WarningMsg,omitempty"`
}

func (s GetExperimentResponseBodyDataResultsTaskIndividualResultList) String() string {
	return tea.Prettify(s)
}

func (s GetExperimentResponseBodyDataResultsTaskIndividualResultList) GoString() string {
	return s.String()
}

func (s *GetExperimentResponseBodyDataResultsTaskIndividualResultList) SetErrorFlag(v bool) *GetExperimentResponseBodyDataResultsTaskIndividualResultList {
	s.ErrorFlag = &v
	return s
}

func (s *GetExperimentResponseBodyDataResultsTaskIndividualResultList) SetErrorMsg(v string) *GetExperimentResponseBodyDataResultsTaskIndividualResultList {
	s.ErrorMsg = &v
	return s
}

func (s *GetExperimentResponseBodyDataResultsTaskIndividualResultList) SetExperimentId(v int64) *GetExperimentResponseBodyDataResultsTaskIndividualResultList {
	s.ExperimentId = &v
	return s
}

func (s *GetExperimentResponseBodyDataResultsTaskIndividualResultList) SetGpuName(v string) *GetExperimentResponseBodyDataResultsTaskIndividualResultList {
	s.GpuName = &v
	return s
}

func (s *GetExperimentResponseBodyDataResultsTaskIndividualResultList) SetGpuNum(v int32) *GetExperimentResponseBodyDataResultsTaskIndividualResultList {
	s.GpuNum = &v
	return s
}

func (s *GetExperimentResponseBodyDataResultsTaskIndividualResultList) SetHostname(v string) *GetExperimentResponseBodyDataResultsTaskIndividualResultList {
	s.Hostname = &v
	return s
}

func (s *GetExperimentResponseBodyDataResultsTaskIndividualResultList) SetPodName(v string) *GetExperimentResponseBodyDataResultsTaskIndividualResultList {
	s.PodName = &v
	return s
}

func (s *GetExperimentResponseBodyDataResultsTaskIndividualResultList) SetSamplesPerSecond(v float64) *GetExperimentResponseBodyDataResultsTaskIndividualResultList {
	s.SamplesPerSecond = &v
	return s
}

func (s *GetExperimentResponseBodyDataResultsTaskIndividualResultList) SetTflops(v float64) *GetExperimentResponseBodyDataResultsTaskIndividualResultList {
	s.Tflops = &v
	return s
}

func (s *GetExperimentResponseBodyDataResultsTaskIndividualResultList) SetWarningFlag(v bool) *GetExperimentResponseBodyDataResultsTaskIndividualResultList {
	s.WarningFlag = &v
	return s
}

func (s *GetExperimentResponseBodyDataResultsTaskIndividualResultList) SetWarningMsg(v string) *GetExperimentResponseBodyDataResultsTaskIndividualResultList {
	s.WarningMsg = &v
	return s
}

type GetExperimentResponseBodyDataResultsWarningBoundList struct {
	// Iteration
	//
	// example:
	//
	// 10
	Iteration *int32 `json:"Iteration,omitempty" xml:"Iteration,omitempty"`
	// LOWER
	//
	// example:
	//
	// 14
	Lower *float64 `json:"Lower,omitempty" xml:"Lower,omitempty"`
	// UPPER
	//
	// example:
	//
	// 56
	Upper *float64 `json:"Upper,omitempty" xml:"Upper,omitempty"`
}

func (s GetExperimentResponseBodyDataResultsWarningBoundList) String() string {
	return tea.Prettify(s)
}

func (s GetExperimentResponseBodyDataResultsWarningBoundList) GoString() string {
	return s.String()
}

func (s *GetExperimentResponseBodyDataResultsWarningBoundList) SetIteration(v int32) *GetExperimentResponseBodyDataResultsWarningBoundList {
	s.Iteration = &v
	return s
}

func (s *GetExperimentResponseBodyDataResultsWarningBoundList) SetLower(v float64) *GetExperimentResponseBodyDataResultsWarningBoundList {
	s.Lower = &v
	return s
}

func (s *GetExperimentResponseBodyDataResultsWarningBoundList) SetUpper(v float64) *GetExperimentResponseBodyDataResultsWarningBoundList {
	s.Upper = &v
	return s
}

type GetExperimentResponseBodyDataResultsWarningWorker struct {
	// Whether there is an error
	//
	// example:
	//
	// true
	ErrorFlag *bool `json:"ErrorFlag,omitempty" xml:"ErrorFlag,omitempty"`
	// Error message
	//
	// example:
	//
	// error message
	ErrorMsg *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	// Experiment ID
	//
	// example:
	//
	// 9
	ExperimentId *int64 `json:"ExperimentId,omitempty" xml:"ExperimentId,omitempty"`
	// GPU name
	//
	// example:
	//
	// 8x OAM 810 GPU
	GpuName *string `json:"GpuName,omitempty" xml:"GpuName,omitempty"`
	// Number of GPUs
	//
	// example:
	//
	// 8
	GpuNum *int32 `json:"GpuNum,omitempty" xml:"GpuNum,omitempty"`
	// Service address
	//
	// example:
	//
	// whza008403
	Hostname *string `json:"Hostname,omitempty" xml:"Hostname,omitempty"`
	// Pod name.
	//
	// example:
	//
	// fluxserv-6fc89b45cf-w8wq6
	PodName *string `json:"PodName,omitempty" xml:"PodName,omitempty"`
	// Throughput
	//
	// example:
	//
	// 15
	SamplesPerSecond *float64 `json:"SamplesPerSecond,omitempty" xml:"SamplesPerSecond,omitempty"`
	// TFLOPS value
	//
	// example:
	//
	// 14
	Tflops *float64 `json:"Tflops,omitempty" xml:"Tflops,omitempty"`
	// Whether there is an alarm
	//
	// example:
	//
	// true
	WarningFlag *bool `json:"WarningFlag,omitempty" xml:"WarningFlag,omitempty"`
	// Alarm message
	//
	// example:
	//
	// warging message
	WarningMsg *string `json:"WarningMsg,omitempty" xml:"WarningMsg,omitempty"`
}

func (s GetExperimentResponseBodyDataResultsWarningWorker) String() string {
	return tea.Prettify(s)
}

func (s GetExperimentResponseBodyDataResultsWarningWorker) GoString() string {
	return s.String()
}

func (s *GetExperimentResponseBodyDataResultsWarningWorker) SetErrorFlag(v bool) *GetExperimentResponseBodyDataResultsWarningWorker {
	s.ErrorFlag = &v
	return s
}

func (s *GetExperimentResponseBodyDataResultsWarningWorker) SetErrorMsg(v string) *GetExperimentResponseBodyDataResultsWarningWorker {
	s.ErrorMsg = &v
	return s
}

func (s *GetExperimentResponseBodyDataResultsWarningWorker) SetExperimentId(v int64) *GetExperimentResponseBodyDataResultsWarningWorker {
	s.ExperimentId = &v
	return s
}

func (s *GetExperimentResponseBodyDataResultsWarningWorker) SetGpuName(v string) *GetExperimentResponseBodyDataResultsWarningWorker {
	s.GpuName = &v
	return s
}

func (s *GetExperimentResponseBodyDataResultsWarningWorker) SetGpuNum(v int32) *GetExperimentResponseBodyDataResultsWarningWorker {
	s.GpuNum = &v
	return s
}

func (s *GetExperimentResponseBodyDataResultsWarningWorker) SetHostname(v string) *GetExperimentResponseBodyDataResultsWarningWorker {
	s.Hostname = &v
	return s
}

func (s *GetExperimentResponseBodyDataResultsWarningWorker) SetPodName(v string) *GetExperimentResponseBodyDataResultsWarningWorker {
	s.PodName = &v
	return s
}

func (s *GetExperimentResponseBodyDataResultsWarningWorker) SetSamplesPerSecond(v float64) *GetExperimentResponseBodyDataResultsWarningWorker {
	s.SamplesPerSecond = &v
	return s
}

func (s *GetExperimentResponseBodyDataResultsWarningWorker) SetTflops(v float64) *GetExperimentResponseBodyDataResultsWarningWorker {
	s.Tflops = &v
	return s
}

func (s *GetExperimentResponseBodyDataResultsWarningWorker) SetWarningFlag(v bool) *GetExperimentResponseBodyDataResultsWarningWorker {
	s.WarningFlag = &v
	return s
}

func (s *GetExperimentResponseBodyDataResultsWarningWorker) SetWarningMsg(v string) *GetExperimentResponseBodyDataResultsWarningWorker {
	s.WarningMsg = &v
	return s
}

type GetExperimentResponseBodyDataTask struct {
	// Creation time
	//
	// example:
	//
	// 2024-03-05 18:24:08
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// End time
	//
	// example:
	//
	// 2024-03-05 18:34:08
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// Experiment parameters
	Params map[string]*string `json:"Params,omitempty" xml:"Params,omitempty"`
	// Scene
	//
	// example:
	//
	// baseline
	Scene *string `json:"Scene,omitempty" xml:"Scene,omitempty"`
	// Start time
	//
	// example:
	//
	// 2024-03-05 18:24:08
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// Status
	//
	// example:
	//
	// success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Task ID
	//
	// example:
	//
	// 167420
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// Update time
	//
	// example:
	//
	// 2024-03-05 18:24:08
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s GetExperimentResponseBodyDataTask) String() string {
	return tea.Prettify(s)
}

func (s GetExperimentResponseBodyDataTask) GoString() string {
	return s.String()
}

func (s *GetExperimentResponseBodyDataTask) SetCreateTime(v int64) *GetExperimentResponseBodyDataTask {
	s.CreateTime = &v
	return s
}

func (s *GetExperimentResponseBodyDataTask) SetEndTime(v int64) *GetExperimentResponseBodyDataTask {
	s.EndTime = &v
	return s
}

func (s *GetExperimentResponseBodyDataTask) SetParams(v map[string]*string) *GetExperimentResponseBodyDataTask {
	s.Params = v
	return s
}

func (s *GetExperimentResponseBodyDataTask) SetScene(v string) *GetExperimentResponseBodyDataTask {
	s.Scene = &v
	return s
}

func (s *GetExperimentResponseBodyDataTask) SetStartTime(v int64) *GetExperimentResponseBodyDataTask {
	s.StartTime = &v
	return s
}

func (s *GetExperimentResponseBodyDataTask) SetStatus(v string) *GetExperimentResponseBodyDataTask {
	s.Status = &v
	return s
}

func (s *GetExperimentResponseBodyDataTask) SetTaskId(v int64) *GetExperimentResponseBodyDataTask {
	s.TaskId = &v
	return s
}

func (s *GetExperimentResponseBodyDataTask) SetUpdateTime(v int64) *GetExperimentResponseBodyDataTask {
	s.UpdateTime = &v
	return s
}

type GetExperimentResponseBodyDataWorkload struct {
	// Default CPU allocation
	//
	// example:
	//
	// 90
	DefaultCpuPerWorker *int32 `json:"DefaultCpuPerWorker,omitempty" xml:"DefaultCpuPerWorker,omitempty"`
	// Default GPU allocation
	//
	// example:
	//
	// 8
	DefaultGpuPerWorker *int32 `json:"DefaultGpuPerWorker,omitempty" xml:"DefaultGpuPerWorker,omitempty"`
	// Default memory (GB) allocation
	//
	// example:
	//
	// 500
	DefaultMemoryPerWorker *int32 `json:"DefaultMemoryPerWorker,omitempty" xml:"DefaultMemoryPerWorker,omitempty"`
	// Default shared memory (GB) allocation
	//
	// example:
	//
	// 500
	DefaultShareMemory *int32 `json:"DefaultShareMemory,omitempty" xml:"DefaultShareMemory,omitempty"`
	// Workload cluster, AI, GPU
	//
	// example:
	//
	// AI
	Family *string `json:"Family,omitempty" xml:"Family,omitempty"`
	// JobKind
	//
	// example:
	//
	// PyTorchJob
	JobKind *string `json:"JobKind,omitempty" xml:"JobKind,omitempty"`
	// Parameter settings
	ParamSettings []*GetExperimentResponseBodyDataWorkloadParamSettings `json:"ParamSettings,omitempty" xml:"ParamSettings,omitempty" type:"Repeated"`
	// Workload usage scenario
	//
	// example:
	//
	// NLP-LLM
	Scene *string `json:"Scene,omitempty" xml:"Scene,omitempty"`
	// Scope
	//
	// example:
	//
	// common
	Scope *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
	// Static configuration
	StaticConfig *GetExperimentResponseBodyDataWorkloadStaticConfig `json:"StaticConfig,omitempty" xml:"StaticConfig,omitempty" type:"Struct"`
	// Version ID
	//
	// example:
	//
	// 1
	VersionId *int64 `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
	// Workload description
	//
	// example:
	//
	// test
	WorkloadDescription *string `json:"WorkloadDescription,omitempty" xml:"WorkloadDescription,omitempty"`
	// Workload ID
	//
	// example:
	//
	// 13
	WorkloadId *int64 `json:"WorkloadId,omitempty" xml:"WorkloadId,omitempty"`
	// Workload name
	//
	// example:
	//
	// test
	WorkloadName *string `json:"WorkloadName,omitempty" xml:"WorkloadName,omitempty"`
	// Workload name
	//
	// example:
	//
	// AI
	WorkloadType *string `json:"WorkloadType,omitempty" xml:"WorkloadType,omitempty"`
}

func (s GetExperimentResponseBodyDataWorkload) String() string {
	return tea.Prettify(s)
}

func (s GetExperimentResponseBodyDataWorkload) GoString() string {
	return s.String()
}

func (s *GetExperimentResponseBodyDataWorkload) SetDefaultCpuPerWorker(v int32) *GetExperimentResponseBodyDataWorkload {
	s.DefaultCpuPerWorker = &v
	return s
}

func (s *GetExperimentResponseBodyDataWorkload) SetDefaultGpuPerWorker(v int32) *GetExperimentResponseBodyDataWorkload {
	s.DefaultGpuPerWorker = &v
	return s
}

func (s *GetExperimentResponseBodyDataWorkload) SetDefaultMemoryPerWorker(v int32) *GetExperimentResponseBodyDataWorkload {
	s.DefaultMemoryPerWorker = &v
	return s
}

func (s *GetExperimentResponseBodyDataWorkload) SetDefaultShareMemory(v int32) *GetExperimentResponseBodyDataWorkload {
	s.DefaultShareMemory = &v
	return s
}

func (s *GetExperimentResponseBodyDataWorkload) SetFamily(v string) *GetExperimentResponseBodyDataWorkload {
	s.Family = &v
	return s
}

func (s *GetExperimentResponseBodyDataWorkload) SetJobKind(v string) *GetExperimentResponseBodyDataWorkload {
	s.JobKind = &v
	return s
}

func (s *GetExperimentResponseBodyDataWorkload) SetParamSettings(v []*GetExperimentResponseBodyDataWorkloadParamSettings) *GetExperimentResponseBodyDataWorkload {
	s.ParamSettings = v
	return s
}

func (s *GetExperimentResponseBodyDataWorkload) SetScene(v string) *GetExperimentResponseBodyDataWorkload {
	s.Scene = &v
	return s
}

func (s *GetExperimentResponseBodyDataWorkload) SetScope(v string) *GetExperimentResponseBodyDataWorkload {
	s.Scope = &v
	return s
}

func (s *GetExperimentResponseBodyDataWorkload) SetStaticConfig(v *GetExperimentResponseBodyDataWorkloadStaticConfig) *GetExperimentResponseBodyDataWorkload {
	s.StaticConfig = v
	return s
}

func (s *GetExperimentResponseBodyDataWorkload) SetVersionId(v int64) *GetExperimentResponseBodyDataWorkload {
	s.VersionId = &v
	return s
}

func (s *GetExperimentResponseBodyDataWorkload) SetWorkloadDescription(v string) *GetExperimentResponseBodyDataWorkload {
	s.WorkloadDescription = &v
	return s
}

func (s *GetExperimentResponseBodyDataWorkload) SetWorkloadId(v int64) *GetExperimentResponseBodyDataWorkload {
	s.WorkloadId = &v
	return s
}

func (s *GetExperimentResponseBodyDataWorkload) SetWorkloadName(v string) *GetExperimentResponseBodyDataWorkload {
	s.WorkloadName = &v
	return s
}

func (s *GetExperimentResponseBodyDataWorkload) SetWorkloadType(v string) *GetExperimentResponseBodyDataWorkload {
	s.WorkloadType = &v
	return s
}

type GetExperimentResponseBodyDataWorkloadParamSettings struct {
	// Default parameter value
	//
	// example:
	//
	// 100
	DefaultValue *string `json:"DefaultValue,omitempty" xml:"DefaultValue,omitempty"`
	// Parameter description
	//
	// example:
	//
	// number
	ParamDesc *string `json:"ParamDesc,omitempty" xml:"ParamDesc,omitempty"`
	// Parameter name
	//
	// example:
	//
	// ITERATION
	ParamName *string `json:"ParamName,omitempty" xml:"ParamName,omitempty"`
	// Parameter regular expression
	//
	// example:
	//
	// [0-9]+
	ParamRegex *string `json:"ParamRegex,omitempty" xml:"ParamRegex,omitempty"`
	// Parameter type
	//
	// example:
	//
	// number
	ParamType *string `json:"ParamType,omitempty" xml:"ParamType,omitempty"`
	// Parameter value
	//
	// example:
	//
	// 100
	ParamValue *string `json:"ParamValue,omitempty" xml:"ParamValue,omitempty"`
}

func (s GetExperimentResponseBodyDataWorkloadParamSettings) String() string {
	return tea.Prettify(s)
}

func (s GetExperimentResponseBodyDataWorkloadParamSettings) GoString() string {
	return s.String()
}

func (s *GetExperimentResponseBodyDataWorkloadParamSettings) SetDefaultValue(v string) *GetExperimentResponseBodyDataWorkloadParamSettings {
	s.DefaultValue = &v
	return s
}

func (s *GetExperimentResponseBodyDataWorkloadParamSettings) SetParamDesc(v string) *GetExperimentResponseBodyDataWorkloadParamSettings {
	s.ParamDesc = &v
	return s
}

func (s *GetExperimentResponseBodyDataWorkloadParamSettings) SetParamName(v string) *GetExperimentResponseBodyDataWorkloadParamSettings {
	s.ParamName = &v
	return s
}

func (s *GetExperimentResponseBodyDataWorkloadParamSettings) SetParamRegex(v string) *GetExperimentResponseBodyDataWorkloadParamSettings {
	s.ParamRegex = &v
	return s
}

func (s *GetExperimentResponseBodyDataWorkloadParamSettings) SetParamType(v string) *GetExperimentResponseBodyDataWorkloadParamSettings {
	s.ParamType = &v
	return s
}

func (s *GetExperimentResponseBodyDataWorkloadParamSettings) SetParamValue(v string) *GetExperimentResponseBodyDataWorkloadParamSettings {
	s.ParamValue = &v
	return s
}

type GetExperimentResponseBodyDataWorkloadStaticConfig struct {
	// Framework
	//
	// example:
	//
	// pyTorch
	FrameWork *string `json:"FrameWork,omitempty" xml:"FrameWork,omitempty"`
	// Operating system
	//
	// example:
	//
	// linux
	Os *string `json:"Os,omitempty" xml:"Os,omitempty"`
	// Number of parameters
	//
	// example:
	//
	// 7B
	Parameters *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// Software stack
	//
	// example:
	//
	// python
	SoftwareStack *string `json:"SoftwareStack,omitempty" xml:"SoftwareStack,omitempty"`
}

func (s GetExperimentResponseBodyDataWorkloadStaticConfig) String() string {
	return tea.Prettify(s)
}

func (s GetExperimentResponseBodyDataWorkloadStaticConfig) GoString() string {
	return s.String()
}

func (s *GetExperimentResponseBodyDataWorkloadStaticConfig) SetFrameWork(v string) *GetExperimentResponseBodyDataWorkloadStaticConfig {
	s.FrameWork = &v
	return s
}

func (s *GetExperimentResponseBodyDataWorkloadStaticConfig) SetOs(v string) *GetExperimentResponseBodyDataWorkloadStaticConfig {
	s.Os = &v
	return s
}

func (s *GetExperimentResponseBodyDataWorkloadStaticConfig) SetParameters(v string) *GetExperimentResponseBodyDataWorkloadStaticConfig {
	s.Parameters = &v
	return s
}

func (s *GetExperimentResponseBodyDataWorkloadStaticConfig) SetSoftwareStack(v string) *GetExperimentResponseBodyDataWorkloadStaticConfig {
	s.SoftwareStack = &v
	return s
}

type GetExperimentResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetExperimentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetExperimentResponse) String() string {
	return tea.Prettify(s)
}

func (s GetExperimentResponse) GoString() string {
	return s.String()
}

func (s *GetExperimentResponse) SetHeaders(v map[string]*string) *GetExperimentResponse {
	s.Headers = v
	return s
}

func (s *GetExperimentResponse) SetStatusCode(v int32) *GetExperimentResponse {
	s.StatusCode = &v
	return s
}

func (s *GetExperimentResponse) SetBody(v *GetExperimentResponseBody) *GetExperimentResponse {
	s.Body = v
	return s
}

type GetExperimentPlanRequest struct {
	// Plan ID
	//
	// This parameter is required.
	//
	// example:
	//
	// 189
	PlanId *int64 `json:"PlanId,omitempty" xml:"PlanId,omitempty"`
}

func (s GetExperimentPlanRequest) String() string {
	return tea.Prettify(s)
}

func (s GetExperimentPlanRequest) GoString() string {
	return s.String()
}

func (s *GetExperimentPlanRequest) SetPlanId(v int64) *GetExperimentPlanRequest {
	s.PlanId = &v
	return s
}

type GetExperimentPlanResponseBody struct {
	// Access denied information
	//
	// example:
	//
	// {}
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// Data
	Data *GetExperimentPlanResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// Request ID
	//
	// example:
	//
	// 6DBAC169-93D1-5DCD-8109-30FB623B3197
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Total count of the query
	//
	// example:
	//
	// 0
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s GetExperimentPlanResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetExperimentPlanResponseBody) GoString() string {
	return s.String()
}

func (s *GetExperimentPlanResponseBody) SetAccessDeniedDetail(v string) *GetExperimentPlanResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *GetExperimentPlanResponseBody) SetData(v *GetExperimentPlanResponseBodyData) *GetExperimentPlanResponseBody {
	s.Data = v
	return s
}

func (s *GetExperimentPlanResponseBody) SetRequestId(v string) *GetExperimentPlanResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetExperimentPlanResponseBody) SetTotalCount(v int64) *GetExperimentPlanResponseBody {
	s.TotalCount = &v
	return s
}

type GetExperimentPlanResponseBodyData struct {
	// Creation time
	//
	// example:
	//
	// 2024-07-07 02:08:54
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// Plan ID
	//
	// example:
	//
	// 189
	PlanId *int64 `json:"PlanId,omitempty" xml:"PlanId,omitempty"`
	// Test plan pipeline
	PlanPipeline []*GetExperimentPlanResponseBodyDataPlanPipeline `json:"PlanPipeline,omitempty" xml:"PlanPipeline,omitempty" type:"Repeated"`
	// Resource group ID
	//
	// example:
	//
	// rg-acfmvmpzi7lmxhq
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// Associated resource ID
	//
	// example:
	//
	// 260860230684
	ResourceId *int64 `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The tag.
	Tags []*GetExperimentPlanResponseBodyDataTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// Associated test plan template ID
	//
	// example:
	//
	// 2162
	TemplateId *int64 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// Associated test plan template name
	//
	// example:
	//
	// MM
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// Update time
	//
	// example:
	//
	// 2024-07-07 02:08:54
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s GetExperimentPlanResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetExperimentPlanResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetExperimentPlanResponseBodyData) SetCreateTime(v string) *GetExperimentPlanResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *GetExperimentPlanResponseBodyData) SetPlanId(v int64) *GetExperimentPlanResponseBodyData {
	s.PlanId = &v
	return s
}

func (s *GetExperimentPlanResponseBodyData) SetPlanPipeline(v []*GetExperimentPlanResponseBodyDataPlanPipeline) *GetExperimentPlanResponseBodyData {
	s.PlanPipeline = v
	return s
}

func (s *GetExperimentPlanResponseBodyData) SetResourceGroupId(v string) *GetExperimentPlanResponseBodyData {
	s.ResourceGroupId = &v
	return s
}

func (s *GetExperimentPlanResponseBodyData) SetResourceId(v int64) *GetExperimentPlanResponseBodyData {
	s.ResourceId = &v
	return s
}

func (s *GetExperimentPlanResponseBodyData) SetTags(v []*GetExperimentPlanResponseBodyDataTags) *GetExperimentPlanResponseBodyData {
	s.Tags = v
	return s
}

func (s *GetExperimentPlanResponseBodyData) SetTemplateId(v int64) *GetExperimentPlanResponseBodyData {
	s.TemplateId = &v
	return s
}

func (s *GetExperimentPlanResponseBodyData) SetTemplateName(v string) *GetExperimentPlanResponseBodyData {
	s.TemplateName = &v
	return s
}

func (s *GetExperimentPlanResponseBodyData) SetUpdateTime(v string) *GetExperimentPlanResponseBodyData {
	s.UpdateTime = &v
	return s
}

type GetExperimentPlanResponseBodyDataPlanPipeline struct {
	// Configured environment parameters
	EnvParams *GetExperimentPlanResponseBodyDataPlanPipelineEnvParams `json:"EnvParams,omitempty" xml:"EnvParams,omitempty" type:"Struct"`
	// Node order number
	//
	// example:
	//
	// 1
	PipelineOrder *int32 `json:"PipelineOrder,omitempty" xml:"PipelineOrder,omitempty"`
	// Resource ID
	//
	// example:
	//
	// 36
	ResourceId *int64 `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// Resource name
	//
	// example:
	//
	// PPU
	ResourceName *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	// Usage scenario, e.g., "baseline"
	//
	// example:
	//
	// baseline
	Scene *string `json:"Scene,omitempty" xml:"Scene,omitempty"`
	// Configured workload parameters
	SettingParams map[string]*string `json:"SettingParams,omitempty" xml:"SettingParams,omitempty"`
	// Workload ID
	//
	// example:
	//
	// 14
	WorkloadId *int64 `json:"WorkloadId,omitempty" xml:"WorkloadId,omitempty"`
	// Workload name
	//
	// example:
	//
	// test
	WorkloadName *string `json:"WorkloadName,omitempty" xml:"WorkloadName,omitempty"`
}

func (s GetExperimentPlanResponseBodyDataPlanPipeline) String() string {
	return tea.Prettify(s)
}

func (s GetExperimentPlanResponseBodyDataPlanPipeline) GoString() string {
	return s.String()
}

func (s *GetExperimentPlanResponseBodyDataPlanPipeline) SetEnvParams(v *GetExperimentPlanResponseBodyDataPlanPipelineEnvParams) *GetExperimentPlanResponseBodyDataPlanPipeline {
	s.EnvParams = v
	return s
}

func (s *GetExperimentPlanResponseBodyDataPlanPipeline) SetPipelineOrder(v int32) *GetExperimentPlanResponseBodyDataPlanPipeline {
	s.PipelineOrder = &v
	return s
}

func (s *GetExperimentPlanResponseBodyDataPlanPipeline) SetResourceId(v int64) *GetExperimentPlanResponseBodyDataPlanPipeline {
	s.ResourceId = &v
	return s
}

func (s *GetExperimentPlanResponseBodyDataPlanPipeline) SetResourceName(v string) *GetExperimentPlanResponseBodyDataPlanPipeline {
	s.ResourceName = &v
	return s
}

func (s *GetExperimentPlanResponseBodyDataPlanPipeline) SetScene(v string) *GetExperimentPlanResponseBodyDataPlanPipeline {
	s.Scene = &v
	return s
}

func (s *GetExperimentPlanResponseBodyDataPlanPipeline) SetSettingParams(v map[string]*string) *GetExperimentPlanResponseBodyDataPlanPipeline {
	s.SettingParams = v
	return s
}

func (s *GetExperimentPlanResponseBodyDataPlanPipeline) SetWorkloadId(v int64) *GetExperimentPlanResponseBodyDataPlanPipeline {
	s.WorkloadId = &v
	return s
}

func (s *GetExperimentPlanResponseBodyDataPlanPipeline) SetWorkloadName(v string) *GetExperimentPlanResponseBodyDataPlanPipeline {
	s.WorkloadName = &v
	return s
}

type GetExperimentPlanResponseBodyDataPlanPipelineEnvParams struct {
	// CPU allocation
	//
	// example:
	//
	// 90
	CpuPerWorker *int32 `json:"CpuPerWorker,omitempty" xml:"CpuPerWorker,omitempty"`
	// CUDA version
	//
	// example:
	//
	// 1.0.0
	CudaVersion *string `json:"CudaVersion,omitempty" xml:"CudaVersion,omitempty"`
	// Additional parameters
	ExtendParam map[string]*string `json:"ExtendParam,omitempty" xml:"ExtendParam,omitempty"`
	// GPU driver version
	//
	// example:
	//
	// 1.0.0
	GpuDriverVersion *string `json:"GpuDriverVersion,omitempty" xml:"GpuDriverVersion,omitempty"`
	// Number of GPUs allocated
	//
	// example:
	//
	// 8
	GpuPerWorker *int32 `json:"GpuPerWorker,omitempty" xml:"GpuPerWorker,omitempty"`
	// Memory GB allocation
	//
	// example:
	//
	// 500
	MemoryPerWorker *int32 `json:"MemoryPerWorker,omitempty" xml:"MemoryPerWorker,omitempty"`
	// NCCL version
	//
	// example:
	//
	// 1.0.0
	NCCLVersion *string `json:"NCCLVersion,omitempty" xml:"NCCLVersion,omitempty"`
	// PyTorch version
	//
	// example:
	//
	// 1.0.0
	PyTorchVersion *string `json:"PyTorchVersion,omitempty" xml:"PyTorchVersion,omitempty"`
	// Specified nodes
	ResourceNodes []*GetExperimentPlanResponseBodyDataPlanPipelineEnvParamsResourceNodes `json:"ResourceNodes,omitempty" xml:"ResourceNodes,omitempty" type:"Repeated"`
	// Shared memory GB allocation
	//
	// example:
	//
	// 500
	ShareMemory *int32 `json:"ShareMemory,omitempty" xml:"ShareMemory,omitempty"`
	// Number of nodes
	//
	// example:
	//
	// 1
	WorkerNum *int32 `json:"WorkerNum,omitempty" xml:"WorkerNum,omitempty"`
}

func (s GetExperimentPlanResponseBodyDataPlanPipelineEnvParams) String() string {
	return tea.Prettify(s)
}

func (s GetExperimentPlanResponseBodyDataPlanPipelineEnvParams) GoString() string {
	return s.String()
}

func (s *GetExperimentPlanResponseBodyDataPlanPipelineEnvParams) SetCpuPerWorker(v int32) *GetExperimentPlanResponseBodyDataPlanPipelineEnvParams {
	s.CpuPerWorker = &v
	return s
}

func (s *GetExperimentPlanResponseBodyDataPlanPipelineEnvParams) SetCudaVersion(v string) *GetExperimentPlanResponseBodyDataPlanPipelineEnvParams {
	s.CudaVersion = &v
	return s
}

func (s *GetExperimentPlanResponseBodyDataPlanPipelineEnvParams) SetExtendParam(v map[string]*string) *GetExperimentPlanResponseBodyDataPlanPipelineEnvParams {
	s.ExtendParam = v
	return s
}

func (s *GetExperimentPlanResponseBodyDataPlanPipelineEnvParams) SetGpuDriverVersion(v string) *GetExperimentPlanResponseBodyDataPlanPipelineEnvParams {
	s.GpuDriverVersion = &v
	return s
}

func (s *GetExperimentPlanResponseBodyDataPlanPipelineEnvParams) SetGpuPerWorker(v int32) *GetExperimentPlanResponseBodyDataPlanPipelineEnvParams {
	s.GpuPerWorker = &v
	return s
}

func (s *GetExperimentPlanResponseBodyDataPlanPipelineEnvParams) SetMemoryPerWorker(v int32) *GetExperimentPlanResponseBodyDataPlanPipelineEnvParams {
	s.MemoryPerWorker = &v
	return s
}

func (s *GetExperimentPlanResponseBodyDataPlanPipelineEnvParams) SetNCCLVersion(v string) *GetExperimentPlanResponseBodyDataPlanPipelineEnvParams {
	s.NCCLVersion = &v
	return s
}

func (s *GetExperimentPlanResponseBodyDataPlanPipelineEnvParams) SetPyTorchVersion(v string) *GetExperimentPlanResponseBodyDataPlanPipelineEnvParams {
	s.PyTorchVersion = &v
	return s
}

func (s *GetExperimentPlanResponseBodyDataPlanPipelineEnvParams) SetResourceNodes(v []*GetExperimentPlanResponseBodyDataPlanPipelineEnvParamsResourceNodes) *GetExperimentPlanResponseBodyDataPlanPipelineEnvParams {
	s.ResourceNodes = v
	return s
}

func (s *GetExperimentPlanResponseBodyDataPlanPipelineEnvParams) SetShareMemory(v int32) *GetExperimentPlanResponseBodyDataPlanPipelineEnvParams {
	s.ShareMemory = &v
	return s
}

func (s *GetExperimentPlanResponseBodyDataPlanPipelineEnvParams) SetWorkerNum(v int32) *GetExperimentPlanResponseBodyDataPlanPipelineEnvParams {
	s.WorkerNum = &v
	return s
}

type GetExperimentPlanResponseBodyDataPlanPipelineEnvParamsResourceNodes struct {
	// Node name
	//
	// example:
	//
	// ods_galaxy_gateway_tickets
	NodeName *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	// Requested CPU
	//
	// example:
	//
	// 90
	RequestCPU *int32 `json:"RequestCPU,omitempty" xml:"RequestCPU,omitempty"`
	// Requested GPU
	//
	// example:
	//
	// 8
	RequestGPU *int32 `json:"RequestGPU,omitempty" xml:"RequestGPU,omitempty"`
	// Memory of the current request
	//
	// example:
	//
	// 500
	RequestMemory *int32 `json:"RequestMemory,omitempty" xml:"RequestMemory,omitempty"`
	// Total CPU
	//
	// example:
	//
	// 90
	TotalCPU *int32 `json:"TotalCPU,omitempty" xml:"TotalCPU,omitempty"`
	// Total GPU
	//
	// example:
	//
	// 8
	TotalGPU *int32 `json:"TotalGPU,omitempty" xml:"TotalGPU,omitempty"`
	// Total memory
	//
	// example:
	//
	// 500
	TotalMemory *int64 `json:"TotalMemory,omitempty" xml:"TotalMemory,omitempty"`
}

func (s GetExperimentPlanResponseBodyDataPlanPipelineEnvParamsResourceNodes) String() string {
	return tea.Prettify(s)
}

func (s GetExperimentPlanResponseBodyDataPlanPipelineEnvParamsResourceNodes) GoString() string {
	return s.String()
}

func (s *GetExperimentPlanResponseBodyDataPlanPipelineEnvParamsResourceNodes) SetNodeName(v string) *GetExperimentPlanResponseBodyDataPlanPipelineEnvParamsResourceNodes {
	s.NodeName = &v
	return s
}

func (s *GetExperimentPlanResponseBodyDataPlanPipelineEnvParamsResourceNodes) SetRequestCPU(v int32) *GetExperimentPlanResponseBodyDataPlanPipelineEnvParamsResourceNodes {
	s.RequestCPU = &v
	return s
}

func (s *GetExperimentPlanResponseBodyDataPlanPipelineEnvParamsResourceNodes) SetRequestGPU(v int32) *GetExperimentPlanResponseBodyDataPlanPipelineEnvParamsResourceNodes {
	s.RequestGPU = &v
	return s
}

func (s *GetExperimentPlanResponseBodyDataPlanPipelineEnvParamsResourceNodes) SetRequestMemory(v int32) *GetExperimentPlanResponseBodyDataPlanPipelineEnvParamsResourceNodes {
	s.RequestMemory = &v
	return s
}

func (s *GetExperimentPlanResponseBodyDataPlanPipelineEnvParamsResourceNodes) SetTotalCPU(v int32) *GetExperimentPlanResponseBodyDataPlanPipelineEnvParamsResourceNodes {
	s.TotalCPU = &v
	return s
}

func (s *GetExperimentPlanResponseBodyDataPlanPipelineEnvParamsResourceNodes) SetTotalGPU(v int32) *GetExperimentPlanResponseBodyDataPlanPipelineEnvParamsResourceNodes {
	s.TotalGPU = &v
	return s
}

func (s *GetExperimentPlanResponseBodyDataPlanPipelineEnvParamsResourceNodes) SetTotalMemory(v int64) *GetExperimentPlanResponseBodyDataPlanPipelineEnvParamsResourceNodes {
	s.TotalMemory = &v
	return s
}

type GetExperimentPlanResponseBodyDataTags struct {
	// The tag key.
	//
	// example:
	//
	// acs:testLXP:test-quota40-19
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The tag value.
	//
	// example:
	//
	// 000088aabb0019e4
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s GetExperimentPlanResponseBodyDataTags) String() string {
	return tea.Prettify(s)
}

func (s GetExperimentPlanResponseBodyDataTags) GoString() string {
	return s.String()
}

func (s *GetExperimentPlanResponseBodyDataTags) SetTagKey(v string) *GetExperimentPlanResponseBodyDataTags {
	s.TagKey = &v
	return s
}

func (s *GetExperimentPlanResponseBodyDataTags) SetTagValue(v string) *GetExperimentPlanResponseBodyDataTags {
	s.TagValue = &v
	return s
}

type GetExperimentPlanResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetExperimentPlanResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetExperimentPlanResponse) String() string {
	return tea.Prettify(s)
}

func (s GetExperimentPlanResponse) GoString() string {
	return s.String()
}

func (s *GetExperimentPlanResponse) SetHeaders(v map[string]*string) *GetExperimentPlanResponse {
	s.Headers = v
	return s
}

func (s *GetExperimentPlanResponse) SetStatusCode(v int32) *GetExperimentPlanResponse {
	s.StatusCode = &v
	return s
}

func (s *GetExperimentPlanResponse) SetBody(v *GetExperimentPlanResponseBody) *GetExperimentPlanResponse {
	s.Body = v
	return s
}

type GetExperimentPlanTemplateRequest struct {
	// Template ID
	//
	// This parameter is required.
	//
	// example:
	//
	// 315797
	TemplateId *int64 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s GetExperimentPlanTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s GetExperimentPlanTemplateRequest) GoString() string {
	return s.String()
}

func (s *GetExperimentPlanTemplateRequest) SetTemplateId(v int64) *GetExperimentPlanTemplateRequest {
	s.TemplateId = &v
	return s
}

type GetExperimentPlanTemplateResponseBody struct {
	// Access denied information
	//
	// example:
	//
	// {}
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// Data
	Data *GetExperimentPlanTemplateResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// Request ID
	//
	// example:
	//
	// 5514CB39-B7C0-5B89-8534-2DE1E0F2B7AB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Total
	//
	// example:
	//
	// 0
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s GetExperimentPlanTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetExperimentPlanTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *GetExperimentPlanTemplateResponseBody) SetAccessDeniedDetail(v string) *GetExperimentPlanTemplateResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *GetExperimentPlanTemplateResponseBody) SetData(v *GetExperimentPlanTemplateResponseBodyData) *GetExperimentPlanTemplateResponseBody {
	s.Data = v
	return s
}

func (s *GetExperimentPlanTemplateResponseBody) SetRequestId(v string) *GetExperimentPlanTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetExperimentPlanTemplateResponseBody) SetTotalCount(v int64) *GetExperimentPlanTemplateResponseBody {
	s.TotalCount = &v
	return s
}

type GetExperimentPlanTemplateResponseBodyData struct {
	// Creation Time
	//
	// example:
	//
	// 2024-11-29 02:16:35
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// Primary account UID
	//
	// example:
	//
	// 12312312312312
	CreatorUid *int64 `json:"CreatorUid,omitempty" xml:"CreatorUid,omitempty"`
	// Whether deleted
	//
	// example:
	//
	// 0
	IsDelete *int32 `json:"IsDelete,omitempty" xml:"IsDelete,omitempty"`
	// Privacy Level
	//
	// example:
	//
	// private
	PrivacyLevel *string `json:"PrivacyLevel,omitempty" xml:"PrivacyLevel,omitempty"`
	// Template Code
	//
	// example:
	//
	// 464086216
	TemplateCode *int64 `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
	// Template Description
	//
	// example:
	//
	// test
	TemplateDescription *string `json:"TemplateDescription,omitempty" xml:"TemplateDescription,omitempty"`
	// Template ID
	//
	// example:
	//
	// 17615126
	TemplateId *int64 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// Template Name
	//
	// example:
	//
	// Test
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// Template Pipeline
	TemplatePipelineParam []*GetExperimentPlanTemplateResponseBodyDataTemplatePipelineParam `json:"TemplatePipelineParam,omitempty" xml:"TemplatePipelineParam,omitempty" type:"Repeated"`
	// Update Time
	//
	// example:
	//
	// 2024-10-22 10:18:10
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// Version ID
	//
	// example:
	//
	// 1
	VersionId *int64 `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
}

func (s GetExperimentPlanTemplateResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetExperimentPlanTemplateResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetExperimentPlanTemplateResponseBodyData) SetCreateTime(v string) *GetExperimentPlanTemplateResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *GetExperimentPlanTemplateResponseBodyData) SetCreatorUid(v int64) *GetExperimentPlanTemplateResponseBodyData {
	s.CreatorUid = &v
	return s
}

func (s *GetExperimentPlanTemplateResponseBodyData) SetIsDelete(v int32) *GetExperimentPlanTemplateResponseBodyData {
	s.IsDelete = &v
	return s
}

func (s *GetExperimentPlanTemplateResponseBodyData) SetPrivacyLevel(v string) *GetExperimentPlanTemplateResponseBodyData {
	s.PrivacyLevel = &v
	return s
}

func (s *GetExperimentPlanTemplateResponseBodyData) SetTemplateCode(v int64) *GetExperimentPlanTemplateResponseBodyData {
	s.TemplateCode = &v
	return s
}

func (s *GetExperimentPlanTemplateResponseBodyData) SetTemplateDescription(v string) *GetExperimentPlanTemplateResponseBodyData {
	s.TemplateDescription = &v
	return s
}

func (s *GetExperimentPlanTemplateResponseBodyData) SetTemplateId(v int64) *GetExperimentPlanTemplateResponseBodyData {
	s.TemplateId = &v
	return s
}

func (s *GetExperimentPlanTemplateResponseBodyData) SetTemplateName(v string) *GetExperimentPlanTemplateResponseBodyData {
	s.TemplateName = &v
	return s
}

func (s *GetExperimentPlanTemplateResponseBodyData) SetTemplatePipelineParam(v []*GetExperimentPlanTemplateResponseBodyDataTemplatePipelineParam) *GetExperimentPlanTemplateResponseBodyData {
	s.TemplatePipelineParam = v
	return s
}

func (s *GetExperimentPlanTemplateResponseBodyData) SetUpdateTime(v string) *GetExperimentPlanTemplateResponseBodyData {
	s.UpdateTime = &v
	return s
}

func (s *GetExperimentPlanTemplateResponseBodyData) SetVersionId(v int64) *GetExperimentPlanTemplateResponseBodyData {
	s.VersionId = &v
	return s
}

type GetExperimentPlanTemplateResponseBodyDataTemplatePipelineParam struct {
	// Configured environment parameters
	EnvParams *GetExperimentPlanTemplateResponseBodyDataTemplatePipelineParamEnvParams `json:"EnvParams,omitempty" xml:"EnvParams,omitempty" type:"Struct"`
	// Node sequence number
	//
	// example:
	//
	// 1
	PipelineOrder *int32 `json:"PipelineOrder,omitempty" xml:"PipelineOrder,omitempty"`
	// Usage scenario, e.g., "baseline"
	//
	// example:
	//
	// baseline
	Scene *string `json:"Scene,omitempty" xml:"Scene,omitempty"`
	// Configured workload parameters
	SettingParams map[string]*string `json:"SettingParams,omitempty" xml:"SettingParams,omitempty"`
	// Workload ID
	//
	// example:
	//
	// 13
	WorkloadId *int64 `json:"WorkloadId,omitempty" xml:"WorkloadId,omitempty"`
	// Workload Name
	//
	// example:
	//
	// test
	WorkloadName *string `json:"WorkloadName,omitempty" xml:"WorkloadName,omitempty"`
}

func (s GetExperimentPlanTemplateResponseBodyDataTemplatePipelineParam) String() string {
	return tea.Prettify(s)
}

func (s GetExperimentPlanTemplateResponseBodyDataTemplatePipelineParam) GoString() string {
	return s.String()
}

func (s *GetExperimentPlanTemplateResponseBodyDataTemplatePipelineParam) SetEnvParams(v *GetExperimentPlanTemplateResponseBodyDataTemplatePipelineParamEnvParams) *GetExperimentPlanTemplateResponseBodyDataTemplatePipelineParam {
	s.EnvParams = v
	return s
}

func (s *GetExperimentPlanTemplateResponseBodyDataTemplatePipelineParam) SetPipelineOrder(v int32) *GetExperimentPlanTemplateResponseBodyDataTemplatePipelineParam {
	s.PipelineOrder = &v
	return s
}

func (s *GetExperimentPlanTemplateResponseBodyDataTemplatePipelineParam) SetScene(v string) *GetExperimentPlanTemplateResponseBodyDataTemplatePipelineParam {
	s.Scene = &v
	return s
}

func (s *GetExperimentPlanTemplateResponseBodyDataTemplatePipelineParam) SetSettingParams(v map[string]*string) *GetExperimentPlanTemplateResponseBodyDataTemplatePipelineParam {
	s.SettingParams = v
	return s
}

func (s *GetExperimentPlanTemplateResponseBodyDataTemplatePipelineParam) SetWorkloadId(v int64) *GetExperimentPlanTemplateResponseBodyDataTemplatePipelineParam {
	s.WorkloadId = &v
	return s
}

func (s *GetExperimentPlanTemplateResponseBodyDataTemplatePipelineParam) SetWorkloadName(v string) *GetExperimentPlanTemplateResponseBodyDataTemplatePipelineParam {
	s.WorkloadName = &v
	return s
}

type GetExperimentPlanTemplateResponseBodyDataTemplatePipelineParamEnvParams struct {
	// CPU allocation
	//
	// example:
	//
	// 90
	CpuPerWorker *int32 `json:"CpuPerWorker,omitempty" xml:"CpuPerWorker,omitempty"`
	// CUDA version
	//
	// example:
	//
	// 1.0.0
	CudaVersion *string `json:"CudaVersion,omitempty" xml:"CudaVersion,omitempty"`
	// Additional parameters
	ExtendParam map[string]*string `json:"ExtendParam,omitempty" xml:"ExtendParam,omitempty"`
	// GPU driver version
	//
	// example:
	//
	// 1.0.0
	GpuDriverVersion *string `json:"GpuDriverVersion,omitempty" xml:"GpuDriverVersion,omitempty"`
	// GPU allocation
	//
	// example:
	//
	// 8
	GpuPerWorker *int32 `json:"GpuPerWorker,omitempty" xml:"GpuPerWorker,omitempty"`
	// Allocated memory in GB
	//
	// example:
	//
	// 500
	MemoryPerWorker *int32 `json:"MemoryPerWorker,omitempty" xml:"MemoryPerWorker,omitempty"`
	// NCCL version
	//
	// example:
	//
	// 1.0.0
	NCCLVersion *string `json:"NCCLVersion,omitempty" xml:"NCCLVersion,omitempty"`
	// PyTorch version
	//
	// example:
	//
	// 1.0.0
	PyTorchVersion *string `json:"PyTorchVersion,omitempty" xml:"PyTorchVersion,omitempty"`
	// Specified nodes
	ResourceNodes []*GetExperimentPlanTemplateResponseBodyDataTemplatePipelineParamEnvParamsResourceNodes `json:"ResourceNodes,omitempty" xml:"ResourceNodes,omitempty" type:"Repeated"`
	// Shared memory in GB
	//
	// example:
	//
	// 500
	ShareMemory *int32 `json:"ShareMemory,omitempty" xml:"ShareMemory,omitempty"`
	// Number of nodes
	//
	// example:
	//
	// 1
	WorkerNum *int32 `json:"WorkerNum,omitempty" xml:"WorkerNum,omitempty"`
}

func (s GetExperimentPlanTemplateResponseBodyDataTemplatePipelineParamEnvParams) String() string {
	return tea.Prettify(s)
}

func (s GetExperimentPlanTemplateResponseBodyDataTemplatePipelineParamEnvParams) GoString() string {
	return s.String()
}

func (s *GetExperimentPlanTemplateResponseBodyDataTemplatePipelineParamEnvParams) SetCpuPerWorker(v int32) *GetExperimentPlanTemplateResponseBodyDataTemplatePipelineParamEnvParams {
	s.CpuPerWorker = &v
	return s
}

func (s *GetExperimentPlanTemplateResponseBodyDataTemplatePipelineParamEnvParams) SetCudaVersion(v string) *GetExperimentPlanTemplateResponseBodyDataTemplatePipelineParamEnvParams {
	s.CudaVersion = &v
	return s
}

func (s *GetExperimentPlanTemplateResponseBodyDataTemplatePipelineParamEnvParams) SetExtendParam(v map[string]*string) *GetExperimentPlanTemplateResponseBodyDataTemplatePipelineParamEnvParams {
	s.ExtendParam = v
	return s
}

func (s *GetExperimentPlanTemplateResponseBodyDataTemplatePipelineParamEnvParams) SetGpuDriverVersion(v string) *GetExperimentPlanTemplateResponseBodyDataTemplatePipelineParamEnvParams {
	s.GpuDriverVersion = &v
	return s
}

func (s *GetExperimentPlanTemplateResponseBodyDataTemplatePipelineParamEnvParams) SetGpuPerWorker(v int32) *GetExperimentPlanTemplateResponseBodyDataTemplatePipelineParamEnvParams {
	s.GpuPerWorker = &v
	return s
}

func (s *GetExperimentPlanTemplateResponseBodyDataTemplatePipelineParamEnvParams) SetMemoryPerWorker(v int32) *GetExperimentPlanTemplateResponseBodyDataTemplatePipelineParamEnvParams {
	s.MemoryPerWorker = &v
	return s
}

func (s *GetExperimentPlanTemplateResponseBodyDataTemplatePipelineParamEnvParams) SetNCCLVersion(v string) *GetExperimentPlanTemplateResponseBodyDataTemplatePipelineParamEnvParams {
	s.NCCLVersion = &v
	return s
}

func (s *GetExperimentPlanTemplateResponseBodyDataTemplatePipelineParamEnvParams) SetPyTorchVersion(v string) *GetExperimentPlanTemplateResponseBodyDataTemplatePipelineParamEnvParams {
	s.PyTorchVersion = &v
	return s
}

func (s *GetExperimentPlanTemplateResponseBodyDataTemplatePipelineParamEnvParams) SetResourceNodes(v []*GetExperimentPlanTemplateResponseBodyDataTemplatePipelineParamEnvParamsResourceNodes) *GetExperimentPlanTemplateResponseBodyDataTemplatePipelineParamEnvParams {
	s.ResourceNodes = v
	return s
}

func (s *GetExperimentPlanTemplateResponseBodyDataTemplatePipelineParamEnvParams) SetShareMemory(v int32) *GetExperimentPlanTemplateResponseBodyDataTemplatePipelineParamEnvParams {
	s.ShareMemory = &v
	return s
}

func (s *GetExperimentPlanTemplateResponseBodyDataTemplatePipelineParamEnvParams) SetWorkerNum(v int32) *GetExperimentPlanTemplateResponseBodyDataTemplatePipelineParamEnvParams {
	s.WorkerNum = &v
	return s
}

type GetExperimentPlanTemplateResponseBodyDataTemplatePipelineParamEnvParamsResourceNodes struct {
	// Node name
	//
	// example:
	//
	// exclusive_coud
	NodeName *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	// 当前请求的cpu
	//
	// example:
	//
	// 10
	RequestCPU *int32 `json:"RequestCPU,omitempty" xml:"RequestCPU,omitempty"`
	// Requested GPU
	//
	// example:
	//
	// 10
	RequestGPU *int32 `json:"RequestGPU,omitempty" xml:"RequestGPU,omitempty"`
	// Requested memory
	//
	// example:
	//
	// 1024
	RequestMemory *int32 `json:"RequestMemory,omitempty" xml:"RequestMemory,omitempty"`
	// Total CPU
	//
	// example:
	//
	// 100
	TotalCPU *int32 `json:"TotalCPU,omitempty" xml:"TotalCPU,omitempty"`
	// Total GPU
	//
	// example:
	//
	// 100
	TotalGPU *int32 `json:"TotalGPU,omitempty" xml:"TotalGPU,omitempty"`
	// Total memory
	//
	// example:
	//
	// 2048
	TotalMemory *int64 `json:"TotalMemory,omitempty" xml:"TotalMemory,omitempty"`
}

func (s GetExperimentPlanTemplateResponseBodyDataTemplatePipelineParamEnvParamsResourceNodes) String() string {
	return tea.Prettify(s)
}

func (s GetExperimentPlanTemplateResponseBodyDataTemplatePipelineParamEnvParamsResourceNodes) GoString() string {
	return s.String()
}

func (s *GetExperimentPlanTemplateResponseBodyDataTemplatePipelineParamEnvParamsResourceNodes) SetNodeName(v string) *GetExperimentPlanTemplateResponseBodyDataTemplatePipelineParamEnvParamsResourceNodes {
	s.NodeName = &v
	return s
}

func (s *GetExperimentPlanTemplateResponseBodyDataTemplatePipelineParamEnvParamsResourceNodes) SetRequestCPU(v int32) *GetExperimentPlanTemplateResponseBodyDataTemplatePipelineParamEnvParamsResourceNodes {
	s.RequestCPU = &v
	return s
}

func (s *GetExperimentPlanTemplateResponseBodyDataTemplatePipelineParamEnvParamsResourceNodes) SetRequestGPU(v int32) *GetExperimentPlanTemplateResponseBodyDataTemplatePipelineParamEnvParamsResourceNodes {
	s.RequestGPU = &v
	return s
}

func (s *GetExperimentPlanTemplateResponseBodyDataTemplatePipelineParamEnvParamsResourceNodes) SetRequestMemory(v int32) *GetExperimentPlanTemplateResponseBodyDataTemplatePipelineParamEnvParamsResourceNodes {
	s.RequestMemory = &v
	return s
}

func (s *GetExperimentPlanTemplateResponseBodyDataTemplatePipelineParamEnvParamsResourceNodes) SetTotalCPU(v int32) *GetExperimentPlanTemplateResponseBodyDataTemplatePipelineParamEnvParamsResourceNodes {
	s.TotalCPU = &v
	return s
}

func (s *GetExperimentPlanTemplateResponseBodyDataTemplatePipelineParamEnvParamsResourceNodes) SetTotalGPU(v int32) *GetExperimentPlanTemplateResponseBodyDataTemplatePipelineParamEnvParamsResourceNodes {
	s.TotalGPU = &v
	return s
}

func (s *GetExperimentPlanTemplateResponseBodyDataTemplatePipelineParamEnvParamsResourceNodes) SetTotalMemory(v int64) *GetExperimentPlanTemplateResponseBodyDataTemplatePipelineParamEnvParamsResourceNodes {
	s.TotalMemory = &v
	return s
}

type GetExperimentPlanTemplateResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetExperimentPlanTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetExperimentPlanTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s GetExperimentPlanTemplateResponse) GoString() string {
	return s.String()
}

func (s *GetExperimentPlanTemplateResponse) SetHeaders(v map[string]*string) *GetExperimentPlanTemplateResponse {
	s.Headers = v
	return s
}

func (s *GetExperimentPlanTemplateResponse) SetStatusCode(v int32) *GetExperimentPlanTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *GetExperimentPlanTemplateResponse) SetBody(v *GetExperimentPlanTemplateResponseBody) *GetExperimentPlanTemplateResponse {
	s.Body = v
	return s
}

type GetExperimentResultDataRequest struct {
	// Experiment ID
	//
	// This parameter is required.
	//
	// example:
	//
	// 234
	ExperimentId *int64 `json:"ExperimentId,omitempty" xml:"ExperimentId,omitempty"`
	// Hostname
	//
	// example:
	//
	// iZj6ccwd7zwfms7hzaz2riZ
	Hostname *string `json:"Hostname,omitempty" xml:"Hostname,omitempty"`
	// Resource Group Id
	//
	// example:
	//
	// rg-sfjgskdfj3k4
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// Workload Type
	//
	// example:
	//
	// AI
	WorkloadType *string `json:"WorkloadType,omitempty" xml:"WorkloadType,omitempty"`
}

func (s GetExperimentResultDataRequest) String() string {
	return tea.Prettify(s)
}

func (s GetExperimentResultDataRequest) GoString() string {
	return s.String()
}

func (s *GetExperimentResultDataRequest) SetExperimentId(v int64) *GetExperimentResultDataRequest {
	s.ExperimentId = &v
	return s
}

func (s *GetExperimentResultDataRequest) SetHostname(v string) *GetExperimentResultDataRequest {
	s.Hostname = &v
	return s
}

func (s *GetExperimentResultDataRequest) SetResourceGroupId(v string) *GetExperimentResultDataRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *GetExperimentResultDataRequest) SetWorkloadType(v string) *GetExperimentResultDataRequest {
	s.WorkloadType = &v
	return s
}

type GetExperimentResultDataResponseBody struct {
	// Access Denied Details
	//
	// example:
	//
	// {}
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// Data
	Data []*GetExperimentResultDataResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// Request ID
	//
	// example:
	//
	// C1D34EC2-AB13-56F4-8322-F15AE563EA04
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Total Count of Queries
	//
	// example:
	//
	// 0
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s GetExperimentResultDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetExperimentResultDataResponseBody) GoString() string {
	return s.String()
}

func (s *GetExperimentResultDataResponseBody) SetAccessDeniedDetail(v string) *GetExperimentResultDataResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *GetExperimentResultDataResponseBody) SetData(v []*GetExperimentResultDataResponseBodyData) *GetExperimentResultDataResponseBody {
	s.Data = v
	return s
}

func (s *GetExperimentResultDataResponseBody) SetRequestId(v string) *GetExperimentResultDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetExperimentResultDataResponseBody) SetTotalCount(v int64) *GetExperimentResultDataResponseBody {
	s.TotalCount = &v
	return s
}

type GetExperimentResultDataResponseBodyData struct {
	// Number of GPUs
	//
	// example:
	//
	// 8
	GpuNum *string `json:"GpuNum,omitempty" xml:"GpuNum,omitempty"`
	// Host IP
	//
	// example:
	//
	// p-jt-waf-app1
	Hostname *string `json:"Hostname,omitempty" xml:"Hostname,omitempty"`
	// List of Metrics Information
	MetricsInfos []*GetExperimentResultDataResponseBodyDataMetricsInfos `json:"MetricsInfos,omitempty" xml:"MetricsInfos,omitempty" type:"Repeated"`
	// Pod Name
	//
	// example:
	//
	// hzs-forge-sdxl-online-7ff4d86444-pc95h
	PodName *string `json:"PodName,omitempty" xml:"PodName,omitempty"`
}

func (s GetExperimentResultDataResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetExperimentResultDataResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetExperimentResultDataResponseBodyData) SetGpuNum(v string) *GetExperimentResultDataResponseBodyData {
	s.GpuNum = &v
	return s
}

func (s *GetExperimentResultDataResponseBodyData) SetHostname(v string) *GetExperimentResultDataResponseBodyData {
	s.Hostname = &v
	return s
}

func (s *GetExperimentResultDataResponseBodyData) SetMetricsInfos(v []*GetExperimentResultDataResponseBodyDataMetricsInfos) *GetExperimentResultDataResponseBodyData {
	s.MetricsInfos = v
	return s
}

func (s *GetExperimentResultDataResponseBodyData) SetPodName(v string) *GetExperimentResultDataResponseBodyData {
	s.PodName = &v
	return s
}

type GetExperimentResultDataResponseBodyDataMetricsInfos struct {
	// GPU
	//
	// example:
	//
	// 8
	GpuNum *string `json:"Gpu_num,omitempty" xml:"Gpu_num,omitempty"`
	// Iteration
	//
	// example:
	//
	// 100
	Iteration *int64 `json:"Iteration,omitempty" xml:"Iteration,omitempty"`
	// TFLOPS
	//
	// example:
	//
	// 43
	Tflops *float64 `json:"Tflops,omitempty" xml:"Tflops,omitempty"`
	// Operation Timestamp
	//
	// example:
	//
	// 1715393860
	Timestamp *int64 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	// Metric Value
	//
	// example:
	//
	// 126
	Value *float64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetExperimentResultDataResponseBodyDataMetricsInfos) String() string {
	return tea.Prettify(s)
}

func (s GetExperimentResultDataResponseBodyDataMetricsInfos) GoString() string {
	return s.String()
}

func (s *GetExperimentResultDataResponseBodyDataMetricsInfos) SetGpuNum(v string) *GetExperimentResultDataResponseBodyDataMetricsInfos {
	s.GpuNum = &v
	return s
}

func (s *GetExperimentResultDataResponseBodyDataMetricsInfos) SetIteration(v int64) *GetExperimentResultDataResponseBodyDataMetricsInfos {
	s.Iteration = &v
	return s
}

func (s *GetExperimentResultDataResponseBodyDataMetricsInfos) SetTflops(v float64) *GetExperimentResultDataResponseBodyDataMetricsInfos {
	s.Tflops = &v
	return s
}

func (s *GetExperimentResultDataResponseBodyDataMetricsInfos) SetTimestamp(v int64) *GetExperimentResultDataResponseBodyDataMetricsInfos {
	s.Timestamp = &v
	return s
}

func (s *GetExperimentResultDataResponseBodyDataMetricsInfos) SetValue(v float64) *GetExperimentResultDataResponseBodyDataMetricsInfos {
	s.Value = &v
	return s
}

type GetExperimentResultDataResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetExperimentResultDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetExperimentResultDataResponse) String() string {
	return tea.Prettify(s)
}

func (s GetExperimentResultDataResponse) GoString() string {
	return s.String()
}

func (s *GetExperimentResultDataResponse) SetHeaders(v map[string]*string) *GetExperimentResultDataResponse {
	s.Headers = v
	return s
}

func (s *GetExperimentResultDataResponse) SetStatusCode(v int32) *GetExperimentResultDataResponse {
	s.StatusCode = &v
	return s
}

func (s *GetExperimentResultDataResponse) SetBody(v *GetExperimentResultDataResponseBody) *GetExperimentResultDataResponse {
	s.Body = v
	return s
}

type GetResourceRequest struct {
	// The cluster ID of Lingjun
	//
	// This parameter is required.
	//
	// example:
	//
	// ehpc-bj-uo8f26cpmo
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s GetResourceRequest) String() string {
	return tea.Prettify(s)
}

func (s GetResourceRequest) GoString() string {
	return s.String()
}

func (s *GetResourceRequest) SetClusterId(v string) *GetResourceRequest {
	s.ClusterId = &v
	return s
}

type GetResourceResponseBody struct {
	// Access denied information
	//
	// example:
	//
	// {}
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// Data
	Data *GetResourceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// Request ID
	//
	// example:
	//
	// 25859897-35C8-5015-8365-7A3CE52F4854
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Total count of the query
	//
	// example:
	//
	// 0
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s GetResourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetResourceResponseBody) GoString() string {
	return s.String()
}

func (s *GetResourceResponseBody) SetAccessDeniedDetail(v string) *GetResourceResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *GetResourceResponseBody) SetData(v *GetResourceResponseBodyData) *GetResourceResponseBody {
	s.Data = v
	return s
}

func (s *GetResourceResponseBody) SetRequestId(v string) *GetResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetResourceResponseBody) SetTotalCount(v int64) *GetResourceResponseBody {
	s.TotalCount = &v
	return s
}

type GetResourceResponseBodyData struct {
	// Cluster description
	//
	// example:
	//
	// test
	ClusterDesc *string `json:"ClusterDesc,omitempty" xml:"ClusterDesc,omitempty"`
	// Cluster ID
	//
	// example:
	//
	// 3123121223
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// Cluster name
	//
	// example:
	//
	// main_cluster
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// Used CPU
	//
	// example:
	//
	// 90
	CpuCoreLimit *int32 `json:"CpuCoreLimit,omitempty" xml:"CpuCoreLimit,omitempty"`
	// Used GPU
	//
	// example:
	//
	// 8
	GpuLimit *int32 `json:"GpuLimit,omitempty" xml:"GpuLimit,omitempty"`
	// Machine type
	MachineType *GetResourceResponseBodyDataMachineType `json:"MachineType,omitempty" xml:"MachineType,omitempty" type:"Struct"`
	// Used memory
	//
	// example:
	//
	// 90
	MaxCpuCore *int32 `json:"MaxCpuCore,omitempty" xml:"MaxCpuCore,omitempty"`
	// Used memory
	//
	// example:
	//
	// 8
	MaxGpu *int32 `json:"MaxGpu,omitempty" xml:"MaxGpu,omitempty"`
	// Used memory
	//
	// example:
	//
	// 500
	MaxMemory *int64 `json:"MaxMemory,omitempty" xml:"MaxMemory,omitempty"`
	// Used memory
	//
	// example:
	//
	// 500
	MemoryLimit *int64 `json:"MemoryLimit,omitempty" xml:"MemoryLimit,omitempty"`
	// Cluster ID
	//
	// example:
	//
	// 189
	ResourceId *int64 `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// List of resource nodes
	ResourceNodes []*GetResourceResponseBodyDataResourceNodes `json:"ResourceNodes,omitempty" xml:"ResourceNodes,omitempty" type:"Repeated"`
	// User authorization parameters
	UserAccessParam *GetResourceResponseBodyDataUserAccessParam `json:"UserAccessParam,omitempty" xml:"UserAccessParam,omitempty" type:"Struct"`
}

func (s GetResourceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetResourceResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetResourceResponseBodyData) SetClusterDesc(v string) *GetResourceResponseBodyData {
	s.ClusterDesc = &v
	return s
}

func (s *GetResourceResponseBodyData) SetClusterId(v string) *GetResourceResponseBodyData {
	s.ClusterId = &v
	return s
}

func (s *GetResourceResponseBodyData) SetClusterName(v string) *GetResourceResponseBodyData {
	s.ClusterName = &v
	return s
}

func (s *GetResourceResponseBodyData) SetCpuCoreLimit(v int32) *GetResourceResponseBodyData {
	s.CpuCoreLimit = &v
	return s
}

func (s *GetResourceResponseBodyData) SetGpuLimit(v int32) *GetResourceResponseBodyData {
	s.GpuLimit = &v
	return s
}

func (s *GetResourceResponseBodyData) SetMachineType(v *GetResourceResponseBodyDataMachineType) *GetResourceResponseBodyData {
	s.MachineType = v
	return s
}

func (s *GetResourceResponseBodyData) SetMaxCpuCore(v int32) *GetResourceResponseBodyData {
	s.MaxCpuCore = &v
	return s
}

func (s *GetResourceResponseBodyData) SetMaxGpu(v int32) *GetResourceResponseBodyData {
	s.MaxGpu = &v
	return s
}

func (s *GetResourceResponseBodyData) SetMaxMemory(v int64) *GetResourceResponseBodyData {
	s.MaxMemory = &v
	return s
}

func (s *GetResourceResponseBodyData) SetMemoryLimit(v int64) *GetResourceResponseBodyData {
	s.MemoryLimit = &v
	return s
}

func (s *GetResourceResponseBodyData) SetResourceId(v int64) *GetResourceResponseBodyData {
	s.ResourceId = &v
	return s
}

func (s *GetResourceResponseBodyData) SetResourceNodes(v []*GetResourceResponseBodyDataResourceNodes) *GetResourceResponseBodyData {
	s.ResourceNodes = v
	return s
}

func (s *GetResourceResponseBodyData) SetUserAccessParam(v *GetResourceResponseBodyDataUserAccessParam) *GetResourceResponseBodyData {
	s.UserAccessParam = v
	return s
}

type GetResourceResponseBodyDataMachineType struct {
	// Number of network bonds
	//
	// example:
	//
	// 5
	BondNum *int32 `json:"BondNum,omitempty" xml:"BondNum,omitempty"`
	// CPU information
	//
	// example:
	//
	// 2x Intel Saphhire Rapid 8469C 48C CPU
	CpuInfo *string `json:"CpuInfo,omitempty" xml:"CpuInfo,omitempty"`
	// Disk information
	//
	// example:
	//
	// 2x 480GB SATA SSD\\n4x 3.84TB NVMe SSD
	DiskInfo *string `json:"DiskInfo,omitempty" xml:"DiskInfo,omitempty"`
	// GPU information
	//
	// example:
	//
	// 8x OAM 810 GPU
	GpuInfo *string `json:"GpuInfo,omitempty" xml:"GpuInfo,omitempty"`
	// Memory information
	//
	// example:
	//
	// 32x 64GB DDR4 4800 Memory
	MemoryInfo *string `json:"MemoryInfo,omitempty" xml:"MemoryInfo,omitempty"`
	// Specification name
	//
	// example:
	//
	// efg2.p8en
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Network information
	//
	// example:
	//
	// 1x 200Gbps Dual Port BF3 DPU for VPC\\n4x 200Gbps Dual Port EIC
	NetworkInfo *string `json:"NetworkInfo,omitempty" xml:"NetworkInfo,omitempty"`
	// Network mode
	//
	// example:
	//
	// 2
	NetworkMode *string `json:"NetworkMode,omitempty" xml:"NetworkMode,omitempty"`
	// Number of nodes
	//
	// example:
	//
	// 1
	NodeCount *int32 `json:"NodeCount,omitempty" xml:"NodeCount,omitempty"`
	// Type
	//
	// example:
	//
	// Private
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetResourceResponseBodyDataMachineType) String() string {
	return tea.Prettify(s)
}

func (s GetResourceResponseBodyDataMachineType) GoString() string {
	return s.String()
}

func (s *GetResourceResponseBodyDataMachineType) SetBondNum(v int32) *GetResourceResponseBodyDataMachineType {
	s.BondNum = &v
	return s
}

func (s *GetResourceResponseBodyDataMachineType) SetCpuInfo(v string) *GetResourceResponseBodyDataMachineType {
	s.CpuInfo = &v
	return s
}

func (s *GetResourceResponseBodyDataMachineType) SetDiskInfo(v string) *GetResourceResponseBodyDataMachineType {
	s.DiskInfo = &v
	return s
}

func (s *GetResourceResponseBodyDataMachineType) SetGpuInfo(v string) *GetResourceResponseBodyDataMachineType {
	s.GpuInfo = &v
	return s
}

func (s *GetResourceResponseBodyDataMachineType) SetMemoryInfo(v string) *GetResourceResponseBodyDataMachineType {
	s.MemoryInfo = &v
	return s
}

func (s *GetResourceResponseBodyDataMachineType) SetName(v string) *GetResourceResponseBodyDataMachineType {
	s.Name = &v
	return s
}

func (s *GetResourceResponseBodyDataMachineType) SetNetworkInfo(v string) *GetResourceResponseBodyDataMachineType {
	s.NetworkInfo = &v
	return s
}

func (s *GetResourceResponseBodyDataMachineType) SetNetworkMode(v string) *GetResourceResponseBodyDataMachineType {
	s.NetworkMode = &v
	return s
}

func (s *GetResourceResponseBodyDataMachineType) SetNodeCount(v int32) *GetResourceResponseBodyDataMachineType {
	s.NodeCount = &v
	return s
}

func (s *GetResourceResponseBodyDataMachineType) SetType(v string) *GetResourceResponseBodyDataMachineType {
	s.Type = &v
	return s
}

type GetResourceResponseBodyDataResourceNodes struct {
	// Node name
	//
	// example:
	//
	// lingj19q90jp66nq-mg2pa0p2l2bipnsi-17
	NodeName *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
}

func (s GetResourceResponseBodyDataResourceNodes) String() string {
	return tea.Prettify(s)
}

func (s GetResourceResponseBodyDataResourceNodes) GoString() string {
	return s.String()
}

func (s *GetResourceResponseBodyDataResourceNodes) SetNodeName(v string) *GetResourceResponseBodyDataResourceNodes {
	s.NodeName = &v
	return s
}

type GetResourceResponseBodyDataUserAccessParam struct {
	// User ID
	//
	// example:
	//
	// dev
	AccessId *string `json:"AccessId,omitempty" xml:"AccessId,omitempty"`
	// User key
	//
	// example:
	//
	// test
	AccessKey *string `json:"AccessKey,omitempty" xml:"AccessKey,omitempty"`
	// Endpoint
	//
	// example:
	//
	// test
	Endpoint *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	// Workspace ID
	//
	// example:
	//
	// test
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s GetResourceResponseBodyDataUserAccessParam) String() string {
	return tea.Prettify(s)
}

func (s GetResourceResponseBodyDataUserAccessParam) GoString() string {
	return s.String()
}

func (s *GetResourceResponseBodyDataUserAccessParam) SetAccessId(v string) *GetResourceResponseBodyDataUserAccessParam {
	s.AccessId = &v
	return s
}

func (s *GetResourceResponseBodyDataUserAccessParam) SetAccessKey(v string) *GetResourceResponseBodyDataUserAccessParam {
	s.AccessKey = &v
	return s
}

func (s *GetResourceResponseBodyDataUserAccessParam) SetEndpoint(v string) *GetResourceResponseBodyDataUserAccessParam {
	s.Endpoint = &v
	return s
}

func (s *GetResourceResponseBodyDataUserAccessParam) SetWorkspaceId(v string) *GetResourceResponseBodyDataUserAccessParam {
	s.WorkspaceId = &v
	return s
}

type GetResourceResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetResourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetResourceResponse) String() string {
	return tea.Prettify(s)
}

func (s GetResourceResponse) GoString() string {
	return s.String()
}

func (s *GetResourceResponse) SetHeaders(v map[string]*string) *GetResourceResponse {
	s.Headers = v
	return s
}

func (s *GetResourceResponse) SetStatusCode(v int32) *GetResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetResourceResponse) SetBody(v *GetResourceResponseBody) *GetResourceResponse {
	s.Body = v
	return s
}

type GetResourcePredictResultRequest struct {
	// Resource ID
	//
	// example:
	//
	// 36
	ResourceId *int64 `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// Template ID
	//
	// This parameter is required.
	//
	// example:
	//
	// 315797
	TemplateId *int64 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s GetResourcePredictResultRequest) String() string {
	return tea.Prettify(s)
}

func (s GetResourcePredictResultRequest) GoString() string {
	return s.String()
}

func (s *GetResourcePredictResultRequest) SetResourceId(v int64) *GetResourcePredictResultRequest {
	s.ResourceId = &v
	return s
}

func (s *GetResourcePredictResultRequest) SetTemplateId(v int64) *GetResourcePredictResultRequest {
	s.TemplateId = &v
	return s
}

type GetResourcePredictResultResponseBody struct {
	// Data
	//
	// example:
	//
	// 2
	Data *int64 `json:"Data,omitempty" xml:"Data,omitempty"`
	// Request ID
	//
	// example:
	//
	// 5514CB39-B7C0-5B89-8534-2DE1E0F2B7AB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// total
	//
	// example:
	//
	// 1
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s GetResourcePredictResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetResourcePredictResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetResourcePredictResultResponseBody) SetData(v int64) *GetResourcePredictResultResponseBody {
	s.Data = &v
	return s
}

func (s *GetResourcePredictResultResponseBody) SetRequestId(v string) *GetResourcePredictResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetResourcePredictResultResponseBody) SetTotalCount(v int64) *GetResourcePredictResultResponseBody {
	s.TotalCount = &v
	return s
}

type GetResourcePredictResultResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetResourcePredictResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetResourcePredictResultResponse) String() string {
	return tea.Prettify(s)
}

func (s GetResourcePredictResultResponse) GoString() string {
	return s.String()
}

func (s *GetResourcePredictResultResponse) SetHeaders(v map[string]*string) *GetResourcePredictResultResponse {
	s.Headers = v
	return s
}

func (s *GetResourcePredictResultResponse) SetStatusCode(v int32) *GetResourcePredictResultResponse {
	s.StatusCode = &v
	return s
}

func (s *GetResourcePredictResultResponse) SetBody(v *GetResourcePredictResultResponseBody) *GetResourcePredictResultResponse {
	s.Body = v
	return s
}

type GetWorkloadRequest struct {
	// Workload ID
	//
	// This parameter is required.
	//
	// example:
	//
	// 13
	WorkloadId *int64 `json:"WorkloadId,omitempty" xml:"WorkloadId,omitempty"`
}

func (s GetWorkloadRequest) String() string {
	return tea.Prettify(s)
}

func (s GetWorkloadRequest) GoString() string {
	return s.String()
}

func (s *GetWorkloadRequest) SetWorkloadId(v int64) *GetWorkloadRequest {
	s.WorkloadId = &v
	return s
}

type GetWorkloadResponseBody struct {
	// Access Denied Information
	//
	// example:
	//
	// {}
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// Data
	Data *GetWorkloadResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// Request ID
	//
	// example:
	//
	// E67E2E4C-2B47-5C55-AA17-1D771E070AEF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// total
	//
	// example:
	//
	// 0
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s GetWorkloadResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetWorkloadResponseBody) GoString() string {
	return s.String()
}

func (s *GetWorkloadResponseBody) SetAccessDeniedDetail(v string) *GetWorkloadResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *GetWorkloadResponseBody) SetData(v *GetWorkloadResponseBodyData) *GetWorkloadResponseBody {
	s.Data = v
	return s
}

func (s *GetWorkloadResponseBody) SetRequestId(v string) *GetWorkloadResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetWorkloadResponseBody) SetTotalCount(v int64) *GetWorkloadResponseBody {
	s.TotalCount = &v
	return s
}

type GetWorkloadResponseBodyData struct {
	// Default CPU Allocation per Worker
	//
	// example:
	//
	// 90
	DefaultCpuPerWorker *int32 `json:"DefaultCpuPerWorker,omitempty" xml:"DefaultCpuPerWorker,omitempty"`
	// Default GPU Allocation per Worker
	//
	// example:
	//
	// 8
	DefaultGpuPerWorker *int32 `json:"DefaultGpuPerWorker,omitempty" xml:"DefaultGpuPerWorker,omitempty"`
	// Default Memory (GB) Allocation per Worker
	//
	// example:
	//
	// 500
	DefaultMemoryPerWorker *int32 `json:"DefaultMemoryPerWorker,omitempty" xml:"DefaultMemoryPerWorker,omitempty"`
	// Default Shared Memory (GB) Allocation
	//
	// example:
	//
	// 500
	DefaultShareMemory *int32 `json:"DefaultShareMemory,omitempty" xml:"DefaultShareMemory,omitempty"`
	// Workload Cluster, e.g., AI, GPU
	//
	// example:
	//
	// AI
	Family *string `json:"Family,omitempty" xml:"Family,omitempty"`
	// Training Job Type
	//
	// example:
	//
	// PyTorchJob
	JobKind *string `json:"JobKind,omitempty" xml:"JobKind,omitempty"`
	// Parameter Settings
	ParamSettings []*GetWorkloadResponseBodyDataParamSettings `json:"ParamSettings,omitempty" xml:"ParamSettings,omitempty" type:"Repeated"`
	// Workload Usage Scenario
	//
	// example:
	//
	// NLP-LLM
	Scene *string `json:"Scene,omitempty" xml:"Scene,omitempty"`
	// Scope Identifier for Workload Usage
	//
	// example:
	//
	// common
	Scope *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
	// Static Configuration
	StaticConfig *GetWorkloadResponseBodyDataStaticConfig `json:"StaticConfig,omitempty" xml:"StaticConfig,omitempty" type:"Struct"`
	// Version ID
	//
	// example:
	//
	// 1
	VersionId *int64 `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
	// Workload Description
	//
	// example:
	//
	// test
	WorkloadDescription *string `json:"WorkloadDescription,omitempty" xml:"WorkloadDescription,omitempty"`
	// Workload ID
	//
	// example:
	//
	// 13
	WorkloadId *int64 `json:"WorkloadId,omitempty" xml:"WorkloadId,omitempty"`
	// Workload Name
	//
	// example:
	//
	// test
	WorkloadName *string `json:"WorkloadName,omitempty" xml:"WorkloadName,omitempty"`
	// Workload Type
	//
	// example:
	//
	// AI
	WorkloadType *string `json:"WorkloadType,omitempty" xml:"WorkloadType,omitempty"`
}

func (s GetWorkloadResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetWorkloadResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetWorkloadResponseBodyData) SetDefaultCpuPerWorker(v int32) *GetWorkloadResponseBodyData {
	s.DefaultCpuPerWorker = &v
	return s
}

func (s *GetWorkloadResponseBodyData) SetDefaultGpuPerWorker(v int32) *GetWorkloadResponseBodyData {
	s.DefaultGpuPerWorker = &v
	return s
}

func (s *GetWorkloadResponseBodyData) SetDefaultMemoryPerWorker(v int32) *GetWorkloadResponseBodyData {
	s.DefaultMemoryPerWorker = &v
	return s
}

func (s *GetWorkloadResponseBodyData) SetDefaultShareMemory(v int32) *GetWorkloadResponseBodyData {
	s.DefaultShareMemory = &v
	return s
}

func (s *GetWorkloadResponseBodyData) SetFamily(v string) *GetWorkloadResponseBodyData {
	s.Family = &v
	return s
}

func (s *GetWorkloadResponseBodyData) SetJobKind(v string) *GetWorkloadResponseBodyData {
	s.JobKind = &v
	return s
}

func (s *GetWorkloadResponseBodyData) SetParamSettings(v []*GetWorkloadResponseBodyDataParamSettings) *GetWorkloadResponseBodyData {
	s.ParamSettings = v
	return s
}

func (s *GetWorkloadResponseBodyData) SetScene(v string) *GetWorkloadResponseBodyData {
	s.Scene = &v
	return s
}

func (s *GetWorkloadResponseBodyData) SetScope(v string) *GetWorkloadResponseBodyData {
	s.Scope = &v
	return s
}

func (s *GetWorkloadResponseBodyData) SetStaticConfig(v *GetWorkloadResponseBodyDataStaticConfig) *GetWorkloadResponseBodyData {
	s.StaticConfig = v
	return s
}

func (s *GetWorkloadResponseBodyData) SetVersionId(v int64) *GetWorkloadResponseBodyData {
	s.VersionId = &v
	return s
}

func (s *GetWorkloadResponseBodyData) SetWorkloadDescription(v string) *GetWorkloadResponseBodyData {
	s.WorkloadDescription = &v
	return s
}

func (s *GetWorkloadResponseBodyData) SetWorkloadId(v int64) *GetWorkloadResponseBodyData {
	s.WorkloadId = &v
	return s
}

func (s *GetWorkloadResponseBodyData) SetWorkloadName(v string) *GetWorkloadResponseBodyData {
	s.WorkloadName = &v
	return s
}

func (s *GetWorkloadResponseBodyData) SetWorkloadType(v string) *GetWorkloadResponseBodyData {
	s.WorkloadType = &v
	return s
}

type GetWorkloadResponseBodyDataParamSettings struct {
	// Default Parameter Value
	//
	// example:
	//
	// 100
	DefaultValue *string `json:"DefaultValue,omitempty" xml:"DefaultValue,omitempty"`
	// Parameter Description
	//
	// example:
	//
	// number
	ParamDesc *string `json:"ParamDesc,omitempty" xml:"ParamDesc,omitempty"`
	// Parameter Name
	//
	// example:
	//
	// ITERATION
	ParamName *string `json:"ParamName,omitempty" xml:"ParamName,omitempty"`
	// Parameter Regular Expression
	//
	// example:
	//
	// [0-9]+
	ParamRegex *string `json:"ParamRegex,omitempty" xml:"ParamRegex,omitempty"`
	// Parameter type
	//
	// example:
	//
	// number
	ParamType *string `json:"ParamType,omitempty" xml:"ParamType,omitempty"`
	// Parameter Value
	//
	// example:
	//
	// 100
	ParamValue *string `json:"ParamValue,omitempty" xml:"ParamValue,omitempty"`
}

func (s GetWorkloadResponseBodyDataParamSettings) String() string {
	return tea.Prettify(s)
}

func (s GetWorkloadResponseBodyDataParamSettings) GoString() string {
	return s.String()
}

func (s *GetWorkloadResponseBodyDataParamSettings) SetDefaultValue(v string) *GetWorkloadResponseBodyDataParamSettings {
	s.DefaultValue = &v
	return s
}

func (s *GetWorkloadResponseBodyDataParamSettings) SetParamDesc(v string) *GetWorkloadResponseBodyDataParamSettings {
	s.ParamDesc = &v
	return s
}

func (s *GetWorkloadResponseBodyDataParamSettings) SetParamName(v string) *GetWorkloadResponseBodyDataParamSettings {
	s.ParamName = &v
	return s
}

func (s *GetWorkloadResponseBodyDataParamSettings) SetParamRegex(v string) *GetWorkloadResponseBodyDataParamSettings {
	s.ParamRegex = &v
	return s
}

func (s *GetWorkloadResponseBodyDataParamSettings) SetParamType(v string) *GetWorkloadResponseBodyDataParamSettings {
	s.ParamType = &v
	return s
}

func (s *GetWorkloadResponseBodyDataParamSettings) SetParamValue(v string) *GetWorkloadResponseBodyDataParamSettings {
	s.ParamValue = &v
	return s
}

type GetWorkloadResponseBodyDataStaticConfig struct {
	// Framework
	//
	// example:
	//
	// PyTorch
	FrameWork *string `json:"FrameWork,omitempty" xml:"FrameWork,omitempty"`
	// Operating System
	//
	// example:
	//
	// linux
	Os *string `json:"Os,omitempty" xml:"Os,omitempty"`
	// Parameter Volume
	//
	// example:
	//
	// 7B
	Parameters *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// Software Stack
	//
	// example:
	//
	// python
	SoftwareStack *string `json:"SoftwareStack,omitempty" xml:"SoftwareStack,omitempty"`
}

func (s GetWorkloadResponseBodyDataStaticConfig) String() string {
	return tea.Prettify(s)
}

func (s GetWorkloadResponseBodyDataStaticConfig) GoString() string {
	return s.String()
}

func (s *GetWorkloadResponseBodyDataStaticConfig) SetFrameWork(v string) *GetWorkloadResponseBodyDataStaticConfig {
	s.FrameWork = &v
	return s
}

func (s *GetWorkloadResponseBodyDataStaticConfig) SetOs(v string) *GetWorkloadResponseBodyDataStaticConfig {
	s.Os = &v
	return s
}

func (s *GetWorkloadResponseBodyDataStaticConfig) SetParameters(v string) *GetWorkloadResponseBodyDataStaticConfig {
	s.Parameters = &v
	return s
}

func (s *GetWorkloadResponseBodyDataStaticConfig) SetSoftwareStack(v string) *GetWorkloadResponseBodyDataStaticConfig {
	s.SoftwareStack = &v
	return s
}

type GetWorkloadResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetWorkloadResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetWorkloadResponse) String() string {
	return tea.Prettify(s)
}

func (s GetWorkloadResponse) GoString() string {
	return s.String()
}

func (s *GetWorkloadResponse) SetHeaders(v map[string]*string) *GetWorkloadResponse {
	s.Headers = v
	return s
}

func (s *GetWorkloadResponse) SetStatusCode(v int32) *GetWorkloadResponse {
	s.StatusCode = &v
	return s
}

func (s *GetWorkloadResponse) SetBody(v *GetWorkloadResponseBody) *GetWorkloadResponse {
	s.Body = v
	return s
}

type ListExperimentPlanTemplatesRequest struct {
	// The sharing level of the template, default is private, options are public and private.
	//
	// example:
	//
	// private
	PrivacyLevel *string `json:"PrivacyLevel,omitempty" xml:"PrivacyLevel,omitempty"`
}

func (s ListExperimentPlanTemplatesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListExperimentPlanTemplatesRequest) GoString() string {
	return s.String()
}

func (s *ListExperimentPlanTemplatesRequest) SetPrivacyLevel(v string) *ListExperimentPlanTemplatesRequest {
	s.PrivacyLevel = &v
	return s
}

type ListExperimentPlanTemplatesResponseBody struct {
	// Data
	Data []*ListExperimentPlanTemplatesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// Request ID
	//
	// example:
	//
	// AAE4AFED-7AE6-52FA-80B6-448E63760552
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Total
	//
	// example:
	//
	// 0
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListExperimentPlanTemplatesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListExperimentPlanTemplatesResponseBody) GoString() string {
	return s.String()
}

func (s *ListExperimentPlanTemplatesResponseBody) SetData(v []*ListExperimentPlanTemplatesResponseBodyData) *ListExperimentPlanTemplatesResponseBody {
	s.Data = v
	return s
}

func (s *ListExperimentPlanTemplatesResponseBody) SetRequestId(v string) *ListExperimentPlanTemplatesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListExperimentPlanTemplatesResponseBody) SetTotalCount(v int64) *ListExperimentPlanTemplatesResponseBody {
	s.TotalCount = &v
	return s
}

type ListExperimentPlanTemplatesResponseBodyData struct {
	// Creation time
	//
	// example:
	//
	// 2024-11-29 02:16:35
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// Primary account UID
	//
	// example:
	//
	// 178154411231232
	CreatorUid *int64 `json:"CreatorUid,omitempty" xml:"CreatorUid,omitempty"`
	// Whether it is deleted
	//
	// example:
	//
	// 0
	IsDelete *int32 `json:"IsDelete,omitempty" xml:"IsDelete,omitempty"`
	// Privacy level
	//
	// example:
	//
	// private
	PrivacyLevel *string `json:"PrivacyLevel,omitempty" xml:"PrivacyLevel,omitempty"`
	// The template code.
	//
	// example:
	//
	// 475315534
	TemplateCode *int64 `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
	// Template description
	//
	// example:
	//
	// test
	TemplateDescription *string `json:"TemplateDescription,omitempty" xml:"TemplateDescription,omitempty"`
	// Template ID
	//
	// example:
	//
	// 17815441
	TemplateId *int64 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// Template name
	//
	// example:
	//
	// test
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// Template pipeline
	TemplatePipelineParam []*ListExperimentPlanTemplatesResponseBodyDataTemplatePipelineParam `json:"TemplatePipelineParam,omitempty" xml:"TemplatePipelineParam,omitempty" type:"Repeated"`
	// Update time
	//
	// example:
	//
	// 2024-11-29 02:16:35
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// Version ID
	//
	// example:
	//
	// 1
	VersionId *int64 `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
}

func (s ListExperimentPlanTemplatesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListExperimentPlanTemplatesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListExperimentPlanTemplatesResponseBodyData) SetCreateTime(v string) *ListExperimentPlanTemplatesResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *ListExperimentPlanTemplatesResponseBodyData) SetCreatorUid(v int64) *ListExperimentPlanTemplatesResponseBodyData {
	s.CreatorUid = &v
	return s
}

func (s *ListExperimentPlanTemplatesResponseBodyData) SetIsDelete(v int32) *ListExperimentPlanTemplatesResponseBodyData {
	s.IsDelete = &v
	return s
}

func (s *ListExperimentPlanTemplatesResponseBodyData) SetPrivacyLevel(v string) *ListExperimentPlanTemplatesResponseBodyData {
	s.PrivacyLevel = &v
	return s
}

func (s *ListExperimentPlanTemplatesResponseBodyData) SetTemplateCode(v int64) *ListExperimentPlanTemplatesResponseBodyData {
	s.TemplateCode = &v
	return s
}

func (s *ListExperimentPlanTemplatesResponseBodyData) SetTemplateDescription(v string) *ListExperimentPlanTemplatesResponseBodyData {
	s.TemplateDescription = &v
	return s
}

func (s *ListExperimentPlanTemplatesResponseBodyData) SetTemplateId(v int64) *ListExperimentPlanTemplatesResponseBodyData {
	s.TemplateId = &v
	return s
}

func (s *ListExperimentPlanTemplatesResponseBodyData) SetTemplateName(v string) *ListExperimentPlanTemplatesResponseBodyData {
	s.TemplateName = &v
	return s
}

func (s *ListExperimentPlanTemplatesResponseBodyData) SetTemplatePipelineParam(v []*ListExperimentPlanTemplatesResponseBodyDataTemplatePipelineParam) *ListExperimentPlanTemplatesResponseBodyData {
	s.TemplatePipelineParam = v
	return s
}

func (s *ListExperimentPlanTemplatesResponseBodyData) SetUpdateTime(v string) *ListExperimentPlanTemplatesResponseBodyData {
	s.UpdateTime = &v
	return s
}

func (s *ListExperimentPlanTemplatesResponseBodyData) SetVersionId(v int64) *ListExperimentPlanTemplatesResponseBodyData {
	s.VersionId = &v
	return s
}

type ListExperimentPlanTemplatesResponseBodyDataTemplatePipelineParam struct {
	// Configured environment parameters
	EnvParams *ListExperimentPlanTemplatesResponseBodyDataTemplatePipelineParamEnvParams `json:"EnvParams,omitempty" xml:"EnvParams,omitempty" type:"Struct"`
	// Node sequence number
	//
	// example:
	//
	// 1
	PipelineOrder *int32 `json:"PipelineOrder,omitempty" xml:"PipelineOrder,omitempty"`
	// Usage scenario, e.g., "baseline"
	//
	// example:
	//
	// baseline
	Scene *string `json:"Scene,omitempty" xml:"Scene,omitempty"`
	// Configured workload parameters
	SettingParams map[string]*string `json:"SettingParams,omitempty" xml:"SettingParams,omitempty"`
	// Workload ID
	//
	// example:
	//
	// 13
	WorkloadId *int64 `json:"WorkloadId,omitempty" xml:"WorkloadId,omitempty"`
	// Workload name
	//
	// example:
	//
	// test
	WorkloadName *string `json:"WorkloadName,omitempty" xml:"WorkloadName,omitempty"`
}

func (s ListExperimentPlanTemplatesResponseBodyDataTemplatePipelineParam) String() string {
	return tea.Prettify(s)
}

func (s ListExperimentPlanTemplatesResponseBodyDataTemplatePipelineParam) GoString() string {
	return s.String()
}

func (s *ListExperimentPlanTemplatesResponseBodyDataTemplatePipelineParam) SetEnvParams(v *ListExperimentPlanTemplatesResponseBodyDataTemplatePipelineParamEnvParams) *ListExperimentPlanTemplatesResponseBodyDataTemplatePipelineParam {
	s.EnvParams = v
	return s
}

func (s *ListExperimentPlanTemplatesResponseBodyDataTemplatePipelineParam) SetPipelineOrder(v int32) *ListExperimentPlanTemplatesResponseBodyDataTemplatePipelineParam {
	s.PipelineOrder = &v
	return s
}

func (s *ListExperimentPlanTemplatesResponseBodyDataTemplatePipelineParam) SetScene(v string) *ListExperimentPlanTemplatesResponseBodyDataTemplatePipelineParam {
	s.Scene = &v
	return s
}

func (s *ListExperimentPlanTemplatesResponseBodyDataTemplatePipelineParam) SetSettingParams(v map[string]*string) *ListExperimentPlanTemplatesResponseBodyDataTemplatePipelineParam {
	s.SettingParams = v
	return s
}

func (s *ListExperimentPlanTemplatesResponseBodyDataTemplatePipelineParam) SetWorkloadId(v int64) *ListExperimentPlanTemplatesResponseBodyDataTemplatePipelineParam {
	s.WorkloadId = &v
	return s
}

func (s *ListExperimentPlanTemplatesResponseBodyDataTemplatePipelineParam) SetWorkloadName(v string) *ListExperimentPlanTemplatesResponseBodyDataTemplatePipelineParam {
	s.WorkloadName = &v
	return s
}

type ListExperimentPlanTemplatesResponseBodyDataTemplatePipelineParamEnvParams struct {
	// CPU allocation
	//
	// example:
	//
	// 90
	CpuPerWorker *int32 `json:"CpuPerWorker,omitempty" xml:"CpuPerWorker,omitempty"`
	// Cuda Version
	//
	// example:
	//
	// 1.0.0
	CudaVersion *string `json:"CudaVersion,omitempty" xml:"CudaVersion,omitempty"`
	// The version of the GPU driver.
	//
	// example:
	//
	// 1.0.0
	GpuDriverVersion *string `json:"GpuDriverVersion,omitempty" xml:"GpuDriverVersion,omitempty"`
	// GPU allocation
	//
	// example:
	//
	// 8
	GpuPerWorker *int32 `json:"GpuPerWorker,omitempty" xml:"GpuPerWorker,omitempty"`
	// Allocated memory in GB
	//
	// example:
	//
	// 500
	MemoryPerWorker *int32 `json:"MemoryPerWorker,omitempty" xml:"MemoryPerWorker,omitempty"`
	// NCCL Version
	//
	// example:
	//
	// 1.0.0
	NCCLVersion *string `json:"NCCLVersion,omitempty" xml:"NCCLVersion,omitempty"`
	// PyTorch Version
	//
	// example:
	//
	// 1.0.0
	PyTorchVersion *string `json:"PyTorchVersion,omitempty" xml:"PyTorchVersion,omitempty"`
	// Allocated shared memory in GB
	//
	// example:
	//
	// 500
	ShareMemory *int32 `json:"ShareMemory,omitempty" xml:"ShareMemory,omitempty"`
	// Number of nodes
	//
	// example:
	//
	// 1
	WorkerNum *int32 `json:"WorkerNum,omitempty" xml:"WorkerNum,omitempty"`
}

func (s ListExperimentPlanTemplatesResponseBodyDataTemplatePipelineParamEnvParams) String() string {
	return tea.Prettify(s)
}

func (s ListExperimentPlanTemplatesResponseBodyDataTemplatePipelineParamEnvParams) GoString() string {
	return s.String()
}

func (s *ListExperimentPlanTemplatesResponseBodyDataTemplatePipelineParamEnvParams) SetCpuPerWorker(v int32) *ListExperimentPlanTemplatesResponseBodyDataTemplatePipelineParamEnvParams {
	s.CpuPerWorker = &v
	return s
}

func (s *ListExperimentPlanTemplatesResponseBodyDataTemplatePipelineParamEnvParams) SetCudaVersion(v string) *ListExperimentPlanTemplatesResponseBodyDataTemplatePipelineParamEnvParams {
	s.CudaVersion = &v
	return s
}

func (s *ListExperimentPlanTemplatesResponseBodyDataTemplatePipelineParamEnvParams) SetGpuDriverVersion(v string) *ListExperimentPlanTemplatesResponseBodyDataTemplatePipelineParamEnvParams {
	s.GpuDriverVersion = &v
	return s
}

func (s *ListExperimentPlanTemplatesResponseBodyDataTemplatePipelineParamEnvParams) SetGpuPerWorker(v int32) *ListExperimentPlanTemplatesResponseBodyDataTemplatePipelineParamEnvParams {
	s.GpuPerWorker = &v
	return s
}

func (s *ListExperimentPlanTemplatesResponseBodyDataTemplatePipelineParamEnvParams) SetMemoryPerWorker(v int32) *ListExperimentPlanTemplatesResponseBodyDataTemplatePipelineParamEnvParams {
	s.MemoryPerWorker = &v
	return s
}

func (s *ListExperimentPlanTemplatesResponseBodyDataTemplatePipelineParamEnvParams) SetNCCLVersion(v string) *ListExperimentPlanTemplatesResponseBodyDataTemplatePipelineParamEnvParams {
	s.NCCLVersion = &v
	return s
}

func (s *ListExperimentPlanTemplatesResponseBodyDataTemplatePipelineParamEnvParams) SetPyTorchVersion(v string) *ListExperimentPlanTemplatesResponseBodyDataTemplatePipelineParamEnvParams {
	s.PyTorchVersion = &v
	return s
}

func (s *ListExperimentPlanTemplatesResponseBodyDataTemplatePipelineParamEnvParams) SetShareMemory(v int32) *ListExperimentPlanTemplatesResponseBodyDataTemplatePipelineParamEnvParams {
	s.ShareMemory = &v
	return s
}

func (s *ListExperimentPlanTemplatesResponseBodyDataTemplatePipelineParamEnvParams) SetWorkerNum(v int32) *ListExperimentPlanTemplatesResponseBodyDataTemplatePipelineParamEnvParams {
	s.WorkerNum = &v
	return s
}

type ListExperimentPlanTemplatesResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListExperimentPlanTemplatesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListExperimentPlanTemplatesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListExperimentPlanTemplatesResponse) GoString() string {
	return s.String()
}

func (s *ListExperimentPlanTemplatesResponse) SetHeaders(v map[string]*string) *ListExperimentPlanTemplatesResponse {
	s.Headers = v
	return s
}

func (s *ListExperimentPlanTemplatesResponse) SetStatusCode(v int32) *ListExperimentPlanTemplatesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListExperimentPlanTemplatesResponse) SetBody(v *ListExperimentPlanTemplatesResponseBody) *ListExperimentPlanTemplatesResponse {
	s.Body = v
	return s
}

type ListExperimentPlansRequest struct {
	// Creation Time Sorting Rule
	//
	// example:
	//
	// desc
	CreatTimeOrder *string `json:"CreatTimeOrder,omitempty" xml:"CreatTimeOrder,omitempty"`
	// End Time Sorting Rule
	//
	// example:
	//
	// desc
	EndTimeOrder *string `json:"EndTimeOrder,omitempty" xml:"EndTimeOrder,omitempty"`
	// Page Number
	//
	// example:
	//
	// 1
	Page *int32 `json:"Page,omitempty" xml:"Page,omitempty"`
	// Execution Status
	PlanTaskStatus []*string `json:"PlanTaskStatus,omitempty" xml:"PlanTaskStatus,omitempty" type:"Repeated"`
	// Resource Group ID
	//
	// example:
	//
	// rg-aekzij65sf2rr5i
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// Resource ID
	//
	// example:
	//
	// 189
	ResourceId *int64 `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// Resource
	ResourceName []*string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty" type:"Repeated"`
	// Number of Items
	//
	// example:
	//
	// 100
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
	// Start Time Sorting Rule
	//
	// example:
	//
	// desc
	StartTimeOrder *string `json:"StartTimeOrder,omitempty" xml:"StartTimeOrder,omitempty"`
	// The tags.
	Tag []*ListExperimentPlansRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// Template Id
	//
	// example:
	//
	// 96
	TemplateId *int64 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s ListExperimentPlansRequest) String() string {
	return tea.Prettify(s)
}

func (s ListExperimentPlansRequest) GoString() string {
	return s.String()
}

func (s *ListExperimentPlansRequest) SetCreatTimeOrder(v string) *ListExperimentPlansRequest {
	s.CreatTimeOrder = &v
	return s
}

func (s *ListExperimentPlansRequest) SetEndTimeOrder(v string) *ListExperimentPlansRequest {
	s.EndTimeOrder = &v
	return s
}

func (s *ListExperimentPlansRequest) SetPage(v int32) *ListExperimentPlansRequest {
	s.Page = &v
	return s
}

func (s *ListExperimentPlansRequest) SetPlanTaskStatus(v []*string) *ListExperimentPlansRequest {
	s.PlanTaskStatus = v
	return s
}

func (s *ListExperimentPlansRequest) SetResourceGroupId(v string) *ListExperimentPlansRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListExperimentPlansRequest) SetResourceId(v int64) *ListExperimentPlansRequest {
	s.ResourceId = &v
	return s
}

func (s *ListExperimentPlansRequest) SetResourceName(v []*string) *ListExperimentPlansRequest {
	s.ResourceName = v
	return s
}

func (s *ListExperimentPlansRequest) SetSize(v int32) *ListExperimentPlansRequest {
	s.Size = &v
	return s
}

func (s *ListExperimentPlansRequest) SetStartTimeOrder(v string) *ListExperimentPlansRequest {
	s.StartTimeOrder = &v
	return s
}

func (s *ListExperimentPlansRequest) SetTag(v []*ListExperimentPlansRequestTag) *ListExperimentPlansRequest {
	s.Tag = v
	return s
}

func (s *ListExperimentPlansRequest) SetTemplateId(v int64) *ListExperimentPlansRequest {
	s.TemplateId = &v
	return s
}

type ListExperimentPlansRequestTag struct {
	// The tag key.
	//
	// example:
	//
	// owner
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// Tag value
	//
	// example:
	//
	// test
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListExperimentPlansRequestTag) String() string {
	return tea.Prettify(s)
}

func (s ListExperimentPlansRequestTag) GoString() string {
	return s.String()
}

func (s *ListExperimentPlansRequestTag) SetKey(v string) *ListExperimentPlansRequestTag {
	s.Key = &v
	return s
}

func (s *ListExperimentPlansRequestTag) SetValue(v string) *ListExperimentPlansRequestTag {
	s.Value = &v
	return s
}

type ListExperimentPlansShrinkRequest struct {
	// Creation Time Sorting Rule
	//
	// example:
	//
	// desc
	CreatTimeOrder *string `json:"CreatTimeOrder,omitempty" xml:"CreatTimeOrder,omitempty"`
	// End Time Sorting Rule
	//
	// example:
	//
	// desc
	EndTimeOrder *string `json:"EndTimeOrder,omitempty" xml:"EndTimeOrder,omitempty"`
	// Page Number
	//
	// example:
	//
	// 1
	Page *int32 `json:"Page,omitempty" xml:"Page,omitempty"`
	// Execution Status
	PlanTaskStatusShrink *string `json:"PlanTaskStatus,omitempty" xml:"PlanTaskStatus,omitempty"`
	// Resource Group ID
	//
	// example:
	//
	// rg-aekzij65sf2rr5i
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// Resource ID
	//
	// example:
	//
	// 189
	ResourceId *int64 `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// Resource
	ResourceNameShrink *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	// Number of Items
	//
	// example:
	//
	// 100
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
	// Start Time Sorting Rule
	//
	// example:
	//
	// desc
	StartTimeOrder *string `json:"StartTimeOrder,omitempty" xml:"StartTimeOrder,omitempty"`
	// The tags.
	TagShrink *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	// Template Id
	//
	// example:
	//
	// 96
	TemplateId *int64 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s ListExperimentPlansShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListExperimentPlansShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListExperimentPlansShrinkRequest) SetCreatTimeOrder(v string) *ListExperimentPlansShrinkRequest {
	s.CreatTimeOrder = &v
	return s
}

func (s *ListExperimentPlansShrinkRequest) SetEndTimeOrder(v string) *ListExperimentPlansShrinkRequest {
	s.EndTimeOrder = &v
	return s
}

func (s *ListExperimentPlansShrinkRequest) SetPage(v int32) *ListExperimentPlansShrinkRequest {
	s.Page = &v
	return s
}

func (s *ListExperimentPlansShrinkRequest) SetPlanTaskStatusShrink(v string) *ListExperimentPlansShrinkRequest {
	s.PlanTaskStatusShrink = &v
	return s
}

func (s *ListExperimentPlansShrinkRequest) SetResourceGroupId(v string) *ListExperimentPlansShrinkRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListExperimentPlansShrinkRequest) SetResourceId(v int64) *ListExperimentPlansShrinkRequest {
	s.ResourceId = &v
	return s
}

func (s *ListExperimentPlansShrinkRequest) SetResourceNameShrink(v string) *ListExperimentPlansShrinkRequest {
	s.ResourceNameShrink = &v
	return s
}

func (s *ListExperimentPlansShrinkRequest) SetSize(v int32) *ListExperimentPlansShrinkRequest {
	s.Size = &v
	return s
}

func (s *ListExperimentPlansShrinkRequest) SetStartTimeOrder(v string) *ListExperimentPlansShrinkRequest {
	s.StartTimeOrder = &v
	return s
}

func (s *ListExperimentPlansShrinkRequest) SetTagShrink(v string) *ListExperimentPlansShrinkRequest {
	s.TagShrink = &v
	return s
}

func (s *ListExperimentPlansShrinkRequest) SetTemplateId(v int64) *ListExperimentPlansShrinkRequest {
	s.TemplateId = &v
	return s
}

type ListExperimentPlansResponseBody struct {
	// Access Denied Detail
	//
	// example:
	//
	// {}
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// Data
	Data []*ListExperimentPlansResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// Request ID
	//
	// example:
	//
	// FA9F1DE7-103B-5C18-AE9E-EB2BFF11C685
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Total
	//
	// example:
	//
	// 0
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListExperimentPlansResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListExperimentPlansResponseBody) GoString() string {
	return s.String()
}

func (s *ListExperimentPlansResponseBody) SetAccessDeniedDetail(v string) *ListExperimentPlansResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ListExperimentPlansResponseBody) SetData(v []*ListExperimentPlansResponseBodyData) *ListExperimentPlansResponseBody {
	s.Data = v
	return s
}

func (s *ListExperimentPlansResponseBody) SetRequestId(v string) *ListExperimentPlansResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListExperimentPlansResponseBody) SetTotalCount(v int64) *ListExperimentPlansResponseBody {
	s.TotalCount = &v
	return s
}

type ListExperimentPlansResponseBodyData struct {
	// Creation Time
	//
	// example:
	//
	// 2024-07-08 10:12:42
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// End Time
	//
	// example:
	//
	// 2024-07-08 10:22:42
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// Plan ID
	//
	// example:
	//
	// 189
	PlanId *int64 `json:"PlanId,omitempty" xml:"PlanId,omitempty"`
	// Test Plan Execution Status
	PlanTaskStatus map[string]*int32 `json:"PlanTaskStatus,omitempty" xml:"PlanTaskStatus,omitempty"`
	// Resource Group ID
	//
	// example:
	//
	// rg-aek5behqmwbfhuy
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The resource ID.
	//
	// example:
	//
	// 189
	ResourceId *int64 `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// Associated Resource Name
	//
	// example:
	//
	// q_ecs_xpec_postpay_c
	ResourceName *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	// Start Time
	//
	// example:
	//
	// 2024-07-08 10:12:42
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The tag.
	Tags []*ListExperimentPlansResponseBodyDataTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// Associated Test Plan Template ID
	//
	// example:
	//
	// 6
	TemplateId *int64 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// Associated Test Plan Template Name
	//
	// example:
	//
	// test
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// Update Time
	//
	// example:
	//
	// 2024-07-08 10:12:42
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListExperimentPlansResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListExperimentPlansResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListExperimentPlansResponseBodyData) SetCreateTime(v string) *ListExperimentPlansResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *ListExperimentPlansResponseBodyData) SetEndTime(v string) *ListExperimentPlansResponseBodyData {
	s.EndTime = &v
	return s
}

func (s *ListExperimentPlansResponseBodyData) SetPlanId(v int64) *ListExperimentPlansResponseBodyData {
	s.PlanId = &v
	return s
}

func (s *ListExperimentPlansResponseBodyData) SetPlanTaskStatus(v map[string]*int32) *ListExperimentPlansResponseBodyData {
	s.PlanTaskStatus = v
	return s
}

func (s *ListExperimentPlansResponseBodyData) SetResourceGroupId(v string) *ListExperimentPlansResponseBodyData {
	s.ResourceGroupId = &v
	return s
}

func (s *ListExperimentPlansResponseBodyData) SetResourceId(v int64) *ListExperimentPlansResponseBodyData {
	s.ResourceId = &v
	return s
}

func (s *ListExperimentPlansResponseBodyData) SetResourceName(v string) *ListExperimentPlansResponseBodyData {
	s.ResourceName = &v
	return s
}

func (s *ListExperimentPlansResponseBodyData) SetStartTime(v string) *ListExperimentPlansResponseBodyData {
	s.StartTime = &v
	return s
}

func (s *ListExperimentPlansResponseBodyData) SetTags(v []*ListExperimentPlansResponseBodyDataTags) *ListExperimentPlansResponseBodyData {
	s.Tags = v
	return s
}

func (s *ListExperimentPlansResponseBodyData) SetTemplateId(v int64) *ListExperimentPlansResponseBodyData {
	s.TemplateId = &v
	return s
}

func (s *ListExperimentPlansResponseBodyData) SetTemplateName(v string) *ListExperimentPlansResponseBodyData {
	s.TemplateName = &v
	return s
}

func (s *ListExperimentPlansResponseBodyData) SetUpdateTime(v string) *ListExperimentPlansResponseBodyData {
	s.UpdateTime = &v
	return s
}

type ListExperimentPlansResponseBodyDataTags struct {
	// The tag key.
	//
	// example:
	//
	// owner
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The tag value.
	//
	// example:
	//
	// test
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s ListExperimentPlansResponseBodyDataTags) String() string {
	return tea.Prettify(s)
}

func (s ListExperimentPlansResponseBodyDataTags) GoString() string {
	return s.String()
}

func (s *ListExperimentPlansResponseBodyDataTags) SetTagKey(v string) *ListExperimentPlansResponseBodyDataTags {
	s.TagKey = &v
	return s
}

func (s *ListExperimentPlansResponseBodyDataTags) SetTagValue(v string) *ListExperimentPlansResponseBodyDataTags {
	s.TagValue = &v
	return s
}

type ListExperimentPlansResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListExperimentPlansResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListExperimentPlansResponse) String() string {
	return tea.Prettify(s)
}

func (s ListExperimentPlansResponse) GoString() string {
	return s.String()
}

func (s *ListExperimentPlansResponse) SetHeaders(v map[string]*string) *ListExperimentPlansResponse {
	s.Headers = v
	return s
}

func (s *ListExperimentPlansResponse) SetStatusCode(v int32) *ListExperimentPlansResponse {
	s.StatusCode = &v
	return s
}

func (s *ListExperimentPlansResponse) SetBody(v *ListExperimentPlansResponseBody) *ListExperimentPlansResponse {
	s.Body = v
	return s
}

type ListExperimentsRequest struct {
	// Order
	//
	// example:
	//
	// 1
	Order *int32 `json:"Order,omitempty" xml:"Order,omitempty"`
	// Plan ID
	//
	// example:
	//
	// 189
	PlanId *int64 `json:"PlanId,omitempty" xml:"PlanId,omitempty"`
	// 资源组id
	//
	// example:
	//
	// rg-uo8f26cpmo
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s ListExperimentsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListExperimentsRequest) GoString() string {
	return s.String()
}

func (s *ListExperimentsRequest) SetOrder(v int32) *ListExperimentsRequest {
	s.Order = &v
	return s
}

func (s *ListExperimentsRequest) SetPlanId(v int64) *ListExperimentsRequest {
	s.PlanId = &v
	return s
}

func (s *ListExperimentsRequest) SetResourceGroupId(v string) *ListExperimentsRequest {
	s.ResourceGroupId = &v
	return s
}

type ListExperimentsResponseBody struct {
	// Access denied information
	//
	// example:
	//
	// {}
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// Data
	Data []*ListExperimentsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// Request ID
	//
	// example:
	//
	// 5514CB39-B7C0-5B89-8534-2DE1E0F2B7AB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Total
	//
	// example:
	//
	// 0
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListExperimentsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListExperimentsResponseBody) GoString() string {
	return s.String()
}

func (s *ListExperimentsResponseBody) SetAccessDeniedDetail(v string) *ListExperimentsResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ListExperimentsResponseBody) SetData(v []*ListExperimentsResponseBodyData) *ListExperimentsResponseBody {
	s.Data = v
	return s
}

func (s *ListExperimentsResponseBody) SetRequestId(v string) *ListExperimentsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListExperimentsResponseBody) SetTotalCount(v int64) *ListExperimentsResponseBody {
	s.TotalCount = &v
	return s
}

type ListExperimentsResponseBodyData struct {
	// Creation time
	//
	// example:
	//
	// 2024-10-22 10:18:10
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// Task end time
	//
	// example:
	//
	// 2024-10-22 10:28:10
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// Environment parameters in operation
	EnvParams *ListExperimentsResponseBodyDataEnvParams `json:"EnvParams,omitempty" xml:"EnvParams,omitempty" type:"Struct"`
	// Experiment ID
	//
	// example:
	//
	// 1684537476910997506
	ExperimentId *int64 `json:"ExperimentId,omitempty" xml:"ExperimentId,omitempty"`
	// Experiment name
	//
	// example:
	//
	// test
	ExperimentName *string `json:"ExperimentName,omitempty" xml:"ExperimentName,omitempty"`
	// Experiment type
	//
	// example:
	//
	// AI
	ExperimentType *string `json:"ExperimentType,omitempty" xml:"ExperimentType,omitempty"`
	// Parsed load parameters
	GetParams map[string]*string `json:"GetParams,omitempty" xml:"GetParams,omitempty"`
	// Resource name
	//
	// example:
	//
	// ecs.r8y.4xlarge
	ResourceName *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	// Task results
	Results *ListExperimentsResponseBodyDataResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Struct"`
	// Load parameters in operation
	SetParams map[string]*string `json:"SetParams,omitempty" xml:"SetParams,omitempty"`
	// Task start time
	//
	// example:
	//
	// 2024-10-22 10:18:10
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// Status
	//
	// example:
	//
	// RUNNING
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Update time
	//
	// example:
	//
	// 2024-10-22 10:18:10
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// Workload name
	//
	// example:
	//
	// test
	WorkloadName *string `json:"WorkloadName,omitempty" xml:"WorkloadName,omitempty"`
}

func (s ListExperimentsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListExperimentsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListExperimentsResponseBodyData) SetCreateTime(v int64) *ListExperimentsResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *ListExperimentsResponseBodyData) SetEndTime(v string) *ListExperimentsResponseBodyData {
	s.EndTime = &v
	return s
}

func (s *ListExperimentsResponseBodyData) SetEnvParams(v *ListExperimentsResponseBodyDataEnvParams) *ListExperimentsResponseBodyData {
	s.EnvParams = v
	return s
}

func (s *ListExperimentsResponseBodyData) SetExperimentId(v int64) *ListExperimentsResponseBodyData {
	s.ExperimentId = &v
	return s
}

func (s *ListExperimentsResponseBodyData) SetExperimentName(v string) *ListExperimentsResponseBodyData {
	s.ExperimentName = &v
	return s
}

func (s *ListExperimentsResponseBodyData) SetExperimentType(v string) *ListExperimentsResponseBodyData {
	s.ExperimentType = &v
	return s
}

func (s *ListExperimentsResponseBodyData) SetGetParams(v map[string]*string) *ListExperimentsResponseBodyData {
	s.GetParams = v
	return s
}

func (s *ListExperimentsResponseBodyData) SetResourceName(v string) *ListExperimentsResponseBodyData {
	s.ResourceName = &v
	return s
}

func (s *ListExperimentsResponseBodyData) SetResults(v *ListExperimentsResponseBodyDataResults) *ListExperimentsResponseBodyData {
	s.Results = v
	return s
}

func (s *ListExperimentsResponseBodyData) SetSetParams(v map[string]*string) *ListExperimentsResponseBodyData {
	s.SetParams = v
	return s
}

func (s *ListExperimentsResponseBodyData) SetStartTime(v string) *ListExperimentsResponseBodyData {
	s.StartTime = &v
	return s
}

func (s *ListExperimentsResponseBodyData) SetStatus(v string) *ListExperimentsResponseBodyData {
	s.Status = &v
	return s
}

func (s *ListExperimentsResponseBodyData) SetUpdateTime(v int64) *ListExperimentsResponseBodyData {
	s.UpdateTime = &v
	return s
}

func (s *ListExperimentsResponseBodyData) SetWorkloadName(v string) *ListExperimentsResponseBodyData {
	s.WorkloadName = &v
	return s
}

type ListExperimentsResponseBodyDataEnvParams struct {
	// Number of CPUs allocated
	//
	// example:
	//
	// 90
	CpuPerWorker *int32 `json:"CpuPerWorker,omitempty" xml:"CpuPerWorker,omitempty"`
	// CUDA version
	//
	// example:
	//
	// 1.0.0
	CudaVersion *string `json:"CudaVersion,omitempty" xml:"CudaVersion,omitempty"`
	// Additional parameters
	ExtendParam map[string]*string `json:"ExtendParam,omitempty" xml:"ExtendParam,omitempty"`
	// GPU driver version
	//
	// example:
	//
	// 1.0.0
	GpuDriverVersion *string `json:"GpuDriverVersion,omitempty" xml:"GpuDriverVersion,omitempty"`
	// Number of GPUs allocated
	//
	// example:
	//
	// 8
	GpuPerWorker *int32 `json:"GpuPerWorker,omitempty" xml:"GpuPerWorker,omitempty"`
	// Amount of memory (GB) allocated
	//
	// example:
	//
	// 500
	MemoryPerWorker *int32 `json:"MemoryPerWorker,omitempty" xml:"MemoryPerWorker,omitempty"`
	// NCCL version
	//
	// example:
	//
	// 1.0.0
	NCCLVersion *string `json:"NCCLVersion,omitempty" xml:"NCCLVersion,omitempty"`
	// PyTorch version
	//
	// example:
	//
	// 1.0.0
	PyTorchVersion *string `json:"PyTorchVersion,omitempty" xml:"PyTorchVersion,omitempty"`
	// Specified nodes
	ResourceNodes []*ListExperimentsResponseBodyDataEnvParamsResourceNodes `json:"ResourceNodes,omitempty" xml:"ResourceNodes,omitempty" type:"Repeated"`
	// Amount of shared memory (GB) allocated
	//
	// example:
	//
	// 500
	ShareMemory *int32 `json:"ShareMemory,omitempty" xml:"ShareMemory,omitempty"`
	// Number of nodes
	//
	// example:
	//
	// 1
	WorkerNum *int32 `json:"WorkerNum,omitempty" xml:"WorkerNum,omitempty"`
}

func (s ListExperimentsResponseBodyDataEnvParams) String() string {
	return tea.Prettify(s)
}

func (s ListExperimentsResponseBodyDataEnvParams) GoString() string {
	return s.String()
}

func (s *ListExperimentsResponseBodyDataEnvParams) SetCpuPerWorker(v int32) *ListExperimentsResponseBodyDataEnvParams {
	s.CpuPerWorker = &v
	return s
}

func (s *ListExperimentsResponseBodyDataEnvParams) SetCudaVersion(v string) *ListExperimentsResponseBodyDataEnvParams {
	s.CudaVersion = &v
	return s
}

func (s *ListExperimentsResponseBodyDataEnvParams) SetExtendParam(v map[string]*string) *ListExperimentsResponseBodyDataEnvParams {
	s.ExtendParam = v
	return s
}

func (s *ListExperimentsResponseBodyDataEnvParams) SetGpuDriverVersion(v string) *ListExperimentsResponseBodyDataEnvParams {
	s.GpuDriverVersion = &v
	return s
}

func (s *ListExperimentsResponseBodyDataEnvParams) SetGpuPerWorker(v int32) *ListExperimentsResponseBodyDataEnvParams {
	s.GpuPerWorker = &v
	return s
}

func (s *ListExperimentsResponseBodyDataEnvParams) SetMemoryPerWorker(v int32) *ListExperimentsResponseBodyDataEnvParams {
	s.MemoryPerWorker = &v
	return s
}

func (s *ListExperimentsResponseBodyDataEnvParams) SetNCCLVersion(v string) *ListExperimentsResponseBodyDataEnvParams {
	s.NCCLVersion = &v
	return s
}

func (s *ListExperimentsResponseBodyDataEnvParams) SetPyTorchVersion(v string) *ListExperimentsResponseBodyDataEnvParams {
	s.PyTorchVersion = &v
	return s
}

func (s *ListExperimentsResponseBodyDataEnvParams) SetResourceNodes(v []*ListExperimentsResponseBodyDataEnvParamsResourceNodes) *ListExperimentsResponseBodyDataEnvParams {
	s.ResourceNodes = v
	return s
}

func (s *ListExperimentsResponseBodyDataEnvParams) SetShareMemory(v int32) *ListExperimentsResponseBodyDataEnvParams {
	s.ShareMemory = &v
	return s
}

func (s *ListExperimentsResponseBodyDataEnvParams) SetWorkerNum(v int32) *ListExperimentsResponseBodyDataEnvParams {
	s.WorkerNum = &v
	return s
}

type ListExperimentsResponseBodyDataEnvParamsResourceNodes struct {
	// Node name
	//
	// example:
	//
	// lingj1xxnjt1k4nv-mg18v52pydyuumae-0
	NodeName *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	// Requested CPU
	//
	// example:
	//
	// 90
	RequestCPU *int32 `json:"RequestCPU,omitempty" xml:"RequestCPU,omitempty"`
	// Requested GPU
	//
	// example:
	//
	// 8
	RequestGPU *int32 `json:"RequestGPU,omitempty" xml:"RequestGPU,omitempty"`
	// Requested memory
	//
	// example:
	//
	// 500
	RequestMemory *int32 `json:"RequestMemory,omitempty" xml:"RequestMemory,omitempty"`
	// Total CPU
	//
	// example:
	//
	// 90
	TotalCPU *int32 `json:"TotalCPU,omitempty" xml:"TotalCPU,omitempty"`
	// Total GPU
	//
	// example:
	//
	// 8
	TotalGPU *int32 `json:"TotalGPU,omitempty" xml:"TotalGPU,omitempty"`
	// Total memory
	//
	// example:
	//
	// 500
	TotalMemory *int64 `json:"TotalMemory,omitempty" xml:"TotalMemory,omitempty"`
}

func (s ListExperimentsResponseBodyDataEnvParamsResourceNodes) String() string {
	return tea.Prettify(s)
}

func (s ListExperimentsResponseBodyDataEnvParamsResourceNodes) GoString() string {
	return s.String()
}

func (s *ListExperimentsResponseBodyDataEnvParamsResourceNodes) SetNodeName(v string) *ListExperimentsResponseBodyDataEnvParamsResourceNodes {
	s.NodeName = &v
	return s
}

func (s *ListExperimentsResponseBodyDataEnvParamsResourceNodes) SetRequestCPU(v int32) *ListExperimentsResponseBodyDataEnvParamsResourceNodes {
	s.RequestCPU = &v
	return s
}

func (s *ListExperimentsResponseBodyDataEnvParamsResourceNodes) SetRequestGPU(v int32) *ListExperimentsResponseBodyDataEnvParamsResourceNodes {
	s.RequestGPU = &v
	return s
}

func (s *ListExperimentsResponseBodyDataEnvParamsResourceNodes) SetRequestMemory(v int32) *ListExperimentsResponseBodyDataEnvParamsResourceNodes {
	s.RequestMemory = &v
	return s
}

func (s *ListExperimentsResponseBodyDataEnvParamsResourceNodes) SetTotalCPU(v int32) *ListExperimentsResponseBodyDataEnvParamsResourceNodes {
	s.TotalCPU = &v
	return s
}

func (s *ListExperimentsResponseBodyDataEnvParamsResourceNodes) SetTotalGPU(v int32) *ListExperimentsResponseBodyDataEnvParamsResourceNodes {
	s.TotalGPU = &v
	return s
}

func (s *ListExperimentsResponseBodyDataEnvParamsResourceNodes) SetTotalMemory(v int64) *ListExperimentsResponseBodyDataEnvParamsResourceNodes {
	s.TotalMemory = &v
	return s
}

type ListExperimentsResponseBodyDataResults struct {
	// Duration
	//
	// example:
	//
	// 20
	Duration *float64 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// Error nodes
	ErrorWorker []*ListExperimentsResponseBodyDataResultsErrorWorker `json:"ErrorWorker,omitempty" xml:"ErrorWorker,omitempty" type:"Repeated"`
	// Parameter name
	//
	// example:
	//
	// 440
	ExperimentId *int64 `json:"ExperimentId,omitempty" xml:"ExperimentId,omitempty"`
	// MFU
	//
	// example:
	//
	// 34
	Mfu *float64 `json:"Mfu,omitempty" xml:"Mfu,omitempty"`
	// Samples per second
	//
	// example:
	//
	// 10
	SamplesPerSecond *float64 `json:"SamplesPerSecond,omitempty" xml:"SamplesPerSecond,omitempty"`
	// Seconds per iteration
	//
	// example:
	//
	// 89
	SecondsPerIteration *float64 `json:"SecondsPerIteration,omitempty" xml:"SecondsPerIteration,omitempty"`
	// Warning worker
	WarningWorker []*ListExperimentsResponseBodyDataResultsWarningWorker `json:"WarningWorker,omitempty" xml:"WarningWorker,omitempty" type:"Repeated"`
}

func (s ListExperimentsResponseBodyDataResults) String() string {
	return tea.Prettify(s)
}

func (s ListExperimentsResponseBodyDataResults) GoString() string {
	return s.String()
}

func (s *ListExperimentsResponseBodyDataResults) SetDuration(v float64) *ListExperimentsResponseBodyDataResults {
	s.Duration = &v
	return s
}

func (s *ListExperimentsResponseBodyDataResults) SetErrorWorker(v []*ListExperimentsResponseBodyDataResultsErrorWorker) *ListExperimentsResponseBodyDataResults {
	s.ErrorWorker = v
	return s
}

func (s *ListExperimentsResponseBodyDataResults) SetExperimentId(v int64) *ListExperimentsResponseBodyDataResults {
	s.ExperimentId = &v
	return s
}

func (s *ListExperimentsResponseBodyDataResults) SetMfu(v float64) *ListExperimentsResponseBodyDataResults {
	s.Mfu = &v
	return s
}

func (s *ListExperimentsResponseBodyDataResults) SetSamplesPerSecond(v float64) *ListExperimentsResponseBodyDataResults {
	s.SamplesPerSecond = &v
	return s
}

func (s *ListExperimentsResponseBodyDataResults) SetSecondsPerIteration(v float64) *ListExperimentsResponseBodyDataResults {
	s.SecondsPerIteration = &v
	return s
}

func (s *ListExperimentsResponseBodyDataResults) SetWarningWorker(v []*ListExperimentsResponseBodyDataResultsWarningWorker) *ListExperimentsResponseBodyDataResults {
	s.WarningWorker = v
	return s
}

type ListExperimentsResponseBodyDataResultsErrorWorker struct {
	// Whether there is an error
	//
	// example:
	//
	// false
	ErrorFlag *bool `json:"ErrorFlag,omitempty" xml:"ErrorFlag,omitempty"`
	// Error information
	//
	// example:
	//
	// error msg
	ErrorMsg *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	// Experiment ID
	//
	// example:
	//
	// 176
	ExperimentId *int64 `json:"ExperimentId,omitempty" xml:"ExperimentId,omitempty"`
	// GPU name
	//
	// example:
	//
	// 8x OAM 810 GPU
	GpuName *string `json:"GpuName,omitempty" xml:"GpuName,omitempty"`
	// Number of GPUs
	//
	// example:
	//
	// 8
	GpuNum *int32 `json:"GpuNum,omitempty" xml:"GpuNum,omitempty"`
	// Host IP
	//
	// example:
	//
	// etcd_cluster_c0n2
	Hostname *string `json:"Hostname,omitempty" xml:"Hostname,omitempty"`
	// Pod name
	//
	// example:
	//
	// fluxserv-6fc89b45cf-w8wq6
	PodName *string `json:"PodName,omitempty" xml:"PodName,omitempty"`
	// Throughput
	//
	// example:
	//
	// 65
	SamplesPerSecond *float64 `json:"SamplesPerSecond,omitempty" xml:"SamplesPerSecond,omitempty"`
	// TFLOPS value
	//
	// example:
	//
	// 42
	Tflops *float64 `json:"Tflops,omitempty" xml:"Tflops,omitempty"`
	// Whether there is an alarm
	//
	// example:
	//
	// false
	WarningFlag *bool `json:"WarningFlag,omitempty" xml:"WarningFlag,omitempty"`
	// Alarm information
	//
	// example:
	//
	// warning msg
	WarningMsg *string `json:"WarningMsg,omitempty" xml:"WarningMsg,omitempty"`
}

func (s ListExperimentsResponseBodyDataResultsErrorWorker) String() string {
	return tea.Prettify(s)
}

func (s ListExperimentsResponseBodyDataResultsErrorWorker) GoString() string {
	return s.String()
}

func (s *ListExperimentsResponseBodyDataResultsErrorWorker) SetErrorFlag(v bool) *ListExperimentsResponseBodyDataResultsErrorWorker {
	s.ErrorFlag = &v
	return s
}

func (s *ListExperimentsResponseBodyDataResultsErrorWorker) SetErrorMsg(v string) *ListExperimentsResponseBodyDataResultsErrorWorker {
	s.ErrorMsg = &v
	return s
}

func (s *ListExperimentsResponseBodyDataResultsErrorWorker) SetExperimentId(v int64) *ListExperimentsResponseBodyDataResultsErrorWorker {
	s.ExperimentId = &v
	return s
}

func (s *ListExperimentsResponseBodyDataResultsErrorWorker) SetGpuName(v string) *ListExperimentsResponseBodyDataResultsErrorWorker {
	s.GpuName = &v
	return s
}

func (s *ListExperimentsResponseBodyDataResultsErrorWorker) SetGpuNum(v int32) *ListExperimentsResponseBodyDataResultsErrorWorker {
	s.GpuNum = &v
	return s
}

func (s *ListExperimentsResponseBodyDataResultsErrorWorker) SetHostname(v string) *ListExperimentsResponseBodyDataResultsErrorWorker {
	s.Hostname = &v
	return s
}

func (s *ListExperimentsResponseBodyDataResultsErrorWorker) SetPodName(v string) *ListExperimentsResponseBodyDataResultsErrorWorker {
	s.PodName = &v
	return s
}

func (s *ListExperimentsResponseBodyDataResultsErrorWorker) SetSamplesPerSecond(v float64) *ListExperimentsResponseBodyDataResultsErrorWorker {
	s.SamplesPerSecond = &v
	return s
}

func (s *ListExperimentsResponseBodyDataResultsErrorWorker) SetTflops(v float64) *ListExperimentsResponseBodyDataResultsErrorWorker {
	s.Tflops = &v
	return s
}

func (s *ListExperimentsResponseBodyDataResultsErrorWorker) SetWarningFlag(v bool) *ListExperimentsResponseBodyDataResultsErrorWorker {
	s.WarningFlag = &v
	return s
}

func (s *ListExperimentsResponseBodyDataResultsErrorWorker) SetWarningMsg(v string) *ListExperimentsResponseBodyDataResultsErrorWorker {
	s.WarningMsg = &v
	return s
}

type ListExperimentsResponseBodyDataResultsWarningWorker struct {
	// Whether there is an error
	//
	// example:
	//
	// false
	ErrorFlag *bool `json:"ErrorFlag,omitempty" xml:"ErrorFlag,omitempty"`
	// Error message
	//
	// example:
	//
	// error msg
	ErrorMsg *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	// Experiment ID
	//
	// example:
	//
	// 113
	ExperimentId *int64 `json:"ExperimentId,omitempty" xml:"ExperimentId,omitempty"`
	// GPU name
	//
	// example:
	//
	// 8x OAM 810 GPU
	GpuName *string `json:"GpuName,omitempty" xml:"GpuName,omitempty"`
	// Number of GPUs
	//
	// example:
	//
	// 90
	GpuNum *int32 `json:"GpuNum,omitempty" xml:"GpuNum,omitempty"`
	// Host IP
	//
	// example:
	//
	// 101.66.165.102
	Hostname *string `json:"Hostname,omitempty" xml:"Hostname,omitempty"`
	// Pod name
	//
	// example:
	//
	// hzs-forge-sdxl-online-7ff4d86444-pc95h
	PodName *string `json:"PodName,omitempty" xml:"PodName,omitempty"`
	// Throughput
	//
	// example:
	//
	// 53
	SamplesPerSecond *float64 `json:"SamplesPerSecond,omitempty" xml:"SamplesPerSecond,omitempty"`
	// TFLOPS value
	//
	// example:
	//
	// 43
	Tflops *float64 `json:"Tflops,omitempty" xml:"Tflops,omitempty"`
	// Whether there is an alarm
	//
	// example:
	//
	// false
	WarningFlag *bool `json:"WarningFlag,omitempty" xml:"WarningFlag,omitempty"`
	// Alarm message
	//
	// example:
	//
	// warning msg
	WarningMsg *string `json:"WarningMsg,omitempty" xml:"WarningMsg,omitempty"`
}

func (s ListExperimentsResponseBodyDataResultsWarningWorker) String() string {
	return tea.Prettify(s)
}

func (s ListExperimentsResponseBodyDataResultsWarningWorker) GoString() string {
	return s.String()
}

func (s *ListExperimentsResponseBodyDataResultsWarningWorker) SetErrorFlag(v bool) *ListExperimentsResponseBodyDataResultsWarningWorker {
	s.ErrorFlag = &v
	return s
}

func (s *ListExperimentsResponseBodyDataResultsWarningWorker) SetErrorMsg(v string) *ListExperimentsResponseBodyDataResultsWarningWorker {
	s.ErrorMsg = &v
	return s
}

func (s *ListExperimentsResponseBodyDataResultsWarningWorker) SetExperimentId(v int64) *ListExperimentsResponseBodyDataResultsWarningWorker {
	s.ExperimentId = &v
	return s
}

func (s *ListExperimentsResponseBodyDataResultsWarningWorker) SetGpuName(v string) *ListExperimentsResponseBodyDataResultsWarningWorker {
	s.GpuName = &v
	return s
}

func (s *ListExperimentsResponseBodyDataResultsWarningWorker) SetGpuNum(v int32) *ListExperimentsResponseBodyDataResultsWarningWorker {
	s.GpuNum = &v
	return s
}

func (s *ListExperimentsResponseBodyDataResultsWarningWorker) SetHostname(v string) *ListExperimentsResponseBodyDataResultsWarningWorker {
	s.Hostname = &v
	return s
}

func (s *ListExperimentsResponseBodyDataResultsWarningWorker) SetPodName(v string) *ListExperimentsResponseBodyDataResultsWarningWorker {
	s.PodName = &v
	return s
}

func (s *ListExperimentsResponseBodyDataResultsWarningWorker) SetSamplesPerSecond(v float64) *ListExperimentsResponseBodyDataResultsWarningWorker {
	s.SamplesPerSecond = &v
	return s
}

func (s *ListExperimentsResponseBodyDataResultsWarningWorker) SetTflops(v float64) *ListExperimentsResponseBodyDataResultsWarningWorker {
	s.Tflops = &v
	return s
}

func (s *ListExperimentsResponseBodyDataResultsWarningWorker) SetWarningFlag(v bool) *ListExperimentsResponseBodyDataResultsWarningWorker {
	s.WarningFlag = &v
	return s
}

func (s *ListExperimentsResponseBodyDataResultsWarningWorker) SetWarningMsg(v string) *ListExperimentsResponseBodyDataResultsWarningWorker {
	s.WarningMsg = &v
	return s
}

type ListExperimentsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListExperimentsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListExperimentsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListExperimentsResponse) GoString() string {
	return s.String()
}

func (s *ListExperimentsResponse) SetHeaders(v map[string]*string) *ListExperimentsResponse {
	s.Headers = v
	return s
}

func (s *ListExperimentsResponse) SetStatusCode(v int32) *ListExperimentsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListExperimentsResponse) SetBody(v *ListExperimentsResponseBody) *ListExperimentsResponse {
	s.Body = v
	return s
}

type ListTagResourcesRequest struct {
	// Next token for the next query
	//
	// example:
	//
	// F0lqbr2JpLDppro1RahGKViWtqXr3L28cePimcRn
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// ResourceId
	ResourceId []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	// Resource type
	//
	// This parameter is required.
	//
	// example:
	//
	// ExperimentPlan
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The list of tags to be queried. The value range for N is 1~20.
	Tag []*ListTagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
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

func (s *ListTagResourcesRequest) SetResourceId(v []*string) *ListTagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *ListTagResourcesRequest) SetResourceType(v string) *ListTagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *ListTagResourcesRequest) SetTag(v []*ListTagResourcesRequestTag) *ListTagResourcesRequest {
	s.Tag = v
	return s
}

type ListTagResourcesRequestTag struct {
	// Tag key, with n in the range [1, 20].
	//
	// example:
	//
	// owner
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// Tag value
	//
	// example:
	//
	// syg
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListTagResourcesRequestTag) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesRequestTag) GoString() string {
	return s.String()
}

func (s *ListTagResourcesRequestTag) SetKey(v string) *ListTagResourcesRequestTag {
	s.Key = &v
	return s
}

func (s *ListTagResourcesRequestTag) SetValue(v string) *ListTagResourcesRequestTag {
	s.Value = &v
	return s
}

type ListTagResourcesResponseBody struct {
	// Access denied information
	//
	// example:
	//
	// {}
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// Next token for the next query. An empty NextToken indicates there are no more results.
	//
	// example:
	//
	// uPZbmbpgxp2/6vNWNPoase3Eqy+gL9pdDBH7KGZXMuZ9GxmBbMJcTP/dlrNqRaWF
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// Request ID
	//
	// example:
	//
	// 5514CB39-B7C0-5B89-8534-2DE1E0F2B7AB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// List of resources
	TagResources []*ListTagResourcesResponseBodyTagResources `json:"TagResources,omitempty" xml:"TagResources,omitempty" type:"Repeated"`
	// Total
	//
	// example:
	//
	// 0
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListTagResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBody) SetAccessDeniedDetail(v string) *ListTagResourcesResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ListTagResourcesResponseBody) SetNextToken(v string) *ListTagResourcesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListTagResourcesResponseBody) SetRequestId(v string) *ListTagResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTagResourcesResponseBody) SetTagResources(v []*ListTagResourcesResponseBodyTagResources) *ListTagResourcesResponseBody {
	s.TagResources = v
	return s
}

func (s *ListTagResourcesResponseBody) SetTotalCount(v int64) *ListTagResourcesResponseBody {
	s.TotalCount = &v
	return s
}

type ListTagResourcesResponseBodyTagResources struct {
	// ResourceId
	//
	// example:
	//
	// 189
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// Resource type
	//
	// example:
	//
	// ExperimentPlan
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// Tag key
	//
	// example:
	//
	// owner
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// Tag value
	//
	// example:
	//
	// syg
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s ListTagResourcesResponseBodyTagResources) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponseBodyTagResources) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBodyTagResources) SetResourceId(v string) *ListTagResourcesResponseBodyTagResources {
	s.ResourceId = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResources) SetResourceType(v string) *ListTagResourcesResponseBodyTagResources {
	s.ResourceType = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResources) SetTagKey(v string) *ListTagResourcesResponseBodyTagResources {
	s.TagKey = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResources) SetTagValue(v string) *ListTagResourcesResponseBodyTagResources {
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

type ListWorkloadsRequest struct {
	// Scope
	//
	// example:
	//
	// common
	Scope *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
}

func (s ListWorkloadsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListWorkloadsRequest) GoString() string {
	return s.String()
}

func (s *ListWorkloadsRequest) SetScope(v string) *ListWorkloadsRequest {
	s.Scope = &v
	return s
}

type ListWorkloadsResponseBody struct {
	// Access Denied Information
	//
	// example:
	//
	// {}
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// Data
	Data []*ListWorkloadsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// Request ID
	//
	// example:
	//
	// 4AC08332-436C-57A3-9FBA-26772B1A9901
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// total
	//
	// example:
	//
	// 1
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListWorkloadsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListWorkloadsResponseBody) GoString() string {
	return s.String()
}

func (s *ListWorkloadsResponseBody) SetAccessDeniedDetail(v string) *ListWorkloadsResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ListWorkloadsResponseBody) SetData(v []*ListWorkloadsResponseBodyData) *ListWorkloadsResponseBody {
	s.Data = v
	return s
}

func (s *ListWorkloadsResponseBody) SetRequestId(v string) *ListWorkloadsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListWorkloadsResponseBody) SetTotalCount(v int64) *ListWorkloadsResponseBody {
	s.TotalCount = &v
	return s
}

type ListWorkloadsResponseBodyData struct {
	// Default CPU Allocation
	//
	// example:
	//
	// 90
	DefaultCpuPerWorker *int32 `json:"DefaultCpuPerWorker,omitempty" xml:"DefaultCpuPerWorker,omitempty"`
	// Default GPU Allocation
	//
	// example:
	//
	// 8
	DefaultGpuPerWorker *int32 `json:"DefaultGpuPerWorker,omitempty" xml:"DefaultGpuPerWorker,omitempty"`
	// Default Memory (GB) Allocation
	//
	// example:
	//
	// 500
	DefaultMemoryPerWorker *int32 `json:"DefaultMemoryPerWorker,omitempty" xml:"DefaultMemoryPerWorker,omitempty"`
	// Default Shared Memory (GB) Allocation
	//
	// example:
	//
	// 500
	DefaultShareMemory *int32 `json:"DefaultShareMemory,omitempty" xml:"DefaultShareMemory,omitempty"`
	// Workload Cluster, AI, GPU
	//
	// example:
	//
	// AI
	Family *string `json:"Family,omitempty" xml:"Family,omitempty"`
	// Training Job Type
	//
	// example:
	//
	// PyTorchJob
	JobKind *string `json:"JobKind,omitempty" xml:"JobKind,omitempty"`
	// Parameter Settings
	ParamSettings []*ListWorkloadsResponseBodyDataParamSettings `json:"ParamSettings,omitempty" xml:"ParamSettings,omitempty" type:"Repeated"`
	// Workload Usage Scenario
	//
	// example:
	//
	// NLP-LLM
	Scene *string `json:"Scene,omitempty" xml:"Scene,omitempty"`
	// Scope Identifier for Workload Usage
	//
	// example:
	//
	// common
	Scope *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
	// Static Configuration
	StaticConfig *ListWorkloadsResponseBodyDataStaticConfig `json:"StaticConfig,omitempty" xml:"StaticConfig,omitempty" type:"Struct"`
	// Version ID
	//
	// example:
	//
	// 1
	VersionId *int64 `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
	// Workload Description
	//
	// example:
	//
	// test
	WorkloadDescription *string `json:"WorkloadDescription,omitempty" xml:"WorkloadDescription,omitempty"`
	// Workload ID
	//
	// example:
	//
	// 13
	WorkloadId *int64 `json:"WorkloadId,omitempty" xml:"WorkloadId,omitempty"`
	// Workload Name
	//
	// example:
	//
	// test
	WorkloadName *string `json:"WorkloadName,omitempty" xml:"WorkloadName,omitempty"`
	// Workload Type
	//
	// example:
	//
	// AI
	WorkloadType *string `json:"WorkloadType,omitempty" xml:"WorkloadType,omitempty"`
}

func (s ListWorkloadsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListWorkloadsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListWorkloadsResponseBodyData) SetDefaultCpuPerWorker(v int32) *ListWorkloadsResponseBodyData {
	s.DefaultCpuPerWorker = &v
	return s
}

func (s *ListWorkloadsResponseBodyData) SetDefaultGpuPerWorker(v int32) *ListWorkloadsResponseBodyData {
	s.DefaultGpuPerWorker = &v
	return s
}

func (s *ListWorkloadsResponseBodyData) SetDefaultMemoryPerWorker(v int32) *ListWorkloadsResponseBodyData {
	s.DefaultMemoryPerWorker = &v
	return s
}

func (s *ListWorkloadsResponseBodyData) SetDefaultShareMemory(v int32) *ListWorkloadsResponseBodyData {
	s.DefaultShareMemory = &v
	return s
}

func (s *ListWorkloadsResponseBodyData) SetFamily(v string) *ListWorkloadsResponseBodyData {
	s.Family = &v
	return s
}

func (s *ListWorkloadsResponseBodyData) SetJobKind(v string) *ListWorkloadsResponseBodyData {
	s.JobKind = &v
	return s
}

func (s *ListWorkloadsResponseBodyData) SetParamSettings(v []*ListWorkloadsResponseBodyDataParamSettings) *ListWorkloadsResponseBodyData {
	s.ParamSettings = v
	return s
}

func (s *ListWorkloadsResponseBodyData) SetScene(v string) *ListWorkloadsResponseBodyData {
	s.Scene = &v
	return s
}

func (s *ListWorkloadsResponseBodyData) SetScope(v string) *ListWorkloadsResponseBodyData {
	s.Scope = &v
	return s
}

func (s *ListWorkloadsResponseBodyData) SetStaticConfig(v *ListWorkloadsResponseBodyDataStaticConfig) *ListWorkloadsResponseBodyData {
	s.StaticConfig = v
	return s
}

func (s *ListWorkloadsResponseBodyData) SetVersionId(v int64) *ListWorkloadsResponseBodyData {
	s.VersionId = &v
	return s
}

func (s *ListWorkloadsResponseBodyData) SetWorkloadDescription(v string) *ListWorkloadsResponseBodyData {
	s.WorkloadDescription = &v
	return s
}

func (s *ListWorkloadsResponseBodyData) SetWorkloadId(v int64) *ListWorkloadsResponseBodyData {
	s.WorkloadId = &v
	return s
}

func (s *ListWorkloadsResponseBodyData) SetWorkloadName(v string) *ListWorkloadsResponseBodyData {
	s.WorkloadName = &v
	return s
}

func (s *ListWorkloadsResponseBodyData) SetWorkloadType(v string) *ListWorkloadsResponseBodyData {
	s.WorkloadType = &v
	return s
}

type ListWorkloadsResponseBodyDataParamSettings struct {
	// Default Parameter Value
	//
	// example:
	//
	// 100
	DefaultValue *string `json:"DefaultValue,omitempty" xml:"DefaultValue,omitempty"`
	// Parameter Description
	//
	// example:
	//
	// number
	ParamDesc *string `json:"ParamDesc,omitempty" xml:"ParamDesc,omitempty"`
	// Parameter Name
	//
	// example:
	//
	// ITERATION
	ParamName *string `json:"ParamName,omitempty" xml:"ParamName,omitempty"`
	// Parameter Regular Expression
	//
	// example:
	//
	// [0-9]+
	ParamRegex *string `json:"ParamRegex,omitempty" xml:"ParamRegex,omitempty"`
	// Parameter type
	//
	// example:
	//
	// number
	ParamType *string `json:"ParamType,omitempty" xml:"ParamType,omitempty"`
	// Parameter Value
	//
	// example:
	//
	// 100
	ParamValue *string `json:"ParamValue,omitempty" xml:"ParamValue,omitempty"`
}

func (s ListWorkloadsResponseBodyDataParamSettings) String() string {
	return tea.Prettify(s)
}

func (s ListWorkloadsResponseBodyDataParamSettings) GoString() string {
	return s.String()
}

func (s *ListWorkloadsResponseBodyDataParamSettings) SetDefaultValue(v string) *ListWorkloadsResponseBodyDataParamSettings {
	s.DefaultValue = &v
	return s
}

func (s *ListWorkloadsResponseBodyDataParamSettings) SetParamDesc(v string) *ListWorkloadsResponseBodyDataParamSettings {
	s.ParamDesc = &v
	return s
}

func (s *ListWorkloadsResponseBodyDataParamSettings) SetParamName(v string) *ListWorkloadsResponseBodyDataParamSettings {
	s.ParamName = &v
	return s
}

func (s *ListWorkloadsResponseBodyDataParamSettings) SetParamRegex(v string) *ListWorkloadsResponseBodyDataParamSettings {
	s.ParamRegex = &v
	return s
}

func (s *ListWorkloadsResponseBodyDataParamSettings) SetParamType(v string) *ListWorkloadsResponseBodyDataParamSettings {
	s.ParamType = &v
	return s
}

func (s *ListWorkloadsResponseBodyDataParamSettings) SetParamValue(v string) *ListWorkloadsResponseBodyDataParamSettings {
	s.ParamValue = &v
	return s
}

type ListWorkloadsResponseBodyDataStaticConfig struct {
	// Framework
	//
	// example:
	//
	// PyTorch
	FrameWork *string `json:"FrameWork,omitempty" xml:"FrameWork,omitempty"`
	// Operating System
	//
	// example:
	//
	// linux
	Os *string `json:"Os,omitempty" xml:"Os,omitempty"`
	// Number of Parameters
	//
	// example:
	//
	// 7B
	Parameters *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// Software Stack
	//
	// example:
	//
	// python
	SoftwareStack *string `json:"SoftwareStack,omitempty" xml:"SoftwareStack,omitempty"`
}

func (s ListWorkloadsResponseBodyDataStaticConfig) String() string {
	return tea.Prettify(s)
}

func (s ListWorkloadsResponseBodyDataStaticConfig) GoString() string {
	return s.String()
}

func (s *ListWorkloadsResponseBodyDataStaticConfig) SetFrameWork(v string) *ListWorkloadsResponseBodyDataStaticConfig {
	s.FrameWork = &v
	return s
}

func (s *ListWorkloadsResponseBodyDataStaticConfig) SetOs(v string) *ListWorkloadsResponseBodyDataStaticConfig {
	s.Os = &v
	return s
}

func (s *ListWorkloadsResponseBodyDataStaticConfig) SetParameters(v string) *ListWorkloadsResponseBodyDataStaticConfig {
	s.Parameters = &v
	return s
}

func (s *ListWorkloadsResponseBodyDataStaticConfig) SetSoftwareStack(v string) *ListWorkloadsResponseBodyDataStaticConfig {
	s.SoftwareStack = &v
	return s
}

type ListWorkloadsResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListWorkloadsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListWorkloadsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListWorkloadsResponse) GoString() string {
	return s.String()
}

func (s *ListWorkloadsResponse) SetHeaders(v map[string]*string) *ListWorkloadsResponse {
	s.Headers = v
	return s
}

func (s *ListWorkloadsResponse) SetStatusCode(v int32) *ListWorkloadsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListWorkloadsResponse) SetBody(v *ListWorkloadsResponseBody) *ListWorkloadsResponse {
	s.Body = v
	return s
}

type StopExperimentRequest struct {
	// Plan ID
	//
	// This parameter is required.
	//
	// example:
	//
	// 234
	ExperimentId *int64 `json:"ExperimentId,omitempty" xml:"ExperimentId,omitempty"`
	// Resource Group Id
	//
	// example:
	//
	// rg-kdsflsdfj23m
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s StopExperimentRequest) String() string {
	return tea.Prettify(s)
}

func (s StopExperimentRequest) GoString() string {
	return s.String()
}

func (s *StopExperimentRequest) SetExperimentId(v int64) *StopExperimentRequest {
	s.ExperimentId = &v
	return s
}

func (s *StopExperimentRequest) SetResourceGroupId(v string) *StopExperimentRequest {
	s.ResourceGroupId = &v
	return s
}

type StopExperimentResponseBody struct {
	// Access denied information
	//
	// example:
	//
	// {}
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// Data
	//
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// Request ID
	//
	// example:
	//
	// 5514CB39-B7C0-5B89-8534-2DE1E0F2B7AB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Total number of queries
	//
	// example:
	//
	// 0
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s StopExperimentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StopExperimentResponseBody) GoString() string {
	return s.String()
}

func (s *StopExperimentResponseBody) SetAccessDeniedDetail(v string) *StopExperimentResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *StopExperimentResponseBody) SetData(v bool) *StopExperimentResponseBody {
	s.Data = &v
	return s
}

func (s *StopExperimentResponseBody) SetRequestId(v string) *StopExperimentResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopExperimentResponseBody) SetTotalCount(v int64) *StopExperimentResponseBody {
	s.TotalCount = &v
	return s
}

type StopExperimentResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopExperimentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopExperimentResponse) String() string {
	return tea.Prettify(s)
}

func (s StopExperimentResponse) GoString() string {
	return s.String()
}

func (s *StopExperimentResponse) SetHeaders(v map[string]*string) *StopExperimentResponse {
	s.Headers = v
	return s
}

func (s *StopExperimentResponse) SetStatusCode(v int32) *StopExperimentResponse {
	s.StatusCode = &v
	return s
}

func (s *StopExperimentResponse) SetBody(v *StopExperimentResponseBody) *StopExperimentResponse {
	s.Body = v
	return s
}

type TagResourcesRequest struct {
	// ResourceId
	//
	// This parameter is required.
	ResourceId []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	// Resource type
	//
	// This parameter is required.
	//
	// example:
	//
	// ExperimentPlan
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// List of tags, up to 20.
	//
	// This parameter is required.
	Tag []*TagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s TagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesRequest) GoString() string {
	return s.String()
}

func (s *TagResourcesRequest) SetResourceId(v []*string) *TagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *TagResourcesRequest) SetResourceType(v string) *TagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *TagResourcesRequest) SetTag(v []*TagResourcesRequestTag) *TagResourcesRequest {
	s.Tag = v
	return s
}

type TagResourcesRequestTag struct {
	// Tag key.
	//
	// example:
	//
	// owner
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// Tag value
	//
	// example:
	//
	// syg
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s TagResourcesRequestTag) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesRequestTag) GoString() string {
	return s.String()
}

func (s *TagResourcesRequestTag) SetKey(v string) *TagResourcesRequestTag {
	s.Key = &v
	return s
}

func (s *TagResourcesRequestTag) SetValue(v string) *TagResourcesRequestTag {
	s.Value = &v
	return s
}

type TagResourcesResponseBody struct {
	// Access denied details
	//
	// example:
	//
	// {}
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// Data
	//
	// example:
	//
	// []
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// Request ID
	//
	// example:
	//
	// E67E2E4C-2B47-5C55-AA17-1D771E070AEF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Total
	//
	// example:
	//
	// 0
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s TagResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *TagResourcesResponseBody) SetAccessDeniedDetail(v string) *TagResourcesResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *TagResourcesResponseBody) SetData(v string) *TagResourcesResponseBody {
	s.Data = &v
	return s
}

func (s *TagResourcesResponseBody) SetRequestId(v string) *TagResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *TagResourcesResponseBody) SetTotalCount(v int64) *TagResourcesResponseBody {
	s.TotalCount = &v
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

type UntagResourcesRequest struct {
	// Whether to delete all, only effective when TagKey.N is empty. Allowed values: true, false, True, False. Default is false.
	//
	// example:
	//
	// true
	All *bool `json:"All,omitempty" xml:"All,omitempty"`
	// Resource ID
	//
	// This parameter is required.
	ResourceId []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	// Resource type
	//
	// This parameter is required.
	//
	// example:
	//
	// ExperimentPlan
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// Tag key group, up to 20 items
	//
	// This parameter is required.
	TagKey []*string `json:"TagKey,omitempty" xml:"TagKey,omitempty" type:"Repeated"`
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

func (s *UntagResourcesRequest) SetResourceId(v []*string) *UntagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *UntagResourcesRequest) SetResourceType(v string) *UntagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *UntagResourcesRequest) SetTagKey(v []*string) *UntagResourcesRequest {
	s.TagKey = v
	return s
}

type UntagResourcesResponseBody struct {
	// Access denied information
	//
	// example:
	//
	// {}
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// Data
	//
	// example:
	//
	// []
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// Request ID
	//
	// example:
	//
	// 25859897-35C8-5015-8365-7A3CE52F4854
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Total
	//
	// example:
	//
	// 0
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s UntagResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *UntagResourcesResponseBody) SetAccessDeniedDetail(v string) *UntagResourcesResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *UntagResourcesResponseBody) SetData(v string) *UntagResourcesResponseBody {
	s.Data = &v
	return s
}

func (s *UntagResourcesResponseBody) SetRequestId(v string) *UntagResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *UntagResourcesResponseBody) SetTotalCount(v int64) *UntagResourcesResponseBody {
	s.TotalCount = &v
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

type UpdateExperimentPlanRequest struct {
	// Experiment plan ID
	//
	// This parameter is required.
	//
	// example:
	//
	// 189
	PlanId *int64 `json:"PlanId,omitempty" xml:"PlanId,omitempty"`
	// Experiment plan name
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	PlanTemplateName *string `json:"PlanTemplateName,omitempty" xml:"PlanTemplateName,omitempty"`
}

func (s UpdateExperimentPlanRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateExperimentPlanRequest) GoString() string {
	return s.String()
}

func (s *UpdateExperimentPlanRequest) SetPlanId(v int64) *UpdateExperimentPlanRequest {
	s.PlanId = &v
	return s
}

func (s *UpdateExperimentPlanRequest) SetPlanTemplateName(v string) *UpdateExperimentPlanRequest {
	s.PlanTemplateName = &v
	return s
}

type UpdateExperimentPlanResponseBody struct {
	// Access denied information
	//
	// example:
	//
	// {}
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// Data
	//
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// Request ID
	//
	// example:
	//
	// 5514CB39-B7C0-5B89-8534-2DE1E0F2B7AB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Total
	//
	// example:
	//
	// 0
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s UpdateExperimentPlanResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateExperimentPlanResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateExperimentPlanResponseBody) SetAccessDeniedDetail(v string) *UpdateExperimentPlanResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *UpdateExperimentPlanResponseBody) SetData(v bool) *UpdateExperimentPlanResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateExperimentPlanResponseBody) SetRequestId(v string) *UpdateExperimentPlanResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateExperimentPlanResponseBody) SetTotalCount(v int64) *UpdateExperimentPlanResponseBody {
	s.TotalCount = &v
	return s
}

type UpdateExperimentPlanResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateExperimentPlanResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateExperimentPlanResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateExperimentPlanResponse) GoString() string {
	return s.String()
}

func (s *UpdateExperimentPlanResponse) SetHeaders(v map[string]*string) *UpdateExperimentPlanResponse {
	s.Headers = v
	return s
}

func (s *UpdateExperimentPlanResponse) SetStatusCode(v int32) *UpdateExperimentPlanResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateExperimentPlanResponse) SetBody(v *UpdateExperimentPlanResponseBody) *UpdateExperimentPlanResponse {
	s.Body = v
	return s
}

type UpdateExperimentPlanTemplateRequest struct {
	// Template code
	//
	// This parameter is required.
	//
	// example:
	//
	// 349623
	TemplateId *int64 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// Template pipeline
	//
	// This parameter is required.
	TemplatePipeline []*UpdateExperimentPlanTemplateRequestTemplatePipeline `json:"TemplatePipeline,omitempty" xml:"TemplatePipeline,omitempty" type:"Repeated"`
}

func (s UpdateExperimentPlanTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateExperimentPlanTemplateRequest) GoString() string {
	return s.String()
}

func (s *UpdateExperimentPlanTemplateRequest) SetTemplateId(v int64) *UpdateExperimentPlanTemplateRequest {
	s.TemplateId = &v
	return s
}

func (s *UpdateExperimentPlanTemplateRequest) SetTemplatePipeline(v []*UpdateExperimentPlanTemplateRequestTemplatePipeline) *UpdateExperimentPlanTemplateRequest {
	s.TemplatePipeline = v
	return s
}

type UpdateExperimentPlanTemplateRequestTemplatePipeline struct {
	// Configured environment parameters
	//
	// This parameter is required.
	EnvParams *UpdateExperimentPlanTemplateRequestTemplatePipelineEnvParams `json:"EnvParams,omitempty" xml:"EnvParams,omitempty" type:"Struct"`
	// Node order number
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PipelineOrder *int32 `json:"PipelineOrder,omitempty" xml:"PipelineOrder,omitempty"`
	// Usage scenario, e.g., "baseline"
	//
	// This parameter is required.
	//
	// example:
	//
	// baseline
	Scene *string `json:"Scene,omitempty" xml:"Scene,omitempty"`
	// Configured workload parameters
	SettingParams map[string]*string `json:"SettingParams,omitempty" xml:"SettingParams,omitempty"`
	// Workload ID
	//
	// This parameter is required.
	//
	// example:
	//
	// 14
	WorkloadId *int64 `json:"WorkloadId,omitempty" xml:"WorkloadId,omitempty"`
	// Workload name
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	WorkloadName *string `json:"WorkloadName,omitempty" xml:"WorkloadName,omitempty"`
}

func (s UpdateExperimentPlanTemplateRequestTemplatePipeline) String() string {
	return tea.Prettify(s)
}

func (s UpdateExperimentPlanTemplateRequestTemplatePipeline) GoString() string {
	return s.String()
}

func (s *UpdateExperimentPlanTemplateRequestTemplatePipeline) SetEnvParams(v *UpdateExperimentPlanTemplateRequestTemplatePipelineEnvParams) *UpdateExperimentPlanTemplateRequestTemplatePipeline {
	s.EnvParams = v
	return s
}

func (s *UpdateExperimentPlanTemplateRequestTemplatePipeline) SetPipelineOrder(v int32) *UpdateExperimentPlanTemplateRequestTemplatePipeline {
	s.PipelineOrder = &v
	return s
}

func (s *UpdateExperimentPlanTemplateRequestTemplatePipeline) SetScene(v string) *UpdateExperimentPlanTemplateRequestTemplatePipeline {
	s.Scene = &v
	return s
}

func (s *UpdateExperimentPlanTemplateRequestTemplatePipeline) SetSettingParams(v map[string]*string) *UpdateExperimentPlanTemplateRequestTemplatePipeline {
	s.SettingParams = v
	return s
}

func (s *UpdateExperimentPlanTemplateRequestTemplatePipeline) SetWorkloadId(v int64) *UpdateExperimentPlanTemplateRequestTemplatePipeline {
	s.WorkloadId = &v
	return s
}

func (s *UpdateExperimentPlanTemplateRequestTemplatePipeline) SetWorkloadName(v string) *UpdateExperimentPlanTemplateRequestTemplatePipeline {
	s.WorkloadName = &v
	return s
}

type UpdateExperimentPlanTemplateRequestTemplatePipelineEnvParams struct {
	// CPU allocation count
	//
	// This parameter is required.
	//
	// example:
	//
	// 90
	CpuPerWorker *int32 `json:"CpuPerWorker,omitempty" xml:"CpuPerWorker,omitempty"`
	// CUDA version
	//
	// example:
	//
	// 1.0.0
	CudaVersion *string `json:"CudaVersion,omitempty" xml:"CudaVersion,omitempty"`
	// GPU driver version
	//
	// example:
	//
	// 1.0.0
	GpuDriverVersion *string `json:"GpuDriverVersion,omitempty" xml:"GpuDriverVersion,omitempty"`
	// GPU allocation count
	//
	// This parameter is required.
	//
	// example:
	//
	// 8
	GpuPerWorker *int32 `json:"GpuPerWorker,omitempty" xml:"GpuPerWorker,omitempty"`
	// Memory GB allocation count
	//
	// This parameter is required.
	//
	// example:
	//
	// 500
	MemoryPerWorker *int32 `json:"MemoryPerWorker,omitempty" xml:"MemoryPerWorker,omitempty"`
	// NCCL version
	//
	// example:
	//
	// 1.0.0
	NCCLVersion *string `json:"NCCLVersion,omitempty" xml:"NCCLVersion,omitempty"`
	// PyTorch version
	//
	// example:
	//
	// 1.0.0
	PyTorchVersion *string `json:"PyTorchVersion,omitempty" xml:"PyTorchVersion,omitempty"`
	// Shared memory GB allocation count
	//
	// This parameter is required.
	//
	// example:
	//
	// 500
	ShareMemory *int32 `json:"ShareMemory,omitempty" xml:"ShareMemory,omitempty"`
	// Number of nodes
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	WorkerNum *int32 `json:"WorkerNum,omitempty" xml:"WorkerNum,omitempty"`
}

func (s UpdateExperimentPlanTemplateRequestTemplatePipelineEnvParams) String() string {
	return tea.Prettify(s)
}

func (s UpdateExperimentPlanTemplateRequestTemplatePipelineEnvParams) GoString() string {
	return s.String()
}

func (s *UpdateExperimentPlanTemplateRequestTemplatePipelineEnvParams) SetCpuPerWorker(v int32) *UpdateExperimentPlanTemplateRequestTemplatePipelineEnvParams {
	s.CpuPerWorker = &v
	return s
}

func (s *UpdateExperimentPlanTemplateRequestTemplatePipelineEnvParams) SetCudaVersion(v string) *UpdateExperimentPlanTemplateRequestTemplatePipelineEnvParams {
	s.CudaVersion = &v
	return s
}

func (s *UpdateExperimentPlanTemplateRequestTemplatePipelineEnvParams) SetGpuDriverVersion(v string) *UpdateExperimentPlanTemplateRequestTemplatePipelineEnvParams {
	s.GpuDriverVersion = &v
	return s
}

func (s *UpdateExperimentPlanTemplateRequestTemplatePipelineEnvParams) SetGpuPerWorker(v int32) *UpdateExperimentPlanTemplateRequestTemplatePipelineEnvParams {
	s.GpuPerWorker = &v
	return s
}

func (s *UpdateExperimentPlanTemplateRequestTemplatePipelineEnvParams) SetMemoryPerWorker(v int32) *UpdateExperimentPlanTemplateRequestTemplatePipelineEnvParams {
	s.MemoryPerWorker = &v
	return s
}

func (s *UpdateExperimentPlanTemplateRequestTemplatePipelineEnvParams) SetNCCLVersion(v string) *UpdateExperimentPlanTemplateRequestTemplatePipelineEnvParams {
	s.NCCLVersion = &v
	return s
}

func (s *UpdateExperimentPlanTemplateRequestTemplatePipelineEnvParams) SetPyTorchVersion(v string) *UpdateExperimentPlanTemplateRequestTemplatePipelineEnvParams {
	s.PyTorchVersion = &v
	return s
}

func (s *UpdateExperimentPlanTemplateRequestTemplatePipelineEnvParams) SetShareMemory(v int32) *UpdateExperimentPlanTemplateRequestTemplatePipelineEnvParams {
	s.ShareMemory = &v
	return s
}

func (s *UpdateExperimentPlanTemplateRequestTemplatePipelineEnvParams) SetWorkerNum(v int32) *UpdateExperimentPlanTemplateRequestTemplatePipelineEnvParams {
	s.WorkerNum = &v
	return s
}

type UpdateExperimentPlanTemplateShrinkRequest struct {
	// Template code
	//
	// This parameter is required.
	//
	// example:
	//
	// 349623
	TemplateId *int64 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// Template pipeline
	//
	// This parameter is required.
	TemplatePipelineShrink *string `json:"TemplatePipeline,omitempty" xml:"TemplatePipeline,omitempty"`
}

func (s UpdateExperimentPlanTemplateShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateExperimentPlanTemplateShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateExperimentPlanTemplateShrinkRequest) SetTemplateId(v int64) *UpdateExperimentPlanTemplateShrinkRequest {
	s.TemplateId = &v
	return s
}

func (s *UpdateExperimentPlanTemplateShrinkRequest) SetTemplatePipelineShrink(v string) *UpdateExperimentPlanTemplateShrinkRequest {
	s.TemplatePipelineShrink = &v
	return s
}

type UpdateExperimentPlanTemplateResponseBody struct {
	// Access denied information
	//
	// example:
	//
	// {}
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// Data
	Data *UpdateExperimentPlanTemplateResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// Request ID
	//
	// example:
	//
	// 5514CB39-B7C0-5B89-8534-2DE1E0F2B7AB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Total
	//
	// example:
	//
	// 0
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s UpdateExperimentPlanTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateExperimentPlanTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateExperimentPlanTemplateResponseBody) SetAccessDeniedDetail(v string) *UpdateExperimentPlanTemplateResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *UpdateExperimentPlanTemplateResponseBody) SetData(v *UpdateExperimentPlanTemplateResponseBodyData) *UpdateExperimentPlanTemplateResponseBody {
	s.Data = v
	return s
}

func (s *UpdateExperimentPlanTemplateResponseBody) SetRequestId(v string) *UpdateExperimentPlanTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateExperimentPlanTemplateResponseBody) SetTotalCount(v int64) *UpdateExperimentPlanTemplateResponseBody {
	s.TotalCount = &v
	return s
}

type UpdateExperimentPlanTemplateResponseBodyData struct {
	// Create Time
	//
	// example:
	//
	// 2024-10-22 10:18:10
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// Primary account UID
	//
	// example:
	//
	// 12312312312312
	CreatorUid *int64 `json:"CreatorUid,omitempty" xml:"CreatorUid,omitempty"`
	// Whether it is deleted
	//
	// example:
	//
	// 0
	IsDelete *int32 `json:"IsDelete,omitempty" xml:"IsDelete,omitempty"`
	// Privacy Level
	//
	// example:
	//
	// private
	PrivacyLevel *string `json:"PrivacyLevel,omitempty" xml:"PrivacyLevel,omitempty"`
	// Template code
	//
	// example:
	//
	// 472840184
	TemplateCode *int64 `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
	// Template Description
	//
	// example:
	//
	// test
	TemplateDescription *string `json:"TemplateDescription,omitempty" xml:"TemplateDescription,omitempty"`
	// Template ID
	//
	// example:
	//
	// 17815441
	TemplateId *int64 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// Template Name
	//
	// example:
	//
	// test
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// Template Pipeline
	TemplatePipelineParam []*UpdateExperimentPlanTemplateResponseBodyDataTemplatePipelineParam `json:"TemplatePipelineParam,omitempty" xml:"TemplatePipelineParam,omitempty" type:"Repeated"`
	// Update Time
	//
	// example:
	//
	// 2024-07-07 02:08:54
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// Version ID
	//
	// example:
	//
	// 1
	VersionId *int64 `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
}

func (s UpdateExperimentPlanTemplateResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s UpdateExperimentPlanTemplateResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateExperimentPlanTemplateResponseBodyData) SetCreateTime(v string) *UpdateExperimentPlanTemplateResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *UpdateExperimentPlanTemplateResponseBodyData) SetCreatorUid(v int64) *UpdateExperimentPlanTemplateResponseBodyData {
	s.CreatorUid = &v
	return s
}

func (s *UpdateExperimentPlanTemplateResponseBodyData) SetIsDelete(v int32) *UpdateExperimentPlanTemplateResponseBodyData {
	s.IsDelete = &v
	return s
}

func (s *UpdateExperimentPlanTemplateResponseBodyData) SetPrivacyLevel(v string) *UpdateExperimentPlanTemplateResponseBodyData {
	s.PrivacyLevel = &v
	return s
}

func (s *UpdateExperimentPlanTemplateResponseBodyData) SetTemplateCode(v int64) *UpdateExperimentPlanTemplateResponseBodyData {
	s.TemplateCode = &v
	return s
}

func (s *UpdateExperimentPlanTemplateResponseBodyData) SetTemplateDescription(v string) *UpdateExperimentPlanTemplateResponseBodyData {
	s.TemplateDescription = &v
	return s
}

func (s *UpdateExperimentPlanTemplateResponseBodyData) SetTemplateId(v int64) *UpdateExperimentPlanTemplateResponseBodyData {
	s.TemplateId = &v
	return s
}

func (s *UpdateExperimentPlanTemplateResponseBodyData) SetTemplateName(v string) *UpdateExperimentPlanTemplateResponseBodyData {
	s.TemplateName = &v
	return s
}

func (s *UpdateExperimentPlanTemplateResponseBodyData) SetTemplatePipelineParam(v []*UpdateExperimentPlanTemplateResponseBodyDataTemplatePipelineParam) *UpdateExperimentPlanTemplateResponseBodyData {
	s.TemplatePipelineParam = v
	return s
}

func (s *UpdateExperimentPlanTemplateResponseBodyData) SetUpdateTime(v string) *UpdateExperimentPlanTemplateResponseBodyData {
	s.UpdateTime = &v
	return s
}

func (s *UpdateExperimentPlanTemplateResponseBodyData) SetVersionId(v int64) *UpdateExperimentPlanTemplateResponseBodyData {
	s.VersionId = &v
	return s
}

type UpdateExperimentPlanTemplateResponseBodyDataTemplatePipelineParam struct {
	// Configured Environment Parameters
	EnvParams *UpdateExperimentPlanTemplateResponseBodyDataTemplatePipelineParamEnvParams `json:"EnvParams,omitempty" xml:"EnvParams,omitempty" type:"Struct"`
	// Node sequence number
	//
	// example:
	//
	// 1
	PipelineOrder *int32 `json:"PipelineOrder,omitempty" xml:"PipelineOrder,omitempty"`
	// Usage Scenario, e.g., "baseline"
	//
	// example:
	//
	// baseline
	Scene *string `json:"Scene,omitempty" xml:"Scene,omitempty"`
	// Configured Workload Parameters
	SettingParams map[string]*string `json:"SettingParams,omitempty" xml:"SettingParams,omitempty"`
	// Workload ID
	//
	// example:
	//
	// 13
	WorkloadId *int64 `json:"WorkloadId,omitempty" xml:"WorkloadId,omitempty"`
	// Workload Name
	//
	// example:
	//
	// test
	WorkloadName *string `json:"WorkloadName,omitempty" xml:"WorkloadName,omitempty"`
}

func (s UpdateExperimentPlanTemplateResponseBodyDataTemplatePipelineParam) String() string {
	return tea.Prettify(s)
}

func (s UpdateExperimentPlanTemplateResponseBodyDataTemplatePipelineParam) GoString() string {
	return s.String()
}

func (s *UpdateExperimentPlanTemplateResponseBodyDataTemplatePipelineParam) SetEnvParams(v *UpdateExperimentPlanTemplateResponseBodyDataTemplatePipelineParamEnvParams) *UpdateExperimentPlanTemplateResponseBodyDataTemplatePipelineParam {
	s.EnvParams = v
	return s
}

func (s *UpdateExperimentPlanTemplateResponseBodyDataTemplatePipelineParam) SetPipelineOrder(v int32) *UpdateExperimentPlanTemplateResponseBodyDataTemplatePipelineParam {
	s.PipelineOrder = &v
	return s
}

func (s *UpdateExperimentPlanTemplateResponseBodyDataTemplatePipelineParam) SetScene(v string) *UpdateExperimentPlanTemplateResponseBodyDataTemplatePipelineParam {
	s.Scene = &v
	return s
}

func (s *UpdateExperimentPlanTemplateResponseBodyDataTemplatePipelineParam) SetSettingParams(v map[string]*string) *UpdateExperimentPlanTemplateResponseBodyDataTemplatePipelineParam {
	s.SettingParams = v
	return s
}

func (s *UpdateExperimentPlanTemplateResponseBodyDataTemplatePipelineParam) SetWorkloadId(v int64) *UpdateExperimentPlanTemplateResponseBodyDataTemplatePipelineParam {
	s.WorkloadId = &v
	return s
}

func (s *UpdateExperimentPlanTemplateResponseBodyDataTemplatePipelineParam) SetWorkloadName(v string) *UpdateExperimentPlanTemplateResponseBodyDataTemplatePipelineParam {
	s.WorkloadName = &v
	return s
}

type UpdateExperimentPlanTemplateResponseBodyDataTemplatePipelineParamEnvParams struct {
	// CPU Allocation
	//
	// example:
	//
	// 90
	CpuPerWorker *int32 `json:"CpuPerWorker,omitempty" xml:"CpuPerWorker,omitempty"`
	// CUDA Version
	//
	// example:
	//
	// 1.0.0
	CudaVersion *string `json:"CudaVersion,omitempty" xml:"CudaVersion,omitempty"`
	// Extend Param
	ExtendParam map[string]*string `json:"ExtendParam,omitempty" xml:"ExtendParam,omitempty"`
	// GPU Driver Version
	//
	// example:
	//
	// 1.0.0
	GpuDriverVersion *string `json:"GpuDriverVersion,omitempty" xml:"GpuDriverVersion,omitempty"`
	// GPU Allocation
	//
	// example:
	//
	// 8
	GpuPerWorker *int32 `json:"GpuPerWorker,omitempty" xml:"GpuPerWorker,omitempty"`
	// Memory (GB) Allocation
	//
	// example:
	//
	// 500
	MemoryPerWorker *int32 `json:"MemoryPerWorker,omitempty" xml:"MemoryPerWorker,omitempty"`
	// NCCL Version
	//
	// example:
	//
	// 1.0.0
	NCCLVersion *string `json:"NCCLVersion,omitempty" xml:"NCCLVersion,omitempty"`
	// PyTorch Version
	//
	// example:
	//
	// 1.0.0
	PyTorchVersion *string `json:"PyTorchVersion,omitempty" xml:"PyTorchVersion,omitempty"`
	// Specified Nodes
	ResourceNodes []*UpdateExperimentPlanTemplateResponseBodyDataTemplatePipelineParamEnvParamsResourceNodes `json:"ResourceNodes,omitempty" xml:"ResourceNodes,omitempty" type:"Repeated"`
	// Shared Memory (GB) Allocation
	//
	// example:
	//
	// 500
	ShareMemory *int32 `json:"ShareMemory,omitempty" xml:"ShareMemory,omitempty"`
	// Number of Nodes
	//
	// example:
	//
	// 1
	WorkerNum *int32 `json:"WorkerNum,omitempty" xml:"WorkerNum,omitempty"`
}

func (s UpdateExperimentPlanTemplateResponseBodyDataTemplatePipelineParamEnvParams) String() string {
	return tea.Prettify(s)
}

func (s UpdateExperimentPlanTemplateResponseBodyDataTemplatePipelineParamEnvParams) GoString() string {
	return s.String()
}

func (s *UpdateExperimentPlanTemplateResponseBodyDataTemplatePipelineParamEnvParams) SetCpuPerWorker(v int32) *UpdateExperimentPlanTemplateResponseBodyDataTemplatePipelineParamEnvParams {
	s.CpuPerWorker = &v
	return s
}

func (s *UpdateExperimentPlanTemplateResponseBodyDataTemplatePipelineParamEnvParams) SetCudaVersion(v string) *UpdateExperimentPlanTemplateResponseBodyDataTemplatePipelineParamEnvParams {
	s.CudaVersion = &v
	return s
}

func (s *UpdateExperimentPlanTemplateResponseBodyDataTemplatePipelineParamEnvParams) SetExtendParam(v map[string]*string) *UpdateExperimentPlanTemplateResponseBodyDataTemplatePipelineParamEnvParams {
	s.ExtendParam = v
	return s
}

func (s *UpdateExperimentPlanTemplateResponseBodyDataTemplatePipelineParamEnvParams) SetGpuDriverVersion(v string) *UpdateExperimentPlanTemplateResponseBodyDataTemplatePipelineParamEnvParams {
	s.GpuDriverVersion = &v
	return s
}

func (s *UpdateExperimentPlanTemplateResponseBodyDataTemplatePipelineParamEnvParams) SetGpuPerWorker(v int32) *UpdateExperimentPlanTemplateResponseBodyDataTemplatePipelineParamEnvParams {
	s.GpuPerWorker = &v
	return s
}

func (s *UpdateExperimentPlanTemplateResponseBodyDataTemplatePipelineParamEnvParams) SetMemoryPerWorker(v int32) *UpdateExperimentPlanTemplateResponseBodyDataTemplatePipelineParamEnvParams {
	s.MemoryPerWorker = &v
	return s
}

func (s *UpdateExperimentPlanTemplateResponseBodyDataTemplatePipelineParamEnvParams) SetNCCLVersion(v string) *UpdateExperimentPlanTemplateResponseBodyDataTemplatePipelineParamEnvParams {
	s.NCCLVersion = &v
	return s
}

func (s *UpdateExperimentPlanTemplateResponseBodyDataTemplatePipelineParamEnvParams) SetPyTorchVersion(v string) *UpdateExperimentPlanTemplateResponseBodyDataTemplatePipelineParamEnvParams {
	s.PyTorchVersion = &v
	return s
}

func (s *UpdateExperimentPlanTemplateResponseBodyDataTemplatePipelineParamEnvParams) SetResourceNodes(v []*UpdateExperimentPlanTemplateResponseBodyDataTemplatePipelineParamEnvParamsResourceNodes) *UpdateExperimentPlanTemplateResponseBodyDataTemplatePipelineParamEnvParams {
	s.ResourceNodes = v
	return s
}

func (s *UpdateExperimentPlanTemplateResponseBodyDataTemplatePipelineParamEnvParams) SetShareMemory(v int32) *UpdateExperimentPlanTemplateResponseBodyDataTemplatePipelineParamEnvParams {
	s.ShareMemory = &v
	return s
}

func (s *UpdateExperimentPlanTemplateResponseBodyDataTemplatePipelineParamEnvParams) SetWorkerNum(v int32) *UpdateExperimentPlanTemplateResponseBodyDataTemplatePipelineParamEnvParams {
	s.WorkerNum = &v
	return s
}

type UpdateExperimentPlanTemplateResponseBodyDataTemplatePipelineParamEnvParamsResourceNodes struct {
	// Node Name
	//
	// example:
	//
	// honeypot
	NodeName *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	// Requested CPU
	//
	// example:
	//
	// 10
	RequestCPU *int32 `json:"RequestCPU,omitempty" xml:"RequestCPU,omitempty"`
	// Requested GPU
	//
	// example:
	//
	// 10
	RequestGPU *int32 `json:"RequestGPU,omitempty" xml:"RequestGPU,omitempty"`
	// Requested Memory
	//
	// example:
	//
	// 10
	RequestMemory *int32 `json:"RequestMemory,omitempty" xml:"RequestMemory,omitempty"`
	// Total CPU
	//
	// example:
	//
	// 100
	TotalCPU *int32 `json:"TotalCPU,omitempty" xml:"TotalCPU,omitempty"`
	// Total GPU
	//
	// example:
	//
	// 100
	TotalGPU *int32 `json:"TotalGPU,omitempty" xml:"TotalGPU,omitempty"`
	// Total Memory
	//
	// example:
	//
	// 100
	TotalMemory *int64 `json:"TotalMemory,omitempty" xml:"TotalMemory,omitempty"`
}

func (s UpdateExperimentPlanTemplateResponseBodyDataTemplatePipelineParamEnvParamsResourceNodes) String() string {
	return tea.Prettify(s)
}

func (s UpdateExperimentPlanTemplateResponseBodyDataTemplatePipelineParamEnvParamsResourceNodes) GoString() string {
	return s.String()
}

func (s *UpdateExperimentPlanTemplateResponseBodyDataTemplatePipelineParamEnvParamsResourceNodes) SetNodeName(v string) *UpdateExperimentPlanTemplateResponseBodyDataTemplatePipelineParamEnvParamsResourceNodes {
	s.NodeName = &v
	return s
}

func (s *UpdateExperimentPlanTemplateResponseBodyDataTemplatePipelineParamEnvParamsResourceNodes) SetRequestCPU(v int32) *UpdateExperimentPlanTemplateResponseBodyDataTemplatePipelineParamEnvParamsResourceNodes {
	s.RequestCPU = &v
	return s
}

func (s *UpdateExperimentPlanTemplateResponseBodyDataTemplatePipelineParamEnvParamsResourceNodes) SetRequestGPU(v int32) *UpdateExperimentPlanTemplateResponseBodyDataTemplatePipelineParamEnvParamsResourceNodes {
	s.RequestGPU = &v
	return s
}

func (s *UpdateExperimentPlanTemplateResponseBodyDataTemplatePipelineParamEnvParamsResourceNodes) SetRequestMemory(v int32) *UpdateExperimentPlanTemplateResponseBodyDataTemplatePipelineParamEnvParamsResourceNodes {
	s.RequestMemory = &v
	return s
}

func (s *UpdateExperimentPlanTemplateResponseBodyDataTemplatePipelineParamEnvParamsResourceNodes) SetTotalCPU(v int32) *UpdateExperimentPlanTemplateResponseBodyDataTemplatePipelineParamEnvParamsResourceNodes {
	s.TotalCPU = &v
	return s
}

func (s *UpdateExperimentPlanTemplateResponseBodyDataTemplatePipelineParamEnvParamsResourceNodes) SetTotalGPU(v int32) *UpdateExperimentPlanTemplateResponseBodyDataTemplatePipelineParamEnvParamsResourceNodes {
	s.TotalGPU = &v
	return s
}

func (s *UpdateExperimentPlanTemplateResponseBodyDataTemplatePipelineParamEnvParamsResourceNodes) SetTotalMemory(v int64) *UpdateExperimentPlanTemplateResponseBodyDataTemplatePipelineParamEnvParamsResourceNodes {
	s.TotalMemory = &v
	return s
}

type UpdateExperimentPlanTemplateResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateExperimentPlanTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateExperimentPlanTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateExperimentPlanTemplateResponse) GoString() string {
	return s.String()
}

func (s *UpdateExperimentPlanTemplateResponse) SetHeaders(v map[string]*string) *UpdateExperimentPlanTemplateResponse {
	s.Headers = v
	return s
}

func (s *UpdateExperimentPlanTemplateResponse) SetStatusCode(v int32) *UpdateExperimentPlanTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateExperimentPlanTemplateResponse) SetBody(v *UpdateExperimentPlanTemplateResponseBody) *UpdateExperimentPlanTemplateResponse {
	s.Body = v
	return s
}

type ValidateResourceRequest struct {
	// Resource ID
	//
	// This parameter is required.
	//
	// example:
	//
	// ehpc-sh-ouypm5aucy
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// User Authorization Parameters
	UserAccessParam *ValidateResourceRequestUserAccessParam `json:"UserAccessParam,omitempty" xml:"UserAccessParam,omitempty" type:"Struct"`
}

func (s ValidateResourceRequest) String() string {
	return tea.Prettify(s)
}

func (s ValidateResourceRequest) GoString() string {
	return s.String()
}

func (s *ValidateResourceRequest) SetClusterId(v string) *ValidateResourceRequest {
	s.ClusterId = &v
	return s
}

func (s *ValidateResourceRequest) SetUserAccessParam(v *ValidateResourceRequestUserAccessParam) *ValidateResourceRequest {
	s.UserAccessParam = v
	return s
}

type ValidateResourceRequestUserAccessParam struct {
	// User ID
	//
	// example:
	//
	// dev
	AccessId *string `json:"AccessId,omitempty" xml:"AccessId,omitempty"`
	// User Key
	//
	// example:
	//
	// test
	AccessKey *string `json:"AccessKey,omitempty" xml:"AccessKey,omitempty"`
	// Endpoint
	//
	// example:
	//
	// test
	Endpoint *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	// Workspace ID
	//
	// example:
	//
	// test
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ValidateResourceRequestUserAccessParam) String() string {
	return tea.Prettify(s)
}

func (s ValidateResourceRequestUserAccessParam) GoString() string {
	return s.String()
}

func (s *ValidateResourceRequestUserAccessParam) SetAccessId(v string) *ValidateResourceRequestUserAccessParam {
	s.AccessId = &v
	return s
}

func (s *ValidateResourceRequestUserAccessParam) SetAccessKey(v string) *ValidateResourceRequestUserAccessParam {
	s.AccessKey = &v
	return s
}

func (s *ValidateResourceRequestUserAccessParam) SetEndpoint(v string) *ValidateResourceRequestUserAccessParam {
	s.Endpoint = &v
	return s
}

func (s *ValidateResourceRequestUserAccessParam) SetWorkspaceId(v string) *ValidateResourceRequestUserAccessParam {
	s.WorkspaceId = &v
	return s
}

type ValidateResourceShrinkRequest struct {
	// Resource ID
	//
	// This parameter is required.
	//
	// example:
	//
	// ehpc-sh-ouypm5aucy
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// User Authorization Parameters
	UserAccessParamShrink *string `json:"UserAccessParam,omitempty" xml:"UserAccessParam,omitempty"`
}

func (s ValidateResourceShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ValidateResourceShrinkRequest) GoString() string {
	return s.String()
}

func (s *ValidateResourceShrinkRequest) SetClusterId(v string) *ValidateResourceShrinkRequest {
	s.ClusterId = &v
	return s
}

func (s *ValidateResourceShrinkRequest) SetUserAccessParamShrink(v string) *ValidateResourceShrinkRequest {
	s.UserAccessParamShrink = &v
	return s
}

type ValidateResourceResponseBody struct {
	// Data
	//
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// Request Id
	//
	// example:
	//
	// 5514CB39-B7C0-5B89-8534-2DE1E0F2B7AB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Total
	//
	// example:
	//
	// 0
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ValidateResourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ValidateResourceResponseBody) GoString() string {
	return s.String()
}

func (s *ValidateResourceResponseBody) SetData(v bool) *ValidateResourceResponseBody {
	s.Data = &v
	return s
}

func (s *ValidateResourceResponseBody) SetRequestId(v string) *ValidateResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ValidateResourceResponseBody) SetTotalCount(v int64) *ValidateResourceResponseBody {
	s.TotalCount = &v
	return s
}

type ValidateResourceResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ValidateResourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ValidateResourceResponse) String() string {
	return tea.Prettify(s)
}

func (s ValidateResourceResponse) GoString() string {
	return s.String()
}

func (s *ValidateResourceResponse) SetHeaders(v map[string]*string) *ValidateResourceResponse {
	s.Headers = v
	return s
}

func (s *ValidateResourceResponse) SetStatusCode(v int32) *ValidateResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *ValidateResourceResponse) SetBody(v *ValidateResourceResponseBody) *ValidateResourceResponse {
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
	client.EndpointRule = tea.String("")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("eflo-cnp"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// # Change resource group
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
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
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
		Version:     tea.String("2023-08-28"),
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
// # Change resource group
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
// deleteSlrEfloCnpForDeleting
//
// @param request - CheckServiceLinkedRoleEfloCnpForDeletingRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckServiceLinkedRoleEfloCnpForDeletingResponse
func (client *Client) CheckServiceLinkedRoleEfloCnpForDeletingWithOptions(request *CheckServiceLinkedRoleEfloCnpForDeletingRequest, runtime *util.RuntimeOptions) (_result *CheckServiceLinkedRoleEfloCnpForDeletingResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountId)) {
		query["AccountId"] = request.AccountId
	}

	if !tea.BoolValue(util.IsUnset(request.DeletionTaskId)) {
		query["DeletionTaskId"] = request.DeletionTaskId
	}

	if !tea.BoolValue(util.IsUnset(request.RoleArn)) {
		query["RoleArn"] = request.RoleArn
	}

	if !tea.BoolValue(util.IsUnset(request.SPIRegionId)) {
		query["SPIRegionId"] = request.SPIRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceName)) {
		query["ServiceName"] = request.ServiceName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CheckServiceLinkedRoleEfloCnpForDeleting"),
		Version:     tea.String("2023-08-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CheckServiceLinkedRoleEfloCnpForDeletingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// deleteSlrEfloCnpForDeleting
//
// @param request - CheckServiceLinkedRoleEfloCnpForDeletingRequest
//
// @return CheckServiceLinkedRoleEfloCnpForDeletingResponse
func (client *Client) CheckServiceLinkedRoleEfloCnpForDeleting(request *CheckServiceLinkedRoleEfloCnpForDeletingRequest) (_result *CheckServiceLinkedRoleEfloCnpForDeletingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CheckServiceLinkedRoleEfloCnpForDeletingResponse{}
	_body, _err := client.CheckServiceLinkedRoleEfloCnpForDeletingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Create Experiment Plan
//
// @param tmpReq - CreateExperimentPlanRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateExperimentPlanResponse
func (client *Client) CreateExperimentPlanWithOptions(tmpReq *CreateExperimentPlanRequest, runtime *util.RuntimeOptions) (_result *CreateExperimentPlanResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateExperimentPlanShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.ExternalParams)) {
		request.ExternalParamsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ExternalParams, tea.String("ExternalParams"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ExternalParamsShrink)) {
		query["ExternalParams"] = request.ExternalParamsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.PlanTemplateName)) {
		query["PlanTemplateName"] = request.PlanTemplateName
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateId)) {
		query["TemplateId"] = request.TemplateId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateExperimentPlan"),
		Version:     tea.String("2023-08-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateExperimentPlanResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Create Experiment Plan
//
// @param request - CreateExperimentPlanRequest
//
// @return CreateExperimentPlanResponse
func (client *Client) CreateExperimentPlan(request *CreateExperimentPlanRequest) (_result *CreateExperimentPlanResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateExperimentPlanResponse{}
	_body, _err := client.CreateExperimentPlanWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Create/Update Test Plan Template
//
// @param tmpReq - CreateExperimentPlanTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateExperimentPlanTemplateResponse
func (client *Client) CreateExperimentPlanTemplateWithOptions(tmpReq *CreateExperimentPlanTemplateRequest, runtime *util.RuntimeOptions) (_result *CreateExperimentPlanTemplateResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateExperimentPlanTemplateShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.TemplatePipeline)) {
		request.TemplatePipelineShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TemplatePipeline, tea.String("TemplatePipeline"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PrivacyLevel)) {
		query["PrivacyLevel"] = request.PrivacyLevel
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateDescription)) {
		query["TemplateDescription"] = request.TemplateDescription
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateId)) {
		query["TemplateId"] = request.TemplateId
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateName)) {
		query["TemplateName"] = request.TemplateName
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TemplatePipelineShrink)) {
		body["TemplatePipeline"] = request.TemplatePipelineShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateExperimentPlanTemplate"),
		Version:     tea.String("2023-08-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateExperimentPlanTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Create/Update Test Plan Template
//
// @param request - CreateExperimentPlanTemplateRequest
//
// @return CreateExperimentPlanTemplateResponse
func (client *Client) CreateExperimentPlanTemplate(request *CreateExperimentPlanTemplateRequest) (_result *CreateExperimentPlanTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateExperimentPlanTemplateResponse{}
	_body, _err := client.CreateExperimentPlanTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Create Evaluation Resource
//
// @param tmpReq - CreateResourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateResourceResponse
func (client *Client) CreateResourceWithOptions(tmpReq *CreateResourceRequest, runtime *util.RuntimeOptions) (_result *CreateResourceResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateResourceShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.MachineTypes)) {
		request.MachineTypesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.MachineTypes, tea.String("MachineTypes"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.UserAccessParam)) {
		request.UserAccessParamShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UserAccessParam, tea.String("UserAccessParam"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterDesc)) {
		query["ClusterDesc"] = request.ClusterDesc
	}

	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.ClusterName)) {
		query["ClusterName"] = request.ClusterName
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MachineTypesShrink)) {
		body["MachineTypes"] = request.MachineTypesShrink
	}

	if !tea.BoolValue(util.IsUnset(request.UserAccessParamShrink)) {
		body["UserAccessParam"] = request.UserAccessParamShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateResource"),
		Version:     tea.String("2023-08-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateResourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Create Evaluation Resource
//
// @param request - CreateResourceRequest
//
// @return CreateResourceResponse
func (client *Client) CreateResource(request *CreateResourceRequest) (_result *CreateResourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateResourceResponse{}
	_body, _err := client.CreateResourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Delete Experiment
//
// @param request - DeleteExperimentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteExperimentResponse
func (client *Client) DeleteExperimentWithOptions(request *DeleteExperimentRequest, runtime *util.RuntimeOptions) (_result *DeleteExperimentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ExperimentId)) {
		query["ExperimentId"] = request.ExperimentId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteExperiment"),
		Version:     tea.String("2023-08-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteExperimentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Delete Experiment
//
// @param request - DeleteExperimentRequest
//
// @return DeleteExperimentResponse
func (client *Client) DeleteExperiment(request *DeleteExperimentRequest) (_result *DeleteExperimentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteExperimentResponse{}
	_body, _err := client.DeleteExperimentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取实验计划详情
//
// @param request - DeleteExperimentPlanRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteExperimentPlanResponse
func (client *Client) DeleteExperimentPlanWithOptions(request *DeleteExperimentPlanRequest, runtime *util.RuntimeOptions) (_result *DeleteExperimentPlanResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PlanId)) {
		query["PlanId"] = request.PlanId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteExperimentPlan"),
		Version:     tea.String("2023-08-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteExperimentPlanResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取实验计划详情
//
// @param request - DeleteExperimentPlanRequest
//
// @return DeleteExperimentPlanResponse
func (client *Client) DeleteExperimentPlan(request *DeleteExperimentPlanRequest) (_result *DeleteExperimentPlanResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteExperimentPlanResponse{}
	_body, _err := client.DeleteExperimentPlanWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Delete Test Plan Template
//
// @param request - DeleteExperimentPlanTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteExperimentPlanTemplateResponse
func (client *Client) DeleteExperimentPlanTemplateWithOptions(request *DeleteExperimentPlanTemplateRequest, runtime *util.RuntimeOptions) (_result *DeleteExperimentPlanTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TemplateId)) {
		query["TemplateId"] = request.TemplateId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteExperimentPlanTemplate"),
		Version:     tea.String("2023-08-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteExperimentPlanTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Delete Test Plan Template
//
// @param request - DeleteExperimentPlanTemplateRequest
//
// @return DeleteExperimentPlanTemplateResponse
func (client *Client) DeleteExperimentPlanTemplate(request *DeleteExperimentPlanTemplateRequest) (_result *DeleteExperimentPlanTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteExperimentPlanTemplateResponse{}
	_body, _err := client.DeleteExperimentPlanTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Get Experiment Details
//
// @param request - GetExperimentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetExperimentResponse
func (client *Client) GetExperimentWithOptions(request *GetExperimentRequest, runtime *util.RuntimeOptions) (_result *GetExperimentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ExperimentId)) {
		query["ExperimentId"] = request.ExperimentId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetExperiment"),
		Version:     tea.String("2023-08-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetExperimentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Get Experiment Details
//
// @param request - GetExperimentRequest
//
// @return GetExperimentResponse
func (client *Client) GetExperiment(request *GetExperimentRequest) (_result *GetExperimentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetExperimentResponse{}
	_body, _err := client.GetExperimentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Get Experiment Plan Details
//
// @param request - GetExperimentPlanRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetExperimentPlanResponse
func (client *Client) GetExperimentPlanWithOptions(request *GetExperimentPlanRequest, runtime *util.RuntimeOptions) (_result *GetExperimentPlanResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PlanId)) {
		query["PlanId"] = request.PlanId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetExperimentPlan"),
		Version:     tea.String("2023-08-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetExperimentPlanResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Get Experiment Plan Details
//
// @param request - GetExperimentPlanRequest
//
// @return GetExperimentPlanResponse
func (client *Client) GetExperimentPlan(request *GetExperimentPlanRequest) (_result *GetExperimentPlanResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetExperimentPlanResponse{}
	_body, _err := client.GetExperimentPlanWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Query Test Plan Template Details
//
// @param request - GetExperimentPlanTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetExperimentPlanTemplateResponse
func (client *Client) GetExperimentPlanTemplateWithOptions(request *GetExperimentPlanTemplateRequest, runtime *util.RuntimeOptions) (_result *GetExperimentPlanTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TemplateId)) {
		query["TemplateId"] = request.TemplateId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetExperimentPlanTemplate"),
		Version:     tea.String("2023-08-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetExperimentPlanTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query Test Plan Template Details
//
// @param request - GetExperimentPlanTemplateRequest
//
// @return GetExperimentPlanTemplateResponse
func (client *Client) GetExperimentPlanTemplate(request *GetExperimentPlanTemplateRequest) (_result *GetExperimentPlanTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetExperimentPlanTemplateResponse{}
	_body, _err := client.GetExperimentPlanTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Fetch Experiment Result Data
//
// @param request - GetExperimentResultDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetExperimentResultDataResponse
func (client *Client) GetExperimentResultDataWithOptions(request *GetExperimentResultDataRequest, runtime *util.RuntimeOptions) (_result *GetExperimentResultDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ExperimentId)) {
		query["ExperimentId"] = request.ExperimentId
	}

	if !tea.BoolValue(util.IsUnset(request.Hostname)) {
		query["Hostname"] = request.Hostname
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.WorkloadType)) {
		query["WorkloadType"] = request.WorkloadType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetExperimentResultData"),
		Version:     tea.String("2023-08-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetExperimentResultDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Fetch Experiment Result Data
//
// @param request - GetExperimentResultDataRequest
//
// @return GetExperimentResultDataResponse
func (client *Client) GetExperimentResultData(request *GetExperimentResultDataRequest) (_result *GetExperimentResultDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetExperimentResultDataResponse{}
	_body, _err := client.GetExperimentResultDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Get Resource Information
//
// @param request - GetResourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetResourceResponse
func (client *Client) GetResourceWithOptions(request *GetResourceRequest, runtime *util.RuntimeOptions) (_result *GetResourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetResource"),
		Version:     tea.String("2023-08-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetResourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Get Resource Information
//
// @param request - GetResourceRequest
//
// @return GetResourceResponse
func (client *Client) GetResource(request *GetResourceRequest) (_result *GetResourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetResourceResponse{}
	_body, _err := client.GetResourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Query the resource prediction results of the test plan template
//
// @param request - GetResourcePredictResultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetResourcePredictResultResponse
func (client *Client) GetResourcePredictResultWithOptions(request *GetResourcePredictResultRequest, runtime *util.RuntimeOptions) (_result *GetResourcePredictResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateId)) {
		query["TemplateId"] = request.TemplateId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetResourcePredictResult"),
		Version:     tea.String("2023-08-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetResourcePredictResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query the resource prediction results of the test plan template
//
// @param request - GetResourcePredictResultRequest
//
// @return GetResourcePredictResultResponse
func (client *Client) GetResourcePredictResult(request *GetResourcePredictResultRequest) (_result *GetResourcePredictResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetResourcePredictResultResponse{}
	_body, _err := client.GetResourcePredictResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Retrieve workload information by ID
//
// @param request - GetWorkloadRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetWorkloadResponse
func (client *Client) GetWorkloadWithOptions(request *GetWorkloadRequest, runtime *util.RuntimeOptions) (_result *GetWorkloadResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.WorkloadId)) {
		query["WorkloadId"] = request.WorkloadId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetWorkload"),
		Version:     tea.String("2023-08-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetWorkloadResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Retrieve workload information by ID
//
// @param request - GetWorkloadRequest
//
// @return GetWorkloadResponse
func (client *Client) GetWorkload(request *GetWorkloadRequest) (_result *GetWorkloadResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetWorkloadResponse{}
	_body, _err := client.GetWorkloadWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Query Test Plan Template List
//
// @param request - ListExperimentPlanTemplatesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListExperimentPlanTemplatesResponse
func (client *Client) ListExperimentPlanTemplatesWithOptions(request *ListExperimentPlanTemplatesRequest, runtime *util.RuntimeOptions) (_result *ListExperimentPlanTemplatesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PrivacyLevel)) {
		query["PrivacyLevel"] = request.PrivacyLevel
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListExperimentPlanTemplates"),
		Version:     tea.String("2023-08-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListExperimentPlanTemplatesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query Test Plan Template List
//
// @param request - ListExperimentPlanTemplatesRequest
//
// @return ListExperimentPlanTemplatesResponse
func (client *Client) ListExperimentPlanTemplates(request *ListExperimentPlanTemplatesRequest) (_result *ListExperimentPlanTemplatesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListExperimentPlanTemplatesResponse{}
	_body, _err := client.ListExperimentPlanTemplatesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Query Experiment Plan List
//
// @param tmpReq - ListExperimentPlansRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListExperimentPlansResponse
func (client *Client) ListExperimentPlansWithOptions(tmpReq *ListExperimentPlansRequest, runtime *util.RuntimeOptions) (_result *ListExperimentPlansResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ListExperimentPlansShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.PlanTaskStatus)) {
		request.PlanTaskStatusShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.PlanTaskStatus, tea.String("PlanTaskStatus"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.ResourceName)) {
		request.ResourceNameShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ResourceName, tea.String("ResourceName"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Tag)) {
		request.TagShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tag, tea.String("Tag"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CreatTimeOrder)) {
		query["CreatTimeOrder"] = request.CreatTimeOrder
	}

	if !tea.BoolValue(util.IsUnset(request.EndTimeOrder)) {
		query["EndTimeOrder"] = request.EndTimeOrder
	}

	if !tea.BoolValue(util.IsUnset(request.Page)) {
		query["Page"] = request.Page
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.Size)) {
		query["Size"] = request.Size
	}

	if !tea.BoolValue(util.IsUnset(request.StartTimeOrder)) {
		query["StartTimeOrder"] = request.StartTimeOrder
	}

	if !tea.BoolValue(util.IsUnset(request.TagShrink)) {
		query["Tag"] = request.TagShrink
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateId)) {
		query["TemplateId"] = request.TemplateId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PlanTaskStatusShrink)) {
		body["PlanTaskStatus"] = request.PlanTaskStatusShrink
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceNameShrink)) {
		body["ResourceName"] = request.ResourceNameShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListExperimentPlans"),
		Version:     tea.String("2023-08-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListExperimentPlansResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query Experiment Plan List
//
// @param request - ListExperimentPlansRequest
//
// @return ListExperimentPlansResponse
func (client *Client) ListExperimentPlans(request *ListExperimentPlansRequest) (_result *ListExperimentPlansResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListExperimentPlansResponse{}
	_body, _err := client.ListExperimentPlansWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Query the experiment list based on the plan ID
//
// @param request - ListExperimentsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListExperimentsResponse
func (client *Client) ListExperimentsWithOptions(request *ListExperimentsRequest, runtime *util.RuntimeOptions) (_result *ListExperimentsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Order)) {
		query["Order"] = request.Order
	}

	if !tea.BoolValue(util.IsUnset(request.PlanId)) {
		query["PlanId"] = request.PlanId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListExperiments"),
		Version:     tea.String("2023-08-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListExperimentsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query the experiment list based on the plan ID
//
// @param request - ListExperimentsRequest
//
// @return ListExperimentsResponse
func (client *Client) ListExperiments(request *ListExperimentsRequest) (_result *ListExperimentsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListExperimentsResponse{}
	_body, _err := client.ListExperimentsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Query Resource Tags
//
// @param request - ListTagResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTagResourcesResponse
func (client *Client) ListTagResourcesWithOptions(request *ListTagResourcesRequest, runtime *util.RuntimeOptions) (_result *ListTagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTagResources"),
		Version:     tea.String("2023-08-28"),
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
// # Query Resource Tags
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
// # Get Workload List
//
// @param request - ListWorkloadsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListWorkloadsResponse
func (client *Client) ListWorkloadsWithOptions(request *ListWorkloadsRequest, runtime *util.RuntimeOptions) (_result *ListWorkloadsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Scope)) {
		query["Scope"] = request.Scope
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListWorkloads"),
		Version:     tea.String("2023-08-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListWorkloadsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Get Workload List
//
// @param request - ListWorkloadsRequest
//
// @return ListWorkloadsResponse
func (client *Client) ListWorkloads(request *ListWorkloadsRequest) (_result *ListWorkloadsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListWorkloadsResponse{}
	_body, _err := client.ListWorkloadsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Stop Experiment
//
// @param request - StopExperimentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopExperimentResponse
func (client *Client) StopExperimentWithOptions(request *StopExperimentRequest, runtime *util.RuntimeOptions) (_result *StopExperimentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ExperimentId)) {
		query["ExperimentId"] = request.ExperimentId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("StopExperiment"),
		Version:     tea.String("2023-08-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &StopExperimentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Stop Experiment
//
// @param request - StopExperimentRequest
//
// @return StopExperimentResponse
func (client *Client) StopExperiment(request *StopExperimentRequest) (_result *StopExperimentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StopExperimentResponse{}
	_body, _err := client.StopExperimentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Tag Resources with User Labels
//
// @param request - TagResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TagResourcesResponse
func (client *Client) TagResourcesWithOptions(request *TagResourcesRequest, runtime *util.RuntimeOptions) (_result *TagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("TagResources"),
		Version:     tea.String("2023-08-28"),
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
// # Tag Resources with User Labels
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
// # Remove User Tags from Resources
//
// @param request - UntagResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UntagResourcesResponse
func (client *Client) UntagResourcesWithOptions(request *UntagResourcesRequest, runtime *util.RuntimeOptions) (_result *UntagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.All)) {
		query["All"] = request.All
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.TagKey)) {
		query["TagKey"] = request.TagKey
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UntagResources"),
		Version:     tea.String("2023-08-28"),
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
// # Remove User Tags from Resources
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
// # Update Experiment Plan
//
// @param request - UpdateExperimentPlanRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateExperimentPlanResponse
func (client *Client) UpdateExperimentPlanWithOptions(request *UpdateExperimentPlanRequest, runtime *util.RuntimeOptions) (_result *UpdateExperimentPlanResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PlanId)) {
		query["PlanId"] = request.PlanId
	}

	if !tea.BoolValue(util.IsUnset(request.PlanTemplateName)) {
		query["PlanTemplateName"] = request.PlanTemplateName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateExperimentPlan"),
		Version:     tea.String("2023-08-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateExperimentPlanResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Update Experiment Plan
//
// @param request - UpdateExperimentPlanRequest
//
// @return UpdateExperimentPlanResponse
func (client *Client) UpdateExperimentPlan(request *UpdateExperimentPlanRequest) (_result *UpdateExperimentPlanResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateExperimentPlanResponse{}
	_body, _err := client.UpdateExperimentPlanWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Update Test Plan Template
//
// @param tmpReq - UpdateExperimentPlanTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateExperimentPlanTemplateResponse
func (client *Client) UpdateExperimentPlanTemplateWithOptions(tmpReq *UpdateExperimentPlanTemplateRequest, runtime *util.RuntimeOptions) (_result *UpdateExperimentPlanTemplateResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &UpdateExperimentPlanTemplateShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.TemplatePipeline)) {
		request.TemplatePipelineShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TemplatePipeline, tea.String("TemplatePipeline"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TemplateId)) {
		query["TemplateId"] = request.TemplateId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TemplatePipelineShrink)) {
		body["TemplatePipeline"] = request.TemplatePipelineShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateExperimentPlanTemplate"),
		Version:     tea.String("2023-08-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateExperimentPlanTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Update Test Plan Template
//
// @param request - UpdateExperimentPlanTemplateRequest
//
// @return UpdateExperimentPlanTemplateResponse
func (client *Client) UpdateExperimentPlanTemplate(request *UpdateExperimentPlanTemplateRequest) (_result *UpdateExperimentPlanTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateExperimentPlanTemplateResponse{}
	_body, _err := client.UpdateExperimentPlanTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Resource Connectivity Test
//
// @param tmpReq - ValidateResourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ValidateResourceResponse
func (client *Client) ValidateResourceWithOptions(tmpReq *ValidateResourceRequest, runtime *util.RuntimeOptions) (_result *ValidateResourceResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ValidateResourceShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.UserAccessParam)) {
		request.UserAccessParamShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UserAccessParam, tea.String("UserAccessParam"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		query["ClusterId"] = request.ClusterId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UserAccessParamShrink)) {
		body["UserAccessParam"] = request.UserAccessParamShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ValidateResource"),
		Version:     tea.String("2023-08-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ValidateResourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Resource Connectivity Test
//
// @param request - ValidateResourceRequest
//
// @return ValidateResourceResponse
func (client *Client) ValidateResource(request *ValidateResourceRequest) (_result *ValidateResourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ValidateResourceResponse{}
	_body, _err := client.ValidateResourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
