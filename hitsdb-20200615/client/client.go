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
	// The ID of the resource group into which you want to change.
	//
	// This parameter is required.
	//
	// example:
	//
	// rg-aek2i6wee****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The resource ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ld-bp17j28j2y7pm****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
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
	// The detailed reason why the access was denied.
	//
	// example:
	//
	// {"AuthAction":"xxx","AuthPrincipalDisplayName":"222","AuthPrincipalOwnerId":"111","AuthPrincipalType":"SubUser",,"NoPermissionType":"ImplicitDeny","PolicyType":"AccountLevelIdentityBasedPolicy","EncodedDiagnosticMessage":"xxxxxx"}
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// Id of the request
	//
	// example:
	//
	// FAED4C02-AF99-5015-A075-692DE9C99630
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ChangeResourceGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ChangeResourceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ChangeResourceGroupResponseBody) SetAccessDeniedDetail(v string) *ChangeResourceGroupResponseBody {
	s.AccessDeniedDetail = &v
	return s
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

type CheckLdpsColumnarIndexStatusRequest struct {
	// This parameter is required.
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s CheckLdpsColumnarIndexStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s CheckLdpsColumnarIndexStatusRequest) GoString() string {
	return s.String()
}

func (s *CheckLdpsColumnarIndexStatusRequest) SetInstanceId(v string) *CheckLdpsColumnarIndexStatusRequest {
	s.InstanceId = &v
	return s
}

func (s *CheckLdpsColumnarIndexStatusRequest) SetOwnerAccount(v string) *CheckLdpsColumnarIndexStatusRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CheckLdpsColumnarIndexStatusRequest) SetOwnerId(v int64) *CheckLdpsColumnarIndexStatusRequest {
	s.OwnerId = &v
	return s
}

func (s *CheckLdpsColumnarIndexStatusRequest) SetResourceOwnerAccount(v string) *CheckLdpsColumnarIndexStatusRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CheckLdpsColumnarIndexStatusRequest) SetResourceOwnerId(v int64) *CheckLdpsColumnarIndexStatusRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CheckLdpsColumnarIndexStatusRequest) SetSecurityToken(v string) *CheckLdpsColumnarIndexStatusRequest {
	s.SecurityToken = &v
	return s
}

type CheckLdpsColumnarIndexStatusResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	Opened             *bool   `json:"Opened,omitempty" xml:"Opened,omitempty"`
	RequestId          *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CheckLdpsColumnarIndexStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CheckLdpsColumnarIndexStatusResponseBody) GoString() string {
	return s.String()
}

func (s *CheckLdpsColumnarIndexStatusResponseBody) SetAccessDeniedDetail(v string) *CheckLdpsColumnarIndexStatusResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CheckLdpsColumnarIndexStatusResponseBody) SetOpened(v bool) *CheckLdpsColumnarIndexStatusResponseBody {
	s.Opened = &v
	return s
}

func (s *CheckLdpsColumnarIndexStatusResponseBody) SetRequestId(v string) *CheckLdpsColumnarIndexStatusResponseBody {
	s.RequestId = &v
	return s
}

type CheckLdpsColumnarIndexStatusResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckLdpsColumnarIndexStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckLdpsColumnarIndexStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s CheckLdpsColumnarIndexStatusResponse) GoString() string {
	return s.String()
}

func (s *CheckLdpsColumnarIndexStatusResponse) SetHeaders(v map[string]*string) *CheckLdpsColumnarIndexStatusResponse {
	s.Headers = v
	return s
}

func (s *CheckLdpsColumnarIndexStatusResponse) SetStatusCode(v int32) *CheckLdpsColumnarIndexStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckLdpsColumnarIndexStatusResponse) SetBody(v *CheckLdpsColumnarIndexStatusResponseBody) *CheckLdpsColumnarIndexStatusResponse {
	s.Body = v
	return s
}

type CreateAutoScalingConfigRequest struct {
	// This parameter is required.
	ConfigName         *string `json:"ConfigName,omitempty" xml:"ConfigName,omitempty"`
	EffectiveTimeEnd   *string `json:"EffectiveTimeEnd,omitempty" xml:"EffectiveTimeEnd,omitempty"`
	EffectiveTimeStart *string `json:"EffectiveTimeStart,omitempty" xml:"EffectiveTimeStart,omitempty"`
	Enabled            *bool   `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// This parameter is required.
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// This parameter is required.
	InstanceId           *string                                        `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	NodesMax             *int32                                         `json:"NodesMax,omitempty" xml:"NodesMax,omitempty"`
	NodesMin             *int32                                         `json:"NodesMin,omitempty" xml:"NodesMin,omitempty"`
	OwnerAccount         *string                                        `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64                                         `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string                                        `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64                                         `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ScaleRuleList        []*CreateAutoScalingConfigRequestScaleRuleList `json:"ScaleRuleList,omitempty" xml:"ScaleRuleList,omitempty" type:"Repeated"`
	// This parameter is required.
	ScaleType     *string `json:"ScaleType,omitempty" xml:"ScaleType,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// This parameter is required.
	SpecId *string `json:"SpecId,omitempty" xml:"SpecId,omitempty"`
}

func (s CreateAutoScalingConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAutoScalingConfigRequest) GoString() string {
	return s.String()
}

func (s *CreateAutoScalingConfigRequest) SetConfigName(v string) *CreateAutoScalingConfigRequest {
	s.ConfigName = &v
	return s
}

func (s *CreateAutoScalingConfigRequest) SetEffectiveTimeEnd(v string) *CreateAutoScalingConfigRequest {
	s.EffectiveTimeEnd = &v
	return s
}

func (s *CreateAutoScalingConfigRequest) SetEffectiveTimeStart(v string) *CreateAutoScalingConfigRequest {
	s.EffectiveTimeStart = &v
	return s
}

func (s *CreateAutoScalingConfigRequest) SetEnabled(v bool) *CreateAutoScalingConfigRequest {
	s.Enabled = &v
	return s
}

func (s *CreateAutoScalingConfigRequest) SetEngine(v string) *CreateAutoScalingConfigRequest {
	s.Engine = &v
	return s
}

func (s *CreateAutoScalingConfigRequest) SetInstanceId(v string) *CreateAutoScalingConfigRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateAutoScalingConfigRequest) SetNodesMax(v int32) *CreateAutoScalingConfigRequest {
	s.NodesMax = &v
	return s
}

func (s *CreateAutoScalingConfigRequest) SetNodesMin(v int32) *CreateAutoScalingConfigRequest {
	s.NodesMin = &v
	return s
}

func (s *CreateAutoScalingConfigRequest) SetOwnerAccount(v string) *CreateAutoScalingConfigRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateAutoScalingConfigRequest) SetOwnerId(v int64) *CreateAutoScalingConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateAutoScalingConfigRequest) SetResourceOwnerAccount(v string) *CreateAutoScalingConfigRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateAutoScalingConfigRequest) SetResourceOwnerId(v int64) *CreateAutoScalingConfigRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateAutoScalingConfigRequest) SetScaleRuleList(v []*CreateAutoScalingConfigRequestScaleRuleList) *CreateAutoScalingConfigRequest {
	s.ScaleRuleList = v
	return s
}

func (s *CreateAutoScalingConfigRequest) SetScaleType(v string) *CreateAutoScalingConfigRequest {
	s.ScaleType = &v
	return s
}

func (s *CreateAutoScalingConfigRequest) SetSecurityToken(v string) *CreateAutoScalingConfigRequest {
	s.SecurityToken = &v
	return s
}

func (s *CreateAutoScalingConfigRequest) SetSpecId(v string) *CreateAutoScalingConfigRequest {
	s.SpecId = &v
	return s
}

type CreateAutoScalingConfigRequestScaleRuleList struct {
	ConfigId          *string `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	Enabled           *bool   `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	EndTime           *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	InstanceId        *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	ObservationWindow *int32  `json:"ObservationWindow,omitempty" xml:"ObservationWindow,omitempty"`
	OperationType     *string `json:"OperationType,omitempty" xml:"OperationType,omitempty"`
	RuleId            *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	RuleName          *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	RuleType          *string `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
	ScaleInStep       *int32  `json:"ScaleInStep,omitempty" xml:"ScaleInStep,omitempty"`
	ScaleOutStep      *int32  `json:"ScaleOutStep,omitempty" xml:"ScaleOutStep,omitempty"`
	SilenceTime       *int32  `json:"SilenceTime,omitempty" xml:"SilenceTime,omitempty"`
	StartTime         *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	TargetMetric      *string `json:"TargetMetric,omitempty" xml:"TargetMetric,omitempty"`
	TargetNodes       *int32  `json:"TargetNodes,omitempty" xml:"TargetNodes,omitempty"`
	ThresholdLower    *int32  `json:"ThresholdLower,omitempty" xml:"ThresholdLower,omitempty"`
	ThresholdUpper    *int32  `json:"ThresholdUpper,omitempty" xml:"ThresholdUpper,omitempty"`
	TriggerCronExpr   *string `json:"TriggerCronExpr,omitempty" xml:"TriggerCronExpr,omitempty"`
}

func (s CreateAutoScalingConfigRequestScaleRuleList) String() string {
	return tea.Prettify(s)
}

func (s CreateAutoScalingConfigRequestScaleRuleList) GoString() string {
	return s.String()
}

func (s *CreateAutoScalingConfigRequestScaleRuleList) SetConfigId(v string) *CreateAutoScalingConfigRequestScaleRuleList {
	s.ConfigId = &v
	return s
}

func (s *CreateAutoScalingConfigRequestScaleRuleList) SetEnabled(v bool) *CreateAutoScalingConfigRequestScaleRuleList {
	s.Enabled = &v
	return s
}

func (s *CreateAutoScalingConfigRequestScaleRuleList) SetEndTime(v string) *CreateAutoScalingConfigRequestScaleRuleList {
	s.EndTime = &v
	return s
}

func (s *CreateAutoScalingConfigRequestScaleRuleList) SetInstanceId(v string) *CreateAutoScalingConfigRequestScaleRuleList {
	s.InstanceId = &v
	return s
}

func (s *CreateAutoScalingConfigRequestScaleRuleList) SetObservationWindow(v int32) *CreateAutoScalingConfigRequestScaleRuleList {
	s.ObservationWindow = &v
	return s
}

func (s *CreateAutoScalingConfigRequestScaleRuleList) SetOperationType(v string) *CreateAutoScalingConfigRequestScaleRuleList {
	s.OperationType = &v
	return s
}

func (s *CreateAutoScalingConfigRequestScaleRuleList) SetRuleId(v string) *CreateAutoScalingConfigRequestScaleRuleList {
	s.RuleId = &v
	return s
}

func (s *CreateAutoScalingConfigRequestScaleRuleList) SetRuleName(v string) *CreateAutoScalingConfigRequestScaleRuleList {
	s.RuleName = &v
	return s
}

func (s *CreateAutoScalingConfigRequestScaleRuleList) SetRuleType(v string) *CreateAutoScalingConfigRequestScaleRuleList {
	s.RuleType = &v
	return s
}

func (s *CreateAutoScalingConfigRequestScaleRuleList) SetScaleInStep(v int32) *CreateAutoScalingConfigRequestScaleRuleList {
	s.ScaleInStep = &v
	return s
}

func (s *CreateAutoScalingConfigRequestScaleRuleList) SetScaleOutStep(v int32) *CreateAutoScalingConfigRequestScaleRuleList {
	s.ScaleOutStep = &v
	return s
}

func (s *CreateAutoScalingConfigRequestScaleRuleList) SetSilenceTime(v int32) *CreateAutoScalingConfigRequestScaleRuleList {
	s.SilenceTime = &v
	return s
}

func (s *CreateAutoScalingConfigRequestScaleRuleList) SetStartTime(v string) *CreateAutoScalingConfigRequestScaleRuleList {
	s.StartTime = &v
	return s
}

func (s *CreateAutoScalingConfigRequestScaleRuleList) SetTargetMetric(v string) *CreateAutoScalingConfigRequestScaleRuleList {
	s.TargetMetric = &v
	return s
}

func (s *CreateAutoScalingConfigRequestScaleRuleList) SetTargetNodes(v int32) *CreateAutoScalingConfigRequestScaleRuleList {
	s.TargetNodes = &v
	return s
}

func (s *CreateAutoScalingConfigRequestScaleRuleList) SetThresholdLower(v int32) *CreateAutoScalingConfigRequestScaleRuleList {
	s.ThresholdLower = &v
	return s
}

func (s *CreateAutoScalingConfigRequestScaleRuleList) SetThresholdUpper(v int32) *CreateAutoScalingConfigRequestScaleRuleList {
	s.ThresholdUpper = &v
	return s
}

func (s *CreateAutoScalingConfigRequestScaleRuleList) SetTriggerCronExpr(v string) *CreateAutoScalingConfigRequestScaleRuleList {
	s.TriggerCronExpr = &v
	return s
}

type CreateAutoScalingConfigShrinkRequest struct {
	// This parameter is required.
	ConfigName         *string `json:"ConfigName,omitempty" xml:"ConfigName,omitempty"`
	EffectiveTimeEnd   *string `json:"EffectiveTimeEnd,omitempty" xml:"EffectiveTimeEnd,omitempty"`
	EffectiveTimeStart *string `json:"EffectiveTimeStart,omitempty" xml:"EffectiveTimeStart,omitempty"`
	Enabled            *bool   `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// This parameter is required.
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// This parameter is required.
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	NodesMax             *int32  `json:"NodesMax,omitempty" xml:"NodesMax,omitempty"`
	NodesMin             *int32  `json:"NodesMin,omitempty" xml:"NodesMin,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ScaleRuleListShrink  *string `json:"ScaleRuleList,omitempty" xml:"ScaleRuleList,omitempty"`
	// This parameter is required.
	ScaleType     *string `json:"ScaleType,omitempty" xml:"ScaleType,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// This parameter is required.
	SpecId *string `json:"SpecId,omitempty" xml:"SpecId,omitempty"`
}

func (s CreateAutoScalingConfigShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAutoScalingConfigShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateAutoScalingConfigShrinkRequest) SetConfigName(v string) *CreateAutoScalingConfigShrinkRequest {
	s.ConfigName = &v
	return s
}

func (s *CreateAutoScalingConfigShrinkRequest) SetEffectiveTimeEnd(v string) *CreateAutoScalingConfigShrinkRequest {
	s.EffectiveTimeEnd = &v
	return s
}

func (s *CreateAutoScalingConfigShrinkRequest) SetEffectiveTimeStart(v string) *CreateAutoScalingConfigShrinkRequest {
	s.EffectiveTimeStart = &v
	return s
}

func (s *CreateAutoScalingConfigShrinkRequest) SetEnabled(v bool) *CreateAutoScalingConfigShrinkRequest {
	s.Enabled = &v
	return s
}

func (s *CreateAutoScalingConfigShrinkRequest) SetEngine(v string) *CreateAutoScalingConfigShrinkRequest {
	s.Engine = &v
	return s
}

func (s *CreateAutoScalingConfigShrinkRequest) SetInstanceId(v string) *CreateAutoScalingConfigShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateAutoScalingConfigShrinkRequest) SetNodesMax(v int32) *CreateAutoScalingConfigShrinkRequest {
	s.NodesMax = &v
	return s
}

func (s *CreateAutoScalingConfigShrinkRequest) SetNodesMin(v int32) *CreateAutoScalingConfigShrinkRequest {
	s.NodesMin = &v
	return s
}

func (s *CreateAutoScalingConfigShrinkRequest) SetOwnerAccount(v string) *CreateAutoScalingConfigShrinkRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateAutoScalingConfigShrinkRequest) SetOwnerId(v int64) *CreateAutoScalingConfigShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateAutoScalingConfigShrinkRequest) SetResourceOwnerAccount(v string) *CreateAutoScalingConfigShrinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateAutoScalingConfigShrinkRequest) SetResourceOwnerId(v int64) *CreateAutoScalingConfigShrinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateAutoScalingConfigShrinkRequest) SetScaleRuleListShrink(v string) *CreateAutoScalingConfigShrinkRequest {
	s.ScaleRuleListShrink = &v
	return s
}

func (s *CreateAutoScalingConfigShrinkRequest) SetScaleType(v string) *CreateAutoScalingConfigShrinkRequest {
	s.ScaleType = &v
	return s
}

func (s *CreateAutoScalingConfigShrinkRequest) SetSecurityToken(v string) *CreateAutoScalingConfigShrinkRequest {
	s.SecurityToken = &v
	return s
}

func (s *CreateAutoScalingConfigShrinkRequest) SetSpecId(v string) *CreateAutoScalingConfigShrinkRequest {
	s.SpecId = &v
	return s
}

type CreateAutoScalingConfigResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateAutoScalingConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateAutoScalingConfigResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAutoScalingConfigResponseBody) SetCode(v string) *CreateAutoScalingConfigResponseBody {
	s.Code = &v
	return s
}

func (s *CreateAutoScalingConfigResponseBody) SetHttpStatusCode(v int32) *CreateAutoScalingConfigResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateAutoScalingConfigResponseBody) SetMessage(v string) *CreateAutoScalingConfigResponseBody {
	s.Message = &v
	return s
}

func (s *CreateAutoScalingConfigResponseBody) SetRequestId(v string) *CreateAutoScalingConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAutoScalingConfigResponseBody) SetSuccess(v bool) *CreateAutoScalingConfigResponseBody {
	s.Success = &v
	return s
}

type CreateAutoScalingConfigResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAutoScalingConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAutoScalingConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateAutoScalingConfigResponse) GoString() string {
	return s.String()
}

func (s *CreateAutoScalingConfigResponse) SetHeaders(v map[string]*string) *CreateAutoScalingConfigResponse {
	s.Headers = v
	return s
}

func (s *CreateAutoScalingConfigResponse) SetStatusCode(v int32) *CreateAutoScalingConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAutoScalingConfigResponse) SetBody(v *CreateAutoScalingConfigResponseBody) *CreateAutoScalingConfigResponse {
	s.Body = v
	return s
}

type CreateAutoScalingRuleRequest struct {
	// This parameter is required.
	ConfigId *string `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	Enabled  *bool   `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	EndTime  *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// This parameter is required.
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	ObservationWindow    *int32  `json:"ObservationWindow,omitempty" xml:"ObservationWindow,omitempty"`
	OperationType        *string `json:"OperationType,omitempty" xml:"OperationType,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// This parameter is required.
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// This parameter is required.
	RuleType        *string `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
	ScaleInStep     *int32  `json:"ScaleInStep,omitempty" xml:"ScaleInStep,omitempty"`
	ScaleOutStep    *int32  `json:"ScaleOutStep,omitempty" xml:"ScaleOutStep,omitempty"`
	SecurityToken   *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	SilenceTime     *int32  `json:"SilenceTime,omitempty" xml:"SilenceTime,omitempty"`
	StartTime       *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	TargetMetric    *string `json:"TargetMetric,omitempty" xml:"TargetMetric,omitempty"`
	TargetNodes     *int32  `json:"TargetNodes,omitempty" xml:"TargetNodes,omitempty"`
	ThresholdLower  *int32  `json:"ThresholdLower,omitempty" xml:"ThresholdLower,omitempty"`
	ThresholdUpper  *int32  `json:"ThresholdUpper,omitempty" xml:"ThresholdUpper,omitempty"`
	TriggerCronExpr *string `json:"TriggerCronExpr,omitempty" xml:"TriggerCronExpr,omitempty"`
}

func (s CreateAutoScalingRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAutoScalingRuleRequest) GoString() string {
	return s.String()
}

func (s *CreateAutoScalingRuleRequest) SetConfigId(v string) *CreateAutoScalingRuleRequest {
	s.ConfigId = &v
	return s
}

func (s *CreateAutoScalingRuleRequest) SetEnabled(v bool) *CreateAutoScalingRuleRequest {
	s.Enabled = &v
	return s
}

func (s *CreateAutoScalingRuleRequest) SetEndTime(v string) *CreateAutoScalingRuleRequest {
	s.EndTime = &v
	return s
}

func (s *CreateAutoScalingRuleRequest) SetInstanceId(v string) *CreateAutoScalingRuleRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateAutoScalingRuleRequest) SetObservationWindow(v int32) *CreateAutoScalingRuleRequest {
	s.ObservationWindow = &v
	return s
}

func (s *CreateAutoScalingRuleRequest) SetOperationType(v string) *CreateAutoScalingRuleRequest {
	s.OperationType = &v
	return s
}

func (s *CreateAutoScalingRuleRequest) SetOwnerAccount(v string) *CreateAutoScalingRuleRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateAutoScalingRuleRequest) SetOwnerId(v int64) *CreateAutoScalingRuleRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateAutoScalingRuleRequest) SetResourceOwnerAccount(v string) *CreateAutoScalingRuleRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateAutoScalingRuleRequest) SetResourceOwnerId(v int64) *CreateAutoScalingRuleRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateAutoScalingRuleRequest) SetRuleName(v string) *CreateAutoScalingRuleRequest {
	s.RuleName = &v
	return s
}

func (s *CreateAutoScalingRuleRequest) SetRuleType(v string) *CreateAutoScalingRuleRequest {
	s.RuleType = &v
	return s
}

func (s *CreateAutoScalingRuleRequest) SetScaleInStep(v int32) *CreateAutoScalingRuleRequest {
	s.ScaleInStep = &v
	return s
}

func (s *CreateAutoScalingRuleRequest) SetScaleOutStep(v int32) *CreateAutoScalingRuleRequest {
	s.ScaleOutStep = &v
	return s
}

func (s *CreateAutoScalingRuleRequest) SetSecurityToken(v string) *CreateAutoScalingRuleRequest {
	s.SecurityToken = &v
	return s
}

func (s *CreateAutoScalingRuleRequest) SetSilenceTime(v int32) *CreateAutoScalingRuleRequest {
	s.SilenceTime = &v
	return s
}

func (s *CreateAutoScalingRuleRequest) SetStartTime(v string) *CreateAutoScalingRuleRequest {
	s.StartTime = &v
	return s
}

func (s *CreateAutoScalingRuleRequest) SetTargetMetric(v string) *CreateAutoScalingRuleRequest {
	s.TargetMetric = &v
	return s
}

func (s *CreateAutoScalingRuleRequest) SetTargetNodes(v int32) *CreateAutoScalingRuleRequest {
	s.TargetNodes = &v
	return s
}

func (s *CreateAutoScalingRuleRequest) SetThresholdLower(v int32) *CreateAutoScalingRuleRequest {
	s.ThresholdLower = &v
	return s
}

func (s *CreateAutoScalingRuleRequest) SetThresholdUpper(v int32) *CreateAutoScalingRuleRequest {
	s.ThresholdUpper = &v
	return s
}

func (s *CreateAutoScalingRuleRequest) SetTriggerCronExpr(v string) *CreateAutoScalingRuleRequest {
	s.TriggerCronExpr = &v
	return s
}

type CreateAutoScalingRuleResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateAutoScalingRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateAutoScalingRuleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAutoScalingRuleResponseBody) SetCode(v string) *CreateAutoScalingRuleResponseBody {
	s.Code = &v
	return s
}

func (s *CreateAutoScalingRuleResponseBody) SetHttpStatusCode(v int32) *CreateAutoScalingRuleResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateAutoScalingRuleResponseBody) SetMessage(v string) *CreateAutoScalingRuleResponseBody {
	s.Message = &v
	return s
}

func (s *CreateAutoScalingRuleResponseBody) SetRequestId(v string) *CreateAutoScalingRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAutoScalingRuleResponseBody) SetSuccess(v bool) *CreateAutoScalingRuleResponseBody {
	s.Success = &v
	return s
}

type CreateAutoScalingRuleResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAutoScalingRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAutoScalingRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateAutoScalingRuleResponse) GoString() string {
	return s.String()
}

func (s *CreateAutoScalingRuleResponse) SetHeaders(v map[string]*string) *CreateAutoScalingRuleResponse {
	s.Headers = v
	return s
}

func (s *CreateAutoScalingRuleResponse) SetStatusCode(v int32) *CreateAutoScalingRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAutoScalingRuleResponse) SetBody(v *CreateAutoScalingRuleResponseBody) *CreateAutoScalingRuleResponse {
	s.Body = v
	return s
}

type CreateLdpsComputeGroupRequest struct {
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// This parameter is required.
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	Properties           *string `json:"Properties,omitempty" xml:"Properties,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s CreateLdpsComputeGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateLdpsComputeGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateLdpsComputeGroupRequest) SetGroupName(v string) *CreateLdpsComputeGroupRequest {
	s.GroupName = &v
	return s
}

func (s *CreateLdpsComputeGroupRequest) SetInstanceId(v string) *CreateLdpsComputeGroupRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateLdpsComputeGroupRequest) SetOwnerAccount(v string) *CreateLdpsComputeGroupRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateLdpsComputeGroupRequest) SetOwnerId(v int64) *CreateLdpsComputeGroupRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateLdpsComputeGroupRequest) SetProperties(v string) *CreateLdpsComputeGroupRequest {
	s.Properties = &v
	return s
}

func (s *CreateLdpsComputeGroupRequest) SetRegionId(v string) *CreateLdpsComputeGroupRequest {
	s.RegionId = &v
	return s
}

func (s *CreateLdpsComputeGroupRequest) SetResourceOwnerAccount(v string) *CreateLdpsComputeGroupRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateLdpsComputeGroupRequest) SetResourceOwnerId(v int64) *CreateLdpsComputeGroupRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateLdpsComputeGroupRequest) SetSecurityToken(v string) *CreateLdpsComputeGroupRequest {
	s.SecurityToken = &v
	return s
}

type CreateLdpsComputeGroupResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	RequestId          *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateLdpsComputeGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateLdpsComputeGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateLdpsComputeGroupResponseBody) SetAccessDeniedDetail(v string) *CreateLdpsComputeGroupResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CreateLdpsComputeGroupResponseBody) SetRequestId(v string) *CreateLdpsComputeGroupResponseBody {
	s.RequestId = &v
	return s
}

type CreateLdpsComputeGroupResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateLdpsComputeGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateLdpsComputeGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateLdpsComputeGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateLdpsComputeGroupResponse) SetHeaders(v map[string]*string) *CreateLdpsComputeGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateLdpsComputeGroupResponse) SetStatusCode(v int32) *CreateLdpsComputeGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateLdpsComputeGroupResponse) SetBody(v *CreateLdpsComputeGroupResponseBody) *CreateLdpsComputeGroupResponse {
	s.Body = v
	return s
}

type CreateLindormInstanceRequest struct {
	// The ID of the vSwitch that is specified for the zone for the coordinate node of the instance. The vSwitch must be deployed in the zone specified by the ArbiterZoneId parameter. **This parameter is required if you want to create a multi-zone instance**.
	//
	// example:
	//
	// vsw-uf6664pqjawb87k36****
	ArbiterVSwitchId *string `json:"ArbiterVSwitchId,omitempty" xml:"ArbiterVSwitchId,omitempty"`
	// The ID of the zone for the coordinate node of the instance. **This parameter is required if you want to create a multi-zone instance**.
	//
	// example:
	//
	// cn-shanghai-g
	ArbiterZoneId *string `json:"ArbiterZoneId,omitempty" xml:"ArbiterZoneId,omitempty"`
	// The architecture of the instance. Valid values:
	//
	// 	- **1.0**: The instance that you want to create is a single-zone instance.
	//
	// 	- **2.0**: The instance that you want to create is a multi-zone instance.
	//
	// By default, the value of this parameter is 1.0. To create a multi-zone instance, set this parameter to 2.0. **This parameter is required if you want to create a multi-zone instance**.
	//
	// example:
	//
	// 2.0
	ArchVersion *string `json:"ArchVersion,omitempty" xml:"ArchVersion,omitempty"`
	// The auto-renewal duration. Unit: month.
	//
	// Valid values: **1*	- to **12**.
	//
	// >  This parameter is available only when the **AutoRenewal*	- parameter is set to **true**.
	//
	// example:
	//
	// 1
	AutoRenewDuration *string `json:"AutoRenewDuration,omitempty" xml:"AutoRenewDuration,omitempty"`
	// Specifies whether to enable auto-renewal for the instance. Valid values:
	//
	// 	- **true**: enables auto-renewal.
	//
	// 	- **false**: disables auto-renewal.
	//
	// Default value: false.
	//
	// >  This parameter is available only when the **PayType*	- parameter is set to **PREPAY**.
	//
	// example:
	//
	// false
	AutoRenewal *bool `json:"AutoRenewal,omitempty" xml:"AutoRenewal,omitempty"`
	// The cold storage capacity of the instance. By default, if you leave this parameter unspecified, cold storage is not enabled for the instance. Unit: GB. Valid values: **800*	- to **1000000**.
	//
	// example:
	//
	// 800
	ColdStorage *int32 `json:"ColdStorage,omitempty" xml:"ColdStorage,omitempty"`
	// The storage capacity of the disk of a single core node. Valid values: 400 to 64000. Unit: GB. **This parameter is required if you want to create a multi-zone instance**.
	//
	// example:
	//
	// 400
	CoreSingleStorage *int32 `json:"CoreSingleStorage,omitempty" xml:"CoreSingleStorage,omitempty"`
	// The specification of the nodes in the instance if you set DiskCategory to local_ssd_pro or local_hdd_pro.
	//
	// When DiskCategory is set to local_ssd_pro, you can set this parameter to the following values:
	//
	// 	- **lindorm.i2.xlarge**: Each node has 4 dedicated CPU cores and 32 GB of dedicated memory.
	//
	// 	- **lindorm.i2.2xlarge**: Each node has 8 dedicated CPU cores and 64 GB of dedicated memory.
	//
	// 	- **lindorm.i2.4xlarge**: Each node has 16 dedicated CPU cores and 128 GB of dedicated memory.
	//
	// 	- **lindorm.i2.8xlarge**: Each node has 32 dedicated CPU cores and 256 GB of dedicated memory.
	//
	// When DiskCategory is set to local_hdd_pro, you can set this parameter to the following values:
	//
	// 	- **lindorm.d1.2xlarge**: Each node has 8 dedicated CPU cores and 32 GB of dedicated memory.
	//
	// 	- **lindorm.d1.4xlarge**: Each node has 16 dedicated CPU cores and 64 GB of dedicated memory.
	//
	// 	- **lindorm.d1.6xlarge**: Each node has 24 dedicated CPU cores and 96 GB of dedicated memory.
	//
	// example:
	//
	// lindorm.i2.xlarge
	CoreSpec *string `json:"CoreSpec,omitempty" xml:"CoreSpec,omitempty"`
	// The storage type of the instance. Valid values:
	//
	// 	- **cloud_efficiency**: This instance uses the Standard type of storage.
	//
	// 	- **cloud_ssd**: This instance uses the Performance type of storage.
	//
	// 	- **capacity_cloud_storage**: This instance uses the Capacity type of storage.
	//
	// 	- **local_ssd_pro**: This instance uses local SSDs.
	//
	// 	- **local_hdd_pro**: This instance uses local HDDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// cloud_efficiency
	DiskCategory *string `json:"DiskCategory,omitempty" xml:"DiskCategory,omitempty"`
	// The subscription period of the instance. The valid values of this parameter depend on the value of the PricingCycle parameter.
	//
	// 	- If PricingCycle is set to **Month**, set this parameter to an integer that ranges from **1*	- to **9**.
	//
	// 	- If PricingCycle is set to **Year**, set this parameter to an integer that ranges from **1*	- to **3**.
	//
	// > This parameter is available and required when the PayType parameter is set to **PREPAY**.
	//
	// example:
	//
	// 1
	Duration *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The number of LindormDFS nodes in the instance. The valid values of this parameter depend on the value of the PayType parameter.
	//
	// 	- If the PayType parameter is set to **PREPAY**, set this parameter to an integer that ranges from **0*	- to **60**.
	//
	// 	- If the PayType parameter is set to **POSTPAY**, set this parameter to an integer that ranges from **0*	- to **8**.
	//
	// example:
	//
	// 2
	FilestoreNum *int32 `json:"FilestoreNum,omitempty" xml:"FilestoreNum,omitempty"`
	// The specification of LindormDFS nodes in the instance. Set the value of this parameter to **lindorm.c.xlarge**, which indicates that each node has 4 dedicated CPU cores and 8 GB of dedicated memory.
	//
	// example:
	//
	// lindorm.c.xlarge
	FilestoreSpec *string `json:"FilestoreSpec,omitempty" xml:"FilestoreSpec,omitempty"`
	// The name of the instance that you want to create.
	//
	// example:
	//
	// lindorm_test
	InstanceAlias *string `json:"InstanceAlias,omitempty" xml:"InstanceAlias,omitempty"`
	// The storage capacity of the instance you want to create. Unit: GB.
	//
	// example:
	//
	// 480
	InstanceStorage *string `json:"InstanceStorage,omitempty" xml:"InstanceStorage,omitempty"`
	// The number of LindormTable nodes in the instance. The valid values of this parameter depend on the value of the PayType parameter.
	//
	// 	- If the PayType parameter is set to **PREPAY**, set this parameter to an integer that ranges from **0*	- to **90**.
	//
	// 	- If the PayType parameter is set to **POSTPAY**, set this parameter to an integer that ranges from **0*	- to **400**.
	//
	// **This parameter is required if you want to create a multi-zone instance**.  The valid values of this parameter range from 4 to 400 if you want to create a multi-zone instance.
	//
	// example:
	//
	// 2
	LindormNum *int32 `json:"LindormNum,omitempty" xml:"LindormNum,omitempty"`
	// The specification of LindormTable nodes in the instance. Valid values:
	//
	// 	- **lindorm.c.xlarge**: Each node has 4 dedicated CPU cores and 8 GB of dedicated memory.
	//
	// 	- **lindorm.c.2xlarge**: Each node has 8 dedicated CPU cores and 16 GB of dedicated memory.
	//
	// 	- **lindorm.c.4xlarge**: Each node has 16 dedicated CPU cores and 32 GB of dedicated memory.
	//
	// 	- **lindorm.c.8xlarge**: Each node has 32 dedicated CPU cores and 64 GB of dedicated memory.
	//
	// example:
	//
	// lindorm.c.xlarge
	LindormSpec *string `json:"LindormSpec,omitempty" xml:"LindormSpec,omitempty"`
	// The disk type of the log nodes. Valid values:
	//
	// 	- **cloud_efficiency**: This instance uses the Standard type of storage.
	//
	// 	- **cloud_ssd**: This instance uses the Performance type of storage.
	//
	// **This parameter is required if you want to create a multi-zone instance**.
	//
	// example:
	//
	// cloud_ssd
	LogDiskCategory *string `json:"LogDiskCategory,omitempty" xml:"LogDiskCategory,omitempty"`
	// The number of the log nodes. Valid values: 4 to 400. **This parameter is required if you want to create a multi-zone instance**.
	//
	// example:
	//
	// 4
	LogNum *int32 `json:"LogNum,omitempty" xml:"LogNum,omitempty"`
	// The storage capacity of the disk of a single log node. Valid values: 400 to 64000. Unit: GB. **This parameter is required if you want to create a multi-zone instance**.
	//
	// example:
	//
	// 400
	LogSingleStorage *int32 `json:"LogSingleStorage,omitempty" xml:"LogSingleStorage,omitempty"`
	// The type of the log nodes. Valid values:
	//
	// 	- **lindorm.sn1.xlarge**: Each node has 4 dedicated CPU cores and 8 GB of dedicated memory.
	//
	// 	- **lindorm.sn1.2xlarge**: Each node has 8 dedicated CPU cores and 16 GB of dedicated memory.
	//
	// **This parameter is required if you want to create a multi-zone instance**.
	//
	// example:
	//
	// lindorm.sn1.large
	LogSpec *string `json:"LogSpec,omitempty" xml:"LogSpec,omitempty"`
	// The number of LTS nodes in the instance. Valid values: **0*	- to **60**.
	//
	// example:
	//
	// 2
	LtsNum *string `json:"LtsNum,omitempty" xml:"LtsNum,omitempty"`
	// The specification of LTS nodes in the instance. Valid values:
	//
	// 	- **lindorm.c.xlarge**: Each node has 4 dedicated CPU cores and 8 GB of dedicated memory.
	//
	// 	- **lindorm.g.xlarge**: Each node has 4 dedicated CPU cores and 16 GB of dedicated memory.
	//
	// 	- **lindorm.c.2xlarge**: Each node has 8 dedicated CPU cores and 16 GB of dedicated memory.
	//
	// 	- **lindorm.g.2xlarge**: Each node has 8 dedicated CPU cores and 32 GB of dedicated memory.
	//
	// 	- **lindorm.c.4xlarge**: Each node has 16 dedicated CPU cores and 32 GB of dedicated memory.
	//
	// 	- **lindorm.g.4xlarge**: Each node has 16 dedicated CPU cores and 64 GB of dedicated memory.
	//
	// 	- **lindorm.c.8xlarge**: Each node has 32 dedicated CPU cores and 64 GB of dedicated memory.
	//
	// 	- **lindorm.g.8xlarge**: Each node has 32 dedicated CPU cores and 128 GB of dedicated memory.
	//
	// example:
	//
	// lindorm.g.xlarge
	LtsSpec *string `json:"LtsSpec,omitempty" xml:"LtsSpec,omitempty"`
	// The combinations of zones that are available for the multi-zone instance. You can go to the purchase page of Lindorm to view the supported zone combinations.
	//
	// 	- **ap-southeast-5abc-aliyun**: Zone A+B+C in the Indonesia (Jakarta) region.
	//
	// 	- **cn-hangzhou-ehi-aliyun**: Zone E+H+I in the China (Hangzhou) region.
	//
	// 	- **cn-beijing-acd-aliyun**: Zone A+C+D in the China (Beijing) region.
	//
	// 	- **ap-southeast-1-abc-aliyun**: Zone A+B+C in the Singapore region.
	//
	// 	- **cn-zhangjiakou-abc-aliyun**: Zone A+B+C in the China (Zhangjiakou) region.
	//
	// 	- **cn-shanghai-efg-aliyun**: Zone E+F+G in the China (Shanghai) region.
	//
	// 	- **cn-shanghai-abd-aliyun**: Zone A+B+D in the China (Shanghai) region.
	//
	// 	- **cn-hangzhou-bef-aliyun**: Zone B+E+F in the China (Hangzhou) region.
	//
	// 	- **cn-hangzhou-bce-aliyun**: Zone B+C+E in the China (Hangzhou) region.
	//
	// 	- **cn-beijing-fgh-aliyun**: Zone F+G+H in the China (Beijing) region.
	//
	// 	- **cn-shenzhen-abc-aliyun**: Zone A+B+C in the China (Shenzhen) region.
	//
	// **This parameter is required if you want to create a multi-zone instance**.
	//
	// example:
	//
	// cn-shanghai-efg-aliyun
	MultiZoneCombination *string `json:"MultiZoneCombination,omitempty" xml:"MultiZoneCombination,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The billing method of the instance you want to create. Valid values:
	//
	// 	- **PREPAY**: subscription.
	//
	// 	- **POSTPAY**: pay-as-you-go.
	//
	// This parameter is required.
	//
	// example:
	//
	// POSTPAY
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// The period based on which you are charged for the instance. Valid values:
	//
	// 	- **Month**: You are charged for the instance on a monthly basis.
	//
	// 	- **Year**: You are charged for the instance on a yearly basis.
	//
	// > This parameter is available and required when the PayType parameter is set to **PREPAY**.
	//
	// example:
	//
	// Month
	PricingCycle *string `json:"PricingCycle,omitempty" xml:"PricingCycle,omitempty"`
	// The ID of the vSwitch that is specified for the secondary zone of the instance. The vSwitch must be deployed in the zone specified by the StandbyZoneId parameter. **This parameter is required if you want to create a multi-zone instance**.
	//
	// example:
	//
	// vsw-uf6fdqa7c0pipnqzq****
	PrimaryVSwitchId *string `json:"PrimaryVSwitchId,omitempty" xml:"PrimaryVSwitchId,omitempty"`
	// Multi-zone instance, availability zone ID of the primary zone. **This parameter is required if you need to create a multi-zone instance.**
	//
	// example:
	//
	// cn-shanghai-e
	PrimaryZoneId *string `json:"PrimaryZoneId,omitempty" xml:"PrimaryZoneId,omitempty"`
	// The ID of the region in which you want to create the instance. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/426062.html) operation to query the region in which you can create the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the Lindorm instance belongs.
	//
	// example:
	//
	// rg-aek2i6weeb4nfii
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The number of LindormSearch nodes in the instance. Valid values: integers from **0*	- to **60**.
	//
	// example:
	//
	// 2
	SolrNum *int32 `json:"SolrNum,omitempty" xml:"SolrNum,omitempty"`
	// The specification of the LindormSearch nodes in the instance. Valid values:
	//
	// 	- **lindorm.g.xlarge**: Each node has 4 dedicated CPU cores and 16 GB of dedicated memory.
	//
	// 	- **lindorm.g.2xlarge**: Each node has 8 dedicated CPU cores and 32 GB of dedicated memory.
	//
	// 	- **lindorm.g.4xlarge**: Each node has 16 dedicated CPU cores and 64 GB of dedicated memory.
	//
	// 	- **lindorm.g.8xlarge**: Each node has 32 dedicated CPU cores and 128 GB of dedicated memory.
	//
	// example:
	//
	// lindorm.g.xlarge
	SolrSpec *string `json:"SolrSpec,omitempty" xml:"SolrSpec,omitempty"`
	// The ID of the vSwitch that is specified for the secondary zone of the instance. The vSwitch must be deployed in the zone specified by the StandbyZoneId parameter. **This parameter is required if you want to create a multi-zone instance**.
	//
	// example:
	//
	// vsw-2zec0kcn08cgdtr6****
	StandbyVSwitchId *string `json:"StandbyVSwitchId,omitempty" xml:"StandbyVSwitchId,omitempty"`
	// The ID of the secondary zone of the instance. **This parameter is required if you want to create a multi-zone instance**.
	//
	// example:
	//
	// cn-shanghai-f
	StandbyZoneId *string `json:"StandbyZoneId,omitempty" xml:"StandbyZoneId,omitempty"`
	// The number of LindormStream nodes in the instance. Valid values: integers from **0*	- to **60**.
	//
	// example:
	//
	// 2
	StreamNum *int32 `json:"StreamNum,omitempty" xml:"StreamNum,omitempty"`
	// The specification of the LindormStream nodes in the instance. Valid values:
	//
	// 	- **lindorm.g.xlarge**: Each node has 4 dedicated CPU cores and 16 GB of dedicated memory.
	//
	// 	- **lindorm.g.2xlarge**: Each node has 8 dedicated CPU cores and 32 GB of dedicated memory.
	//
	// 	- **lindorm.g.4xlarge**: Each node has 16 dedicated CPU cores and 64 GB of dedicated memory.
	//
	// 	- **lindorm.g.8xlarge**: Each node has 32 dedicated CPU cores and 128 GB of dedicated memory.
	//
	// example:
	//
	// lindorm.g.xlarge
	StreamSpec *string `json:"StreamSpec,omitempty" xml:"StreamSpec,omitempty"`
	// The number of the LindormTSDB nodes in the instance. The valid values of this parameter depend on the value of the PayType parameter.
	//
	// 	- If the PayType parameter is set to **PREPAY**, set this parameter to an integer that ranges from **0*	- to **24**.
	//
	// 	- If the PayType parameter is set to **POSTPAY**, set this parameter to an integer that ranges from **0*	- to **32**.
	//
	// example:
	//
	// 2
	TsdbNum *int32 `json:"TsdbNum,omitempty" xml:"TsdbNum,omitempty"`
	// The specification of the LindormTSDB nodes in the instance. Valid values:
	//
	// 	- **lindorm.g.xlarge**: Each node has 4 dedicated CPU cores and 16 GB of dedicated memory.
	//
	// 	- **lindorm.g.2xlarge**: Each node has 8 dedicated CPU cores and 32 GB of dedicated memory.
	//
	// 	- **lindorm.g.4xlarge**: Each node has 16 dedicated CPU cores and 64 GB of dedicated memory.
	//
	// 	- **lindorm.g.8xlarge**: Each node has 32 dedicated CPU cores and 128 GB of dedicated memory.
	//
	// example:
	//
	// lindorm.g.xlarge
	TsdbSpec *string `json:"TsdbSpec,omitempty" xml:"TsdbSpec,omitempty"`
	// The ID of the VPC in which you want to create the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpc-bp1nme44gek34slfc****
	VPCId *string `json:"VPCId,omitempty" xml:"VPCId,omitempty"`
	// The ID of the vSwitch to which you want the instance to connect.
	//
	// This parameter is required.
	//
	// example:
	//
	// vsw-bp1e7clcw529l773d****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The ID of the zone in which you want to create the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai-f
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s CreateLindormInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateLindormInstanceRequest) GoString() string {
	return s.String()
}

func (s *CreateLindormInstanceRequest) SetArbiterVSwitchId(v string) *CreateLindormInstanceRequest {
	s.ArbiterVSwitchId = &v
	return s
}

func (s *CreateLindormInstanceRequest) SetArbiterZoneId(v string) *CreateLindormInstanceRequest {
	s.ArbiterZoneId = &v
	return s
}

func (s *CreateLindormInstanceRequest) SetArchVersion(v string) *CreateLindormInstanceRequest {
	s.ArchVersion = &v
	return s
}

func (s *CreateLindormInstanceRequest) SetAutoRenewDuration(v string) *CreateLindormInstanceRequest {
	s.AutoRenewDuration = &v
	return s
}

func (s *CreateLindormInstanceRequest) SetAutoRenewal(v bool) *CreateLindormInstanceRequest {
	s.AutoRenewal = &v
	return s
}

func (s *CreateLindormInstanceRequest) SetColdStorage(v int32) *CreateLindormInstanceRequest {
	s.ColdStorage = &v
	return s
}

func (s *CreateLindormInstanceRequest) SetCoreSingleStorage(v int32) *CreateLindormInstanceRequest {
	s.CoreSingleStorage = &v
	return s
}

func (s *CreateLindormInstanceRequest) SetCoreSpec(v string) *CreateLindormInstanceRequest {
	s.CoreSpec = &v
	return s
}

func (s *CreateLindormInstanceRequest) SetDiskCategory(v string) *CreateLindormInstanceRequest {
	s.DiskCategory = &v
	return s
}

func (s *CreateLindormInstanceRequest) SetDuration(v string) *CreateLindormInstanceRequest {
	s.Duration = &v
	return s
}

func (s *CreateLindormInstanceRequest) SetFilestoreNum(v int32) *CreateLindormInstanceRequest {
	s.FilestoreNum = &v
	return s
}

func (s *CreateLindormInstanceRequest) SetFilestoreSpec(v string) *CreateLindormInstanceRequest {
	s.FilestoreSpec = &v
	return s
}

func (s *CreateLindormInstanceRequest) SetInstanceAlias(v string) *CreateLindormInstanceRequest {
	s.InstanceAlias = &v
	return s
}

func (s *CreateLindormInstanceRequest) SetInstanceStorage(v string) *CreateLindormInstanceRequest {
	s.InstanceStorage = &v
	return s
}

func (s *CreateLindormInstanceRequest) SetLindormNum(v int32) *CreateLindormInstanceRequest {
	s.LindormNum = &v
	return s
}

func (s *CreateLindormInstanceRequest) SetLindormSpec(v string) *CreateLindormInstanceRequest {
	s.LindormSpec = &v
	return s
}

func (s *CreateLindormInstanceRequest) SetLogDiskCategory(v string) *CreateLindormInstanceRequest {
	s.LogDiskCategory = &v
	return s
}

func (s *CreateLindormInstanceRequest) SetLogNum(v int32) *CreateLindormInstanceRequest {
	s.LogNum = &v
	return s
}

func (s *CreateLindormInstanceRequest) SetLogSingleStorage(v int32) *CreateLindormInstanceRequest {
	s.LogSingleStorage = &v
	return s
}

func (s *CreateLindormInstanceRequest) SetLogSpec(v string) *CreateLindormInstanceRequest {
	s.LogSpec = &v
	return s
}

func (s *CreateLindormInstanceRequest) SetLtsNum(v string) *CreateLindormInstanceRequest {
	s.LtsNum = &v
	return s
}

func (s *CreateLindormInstanceRequest) SetLtsSpec(v string) *CreateLindormInstanceRequest {
	s.LtsSpec = &v
	return s
}

func (s *CreateLindormInstanceRequest) SetMultiZoneCombination(v string) *CreateLindormInstanceRequest {
	s.MultiZoneCombination = &v
	return s
}

func (s *CreateLindormInstanceRequest) SetOwnerAccount(v string) *CreateLindormInstanceRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateLindormInstanceRequest) SetOwnerId(v int64) *CreateLindormInstanceRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateLindormInstanceRequest) SetPayType(v string) *CreateLindormInstanceRequest {
	s.PayType = &v
	return s
}

func (s *CreateLindormInstanceRequest) SetPricingCycle(v string) *CreateLindormInstanceRequest {
	s.PricingCycle = &v
	return s
}

func (s *CreateLindormInstanceRequest) SetPrimaryVSwitchId(v string) *CreateLindormInstanceRequest {
	s.PrimaryVSwitchId = &v
	return s
}

func (s *CreateLindormInstanceRequest) SetPrimaryZoneId(v string) *CreateLindormInstanceRequest {
	s.PrimaryZoneId = &v
	return s
}

func (s *CreateLindormInstanceRequest) SetRegionId(v string) *CreateLindormInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *CreateLindormInstanceRequest) SetResourceGroupId(v string) *CreateLindormInstanceRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateLindormInstanceRequest) SetResourceOwnerAccount(v string) *CreateLindormInstanceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateLindormInstanceRequest) SetResourceOwnerId(v int64) *CreateLindormInstanceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateLindormInstanceRequest) SetSecurityToken(v string) *CreateLindormInstanceRequest {
	s.SecurityToken = &v
	return s
}

func (s *CreateLindormInstanceRequest) SetSolrNum(v int32) *CreateLindormInstanceRequest {
	s.SolrNum = &v
	return s
}

func (s *CreateLindormInstanceRequest) SetSolrSpec(v string) *CreateLindormInstanceRequest {
	s.SolrSpec = &v
	return s
}

func (s *CreateLindormInstanceRequest) SetStandbyVSwitchId(v string) *CreateLindormInstanceRequest {
	s.StandbyVSwitchId = &v
	return s
}

func (s *CreateLindormInstanceRequest) SetStandbyZoneId(v string) *CreateLindormInstanceRequest {
	s.StandbyZoneId = &v
	return s
}

func (s *CreateLindormInstanceRequest) SetStreamNum(v int32) *CreateLindormInstanceRequest {
	s.StreamNum = &v
	return s
}

func (s *CreateLindormInstanceRequest) SetStreamSpec(v string) *CreateLindormInstanceRequest {
	s.StreamSpec = &v
	return s
}

func (s *CreateLindormInstanceRequest) SetTsdbNum(v int32) *CreateLindormInstanceRequest {
	s.TsdbNum = &v
	return s
}

func (s *CreateLindormInstanceRequest) SetTsdbSpec(v string) *CreateLindormInstanceRequest {
	s.TsdbSpec = &v
	return s
}

func (s *CreateLindormInstanceRequest) SetVPCId(v string) *CreateLindormInstanceRequest {
	s.VPCId = &v
	return s
}

func (s *CreateLindormInstanceRequest) SetVSwitchId(v string) *CreateLindormInstanceRequest {
	s.VSwitchId = &v
	return s
}

func (s *CreateLindormInstanceRequest) SetZoneId(v string) *CreateLindormInstanceRequest {
	s.ZoneId = &v
	return s
}

type CreateLindormInstanceResponseBody struct {
	// The details about the access denial.
	//
	// example:
	//
	// {"AuthAction":"xxx","AuthPrincipalDisplayName":"222","AuthPrincipalOwnerId":"111","AuthPrincipalType":"SubUser",,"NoPermissionType":"ImplicitDeny","PolicyType":"AccountLevelIdentityBasedPolicy","EncodedDiagnosticMessage":"xxxxxx"}
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The ID of the Lindorm instance that is created.
	//
	// example:
	//
	// ld-bp1o3y0yme2i2****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the order.
	//
	// example:
	//
	// 111111111111111
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 93BE8227-3406-4D7A-883D-9A421D42****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateLindormInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateLindormInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateLindormInstanceResponseBody) SetAccessDeniedDetail(v string) *CreateLindormInstanceResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CreateLindormInstanceResponseBody) SetInstanceId(v string) *CreateLindormInstanceResponseBody {
	s.InstanceId = &v
	return s
}

func (s *CreateLindormInstanceResponseBody) SetOrderId(v int64) *CreateLindormInstanceResponseBody {
	s.OrderId = &v
	return s
}

func (s *CreateLindormInstanceResponseBody) SetRequestId(v string) *CreateLindormInstanceResponseBody {
	s.RequestId = &v
	return s
}

type CreateLindormInstanceResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateLindormInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateLindormInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateLindormInstanceResponse) GoString() string {
	return s.String()
}

func (s *CreateLindormInstanceResponse) SetHeaders(v map[string]*string) *CreateLindormInstanceResponse {
	s.Headers = v
	return s
}

func (s *CreateLindormInstanceResponse) SetStatusCode(v int32) *CreateLindormInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateLindormInstanceResponse) SetBody(v *CreateLindormInstanceResponseBody) *CreateLindormInstanceResponse {
	s.Body = v
	return s
}

type CreateLindormV2InstanceRequest struct {
	ArbiterVSwitchId      *string `json:"ArbiterVSwitchId,omitempty" xml:"ArbiterVSwitchId,omitempty"`
	ArbiterZoneId         *string `json:"ArbiterZoneId,omitempty" xml:"ArbiterZoneId,omitempty"`
	ArchVersion           *string `json:"ArchVersion,omitempty" xml:"ArchVersion,omitempty"`
	AutoRenewDuration     *string `json:"AutoRenewDuration,omitempty" xml:"AutoRenewDuration,omitempty"`
	AutoRenewal           *bool   `json:"AutoRenewal,omitempty" xml:"AutoRenewal,omitempty"`
	CapacityStorageSize   *int32  `json:"CapacityStorageSize,omitempty" xml:"CapacityStorageSize,omitempty"`
	CloudStorageSize      *int32  `json:"CloudStorageSize,omitempty" xml:"CloudStorageSize,omitempty"`
	CloudStorageType      *string `json:"CloudStorageType,omitempty" xml:"CloudStorageType,omitempty"`
	ClusterMode           *string `json:"ClusterMode,omitempty" xml:"ClusterMode,omitempty"`
	ClusterPattern        *string `json:"ClusterPattern,omitempty" xml:"ClusterPattern,omitempty"`
	Duration              *int32  `json:"Duration,omitempty" xml:"Duration,omitempty"`
	EnableCapacityStorage *bool   `json:"EnableCapacityStorage,omitempty" xml:"EnableCapacityStorage,omitempty"`
	// This parameter is required.
	EngineList    []*CreateLindormV2InstanceRequestEngineList `json:"EngineList,omitempty" xml:"EngineList,omitempty" type:"Repeated"`
	InstanceAlias *string                                     `json:"InstanceAlias,omitempty" xml:"InstanceAlias,omitempty"`
	OwnerAccount  *string                                     `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId       *int64                                      `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	PayType          *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	PricingCycle     *string `json:"PricingCycle,omitempty" xml:"PricingCycle,omitempty"`
	PrimaryVSwitchId *string `json:"PrimaryVSwitchId,omitempty" xml:"PrimaryVSwitchId,omitempty"`
	PrimaryZoneId    *string `json:"PrimaryZoneId,omitempty" xml:"PrimaryZoneId,omitempty"`
	// This parameter is required.
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	StandbyVSwitchId     *string `json:"StandbyVSwitchId,omitempty" xml:"StandbyVSwitchId,omitempty"`
	StandbyZoneId        *string `json:"StandbyZoneId,omitempty" xml:"StandbyZoneId,omitempty"`
	// This parameter is required.
	VPCId     *string `json:"VPCId,omitempty" xml:"VPCId,omitempty"`
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// This parameter is required.
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s CreateLindormV2InstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateLindormV2InstanceRequest) GoString() string {
	return s.String()
}

func (s *CreateLindormV2InstanceRequest) SetArbiterVSwitchId(v string) *CreateLindormV2InstanceRequest {
	s.ArbiterVSwitchId = &v
	return s
}

func (s *CreateLindormV2InstanceRequest) SetArbiterZoneId(v string) *CreateLindormV2InstanceRequest {
	s.ArbiterZoneId = &v
	return s
}

func (s *CreateLindormV2InstanceRequest) SetArchVersion(v string) *CreateLindormV2InstanceRequest {
	s.ArchVersion = &v
	return s
}

func (s *CreateLindormV2InstanceRequest) SetAutoRenewDuration(v string) *CreateLindormV2InstanceRequest {
	s.AutoRenewDuration = &v
	return s
}

func (s *CreateLindormV2InstanceRequest) SetAutoRenewal(v bool) *CreateLindormV2InstanceRequest {
	s.AutoRenewal = &v
	return s
}

func (s *CreateLindormV2InstanceRequest) SetCapacityStorageSize(v int32) *CreateLindormV2InstanceRequest {
	s.CapacityStorageSize = &v
	return s
}

func (s *CreateLindormV2InstanceRequest) SetCloudStorageSize(v int32) *CreateLindormV2InstanceRequest {
	s.CloudStorageSize = &v
	return s
}

func (s *CreateLindormV2InstanceRequest) SetCloudStorageType(v string) *CreateLindormV2InstanceRequest {
	s.CloudStorageType = &v
	return s
}

func (s *CreateLindormV2InstanceRequest) SetClusterMode(v string) *CreateLindormV2InstanceRequest {
	s.ClusterMode = &v
	return s
}

func (s *CreateLindormV2InstanceRequest) SetClusterPattern(v string) *CreateLindormV2InstanceRequest {
	s.ClusterPattern = &v
	return s
}

func (s *CreateLindormV2InstanceRequest) SetDuration(v int32) *CreateLindormV2InstanceRequest {
	s.Duration = &v
	return s
}

func (s *CreateLindormV2InstanceRequest) SetEnableCapacityStorage(v bool) *CreateLindormV2InstanceRequest {
	s.EnableCapacityStorage = &v
	return s
}

func (s *CreateLindormV2InstanceRequest) SetEngineList(v []*CreateLindormV2InstanceRequestEngineList) *CreateLindormV2InstanceRequest {
	s.EngineList = v
	return s
}

func (s *CreateLindormV2InstanceRequest) SetInstanceAlias(v string) *CreateLindormV2InstanceRequest {
	s.InstanceAlias = &v
	return s
}

func (s *CreateLindormV2InstanceRequest) SetOwnerAccount(v string) *CreateLindormV2InstanceRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateLindormV2InstanceRequest) SetOwnerId(v int64) *CreateLindormV2InstanceRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateLindormV2InstanceRequest) SetPayType(v string) *CreateLindormV2InstanceRequest {
	s.PayType = &v
	return s
}

func (s *CreateLindormV2InstanceRequest) SetPricingCycle(v string) *CreateLindormV2InstanceRequest {
	s.PricingCycle = &v
	return s
}

func (s *CreateLindormV2InstanceRequest) SetPrimaryVSwitchId(v string) *CreateLindormV2InstanceRequest {
	s.PrimaryVSwitchId = &v
	return s
}

func (s *CreateLindormV2InstanceRequest) SetPrimaryZoneId(v string) *CreateLindormV2InstanceRequest {
	s.PrimaryZoneId = &v
	return s
}

func (s *CreateLindormV2InstanceRequest) SetRegionId(v string) *CreateLindormV2InstanceRequest {
	s.RegionId = &v
	return s
}

func (s *CreateLindormV2InstanceRequest) SetResourceGroupId(v string) *CreateLindormV2InstanceRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateLindormV2InstanceRequest) SetResourceOwnerAccount(v string) *CreateLindormV2InstanceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateLindormV2InstanceRequest) SetResourceOwnerId(v int64) *CreateLindormV2InstanceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateLindormV2InstanceRequest) SetSecurityToken(v string) *CreateLindormV2InstanceRequest {
	s.SecurityToken = &v
	return s
}

func (s *CreateLindormV2InstanceRequest) SetStandbyVSwitchId(v string) *CreateLindormV2InstanceRequest {
	s.StandbyVSwitchId = &v
	return s
}

func (s *CreateLindormV2InstanceRequest) SetStandbyZoneId(v string) *CreateLindormV2InstanceRequest {
	s.StandbyZoneId = &v
	return s
}

func (s *CreateLindormV2InstanceRequest) SetVPCId(v string) *CreateLindormV2InstanceRequest {
	s.VPCId = &v
	return s
}

func (s *CreateLindormV2InstanceRequest) SetVSwitchId(v string) *CreateLindormV2InstanceRequest {
	s.VSwitchId = &v
	return s
}

func (s *CreateLindormV2InstanceRequest) SetZoneId(v string) *CreateLindormV2InstanceRequest {
	s.ZoneId = &v
	return s
}

type CreateLindormV2InstanceRequestEngineList struct {
	// This parameter is required.
	EngineType    *string                                                  `json:"EngineType,omitempty" xml:"EngineType,omitempty"`
	NodeGroupList []*CreateLindormV2InstanceRequestEngineListNodeGroupList `json:"NodeGroupList,omitempty" xml:"NodeGroupList,omitempty" type:"Repeated"`
}

func (s CreateLindormV2InstanceRequestEngineList) String() string {
	return tea.Prettify(s)
}

func (s CreateLindormV2InstanceRequestEngineList) GoString() string {
	return s.String()
}

func (s *CreateLindormV2InstanceRequestEngineList) SetEngineType(v string) *CreateLindormV2InstanceRequestEngineList {
	s.EngineType = &v
	return s
}

func (s *CreateLindormV2InstanceRequestEngineList) SetNodeGroupList(v []*CreateLindormV2InstanceRequestEngineListNodeGroupList) *CreateLindormV2InstanceRequestEngineList {
	s.NodeGroupList = v
	return s
}

type CreateLindormV2InstanceRequestEngineListNodeGroupList struct {
	// This parameter is required.
	NodeCount    *int32  `json:"NodeCount,omitempty" xml:"NodeCount,omitempty"`
	NodeDiskSize *int32  `json:"NodeDiskSize,omitempty" xml:"NodeDiskSize,omitempty"`
	NodeDiskType *string `json:"NodeDiskType,omitempty" xml:"NodeDiskType,omitempty"`
	// This parameter is required.
	NodeSpec          *string `json:"NodeSpec,omitempty" xml:"NodeSpec,omitempty"`
	ResourceGroupName *string `json:"ResourceGroupName,omitempty" xml:"ResourceGroupName,omitempty"`
}

func (s CreateLindormV2InstanceRequestEngineListNodeGroupList) String() string {
	return tea.Prettify(s)
}

func (s CreateLindormV2InstanceRequestEngineListNodeGroupList) GoString() string {
	return s.String()
}

func (s *CreateLindormV2InstanceRequestEngineListNodeGroupList) SetNodeCount(v int32) *CreateLindormV2InstanceRequestEngineListNodeGroupList {
	s.NodeCount = &v
	return s
}

func (s *CreateLindormV2InstanceRequestEngineListNodeGroupList) SetNodeDiskSize(v int32) *CreateLindormV2InstanceRequestEngineListNodeGroupList {
	s.NodeDiskSize = &v
	return s
}

func (s *CreateLindormV2InstanceRequestEngineListNodeGroupList) SetNodeDiskType(v string) *CreateLindormV2InstanceRequestEngineListNodeGroupList {
	s.NodeDiskType = &v
	return s
}

func (s *CreateLindormV2InstanceRequestEngineListNodeGroupList) SetNodeSpec(v string) *CreateLindormV2InstanceRequestEngineListNodeGroupList {
	s.NodeSpec = &v
	return s
}

func (s *CreateLindormV2InstanceRequestEngineListNodeGroupList) SetResourceGroupName(v string) *CreateLindormV2InstanceRequestEngineListNodeGroupList {
	s.ResourceGroupName = &v
	return s
}

type CreateLindormV2InstanceResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	InstanceId         *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OrderId            *int64  `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	RequestId          *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateLindormV2InstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateLindormV2InstanceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateLindormV2InstanceResponseBody) SetAccessDeniedDetail(v string) *CreateLindormV2InstanceResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CreateLindormV2InstanceResponseBody) SetInstanceId(v string) *CreateLindormV2InstanceResponseBody {
	s.InstanceId = &v
	return s
}

func (s *CreateLindormV2InstanceResponseBody) SetOrderId(v int64) *CreateLindormV2InstanceResponseBody {
	s.OrderId = &v
	return s
}

func (s *CreateLindormV2InstanceResponseBody) SetRequestId(v string) *CreateLindormV2InstanceResponseBody {
	s.RequestId = &v
	return s
}

type CreateLindormV2InstanceResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateLindormV2InstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateLindormV2InstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateLindormV2InstanceResponse) GoString() string {
	return s.String()
}

func (s *CreateLindormV2InstanceResponse) SetHeaders(v map[string]*string) *CreateLindormV2InstanceResponse {
	s.Headers = v
	return s
}

func (s *CreateLindormV2InstanceResponse) SetStatusCode(v int32) *CreateLindormV2InstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateLindormV2InstanceResponse) SetBody(v *CreateLindormV2InstanceResponseBody) *CreateLindormV2InstanceResponse {
	s.Body = v
	return s
}

type DeleteAutoScalingConfigRequest struct {
	// This parameter is required.
	ConfigId *string `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// This parameter is required.
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DeleteAutoScalingConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteAutoScalingConfigRequest) GoString() string {
	return s.String()
}

func (s *DeleteAutoScalingConfigRequest) SetConfigId(v string) *DeleteAutoScalingConfigRequest {
	s.ConfigId = &v
	return s
}

func (s *DeleteAutoScalingConfigRequest) SetInstanceId(v string) *DeleteAutoScalingConfigRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteAutoScalingConfigRequest) SetOwnerAccount(v string) *DeleteAutoScalingConfigRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteAutoScalingConfigRequest) SetOwnerId(v int64) *DeleteAutoScalingConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteAutoScalingConfigRequest) SetResourceOwnerAccount(v string) *DeleteAutoScalingConfigRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteAutoScalingConfigRequest) SetResourceOwnerId(v int64) *DeleteAutoScalingConfigRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteAutoScalingConfigRequest) SetSecurityToken(v string) *DeleteAutoScalingConfigRequest {
	s.SecurityToken = &v
	return s
}

type DeleteAutoScalingConfigResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteAutoScalingConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteAutoScalingConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAutoScalingConfigResponseBody) SetCode(v string) *DeleteAutoScalingConfigResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteAutoScalingConfigResponseBody) SetHttpStatusCode(v int32) *DeleteAutoScalingConfigResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteAutoScalingConfigResponseBody) SetMessage(v string) *DeleteAutoScalingConfigResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteAutoScalingConfigResponseBody) SetRequestId(v string) *DeleteAutoScalingConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAutoScalingConfigResponseBody) SetSuccess(v bool) *DeleteAutoScalingConfigResponseBody {
	s.Success = &v
	return s
}

type DeleteAutoScalingConfigResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAutoScalingConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAutoScalingConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteAutoScalingConfigResponse) GoString() string {
	return s.String()
}

func (s *DeleteAutoScalingConfigResponse) SetHeaders(v map[string]*string) *DeleteAutoScalingConfigResponse {
	s.Headers = v
	return s
}

func (s *DeleteAutoScalingConfigResponse) SetStatusCode(v int32) *DeleteAutoScalingConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAutoScalingConfigResponse) SetBody(v *DeleteAutoScalingConfigResponseBody) *DeleteAutoScalingConfigResponse {
	s.Body = v
	return s
}

type DeleteAutoScalingRuleRequest struct {
	// This parameter is required.
	ConfigId *string `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// This parameter is required.
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// This parameter is required.
	RuleId        *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DeleteAutoScalingRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteAutoScalingRuleRequest) GoString() string {
	return s.String()
}

func (s *DeleteAutoScalingRuleRequest) SetConfigId(v string) *DeleteAutoScalingRuleRequest {
	s.ConfigId = &v
	return s
}

func (s *DeleteAutoScalingRuleRequest) SetInstanceId(v string) *DeleteAutoScalingRuleRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteAutoScalingRuleRequest) SetOwnerAccount(v string) *DeleteAutoScalingRuleRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteAutoScalingRuleRequest) SetOwnerId(v int64) *DeleteAutoScalingRuleRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteAutoScalingRuleRequest) SetResourceOwnerAccount(v string) *DeleteAutoScalingRuleRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteAutoScalingRuleRequest) SetResourceOwnerId(v int64) *DeleteAutoScalingRuleRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteAutoScalingRuleRequest) SetRuleId(v string) *DeleteAutoScalingRuleRequest {
	s.RuleId = &v
	return s
}

func (s *DeleteAutoScalingRuleRequest) SetSecurityToken(v string) *DeleteAutoScalingRuleRequest {
	s.SecurityToken = &v
	return s
}

type DeleteAutoScalingRuleResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteAutoScalingRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteAutoScalingRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAutoScalingRuleResponseBody) SetCode(v string) *DeleteAutoScalingRuleResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteAutoScalingRuleResponseBody) SetHttpStatusCode(v int32) *DeleteAutoScalingRuleResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteAutoScalingRuleResponseBody) SetMessage(v string) *DeleteAutoScalingRuleResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteAutoScalingRuleResponseBody) SetRequestId(v string) *DeleteAutoScalingRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAutoScalingRuleResponseBody) SetSuccess(v bool) *DeleteAutoScalingRuleResponseBody {
	s.Success = &v
	return s
}

type DeleteAutoScalingRuleResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAutoScalingRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAutoScalingRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteAutoScalingRuleResponse) GoString() string {
	return s.String()
}

func (s *DeleteAutoScalingRuleResponse) SetHeaders(v map[string]*string) *DeleteAutoScalingRuleResponse {
	s.Headers = v
	return s
}

func (s *DeleteAutoScalingRuleResponse) SetStatusCode(v int32) *DeleteAutoScalingRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAutoScalingRuleResponse) SetBody(v *DeleteAutoScalingRuleResponseBody) *DeleteAutoScalingRuleResponse {
	s.Body = v
	return s
}

type DeleteCustomResourceRequest struct {
	// This parameter is required.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	Name                 *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DeleteCustomResourceRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteCustomResourceRequest) GoString() string {
	return s.String()
}

func (s *DeleteCustomResourceRequest) SetInstanceId(v string) *DeleteCustomResourceRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteCustomResourceRequest) SetName(v string) *DeleteCustomResourceRequest {
	s.Name = &v
	return s
}

func (s *DeleteCustomResourceRequest) SetOwnerAccount(v string) *DeleteCustomResourceRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteCustomResourceRequest) SetOwnerId(v int64) *DeleteCustomResourceRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteCustomResourceRequest) SetRegionId(v string) *DeleteCustomResourceRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteCustomResourceRequest) SetResourceOwnerAccount(v string) *DeleteCustomResourceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteCustomResourceRequest) SetResourceOwnerId(v int64) *DeleteCustomResourceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteCustomResourceRequest) SetSecurityToken(v string) *DeleteCustomResourceRequest {
	s.SecurityToken = &v
	return s
}

type DeleteCustomResourceResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteCustomResourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteCustomResourceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCustomResourceResponseBody) SetRequestId(v string) *DeleteCustomResourceResponseBody {
	s.RequestId = &v
	return s
}

type DeleteCustomResourceResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteCustomResourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteCustomResourceResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteCustomResourceResponse) GoString() string {
	return s.String()
}

func (s *DeleteCustomResourceResponse) SetHeaders(v map[string]*string) *DeleteCustomResourceResponse {
	s.Headers = v
	return s
}

func (s *DeleteCustomResourceResponse) SetStatusCode(v int32) *DeleteCustomResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteCustomResourceResponse) SetBody(v *DeleteCustomResourceResponseBody) *DeleteCustomResourceResponse {
	s.Body = v
	return s
}

type DeleteLdpsComputeGroupRequest struct {
	// This parameter is required.
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// This parameter is required.
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DeleteLdpsComputeGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteLdpsComputeGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteLdpsComputeGroupRequest) SetGroupName(v string) *DeleteLdpsComputeGroupRequest {
	s.GroupName = &v
	return s
}

func (s *DeleteLdpsComputeGroupRequest) SetInstanceId(v string) *DeleteLdpsComputeGroupRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteLdpsComputeGroupRequest) SetOwnerAccount(v string) *DeleteLdpsComputeGroupRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteLdpsComputeGroupRequest) SetOwnerId(v int64) *DeleteLdpsComputeGroupRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteLdpsComputeGroupRequest) SetRegionId(v string) *DeleteLdpsComputeGroupRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteLdpsComputeGroupRequest) SetResourceOwnerAccount(v string) *DeleteLdpsComputeGroupRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteLdpsComputeGroupRequest) SetResourceOwnerId(v int64) *DeleteLdpsComputeGroupRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteLdpsComputeGroupRequest) SetSecurityToken(v string) *DeleteLdpsComputeGroupRequest {
	s.SecurityToken = &v
	return s
}

type DeleteLdpsComputeGroupResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	RequestId          *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteLdpsComputeGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteLdpsComputeGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteLdpsComputeGroupResponseBody) SetAccessDeniedDetail(v string) *DeleteLdpsComputeGroupResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *DeleteLdpsComputeGroupResponseBody) SetRequestId(v string) *DeleteLdpsComputeGroupResponseBody {
	s.RequestId = &v
	return s
}

type DeleteLdpsComputeGroupResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteLdpsComputeGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteLdpsComputeGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteLdpsComputeGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteLdpsComputeGroupResponse) SetHeaders(v map[string]*string) *DeleteLdpsComputeGroupResponse {
	s.Headers = v
	return s
}

func (s *DeleteLdpsComputeGroupResponse) SetStatusCode(v int32) *DeleteLdpsComputeGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteLdpsComputeGroupResponse) SetBody(v *DeleteLdpsComputeGroupResponseBody) *DeleteLdpsComputeGroupResponse {
	s.Body = v
	return s
}

type DeployLdpsSemiManagedComponentRequest struct {
	ComponentName *string `json:"ComponentName,omitempty" xml:"ComponentName,omitempty"`
	// This parameter is required.
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DeployLdpsSemiManagedComponentRequest) String() string {
	return tea.Prettify(s)
}

func (s DeployLdpsSemiManagedComponentRequest) GoString() string {
	return s.String()
}

func (s *DeployLdpsSemiManagedComponentRequest) SetComponentName(v string) *DeployLdpsSemiManagedComponentRequest {
	s.ComponentName = &v
	return s
}

func (s *DeployLdpsSemiManagedComponentRequest) SetInstanceId(v string) *DeployLdpsSemiManagedComponentRequest {
	s.InstanceId = &v
	return s
}

func (s *DeployLdpsSemiManagedComponentRequest) SetOwnerAccount(v string) *DeployLdpsSemiManagedComponentRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeployLdpsSemiManagedComponentRequest) SetOwnerId(v int64) *DeployLdpsSemiManagedComponentRequest {
	s.OwnerId = &v
	return s
}

func (s *DeployLdpsSemiManagedComponentRequest) SetRegionId(v string) *DeployLdpsSemiManagedComponentRequest {
	s.RegionId = &v
	return s
}

func (s *DeployLdpsSemiManagedComponentRequest) SetResourceOwnerAccount(v string) *DeployLdpsSemiManagedComponentRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeployLdpsSemiManagedComponentRequest) SetResourceOwnerId(v int64) *DeployLdpsSemiManagedComponentRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeployLdpsSemiManagedComponentRequest) SetSecurityToken(v string) *DeployLdpsSemiManagedComponentRequest {
	s.SecurityToken = &v
	return s
}

type DeployLdpsSemiManagedComponentResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeployLdpsSemiManagedComponentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeployLdpsSemiManagedComponentResponseBody) GoString() string {
	return s.String()
}

func (s *DeployLdpsSemiManagedComponentResponseBody) SetRequestId(v string) *DeployLdpsSemiManagedComponentResponseBody {
	s.RequestId = &v
	return s
}

type DeployLdpsSemiManagedComponentResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeployLdpsSemiManagedComponentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeployLdpsSemiManagedComponentResponse) String() string {
	return tea.Prettify(s)
}

func (s DeployLdpsSemiManagedComponentResponse) GoString() string {
	return s.String()
}

func (s *DeployLdpsSemiManagedComponentResponse) SetHeaders(v map[string]*string) *DeployLdpsSemiManagedComponentResponse {
	s.Headers = v
	return s
}

func (s *DeployLdpsSemiManagedComponentResponse) SetStatusCode(v int32) *DeployLdpsSemiManagedComponentResponse {
	s.StatusCode = &v
	return s
}

func (s *DeployLdpsSemiManagedComponentResponse) SetBody(v *DeployLdpsSemiManagedComponentResponseBody) *DeployLdpsSemiManagedComponentResponse {
	s.Body = v
	return s
}

type DescribeRegionsRequest struct {
	// The display language of the regions in the returned results. Valid values:
	//
	// 	- **zh-CN*	- (default): Chinese.
	//
	// 	- **en-US**: English.
	//
	// example:
	//
	// zh-CN
	AcceptLanguage       *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
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

func (s *DescribeRegionsRequest) SetOwnerAccount(v string) *DescribeRegionsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeRegionsRequest) SetOwnerId(v int64) *DescribeRegionsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeRegionsRequest) SetResourceOwnerAccount(v string) *DescribeRegionsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeRegionsRequest) SetResourceOwnerId(v int64) *DescribeRegionsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeRegionsRequest) SetSecurityToken(v string) *DescribeRegionsRequest {
	s.SecurityToken = &v
	return s
}

type DescribeRegionsResponseBody struct {
	// The regions supported by Lindorm.
	Regions []*DescribeRegionsResponseBodyRegions `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 73F6E6DA-9AE5-5548-9E07-761A554DAF2E
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
	LocalName *string `json:"LocalName,omitempty" xml:"LocalName,omitempty"`
	// The endpoint for the region.
	//
	// example:
	//
	// hitsdb.aliyuncs.com
	RegionEndpoint *string `json:"RegionEndpoint,omitempty" xml:"RegionEndpoint,omitempty"`
	// The ID of the region.
	//
	// example:
	//
	// cn-hangzhou
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

type GetAutoScalingConfigRequest struct {
	// This parameter is required.
	ConfigId *string `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// This parameter is required.
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s GetAutoScalingConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAutoScalingConfigRequest) GoString() string {
	return s.String()
}

func (s *GetAutoScalingConfigRequest) SetConfigId(v string) *GetAutoScalingConfigRequest {
	s.ConfigId = &v
	return s
}

func (s *GetAutoScalingConfigRequest) SetInstanceId(v string) *GetAutoScalingConfigRequest {
	s.InstanceId = &v
	return s
}

func (s *GetAutoScalingConfigRequest) SetOwnerAccount(v string) *GetAutoScalingConfigRequest {
	s.OwnerAccount = &v
	return s
}

func (s *GetAutoScalingConfigRequest) SetOwnerId(v int64) *GetAutoScalingConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *GetAutoScalingConfigRequest) SetResourceOwnerAccount(v string) *GetAutoScalingConfigRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetAutoScalingConfigRequest) SetResourceOwnerId(v int64) *GetAutoScalingConfigRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetAutoScalingConfigRequest) SetSecurityToken(v string) *GetAutoScalingConfigRequest {
	s.SecurityToken = &v
	return s
}

type GetAutoScalingConfigResponseBody struct {
	AccessDeniedDetail *string                               `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	Code               *string                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Data               *GetAutoScalingConfigResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	DynamicCode        *string                               `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	DynamicMessage     *string                               `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	HttpStatusCode     *int32                                `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message            *string                               `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId          *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success            *bool                                 `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetAutoScalingConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAutoScalingConfigResponseBody) GoString() string {
	return s.String()
}

func (s *GetAutoScalingConfigResponseBody) SetAccessDeniedDetail(v string) *GetAutoScalingConfigResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *GetAutoScalingConfigResponseBody) SetCode(v string) *GetAutoScalingConfigResponseBody {
	s.Code = &v
	return s
}

func (s *GetAutoScalingConfigResponseBody) SetData(v *GetAutoScalingConfigResponseBodyData) *GetAutoScalingConfigResponseBody {
	s.Data = v
	return s
}

func (s *GetAutoScalingConfigResponseBody) SetDynamicCode(v string) *GetAutoScalingConfigResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *GetAutoScalingConfigResponseBody) SetDynamicMessage(v string) *GetAutoScalingConfigResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *GetAutoScalingConfigResponseBody) SetHttpStatusCode(v int32) *GetAutoScalingConfigResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetAutoScalingConfigResponseBody) SetMessage(v string) *GetAutoScalingConfigResponseBody {
	s.Message = &v
	return s
}

func (s *GetAutoScalingConfigResponseBody) SetRequestId(v string) *GetAutoScalingConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAutoScalingConfigResponseBody) SetSuccess(v bool) *GetAutoScalingConfigResponseBody {
	s.Success = &v
	return s
}

type GetAutoScalingConfigResponseBodyData struct {
	ConfigId           *string                                              `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	ConfigName         *string                                              `json:"ConfigName,omitempty" xml:"ConfigName,omitempty"`
	EffectiveTimeEnd   *string                                              `json:"EffectiveTimeEnd,omitempty" xml:"EffectiveTimeEnd,omitempty"`
	EffectiveTimeStart *string                                              `json:"EffectiveTimeStart,omitempty" xml:"EffectiveTimeStart,omitempty"`
	Enabled            *bool                                                `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	Engine             *string                                              `json:"Engine,omitempty" xml:"Engine,omitempty"`
	InstanceId         *string                                              `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	NodesMax           *int32                                               `json:"NodesMax,omitempty" xml:"NodesMax,omitempty"`
	NodesMin           *int32                                               `json:"NodesMin,omitempty" xml:"NodesMin,omitempty"`
	ScaleRuleList      []*GetAutoScalingConfigResponseBodyDataScaleRuleList `json:"ScaleRuleList,omitempty" xml:"ScaleRuleList,omitempty" type:"Repeated"`
	ScaleType          *string                                              `json:"ScaleType,omitempty" xml:"ScaleType,omitempty"`
	SpecId             *string                                              `json:"SpecId,omitempty" xml:"SpecId,omitempty"`
}

func (s GetAutoScalingConfigResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetAutoScalingConfigResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetAutoScalingConfigResponseBodyData) SetConfigId(v string) *GetAutoScalingConfigResponseBodyData {
	s.ConfigId = &v
	return s
}

func (s *GetAutoScalingConfigResponseBodyData) SetConfigName(v string) *GetAutoScalingConfigResponseBodyData {
	s.ConfigName = &v
	return s
}

func (s *GetAutoScalingConfigResponseBodyData) SetEffectiveTimeEnd(v string) *GetAutoScalingConfigResponseBodyData {
	s.EffectiveTimeEnd = &v
	return s
}

func (s *GetAutoScalingConfigResponseBodyData) SetEffectiveTimeStart(v string) *GetAutoScalingConfigResponseBodyData {
	s.EffectiveTimeStart = &v
	return s
}

func (s *GetAutoScalingConfigResponseBodyData) SetEnabled(v bool) *GetAutoScalingConfigResponseBodyData {
	s.Enabled = &v
	return s
}

func (s *GetAutoScalingConfigResponseBodyData) SetEngine(v string) *GetAutoScalingConfigResponseBodyData {
	s.Engine = &v
	return s
}

func (s *GetAutoScalingConfigResponseBodyData) SetInstanceId(v string) *GetAutoScalingConfigResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *GetAutoScalingConfigResponseBodyData) SetNodesMax(v int32) *GetAutoScalingConfigResponseBodyData {
	s.NodesMax = &v
	return s
}

func (s *GetAutoScalingConfigResponseBodyData) SetNodesMin(v int32) *GetAutoScalingConfigResponseBodyData {
	s.NodesMin = &v
	return s
}

func (s *GetAutoScalingConfigResponseBodyData) SetScaleRuleList(v []*GetAutoScalingConfigResponseBodyDataScaleRuleList) *GetAutoScalingConfigResponseBodyData {
	s.ScaleRuleList = v
	return s
}

func (s *GetAutoScalingConfigResponseBodyData) SetScaleType(v string) *GetAutoScalingConfigResponseBodyData {
	s.ScaleType = &v
	return s
}

func (s *GetAutoScalingConfigResponseBodyData) SetSpecId(v string) *GetAutoScalingConfigResponseBodyData {
	s.SpecId = &v
	return s
}

type GetAutoScalingConfigResponseBodyDataScaleRuleList struct {
	ConfigId          *string `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	Enabled           *bool   `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	EndTime           *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	InstanceId        *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	ObservationWindow *int32  `json:"ObservationWindow,omitempty" xml:"ObservationWindow,omitempty"`
	OperationType     *string `json:"OperationType,omitempty" xml:"OperationType,omitempty"`
	RuleId            *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	RuleName          *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	RuleType          *string `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
	ScaleInStep       *int32  `json:"ScaleInStep,omitempty" xml:"ScaleInStep,omitempty"`
	ScaleOutStep      *int32  `json:"ScaleOutStep,omitempty" xml:"ScaleOutStep,omitempty"`
	SilenceTime       *int32  `json:"SilenceTime,omitempty" xml:"SilenceTime,omitempty"`
	StartTime         *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	TargetMetric      *string `json:"TargetMetric,omitempty" xml:"TargetMetric,omitempty"`
	TargetNodes       *int32  `json:"TargetNodes,omitempty" xml:"TargetNodes,omitempty"`
	ThresholdLower    *int32  `json:"ThresholdLower,omitempty" xml:"ThresholdLower,omitempty"`
	ThresholdUpper    *int32  `json:"ThresholdUpper,omitempty" xml:"ThresholdUpper,omitempty"`
	TriggerCronExpr   *string `json:"TriggerCronExpr,omitempty" xml:"TriggerCronExpr,omitempty"`
}

func (s GetAutoScalingConfigResponseBodyDataScaleRuleList) String() string {
	return tea.Prettify(s)
}

func (s GetAutoScalingConfigResponseBodyDataScaleRuleList) GoString() string {
	return s.String()
}

func (s *GetAutoScalingConfigResponseBodyDataScaleRuleList) SetConfigId(v string) *GetAutoScalingConfigResponseBodyDataScaleRuleList {
	s.ConfigId = &v
	return s
}

func (s *GetAutoScalingConfigResponseBodyDataScaleRuleList) SetEnabled(v bool) *GetAutoScalingConfigResponseBodyDataScaleRuleList {
	s.Enabled = &v
	return s
}

func (s *GetAutoScalingConfigResponseBodyDataScaleRuleList) SetEndTime(v string) *GetAutoScalingConfigResponseBodyDataScaleRuleList {
	s.EndTime = &v
	return s
}

func (s *GetAutoScalingConfigResponseBodyDataScaleRuleList) SetInstanceId(v string) *GetAutoScalingConfigResponseBodyDataScaleRuleList {
	s.InstanceId = &v
	return s
}

func (s *GetAutoScalingConfigResponseBodyDataScaleRuleList) SetObservationWindow(v int32) *GetAutoScalingConfigResponseBodyDataScaleRuleList {
	s.ObservationWindow = &v
	return s
}

func (s *GetAutoScalingConfigResponseBodyDataScaleRuleList) SetOperationType(v string) *GetAutoScalingConfigResponseBodyDataScaleRuleList {
	s.OperationType = &v
	return s
}

func (s *GetAutoScalingConfigResponseBodyDataScaleRuleList) SetRuleId(v string) *GetAutoScalingConfigResponseBodyDataScaleRuleList {
	s.RuleId = &v
	return s
}

func (s *GetAutoScalingConfigResponseBodyDataScaleRuleList) SetRuleName(v string) *GetAutoScalingConfigResponseBodyDataScaleRuleList {
	s.RuleName = &v
	return s
}

func (s *GetAutoScalingConfigResponseBodyDataScaleRuleList) SetRuleType(v string) *GetAutoScalingConfigResponseBodyDataScaleRuleList {
	s.RuleType = &v
	return s
}

func (s *GetAutoScalingConfigResponseBodyDataScaleRuleList) SetScaleInStep(v int32) *GetAutoScalingConfigResponseBodyDataScaleRuleList {
	s.ScaleInStep = &v
	return s
}

func (s *GetAutoScalingConfigResponseBodyDataScaleRuleList) SetScaleOutStep(v int32) *GetAutoScalingConfigResponseBodyDataScaleRuleList {
	s.ScaleOutStep = &v
	return s
}

func (s *GetAutoScalingConfigResponseBodyDataScaleRuleList) SetSilenceTime(v int32) *GetAutoScalingConfigResponseBodyDataScaleRuleList {
	s.SilenceTime = &v
	return s
}

func (s *GetAutoScalingConfigResponseBodyDataScaleRuleList) SetStartTime(v string) *GetAutoScalingConfigResponseBodyDataScaleRuleList {
	s.StartTime = &v
	return s
}

func (s *GetAutoScalingConfigResponseBodyDataScaleRuleList) SetTargetMetric(v string) *GetAutoScalingConfigResponseBodyDataScaleRuleList {
	s.TargetMetric = &v
	return s
}

func (s *GetAutoScalingConfigResponseBodyDataScaleRuleList) SetTargetNodes(v int32) *GetAutoScalingConfigResponseBodyDataScaleRuleList {
	s.TargetNodes = &v
	return s
}

func (s *GetAutoScalingConfigResponseBodyDataScaleRuleList) SetThresholdLower(v int32) *GetAutoScalingConfigResponseBodyDataScaleRuleList {
	s.ThresholdLower = &v
	return s
}

func (s *GetAutoScalingConfigResponseBodyDataScaleRuleList) SetThresholdUpper(v int32) *GetAutoScalingConfigResponseBodyDataScaleRuleList {
	s.ThresholdUpper = &v
	return s
}

func (s *GetAutoScalingConfigResponseBodyDataScaleRuleList) SetTriggerCronExpr(v string) *GetAutoScalingConfigResponseBodyDataScaleRuleList {
	s.TriggerCronExpr = &v
	return s
}

type GetAutoScalingConfigResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAutoScalingConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAutoScalingConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAutoScalingConfigResponse) GoString() string {
	return s.String()
}

func (s *GetAutoScalingConfigResponse) SetHeaders(v map[string]*string) *GetAutoScalingConfigResponse {
	s.Headers = v
	return s
}

func (s *GetAutoScalingConfigResponse) SetStatusCode(v int32) *GetAutoScalingConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAutoScalingConfigResponse) SetBody(v *GetAutoScalingConfigResponseBody) *GetAutoScalingConfigResponse {
	s.Body = v
	return s
}

type GetAutoScalingRuleRequest struct {
	// This parameter is required.
	ConfigId *string `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// This parameter is required.
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// This parameter is required.
	RuleId        *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s GetAutoScalingRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAutoScalingRuleRequest) GoString() string {
	return s.String()
}

func (s *GetAutoScalingRuleRequest) SetConfigId(v string) *GetAutoScalingRuleRequest {
	s.ConfigId = &v
	return s
}

func (s *GetAutoScalingRuleRequest) SetInstanceId(v string) *GetAutoScalingRuleRequest {
	s.InstanceId = &v
	return s
}

func (s *GetAutoScalingRuleRequest) SetOwnerAccount(v string) *GetAutoScalingRuleRequest {
	s.OwnerAccount = &v
	return s
}

func (s *GetAutoScalingRuleRequest) SetOwnerId(v int64) *GetAutoScalingRuleRequest {
	s.OwnerId = &v
	return s
}

func (s *GetAutoScalingRuleRequest) SetResourceOwnerAccount(v string) *GetAutoScalingRuleRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetAutoScalingRuleRequest) SetResourceOwnerId(v int64) *GetAutoScalingRuleRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetAutoScalingRuleRequest) SetRuleId(v string) *GetAutoScalingRuleRequest {
	s.RuleId = &v
	return s
}

func (s *GetAutoScalingRuleRequest) SetSecurityToken(v string) *GetAutoScalingRuleRequest {
	s.SecurityToken = &v
	return s
}

type GetAutoScalingRuleResponseBody struct {
	AccessDeniedDetail *string                             `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	Code               *string                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Data               *GetAutoScalingRuleResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	DynamicCode        *string                             `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	DynamicMessage     *string                             `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	HttpStatusCode     *int32                              `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message            *string                             `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId          *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success            *bool                               `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetAutoScalingRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAutoScalingRuleResponseBody) GoString() string {
	return s.String()
}

func (s *GetAutoScalingRuleResponseBody) SetAccessDeniedDetail(v string) *GetAutoScalingRuleResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *GetAutoScalingRuleResponseBody) SetCode(v string) *GetAutoScalingRuleResponseBody {
	s.Code = &v
	return s
}

func (s *GetAutoScalingRuleResponseBody) SetData(v *GetAutoScalingRuleResponseBodyData) *GetAutoScalingRuleResponseBody {
	s.Data = v
	return s
}

func (s *GetAutoScalingRuleResponseBody) SetDynamicCode(v string) *GetAutoScalingRuleResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *GetAutoScalingRuleResponseBody) SetDynamicMessage(v string) *GetAutoScalingRuleResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *GetAutoScalingRuleResponseBody) SetHttpStatusCode(v int32) *GetAutoScalingRuleResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetAutoScalingRuleResponseBody) SetMessage(v string) *GetAutoScalingRuleResponseBody {
	s.Message = &v
	return s
}

func (s *GetAutoScalingRuleResponseBody) SetRequestId(v string) *GetAutoScalingRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAutoScalingRuleResponseBody) SetSuccess(v bool) *GetAutoScalingRuleResponseBody {
	s.Success = &v
	return s
}

type GetAutoScalingRuleResponseBodyData struct {
	ConfigId          *string `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	Enabled           *bool   `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	EndTime           *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	InstanceId        *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	ObservationWindow *int32  `json:"ObservationWindow,omitempty" xml:"ObservationWindow,omitempty"`
	OperationType     *string `json:"OperationType,omitempty" xml:"OperationType,omitempty"`
	RuleId            *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	RuleName          *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	RuleType          *string `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
	ScaleInStep       *int32  `json:"ScaleInStep,omitempty" xml:"ScaleInStep,omitempty"`
	ScaleOutStep      *int32  `json:"ScaleOutStep,omitempty" xml:"ScaleOutStep,omitempty"`
	SilenceTime       *int32  `json:"SilenceTime,omitempty" xml:"SilenceTime,omitempty"`
	StartTime         *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	TargetMetric      *string `json:"TargetMetric,omitempty" xml:"TargetMetric,omitempty"`
	TargetNodes       *int32  `json:"TargetNodes,omitempty" xml:"TargetNodes,omitempty"`
	ThresholdLower    *int32  `json:"ThresholdLower,omitempty" xml:"ThresholdLower,omitempty"`
	ThresholdUpper    *int32  `json:"ThresholdUpper,omitempty" xml:"ThresholdUpper,omitempty"`
	TriggerCronExpr   *string `json:"TriggerCronExpr,omitempty" xml:"TriggerCronExpr,omitempty"`
}

func (s GetAutoScalingRuleResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetAutoScalingRuleResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetAutoScalingRuleResponseBodyData) SetConfigId(v string) *GetAutoScalingRuleResponseBodyData {
	s.ConfigId = &v
	return s
}

func (s *GetAutoScalingRuleResponseBodyData) SetEnabled(v bool) *GetAutoScalingRuleResponseBodyData {
	s.Enabled = &v
	return s
}

func (s *GetAutoScalingRuleResponseBodyData) SetEndTime(v string) *GetAutoScalingRuleResponseBodyData {
	s.EndTime = &v
	return s
}

func (s *GetAutoScalingRuleResponseBodyData) SetInstanceId(v string) *GetAutoScalingRuleResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *GetAutoScalingRuleResponseBodyData) SetObservationWindow(v int32) *GetAutoScalingRuleResponseBodyData {
	s.ObservationWindow = &v
	return s
}

func (s *GetAutoScalingRuleResponseBodyData) SetOperationType(v string) *GetAutoScalingRuleResponseBodyData {
	s.OperationType = &v
	return s
}

func (s *GetAutoScalingRuleResponseBodyData) SetRuleId(v string) *GetAutoScalingRuleResponseBodyData {
	s.RuleId = &v
	return s
}

func (s *GetAutoScalingRuleResponseBodyData) SetRuleName(v string) *GetAutoScalingRuleResponseBodyData {
	s.RuleName = &v
	return s
}

func (s *GetAutoScalingRuleResponseBodyData) SetRuleType(v string) *GetAutoScalingRuleResponseBodyData {
	s.RuleType = &v
	return s
}

func (s *GetAutoScalingRuleResponseBodyData) SetScaleInStep(v int32) *GetAutoScalingRuleResponseBodyData {
	s.ScaleInStep = &v
	return s
}

func (s *GetAutoScalingRuleResponseBodyData) SetScaleOutStep(v int32) *GetAutoScalingRuleResponseBodyData {
	s.ScaleOutStep = &v
	return s
}

func (s *GetAutoScalingRuleResponseBodyData) SetSilenceTime(v int32) *GetAutoScalingRuleResponseBodyData {
	s.SilenceTime = &v
	return s
}

func (s *GetAutoScalingRuleResponseBodyData) SetStartTime(v string) *GetAutoScalingRuleResponseBodyData {
	s.StartTime = &v
	return s
}

func (s *GetAutoScalingRuleResponseBodyData) SetTargetMetric(v string) *GetAutoScalingRuleResponseBodyData {
	s.TargetMetric = &v
	return s
}

func (s *GetAutoScalingRuleResponseBodyData) SetTargetNodes(v int32) *GetAutoScalingRuleResponseBodyData {
	s.TargetNodes = &v
	return s
}

func (s *GetAutoScalingRuleResponseBodyData) SetThresholdLower(v int32) *GetAutoScalingRuleResponseBodyData {
	s.ThresholdLower = &v
	return s
}

func (s *GetAutoScalingRuleResponseBodyData) SetThresholdUpper(v int32) *GetAutoScalingRuleResponseBodyData {
	s.ThresholdUpper = &v
	return s
}

func (s *GetAutoScalingRuleResponseBodyData) SetTriggerCronExpr(v string) *GetAutoScalingRuleResponseBodyData {
	s.TriggerCronExpr = &v
	return s
}

type GetAutoScalingRuleResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAutoScalingRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAutoScalingRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAutoScalingRuleResponse) GoString() string {
	return s.String()
}

func (s *GetAutoScalingRuleResponse) SetHeaders(v map[string]*string) *GetAutoScalingRuleResponse {
	s.Headers = v
	return s
}

func (s *GetAutoScalingRuleResponse) SetStatusCode(v int32) *GetAutoScalingRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAutoScalingRuleResponse) SetBody(v *GetAutoScalingRuleResponseBody) *GetAutoScalingRuleResponse {
	s.Body = v
	return s
}

type GetClientSourceIpRequest struct {
	// This parameter is required.
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s GetClientSourceIpRequest) String() string {
	return tea.Prettify(s)
}

func (s GetClientSourceIpRequest) GoString() string {
	return s.String()
}

func (s *GetClientSourceIpRequest) SetInstanceId(v string) *GetClientSourceIpRequest {
	s.InstanceId = &v
	return s
}

func (s *GetClientSourceIpRequest) SetOwnerAccount(v string) *GetClientSourceIpRequest {
	s.OwnerAccount = &v
	return s
}

func (s *GetClientSourceIpRequest) SetOwnerId(v int64) *GetClientSourceIpRequest {
	s.OwnerId = &v
	return s
}

func (s *GetClientSourceIpRequest) SetRegionId(v string) *GetClientSourceIpRequest {
	s.RegionId = &v
	return s
}

func (s *GetClientSourceIpRequest) SetResourceOwnerAccount(v string) *GetClientSourceIpRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetClientSourceIpRequest) SetResourceOwnerId(v int64) *GetClientSourceIpRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetClientSourceIpRequest) SetSecurityToken(v string) *GetClientSourceIpRequest {
	s.SecurityToken = &v
	return s
}

type GetClientSourceIpResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	ClientIp           *string `json:"ClientIp,omitempty" xml:"ClientIp,omitempty"`
	RequestId          *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetClientSourceIpResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetClientSourceIpResponseBody) GoString() string {
	return s.String()
}

func (s *GetClientSourceIpResponseBody) SetAccessDeniedDetail(v string) *GetClientSourceIpResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *GetClientSourceIpResponseBody) SetClientIp(v string) *GetClientSourceIpResponseBody {
	s.ClientIp = &v
	return s
}

func (s *GetClientSourceIpResponseBody) SetRequestId(v string) *GetClientSourceIpResponseBody {
	s.RequestId = &v
	return s
}

type GetClientSourceIpResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetClientSourceIpResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetClientSourceIpResponse) String() string {
	return tea.Prettify(s)
}

func (s GetClientSourceIpResponse) GoString() string {
	return s.String()
}

func (s *GetClientSourceIpResponse) SetHeaders(v map[string]*string) *GetClientSourceIpResponse {
	s.Headers = v
	return s
}

func (s *GetClientSourceIpResponse) SetStatusCode(v int32) *GetClientSourceIpResponse {
	s.StatusCode = &v
	return s
}

func (s *GetClientSourceIpResponse) SetBody(v *GetClientSourceIpResponseBody) *GetClientSourceIpResponse {
	s.Body = v
	return s
}

type GetEngineDefaultAuthRequest struct {
	// This parameter is required.
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s GetEngineDefaultAuthRequest) String() string {
	return tea.Prettify(s)
}

func (s GetEngineDefaultAuthRequest) GoString() string {
	return s.String()
}

func (s *GetEngineDefaultAuthRequest) SetInstanceId(v string) *GetEngineDefaultAuthRequest {
	s.InstanceId = &v
	return s
}

func (s *GetEngineDefaultAuthRequest) SetOwnerAccount(v string) *GetEngineDefaultAuthRequest {
	s.OwnerAccount = &v
	return s
}

func (s *GetEngineDefaultAuthRequest) SetOwnerId(v int64) *GetEngineDefaultAuthRequest {
	s.OwnerId = &v
	return s
}

func (s *GetEngineDefaultAuthRequest) SetRegionId(v string) *GetEngineDefaultAuthRequest {
	s.RegionId = &v
	return s
}

func (s *GetEngineDefaultAuthRequest) SetResourceOwnerAccount(v string) *GetEngineDefaultAuthRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetEngineDefaultAuthRequest) SetResourceOwnerId(v int64) *GetEngineDefaultAuthRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetEngineDefaultAuthRequest) SetSecurityToken(v string) *GetEngineDefaultAuthRequest {
	s.SecurityToken = &v
	return s
}

type GetEngineDefaultAuthResponseBody struct {
	AccessDeniedDetail *string                                      `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	AuthInfos          []*GetEngineDefaultAuthResponseBodyAuthInfos `json:"AuthInfos,omitempty" xml:"AuthInfos,omitempty" type:"Repeated"`
	InstanceId         *string                                      `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RequestId          *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetEngineDefaultAuthResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetEngineDefaultAuthResponseBody) GoString() string {
	return s.String()
}

func (s *GetEngineDefaultAuthResponseBody) SetAccessDeniedDetail(v string) *GetEngineDefaultAuthResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *GetEngineDefaultAuthResponseBody) SetAuthInfos(v []*GetEngineDefaultAuthResponseBodyAuthInfos) *GetEngineDefaultAuthResponseBody {
	s.AuthInfos = v
	return s
}

func (s *GetEngineDefaultAuthResponseBody) SetInstanceId(v string) *GetEngineDefaultAuthResponseBody {
	s.InstanceId = &v
	return s
}

func (s *GetEngineDefaultAuthResponseBody) SetRequestId(v string) *GetEngineDefaultAuthResponseBody {
	s.RequestId = &v
	return s
}

type GetEngineDefaultAuthResponseBodyAuthInfos struct {
	Engine   *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
}

func (s GetEngineDefaultAuthResponseBodyAuthInfos) String() string {
	return tea.Prettify(s)
}

func (s GetEngineDefaultAuthResponseBodyAuthInfos) GoString() string {
	return s.String()
}

func (s *GetEngineDefaultAuthResponseBodyAuthInfos) SetEngine(v string) *GetEngineDefaultAuthResponseBodyAuthInfos {
	s.Engine = &v
	return s
}

func (s *GetEngineDefaultAuthResponseBodyAuthInfos) SetPassword(v string) *GetEngineDefaultAuthResponseBodyAuthInfos {
	s.Password = &v
	return s
}

func (s *GetEngineDefaultAuthResponseBodyAuthInfos) SetUsername(v string) *GetEngineDefaultAuthResponseBodyAuthInfos {
	s.Username = &v
	return s
}

type GetEngineDefaultAuthResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetEngineDefaultAuthResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetEngineDefaultAuthResponse) String() string {
	return tea.Prettify(s)
}

func (s GetEngineDefaultAuthResponse) GoString() string {
	return s.String()
}

func (s *GetEngineDefaultAuthResponse) SetHeaders(v map[string]*string) *GetEngineDefaultAuthResponse {
	s.Headers = v
	return s
}

func (s *GetEngineDefaultAuthResponse) SetStatusCode(v int32) *GetEngineDefaultAuthResponse {
	s.StatusCode = &v
	return s
}

func (s *GetEngineDefaultAuthResponse) SetBody(v *GetEngineDefaultAuthResponseBody) *GetEngineDefaultAuthResponse {
	s.Body = v
	return s
}

type GetInstanceIpWhiteListRequest struct {
	// The ID of the instance whose whitelists you want to query. You can call the [GetLindormInstanceList](https://help.aliyun.com/document_detail/426068.html) operation to obtain the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ld-bp1z3506imz2f****
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s GetInstanceIpWhiteListRequest) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceIpWhiteListRequest) GoString() string {
	return s.String()
}

func (s *GetInstanceIpWhiteListRequest) SetInstanceId(v string) *GetInstanceIpWhiteListRequest {
	s.InstanceId = &v
	return s
}

func (s *GetInstanceIpWhiteListRequest) SetOwnerAccount(v string) *GetInstanceIpWhiteListRequest {
	s.OwnerAccount = &v
	return s
}

func (s *GetInstanceIpWhiteListRequest) SetOwnerId(v int64) *GetInstanceIpWhiteListRequest {
	s.OwnerId = &v
	return s
}

func (s *GetInstanceIpWhiteListRequest) SetResourceOwnerAccount(v string) *GetInstanceIpWhiteListRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetInstanceIpWhiteListRequest) SetResourceOwnerId(v int64) *GetInstanceIpWhiteListRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetInstanceIpWhiteListRequest) SetSecurityToken(v string) *GetInstanceIpWhiteListRequest {
	s.SecurityToken = &v
	return s
}

type GetInstanceIpWhiteListResponseBody struct {
	// The details about the access denial.
	//
	// example:
	//
	// {"AuthAction":"xxx","AuthPrincipalDisplayName":"222","AuthPrincipalOwnerId":"111","AuthPrincipalType":"SubUser",,"NoPermissionType":"ImplicitDeny","PolicyType":"AccountLevelIdentityBasedPolicy","EncodedDiagnosticMessage":"xxxxxx"}
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The details about the IP address whitelists.
	GroupList []*GetInstanceIpWhiteListResponseBodyGroupList `json:"GroupList,omitempty" xml:"GroupList,omitempty" type:"Repeated"`
	// The ID of the Lindorm instance.
	//
	// example:
	//
	// ld-bp1z3506imz2f****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The list of IP addresses in the whitelist of the instance.
	IpList []*string `json:"IpList,omitempty" xml:"IpList,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 1D1F6F4D-9203-53E7-84E9-5376B4657E63
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetInstanceIpWhiteListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceIpWhiteListResponseBody) GoString() string {
	return s.String()
}

func (s *GetInstanceIpWhiteListResponseBody) SetAccessDeniedDetail(v string) *GetInstanceIpWhiteListResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *GetInstanceIpWhiteListResponseBody) SetGroupList(v []*GetInstanceIpWhiteListResponseBodyGroupList) *GetInstanceIpWhiteListResponseBody {
	s.GroupList = v
	return s
}

func (s *GetInstanceIpWhiteListResponseBody) SetInstanceId(v string) *GetInstanceIpWhiteListResponseBody {
	s.InstanceId = &v
	return s
}

func (s *GetInstanceIpWhiteListResponseBody) SetIpList(v []*string) *GetInstanceIpWhiteListResponseBody {
	s.IpList = v
	return s
}

func (s *GetInstanceIpWhiteListResponseBody) SetRequestId(v string) *GetInstanceIpWhiteListResponseBody {
	s.RequestId = &v
	return s
}

type GetInstanceIpWhiteListResponseBodyGroupList struct {
	// The name of the IP address whitelist.
	//
	// example:
	//
	// test
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The IP addresses in the whitelist.
	//
	// example:
	//
	// 192.168.1.0/24
	SecurityIpList *string `json:"SecurityIpList,omitempty" xml:"SecurityIpList,omitempty"`
}

func (s GetInstanceIpWhiteListResponseBodyGroupList) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceIpWhiteListResponseBodyGroupList) GoString() string {
	return s.String()
}

func (s *GetInstanceIpWhiteListResponseBodyGroupList) SetGroupName(v string) *GetInstanceIpWhiteListResponseBodyGroupList {
	s.GroupName = &v
	return s
}

func (s *GetInstanceIpWhiteListResponseBodyGroupList) SetSecurityIpList(v string) *GetInstanceIpWhiteListResponseBodyGroupList {
	s.SecurityIpList = &v
	return s
}

type GetInstanceIpWhiteListResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetInstanceIpWhiteListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetInstanceIpWhiteListResponse) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceIpWhiteListResponse) GoString() string {
	return s.String()
}

func (s *GetInstanceIpWhiteListResponse) SetHeaders(v map[string]*string) *GetInstanceIpWhiteListResponse {
	s.Headers = v
	return s
}

func (s *GetInstanceIpWhiteListResponse) SetStatusCode(v int32) *GetInstanceIpWhiteListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetInstanceIpWhiteListResponse) SetBody(v *GetInstanceIpWhiteListResponseBody) *GetInstanceIpWhiteListResponse {
	s.Body = v
	return s
}

type GetInstanceSecurityGroupsRequest struct {
	// This parameter is required.
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s GetInstanceSecurityGroupsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceSecurityGroupsRequest) GoString() string {
	return s.String()
}

func (s *GetInstanceSecurityGroupsRequest) SetInstanceId(v string) *GetInstanceSecurityGroupsRequest {
	s.InstanceId = &v
	return s
}

func (s *GetInstanceSecurityGroupsRequest) SetOwnerAccount(v string) *GetInstanceSecurityGroupsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *GetInstanceSecurityGroupsRequest) SetOwnerId(v int64) *GetInstanceSecurityGroupsRequest {
	s.OwnerId = &v
	return s
}

func (s *GetInstanceSecurityGroupsRequest) SetResourceOwnerAccount(v string) *GetInstanceSecurityGroupsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetInstanceSecurityGroupsRequest) SetResourceOwnerId(v int64) *GetInstanceSecurityGroupsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetInstanceSecurityGroupsRequest) SetSecurityToken(v string) *GetInstanceSecurityGroupsRequest {
	s.SecurityToken = &v
	return s
}

type GetInstanceSecurityGroupsResponseBody struct {
	AccessDeniedDetail *string   `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	InstanceId         *string   `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RequestId          *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SecurityGroups     []*string `json:"SecurityGroups,omitempty" xml:"SecurityGroups,omitempty" type:"Repeated"`
}

func (s GetInstanceSecurityGroupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceSecurityGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *GetInstanceSecurityGroupsResponseBody) SetAccessDeniedDetail(v string) *GetInstanceSecurityGroupsResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *GetInstanceSecurityGroupsResponseBody) SetInstanceId(v string) *GetInstanceSecurityGroupsResponseBody {
	s.InstanceId = &v
	return s
}

func (s *GetInstanceSecurityGroupsResponseBody) SetRequestId(v string) *GetInstanceSecurityGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetInstanceSecurityGroupsResponseBody) SetSecurityGroups(v []*string) *GetInstanceSecurityGroupsResponseBody {
	s.SecurityGroups = v
	return s
}

type GetInstanceSecurityGroupsResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetInstanceSecurityGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetInstanceSecurityGroupsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceSecurityGroupsResponse) GoString() string {
	return s.String()
}

func (s *GetInstanceSecurityGroupsResponse) SetHeaders(v map[string]*string) *GetInstanceSecurityGroupsResponse {
	s.Headers = v
	return s
}

func (s *GetInstanceSecurityGroupsResponse) SetStatusCode(v int32) *GetInstanceSecurityGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetInstanceSecurityGroupsResponse) SetBody(v *GetInstanceSecurityGroupsResponseBody) *GetInstanceSecurityGroupsResponse {
	s.Body = v
	return s
}

type GetLdpsComputeGroupRequest struct {
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// This parameter is required.
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s GetLdpsComputeGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s GetLdpsComputeGroupRequest) GoString() string {
	return s.String()
}

func (s *GetLdpsComputeGroupRequest) SetGroupName(v string) *GetLdpsComputeGroupRequest {
	s.GroupName = &v
	return s
}

func (s *GetLdpsComputeGroupRequest) SetInstanceId(v string) *GetLdpsComputeGroupRequest {
	s.InstanceId = &v
	return s
}

func (s *GetLdpsComputeGroupRequest) SetOwnerAccount(v string) *GetLdpsComputeGroupRequest {
	s.OwnerAccount = &v
	return s
}

func (s *GetLdpsComputeGroupRequest) SetOwnerId(v int64) *GetLdpsComputeGroupRequest {
	s.OwnerId = &v
	return s
}

func (s *GetLdpsComputeGroupRequest) SetRegionId(v string) *GetLdpsComputeGroupRequest {
	s.RegionId = &v
	return s
}

func (s *GetLdpsComputeGroupRequest) SetResourceOwnerAccount(v string) *GetLdpsComputeGroupRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetLdpsComputeGroupRequest) SetResourceOwnerId(v int64) *GetLdpsComputeGroupRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetLdpsComputeGroupRequest) SetSecurityToken(v string) *GetLdpsComputeGroupRequest {
	s.SecurityToken = &v
	return s
}

type GetLdpsComputeGroupResponseBody struct {
	GroupName  *string                `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	Properties map[string]interface{} `json:"Properties,omitempty" xml:"Properties,omitempty"`
	RequestId  *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetLdpsComputeGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetLdpsComputeGroupResponseBody) GoString() string {
	return s.String()
}

func (s *GetLdpsComputeGroupResponseBody) SetGroupName(v string) *GetLdpsComputeGroupResponseBody {
	s.GroupName = &v
	return s
}

func (s *GetLdpsComputeGroupResponseBody) SetProperties(v map[string]interface{}) *GetLdpsComputeGroupResponseBody {
	s.Properties = v
	return s
}

func (s *GetLdpsComputeGroupResponseBody) SetRequestId(v string) *GetLdpsComputeGroupResponseBody {
	s.RequestId = &v
	return s
}

type GetLdpsComputeGroupResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetLdpsComputeGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetLdpsComputeGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s GetLdpsComputeGroupResponse) GoString() string {
	return s.String()
}

func (s *GetLdpsComputeGroupResponse) SetHeaders(v map[string]*string) *GetLdpsComputeGroupResponse {
	s.Headers = v
	return s
}

func (s *GetLdpsComputeGroupResponse) SetStatusCode(v int32) *GetLdpsComputeGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *GetLdpsComputeGroupResponse) SetBody(v *GetLdpsComputeGroupResponseBody) *GetLdpsComputeGroupResponse {
	s.Body = v
	return s
}

type GetLdpsNamespacedQuotaRequest struct {
	// This parameter is required.
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Namespace            *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s GetLdpsNamespacedQuotaRequest) String() string {
	return tea.Prettify(s)
}

func (s GetLdpsNamespacedQuotaRequest) GoString() string {
	return s.String()
}

func (s *GetLdpsNamespacedQuotaRequest) SetInstanceId(v string) *GetLdpsNamespacedQuotaRequest {
	s.InstanceId = &v
	return s
}

func (s *GetLdpsNamespacedQuotaRequest) SetNamespace(v string) *GetLdpsNamespacedQuotaRequest {
	s.Namespace = &v
	return s
}

func (s *GetLdpsNamespacedQuotaRequest) SetOwnerAccount(v string) *GetLdpsNamespacedQuotaRequest {
	s.OwnerAccount = &v
	return s
}

func (s *GetLdpsNamespacedQuotaRequest) SetOwnerId(v int64) *GetLdpsNamespacedQuotaRequest {
	s.OwnerId = &v
	return s
}

func (s *GetLdpsNamespacedQuotaRequest) SetRegionId(v string) *GetLdpsNamespacedQuotaRequest {
	s.RegionId = &v
	return s
}

func (s *GetLdpsNamespacedQuotaRequest) SetResourceOwnerAccount(v string) *GetLdpsNamespacedQuotaRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetLdpsNamespacedQuotaRequest) SetResourceOwnerId(v int64) *GetLdpsNamespacedQuotaRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetLdpsNamespacedQuotaRequest) SetSecurityToken(v string) *GetLdpsNamespacedQuotaRequest {
	s.SecurityToken = &v
	return s
}

type GetLdpsNamespacedQuotaResponseBody struct {
	NamespacedQuotas []*GetLdpsNamespacedQuotaResponseBodyNamespacedQuotas `json:"NamespacedQuotas,omitempty" xml:"NamespacedQuotas,omitempty" type:"Repeated"`
	RequestId        *string                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetLdpsNamespacedQuotaResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetLdpsNamespacedQuotaResponseBody) GoString() string {
	return s.String()
}

func (s *GetLdpsNamespacedQuotaResponseBody) SetNamespacedQuotas(v []*GetLdpsNamespacedQuotaResponseBodyNamespacedQuotas) *GetLdpsNamespacedQuotaResponseBody {
	s.NamespacedQuotas = v
	return s
}

func (s *GetLdpsNamespacedQuotaResponseBody) SetRequestId(v string) *GetLdpsNamespacedQuotaResponseBody {
	s.RequestId = &v
	return s
}

type GetLdpsNamespacedQuotaResponseBodyNamespacedQuotas struct {
	CpuAmount    *string `json:"CpuAmount,omitempty" xml:"CpuAmount,omitempty"`
	MemoryAmount *string `json:"MemoryAmount,omitempty" xml:"MemoryAmount,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	UsedCpu      *string `json:"UsedCpu,omitempty" xml:"UsedCpu,omitempty"`
	UsedMemory   *string `json:"UsedMemory,omitempty" xml:"UsedMemory,omitempty"`
}

func (s GetLdpsNamespacedQuotaResponseBodyNamespacedQuotas) String() string {
	return tea.Prettify(s)
}

func (s GetLdpsNamespacedQuotaResponseBodyNamespacedQuotas) GoString() string {
	return s.String()
}

func (s *GetLdpsNamespacedQuotaResponseBodyNamespacedQuotas) SetCpuAmount(v string) *GetLdpsNamespacedQuotaResponseBodyNamespacedQuotas {
	s.CpuAmount = &v
	return s
}

func (s *GetLdpsNamespacedQuotaResponseBodyNamespacedQuotas) SetMemoryAmount(v string) *GetLdpsNamespacedQuotaResponseBodyNamespacedQuotas {
	s.MemoryAmount = &v
	return s
}

func (s *GetLdpsNamespacedQuotaResponseBodyNamespacedQuotas) SetName(v string) *GetLdpsNamespacedQuotaResponseBodyNamespacedQuotas {
	s.Name = &v
	return s
}

func (s *GetLdpsNamespacedQuotaResponseBodyNamespacedQuotas) SetUsedCpu(v string) *GetLdpsNamespacedQuotaResponseBodyNamespacedQuotas {
	s.UsedCpu = &v
	return s
}

func (s *GetLdpsNamespacedQuotaResponseBodyNamespacedQuotas) SetUsedMemory(v string) *GetLdpsNamespacedQuotaResponseBodyNamespacedQuotas {
	s.UsedMemory = &v
	return s
}

type GetLdpsNamespacedQuotaResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetLdpsNamespacedQuotaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetLdpsNamespacedQuotaResponse) String() string {
	return tea.Prettify(s)
}

func (s GetLdpsNamespacedQuotaResponse) GoString() string {
	return s.String()
}

func (s *GetLdpsNamespacedQuotaResponse) SetHeaders(v map[string]*string) *GetLdpsNamespacedQuotaResponse {
	s.Headers = v
	return s
}

func (s *GetLdpsNamespacedQuotaResponse) SetStatusCode(v int32) *GetLdpsNamespacedQuotaResponse {
	s.StatusCode = &v
	return s
}

func (s *GetLdpsNamespacedQuotaResponse) SetBody(v *GetLdpsNamespacedQuotaResponseBody) *GetLdpsNamespacedQuotaResponse {
	s.Body = v
	return s
}

type GetLdpsResourceCostRequest struct {
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// This parameter is required.
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	JobId                *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	StartTime            *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s GetLdpsResourceCostRequest) String() string {
	return tea.Prettify(s)
}

func (s GetLdpsResourceCostRequest) GoString() string {
	return s.String()
}

func (s *GetLdpsResourceCostRequest) SetEndTime(v int64) *GetLdpsResourceCostRequest {
	s.EndTime = &v
	return s
}

func (s *GetLdpsResourceCostRequest) SetInstanceId(v string) *GetLdpsResourceCostRequest {
	s.InstanceId = &v
	return s
}

func (s *GetLdpsResourceCostRequest) SetJobId(v string) *GetLdpsResourceCostRequest {
	s.JobId = &v
	return s
}

func (s *GetLdpsResourceCostRequest) SetOwnerAccount(v string) *GetLdpsResourceCostRequest {
	s.OwnerAccount = &v
	return s
}

func (s *GetLdpsResourceCostRequest) SetOwnerId(v int64) *GetLdpsResourceCostRequest {
	s.OwnerId = &v
	return s
}

func (s *GetLdpsResourceCostRequest) SetRegionId(v string) *GetLdpsResourceCostRequest {
	s.RegionId = &v
	return s
}

func (s *GetLdpsResourceCostRequest) SetResourceOwnerAccount(v string) *GetLdpsResourceCostRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetLdpsResourceCostRequest) SetResourceOwnerId(v int64) *GetLdpsResourceCostRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetLdpsResourceCostRequest) SetSecurityToken(v string) *GetLdpsResourceCostRequest {
	s.SecurityToken = &v
	return s
}

func (s *GetLdpsResourceCostRequest) SetStartTime(v int64) *GetLdpsResourceCostRequest {
	s.StartTime = &v
	return s
}

type GetLdpsResourceCostResponseBody struct {
	EndTime       *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	InstanceId    *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	JobId         *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StartTime     *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	TotalResource *int64  `json:"TotalResource,omitempty" xml:"TotalResource,omitempty"`
}

func (s GetLdpsResourceCostResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetLdpsResourceCostResponseBody) GoString() string {
	return s.String()
}

func (s *GetLdpsResourceCostResponseBody) SetEndTime(v int64) *GetLdpsResourceCostResponseBody {
	s.EndTime = &v
	return s
}

func (s *GetLdpsResourceCostResponseBody) SetInstanceId(v string) *GetLdpsResourceCostResponseBody {
	s.InstanceId = &v
	return s
}

func (s *GetLdpsResourceCostResponseBody) SetJobId(v string) *GetLdpsResourceCostResponseBody {
	s.JobId = &v
	return s
}

func (s *GetLdpsResourceCostResponseBody) SetRequestId(v string) *GetLdpsResourceCostResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetLdpsResourceCostResponseBody) SetStartTime(v int64) *GetLdpsResourceCostResponseBody {
	s.StartTime = &v
	return s
}

func (s *GetLdpsResourceCostResponseBody) SetTotalResource(v int64) *GetLdpsResourceCostResponseBody {
	s.TotalResource = &v
	return s
}

type GetLdpsResourceCostResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetLdpsResourceCostResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetLdpsResourceCostResponse) String() string {
	return tea.Prettify(s)
}

func (s GetLdpsResourceCostResponse) GoString() string {
	return s.String()
}

func (s *GetLdpsResourceCostResponse) SetHeaders(v map[string]*string) *GetLdpsResourceCostResponse {
	s.Headers = v
	return s
}

func (s *GetLdpsResourceCostResponse) SetStatusCode(v int32) *GetLdpsResourceCostResponse {
	s.StatusCode = &v
	return s
}

func (s *GetLdpsResourceCostResponse) SetBody(v *GetLdpsResourceCostResponseBody) *GetLdpsResourceCostResponse {
	s.Body = v
	return s
}

type GetLindormFsUsedDetailRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ld-xxxx
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s GetLindormFsUsedDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s GetLindormFsUsedDetailRequest) GoString() string {
	return s.String()
}

func (s *GetLindormFsUsedDetailRequest) SetInstanceId(v string) *GetLindormFsUsedDetailRequest {
	s.InstanceId = &v
	return s
}

func (s *GetLindormFsUsedDetailRequest) SetOwnerAccount(v string) *GetLindormFsUsedDetailRequest {
	s.OwnerAccount = &v
	return s
}

func (s *GetLindormFsUsedDetailRequest) SetOwnerId(v int64) *GetLindormFsUsedDetailRequest {
	s.OwnerId = &v
	return s
}

func (s *GetLindormFsUsedDetailRequest) SetRegionId(v string) *GetLindormFsUsedDetailRequest {
	s.RegionId = &v
	return s
}

func (s *GetLindormFsUsedDetailRequest) SetResourceOwnerAccount(v string) *GetLindormFsUsedDetailRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetLindormFsUsedDetailRequest) SetResourceOwnerId(v int64) *GetLindormFsUsedDetailRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetLindormFsUsedDetailRequest) SetSecurityToken(v string) *GetLindormFsUsedDetailRequest {
	s.SecurityToken = &v
	return s
}

type GetLindormFsUsedDetailResponseBody struct {
	// example:
	//
	// {}
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// 85899345920
	FsCapacity *string `json:"FsCapacity,omitempty" xml:"FsCapacity,omitempty"`
	// example:
	//
	// 85899345920
	FsCapacityCold *string `json:"FsCapacityCold,omitempty" xml:"FsCapacityCold,omitempty"`
	// example:
	//
	// 85899345920
	FsCapacityHot *string `json:"FsCapacityHot,omitempty" xml:"FsCapacityHot,omitempty"`
	// example:
	//
	// 33269
	FsUsedCold *string `json:"FsUsedCold,omitempty" xml:"FsUsedCold,omitempty"`
	// example:
	//
	// 33269
	FsUsedColdOnLindormSearch *string `json:"FsUsedColdOnLindormSearch,omitempty" xml:"FsUsedColdOnLindormSearch,omitempty"`
	// example:
	//
	// 33269
	FsUsedColdOnLindormTSDB *string `json:"FsUsedColdOnLindormTSDB,omitempty" xml:"FsUsedColdOnLindormTSDB,omitempty"`
	// example:
	//
	// 33269
	FsUsedColdOnLindormTable *string `json:"FsUsedColdOnLindormTable,omitempty" xml:"FsUsedColdOnLindormTable,omitempty"`
	// example:
	//
	// 33269
	FsUsedHot *string `json:"FsUsedHot,omitempty" xml:"FsUsedHot,omitempty"`
	// example:
	//
	// 33269
	FsUsedHotOnLindormSearch *string `json:"FsUsedHotOnLindormSearch,omitempty" xml:"FsUsedHotOnLindormSearch,omitempty"`
	// example:
	//
	// 33269
	FsUsedHotOnLindormTSDB *string `json:"FsUsedHotOnLindormTSDB,omitempty" xml:"FsUsedHotOnLindormTSDB,omitempty"`
	// example:
	//
	// 33269
	FsUsedHotOnLindormTable *string `json:"FsUsedHotOnLindormTable,omitempty" xml:"FsUsedHotOnLindormTable,omitempty"`
	// example:
	//
	// 33269
	FsUsedOnLindormSearch *string `json:"FsUsedOnLindormSearch,omitempty" xml:"FsUsedOnLindormSearch,omitempty"`
	// example:
	//
	// 33269
	FsUsedOnLindormTSDB *string `json:"FsUsedOnLindormTSDB,omitempty" xml:"FsUsedOnLindormTSDB,omitempty"`
	// example:
	//
	// 33269
	FsUsedOnLindormTable *string `json:"FsUsedOnLindormTable,omitempty" xml:"FsUsedOnLindormTable,omitempty"`
	// example:
	//
	// 33269
	FsUsedOnLindormTableData *string `json:"FsUsedOnLindormTableData,omitempty" xml:"FsUsedOnLindormTableData,omitempty"`
	// example:
	//
	// 33269
	FsUsedOnLindormTableWAL *string                                                `json:"FsUsedOnLindormTableWAL,omitempty" xml:"FsUsedOnLindormTableWAL,omitempty"`
	LStorageUsageList       []*GetLindormFsUsedDetailResponseBodyLStorageUsageList `json:"LStorageUsageList,omitempty" xml:"LStorageUsageList,omitempty" type:"Repeated"`
	// example:
	//
	// 4F23D50C-400C-592C-9486-9D1E10179065
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Valid *string `json:"Valid,omitempty" xml:"Valid,omitempty"`
}

func (s GetLindormFsUsedDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetLindormFsUsedDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GetLindormFsUsedDetailResponseBody) SetAccessDeniedDetail(v string) *GetLindormFsUsedDetailResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *GetLindormFsUsedDetailResponseBody) SetFsCapacity(v string) *GetLindormFsUsedDetailResponseBody {
	s.FsCapacity = &v
	return s
}

func (s *GetLindormFsUsedDetailResponseBody) SetFsCapacityCold(v string) *GetLindormFsUsedDetailResponseBody {
	s.FsCapacityCold = &v
	return s
}

func (s *GetLindormFsUsedDetailResponseBody) SetFsCapacityHot(v string) *GetLindormFsUsedDetailResponseBody {
	s.FsCapacityHot = &v
	return s
}

func (s *GetLindormFsUsedDetailResponseBody) SetFsUsedCold(v string) *GetLindormFsUsedDetailResponseBody {
	s.FsUsedCold = &v
	return s
}

func (s *GetLindormFsUsedDetailResponseBody) SetFsUsedColdOnLindormSearch(v string) *GetLindormFsUsedDetailResponseBody {
	s.FsUsedColdOnLindormSearch = &v
	return s
}

func (s *GetLindormFsUsedDetailResponseBody) SetFsUsedColdOnLindormTSDB(v string) *GetLindormFsUsedDetailResponseBody {
	s.FsUsedColdOnLindormTSDB = &v
	return s
}

func (s *GetLindormFsUsedDetailResponseBody) SetFsUsedColdOnLindormTable(v string) *GetLindormFsUsedDetailResponseBody {
	s.FsUsedColdOnLindormTable = &v
	return s
}

func (s *GetLindormFsUsedDetailResponseBody) SetFsUsedHot(v string) *GetLindormFsUsedDetailResponseBody {
	s.FsUsedHot = &v
	return s
}

func (s *GetLindormFsUsedDetailResponseBody) SetFsUsedHotOnLindormSearch(v string) *GetLindormFsUsedDetailResponseBody {
	s.FsUsedHotOnLindormSearch = &v
	return s
}

func (s *GetLindormFsUsedDetailResponseBody) SetFsUsedHotOnLindormTSDB(v string) *GetLindormFsUsedDetailResponseBody {
	s.FsUsedHotOnLindormTSDB = &v
	return s
}

func (s *GetLindormFsUsedDetailResponseBody) SetFsUsedHotOnLindormTable(v string) *GetLindormFsUsedDetailResponseBody {
	s.FsUsedHotOnLindormTable = &v
	return s
}

func (s *GetLindormFsUsedDetailResponseBody) SetFsUsedOnLindormSearch(v string) *GetLindormFsUsedDetailResponseBody {
	s.FsUsedOnLindormSearch = &v
	return s
}

func (s *GetLindormFsUsedDetailResponseBody) SetFsUsedOnLindormTSDB(v string) *GetLindormFsUsedDetailResponseBody {
	s.FsUsedOnLindormTSDB = &v
	return s
}

func (s *GetLindormFsUsedDetailResponseBody) SetFsUsedOnLindormTable(v string) *GetLindormFsUsedDetailResponseBody {
	s.FsUsedOnLindormTable = &v
	return s
}

func (s *GetLindormFsUsedDetailResponseBody) SetFsUsedOnLindormTableData(v string) *GetLindormFsUsedDetailResponseBody {
	s.FsUsedOnLindormTableData = &v
	return s
}

func (s *GetLindormFsUsedDetailResponseBody) SetFsUsedOnLindormTableWAL(v string) *GetLindormFsUsedDetailResponseBody {
	s.FsUsedOnLindormTableWAL = &v
	return s
}

func (s *GetLindormFsUsedDetailResponseBody) SetLStorageUsageList(v []*GetLindormFsUsedDetailResponseBodyLStorageUsageList) *GetLindormFsUsedDetailResponseBody {
	s.LStorageUsageList = v
	return s
}

func (s *GetLindormFsUsedDetailResponseBody) SetRequestId(v string) *GetLindormFsUsedDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetLindormFsUsedDetailResponseBody) SetValid(v string) *GetLindormFsUsedDetailResponseBody {
	s.Valid = &v
	return s
}

type GetLindormFsUsedDetailResponseBodyLStorageUsageList struct {
	// example:
	//
	// 85899345920
	Capacity *string `json:"Capacity,omitempty" xml:"Capacity,omitempty"`
	// example:
	//
	// StandardCloudStorage
	DiskType *string `json:"DiskType,omitempty" xml:"DiskType,omitempty"`
	// example:
	//
	// 33269
	Used *string `json:"Used,omitempty" xml:"Used,omitempty"`
	// example:
	//
	// 33269
	UsedLindormSearch *string `json:"UsedLindormSearch,omitempty" xml:"UsedLindormSearch,omitempty"`
	// example:
	//
	// 33269
	UsedLindormSpark *string `json:"UsedLindormSpark,omitempty" xml:"UsedLindormSpark,omitempty"`
	// example:
	//
	// 33269
	UsedLindormTable *string `json:"UsedLindormTable,omitempty" xml:"UsedLindormTable,omitempty"`
	// example:
	//
	// 33269
	UsedLindormTsdb *string `json:"UsedLindormTsdb,omitempty" xml:"UsedLindormTsdb,omitempty"`
	// example:
	//
	// 33269
	UsedOther *string `json:"UsedOther,omitempty" xml:"UsedOther,omitempty"`
}

func (s GetLindormFsUsedDetailResponseBodyLStorageUsageList) String() string {
	return tea.Prettify(s)
}

func (s GetLindormFsUsedDetailResponseBodyLStorageUsageList) GoString() string {
	return s.String()
}

func (s *GetLindormFsUsedDetailResponseBodyLStorageUsageList) SetCapacity(v string) *GetLindormFsUsedDetailResponseBodyLStorageUsageList {
	s.Capacity = &v
	return s
}

func (s *GetLindormFsUsedDetailResponseBodyLStorageUsageList) SetDiskType(v string) *GetLindormFsUsedDetailResponseBodyLStorageUsageList {
	s.DiskType = &v
	return s
}

func (s *GetLindormFsUsedDetailResponseBodyLStorageUsageList) SetUsed(v string) *GetLindormFsUsedDetailResponseBodyLStorageUsageList {
	s.Used = &v
	return s
}

func (s *GetLindormFsUsedDetailResponseBodyLStorageUsageList) SetUsedLindormSearch(v string) *GetLindormFsUsedDetailResponseBodyLStorageUsageList {
	s.UsedLindormSearch = &v
	return s
}

func (s *GetLindormFsUsedDetailResponseBodyLStorageUsageList) SetUsedLindormSpark(v string) *GetLindormFsUsedDetailResponseBodyLStorageUsageList {
	s.UsedLindormSpark = &v
	return s
}

func (s *GetLindormFsUsedDetailResponseBodyLStorageUsageList) SetUsedLindormTable(v string) *GetLindormFsUsedDetailResponseBodyLStorageUsageList {
	s.UsedLindormTable = &v
	return s
}

func (s *GetLindormFsUsedDetailResponseBodyLStorageUsageList) SetUsedLindormTsdb(v string) *GetLindormFsUsedDetailResponseBodyLStorageUsageList {
	s.UsedLindormTsdb = &v
	return s
}

func (s *GetLindormFsUsedDetailResponseBodyLStorageUsageList) SetUsedOther(v string) *GetLindormFsUsedDetailResponseBodyLStorageUsageList {
	s.UsedOther = &v
	return s
}

type GetLindormFsUsedDetailResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetLindormFsUsedDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetLindormFsUsedDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s GetLindormFsUsedDetailResponse) GoString() string {
	return s.String()
}

func (s *GetLindormFsUsedDetailResponse) SetHeaders(v map[string]*string) *GetLindormFsUsedDetailResponse {
	s.Headers = v
	return s
}

func (s *GetLindormFsUsedDetailResponse) SetStatusCode(v int32) *GetLindormFsUsedDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *GetLindormFsUsedDetailResponse) SetBody(v *GetLindormFsUsedDetailResponseBody) *GetLindormFsUsedDetailResponse {
	s.Body = v
	return s
}

type GetLindormInstanceRequest struct {
	// The disk type of the log nodes. This parameter is returned only for multi-zone instances. Valid values:
	//
	// 	- **cloud_efficiency**: The nodes use the Standard type of storage.
	//
	// 	- **cloud_ssd**: The nodes use the Performance type of storage.
	//
	// This parameter is required.
	//
	// example:
	//
	// ld-bp1o3y0yme2i2****
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s GetLindormInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s GetLindormInstanceRequest) GoString() string {
	return s.String()
}

func (s *GetLindormInstanceRequest) SetInstanceId(v string) *GetLindormInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *GetLindormInstanceRequest) SetOwnerAccount(v string) *GetLindormInstanceRequest {
	s.OwnerAccount = &v
	return s
}

func (s *GetLindormInstanceRequest) SetOwnerId(v int64) *GetLindormInstanceRequest {
	s.OwnerId = &v
	return s
}

func (s *GetLindormInstanceRequest) SetResourceOwnerAccount(v string) *GetLindormInstanceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetLindormInstanceRequest) SetResourceOwnerId(v int64) *GetLindormInstanceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetLindormInstanceRequest) SetSecurityToken(v string) *GetLindormInstanceRequest {
	s.SecurityToken = &v
	return s
}

type GetLindormInstanceResponseBody struct {
	// 16-digit AliUid of the Alibaba Cloud primary account (main account).
	//
	// example:
	//
	// 164901546557****
	AliUid *int64 `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	// Multi-AZ instance, coordinating Availability Zone virtual switch ID, which must be located in the Availability Zone corresponding to ArbiterZoneId.
	//
	// example:
	//
	// vsw-uf6664pqjawb87k36****
	ArbiterVSwitchId *string `json:"ArbiterVSwitchId,omitempty" xml:"ArbiterVSwitchId,omitempty"`
	// Multi-zone instance, coordinating Availability Zone ID.
	//
	// example:
	//
	// cn-shanghai-g
	ArbiterZoneId *string `json:"ArbiterZoneId,omitempty" xml:"ArbiterZoneId,omitempty"`
	// The architecture of the instance. Valid values:
	//
	// 	- **1.0**: The instance is deployed in a single zone.
	//
	// 	- **2.0**: The instance is deployed across multiple zones.
	//
	// example:
	//
	// 1.0
	ArchVersion *string `json:"ArchVersion,omitempty" xml:"ArchVersion,omitempty"`
	// The Archive storage size of the instance.
	//
	// example:
	//
	// 0GB
	ArchiveStorage *int32 `json:"ArchiveStorage,omitempty" xml:"ArchiveStorage,omitempty"`
	// Indicates whether auto-renewal is enabled, with the following returns:
	//
	// - **true**: Enabled. - **false**: Disabled.
	//
	// > This parameter is returned when the instance\\"s payment type is prepaid.
	//
	// example:
	//
	// false
	AutoRenew *bool `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	// The Capacity storage size of the instance.
	//
	// example:
	//
	// 0GB
	ColdStorage *int32 `json:"ColdStorage,omitempty" xml:"ColdStorage,omitempty"`
	// The disk type of the core nodes. This parameter is returned only for multi-zone instances. Valid values:
	//
	// 	- **cloud_efficiency**: This instance uses the Standard type of storage.
	//
	// 	- **cloud_ssd**: This instance uses the Performance type of storage.
	//
	// 	- **cloud_essd**: This instance uses ESSDs for storage.
	//
	// 	- **cloud_essd_pl0**: This instance uses PL0 ESSDs for storage.
	//
	// example:
	//
	// cloud_efficiency
	CoreDiskCategory *string `json:"CoreDiskCategory,omitempty" xml:"CoreDiskCategory,omitempty"`
	// Multi-zone instance, number of core nodes.
	//
	// example:
	//
	// 4
	CoreNum *int32 `json:"CoreNum,omitempty" xml:"CoreNum,omitempty"`
	// Multi-zone instance, core single-node disk capacity.
	//
	// example:
	//
	// 400
	CoreSingleStorage *int32 `json:"CoreSingleStorage,omitempty" xml:"CoreSingleStorage,omitempty"`
	// Multi-zone instance, core node specification.
	//
	// example:
	//
	// lindorm.g.xlarge
	CoreSpec *string `json:"CoreSpec,omitempty" xml:"CoreSpec,omitempty"`
	// The timestamp in milliseconds between the instance creation time and 1970-01-01 00:00:00.
	//
	// example:
	//
	// 1627290664000
	CreateMilliseconds *int64 `json:"CreateMilliseconds,omitempty" xml:"CreateMilliseconds,omitempty"`
	// The storage capacity of the disk of a single log node. This parameter is returned only for multi-zone instances.
	//
	// example:
	//
	// 2021-07-26 17:10:26
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// Indicates whether deletion protection is enabled, returning:
	//
	// - **true**: Enabled. - **false**: Disabled.
	//
	// example:
	//
	// false
	DeletionProtection *string `json:"DeletionProtection,omitempty" xml:"DeletionProtection,omitempty"`
	// The storage type of the instance. Valid values:
	//
	// 	- **cloud_efficiency**: This instance uses the Standard type of storage.
	//
	// 	- **cloud_ssd**: This instance uses the Performance type of storage.
	//
	// 	- **cloud_essd**: This instance uses ESSDs for storage.
	//
	// 	- **cloud_essd_pl0**: This instance uses PL0 ESSDs for storage.
	//
	// 	- **capacity_cloud_storage**: This instance uses the Capacity type of storage.
	//
	// 	- **local_ssd_pro**: This instance uses local SSDs for storage.
	//
	// 	- **local_hdd_pro**: This instance uses local HDDs for storage.
	//
	// example:
	//
	// cloud_efficiency
	DiskCategory *string `json:"DiskCategory,omitempty" xml:"DiskCategory,omitempty"`
	// The threshold for disk space.
	//
	// example:
	//
	// 80%
	DiskThreshold *string `json:"DiskThreshold,omitempty" xml:"DiskThreshold,omitempty"`
	// Disk space usage rate.
	//
	// example:
	//
	// 0.0%
	DiskUsage *string `json:"DiskUsage,omitempty" xml:"DiskUsage,omitempty"`
	// Indicates whether LBlob is enabled for the instance. Valid values:
	//
	// true: LBlob is enabled for the instance. false: LBlob is not enabled for the instance.
	//
	// example:
	//
	// true
	EnableBlob *bool `json:"EnableBlob,omitempty" xml:"EnableBlob,omitempty"`
	// Indicates whether the data subscription feature for the instance is enabled. Returns:
	//
	// - **true**: Enabled. - **false**: Not enabled.
	//
	// example:
	//
	// false
	EnableCdc *bool `json:"EnableCdc,omitempty" xml:"EnableCdc,omitempty"`
	// Indicates whether the instance\\"s compute engine is enabled, returning:
	//
	// - **true**: Enabled. - **false**: Not enabled.
	//
	// example:
	//
	// true
	EnableCompute *bool `json:"EnableCompute,omitempty" xml:"EnableCompute,omitempty"`
	// Indicates whether the Key Management Service (KMS) is enabled, returning:
	//
	// - **true**: Enabled. - **false**: Disabled.
	//
	// example:
	//
	// false
	EnableKms *bool `json:"EnableKms,omitempty" xml:"EnableKms,omitempty"`
	// Indicates whether the wide-table engine supports Thrift and CQL protocols. If not supported, the SwitchLProxyService interface can be used to enable or disable.
	//
	// True indicates support
	//
	// False indicates no support
	//
	// example:
	//
	// False
	EnableLProxy *bool `json:"EnableLProxy,omitempty" xml:"EnableLProxy,omitempty"`
	// Indicates whether the LTS engine is activated for the instance. Valid values:
	//
	// 	- **true**: The LTS engine is activated for the instance.
	//
	// 	- **false**: The LTS engine is not activated for the instance.
	//
	// example:
	//
	// true
	EnableLTS *bool `json:"EnableLTS,omitempty" xml:"EnableLTS,omitempty"`
	// Indicates whether LindormTable of the instance supports LindormSQL V3 that is compatible with MySQL. By default, LindormTable of instances that are purchased after October 24, 2023 supports LindormSQL V3. If your instance is purchased before this date and want to enable LindormSQL V3, contact the technical support.
	//
	// 	- True: LindormTable supports LindormSQL V3.
	//
	// 	- False: LindormTable does not support LindormSQL V3.
	//
	// example:
	//
	// True
	EnableLsqlVersionV3 *bool `json:"EnableLsqlVersionV3,omitempty" xml:"EnableLsqlVersionV3,omitempty"`
	// Indicates whether AI control nodes are enabled for the instance.
	//
	// 	- True: AI control nodes are enabled for the instance.
	//
	// 	- False: AI control nodes are not enabled for the instance.
	//
	// example:
	//
	// False
	EnableMLCtrl *bool `json:"EnableMLCtrl,omitempty" xml:"EnableMLCtrl,omitempty"`
	// Indicates whether SSL link encryption is enabled, returning:
	//
	// - **true**: Enabled. - **false**: Disabled.
	//
	// example:
	//
	// false
	EnableSSL *bool `json:"EnableSSL,omitempty" xml:"EnableSSL,omitempty"`
	// Whether to enable the Compute Engine History Server.
	//
	// example:
	//
	// true
	EnableShs *bool `json:"EnableShs,omitempty" xml:"EnableShs,omitempty"`
	// Indicates whether the Transparent Data Encryption (TDE) is enabled, returning:
	//
	// - **true**: Enabled. - **false**: Disabled.
	//
	// example:
	//
	// false
	EnableStoreTDE *bool `json:"EnableStoreTDE,omitempty" xml:"EnableStoreTDE,omitempty"`
	// Indicates whether the instance has the stream engine enabled. Return values:
	//
	// - **true**: Stream engine is enabled. - **false**: Stream engine is not enabled.
	//
	// example:
	//
	// true
	EnableStream *bool `json:"EnableStream,omitempty" xml:"EnableStream,omitempty"`
	// The latest version number of the engine.
	EngineList []*GetLindormInstanceResponseBodyEngineList `json:"EngineList,omitempty" xml:"EngineList,omitempty" type:"Repeated"`
	// Supported engine types, the return value is obtained by performing addition operations on the values of the following engine types.
	//
	// - 1: Search Engine - 2: Time Series Engine - 4: Wide Table Engine - 8: File Engine
	//
	// > For example: If EngineType is 15, where 15 = 8 + 4 + 2 + 1, it indicates that the instance supports Search Engine, Time Series Engine, Wide Table Engine, and File Engine. If EngineType is 6, where 6 = 4 + 2, it signifies that the instance supports Time Series Engine and Wide Table Engine.
	//
	// example:
	//
	// 15
	EngineType *int32 `json:"EngineType,omitempty" xml:"EngineType,omitempty"`
	// Expiration time of the instance, format: **yyyy-MM-dd HH:mm:ss**.
	//
	// > This parameter is only returned when the payment type is pre-paid.
	//
	// example:
	//
	// 2021-08-27 00:00:00
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The millisecond value between the instance expiration time and 1970-01-01 00:00:00.
	//
	// example:
	//
	// 1629993600000
	ExpiredMilliseconds *int64 `json:"ExpiredMilliseconds,omitempty" xml:"ExpiredMilliseconds,omitempty"`
	// Instance name.
	//
	// example:
	//
	// test0726
	InstanceAlias *string `json:"InstanceAlias,omitempty" xml:"InstanceAlias,omitempty"`
	// Instance ID.
	//
	// example:
	//
	// ld-bp1o3y0yme2i2****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The status of the instance. Valid values:
	//
	// 	- **CREATING**: The instance is being created.
	//
	// 	- **ACTIVATION**: The instance is running.
	//
	// 	- **COLD_EXPANDING**: The Capacity storage of the instance is being scaled up.
	//
	// 	- **MINOR_VERSION_TRANSING**: The minor version of the instance is being updated.
	//
	// 	- **RESIZING**: The nodes in the instance are being scaled up.
	//
	// 	- **SHRINKING**: The nodes in the instance are being scaled down.
	//
	// 	- **CLASS_CHANGING**: The specification of the instance is being changed.
	//
	// 	- **SSL_SWITCHING: SSL**: The SSL configurations of the instance are being changed.
	//
	// 	- **CDC_OPENING**: Data subscription is being enabled for the instance.
	//
	// 	- **TRANSFER**: The data of the instance is being transferred.
	//
	// 	- **DATABASE_TRANSFER**: The data of the instance is being transferred to databases.
	//
	// 	- **GUARD_CREATING**: A disaster recovery instance is being created.
	//
	// 	- **BACKUP_RECOVERING**: The data of the instance is being restored from a backup.
	//
	// 	- **DATABASE_IMPORTING**: Data is being imported to the instance.
	//
	// 	- **NET_MODIFYING**: The network configurations of the instance are being changed.
	//
	// 	- **NET_SWITCHING**: The network of the instance is being switched between a virtual private cloud (VPC) and the Internet.
	//
	// 	- **NET_CREATING**: The connection to the instance is being created.
	//
	// 	- **NET_DELETING**: The connection to the instance is being deleted.
	//
	// 	- **DELETING**: The instance is being deleted.
	//
	// 	- **RESTARTING**: The instance is restarting.
	//
	// 	- **LOCKED**: The instance is locked because it expires.
	//
	// example:
	//
	// ACTIVATION
	InstanceStatus *string `json:"InstanceStatus,omitempty" xml:"InstanceStatus,omitempty"`
	// Instance\\"s storage capacity.
	//
	// example:
	//
	// 480
	InstanceStorage *string `json:"InstanceStorage,omitempty" xml:"InstanceStorage,omitempty"`
	// Multi-zone instance, log node disk type, returns:
	//
	// - **cloud_efficiency**：Standard cloud storage. - **cloud_ssd**：Performance cloud storage.
	//
	// example:
	//
	// cloud_ssd
	LogDiskCategory *string `json:"LogDiskCategory,omitempty" xml:"LogDiskCategory,omitempty"`
	// Multi-zone instance, number of log nodes.
	//
	// example:
	//
	// 4
	LogNum *int32 `json:"LogNum,omitempty" xml:"LogNum,omitempty"`
	// The storage capacity of the disk of a single log node. This parameter is returned only for multi-zone instances.
	//
	// example:
	//
	// 400GB
	LogSingleStorage *int32 `json:"LogSingleStorage,omitempty" xml:"LogSingleStorage,omitempty"`
	// Multi-zone instance, log node specification.
	//
	// example:
	//
	// lindorm.sn1.large
	LogSpec *string `json:"LogSpec,omitempty" xml:"LogSpec,omitempty"`
	// Maintainable end time.
	//
	// example:
	//
	// 20:00Z
	MaintainEndTime *string `json:"MaintainEndTime,omitempty" xml:"MaintainEndTime,omitempty"`
	// Maintainable start time.
	//
	// example:
	//
	// 00:00Z
	MaintainStartTime *string `json:"MaintainStartTime,omitempty" xml:"MaintainStartTime,omitempty"`
	// Multi-zone combinations. For support details on zone combinations, please refer to the product page.
	//
	// - **ap-southeast-5abc-aliyun**: Indonesia (Jakarta) A+B+C. - **cn-hangzhou-ehi-aliyun**: East China 1 (Hangzhou) E+H+I. - **cn-beijing-acd-aliyun**: North China 2 (Beijing) A+C+D. - **ap-southeast-1-abc-aliyun**: Singapore A+B+C. - **cn-zhangjiakou-abc-aliyun**: North China 3 (Zhangjiakou) A+B+C. - **cn-shanghai-efg-aliyun**: East China 2 (Shanghai) E+F+G. - **cn-shanghai-abd-aliyun**: East China 2 (Shanghai) A+B+D. - **cn-hangzhou-bef-aliyun**: East China 1 (Hangzhou) B+E+F. - **cn-hangzhou-bce-aliyun**: East China 1 (Hangzhou) B+C+E. - **cn-beijing-fgh-aliyun**: North China 2 (Beijing) F+G+H. - **cn-shenzhen-abc-aliyun**: South China 1 (Shenzhen) A+B+C.
	//
	// example:
	//
	// cn-shanghai-efg-aliyun
	MultiZoneCombination *string `json:"MultiZoneCombination,omitempty" xml:"MultiZoneCombination,omitempty"`
	// Instance\\"s network type.
	//
	// example:
	//
	// vpc
	NetworkType *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	// 400
	//
	// example:
	//
	// POSTPAY
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// Multi-zone instance, the virtual switch ID of the primary availability zone, which must be in the availability zone corresponding to PrimaryZoneId.
	//
	// example:
	//
	// vsw-uf6fdqa7c0pipnqzq****
	PrimaryVSwitchId *string `json:"PrimaryVSwitchId,omitempty" xml:"PrimaryVSwitchId,omitempty"`
	// Multi-zone instance, availability zone ID of the primary zone.
	//
	// example:
	//
	// cn-shanghai-e
	PrimaryZoneId *string `json:"PrimaryZoneId,omitempty" xml:"PrimaryZoneId,omitempty"`
	// Region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Request ID.
	//
	// example:
	//
	// 633F1BE4-C8DA-5744-8FDF-A3075C3FE37F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Resource group ID.
	//
	// example:
	//
	// rg-aek2wvd6oia****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// Instance type, valid values:
	//
	// - **lindorm**：represents a Lindorm single-zone instance. - **lindorm_multizone**：represents a Lindorm multi-zone instance. - **serverless_lindorm**：represents a Lindorm Serverless instance. - **lindorm_standalone**：represents a Lindorm standalone instance. - **lts**：represents the Lindorm Data Channel Service type.
	//
	// example:
	//
	// lindorm
	ServiceType *string `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
	// Multi-zone instance, the virtual switch ID of the backup availability zone, which must be in the availability zone corresponding to StandbyZoneId.
	//
	// example:
	//
	// vsw-2zec0kcn08cgdtr6****
	StandbyVSwitchId *string `json:"StandbyVSwitchId,omitempty" xml:"StandbyVSwitchId,omitempty"`
	// Multi-zone instance, backup availability zone\\"s availability zone ID.
	//
	// example:
	//
	// cn-shanghai-f
	StandbyZoneId *string `json:"StandbyZoneId,omitempty" xml:"StandbyZoneId,omitempty"`
	// The type of the log nodes. This parameter is returned only for multi-zone instances.
	//
	// example:
	//
	// vpc-bp1n3i15v90el48nx****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The number of the log nodes. This parameter is returned only for multi-zone instances.
	//
	// example:
	//
	// vsw-bp1vbjzmod9q3l9eo****
	VswitchId *string `json:"VswitchId,omitempty" xml:"VswitchId,omitempty"`
	// Availability Zone ID.
	//
	// example:
	//
	// cn-hangzhou-h
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s GetLindormInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetLindormInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *GetLindormInstanceResponseBody) SetAliUid(v int64) *GetLindormInstanceResponseBody {
	s.AliUid = &v
	return s
}

func (s *GetLindormInstanceResponseBody) SetArbiterVSwitchId(v string) *GetLindormInstanceResponseBody {
	s.ArbiterVSwitchId = &v
	return s
}

func (s *GetLindormInstanceResponseBody) SetArbiterZoneId(v string) *GetLindormInstanceResponseBody {
	s.ArbiterZoneId = &v
	return s
}

func (s *GetLindormInstanceResponseBody) SetArchVersion(v string) *GetLindormInstanceResponseBody {
	s.ArchVersion = &v
	return s
}

func (s *GetLindormInstanceResponseBody) SetArchiveStorage(v int32) *GetLindormInstanceResponseBody {
	s.ArchiveStorage = &v
	return s
}

func (s *GetLindormInstanceResponseBody) SetAutoRenew(v bool) *GetLindormInstanceResponseBody {
	s.AutoRenew = &v
	return s
}

func (s *GetLindormInstanceResponseBody) SetColdStorage(v int32) *GetLindormInstanceResponseBody {
	s.ColdStorage = &v
	return s
}

func (s *GetLindormInstanceResponseBody) SetCoreDiskCategory(v string) *GetLindormInstanceResponseBody {
	s.CoreDiskCategory = &v
	return s
}

func (s *GetLindormInstanceResponseBody) SetCoreNum(v int32) *GetLindormInstanceResponseBody {
	s.CoreNum = &v
	return s
}

func (s *GetLindormInstanceResponseBody) SetCoreSingleStorage(v int32) *GetLindormInstanceResponseBody {
	s.CoreSingleStorage = &v
	return s
}

func (s *GetLindormInstanceResponseBody) SetCoreSpec(v string) *GetLindormInstanceResponseBody {
	s.CoreSpec = &v
	return s
}

func (s *GetLindormInstanceResponseBody) SetCreateMilliseconds(v int64) *GetLindormInstanceResponseBody {
	s.CreateMilliseconds = &v
	return s
}

func (s *GetLindormInstanceResponseBody) SetCreateTime(v string) *GetLindormInstanceResponseBody {
	s.CreateTime = &v
	return s
}

func (s *GetLindormInstanceResponseBody) SetDeletionProtection(v string) *GetLindormInstanceResponseBody {
	s.DeletionProtection = &v
	return s
}

func (s *GetLindormInstanceResponseBody) SetDiskCategory(v string) *GetLindormInstanceResponseBody {
	s.DiskCategory = &v
	return s
}

func (s *GetLindormInstanceResponseBody) SetDiskThreshold(v string) *GetLindormInstanceResponseBody {
	s.DiskThreshold = &v
	return s
}

func (s *GetLindormInstanceResponseBody) SetDiskUsage(v string) *GetLindormInstanceResponseBody {
	s.DiskUsage = &v
	return s
}

func (s *GetLindormInstanceResponseBody) SetEnableBlob(v bool) *GetLindormInstanceResponseBody {
	s.EnableBlob = &v
	return s
}

func (s *GetLindormInstanceResponseBody) SetEnableCdc(v bool) *GetLindormInstanceResponseBody {
	s.EnableCdc = &v
	return s
}

func (s *GetLindormInstanceResponseBody) SetEnableCompute(v bool) *GetLindormInstanceResponseBody {
	s.EnableCompute = &v
	return s
}

func (s *GetLindormInstanceResponseBody) SetEnableKms(v bool) *GetLindormInstanceResponseBody {
	s.EnableKms = &v
	return s
}

func (s *GetLindormInstanceResponseBody) SetEnableLProxy(v bool) *GetLindormInstanceResponseBody {
	s.EnableLProxy = &v
	return s
}

func (s *GetLindormInstanceResponseBody) SetEnableLTS(v bool) *GetLindormInstanceResponseBody {
	s.EnableLTS = &v
	return s
}

func (s *GetLindormInstanceResponseBody) SetEnableLsqlVersionV3(v bool) *GetLindormInstanceResponseBody {
	s.EnableLsqlVersionV3 = &v
	return s
}

func (s *GetLindormInstanceResponseBody) SetEnableMLCtrl(v bool) *GetLindormInstanceResponseBody {
	s.EnableMLCtrl = &v
	return s
}

func (s *GetLindormInstanceResponseBody) SetEnableSSL(v bool) *GetLindormInstanceResponseBody {
	s.EnableSSL = &v
	return s
}

func (s *GetLindormInstanceResponseBody) SetEnableShs(v bool) *GetLindormInstanceResponseBody {
	s.EnableShs = &v
	return s
}

func (s *GetLindormInstanceResponseBody) SetEnableStoreTDE(v bool) *GetLindormInstanceResponseBody {
	s.EnableStoreTDE = &v
	return s
}

func (s *GetLindormInstanceResponseBody) SetEnableStream(v bool) *GetLindormInstanceResponseBody {
	s.EnableStream = &v
	return s
}

func (s *GetLindormInstanceResponseBody) SetEngineList(v []*GetLindormInstanceResponseBodyEngineList) *GetLindormInstanceResponseBody {
	s.EngineList = v
	return s
}

func (s *GetLindormInstanceResponseBody) SetEngineType(v int32) *GetLindormInstanceResponseBody {
	s.EngineType = &v
	return s
}

func (s *GetLindormInstanceResponseBody) SetExpireTime(v string) *GetLindormInstanceResponseBody {
	s.ExpireTime = &v
	return s
}

func (s *GetLindormInstanceResponseBody) SetExpiredMilliseconds(v int64) *GetLindormInstanceResponseBody {
	s.ExpiredMilliseconds = &v
	return s
}

func (s *GetLindormInstanceResponseBody) SetInstanceAlias(v string) *GetLindormInstanceResponseBody {
	s.InstanceAlias = &v
	return s
}

func (s *GetLindormInstanceResponseBody) SetInstanceId(v string) *GetLindormInstanceResponseBody {
	s.InstanceId = &v
	return s
}

func (s *GetLindormInstanceResponseBody) SetInstanceStatus(v string) *GetLindormInstanceResponseBody {
	s.InstanceStatus = &v
	return s
}

func (s *GetLindormInstanceResponseBody) SetInstanceStorage(v string) *GetLindormInstanceResponseBody {
	s.InstanceStorage = &v
	return s
}

func (s *GetLindormInstanceResponseBody) SetLogDiskCategory(v string) *GetLindormInstanceResponseBody {
	s.LogDiskCategory = &v
	return s
}

func (s *GetLindormInstanceResponseBody) SetLogNum(v int32) *GetLindormInstanceResponseBody {
	s.LogNum = &v
	return s
}

func (s *GetLindormInstanceResponseBody) SetLogSingleStorage(v int32) *GetLindormInstanceResponseBody {
	s.LogSingleStorage = &v
	return s
}

func (s *GetLindormInstanceResponseBody) SetLogSpec(v string) *GetLindormInstanceResponseBody {
	s.LogSpec = &v
	return s
}

func (s *GetLindormInstanceResponseBody) SetMaintainEndTime(v string) *GetLindormInstanceResponseBody {
	s.MaintainEndTime = &v
	return s
}

func (s *GetLindormInstanceResponseBody) SetMaintainStartTime(v string) *GetLindormInstanceResponseBody {
	s.MaintainStartTime = &v
	return s
}

func (s *GetLindormInstanceResponseBody) SetMultiZoneCombination(v string) *GetLindormInstanceResponseBody {
	s.MultiZoneCombination = &v
	return s
}

func (s *GetLindormInstanceResponseBody) SetNetworkType(v string) *GetLindormInstanceResponseBody {
	s.NetworkType = &v
	return s
}

func (s *GetLindormInstanceResponseBody) SetPayType(v string) *GetLindormInstanceResponseBody {
	s.PayType = &v
	return s
}

func (s *GetLindormInstanceResponseBody) SetPrimaryVSwitchId(v string) *GetLindormInstanceResponseBody {
	s.PrimaryVSwitchId = &v
	return s
}

func (s *GetLindormInstanceResponseBody) SetPrimaryZoneId(v string) *GetLindormInstanceResponseBody {
	s.PrimaryZoneId = &v
	return s
}

func (s *GetLindormInstanceResponseBody) SetRegionId(v string) *GetLindormInstanceResponseBody {
	s.RegionId = &v
	return s
}

func (s *GetLindormInstanceResponseBody) SetRequestId(v string) *GetLindormInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetLindormInstanceResponseBody) SetResourceGroupId(v string) *GetLindormInstanceResponseBody {
	s.ResourceGroupId = &v
	return s
}

func (s *GetLindormInstanceResponseBody) SetServiceType(v string) *GetLindormInstanceResponseBody {
	s.ServiceType = &v
	return s
}

func (s *GetLindormInstanceResponseBody) SetStandbyVSwitchId(v string) *GetLindormInstanceResponseBody {
	s.StandbyVSwitchId = &v
	return s
}

func (s *GetLindormInstanceResponseBody) SetStandbyZoneId(v string) *GetLindormInstanceResponseBody {
	s.StandbyZoneId = &v
	return s
}

func (s *GetLindormInstanceResponseBody) SetVpcId(v string) *GetLindormInstanceResponseBody {
	s.VpcId = &v
	return s
}

func (s *GetLindormInstanceResponseBody) SetVswitchId(v string) *GetLindormInstanceResponseBody {
	s.VswitchId = &v
	return s
}

func (s *GetLindormInstanceResponseBody) SetZoneId(v string) *GetLindormInstanceResponseBody {
	s.ZoneId = &v
	return s
}

type GetLindormInstanceResponseBodyEngineList struct {
	// The number of engine nodes.
	//
	// example:
	//
	// 2
	CoreCount *string `json:"CoreCount,omitempty" xml:"CoreCount,omitempty"`
	// The number of CPU cores on the engine node.
	//
	// example:
	//
	// 4
	CpuCount *string `json:"CpuCount,omitempty" xml:"CpuCount,omitempty"`
	// The engine type. Valid values:
	//
	// 	- **lindorm**: LindormTable.
	//
	// 	- **tsdb**: LindormTSDB.
	//
	// 	- **solr**: LindormSearch.
	//
	// 	- **store**: LindormDFS.
	//
	// 	- **bds**: Lindorm Tunnel Service (LTS).
	//
	// 	- **compute**: Lindorm Distributed Processing System (LDPS).
	//
	// example:
	//
	// lindorm
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// Indicates whether the version of the engine is the latest. Valid values:
	//
	// 	- **true**: The version of the engine is the latest.
	//
	// 	- **false**: The version of the engine is not the latest.
	//
	// example:
	//
	// false
	IsLastVersion *bool `json:"IsLastVersion,omitempty" xml:"IsLastVersion,omitempty"`
	// The latest version number of the engine.
	//
	// example:
	//
	// 2.2.19.2
	LatestVersion *string `json:"LatestVersion,omitempty" xml:"LatestVersion,omitempty"`
	// The memory size of the engine nodes.
	//
	// example:
	//
	// 8GB
	MemorySize *string `json:"MemorySize,omitempty" xml:"MemorySize,omitempty"`
	// The version of the engine.
	//
	// example:
	//
	// 2.2.3
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s GetLindormInstanceResponseBodyEngineList) String() string {
	return tea.Prettify(s)
}

func (s GetLindormInstanceResponseBodyEngineList) GoString() string {
	return s.String()
}

func (s *GetLindormInstanceResponseBodyEngineList) SetCoreCount(v string) *GetLindormInstanceResponseBodyEngineList {
	s.CoreCount = &v
	return s
}

func (s *GetLindormInstanceResponseBodyEngineList) SetCpuCount(v string) *GetLindormInstanceResponseBodyEngineList {
	s.CpuCount = &v
	return s
}

func (s *GetLindormInstanceResponseBodyEngineList) SetEngine(v string) *GetLindormInstanceResponseBodyEngineList {
	s.Engine = &v
	return s
}

func (s *GetLindormInstanceResponseBodyEngineList) SetIsLastVersion(v bool) *GetLindormInstanceResponseBodyEngineList {
	s.IsLastVersion = &v
	return s
}

func (s *GetLindormInstanceResponseBodyEngineList) SetLatestVersion(v string) *GetLindormInstanceResponseBodyEngineList {
	s.LatestVersion = &v
	return s
}

func (s *GetLindormInstanceResponseBodyEngineList) SetMemorySize(v string) *GetLindormInstanceResponseBodyEngineList {
	s.MemorySize = &v
	return s
}

func (s *GetLindormInstanceResponseBodyEngineList) SetVersion(v string) *GetLindormInstanceResponseBodyEngineList {
	s.Version = &v
	return s
}

type GetLindormInstanceResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetLindormInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetLindormInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s GetLindormInstanceResponse) GoString() string {
	return s.String()
}

func (s *GetLindormInstanceResponse) SetHeaders(v map[string]*string) *GetLindormInstanceResponse {
	s.Headers = v
	return s
}

func (s *GetLindormInstanceResponse) SetStatusCode(v int32) *GetLindormInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetLindormInstanceResponse) SetBody(v *GetLindormInstanceResponseBody) *GetLindormInstanceResponse {
	s.Body = v
	return s
}

type GetLindormInstanceEngineListRequest struct {
	// Instance ID, which can be obtained by calling the [GetLindormInstanceList](https://help.aliyun.com/document_detail/426069.html) interface.
	//
	// This parameter is required.
	//
	// example:
	//
	// ld-bp1nq34mv3smk****
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// Region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s GetLindormInstanceEngineListRequest) String() string {
	return tea.Prettify(s)
}

func (s GetLindormInstanceEngineListRequest) GoString() string {
	return s.String()
}

func (s *GetLindormInstanceEngineListRequest) SetInstanceId(v string) *GetLindormInstanceEngineListRequest {
	s.InstanceId = &v
	return s
}

func (s *GetLindormInstanceEngineListRequest) SetOwnerAccount(v string) *GetLindormInstanceEngineListRequest {
	s.OwnerAccount = &v
	return s
}

func (s *GetLindormInstanceEngineListRequest) SetOwnerId(v int64) *GetLindormInstanceEngineListRequest {
	s.OwnerId = &v
	return s
}

func (s *GetLindormInstanceEngineListRequest) SetRegionId(v string) *GetLindormInstanceEngineListRequest {
	s.RegionId = &v
	return s
}

func (s *GetLindormInstanceEngineListRequest) SetResourceOwnerAccount(v string) *GetLindormInstanceEngineListRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetLindormInstanceEngineListRequest) SetResourceOwnerId(v int64) *GetLindormInstanceEngineListRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetLindormInstanceEngineListRequest) SetSecurityToken(v string) *GetLindormInstanceEngineListRequest {
	s.SecurityToken = &v
	return s
}

type GetLindormInstanceEngineListResponseBody struct {
	// The details about the access denial.
	//
	// example:
	//
	// {"AuthAction":"xxx","AuthPrincipalDisplayName":"222","AuthPrincipalOwnerId":"111","AuthPrincipalType":"SubUser",,"NoPermissionType":"ImplicitDeny","PolicyType":"AccountLevelIdentityBasedPolicy","EncodedDiagnosticMessage":"xxxxxx"}
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The list of engines that can run on the specified instance.
	EngineList []*GetLindormInstanceEngineListResponseBodyEngineList `json:"EngineList,omitempty" xml:"EngineList,omitempty" type:"Repeated"`
	// Instance ID.
	//
	// example:
	//
	// ld-bp1nq34mv3smk****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Request ID.
	//
	// example:
	//
	// B496BA0E-520C-59FC-BA04-196D8F3B07EF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetLindormInstanceEngineListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetLindormInstanceEngineListResponseBody) GoString() string {
	return s.String()
}

func (s *GetLindormInstanceEngineListResponseBody) SetAccessDeniedDetail(v string) *GetLindormInstanceEngineListResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *GetLindormInstanceEngineListResponseBody) SetEngineList(v []*GetLindormInstanceEngineListResponseBodyEngineList) *GetLindormInstanceEngineListResponseBody {
	s.EngineList = v
	return s
}

func (s *GetLindormInstanceEngineListResponseBody) SetInstanceId(v string) *GetLindormInstanceEngineListResponseBody {
	s.InstanceId = &v
	return s
}

func (s *GetLindormInstanceEngineListResponseBody) SetRequestId(v string) *GetLindormInstanceEngineListResponseBody {
	s.RequestId = &v
	return s
}

type GetLindormInstanceEngineListResponseBodyEngineList struct {
	// The type of engine that can run on the instance. Valid values:
	//
	// 	- **lindorm**: LindormTable.
	//
	// 	- **tsdb**: LindormTSDB.
	//
	// 	- **solr**: LindormSearch.
	//
	// 	- **store**: LindormDFS.
	//
	// example:
	//
	// lindorm
	EngineType *string `json:"EngineType,omitempty" xml:"EngineType,omitempty"`
	// The list of connection information about the engine.
	NetInfoList []*GetLindormInstanceEngineListResponseBodyEngineListNetInfoList `json:"NetInfoList,omitempty" xml:"NetInfoList,omitempty" type:"Repeated"`
}

func (s GetLindormInstanceEngineListResponseBodyEngineList) String() string {
	return tea.Prettify(s)
}

func (s GetLindormInstanceEngineListResponseBodyEngineList) GoString() string {
	return s.String()
}

func (s *GetLindormInstanceEngineListResponseBodyEngineList) SetEngineType(v string) *GetLindormInstanceEngineListResponseBodyEngineList {
	s.EngineType = &v
	return s
}

func (s *GetLindormInstanceEngineListResponseBodyEngineList) SetNetInfoList(v []*GetLindormInstanceEngineListResponseBodyEngineListNetInfoList) *GetLindormInstanceEngineListResponseBodyEngineList {
	s.NetInfoList = v
	return s
}

type GetLindormInstanceEngineListResponseBodyEngineListNetInfoList struct {
	// The method by which the connection information can be used to access LindormTable. Valid values:
	//
	// 	- **0**: The default value. This value can be ignored.
	//
	// 	- **1**: The connection information can be used to access LindormTable by using ApsaraDB for HBase API for Java.
	//
	// 	- **2**: The connection information can be used to access LindormTable by using ApsaraDB for HBase API for a non-Java language.
	//
	// 	- **3**: The connection information can be used to access LindormTable by using the LindormTable endpoint for CQL.
	//
	// 	- **4**: The connection information can be used to access LindormTable by using the LindormTable endpoint for SQL.
	//
	// 	- **5**: The connection information can be used to access Lindorm by using the LindormTable endpoint for Amazon S3.
	//
	// 	- **6**: The connection information can be used to access Lindorm by using the LindormTable endpoint for MySQL.
	//
	// example:
	//
	// 1
	AccessType *int32 `json:"AccessType,omitempty" xml:"AccessType,omitempty"`
	// The endpoint that is used to connect to the engine.
	//
	// example:
	//
	// ld-bp1nq34mv3smk****-proxy-lindorm.lindorm.rds.aliyuncs.com
	ConnectionString *string `json:"ConnectionString,omitempty" xml:"ConnectionString,omitempty"`
	// The network type of the endpoint. Valid values:
	//
	// 	- **0**: Internet
	//
	// 	- **2**: virtual private cloud (VPC)
	//
	// example:
	//
	// 2
	NetType *string `json:"NetType,omitempty" xml:"NetType,omitempty"`
	// The port number used to connect to the engine.
	//
	// example:
	//
	// 30020
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
}

func (s GetLindormInstanceEngineListResponseBodyEngineListNetInfoList) String() string {
	return tea.Prettify(s)
}

func (s GetLindormInstanceEngineListResponseBodyEngineListNetInfoList) GoString() string {
	return s.String()
}

func (s *GetLindormInstanceEngineListResponseBodyEngineListNetInfoList) SetAccessType(v int32) *GetLindormInstanceEngineListResponseBodyEngineListNetInfoList {
	s.AccessType = &v
	return s
}

func (s *GetLindormInstanceEngineListResponseBodyEngineListNetInfoList) SetConnectionString(v string) *GetLindormInstanceEngineListResponseBodyEngineListNetInfoList {
	s.ConnectionString = &v
	return s
}

func (s *GetLindormInstanceEngineListResponseBodyEngineListNetInfoList) SetNetType(v string) *GetLindormInstanceEngineListResponseBodyEngineListNetInfoList {
	s.NetType = &v
	return s
}

func (s *GetLindormInstanceEngineListResponseBodyEngineListNetInfoList) SetPort(v int32) *GetLindormInstanceEngineListResponseBodyEngineListNetInfoList {
	s.Port = &v
	return s
}

type GetLindormInstanceEngineListResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetLindormInstanceEngineListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetLindormInstanceEngineListResponse) String() string {
	return tea.Prettify(s)
}

func (s GetLindormInstanceEngineListResponse) GoString() string {
	return s.String()
}

func (s *GetLindormInstanceEngineListResponse) SetHeaders(v map[string]*string) *GetLindormInstanceEngineListResponse {
	s.Headers = v
	return s
}

func (s *GetLindormInstanceEngineListResponse) SetStatusCode(v int32) *GetLindormInstanceEngineListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetLindormInstanceEngineListResponse) SetBody(v *GetLindormInstanceEngineListResponseBody) *GetLindormInstanceEngineListResponse {
	s.Body = v
	return s
}

type GetLindormInstanceListRequest struct {
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The number of the pages to return,
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of instances to return on each page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The keyword contained in the names of Lindorm instances you want to query. Fuzzy queries based on the keyword is supported.
	//
	// example:
	//
	// test
	QueryStr *string `json:"QueryStr,omitempty" xml:"QueryStr,omitempty"`
	// The ID of the region in which the instances that you want to query is located. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/426062.html) operation to query the most recent region list.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the instance belongs.
	//
	// example:
	//
	// rg-aek3b63arvg27vi
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The series of instances that you want to query. Valid values:
	//
	// 	- **lindorm**: The instance is a single-zone Lindorm instance.
	//
	// 	- **lindorm_multizone**: The instance is a multi-zone Lindorm instance.
	//
	// 	- **serverless_lindorm**: The instance is a Lindorm Serverless instance.
	//
	// 	- **lindorm_standalone**: The instance is a single-node Lindorm instance.
	//
	// 	- **lts**: The instance is an LTS instance.
	//
	// example:
	//
	// lindorm
	ServiceType *string `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
	// The engine supported by the instances that you want to query. The engines are indicated by different numbers:
	//
	// 	- **1**: LindormSearch.
	//
	// 	- **2**: LindormTSDB.
	//
	// 	- **4**: LindormTable.
	//
	// 	- **8**: LindormDFS.
	//
	// > The value of this parameter is the sum of all numbers that indicate the engines supported by the instance. For example, if you set the value of this parameter to 15, which is the sum of 1, 2, 4, and 8, this operation queries instances that support all four engines. If you set the value of this parameter to 6, which is the sum of 2 and 4, this operation queries instances that support LindormTSDB and LindormTable.
	//
	// example:
	//
	// 15
	SupportEngine *int32 `json:"SupportEngine,omitempty" xml:"SupportEngine,omitempty"`
	// The list of tags associated with the specified instances.
	Tag []*GetLindormInstanceListRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s GetLindormInstanceListRequest) String() string {
	return tea.Prettify(s)
}

func (s GetLindormInstanceListRequest) GoString() string {
	return s.String()
}

func (s *GetLindormInstanceListRequest) SetOwnerAccount(v string) *GetLindormInstanceListRequest {
	s.OwnerAccount = &v
	return s
}

func (s *GetLindormInstanceListRequest) SetOwnerId(v int64) *GetLindormInstanceListRequest {
	s.OwnerId = &v
	return s
}

func (s *GetLindormInstanceListRequest) SetPageNumber(v int32) *GetLindormInstanceListRequest {
	s.PageNumber = &v
	return s
}

func (s *GetLindormInstanceListRequest) SetPageSize(v int32) *GetLindormInstanceListRequest {
	s.PageSize = &v
	return s
}

func (s *GetLindormInstanceListRequest) SetQueryStr(v string) *GetLindormInstanceListRequest {
	s.QueryStr = &v
	return s
}

func (s *GetLindormInstanceListRequest) SetRegionId(v string) *GetLindormInstanceListRequest {
	s.RegionId = &v
	return s
}

func (s *GetLindormInstanceListRequest) SetResourceGroupId(v string) *GetLindormInstanceListRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *GetLindormInstanceListRequest) SetResourceOwnerAccount(v string) *GetLindormInstanceListRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetLindormInstanceListRequest) SetResourceOwnerId(v int64) *GetLindormInstanceListRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetLindormInstanceListRequest) SetSecurityToken(v string) *GetLindormInstanceListRequest {
	s.SecurityToken = &v
	return s
}

func (s *GetLindormInstanceListRequest) SetServiceType(v string) *GetLindormInstanceListRequest {
	s.ServiceType = &v
	return s
}

func (s *GetLindormInstanceListRequest) SetSupportEngine(v int32) *GetLindormInstanceListRequest {
	s.SupportEngine = &v
	return s
}

func (s *GetLindormInstanceListRequest) SetTag(v []*GetLindormInstanceListRequestTag) *GetLindormInstanceListRequest {
	s.Tag = v
	return s
}

type GetLindormInstanceListRequestTag struct {
	// The key of tag N of the instances you want to query. You can specify 1 to 20 tag keys.
	//
	// > You can specify the keys of multiple tags. For example, you can specify the key of the first tag in the first key-value pair contained in the value of this parameter and specify the key of the second tag in the second key-value pair.
	//
	// example:
	//
	// test
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of tag N of the instances you want to query. You can specify 1 to 20 tag values.
	//
	// > You can specify the values of multiple tags. For example, you can specify the value of the first tag in the first key-value pair contained in the value of this parameter and specify the value of the second tag in the second key-value pair.
	//
	// example:
	//
	// 2.2.18
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetLindormInstanceListRequestTag) String() string {
	return tea.Prettify(s)
}

func (s GetLindormInstanceListRequestTag) GoString() string {
	return s.String()
}

func (s *GetLindormInstanceListRequestTag) SetKey(v string) *GetLindormInstanceListRequestTag {
	s.Key = &v
	return s
}

func (s *GetLindormInstanceListRequestTag) SetValue(v string) *GetLindormInstanceListRequestTag {
	s.Value = &v
	return s
}

type GetLindormInstanceListResponseBody struct {
	// The list of instance.
	InstanceList []*GetLindormInstanceListResponseBodyInstanceList `json:"InstanceList,omitempty" xml:"InstanceList,omitempty" type:"Repeated"`
	// The number of returned pages.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of instances that are returned on each page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 1CA1FAFD-E8DC-51C2-AA7E-CA6E2D049BA0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of returned instances.
	//
	// example:
	//
	// 1
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s GetLindormInstanceListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetLindormInstanceListResponseBody) GoString() string {
	return s.String()
}

func (s *GetLindormInstanceListResponseBody) SetInstanceList(v []*GetLindormInstanceListResponseBodyInstanceList) *GetLindormInstanceListResponseBody {
	s.InstanceList = v
	return s
}

func (s *GetLindormInstanceListResponseBody) SetPageNumber(v int32) *GetLindormInstanceListResponseBody {
	s.PageNumber = &v
	return s
}

func (s *GetLindormInstanceListResponseBody) SetPageSize(v int32) *GetLindormInstanceListResponseBody {
	s.PageSize = &v
	return s
}

func (s *GetLindormInstanceListResponseBody) SetRequestId(v string) *GetLindormInstanceListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetLindormInstanceListResponseBody) SetTotal(v int32) *GetLindormInstanceListResponseBody {
	s.Total = &v
	return s
}

type GetLindormInstanceListResponseBodyInstanceList struct {
	// The 16-digit AliUid of the Alibaba Cloud account that owns the instance.
	//
	// example:
	//
	// 164901546557****
	AliUid *int64 `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	// The time when the instance is created. This value is a UNIX timestamp that indicates the number of milliseconds that have elapsed since the epoch time January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1631772842000
	CreateMilliseconds *int64 `json:"CreateMilliseconds,omitempty" xml:"CreateMilliseconds,omitempty"`
	// The time when the instance is created.
	//
	// example:
	//
	// 2021-09-16 14:13:13
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// Indicates whether the column storage engine is enabled, returning:
	//
	// - **true**: Enabled. - **false**: Not enabled.
	//
	// example:
	//
	// true
	EnableColumn *bool `json:"EnableColumn,omitempty" xml:"EnableColumn,omitempty"`
	// Indicates whether LDPS is activated for the instance. Valid values:
	//
	// 	- **true**: LDPS is activated for the instance.
	//
	// 	- **false**: LDPS is not activated for the instance.
	//
	// example:
	//
	// true
	EnableCompute *bool `json:"EnableCompute,omitempty" xml:"EnableCompute,omitempty"`
	// Indicates whether the LTS engine is enabled, returning:
	//
	// - **true**: Enabled. - **false**: Not enabled.
	//
	// example:
	//
	// true
	EnableLts *bool `json:"EnableLts,omitempty" xml:"EnableLts,omitempty"`
	// Indicates whether the message engine is enabled, returning:
	//
	// - **true**: Enabled. - **false**: Not enabled.
	//
	// example:
	//
	// true
	EnableMessage *bool `json:"EnableMessage,omitempty" xml:"EnableMessage,omitempty"`
	EnableRow     *bool `json:"EnableRow,omitempty" xml:"EnableRow,omitempty"`
	// Indicates whether the Lindorm streaming engine is activated for the instance. Valid values:
	//
	// 	- **true**: The Lindorm streaming engine is activated for the instance.
	//
	// 	- **false**: The Lindorm streaming engine is not activated for the instance.
	//
	// example:
	//
	// true
	EnableStream *bool `json:"EnableStream,omitempty" xml:"EnableStream,omitempty"`
	// Whether the vector engine is enabled, returns:
	//
	// - **true**: Enabled. - **false**: Not enabled.
	//
	// example:
	//
	// true
	EnableVector *bool `json:"EnableVector,omitempty" xml:"EnableVector,omitempty"`
	// The engine supported by the instance. The engines are indicated by different numbers:
	//
	// 	- **1**: LindormSearch.
	//
	// 	- **2**: LindormTSDB.
	//
	// 	- **4**: LindormTable.
	//
	// 	- **8**: LindormDFS.
	//
	// > The value of this parameter is the sum of all numbers that indicate the engines supported by the instance. For example, if the value of this parameter is 15, which is the sum of 1, 2, 4, and 8, the instance supports all four engines. If the value of this parameter is 6, which is the sum of 2 and 4, the instance supports LindormTSDB and LindormTable.
	//
	// example:
	//
	// 15
	EngineType *string `json:"EngineType,omitempty" xml:"EngineType,omitempty"`
	// The time when the instance expires.
	//
	// > This parameter is returned only if the billing method of the instance is subscription.
	//
	// example:
	//
	// 2022-04-26 00:00:00
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The time when the instance expires. This value is a UNIX timestamp that indicates the number of milliseconds that have elapsed since the epoch time January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1650902400000
	ExpiredMilliseconds *int64 `json:"ExpiredMilliseconds,omitempty" xml:"ExpiredMilliseconds,omitempty"`
	// The name of the VPC.
	//
	// example:
	//
	// test
	InstanceAlias *string `json:"InstanceAlias,omitempty" xml:"InstanceAlias,omitempty"`
	// The ID of the instance
	//
	// example:
	//
	// ld-bp17pwu1541ia****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The status of the instance. Valid values:
	//
	// 	- **CREATING**: The instance is being created.
	//
	// 	- **ACTIVATION**: The instance is running.
	//
	// 	- **COLD_EXPANDING**: The Capacity storage of the instance is being scaled up.
	//
	// 	- **MINOR_VERSION_TRANSING**: The minor version of the instance is being updated.
	//
	// 	- **RESIZING**: The nodes in the instance are being scaled up.
	//
	// 	- **SHRINKING**: The nodes in the instance are being scaled down.
	//
	// 	- **CLASS_CHANGING**: The specification of the instance is being changed.
	//
	// 	- **SSL_SWITCHING: SSL**: The SSL configurations of the instance are being changed.
	//
	// 	- **CDC_OPENING**: Data subscription is being enabled for the instance.
	//
	// 	- **TRANSFER**: The data of the instance is being transferred.
	//
	// 	- **DATABASE_TRANSFER**: The data of the instance is being transferred to databases.
	//
	// 	- **GUARD_CREATING**: A disaster recovery instance is being created.
	//
	// 	- **BACKUP_RECOVERING**: The data of the instance is being restored from a backup.
	//
	// 	- **DATABASE_IMPORTING**: Data is being imported to the instance.
	//
	// 	- **NET_MODIFYING**: The network configurations of the instance are being changed.
	//
	// 	- **NET_SWITCHING**: The network of the instance is being switched between a virtual private cloud (VPC) and the Internet.
	//
	// 	- **NET_CREATING**: The connection to the instance is being created.
	//
	// 	- **NET_DELETING**: The connection to the instance is being deleted.
	//
	// 	- **DELETING**: The instance is being deleted.
	//
	// 	- **RESTARTING**: The instance is restarting.
	//
	// 	- **LOCKED**: The instance is locked because it expires.
	//
	// example:
	//
	// ACTIVATION
	InstanceStatus *string `json:"InstanceStatus,omitempty" xml:"InstanceStatus,omitempty"`
	// The storage capacity of the instance.
	//
	// example:
	//
	// 960
	InstanceStorage *string `json:"InstanceStorage,omitempty" xml:"InstanceStorage,omitempty"`
	// The network type of the instance.
	//
	// example:
	//
	// vpc
	NetworkType *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	// The billing method of the instance. Valid values:
	//
	// 	- **PREPAY**: subscription.
	//
	// 	- **POSTPAY**: pay-as-you-go.
	//
	// example:
	//
	// PREPAY
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// The region ID of the instance.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the instance belongs.
	//
	// example:
	//
	// rg-aekzledqeat****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The series of the instance. Valid values:
	//
	// 	- **lindorm**: The instance is a Lindorm instance.
	//
	// 	- **serverless_lindorm**: The instance is a Lindorm Serverless instance.
	//
	// 	- **lindorm_standalone**: The instance is a single-node Lindorm instance.
	//
	// 	- **lts**: The instance is an LTS instance.
	//
	// example:
	//
	// lindorm
	ServiceType *string `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
	// The list of tags associated with the specified instances.
	Tags []*GetLindormInstanceListResponseBodyInstanceListTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The ID of the VPC in which the instance is deployed.
	//
	// example:
	//
	// vpc-bp1n3i15v90el48nx****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The ID of the zone in which the instance is created.
	//
	// example:
	//
	// cn-hangzhou-h
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s GetLindormInstanceListResponseBodyInstanceList) String() string {
	return tea.Prettify(s)
}

func (s GetLindormInstanceListResponseBodyInstanceList) GoString() string {
	return s.String()
}

func (s *GetLindormInstanceListResponseBodyInstanceList) SetAliUid(v int64) *GetLindormInstanceListResponseBodyInstanceList {
	s.AliUid = &v
	return s
}

func (s *GetLindormInstanceListResponseBodyInstanceList) SetCreateMilliseconds(v int64) *GetLindormInstanceListResponseBodyInstanceList {
	s.CreateMilliseconds = &v
	return s
}

func (s *GetLindormInstanceListResponseBodyInstanceList) SetCreateTime(v string) *GetLindormInstanceListResponseBodyInstanceList {
	s.CreateTime = &v
	return s
}

func (s *GetLindormInstanceListResponseBodyInstanceList) SetEnableColumn(v bool) *GetLindormInstanceListResponseBodyInstanceList {
	s.EnableColumn = &v
	return s
}

func (s *GetLindormInstanceListResponseBodyInstanceList) SetEnableCompute(v bool) *GetLindormInstanceListResponseBodyInstanceList {
	s.EnableCompute = &v
	return s
}

func (s *GetLindormInstanceListResponseBodyInstanceList) SetEnableLts(v bool) *GetLindormInstanceListResponseBodyInstanceList {
	s.EnableLts = &v
	return s
}

func (s *GetLindormInstanceListResponseBodyInstanceList) SetEnableMessage(v bool) *GetLindormInstanceListResponseBodyInstanceList {
	s.EnableMessage = &v
	return s
}

func (s *GetLindormInstanceListResponseBodyInstanceList) SetEnableRow(v bool) *GetLindormInstanceListResponseBodyInstanceList {
	s.EnableRow = &v
	return s
}

func (s *GetLindormInstanceListResponseBodyInstanceList) SetEnableStream(v bool) *GetLindormInstanceListResponseBodyInstanceList {
	s.EnableStream = &v
	return s
}

func (s *GetLindormInstanceListResponseBodyInstanceList) SetEnableVector(v bool) *GetLindormInstanceListResponseBodyInstanceList {
	s.EnableVector = &v
	return s
}

func (s *GetLindormInstanceListResponseBodyInstanceList) SetEngineType(v string) *GetLindormInstanceListResponseBodyInstanceList {
	s.EngineType = &v
	return s
}

func (s *GetLindormInstanceListResponseBodyInstanceList) SetExpireTime(v string) *GetLindormInstanceListResponseBodyInstanceList {
	s.ExpireTime = &v
	return s
}

func (s *GetLindormInstanceListResponseBodyInstanceList) SetExpiredMilliseconds(v int64) *GetLindormInstanceListResponseBodyInstanceList {
	s.ExpiredMilliseconds = &v
	return s
}

func (s *GetLindormInstanceListResponseBodyInstanceList) SetInstanceAlias(v string) *GetLindormInstanceListResponseBodyInstanceList {
	s.InstanceAlias = &v
	return s
}

func (s *GetLindormInstanceListResponseBodyInstanceList) SetInstanceId(v string) *GetLindormInstanceListResponseBodyInstanceList {
	s.InstanceId = &v
	return s
}

func (s *GetLindormInstanceListResponseBodyInstanceList) SetInstanceStatus(v string) *GetLindormInstanceListResponseBodyInstanceList {
	s.InstanceStatus = &v
	return s
}

func (s *GetLindormInstanceListResponseBodyInstanceList) SetInstanceStorage(v string) *GetLindormInstanceListResponseBodyInstanceList {
	s.InstanceStorage = &v
	return s
}

func (s *GetLindormInstanceListResponseBodyInstanceList) SetNetworkType(v string) *GetLindormInstanceListResponseBodyInstanceList {
	s.NetworkType = &v
	return s
}

func (s *GetLindormInstanceListResponseBodyInstanceList) SetPayType(v string) *GetLindormInstanceListResponseBodyInstanceList {
	s.PayType = &v
	return s
}

func (s *GetLindormInstanceListResponseBodyInstanceList) SetRegionId(v string) *GetLindormInstanceListResponseBodyInstanceList {
	s.RegionId = &v
	return s
}

func (s *GetLindormInstanceListResponseBodyInstanceList) SetResourceGroupId(v string) *GetLindormInstanceListResponseBodyInstanceList {
	s.ResourceGroupId = &v
	return s
}

func (s *GetLindormInstanceListResponseBodyInstanceList) SetServiceType(v string) *GetLindormInstanceListResponseBodyInstanceList {
	s.ServiceType = &v
	return s
}

func (s *GetLindormInstanceListResponseBodyInstanceList) SetTags(v []*GetLindormInstanceListResponseBodyInstanceListTags) *GetLindormInstanceListResponseBodyInstanceList {
	s.Tags = v
	return s
}

func (s *GetLindormInstanceListResponseBodyInstanceList) SetVpcId(v string) *GetLindormInstanceListResponseBodyInstanceList {
	s.VpcId = &v
	return s
}

func (s *GetLindormInstanceListResponseBodyInstanceList) SetZoneId(v string) *GetLindormInstanceListResponseBodyInstanceList {
	s.ZoneId = &v
	return s
}

type GetLindormInstanceListResponseBodyInstanceListTags struct {
	// The key of the tag.
	//
	// example:
	//
	// test
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag.
	//
	// example:
	//
	// 2.2.18
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetLindormInstanceListResponseBodyInstanceListTags) String() string {
	return tea.Prettify(s)
}

func (s GetLindormInstanceListResponseBodyInstanceListTags) GoString() string {
	return s.String()
}

func (s *GetLindormInstanceListResponseBodyInstanceListTags) SetKey(v string) *GetLindormInstanceListResponseBodyInstanceListTags {
	s.Key = &v
	return s
}

func (s *GetLindormInstanceListResponseBodyInstanceListTags) SetValue(v string) *GetLindormInstanceListResponseBodyInstanceListTags {
	s.Value = &v
	return s
}

type GetLindormInstanceListResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetLindormInstanceListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetLindormInstanceListResponse) String() string {
	return tea.Prettify(s)
}

func (s GetLindormInstanceListResponse) GoString() string {
	return s.String()
}

func (s *GetLindormInstanceListResponse) SetHeaders(v map[string]*string) *GetLindormInstanceListResponse {
	s.Headers = v
	return s
}

func (s *GetLindormInstanceListResponse) SetStatusCode(v int32) *GetLindormInstanceListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetLindormInstanceListResponse) SetBody(v *GetLindormInstanceListResponseBody) *GetLindormInstanceListResponse {
	s.Body = v
	return s
}

type GetLindormV2InstanceRequest struct {
	// This parameter is required.
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s GetLindormV2InstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s GetLindormV2InstanceRequest) GoString() string {
	return s.String()
}

func (s *GetLindormV2InstanceRequest) SetInstanceId(v string) *GetLindormV2InstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *GetLindormV2InstanceRequest) SetOwnerAccount(v string) *GetLindormV2InstanceRequest {
	s.OwnerAccount = &v
	return s
}

func (s *GetLindormV2InstanceRequest) SetOwnerId(v int64) *GetLindormV2InstanceRequest {
	s.OwnerId = &v
	return s
}

func (s *GetLindormV2InstanceRequest) SetResourceOwnerAccount(v string) *GetLindormV2InstanceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetLindormV2InstanceRequest) SetResourceOwnerId(v int64) *GetLindormV2InstanceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetLindormV2InstanceRequest) SetSecurityToken(v string) *GetLindormV2InstanceRequest {
	s.SecurityToken = &v
	return s
}

type GetLindormV2InstanceResponseBody struct {
	AliUid              *int64                                         `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	AutoRenew           *bool                                          `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	ColdStorage         *int32                                         `json:"ColdStorage,omitempty" xml:"ColdStorage,omitempty"`
	CreateMilliseconds  *int64                                         `json:"CreateMilliseconds,omitempty" xml:"CreateMilliseconds,omitempty"`
	DeletionProtection  *string                                        `json:"DeletionProtection,omitempty" xml:"DeletionProtection,omitempty"`
	DiskCategory        *string                                        `json:"DiskCategory,omitempty" xml:"DiskCategory,omitempty"`
	DiskThreshold       *string                                        `json:"DiskThreshold,omitempty" xml:"DiskThreshold,omitempty"`
	DiskUsage           *string                                        `json:"DiskUsage,omitempty" xml:"DiskUsage,omitempty"`
	EnableCompute       *bool                                          `json:"EnableCompute,omitempty" xml:"EnableCompute,omitempty"`
	EngineList          []*GetLindormV2InstanceResponseBodyEngineList  `json:"EngineList,omitempty" xml:"EngineList,omitempty" type:"Repeated"`
	ExpiredMilliseconds *int64                                         `json:"ExpiredMilliseconds,omitempty" xml:"ExpiredMilliseconds,omitempty"`
	InitialRootPassword *string                                        `json:"InitialRootPassword,omitempty" xml:"InitialRootPassword,omitempty"`
	InstanceAlias       *string                                        `json:"InstanceAlias,omitempty" xml:"InstanceAlias,omitempty"`
	InstanceId          *string                                        `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceStatus      *string                                        `json:"InstanceStatus,omitempty" xml:"InstanceStatus,omitempty"`
	InstanceType        *string                                        `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	MaintainEndTime     *string                                        `json:"MaintainEndTime,omitempty" xml:"MaintainEndTime,omitempty"`
	MaintainStartTime   *string                                        `json:"MaintainStartTime,omitempty" xml:"MaintainStartTime,omitempty"`
	NetworkType         *string                                        `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	PayType             *string                                        `json:"PayType,omitempty" xml:"PayType,omitempty"`
	RegionId            *string                                        `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RequestId           *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResourceGroupId     *string                                        `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ServiceType         *string                                        `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
	StorageUsage        *GetLindormV2InstanceResponseBodyStorageUsage  `json:"StorageUsage,omitempty" xml:"StorageUsage,omitempty" type:"Struct"`
	VpcId               *string                                        `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	VswitchId           *string                                        `json:"VswitchId,omitempty" xml:"VswitchId,omitempty"`
	WhiteIpList         []*GetLindormV2InstanceResponseBodyWhiteIpList `json:"WhiteIpList,omitempty" xml:"WhiteIpList,omitempty" type:"Repeated"`
	ZoneEngineInfoMap   map[string]interface{}                         `json:"ZoneEngineInfoMap,omitempty" xml:"ZoneEngineInfoMap,omitempty"`
	ZoneId              *string                                        `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s GetLindormV2InstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetLindormV2InstanceResponseBody) GoString() string {
	return s.String()
}

func (s *GetLindormV2InstanceResponseBody) SetAliUid(v int64) *GetLindormV2InstanceResponseBody {
	s.AliUid = &v
	return s
}

func (s *GetLindormV2InstanceResponseBody) SetAutoRenew(v bool) *GetLindormV2InstanceResponseBody {
	s.AutoRenew = &v
	return s
}

func (s *GetLindormV2InstanceResponseBody) SetColdStorage(v int32) *GetLindormV2InstanceResponseBody {
	s.ColdStorage = &v
	return s
}

func (s *GetLindormV2InstanceResponseBody) SetCreateMilliseconds(v int64) *GetLindormV2InstanceResponseBody {
	s.CreateMilliseconds = &v
	return s
}

func (s *GetLindormV2InstanceResponseBody) SetDeletionProtection(v string) *GetLindormV2InstanceResponseBody {
	s.DeletionProtection = &v
	return s
}

func (s *GetLindormV2InstanceResponseBody) SetDiskCategory(v string) *GetLindormV2InstanceResponseBody {
	s.DiskCategory = &v
	return s
}

func (s *GetLindormV2InstanceResponseBody) SetDiskThreshold(v string) *GetLindormV2InstanceResponseBody {
	s.DiskThreshold = &v
	return s
}

func (s *GetLindormV2InstanceResponseBody) SetDiskUsage(v string) *GetLindormV2InstanceResponseBody {
	s.DiskUsage = &v
	return s
}

func (s *GetLindormV2InstanceResponseBody) SetEnableCompute(v bool) *GetLindormV2InstanceResponseBody {
	s.EnableCompute = &v
	return s
}

func (s *GetLindormV2InstanceResponseBody) SetEngineList(v []*GetLindormV2InstanceResponseBodyEngineList) *GetLindormV2InstanceResponseBody {
	s.EngineList = v
	return s
}

func (s *GetLindormV2InstanceResponseBody) SetExpiredMilliseconds(v int64) *GetLindormV2InstanceResponseBody {
	s.ExpiredMilliseconds = &v
	return s
}

func (s *GetLindormV2InstanceResponseBody) SetInitialRootPassword(v string) *GetLindormV2InstanceResponseBody {
	s.InitialRootPassword = &v
	return s
}

func (s *GetLindormV2InstanceResponseBody) SetInstanceAlias(v string) *GetLindormV2InstanceResponseBody {
	s.InstanceAlias = &v
	return s
}

func (s *GetLindormV2InstanceResponseBody) SetInstanceId(v string) *GetLindormV2InstanceResponseBody {
	s.InstanceId = &v
	return s
}

func (s *GetLindormV2InstanceResponseBody) SetInstanceStatus(v string) *GetLindormV2InstanceResponseBody {
	s.InstanceStatus = &v
	return s
}

func (s *GetLindormV2InstanceResponseBody) SetInstanceType(v string) *GetLindormV2InstanceResponseBody {
	s.InstanceType = &v
	return s
}

func (s *GetLindormV2InstanceResponseBody) SetMaintainEndTime(v string) *GetLindormV2InstanceResponseBody {
	s.MaintainEndTime = &v
	return s
}

func (s *GetLindormV2InstanceResponseBody) SetMaintainStartTime(v string) *GetLindormV2InstanceResponseBody {
	s.MaintainStartTime = &v
	return s
}

func (s *GetLindormV2InstanceResponseBody) SetNetworkType(v string) *GetLindormV2InstanceResponseBody {
	s.NetworkType = &v
	return s
}

func (s *GetLindormV2InstanceResponseBody) SetPayType(v string) *GetLindormV2InstanceResponseBody {
	s.PayType = &v
	return s
}

func (s *GetLindormV2InstanceResponseBody) SetRegionId(v string) *GetLindormV2InstanceResponseBody {
	s.RegionId = &v
	return s
}

func (s *GetLindormV2InstanceResponseBody) SetRequestId(v string) *GetLindormV2InstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetLindormV2InstanceResponseBody) SetResourceGroupId(v string) *GetLindormV2InstanceResponseBody {
	s.ResourceGroupId = &v
	return s
}

func (s *GetLindormV2InstanceResponseBody) SetServiceType(v string) *GetLindormV2InstanceResponseBody {
	s.ServiceType = &v
	return s
}

func (s *GetLindormV2InstanceResponseBody) SetStorageUsage(v *GetLindormV2InstanceResponseBodyStorageUsage) *GetLindormV2InstanceResponseBody {
	s.StorageUsage = v
	return s
}

func (s *GetLindormV2InstanceResponseBody) SetVpcId(v string) *GetLindormV2InstanceResponseBody {
	s.VpcId = &v
	return s
}

func (s *GetLindormV2InstanceResponseBody) SetVswitchId(v string) *GetLindormV2InstanceResponseBody {
	s.VswitchId = &v
	return s
}

func (s *GetLindormV2InstanceResponseBody) SetWhiteIpList(v []*GetLindormV2InstanceResponseBodyWhiteIpList) *GetLindormV2InstanceResponseBody {
	s.WhiteIpList = v
	return s
}

func (s *GetLindormV2InstanceResponseBody) SetZoneEngineInfoMap(v map[string]interface{}) *GetLindormV2InstanceResponseBody {
	s.ZoneEngineInfoMap = v
	return s
}

func (s *GetLindormV2InstanceResponseBody) SetZoneId(v string) *GetLindormV2InstanceResponseBody {
	s.ZoneId = &v
	return s
}

type GetLindormV2InstanceResponseBodyEngineList struct {
	ConnectAddressList []*GetLindormV2InstanceResponseBodyEngineListConnectAddressList `json:"ConnectAddressList,omitempty" xml:"ConnectAddressList,omitempty" type:"Repeated"`
	Engine             *string                                                         `json:"Engine,omitempty" xml:"Engine,omitempty"`
	IsLastVersion      *bool                                                           `json:"IsLastVersion,omitempty" xml:"IsLastVersion,omitempty"`
	LatestVersion      *string                                                         `json:"LatestVersion,omitempty" xml:"LatestVersion,omitempty"`
	NodeGroup          []*GetLindormV2InstanceResponseBodyEngineListNodeGroup          `json:"NodeGroup,omitempty" xml:"NodeGroup,omitempty" type:"Repeated"`
	Version            *string                                                         `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s GetLindormV2InstanceResponseBodyEngineList) String() string {
	return tea.Prettify(s)
}

func (s GetLindormV2InstanceResponseBodyEngineList) GoString() string {
	return s.String()
}

func (s *GetLindormV2InstanceResponseBodyEngineList) SetConnectAddressList(v []*GetLindormV2InstanceResponseBodyEngineListConnectAddressList) *GetLindormV2InstanceResponseBodyEngineList {
	s.ConnectAddressList = v
	return s
}

func (s *GetLindormV2InstanceResponseBodyEngineList) SetEngine(v string) *GetLindormV2InstanceResponseBodyEngineList {
	s.Engine = &v
	return s
}

func (s *GetLindormV2InstanceResponseBodyEngineList) SetIsLastVersion(v bool) *GetLindormV2InstanceResponseBodyEngineList {
	s.IsLastVersion = &v
	return s
}

func (s *GetLindormV2InstanceResponseBodyEngineList) SetLatestVersion(v string) *GetLindormV2InstanceResponseBodyEngineList {
	s.LatestVersion = &v
	return s
}

func (s *GetLindormV2InstanceResponseBodyEngineList) SetNodeGroup(v []*GetLindormV2InstanceResponseBodyEngineListNodeGroup) *GetLindormV2InstanceResponseBodyEngineList {
	s.NodeGroup = v
	return s
}

func (s *GetLindormV2InstanceResponseBodyEngineList) SetVersion(v string) *GetLindormV2InstanceResponseBodyEngineList {
	s.Version = &v
	return s
}

type GetLindormV2InstanceResponseBodyEngineListConnectAddressList struct {
	Address *string `json:"Address,omitempty" xml:"Address,omitempty"`
	Port    *string `json:"Port,omitempty" xml:"Port,omitempty"`
	Type    *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetLindormV2InstanceResponseBodyEngineListConnectAddressList) String() string {
	return tea.Prettify(s)
}

func (s GetLindormV2InstanceResponseBodyEngineListConnectAddressList) GoString() string {
	return s.String()
}

func (s *GetLindormV2InstanceResponseBodyEngineListConnectAddressList) SetAddress(v string) *GetLindormV2InstanceResponseBodyEngineListConnectAddressList {
	s.Address = &v
	return s
}

func (s *GetLindormV2InstanceResponseBodyEngineListConnectAddressList) SetPort(v string) *GetLindormV2InstanceResponseBodyEngineListConnectAddressList {
	s.Port = &v
	return s
}

func (s *GetLindormV2InstanceResponseBodyEngineListConnectAddressList) SetType(v string) *GetLindormV2InstanceResponseBodyEngineListConnectAddressList {
	s.Type = &v
	return s
}

type GetLindormV2InstanceResponseBodyEngineListNodeGroup struct {
	Category              *string `json:"Category,omitempty" xml:"Category,omitempty"`
	CpuCoreCount          *int32  `json:"CpuCoreCount,omitempty" xml:"CpuCoreCount,omitempty"`
	EnableAttachLocalDisk *bool   `json:"EnableAttachLocalDisk,omitempty" xml:"EnableAttachLocalDisk,omitempty"`
	LocalDiskCapacity     *int64  `json:"LocalDiskCapacity,omitempty" xml:"LocalDiskCapacity,omitempty"`
	LocalDiskCategory     *string `json:"LocalDiskCategory,omitempty" xml:"LocalDiskCategory,omitempty"`
	MemorySizeGiB         *int32  `json:"MemorySizeGiB,omitempty" xml:"MemorySizeGiB,omitempty"`
	NodeSpec              *string `json:"NodeSpec,omitempty" xml:"NodeSpec,omitempty"`
	Quantity              *int32  `json:"Quantity,omitempty" xml:"Quantity,omitempty"`
	ResourceGroupName     *string `json:"ResourceGroupName,omitempty" xml:"ResourceGroupName,omitempty"`
	SpecId                *string `json:"SpecId,omitempty" xml:"SpecId,omitempty"`
	Status                *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetLindormV2InstanceResponseBodyEngineListNodeGroup) String() string {
	return tea.Prettify(s)
}

func (s GetLindormV2InstanceResponseBodyEngineListNodeGroup) GoString() string {
	return s.String()
}

func (s *GetLindormV2InstanceResponseBodyEngineListNodeGroup) SetCategory(v string) *GetLindormV2InstanceResponseBodyEngineListNodeGroup {
	s.Category = &v
	return s
}

func (s *GetLindormV2InstanceResponseBodyEngineListNodeGroup) SetCpuCoreCount(v int32) *GetLindormV2InstanceResponseBodyEngineListNodeGroup {
	s.CpuCoreCount = &v
	return s
}

func (s *GetLindormV2InstanceResponseBodyEngineListNodeGroup) SetEnableAttachLocalDisk(v bool) *GetLindormV2InstanceResponseBodyEngineListNodeGroup {
	s.EnableAttachLocalDisk = &v
	return s
}

func (s *GetLindormV2InstanceResponseBodyEngineListNodeGroup) SetLocalDiskCapacity(v int64) *GetLindormV2InstanceResponseBodyEngineListNodeGroup {
	s.LocalDiskCapacity = &v
	return s
}

func (s *GetLindormV2InstanceResponseBodyEngineListNodeGroup) SetLocalDiskCategory(v string) *GetLindormV2InstanceResponseBodyEngineListNodeGroup {
	s.LocalDiskCategory = &v
	return s
}

func (s *GetLindormV2InstanceResponseBodyEngineListNodeGroup) SetMemorySizeGiB(v int32) *GetLindormV2InstanceResponseBodyEngineListNodeGroup {
	s.MemorySizeGiB = &v
	return s
}

func (s *GetLindormV2InstanceResponseBodyEngineListNodeGroup) SetNodeSpec(v string) *GetLindormV2InstanceResponseBodyEngineListNodeGroup {
	s.NodeSpec = &v
	return s
}

func (s *GetLindormV2InstanceResponseBodyEngineListNodeGroup) SetQuantity(v int32) *GetLindormV2InstanceResponseBodyEngineListNodeGroup {
	s.Quantity = &v
	return s
}

func (s *GetLindormV2InstanceResponseBodyEngineListNodeGroup) SetResourceGroupName(v string) *GetLindormV2InstanceResponseBodyEngineListNodeGroup {
	s.ResourceGroupName = &v
	return s
}

func (s *GetLindormV2InstanceResponseBodyEngineListNodeGroup) SetSpecId(v string) *GetLindormV2InstanceResponseBodyEngineListNodeGroup {
	s.SpecId = &v
	return s
}

func (s *GetLindormV2InstanceResponseBodyEngineListNodeGroup) SetStatus(v string) *GetLindormV2InstanceResponseBodyEngineListNodeGroup {
	s.Status = &v
	return s
}

type GetLindormV2InstanceResponseBodyStorageUsage struct {
	CapacityByDiskCategory []map[string]interface{} `json:"CapacityByDiskCategory,omitempty" xml:"CapacityByDiskCategory,omitempty" type:"Repeated"`
	EngineUsage            map[string]interface{}   `json:"EngineUsage,omitempty" xml:"EngineUsage,omitempty"`
}

func (s GetLindormV2InstanceResponseBodyStorageUsage) String() string {
	return tea.Prettify(s)
}

func (s GetLindormV2InstanceResponseBodyStorageUsage) GoString() string {
	return s.String()
}

func (s *GetLindormV2InstanceResponseBodyStorageUsage) SetCapacityByDiskCategory(v []map[string]interface{}) *GetLindormV2InstanceResponseBodyStorageUsage {
	s.CapacityByDiskCategory = v
	return s
}

func (s *GetLindormV2InstanceResponseBodyStorageUsage) SetEngineUsage(v map[string]interface{}) *GetLindormV2InstanceResponseBodyStorageUsage {
	s.EngineUsage = v
	return s
}

type GetLindormV2InstanceResponseBodyWhiteIpList struct {
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	IpList    *string `json:"IpList,omitempty" xml:"IpList,omitempty"`
}

func (s GetLindormV2InstanceResponseBodyWhiteIpList) String() string {
	return tea.Prettify(s)
}

func (s GetLindormV2InstanceResponseBodyWhiteIpList) GoString() string {
	return s.String()
}

func (s *GetLindormV2InstanceResponseBodyWhiteIpList) SetGroupName(v string) *GetLindormV2InstanceResponseBodyWhiteIpList {
	s.GroupName = &v
	return s
}

func (s *GetLindormV2InstanceResponseBodyWhiteIpList) SetIpList(v string) *GetLindormV2InstanceResponseBodyWhiteIpList {
	s.IpList = &v
	return s
}

type GetLindormV2InstanceResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetLindormV2InstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetLindormV2InstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s GetLindormV2InstanceResponse) GoString() string {
	return s.String()
}

func (s *GetLindormV2InstanceResponse) SetHeaders(v map[string]*string) *GetLindormV2InstanceResponse {
	s.Headers = v
	return s
}

func (s *GetLindormV2InstanceResponse) SetStatusCode(v int32) *GetLindormV2InstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetLindormV2InstanceResponse) SetBody(v *GetLindormV2InstanceResponseBody) *GetLindormV2InstanceResponse {
	s.Body = v
	return s
}

type GetLindormV2InstanceEngineListRequest struct {
	// This parameter is required.
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s GetLindormV2InstanceEngineListRequest) String() string {
	return tea.Prettify(s)
}

func (s GetLindormV2InstanceEngineListRequest) GoString() string {
	return s.String()
}

func (s *GetLindormV2InstanceEngineListRequest) SetInstanceId(v string) *GetLindormV2InstanceEngineListRequest {
	s.InstanceId = &v
	return s
}

func (s *GetLindormV2InstanceEngineListRequest) SetOwnerAccount(v string) *GetLindormV2InstanceEngineListRequest {
	s.OwnerAccount = &v
	return s
}

func (s *GetLindormV2InstanceEngineListRequest) SetOwnerId(v int64) *GetLindormV2InstanceEngineListRequest {
	s.OwnerId = &v
	return s
}

func (s *GetLindormV2InstanceEngineListRequest) SetRegionId(v string) *GetLindormV2InstanceEngineListRequest {
	s.RegionId = &v
	return s
}

func (s *GetLindormV2InstanceEngineListRequest) SetResourceOwnerAccount(v string) *GetLindormV2InstanceEngineListRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetLindormV2InstanceEngineListRequest) SetResourceOwnerId(v int64) *GetLindormV2InstanceEngineListRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetLindormV2InstanceEngineListRequest) SetSecurityToken(v string) *GetLindormV2InstanceEngineListRequest {
	s.SecurityToken = &v
	return s
}

type GetLindormV2InstanceEngineListResponseBody struct {
	AccessDeniedDetail *string                                                 `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	EngineList         []*GetLindormV2InstanceEngineListResponseBodyEngineList `json:"EngineList,omitempty" xml:"EngineList,omitempty" type:"Repeated"`
	InstanceId         *string                                                 `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RequestId          *string                                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetLindormV2InstanceEngineListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetLindormV2InstanceEngineListResponseBody) GoString() string {
	return s.String()
}

func (s *GetLindormV2InstanceEngineListResponseBody) SetAccessDeniedDetail(v string) *GetLindormV2InstanceEngineListResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *GetLindormV2InstanceEngineListResponseBody) SetEngineList(v []*GetLindormV2InstanceEngineListResponseBodyEngineList) *GetLindormV2InstanceEngineListResponseBody {
	s.EngineList = v
	return s
}

func (s *GetLindormV2InstanceEngineListResponseBody) SetInstanceId(v string) *GetLindormV2InstanceEngineListResponseBody {
	s.InstanceId = &v
	return s
}

func (s *GetLindormV2InstanceEngineListResponseBody) SetRequestId(v string) *GetLindormV2InstanceEngineListResponseBody {
	s.RequestId = &v
	return s
}

type GetLindormV2InstanceEngineListResponseBodyEngineList struct {
	EngineType  *string                                                            `json:"EngineType,omitempty" xml:"EngineType,omitempty"`
	NetInfoList []*GetLindormV2InstanceEngineListResponseBodyEngineListNetInfoList `json:"NetInfoList,omitempty" xml:"NetInfoList,omitempty" type:"Repeated"`
}

func (s GetLindormV2InstanceEngineListResponseBodyEngineList) String() string {
	return tea.Prettify(s)
}

func (s GetLindormV2InstanceEngineListResponseBodyEngineList) GoString() string {
	return s.String()
}

func (s *GetLindormV2InstanceEngineListResponseBodyEngineList) SetEngineType(v string) *GetLindormV2InstanceEngineListResponseBodyEngineList {
	s.EngineType = &v
	return s
}

func (s *GetLindormV2InstanceEngineListResponseBodyEngineList) SetNetInfoList(v []*GetLindormV2InstanceEngineListResponseBodyEngineListNetInfoList) *GetLindormV2InstanceEngineListResponseBodyEngineList {
	s.NetInfoList = v
	return s
}

type GetLindormV2InstanceEngineListResponseBodyEngineListNetInfoList struct {
	AccessType       *int32  `json:"AccessType,omitempty" xml:"AccessType,omitempty"`
	ConnectionString *string `json:"ConnectionString,omitempty" xml:"ConnectionString,omitempty"`
	NetType          *string `json:"NetType,omitempty" xml:"NetType,omitempty"`
	Port             *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
}

func (s GetLindormV2InstanceEngineListResponseBodyEngineListNetInfoList) String() string {
	return tea.Prettify(s)
}

func (s GetLindormV2InstanceEngineListResponseBodyEngineListNetInfoList) GoString() string {
	return s.String()
}

func (s *GetLindormV2InstanceEngineListResponseBodyEngineListNetInfoList) SetAccessType(v int32) *GetLindormV2InstanceEngineListResponseBodyEngineListNetInfoList {
	s.AccessType = &v
	return s
}

func (s *GetLindormV2InstanceEngineListResponseBodyEngineListNetInfoList) SetConnectionString(v string) *GetLindormV2InstanceEngineListResponseBodyEngineListNetInfoList {
	s.ConnectionString = &v
	return s
}

func (s *GetLindormV2InstanceEngineListResponseBodyEngineListNetInfoList) SetNetType(v string) *GetLindormV2InstanceEngineListResponseBodyEngineListNetInfoList {
	s.NetType = &v
	return s
}

func (s *GetLindormV2InstanceEngineListResponseBodyEngineListNetInfoList) SetPort(v int32) *GetLindormV2InstanceEngineListResponseBodyEngineListNetInfoList {
	s.Port = &v
	return s
}

type GetLindormV2InstanceEngineListResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetLindormV2InstanceEngineListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetLindormV2InstanceEngineListResponse) String() string {
	return tea.Prettify(s)
}

func (s GetLindormV2InstanceEngineListResponse) GoString() string {
	return s.String()
}

func (s *GetLindormV2InstanceEngineListResponse) SetHeaders(v map[string]*string) *GetLindormV2InstanceEngineListResponse {
	s.Headers = v
	return s
}

func (s *GetLindormV2InstanceEngineListResponse) SetStatusCode(v int32) *GetLindormV2InstanceEngineListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetLindormV2InstanceEngineListResponse) SetBody(v *GetLindormV2InstanceEngineListResponseBody) *GetLindormV2InstanceEngineListResponse {
	s.Body = v
	return s
}

type GetLindormV2StorageUsageRequest struct {
	// This parameter is required.
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s GetLindormV2StorageUsageRequest) String() string {
	return tea.Prettify(s)
}

func (s GetLindormV2StorageUsageRequest) GoString() string {
	return s.String()
}

func (s *GetLindormV2StorageUsageRequest) SetInstanceId(v string) *GetLindormV2StorageUsageRequest {
	s.InstanceId = &v
	return s
}

func (s *GetLindormV2StorageUsageRequest) SetOwnerAccount(v string) *GetLindormV2StorageUsageRequest {
	s.OwnerAccount = &v
	return s
}

func (s *GetLindormV2StorageUsageRequest) SetOwnerId(v int64) *GetLindormV2StorageUsageRequest {
	s.OwnerId = &v
	return s
}

func (s *GetLindormV2StorageUsageRequest) SetResourceOwnerAccount(v string) *GetLindormV2StorageUsageRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetLindormV2StorageUsageRequest) SetResourceOwnerId(v int64) *GetLindormV2StorageUsageRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetLindormV2StorageUsageRequest) SetSecurityToken(v string) *GetLindormV2StorageUsageRequest {
	s.SecurityToken = &v
	return s
}

type GetLindormV2StorageUsageResponseBody struct {
	AccessDeniedDetail     *string                  `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	CapacityByDiskCategory []map[string]interface{} `json:"CapacityByDiskCategory,omitempty" xml:"CapacityByDiskCategory,omitempty" type:"Repeated"`
	InstanceStorageZoneMap map[string]interface{}   `json:"InstanceStorageZoneMap,omitempty" xml:"InstanceStorageZoneMap,omitempty"`
	RequestId              *string                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	UsageByDiskCategory    []map[string]interface{} `json:"UsageByDiskCategory,omitempty" xml:"UsageByDiskCategory,omitempty" type:"Repeated"`
}

func (s GetLindormV2StorageUsageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetLindormV2StorageUsageResponseBody) GoString() string {
	return s.String()
}

func (s *GetLindormV2StorageUsageResponseBody) SetAccessDeniedDetail(v string) *GetLindormV2StorageUsageResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *GetLindormV2StorageUsageResponseBody) SetCapacityByDiskCategory(v []map[string]interface{}) *GetLindormV2StorageUsageResponseBody {
	s.CapacityByDiskCategory = v
	return s
}

func (s *GetLindormV2StorageUsageResponseBody) SetInstanceStorageZoneMap(v map[string]interface{}) *GetLindormV2StorageUsageResponseBody {
	s.InstanceStorageZoneMap = v
	return s
}

func (s *GetLindormV2StorageUsageResponseBody) SetRequestId(v string) *GetLindormV2StorageUsageResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetLindormV2StorageUsageResponseBody) SetUsageByDiskCategory(v []map[string]interface{}) *GetLindormV2StorageUsageResponseBody {
	s.UsageByDiskCategory = v
	return s
}

type GetLindormV2StorageUsageResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetLindormV2StorageUsageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetLindormV2StorageUsageResponse) String() string {
	return tea.Prettify(s)
}

func (s GetLindormV2StorageUsageResponse) GoString() string {
	return s.String()
}

func (s *GetLindormV2StorageUsageResponse) SetHeaders(v map[string]*string) *GetLindormV2StorageUsageResponse {
	s.Headers = v
	return s
}

func (s *GetLindormV2StorageUsageResponse) SetStatusCode(v int32) *GetLindormV2StorageUsageResponse {
	s.StatusCode = &v
	return s
}

func (s *GetLindormV2StorageUsageResponse) SetBody(v *GetLindormV2StorageUsageResponseBody) *GetLindormV2StorageUsageResponse {
	s.Body = v
	return s
}

type ListAutoScalingConfigsRequest struct {
	// This parameter is required.
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s ListAutoScalingConfigsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAutoScalingConfigsRequest) GoString() string {
	return s.String()
}

func (s *ListAutoScalingConfigsRequest) SetInstanceId(v string) *ListAutoScalingConfigsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListAutoScalingConfigsRequest) SetOwnerAccount(v string) *ListAutoScalingConfigsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ListAutoScalingConfigsRequest) SetOwnerId(v int64) *ListAutoScalingConfigsRequest {
	s.OwnerId = &v
	return s
}

func (s *ListAutoScalingConfigsRequest) SetResourceOwnerAccount(v string) *ListAutoScalingConfigsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListAutoScalingConfigsRequest) SetResourceOwnerId(v int64) *ListAutoScalingConfigsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListAutoScalingConfigsRequest) SetSecurityToken(v string) *ListAutoScalingConfigsRequest {
	s.SecurityToken = &v
	return s
}

type ListAutoScalingConfigsResponseBody struct {
	AccessDeniedDetail *string                                 `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	Code               *string                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Data               *ListAutoScalingConfigsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	DynamicCode        *string                                 `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	DynamicMessage     *string                                 `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	HttpStatusCode     *int32                                  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message            *string                                 `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId          *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success            *bool                                   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListAutoScalingConfigsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAutoScalingConfigsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAutoScalingConfigsResponseBody) SetAccessDeniedDetail(v string) *ListAutoScalingConfigsResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ListAutoScalingConfigsResponseBody) SetCode(v string) *ListAutoScalingConfigsResponseBody {
	s.Code = &v
	return s
}

func (s *ListAutoScalingConfigsResponseBody) SetData(v *ListAutoScalingConfigsResponseBodyData) *ListAutoScalingConfigsResponseBody {
	s.Data = v
	return s
}

func (s *ListAutoScalingConfigsResponseBody) SetDynamicCode(v string) *ListAutoScalingConfigsResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *ListAutoScalingConfigsResponseBody) SetDynamicMessage(v string) *ListAutoScalingConfigsResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *ListAutoScalingConfigsResponseBody) SetHttpStatusCode(v int32) *ListAutoScalingConfigsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListAutoScalingConfigsResponseBody) SetMessage(v string) *ListAutoScalingConfigsResponseBody {
	s.Message = &v
	return s
}

func (s *ListAutoScalingConfigsResponseBody) SetRequestId(v string) *ListAutoScalingConfigsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAutoScalingConfigsResponseBody) SetSuccess(v bool) *ListAutoScalingConfigsResponseBody {
	s.Success = &v
	return s
}

type ListAutoScalingConfigsResponseBodyData struct {
	ScaleConfigs []*ListAutoScalingConfigsResponseBodyDataScaleConfigs `json:"ScaleConfigs,omitempty" xml:"ScaleConfigs,omitempty" type:"Repeated"`
}

func (s ListAutoScalingConfigsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListAutoScalingConfigsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListAutoScalingConfigsResponseBodyData) SetScaleConfigs(v []*ListAutoScalingConfigsResponseBodyDataScaleConfigs) *ListAutoScalingConfigsResponseBodyData {
	s.ScaleConfigs = v
	return s
}

type ListAutoScalingConfigsResponseBodyDataScaleConfigs struct {
	ConfigId           *string                                                            `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	ConfigName         *string                                                            `json:"ConfigName,omitempty" xml:"ConfigName,omitempty"`
	EffectiveTimeEnd   *string                                                            `json:"EffectiveTimeEnd,omitempty" xml:"EffectiveTimeEnd,omitempty"`
	EffectiveTimeStart *string                                                            `json:"EffectiveTimeStart,omitempty" xml:"EffectiveTimeStart,omitempty"`
	Enabled            *bool                                                              `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	Engine             *string                                                            `json:"Engine,omitempty" xml:"Engine,omitempty"`
	InstanceId         *string                                                            `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	NodesMax           *int32                                                             `json:"NodesMax,omitempty" xml:"NodesMax,omitempty"`
	NodesMin           *int32                                                             `json:"NodesMin,omitempty" xml:"NodesMin,omitempty"`
	ScaleRuleList      []*ListAutoScalingConfigsResponseBodyDataScaleConfigsScaleRuleList `json:"ScaleRuleList,omitempty" xml:"ScaleRuleList,omitempty" type:"Repeated"`
	ScaleType          *string                                                            `json:"ScaleType,omitempty" xml:"ScaleType,omitempty"`
	SpecId             *string                                                            `json:"SpecId,omitempty" xml:"SpecId,omitempty"`
}

func (s ListAutoScalingConfigsResponseBodyDataScaleConfigs) String() string {
	return tea.Prettify(s)
}

func (s ListAutoScalingConfigsResponseBodyDataScaleConfigs) GoString() string {
	return s.String()
}

func (s *ListAutoScalingConfigsResponseBodyDataScaleConfigs) SetConfigId(v string) *ListAutoScalingConfigsResponseBodyDataScaleConfigs {
	s.ConfigId = &v
	return s
}

func (s *ListAutoScalingConfigsResponseBodyDataScaleConfigs) SetConfigName(v string) *ListAutoScalingConfigsResponseBodyDataScaleConfigs {
	s.ConfigName = &v
	return s
}

func (s *ListAutoScalingConfigsResponseBodyDataScaleConfigs) SetEffectiveTimeEnd(v string) *ListAutoScalingConfigsResponseBodyDataScaleConfigs {
	s.EffectiveTimeEnd = &v
	return s
}

func (s *ListAutoScalingConfigsResponseBodyDataScaleConfigs) SetEffectiveTimeStart(v string) *ListAutoScalingConfigsResponseBodyDataScaleConfigs {
	s.EffectiveTimeStart = &v
	return s
}

func (s *ListAutoScalingConfigsResponseBodyDataScaleConfigs) SetEnabled(v bool) *ListAutoScalingConfigsResponseBodyDataScaleConfigs {
	s.Enabled = &v
	return s
}

func (s *ListAutoScalingConfigsResponseBodyDataScaleConfigs) SetEngine(v string) *ListAutoScalingConfigsResponseBodyDataScaleConfigs {
	s.Engine = &v
	return s
}

func (s *ListAutoScalingConfigsResponseBodyDataScaleConfigs) SetInstanceId(v string) *ListAutoScalingConfigsResponseBodyDataScaleConfigs {
	s.InstanceId = &v
	return s
}

func (s *ListAutoScalingConfigsResponseBodyDataScaleConfigs) SetNodesMax(v int32) *ListAutoScalingConfigsResponseBodyDataScaleConfigs {
	s.NodesMax = &v
	return s
}

func (s *ListAutoScalingConfigsResponseBodyDataScaleConfigs) SetNodesMin(v int32) *ListAutoScalingConfigsResponseBodyDataScaleConfigs {
	s.NodesMin = &v
	return s
}

func (s *ListAutoScalingConfigsResponseBodyDataScaleConfigs) SetScaleRuleList(v []*ListAutoScalingConfigsResponseBodyDataScaleConfigsScaleRuleList) *ListAutoScalingConfigsResponseBodyDataScaleConfigs {
	s.ScaleRuleList = v
	return s
}

func (s *ListAutoScalingConfigsResponseBodyDataScaleConfigs) SetScaleType(v string) *ListAutoScalingConfigsResponseBodyDataScaleConfigs {
	s.ScaleType = &v
	return s
}

func (s *ListAutoScalingConfigsResponseBodyDataScaleConfigs) SetSpecId(v string) *ListAutoScalingConfigsResponseBodyDataScaleConfigs {
	s.SpecId = &v
	return s
}

type ListAutoScalingConfigsResponseBodyDataScaleConfigsScaleRuleList struct {
	ConfigId          *string `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	Enabled           *bool   `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	EndTime           *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	InstanceId        *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	ObservationWindow *int32  `json:"ObservationWindow,omitempty" xml:"ObservationWindow,omitempty"`
	OperationType     *string `json:"OperationType,omitempty" xml:"OperationType,omitempty"`
	RuleId            *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	RuleName          *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	RuleType          *string `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
	ScaleInStep       *int32  `json:"ScaleInStep,omitempty" xml:"ScaleInStep,omitempty"`
	ScaleOutStep      *int32  `json:"ScaleOutStep,omitempty" xml:"ScaleOutStep,omitempty"`
	SilenceTime       *int32  `json:"SilenceTime,omitempty" xml:"SilenceTime,omitempty"`
	StartTime         *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	TargetMetric      *string `json:"TargetMetric,omitempty" xml:"TargetMetric,omitempty"`
	TargetNodes       *int32  `json:"TargetNodes,omitempty" xml:"TargetNodes,omitempty"`
	ThresholdLower    *int32  `json:"ThresholdLower,omitempty" xml:"ThresholdLower,omitempty"`
	ThresholdUpper    *int32  `json:"ThresholdUpper,omitempty" xml:"ThresholdUpper,omitempty"`
	TriggerCronExpr   *string `json:"TriggerCronExpr,omitempty" xml:"TriggerCronExpr,omitempty"`
}

func (s ListAutoScalingConfigsResponseBodyDataScaleConfigsScaleRuleList) String() string {
	return tea.Prettify(s)
}

func (s ListAutoScalingConfigsResponseBodyDataScaleConfigsScaleRuleList) GoString() string {
	return s.String()
}

func (s *ListAutoScalingConfigsResponseBodyDataScaleConfigsScaleRuleList) SetConfigId(v string) *ListAutoScalingConfigsResponseBodyDataScaleConfigsScaleRuleList {
	s.ConfigId = &v
	return s
}

func (s *ListAutoScalingConfigsResponseBodyDataScaleConfigsScaleRuleList) SetEnabled(v bool) *ListAutoScalingConfigsResponseBodyDataScaleConfigsScaleRuleList {
	s.Enabled = &v
	return s
}

func (s *ListAutoScalingConfigsResponseBodyDataScaleConfigsScaleRuleList) SetEndTime(v string) *ListAutoScalingConfigsResponseBodyDataScaleConfigsScaleRuleList {
	s.EndTime = &v
	return s
}

func (s *ListAutoScalingConfigsResponseBodyDataScaleConfigsScaleRuleList) SetInstanceId(v string) *ListAutoScalingConfigsResponseBodyDataScaleConfigsScaleRuleList {
	s.InstanceId = &v
	return s
}

func (s *ListAutoScalingConfigsResponseBodyDataScaleConfigsScaleRuleList) SetObservationWindow(v int32) *ListAutoScalingConfigsResponseBodyDataScaleConfigsScaleRuleList {
	s.ObservationWindow = &v
	return s
}

func (s *ListAutoScalingConfigsResponseBodyDataScaleConfigsScaleRuleList) SetOperationType(v string) *ListAutoScalingConfigsResponseBodyDataScaleConfigsScaleRuleList {
	s.OperationType = &v
	return s
}

func (s *ListAutoScalingConfigsResponseBodyDataScaleConfigsScaleRuleList) SetRuleId(v string) *ListAutoScalingConfigsResponseBodyDataScaleConfigsScaleRuleList {
	s.RuleId = &v
	return s
}

func (s *ListAutoScalingConfigsResponseBodyDataScaleConfigsScaleRuleList) SetRuleName(v string) *ListAutoScalingConfigsResponseBodyDataScaleConfigsScaleRuleList {
	s.RuleName = &v
	return s
}

func (s *ListAutoScalingConfigsResponseBodyDataScaleConfigsScaleRuleList) SetRuleType(v string) *ListAutoScalingConfigsResponseBodyDataScaleConfigsScaleRuleList {
	s.RuleType = &v
	return s
}

func (s *ListAutoScalingConfigsResponseBodyDataScaleConfigsScaleRuleList) SetScaleInStep(v int32) *ListAutoScalingConfigsResponseBodyDataScaleConfigsScaleRuleList {
	s.ScaleInStep = &v
	return s
}

func (s *ListAutoScalingConfigsResponseBodyDataScaleConfigsScaleRuleList) SetScaleOutStep(v int32) *ListAutoScalingConfigsResponseBodyDataScaleConfigsScaleRuleList {
	s.ScaleOutStep = &v
	return s
}

func (s *ListAutoScalingConfigsResponseBodyDataScaleConfigsScaleRuleList) SetSilenceTime(v int32) *ListAutoScalingConfigsResponseBodyDataScaleConfigsScaleRuleList {
	s.SilenceTime = &v
	return s
}

func (s *ListAutoScalingConfigsResponseBodyDataScaleConfigsScaleRuleList) SetStartTime(v string) *ListAutoScalingConfigsResponseBodyDataScaleConfigsScaleRuleList {
	s.StartTime = &v
	return s
}

func (s *ListAutoScalingConfigsResponseBodyDataScaleConfigsScaleRuleList) SetTargetMetric(v string) *ListAutoScalingConfigsResponseBodyDataScaleConfigsScaleRuleList {
	s.TargetMetric = &v
	return s
}

func (s *ListAutoScalingConfigsResponseBodyDataScaleConfigsScaleRuleList) SetTargetNodes(v int32) *ListAutoScalingConfigsResponseBodyDataScaleConfigsScaleRuleList {
	s.TargetNodes = &v
	return s
}

func (s *ListAutoScalingConfigsResponseBodyDataScaleConfigsScaleRuleList) SetThresholdLower(v int32) *ListAutoScalingConfigsResponseBodyDataScaleConfigsScaleRuleList {
	s.ThresholdLower = &v
	return s
}

func (s *ListAutoScalingConfigsResponseBodyDataScaleConfigsScaleRuleList) SetThresholdUpper(v int32) *ListAutoScalingConfigsResponseBodyDataScaleConfigsScaleRuleList {
	s.ThresholdUpper = &v
	return s
}

func (s *ListAutoScalingConfigsResponseBodyDataScaleConfigsScaleRuleList) SetTriggerCronExpr(v string) *ListAutoScalingConfigsResponseBodyDataScaleConfigsScaleRuleList {
	s.TriggerCronExpr = &v
	return s
}

type ListAutoScalingConfigsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAutoScalingConfigsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAutoScalingConfigsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAutoScalingConfigsResponse) GoString() string {
	return s.String()
}

func (s *ListAutoScalingConfigsResponse) SetHeaders(v map[string]*string) *ListAutoScalingConfigsResponse {
	s.Headers = v
	return s
}

func (s *ListAutoScalingConfigsResponse) SetStatusCode(v int32) *ListAutoScalingConfigsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAutoScalingConfigsResponse) SetBody(v *ListAutoScalingConfigsResponseBody) *ListAutoScalingConfigsResponse {
	s.Body = v
	return s
}

type ListAutoScalingRecordsRequest struct {
	// This parameter is required.
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// This parameter is required.
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s ListAutoScalingRecordsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAutoScalingRecordsRequest) GoString() string {
	return s.String()
}

func (s *ListAutoScalingRecordsRequest) SetInstanceId(v string) *ListAutoScalingRecordsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListAutoScalingRecordsRequest) SetOwnerAccount(v string) *ListAutoScalingRecordsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ListAutoScalingRecordsRequest) SetOwnerId(v int64) *ListAutoScalingRecordsRequest {
	s.OwnerId = &v
	return s
}

func (s *ListAutoScalingRecordsRequest) SetPageNum(v int32) *ListAutoScalingRecordsRequest {
	s.PageNum = &v
	return s
}

func (s *ListAutoScalingRecordsRequest) SetPageSize(v int32) *ListAutoScalingRecordsRequest {
	s.PageSize = &v
	return s
}

func (s *ListAutoScalingRecordsRequest) SetResourceOwnerAccount(v string) *ListAutoScalingRecordsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListAutoScalingRecordsRequest) SetResourceOwnerId(v int64) *ListAutoScalingRecordsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListAutoScalingRecordsRequest) SetSecurityToken(v string) *ListAutoScalingRecordsRequest {
	s.SecurityToken = &v
	return s
}

type ListAutoScalingRecordsResponseBody struct {
	AccessDeniedDetail *string                                 `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	Code               *string                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Data               *ListAutoScalingRecordsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	DynamicCode        *string                                 `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	DynamicMessage     *string                                 `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	HttpStatusCode     *int32                                  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message            *string                                 `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId          *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success            *bool                                   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListAutoScalingRecordsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAutoScalingRecordsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAutoScalingRecordsResponseBody) SetAccessDeniedDetail(v string) *ListAutoScalingRecordsResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ListAutoScalingRecordsResponseBody) SetCode(v string) *ListAutoScalingRecordsResponseBody {
	s.Code = &v
	return s
}

func (s *ListAutoScalingRecordsResponseBody) SetData(v *ListAutoScalingRecordsResponseBodyData) *ListAutoScalingRecordsResponseBody {
	s.Data = v
	return s
}

func (s *ListAutoScalingRecordsResponseBody) SetDynamicCode(v string) *ListAutoScalingRecordsResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *ListAutoScalingRecordsResponseBody) SetDynamicMessage(v string) *ListAutoScalingRecordsResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *ListAutoScalingRecordsResponseBody) SetHttpStatusCode(v int32) *ListAutoScalingRecordsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListAutoScalingRecordsResponseBody) SetMessage(v string) *ListAutoScalingRecordsResponseBody {
	s.Message = &v
	return s
}

func (s *ListAutoScalingRecordsResponseBody) SetRequestId(v string) *ListAutoScalingRecordsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAutoScalingRecordsResponseBody) SetSuccess(v bool) *ListAutoScalingRecordsResponseBody {
	s.Success = &v
	return s
}

type ListAutoScalingRecordsResponseBodyData struct {
	PageNum      *int32                                                `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize     *int32                                                `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ScaleRecords []*ListAutoScalingRecordsResponseBodyDataScaleRecords `json:"ScaleRecords,omitempty" xml:"ScaleRecords,omitempty" type:"Repeated"`
	TotalNum     *int32                                                `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
	TotalPage    *int32                                                `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
}

func (s ListAutoScalingRecordsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListAutoScalingRecordsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListAutoScalingRecordsResponseBodyData) SetPageNum(v int32) *ListAutoScalingRecordsResponseBodyData {
	s.PageNum = &v
	return s
}

func (s *ListAutoScalingRecordsResponseBodyData) SetPageSize(v int32) *ListAutoScalingRecordsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListAutoScalingRecordsResponseBodyData) SetScaleRecords(v []*ListAutoScalingRecordsResponseBodyDataScaleRecords) *ListAutoScalingRecordsResponseBodyData {
	s.ScaleRecords = v
	return s
}

func (s *ListAutoScalingRecordsResponseBodyData) SetTotalNum(v int32) *ListAutoScalingRecordsResponseBodyData {
	s.TotalNum = &v
	return s
}

func (s *ListAutoScalingRecordsResponseBodyData) SetTotalPage(v int32) *ListAutoScalingRecordsResponseBodyData {
	s.TotalPage = &v
	return s
}

type ListAutoScalingRecordsResponseBodyDataScaleRecords struct {
	Detail       *string `json:"Detail,omitempty" xml:"Detail,omitempty"`
	EndTime      *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Id           *string `json:"Id,omitempty" xml:"Id,omitempty"`
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OldValue     *string `json:"OldValue,omitempty" xml:"OldValue,omitempty"`
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	SpecGroupId  *string `json:"SpecGroupId,omitempty" xml:"SpecGroupId,omitempty"`
	StartTime    *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Strategy     *string `json:"Strategy,omitempty" xml:"Strategy,omitempty"`
	TargetValue  *string `json:"TargetValue,omitempty" xml:"TargetValue,omitempty"`
}

func (s ListAutoScalingRecordsResponseBodyDataScaleRecords) String() string {
	return tea.Prettify(s)
}

func (s ListAutoScalingRecordsResponseBodyDataScaleRecords) GoString() string {
	return s.String()
}

func (s *ListAutoScalingRecordsResponseBodyDataScaleRecords) SetDetail(v string) *ListAutoScalingRecordsResponseBodyDataScaleRecords {
	s.Detail = &v
	return s
}

func (s *ListAutoScalingRecordsResponseBodyDataScaleRecords) SetEndTime(v string) *ListAutoScalingRecordsResponseBodyDataScaleRecords {
	s.EndTime = &v
	return s
}

func (s *ListAutoScalingRecordsResponseBodyDataScaleRecords) SetId(v string) *ListAutoScalingRecordsResponseBodyDataScaleRecords {
	s.Id = &v
	return s
}

func (s *ListAutoScalingRecordsResponseBodyDataScaleRecords) SetInstanceId(v string) *ListAutoScalingRecordsResponseBodyDataScaleRecords {
	s.InstanceId = &v
	return s
}

func (s *ListAutoScalingRecordsResponseBodyDataScaleRecords) SetOldValue(v string) *ListAutoScalingRecordsResponseBodyDataScaleRecords {
	s.OldValue = &v
	return s
}

func (s *ListAutoScalingRecordsResponseBodyDataScaleRecords) SetResourceType(v string) *ListAutoScalingRecordsResponseBodyDataScaleRecords {
	s.ResourceType = &v
	return s
}

func (s *ListAutoScalingRecordsResponseBodyDataScaleRecords) SetSpecGroupId(v string) *ListAutoScalingRecordsResponseBodyDataScaleRecords {
	s.SpecGroupId = &v
	return s
}

func (s *ListAutoScalingRecordsResponseBodyDataScaleRecords) SetStartTime(v string) *ListAutoScalingRecordsResponseBodyDataScaleRecords {
	s.StartTime = &v
	return s
}

func (s *ListAutoScalingRecordsResponseBodyDataScaleRecords) SetStatus(v string) *ListAutoScalingRecordsResponseBodyDataScaleRecords {
	s.Status = &v
	return s
}

func (s *ListAutoScalingRecordsResponseBodyDataScaleRecords) SetStrategy(v string) *ListAutoScalingRecordsResponseBodyDataScaleRecords {
	s.Strategy = &v
	return s
}

func (s *ListAutoScalingRecordsResponseBodyDataScaleRecords) SetTargetValue(v string) *ListAutoScalingRecordsResponseBodyDataScaleRecords {
	s.TargetValue = &v
	return s
}

type ListAutoScalingRecordsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAutoScalingRecordsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAutoScalingRecordsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAutoScalingRecordsResponse) GoString() string {
	return s.String()
}

func (s *ListAutoScalingRecordsResponse) SetHeaders(v map[string]*string) *ListAutoScalingRecordsResponse {
	s.Headers = v
	return s
}

func (s *ListAutoScalingRecordsResponse) SetStatusCode(v int32) *ListAutoScalingRecordsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAutoScalingRecordsResponse) SetBody(v *ListAutoScalingRecordsResponseBody) *ListAutoScalingRecordsResponse {
	s.Body = v
	return s
}

type ListAutoScalingRulesRequest struct {
	// This parameter is required.
	ConfigId *string `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// This parameter is required.
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s ListAutoScalingRulesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAutoScalingRulesRequest) GoString() string {
	return s.String()
}

func (s *ListAutoScalingRulesRequest) SetConfigId(v string) *ListAutoScalingRulesRequest {
	s.ConfigId = &v
	return s
}

func (s *ListAutoScalingRulesRequest) SetInstanceId(v string) *ListAutoScalingRulesRequest {
	s.InstanceId = &v
	return s
}

func (s *ListAutoScalingRulesRequest) SetOwnerAccount(v string) *ListAutoScalingRulesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ListAutoScalingRulesRequest) SetOwnerId(v int64) *ListAutoScalingRulesRequest {
	s.OwnerId = &v
	return s
}

func (s *ListAutoScalingRulesRequest) SetResourceOwnerAccount(v string) *ListAutoScalingRulesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListAutoScalingRulesRequest) SetResourceOwnerId(v int64) *ListAutoScalingRulesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListAutoScalingRulesRequest) SetSecurityToken(v string) *ListAutoScalingRulesRequest {
	s.SecurityToken = &v
	return s
}

type ListAutoScalingRulesResponseBody struct {
	AccessDeniedDetail *string                               `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	Code               *string                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Data               *ListAutoScalingRulesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	DynamicCode        *string                               `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	DynamicMessage     *string                               `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	HttpStatusCode     *int32                                `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message            *string                               `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId          *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success            *bool                                 `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListAutoScalingRulesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAutoScalingRulesResponseBody) GoString() string {
	return s.String()
}

func (s *ListAutoScalingRulesResponseBody) SetAccessDeniedDetail(v string) *ListAutoScalingRulesResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ListAutoScalingRulesResponseBody) SetCode(v string) *ListAutoScalingRulesResponseBody {
	s.Code = &v
	return s
}

func (s *ListAutoScalingRulesResponseBody) SetData(v *ListAutoScalingRulesResponseBodyData) *ListAutoScalingRulesResponseBody {
	s.Data = v
	return s
}

func (s *ListAutoScalingRulesResponseBody) SetDynamicCode(v string) *ListAutoScalingRulesResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *ListAutoScalingRulesResponseBody) SetDynamicMessage(v string) *ListAutoScalingRulesResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *ListAutoScalingRulesResponseBody) SetHttpStatusCode(v int32) *ListAutoScalingRulesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListAutoScalingRulesResponseBody) SetMessage(v string) *ListAutoScalingRulesResponseBody {
	s.Message = &v
	return s
}

func (s *ListAutoScalingRulesResponseBody) SetRequestId(v string) *ListAutoScalingRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAutoScalingRulesResponseBody) SetSuccess(v bool) *ListAutoScalingRulesResponseBody {
	s.Success = &v
	return s
}

type ListAutoScalingRulesResponseBodyData struct {
	ScaleRules []*ListAutoScalingRulesResponseBodyDataScaleRules `json:"ScaleRules,omitempty" xml:"ScaleRules,omitempty" type:"Repeated"`
}

func (s ListAutoScalingRulesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListAutoScalingRulesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListAutoScalingRulesResponseBodyData) SetScaleRules(v []*ListAutoScalingRulesResponseBodyDataScaleRules) *ListAutoScalingRulesResponseBodyData {
	s.ScaleRules = v
	return s
}

type ListAutoScalingRulesResponseBodyDataScaleRules struct {
	ConfigId          *string `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	Enabled           *bool   `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	EndTime           *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	InstanceId        *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	ObservationWindow *int32  `json:"ObservationWindow,omitempty" xml:"ObservationWindow,omitempty"`
	OperationType     *string `json:"OperationType,omitempty" xml:"OperationType,omitempty"`
	RuleId            *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	RuleName          *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	RuleType          *string `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
	ScaleInStep       *int32  `json:"ScaleInStep,omitempty" xml:"ScaleInStep,omitempty"`
	ScaleOutStep      *int32  `json:"ScaleOutStep,omitempty" xml:"ScaleOutStep,omitempty"`
	SilenceTime       *int32  `json:"SilenceTime,omitempty" xml:"SilenceTime,omitempty"`
	StartTime         *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	TargetMetric      *string `json:"TargetMetric,omitempty" xml:"TargetMetric,omitempty"`
	TargetNodes       *int32  `json:"TargetNodes,omitempty" xml:"TargetNodes,omitempty"`
	ThresholdLower    *int32  `json:"ThresholdLower,omitempty" xml:"ThresholdLower,omitempty"`
	ThresholdUpper    *int32  `json:"ThresholdUpper,omitempty" xml:"ThresholdUpper,omitempty"`
	TriggerCronExpr   *string `json:"TriggerCronExpr,omitempty" xml:"TriggerCronExpr,omitempty"`
}

func (s ListAutoScalingRulesResponseBodyDataScaleRules) String() string {
	return tea.Prettify(s)
}

func (s ListAutoScalingRulesResponseBodyDataScaleRules) GoString() string {
	return s.String()
}

func (s *ListAutoScalingRulesResponseBodyDataScaleRules) SetConfigId(v string) *ListAutoScalingRulesResponseBodyDataScaleRules {
	s.ConfigId = &v
	return s
}

func (s *ListAutoScalingRulesResponseBodyDataScaleRules) SetEnabled(v bool) *ListAutoScalingRulesResponseBodyDataScaleRules {
	s.Enabled = &v
	return s
}

func (s *ListAutoScalingRulesResponseBodyDataScaleRules) SetEndTime(v string) *ListAutoScalingRulesResponseBodyDataScaleRules {
	s.EndTime = &v
	return s
}

func (s *ListAutoScalingRulesResponseBodyDataScaleRules) SetInstanceId(v string) *ListAutoScalingRulesResponseBodyDataScaleRules {
	s.InstanceId = &v
	return s
}

func (s *ListAutoScalingRulesResponseBodyDataScaleRules) SetObservationWindow(v int32) *ListAutoScalingRulesResponseBodyDataScaleRules {
	s.ObservationWindow = &v
	return s
}

func (s *ListAutoScalingRulesResponseBodyDataScaleRules) SetOperationType(v string) *ListAutoScalingRulesResponseBodyDataScaleRules {
	s.OperationType = &v
	return s
}

func (s *ListAutoScalingRulesResponseBodyDataScaleRules) SetRuleId(v string) *ListAutoScalingRulesResponseBodyDataScaleRules {
	s.RuleId = &v
	return s
}

func (s *ListAutoScalingRulesResponseBodyDataScaleRules) SetRuleName(v string) *ListAutoScalingRulesResponseBodyDataScaleRules {
	s.RuleName = &v
	return s
}

func (s *ListAutoScalingRulesResponseBodyDataScaleRules) SetRuleType(v string) *ListAutoScalingRulesResponseBodyDataScaleRules {
	s.RuleType = &v
	return s
}

func (s *ListAutoScalingRulesResponseBodyDataScaleRules) SetScaleInStep(v int32) *ListAutoScalingRulesResponseBodyDataScaleRules {
	s.ScaleInStep = &v
	return s
}

func (s *ListAutoScalingRulesResponseBodyDataScaleRules) SetScaleOutStep(v int32) *ListAutoScalingRulesResponseBodyDataScaleRules {
	s.ScaleOutStep = &v
	return s
}

func (s *ListAutoScalingRulesResponseBodyDataScaleRules) SetSilenceTime(v int32) *ListAutoScalingRulesResponseBodyDataScaleRules {
	s.SilenceTime = &v
	return s
}

func (s *ListAutoScalingRulesResponseBodyDataScaleRules) SetStartTime(v string) *ListAutoScalingRulesResponseBodyDataScaleRules {
	s.StartTime = &v
	return s
}

func (s *ListAutoScalingRulesResponseBodyDataScaleRules) SetTargetMetric(v string) *ListAutoScalingRulesResponseBodyDataScaleRules {
	s.TargetMetric = &v
	return s
}

func (s *ListAutoScalingRulesResponseBodyDataScaleRules) SetTargetNodes(v int32) *ListAutoScalingRulesResponseBodyDataScaleRules {
	s.TargetNodes = &v
	return s
}

func (s *ListAutoScalingRulesResponseBodyDataScaleRules) SetThresholdLower(v int32) *ListAutoScalingRulesResponseBodyDataScaleRules {
	s.ThresholdLower = &v
	return s
}

func (s *ListAutoScalingRulesResponseBodyDataScaleRules) SetThresholdUpper(v int32) *ListAutoScalingRulesResponseBodyDataScaleRules {
	s.ThresholdUpper = &v
	return s
}

func (s *ListAutoScalingRulesResponseBodyDataScaleRules) SetTriggerCronExpr(v string) *ListAutoScalingRulesResponseBodyDataScaleRules {
	s.TriggerCronExpr = &v
	return s
}

type ListAutoScalingRulesResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAutoScalingRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAutoScalingRulesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAutoScalingRulesResponse) GoString() string {
	return s.String()
}

func (s *ListAutoScalingRulesResponse) SetHeaders(v map[string]*string) *ListAutoScalingRulesResponse {
	s.Headers = v
	return s
}

func (s *ListAutoScalingRulesResponse) SetStatusCode(v int32) *ListAutoScalingRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAutoScalingRulesResponse) SetBody(v *ListAutoScalingRulesResponseBody) *ListAutoScalingRulesResponse {
	s.Body = v
	return s
}

type ListLdpsComputeGroupsRequest struct {
	// This parameter is required.
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s ListLdpsComputeGroupsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListLdpsComputeGroupsRequest) GoString() string {
	return s.String()
}

func (s *ListLdpsComputeGroupsRequest) SetInstanceId(v string) *ListLdpsComputeGroupsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListLdpsComputeGroupsRequest) SetOwnerAccount(v string) *ListLdpsComputeGroupsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ListLdpsComputeGroupsRequest) SetOwnerId(v int64) *ListLdpsComputeGroupsRequest {
	s.OwnerId = &v
	return s
}

func (s *ListLdpsComputeGroupsRequest) SetRegionId(v string) *ListLdpsComputeGroupsRequest {
	s.RegionId = &v
	return s
}

func (s *ListLdpsComputeGroupsRequest) SetResourceOwnerAccount(v string) *ListLdpsComputeGroupsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListLdpsComputeGroupsRequest) SetResourceOwnerId(v int64) *ListLdpsComputeGroupsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListLdpsComputeGroupsRequest) SetSecurityToken(v string) *ListLdpsComputeGroupsRequest {
	s.SecurityToken = &v
	return s
}

type ListLdpsComputeGroupsResponseBody struct {
	AccessDeniedDetail *string                                       `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	GroupList          []*ListLdpsComputeGroupsResponseBodyGroupList `json:"GroupList,omitempty" xml:"GroupList,omitempty" type:"Repeated"`
	RequestId          *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListLdpsComputeGroupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListLdpsComputeGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *ListLdpsComputeGroupsResponseBody) SetAccessDeniedDetail(v string) *ListLdpsComputeGroupsResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ListLdpsComputeGroupsResponseBody) SetGroupList(v []*ListLdpsComputeGroupsResponseBodyGroupList) *ListLdpsComputeGroupsResponseBody {
	s.GroupList = v
	return s
}

func (s *ListLdpsComputeGroupsResponseBody) SetRequestId(v string) *ListLdpsComputeGroupsResponseBody {
	s.RequestId = &v
	return s
}

type ListLdpsComputeGroupsResponseBodyGroupList struct {
	ExceptionInfo *string                `json:"ExceptionInfo,omitempty" xml:"ExceptionInfo,omitempty"`
	GroupName     *string                `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	IsDefault     *bool                  `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	Properties    map[string]interface{} `json:"Properties,omitempty" xml:"Properties,omitempty"`
	State         *string                `json:"State,omitempty" xml:"State,omitempty"`
	WebUI         *string                `json:"WebUI,omitempty" xml:"WebUI,omitempty"`
}

func (s ListLdpsComputeGroupsResponseBodyGroupList) String() string {
	return tea.Prettify(s)
}

func (s ListLdpsComputeGroupsResponseBodyGroupList) GoString() string {
	return s.String()
}

func (s *ListLdpsComputeGroupsResponseBodyGroupList) SetExceptionInfo(v string) *ListLdpsComputeGroupsResponseBodyGroupList {
	s.ExceptionInfo = &v
	return s
}

func (s *ListLdpsComputeGroupsResponseBodyGroupList) SetGroupName(v string) *ListLdpsComputeGroupsResponseBodyGroupList {
	s.GroupName = &v
	return s
}

func (s *ListLdpsComputeGroupsResponseBodyGroupList) SetIsDefault(v bool) *ListLdpsComputeGroupsResponseBodyGroupList {
	s.IsDefault = &v
	return s
}

func (s *ListLdpsComputeGroupsResponseBodyGroupList) SetProperties(v map[string]interface{}) *ListLdpsComputeGroupsResponseBodyGroupList {
	s.Properties = v
	return s
}

func (s *ListLdpsComputeGroupsResponseBodyGroupList) SetState(v string) *ListLdpsComputeGroupsResponseBodyGroupList {
	s.State = &v
	return s
}

func (s *ListLdpsComputeGroupsResponseBodyGroupList) SetWebUI(v string) *ListLdpsComputeGroupsResponseBodyGroupList {
	s.WebUI = &v
	return s
}

type ListLdpsComputeGroupsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListLdpsComputeGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListLdpsComputeGroupsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListLdpsComputeGroupsResponse) GoString() string {
	return s.String()
}

func (s *ListLdpsComputeGroupsResponse) SetHeaders(v map[string]*string) *ListLdpsComputeGroupsResponse {
	s.Headers = v
	return s
}

func (s *ListLdpsComputeGroupsResponse) SetStatusCode(v int32) *ListLdpsComputeGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListLdpsComputeGroupsResponse) SetBody(v *ListLdpsComputeGroupsResponseBody) *ListLdpsComputeGroupsResponse {
	s.Body = v
	return s
}

type ListTagResourcesRequest struct {
	// The token used to start the next query to retrieve more results.
	//
	// > This parameter is not required in the first query. If not all results are returned in one query, you can pass in the **NextToken*	- value returned for the query to perform the next query.
	//
	// example:
	//
	// 212db86****
	NextToken    *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region in which the instances whose tags you want to query are located. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/426062.html) operation to query the region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The list of resource IDs.
	ResourceId           []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	ResourceOwnerAccount *string   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64    `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The resource type. Set the value to **INSTANCE**.
	//
	// This parameter is required.
	//
	// example:
	//
	// INSTANCE
	ResourceType  *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The list of tags associated with the instances you want to query.
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

func (s *ListTagResourcesRequest) SetOwnerAccount(v string) *ListTagResourcesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ListTagResourcesRequest) SetOwnerId(v int64) *ListTagResourcesRequest {
	s.OwnerId = &v
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

func (s *ListTagResourcesRequest) SetResourceOwnerAccount(v string) *ListTagResourcesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListTagResourcesRequest) SetResourceOwnerId(v int64) *ListTagResourcesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListTagResourcesRequest) SetResourceType(v string) *ListTagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *ListTagResourcesRequest) SetSecurityToken(v string) *ListTagResourcesRequest {
	s.SecurityToken = &v
	return s
}

func (s *ListTagResourcesRequest) SetTag(v []*ListTagResourcesRequestTag) *ListTagResourcesRequest {
	s.Tag = v
	return s
}

type ListTagResourcesRequestTag struct {
	// The keys of the tags associated with the instances you want to query.
	//
	// > You can specify the keys of multiple tags. For example, you can specify the key of the first tag in the first key-value pair contained in the value of this parameter and specify the key of the second tag in the second key-value pair.
	//
	// example:
	//
	// test
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The values of the tags associated with the instances you want to query.
	//
	// > You can specify the values of multiple tags. For example, you can specify the value of the first tag in the first key-value pair contained in the value of this parameter and specify the value of the second tag in the second key-value pair.
	//
	// example:
	//
	// 2.2.8
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
	// The token used to start the next query.
	//
	// > If not all results are returned in the first query, this parameter is returned. You can pass in the returned value of this parameter for the next query.
	//
	// example:
	//
	// 212db86****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 05CB115C-91CB-529F-9098-50C1F6CB3BD3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The list of resources.
	TagResources []*ListTagResourcesResponseBodyTagResources `json:"TagResources,omitempty" xml:"TagResources,omitempty" type:"Repeated"`
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

func (s *ListTagResourcesResponseBody) SetTagResources(v []*ListTagResourcesResponseBodyTagResources) *ListTagResourcesResponseBody {
	s.TagResources = v
	return s
}

type ListTagResourcesResponseBodyTagResources struct {
	// The ID of the resource, which is the ID of the instance.
	//
	// example:
	//
	// ld-bp17j28j2y7pm****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The type of the resources. The returned value is fixed to **ALIYUN::HITSDB::INSTANCE**.
	//
	// example:
	//
	// ALIYUN::HITSDB::INSTANCE
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The key of the tag associated with the instance.
	//
	// example:
	//
	// test
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The value of the tag associated with the instance.
	//
	// example:
	//
	// 2.2.8
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

type ModifyAutoScalingConfigRequest struct {
	// This parameter is required.
	ConfigId           *string `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	ConfigName         *string `json:"ConfigName,omitempty" xml:"ConfigName,omitempty"`
	EffectiveTimeEnd   *string `json:"EffectiveTimeEnd,omitempty" xml:"EffectiveTimeEnd,omitempty"`
	EffectiveTimeStart *string `json:"EffectiveTimeStart,omitempty" xml:"EffectiveTimeStart,omitempty"`
	Enabled            *bool   `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	Engine             *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// This parameter is required.
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	NodesMax             *int32  `json:"NodesMax,omitempty" xml:"NodesMax,omitempty"`
	NodesMin             *int32  `json:"NodesMin,omitempty" xml:"NodesMin,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ScaleType            *string `json:"ScaleType,omitempty" xml:"ScaleType,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	SpecId               *string `json:"SpecId,omitempty" xml:"SpecId,omitempty"`
}

func (s ModifyAutoScalingConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyAutoScalingConfigRequest) GoString() string {
	return s.String()
}

func (s *ModifyAutoScalingConfigRequest) SetConfigId(v string) *ModifyAutoScalingConfigRequest {
	s.ConfigId = &v
	return s
}

func (s *ModifyAutoScalingConfigRequest) SetConfigName(v string) *ModifyAutoScalingConfigRequest {
	s.ConfigName = &v
	return s
}

func (s *ModifyAutoScalingConfigRequest) SetEffectiveTimeEnd(v string) *ModifyAutoScalingConfigRequest {
	s.EffectiveTimeEnd = &v
	return s
}

func (s *ModifyAutoScalingConfigRequest) SetEffectiveTimeStart(v string) *ModifyAutoScalingConfigRequest {
	s.EffectiveTimeStart = &v
	return s
}

func (s *ModifyAutoScalingConfigRequest) SetEnabled(v bool) *ModifyAutoScalingConfigRequest {
	s.Enabled = &v
	return s
}

func (s *ModifyAutoScalingConfigRequest) SetEngine(v string) *ModifyAutoScalingConfigRequest {
	s.Engine = &v
	return s
}

func (s *ModifyAutoScalingConfigRequest) SetInstanceId(v string) *ModifyAutoScalingConfigRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyAutoScalingConfigRequest) SetNodesMax(v int32) *ModifyAutoScalingConfigRequest {
	s.NodesMax = &v
	return s
}

func (s *ModifyAutoScalingConfigRequest) SetNodesMin(v int32) *ModifyAutoScalingConfigRequest {
	s.NodesMin = &v
	return s
}

func (s *ModifyAutoScalingConfigRequest) SetOwnerAccount(v string) *ModifyAutoScalingConfigRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyAutoScalingConfigRequest) SetOwnerId(v int64) *ModifyAutoScalingConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyAutoScalingConfigRequest) SetResourceOwnerAccount(v string) *ModifyAutoScalingConfigRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyAutoScalingConfigRequest) SetResourceOwnerId(v int64) *ModifyAutoScalingConfigRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyAutoScalingConfigRequest) SetScaleType(v string) *ModifyAutoScalingConfigRequest {
	s.ScaleType = &v
	return s
}

func (s *ModifyAutoScalingConfigRequest) SetSecurityToken(v string) *ModifyAutoScalingConfigRequest {
	s.SecurityToken = &v
	return s
}

func (s *ModifyAutoScalingConfigRequest) SetSpecId(v string) *ModifyAutoScalingConfigRequest {
	s.SpecId = &v
	return s
}

type ModifyAutoScalingConfigResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyAutoScalingConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyAutoScalingConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyAutoScalingConfigResponseBody) SetCode(v string) *ModifyAutoScalingConfigResponseBody {
	s.Code = &v
	return s
}

func (s *ModifyAutoScalingConfigResponseBody) SetHttpStatusCode(v int32) *ModifyAutoScalingConfigResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModifyAutoScalingConfigResponseBody) SetMessage(v string) *ModifyAutoScalingConfigResponseBody {
	s.Message = &v
	return s
}

func (s *ModifyAutoScalingConfigResponseBody) SetRequestId(v string) *ModifyAutoScalingConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyAutoScalingConfigResponseBody) SetSuccess(v bool) *ModifyAutoScalingConfigResponseBody {
	s.Success = &v
	return s
}

type ModifyAutoScalingConfigResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyAutoScalingConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyAutoScalingConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyAutoScalingConfigResponse) GoString() string {
	return s.String()
}

func (s *ModifyAutoScalingConfigResponse) SetHeaders(v map[string]*string) *ModifyAutoScalingConfigResponse {
	s.Headers = v
	return s
}

func (s *ModifyAutoScalingConfigResponse) SetStatusCode(v int32) *ModifyAutoScalingConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyAutoScalingConfigResponse) SetBody(v *ModifyAutoScalingConfigResponseBody) *ModifyAutoScalingConfigResponse {
	s.Body = v
	return s
}

type ModifyAutoScalingRuleRequest struct {
	// This parameter is required.
	ConfigId *string `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	Enabled  *bool   `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	EndTime  *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// This parameter is required.
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	ObservationWindow    *int32  `json:"ObservationWindow,omitempty" xml:"ObservationWindow,omitempty"`
	OperationType        *string `json:"OperationType,omitempty" xml:"OperationType,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// This parameter is required.
	RuleId          *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	RuleName        *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	RuleType        *string `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
	ScaleInStep     *int32  `json:"ScaleInStep,omitempty" xml:"ScaleInStep,omitempty"`
	ScaleOutStep    *int32  `json:"ScaleOutStep,omitempty" xml:"ScaleOutStep,omitempty"`
	SecurityToken   *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	SilenceTime     *int32  `json:"SilenceTime,omitempty" xml:"SilenceTime,omitempty"`
	StartTime       *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	TargetMetric    *string `json:"TargetMetric,omitempty" xml:"TargetMetric,omitempty"`
	TargetNodes     *int32  `json:"TargetNodes,omitempty" xml:"TargetNodes,omitempty"`
	ThresholdLower  *int32  `json:"ThresholdLower,omitempty" xml:"ThresholdLower,omitempty"`
	ThresholdUpper  *int32  `json:"ThresholdUpper,omitempty" xml:"ThresholdUpper,omitempty"`
	TriggerCronExpr *string `json:"TriggerCronExpr,omitempty" xml:"TriggerCronExpr,omitempty"`
}

func (s ModifyAutoScalingRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyAutoScalingRuleRequest) GoString() string {
	return s.String()
}

func (s *ModifyAutoScalingRuleRequest) SetConfigId(v string) *ModifyAutoScalingRuleRequest {
	s.ConfigId = &v
	return s
}

func (s *ModifyAutoScalingRuleRequest) SetEnabled(v bool) *ModifyAutoScalingRuleRequest {
	s.Enabled = &v
	return s
}

func (s *ModifyAutoScalingRuleRequest) SetEndTime(v string) *ModifyAutoScalingRuleRequest {
	s.EndTime = &v
	return s
}

func (s *ModifyAutoScalingRuleRequest) SetInstanceId(v string) *ModifyAutoScalingRuleRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyAutoScalingRuleRequest) SetObservationWindow(v int32) *ModifyAutoScalingRuleRequest {
	s.ObservationWindow = &v
	return s
}

func (s *ModifyAutoScalingRuleRequest) SetOperationType(v string) *ModifyAutoScalingRuleRequest {
	s.OperationType = &v
	return s
}

func (s *ModifyAutoScalingRuleRequest) SetOwnerAccount(v string) *ModifyAutoScalingRuleRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyAutoScalingRuleRequest) SetOwnerId(v int64) *ModifyAutoScalingRuleRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyAutoScalingRuleRequest) SetResourceOwnerAccount(v string) *ModifyAutoScalingRuleRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyAutoScalingRuleRequest) SetResourceOwnerId(v int64) *ModifyAutoScalingRuleRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyAutoScalingRuleRequest) SetRuleId(v string) *ModifyAutoScalingRuleRequest {
	s.RuleId = &v
	return s
}

func (s *ModifyAutoScalingRuleRequest) SetRuleName(v string) *ModifyAutoScalingRuleRequest {
	s.RuleName = &v
	return s
}

func (s *ModifyAutoScalingRuleRequest) SetRuleType(v string) *ModifyAutoScalingRuleRequest {
	s.RuleType = &v
	return s
}

func (s *ModifyAutoScalingRuleRequest) SetScaleInStep(v int32) *ModifyAutoScalingRuleRequest {
	s.ScaleInStep = &v
	return s
}

func (s *ModifyAutoScalingRuleRequest) SetScaleOutStep(v int32) *ModifyAutoScalingRuleRequest {
	s.ScaleOutStep = &v
	return s
}

func (s *ModifyAutoScalingRuleRequest) SetSecurityToken(v string) *ModifyAutoScalingRuleRequest {
	s.SecurityToken = &v
	return s
}

func (s *ModifyAutoScalingRuleRequest) SetSilenceTime(v int32) *ModifyAutoScalingRuleRequest {
	s.SilenceTime = &v
	return s
}

func (s *ModifyAutoScalingRuleRequest) SetStartTime(v string) *ModifyAutoScalingRuleRequest {
	s.StartTime = &v
	return s
}

func (s *ModifyAutoScalingRuleRequest) SetTargetMetric(v string) *ModifyAutoScalingRuleRequest {
	s.TargetMetric = &v
	return s
}

func (s *ModifyAutoScalingRuleRequest) SetTargetNodes(v int32) *ModifyAutoScalingRuleRequest {
	s.TargetNodes = &v
	return s
}

func (s *ModifyAutoScalingRuleRequest) SetThresholdLower(v int32) *ModifyAutoScalingRuleRequest {
	s.ThresholdLower = &v
	return s
}

func (s *ModifyAutoScalingRuleRequest) SetThresholdUpper(v int32) *ModifyAutoScalingRuleRequest {
	s.ThresholdUpper = &v
	return s
}

func (s *ModifyAutoScalingRuleRequest) SetTriggerCronExpr(v string) *ModifyAutoScalingRuleRequest {
	s.TriggerCronExpr = &v
	return s
}

type ModifyAutoScalingRuleResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyAutoScalingRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyAutoScalingRuleResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyAutoScalingRuleResponseBody) SetCode(v string) *ModifyAutoScalingRuleResponseBody {
	s.Code = &v
	return s
}

func (s *ModifyAutoScalingRuleResponseBody) SetHttpStatusCode(v int32) *ModifyAutoScalingRuleResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModifyAutoScalingRuleResponseBody) SetMessage(v string) *ModifyAutoScalingRuleResponseBody {
	s.Message = &v
	return s
}

func (s *ModifyAutoScalingRuleResponseBody) SetRequestId(v string) *ModifyAutoScalingRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyAutoScalingRuleResponseBody) SetSuccess(v bool) *ModifyAutoScalingRuleResponseBody {
	s.Success = &v
	return s
}

type ModifyAutoScalingRuleResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyAutoScalingRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyAutoScalingRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyAutoScalingRuleResponse) GoString() string {
	return s.String()
}

func (s *ModifyAutoScalingRuleResponse) SetHeaders(v map[string]*string) *ModifyAutoScalingRuleResponse {
	s.Headers = v
	return s
}

func (s *ModifyAutoScalingRuleResponse) SetStatusCode(v int32) *ModifyAutoScalingRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyAutoScalingRuleResponse) SetBody(v *ModifyAutoScalingRuleResponseBody) *ModifyAutoScalingRuleResponse {
	s.Body = v
	return s
}

type ModifyInstancePayTypeRequest struct {
	// The subscription duration of the instance. The parameter is required if the instance is an subscription instance.
	//
	// 	- If PricingCycle is set to Month, set this parameter to an integer that ranges from 1 to 9.
	//
	// 	- If PricingCycle is set to Year, set this parameter to an integer that ranges from 1 to 3.
	//
	// example:
	//
	// 1
	Duration *int32 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// ld-bp1z3506imz2f****
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The billing method of the instance. Valid values:
	//
	// 	- **PREPAY**: subscription.
	//
	// 	- **POSTPAY**: pay-as-you-go.
	//
	// This parameter is required.
	//
	// example:
	//
	// POSTPAY
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// The unit of the subscription duration for the instance. Valid values:
	//
	// 	- Month
	//
	// 	- Year
	//
	// example:
	//
	// Month
	PricingCycle         *string `json:"PricingCycle,omitempty" xml:"PricingCycle,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s ModifyInstancePayTypeRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyInstancePayTypeRequest) GoString() string {
	return s.String()
}

func (s *ModifyInstancePayTypeRequest) SetDuration(v int32) *ModifyInstancePayTypeRequest {
	s.Duration = &v
	return s
}

func (s *ModifyInstancePayTypeRequest) SetInstanceId(v string) *ModifyInstancePayTypeRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyInstancePayTypeRequest) SetOwnerAccount(v string) *ModifyInstancePayTypeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyInstancePayTypeRequest) SetOwnerId(v int64) *ModifyInstancePayTypeRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyInstancePayTypeRequest) SetPayType(v string) *ModifyInstancePayTypeRequest {
	s.PayType = &v
	return s
}

func (s *ModifyInstancePayTypeRequest) SetPricingCycle(v string) *ModifyInstancePayTypeRequest {
	s.PricingCycle = &v
	return s
}

func (s *ModifyInstancePayTypeRequest) SetResourceOwnerAccount(v string) *ModifyInstancePayTypeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyInstancePayTypeRequest) SetResourceOwnerId(v int64) *ModifyInstancePayTypeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyInstancePayTypeRequest) SetSecurityToken(v string) *ModifyInstancePayTypeRequest {
	s.SecurityToken = &v
	return s
}

type ModifyInstancePayTypeResponseBody struct {
	// The detailed reason why the access was denied.
	//
	// example:
	//
	// {"AuthAction":"xxx","AuthPrincipalDisplayName":"222","AuthPrincipalOwnerId":"111","AuthPrincipalType":"SubUser",,"NoPermissionType":"ImplicitDeny","PolicyType":"AccountLevelIdentityBasedPolicy","EncodedDiagnosticMessage":"xxxxxx"}
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// ld-bp1z3506imz2f****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the order.
	//
	// example:
	//
	// 211662251220224
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 587BCA54-50DA-4885-ADE9-80A848339151
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyInstancePayTypeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyInstancePayTypeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyInstancePayTypeResponseBody) SetAccessDeniedDetail(v string) *ModifyInstancePayTypeResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ModifyInstancePayTypeResponseBody) SetInstanceId(v string) *ModifyInstancePayTypeResponseBody {
	s.InstanceId = &v
	return s
}

func (s *ModifyInstancePayTypeResponseBody) SetOrderId(v int64) *ModifyInstancePayTypeResponseBody {
	s.OrderId = &v
	return s
}

func (s *ModifyInstancePayTypeResponseBody) SetRequestId(v string) *ModifyInstancePayTypeResponseBody {
	s.RequestId = &v
	return s
}

type ModifyInstancePayTypeResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyInstancePayTypeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyInstancePayTypeResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyInstancePayTypeResponse) GoString() string {
	return s.String()
}

func (s *ModifyInstancePayTypeResponse) SetHeaders(v map[string]*string) *ModifyInstancePayTypeResponse {
	s.Headers = v
	return s
}

func (s *ModifyInstancePayTypeResponse) SetStatusCode(v int32) *ModifyInstancePayTypeResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyInstancePayTypeResponse) SetBody(v *ModifyInstancePayTypeResponseBody) *ModifyInstancePayTypeResponse {
	s.Body = v
	return s
}

type ModifyLindormV2InstanceRequest struct {
	CloudStorageSize *int64  `json:"CloudStorageSize,omitempty" xml:"CloudStorageSize,omitempty"`
	CloudStorageType *string `json:"CloudStorageType,omitempty" xml:"CloudStorageType,omitempty"`
	EngineType       *string `json:"EngineType,omitempty" xml:"EngineType,omitempty"`
	// This parameter is required.
	InstanceId    *string                                        `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	NodeGroupList []*ModifyLindormV2InstanceRequestNodeGroupList `json:"NodeGroupList,omitempty" xml:"NodeGroupList,omitempty" type:"Repeated"`
	OwnerAccount  *string                                        `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId       *int64                                         `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// This parameter is required.
	UpgradeType *string `json:"UpgradeType,omitempty" xml:"UpgradeType,omitempty"`
}

func (s ModifyLindormV2InstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyLindormV2InstanceRequest) GoString() string {
	return s.String()
}

func (s *ModifyLindormV2InstanceRequest) SetCloudStorageSize(v int64) *ModifyLindormV2InstanceRequest {
	s.CloudStorageSize = &v
	return s
}

func (s *ModifyLindormV2InstanceRequest) SetCloudStorageType(v string) *ModifyLindormV2InstanceRequest {
	s.CloudStorageType = &v
	return s
}

func (s *ModifyLindormV2InstanceRequest) SetEngineType(v string) *ModifyLindormV2InstanceRequest {
	s.EngineType = &v
	return s
}

func (s *ModifyLindormV2InstanceRequest) SetInstanceId(v string) *ModifyLindormV2InstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyLindormV2InstanceRequest) SetNodeGroupList(v []*ModifyLindormV2InstanceRequestNodeGroupList) *ModifyLindormV2InstanceRequest {
	s.NodeGroupList = v
	return s
}

func (s *ModifyLindormV2InstanceRequest) SetOwnerAccount(v string) *ModifyLindormV2InstanceRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyLindormV2InstanceRequest) SetOwnerId(v int64) *ModifyLindormV2InstanceRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyLindormV2InstanceRequest) SetRegionId(v string) *ModifyLindormV2InstanceRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyLindormV2InstanceRequest) SetResourceOwnerAccount(v string) *ModifyLindormV2InstanceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyLindormV2InstanceRequest) SetResourceOwnerId(v int64) *ModifyLindormV2InstanceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyLindormV2InstanceRequest) SetSecurityToken(v string) *ModifyLindormV2InstanceRequest {
	s.SecurityToken = &v
	return s
}

func (s *ModifyLindormV2InstanceRequest) SetUpgradeType(v string) *ModifyLindormV2InstanceRequest {
	s.UpgradeType = &v
	return s
}

type ModifyLindormV2InstanceRequestNodeGroupList struct {
	GroupId           *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	NodeCount         *string `json:"NodeCount,omitempty" xml:"NodeCount,omitempty"`
	NodeDiskSize      *int64  `json:"NodeDiskSize,omitempty" xml:"NodeDiskSize,omitempty"`
	NodeDiskType      *string `json:"NodeDiskType,omitempty" xml:"NodeDiskType,omitempty"`
	NodeSpec          *string `json:"NodeSpec,omitempty" xml:"NodeSpec,omitempty"`
	ResourceGroupName *string `json:"ResourceGroupName,omitempty" xml:"ResourceGroupName,omitempty"`
}

func (s ModifyLindormV2InstanceRequestNodeGroupList) String() string {
	return tea.Prettify(s)
}

func (s ModifyLindormV2InstanceRequestNodeGroupList) GoString() string {
	return s.String()
}

func (s *ModifyLindormV2InstanceRequestNodeGroupList) SetGroupId(v string) *ModifyLindormV2InstanceRequestNodeGroupList {
	s.GroupId = &v
	return s
}

func (s *ModifyLindormV2InstanceRequestNodeGroupList) SetNodeCount(v string) *ModifyLindormV2InstanceRequestNodeGroupList {
	s.NodeCount = &v
	return s
}

func (s *ModifyLindormV2InstanceRequestNodeGroupList) SetNodeDiskSize(v int64) *ModifyLindormV2InstanceRequestNodeGroupList {
	s.NodeDiskSize = &v
	return s
}

func (s *ModifyLindormV2InstanceRequestNodeGroupList) SetNodeDiskType(v string) *ModifyLindormV2InstanceRequestNodeGroupList {
	s.NodeDiskType = &v
	return s
}

func (s *ModifyLindormV2InstanceRequestNodeGroupList) SetNodeSpec(v string) *ModifyLindormV2InstanceRequestNodeGroupList {
	s.NodeSpec = &v
	return s
}

func (s *ModifyLindormV2InstanceRequestNodeGroupList) SetResourceGroupName(v string) *ModifyLindormV2InstanceRequestNodeGroupList {
	s.ResourceGroupName = &v
	return s
}

type ModifyLindormV2InstanceResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	InstanceId         *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OrderId            *int64  `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	RequestId          *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyLindormV2InstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyLindormV2InstanceResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyLindormV2InstanceResponseBody) SetAccessDeniedDetail(v string) *ModifyLindormV2InstanceResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ModifyLindormV2InstanceResponseBody) SetInstanceId(v string) *ModifyLindormV2InstanceResponseBody {
	s.InstanceId = &v
	return s
}

func (s *ModifyLindormV2InstanceResponseBody) SetOrderId(v int64) *ModifyLindormV2InstanceResponseBody {
	s.OrderId = &v
	return s
}

func (s *ModifyLindormV2InstanceResponseBody) SetRequestId(v string) *ModifyLindormV2InstanceResponseBody {
	s.RequestId = &v
	return s
}

type ModifyLindormV2InstanceResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyLindormV2InstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyLindormV2InstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyLindormV2InstanceResponse) GoString() string {
	return s.String()
}

func (s *ModifyLindormV2InstanceResponse) SetHeaders(v map[string]*string) *ModifyLindormV2InstanceResponse {
	s.Headers = v
	return s
}

func (s *ModifyLindormV2InstanceResponse) SetStatusCode(v int32) *ModifyLindormV2InstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyLindormV2InstanceResponse) SetBody(v *ModifyLindormV2InstanceResponseBody) *ModifyLindormV2InstanceResponse {
	s.Body = v
	return s
}

type ModifyLindormV2WhiteIpListRequest struct {
	DeleteGroup *bool `json:"DeleteGroup,omitempty" xml:"DeleteGroup,omitempty"`
	// This parameter is required.
	GroupName            *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// This parameter is required.
	WhiteIpList *string `json:"WhiteIpList,omitempty" xml:"WhiteIpList,omitempty"`
}

func (s ModifyLindormV2WhiteIpListRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyLindormV2WhiteIpListRequest) GoString() string {
	return s.String()
}

func (s *ModifyLindormV2WhiteIpListRequest) SetDeleteGroup(v bool) *ModifyLindormV2WhiteIpListRequest {
	s.DeleteGroup = &v
	return s
}

func (s *ModifyLindormV2WhiteIpListRequest) SetGroupName(v string) *ModifyLindormV2WhiteIpListRequest {
	s.GroupName = &v
	return s
}

func (s *ModifyLindormV2WhiteIpListRequest) SetInstanceId(v string) *ModifyLindormV2WhiteIpListRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyLindormV2WhiteIpListRequest) SetOwnerAccount(v string) *ModifyLindormV2WhiteIpListRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyLindormV2WhiteIpListRequest) SetOwnerId(v int64) *ModifyLindormV2WhiteIpListRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyLindormV2WhiteIpListRequest) SetRegionId(v string) *ModifyLindormV2WhiteIpListRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyLindormV2WhiteIpListRequest) SetResourceOwnerAccount(v string) *ModifyLindormV2WhiteIpListRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyLindormV2WhiteIpListRequest) SetResourceOwnerId(v int64) *ModifyLindormV2WhiteIpListRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyLindormV2WhiteIpListRequest) SetSecurityToken(v string) *ModifyLindormV2WhiteIpListRequest {
	s.SecurityToken = &v
	return s
}

func (s *ModifyLindormV2WhiteIpListRequest) SetWhiteIpList(v string) *ModifyLindormV2WhiteIpListRequest {
	s.WhiteIpList = &v
	return s
}

type ModifyLindormV2WhiteIpListResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	RequestId          *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyLindormV2WhiteIpListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyLindormV2WhiteIpListResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyLindormV2WhiteIpListResponseBody) SetAccessDeniedDetail(v string) *ModifyLindormV2WhiteIpListResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ModifyLindormV2WhiteIpListResponseBody) SetRequestId(v string) *ModifyLindormV2WhiteIpListResponseBody {
	s.RequestId = &v
	return s
}

type ModifyLindormV2WhiteIpListResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyLindormV2WhiteIpListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyLindormV2WhiteIpListResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyLindormV2WhiteIpListResponse) GoString() string {
	return s.String()
}

func (s *ModifyLindormV2WhiteIpListResponse) SetHeaders(v map[string]*string) *ModifyLindormV2WhiteIpListResponse {
	s.Headers = v
	return s
}

func (s *ModifyLindormV2WhiteIpListResponse) SetStatusCode(v int32) *ModifyLindormV2WhiteIpListResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyLindormV2WhiteIpListResponse) SetBody(v *ModifyLindormV2WhiteIpListResponseBody) *ModifyLindormV2WhiteIpListResponse {
	s.Body = v
	return s
}

type OpenComputeEngineRequest struct {
	CpuLimit *string `json:"CpuLimit,omitempty" xml:"CpuLimit,omitempty"`
	// This parameter is required.
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	MemoryLimit          *string `json:"MemoryLimit,omitempty" xml:"MemoryLimit,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s OpenComputeEngineRequest) String() string {
	return tea.Prettify(s)
}

func (s OpenComputeEngineRequest) GoString() string {
	return s.String()
}

func (s *OpenComputeEngineRequest) SetCpuLimit(v string) *OpenComputeEngineRequest {
	s.CpuLimit = &v
	return s
}

func (s *OpenComputeEngineRequest) SetInstanceId(v string) *OpenComputeEngineRequest {
	s.InstanceId = &v
	return s
}

func (s *OpenComputeEngineRequest) SetMemoryLimit(v string) *OpenComputeEngineRequest {
	s.MemoryLimit = &v
	return s
}

func (s *OpenComputeEngineRequest) SetOwnerAccount(v string) *OpenComputeEngineRequest {
	s.OwnerAccount = &v
	return s
}

func (s *OpenComputeEngineRequest) SetOwnerId(v int64) *OpenComputeEngineRequest {
	s.OwnerId = &v
	return s
}

func (s *OpenComputeEngineRequest) SetResourceOwnerAccount(v string) *OpenComputeEngineRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *OpenComputeEngineRequest) SetResourceOwnerId(v int64) *OpenComputeEngineRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *OpenComputeEngineRequest) SetSecurityToken(v string) *OpenComputeEngineRequest {
	s.SecurityToken = &v
	return s
}

type OpenComputeEngineResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	RequestId          *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s OpenComputeEngineResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OpenComputeEngineResponseBody) GoString() string {
	return s.String()
}

func (s *OpenComputeEngineResponseBody) SetAccessDeniedDetail(v string) *OpenComputeEngineResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *OpenComputeEngineResponseBody) SetRequestId(v string) *OpenComputeEngineResponseBody {
	s.RequestId = &v
	return s
}

type OpenComputeEngineResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OpenComputeEngineResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OpenComputeEngineResponse) String() string {
	return tea.Prettify(s)
}

func (s OpenComputeEngineResponse) GoString() string {
	return s.String()
}

func (s *OpenComputeEngineResponse) SetHeaders(v map[string]*string) *OpenComputeEngineResponse {
	s.Headers = v
	return s
}

func (s *OpenComputeEngineResponse) SetStatusCode(v int32) *OpenComputeEngineResponse {
	s.StatusCode = &v
	return s
}

func (s *OpenComputeEngineResponse) SetBody(v *OpenComputeEngineResponseBody) *OpenComputeEngineResponse {
	s.Body = v
	return s
}

type OpenComputePreCheckRequest struct {
	CpuLimit *string `json:"CpuLimit,omitempty" xml:"CpuLimit,omitempty"`
	// This parameter is required.
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	MemoryLimit          *string `json:"MemoryLimit,omitempty" xml:"MemoryLimit,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s OpenComputePreCheckRequest) String() string {
	return tea.Prettify(s)
}

func (s OpenComputePreCheckRequest) GoString() string {
	return s.String()
}

func (s *OpenComputePreCheckRequest) SetCpuLimit(v string) *OpenComputePreCheckRequest {
	s.CpuLimit = &v
	return s
}

func (s *OpenComputePreCheckRequest) SetInstanceId(v string) *OpenComputePreCheckRequest {
	s.InstanceId = &v
	return s
}

func (s *OpenComputePreCheckRequest) SetMemoryLimit(v string) *OpenComputePreCheckRequest {
	s.MemoryLimit = &v
	return s
}

func (s *OpenComputePreCheckRequest) SetOwnerAccount(v string) *OpenComputePreCheckRequest {
	s.OwnerAccount = &v
	return s
}

func (s *OpenComputePreCheckRequest) SetOwnerId(v int64) *OpenComputePreCheckRequest {
	s.OwnerId = &v
	return s
}

func (s *OpenComputePreCheckRequest) SetResourceOwnerAccount(v string) *OpenComputePreCheckRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *OpenComputePreCheckRequest) SetResourceOwnerId(v int64) *OpenComputePreCheckRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *OpenComputePreCheckRequest) SetSecurityToken(v string) *OpenComputePreCheckRequest {
	s.SecurityToken = &v
	return s
}

type OpenComputePreCheckResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	RequestId          *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s OpenComputePreCheckResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OpenComputePreCheckResponseBody) GoString() string {
	return s.String()
}

func (s *OpenComputePreCheckResponseBody) SetAccessDeniedDetail(v string) *OpenComputePreCheckResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *OpenComputePreCheckResponseBody) SetRequestId(v string) *OpenComputePreCheckResponseBody {
	s.RequestId = &v
	return s
}

type OpenComputePreCheckResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OpenComputePreCheckResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OpenComputePreCheckResponse) String() string {
	return tea.Prettify(s)
}

func (s OpenComputePreCheckResponse) GoString() string {
	return s.String()
}

func (s *OpenComputePreCheckResponse) SetHeaders(v map[string]*string) *OpenComputePreCheckResponse {
	s.Headers = v
	return s
}

func (s *OpenComputePreCheckResponse) SetStatusCode(v int32) *OpenComputePreCheckResponse {
	s.StatusCode = &v
	return s
}

func (s *OpenComputePreCheckResponse) SetBody(v *OpenComputePreCheckResponseBody) *OpenComputePreCheckResponse {
	s.Body = v
	return s
}

type ReleaseLindormInstanceRequest struct {
	// Specifies whether to release the instance immediately. If you set this parameter to false, data in the released instance is retained for seven days before it is completely deleted. If you set this parameter to true, data in the released instance is immediately deleted. The default value is false.
	//
	// example:
	//
	// false
	Immediately *bool `json:"Immediately,omitempty" xml:"Immediately,omitempty"`
	// Instance ID, which can be obtained by calling the [GetLindormInstanceList](https://help.aliyun.com/document_detail/426069.html) interface.
	//
	// This parameter is required.
	//
	// example:
	//
	// ld-bp1z3506imz2f****
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s ReleaseLindormInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s ReleaseLindormInstanceRequest) GoString() string {
	return s.String()
}

func (s *ReleaseLindormInstanceRequest) SetImmediately(v bool) *ReleaseLindormInstanceRequest {
	s.Immediately = &v
	return s
}

func (s *ReleaseLindormInstanceRequest) SetInstanceId(v string) *ReleaseLindormInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *ReleaseLindormInstanceRequest) SetOwnerAccount(v string) *ReleaseLindormInstanceRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ReleaseLindormInstanceRequest) SetOwnerId(v int64) *ReleaseLindormInstanceRequest {
	s.OwnerId = &v
	return s
}

func (s *ReleaseLindormInstanceRequest) SetResourceOwnerAccount(v string) *ReleaseLindormInstanceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ReleaseLindormInstanceRequest) SetResourceOwnerId(v int64) *ReleaseLindormInstanceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ReleaseLindormInstanceRequest) SetSecurityToken(v string) *ReleaseLindormInstanceRequest {
	s.SecurityToken = &v
	return s
}

type ReleaseLindormInstanceResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// F67BFFF3-F5C2-45B5-9C28-6E4A1E51****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ReleaseLindormInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ReleaseLindormInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *ReleaseLindormInstanceResponseBody) SetRequestId(v string) *ReleaseLindormInstanceResponseBody {
	s.RequestId = &v
	return s
}

type ReleaseLindormInstanceResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReleaseLindormInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReleaseLindormInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s ReleaseLindormInstanceResponse) GoString() string {
	return s.String()
}

func (s *ReleaseLindormInstanceResponse) SetHeaders(v map[string]*string) *ReleaseLindormInstanceResponse {
	s.Headers = v
	return s
}

func (s *ReleaseLindormInstanceResponse) SetStatusCode(v int32) *ReleaseLindormInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *ReleaseLindormInstanceResponse) SetBody(v *ReleaseLindormInstanceResponseBody) *ReleaseLindormInstanceResponse {
	s.Body = v
	return s
}

type ReleaseLindormV2InstanceRequest struct {
	Immediately *bool `json:"Immediately,omitempty" xml:"Immediately,omitempty"`
	// This parameter is required.
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s ReleaseLindormV2InstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s ReleaseLindormV2InstanceRequest) GoString() string {
	return s.String()
}

func (s *ReleaseLindormV2InstanceRequest) SetImmediately(v bool) *ReleaseLindormV2InstanceRequest {
	s.Immediately = &v
	return s
}

func (s *ReleaseLindormV2InstanceRequest) SetInstanceId(v string) *ReleaseLindormV2InstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *ReleaseLindormV2InstanceRequest) SetOwnerAccount(v string) *ReleaseLindormV2InstanceRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ReleaseLindormV2InstanceRequest) SetOwnerId(v int64) *ReleaseLindormV2InstanceRequest {
	s.OwnerId = &v
	return s
}

func (s *ReleaseLindormV2InstanceRequest) SetResourceOwnerAccount(v string) *ReleaseLindormV2InstanceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ReleaseLindormV2InstanceRequest) SetResourceOwnerId(v int64) *ReleaseLindormV2InstanceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ReleaseLindormV2InstanceRequest) SetSecurityToken(v string) *ReleaseLindormV2InstanceRequest {
	s.SecurityToken = &v
	return s
}

type ReleaseLindormV2InstanceResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	RequestId          *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ReleaseLindormV2InstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ReleaseLindormV2InstanceResponseBody) GoString() string {
	return s.String()
}

func (s *ReleaseLindormV2InstanceResponseBody) SetAccessDeniedDetail(v string) *ReleaseLindormV2InstanceResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ReleaseLindormV2InstanceResponseBody) SetRequestId(v string) *ReleaseLindormV2InstanceResponseBody {
	s.RequestId = &v
	return s
}

type ReleaseLindormV2InstanceResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReleaseLindormV2InstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReleaseLindormV2InstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s ReleaseLindormV2InstanceResponse) GoString() string {
	return s.String()
}

func (s *ReleaseLindormV2InstanceResponse) SetHeaders(v map[string]*string) *ReleaseLindormV2InstanceResponse {
	s.Headers = v
	return s
}

func (s *ReleaseLindormV2InstanceResponse) SetStatusCode(v int32) *ReleaseLindormV2InstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *ReleaseLindormV2InstanceResponse) SetBody(v *ReleaseLindormV2InstanceResponseBody) *ReleaseLindormV2InstanceResponse {
	s.Body = v
	return s
}

type RenewLindormInstanceRequest struct {
	// The subscription duration of the instance. The valid values of this parameter depend on the value of the PricingCycle parameter.
	//
	// 	- If PricingCycle is set to **Month**, set this parameter to an integer that ranges from **1*	- to **9**.
	//
	// 	- If PricingCycle is set to **Year**, set this parameter to an integer that ranges from **1*	- to **3**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	Duration *int32 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The ID of the instance that you want to renew. You can call the [GetLindormInstanceList](https://help.aliyun.com/document_detail/426069.html) operation to obtain the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ld-bp1z3506imz2f****
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The period based on which you are charged for the instance. Valid values:
	//
	// 	- **Month**: You are charged for the instance based on months.
	//
	// 	- **Year**: You are charged for the instance based on years.
	//
	// This parameter is required.
	//
	// example:
	//
	// Month
	PricingCycle *string `json:"PricingCycle,omitempty" xml:"PricingCycle,omitempty"`
	// The ID of the region in which the instance that you want to renew is located. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/426062.html) operation to query the region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s RenewLindormInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s RenewLindormInstanceRequest) GoString() string {
	return s.String()
}

func (s *RenewLindormInstanceRequest) SetDuration(v int32) *RenewLindormInstanceRequest {
	s.Duration = &v
	return s
}

func (s *RenewLindormInstanceRequest) SetInstanceId(v string) *RenewLindormInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *RenewLindormInstanceRequest) SetOwnerAccount(v string) *RenewLindormInstanceRequest {
	s.OwnerAccount = &v
	return s
}

func (s *RenewLindormInstanceRequest) SetOwnerId(v int64) *RenewLindormInstanceRequest {
	s.OwnerId = &v
	return s
}

func (s *RenewLindormInstanceRequest) SetPricingCycle(v string) *RenewLindormInstanceRequest {
	s.PricingCycle = &v
	return s
}

func (s *RenewLindormInstanceRequest) SetRegionId(v string) *RenewLindormInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *RenewLindormInstanceRequest) SetResourceOwnerAccount(v string) *RenewLindormInstanceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *RenewLindormInstanceRequest) SetResourceOwnerId(v int64) *RenewLindormInstanceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *RenewLindormInstanceRequest) SetSecurityToken(v string) *RenewLindormInstanceRequest {
	s.SecurityToken = &v
	return s
}

type RenewLindormInstanceResponseBody struct {
	// The details about the access denial.
	//
	// example:
	//
	// {"AuthAction":"xxx","AuthPrincipalDisplayName":"222","AuthPrincipalOwnerId":"111","AuthPrincipalType":"SubUser",,"NoPermissionType":"ImplicitDeny","PolicyType":"AccountLevelIdentityBasedPolicy","EncodedDiagnosticMessage":"xxxxxx"}
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// ld-bp1z3506imz2f****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the order. You can obtain the order ID on the Orders page of the Expenses and Costs console.
	//
	// example:
	//
	// 213465921640411
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 1556DCB0-043A-4444-8BD9-CF4A68E7EE64
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RenewLindormInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RenewLindormInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *RenewLindormInstanceResponseBody) SetAccessDeniedDetail(v string) *RenewLindormInstanceResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *RenewLindormInstanceResponseBody) SetInstanceId(v string) *RenewLindormInstanceResponseBody {
	s.InstanceId = &v
	return s
}

func (s *RenewLindormInstanceResponseBody) SetOrderId(v int64) *RenewLindormInstanceResponseBody {
	s.OrderId = &v
	return s
}

func (s *RenewLindormInstanceResponseBody) SetRequestId(v string) *RenewLindormInstanceResponseBody {
	s.RequestId = &v
	return s
}

type RenewLindormInstanceResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RenewLindormInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RenewLindormInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s RenewLindormInstanceResponse) GoString() string {
	return s.String()
}

func (s *RenewLindormInstanceResponse) SetHeaders(v map[string]*string) *RenewLindormInstanceResponse {
	s.Headers = v
	return s
}

func (s *RenewLindormInstanceResponse) SetStatusCode(v int32) *RenewLindormInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *RenewLindormInstanceResponse) SetBody(v *RenewLindormInstanceResponseBody) *RenewLindormInstanceResponse {
	s.Body = v
	return s
}

type RestartLdpsComputeGroupRequest struct {
	// This parameter is required.
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// This parameter is required.
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s RestartLdpsComputeGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s RestartLdpsComputeGroupRequest) GoString() string {
	return s.String()
}

func (s *RestartLdpsComputeGroupRequest) SetGroupName(v string) *RestartLdpsComputeGroupRequest {
	s.GroupName = &v
	return s
}

func (s *RestartLdpsComputeGroupRequest) SetInstanceId(v string) *RestartLdpsComputeGroupRequest {
	s.InstanceId = &v
	return s
}

func (s *RestartLdpsComputeGroupRequest) SetOwnerAccount(v string) *RestartLdpsComputeGroupRequest {
	s.OwnerAccount = &v
	return s
}

func (s *RestartLdpsComputeGroupRequest) SetOwnerId(v int64) *RestartLdpsComputeGroupRequest {
	s.OwnerId = &v
	return s
}

func (s *RestartLdpsComputeGroupRequest) SetRegionId(v string) *RestartLdpsComputeGroupRequest {
	s.RegionId = &v
	return s
}

func (s *RestartLdpsComputeGroupRequest) SetResourceOwnerAccount(v string) *RestartLdpsComputeGroupRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *RestartLdpsComputeGroupRequest) SetResourceOwnerId(v int64) *RestartLdpsComputeGroupRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *RestartLdpsComputeGroupRequest) SetSecurityToken(v string) *RestartLdpsComputeGroupRequest {
	s.SecurityToken = &v
	return s
}

type RestartLdpsComputeGroupResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	RequestId          *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RestartLdpsComputeGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RestartLdpsComputeGroupResponseBody) GoString() string {
	return s.String()
}

func (s *RestartLdpsComputeGroupResponseBody) SetAccessDeniedDetail(v string) *RestartLdpsComputeGroupResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *RestartLdpsComputeGroupResponseBody) SetRequestId(v string) *RestartLdpsComputeGroupResponseBody {
	s.RequestId = &v
	return s
}

type RestartLdpsComputeGroupResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RestartLdpsComputeGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RestartLdpsComputeGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s RestartLdpsComputeGroupResponse) GoString() string {
	return s.String()
}

func (s *RestartLdpsComputeGroupResponse) SetHeaders(v map[string]*string) *RestartLdpsComputeGroupResponse {
	s.Headers = v
	return s
}

func (s *RestartLdpsComputeGroupResponse) SetStatusCode(v int32) *RestartLdpsComputeGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *RestartLdpsComputeGroupResponse) SetBody(v *RestartLdpsComputeGroupResponseBody) *RestartLdpsComputeGroupResponse {
	s.Body = v
	return s
}

type SetDefaultOlapComputeGroupRequest struct {
	// This parameter is required.
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// This parameter is required.
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	IsDefault            *bool   `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s SetDefaultOlapComputeGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s SetDefaultOlapComputeGroupRequest) GoString() string {
	return s.String()
}

func (s *SetDefaultOlapComputeGroupRequest) SetGroupName(v string) *SetDefaultOlapComputeGroupRequest {
	s.GroupName = &v
	return s
}

func (s *SetDefaultOlapComputeGroupRequest) SetInstanceId(v string) *SetDefaultOlapComputeGroupRequest {
	s.InstanceId = &v
	return s
}

func (s *SetDefaultOlapComputeGroupRequest) SetIsDefault(v bool) *SetDefaultOlapComputeGroupRequest {
	s.IsDefault = &v
	return s
}

func (s *SetDefaultOlapComputeGroupRequest) SetOwnerAccount(v string) *SetDefaultOlapComputeGroupRequest {
	s.OwnerAccount = &v
	return s
}

func (s *SetDefaultOlapComputeGroupRequest) SetOwnerId(v int64) *SetDefaultOlapComputeGroupRequest {
	s.OwnerId = &v
	return s
}

func (s *SetDefaultOlapComputeGroupRequest) SetRegionId(v string) *SetDefaultOlapComputeGroupRequest {
	s.RegionId = &v
	return s
}

func (s *SetDefaultOlapComputeGroupRequest) SetResourceOwnerAccount(v string) *SetDefaultOlapComputeGroupRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SetDefaultOlapComputeGroupRequest) SetResourceOwnerId(v int64) *SetDefaultOlapComputeGroupRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *SetDefaultOlapComputeGroupRequest) SetSecurityToken(v string) *SetDefaultOlapComputeGroupRequest {
	s.SecurityToken = &v
	return s
}

type SetDefaultOlapComputeGroupResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetDefaultOlapComputeGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetDefaultOlapComputeGroupResponseBody) GoString() string {
	return s.String()
}

func (s *SetDefaultOlapComputeGroupResponseBody) SetRequestId(v string) *SetDefaultOlapComputeGroupResponseBody {
	s.RequestId = &v
	return s
}

type SetDefaultOlapComputeGroupResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetDefaultOlapComputeGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetDefaultOlapComputeGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s SetDefaultOlapComputeGroupResponse) GoString() string {
	return s.String()
}

func (s *SetDefaultOlapComputeGroupResponse) SetHeaders(v map[string]*string) *SetDefaultOlapComputeGroupResponse {
	s.Headers = v
	return s
}

func (s *SetDefaultOlapComputeGroupResponse) SetStatusCode(v int32) *SetDefaultOlapComputeGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *SetDefaultOlapComputeGroupResponse) SetBody(v *SetDefaultOlapComputeGroupResponseBody) *SetDefaultOlapComputeGroupResponse {
	s.Body = v
	return s
}

type SwitchLSQLV3MySQLServiceRequest struct {
	// The type of the operation. Valid value:
	//
	// 	- 1: enables the MySQL compatibility feature.
	//
	// 	- 0: disables the MySQL compatibility feature.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	ActionType *int32 `json:"ActionType,omitempty" xml:"ActionType,omitempty"`
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ld-bp1o3y0yme2i2****
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s SwitchLSQLV3MySQLServiceRequest) String() string {
	return tea.Prettify(s)
}

func (s SwitchLSQLV3MySQLServiceRequest) GoString() string {
	return s.String()
}

func (s *SwitchLSQLV3MySQLServiceRequest) SetActionType(v int32) *SwitchLSQLV3MySQLServiceRequest {
	s.ActionType = &v
	return s
}

func (s *SwitchLSQLV3MySQLServiceRequest) SetInstanceId(v string) *SwitchLSQLV3MySQLServiceRequest {
	s.InstanceId = &v
	return s
}

func (s *SwitchLSQLV3MySQLServiceRequest) SetOwnerAccount(v string) *SwitchLSQLV3MySQLServiceRequest {
	s.OwnerAccount = &v
	return s
}

func (s *SwitchLSQLV3MySQLServiceRequest) SetOwnerId(v int64) *SwitchLSQLV3MySQLServiceRequest {
	s.OwnerId = &v
	return s
}

func (s *SwitchLSQLV3MySQLServiceRequest) SetResourceOwnerAccount(v string) *SwitchLSQLV3MySQLServiceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SwitchLSQLV3MySQLServiceRequest) SetResourceOwnerId(v int64) *SwitchLSQLV3MySQLServiceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *SwitchLSQLV3MySQLServiceRequest) SetSecurityToken(v string) *SwitchLSQLV3MySQLServiceRequest {
	s.SecurityToken = &v
	return s
}

type SwitchLSQLV3MySQLServiceResponseBody struct {
	// The details about the access denial.
	//
	// example:
	//
	// {"AuthAction":"xxx","AuthPrincipalDisplayName":"222","AuthPrincipalOwnerId":"111","AuthPrincipalType":"SubUser",,"NoPermissionType":"ImplicitDeny","PolicyType":"AccountLevelIdentityBasedPolicy","EncodedDiagnosticMessage":"xxxxxx"}
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 1556DCB0-043A-4444-8BD9-CF4A68E7EE64
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SwitchLSQLV3MySQLServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SwitchLSQLV3MySQLServiceResponseBody) GoString() string {
	return s.String()
}

func (s *SwitchLSQLV3MySQLServiceResponseBody) SetAccessDeniedDetail(v string) *SwitchLSQLV3MySQLServiceResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *SwitchLSQLV3MySQLServiceResponseBody) SetRequestId(v string) *SwitchLSQLV3MySQLServiceResponseBody {
	s.RequestId = &v
	return s
}

type SwitchLSQLV3MySQLServiceResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SwitchLSQLV3MySQLServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SwitchLSQLV3MySQLServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s SwitchLSQLV3MySQLServiceResponse) GoString() string {
	return s.String()
}

func (s *SwitchLSQLV3MySQLServiceResponse) SetHeaders(v map[string]*string) *SwitchLSQLV3MySQLServiceResponse {
	s.Headers = v
	return s
}

func (s *SwitchLSQLV3MySQLServiceResponse) SetStatusCode(v int32) *SwitchLSQLV3MySQLServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *SwitchLSQLV3MySQLServiceResponse) SetBody(v *SwitchLSQLV3MySQLServiceResponseBody) *SwitchLSQLV3MySQLServiceResponse {
	s.Body = v
	return s
}

type TagResourcesRequest struct {
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region in which the instances you want to associate tags with are located. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/426062.html) operation to query the region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The list of resource IDs.
	//
	// This parameter is required.
	ResourceId           []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	ResourceOwnerAccount *string   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64    `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The resource type. Set the value to **INSTANCE**.
	//
	// This parameter is required.
	//
	// example:
	//
	// INSTANCE
	ResourceType  *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The tags that you want to associate with the resource.
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

func (s *TagResourcesRequest) SetOwnerAccount(v string) *TagResourcesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *TagResourcesRequest) SetOwnerId(v int64) *TagResourcesRequest {
	s.OwnerId = &v
	return s
}

func (s *TagResourcesRequest) SetRegionId(v string) *TagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *TagResourcesRequest) SetResourceId(v []*string) *TagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *TagResourcesRequest) SetResourceOwnerAccount(v string) *TagResourcesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *TagResourcesRequest) SetResourceOwnerId(v int64) *TagResourcesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *TagResourcesRequest) SetResourceType(v string) *TagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *TagResourcesRequest) SetSecurityToken(v string) *TagResourcesRequest {
	s.SecurityToken = &v
	return s
}

func (s *TagResourcesRequest) SetTag(v []*TagResourcesRequestTag) *TagResourcesRequest {
	s.Tag = v
	return s
}

type TagResourcesRequestTag struct {
	// The key of the tag that you want to associate with the resource.
	//
	// > You can specify the keys of multiple tags. For example, you can specify the key of the first tag in the first key-value pair contained in the value of this parameter and specify the key of the second tag in the second key-value pair.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag that you want to associate with the resource.
	//
	// > You can specify the values of multiple tags. For example, you can specify the value of the first tag in the first key-value pair contained in the value of this parameter and specify the value of the second tag in the second key-value pair.
	//
	// example:
	//
	// 2.2.8
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
	// The ID of the request.
	//
	// example:
	//
	// 4F23D50C-400C-592C-9486-9D1E10179065
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
	// Specifies whether to remove all tags from the instances. Valid values:
	//
	// 	- **true**: Remove all tags from the instances.
	//
	// 	- **false**: Do not remove all tags from the instances.
	//
	// >  The default value of this parameter is false.
	//
	//
	//
	// 	- If you specify this parameter together with the TagKey parameter, this parameter does not take effect.
	//
	// example:
	//
	// false
	All          *bool   `json:"All,omitempty" xml:"All,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The IDs of instances.
	//
	// This parameter is required.
	ResourceId           []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	ResourceOwnerAccount *string   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64    `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The resource type. Set the value to **INSTANCE**.
	//
	// This parameter is required.
	//
	// example:
	//
	// INSTANCE
	ResourceType  *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The list of keys of the tags that you want to remove.
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

func (s *UntagResourcesRequest) SetOwnerAccount(v string) *UntagResourcesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *UntagResourcesRequest) SetOwnerId(v int64) *UntagResourcesRequest {
	s.OwnerId = &v
	return s
}

func (s *UntagResourcesRequest) SetResourceId(v []*string) *UntagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *UntagResourcesRequest) SetResourceOwnerAccount(v string) *UntagResourcesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UntagResourcesRequest) SetResourceOwnerId(v int64) *UntagResourcesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UntagResourcesRequest) SetResourceType(v string) *UntagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *UntagResourcesRequest) SetSecurityToken(v string) *UntagResourcesRequest {
	s.SecurityToken = &v
	return s
}

func (s *UntagResourcesRequest) SetTagKey(v []*string) *UntagResourcesRequest {
	s.TagKey = v
	return s
}

type UntagResourcesResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 8CACBBCE-7519-545C-8695-86D4F09CED7E
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

type UpdateInstanceIpWhiteListRequest struct {
	// Specifies whether to clear all IP addresses and CIDR blocks in the whitelist.
	//
	// example:
	//
	// false
	Delete *bool `json:"Delete,omitempty" xml:"Delete,omitempty"`
	// The name of the IP whitelist. Default value: user.
	//
	// example:
	//
	// test_group
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The ID of the instance for which you want to configure a whitelist. You can call the [GetLindormInstanceList](https://help.aliyun.com/document_detail/426069.html) operation to obtain the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ld-bp1z3506imz2f****
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The IP addresses or CIDR blocks that you want to add to the whitelist.
	//
	// >  If you add 127.0.0.1 to the whitelist, all IP addresses cannot be used to access the Lindorm instance. If you add the CIDR block 192.168.0.0/24 to the whitelist, you can use all IP addresses in the CIDR block to access the Lindorm instance. Separate multiple IP addresses or CIDR blocks with commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// 106.11.XX.XX/24
	SecurityIpList *string `json:"SecurityIpList,omitempty" xml:"SecurityIpList,omitempty"`
	SecurityToken  *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s UpdateInstanceIpWhiteListRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateInstanceIpWhiteListRequest) GoString() string {
	return s.String()
}

func (s *UpdateInstanceIpWhiteListRequest) SetDelete(v bool) *UpdateInstanceIpWhiteListRequest {
	s.Delete = &v
	return s
}

func (s *UpdateInstanceIpWhiteListRequest) SetGroupName(v string) *UpdateInstanceIpWhiteListRequest {
	s.GroupName = &v
	return s
}

func (s *UpdateInstanceIpWhiteListRequest) SetInstanceId(v string) *UpdateInstanceIpWhiteListRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateInstanceIpWhiteListRequest) SetOwnerAccount(v string) *UpdateInstanceIpWhiteListRequest {
	s.OwnerAccount = &v
	return s
}

func (s *UpdateInstanceIpWhiteListRequest) SetOwnerId(v int64) *UpdateInstanceIpWhiteListRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateInstanceIpWhiteListRequest) SetResourceOwnerAccount(v string) *UpdateInstanceIpWhiteListRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UpdateInstanceIpWhiteListRequest) SetResourceOwnerId(v int64) *UpdateInstanceIpWhiteListRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UpdateInstanceIpWhiteListRequest) SetSecurityIpList(v string) *UpdateInstanceIpWhiteListRequest {
	s.SecurityIpList = &v
	return s
}

func (s *UpdateInstanceIpWhiteListRequest) SetSecurityToken(v string) *UpdateInstanceIpWhiteListRequest {
	s.SecurityToken = &v
	return s
}

type UpdateInstanceIpWhiteListResponseBody struct {
	// The details about the access denial.
	//
	// example:
	//
	// {"AuthAction":"xxx","AuthPrincipalDisplayName":"222","AuthPrincipalOwnerId":"111","AuthPrincipalType":"SubUser",,"NoPermissionType":"ImplicitDeny","PolicyType":"AccountLevelIdentityBasedPolicy","EncodedDiagnosticMessage":"xxxxxx"}
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 4944539D-D27C-458D-95F1-2DCEB5E0EED5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateInstanceIpWhiteListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateInstanceIpWhiteListResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateInstanceIpWhiteListResponseBody) SetAccessDeniedDetail(v string) *UpdateInstanceIpWhiteListResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *UpdateInstanceIpWhiteListResponseBody) SetRequestId(v string) *UpdateInstanceIpWhiteListResponseBody {
	s.RequestId = &v
	return s
}

type UpdateInstanceIpWhiteListResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateInstanceIpWhiteListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateInstanceIpWhiteListResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateInstanceIpWhiteListResponse) GoString() string {
	return s.String()
}

func (s *UpdateInstanceIpWhiteListResponse) SetHeaders(v map[string]*string) *UpdateInstanceIpWhiteListResponse {
	s.Headers = v
	return s
}

func (s *UpdateInstanceIpWhiteListResponse) SetStatusCode(v int32) *UpdateInstanceIpWhiteListResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateInstanceIpWhiteListResponse) SetBody(v *UpdateInstanceIpWhiteListResponseBody) *UpdateInstanceIpWhiteListResponse {
	s.Body = v
	return s
}

type UpdateInstanceSecurityGroupsRequest struct {
	// This parameter is required.
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// This parameter is required.
	SecurityGroups *string `json:"SecurityGroups,omitempty" xml:"SecurityGroups,omitempty"`
	SecurityToken  *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s UpdateInstanceSecurityGroupsRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateInstanceSecurityGroupsRequest) GoString() string {
	return s.String()
}

func (s *UpdateInstanceSecurityGroupsRequest) SetInstanceId(v string) *UpdateInstanceSecurityGroupsRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateInstanceSecurityGroupsRequest) SetOwnerAccount(v string) *UpdateInstanceSecurityGroupsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *UpdateInstanceSecurityGroupsRequest) SetOwnerId(v int64) *UpdateInstanceSecurityGroupsRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateInstanceSecurityGroupsRequest) SetResourceOwnerAccount(v string) *UpdateInstanceSecurityGroupsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UpdateInstanceSecurityGroupsRequest) SetResourceOwnerId(v int64) *UpdateInstanceSecurityGroupsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UpdateInstanceSecurityGroupsRequest) SetSecurityGroups(v string) *UpdateInstanceSecurityGroupsRequest {
	s.SecurityGroups = &v
	return s
}

func (s *UpdateInstanceSecurityGroupsRequest) SetSecurityToken(v string) *UpdateInstanceSecurityGroupsRequest {
	s.SecurityToken = &v
	return s
}

type UpdateInstanceSecurityGroupsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateInstanceSecurityGroupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateInstanceSecurityGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateInstanceSecurityGroupsResponseBody) SetRequestId(v string) *UpdateInstanceSecurityGroupsResponseBody {
	s.RequestId = &v
	return s
}

type UpdateInstanceSecurityGroupsResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateInstanceSecurityGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateInstanceSecurityGroupsResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateInstanceSecurityGroupsResponse) GoString() string {
	return s.String()
}

func (s *UpdateInstanceSecurityGroupsResponse) SetHeaders(v map[string]*string) *UpdateInstanceSecurityGroupsResponse {
	s.Headers = v
	return s
}

func (s *UpdateInstanceSecurityGroupsResponse) SetStatusCode(v int32) *UpdateInstanceSecurityGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateInstanceSecurityGroupsResponse) SetBody(v *UpdateInstanceSecurityGroupsResponseBody) *UpdateInstanceSecurityGroupsResponse {
	s.Body = v
	return s
}

type UpdateLdpsComputeGroupRequest struct {
	// This parameter is required.
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// This parameter is required.
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	Properties           *string `json:"Properties,omitempty" xml:"Properties,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s UpdateLdpsComputeGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateLdpsComputeGroupRequest) GoString() string {
	return s.String()
}

func (s *UpdateLdpsComputeGroupRequest) SetGroupName(v string) *UpdateLdpsComputeGroupRequest {
	s.GroupName = &v
	return s
}

func (s *UpdateLdpsComputeGroupRequest) SetInstanceId(v string) *UpdateLdpsComputeGroupRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateLdpsComputeGroupRequest) SetOwnerAccount(v string) *UpdateLdpsComputeGroupRequest {
	s.OwnerAccount = &v
	return s
}

func (s *UpdateLdpsComputeGroupRequest) SetOwnerId(v int64) *UpdateLdpsComputeGroupRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateLdpsComputeGroupRequest) SetProperties(v string) *UpdateLdpsComputeGroupRequest {
	s.Properties = &v
	return s
}

func (s *UpdateLdpsComputeGroupRequest) SetRegionId(v string) *UpdateLdpsComputeGroupRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateLdpsComputeGroupRequest) SetResourceOwnerAccount(v string) *UpdateLdpsComputeGroupRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UpdateLdpsComputeGroupRequest) SetResourceOwnerId(v int64) *UpdateLdpsComputeGroupRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UpdateLdpsComputeGroupRequest) SetSecurityToken(v string) *UpdateLdpsComputeGroupRequest {
	s.SecurityToken = &v
	return s
}

type UpdateLdpsComputeGroupResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	RequestId          *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateLdpsComputeGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateLdpsComputeGroupResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateLdpsComputeGroupResponseBody) SetAccessDeniedDetail(v string) *UpdateLdpsComputeGroupResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *UpdateLdpsComputeGroupResponseBody) SetRequestId(v string) *UpdateLdpsComputeGroupResponseBody {
	s.RequestId = &v
	return s
}

type UpdateLdpsComputeGroupResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateLdpsComputeGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateLdpsComputeGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateLdpsComputeGroupResponse) GoString() string {
	return s.String()
}

func (s *UpdateLdpsComputeGroupResponse) SetHeaders(v map[string]*string) *UpdateLdpsComputeGroupResponse {
	s.Headers = v
	return s
}

func (s *UpdateLdpsComputeGroupResponse) SetStatusCode(v int32) *UpdateLdpsComputeGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateLdpsComputeGroupResponse) SetBody(v *UpdateLdpsComputeGroupResponseBody) *UpdateLdpsComputeGroupResponse {
	s.Body = v
	return s
}

type UpdateLindormV2InstanceParameterRequest struct {
	// This parameter is required.
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	ParameterKey *string `json:"ParameterKey,omitempty" xml:"ParameterKey,omitempty"`
	// This parameter is required.
	ParameterValue       *string `json:"ParameterValue,omitempty" xml:"ParameterValue,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	UpdateType           *string `json:"UpdateType,omitempty" xml:"UpdateType,omitempty"`
}

func (s UpdateLindormV2InstanceParameterRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateLindormV2InstanceParameterRequest) GoString() string {
	return s.String()
}

func (s *UpdateLindormV2InstanceParameterRequest) SetInstanceId(v string) *UpdateLindormV2InstanceParameterRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateLindormV2InstanceParameterRequest) SetOwnerAccount(v string) *UpdateLindormV2InstanceParameterRequest {
	s.OwnerAccount = &v
	return s
}

func (s *UpdateLindormV2InstanceParameterRequest) SetOwnerId(v int64) *UpdateLindormV2InstanceParameterRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateLindormV2InstanceParameterRequest) SetParameterKey(v string) *UpdateLindormV2InstanceParameterRequest {
	s.ParameterKey = &v
	return s
}

func (s *UpdateLindormV2InstanceParameterRequest) SetParameterValue(v string) *UpdateLindormV2InstanceParameterRequest {
	s.ParameterValue = &v
	return s
}

func (s *UpdateLindormV2InstanceParameterRequest) SetResourceOwnerAccount(v string) *UpdateLindormV2InstanceParameterRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UpdateLindormV2InstanceParameterRequest) SetResourceOwnerId(v int64) *UpdateLindormV2InstanceParameterRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UpdateLindormV2InstanceParameterRequest) SetSecurityToken(v string) *UpdateLindormV2InstanceParameterRequest {
	s.SecurityToken = &v
	return s
}

func (s *UpdateLindormV2InstanceParameterRequest) SetUpdateType(v string) *UpdateLindormV2InstanceParameterRequest {
	s.UpdateType = &v
	return s
}

type UpdateLindormV2InstanceParameterResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	RequestId          *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateLindormV2InstanceParameterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateLindormV2InstanceParameterResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateLindormV2InstanceParameterResponseBody) SetAccessDeniedDetail(v string) *UpdateLindormV2InstanceParameterResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *UpdateLindormV2InstanceParameterResponseBody) SetRequestId(v string) *UpdateLindormV2InstanceParameterResponseBody {
	s.RequestId = &v
	return s
}

type UpdateLindormV2InstanceParameterResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateLindormV2InstanceParameterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateLindormV2InstanceParameterResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateLindormV2InstanceParameterResponse) GoString() string {
	return s.String()
}

func (s *UpdateLindormV2InstanceParameterResponse) SetHeaders(v map[string]*string) *UpdateLindormV2InstanceParameterResponse {
	s.Headers = v
	return s
}

func (s *UpdateLindormV2InstanceParameterResponse) SetStatusCode(v int32) *UpdateLindormV2InstanceParameterResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateLindormV2InstanceParameterResponse) SetBody(v *UpdateLindormV2InstanceParameterResponseBody) *UpdateLindormV2InstanceParameterResponse {
	s.Body = v
	return s
}

type UpgradeLindormInstanceRequest struct {
	// The storage capacity of the instance after it is upgraded. Unit: GB. Valid values: **480*	- to **1017600**.
	//
	// example:
	//
	// 480
	ClusterStorage *int32 `json:"ClusterStorage,omitempty" xml:"ClusterStorage,omitempty"`
	// The cold storage capacity of the instance after it is upgraded. Unit: GB. Valid values: **800*	- to **1000000**.
	//
	// example:
	//
	// 800
	ColdStorage *int32 `json:"ColdStorage,omitempty" xml:"ColdStorage,omitempty"`
	// The storage capacity of a single core node in the instance after the instance upgraded. This parameter is available only if the instance you want to upgrade is a multi-zone instance. Unit: GB. Valid values: 400 to 64000. **This parameter is optional**.
	//
	// example:
	//
	// 400
	CoreSingleStorage *int32 `json:"CoreSingleStorage,omitempty" xml:"CoreSingleStorage,omitempty"`
	// The number of LindormDFS nodes in the instance after the instance is upgraded. Valid values: integers from **0*	- to **60**.
	//
	// example:
	//
	// 2
	FilestoreNum *int32 `json:"FilestoreNum,omitempty" xml:"FilestoreNum,omitempty"`
	// The specification of LindormDFS nodes in the instance after the instance is upgraded. Valid values:
	//
	// 	- **lindorm.g.xlarge**: Each node has 4 dedicated CPU cores and 16 GB of dedicated memory.
	//
	// 	- **lindorm.g.2xlarge**: Each node has 8 dedicated CPU cores and 32 GB of dedicated memory.
	//
	// 	- **lindorm.g.4xlarge**: Each node has 16 dedicated CPU cores and 64 GB of dedicated memory.
	//
	// 	- **lindorm.g.8xlarge**: Each node has 32 dedicated CPU cores and 128 GB of dedicated memory.
	//
	// example:
	//
	// lindorm.g.xlarge
	FilestoreSpec *string `json:"FilestoreSpec,omitempty" xml:"FilestoreSpec,omitempty"`
	// The ID of the instance that you want to upgrade, scale up, or enable cold storage. You can call the [GetLindormInstanceList](https://help.aliyun.com/document_detail/426069.html) operation to query the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ld-bp1o3y0yme2i2****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The number of LindormTable nodes in the instance after the instance is upgraded. Valid values: integers from **0*	- to **90**.
	//
	// > This parameter must be specified together with the LindormSpec parameter.
	//
	// example:
	//
	// 2
	LindormNum *int32 `json:"LindormNum,omitempty" xml:"LindormNum,omitempty"`
	// The specification of LindormTable nodes in the instance after the instance is upgraded. Valid values:
	//
	// 	- **lindorm.c.xlarge**: Each node has 4 dedicated CPU cores and 8 GB of dedicated memory.
	//
	// 	- **lindorm.c.2xlarge**: Each node has 8 dedicated CPU cores and 16 GB of dedicated memory.
	//
	// 	- **lindorm.c.4xlarge**: Each node has 16 dedicated CPU cores and 32 GB of dedicated memory.
	//
	// 	- **lindorm.c.8xlarge**: Each node has 32 dedicated CPU cores and 64 GB of dedicated memory.
	//
	// example:
	//
	// lindorm.c.xlarge
	LindormSpec *string `json:"LindormSpec,omitempty" xml:"LindormSpec,omitempty"`
	// The number of log nodes in the instance after the instance is upgraded. This parameter is available only if the instance you want to upgrade is a multi-zone instance. **This parameter is optional**.
	//
	// example:
	//
	// 4
	LogNum *int32 `json:"LogNum,omitempty" xml:"LogNum,omitempty"`
	// The storage capacity of a single log node in the instance after the instance upgraded. This parameter is available only if the instance you want to upgrade is a multi-zone instance. **This parameter is optional**.
	//
	// example:
	//
	// 400
	LogSingleStorage *int32 `json:"LogSingleStorage,omitempty" xml:"LogSingleStorage,omitempty"`
	// The specification of log nodes in the instance after the instance is upgraded. This parameter is available only if the instance you want to upgrade is a multi-zone instance. Valid values:
	//
	// 	- **lindorm.sn1.large**: Each node has 4 dedicated CPU cores and 8 GB of dedicated memory.
	//
	// 	- **lindorm.sn1.2xlarge**: Each node has 8 dedicated CPU cores and 16 GB of dedicated memory.
	//
	// **This parameter is optional**.
	//
	// example:
	//
	// lindorm.sn1.large
	LogSpec *string `json:"LogSpec,omitempty" xml:"LogSpec,omitempty"`
	// The number of LTS nodes in the instance after the instance is upgraded. Valid values: integers from **0*	- to **50**.
	//
	// example:
	//
	// 2
	LtsCoreNum *int32 `json:"LtsCoreNum,omitempty" xml:"LtsCoreNum,omitempty"`
	// The specification of Lindorm Tunnel Service (LTS) nodes in the instance after the instance is upgraded. Valid values:
	//
	// 	- **lindorm.g.xlarge**: Each node has 4 dedicated CPU cores and 16 GB of dedicated memory.
	//
	// 	- **lindorm.g.2xlarge**: Each node has 8 dedicated CPU cores and 32 GB of dedicated memory.
	//
	// example:
	//
	// lindorm.g.xlarge
	LtsCoreSpec  *string `json:"LtsCoreSpec,omitempty" xml:"LtsCoreSpec,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region in which the instance that you want to upgrade, scale up, or enable cold storage is located. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/426062.html) operation to query the region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The number of LindormSearch nodes in the instance after the instance is upgraded. Valid values: integers from **0*	- to **60**.
	//
	// example:
	//
	// 2
	SolrNum *int32 `json:"SolrNum,omitempty" xml:"SolrNum,omitempty"`
	// The specification of LindormSearch nodes in the instance after the instance is upgraded. Valid values:
	//
	// 	- **lindorm.g.xlarge**: Each node has 4 dedicated CPU cores and 16 GB of dedicated memory.
	//
	// 	- **lindorm.g.2xlarge**: Each node has 8 dedicated CPU cores and 32 GB of dedicated memory.
	//
	// 	- **lindorm.g.4xlarge**: Each node has 16 dedicated CPU cores and 64 GB of dedicated memory.
	//
	// 	- **lindorm.g.8xlarge**: Each node has 32 dedicated CPU cores and 128 GB of dedicated memory.
	//
	// example:
	//
	// lindorm.g.xlarge
	SolrSpec *string `json:"SolrSpec,omitempty" xml:"SolrSpec,omitempty"`
	// The number of LindormStream nodes in the instance after the instance is upgraded. Valid values: integers from **0*	- to **60**.
	//
	// example:
	//
	// 2
	StreamNum *int32 `json:"StreamNum,omitempty" xml:"StreamNum,omitempty"`
	// The specification of LindormStream nodes in the instance after the instance is upgraded. Valid values:
	//
	// 	- **lindorm.g.xlarge**: Each node has 4 dedicated CPU cores and 16 GB of dedicated memory.
	//
	// 	- **lindorm.g.2xlarge**: Each node has 8 dedicated CPU cores and 32 GB of dedicated memory.
	//
	// 	- **lindorm.g.4xlarge**: Each node has 16 dedicated CPU cores and 64 GB of dedicated memory.
	//
	// 	- **lindorm.g.8xlarge**: Each node has 32 dedicated CPU cores and 128 GB of dedicated memory.
	//
	// example:
	//
	// lindorm.g.xlarge
	StreamSpec *string `json:"StreamSpec,omitempty" xml:"StreamSpec,omitempty"`
	// The number of LindormTSDB nodes in the instance after the instance is upgraded. Valid values: integers from **0*	- to **24**.
	//
	// example:
	//
	// 2
	TsdbNum *int32 `json:"TsdbNum,omitempty" xml:"TsdbNum,omitempty"`
	// The specification of LindormTSDB nodes in the instance after the instance is upgraded. Valid values:
	//
	// 	- **lindorm.g.xlarge**: Each node has 4 dedicated CPU cores and 16 GB of dedicated memory.
	//
	// 	- **lindorm.g.2xlarge**: Each node has 8 dedicated CPU cores and 32 GB of dedicated memory.
	//
	// 	- **lindorm.g.4xlarge**: Each node has 16 dedicated CPU cores and 64 GB of dedicated memory.
	//
	// 	- **lindorm.g.8xlarge**: Each node has 32 dedicated CPU cores and 128 GB of dedicated memory.
	//
	// example:
	//
	// lindorm.g.xlarge
	TsdbSpec *string `json:"TsdbSpec,omitempty" xml:"TsdbSpec,omitempty"`
	// The upgrade type of the operation. For more information about upgrade types, see the UpgradeType parameters section.
	//
	// This parameter is required.
	//
	// example:
	//
	// upgrade-cold-storage
	UpgradeType *string `json:"UpgradeType,omitempty" xml:"UpgradeType,omitempty"`
	// The ID of the zone in which the instance that you want to upgrade, scale up, or enable cold storage is located. You can call the [GetLindormInstance](https://help.aliyun.com/document_detail/426067.html) operation to query the zone ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai-f
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s UpgradeLindormInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s UpgradeLindormInstanceRequest) GoString() string {
	return s.String()
}

func (s *UpgradeLindormInstanceRequest) SetClusterStorage(v int32) *UpgradeLindormInstanceRequest {
	s.ClusterStorage = &v
	return s
}

func (s *UpgradeLindormInstanceRequest) SetColdStorage(v int32) *UpgradeLindormInstanceRequest {
	s.ColdStorage = &v
	return s
}

func (s *UpgradeLindormInstanceRequest) SetCoreSingleStorage(v int32) *UpgradeLindormInstanceRequest {
	s.CoreSingleStorage = &v
	return s
}

func (s *UpgradeLindormInstanceRequest) SetFilestoreNum(v int32) *UpgradeLindormInstanceRequest {
	s.FilestoreNum = &v
	return s
}

func (s *UpgradeLindormInstanceRequest) SetFilestoreSpec(v string) *UpgradeLindormInstanceRequest {
	s.FilestoreSpec = &v
	return s
}

func (s *UpgradeLindormInstanceRequest) SetInstanceId(v string) *UpgradeLindormInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *UpgradeLindormInstanceRequest) SetLindormNum(v int32) *UpgradeLindormInstanceRequest {
	s.LindormNum = &v
	return s
}

func (s *UpgradeLindormInstanceRequest) SetLindormSpec(v string) *UpgradeLindormInstanceRequest {
	s.LindormSpec = &v
	return s
}

func (s *UpgradeLindormInstanceRequest) SetLogNum(v int32) *UpgradeLindormInstanceRequest {
	s.LogNum = &v
	return s
}

func (s *UpgradeLindormInstanceRequest) SetLogSingleStorage(v int32) *UpgradeLindormInstanceRequest {
	s.LogSingleStorage = &v
	return s
}

func (s *UpgradeLindormInstanceRequest) SetLogSpec(v string) *UpgradeLindormInstanceRequest {
	s.LogSpec = &v
	return s
}

func (s *UpgradeLindormInstanceRequest) SetLtsCoreNum(v int32) *UpgradeLindormInstanceRequest {
	s.LtsCoreNum = &v
	return s
}

func (s *UpgradeLindormInstanceRequest) SetLtsCoreSpec(v string) *UpgradeLindormInstanceRequest {
	s.LtsCoreSpec = &v
	return s
}

func (s *UpgradeLindormInstanceRequest) SetOwnerAccount(v string) *UpgradeLindormInstanceRequest {
	s.OwnerAccount = &v
	return s
}

func (s *UpgradeLindormInstanceRequest) SetOwnerId(v int64) *UpgradeLindormInstanceRequest {
	s.OwnerId = &v
	return s
}

func (s *UpgradeLindormInstanceRequest) SetRegionId(v string) *UpgradeLindormInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *UpgradeLindormInstanceRequest) SetResourceOwnerAccount(v string) *UpgradeLindormInstanceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UpgradeLindormInstanceRequest) SetResourceOwnerId(v int64) *UpgradeLindormInstanceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UpgradeLindormInstanceRequest) SetSecurityToken(v string) *UpgradeLindormInstanceRequest {
	s.SecurityToken = &v
	return s
}

func (s *UpgradeLindormInstanceRequest) SetSolrNum(v int32) *UpgradeLindormInstanceRequest {
	s.SolrNum = &v
	return s
}

func (s *UpgradeLindormInstanceRequest) SetSolrSpec(v string) *UpgradeLindormInstanceRequest {
	s.SolrSpec = &v
	return s
}

func (s *UpgradeLindormInstanceRequest) SetStreamNum(v int32) *UpgradeLindormInstanceRequest {
	s.StreamNum = &v
	return s
}

func (s *UpgradeLindormInstanceRequest) SetStreamSpec(v string) *UpgradeLindormInstanceRequest {
	s.StreamSpec = &v
	return s
}

func (s *UpgradeLindormInstanceRequest) SetTsdbNum(v int32) *UpgradeLindormInstanceRequest {
	s.TsdbNum = &v
	return s
}

func (s *UpgradeLindormInstanceRequest) SetTsdbSpec(v string) *UpgradeLindormInstanceRequest {
	s.TsdbSpec = &v
	return s
}

func (s *UpgradeLindormInstanceRequest) SetUpgradeType(v string) *UpgradeLindormInstanceRequest {
	s.UpgradeType = &v
	return s
}

func (s *UpgradeLindormInstanceRequest) SetZoneId(v string) *UpgradeLindormInstanceRequest {
	s.ZoneId = &v
	return s
}

type UpgradeLindormInstanceResponseBody struct {
	// The ID of the order.
	//
	// example:
	//
	// 111111111111111
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 2A7D4F9D-AA26-4E15-A2B1-3E4792C6****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpgradeLindormInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpgradeLindormInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *UpgradeLindormInstanceResponseBody) SetOrderId(v int64) *UpgradeLindormInstanceResponseBody {
	s.OrderId = &v
	return s
}

func (s *UpgradeLindormInstanceResponseBody) SetRequestId(v string) *UpgradeLindormInstanceResponseBody {
	s.RequestId = &v
	return s
}

type UpgradeLindormInstanceResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpgradeLindormInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpgradeLindormInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s UpgradeLindormInstanceResponse) GoString() string {
	return s.String()
}

func (s *UpgradeLindormInstanceResponse) SetHeaders(v map[string]*string) *UpgradeLindormInstanceResponse {
	s.Headers = v
	return s
}

func (s *UpgradeLindormInstanceResponse) SetStatusCode(v int32) *UpgradeLindormInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *UpgradeLindormInstanceResponse) SetBody(v *UpgradeLindormInstanceResponseBody) *UpgradeLindormInstanceResponse {
	s.Body = v
	return s
}

type UpgradeLindormV2StreamEngineRequest struct {
	// This parameter is required.
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	Quantity             *int32  `json:"Quantity,omitempty" xml:"Quantity,omitempty"`
	ResourceGroupName    *string `json:"ResourceGroupName,omitempty" xml:"ResourceGroupName,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// This parameter is required.
	Spec *string `json:"Spec,omitempty" xml:"Spec,omitempty"`
	// This parameter is required.
	SpecId *string `json:"SpecId,omitempty" xml:"SpecId,omitempty"`
	// This parameter is required.
	UpgradeType *string `json:"UpgradeType,omitempty" xml:"UpgradeType,omitempty"`
}

func (s UpgradeLindormV2StreamEngineRequest) String() string {
	return tea.Prettify(s)
}

func (s UpgradeLindormV2StreamEngineRequest) GoString() string {
	return s.String()
}

func (s *UpgradeLindormV2StreamEngineRequest) SetInstanceId(v string) *UpgradeLindormV2StreamEngineRequest {
	s.InstanceId = &v
	return s
}

func (s *UpgradeLindormV2StreamEngineRequest) SetOwnerAccount(v string) *UpgradeLindormV2StreamEngineRequest {
	s.OwnerAccount = &v
	return s
}

func (s *UpgradeLindormV2StreamEngineRequest) SetOwnerId(v int64) *UpgradeLindormV2StreamEngineRequest {
	s.OwnerId = &v
	return s
}

func (s *UpgradeLindormV2StreamEngineRequest) SetQuantity(v int32) *UpgradeLindormV2StreamEngineRequest {
	s.Quantity = &v
	return s
}

func (s *UpgradeLindormV2StreamEngineRequest) SetResourceGroupName(v string) *UpgradeLindormV2StreamEngineRequest {
	s.ResourceGroupName = &v
	return s
}

func (s *UpgradeLindormV2StreamEngineRequest) SetResourceOwnerAccount(v string) *UpgradeLindormV2StreamEngineRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UpgradeLindormV2StreamEngineRequest) SetResourceOwnerId(v int64) *UpgradeLindormV2StreamEngineRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UpgradeLindormV2StreamEngineRequest) SetSecurityToken(v string) *UpgradeLindormV2StreamEngineRequest {
	s.SecurityToken = &v
	return s
}

func (s *UpgradeLindormV2StreamEngineRequest) SetSpec(v string) *UpgradeLindormV2StreamEngineRequest {
	s.Spec = &v
	return s
}

func (s *UpgradeLindormV2StreamEngineRequest) SetSpecId(v string) *UpgradeLindormV2StreamEngineRequest {
	s.SpecId = &v
	return s
}

func (s *UpgradeLindormV2StreamEngineRequest) SetUpgradeType(v string) *UpgradeLindormV2StreamEngineRequest {
	s.UpgradeType = &v
	return s
}

type UpgradeLindormV2StreamEngineResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	RequestId          *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpgradeLindormV2StreamEngineResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpgradeLindormV2StreamEngineResponseBody) GoString() string {
	return s.String()
}

func (s *UpgradeLindormV2StreamEngineResponseBody) SetAccessDeniedDetail(v string) *UpgradeLindormV2StreamEngineResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *UpgradeLindormV2StreamEngineResponseBody) SetRequestId(v string) *UpgradeLindormV2StreamEngineResponseBody {
	s.RequestId = &v
	return s
}

type UpgradeLindormV2StreamEngineResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpgradeLindormV2StreamEngineResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpgradeLindormV2StreamEngineResponse) String() string {
	return tea.Prettify(s)
}

func (s UpgradeLindormV2StreamEngineResponse) GoString() string {
	return s.String()
}

func (s *UpgradeLindormV2StreamEngineResponse) SetHeaders(v map[string]*string) *UpgradeLindormV2StreamEngineResponse {
	s.Headers = v
	return s
}

func (s *UpgradeLindormV2StreamEngineResponse) SetStatusCode(v int32) *UpgradeLindormV2StreamEngineResponse {
	s.StatusCode = &v
	return s
}

func (s *UpgradeLindormV2StreamEngineResponse) SetBody(v *UpgradeLindormV2StreamEngineResponseBody) *UpgradeLindormV2StreamEngineResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("hitsdb"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// Changes a resource group to another.
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
		Version:     tea.String("2020-06-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ChangeResourceGroupResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ChangeResourceGroupResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Changes a resource group to another.
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

// @param request - CheckLdpsColumnarIndexStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckLdpsColumnarIndexStatusResponse
func (client *Client) CheckLdpsColumnarIndexStatusWithOptions(request *CheckLdpsColumnarIndexStatusRequest, runtime *util.RuntimeOptions) (_result *CheckLdpsColumnarIndexStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CheckLdpsColumnarIndexStatus"),
		Version:     tea.String("2020-06-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &CheckLdpsColumnarIndexStatusResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &CheckLdpsColumnarIndexStatusResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @param request - CheckLdpsColumnarIndexStatusRequest
//
// @return CheckLdpsColumnarIndexStatusResponse
func (client *Client) CheckLdpsColumnarIndexStatus(request *CheckLdpsColumnarIndexStatusRequest) (_result *CheckLdpsColumnarIndexStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CheckLdpsColumnarIndexStatusResponse{}
	_body, _err := client.CheckLdpsColumnarIndexStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param tmpReq - CreateAutoScalingConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAutoScalingConfigResponse
func (client *Client) CreateAutoScalingConfigWithOptions(tmpReq *CreateAutoScalingConfigRequest, runtime *util.RuntimeOptions) (_result *CreateAutoScalingConfigResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateAutoScalingConfigShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.ScaleRuleList)) {
		request.ScaleRuleListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ScaleRuleList, tea.String("ScaleRuleList"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConfigName)) {
		query["ConfigName"] = request.ConfigName
	}

	if !tea.BoolValue(util.IsUnset(request.EffectiveTimeEnd)) {
		query["EffectiveTimeEnd"] = request.EffectiveTimeEnd
	}

	if !tea.BoolValue(util.IsUnset(request.EffectiveTimeStart)) {
		query["EffectiveTimeStart"] = request.EffectiveTimeStart
	}

	if !tea.BoolValue(util.IsUnset(request.Enabled)) {
		query["Enabled"] = request.Enabled
	}

	if !tea.BoolValue(util.IsUnset(request.Engine)) {
		query["Engine"] = request.Engine
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.NodesMax)) {
		query["NodesMax"] = request.NodesMax
	}

	if !tea.BoolValue(util.IsUnset(request.NodesMin)) {
		query["NodesMin"] = request.NodesMin
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ScaleRuleListShrink)) {
		query["ScaleRuleList"] = request.ScaleRuleListShrink
	}

	if !tea.BoolValue(util.IsUnset(request.ScaleType)) {
		query["ScaleType"] = request.ScaleType
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !tea.BoolValue(util.IsUnset(request.SpecId)) {
		query["SpecId"] = request.SpecId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateAutoScalingConfig"),
		Version:     tea.String("2020-06-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &CreateAutoScalingConfigResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &CreateAutoScalingConfigResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @param request - CreateAutoScalingConfigRequest
//
// @return CreateAutoScalingConfigResponse
func (client *Client) CreateAutoScalingConfig(request *CreateAutoScalingConfigRequest) (_result *CreateAutoScalingConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateAutoScalingConfigResponse{}
	_body, _err := client.CreateAutoScalingConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - CreateAutoScalingRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAutoScalingRuleResponse
func (client *Client) CreateAutoScalingRuleWithOptions(request *CreateAutoScalingRuleRequest, runtime *util.RuntimeOptions) (_result *CreateAutoScalingRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConfigId)) {
		query["ConfigId"] = request.ConfigId
	}

	if !tea.BoolValue(util.IsUnset(request.Enabled)) {
		query["Enabled"] = request.Enabled
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.ObservationWindow)) {
		query["ObservationWindow"] = request.ObservationWindow
	}

	if !tea.BoolValue(util.IsUnset(request.OperationType)) {
		query["OperationType"] = request.OperationType
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RuleName)) {
		query["RuleName"] = request.RuleName
	}

	if !tea.BoolValue(util.IsUnset(request.RuleType)) {
		query["RuleType"] = request.RuleType
	}

	if !tea.BoolValue(util.IsUnset(request.ScaleInStep)) {
		query["ScaleInStep"] = request.ScaleInStep
	}

	if !tea.BoolValue(util.IsUnset(request.ScaleOutStep)) {
		query["ScaleOutStep"] = request.ScaleOutStep
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !tea.BoolValue(util.IsUnset(request.SilenceTime)) {
		query["SilenceTime"] = request.SilenceTime
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.TargetMetric)) {
		query["TargetMetric"] = request.TargetMetric
	}

	if !tea.BoolValue(util.IsUnset(request.TargetNodes)) {
		query["TargetNodes"] = request.TargetNodes
	}

	if !tea.BoolValue(util.IsUnset(request.ThresholdLower)) {
		query["ThresholdLower"] = request.ThresholdLower
	}

	if !tea.BoolValue(util.IsUnset(request.ThresholdUpper)) {
		query["ThresholdUpper"] = request.ThresholdUpper
	}

	if !tea.BoolValue(util.IsUnset(request.TriggerCronExpr)) {
		query["TriggerCronExpr"] = request.TriggerCronExpr
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateAutoScalingRule"),
		Version:     tea.String("2020-06-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &CreateAutoScalingRuleResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &CreateAutoScalingRuleResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @param request - CreateAutoScalingRuleRequest
//
// @return CreateAutoScalingRuleResponse
func (client *Client) CreateAutoScalingRule(request *CreateAutoScalingRuleRequest) (_result *CreateAutoScalingRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateAutoScalingRuleResponse{}
	_body, _err := client.CreateAutoScalingRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - CreateLdpsComputeGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateLdpsComputeGroupResponse
func (client *Client) CreateLdpsComputeGroupWithOptions(request *CreateLdpsComputeGroupRequest, runtime *util.RuntimeOptions) (_result *CreateLdpsComputeGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GroupName)) {
		query["GroupName"] = request.GroupName
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.Properties)) {
		query["Properties"] = request.Properties
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateLdpsComputeGroup"),
		Version:     tea.String("2020-06-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &CreateLdpsComputeGroupResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &CreateLdpsComputeGroupResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @param request - CreateLdpsComputeGroupRequest
//
// @return CreateLdpsComputeGroupResponse
func (client *Client) CreateLdpsComputeGroup(request *CreateLdpsComputeGroupRequest) (_result *CreateLdpsComputeGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateLdpsComputeGroupResponse{}
	_body, _err := client.CreateLdpsComputeGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a Lindorm instance.
//
// Description:
//
// You must select at least one engine when you create a Lindorm instance. For more information about how to select the storage type and engine type when you create a Lindorm instance, see [Select engine types](https://help.aliyun.com/document_detail/181971.html) and [Select storage types](https://help.aliyun.com/document_detail/174643.html).
//
// @param request - CreateLindormInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateLindormInstanceResponse
func (client *Client) CreateLindormInstanceWithOptions(request *CreateLindormInstanceRequest, runtime *util.RuntimeOptions) (_result *CreateLindormInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ArbiterVSwitchId)) {
		query["ArbiterVSwitchId"] = request.ArbiterVSwitchId
	}

	if !tea.BoolValue(util.IsUnset(request.ArbiterZoneId)) {
		query["ArbiterZoneId"] = request.ArbiterZoneId
	}

	if !tea.BoolValue(util.IsUnset(request.ArchVersion)) {
		query["ArchVersion"] = request.ArchVersion
	}

	if !tea.BoolValue(util.IsUnset(request.AutoRenewDuration)) {
		query["AutoRenewDuration"] = request.AutoRenewDuration
	}

	if !tea.BoolValue(util.IsUnset(request.AutoRenewal)) {
		query["AutoRenewal"] = request.AutoRenewal
	}

	if !tea.BoolValue(util.IsUnset(request.ColdStorage)) {
		query["ColdStorage"] = request.ColdStorage
	}

	if !tea.BoolValue(util.IsUnset(request.CoreSingleStorage)) {
		query["CoreSingleStorage"] = request.CoreSingleStorage
	}

	if !tea.BoolValue(util.IsUnset(request.CoreSpec)) {
		query["CoreSpec"] = request.CoreSpec
	}

	if !tea.BoolValue(util.IsUnset(request.DiskCategory)) {
		query["DiskCategory"] = request.DiskCategory
	}

	if !tea.BoolValue(util.IsUnset(request.Duration)) {
		query["Duration"] = request.Duration
	}

	if !tea.BoolValue(util.IsUnset(request.FilestoreNum)) {
		query["FilestoreNum"] = request.FilestoreNum
	}

	if !tea.BoolValue(util.IsUnset(request.FilestoreSpec)) {
		query["FilestoreSpec"] = request.FilestoreSpec
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceAlias)) {
		query["InstanceAlias"] = request.InstanceAlias
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceStorage)) {
		query["InstanceStorage"] = request.InstanceStorage
	}

	if !tea.BoolValue(util.IsUnset(request.LindormNum)) {
		query["LindormNum"] = request.LindormNum
	}

	if !tea.BoolValue(util.IsUnset(request.LindormSpec)) {
		query["LindormSpec"] = request.LindormSpec
	}

	if !tea.BoolValue(util.IsUnset(request.LogDiskCategory)) {
		query["LogDiskCategory"] = request.LogDiskCategory
	}

	if !tea.BoolValue(util.IsUnset(request.LogNum)) {
		query["LogNum"] = request.LogNum
	}

	if !tea.BoolValue(util.IsUnset(request.LogSingleStorage)) {
		query["LogSingleStorage"] = request.LogSingleStorage
	}

	if !tea.BoolValue(util.IsUnset(request.LogSpec)) {
		query["LogSpec"] = request.LogSpec
	}

	if !tea.BoolValue(util.IsUnset(request.LtsNum)) {
		query["LtsNum"] = request.LtsNum
	}

	if !tea.BoolValue(util.IsUnset(request.LtsSpec)) {
		query["LtsSpec"] = request.LtsSpec
	}

	if !tea.BoolValue(util.IsUnset(request.MultiZoneCombination)) {
		query["MultiZoneCombination"] = request.MultiZoneCombination
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PayType)) {
		query["PayType"] = request.PayType
	}

	if !tea.BoolValue(util.IsUnset(request.PricingCycle)) {
		query["PricingCycle"] = request.PricingCycle
	}

	if !tea.BoolValue(util.IsUnset(request.PrimaryVSwitchId)) {
		query["PrimaryVSwitchId"] = request.PrimaryVSwitchId
	}

	if !tea.BoolValue(util.IsUnset(request.PrimaryZoneId)) {
		query["PrimaryZoneId"] = request.PrimaryZoneId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !tea.BoolValue(util.IsUnset(request.SolrNum)) {
		query["SolrNum"] = request.SolrNum
	}

	if !tea.BoolValue(util.IsUnset(request.SolrSpec)) {
		query["SolrSpec"] = request.SolrSpec
	}

	if !tea.BoolValue(util.IsUnset(request.StandbyVSwitchId)) {
		query["StandbyVSwitchId"] = request.StandbyVSwitchId
	}

	if !tea.BoolValue(util.IsUnset(request.StandbyZoneId)) {
		query["StandbyZoneId"] = request.StandbyZoneId
	}

	if !tea.BoolValue(util.IsUnset(request.StreamNum)) {
		query["StreamNum"] = request.StreamNum
	}

	if !tea.BoolValue(util.IsUnset(request.StreamSpec)) {
		query["StreamSpec"] = request.StreamSpec
	}

	if !tea.BoolValue(util.IsUnset(request.TsdbNum)) {
		query["TsdbNum"] = request.TsdbNum
	}

	if !tea.BoolValue(util.IsUnset(request.TsdbSpec)) {
		query["TsdbSpec"] = request.TsdbSpec
	}

	if !tea.BoolValue(util.IsUnset(request.VPCId)) {
		query["VPCId"] = request.VPCId
	}

	if !tea.BoolValue(util.IsUnset(request.VSwitchId)) {
		query["VSwitchId"] = request.VSwitchId
	}

	if !tea.BoolValue(util.IsUnset(request.ZoneId)) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateLindormInstance"),
		Version:     tea.String("2020-06-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &CreateLindormInstanceResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &CreateLindormInstanceResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Creates a Lindorm instance.
//
// Description:
//
// You must select at least one engine when you create a Lindorm instance. For more information about how to select the storage type and engine type when you create a Lindorm instance, see [Select engine types](https://help.aliyun.com/document_detail/181971.html) and [Select storage types](https://help.aliyun.com/document_detail/174643.html).
//
// @param request - CreateLindormInstanceRequest
//
// @return CreateLindormInstanceResponse
func (client *Client) CreateLindormInstance(request *CreateLindormInstanceRequest) (_result *CreateLindormInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateLindormInstanceResponse{}
	_body, _err := client.CreateLindormInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - CreateLindormV2InstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateLindormV2InstanceResponse
func (client *Client) CreateLindormV2InstanceWithOptions(request *CreateLindormV2InstanceRequest, runtime *util.RuntimeOptions) (_result *CreateLindormV2InstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ArbiterVSwitchId)) {
		query["ArbiterVSwitchId"] = request.ArbiterVSwitchId
	}

	if !tea.BoolValue(util.IsUnset(request.ArbiterZoneId)) {
		query["ArbiterZoneId"] = request.ArbiterZoneId
	}

	if !tea.BoolValue(util.IsUnset(request.ArchVersion)) {
		query["ArchVersion"] = request.ArchVersion
	}

	if !tea.BoolValue(util.IsUnset(request.AutoRenewDuration)) {
		query["AutoRenewDuration"] = request.AutoRenewDuration
	}

	if !tea.BoolValue(util.IsUnset(request.AutoRenewal)) {
		query["AutoRenewal"] = request.AutoRenewal
	}

	if !tea.BoolValue(util.IsUnset(request.CapacityStorageSize)) {
		query["CapacityStorageSize"] = request.CapacityStorageSize
	}

	if !tea.BoolValue(util.IsUnset(request.CloudStorageSize)) {
		query["CloudStorageSize"] = request.CloudStorageSize
	}

	if !tea.BoolValue(util.IsUnset(request.CloudStorageType)) {
		query["CloudStorageType"] = request.CloudStorageType
	}

	if !tea.BoolValue(util.IsUnset(request.ClusterMode)) {
		query["ClusterMode"] = request.ClusterMode
	}

	if !tea.BoolValue(util.IsUnset(request.ClusterPattern)) {
		query["ClusterPattern"] = request.ClusterPattern
	}

	if !tea.BoolValue(util.IsUnset(request.Duration)) {
		query["Duration"] = request.Duration
	}

	if !tea.BoolValue(util.IsUnset(request.EnableCapacityStorage)) {
		query["EnableCapacityStorage"] = request.EnableCapacityStorage
	}

	if !tea.BoolValue(util.IsUnset(request.EngineList)) {
		query["EngineList"] = request.EngineList
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceAlias)) {
		query["InstanceAlias"] = request.InstanceAlias
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PayType)) {
		query["PayType"] = request.PayType
	}

	if !tea.BoolValue(util.IsUnset(request.PricingCycle)) {
		query["PricingCycle"] = request.PricingCycle
	}

	if !tea.BoolValue(util.IsUnset(request.PrimaryVSwitchId)) {
		query["PrimaryVSwitchId"] = request.PrimaryVSwitchId
	}

	if !tea.BoolValue(util.IsUnset(request.PrimaryZoneId)) {
		query["PrimaryZoneId"] = request.PrimaryZoneId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !tea.BoolValue(util.IsUnset(request.StandbyVSwitchId)) {
		query["StandbyVSwitchId"] = request.StandbyVSwitchId
	}

	if !tea.BoolValue(util.IsUnset(request.StandbyZoneId)) {
		query["StandbyZoneId"] = request.StandbyZoneId
	}

	if !tea.BoolValue(util.IsUnset(request.VPCId)) {
		query["VPCId"] = request.VPCId
	}

	if !tea.BoolValue(util.IsUnset(request.VSwitchId)) {
		query["VSwitchId"] = request.VSwitchId
	}

	if !tea.BoolValue(util.IsUnset(request.ZoneId)) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateLindormV2Instance"),
		Version:     tea.String("2020-06-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &CreateLindormV2InstanceResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &CreateLindormV2InstanceResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @param request - CreateLindormV2InstanceRequest
//
// @return CreateLindormV2InstanceResponse
func (client *Client) CreateLindormV2Instance(request *CreateLindormV2InstanceRequest) (_result *CreateLindormV2InstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateLindormV2InstanceResponse{}
	_body, _err := client.CreateLindormV2InstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DeleteAutoScalingConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAutoScalingConfigResponse
func (client *Client) DeleteAutoScalingConfigWithOptions(request *DeleteAutoScalingConfigRequest, runtime *util.RuntimeOptions) (_result *DeleteAutoScalingConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConfigId)) {
		query["ConfigId"] = request.ConfigId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteAutoScalingConfig"),
		Version:     tea.String("2020-06-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DeleteAutoScalingConfigResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DeleteAutoScalingConfigResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @param request - DeleteAutoScalingConfigRequest
//
// @return DeleteAutoScalingConfigResponse
func (client *Client) DeleteAutoScalingConfig(request *DeleteAutoScalingConfigRequest) (_result *DeleteAutoScalingConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteAutoScalingConfigResponse{}
	_body, _err := client.DeleteAutoScalingConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DeleteAutoScalingRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAutoScalingRuleResponse
func (client *Client) DeleteAutoScalingRuleWithOptions(request *DeleteAutoScalingRuleRequest, runtime *util.RuntimeOptions) (_result *DeleteAutoScalingRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConfigId)) {
		query["ConfigId"] = request.ConfigId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RuleId)) {
		query["RuleId"] = request.RuleId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteAutoScalingRule"),
		Version:     tea.String("2020-06-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DeleteAutoScalingRuleResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DeleteAutoScalingRuleResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @param request - DeleteAutoScalingRuleRequest
//
// @return DeleteAutoScalingRuleResponse
func (client *Client) DeleteAutoScalingRule(request *DeleteAutoScalingRuleRequest) (_result *DeleteAutoScalingRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteAutoScalingRuleResponse{}
	_body, _err := client.DeleteAutoScalingRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DeleteCustomResourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCustomResourceResponse
func (client *Client) DeleteCustomResourceWithOptions(request *DeleteCustomResourceRequest, runtime *util.RuntimeOptions) (_result *DeleteCustomResourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteCustomResource"),
		Version:     tea.String("2020-06-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DeleteCustomResourceResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DeleteCustomResourceResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @param request - DeleteCustomResourceRequest
//
// @return DeleteCustomResourceResponse
func (client *Client) DeleteCustomResource(request *DeleteCustomResourceRequest) (_result *DeleteCustomResourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteCustomResourceResponse{}
	_body, _err := client.DeleteCustomResourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DeleteLdpsComputeGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteLdpsComputeGroupResponse
func (client *Client) DeleteLdpsComputeGroupWithOptions(request *DeleteLdpsComputeGroupRequest, runtime *util.RuntimeOptions) (_result *DeleteLdpsComputeGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GroupName)) {
		query["GroupName"] = request.GroupName
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteLdpsComputeGroup"),
		Version:     tea.String("2020-06-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DeleteLdpsComputeGroupResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DeleteLdpsComputeGroupResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @param request - DeleteLdpsComputeGroupRequest
//
// @return DeleteLdpsComputeGroupResponse
func (client *Client) DeleteLdpsComputeGroup(request *DeleteLdpsComputeGroupRequest) (_result *DeleteLdpsComputeGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteLdpsComputeGroupResponse{}
	_body, _err := client.DeleteLdpsComputeGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DeployLdpsSemiManagedComponentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeployLdpsSemiManagedComponentResponse
func (client *Client) DeployLdpsSemiManagedComponentWithOptions(request *DeployLdpsSemiManagedComponentRequest, runtime *util.RuntimeOptions) (_result *DeployLdpsSemiManagedComponentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ComponentName)) {
		query["ComponentName"] = request.ComponentName
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeployLdpsSemiManagedComponent"),
		Version:     tea.String("2020-06-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DeployLdpsSemiManagedComponentResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DeployLdpsSemiManagedComponentResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @param request - DeployLdpsSemiManagedComponentRequest
//
// @return DeployLdpsSemiManagedComponentResponse
func (client *Client) DeployLdpsSemiManagedComponent(request *DeployLdpsSemiManagedComponentRequest) (_result *DeployLdpsSemiManagedComponentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeployLdpsSemiManagedComponentResponse{}
	_body, _err := client.DeployLdpsSemiManagedComponentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains the regions supported by Lindorm.
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

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeRegions"),
		Version:     tea.String("2020-06-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DescribeRegionsResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DescribeRegionsResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Obtains the regions supported by Lindorm.
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

// @param request - GetAutoScalingConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAutoScalingConfigResponse
func (client *Client) GetAutoScalingConfigWithOptions(request *GetAutoScalingConfigRequest, runtime *util.RuntimeOptions) (_result *GetAutoScalingConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConfigId)) {
		query["ConfigId"] = request.ConfigId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetAutoScalingConfig"),
		Version:     tea.String("2020-06-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetAutoScalingConfigResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetAutoScalingConfigResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @param request - GetAutoScalingConfigRequest
//
// @return GetAutoScalingConfigResponse
func (client *Client) GetAutoScalingConfig(request *GetAutoScalingConfigRequest) (_result *GetAutoScalingConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAutoScalingConfigResponse{}
	_body, _err := client.GetAutoScalingConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - GetAutoScalingRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAutoScalingRuleResponse
func (client *Client) GetAutoScalingRuleWithOptions(request *GetAutoScalingRuleRequest, runtime *util.RuntimeOptions) (_result *GetAutoScalingRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConfigId)) {
		query["ConfigId"] = request.ConfigId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RuleId)) {
		query["RuleId"] = request.RuleId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetAutoScalingRule"),
		Version:     tea.String("2020-06-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetAutoScalingRuleResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetAutoScalingRuleResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @param request - GetAutoScalingRuleRequest
//
// @return GetAutoScalingRuleResponse
func (client *Client) GetAutoScalingRule(request *GetAutoScalingRuleRequest) (_result *GetAutoScalingRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAutoScalingRuleResponse{}
	_body, _err := client.GetAutoScalingRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - GetClientSourceIpRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetClientSourceIpResponse
func (client *Client) GetClientSourceIpWithOptions(request *GetClientSourceIpRequest, runtime *util.RuntimeOptions) (_result *GetClientSourceIpResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetClientSourceIp"),
		Version:     tea.String("2020-06-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetClientSourceIpResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetClientSourceIpResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @param request - GetClientSourceIpRequest
//
// @return GetClientSourceIpResponse
func (client *Client) GetClientSourceIp(request *GetClientSourceIpRequest) (_result *GetClientSourceIpResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetClientSourceIpResponse{}
	_body, _err := client.GetClientSourceIpWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - GetEngineDefaultAuthRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetEngineDefaultAuthResponse
func (client *Client) GetEngineDefaultAuthWithOptions(request *GetEngineDefaultAuthRequest, runtime *util.RuntimeOptions) (_result *GetEngineDefaultAuthResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetEngineDefaultAuth"),
		Version:     tea.String("2020-06-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetEngineDefaultAuthResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetEngineDefaultAuthResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @param request - GetEngineDefaultAuthRequest
//
// @return GetEngineDefaultAuthResponse
func (client *Client) GetEngineDefaultAuth(request *GetEngineDefaultAuthRequest) (_result *GetEngineDefaultAuthResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetEngineDefaultAuthResponse{}
	_body, _err := client.GetEngineDefaultAuthWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the whitelists configured for a Lindorm instance.
//
// @param request - GetInstanceIpWhiteListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetInstanceIpWhiteListResponse
func (client *Client) GetInstanceIpWhiteListWithOptions(request *GetInstanceIpWhiteListRequest, runtime *util.RuntimeOptions) (_result *GetInstanceIpWhiteListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetInstanceIpWhiteList"),
		Version:     tea.String("2020-06-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetInstanceIpWhiteListResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetInstanceIpWhiteListResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Queries the whitelists configured for a Lindorm instance.
//
// @param request - GetInstanceIpWhiteListRequest
//
// @return GetInstanceIpWhiteListResponse
func (client *Client) GetInstanceIpWhiteList(request *GetInstanceIpWhiteListRequest) (_result *GetInstanceIpWhiteListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetInstanceIpWhiteListResponse{}
	_body, _err := client.GetInstanceIpWhiteListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - GetInstanceSecurityGroupsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetInstanceSecurityGroupsResponse
func (client *Client) GetInstanceSecurityGroupsWithOptions(request *GetInstanceSecurityGroupsRequest, runtime *util.RuntimeOptions) (_result *GetInstanceSecurityGroupsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetInstanceSecurityGroups"),
		Version:     tea.String("2020-06-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetInstanceSecurityGroupsResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetInstanceSecurityGroupsResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @param request - GetInstanceSecurityGroupsRequest
//
// @return GetInstanceSecurityGroupsResponse
func (client *Client) GetInstanceSecurityGroups(request *GetInstanceSecurityGroupsRequest) (_result *GetInstanceSecurityGroupsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetInstanceSecurityGroupsResponse{}
	_body, _err := client.GetInstanceSecurityGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - GetLdpsComputeGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetLdpsComputeGroupResponse
func (client *Client) GetLdpsComputeGroupWithOptions(request *GetLdpsComputeGroupRequest, runtime *util.RuntimeOptions) (_result *GetLdpsComputeGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GroupName)) {
		query["GroupName"] = request.GroupName
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetLdpsComputeGroup"),
		Version:     tea.String("2020-06-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetLdpsComputeGroupResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetLdpsComputeGroupResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @param request - GetLdpsComputeGroupRequest
//
// @return GetLdpsComputeGroupResponse
func (client *Client) GetLdpsComputeGroup(request *GetLdpsComputeGroupRequest) (_result *GetLdpsComputeGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetLdpsComputeGroupResponse{}
	_body, _err := client.GetLdpsComputeGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - GetLdpsNamespacedQuotaRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetLdpsNamespacedQuotaResponse
func (client *Client) GetLdpsNamespacedQuotaWithOptions(request *GetLdpsNamespacedQuotaRequest, runtime *util.RuntimeOptions) (_result *GetLdpsNamespacedQuotaResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Namespace)) {
		query["Namespace"] = request.Namespace
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetLdpsNamespacedQuota"),
		Version:     tea.String("2020-06-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetLdpsNamespacedQuotaResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetLdpsNamespacedQuotaResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @param request - GetLdpsNamespacedQuotaRequest
//
// @return GetLdpsNamespacedQuotaResponse
func (client *Client) GetLdpsNamespacedQuota(request *GetLdpsNamespacedQuotaRequest) (_result *GetLdpsNamespacedQuotaResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetLdpsNamespacedQuotaResponse{}
	_body, _err := client.GetLdpsNamespacedQuotaWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - GetLdpsResourceCostRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetLdpsResourceCostResponse
func (client *Client) GetLdpsResourceCostWithOptions(request *GetLdpsResourceCostRequest, runtime *util.RuntimeOptions) (_result *GetLdpsResourceCostResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.JobId)) {
		query["JobId"] = request.JobId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetLdpsResourceCost"),
		Version:     tea.String("2020-06-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetLdpsResourceCostResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetLdpsResourceCostResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @param request - GetLdpsResourceCostRequest
//
// @return GetLdpsResourceCostResponse
func (client *Client) GetLdpsResourceCost(request *GetLdpsResourceCostRequest) (_result *GetLdpsResourceCostResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetLdpsResourceCostResponse{}
	_body, _err := client.GetLdpsResourceCostWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - GetLindormFsUsedDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetLindormFsUsedDetailResponse
func (client *Client) GetLindormFsUsedDetailWithOptions(request *GetLindormFsUsedDetailRequest, runtime *util.RuntimeOptions) (_result *GetLindormFsUsedDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetLindormFsUsedDetail"),
		Version:     tea.String("2020-06-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetLindormFsUsedDetailResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetLindormFsUsedDetailResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @param request - GetLindormFsUsedDetailRequest
//
// @return GetLindormFsUsedDetailResponse
func (client *Client) GetLindormFsUsedDetail(request *GetLindormFsUsedDetailRequest) (_result *GetLindormFsUsedDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetLindormFsUsedDetailResponse{}
	_body, _err := client.GetLindormFsUsedDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains the detailed information about a Lindorm instance, including the instance type, billing method, and VPC.
//
// @param request - GetLindormInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetLindormInstanceResponse
func (client *Client) GetLindormInstanceWithOptions(request *GetLindormInstanceRequest, runtime *util.RuntimeOptions) (_result *GetLindormInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetLindormInstance"),
		Version:     tea.String("2020-06-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetLindormInstanceResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetLindormInstanceResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Obtains the detailed information about a Lindorm instance, including the instance type, billing method, and VPC.
//
// @param request - GetLindormInstanceRequest
//
// @return GetLindormInstanceResponse
func (client *Client) GetLindormInstance(request *GetLindormInstanceRequest) (_result *GetLindormInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetLindormInstanceResponse{}
	_body, _err := client.GetLindormInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains the engine types supported by the specified Lindorm instance.
//
// @param request - GetLindormInstanceEngineListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetLindormInstanceEngineListResponse
func (client *Client) GetLindormInstanceEngineListWithOptions(request *GetLindormInstanceEngineListRequest, runtime *util.RuntimeOptions) (_result *GetLindormInstanceEngineListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetLindormInstanceEngineList"),
		Version:     tea.String("2020-06-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetLindormInstanceEngineListResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetLindormInstanceEngineListResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Obtains the engine types supported by the specified Lindorm instance.
//
// @param request - GetLindormInstanceEngineListRequest
//
// @return GetLindormInstanceEngineListResponse
func (client *Client) GetLindormInstanceEngineList(request *GetLindormInstanceEngineListRequest) (_result *GetLindormInstanceEngineListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetLindormInstanceEngineListResponse{}
	_body, _err := client.GetLindormInstanceEngineListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the instances that meet the specified conditions.
//
// @param request - GetLindormInstanceListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetLindormInstanceListResponse
func (client *Client) GetLindormInstanceListWithOptions(request *GetLindormInstanceListRequest, runtime *util.RuntimeOptions) (_result *GetLindormInstanceListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.QueryStr)) {
		query["QueryStr"] = request.QueryStr
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceType)) {
		query["ServiceType"] = request.ServiceType
	}

	if !tea.BoolValue(util.IsUnset(request.SupportEngine)) {
		query["SupportEngine"] = request.SupportEngine
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetLindormInstanceList"),
		Version:     tea.String("2020-06-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetLindormInstanceListResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetLindormInstanceListResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Queries the instances that meet the specified conditions.
//
// @param request - GetLindormInstanceListRequest
//
// @return GetLindormInstanceListResponse
func (client *Client) GetLindormInstanceList(request *GetLindormInstanceListRequest) (_result *GetLindormInstanceListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetLindormInstanceListResponse{}
	_body, _err := client.GetLindormInstanceListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - GetLindormV2InstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetLindormV2InstanceResponse
func (client *Client) GetLindormV2InstanceWithOptions(request *GetLindormV2InstanceRequest, runtime *util.RuntimeOptions) (_result *GetLindormV2InstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetLindormV2Instance"),
		Version:     tea.String("2020-06-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetLindormV2InstanceResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetLindormV2InstanceResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @param request - GetLindormV2InstanceRequest
//
// @return GetLindormV2InstanceResponse
func (client *Client) GetLindormV2Instance(request *GetLindormV2InstanceRequest) (_result *GetLindormV2InstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetLindormV2InstanceResponse{}
	_body, _err := client.GetLindormV2InstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - GetLindormV2InstanceEngineListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetLindormV2InstanceEngineListResponse
func (client *Client) GetLindormV2InstanceEngineListWithOptions(request *GetLindormV2InstanceEngineListRequest, runtime *util.RuntimeOptions) (_result *GetLindormV2InstanceEngineListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetLindormV2InstanceEngineList"),
		Version:     tea.String("2020-06-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetLindormV2InstanceEngineListResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetLindormV2InstanceEngineListResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @param request - GetLindormV2InstanceEngineListRequest
//
// @return GetLindormV2InstanceEngineListResponse
func (client *Client) GetLindormV2InstanceEngineList(request *GetLindormV2InstanceEngineListRequest) (_result *GetLindormV2InstanceEngineListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetLindormV2InstanceEngineListResponse{}
	_body, _err := client.GetLindormV2InstanceEngineListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - GetLindormV2StorageUsageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetLindormV2StorageUsageResponse
func (client *Client) GetLindormV2StorageUsageWithOptions(request *GetLindormV2StorageUsageRequest, runtime *util.RuntimeOptions) (_result *GetLindormV2StorageUsageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetLindormV2StorageUsage"),
		Version:     tea.String("2020-06-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetLindormV2StorageUsageResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetLindormV2StorageUsageResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @param request - GetLindormV2StorageUsageRequest
//
// @return GetLindormV2StorageUsageResponse
func (client *Client) GetLindormV2StorageUsage(request *GetLindormV2StorageUsageRequest) (_result *GetLindormV2StorageUsageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetLindormV2StorageUsageResponse{}
	_body, _err := client.GetLindormV2StorageUsageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ListAutoScalingConfigsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAutoScalingConfigsResponse
func (client *Client) ListAutoScalingConfigsWithOptions(request *ListAutoScalingConfigsRequest, runtime *util.RuntimeOptions) (_result *ListAutoScalingConfigsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListAutoScalingConfigs"),
		Version:     tea.String("2020-06-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListAutoScalingConfigsResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListAutoScalingConfigsResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @param request - ListAutoScalingConfigsRequest
//
// @return ListAutoScalingConfigsResponse
func (client *Client) ListAutoScalingConfigs(request *ListAutoScalingConfigsRequest) (_result *ListAutoScalingConfigsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListAutoScalingConfigsResponse{}
	_body, _err := client.ListAutoScalingConfigsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ListAutoScalingRecordsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAutoScalingRecordsResponse
func (client *Client) ListAutoScalingRecordsWithOptions(request *ListAutoScalingRecordsRequest, runtime *util.RuntimeOptions) (_result *ListAutoScalingRecordsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNum)) {
		query["PageNum"] = request.PageNum
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListAutoScalingRecords"),
		Version:     tea.String("2020-06-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListAutoScalingRecordsResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListAutoScalingRecordsResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @param request - ListAutoScalingRecordsRequest
//
// @return ListAutoScalingRecordsResponse
func (client *Client) ListAutoScalingRecords(request *ListAutoScalingRecordsRequest) (_result *ListAutoScalingRecordsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListAutoScalingRecordsResponse{}
	_body, _err := client.ListAutoScalingRecordsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ListAutoScalingRulesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAutoScalingRulesResponse
func (client *Client) ListAutoScalingRulesWithOptions(request *ListAutoScalingRulesRequest, runtime *util.RuntimeOptions) (_result *ListAutoScalingRulesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConfigId)) {
		query["ConfigId"] = request.ConfigId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListAutoScalingRules"),
		Version:     tea.String("2020-06-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListAutoScalingRulesResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListAutoScalingRulesResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @param request - ListAutoScalingRulesRequest
//
// @return ListAutoScalingRulesResponse
func (client *Client) ListAutoScalingRules(request *ListAutoScalingRulesRequest) (_result *ListAutoScalingRulesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListAutoScalingRulesResponse{}
	_body, _err := client.ListAutoScalingRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ListLdpsComputeGroupsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListLdpsComputeGroupsResponse
func (client *Client) ListLdpsComputeGroupsWithOptions(request *ListLdpsComputeGroupsRequest, runtime *util.RuntimeOptions) (_result *ListLdpsComputeGroupsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListLdpsComputeGroups"),
		Version:     tea.String("2020-06-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListLdpsComputeGroupsResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListLdpsComputeGroupsResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @param request - ListLdpsComputeGroupsRequest
//
// @return ListLdpsComputeGroupsResponse
func (client *Client) ListLdpsComputeGroups(request *ListLdpsComputeGroupsRequest) (_result *ListLdpsComputeGroupsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListLdpsComputeGroupsResponse{}
	_body, _err := client.ListLdpsComputeGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the tags associated with the specified Lindorm instance.
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

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTagResources"),
		Version:     tea.String("2020-06-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListTagResourcesResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListTagResourcesResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Queries the tags associated with the specified Lindorm instance.
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

// @param request - ModifyAutoScalingConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyAutoScalingConfigResponse
func (client *Client) ModifyAutoScalingConfigWithOptions(request *ModifyAutoScalingConfigRequest, runtime *util.RuntimeOptions) (_result *ModifyAutoScalingConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConfigId)) {
		query["ConfigId"] = request.ConfigId
	}

	if !tea.BoolValue(util.IsUnset(request.ConfigName)) {
		query["ConfigName"] = request.ConfigName
	}

	if !tea.BoolValue(util.IsUnset(request.EffectiveTimeEnd)) {
		query["EffectiveTimeEnd"] = request.EffectiveTimeEnd
	}

	if !tea.BoolValue(util.IsUnset(request.EffectiveTimeStart)) {
		query["EffectiveTimeStart"] = request.EffectiveTimeStart
	}

	if !tea.BoolValue(util.IsUnset(request.Enabled)) {
		query["Enabled"] = request.Enabled
	}

	if !tea.BoolValue(util.IsUnset(request.Engine)) {
		query["Engine"] = request.Engine
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.NodesMax)) {
		query["NodesMax"] = request.NodesMax
	}

	if !tea.BoolValue(util.IsUnset(request.NodesMin)) {
		query["NodesMin"] = request.NodesMin
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ScaleType)) {
		query["ScaleType"] = request.ScaleType
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !tea.BoolValue(util.IsUnset(request.SpecId)) {
		query["SpecId"] = request.SpecId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyAutoScalingConfig"),
		Version:     tea.String("2020-06-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ModifyAutoScalingConfigResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ModifyAutoScalingConfigResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @param request - ModifyAutoScalingConfigRequest
//
// @return ModifyAutoScalingConfigResponse
func (client *Client) ModifyAutoScalingConfig(request *ModifyAutoScalingConfigRequest) (_result *ModifyAutoScalingConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyAutoScalingConfigResponse{}
	_body, _err := client.ModifyAutoScalingConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ModifyAutoScalingRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyAutoScalingRuleResponse
func (client *Client) ModifyAutoScalingRuleWithOptions(request *ModifyAutoScalingRuleRequest, runtime *util.RuntimeOptions) (_result *ModifyAutoScalingRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConfigId)) {
		query["ConfigId"] = request.ConfigId
	}

	if !tea.BoolValue(util.IsUnset(request.Enabled)) {
		query["Enabled"] = request.Enabled
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.ObservationWindow)) {
		query["ObservationWindow"] = request.ObservationWindow
	}

	if !tea.BoolValue(util.IsUnset(request.OperationType)) {
		query["OperationType"] = request.OperationType
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RuleId)) {
		query["RuleId"] = request.RuleId
	}

	if !tea.BoolValue(util.IsUnset(request.RuleName)) {
		query["RuleName"] = request.RuleName
	}

	if !tea.BoolValue(util.IsUnset(request.RuleType)) {
		query["RuleType"] = request.RuleType
	}

	if !tea.BoolValue(util.IsUnset(request.ScaleInStep)) {
		query["ScaleInStep"] = request.ScaleInStep
	}

	if !tea.BoolValue(util.IsUnset(request.ScaleOutStep)) {
		query["ScaleOutStep"] = request.ScaleOutStep
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !tea.BoolValue(util.IsUnset(request.SilenceTime)) {
		query["SilenceTime"] = request.SilenceTime
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.TargetMetric)) {
		query["TargetMetric"] = request.TargetMetric
	}

	if !tea.BoolValue(util.IsUnset(request.TargetNodes)) {
		query["TargetNodes"] = request.TargetNodes
	}

	if !tea.BoolValue(util.IsUnset(request.ThresholdLower)) {
		query["ThresholdLower"] = request.ThresholdLower
	}

	if !tea.BoolValue(util.IsUnset(request.ThresholdUpper)) {
		query["ThresholdUpper"] = request.ThresholdUpper
	}

	if !tea.BoolValue(util.IsUnset(request.TriggerCronExpr)) {
		query["TriggerCronExpr"] = request.TriggerCronExpr
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyAutoScalingRule"),
		Version:     tea.String("2020-06-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ModifyAutoScalingRuleResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ModifyAutoScalingRuleResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @param request - ModifyAutoScalingRuleRequest
//
// @return ModifyAutoScalingRuleResponse
func (client *Client) ModifyAutoScalingRule(request *ModifyAutoScalingRuleRequest) (_result *ModifyAutoScalingRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyAutoScalingRuleResponse{}
	_body, _err := client.ModifyAutoScalingRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Changes the billing method of the specified Lindorm instance.
//
// Description:
//
// You can call this operation to change the billing method of an instance to subscription or pay-as-you-go.
//
// Before you call this operation, make sure that you fully understand the billing methods and [pricing](https://www.alibabacloud.com/en/pricing-calculator?spm=a2c63.p38356.0.0.2b024c2adcHeXL&_p_lc=1#/commodity/hitsdb_lindormpre_public_intl) of Lindorm. Published on only international site (alibabacloud.com).
//
// @param request - ModifyInstancePayTypeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyInstancePayTypeResponse
func (client *Client) ModifyInstancePayTypeWithOptions(request *ModifyInstancePayTypeRequest, runtime *util.RuntimeOptions) (_result *ModifyInstancePayTypeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Duration)) {
		query["Duration"] = request.Duration
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PayType)) {
		query["PayType"] = request.PayType
	}

	if !tea.BoolValue(util.IsUnset(request.PricingCycle)) {
		query["PricingCycle"] = request.PricingCycle
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyInstancePayType"),
		Version:     tea.String("2020-06-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ModifyInstancePayTypeResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ModifyInstancePayTypeResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Changes the billing method of the specified Lindorm instance.
//
// Description:
//
// You can call this operation to change the billing method of an instance to subscription or pay-as-you-go.
//
// Before you call this operation, make sure that you fully understand the billing methods and [pricing](https://www.alibabacloud.com/en/pricing-calculator?spm=a2c63.p38356.0.0.2b024c2adcHeXL&_p_lc=1#/commodity/hitsdb_lindormpre_public_intl) of Lindorm. Published on only international site (alibabacloud.com).
//
// @param request - ModifyInstancePayTypeRequest
//
// @return ModifyInstancePayTypeResponse
func (client *Client) ModifyInstancePayType(request *ModifyInstancePayTypeRequest) (_result *ModifyInstancePayTypeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyInstancePayTypeResponse{}
	_body, _err := client.ModifyInstancePayTypeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ModifyLindormV2InstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyLindormV2InstanceResponse
func (client *Client) ModifyLindormV2InstanceWithOptions(request *ModifyLindormV2InstanceRequest, runtime *util.RuntimeOptions) (_result *ModifyLindormV2InstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CloudStorageSize)) {
		query["CloudStorageSize"] = request.CloudStorageSize
	}

	if !tea.BoolValue(util.IsUnset(request.CloudStorageType)) {
		query["CloudStorageType"] = request.CloudStorageType
	}

	if !tea.BoolValue(util.IsUnset(request.EngineType)) {
		query["EngineType"] = request.EngineType
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.NodeGroupList)) {
		query["NodeGroupList"] = request.NodeGroupList
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !tea.BoolValue(util.IsUnset(request.UpgradeType)) {
		query["UpgradeType"] = request.UpgradeType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyLindormV2Instance"),
		Version:     tea.String("2020-06-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ModifyLindormV2InstanceResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ModifyLindormV2InstanceResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @param request - ModifyLindormV2InstanceRequest
//
// @return ModifyLindormV2InstanceResponse
func (client *Client) ModifyLindormV2Instance(request *ModifyLindormV2InstanceRequest) (_result *ModifyLindormV2InstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyLindormV2InstanceResponse{}
	_body, _err := client.ModifyLindormV2InstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ModifyLindormV2WhiteIpListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyLindormV2WhiteIpListResponse
func (client *Client) ModifyLindormV2WhiteIpListWithOptions(request *ModifyLindormV2WhiteIpListRequest, runtime *util.RuntimeOptions) (_result *ModifyLindormV2WhiteIpListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeleteGroup)) {
		query["DeleteGroup"] = request.DeleteGroup
	}

	if !tea.BoolValue(util.IsUnset(request.GroupName)) {
		query["GroupName"] = request.GroupName
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !tea.BoolValue(util.IsUnset(request.WhiteIpList)) {
		query["WhiteIpList"] = request.WhiteIpList
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyLindormV2WhiteIpList"),
		Version:     tea.String("2020-06-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ModifyLindormV2WhiteIpListResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ModifyLindormV2WhiteIpListResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @param request - ModifyLindormV2WhiteIpListRequest
//
// @return ModifyLindormV2WhiteIpListResponse
func (client *Client) ModifyLindormV2WhiteIpList(request *ModifyLindormV2WhiteIpListRequest) (_result *ModifyLindormV2WhiteIpListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyLindormV2WhiteIpListResponse{}
	_body, _err := client.ModifyLindormV2WhiteIpListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - OpenComputeEngineRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OpenComputeEngineResponse
func (client *Client) OpenComputeEngineWithOptions(request *OpenComputeEngineRequest, runtime *util.RuntimeOptions) (_result *OpenComputeEngineResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CpuLimit)) {
		query["CpuLimit"] = request.CpuLimit
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.MemoryLimit)) {
		query["MemoryLimit"] = request.MemoryLimit
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("OpenComputeEngine"),
		Version:     tea.String("2020-06-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &OpenComputeEngineResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &OpenComputeEngineResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @param request - OpenComputeEngineRequest
//
// @return OpenComputeEngineResponse
func (client *Client) OpenComputeEngine(request *OpenComputeEngineRequest) (_result *OpenComputeEngineResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &OpenComputeEngineResponse{}
	_body, _err := client.OpenComputeEngineWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - OpenComputePreCheckRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OpenComputePreCheckResponse
func (client *Client) OpenComputePreCheckWithOptions(request *OpenComputePreCheckRequest, runtime *util.RuntimeOptions) (_result *OpenComputePreCheckResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CpuLimit)) {
		query["CpuLimit"] = request.CpuLimit
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.MemoryLimit)) {
		query["MemoryLimit"] = request.MemoryLimit
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("OpenComputePreCheck"),
		Version:     tea.String("2020-06-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &OpenComputePreCheckResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &OpenComputePreCheckResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @param request - OpenComputePreCheckRequest
//
// @return OpenComputePreCheckResponse
func (client *Client) OpenComputePreCheck(request *OpenComputePreCheckRequest) (_result *OpenComputePreCheckResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &OpenComputePreCheckResponse{}
	_body, _err := client.OpenComputePreCheckWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Releases a Lindorm instance.
//
// @param request - ReleaseLindormInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReleaseLindormInstanceResponse
func (client *Client) ReleaseLindormInstanceWithOptions(request *ReleaseLindormInstanceRequest, runtime *util.RuntimeOptions) (_result *ReleaseLindormInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Immediately)) {
		query["Immediately"] = request.Immediately
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ReleaseLindormInstance"),
		Version:     tea.String("2020-06-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ReleaseLindormInstanceResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ReleaseLindormInstanceResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Releases a Lindorm instance.
//
// @param request - ReleaseLindormInstanceRequest
//
// @return ReleaseLindormInstanceResponse
func (client *Client) ReleaseLindormInstance(request *ReleaseLindormInstanceRequest) (_result *ReleaseLindormInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ReleaseLindormInstanceResponse{}
	_body, _err := client.ReleaseLindormInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ReleaseLindormV2InstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReleaseLindormV2InstanceResponse
func (client *Client) ReleaseLindormV2InstanceWithOptions(request *ReleaseLindormV2InstanceRequest, runtime *util.RuntimeOptions) (_result *ReleaseLindormV2InstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Immediately)) {
		query["Immediately"] = request.Immediately
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ReleaseLindormV2Instance"),
		Version:     tea.String("2020-06-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ReleaseLindormV2InstanceResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ReleaseLindormV2InstanceResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @param request - ReleaseLindormV2InstanceRequest
//
// @return ReleaseLindormV2InstanceResponse
func (client *Client) ReleaseLindormV2Instance(request *ReleaseLindormV2InstanceRequest) (_result *ReleaseLindormV2InstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ReleaseLindormV2InstanceResponse{}
	_body, _err := client.ReleaseLindormV2InstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Renews a subscription Lindorm instance.
//
// Description:
//
// You can call this operation to renew a subscription Lindorm instance for 1 to 9 months or 1 to 3 years.
//
// Before you call this operation, make sure that you fully understand the billing methods and pricing of Lindorm.
//
// @param request - RenewLindormInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RenewLindormInstanceResponse
func (client *Client) RenewLindormInstanceWithOptions(request *RenewLindormInstanceRequest, runtime *util.RuntimeOptions) (_result *RenewLindormInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Duration)) {
		query["Duration"] = request.Duration
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PricingCycle)) {
		query["PricingCycle"] = request.PricingCycle
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RenewLindormInstance"),
		Version:     tea.String("2020-06-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &RenewLindormInstanceResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &RenewLindormInstanceResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Renews a subscription Lindorm instance.
//
// Description:
//
// You can call this operation to renew a subscription Lindorm instance for 1 to 9 months or 1 to 3 years.
//
// Before you call this operation, make sure that you fully understand the billing methods and pricing of Lindorm.
//
// @param request - RenewLindormInstanceRequest
//
// @return RenewLindormInstanceResponse
func (client *Client) RenewLindormInstance(request *RenewLindormInstanceRequest) (_result *RenewLindormInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RenewLindormInstanceResponse{}
	_body, _err := client.RenewLindormInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - RestartLdpsComputeGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RestartLdpsComputeGroupResponse
func (client *Client) RestartLdpsComputeGroupWithOptions(request *RestartLdpsComputeGroupRequest, runtime *util.RuntimeOptions) (_result *RestartLdpsComputeGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GroupName)) {
		query["GroupName"] = request.GroupName
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RestartLdpsComputeGroup"),
		Version:     tea.String("2020-06-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &RestartLdpsComputeGroupResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &RestartLdpsComputeGroupResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @param request - RestartLdpsComputeGroupRequest
//
// @return RestartLdpsComputeGroupResponse
func (client *Client) RestartLdpsComputeGroup(request *RestartLdpsComputeGroupRequest) (_result *RestartLdpsComputeGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RestartLdpsComputeGroupResponse{}
	_body, _err := client.RestartLdpsComputeGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - SetDefaultOlapComputeGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetDefaultOlapComputeGroupResponse
func (client *Client) SetDefaultOlapComputeGroupWithOptions(request *SetDefaultOlapComputeGroupRequest, runtime *util.RuntimeOptions) (_result *SetDefaultOlapComputeGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GroupName)) {
		query["GroupName"] = request.GroupName
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.IsDefault)) {
		query["IsDefault"] = request.IsDefault
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SetDefaultOlapComputeGroup"),
		Version:     tea.String("2020-06-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &SetDefaultOlapComputeGroupResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &SetDefaultOlapComputeGroupResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @param request - SetDefaultOlapComputeGroupRequest
//
// @return SetDefaultOlapComputeGroupResponse
func (client *Client) SetDefaultOlapComputeGroup(request *SetDefaultOlapComputeGroupRequest) (_result *SetDefaultOlapComputeGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetDefaultOlapComputeGroupResponse{}
	_body, _err := client.SetDefaultOlapComputeGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Enables or disables the MySQL compatibility feature for a Lindorm instance.
//
// Description:
//
// Prerequisites
//
// 	- The LindormTable version of your instance is 2.6.0 or later.
//
// 	- The LindormTable of your instance supports LindormSQL V3. The value of the EnableLsqlVersionV3 parameter in the response of the GetLindormInstance operation is true for Lindorm instances purchased after Oct 24, 2023, which indicates that LindormSQL is supported by these instances by default. If you want to enable LindormSQL for instances purchased before Oct 24, 2023, contact the on-duty technical support.
//
// You can enable the MySQL compatibility feature for a Lindorm instance only when the instance meets the preceding requirements.
//
// @param request - SwitchLSQLV3MySQLServiceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SwitchLSQLV3MySQLServiceResponse
func (client *Client) SwitchLSQLV3MySQLServiceWithOptions(request *SwitchLSQLV3MySQLServiceRequest, runtime *util.RuntimeOptions) (_result *SwitchLSQLV3MySQLServiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ActionType)) {
		query["ActionType"] = request.ActionType
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SwitchLSQLV3MySQLService"),
		Version:     tea.String("2020-06-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &SwitchLSQLV3MySQLServiceResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &SwitchLSQLV3MySQLServiceResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Enables or disables the MySQL compatibility feature for a Lindorm instance.
//
// Description:
//
// Prerequisites
//
// 	- The LindormTable version of your instance is 2.6.0 or later.
//
// 	- The LindormTable of your instance supports LindormSQL V3. The value of the EnableLsqlVersionV3 parameter in the response of the GetLindormInstance operation is true for Lindorm instances purchased after Oct 24, 2023, which indicates that LindormSQL is supported by these instances by default. If you want to enable LindormSQL for instances purchased before Oct 24, 2023, contact the on-duty technical support.
//
// You can enable the MySQL compatibility feature for a Lindorm instance only when the instance meets the preceding requirements.
//
// @param request - SwitchLSQLV3MySQLServiceRequest
//
// @return SwitchLSQLV3MySQLServiceResponse
func (client *Client) SwitchLSQLV3MySQLService(request *SwitchLSQLV3MySQLServiceRequest) (_result *SwitchLSQLV3MySQLServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SwitchLSQLV3MySQLServiceResponse{}
	_body, _err := client.SwitchLSQLV3MySQLServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Associates tags with a single or multiple Lindorm instances.
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
	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("TagResources"),
		Version:     tea.String("2020-06-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &TagResourcesResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &TagResourcesResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Associates tags with a single or multiple Lindorm instances.
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
// Removes tags from a Lindorm instance.
//
// Description:
//
// If a tag is not added to any Lindorm instance, it is deleted.
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

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !tea.BoolValue(util.IsUnset(request.TagKey)) {
		query["TagKey"] = request.TagKey
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UntagResources"),
		Version:     tea.String("2020-06-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &UntagResourcesResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &UntagResourcesResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Removes tags from a Lindorm instance.
//
// Description:
//
// If a tag is not added to any Lindorm instance, it is deleted.
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
// Configures an IP address whitelist for a Lindorm instance.
//
// @param request - UpdateInstanceIpWhiteListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateInstanceIpWhiteListResponse
func (client *Client) UpdateInstanceIpWhiteListWithOptions(request *UpdateInstanceIpWhiteListRequest, runtime *util.RuntimeOptions) (_result *UpdateInstanceIpWhiteListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Delete)) {
		query["Delete"] = request.Delete
	}

	if !tea.BoolValue(util.IsUnset(request.GroupName)) {
		query["GroupName"] = request.GroupName
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityIpList)) {
		query["SecurityIpList"] = request.SecurityIpList
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateInstanceIpWhiteList"),
		Version:     tea.String("2020-06-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &UpdateInstanceIpWhiteListResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &UpdateInstanceIpWhiteListResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Configures an IP address whitelist for a Lindorm instance.
//
// @param request - UpdateInstanceIpWhiteListRequest
//
// @return UpdateInstanceIpWhiteListResponse
func (client *Client) UpdateInstanceIpWhiteList(request *UpdateInstanceIpWhiteListRequest) (_result *UpdateInstanceIpWhiteListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateInstanceIpWhiteListResponse{}
	_body, _err := client.UpdateInstanceIpWhiteListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - UpdateInstanceSecurityGroupsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateInstanceSecurityGroupsResponse
func (client *Client) UpdateInstanceSecurityGroupsWithOptions(request *UpdateInstanceSecurityGroupsRequest, runtime *util.RuntimeOptions) (_result *UpdateInstanceSecurityGroupsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityGroups)) {
		query["SecurityGroups"] = request.SecurityGroups
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateInstanceSecurityGroups"),
		Version:     tea.String("2020-06-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &UpdateInstanceSecurityGroupsResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &UpdateInstanceSecurityGroupsResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @param request - UpdateInstanceSecurityGroupsRequest
//
// @return UpdateInstanceSecurityGroupsResponse
func (client *Client) UpdateInstanceSecurityGroups(request *UpdateInstanceSecurityGroupsRequest) (_result *UpdateInstanceSecurityGroupsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateInstanceSecurityGroupsResponse{}
	_body, _err := client.UpdateInstanceSecurityGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - UpdateLdpsComputeGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateLdpsComputeGroupResponse
func (client *Client) UpdateLdpsComputeGroupWithOptions(request *UpdateLdpsComputeGroupRequest, runtime *util.RuntimeOptions) (_result *UpdateLdpsComputeGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GroupName)) {
		query["GroupName"] = request.GroupName
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.Properties)) {
		query["Properties"] = request.Properties
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateLdpsComputeGroup"),
		Version:     tea.String("2020-06-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &UpdateLdpsComputeGroupResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &UpdateLdpsComputeGroupResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @param request - UpdateLdpsComputeGroupRequest
//
// @return UpdateLdpsComputeGroupResponse
func (client *Client) UpdateLdpsComputeGroup(request *UpdateLdpsComputeGroupRequest) (_result *UpdateLdpsComputeGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateLdpsComputeGroupResponse{}
	_body, _err := client.UpdateLdpsComputeGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - UpdateLindormV2InstanceParameterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateLindormV2InstanceParameterResponse
func (client *Client) UpdateLindormV2InstanceParameterWithOptions(request *UpdateLindormV2InstanceParameterRequest, runtime *util.RuntimeOptions) (_result *UpdateLindormV2InstanceParameterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ParameterKey)) {
		query["ParameterKey"] = request.ParameterKey
	}

	if !tea.BoolValue(util.IsUnset(request.ParameterValue)) {
		query["ParameterValue"] = request.ParameterValue
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !tea.BoolValue(util.IsUnset(request.UpdateType)) {
		query["UpdateType"] = request.UpdateType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateLindormV2InstanceParameter"),
		Version:     tea.String("2020-06-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &UpdateLindormV2InstanceParameterResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &UpdateLindormV2InstanceParameterResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @param request - UpdateLindormV2InstanceParameterRequest
//
// @return UpdateLindormV2InstanceParameterResponse
func (client *Client) UpdateLindormV2InstanceParameter(request *UpdateLindormV2InstanceParameterRequest) (_result *UpdateLindormV2InstanceParameterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateLindormV2InstanceParameterResponse{}
	_body, _err := client.UpdateLindormV2InstanceParameterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Upgrades, scales up, or enable cold storage for a Lindorm instance.
//
// Description:
//
// For more information about how to select the storage type and engine type when you create a Lindorm instance, see [Select engine typpes](https://help.aliyun.com/document_detail/181971.html) and [Select storage types](https://help.aliyun.com/document_detail/174643.html).
//
// @param request - UpgradeLindormInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpgradeLindormInstanceResponse
func (client *Client) UpgradeLindormInstanceWithOptions(request *UpgradeLindormInstanceRequest, runtime *util.RuntimeOptions) (_result *UpgradeLindormInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterStorage)) {
		query["ClusterStorage"] = request.ClusterStorage
	}

	if !tea.BoolValue(util.IsUnset(request.ColdStorage)) {
		query["ColdStorage"] = request.ColdStorage
	}

	if !tea.BoolValue(util.IsUnset(request.CoreSingleStorage)) {
		query["CoreSingleStorage"] = request.CoreSingleStorage
	}

	if !tea.BoolValue(util.IsUnset(request.FilestoreNum)) {
		query["FilestoreNum"] = request.FilestoreNum
	}

	if !tea.BoolValue(util.IsUnset(request.FilestoreSpec)) {
		query["FilestoreSpec"] = request.FilestoreSpec
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.LindormNum)) {
		query["LindormNum"] = request.LindormNum
	}

	if !tea.BoolValue(util.IsUnset(request.LindormSpec)) {
		query["LindormSpec"] = request.LindormSpec
	}

	if !tea.BoolValue(util.IsUnset(request.LogNum)) {
		query["LogNum"] = request.LogNum
	}

	if !tea.BoolValue(util.IsUnset(request.LogSingleStorage)) {
		query["LogSingleStorage"] = request.LogSingleStorage
	}

	if !tea.BoolValue(util.IsUnset(request.LogSpec)) {
		query["LogSpec"] = request.LogSpec
	}

	if !tea.BoolValue(util.IsUnset(request.LtsCoreNum)) {
		query["LtsCoreNum"] = request.LtsCoreNum
	}

	if !tea.BoolValue(util.IsUnset(request.LtsCoreSpec)) {
		query["LtsCoreSpec"] = request.LtsCoreSpec
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !tea.BoolValue(util.IsUnset(request.SolrNum)) {
		query["SolrNum"] = request.SolrNum
	}

	if !tea.BoolValue(util.IsUnset(request.SolrSpec)) {
		query["SolrSpec"] = request.SolrSpec
	}

	if !tea.BoolValue(util.IsUnset(request.StreamNum)) {
		query["StreamNum"] = request.StreamNum
	}

	if !tea.BoolValue(util.IsUnset(request.StreamSpec)) {
		query["StreamSpec"] = request.StreamSpec
	}

	if !tea.BoolValue(util.IsUnset(request.TsdbNum)) {
		query["TsdbNum"] = request.TsdbNum
	}

	if !tea.BoolValue(util.IsUnset(request.TsdbSpec)) {
		query["TsdbSpec"] = request.TsdbSpec
	}

	if !tea.BoolValue(util.IsUnset(request.UpgradeType)) {
		query["UpgradeType"] = request.UpgradeType
	}

	if !tea.BoolValue(util.IsUnset(request.ZoneId)) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpgradeLindormInstance"),
		Version:     tea.String("2020-06-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &UpgradeLindormInstanceResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &UpgradeLindormInstanceResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// Upgrades, scales up, or enable cold storage for a Lindorm instance.
//
// Description:
//
// For more information about how to select the storage type and engine type when you create a Lindorm instance, see [Select engine typpes](https://help.aliyun.com/document_detail/181971.html) and [Select storage types](https://help.aliyun.com/document_detail/174643.html).
//
// @param request - UpgradeLindormInstanceRequest
//
// @return UpgradeLindormInstanceResponse
func (client *Client) UpgradeLindormInstance(request *UpgradeLindormInstanceRequest) (_result *UpgradeLindormInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpgradeLindormInstanceResponse{}
	_body, _err := client.UpgradeLindormInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - UpgradeLindormV2StreamEngineRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpgradeLindormV2StreamEngineResponse
func (client *Client) UpgradeLindormV2StreamEngineWithOptions(request *UpgradeLindormV2StreamEngineRequest, runtime *util.RuntimeOptions) (_result *UpgradeLindormV2StreamEngineResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.Quantity)) {
		query["Quantity"] = request.Quantity
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupName)) {
		query["ResourceGroupName"] = request.ResourceGroupName
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityToken)) {
		query["SecurityToken"] = request.SecurityToken
	}

	if !tea.BoolValue(util.IsUnset(request.Spec)) {
		query["Spec"] = request.Spec
	}

	if !tea.BoolValue(util.IsUnset(request.SpecId)) {
		query["SpecId"] = request.SpecId
	}

	if !tea.BoolValue(util.IsUnset(request.UpgradeType)) {
		query["UpgradeType"] = request.UpgradeType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpgradeLindormV2StreamEngine"),
		Version:     tea.String("2020-06-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &UpgradeLindormV2StreamEngineResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &UpgradeLindormV2StreamEngineResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// @param request - UpgradeLindormV2StreamEngineRequest
//
// @return UpgradeLindormV2StreamEngineResponse
func (client *Client) UpgradeLindormV2StreamEngine(request *UpgradeLindormV2StreamEngineRequest) (_result *UpgradeLindormV2StreamEngineResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpgradeLindormV2StreamEngineResponse{}
	_body, _err := client.UpgradeLindormV2StreamEngineWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
