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
	// example:
	//
	// 54
	ExperimentId *int64 `json:"ExperimentId,omitempty" xml:"ExperimentId,omitempty"`
	// example:
	//
	// p-jt-waf-app1
	Hostname *string `json:"Hostname,omitempty" xml:"Hostname,omitempty"`
	// example:
	//
	// fluxserv-6fc89b45cf-w8wq6
	PodName *string `json:"PodName,omitempty" xml:"PodName,omitempty"`
	// example:
	//
	// 8
	GpuNum *int32 `json:"GpuNum,omitempty" xml:"GpuNum,omitempty"`
	// example:
	//
	// 8x OAM 810 GPU
	GpuName *string `json:"GpuName,omitempty" xml:"GpuName,omitempty"`
	// example:
	//
	// false
	WarningFlag *bool   `json:"WarningFlag,omitempty" xml:"WarningFlag,omitempty"`
	WarningMsg  *string `json:"WarningMsg,omitempty" xml:"WarningMsg,omitempty"`
	// example:
	//
	// false
	ErrorFlag *bool   `json:"ErrorFlag,omitempty" xml:"ErrorFlag,omitempty"`
	ErrorMsg  *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	// TFLOPS
	//
	// example:
	//
	// 45
	Tflops *float64 `json:"Tflops,omitempty" xml:"Tflops,omitempty"`
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

type CreateExperimentPlanRequest struct {
	// Additional parameters
	//
	// example:
	//
	// {}
	ExternalParams map[string]interface{} `json:"ExternalParams,omitempty" xml:"ExternalParams,omitempty"`
	// Resource group ID
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
	// Template ID
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

func (s *CreateExperimentPlanRequest) SetResourceGroupId(v string) *CreateExperimentPlanRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateExperimentPlanRequest) SetResourceId(v int64) *CreateExperimentPlanRequest {
	s.ResourceId = &v
	return s
}

func (s *CreateExperimentPlanRequest) SetTemplateId(v int64) *CreateExperimentPlanRequest {
	s.TemplateId = &v
	return s
}

type CreateExperimentPlanShrinkRequest struct {
	// Additional parameters
	//
	// example:
	//
	// {}
	ExternalParamsShrink *string `json:"ExternalParams,omitempty" xml:"ExternalParams,omitempty"`
	// Resource group ID
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
	// Template ID
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

func (s *CreateExperimentPlanShrinkRequest) SetResourceGroupId(v string) *CreateExperimentPlanShrinkRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateExperimentPlanShrinkRequest) SetResourceId(v int64) *CreateExperimentPlanShrinkRequest {
	s.ResourceId = &v
	return s
}

