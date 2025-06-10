// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type RepoConfiguration struct {
	ArtifactBuildRuleParameters *RepoConfigurationArtifactBuildRuleParameters `json:"ArtifactBuildRuleParameters,omitempty" xml:"ArtifactBuildRuleParameters,omitempty" type:"Struct"`
	// This parameter is required.
	RepoType *string `json:"RepoType,omitempty" xml:"RepoType,omitempty"`
	// This parameter is required.
	TagImmutability *bool `json:"TagImmutability,omitempty" xml:"TagImmutability,omitempty"`
}

func (s RepoConfiguration) String() string {
	return tea.Prettify(s)
}

func (s RepoConfiguration) GoString() string {
	return s.String()
}

func (s *RepoConfiguration) SetArtifactBuildRuleParameters(v *RepoConfigurationArtifactBuildRuleParameters) *RepoConfiguration {
	s.ArtifactBuildRuleParameters = v
	return s
}

func (s *RepoConfiguration) SetRepoType(v string) *RepoConfiguration {
	s.RepoType = &v
	return s
}

func (s *RepoConfiguration) SetTagImmutability(v bool) *RepoConfiguration {
	s.TagImmutability = &v
	return s
}

type RepoConfigurationArtifactBuildRuleParameters struct {
	// This parameter is required.
	ImageIndexOnly *bool   `json:"ImageIndexOnly,omitempty" xml:"ImageIndexOnly,omitempty"`
	PriorityFile   *string `json:"PriorityFile,omitempty" xml:"PriorityFile,omitempty"`
}

func (s RepoConfigurationArtifactBuildRuleParameters) String() string {
	return tea.Prettify(s)
}

func (s RepoConfigurationArtifactBuildRuleParameters) GoString() string {
	return s.String()
}

func (s *RepoConfigurationArtifactBuildRuleParameters) SetImageIndexOnly(v bool) *RepoConfigurationArtifactBuildRuleParameters {
	s.ImageIndexOnly = &v
	return s
}

func (s *RepoConfigurationArtifactBuildRuleParameters) SetPriorityFile(v string) *RepoConfigurationArtifactBuildRuleParameters {
	s.PriorityFile = &v
	return s
}

type CancelArtifactBuildTaskRequest struct {
	// The ID of the artifact building task.
	//
	// This parameter is required.
	//
	// example:
	//
	// i2ei-12****
	BuildTaskId *string `json:"BuildTaskId,omitempty" xml:"BuildTaskId,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cri-shac42yvqzvq****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s CancelArtifactBuildTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s CancelArtifactBuildTaskRequest) GoString() string {
	return s.String()
}

func (s *CancelArtifactBuildTaskRequest) SetBuildTaskId(v string) *CancelArtifactBuildTaskRequest {
	s.BuildTaskId = &v
	return s
}

func (s *CancelArtifactBuildTaskRequest) SetInstanceId(v string) *CancelArtifactBuildTaskRequest {
	s.InstanceId = &v
	return s
}

type CancelArtifactBuildTaskResponseBody struct {
	// The return value.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// 	- `true`: The request is successful.
	//
	// 	- `false`: The request fails.
	//
	// example:
	//
	// true
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// C4C7DD0C-C9D6-437A-A7EE-121EFD70D002
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CancelArtifactBuildTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CancelArtifactBuildTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CancelArtifactBuildTaskResponseBody) SetCode(v string) *CancelArtifactBuildTaskResponseBody {
	s.Code = &v
	return s
}

func (s *CancelArtifactBuildTaskResponseBody) SetIsSuccess(v bool) *CancelArtifactBuildTaskResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *CancelArtifactBuildTaskResponseBody) SetRequestId(v string) *CancelArtifactBuildTaskResponseBody {
	s.RequestId = &v
	return s
}

type CancelArtifactBuildTaskResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CancelArtifactBuildTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CancelArtifactBuildTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s CancelArtifactBuildTaskResponse) GoString() string {
	return s.String()
}

func (s *CancelArtifactBuildTaskResponse) SetHeaders(v map[string]*string) *CancelArtifactBuildTaskResponse {
	s.Headers = v
	return s
}

func (s *CancelArtifactBuildTaskResponse) SetStatusCode(v int32) *CancelArtifactBuildTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelArtifactBuildTaskResponse) SetBody(v *CancelArtifactBuildTaskResponseBody) *CancelArtifactBuildTaskResponse {
	s.Body = v
	return s
}

type CancelRepoBuildRecordRequest struct {
	// The ID of the image building record.
	//
	// This parameter is required.
	//
	// example:
	//
	// 74FDBA62-30C0-4F22-BE7B-F1D36FD1****
	BuildRecordId *string `json:"BuildRecordId,omitempty" xml:"BuildRecordId,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cri-kmsiwlxxdcva****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the image repository.
	//
	// This parameter is required.
	//
	// example:
	//
	// crr-tquyps22md8p****
	RepoId *string `json:"RepoId,omitempty" xml:"RepoId,omitempty"`
}

func (s CancelRepoBuildRecordRequest) String() string {
	return tea.Prettify(s)
}

func (s CancelRepoBuildRecordRequest) GoString() string {
	return s.String()
}

func (s *CancelRepoBuildRecordRequest) SetBuildRecordId(v string) *CancelRepoBuildRecordRequest {
	s.BuildRecordId = &v
	return s
}

func (s *CancelRepoBuildRecordRequest) SetInstanceId(v string) *CancelRepoBuildRecordRequest {
	s.InstanceId = &v
	return s
}

func (s *CancelRepoBuildRecordRequest) SetRepoId(v string) *CancelRepoBuildRecordRequest {
	s.RepoId = &v
	return s
}

type CancelRepoBuildRecordResponseBody struct {
	// The return value.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 4CE1F661-75DD-4EBD-A4AD-057B26834ABB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CancelRepoBuildRecordResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CancelRepoBuildRecordResponseBody) GoString() string {
	return s.String()
}

func (s *CancelRepoBuildRecordResponseBody) SetCode(v string) *CancelRepoBuildRecordResponseBody {
	s.Code = &v
	return s
}

func (s *CancelRepoBuildRecordResponseBody) SetIsSuccess(v bool) *CancelRepoBuildRecordResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *CancelRepoBuildRecordResponseBody) SetRequestId(v string) *CancelRepoBuildRecordResponseBody {
	s.RequestId = &v
	return s
}

type CancelRepoBuildRecordResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CancelRepoBuildRecordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CancelRepoBuildRecordResponse) String() string {
	return tea.Prettify(s)
}

func (s CancelRepoBuildRecordResponse) GoString() string {
	return s.String()
}

func (s *CancelRepoBuildRecordResponse) SetHeaders(v map[string]*string) *CancelRepoBuildRecordResponse {
	s.Headers = v
	return s
}

func (s *CancelRepoBuildRecordResponse) SetStatusCode(v int32) *CancelRepoBuildRecordResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelRepoBuildRecordResponse) SetBody(v *CancelRepoBuildRecordResponseBody) *CancelRepoBuildRecordResponse {
	s.Body = v
	return s
}

type CancelRepoSyncTaskRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cri-xkx6vujuhay0****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the replication task.
	//
	// This parameter is required.
	//
	// example:
	//
	// rst-biu4u4pm4it5****
	SyncTaskId *string `json:"SyncTaskId,omitempty" xml:"SyncTaskId,omitempty"`
}

func (s CancelRepoSyncTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s CancelRepoSyncTaskRequest) GoString() string {
	return s.String()
}

func (s *CancelRepoSyncTaskRequest) SetInstanceId(v string) *CancelRepoSyncTaskRequest {
	s.InstanceId = &v
	return s
}

func (s *CancelRepoSyncTaskRequest) SetSyncTaskId(v string) *CancelRepoSyncTaskRequest {
	s.SyncTaskId = &v
	return s
}

type CancelRepoSyncTaskResponseBody struct {
	// The response code.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the CancelRepoSyncTask request is successful. Valid values:
	//
	// 	- `true`: successful
	//
	// 	- `false`: failed
	//
	// example:
	//
	// true
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The request ID.
	//
	// example:
	//
	// EB9C5722-51E2-4497-A573-575B0CA5CE0C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CancelRepoSyncTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CancelRepoSyncTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CancelRepoSyncTaskResponseBody) SetCode(v string) *CancelRepoSyncTaskResponseBody {
	s.Code = &v
	return s
}

func (s *CancelRepoSyncTaskResponseBody) SetIsSuccess(v bool) *CancelRepoSyncTaskResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *CancelRepoSyncTaskResponseBody) SetRequestId(v string) *CancelRepoSyncTaskResponseBody {
	s.RequestId = &v
	return s
}

type CancelRepoSyncTaskResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CancelRepoSyncTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CancelRepoSyncTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s CancelRepoSyncTaskResponse) GoString() string {
	return s.String()
}

func (s *CancelRepoSyncTaskResponse) SetHeaders(v map[string]*string) *CancelRepoSyncTaskResponse {
	s.Headers = v
	return s
}

func (s *CancelRepoSyncTaskResponse) SetStatusCode(v int32) *CancelRepoSyncTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelRepoSyncTaskResponse) SetBody(v *CancelRepoSyncTaskResponseBody) *CancelRepoSyncTaskResponse {
	s.Body = v
	return s
}

type ChangeResourceGroupRequest struct {
	// The ID of the resource group to which you want to move the resource.
	//
	// This parameter is required.
	//
	// example:
	//
	// rg-aekz5nlvlaksnvi
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The ID of the resource.
	//
	// This parameter is required.
	//
	// example:
	//
	// cri-8qong6ve5p3mhlgt
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The ID of the region.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shenzhen-finance-1
	ResourceRegionId *string `json:"ResourceRegionId,omitempty" xml:"ResourceRegionId,omitempty"`
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

func (s *ChangeResourceGroupRequest) SetResourceRegionId(v string) *ChangeResourceGroupRequest {
	s.ResourceRegionId = &v
	return s
}

type ChangeResourceGroupResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 314CB661-35A5-5F01-A623-3EC6F87FF52F
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

type CreateArtifactBuildRuleRequest struct {
	// The type of the artifact.
	//
	// 	- `ACCELERATED_IMAGE`: accelerated images.
	//
	// This parameter is required.
	//
	// example:
	//
	// ACCELERATED_IMAGE
	ArtifactType *string `json:"ArtifactType,omitempty" xml:"ArtifactType,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cri-cxreylqvcyje****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Additional parameters.
	//
	// example:
	//
	// {}
	Parameters map[string]interface{} `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// The ID of the effective range of the rule.
	//
	// 	- Set the value to the ID of the image repository.
	//
	// This parameter is required.
	//
	// example:
	//
	// crr-pmajihou6cg0****
	ScopeId *string `json:"ScopeId,omitempty" xml:"ScopeId,omitempty"`
	// The effective range of the rule. Valid values:
	//
	// 	- `REPOSITORY`
	//
	// This parameter is required.
	//
	// example:
	//
	// REPOSITORY
	ScopeType *string `json:"ScopeType,omitempty" xml:"ScopeType,omitempty"`
}

func (s CreateArtifactBuildRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateArtifactBuildRuleRequest) GoString() string {
	return s.String()
}

func (s *CreateArtifactBuildRuleRequest) SetArtifactType(v string) *CreateArtifactBuildRuleRequest {
	s.ArtifactType = &v
	return s
}

func (s *CreateArtifactBuildRuleRequest) SetInstanceId(v string) *CreateArtifactBuildRuleRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateArtifactBuildRuleRequest) SetParameters(v map[string]interface{}) *CreateArtifactBuildRuleRequest {
	s.Parameters = v
	return s
}

func (s *CreateArtifactBuildRuleRequest) SetScopeId(v string) *CreateArtifactBuildRuleRequest {
	s.ScopeId = &v
	return s
}

func (s *CreateArtifactBuildRuleRequest) SetScopeType(v string) *CreateArtifactBuildRuleRequest {
	s.ScopeType = &v
	return s
}

type CreateArtifactBuildRuleShrinkRequest struct {
	// The type of the artifact.
	//
	// 	- `ACCELERATED_IMAGE`: accelerated images.
	//
	// This parameter is required.
	//
	// example:
	//
	// ACCELERATED_IMAGE
	ArtifactType *string `json:"ArtifactType,omitempty" xml:"ArtifactType,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cri-cxreylqvcyje****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Additional parameters.
	//
	// example:
	//
	// {}
	ParametersShrink *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// The ID of the effective range of the rule.
	//
	// 	- Set the value to the ID of the image repository.
	//
	// This parameter is required.
	//
	// example:
	//
	// crr-pmajihou6cg0****
	ScopeId *string `json:"ScopeId,omitempty" xml:"ScopeId,omitempty"`
	// The effective range of the rule. Valid values:
	//
	// 	- `REPOSITORY`
	//
	// This parameter is required.
	//
	// example:
	//
	// REPOSITORY
	ScopeType *string `json:"ScopeType,omitempty" xml:"ScopeType,omitempty"`
}

func (s CreateArtifactBuildRuleShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateArtifactBuildRuleShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateArtifactBuildRuleShrinkRequest) SetArtifactType(v string) *CreateArtifactBuildRuleShrinkRequest {
	s.ArtifactType = &v
	return s
}

func (s *CreateArtifactBuildRuleShrinkRequest) SetInstanceId(v string) *CreateArtifactBuildRuleShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateArtifactBuildRuleShrinkRequest) SetParametersShrink(v string) *CreateArtifactBuildRuleShrinkRequest {
	s.ParametersShrink = &v
	return s
}

func (s *CreateArtifactBuildRuleShrinkRequest) SetScopeId(v string) *CreateArtifactBuildRuleShrinkRequest {
	s.ScopeId = &v
	return s
}

func (s *CreateArtifactBuildRuleShrinkRequest) SetScopeType(v string) *CreateArtifactBuildRuleShrinkRequest {
	s.ScopeType = &v
	return s
}

type CreateArtifactBuildRuleResponseBody struct {
	// The ID of the accelerated image building rule.
	//
	// example:
	//
	// crabr-7dfa5qye5****
	BuildRuleId *string `json:"BuildRuleId,omitempty" xml:"BuildRuleId,omitempty"`
	// The return value.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The request ID.
	//
	// example:
	//
	// C8E90AB5-0A96-5D12-9E59-11EE46360642
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateArtifactBuildRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateArtifactBuildRuleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateArtifactBuildRuleResponseBody) SetBuildRuleId(v string) *CreateArtifactBuildRuleResponseBody {
	s.BuildRuleId = &v
	return s
}

func (s *CreateArtifactBuildRuleResponseBody) SetCode(v string) *CreateArtifactBuildRuleResponseBody {
	s.Code = &v
	return s
}

func (s *CreateArtifactBuildRuleResponseBody) SetIsSuccess(v bool) *CreateArtifactBuildRuleResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *CreateArtifactBuildRuleResponseBody) SetRequestId(v string) *CreateArtifactBuildRuleResponseBody {
	s.RequestId = &v
	return s
}

type CreateArtifactBuildRuleResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateArtifactBuildRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateArtifactBuildRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateArtifactBuildRuleResponse) GoString() string {
	return s.String()
}

func (s *CreateArtifactBuildRuleResponse) SetHeaders(v map[string]*string) *CreateArtifactBuildRuleResponse {
	s.Headers = v
	return s
}

func (s *CreateArtifactBuildRuleResponse) SetStatusCode(v int32) *CreateArtifactBuildRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateArtifactBuildRuleResponse) SetBody(v *CreateArtifactBuildRuleResponseBody) *CreateArtifactBuildRuleResponse {
	s.Body = v
	return s
}

type CreateArtifactLifecycleRuleRequest struct {
	// Specify whether to automatically execute the lifecycle management rule.
	//
	// example:
	//
	// false
	Auto *bool `json:"Auto,omitempty" xml:"Auto,omitempty"`
	// Specify whether to enable lifecycle management for the artifact.
	//
	// example:
	//
	// true
	EnableDeleteTag *bool `json:"EnableDeleteTag,omitempty" xml:"EnableDeleteTag,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cri-gbwfk7qbgrxe****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the namespace.
	//
	// example:
	//
	// dev-backend
	NamespaceName *string `json:"NamespaceName,omitempty" xml:"NamespaceName,omitempty"`
	// The name of the image repository.
	//
	// example:
	//
	// test_1
	RepoName *string `json:"RepoName,omitempty" xml:"RepoName,omitempty"`
	// The number of images that you want to retain.
	//
	// example:
	//
	// 30
	RetentionTagCount *int64 `json:"RetentionTagCount,omitempty" xml:"RetentionTagCount,omitempty"`
	// The execution cycle of the lifecycle management rule.
	//
	// example:
	//
	// WEEK
	ScheduleTime *string `json:"ScheduleTime,omitempty" xml:"ScheduleTime,omitempty"`
	// The deletion scope.
	//
	// example:
	//
	// INSTANCE
	Scope *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
	// The regular expression that is used to indicate which image tags are retained.
	//
	// example:
	//
	// release-.*
	TagRegexp *string `json:"TagRegexp,omitempty" xml:"TagRegexp,omitempty"`
}

func (s CreateArtifactLifecycleRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateArtifactLifecycleRuleRequest) GoString() string {
	return s.String()
}

func (s *CreateArtifactLifecycleRuleRequest) SetAuto(v bool) *CreateArtifactLifecycleRuleRequest {
	s.Auto = &v
	return s
}

func (s *CreateArtifactLifecycleRuleRequest) SetEnableDeleteTag(v bool) *CreateArtifactLifecycleRuleRequest {
	s.EnableDeleteTag = &v
	return s
}

func (s *CreateArtifactLifecycleRuleRequest) SetInstanceId(v string) *CreateArtifactLifecycleRuleRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateArtifactLifecycleRuleRequest) SetNamespaceName(v string) *CreateArtifactLifecycleRuleRequest {
	s.NamespaceName = &v
	return s
}

func (s *CreateArtifactLifecycleRuleRequest) SetRepoName(v string) *CreateArtifactLifecycleRuleRequest {
	s.RepoName = &v
	return s
}

func (s *CreateArtifactLifecycleRuleRequest) SetRetentionTagCount(v int64) *CreateArtifactLifecycleRuleRequest {
	s.RetentionTagCount = &v
	return s
}

func (s *CreateArtifactLifecycleRuleRequest) SetScheduleTime(v string) *CreateArtifactLifecycleRuleRequest {
	s.ScheduleTime = &v
	return s
}

func (s *CreateArtifactLifecycleRuleRequest) SetScope(v string) *CreateArtifactLifecycleRuleRequest {
	s.Scope = &v
	return s
}

func (s *CreateArtifactLifecycleRuleRequest) SetTagRegexp(v string) *CreateArtifactLifecycleRuleRequest {
	s.TagRegexp = &v
	return s
}

type CreateArtifactLifecycleRuleResponseBody struct {
	// The return value.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// 	- `true`: The request is successful.
	//
	// 	- `false`: The request fails.
	//
	// example:
	//
	// True
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0AA66379-B880-5123-9F6A-96BB25D****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The rule ID.
	//
	// example:
	//
	// cralr-b6thg027zmk1****
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s CreateArtifactLifecycleRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateArtifactLifecycleRuleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateArtifactLifecycleRuleResponseBody) SetCode(v string) *CreateArtifactLifecycleRuleResponseBody {
	s.Code = &v
	return s
}

func (s *CreateArtifactLifecycleRuleResponseBody) SetIsSuccess(v bool) *CreateArtifactLifecycleRuleResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *CreateArtifactLifecycleRuleResponseBody) SetRequestId(v string) *CreateArtifactLifecycleRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateArtifactLifecycleRuleResponseBody) SetRuleId(v string) *CreateArtifactLifecycleRuleResponseBody {
	s.RuleId = &v
	return s
}

type CreateArtifactLifecycleRuleResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateArtifactLifecycleRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateArtifactLifecycleRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateArtifactLifecycleRuleResponse) GoString() string {
	return s.String()
}

func (s *CreateArtifactLifecycleRuleResponse) SetHeaders(v map[string]*string) *CreateArtifactLifecycleRuleResponse {
	s.Headers = v
	return s
}

func (s *CreateArtifactLifecycleRuleResponse) SetStatusCode(v int32) *CreateArtifactLifecycleRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateArtifactLifecycleRuleResponse) SetBody(v *CreateArtifactLifecycleRuleResponseBody) *CreateArtifactLifecycleRuleResponse {
	s.Body = v
	return s
}

type CreateArtifactSubscriptionRuleRequest struct {
	// Indicates whether an acceleration link is enabled for image subscription. The subscription acceleration feature is in public preview. The feature is optimized based on scheduling policies and network links to accelerate image subscription.
	//
	// example:
	//
	// true
	Accelerate *bool `json:"Accelerate,omitempty" xml:"Accelerate,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cri-c0o11woew0k****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the Container Registry namespace.
	//
	// This parameter is required.
	//
	// example:
	//
	// test-ns
	NamespaceName *string `json:"NamespaceName,omitempty" xml:"NamespaceName,omitempty"`
	// Indicates whether the original image is overwritten.
	//
	// example:
	//
	// true
	Override *bool `json:"Override,omitempty" xml:"Override,omitempty"`
	// The operating system and architecture. If the source repository contains a multi-arch image, only the specified operating system and architecture are subscribed to the destination repository of the Enterprise Edition instance. Subscribe to the destination repository of an Enterprise Edition instance
	//
	// This parameter is required.
	Platform []*string `json:"Platform,omitempty" xml:"Platform,omitempty" type:"Repeated"`
	// The name of the Container Registry repository.
	//
	// This parameter is required.
	//
	// example:
	//
	// test-repo
	RepoName *string `json:"RepoName,omitempty" xml:"RepoName,omitempty"`
	// The source namespace.
	//
	// example:
	//
	// library
	SourceNamespaceName *string `json:"SourceNamespaceName,omitempty" xml:"SourceNamespaceName,omitempty"`
	// The source of the artifact.
	//
	// Valid values:
	//
	// 	- DOCKER_HUB: Docker Hub
	//
	// 	- GCR: GCR
	//
	// 	- QUAY: Quay.io
	//
	// This parameter is required.
	//
	// example:
	//
	// DOCKER_HUB
	SourceProvider *string `json:"SourceProvider,omitempty" xml:"SourceProvider,omitempty"`
	// The source repository.
	//
	// This parameter is required.
	//
	// example:
	//
	// nginx
	SourceRepoName *string `json:"SourceRepoName,omitempty" xml:"SourceRepoName,omitempty"`
	// The number of subscribed images.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	TagCount *int64 `json:"TagCount,omitempty" xml:"TagCount,omitempty"`
	// The image tag in the subscription source repository. Regular expressions are supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// release-v.*
	TagRegexp *string `json:"TagRegexp,omitempty" xml:"TagRegexp,omitempty"`
}

func (s CreateArtifactSubscriptionRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateArtifactSubscriptionRuleRequest) GoString() string {
	return s.String()
}

func (s *CreateArtifactSubscriptionRuleRequest) SetAccelerate(v bool) *CreateArtifactSubscriptionRuleRequest {
	s.Accelerate = &v
	return s
}

func (s *CreateArtifactSubscriptionRuleRequest) SetInstanceId(v string) *CreateArtifactSubscriptionRuleRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateArtifactSubscriptionRuleRequest) SetNamespaceName(v string) *CreateArtifactSubscriptionRuleRequest {
	s.NamespaceName = &v
	return s
}

func (s *CreateArtifactSubscriptionRuleRequest) SetOverride(v bool) *CreateArtifactSubscriptionRuleRequest {
	s.Override = &v
	return s
}

func (s *CreateArtifactSubscriptionRuleRequest) SetPlatform(v []*string) *CreateArtifactSubscriptionRuleRequest {
	s.Platform = v
	return s
}

func (s *CreateArtifactSubscriptionRuleRequest) SetRepoName(v string) *CreateArtifactSubscriptionRuleRequest {
	s.RepoName = &v
	return s
}

func (s *CreateArtifactSubscriptionRuleRequest) SetSourceNamespaceName(v string) *CreateArtifactSubscriptionRuleRequest {
	s.SourceNamespaceName = &v
	return s
}

func (s *CreateArtifactSubscriptionRuleRequest) SetSourceProvider(v string) *CreateArtifactSubscriptionRuleRequest {
	s.SourceProvider = &v
	return s
}

func (s *CreateArtifactSubscriptionRuleRequest) SetSourceRepoName(v string) *CreateArtifactSubscriptionRuleRequest {
	s.SourceRepoName = &v
	return s
}

func (s *CreateArtifactSubscriptionRuleRequest) SetTagCount(v int64) *CreateArtifactSubscriptionRuleRequest {
	s.TagCount = &v
	return s
}

func (s *CreateArtifactSubscriptionRuleRequest) SetTagRegexp(v string) *CreateArtifactSubscriptionRuleRequest {
	s.TagRegexp = &v
	return s
}

type CreateArtifactSubscriptionRuleResponseBody struct {
	// The return value.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicate whether the API request is successful. Valid values:
	//
	// 	- `true`: The request is successful.
	//
	// 	- `false`: The request fails.
	//
	// example:
	//
	// true
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 02B27D80-FD32-5155-931A-93700779BB9E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the artifact subscription rule.
	//
	// example:
	//
	// crasr-lxdfele7dg4****
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s CreateArtifactSubscriptionRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateArtifactSubscriptionRuleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateArtifactSubscriptionRuleResponseBody) SetCode(v string) *CreateArtifactSubscriptionRuleResponseBody {
	s.Code = &v
	return s
}

func (s *CreateArtifactSubscriptionRuleResponseBody) SetIsSuccess(v bool) *CreateArtifactSubscriptionRuleResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *CreateArtifactSubscriptionRuleResponseBody) SetRequestId(v string) *CreateArtifactSubscriptionRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateArtifactSubscriptionRuleResponseBody) SetRuleId(v string) *CreateArtifactSubscriptionRuleResponseBody {
	s.RuleId = &v
	return s
}

type CreateArtifactSubscriptionRuleResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateArtifactSubscriptionRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateArtifactSubscriptionRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateArtifactSubscriptionRuleResponse) GoString() string {
	return s.String()
}

func (s *CreateArtifactSubscriptionRuleResponse) SetHeaders(v map[string]*string) *CreateArtifactSubscriptionRuleResponse {
	s.Headers = v
	return s
}

func (s *CreateArtifactSubscriptionRuleResponse) SetStatusCode(v int32) *CreateArtifactSubscriptionRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateArtifactSubscriptionRuleResponse) SetBody(v *CreateArtifactSubscriptionRuleResponseBody) *CreateArtifactSubscriptionRuleResponse {
	s.Body = v
	return s
}

type CreateArtifactSubscriptionTaskRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cri-4ec5xvj4j0l****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The rule ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// crasr-88s7vmelc3m****
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s CreateArtifactSubscriptionTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateArtifactSubscriptionTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateArtifactSubscriptionTaskRequest) SetInstanceId(v string) *CreateArtifactSubscriptionTaskRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateArtifactSubscriptionTaskRequest) SetRuleId(v string) *CreateArtifactSubscriptionTaskRequest {
	s.RuleId = &v
	return s
}

type CreateArtifactSubscriptionTaskResponseBody struct {
	// The return value.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the API request is successful. Valid values:
	//
	// 	- `true`: The request is successful.
	//
	// 	- `false`: The request fails.
	//
	// example:
	//
	// true
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 132D314B-BDD4-564C-89FE-3E2BAE115239
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The task ID.
	//
	// example:
	//
	// crast-40le4es9yh0p****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CreateArtifactSubscriptionTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateArtifactSubscriptionTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateArtifactSubscriptionTaskResponseBody) SetCode(v string) *CreateArtifactSubscriptionTaskResponseBody {
	s.Code = &v
	return s
}

func (s *CreateArtifactSubscriptionTaskResponseBody) SetIsSuccess(v bool) *CreateArtifactSubscriptionTaskResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *CreateArtifactSubscriptionTaskResponseBody) SetRequestId(v string) *CreateArtifactSubscriptionTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateArtifactSubscriptionTaskResponseBody) SetTaskId(v string) *CreateArtifactSubscriptionTaskResponseBody {
	s.TaskId = &v
	return s
}

type CreateArtifactSubscriptionTaskResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateArtifactSubscriptionTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateArtifactSubscriptionTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateArtifactSubscriptionTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateArtifactSubscriptionTaskResponse) SetHeaders(v map[string]*string) *CreateArtifactSubscriptionTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateArtifactSubscriptionTaskResponse) SetStatusCode(v int32) *CreateArtifactSubscriptionTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateArtifactSubscriptionTaskResponse) SetBody(v *CreateArtifactSubscriptionTaskResponseBody) *CreateArtifactSubscriptionTaskResponse {
	s.Body = v
	return s
}

type CreateBuildRecordByRecordRequest struct {
	// The ID of the image building record.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0A311FC5-B8C6-4332-80E4-539EB73****
	BuildRecordId *string `json:"BuildRecordId,omitempty" xml:"BuildRecordId,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cri-hpdfkc6utbaq****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the image repository.
	//
	// This parameter is required.
	//
	// example:
	//
	// crr-hnoq7j93or3k****
	RepoId *string `json:"RepoId,omitempty" xml:"RepoId,omitempty"`
}

func (s CreateBuildRecordByRecordRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateBuildRecordByRecordRequest) GoString() string {
	return s.String()
}

func (s *CreateBuildRecordByRecordRequest) SetBuildRecordId(v string) *CreateBuildRecordByRecordRequest {
	s.BuildRecordId = &v
	return s
}

func (s *CreateBuildRecordByRecordRequest) SetInstanceId(v string) *CreateBuildRecordByRecordRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateBuildRecordByRecordRequest) SetRepoId(v string) *CreateBuildRecordByRecordRequest {
	s.RepoId = &v
	return s
}

type CreateBuildRecordByRecordResponseBody struct {
	// The ID of the image building record.
	//
	// example:
	//
	// crbr-ly77w5i3t31f****
	BuildRecordId *string `json:"BuildRecordId,omitempty" xml:"BuildRecordId,omitempty"`
	// The HTTP status code. The status code 200 indicates that the request is successful.\\
	//
	// Other status codes indicate that the request failed.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// 	- `true`: The request is successful.
	//
	// 	- `false`: The request fails.
	//
	// example:
	//
	// true
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 4CE1F661-75DD-4EBD-A4AD-057B26834ABB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateBuildRecordByRecordResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateBuildRecordByRecordResponseBody) GoString() string {
	return s.String()
}

func (s *CreateBuildRecordByRecordResponseBody) SetBuildRecordId(v string) *CreateBuildRecordByRecordResponseBody {
	s.BuildRecordId = &v
	return s
}

func (s *CreateBuildRecordByRecordResponseBody) SetCode(v string) *CreateBuildRecordByRecordResponseBody {
	s.Code = &v
	return s
}

func (s *CreateBuildRecordByRecordResponseBody) SetIsSuccess(v bool) *CreateBuildRecordByRecordResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *CreateBuildRecordByRecordResponseBody) SetRequestId(v string) *CreateBuildRecordByRecordResponseBody {
	s.RequestId = &v
	return s
}

type CreateBuildRecordByRecordResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateBuildRecordByRecordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateBuildRecordByRecordResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateBuildRecordByRecordResponse) GoString() string {
	return s.String()
}

func (s *CreateBuildRecordByRecordResponse) SetHeaders(v map[string]*string) *CreateBuildRecordByRecordResponse {
	s.Headers = v
	return s
}

func (s *CreateBuildRecordByRecordResponse) SetStatusCode(v int32) *CreateBuildRecordByRecordResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateBuildRecordByRecordResponse) SetBody(v *CreateBuildRecordByRecordResponseBody) *CreateBuildRecordByRecordResponse {
	s.Body = v
	return s
}

type CreateBuildRecordByRuleRequest struct {
	// The ID of the image building rule.
	//
	// This parameter is required.
	//
	// example:
	//
	// crbr-1j95g4bu2s1i****
	BuildRuleId *string `json:"BuildRuleId,omitempty" xml:"BuildRuleId,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cri-asd6vujuhay0****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the image repository.
	//
	// This parameter is required.
	//
	// example:
	//
	// crr-8dz3aedjqlmk****
	RepoId *string `json:"RepoId,omitempty" xml:"RepoId,omitempty"`
}

func (s CreateBuildRecordByRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateBuildRecordByRuleRequest) GoString() string {
	return s.String()
}

func (s *CreateBuildRecordByRuleRequest) SetBuildRuleId(v string) *CreateBuildRecordByRuleRequest {
	s.BuildRuleId = &v
	return s
}

func (s *CreateBuildRecordByRuleRequest) SetInstanceId(v string) *CreateBuildRecordByRuleRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateBuildRecordByRuleRequest) SetRepoId(v string) *CreateBuildRecordByRuleRequest {
	s.RepoId = &v
	return s
}

type CreateBuildRecordByRuleResponseBody struct {
	// The ID of the image building record.
	//
	// example:
	//
	// 0A311FC5-B8C6-4332-80E4-539EB73****
	BuildRecordId *string `json:"BuildRecordId,omitempty" xml:"BuildRecordId,omitempty"`
	// The return value.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The request ID.
	//
	// example:
	//
	// B01B8857-A16E-40E9-A37E-764F15776FAA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateBuildRecordByRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateBuildRecordByRuleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateBuildRecordByRuleResponseBody) SetBuildRecordId(v string) *CreateBuildRecordByRuleResponseBody {
	s.BuildRecordId = &v
	return s
}

func (s *CreateBuildRecordByRuleResponseBody) SetCode(v string) *CreateBuildRecordByRuleResponseBody {
	s.Code = &v
	return s
}

func (s *CreateBuildRecordByRuleResponseBody) SetIsSuccess(v bool) *CreateBuildRecordByRuleResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *CreateBuildRecordByRuleResponseBody) SetRequestId(v string) *CreateBuildRecordByRuleResponseBody {
	s.RequestId = &v
	return s
}

type CreateBuildRecordByRuleResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateBuildRecordByRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateBuildRecordByRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateBuildRecordByRuleResponse) GoString() string {
	return s.String()
}

func (s *CreateBuildRecordByRuleResponse) SetHeaders(v map[string]*string) *CreateBuildRecordByRuleResponse {
	s.Headers = v
	return s
}

func (s *CreateBuildRecordByRuleResponse) SetStatusCode(v int32) *CreateBuildRecordByRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateBuildRecordByRuleResponse) SetBody(v *CreateBuildRecordByRuleResponseBody) *CreateBuildRecordByRuleResponse {
	s.Body = v
	return s
}

type CreateChainRequest struct {
	// The configuration of the delivery chain in the JSON format.
	//
	// example:
	//
	// chainconfig
	ChainConfig *string `json:"ChainConfig,omitempty" xml:"ChainConfig,omitempty"`
	// The description of the delivery chain.
	//
	// example:
	//
	// description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cri-4cdrlqmhn4gm****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the delivery chain.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The name of the repository.
	//
	// example:
	//
	// repo1
	RepoName *string `json:"RepoName,omitempty" xml:"RepoName,omitempty"`
	// The name of the namespace.
	//
	// example:
	//
	// ns1
	RepoNamespaceName *string `json:"RepoNamespaceName,omitempty" xml:"RepoNamespaceName,omitempty"`
	// Repositories in which the delivery chain does not take effect.
	ScopeExclude []*string `json:"ScopeExclude,omitempty" xml:"ScopeExclude,omitempty" type:"Repeated"`
}

func (s CreateChainRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateChainRequest) GoString() string {
	return s.String()
}

func (s *CreateChainRequest) SetChainConfig(v string) *CreateChainRequest {
	s.ChainConfig = &v
	return s
}

func (s *CreateChainRequest) SetDescription(v string) *CreateChainRequest {
	s.Description = &v
	return s
}

func (s *CreateChainRequest) SetInstanceId(v string) *CreateChainRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateChainRequest) SetName(v string) *CreateChainRequest {
	s.Name = &v
	return s
}

func (s *CreateChainRequest) SetRepoName(v string) *CreateChainRequest {
	s.RepoName = &v
	return s
}

func (s *CreateChainRequest) SetRepoNamespaceName(v string) *CreateChainRequest {
	s.RepoNamespaceName = &v
	return s
}

func (s *CreateChainRequest) SetScopeExclude(v []*string) *CreateChainRequest {
	s.ScopeExclude = v
	return s
}

type CreateChainResponseBody struct {
	// The ID of the delivery chain.
	//
	// example:
	//
	// chi-02ymhtwl3cq8****
	ChainId *string `json:"ChainId,omitempty" xml:"ChainId,omitempty"`
	// The return value.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 4BC03B36-E515-5806-99AC-268AE3C0****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateChainResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateChainResponseBody) GoString() string {
	return s.String()
}

func (s *CreateChainResponseBody) SetChainId(v string) *CreateChainResponseBody {
	s.ChainId = &v
	return s
}

func (s *CreateChainResponseBody) SetCode(v string) *CreateChainResponseBody {
	s.Code = &v
	return s
}

func (s *CreateChainResponseBody) SetIsSuccess(v bool) *CreateChainResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *CreateChainResponseBody) SetRequestId(v string) *CreateChainResponseBody {
	s.RequestId = &v
	return s
}

type CreateChainResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateChainResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateChainResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateChainResponse) GoString() string {
	return s.String()
}

func (s *CreateChainResponse) SetHeaders(v map[string]*string) *CreateChainResponse {
	s.Headers = v
	return s
}

func (s *CreateChainResponse) SetStatusCode(v int32) *CreateChainResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateChainResponse) SetBody(v *CreateChainResponseBody) *CreateChainResponse {
	s.Body = v
	return s
}

type CreateChartNamespaceRequest struct {
	// Specifies whether to automatically create repositories in the namespace. Valid values:
	//
	// \\-`  true `: automatically creates repositories in the namespace.
	//
	// \\-`  false `: does not automatically create repositories in the namespace.
	//
	// example:
	//
	// true
	AutoCreateRepo *bool `json:"AutoCreateRepo,omitempty" xml:"AutoCreateRepo,omitempty"`
	// The default repository type. Valid values:
	//
	// 	- `PUBLIC`: a public repository
	//
	// 	- `PRIVATE`: a private repository
	//
	// example:
	//
	// PUBLIC
	DefaultRepoType *string `json:"DefaultRepoType,omitempty" xml:"DefaultRepoType,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cri-xkx6vujuhay0****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the namespace.
	//
	// This parameter is required.
	//
	// example:
	//
	// namespace01
	NamespaceName *string `json:"NamespaceName,omitempty" xml:"NamespaceName,omitempty"`
}

func (s CreateChartNamespaceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateChartNamespaceRequest) GoString() string {
	return s.String()
}

func (s *CreateChartNamespaceRequest) SetAutoCreateRepo(v bool) *CreateChartNamespaceRequest {
	s.AutoCreateRepo = &v
	return s
}

func (s *CreateChartNamespaceRequest) SetDefaultRepoType(v string) *CreateChartNamespaceRequest {
	s.DefaultRepoType = &v
	return s
}

func (s *CreateChartNamespaceRequest) SetInstanceId(v string) *CreateChartNamespaceRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateChartNamespaceRequest) SetNamespaceName(v string) *CreateChartNamespaceRequest {
	s.NamespaceName = &v
	return s
}

type CreateChartNamespaceResponseBody struct {
	// The return value.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// 	- `true`: The request is successful.
	//
	// 	- `false`: The request fails.
	//
	// example:
	//
	// true
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 724402D0-75CD-4794-BC20-7D3720823AE0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateChartNamespaceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateChartNamespaceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateChartNamespaceResponseBody) SetCode(v string) *CreateChartNamespaceResponseBody {
	s.Code = &v
	return s
}

func (s *CreateChartNamespaceResponseBody) SetIsSuccess(v bool) *CreateChartNamespaceResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *CreateChartNamespaceResponseBody) SetRequestId(v string) *CreateChartNamespaceResponseBody {
	s.RequestId = &v
	return s
}

type CreateChartNamespaceResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateChartNamespaceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateChartNamespaceResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateChartNamespaceResponse) GoString() string {
	return s.String()
}

func (s *CreateChartNamespaceResponse) SetHeaders(v map[string]*string) *CreateChartNamespaceResponse {
	s.Headers = v
	return s
}

func (s *CreateChartNamespaceResponse) SetStatusCode(v int32) *CreateChartNamespaceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateChartNamespaceResponse) SetBody(v *CreateChartNamespaceResponseBody) *CreateChartNamespaceResponse {
	s.Body = v
	return s
}

type CreateChartRepositoryRequest struct {
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cri-xkx6vujuhay0****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the repository.
	//
	// This parameter is required.
	//
	// example:
	//
	// repo01
	RepoName *string `json:"RepoName,omitempty" xml:"RepoName,omitempty"`
	// The name of the namespace to which the repository belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// namespace01
	RepoNamespaceName *string `json:"RepoNamespaceName,omitempty" xml:"RepoNamespaceName,omitempty"`
	// The default repository type. Valid values:
	//
	// 	- `PUBLIC`: a public repository.
	//
	// 	- `PRIVATE`: a private repository.
	//
	// You can specify the RepoType or Summary parameter. The RepoType parameter is optional.
	//
	// example:
	//
	// PUBLIC
	RepoType *string `json:"RepoType,omitempty" xml:"RepoType,omitempty"`
	// The summary of the repository.
	//
	// example:
	//
	// summary
	Summary *string `json:"Summary,omitempty" xml:"Summary,omitempty"`
}

func (s CreateChartRepositoryRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateChartRepositoryRequest) GoString() string {
	return s.String()
}

func (s *CreateChartRepositoryRequest) SetInstanceId(v string) *CreateChartRepositoryRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateChartRepositoryRequest) SetRepoName(v string) *CreateChartRepositoryRequest {
	s.RepoName = &v
	return s
}

func (s *CreateChartRepositoryRequest) SetRepoNamespaceName(v string) *CreateChartRepositoryRequest {
	s.RepoNamespaceName = &v
	return s
}

func (s *CreateChartRepositoryRequest) SetRepoType(v string) *CreateChartRepositoryRequest {
	s.RepoType = &v
	return s
}

func (s *CreateChartRepositoryRequest) SetSummary(v string) *CreateChartRepositoryRequest {
	s.Summary = &v
	return s
}

type CreateChartRepositoryResponseBody struct {
	// The return value.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// 	- `true`: The request is successful.
	//
	// 	- `false`: The request fails.
	//
	// example:
	//
	// true
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The ID of the repository.
	//
	// example:
	//
	// crcr-2micxey5ewj4****
	RepoId *string `json:"RepoId,omitempty" xml:"RepoId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 60390244-A483-491A-B41D-F866C95380A1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateChartRepositoryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateChartRepositoryResponseBody) GoString() string {
	return s.String()
}

func (s *CreateChartRepositoryResponseBody) SetCode(v string) *CreateChartRepositoryResponseBody {
	s.Code = &v
	return s
}

func (s *CreateChartRepositoryResponseBody) SetIsSuccess(v bool) *CreateChartRepositoryResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *CreateChartRepositoryResponseBody) SetRepoId(v string) *CreateChartRepositoryResponseBody {
	s.RepoId = &v
	return s
}

func (s *CreateChartRepositoryResponseBody) SetRequestId(v string) *CreateChartRepositoryResponseBody {
	s.RequestId = &v
	return s
}

type CreateChartRepositoryResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateChartRepositoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateChartRepositoryResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateChartRepositoryResponse) GoString() string {
	return s.String()
}

func (s *CreateChartRepositoryResponse) SetHeaders(v map[string]*string) *CreateChartRepositoryResponse {
	s.Headers = v
	return s
}

func (s *CreateChartRepositoryResponse) SetStatusCode(v int32) *CreateChartRepositoryResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateChartRepositoryResponse) SetBody(v *CreateChartRepositoryResponseBody) *CreateChartRepositoryResponse {
	s.Body = v
	return s
}

type CreateInstanceEndpointAclPolicyRequest struct {
	// The description.
	//
	// example:
	//
	// test
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The type of the endpoint. Set the value to Internet.
	//
	// This parameter is required.
	//
	// example:
	//
	// internet
	EndpointType *string `json:"EndpointType,omitempty" xml:"EndpointType,omitempty"`
	// The CIDR block that is accessible.
	//
	// This parameter is required.
	//
	// example:
	//
	// 192.168.1.1/32
	Entry *string `json:"Entry,omitempty" xml:"Entry,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cri-xkx6vujuhay0****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the module that you want to access. Valid values:
	//
	// 	- `Registry`: the image repository.
	//
	// 	- `Chart`: a Helm chart.
	//
	// example:
	//
	// Registry
	ModuleName *string `json:"ModuleName,omitempty" xml:"ModuleName,omitempty"`
}

func (s CreateInstanceEndpointAclPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceEndpointAclPolicyRequest) GoString() string {
	return s.String()
}

func (s *CreateInstanceEndpointAclPolicyRequest) SetComment(v string) *CreateInstanceEndpointAclPolicyRequest {
	s.Comment = &v
	return s
}

func (s *CreateInstanceEndpointAclPolicyRequest) SetEndpointType(v string) *CreateInstanceEndpointAclPolicyRequest {
	s.EndpointType = &v
	return s
}

func (s *CreateInstanceEndpointAclPolicyRequest) SetEntry(v string) *CreateInstanceEndpointAclPolicyRequest {
	s.Entry = &v
	return s
}

func (s *CreateInstanceEndpointAclPolicyRequest) SetInstanceId(v string) *CreateInstanceEndpointAclPolicyRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateInstanceEndpointAclPolicyRequest) SetModuleName(v string) *CreateInstanceEndpointAclPolicyRequest {
	s.ModuleName = &v
	return s
}

type CreateInstanceEndpointAclPolicyResponseBody struct {
	// The return value.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// 	- `true`: The request is successful.
	//
	// 	- `false`: The request fails.
	//
	// example:
	//
	// true
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// D735C5EC-4206-4F48-A090-307BF56BEB99
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateInstanceEndpointAclPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceEndpointAclPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *CreateInstanceEndpointAclPolicyResponseBody) SetCode(v string) *CreateInstanceEndpointAclPolicyResponseBody {
	s.Code = &v
	return s
}

func (s *CreateInstanceEndpointAclPolicyResponseBody) SetIsSuccess(v bool) *CreateInstanceEndpointAclPolicyResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *CreateInstanceEndpointAclPolicyResponseBody) SetRequestId(v string) *CreateInstanceEndpointAclPolicyResponseBody {
	s.RequestId = &v
	return s
}

type CreateInstanceEndpointAclPolicyResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateInstanceEndpointAclPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateInstanceEndpointAclPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceEndpointAclPolicyResponse) GoString() string {
	return s.String()
}

func (s *CreateInstanceEndpointAclPolicyResponse) SetHeaders(v map[string]*string) *CreateInstanceEndpointAclPolicyResponse {
	s.Headers = v
	return s
}

func (s *CreateInstanceEndpointAclPolicyResponse) SetStatusCode(v int32) *CreateInstanceEndpointAclPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateInstanceEndpointAclPolicyResponse) SetBody(v *CreateInstanceEndpointAclPolicyResponseBody) *CreateInstanceEndpointAclPolicyResponse {
	s.Body = v
	return s
}

type CreateInstanceVpcEndpointLinkedVpcRequest struct {
	// Specifies whether to automatically create Alibaba Cloud DNS PrivateZone records. Valid values:
	//
	// >  If you enable automatic creation of PrivateZone records, a PrivateZone record is automatically created when you associate a VPC with the instance.
	//
	// 	- `true`
	//
	// 	- `false`
	//
	// example:
	//
	// false
	EnableCreateDNSRecordInPvzt *bool `json:"EnableCreateDNSRecordInPvzt,omitempty" xml:"EnableCreateDNSRecordInPvzt,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cri-xkx6vujuhay0****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the module that you want to access. Valid values:
	//
	// 	- `Registry`: image repositories.
	//
	// 	- `Chart`: Helm charts.
	//
	// example:
	//
	// Registry
	ModuleName *string `json:"ModuleName,omitempty" xml:"ModuleName,omitempty"`
	// The VPC ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpc-uf6pa68zxnnlc48dd****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The ID of the vSwitch that is associated with the specified VPC.
	//
	// This parameter is required.
	//
	// example:
	//
	// vsw-uf6u0kn8x2gbzxfn2****
	VswitchId *string `json:"VswitchId,omitempty" xml:"VswitchId,omitempty"`
}

func (s CreateInstanceVpcEndpointLinkedVpcRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceVpcEndpointLinkedVpcRequest) GoString() string {
	return s.String()
}

func (s *CreateInstanceVpcEndpointLinkedVpcRequest) SetEnableCreateDNSRecordInPvzt(v bool) *CreateInstanceVpcEndpointLinkedVpcRequest {
	s.EnableCreateDNSRecordInPvzt = &v
	return s
}

func (s *CreateInstanceVpcEndpointLinkedVpcRequest) SetInstanceId(v string) *CreateInstanceVpcEndpointLinkedVpcRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateInstanceVpcEndpointLinkedVpcRequest) SetModuleName(v string) *CreateInstanceVpcEndpointLinkedVpcRequest {
	s.ModuleName = &v
	return s
}

func (s *CreateInstanceVpcEndpointLinkedVpcRequest) SetVpcId(v string) *CreateInstanceVpcEndpointLinkedVpcRequest {
	s.VpcId = &v
	return s
}

func (s *CreateInstanceVpcEndpointLinkedVpcRequest) SetVswitchId(v string) *CreateInstanceVpcEndpointLinkedVpcRequest {
	s.VswitchId = &v
	return s
}

type CreateInstanceVpcEndpointLinkedVpcResponseBody struct {
	// The return value.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the request is successful.
	//
	// 	- `true`: The request is successful.
	//
	// 	- `false`: The request fails.
	//
	// example:
	//
	// true
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The request ID.
	//
	// example:
	//
	// D4978DCC-ECBD-40B0-A714-EE6959B22C77
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateInstanceVpcEndpointLinkedVpcResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceVpcEndpointLinkedVpcResponseBody) GoString() string {
	return s.String()
}

func (s *CreateInstanceVpcEndpointLinkedVpcResponseBody) SetCode(v string) *CreateInstanceVpcEndpointLinkedVpcResponseBody {
	s.Code = &v
	return s
}

func (s *CreateInstanceVpcEndpointLinkedVpcResponseBody) SetIsSuccess(v bool) *CreateInstanceVpcEndpointLinkedVpcResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *CreateInstanceVpcEndpointLinkedVpcResponseBody) SetRequestId(v string) *CreateInstanceVpcEndpointLinkedVpcResponseBody {
	s.RequestId = &v
	return s
}

type CreateInstanceVpcEndpointLinkedVpcResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateInstanceVpcEndpointLinkedVpcResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateInstanceVpcEndpointLinkedVpcResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceVpcEndpointLinkedVpcResponse) GoString() string {
	return s.String()
}

func (s *CreateInstanceVpcEndpointLinkedVpcResponse) SetHeaders(v map[string]*string) *CreateInstanceVpcEndpointLinkedVpcResponse {
	s.Headers = v
	return s
}

func (s *CreateInstanceVpcEndpointLinkedVpcResponse) SetStatusCode(v int32) *CreateInstanceVpcEndpointLinkedVpcResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateInstanceVpcEndpointLinkedVpcResponse) SetBody(v *CreateInstanceVpcEndpointLinkedVpcResponseBody) *CreateInstanceVpcEndpointLinkedVpcResponse {
	s.Body = v
	return s
}

type CreateNamespaceRequest struct {
	// Specifies whether to automatically create an image repository in the namespace.
	//
	// example:
	//
	// true
	AutoCreateRepo           *bool              `json:"AutoCreateRepo,omitempty" xml:"AutoCreateRepo,omitempty"`
	DefaultRepoConfiguration *RepoConfiguration `json:"DefaultRepoConfiguration,omitempty" xml:"DefaultRepoConfiguration,omitempty"`
	// Deprecated
	//
	// The default type of the repositories that are automatically created in the namespace. Valid values:
	//
	// 	- `PUBLIC`: public repositories
	//
	// 	- `PRIVATE`: private repositories.
	//
	// example:
	//
	// PUBLIC
	DefaultRepoType *string `json:"DefaultRepoType,omitempty" xml:"DefaultRepoType,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cri-xkx6vujuhay0****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the namespace. The name must be 2 to 120 characters in length, and can contain lowercase letters, digits, and the following delimiters: underscores (_), hyphens (-), and periods (.). The name cannot start or end with a delimiter.
	//
	// This parameter is required.
	//
	// example:
	//
	// namespace1
	NamespaceName *string `json:"NamespaceName,omitempty" xml:"NamespaceName,omitempty"`
}

func (s CreateNamespaceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateNamespaceRequest) GoString() string {
	return s.String()
}

func (s *CreateNamespaceRequest) SetAutoCreateRepo(v bool) *CreateNamespaceRequest {
	s.AutoCreateRepo = &v
	return s
}

func (s *CreateNamespaceRequest) SetDefaultRepoConfiguration(v *RepoConfiguration) *CreateNamespaceRequest {
	s.DefaultRepoConfiguration = v
	return s
}

func (s *CreateNamespaceRequest) SetDefaultRepoType(v string) *CreateNamespaceRequest {
	s.DefaultRepoType = &v
	return s
}

func (s *CreateNamespaceRequest) SetInstanceId(v string) *CreateNamespaceRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateNamespaceRequest) SetNamespaceName(v string) *CreateNamespaceRequest {
	s.NamespaceName = &v
	return s
}

type CreateNamespaceShrinkRequest struct {
	// Specifies whether to automatically create an image repository in the namespace.
	//
	// example:
	//
	// true
	AutoCreateRepo                 *bool   `json:"AutoCreateRepo,omitempty" xml:"AutoCreateRepo,omitempty"`
	DefaultRepoConfigurationShrink *string `json:"DefaultRepoConfiguration,omitempty" xml:"DefaultRepoConfiguration,omitempty"`
	// Deprecated
	//
	// The default type of the repositories that are automatically created in the namespace. Valid values:
	//
	// 	- `PUBLIC`: public repositories
	//
	// 	- `PRIVATE`: private repositories.
	//
	// example:
	//
	// PUBLIC
	DefaultRepoType *string `json:"DefaultRepoType,omitempty" xml:"DefaultRepoType,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cri-xkx6vujuhay0****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the namespace. The name must be 2 to 120 characters in length, and can contain lowercase letters, digits, and the following delimiters: underscores (_), hyphens (-), and periods (.). The name cannot start or end with a delimiter.
	//
	// This parameter is required.
	//
	// example:
	//
	// namespace1
	NamespaceName *string `json:"NamespaceName,omitempty" xml:"NamespaceName,omitempty"`
}

func (s CreateNamespaceShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateNamespaceShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateNamespaceShrinkRequest) SetAutoCreateRepo(v bool) *CreateNamespaceShrinkRequest {
	s.AutoCreateRepo = &v
	return s
}

func (s *CreateNamespaceShrinkRequest) SetDefaultRepoConfigurationShrink(v string) *CreateNamespaceShrinkRequest {
	s.DefaultRepoConfigurationShrink = &v
	return s
}

func (s *CreateNamespaceShrinkRequest) SetDefaultRepoType(v string) *CreateNamespaceShrinkRequest {
	s.DefaultRepoType = &v
	return s
}

func (s *CreateNamespaceShrinkRequest) SetInstanceId(v string) *CreateNamespaceShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateNamespaceShrinkRequest) SetNamespaceName(v string) *CreateNamespaceShrinkRequest {
	s.NamespaceName = &v
	return s
}

type CreateNamespaceResponseBody struct {
	// The return value.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// BC648259-91A7-4502-BED3-EDF64361FA83
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateNamespaceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateNamespaceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateNamespaceResponseBody) SetCode(v string) *CreateNamespaceResponseBody {
	s.Code = &v
	return s
}

func (s *CreateNamespaceResponseBody) SetIsSuccess(v bool) *CreateNamespaceResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *CreateNamespaceResponseBody) SetRequestId(v string) *CreateNamespaceResponseBody {
	s.RequestId = &v
	return s
}

type CreateNamespaceResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateNamespaceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateNamespaceResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateNamespaceResponse) GoString() string {
	return s.String()
}

func (s *CreateNamespaceResponse) SetHeaders(v map[string]*string) *CreateNamespaceResponse {
	s.Headers = v
	return s
}

func (s *CreateNamespaceResponse) SetStatusCode(v int32) *CreateNamespaceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateNamespaceResponse) SetBody(v *CreateNamespaceResponseBody) *CreateNamespaceResponse {
	s.Body = v
	return s
}

type CreateRepoBuildRuleRequest struct {
	// Building arguments.
	BuildArgs []*string `json:"BuildArgs,omitempty" xml:"BuildArgs,omitempty" type:"Repeated"`
	// The path of the Dockerfile.
	//
	// example:
	//
	// /
	DockerfileLocation *string `json:"DockerfileLocation,omitempty" xml:"DockerfileLocation,omitempty"`
	// The name of the Dockerfile.
	//
	// example:
	//
	// Dockerfile
	DockerfileName *string `json:"DockerfileName,omitempty" xml:"DockerfileName,omitempty"`
	// The tag of the image.
	//
	// This parameter is required.
	//
	// example:
	//
	// v0.9.5
	ImageTag *string `json:"ImageTag,omitempty" xml:"ImageTag,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cri-xkx6vujuhay0****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Architecture for image building. Valid values:
	//
	// 	- `linux/amd64`
	//
	// 	- `linux/arm64`
	//
	// 	- `linux/386`
	//
	// 	- `linux/arm/v7`
	//
	// 	- `inux/arm/v6`
	//
	// Default value: `linux/amd64`
	Platforms []*string `json:"Platforms,omitempty" xml:"Platforms,omitempty" type:"Repeated"`
	// The name of the push that triggers the building rule.
	//
	// This parameter is required.
	//
	// example:
	//
	// master
	PushName *string `json:"PushName,omitempty" xml:"PushName,omitempty"`
	// The type of the push that triggers the building rule. Valid values:
	//
	// 	- `GIT_TAG`: tag push
	//
	// 	- `GIT_BRANCH`: branch push
	//
	// This parameter is required.
	//
	// example:
	//
	// GIT_BRANCH
	PushType *string `json:"PushType,omitempty" xml:"PushType,omitempty"`
	// The ID of the image repository.
	//
	// This parameter is required.
	//
	// example:
	//
	// crr-8dz3aedjqlmk****
	RepoId *string `json:"RepoId,omitempty" xml:"RepoId,omitempty"`
}

func (s CreateRepoBuildRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateRepoBuildRuleRequest) GoString() string {
	return s.String()
}

func (s *CreateRepoBuildRuleRequest) SetBuildArgs(v []*string) *CreateRepoBuildRuleRequest {
	s.BuildArgs = v
	return s
}

func (s *CreateRepoBuildRuleRequest) SetDockerfileLocation(v string) *CreateRepoBuildRuleRequest {
	s.DockerfileLocation = &v
	return s
}

func (s *CreateRepoBuildRuleRequest) SetDockerfileName(v string) *CreateRepoBuildRuleRequest {
	s.DockerfileName = &v
	return s
}

func (s *CreateRepoBuildRuleRequest) SetImageTag(v string) *CreateRepoBuildRuleRequest {
	s.ImageTag = &v
	return s
}

func (s *CreateRepoBuildRuleRequest) SetInstanceId(v string) *CreateRepoBuildRuleRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateRepoBuildRuleRequest) SetPlatforms(v []*string) *CreateRepoBuildRuleRequest {
	s.Platforms = v
	return s
}

func (s *CreateRepoBuildRuleRequest) SetPushName(v string) *CreateRepoBuildRuleRequest {
	s.PushName = &v
	return s
}

func (s *CreateRepoBuildRuleRequest) SetPushType(v string) *CreateRepoBuildRuleRequest {
	s.PushType = &v
	return s
}

func (s *CreateRepoBuildRuleRequest) SetRepoId(v string) *CreateRepoBuildRuleRequest {
	s.RepoId = &v
	return s
}

type CreateRepoBuildRuleResponseBody struct {
	// The ID of the building rule.
	//
	// example:
	//
	// crbr-ly77w5i3t31f****
	BuildRuleId *string `json:"BuildRuleId,omitempty" xml:"BuildRuleId,omitempty"`
	// The return value.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// 	- `true`: The request is successful.
	//
	// 	- `false`: The request fails.
	//
	// example:
	//
	// true
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 4CE1F661-75DD-4EBD-A4AD-057B26834ABB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateRepoBuildRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateRepoBuildRuleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRepoBuildRuleResponseBody) SetBuildRuleId(v string) *CreateRepoBuildRuleResponseBody {
	s.BuildRuleId = &v
	return s
}

func (s *CreateRepoBuildRuleResponseBody) SetCode(v string) *CreateRepoBuildRuleResponseBody {
	s.Code = &v
	return s
}

func (s *CreateRepoBuildRuleResponseBody) SetIsSuccess(v bool) *CreateRepoBuildRuleResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *CreateRepoBuildRuleResponseBody) SetRequestId(v string) *CreateRepoBuildRuleResponseBody {
	s.RequestId = &v
	return s
}

type CreateRepoBuildRuleResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateRepoBuildRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateRepoBuildRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateRepoBuildRuleResponse) GoString() string {
	return s.String()
}

func (s *CreateRepoBuildRuleResponse) SetHeaders(v map[string]*string) *CreateRepoBuildRuleResponse {
	s.Headers = v
	return s
}

func (s *CreateRepoBuildRuleResponse) SetStatusCode(v int32) *CreateRepoBuildRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateRepoBuildRuleResponse) SetBody(v *CreateRepoBuildRuleResponseBody) *CreateRepoBuildRuleResponse {
	s.Body = v
	return s
}

type CreateRepoSourceCodeRepoRequest struct {
	// Specifies whether to trigger image building when source code is committed. Valid values:
	//
	// 	- `true`: triggers image building when source code is committed.
	//
	// 	- `false`: does not trigger image building when source code is committed.
	//
	// example:
	//
	// true
	AutoBuild *bool `json:"AutoBuild,omitempty" xml:"AutoBuild,omitempty"`
	// The name of the source code repository.
	//
	// This parameter is required.
	//
	// example:
	//
	// repo
	CodeRepoName *string `json:"CodeRepoName,omitempty" xml:"CodeRepoName,omitempty"`
	// The namespace to which the source code repository belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// namespace
	CodeRepoNamespaceName *string `json:"CodeRepoNamespaceName,omitempty" xml:"CodeRepoNamespaceName,omitempty"`
	// The type of the source code hosting platform. Valid values: `GITHUB`, `GITLAB`, `GITEE`, `CODE`, and `CODEUP`.
	//
	// This parameter is required.
	//
	// example:
	//
	// GITHUB
	CodeRepoType *string `json:"CodeRepoType,omitempty" xml:"CodeRepoType,omitempty"`
	// Specifies whether to disable building caches. Valid values:
	//
	// 	- `true`: disables building caches.
	//
	// 	- `false`: enables building caches.
	//
	// example:
	//
	// false
	DisableCacheBuild *bool `json:"DisableCacheBuild,omitempty" xml:"DisableCacheBuild,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cri-shac42yvqzvq****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Specifies whether to enable Build With Servers Deployed Outside Chinese Mainland. Valid values:
	//
	// 	- `true`: enables Build With Servers Deployed Outside Chinese Mainland.
	//
	// 	- `false`: does not enable Build With Servers Deployed Outside Chinese Mainland.
	//
	// example:
	//
	// false
	OverseaBuild *bool `json:"OverseaBuild,omitempty" xml:"OverseaBuild,omitempty"`
	// The ID of the image repository.
	//
	// This parameter is required.
	//
	// example:
	//
	// crr-gzsrlevmvoaq****
	RepoId *string `json:"RepoId,omitempty" xml:"RepoId,omitempty"`
}

func (s CreateRepoSourceCodeRepoRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateRepoSourceCodeRepoRequest) GoString() string {
	return s.String()
}

func (s *CreateRepoSourceCodeRepoRequest) SetAutoBuild(v bool) *CreateRepoSourceCodeRepoRequest {
	s.AutoBuild = &v
	return s
}

func (s *CreateRepoSourceCodeRepoRequest) SetCodeRepoName(v string) *CreateRepoSourceCodeRepoRequest {
	s.CodeRepoName = &v
	return s
}

func (s *CreateRepoSourceCodeRepoRequest) SetCodeRepoNamespaceName(v string) *CreateRepoSourceCodeRepoRequest {
	s.CodeRepoNamespaceName = &v
	return s
}

func (s *CreateRepoSourceCodeRepoRequest) SetCodeRepoType(v string) *CreateRepoSourceCodeRepoRequest {
	s.CodeRepoType = &v
	return s
}

func (s *CreateRepoSourceCodeRepoRequest) SetDisableCacheBuild(v bool) *CreateRepoSourceCodeRepoRequest {
	s.DisableCacheBuild = &v
	return s
}

func (s *CreateRepoSourceCodeRepoRequest) SetInstanceId(v string) *CreateRepoSourceCodeRepoRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateRepoSourceCodeRepoRequest) SetOverseaBuild(v bool) *CreateRepoSourceCodeRepoRequest {
	s.OverseaBuild = &v
	return s
}

func (s *CreateRepoSourceCodeRepoRequest) SetRepoId(v string) *CreateRepoSourceCodeRepoRequest {
	s.RepoId = &v
	return s
}

type CreateRepoSourceCodeRepoResponseBody struct {
	// The return value.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 4CE1F661-75DD-4EBD-A4AD-057B26834ABB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateRepoSourceCodeRepoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateRepoSourceCodeRepoResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRepoSourceCodeRepoResponseBody) SetCode(v string) *CreateRepoSourceCodeRepoResponseBody {
	s.Code = &v
	return s
}

func (s *CreateRepoSourceCodeRepoResponseBody) SetIsSuccess(v bool) *CreateRepoSourceCodeRepoResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *CreateRepoSourceCodeRepoResponseBody) SetRequestId(v string) *CreateRepoSourceCodeRepoResponseBody {
	s.RequestId = &v
	return s
}

type CreateRepoSourceCodeRepoResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateRepoSourceCodeRepoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateRepoSourceCodeRepoResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateRepoSourceCodeRepoResponse) GoString() string {
	return s.String()
}

func (s *CreateRepoSourceCodeRepoResponse) SetHeaders(v map[string]*string) *CreateRepoSourceCodeRepoResponse {
	s.Headers = v
	return s
}

func (s *CreateRepoSourceCodeRepoResponse) SetStatusCode(v int32) *CreateRepoSourceCodeRepoResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateRepoSourceCodeRepoResponse) SetBody(v *CreateRepoSourceCodeRepoResponseBody) *CreateRepoSourceCodeRepoResponse {
	s.Body = v
	return s
}

type CreateRepoSyncRuleRequest struct {
	// The source instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cri-hpdfkc6utbaq****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The namespace name of the source instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// ns1
	NamespaceName *string `json:"NamespaceName,omitempty" xml:"NamespaceName,omitempty"`
	// The name of the image repository in the source instance.
	//
	// example:
	//
	// repo1
	RepoName *string `json:"RepoName,omitempty" xml:"RepoName,omitempty"`
	// The regular expression that is used to filter repositories.
	//
	// >  This parameter is valid only when SyncScope is set to `NAMESPACE`.
	//
	// example:
	//
	// .*
	RepoNameFilter *string `json:"RepoNameFilter,omitempty" xml:"RepoNameFilter,omitempty"`
	// The name of the image synchronization rule.
	//
	// This parameter is required.
	//
	// example:
	//
	// rule
	SyncRuleName *string `json:"SyncRuleName,omitempty" xml:"SyncRuleName,omitempty"`
	// The synchronization scope. Valid values:
	//
	// 	- `REPO`: synchronizes the image tags in an image repository that meet the synchronization rule.
	//
	// 	- `NAMESPACE`: synchronizes the image tags in a namespace that meet the synchronization rule.
	//
	// This parameter is required.
	//
	// example:
	//
	// REPO
	SyncScope *string `json:"SyncScope,omitempty" xml:"SyncScope,omitempty"`
	// The mode of triggering the synchronization rule. Valid values:
	//
	// 	- `INITIATIVE`: manually triggers the synchronization rule.
	//
	// 	- `PASSIVE`: automatically triggers the synchronization rule.
	//
	// example:
	//
	// PASSIVE
	SyncTrigger *string `json:"SyncTrigger,omitempty" xml:"SyncTrigger,omitempty"`
	// The regular expression that is used to filter image tags.
	//
	// This parameter is required.
	//
	// example:
	//
	// .*
	TagFilter *string `json:"TagFilter,omitempty" xml:"TagFilter,omitempty"`
	// The destination instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cri-ibxs3piklys3****
	TargetInstanceId *string `json:"TargetInstanceId,omitempty" xml:"TargetInstanceId,omitempty"`
	// The namespace name of the destination instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// ns1
	TargetNamespaceName *string `json:"TargetNamespaceName,omitempty" xml:"TargetNamespaceName,omitempty"`
	// The region ID of the destination instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	TargetRegionId *string `json:"TargetRegionId,omitempty" xml:"TargetRegionId,omitempty"`
	// The name of the image repository in the destination instance.
	//
	// example:
	//
	// repo1
	TargetRepoName *string `json:"TargetRepoName,omitempty" xml:"TargetRepoName,omitempty"`
	// The user ID (UID) of the account to which the destination instance belongs.
	//
	// >  If you synchronize images across accounts, you must use the UID.
	//
	// example:
	//
	// 12645940***
	TargetUserId *string `json:"TargetUserId,omitempty" xml:"TargetUserId,omitempty"`
}

func (s CreateRepoSyncRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateRepoSyncRuleRequest) GoString() string {
	return s.String()
}

func (s *CreateRepoSyncRuleRequest) SetInstanceId(v string) *CreateRepoSyncRuleRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateRepoSyncRuleRequest) SetNamespaceName(v string) *CreateRepoSyncRuleRequest {
	s.NamespaceName = &v
	return s
}

func (s *CreateRepoSyncRuleRequest) SetRepoName(v string) *CreateRepoSyncRuleRequest {
	s.RepoName = &v
	return s
}

func (s *CreateRepoSyncRuleRequest) SetRepoNameFilter(v string) *CreateRepoSyncRuleRequest {
	s.RepoNameFilter = &v
	return s
}

func (s *CreateRepoSyncRuleRequest) SetSyncRuleName(v string) *CreateRepoSyncRuleRequest {
	s.SyncRuleName = &v
	return s
}

func (s *CreateRepoSyncRuleRequest) SetSyncScope(v string) *CreateRepoSyncRuleRequest {
	s.SyncScope = &v
	return s
}

func (s *CreateRepoSyncRuleRequest) SetSyncTrigger(v string) *CreateRepoSyncRuleRequest {
	s.SyncTrigger = &v
	return s
}

func (s *CreateRepoSyncRuleRequest) SetTagFilter(v string) *CreateRepoSyncRuleRequest {
	s.TagFilter = &v
	return s
}

func (s *CreateRepoSyncRuleRequest) SetTargetInstanceId(v string) *CreateRepoSyncRuleRequest {
	s.TargetInstanceId = &v
	return s
}

func (s *CreateRepoSyncRuleRequest) SetTargetNamespaceName(v string) *CreateRepoSyncRuleRequest {
	s.TargetNamespaceName = &v
	return s
}

func (s *CreateRepoSyncRuleRequest) SetTargetRegionId(v string) *CreateRepoSyncRuleRequest {
	s.TargetRegionId = &v
	return s
}

func (s *CreateRepoSyncRuleRequest) SetTargetRepoName(v string) *CreateRepoSyncRuleRequest {
	s.TargetRepoName = &v
	return s
}

func (s *CreateRepoSyncRuleRequest) SetTargetUserId(v string) *CreateRepoSyncRuleRequest {
	s.TargetUserId = &v
	return s
}

type CreateRepoSyncRuleResponseBody struct {
	// The HTTP status code.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 8F8A0BA6-7F06-4BAE-B147-10BD6A25****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the synchronization rule.
	//
	// example:
	//
	// crsr-gk5p2ns1kzns****
	SyncRuleId *string `json:"SyncRuleId,omitempty" xml:"SyncRuleId,omitempty"`
}

func (s CreateRepoSyncRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateRepoSyncRuleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRepoSyncRuleResponseBody) SetCode(v string) *CreateRepoSyncRuleResponseBody {
	s.Code = &v
	return s
}

func (s *CreateRepoSyncRuleResponseBody) SetIsSuccess(v bool) *CreateRepoSyncRuleResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *CreateRepoSyncRuleResponseBody) SetRequestId(v string) *CreateRepoSyncRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateRepoSyncRuleResponseBody) SetSyncRuleId(v string) *CreateRepoSyncRuleResponseBody {
	s.SyncRuleId = &v
	return s
}

type CreateRepoSyncRuleResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateRepoSyncRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateRepoSyncRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateRepoSyncRuleResponse) GoString() string {
	return s.String()
}

func (s *CreateRepoSyncRuleResponse) SetHeaders(v map[string]*string) *CreateRepoSyncRuleResponse {
	s.Headers = v
	return s
}

func (s *CreateRepoSyncRuleResponse) SetStatusCode(v int32) *CreateRepoSyncRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateRepoSyncRuleResponse) SetBody(v *CreateRepoSyncRuleResponseBody) *CreateRepoSyncRuleResponse {
	s.Body = v
	return s
}

type CreateRepoSyncTaskRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// cri-hpdfkc6utbaq****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// true
	Override *bool `json:"Override,omitempty" xml:"Override,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// crr-iql7jalx4g0****
	RepoId *string `json:"RepoId,omitempty" xml:"RepoId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// tag1
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cri-ibxs3piklys3****
	TargetInstanceId *string `json:"TargetInstanceId,omitempty" xml:"TargetInstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ns1
	TargetNamespace *string `json:"TargetNamespace,omitempty" xml:"TargetNamespace,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	TargetRegionId *string `json:"TargetRegionId,omitempty" xml:"TargetRegionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// repo1
	TargetRepoName *string `json:"TargetRepoName,omitempty" xml:"TargetRepoName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// tag1
	TargetTag *string `json:"TargetTag,omitempty" xml:"TargetTag,omitempty"`
	// example:
	//
	// 12345***
	TargetUserId *string `json:"TargetUserId,omitempty" xml:"TargetUserId,omitempty"`
}

func (s CreateRepoSyncTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateRepoSyncTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateRepoSyncTaskRequest) SetInstanceId(v string) *CreateRepoSyncTaskRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateRepoSyncTaskRequest) SetOverride(v bool) *CreateRepoSyncTaskRequest {
	s.Override = &v
	return s
}

func (s *CreateRepoSyncTaskRequest) SetRepoId(v string) *CreateRepoSyncTaskRequest {
	s.RepoId = &v
	return s
}

func (s *CreateRepoSyncTaskRequest) SetTag(v string) *CreateRepoSyncTaskRequest {
	s.Tag = &v
	return s
}

func (s *CreateRepoSyncTaskRequest) SetTargetInstanceId(v string) *CreateRepoSyncTaskRequest {
	s.TargetInstanceId = &v
	return s
}

func (s *CreateRepoSyncTaskRequest) SetTargetNamespace(v string) *CreateRepoSyncTaskRequest {
	s.TargetNamespace = &v
	return s
}

func (s *CreateRepoSyncTaskRequest) SetTargetRegionId(v string) *CreateRepoSyncTaskRequest {
	s.TargetRegionId = &v
	return s
}

func (s *CreateRepoSyncTaskRequest) SetTargetRepoName(v string) *CreateRepoSyncTaskRequest {
	s.TargetRepoName = &v
	return s
}

func (s *CreateRepoSyncTaskRequest) SetTargetTag(v string) *CreateRepoSyncTaskRequest {
	s.TargetTag = &v
	return s
}

func (s *CreateRepoSyncTaskRequest) SetTargetUserId(v string) *CreateRepoSyncTaskRequest {
	s.TargetUserId = &v
	return s
}

type CreateRepoSyncTaskResponseBody struct {
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// true
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// example:
	//
	// 8F8A0BA6-7F06-4BAE-B147-10BD6A25****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// rst-gbch330f0c****
	SyncTaskId *string `json:"SyncTaskId,omitempty" xml:"SyncTaskId,omitempty"`
}

func (s CreateRepoSyncTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateRepoSyncTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRepoSyncTaskResponseBody) SetCode(v string) *CreateRepoSyncTaskResponseBody {
	s.Code = &v
	return s
}

func (s *CreateRepoSyncTaskResponseBody) SetIsSuccess(v bool) *CreateRepoSyncTaskResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *CreateRepoSyncTaskResponseBody) SetRequestId(v string) *CreateRepoSyncTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateRepoSyncTaskResponseBody) SetSyncTaskId(v string) *CreateRepoSyncTaskResponseBody {
	s.SyncTaskId = &v
	return s
}

type CreateRepoSyncTaskResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateRepoSyncTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateRepoSyncTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateRepoSyncTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateRepoSyncTaskResponse) SetHeaders(v map[string]*string) *CreateRepoSyncTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateRepoSyncTaskResponse) SetStatusCode(v int32) *CreateRepoSyncTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateRepoSyncTaskResponse) SetBody(v *CreateRepoSyncTaskResponseBody) *CreateRepoSyncTaskResponse {
	s.Body = v
	return s
}

type CreateRepoSyncTaskByRuleRequest struct {
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cri-hpdfkc6utbaq****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the image repository.
	//
	// This parameter is required.
	//
	// example:
	//
	// crr-hnoq7j93or3k****
	RepoId *string `json:"RepoId,omitempty" xml:"RepoId,omitempty"`
	// The ID of the synchronization rule.
	//
	// This parameter is required.
	//
	// example:
	//
	// crsr-o8n4dijbumgq****
	SyncRuleId *string `json:"SyncRuleId,omitempty" xml:"SyncRuleId,omitempty"`
	// The version of the image to be synchronized.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1.24
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
}

func (s CreateRepoSyncTaskByRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateRepoSyncTaskByRuleRequest) GoString() string {
	return s.String()
}

func (s *CreateRepoSyncTaskByRuleRequest) SetInstanceId(v string) *CreateRepoSyncTaskByRuleRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateRepoSyncTaskByRuleRequest) SetRepoId(v string) *CreateRepoSyncTaskByRuleRequest {
	s.RepoId = &v
	return s
}

func (s *CreateRepoSyncTaskByRuleRequest) SetSyncRuleId(v string) *CreateRepoSyncTaskByRuleRequest {
	s.SyncRuleId = &v
	return s
}

func (s *CreateRepoSyncTaskByRuleRequest) SetTag(v string) *CreateRepoSyncTaskByRuleRequest {
	s.Tag = &v
	return s
}

type CreateRepoSyncTaskByRuleResponseBody struct {
	// The return value.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// 	- `true`: The request is successful.
	//
	// 	- `false`: The request fails.
	//
	// example:
	//
	// true
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 17A4C658-AE8F-4A08-821F-EDCB5FC74EE8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the synchronization task.
	//
	// example:
	//
	// rst-biu4u4pm4it5****
	SyncTaskId *string `json:"SyncTaskId,omitempty" xml:"SyncTaskId,omitempty"`
}

func (s CreateRepoSyncTaskByRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateRepoSyncTaskByRuleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRepoSyncTaskByRuleResponseBody) SetCode(v string) *CreateRepoSyncTaskByRuleResponseBody {
	s.Code = &v
	return s
}

func (s *CreateRepoSyncTaskByRuleResponseBody) SetIsSuccess(v bool) *CreateRepoSyncTaskByRuleResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *CreateRepoSyncTaskByRuleResponseBody) SetRequestId(v string) *CreateRepoSyncTaskByRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateRepoSyncTaskByRuleResponseBody) SetSyncTaskId(v string) *CreateRepoSyncTaskByRuleResponseBody {
	s.SyncTaskId = &v
	return s
}

type CreateRepoSyncTaskByRuleResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateRepoSyncTaskByRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateRepoSyncTaskByRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateRepoSyncTaskByRuleResponse) GoString() string {
	return s.String()
}

func (s *CreateRepoSyncTaskByRuleResponse) SetHeaders(v map[string]*string) *CreateRepoSyncTaskByRuleResponse {
	s.Headers = v
	return s
}

func (s *CreateRepoSyncTaskByRuleResponse) SetStatusCode(v int32) *CreateRepoSyncTaskByRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateRepoSyncTaskByRuleResponse) SetBody(v *CreateRepoSyncTaskByRuleResponseBody) *CreateRepoSyncTaskByRuleResponse {
	s.Body = v
	return s
}

type CreateRepoTagRequest struct {
	// The source image tag.
	//
	// This parameter is required.
	//
	// example:
	//
	// v1
	FromTag *string `json:"FromTag,omitempty" xml:"FromTag,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cri-shac42yvqzv****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the namespace.
	//
	// This parameter is required.
	//
	// example:
	//
	// ns
	NamespaceName *string `json:"NamespaceName,omitempty" xml:"NamespaceName,omitempty"`
	// The name of the image repository.
	//
	// This parameter is required.
	//
	// example:
	//
	// repo1
	RepoName *string `json:"RepoName,omitempty" xml:"RepoName,omitempty"`
	// The image tag that you want to create.
	//
	// This parameter is required.
	//
	// example:
	//
	// v2
	ToTag *string `json:"ToTag,omitempty" xml:"ToTag,omitempty"`
}

func (s CreateRepoTagRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateRepoTagRequest) GoString() string {
	return s.String()
}

func (s *CreateRepoTagRequest) SetFromTag(v string) *CreateRepoTagRequest {
	s.FromTag = &v
	return s
}

func (s *CreateRepoTagRequest) SetInstanceId(v string) *CreateRepoTagRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateRepoTagRequest) SetNamespaceName(v string) *CreateRepoTagRequest {
	s.NamespaceName = &v
	return s
}

func (s *CreateRepoTagRequest) SetRepoName(v string) *CreateRepoTagRequest {
	s.RepoName = &v
	return s
}

func (s *CreateRepoTagRequest) SetToTag(v string) *CreateRepoTagRequest {
	s.ToTag = &v
	return s
}

type CreateRepoTagResponseBody struct {
	// The return value.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// C4C7DD0C-C9D6-437A-A7EE-8BY*****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateRepoTagResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateRepoTagResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRepoTagResponseBody) SetCode(v string) *CreateRepoTagResponseBody {
	s.Code = &v
	return s
}

func (s *CreateRepoTagResponseBody) SetIsSuccess(v bool) *CreateRepoTagResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *CreateRepoTagResponseBody) SetRequestId(v string) *CreateRepoTagResponseBody {
	s.RequestId = &v
	return s
}

type CreateRepoTagResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateRepoTagResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateRepoTagResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateRepoTagResponse) GoString() string {
	return s.String()
}

func (s *CreateRepoTagResponse) SetHeaders(v map[string]*string) *CreateRepoTagResponse {
	s.Headers = v
	return s
}

func (s *CreateRepoTagResponse) SetStatusCode(v int32) *CreateRepoTagResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateRepoTagResponse) SetBody(v *CreateRepoTagResponseBody) *CreateRepoTagResponse {
	s.Body = v
	return s
}

type CreateRepoTagScanTaskRequest struct {
	// The digest of the image.
	//
	// example:
	//
	// sha256:815386ebbe9a3490f38785ab11bda34ec8dacf4634af77b8912832d4f85dca04
	Digest *string `json:"Digest,omitempty" xml:"Digest,omitempty"`
	// The ID of the Container Registry instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cri-xkx6vujuhay0****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the image repository.
	//
	// This parameter is required.
	//
	// example:
	//
	// crr-xwvi3osiy4ff****
	RepoId *string `json:"RepoId,omitempty" xml:"RepoId,omitempty"`
	// The type of the scanning engine.
	//
	// 	- `SAS_SCAN_SERVICE`: Security Center scan engine (paid service)
	//
	// 	- `ACR_SCAN_SERVICE`: Container Registry scan engine
	//
	// example:
	//
	// ACR_SCAN_SERVICE
	ScanService *string `json:"ScanService,omitempty" xml:"ScanService,omitempty"`
	ScanType    *string `json:"ScanType,omitempty" xml:"ScanType,omitempty"`
	// The image version.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1.24
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
}

func (s CreateRepoTagScanTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateRepoTagScanTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateRepoTagScanTaskRequest) SetDigest(v string) *CreateRepoTagScanTaskRequest {
	s.Digest = &v
	return s
}

func (s *CreateRepoTagScanTaskRequest) SetInstanceId(v string) *CreateRepoTagScanTaskRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateRepoTagScanTaskRequest) SetRepoId(v string) *CreateRepoTagScanTaskRequest {
	s.RepoId = &v
	return s
}

func (s *CreateRepoTagScanTaskRequest) SetScanService(v string) *CreateRepoTagScanTaskRequest {
	s.ScanService = &v
	return s
}

func (s *CreateRepoTagScanTaskRequest) SetScanType(v string) *CreateRepoTagScanTaskRequest {
	s.ScanType = &v
	return s
}

func (s *CreateRepoTagScanTaskRequest) SetTag(v string) *CreateRepoTagScanTaskRequest {
	s.Tag = &v
	return s
}

type CreateRepoTagScanTaskResponseBody struct {
	// The return value.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the API request is successful. Valid values:
	//
	// 	- `true`: The request is successful.
	//
	// 	- `false`: The request fails.
	//
	// example:
	//
	// true
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The request ID.
	//
	// example:
	//
	// BC648259-91A7-4502-BED3-EDF64361FA83
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateRepoTagScanTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateRepoTagScanTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRepoTagScanTaskResponseBody) SetCode(v string) *CreateRepoTagScanTaskResponseBody {
	s.Code = &v
	return s
}

func (s *CreateRepoTagScanTaskResponseBody) SetIsSuccess(v bool) *CreateRepoTagScanTaskResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *CreateRepoTagScanTaskResponseBody) SetRequestId(v string) *CreateRepoTagScanTaskResponseBody {
	s.RequestId = &v
	return s
}

type CreateRepoTagScanTaskResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateRepoTagScanTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateRepoTagScanTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateRepoTagScanTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateRepoTagScanTaskResponse) SetHeaders(v map[string]*string) *CreateRepoTagScanTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateRepoTagScanTaskResponse) SetStatusCode(v int32) *CreateRepoTagScanTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateRepoTagScanTaskResponse) SetBody(v *CreateRepoTagScanTaskResponseBody) *CreateRepoTagScanTaskResponse {
	s.Body = v
	return s
}

type CreateRepoTriggerRequest struct {
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cri-xkx6vujuhay0****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the image repository.
	//
	// This parameter is required.
	//
	// example:
	//
	// crr-xwvi3osiy4ff****
	RepoId *string `json:"RepoId,omitempty" xml:"RepoId,omitempty"`
	// The name of the trigger.
	//
	// This parameter is required.
	//
	// example:
	//
	// trigger1
	TriggerName *string `json:"TriggerName,omitempty" xml:"TriggerName,omitempty"`
	// The image tag based on which the trigger is set.
	//
	// >
	//
	// 	- If `TriggerType` is set to `ALL`, `TriggerTag` can be set to a string or an array, for example, `*`.
	//
	// 	- If `TriggerType` is set to `TAG_LIST`, `TriggerTag` must be set to an array, for example, `[1]`.
	//
	// 	- If `TriggerType` is set to `TAG_REG_EXP`, `TriggerTag` must be set to a string, for example, `*`.
	//
	// example:
	//
	// [1]
	TriggerTag *string `json:"TriggerTag,omitempty" xml:"TriggerTag,omitempty"`
	// The type of the trigger. Valid values:
	//
	// 	- `ALL`: a trigger that supports both tags and regular expressions.
	//
	// 	- `TAG_LIST`: a tag-based trigger.
	//
	// 	- `TAG_REG_EXP`: a regular expression-based trigger.
	//
	// This parameter is required.
	//
	// example:
	//
	// ALL
	TriggerType *string `json:"TriggerType,omitempty" xml:"TriggerType,omitempty"`
	// The URL of the trigger.
	//
	// This parameter is required.
	//
	// example:
	//
	// http://www.mysite.com
	TriggerUrl *string `json:"TriggerUrl,omitempty" xml:"TriggerUrl,omitempty"`
}

func (s CreateRepoTriggerRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateRepoTriggerRequest) GoString() string {
	return s.String()
}

func (s *CreateRepoTriggerRequest) SetInstanceId(v string) *CreateRepoTriggerRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateRepoTriggerRequest) SetRepoId(v string) *CreateRepoTriggerRequest {
	s.RepoId = &v
	return s
}

func (s *CreateRepoTriggerRequest) SetTriggerName(v string) *CreateRepoTriggerRequest {
	s.TriggerName = &v
	return s
}

func (s *CreateRepoTriggerRequest) SetTriggerTag(v string) *CreateRepoTriggerRequest {
	s.TriggerTag = &v
	return s
}

func (s *CreateRepoTriggerRequest) SetTriggerType(v string) *CreateRepoTriggerRequest {
	s.TriggerType = &v
	return s
}

func (s *CreateRepoTriggerRequest) SetTriggerUrl(v string) *CreateRepoTriggerRequest {
	s.TriggerUrl = &v
	return s
}

type CreateRepoTriggerResponseBody struct {
	// The return value.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// B79F5E0E-8770-407D-BCB6-ECF4BA9C****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the trigger.
	//
	// example:
	//
	// crw-0z4pf81pgz35****
	TriggerId *string `json:"TriggerId,omitempty" xml:"TriggerId,omitempty"`
}

func (s CreateRepoTriggerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateRepoTriggerResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRepoTriggerResponseBody) SetCode(v string) *CreateRepoTriggerResponseBody {
	s.Code = &v
	return s
}

func (s *CreateRepoTriggerResponseBody) SetIsSuccess(v bool) *CreateRepoTriggerResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *CreateRepoTriggerResponseBody) SetRequestId(v string) *CreateRepoTriggerResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateRepoTriggerResponseBody) SetTriggerId(v string) *CreateRepoTriggerResponseBody {
	s.TriggerId = &v
	return s
}

type CreateRepoTriggerResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateRepoTriggerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateRepoTriggerResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateRepoTriggerResponse) GoString() string {
	return s.String()
}

func (s *CreateRepoTriggerResponse) SetHeaders(v map[string]*string) *CreateRepoTriggerResponse {
	s.Headers = v
	return s
}

func (s *CreateRepoTriggerResponse) SetStatusCode(v int32) *CreateRepoTriggerResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateRepoTriggerResponse) SetBody(v *CreateRepoTriggerResponseBody) *CreateRepoTriggerResponse {
	s.Body = v
	return s
}

type CreateRepositoryRequest struct {
	// The description of the repository.
	//
	// example:
	//
	// repo1
	Detail *string `json:"Detail,omitempty" xml:"Detail,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cri-xkx6vujuhay0****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the image repository.
	//
	// This parameter is required.
	//
	// example:
	//
	// repo1
	RepoName *string `json:"RepoName,omitempty" xml:"RepoName,omitempty"`
	// The name of the namespace to which the image repository belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// namespace01
	RepoNamespaceName *string `json:"RepoNamespaceName,omitempty" xml:"RepoNamespaceName,omitempty"`
	// The type of the repository. Valid values:
	//
	// 	- `PUBLIC`: The repository is a public repository.
	//
	// 	- `PRIVATE`: The repository is a private repository.
	//
	// This parameter is required.
	//
	// example:
	//
	// PRIVATE
	RepoType *string `json:"RepoType,omitempty" xml:"RepoType,omitempty"`
	// The summary about the repository.
	//
	// This parameter is required.
	//
	// example:
	//
	// repo1
	Summary *string `json:"Summary,omitempty" xml:"Summary,omitempty"`
	// Specifies whether to enable the feature of image tag immutability. Valid values:
	//
	// 	- `true`: enables the feature of image tag immutability.
	//
	// 	- `false`: disables the feature of image tag immutability.
	//
	// example:
	//
	// true
	TagImmutability *bool `json:"TagImmutability,omitempty" xml:"TagImmutability,omitempty"`
}

func (s CreateRepositoryRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateRepositoryRequest) GoString() string {
	return s.String()
}

func (s *CreateRepositoryRequest) SetDetail(v string) *CreateRepositoryRequest {
	s.Detail = &v
	return s
}

func (s *CreateRepositoryRequest) SetInstanceId(v string) *CreateRepositoryRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateRepositoryRequest) SetRepoName(v string) *CreateRepositoryRequest {
	s.RepoName = &v
	return s
}

func (s *CreateRepositoryRequest) SetRepoNamespaceName(v string) *CreateRepositoryRequest {
	s.RepoNamespaceName = &v
	return s
}

func (s *CreateRepositoryRequest) SetRepoType(v string) *CreateRepositoryRequest {
	s.RepoType = &v
	return s
}

func (s *CreateRepositoryRequest) SetSummary(v string) *CreateRepositoryRequest {
	s.Summary = &v
	return s
}

func (s *CreateRepositoryRequest) SetTagImmutability(v bool) *CreateRepositoryRequest {
	s.TagImmutability = &v
	return s
}

type CreateRepositoryResponseBody struct {
	// The return value.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The ID of the image repository.
	//
	// example:
	//
	// crr-xwvi3osiy4ff****
	RepoId *string `json:"RepoId,omitempty" xml:"RepoId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 886FB272-15C3-44FC-AA54-F4ABD5B93A28
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateRepositoryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateRepositoryResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRepositoryResponseBody) SetCode(v string) *CreateRepositoryResponseBody {
	s.Code = &v
	return s
}

func (s *CreateRepositoryResponseBody) SetIsSuccess(v bool) *CreateRepositoryResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *CreateRepositoryResponseBody) SetRepoId(v string) *CreateRepositoryResponseBody {
	s.RepoId = &v
	return s
}

func (s *CreateRepositoryResponseBody) SetRequestId(v string) *CreateRepositoryResponseBody {
	s.RequestId = &v
	return s
}

type CreateRepositoryResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateRepositoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateRepositoryResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateRepositoryResponse) GoString() string {
	return s.String()
}

func (s *CreateRepositoryResponse) SetHeaders(v map[string]*string) *CreateRepositoryResponse {
	s.Headers = v
	return s
}

func (s *CreateRepositoryResponse) SetStatusCode(v int32) *CreateRepositoryResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateRepositoryResponse) SetBody(v *CreateRepositoryResponseBody) *CreateRepositoryResponse {
	s.Body = v
	return s
}

type DeleteArtifactLifecycleRuleRequest struct {
	// The ID of the Container Registry instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cri-brlg4cbj2ylkrqqq
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The rule ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cralr-3v8pao9k7chb8q62
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s DeleteArtifactLifecycleRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteArtifactLifecycleRuleRequest) GoString() string {
	return s.String()
}

func (s *DeleteArtifactLifecycleRuleRequest) SetInstanceId(v string) *DeleteArtifactLifecycleRuleRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteArtifactLifecycleRuleRequest) SetRuleId(v string) *DeleteArtifactLifecycleRuleRequest {
	s.RuleId = &v
	return s
}

type DeleteArtifactLifecycleRuleResponseBody struct {
	// The return value.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// 	- `true`: The request is successful.
	//
	// 	- `false`: The request fails.
	//
	// example:
	//
	// True
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 001AB638-C99B-5A27-8AC9-B2DBABFFEBB5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteArtifactLifecycleRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteArtifactLifecycleRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteArtifactLifecycleRuleResponseBody) SetCode(v string) *DeleteArtifactLifecycleRuleResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteArtifactLifecycleRuleResponseBody) SetIsSuccess(v bool) *DeleteArtifactLifecycleRuleResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *DeleteArtifactLifecycleRuleResponseBody) SetRequestId(v string) *DeleteArtifactLifecycleRuleResponseBody {
	s.RequestId = &v
	return s
}

type DeleteArtifactLifecycleRuleResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteArtifactLifecycleRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteArtifactLifecycleRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteArtifactLifecycleRuleResponse) GoString() string {
	return s.String()
}

func (s *DeleteArtifactLifecycleRuleResponse) SetHeaders(v map[string]*string) *DeleteArtifactLifecycleRuleResponse {
	s.Headers = v
	return s
}

func (s *DeleteArtifactLifecycleRuleResponse) SetStatusCode(v int32) *DeleteArtifactLifecycleRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteArtifactLifecycleRuleResponse) SetBody(v *DeleteArtifactLifecycleRuleResponseBody) *DeleteArtifactLifecycleRuleResponse {
	s.Body = v
	return s
}

type DeleteArtifactSubscriptionRuleRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cri-c0o11woew0k****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The rule ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// crasr-mdbpung4i1rm****
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s DeleteArtifactSubscriptionRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteArtifactSubscriptionRuleRequest) GoString() string {
	return s.String()
}

func (s *DeleteArtifactSubscriptionRuleRequest) SetInstanceId(v string) *DeleteArtifactSubscriptionRuleRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteArtifactSubscriptionRuleRequest) SetRuleId(v string) *DeleteArtifactSubscriptionRuleRequest {
	s.RuleId = &v
	return s
}

type DeleteArtifactSubscriptionRuleResponseBody struct {
	// The return value.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the API request is successful. Valid values:
	//
	// 	- `true`: The request is successful.
	//
	// 	- `false`: The request fails.
	//
	// example:
	//
	// true
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 02B27D80-FD32-5155-931A-93700779BB9E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteArtifactSubscriptionRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteArtifactSubscriptionRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteArtifactSubscriptionRuleResponseBody) SetCode(v string) *DeleteArtifactSubscriptionRuleResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteArtifactSubscriptionRuleResponseBody) SetIsSuccess(v bool) *DeleteArtifactSubscriptionRuleResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *DeleteArtifactSubscriptionRuleResponseBody) SetRequestId(v string) *DeleteArtifactSubscriptionRuleResponseBody {
	s.RequestId = &v
	return s
}

type DeleteArtifactSubscriptionRuleResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteArtifactSubscriptionRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteArtifactSubscriptionRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteArtifactSubscriptionRuleResponse) GoString() string {
	return s.String()
}

func (s *DeleteArtifactSubscriptionRuleResponse) SetHeaders(v map[string]*string) *DeleteArtifactSubscriptionRuleResponse {
	s.Headers = v
	return s
}

func (s *DeleteArtifactSubscriptionRuleResponse) SetStatusCode(v int32) *DeleteArtifactSubscriptionRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteArtifactSubscriptionRuleResponse) SetBody(v *DeleteArtifactSubscriptionRuleResponseBody) *DeleteArtifactSubscriptionRuleResponse {
	s.Body = v
	return s
}

type DeleteChainRequest struct {
	// The ID of the delivery pipeline.
	//
	// This parameter is required.
	//
	// example:
	//
	// chi-02ymhtwl3cq8****
	ChainId *string `json:"ChainId,omitempty" xml:"ChainId,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cri-4cdrlqmhn4gm****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DeleteChainRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteChainRequest) GoString() string {
	return s.String()
}

func (s *DeleteChainRequest) SetChainId(v string) *DeleteChainRequest {
	s.ChainId = &v
	return s
}

func (s *DeleteChainRequest) SetInstanceId(v string) *DeleteChainRequest {
	s.InstanceId = &v
	return s
}

type DeleteChainResponseBody struct {
	// The return value.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// DB1809A8-E1C8-5707-BAF8-D4FC1C11****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteChainResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteChainResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteChainResponseBody) SetCode(v string) *DeleteChainResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteChainResponseBody) SetIsSuccess(v bool) *DeleteChainResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *DeleteChainResponseBody) SetRequestId(v string) *DeleteChainResponseBody {
	s.RequestId = &v
	return s
}

type DeleteChainResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteChainResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteChainResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteChainResponse) GoString() string {
	return s.String()
}

func (s *DeleteChainResponse) SetHeaders(v map[string]*string) *DeleteChainResponse {
	s.Headers = v
	return s
}

func (s *DeleteChainResponse) SetStatusCode(v int32) *DeleteChainResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteChainResponse) SetBody(v *DeleteChainResponseBody) *DeleteChainResponse {
	s.Body = v
	return s
}

type DeleteChartNamespaceRequest struct {
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cri-xkx6vujuhay0****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the chart namespace that you want to delete.
	//
	// This parameter is required.
	//
	// example:
	//
	// ns2
	NamespaceName *string `json:"NamespaceName,omitempty" xml:"NamespaceName,omitempty"`
}

func (s DeleteChartNamespaceRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteChartNamespaceRequest) GoString() string {
	return s.String()
}

func (s *DeleteChartNamespaceRequest) SetInstanceId(v string) *DeleteChartNamespaceRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteChartNamespaceRequest) SetNamespaceName(v string) *DeleteChartNamespaceRequest {
	s.NamespaceName = &v
	return s
}

type DeleteChartNamespaceResponseBody struct {
	// The return value.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// 	- `true`: The request is successful.
	//
	// 	- `false`: The request fails.
	//
	// example:
	//
	// true
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// FEC62DF1-1394-467F-A69F-4BC1BA29F383
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteChartNamespaceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteChartNamespaceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteChartNamespaceResponseBody) SetCode(v string) *DeleteChartNamespaceResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteChartNamespaceResponseBody) SetIsSuccess(v bool) *DeleteChartNamespaceResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *DeleteChartNamespaceResponseBody) SetRequestId(v string) *DeleteChartNamespaceResponseBody {
	s.RequestId = &v
	return s
}

type DeleteChartNamespaceResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteChartNamespaceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteChartNamespaceResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteChartNamespaceResponse) GoString() string {
	return s.String()
}

func (s *DeleteChartNamespaceResponse) SetHeaders(v map[string]*string) *DeleteChartNamespaceResponse {
	s.Headers = v
	return s
}

func (s *DeleteChartNamespaceResponse) SetStatusCode(v int32) *DeleteChartNamespaceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteChartNamespaceResponse) SetBody(v *DeleteChartNamespaceResponseBody) *DeleteChartNamespaceResponse {
	s.Body = v
	return s
}

type DeleteChartReleaseRequest struct {
	// The name of the chart.
	//
	// This parameter is required.
	//
	// example:
	//
	// chart3
	Chart *string `json:"Chart,omitempty" xml:"Chart,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cri-xkx6vujuhay0****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The version of the chart that you want to delete.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0.1.0
	Release *string `json:"Release,omitempty" xml:"Release,omitempty"`
	// The name of the repository.
	//
	// This parameter is required.
	//
	// example:
	//
	// repo1
	RepoName *string `json:"RepoName,omitempty" xml:"RepoName,omitempty"`
	// The name of the namespace.
	//
	// This parameter is required.
	//
	// example:
	//
	// ns1
	RepoNamespaceName *string `json:"RepoNamespaceName,omitempty" xml:"RepoNamespaceName,omitempty"`
}

func (s DeleteChartReleaseRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteChartReleaseRequest) GoString() string {
	return s.String()
}

func (s *DeleteChartReleaseRequest) SetChart(v string) *DeleteChartReleaseRequest {
	s.Chart = &v
	return s
}

func (s *DeleteChartReleaseRequest) SetInstanceId(v string) *DeleteChartReleaseRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteChartReleaseRequest) SetRelease(v string) *DeleteChartReleaseRequest {
	s.Release = &v
	return s
}

func (s *DeleteChartReleaseRequest) SetRepoName(v string) *DeleteChartReleaseRequest {
	s.RepoName = &v
	return s
}

func (s *DeleteChartReleaseRequest) SetRepoNamespaceName(v string) *DeleteChartReleaseRequest {
	s.RepoNamespaceName = &v
	return s
}

type DeleteChartReleaseResponseBody struct {
	// The return value.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// C2D6CE47-6DEF-45F4-A1AC-90F3AFBA751F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteChartReleaseResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteChartReleaseResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteChartReleaseResponseBody) SetCode(v string) *DeleteChartReleaseResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteChartReleaseResponseBody) SetIsSuccess(v bool) *DeleteChartReleaseResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *DeleteChartReleaseResponseBody) SetRequestId(v string) *DeleteChartReleaseResponseBody {
	s.RequestId = &v
	return s
}

type DeleteChartReleaseResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteChartReleaseResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteChartReleaseResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteChartReleaseResponse) GoString() string {
	return s.String()
}

func (s *DeleteChartReleaseResponse) SetHeaders(v map[string]*string) *DeleteChartReleaseResponse {
	s.Headers = v
	return s
}

func (s *DeleteChartReleaseResponse) SetStatusCode(v int32) *DeleteChartReleaseResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteChartReleaseResponse) SetBody(v *DeleteChartReleaseResponseBody) *DeleteChartReleaseResponse {
	s.Body = v
	return s
}

type DeleteChartRepositoryRequest struct {
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cri-xkx6vujuhay****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the repository.
	//
	// This parameter is required.
	//
	// example:
	//
	// repo01
	RepoName *string `json:"RepoName,omitempty" xml:"RepoName,omitempty"`
	// The name of the namespace to which the repository belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// namespace01
	RepoNamespaceName *string `json:"RepoNamespaceName,omitempty" xml:"RepoNamespaceName,omitempty"`
}

func (s DeleteChartRepositoryRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteChartRepositoryRequest) GoString() string {
	return s.String()
}

func (s *DeleteChartRepositoryRequest) SetInstanceId(v string) *DeleteChartRepositoryRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteChartRepositoryRequest) SetRepoName(v string) *DeleteChartRepositoryRequest {
	s.RepoName = &v
	return s
}

func (s *DeleteChartRepositoryRequest) SetRepoNamespaceName(v string) *DeleteChartRepositoryRequest {
	s.RepoNamespaceName = &v
	return s
}

type DeleteChartRepositoryResponseBody struct {
	// The return value.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// 	- `true`: The request is successful.
	//
	// 	- `false`: The request fails.
	//
	// example:
	//
	// true
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 12589EF7-96E2-4554-AAD7-F7209E88CAD3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteChartRepositoryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteChartRepositoryResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteChartRepositoryResponseBody) SetCode(v string) *DeleteChartRepositoryResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteChartRepositoryResponseBody) SetIsSuccess(v bool) *DeleteChartRepositoryResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *DeleteChartRepositoryResponseBody) SetRequestId(v string) *DeleteChartRepositoryResponseBody {
	s.RequestId = &v
	return s
}

type DeleteChartRepositoryResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteChartRepositoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteChartRepositoryResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteChartRepositoryResponse) GoString() string {
	return s.String()
}

func (s *DeleteChartRepositoryResponse) SetHeaders(v map[string]*string) *DeleteChartRepositoryResponse {
	s.Headers = v
	return s
}

func (s *DeleteChartRepositoryResponse) SetStatusCode(v int32) *DeleteChartRepositoryResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteChartRepositoryResponse) SetBody(v *DeleteChartRepositoryResponseBody) *DeleteChartRepositoryResponse {
	s.Body = v
	return s
}

type DeleteEventCenterRuleRequest struct {
	// The ID of the instance.
	//
	// example:
	//
	// cri-xkx6vujuhay0****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the event notification rule.
	//
	// example:
	//
	// crecr-n6pbhgjx*****
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s DeleteEventCenterRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteEventCenterRuleRequest) GoString() string {
	return s.String()
}

func (s *DeleteEventCenterRuleRequest) SetInstanceId(v string) *DeleteEventCenterRuleRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteEventCenterRuleRequest) SetRuleId(v string) *DeleteEventCenterRuleRequest {
	s.RuleId = &v
	return s
}

type DeleteEventCenterRuleResponseBody struct {
	// The return value.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 031572FA-7D8F-4C05-B790-1071E0E05DE6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteEventCenterRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteEventCenterRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteEventCenterRuleResponseBody) SetCode(v string) *DeleteEventCenterRuleResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteEventCenterRuleResponseBody) SetRequestId(v string) *DeleteEventCenterRuleResponseBody {
	s.RequestId = &v
	return s
}

type DeleteEventCenterRuleResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteEventCenterRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteEventCenterRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteEventCenterRuleResponse) GoString() string {
	return s.String()
}

func (s *DeleteEventCenterRuleResponse) SetHeaders(v map[string]*string) *DeleteEventCenterRuleResponse {
	s.Headers = v
	return s
}

func (s *DeleteEventCenterRuleResponse) SetStatusCode(v int32) *DeleteEventCenterRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteEventCenterRuleResponse) SetBody(v *DeleteEventCenterRuleResponseBody) *DeleteEventCenterRuleResponse {
	s.Body = v
	return s
}

type DeleteInstanceEndpointAclPolicyRequest struct {
	// The type of the endpoint. Set the value to Internet.
	//
	// This parameter is required.
	//
	// example:
	//
	// internet
	EndpointType *string `json:"EndpointType,omitempty" xml:"EndpointType,omitempty"`
	// The CIDR block.
	//
	// This parameter is required.
	//
	// example:
	//
	// 127.0.0.1/32
	Entry *string `json:"Entry,omitempty" xml:"Entry,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cri-xkx6vujuhay0****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the module that you want to access. Valid values:
	//
	// 	- `Registry`: the image repository.
	//
	// 	- `Chart`: a Helm chart.
	//
	// example:
	//
	// Chart
	ModuleName *string `json:"ModuleName,omitempty" xml:"ModuleName,omitempty"`
}

func (s DeleteInstanceEndpointAclPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteInstanceEndpointAclPolicyRequest) GoString() string {
	return s.String()
}

func (s *DeleteInstanceEndpointAclPolicyRequest) SetEndpointType(v string) *DeleteInstanceEndpointAclPolicyRequest {
	s.EndpointType = &v
	return s
}

func (s *DeleteInstanceEndpointAclPolicyRequest) SetEntry(v string) *DeleteInstanceEndpointAclPolicyRequest {
	s.Entry = &v
	return s
}

func (s *DeleteInstanceEndpointAclPolicyRequest) SetInstanceId(v string) *DeleteInstanceEndpointAclPolicyRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteInstanceEndpointAclPolicyRequest) SetModuleName(v string) *DeleteInstanceEndpointAclPolicyRequest {
	s.ModuleName = &v
	return s
}

type DeleteInstanceEndpointAclPolicyResponseBody struct {
	// The return value.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// 	- `true`: The request is successful.
	//
	// 	- `false`: The request fails.
	//
	// example:
	//
	// true
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// BDB1F145-F0FF-44E9-AADF-A678642A7C7D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteInstanceEndpointAclPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteInstanceEndpointAclPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteInstanceEndpointAclPolicyResponseBody) SetCode(v string) *DeleteInstanceEndpointAclPolicyResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteInstanceEndpointAclPolicyResponseBody) SetIsSuccess(v bool) *DeleteInstanceEndpointAclPolicyResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *DeleteInstanceEndpointAclPolicyResponseBody) SetRequestId(v string) *DeleteInstanceEndpointAclPolicyResponseBody {
	s.RequestId = &v
	return s
}

type DeleteInstanceEndpointAclPolicyResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteInstanceEndpointAclPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteInstanceEndpointAclPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteInstanceEndpointAclPolicyResponse) GoString() string {
	return s.String()
}

func (s *DeleteInstanceEndpointAclPolicyResponse) SetHeaders(v map[string]*string) *DeleteInstanceEndpointAclPolicyResponse {
	s.Headers = v
	return s
}

func (s *DeleteInstanceEndpointAclPolicyResponse) SetStatusCode(v int32) *DeleteInstanceEndpointAclPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteInstanceEndpointAclPolicyResponse) SetBody(v *DeleteInstanceEndpointAclPolicyResponseBody) *DeleteInstanceEndpointAclPolicyResponse {
	s.Body = v
	return s
}

type DeleteInstanceVpcEndpointLinkedVpcRequest struct {
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cri-xkx6vujuhay0****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the module that you want to access. Valid values:
	//
	// 	- `Registry`: the image repository.
	//
	// 	- `Chart`: a Helm chart.
	//
	// example:
	//
	// Chart
	ModuleName *string `json:"ModuleName,omitempty" xml:"ModuleName,omitempty"`
	// The ID of the VPC.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpc-uf6pa68zxnnlc48dd****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The ID of the vSwitch.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpc-uf6pa68zxnnlc48dd****
	VswitchId *string `json:"VswitchId,omitempty" xml:"VswitchId,omitempty"`
}

func (s DeleteInstanceVpcEndpointLinkedVpcRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteInstanceVpcEndpointLinkedVpcRequest) GoString() string {
	return s.String()
}

func (s *DeleteInstanceVpcEndpointLinkedVpcRequest) SetInstanceId(v string) *DeleteInstanceVpcEndpointLinkedVpcRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteInstanceVpcEndpointLinkedVpcRequest) SetModuleName(v string) *DeleteInstanceVpcEndpointLinkedVpcRequest {
	s.ModuleName = &v
	return s
}

func (s *DeleteInstanceVpcEndpointLinkedVpcRequest) SetVpcId(v string) *DeleteInstanceVpcEndpointLinkedVpcRequest {
	s.VpcId = &v
	return s
}

func (s *DeleteInstanceVpcEndpointLinkedVpcRequest) SetVswitchId(v string) *DeleteInstanceVpcEndpointLinkedVpcRequest {
	s.VswitchId = &v
	return s
}

type DeleteInstanceVpcEndpointLinkedVpcResponseBody struct {
	// The return value.
	//
	// example:
	//
	// true
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// 	- `true`: The request is successful.
	//
	// 	- `false`: The request fails.
	//
	// example:
	//
	// true
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 20FE7A66-0044-4E23-BBEC-C434EADBD7AF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteInstanceVpcEndpointLinkedVpcResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteInstanceVpcEndpointLinkedVpcResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteInstanceVpcEndpointLinkedVpcResponseBody) SetCode(v string) *DeleteInstanceVpcEndpointLinkedVpcResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteInstanceVpcEndpointLinkedVpcResponseBody) SetIsSuccess(v bool) *DeleteInstanceVpcEndpointLinkedVpcResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *DeleteInstanceVpcEndpointLinkedVpcResponseBody) SetRequestId(v string) *DeleteInstanceVpcEndpointLinkedVpcResponseBody {
	s.RequestId = &v
	return s
}

type DeleteInstanceVpcEndpointLinkedVpcResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteInstanceVpcEndpointLinkedVpcResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteInstanceVpcEndpointLinkedVpcResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteInstanceVpcEndpointLinkedVpcResponse) GoString() string {
	return s.String()
}

func (s *DeleteInstanceVpcEndpointLinkedVpcResponse) SetHeaders(v map[string]*string) *DeleteInstanceVpcEndpointLinkedVpcResponse {
	s.Headers = v
	return s
}

func (s *DeleteInstanceVpcEndpointLinkedVpcResponse) SetStatusCode(v int32) *DeleteInstanceVpcEndpointLinkedVpcResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteInstanceVpcEndpointLinkedVpcResponse) SetBody(v *DeleteInstanceVpcEndpointLinkedVpcResponseBody) *DeleteInstanceVpcEndpointLinkedVpcResponse {
	s.Body = v
	return s
}

type DeleteNamespaceRequest struct {
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cri-xkx6vujuhay0****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the namespace.
	//
	// This parameter is required.
	//
	// example:
	//
	// ns3
	NamespaceName *string `json:"NamespaceName,omitempty" xml:"NamespaceName,omitempty"`
}

func (s DeleteNamespaceRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteNamespaceRequest) GoString() string {
	return s.String()
}

func (s *DeleteNamespaceRequest) SetInstanceId(v string) *DeleteNamespaceRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteNamespaceRequest) SetNamespaceName(v string) *DeleteNamespaceRequest {
	s.NamespaceName = &v
	return s
}

type DeleteNamespaceResponseBody struct {
	// The return value.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// BA08C185-8F76-48D7-ACB3-BA11BF2778F9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteNamespaceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteNamespaceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteNamespaceResponseBody) SetCode(v string) *DeleteNamespaceResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteNamespaceResponseBody) SetIsSuccess(v bool) *DeleteNamespaceResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *DeleteNamespaceResponseBody) SetRequestId(v string) *DeleteNamespaceResponseBody {
	s.RequestId = &v
	return s
}

type DeleteNamespaceResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteNamespaceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteNamespaceResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteNamespaceResponse) GoString() string {
	return s.String()
}

func (s *DeleteNamespaceResponse) SetHeaders(v map[string]*string) *DeleteNamespaceResponse {
	s.Headers = v
	return s
}

func (s *DeleteNamespaceResponse) SetStatusCode(v int32) *DeleteNamespaceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteNamespaceResponse) SetBody(v *DeleteNamespaceResponseBody) *DeleteNamespaceResponse {
	s.Body = v
	return s
}

type DeleteRepoBuildRuleRequest struct {
	// The ID of the image building rule.
	//
	// This parameter is required.
	//
	// example:
	//
	// crbr-36tffn0kouvi****
	BuildRuleId *string `json:"BuildRuleId,omitempty" xml:"BuildRuleId,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cri-xkx6vujuhay0****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the image repository.
	//
	// This parameter is required.
	//
	// example:
	//
	// crr-xwvi3osiy4ff****
	RepoId *string `json:"RepoId,omitempty" xml:"RepoId,omitempty"`
}

func (s DeleteRepoBuildRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteRepoBuildRuleRequest) GoString() string {
	return s.String()
}

func (s *DeleteRepoBuildRuleRequest) SetBuildRuleId(v string) *DeleteRepoBuildRuleRequest {
	s.BuildRuleId = &v
	return s
}

func (s *DeleteRepoBuildRuleRequest) SetInstanceId(v string) *DeleteRepoBuildRuleRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteRepoBuildRuleRequest) SetRepoId(v string) *DeleteRepoBuildRuleRequest {
	s.RepoId = &v
	return s
}

type DeleteRepoBuildRuleResponseBody struct {
	// The return value.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 2E3F55BF-FA7B-454E-B2C6-85265E243ADC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteRepoBuildRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteRepoBuildRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteRepoBuildRuleResponseBody) SetCode(v string) *DeleteRepoBuildRuleResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteRepoBuildRuleResponseBody) SetIsSuccess(v bool) *DeleteRepoBuildRuleResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *DeleteRepoBuildRuleResponseBody) SetRequestId(v string) *DeleteRepoBuildRuleResponseBody {
	s.RequestId = &v
	return s
}

type DeleteRepoBuildRuleResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteRepoBuildRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteRepoBuildRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteRepoBuildRuleResponse) GoString() string {
	return s.String()
}

func (s *DeleteRepoBuildRuleResponse) SetHeaders(v map[string]*string) *DeleteRepoBuildRuleResponse {
	s.Headers = v
	return s
}

func (s *DeleteRepoBuildRuleResponse) SetStatusCode(v int32) *DeleteRepoBuildRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteRepoBuildRuleResponse) SetBody(v *DeleteRepoBuildRuleResponseBody) *DeleteRepoBuildRuleResponse {
	s.Body = v
	return s
}

type DeleteRepoSyncRuleRequest struct {
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cri-hpdfkc6utbaq****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the synchronization rule.
	//
	// This parameter is required.
	//
	// example:
	//
	// crsr-gk5p2ns1kzns****
	SyncRuleId *string `json:"SyncRuleId,omitempty" xml:"SyncRuleId,omitempty"`
}

func (s DeleteRepoSyncRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteRepoSyncRuleRequest) GoString() string {
	return s.String()
}

func (s *DeleteRepoSyncRuleRequest) SetInstanceId(v string) *DeleteRepoSyncRuleRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteRepoSyncRuleRequest) SetSyncRuleId(v string) *DeleteRepoSyncRuleRequest {
	s.SyncRuleId = &v
	return s
}

type DeleteRepoSyncRuleResponseBody struct {
	// The return value.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 72DD4198-1BB9-47A3-BC01-EAD1A6D5E173
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteRepoSyncRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteRepoSyncRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteRepoSyncRuleResponseBody) SetCode(v string) *DeleteRepoSyncRuleResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteRepoSyncRuleResponseBody) SetIsSuccess(v bool) *DeleteRepoSyncRuleResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *DeleteRepoSyncRuleResponseBody) SetRequestId(v string) *DeleteRepoSyncRuleResponseBody {
	s.RequestId = &v
	return s
}

type DeleteRepoSyncRuleResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteRepoSyncRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteRepoSyncRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteRepoSyncRuleResponse) GoString() string {
	return s.String()
}

func (s *DeleteRepoSyncRuleResponse) SetHeaders(v map[string]*string) *DeleteRepoSyncRuleResponse {
	s.Headers = v
	return s
}

func (s *DeleteRepoSyncRuleResponse) SetStatusCode(v int32) *DeleteRepoSyncRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteRepoSyncRuleResponse) SetBody(v *DeleteRepoSyncRuleResponseBody) *DeleteRepoSyncRuleResponse {
	s.Body = v
	return s
}

type DeleteRepoTagRequest struct {
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cri-xkx6vujuhay****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the image repository.
	//
	// This parameter is required.
	//
	// example:
	//
	// crr-xwvi3osiy4ff****
	RepoId *string `json:"RepoId,omitempty" xml:"RepoId,omitempty"`
	// The tag of the image.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1.24
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
}

func (s DeleteRepoTagRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteRepoTagRequest) GoString() string {
	return s.String()
}

func (s *DeleteRepoTagRequest) SetInstanceId(v string) *DeleteRepoTagRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteRepoTagRequest) SetRepoId(v string) *DeleteRepoTagRequest {
	s.RepoId = &v
	return s
}

func (s *DeleteRepoTagRequest) SetTag(v string) *DeleteRepoTagRequest {
	s.Tag = &v
	return s
}

type DeleteRepoTagResponseBody struct {
	// The return value.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 96E66B3A-C81A-48BE-ACD6-C0AB1F9313C0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteRepoTagResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteRepoTagResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteRepoTagResponseBody) SetCode(v string) *DeleteRepoTagResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteRepoTagResponseBody) SetIsSuccess(v bool) *DeleteRepoTagResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *DeleteRepoTagResponseBody) SetRequestId(v string) *DeleteRepoTagResponseBody {
	s.RequestId = &v
	return s
}

type DeleteRepoTagResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteRepoTagResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteRepoTagResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteRepoTagResponse) GoString() string {
	return s.String()
}

func (s *DeleteRepoTagResponse) SetHeaders(v map[string]*string) *DeleteRepoTagResponse {
	s.Headers = v
	return s
}

func (s *DeleteRepoTagResponse) SetStatusCode(v int32) *DeleteRepoTagResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteRepoTagResponse) SetBody(v *DeleteRepoTagResponseBody) *DeleteRepoTagResponse {
	s.Body = v
	return s
}

type DeleteRepoTriggerRequest struct {
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cri-xkx6vujuhay0****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the image repository.
	//
	// This parameter is required.
	//
	// example:
	//
	// crr-xwvi3osiy4ff****
	RepoId *string `json:"RepoId,omitempty" xml:"RepoId,omitempty"`
	// The ID of the trigger.
	//
	// This parameter is required.
	//
	// example:
	//
	// crw-0z4pf81pgz35****
	TriggerId *string `json:"TriggerId,omitempty" xml:"TriggerId,omitempty"`
}

func (s DeleteRepoTriggerRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteRepoTriggerRequest) GoString() string {
	return s.String()
}

func (s *DeleteRepoTriggerRequest) SetInstanceId(v string) *DeleteRepoTriggerRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteRepoTriggerRequest) SetRepoId(v string) *DeleteRepoTriggerRequest {
	s.RepoId = &v
	return s
}

func (s *DeleteRepoTriggerRequest) SetTriggerId(v string) *DeleteRepoTriggerRequest {
	s.TriggerId = &v
	return s
}

type DeleteRepoTriggerResponseBody struct {
	// The return value.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// 	- `true`: The request is successful.
	//
	// 	- `false`: The request fails.
	//
	// example:
	//
	// true
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 85180AE4-9A57-48F8-9EF9-68ECCE54B552
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteRepoTriggerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteRepoTriggerResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteRepoTriggerResponseBody) SetCode(v string) *DeleteRepoTriggerResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteRepoTriggerResponseBody) SetIsSuccess(v bool) *DeleteRepoTriggerResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *DeleteRepoTriggerResponseBody) SetRequestId(v string) *DeleteRepoTriggerResponseBody {
	s.RequestId = &v
	return s
}

type DeleteRepoTriggerResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteRepoTriggerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteRepoTriggerResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteRepoTriggerResponse) GoString() string {
	return s.String()
}

func (s *DeleteRepoTriggerResponse) SetHeaders(v map[string]*string) *DeleteRepoTriggerResponse {
	s.Headers = v
	return s
}

func (s *DeleteRepoTriggerResponse) SetStatusCode(v int32) *DeleteRepoTriggerResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteRepoTriggerResponse) SetBody(v *DeleteRepoTriggerResponseBody) *DeleteRepoTriggerResponse {
	s.Body = v
	return s
}

type DeleteRepositoryRequest struct {
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cri-xkx6vujuhay0****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the image repository.
	//
	// example:
	//
	// crr-l4933wbcmun2****
	RepoId *string `json:"RepoId,omitempty" xml:"RepoId,omitempty"`
	// The name of the repository.
	//
	// example:
	//
	// test-repo
	RepoName *string `json:"RepoName,omitempty" xml:"RepoName,omitempty"`
	// The name of the namespace to which the repository belongs.
	//
	// example:
	//
	// test-namespace
	RepoNamespaceName *string `json:"RepoNamespaceName,omitempty" xml:"RepoNamespaceName,omitempty"`
}

func (s DeleteRepositoryRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteRepositoryRequest) GoString() string {
	return s.String()
}

func (s *DeleteRepositoryRequest) SetInstanceId(v string) *DeleteRepositoryRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteRepositoryRequest) SetRepoId(v string) *DeleteRepositoryRequest {
	s.RepoId = &v
	return s
}

func (s *DeleteRepositoryRequest) SetRepoName(v string) *DeleteRepositoryRequest {
	s.RepoName = &v
	return s
}

func (s *DeleteRepositoryRequest) SetRepoNamespaceName(v string) *DeleteRepositoryRequest {
	s.RepoNamespaceName = &v
	return s
}

type DeleteRepositoryResponseBody struct {
	// Return values
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// 	- `true`: The request is successful.
	//
	// 	- `false`: The request fails.
	//
	// example:
	//
	// true
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 47DD9D56-09A0-4C52-B520-C3805DBAB96B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteRepositoryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteRepositoryResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteRepositoryResponseBody) SetCode(v string) *DeleteRepositoryResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteRepositoryResponseBody) SetIsSuccess(v bool) *DeleteRepositoryResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *DeleteRepositoryResponseBody) SetRequestId(v string) *DeleteRepositoryResponseBody {
	s.RequestId = &v
	return s
}

type DeleteRepositoryResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteRepositoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteRepositoryResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteRepositoryResponse) GoString() string {
	return s.String()
}

func (s *DeleteRepositoryResponse) SetHeaders(v map[string]*string) *DeleteRepositoryResponse {
	s.Headers = v
	return s
}

func (s *DeleteRepositoryResponse) SetStatusCode(v int32) *DeleteRepositoryResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteRepositoryResponse) SetBody(v *DeleteRepositoryResponseBody) *DeleteRepositoryResponse {
	s.Body = v
	return s
}

type GetArtifactBuildRuleRequest struct {
	// example:
	//
	// ACCELERATED_IMAGE
	ArtifactType *string `json:"ArtifactType,omitempty" xml:"ArtifactType,omitempty"`
	// example:
	//
	// crabr-o2670wqz2n70****
	BuildRuleId *string `json:"BuildRuleId,omitempty" xml:"BuildRuleId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cri-xkx6vujuhay0****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// crr-8dz3aedjqlmk****
	ScopeId *string `json:"ScopeId,omitempty" xml:"ScopeId,omitempty"`
	// example:
	//
	// REPOSITORY
	ScopeType *string `json:"ScopeType,omitempty" xml:"ScopeType,omitempty"`
}

func (s GetArtifactBuildRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s GetArtifactBuildRuleRequest) GoString() string {
	return s.String()
}

func (s *GetArtifactBuildRuleRequest) SetArtifactType(v string) *GetArtifactBuildRuleRequest {
	s.ArtifactType = &v
	return s
}

func (s *GetArtifactBuildRuleRequest) SetBuildRuleId(v string) *GetArtifactBuildRuleRequest {
	s.BuildRuleId = &v
	return s
}

func (s *GetArtifactBuildRuleRequest) SetInstanceId(v string) *GetArtifactBuildRuleRequest {
	s.InstanceId = &v
	return s
}

func (s *GetArtifactBuildRuleRequest) SetScopeId(v string) *GetArtifactBuildRuleRequest {
	s.ScopeId = &v
	return s
}

func (s *GetArtifactBuildRuleRequest) SetScopeType(v string) *GetArtifactBuildRuleRequest {
	s.ScopeType = &v
	return s
}

type GetArtifactBuildRuleResponseBody struct {
	// example:
	//
	// ACCELERATED_IMAGE
	ArtifactType *string `json:"ArtifactType,omitempty" xml:"ArtifactType,omitempty"`
	// example:
	//
	// crabr-o2670wqz2n70****
	BuildRuleId *string `json:"BuildRuleId,omitempty" xml:"BuildRuleId,omitempty"`
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// true
	IsSuccess  *bool                                       `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	Parameters *GetArtifactBuildRuleResponseBodyParameters `json:"Parameters,omitempty" xml:"Parameters,omitempty" type:"Struct"`
	// example:
	//
	// 7A3E98F6-296C-54AC-A612-B75E7777D4C1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// crr-8dz3aedjqlmk****
	ScopeId *string `json:"ScopeId,omitempty" xml:"ScopeId,omitempty"`
	// example:
	//
	// REPOSITORY
	ScopeType *string `json:"ScopeType,omitempty" xml:"ScopeType,omitempty"`
}

func (s GetArtifactBuildRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetArtifactBuildRuleResponseBody) GoString() string {
	return s.String()
}

func (s *GetArtifactBuildRuleResponseBody) SetArtifactType(v string) *GetArtifactBuildRuleResponseBody {
	s.ArtifactType = &v
	return s
}

func (s *GetArtifactBuildRuleResponseBody) SetBuildRuleId(v string) *GetArtifactBuildRuleResponseBody {
	s.BuildRuleId = &v
	return s
}

func (s *GetArtifactBuildRuleResponseBody) SetCode(v string) *GetArtifactBuildRuleResponseBody {
	s.Code = &v
	return s
}

func (s *GetArtifactBuildRuleResponseBody) SetIsSuccess(v bool) *GetArtifactBuildRuleResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *GetArtifactBuildRuleResponseBody) SetParameters(v *GetArtifactBuildRuleResponseBodyParameters) *GetArtifactBuildRuleResponseBody {
	s.Parameters = v
	return s
}

func (s *GetArtifactBuildRuleResponseBody) SetRequestId(v string) *GetArtifactBuildRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetArtifactBuildRuleResponseBody) SetScopeId(v string) *GetArtifactBuildRuleResponseBody {
	s.ScopeId = &v
	return s
}

func (s *GetArtifactBuildRuleResponseBody) SetScopeType(v string) *GetArtifactBuildRuleResponseBody {
	s.ScopeType = &v
	return s
}

type GetArtifactBuildRuleResponseBodyParameters struct {
	ImageIndexOnly *bool   `json:"ImageIndexOnly,omitempty" xml:"ImageIndexOnly,omitempty"`
	PriorityFile   *string `json:"PriorityFile,omitempty" xml:"PriorityFile,omitempty"`
}

func (s GetArtifactBuildRuleResponseBodyParameters) String() string {
	return tea.Prettify(s)
}

func (s GetArtifactBuildRuleResponseBodyParameters) GoString() string {
	return s.String()
}

func (s *GetArtifactBuildRuleResponseBodyParameters) SetImageIndexOnly(v bool) *GetArtifactBuildRuleResponseBodyParameters {
	s.ImageIndexOnly = &v
	return s
}

func (s *GetArtifactBuildRuleResponseBodyParameters) SetPriorityFile(v string) *GetArtifactBuildRuleResponseBodyParameters {
	s.PriorityFile = &v
	return s
}

type GetArtifactBuildRuleResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetArtifactBuildRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetArtifactBuildRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s GetArtifactBuildRuleResponse) GoString() string {
	return s.String()
}

func (s *GetArtifactBuildRuleResponse) SetHeaders(v map[string]*string) *GetArtifactBuildRuleResponse {
	s.Headers = v
	return s
}

func (s *GetArtifactBuildRuleResponse) SetStatusCode(v int32) *GetArtifactBuildRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *GetArtifactBuildRuleResponse) SetBody(v *GetArtifactBuildRuleResponseBody) *GetArtifactBuildRuleResponse {
	s.Body = v
	return s
}

type GetArtifactBuildTaskRequest struct {
	// The ID of the artifact building task.
	//
	// This parameter is required.
	//
	// example:
	//
	// i2a-1yu****
	BuildTaskId *string `json:"BuildTaskId,omitempty" xml:"BuildTaskId,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cri-shac42yvqzvq****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetArtifactBuildTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s GetArtifactBuildTaskRequest) GoString() string {
	return s.String()
}

func (s *GetArtifactBuildTaskRequest) SetBuildTaskId(v string) *GetArtifactBuildTaskRequest {
	s.BuildTaskId = &v
	return s
}

func (s *GetArtifactBuildTaskRequest) SetInstanceId(v string) *GetArtifactBuildTaskRequest {
	s.InstanceId = &v
	return s
}

type GetArtifactBuildTaskResponseBody struct {
	// The type of the artifact building task. Valid values:
	//
	// 	- `IMAGE_TO_ACCELERATED_IMAGE`: builds accelerated images for Container Service for Kubernetes (ACK) clusters.
	//
	// 	- `IMAGE_TO_ECI_ACCELERATED_IMAGE`: builds accelerated images for elastic container instances.
	//
	// example:
	//
	// IMAGE_TO_ACCELERATED_IMAGE
	ArtifactBuildType *string `json:"ArtifactBuildType,omitempty" xml:"ArtifactBuildType,omitempty"`
	// The ID of the artifact building task.
	//
	// example:
	//
	// i2a-1yu****
	BuildTaskId *string `json:"BuildTaskId,omitempty" xml:"BuildTaskId,omitempty"`
	// The return value.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The time when the artifact building task ends.
	//
	// example:
	//
	// 156871880
	EndTime      *int32    `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Instructions []*string `json:"Instructions,omitempty" xml:"Instructions,omitempty" type:"Repeated"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// C4C7DD0C-C9D6-437A-A7EE-121EFD70D002
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the source artifact.
	SourceArtifact *GetArtifactBuildTaskResponseBodySourceArtifact `json:"SourceArtifact,omitempty" xml:"SourceArtifact,omitempty" type:"Struct"`
	// The time when the artifact building task starts.
	//
	// example:
	//
	// 156871881
	StartTime *int32 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The artifact that is built in the task.
	TargetArtifact *GetArtifactBuildTaskResponseBodyTargetArtifact `json:"TargetArtifact,omitempty" xml:"TargetArtifact,omitempty" type:"Struct"`
	// The status of the artifact that is built in the task. Valid values:
	//
	// 	- `PENDING`: The artifact is being scheduled.
	//
	// 	- `BUILDING`: The artifact is being built.
	//
	// 	- `SUCCESS`: The artifact is built.
	//
	// 	- `FAILED`: The artifact fails to be built.
	//
	// example:
	//
	// BUILDING
	TaskStatus *string `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
}

func (s GetArtifactBuildTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetArtifactBuildTaskResponseBody) GoString() string {
	return s.String()
}

func (s *GetArtifactBuildTaskResponseBody) SetArtifactBuildType(v string) *GetArtifactBuildTaskResponseBody {
	s.ArtifactBuildType = &v
	return s
}

func (s *GetArtifactBuildTaskResponseBody) SetBuildTaskId(v string) *GetArtifactBuildTaskResponseBody {
	s.BuildTaskId = &v
	return s
}

func (s *GetArtifactBuildTaskResponseBody) SetCode(v string) *GetArtifactBuildTaskResponseBody {
	s.Code = &v
	return s
}

func (s *GetArtifactBuildTaskResponseBody) SetEndTime(v int32) *GetArtifactBuildTaskResponseBody {
	s.EndTime = &v
	return s
}

func (s *GetArtifactBuildTaskResponseBody) SetInstructions(v []*string) *GetArtifactBuildTaskResponseBody {
	s.Instructions = v
	return s
}

func (s *GetArtifactBuildTaskResponseBody) SetIsSuccess(v bool) *GetArtifactBuildTaskResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *GetArtifactBuildTaskResponseBody) SetRequestId(v string) *GetArtifactBuildTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetArtifactBuildTaskResponseBody) SetSourceArtifact(v *GetArtifactBuildTaskResponseBodySourceArtifact) *GetArtifactBuildTaskResponseBody {
	s.SourceArtifact = v
	return s
}

func (s *GetArtifactBuildTaskResponseBody) SetStartTime(v int32) *GetArtifactBuildTaskResponseBody {
	s.StartTime = &v
	return s
}

func (s *GetArtifactBuildTaskResponseBody) SetTargetArtifact(v *GetArtifactBuildTaskResponseBodyTargetArtifact) *GetArtifactBuildTaskResponseBody {
	s.TargetArtifact = v
	return s
}

func (s *GetArtifactBuildTaskResponseBody) SetTaskStatus(v string) *GetArtifactBuildTaskResponseBody {
	s.TaskStatus = &v
	return s
}

type GetArtifactBuildTaskResponseBodySourceArtifact struct {
	// The type of the artifact that is built in the task. The value can only be IMAGE.
	//
	// example:
	//
	// IMAGE
	ArtifactType *string `json:"ArtifactType,omitempty" xml:"ArtifactType,omitempty"`
	// The ID of the repository to which the source artifact belongs. The repository can only be an image repository.
	//
	// example:
	//
	// cri-shac42yvqzvq****
	RepoId *string `json:"RepoId,omitempty" xml:"RepoId,omitempty"`
	// The version of the artifact. The artifact can only be an image.
	//
	// example:
	//
	// latest
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s GetArtifactBuildTaskResponseBodySourceArtifact) String() string {
	return tea.Prettify(s)
}

func (s GetArtifactBuildTaskResponseBodySourceArtifact) GoString() string {
	return s.String()
}

func (s *GetArtifactBuildTaskResponseBodySourceArtifact) SetArtifactType(v string) *GetArtifactBuildTaskResponseBodySourceArtifact {
	s.ArtifactType = &v
	return s
}

func (s *GetArtifactBuildTaskResponseBodySourceArtifact) SetRepoId(v string) *GetArtifactBuildTaskResponseBodySourceArtifact {
	s.RepoId = &v
	return s
}

func (s *GetArtifactBuildTaskResponseBodySourceArtifact) SetVersion(v string) *GetArtifactBuildTaskResponseBodySourceArtifact {
	s.Version = &v
	return s
}

type GetArtifactBuildTaskResponseBodyTargetArtifact struct {
	// The type of the artifact that is built in the task. The value can only be IMAGE.
	//
	// example:
	//
	// IMAGE
	ArtifactType *string `json:"ArtifactType,omitempty" xml:"ArtifactType,omitempty"`
	// The ID of the repository to which the artifact that is built in the task belongs. The repository can only be an image repository. The value is the same as the ID of the repository to which the source artifact belongs.
	//
	// example:
	//
	// crr-1234567
	RepoId *string `json:"RepoId,omitempty" xml:"RepoId,omitempty"`
	// The version of the artifact that is built in the task. The artifact can only be an image.
	//
	// example:
	//
	// latest_accelerated
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s GetArtifactBuildTaskResponseBodyTargetArtifact) String() string {
	return tea.Prettify(s)
}

func (s GetArtifactBuildTaskResponseBodyTargetArtifact) GoString() string {
	return s.String()
}

func (s *GetArtifactBuildTaskResponseBodyTargetArtifact) SetArtifactType(v string) *GetArtifactBuildTaskResponseBodyTargetArtifact {
	s.ArtifactType = &v
	return s
}

func (s *GetArtifactBuildTaskResponseBodyTargetArtifact) SetRepoId(v string) *GetArtifactBuildTaskResponseBodyTargetArtifact {
	s.RepoId = &v
	return s
}

func (s *GetArtifactBuildTaskResponseBodyTargetArtifact) SetVersion(v string) *GetArtifactBuildTaskResponseBodyTargetArtifact {
	s.Version = &v
	return s
}

type GetArtifactBuildTaskResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetArtifactBuildTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetArtifactBuildTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s GetArtifactBuildTaskResponse) GoString() string {
	return s.String()
}

func (s *GetArtifactBuildTaskResponse) SetHeaders(v map[string]*string) *GetArtifactBuildTaskResponse {
	s.Headers = v
	return s
}

func (s *GetArtifactBuildTaskResponse) SetStatusCode(v int32) *GetArtifactBuildTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *GetArtifactBuildTaskResponse) SetBody(v *GetArtifactBuildTaskResponseBody) *GetArtifactBuildTaskResponse {
	s.Body = v
	return s
}

type GetArtifactLifecycleRuleRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cri-hpdfkc6utbaq****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The rule ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cralr-a18bkiajy81****
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s GetArtifactLifecycleRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s GetArtifactLifecycleRuleRequest) GoString() string {
	return s.String()
}

func (s *GetArtifactLifecycleRuleRequest) SetInstanceId(v string) *GetArtifactLifecycleRuleRequest {
	s.InstanceId = &v
	return s
}

func (s *GetArtifactLifecycleRuleRequest) SetRuleId(v string) *GetArtifactLifecycleRuleRequest {
	s.RuleId = &v
	return s
}

type GetArtifactLifecycleRuleResponseBody struct {
	// Indicates whether the lifecycle management rule is automatically executed.
	//
	// example:
	//
	// true
	Auto *bool `json:"Auto,omitempty" xml:"Auto,omitempty"`
	// The return value.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The time when the lifecycle management rule was created.
	//
	// example:
	//
	// 1571926439000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// Indicates whether lifecycle management is enabled for the artifact.
	//
	// example:
	//
	// true
	EnableDeleteTag *bool `json:"EnableDeleteTag,omitempty" xml:"EnableDeleteTag,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// cri-xkx6vujuhay0****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Indicates whether the API request is successful. Valid values:
	//
	// 	- `true`: The request is successful.
	//
	// 	- `false`: The request fails.
	//
	// example:
	//
	// true
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The time when the lifecycle management rule was last modified.
	//
	// example:
	//
	// 1638259914000
	ModifiedTime *int64 `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// The name of the namespace.
	//
	// example:
	//
	// test-namespace
	NamespaceName *string `json:"NamespaceName,omitempty" xml:"NamespaceName,omitempty"`
	// The time when the lifecycle management rule is next executed.
	//
	// example:
	//
	// 1701878400000
	NextTime *int64 `json:"NextTime,omitempty" xml:"NextTime,omitempty"`
	// The name of the image repository.
	//
	// example:
	//
	// test-repo
	RepoName *string `json:"RepoName,omitempty" xml:"RepoName,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 724402D0-75CD-4794-BC20-7D37208****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of retained images.
	//
	// example:
	//
	// 30
	RetentionTagCount *int64 `json:"RetentionTagCount,omitempty" xml:"RetentionTagCount,omitempty"`
	// The rule ID.
	//
	// example:
	//
	// cralr-a18bkiajy8****
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// The execution cycle of the lifecycle management rule.
	//
	// example:
	//
	// WEEK
	ScheduleTime *string `json:"ScheduleTime,omitempty" xml:"ScheduleTime,omitempty"`
	// The deletion scope of artifacts.
	//
	// example:
	//
	// INSTANCE
	Scope *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
	// The regular expression that indicates which image tags are retained.
	//
	// example:
	//
	// .*-alpine
	TagRegexp *string `json:"TagRegexp,omitempty" xml:"TagRegexp,omitempty"`
}

func (s GetArtifactLifecycleRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetArtifactLifecycleRuleResponseBody) GoString() string {
	return s.String()
}

func (s *GetArtifactLifecycleRuleResponseBody) SetAuto(v bool) *GetArtifactLifecycleRuleResponseBody {
	s.Auto = &v
	return s
}

func (s *GetArtifactLifecycleRuleResponseBody) SetCode(v string) *GetArtifactLifecycleRuleResponseBody {
	s.Code = &v
	return s
}

func (s *GetArtifactLifecycleRuleResponseBody) SetCreateTime(v int64) *GetArtifactLifecycleRuleResponseBody {
	s.CreateTime = &v
	return s
}

func (s *GetArtifactLifecycleRuleResponseBody) SetEnableDeleteTag(v bool) *GetArtifactLifecycleRuleResponseBody {
	s.EnableDeleteTag = &v
	return s
}

func (s *GetArtifactLifecycleRuleResponseBody) SetInstanceId(v string) *GetArtifactLifecycleRuleResponseBody {
	s.InstanceId = &v
	return s
}

func (s *GetArtifactLifecycleRuleResponseBody) SetIsSuccess(v bool) *GetArtifactLifecycleRuleResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *GetArtifactLifecycleRuleResponseBody) SetModifiedTime(v int64) *GetArtifactLifecycleRuleResponseBody {
	s.ModifiedTime = &v
	return s
}

func (s *GetArtifactLifecycleRuleResponseBody) SetNamespaceName(v string) *GetArtifactLifecycleRuleResponseBody {
	s.NamespaceName = &v
	return s
}

func (s *GetArtifactLifecycleRuleResponseBody) SetNextTime(v int64) *GetArtifactLifecycleRuleResponseBody {
	s.NextTime = &v
	return s
}

func (s *GetArtifactLifecycleRuleResponseBody) SetRepoName(v string) *GetArtifactLifecycleRuleResponseBody {
	s.RepoName = &v
	return s
}

func (s *GetArtifactLifecycleRuleResponseBody) SetRequestId(v string) *GetArtifactLifecycleRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetArtifactLifecycleRuleResponseBody) SetRetentionTagCount(v int64) *GetArtifactLifecycleRuleResponseBody {
	s.RetentionTagCount = &v
	return s
}

func (s *GetArtifactLifecycleRuleResponseBody) SetRuleId(v string) *GetArtifactLifecycleRuleResponseBody {
	s.RuleId = &v
	return s
}

func (s *GetArtifactLifecycleRuleResponseBody) SetScheduleTime(v string) *GetArtifactLifecycleRuleResponseBody {
	s.ScheduleTime = &v
	return s
}

func (s *GetArtifactLifecycleRuleResponseBody) SetScope(v string) *GetArtifactLifecycleRuleResponseBody {
	s.Scope = &v
	return s
}

func (s *GetArtifactLifecycleRuleResponseBody) SetTagRegexp(v string) *GetArtifactLifecycleRuleResponseBody {
	s.TagRegexp = &v
	return s
}

type GetArtifactLifecycleRuleResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetArtifactLifecycleRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetArtifactLifecycleRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s GetArtifactLifecycleRuleResponse) GoString() string {
	return s.String()
}

func (s *GetArtifactLifecycleRuleResponse) SetHeaders(v map[string]*string) *GetArtifactLifecycleRuleResponse {
	s.Headers = v
	return s
}

func (s *GetArtifactLifecycleRuleResponse) SetStatusCode(v int32) *GetArtifactLifecycleRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *GetArtifactLifecycleRuleResponse) SetBody(v *GetArtifactLifecycleRuleResponseBody) *GetArtifactLifecycleRuleResponse {
	s.Body = v
	return s
}

type GetArtifactSubscriptionRuleRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cri-c0o11woew0k****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The rule ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// crasr-mdbpung4i1rm****
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s GetArtifactSubscriptionRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s GetArtifactSubscriptionRuleRequest) GoString() string {
	return s.String()
}

func (s *GetArtifactSubscriptionRuleRequest) SetInstanceId(v string) *GetArtifactSubscriptionRuleRequest {
	s.InstanceId = &v
	return s
}

func (s *GetArtifactSubscriptionRuleRequest) SetRuleId(v string) *GetArtifactSubscriptionRuleRequest {
	s.RuleId = &v
	return s
}

type GetArtifactSubscriptionRuleResponseBody struct {
	// Indicates whether an acceleration link is enabled for image subscription. The subscription acceleration feature is in public preview. The feature is optimized based on scheduling policies and network links to accelerate image subscription.
	//
	// example:
	//
	// true
	Accelerate *bool `json:"Accelerate,omitempty" xml:"Accelerate,omitempty"`
	// The return value.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The time when the subscription rule was created.
	//
	// example:
	//
	// 1570759546000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// cri-hpdfkc6utbaq****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Indicates whether the API request is successful. Valid values:
	//
	// 	- `true`: The request is successful.
	//
	// 	- `false`: The request fails.
	//
	// example:
	//
	// true
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The time when the subscription rule was modified.
	//
	// example:
	//
	// 1638259914000
	ModifiedTime *int64 `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// The name of the Container Registry namespace.
	//
	// example:
	//
	// test-ns
	NamespaceName *string `json:"NamespaceName,omitempty" xml:"NamespaceName,omitempty"`
	// Indicates whether the original image is overwritten.
	//
	// example:
	//
	// true
	Override *bool `json:"Override,omitempty" xml:"Override,omitempty"`
	// The operating system and architecture. If the source repository contains multi-arch images, only the images with the specified operating system and architecture are subscribed to the destination repository of the Enterprise Edition instance.
	Platform []*string `json:"Platform,omitempty" xml:"Platform,omitempty" type:"Repeated"`
	// The name of the Container Registry repository.
	//
	// example:
	//
	// test-repo
	RepoName *string `json:"RepoName,omitempty" xml:"RepoName,omitempty"`
	// The request ID.
	//
	// example:
	//
	// D4978DCC-ECBD-40B0-A714-EE6959B22C77
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The rule ID.
	//
	// example:
	//
	// crasr-mdbpung4i1rm****
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// The name of the source namespace.
	//
	// example:
	//
	// library
	SourceNamespaceName *string `json:"SourceNamespaceName,omitempty" xml:"SourceNamespaceName,omitempty"`
	// The source of the artifact.
	//
	// Valid values:
	//
	// 	- DOCKER_HUB: Docker Hub
	//
	// 	- GCR: GCR
	//
	// 	- QUAY: Quay.io
	//
	// example:
	//
	// DOCKER_HUB
	SourceProvider *string `json:"SourceProvider,omitempty" xml:"SourceProvider,omitempty"`
	// The source repository.
	//
	// example:
	//
	// nginx
	SourceRepoName *string `json:"SourceRepoName,omitempty" xml:"SourceRepoName,omitempty"`
	// The number of subscribed images.
	//
	// example:
	//
	// 1
	TagCount *int64 `json:"TagCount,omitempty" xml:"TagCount,omitempty"`
	// The image tag in the subscription source repository. Regular expressions are supported.
	//
	// example:
	//
	// release-v.*
	TagRegexp *string `json:"TagRegexp,omitempty" xml:"TagRegexp,omitempty"`
}

func (s GetArtifactSubscriptionRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetArtifactSubscriptionRuleResponseBody) GoString() string {
	return s.String()
}

func (s *GetArtifactSubscriptionRuleResponseBody) SetAccelerate(v bool) *GetArtifactSubscriptionRuleResponseBody {
	s.Accelerate = &v
	return s
}

func (s *GetArtifactSubscriptionRuleResponseBody) SetCode(v string) *GetArtifactSubscriptionRuleResponseBody {
	s.Code = &v
	return s
}

func (s *GetArtifactSubscriptionRuleResponseBody) SetCreateTime(v int64) *GetArtifactSubscriptionRuleResponseBody {
	s.CreateTime = &v
	return s
}

func (s *GetArtifactSubscriptionRuleResponseBody) SetInstanceId(v string) *GetArtifactSubscriptionRuleResponseBody {
	s.InstanceId = &v
	return s
}

func (s *GetArtifactSubscriptionRuleResponseBody) SetIsSuccess(v bool) *GetArtifactSubscriptionRuleResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *GetArtifactSubscriptionRuleResponseBody) SetModifiedTime(v int64) *GetArtifactSubscriptionRuleResponseBody {
	s.ModifiedTime = &v
	return s
}

func (s *GetArtifactSubscriptionRuleResponseBody) SetNamespaceName(v string) *GetArtifactSubscriptionRuleResponseBody {
	s.NamespaceName = &v
	return s
}

func (s *GetArtifactSubscriptionRuleResponseBody) SetOverride(v bool) *GetArtifactSubscriptionRuleResponseBody {
	s.Override = &v
	return s
}

func (s *GetArtifactSubscriptionRuleResponseBody) SetPlatform(v []*string) *GetArtifactSubscriptionRuleResponseBody {
	s.Platform = v
	return s
}

func (s *GetArtifactSubscriptionRuleResponseBody) SetRepoName(v string) *GetArtifactSubscriptionRuleResponseBody {
	s.RepoName = &v
	return s
}

func (s *GetArtifactSubscriptionRuleResponseBody) SetRequestId(v string) *GetArtifactSubscriptionRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetArtifactSubscriptionRuleResponseBody) SetRuleId(v string) *GetArtifactSubscriptionRuleResponseBody {
	s.RuleId = &v
	return s
}

func (s *GetArtifactSubscriptionRuleResponseBody) SetSourceNamespaceName(v string) *GetArtifactSubscriptionRuleResponseBody {
	s.SourceNamespaceName = &v
	return s
}

func (s *GetArtifactSubscriptionRuleResponseBody) SetSourceProvider(v string) *GetArtifactSubscriptionRuleResponseBody {
	s.SourceProvider = &v
	return s
}

func (s *GetArtifactSubscriptionRuleResponseBody) SetSourceRepoName(v string) *GetArtifactSubscriptionRuleResponseBody {
	s.SourceRepoName = &v
	return s
}

func (s *GetArtifactSubscriptionRuleResponseBody) SetTagCount(v int64) *GetArtifactSubscriptionRuleResponseBody {
	s.TagCount = &v
	return s
}

func (s *GetArtifactSubscriptionRuleResponseBody) SetTagRegexp(v string) *GetArtifactSubscriptionRuleResponseBody {
	s.TagRegexp = &v
	return s
}

type GetArtifactSubscriptionRuleResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetArtifactSubscriptionRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetArtifactSubscriptionRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s GetArtifactSubscriptionRuleResponse) GoString() string {
	return s.String()
}

func (s *GetArtifactSubscriptionRuleResponse) SetHeaders(v map[string]*string) *GetArtifactSubscriptionRuleResponse {
	s.Headers = v
	return s
}

func (s *GetArtifactSubscriptionRuleResponse) SetStatusCode(v int32) *GetArtifactSubscriptionRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *GetArtifactSubscriptionRuleResponse) SetBody(v *GetArtifactSubscriptionRuleResponseBody) *GetArtifactSubscriptionRuleResponse {
	s.Body = v
	return s
}

type GetArtifactSubscriptionTaskRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cri-xkx6vujuhay0****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The task ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// crast-40le4es9yh0p****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s GetArtifactSubscriptionTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s GetArtifactSubscriptionTaskRequest) GoString() string {
	return s.String()
}

func (s *GetArtifactSubscriptionTaskRequest) SetInstanceId(v string) *GetArtifactSubscriptionTaskRequest {
	s.InstanceId = &v
	return s
}

func (s *GetArtifactSubscriptionTaskRequest) SetTaskId(v string) *GetArtifactSubscriptionTaskRequest {
	s.TaskId = &v
	return s
}

type GetArtifactSubscriptionTaskResponseBody struct {
	// The artifact type.
	//
	// example:
	//
	// IMAGE
	ArtifactType *string `json:"ArtifactType,omitempty" xml:"ArtifactType,omitempty"`
	// The return value.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The end time of the artifact subscription task.
	//
	// example:
	//
	// 1691979202000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// cri-xkx6vujuhay0****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Indicates whether the API request is successful. Valid values:
	//
	// 	- `true`: The request is successful.
	//
	// 	- `false`: The request fails.
	//
	// example:
	//
	// true
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The return message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The name of the Container Registry namespace.
	//
	// example:
	//
	// test
	NamespaceName *string `json:"NamespaceName,omitempty" xml:"NamespaceName,omitempty"`
	// The name of the Container Registry repository.
	//
	// example:
	//
	// test-repo
	RepoName *string `json:"RepoName,omitempty" xml:"RepoName,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 12589EF7-96E2-4554-AAD7-F7209E88CAD3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The name of the source namespace.
	//
	// example:
	//
	// library
	SourceNamespaceName *string `json:"SourceNamespaceName,omitempty" xml:"SourceNamespaceName,omitempty"`
	// The artifact source.
	//
	// example:
	//
	// DOCKER_HUB
	SourceProvider *string `json:"SourceProvider,omitempty" xml:"SourceProvider,omitempty"`
	// The name of the source repository.
	//
	// example:
	//
	// nginx
	SourceRepoName *string `json:"SourceRepoName,omitempty" xml:"SourceRepoName,omitempty"`
	// The type of the source repository.
	//
	// example:
	//
	// PUBLIC
	SourceRepoType *string `json:"SourceRepoType,omitempty" xml:"SourceRepoType,omitempty"`
	// The start time of the artifact subscription task.
	//
	// example:
	//
	// 1568718468000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The total subscribed tags.
	//
	// example:
	//
	// 1
	TagSubscriptionCount *int64 `json:"TagSubscriptionCount,omitempty" xml:"TagSubscriptionCount,omitempty"`
	// The total number of tags.
	//
	// example:
	//
	// 6
	TagTotalCount *int64 `json:"TagTotalCount,omitempty" xml:"TagTotalCount,omitempty"`
	// The task ID.
	//
	// example:
	//
	// crast-40le4es9yh0p****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The task results.
	//
	// example:
	//
	// SUCCESS
	TaskResult *string `json:"TaskResult,omitempty" xml:"TaskResult,omitempty"`
	// The status of the task.
	//
	// example:
	//
	// RUNNING
	TaskStatus *string `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	// The type of the task. Valid values:
	//
	// example:
	//
	// AUTO
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s GetArtifactSubscriptionTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetArtifactSubscriptionTaskResponseBody) GoString() string {
	return s.String()
}

func (s *GetArtifactSubscriptionTaskResponseBody) SetArtifactType(v string) *GetArtifactSubscriptionTaskResponseBody {
	s.ArtifactType = &v
	return s
}

func (s *GetArtifactSubscriptionTaskResponseBody) SetCode(v string) *GetArtifactSubscriptionTaskResponseBody {
	s.Code = &v
	return s
}

func (s *GetArtifactSubscriptionTaskResponseBody) SetEndTime(v int64) *GetArtifactSubscriptionTaskResponseBody {
	s.EndTime = &v
	return s
}

func (s *GetArtifactSubscriptionTaskResponseBody) SetInstanceId(v string) *GetArtifactSubscriptionTaskResponseBody {
	s.InstanceId = &v
	return s
}

func (s *GetArtifactSubscriptionTaskResponseBody) SetIsSuccess(v bool) *GetArtifactSubscriptionTaskResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *GetArtifactSubscriptionTaskResponseBody) SetMessage(v string) *GetArtifactSubscriptionTaskResponseBody {
	s.Message = &v
	return s
}

func (s *GetArtifactSubscriptionTaskResponseBody) SetNamespaceName(v string) *GetArtifactSubscriptionTaskResponseBody {
	s.NamespaceName = &v
	return s
}

func (s *GetArtifactSubscriptionTaskResponseBody) SetRepoName(v string) *GetArtifactSubscriptionTaskResponseBody {
	s.RepoName = &v
	return s
}

func (s *GetArtifactSubscriptionTaskResponseBody) SetRequestId(v string) *GetArtifactSubscriptionTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetArtifactSubscriptionTaskResponseBody) SetSourceNamespaceName(v string) *GetArtifactSubscriptionTaskResponseBody {
	s.SourceNamespaceName = &v
	return s
}

func (s *GetArtifactSubscriptionTaskResponseBody) SetSourceProvider(v string) *GetArtifactSubscriptionTaskResponseBody {
	s.SourceProvider = &v
	return s
}

func (s *GetArtifactSubscriptionTaskResponseBody) SetSourceRepoName(v string) *GetArtifactSubscriptionTaskResponseBody {
	s.SourceRepoName = &v
	return s
}

func (s *GetArtifactSubscriptionTaskResponseBody) SetSourceRepoType(v string) *GetArtifactSubscriptionTaskResponseBody {
	s.SourceRepoType = &v
	return s
}

func (s *GetArtifactSubscriptionTaskResponseBody) SetStartTime(v int64) *GetArtifactSubscriptionTaskResponseBody {
	s.StartTime = &v
	return s
}

func (s *GetArtifactSubscriptionTaskResponseBody) SetTagSubscriptionCount(v int64) *GetArtifactSubscriptionTaskResponseBody {
	s.TagSubscriptionCount = &v
	return s
}

func (s *GetArtifactSubscriptionTaskResponseBody) SetTagTotalCount(v int64) *GetArtifactSubscriptionTaskResponseBody {
	s.TagTotalCount = &v
	return s
}

func (s *GetArtifactSubscriptionTaskResponseBody) SetTaskId(v string) *GetArtifactSubscriptionTaskResponseBody {
	s.TaskId = &v
	return s
}

func (s *GetArtifactSubscriptionTaskResponseBody) SetTaskResult(v string) *GetArtifactSubscriptionTaskResponseBody {
	s.TaskResult = &v
	return s
}

func (s *GetArtifactSubscriptionTaskResponseBody) SetTaskStatus(v string) *GetArtifactSubscriptionTaskResponseBody {
	s.TaskStatus = &v
	return s
}

func (s *GetArtifactSubscriptionTaskResponseBody) SetTaskType(v string) *GetArtifactSubscriptionTaskResponseBody {
	s.TaskType = &v
	return s
}

type GetArtifactSubscriptionTaskResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetArtifactSubscriptionTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetArtifactSubscriptionTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s GetArtifactSubscriptionTaskResponse) GoString() string {
	return s.String()
}

func (s *GetArtifactSubscriptionTaskResponse) SetHeaders(v map[string]*string) *GetArtifactSubscriptionTaskResponse {
	s.Headers = v
	return s
}

func (s *GetArtifactSubscriptionTaskResponse) SetStatusCode(v int32) *GetArtifactSubscriptionTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *GetArtifactSubscriptionTaskResponse) SetBody(v *GetArtifactSubscriptionTaskResponseBody) *GetArtifactSubscriptionTaskResponse {
	s.Body = v
	return s
}

type GetArtifactSubscriptionTaskResultRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cri-90fxryf9pwf****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The task ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// crast-y64sq01bgad****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s GetArtifactSubscriptionTaskResultRequest) String() string {
	return tea.Prettify(s)
}

func (s GetArtifactSubscriptionTaskResultRequest) GoString() string {
	return s.String()
}

func (s *GetArtifactSubscriptionTaskResultRequest) SetInstanceId(v string) *GetArtifactSubscriptionTaskResultRequest {
	s.InstanceId = &v
	return s
}

func (s *GetArtifactSubscriptionTaskResultRequest) SetPageNo(v int32) *GetArtifactSubscriptionTaskResultRequest {
	s.PageNo = &v
	return s
}

func (s *GetArtifactSubscriptionTaskResultRequest) SetPageSize(v int32) *GetArtifactSubscriptionTaskResultRequest {
	s.PageSize = &v
	return s
}

func (s *GetArtifactSubscriptionTaskResultRequest) SetTaskId(v string) *GetArtifactSubscriptionTaskResultRequest {
	s.TaskId = &v
	return s
}

type GetArtifactSubscriptionTaskResultResponseBody struct {
	// The return value.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// 	- `true`: The request is successful.
	//
	// 	- `false`: The request fails.
	//
	// example:
	//
	// true
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0A8768F6-9B47-5127-A075-9CFB9F79181F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The result of the artifact subscription task.
	TaskResults []*GetArtifactSubscriptionTaskResultResponseBodyTaskResults `json:"TaskResults,omitempty" xml:"TaskResults,omitempty" type:"Repeated"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s GetArtifactSubscriptionTaskResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetArtifactSubscriptionTaskResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetArtifactSubscriptionTaskResultResponseBody) SetCode(v string) *GetArtifactSubscriptionTaskResultResponseBody {
	s.Code = &v
	return s
}

func (s *GetArtifactSubscriptionTaskResultResponseBody) SetIsSuccess(v bool) *GetArtifactSubscriptionTaskResultResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *GetArtifactSubscriptionTaskResultResponseBody) SetPageNo(v int32) *GetArtifactSubscriptionTaskResultResponseBody {
	s.PageNo = &v
	return s
}

func (s *GetArtifactSubscriptionTaskResultResponseBody) SetPageSize(v int32) *GetArtifactSubscriptionTaskResultResponseBody {
	s.PageSize = &v
	return s
}

func (s *GetArtifactSubscriptionTaskResultResponseBody) SetRequestId(v string) *GetArtifactSubscriptionTaskResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetArtifactSubscriptionTaskResultResponseBody) SetTaskResults(v []*GetArtifactSubscriptionTaskResultResponseBodyTaskResults) *GetArtifactSubscriptionTaskResultResponseBody {
	s.TaskResults = v
	return s
}

func (s *GetArtifactSubscriptionTaskResultResponseBody) SetTotalCount(v int32) *GetArtifactSubscriptionTaskResultResponseBody {
	s.TotalCount = &v
	return s
}

type GetArtifactSubscriptionTaskResultResponseBodyTaskResults struct {
	// The end time of the subscription task.
	//
	// example:
	//
	// 1692756630000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// cri-isj2wgaw4z9****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the namespace.
	//
	// example:
	//
	// test-ns
	NamespaceName *string `json:"NamespaceName,omitempty" xml:"NamespaceName,omitempty"`
	// The name of the repository.
	//
	// example:
	//
	// test-reop
	RepoName *string `json:"RepoName,omitempty" xml:"RepoName,omitempty"`
	// The result of the task.
	//
	// example:
	//
	// SUCCESS
	Result *string `json:"Result,omitempty" xml:"Result,omitempty"`
	// The start time of the subscription task.
	//
	// example:
	//
	// 1691719501000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The status of the task.
	//
	// example:
	//
	// COMPLETED
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The image tag.
	//
	// example:
	//
	// v2.0
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	// The task ID.
	//
	// example:
	//
	// crast-wkpfwqozjiq****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s GetArtifactSubscriptionTaskResultResponseBodyTaskResults) String() string {
	return tea.Prettify(s)
}

func (s GetArtifactSubscriptionTaskResultResponseBodyTaskResults) GoString() string {
	return s.String()
}

func (s *GetArtifactSubscriptionTaskResultResponseBodyTaskResults) SetEndTime(v int64) *GetArtifactSubscriptionTaskResultResponseBodyTaskResults {
	s.EndTime = &v
	return s
}

func (s *GetArtifactSubscriptionTaskResultResponseBodyTaskResults) SetInstanceId(v string) *GetArtifactSubscriptionTaskResultResponseBodyTaskResults {
	s.InstanceId = &v
	return s
}

func (s *GetArtifactSubscriptionTaskResultResponseBodyTaskResults) SetNamespaceName(v string) *GetArtifactSubscriptionTaskResultResponseBodyTaskResults {
	s.NamespaceName = &v
	return s
}

func (s *GetArtifactSubscriptionTaskResultResponseBodyTaskResults) SetRepoName(v string) *GetArtifactSubscriptionTaskResultResponseBodyTaskResults {
	s.RepoName = &v
	return s
}

func (s *GetArtifactSubscriptionTaskResultResponseBodyTaskResults) SetResult(v string) *GetArtifactSubscriptionTaskResultResponseBodyTaskResults {
	s.Result = &v
	return s
}

func (s *GetArtifactSubscriptionTaskResultResponseBodyTaskResults) SetStartTime(v int64) *GetArtifactSubscriptionTaskResultResponseBodyTaskResults {
	s.StartTime = &v
	return s
}

func (s *GetArtifactSubscriptionTaskResultResponseBodyTaskResults) SetStatus(v string) *GetArtifactSubscriptionTaskResultResponseBodyTaskResults {
	s.Status = &v
	return s
}

func (s *GetArtifactSubscriptionTaskResultResponseBodyTaskResults) SetTag(v string) *GetArtifactSubscriptionTaskResultResponseBodyTaskResults {
	s.Tag = &v
	return s
}

func (s *GetArtifactSubscriptionTaskResultResponseBodyTaskResults) SetTaskId(v string) *GetArtifactSubscriptionTaskResultResponseBodyTaskResults {
	s.TaskId = &v
	return s
}

type GetArtifactSubscriptionTaskResultResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetArtifactSubscriptionTaskResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetArtifactSubscriptionTaskResultResponse) String() string {
	return tea.Prettify(s)
}

func (s GetArtifactSubscriptionTaskResultResponse) GoString() string {
	return s.String()
}

func (s *GetArtifactSubscriptionTaskResultResponse) SetHeaders(v map[string]*string) *GetArtifactSubscriptionTaskResultResponse {
	s.Headers = v
	return s
}

func (s *GetArtifactSubscriptionTaskResultResponse) SetStatusCode(v int32) *GetArtifactSubscriptionTaskResultResponse {
	s.StatusCode = &v
	return s
}

func (s *GetArtifactSubscriptionTaskResultResponse) SetBody(v *GetArtifactSubscriptionTaskResultResponseBody) *GetArtifactSubscriptionTaskResultResponse {
	s.Body = v
	return s
}

type GetAuthorizationTokenRequest struct {
	// The ID of the request.
	//
	// This parameter is required.
	//
	// example:
	//
	// cri-kmsiwlxxdcvaduwb
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetAuthorizationTokenRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAuthorizationTokenRequest) GoString() string {
	return s.String()
}

func (s *GetAuthorizationTokenRequest) SetInstanceId(v string) *GetAuthorizationTokenRequest {
	s.InstanceId = &v
	return s
}

type GetAuthorizationTokenResponseBody struct {
	// The temporary password returned after you call this API operation is a Security Token Service (STS) token whose validity period is 1 hour. Take note of the following items when you log on to Container Registry instances by using an STS token:
	//
	// 	- If the STS token belongs to an Alibaba Cloud account, you can use the STS token to log on to all Container Registry instances that belong to the Alibaba Cloud account.
	//
	// 	- If the STS token belongs to a Resource Access Management (RAM) user, you can use the STS token to log on to all Container Registry instances that belong to the RAM user.
	//
	// 	- You can use an STS token to access only Container Registry instances to which the STS token is scoped.
	//
	// example:
	//
	// shaunadadakks:uuczxnjcyeyhdjadkkajsjdjadhyucb
	AuthorizationToken *string `json:"AuthorizationToken,omitempty" xml:"AuthorizationToken,omitempty"`
	// Indicates whether the API call is successful.
	//
	// 	- `true`: successful
	//
	// 	- `false`: failed
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The return value.
	//
	// example:
	//
	// 1571242083000
	ExpireTime *int64 `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The username that is used to log on to the Container Registry instance.
	//
	// example:
	//
	// true
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The timestamp when the temporary password expires. Unit: milliseconds.
	//
	// example:
	//
	// E069EB86-E6AD-4A98-ADDE-0E993390239A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The password that is used to log on to the Container Registry instance.
	//
	// example:
	//
	// temp_user_cr
	TempUsername *string `json:"TempUsername,omitempty" xml:"TempUsername,omitempty"`
}

func (s GetAuthorizationTokenResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAuthorizationTokenResponseBody) GoString() string {
	return s.String()
}

func (s *GetAuthorizationTokenResponseBody) SetAuthorizationToken(v string) *GetAuthorizationTokenResponseBody {
	s.AuthorizationToken = &v
	return s
}

func (s *GetAuthorizationTokenResponseBody) SetCode(v string) *GetAuthorizationTokenResponseBody {
	s.Code = &v
	return s
}

func (s *GetAuthorizationTokenResponseBody) SetExpireTime(v int64) *GetAuthorizationTokenResponseBody {
	s.ExpireTime = &v
	return s
}

func (s *GetAuthorizationTokenResponseBody) SetIsSuccess(v bool) *GetAuthorizationTokenResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *GetAuthorizationTokenResponseBody) SetRequestId(v string) *GetAuthorizationTokenResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAuthorizationTokenResponseBody) SetTempUsername(v string) *GetAuthorizationTokenResponseBody {
	s.TempUsername = &v
	return s
}

type GetAuthorizationTokenResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAuthorizationTokenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAuthorizationTokenResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAuthorizationTokenResponse) GoString() string {
	return s.String()
}

func (s *GetAuthorizationTokenResponse) SetHeaders(v map[string]*string) *GetAuthorizationTokenResponse {
	s.Headers = v
	return s
}

func (s *GetAuthorizationTokenResponse) SetStatusCode(v int32) *GetAuthorizationTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAuthorizationTokenResponse) SetBody(v *GetAuthorizationTokenResponseBody) *GetAuthorizationTokenResponse {
	s.Body = v
	return s
}

type GetChainRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// chi-0ops0gsmw5x2****
	ChainId *string `json:"ChainId,omitempty" xml:"ChainId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cri-4cdrlqmhn4gm****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetChainRequest) String() string {
	return tea.Prettify(s)
}

func (s GetChainRequest) GoString() string {
	return s.String()
}

func (s *GetChainRequest) SetChainId(v string) *GetChainRequest {
	s.ChainId = &v
	return s
}

func (s *GetChainRequest) SetInstanceId(v string) *GetChainRequest {
	s.InstanceId = &v
	return s
}

type GetChainResponseBody struct {
	ChainConfig *GetChainResponseBodyChainConfig `json:"ChainConfig,omitempty" xml:"ChainConfig,omitempty" type:"Struct"`
	// example:
	//
	// chi-0ops0gsmw5x2****
	ChainId *string `json:"ChainId,omitempty" xml:"ChainId,omitempty"`
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 1638255427000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// cri-4cdrlqmhn4gm****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// true
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// example:
	//
	// 1638259914000
	ModifiedTime *int64 `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// C87993B5-7D61-5CAC-8D64-1AC732DD69FF
	RequestId    *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ScopeExclude []*string `json:"ScopeExclude,omitempty" xml:"ScopeExclude,omitempty" type:"Repeated"`
	// example:
	//
	// crr-nyrh2oko32xb****
	ScopeId *string `json:"ScopeId,omitempty" xml:"ScopeId,omitempty"`
	// example:
	//
	// REPOSITORY
	ScopeType *string `json:"ScopeType,omitempty" xml:"ScopeType,omitempty"`
}

func (s GetChainResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetChainResponseBody) GoString() string {
	return s.String()
}

func (s *GetChainResponseBody) SetChainConfig(v *GetChainResponseBodyChainConfig) *GetChainResponseBody {
	s.ChainConfig = v
	return s
}

func (s *GetChainResponseBody) SetChainId(v string) *GetChainResponseBody {
	s.ChainId = &v
	return s
}

func (s *GetChainResponseBody) SetCode(v string) *GetChainResponseBody {
	s.Code = &v
	return s
}

func (s *GetChainResponseBody) SetCreateTime(v int64) *GetChainResponseBody {
	s.CreateTime = &v
	return s
}

func (s *GetChainResponseBody) SetDescription(v string) *GetChainResponseBody {
	s.Description = &v
	return s
}

func (s *GetChainResponseBody) SetInstanceId(v string) *GetChainResponseBody {
	s.InstanceId = &v
	return s
}

func (s *GetChainResponseBody) SetIsSuccess(v bool) *GetChainResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *GetChainResponseBody) SetModifiedTime(v int64) *GetChainResponseBody {
	s.ModifiedTime = &v
	return s
}

func (s *GetChainResponseBody) SetName(v string) *GetChainResponseBody {
	s.Name = &v
	return s
}

func (s *GetChainResponseBody) SetRequestId(v string) *GetChainResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetChainResponseBody) SetScopeExclude(v []*string) *GetChainResponseBody {
	s.ScopeExclude = v
	return s
}

func (s *GetChainResponseBody) SetScopeId(v string) *GetChainResponseBody {
	s.ScopeId = &v
	return s
}

func (s *GetChainResponseBody) SetScopeType(v string) *GetChainResponseBody {
	s.ScopeType = &v
	return s
}

type GetChainResponseBodyChainConfig struct {
	// example:
	//
	// cci-lz3ycgo69ukt****
	ChainConfigId *string `json:"ChainConfigId,omitempty" xml:"ChainConfigId,omitempty"`
	// example:
	//
	// true
	IsActive *bool                                     `json:"IsActive,omitempty" xml:"IsActive,omitempty"`
	Nodes    []*GetChainResponseBodyChainConfigNodes   `json:"Nodes,omitempty" xml:"Nodes,omitempty" type:"Repeated"`
	Routers  []*GetChainResponseBodyChainConfigRouters `json:"Routers,omitempty" xml:"Routers,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s GetChainResponseBodyChainConfig) String() string {
	return tea.Prettify(s)
}

func (s GetChainResponseBodyChainConfig) GoString() string {
	return s.String()
}

func (s *GetChainResponseBodyChainConfig) SetChainConfigId(v string) *GetChainResponseBodyChainConfig {
	s.ChainConfigId = &v
	return s
}

func (s *GetChainResponseBodyChainConfig) SetIsActive(v bool) *GetChainResponseBodyChainConfig {
	s.IsActive = &v
	return s
}

func (s *GetChainResponseBodyChainConfig) SetNodes(v []*GetChainResponseBodyChainConfigNodes) *GetChainResponseBodyChainConfig {
	s.Nodes = v
	return s
}

func (s *GetChainResponseBodyChainConfig) SetRouters(v []*GetChainResponseBodyChainConfigRouters) *GetChainResponseBodyChainConfig {
	s.Routers = v
	return s
}

func (s *GetChainResponseBodyChainConfig) SetVersion(v string) *GetChainResponseBodyChainConfig {
	s.Version = &v
	return s
}

type GetChainResponseBodyChainConfigNodes struct {
	// example:
	//
	// true
	Enable     *bool                                           `json:"Enable,omitempty" xml:"Enable,omitempty"`
	NodeConfig *GetChainResponseBodyChainConfigNodesNodeConfig `json:"NodeConfig,omitempty" xml:"NodeConfig,omitempty" type:"Struct"`
	// example:
	//
	// VULNERABILITY_SCANNING
	NodeName *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
}

func (s GetChainResponseBodyChainConfigNodes) String() string {
	return tea.Prettify(s)
}

func (s GetChainResponseBodyChainConfigNodes) GoString() string {
	return s.String()
}

func (s *GetChainResponseBodyChainConfigNodes) SetEnable(v bool) *GetChainResponseBodyChainConfigNodes {
	s.Enable = &v
	return s
}

func (s *GetChainResponseBodyChainConfigNodes) SetNodeConfig(v *GetChainResponseBodyChainConfigNodesNodeConfig) *GetChainResponseBodyChainConfigNodes {
	s.NodeConfig = v
	return s
}

func (s *GetChainResponseBodyChainConfigNodes) SetNodeName(v string) *GetChainResponseBodyChainConfigNodes {
	s.NodeName = &v
	return s
}

type GetChainResponseBodyChainConfigNodesNodeConfig struct {
	DenyPolicy *GetChainResponseBodyChainConfigNodesNodeConfigDenyPolicy `json:"DenyPolicy,omitempty" xml:"DenyPolicy,omitempty" type:"Struct"`
	// example:
	//
	// 3
	Retry *int32 `json:"Retry,omitempty" xml:"Retry,omitempty"`
	// example:
	//
	// ACR_SCAN_SERVICE
	ScanEngine *string `json:"ScanEngine,omitempty" xml:"ScanEngine,omitempty"`
	Timeout    *int64  `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
}

func (s GetChainResponseBodyChainConfigNodesNodeConfig) String() string {
	return tea.Prettify(s)
}

func (s GetChainResponseBodyChainConfigNodesNodeConfig) GoString() string {
	return s.String()
}

func (s *GetChainResponseBodyChainConfigNodesNodeConfig) SetDenyPolicy(v *GetChainResponseBodyChainConfigNodesNodeConfigDenyPolicy) *GetChainResponseBodyChainConfigNodesNodeConfig {
	s.DenyPolicy = v
	return s
}

func (s *GetChainResponseBodyChainConfigNodesNodeConfig) SetRetry(v int32) *GetChainResponseBodyChainConfigNodesNodeConfig {
	s.Retry = &v
	return s
}

func (s *GetChainResponseBodyChainConfigNodesNodeConfig) SetScanEngine(v string) *GetChainResponseBodyChainConfigNodesNodeConfig {
	s.ScanEngine = &v
	return s
}

func (s *GetChainResponseBodyChainConfigNodesNodeConfig) SetTimeout(v int64) *GetChainResponseBodyChainConfigNodesNodeConfig {
	s.Timeout = &v
	return s
}

type GetChainResponseBodyChainConfigNodesNodeConfigDenyPolicy struct {
	// example:
	//
	// BLOCK
	Action *string `json:"Action,omitempty" xml:"Action,omitempty"`
	// example:
	//
	// identification,hc_image_exploit
	BaselineList *string `json:"BaselineList,omitempty" xml:"BaselineList,omitempty"`
	// example:
	//
	// 10
	IssueCount *string `json:"IssueCount,omitempty" xml:"IssueCount,omitempty"`
	// example:
	//
	// HIGH
	IssueLevel *string `json:"IssueLevel,omitempty" xml:"IssueLevel,omitempty"`
	// example:
	//
	// CVE-2020-8286,CVE-2020-8285
	IssueList *string `json:"IssueList,omitempty" xml:"IssueList,omitempty"`
	// example:
	//
	// AND
	Logic *string `json:"Logic,omitempty" xml:"Logic,omitempty"`
	// example:
	//
	// mutate_cockhorse,abnormal_program
	MaliciousList *string `json:"MaliciousList,omitempty" xml:"MaliciousList,omitempty"`
}

func (s GetChainResponseBodyChainConfigNodesNodeConfigDenyPolicy) String() string {
	return tea.Prettify(s)
}

func (s GetChainResponseBodyChainConfigNodesNodeConfigDenyPolicy) GoString() string {
	return s.String()
}

func (s *GetChainResponseBodyChainConfigNodesNodeConfigDenyPolicy) SetAction(v string) *GetChainResponseBodyChainConfigNodesNodeConfigDenyPolicy {
	s.Action = &v
	return s
}

func (s *GetChainResponseBodyChainConfigNodesNodeConfigDenyPolicy) SetBaselineList(v string) *GetChainResponseBodyChainConfigNodesNodeConfigDenyPolicy {
	s.BaselineList = &v
	return s
}

func (s *GetChainResponseBodyChainConfigNodesNodeConfigDenyPolicy) SetIssueCount(v string) *GetChainResponseBodyChainConfigNodesNodeConfigDenyPolicy {
	s.IssueCount = &v
	return s
}

func (s *GetChainResponseBodyChainConfigNodesNodeConfigDenyPolicy) SetIssueLevel(v string) *GetChainResponseBodyChainConfigNodesNodeConfigDenyPolicy {
	s.IssueLevel = &v
	return s
}

func (s *GetChainResponseBodyChainConfigNodesNodeConfigDenyPolicy) SetIssueList(v string) *GetChainResponseBodyChainConfigNodesNodeConfigDenyPolicy {
	s.IssueList = &v
	return s
}

func (s *GetChainResponseBodyChainConfigNodesNodeConfigDenyPolicy) SetLogic(v string) *GetChainResponseBodyChainConfigNodesNodeConfigDenyPolicy {
	s.Logic = &v
	return s
}

func (s *GetChainResponseBodyChainConfigNodesNodeConfigDenyPolicy) SetMaliciousList(v string) *GetChainResponseBodyChainConfigNodesNodeConfigDenyPolicy {
	s.MaliciousList = &v
	return s
}

type GetChainResponseBodyChainConfigRouters struct {
	From *GetChainResponseBodyChainConfigRoutersFrom `json:"From,omitempty" xml:"From,omitempty" type:"Struct"`
	To   *GetChainResponseBodyChainConfigRoutersTo   `json:"To,omitempty" xml:"To,omitempty" type:"Struct"`
}

func (s GetChainResponseBodyChainConfigRouters) String() string {
	return tea.Prettify(s)
}

func (s GetChainResponseBodyChainConfigRouters) GoString() string {
	return s.String()
}

func (s *GetChainResponseBodyChainConfigRouters) SetFrom(v *GetChainResponseBodyChainConfigRoutersFrom) *GetChainResponseBodyChainConfigRouters {
	s.From = v
	return s
}

func (s *GetChainResponseBodyChainConfigRouters) SetTo(v *GetChainResponseBodyChainConfigRoutersTo) *GetChainResponseBodyChainConfigRouters {
	s.To = v
	return s
}

type GetChainResponseBodyChainConfigRoutersFrom struct {
	// example:
	//
	// DOCKER_IMAGE_BUILD
	NodeName *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
}

func (s GetChainResponseBodyChainConfigRoutersFrom) String() string {
	return tea.Prettify(s)
}

func (s GetChainResponseBodyChainConfigRoutersFrom) GoString() string {
	return s.String()
}

func (s *GetChainResponseBodyChainConfigRoutersFrom) SetNodeName(v string) *GetChainResponseBodyChainConfigRoutersFrom {
	s.NodeName = &v
	return s
}

type GetChainResponseBodyChainConfigRoutersTo struct {
	// example:
	//
	// DOCKER_IMAGE_PUSH
	NodeName *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
}

func (s GetChainResponseBodyChainConfigRoutersTo) String() string {
	return tea.Prettify(s)
}

func (s GetChainResponseBodyChainConfigRoutersTo) GoString() string {
	return s.String()
}

func (s *GetChainResponseBodyChainConfigRoutersTo) SetNodeName(v string) *GetChainResponseBodyChainConfigRoutersTo {
	s.NodeName = &v
	return s
}

type GetChainResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetChainResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetChainResponse) String() string {
	return tea.Prettify(s)
}

func (s GetChainResponse) GoString() string {
	return s.String()
}

func (s *GetChainResponse) SetHeaders(v map[string]*string) *GetChainResponse {
	s.Headers = v
	return s
}

func (s *GetChainResponse) SetStatusCode(v int32) *GetChainResponse {
	s.StatusCode = &v
	return s
}

func (s *GetChainResponse) SetBody(v *GetChainResponseBody) *GetChainResponse {
	s.Body = v
	return s
}

type GetChartNamespaceRequest struct {
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cri-xkx6vujuhay0****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the namespace.
	//
	// This parameter is required.
	//
	// example:
	//
	// ns1
	NamespaceName *string `json:"NamespaceName,omitempty" xml:"NamespaceName,omitempty"`
}

func (s GetChartNamespaceRequest) String() string {
	return tea.Prettify(s)
}

func (s GetChartNamespaceRequest) GoString() string {
	return s.String()
}

func (s *GetChartNamespaceRequest) SetInstanceId(v string) *GetChartNamespaceRequest {
	s.InstanceId = &v
	return s
}

func (s *GetChartNamespaceRequest) SetNamespaceName(v string) *GetChartNamespaceRequest {
	s.NamespaceName = &v
	return s
}

type GetChartNamespaceResponseBody struct {
	// Indicates whether a repository was automatically created in the namespace. Valid values:
	//
	// 	- `true`: A repository was automatically created in the namespace.
	//
	// 	- `false`: No repository was automatically created in the namespace.
	//
	// example:
	//
	// true
	AutoCreateRepo *bool `json:"AutoCreateRepo,omitempty" xml:"AutoCreateRepo,omitempty"`
	// The return value.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The default repository type. Valid values:
	//
	// 	- `PUBLIC`: a public repository.
	//
	// 	- `PRIVATE`: a private repository.
	//
	// example:
	//
	// PRIVATE
	DefaultRepoType *string `json:"DefaultRepoType,omitempty" xml:"DefaultRepoType,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// cri-xkx6vujuhay0****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// 	- `true`: The request is successful.
	//
	// 	- `false`: The request fails.
	//
	// example:
	//
	// true
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The ID of the namespace.
	//
	// example:
	//
	// crcn-43dhbjbyt2xl****
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	// The name of the namespace.
	//
	// example:
	//
	// ns1
	NamespaceName *string `json:"NamespaceName,omitempty" xml:"NamespaceName,omitempty"`
	// The status of the namespace. Valid values:
	//
	// 	- `NORMAL`: The namespace is normal.
	//
	// 	- `DELETING`: The namespace is being deleted.
	//
	// example:
	//
	// NORMAL
	NamespaceStatus *string `json:"NamespaceStatus,omitempty" xml:"NamespaceStatus,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// CD71CF13-93AA-4805-848B-69B2DD543A9A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the resource group to which the instance belongs.
	//
	// example:
	//
	// rg-acfmv36i4is****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s GetChartNamespaceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetChartNamespaceResponseBody) GoString() string {
	return s.String()
}

func (s *GetChartNamespaceResponseBody) SetAutoCreateRepo(v bool) *GetChartNamespaceResponseBody {
	s.AutoCreateRepo = &v
	return s
}

func (s *GetChartNamespaceResponseBody) SetCode(v string) *GetChartNamespaceResponseBody {
	s.Code = &v
	return s
}

func (s *GetChartNamespaceResponseBody) SetDefaultRepoType(v string) *GetChartNamespaceResponseBody {
	s.DefaultRepoType = &v
	return s
}

func (s *GetChartNamespaceResponseBody) SetInstanceId(v string) *GetChartNamespaceResponseBody {
	s.InstanceId = &v
	return s
}

func (s *GetChartNamespaceResponseBody) SetIsSuccess(v bool) *GetChartNamespaceResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *GetChartNamespaceResponseBody) SetNamespaceId(v string) *GetChartNamespaceResponseBody {
	s.NamespaceId = &v
	return s
}

func (s *GetChartNamespaceResponseBody) SetNamespaceName(v string) *GetChartNamespaceResponseBody {
	s.NamespaceName = &v
	return s
}

func (s *GetChartNamespaceResponseBody) SetNamespaceStatus(v string) *GetChartNamespaceResponseBody {
	s.NamespaceStatus = &v
	return s
}

func (s *GetChartNamespaceResponseBody) SetRequestId(v string) *GetChartNamespaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetChartNamespaceResponseBody) SetResourceGroupId(v string) *GetChartNamespaceResponseBody {
	s.ResourceGroupId = &v
	return s
}

type GetChartNamespaceResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetChartNamespaceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetChartNamespaceResponse) String() string {
	return tea.Prettify(s)
}

func (s GetChartNamespaceResponse) GoString() string {
	return s.String()
}

func (s *GetChartNamespaceResponse) SetHeaders(v map[string]*string) *GetChartNamespaceResponse {
	s.Headers = v
	return s
}

func (s *GetChartNamespaceResponse) SetStatusCode(v int32) *GetChartNamespaceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetChartNamespaceResponse) SetBody(v *GetChartNamespaceResponseBody) *GetChartNamespaceResponse {
	s.Body = v
	return s
}

type GetChartRepositoryRequest struct {
	// The ID of the Container Registry instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cri-kmsiwlxxdcva****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the repository.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	RepoName *string `json:"RepoName,omitempty" xml:"RepoName,omitempty"`
	// The name of the namespace.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	RepoNamespaceName *string `json:"RepoNamespaceName,omitempty" xml:"RepoNamespaceName,omitempty"`
}

func (s GetChartRepositoryRequest) String() string {
	return tea.Prettify(s)
}

func (s GetChartRepositoryRequest) GoString() string {
	return s.String()
}

func (s *GetChartRepositoryRequest) SetInstanceId(v string) *GetChartRepositoryRequest {
	s.InstanceId = &v
	return s
}

func (s *GetChartRepositoryRequest) SetRepoName(v string) *GetChartRepositoryRequest {
	s.RepoName = &v
	return s
}

func (s *GetChartRepositoryRequest) SetRepoNamespaceName(v string) *GetChartRepositoryRequest {
	s.RepoNamespaceName = &v
	return s
}

type GetChartRepositoryResponseBody struct {
	// The return value.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The time when the chart repository was created.
	//
	// example:
	//
	// 1563767620000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// cri-kmsiwlxxdcva****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// 	- `true`: The request is successful.
	//
	// 	- `false`: The request fails.
	//
	// example:
	//
	// true
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The time when the chart repository was last modified.
	//
	// example:
	//
	// 1563767700000
	ModifiedTime *int64 `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// The ID of the chart repository.
	//
	// example:
	//
	// crcr-c7letfwev5oq****
	RepoId *string `json:"RepoId,omitempty" xml:"RepoId,omitempty"`
	// The name of the chart repository.
	//
	// example:
	//
	// test
	RepoName *string `json:"RepoName,omitempty" xml:"RepoName,omitempty"`
	// The name of the namespace to which the chart repository belongs.
	//
	// example:
	//
	// test
	RepoNamespaceName *string `json:"RepoNamespaceName,omitempty" xml:"RepoNamespaceName,omitempty"`
	// The status of the chart repository. Valid values:
	//
	// 	- `NORMAL`: The repository is normal.
	//
	// 	- `DELETING`: The repository is being deleted.
	//
	// example:
	//
	// NORMAL
	RepoStatus *string `json:"RepoStatus,omitempty" xml:"RepoStatus,omitempty"`
	// The type of the chart repository. Valid values:
	//
	// 	- `PUBLIC`: a public repository
	//
	// 	- `PRIVATE`: a private repository
	//
	// example:
	//
	// PUBLIC
	RepoType *string `json:"RepoType,omitempty" xml:"RepoType,omitempty"`
	// The request ID.
	//
	// example:
	//
	// A3F6AB56-DEF4-4FF5-8DE4-680362C0E21F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfmv36i4is****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The summary about the chart repository.
	//
	// example:
	//
	// test
	Summary *string `json:"Summary,omitempty" xml:"Summary,omitempty"`
}

func (s GetChartRepositoryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetChartRepositoryResponseBody) GoString() string {
	return s.String()
}

func (s *GetChartRepositoryResponseBody) SetCode(v string) *GetChartRepositoryResponseBody {
	s.Code = &v
	return s
}

func (s *GetChartRepositoryResponseBody) SetCreateTime(v int64) *GetChartRepositoryResponseBody {
	s.CreateTime = &v
	return s
}

func (s *GetChartRepositoryResponseBody) SetInstanceId(v string) *GetChartRepositoryResponseBody {
	s.InstanceId = &v
	return s
}

func (s *GetChartRepositoryResponseBody) SetIsSuccess(v bool) *GetChartRepositoryResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *GetChartRepositoryResponseBody) SetModifiedTime(v int64) *GetChartRepositoryResponseBody {
	s.ModifiedTime = &v
	return s
}

func (s *GetChartRepositoryResponseBody) SetRepoId(v string) *GetChartRepositoryResponseBody {
	s.RepoId = &v
	return s
}

func (s *GetChartRepositoryResponseBody) SetRepoName(v string) *GetChartRepositoryResponseBody {
	s.RepoName = &v
	return s
}

func (s *GetChartRepositoryResponseBody) SetRepoNamespaceName(v string) *GetChartRepositoryResponseBody {
	s.RepoNamespaceName = &v
	return s
}

func (s *GetChartRepositoryResponseBody) SetRepoStatus(v string) *GetChartRepositoryResponseBody {
	s.RepoStatus = &v
	return s
}

func (s *GetChartRepositoryResponseBody) SetRepoType(v string) *GetChartRepositoryResponseBody {
	s.RepoType = &v
	return s
}

func (s *GetChartRepositoryResponseBody) SetRequestId(v string) *GetChartRepositoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetChartRepositoryResponseBody) SetResourceGroupId(v string) *GetChartRepositoryResponseBody {
	s.ResourceGroupId = &v
	return s
}

func (s *GetChartRepositoryResponseBody) SetSummary(v string) *GetChartRepositoryResponseBody {
	s.Summary = &v
	return s
}

type GetChartRepositoryResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetChartRepositoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetChartRepositoryResponse) String() string {
	return tea.Prettify(s)
}

func (s GetChartRepositoryResponse) GoString() string {
	return s.String()
}

func (s *GetChartRepositoryResponse) SetHeaders(v map[string]*string) *GetChartRepositoryResponse {
	s.Headers = v
	return s
}

func (s *GetChartRepositoryResponse) SetStatusCode(v int32) *GetChartRepositoryResponse {
	s.StatusCode = &v
	return s
}

func (s *GetChartRepositoryResponse) SetBody(v *GetChartRepositoryResponseBody) *GetChartRepositoryResponse {
	s.Body = v
	return s
}

type GetInstanceRequest struct {
	// The ID of the Container Registry instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cri-xkx6vujuhay0****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceRequest) GoString() string {
	return s.String()
}

func (s *GetInstanceRequest) SetInstanceId(v string) *GetInstanceRequest {
	s.InstanceId = &v
	return s
}

type GetInstanceResponseBody struct {
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The time when the instance was created.
	//
	// example:
	//
	// 1571926439000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The ID of the Container Registry instance.
	//
	// example:
	//
	// cri-xkx6vujuhay0****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The issue occurs on the instance.
	//
	// example:
	//
	// The issue occurs on the instance. Valid values: OSS_TOO_MANY_BUCKETS: The number of Object Storage Service (OSS) buckets exceeds the upper limit. OSS_BUCKET_ALREADY_EXISTS: An OSS bucket that has the duplicate name already exists. OSS_SERVICE_ROLE_UNAUTHORIZED: The OSS service-linked role is not granted permissions. USER_NOT_REGISTERED_BY_REAL_NAME: The Alibaba Cloud account has not passed a real name verification.
	InstanceIssue *string `json:"InstanceIssue,omitempty" xml:"InstanceIssue,omitempty"`
	// The name of the instance.
	//
	// example:
	//
	// shanghai-instance1
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The edition of the instance. Valid values: Enterprise_Basic: Basic Edition instances Enterprise_Standard: Standard Edition instances Enterprise_Advanced: Advanced Edition instances
	//
	// example:
	//
	// Enterprise_Basic
	InstanceSpecification *string `json:"InstanceSpecification,omitempty" xml:"InstanceSpecification,omitempty"`
	// The status of the instance. Valid values:
	//
	// 	- `PENDING`: The instance is being initialized.
	//
	// 	- `INIT_ERROR`: The instance failed to be initialized.
	//
	// 	- `STARTING`: The instance is being started.
	//
	// 	- `RUNNING`: The instance is running.
	//
	// 	- `STOPPING`: The instance is being stopped.
	//
	// 	- `STOPPED`: The instance is stopped.
	//
	// 	- `DELETING`: The instance is being deleted.
	//
	// 	- `DELETED`: The instance is deleted.
	//
	// example:
	//
	// RUNNING
	InstanceStatus *string `json:"InstanceStatus,omitempty" xml:"InstanceStatus,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// 	- `true`: The request is successful.
	//
	// 	- `false`: The request fails.
	//
	// example:
	//
	// true
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The time when the instance was last modified.
	//
	// example:
	//
	// 1571926560000
	ModifiedTime *int64 `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 6EF34B18-4228-470C-860C-D28597CF010E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the resource group to which the instance belongs.
	//
	// example:
	//
	// rg-acfmv36i4isx****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The tags of the instance.
	Tags []*GetInstanceResponseBodyTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s GetInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *GetInstanceResponseBody) SetCode(v string) *GetInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *GetInstanceResponseBody) SetCreateTime(v int64) *GetInstanceResponseBody {
	s.CreateTime = &v
	return s
}

func (s *GetInstanceResponseBody) SetInstanceId(v string) *GetInstanceResponseBody {
	s.InstanceId = &v
	return s
}

func (s *GetInstanceResponseBody) SetInstanceIssue(v string) *GetInstanceResponseBody {
	s.InstanceIssue = &v
	return s
}

func (s *GetInstanceResponseBody) SetInstanceName(v string) *GetInstanceResponseBody {
	s.InstanceName = &v
	return s
}

func (s *GetInstanceResponseBody) SetInstanceSpecification(v string) *GetInstanceResponseBody {
	s.InstanceSpecification = &v
	return s
}

func (s *GetInstanceResponseBody) SetInstanceStatus(v string) *GetInstanceResponseBody {
	s.InstanceStatus = &v
	return s
}

func (s *GetInstanceResponseBody) SetIsSuccess(v bool) *GetInstanceResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *GetInstanceResponseBody) SetModifiedTime(v int64) *GetInstanceResponseBody {
	s.ModifiedTime = &v
	return s
}

func (s *GetInstanceResponseBody) SetRequestId(v string) *GetInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetInstanceResponseBody) SetResourceGroupId(v string) *GetInstanceResponseBody {
	s.ResourceGroupId = &v
	return s
}

func (s *GetInstanceResponseBody) SetTags(v []*GetInstanceResponseBodyTags) *GetInstanceResponseBody {
	s.Tags = v
	return s
}

type GetInstanceResponseBodyTags struct {
	// The tag key.
	//
	// example:
	//
	// test_key
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The tag value.
	//
	// example:
	//
	// test_value
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s GetInstanceResponseBodyTags) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceResponseBodyTags) GoString() string {
	return s.String()
}

func (s *GetInstanceResponseBodyTags) SetTagKey(v string) *GetInstanceResponseBodyTags {
	s.TagKey = &v
	return s
}

func (s *GetInstanceResponseBodyTags) SetTagValue(v string) *GetInstanceResponseBodyTags {
	s.TagValue = &v
	return s
}

type GetInstanceResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceResponse) GoString() string {
	return s.String()
}

func (s *GetInstanceResponse) SetHeaders(v map[string]*string) *GetInstanceResponse {
	s.Headers = v
	return s
}

func (s *GetInstanceResponse) SetStatusCode(v int32) *GetInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetInstanceResponse) SetBody(v *GetInstanceResponseBody) *GetInstanceResponse {
	s.Body = v
	return s
}

type GetInstanceCountResponseBody struct {
	// Return value
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Number of instances
	//
	// example:
	//
	// 5
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// Indicates whether the API call was successful. Values:
	//
	// - `true`: The API call was successful.
	//
	// - `false`: The API call failed.
	//
	// example:
	//
	// true
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// Request ID
	//
	// example:
	//
	// BC648259-91A7-4502-BED3-EDF64361FA83
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetInstanceCountResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceCountResponseBody) GoString() string {
	return s.String()
}

func (s *GetInstanceCountResponseBody) SetCode(v string) *GetInstanceCountResponseBody {
	s.Code = &v
	return s
}

func (s *GetInstanceCountResponseBody) SetCount(v int32) *GetInstanceCountResponseBody {
	s.Count = &v
	return s
}

func (s *GetInstanceCountResponseBody) SetIsSuccess(v bool) *GetInstanceCountResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *GetInstanceCountResponseBody) SetRequestId(v string) *GetInstanceCountResponseBody {
	s.RequestId = &v
	return s
}

type GetInstanceCountResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetInstanceCountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetInstanceCountResponse) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceCountResponse) GoString() string {
	return s.String()
}

func (s *GetInstanceCountResponse) SetHeaders(v map[string]*string) *GetInstanceCountResponse {
	s.Headers = v
	return s
}

func (s *GetInstanceCountResponse) SetStatusCode(v int32) *GetInstanceCountResponse {
	s.StatusCode = &v
	return s
}

func (s *GetInstanceCountResponse) SetBody(v *GetInstanceCountResponseBody) *GetInstanceCountResponse {
	s.Body = v
	return s
}

type GetInstanceEndpointRequest struct {
	// The type of the endpoint. Set the value to Internet.
	//
	// This parameter is required.
	//
	// example:
	//
	// internet
	EndpointType *string `json:"EndpointType,omitempty" xml:"EndpointType,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cri-xkx6vujuhay0****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the module that you want to access. Valid values:
	//
	// 	- `Registry`: the image repository.
	//
	// 	- `Chart`: a Helm chart.
	//
	// example:
	//
	// Registry
	ModuleName *string `json:"ModuleName,omitempty" xml:"ModuleName,omitempty"`
}

func (s GetInstanceEndpointRequest) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceEndpointRequest) GoString() string {
	return s.String()
}

func (s *GetInstanceEndpointRequest) SetEndpointType(v string) *GetInstanceEndpointRequest {
	s.EndpointType = &v
	return s
}

func (s *GetInstanceEndpointRequest) SetInstanceId(v string) *GetInstanceEndpointRequest {
	s.InstanceId = &v
	return s
}

func (s *GetInstanceEndpointRequest) SetModuleName(v string) *GetInstanceEndpointRequest {
	s.ModuleName = &v
	return s
}

type GetInstanceEndpointResponseBody struct {
	// Indicates whether the access control list (ACL) feature is enabled.
	//
	// example:
	//
	// true
	AclEnable *bool `json:"AclEnable,omitempty" xml:"AclEnable,omitempty"`
	// The ACLs.
	AclEntries []*GetInstanceEndpointResponseBodyAclEntries `json:"AclEntries,omitempty" xml:"AclEntries,omitempty" type:"Repeated"`
	// The return value.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Domain names.
	Domains []*GetInstanceEndpointResponseBodyDomains `json:"Domains,omitempty" xml:"Domains,omitempty" type:"Repeated"`
	// Indicates whether the ACL feature is enabled.
	//
	// example:
	//
	// true
	Enable *bool `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 8F3D5EC5-39D1-4C53-A198-48C54C658FA3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The status of the instance.
	//
	// example:
	//
	// RUNNING
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetInstanceEndpointResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceEndpointResponseBody) GoString() string {
	return s.String()
}

func (s *GetInstanceEndpointResponseBody) SetAclEnable(v bool) *GetInstanceEndpointResponseBody {
	s.AclEnable = &v
	return s
}

func (s *GetInstanceEndpointResponseBody) SetAclEntries(v []*GetInstanceEndpointResponseBodyAclEntries) *GetInstanceEndpointResponseBody {
	s.AclEntries = v
	return s
}

func (s *GetInstanceEndpointResponseBody) SetCode(v string) *GetInstanceEndpointResponseBody {
	s.Code = &v
	return s
}

func (s *GetInstanceEndpointResponseBody) SetDomains(v []*GetInstanceEndpointResponseBodyDomains) *GetInstanceEndpointResponseBody {
	s.Domains = v
	return s
}

func (s *GetInstanceEndpointResponseBody) SetEnable(v bool) *GetInstanceEndpointResponseBody {
	s.Enable = &v
	return s
}

func (s *GetInstanceEndpointResponseBody) SetIsSuccess(v bool) *GetInstanceEndpointResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *GetInstanceEndpointResponseBody) SetRequestId(v string) *GetInstanceEndpointResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetInstanceEndpointResponseBody) SetStatus(v string) *GetInstanceEndpointResponseBody {
	s.Status = &v
	return s
}

type GetInstanceEndpointResponseBodyAclEntries struct {
	// Remarks for public IP address whitelists.
	//
	// example:
	//
	// 1
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The public IP address whitelist.
	//
	// example:
	//
	// 192.168.1.0/24
	Entry *string `json:"Entry,omitempty" xml:"Entry,omitempty"`
}

func (s GetInstanceEndpointResponseBodyAclEntries) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceEndpointResponseBodyAclEntries) GoString() string {
	return s.String()
}

func (s *GetInstanceEndpointResponseBodyAclEntries) SetComment(v string) *GetInstanceEndpointResponseBodyAclEntries {
	s.Comment = &v
	return s
}

func (s *GetInstanceEndpointResponseBodyAclEntries) SetEntry(v string) *GetInstanceEndpointResponseBodyAclEntries {
	s.Entry = &v
	return s
}

type GetInstanceEndpointResponseBodyDomains struct {
	// The domain name that is used to access the Container Registry Enterprise Edition instance.
	//
	// example:
	//
	// shanghai-instance1-registry.cn-shanghai.cr.aliyuncs.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The type of the domain name. Valid values:
	//
	// 	- `SYSTEM`: a system domain name.
	//
	// 	- `USER`: a user domain name.
	//
	// example:
	//
	// SYSTEM
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetInstanceEndpointResponseBodyDomains) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceEndpointResponseBodyDomains) GoString() string {
	return s.String()
}

func (s *GetInstanceEndpointResponseBodyDomains) SetDomain(v string) *GetInstanceEndpointResponseBodyDomains {
	s.Domain = &v
	return s
}

func (s *GetInstanceEndpointResponseBodyDomains) SetType(v string) *GetInstanceEndpointResponseBodyDomains {
	s.Type = &v
	return s
}

type GetInstanceEndpointResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetInstanceEndpointResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetInstanceEndpointResponse) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceEndpointResponse) GoString() string {
	return s.String()
}

func (s *GetInstanceEndpointResponse) SetHeaders(v map[string]*string) *GetInstanceEndpointResponse {
	s.Headers = v
	return s
}

func (s *GetInstanceEndpointResponse) SetStatusCode(v int32) *GetInstanceEndpointResponse {
	s.StatusCode = &v
	return s
}

func (s *GetInstanceEndpointResponse) SetBody(v *GetInstanceEndpointResponseBody) *GetInstanceEndpointResponse {
	s.Body = v
	return s
}

type GetInstanceUsageRequest struct {
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cri-xkx6vujuhay0****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetInstanceUsageRequest) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceUsageRequest) GoString() string {
	return s.String()
}

func (s *GetInstanceUsageRequest) SetInstanceId(v string) *GetInstanceUsageRequest {
	s.InstanceId = &v
	return s
}

type GetInstanceUsageResponseBody struct {
	// The quota of chart namespaces.
	//
	// example:
	//
	// 50
	ChartNamespaceQuota *string `json:"ChartNamespaceQuota,omitempty" xml:"ChartNamespaceQuota,omitempty"`
	// The number of chart namespaces that are created in the instance.
	//
	// example:
	//
	// 2
	ChartNamespaceUsage *string `json:"ChartNamespaceUsage,omitempty" xml:"ChartNamespaceUsage,omitempty"`
	// The quota of chart repositories for the instance.
	//
	// example:
	//
	// 5000
	ChartRepoQuota *string `json:"ChartRepoQuota,omitempty" xml:"ChartRepoQuota,omitempty"`
	// The number of chart repositories that are created.
	//
	// example:
	//
	// 5
	ChartRepoUsage *string `json:"ChartRepoUsage,omitempty" xml:"ChartRepoUsage,omitempty"`
	// The return value.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// 	- `true`: The request is successful.
	//
	// 	- `false`: The request fails.
	//
	// example:
	//
	// true
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The quota of image namespaces for the instance.
	//
	// example:
	//
	// 100
	NamespaceQuota *string `json:"NamespaceQuota,omitempty" xml:"NamespaceQuota,omitempty"`
	// The number of image namespaces that are created in the instance.
	//
	// example:
	//
	// 4
	NamespaceUsage *string `json:"NamespaceUsage,omitempty" xml:"NamespaceUsage,omitempty"`
	// The quota of image repositories for the instance.
	//
	// example:
	//
	// 1000
	RepoQuota *string `json:"RepoQuota,omitempty" xml:"RepoQuota,omitempty"`
	// The number of image repositories that are created in the instance.
	//
	// example:
	//
	// 2
	RepoUsage *string `json:"RepoUsage,omitempty" xml:"RepoUsage,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// A726E801-7FCF-43F9-AF1C-51B3E65D3E7A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// VPC quota
	//
	// example:
	//
	// 5
	VpcQuota *string `json:"VpcQuota,omitempty" xml:"VpcQuota,omitempty"`
	// Number of bound VPCs
	//
	// example:
	//
	// 2
	VpcUsage *string `json:"VpcUsage,omitempty" xml:"VpcUsage,omitempty"`
}

func (s GetInstanceUsageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceUsageResponseBody) GoString() string {
	return s.String()
}

func (s *GetInstanceUsageResponseBody) SetChartNamespaceQuota(v string) *GetInstanceUsageResponseBody {
	s.ChartNamespaceQuota = &v
	return s
}

func (s *GetInstanceUsageResponseBody) SetChartNamespaceUsage(v string) *GetInstanceUsageResponseBody {
	s.ChartNamespaceUsage = &v
	return s
}

func (s *GetInstanceUsageResponseBody) SetChartRepoQuota(v string) *GetInstanceUsageResponseBody {
	s.ChartRepoQuota = &v
	return s
}

func (s *GetInstanceUsageResponseBody) SetChartRepoUsage(v string) *GetInstanceUsageResponseBody {
	s.ChartRepoUsage = &v
	return s
}

func (s *GetInstanceUsageResponseBody) SetCode(v string) *GetInstanceUsageResponseBody {
	s.Code = &v
	return s
}

func (s *GetInstanceUsageResponseBody) SetIsSuccess(v bool) *GetInstanceUsageResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *GetInstanceUsageResponseBody) SetNamespaceQuota(v string) *GetInstanceUsageResponseBody {
	s.NamespaceQuota = &v
	return s
}

func (s *GetInstanceUsageResponseBody) SetNamespaceUsage(v string) *GetInstanceUsageResponseBody {
	s.NamespaceUsage = &v
	return s
}

func (s *GetInstanceUsageResponseBody) SetRepoQuota(v string) *GetInstanceUsageResponseBody {
	s.RepoQuota = &v
	return s
}

func (s *GetInstanceUsageResponseBody) SetRepoUsage(v string) *GetInstanceUsageResponseBody {
	s.RepoUsage = &v
	return s
}

func (s *GetInstanceUsageResponseBody) SetRequestId(v string) *GetInstanceUsageResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetInstanceUsageResponseBody) SetVpcQuota(v string) *GetInstanceUsageResponseBody {
	s.VpcQuota = &v
	return s
}

func (s *GetInstanceUsageResponseBody) SetVpcUsage(v string) *GetInstanceUsageResponseBody {
	s.VpcUsage = &v
	return s
}

type GetInstanceUsageResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetInstanceUsageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetInstanceUsageResponse) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceUsageResponse) GoString() string {
	return s.String()
}

func (s *GetInstanceUsageResponse) SetHeaders(v map[string]*string) *GetInstanceUsageResponse {
	s.Headers = v
	return s
}

func (s *GetInstanceUsageResponse) SetStatusCode(v int32) *GetInstanceUsageResponse {
	s.StatusCode = &v
	return s
}

func (s *GetInstanceUsageResponse) SetBody(v *GetInstanceUsageResponseBody) *GetInstanceUsageResponse {
	s.Body = v
	return s
}

type GetInstanceVpcEndpointRequest struct {
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cri-kmsiwlxxdcva****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the module that you want to access. Valid values:
	//
	// 	- `Registry`: the image repository.
	//
	// 	- `Chart`: a Helm chart.
	//
	// example:
	//
	// Chart
	ModuleName *string `json:"ModuleName,omitempty" xml:"ModuleName,omitempty"`
}

func (s GetInstanceVpcEndpointRequest) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceVpcEndpointRequest) GoString() string {
	return s.String()
}

func (s *GetInstanceVpcEndpointRequest) SetInstanceId(v string) *GetInstanceVpcEndpointRequest {
	s.InstanceId = &v
	return s
}

func (s *GetInstanceVpcEndpointRequest) SetModuleName(v string) *GetInstanceVpcEndpointRequest {
	s.ModuleName = &v
	return s
}

type GetInstanceVpcEndpointResponseBody struct {
	// The return value.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Domain names.
	Domains []*string `json:"Domains,omitempty" xml:"Domains,omitempty" type:"Repeated"`
	// Indicates whether the VPC endpoint is enabled. Valid values:
	//
	// 	- `true`
	//
	// 	- `false`
	//
	// example:
	//
	// true
	Enable *bool `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// 	- `true`: The request is successful.
	//
	// 	- `false`: The request fails.
	//
	// example:
	//
	// true
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The VPCs that are added to the access control list.
	LinkedVpcs []*GetInstanceVpcEndpointResponseBodyLinkedVpcs `json:"LinkedVpcs,omitempty" xml:"LinkedVpcs,omitempty" type:"Repeated"`
	// The name of the modules that can be accessed. Valid values:
	//
	// 	- `Registry`: image repositories.
	//
	// 	- `Chart`: Helm charts.
	//
	// example:
	//
	// Registry
	ModuleName *string `json:"ModuleName,omitempty" xml:"ModuleName,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// BAE9349D-A587-4F9A-B574-9DA0EF2638D1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetInstanceVpcEndpointResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceVpcEndpointResponseBody) GoString() string {
	return s.String()
}

func (s *GetInstanceVpcEndpointResponseBody) SetCode(v string) *GetInstanceVpcEndpointResponseBody {
	s.Code = &v
	return s
}

func (s *GetInstanceVpcEndpointResponseBody) SetDomains(v []*string) *GetInstanceVpcEndpointResponseBody {
	s.Domains = v
	return s
}

func (s *GetInstanceVpcEndpointResponseBody) SetEnable(v bool) *GetInstanceVpcEndpointResponseBody {
	s.Enable = &v
	return s
}

func (s *GetInstanceVpcEndpointResponseBody) SetIsSuccess(v bool) *GetInstanceVpcEndpointResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *GetInstanceVpcEndpointResponseBody) SetLinkedVpcs(v []*GetInstanceVpcEndpointResponseBodyLinkedVpcs) *GetInstanceVpcEndpointResponseBody {
	s.LinkedVpcs = v
	return s
}

func (s *GetInstanceVpcEndpointResponseBody) SetModuleName(v string) *GetInstanceVpcEndpointResponseBody {
	s.ModuleName = &v
	return s
}

func (s *GetInstanceVpcEndpointResponseBody) SetRequestId(v string) *GetInstanceVpcEndpointResponseBody {
	s.RequestId = &v
	return s
}

type GetInstanceVpcEndpointResponseBodyLinkedVpcs struct {
	// Indicates whether the VPC is the default VPC over which the Container Registry instance is accessed.
	//
	// example:
	//
	// false
	DefaultAccess *bool `json:"DefaultAccess,omitempty" xml:"DefaultAccess,omitempty"`
	// IP address.
	//
	// example:
	//
	// 192.168.10.11
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// The status of the VPC. Valid values:
	//
	// 	- `CREATING`
	//
	// 	- `RUNNING`
	//
	// example:
	//
	// CREATING
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// VPC ID
	//
	// example:
	//
	// vpc-uf6aamu2nomfr1thd****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The ID of the vSwitch.
	//
	// example:
	//
	// vsw-uf62m5vmxl2e72dk7****
	VswitchId *string `json:"VswitchId,omitempty" xml:"VswitchId,omitempty"`
}

func (s GetInstanceVpcEndpointResponseBodyLinkedVpcs) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceVpcEndpointResponseBodyLinkedVpcs) GoString() string {
	return s.String()
}

func (s *GetInstanceVpcEndpointResponseBodyLinkedVpcs) SetDefaultAccess(v bool) *GetInstanceVpcEndpointResponseBodyLinkedVpcs {
	s.DefaultAccess = &v
	return s
}

func (s *GetInstanceVpcEndpointResponseBodyLinkedVpcs) SetIp(v string) *GetInstanceVpcEndpointResponseBodyLinkedVpcs {
	s.Ip = &v
	return s
}

func (s *GetInstanceVpcEndpointResponseBodyLinkedVpcs) SetStatus(v string) *GetInstanceVpcEndpointResponseBodyLinkedVpcs {
	s.Status = &v
	return s
}

func (s *GetInstanceVpcEndpointResponseBodyLinkedVpcs) SetVpcId(v string) *GetInstanceVpcEndpointResponseBodyLinkedVpcs {
	s.VpcId = &v
	return s
}

func (s *GetInstanceVpcEndpointResponseBodyLinkedVpcs) SetVswitchId(v string) *GetInstanceVpcEndpointResponseBodyLinkedVpcs {
	s.VswitchId = &v
	return s
}

type GetInstanceVpcEndpointResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetInstanceVpcEndpointResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetInstanceVpcEndpointResponse) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceVpcEndpointResponse) GoString() string {
	return s.String()
}

func (s *GetInstanceVpcEndpointResponse) SetHeaders(v map[string]*string) *GetInstanceVpcEndpointResponse {
	s.Headers = v
	return s
}

func (s *GetInstanceVpcEndpointResponse) SetStatusCode(v int32) *GetInstanceVpcEndpointResponse {
	s.StatusCode = &v
	return s
}

func (s *GetInstanceVpcEndpointResponse) SetBody(v *GetInstanceVpcEndpointResponseBody) *GetInstanceVpcEndpointResponse {
	s.Body = v
	return s
}

type GetNamespaceRequest struct {
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cri-kmsiwlxxdcva****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the namespace.
	//
	// example:
	//
	// crn-tiw8t3f8i5lta****
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	// The name of the namespace.
	//
	// example:
	//
	// test
	NamespaceName *string `json:"NamespaceName,omitempty" xml:"NamespaceName,omitempty"`
}

func (s GetNamespaceRequest) String() string {
	return tea.Prettify(s)
}

func (s GetNamespaceRequest) GoString() string {
	return s.String()
}

func (s *GetNamespaceRequest) SetInstanceId(v string) *GetNamespaceRequest {
	s.InstanceId = &v
	return s
}

func (s *GetNamespaceRequest) SetNamespaceId(v string) *GetNamespaceRequest {
	s.NamespaceId = &v
	return s
}

func (s *GetNamespaceRequest) SetNamespaceName(v string) *GetNamespaceRequest {
	s.NamespaceName = &v
	return s
}

type GetNamespaceResponseBody struct {
	// Indicates whether a repository is automatically created when an image is pushed to the namespace.
	//
	// example:
	//
	// true
	AutoCreateRepo *bool `json:"AutoCreateRepo,omitempty" xml:"AutoCreateRepo,omitempty"`
	// The return value.
	//
	// example:
	//
	// success
	Code                     *string            `json:"Code,omitempty" xml:"Code,omitempty"`
	DefaultRepoConfiguration *RepoConfiguration `json:"DefaultRepoConfiguration,omitempty" xml:"DefaultRepoConfiguration,omitempty"`
	// Deprecated
	//
	// The default type of repositories in the namespace. Valid values:
	//
	// 	- PUBLIC: public repositories.
	//
	// 	- PRIVATE: private repositories.
	//
	// example:
	//
	// PUBLIC
	DefaultRepoType *string `json:"DefaultRepoType,omitempty" xml:"DefaultRepoType,omitempty"`
	// The ID of the Container Registry instance.
	//
	// example:
	//
	// cri-kmsiwlxxdcva****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The ID of the namespace.
	//
	// example:
	//
	// crn-tiw8t3f8i5lt****
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	// The name of the namespace.
	//
	// example:
	//
	// test
	NamespaceName *string `json:"NamespaceName,omitempty" xml:"NamespaceName,omitempty"`
	// The status of the namespace.
	//
	// 	- NORMAL
	//
	// 	- DELETING
	//
	// example:
	//
	// NORMAL
	NamespaceStatus *string `json:"NamespaceStatus,omitempty" xml:"NamespaceStatus,omitempty"`
	// The request ID.
	//
	// example:
	//
	// E4BC9E21-8AA5-4582-83C1-C1209AB8196F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the resource group to which the namespace belongs.
	//
	// example:
	//
	// rg-acfmv36i4is****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s GetNamespaceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetNamespaceResponseBody) GoString() string {
	return s.String()
}

func (s *GetNamespaceResponseBody) SetAutoCreateRepo(v bool) *GetNamespaceResponseBody {
	s.AutoCreateRepo = &v
	return s
}

func (s *GetNamespaceResponseBody) SetCode(v string) *GetNamespaceResponseBody {
	s.Code = &v
	return s
}

func (s *GetNamespaceResponseBody) SetDefaultRepoConfiguration(v *RepoConfiguration) *GetNamespaceResponseBody {
	s.DefaultRepoConfiguration = v
	return s
}

func (s *GetNamespaceResponseBody) SetDefaultRepoType(v string) *GetNamespaceResponseBody {
	s.DefaultRepoType = &v
	return s
}

func (s *GetNamespaceResponseBody) SetInstanceId(v string) *GetNamespaceResponseBody {
	s.InstanceId = &v
	return s
}

func (s *GetNamespaceResponseBody) SetIsSuccess(v bool) *GetNamespaceResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *GetNamespaceResponseBody) SetNamespaceId(v string) *GetNamespaceResponseBody {
	s.NamespaceId = &v
	return s
}

func (s *GetNamespaceResponseBody) SetNamespaceName(v string) *GetNamespaceResponseBody {
	s.NamespaceName = &v
	return s
}

func (s *GetNamespaceResponseBody) SetNamespaceStatus(v string) *GetNamespaceResponseBody {
	s.NamespaceStatus = &v
	return s
}

func (s *GetNamespaceResponseBody) SetRequestId(v string) *GetNamespaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetNamespaceResponseBody) SetResourceGroupId(v string) *GetNamespaceResponseBody {
	s.ResourceGroupId = &v
	return s
}

type GetNamespaceResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetNamespaceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetNamespaceResponse) String() string {
	return tea.Prettify(s)
}

func (s GetNamespaceResponse) GoString() string {
	return s.String()
}

func (s *GetNamespaceResponse) SetHeaders(v map[string]*string) *GetNamespaceResponse {
	s.Headers = v
	return s
}

func (s *GetNamespaceResponse) SetStatusCode(v int32) *GetNamespaceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetNamespaceResponse) SetBody(v *GetNamespaceResponseBody) *GetNamespaceResponse {
	s.Body = v
	return s
}

type GetRepoBuildRecordRequest struct {
	// The ID of the image building record.
	//
	// This parameter is required.
	//
	// example:
	//
	// a78ec6fb-16ea-4649-93b7-f52afba7d****
	BuildRecordId *string `json:"BuildRecordId,omitempty" xml:"BuildRecordId,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cri-kmsiwlxxdcva****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetRepoBuildRecordRequest) String() string {
	return tea.Prettify(s)
}

func (s GetRepoBuildRecordRequest) GoString() string {
	return s.String()
}

func (s *GetRepoBuildRecordRequest) SetBuildRecordId(v string) *GetRepoBuildRecordRequest {
	s.BuildRecordId = &v
	return s
}

func (s *GetRepoBuildRecordRequest) SetInstanceId(v string) *GetRepoBuildRecordRequest {
	s.InstanceId = &v
	return s
}

type GetRepoBuildRecordResponseBody struct {
	// The ID of the image building record.
	//
	// example:
	//
	// 79174CBA-8556-443A-8976-22C922D7****
	BuildRecordId *string `json:"BuildRecordId,omitempty" xml:"BuildRecordId,omitempty"`
	// The return value.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The time when the image building was completed.
	//
	// example:
	//
	// 1568718698000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The information about the image.
	Image *GetRepoBuildRecordResponseBodyImage `json:"Image,omitempty" xml:"Image,omitempty" type:"Struct"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// a78ec6fb-16ea-4649-93b7-f52afba7d9de1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The time when the image building started.
	//
	// example:
	//
	// 1568718468000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The status of the instance.
	//
	// example:
	//
	// true
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetRepoBuildRecordResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetRepoBuildRecordResponseBody) GoString() string {
	return s.String()
}

func (s *GetRepoBuildRecordResponseBody) SetBuildRecordId(v string) *GetRepoBuildRecordResponseBody {
	s.BuildRecordId = &v
	return s
}

func (s *GetRepoBuildRecordResponseBody) SetCode(v string) *GetRepoBuildRecordResponseBody {
	s.Code = &v
	return s
}

func (s *GetRepoBuildRecordResponseBody) SetEndTime(v int64) *GetRepoBuildRecordResponseBody {
	s.EndTime = &v
	return s
}

func (s *GetRepoBuildRecordResponseBody) SetImage(v *GetRepoBuildRecordResponseBodyImage) *GetRepoBuildRecordResponseBody {
	s.Image = v
	return s
}

func (s *GetRepoBuildRecordResponseBody) SetIsSuccess(v bool) *GetRepoBuildRecordResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *GetRepoBuildRecordResponseBody) SetRequestId(v string) *GetRepoBuildRecordResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRepoBuildRecordResponseBody) SetStartTime(v int64) *GetRepoBuildRecordResponseBody {
	s.StartTime = &v
	return s
}

func (s *GetRepoBuildRecordResponseBody) SetStatus(v string) *GetRepoBuildRecordResponseBody {
	s.Status = &v
	return s
}

type GetRepoBuildRecordResponseBodyImage struct {
	// The tag of the image.
	//
	// example:
	//
	// master
	ImageTag *string `json:"ImageTag,omitempty" xml:"ImageTag,omitempty"`
	// The name of the image repository.
	//
	// example:
	//
	// test
	RepoName *string `json:"RepoName,omitempty" xml:"RepoName,omitempty"`
	// The name of the namespace to which the image repository belongs.
	//
	// example:
	//
	// test
	RepoNamespaceName *string `json:"RepoNamespaceName,omitempty" xml:"RepoNamespaceName,omitempty"`
}

func (s GetRepoBuildRecordResponseBodyImage) String() string {
	return tea.Prettify(s)
}

func (s GetRepoBuildRecordResponseBodyImage) GoString() string {
	return s.String()
}

func (s *GetRepoBuildRecordResponseBodyImage) SetImageTag(v string) *GetRepoBuildRecordResponseBodyImage {
	s.ImageTag = &v
	return s
}

func (s *GetRepoBuildRecordResponseBodyImage) SetRepoName(v string) *GetRepoBuildRecordResponseBodyImage {
	s.RepoName = &v
	return s
}

func (s *GetRepoBuildRecordResponseBodyImage) SetRepoNamespaceName(v string) *GetRepoBuildRecordResponseBodyImage {
	s.RepoNamespaceName = &v
	return s
}

type GetRepoBuildRecordResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetRepoBuildRecordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetRepoBuildRecordResponse) String() string {
	return tea.Prettify(s)
}

func (s GetRepoBuildRecordResponse) GoString() string {
	return s.String()
}

func (s *GetRepoBuildRecordResponse) SetHeaders(v map[string]*string) *GetRepoBuildRecordResponse {
	s.Headers = v
	return s
}

func (s *GetRepoBuildRecordResponse) SetStatusCode(v int32) *GetRepoBuildRecordResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRepoBuildRecordResponse) SetBody(v *GetRepoBuildRecordResponseBody) *GetRepoBuildRecordResponse {
	s.Body = v
	return s
}

type GetRepoBuildRecordStatusRequest struct {
	// The ID of the image building record.
	//
	// This parameter is required.
	//
	// example:
	//
	// a78ec6fb-16ea-4649-93b7-f52afba7d****
	BuildRecordId *string `json:"BuildRecordId,omitempty" xml:"BuildRecordId,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cri-kmsiwlxxdcva****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the image repository.
	//
	// This parameter is required.
	//
	// example:
	//
	// crr-jnzm47ihjmgc****
	RepoId *string `json:"RepoId,omitempty" xml:"RepoId,omitempty"`
}

func (s GetRepoBuildRecordStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s GetRepoBuildRecordStatusRequest) GoString() string {
	return s.String()
}

func (s *GetRepoBuildRecordStatusRequest) SetBuildRecordId(v string) *GetRepoBuildRecordStatusRequest {
	s.BuildRecordId = &v
	return s
}

func (s *GetRepoBuildRecordStatusRequest) SetInstanceId(v string) *GetRepoBuildRecordStatusRequest {
	s.InstanceId = &v
	return s
}

func (s *GetRepoBuildRecordStatusRequest) SetRepoId(v string) *GetRepoBuildRecordStatusRequest {
	s.RepoId = &v
	return s
}

type GetRepoBuildRecordStatusResponseBody struct {
	// The status of the image building.
	//
	// example:
	//
	// success
	BuildStatus *string `json:"BuildStatus,omitempty" xml:"BuildStatus,omitempty"`
	// The return value.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 79174CBA-8556-443A-8976-22C922D7BE37
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetRepoBuildRecordStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetRepoBuildRecordStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetRepoBuildRecordStatusResponseBody) SetBuildStatus(v string) *GetRepoBuildRecordStatusResponseBody {
	s.BuildStatus = &v
	return s
}

func (s *GetRepoBuildRecordStatusResponseBody) SetCode(v string) *GetRepoBuildRecordStatusResponseBody {
	s.Code = &v
	return s
}

func (s *GetRepoBuildRecordStatusResponseBody) SetIsSuccess(v bool) *GetRepoBuildRecordStatusResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *GetRepoBuildRecordStatusResponseBody) SetRequestId(v string) *GetRepoBuildRecordStatusResponseBody {
	s.RequestId = &v
	return s
}

type GetRepoBuildRecordStatusResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetRepoBuildRecordStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetRepoBuildRecordStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s GetRepoBuildRecordStatusResponse) GoString() string {
	return s.String()
}

func (s *GetRepoBuildRecordStatusResponse) SetHeaders(v map[string]*string) *GetRepoBuildRecordStatusResponse {
	s.Headers = v
	return s
}

func (s *GetRepoBuildRecordStatusResponse) SetStatusCode(v int32) *GetRepoBuildRecordStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRepoBuildRecordStatusResponse) SetBody(v *GetRepoBuildRecordStatusResponseBody) *GetRepoBuildRecordStatusResponse {
	s.Body = v
	return s
}

type GetRepoSourceCodeRepoRequest struct {
	// The ID of the Container Registry instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cri-shac42yvqzvq****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the repository.
	//
	// This parameter is required.
	//
	// example:
	//
	// crr-gzsrlevmvoaq****
	RepoId *string `json:"RepoId,omitempty" xml:"RepoId,omitempty"`
}

func (s GetRepoSourceCodeRepoRequest) String() string {
	return tea.Prettify(s)
}

func (s GetRepoSourceCodeRepoRequest) GoString() string {
	return s.String()
}

func (s *GetRepoSourceCodeRepoRequest) SetInstanceId(v string) *GetRepoSourceCodeRepoRequest {
	s.InstanceId = &v
	return s
}

func (s *GetRepoSourceCodeRepoRequest) SetRepoId(v string) *GetRepoSourceCodeRepoRequest {
	s.RepoId = &v
	return s
}

type GetRepoSourceCodeRepoResponseBody struct {
	// Indicates whether image building is automatically triggered when source code is committed. Valid values:
	//
	// 	- `true`: Image building is automatically triggered when source code is committed.
	//
	// 	- `false`: Image building is not triggered when source code is committed.
	//
	// example:
	//
	// true
	AutoBuild *string `json:"AutoBuild,omitempty" xml:"AutoBuild,omitempty"`
	// The response code.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The address of the source code repository.
	//
	// example:
	//
	// https://github.com
	CodeRepoDomain *string `json:"CodeRepoDomain,omitempty" xml:"CodeRepoDomain,omitempty"`
	// The name of the source code repository.
	//
	// example:
	//
	// repo
	CodeRepoName *string `json:"CodeRepoName,omitempty" xml:"CodeRepoName,omitempty"`
	// The namespace to which the source code repository belongs.
	//
	// example:
	//
	// namespace
	CodeRepoNamespaceName *string `json:"CodeRepoNamespaceName,omitempty" xml:"CodeRepoNamespaceName,omitempty"`
	// The type of the code hosting platform. Valid values: `GITHUB`, `GITLAB`, `GITEE`, `CODE`, and `CODEUP`.
	//
	// example:
	//
	// GITHUB
	CodeRepoType *string `json:"CodeRepoType,omitempty" xml:"CodeRepoType,omitempty"`
	// Indicates whether build cache is disabled. Valid values:
	//
	// 	- `true`: Build cache is disabled.
	//
	// 	- `false`: Build cache is enabled.
	//
	// example:
	//
	// false
	DisableCacheBuild *string `json:"DisableCacheBuild,omitempty" xml:"DisableCacheBuild,omitempty"`
	// Indicates whether the API call is successful. Valid values:
	//
	// 	- `true`: successful
	//
	// 	- `false`: failed
	//
	// example:
	//
	// true
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// Indicates whether image building is accelerated for servers outside the Chinese mainland. Valid values:
	//
	// 	- `true`: Image building is accelerated for servers outside the Chinese mainland.
	//
	// 	- `false`: Image building is not accelerated for servers outside the Chinese mainland.
	//
	// example:
	//
	// false
	OverseaBuild *string `json:"OverseaBuild,omitempty" xml:"OverseaBuild,omitempty"`
	// The ID of the repository.
	//
	// example:
	//
	// crr-gzsrlevmvoaq****
	RepoId *string `json:"RepoId,omitempty" xml:"RepoId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 4CE1F661-75DD-4EBD-A4AD-057B26834ABB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetRepoSourceCodeRepoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetRepoSourceCodeRepoResponseBody) GoString() string {
	return s.String()
}

func (s *GetRepoSourceCodeRepoResponseBody) SetAutoBuild(v string) *GetRepoSourceCodeRepoResponseBody {
	s.AutoBuild = &v
	return s
}

func (s *GetRepoSourceCodeRepoResponseBody) SetCode(v string) *GetRepoSourceCodeRepoResponseBody {
	s.Code = &v
	return s
}

func (s *GetRepoSourceCodeRepoResponseBody) SetCodeRepoDomain(v string) *GetRepoSourceCodeRepoResponseBody {
	s.CodeRepoDomain = &v
	return s
}

func (s *GetRepoSourceCodeRepoResponseBody) SetCodeRepoName(v string) *GetRepoSourceCodeRepoResponseBody {
	s.CodeRepoName = &v
	return s
}

func (s *GetRepoSourceCodeRepoResponseBody) SetCodeRepoNamespaceName(v string) *GetRepoSourceCodeRepoResponseBody {
	s.CodeRepoNamespaceName = &v
	return s
}

func (s *GetRepoSourceCodeRepoResponseBody) SetCodeRepoType(v string) *GetRepoSourceCodeRepoResponseBody {
	s.CodeRepoType = &v
	return s
}

func (s *GetRepoSourceCodeRepoResponseBody) SetDisableCacheBuild(v string) *GetRepoSourceCodeRepoResponseBody {
	s.DisableCacheBuild = &v
	return s
}

func (s *GetRepoSourceCodeRepoResponseBody) SetIsSuccess(v bool) *GetRepoSourceCodeRepoResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *GetRepoSourceCodeRepoResponseBody) SetOverseaBuild(v string) *GetRepoSourceCodeRepoResponseBody {
	s.OverseaBuild = &v
	return s
}

func (s *GetRepoSourceCodeRepoResponseBody) SetRepoId(v string) *GetRepoSourceCodeRepoResponseBody {
	s.RepoId = &v
	return s
}

func (s *GetRepoSourceCodeRepoResponseBody) SetRequestId(v string) *GetRepoSourceCodeRepoResponseBody {
	s.RequestId = &v
	return s
}

type GetRepoSourceCodeRepoResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetRepoSourceCodeRepoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetRepoSourceCodeRepoResponse) String() string {
	return tea.Prettify(s)
}

func (s GetRepoSourceCodeRepoResponse) GoString() string {
	return s.String()
}

func (s *GetRepoSourceCodeRepoResponse) SetHeaders(v map[string]*string) *GetRepoSourceCodeRepoResponse {
	s.Headers = v
	return s
}

func (s *GetRepoSourceCodeRepoResponse) SetStatusCode(v int32) *GetRepoSourceCodeRepoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRepoSourceCodeRepoResponse) SetBody(v *GetRepoSourceCodeRepoResponseBody) *GetRepoSourceCodeRepoResponse {
	s.Body = v
	return s
}

type GetRepoSyncTaskRequest struct {
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cri-sgedpenzw80e****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the synchronization task.
	//
	// This parameter is required.
	//
	// example:
	//
	// rst-zxjkiv5oil6f****
	SyncTaskId *string `json:"SyncTaskId,omitempty" xml:"SyncTaskId,omitempty"`
}

func (s GetRepoSyncTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s GetRepoSyncTaskRequest) GoString() string {
	return s.String()
}

func (s *GetRepoSyncTaskRequest) SetInstanceId(v string) *GetRepoSyncTaskRequest {
	s.InstanceId = &v
	return s
}

func (s *GetRepoSyncTaskRequest) SetSyncTaskId(v string) *GetRepoSyncTaskRequest {
	s.SyncTaskId = &v
	return s
}

type GetRepoSyncTaskResponseBody struct {
	// The return value.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the synchronization task is performed across Alibaba Cloud accounts.
	//
	// example:
	//
	// true
	CrossUser *bool `json:"CrossUser,omitempty" xml:"CrossUser,omitempty"`
	// The source address of the image.
	ImageFrom *GetRepoSyncTaskResponseBodyImageFrom `json:"ImageFrom,omitempty" xml:"ImageFrom,omitempty" type:"Struct"`
	// The destination address of the image.
	ImageTo *GetRepoSyncTaskResponseBodyImageTo `json:"ImageTo,omitempty" xml:"ImageTo,omitempty" type:"Struct"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The synchronization tasks for the image layer.
	LayerTasks []*GetRepoSyncTaskResponseBodyLayerTasks `json:"LayerTasks,omitempty" xml:"LayerTasks,omitempty" type:"Repeated"`
	// The synchronization progress. Valid values:
	//
	// 	- `0`: The synchronization starts or failed.
	//
	// 	- `1`: The synchronization is successful.
	//
	// example:
	//
	// 1
	Progress *int64 `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// A6DEF8B0-5D45-46D6-867D-8C7FF0966B07
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the synchronization task in which multiple images are synchronized at a time.
	//
	// example:
	//
	// a9434731-95ef-4087-9cf4-369c8e90****
	SyncBatchTaskId *string `json:"SyncBatchTaskId,omitempty" xml:"SyncBatchTaskId,omitempty"`
	// The ID of the synchronization rule.
	//
	// example:
	//
	// crsr-cllro6ho3wne****
	SyncRuleId *string `json:"SyncRuleId,omitempty" xml:"SyncRuleId,omitempty"`
	// The ID of the synchronization task.
	//
	// example:
	//
	// rst-zxjkiv5oil6f****
	SyncTaskId *string `json:"SyncTaskId,omitempty" xml:"SyncTaskId,omitempty"`
	// Indicates whether transfer acceleration is enabled in the synchronization process.
	//
	// example:
	//
	// true
	SyncTransAccelerate *bool `json:"SyncTransAccelerate,omitempty" xml:"SyncTransAccelerate,omitempty"`
	// The size of the image layer that is synchronized. Unit: bytes.
	//
	// example:
	//
	// 23655489
	SyncedSize *int64 `json:"SyncedSize,omitempty" xml:"SyncedSize,omitempty"`
	// The error message that is returned if the synchronization task fails.
	//
	// >  The system uses this parameter to return an error message if the synchronization task fails.
	//
	// Valid values:
	//
	// 	- OSS_POLICY_UNAUTHORIZED: Container Registry is not granted permissions to use Object Storage Service (OSS).
	//
	// 	- TAG_CONFLICT: The destination repository contains an image that has the same tag as the source image, and image tag immutability is enabled for the destination repository.
	//
	// 	- UNSUPPORTED_FORMAT: The manifest and config formats of the image to be synchronized are not supported.
	//
	// 	- INTERNAL_ERROR: The synchronization task failed due to internal issues on the server.
	//
	// 	- NETWORK_ERROR: The synchronization task failed due to unstable network connection.
	//
	// 	- DATA_LENGTH_EXCEEDED: The manifest or config of the image is oversized.
	//
	// example:
	//
	// NETWORK_ERROR
	TaskIssue *string `json:"TaskIssue,omitempty" xml:"TaskIssue,omitempty"`
	// The status of the task. Valid values:
	//
	// example:
	//
	// SUCCESS
	TaskStatus *string `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	// The policy that is used to trigger the synchronization task.
	//
	// example:
	//
	// null
	TaskTrigger *string `json:"TaskTrigger,omitempty" xml:"TaskTrigger,omitempty"`
}

func (s GetRepoSyncTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetRepoSyncTaskResponseBody) GoString() string {
	return s.String()
}

func (s *GetRepoSyncTaskResponseBody) SetCode(v string) *GetRepoSyncTaskResponseBody {
	s.Code = &v
	return s
}

func (s *GetRepoSyncTaskResponseBody) SetCrossUser(v bool) *GetRepoSyncTaskResponseBody {
	s.CrossUser = &v
	return s
}

func (s *GetRepoSyncTaskResponseBody) SetImageFrom(v *GetRepoSyncTaskResponseBodyImageFrom) *GetRepoSyncTaskResponseBody {
	s.ImageFrom = v
	return s
}

func (s *GetRepoSyncTaskResponseBody) SetImageTo(v *GetRepoSyncTaskResponseBodyImageTo) *GetRepoSyncTaskResponseBody {
	s.ImageTo = v
	return s
}

func (s *GetRepoSyncTaskResponseBody) SetIsSuccess(v bool) *GetRepoSyncTaskResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *GetRepoSyncTaskResponseBody) SetLayerTasks(v []*GetRepoSyncTaskResponseBodyLayerTasks) *GetRepoSyncTaskResponseBody {
	s.LayerTasks = v
	return s
}

func (s *GetRepoSyncTaskResponseBody) SetProgress(v int64) *GetRepoSyncTaskResponseBody {
	s.Progress = &v
	return s
}

func (s *GetRepoSyncTaskResponseBody) SetRequestId(v string) *GetRepoSyncTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRepoSyncTaskResponseBody) SetSyncBatchTaskId(v string) *GetRepoSyncTaskResponseBody {
	s.SyncBatchTaskId = &v
	return s
}

func (s *GetRepoSyncTaskResponseBody) SetSyncRuleId(v string) *GetRepoSyncTaskResponseBody {
	s.SyncRuleId = &v
	return s
}

func (s *GetRepoSyncTaskResponseBody) SetSyncTaskId(v string) *GetRepoSyncTaskResponseBody {
	s.SyncTaskId = &v
	return s
}

func (s *GetRepoSyncTaskResponseBody) SetSyncTransAccelerate(v bool) *GetRepoSyncTaskResponseBody {
	s.SyncTransAccelerate = &v
	return s
}

func (s *GetRepoSyncTaskResponseBody) SetSyncedSize(v int64) *GetRepoSyncTaskResponseBody {
	s.SyncedSize = &v
	return s
}

func (s *GetRepoSyncTaskResponseBody) SetTaskIssue(v string) *GetRepoSyncTaskResponseBody {
	s.TaskIssue = &v
	return s
}

func (s *GetRepoSyncTaskResponseBody) SetTaskStatus(v string) *GetRepoSyncTaskResponseBody {
	s.TaskStatus = &v
	return s
}

func (s *GetRepoSyncTaskResponseBody) SetTaskTrigger(v string) *GetRepoSyncTaskResponseBody {
	s.TaskTrigger = &v
	return s
}

type GetRepoSyncTaskResponseBodyImageFrom struct {
	// The tag of the image.
	//
	// example:
	//
	// master
	ImageTag *string `json:"ImageTag,omitempty" xml:"ImageTag,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// cri-sgedpenzw80e****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the image repository.
	//
	// example:
	//
	// test
	RepoName *string `json:"RepoName,omitempty" xml:"RepoName,omitempty"`
	// The name of the namespace.
	//
	// example:
	//
	// test
	RepoNamespaceName *string `json:"RepoNamespaceName,omitempty" xml:"RepoNamespaceName,omitempty"`
}

func (s GetRepoSyncTaskResponseBodyImageFrom) String() string {
	return tea.Prettify(s)
}

func (s GetRepoSyncTaskResponseBodyImageFrom) GoString() string {
	return s.String()
}

func (s *GetRepoSyncTaskResponseBodyImageFrom) SetImageTag(v string) *GetRepoSyncTaskResponseBodyImageFrom {
	s.ImageTag = &v
	return s
}

func (s *GetRepoSyncTaskResponseBodyImageFrom) SetInstanceId(v string) *GetRepoSyncTaskResponseBodyImageFrom {
	s.InstanceId = &v
	return s
}

func (s *GetRepoSyncTaskResponseBodyImageFrom) SetRegionId(v string) *GetRepoSyncTaskResponseBodyImageFrom {
	s.RegionId = &v
	return s
}

func (s *GetRepoSyncTaskResponseBodyImageFrom) SetRepoName(v string) *GetRepoSyncTaskResponseBodyImageFrom {
	s.RepoName = &v
	return s
}

func (s *GetRepoSyncTaskResponseBodyImageFrom) SetRepoNamespaceName(v string) *GetRepoSyncTaskResponseBodyImageFrom {
	s.RepoNamespaceName = &v
	return s
}

type GetRepoSyncTaskResponseBodyImageTo struct {
	// The tag of the image.
	//
	// example:
	//
	// master
	ImageTag *string `json:"ImageTag,omitempty" xml:"ImageTag,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// cri-leqzomz5vijc****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// eu-west-1
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the image repository.
	//
	// example:
	//
	// test
	RepoName *string `json:"RepoName,omitempty" xml:"RepoName,omitempty"`
	// The name of the namespace.
	//
	// example:
	//
	// test
	RepoNamespaceName *string `json:"RepoNamespaceName,omitempty" xml:"RepoNamespaceName,omitempty"`
}

func (s GetRepoSyncTaskResponseBodyImageTo) String() string {
	return tea.Prettify(s)
}

func (s GetRepoSyncTaskResponseBodyImageTo) GoString() string {
	return s.String()
}

func (s *GetRepoSyncTaskResponseBodyImageTo) SetImageTag(v string) *GetRepoSyncTaskResponseBodyImageTo {
	s.ImageTag = &v
	return s
}

func (s *GetRepoSyncTaskResponseBodyImageTo) SetInstanceId(v string) *GetRepoSyncTaskResponseBodyImageTo {
	s.InstanceId = &v
	return s
}

func (s *GetRepoSyncTaskResponseBodyImageTo) SetRegionId(v string) *GetRepoSyncTaskResponseBodyImageTo {
	s.RegionId = &v
	return s
}

func (s *GetRepoSyncTaskResponseBodyImageTo) SetRepoName(v string) *GetRepoSyncTaskResponseBodyImageTo {
	s.RepoName = &v
	return s
}

func (s *GetRepoSyncTaskResponseBodyImageTo) SetRepoNamespaceName(v string) *GetRepoSyncTaskResponseBodyImageTo {
	s.RepoNamespaceName = &v
	return s
}

type GetRepoSyncTaskResponseBodyLayerTasks struct {
	// The digest of the artifact.
	//
	// example:
	//
	// sha256:36fb85fcb5e919cb60e782397a6be04201868fe7b38ef7669fc01caec1c8fc4e
	ArtifactDigest *string `json:"ArtifactDigest,omitempty" xml:"ArtifactDigest,omitempty"`
	// The digest of the image layer.
	//
	// example:
	//
	// sha256:36fb85fcb5e919cb60e782397a6be04201868fe7b38ef7669fc01caec1c8fc4e
	Digest *string `json:"Digest,omitempty" xml:"Digest,omitempty"`
	// The size of synchronized image layers.
	//
	// example:
	//
	// 23655489
	Size *int64 `json:"Size,omitempty" xml:"Size,omitempty"`
	// The ID of the synchronization task for the image layer.
	//
	// example:
	//
	// rslt-074x4q20fx2d****
	SyncLayerTaskId *string `json:"SyncLayerTaskId,omitempty" xml:"SyncLayerTaskId,omitempty"`
	// The size of the image layer that is synchronized.
	//
	// example:
	//
	// 23655489
	SyncedSize *int64 `json:"SyncedSize,omitempty" xml:"SyncedSize,omitempty"`
	// The status of the synchronization task. Valid values:
	//
	// example:
	//
	// SUCCESS
	TaskStatus *string `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
}

func (s GetRepoSyncTaskResponseBodyLayerTasks) String() string {
	return tea.Prettify(s)
}

func (s GetRepoSyncTaskResponseBodyLayerTasks) GoString() string {
	return s.String()
}

func (s *GetRepoSyncTaskResponseBodyLayerTasks) SetArtifactDigest(v string) *GetRepoSyncTaskResponseBodyLayerTasks {
	s.ArtifactDigest = &v
	return s
}

func (s *GetRepoSyncTaskResponseBodyLayerTasks) SetDigest(v string) *GetRepoSyncTaskResponseBodyLayerTasks {
	s.Digest = &v
	return s
}

func (s *GetRepoSyncTaskResponseBodyLayerTasks) SetSize(v int64) *GetRepoSyncTaskResponseBodyLayerTasks {
	s.Size = &v
	return s
}

func (s *GetRepoSyncTaskResponseBodyLayerTasks) SetSyncLayerTaskId(v string) *GetRepoSyncTaskResponseBodyLayerTasks {
	s.SyncLayerTaskId = &v
	return s
}

func (s *GetRepoSyncTaskResponseBodyLayerTasks) SetSyncedSize(v int64) *GetRepoSyncTaskResponseBodyLayerTasks {
	s.SyncedSize = &v
	return s
}

func (s *GetRepoSyncTaskResponseBodyLayerTasks) SetTaskStatus(v string) *GetRepoSyncTaskResponseBodyLayerTasks {
	s.TaskStatus = &v
	return s
}

type GetRepoSyncTaskResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetRepoSyncTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetRepoSyncTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s GetRepoSyncTaskResponse) GoString() string {
	return s.String()
}

func (s *GetRepoSyncTaskResponse) SetHeaders(v map[string]*string) *GetRepoSyncTaskResponse {
	s.Headers = v
	return s
}

func (s *GetRepoSyncTaskResponse) SetStatusCode(v int32) *GetRepoSyncTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRepoSyncTaskResponse) SetBody(v *GetRepoSyncTaskResponseBody) *GetRepoSyncTaskResponse {
	s.Body = v
	return s
}

type GetRepoTagRequest struct {
	// The return value of status code.
	//
	// This parameter is required.
	//
	// example:
	//
	// cri-kmsiwlxxdcva****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The operation that you want to perform. Set the value to **GetRepoTag**.
	//
	// This parameter is required.
	//
	// example:
	//
	// crr-tquyps22md8p****
	RepoId *string `json:"RepoId,omitempty" xml:"RepoId,omitempty"`
	// The number of milliseconds that have elapsed since the image was created.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1.0
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
}

func (s GetRepoTagRequest) String() string {
	return tea.Prettify(s)
}

func (s GetRepoTagRequest) GoString() string {
	return s.String()
}

func (s *GetRepoTagRequest) SetInstanceId(v string) *GetRepoTagRequest {
	s.InstanceId = &v
	return s
}

func (s *GetRepoTagRequest) SetRepoId(v string) *GetRepoTagRequest {
	s.RepoId = &v
	return s
}

func (s *GetRepoTagRequest) SetTag(v string) *GetRepoTagRequest {
	s.Tag = &v
	return s
}

type GetRepoTagResponseBody struct {
	// The ID of the image.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The size of the image. Unit: Bytes.
	//
	// example:
	//
	// 67bfbcc12b67936ec7f867927817cbb071832b873dbcaed312a1930ba5f1****
	Digest *string `json:"Digest,omitempty" xml:"Digest,omitempty"`
	// crr-tquyps22md8p****
	//
	// example:
	//
	// 1572839125000
	ImageCreate *int64 `json:"ImageCreate,omitempty" xml:"ImageCreate,omitempty"`
	// example:
	//
	// 45023655bf39c382e26a8607d057c27871dee163c1ecf48cc1ebf2a1****
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The number of milliseconds that have elapsed since the image was last updated.
	//
	// example:
	//
	// 27107966
	ImageSize *int64 `json:"ImageSize,omitempty" xml:"ImageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 1572875608000
	ImageUpdate *int64 `json:"ImageUpdate,omitempty" xml:"ImageUpdate,omitempty"`
	// The status of the image. Valid values:
	//
	// 	- `NORMAL`: The image is normal.
	//
	// 	- `DELETING`: The image is being deleted.
	//
	// example:
	//
	// true
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// 1.0
	//
	// example:
	//
	// 031572FA-7D8F-4C05-B790-1071E0E05DE6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// NORMAL
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The version of the repository.
	//
	// example:
	//
	// 1.0
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
}

func (s GetRepoTagResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetRepoTagResponseBody) GoString() string {
	return s.String()
}

func (s *GetRepoTagResponseBody) SetCode(v string) *GetRepoTagResponseBody {
	s.Code = &v
	return s
}

func (s *GetRepoTagResponseBody) SetDigest(v string) *GetRepoTagResponseBody {
	s.Digest = &v
	return s
}

func (s *GetRepoTagResponseBody) SetImageCreate(v int64) *GetRepoTagResponseBody {
	s.ImageCreate = &v
	return s
}

func (s *GetRepoTagResponseBody) SetImageId(v string) *GetRepoTagResponseBody {
	s.ImageId = &v
	return s
}

func (s *GetRepoTagResponseBody) SetImageSize(v int64) *GetRepoTagResponseBody {
	s.ImageSize = &v
	return s
}

func (s *GetRepoTagResponseBody) SetImageUpdate(v int64) *GetRepoTagResponseBody {
	s.ImageUpdate = &v
	return s
}

func (s *GetRepoTagResponseBody) SetIsSuccess(v bool) *GetRepoTagResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *GetRepoTagResponseBody) SetRequestId(v string) *GetRepoTagResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRepoTagResponseBody) SetStatus(v string) *GetRepoTagResponseBody {
	s.Status = &v
	return s
}

func (s *GetRepoTagResponseBody) SetTag(v string) *GetRepoTagResponseBody {
	s.Tag = &v
	return s
}

type GetRepoTagResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetRepoTagResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetRepoTagResponse) String() string {
	return tea.Prettify(s)
}

func (s GetRepoTagResponse) GoString() string {
	return s.String()
}

func (s *GetRepoTagResponse) SetHeaders(v map[string]*string) *GetRepoTagResponse {
	s.Headers = v
	return s
}

func (s *GetRepoTagResponse) SetStatusCode(v int32) *GetRepoTagResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRepoTagResponse) SetBody(v *GetRepoTagResponseBody) *GetRepoTagResponse {
	s.Body = v
	return s
}

type GetRepoTagScanStatusRequest struct {
	// The image digest.
	//
	// example:
	//
	// 67bfbcc12b67936ec7f867927817cbb071832b873dbcaed312a1930ba5f1d529
	Digest *string `json:"Digest,omitempty" xml:"Digest,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cri-2j88dtld8yel****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the image repository.
	//
	// example:
	//
	// crr-uf082u9dg8do****
	RepoId *string `json:"RepoId,omitempty" xml:"RepoId,omitempty"`
	// The ID of the image scan task.
	//
	// example:
	//
	// 838152F9-F725-5A52-A344-8972D65AC045
	ScanTaskId *string `json:"ScanTaskId,omitempty" xml:"ScanTaskId,omitempty"`
	ScanType   *string `json:"ScanType,omitempty" xml:"ScanType,omitempty"`
	// The image tag.
	//
	// example:
	//
	// 1
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
}

func (s GetRepoTagScanStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s GetRepoTagScanStatusRequest) GoString() string {
	return s.String()
}

func (s *GetRepoTagScanStatusRequest) SetDigest(v string) *GetRepoTagScanStatusRequest {
	s.Digest = &v
	return s
}

func (s *GetRepoTagScanStatusRequest) SetInstanceId(v string) *GetRepoTagScanStatusRequest {
	s.InstanceId = &v
	return s
}

func (s *GetRepoTagScanStatusRequest) SetRepoId(v string) *GetRepoTagScanStatusRequest {
	s.RepoId = &v
	return s
}

func (s *GetRepoTagScanStatusRequest) SetScanTaskId(v string) *GetRepoTagScanStatusRequest {
	s.ScanTaskId = &v
	return s
}

func (s *GetRepoTagScanStatusRequest) SetScanType(v string) *GetRepoTagScanStatusRequest {
	s.ScanType = &v
	return s
}

func (s *GetRepoTagScanStatusRequest) SetTag(v string) *GetRepoTagScanStatusRequest {
	s.Tag = &v
	return s
}

type GetRepoTagScanStatusResponseBody struct {
	// The HTTP status code.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// 	- `true`: The request is successful.
	//
	// 	- `false`: The request fails.
	//
	// example:
	//
	// true
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The request ID.
	//
	// example:
	//
	// BC648259-91A7-4502-BED3-EDF64361FA83
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The type of the scanning engine.
	//
	// 	- `ACR_SCAN_SERVICE`: Trivy scan engine provided by Container Registry
	//
	// 	- `SAS_SCAN_SERVICE`: Security Center scan engine
	//
	// example:
	//
	// ACR_SCAN_SERVICE
	ScanService *string `json:"ScanService,omitempty" xml:"ScanService,omitempty"`
	// The scanning status of the image tag. Valid values:
	//
	// 	- `SCANNING`: The image tag is being scanned.
	//
	// 	- `COMPLETE`: The scanning of the image tag is complete.
	//
	// 	- `FAILED`: The image tag failed to be scanned.
	//
	// 	- `RETRYING`: The system is retrying to scan the image tag.
	//
	// example:
	//
	// COMPLETE
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetRepoTagScanStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetRepoTagScanStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetRepoTagScanStatusResponseBody) SetCode(v string) *GetRepoTagScanStatusResponseBody {
	s.Code = &v
	return s
}

func (s *GetRepoTagScanStatusResponseBody) SetIsSuccess(v bool) *GetRepoTagScanStatusResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *GetRepoTagScanStatusResponseBody) SetRequestId(v string) *GetRepoTagScanStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRepoTagScanStatusResponseBody) SetScanService(v string) *GetRepoTagScanStatusResponseBody {
	s.ScanService = &v
	return s
}

func (s *GetRepoTagScanStatusResponseBody) SetStatus(v string) *GetRepoTagScanStatusResponseBody {
	s.Status = &v
	return s
}

type GetRepoTagScanStatusResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetRepoTagScanStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetRepoTagScanStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s GetRepoTagScanStatusResponse) GoString() string {
	return s.String()
}

func (s *GetRepoTagScanStatusResponse) SetHeaders(v map[string]*string) *GetRepoTagScanStatusResponse {
	s.Headers = v
	return s
}

func (s *GetRepoTagScanStatusResponse) SetStatusCode(v int32) *GetRepoTagScanStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRepoTagScanStatusResponse) SetBody(v *GetRepoTagScanStatusResponseBody) *GetRepoTagScanStatusResponse {
	s.Body = v
	return s
}

type GetRepoTagScanSummaryRequest struct {
	// The number of unknown-severity vulnerabilities.
	//
	// example:
	//
	// sha256:c9f370a4eb1c00d0b0d7212a0a9fa4a7697756c90f0f680afaf9737a25725f4c
	Digest *string `json:"Digest,omitempty" xml:"Digest,omitempty"`
	// The ID of the image repository.
	//
	// This parameter is required.
	//
	// example:
	//
	// cri-2j88dtld8yel****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the image tag.
	//
	// example:
	//
	// crr-c2i5yk6h6pu9d5o8
	RepoId *string `json:"RepoId,omitempty" xml:"RepoId,omitempty"`
	// The digest of the image.
	//
	// example:
	//
	// 47A3E5A3-6AD4-5F02-93B8-59F778AE25D4
	ScanTaskId *string `json:"ScanTaskId,omitempty" xml:"ScanTaskId,omitempty"`
	// The ID of the security scan task.
	//
	// example:
	//
	// 1
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
}

func (s GetRepoTagScanSummaryRequest) String() string {
	return tea.Prettify(s)
}

func (s GetRepoTagScanSummaryRequest) GoString() string {
	return s.String()
}

func (s *GetRepoTagScanSummaryRequest) SetDigest(v string) *GetRepoTagScanSummaryRequest {
	s.Digest = &v
	return s
}

func (s *GetRepoTagScanSummaryRequest) SetInstanceId(v string) *GetRepoTagScanSummaryRequest {
	s.InstanceId = &v
	return s
}

func (s *GetRepoTagScanSummaryRequest) SetRepoId(v string) *GetRepoTagScanSummaryRequest {
	s.RepoId = &v
	return s
}

func (s *GetRepoTagScanSummaryRequest) SetScanTaskId(v string) *GetRepoTagScanSummaryRequest {
	s.ScanTaskId = &v
	return s
}

func (s *GetRepoTagScanSummaryRequest) SetTag(v string) *GetRepoTagScanSummaryRequest {
	s.Tag = &v
	return s
}

type GetRepoTagScanSummaryResponseBody struct {
	// The number of medium-severity vulnerabilities.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The number of low-severity vulnerabilities.
	//
	// example:
	//
	// 22
	HighSeverity *int32 `json:"HighSeverity,omitempty" xml:"HighSeverity,omitempty"`
	// The number of high-severity vulnerabilities.
	//
	// example:
	//
	// true
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// example:
	//
	// 89
	LowSeverity *int32 `json:"LowSeverity,omitempty" xml:"LowSeverity,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// 	- `true`: The request is successful.
	//
	// 	- `false`: The request fails.
	//
	// example:
	//
	// 81
	MediumSeverity *int32 `json:"MediumSeverity,omitempty" xml:"MediumSeverity,omitempty"`
	// The total number of vulnerabilities detected on images.
	//
	// example:
	//
	// BC648259-91A7-4502-BED3-EDF64361FA83
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The return value.
	//
	// example:
	//
	// 196
	TotalSeverity *int32 `json:"TotalSeverity,omitempty" xml:"TotalSeverity,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 4
	UnknownSeverity *int32 `json:"UnknownSeverity,omitempty" xml:"UnknownSeverity,omitempty"`
}

func (s GetRepoTagScanSummaryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetRepoTagScanSummaryResponseBody) GoString() string {
	return s.String()
}

func (s *GetRepoTagScanSummaryResponseBody) SetCode(v string) *GetRepoTagScanSummaryResponseBody {
	s.Code = &v
	return s
}

func (s *GetRepoTagScanSummaryResponseBody) SetHighSeverity(v int32) *GetRepoTagScanSummaryResponseBody {
	s.HighSeverity = &v
	return s
}

func (s *GetRepoTagScanSummaryResponseBody) SetIsSuccess(v bool) *GetRepoTagScanSummaryResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *GetRepoTagScanSummaryResponseBody) SetLowSeverity(v int32) *GetRepoTagScanSummaryResponseBody {
	s.LowSeverity = &v
	return s
}

func (s *GetRepoTagScanSummaryResponseBody) SetMediumSeverity(v int32) *GetRepoTagScanSummaryResponseBody {
	s.MediumSeverity = &v
	return s
}

func (s *GetRepoTagScanSummaryResponseBody) SetRequestId(v string) *GetRepoTagScanSummaryResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRepoTagScanSummaryResponseBody) SetTotalSeverity(v int32) *GetRepoTagScanSummaryResponseBody {
	s.TotalSeverity = &v
	return s
}

func (s *GetRepoTagScanSummaryResponseBody) SetUnknownSeverity(v int32) *GetRepoTagScanSummaryResponseBody {
	s.UnknownSeverity = &v
	return s
}

type GetRepoTagScanSummaryResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetRepoTagScanSummaryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetRepoTagScanSummaryResponse) String() string {
	return tea.Prettify(s)
}

func (s GetRepoTagScanSummaryResponse) GoString() string {
	return s.String()
}

func (s *GetRepoTagScanSummaryResponse) SetHeaders(v map[string]*string) *GetRepoTagScanSummaryResponse {
	s.Headers = v
	return s
}

func (s *GetRepoTagScanSummaryResponse) SetStatusCode(v int32) *GetRepoTagScanSummaryResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRepoTagScanSummaryResponse) SetBody(v *GetRepoTagScanSummaryResponseBody) *GetRepoTagScanSummaryResponse {
	s.Body = v
	return s
}

type GetRepositoryRequest struct {
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cri-kmsiwlxxdcva****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the repository.
	//
	// example:
	//
	// crr-03cuozrsqhkw****
	RepoId *string `json:"RepoId,omitempty" xml:"RepoId,omitempty"`
	// The name of the repository.
	//
	// example:
	//
	// test
	RepoName *string `json:"RepoName,omitempty" xml:"RepoName,omitempty"`
	// The name of the namespace to which the repository belongs.
	//
	// example:
	//
	// test
	RepoNamespaceName *string `json:"RepoNamespaceName,omitempty" xml:"RepoNamespaceName,omitempty"`
}

func (s GetRepositoryRequest) String() string {
	return tea.Prettify(s)
}

func (s GetRepositoryRequest) GoString() string {
	return s.String()
}

func (s *GetRepositoryRequest) SetInstanceId(v string) *GetRepositoryRequest {
	s.InstanceId = &v
	return s
}

func (s *GetRepositoryRequest) SetRepoId(v string) *GetRepositoryRequest {
	s.RepoId = &v
	return s
}

func (s *GetRepositoryRequest) SetRepoName(v string) *GetRepositoryRequest {
	s.RepoName = &v
	return s
}

func (s *GetRepositoryRequest) SetRepoNamespaceName(v string) *GetRepositoryRequest {
	s.RepoNamespaceName = &v
	return s
}

type GetRepositoryResponseBody struct {
	// The return value.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The time when the repository was created.
	//
	// example:
	//
	// 1570759546000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The details of the repository.
	//
	// example:
	//
	// test
	Detail *string `json:"Detail,omitempty" xml:"Detail,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// cri-kmsiwlxxdcva****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The time when the repository was last modified.
	//
	// example:
	//
	// 1570759546100
	ModifiedTime *int64 `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// Indicates how the repository was created. Valid values:
	//
	// 	- `MANUAL`: The repository was manually created.
	//
	// 	- `AUTO`: The repository was automatically created.
	//
	// example:
	//
	// MANUAL
	RepoBuildType *string `json:"RepoBuildType,omitempty" xml:"RepoBuildType,omitempty"`
	// The ID of the repository.
	//
	// example:
	//
	// crr-l5eoubonp0l****
	RepoId *string `json:"RepoId,omitempty" xml:"RepoId,omitempty"`
	// The name of the repository.
	//
	// example:
	//
	// test
	RepoName *string `json:"RepoName,omitempty" xml:"RepoName,omitempty"`
	// The name of the namespace.
	//
	// example:
	//
	// test
	RepoNamespaceName *string `json:"RepoNamespaceName,omitempty" xml:"RepoNamespaceName,omitempty"`
	// The status of the repository.
	//
	// example:
	//
	// NORMAL
	RepoStatus *string `json:"RepoStatus,omitempty" xml:"RepoStatus,omitempty"`
	// The type of the repository. Valid values:
	//
	// 	- `PUBLIC`: public repository.
	//
	// 	- `PRIVATE`: private repository.
	//
	// example:
	//
	// PRIVATE
	RepoType *string `json:"RepoType,omitempty" xml:"RepoType,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 915E6734-3E50-4640-8DBA-126D2D94DE29
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfmv36i4is****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The summary of the repository.
	//
	// example:
	//
	// Automatically created repository
	Summary *string `json:"Summary,omitempty" xml:"Summary,omitempty"`
	// Indicates whether the feature of image tag immutability is enabled. Valid values:
	//
	// 	- `true`: The feature of image tag immutability is enabled.
	//
	// 	- `false`: The feature of image tag immutability is disabled.
	//
	// example:
	//
	// true
	TagImmutability *bool `json:"TagImmutability,omitempty" xml:"TagImmutability,omitempty"`
}

func (s GetRepositoryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetRepositoryResponseBody) GoString() string {
	return s.String()
}

func (s *GetRepositoryResponseBody) SetCode(v string) *GetRepositoryResponseBody {
	s.Code = &v
	return s
}

func (s *GetRepositoryResponseBody) SetCreateTime(v int64) *GetRepositoryResponseBody {
	s.CreateTime = &v
	return s
}

func (s *GetRepositoryResponseBody) SetDetail(v string) *GetRepositoryResponseBody {
	s.Detail = &v
	return s
}

func (s *GetRepositoryResponseBody) SetInstanceId(v string) *GetRepositoryResponseBody {
	s.InstanceId = &v
	return s
}

func (s *GetRepositoryResponseBody) SetIsSuccess(v bool) *GetRepositoryResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *GetRepositoryResponseBody) SetModifiedTime(v int64) *GetRepositoryResponseBody {
	s.ModifiedTime = &v
	return s
}

func (s *GetRepositoryResponseBody) SetRepoBuildType(v string) *GetRepositoryResponseBody {
	s.RepoBuildType = &v
	return s
}

func (s *GetRepositoryResponseBody) SetRepoId(v string) *GetRepositoryResponseBody {
	s.RepoId = &v
	return s
}

func (s *GetRepositoryResponseBody) SetRepoName(v string) *GetRepositoryResponseBody {
	s.RepoName = &v
	return s
}

func (s *GetRepositoryResponseBody) SetRepoNamespaceName(v string) *GetRepositoryResponseBody {
	s.RepoNamespaceName = &v
	return s
}

func (s *GetRepositoryResponseBody) SetRepoStatus(v string) *GetRepositoryResponseBody {
	s.RepoStatus = &v
	return s
}

func (s *GetRepositoryResponseBody) SetRepoType(v string) *GetRepositoryResponseBody {
	s.RepoType = &v
	return s
}

func (s *GetRepositoryResponseBody) SetRequestId(v string) *GetRepositoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRepositoryResponseBody) SetResourceGroupId(v string) *GetRepositoryResponseBody {
	s.ResourceGroupId = &v
	return s
}

func (s *GetRepositoryResponseBody) SetSummary(v string) *GetRepositoryResponseBody {
	s.Summary = &v
	return s
}

func (s *GetRepositoryResponseBody) SetTagImmutability(v bool) *GetRepositoryResponseBody {
	s.TagImmutability = &v
	return s
}

type GetRepositoryResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetRepositoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetRepositoryResponse) String() string {
	return tea.Prettify(s)
}

func (s GetRepositoryResponse) GoString() string {
	return s.String()
}

func (s *GetRepositoryResponse) SetHeaders(v map[string]*string) *GetRepositoryResponse {
	s.Headers = v
	return s
}

func (s *GetRepositoryResponse) SetStatusCode(v int32) *GetRepositoryResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRepositoryResponse) SetBody(v *GetRepositoryResponseBody) *GetRepositoryResponse {
	s.Body = v
	return s
}

type ListArtifactBuildTaskLogRequest struct {
	// The ID of the artifact build task.
	//
	// This parameter is required.
	//
	// example:
	//
	// i2a-1yu****
	BuildTaskId *string `json:"BuildTaskId,omitempty" xml:"BuildTaskId,omitempty"`
	// The ID of the Container Registry instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cri-shac42yvqzvq****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The number of the page to return.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	Page *int32 `json:"Page,omitempty" xml:"Page,omitempty"`
	// The number of entries per page. Maximum value: 100. If you specify a value greater than 100 for this parameter, the system reports a parameter error or uses 100 as the maximum value.
	//
	// This parameter is required.
	//
	// example:
	//
	// 100
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListArtifactBuildTaskLogRequest) String() string {
	return tea.Prettify(s)
}

func (s ListArtifactBuildTaskLogRequest) GoString() string {
	return s.String()
}

func (s *ListArtifactBuildTaskLogRequest) SetBuildTaskId(v string) *ListArtifactBuildTaskLogRequest {
	s.BuildTaskId = &v
	return s
}

func (s *ListArtifactBuildTaskLogRequest) SetInstanceId(v string) *ListArtifactBuildTaskLogRequest {
	s.InstanceId = &v
	return s
}

func (s *ListArtifactBuildTaskLogRequest) SetPage(v int32) *ListArtifactBuildTaskLogRequest {
	s.Page = &v
	return s
}

func (s *ListArtifactBuildTaskLogRequest) SetPageSize(v int32) *ListArtifactBuildTaskLogRequest {
	s.PageSize = &v
	return s
}

type ListArtifactBuildTaskLogResponseBody struct {
	// The log entries of the artifact build task.
	BuildTaskLogs []*ListArtifactBuildTaskLogResponseBodyBuildTaskLogs `json:"BuildTaskLogs,omitempty" xml:"BuildTaskLogs,omitempty" type:"Repeated"`
	// The response code.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the API request is successful. Valid values:
	//
	// 	- `true`: successful
	//
	// 	- `false`: failed
	//
	// example:
	//
	// true
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The request ID.
	//
	// example:
	//
	// C4C7DD0C-C9D6-437A-A7EE-121EFD70D002
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of log entries.
	//
	// example:
	//
	// 200
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListArtifactBuildTaskLogResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListArtifactBuildTaskLogResponseBody) GoString() string {
	return s.String()
}

func (s *ListArtifactBuildTaskLogResponseBody) SetBuildTaskLogs(v []*ListArtifactBuildTaskLogResponseBodyBuildTaskLogs) *ListArtifactBuildTaskLogResponseBody {
	s.BuildTaskLogs = v
	return s
}

func (s *ListArtifactBuildTaskLogResponseBody) SetCode(v string) *ListArtifactBuildTaskLogResponseBody {
	s.Code = &v
	return s
}

func (s *ListArtifactBuildTaskLogResponseBody) SetIsSuccess(v bool) *ListArtifactBuildTaskLogResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *ListArtifactBuildTaskLogResponseBody) SetRequestId(v string) *ListArtifactBuildTaskLogResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListArtifactBuildTaskLogResponseBody) SetTotalCount(v int32) *ListArtifactBuildTaskLogResponseBody {
	s.TotalCount = &v
	return s
}

type ListArtifactBuildTaskLogResponseBodyBuildTaskLogs struct {
	// The row number of the log entry.
	//
	// example:
	//
	// 1
	LineNumber *int32 `json:"LineNumber,omitempty" xml:"LineNumber,omitempty"`
	// The log data.
	//
	// example:
	//
	// Start Build
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s ListArtifactBuildTaskLogResponseBodyBuildTaskLogs) String() string {
	return tea.Prettify(s)
}

func (s ListArtifactBuildTaskLogResponseBodyBuildTaskLogs) GoString() string {
	return s.String()
}

func (s *ListArtifactBuildTaskLogResponseBodyBuildTaskLogs) SetLineNumber(v int32) *ListArtifactBuildTaskLogResponseBodyBuildTaskLogs {
	s.LineNumber = &v
	return s
}

func (s *ListArtifactBuildTaskLogResponseBodyBuildTaskLogs) SetMessage(v string) *ListArtifactBuildTaskLogResponseBodyBuildTaskLogs {
	s.Message = &v
	return s
}

type ListArtifactBuildTaskLogResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListArtifactBuildTaskLogResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListArtifactBuildTaskLogResponse) String() string {
	return tea.Prettify(s)
}

func (s ListArtifactBuildTaskLogResponse) GoString() string {
	return s.String()
}

func (s *ListArtifactBuildTaskLogResponse) SetHeaders(v map[string]*string) *ListArtifactBuildTaskLogResponse {
	s.Headers = v
	return s
}

func (s *ListArtifactBuildTaskLogResponse) SetStatusCode(v int32) *ListArtifactBuildTaskLogResponse {
	s.StatusCode = &v
	return s
}

func (s *ListArtifactBuildTaskLogResponse) SetBody(v *ListArtifactBuildTaskLogResponseBody) *ListArtifactBuildTaskLogResponse {
	s.Body = v
	return s
}

type ListArtifactLifecycleRuleRequest struct {
	// Specifies whether to enable lifecycle management for the artifact.
	//
	// example:
	//
	// true
	EnableDeleteTag *bool `json:"EnableDeleteTag,omitempty" xml:"EnableDeleteTag,omitempty"`
	// The ID of the Container Registry Enterprise Edition instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cri-eztul9ucz76q****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page. Maximum value: 100. If you specify a value greater than 100 for this parameter, the system reports a parameter error or uses 100 as the maximum value.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListArtifactLifecycleRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s ListArtifactLifecycleRuleRequest) GoString() string {
	return s.String()
}

func (s *ListArtifactLifecycleRuleRequest) SetEnableDeleteTag(v bool) *ListArtifactLifecycleRuleRequest {
	s.EnableDeleteTag = &v
	return s
}

func (s *ListArtifactLifecycleRuleRequest) SetInstanceId(v string) *ListArtifactLifecycleRuleRequest {
	s.InstanceId = &v
	return s
}

func (s *ListArtifactLifecycleRuleRequest) SetPageNo(v int32) *ListArtifactLifecycleRuleRequest {
	s.PageNo = &v
	return s
}

func (s *ListArtifactLifecycleRuleRequest) SetPageSize(v int32) *ListArtifactLifecycleRuleRequest {
	s.PageSize = &v
	return s
}

type ListArtifactLifecycleRuleResponseBody struct {
	// The return value.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// 	- `true`: The request is successful.
	//
	// 	- `false`: The request fails.
	//
	// example:
	//
	// true
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// F92D82F9-A4C4-5A4A-97B9-E495BF1B****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// _
	Rules []*ListArtifactLifecycleRuleResponseBodyRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
	// The total number of entries returned.
	//
	// example:
	//
	// 39
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListArtifactLifecycleRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListArtifactLifecycleRuleResponseBody) GoString() string {
	return s.String()
}

func (s *ListArtifactLifecycleRuleResponseBody) SetCode(v string) *ListArtifactLifecycleRuleResponseBody {
	s.Code = &v
	return s
}

func (s *ListArtifactLifecycleRuleResponseBody) SetIsSuccess(v bool) *ListArtifactLifecycleRuleResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *ListArtifactLifecycleRuleResponseBody) SetPageNo(v int32) *ListArtifactLifecycleRuleResponseBody {
	s.PageNo = &v
	return s
}

func (s *ListArtifactLifecycleRuleResponseBody) SetPageSize(v int32) *ListArtifactLifecycleRuleResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListArtifactLifecycleRuleResponseBody) SetRequestId(v string) *ListArtifactLifecycleRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListArtifactLifecycleRuleResponseBody) SetRules(v []*ListArtifactLifecycleRuleResponseBodyRules) *ListArtifactLifecycleRuleResponseBody {
	s.Rules = v
	return s
}

func (s *ListArtifactLifecycleRuleResponseBody) SetTotalCount(v int32) *ListArtifactLifecycleRuleResponseBody {
	s.TotalCount = &v
	return s
}

type ListArtifactLifecycleRuleResponseBodyRules struct {
	// Indicates whether the lifecycle management rule is automatically executed.
	//
	// example:
	//
	// false
	Auto *bool `json:"Auto,omitempty" xml:"Auto,omitempty"`
	// The time when the lifecycle management rule was created.
	//
	// example:
	//
	// 1638187989000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// Indicates whether lifecycle management is enabled for the artifact.
	//
	// example:
	//
	// true
	EnableDeleteTag *bool `json:"EnableDeleteTag,omitempty" xml:"EnableDeleteTag,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// cri-brlg4cbj2yl****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The time when the lifecycle management rule was last modified.
	//
	// example:
	//
	// 1678341923385
	ModifiedTime *int64 `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// The name of the namespace.
	//
	// example:
	//
	// test-ns
	NamespaceName *string `json:"NamespaceName,omitempty" xml:"NamespaceName,omitempty"`
	// The time when the lifecycle management rule is next executed.
	//
	// example:
	//
	// 1638187989000
	NextTime *int64 `json:"NextTime,omitempty" xml:"NextTime,omitempty"`
	// The name of the image repository.
	//
	// example:
	//
	// test_1
	RepoName *string `json:"RepoName,omitempty" xml:"RepoName,omitempty"`
	// The number of retained images.
	//
	// example:
	//
	// 30
	RetentionTagCount *int64 `json:"RetentionTagCount,omitempty" xml:"RetentionTagCount,omitempty"`
	// The rule ID.
	//
	// example:
	//
	// cralr-yqx1q5sir6d****
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// The execution cycle of the lifecycle management rule.
	//
	// example:
	//
	// WEEK
	ScheduleTime *string `json:"ScheduleTime,omitempty" xml:"ScheduleTime,omitempty"`
	// The deletion scope of artifacts.
	//
	// example:
	//
	// INSTANCE
	Scope *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
	// The regular expression that indicates which image tags are retained.
	//
	// example:
	//
	// .*-alpine
	TagRegexp *string `json:"TagRegexp,omitempty" xml:"TagRegexp,omitempty"`
}

func (s ListArtifactLifecycleRuleResponseBodyRules) String() string {
	return tea.Prettify(s)
}

func (s ListArtifactLifecycleRuleResponseBodyRules) GoString() string {
	return s.String()
}

func (s *ListArtifactLifecycleRuleResponseBodyRules) SetAuto(v bool) *ListArtifactLifecycleRuleResponseBodyRules {
	s.Auto = &v
	return s
}

func (s *ListArtifactLifecycleRuleResponseBodyRules) SetCreateTime(v int64) *ListArtifactLifecycleRuleResponseBodyRules {
	s.CreateTime = &v
	return s
}

func (s *ListArtifactLifecycleRuleResponseBodyRules) SetEnableDeleteTag(v bool) *ListArtifactLifecycleRuleResponseBodyRules {
	s.EnableDeleteTag = &v
	return s
}

func (s *ListArtifactLifecycleRuleResponseBodyRules) SetInstanceId(v string) *ListArtifactLifecycleRuleResponseBodyRules {
	s.InstanceId = &v
	return s
}

func (s *ListArtifactLifecycleRuleResponseBodyRules) SetModifiedTime(v int64) *ListArtifactLifecycleRuleResponseBodyRules {
	s.ModifiedTime = &v
	return s
}

func (s *ListArtifactLifecycleRuleResponseBodyRules) SetNamespaceName(v string) *ListArtifactLifecycleRuleResponseBodyRules {
	s.NamespaceName = &v
	return s
}

func (s *ListArtifactLifecycleRuleResponseBodyRules) SetNextTime(v int64) *ListArtifactLifecycleRuleResponseBodyRules {
	s.NextTime = &v
	return s
}

func (s *ListArtifactLifecycleRuleResponseBodyRules) SetRepoName(v string) *ListArtifactLifecycleRuleResponseBodyRules {
	s.RepoName = &v
	return s
}

func (s *ListArtifactLifecycleRuleResponseBodyRules) SetRetentionTagCount(v int64) *ListArtifactLifecycleRuleResponseBodyRules {
	s.RetentionTagCount = &v
	return s
}

func (s *ListArtifactLifecycleRuleResponseBodyRules) SetRuleId(v string) *ListArtifactLifecycleRuleResponseBodyRules {
	s.RuleId = &v
	return s
}

func (s *ListArtifactLifecycleRuleResponseBodyRules) SetScheduleTime(v string) *ListArtifactLifecycleRuleResponseBodyRules {
	s.ScheduleTime = &v
	return s
}

func (s *ListArtifactLifecycleRuleResponseBodyRules) SetScope(v string) *ListArtifactLifecycleRuleResponseBodyRules {
	s.Scope = &v
	return s
}

func (s *ListArtifactLifecycleRuleResponseBodyRules) SetTagRegexp(v string) *ListArtifactLifecycleRuleResponseBodyRules {
	s.TagRegexp = &v
	return s
}

type ListArtifactLifecycleRuleResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListArtifactLifecycleRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListArtifactLifecycleRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s ListArtifactLifecycleRuleResponse) GoString() string {
	return s.String()
}

func (s *ListArtifactLifecycleRuleResponse) SetHeaders(v map[string]*string) *ListArtifactLifecycleRuleResponse {
	s.Headers = v
	return s
}

func (s *ListArtifactLifecycleRuleResponse) SetStatusCode(v int32) *ListArtifactLifecycleRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *ListArtifactLifecycleRuleResponse) SetBody(v *ListArtifactLifecycleRuleResponseBody) *ListArtifactLifecycleRuleResponse {
	s.Body = v
	return s
}

type ListArtifactSubscriptionRuleRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cri-c0o11woew0k****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page. Maximum value: 100. If you specify a value greater than 100 for this parameter, the system reports a parameter error or uses 100 as the maximum value.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListArtifactSubscriptionRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s ListArtifactSubscriptionRuleRequest) GoString() string {
	return s.String()
}

func (s *ListArtifactSubscriptionRuleRequest) SetInstanceId(v string) *ListArtifactSubscriptionRuleRequest {
	s.InstanceId = &v
	return s
}

func (s *ListArtifactSubscriptionRuleRequest) SetPageNo(v int32) *ListArtifactSubscriptionRuleRequest {
	s.PageNo = &v
	return s
}

func (s *ListArtifactSubscriptionRuleRequest) SetPageSize(v int32) *ListArtifactSubscriptionRuleRequest {
	s.PageSize = &v
	return s
}

type ListArtifactSubscriptionRuleResponseBody struct {
	// The return value.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// 	- `true`
	//
	// 	- `false`
	//
	// example:
	//
	// true
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 07FC5654-C82A-59FA-A9D1-78B4EE443F86
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The queried artifact subscription rules.
	Rules []*ListArtifactSubscriptionRuleResponseBodyRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
	// The total number of entries returned.
	//
	// example:
	//
	// 13
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListArtifactSubscriptionRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListArtifactSubscriptionRuleResponseBody) GoString() string {
	return s.String()
}

func (s *ListArtifactSubscriptionRuleResponseBody) SetCode(v string) *ListArtifactSubscriptionRuleResponseBody {
	s.Code = &v
	return s
}

func (s *ListArtifactSubscriptionRuleResponseBody) SetIsSuccess(v bool) *ListArtifactSubscriptionRuleResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *ListArtifactSubscriptionRuleResponseBody) SetPageNo(v int32) *ListArtifactSubscriptionRuleResponseBody {
	s.PageNo = &v
	return s
}

func (s *ListArtifactSubscriptionRuleResponseBody) SetPageSize(v int32) *ListArtifactSubscriptionRuleResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListArtifactSubscriptionRuleResponseBody) SetRequestId(v string) *ListArtifactSubscriptionRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListArtifactSubscriptionRuleResponseBody) SetRules(v []*ListArtifactSubscriptionRuleResponseBodyRules) *ListArtifactSubscriptionRuleResponseBody {
	s.Rules = v
	return s
}

func (s *ListArtifactSubscriptionRuleResponseBody) SetTotalCount(v int32) *ListArtifactSubscriptionRuleResponseBody {
	s.TotalCount = &v
	return s
}

type ListArtifactSubscriptionRuleResponseBodyRules struct {
	// Indicates whether an acceleration link is enabled for image subscription. The subscription acceleration feature is in public preview. The feature is optimized based on scheduling policies and network links to accelerate image subscription.
	//
	// example:
	//
	// true
	Accelerate *bool `json:"Accelerate,omitempty" xml:"Accelerate,omitempty"`
	// The time when the subscription rule was created.
	//
	// example:
	//
	// 1638187989000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// cri-brlg4cbj2yl****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The time when the subscription rule was modified.
	//
	// example:
	//
	// 1678341923385
	ModifiedTime *int64 `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// The name of the source namespace.
	//
	// example:
	//
	// test-ns
	NamespaceName *string `json:"NamespaceName,omitempty" xml:"NamespaceName,omitempty"`
	// Indicates whether the original image is overwritten.
	//
	// example:
	//
	// true
	Override *bool `json:"Override,omitempty" xml:"Override,omitempty"`
	// The operating system and architecture. If the source repository contains a multi-arch image, only the images with the specified operating system and architecture are subscribed to the destination repository of the Enterprise Edition instance.
	Platform []*string `json:"Platform,omitempty" xml:"Platform,omitempty" type:"Repeated"`
	// The name of the source repository.
	//
	// example:
	//
	// test-repo
	RepoName *string `json:"RepoName,omitempty" xml:"RepoName,omitempty"`
	// The rule ID.
	//
	// example:
	//
	// crasr-mdbpung4i1rm****
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// The source namespace.
	//
	// example:
	//
	// library
	SourceNamespaceName *string `json:"SourceNamespaceName,omitempty" xml:"SourceNamespaceName,omitempty"`
	// The source of the artifact.
	//
	// Valid values:
	//
	// 	- DOCKER_HUB: Docker Hub
	//
	// 	- GCR: GCR
	//
	// 	- QUAY: Quay.io
	//
	// example:
	//
	// DOCKER_HUB
	SourceProvider *string `json:"SourceProvider,omitempty" xml:"SourceProvider,omitempty"`
	// The source repository.
	//
	// example:
	//
	// nginx
	SourceRepoName *string `json:"SourceRepoName,omitempty" xml:"SourceRepoName,omitempty"`
	// The number of subscribed images.
	//
	// example:
	//
	// 1
	TagCount *int64 `json:"TagCount,omitempty" xml:"TagCount,omitempty"`
	// The image tag in the subscription source repository. Regular expressions are supported.
	//
	// example:
	//
	// release.*
	TagRegexp *string `json:"TagRegexp,omitempty" xml:"TagRegexp,omitempty"`
}

func (s ListArtifactSubscriptionRuleResponseBodyRules) String() string {
	return tea.Prettify(s)
}

func (s ListArtifactSubscriptionRuleResponseBodyRules) GoString() string {
	return s.String()
}

func (s *ListArtifactSubscriptionRuleResponseBodyRules) SetAccelerate(v bool) *ListArtifactSubscriptionRuleResponseBodyRules {
	s.Accelerate = &v
	return s
}

func (s *ListArtifactSubscriptionRuleResponseBodyRules) SetCreateTime(v int64) *ListArtifactSubscriptionRuleResponseBodyRules {
	s.CreateTime = &v
	return s
}

func (s *ListArtifactSubscriptionRuleResponseBodyRules) SetInstanceId(v string) *ListArtifactSubscriptionRuleResponseBodyRules {
	s.InstanceId = &v
	return s
}

func (s *ListArtifactSubscriptionRuleResponseBodyRules) SetModifiedTime(v int64) *ListArtifactSubscriptionRuleResponseBodyRules {
	s.ModifiedTime = &v
	return s
}

func (s *ListArtifactSubscriptionRuleResponseBodyRules) SetNamespaceName(v string) *ListArtifactSubscriptionRuleResponseBodyRules {
	s.NamespaceName = &v
	return s
}

func (s *ListArtifactSubscriptionRuleResponseBodyRules) SetOverride(v bool) *ListArtifactSubscriptionRuleResponseBodyRules {
	s.Override = &v
	return s
}

func (s *ListArtifactSubscriptionRuleResponseBodyRules) SetPlatform(v []*string) *ListArtifactSubscriptionRuleResponseBodyRules {
	s.Platform = v
	return s
}

func (s *ListArtifactSubscriptionRuleResponseBodyRules) SetRepoName(v string) *ListArtifactSubscriptionRuleResponseBodyRules {
	s.RepoName = &v
	return s
}

func (s *ListArtifactSubscriptionRuleResponseBodyRules) SetRuleId(v string) *ListArtifactSubscriptionRuleResponseBodyRules {
	s.RuleId = &v
	return s
}

func (s *ListArtifactSubscriptionRuleResponseBodyRules) SetSourceNamespaceName(v string) *ListArtifactSubscriptionRuleResponseBodyRules {
	s.SourceNamespaceName = &v
	return s
}

func (s *ListArtifactSubscriptionRuleResponseBodyRules) SetSourceProvider(v string) *ListArtifactSubscriptionRuleResponseBodyRules {
	s.SourceProvider = &v
	return s
}

func (s *ListArtifactSubscriptionRuleResponseBodyRules) SetSourceRepoName(v string) *ListArtifactSubscriptionRuleResponseBodyRules {
	s.SourceRepoName = &v
	return s
}

func (s *ListArtifactSubscriptionRuleResponseBodyRules) SetTagCount(v int64) *ListArtifactSubscriptionRuleResponseBodyRules {
	s.TagCount = &v
	return s
}

func (s *ListArtifactSubscriptionRuleResponseBodyRules) SetTagRegexp(v string) *ListArtifactSubscriptionRuleResponseBodyRules {
	s.TagRegexp = &v
	return s
}

type ListArtifactSubscriptionRuleResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListArtifactSubscriptionRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListArtifactSubscriptionRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s ListArtifactSubscriptionRuleResponse) GoString() string {
	return s.String()
}

func (s *ListArtifactSubscriptionRuleResponse) SetHeaders(v map[string]*string) *ListArtifactSubscriptionRuleResponse {
	s.Headers = v
	return s
}

func (s *ListArtifactSubscriptionRuleResponse) SetStatusCode(v int32) *ListArtifactSubscriptionRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *ListArtifactSubscriptionRuleResponse) SetBody(v *ListArtifactSubscriptionRuleResponseBody) *ListArtifactSubscriptionRuleResponse {
	s.Body = v
	return s
}

type ListArtifactSubscriptionTaskRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cri-m9ob8792vm****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListArtifactSubscriptionTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s ListArtifactSubscriptionTaskRequest) GoString() string {
	return s.String()
}

func (s *ListArtifactSubscriptionTaskRequest) SetInstanceId(v string) *ListArtifactSubscriptionTaskRequest {
	s.InstanceId = &v
	return s
}

func (s *ListArtifactSubscriptionTaskRequest) SetPageNo(v int32) *ListArtifactSubscriptionTaskRequest {
	s.PageNo = &v
	return s
}

func (s *ListArtifactSubscriptionTaskRequest) SetPageSize(v int32) *ListArtifactSubscriptionTaskRequest {
	s.PageSize = &v
	return s
}

type ListArtifactSubscriptionTaskResponseBody struct {
	// The return value.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the API request is successful. Valid values:
	//
	// 	- `true`: The request is successful.
	//
	// 	- `false`: The request fails.
	//
	// example:
	//
	// true
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 81E7A039-A4EF-57D9-A100-88E5DCEF9D56
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The queried artifact subscription tasks.
	Tasks []*ListArtifactSubscriptionTaskResponseBodyTasks `json:"Tasks,omitempty" xml:"Tasks,omitempty" type:"Repeated"`
	// The total number of entries returned.
	//
	// example:
	//
	// 5
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListArtifactSubscriptionTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListArtifactSubscriptionTaskResponseBody) GoString() string {
	return s.String()
}

func (s *ListArtifactSubscriptionTaskResponseBody) SetCode(v string) *ListArtifactSubscriptionTaskResponseBody {
	s.Code = &v
	return s
}

func (s *ListArtifactSubscriptionTaskResponseBody) SetIsSuccess(v bool) *ListArtifactSubscriptionTaskResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *ListArtifactSubscriptionTaskResponseBody) SetPageNo(v int32) *ListArtifactSubscriptionTaskResponseBody {
	s.PageNo = &v
	return s
}

func (s *ListArtifactSubscriptionTaskResponseBody) SetPageSize(v int32) *ListArtifactSubscriptionTaskResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListArtifactSubscriptionTaskResponseBody) SetRequestId(v string) *ListArtifactSubscriptionTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListArtifactSubscriptionTaskResponseBody) SetTasks(v []*ListArtifactSubscriptionTaskResponseBodyTasks) *ListArtifactSubscriptionTaskResponseBody {
	s.Tasks = v
	return s
}

func (s *ListArtifactSubscriptionTaskResponseBody) SetTotalCount(v int32) *ListArtifactSubscriptionTaskResponseBody {
	s.TotalCount = &v
	return s
}

type ListArtifactSubscriptionTaskResponseBodyTasks struct {
	// The type of the artifact.
	//
	// example:
	//
	// IMAGE
	ArtifactType *string `json:"ArtifactType,omitempty" xml:"ArtifactType,omitempty"`
	// The end time of the artifact subscription task.
	//
	// example:
	//
	// 1692756630000
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// cri-7pd01myak****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The name of the namespace.
	//
	// example:
	//
	// test-ns
	NamespaceName *string `json:"NamespaceName,omitempty" xml:"NamespaceName,omitempty"`
	// The name of the repository.
	//
	// example:
	//
	// test-repo
	RepoName *string `json:"RepoName,omitempty" xml:"RepoName,omitempty"`
	// The name of the source namespace.
	//
	// example:
	//
	// library
	SourceNamespaceName *string `json:"SourceNamespaceName,omitempty" xml:"SourceNamespaceName,omitempty"`
	// The artifact source.
	//
	// example:
	//
	// DOCKER_HUB
	SourceProvider *string `json:"SourceProvider,omitempty" xml:"SourceProvider,omitempty"`
	// The name of the source repository.
	//
	// example:
	//
	// nginx
	SourceRepoName *string `json:"SourceRepoName,omitempty" xml:"SourceRepoName,omitempty"`
	// The type of the source artifact.
	//
	// example:
	//
	// PUBLIC
	SourceRepoType *string `json:"SourceRepoType,omitempty" xml:"SourceRepoType,omitempty"`
	// The start time of the artifact subscription task.
	//
	// example:
	//
	// 1695348301000
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The total number of subscribed tags.
	//
	// example:
	//
	// 3
	TagSubscriptionCount *int64 `json:"TagSubscriptionCount,omitempty" xml:"TagSubscriptionCount,omitempty"`
	// The total number of tags.
	//
	// example:
	//
	// 311
	TagTotalCount *int64 `json:"TagTotalCount,omitempty" xml:"TagTotalCount,omitempty"`
	// The task ID.
	//
	// example:
	//
	// crast-40le4es9yh0p****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The task result.
	//
	// example:
	//
	// SUCCESS
	TaskResult *string `json:"TaskResult,omitempty" xml:"TaskResult,omitempty"`
	// The status of the task.
	//
	// example:
	//
	// RUNNING
	TaskStatus *string `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	// The type of the task.
	//
	// example:
	//
	// AUTO
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s ListArtifactSubscriptionTaskResponseBodyTasks) String() string {
	return tea.Prettify(s)
}

func (s ListArtifactSubscriptionTaskResponseBodyTasks) GoString() string {
	return s.String()
}

func (s *ListArtifactSubscriptionTaskResponseBodyTasks) SetArtifactType(v string) *ListArtifactSubscriptionTaskResponseBodyTasks {
	s.ArtifactType = &v
	return s
}

func (s *ListArtifactSubscriptionTaskResponseBodyTasks) SetEndTime(v string) *ListArtifactSubscriptionTaskResponseBodyTasks {
	s.EndTime = &v
	return s
}

func (s *ListArtifactSubscriptionTaskResponseBodyTasks) SetInstanceId(v string) *ListArtifactSubscriptionTaskResponseBodyTasks {
	s.InstanceId = &v
	return s
}

func (s *ListArtifactSubscriptionTaskResponseBodyTasks) SetMessage(v string) *ListArtifactSubscriptionTaskResponseBodyTasks {
	s.Message = &v
	return s
}

func (s *ListArtifactSubscriptionTaskResponseBodyTasks) SetNamespaceName(v string) *ListArtifactSubscriptionTaskResponseBodyTasks {
	s.NamespaceName = &v
	return s
}

func (s *ListArtifactSubscriptionTaskResponseBodyTasks) SetRepoName(v string) *ListArtifactSubscriptionTaskResponseBodyTasks {
	s.RepoName = &v
	return s
}

func (s *ListArtifactSubscriptionTaskResponseBodyTasks) SetSourceNamespaceName(v string) *ListArtifactSubscriptionTaskResponseBodyTasks {
	s.SourceNamespaceName = &v
	return s
}

func (s *ListArtifactSubscriptionTaskResponseBodyTasks) SetSourceProvider(v string) *ListArtifactSubscriptionTaskResponseBodyTasks {
	s.SourceProvider = &v
	return s
}

func (s *ListArtifactSubscriptionTaskResponseBodyTasks) SetSourceRepoName(v string) *ListArtifactSubscriptionTaskResponseBodyTasks {
	s.SourceRepoName = &v
	return s
}

func (s *ListArtifactSubscriptionTaskResponseBodyTasks) SetSourceRepoType(v string) *ListArtifactSubscriptionTaskResponseBodyTasks {
	s.SourceRepoType = &v
	return s
}

func (s *ListArtifactSubscriptionTaskResponseBodyTasks) SetStartTime(v string) *ListArtifactSubscriptionTaskResponseBodyTasks {
	s.StartTime = &v
	return s
}

func (s *ListArtifactSubscriptionTaskResponseBodyTasks) SetTagSubscriptionCount(v int64) *ListArtifactSubscriptionTaskResponseBodyTasks {
	s.TagSubscriptionCount = &v
	return s
}

func (s *ListArtifactSubscriptionTaskResponseBodyTasks) SetTagTotalCount(v int64) *ListArtifactSubscriptionTaskResponseBodyTasks {
	s.TagTotalCount = &v
	return s
}

func (s *ListArtifactSubscriptionTaskResponseBodyTasks) SetTaskId(v string) *ListArtifactSubscriptionTaskResponseBodyTasks {
	s.TaskId = &v
	return s
}

func (s *ListArtifactSubscriptionTaskResponseBodyTasks) SetTaskResult(v string) *ListArtifactSubscriptionTaskResponseBodyTasks {
	s.TaskResult = &v
	return s
}

func (s *ListArtifactSubscriptionTaskResponseBodyTasks) SetTaskStatus(v string) *ListArtifactSubscriptionTaskResponseBodyTasks {
	s.TaskStatus = &v
	return s
}

func (s *ListArtifactSubscriptionTaskResponseBodyTasks) SetTaskType(v string) *ListArtifactSubscriptionTaskResponseBodyTasks {
	s.TaskType = &v
	return s
}

type ListArtifactSubscriptionTaskResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListArtifactSubscriptionTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListArtifactSubscriptionTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s ListArtifactSubscriptionTaskResponse) GoString() string {
	return s.String()
}

func (s *ListArtifactSubscriptionTaskResponse) SetHeaders(v map[string]*string) *ListArtifactSubscriptionTaskResponse {
	s.Headers = v
	return s
}

func (s *ListArtifactSubscriptionTaskResponse) SetStatusCode(v int32) *ListArtifactSubscriptionTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *ListArtifactSubscriptionTaskResponse) SetBody(v *ListArtifactSubscriptionTaskResponseBody) *ListArtifactSubscriptionTaskResponse {
	s.Body = v
	return s
}

type ListChainRequest struct {
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cri-4cdrlqmhn4gm****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The number of the page to return.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries to return on each page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The name of the image repository.
	//
	// example:
	//
	// repo1
	RepoName *string `json:"RepoName,omitempty" xml:"RepoName,omitempty"`
	// The name of the namespace.
	//
	// example:
	//
	// ns1
	RepoNamespaceName *string `json:"RepoNamespaceName,omitempty" xml:"RepoNamespaceName,omitempty"`
}

func (s ListChainRequest) String() string {
	return tea.Prettify(s)
}

func (s ListChainRequest) GoString() string {
	return s.String()
}

func (s *ListChainRequest) SetInstanceId(v string) *ListChainRequest {
	s.InstanceId = &v
	return s
}

func (s *ListChainRequest) SetPageNo(v int32) *ListChainRequest {
	s.PageNo = &v
	return s
}

func (s *ListChainRequest) SetPageSize(v int32) *ListChainRequest {
	s.PageSize = &v
	return s
}

func (s *ListChainRequest) SetRepoName(v string) *ListChainRequest {
	s.RepoName = &v
	return s
}

func (s *ListChainRequest) SetRepoNamespaceName(v string) *ListChainRequest {
	s.RepoNamespaceName = &v
	return s
}

type ListChainResponseBody struct {
	// The list of delivery chains.
	Chains []*ListChainResponseBodyChains `json:"Chains,omitempty" xml:"Chains,omitempty" type:"Repeated"`
	// The return value.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 1
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 85A99B10-3926-5201-958E-C06FA47F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of delivery chains.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListChainResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListChainResponseBody) GoString() string {
	return s.String()
}

func (s *ListChainResponseBody) SetChains(v []*ListChainResponseBodyChains) *ListChainResponseBody {
	s.Chains = v
	return s
}

func (s *ListChainResponseBody) SetCode(v string) *ListChainResponseBody {
	s.Code = &v
	return s
}

func (s *ListChainResponseBody) SetIsSuccess(v bool) *ListChainResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *ListChainResponseBody) SetPageNo(v int32) *ListChainResponseBody {
	s.PageNo = &v
	return s
}

func (s *ListChainResponseBody) SetPageSize(v int32) *ListChainResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListChainResponseBody) SetRequestId(v string) *ListChainResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListChainResponseBody) SetTotalCount(v int32) *ListChainResponseBody {
	s.TotalCount = &v
	return s
}

type ListChainResponseBodyChains struct {
	// The ID of the delivery chain.
	//
	// example:
	//
	// chi-0ops0gsmw5x2****
	ChainId *string `json:"ChainId,omitempty" xml:"ChainId,omitempty"`
	// The time when the delivery chain was created.
	//
	// example:
	//
	// 1638255427000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the delivery chain.
	//
	// example:
	//
	// description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// cri-4cdrlqmhn4gm****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The time when the delivery chain was last modified.
	//
	// example:
	//
	// 1638259914000
	ModifiedTime *int64 `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// The name of the delivery chain.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Repositories to which the delivery chain does not apply.
	ScopeExclude []*string `json:"ScopeExclude,omitempty" xml:"ScopeExclude,omitempty" type:"Repeated"`
	// The ID of the delivery chain scope.
	//
	// example:
	//
	// crr-nyrh2oko32xb****
	ScopeId *string `json:"ScopeId,omitempty" xml:"ScopeId,omitempty"`
	// The type of the delivery chain scope.
	//
	// example:
	//
	// REPOSITORY
	ScopeType *string `json:"ScopeType,omitempty" xml:"ScopeType,omitempty"`
}

func (s ListChainResponseBodyChains) String() string {
	return tea.Prettify(s)
}

func (s ListChainResponseBodyChains) GoString() string {
	return s.String()
}

func (s *ListChainResponseBodyChains) SetChainId(v string) *ListChainResponseBodyChains {
	s.ChainId = &v
	return s
}

func (s *ListChainResponseBodyChains) SetCreateTime(v int64) *ListChainResponseBodyChains {
	s.CreateTime = &v
	return s
}

func (s *ListChainResponseBodyChains) SetDescription(v string) *ListChainResponseBodyChains {
	s.Description = &v
	return s
}

func (s *ListChainResponseBodyChains) SetInstanceId(v string) *ListChainResponseBodyChains {
	s.InstanceId = &v
	return s
}

func (s *ListChainResponseBodyChains) SetModifiedTime(v int64) *ListChainResponseBodyChains {
	s.ModifiedTime = &v
	return s
}

func (s *ListChainResponseBodyChains) SetName(v string) *ListChainResponseBodyChains {
	s.Name = &v
	return s
}

func (s *ListChainResponseBodyChains) SetScopeExclude(v []*string) *ListChainResponseBodyChains {
	s.ScopeExclude = v
	return s
}

func (s *ListChainResponseBodyChains) SetScopeId(v string) *ListChainResponseBodyChains {
	s.ScopeId = &v
	return s
}

func (s *ListChainResponseBodyChains) SetScopeType(v string) *ListChainResponseBodyChains {
	s.ScopeType = &v
	return s
}

type ListChainResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListChainResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListChainResponse) String() string {
	return tea.Prettify(s)
}

func (s ListChainResponse) GoString() string {
	return s.String()
}

func (s *ListChainResponse) SetHeaders(v map[string]*string) *ListChainResponse {
	s.Headers = v
	return s
}

func (s *ListChainResponse) SetStatusCode(v int32) *ListChainResponse {
	s.StatusCode = &v
	return s
}

func (s *ListChainResponse) SetBody(v *ListChainResponseBody) *ListChainResponse {
	s.Body = v
	return s
}

type ListChainInstanceRequest struct {
	// The operation that you want to perform. Set this parameter to **ListChainInstance**.
	//
	// This parameter is required.
	//
	// example:
	//
	// cri-kmsiwlxxdcva****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The time when the delivery chain started.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The name of the image repository.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The time when the delivery chain is completed.
	//
	// example:
	//
	// test-repo
	RepoName *string `json:"RepoName,omitempty" xml:"RepoName,omitempty"`
	// The name of the delivery chain.
	//
	// example:
	//
	// test-namespace
	RepoNamespaceName *string `json:"RepoNamespaceName,omitempty" xml:"RepoNamespaceName,omitempty"`
}

func (s ListChainInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s ListChainInstanceRequest) GoString() string {
	return s.String()
}

func (s *ListChainInstanceRequest) SetInstanceId(v string) *ListChainInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *ListChainInstanceRequest) SetPageNo(v int32) *ListChainInstanceRequest {
	s.PageNo = &v
	return s
}

func (s *ListChainInstanceRequest) SetPageSize(v int32) *ListChainInstanceRequest {
	s.PageSize = &v
	return s
}

func (s *ListChainInstanceRequest) SetRepoName(v string) *ListChainInstanceRequest {
	s.RepoName = &v
	return s
}

func (s *ListChainInstanceRequest) SetRepoNamespaceName(v string) *ListChainInstanceRequest {
	s.RepoNamespaceName = &v
	return s
}

type ListChainInstanceResponseBody struct {
	// The number of entries to return on each page.
	ChainInstances []*ListChainInstanceResponseBodyChainInstances `json:"ChainInstances,omitempty" xml:"ChainInstances,omitempty" type:"Repeated"`
	// The version of the delivery chain.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The page number of the page to return.
	//
	// example:
	//
	// cri-kmsiwlxxdcva****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The execution record of the delivery chain.
	//
	// example:
	//
	// true
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// 30
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// Indicates whether the operation is successful.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the Container Registry instance.
	//
	// example:
	//
	// 838D1602-6D8F-47FB-B60A-656645D2****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The name of the repository.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListChainInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListChainInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *ListChainInstanceResponseBody) SetChainInstances(v []*ListChainInstanceResponseBodyChainInstances) *ListChainInstanceResponseBody {
	s.ChainInstances = v
	return s
}

func (s *ListChainInstanceResponseBody) SetCode(v string) *ListChainInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *ListChainInstanceResponseBody) SetInstanceId(v string) *ListChainInstanceResponseBody {
	s.InstanceId = &v
	return s
}

func (s *ListChainInstanceResponseBody) SetIsSuccess(v bool) *ListChainInstanceResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *ListChainInstanceResponseBody) SetPageNo(v int32) *ListChainInstanceResponseBody {
	s.PageNo = &v
	return s
}

func (s *ListChainInstanceResponseBody) SetPageSize(v int32) *ListChainInstanceResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListChainInstanceResponseBody) SetRequestId(v string) *ListChainInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListChainInstanceResponseBody) SetTotalCount(v int32) *ListChainInstanceResponseBody {
	s.TotalCount = &v
	return s
}

type ListChainInstanceResponseBodyChainInstances struct {
	// The name of the namespace.
	Chain *ListChainInstanceResponseBodyChainInstancesChain `json:"Chain,omitempty" xml:"Chain,omitempty" type:"Struct"`
	// 1
	//
	// example:
	//
	// F4CF4DDB-BEF2-5575-****-*******
	ChainInstanceId *string `json:"ChainInstanceId,omitempty" xml:"ChainInstanceId,omitempty"`
	// The ID of the Container Registry instance.
	//
	// example:
	//
	// 1636685856000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The ID of the delivery chain.
	//
	// example:
	//
	// test-repo
	RepoName *string `json:"RepoName,omitempty" xml:"RepoName,omitempty"`
	// The execution result of the delivery chain. Valid values:
	//
	// 	- `SUCCESS`
	//
	// 	- `FAILED`
	//
	// 	- `CANCELED`
	//
	// 	- `DENIED`
	//
	// example:
	//
	// test-ns
	RepoNamespaceName *string `json:"RepoNamespaceName,omitempty" xml:"RepoNamespaceName,omitempty"`
	// The list of the execution records of delivery chains.
	//
	// example:
	//
	// SUCCESS
	Result *string `json:"Result,omitempty" xml:"Result,omitempty"`
	// test-repo
	//
	// example:
	//
	// 1636685776000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The status of the delivery chain. Valid values:
	//
	// 	- `RUNNING`
	//
	// 	- `COMPLETE`
	//
	// 	- `CANCELING`
	//
	// 	- `CANCELED`
	//
	// example:
	//
	// COMPLETE
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListChainInstanceResponseBodyChainInstances) String() string {
	return tea.Prettify(s)
}

func (s ListChainInstanceResponseBodyChainInstances) GoString() string {
	return s.String()
}

func (s *ListChainInstanceResponseBodyChainInstances) SetChain(v *ListChainInstanceResponseBodyChainInstancesChain) *ListChainInstanceResponseBodyChainInstances {
	s.Chain = v
	return s
}

func (s *ListChainInstanceResponseBodyChainInstances) SetChainInstanceId(v string) *ListChainInstanceResponseBodyChainInstances {
	s.ChainInstanceId = &v
	return s
}

func (s *ListChainInstanceResponseBodyChainInstances) SetEndTime(v int64) *ListChainInstanceResponseBodyChainInstances {
	s.EndTime = &v
	return s
}

func (s *ListChainInstanceResponseBodyChainInstances) SetRepoName(v string) *ListChainInstanceResponseBodyChainInstances {
	s.RepoName = &v
	return s
}

func (s *ListChainInstanceResponseBodyChainInstances) SetRepoNamespaceName(v string) *ListChainInstanceResponseBodyChainInstances {
	s.RepoNamespaceName = &v
	return s
}

func (s *ListChainInstanceResponseBodyChainInstances) SetResult(v string) *ListChainInstanceResponseBodyChainInstances {
	s.Result = &v
	return s
}

func (s *ListChainInstanceResponseBodyChainInstances) SetStartTime(v int64) *ListChainInstanceResponseBodyChainInstances {
	s.StartTime = &v
	return s
}

func (s *ListChainInstanceResponseBodyChainInstances) SetStatus(v string) *ListChainInstanceResponseBodyChainInstances {
	s.Status = &v
	return s
}

type ListChainInstanceResponseBodyChainInstancesChain struct {
	// The name of the namespace.
	//
	// example:
	//
	// chi-m42gbku0****
	ChainId *string `json:"ChainId,omitempty" xml:"ChainId,omitempty"`
	// The number of entries returned on each page.
	//
	// example:
	//
	// test-chain
	ChainName *string `json:"ChainName,omitempty" xml:"ChainName,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 1
	Version *int64 `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s ListChainInstanceResponseBodyChainInstancesChain) String() string {
	return tea.Prettify(s)
}

func (s ListChainInstanceResponseBodyChainInstancesChain) GoString() string {
	return s.String()
}

func (s *ListChainInstanceResponseBodyChainInstancesChain) SetChainId(v string) *ListChainInstanceResponseBodyChainInstancesChain {
	s.ChainId = &v
	return s
}

func (s *ListChainInstanceResponseBodyChainInstancesChain) SetChainName(v string) *ListChainInstanceResponseBodyChainInstancesChain {
	s.ChainName = &v
	return s
}

func (s *ListChainInstanceResponseBodyChainInstancesChain) SetVersion(v int64) *ListChainInstanceResponseBodyChainInstancesChain {
	s.Version = &v
	return s
}

type ListChainInstanceResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListChainInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListChainInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s ListChainInstanceResponse) GoString() string {
	return s.String()
}

func (s *ListChainInstanceResponse) SetHeaders(v map[string]*string) *ListChainInstanceResponse {
	s.Headers = v
	return s
}

func (s *ListChainInstanceResponse) SetStatusCode(v int32) *ListChainInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *ListChainInstanceResponse) SetBody(v *ListChainInstanceResponseBody) *ListChainInstanceResponse {
	s.Body = v
	return s
}

type ListChartNamespaceRequest struct {
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cri-kmsiwlxxdcva****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the namespace.
	//
	// example:
	//
	// test
	NamespaceName *string `json:"NamespaceName,omitempty" xml:"NamespaceName,omitempty"`
	// The status of the namespace. Valid values:
	//
	// 	- `NORMAL`: The namespace is normal.
	//
	// 	- `DELETING`: The namespace is being deleted.
	//
	// example:
	//
	// NORMAL
	NamespaceStatus *string `json:"NamespaceStatus,omitempty" xml:"NamespaceStatus,omitempty"`
	// The number of the page to return.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries to return on each page.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListChartNamespaceRequest) String() string {
	return tea.Prettify(s)
}

func (s ListChartNamespaceRequest) GoString() string {
	return s.String()
}

func (s *ListChartNamespaceRequest) SetInstanceId(v string) *ListChartNamespaceRequest {
	s.InstanceId = &v
	return s
}

func (s *ListChartNamespaceRequest) SetNamespaceName(v string) *ListChartNamespaceRequest {
	s.NamespaceName = &v
	return s
}

func (s *ListChartNamespaceRequest) SetNamespaceStatus(v string) *ListChartNamespaceRequest {
	s.NamespaceStatus = &v
	return s
}

func (s *ListChartNamespaceRequest) SetPageNo(v int32) *ListChartNamespaceRequest {
	s.PageNo = &v
	return s
}

func (s *ListChartNamespaceRequest) SetPageSize(v int32) *ListChartNamespaceRequest {
	s.PageSize = &v
	return s
}

type ListChartNamespaceResponseBody struct {
	// The return value.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The namespaces.
	Namespaces []*ListChartNamespaceResponseBodyNamespaces `json:"Namespaces,omitempty" xml:"Namespaces,omitempty" type:"Repeated"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// F56D589D-AF7F-4900-BA46-62C780AC2C10
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of returned entries.
	//
	// example:
	//
	// 1
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListChartNamespaceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListChartNamespaceResponseBody) GoString() string {
	return s.String()
}

func (s *ListChartNamespaceResponseBody) SetCode(v string) *ListChartNamespaceResponseBody {
	s.Code = &v
	return s
}

func (s *ListChartNamespaceResponseBody) SetIsSuccess(v bool) *ListChartNamespaceResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *ListChartNamespaceResponseBody) SetNamespaces(v []*ListChartNamespaceResponseBodyNamespaces) *ListChartNamespaceResponseBody {
	s.Namespaces = v
	return s
}

func (s *ListChartNamespaceResponseBody) SetPageNo(v int32) *ListChartNamespaceResponseBody {
	s.PageNo = &v
	return s
}

func (s *ListChartNamespaceResponseBody) SetPageSize(v int32) *ListChartNamespaceResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListChartNamespaceResponseBody) SetRequestId(v string) *ListChartNamespaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListChartNamespaceResponseBody) SetTotalCount(v string) *ListChartNamespaceResponseBody {
	s.TotalCount = &v
	return s
}

type ListChartNamespaceResponseBodyNamespaces struct {
	// Indicates whether a repository was automatically created when a chart is pushed to the namespace.
	//
	// example:
	//
	// true
	AutoCreateRepo *bool `json:"AutoCreateRepo,omitempty" xml:"AutoCreateRepo,omitempty"`
	// The default repository type. Valid values:
	//
	// 	- `PUBLIC`: a public repository
	//
	// 	- `PRIVATE`: a private repository
	//
	// example:
	//
	// PUBLIC
	DefaultRepoType *string `json:"DefaultRepoType,omitempty" xml:"DefaultRepoType,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// cri-kmsiwlxxdcva****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the namespace.
	//
	// example:
	//
	// null
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	// The name of the namespace.
	//
	// example:
	//
	// test
	NamespaceName *string `json:"NamespaceName,omitempty" xml:"NamespaceName,omitempty"`
	// The status of the namespace. Valid values:
	//
	// 	- `NORMAL`: The namespace is normal.
	//
	// 	- `DELETING`: The namespace is being deleted.
	//
	// example:
	//
	// NORMAL
	NamespaceStatus *string `json:"NamespaceStatus,omitempty" xml:"NamespaceStatus,omitempty"`
	// The ID of the resource group to which the instance belongs.
	//
	// example:
	//
	// rg-acfm4n5kzyf****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s ListChartNamespaceResponseBodyNamespaces) String() string {
	return tea.Prettify(s)
}

func (s ListChartNamespaceResponseBodyNamespaces) GoString() string {
	return s.String()
}

func (s *ListChartNamespaceResponseBodyNamespaces) SetAutoCreateRepo(v bool) *ListChartNamespaceResponseBodyNamespaces {
	s.AutoCreateRepo = &v
	return s
}

func (s *ListChartNamespaceResponseBodyNamespaces) SetDefaultRepoType(v string) *ListChartNamespaceResponseBodyNamespaces {
	s.DefaultRepoType = &v
	return s
}

func (s *ListChartNamespaceResponseBodyNamespaces) SetInstanceId(v string) *ListChartNamespaceResponseBodyNamespaces {
	s.InstanceId = &v
	return s
}

func (s *ListChartNamespaceResponseBodyNamespaces) SetNamespaceId(v string) *ListChartNamespaceResponseBodyNamespaces {
	s.NamespaceId = &v
	return s
}

func (s *ListChartNamespaceResponseBodyNamespaces) SetNamespaceName(v string) *ListChartNamespaceResponseBodyNamespaces {
	s.NamespaceName = &v
	return s
}

func (s *ListChartNamespaceResponseBodyNamespaces) SetNamespaceStatus(v string) *ListChartNamespaceResponseBodyNamespaces {
	s.NamespaceStatus = &v
	return s
}

func (s *ListChartNamespaceResponseBodyNamespaces) SetResourceGroupId(v string) *ListChartNamespaceResponseBodyNamespaces {
	s.ResourceGroupId = &v
	return s
}

type ListChartNamespaceResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListChartNamespaceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListChartNamespaceResponse) String() string {
	return tea.Prettify(s)
}

func (s ListChartNamespaceResponse) GoString() string {
	return s.String()
}

func (s *ListChartNamespaceResponse) SetHeaders(v map[string]*string) *ListChartNamespaceResponse {
	s.Headers = v
	return s
}

func (s *ListChartNamespaceResponse) SetStatusCode(v int32) *ListChartNamespaceResponse {
	s.StatusCode = &v
	return s
}

func (s *ListChartNamespaceResponse) SetBody(v *ListChartNamespaceResponseBody) *ListChartNamespaceResponse {
	s.Body = v
	return s
}

type ListChartReleaseRequest struct {
	// The chart whose versions you want to query.
	//
	// example:
	//
	// null
	Chart *string `json:"Chart,omitempty" xml:"Chart,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cri-xkx6vujuhay0****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The number of the page to return.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries to return on each page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The name of the repository.
	//
	// This parameter is required.
	//
	// example:
	//
	// repo1
	RepoName *string `json:"RepoName,omitempty" xml:"RepoName,omitempty"`
	// The name of the namespace.
	//
	// This parameter is required.
	//
	// example:
	//
	// ns1
	RepoNamespaceName *string `json:"RepoNamespaceName,omitempty" xml:"RepoNamespaceName,omitempty"`
}

func (s ListChartReleaseRequest) String() string {
	return tea.Prettify(s)
}

func (s ListChartReleaseRequest) GoString() string {
	return s.String()
}

func (s *ListChartReleaseRequest) SetChart(v string) *ListChartReleaseRequest {
	s.Chart = &v
	return s
}

func (s *ListChartReleaseRequest) SetInstanceId(v string) *ListChartReleaseRequest {
	s.InstanceId = &v
	return s
}

func (s *ListChartReleaseRequest) SetPageNo(v int32) *ListChartReleaseRequest {
	s.PageNo = &v
	return s
}

func (s *ListChartReleaseRequest) SetPageSize(v int32) *ListChartReleaseRequest {
	s.PageSize = &v
	return s
}

func (s *ListChartReleaseRequest) SetRepoName(v string) *ListChartReleaseRequest {
	s.RepoName = &v
	return s
}

func (s *ListChartReleaseRequest) SetRepoNamespaceName(v string) *ListChartReleaseRequest {
	s.RepoNamespaceName = &v
	return s
}

type ListChartReleaseResponseBody struct {
	// The list of chart versions.
	ChartReleases []*ListChartReleaseResponseBodyChartReleases `json:"ChartReleases,omitempty" xml:"ChartReleases,omitempty" type:"Repeated"`
	// The return value.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// F68823F6-F1B5-4A4E-8421-A83CAB8F2963
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of returned entries.
	//
	// example:
	//
	// 1
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListChartReleaseResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListChartReleaseResponseBody) GoString() string {
	return s.String()
}

func (s *ListChartReleaseResponseBody) SetChartReleases(v []*ListChartReleaseResponseBodyChartReleases) *ListChartReleaseResponseBody {
	s.ChartReleases = v
	return s
}

func (s *ListChartReleaseResponseBody) SetCode(v string) *ListChartReleaseResponseBody {
	s.Code = &v
	return s
}

func (s *ListChartReleaseResponseBody) SetIsSuccess(v bool) *ListChartReleaseResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *ListChartReleaseResponseBody) SetPageNo(v int32) *ListChartReleaseResponseBody {
	s.PageNo = &v
	return s
}

func (s *ListChartReleaseResponseBody) SetPageSize(v int32) *ListChartReleaseResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListChartReleaseResponseBody) SetRequestId(v string) *ListChartReleaseResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListChartReleaseResponseBody) SetTotalCount(v string) *ListChartReleaseResponseBody {
	s.TotalCount = &v
	return s
}

type ListChartReleaseResponseBodyChartReleases struct {
	// The name of the chart version.
	//
	// example:
	//
	// chart1
	Chart *string `json:"Chart,omitempty" xml:"Chart,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// cri-xkx6vujuhay0****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The time when the chart was last modified.
	//
	// example:
	//
	// 1571930323000
	ModifiedTime *int64 `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// The version number of the chart.
	//
	// example:
	//
	// 0.1.0
	Release *string `json:"Release,omitempty" xml:"Release,omitempty"`
	// The ID of the chart repository.
	//
	// example:
	//
	// crcr-gpsu7b8chmxk****
	RepoId *string `json:"RepoId,omitempty" xml:"RepoId,omitempty"`
	// The size of the chart of the version. Unit: bytes.
	//
	// example:
	//
	// 2826
	Size *string `json:"Size,omitempty" xml:"Size,omitempty"`
	// The status of the chart.
	//
	// example:
	//
	// ENABLED
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListChartReleaseResponseBodyChartReleases) String() string {
	return tea.Prettify(s)
}

func (s ListChartReleaseResponseBodyChartReleases) GoString() string {
	return s.String()
}

func (s *ListChartReleaseResponseBodyChartReleases) SetChart(v string) *ListChartReleaseResponseBodyChartReleases {
	s.Chart = &v
	return s
}

func (s *ListChartReleaseResponseBodyChartReleases) SetInstanceId(v string) *ListChartReleaseResponseBodyChartReleases {
	s.InstanceId = &v
	return s
}

func (s *ListChartReleaseResponseBodyChartReleases) SetModifiedTime(v int64) *ListChartReleaseResponseBodyChartReleases {
	s.ModifiedTime = &v
	return s
}

func (s *ListChartReleaseResponseBodyChartReleases) SetRelease(v string) *ListChartReleaseResponseBodyChartReleases {
	s.Release = &v
	return s
}

func (s *ListChartReleaseResponseBodyChartReleases) SetRepoId(v string) *ListChartReleaseResponseBodyChartReleases {
	s.RepoId = &v
	return s
}

func (s *ListChartReleaseResponseBodyChartReleases) SetSize(v string) *ListChartReleaseResponseBodyChartReleases {
	s.Size = &v
	return s
}

func (s *ListChartReleaseResponseBodyChartReleases) SetStatus(v string) *ListChartReleaseResponseBodyChartReleases {
	s.Status = &v
	return s
}

type ListChartReleaseResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListChartReleaseResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListChartReleaseResponse) String() string {
	return tea.Prettify(s)
}

func (s ListChartReleaseResponse) GoString() string {
	return s.String()
}

func (s *ListChartReleaseResponse) SetHeaders(v map[string]*string) *ListChartReleaseResponse {
	s.Headers = v
	return s
}

func (s *ListChartReleaseResponse) SetStatusCode(v int32) *ListChartReleaseResponse {
	s.StatusCode = &v
	return s
}

func (s *ListChartReleaseResponse) SetBody(v *ListChartReleaseResponseBody) *ListChartReleaseResponse {
	s.Body = v
	return s
}

type ListChartRepositoryRequest struct {
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cri-xkx6vujuhay0****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The number of the page to return.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries to return on each page.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The name of the repository.
	//
	// example:
	//
	// ns1
	RepoName *string `json:"RepoName,omitempty" xml:"RepoName,omitempty"`
	// The name of the namespace to which the repository belongs.
	//
	// example:
	//
	// repo1
	RepoNamespaceName *string `json:"RepoNamespaceName,omitempty" xml:"RepoNamespaceName,omitempty"`
	// The status of the chart repositories that you want to query. Valid values:
	//
	// 	- `ALL`: query repositories of all status.
	//
	// 	- `NORMAL`: query normal repositories.
	//
	// 	- `DELETING`: query repositories that are being deleted.
	//
	// example:
	//
	// ALL
	RepoStatus *string `json:"RepoStatus,omitempty" xml:"RepoStatus,omitempty"`
}

func (s ListChartRepositoryRequest) String() string {
	return tea.Prettify(s)
}

func (s ListChartRepositoryRequest) GoString() string {
	return s.String()
}

func (s *ListChartRepositoryRequest) SetInstanceId(v string) *ListChartRepositoryRequest {
	s.InstanceId = &v
	return s
}

func (s *ListChartRepositoryRequest) SetPageNo(v int32) *ListChartRepositoryRequest {
	s.PageNo = &v
	return s
}

func (s *ListChartRepositoryRequest) SetPageSize(v int32) *ListChartRepositoryRequest {
	s.PageSize = &v
	return s
}

func (s *ListChartRepositoryRequest) SetRepoName(v string) *ListChartRepositoryRequest {
	s.RepoName = &v
	return s
}

func (s *ListChartRepositoryRequest) SetRepoNamespaceName(v string) *ListChartRepositoryRequest {
	s.RepoNamespaceName = &v
	return s
}

func (s *ListChartRepositoryRequest) SetRepoStatus(v string) *ListChartRepositoryRequest {
	s.RepoStatus = &v
	return s
}

type ListChartRepositoryResponseBody struct {
	// The return value.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The queried repositories.
	Repositories []*ListChartRepositoryResponseBodyRepositories `json:"Repositories,omitempty" xml:"Repositories,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 0AB62FB8-6873-4032-8515-4578D27523B7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of returned entries.
	//
	// example:
	//
	// 1
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListChartRepositoryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListChartRepositoryResponseBody) GoString() string {
	return s.String()
}

func (s *ListChartRepositoryResponseBody) SetCode(v string) *ListChartRepositoryResponseBody {
	s.Code = &v
	return s
}

func (s *ListChartRepositoryResponseBody) SetIsSuccess(v bool) *ListChartRepositoryResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *ListChartRepositoryResponseBody) SetPageNo(v int32) *ListChartRepositoryResponseBody {
	s.PageNo = &v
	return s
}

func (s *ListChartRepositoryResponseBody) SetPageSize(v int32) *ListChartRepositoryResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListChartRepositoryResponseBody) SetRepositories(v []*ListChartRepositoryResponseBodyRepositories) *ListChartRepositoryResponseBody {
	s.Repositories = v
	return s
}

func (s *ListChartRepositoryResponseBody) SetRequestId(v string) *ListChartRepositoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListChartRepositoryResponseBody) SetTotalCount(v string) *ListChartRepositoryResponseBody {
	s.TotalCount = &v
	return s
}

type ListChartRepositoryResponseBodyRepositories struct {
	// The time when the repository was created.
	//
	// example:
	//
	// 1571926644000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// cri-xkx6vujuhay0****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The time when the repository was last modified.
	//
	// example:
	//
	// 1571930329000
	ModifiedTime *int64 `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// The ID of the repository.
	//
	// example:
	//
	// crcr-gpsu7b8chmxk****
	RepoId *string `json:"RepoId,omitempty" xml:"RepoId,omitempty"`
	// The name of the repository.
	//
	// example:
	//
	// repo1
	RepoName *string `json:"RepoName,omitempty" xml:"RepoName,omitempty"`
	// The name of the namespace to which the repository belongs.
	//
	// example:
	//
	// ns1
	RepoNamespaceName *string `json:"RepoNamespaceName,omitempty" xml:"RepoNamespaceName,omitempty"`
	// The status of the repository. Valid values:
	//
	// 	- `NORMAL`: The repository is normal.
	//
	// 	- `DELETING`: The repository is being deleted.
	//
	// example:
	//
	// NORMAL
	RepoStatus *string `json:"RepoStatus,omitempty" xml:"RepoStatus,omitempty"`
	// The type of the repository. Valid values:
	//
	// 	- `PRIVATE`: a private repository
	//
	// 	- `PUBLIC`: a public repository
	//
	// example:
	//
	// PUBLIC
	RepoType *string `json:"RepoType,omitempty" xml:"RepoType,omitempty"`
	// The ID of the resource group to which the repository belongs.
	//
	// example:
	//
	// rg-aek2ikd5rxxxxxx
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The summary about the repository.
	//
	// example:
	//
	// test
	Summary *string `json:"Summary,omitempty" xml:"Summary,omitempty"`
}

func (s ListChartRepositoryResponseBodyRepositories) String() string {
	return tea.Prettify(s)
}

func (s ListChartRepositoryResponseBodyRepositories) GoString() string {
	return s.String()
}

func (s *ListChartRepositoryResponseBodyRepositories) SetCreateTime(v int64) *ListChartRepositoryResponseBodyRepositories {
	s.CreateTime = &v
	return s
}

func (s *ListChartRepositoryResponseBodyRepositories) SetInstanceId(v string) *ListChartRepositoryResponseBodyRepositories {
	s.InstanceId = &v
	return s
}

func (s *ListChartRepositoryResponseBodyRepositories) SetModifiedTime(v int64) *ListChartRepositoryResponseBodyRepositories {
	s.ModifiedTime = &v
	return s
}

func (s *ListChartRepositoryResponseBodyRepositories) SetRepoId(v string) *ListChartRepositoryResponseBodyRepositories {
	s.RepoId = &v
	return s
}

func (s *ListChartRepositoryResponseBodyRepositories) SetRepoName(v string) *ListChartRepositoryResponseBodyRepositories {
	s.RepoName = &v
	return s
}

func (s *ListChartRepositoryResponseBodyRepositories) SetRepoNamespaceName(v string) *ListChartRepositoryResponseBodyRepositories {
	s.RepoNamespaceName = &v
	return s
}

func (s *ListChartRepositoryResponseBodyRepositories) SetRepoStatus(v string) *ListChartRepositoryResponseBodyRepositories {
	s.RepoStatus = &v
	return s
}

func (s *ListChartRepositoryResponseBodyRepositories) SetRepoType(v string) *ListChartRepositoryResponseBodyRepositories {
	s.RepoType = &v
	return s
}

func (s *ListChartRepositoryResponseBodyRepositories) SetResourceGroupId(v string) *ListChartRepositoryResponseBodyRepositories {
	s.ResourceGroupId = &v
	return s
}

func (s *ListChartRepositoryResponseBodyRepositories) SetSummary(v string) *ListChartRepositoryResponseBodyRepositories {
	s.Summary = &v
	return s
}

type ListChartRepositoryResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListChartRepositoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListChartRepositoryResponse) String() string {
	return tea.Prettify(s)
}

func (s ListChartRepositoryResponse) GoString() string {
	return s.String()
}

func (s *ListChartRepositoryResponse) SetHeaders(v map[string]*string) *ListChartRepositoryResponse {
	s.Headers = v
	return s
}

func (s *ListChartRepositoryResponse) SetStatusCode(v int32) *ListChartRepositoryResponse {
	s.StatusCode = &v
	return s
}

func (s *ListChartRepositoryResponse) SetBody(v *ListChartRepositoryResponseBody) *ListChartRepositoryResponse {
	s.Body = v
	return s
}

type ListEventCenterRecordRequest struct {
	// The type of the event. Valid values:
	//
	// 	- `cr:Artifact:DeliveryChainCompleted`: The delivery chain is processed.
	//
	// 	- `cr:Artifact:SynchronizationCompleted`: The image is replicated.
	//
	// 	- `cr:Artifact:BuildCompleted`: The image is built.
	//
	// 	- `cr:Artifact:ScanCompleted`: The image is scanned.
	//
	// 	- `cr:Artifact:SigningCompleted`: The image is signed.
	//
	// example:
	//
	// cr:Artifact:DeliveryChainCompleted
	EventType *string `json:"EventType,omitempty" xml:"EventType,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cri-kmsiwlxxdcva****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The number of the page to return.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries to return on each page.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The name of the repository.
	//
	// example:
	//
	// test
	RepoName *string `json:"RepoName,omitempty" xml:"RepoName,omitempty"`
	// The name of the namespace to which the repository belongs.
	//
	// example:
	//
	// test
	RepoNamespaceName *string `json:"RepoNamespaceName,omitempty" xml:"RepoNamespaceName,omitempty"`
	// The ID of the event notification rule.
	//
	// example:
	//
	// crecr-n6pbhgjxtla***
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s ListEventCenterRecordRequest) String() string {
	return tea.Prettify(s)
}

func (s ListEventCenterRecordRequest) GoString() string {
	return s.String()
}

func (s *ListEventCenterRecordRequest) SetEventType(v string) *ListEventCenterRecordRequest {
	s.EventType = &v
	return s
}

func (s *ListEventCenterRecordRequest) SetInstanceId(v string) *ListEventCenterRecordRequest {
	s.InstanceId = &v
	return s
}

func (s *ListEventCenterRecordRequest) SetPageNo(v int32) *ListEventCenterRecordRequest {
	s.PageNo = &v
	return s
}

func (s *ListEventCenterRecordRequest) SetPageSize(v int32) *ListEventCenterRecordRequest {
	s.PageSize = &v
	return s
}

func (s *ListEventCenterRecordRequest) SetRepoName(v string) *ListEventCenterRecordRequest {
	s.RepoName = &v
	return s
}

func (s *ListEventCenterRecordRequest) SetRepoNamespaceName(v string) *ListEventCenterRecordRequest {
	s.RepoNamespaceName = &v
	return s
}

func (s *ListEventCenterRecordRequest) SetRuleId(v string) *ListEventCenterRecordRequest {
	s.RuleId = &v
	return s
}

type ListEventCenterRecordResponseBody struct {
	// The return value.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The list of historical events.
	//
	// example:
	//
	// []
	Records []*ListEventCenterRecordResponseBodyRecords `json:"Records,omitempty" xml:"Records,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 665C7A5E-BAEC-5BCD-AF9F-5F9260D672BF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total entries of historical events.
	//
	// example:
	//
	// 50
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListEventCenterRecordResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListEventCenterRecordResponseBody) GoString() string {
	return s.String()
}

func (s *ListEventCenterRecordResponseBody) SetCode(v string) *ListEventCenterRecordResponseBody {
	s.Code = &v
	return s
}

func (s *ListEventCenterRecordResponseBody) SetIsSuccess(v bool) *ListEventCenterRecordResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *ListEventCenterRecordResponseBody) SetPageNo(v int32) *ListEventCenterRecordResponseBody {
	s.PageNo = &v
	return s
}

func (s *ListEventCenterRecordResponseBody) SetPageSize(v int32) *ListEventCenterRecordResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListEventCenterRecordResponseBody) SetRecords(v []*ListEventCenterRecordResponseBodyRecords) *ListEventCenterRecordResponseBody {
	s.Records = v
	return s
}

func (s *ListEventCenterRecordResponseBody) SetRequestId(v string) *ListEventCenterRecordResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListEventCenterRecordResponseBody) SetTotalCount(v int32) *ListEventCenterRecordResponseBody {
	s.TotalCount = &v
	return s
}

type ListEventCenterRecordResponseBodyRecords struct {
	// The time when the event was created.
	//
	// example:
	//
	// 1638188622000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The event notification channel.
	//
	// example:
	//
	// EVENT_BRIDGE
	EventChannel *string `json:"EventChannel,omitempty" xml:"EventChannel,omitempty"`
	// The ID of the event notification.
	//
	// example:
	//
	// 7d478419-61df-49e5-b92b-30ce730c2127
	EventNotifyId *string `json:"EventNotifyId,omitempty" xml:"EventNotifyId,omitempty"`
	// The notification method. Valid values:
	//
	// 	- `http`: The notification is sent over HTTP.
	//
	// 	- `https`: The notification is sent over HTTPS.
	//
	// 	- `dingding`: The notification is sent by using DingTalk.
	//
	// example:
	//
	// http
	EventNotifyMethod *string `json:"EventNotifyMethod,omitempty" xml:"EventNotifyMethod,omitempty"`
	// The type of the event.
	//
	// example:
	//
	// cr:Artifact:DeliveryChainCompleted
	EventType *string `json:"EventType,omitempty" xml:"EventType,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// cri-gl34plsa****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The namespace.
	//
	// example:
	//
	// mychain
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The ID of the event record.
	//
	// example:
	//
	// crecrr-ctdbzwtkpr*****
	RecordId *string `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
	// The name of the image repository.
	//
	// example:
	//
	// ruby-2.4.0
	RepoName *string `json:"RepoName,omitempty" xml:"RepoName,omitempty"`
	// The ID of the event notification rule.
	//
	// example:
	//
	// crecr-n6pbhgjxtla*****
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// The name of the event notification rule.
	//
	// example:
	//
	// chain-demo
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The tags.
	//
	// example:
	//
	// ruby-2.4.0
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	// The time when the event was last updated.
	//
	// example:
	//
	// 1638188622000
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListEventCenterRecordResponseBodyRecords) String() string {
	return tea.Prettify(s)
}

func (s ListEventCenterRecordResponseBodyRecords) GoString() string {
	return s.String()
}

func (s *ListEventCenterRecordResponseBodyRecords) SetCreateTime(v int64) *ListEventCenterRecordResponseBodyRecords {
	s.CreateTime = &v
	return s
}

func (s *ListEventCenterRecordResponseBodyRecords) SetEventChannel(v string) *ListEventCenterRecordResponseBodyRecords {
	s.EventChannel = &v
	return s
}

func (s *ListEventCenterRecordResponseBodyRecords) SetEventNotifyId(v string) *ListEventCenterRecordResponseBodyRecords {
	s.EventNotifyId = &v
	return s
}

func (s *ListEventCenterRecordResponseBodyRecords) SetEventNotifyMethod(v string) *ListEventCenterRecordResponseBodyRecords {
	s.EventNotifyMethod = &v
	return s
}

func (s *ListEventCenterRecordResponseBodyRecords) SetEventType(v string) *ListEventCenterRecordResponseBodyRecords {
	s.EventType = &v
	return s
}

func (s *ListEventCenterRecordResponseBodyRecords) SetInstanceId(v string) *ListEventCenterRecordResponseBodyRecords {
	s.InstanceId = &v
	return s
}

func (s *ListEventCenterRecordResponseBodyRecords) SetNamespace(v string) *ListEventCenterRecordResponseBodyRecords {
	s.Namespace = &v
	return s
}

func (s *ListEventCenterRecordResponseBodyRecords) SetRecordId(v string) *ListEventCenterRecordResponseBodyRecords {
	s.RecordId = &v
	return s
}

func (s *ListEventCenterRecordResponseBodyRecords) SetRepoName(v string) *ListEventCenterRecordResponseBodyRecords {
	s.RepoName = &v
	return s
}

func (s *ListEventCenterRecordResponseBodyRecords) SetRuleId(v string) *ListEventCenterRecordResponseBodyRecords {
	s.RuleId = &v
	return s
}

func (s *ListEventCenterRecordResponseBodyRecords) SetRuleName(v string) *ListEventCenterRecordResponseBodyRecords {
	s.RuleName = &v
	return s
}

func (s *ListEventCenterRecordResponseBodyRecords) SetTag(v string) *ListEventCenterRecordResponseBodyRecords {
	s.Tag = &v
	return s
}

func (s *ListEventCenterRecordResponseBodyRecords) SetUpdateTime(v int64) *ListEventCenterRecordResponseBodyRecords {
	s.UpdateTime = &v
	return s
}

type ListEventCenterRecordResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListEventCenterRecordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListEventCenterRecordResponse) String() string {
	return tea.Prettify(s)
}

func (s ListEventCenterRecordResponse) GoString() string {
	return s.String()
}

func (s *ListEventCenterRecordResponse) SetHeaders(v map[string]*string) *ListEventCenterRecordResponse {
	s.Headers = v
	return s
}

func (s *ListEventCenterRecordResponse) SetStatusCode(v int32) *ListEventCenterRecordResponse {
	s.StatusCode = &v
	return s
}

func (s *ListEventCenterRecordResponse) SetBody(v *ListEventCenterRecordResponseBody) *ListEventCenterRecordResponse {
	s.Body = v
	return s
}

type ListEventCenterRuleNameRequest struct {
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cri-kmsiwlxxdcva****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s ListEventCenterRuleNameRequest) String() string {
	return tea.Prettify(s)
}

func (s ListEventCenterRuleNameRequest) GoString() string {
	return s.String()
}

func (s *ListEventCenterRuleNameRequest) SetInstanceId(v string) *ListEventCenterRuleNameRequest {
	s.InstanceId = &v
	return s
}

type ListEventCenterRuleNameResponseBody struct {
	// The return value.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 031572FA-7D8F-4C05-B790-1071E0E05DE6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The list of names of event notification rules.
	//
	// example:
	//
	// [{\\"RuleName\\": \\"mlf\\", \\"RuleId\\": \\"crecr-73q93pgljm1pc2fp\\"}]
	RuleNames []*ListEventCenterRuleNameResponseBodyRuleNames `json:"RuleNames,omitempty" xml:"RuleNames,omitempty" type:"Repeated"`
}

func (s ListEventCenterRuleNameResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListEventCenterRuleNameResponseBody) GoString() string {
	return s.String()
}

func (s *ListEventCenterRuleNameResponseBody) SetCode(v string) *ListEventCenterRuleNameResponseBody {
	s.Code = &v
	return s
}

func (s *ListEventCenterRuleNameResponseBody) SetIsSuccess(v bool) *ListEventCenterRuleNameResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *ListEventCenterRuleNameResponseBody) SetRequestId(v string) *ListEventCenterRuleNameResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListEventCenterRuleNameResponseBody) SetRuleNames(v []*ListEventCenterRuleNameResponseBodyRuleNames) *ListEventCenterRuleNameResponseBody {
	s.RuleNames = v
	return s
}

type ListEventCenterRuleNameResponseBodyRuleNames struct {
	// The ID of the event notification rule.
	//
	// example:
	//
	// crecr-n6pbhgjxtl*****
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// The name of the event notification rule.
	//
	// example:
	//
	// test-chain
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
}

func (s ListEventCenterRuleNameResponseBodyRuleNames) String() string {
	return tea.Prettify(s)
}

func (s ListEventCenterRuleNameResponseBodyRuleNames) GoString() string {
	return s.String()
}

func (s *ListEventCenterRuleNameResponseBodyRuleNames) SetRuleId(v string) *ListEventCenterRuleNameResponseBodyRuleNames {
	s.RuleId = &v
	return s
}

func (s *ListEventCenterRuleNameResponseBodyRuleNames) SetRuleName(v string) *ListEventCenterRuleNameResponseBodyRuleNames {
	s.RuleName = &v
	return s
}

type ListEventCenterRuleNameResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListEventCenterRuleNameResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListEventCenterRuleNameResponse) String() string {
	return tea.Prettify(s)
}

func (s ListEventCenterRuleNameResponse) GoString() string {
	return s.String()
}

func (s *ListEventCenterRuleNameResponse) SetHeaders(v map[string]*string) *ListEventCenterRuleNameResponse {
	s.Headers = v
	return s
}

func (s *ListEventCenterRuleNameResponse) SetStatusCode(v int32) *ListEventCenterRuleNameResponse {
	s.StatusCode = &v
	return s
}

func (s *ListEventCenterRuleNameResponse) SetBody(v *ListEventCenterRuleNameResponseBody) *ListEventCenterRuleNameResponse {
	s.Body = v
	return s
}

type ListInstanceRequest struct {
	// Deprecated
	//
	// The instance name.
	//
	// example:
	//
	// test
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The status of the instance. Valid values:
	//
	// 	- `PENDING`: The instance is being initialized.
	//
	// 	- `INIT_ERROR`: The initialization of the instance fails.
	//
	// 	- `STARTING`: The instance is being started.
	//
	// 	- `RUNNING`: The instance is running.
	//
	// 	- `STOPPING`: The instance is being stopped.
	//
	// 	- `STOPPED`: The instance is stopped.
	//
	// 	- `DELETING`: The instance is being deleted.
	//
	// 	- `DELETED`: The instance is deleted.
	//
	// example:
	//
	// RUNNING
	InstanceStatus *string `json:"InstanceStatus,omitempty" xml:"InstanceStatus,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the resource group to which the instance belongs.
	//
	// example:
	//
	// rg-acfmv36i4is****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s ListInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s ListInstanceRequest) GoString() string {
	return s.String()
}

func (s *ListInstanceRequest) SetInstanceName(v string) *ListInstanceRequest {
	s.InstanceName = &v
	return s
}

func (s *ListInstanceRequest) SetInstanceStatus(v string) *ListInstanceRequest {
	s.InstanceStatus = &v
	return s
}

func (s *ListInstanceRequest) SetPageNo(v int32) *ListInstanceRequest {
	s.PageNo = &v
	return s
}

func (s *ListInstanceRequest) SetPageSize(v int32) *ListInstanceRequest {
	s.PageSize = &v
	return s
}

func (s *ListInstanceRequest) SetResourceGroupId(v string) *ListInstanceRequest {
	s.ResourceGroupId = &v
	return s
}

type ListInstanceResponseBody struct {
	// The return value.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The queried instances.
	Instances []*ListInstanceResponseBodyInstances `json:"Instances,omitempty" xml:"Instances,omitempty" type:"Repeated"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page. Default value: 30.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// A2A9BA68-B264-4953-9154-CE61B1C03BA6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 12121
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *ListInstanceResponseBody) SetCode(v string) *ListInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *ListInstanceResponseBody) SetInstances(v []*ListInstanceResponseBodyInstances) *ListInstanceResponseBody {
	s.Instances = v
	return s
}

func (s *ListInstanceResponseBody) SetIsSuccess(v bool) *ListInstanceResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *ListInstanceResponseBody) SetPageNo(v int32) *ListInstanceResponseBody {
	s.PageNo = &v
	return s
}

func (s *ListInstanceResponseBody) SetPageSize(v int32) *ListInstanceResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListInstanceResponseBody) SetRequestId(v string) *ListInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListInstanceResponseBody) SetTotalCount(v int32) *ListInstanceResponseBody {
	s.TotalCount = &v
	return s
}

type ListInstanceResponseBodyInstances struct {
	// The time when the instance was created.
	//
	// example:
	//
	// 1562849679000
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// cri-sgedpenzw80e****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The issue occurs on the instance.
	//
	// example:
	//
	// oss bucket already exists
	InstanceIssue *string `json:"InstanceIssue,omitempty" xml:"InstanceIssue,omitempty"`
	// The name of the instance.
	//
	// example:
	//
	// test
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The edition of the Container Registry Enterprise Edition instance.
	//
	// example:
	//
	// Enterprise_Basic
	InstanceSpecification *string `json:"InstanceSpecification,omitempty" xml:"InstanceSpecification,omitempty"`
	// The status of the instance.
	//
	// example:
	//
	// RUNNING
	InstanceStatus *string `json:"InstanceStatus,omitempty" xml:"InstanceStatus,omitempty"`
	// The time when the instance was last modified.
	//
	// example:
	//
	// 1562849760000
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// The region ID of the instance.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the instance belongs.
	//
	// example:
	//
	// rg-aek2h3aexpy****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The tags of the instance.
	Tags []*ListInstanceResponseBodyInstancesTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s ListInstanceResponseBodyInstances) String() string {
	return tea.Prettify(s)
}

func (s ListInstanceResponseBodyInstances) GoString() string {
	return s.String()
}

func (s *ListInstanceResponseBodyInstances) SetCreateTime(v string) *ListInstanceResponseBodyInstances {
	s.CreateTime = &v
	return s
}

func (s *ListInstanceResponseBodyInstances) SetInstanceId(v string) *ListInstanceResponseBodyInstances {
	s.InstanceId = &v
	return s
}

func (s *ListInstanceResponseBodyInstances) SetInstanceIssue(v string) *ListInstanceResponseBodyInstances {
	s.InstanceIssue = &v
	return s
}

func (s *ListInstanceResponseBodyInstances) SetInstanceName(v string) *ListInstanceResponseBodyInstances {
	s.InstanceName = &v
	return s
}

func (s *ListInstanceResponseBodyInstances) SetInstanceSpecification(v string) *ListInstanceResponseBodyInstances {
	s.InstanceSpecification = &v
	return s
}

func (s *ListInstanceResponseBodyInstances) SetInstanceStatus(v string) *ListInstanceResponseBodyInstances {
	s.InstanceStatus = &v
	return s
}

func (s *ListInstanceResponseBodyInstances) SetModifiedTime(v string) *ListInstanceResponseBodyInstances {
	s.ModifiedTime = &v
	return s
}

func (s *ListInstanceResponseBodyInstances) SetRegionId(v string) *ListInstanceResponseBodyInstances {
	s.RegionId = &v
	return s
}

func (s *ListInstanceResponseBodyInstances) SetResourceGroupId(v string) *ListInstanceResponseBodyInstances {
	s.ResourceGroupId = &v
	return s
}

func (s *ListInstanceResponseBodyInstances) SetTags(v []*ListInstanceResponseBodyInstancesTags) *ListInstanceResponseBodyInstances {
	s.Tags = v
	return s
}

type ListInstanceResponseBodyInstancesTags struct {
	// The tag key.
	//
	// example:
	//
	// test_key
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The tag value.
	//
	// example:
	//
	// test_value
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s ListInstanceResponseBodyInstancesTags) String() string {
	return tea.Prettify(s)
}

func (s ListInstanceResponseBodyInstancesTags) GoString() string {
	return s.String()
}

func (s *ListInstanceResponseBodyInstancesTags) SetTagKey(v string) *ListInstanceResponseBodyInstancesTags {
	s.TagKey = &v
	return s
}

func (s *ListInstanceResponseBodyInstancesTags) SetTagValue(v string) *ListInstanceResponseBodyInstancesTags {
	s.TagValue = &v
	return s
}

type ListInstanceResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s ListInstanceResponse) GoString() string {
	return s.String()
}

func (s *ListInstanceResponse) SetHeaders(v map[string]*string) *ListInstanceResponse {
	s.Headers = v
	return s
}

func (s *ListInstanceResponse) SetStatusCode(v int32) *ListInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *ListInstanceResponse) SetBody(v *ListInstanceResponseBody) *ListInstanceResponse {
	s.Body = v
	return s
}

type ListInstanceEndpointRequest struct {
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cri-kmsiwlxxdcva****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the module that you want to access. Valid values:
	//
	// 	- `Registry`: image repositories.
	//
	// 	- `Chart`: Helm charts.
	//
	// example:
	//
	// Chart
	ModuleName *string `json:"ModuleName,omitempty" xml:"ModuleName,omitempty"`
	// Specify whether to return endpoints in simple mode. If yes, no access control information about the endpoint is returned.
	//
	// example:
	//
	// false
	Summary *bool `json:"Summary,omitempty" xml:"Summary,omitempty"`
}

func (s ListInstanceEndpointRequest) String() string {
	return tea.Prettify(s)
}

func (s ListInstanceEndpointRequest) GoString() string {
	return s.String()
}

func (s *ListInstanceEndpointRequest) SetInstanceId(v string) *ListInstanceEndpointRequest {
	s.InstanceId = &v
	return s
}

func (s *ListInstanceEndpointRequest) SetModuleName(v string) *ListInstanceEndpointRequest {
	s.ModuleName = &v
	return s
}

func (s *ListInstanceEndpointRequest) SetSummary(v bool) *ListInstanceEndpointRequest {
	s.Summary = &v
	return s
}

type ListInstanceEndpointResponseBody struct {
	// The return value.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The endpoints of the instance.
	Endpoints []*ListInstanceEndpointResponseBodyEndpoints `json:"Endpoints,omitempty" xml:"Endpoints,omitempty" type:"Repeated"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1B21A877-66A2-4095-90EB-20A7781A4A67
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListInstanceEndpointResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListInstanceEndpointResponseBody) GoString() string {
	return s.String()
}

func (s *ListInstanceEndpointResponseBody) SetCode(v string) *ListInstanceEndpointResponseBody {
	s.Code = &v
	return s
}

func (s *ListInstanceEndpointResponseBody) SetEndpoints(v []*ListInstanceEndpointResponseBodyEndpoints) *ListInstanceEndpointResponseBody {
	s.Endpoints = v
	return s
}

func (s *ListInstanceEndpointResponseBody) SetIsSuccess(v bool) *ListInstanceEndpointResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *ListInstanceEndpointResponseBody) SetRequestId(v string) *ListInstanceEndpointResponseBody {
	s.RequestId = &v
	return s
}

type ListInstanceEndpointResponseBodyEndpoints struct {
	// Indicates whether the endpoint is added to an access control list (ACL).
	//
	// example:
	//
	// true
	AclEnable *bool `json:"AclEnable,omitempty" xml:"AclEnable,omitempty"`
	// The ACLs that are configured for the instance.
	AclEntries []*ListInstanceEndpointResponseBodyEndpointsAclEntries `json:"AclEntries,omitempty" xml:"AclEntries,omitempty" type:"Repeated"`
	// The list of domain names of the Container Registry instance.
	Domains []*ListInstanceEndpointResponseBodyEndpointsDomains `json:"Domains,omitempty" xml:"Domains,omitempty" type:"Repeated"`
	// Indicates whether the endpoint is added to an ACL.
	//
	// example:
	//
	// true
	Enable *bool `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// The type of the endpoint.
	//
	// example:
	//
	// internet
	EndpointType *string `json:"EndpointType,omitempty" xml:"EndpointType,omitempty"`
	// The VPCs associated with the instance.
	LinkedVpcs []*ListInstanceEndpointResponseBodyEndpointsLinkedVpcs `json:"LinkedVpcs,omitempty" xml:"LinkedVpcs,omitempty" type:"Repeated"`
	// The status of the endpoint.
	//
	// example:
	//
	// RUNNING
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListInstanceEndpointResponseBodyEndpoints) String() string {
	return tea.Prettify(s)
}

func (s ListInstanceEndpointResponseBodyEndpoints) GoString() string {
	return s.String()
}

func (s *ListInstanceEndpointResponseBodyEndpoints) SetAclEnable(v bool) *ListInstanceEndpointResponseBodyEndpoints {
	s.AclEnable = &v
	return s
}

func (s *ListInstanceEndpointResponseBodyEndpoints) SetAclEntries(v []*ListInstanceEndpointResponseBodyEndpointsAclEntries) *ListInstanceEndpointResponseBodyEndpoints {
	s.AclEntries = v
	return s
}

func (s *ListInstanceEndpointResponseBodyEndpoints) SetDomains(v []*ListInstanceEndpointResponseBodyEndpointsDomains) *ListInstanceEndpointResponseBodyEndpoints {
	s.Domains = v
	return s
}

func (s *ListInstanceEndpointResponseBodyEndpoints) SetEnable(v bool) *ListInstanceEndpointResponseBodyEndpoints {
	s.Enable = &v
	return s
}

func (s *ListInstanceEndpointResponseBodyEndpoints) SetEndpointType(v string) *ListInstanceEndpointResponseBodyEndpoints {
	s.EndpointType = &v
	return s
}

func (s *ListInstanceEndpointResponseBodyEndpoints) SetLinkedVpcs(v []*ListInstanceEndpointResponseBodyEndpointsLinkedVpcs) *ListInstanceEndpointResponseBodyEndpoints {
	s.LinkedVpcs = v
	return s
}

func (s *ListInstanceEndpointResponseBodyEndpoints) SetStatus(v string) *ListInstanceEndpointResponseBodyEndpoints {
	s.Status = &v
	return s
}

type ListInstanceEndpointResponseBodyEndpointsAclEntries struct {
	// The information about the ACL.
	//
	// example:
	//
	// null
	Entry *string `json:"Entry,omitempty" xml:"Entry,omitempty"`
}

func (s ListInstanceEndpointResponseBodyEndpointsAclEntries) String() string {
	return tea.Prettify(s)
}

func (s ListInstanceEndpointResponseBodyEndpointsAclEntries) GoString() string {
	return s.String()
}

func (s *ListInstanceEndpointResponseBodyEndpointsAclEntries) SetEntry(v string) *ListInstanceEndpointResponseBodyEndpointsAclEntries {
	s.Entry = &v
	return s
}

type ListInstanceEndpointResponseBodyEndpointsDomains struct {
	// The domain name of the Container Registry instance.
	//
	// example:
	//
	// t****-registry.cn-shanghai.cr.aliyuncs.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The type of the domain name. Valid values:
	//
	// 	- SYSTEM: system domain name.
	//
	// 	- USER: user domain name.
	//
	// example:
	//
	// SYSTEM
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListInstanceEndpointResponseBodyEndpointsDomains) String() string {
	return tea.Prettify(s)
}

func (s ListInstanceEndpointResponseBodyEndpointsDomains) GoString() string {
	return s.String()
}

func (s *ListInstanceEndpointResponseBodyEndpointsDomains) SetDomain(v string) *ListInstanceEndpointResponseBodyEndpointsDomains {
	s.Domain = &v
	return s
}

func (s *ListInstanceEndpointResponseBodyEndpointsDomains) SetType(v string) *ListInstanceEndpointResponseBodyEndpointsDomains {
	s.Type = &v
	return s
}

type ListInstanceEndpointResponseBodyEndpointsLinkedVpcs struct {
	// VPC ID
	//
	// example:
	//
	// null
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s ListInstanceEndpointResponseBodyEndpointsLinkedVpcs) String() string {
	return tea.Prettify(s)
}

func (s ListInstanceEndpointResponseBodyEndpointsLinkedVpcs) GoString() string {
	return s.String()
}

func (s *ListInstanceEndpointResponseBodyEndpointsLinkedVpcs) SetVpcId(v string) *ListInstanceEndpointResponseBodyEndpointsLinkedVpcs {
	s.VpcId = &v
	return s
}

type ListInstanceEndpointResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListInstanceEndpointResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListInstanceEndpointResponse) String() string {
	return tea.Prettify(s)
}

func (s ListInstanceEndpointResponse) GoString() string {
	return s.String()
}

func (s *ListInstanceEndpointResponse) SetHeaders(v map[string]*string) *ListInstanceEndpointResponse {
	s.Headers = v
	return s
}

func (s *ListInstanceEndpointResponse) SetStatusCode(v int32) *ListInstanceEndpointResponse {
	s.StatusCode = &v
	return s
}

func (s *ListInstanceEndpointResponse) SetBody(v *ListInstanceEndpointResponseBody) *ListInstanceEndpointResponse {
	s.Body = v
	return s
}

type ListInstanceRegionRequest struct {
	// The language used for response parameters. Set this parameter to `zh-CN`.
	//
	// example:
	//
	// zh-CN
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s ListInstanceRegionRequest) String() string {
	return tea.Prettify(s)
}

func (s ListInstanceRegionRequest) GoString() string {
	return s.String()
}

func (s *ListInstanceRegionRequest) SetLang(v string) *ListInstanceRegionRequest {
	s.Lang = &v
	return s
}

type ListInstanceRegionResponseBody struct {
	// The return value.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// 	- `true`: The request is successful.
	//
	// 	- `false`: The request fails.
	//
	// example:
	//
	// true
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The list of regions.
	Regions []*ListInstanceRegionResponseBodyRegions `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 11F182E1-0F84-4F5B-8D3B-61E991482727
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListInstanceRegionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListInstanceRegionResponseBody) GoString() string {
	return s.String()
}

func (s *ListInstanceRegionResponseBody) SetCode(v string) *ListInstanceRegionResponseBody {
	s.Code = &v
	return s
}

func (s *ListInstanceRegionResponseBody) SetIsSuccess(v bool) *ListInstanceRegionResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *ListInstanceRegionResponseBody) SetRegions(v []*ListInstanceRegionResponseBodyRegions) *ListInstanceRegionResponseBody {
	s.Regions = v
	return s
}

func (s *ListInstanceRegionResponseBody) SetRequestId(v string) *ListInstanceRegionResponseBody {
	s.RequestId = &v
	return s
}

type ListInstanceRegionResponseBodyRegions struct {
	// The name of the region.
	//
	// example:
	//
	// China (Shenzhen)
	LocalName *string `json:"LocalName,omitempty" xml:"LocalName,omitempty"`
	// The ID of the region.
	//
	// example:
	//
	// cn-shenzhen
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListInstanceRegionResponseBodyRegions) String() string {
	return tea.Prettify(s)
}

func (s ListInstanceRegionResponseBodyRegions) GoString() string {
	return s.String()
}

func (s *ListInstanceRegionResponseBodyRegions) SetLocalName(v string) *ListInstanceRegionResponseBodyRegions {
	s.LocalName = &v
	return s
}

func (s *ListInstanceRegionResponseBodyRegions) SetRegionId(v string) *ListInstanceRegionResponseBodyRegions {
	s.RegionId = &v
	return s
}

type ListInstanceRegionResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListInstanceRegionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListInstanceRegionResponse) String() string {
	return tea.Prettify(s)
}

func (s ListInstanceRegionResponse) GoString() string {
	return s.String()
}

func (s *ListInstanceRegionResponse) SetHeaders(v map[string]*string) *ListInstanceRegionResponse {
	s.Headers = v
	return s
}

func (s *ListInstanceRegionResponse) SetStatusCode(v int32) *ListInstanceRegionResponse {
	s.StatusCode = &v
	return s
}

func (s *ListInstanceRegionResponse) SetBody(v *ListInstanceRegionResponseBody) *ListInstanceRegionResponse {
	s.Body = v
	return s
}

type ListNamespaceRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cri-94klsruryslx****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The namespace name.
	//
	// example:
	//
	// test-namespace
	NamespaceName *string `json:"NamespaceName,omitempty" xml:"NamespaceName,omitempty"`
	// The status of the namespace. Valid values:
	//
	// 	- `NORMAL`
	//
	// 	- `DELETING`
	//
	// example:
	//
	// NORMAL
	NamespaceStatus *string `json:"NamespaceStatus,omitempty" xml:"NamespaceStatus,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListNamespaceRequest) String() string {
	return tea.Prettify(s)
}

func (s ListNamespaceRequest) GoString() string {
	return s.String()
}

func (s *ListNamespaceRequest) SetInstanceId(v string) *ListNamespaceRequest {
	s.InstanceId = &v
	return s
}

func (s *ListNamespaceRequest) SetNamespaceName(v string) *ListNamespaceRequest {
	s.NamespaceName = &v
	return s
}

func (s *ListNamespaceRequest) SetNamespaceStatus(v string) *ListNamespaceRequest {
	s.NamespaceStatus = &v
	return s
}

func (s *ListNamespaceRequest) SetPageNo(v int32) *ListNamespaceRequest {
	s.PageNo = &v
	return s
}

func (s *ListNamespaceRequest) SetPageSize(v int32) *ListNamespaceRequest {
	s.PageSize = &v
	return s
}

type ListNamespaceResponseBody struct {
	// The HTTP status code.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// 	- `true`: The request is successful.
	//
	// 	- `false`: The request fails.
	//
	// example:
	//
	// true
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The queried namespaces.
	Namespaces []*ListNamespaceResponseBodyNamespaces `json:"Namespaces,omitempty" xml:"Namespaces,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// B7E5FCA5-55ED-451C-9649-0BB2B93387D0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of the queried namespaces.
	//
	// example:
	//
	// 1
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListNamespaceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListNamespaceResponseBody) GoString() string {
	return s.String()
}

func (s *ListNamespaceResponseBody) SetCode(v string) *ListNamespaceResponseBody {
	s.Code = &v
	return s
}

func (s *ListNamespaceResponseBody) SetIsSuccess(v bool) *ListNamespaceResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *ListNamespaceResponseBody) SetNamespaces(v []*ListNamespaceResponseBodyNamespaces) *ListNamespaceResponseBody {
	s.Namespaces = v
	return s
}

func (s *ListNamespaceResponseBody) SetPageNo(v int32) *ListNamespaceResponseBody {
	s.PageNo = &v
	return s
}

func (s *ListNamespaceResponseBody) SetPageSize(v int32) *ListNamespaceResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListNamespaceResponseBody) SetRequestId(v string) *ListNamespaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListNamespaceResponseBody) SetTotalCount(v string) *ListNamespaceResponseBody {
	s.TotalCount = &v
	return s
}

type ListNamespaceResponseBodyNamespaces struct {
	// Indicates whether the automatically creating repositories feature is enabled for the namespace.
	//
	// example:
	//
	// true
	AutoCreateRepo           *bool              `json:"AutoCreateRepo,omitempty" xml:"AutoCreateRepo,omitempty"`
	DefaultRepoConfiguration *RepoConfiguration `json:"DefaultRepoConfiguration,omitempty" xml:"DefaultRepoConfiguration,omitempty"`
	// Deprecated
	//
	// The default type of repositories in the namespace. Valid values:
	//
	// 	- `PUBLIC`: public repositories.
	//
	// 	- `PRIVATE`: private repositories.
	//
	// example:
	//
	// PUBLIC
	DefaultRepoType *string `json:"DefaultRepoType,omitempty" xml:"DefaultRepoType,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// cri-94klsruryslx****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The namespace ID.
	//
	// example:
	//
	// crn-tiw8t3f8i5lt****
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	// The namespace name.
	//
	// example:
	//
	// test
	NamespaceName *string `json:"NamespaceName,omitempty" xml:"NamespaceName,omitempty"`
	// The status of the namespace. Valid values:
	//
	// 	- `NORMAL`: The namespace is normal.
	//
	// 	- `DELETING`: The namespace is being deleted.
	//
	// example:
	//
	// NORMAL
	NamespaceStatus *string `json:"NamespaceStatus,omitempty" xml:"NamespaceStatus,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfm4n5kzyf2fbi
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s ListNamespaceResponseBodyNamespaces) String() string {
	return tea.Prettify(s)
}

func (s ListNamespaceResponseBodyNamespaces) GoString() string {
	return s.String()
}

func (s *ListNamespaceResponseBodyNamespaces) SetAutoCreateRepo(v bool) *ListNamespaceResponseBodyNamespaces {
	s.AutoCreateRepo = &v
	return s
}

func (s *ListNamespaceResponseBodyNamespaces) SetDefaultRepoConfiguration(v *RepoConfiguration) *ListNamespaceResponseBodyNamespaces {
	s.DefaultRepoConfiguration = v
	return s
}

func (s *ListNamespaceResponseBodyNamespaces) SetDefaultRepoType(v string) *ListNamespaceResponseBodyNamespaces {
	s.DefaultRepoType = &v
	return s
}

func (s *ListNamespaceResponseBodyNamespaces) SetInstanceId(v string) *ListNamespaceResponseBodyNamespaces {
	s.InstanceId = &v
	return s
}

func (s *ListNamespaceResponseBodyNamespaces) SetNamespaceId(v string) *ListNamespaceResponseBodyNamespaces {
	s.NamespaceId = &v
	return s
}

func (s *ListNamespaceResponseBodyNamespaces) SetNamespaceName(v string) *ListNamespaceResponseBodyNamespaces {
	s.NamespaceName = &v
	return s
}

func (s *ListNamespaceResponseBodyNamespaces) SetNamespaceStatus(v string) *ListNamespaceResponseBodyNamespaces {
	s.NamespaceStatus = &v
	return s
}

func (s *ListNamespaceResponseBodyNamespaces) SetResourceGroupId(v string) *ListNamespaceResponseBodyNamespaces {
	s.ResourceGroupId = &v
	return s
}

type ListNamespaceResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListNamespaceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListNamespaceResponse) String() string {
	return tea.Prettify(s)
}

func (s ListNamespaceResponse) GoString() string {
	return s.String()
}

func (s *ListNamespaceResponse) SetHeaders(v map[string]*string) *ListNamespaceResponse {
	s.Headers = v
	return s
}

func (s *ListNamespaceResponse) SetStatusCode(v int32) *ListNamespaceResponse {
	s.StatusCode = &v
	return s
}

func (s *ListNamespaceResponse) SetBody(v *ListNamespaceResponseBody) *ListNamespaceResponse {
	s.Body = v
	return s
}

type ListRepoBuildRecordRequest struct {
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cri-kmsiwlxxdcva****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The number of the page to return.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries to return on each page.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the repository.
	//
	// This parameter is required.
	//
	// example:
	//
	// crr-tquyps22md8****
	RepoId *string `json:"RepoId,omitempty" xml:"RepoId,omitempty"`
}

func (s ListRepoBuildRecordRequest) String() string {
	return tea.Prettify(s)
}

func (s ListRepoBuildRecordRequest) GoString() string {
	return s.String()
}

func (s *ListRepoBuildRecordRequest) SetInstanceId(v string) *ListRepoBuildRecordRequest {
	s.InstanceId = &v
	return s
}

func (s *ListRepoBuildRecordRequest) SetPageNo(v int32) *ListRepoBuildRecordRequest {
	s.PageNo = &v
	return s
}

func (s *ListRepoBuildRecordRequest) SetPageSize(v int32) *ListRepoBuildRecordRequest {
	s.PageSize = &v
	return s
}

func (s *ListRepoBuildRecordRequest) SetRepoId(v string) *ListRepoBuildRecordRequest {
	s.RepoId = &v
	return s
}

type ListRepoBuildRecordResponseBody struct {
	// The list of image building records.
	BuildRecords []*ListRepoBuildRecordResponseBodyBuildRecords `json:"BuildRecords,omitempty" xml:"BuildRecords,omitempty" type:"Repeated"`
	// The return value.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// 	- `true`: The request is successful.
	//
	// 	- `false`: The request fails.
	//
	// example:
	//
	// true
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 9D23DEDF-E91D-434B-B7D5-9D12C648D166
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of returned entries.
	//
	// example:
	//
	// 1
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListRepoBuildRecordResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListRepoBuildRecordResponseBody) GoString() string {
	return s.String()
}

func (s *ListRepoBuildRecordResponseBody) SetBuildRecords(v []*ListRepoBuildRecordResponseBodyBuildRecords) *ListRepoBuildRecordResponseBody {
	s.BuildRecords = v
	return s
}

func (s *ListRepoBuildRecordResponseBody) SetCode(v string) *ListRepoBuildRecordResponseBody {
	s.Code = &v
	return s
}

func (s *ListRepoBuildRecordResponseBody) SetIsSuccess(v bool) *ListRepoBuildRecordResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *ListRepoBuildRecordResponseBody) SetPageNo(v int32) *ListRepoBuildRecordResponseBody {
	s.PageNo = &v
	return s
}

func (s *ListRepoBuildRecordResponseBody) SetPageSize(v int32) *ListRepoBuildRecordResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListRepoBuildRecordResponseBody) SetRequestId(v string) *ListRepoBuildRecordResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRepoBuildRecordResponseBody) SetTotalCount(v string) *ListRepoBuildRecordResponseBody {
	s.TotalCount = &v
	return s
}

type ListRepoBuildRecordResponseBodyBuildRecords struct {
	// The ID of the image building record.
	//
	// example:
	//
	// 537e08ab-735e-415f-b7c2-160eb87f8****
	BuildRecordId *string `json:"BuildRecordId,omitempty" xml:"BuildRecordId,omitempty"`
	// The status of the image building.
	//
	// example:
	//
	// SUCCESS
	BuildStatus *string `json:"BuildStatus,omitempty" xml:"BuildStatus,omitempty"`
	// The time when the image building ended.
	//
	// example:
	//
	// 1572875610000
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The information about the image.
	Image *ListRepoBuildRecordResponseBodyBuildRecordsImage `json:"Image,omitempty" xml:"Image,omitempty" type:"Struct"`
	// The time when the image building started.
	//
	// example:
	//
	// 1572872207000
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ListRepoBuildRecordResponseBodyBuildRecords) String() string {
	return tea.Prettify(s)
}

func (s ListRepoBuildRecordResponseBodyBuildRecords) GoString() string {
	return s.String()
}

func (s *ListRepoBuildRecordResponseBodyBuildRecords) SetBuildRecordId(v string) *ListRepoBuildRecordResponseBodyBuildRecords {
	s.BuildRecordId = &v
	return s
}

func (s *ListRepoBuildRecordResponseBodyBuildRecords) SetBuildStatus(v string) *ListRepoBuildRecordResponseBodyBuildRecords {
	s.BuildStatus = &v
	return s
}

func (s *ListRepoBuildRecordResponseBodyBuildRecords) SetEndTime(v string) *ListRepoBuildRecordResponseBodyBuildRecords {
	s.EndTime = &v
	return s
}

func (s *ListRepoBuildRecordResponseBodyBuildRecords) SetImage(v *ListRepoBuildRecordResponseBodyBuildRecordsImage) *ListRepoBuildRecordResponseBodyBuildRecords {
	s.Image = v
	return s
}

func (s *ListRepoBuildRecordResponseBodyBuildRecords) SetStartTime(v string) *ListRepoBuildRecordResponseBodyBuildRecords {
	s.StartTime = &v
	return s
}

type ListRepoBuildRecordResponseBodyBuildRecordsImage struct {
	// The tag of the image.
	//
	// example:
	//
	// v0.1
	ImageTag *string `json:"ImageTag,omitempty" xml:"ImageTag,omitempty"`
	// The ID of the repository.
	//
	// example:
	//
	// crr-gzsrlevmvoaq****
	RepoId *string `json:"RepoId,omitempty" xml:"RepoId,omitempty"`
	// The name of the repository.
	//
	// example:
	//
	// test
	RepoName *string `json:"RepoName,omitempty" xml:"RepoName,omitempty"`
	// The name of the namespace to which the repository belongs.
	//
	// example:
	//
	// test
	RepoNamespaceName *string `json:"RepoNamespaceName,omitempty" xml:"RepoNamespaceName,omitempty"`
}

func (s ListRepoBuildRecordResponseBodyBuildRecordsImage) String() string {
	return tea.Prettify(s)
}

func (s ListRepoBuildRecordResponseBodyBuildRecordsImage) GoString() string {
	return s.String()
}

func (s *ListRepoBuildRecordResponseBodyBuildRecordsImage) SetImageTag(v string) *ListRepoBuildRecordResponseBodyBuildRecordsImage {
	s.ImageTag = &v
	return s
}

func (s *ListRepoBuildRecordResponseBodyBuildRecordsImage) SetRepoId(v string) *ListRepoBuildRecordResponseBodyBuildRecordsImage {
	s.RepoId = &v
	return s
}

func (s *ListRepoBuildRecordResponseBodyBuildRecordsImage) SetRepoName(v string) *ListRepoBuildRecordResponseBodyBuildRecordsImage {
	s.RepoName = &v
	return s
}

func (s *ListRepoBuildRecordResponseBodyBuildRecordsImage) SetRepoNamespaceName(v string) *ListRepoBuildRecordResponseBodyBuildRecordsImage {
	s.RepoNamespaceName = &v
	return s
}

type ListRepoBuildRecordResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListRepoBuildRecordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListRepoBuildRecordResponse) String() string {
	return tea.Prettify(s)
}

func (s ListRepoBuildRecordResponse) GoString() string {
	return s.String()
}

func (s *ListRepoBuildRecordResponse) SetHeaders(v map[string]*string) *ListRepoBuildRecordResponse {
	s.Headers = v
	return s
}

func (s *ListRepoBuildRecordResponse) SetStatusCode(v int32) *ListRepoBuildRecordResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRepoBuildRecordResponse) SetBody(v *ListRepoBuildRecordResponseBody) *ListRepoBuildRecordResponse {
	s.Body = v
	return s
}

type ListRepoBuildRecordLogRequest struct {
	// The ID of the image building record.
	//
	// This parameter is required.
	//
	// example:
	//
	// C5B4D5D7-A1C6-4E9B-ABD2-401361C4****
	BuildRecordId *string `json:"BuildRecordId,omitempty" xml:"BuildRecordId,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cri-nmbv37dlv5d3****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The offset of log lines.
	//
	// example:
	//
	// 0
	Offset *int32 `json:"Offset,omitempty" xml:"Offset,omitempty"`
	// The ID of the image repository.
	//
	// example:
	//
	// crr-z4dvahhku9wv4****
	RepoId *string `json:"RepoId,omitempty" xml:"RepoId,omitempty"`
}

func (s ListRepoBuildRecordLogRequest) String() string {
	return tea.Prettify(s)
}

func (s ListRepoBuildRecordLogRequest) GoString() string {
	return s.String()
}

func (s *ListRepoBuildRecordLogRequest) SetBuildRecordId(v string) *ListRepoBuildRecordLogRequest {
	s.BuildRecordId = &v
	return s
}

func (s *ListRepoBuildRecordLogRequest) SetInstanceId(v string) *ListRepoBuildRecordLogRequest {
	s.InstanceId = &v
	return s
}

func (s *ListRepoBuildRecordLogRequest) SetOffset(v int32) *ListRepoBuildRecordLogRequest {
	s.Offset = &v
	return s
}

func (s *ListRepoBuildRecordLogRequest) SetRepoId(v string) *ListRepoBuildRecordLogRequest {
	s.RepoId = &v
	return s
}

type ListRepoBuildRecordLogResponseBody struct {
	// The log content of the image building record.
	BuildRecordLogs []*ListRepoBuildRecordLogResponseBodyBuildRecordLogs `json:"BuildRecordLogs,omitempty" xml:"BuildRecordLogs,omitempty" type:"Repeated"`
	// The return value.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 100
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 4CE1F661-75DD-4EBD-A4AD-057B26834ABB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of returned entries.
	//
	// example:
	//
	// 1000
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListRepoBuildRecordLogResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListRepoBuildRecordLogResponseBody) GoString() string {
	return s.String()
}

func (s *ListRepoBuildRecordLogResponseBody) SetBuildRecordLogs(v []*ListRepoBuildRecordLogResponseBodyBuildRecordLogs) *ListRepoBuildRecordLogResponseBody {
	s.BuildRecordLogs = v
	return s
}

func (s *ListRepoBuildRecordLogResponseBody) SetCode(v string) *ListRepoBuildRecordLogResponseBody {
	s.Code = &v
	return s
}

func (s *ListRepoBuildRecordLogResponseBody) SetIsSuccess(v bool) *ListRepoBuildRecordLogResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *ListRepoBuildRecordLogResponseBody) SetPageNo(v int32) *ListRepoBuildRecordLogResponseBody {
	s.PageNo = &v
	return s
}

func (s *ListRepoBuildRecordLogResponseBody) SetPageSize(v int32) *ListRepoBuildRecordLogResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListRepoBuildRecordLogResponseBody) SetRequestId(v string) *ListRepoBuildRecordLogResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRepoBuildRecordLogResponseBody) SetTotalCount(v string) *ListRepoBuildRecordLogResponseBody {
	s.TotalCount = &v
	return s
}

type ListRepoBuildRecordLogResponseBodyBuildRecordLogs struct {
	// The stage of the building that is recorded in the log entry.
	//
	// example:
	//
	// GIT_CLONE
	BuildStage *string `json:"BuildStage,omitempty" xml:"BuildStage,omitempty"`
	// The line number of the log entry.
	//
	// example:
	//
	// 2
	LineNumber *int32 `json:"LineNumber,omitempty" xml:"LineNumber,omitempty"`
	// The content of the log.
	//
	// example:
	//
	// fetch stage begin
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s ListRepoBuildRecordLogResponseBodyBuildRecordLogs) String() string {
	return tea.Prettify(s)
}

func (s ListRepoBuildRecordLogResponseBodyBuildRecordLogs) GoString() string {
	return s.String()
}

func (s *ListRepoBuildRecordLogResponseBodyBuildRecordLogs) SetBuildStage(v string) *ListRepoBuildRecordLogResponseBodyBuildRecordLogs {
	s.BuildStage = &v
	return s
}

func (s *ListRepoBuildRecordLogResponseBodyBuildRecordLogs) SetLineNumber(v int32) *ListRepoBuildRecordLogResponseBodyBuildRecordLogs {
	s.LineNumber = &v
	return s
}

func (s *ListRepoBuildRecordLogResponseBodyBuildRecordLogs) SetMessage(v string) *ListRepoBuildRecordLogResponseBodyBuildRecordLogs {
	s.Message = &v
	return s
}

type ListRepoBuildRecordLogResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListRepoBuildRecordLogResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListRepoBuildRecordLogResponse) String() string {
	return tea.Prettify(s)
}

func (s ListRepoBuildRecordLogResponse) GoString() string {
	return s.String()
}

func (s *ListRepoBuildRecordLogResponse) SetHeaders(v map[string]*string) *ListRepoBuildRecordLogResponse {
	s.Headers = v
	return s
}

func (s *ListRepoBuildRecordLogResponse) SetStatusCode(v int32) *ListRepoBuildRecordLogResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRepoBuildRecordLogResponse) SetBody(v *ListRepoBuildRecordLogResponseBody) *ListRepoBuildRecordLogResponse {
	s.Body = v
	return s
}

type ListRepoBuildRuleRequest struct {
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cri-kmsiwlxxdcva****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The number of the page to return.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries to return on each page.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the image repository.
	//
	// This parameter is required.
	//
	// example:
	//
	// crr-tquyps22md8****
	RepoId *string `json:"RepoId,omitempty" xml:"RepoId,omitempty"`
}

func (s ListRepoBuildRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s ListRepoBuildRuleRequest) GoString() string {
	return s.String()
}

func (s *ListRepoBuildRuleRequest) SetInstanceId(v string) *ListRepoBuildRuleRequest {
	s.InstanceId = &v
	return s
}

func (s *ListRepoBuildRuleRequest) SetPageNo(v int32) *ListRepoBuildRuleRequest {
	s.PageNo = &v
	return s
}

func (s *ListRepoBuildRuleRequest) SetPageSize(v int32) *ListRepoBuildRuleRequest {
	s.PageSize = &v
	return s
}

func (s *ListRepoBuildRuleRequest) SetRepoId(v string) *ListRepoBuildRuleRequest {
	s.RepoId = &v
	return s
}

type ListRepoBuildRuleResponseBody struct {
	// The list of image building rules.
	BuildRules []*ListRepoBuildRuleResponseBodyBuildRules `json:"BuildRules,omitempty" xml:"BuildRules,omitempty" type:"Repeated"`
	// The return value.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// 	- `true`: The request is successful.
	//
	// 	- `false`: The request fails.
	//
	// example:
	//
	// true
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 42D782C8-E8F6-4A32-BEA0-6A6AC854C22A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of returned entries.
	//
	// example:
	//
	// 1
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListRepoBuildRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListRepoBuildRuleResponseBody) GoString() string {
	return s.String()
}

func (s *ListRepoBuildRuleResponseBody) SetBuildRules(v []*ListRepoBuildRuleResponseBodyBuildRules) *ListRepoBuildRuleResponseBody {
	s.BuildRules = v
	return s
}

func (s *ListRepoBuildRuleResponseBody) SetCode(v string) *ListRepoBuildRuleResponseBody {
	s.Code = &v
	return s
}

func (s *ListRepoBuildRuleResponseBody) SetIsSuccess(v bool) *ListRepoBuildRuleResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *ListRepoBuildRuleResponseBody) SetPageNo(v int32) *ListRepoBuildRuleResponseBody {
	s.PageNo = &v
	return s
}

func (s *ListRepoBuildRuleResponseBody) SetPageSize(v int32) *ListRepoBuildRuleResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListRepoBuildRuleResponseBody) SetRequestId(v string) *ListRepoBuildRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRepoBuildRuleResponseBody) SetTotalCount(v string) *ListRepoBuildRuleResponseBody {
	s.TotalCount = &v
	return s
}

type ListRepoBuildRuleResponseBodyBuildRules struct {
	BuildArgs []*string `json:"BuildArgs,omitempty" xml:"BuildArgs,omitempty" type:"Repeated"`
	// The ID of the image building rule.
	//
	// example:
	//
	// crbr-khys0nd3asbe****
	BuildRuleId *string `json:"BuildRuleId,omitempty" xml:"BuildRuleId,omitempty"`
	// The directory of the Dockerfile.
	//
	// example:
	//
	// /
	DockerfileLocation *string `json:"DockerfileLocation,omitempty" xml:"DockerfileLocation,omitempty"`
	// The name of the Dockerfile.
	//
	// example:
	//
	// Dockerfile
	DockerfileName *string `json:"DockerfileName,omitempty" xml:"DockerfileName,omitempty"`
	// The tag of the image.
	//
	// example:
	//
	// v0.1
	ImageTag  *string   `json:"ImageTag,omitempty" xml:"ImageTag,omitempty"`
	Platforms []*string `json:"Platforms,omitempty" xml:"Platforms,omitempty" type:"Repeated"`
	// The name of the push that triggers the building rule.
	//
	// example:
	//
	// v0.1
	PushName *string `json:"PushName,omitempty" xml:"PushName,omitempty"`
	// The type of the push that triggers the image building rule. Valid values:
	//
	// 	- GIT_BRANCH: branch push
	//
	// 	- GIT_TAG: tag push
	//
	// example:
	//
	// GIT_BRANCH
	PushType *string `json:"PushType,omitempty" xml:"PushType,omitempty"`
}

func (s ListRepoBuildRuleResponseBodyBuildRules) String() string {
	return tea.Prettify(s)
}

func (s ListRepoBuildRuleResponseBodyBuildRules) GoString() string {
	return s.String()
}

func (s *ListRepoBuildRuleResponseBodyBuildRules) SetBuildArgs(v []*string) *ListRepoBuildRuleResponseBodyBuildRules {
	s.BuildArgs = v
	return s
}

func (s *ListRepoBuildRuleResponseBodyBuildRules) SetBuildRuleId(v string) *ListRepoBuildRuleResponseBodyBuildRules {
	s.BuildRuleId = &v
	return s
}

func (s *ListRepoBuildRuleResponseBodyBuildRules) SetDockerfileLocation(v string) *ListRepoBuildRuleResponseBodyBuildRules {
	s.DockerfileLocation = &v
	return s
}

func (s *ListRepoBuildRuleResponseBodyBuildRules) SetDockerfileName(v string) *ListRepoBuildRuleResponseBodyBuildRules {
	s.DockerfileName = &v
	return s
}

func (s *ListRepoBuildRuleResponseBodyBuildRules) SetImageTag(v string) *ListRepoBuildRuleResponseBodyBuildRules {
	s.ImageTag = &v
	return s
}

func (s *ListRepoBuildRuleResponseBodyBuildRules) SetPlatforms(v []*string) *ListRepoBuildRuleResponseBodyBuildRules {
	s.Platforms = v
	return s
}

func (s *ListRepoBuildRuleResponseBodyBuildRules) SetPushName(v string) *ListRepoBuildRuleResponseBodyBuildRules {
	s.PushName = &v
	return s
}

func (s *ListRepoBuildRuleResponseBodyBuildRules) SetPushType(v string) *ListRepoBuildRuleResponseBodyBuildRules {
	s.PushType = &v
	return s
}

type ListRepoBuildRuleResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListRepoBuildRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListRepoBuildRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s ListRepoBuildRuleResponse) GoString() string {
	return s.String()
}

func (s *ListRepoBuildRuleResponse) SetHeaders(v map[string]*string) *ListRepoBuildRuleResponse {
	s.Headers = v
	return s
}

func (s *ListRepoBuildRuleResponse) SetStatusCode(v int32) *ListRepoBuildRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRepoBuildRuleResponse) SetBody(v *ListRepoBuildRuleResponseBody) *ListRepoBuildRuleResponse {
	s.Body = v
	return s
}

type ListRepoSyncRuleRequest struct {
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cri-kmsiwlxxdcva****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the namespace.
	//
	// example:
	//
	// test-namespace
	NamespaceName *string `json:"NamespaceName,omitempty" xml:"NamespaceName,omitempty"`
	// The number of the page to return.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries to return on each page.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The name of the image repository.
	//
	// example:
	//
	// test-repo
	RepoName *string `json:"RepoName,omitempty" xml:"RepoName,omitempty"`
	// The ID of the destination instance.
	//
	// example:
	//
	// cri-k77rd2eo9ztt****
	TargetInstanceId *string `json:"TargetInstanceId,omitempty" xml:"TargetInstanceId,omitempty"`
	// The region ID of the destination instance.
	//
	// example:
	//
	// cn-shenzhen
	TargetRegionId *string `json:"TargetRegionId,omitempty" xml:"TargetRegionId,omitempty"`
}

func (s ListRepoSyncRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s ListRepoSyncRuleRequest) GoString() string {
	return s.String()
}

func (s *ListRepoSyncRuleRequest) SetInstanceId(v string) *ListRepoSyncRuleRequest {
	s.InstanceId = &v
	return s
}

func (s *ListRepoSyncRuleRequest) SetNamespaceName(v string) *ListRepoSyncRuleRequest {
	s.NamespaceName = &v
	return s
}

func (s *ListRepoSyncRuleRequest) SetPageNo(v int32) *ListRepoSyncRuleRequest {
	s.PageNo = &v
	return s
}

func (s *ListRepoSyncRuleRequest) SetPageSize(v int32) *ListRepoSyncRuleRequest {
	s.PageSize = &v
	return s
}

func (s *ListRepoSyncRuleRequest) SetRepoName(v string) *ListRepoSyncRuleRequest {
	s.RepoName = &v
	return s
}

func (s *ListRepoSyncRuleRequest) SetTargetInstanceId(v string) *ListRepoSyncRuleRequest {
	s.TargetInstanceId = &v
	return s
}

func (s *ListRepoSyncRuleRequest) SetTargetRegionId(v string) *ListRepoSyncRuleRequest {
	s.TargetRegionId = &v
	return s
}

type ListRepoSyncRuleResponseBody struct {
	// The return value.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 838D1602-6D8F-47FB-B60A-656645D2****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The queried synchronization rules.
	SyncRules []*ListRepoSyncRuleResponseBodySyncRules `json:"SyncRules,omitempty" xml:"SyncRules,omitempty" type:"Repeated"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListRepoSyncRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListRepoSyncRuleResponseBody) GoString() string {
	return s.String()
}

func (s *ListRepoSyncRuleResponseBody) SetCode(v string) *ListRepoSyncRuleResponseBody {
	s.Code = &v
	return s
}

func (s *ListRepoSyncRuleResponseBody) SetIsSuccess(v bool) *ListRepoSyncRuleResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *ListRepoSyncRuleResponseBody) SetPageNo(v int32) *ListRepoSyncRuleResponseBody {
	s.PageNo = &v
	return s
}

func (s *ListRepoSyncRuleResponseBody) SetPageSize(v int32) *ListRepoSyncRuleResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListRepoSyncRuleResponseBody) SetRequestId(v string) *ListRepoSyncRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRepoSyncRuleResponseBody) SetSyncRules(v []*ListRepoSyncRuleResponseBodySyncRules) *ListRepoSyncRuleResponseBody {
	s.SyncRules = v
	return s
}

func (s *ListRepoSyncRuleResponseBody) SetTotalCount(v int32) *ListRepoSyncRuleResponseBody {
	s.TotalCount = &v
	return s
}

type ListRepoSyncRuleResponseBodySyncRules struct {
	// The time when the synchronization rule was created.
	//
	// example:
	//
	// 1572604642000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// Indicates whether the synchronization is performed across Alibaba Cloud accounts. Valid values:
	//
	// 	- `true`: Images are synchronized across Alibaba Cloud accounts.
	//
	// 	- `false`: Images are synchronized within the same Alibaba Cloud account.
	//
	// Default value: `false`.
	//
	// example:
	//
	// true
	CrossUser *bool `json:"CrossUser,omitempty" xml:"CrossUser,omitempty"`
	// The ID of the source instance.
	//
	// example:
	//
	// cri-kmsiwlxxdcva****
	LocalInstanceId *string `json:"LocalInstanceId,omitempty" xml:"LocalInstanceId,omitempty"`
	// The name of the namespace in the source instance.
	//
	// example:
	//
	// test
	LocalNamespaceName *string `json:"LocalNamespaceName,omitempty" xml:"LocalNamespaceName,omitempty"`
	// The region ID of the source instance.
	//
	// example:
	//
	// cn-shanghai
	LocalRegionId *string `json:"LocalRegionId,omitempty" xml:"LocalRegionId,omitempty"`
	// The name of the repository in the source instance.
	//
	// example:
	//
	// test-repo-local
	LocalRepoName *string `json:"LocalRepoName,omitempty" xml:"LocalRepoName,omitempty"`
	// The time when the synchronization rule was last modified.
	//
	// example:
	//
	// 1572604642000
	ModifiedTime *int64 `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// The regular expression that is used to filter repositories.
	//
	// >  This parameter is valid only when SyncScope is set to `NAMESPACE`.
	//
	// example:
	//
	// .*
	RepoNameFilter *string `json:"RepoNameFilter,omitempty" xml:"RepoNameFilter,omitempty"`
	// The synchronization direction. Valid values:
	//
	// 	- `FROM`: Images are synchronized from the source instance to the destination instance.
	//
	// 	- `TO`: Images are synchronized from the destination instance to the source instance.
	//
	// example:
	//
	// FROM
	SyncDirection *string `json:"SyncDirection,omitempty" xml:"SyncDirection,omitempty"`
	// The ID of the synchronization rule.
	//
	// example:
	//
	// crsr-7lph66uloi6h****
	SyncRuleId *string `json:"SyncRuleId,omitempty" xml:"SyncRuleId,omitempty"`
	// The name of the synchronization rule.
	//
	// example:
	//
	// sync-rule-1
	SyncRuleName *string `json:"SyncRuleName,omitempty" xml:"SyncRuleName,omitempty"`
	// The synchronization scope. Valid values:
	//
	// 	- `NAMESPACE`: synchronizes the image tags in a namespace that meet the synchronization rule.
	//
	// 	- `REPO`: synchronizes the image tags in an image repository that meet the synchronization rule.
	//
	// example:
	//
	// NAMESPACE
	SyncScope *string `json:"SyncScope,omitempty" xml:"SyncScope,omitempty"`
	// The policy that is applied to trigger the synchronization rule. Valid values:
	//
	// 	- `INITIATIVE`: The synchronization rule is positively triggered.
	//
	// 	- `PASSIVE`: The synchronization rule is passively triggered.
	//
	// example:
	//
	// PASSIVE
	SyncTrigger *string `json:"SyncTrigger,omitempty" xml:"SyncTrigger,omitempty"`
	// The regular expression that is used to filter image tags.
	//
	// example:
	//
	// .*
	TagFilter *string `json:"TagFilter,omitempty" xml:"TagFilter,omitempty"`
	// The ID of the destination instance.
	//
	// example:
	//
	// cri-k77rd2eo9ztt****
	TargetInstanceId *string `json:"TargetInstanceId,omitempty" xml:"TargetInstanceId,omitempty"`
	// The name of the namespace in the destination instance.
	//
	// example:
	//
	// test
	TargetNamespaceName *string `json:"TargetNamespaceName,omitempty" xml:"TargetNamespaceName,omitempty"`
	// The region ID of the destination instance.
	//
	// example:
	//
	// cn-shenzhen
	TargetRegionId *string `json:"TargetRegionId,omitempty" xml:"TargetRegionId,omitempty"`
	// The name of the repository in the destination instance.
	//
	// example:
	//
	// test-repo-target
	TargetRepoName *string `json:"TargetRepoName,omitempty" xml:"TargetRepoName,omitempty"`
}

func (s ListRepoSyncRuleResponseBodySyncRules) String() string {
	return tea.Prettify(s)
}

func (s ListRepoSyncRuleResponseBodySyncRules) GoString() string {
	return s.String()
}

func (s *ListRepoSyncRuleResponseBodySyncRules) SetCreateTime(v int64) *ListRepoSyncRuleResponseBodySyncRules {
	s.CreateTime = &v
	return s
}

func (s *ListRepoSyncRuleResponseBodySyncRules) SetCrossUser(v bool) *ListRepoSyncRuleResponseBodySyncRules {
	s.CrossUser = &v
	return s
}

func (s *ListRepoSyncRuleResponseBodySyncRules) SetLocalInstanceId(v string) *ListRepoSyncRuleResponseBodySyncRules {
	s.LocalInstanceId = &v
	return s
}

func (s *ListRepoSyncRuleResponseBodySyncRules) SetLocalNamespaceName(v string) *ListRepoSyncRuleResponseBodySyncRules {
	s.LocalNamespaceName = &v
	return s
}

func (s *ListRepoSyncRuleResponseBodySyncRules) SetLocalRegionId(v string) *ListRepoSyncRuleResponseBodySyncRules {
	s.LocalRegionId = &v
	return s
}

func (s *ListRepoSyncRuleResponseBodySyncRules) SetLocalRepoName(v string) *ListRepoSyncRuleResponseBodySyncRules {
	s.LocalRepoName = &v
	return s
}

func (s *ListRepoSyncRuleResponseBodySyncRules) SetModifiedTime(v int64) *ListRepoSyncRuleResponseBodySyncRules {
	s.ModifiedTime = &v
	return s
}

func (s *ListRepoSyncRuleResponseBodySyncRules) SetRepoNameFilter(v string) *ListRepoSyncRuleResponseBodySyncRules {
	s.RepoNameFilter = &v
	return s
}

func (s *ListRepoSyncRuleResponseBodySyncRules) SetSyncDirection(v string) *ListRepoSyncRuleResponseBodySyncRules {
	s.SyncDirection = &v
	return s
}

func (s *ListRepoSyncRuleResponseBodySyncRules) SetSyncRuleId(v string) *ListRepoSyncRuleResponseBodySyncRules {
	s.SyncRuleId = &v
	return s
}

func (s *ListRepoSyncRuleResponseBodySyncRules) SetSyncRuleName(v string) *ListRepoSyncRuleResponseBodySyncRules {
	s.SyncRuleName = &v
	return s
}

func (s *ListRepoSyncRuleResponseBodySyncRules) SetSyncScope(v string) *ListRepoSyncRuleResponseBodySyncRules {
	s.SyncScope = &v
	return s
}

func (s *ListRepoSyncRuleResponseBodySyncRules) SetSyncTrigger(v string) *ListRepoSyncRuleResponseBodySyncRules {
	s.SyncTrigger = &v
	return s
}

func (s *ListRepoSyncRuleResponseBodySyncRules) SetTagFilter(v string) *ListRepoSyncRuleResponseBodySyncRules {
	s.TagFilter = &v
	return s
}

func (s *ListRepoSyncRuleResponseBodySyncRules) SetTargetInstanceId(v string) *ListRepoSyncRuleResponseBodySyncRules {
	s.TargetInstanceId = &v
	return s
}

func (s *ListRepoSyncRuleResponseBodySyncRules) SetTargetNamespaceName(v string) *ListRepoSyncRuleResponseBodySyncRules {
	s.TargetNamespaceName = &v
	return s
}

func (s *ListRepoSyncRuleResponseBodySyncRules) SetTargetRegionId(v string) *ListRepoSyncRuleResponseBodySyncRules {
	s.TargetRegionId = &v
	return s
}

func (s *ListRepoSyncRuleResponseBodySyncRules) SetTargetRepoName(v string) *ListRepoSyncRuleResponseBodySyncRules {
	s.TargetRepoName = &v
	return s
}

type ListRepoSyncRuleResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListRepoSyncRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListRepoSyncRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s ListRepoSyncRuleResponse) GoString() string {
	return s.String()
}

func (s *ListRepoSyncRuleResponse) SetHeaders(v map[string]*string) *ListRepoSyncRuleResponse {
	s.Headers = v
	return s
}

func (s *ListRepoSyncRuleResponse) SetStatusCode(v int32) *ListRepoSyncRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRepoSyncRuleResponse) SetBody(v *ListRepoSyncRuleResponseBody) *ListRepoSyncRuleResponse {
	s.Body = v
	return s
}

type ListRepoSyncTaskRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cri-kmsiwlxxdcva****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The repository name.
	//
	// example:
	//
	// test
	RepoName *string `json:"RepoName,omitempty" xml:"RepoName,omitempty"`
	// The name of the namespace to which the repository belongs.
	//
	// example:
	//
	// ns
	RepoNamespaceName *string `json:"RepoNamespaceName,omitempty" xml:"RepoNamespaceName,omitempty"`
	// The ID of the synchronization task record, which is the same as SyncBatchTaskId in the response.
	//
	// >  If an image meets multiple synchronization rules and multiple synchronization tasks are generated for the image, these synchronization tasks use the same SyncBatchTaskId.
	//
	// example:
	//
	// crsr-7lph66uloi6h****
	SyncRecordId *string `json:"SyncRecordId,omitempty" xml:"SyncRecordId,omitempty"`
	// The image tag.
	//
	// example:
	//
	// nginx
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
}

func (s ListRepoSyncTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s ListRepoSyncTaskRequest) GoString() string {
	return s.String()
}

func (s *ListRepoSyncTaskRequest) SetInstanceId(v string) *ListRepoSyncTaskRequest {
	s.InstanceId = &v
	return s
}

func (s *ListRepoSyncTaskRequest) SetPageNo(v int32) *ListRepoSyncTaskRequest {
	s.PageNo = &v
	return s
}

func (s *ListRepoSyncTaskRequest) SetPageSize(v int32) *ListRepoSyncTaskRequest {
	s.PageSize = &v
	return s
}

func (s *ListRepoSyncTaskRequest) SetRepoName(v string) *ListRepoSyncTaskRequest {
	s.RepoName = &v
	return s
}

func (s *ListRepoSyncTaskRequest) SetRepoNamespaceName(v string) *ListRepoSyncTaskRequest {
	s.RepoNamespaceName = &v
	return s
}

func (s *ListRepoSyncTaskRequest) SetSyncRecordId(v string) *ListRepoSyncTaskRequest {
	s.SyncRecordId = &v
	return s
}

func (s *ListRepoSyncTaskRequest) SetTag(v string) *ListRepoSyncTaskRequest {
	s.Tag = &v
	return s
}

type ListRepoSyncTaskResponseBody struct {
	// The HTTP status code.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 7640819A-FB5B-4E25-A227-97717F62****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The queried synchronization tasks.
	SyncTasks []*ListRepoSyncTaskResponseBodySyncTasks `json:"SyncTasks,omitempty" xml:"SyncTasks,omitempty" type:"Repeated"`
	// The total number of the queried synchronization tasks.
	//
	// example:
	//
	// 1
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListRepoSyncTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListRepoSyncTaskResponseBody) GoString() string {
	return s.String()
}

func (s *ListRepoSyncTaskResponseBody) SetCode(v string) *ListRepoSyncTaskResponseBody {
	s.Code = &v
	return s
}

func (s *ListRepoSyncTaskResponseBody) SetIsSuccess(v bool) *ListRepoSyncTaskResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *ListRepoSyncTaskResponseBody) SetPageNo(v int32) *ListRepoSyncTaskResponseBody {
	s.PageNo = &v
	return s
}

func (s *ListRepoSyncTaskResponseBody) SetPageSize(v int32) *ListRepoSyncTaskResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListRepoSyncTaskResponseBody) SetRequestId(v string) *ListRepoSyncTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRepoSyncTaskResponseBody) SetSyncTasks(v []*ListRepoSyncTaskResponseBodySyncTasks) *ListRepoSyncTaskResponseBody {
	s.SyncTasks = v
	return s
}

func (s *ListRepoSyncTaskResponseBody) SetTotalCount(v string) *ListRepoSyncTaskResponseBody {
	s.TotalCount = &v
	return s
}

type ListRepoSyncTaskResponseBodySyncTasks struct {
	// The time when the synchronization task was created.
	//
	// example:
	//
	// 1572839126000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// Indicates whether the synchronization task is performed across Alibaba Cloud accounts. Valid values:
	//
	// 	- `true`: The image synchronization task is performed across accounts.
	//
	// 	- `false`: The image synchronization task is performed within the same account.
	//
	// Default value: `false`.
	//
	// example:
	//
	// true
	CrossUser *bool `json:"CrossUser,omitempty" xml:"CrossUser,omitempty"`
	// Indicates whether a custom synchronization link is used.
	//
	// example:
	//
	// true
	CustomLink *bool `json:"CustomLink,omitempty" xml:"CustomLink,omitempty"`
	// The information about the source image.
	ImageFrom *ListRepoSyncTaskResponseBodySyncTasksImageFrom `json:"ImageFrom,omitempty" xml:"ImageFrom,omitempty" type:"Struct"`
	// The information about the destination image.
	ImageTo *ListRepoSyncTaskResponseBodySyncTasksImageTo `json:"ImageTo,omitempty" xml:"ImageTo,omitempty" type:"Struct"`
	// The time when the synchronization task was last modified.
	//
	// example:
	//
	// 1572839133000
	ModifedTime *int64 `json:"ModifedTime,omitempty" xml:"ModifedTime,omitempty"`
	// The ID of the image synchronization batch tasks, which is the same as the value of SyncRecordId in the request.
	//
	// >  If an image meets multiple synchronization rules and multiple synchronization tasks are generated for the image, these synchronization tasks use the same SyncBatchTaskId.
	//
	// example:
	//
	// 15DEEB56-9271-4FDD-AC4D-C3A5CC2C****
	SyncBatchTaskId *string `json:"SyncBatchTaskId,omitempty" xml:"SyncBatchTaskId,omitempty"`
	// The ID of the synchronization rule.
	//
	// example:
	//
	// crsr-7lph66uloi6h****
	SyncRuleId *string `json:"SyncRuleId,omitempty" xml:"SyncRuleId,omitempty"`
	// The ID of the synchronization task.
	//
	// example:
	//
	// rst-4kfd7fk6pohk****
	SyncTaskId *string `json:"SyncTaskId,omitempty" xml:"SyncTaskId,omitempty"`
	// Indicates whether the synchronization transfer acceleration feature is enabled for the synchronization task.
	//
	// example:
	//
	// true
	SyncTransAccelerate *bool `json:"SyncTransAccelerate,omitempty" xml:"SyncTransAccelerate,omitempty"`
	// The error message that is returned if the synchronization task fails.
	//
	// >  The system uses this parameter to return an error message if the synchronization task fails.
	//
	// Valid value:
	//
	// 	- OSS_POLICY_UNAUTHORIZED: Container Registry is not granted permissions to access Object Storage Service (OSS).
	//
	// 	- TAG_CONFLICT: The destination repository contains an image that has the same tag as the source image, and image tag immutability is enabled for the destination repository.
	//
	// 	- UNSUPPORTED_FORMAT: The manifest or config format of the image to be synchronized is not supported.
	//
	// 	- INTERNAL_ERROR: The synchronization task failed due to internal issues on the server.
	//
	// 	- NETWORK_ERROR: The synchronization task failed due to unstable network connection.
	//
	// 	- DATA_LENGTH_EXCEEDED: The manifest or config of the image is oversized.
	//
	// example:
	//
	// NETWORK_ERROR
	TaskIssue *string `json:"TaskIssue,omitempty" xml:"TaskIssue,omitempty"`
	// The status of the synchronization task.
	//
	// example:
	//
	// SUCCESS
	TaskStatus *string `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	// The policy that is configured to trigger the synchronization task. Valid values:
	//
	// 	- `PASSIVE`: automatically triggers the synchronization task.
	//
	// 	- `INITIATIVE`: manually triggers the synchronization task.
	//
	// Default value: `PASSIVE`.
	//
	// example:
	//
	// PASSIVE
	TaskTrigger *string `json:"TaskTrigger,omitempty" xml:"TaskTrigger,omitempty"`
}

func (s ListRepoSyncTaskResponseBodySyncTasks) String() string {
	return tea.Prettify(s)
}

func (s ListRepoSyncTaskResponseBodySyncTasks) GoString() string {
	return s.String()
}

func (s *ListRepoSyncTaskResponseBodySyncTasks) SetCreateTime(v int64) *ListRepoSyncTaskResponseBodySyncTasks {
	s.CreateTime = &v
	return s
}

func (s *ListRepoSyncTaskResponseBodySyncTasks) SetCrossUser(v bool) *ListRepoSyncTaskResponseBodySyncTasks {
	s.CrossUser = &v
	return s
}

func (s *ListRepoSyncTaskResponseBodySyncTasks) SetCustomLink(v bool) *ListRepoSyncTaskResponseBodySyncTasks {
	s.CustomLink = &v
	return s
}

func (s *ListRepoSyncTaskResponseBodySyncTasks) SetImageFrom(v *ListRepoSyncTaskResponseBodySyncTasksImageFrom) *ListRepoSyncTaskResponseBodySyncTasks {
	s.ImageFrom = v
	return s
}

func (s *ListRepoSyncTaskResponseBodySyncTasks) SetImageTo(v *ListRepoSyncTaskResponseBodySyncTasksImageTo) *ListRepoSyncTaskResponseBodySyncTasks {
	s.ImageTo = v
	return s
}

func (s *ListRepoSyncTaskResponseBodySyncTasks) SetModifedTime(v int64) *ListRepoSyncTaskResponseBodySyncTasks {
	s.ModifedTime = &v
	return s
}

func (s *ListRepoSyncTaskResponseBodySyncTasks) SetSyncBatchTaskId(v string) *ListRepoSyncTaskResponseBodySyncTasks {
	s.SyncBatchTaskId = &v
	return s
}

func (s *ListRepoSyncTaskResponseBodySyncTasks) SetSyncRuleId(v string) *ListRepoSyncTaskResponseBodySyncTasks {
	s.SyncRuleId = &v
	return s
}

func (s *ListRepoSyncTaskResponseBodySyncTasks) SetSyncTaskId(v string) *ListRepoSyncTaskResponseBodySyncTasks {
	s.SyncTaskId = &v
	return s
}

func (s *ListRepoSyncTaskResponseBodySyncTasks) SetSyncTransAccelerate(v bool) *ListRepoSyncTaskResponseBodySyncTasks {
	s.SyncTransAccelerate = &v
	return s
}

func (s *ListRepoSyncTaskResponseBodySyncTasks) SetTaskIssue(v string) *ListRepoSyncTaskResponseBodySyncTasks {
	s.TaskIssue = &v
	return s
}

func (s *ListRepoSyncTaskResponseBodySyncTasks) SetTaskStatus(v string) *ListRepoSyncTaskResponseBodySyncTasks {
	s.TaskStatus = &v
	return s
}

func (s *ListRepoSyncTaskResponseBodySyncTasks) SetTaskTrigger(v string) *ListRepoSyncTaskResponseBodySyncTasks {
	s.TaskTrigger = &v
	return s
}

type ListRepoSyncTaskResponseBodySyncTasksImageFrom struct {
	// The image tag.
	//
	// example:
	//
	// v0.1
	ImageTag *string `json:"ImageTag,omitempty" xml:"ImageTag,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// cri-kmsiwlxxdcva****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The repository name.
	//
	// example:
	//
	// test
	RepoName *string `json:"RepoName,omitempty" xml:"RepoName,omitempty"`
	// The namespace to which the repository belongs.
	//
	// example:
	//
	// test
	RepoNamespaceName *string `json:"RepoNamespaceName,omitempty" xml:"RepoNamespaceName,omitempty"`
}

func (s ListRepoSyncTaskResponseBodySyncTasksImageFrom) String() string {
	return tea.Prettify(s)
}

func (s ListRepoSyncTaskResponseBodySyncTasksImageFrom) GoString() string {
	return s.String()
}

func (s *ListRepoSyncTaskResponseBodySyncTasksImageFrom) SetImageTag(v string) *ListRepoSyncTaskResponseBodySyncTasksImageFrom {
	s.ImageTag = &v
	return s
}

func (s *ListRepoSyncTaskResponseBodySyncTasksImageFrom) SetInstanceId(v string) *ListRepoSyncTaskResponseBodySyncTasksImageFrom {
	s.InstanceId = &v
	return s
}

func (s *ListRepoSyncTaskResponseBodySyncTasksImageFrom) SetRegionId(v string) *ListRepoSyncTaskResponseBodySyncTasksImageFrom {
	s.RegionId = &v
	return s
}

func (s *ListRepoSyncTaskResponseBodySyncTasksImageFrom) SetRepoName(v string) *ListRepoSyncTaskResponseBodySyncTasksImageFrom {
	s.RepoName = &v
	return s
}

func (s *ListRepoSyncTaskResponseBodySyncTasksImageFrom) SetRepoNamespaceName(v string) *ListRepoSyncTaskResponseBodySyncTasksImageFrom {
	s.RepoNamespaceName = &v
	return s
}

type ListRepoSyncTaskResponseBodySyncTasksImageTo struct {
	// The image tag.
	//
	// example:
	//
	// v0.1
	ImageTag *string `json:"ImageTag,omitempty" xml:"ImageTag,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// cri-k77rd2eo9zttneqo
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-shenzhen
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The repository name.
	//
	// example:
	//
	// test
	RepoName *string `json:"RepoName,omitempty" xml:"RepoName,omitempty"`
	// The namespace to which the repository belongs.
	//
	// example:
	//
	// test
	RepoNamespaceName *string `json:"RepoNamespaceName,omitempty" xml:"RepoNamespaceName,omitempty"`
}

func (s ListRepoSyncTaskResponseBodySyncTasksImageTo) String() string {
	return tea.Prettify(s)
}

func (s ListRepoSyncTaskResponseBodySyncTasksImageTo) GoString() string {
	return s.String()
}

func (s *ListRepoSyncTaskResponseBodySyncTasksImageTo) SetImageTag(v string) *ListRepoSyncTaskResponseBodySyncTasksImageTo {
	s.ImageTag = &v
	return s
}

func (s *ListRepoSyncTaskResponseBodySyncTasksImageTo) SetInstanceId(v string) *ListRepoSyncTaskResponseBodySyncTasksImageTo {
	s.InstanceId = &v
	return s
}

func (s *ListRepoSyncTaskResponseBodySyncTasksImageTo) SetRegionId(v string) *ListRepoSyncTaskResponseBodySyncTasksImageTo {
	s.RegionId = &v
	return s
}

func (s *ListRepoSyncTaskResponseBodySyncTasksImageTo) SetRepoName(v string) *ListRepoSyncTaskResponseBodySyncTasksImageTo {
	s.RepoName = &v
	return s
}

func (s *ListRepoSyncTaskResponseBodySyncTasksImageTo) SetRepoNamespaceName(v string) *ListRepoSyncTaskResponseBodySyncTasksImageTo {
	s.RepoNamespaceName = &v
	return s
}

type ListRepoSyncTaskResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListRepoSyncTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListRepoSyncTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s ListRepoSyncTaskResponse) GoString() string {
	return s.String()
}

func (s *ListRepoSyncTaskResponse) SetHeaders(v map[string]*string) *ListRepoSyncTaskResponse {
	s.Headers = v
	return s
}

func (s *ListRepoSyncTaskResponse) SetStatusCode(v int32) *ListRepoSyncTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRepoSyncTaskResponse) SetBody(v *ListRepoSyncTaskResponseBody) *ListRepoSyncTaskResponse {
	s.Body = v
	return s
}

type ListRepoTagRequest struct {
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cri-kmsiwlxxdcva****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The number of the page to return.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries to return on each page.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the repository.
	//
	// This parameter is required.
	//
	// example:
	//
	// crr-tquyps22md8p****
	RepoId *string `json:"RepoId,omitempty" xml:"RepoId,omitempty"`
}

func (s ListRepoTagRequest) String() string {
	return tea.Prettify(s)
}

func (s ListRepoTagRequest) GoString() string {
	return s.String()
}

func (s *ListRepoTagRequest) SetInstanceId(v string) *ListRepoTagRequest {
	s.InstanceId = &v
	return s
}

func (s *ListRepoTagRequest) SetPageNo(v int32) *ListRepoTagRequest {
	s.PageNo = &v
	return s
}

func (s *ListRepoTagRequest) SetPageSize(v int32) *ListRepoTagRequest {
	s.PageSize = &v
	return s
}

func (s *ListRepoTagRequest) SetRepoId(v string) *ListRepoTagRequest {
	s.RepoId = &v
	return s
}

type ListRepoTagResponseBody struct {
	// The return value.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The images.
	Images []*ListRepoTagResponseBodyImages `json:"Images,omitempty" xml:"Images,omitempty" type:"Repeated"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 031572FA-7D8F-4C05-B790-1071E0E05DE6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of returned entries.
	//
	// example:
	//
	// 1
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListRepoTagResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListRepoTagResponseBody) GoString() string {
	return s.String()
}

func (s *ListRepoTagResponseBody) SetCode(v string) *ListRepoTagResponseBody {
	s.Code = &v
	return s
}

func (s *ListRepoTagResponseBody) SetImages(v []*ListRepoTagResponseBodyImages) *ListRepoTagResponseBody {
	s.Images = v
	return s
}

func (s *ListRepoTagResponseBody) SetIsSuccess(v bool) *ListRepoTagResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *ListRepoTagResponseBody) SetPageNo(v int32) *ListRepoTagResponseBody {
	s.PageNo = &v
	return s
}

func (s *ListRepoTagResponseBody) SetPageSize(v int32) *ListRepoTagResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListRepoTagResponseBody) SetRequestId(v string) *ListRepoTagResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRepoTagResponseBody) SetTotalCount(v string) *ListRepoTagResponseBody {
	s.TotalCount = &v
	return s
}

type ListRepoTagResponseBodyImages struct {
	// The digest of the image.
	//
	// example:
	//
	// 67bfbcc12b67936ec7f867927817cbb071832b873dbcaed312a1930ba5f1****
	Digest *string `json:"Digest,omitempty" xml:"Digest,omitempty"`
	// The time when the image was created.
	//
	// example:
	//
	// 1572839125000
	ImageCreate *string `json:"ImageCreate,omitempty" xml:"ImageCreate,omitempty"`
	// The ID of the image.
	//
	// example:
	//
	// 45023655bf39c382e26a8607d057c27871dee163c1ecf48cc1ebf2a1****
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The size of the image.
	//
	// example:
	//
	// 27107966
	ImageSize *int64 `json:"ImageSize,omitempty" xml:"ImageSize,omitempty"`
	// The time when the image was last updated.
	//
	// example:
	//
	// 1572875608000
	ImageUpdate *string `json:"ImageUpdate,omitempty" xml:"ImageUpdate,omitempty"`
	// The status of the image.
	//
	// example:
	//
	// NORMAL
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tag of the image.
	//
	// example:
	//
	// v0.1
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
}

func (s ListRepoTagResponseBodyImages) String() string {
	return tea.Prettify(s)
}

func (s ListRepoTagResponseBodyImages) GoString() string {
	return s.String()
}

func (s *ListRepoTagResponseBodyImages) SetDigest(v string) *ListRepoTagResponseBodyImages {
	s.Digest = &v
	return s
}

func (s *ListRepoTagResponseBodyImages) SetImageCreate(v string) *ListRepoTagResponseBodyImages {
	s.ImageCreate = &v
	return s
}

func (s *ListRepoTagResponseBodyImages) SetImageId(v string) *ListRepoTagResponseBodyImages {
	s.ImageId = &v
	return s
}

func (s *ListRepoTagResponseBodyImages) SetImageSize(v int64) *ListRepoTagResponseBodyImages {
	s.ImageSize = &v
	return s
}

func (s *ListRepoTagResponseBodyImages) SetImageUpdate(v string) *ListRepoTagResponseBodyImages {
	s.ImageUpdate = &v
	return s
}

func (s *ListRepoTagResponseBodyImages) SetStatus(v string) *ListRepoTagResponseBodyImages {
	s.Status = &v
	return s
}

func (s *ListRepoTagResponseBodyImages) SetTag(v string) *ListRepoTagResponseBodyImages {
	s.Tag = &v
	return s
}

type ListRepoTagResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListRepoTagResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListRepoTagResponse) String() string {
	return tea.Prettify(s)
}

func (s ListRepoTagResponse) GoString() string {
	return s.String()
}

func (s *ListRepoTagResponse) SetHeaders(v map[string]*string) *ListRepoTagResponse {
	s.Headers = v
	return s
}

func (s *ListRepoTagResponse) SetStatusCode(v int32) *ListRepoTagResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRepoTagResponse) SetBody(v *ListRepoTagResponseBody) *ListRepoTagResponse {
	s.Body = v
	return s
}

type ListRepoTagScanResultRequest struct {
	// The digest of the image.
	//
	// example:
	//
	// sha256:6b0b094f8a904f8fb6602427aed0d1fa
	Digest *string `json:"Digest,omitempty" xml:"Digest,omitempty"`
	// The parameter whose value that you want to query. Fox example, if the value is `FixCmd`, only the `FixCmd` parameter is returned.
	//
	// example:
	//
	// FixCmd
	FilterValue *string `json:"FilterValue,omitempty" xml:"FilterValue,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cri-2j88dtld8yel****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The number of the page to return.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries to return on each page.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the image repository.
	//
	// example:
	//
	// crr-uf082u9dg8do****
	RepoId *string `json:"RepoId,omitempty" xml:"RepoId,omitempty"`
	// The ID of the security scan task.
	//
	// example:
	//
	// 6b0b094f-8a90-4f8f-b660-2427aed0****
	ScanTaskId *string `json:"ScanTaskId,omitempty" xml:"ScanTaskId,omitempty"`
	// The type of the vulnerability. Valid values:
	//
	// 	- `cve`: image system vulnerability
	//
	// 	- `sca`: image application vulnerability
	//
	// example:
	//
	// sca
	ScanType *string `json:"ScanType,omitempty" xml:"ScanType,omitempty"`
	// The severity of the vulnerability. Valid values:
	//
	// 	- `High`
	//
	// 	- `Medium`
	//
	// 	- `Low`
	//
	// 	- `Unknown`
	//
	// example:
	//
	// High
	Severity *string `json:"Severity,omitempty" xml:"Severity,omitempty"`
	// The name of the image tag.
	//
	// example:
	//
	// 1
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	// The keyword for fuzzy search used in scanning. The value can be a CVE name.
	//
	// example:
	//
	// CVE-2021
	VulQueryKey *string `json:"VulQueryKey,omitempty" xml:"VulQueryKey,omitempty"`
}

func (s ListRepoTagScanResultRequest) String() string {
	return tea.Prettify(s)
}

func (s ListRepoTagScanResultRequest) GoString() string {
	return s.String()
}

func (s *ListRepoTagScanResultRequest) SetDigest(v string) *ListRepoTagScanResultRequest {
	s.Digest = &v
	return s
}

func (s *ListRepoTagScanResultRequest) SetFilterValue(v string) *ListRepoTagScanResultRequest {
	s.FilterValue = &v
	return s
}

func (s *ListRepoTagScanResultRequest) SetInstanceId(v string) *ListRepoTagScanResultRequest {
	s.InstanceId = &v
	return s
}

func (s *ListRepoTagScanResultRequest) SetPageNo(v int32) *ListRepoTagScanResultRequest {
	s.PageNo = &v
	return s
}

func (s *ListRepoTagScanResultRequest) SetPageSize(v int32) *ListRepoTagScanResultRequest {
	s.PageSize = &v
	return s
}

func (s *ListRepoTagScanResultRequest) SetRepoId(v string) *ListRepoTagScanResultRequest {
	s.RepoId = &v
	return s
}

func (s *ListRepoTagScanResultRequest) SetScanTaskId(v string) *ListRepoTagScanResultRequest {
	s.ScanTaskId = &v
	return s
}

func (s *ListRepoTagScanResultRequest) SetScanType(v string) *ListRepoTagScanResultRequest {
	s.ScanType = &v
	return s
}

func (s *ListRepoTagScanResultRequest) SetSeverity(v string) *ListRepoTagScanResultRequest {
	s.Severity = &v
	return s
}

func (s *ListRepoTagScanResultRequest) SetTag(v string) *ListRepoTagScanResultRequest {
	s.Tag = &v
	return s
}

func (s *ListRepoTagScanResultRequest) SetVulQueryKey(v string) *ListRepoTagScanResultRequest {
	s.VulQueryKey = &v
	return s
}

type ListRepoTagScanResultResponseBody struct {
	// The return value.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// 	- `true`: The request is successful.
	//
	// 	- `false`: The request failed.
	//
	// example:
	//
	// true
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The number of the returned page.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 56B5C92F-F5D9-46E0-823F-EC71D1892DAA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of vulnerabilities detected on images.
	//
	// example:
	//
	// 196
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The details about the detected vulnerabilities.
	Vulnerabilities []*ListRepoTagScanResultResponseBodyVulnerabilities `json:"Vulnerabilities,omitempty" xml:"Vulnerabilities,omitempty" type:"Repeated"`
}

func (s ListRepoTagScanResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListRepoTagScanResultResponseBody) GoString() string {
	return s.String()
}

func (s *ListRepoTagScanResultResponseBody) SetCode(v string) *ListRepoTagScanResultResponseBody {
	s.Code = &v
	return s
}

func (s *ListRepoTagScanResultResponseBody) SetIsSuccess(v bool) *ListRepoTagScanResultResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *ListRepoTagScanResultResponseBody) SetPageNo(v int32) *ListRepoTagScanResultResponseBody {
	s.PageNo = &v
	return s
}

func (s *ListRepoTagScanResultResponseBody) SetPageSize(v int32) *ListRepoTagScanResultResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListRepoTagScanResultResponseBody) SetRequestId(v string) *ListRepoTagScanResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRepoTagScanResultResponseBody) SetTotalCount(v int32) *ListRepoTagScanResultResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListRepoTagScanResultResponseBody) SetVulnerabilities(v []*ListRepoTagScanResultResponseBodyVulnerabilities) *ListRepoTagScanResultResponseBody {
	s.Vulnerabilities = v
	return s
}

type ListRepoTagScanResultResponseBodyVulnerabilities struct {
	// The ID of the image layer where the vulnerability was detected.
	//
	// example:
	//
	// sha256:123456717b8e40b6480979b739010d8d549989602bcdd07922119aec6f9dbe57
	AddedBy *string `json:"AddedBy,omitempty" xml:"AddedBy,omitempty"`
	// The name of the vulnerability.
	//
	// example:
	//
	// Vulnerability
	AliasName *string `json:"AliasName,omitempty" xml:"AliasName,omitempty"`
	// The URL of the vulnerability.
	//
	// example:
	//
	// https://security-tracker.debian.org/tracker/CVE-2009-5155
	CveLink *string `json:"CveLink,omitempty" xml:"CveLink,omitempty"`
	// The directory of the vulnerability.
	//
	// example:
	//
	// /test.txt
	CveLocation *string `json:"CveLocation,omitempty" xml:"CveLocation,omitempty"`
	// The name of the vulnerability.
	//
	// example:
	//
	// CVE-2009-5155
	CveName *string `json:"CveName,omitempty" xml:"CveName,omitempty"`
	// The description of the vulnerability.
	//
	// example:
	//
	// description.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The cause of the vulnerability.
	//
	// example:
	//
	// eglibc
	Feature *string `json:"Feature,omitempty" xml:"Feature,omitempty"`
	// The command used to fix the vulnerability.
	//
	// example:
	//
	// yum install -y xxx
	FixCmd *string `json:"FixCmd,omitempty" xml:"FixCmd,omitempty"`
	// The type of the vulnerability. Valid values:
	//
	// 	- `cve`: image system vulnerability
	//
	// 	- `sca`: image application vulnerability
	//
	// example:
	//
	// cve
	ScanType *string `json:"ScanType,omitempty" xml:"ScanType,omitempty"`
	// The severity of the vulnerability.
	//
	// example:
	//
	// Medium
	Severity *string `json:"Severity,omitempty" xml:"Severity,omitempty"`
	// The version of the vulnerability.
	//
	// example:
	//
	// 2.19-6.9
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
	// The version where the vulnerability was fixed.
	//
	// example:
	//
	// 2.19-18+deb8u5
	VersionFixed *string `json:"VersionFixed,omitempty" xml:"VersionFixed,omitempty"`
	// The format of the vulnerability.
	//
	// example:
	//
	// dpkg
	VersionFormat *string `json:"VersionFormat,omitempty" xml:"VersionFormat,omitempty"`
}

func (s ListRepoTagScanResultResponseBodyVulnerabilities) String() string {
	return tea.Prettify(s)
}

func (s ListRepoTagScanResultResponseBodyVulnerabilities) GoString() string {
	return s.String()
}

func (s *ListRepoTagScanResultResponseBodyVulnerabilities) SetAddedBy(v string) *ListRepoTagScanResultResponseBodyVulnerabilities {
	s.AddedBy = &v
	return s
}

func (s *ListRepoTagScanResultResponseBodyVulnerabilities) SetAliasName(v string) *ListRepoTagScanResultResponseBodyVulnerabilities {
	s.AliasName = &v
	return s
}

func (s *ListRepoTagScanResultResponseBodyVulnerabilities) SetCveLink(v string) *ListRepoTagScanResultResponseBodyVulnerabilities {
	s.CveLink = &v
	return s
}

func (s *ListRepoTagScanResultResponseBodyVulnerabilities) SetCveLocation(v string) *ListRepoTagScanResultResponseBodyVulnerabilities {
	s.CveLocation = &v
	return s
}

func (s *ListRepoTagScanResultResponseBodyVulnerabilities) SetCveName(v string) *ListRepoTagScanResultResponseBodyVulnerabilities {
	s.CveName = &v
	return s
}

func (s *ListRepoTagScanResultResponseBodyVulnerabilities) SetDescription(v string) *ListRepoTagScanResultResponseBodyVulnerabilities {
	s.Description = &v
	return s
}

func (s *ListRepoTagScanResultResponseBodyVulnerabilities) SetFeature(v string) *ListRepoTagScanResultResponseBodyVulnerabilities {
	s.Feature = &v
	return s
}

func (s *ListRepoTagScanResultResponseBodyVulnerabilities) SetFixCmd(v string) *ListRepoTagScanResultResponseBodyVulnerabilities {
	s.FixCmd = &v
	return s
}

func (s *ListRepoTagScanResultResponseBodyVulnerabilities) SetScanType(v string) *ListRepoTagScanResultResponseBodyVulnerabilities {
	s.ScanType = &v
	return s
}

func (s *ListRepoTagScanResultResponseBodyVulnerabilities) SetSeverity(v string) *ListRepoTagScanResultResponseBodyVulnerabilities {
	s.Severity = &v
	return s
}

func (s *ListRepoTagScanResultResponseBodyVulnerabilities) SetVersion(v string) *ListRepoTagScanResultResponseBodyVulnerabilities {
	s.Version = &v
	return s
}

func (s *ListRepoTagScanResultResponseBodyVulnerabilities) SetVersionFixed(v string) *ListRepoTagScanResultResponseBodyVulnerabilities {
	s.VersionFixed = &v
	return s
}

func (s *ListRepoTagScanResultResponseBodyVulnerabilities) SetVersionFormat(v string) *ListRepoTagScanResultResponseBodyVulnerabilities {
	s.VersionFormat = &v
	return s
}

type ListRepoTagScanResultResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListRepoTagScanResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListRepoTagScanResultResponse) String() string {
	return tea.Prettify(s)
}

func (s ListRepoTagScanResultResponse) GoString() string {
	return s.String()
}

func (s *ListRepoTagScanResultResponse) SetHeaders(v map[string]*string) *ListRepoTagScanResultResponse {
	s.Headers = v
	return s
}

func (s *ListRepoTagScanResultResponse) SetStatusCode(v int32) *ListRepoTagScanResultResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRepoTagScanResultResponse) SetBody(v *ListRepoTagScanResultResponseBody) *ListRepoTagScanResultResponse {
	s.Body = v
	return s
}

type ListRepoTriggerRequest struct {
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cri-kmsiwlxxdcva****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the repository.
	//
	// This parameter is required.
	//
	// example:
	//
	// crr-tquyps22md8p****
	RepoId *string `json:"RepoId,omitempty" xml:"RepoId,omitempty"`
}

func (s ListRepoTriggerRequest) String() string {
	return tea.Prettify(s)
}

func (s ListRepoTriggerRequest) GoString() string {
	return s.String()
}

func (s *ListRepoTriggerRequest) SetInstanceId(v string) *ListRepoTriggerRequest {
	s.InstanceId = &v
	return s
}

func (s *ListRepoTriggerRequest) SetRepoId(v string) *ListRepoTriggerRequest {
	s.RepoId = &v
	return s
}

type ListRepoTriggerResponseBody struct {
	// The return value.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// 	- `true`: The request is successful.
	//
	// 	- `false`: The request fails.
	//
	// example:
	//
	// true
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 2CA76D52-A8F0-4D0B-854E-FBD9F6C99049
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The triggers of the repository.
	Triggers []*ListRepoTriggerResponseBodyTriggers `json:"Triggers,omitempty" xml:"Triggers,omitempty" type:"Repeated"`
}

func (s ListRepoTriggerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListRepoTriggerResponseBody) GoString() string {
	return s.String()
}

func (s *ListRepoTriggerResponseBody) SetCode(v string) *ListRepoTriggerResponseBody {
	s.Code = &v
	return s
}

func (s *ListRepoTriggerResponseBody) SetIsSuccess(v bool) *ListRepoTriggerResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *ListRepoTriggerResponseBody) SetRequestId(v string) *ListRepoTriggerResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRepoTriggerResponseBody) SetTriggers(v []*ListRepoTriggerResponseBodyTriggers) *ListRepoTriggerResponseBody {
	s.Triggers = v
	return s
}

type ListRepoTriggerResponseBodyTriggers struct {
	// The type of the event that activates the trigger. Valid values:
	//
	// 	- `BUILD_SUCCESS`: The trigger is activated when an image building task is successful.
	//
	// 	- `BUILD_Fail`: The trigger is activated when an image building task fails.
	//
	// example:
	//
	// BUILD_SUCCESS
	RepoEvent *string `json:"RepoEvent,omitempty" xml:"RepoEvent,omitempty"`
	// The ID of the trigger.
	//
	// example:
	//
	// crw-vriyql9eq7ep****
	TriggerId *string `json:"TriggerId,omitempty" xml:"TriggerId,omitempty"`
	// The name of the trigger.
	//
	// example:
	//
	// test
	TriggerName *string `json:"TriggerName,omitempty" xml:"TriggerName,omitempty"`
	// The image tag based on which the trigger is set.
	//
	// example:
	//
	// *
	TriggerTag *string `json:"TriggerTag,omitempty" xml:"TriggerTag,omitempty"`
	// The type of the trigger. Valid values:
	//
	// 	- `ALL`: a trigger that supports both tags and regular expressions.
	//
	// 	- `TAG_LISTTAG`: a tag-based trigger.
	//
	// 	- `TAG_REG_EXP`: a regular expression-based trigger.
	//
	// example:
	//
	// ALL
	TriggerType *string `json:"TriggerType,omitempty" xml:"TriggerType,omitempty"`
	// The URL of the trigger.
	//
	// example:
	//
	// https://www.test.com
	TriggerUrl *string `json:"TriggerUrl,omitempty" xml:"TriggerUrl,omitempty"`
}

func (s ListRepoTriggerResponseBodyTriggers) String() string {
	return tea.Prettify(s)
}

func (s ListRepoTriggerResponseBodyTriggers) GoString() string {
	return s.String()
}

func (s *ListRepoTriggerResponseBodyTriggers) SetRepoEvent(v string) *ListRepoTriggerResponseBodyTriggers {
	s.RepoEvent = &v
	return s
}

func (s *ListRepoTriggerResponseBodyTriggers) SetTriggerId(v string) *ListRepoTriggerResponseBodyTriggers {
	s.TriggerId = &v
	return s
}

func (s *ListRepoTriggerResponseBodyTriggers) SetTriggerName(v string) *ListRepoTriggerResponseBodyTriggers {
	s.TriggerName = &v
	return s
}

func (s *ListRepoTriggerResponseBodyTriggers) SetTriggerTag(v string) *ListRepoTriggerResponseBodyTriggers {
	s.TriggerTag = &v
	return s
}

func (s *ListRepoTriggerResponseBodyTriggers) SetTriggerType(v string) *ListRepoTriggerResponseBodyTriggers {
	s.TriggerType = &v
	return s
}

func (s *ListRepoTriggerResponseBodyTriggers) SetTriggerUrl(v string) *ListRepoTriggerResponseBodyTriggers {
	s.TriggerUrl = &v
	return s
}

type ListRepoTriggerResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListRepoTriggerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListRepoTriggerResponse) String() string {
	return tea.Prettify(s)
}

func (s ListRepoTriggerResponse) GoString() string {
	return s.String()
}

func (s *ListRepoTriggerResponse) SetHeaders(v map[string]*string) *ListRepoTriggerResponse {
	s.Headers = v
	return s
}

func (s *ListRepoTriggerResponse) SetStatusCode(v int32) *ListRepoTriggerResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRepoTriggerResponse) SetBody(v *ListRepoTriggerResponseBody) *ListRepoTriggerResponse {
	s.Body = v
	return s
}

type ListRepositoryRequest struct {
	// The ID of the Container Registry instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cri-kmsiwlxxdcva****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page. Maximum value: 100. If you specify a value larger than 100 for this parameter, the system reports a parameter error or uses 100 as the maximum value.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The name of the repository.
	//
	// example:
	//
	// repo-test
	RepoName *string `json:"RepoName,omitempty" xml:"RepoName,omitempty"`
	// The name of the namespace to which the repository belongs.
	//
	// example:
	//
	// repo-namespace-test
	RepoNamespaceName *string `json:"RepoNamespaceName,omitempty" xml:"RepoNamespaceName,omitempty"`
	// example:
	//
	// ALL
	RepoStatus *string `json:"RepoStatus,omitempty" xml:"RepoStatus,omitempty"`
}

func (s ListRepositoryRequest) String() string {
	return tea.Prettify(s)
}

func (s ListRepositoryRequest) GoString() string {
	return s.String()
}

func (s *ListRepositoryRequest) SetInstanceId(v string) *ListRepositoryRequest {
	s.InstanceId = &v
	return s
}

func (s *ListRepositoryRequest) SetPageNo(v int32) *ListRepositoryRequest {
	s.PageNo = &v
	return s
}

func (s *ListRepositoryRequest) SetPageSize(v int32) *ListRepositoryRequest {
	s.PageSize = &v
	return s
}

func (s *ListRepositoryRequest) SetRepoName(v string) *ListRepositoryRequest {
	s.RepoName = &v
	return s
}

func (s *ListRepositoryRequest) SetRepoNamespaceName(v string) *ListRepositoryRequest {
	s.RepoNamespaceName = &v
	return s
}

func (s *ListRepositoryRequest) SetRepoStatus(v string) *ListRepositoryRequest {
	s.RepoStatus = &v
	return s
}

type ListRepositoryResponseBody struct {
	// The return value.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The information about the repositories.
	Repositories []*ListRepositoryResponseBodyRepositories `json:"Repositories,omitempty" xml:"Repositories,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 5241C090-DA69-4B0F-8E3F-2F24FDE1110E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of the queried image repositories.
	//
	// example:
	//
	// 1
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListRepositoryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListRepositoryResponseBody) GoString() string {
	return s.String()
}

func (s *ListRepositoryResponseBody) SetCode(v string) *ListRepositoryResponseBody {
	s.Code = &v
	return s
}

func (s *ListRepositoryResponseBody) SetIsSuccess(v bool) *ListRepositoryResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *ListRepositoryResponseBody) SetPageNo(v int32) *ListRepositoryResponseBody {
	s.PageNo = &v
	return s
}

func (s *ListRepositoryResponseBody) SetPageSize(v int32) *ListRepositoryResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListRepositoryResponseBody) SetRepositories(v []*ListRepositoryResponseBodyRepositories) *ListRepositoryResponseBody {
	s.Repositories = v
	return s
}

func (s *ListRepositoryResponseBody) SetRequestId(v string) *ListRepositoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRepositoryResponseBody) SetTotalCount(v string) *ListRepositoryResponseBody {
	s.TotalCount = &v
	return s
}

type ListRepositoryResponseBodyRepositories struct {
	// The time when the repository was created.
	//
	// example:
	//
	// 1564153576000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The ID of the Container Registry instance to which the repository belongs.
	//
	// example:
	//
	// cri-kmsiwlxxdcv****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The time when the repository was last modified.
	//
	// example:
	//
	// 1564153576000
	ModifiedTime *int64 `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// The type of the repository building. Valid values:
	//
	// 	- `AUTO`: The repository is automatically built.
	//
	// 	- `MANUAL`: The repository is manually built.
	//
	// example:
	//
	// MANUAL
	RepoBuildType *string `json:"RepoBuildType,omitempty" xml:"RepoBuildType,omitempty"`
	// The ID of the repository.
	//
	// example:
	//
	// crr-03cuozrsqhkw****
	RepoId *string `json:"RepoId,omitempty" xml:"RepoId,omitempty"`
	// The name of the repository.
	//
	// example:
	//
	// test
	RepoName *string `json:"RepoName,omitempty" xml:"RepoName,omitempty"`
	// The name of the namespace to which the repository belongs.
	//
	// example:
	//
	// test
	RepoNamespaceName *string `json:"RepoNamespaceName,omitempty" xml:"RepoNamespaceName,omitempty"`
	// The status of the repository.
	//
	// example:
	//
	// NORMAL
	RepoStatus *string `json:"RepoStatus,omitempty" xml:"RepoStatus,omitempty"`
	// The type of the repository. Valid values:
	//
	// 	- `PUBLIC`
	//
	// 	- `PRIVATE`
	//
	// example:
	//
	// PRIVATE
	RepoType *string `json:"RepoType,omitempty" xml:"RepoType,omitempty"`
	// The ID of the resource group to which the repository belongs.
	//
	// example:
	//
	// rg-acfm4n5kzyfxxxx
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The summary of the repository.
	//
	// example:
	//
	// test OK
	Summary *string `json:"Summary,omitempty" xml:"Summary,omitempty"`
	// Indicates whether the feature of image tag immutability is enabled for the repository.
	//
	// example:
	//
	// true
	TagImmutability *bool `json:"TagImmutability,omitempty" xml:"TagImmutability,omitempty"`
}

func (s ListRepositoryResponseBodyRepositories) String() string {
	return tea.Prettify(s)
}

func (s ListRepositoryResponseBodyRepositories) GoString() string {
	return s.String()
}

func (s *ListRepositoryResponseBodyRepositories) SetCreateTime(v int64) *ListRepositoryResponseBodyRepositories {
	s.CreateTime = &v
	return s
}

func (s *ListRepositoryResponseBodyRepositories) SetInstanceId(v string) *ListRepositoryResponseBodyRepositories {
	s.InstanceId = &v
	return s
}

func (s *ListRepositoryResponseBodyRepositories) SetModifiedTime(v int64) *ListRepositoryResponseBodyRepositories {
	s.ModifiedTime = &v
	return s
}

func (s *ListRepositoryResponseBodyRepositories) SetRepoBuildType(v string) *ListRepositoryResponseBodyRepositories {
	s.RepoBuildType = &v
	return s
}

func (s *ListRepositoryResponseBodyRepositories) SetRepoId(v string) *ListRepositoryResponseBodyRepositories {
	s.RepoId = &v
	return s
}

func (s *ListRepositoryResponseBodyRepositories) SetRepoName(v string) *ListRepositoryResponseBodyRepositories {
	s.RepoName = &v
	return s
}

func (s *ListRepositoryResponseBodyRepositories) SetRepoNamespaceName(v string) *ListRepositoryResponseBodyRepositories {
	s.RepoNamespaceName = &v
	return s
}

func (s *ListRepositoryResponseBodyRepositories) SetRepoStatus(v string) *ListRepositoryResponseBodyRepositories {
	s.RepoStatus = &v
	return s
}

func (s *ListRepositoryResponseBodyRepositories) SetRepoType(v string) *ListRepositoryResponseBodyRepositories {
	s.RepoType = &v
	return s
}

func (s *ListRepositoryResponseBodyRepositories) SetResourceGroupId(v string) *ListRepositoryResponseBodyRepositories {
	s.ResourceGroupId = &v
	return s
}

func (s *ListRepositoryResponseBodyRepositories) SetSummary(v string) *ListRepositoryResponseBodyRepositories {
	s.Summary = &v
	return s
}

func (s *ListRepositoryResponseBodyRepositories) SetTagImmutability(v bool) *ListRepositoryResponseBodyRepositories {
	s.TagImmutability = &v
	return s
}

type ListRepositoryResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListRepositoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListRepositoryResponse) String() string {
	return tea.Prettify(s)
}

func (s ListRepositoryResponse) GoString() string {
	return s.String()
}

func (s *ListRepositoryResponse) SetHeaders(v map[string]*string) *ListRepositoryResponse {
	s.Headers = v
	return s
}

func (s *ListRepositoryResponse) SetStatusCode(v int32) *ListRepositoryResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRepositoryResponse) SetBody(v *ListRepositoryResponseBody) *ListRepositoryResponse {
	s.Body = v
	return s
}

type ListScanBaselineByTaskRequest struct {
	// The digest value of the image.
	//
	// example:
	//
	// sha256:1c89806cfaf66d2990e2cf1131ebd56ff24b133745a33abf1228*************
	Digest *string `json:"Digest,omitempty" xml:"Digest,omitempty"`
	// The ID of the Container Registry instance.
	//
	// example:
	//
	// cri-***********
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The level of the baseline risk.
	//
	// example:
	//
	// High
	Level *string `json:"Level,omitempty" xml:"Level,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the image repository.
	//
	// example:
	//
	// crr-**************
	RepoId *string `json:"RepoId,omitempty" xml:"RepoId,omitempty"`
	// The ID of the image scan task.
	//
	// example:
	//
	// 3e526d7e-4b45-4703-b046-***********
	ScanTaskId *string `json:"ScanTaskId,omitempty" xml:"ScanTaskId,omitempty"`
	// The image version.
	//
	// example:
	//
	// 1.1.36
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
}

func (s ListScanBaselineByTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s ListScanBaselineByTaskRequest) GoString() string {
	return s.String()
}

func (s *ListScanBaselineByTaskRequest) SetDigest(v string) *ListScanBaselineByTaskRequest {
	s.Digest = &v
	return s
}

func (s *ListScanBaselineByTaskRequest) SetInstanceId(v string) *ListScanBaselineByTaskRequest {
	s.InstanceId = &v
	return s
}

func (s *ListScanBaselineByTaskRequest) SetLevel(v string) *ListScanBaselineByTaskRequest {
	s.Level = &v
	return s
}

func (s *ListScanBaselineByTaskRequest) SetPageNo(v int32) *ListScanBaselineByTaskRequest {
	s.PageNo = &v
	return s
}

func (s *ListScanBaselineByTaskRequest) SetPageSize(v int32) *ListScanBaselineByTaskRequest {
	s.PageSize = &v
	return s
}

func (s *ListScanBaselineByTaskRequest) SetRepoId(v string) *ListScanBaselineByTaskRequest {
	s.RepoId = &v
	return s
}

func (s *ListScanBaselineByTaskRequest) SetScanTaskId(v string) *ListScanBaselineByTaskRequest {
	s.ScanTaskId = &v
	return s
}

func (s *ListScanBaselineByTaskRequest) SetTag(v string) *ListScanBaselineByTaskRequest {
	s.Tag = &v
	return s
}

type ListScanBaselineByTaskResponseBody struct {
	// The return value.
	//
	// example:
	//
	// success
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the API request was successful. Valid values:
	//
	// 	- `true`: successful
	//
	// 	- `false`: failed
	//
	// example:
	//
	// true
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 5259118F-79E2-57E9-9AEA-551586F4FAED
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The scanned baseline risks.
	ScanBaselines []*ListScanBaselineByTaskResponseBodyScanBaselines `json:"ScanBaselines,omitempty" xml:"ScanBaselines,omitempty" type:"Repeated"`
	// The total number of entries returned.
	//
	// example:
	//
	// 2
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListScanBaselineByTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListScanBaselineByTaskResponseBody) GoString() string {
	return s.String()
}

func (s *ListScanBaselineByTaskResponseBody) SetCode(v int32) *ListScanBaselineByTaskResponseBody {
	s.Code = &v
	return s
}

func (s *ListScanBaselineByTaskResponseBody) SetIsSuccess(v bool) *ListScanBaselineByTaskResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *ListScanBaselineByTaskResponseBody) SetPageNo(v int32) *ListScanBaselineByTaskResponseBody {
	s.PageNo = &v
	return s
}

func (s *ListScanBaselineByTaskResponseBody) SetPageSize(v int32) *ListScanBaselineByTaskResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListScanBaselineByTaskResponseBody) SetRequestId(v string) *ListScanBaselineByTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListScanBaselineByTaskResponseBody) SetScanBaselines(v []*ListScanBaselineByTaskResponseBodyScanBaselines) *ListScanBaselineByTaskResponseBody {
	s.ScanBaselines = v
	return s
}

func (s *ListScanBaselineByTaskResponseBody) SetTotalCount(v int32) *ListScanBaselineByTaskResponseBody {
	s.TotalCount = &v
	return s
}

type ListScanBaselineByTaskResponseBodyScanBaselines struct {
	// The category to which the baseline risk belongs.
	BaselineClassAlias *string `json:"BaselineClassAlias,omitempty" xml:"BaselineClassAlias,omitempty"`
	// Suggestions about how to fix the baseline risk.
	BaselineDetailAdvice *string `json:"BaselineDetailAdvice,omitempty" xml:"BaselineDetailAdvice,omitempty"`
	// The description of the baseline risk.
	BaselineDetailDescription *string `json:"BaselineDetailDescription,omitempty" xml:"BaselineDetailDescription,omitempty"`
	// The path and content of the baseline risk.
	//
	// example:
	//
	// usr/local/www/project/environments/dev/common/config/paramsxxx
	BaselineDetailPrompt *string `json:"BaselineDetailPrompt,omitempty" xml:"BaselineDetailPrompt,omitempty"`
	// The number of baseline risks.
	//
	// example:
	//
	// 1
	BaselineItemCount *int32 `json:"BaselineItemCount,omitempty" xml:"BaselineItemCount,omitempty"`
	// The name of the baseline risk.
	BaselineNameAlias *string `json:"BaselineNameAlias,omitempty" xml:"BaselineNameAlias,omitempty"`
	// The key of the baseline name.
	//
	// example:
	//
	// ak_leak
	BaselineNameKey *string `json:"BaselineNameKey,omitempty" xml:"BaselineNameKey,omitempty"`
	// The severity of the baseline risk.
	//
	// example:
	//
	// high
	BaselineNameLevel *string `json:"BaselineNameLevel,omitempty" xml:"BaselineNameLevel,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 1695090008000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The time of the first scan.
	//
	// example:
	//
	// 2024-04-10 15:33:26
	FirstScanTime *int64 `json:"FirstScanTime,omitempty" xml:"FirstScanTime,omitempty"`
	// High risk quantity.
	//
	// example:
	//
	// 1
	HighRiskItemCount *int32 `json:"HighRiskItemCount,omitempty" xml:"HighRiskItemCount,omitempty"`
	// Low risk quantity.
	//
	// example:
	//
	// 1
	LowRiskItemCount *int32 `json:"LowRiskItemCount,omitempty" xml:"LowRiskItemCount,omitempty"`
	// Medium risk quantity.
	//
	// example:
	//
	// 1
	MiddleRiskItemCount *int32 `json:"MiddleRiskItemCount,omitempty" xml:"MiddleRiskItemCount,omitempty"`
	// The ID of the image scan task.
	//
	// example:
	//
	// 2328fcaa-f28a-405d-a357-asdvfrew23
	ScanTaskId *string `json:"ScanTaskId,omitempty" xml:"ScanTaskId,omitempty"`
	// The update time.
	//
	// example:
	//
	// 1684220824226
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListScanBaselineByTaskResponseBodyScanBaselines) String() string {
	return tea.Prettify(s)
}

func (s ListScanBaselineByTaskResponseBodyScanBaselines) GoString() string {
	return s.String()
}

func (s *ListScanBaselineByTaskResponseBodyScanBaselines) SetBaselineClassAlias(v string) *ListScanBaselineByTaskResponseBodyScanBaselines {
	s.BaselineClassAlias = &v
	return s
}

func (s *ListScanBaselineByTaskResponseBodyScanBaselines) SetBaselineDetailAdvice(v string) *ListScanBaselineByTaskResponseBodyScanBaselines {
	s.BaselineDetailAdvice = &v
	return s
}

func (s *ListScanBaselineByTaskResponseBodyScanBaselines) SetBaselineDetailDescription(v string) *ListScanBaselineByTaskResponseBodyScanBaselines {
	s.BaselineDetailDescription = &v
	return s
}

func (s *ListScanBaselineByTaskResponseBodyScanBaselines) SetBaselineDetailPrompt(v string) *ListScanBaselineByTaskResponseBodyScanBaselines {
	s.BaselineDetailPrompt = &v
	return s
}

func (s *ListScanBaselineByTaskResponseBodyScanBaselines) SetBaselineItemCount(v int32) *ListScanBaselineByTaskResponseBodyScanBaselines {
	s.BaselineItemCount = &v
	return s
}

func (s *ListScanBaselineByTaskResponseBodyScanBaselines) SetBaselineNameAlias(v string) *ListScanBaselineByTaskResponseBodyScanBaselines {
	s.BaselineNameAlias = &v
	return s
}

func (s *ListScanBaselineByTaskResponseBodyScanBaselines) SetBaselineNameKey(v string) *ListScanBaselineByTaskResponseBodyScanBaselines {
	s.BaselineNameKey = &v
	return s
}

func (s *ListScanBaselineByTaskResponseBodyScanBaselines) SetBaselineNameLevel(v string) *ListScanBaselineByTaskResponseBodyScanBaselines {
	s.BaselineNameLevel = &v
	return s
}

func (s *ListScanBaselineByTaskResponseBodyScanBaselines) SetCreateTime(v int64) *ListScanBaselineByTaskResponseBodyScanBaselines {
	s.CreateTime = &v
	return s
}

func (s *ListScanBaselineByTaskResponseBodyScanBaselines) SetFirstScanTime(v int64) *ListScanBaselineByTaskResponseBodyScanBaselines {
	s.FirstScanTime = &v
	return s
}

func (s *ListScanBaselineByTaskResponseBodyScanBaselines) SetHighRiskItemCount(v int32) *ListScanBaselineByTaskResponseBodyScanBaselines {
	s.HighRiskItemCount = &v
	return s
}

func (s *ListScanBaselineByTaskResponseBodyScanBaselines) SetLowRiskItemCount(v int32) *ListScanBaselineByTaskResponseBodyScanBaselines {
	s.LowRiskItemCount = &v
	return s
}

func (s *ListScanBaselineByTaskResponseBodyScanBaselines) SetMiddleRiskItemCount(v int32) *ListScanBaselineByTaskResponseBodyScanBaselines {
	s.MiddleRiskItemCount = &v
	return s
}

func (s *ListScanBaselineByTaskResponseBodyScanBaselines) SetScanTaskId(v string) *ListScanBaselineByTaskResponseBodyScanBaselines {
	s.ScanTaskId = &v
	return s
}

func (s *ListScanBaselineByTaskResponseBodyScanBaselines) SetUpdateTime(v int64) *ListScanBaselineByTaskResponseBodyScanBaselines {
	s.UpdateTime = &v
	return s
}

type ListScanBaselineByTaskResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListScanBaselineByTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListScanBaselineByTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s ListScanBaselineByTaskResponse) GoString() string {
	return s.String()
}

func (s *ListScanBaselineByTaskResponse) SetHeaders(v map[string]*string) *ListScanBaselineByTaskResponse {
	s.Headers = v
	return s
}

func (s *ListScanBaselineByTaskResponse) SetStatusCode(v int32) *ListScanBaselineByTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *ListScanBaselineByTaskResponse) SetBody(v *ListScanBaselineByTaskResponseBody) *ListScanBaselineByTaskResponse {
	s.Body = v
	return s
}

type ListScanMaliciousFileByTaskRequest struct {
	// The image digest.
	//
	// example:
	//
	// sha256:aa4bffff6406785e930dab1f94c9a1297ba22xxxx71d71504a015764*********
	Digest *string `json:"Digest,omitempty" xml:"Digest,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// cri-gu94qynvpwk*****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The severity of the malicious file.
	//
	// example:
	//
	// High
	Level *string `json:"Level,omitempty" xml:"Level,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page. Maximum value: 100. If you specify a value greater than 100 for this parameter, the system reports a parameter error or uses 100 as the maximum value.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The image repository ID.
	//
	// example:
	//
	// crr-h1y4l279wb8*****
	RepoId *string `json:"RepoId,omitempty" xml:"RepoId,omitempty"`
	// The ID of the image scan task.
	//
	// example:
	//
	// 79ee6bc2-3288-4c56-b967-**********
	ScanTaskId *string `json:"ScanTaskId,omitempty" xml:"ScanTaskId,omitempty"`
	// The image tag.
	//
	// example:
	//
	// V6.11
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
}

func (s ListScanMaliciousFileByTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s ListScanMaliciousFileByTaskRequest) GoString() string {
	return s.String()
}

func (s *ListScanMaliciousFileByTaskRequest) SetDigest(v string) *ListScanMaliciousFileByTaskRequest {
	s.Digest = &v
	return s
}

func (s *ListScanMaliciousFileByTaskRequest) SetInstanceId(v string) *ListScanMaliciousFileByTaskRequest {
	s.InstanceId = &v
	return s
}

func (s *ListScanMaliciousFileByTaskRequest) SetLevel(v string) *ListScanMaliciousFileByTaskRequest {
	s.Level = &v
	return s
}

func (s *ListScanMaliciousFileByTaskRequest) SetPageNo(v int32) *ListScanMaliciousFileByTaskRequest {
	s.PageNo = &v
	return s
}

func (s *ListScanMaliciousFileByTaskRequest) SetPageSize(v int32) *ListScanMaliciousFileByTaskRequest {
	s.PageSize = &v
	return s
}

func (s *ListScanMaliciousFileByTaskRequest) SetRepoId(v string) *ListScanMaliciousFileByTaskRequest {
	s.RepoId = &v
	return s
}

func (s *ListScanMaliciousFileByTaskRequest) SetScanTaskId(v string) *ListScanMaliciousFileByTaskRequest {
	s.ScanTaskId = &v
	return s
}

func (s *ListScanMaliciousFileByTaskRequest) SetTag(v string) *ListScanMaliciousFileByTaskRequest {
	s.Tag = &v
	return s
}

type ListScanMaliciousFileByTaskResponseBody struct {
	// The return value.
	//
	// example:
	//
	// success
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the API request is successful. Valid values:
	//
	// 	- `true`: successful
	//
	// 	- `false`: failed
	//
	// example:
	//
	// true
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 52AE49C8-B91A-5C1A-821F-C34324B42F7C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The queried malicious files.
	ScanMaliciousFiles []*ListScanMaliciousFileByTaskResponseBodyScanMaliciousFiles `json:"ScanMaliciousFiles,omitempty" xml:"ScanMaliciousFiles,omitempty" type:"Repeated"`
	// The total number of entries returned.
	//
	// example:
	//
	// 13
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListScanMaliciousFileByTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListScanMaliciousFileByTaskResponseBody) GoString() string {
	return s.String()
}

func (s *ListScanMaliciousFileByTaskResponseBody) SetCode(v int32) *ListScanMaliciousFileByTaskResponseBody {
	s.Code = &v
	return s
}

func (s *ListScanMaliciousFileByTaskResponseBody) SetIsSuccess(v bool) *ListScanMaliciousFileByTaskResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *ListScanMaliciousFileByTaskResponseBody) SetPageNo(v int32) *ListScanMaliciousFileByTaskResponseBody {
	s.PageNo = &v
	return s
}

func (s *ListScanMaliciousFileByTaskResponseBody) SetPageSize(v int32) *ListScanMaliciousFileByTaskResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListScanMaliciousFileByTaskResponseBody) SetRequestId(v string) *ListScanMaliciousFileByTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListScanMaliciousFileByTaskResponseBody) SetScanMaliciousFiles(v []*ListScanMaliciousFileByTaskResponseBodyScanMaliciousFiles) *ListScanMaliciousFileByTaskResponseBody {
	s.ScanMaliciousFiles = v
	return s
}

func (s *ListScanMaliciousFileByTaskResponseBody) SetTotalCount(v int32) *ListScanMaliciousFileByTaskResponseBody {
	s.TotalCount = &v
	return s
}

type ListScanMaliciousFileByTaskResponseBodyScanMaliciousFiles struct {
	// The time when the image was created.
	//
	// example:
	//
	// 2023-04-10 11:42:06
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The path of the file.
	//
	// example:
	//
	// tenant/0000000000000000/
	FilePath *string `json:"FilePath,omitempty" xml:"FilePath,omitempty"`
	// The time when the image was first scanned.
	//
	// example:
	//
	// 2023-04-10 11:42:06
	FirstScanTime *int64 `json:"FirstScanTime,omitempty" xml:"FirstScanTime,omitempty"`
	// The severity of the malicious file.
	//
	// example:
	//
	// remind
	Level *string `json:"Level,omitempty" xml:"Level,omitempty"`
	// The MD5 hash value of the malicious file.
	//
	// example:
	//
	// e76c9759783cbbc9be0ff91ca3xxxxxx
	MaliciousMd5 *string `json:"MaliciousMd5,omitempty" xml:"MaliciousMd5,omitempty"`
	// The type of the malicious file.
	//
	// example:
	//
	// Suspected to contain Webshell code
	MaliciousName *string `json:"MaliciousName,omitempty" xml:"MaliciousName,omitempty"`
	// The ID of the image scan task.
	//
	// example:
	//
	// fe2d8980-de45-4f55-b57d-e438e6d2e972
	ScanTaskId *string `json:"ScanTaskId,omitempty" xml:"ScanTaskId,omitempty"`
	// The time when the image was updated.
	//
	// example:
	//
	// 2023-04-10 11:42:06
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListScanMaliciousFileByTaskResponseBodyScanMaliciousFiles) String() string {
	return tea.Prettify(s)
}

func (s ListScanMaliciousFileByTaskResponseBodyScanMaliciousFiles) GoString() string {
	return s.String()
}

func (s *ListScanMaliciousFileByTaskResponseBodyScanMaliciousFiles) SetCreateTime(v int64) *ListScanMaliciousFileByTaskResponseBodyScanMaliciousFiles {
	s.CreateTime = &v
	return s
}

func (s *ListScanMaliciousFileByTaskResponseBodyScanMaliciousFiles) SetFilePath(v string) *ListScanMaliciousFileByTaskResponseBodyScanMaliciousFiles {
	s.FilePath = &v
	return s
}

func (s *ListScanMaliciousFileByTaskResponseBodyScanMaliciousFiles) SetFirstScanTime(v int64) *ListScanMaliciousFileByTaskResponseBodyScanMaliciousFiles {
	s.FirstScanTime = &v
	return s
}

func (s *ListScanMaliciousFileByTaskResponseBodyScanMaliciousFiles) SetLevel(v string) *ListScanMaliciousFileByTaskResponseBodyScanMaliciousFiles {
	s.Level = &v
	return s
}

func (s *ListScanMaliciousFileByTaskResponseBodyScanMaliciousFiles) SetMaliciousMd5(v string) *ListScanMaliciousFileByTaskResponseBodyScanMaliciousFiles {
	s.MaliciousMd5 = &v
	return s
}

func (s *ListScanMaliciousFileByTaskResponseBodyScanMaliciousFiles) SetMaliciousName(v string) *ListScanMaliciousFileByTaskResponseBodyScanMaliciousFiles {
	s.MaliciousName = &v
	return s
}

func (s *ListScanMaliciousFileByTaskResponseBodyScanMaliciousFiles) SetScanTaskId(v string) *ListScanMaliciousFileByTaskResponseBodyScanMaliciousFiles {
	s.ScanTaskId = &v
	return s
}

func (s *ListScanMaliciousFileByTaskResponseBodyScanMaliciousFiles) SetUpdateTime(v int64) *ListScanMaliciousFileByTaskResponseBodyScanMaliciousFiles {
	s.UpdateTime = &v
	return s
}

type ListScanMaliciousFileByTaskResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListScanMaliciousFileByTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListScanMaliciousFileByTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s ListScanMaliciousFileByTaskResponse) GoString() string {
	return s.String()
}

func (s *ListScanMaliciousFileByTaskResponse) SetHeaders(v map[string]*string) *ListScanMaliciousFileByTaskResponse {
	s.Headers = v
	return s
}

func (s *ListScanMaliciousFileByTaskResponse) SetStatusCode(v int32) *ListScanMaliciousFileByTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *ListScanMaliciousFileByTaskResponse) SetBody(v *ListScanMaliciousFileByTaskResponseBody) *ListScanMaliciousFileByTaskResponse {
	s.Body = v
	return s
}

type ListTagResourcesRequest struct {
	// The pagination token that is used in the next request to retrieve a new page of results. You do not need to specify this parameter for the first request or if no next query exists. You must specify the token that is obtained from the previous query as the value of NextToken.
	//
	// example:
	//
	// AAAAAfj+3fkqd8igM6VLaQjlaYc=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the region in which the resources are deployed.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource IDs. You can specify a maximum of 20 resource IDs.
	ResourceId []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	// The type of the resources. Instance resources are supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// Instance
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The tag list.
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

func (s *ListTagResourcesRequest) SetRegionId(v string) *ListTagResourcesRequest {
	s.RegionId = &v
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
	// The tag key of the resource. You can specify up to 20 tag keys. The tag key cannot be an empty string.
	//
	// A tag key can be up to 128 characters in length. It cannot start with acs: or aliyun, and cannot contain http:// or https://.
	//
	// example:
	//
	// test-key
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the resource. You can specify up to 20 tag values. The tag value can be an empty string.
	//
	// It can be up to 128 characters in length. It cannot start with acs: or aliyun and cannot contain http:// or https://.
	//
	// example:
	//
	// test-val
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
	// A pagination token. It can be used in the next request to retrieve a new page of results. If NextToken is empty, no next page exists.
	//
	// example:
	//
	// AAAAAfj+3fkqd8igM6VLaQjlaYc=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 7FF809ED-B42F-5AC3-9A17-CFE14BE32A8E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The tags that are added to the resource.
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
	// The resource ID.
	//
	// example:
	//
	// cri-w19e7qr2wibl****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The resource type.
	//
	// example:
	//
	// Instance
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The tag key.
	//
	// example:
	//
	// test-key
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The tag value.
	//
	// example:
	//
	// test-val
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

type ResetLoginPasswordRequest struct {
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cri-kmsiwlxxdcva****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The new password that you specify to log on to the instance. The password must be 8 to 32 bits in length, and must contain at least two of the following character types: letters, special characters, and digits.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
}

func (s ResetLoginPasswordRequest) String() string {
	return tea.Prettify(s)
}

func (s ResetLoginPasswordRequest) GoString() string {
	return s.String()
}

func (s *ResetLoginPasswordRequest) SetInstanceId(v string) *ResetLoginPasswordRequest {
	s.InstanceId = &v
	return s
}

func (s *ResetLoginPasswordRequest) SetPassword(v string) *ResetLoginPasswordRequest {
	s.Password = &v
	return s
}

type ResetLoginPasswordResponseBody struct {
	// The return value.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// 	- `true`: The request is successful.
	//
	// 	- `false`: The request fails.
	//
	// example:
	//
	// true
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// EB9C5722-51E2-4497-A573-575B0CA5CE0C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ResetLoginPasswordResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ResetLoginPasswordResponseBody) GoString() string {
	return s.String()
}

func (s *ResetLoginPasswordResponseBody) SetCode(v string) *ResetLoginPasswordResponseBody {
	s.Code = &v
	return s
}

func (s *ResetLoginPasswordResponseBody) SetIsSuccess(v bool) *ResetLoginPasswordResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *ResetLoginPasswordResponseBody) SetRequestId(v string) *ResetLoginPasswordResponseBody {
	s.RequestId = &v
	return s
}

type ResetLoginPasswordResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ResetLoginPasswordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ResetLoginPasswordResponse) String() string {
	return tea.Prettify(s)
}

func (s ResetLoginPasswordResponse) GoString() string {
	return s.String()
}

func (s *ResetLoginPasswordResponse) SetHeaders(v map[string]*string) *ResetLoginPasswordResponse {
	s.Headers = v
	return s
}

func (s *ResetLoginPasswordResponse) SetStatusCode(v int32) *ResetLoginPasswordResponse {
	s.StatusCode = &v
	return s
}

func (s *ResetLoginPasswordResponse) SetBody(v *ResetLoginPasswordResponseBody) *ResetLoginPasswordResponse {
	s.Body = v
	return s
}

type TagResourcesRequest struct {
	// The region ID of the resources.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource IDs. You can specify a maximum of 20 resource IDs.
	//
	// This parameter is required.
	ResourceId []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	// The type of the resource. Instance resources are supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// Instance
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The tag list.
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

func (s *TagResourcesRequest) SetRegionId(v string) *TagResourcesRequest {
	s.RegionId = &v
	return s
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
	// The tag key of the resource. You can specify up to 20 tag keys. The tag key cannot be an empty string.
	//
	// A tag key can be up to 128 characters in length. It cannot start with acs: or aliyun, and cannot contain http:// or https://.
	//
	// example:
	//
	// test-key
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the resource. You can specify up to 20 tag values. The tag value can be an empty string.
	//
	// A tag value can be up to 128 characters in length. It cannot start with acs: or aliyun, and cannot contain http:// or https://.
	//
	// example:
	//
	// test-val
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
	// The request ID.
	//
	// example:
	//
	// E9A586D0-3977-5C28-A44D-55D3A9CD53CC
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

type UntagResourcesRequest struct {
	// Specifies whether to remove all tags from the resource. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// >If you specify both this parameter and the TagKey parameter, this parameter does not take effect.
	//
	// example:
	//
	// true
	All *bool `json:"All,omitempty" xml:"All,omitempty"`
	// The ID of the region in which the resources are deployed.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource IDs. You can specify a maximum of 20 resource IDs.
	ResourceId []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	// The type of the resources. Instance resources are supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// Instance
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The key of tag N added to the resource. Valid values of N: 1 to 20.
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

func (s *UntagResourcesRequest) SetRegionId(v string) *UntagResourcesRequest {
	s.RegionId = &v
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
	// The request ID.
	//
	// example:
	//
	// 724402D0-75CD-4794-BC20-7D3720823AE0
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

type UpdateArtifactLifecycleRuleRequest struct {
	// Specifies whether to automatically execute the lifecycle management rule.
	//
	// example:
	//
	// false
	Auto *bool `json:"Auto,omitempty" xml:"Auto,omitempty"`
	// Specifies whether to enable lifecycle management for the artifact.
	//
	// example:
	//
	// true
	EnableDeleteTag *bool `json:"EnableDeleteTag,omitempty" xml:"EnableDeleteTag,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cri-r6ym0lerldp****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the namespace.
	//
	// example:
	//
	// test-ns
	NamespaceName *string `json:"NamespaceName,omitempty" xml:"NamespaceName,omitempty"`
	// The name of the image repository.
	//
	// example:
	//
	// test_1
	RepoName *string `json:"RepoName,omitempty" xml:"RepoName,omitempty"`
	// The number of images that you want to retain.
	//
	// example:
	//
	// 30
	RetentionTagCount *int64 `json:"RetentionTagCount,omitempty" xml:"RetentionTagCount,omitempty"`
	// The rule ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cralr-luq6qiegzvx****
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// The execution cycle of the lifecycle management rule.
	//
	// example:
	//
	// WEEK
	ScheduleTime *string `json:"ScheduleTime,omitempty" xml:"ScheduleTime,omitempty"`
	// The deletion scope of artifacts.
	//
	// example:
	//
	// REPO
	Scope *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
	// The regular expression that indicates which image tags you want to retain.
	//
	// example:
	//
	// .*production_.*
	TagRegexp *string `json:"TagRegexp,omitempty" xml:"TagRegexp,omitempty"`
}

func (s UpdateArtifactLifecycleRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateArtifactLifecycleRuleRequest) GoString() string {
	return s.String()
}

func (s *UpdateArtifactLifecycleRuleRequest) SetAuto(v bool) *UpdateArtifactLifecycleRuleRequest {
	s.Auto = &v
	return s
}

func (s *UpdateArtifactLifecycleRuleRequest) SetEnableDeleteTag(v bool) *UpdateArtifactLifecycleRuleRequest {
	s.EnableDeleteTag = &v
	return s
}

func (s *UpdateArtifactLifecycleRuleRequest) SetInstanceId(v string) *UpdateArtifactLifecycleRuleRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateArtifactLifecycleRuleRequest) SetNamespaceName(v string) *UpdateArtifactLifecycleRuleRequest {
	s.NamespaceName = &v
	return s
}

func (s *UpdateArtifactLifecycleRuleRequest) SetRepoName(v string) *UpdateArtifactLifecycleRuleRequest {
	s.RepoName = &v
	return s
}

func (s *UpdateArtifactLifecycleRuleRequest) SetRetentionTagCount(v int64) *UpdateArtifactLifecycleRuleRequest {
	s.RetentionTagCount = &v
	return s
}

func (s *UpdateArtifactLifecycleRuleRequest) SetRuleId(v string) *UpdateArtifactLifecycleRuleRequest {
	s.RuleId = &v
	return s
}

func (s *UpdateArtifactLifecycleRuleRequest) SetScheduleTime(v string) *UpdateArtifactLifecycleRuleRequest {
	s.ScheduleTime = &v
	return s
}

func (s *UpdateArtifactLifecycleRuleRequest) SetScope(v string) *UpdateArtifactLifecycleRuleRequest {
	s.Scope = &v
	return s
}

func (s *UpdateArtifactLifecycleRuleRequest) SetTagRegexp(v string) *UpdateArtifactLifecycleRuleRequest {
	s.TagRegexp = &v
	return s
}

type UpdateArtifactLifecycleRuleResponseBody struct {
	// The return value.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the API request is successful. Valid values:
	//
	// 	- `true`: The request is successful.
	//
	// 	- `false`: The request fails.
	//
	// example:
	//
	// true
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The request ID.
	//
	// example:
	//
	// BF92FC2E-455F-5600-A276-D2150A59****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateArtifactLifecycleRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateArtifactLifecycleRuleResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateArtifactLifecycleRuleResponseBody) SetCode(v string) *UpdateArtifactLifecycleRuleResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateArtifactLifecycleRuleResponseBody) SetIsSuccess(v bool) *UpdateArtifactLifecycleRuleResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *UpdateArtifactLifecycleRuleResponseBody) SetRequestId(v string) *UpdateArtifactLifecycleRuleResponseBody {
	s.RequestId = &v
	return s
}

type UpdateArtifactLifecycleRuleResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateArtifactLifecycleRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateArtifactLifecycleRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateArtifactLifecycleRuleResponse) GoString() string {
	return s.String()
}

func (s *UpdateArtifactLifecycleRuleResponse) SetHeaders(v map[string]*string) *UpdateArtifactLifecycleRuleResponse {
	s.Headers = v
	return s
}

func (s *UpdateArtifactLifecycleRuleResponse) SetStatusCode(v int32) *UpdateArtifactLifecycleRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateArtifactLifecycleRuleResponse) SetBody(v *UpdateArtifactLifecycleRuleResponseBody) *UpdateArtifactLifecycleRuleResponse {
	s.Body = v
	return s
}

type UpdateArtifactSubscriptionRuleRequest struct {
	// Specifies whether to enable an acceleration link for image subscription. The subscription acceleration feature is in public preview. The feature is optimized based on scheduling policies and network links to accelerate image subscription.
	//
	// example:
	//
	// true
	Accelerate *string `json:"Accelerate,omitempty" xml:"Accelerate,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cri-c0o11woew0k****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the Container Registry namespace.
	//
	// example:
	//
	// test-ns
	NamespaceName *string `json:"NamespaceName,omitempty" xml:"NamespaceName,omitempty"`
	// Specifies whether to overwrite the original image.
	//
	// example:
	//
	// true
	Override *string `json:"Override,omitempty" xml:"Override,omitempty"`
	// The operating system and architecture. If the source repository contains multi-arch images, only the images with the specified operating system and architecture are subscribed to the destination repository of the Enterprise Edition instance.
	Platform []*string `json:"Platform,omitempty" xml:"Platform,omitempty" type:"Repeated"`
	// The name of the Container Registry repository.
	//
	// example:
	//
	// test-repo
	RepoName *string `json:"RepoName,omitempty" xml:"RepoName,omitempty"`
	// The rule ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// crasr-mdbpung4i1rm****
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// The name of the source namespace.
	//
	// example:
	//
	// library
	SourceNamespaceName *string `json:"SourceNamespaceName,omitempty" xml:"SourceNamespaceName,omitempty"`
	// The source of the artifacts.
	//
	// Valid values:
	//
	// 	- DOCKER_HUB: Docker Hub
	//
	// 	- GCR: GCR
	//
	// 	- QUAY: Quay.io
	//
	// example:
	//
	// DOCKER_HUB
	SourceProvider *string `json:"SourceProvider,omitempty" xml:"SourceProvider,omitempty"`
	// The source repository.
	//
	// example:
	//
	// nginx
	SourceRepoName *string `json:"SourceRepoName,omitempty" xml:"SourceRepoName,omitempty"`
	// The number of subscribed images.
	//
	// if can be null:
	// false
	//
	// example:
	//
	// 1
	TagCount *int64 `json:"TagCount,omitempty" xml:"TagCount,omitempty"`
	// The image tags in the subscription source repository. Regular expressions are supported.
	//
	// example:
	//
	// release-v.*
	TagRegexp *string `json:"TagRegexp,omitempty" xml:"TagRegexp,omitempty"`
}

func (s UpdateArtifactSubscriptionRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateArtifactSubscriptionRuleRequest) GoString() string {
	return s.String()
}

func (s *UpdateArtifactSubscriptionRuleRequest) SetAccelerate(v string) *UpdateArtifactSubscriptionRuleRequest {
	s.Accelerate = &v
	return s
}

func (s *UpdateArtifactSubscriptionRuleRequest) SetInstanceId(v string) *UpdateArtifactSubscriptionRuleRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateArtifactSubscriptionRuleRequest) SetNamespaceName(v string) *UpdateArtifactSubscriptionRuleRequest {
	s.NamespaceName = &v
	return s
}

func (s *UpdateArtifactSubscriptionRuleRequest) SetOverride(v string) *UpdateArtifactSubscriptionRuleRequest {
	s.Override = &v
	return s
}

func (s *UpdateArtifactSubscriptionRuleRequest) SetPlatform(v []*string) *UpdateArtifactSubscriptionRuleRequest {
	s.Platform = v
	return s
}

func (s *UpdateArtifactSubscriptionRuleRequest) SetRepoName(v string) *UpdateArtifactSubscriptionRuleRequest {
	s.RepoName = &v
	return s
}

func (s *UpdateArtifactSubscriptionRuleRequest) SetRuleId(v string) *UpdateArtifactSubscriptionRuleRequest {
	s.RuleId = &v
	return s
}

func (s *UpdateArtifactSubscriptionRuleRequest) SetSourceNamespaceName(v string) *UpdateArtifactSubscriptionRuleRequest {
	s.SourceNamespaceName = &v
	return s
}

func (s *UpdateArtifactSubscriptionRuleRequest) SetSourceProvider(v string) *UpdateArtifactSubscriptionRuleRequest {
	s.SourceProvider = &v
	return s
}

func (s *UpdateArtifactSubscriptionRuleRequest) SetSourceRepoName(v string) *UpdateArtifactSubscriptionRuleRequest {
	s.SourceRepoName = &v
	return s
}

func (s *UpdateArtifactSubscriptionRuleRequest) SetTagCount(v int64) *UpdateArtifactSubscriptionRuleRequest {
	s.TagCount = &v
	return s
}

func (s *UpdateArtifactSubscriptionRuleRequest) SetTagRegexp(v string) *UpdateArtifactSubscriptionRuleRequest {
	s.TagRegexp = &v
	return s
}

type UpdateArtifactSubscriptionRuleResponseBody struct {
	// The return value.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the API request is successful. Valid values:
	//
	// 	- `true`: The request is successful.
	//
	// 	- `false`: The request fails.
	//
	// example:
	//
	// true
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 02B27D80-FD32-5155-931A-93700779BB9E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateArtifactSubscriptionRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateArtifactSubscriptionRuleResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateArtifactSubscriptionRuleResponseBody) SetCode(v string) *UpdateArtifactSubscriptionRuleResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateArtifactSubscriptionRuleResponseBody) SetIsSuccess(v bool) *UpdateArtifactSubscriptionRuleResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *UpdateArtifactSubscriptionRuleResponseBody) SetRequestId(v string) *UpdateArtifactSubscriptionRuleResponseBody {
	s.RequestId = &v
	return s
}

type UpdateArtifactSubscriptionRuleResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateArtifactSubscriptionRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateArtifactSubscriptionRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateArtifactSubscriptionRuleResponse) GoString() string {
	return s.String()
}

func (s *UpdateArtifactSubscriptionRuleResponse) SetHeaders(v map[string]*string) *UpdateArtifactSubscriptionRuleResponse {
	s.Headers = v
	return s
}

func (s *UpdateArtifactSubscriptionRuleResponse) SetStatusCode(v int32) *UpdateArtifactSubscriptionRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateArtifactSubscriptionRuleResponse) SetBody(v *UpdateArtifactSubscriptionRuleResponseBody) *UpdateArtifactSubscriptionRuleResponse {
	s.Body = v
	return s
}

type UpdateChainRequest struct {
	// The configuration of the delivery chain in the JSON format.
	//
	// This parameter is required.
	//
	// example:
	//
	// chainconfig
	ChainConfig *string `json:"ChainConfig,omitempty" xml:"ChainConfig,omitempty"`
	// The ID of the delivery chain.
	//
	// This parameter is required.
	//
	// example:
	//
	// chi-02ymhtwl3cq8****
	ChainId *string `json:"ChainId,omitempty" xml:"ChainId,omitempty"`
	// The description of the delivery chain.
	//
	// example:
	//
	// description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cri-4cdrlqmhn4gm****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the delivery chain.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Repositories in which the delivery chain does not take effect.
	ScopeExclude []*string `json:"ScopeExclude,omitempty" xml:"ScopeExclude,omitempty" type:"Repeated"`
}

func (s UpdateChainRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateChainRequest) GoString() string {
	return s.String()
}

func (s *UpdateChainRequest) SetChainConfig(v string) *UpdateChainRequest {
	s.ChainConfig = &v
	return s
}

func (s *UpdateChainRequest) SetChainId(v string) *UpdateChainRequest {
	s.ChainId = &v
	return s
}

func (s *UpdateChainRequest) SetDescription(v string) *UpdateChainRequest {
	s.Description = &v
	return s
}

func (s *UpdateChainRequest) SetInstanceId(v string) *UpdateChainRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateChainRequest) SetName(v string) *UpdateChainRequest {
	s.Name = &v
	return s
}

func (s *UpdateChainRequest) SetScopeExclude(v []*string) *UpdateChainRequest {
	s.ScopeExclude = v
	return s
}

type UpdateChainResponseBody struct {
	// The return value.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 85A99B10-3926-5201-958E-C06FA47F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateChainResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateChainResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateChainResponseBody) SetCode(v string) *UpdateChainResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateChainResponseBody) SetIsSuccess(v bool) *UpdateChainResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *UpdateChainResponseBody) SetRequestId(v string) *UpdateChainResponseBody {
	s.RequestId = &v
	return s
}

type UpdateChainResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateChainResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateChainResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateChainResponse) GoString() string {
	return s.String()
}

func (s *UpdateChainResponse) SetHeaders(v map[string]*string) *UpdateChainResponse {
	s.Headers = v
	return s
}

func (s *UpdateChainResponse) SetStatusCode(v int32) *UpdateChainResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateChainResponse) SetBody(v *UpdateChainResponseBody) *UpdateChainResponse {
	s.Body = v
	return s
}

type UpdateChartNamespaceRequest struct {
	// Specifies whether to automatically create repositories in the namespace. Valid values:
	//
	// 	- `true`: automatically creates repositories in the namespace.
	//
	// 	- `false`: does not automatically create repositories in the namespace.
	//
	// example:
	//
	// true
	AutoCreateRepo *bool `json:"AutoCreateRepo,omitempty" xml:"AutoCreateRepo,omitempty"`
	// The default type of the repository. Valid values:
	//
	// 	- `PUBLIC`: a public repository
	//
	// 	- `PRIVATE`: a private repository
	//
	// example:
	//
	// PUBLIC
	DefaultRepoType *string `json:"DefaultRepoType,omitempty" xml:"DefaultRepoType,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cri-kmsiwlxxdcva****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the namespace to which the repository belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	NamespaceName *string `json:"NamespaceName,omitempty" xml:"NamespaceName,omitempty"`
}

func (s UpdateChartNamespaceRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateChartNamespaceRequest) GoString() string {
	return s.String()
}

func (s *UpdateChartNamespaceRequest) SetAutoCreateRepo(v bool) *UpdateChartNamespaceRequest {
	s.AutoCreateRepo = &v
	return s
}

func (s *UpdateChartNamespaceRequest) SetDefaultRepoType(v string) *UpdateChartNamespaceRequest {
	s.DefaultRepoType = &v
	return s
}

func (s *UpdateChartNamespaceRequest) SetInstanceId(v string) *UpdateChartNamespaceRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateChartNamespaceRequest) SetNamespaceName(v string) *UpdateChartNamespaceRequest {
	s.NamespaceName = &v
	return s
}

type UpdateChartNamespaceResponseBody struct {
	// The return value.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the API request is successful. Valid values:
	//
	// 	- `true`: successful
	//
	// 	- `false`: failed
	//
	// example:
	//
	// true
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 6877B80A-2895-44C4-BC9E-703B157DEE66
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateChartNamespaceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateChartNamespaceResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateChartNamespaceResponseBody) SetCode(v string) *UpdateChartNamespaceResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateChartNamespaceResponseBody) SetIsSuccess(v bool) *UpdateChartNamespaceResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *UpdateChartNamespaceResponseBody) SetRequestId(v string) *UpdateChartNamespaceResponseBody {
	s.RequestId = &v
	return s
}

type UpdateChartNamespaceResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateChartNamespaceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateChartNamespaceResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateChartNamespaceResponse) GoString() string {
	return s.String()
}

func (s *UpdateChartNamespaceResponse) SetHeaders(v map[string]*string) *UpdateChartNamespaceResponse {
	s.Headers = v
	return s
}

func (s *UpdateChartNamespaceResponse) SetStatusCode(v int32) *UpdateChartNamespaceResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateChartNamespaceResponse) SetBody(v *UpdateChartNamespaceResponseBody) *UpdateChartNamespaceResponse {
	s.Body = v
	return s
}

type UpdateChartRepositoryRequest struct {
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cri-kmsiwlxxdcva****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the repository.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	RepoName *string `json:"RepoName,omitempty" xml:"RepoName,omitempty"`
	// The name of the namespace to which the repository belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	RepoNamespaceName *string `json:"RepoNamespaceName,omitempty" xml:"RepoNamespaceName,omitempty"`
	// The type of the repository. Valid values:
	//
	// 	- `PUBLIC`: a public repository.
	//
	// 	- `PRIVATE`: a private repository.
	//
	// example:
	//
	// PUBLIC
	RepoType *string `json:"RepoType,omitempty" xml:"RepoType,omitempty"`
	// The summary of the repository.
	//
	// example:
	//
	// test
	Summary *string `json:"Summary,omitempty" xml:"Summary,omitempty"`
}

func (s UpdateChartRepositoryRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateChartRepositoryRequest) GoString() string {
	return s.String()
}

func (s *UpdateChartRepositoryRequest) SetInstanceId(v string) *UpdateChartRepositoryRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateChartRepositoryRequest) SetRepoName(v string) *UpdateChartRepositoryRequest {
	s.RepoName = &v
	return s
}

func (s *UpdateChartRepositoryRequest) SetRepoNamespaceName(v string) *UpdateChartRepositoryRequest {
	s.RepoNamespaceName = &v
	return s
}

func (s *UpdateChartRepositoryRequest) SetRepoType(v string) *UpdateChartRepositoryRequest {
	s.RepoType = &v
	return s
}

func (s *UpdateChartRepositoryRequest) SetSummary(v string) *UpdateChartRepositoryRequest {
	s.Summary = &v
	return s
}

type UpdateChartRepositoryResponseBody struct {
	// The return value.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// 	- `true`: The request is successful.
	//
	// 	- `false`: The request fails.
	//
	// example:
	//
	// true
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// EB9C5722-51E2-4497-A573-575B0CA5CE0C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateChartRepositoryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateChartRepositoryResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateChartRepositoryResponseBody) SetCode(v string) *UpdateChartRepositoryResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateChartRepositoryResponseBody) SetIsSuccess(v bool) *UpdateChartRepositoryResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *UpdateChartRepositoryResponseBody) SetRequestId(v string) *UpdateChartRepositoryResponseBody {
	s.RequestId = &v
	return s
}

type UpdateChartRepositoryResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateChartRepositoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateChartRepositoryResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateChartRepositoryResponse) GoString() string {
	return s.String()
}

func (s *UpdateChartRepositoryResponse) SetHeaders(v map[string]*string) *UpdateChartRepositoryResponse {
	s.Headers = v
	return s
}

func (s *UpdateChartRepositoryResponse) SetStatusCode(v int32) *UpdateChartRepositoryResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateChartRepositoryResponse) SetBody(v *UpdateChartRepositoryResponseBody) *UpdateChartRepositoryResponse {
	s.Body = v
	return s
}

type UpdateEventCenterRuleRequest struct {
	// The event notification channel.
	//
	// example:
	//
	// EVENT_BRIDGE
	EventChannel *string `json:"EventChannel,omitempty" xml:"EventChannel,omitempty"`
	// The event configuration.
	//
	// example:
	//
	// {
	//
	//         "notifyMethod":"http",
	//
	//         "notifyConfig":{
	//
	//             "Url":"http://www.aliyundoc.com",
	//
	//             "id":"MaAV3HgTkO5Fh8l1V********",
	//
	//         },
	//
	//         "notifyFilter":{}
	//
	//     }
	EventConfig *string `json:"EventConfig,omitempty" xml:"EventConfig,omitempty"`
	// The event scope. Valid values:
	//
	// 	- `INSTANCE`
	//
	// 	- `NAMESPACE`
	//
	// 	- `REPO`
	//
	// Default value: `INSTANCE`
	//
	// example:
	//
	// INSTANCE
	EventScope *string `json:"EventScope,omitempty" xml:"EventScope,omitempty"`
	// The type of the event. Valid values:
	//
	// 	- `cr:Artifact:DeliveryChainCompleted`: The delivery chain is processed.
	//
	// 	- `cr:Artifact:SynchronizationCompleted`: The image is replicated.
	//
	// 	- `cr:Artifact:BuildCompleted`: The image is built.
	//
	// 	- `cr:Artifact:ScanCompleted`: The image is scanned.
	//
	// 	- `cr:Artifact:SigningCompleted`: The image is signed.
	//
	// example:
	//
	// cr:Artifact:DeliveryChainCompleted
	EventType *string `json:"EventType,omitempty" xml:"EventType,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cri-kmsiwlxxdcva****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The namespaces to which the event notification rule applies.
	//
	// example:
	//
	// ns
	Namespaces []*string `json:"Namespaces,omitempty" xml:"Namespaces,omitempty" type:"Repeated"`
	// The names of the repositories to which the event notification rule applies.
	//
	// example:
	//
	// reponame
	RepoNames []*string `json:"RepoNames,omitempty" xml:"RepoNames,omitempty" type:"Repeated"`
	// The regular expression for image tags.
	//
	// example:
	//
	// .*
	RepoTagFilterPattern *string `json:"RepoTagFilterPattern,omitempty" xml:"RepoTagFilterPattern,omitempty"`
	// The ID of the event notification rule.
	//
	// This parameter is required.
	//
	// example:
	//
	// crecr-n6pbhgjxt*****
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// The name of the event notification rule.
	//
	// example:
	//
	// chain-demo
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
}

func (s UpdateEventCenterRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateEventCenterRuleRequest) GoString() string {
	return s.String()
}

func (s *UpdateEventCenterRuleRequest) SetEventChannel(v string) *UpdateEventCenterRuleRequest {
	s.EventChannel = &v
	return s
}

func (s *UpdateEventCenterRuleRequest) SetEventConfig(v string) *UpdateEventCenterRuleRequest {
	s.EventConfig = &v
	return s
}

func (s *UpdateEventCenterRuleRequest) SetEventScope(v string) *UpdateEventCenterRuleRequest {
	s.EventScope = &v
	return s
}

func (s *UpdateEventCenterRuleRequest) SetEventType(v string) *UpdateEventCenterRuleRequest {
	s.EventType = &v
	return s
}

func (s *UpdateEventCenterRuleRequest) SetInstanceId(v string) *UpdateEventCenterRuleRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateEventCenterRuleRequest) SetNamespaces(v []*string) *UpdateEventCenterRuleRequest {
	s.Namespaces = v
	return s
}

func (s *UpdateEventCenterRuleRequest) SetRepoNames(v []*string) *UpdateEventCenterRuleRequest {
	s.RepoNames = v
	return s
}

func (s *UpdateEventCenterRuleRequest) SetRepoTagFilterPattern(v string) *UpdateEventCenterRuleRequest {
	s.RepoTagFilterPattern = &v
	return s
}

func (s *UpdateEventCenterRuleRequest) SetRuleId(v string) *UpdateEventCenterRuleRequest {
	s.RuleId = &v
	return s
}

func (s *UpdateEventCenterRuleRequest) SetRuleName(v string) *UpdateEventCenterRuleRequest {
	s.RuleName = &v
	return s
}

type UpdateEventCenterRuleShrinkRequest struct {
	// The event notification channel.
	//
	// example:
	//
	// EVENT_BRIDGE
	EventChannel *string `json:"EventChannel,omitempty" xml:"EventChannel,omitempty"`
	// The event configuration.
	//
	// example:
	//
	// {
	//
	//         "notifyMethod":"http",
	//
	//         "notifyConfig":{
	//
	//             "Url":"http://www.aliyundoc.com",
	//
	//             "id":"MaAV3HgTkO5Fh8l1V********",
	//
	//         },
	//
	//         "notifyFilter":{}
	//
	//     }
	EventConfig *string `json:"EventConfig,omitempty" xml:"EventConfig,omitempty"`
	// The event scope. Valid values:
	//
	// 	- `INSTANCE`
	//
	// 	- `NAMESPACE`
	//
	// 	- `REPO`
	//
	// Default value: `INSTANCE`
	//
	// example:
	//
	// INSTANCE
	EventScope *string `json:"EventScope,omitempty" xml:"EventScope,omitempty"`
	// The type of the event. Valid values:
	//
	// 	- `cr:Artifact:DeliveryChainCompleted`: The delivery chain is processed.
	//
	// 	- `cr:Artifact:SynchronizationCompleted`: The image is replicated.
	//
	// 	- `cr:Artifact:BuildCompleted`: The image is built.
	//
	// 	- `cr:Artifact:ScanCompleted`: The image is scanned.
	//
	// 	- `cr:Artifact:SigningCompleted`: The image is signed.
	//
	// example:
	//
	// cr:Artifact:DeliveryChainCompleted
	EventType *string `json:"EventType,omitempty" xml:"EventType,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cri-kmsiwlxxdcva****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The namespaces to which the event notification rule applies.
	//
	// example:
	//
	// ns
	NamespacesShrink *string `json:"Namespaces,omitempty" xml:"Namespaces,omitempty"`
	// The names of the repositories to which the event notification rule applies.
	//
	// example:
	//
	// reponame
	RepoNamesShrink *string `json:"RepoNames,omitempty" xml:"RepoNames,omitempty"`
	// The regular expression for image tags.
	//
	// example:
	//
	// .*
	RepoTagFilterPattern *string `json:"RepoTagFilterPattern,omitempty" xml:"RepoTagFilterPattern,omitempty"`
	// The ID of the event notification rule.
	//
	// This parameter is required.
	//
	// example:
	//
	// crecr-n6pbhgjxt*****
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// The name of the event notification rule.
	//
	// example:
	//
	// chain-demo
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
}

func (s UpdateEventCenterRuleShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateEventCenterRuleShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateEventCenterRuleShrinkRequest) SetEventChannel(v string) *UpdateEventCenterRuleShrinkRequest {
	s.EventChannel = &v
	return s
}

func (s *UpdateEventCenterRuleShrinkRequest) SetEventConfig(v string) *UpdateEventCenterRuleShrinkRequest {
	s.EventConfig = &v
	return s
}

func (s *UpdateEventCenterRuleShrinkRequest) SetEventScope(v string) *UpdateEventCenterRuleShrinkRequest {
	s.EventScope = &v
	return s
}

func (s *UpdateEventCenterRuleShrinkRequest) SetEventType(v string) *UpdateEventCenterRuleShrinkRequest {
	s.EventType = &v
	return s
}

func (s *UpdateEventCenterRuleShrinkRequest) SetInstanceId(v string) *UpdateEventCenterRuleShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateEventCenterRuleShrinkRequest) SetNamespacesShrink(v string) *UpdateEventCenterRuleShrinkRequest {
	s.NamespacesShrink = &v
	return s
}

func (s *UpdateEventCenterRuleShrinkRequest) SetRepoNamesShrink(v string) *UpdateEventCenterRuleShrinkRequest {
	s.RepoNamesShrink = &v
	return s
}

func (s *UpdateEventCenterRuleShrinkRequest) SetRepoTagFilterPattern(v string) *UpdateEventCenterRuleShrinkRequest {
	s.RepoTagFilterPattern = &v
	return s
}

func (s *UpdateEventCenterRuleShrinkRequest) SetRuleId(v string) *UpdateEventCenterRuleShrinkRequest {
	s.RuleId = &v
	return s
}

func (s *UpdateEventCenterRuleShrinkRequest) SetRuleName(v string) *UpdateEventCenterRuleShrinkRequest {
	s.RuleName = &v
	return s
}

type UpdateEventCenterRuleResponseBody struct {
	// The status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 031572FA-7D8F-4C05-B790-1071E0E05DE6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the event notification rule.
	//
	// example:
	//
	// crecr-n6pbhgjxt*****
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s UpdateEventCenterRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateEventCenterRuleResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateEventCenterRuleResponseBody) SetCode(v int32) *UpdateEventCenterRuleResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateEventCenterRuleResponseBody) SetRequestId(v string) *UpdateEventCenterRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateEventCenterRuleResponseBody) SetRuleId(v string) *UpdateEventCenterRuleResponseBody {
	s.RuleId = &v
	return s
}

type UpdateEventCenterRuleResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateEventCenterRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateEventCenterRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateEventCenterRuleResponse) GoString() string {
	return s.String()
}

func (s *UpdateEventCenterRuleResponse) SetHeaders(v map[string]*string) *UpdateEventCenterRuleResponse {
	s.Headers = v
	return s
}

func (s *UpdateEventCenterRuleResponse) SetStatusCode(v int32) *UpdateEventCenterRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateEventCenterRuleResponse) SetBody(v *UpdateEventCenterRuleResponseBody) *UpdateEventCenterRuleResponse {
	s.Body = v
	return s
}

type UpdateInstanceEndpointStatusRequest struct {
	// Specifies whether to enable the instance endpoint. Valid values:
	//
	// 	- `true`: enables the instance endpoint.
	//
	// 	- `false`: disables the instance endpoint
	//
	// This parameter is required.
	//
	// example:
	//
	// false
	Enable *bool `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// The type of the endpoint. Set the value to Internet.
	//
	// This parameter is required.
	//
	// example:
	//
	// internet
	EndpointType *string `json:"EndpointType,omitempty" xml:"EndpointType,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cri-kmsiwlxxdcva****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the module that you want to access. Valid values:
	//
	// 	- `Registry`: the image repository.
	//
	// 	- `Chart`: a Helm chart.
	//
	// example:
	//
	// Chart
	ModuleName *string `json:"ModuleName,omitempty" xml:"ModuleName,omitempty"`
}

func (s UpdateInstanceEndpointStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateInstanceEndpointStatusRequest) GoString() string {
	return s.String()
}

func (s *UpdateInstanceEndpointStatusRequest) SetEnable(v bool) *UpdateInstanceEndpointStatusRequest {
	s.Enable = &v
	return s
}

func (s *UpdateInstanceEndpointStatusRequest) SetEndpointType(v string) *UpdateInstanceEndpointStatusRequest {
	s.EndpointType = &v
	return s
}

func (s *UpdateInstanceEndpointStatusRequest) SetInstanceId(v string) *UpdateInstanceEndpointStatusRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateInstanceEndpointStatusRequest) SetModuleName(v string) *UpdateInstanceEndpointStatusRequest {
	s.ModuleName = &v
	return s
}

type UpdateInstanceEndpointStatusResponseBody struct {
	// The return value.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// 	- `true`: The request is successful.
	//
	// 	- `false`: The request fails.
	//
	// example:
	//
	// true
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 2FC14396-A16A-42BA-AAE4-BB94D956DF09
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateInstanceEndpointStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateInstanceEndpointStatusResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateInstanceEndpointStatusResponseBody) SetCode(v string) *UpdateInstanceEndpointStatusResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateInstanceEndpointStatusResponseBody) SetIsSuccess(v bool) *UpdateInstanceEndpointStatusResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *UpdateInstanceEndpointStatusResponseBody) SetRequestId(v string) *UpdateInstanceEndpointStatusResponseBody {
	s.RequestId = &v
	return s
}

type UpdateInstanceEndpointStatusResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateInstanceEndpointStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateInstanceEndpointStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateInstanceEndpointStatusResponse) GoString() string {
	return s.String()
}

func (s *UpdateInstanceEndpointStatusResponse) SetHeaders(v map[string]*string) *UpdateInstanceEndpointStatusResponse {
	s.Headers = v
	return s
}

func (s *UpdateInstanceEndpointStatusResponse) SetStatusCode(v int32) *UpdateInstanceEndpointStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateInstanceEndpointStatusResponse) SetBody(v *UpdateInstanceEndpointStatusResponseBody) *UpdateInstanceEndpointStatusResponse {
	s.Body = v
	return s
}

type UpdateNamespaceRequest struct {
	// Specifies whether to automatically create a repository when an image is pushed to the namespace.
	//
	// example:
	//
	// true
	AutoCreateRepo           *bool              `json:"AutoCreateRepo,omitempty" xml:"AutoCreateRepo,omitempty"`
	DefaultRepoConfiguration *RepoConfiguration `json:"DefaultRepoConfiguration,omitempty" xml:"DefaultRepoConfiguration,omitempty"`
	// Deprecated
	//
	// The default type of the repository. Valid values:
	//
	// 	- `PUBLIC`: The repository is a public repository.
	//
	// 	- `PRIVATE`: The repository is a private repository.
	//
	// example:
	//
	// PRIVATE
	DefaultRepoType *string `json:"DefaultRepoType,omitempty" xml:"DefaultRepoType,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cri-kmsiwlxxdcva****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the namespace.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	NamespaceName *string `json:"NamespaceName,omitempty" xml:"NamespaceName,omitempty"`
}

func (s UpdateNamespaceRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateNamespaceRequest) GoString() string {
	return s.String()
}

func (s *UpdateNamespaceRequest) SetAutoCreateRepo(v bool) *UpdateNamespaceRequest {
	s.AutoCreateRepo = &v
	return s
}

func (s *UpdateNamespaceRequest) SetDefaultRepoConfiguration(v *RepoConfiguration) *UpdateNamespaceRequest {
	s.DefaultRepoConfiguration = v
	return s
}

func (s *UpdateNamespaceRequest) SetDefaultRepoType(v string) *UpdateNamespaceRequest {
	s.DefaultRepoType = &v
	return s
}

func (s *UpdateNamespaceRequest) SetInstanceId(v string) *UpdateNamespaceRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateNamespaceRequest) SetNamespaceName(v string) *UpdateNamespaceRequest {
	s.NamespaceName = &v
	return s
}

type UpdateNamespaceShrinkRequest struct {
	// Specifies whether to automatically create a repository when an image is pushed to the namespace.
	//
	// example:
	//
	// true
	AutoCreateRepo                 *bool   `json:"AutoCreateRepo,omitempty" xml:"AutoCreateRepo,omitempty"`
	DefaultRepoConfigurationShrink *string `json:"DefaultRepoConfiguration,omitempty" xml:"DefaultRepoConfiguration,omitempty"`
	// Deprecated
	//
	// The default type of the repository. Valid values:
	//
	// 	- `PUBLIC`: The repository is a public repository.
	//
	// 	- `PRIVATE`: The repository is a private repository.
	//
	// example:
	//
	// PRIVATE
	DefaultRepoType *string `json:"DefaultRepoType,omitempty" xml:"DefaultRepoType,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cri-kmsiwlxxdcva****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the namespace.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	NamespaceName *string `json:"NamespaceName,omitempty" xml:"NamespaceName,omitempty"`
}

func (s UpdateNamespaceShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateNamespaceShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateNamespaceShrinkRequest) SetAutoCreateRepo(v bool) *UpdateNamespaceShrinkRequest {
	s.AutoCreateRepo = &v
	return s
}

func (s *UpdateNamespaceShrinkRequest) SetDefaultRepoConfigurationShrink(v string) *UpdateNamespaceShrinkRequest {
	s.DefaultRepoConfigurationShrink = &v
	return s
}

func (s *UpdateNamespaceShrinkRequest) SetDefaultRepoType(v string) *UpdateNamespaceShrinkRequest {
	s.DefaultRepoType = &v
	return s
}

func (s *UpdateNamespaceShrinkRequest) SetInstanceId(v string) *UpdateNamespaceShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateNamespaceShrinkRequest) SetNamespaceName(v string) *UpdateNamespaceShrinkRequest {
	s.NamespaceName = &v
	return s
}

type UpdateNamespaceResponseBody struct {
	// The return value.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 90B8475C-C066-4B92-946E-4D0DECB514E8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateNamespaceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateNamespaceResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateNamespaceResponseBody) SetCode(v string) *UpdateNamespaceResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateNamespaceResponseBody) SetIsSuccess(v bool) *UpdateNamespaceResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *UpdateNamespaceResponseBody) SetRequestId(v string) *UpdateNamespaceResponseBody {
	s.RequestId = &v
	return s
}

type UpdateNamespaceResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateNamespaceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateNamespaceResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateNamespaceResponse) GoString() string {
	return s.String()
}

func (s *UpdateNamespaceResponse) SetHeaders(v map[string]*string) *UpdateNamespaceResponse {
	s.Headers = v
	return s
}

func (s *UpdateNamespaceResponse) SetStatusCode(v int32) *UpdateNamespaceResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateNamespaceResponse) SetBody(v *UpdateNamespaceResponseBody) *UpdateNamespaceResponse {
	s.Body = v
	return s
}

type UpdateRepoBuildRuleRequest struct {
	// Building arguments.
	BuildArgs []*string `json:"BuildArgs,omitempty" xml:"BuildArgs,omitempty" type:"Repeated"`
	// The ID of the building rule.
	//
	// This parameter is required.
	//
	// example:
	//
	// crbr-ly77w5i3t31f****
	BuildRuleId *string `json:"BuildRuleId,omitempty" xml:"BuildRuleId,omitempty"`
	// The path of the Dockerfile.
	//
	// example:
	//
	// /
	DockerfileLocation *string `json:"DockerfileLocation,omitempty" xml:"DockerfileLocation,omitempty"`
	// The name of the Dockerfile.
	//
	// example:
	//
	// Dockerfile
	DockerfileName *string `json:"DockerfileName,omitempty" xml:"DockerfileName,omitempty"`
	// The tag of the image.
	//
	// example:
	//
	// v0.9.5
	ImageTag *string `json:"ImageTag,omitempty" xml:"ImageTag,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cri-kmsiwlxxdcva****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Architecture for image building. Valid values:
	//
	// 	- `linux/amd64`
	//
	// 	- `linux/arm64`
	//
	// 	- `linux/386`
	//
	// 	- `linux/arm/v7`
	//
	// 	- `linux/arm/v6`
	//
	// Default value: `linux/amd64`
	//
	// example:
	//
	// linux/amd64
	Platforms []*string `json:"Platforms,omitempty" xml:"Platforms,omitempty" type:"Repeated"`
	// The name of the push that triggers the building rule.
	//
	// example:
	//
	// master
	PushName *string `json:"PushName,omitempty" xml:"PushName,omitempty"`
	// The type of the push that triggers the building rule. Valid values:
	//
	// 	- `GIT_TAG`: tag push
	//
	// 	- `GIT_BRANCH`: branch push
	//
	// example:
	//
	// GIT_BRANCH
	PushType *string `json:"PushType,omitempty" xml:"PushType,omitempty"`
	// The ID of the image repository.
	//
	// This parameter is required.
	//
	// example:
	//
	// crr-tquyps22md8p****
	RepoId *string `json:"RepoId,omitempty" xml:"RepoId,omitempty"`
}

func (s UpdateRepoBuildRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateRepoBuildRuleRequest) GoString() string {
	return s.String()
}

func (s *UpdateRepoBuildRuleRequest) SetBuildArgs(v []*string) *UpdateRepoBuildRuleRequest {
	s.BuildArgs = v
	return s
}

func (s *UpdateRepoBuildRuleRequest) SetBuildRuleId(v string) *UpdateRepoBuildRuleRequest {
	s.BuildRuleId = &v
	return s
}

func (s *UpdateRepoBuildRuleRequest) SetDockerfileLocation(v string) *UpdateRepoBuildRuleRequest {
	s.DockerfileLocation = &v
	return s
}

func (s *UpdateRepoBuildRuleRequest) SetDockerfileName(v string) *UpdateRepoBuildRuleRequest {
	s.DockerfileName = &v
	return s
}

func (s *UpdateRepoBuildRuleRequest) SetImageTag(v string) *UpdateRepoBuildRuleRequest {
	s.ImageTag = &v
	return s
}

func (s *UpdateRepoBuildRuleRequest) SetInstanceId(v string) *UpdateRepoBuildRuleRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateRepoBuildRuleRequest) SetPlatforms(v []*string) *UpdateRepoBuildRuleRequest {
	s.Platforms = v
	return s
}

func (s *UpdateRepoBuildRuleRequest) SetPushName(v string) *UpdateRepoBuildRuleRequest {
	s.PushName = &v
	return s
}

func (s *UpdateRepoBuildRuleRequest) SetPushType(v string) *UpdateRepoBuildRuleRequest {
	s.PushType = &v
	return s
}

func (s *UpdateRepoBuildRuleRequest) SetRepoId(v string) *UpdateRepoBuildRuleRequest {
	s.RepoId = &v
	return s
}

type UpdateRepoBuildRuleResponseBody struct {
	// The ID of the building rule.
	//
	// example:
	//
	// crbr-ly77w5i3t31f****
	BuildRuleId *string `json:"BuildRuleId,omitempty" xml:"BuildRuleId,omitempty"`
	// The return value.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// 	- `true`: The request is successful.
	//
	// 	- `false`: The request fails.
	//
	// example:
	//
	// true
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// BC648259-91A7-4502-BED3-EDF64361FA83
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateRepoBuildRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateRepoBuildRuleResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateRepoBuildRuleResponseBody) SetBuildRuleId(v string) *UpdateRepoBuildRuleResponseBody {
	s.BuildRuleId = &v
	return s
}

func (s *UpdateRepoBuildRuleResponseBody) SetCode(v string) *UpdateRepoBuildRuleResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateRepoBuildRuleResponseBody) SetIsSuccess(v bool) *UpdateRepoBuildRuleResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *UpdateRepoBuildRuleResponseBody) SetRequestId(v string) *UpdateRepoBuildRuleResponseBody {
	s.RequestId = &v
	return s
}

type UpdateRepoBuildRuleResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateRepoBuildRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateRepoBuildRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateRepoBuildRuleResponse) GoString() string {
	return s.String()
}

func (s *UpdateRepoBuildRuleResponse) SetHeaders(v map[string]*string) *UpdateRepoBuildRuleResponse {
	s.Headers = v
	return s
}

func (s *UpdateRepoBuildRuleResponse) SetStatusCode(v int32) *UpdateRepoBuildRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateRepoBuildRuleResponse) SetBody(v *UpdateRepoBuildRuleResponseBody) *UpdateRepoBuildRuleResponse {
	s.Body = v
	return s
}

type UpdateRepoSourceCodeRepoRequest struct {
	// Specifies whether to enable automatic image building when code is committed. Valid values:
	//
	// 	- `true`: enables automatic image building when code is committed.
	//
	// 	- `false`: disables automatic image building when code is committed.
	//
	// example:
	//
	// true
	AutoBuild *string `json:"AutoBuild,omitempty" xml:"AutoBuild,omitempty"`
	// The ID of the source code repository.
	//
	// example:
	//
	// crr-cp7d6sget5r****
	CodeRepoId *string `json:"CodeRepoId,omitempty" xml:"CodeRepoId,omitempty"`
	// The name of the source code repository.
	//
	// This parameter is required.
	//
	// example:
	//
	// repo
	CodeRepoName *string `json:"CodeRepoName,omitempty" xml:"CodeRepoName,omitempty"`
	// The namespace to which the source code repository belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// namespace
	CodeRepoNamespaceName *string `json:"CodeRepoNamespaceName,omitempty" xml:"CodeRepoNamespaceName,omitempty"`
	// The type of the source code hosting platform. Valid values: GITHUB, GITLAB, GITEE, CODE, and CODEUP.
	//
	// This parameter is required.
	//
	// example:
	//
	// GITHUB
	CodeRepoType *string `json:"CodeRepoType,omitempty" xml:"CodeRepoType,omitempty"`
	// Specifies whether to disable building caches. Valid values:
	//
	// 	- `true`: disables building caches.
	//
	// 	- `false`: enables building caches.
	//
	// example:
	//
	// false
	DisableCacheBuild *string `json:"DisableCacheBuild,omitempty" xml:"DisableCacheBuild,omitempty"`
	// The ID of the Container Registry Enterprise Edition instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cri-shac42yvqzvq****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Specifies whether to enable Build With Servers Deployed Outside Chinese Mainland. Valid values:
	//
	// 	- `true`: enables Build With Servers Deployed Outside Chinese Mainland.
	//
	// 	- `false`: disables Build With Servers Deployed Outside Chinese Mainland.
	//
	// example:
	//
	// false
	OverseaBuild *string `json:"OverseaBuild,omitempty" xml:"OverseaBuild,omitempty"`
	// The ID of the image repository.
	//
	// This parameter is required.
	//
	// example:
	//
	// crr-gzsrlevmvoa****
	RepoId *string `json:"RepoId,omitempty" xml:"RepoId,omitempty"`
}

func (s UpdateRepoSourceCodeRepoRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateRepoSourceCodeRepoRequest) GoString() string {
	return s.String()
}

func (s *UpdateRepoSourceCodeRepoRequest) SetAutoBuild(v string) *UpdateRepoSourceCodeRepoRequest {
	s.AutoBuild = &v
	return s
}

func (s *UpdateRepoSourceCodeRepoRequest) SetCodeRepoId(v string) *UpdateRepoSourceCodeRepoRequest {
	s.CodeRepoId = &v
	return s
}

func (s *UpdateRepoSourceCodeRepoRequest) SetCodeRepoName(v string) *UpdateRepoSourceCodeRepoRequest {
	s.CodeRepoName = &v
	return s
}

func (s *UpdateRepoSourceCodeRepoRequest) SetCodeRepoNamespaceName(v string) *UpdateRepoSourceCodeRepoRequest {
	s.CodeRepoNamespaceName = &v
	return s
}

func (s *UpdateRepoSourceCodeRepoRequest) SetCodeRepoType(v string) *UpdateRepoSourceCodeRepoRequest {
	s.CodeRepoType = &v
	return s
}

func (s *UpdateRepoSourceCodeRepoRequest) SetDisableCacheBuild(v string) *UpdateRepoSourceCodeRepoRequest {
	s.DisableCacheBuild = &v
	return s
}

func (s *UpdateRepoSourceCodeRepoRequest) SetInstanceId(v string) *UpdateRepoSourceCodeRepoRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateRepoSourceCodeRepoRequest) SetOverseaBuild(v string) *UpdateRepoSourceCodeRepoRequest {
	s.OverseaBuild = &v
	return s
}

func (s *UpdateRepoSourceCodeRepoRequest) SetRepoId(v string) *UpdateRepoSourceCodeRepoRequest {
	s.RepoId = &v
	return s
}

type UpdateRepoSourceCodeRepoResponseBody struct {
	// The return value.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// F56D589D-AF7F-4900-BA46-62C780AC2C10
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateRepoSourceCodeRepoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateRepoSourceCodeRepoResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateRepoSourceCodeRepoResponseBody) SetCode(v string) *UpdateRepoSourceCodeRepoResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateRepoSourceCodeRepoResponseBody) SetIsSuccess(v bool) *UpdateRepoSourceCodeRepoResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *UpdateRepoSourceCodeRepoResponseBody) SetRequestId(v string) *UpdateRepoSourceCodeRepoResponseBody {
	s.RequestId = &v
	return s
}

type UpdateRepoSourceCodeRepoResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateRepoSourceCodeRepoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateRepoSourceCodeRepoResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateRepoSourceCodeRepoResponse) GoString() string {
	return s.String()
}

func (s *UpdateRepoSourceCodeRepoResponse) SetHeaders(v map[string]*string) *UpdateRepoSourceCodeRepoResponse {
	s.Headers = v
	return s
}

func (s *UpdateRepoSourceCodeRepoResponse) SetStatusCode(v int32) *UpdateRepoSourceCodeRepoResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateRepoSourceCodeRepoResponse) SetBody(v *UpdateRepoSourceCodeRepoResponseBody) *UpdateRepoSourceCodeRepoResponse {
	s.Body = v
	return s
}

type UpdateRepoTriggerRequest struct {
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cri-kmsiwlxxdcva****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the image repository.
	//
	// This parameter is required.
	//
	// example:
	//
	// crr-tquyps22md8p****
	RepoId *string `json:"RepoId,omitempty" xml:"RepoId,omitempty"`
	// The ID of the trigger.
	//
	// This parameter is required.
	//
	// example:
	//
	// crw-k7bdx4kt52ty****
	TriggerId *string `json:"TriggerId,omitempty" xml:"TriggerId,omitempty"`
	// The name of the trigger.
	//
	// You can specify the TriggerName or TriggerUrl parameter. The TriggerName parameter is optional.
	//
	// example:
	//
	// test_trigger
	TriggerName *string `json:"TriggerName,omitempty" xml:"TriggerName,omitempty"`
	// The image tag based on which the trigger is set.
	//
	// example:
	//
	// master
	TriggerTag *string `json:"TriggerTag,omitempty" xml:"TriggerTag,omitempty"`
	// The type of the trigger. Valid values:
	//
	// 	- `ALL`: a trigger that supports both tags and regular expressions.
	//
	// 	- `TAG_LISTTAG`: a tag-based trigger.
	//
	// 	- `TAG_REG_EXP`: a regular expression-based trigger.
	//
	// example:
	//
	// TAG_LIST
	TriggerType *string `json:"TriggerType,omitempty" xml:"TriggerType,omitempty"`
	// The URL of the trigger.
	//
	// example:
	//
	// https://www.test.com
	TriggerUrl *string `json:"TriggerUrl,omitempty" xml:"TriggerUrl,omitempty"`
}

func (s UpdateRepoTriggerRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateRepoTriggerRequest) GoString() string {
	return s.String()
}

func (s *UpdateRepoTriggerRequest) SetInstanceId(v string) *UpdateRepoTriggerRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateRepoTriggerRequest) SetRepoId(v string) *UpdateRepoTriggerRequest {
	s.RepoId = &v
	return s
}

func (s *UpdateRepoTriggerRequest) SetTriggerId(v string) *UpdateRepoTriggerRequest {
	s.TriggerId = &v
	return s
}

func (s *UpdateRepoTriggerRequest) SetTriggerName(v string) *UpdateRepoTriggerRequest {
	s.TriggerName = &v
	return s
}

func (s *UpdateRepoTriggerRequest) SetTriggerTag(v string) *UpdateRepoTriggerRequest {
	s.TriggerTag = &v
	return s
}

func (s *UpdateRepoTriggerRequest) SetTriggerType(v string) *UpdateRepoTriggerRequest {
	s.TriggerType = &v
	return s
}

func (s *UpdateRepoTriggerRequest) SetTriggerUrl(v string) *UpdateRepoTriggerRequest {
	s.TriggerUrl = &v
	return s
}

type UpdateRepoTriggerResponseBody struct {
	// The return value.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 32535049-ED91-4589-98C0-7C88766EDF1F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateRepoTriggerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateRepoTriggerResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateRepoTriggerResponseBody) SetCode(v string) *UpdateRepoTriggerResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateRepoTriggerResponseBody) SetIsSuccess(v bool) *UpdateRepoTriggerResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *UpdateRepoTriggerResponseBody) SetRequestId(v string) *UpdateRepoTriggerResponseBody {
	s.RequestId = &v
	return s
}

type UpdateRepoTriggerResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateRepoTriggerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateRepoTriggerResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateRepoTriggerResponse) GoString() string {
	return s.String()
}

func (s *UpdateRepoTriggerResponse) SetHeaders(v map[string]*string) *UpdateRepoTriggerResponse {
	s.Headers = v
	return s
}

func (s *UpdateRepoTriggerResponse) SetStatusCode(v int32) *UpdateRepoTriggerResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateRepoTriggerResponse) SetBody(v *UpdateRepoTriggerResponseBody) *UpdateRepoTriggerResponse {
	s.Body = v
	return s
}

type UpdateRepositoryRequest struct {
	// example:
	//
	// repo-for-test
	Detail *string `json:"Detail,omitempty" xml:"Detail,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cri-kmsiwlxxdcva****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// crr-tquyps22md8p****
	RepoId *string `json:"RepoId,omitempty" xml:"RepoId,omitempty"`
	// example:
	//
	// dsp/domain-microapp
	RepoName *string `json:"RepoName,omitempty" xml:"RepoName,omitempty"`
	// example:
	//
	// ejiayou-other
	RepoNamespaceName *string `json:"RepoNamespaceName,omitempty" xml:"RepoNamespaceName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// PUBLIC
	RepoType *string `json:"RepoType,omitempty" xml:"RepoType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test repo
	Summary *string `json:"Summary,omitempty" xml:"Summary,omitempty"`
	// example:
	//
	// true
	TagImmutability *bool `json:"TagImmutability,omitempty" xml:"TagImmutability,omitempty"`
}

func (s UpdateRepositoryRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateRepositoryRequest) GoString() string {
	return s.String()
}

func (s *UpdateRepositoryRequest) SetDetail(v string) *UpdateRepositoryRequest {
	s.Detail = &v
	return s
}

func (s *UpdateRepositoryRequest) SetInstanceId(v string) *UpdateRepositoryRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateRepositoryRequest) SetRepoId(v string) *UpdateRepositoryRequest {
	s.RepoId = &v
	return s
}

func (s *UpdateRepositoryRequest) SetRepoName(v string) *UpdateRepositoryRequest {
	s.RepoName = &v
	return s
}

func (s *UpdateRepositoryRequest) SetRepoNamespaceName(v string) *UpdateRepositoryRequest {
	s.RepoNamespaceName = &v
	return s
}

func (s *UpdateRepositoryRequest) SetRepoType(v string) *UpdateRepositoryRequest {
	s.RepoType = &v
	return s
}

func (s *UpdateRepositoryRequest) SetSummary(v string) *UpdateRepositoryRequest {
	s.Summary = &v
	return s
}

func (s *UpdateRepositoryRequest) SetTagImmutability(v bool) *UpdateRepositoryRequest {
	s.TagImmutability = &v
	return s
}

type UpdateRepositoryResponseBody struct {
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// true
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// example:
	//
	// CC43EC6B-0DD4-40AE-8811-B0519617051A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateRepositoryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateRepositoryResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateRepositoryResponseBody) SetCode(v string) *UpdateRepositoryResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateRepositoryResponseBody) SetIsSuccess(v bool) *UpdateRepositoryResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *UpdateRepositoryResponseBody) SetRequestId(v string) *UpdateRepositoryResponseBody {
	s.RequestId = &v
	return s
}

type UpdateRepositoryResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateRepositoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateRepositoryResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateRepositoryResponse) GoString() string {
	return s.String()
}

func (s *UpdateRepositoryResponse) SetHeaders(v map[string]*string) *UpdateRepositoryResponse {
	s.Headers = v
	return s
}

func (s *UpdateRepositoryResponse) SetStatusCode(v int32) *UpdateRepositoryResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateRepositoryResponse) SetBody(v *UpdateRepositoryResponseBody) *UpdateRepositoryResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("cr"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// Cancels an artifact building task.
//
// @param request - CancelArtifactBuildTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelArtifactBuildTaskResponse
func (client *Client) CancelArtifactBuildTaskWithOptions(request *CancelArtifactBuildTaskRequest, runtime *util.RuntimeOptions) (_result *CancelArtifactBuildTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BuildTaskId)) {
		query["BuildTaskId"] = request.BuildTaskId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CancelArtifactBuildTask"),
		Version:     tea.String("2018-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CancelArtifactBuildTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Cancels an artifact building task.
//
// @param request - CancelArtifactBuildTaskRequest
//
// @return CancelArtifactBuildTaskResponse
func (client *Client) CancelArtifactBuildTask(request *CancelArtifactBuildTaskRequest) (_result *CancelArtifactBuildTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CancelArtifactBuildTaskResponse{}
	_body, _err := client.CancelArtifactBuildTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Cancels an image building task of a repository.
//
// @param request - CancelRepoBuildRecordRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelRepoBuildRecordResponse
func (client *Client) CancelRepoBuildRecordWithOptions(request *CancelRepoBuildRecordRequest, runtime *util.RuntimeOptions) (_result *CancelRepoBuildRecordResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BuildRecordId)) {
		query["BuildRecordId"] = request.BuildRecordId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RepoId)) {
		query["RepoId"] = request.RepoId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CancelRepoBuildRecord"),
		Version:     tea.String("2018-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CancelRepoBuildRecordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Cancels an image building task of a repository.
//
// @param request - CancelRepoBuildRecordRequest
//
// @return CancelRepoBuildRecordResponse
func (client *Client) CancelRepoBuildRecord(request *CancelRepoBuildRecordRequest) (_result *CancelRepoBuildRecordResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CancelRepoBuildRecordResponse{}
	_body, _err := client.CancelRepoBuildRecordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Cancels a single replication task.
//
// @param request - CancelRepoSyncTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelRepoSyncTaskResponse
func (client *Client) CancelRepoSyncTaskWithOptions(request *CancelRepoSyncTaskRequest, runtime *util.RuntimeOptions) (_result *CancelRepoSyncTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.SyncTaskId)) {
		query["SyncTaskId"] = request.SyncTaskId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CancelRepoSyncTask"),
		Version:     tea.String("2018-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CancelRepoSyncTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Cancels a single replication task.
//
// @param request - CancelRepoSyncTaskRequest
//
// @return CancelRepoSyncTaskResponse
func (client *Client) CancelRepoSyncTask(request *CancelRepoSyncTaskRequest) (_result *CancelRepoSyncTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CancelRepoSyncTaskResponse{}
	_body, _err := client.CancelRepoSyncTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Changes the resource group to which a resource belongs.
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

	if !tea.BoolValue(util.IsUnset(request.ResourceRegionId)) {
		query["ResourceRegionId"] = request.ResourceRegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ChangeResourceGroup"),
		Version:     tea.String("2018-12-01"),
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
// Changes the resource group to which a resource belongs.
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
// Create image repository acceleration rules for image building.
//
// Description:
//
// You can create building rules of accelerated images only for image repositories of Container Registry Advanced Edition instances. You cannot create building rules of accelerated images for image repositories of Container Registry Basic Edition instances. For more information, see [Specifications of different editions](https://www.alibabacloud.com/help/en/acr/product-overview/what-is-container-registry?spm=openapi-amp.newDocPublishment.0.0.bf82281fRj7rmV#section-n3q-ps7-x6k).
//
// Accelerated images are not supported in Alibaba Finance Cloud regions or Alibaba Gov Cloud regions.
//
// @param tmpReq - CreateArtifactBuildRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateArtifactBuildRuleResponse
func (client *Client) CreateArtifactBuildRuleWithOptions(tmpReq *CreateArtifactBuildRuleRequest, runtime *util.RuntimeOptions) (_result *CreateArtifactBuildRuleResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateArtifactBuildRuleShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Parameters)) {
		request.ParametersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Parameters, tea.String("Parameters"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ArtifactType)) {
		query["ArtifactType"] = request.ArtifactType
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.ParametersShrink)) {
		query["Parameters"] = request.ParametersShrink
	}

	if !tea.BoolValue(util.IsUnset(request.ScopeId)) {
		query["ScopeId"] = request.ScopeId
	}

	if !tea.BoolValue(util.IsUnset(request.ScopeType)) {
		query["ScopeType"] = request.ScopeType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateArtifactBuildRule"),
		Version:     tea.String("2018-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateArtifactBuildRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Create image repository acceleration rules for image building.
//
// Description:
//
// You can create building rules of accelerated images only for image repositories of Container Registry Advanced Edition instances. You cannot create building rules of accelerated images for image repositories of Container Registry Basic Edition instances. For more information, see [Specifications of different editions](https://www.alibabacloud.com/help/en/acr/product-overview/what-is-container-registry?spm=openapi-amp.newDocPublishment.0.0.bf82281fRj7rmV#section-n3q-ps7-x6k).
//
// Accelerated images are not supported in Alibaba Finance Cloud regions or Alibaba Gov Cloud regions.
//
// @param request - CreateArtifactBuildRuleRequest
//
// @return CreateArtifactBuildRuleResponse
func (client *Client) CreateArtifactBuildRule(request *CreateArtifactBuildRuleRequest) (_result *CreateArtifactBuildRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateArtifactBuildRuleResponse{}
	_body, _err := client.CreateArtifactBuildRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a lifecycle management rule for an artifact.
//
// @param request - CreateArtifactLifecycleRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateArtifactLifecycleRuleResponse
func (client *Client) CreateArtifactLifecycleRuleWithOptions(request *CreateArtifactLifecycleRuleRequest, runtime *util.RuntimeOptions) (_result *CreateArtifactLifecycleRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Auto)) {
		query["Auto"] = request.Auto
	}

	if !tea.BoolValue(util.IsUnset(request.EnableDeleteTag)) {
		query["EnableDeleteTag"] = request.EnableDeleteTag
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.NamespaceName)) {
		query["NamespaceName"] = request.NamespaceName
	}

	if !tea.BoolValue(util.IsUnset(request.RepoName)) {
		query["RepoName"] = request.RepoName
	}

	if !tea.BoolValue(util.IsUnset(request.RetentionTagCount)) {
		query["RetentionTagCount"] = request.RetentionTagCount
	}

	if !tea.BoolValue(util.IsUnset(request.ScheduleTime)) {
		query["ScheduleTime"] = request.ScheduleTime
	}

	if !tea.BoolValue(util.IsUnset(request.Scope)) {
		query["Scope"] = request.Scope
	}

	if !tea.BoolValue(util.IsUnset(request.TagRegexp)) {
		query["TagRegexp"] = request.TagRegexp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateArtifactLifecycleRule"),
		Version:     tea.String("2018-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateArtifactLifecycleRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a lifecycle management rule for an artifact.
//
// @param request - CreateArtifactLifecycleRuleRequest
//
// @return CreateArtifactLifecycleRuleResponse
func (client *Client) CreateArtifactLifecycleRule(request *CreateArtifactLifecycleRuleRequest) (_result *CreateArtifactLifecycleRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateArtifactLifecycleRuleResponse{}
	_body, _err := client.CreateArtifactLifecycleRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an artifact subscription rule.
//
// @param request - CreateArtifactSubscriptionRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateArtifactSubscriptionRuleResponse
func (client *Client) CreateArtifactSubscriptionRuleWithOptions(request *CreateArtifactSubscriptionRuleRequest, runtime *util.RuntimeOptions) (_result *CreateArtifactSubscriptionRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Accelerate)) {
		query["Accelerate"] = request.Accelerate
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.NamespaceName)) {
		query["NamespaceName"] = request.NamespaceName
	}

	if !tea.BoolValue(util.IsUnset(request.Override)) {
		query["Override"] = request.Override
	}

	if !tea.BoolValue(util.IsUnset(request.Platform)) {
		query["Platform"] = request.Platform
	}

	if !tea.BoolValue(util.IsUnset(request.RepoName)) {
		query["RepoName"] = request.RepoName
	}

	if !tea.BoolValue(util.IsUnset(request.SourceNamespaceName)) {
		query["SourceNamespaceName"] = request.SourceNamespaceName
	}

	if !tea.BoolValue(util.IsUnset(request.SourceProvider)) {
		query["SourceProvider"] = request.SourceProvider
	}

	if !tea.BoolValue(util.IsUnset(request.SourceRepoName)) {
		query["SourceRepoName"] = request.SourceRepoName
	}

	if !tea.BoolValue(util.IsUnset(request.TagCount)) {
		query["TagCount"] = request.TagCount
	}

	if !tea.BoolValue(util.IsUnset(request.TagRegexp)) {
		query["TagRegexp"] = request.TagRegexp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateArtifactSubscriptionRule"),
		Version:     tea.String("2018-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateArtifactSubscriptionRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an artifact subscription rule.
//
// @param request - CreateArtifactSubscriptionRuleRequest
//
// @return CreateArtifactSubscriptionRuleResponse
func (client *Client) CreateArtifactSubscriptionRule(request *CreateArtifactSubscriptionRuleRequest) (_result *CreateArtifactSubscriptionRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateArtifactSubscriptionRuleResponse{}
	_body, _err := client.CreateArtifactSubscriptionRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an artifact subscription task.
//
// @param request - CreateArtifactSubscriptionTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateArtifactSubscriptionTaskResponse
func (client *Client) CreateArtifactSubscriptionTaskWithOptions(request *CreateArtifactSubscriptionTaskRequest, runtime *util.RuntimeOptions) (_result *CreateArtifactSubscriptionTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RuleId)) {
		query["RuleId"] = request.RuleId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateArtifactSubscriptionTask"),
		Version:     tea.String("2018-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateArtifactSubscriptionTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an artifact subscription task.
//
// @param request - CreateArtifactSubscriptionTaskRequest
//
// @return CreateArtifactSubscriptionTaskResponse
func (client *Client) CreateArtifactSubscriptionTask(request *CreateArtifactSubscriptionTaskRequest) (_result *CreateArtifactSubscriptionTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateArtifactSubscriptionTaskResponse{}
	_body, _err := client.CreateArtifactSubscriptionTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an image building record based on an existing record.
//
// @param request - CreateBuildRecordByRecordRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateBuildRecordByRecordResponse
func (client *Client) CreateBuildRecordByRecordWithOptions(request *CreateBuildRecordByRecordRequest, runtime *util.RuntimeOptions) (_result *CreateBuildRecordByRecordResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BuildRecordId)) {
		query["BuildRecordId"] = request.BuildRecordId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RepoId)) {
		query["RepoId"] = request.RepoId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateBuildRecordByRecord"),
		Version:     tea.String("2018-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateBuildRecordByRecordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an image building record based on an existing record.
//
// @param request - CreateBuildRecordByRecordRequest
//
// @return CreateBuildRecordByRecordResponse
func (client *Client) CreateBuildRecordByRecord(request *CreateBuildRecordByRecordRequest) (_result *CreateBuildRecordByRecordResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateBuildRecordByRecordResponse{}
	_body, _err := client.CreateBuildRecordByRecordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an image building record based on a rule.
//
// @param request - CreateBuildRecordByRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateBuildRecordByRuleResponse
func (client *Client) CreateBuildRecordByRuleWithOptions(request *CreateBuildRecordByRuleRequest, runtime *util.RuntimeOptions) (_result *CreateBuildRecordByRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BuildRuleId)) {
		query["BuildRuleId"] = request.BuildRuleId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RepoId)) {
		query["RepoId"] = request.RepoId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateBuildRecordByRule"),
		Version:     tea.String("2018-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateBuildRecordByRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an image building record based on a rule.
//
// @param request - CreateBuildRecordByRuleRequest
//
// @return CreateBuildRecordByRuleResponse
func (client *Client) CreateBuildRecordByRule(request *CreateBuildRecordByRuleRequest) (_result *CreateBuildRecordByRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateBuildRecordByRuleResponse{}
	_body, _err := client.CreateBuildRecordByRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a delivery chain.
//
// @param request - CreateChainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateChainResponse
func (client *Client) CreateChainWithOptions(request *CreateChainRequest, runtime *util.RuntimeOptions) (_result *CreateChainResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ChainConfig)) {
		query["ChainConfig"] = request.ChainConfig
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.RepoName)) {
		query["RepoName"] = request.RepoName
	}

	if !tea.BoolValue(util.IsUnset(request.RepoNamespaceName)) {
		query["RepoNamespaceName"] = request.RepoNamespaceName
	}

	if !tea.BoolValue(util.IsUnset(request.ScopeExclude)) {
		query["ScopeExclude"] = request.ScopeExclude
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateChain"),
		Version:     tea.String("2018-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateChainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a delivery chain.
//
// @param request - CreateChainRequest
//
// @return CreateChainResponse
func (client *Client) CreateChain(request *CreateChainRequest) (_result *CreateChainResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateChainResponse{}
	_body, _err := client.CreateChainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a chart namespace in an instance.
//
// @param request - CreateChartNamespaceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateChartNamespaceResponse
func (client *Client) CreateChartNamespaceWithOptions(request *CreateChartNamespaceRequest, runtime *util.RuntimeOptions) (_result *CreateChartNamespaceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AutoCreateRepo)) {
		query["AutoCreateRepo"] = request.AutoCreateRepo
	}

	if !tea.BoolValue(util.IsUnset(request.DefaultRepoType)) {
		query["DefaultRepoType"] = request.DefaultRepoType
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.NamespaceName)) {
		query["NamespaceName"] = request.NamespaceName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateChartNamespace"),
		Version:     tea.String("2018-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateChartNamespaceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a chart namespace in an instance.
//
// @param request - CreateChartNamespaceRequest
//
// @return CreateChartNamespaceResponse
func (client *Client) CreateChartNamespace(request *CreateChartNamespaceRequest) (_result *CreateChartNamespaceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateChartNamespaceResponse{}
	_body, _err := client.CreateChartNamespaceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a chart repository.
//
// @param request - CreateChartRepositoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateChartRepositoryResponse
func (client *Client) CreateChartRepositoryWithOptions(request *CreateChartRepositoryRequest, runtime *util.RuntimeOptions) (_result *CreateChartRepositoryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RepoName)) {
		query["RepoName"] = request.RepoName
	}

	if !tea.BoolValue(util.IsUnset(request.RepoNamespaceName)) {
		query["RepoNamespaceName"] = request.RepoNamespaceName
	}

	if !tea.BoolValue(util.IsUnset(request.RepoType)) {
		query["RepoType"] = request.RepoType
	}

	if !tea.BoolValue(util.IsUnset(request.Summary)) {
		query["Summary"] = request.Summary
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateChartRepository"),
		Version:     tea.String("2018-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateChartRepositoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a chart repository.
//
// @param request - CreateChartRepositoryRequest
//
// @return CreateChartRepositoryResponse
func (client *Client) CreateChartRepository(request *CreateChartRepositoryRequest) (_result *CreateChartRepositoryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateChartRepositoryResponse{}
	_body, _err := client.CreateChartRepositoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a whitelist policy for the public endpoint of the instance.
//
// @param request - CreateInstanceEndpointAclPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateInstanceEndpointAclPolicyResponse
func (client *Client) CreateInstanceEndpointAclPolicyWithOptions(request *CreateInstanceEndpointAclPolicyRequest, runtime *util.RuntimeOptions) (_result *CreateInstanceEndpointAclPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Comment)) {
		query["Comment"] = request.Comment
	}

	if !tea.BoolValue(util.IsUnset(request.EndpointType)) {
		query["EndpointType"] = request.EndpointType
	}

	if !tea.BoolValue(util.IsUnset(request.Entry)) {
		query["Entry"] = request.Entry
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.ModuleName)) {
		query["ModuleName"] = request.ModuleName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateInstanceEndpointAclPolicy"),
		Version:     tea.String("2018-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateInstanceEndpointAclPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a whitelist policy for the public endpoint of the instance.
//
// @param request - CreateInstanceEndpointAclPolicyRequest
//
// @return CreateInstanceEndpointAclPolicyResponse
func (client *Client) CreateInstanceEndpointAclPolicy(request *CreateInstanceEndpointAclPolicyRequest) (_result *CreateInstanceEndpointAclPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateInstanceEndpointAclPolicyResponse{}
	_body, _err := client.CreateInstanceEndpointAclPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Associates a virtual private cloud (VPC) with a Container Registry instance.
//
// Description:
//
// The VPC quota must be purchased separately.
//
// @param request - CreateInstanceVpcEndpointLinkedVpcRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateInstanceVpcEndpointLinkedVpcResponse
func (client *Client) CreateInstanceVpcEndpointLinkedVpcWithOptions(request *CreateInstanceVpcEndpointLinkedVpcRequest, runtime *util.RuntimeOptions) (_result *CreateInstanceVpcEndpointLinkedVpcResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EnableCreateDNSRecordInPvzt)) {
		query["EnableCreateDNSRecordInPvzt"] = request.EnableCreateDNSRecordInPvzt
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.ModuleName)) {
		query["ModuleName"] = request.ModuleName
	}

	if !tea.BoolValue(util.IsUnset(request.VpcId)) {
		query["VpcId"] = request.VpcId
	}

	if !tea.BoolValue(util.IsUnset(request.VswitchId)) {
		query["VswitchId"] = request.VswitchId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateInstanceVpcEndpointLinkedVpc"),
		Version:     tea.String("2018-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateInstanceVpcEndpointLinkedVpcResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Associates a virtual private cloud (VPC) with a Container Registry instance.
//
// Description:
//
// The VPC quota must be purchased separately.
//
// @param request - CreateInstanceVpcEndpointLinkedVpcRequest
//
// @return CreateInstanceVpcEndpointLinkedVpcResponse
func (client *Client) CreateInstanceVpcEndpointLinkedVpc(request *CreateInstanceVpcEndpointLinkedVpcRequest) (_result *CreateInstanceVpcEndpointLinkedVpcResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateInstanceVpcEndpointLinkedVpcResponse{}
	_body, _err := client.CreateInstanceVpcEndpointLinkedVpcWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a namespace of image repositories.
//
// @param tmpReq - CreateNamespaceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateNamespaceResponse
func (client *Client) CreateNamespaceWithOptions(tmpReq *CreateNamespaceRequest, runtime *util.RuntimeOptions) (_result *CreateNamespaceResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateNamespaceShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.DefaultRepoConfiguration)) {
		request.DefaultRepoConfigurationShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DefaultRepoConfiguration, tea.String("DefaultRepoConfiguration"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AutoCreateRepo)) {
		query["AutoCreateRepo"] = request.AutoCreateRepo
	}

	if !tea.BoolValue(util.IsUnset(request.DefaultRepoConfigurationShrink)) {
		query["DefaultRepoConfiguration"] = request.DefaultRepoConfigurationShrink
	}

	if !tea.BoolValue(util.IsUnset(request.DefaultRepoType)) {
		query["DefaultRepoType"] = request.DefaultRepoType
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.NamespaceName)) {
		query["NamespaceName"] = request.NamespaceName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateNamespace"),
		Version:     tea.String("2018-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateNamespaceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a namespace of image repositories.
//
// @param request - CreateNamespaceRequest
//
// @return CreateNamespaceResponse
func (client *Client) CreateNamespace(request *CreateNamespaceRequest) (_result *CreateNamespaceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateNamespaceResponse{}
	_body, _err := client.CreateNamespaceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an image building rule for a repository.
//
// @param request - CreateRepoBuildRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateRepoBuildRuleResponse
func (client *Client) CreateRepoBuildRuleWithOptions(request *CreateRepoBuildRuleRequest, runtime *util.RuntimeOptions) (_result *CreateRepoBuildRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BuildArgs)) {
		query["BuildArgs"] = request.BuildArgs
	}

	if !tea.BoolValue(util.IsUnset(request.DockerfileLocation)) {
		query["DockerfileLocation"] = request.DockerfileLocation
	}

	if !tea.BoolValue(util.IsUnset(request.DockerfileName)) {
		query["DockerfileName"] = request.DockerfileName
	}

	if !tea.BoolValue(util.IsUnset(request.ImageTag)) {
		query["ImageTag"] = request.ImageTag
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Platforms)) {
		query["Platforms"] = request.Platforms
	}

	if !tea.BoolValue(util.IsUnset(request.PushName)) {
		query["PushName"] = request.PushName
	}

	if !tea.BoolValue(util.IsUnset(request.PushType)) {
		query["PushType"] = request.PushType
	}

	if !tea.BoolValue(util.IsUnset(request.RepoId)) {
		query["RepoId"] = request.RepoId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateRepoBuildRule"),
		Version:     tea.String("2018-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateRepoBuildRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an image building rule for a repository.
//
// @param request - CreateRepoBuildRuleRequest
//
// @return CreateRepoBuildRuleResponse
func (client *Client) CreateRepoBuildRule(request *CreateRepoBuildRuleRequest) (_result *CreateRepoBuildRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateRepoBuildRuleResponse{}
	_body, _err := client.CreateRepoBuildRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Binds a source code repository to an image repository.
//
// @param request - CreateRepoSourceCodeRepoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateRepoSourceCodeRepoResponse
func (client *Client) CreateRepoSourceCodeRepoWithOptions(request *CreateRepoSourceCodeRepoRequest, runtime *util.RuntimeOptions) (_result *CreateRepoSourceCodeRepoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AutoBuild)) {
		query["AutoBuild"] = request.AutoBuild
	}

	if !tea.BoolValue(util.IsUnset(request.CodeRepoName)) {
		query["CodeRepoName"] = request.CodeRepoName
	}

	if !tea.BoolValue(util.IsUnset(request.CodeRepoNamespaceName)) {
		query["CodeRepoNamespaceName"] = request.CodeRepoNamespaceName
	}

	if !tea.BoolValue(util.IsUnset(request.CodeRepoType)) {
		query["CodeRepoType"] = request.CodeRepoType
	}

	if !tea.BoolValue(util.IsUnset(request.DisableCacheBuild)) {
		query["DisableCacheBuild"] = request.DisableCacheBuild
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.OverseaBuild)) {
		query["OverseaBuild"] = request.OverseaBuild
	}

	if !tea.BoolValue(util.IsUnset(request.RepoId)) {
		query["RepoId"] = request.RepoId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateRepoSourceCodeRepo"),
		Version:     tea.String("2018-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateRepoSourceCodeRepoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Binds a source code repository to an image repository.
//
// @param request - CreateRepoSourceCodeRepoRequest
//
// @return CreateRepoSourceCodeRepoResponse
func (client *Client) CreateRepoSourceCodeRepo(request *CreateRepoSourceCodeRepoRequest) (_result *CreateRepoSourceCodeRepoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateRepoSourceCodeRepoResponse{}
	_body, _err := client.CreateRepoSourceCodeRepoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an image synchronization rule for an image repository.
//
// @param request - CreateRepoSyncRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateRepoSyncRuleResponse
func (client *Client) CreateRepoSyncRuleWithOptions(request *CreateRepoSyncRuleRequest, runtime *util.RuntimeOptions) (_result *CreateRepoSyncRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.NamespaceName)) {
		query["NamespaceName"] = request.NamespaceName
	}

	if !tea.BoolValue(util.IsUnset(request.RepoName)) {
		query["RepoName"] = request.RepoName
	}

	if !tea.BoolValue(util.IsUnset(request.RepoNameFilter)) {
		query["RepoNameFilter"] = request.RepoNameFilter
	}

	if !tea.BoolValue(util.IsUnset(request.SyncRuleName)) {
		query["SyncRuleName"] = request.SyncRuleName
	}

	if !tea.BoolValue(util.IsUnset(request.SyncScope)) {
		query["SyncScope"] = request.SyncScope
	}

	if !tea.BoolValue(util.IsUnset(request.SyncTrigger)) {
		query["SyncTrigger"] = request.SyncTrigger
	}

	if !tea.BoolValue(util.IsUnset(request.TagFilter)) {
		query["TagFilter"] = request.TagFilter
	}

	if !tea.BoolValue(util.IsUnset(request.TargetInstanceId)) {
		query["TargetInstanceId"] = request.TargetInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.TargetNamespaceName)) {
		query["TargetNamespaceName"] = request.TargetNamespaceName
	}

	if !tea.BoolValue(util.IsUnset(request.TargetRegionId)) {
		query["TargetRegionId"] = request.TargetRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.TargetRepoName)) {
		query["TargetRepoName"] = request.TargetRepoName
	}

	if !tea.BoolValue(util.IsUnset(request.TargetUserId)) {
		query["TargetUserId"] = request.TargetUserId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateRepoSyncRule"),
		Version:     tea.String("2018-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateRepoSyncRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an image synchronization rule for an image repository.
//
// @param request - CreateRepoSyncRuleRequest
//
// @return CreateRepoSyncRuleResponse
func (client *Client) CreateRepoSyncRule(request *CreateRepoSyncRuleRequest) (_result *CreateRepoSyncRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateRepoSyncRuleResponse{}
	_body, _err := client.CreateRepoSyncRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - CreateRepoSyncTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateRepoSyncTaskResponse
func (client *Client) CreateRepoSyncTaskWithOptions(request *CreateRepoSyncTaskRequest, runtime *util.RuntimeOptions) (_result *CreateRepoSyncTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Override)) {
		query["Override"] = request.Override
	}

	if !tea.BoolValue(util.IsUnset(request.RepoId)) {
		query["RepoId"] = request.RepoId
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	if !tea.BoolValue(util.IsUnset(request.TargetInstanceId)) {
		query["TargetInstanceId"] = request.TargetInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.TargetNamespace)) {
		query["TargetNamespace"] = request.TargetNamespace
	}

	if !tea.BoolValue(util.IsUnset(request.TargetRegionId)) {
		query["TargetRegionId"] = request.TargetRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.TargetRepoName)) {
		query["TargetRepoName"] = request.TargetRepoName
	}

	if !tea.BoolValue(util.IsUnset(request.TargetTag)) {
		query["TargetTag"] = request.TargetTag
	}

	if !tea.BoolValue(util.IsUnset(request.TargetUserId)) {
		query["TargetUserId"] = request.TargetUserId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateRepoSyncTask"),
		Version:     tea.String("2018-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateRepoSyncTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - CreateRepoSyncTaskRequest
//
// @return CreateRepoSyncTaskResponse
func (client *Client) CreateRepoSyncTask(request *CreateRepoSyncTaskRequest) (_result *CreateRepoSyncTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateRepoSyncTaskResponse{}
	_body, _err := client.CreateRepoSyncTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an image replication task based on a manual replication rule.
//
// @param request - CreateRepoSyncTaskByRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateRepoSyncTaskByRuleResponse
func (client *Client) CreateRepoSyncTaskByRuleWithOptions(request *CreateRepoSyncTaskByRuleRequest, runtime *util.RuntimeOptions) (_result *CreateRepoSyncTaskByRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RepoId)) {
		query["RepoId"] = request.RepoId
	}

	if !tea.BoolValue(util.IsUnset(request.SyncRuleId)) {
		query["SyncRuleId"] = request.SyncRuleId
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateRepoSyncTaskByRule"),
		Version:     tea.String("2018-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateRepoSyncTaskByRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an image replication task based on a manual replication rule.
//
// @param request - CreateRepoSyncTaskByRuleRequest
//
// @return CreateRepoSyncTaskByRuleResponse
func (client *Client) CreateRepoSyncTaskByRule(request *CreateRepoSyncTaskByRuleRequest) (_result *CreateRepoSyncTaskByRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateRepoSyncTaskByRuleResponse{}
	_body, _err := client.CreateRepoSyncTaskByRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an image tag based on an existing image tag in an image repository.
//
// @param request - CreateRepoTagRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateRepoTagResponse
func (client *Client) CreateRepoTagWithOptions(request *CreateRepoTagRequest, runtime *util.RuntimeOptions) (_result *CreateRepoTagResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FromTag)) {
		query["FromTag"] = request.FromTag
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.NamespaceName)) {
		query["NamespaceName"] = request.NamespaceName
	}

	if !tea.BoolValue(util.IsUnset(request.RepoName)) {
		query["RepoName"] = request.RepoName
	}

	if !tea.BoolValue(util.IsUnset(request.ToTag)) {
		query["ToTag"] = request.ToTag
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateRepoTag"),
		Version:     tea.String("2018-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateRepoTagResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an image tag based on an existing image tag in an image repository.
//
// @param request - CreateRepoTagRequest
//
// @return CreateRepoTagResponse
func (client *Client) CreateRepoTag(request *CreateRepoTagRequest) (_result *CreateRepoTagResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateRepoTagResponse{}
	_body, _err := client.CreateRepoTagWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an image scan task.
//
// @param request - CreateRepoTagScanTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateRepoTagScanTaskResponse
func (client *Client) CreateRepoTagScanTaskWithOptions(request *CreateRepoTagScanTaskRequest, runtime *util.RuntimeOptions) (_result *CreateRepoTagScanTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Digest)) {
		query["Digest"] = request.Digest
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RepoId)) {
		query["RepoId"] = request.RepoId
	}

	if !tea.BoolValue(util.IsUnset(request.ScanService)) {
		query["ScanService"] = request.ScanService
	}

	if !tea.BoolValue(util.IsUnset(request.ScanType)) {
		query["ScanType"] = request.ScanType
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateRepoTagScanTask"),
		Version:     tea.String("2018-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateRepoTagScanTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an image scan task.
//
// @param request - CreateRepoTagScanTaskRequest
//
// @return CreateRepoTagScanTaskResponse
func (client *Client) CreateRepoTagScanTask(request *CreateRepoTagScanTaskRequest) (_result *CreateRepoTagScanTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateRepoTagScanTaskResponse{}
	_body, _err := client.CreateRepoTagScanTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a trigger for a repository.
//
// @param request - CreateRepoTriggerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateRepoTriggerResponse
func (client *Client) CreateRepoTriggerWithOptions(request *CreateRepoTriggerRequest, runtime *util.RuntimeOptions) (_result *CreateRepoTriggerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RepoId)) {
		query["RepoId"] = request.RepoId
	}

	if !tea.BoolValue(util.IsUnset(request.TriggerName)) {
		query["TriggerName"] = request.TriggerName
	}

	if !tea.BoolValue(util.IsUnset(request.TriggerTag)) {
		query["TriggerTag"] = request.TriggerTag
	}

	if !tea.BoolValue(util.IsUnset(request.TriggerType)) {
		query["TriggerType"] = request.TriggerType
	}

	if !tea.BoolValue(util.IsUnset(request.TriggerUrl)) {
		query["TriggerUrl"] = request.TriggerUrl
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateRepoTrigger"),
		Version:     tea.String("2018-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateRepoTriggerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a trigger for a repository.
//
// @param request - CreateRepoTriggerRequest
//
// @return CreateRepoTriggerResponse
func (client *Client) CreateRepoTrigger(request *CreateRepoTriggerRequest) (_result *CreateRepoTriggerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateRepoTriggerResponse{}
	_body, _err := client.CreateRepoTriggerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an image repository.
//
// @param request - CreateRepositoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateRepositoryResponse
func (client *Client) CreateRepositoryWithOptions(request *CreateRepositoryRequest, runtime *util.RuntimeOptions) (_result *CreateRepositoryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Detail)) {
		query["Detail"] = request.Detail
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RepoName)) {
		query["RepoName"] = request.RepoName
	}

	if !tea.BoolValue(util.IsUnset(request.RepoNamespaceName)) {
		query["RepoNamespaceName"] = request.RepoNamespaceName
	}

	if !tea.BoolValue(util.IsUnset(request.RepoType)) {
		query["RepoType"] = request.RepoType
	}

	if !tea.BoolValue(util.IsUnset(request.Summary)) {
		query["Summary"] = request.Summary
	}

	if !tea.BoolValue(util.IsUnset(request.TagImmutability)) {
		query["TagImmutability"] = request.TagImmutability
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateRepository"),
		Version:     tea.String("2018-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateRepositoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an image repository.
//
// @param request - CreateRepositoryRequest
//
// @return CreateRepositoryResponse
func (client *Client) CreateRepository(request *CreateRepositoryRequest) (_result *CreateRepositoryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateRepositoryResponse{}
	_body, _err := client.CreateRepositoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes an artifact lifecycle management rule.
//
// @param request - DeleteArtifactLifecycleRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteArtifactLifecycleRuleResponse
func (client *Client) DeleteArtifactLifecycleRuleWithOptions(request *DeleteArtifactLifecycleRuleRequest, runtime *util.RuntimeOptions) (_result *DeleteArtifactLifecycleRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RuleId)) {
		query["RuleId"] = request.RuleId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteArtifactLifecycleRule"),
		Version:     tea.String("2018-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteArtifactLifecycleRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an artifact lifecycle management rule.
//
// @param request - DeleteArtifactLifecycleRuleRequest
//
// @return DeleteArtifactLifecycleRuleResponse
func (client *Client) DeleteArtifactLifecycleRule(request *DeleteArtifactLifecycleRuleRequest) (_result *DeleteArtifactLifecycleRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteArtifactLifecycleRuleResponse{}
	_body, _err := client.DeleteArtifactLifecycleRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes an artifact subscription rule.
//
// @param request - DeleteArtifactSubscriptionRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteArtifactSubscriptionRuleResponse
func (client *Client) DeleteArtifactSubscriptionRuleWithOptions(request *DeleteArtifactSubscriptionRuleRequest, runtime *util.RuntimeOptions) (_result *DeleteArtifactSubscriptionRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RuleId)) {
		query["RuleId"] = request.RuleId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteArtifactSubscriptionRule"),
		Version:     tea.String("2018-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteArtifactSubscriptionRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an artifact subscription rule.
//
// @param request - DeleteArtifactSubscriptionRuleRequest
//
// @return DeleteArtifactSubscriptionRuleResponse
func (client *Client) DeleteArtifactSubscriptionRule(request *DeleteArtifactSubscriptionRuleRequest) (_result *DeleteArtifactSubscriptionRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteArtifactSubscriptionRuleResponse{}
	_body, _err := client.DeleteArtifactSubscriptionRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a delivery pipeline.
//
// @param request - DeleteChainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteChainResponse
func (client *Client) DeleteChainWithOptions(request *DeleteChainRequest, runtime *util.RuntimeOptions) (_result *DeleteChainResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ChainId)) {
		query["ChainId"] = request.ChainId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteChain"),
		Version:     tea.String("2018-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteChainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a delivery pipeline.
//
// @param request - DeleteChainRequest
//
// @return DeleteChainResponse
func (client *Client) DeleteChain(request *DeleteChainRequest) (_result *DeleteChainResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteChainResponse{}
	_body, _err := client.DeleteChainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a chart namespace from an instance.
//
// Description:
//
// >  If you delete a chart namespace, all repositories in the namespace and the charts in all repositories are deleted.
//
// @param request - DeleteChartNamespaceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteChartNamespaceResponse
func (client *Client) DeleteChartNamespaceWithOptions(request *DeleteChartNamespaceRequest, runtime *util.RuntimeOptions) (_result *DeleteChartNamespaceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.NamespaceName)) {
		query["NamespaceName"] = request.NamespaceName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteChartNamespace"),
		Version:     tea.String("2018-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteChartNamespaceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a chart namespace from an instance.
//
// Description:
//
// >  If you delete a chart namespace, all repositories in the namespace and the charts in all repositories are deleted.
//
// @param request - DeleteChartNamespaceRequest
//
// @return DeleteChartNamespaceResponse
func (client *Client) DeleteChartNamespace(request *DeleteChartNamespaceRequest) (_result *DeleteChartNamespaceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteChartNamespaceResponse{}
	_body, _err := client.DeleteChartNamespaceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a chart version from a chart repository.
//
// @param request - DeleteChartReleaseRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteChartReleaseResponse
func (client *Client) DeleteChartReleaseWithOptions(request *DeleteChartReleaseRequest, runtime *util.RuntimeOptions) (_result *DeleteChartReleaseResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Chart)) {
		query["Chart"] = request.Chart
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Release)) {
		query["Release"] = request.Release
	}

	if !tea.BoolValue(util.IsUnset(request.RepoName)) {
		query["RepoName"] = request.RepoName
	}

	if !tea.BoolValue(util.IsUnset(request.RepoNamespaceName)) {
		query["RepoNamespaceName"] = request.RepoNamespaceName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteChartRelease"),
		Version:     tea.String("2018-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteChartReleaseResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a chart version from a chart repository.
//
// @param request - DeleteChartReleaseRequest
//
// @return DeleteChartReleaseResponse
func (client *Client) DeleteChartRelease(request *DeleteChartReleaseRequest) (_result *DeleteChartReleaseResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteChartReleaseResponse{}
	_body, _err := client.DeleteChartReleaseWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a chart repository from an instance.
//
// @param request - DeleteChartRepositoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteChartRepositoryResponse
func (client *Client) DeleteChartRepositoryWithOptions(request *DeleteChartRepositoryRequest, runtime *util.RuntimeOptions) (_result *DeleteChartRepositoryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RepoName)) {
		query["RepoName"] = request.RepoName
	}

	if !tea.BoolValue(util.IsUnset(request.RepoNamespaceName)) {
		query["RepoNamespaceName"] = request.RepoNamespaceName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteChartRepository"),
		Version:     tea.String("2018-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteChartRepositoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a chart repository from an instance.
//
// @param request - DeleteChartRepositoryRequest
//
// @return DeleteChartRepositoryResponse
func (client *Client) DeleteChartRepository(request *DeleteChartRepositoryRequest) (_result *DeleteChartRepositoryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteChartRepositoryResponse{}
	_body, _err := client.DeleteChartRepositoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes an event notification rule.
//
// @param request - DeleteEventCenterRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteEventCenterRuleResponse
func (client *Client) DeleteEventCenterRuleWithOptions(request *DeleteEventCenterRuleRequest, runtime *util.RuntimeOptions) (_result *DeleteEventCenterRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RuleId)) {
		query["RuleId"] = request.RuleId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteEventCenterRule"),
		Version:     tea.String("2018-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteEventCenterRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an event notification rule.
//
// @param request - DeleteEventCenterRuleRequest
//
// @return DeleteEventCenterRuleResponse
func (client *Client) DeleteEventCenterRule(request *DeleteEventCenterRuleRequest) (_result *DeleteEventCenterRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteEventCenterRuleResponse{}
	_body, _err := client.DeleteEventCenterRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a whitelist policy for the public endpoint of an instance.
//
// @param request - DeleteInstanceEndpointAclPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteInstanceEndpointAclPolicyResponse
func (client *Client) DeleteInstanceEndpointAclPolicyWithOptions(request *DeleteInstanceEndpointAclPolicyRequest, runtime *util.RuntimeOptions) (_result *DeleteInstanceEndpointAclPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndpointType)) {
		query["EndpointType"] = request.EndpointType
	}

	if !tea.BoolValue(util.IsUnset(request.Entry)) {
		query["Entry"] = request.Entry
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.ModuleName)) {
		query["ModuleName"] = request.ModuleName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteInstanceEndpointAclPolicy"),
		Version:     tea.String("2018-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteInstanceEndpointAclPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a whitelist policy for the public endpoint of an instance.
//
// @param request - DeleteInstanceEndpointAclPolicyRequest
//
// @return DeleteInstanceEndpointAclPolicyResponse
func (client *Client) DeleteInstanceEndpointAclPolicy(request *DeleteInstanceEndpointAclPolicyRequest) (_result *DeleteInstanceEndpointAclPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteInstanceEndpointAclPolicyResponse{}
	_body, _err := client.DeleteInstanceEndpointAclPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Disassociates a virtual private cloud (VPC) from an instance.
//
// @param request - DeleteInstanceVpcEndpointLinkedVpcRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteInstanceVpcEndpointLinkedVpcResponse
func (client *Client) DeleteInstanceVpcEndpointLinkedVpcWithOptions(request *DeleteInstanceVpcEndpointLinkedVpcRequest, runtime *util.RuntimeOptions) (_result *DeleteInstanceVpcEndpointLinkedVpcResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.ModuleName)) {
		query["ModuleName"] = request.ModuleName
	}

	if !tea.BoolValue(util.IsUnset(request.VpcId)) {
		query["VpcId"] = request.VpcId
	}

	if !tea.BoolValue(util.IsUnset(request.VswitchId)) {
		query["VswitchId"] = request.VswitchId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteInstanceVpcEndpointLinkedVpc"),
		Version:     tea.String("2018-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteInstanceVpcEndpointLinkedVpcResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Disassociates a virtual private cloud (VPC) from an instance.
//
// @param request - DeleteInstanceVpcEndpointLinkedVpcRequest
//
// @return DeleteInstanceVpcEndpointLinkedVpcResponse
func (client *Client) DeleteInstanceVpcEndpointLinkedVpc(request *DeleteInstanceVpcEndpointLinkedVpcRequest) (_result *DeleteInstanceVpcEndpointLinkedVpcResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteInstanceVpcEndpointLinkedVpcResponse{}
	_body, _err := client.DeleteInstanceVpcEndpointLinkedVpcWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a namespace.
//
// Description:
//
// > After you delete a namespace, all repositories in the namespace and all images in these repositories are deleted as well.
//
// @param request - DeleteNamespaceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteNamespaceResponse
func (client *Client) DeleteNamespaceWithOptions(request *DeleteNamespaceRequest, runtime *util.RuntimeOptions) (_result *DeleteNamespaceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.NamespaceName)) {
		query["NamespaceName"] = request.NamespaceName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteNamespace"),
		Version:     tea.String("2018-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteNamespaceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a namespace.
//
// Description:
//
// > After you delete a namespace, all repositories in the namespace and all images in these repositories are deleted as well.
//
// @param request - DeleteNamespaceRequest
//
// @return DeleteNamespaceResponse
func (client *Client) DeleteNamespace(request *DeleteNamespaceRequest) (_result *DeleteNamespaceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteNamespaceResponse{}
	_body, _err := client.DeleteNamespaceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes an image building rule of a repository.
//
// @param request - DeleteRepoBuildRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteRepoBuildRuleResponse
func (client *Client) DeleteRepoBuildRuleWithOptions(request *DeleteRepoBuildRuleRequest, runtime *util.RuntimeOptions) (_result *DeleteRepoBuildRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BuildRuleId)) {
		query["BuildRuleId"] = request.BuildRuleId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RepoId)) {
		query["RepoId"] = request.RepoId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteRepoBuildRule"),
		Version:     tea.String("2018-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteRepoBuildRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an image building rule of a repository.
//
// @param request - DeleteRepoBuildRuleRequest
//
// @return DeleteRepoBuildRuleResponse
func (client *Client) DeleteRepoBuildRule(request *DeleteRepoBuildRuleRequest) (_result *DeleteRepoBuildRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteRepoBuildRuleResponse{}
	_body, _err := client.DeleteRepoBuildRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes an image replication rule of an image repository.
//
// @param request - DeleteRepoSyncRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteRepoSyncRuleResponse
func (client *Client) DeleteRepoSyncRuleWithOptions(request *DeleteRepoSyncRuleRequest, runtime *util.RuntimeOptions) (_result *DeleteRepoSyncRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.SyncRuleId)) {
		query["SyncRuleId"] = request.SyncRuleId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteRepoSyncRule"),
		Version:     tea.String("2018-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteRepoSyncRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an image replication rule of an image repository.
//
// @param request - DeleteRepoSyncRuleRequest
//
// @return DeleteRepoSyncRuleResponse
func (client *Client) DeleteRepoSyncRule(request *DeleteRepoSyncRuleRequest) (_result *DeleteRepoSyncRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteRepoSyncRuleResponse{}
	_body, _err := client.DeleteRepoSyncRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes an image tag.
//
// @param request - DeleteRepoTagRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteRepoTagResponse
func (client *Client) DeleteRepoTagWithOptions(request *DeleteRepoTagRequest, runtime *util.RuntimeOptions) (_result *DeleteRepoTagResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RepoId)) {
		query["RepoId"] = request.RepoId
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteRepoTag"),
		Version:     tea.String("2018-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteRepoTagResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an image tag.
//
// @param request - DeleteRepoTagRequest
//
// @return DeleteRepoTagResponse
func (client *Client) DeleteRepoTag(request *DeleteRepoTagRequest) (_result *DeleteRepoTagResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteRepoTagResponse{}
	_body, _err := client.DeleteRepoTagWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a trigger of an image repository.
//
// @param request - DeleteRepoTriggerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteRepoTriggerResponse
func (client *Client) DeleteRepoTriggerWithOptions(request *DeleteRepoTriggerRequest, runtime *util.RuntimeOptions) (_result *DeleteRepoTriggerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RepoId)) {
		query["RepoId"] = request.RepoId
	}

	if !tea.BoolValue(util.IsUnset(request.TriggerId)) {
		query["TriggerId"] = request.TriggerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteRepoTrigger"),
		Version:     tea.String("2018-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteRepoTriggerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a trigger of an image repository.
//
// @param request - DeleteRepoTriggerRequest
//
// @return DeleteRepoTriggerResponse
func (client *Client) DeleteRepoTrigger(request *DeleteRepoTriggerRequest) (_result *DeleteRepoTriggerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteRepoTriggerResponse{}
	_body, _err := client.DeleteRepoTriggerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes an image repository.
//
// Description:
//
// If you delete a repository, all images in the repository are also deleted.
//
// @param request - DeleteRepositoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteRepositoryResponse
func (client *Client) DeleteRepositoryWithOptions(request *DeleteRepositoryRequest, runtime *util.RuntimeOptions) (_result *DeleteRepositoryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RepoId)) {
		query["RepoId"] = request.RepoId
	}

	if !tea.BoolValue(util.IsUnset(request.RepoName)) {
		query["RepoName"] = request.RepoName
	}

	if !tea.BoolValue(util.IsUnset(request.RepoNamespaceName)) {
		query["RepoNamespaceName"] = request.RepoNamespaceName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteRepository"),
		Version:     tea.String("2018-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteRepositoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an image repository.
//
// Description:
//
// If you delete a repository, all images in the repository are also deleted.
//
// @param request - DeleteRepositoryRequest
//
// @return DeleteRepositoryResponse
func (client *Client) DeleteRepository(request *DeleteRepositoryRequest) (_result *DeleteRepositoryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteRepositoryResponse{}
	_body, _err := client.DeleteRepositoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - GetArtifactBuildRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetArtifactBuildRuleResponse
func (client *Client) GetArtifactBuildRuleWithOptions(request *GetArtifactBuildRuleRequest, runtime *util.RuntimeOptions) (_result *GetArtifactBuildRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetArtifactBuildRule"),
		Version:     tea.String("2018-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetArtifactBuildRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - GetArtifactBuildRuleRequest
//
// @return GetArtifactBuildRuleResponse
func (client *Client) GetArtifactBuildRule(request *GetArtifactBuildRuleRequest) (_result *GetArtifactBuildRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetArtifactBuildRuleResponse{}
	_body, _err := client.GetArtifactBuildRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of an artifact building task.
//
// @param request - GetArtifactBuildTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetArtifactBuildTaskResponse
func (client *Client) GetArtifactBuildTaskWithOptions(request *GetArtifactBuildTaskRequest, runtime *util.RuntimeOptions) (_result *GetArtifactBuildTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetArtifactBuildTask"),
		Version:     tea.String("2018-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetArtifactBuildTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of an artifact building task.
//
// @param request - GetArtifactBuildTaskRequest
//
// @return GetArtifactBuildTaskResponse
func (client *Client) GetArtifactBuildTask(request *GetArtifactBuildTaskRequest) (_result *GetArtifactBuildTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetArtifactBuildTaskResponse{}
	_body, _err := client.GetArtifactBuildTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the lifecycle management rules of an artifact.
//
// @param request - GetArtifactLifecycleRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetArtifactLifecycleRuleResponse
func (client *Client) GetArtifactLifecycleRuleWithOptions(request *GetArtifactLifecycleRuleRequest, runtime *util.RuntimeOptions) (_result *GetArtifactLifecycleRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetArtifactLifecycleRule"),
		Version:     tea.String("2018-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetArtifactLifecycleRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the lifecycle management rules of an artifact.
//
// @param request - GetArtifactLifecycleRuleRequest
//
// @return GetArtifactLifecycleRuleResponse
func (client *Client) GetArtifactLifecycleRule(request *GetArtifactLifecycleRuleRequest) (_result *GetArtifactLifecycleRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetArtifactLifecycleRuleResponse{}
	_body, _err := client.GetArtifactLifecycleRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about an artifact subscription rule.
//
// @param request - GetArtifactSubscriptionRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetArtifactSubscriptionRuleResponse
func (client *Client) GetArtifactSubscriptionRuleWithOptions(request *GetArtifactSubscriptionRuleRequest, runtime *util.RuntimeOptions) (_result *GetArtifactSubscriptionRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetArtifactSubscriptionRule"),
		Version:     tea.String("2018-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetArtifactSubscriptionRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about an artifact subscription rule.
//
// @param request - GetArtifactSubscriptionRuleRequest
//
// @return GetArtifactSubscriptionRuleResponse
func (client *Client) GetArtifactSubscriptionRule(request *GetArtifactSubscriptionRuleRequest) (_result *GetArtifactSubscriptionRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetArtifactSubscriptionRuleResponse{}
	_body, _err := client.GetArtifactSubscriptionRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries an artifact subscription task.
//
// @param request - GetArtifactSubscriptionTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetArtifactSubscriptionTaskResponse
func (client *Client) GetArtifactSubscriptionTaskWithOptions(request *GetArtifactSubscriptionTaskRequest, runtime *util.RuntimeOptions) (_result *GetArtifactSubscriptionTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetArtifactSubscriptionTask"),
		Version:     tea.String("2018-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetArtifactSubscriptionTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries an artifact subscription task.
//
// @param request - GetArtifactSubscriptionTaskRequest
//
// @return GetArtifactSubscriptionTaskResponse
func (client *Client) GetArtifactSubscriptionTask(request *GetArtifactSubscriptionTaskRequest) (_result *GetArtifactSubscriptionTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetArtifactSubscriptionTaskResponse{}
	_body, _err := client.GetArtifactSubscriptionTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of an artifact subscription task.
//
// @param request - GetArtifactSubscriptionTaskResultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetArtifactSubscriptionTaskResultResponse
func (client *Client) GetArtifactSubscriptionTaskResultWithOptions(request *GetArtifactSubscriptionTaskResultRequest, runtime *util.RuntimeOptions) (_result *GetArtifactSubscriptionTaskResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetArtifactSubscriptionTaskResult"),
		Version:     tea.String("2018-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetArtifactSubscriptionTaskResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of an artifact subscription task.
//
// @param request - GetArtifactSubscriptionTaskResultRequest
//
// @return GetArtifactSubscriptionTaskResultResponse
func (client *Client) GetArtifactSubscriptionTaskResult(request *GetArtifactSubscriptionTaskResultRequest) (_result *GetArtifactSubscriptionTaskResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetArtifactSubscriptionTaskResultResponse{}
	_body, _err := client.GetArtifactSubscriptionTaskResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a temporary username and a token that you can use to log on to a Container Registry instance.
//
// Description:
//
// The ID of the Container Registry instance.
//
// @param request - GetAuthorizationTokenRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAuthorizationTokenResponse
func (client *Client) GetAuthorizationTokenWithOptions(request *GetAuthorizationTokenRequest, runtime *util.RuntimeOptions) (_result *GetAuthorizationTokenResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetAuthorizationToken"),
		Version:     tea.String("2018-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAuthorizationTokenResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a temporary username and a token that you can use to log on to a Container Registry instance.
//
// Description:
//
// The ID of the Container Registry instance.
//
// @param request - GetAuthorizationTokenRequest
//
// @return GetAuthorizationTokenResponse
func (client *Client) GetAuthorizationToken(request *GetAuthorizationTokenRequest) (_result *GetAuthorizationTokenResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAuthorizationTokenResponse{}
	_body, _err := client.GetAuthorizationTokenWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - GetChainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetChainResponse
func (client *Client) GetChainWithOptions(request *GetChainRequest, runtime *util.RuntimeOptions) (_result *GetChainResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ChainId)) {
		query["ChainId"] = request.ChainId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetChain"),
		Version:     tea.String("2018-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetChainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - GetChainRequest
//
// @return GetChainResponse
func (client *Client) GetChain(request *GetChainRequest) (_result *GetChainResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetChainResponse{}
	_body, _err := client.GetChainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a chart namespace in an instance.
//
// @param request - GetChartNamespaceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetChartNamespaceResponse
func (client *Client) GetChartNamespaceWithOptions(request *GetChartNamespaceRequest, runtime *util.RuntimeOptions) (_result *GetChartNamespaceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.NamespaceName)) {
		query["NamespaceName"] = request.NamespaceName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetChartNamespace"),
		Version:     tea.String("2018-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetChartNamespaceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a chart namespace in an instance.
//
// @param request - GetChartNamespaceRequest
//
// @return GetChartNamespaceResponse
func (client *Client) GetChartNamespace(request *GetChartNamespaceRequest) (_result *GetChartNamespaceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetChartNamespaceResponse{}
	_body, _err := client.GetChartNamespaceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about a chart repository.
//
// @param request - GetChartRepositoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetChartRepositoryResponse
func (client *Client) GetChartRepositoryWithOptions(request *GetChartRepositoryRequest, runtime *util.RuntimeOptions) (_result *GetChartRepositoryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RepoName)) {
		query["RepoName"] = request.RepoName
	}

	if !tea.BoolValue(util.IsUnset(request.RepoNamespaceName)) {
		query["RepoNamespaceName"] = request.RepoNamespaceName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetChartRepository"),
		Version:     tea.String("2018-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetChartRepositoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a chart repository.
//
// @param request - GetChartRepositoryRequest
//
// @return GetChartRepositoryResponse
func (client *Client) GetChartRepository(request *GetChartRepositoryRequest) (_result *GetChartRepositoryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetChartRepositoryResponse{}
	_body, _err := client.GetChartRepositoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// The ID of the resource group to which the instance belongs.
//
// @param request - GetInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetInstanceResponse
func (client *Client) GetInstanceWithOptions(request *GetInstanceRequest, runtime *util.RuntimeOptions) (_result *GetInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetInstance"),
		Version:     tea.String("2018-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// The ID of the resource group to which the instance belongs.
//
// @param request - GetInstanceRequest
//
// @return GetInstanceResponse
func (client *Client) GetInstance(request *GetInstanceRequest) (_result *GetInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetInstanceResponse{}
	_body, _err := client.GetInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the number of instances.
//
// @param request - GetInstanceCountRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetInstanceCountResponse
func (client *Client) GetInstanceCountWithOptions(runtime *util.RuntimeOptions) (_result *GetInstanceCountResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("GetInstanceCount"),
		Version:     tea.String("2018-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetInstanceCountResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the number of instances.
//
// @return GetInstanceCountResponse
func (client *Client) GetInstanceCount() (_result *GetInstanceCountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetInstanceCountResponse{}
	_body, _err := client.GetInstanceCountWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the endpoint of an instance.
//
// @param request - GetInstanceEndpointRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetInstanceEndpointResponse
func (client *Client) GetInstanceEndpointWithOptions(request *GetInstanceEndpointRequest, runtime *util.RuntimeOptions) (_result *GetInstanceEndpointResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndpointType)) {
		query["EndpointType"] = request.EndpointType
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.ModuleName)) {
		query["ModuleName"] = request.ModuleName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetInstanceEndpoint"),
		Version:     tea.String("2018-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetInstanceEndpointResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the endpoint of an instance.
//
// @param request - GetInstanceEndpointRequest
//
// @return GetInstanceEndpointResponse
func (client *Client) GetInstanceEndpoint(request *GetInstanceEndpointRequest) (_result *GetInstanceEndpointResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetInstanceEndpointResponse{}
	_body, _err := client.GetInstanceEndpointWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the quota usage of an instance.
//
// @param request - GetInstanceUsageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetInstanceUsageResponse
func (client *Client) GetInstanceUsageWithOptions(request *GetInstanceUsageRequest, runtime *util.RuntimeOptions) (_result *GetInstanceUsageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetInstanceUsage"),
		Version:     tea.String("2018-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetInstanceUsageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the quota usage of an instance.
//
// @param request - GetInstanceUsageRequest
//
// @return GetInstanceUsageResponse
func (client *Client) GetInstanceUsage(request *GetInstanceUsageRequest) (_result *GetInstanceUsageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetInstanceUsageResponse{}
	_body, _err := client.GetInstanceUsageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the endpoints of the virtual private clouds (VPCs) in which a Container Registry instance is deployed.
//
// @param request - GetInstanceVpcEndpointRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetInstanceVpcEndpointResponse
func (client *Client) GetInstanceVpcEndpointWithOptions(request *GetInstanceVpcEndpointRequest, runtime *util.RuntimeOptions) (_result *GetInstanceVpcEndpointResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.ModuleName)) {
		query["ModuleName"] = request.ModuleName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetInstanceVpcEndpoint"),
		Version:     tea.String("2018-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetInstanceVpcEndpointResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the endpoints of the virtual private clouds (VPCs) in which a Container Registry instance is deployed.
//
// @param request - GetInstanceVpcEndpointRequest
//
// @return GetInstanceVpcEndpointResponse
func (client *Client) GetInstanceVpcEndpoint(request *GetInstanceVpcEndpointRequest) (_result *GetInstanceVpcEndpointResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetInstanceVpcEndpointResponse{}
	_body, _err := client.GetInstanceVpcEndpointWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about a namespace.
//
// @param request - GetNamespaceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetNamespaceResponse
func (client *Client) GetNamespaceWithOptions(request *GetNamespaceRequest, runtime *util.RuntimeOptions) (_result *GetNamespaceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.NamespaceId)) {
		query["NamespaceId"] = request.NamespaceId
	}

	if !tea.BoolValue(util.IsUnset(request.NamespaceName)) {
		query["NamespaceName"] = request.NamespaceName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetNamespace"),
		Version:     tea.String("2018-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetNamespaceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a namespace.
//
// @param request - GetNamespaceRequest
//
// @return GetNamespaceResponse
func (client *Client) GetNamespace(request *GetNamespaceRequest) (_result *GetNamespaceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetNamespaceResponse{}
	_body, _err := client.GetNamespaceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about image building records of a repository.
//
// Description:
//
// ***
//
// @param request - GetRepoBuildRecordRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRepoBuildRecordResponse
func (client *Client) GetRepoBuildRecordWithOptions(request *GetRepoBuildRecordRequest, runtime *util.RuntimeOptions) (_result *GetRepoBuildRecordResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BuildRecordId)) {
		query["BuildRecordId"] = request.BuildRecordId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetRepoBuildRecord"),
		Version:     tea.String("2018-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetRepoBuildRecordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about image building records of a repository.
//
// Description:
//
// ***
//
// @param request - GetRepoBuildRecordRequest
//
// @return GetRepoBuildRecordResponse
func (client *Client) GetRepoBuildRecord(request *GetRepoBuildRecordRequest) (_result *GetRepoBuildRecordResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetRepoBuildRecordResponse{}
	_body, _err := client.GetRepoBuildRecordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the status of an image building task.
//
// @param request - GetRepoBuildRecordStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRepoBuildRecordStatusResponse
func (client *Client) GetRepoBuildRecordStatusWithOptions(request *GetRepoBuildRecordStatusRequest, runtime *util.RuntimeOptions) (_result *GetRepoBuildRecordStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BuildRecordId)) {
		query["BuildRecordId"] = request.BuildRecordId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RepoId)) {
		query["RepoId"] = request.RepoId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetRepoBuildRecordStatus"),
		Version:     tea.String("2018-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetRepoBuildRecordStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the status of an image building task.
//
// @param request - GetRepoBuildRecordStatusRequest
//
// @return GetRepoBuildRecordStatusResponse
func (client *Client) GetRepoBuildRecordStatus(request *GetRepoBuildRecordStatusRequest) (_result *GetRepoBuildRecordStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetRepoBuildRecordStatusResponse{}
	_body, _err := client.GetRepoBuildRecordStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about the source code repository that is bound to an image repository.
//
// @param request - GetRepoSourceCodeRepoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRepoSourceCodeRepoResponse
func (client *Client) GetRepoSourceCodeRepoWithOptions(request *GetRepoSourceCodeRepoRequest, runtime *util.RuntimeOptions) (_result *GetRepoSourceCodeRepoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RepoId)) {
		query["RepoId"] = request.RepoId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetRepoSourceCodeRepo"),
		Version:     tea.String("2018-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetRepoSourceCodeRepoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about the source code repository that is bound to an image repository.
//
// @param request - GetRepoSourceCodeRepoRequest
//
// @return GetRepoSourceCodeRepoResponse
func (client *Client) GetRepoSourceCodeRepo(request *GetRepoSourceCodeRepoRequest) (_result *GetRepoSourceCodeRepoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetRepoSourceCodeRepoResponse{}
	_body, _err := client.GetRepoSourceCodeRepoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries an image synchronization task in an instance.
//
// @param request - GetRepoSyncTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRepoSyncTaskResponse
func (client *Client) GetRepoSyncTaskWithOptions(request *GetRepoSyncTaskRequest, runtime *util.RuntimeOptions) (_result *GetRepoSyncTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.SyncTaskId)) {
		query["SyncTaskId"] = request.SyncTaskId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetRepoSyncTask"),
		Version:     tea.String("2018-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetRepoSyncTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries an image synchronization task in an instance.
//
// @param request - GetRepoSyncTaskRequest
//
// @return GetRepoSyncTaskResponse
func (client *Client) GetRepoSyncTask(request *GetRepoSyncTaskRequest) (_result *GetRepoSyncTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetRepoSyncTaskResponse{}
	_body, _err := client.GetRepoSyncTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// The version of the repository.
//
// @param request - GetRepoTagRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRepoTagResponse
func (client *Client) GetRepoTagWithOptions(request *GetRepoTagRequest, runtime *util.RuntimeOptions) (_result *GetRepoTagResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetRepoTag"),
		Version:     tea.String("2018-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetRepoTagResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// The version of the repository.
//
// @param request - GetRepoTagRequest
//
// @return GetRepoTagResponse
func (client *Client) GetRepoTag(request *GetRepoTagRequest) (_result *GetRepoTagResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetRepoTagResponse{}
	_body, _err := client.GetRepoTagWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the scanning status of an image tag.
//
// @param request - GetRepoTagScanStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRepoTagScanStatusResponse
func (client *Client) GetRepoTagScanStatusWithOptions(request *GetRepoTagScanStatusRequest, runtime *util.RuntimeOptions) (_result *GetRepoTagScanStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Digest)) {
		query["Digest"] = request.Digest
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RepoId)) {
		query["RepoId"] = request.RepoId
	}

	if !tea.BoolValue(util.IsUnset(request.ScanTaskId)) {
		query["ScanTaskId"] = request.ScanTaskId
	}

	if !tea.BoolValue(util.IsUnset(request.ScanType)) {
		query["ScanType"] = request.ScanType
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetRepoTagScanStatus"),
		Version:     tea.String("2018-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetRepoTagScanStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the scanning status of an image tag.
//
// @param request - GetRepoTagScanStatusRequest
//
// @return GetRepoTagScanStatusResponse
func (client *Client) GetRepoTagScanStatus(request *GetRepoTagScanStatusRequest) (_result *GetRepoTagScanStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetRepoTagScanStatusResponse{}
	_body, _err := client.GetRepoTagScanStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - GetRepoTagScanSummaryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRepoTagScanSummaryResponse
func (client *Client) GetRepoTagScanSummaryWithOptions(request *GetRepoTagScanSummaryRequest, runtime *util.RuntimeOptions) (_result *GetRepoTagScanSummaryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Digest)) {
		query["Digest"] = request.Digest
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RepoId)) {
		query["RepoId"] = request.RepoId
	}

	if !tea.BoolValue(util.IsUnset(request.ScanTaskId)) {
		query["ScanTaskId"] = request.ScanTaskId
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetRepoTagScanSummary"),
		Version:     tea.String("2018-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetRepoTagScanSummaryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - GetRepoTagScanSummaryRequest
//
// @return GetRepoTagScanSummaryResponse
func (client *Client) GetRepoTagScanSummary(request *GetRepoTagScanSummaryRequest) (_result *GetRepoTagScanSummaryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetRepoTagScanSummaryResponse{}
	_body, _err := client.GetRepoTagScanSummaryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries details about an image repository.
//
// @param request - GetRepositoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRepositoryResponse
func (client *Client) GetRepositoryWithOptions(request *GetRepositoryRequest, runtime *util.RuntimeOptions) (_result *GetRepositoryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RepoId)) {
		query["RepoId"] = request.RepoId
	}

	if !tea.BoolValue(util.IsUnset(request.RepoName)) {
		query["RepoName"] = request.RepoName
	}

	if !tea.BoolValue(util.IsUnset(request.RepoNamespaceName)) {
		query["RepoNamespaceName"] = request.RepoNamespaceName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetRepository"),
		Version:     tea.String("2018-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetRepositoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries details about an image repository.
//
// @param request - GetRepositoryRequest
//
// @return GetRepositoryResponse
func (client *Client) GetRepository(request *GetRepositoryRequest) (_result *GetRepositoryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetRepositoryResponse{}
	_body, _err := client.GetRepositoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the log entries of an artifact building task.
//
// @param request - ListArtifactBuildTaskLogRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListArtifactBuildTaskLogResponse
func (client *Client) ListArtifactBuildTaskLogWithOptions(request *ListArtifactBuildTaskLogRequest, runtime *util.RuntimeOptions) (_result *ListArtifactBuildTaskLogResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListArtifactBuildTaskLog"),
		Version:     tea.String("2018-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListArtifactBuildTaskLogResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the log entries of an artifact building task.
//
// @param request - ListArtifactBuildTaskLogRequest
//
// @return ListArtifactBuildTaskLogResponse
func (client *Client) ListArtifactBuildTaskLog(request *ListArtifactBuildTaskLogRequest) (_result *ListArtifactBuildTaskLogResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListArtifactBuildTaskLogResponse{}
	_body, _err := client.ListArtifactBuildTaskLogWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the lifecycle management rules of an artifact.
//
// @param request - ListArtifactLifecycleRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListArtifactLifecycleRuleResponse
func (client *Client) ListArtifactLifecycleRuleWithOptions(request *ListArtifactLifecycleRuleRequest, runtime *util.RuntimeOptions) (_result *ListArtifactLifecycleRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListArtifactLifecycleRule"),
		Version:     tea.String("2018-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListArtifactLifecycleRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the lifecycle management rules of an artifact.
//
// @param request - ListArtifactLifecycleRuleRequest
//
// @return ListArtifactLifecycleRuleResponse
func (client *Client) ListArtifactLifecycleRule(request *ListArtifactLifecycleRuleRequest) (_result *ListArtifactLifecycleRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListArtifactLifecycleRuleResponse{}
	_body, _err := client.ListArtifactLifecycleRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Lists the subscription rules of artifacts.
//
// @param request - ListArtifactSubscriptionRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListArtifactSubscriptionRuleResponse
func (client *Client) ListArtifactSubscriptionRuleWithOptions(request *ListArtifactSubscriptionRuleRequest, runtime *util.RuntimeOptions) (_result *ListArtifactSubscriptionRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListArtifactSubscriptionRule"),
		Version:     tea.String("2018-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListArtifactSubscriptionRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists the subscription rules of artifacts.
//
// @param request - ListArtifactSubscriptionRuleRequest
//
// @return ListArtifactSubscriptionRuleResponse
func (client *Client) ListArtifactSubscriptionRule(request *ListArtifactSubscriptionRuleRequest) (_result *ListArtifactSubscriptionRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListArtifactSubscriptionRuleResponse{}
	_body, _err := client.ListArtifactSubscriptionRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Lists artifact subscription tasks.
//
// @param request - ListArtifactSubscriptionTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListArtifactSubscriptionTaskResponse
func (client *Client) ListArtifactSubscriptionTaskWithOptions(request *ListArtifactSubscriptionTaskRequest, runtime *util.RuntimeOptions) (_result *ListArtifactSubscriptionTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListArtifactSubscriptionTask"),
		Version:     tea.String("2018-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListArtifactSubscriptionTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists artifact subscription tasks.
//
// @param request - ListArtifactSubscriptionTaskRequest
//
// @return ListArtifactSubscriptionTaskResponse
func (client *Client) ListArtifactSubscriptionTask(request *ListArtifactSubscriptionTaskRequest) (_result *ListArtifactSubscriptionTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListArtifactSubscriptionTaskResponse{}
	_body, _err := client.ListArtifactSubscriptionTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries delivery chains.
//
// @param request - ListChainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListChainResponse
func (client *Client) ListChainWithOptions(request *ListChainRequest, runtime *util.RuntimeOptions) (_result *ListChainResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNo)) {
		query["PageNo"] = request.PageNo
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RepoName)) {
		query["RepoName"] = request.RepoName
	}

	if !tea.BoolValue(util.IsUnset(request.RepoNamespaceName)) {
		query["RepoNamespaceName"] = request.RepoNamespaceName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListChain"),
		Version:     tea.String("2018-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListChainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries delivery chains.
//
// @param request - ListChainRequest
//
// @return ListChainResponse
func (client *Client) ListChain(request *ListChainRequest) (_result *ListChainResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListChainResponse{}
	_body, _err := client.ListChainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// The response code.
//
// @param request - ListChainInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListChainInstanceResponse
func (client *Client) ListChainInstanceWithOptions(request *ListChainInstanceRequest, runtime *util.RuntimeOptions) (_result *ListChainInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNo)) {
		query["PageNo"] = request.PageNo
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RepoName)) {
		query["RepoName"] = request.RepoName
	}

	if !tea.BoolValue(util.IsUnset(request.RepoNamespaceName)) {
		query["RepoNamespaceName"] = request.RepoNamespaceName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListChainInstance"),
		Version:     tea.String("2018-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListChainInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// The response code.
//
// @param request - ListChainInstanceRequest
//
// @return ListChainInstanceResponse
func (client *Client) ListChainInstance(request *ListChainInstanceRequest) (_result *ListChainInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListChainInstanceResponse{}
	_body, _err := client.ListChainInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the chart namespaces of a Container Registry instance.
//
// @param request - ListChartNamespaceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListChartNamespaceResponse
func (client *Client) ListChartNamespaceWithOptions(request *ListChartNamespaceRequest, runtime *util.RuntimeOptions) (_result *ListChartNamespaceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.NamespaceName)) {
		query["NamespaceName"] = request.NamespaceName
	}

	if !tea.BoolValue(util.IsUnset(request.NamespaceStatus)) {
		query["NamespaceStatus"] = request.NamespaceStatus
	}

	if !tea.BoolValue(util.IsUnset(request.PageNo)) {
		query["PageNo"] = request.PageNo
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListChartNamespace"),
		Version:     tea.String("2018-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListChartNamespaceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the chart namespaces of a Container Registry instance.
//
// @param request - ListChartNamespaceRequest
//
// @return ListChartNamespaceResponse
func (client *Client) ListChartNamespace(request *ListChartNamespaceRequest) (_result *ListChartNamespaceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListChartNamespaceResponse{}
	_body, _err := client.ListChartNamespaceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the versions of a chart in a chart repository.
//
// @param request - ListChartReleaseRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListChartReleaseResponse
func (client *Client) ListChartReleaseWithOptions(request *ListChartReleaseRequest, runtime *util.RuntimeOptions) (_result *ListChartReleaseResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Chart)) {
		query["Chart"] = request.Chart
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNo)) {
		query["PageNo"] = request.PageNo
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RepoName)) {
		query["RepoName"] = request.RepoName
	}

	if !tea.BoolValue(util.IsUnset(request.RepoNamespaceName)) {
		query["RepoNamespaceName"] = request.RepoNamespaceName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListChartRelease"),
		Version:     tea.String("2018-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListChartReleaseResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the versions of a chart in a chart repository.
//
// @param request - ListChartReleaseRequest
//
// @return ListChartReleaseResponse
func (client *Client) ListChartRelease(request *ListChartReleaseRequest) (_result *ListChartReleaseResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListChartReleaseResponse{}
	_body, _err := client.ListChartReleaseWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the chart repositories of a Container Registry instance.
//
// @param request - ListChartRepositoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListChartRepositoryResponse
func (client *Client) ListChartRepositoryWithOptions(request *ListChartRepositoryRequest, runtime *util.RuntimeOptions) (_result *ListChartRepositoryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNo)) {
		query["PageNo"] = request.PageNo
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RepoName)) {
		query["RepoName"] = request.RepoName
	}

	if !tea.BoolValue(util.IsUnset(request.RepoNamespaceName)) {
		query["RepoNamespaceName"] = request.RepoNamespaceName
	}

	if !tea.BoolValue(util.IsUnset(request.RepoStatus)) {
		query["RepoStatus"] = request.RepoStatus
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListChartRepository"),
		Version:     tea.String("2018-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListChartRepositoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the chart repositories of a Container Registry instance.
//
// @param request - ListChartRepositoryRequest
//
// @return ListChartRepositoryResponse
func (client *Client) ListChartRepository(request *ListChartRepositoryRequest) (_result *ListChartRepositoryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListChartRepositoryResponse{}
	_body, _err := client.ListChartRepositoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the historical events of an event rule.
//
// @param request - ListEventCenterRecordRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListEventCenterRecordResponse
func (client *Client) ListEventCenterRecordWithOptions(request *ListEventCenterRecordRequest, runtime *util.RuntimeOptions) (_result *ListEventCenterRecordResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListEventCenterRecord"),
		Version:     tea.String("2018-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListEventCenterRecordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the historical events of an event rule.
//
// @param request - ListEventCenterRecordRequest
//
// @return ListEventCenterRecordResponse
func (client *Client) ListEventCenterRecord(request *ListEventCenterRecordRequest) (_result *ListEventCenterRecordResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListEventCenterRecordResponse{}
	_body, _err := client.ListEventCenterRecordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the names of event notification rules.
//
// @param request - ListEventCenterRuleNameRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListEventCenterRuleNameResponse
func (client *Client) ListEventCenterRuleNameWithOptions(request *ListEventCenterRuleNameRequest, runtime *util.RuntimeOptions) (_result *ListEventCenterRuleNameResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListEventCenterRuleName"),
		Version:     tea.String("2018-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListEventCenterRuleNameResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the names of event notification rules.
//
// @param request - ListEventCenterRuleNameRequest
//
// @return ListEventCenterRuleNameResponse
func (client *Client) ListEventCenterRuleName(request *ListEventCenterRuleNameRequest) (_result *ListEventCenterRuleNameResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListEventCenterRuleNameResponse{}
	_body, _err := client.ListEventCenterRuleNameWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries Container Registry instances.
//
// @param request - ListInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListInstanceResponse
func (client *Client) ListInstanceWithOptions(request *ListInstanceRequest, runtime *util.RuntimeOptions) (_result *ListInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceName)) {
		query["InstanceName"] = request.InstanceName
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceStatus)) {
		query["InstanceStatus"] = request.InstanceStatus
	}

	if !tea.BoolValue(util.IsUnset(request.PageNo)) {
		query["PageNo"] = request.PageNo
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListInstance"),
		Version:     tea.String("2018-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries Container Registry instances.
//
// @param request - ListInstanceRequest
//
// @return ListInstanceResponse
func (client *Client) ListInstance(request *ListInstanceRequest) (_result *ListInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListInstanceResponse{}
	_body, _err := client.ListInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the endpoints of a Container Registry instance.
//
// @param request - ListInstanceEndpointRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListInstanceEndpointResponse
func (client *Client) ListInstanceEndpointWithOptions(request *ListInstanceEndpointRequest, runtime *util.RuntimeOptions) (_result *ListInstanceEndpointResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.ModuleName)) {
		query["ModuleName"] = request.ModuleName
	}

	if !tea.BoolValue(util.IsUnset(request.Summary)) {
		query["Summary"] = request.Summary
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListInstanceEndpoint"),
		Version:     tea.String("2018-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListInstanceEndpointResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the endpoints of a Container Registry instance.
//
// @param request - ListInstanceEndpointRequest
//
// @return ListInstanceEndpointResponse
func (client *Client) ListInstanceEndpoint(request *ListInstanceEndpointRequest) (_result *ListInstanceEndpointResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListInstanceEndpointResponse{}
	_body, _err := client.ListInstanceEndpointWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries regions in which you can create Container Registry instances.
//
// @param request - ListInstanceRegionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListInstanceRegionResponse
func (client *Client) ListInstanceRegionWithOptions(request *ListInstanceRegionRequest, runtime *util.RuntimeOptions) (_result *ListInstanceRegionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Lang)) {
		query["Lang"] = request.Lang
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListInstanceRegion"),
		Version:     tea.String("2018-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListInstanceRegionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries regions in which you can create Container Registry instances.
//
// @param request - ListInstanceRegionRequest
//
// @return ListInstanceRegionResponse
func (client *Client) ListInstanceRegion(request *ListInstanceRegionRequest) (_result *ListInstanceRegionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListInstanceRegionResponse{}
	_body, _err := client.ListInstanceRegionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries namespaces in a Container Registry instance.
//
// @param request - ListNamespaceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListNamespaceResponse
func (client *Client) ListNamespaceWithOptions(request *ListNamespaceRequest, runtime *util.RuntimeOptions) (_result *ListNamespaceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.NamespaceName)) {
		query["NamespaceName"] = request.NamespaceName
	}

	if !tea.BoolValue(util.IsUnset(request.NamespaceStatus)) {
		query["NamespaceStatus"] = request.NamespaceStatus
	}

	if !tea.BoolValue(util.IsUnset(request.PageNo)) {
		query["PageNo"] = request.PageNo
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListNamespace"),
		Version:     tea.String("2018-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListNamespaceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries namespaces in a Container Registry instance.
//
// @param request - ListNamespaceRequest
//
// @return ListNamespaceResponse
func (client *Client) ListNamespace(request *ListNamespaceRequest) (_result *ListNamespaceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListNamespaceResponse{}
	_body, _err := client.ListNamespaceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries image building records of an image repository.
//
// @param request - ListRepoBuildRecordRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRepoBuildRecordResponse
func (client *Client) ListRepoBuildRecordWithOptions(request *ListRepoBuildRecordRequest, runtime *util.RuntimeOptions) (_result *ListRepoBuildRecordResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNo)) {
		query["PageNo"] = request.PageNo
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RepoId)) {
		query["RepoId"] = request.RepoId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListRepoBuildRecord"),
		Version:     tea.String("2018-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListRepoBuildRecordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries image building records of an image repository.
//
// @param request - ListRepoBuildRecordRequest
//
// @return ListRepoBuildRecordResponse
func (client *Client) ListRepoBuildRecord(request *ListRepoBuildRecordRequest) (_result *ListRepoBuildRecordResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListRepoBuildRecordResponse{}
	_body, _err := client.ListRepoBuildRecordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the log of an image building record.
//
// @param request - ListRepoBuildRecordLogRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRepoBuildRecordLogResponse
func (client *Client) ListRepoBuildRecordLogWithOptions(request *ListRepoBuildRecordLogRequest, runtime *util.RuntimeOptions) (_result *ListRepoBuildRecordLogResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BuildRecordId)) {
		query["BuildRecordId"] = request.BuildRecordId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Offset)) {
		query["Offset"] = request.Offset
	}

	if !tea.BoolValue(util.IsUnset(request.RepoId)) {
		query["RepoId"] = request.RepoId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListRepoBuildRecordLog"),
		Version:     tea.String("2018-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListRepoBuildRecordLogResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the log of an image building record.
//
// @param request - ListRepoBuildRecordLogRequest
//
// @return ListRepoBuildRecordLogResponse
func (client *Client) ListRepoBuildRecordLog(request *ListRepoBuildRecordLogRequest) (_result *ListRepoBuildRecordLogResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListRepoBuildRecordLogResponse{}
	_body, _err := client.ListRepoBuildRecordLogWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries image building rules of a repository.
//
// @param request - ListRepoBuildRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRepoBuildRuleResponse
func (client *Client) ListRepoBuildRuleWithOptions(request *ListRepoBuildRuleRequest, runtime *util.RuntimeOptions) (_result *ListRepoBuildRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNo)) {
		query["PageNo"] = request.PageNo
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RepoId)) {
		query["RepoId"] = request.RepoId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListRepoBuildRule"),
		Version:     tea.String("2018-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListRepoBuildRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries image building rules of a repository.
//
// @param request - ListRepoBuildRuleRequest
//
// @return ListRepoBuildRuleResponse
func (client *Client) ListRepoBuildRule(request *ListRepoBuildRuleRequest) (_result *ListRepoBuildRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListRepoBuildRuleResponse{}
	_body, _err := client.ListRepoBuildRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries image synchronization rules of a repository.
//
// @param request - ListRepoSyncRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRepoSyncRuleResponse
func (client *Client) ListRepoSyncRuleWithOptions(request *ListRepoSyncRuleRequest, runtime *util.RuntimeOptions) (_result *ListRepoSyncRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.NamespaceName)) {
		query["NamespaceName"] = request.NamespaceName
	}

	if !tea.BoolValue(util.IsUnset(request.PageNo)) {
		query["PageNo"] = request.PageNo
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RepoName)) {
		query["RepoName"] = request.RepoName
	}

	if !tea.BoolValue(util.IsUnset(request.TargetInstanceId)) {
		query["TargetInstanceId"] = request.TargetInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.TargetRegionId)) {
		query["TargetRegionId"] = request.TargetRegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListRepoSyncRule"),
		Version:     tea.String("2018-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListRepoSyncRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries image synchronization rules of a repository.
//
// @param request - ListRepoSyncRuleRequest
//
// @return ListRepoSyncRuleResponse
func (client *Client) ListRepoSyncRule(request *ListRepoSyncRuleRequest) (_result *ListRepoSyncRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListRepoSyncRuleResponse{}
	_body, _err := client.ListRepoSyncRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries image synchronization tasks in an image repository.
//
// @param request - ListRepoSyncTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRepoSyncTaskResponse
func (client *Client) ListRepoSyncTaskWithOptions(request *ListRepoSyncTaskRequest, runtime *util.RuntimeOptions) (_result *ListRepoSyncTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNo)) {
		query["PageNo"] = request.PageNo
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RepoName)) {
		query["RepoName"] = request.RepoName
	}

	if !tea.BoolValue(util.IsUnset(request.RepoNamespaceName)) {
		query["RepoNamespaceName"] = request.RepoNamespaceName
	}

	if !tea.BoolValue(util.IsUnset(request.SyncRecordId)) {
		query["SyncRecordId"] = request.SyncRecordId
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListRepoSyncTask"),
		Version:     tea.String("2018-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListRepoSyncTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries image synchronization tasks in an image repository.
//
// @param request - ListRepoSyncTaskRequest
//
// @return ListRepoSyncTaskResponse
func (client *Client) ListRepoSyncTask(request *ListRepoSyncTaskRequest) (_result *ListRepoSyncTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListRepoSyncTaskResponse{}
	_body, _err := client.ListRepoSyncTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries image tags in a repository.
//
// @param request - ListRepoTagRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRepoTagResponse
func (client *Client) ListRepoTagWithOptions(request *ListRepoTagRequest, runtime *util.RuntimeOptions) (_result *ListRepoTagResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNo)) {
		query["PageNo"] = request.PageNo
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RepoId)) {
		query["RepoId"] = request.RepoId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListRepoTag"),
		Version:     tea.String("2018-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListRepoTagResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries image tags in a repository.
//
// @param request - ListRepoTagRequest
//
// @return ListRepoTagResponse
func (client *Client) ListRepoTag(request *ListRepoTagRequest) (_result *ListRepoTagResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListRepoTagResponse{}
	_body, _err := client.ListRepoTagWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the results of a security scan that is created for an image tag.
//
// @param request - ListRepoTagScanResultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRepoTagScanResultResponse
func (client *Client) ListRepoTagScanResultWithOptions(request *ListRepoTagScanResultRequest, runtime *util.RuntimeOptions) (_result *ListRepoTagScanResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Digest)) {
		query["Digest"] = request.Digest
	}

	if !tea.BoolValue(util.IsUnset(request.FilterValue)) {
		query["FilterValue"] = request.FilterValue
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNo)) {
		query["PageNo"] = request.PageNo
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RepoId)) {
		query["RepoId"] = request.RepoId
	}

	if !tea.BoolValue(util.IsUnset(request.ScanTaskId)) {
		query["ScanTaskId"] = request.ScanTaskId
	}

	if !tea.BoolValue(util.IsUnset(request.ScanType)) {
		query["ScanType"] = request.ScanType
	}

	if !tea.BoolValue(util.IsUnset(request.Severity)) {
		query["Severity"] = request.Severity
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	if !tea.BoolValue(util.IsUnset(request.VulQueryKey)) {
		query["VulQueryKey"] = request.VulQueryKey
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListRepoTagScanResult"),
		Version:     tea.String("2018-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListRepoTagScanResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the results of a security scan that is created for an image tag.
//
// @param request - ListRepoTagScanResultRequest
//
// @return ListRepoTagScanResultResponse
func (client *Client) ListRepoTagScanResult(request *ListRepoTagScanResultRequest) (_result *ListRepoTagScanResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListRepoTagScanResultResponse{}
	_body, _err := client.ListRepoTagScanResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the triggers of a repository.
//
// @param request - ListRepoTriggerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRepoTriggerResponse
func (client *Client) ListRepoTriggerWithOptions(request *ListRepoTriggerRequest, runtime *util.RuntimeOptions) (_result *ListRepoTriggerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RepoId)) {
		query["RepoId"] = request.RepoId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListRepoTrigger"),
		Version:     tea.String("2018-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListRepoTriggerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the triggers of a repository.
//
// @param request - ListRepoTriggerRequest
//
// @return ListRepoTriggerResponse
func (client *Client) ListRepoTrigger(request *ListRepoTriggerRequest) (_result *ListRepoTriggerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListRepoTriggerResponse{}
	_body, _err := client.ListRepoTriggerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries image repositories.
//
// @param request - ListRepositoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRepositoryResponse
func (client *Client) ListRepositoryWithOptions(request *ListRepositoryRequest, runtime *util.RuntimeOptions) (_result *ListRepositoryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNo)) {
		query["PageNo"] = request.PageNo
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RepoName)) {
		query["RepoName"] = request.RepoName
	}

	if !tea.BoolValue(util.IsUnset(request.RepoNamespaceName)) {
		query["RepoNamespaceName"] = request.RepoNamespaceName
	}

	if !tea.BoolValue(util.IsUnset(request.RepoStatus)) {
		query["RepoStatus"] = request.RepoStatus
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListRepository"),
		Version:     tea.String("2018-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListRepositoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries image repositories.
//
// @param request - ListRepositoryRequest
//
// @return ListRepositoryResponse
func (client *Client) ListRepository(request *ListRepositoryRequest) (_result *ListRepositoryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListRepositoryResponse{}
	_body, _err := client.ListRepositoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the baseline risks of a scan task by page.
//
// Description:
//
// Before you call this operation, use a Security Center scan engine to scan the image.
//
// @param request - ListScanBaselineByTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListScanBaselineByTaskResponse
func (client *Client) ListScanBaselineByTaskWithOptions(request *ListScanBaselineByTaskRequest, runtime *util.RuntimeOptions) (_result *ListScanBaselineByTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListScanBaselineByTask"),
		Version:     tea.String("2018-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListScanBaselineByTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the baseline risks of a scan task by page.
//
// Description:
//
// Before you call this operation, use a Security Center scan engine to scan the image.
//
// @param request - ListScanBaselineByTaskRequest
//
// @return ListScanBaselineByTaskResponse
func (client *Client) ListScanBaselineByTask(request *ListScanBaselineByTaskRequest) (_result *ListScanBaselineByTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListScanBaselineByTaskResponse{}
	_body, _err := client.ListScanBaselineByTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the malicious files of a scan task by page.
//
// Description:
//
// Before you call this operation, use a Security Center scan engine to scan the image.
//
// @param request - ListScanMaliciousFileByTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListScanMaliciousFileByTaskResponse
func (client *Client) ListScanMaliciousFileByTaskWithOptions(request *ListScanMaliciousFileByTaskRequest, runtime *util.RuntimeOptions) (_result *ListScanMaliciousFileByTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListScanMaliciousFileByTask"),
		Version:     tea.String("2018-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListScanMaliciousFileByTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the malicious files of a scan task by page.
//
// Description:
//
// Before you call this operation, use a Security Center scan engine to scan the image.
//
// @param request - ListScanMaliciousFileByTaskRequest
//
// @return ListScanMaliciousFileByTaskResponse
func (client *Client) ListScanMaliciousFileByTask(request *ListScanMaliciousFileByTaskRequest) (_result *ListScanMaliciousFileByTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListScanMaliciousFileByTaskResponse{}
	_body, _err := client.ListScanMaliciousFileByTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the tags that are added to cloud resources. Instance resources are supported.
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

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
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
		Version:     tea.String("2018-12-01"),
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
// Queries the tags that are added to cloud resources. Instance resources are supported.
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
// Resets the logon password of a Container Registry instance.
//
// @param request - ResetLoginPasswordRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ResetLoginPasswordResponse
func (client *Client) ResetLoginPasswordWithOptions(request *ResetLoginPasswordRequest, runtime *util.RuntimeOptions) (_result *ResetLoginPasswordResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Password)) {
		query["Password"] = request.Password
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ResetLoginPassword"),
		Version:     tea.String("2018-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ResetLoginPasswordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Resets the logon password of a Container Registry instance.
//
// @param request - ResetLoginPasswordRequest
//
// @return ResetLoginPasswordResponse
func (client *Client) ResetLoginPassword(request *ResetLoginPasswordRequest) (_result *ResetLoginPasswordResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ResetLoginPasswordResponse{}
	_body, _err := client.ResetLoginPasswordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds tags to resources. Instance resources are supported.
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
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
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
		Action:      tea.String("TagResources"),
		Version:     tea.String("2018-12-01"),
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
// Adds tags to resources. Instance resources are supported.
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
// Removes tags from resources. Instance resources are supported.
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

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
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
		Version:     tea.String("2018-12-01"),
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
// Removes tags from resources. Instance resources are supported.
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
// Updates a lifecycle management rule of an artifact.
//
// @param request - UpdateArtifactLifecycleRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateArtifactLifecycleRuleResponse
func (client *Client) UpdateArtifactLifecycleRuleWithOptions(request *UpdateArtifactLifecycleRuleRequest, runtime *util.RuntimeOptions) (_result *UpdateArtifactLifecycleRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Auto)) {
		query["Auto"] = request.Auto
	}

	if !tea.BoolValue(util.IsUnset(request.EnableDeleteTag)) {
		query["EnableDeleteTag"] = request.EnableDeleteTag
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.NamespaceName)) {
		query["NamespaceName"] = request.NamespaceName
	}

	if !tea.BoolValue(util.IsUnset(request.RepoName)) {
		query["RepoName"] = request.RepoName
	}

	if !tea.BoolValue(util.IsUnset(request.RetentionTagCount)) {
		query["RetentionTagCount"] = request.RetentionTagCount
	}

	if !tea.BoolValue(util.IsUnset(request.RuleId)) {
		query["RuleId"] = request.RuleId
	}

	if !tea.BoolValue(util.IsUnset(request.ScheduleTime)) {
		query["ScheduleTime"] = request.ScheduleTime
	}

	if !tea.BoolValue(util.IsUnset(request.Scope)) {
		query["Scope"] = request.Scope
	}

	if !tea.BoolValue(util.IsUnset(request.TagRegexp)) {
		query["TagRegexp"] = request.TagRegexp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateArtifactLifecycleRule"),
		Version:     tea.String("2018-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateArtifactLifecycleRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a lifecycle management rule of an artifact.
//
// @param request - UpdateArtifactLifecycleRuleRequest
//
// @return UpdateArtifactLifecycleRuleResponse
func (client *Client) UpdateArtifactLifecycleRule(request *UpdateArtifactLifecycleRuleRequest) (_result *UpdateArtifactLifecycleRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateArtifactLifecycleRuleResponse{}
	_body, _err := client.UpdateArtifactLifecycleRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates an artifact subscription rule.
//
// @param request - UpdateArtifactSubscriptionRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateArtifactSubscriptionRuleResponse
func (client *Client) UpdateArtifactSubscriptionRuleWithOptions(request *UpdateArtifactSubscriptionRuleRequest, runtime *util.RuntimeOptions) (_result *UpdateArtifactSubscriptionRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Accelerate)) {
		query["Accelerate"] = request.Accelerate
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.NamespaceName)) {
		query["NamespaceName"] = request.NamespaceName
	}

	if !tea.BoolValue(util.IsUnset(request.Override)) {
		query["Override"] = request.Override
	}

	if !tea.BoolValue(util.IsUnset(request.Platform)) {
		query["Platform"] = request.Platform
	}

	if !tea.BoolValue(util.IsUnset(request.RepoName)) {
		query["RepoName"] = request.RepoName
	}

	if !tea.BoolValue(util.IsUnset(request.RuleId)) {
		query["RuleId"] = request.RuleId
	}

	if !tea.BoolValue(util.IsUnset(request.SourceNamespaceName)) {
		query["SourceNamespaceName"] = request.SourceNamespaceName
	}

	if !tea.BoolValue(util.IsUnset(request.SourceProvider)) {
		query["SourceProvider"] = request.SourceProvider
	}

	if !tea.BoolValue(util.IsUnset(request.SourceRepoName)) {
		query["SourceRepoName"] = request.SourceRepoName
	}

	if !tea.BoolValue(util.IsUnset(request.TagCount)) {
		query["TagCount"] = request.TagCount
	}

	if !tea.BoolValue(util.IsUnset(request.TagRegexp)) {
		query["TagRegexp"] = request.TagRegexp
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateArtifactSubscriptionRule"),
		Version:     tea.String("2018-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateArtifactSubscriptionRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates an artifact subscription rule.
//
// @param request - UpdateArtifactSubscriptionRuleRequest
//
// @return UpdateArtifactSubscriptionRuleResponse
func (client *Client) UpdateArtifactSubscriptionRule(request *UpdateArtifactSubscriptionRuleRequest) (_result *UpdateArtifactSubscriptionRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateArtifactSubscriptionRuleResponse{}
	_body, _err := client.UpdateArtifactSubscriptionRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the information about a delivery chain, such as the node execution sequence of the delivery chain.
//
// @param request - UpdateChainRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateChainResponse
func (client *Client) UpdateChainWithOptions(request *UpdateChainRequest, runtime *util.RuntimeOptions) (_result *UpdateChainResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ChainConfig)) {
		query["ChainConfig"] = request.ChainConfig
	}

	if !tea.BoolValue(util.IsUnset(request.ChainId)) {
		query["ChainId"] = request.ChainId
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.ScopeExclude)) {
		query["ScopeExclude"] = request.ScopeExclude
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateChain"),
		Version:     tea.String("2018-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateChainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the information about a delivery chain, such as the node execution sequence of the delivery chain.
//
// @param request - UpdateChainRequest
//
// @return UpdateChainResponse
func (client *Client) UpdateChain(request *UpdateChainRequest) (_result *UpdateChainResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateChainResponse{}
	_body, _err := client.UpdateChainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates a chart namespace.
//
// @param request - UpdateChartNamespaceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateChartNamespaceResponse
func (client *Client) UpdateChartNamespaceWithOptions(request *UpdateChartNamespaceRequest, runtime *util.RuntimeOptions) (_result *UpdateChartNamespaceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AutoCreateRepo)) {
		query["AutoCreateRepo"] = request.AutoCreateRepo
	}

	if !tea.BoolValue(util.IsUnset(request.DefaultRepoType)) {
		query["DefaultRepoType"] = request.DefaultRepoType
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.NamespaceName)) {
		query["NamespaceName"] = request.NamespaceName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateChartNamespace"),
		Version:     tea.String("2018-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateChartNamespaceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a chart namespace.
//
// @param request - UpdateChartNamespaceRequest
//
// @return UpdateChartNamespaceResponse
func (client *Client) UpdateChartNamespace(request *UpdateChartNamespaceRequest) (_result *UpdateChartNamespaceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateChartNamespaceResponse{}
	_body, _err := client.UpdateChartNamespaceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates a chart repository of a Container Registry instance.
//
// @param request - UpdateChartRepositoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateChartRepositoryResponse
func (client *Client) UpdateChartRepositoryWithOptions(request *UpdateChartRepositoryRequest, runtime *util.RuntimeOptions) (_result *UpdateChartRepositoryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RepoName)) {
		query["RepoName"] = request.RepoName
	}

	if !tea.BoolValue(util.IsUnset(request.RepoNamespaceName)) {
		query["RepoNamespaceName"] = request.RepoNamespaceName
	}

	if !tea.BoolValue(util.IsUnset(request.RepoType)) {
		query["RepoType"] = request.RepoType
	}

	if !tea.BoolValue(util.IsUnset(request.Summary)) {
		query["Summary"] = request.Summary
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateChartRepository"),
		Version:     tea.String("2018-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateChartRepositoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a chart repository of a Container Registry instance.
//
// @param request - UpdateChartRepositoryRequest
//
// @return UpdateChartRepositoryResponse
func (client *Client) UpdateChartRepository(request *UpdateChartRepositoryRequest) (_result *UpdateChartRepositoryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateChartRepositoryResponse{}
	_body, _err := client.UpdateChartRepositoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates an event notification rule.
//
// @param tmpReq - UpdateEventCenterRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateEventCenterRuleResponse
func (client *Client) UpdateEventCenterRuleWithOptions(tmpReq *UpdateEventCenterRuleRequest, runtime *util.RuntimeOptions) (_result *UpdateEventCenterRuleResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &UpdateEventCenterRuleShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Namespaces)) {
		request.NamespacesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Namespaces, tea.String("Namespaces"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.RepoNames)) {
		request.RepoNamesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RepoNames, tea.String("RepoNames"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EventChannel)) {
		query["EventChannel"] = request.EventChannel
	}

	if !tea.BoolValue(util.IsUnset(request.EventConfig)) {
		query["EventConfig"] = request.EventConfig
	}

	if !tea.BoolValue(util.IsUnset(request.EventScope)) {
		query["EventScope"] = request.EventScope
	}

	if !tea.BoolValue(util.IsUnset(request.EventType)) {
		query["EventType"] = request.EventType
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.NamespacesShrink)) {
		query["Namespaces"] = request.NamespacesShrink
	}

	if !tea.BoolValue(util.IsUnset(request.RepoNamesShrink)) {
		query["RepoNames"] = request.RepoNamesShrink
	}

	if !tea.BoolValue(util.IsUnset(request.RepoTagFilterPattern)) {
		query["RepoTagFilterPattern"] = request.RepoTagFilterPattern
	}

	if !tea.BoolValue(util.IsUnset(request.RuleId)) {
		query["RuleId"] = request.RuleId
	}

	if !tea.BoolValue(util.IsUnset(request.RuleName)) {
		query["RuleName"] = request.RuleName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateEventCenterRule"),
		Version:     tea.String("2018-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateEventCenterRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates an event notification rule.
//
// @param request - UpdateEventCenterRuleRequest
//
// @return UpdateEventCenterRuleResponse
func (client *Client) UpdateEventCenterRule(request *UpdateEventCenterRuleRequest) (_result *UpdateEventCenterRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateEventCenterRuleResponse{}
	_body, _err := client.UpdateEventCenterRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the status of an instance endpoint.
//
// @param request - UpdateInstanceEndpointStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateInstanceEndpointStatusResponse
func (client *Client) UpdateInstanceEndpointStatusWithOptions(request *UpdateInstanceEndpointStatusRequest, runtime *util.RuntimeOptions) (_result *UpdateInstanceEndpointStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Enable)) {
		query["Enable"] = request.Enable
	}

	if !tea.BoolValue(util.IsUnset(request.EndpointType)) {
		query["EndpointType"] = request.EndpointType
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.ModuleName)) {
		query["ModuleName"] = request.ModuleName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateInstanceEndpointStatus"),
		Version:     tea.String("2018-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateInstanceEndpointStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the status of an instance endpoint.
//
// @param request - UpdateInstanceEndpointStatusRequest
//
// @return UpdateInstanceEndpointStatusResponse
func (client *Client) UpdateInstanceEndpointStatus(request *UpdateInstanceEndpointStatusRequest) (_result *UpdateInstanceEndpointStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateInstanceEndpointStatusResponse{}
	_body, _err := client.UpdateInstanceEndpointStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates a namespace.
//
// @param tmpReq - UpdateNamespaceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateNamespaceResponse
func (client *Client) UpdateNamespaceWithOptions(tmpReq *UpdateNamespaceRequest, runtime *util.RuntimeOptions) (_result *UpdateNamespaceResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &UpdateNamespaceShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.DefaultRepoConfiguration)) {
		request.DefaultRepoConfigurationShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DefaultRepoConfiguration, tea.String("DefaultRepoConfiguration"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AutoCreateRepo)) {
		query["AutoCreateRepo"] = request.AutoCreateRepo
	}

	if !tea.BoolValue(util.IsUnset(request.DefaultRepoConfigurationShrink)) {
		query["DefaultRepoConfiguration"] = request.DefaultRepoConfigurationShrink
	}

	if !tea.BoolValue(util.IsUnset(request.DefaultRepoType)) {
		query["DefaultRepoType"] = request.DefaultRepoType
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.NamespaceName)) {
		query["NamespaceName"] = request.NamespaceName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateNamespace"),
		Version:     tea.String("2018-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateNamespaceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a namespace.
//
// @param request - UpdateNamespaceRequest
//
// @return UpdateNamespaceResponse
func (client *Client) UpdateNamespace(request *UpdateNamespaceRequest) (_result *UpdateNamespaceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateNamespaceResponse{}
	_body, _err := client.UpdateNamespaceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates an image building rule for a repository.
//
// @param request - UpdateRepoBuildRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateRepoBuildRuleResponse
func (client *Client) UpdateRepoBuildRuleWithOptions(request *UpdateRepoBuildRuleRequest, runtime *util.RuntimeOptions) (_result *UpdateRepoBuildRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BuildArgs)) {
		query["BuildArgs"] = request.BuildArgs
	}

	if !tea.BoolValue(util.IsUnset(request.BuildRuleId)) {
		query["BuildRuleId"] = request.BuildRuleId
	}

	if !tea.BoolValue(util.IsUnset(request.DockerfileLocation)) {
		query["DockerfileLocation"] = request.DockerfileLocation
	}

	if !tea.BoolValue(util.IsUnset(request.DockerfileName)) {
		query["DockerfileName"] = request.DockerfileName
	}

	if !tea.BoolValue(util.IsUnset(request.ImageTag)) {
		query["ImageTag"] = request.ImageTag
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Platforms)) {
		query["Platforms"] = request.Platforms
	}

	if !tea.BoolValue(util.IsUnset(request.PushName)) {
		query["PushName"] = request.PushName
	}

	if !tea.BoolValue(util.IsUnset(request.PushType)) {
		query["PushType"] = request.PushType
	}

	if !tea.BoolValue(util.IsUnset(request.RepoId)) {
		query["RepoId"] = request.RepoId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateRepoBuildRule"),
		Version:     tea.String("2018-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateRepoBuildRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates an image building rule for a repository.
//
// @param request - UpdateRepoBuildRuleRequest
//
// @return UpdateRepoBuildRuleResponse
func (client *Client) UpdateRepoBuildRule(request *UpdateRepoBuildRuleRequest) (_result *UpdateRepoBuildRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateRepoBuildRuleResponse{}
	_body, _err := client.UpdateRepoBuildRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the URL of the source code repository that is bound to an image repository.
//
// @param request - UpdateRepoSourceCodeRepoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateRepoSourceCodeRepoResponse
func (client *Client) UpdateRepoSourceCodeRepoWithOptions(request *UpdateRepoSourceCodeRepoRequest, runtime *util.RuntimeOptions) (_result *UpdateRepoSourceCodeRepoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AutoBuild)) {
		query["AutoBuild"] = request.AutoBuild
	}

	if !tea.BoolValue(util.IsUnset(request.CodeRepoId)) {
		query["CodeRepoId"] = request.CodeRepoId
	}

	if !tea.BoolValue(util.IsUnset(request.CodeRepoName)) {
		query["CodeRepoName"] = request.CodeRepoName
	}

	if !tea.BoolValue(util.IsUnset(request.CodeRepoNamespaceName)) {
		query["CodeRepoNamespaceName"] = request.CodeRepoNamespaceName
	}

	if !tea.BoolValue(util.IsUnset(request.CodeRepoType)) {
		query["CodeRepoType"] = request.CodeRepoType
	}

	if !tea.BoolValue(util.IsUnset(request.DisableCacheBuild)) {
		query["DisableCacheBuild"] = request.DisableCacheBuild
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.OverseaBuild)) {
		query["OverseaBuild"] = request.OverseaBuild
	}

	if !tea.BoolValue(util.IsUnset(request.RepoId)) {
		query["RepoId"] = request.RepoId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateRepoSourceCodeRepo"),
		Version:     tea.String("2018-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateRepoSourceCodeRepoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the URL of the source code repository that is bound to an image repository.
//
// @param request - UpdateRepoSourceCodeRepoRequest
//
// @return UpdateRepoSourceCodeRepoResponse
func (client *Client) UpdateRepoSourceCodeRepo(request *UpdateRepoSourceCodeRepoRequest) (_result *UpdateRepoSourceCodeRepoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateRepoSourceCodeRepoResponse{}
	_body, _err := client.UpdateRepoSourceCodeRepoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates a trigger of an image repository.
//
// @param request - UpdateRepoTriggerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateRepoTriggerResponse
func (client *Client) UpdateRepoTriggerWithOptions(request *UpdateRepoTriggerRequest, runtime *util.RuntimeOptions) (_result *UpdateRepoTriggerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RepoId)) {
		query["RepoId"] = request.RepoId
	}

	if !tea.BoolValue(util.IsUnset(request.TriggerId)) {
		query["TriggerId"] = request.TriggerId
	}

	if !tea.BoolValue(util.IsUnset(request.TriggerName)) {
		query["TriggerName"] = request.TriggerName
	}

	if !tea.BoolValue(util.IsUnset(request.TriggerTag)) {
		query["TriggerTag"] = request.TriggerTag
	}

	if !tea.BoolValue(util.IsUnset(request.TriggerType)) {
		query["TriggerType"] = request.TriggerType
	}

	if !tea.BoolValue(util.IsUnset(request.TriggerUrl)) {
		query["TriggerUrl"] = request.TriggerUrl
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateRepoTrigger"),
		Version:     tea.String("2018-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateRepoTriggerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a trigger of an image repository.
//
// @param request - UpdateRepoTriggerRequest
//
// @return UpdateRepoTriggerResponse
func (client *Client) UpdateRepoTrigger(request *UpdateRepoTriggerRequest) (_result *UpdateRepoTriggerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateRepoTriggerResponse{}
	_body, _err := client.UpdateRepoTriggerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// The ID of the request.
//
// @param request - UpdateRepositoryRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateRepositoryResponse
func (client *Client) UpdateRepositoryWithOptions(request *UpdateRepositoryRequest, runtime *util.RuntimeOptions) (_result *UpdateRepositoryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Detail)) {
		query["Detail"] = request.Detail
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.RepoId)) {
		query["RepoId"] = request.RepoId
	}

	if !tea.BoolValue(util.IsUnset(request.RepoName)) {
		query["RepoName"] = request.RepoName
	}

	if !tea.BoolValue(util.IsUnset(request.RepoNamespaceName)) {
		query["RepoNamespaceName"] = request.RepoNamespaceName
	}

	if !tea.BoolValue(util.IsUnset(request.RepoType)) {
		query["RepoType"] = request.RepoType
	}

	if !tea.BoolValue(util.IsUnset(request.Summary)) {
		query["Summary"] = request.Summary
	}

	if !tea.BoolValue(util.IsUnset(request.TagImmutability)) {
		query["TagImmutability"] = request.TagImmutability
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateRepository"),
		Version:     tea.String("2018-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateRepositoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// The ID of the request.
//
// @param request - UpdateRepositoryRequest
//
// @return UpdateRepositoryResponse
func (client *Client) UpdateRepository(request *UpdateRepositoryRequest) (_result *UpdateRepositoryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateRepositoryResponse{}
	_body, _err := client.UpdateRepositoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
