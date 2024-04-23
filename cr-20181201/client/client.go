// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type CancelArtifactBuildTaskRequest struct {
	// The ID of the artifact building task.
	BuildTaskId *string `json:"BuildTaskId,omitempty" xml:"BuildTaskId,omitempty"`
	// The ID of the instance.
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
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// *   `true`: The request is successful.
	// *   `false`: The request fails.
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The ID of the request.
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
	BuildRecordId *string `json:"BuildRecordId,omitempty" xml:"BuildRecordId,omitempty"`
	// The ID of the instance.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the image repository.
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
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the request is successful.
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The ID of the request.
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

type ChangeResourceGroupRequest struct {
	// The ID of the resource group.
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// Id of the request
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The region ID of the resource group.
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
	// *   `ACCELERATED_IMAGE`: accelerated images.
	ArtifactType *string `json:"ArtifactType,omitempty" xml:"ArtifactType,omitempty"`
	// The instance ID.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Additional parameters.
	Parameters map[string]interface{} `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// The ID of the effective range of the rule.
	//
	// *   Set the value to the ID of the image repository.
	ScopeId *string `json:"ScopeId,omitempty" xml:"ScopeId,omitempty"`
	// The effective range of the rule. Valid values:
	//
	// *   `REPOSITORY`
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
	// *   `ACCELERATED_IMAGE`: accelerated images.
	ArtifactType *string `json:"ArtifactType,omitempty" xml:"ArtifactType,omitempty"`
	// The instance ID.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Additional parameters.
	ParametersShrink *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// The ID of the effective range of the rule.
	//
	// *   Set the value to the ID of the image repository.
	ScopeId *string `json:"ScopeId,omitempty" xml:"ScopeId,omitempty"`
	// The effective range of the rule. Valid values:
	//
	// *   `REPOSITORY`
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
	BuildRuleId *string `json:"BuildRuleId,omitempty" xml:"BuildRuleId,omitempty"`
	// The return value.
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the request is successful.
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The request ID.
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
	Auto              *bool   `json:"Auto,omitempty" xml:"Auto,omitempty"`
	EnableDeleteTag   *bool   `json:"EnableDeleteTag,omitempty" xml:"EnableDeleteTag,omitempty"`
	InstanceId        *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	NamespaceName     *string `json:"NamespaceName,omitempty" xml:"NamespaceName,omitempty"`
	RepoName          *string `json:"RepoName,omitempty" xml:"RepoName,omitempty"`
	RetentionTagCount *int64  `json:"RetentionTagCount,omitempty" xml:"RetentionTagCount,omitempty"`
	ScheduleTime      *string `json:"ScheduleTime,omitempty" xml:"ScheduleTime,omitempty"`
	Scope             *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
	TagRegexp         *string `json:"TagRegexp,omitempty" xml:"TagRegexp,omitempty"`
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
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	IsSuccess *bool   `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RuleId    *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
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

type CreateBuildRecordByRecordRequest struct {
	BuildRecordId *string `json:"BuildRecordId,omitempty" xml:"BuildRecordId,omitempty"`
	InstanceId    *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RepoId        *string `json:"RepoId,omitempty" xml:"RepoId,omitempty"`
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
	BuildRecordId *string `json:"BuildRecordId,omitempty" xml:"BuildRecordId,omitempty"`
	Code          *string `json:"Code,omitempty" xml:"Code,omitempty"`
	IsSuccess     *bool   `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	BuildRuleId *string `json:"BuildRuleId,omitempty" xml:"BuildRuleId,omitempty"`
	// The ID of the instance.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the image repository.
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
	BuildRecordId *string `json:"BuildRecordId,omitempty" xml:"BuildRecordId,omitempty"`
	// The return value.
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the request is successful.
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The ID of the request.
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
	ChainConfig *string `json:"ChainConfig,omitempty" xml:"ChainConfig,omitempty"`
	// The description of the delivery chain.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the instance.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the delivery chain.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The name of the repository.
	RepoName *string `json:"RepoName,omitempty" xml:"RepoName,omitempty"`
	// The name of the namespace.
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
	ChainId *string `json:"ChainId,omitempty" xml:"ChainId,omitempty"`
	// The return value.
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the request is successful.
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The ID of the request.
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
	// \-`  true `: automatically creates repositories in the namespace.
	//
	// \-`  false `: does not automatically create repositories in the namespace.
	AutoCreateRepo *bool `json:"AutoCreateRepo,omitempty" xml:"AutoCreateRepo,omitempty"`
	// The default repository type. Valid values:
	//
	// *   `PUBLIC`: a public repository
	// *   `PRIVATE`: a private repository
	DefaultRepoType *string `json:"DefaultRepoType,omitempty" xml:"DefaultRepoType,omitempty"`
	// The ID of the instance.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the namespace.
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
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// *   `true`: The request is successful.
	// *   `false`: The request fails.
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The ID of the request.
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
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the repository.
	RepoName *string `json:"RepoName,omitempty" xml:"RepoName,omitempty"`
	// The name of the namespace to which the repository belongs.
	RepoNamespaceName *string `json:"RepoNamespaceName,omitempty" xml:"RepoNamespaceName,omitempty"`
	// The default repository type. Valid values:
	//
	// *   `PUBLIC`: a public repository.
	// *   `PRIVATE`: a private repository.
	//
	// You can specify the RepoType or Summary parameter. The RepoType parameter is optional.
	RepoType *string `json:"RepoType,omitempty" xml:"RepoType,omitempty"`
	// The summary of the repository.
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
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// *   `true`: The request is successful.
	// *   `false`: The request fails.
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The ID of the repository.
	RepoId *string `json:"RepoId,omitempty" xml:"RepoId,omitempty"`
	// The ID of the request.
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
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The type of the endpoint. Set the value to Internet.
	EndpointType *string `json:"EndpointType,omitempty" xml:"EndpointType,omitempty"`
	// The CIDR block that is accessible.
	Entry *string `json:"Entry,omitempty" xml:"Entry,omitempty"`
	// The ID of the instance.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the module that you want to access. Valid values:
	//
	// *   `Registry`: the image repository.
	// *   `Chart`: a Helm chart.
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
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// *   `true`: The request is successful.
	// *   `false`: The request fails.
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The ID of the request.
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
	// *   `true`
	// *   `false`
	EnableCreateDNSRecordInPvzt *bool `json:"EnableCreateDNSRecordInPvzt,omitempty" xml:"EnableCreateDNSRecordInPvzt,omitempty"`
	// The instance ID.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the module that you want to access. Valid values:
	//
	// *   `Registry`: image repositories.
	// *   `Chart`: Helm charts.
	ModuleName *string `json:"ModuleName,omitempty" xml:"ModuleName,omitempty"`
	// The VPC ID.
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The ID of the vSwitch that is associated with the specified VPC.
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
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the request is successful.
	//
	// *   `true`: The request is successful.
	// *   `false`: The request fails.
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The request ID.
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
	AutoCreateRepo *bool `json:"AutoCreateRepo,omitempty" xml:"AutoCreateRepo,omitempty"`
	// The default type of the repository that is automatically created. Valid values:
	//
	// *   `PUBLIC`: a public repository
	// *   `PRIVATE`: a private repository
	DefaultRepoType *string `json:"DefaultRepoType,omitempty" xml:"DefaultRepoType,omitempty"`
	// The ID of the instance.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the namespace. The name must be 2 to 120 characters in length, and can contain lowercase letters, digits, and the following delimiters: underscores (\_), hyphens (-), and periods (.). The name cannot start or end with a delimiter.
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

type CreateNamespaceResponseBody struct {
	// The return value.
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the request is successful.
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The ID of the request.
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
	DockerfileLocation *string `json:"DockerfileLocation,omitempty" xml:"DockerfileLocation,omitempty"`
	// The name of the Dockerfile.
	DockerfileName *string `json:"DockerfileName,omitempty" xml:"DockerfileName,omitempty"`
	// The tag of the image.
	ImageTag *string `json:"ImageTag,omitempty" xml:"ImageTag,omitempty"`
	// The ID of the instance.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Architecture for image building. Valid values:
	//
	// *   `linux/amd64`
	// *   `linux/arm64`
	// *   `linux/386`
	// *   `linux/arm/v7`
	// *   `inux/arm/v6`
	//
	// Default value: `linux/amd64`
	Platforms []*string `json:"Platforms,omitempty" xml:"Platforms,omitempty" type:"Repeated"`
	// The name of the push that triggers the building rule.
	PushName *string `json:"PushName,omitempty" xml:"PushName,omitempty"`
	// The type of the push that triggers the building rule. Valid values:
	//
	// *   `GIT_TAG`: tag push
	// *   `GIT_BRANCH`: branch push
	PushType *string `json:"PushType,omitempty" xml:"PushType,omitempty"`
	// The ID of the image repository.
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
	BuildRuleId *string `json:"BuildRuleId,omitempty" xml:"BuildRuleId,omitempty"`
	// The return value.
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// *   `true`: The request is successful.
	// *   `false`: The request fails.
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The ID of the request.
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
	// *   `true`: triggers image building when source code is committed.
	// *   `false`: does not trigger image building when source code is committed.
	AutoBuild *bool `json:"AutoBuild,omitempty" xml:"AutoBuild,omitempty"`
	// The name of the source code repository.
	CodeRepoName *string `json:"CodeRepoName,omitempty" xml:"CodeRepoName,omitempty"`
	// The namespace to which the source code repository belongs.
	CodeRepoNamespaceName *string `json:"CodeRepoNamespaceName,omitempty" xml:"CodeRepoNamespaceName,omitempty"`
	// The type of the source code hosting platform. Valid values: `GITHUB`, `GITLAB`, `GITEE`, `CODE`, and `CODEUP`.
	CodeRepoType *string `json:"CodeRepoType,omitempty" xml:"CodeRepoType,omitempty"`
	// Specifies whether to disable building caches. Valid values:
	//
	// *   `true`: disables building caches.
	// *   `false`: enables building caches.
	DisableCacheBuild *bool `json:"DisableCacheBuild,omitempty" xml:"DisableCacheBuild,omitempty"`
	// The ID of the instance.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Specifies whether to enable Build With Servers Deployed Outside Chinese Mainland. Valid values:
	//
	// *   `true`: enables Build With Servers Deployed Outside Chinese Mainland.
	// *   `false`: does not enable Build With Servers Deployed Outside Chinese Mainland.
	OverseaBuild *bool `json:"OverseaBuild,omitempty" xml:"OverseaBuild,omitempty"`
	// The ID of the image repository.
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
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the request is successful.
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The ID of the request.
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
	InstanceId          *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	NamespaceName       *string `json:"NamespaceName,omitempty" xml:"NamespaceName,omitempty"`
	RepoName            *string `json:"RepoName,omitempty" xml:"RepoName,omitempty"`
	SyncRuleName        *string `json:"SyncRuleName,omitempty" xml:"SyncRuleName,omitempty"`
	SyncScope           *string `json:"SyncScope,omitempty" xml:"SyncScope,omitempty"`
	SyncTrigger         *string `json:"SyncTrigger,omitempty" xml:"SyncTrigger,omitempty"`
	TagFilter           *string `json:"TagFilter,omitempty" xml:"TagFilter,omitempty"`
	TargetInstanceId    *string `json:"TargetInstanceId,omitempty" xml:"TargetInstanceId,omitempty"`
	TargetNamespaceName *string `json:"TargetNamespaceName,omitempty" xml:"TargetNamespaceName,omitempty"`
	TargetRegionId      *string `json:"TargetRegionId,omitempty" xml:"TargetRegionId,omitempty"`
	TargetRepoName      *string `json:"TargetRepoName,omitempty" xml:"TargetRepoName,omitempty"`
	TargetUserId        *string `json:"TargetUserId,omitempty" xml:"TargetUserId,omitempty"`
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
	Code       *string `json:"Code,omitempty" xml:"Code,omitempty"`
	IsSuccess  *bool   `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	InstanceId       *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Override         *bool   `json:"Override,omitempty" xml:"Override,omitempty"`
	RepoId           *string `json:"RepoId,omitempty" xml:"RepoId,omitempty"`
	Tag              *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	TargetInstanceId *string `json:"TargetInstanceId,omitempty" xml:"TargetInstanceId,omitempty"`
	TargetNamespace  *string `json:"TargetNamespace,omitempty" xml:"TargetNamespace,omitempty"`
	TargetRegionId   *string `json:"TargetRegionId,omitempty" xml:"TargetRegionId,omitempty"`
	TargetRepoName   *string `json:"TargetRepoName,omitempty" xml:"TargetRepoName,omitempty"`
	TargetTag        *string `json:"TargetTag,omitempty" xml:"TargetTag,omitempty"`
	TargetUserId     *string `json:"TargetUserId,omitempty" xml:"TargetUserId,omitempty"`
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
	Code       *string `json:"Code,omitempty" xml:"Code,omitempty"`
	IsSuccess  *bool   `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the image repository.
	RepoId *string `json:"RepoId,omitempty" xml:"RepoId,omitempty"`
	// The ID of the synchronization rule.
	SyncRuleId *string `json:"SyncRuleId,omitempty" xml:"SyncRuleId,omitempty"`
	// The version of the image to be synchronized.
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
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// *   `true`: The request is successful.
	// *   `false`: The request fails.
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the synchronization task.
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
	FromTag *string `json:"FromTag,omitempty" xml:"FromTag,omitempty"`
	// The ID of the instance.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the namespace.
	NamespaceName *string `json:"NamespaceName,omitempty" xml:"NamespaceName,omitempty"`
	// The name of the image repository.
	RepoName *string `json:"RepoName,omitempty" xml:"RepoName,omitempty"`
	// The image tag that you want to create.
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
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the request is successful.
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The ID of the request.
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
	Digest *string `json:"Digest,omitempty" xml:"Digest,omitempty"`
	// The ID of the instance.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the image repository.
	RepoId *string `json:"RepoId,omitempty" xml:"RepoId,omitempty"`
	// The type of the scanning engine.
	//
	// *   `SAS_SCAN_SERVICE`: Security Center scan engine (paid service)
	// *   `ACR_SCAN_SERVICE`: Container Registry scan engine
	ScanService *string `json:"ScanService,omitempty" xml:"ScanService,omitempty"`
	// The version of the image.
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

func (s *CreateRepoTagScanTaskRequest) SetTag(v string) *CreateRepoTagScanTaskRequest {
	s.Tag = &v
	return s
}

type CreateRepoTagScanTaskResponseBody struct {
	// The return value.
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// *   `true`: The request is successful.
	// *   `false`: The request fails.
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The ID of the request.
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
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the image repository.
	RepoId *string `json:"RepoId,omitempty" xml:"RepoId,omitempty"`
	// The name of the trigger.
	TriggerName *string `json:"TriggerName,omitempty" xml:"TriggerName,omitempty"`
	// The image tag based on which the trigger is set.
	//
	// >
	//
	// *   If `TriggerType` is set to `ALL`, `TriggerTag` can be set to a string or an array, for example, `*`.
	//
	// *   If `TriggerType` is set to `TAG_LIST`, `TriggerTag` must be set to an array, for example, `[1]`.
	// *   If `TriggerType` is set to `TAG_REG_EXP`, `TriggerTag` must be set to a string, for example, `*`.
	TriggerTag *string `json:"TriggerTag,omitempty" xml:"TriggerTag,omitempty"`
	// The type of the trigger. Valid values:
	//
	// *   `ALL`: a trigger that supports both tags and regular expressions.
	// *   `TAG_LIST`: a tag-based trigger.
	// *   `TAG_REG_EXP`: a regular expression-based trigger.
	TriggerType *string `json:"TriggerType,omitempty" xml:"TriggerType,omitempty"`
	// The URL of the trigger.
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
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the request is successful.
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the trigger.
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
	Detail *string `json:"Detail,omitempty" xml:"Detail,omitempty"`
	// The ID of the instance.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the image repository.
	RepoName *string `json:"RepoName,omitempty" xml:"RepoName,omitempty"`
	// The name of the namespace to which the image repository belongs.
	RepoNamespaceName *string `json:"RepoNamespaceName,omitempty" xml:"RepoNamespaceName,omitempty"`
	// The type of the repository. Valid values:
	//
	// *   `PUBLIC`: The repository is a public repository.
	// *   `PRIVATE`: The repository is a private repository.
	RepoType *string `json:"RepoType,omitempty" xml:"RepoType,omitempty"`
	// The summary about the repository.
	Summary *string `json:"Summary,omitempty" xml:"Summary,omitempty"`
	// Specifies whether to enable the feature of image tag immutability. Valid values:
	//
	// *   `true`: enables the feature of image tag immutability.
	// *   `false`: disables the feature of image tag immutability.
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
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the request is successful.
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The ID of the image repository.
	RepoId *string `json:"RepoId,omitempty" xml:"RepoId,omitempty"`
	// The ID of the request.
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
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RuleId     *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
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
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	IsSuccess *bool   `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
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

type DeleteChainRequest struct {
	// The ID of the delivery pipeline.
	ChainId *string `json:"ChainId,omitempty" xml:"ChainId,omitempty"`
	// The ID of the instance.
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
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the request is successful.
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The ID of the request.
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
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the chart namespace that you want to delete.
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
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// *   `true`: The request is successful.
	// *   `false`: The request fails.
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The ID of the request.
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
	Chart *string `json:"Chart,omitempty" xml:"Chart,omitempty"`
	// The ID of the instance.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The version of the chart that you want to delete.
	Release *string `json:"Release,omitempty" xml:"Release,omitempty"`
	// The name of the repository.
	RepoName *string `json:"RepoName,omitempty" xml:"RepoName,omitempty"`
	// The name of the namespace.
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
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the request is successful.
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The ID of the request.
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
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the repository.
	RepoName *string `json:"RepoName,omitempty" xml:"RepoName,omitempty"`
	// The name of the namespace to which the repository belongs.
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
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// *   `true`: The request is successful.
	// *   `false`: The request fails.
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The ID of the request.
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
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the event notification rule.
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
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The ID of the request.
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
	EndpointType *string `json:"EndpointType,omitempty" xml:"EndpointType,omitempty"`
	// The CIDR block.
	Entry *string `json:"Entry,omitempty" xml:"Entry,omitempty"`
	// The ID of the instance.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the module that you want to access. Valid values:
	//
	// *   `Registry`: the image repository.
	// *   `Chart`: a Helm chart.
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
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// *   `true`: The request is successful.
	// *   `false`: The request fails.
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The ID of the request.
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
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the module that you want to access. Valid values:
	//
	// *   `Registry`: the image repository.
	// *   `Chart`: a Helm chart.
	ModuleName *string `json:"ModuleName,omitempty" xml:"ModuleName,omitempty"`
	// The ID of the VPC.
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The ID of the vSwitch.
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
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// *   `true`: The request is successful.
	// *   `false`: The request fails.
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The ID of the request.
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
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the namespace.
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
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the request is successful.
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The ID of the request.
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
	BuildRuleId *string `json:"BuildRuleId,omitempty" xml:"BuildRuleId,omitempty"`
	// The ID of the instance.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the image repository.
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
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the request is successful.
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The ID of the request.
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
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the synchronization rule.
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
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the request is successful.
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The ID of the request.
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
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the image repository.
	RepoId *string `json:"RepoId,omitempty" xml:"RepoId,omitempty"`
	// The tag of the image.
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
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the request is successful.
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The ID of the request.
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
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the image repository.
	RepoId *string `json:"RepoId,omitempty" xml:"RepoId,omitempty"`
	// The ID of the trigger.
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
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// *   `true`: The request is successful.
	// *   `false`: The request fails.
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The ID of the request.
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
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the image repository.
	RepoId *string `json:"RepoId,omitempty" xml:"RepoId,omitempty"`
	// The name of the repository.
	RepoName *string `json:"RepoName,omitempty" xml:"RepoName,omitempty"`
	// The name of the namespace to which the repository belongs.
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
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// *   `true`: The request is successful.
	// *   `false`: The request fails.
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The request ID.
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
	ArtifactType *string `json:"ArtifactType,omitempty" xml:"ArtifactType,omitempty"`
	BuildRuleId  *string `json:"BuildRuleId,omitempty" xml:"BuildRuleId,omitempty"`
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	ScopeId      *string `json:"ScopeId,omitempty" xml:"ScopeId,omitempty"`
	ScopeType    *string `json:"ScopeType,omitempty" xml:"ScopeType,omitempty"`
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
	ArtifactType *string                                     `json:"ArtifactType,omitempty" xml:"ArtifactType,omitempty"`
	BuildRuleId  *string                                     `json:"BuildRuleId,omitempty" xml:"BuildRuleId,omitempty"`
	Code         *string                                     `json:"Code,omitempty" xml:"Code,omitempty"`
	IsSuccess    *bool                                       `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	Parameters   *GetArtifactBuildRuleResponseBodyParameters `json:"Parameters,omitempty" xml:"Parameters,omitempty" type:"Struct"`
	RequestId    *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ScopeId      *string                                     `json:"ScopeId,omitempty" xml:"ScopeId,omitempty"`
	ScopeType    *string                                     `json:"ScopeType,omitempty" xml:"ScopeType,omitempty"`
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
	ImageIndexOnly *bool `json:"ImageIndexOnly,omitempty" xml:"ImageIndexOnly,omitempty"`
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
	BuildTaskId *string `json:"BuildTaskId,omitempty" xml:"BuildTaskId,omitempty"`
	// The ID of the instance.
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
	// *   `IMAGE_TO_ACCELERATED_IMAGE`: builds accelerated images for Container Service for Kubernetes (ACK) clusters.
	// *   `IMAGE_TO_ECI_ACCELERATED_IMAGE`: builds accelerated images for elastic container instances.
	ArtifactBuildType *string `json:"ArtifactBuildType,omitempty" xml:"ArtifactBuildType,omitempty"`
	// The ID of the artifact building task.
	BuildTaskId *string `json:"BuildTaskId,omitempty" xml:"BuildTaskId,omitempty"`
	// The return value.
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The time when the artifact building task ends.
	EndTime      *int32    `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Instructions []*string `json:"Instructions,omitempty" xml:"Instructions,omitempty" type:"Repeated"`
	// Indicates whether the request is successful.
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the source artifact.
	SourceArtifact *GetArtifactBuildTaskResponseBodySourceArtifact `json:"SourceArtifact,omitempty" xml:"SourceArtifact,omitempty" type:"Struct"`
	// The time when the artifact building task starts.
	StartTime *int32 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The artifact that is built in the task.
	TargetArtifact *GetArtifactBuildTaskResponseBodyTargetArtifact `json:"TargetArtifact,omitempty" xml:"TargetArtifact,omitempty" type:"Struct"`
	// The status of the artifact that is built in the task. Valid values:
	//
	// *   `PENDING`: The artifact is being scheduled.
	// *   `BUILDING`: The artifact is being built.
	// *   `SUCCESS`: The artifact is built.
	// *   `FAILED`: The artifact fails to be built.
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
	ArtifactType *string `json:"ArtifactType,omitempty" xml:"ArtifactType,omitempty"`
	// The ID of the repository to which the source artifact belongs. The repository can only be an image repository.
	RepoId *string `json:"RepoId,omitempty" xml:"RepoId,omitempty"`
	// The version of the artifact. The artifact can only be an image.
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
	ArtifactType *string `json:"ArtifactType,omitempty" xml:"ArtifactType,omitempty"`
	// The ID of the repository to which the artifact that is built in the task belongs. The repository can only be an image repository. The value is the same as the ID of the repository to which the source artifact belongs.
	RepoId *string `json:"RepoId,omitempty" xml:"RepoId,omitempty"`
	// The version of the artifact that is built in the task. The artifact can only be an image.
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
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RuleId     *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
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
	Auto              *bool   `json:"Auto,omitempty" xml:"Auto,omitempty"`
	Code              *string `json:"Code,omitempty" xml:"Code,omitempty"`
	CreateTime        *int64  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	EnableDeleteTag   *bool   `json:"EnableDeleteTag,omitempty" xml:"EnableDeleteTag,omitempty"`
	InstanceId        *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	IsSuccess         *bool   `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	ModifiedTime      *int64  `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	NamespaceName     *string `json:"NamespaceName,omitempty" xml:"NamespaceName,omitempty"`
	NextTime          *int64  `json:"NextTime,omitempty" xml:"NextTime,omitempty"`
	RepoName          *string `json:"RepoName,omitempty" xml:"RepoName,omitempty"`
	RequestId         *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RetentionTagCount *int64  `json:"RetentionTagCount,omitempty" xml:"RetentionTagCount,omitempty"`
	RuleId            *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	ScheduleTime      *string `json:"ScheduleTime,omitempty" xml:"ScheduleTime,omitempty"`
	Scope             *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
	TagRegexp         *string `json:"TagRegexp,omitempty" xml:"TagRegexp,omitempty"`
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

type GetAuthorizationTokenRequest struct {
	// The ID of the request.
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
	// *   If the STS token belongs to an Alibaba Cloud account, you can use the STS token to log on to all Container Registry instances that belong to the Alibaba Cloud account.
	// *   If the STS token belongs to a Resource Access Management (RAM) user, you can use the STS token to log on to all Container Registry instances that belong to the RAM user.
	// *   You can use an STS token to access only Container Registry instances to which the STS token is scoped.
	AuthorizationToken *string `json:"AuthorizationToken,omitempty" xml:"AuthorizationToken,omitempty"`
	// Indicates whether the API call is successful.
	//
	// *   `true`: successful
	// *   `false`: failed
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The return value.
	ExpireTime *int64 `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The username that is used to log on to the Container Registry instance.
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The timestamp when the temporary password expires. Unit: milliseconds.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The password that is used to log on to the Container Registry instance.
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
	ChainId    *string `json:"ChainId,omitempty" xml:"ChainId,omitempty"`
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
	ChainConfig  *GetChainResponseBodyChainConfig `json:"ChainConfig,omitempty" xml:"ChainConfig,omitempty" type:"Struct"`
	ChainId      *string                          `json:"ChainId,omitempty" xml:"ChainId,omitempty"`
	Code         *string                          `json:"Code,omitempty" xml:"Code,omitempty"`
	CreateTime   *int64                           `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description  *string                          `json:"Description,omitempty" xml:"Description,omitempty"`
	InstanceId   *string                          `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	IsSuccess    *bool                            `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	ModifiedTime *int64                           `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	Name         *string                          `json:"Name,omitempty" xml:"Name,omitempty"`
	RequestId    *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ScopeExclude []*string                        `json:"ScopeExclude,omitempty" xml:"ScopeExclude,omitempty" type:"Repeated"`
	ScopeId      *string                          `json:"ScopeId,omitempty" xml:"ScopeId,omitempty"`
	ScopeType    *string                          `json:"ScopeType,omitempty" xml:"ScopeType,omitempty"`
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
	ChainConfigId *string                                   `json:"ChainConfigId,omitempty" xml:"ChainConfigId,omitempty"`
	IsActive      *bool                                     `json:"IsActive,omitempty" xml:"IsActive,omitempty"`
	Nodes         []*GetChainResponseBodyChainConfigNodes   `json:"Nodes,omitempty" xml:"Nodes,omitempty" type:"Repeated"`
	Routers       []*GetChainResponseBodyChainConfigRouters `json:"Routers,omitempty" xml:"Routers,omitempty" type:"Repeated"`
	Version       *string                                   `json:"Version,omitempty" xml:"Version,omitempty"`
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
	Enable     *bool                                           `json:"Enable,omitempty" xml:"Enable,omitempty"`
	NodeConfig *GetChainResponseBodyChainConfigNodesNodeConfig `json:"NodeConfig,omitempty" xml:"NodeConfig,omitempty" type:"Struct"`
	NodeName   *string                                         `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
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
	Retry      *int32                                                    `json:"Retry,omitempty" xml:"Retry,omitempty"`
	ScanEngine *string                                                   `json:"ScanEngine,omitempty" xml:"ScanEngine,omitempty"`
	Timeout    *int64                                                    `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
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
	Action        *string `json:"Action,omitempty" xml:"Action,omitempty"`
	BaselineList  *string `json:"BaselineList,omitempty" xml:"BaselineList,omitempty"`
	IssueCount    *string `json:"IssueCount,omitempty" xml:"IssueCount,omitempty"`
	IssueLevel    *string `json:"IssueLevel,omitempty" xml:"IssueLevel,omitempty"`
	IssueList     *string `json:"IssueList,omitempty" xml:"IssueList,omitempty"`
	Logic         *string `json:"Logic,omitempty" xml:"Logic,omitempty"`
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
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the namespace.
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
	// *   `true`: A repository was automatically created in the namespace.
	// *   `false`: No repository was automatically created in the namespace.
	AutoCreateRepo *bool `json:"AutoCreateRepo,omitempty" xml:"AutoCreateRepo,omitempty"`
	// The return value.
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The default repository type. Valid values:
	//
	// *   `PUBLIC`: a public repository.
	// *   `PRIVATE`: a private repository.
	DefaultRepoType *string `json:"DefaultRepoType,omitempty" xml:"DefaultRepoType,omitempty"`
	// The ID of the instance.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// *   `true`: The request is successful.
	// *   `false`: The request fails.
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The ID of the namespace.
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	// The name of the namespace.
	NamespaceName *string `json:"NamespaceName,omitempty" xml:"NamespaceName,omitempty"`
	// The status of the namespace. Valid values:
	//
	// *   `NORMAL`: The namespace is normal.
	// *   `DELETING`: The namespace is being deleted.
	NamespaceStatus *string `json:"NamespaceStatus,omitempty" xml:"NamespaceStatus,omitempty"`
	// The ID of the request.
	RequestId       *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// The ID of the instance.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the repository.
	RepoName *string `json:"RepoName,omitempty" xml:"RepoName,omitempty"`
	// The name of the namespace.
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
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The time when the chart repository was created.
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The ID of the instance.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// *   `true`: The request is successful.
	// *   `false`: The request fails.
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The time when the chart repository was last modified.
	ModifiedTime *int64 `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// The ID of the chart repository.
	RepoId *string `json:"RepoId,omitempty" xml:"RepoId,omitempty"`
	// The name of the chart repository.
	RepoName *string `json:"RepoName,omitempty" xml:"RepoName,omitempty"`
	// The name of the namespace to which the chart repository belongs.
	RepoNamespaceName *string `json:"RepoNamespaceName,omitempty" xml:"RepoNamespaceName,omitempty"`
	// The status of the chart repository. Valid values:
	//
	// *   `NORMAL`: The repository is normal.
	// *   `DELETING`: The repository is being deleted.
	RepoStatus *string `json:"RepoStatus,omitempty" xml:"RepoStatus,omitempty"`
	// The type of the chart repository. Valid values:
	//
	// *   `PUBLIC`: a public repository
	// *   `PRIVATE`: a private repository
	RepoType *string `json:"RepoType,omitempty" xml:"RepoType,omitempty"`
	// The ID of the request.
	RequestId       *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The summary about the chart repository.
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
	Code                  *string                        `json:"Code,omitempty" xml:"Code,omitempty"`
	CreateTime            *int64                         `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	InstanceId            *string                        `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceIssue         *string                        `json:"InstanceIssue,omitempty" xml:"InstanceIssue,omitempty"`
	InstanceName          *string                        `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	InstanceSpecification *string                        `json:"InstanceSpecification,omitempty" xml:"InstanceSpecification,omitempty"`
	InstanceStatus        *string                        `json:"InstanceStatus,omitempty" xml:"InstanceStatus,omitempty"`
	IsSuccess             *bool                          `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	ModifiedTime          *int64                         `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	RequestId             *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResourceGroupId       *string                        `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Tags                  []*GetInstanceResponseBodyTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
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
	TagKey   *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
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
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Count     *int32  `json:"Count,omitempty" xml:"Count,omitempty"`
	IsSuccess *bool   `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
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
	EndpointType *string `json:"EndpointType,omitempty" xml:"EndpointType,omitempty"`
	// The ID of the instance.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the module that you want to access. Valid values:
	//
	// *   `Registry`: the image repository.
	// *   `Chart`: a Helm chart.
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
	AclEnable *bool `json:"AclEnable,omitempty" xml:"AclEnable,omitempty"`
	// The ACLs.
	AclEntries []*GetInstanceEndpointResponseBodyAclEntries `json:"AclEntries,omitempty" xml:"AclEntries,omitempty" type:"Repeated"`
	// The return value.
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Domain names.
	Domains []*GetInstanceEndpointResponseBodyDomains `json:"Domains,omitempty" xml:"Domains,omitempty" type:"Repeated"`
	// Indicates whether the ACL feature is enabled.
	Enable *bool `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// Indicates whether the request is successful.
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The status of the instance.
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
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The public IP address whitelist.
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
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The type of the domain name. Valid values:
	//
	// *   `SYSTEM`: a system domain name.
	// *   `USER`: a user domain name.
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
	ChartNamespaceQuota *string `json:"ChartNamespaceQuota,omitempty" xml:"ChartNamespaceQuota,omitempty"`
	// The number of chart namespaces that are created in the instance.
	ChartNamespaceUsage *string `json:"ChartNamespaceUsage,omitempty" xml:"ChartNamespaceUsage,omitempty"`
	// The quota of chart repositories for the instance.
	ChartRepoQuota *string `json:"ChartRepoQuota,omitempty" xml:"ChartRepoQuota,omitempty"`
	// The number of chart repositories that are created.
	ChartRepoUsage *string `json:"ChartRepoUsage,omitempty" xml:"ChartRepoUsage,omitempty"`
	// The return value.
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// *   `true`: The request is successful.
	// *   `false`: The request fails.
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The quota of image namespaces for the instance.
	NamespaceQuota *string `json:"NamespaceQuota,omitempty" xml:"NamespaceQuota,omitempty"`
	// The number of image namespaces that are created in the instance.
	NamespaceUsage *string `json:"NamespaceUsage,omitempty" xml:"NamespaceUsage,omitempty"`
	// The quota of image repositories for the instance.
	RepoQuota *string `json:"RepoQuota,omitempty" xml:"RepoQuota,omitempty"`
	// The number of image repositories that are created in the instance.
	RepoUsage *string `json:"RepoUsage,omitempty" xml:"RepoUsage,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the module that you want to access. Valid values:
	//
	// *   `Registry`: the image repository.
	// *   `Chart`: a Helm chart.
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
	Code    *string   `json:"Code,omitempty" xml:"Code,omitempty"`
	Domains []*string `json:"Domains,omitempty" xml:"Domains,omitempty" type:"Repeated"`
	// Indicates whether the access control list (ACL) feature is enabled. Valid values:
	//
	// *   `true`: The ACL feature is enabled.
	// *   `false`: The ACL feature is disabled.
	Enable *bool `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// *   `true`: The request is successful.
	// *   `false`: The request fails.
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The VPCs in which the instance is deployed.
	LinkedVpcs []*GetInstanceVpcEndpointResponseBodyLinkedVpcs `json:"LinkedVpcs,omitempty" xml:"LinkedVpcs,omitempty" type:"Repeated"`
	ModuleName *string                                         `json:"ModuleName,omitempty" xml:"ModuleName,omitempty"`
	// The ID of the request.
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
	// Indicates whether the default ACL is used.
	DefaultAccess *bool `json:"DefaultAccess,omitempty" xml:"DefaultAccess,omitempty"`
	// IP address.
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// The status of the VPC. Valid values:
	//
	// *   `CREATING`: The VPC is being created.
	// *   `RUNNING`: The VPC is running.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// VPC ID
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The ID of the vSwitch.
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
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the namespace.
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	// The name of the namespace.
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
	AutoCreateRepo *bool `json:"AutoCreateRepo,omitempty" xml:"AutoCreateRepo,omitempty"`
	// The return value.
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The default type of the repository. Valid values:
	//
	// *   PUBLIC: The repository is a public repository.
	// *   PRIVATE: The repository is a private repository.
	DefaultRepoType *string `json:"DefaultRepoType,omitempty" xml:"DefaultRepoType,omitempty"`
	// The ID of the instance.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Indicates whether the request is successful.
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The ID of the namespace.
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	// The name of the namespace.
	NamespaceName *string `json:"NamespaceName,omitempty" xml:"NamespaceName,omitempty"`
	// The status of the namespace.
	NamespaceStatus *string `json:"NamespaceStatus,omitempty" xml:"NamespaceStatus,omitempty"`
	// The ID of the request.
	RequestId       *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	BuildRecordId *string `json:"BuildRecordId,omitempty" xml:"BuildRecordId,omitempty"`
	// The ID of the instance.
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
	BuildRecordId *string `json:"BuildRecordId,omitempty" xml:"BuildRecordId,omitempty"`
	// The return value.
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The time when the image building was completed.
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The information about the image.
	Image *GetRepoBuildRecordResponseBodyImage `json:"Image,omitempty" xml:"Image,omitempty" type:"Struct"`
	// Indicates whether the request is successful.
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The time when the image building started.
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The status of the instance.
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
	ImageTag *string `json:"ImageTag,omitempty" xml:"ImageTag,omitempty"`
	// The name of the image repository.
	RepoName *string `json:"RepoName,omitempty" xml:"RepoName,omitempty"`
	// The name of the namespace to which the image repository belongs.
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
	BuildRecordId *string `json:"BuildRecordId,omitempty" xml:"BuildRecordId,omitempty"`
	// The ID of the instance.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the image repository.
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
	BuildStatus *string `json:"BuildStatus,omitempty" xml:"BuildStatus,omitempty"`
	// The return value.
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the request is successful.
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The ID of the request.
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
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the repository.
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
	// *   `true`: Image building is automatically triggered when source code is committed.
	// *   `false`: Image building is not triggered when source code is committed.
	AutoBuild *string `json:"AutoBuild,omitempty" xml:"AutoBuild,omitempty"`
	// The response code.
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The address of the source code repository.
	CodeRepoDomain *string `json:"CodeRepoDomain,omitempty" xml:"CodeRepoDomain,omitempty"`
	// The name of the source code repository.
	CodeRepoName *string `json:"CodeRepoName,omitempty" xml:"CodeRepoName,omitempty"`
	// The namespace to which the source code repository belongs.
	CodeRepoNamespaceName *string `json:"CodeRepoNamespaceName,omitempty" xml:"CodeRepoNamespaceName,omitempty"`
	// The type of the code hosting platform. Valid values: `GITHUB`, `GITLAB`, `GITEE`, `CODE`, and `CODEUP`.
	CodeRepoType *string `json:"CodeRepoType,omitempty" xml:"CodeRepoType,omitempty"`
	// Indicates whether build cache is disabled. Valid values:
	//
	// *   `true`: Build cache is disabled.
	// *   `false`: Build cache is enabled.
	DisableCacheBuild *string `json:"DisableCacheBuild,omitempty" xml:"DisableCacheBuild,omitempty"`
	// Indicates whether the API call is successful. Valid values:
	//
	// *   `true`: successful
	// *   `false`: failed
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// Indicates whether image building is accelerated for servers outside the Chinese mainland. Valid values:
	//
	// *   `true`: Image building is accelerated for servers outside the Chinese mainland.
	// *   `false`: Image building is not accelerated for servers outside the Chinese mainland.
	OverseaBuild *string `json:"OverseaBuild,omitempty" xml:"OverseaBuild,omitempty"`
	// The ID of the repository.
	RepoId *string `json:"RepoId,omitempty" xml:"RepoId,omitempty"`
	// The ID of the request.
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
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the synchronization task.
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
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the synchronization task is performed across Alibaba Cloud accounts.
	CrossUser *bool `json:"CrossUser,omitempty" xml:"CrossUser,omitempty"`
	// The source address of the image.
	ImageFrom *GetRepoSyncTaskResponseBodyImageFrom `json:"ImageFrom,omitempty" xml:"ImageFrom,omitempty" type:"Struct"`
	// The destination address of the image.
	ImageTo *GetRepoSyncTaskResponseBodyImageTo `json:"ImageTo,omitempty" xml:"ImageTo,omitempty" type:"Struct"`
	// Indicates whether the request is successful.
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The synchronization tasks for the image layer.
	LayerTasks []*GetRepoSyncTaskResponseBodyLayerTasks `json:"LayerTasks,omitempty" xml:"LayerTasks,omitempty" type:"Repeated"`
	// The synchronization progress. Valid values:
	//
	// *   `0`: The synchronization starts or failed.
	// *   `1`: The synchronization is successful.
	Progress *int64 `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the synchronization task in which multiple images are synchronized at a time.
	SyncBatchTaskId *string `json:"SyncBatchTaskId,omitempty" xml:"SyncBatchTaskId,omitempty"`
	// The ID of the synchronization rule.
	SyncRuleId *string `json:"SyncRuleId,omitempty" xml:"SyncRuleId,omitempty"`
	// The ID of the synchronization task.
	SyncTaskId *string `json:"SyncTaskId,omitempty" xml:"SyncTaskId,omitempty"`
	// Indicates whether transfer acceleration is enabled in the synchronization process.
	SyncTransAccelerate *bool `json:"SyncTransAccelerate,omitempty" xml:"SyncTransAccelerate,omitempty"`
	// The size of the image layer that is synchronized. Unit: bytes.
	SyncedSize *int64 `json:"SyncedSize,omitempty" xml:"SyncedSize,omitempty"`
	// The status of the task. Valid values:
	TaskStatus *string `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	// The policy that is used to trigger the synchronization task.
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
	ImageTag *string `json:"ImageTag,omitempty" xml:"ImageTag,omitempty"`
	// The ID of the instance.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the image repository.
	RepoName *string `json:"RepoName,omitempty" xml:"RepoName,omitempty"`
	// The name of the namespace.
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
	ImageTag *string `json:"ImageTag,omitempty" xml:"ImageTag,omitempty"`
	// The ID of the instance.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the image repository.
	RepoName *string `json:"RepoName,omitempty" xml:"RepoName,omitempty"`
	// The name of the namespace.
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
	ArtifactDigest *string `json:"ArtifactDigest,omitempty" xml:"ArtifactDigest,omitempty"`
	// The digest of the image layer.
	Digest *string `json:"Digest,omitempty" xml:"Digest,omitempty"`
	// The size of synchronized image layers.
	Size *int64 `json:"Size,omitempty" xml:"Size,omitempty"`
	// The ID of the synchronization task for the image layer.
	SyncLayerTaskId *string `json:"SyncLayerTaskId,omitempty" xml:"SyncLayerTaskId,omitempty"`
	// The size of the image layer that is synchronized.
	SyncedSize *int64 `json:"SyncedSize,omitempty" xml:"SyncedSize,omitempty"`
	// The status of the synchronization task. Valid values:
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
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The operation that you want to perform. Set the value to **GetRepoTag**.
	RepoId *string `json:"RepoId,omitempty" xml:"RepoId,omitempty"`
	// The number of milliseconds that have elapsed since the image was created.
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
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The size of the image. Unit: Bytes.
	Digest *string `json:"Digest,omitempty" xml:"Digest,omitempty"`
	// crr-tquyps22md8p****
	ImageCreate *int64  `json:"ImageCreate,omitempty" xml:"ImageCreate,omitempty"`
	ImageId     *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The number of milliseconds that have elapsed since the image was last updated.
	ImageSize *int64 `json:"ImageSize,omitempty" xml:"ImageSize,omitempty"`
	// The ID of the request.
	ImageUpdate *int64 `json:"ImageUpdate,omitempty" xml:"ImageUpdate,omitempty"`
	// The status of the image. Valid values:
	//
	// *   `NORMAL`: The image is normal.
	// *   `DELETING`: The image is being deleted.
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// 1.0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the instance.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The version of the repository.
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

type GetRepoTagLayersRequest struct {
	// The digest of the image.
	Digest *string `json:"Digest,omitempty" xml:"Digest,omitempty"`
	// The ID of the instance.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the image repository.
	RepoId *string `json:"RepoId,omitempty" xml:"RepoId,omitempty"`
	// The tag of the image.
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
}

func (s GetRepoTagLayersRequest) String() string {
	return tea.Prettify(s)
}

func (s GetRepoTagLayersRequest) GoString() string {
	return s.String()
}

func (s *GetRepoTagLayersRequest) SetDigest(v string) *GetRepoTagLayersRequest {
	s.Digest = &v
	return s
}

func (s *GetRepoTagLayersRequest) SetInstanceId(v string) *GetRepoTagLayersRequest {
	s.InstanceId = &v
	return s
}

func (s *GetRepoTagLayersRequest) SetRepoId(v string) *GetRepoTagLayersRequest {
	s.RepoId = &v
	return s
}

func (s *GetRepoTagLayersRequest) SetTag(v string) *GetRepoTagLayersRequest {
	s.Tag = &v
	return s
}

type GetRepoTagLayersResponseBody struct {
	// The return value.
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// *   `true`: The request is successful.
	// *   `false`: The request fails.
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The queried image layers.
	Layers []*GetRepoTagLayersResponseBodyLayers `json:"Layers,omitempty" xml:"Layers,omitempty" type:"Repeated"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetRepoTagLayersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetRepoTagLayersResponseBody) GoString() string {
	return s.String()
}

func (s *GetRepoTagLayersResponseBody) SetCode(v string) *GetRepoTagLayersResponseBody {
	s.Code = &v
	return s
}

func (s *GetRepoTagLayersResponseBody) SetIsSuccess(v bool) *GetRepoTagLayersResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *GetRepoTagLayersResponseBody) SetLayers(v []*GetRepoTagLayersResponseBodyLayers) *GetRepoTagLayersResponseBody {
	s.Layers = v
	return s
}

func (s *GetRepoTagLayersResponseBody) SetRequestId(v string) *GetRepoTagLayersResponseBody {
	s.RequestId = &v
	return s
}

type GetRepoTagLayersResponseBodyLayers struct {
	// The digest of a single image layer.
	BlobDigest *string `json:"BlobDigest,omitempty" xml:"BlobDigest,omitempty"`
	// The size of the image layer.
	BlobSize *int64 `json:"BlobSize,omitempty" xml:"BlobSize,omitempty"`
	// Operation on the image layer.
	LayerCMD *string `json:"LayerCMD,omitempty" xml:"LayerCMD,omitempty"`
	// The sequence number of the layer stack.
	LayerIndex *int32 `json:"LayerIndex,omitempty" xml:"LayerIndex,omitempty"`
	// The command for the image layer.
	LayerInstruction *string `json:"LayerInstruction,omitempty" xml:"LayerInstruction,omitempty"`
}

func (s GetRepoTagLayersResponseBodyLayers) String() string {
	return tea.Prettify(s)
}

func (s GetRepoTagLayersResponseBodyLayers) GoString() string {
	return s.String()
}

func (s *GetRepoTagLayersResponseBodyLayers) SetBlobDigest(v string) *GetRepoTagLayersResponseBodyLayers {
	s.BlobDigest = &v
	return s
}

func (s *GetRepoTagLayersResponseBodyLayers) SetBlobSize(v int64) *GetRepoTagLayersResponseBodyLayers {
	s.BlobSize = &v
	return s
}

func (s *GetRepoTagLayersResponseBodyLayers) SetLayerCMD(v string) *GetRepoTagLayersResponseBodyLayers {
	s.LayerCMD = &v
	return s
}

func (s *GetRepoTagLayersResponseBodyLayers) SetLayerIndex(v int32) *GetRepoTagLayersResponseBodyLayers {
	s.LayerIndex = &v
	return s
}

func (s *GetRepoTagLayersResponseBodyLayers) SetLayerInstruction(v string) *GetRepoTagLayersResponseBodyLayers {
	s.LayerInstruction = &v
	return s
}

type GetRepoTagLayersResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetRepoTagLayersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetRepoTagLayersResponse) String() string {
	return tea.Prettify(s)
}

func (s GetRepoTagLayersResponse) GoString() string {
	return s.String()
}

func (s *GetRepoTagLayersResponse) SetHeaders(v map[string]*string) *GetRepoTagLayersResponse {
	s.Headers = v
	return s
}

func (s *GetRepoTagLayersResponse) SetStatusCode(v int32) *GetRepoTagLayersResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRepoTagLayersResponse) SetBody(v *GetRepoTagLayersResponseBody) *GetRepoTagLayersResponse {
	s.Body = v
	return s
}

type GetRepoTagManifestRequest struct {
	// The ID of the instance.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the repository.
	RepoId *string `json:"RepoId,omitempty" xml:"RepoId,omitempty"`
	// The schema version of the manifest. Valid values: 1 and 2.
	SchemaVersion *int32 `json:"SchemaVersion,omitempty" xml:"SchemaVersion,omitempty"`
	// The tag of the image.
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
}

func (s GetRepoTagManifestRequest) String() string {
	return tea.Prettify(s)
}

func (s GetRepoTagManifestRequest) GoString() string {
	return s.String()
}

func (s *GetRepoTagManifestRequest) SetInstanceId(v string) *GetRepoTagManifestRequest {
	s.InstanceId = &v
	return s
}

func (s *GetRepoTagManifestRequest) SetRepoId(v string) *GetRepoTagManifestRequest {
	s.RepoId = &v
	return s
}

func (s *GetRepoTagManifestRequest) SetSchemaVersion(v int32) *GetRepoTagManifestRequest {
	s.SchemaVersion = &v
	return s
}

func (s *GetRepoTagManifestRequest) SetTag(v string) *GetRepoTagManifestRequest {
	s.Tag = &v
	return s
}

type GetRepoTagManifestResponseBody struct {
	// The return value.
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the request is successful.
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The information about the image manifest.
	Manifest *GetRepoTagManifestResponseBodyManifest `json:"Manifest,omitempty" xml:"Manifest,omitempty" type:"Struct"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetRepoTagManifestResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetRepoTagManifestResponseBody) GoString() string {
	return s.String()
}

func (s *GetRepoTagManifestResponseBody) SetCode(v string) *GetRepoTagManifestResponseBody {
	s.Code = &v
	return s
}

func (s *GetRepoTagManifestResponseBody) SetIsSuccess(v bool) *GetRepoTagManifestResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *GetRepoTagManifestResponseBody) SetManifest(v *GetRepoTagManifestResponseBodyManifest) *GetRepoTagManifestResponseBody {
	s.Manifest = v
	return s
}

func (s *GetRepoTagManifestResponseBody) SetRequestId(v string) *GetRepoTagManifestResponseBody {
	s.RequestId = &v
	return s
}

type GetRepoTagManifestResponseBodyManifest struct {
	// Architecture.
	Architecture *string `json:"Architecture,omitempty" xml:"Architecture,omitempty"`
	// The configuration information.
	Config *GetRepoTagManifestResponseBodyManifestConfig `json:"Config,omitempty" xml:"Config,omitempty" type:"Struct"`
	// The digest of the referenced filesystem image layer.
	FsLayers []*GetRepoTagManifestResponseBodyManifestFsLayers `json:"FsLayers,omitempty" xml:"FsLayers,omitempty" type:"Repeated"`
	// A list of unstructured historical data for V1 compatibility.
	History []*GetRepoTagManifestResponseBodyManifestHistory `json:"History,omitempty" xml:"History,omitempty" type:"Repeated"`
	// The information about image layers.
	Layers []*GetRepoTagManifestResponseBodyManifestLayers `json:"Layers,omitempty" xml:"Layers,omitempty" type:"Repeated"`
	// The type.
	MediaType *string `json:"MediaType,omitempty" xml:"MediaType,omitempty"`
	// The name.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The schema version of the manifest.
	SchemaVersion *int32 `json:"SchemaVersion,omitempty" xml:"SchemaVersion,omitempty"`
	// The information about signatures.
	Signatures []*GetRepoTagManifestResponseBodyManifestSignatures `json:"Signatures,omitempty" xml:"Signatures,omitempty" type:"Repeated"`
	// The tag of the image.
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
}

func (s GetRepoTagManifestResponseBodyManifest) String() string {
	return tea.Prettify(s)
}

func (s GetRepoTagManifestResponseBodyManifest) GoString() string {
	return s.String()
}

func (s *GetRepoTagManifestResponseBodyManifest) SetArchitecture(v string) *GetRepoTagManifestResponseBodyManifest {
	s.Architecture = &v
	return s
}

func (s *GetRepoTagManifestResponseBodyManifest) SetConfig(v *GetRepoTagManifestResponseBodyManifestConfig) *GetRepoTagManifestResponseBodyManifest {
	s.Config = v
	return s
}

func (s *GetRepoTagManifestResponseBodyManifest) SetFsLayers(v []*GetRepoTagManifestResponseBodyManifestFsLayers) *GetRepoTagManifestResponseBodyManifest {
	s.FsLayers = v
	return s
}

func (s *GetRepoTagManifestResponseBodyManifest) SetHistory(v []*GetRepoTagManifestResponseBodyManifestHistory) *GetRepoTagManifestResponseBodyManifest {
	s.History = v
	return s
}

func (s *GetRepoTagManifestResponseBodyManifest) SetLayers(v []*GetRepoTagManifestResponseBodyManifestLayers) *GetRepoTagManifestResponseBodyManifest {
	s.Layers = v
	return s
}

func (s *GetRepoTagManifestResponseBodyManifest) SetMediaType(v string) *GetRepoTagManifestResponseBodyManifest {
	s.MediaType = &v
	return s
}

func (s *GetRepoTagManifestResponseBodyManifest) SetName(v string) *GetRepoTagManifestResponseBodyManifest {
	s.Name = &v
	return s
}

func (s *GetRepoTagManifestResponseBodyManifest) SetSchemaVersion(v int32) *GetRepoTagManifestResponseBodyManifest {
	s.SchemaVersion = &v
	return s
}

func (s *GetRepoTagManifestResponseBodyManifest) SetSignatures(v []*GetRepoTagManifestResponseBodyManifestSignatures) *GetRepoTagManifestResponseBodyManifest {
	s.Signatures = v
	return s
}

func (s *GetRepoTagManifestResponseBodyManifest) SetTag(v string) *GetRepoTagManifestResponseBodyManifest {
	s.Tag = &v
	return s
}

type GetRepoTagManifestResponseBodyManifestConfig struct {
	// The digest of the image.
	Digest *string `json:"Digest,omitempty" xml:"Digest,omitempty"`
	// The MIME type of the configuration file.
	MediaType *string `json:"MediaType,omitempty" xml:"MediaType,omitempty"`
	// Size
	Size *int64 `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s GetRepoTagManifestResponseBodyManifestConfig) String() string {
	return tea.Prettify(s)
}

func (s GetRepoTagManifestResponseBodyManifestConfig) GoString() string {
	return s.String()
}

func (s *GetRepoTagManifestResponseBodyManifestConfig) SetDigest(v string) *GetRepoTagManifestResponseBodyManifestConfig {
	s.Digest = &v
	return s
}

func (s *GetRepoTagManifestResponseBodyManifestConfig) SetMediaType(v string) *GetRepoTagManifestResponseBodyManifestConfig {
	s.MediaType = &v
	return s
}

func (s *GetRepoTagManifestResponseBodyManifestConfig) SetSize(v int64) *GetRepoTagManifestResponseBodyManifestConfig {
	s.Size = &v
	return s
}

type GetRepoTagManifestResponseBodyManifestFsLayers struct {
	// A list of filesystem layer blob sums contained in this image.
	BlobSum *string `json:"BlobSum,omitempty" xml:"BlobSum,omitempty"`
}

func (s GetRepoTagManifestResponseBodyManifestFsLayers) String() string {
	return tea.Prettify(s)
}

func (s GetRepoTagManifestResponseBodyManifestFsLayers) GoString() string {
	return s.String()
}

func (s *GetRepoTagManifestResponseBodyManifestFsLayers) SetBlobSum(v string) *GetRepoTagManifestResponseBodyManifestFsLayers {
	s.BlobSum = &v
	return s
}

type GetRepoTagManifestResponseBodyManifestHistory struct {
	// The raw V1 compatibility information.
	V1Compatibility map[string]interface{} `json:"V1Compatibility,omitempty" xml:"V1Compatibility,omitempty"`
}

func (s GetRepoTagManifestResponseBodyManifestHistory) String() string {
	return tea.Prettify(s)
}

func (s GetRepoTagManifestResponseBodyManifestHistory) GoString() string {
	return s.String()
}

func (s *GetRepoTagManifestResponseBodyManifestHistory) SetV1Compatibility(v map[string]interface{}) *GetRepoTagManifestResponseBodyManifestHistory {
	s.V1Compatibility = v
	return s
}

type GetRepoTagManifestResponseBodyManifestLayers struct {
	// The digest of the image.
	Digest *string `json:"Digest,omitempty" xml:"Digest,omitempty"`
	// The MIME type of the configuration file.
	MediaType *string `json:"MediaType,omitempty" xml:"MediaType,omitempty"`
	// Size.
	Size *int64 `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s GetRepoTagManifestResponseBodyManifestLayers) String() string {
	return tea.Prettify(s)
}

func (s GetRepoTagManifestResponseBodyManifestLayers) GoString() string {
	return s.String()
}

func (s *GetRepoTagManifestResponseBodyManifestLayers) SetDigest(v string) *GetRepoTagManifestResponseBodyManifestLayers {
	s.Digest = &v
	return s
}

func (s *GetRepoTagManifestResponseBodyManifestLayers) SetMediaType(v string) *GetRepoTagManifestResponseBodyManifestLayers {
	s.MediaType = &v
	return s
}

func (s *GetRepoTagManifestResponseBodyManifestLayers) SetSize(v int64) *GetRepoTagManifestResponseBodyManifestLayers {
	s.Size = &v
	return s
}

type GetRepoTagManifestResponseBodyManifestSignatures struct {
	// The header information of the signature.
	Header map[string]interface{} `json:"Header,omitempty" xml:"Header,omitempty"`
	// The signed protected header.
	Protected *string `json:"Protected,omitempty" xml:"Protected,omitempty"`
	// The signature for the image manifest.
	Signature *string `json:"Signature,omitempty" xml:"Signature,omitempty"`
}

func (s GetRepoTagManifestResponseBodyManifestSignatures) String() string {
	return tea.Prettify(s)
}

func (s GetRepoTagManifestResponseBodyManifestSignatures) GoString() string {
	return s.String()
}

func (s *GetRepoTagManifestResponseBodyManifestSignatures) SetHeader(v map[string]interface{}) *GetRepoTagManifestResponseBodyManifestSignatures {
	s.Header = v
	return s
}

func (s *GetRepoTagManifestResponseBodyManifestSignatures) SetProtected(v string) *GetRepoTagManifestResponseBodyManifestSignatures {
	s.Protected = &v
	return s
}

func (s *GetRepoTagManifestResponseBodyManifestSignatures) SetSignature(v string) *GetRepoTagManifestResponseBodyManifestSignatures {
	s.Signature = &v
	return s
}

type GetRepoTagManifestResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetRepoTagManifestResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetRepoTagManifestResponse) String() string {
	return tea.Prettify(s)
}

func (s GetRepoTagManifestResponse) GoString() string {
	return s.String()
}

func (s *GetRepoTagManifestResponse) SetHeaders(v map[string]*string) *GetRepoTagManifestResponse {
	s.Headers = v
	return s
}

func (s *GetRepoTagManifestResponse) SetStatusCode(v int32) *GetRepoTagManifestResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRepoTagManifestResponse) SetBody(v *GetRepoTagManifestResponseBody) *GetRepoTagManifestResponse {
	s.Body = v
	return s
}

type GetRepoTagScanStatusRequest struct {
	Digest     *string `json:"Digest,omitempty" xml:"Digest,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RepoId     *string `json:"RepoId,omitempty" xml:"RepoId,omitempty"`
	ScanTaskId *string `json:"ScanTaskId,omitempty" xml:"ScanTaskId,omitempty"`
	Tag        *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
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

func (s *GetRepoTagScanStatusRequest) SetTag(v string) *GetRepoTagScanStatusRequest {
	s.Tag = &v
	return s
}

type GetRepoTagScanStatusResponseBody struct {
	Code        *string `json:"Code,omitempty" xml:"Code,omitempty"`
	IsSuccess   *bool   `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ScanService *string `json:"ScanService,omitempty" xml:"ScanService,omitempty"`
	Status      *string `json:"Status,omitempty" xml:"Status,omitempty"`
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
	Digest *string `json:"Digest,omitempty" xml:"Digest,omitempty"`
	// The ID of the image repository.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the image tag.
	RepoId *string `json:"RepoId,omitempty" xml:"RepoId,omitempty"`
	// The digest of the image.
	ScanTaskId *string `json:"ScanTaskId,omitempty" xml:"ScanTaskId,omitempty"`
	// The ID of the security scan task.
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
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The number of low-severity vulnerabilities.
	HighSeverity *int32 `json:"HighSeverity,omitempty" xml:"HighSeverity,omitempty"`
	// The number of high-severity vulnerabilities.
	IsSuccess   *bool  `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	LowSeverity *int32 `json:"LowSeverity,omitempty" xml:"LowSeverity,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// *   `true`: The request is successful.
	// *   `false`: The request fails.
	MediumSeverity *int32 `json:"MediumSeverity,omitempty" xml:"MediumSeverity,omitempty"`
	// The total number of vulnerabilities detected on images.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The return value.
	TotalSeverity *int32 `json:"TotalSeverity,omitempty" xml:"TotalSeverity,omitempty"`
	// The ID of the request.
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
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the repository.
	RepoId *string `json:"RepoId,omitempty" xml:"RepoId,omitempty"`
	// The name of the repository.
	RepoName *string `json:"RepoName,omitempty" xml:"RepoName,omitempty"`
	// The name of the namespace to which the repository belongs.
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
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The time when the repository was created.
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The details of the repository.
	Detail *string `json:"Detail,omitempty" xml:"Detail,omitempty"`
	// The ID of the instance.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Indicates whether the request is successful.
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The time when the repository was last modified.
	ModifiedTime *int64 `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// Indicates how the repository was created. Valid values:
	//
	// *   `MANUAL`: The repository was manually created.
	// *   `AUTO`: The repository was automatically created.
	RepoBuildType *string `json:"RepoBuildType,omitempty" xml:"RepoBuildType,omitempty"`
	// The ID of the repository.
	RepoId *string `json:"RepoId,omitempty" xml:"RepoId,omitempty"`
	// The name of the repository.
	RepoName *string `json:"RepoName,omitempty" xml:"RepoName,omitempty"`
	// The name of the namespace.
	RepoNamespaceName *string `json:"RepoNamespaceName,omitempty" xml:"RepoNamespaceName,omitempty"`
	// The status of the repository.
	RepoStatus *string `json:"RepoStatus,omitempty" xml:"RepoStatus,omitempty"`
	// The type of the repository. Valid values:
	//
	// *   `PUBLIC`: public repository.
	// *   `PRIVATE`: private repository.
	RepoType *string `json:"RepoType,omitempty" xml:"RepoType,omitempty"`
	// The ID of the request.
	RequestId       *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The summary of the repository.
	Summary *string `json:"Summary,omitempty" xml:"Summary,omitempty"`
	// Indicates whether the feature of image tag immutability is enabled. Valid values:
	//
	// *   `true`: The feature of image tag immutability is enabled.
	// *   `false`: The feature of image tag immutability is disabled.
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
	BuildTaskId *string `json:"BuildTaskId,omitempty" xml:"BuildTaskId,omitempty"`
	// The ID of the Container Registry instance.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The number of the page to return.
	Page *int32 `json:"Page,omitempty" xml:"Page,omitempty"`
	// The number of entries to return on each page.
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
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the API call is successful.
	//
	// *   `true`: successful
	// *   `false`: failed
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of log entries.
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
	LineNumber *int32 `json:"LineNumber,omitempty" xml:"LineNumber,omitempty"`
	// The content of the log entry.
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
	EnableDeleteTag *bool   `json:"EnableDeleteTag,omitempty" xml:"EnableDeleteTag,omitempty"`
	InstanceId      *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	PageNo          *int32  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize        *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
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
	Code       *string                                       `json:"Code,omitempty" xml:"Code,omitempty"`
	IsSuccess  *bool                                         `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	PageNo     *int32                                        `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize   *int32                                        `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Rules      []*ListArtifactLifecycleRuleResponseBodyRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
	TotalCount *int32                                        `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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
	Auto              *bool   `json:"Auto,omitempty" xml:"Auto,omitempty"`
	CreateTime        *int64  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	EnableDeleteTag   *bool   `json:"EnableDeleteTag,omitempty" xml:"EnableDeleteTag,omitempty"`
	InstanceId        *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	ModifiedTime      *int64  `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	NamespaceName     *string `json:"NamespaceName,omitempty" xml:"NamespaceName,omitempty"`
	NextTime          *int64  `json:"NextTime,omitempty" xml:"NextTime,omitempty"`
	RepoName          *string `json:"RepoName,omitempty" xml:"RepoName,omitempty"`
	RetentionTagCount *int64  `json:"RetentionTagCount,omitempty" xml:"RetentionTagCount,omitempty"`
	RuleId            *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	ScheduleTime      *string `json:"ScheduleTime,omitempty" xml:"ScheduleTime,omitempty"`
	Scope             *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
	TagRegexp         *string `json:"TagRegexp,omitempty" xml:"TagRegexp,omitempty"`
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

type ListChainRequest struct {
	// The ID of the instance.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The number of the page to return.
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries to return on each page.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The name of the image repository.
	RepoName *string `json:"RepoName,omitempty" xml:"RepoName,omitempty"`
	// The name of the namespace.
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
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the request is successful.
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The page number of the returned page.
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries returned per page.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of delivery chains.
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
	ChainId *string `json:"ChainId,omitempty" xml:"ChainId,omitempty"`
	// The time when the delivery chain was created.
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the delivery chain.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the instance.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The time when the delivery chain was last modified.
	ModifiedTime *int64 `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// The name of the delivery chain.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Repositories to which the delivery chain does not apply.
	ScopeExclude []*string `json:"ScopeExclude,omitempty" xml:"ScopeExclude,omitempty" type:"Repeated"`
	// The ID of the delivery chain scope.
	ScopeId *string `json:"ScopeId,omitempty" xml:"ScopeId,omitempty"`
	// The type of the delivery chain scope.
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
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The time when the delivery chain started.
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The name of the image repository.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The time when the delivery chain is completed.
	RepoName *string `json:"RepoName,omitempty" xml:"RepoName,omitempty"`
	// The name of the delivery chain.
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
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The page number of the page to return.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The execution record of the delivery chain.
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// 30
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// Indicates whether the operation is successful.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the Container Registry instance.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The name of the repository.
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
	ChainInstanceId *string `json:"ChainInstanceId,omitempty" xml:"ChainInstanceId,omitempty"`
	// The ID of the Container Registry instance.
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The ID of the delivery chain.
	RepoName *string `json:"RepoName,omitempty" xml:"RepoName,omitempty"`
	// The execution result of the delivery chain. Valid values:
	//
	// *   `SUCCESS`
	// *   `FAILED`
	// *   `CANCELED`
	// *   `DENIED`
	RepoNamespaceName *string `json:"RepoNamespaceName,omitempty" xml:"RepoNamespaceName,omitempty"`
	// The list of the execution records of delivery chains.
	Result *string `json:"Result,omitempty" xml:"Result,omitempty"`
	// test-repo
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The status of the delivery chain. Valid values:
	//
	// *   `RUNNING`
	// *   `COMPLETE`
	// *   `CANCELING`
	// *   `CANCELED`
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
	ChainId *string `json:"ChainId,omitempty" xml:"ChainId,omitempty"`
	// The number of entries returned on each page.
	ChainName *string `json:"ChainName,omitempty" xml:"ChainName,omitempty"`
	// The ID of the request.
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
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the namespace.
	NamespaceName *string `json:"NamespaceName,omitempty" xml:"NamespaceName,omitempty"`
	// The status of the namespace. Valid values:
	//
	// *   `NORMAL`: The namespace is normal.
	// *   `DELETING`: The namespace is being deleted.
	NamespaceStatus *string `json:"NamespaceStatus,omitempty" xml:"NamespaceStatus,omitempty"`
	// The number of the page to return.
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries to return on each page.
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
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the request is successful.
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The namespaces.
	Namespaces []*ListChartNamespaceResponseBodyNamespaces `json:"Namespaces,omitempty" xml:"Namespaces,omitempty" type:"Repeated"`
	// The page number of the returned page.
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries returned per page.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of returned entries.
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
	AutoCreateRepo *bool `json:"AutoCreateRepo,omitempty" xml:"AutoCreateRepo,omitempty"`
	// The default repository type. Valid values:
	//
	// *   `PUBLIC`: a public repository
	// *   `PRIVATE`: a private repository
	DefaultRepoType *string `json:"DefaultRepoType,omitempty" xml:"DefaultRepoType,omitempty"`
	// The ID of the instance.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the namespace.
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	// The name of the namespace.
	NamespaceName *string `json:"NamespaceName,omitempty" xml:"NamespaceName,omitempty"`
	// The status of the namespace. Valid values:
	//
	// *   `NORMAL`: The namespace is normal.
	// *   `DELETING`: The namespace is being deleted.
	NamespaceStatus *string `json:"NamespaceStatus,omitempty" xml:"NamespaceStatus,omitempty"`
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
	Chart *string `json:"Chart,omitempty" xml:"Chart,omitempty"`
	// The ID of the instance.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The number of the page to return.
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries to return on each page.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The name of the repository.
	RepoName *string `json:"RepoName,omitempty" xml:"RepoName,omitempty"`
	// The name of the namespace.
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
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the request is successful.
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The page number of the returned page.
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries returned per page.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of returned entries.
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
	Chart *string `json:"Chart,omitempty" xml:"Chart,omitempty"`
	// The ID of the instance.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The time when the chart was last modified.
	ModifiedTime *int64 `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// The version number of the chart.
	Release *string `json:"Release,omitempty" xml:"Release,omitempty"`
	// The ID of the chart repository.
	RepoId *string `json:"RepoId,omitempty" xml:"RepoId,omitempty"`
	// The size of the chart of the version. Unit: bytes.
	Size *string `json:"Size,omitempty" xml:"Size,omitempty"`
	// The status of the chart.
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
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The number of the page to return.
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries to return on each page.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The name of the repository.
	RepoName *string `json:"RepoName,omitempty" xml:"RepoName,omitempty"`
	// The name of the namespace to which the repository belongs.
	RepoNamespaceName *string `json:"RepoNamespaceName,omitempty" xml:"RepoNamespaceName,omitempty"`
	// The status of the chart repositories that you want to query. Valid values:
	//
	// *   `ALL`: query repositories of all status.
	// *   `NORMAL`: query normal repositories.
	// *   `DELETING`: query repositories that are being deleted.
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
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the request is successful.
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The page number of the returned page.
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries returned per page.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The queried repositories.
	Repositories []*ListChartRepositoryResponseBodyRepositories `json:"Repositories,omitempty" xml:"Repositories,omitempty" type:"Repeated"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of returned entries.
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
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The ID of the instance.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The time when the repository was last modified.
	ModifiedTime *int64 `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// The ID of the repository.
	RepoId *string `json:"RepoId,omitempty" xml:"RepoId,omitempty"`
	// The name of the repository.
	RepoName *string `json:"RepoName,omitempty" xml:"RepoName,omitempty"`
	// The name of the namespace to which the repository belongs.
	RepoNamespaceName *string `json:"RepoNamespaceName,omitempty" xml:"RepoNamespaceName,omitempty"`
	// The status of the repository. Valid values:
	//
	// *   `NORMAL`: The repository is normal.
	// *   `DELETING`: The repository is being deleted.
	RepoStatus *string `json:"RepoStatus,omitempty" xml:"RepoStatus,omitempty"`
	// The type of the repository. Valid values:
	//
	// *   `PRIVATE`: a private repository
	// *   `PUBLIC`: a public repository
	RepoType *string `json:"RepoType,omitempty" xml:"RepoType,omitempty"`
	// The ID of the resource group to which the repository belongs.
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The summary about the repository.
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
	// *   `cr:Artifact:DeliveryChainCompleted`: The delivery chain is processed.
	// *   `cr:Artifact:SynchronizationCompleted`: The image is replicated.
	// *   `cr:Artifact:BuildCompleted`: The image is built.
	// *   `cr:Artifact:ScanCompleted`: The image is scanned.
	// *   `cr:Artifact:SigningCompleted`: The image is signed.
	EventType *string `json:"EventType,omitempty" xml:"EventType,omitempty"`
	// The ID of the instance.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The number of the page to return.
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries to return on each page.
	PageSize          *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RepoName          *string `json:"RepoName,omitempty" xml:"RepoName,omitempty"`
	RepoNamespaceName *string `json:"RepoNamespaceName,omitempty" xml:"RepoNamespaceName,omitempty"`
	// The ID of the event notification rule.
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
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the request is successful.
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The page number of the returned page.
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries returned per page.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The list of historical events.
	Records []*ListEventCenterRecordResponseBodyRecords `json:"Records,omitempty" xml:"Records,omitempty" type:"Repeated"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total entries of historical events.
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
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The event notification channel.
	EventChannel *string `json:"EventChannel,omitempty" xml:"EventChannel,omitempty"`
	// The ID of the event notification.
	EventNotifyId *string `json:"EventNotifyId,omitempty" xml:"EventNotifyId,omitempty"`
	// The notification method. Valid values:
	//
	// *   `http`: The notification is sent over HTTP.
	// *   `https`: The notification is sent over HTTPS.
	// *   `dingding`: The notification is sent by using DingTalk.
	EventNotifyMethod *string `json:"EventNotifyMethod,omitempty" xml:"EventNotifyMethod,omitempty"`
	// The type of the event.
	EventType *string `json:"EventType,omitempty" xml:"EventType,omitempty"`
	// The ID of the instance.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The namespace.
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The ID of the event record.
	RecordId *string `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
	// The name of the image repository.
	RepoName *string `json:"RepoName,omitempty" xml:"RepoName,omitempty"`
	// The ID of the event notification rule.
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// The name of the event notification rule.
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The tags.
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	// The time when the event was last updated.
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
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the request is successful.
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The list of names of event notification rules.
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
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// The name of the event notification rule.
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
	// The instance name.
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The status of the instance. Valid values:
	//
	// *   `PENDING`: The instance is being initialized.
	// *   `INIT_ERROR`: The initialization of the instance fails.
	// *   `STARTING`: The instance is being started.
	// *   `RUNNING`: The instance is running.
	// *   `STOPPING`: The instance is being stopped.
	// *   `STOPPED`: The instance is stopped.
	// *   `DELETING`: The instance is being deleted.
	// *   `DELETED`: The instance is deleted.
	InstanceStatus *string `json:"InstanceStatus,omitempty" xml:"InstanceStatus,omitempty"`
	// The page number.
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the resource group to which the instance belongs.
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
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The queried instances.
	Instances []*ListInstanceResponseBodyInstances `json:"Instances,omitempty" xml:"Instances,omitempty" type:"Repeated"`
	// Indicates whether the request is successful.
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The page number. Default value: 1.
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page. Default value: 30.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
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
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The instance ID.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The issue occurs on the instance.
	InstanceIssue *string `json:"InstanceIssue,omitempty" xml:"InstanceIssue,omitempty"`
	// The name of the instance.
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The edition of the Container Registry Enterprise Edition instance.
	InstanceSpecification *string `json:"InstanceSpecification,omitempty" xml:"InstanceSpecification,omitempty"`
	// The status of the instance.
	InstanceStatus *string `json:"InstanceStatus,omitempty" xml:"InstanceStatus,omitempty"`
	// The time when the instance was last modified.
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// The region ID of the instance.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the instance belongs.
	ResourceGroupId *string                                  `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Tags            []*ListInstanceResponseBodyInstancesTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
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
	TagKey   *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
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
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the module that you want to access. Valid values:
	//
	// *   `Registry`: the image repository.
	// *   `Chart`: a Helm chart.
	ModuleName *string `json:"ModuleName,omitempty" xml:"ModuleName,omitempty"`
	Summary    *bool   `json:"Summary,omitempty" xml:"Summary,omitempty"`
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
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The endpoints of the instance.
	Endpoints []*ListInstanceEndpointResponseBodyEndpoints `json:"Endpoints,omitempty" xml:"Endpoints,omitempty" type:"Repeated"`
	// Indicates whether the request is successful.
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The ID of the request.
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
	// Indicates whether the access control list (ACL) feature is enabled.
	AclEnable *bool `json:"AclEnable,omitempty" xml:"AclEnable,omitempty"`
	// The ACL configured for the instance.
	AclEntries []*ListInstanceEndpointResponseBodyEndpointsAclEntries `json:"AclEntries,omitempty" xml:"AclEntries,omitempty" type:"Repeated"`
	// Domain names.
	Domains []*ListInstanceEndpointResponseBodyEndpointsDomains `json:"Domains,omitempty" xml:"Domains,omitempty" type:"Repeated"`
	// Indicates whether the ACL feature is enabled.
	Enable *bool `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// The type of the endpoint.
	EndpointType *string `json:"EndpointType,omitempty" xml:"EndpointType,omitempty"`
	// The virtual private clouds (VPCs) that are associated with the instance.
	LinkedVpcs []*ListInstanceEndpointResponseBodyEndpointsLinkedVpcs `json:"LinkedVpcs,omitempty" xml:"LinkedVpcs,omitempty" type:"Repeated"`
	// The status of the instance.
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
	// Details about the ACL.
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
	// The domain name.
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// Type
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
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// *   `true`: The request is successful.
	// *   `false`: The request fails.
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The list of regions.
	Regions []*ListInstanceRegionResponseBodyRegions `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Repeated"`
	// The ID of the request.
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
	LocalName *string `json:"LocalName,omitempty" xml:"LocalName,omitempty"`
	// The ID of the region.
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
	// The number of the page to return.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The number of entries returned per page.
	NamespaceName *string `json:"NamespaceName,omitempty" xml:"NamespaceName,omitempty"`
	// The ID of the namespace.
	NamespaceStatus *string `json:"NamespaceStatus,omitempty" xml:"NamespaceStatus,omitempty"`
	// The list of namespaces.
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The ID of the request.
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
	Code       *string                                `json:"Code,omitempty" xml:"Code,omitempty"`
	IsSuccess  *bool                                  `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	Namespaces []*ListNamespaceResponseBodyNamespaces `json:"Namespaces,omitempty" xml:"Namespaces,omitempty" type:"Repeated"`
	PageNo     *int32                                 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize   *int32                                 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *string                                `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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
	AutoCreateRepo  *bool   `json:"AutoCreateRepo,omitempty" xml:"AutoCreateRepo,omitempty"`
	DefaultRepoType *string `json:"DefaultRepoType,omitempty" xml:"DefaultRepoType,omitempty"`
	InstanceId      *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	NamespaceId     *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	NamespaceName   *string `json:"NamespaceName,omitempty" xml:"NamespaceName,omitempty"`
	NamespaceStatus *string `json:"NamespaceStatus,omitempty" xml:"NamespaceStatus,omitempty"`
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
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The number of the page to return.
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries to return on each page.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the repository.
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
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// *   `true`: The request is successful.
	// *   `false`: The request fails.
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The page number of the returned page.
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries returned per page.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of returned entries.
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
	BuildRecordId *string `json:"BuildRecordId,omitempty" xml:"BuildRecordId,omitempty"`
	// The status of the image building.
	BuildStatus *string `json:"BuildStatus,omitempty" xml:"BuildStatus,omitempty"`
	// The time when the image building ended.
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The information about the image.
	Image *ListRepoBuildRecordResponseBodyBuildRecordsImage `json:"Image,omitempty" xml:"Image,omitempty" type:"Struct"`
	// The time when the image building started.
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
	ImageTag *string `json:"ImageTag,omitempty" xml:"ImageTag,omitempty"`
	// The ID of the repository.
	RepoId *string `json:"RepoId,omitempty" xml:"RepoId,omitempty"`
	// The name of the repository.
	RepoName *string `json:"RepoName,omitempty" xml:"RepoName,omitempty"`
	// The name of the namespace to which the repository belongs.
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
	BuildRecordId *string `json:"BuildRecordId,omitempty" xml:"BuildRecordId,omitempty"`
	// The ID of the instance.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The offset of log lines.
	Offset *int32 `json:"Offset,omitempty" xml:"Offset,omitempty"`
	// The ID of the image repository.
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
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the request is successful.
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The page number of the returned page.
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries returned per page.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of returned entries.
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
	BuildStage *string `json:"BuildStage,omitempty" xml:"BuildStage,omitempty"`
	// The line number of the log entry.
	LineNumber *int32 `json:"LineNumber,omitempty" xml:"LineNumber,omitempty"`
	// The content of the log.
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
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The number of the page to return.
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries to return on each page.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the image repository.
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
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// *   `true`: The request is successful.
	// *   `false`: The request fails.
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The page number of the returned page.
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries returned per page.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of returned entries.
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
	BuildRuleId *string `json:"BuildRuleId,omitempty" xml:"BuildRuleId,omitempty"`
	// The directory of the Dockerfile.
	DockerfileLocation *string `json:"DockerfileLocation,omitempty" xml:"DockerfileLocation,omitempty"`
	// The name of the Dockerfile.
	DockerfileName *string `json:"DockerfileName,omitempty" xml:"DockerfileName,omitempty"`
	// The tag of the image.
	ImageTag  *string   `json:"ImageTag,omitempty" xml:"ImageTag,omitempty"`
	Platforms []*string `json:"Platforms,omitempty" xml:"Platforms,omitempty" type:"Repeated"`
	// The name of the push that triggers the building rule.
	PushName *string `json:"PushName,omitempty" xml:"PushName,omitempty"`
	// The type of the push that triggers the image building rule. Valid values:
	//
	// *   GIT_BRANCH: branch push
	// *   GIT_TAG: tag push
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
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the namespace.
	NamespaceName *string `json:"NamespaceName,omitempty" xml:"NamespaceName,omitempty"`
	// The number of the page to return.
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries to return on each page.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The name of the image repository.
	RepoName *string `json:"RepoName,omitempty" xml:"RepoName,omitempty"`
	// The ID of the destination instance.
	TargetInstanceId *string `json:"TargetInstanceId,omitempty" xml:"TargetInstanceId,omitempty"`
	// The region ID of the destination instance.
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
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the request is successful.
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The page number of the returned page.
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries returned per page.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The synchronization rules.
	SyncRules []*ListRepoSyncRuleResponseBodySyncRules `json:"SyncRules,omitempty" xml:"SyncRules,omitempty" type:"Repeated"`
	// The total number of returned entries.
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
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// Indicates whether images are synchronized across Alibaba Cloud accounts. Valid values:
	//
	// *   `true`: Images are synchronized across Alibaba Cloud accounts.
	// *   `false`: Images are synchronized within the same Alibaba Cloud account.
	//
	// Default value: `false`
	CrossUser *bool `json:"CrossUser,omitempty" xml:"CrossUser,omitempty"`
	// The ID of the source instance.
	LocalInstanceId *string `json:"LocalInstanceId,omitempty" xml:"LocalInstanceId,omitempty"`
	// The namespace name of the source instance.
	LocalNamespaceName *string `json:"LocalNamespaceName,omitempty" xml:"LocalNamespaceName,omitempty"`
	// The region ID of the source instance.
	LocalRegionId *string `json:"LocalRegionId,omitempty" xml:"LocalRegionId,omitempty"`
	// The image repository name of the source instance.
	LocalRepoName *string `json:"LocalRepoName,omitempty" xml:"LocalRepoName,omitempty"`
	// The time when the synchronization rule was last modified.
	ModifiedTime *int64 `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// The synchronization direction. Valid values:
	//
	// *   `FROM`: Images are synchronized from the source instance to the destination instance.
	// *   `TO`: Images are synchronized from the destination instance to the source instance.
	SyncDirection *string `json:"SyncDirection,omitempty" xml:"SyncDirection,omitempty"`
	// The ID of the synchronization rule.
	SyncRuleId *string `json:"SyncRuleId,omitempty" xml:"SyncRuleId,omitempty"`
	// The name of the synchronization rule.
	SyncRuleName *string `json:"SyncRuleName,omitempty" xml:"SyncRuleName,omitempty"`
	// The synchronization scope. Valid values:
	//
	// *   `NAMESPACE`: synchronizes the image tags in a namespace that meet the synchronization rule.
	// *   `REPO`: synchronizes the image tags in an image repository that meet the synchronization rule.
	SyncScope *string `json:"SyncScope,omitempty" xml:"SyncScope,omitempty"`
	// The policy that is applied to trigger the synchronization rule. Valid values:
	//
	// *   `INITIATIVE`: The synchronization rule is positively triggered.
	// *   `PASSIVE`: The synchronization rule is passively triggered.
	SyncTrigger *string `json:"SyncTrigger,omitempty" xml:"SyncTrigger,omitempty"`
	// The regular expression that is used to filter image tags.
	TagFilter *string `json:"TagFilter,omitempty" xml:"TagFilter,omitempty"`
	// The ID of the destination instance.
	TargetInstanceId *string `json:"TargetInstanceId,omitempty" xml:"TargetInstanceId,omitempty"`
	// The namespace name of the destination instance.
	TargetNamespaceName *string `json:"TargetNamespaceName,omitempty" xml:"TargetNamespaceName,omitempty"`
	// The region ID of the destination instance.
	TargetRegionId *string `json:"TargetRegionId,omitempty" xml:"TargetRegionId,omitempty"`
	// The image repository name of the destination instance.
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
	InstanceId        *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	PageNo            *int32  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize          *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RepoName          *string `json:"RepoName,omitempty" xml:"RepoName,omitempty"`
	RepoNamespaceName *string `json:"RepoNamespaceName,omitempty" xml:"RepoNamespaceName,omitempty"`
	SyncRecordId      *string `json:"SyncRecordId,omitempty" xml:"SyncRecordId,omitempty"`
	Tag               *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
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
	Code       *string                                  `json:"Code,omitempty" xml:"Code,omitempty"`
	IsSuccess  *bool                                    `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	PageNo     *int32                                   `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize   *int32                                   `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SyncTasks  []*ListRepoSyncTaskResponseBodySyncTasks `json:"SyncTasks,omitempty" xml:"SyncTasks,omitempty" type:"Repeated"`
	TotalCount *string                                  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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
	CreateTime          *int64                                          `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	CrossUser           *bool                                           `json:"CrossUser,omitempty" xml:"CrossUser,omitempty"`
	CustomLink          *bool                                           `json:"CustomLink,omitempty" xml:"CustomLink,omitempty"`
	ImageFrom           *ListRepoSyncTaskResponseBodySyncTasksImageFrom `json:"ImageFrom,omitempty" xml:"ImageFrom,omitempty" type:"Struct"`
	ImageTo             *ListRepoSyncTaskResponseBodySyncTasksImageTo   `json:"ImageTo,omitempty" xml:"ImageTo,omitempty" type:"Struct"`
	ModifedTime         *int64                                          `json:"ModifedTime,omitempty" xml:"ModifedTime,omitempty"`
	SyncBatchTaskId     *string                                         `json:"SyncBatchTaskId,omitempty" xml:"SyncBatchTaskId,omitempty"`
	SyncRuleId          *string                                         `json:"SyncRuleId,omitempty" xml:"SyncRuleId,omitempty"`
	SyncTaskId          *string                                         `json:"SyncTaskId,omitempty" xml:"SyncTaskId,omitempty"`
	SyncTransAccelerate *bool                                           `json:"SyncTransAccelerate,omitempty" xml:"SyncTransAccelerate,omitempty"`
	TaskStatus          *string                                         `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	TaskTrigger         *string                                         `json:"TaskTrigger,omitempty" xml:"TaskTrigger,omitempty"`
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

func (s *ListRepoSyncTaskResponseBodySyncTasks) SetTaskStatus(v string) *ListRepoSyncTaskResponseBodySyncTasks {
	s.TaskStatus = &v
	return s
}

func (s *ListRepoSyncTaskResponseBodySyncTasks) SetTaskTrigger(v string) *ListRepoSyncTaskResponseBodySyncTasks {
	s.TaskTrigger = &v
	return s
}

type ListRepoSyncTaskResponseBodySyncTasksImageFrom struct {
	ImageTag          *string `json:"ImageTag,omitempty" xml:"ImageTag,omitempty"`
	InstanceId        *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId          *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RepoName          *string `json:"RepoName,omitempty" xml:"RepoName,omitempty"`
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
	ImageTag          *string `json:"ImageTag,omitempty" xml:"ImageTag,omitempty"`
	InstanceId        *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId          *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RepoName          *string `json:"RepoName,omitempty" xml:"RepoName,omitempty"`
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
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The number of the page to return.
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries to return on each page.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the repository.
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
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The images.
	Images []*ListRepoTagResponseBodyImages `json:"Images,omitempty" xml:"Images,omitempty" type:"Repeated"`
	// Indicates whether the request is successful.
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The page number of the returned page.
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries returned per page.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of returned entries.
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
	Digest *string `json:"Digest,omitempty" xml:"Digest,omitempty"`
	// The time when the image was created.
	ImageCreate *string `json:"ImageCreate,omitempty" xml:"ImageCreate,omitempty"`
	// The ID of the image.
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The size of the image.
	ImageSize *int64 `json:"ImageSize,omitempty" xml:"ImageSize,omitempty"`
	// The time when the image was last updated.
	ImageUpdate *string `json:"ImageUpdate,omitempty" xml:"ImageUpdate,omitempty"`
	// The status of the image.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tag of the image.
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
	Digest *string `json:"Digest,omitempty" xml:"Digest,omitempty"`
	// The parameter whose value that you want to query. Fox example, if the value is `FixCmd`, only the `FixCmd` parameter is returned.
	FilterValue *string `json:"FilterValue,omitempty" xml:"FilterValue,omitempty"`
	// The ID of the instance.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The number of the page to return.
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries to return on each page.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the image repository.
	RepoId *string `json:"RepoId,omitempty" xml:"RepoId,omitempty"`
	// The ID of the security scan task.
	ScanTaskId *string `json:"ScanTaskId,omitempty" xml:"ScanTaskId,omitempty"`
	// The type of the vulnerability. Valid values:
	//
	// *   `cve`: image system vulnerability
	// *   `sca`: image application vulnerability
	ScanType *string `json:"ScanType,omitempty" xml:"ScanType,omitempty"`
	// The severity of the vulnerability. Valid values:
	//
	// *   `High`
	// *   `Medium`
	// *   `Low`
	// *   `Unknown`
	Severity *string `json:"Severity,omitempty" xml:"Severity,omitempty"`
	// The name of the image tag.
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	// The keyword for fuzzy search used in scanning. The value can be a CVE name.
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
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// *   `true`: The request is successful.
	// *   `false`: The request failed.
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The number of the returned page.
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries returned per page.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of vulnerabilities detected on images.
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
	AddedBy *string `json:"AddedBy,omitempty" xml:"AddedBy,omitempty"`
	// The name of the vulnerability.
	AliasName *string `json:"AliasName,omitempty" xml:"AliasName,omitempty"`
	// The URL of the vulnerability.
	CveLink *string `json:"CveLink,omitempty" xml:"CveLink,omitempty"`
	// The directory of the vulnerability.
	CveLocation *string `json:"CveLocation,omitempty" xml:"CveLocation,omitempty"`
	// The name of the vulnerability.
	CveName *string `json:"CveName,omitempty" xml:"CveName,omitempty"`
	// The description of the vulnerability.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The cause of the vulnerability.
	Feature *string `json:"Feature,omitempty" xml:"Feature,omitempty"`
	// The command used to fix the vulnerability.
	FixCmd *string `json:"FixCmd,omitempty" xml:"FixCmd,omitempty"`
	// The type of the vulnerability. Valid values:
	//
	// *   `cve`: image system vulnerability
	// *   `sca`: image application vulnerability
	ScanType *string `json:"ScanType,omitempty" xml:"ScanType,omitempty"`
	// The severity of the vulnerability.
	Severity *string `json:"Severity,omitempty" xml:"Severity,omitempty"`
	// The version of the vulnerability.
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
	// The version where the vulnerability was fixed.
	VersionFixed *string `json:"VersionFixed,omitempty" xml:"VersionFixed,omitempty"`
	// The format of the vulnerability.
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
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the repository.
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
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// *   `true`: The request is successful.
	// *   `false`: The request fails.
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The ID of the request.
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
	// *   `BUILD_SUCCESS`: The trigger is activated when an image building task is successful.
	// *   `BUILD_Fail`: The trigger is activated when an image building task fails.
	RepoEvent *string `json:"RepoEvent,omitempty" xml:"RepoEvent,omitempty"`
	// The ID of the trigger.
	TriggerId *string `json:"TriggerId,omitempty" xml:"TriggerId,omitempty"`
	// The name of the trigger.
	TriggerName *string `json:"TriggerName,omitempty" xml:"TriggerName,omitempty"`
	// The image tag based on which the trigger is set.
	TriggerTag *string `json:"TriggerTag,omitempty" xml:"TriggerTag,omitempty"`
	// The type of the trigger. Valid values:
	//
	// *   `ALL`: a trigger that supports both tags and regular expressions.
	// *   `TAG_LISTTAG`: a tag-based trigger.
	// *   `TAG_REG_EXP`: a regular expression-based trigger.
	TriggerType *string `json:"TriggerType,omitempty" xml:"TriggerType,omitempty"`
	// The URL of the trigger.
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
	InstanceId        *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	PageNo            *int32  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize          *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RepoName          *string `json:"RepoName,omitempty" xml:"RepoName,omitempty"`
	RepoNamespaceName *string `json:"RepoNamespaceName,omitempty" xml:"RepoNamespaceName,omitempty"`
	RepoStatus        *string `json:"RepoStatus,omitempty" xml:"RepoStatus,omitempty"`
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
	Code         *string                                   `json:"Code,omitempty" xml:"Code,omitempty"`
	IsSuccess    *bool                                     `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	PageNo       *int32                                    `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize     *int32                                    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Repositories []*ListRepositoryResponseBodyRepositories `json:"Repositories,omitempty" xml:"Repositories,omitempty" type:"Repeated"`
	RequestId    *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount   *string                                   `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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
	CreateTime        *int64  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	InstanceId        *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	ModifiedTime      *int64  `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	RepoBuildType     *string `json:"RepoBuildType,omitempty" xml:"RepoBuildType,omitempty"`
	RepoId            *string `json:"RepoId,omitempty" xml:"RepoId,omitempty"`
	RepoName          *string `json:"RepoName,omitempty" xml:"RepoName,omitempty"`
	RepoNamespaceName *string `json:"RepoNamespaceName,omitempty" xml:"RepoNamespaceName,omitempty"`
	RepoStatus        *string `json:"RepoStatus,omitempty" xml:"RepoStatus,omitempty"`
	RepoType          *string `json:"RepoType,omitempty" xml:"RepoType,omitempty"`
	ResourceGroupId   *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Summary           *string `json:"Summary,omitempty" xml:"Summary,omitempty"`
	TagImmutability   *bool   `json:"TagImmutability,omitempty" xml:"TagImmutability,omitempty"`
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
	Digest     *string `json:"Digest,omitempty" xml:"Digest,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Level      *string `json:"Level,omitempty" xml:"Level,omitempty"`
	PageNo     *int32  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RepoId     *string `json:"RepoId,omitempty" xml:"RepoId,omitempty"`
	ScanTaskId *string `json:"ScanTaskId,omitempty" xml:"ScanTaskId,omitempty"`
	Tag        *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
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
	Code      *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	IsSuccess *bool  `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	PageNo    *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize  *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Id of the request
	RequestId     *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ScanBaselines []*ListScanBaselineByTaskResponseBodyScanBaselines `json:"ScanBaselines,omitempty" xml:"ScanBaselines,omitempty" type:"Repeated"`
	TotalCount    *int32                                             `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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
	BaselineClassAlias        *string `json:"BaselineClassAlias,omitempty" xml:"BaselineClassAlias,omitempty"`
	BaselineDetailAdvice      *string `json:"BaselineDetailAdvice,omitempty" xml:"BaselineDetailAdvice,omitempty"`
	BaselineDetailDescription *string `json:"BaselineDetailDescription,omitempty" xml:"BaselineDetailDescription,omitempty"`
	BaselineDetailPrompt      *string `json:"BaselineDetailPrompt,omitempty" xml:"BaselineDetailPrompt,omitempty"`
	BaselineItemCount         *int32  `json:"BaselineItemCount,omitempty" xml:"BaselineItemCount,omitempty"`
	BaselineNameAlias         *string `json:"BaselineNameAlias,omitempty" xml:"BaselineNameAlias,omitempty"`
	BaselineNameKey           *string `json:"BaselineNameKey,omitempty" xml:"BaselineNameKey,omitempty"`
	BaselineNameLevel         *string `json:"BaselineNameLevel,omitempty" xml:"BaselineNameLevel,omitempty"`
	CreateTime                *int64  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	FirstScanTime             *int64  `json:"FirstScanTime,omitempty" xml:"FirstScanTime,omitempty"`
	HighRiskItemCount         *int32  `json:"HighRiskItemCount,omitempty" xml:"HighRiskItemCount,omitempty"`
	LowRiskItemCount          *int32  `json:"LowRiskItemCount,omitempty" xml:"LowRiskItemCount,omitempty"`
	MiddleRiskItemCount       *int32  `json:"MiddleRiskItemCount,omitempty" xml:"MiddleRiskItemCount,omitempty"`
	ScanTaskId                *string `json:"ScanTaskId,omitempty" xml:"ScanTaskId,omitempty"`
	UpdateTime                *int64  `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
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
	Digest     *string `json:"Digest,omitempty" xml:"Digest,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Level      *string `json:"Level,omitempty" xml:"Level,omitempty"`
	PageNo     *int32  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RepoId     *string `json:"RepoId,omitempty" xml:"RepoId,omitempty"`
	ScanTaskId *string `json:"ScanTaskId,omitempty" xml:"ScanTaskId,omitempty"`
	Tag        *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
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
	Code               *int32                                                       `json:"Code,omitempty" xml:"Code,omitempty"`
	IsSuccess          *bool                                                        `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	PageNo             *int32                                                       `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize           *int32                                                       `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId          *string                                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ScanMaliciousFiles []*ListScanMaliciousFileByTaskResponseBodyScanMaliciousFiles `json:"ScanMaliciousFiles,omitempty" xml:"ScanMaliciousFiles,omitempty" type:"Repeated"`
	TotalCount         *int32                                                       `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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
	CreateTime    *int64  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	FilePath      *string `json:"FilePath,omitempty" xml:"FilePath,omitempty"`
	FirstScanTime *int64  `json:"FirstScanTime,omitempty" xml:"FirstScanTime,omitempty"`
	Level         *string `json:"Level,omitempty" xml:"Level,omitempty"`
	MaliciousMd5  *string `json:"MaliciousMd5,omitempty" xml:"MaliciousMd5,omitempty"`
	MaliciousName *string `json:"MaliciousName,omitempty" xml:"MaliciousName,omitempty"`
	ScanTaskId    *string `json:"ScanTaskId,omitempty" xml:"ScanTaskId,omitempty"`
	UpdateTime    *int64  `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
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
	NextToken    *string                       `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RegionId     *string                       `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceId   []*string                     `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	ResourceType *string                       `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	Tag          []*ListTagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
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
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
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
	NextToken    *string                                   `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId    *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	ResourceId   *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	TagKey       *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	TagValue     *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
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
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The new password that you specify to log on to the instance. The password must be 8 to 32 bits in length, and must contain at least two of the following character types: letters, special characters, and digits.
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
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// *   `true`: The request is successful.
	// *   `false`: The request fails.
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The ID of the request.
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
	RegionId     *string                   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceId   []*string                 `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	ResourceType *string                   `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	Tag          []*TagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
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
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
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
	All          *bool     `json:"All,omitempty" xml:"All,omitempty"`
	RegionId     *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceId   []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	ResourceType *string   `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	TagKey       []*string `json:"TagKey,omitempty" xml:"TagKey,omitempty" type:"Repeated"`
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
	Auto              *bool   `json:"Auto,omitempty" xml:"Auto,omitempty"`
	EnableDeleteTag   *bool   `json:"EnableDeleteTag,omitempty" xml:"EnableDeleteTag,omitempty"`
	InstanceId        *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	NamespaceName     *string `json:"NamespaceName,omitempty" xml:"NamespaceName,omitempty"`
	RepoName          *string `json:"RepoName,omitempty" xml:"RepoName,omitempty"`
	RetentionTagCount *int64  `json:"RetentionTagCount,omitempty" xml:"RetentionTagCount,omitempty"`
	RuleId            *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	ScheduleTime      *string `json:"ScheduleTime,omitempty" xml:"ScheduleTime,omitempty"`
	Scope             *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
	TagRegexp         *string `json:"TagRegexp,omitempty" xml:"TagRegexp,omitempty"`
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
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	IsSuccess *bool   `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
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

type UpdateChainRequest struct {
	// The configuration of the delivery chain in the JSON format.
	ChainConfig *string `json:"ChainConfig,omitempty" xml:"ChainConfig,omitempty"`
	// The ID of the delivery chain.
	ChainId *string `json:"ChainId,omitempty" xml:"ChainId,omitempty"`
	// The description of the delivery chain.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the instance.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the delivery chain.
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
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the request is successful.
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The ID of the request.
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
	// *   `true`: automatically creates repositories in the namespace.
	// *   `false`: does not automatically create repositories in the namespace.
	AutoCreateRepo *bool `json:"AutoCreateRepo,omitempty" xml:"AutoCreateRepo,omitempty"`
	// The default type of the repository. Valid values:
	//
	// *   `PUBLIC`: a public repository
	// *   `PRIVATE`: a private repository
	DefaultRepoType *string `json:"DefaultRepoType,omitempty" xml:"DefaultRepoType,omitempty"`
	// The ID of the instance.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the namespace to which the repository belongs.
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
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// *   `true`: The request is successful.
	// *   `false`: The request fails.
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The ID of the request.
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
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the repository.
	RepoName *string `json:"RepoName,omitempty" xml:"RepoName,omitempty"`
	// The name of the namespace to which the repository belongs.
	RepoNamespaceName *string `json:"RepoNamespaceName,omitempty" xml:"RepoNamespaceName,omitempty"`
	// The type of the repository. Valid values:
	//
	// *   `PUBLIC`: a public repository.
	// *   `PRIVATE`: a private repository.
	RepoType *string `json:"RepoType,omitempty" xml:"RepoType,omitempty"`
	// The summary of the repository.
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
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// *   `true`: The request is successful.
	// *   `false`: The request fails.
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The ID of the request.
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
	EventChannel *string `json:"EventChannel,omitempty" xml:"EventChannel,omitempty"`
	// The event configuration.
	EventConfig *string `json:"EventConfig,omitempty" xml:"EventConfig,omitempty"`
	// The event scope. Valid values:
	//
	// *   `INSTANCE`
	// *   `NAMESPACE`
	// *   `REPO`
	//
	// Default value: `INSTANCE`
	EventScope *string `json:"EventScope,omitempty" xml:"EventScope,omitempty"`
	// The type of the event. Valid values:
	//
	// *   `cr:Artifact:DeliveryChainCompleted`: The delivery chain is processed.
	// *   `cr:Artifact:SynchronizationCompleted`: The image is replicated.
	// *   `cr:Artifact:BuildCompleted`: The image is built.
	// *   `cr:Artifact:ScanCompleted`: The image is scanned.
	// *   `cr:Artifact:SigningCompleted`: The image is signed.
	EventType *string `json:"EventType,omitempty" xml:"EventType,omitempty"`
	// The ID of the instance.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The namespaces to which the event notification rule applies.
	Namespaces []*string `json:"Namespaces,omitempty" xml:"Namespaces,omitempty" type:"Repeated"`
	// The names of the repositories to which the event notification rule applies.
	RepoNames []*string `json:"RepoNames,omitempty" xml:"RepoNames,omitempty" type:"Repeated"`
	// The regular expression for image tags.
	RepoTagFilterPattern *string `json:"RepoTagFilterPattern,omitempty" xml:"RepoTagFilterPattern,omitempty"`
	// The ID of the event notification rule.
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// The name of the event notification rule.
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
	EventChannel *string `json:"EventChannel,omitempty" xml:"EventChannel,omitempty"`
	// The event configuration.
	EventConfig *string `json:"EventConfig,omitempty" xml:"EventConfig,omitempty"`
	// The event scope. Valid values:
	//
	// *   `INSTANCE`
	// *   `NAMESPACE`
	// *   `REPO`
	//
	// Default value: `INSTANCE`
	EventScope *string `json:"EventScope,omitempty" xml:"EventScope,omitempty"`
	// The type of the event. Valid values:
	//
	// *   `cr:Artifact:DeliveryChainCompleted`: The delivery chain is processed.
	// *   `cr:Artifact:SynchronizationCompleted`: The image is replicated.
	// *   `cr:Artifact:BuildCompleted`: The image is built.
	// *   `cr:Artifact:ScanCompleted`: The image is scanned.
	// *   `cr:Artifact:SigningCompleted`: The image is signed.
	EventType *string `json:"EventType,omitempty" xml:"EventType,omitempty"`
	// The ID of the instance.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The namespaces to which the event notification rule applies.
	NamespacesShrink *string `json:"Namespaces,omitempty" xml:"Namespaces,omitempty"`
	// The names of the repositories to which the event notification rule applies.
	RepoNamesShrink *string `json:"RepoNames,omitempty" xml:"RepoNames,omitempty"`
	// The regular expression for image tags.
	RepoTagFilterPattern *string `json:"RepoTagFilterPattern,omitempty" xml:"RepoTagFilterPattern,omitempty"`
	// The ID of the event notification rule.
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// The name of the event notification rule.
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
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the event notification rule.
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
	// *   `true`: enables the instance endpoint.
	// *   `false`: disables the instance endpoint
	Enable *bool `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// The type of the endpoint. Set the value to Internet.
	EndpointType *string `json:"EndpointType,omitempty" xml:"EndpointType,omitempty"`
	// The ID of the instance.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the module that you want to access. Valid values:
	//
	// *   `Registry`: the image repository.
	// *   `Chart`: a Helm chart.
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
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// *   `true`: The request is successful.
	// *   `false`: The request fails.
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The ID of the request.
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
	AutoCreateRepo *bool `json:"AutoCreateRepo,omitempty" xml:"AutoCreateRepo,omitempty"`
	// The default type of the repository. Valid values:
	//
	// *   `PUBLIC`: The repository is a public repository.
	// *   `PRIVATE`: The repository is a private repository.
	DefaultRepoType *string `json:"DefaultRepoType,omitempty" xml:"DefaultRepoType,omitempty"`
	// The ID of the instance.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the namespace.
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

type UpdateNamespaceResponseBody struct {
	// The return value.
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the request is successful.
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The ID of the request.
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
	BuildRuleId *string `json:"BuildRuleId,omitempty" xml:"BuildRuleId,omitempty"`
	// The path of the Dockerfile.
	DockerfileLocation *string `json:"DockerfileLocation,omitempty" xml:"DockerfileLocation,omitempty"`
	// The name of the Dockerfile.
	DockerfileName *string `json:"DockerfileName,omitempty" xml:"DockerfileName,omitempty"`
	// The tag of the image.
	ImageTag *string `json:"ImageTag,omitempty" xml:"ImageTag,omitempty"`
	// The ID of the instance.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Architecture for image building. Valid values:
	//
	// *   `linux/amd64`
	// *   `linux/arm64`
	// *   `linux/386`
	// *   `linux/arm/v7`
	// *   `linux/arm/v6`
	//
	// Default value: `linux/amd64`
	Platforms []*string `json:"Platforms,omitempty" xml:"Platforms,omitempty" type:"Repeated"`
	// The name of the push that triggers the building rule.
	PushName *string `json:"PushName,omitempty" xml:"PushName,omitempty"`
	// The type of the push that triggers the building rule. Valid values:
	//
	// *   `GIT_TAG`: tag push
	// *   `GIT_BRANCH`: branch push
	PushType *string `json:"PushType,omitempty" xml:"PushType,omitempty"`
	// The ID of the image repository.
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
	BuildRuleId *string `json:"BuildRuleId,omitempty" xml:"BuildRuleId,omitempty"`
	// The return value.
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// *   `true`: The request is successful.
	// *   `false`: The request fails.
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The ID of the request.
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
	// *   `true`: enables automatic image building when code is committed.
	// *   `false`: disables automatic image building when code is committed.
	AutoBuild *string `json:"AutoBuild,omitempty" xml:"AutoBuild,omitempty"`
	// The ID of the source code repository.
	CodeRepoId *string `json:"CodeRepoId,omitempty" xml:"CodeRepoId,omitempty"`
	// The name of the source code repository.
	CodeRepoName *string `json:"CodeRepoName,omitempty" xml:"CodeRepoName,omitempty"`
	// The namespace to which the source code repository belongs.
	CodeRepoNamespaceName *string `json:"CodeRepoNamespaceName,omitempty" xml:"CodeRepoNamespaceName,omitempty"`
	// The type of the source code hosting platform. Valid values: GITHUB, GITLAB, GITEE, CODE, and CODEUP.
	CodeRepoType *string `json:"CodeRepoType,omitempty" xml:"CodeRepoType,omitempty"`
	// Specifies whether to disable building caches. Valid values:
	//
	// *   `true`: disables building caches.
	// *   `false`: enables building caches.
	DisableCacheBuild *string `json:"DisableCacheBuild,omitempty" xml:"DisableCacheBuild,omitempty"`
	// The ID of the Container Registry Enterprise Edition instance.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Specifies whether to enable Build With Servers Deployed Outside Chinese Mainland. Valid values:
	//
	// *   `true`: enables Build With Servers Deployed Outside Chinese Mainland.
	// *   `false`: disables Build With Servers Deployed Outside Chinese Mainland.
	OverseaBuild *string `json:"OverseaBuild,omitempty" xml:"OverseaBuild,omitempty"`
	// The ID of the image repository.
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
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the request is successful.
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The ID of the request.
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
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the image repository.
	RepoId *string `json:"RepoId,omitempty" xml:"RepoId,omitempty"`
	// The ID of the trigger.
	TriggerId *string `json:"TriggerId,omitempty" xml:"TriggerId,omitempty"`
	// The name of the trigger.
	//
	// You can specify the TriggerName or TriggerUrl parameter. The TriggerName parameter is optional.
	TriggerName *string `json:"TriggerName,omitempty" xml:"TriggerName,omitempty"`
	// The image tag based on which the trigger is set.
	TriggerTag *string `json:"TriggerTag,omitempty" xml:"TriggerTag,omitempty"`
	// The type of the trigger. Valid values:
	//
	// *   `ALL`: a trigger that supports both tags and regular expressions.
	// *   `TAG_LISTTAG`: a tag-based trigger.
	// *   `TAG_REG_EXP`: a regular expression-based trigger.
	TriggerType *string `json:"TriggerType,omitempty" xml:"TriggerType,omitempty"`
	// The URL of the trigger.
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
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the request is successful.
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The ID of the request.
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
	Detail            *string `json:"Detail,omitempty" xml:"Detail,omitempty"`
	InstanceId        *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RepoId            *string `json:"RepoId,omitempty" xml:"RepoId,omitempty"`
	RepoName          *string `json:"RepoName,omitempty" xml:"RepoName,omitempty"`
	RepoNamespaceName *string `json:"RepoNamespaceName,omitempty" xml:"RepoNamespaceName,omitempty"`
	RepoType          *string `json:"RepoType,omitempty" xml:"RepoType,omitempty"`
	Summary           *string `json:"Summary,omitempty" xml:"Summary,omitempty"`
	TagImmutability   *bool   `json:"TagImmutability,omitempty" xml:"TagImmutability,omitempty"`
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
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	IsSuccess *bool   `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
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

/**
 * You can create building rules of accelerated images only for image repositories of Container Registry Advanced Edition instances. You cannot create building rules of accelerated images for image repositories of Container Registry Basic Edition instances. For more information, see [Specifications of different editions](https://www.alibabacloud.com/help/zh/container-registry/latest/what-is-container-registry#section-go7-lhg-qbc).
 * Accelerated images can be built in the following regions: China (Hangzhou), China (Shanghai), China (Beijing), China (Shenzhen), China (Guangzhou), China (Zhangjiakou), China (Hong Kong), US (Virginia), US (Silicon Valley), Singapore, Japan (Tokyo), and Malaysia (Kuala Lumpur).
 *
 * @param tmpReq CreateArtifactBuildRuleRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return CreateArtifactBuildRuleResponse
 */
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

/**
 * You can create building rules of accelerated images only for image repositories of Container Registry Advanced Edition instances. You cannot create building rules of accelerated images for image repositories of Container Registry Basic Edition instances. For more information, see [Specifications of different editions](https://www.alibabacloud.com/help/zh/container-registry/latest/what-is-container-registry#section-go7-lhg-qbc).
 * Accelerated images can be built in the following regions: China (Hangzhou), China (Shanghai), China (Beijing), China (Shenzhen), China (Guangzhou), China (Zhangjiakou), China (Hong Kong), US (Virginia), US (Silicon Valley), Singapore, Japan (Tokyo), and Malaysia (Kuala Lumpur).
 *
 * @param request CreateArtifactBuildRuleRequest
 * @return CreateArtifactBuildRuleResponse
 */
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

/**
 * A maximum of three VPCs can be associated with a Container Registry instance. If you want to associate more VPCs, contact Alibaba Cloud technical support.
 *
 * @param request CreateInstanceVpcEndpointLinkedVpcRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return CreateInstanceVpcEndpointLinkedVpcResponse
 */
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

/**
 * A maximum of three VPCs can be associated with a Container Registry instance. If you want to associate more VPCs, contact Alibaba Cloud technical support.
 *
 * @param request CreateInstanceVpcEndpointLinkedVpcRequest
 * @return CreateInstanceVpcEndpointLinkedVpcResponse
 */
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

func (client *Client) CreateNamespaceWithOptions(request *CreateNamespaceRequest, runtime *util.RuntimeOptions) (_result *CreateNamespaceResponse, _err error) {
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

/**
 * >  If you delete a chart namespace, all repositories in the namespace and the charts in all repositories are deleted.
 *
 * @param request DeleteChartNamespaceRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DeleteChartNamespaceResponse
 */
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

/**
 * >  If you delete a chart namespace, all repositories in the namespace and the charts in all repositories are deleted.
 *
 * @param request DeleteChartNamespaceRequest
 * @return DeleteChartNamespaceResponse
 */
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

/**
 * > After you delete a namespace, all repositories in the namespace and all images in these repositories are deleted as well.
 *
 * @param request DeleteNamespaceRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DeleteNamespaceResponse
 */
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

/**
 * > After you delete a namespace, all repositories in the namespace and all images in these repositories are deleted as well.
 *
 * @param request DeleteNamespaceRequest
 * @return DeleteNamespaceResponse
 */
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

/**
 * If you delete a repository, all images in the repository are also deleted.
 *
 * @param request DeleteRepositoryRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DeleteRepositoryResponse
 */
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

/**
 * If you delete a repository, all images in the repository are also deleted.
 *
 * @param request DeleteRepositoryRequest
 * @return DeleteRepositoryResponse
 */
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

/**
 * The ID of the Container Registry instance.
 *
 * @param request GetAuthorizationTokenRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return GetAuthorizationTokenResponse
 */
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

/**
 * The ID of the Container Registry instance.
 *
 * @param request GetAuthorizationTokenRequest
 * @return GetAuthorizationTokenResponse
 */
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

/**
 * ****
 *
 * @param request GetRepoBuildRecordRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return GetRepoBuildRecordResponse
 */
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

/**
 * ****
 *
 * @param request GetRepoBuildRecordRequest
 * @return GetRepoBuildRecordResponse
 */
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

func (client *Client) GetRepoTagLayersWithOptions(request *GetRepoTagLayersRequest, runtime *util.RuntimeOptions) (_result *GetRepoTagLayersResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetRepoTagLayers"),
		Version:     tea.String("2018-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetRepoTagLayersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetRepoTagLayers(request *GetRepoTagLayersRequest) (_result *GetRepoTagLayersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetRepoTagLayersResponse{}
	_body, _err := client.GetRepoTagLayersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetRepoTagManifestWithOptions(request *GetRepoTagManifestRequest, runtime *util.RuntimeOptions) (_result *GetRepoTagManifestResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.SchemaVersion)) {
		query["SchemaVersion"] = request.SchemaVersion
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetRepoTagManifest"),
		Version:     tea.String("2018-12-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetRepoTagManifestResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetRepoTagManifest(request *GetRepoTagManifestRequest) (_result *GetRepoTagManifestResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetRepoTagManifestResponse{}
	_body, _err := client.GetRepoTagManifestWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

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

func (client *Client) UpdateNamespaceWithOptions(request *UpdateNamespaceRequest, runtime *util.RuntimeOptions) (_result *UpdateNamespaceResponse, _err error) {
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