func (s *CreateExperimentPlanShrinkRequest) SetTemplateId(v int64) *CreateExperimentPlanShrinkRequest {
	s.TemplateId = &v
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
	// Total count of the query
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
	// example:
	//
	// private
	PrivacyLevel *string `json:"PrivacyLevel,omitempty" xml:"PrivacyLevel,omitempty"`
	// example:
	//
	// The template installs jdk and tomcat on a new ECS instance.
	TemplateDescription *string `json:"TemplateDescription,omitempty" xml:"TemplateDescription,omitempty"`
	// example:
	//
	// 4724
	TemplateId       *int64                                                 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	TemplateName     *string                                                `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
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
	EnvParams *CreateExperimentPlanTemplateRequestTemplatePipelineEnvParams `json:"EnvParams,omitempty" xml:"EnvParams,omitempty" type:"Struct"`
	// example:
	//
	// 1
	PipelineOrder *int32 `json:"PipelineOrder,omitempty" xml:"PipelineOrder,omitempty"`
	// example:
	//
	// baseline
	Scene         *string            `json:"Scene,omitempty" xml:"Scene,omitempty"`
	SettingParams map[string]*string `json:"SettingParams,omitempty" xml:"SettingParams,omitempty"`
	// example:
	//
	// 14
	WorkloadId *int64 `json:"WorkloadId,omitempty" xml:"WorkloadId,omitempty"`
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
	// GpuDriverVersion
	//
	// example:
	//
	// 1.0.0
	GpuDriverVersion *string `json:"GpuDriverVersion,omitempty" xml:"GpuDriverVersion,omitempty"`
	// example:
	//
	// 8
	GpuPerWorker *int32 `json:"GpuPerWorker,omitempty" xml:"GpuPerWorker,omitempty"`
	// example:
	//
	// 500
	MemoryPerWorker *int32 `json:"MemoryPerWorker,omitempty" xml:"MemoryPerWorker,omitempty"`
	// NCCLVersion
	//
	// example:
	//
	// 1.0.0
	NCCLVersion *string `json:"NCCLVersion,omitempty" xml:"NCCLVersion,omitempty"`
	// PyTorchVersion
	//
	// example:
	//
	// 1.0.0
	PyTorchVersion *string `json:"PyTorchVersion,omitempty" xml:"PyTorchVersion,omitempty"`
	// example:
	//
	// 500
	ShareMemory *int32 `json:"ShareMemory,omitempty" xml:"ShareMemory,omitempty"`
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
	// example:
	//
	// private
	PrivacyLevel *string `json:"PrivacyLevel,omitempty" xml:"PrivacyLevel,omitempty"`
	// example:
	//
	// The template installs jdk and tomcat on a new ECS instance.
	TemplateDescription *string `json:"TemplateDescription,omitempty" xml:"TemplateDescription,omitempty"`
	// example:
	//
	// 4724
	TemplateId             *int64  `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	TemplateName           *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
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
	// example:
	//
	// {}
	AccessDeniedDetail *string                                       `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	Data               *CreateExperimentPlanTemplateResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 5514CB39-B7C0-5B89-8534-2DE1E0F2B7AB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// example:
	//
	// 2024-11-19T02:01:05Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 12312312312312
	CreatorUid *int64 `json:"CreatorUid,omitempty" xml:"CreatorUid,omitempty"`
	// example:
	//
	// 0
	IsDelete *int32 `json:"IsDelete,omitempty" xml:"IsDelete,omitempty"`
	// example:
	//
	// private
	PrivacyLevel        *string `json:"PrivacyLevel,omitempty" xml:"PrivacyLevel,omitempty"`
	TemplateDescription *string `json:"TemplateDescription,omitempty" xml:"TemplateDescription,omitempty"`
	// example:
	//
	// 17615126
	TemplateId            *int64                                                               `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	TemplateName          *string                                                              `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	TemplatePipelineParam []*CreateExperimentPlanTemplateResponseBodyDataTemplatePipelineParam `json:"TemplatePipelineParam,omitempty" xml:"TemplatePipelineParam,omitempty" type:"Repeated"`
	// example:
	//
	// 2023-10-16T01:58Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
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
	EnvParams *CreateExperimentPlanTemplateResponseBodyDataTemplatePipelineParamEnvParams `json:"EnvParams,omitempty" xml:"EnvParams,omitempty" type:"Struct"`
	// example:
	//
	// 1
	PipelineOrder *int32 `json:"PipelineOrder,omitempty" xml:"PipelineOrder,omitempty"`
	// example:
	//
	// baseline
	Scene         *string            `json:"Scene,omitempty" xml:"Scene,omitempty"`
	SettingParams map[string]*string `json:"SettingParams,omitempty" xml:"SettingParams,omitempty"`
	// example:
	//
	// 13
	WorkloadId *int64 `json:"WorkloadId,omitempty" xml:"WorkloadId,omitempty"`
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
	// GpuDriverVersion
	//
	// example:
	//
	// 1.0.0
	GpuDriverVersion *string `json:"GpuDriverVersion,omitempty" xml:"GpuDriverVersion,omitempty"`
	// example:
	//
	// 8
	GpuPerWorker *int32 `json:"GpuPerWorker,omitempty" xml:"GpuPerWorker,omitempty"`
	// example:
	//
	// 500
	MemoryPerWorker *int32 `json:"MemoryPerWorker,omitempty" xml:"MemoryPerWorker,omitempty"`
	// NCCLVersion
	//
	// example:
	//
	// 1.0.0
	NCCLVersion *string `json:"NCCLVersion,omitempty" xml:"NCCLVersion,omitempty"`
	// PyTorchVersion
	//
	// example:
	//
	// 1.0.0
	PyTorchVersion *string `json:"PyTorchVersion,omitempty" xml:"PyTorchVersion,omitempty"`
	// example:
	//
	// 500
	ShareMemory *int32 `json:"ShareMemory,omitempty" xml:"ShareMemory,omitempty"`
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
	// example:
	//
	// ehpc-sh-fj71c0ycfw
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// Cluster Name
	//
	// example:
	//
	// tre-1-ppu
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// Cluster Type
	//
	// example:
	//
	// ACK
	ClusterType *string `json:"ClusterType,omitempty" xml:"ClusterType,omitempty"`
	// Machine Types
	MachineTypes *CreateResourceRequestMachineTypes `json:"MachineTypes,omitempty" xml:"MachineTypes,omitempty" type:"Struct"`
	// Resource Type
	//
	// example:
	//
	// ACK
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// User Access Parameters
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

func (s *CreateResourceRequest) SetClusterType(v string) *CreateResourceRequest {
	s.ClusterType = &v
	return s
}

func (s *CreateResourceRequest) SetMachineTypes(v *CreateResourceRequestMachineTypes) *CreateResourceRequest {
	s.MachineTypes = v
	return s
}

func (s *CreateResourceRequest) SetResourceType(v string) *CreateResourceRequest {
	s.ResourceType = &v
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
	// example:
	//
	// ehpc-sh-fj71c0ycfw
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// Cluster Name
	//
	// example:
	//
	// tre-1-ppu
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// Cluster Type
	//
	// example:
	//
	// ACK
	ClusterType *string `json:"ClusterType,omitempty" xml:"ClusterType,omitempty"`
	// Machine Types
	MachineTypesShrink *string `json:"MachineTypes,omitempty" xml:"MachineTypes,omitempty"`
	// Resource Type
	//
	// example:
	//
	// ACK
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// User Access Parameters
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

func (s *CreateResourceShrinkRequest) SetClusterType(v string) *CreateResourceShrinkRequest {
	s.ClusterType = &v
	return s
}

func (s *CreateResourceShrinkRequest) SetMachineTypesShrink(v string) *CreateResourceShrinkRequest {
	s.MachineTypesShrink = &v
	return s
}

func (s *CreateResourceShrinkRequest) SetResourceType(v string) *CreateResourceShrinkRequest {
	s.ResourceType = &v
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
	// example:
	//
	// 234
	ExperimentId *int64 `json:"ExperimentId,omitempty" xml:"ExperimentId,omitempty"`
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

type DeleteExperimentResponseBody struct {
	// example:
	//
	// {}
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// []
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 5514CB39-B7C0-5B89-8534-2DE1E0F2B7AB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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

type DeleteExperimentPlanTemplateRequest struct {
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
	// example:
	//
	// []
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 4D3FD55F-3BCD-5914-9B74-A1F4961327E7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// example:
	//
	// 234
	ExperimentId *int64 `json:"ExperimentId,omitempty" xml:"ExperimentId,omitempty"`
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

type GetExperimentResponseBody struct {
	// example:
	//
	// {}
	AccessDeniedDetail *string                        `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	Data               *GetExperimentResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// E67E2E4C-2B47-5C55-AA17-1D771E070AEF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// example:
	//
	// 2024-11-29 02:16:35
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 2024-11-29 02:26:35
	EndTime   *string                                 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	EnvParams *GetExperimentResponseBodyDataEnvParams `json:"EnvParams,omitempty" xml:"EnvParams,omitempty" type:"Struct"`
	// example:
	//
	// 1726882991828688898
	ExperimentId *int64 `json:"ExperimentId,omitempty" xml:"ExperimentId,omitempty"`
	// example:
	//
	// test
	ExperimentName *string `json:"ExperimentName,omitempty" xml:"ExperimentName,omitempty"`
	// example:
	//
	// AI
	ExperimentType *string                                `json:"ExperimentType,omitempty" xml:"ExperimentType,omitempty"`
	GetParams      map[string]*string                     `json:"GetParams,omitempty" xml:"GetParams,omitempty"`
	Resource       *GetExperimentResponseBodyDataResource `json:"Resource,omitempty" xml:"Resource,omitempty" type:"Struct"`
	// example:
	//
	// cifnews-guoyuan
	ResourceName *string                               `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	Results      *GetExperimentResponseBodyDataResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Struct"`
	SetParams    map[string]*string                    `json:"SetParams,omitempty" xml:"SetParams,omitempty"`
	// example:
	//
	// 2024-11-29 02:16:35
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// RUNNING
	Status *string                            `json:"Status,omitempty" xml:"Status,omitempty"`
	Task   *GetExperimentResponseBodyDataTask `json:"Task,omitempty" xml:"Task,omitempty" type:"Struct"`
	// example:
	//
	// 2024-11-29 02:16:35
	UpdateTime *int64                                 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	Workload   *GetExperimentResponseBodyDataWorkload `json:"Workload,omitempty" xml:"Workload,omitempty" type:"Struct"`
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
	// example:
	//
	// 90
	CpuPerWorker *int32 `json:"CpuPerWorker,omitempty" xml:"CpuPerWorker,omitempty"`
	// cudaVersion
	//
	// example:
	//
	// 1.0.0
	CudaVersion *string            `json:"CudaVersion,omitempty" xml:"CudaVersion,omitempty"`
	ExtendParam map[string]*string `json:"ExtendParam,omitempty" xml:"ExtendParam,omitempty"`
	// GpuDriverVersion
	//
	// example:
	//
	// 1.0.0
	GpuDriverVersion *string `json:"GpuDriverVersion,omitempty" xml:"GpuDriverVersion,omitempty"`
	// example:
	//
	// 8
	GpuPerWorker *int32 `json:"GpuPerWorker,omitempty" xml:"GpuPerWorker,omitempty"`
	// example:
	//
	// 500
	MemoryPerWorker *int32 `json:"MemoryPerWorker,omitempty" xml:"MemoryPerWorker,omitempty"`
	// NCCLVersion
	//
	// example:
	//
	// 1.0.0
	NCCLVersion *string `json:"NCCLVersion,omitempty" xml:"NCCLVersion,omitempty"`
	// PyTorchVersion
	//
	// example:
	//
	// 1.0.0
	PyTorchVersion *string                                                `json:"PyTorchVersion,omitempty" xml:"PyTorchVersion,omitempty"`
	ResourceNodes  []*GetExperimentResponseBodyDataEnvParamsResourceNodes `json:"ResourceNodes,omitempty" xml:"ResourceNodes,omitempty" type:"Repeated"`
	// example:
	//
	// 500
	ShareMemory *int32 `json:"ShareMemory,omitempty" xml:"ShareMemory,omitempty"`
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
	// example:
	//
	// p-jt-waf-app1
	NodeName *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	// example:
	//
	// 90
	RequestCPU *int32 `json:"RequestCPU,omitempty" xml:"RequestCPU,omitempty"`
	// example:
	//
	// 8
	RequestGPU *int32 `json:"RequestGPU,omitempty" xml:"RequestGPU,omitempty"`
	// example:
	//
	// 500
	RequestMemory *int32 `json:"RequestMemory,omitempty" xml:"RequestMemory,omitempty"`
	// example:
	//
	// 90
	TotalCPU *int32 `json:"TotalCPU,omitempty" xml:"TotalCPU,omitempty"`
	// example:
	//
	// 8
	TotalGPU *int32 `json:"TotalGPU,omitempty" xml:"TotalGPU,omitempty"`
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
	// example:
	//
	// 90
	CpuCoreLimit *int32 `json:"CpuCoreLimit,omitempty" xml:"CpuCoreLimit,omitempty"`
	// example:
	//
	// 8
	GpuLimit    *int32                                            `json:"GpuLimit,omitempty" xml:"GpuLimit,omitempty"`
	MachineType *GetExperimentResponseBodyDataResourceMachineType `json:"MachineType,omitempty" xml:"MachineType,omitempty" type:"Struct"`
	// example:
	//
	// 90
	MaxCpuCore *int32 `json:"MaxCpuCore,omitempty" xml:"MaxCpuCore,omitempty"`
	// example:
	//
	// 8
	MaxGpu *int32 `json:"MaxGpu,omitempty" xml:"MaxGpu,omitempty"`
	// example:
	//
	// 500
	MaxMemory *int64 `json:"MaxMemory,omitempty" xml:"MaxMemory,omitempty"`
	// example:
	//
	// 500
	MemoryLimit *int64 `json:"MemoryLimit,omitempty" xml:"MemoryLimit,omitempty"`
	// example:
	//
	// 189
	ResourceId *int64 `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// example:
	//
	// ecs.r8y.4xlarge
	ResourceName    *string                                               `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	ResourceNodes   []*GetExperimentResponseBodyDataResourceResourceNodes `json:"ResourceNodes,omitempty" xml:"ResourceNodes,omitempty" type:"Repeated"`
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
	// example:
	//
	// 5
	BondNum *int32 `json:"BondNum,omitempty" xml:"BondNum,omitempty"`
	// example:
	//
	// 2x Intel Icelake 8369B 32C CPU
	CpuInfo *string `json:"CpuInfo,omitempty" xml:"CpuInfo,omitempty"`
	// example:
	//
	// 2x 480GB SATA SSD \\n 4x 3.84TB NVMe SSD
	DiskInfo *string `json:"DiskInfo,omitempty" xml:"DiskInfo,omitempty"`
	// example:
	//
	// 8x NVIDIA SXM4 80GB A100 GPU
	GpuInfo *string `json:"GpuInfo,omitempty" xml:"GpuInfo,omitempty"`
	// example:
	//
	// 32x 64GB DDR4 3200 Memory
	MemoryInfo *string `json:"MemoryInfo,omitempty" xml:"MemoryInfo,omitempty"`
	// example:
	//
	// efg1.nvga1n
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 1x 100Gbps DP NIC for VPC \\n 4x 100Gbps DP RoCE NIC
	NetworkInfo *string `json:"NetworkInfo,omitempty" xml:"NetworkInfo,omitempty"`
	// example:
	//
	// 2
	NetworkMode *string `json:"NetworkMode,omitempty" xml:"NetworkMode,omitempty"`
	// example:
	//
	// 1
	NodeCount *int32 `json:"NodeCount,omitempty" xml:"NodeCount,omitempty"`
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
	// example:
	//
	// dev
	AccessId *string `json:"AccessId,omitempty" xml:"AccessId,omitempty"`
	// example:
	//
	// test
	AccessKey *string `json:"AccessKey,omitempty" xml:"AccessKey,omitempty"`
	// endpoint
	//
	// example:
	//
	// test
	Endpoint *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
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
	// example:
	//
	// 764
	Duration    *float64                                           `json:"Duration,omitempty" xml:"Duration,omitempty"`
	ErrorWorker []*GetExperimentResponseBodyDataResultsErrorWorker `json:"ErrorWorker,omitempty" xml:"ErrorWorker,omitempty" type:"Repeated"`
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
	// example:
	//
	// 10
	SamplesPerSecond *float64 `json:"SamplesPerSecond,omitempty" xml:"SamplesPerSecond,omitempty"`
	// example:
	//
	// 1000
	SecondsPerIteration      *float64                                                        `json:"SecondsPerIteration,omitempty" xml:"SecondsPerIteration,omitempty"`
	TaskIndividualResultList []*GetExperimentResponseBodyDataResultsTaskIndividualResultList `json:"TaskIndividualResultList,omitempty" xml:"TaskIndividualResultList,omitempty" type:"Repeated"`
	TaskIndividualResultMap  map[string][]*DataResultsTaskIndividualResultMapValue           `json:"TaskIndividualResultMap,omitempty" xml:"TaskIndividualResultMap,omitempty"`
	WarningBoundList         []*GetExperimentResponseBodyDataResultsWarningBoundList         `json:"WarningBoundList,omitempty" xml:"WarningBoundList,omitempty" type:"Repeated"`
	WarningWorker            []*GetExperimentResponseBodyDataResultsWarningWorker            `json:"WarningWorker,omitempty" xml:"WarningWorker,omitempty" type:"Repeated"`
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
	// example:
	//
	// true
	ErrorFlag *bool `json:"ErrorFlag,omitempty" xml:"ErrorFlag,omitempty"`
	// example:
	//
	// Connection reset
	ErrorMsg *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	// example:
	//
	// 97
	ExperimentId *int64 `json:"ExperimentId,omitempty" xml:"ExperimentId,omitempty"`
	// example:
	//
	// 8x OAM 810 GPU
	GpuName *string `json:"GpuName,omitempty" xml:"GpuName,omitempty"`
	// example:
	//
	// 8
	GpuNum *int32 `json:"GpuNum,omitempty" xml:"GpuNum,omitempty"`
	// example:
	//
	// 60.188.98.209
	Hostname *string `json:"Hostname,omitempty" xml:"Hostname,omitempty"`
	// example:
	//
	// hzs-forge-sdxl-online-7ff4d86444-pc95h
	PodName *string `json:"PodName,omitempty" xml:"PodName,omitempty"`
	// example:
	//
	// 23
	SamplesPerSecond *float64 `json:"SamplesPerSecond,omitempty" xml:"SamplesPerSecond,omitempty"`
	// example:
	//
	// 12
	Tflops *float64 `json:"Tflops,omitempty" xml:"Tflops,omitempty"`
	// example:
	//
	// false
	WarningFlag *bool   `json:"WarningFlag,omitempty" xml:"WarningFlag,omitempty"`
	WarningMsg  *string `json:"WarningMsg,omitempty" xml:"WarningMsg,omitempty"`
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
	// example:
	//
	// false
	ErrorFlag *bool   `json:"ErrorFlag,omitempty" xml:"ErrorFlag,omitempty"`
	ErrorMsg  *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	// example:
	//
	// 48
	ExperimentId *int64 `json:"ExperimentId,omitempty" xml:"ExperimentId,omitempty"`
	// example:
	//
	// 8x OAM 810 GPU
	GpuName *string `json:"GpuName,omitempty" xml:"GpuName,omitempty"`
	// example:
	//
	// 8
	GpuNum *int32 `json:"GpuNum,omitempty" xml:"GpuNum,omitempty"`
	// example:
	//
	// p-jt-waf-app1
	Hostname *string `json:"Hostname,omitempty" xml:"Hostname,omitempty"`
	// example:
	//
	// fluxserv-6fc89b45cf-w8wq6
	PodName *string `json:"PodName,omitempty" xml:"PodName,omitempty"`
	// example:
	//
	// 28
	SamplesPerSecond *float64 `json:"SamplesPerSecond,omitempty" xml:"SamplesPerSecond,omitempty"`
	// example:
	//
	// 16
	Tflops *float64 `json:"Tflops,omitempty" xml:"Tflops,omitempty"`
	// example:
	//
	// false
	WarningFlag *bool   `json:"WarningFlag,omitempty" xml:"WarningFlag,omitempty"`
	WarningMsg  *string `json:"WarningMsg,omitempty" xml:"WarningMsg,omitempty"`
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
	// example:
	//
	// true
	ErrorFlag *bool   `json:"ErrorFlag,omitempty" xml:"ErrorFlag,omitempty"`
	ErrorMsg  *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	// example:
	//
	// 9
	ExperimentId *int64 `json:"ExperimentId,omitempty" xml:"ExperimentId,omitempty"`
	// example:
	//
	// 8x OAM 810 GPU
	GpuName *string `json:"GpuName,omitempty" xml:"GpuName,omitempty"`
	// example:
	//
	// 8
	GpuNum *int32 `json:"GpuNum,omitempty" xml:"GpuNum,omitempty"`
	// example:
	//
	// whza008403
	Hostname *string `json:"Hostname,omitempty" xml:"Hostname,omitempty"`
	// example:
	//
	// fluxserv-6fc89b45cf-w8wq6
	PodName *string `json:"PodName,omitempty" xml:"PodName,omitempty"`
	// example:
	//
	// 15
	SamplesPerSecond *float64 `json:"SamplesPerSecond,omitempty" xml:"SamplesPerSecond,omitempty"`
	// example:
	//
	// 14
	Tflops *float64 `json:"Tflops,omitempty" xml:"Tflops,omitempty"`
	// example:
	//
	// true
	WarningFlag *bool   `json:"WarningFlag,omitempty" xml:"WarningFlag,omitempty"`
	WarningMsg  *string `json:"WarningMsg,omitempty" xml:"WarningMsg,omitempty"`
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
	// example:
	//
	// 2024-03-05 18:24:08
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 2024-03-05 18:34:08
	EndTime *int64             `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Params  map[string]*string `json:"Params,omitempty" xml:"Params,omitempty"`
	// example:
	//
	// baseline
	Scene *string `json:"Scene,omitempty" xml:"Scene,omitempty"`
	// example:
	//
	// 2024-03-05 18:24:08
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 167420
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
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
	// example:
	//
	// 90
	DefaultCpuPerWorker *int32 `json:"DefaultCpuPerWorker,omitempty" xml:"DefaultCpuPerWorker,omitempty"`
	// example:
	//
	// 8
	DefaultGpuPerWorker *int32 `json:"DefaultGpuPerWorker,omitempty" xml:"DefaultGpuPerWorker,omitempty"`
	// example:
	//
	// 500
	DefaultMemoryPerWorker *int32 `json:"DefaultMemoryPerWorker,omitempty" xml:"DefaultMemoryPerWorker,omitempty"`
	// example:
	//
	// 500
	DefaultShareMemory *int32 `json:"DefaultShareMemory,omitempty" xml:"DefaultShareMemory,omitempty"`
	// example:
	//
	// AI
	Family *string `json:"Family,omitempty" xml:"Family,omitempty"`
	// example:
	//
	// PyTorchJob
	JobKind       *string                                               `json:"JobKind,omitempty" xml:"JobKind,omitempty"`
	ParamSettings []*GetExperimentResponseBodyDataWorkloadParamSettings `json:"ParamSettings,omitempty" xml:"ParamSettings,omitempty" type:"Repeated"`
	// example:
	//
	// NLP-LLM
	Scene *string `json:"Scene,omitempty" xml:"Scene,omitempty"`
	// example:
	//
	// common
	Scope        *string                                            `json:"Scope,omitempty" xml:"Scope,omitempty"`
	StaticConfig *GetExperimentResponseBodyDataWorkloadStaticConfig `json:"StaticConfig,omitempty" xml:"StaticConfig,omitempty" type:"Struct"`
	// example:
	//
	// 1
	VersionId *int64 `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
	// example:
	//
	// test
	WorkloadDescription *string `json:"WorkloadDescription,omitempty" xml:"WorkloadDescription,omitempty"`
	// example:
	//
	// 13
	WorkloadId *int64 `json:"WorkloadId,omitempty" xml:"WorkloadId,omitempty"`
	// example:
	//
	// test
	WorkloadName *string `json:"WorkloadName,omitempty" xml:"WorkloadName,omitempty"`
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
	// example:
	//
	// 100
	DefaultValue *string `json:"DefaultValue,omitempty" xml:"DefaultValue,omitempty"`
	ParamDesc    *string `json:"ParamDesc,omitempty" xml:"ParamDesc,omitempty"`
	// example:
	//
	// ITERATION
	ParamName *string `json:"ParamName,omitempty" xml:"ParamName,omitempty"`
	// example:
	//
	// [0-9]+
	ParamRegex *string `json:"ParamRegex,omitempty" xml:"ParamRegex,omitempty"`
	// example:
	//
	// number
	ParamType *string `json:"ParamType,omitempty" xml:"ParamType,omitempty"`
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
	// example:
	//
	// pyTorch
	FrameWork *string `json:"FrameWork,omitempty" xml:"FrameWork,omitempty"`
	// example:
	//
	// linux
	Os *string `json:"Os,omitempty" xml:"Os,omitempty"`
	// example:
	//
	// 7B
	Parameters *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
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

type GetExperimentResultDataRequest struct {
	// example:
	//
	// 234
	ExperimentId *int64 `json:"ExperimentId,omitempty" xml:"ExperimentId,omitempty"`
	// example:
	//
	// iZj6ccwd7zwfms7hzaz2riZ
	Hostname *string `json:"Hostname,omitempty" xml:"Hostname,omitempty"`
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

func (s *GetExperimentResultDataRequest) SetWorkloadType(v string) *GetExperimentResultDataRequest {
	s.WorkloadType = &v
	return s
}

type GetExperimentResultDataResponseBody struct {
	// example:
	//
	// {}
	AccessDeniedDetail *string                                    `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	Data               []*GetExperimentResultDataResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// C1D34EC2-AB13-56F4-8322-F15AE563EA04
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// example:
	//
	// 8
	GpuNum *string `json:"GpuNum,omitempty" xml:"GpuNum,omitempty"`
	// example:
	//
	// p-jt-waf-app1
	Hostname     *string                                                `json:"Hostname,omitempty" xml:"Hostname,omitempty"`
	MetricsInfos []*GetExperimentResultDataResponseBodyDataMetricsInfos `json:"MetricsInfos,omitempty" xml:"MetricsInfos,omitempty" type:"Repeated"`
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
	// gpu
	//
	// example:
	//
	// 8
	GpuNum *string `json:"Gpu_num,omitempty" xml:"Gpu_num,omitempty"`
	// iteration
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
	// example:
	//
	// 1715393860
	Timestamp *int64 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
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
	// Cluster name
	//
	// example:
	//
	// ecs.g6.4xlarge
	ResourceName *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
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

func (s *GetResourceResponseBodyData) SetResourceName(v string) *GetResourceResponseBodyData {
	s.ResourceName = &v
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
	// example:
	//
	// 36
	ResourceId *int64 `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
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
	// example:
	//
	// 2
	Data *int64 `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 5514CB39-B7C0-5B89-8534-2DE1E0F2B7AB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// example:
	//
	// {}
	AccessDeniedDetail *string                      `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	Data               *GetWorkloadResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// E67E2E4C-2B47-5C55-AA17-1D771E070AEF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// example:
	//
	// 90
	DefaultCpuPerWorker *int32 `json:"DefaultCpuPerWorker,omitempty" xml:"DefaultCpuPerWorker,omitempty"`
	// example:
	//
	// 8
	DefaultGpuPerWorker *int32 `json:"DefaultGpuPerWorker,omitempty" xml:"DefaultGpuPerWorker,omitempty"`
	// example:
	//
	// 500
	DefaultMemoryPerWorker *int32 `json:"DefaultMemoryPerWorker,omitempty" xml:"DefaultMemoryPerWorker,omitempty"`
	// example:
	//
	// 500
	DefaultShareMemory *int32 `json:"DefaultShareMemory,omitempty" xml:"DefaultShareMemory,omitempty"`
	// example:
	//
	// AI
	Family *string `json:"Family,omitempty" xml:"Family,omitempty"`
	// example:
	//
	// PyTorchJob
	JobKind       *string                                     `json:"JobKind,omitempty" xml:"JobKind,omitempty"`
	ParamSettings []*GetWorkloadResponseBodyDataParamSettings `json:"ParamSettings,omitempty" xml:"ParamSettings,omitempty" type:"Repeated"`
	// example:
	//
	// NLP-LLM
	Scene *string `json:"Scene,omitempty" xml:"Scene,omitempty"`
	// example:
	//
	// common
	Scope        *string                                  `json:"Scope,omitempty" xml:"Scope,omitempty"`
	StaticConfig *GetWorkloadResponseBodyDataStaticConfig `json:"StaticConfig,omitempty" xml:"StaticConfig,omitempty" type:"Struct"`
	// example:
	//
	// 1
	VersionId *int64 `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
	// example:
	//
	// test
	WorkloadDescription *string `json:"WorkloadDescription,omitempty" xml:"WorkloadDescription,omitempty"`
	// example:
	//
	// 13
	WorkloadId *int64 `json:"WorkloadId,omitempty" xml:"WorkloadId,omitempty"`
	// example:
	//
	// test
	WorkloadName *string `json:"WorkloadName,omitempty" xml:"WorkloadName,omitempty"`
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
	// example:
	//
	// 100
	DefaultValue *string `json:"DefaultValue,omitempty" xml:"DefaultValue,omitempty"`
	ParamDesc    *string `json:"ParamDesc,omitempty" xml:"ParamDesc,omitempty"`
	// example:
	//
	// ITERATION
	ParamName *string `json:"ParamName,omitempty" xml:"ParamName,omitempty"`
	// example:
	//
	// [0-9]+
	ParamRegex *string `json:"ParamRegex,omitempty" xml:"ParamRegex,omitempty"`
	// example:
	//
	// number
	ParamType *string `json:"ParamType,omitempty" xml:"ParamType,omitempty"`
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
	// example:
	//
	// PyTorch
	FrameWork *string `json:"FrameWork,omitempty" xml:"FrameWork,omitempty"`
	// example:
	//
	// linux
	Os *string `json:"Os,omitempty" xml:"Os,omitempty"`
	// example:
	//
	// 7B
	Parameters *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
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

func (s *ListExperimentPlanTemplatesResponseBodyDataTemplatePipelineParamEnvParams) SetGpuPerWorker(v int32) *ListExperimentPlanTemplatesResponseBodyDataTemplatePipelineParamEnvParams {
	s.GpuPerWorker = &v
	return s
}

func (s *ListExperimentPlanTemplatesResponseBodyDataTemplatePipelineParamEnvParams) SetMemoryPerWorker(v int32) *ListExperimentPlanTemplatesResponseBodyDataTemplatePipelineParamEnvParams {
	s.MemoryPerWorker = &v
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

func (s *ListExperimentPlansResponseBodyData) SetResourceName(v string) *ListExperimentPlansResponseBodyData {
	s.ResourceName = &v
	return s
}

func (s *ListExperimentPlansResponseBodyData) SetStartTime(v string) *ListExperimentPlansResponseBodyData {
	s.StartTime = &v
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

type ListWorkloadsRequest struct {
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
	// example:
	//
	// {}
	AccessDeniedDetail *string                          `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	Data               []*ListWorkloadsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// 4AC08332-436C-57A3-9FBA-26772B1A9901
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// example:
	//
	// 90
	DefaultCpuPerWorker *int32 `json:"DefaultCpuPerWorker,omitempty" xml:"DefaultCpuPerWorker,omitempty"`
	// example:
	//
	// 8
	DefaultGpuPerWorker *int32 `json:"DefaultGpuPerWorker,omitempty" xml:"DefaultGpuPerWorker,omitempty"`
	// example:
	//
	// 500
	DefaultMemoryPerWorker *int32 `json:"DefaultMemoryPerWorker,omitempty" xml:"DefaultMemoryPerWorker,omitempty"`
	// example:
	//
	// 500
	DefaultShareMemory *int32 `json:"DefaultShareMemory,omitempty" xml:"DefaultShareMemory,omitempty"`
	// example:
	//
	// AI
	Family *string `json:"Family,omitempty" xml:"Family,omitempty"`
	// example:
	//
	// PyTorchJob
	JobKind       *string                                       `json:"JobKind,omitempty" xml:"JobKind,omitempty"`
	ParamSettings []*ListWorkloadsResponseBodyDataParamSettings `json:"ParamSettings,omitempty" xml:"ParamSettings,omitempty" type:"Repeated"`
	// example:
	//
	// NLP-LLM
	Scene *string `json:"Scene,omitempty" xml:"Scene,omitempty"`
	// example:
	//
	// common
	Scope        *string                                    `json:"Scope,omitempty" xml:"Scope,omitempty"`
	StaticConfig *ListWorkloadsResponseBodyDataStaticConfig `json:"StaticConfig,omitempty" xml:"StaticConfig,omitempty" type:"Struct"`
	// example:
	//
	// 1
	VersionId *int64 `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
	// example:
	//
	// test
	WorkloadDescription *string `json:"WorkloadDescription,omitempty" xml:"WorkloadDescription,omitempty"`
	// example:
	//
	// 13
	WorkloadId *int64 `json:"WorkloadId,omitempty" xml:"WorkloadId,omitempty"`
	// example:
	//
	// test
	WorkloadName *string `json:"WorkloadName,omitempty" xml:"WorkloadName,omitempty"`
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
	// example:
	//
	// 100
	DefaultValue *string `json:"DefaultValue,omitempty" xml:"DefaultValue,omitempty"`
	ParamDesc    *string `json:"ParamDesc,omitempty" xml:"ParamDesc,omitempty"`
	// example:
	//
	// ITERATION
	ParamName *string `json:"ParamName,omitempty" xml:"ParamName,omitempty"`
	// example:
	//
	// [0-9]+
	ParamRegex *string `json:"ParamRegex,omitempty" xml:"ParamRegex,omitempty"`
	// example:
	//
	// number
	ParamType *string `json:"ParamType,omitempty" xml:"ParamType,omitempty"`
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
	// example:
	//
	// PyTorch
	FrameWork *string `json:"FrameWork,omitempty" xml:"FrameWork,omitempty"`
	// example:
	//
	// linux
	Os *string `json:"Os,omitempty" xml:"Os,omitempty"`
	// example:
	//
	// 7B
	Parameters *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
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
	// example:
	//
	// 234
	ExperimentId *int64 `json:"ExperimentId,omitempty" xml:"ExperimentId,omitempty"`
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

type StopExperimentResponseBody struct {
	// example:
	//
	// {}
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 5514CB39-B7C0-5B89-8534-2DE1E0F2B7AB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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

type ValidateResourceRequest struct {
	// Resource ID
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
// 资源转组
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
// 资源转组
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
// Create Experiment Plan
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

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

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
// Create Experiment Plan
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
// 创建/更新测试计划模板
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
// 创建/更新测试计划模板
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
// Create Evaluation Resource
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

	if !tea.BoolValue(util.IsUnset(request.ClusterType)) {
		query["ClusterType"] = request.ClusterType
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
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
// Create Evaluation Resource
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
// 删除实验
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
// 删除实验
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
// 删除测试计划模板
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
// 删除测试计划模板
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
// 获取实验详情
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
// 获取实验详情
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
// Get Experiment Plan Details
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
// Get Experiment Plan Details
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
// 获取实验结果数据
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
// 获取实验结果数据
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
// Get Resource Information
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
// Get Resource Information
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
// 查询测试计划模板资源预测结果
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
// 查询测试计划模板资源预测结果
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
// 通过id获取负载信息
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
// 通过id获取负载信息
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
// Query Test Plan Template List
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
// Query Test Plan Template List
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
// Query Experiment Plan List
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

	if !tea.BoolValue(util.IsUnset(request.Size)) {
		query["Size"] = request.Size
	}

	if !tea.BoolValue(util.IsUnset(request.StartTimeOrder)) {
		query["StartTimeOrder"] = request.StartTimeOrder
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
// Query Experiment Plan List
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
// Query the experiment list based on the plan ID
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
// Query the experiment list based on the plan ID
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
// 获取负载列表
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
// 获取负载列表
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
// 停止实验
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
// 停止实验
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
// Resource Connectivity Test
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
// Resource Connectivity Test
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
