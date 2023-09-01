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

type ActiveConfigRulesRequest struct {
	// The IDs of the rules. Separate multiple rule IDs with commas (,). You can specify a maximum of 20 rule IDs at a time.
	ConfigRuleIds *string `json:"ConfigRuleIds,omitempty" xml:"ConfigRuleIds,omitempty"`
}

func (s ActiveConfigRulesRequest) String() string {
	return tea.Prettify(s)
}

func (s ActiveConfigRulesRequest) GoString() string {
	return s.String()
}

func (s *ActiveConfigRulesRequest) SetConfigRuleIds(v string) *ActiveConfigRulesRequest {
	s.ConfigRuleIds = &v
	return s
}

type ActiveConfigRulesResponseBody struct {
	// The result information about the operation.
	OperateRuleResult *ActiveConfigRulesResponseBodyOperateRuleResult `json:"OperateRuleResult,omitempty" xml:"OperateRuleResult,omitempty" type:"Struct"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ActiveConfigRulesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ActiveConfigRulesResponseBody) GoString() string {
	return s.String()
}

func (s *ActiveConfigRulesResponseBody) SetOperateRuleResult(v *ActiveConfigRulesResponseBodyOperateRuleResult) *ActiveConfigRulesResponseBody {
	s.OperateRuleResult = v
	return s
}

func (s *ActiveConfigRulesResponseBody) SetRequestId(v string) *ActiveConfigRulesResponseBody {
	s.RequestId = &v
	return s
}

type ActiveConfigRulesResponseBodyOperateRuleResult struct {
	// The result information about each rule.
	OperateRuleItemList []*ActiveConfigRulesResponseBodyOperateRuleResultOperateRuleItemList `json:"OperateRuleItemList,omitempty" xml:"OperateRuleItemList,omitempty" type:"Repeated"`
}

func (s ActiveConfigRulesResponseBodyOperateRuleResult) String() string {
	return tea.Prettify(s)
}

func (s ActiveConfigRulesResponseBodyOperateRuleResult) GoString() string {
	return s.String()
}

func (s *ActiveConfigRulesResponseBodyOperateRuleResult) SetOperateRuleItemList(v []*ActiveConfigRulesResponseBodyOperateRuleResultOperateRuleItemList) *ActiveConfigRulesResponseBodyOperateRuleResult {
	s.OperateRuleItemList = v
	return s
}

type ActiveConfigRulesResponseBodyOperateRuleResultOperateRuleItemList struct {
	// The ID of the rule.
	ConfigRuleId *string `json:"ConfigRuleId,omitempty" xml:"ConfigRuleId,omitempty"`
	// The error code.
	//
	// *   If the rule is enabled, no error code is returned.
	// *   If you fail to enable the rule, an error code is returned. For more information, visit the [API Error Center](https://error-center.alibabacloud.com/status/product/Config).
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// Indicates whether the rule is enabled. Valid values:
	//
	// *   true
	// *   false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ActiveConfigRulesResponseBodyOperateRuleResultOperateRuleItemList) String() string {
	return tea.Prettify(s)
}

func (s ActiveConfigRulesResponseBodyOperateRuleResultOperateRuleItemList) GoString() string {
	return s.String()
}

func (s *ActiveConfigRulesResponseBodyOperateRuleResultOperateRuleItemList) SetConfigRuleId(v string) *ActiveConfigRulesResponseBodyOperateRuleResultOperateRuleItemList {
	s.ConfigRuleId = &v
	return s
}

func (s *ActiveConfigRulesResponseBodyOperateRuleResultOperateRuleItemList) SetErrorCode(v string) *ActiveConfigRulesResponseBodyOperateRuleResultOperateRuleItemList {
	s.ErrorCode = &v
	return s
}

func (s *ActiveConfigRulesResponseBodyOperateRuleResultOperateRuleItemList) SetSuccess(v bool) *ActiveConfigRulesResponseBodyOperateRuleResultOperateRuleItemList {
	s.Success = &v
	return s
}

type ActiveConfigRulesResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ActiveConfigRulesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ActiveConfigRulesResponse) String() string {
	return tea.Prettify(s)
}

func (s ActiveConfigRulesResponse) GoString() string {
	return s.String()
}

func (s *ActiveConfigRulesResponse) SetHeaders(v map[string]*string) *ActiveConfigRulesResponse {
	s.Headers = v
	return s
}

func (s *ActiveConfigRulesResponse) SetStatusCode(v int32) *ActiveConfigRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *ActiveConfigRulesResponse) SetBody(v *ActiveConfigRulesResponseBody) *ActiveConfigRulesResponse {
	s.Body = v
	return s
}

type DeleteConfigRulesRequest struct {
	// The IDs of the rules. Separate multiple rule IDs with commas (,). You can specify a maximum of 20 rule IDs at a time.
	ConfigRuleIds *string `json:"ConfigRuleIds,omitempty" xml:"ConfigRuleIds,omitempty"`
}

func (s DeleteConfigRulesRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteConfigRulesRequest) GoString() string {
	return s.String()
}

func (s *DeleteConfigRulesRequest) SetConfigRuleIds(v string) *DeleteConfigRulesRequest {
	s.ConfigRuleIds = &v
	return s
}

type DeleteConfigRulesResponseBody struct {
	// The result information about the operation.
	OperateRuleResult *DeleteConfigRulesResponseBodyOperateRuleResult `json:"OperateRuleResult,omitempty" xml:"OperateRuleResult,omitempty" type:"Struct"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteConfigRulesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteConfigRulesResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteConfigRulesResponseBody) SetOperateRuleResult(v *DeleteConfigRulesResponseBodyOperateRuleResult) *DeleteConfigRulesResponseBody {
	s.OperateRuleResult = v
	return s
}

func (s *DeleteConfigRulesResponseBody) SetRequestId(v string) *DeleteConfigRulesResponseBody {
	s.RequestId = &v
	return s
}

type DeleteConfigRulesResponseBodyOperateRuleResult struct {
	// The result information about each rule.
	OperateRuleItemList []*DeleteConfigRulesResponseBodyOperateRuleResultOperateRuleItemList `json:"OperateRuleItemList,omitempty" xml:"OperateRuleItemList,omitempty" type:"Repeated"`
}

func (s DeleteConfigRulesResponseBodyOperateRuleResult) String() string {
	return tea.Prettify(s)
}

func (s DeleteConfigRulesResponseBodyOperateRuleResult) GoString() string {
	return s.String()
}

func (s *DeleteConfigRulesResponseBodyOperateRuleResult) SetOperateRuleItemList(v []*DeleteConfigRulesResponseBodyOperateRuleResultOperateRuleItemList) *DeleteConfigRulesResponseBodyOperateRuleResult {
	s.OperateRuleItemList = v
	return s
}

type DeleteConfigRulesResponseBodyOperateRuleResultOperateRuleItemList struct {
	// The ID of the rule.
	ConfigRuleId *string `json:"ConfigRuleId,omitempty" xml:"ConfigRuleId,omitempty"`
	// The error code.
	//
	// *   If the rule is deleted, no error code is returned.
	// *   If you fail to delete the rule, an error code is returned. For more information, visit the [API Error Center](https://error-center.alibabacloud.com/status/product/Config).
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// Indicates whether the rule is deleted. Valid values:
	//
	// *   true
	// *   false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteConfigRulesResponseBodyOperateRuleResultOperateRuleItemList) String() string {
	return tea.Prettify(s)
}

func (s DeleteConfigRulesResponseBodyOperateRuleResultOperateRuleItemList) GoString() string {
	return s.String()
}

func (s *DeleteConfigRulesResponseBodyOperateRuleResultOperateRuleItemList) SetConfigRuleId(v string) *DeleteConfigRulesResponseBodyOperateRuleResultOperateRuleItemList {
	s.ConfigRuleId = &v
	return s
}

func (s *DeleteConfigRulesResponseBodyOperateRuleResultOperateRuleItemList) SetErrorCode(v string) *DeleteConfigRulesResponseBodyOperateRuleResultOperateRuleItemList {
	s.ErrorCode = &v
	return s
}

func (s *DeleteConfigRulesResponseBodyOperateRuleResultOperateRuleItemList) SetSuccess(v bool) *DeleteConfigRulesResponseBodyOperateRuleResultOperateRuleItemList {
	s.Success = &v
	return s
}

type DeleteConfigRulesResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteConfigRulesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteConfigRulesResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteConfigRulesResponse) GoString() string {
	return s.String()
}

func (s *DeleteConfigRulesResponse) SetHeaders(v map[string]*string) *DeleteConfigRulesResponse {
	s.Headers = v
	return s
}

func (s *DeleteConfigRulesResponse) SetStatusCode(v int32) *DeleteConfigRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteConfigRulesResponse) SetBody(v *DeleteConfigRulesResponseBody) *DeleteConfigRulesResponse {
	s.Body = v
	return s
}

type DescribeComplianceRequest struct {
	// The compliance evaluation result of the resource. Valid values:
	//
	// *   COMPLIANT: The resource is evaluated to be compliant.
	// *   NON_COMPLIANT: The resource is evaluated to be non-compliant.
	// *   NOT_APPLICABLE: The rule does not apply to the resource.
	// *   INSUFFICIENT_DATA: The resource data is insufficient.
	ComplianceType *string `json:"ComplianceType,omitempty" xml:"ComplianceType,omitempty"`
	// The ID of the rule.
	//
	// If you query compliance evaluation results by resource, you must specify the ConfigRuleId, ResourceType, and ResourceId parameters. Otherwise, the ConfigRuleId parameter becomes invalid.
	ConfigRuleId *string `json:"ConfigRuleId,omitempty" xml:"ConfigRuleId,omitempty"`
	// This parameter is scheduled to be removed before 00:00:00, June 30, 2021. Account group-related APIs will be provided as an alternative before 00:00:00, May 30, 2021. If you are using this parameter, we recommend that you switch to account group-related APIs after 00:00:00, May 30, 2021. For more information, see [Account groups](~~211534~~).
	MemberId *int64 `json:"MemberId,omitempty" xml:"MemberId,omitempty"`
	// This parameter is scheduled to be removed before 00:00:00, June 30, 2021. Account group-related APIs will be provided as an alternative before 00:00:00, May 30, 2021. If you are using this parameter, we recommend that you switch to account group-related APIs after 00:00:00, May 30, 2021. For more information, see [Account groups](~~211534~~).
	MultiAccount *bool `json:"MultiAccount,omitempty" xml:"MultiAccount,omitempty"`
	// The ID of the resource.
	//
	// If you query compliance evaluation results by resource, you must specify the ResourceType and ResourceId parameters.
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The type of the resource.
	//
	// If you query compliance evaluation results by resource, you must specify the ResourceType and ResourceId parameters.
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s DescribeComplianceRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeComplianceRequest) GoString() string {
	return s.String()
}

func (s *DescribeComplianceRequest) SetComplianceType(v string) *DescribeComplianceRequest {
	s.ComplianceType = &v
	return s
}

func (s *DescribeComplianceRequest) SetConfigRuleId(v string) *DescribeComplianceRequest {
	s.ConfigRuleId = &v
	return s
}

func (s *DescribeComplianceRequest) SetMemberId(v int64) *DescribeComplianceRequest {
	s.MemberId = &v
	return s
}

func (s *DescribeComplianceRequest) SetMultiAccount(v bool) *DescribeComplianceRequest {
	s.MultiAccount = &v
	return s
}

func (s *DescribeComplianceRequest) SetResourceId(v string) *DescribeComplianceRequest {
	s.ResourceId = &v
	return s
}

func (s *DescribeComplianceRequest) SetResourceType(v string) *DescribeComplianceRequest {
	s.ResourceType = &v
	return s
}

type DescribeComplianceResponseBody struct {
	// The statistics of compliance evaluations.
	ComplianceResult *DescribeComplianceResponseBodyComplianceResult `json:"ComplianceResult,omitempty" xml:"ComplianceResult,omitempty" type:"Struct"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeComplianceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeComplianceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeComplianceResponseBody) SetComplianceResult(v *DescribeComplianceResponseBodyComplianceResult) *DescribeComplianceResponseBody {
	s.ComplianceResult = v
	return s
}

func (s *DescribeComplianceResponseBody) SetRequestId(v string) *DescribeComplianceResponseBody {
	s.RequestId = &v
	return s
}

type DescribeComplianceResponseBodyComplianceResult struct {
	// The compliance evaluation results based on compliance types.
	Compliances []*DescribeComplianceResponseBodyComplianceResultCompliances `json:"Compliances,omitempty" xml:"Compliances,omitempty" type:"Repeated"`
	// The total number of compliance evaluations.
	//
	// *   This parameter returns the total number of evaluated resources if you set the ResourceId parameter in the request.
	// *   This parameter returns the total number of triggered rules if you set the ConfigRuleId parameter in the request.
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeComplianceResponseBodyComplianceResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeComplianceResponseBodyComplianceResult) GoString() string {
	return s.String()
}

func (s *DescribeComplianceResponseBodyComplianceResult) SetCompliances(v []*DescribeComplianceResponseBodyComplianceResultCompliances) *DescribeComplianceResponseBodyComplianceResult {
	s.Compliances = v
	return s
}

func (s *DescribeComplianceResponseBodyComplianceResult) SetTotalCount(v int64) *DescribeComplianceResponseBodyComplianceResult {
	s.TotalCount = &v
	return s
}

type DescribeComplianceResponseBodyComplianceResultCompliances struct {
	// The compliance evaluation result of the resource. Valid values:
	//
	// *   COMPLIANT: The resource is evaluated to be compliant.
	// *   NON_COMPLIANT: The resource is evaluated to be non-compliant.
	// *   NOT_APPLICABLE: The rule does not apply to the resource.
	// *   INSUFFICIENT_DATA: The resource data is insufficient.
	ComplianceType *string `json:"ComplianceType,omitempty" xml:"ComplianceType,omitempty"`
	// The number of compliance evaluations with the corresponding result.
	//
	// *   This parameter returns the total number of evaluated resources if you set the ResourceId parameter in the request.
	// *   This parameter returns the total number of triggered rules if you set the ConfigRuleId parameter in the request.
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
}

func (s DescribeComplianceResponseBodyComplianceResultCompliances) String() string {
	return tea.Prettify(s)
}

func (s DescribeComplianceResponseBodyComplianceResultCompliances) GoString() string {
	return s.String()
}

func (s *DescribeComplianceResponseBodyComplianceResultCompliances) SetComplianceType(v string) *DescribeComplianceResponseBodyComplianceResultCompliances {
	s.ComplianceType = &v
	return s
}

func (s *DescribeComplianceResponseBodyComplianceResultCompliances) SetCount(v int32) *DescribeComplianceResponseBodyComplianceResultCompliances {
	s.Count = &v
	return s
}

type DescribeComplianceResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeComplianceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeComplianceResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeComplianceResponse) GoString() string {
	return s.String()
}

func (s *DescribeComplianceResponse) SetHeaders(v map[string]*string) *DescribeComplianceResponse {
	s.Headers = v
	return s
}

func (s *DescribeComplianceResponse) SetStatusCode(v int32) *DescribeComplianceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeComplianceResponse) SetBody(v *DescribeComplianceResponseBody) *DescribeComplianceResponse {
	s.Body = v
	return s
}

type DescribeComplianceSummaryRequest struct {
	// This parameter is scheduled to be removed before 00:00:00, June 30, 2021. Account group-related APIs will be provided as an alternative before 00:00:00, May 30, 2021. If you are using this parameter, we recommend that you switch to account group-related APIs after 00:00:00, May 30, 2021. For more information, see [Account groups](~~211534~~).
	MemberId *int64 `json:"MemberId,omitempty" xml:"MemberId,omitempty"`
	// This parameter is scheduled to be removed before 00:00:00, June 30, 2021. Account group-related APIs will be provided as an alternative before 00:00:00, May 30, 2021. If you are using this parameter, we recommend that you switch to account group-related APIs after 00:00:00, May 30, 2021. For more information, see [Account groups](~~211534~~).
	MultiAccount *bool `json:"MultiAccount,omitempty" xml:"MultiAccount,omitempty"`
}

func (s DescribeComplianceSummaryRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeComplianceSummaryRequest) GoString() string {
	return s.String()
}

func (s *DescribeComplianceSummaryRequest) SetMemberId(v int64) *DescribeComplianceSummaryRequest {
	s.MemberId = &v
	return s
}

func (s *DescribeComplianceSummaryRequest) SetMultiAccount(v bool) *DescribeComplianceSummaryRequest {
	s.MultiAccount = &v
	return s
}

type DescribeComplianceSummaryResponseBody struct {
	// The summary of compliance evaluations.
	ComplianceSummary *DescribeComplianceSummaryResponseBodyComplianceSummary `json:"ComplianceSummary,omitempty" xml:"ComplianceSummary,omitempty" type:"Struct"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeComplianceSummaryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeComplianceSummaryResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeComplianceSummaryResponseBody) SetComplianceSummary(v *DescribeComplianceSummaryResponseBodyComplianceSummary) *DescribeComplianceSummaryResponseBody {
	s.ComplianceSummary = v
	return s
}

func (s *DescribeComplianceSummaryResponseBody) SetRequestId(v string) *DescribeComplianceSummaryResponseBody {
	s.RequestId = &v
	return s
}

type DescribeComplianceSummaryResponseBodyComplianceSummary struct {
	// The summary of compliance evaluations from the rule dimension.
	ComplianceSummaryByConfigRule *DescribeComplianceSummaryResponseBodyComplianceSummaryComplianceSummaryByConfigRule `json:"ComplianceSummaryByConfigRule,omitempty" xml:"ComplianceSummaryByConfigRule,omitempty" type:"Struct"`
	// The summary of compliance evaluations from the resource dimension.
	ComplianceSummaryByResource *DescribeComplianceSummaryResponseBodyComplianceSummaryComplianceSummaryByResource `json:"ComplianceSummaryByResource,omitempty" xml:"ComplianceSummaryByResource,omitempty" type:"Struct"`
}

func (s DescribeComplianceSummaryResponseBodyComplianceSummary) String() string {
	return tea.Prettify(s)
}

func (s DescribeComplianceSummaryResponseBodyComplianceSummary) GoString() string {
	return s.String()
}

func (s *DescribeComplianceSummaryResponseBodyComplianceSummary) SetComplianceSummaryByConfigRule(v *DescribeComplianceSummaryResponseBodyComplianceSummaryComplianceSummaryByConfigRule) *DescribeComplianceSummaryResponseBodyComplianceSummary {
	s.ComplianceSummaryByConfigRule = v
	return s
}

func (s *DescribeComplianceSummaryResponseBodyComplianceSummary) SetComplianceSummaryByResource(v *DescribeComplianceSummaryResponseBodyComplianceSummaryComplianceSummaryByResource) *DescribeComplianceSummaryResponseBodyComplianceSummary {
	s.ComplianceSummaryByResource = v
	return s
}

type DescribeComplianceSummaryResponseBodyComplianceSummaryComplianceSummaryByConfigRule struct {
	// The timestamp when the summary was recorded.
	ComplianceSummaryTimestamp *int64 `json:"ComplianceSummaryTimestamp,omitempty" xml:"ComplianceSummaryTimestamp,omitempty"`
	// The number of rules whose evaluation results are compliant.
	CompliantCount *int32 `json:"CompliantCount,omitempty" xml:"CompliantCount,omitempty"`
	// The number of rules whose evaluation results are non-compliant.
	NonCompliantCount *int32 `json:"NonCompliantCount,omitempty" xml:"NonCompliantCount,omitempty"`
	// The total number of rules.
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeComplianceSummaryResponseBodyComplianceSummaryComplianceSummaryByConfigRule) String() string {
	return tea.Prettify(s)
}

func (s DescribeComplianceSummaryResponseBodyComplianceSummaryComplianceSummaryByConfigRule) GoString() string {
	return s.String()
}

func (s *DescribeComplianceSummaryResponseBodyComplianceSummaryComplianceSummaryByConfigRule) SetComplianceSummaryTimestamp(v int64) *DescribeComplianceSummaryResponseBodyComplianceSummaryComplianceSummaryByConfigRule {
	s.ComplianceSummaryTimestamp = &v
	return s
}

func (s *DescribeComplianceSummaryResponseBodyComplianceSummaryComplianceSummaryByConfigRule) SetCompliantCount(v int32) *DescribeComplianceSummaryResponseBodyComplianceSummaryComplianceSummaryByConfigRule {
	s.CompliantCount = &v
	return s
}

func (s *DescribeComplianceSummaryResponseBodyComplianceSummaryComplianceSummaryByConfigRule) SetNonCompliantCount(v int32) *DescribeComplianceSummaryResponseBodyComplianceSummaryComplianceSummaryByConfigRule {
	s.NonCompliantCount = &v
	return s
}

func (s *DescribeComplianceSummaryResponseBodyComplianceSummaryComplianceSummaryByConfigRule) SetTotalCount(v int64) *DescribeComplianceSummaryResponseBodyComplianceSummaryComplianceSummaryByConfigRule {
	s.TotalCount = &v
	return s
}

type DescribeComplianceSummaryResponseBodyComplianceSummaryComplianceSummaryByResource struct {
	// The timestamp when the summary was recorded.
	ComplianceSummaryTimestamp *int64 `json:"ComplianceSummaryTimestamp,omitempty" xml:"ComplianceSummaryTimestamp,omitempty"`
	// The number of resources evaluated as compliant.
	CompliantCount *int32 `json:"CompliantCount,omitempty" xml:"CompliantCount,omitempty"`
	// The number of resources evaluated as non-compliant.
	NonCompliantCount *int32 `json:"NonCompliantCount,omitempty" xml:"NonCompliantCount,omitempty"`
	// The total number of resources.
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeComplianceSummaryResponseBodyComplianceSummaryComplianceSummaryByResource) String() string {
	return tea.Prettify(s)
}

func (s DescribeComplianceSummaryResponseBodyComplianceSummaryComplianceSummaryByResource) GoString() string {
	return s.String()
}

func (s *DescribeComplianceSummaryResponseBodyComplianceSummaryComplianceSummaryByResource) SetComplianceSummaryTimestamp(v int64) *DescribeComplianceSummaryResponseBodyComplianceSummaryComplianceSummaryByResource {
	s.ComplianceSummaryTimestamp = &v
	return s
}

func (s *DescribeComplianceSummaryResponseBodyComplianceSummaryComplianceSummaryByResource) SetCompliantCount(v int32) *DescribeComplianceSummaryResponseBodyComplianceSummaryComplianceSummaryByResource {
	s.CompliantCount = &v
	return s
}

func (s *DescribeComplianceSummaryResponseBodyComplianceSummaryComplianceSummaryByResource) SetNonCompliantCount(v int32) *DescribeComplianceSummaryResponseBodyComplianceSummaryComplianceSummaryByResource {
	s.NonCompliantCount = &v
	return s
}

func (s *DescribeComplianceSummaryResponseBodyComplianceSummaryComplianceSummaryByResource) SetTotalCount(v int64) *DescribeComplianceSummaryResponseBodyComplianceSummaryComplianceSummaryByResource {
	s.TotalCount = &v
	return s
}

type DescribeComplianceSummaryResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeComplianceSummaryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeComplianceSummaryResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeComplianceSummaryResponse) GoString() string {
	return s.String()
}

func (s *DescribeComplianceSummaryResponse) SetHeaders(v map[string]*string) *DescribeComplianceSummaryResponse {
	s.Headers = v
	return s
}

func (s *DescribeComplianceSummaryResponse) SetStatusCode(v int32) *DescribeComplianceSummaryResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeComplianceSummaryResponse) SetBody(v *DescribeComplianceSummaryResponseBody) *DescribeComplianceSummaryResponse {
	s.Body = v
	return s
}

type DescribeConfigRuleRequest struct {
	// The ID of the rule. You can call the [ListConfigRules](~~169607~~) operation to obtain the rule ID.
	ConfigRuleId *string `json:"ConfigRuleId,omitempty" xml:"ConfigRuleId,omitempty"`
	// This parameter is unavailable since 00:00:00 on June 30, 2021. Account group-related operations are provided as an alternative after 00:00:00 on May 30, 2021. If you are using this parameter, we recommend that you switch to account group-related operations after 00:00:00 on May 30, 2021. For information about account groups, see [Overview](~~211534~~).
	MemberId *int64 `json:"MemberId,omitempty" xml:"MemberId,omitempty"`
	// This parameter is unavailable since 00:00:00 on June 30, 2021. Account group-related operations are provided as an alternative after 00:00:00 on May 30, 2021. If you are using this parameter, we recommend that you switch to account group-related operations after 00:00:00 on May 30, 2021. For information about account groups, see [Overview](~~211534~~).
	MultiAccount *bool `json:"MultiAccount,omitempty" xml:"MultiAccount,omitempty"`
}

func (s DescribeConfigRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeConfigRuleRequest) GoString() string {
	return s.String()
}

func (s *DescribeConfigRuleRequest) SetConfigRuleId(v string) *DescribeConfigRuleRequest {
	s.ConfigRuleId = &v
	return s
}

func (s *DescribeConfigRuleRequest) SetMemberId(v int64) *DescribeConfigRuleRequest {
	s.MemberId = &v
	return s
}

func (s *DescribeConfigRuleRequest) SetMultiAccount(v bool) *DescribeConfigRuleRequest {
	s.MultiAccount = &v
	return s
}

type DescribeConfigRuleResponseBody struct {
	// The details of the rule.
	ConfigRule *DescribeConfigRuleResponseBodyConfigRule `json:"ConfigRule,omitempty" xml:"ConfigRule,omitempty" type:"Struct"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeConfigRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeConfigRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeConfigRuleResponseBody) SetConfigRule(v *DescribeConfigRuleResponseBodyConfigRule) *DescribeConfigRuleResponseBody {
	s.ConfigRule = v
	return s
}

func (s *DescribeConfigRuleResponseBody) SetRequestId(v string) *DescribeConfigRuleResponseBody {
	s.RequestId = &v
	return s
}

type DescribeConfigRuleResponseBodyConfigRule struct {
	// The Alibaba Cloud Resource Name (ARN) of the rule.
	ConfigRuleArn *string `json:"ConfigRuleArn,omitempty" xml:"ConfigRuleArn,omitempty"`
	// The information about compliance evaluations performed by the rule.
	ConfigRuleEvaluationStatus *DescribeConfigRuleResponseBodyConfigRuleConfigRuleEvaluationStatus `json:"ConfigRuleEvaluationStatus,omitempty" xml:"ConfigRuleEvaluationStatus,omitempty" type:"Struct"`
	// The ID of the rule.
	ConfigRuleId *string `json:"ConfigRuleId,omitempty" xml:"ConfigRuleId,omitempty"`
	// The name of the evaluation rule.
	ConfigRuleName *string `json:"ConfigRuleName,omitempty" xml:"ConfigRuleName,omitempty"`
	// The status of the rule. Valid values:
	//
	// *   ACTIVE: The rule is enabled.
	// *   EVALUATING: The rule was triggered and is being used to monitor resource configurations.
	// *   INACTIVE: The rule was disabled.
	ConfigRuleState *string `json:"ConfigRuleState,omitempty" xml:"ConfigRuleState,omitempty"`
	// The timestamp when the rule was created.
	CreateTimestamp *int64 `json:"CreateTimestamp,omitempty" xml:"CreateTimestamp,omitempty"`
	// The description of the system defense rule.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The input parameters of the rule.
	InputParameters map[string]interface{} `json:"InputParameters,omitempty" xml:"InputParameters,omitempty"`
	// The details of the managed rule.
	ManagedRule *DescribeConfigRuleResponseBodyConfigRuleManagedRule `json:"ManagedRule,omitempty" xml:"ManagedRule,omitempty" type:"Struct"`
	// The intervals at which the rule is triggered. Valid values:
	//
	// *   One_Hour: 1 hour
	// *   Three_Hours: 3 hours
	// *   Six_Hours: 6 hours
	// *   Twelve_Hours: 12 hours.
	// *   TwentyFour_Hours: 24 hours.
	MaximumExecutionFrequency *string `json:"MaximumExecutionFrequency,omitempty" xml:"MaximumExecutionFrequency,omitempty"`
	// The timestamp when the rule was last modified.
	ModifiedTimestamp *int64 `json:"ModifiedTimestamp,omitempty" xml:"ModifiedTimestamp,omitempty"`
	// The risk level of the resources that are not compliant with the rule. Valid values:
	//
	// *   1: high
	// *   2: medium
	// *   3: low
	RiskLevel *int32 `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
	// The monitoring scope of the rule.
	Scope *DescribeConfigRuleResponseBodyConfigRuleScope `json:"Scope,omitempty" xml:"Scope,omitempty" type:"Struct"`
	// The information about the trigger of the rule.
	Source *DescribeConfigRuleResponseBodyConfigRuleSource `json:"Source,omitempty" xml:"Source,omitempty" type:"Struct"`
}

func (s DescribeConfigRuleResponseBodyConfigRule) String() string {
	return tea.Prettify(s)
}

func (s DescribeConfigRuleResponseBodyConfigRule) GoString() string {
	return s.String()
}

func (s *DescribeConfigRuleResponseBodyConfigRule) SetConfigRuleArn(v string) *DescribeConfigRuleResponseBodyConfigRule {
	s.ConfigRuleArn = &v
	return s
}

func (s *DescribeConfigRuleResponseBodyConfigRule) SetConfigRuleEvaluationStatus(v *DescribeConfigRuleResponseBodyConfigRuleConfigRuleEvaluationStatus) *DescribeConfigRuleResponseBodyConfigRule {
	s.ConfigRuleEvaluationStatus = v
	return s
}

func (s *DescribeConfigRuleResponseBodyConfigRule) SetConfigRuleId(v string) *DescribeConfigRuleResponseBodyConfigRule {
	s.ConfigRuleId = &v
	return s
}

func (s *DescribeConfigRuleResponseBodyConfigRule) SetConfigRuleName(v string) *DescribeConfigRuleResponseBodyConfigRule {
	s.ConfigRuleName = &v
	return s
}

func (s *DescribeConfigRuleResponseBodyConfigRule) SetConfigRuleState(v string) *DescribeConfigRuleResponseBodyConfigRule {
	s.ConfigRuleState = &v
	return s
}

func (s *DescribeConfigRuleResponseBodyConfigRule) SetCreateTimestamp(v int64) *DescribeConfigRuleResponseBodyConfigRule {
	s.CreateTimestamp = &v
	return s
}

func (s *DescribeConfigRuleResponseBodyConfigRule) SetDescription(v string) *DescribeConfigRuleResponseBodyConfigRule {
	s.Description = &v
	return s
}

func (s *DescribeConfigRuleResponseBodyConfigRule) SetInputParameters(v map[string]interface{}) *DescribeConfigRuleResponseBodyConfigRule {
	s.InputParameters = v
	return s
}

func (s *DescribeConfigRuleResponseBodyConfigRule) SetManagedRule(v *DescribeConfigRuleResponseBodyConfigRuleManagedRule) *DescribeConfigRuleResponseBodyConfigRule {
	s.ManagedRule = v
	return s
}

func (s *DescribeConfigRuleResponseBodyConfigRule) SetMaximumExecutionFrequency(v string) *DescribeConfigRuleResponseBodyConfigRule {
	s.MaximumExecutionFrequency = &v
	return s
}

func (s *DescribeConfigRuleResponseBodyConfigRule) SetModifiedTimestamp(v int64) *DescribeConfigRuleResponseBodyConfigRule {
	s.ModifiedTimestamp = &v
	return s
}

func (s *DescribeConfigRuleResponseBodyConfigRule) SetRiskLevel(v int32) *DescribeConfigRuleResponseBodyConfigRule {
	s.RiskLevel = &v
	return s
}

func (s *DescribeConfigRuleResponseBodyConfigRule) SetScope(v *DescribeConfigRuleResponseBodyConfigRuleScope) *DescribeConfigRuleResponseBodyConfigRule {
	s.Scope = v
	return s
}

func (s *DescribeConfigRuleResponseBodyConfigRule) SetSource(v *DescribeConfigRuleResponseBodyConfigRuleSource) *DescribeConfigRuleResponseBodyConfigRule {
	s.Source = v
	return s
}

type DescribeConfigRuleResponseBodyConfigRuleConfigRuleEvaluationStatus struct {
	// The timestamp when the rule was first triggered.
	FirstActivatedTimestamp *int64 `json:"FirstActivatedTimestamp,omitempty" xml:"FirstActivatedTimestamp,omitempty"`
	// Indicates whether resources were evaluated based on the rule. Valid values:
	//
	// *   true
	// *   false
	FirstEvaluationStarted *bool `json:"FirstEvaluationStarted,omitempty" xml:"FirstEvaluationStarted,omitempty"`
	// The error code returned for the last failed compliance evaluation.
	LastErrorCode *string `json:"LastErrorCode,omitempty" xml:"LastErrorCode,omitempty"`
	// The error message returned for the last failed compliance evaluation.
	LastErrorMessage *string `json:"LastErrorMessage,omitempty" xml:"LastErrorMessage,omitempty"`
	// The timestamp when the last failed compliance evaluation of the rule ended.
	LastFailedEvaluationTimestamp *int64 `json:"LastFailedEvaluationTimestamp,omitempty" xml:"LastFailedEvaluationTimestamp,omitempty"`
	// The timestamp when the last failed compliance evaluation of the rule started.
	LastFailedInvocationTimestamp *int64 `json:"LastFailedInvocationTimestamp,omitempty" xml:"LastFailedInvocationTimestamp,omitempty"`
	// The timestamp when the last successful compliance evaluation of the rule ended.
	LastSuccessfulEvaluationTimestamp *int64 `json:"LastSuccessfulEvaluationTimestamp,omitempty" xml:"LastSuccessfulEvaluationTimestamp,omitempty"`
	// The timestamp when the last successful compliance evaluation of the rule started.
	LastSuccessfulInvocationTimestamp *int64 `json:"LastSuccessfulInvocationTimestamp,omitempty" xml:"LastSuccessfulInvocationTimestamp,omitempty"`
}

func (s DescribeConfigRuleResponseBodyConfigRuleConfigRuleEvaluationStatus) String() string {
	return tea.Prettify(s)
}

func (s DescribeConfigRuleResponseBodyConfigRuleConfigRuleEvaluationStatus) GoString() string {
	return s.String()
}

func (s *DescribeConfigRuleResponseBodyConfigRuleConfigRuleEvaluationStatus) SetFirstActivatedTimestamp(v int64) *DescribeConfigRuleResponseBodyConfigRuleConfigRuleEvaluationStatus {
	s.FirstActivatedTimestamp = &v
	return s
}

func (s *DescribeConfigRuleResponseBodyConfigRuleConfigRuleEvaluationStatus) SetFirstEvaluationStarted(v bool) *DescribeConfigRuleResponseBodyConfigRuleConfigRuleEvaluationStatus {
	s.FirstEvaluationStarted = &v
	return s
}

func (s *DescribeConfigRuleResponseBodyConfigRuleConfigRuleEvaluationStatus) SetLastErrorCode(v string) *DescribeConfigRuleResponseBodyConfigRuleConfigRuleEvaluationStatus {
	s.LastErrorCode = &v
	return s
}

func (s *DescribeConfigRuleResponseBodyConfigRuleConfigRuleEvaluationStatus) SetLastErrorMessage(v string) *DescribeConfigRuleResponseBodyConfigRuleConfigRuleEvaluationStatus {
	s.LastErrorMessage = &v
	return s
}

func (s *DescribeConfigRuleResponseBodyConfigRuleConfigRuleEvaluationStatus) SetLastFailedEvaluationTimestamp(v int64) *DescribeConfigRuleResponseBodyConfigRuleConfigRuleEvaluationStatus {
	s.LastFailedEvaluationTimestamp = &v
	return s
}

func (s *DescribeConfigRuleResponseBodyConfigRuleConfigRuleEvaluationStatus) SetLastFailedInvocationTimestamp(v int64) *DescribeConfigRuleResponseBodyConfigRuleConfigRuleEvaluationStatus {
	s.LastFailedInvocationTimestamp = &v
	return s
}

func (s *DescribeConfigRuleResponseBodyConfigRuleConfigRuleEvaluationStatus) SetLastSuccessfulEvaluationTimestamp(v int64) *DescribeConfigRuleResponseBodyConfigRuleConfigRuleEvaluationStatus {
	s.LastSuccessfulEvaluationTimestamp = &v
	return s
}

func (s *DescribeConfigRuleResponseBodyConfigRuleConfigRuleEvaluationStatus) SetLastSuccessfulInvocationTimestamp(v int64) *DescribeConfigRuleResponseBodyConfigRuleConfigRuleEvaluationStatus {
	s.LastSuccessfulInvocationTimestamp = &v
	return s
}

type DescribeConfigRuleResponseBodyConfigRuleManagedRule struct {
	// The settings of the required input parameters for the managed rule.
	CompulsoryInputParameterDetails map[string]interface{} `json:"CompulsoryInputParameterDetails,omitempty" xml:"CompulsoryInputParameterDetails,omitempty"`
	// The description of the managed rule.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The identifier of the managed rule.
	Identifier *string `json:"Identifier,omitempty" xml:"Identifier,omitempty"`
	// The labels.
	Labels []*string `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	// The name of the managed rule.
	ManagedRuleName *string `json:"ManagedRuleName,omitempty" xml:"ManagedRuleName,omitempty"`
	// The settings of the optional input parameters for the managed rule.
	OptionalInputParameterDetails map[string]interface{} `json:"OptionalInputParameterDetails,omitempty" xml:"OptionalInputParameterDetails,omitempty"`
	// The intervals at which the rule is triggered. Valid values:
	//
	// *   One_Hour: 1 hour
	// *   Three_Hours: 3 hours
	// *   Six_Hours: 6 hours
	// *   Twelve_Hours: 12 hours.
	// *   TwentyFour_Hours: 24 hours.
	SourceDetails []*DescribeConfigRuleResponseBodyConfigRuleManagedRuleSourceDetails `json:"SourceDetails,omitempty" xml:"SourceDetails,omitempty" type:"Repeated"`
}

func (s DescribeConfigRuleResponseBodyConfigRuleManagedRule) String() string {
	return tea.Prettify(s)
}

func (s DescribeConfigRuleResponseBodyConfigRuleManagedRule) GoString() string {
	return s.String()
}

func (s *DescribeConfigRuleResponseBodyConfigRuleManagedRule) SetCompulsoryInputParameterDetails(v map[string]interface{}) *DescribeConfigRuleResponseBodyConfigRuleManagedRule {
	s.CompulsoryInputParameterDetails = v
	return s
}

func (s *DescribeConfigRuleResponseBodyConfigRuleManagedRule) SetDescription(v string) *DescribeConfigRuleResponseBodyConfigRuleManagedRule {
	s.Description = &v
	return s
}

func (s *DescribeConfigRuleResponseBodyConfigRuleManagedRule) SetIdentifier(v string) *DescribeConfigRuleResponseBodyConfigRuleManagedRule {
	s.Identifier = &v
	return s
}

func (s *DescribeConfigRuleResponseBodyConfigRuleManagedRule) SetLabels(v []*string) *DescribeConfigRuleResponseBodyConfigRuleManagedRule {
	s.Labels = v
	return s
}

func (s *DescribeConfigRuleResponseBodyConfigRuleManagedRule) SetManagedRuleName(v string) *DescribeConfigRuleResponseBodyConfigRuleManagedRule {
	s.ManagedRuleName = &v
	return s
}

func (s *DescribeConfigRuleResponseBodyConfigRuleManagedRule) SetOptionalInputParameterDetails(v map[string]interface{}) *DescribeConfigRuleResponseBodyConfigRuleManagedRule {
	s.OptionalInputParameterDetails = v
	return s
}

func (s *DescribeConfigRuleResponseBodyConfigRuleManagedRule) SetSourceDetails(v []*DescribeConfigRuleResponseBodyConfigRuleManagedRuleSourceDetails) *DescribeConfigRuleResponseBodyConfigRuleManagedRule {
	s.SourceDetails = v
	return s
}

type DescribeConfigRuleResponseBodyConfigRuleManagedRuleSourceDetails struct {
	// The event source of the managed rule.
	//
	// >  Only events related to Cloud Config are supported. The value is fixed to aliyun.config.
	EventSource *string `json:"EventSource,omitempty" xml:"EventSource,omitempty"`
	// The intervals at which the managed rule is triggered. Valid values:
	//
	// *   One_Hour: 1 hour.
	// *   Three_Hours: 3 hours.
	// *   Six_Hours: 6 hours.
	// *   Twelve_Hours: 12 hours.
	// *   TwentyFour_Hours: 24 hours.
	MaximumExecutionFrequency *string `json:"MaximumExecutionFrequency,omitempty" xml:"MaximumExecutionFrequency,omitempty"`
	// The trigger type of the managed rule. Valid values:
	//
	// *   ConfigurationItemChangeNotification: The managed rule is triggered by configuration changes.
	// *   ScheduledNotification: The managed rule is triggered as scheduled.
	MessageType *string `json:"MessageType,omitempty" xml:"MessageType,omitempty"`
}

func (s DescribeConfigRuleResponseBodyConfigRuleManagedRuleSourceDetails) String() string {
	return tea.Prettify(s)
}

func (s DescribeConfigRuleResponseBodyConfigRuleManagedRuleSourceDetails) GoString() string {
	return s.String()
}

func (s *DescribeConfigRuleResponseBodyConfigRuleManagedRuleSourceDetails) SetEventSource(v string) *DescribeConfigRuleResponseBodyConfigRuleManagedRuleSourceDetails {
	s.EventSource = &v
	return s
}

func (s *DescribeConfigRuleResponseBodyConfigRuleManagedRuleSourceDetails) SetMaximumExecutionFrequency(v string) *DescribeConfigRuleResponseBodyConfigRuleManagedRuleSourceDetails {
	s.MaximumExecutionFrequency = &v
	return s
}

func (s *DescribeConfigRuleResponseBodyConfigRuleManagedRuleSourceDetails) SetMessageType(v string) *DescribeConfigRuleResponseBodyConfigRuleManagedRuleSourceDetails {
	s.MessageType = &v
	return s
}

type DescribeConfigRuleResponseBodyConfigRuleScope struct {
	// The ID of the resource to be evaluated.
	ComplianceResourceId *string `json:"ComplianceResourceId,omitempty" xml:"ComplianceResourceId,omitempty"`
	// The types of the resources to be evaluated.
	ComplianceResourceTypes []*string `json:"ComplianceResourceTypes,omitempty" xml:"ComplianceResourceTypes,omitempty" type:"Repeated"`
}

func (s DescribeConfigRuleResponseBodyConfigRuleScope) String() string {
	return tea.Prettify(s)
}

func (s DescribeConfigRuleResponseBodyConfigRuleScope) GoString() string {
	return s.String()
}

func (s *DescribeConfigRuleResponseBodyConfigRuleScope) SetComplianceResourceId(v string) *DescribeConfigRuleResponseBodyConfigRuleScope {
	s.ComplianceResourceId = &v
	return s
}

func (s *DescribeConfigRuleResponseBodyConfigRuleScope) SetComplianceResourceTypes(v []*string) *DescribeConfigRuleResponseBodyConfigRuleScope {
	s.ComplianceResourceTypes = v
	return s
}

type DescribeConfigRuleResponseBodyConfigRuleSource struct {
	// The identifier of the rule.
	//
	// *   If the rule was created based on a managed rule, the value of this parameter is the name of the managed rule.
	// *   If the rule is a custom rule, the value of this parameter is the Alibaba Cloud Resource Name (ARN) of the relevant function in Function Compute.
	Identifier *string `json:"Identifier,omitempty" xml:"Identifier,omitempty"`
	// The way in which the rule was created. Valid values:
	//
	// *   CUSTOM_FC: The rule is a custom rule.
	// *   ALIYUN: The rule was created based on a managed rule of Alibaba Cloud.
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The name of the input parameter for the rule.
	SourceConditions []*DescribeConfigRuleResponseBodyConfigRuleSourceSourceConditions `json:"SourceConditions,omitempty" xml:"SourceConditions,omitempty" type:"Repeated"`
	// The intervals at which the rule is triggered. Valid values:
	//
	// *   One_Hour: 1 hour
	// *   Three_Hours: 3 hours
	// *   Six_Hours: 6 hours
	// *   Twelve_Hours: 12 hours.
	// *   TwentyFour_Hours: 24 hours.
	SourceDetails []*DescribeConfigRuleResponseBodyConfigRuleSourceSourceDetails `json:"SourceDetails,omitempty" xml:"SourceDetails,omitempty" type:"Repeated"`
}

func (s DescribeConfigRuleResponseBodyConfigRuleSource) String() string {
	return tea.Prettify(s)
}

func (s DescribeConfigRuleResponseBodyConfigRuleSource) GoString() string {
	return s.String()
}

func (s *DescribeConfigRuleResponseBodyConfigRuleSource) SetIdentifier(v string) *DescribeConfigRuleResponseBodyConfigRuleSource {
	s.Identifier = &v
	return s
}

func (s *DescribeConfigRuleResponseBodyConfigRuleSource) SetOwner(v string) *DescribeConfigRuleResponseBodyConfigRuleSource {
	s.Owner = &v
	return s
}

func (s *DescribeConfigRuleResponseBodyConfigRuleSource) SetSourceConditions(v []*DescribeConfigRuleResponseBodyConfigRuleSourceSourceConditions) *DescribeConfigRuleResponseBodyConfigRuleSource {
	s.SourceConditions = v
	return s
}

func (s *DescribeConfigRuleResponseBodyConfigRuleSource) SetSourceDetails(v []*DescribeConfigRuleResponseBodyConfigRuleSourceSourceDetails) *DescribeConfigRuleResponseBodyConfigRuleSource {
	s.SourceDetails = v
	return s
}

type DescribeConfigRuleResponseBodyConfigRuleSourceSourceConditions struct {
	// The expected value of the input parameter.
	DesiredValue *string `json:"DesiredValue,omitempty" xml:"DesiredValue,omitempty"`
	// The name of the input parameter for the rule.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The operator used to compare the actual value against the expected value of the input parameter. The operator varies based on the type of data returned for the SelectPath parameter.
	//
	// *   Valid values for the String data type:
	//
	//     *   StringEquals: The actual value is equal to the expected value.
	//     *   NotStringEquals: The actual value is not equal to the expected value.
	//     *   StringIn: The actual value exists in the expected value.
	//     *   NotStringIn: The actual value does not exist in the expected value.
	//     *   StringContains: The actual value contains the expected value.
	//     *   NotStringContains: The actual value does not contain the expected value.
	//
	// *   Valid values for the Number data type:
	//
	//     *   Equals: The actual value is equal to the expected value.
	//     *   NotEquals: The actual value is not equal to the expected value.
	//     *   Less: The actual value is less than the expected value.
	//     *   LessOrEquals: The actual value is less than or equal to the expected value.
	//     *   Greater: The actual value is greater than the expected value.
	//     *   GreaterOrEquals: The actual value is greater than or equal to the expected value.
	//
	// *   Valid values for the Base64String data type that indicates a Base64-encoded string:
	//
	//     *   Base64Contains: The actual value contains the expected value.
	//     *   NotBase64Contains: The actual value does not contain the expected value.
	//     *   Base64ContainsAll: The actual value contains all characters in the expected value.
	//     *   Base64ExcludeAll: The actual value excludes all characters in the expected value.
	//
	// *   Valid values for the Array data type:
	//
	//     *   Contains: The actual value contains the expected value.
	//     *   NotContains: The actual value does not contain the expected value.
	//     *   In: The actual value exists in the expected value.
	//     *   NotIn: The actual value does not exist in the expected value.
	//     *   ContainsAll: The actual value contains all elements of the expected value.
	//     *   ExcludeAll: The actual value excludes all elements of the expected value.
	//     *   IsEmpty: The actual value is null.
	Operator *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	// The description of the input parameter.
	Tips *string `json:"Tips,omitempty" xml:"Tips,omitempty"`
}

func (s DescribeConfigRuleResponseBodyConfigRuleSourceSourceConditions) String() string {
	return tea.Prettify(s)
}

func (s DescribeConfigRuleResponseBodyConfigRuleSourceSourceConditions) GoString() string {
	return s.String()
}

func (s *DescribeConfigRuleResponseBodyConfigRuleSourceSourceConditions) SetDesiredValue(v string) *DescribeConfigRuleResponseBodyConfigRuleSourceSourceConditions {
	s.DesiredValue = &v
	return s
}

func (s *DescribeConfigRuleResponseBodyConfigRuleSourceSourceConditions) SetName(v string) *DescribeConfigRuleResponseBodyConfigRuleSourceSourceConditions {
	s.Name = &v
	return s
}

func (s *DescribeConfigRuleResponseBodyConfigRuleSourceSourceConditions) SetOperator(v string) *DescribeConfigRuleResponseBodyConfigRuleSourceSourceConditions {
	s.Operator = &v
	return s
}

func (s *DescribeConfigRuleResponseBodyConfigRuleSourceSourceConditions) SetTips(v string) *DescribeConfigRuleResponseBodyConfigRuleSourceSourceConditions {
	s.Tips = &v
	return s
}

type DescribeConfigRuleResponseBodyConfigRuleSourceSourceDetails struct {
	// The event source of the rule.
	//
	// >  Only events related to Cloud Config are supported. The value is fixed to aliyun.config.
	EventSource *string `json:"EventSource,omitempty" xml:"EventSource,omitempty"`
	// The intervals at which the rule is triggered. Valid values:
	//
	// *   One_Hour: 1 hour.
	// *   Three_Hours: 3 hours.
	// *   Six_Hours: 6 hours.
	// *   Twelve_Hours: 12 hours.
	// *   TwentyFour_Hours: 24 hours.
	MaximumExecutionFrequency *string `json:"MaximumExecutionFrequency,omitempty" xml:"MaximumExecutionFrequency,omitempty"`
	// The trigger type of the rule. Valid values:
	//
	// *   ConfigurationItemChangeNotification: The rule is triggered by configuration changes.
	// *   ScheduledNotification: The rule is triggered as scheduled.
	MessageType *string `json:"MessageType,omitempty" xml:"MessageType,omitempty"`
}

func (s DescribeConfigRuleResponseBodyConfigRuleSourceSourceDetails) String() string {
	return tea.Prettify(s)
}

func (s DescribeConfigRuleResponseBodyConfigRuleSourceSourceDetails) GoString() string {
	return s.String()
}

func (s *DescribeConfigRuleResponseBodyConfigRuleSourceSourceDetails) SetEventSource(v string) *DescribeConfigRuleResponseBodyConfigRuleSourceSourceDetails {
	s.EventSource = &v
	return s
}

func (s *DescribeConfigRuleResponseBodyConfigRuleSourceSourceDetails) SetMaximumExecutionFrequency(v string) *DescribeConfigRuleResponseBodyConfigRuleSourceSourceDetails {
	s.MaximumExecutionFrequency = &v
	return s
}

func (s *DescribeConfigRuleResponseBodyConfigRuleSourceSourceDetails) SetMessageType(v string) *DescribeConfigRuleResponseBodyConfigRuleSourceSourceDetails {
	s.MessageType = &v
	return s
}

type DescribeConfigRuleResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeConfigRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeConfigRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeConfigRuleResponse) GoString() string {
	return s.String()
}

func (s *DescribeConfigRuleResponse) SetHeaders(v map[string]*string) *DescribeConfigRuleResponse {
	s.Headers = v
	return s
}

func (s *DescribeConfigRuleResponse) SetStatusCode(v int32) *DescribeConfigRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeConfigRuleResponse) SetBody(v *DescribeConfigRuleResponseBody) *DescribeConfigRuleResponse {
	s.Body = v
	return s
}

type DescribeConfigurationRecorderResponseBody struct {
	// The details of the configuration recorder that monitors resources.
	ConfigurationRecorder *DescribeConfigurationRecorderResponseBodyConfigurationRecorder `json:"ConfigurationRecorder,omitempty" xml:"ConfigurationRecorder,omitempty" type:"Struct"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeConfigurationRecorderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeConfigurationRecorderResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeConfigurationRecorderResponseBody) SetConfigurationRecorder(v *DescribeConfigurationRecorderResponseBodyConfigurationRecorder) *DescribeConfigurationRecorderResponseBody {
	s.ConfigurationRecorder = v
	return s
}

func (s *DescribeConfigurationRecorderResponseBody) SetRequestId(v string) *DescribeConfigurationRecorderResponseBody {
	s.RequestId = &v
	return s
}

type DescribeConfigurationRecorderResponseBodyConfigurationRecorder struct {
	// The ID of your Alibaba Cloud account.
	AccountId *int64 `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The status of the configuration recorder. Valid values:
	//
	// *   REGISTRABLE: The configuration recorder has not been registered.
	// *   BUILDING: The configuration recorder is being deployed.
	// *   REGISTERED: The configuration recorder has been registered.
	// *   REBUILDING: The configuration recorder is being redeployed.
	ConfigurationRecorderStatus *string `json:"ConfigurationRecorderStatus,omitempty" xml:"ConfigurationRecorderStatus,omitempty"`
	// Indicates whether Cloud Config is upgraded to Cloud Config for Enterprise. Valid values:
	//
	// *   REGISTRABLE: Cloud Config is not upgraded to Cloud Config for Enterprise.
	// *   BUILDING: Cloud Config is being upgraded to Cloud Config for Enterprise.
	// *   REGISTERED: Cloud Config is upgraded to Cloud Config for Enterprise.
	OrganizationEnableStatus *string `json:"OrganizationEnableStatus,omitempty" xml:"OrganizationEnableStatus,omitempty"`
	// The ID of the management account.
	//
	// >  This parameter is returned only when the operation is called by a management account.
	OrganizationMasterId *int64 `json:"OrganizationMasterId,omitempty" xml:"OrganizationMasterId,omitempty"`
	// The types of resources that are monitored by Cloud Config.
	ResourceTypes []*string `json:"ResourceTypes,omitempty" xml:"ResourceTypes,omitempty" type:"Repeated"`
}

func (s DescribeConfigurationRecorderResponseBodyConfigurationRecorder) String() string {
	return tea.Prettify(s)
}

func (s DescribeConfigurationRecorderResponseBodyConfigurationRecorder) GoString() string {
	return s.String()
}

func (s *DescribeConfigurationRecorderResponseBodyConfigurationRecorder) SetAccountId(v int64) *DescribeConfigurationRecorderResponseBodyConfigurationRecorder {
	s.AccountId = &v
	return s
}

func (s *DescribeConfigurationRecorderResponseBodyConfigurationRecorder) SetConfigurationRecorderStatus(v string) *DescribeConfigurationRecorderResponseBodyConfigurationRecorder {
	s.ConfigurationRecorderStatus = &v
	return s
}

func (s *DescribeConfigurationRecorderResponseBodyConfigurationRecorder) SetOrganizationEnableStatus(v string) *DescribeConfigurationRecorderResponseBodyConfigurationRecorder {
	s.OrganizationEnableStatus = &v
	return s
}

func (s *DescribeConfigurationRecorderResponseBodyConfigurationRecorder) SetOrganizationMasterId(v int64) *DescribeConfigurationRecorderResponseBodyConfigurationRecorder {
	s.OrganizationMasterId = &v
	return s
}

func (s *DescribeConfigurationRecorderResponseBodyConfigurationRecorder) SetResourceTypes(v []*string) *DescribeConfigurationRecorderResponseBodyConfigurationRecorder {
	s.ResourceTypes = v
	return s
}

type DescribeConfigurationRecorderResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeConfigurationRecorderResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeConfigurationRecorderResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeConfigurationRecorderResponse) GoString() string {
	return s.String()
}

func (s *DescribeConfigurationRecorderResponse) SetHeaders(v map[string]*string) *DescribeConfigurationRecorderResponse {
	s.Headers = v
	return s
}

func (s *DescribeConfigurationRecorderResponse) SetStatusCode(v int32) *DescribeConfigurationRecorderResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeConfigurationRecorderResponse) SetBody(v *DescribeConfigurationRecorderResponseBody) *DescribeConfigurationRecorderResponse {
	s.Body = v
	return s
}

type DescribeDeliveryChannelsRequest struct {
	// The ID of the delivery channel. Separate multiple IDs with commas (,).
	//
	// For more information about how to obtain the ID of a delivery channel, see [PutDeliveryChannel](~~174253~~).
	DeliveryChannelIds *string `json:"DeliveryChannelIds,omitempty" xml:"DeliveryChannelIds,omitempty"`
}

func (s DescribeDeliveryChannelsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDeliveryChannelsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDeliveryChannelsRequest) SetDeliveryChannelIds(v string) *DescribeDeliveryChannelsRequest {
	s.DeliveryChannelIds = &v
	return s
}

type DescribeDeliveryChannelsResponseBody struct {
	// The information about the delivery channel.
	DeliveryChannels []*DescribeDeliveryChannelsResponseBodyDeliveryChannels `json:"DeliveryChannels,omitempty" xml:"DeliveryChannels,omitempty" type:"Repeated"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDeliveryChannelsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDeliveryChannelsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDeliveryChannelsResponseBody) SetDeliveryChannels(v []*DescribeDeliveryChannelsResponseBodyDeliveryChannels) *DescribeDeliveryChannelsResponseBody {
	s.DeliveryChannels = v
	return s
}

func (s *DescribeDeliveryChannelsResponseBody) SetRequestId(v string) *DescribeDeliveryChannelsResponseBody {
	s.RequestId = &v
	return s
}

type DescribeDeliveryChannelsResponseBodyDeliveryChannels struct {
	// Indicates whether the resource change logs are received by the specified destination. If you set this parameter to true, Cloud Config delivers the resource change logs to OSS, Log Service, or MNS when the configurations of the resources change. Valid values:
	//
	// *   true: The resource change logs are received by the specified destination.
	// *   false: The resource change logs are not received by the specified destination.
	ConfigurationItemChangeNotification *bool `json:"ConfigurationItemChangeNotification,omitempty" xml:"ConfigurationItemChangeNotification,omitempty"`
	// Indicates whether the scheduled resource snapshots are received by the OSS bucket. If you set this parameter to true, Cloud Config delivers the scheduled resource snapshots to the specified OSS bucket at 00:00:00 and 12:00:00 every day. Valid values:
	//
	// *   true: The scheduled resource snapshots are received by the specified OSS bucket.
	// *   false: The scheduled resource snapshots are not received by the specified OSS bucket.
	ConfigurationSnapshot *bool `json:"ConfigurationSnapshot,omitempty" xml:"ConfigurationSnapshot,omitempty"`
	// The Alibaba Cloud Resource Name (ARN) of the role assumed by the delivery channel.
	DeliveryChannelAssumeRoleArn *string `json:"DeliveryChannelAssumeRoleArn,omitempty" xml:"DeliveryChannelAssumeRoleArn,omitempty"`
	// The rule attached to the delivery channel. This parameter is applicable only to delivery channels of the MNS type.
	//
	// This parameter indicates the lowest risk level for the events to subscribe to and the resource types for which you subscribe to events.
	//
	// *   The setting of the lowest risk level for the events to subscribe to is in the following format: `{"filterType":"RuleRiskLevel","value":"1","multiple":false}`.
	//
	//     The value field indicates the lowest risk level. Valid values: 1, 2, and 3. A value of 1 indicates the high risk level. A value of 2 indicates the medium risk level. A value of 3 indicates the low risk level.
	//
	// *   The setting of the resource types for which you subscribe to events is in the following format: `{"filterType":"ResourceType","values":["ACS::ACK::Cluster","ACS::ActionTrail::Trail","ACS::CBWP::CommonBandwidthPackage"],"multiple":true}`.
	//
	//     The values field indicates the resource types for which you subscribe to events. Its value is a JSON array.
	//
	//     Example: `[{"filterType":"ResourceType","values":["ACS::ActionTrail::Trail","ACS::CBWP::CommonBandwidthPackage","ACS::CDN::Domain","ACS::CEN::CenBandwidthPackage","ACS::CEN::CenInstance","ACS::CEN::Flowlog","ACS::DdosCoo::Instance"],"multiple":true}]`
	DeliveryChannelCondition *string `json:"DeliveryChannelCondition,omitempty" xml:"DeliveryChannelCondition,omitempty"`
	// The ID of the delivery channel.
	DeliveryChannelId *string `json:"DeliveryChannelId,omitempty" xml:"DeliveryChannelId,omitempty"`
	// The name of the delivery channel.
	DeliveryChannelName *string `json:"DeliveryChannelName,omitempty" xml:"DeliveryChannelName,omitempty"`
	// The ARN of the delivery destination.
	//
	// *   If the DeliveryChannelType parameter is set to OSS, the value of this parameter is the ARN of the destination OSS bucket.
	// *   If the DeliveryChannelType parameter is set to MNS, the value of this parameter is the ARN of the destination MNS topic.
	// *   If the DeliveryChannelType parameter is set to SLS, the value is the ARN of the destination Log Service Logstore.
	DeliveryChannelTargetArn *string `json:"DeliveryChannelTargetArn,omitempty" xml:"DeliveryChannelTargetArn,omitempty"`
	// The type of the delivery channel. Valid values:
	//
	// *   OSS: Object Storage Service (OSS)
	// *   MNS: Message Service (MNS)
	// *   SLS: Log Service
	DeliveryChannelType *string `json:"DeliveryChannelType,omitempty" xml:"DeliveryChannelType,omitempty"`
	// The description of the delivery channel.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Indicates whether the resource non-compliance events are received by the specified destination. If you set this parameter to true, Cloud Config delivers resource non-compliance events to Log Service or MNS when resources are evaluated as non-compliant. Valid values:
	//
	// *   true: The resource non-compliance events are received by the specified destination.
	// *   false: The resource non-compliance events are not received by the specified destination.
	NonCompliantNotification *bool `json:"NonCompliantNotification,omitempty" xml:"NonCompliantNotification,omitempty"`
	// The ARN of the OSS bucket to which the delivered data is transferred when the size of the data exceeds the specified upper limit of the delivery channel.
	OversizedDataOSSTargetArn *string `json:"OversizedDataOSSTargetArn,omitempty" xml:"OversizedDataOSSTargetArn,omitempty"`
	// The status of the delivery channel. Valid values:
	//
	// *   0: The delivery channel is disabled.
	// *   1: The delivery channel is enabled.
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeDeliveryChannelsResponseBodyDeliveryChannels) String() string {
	return tea.Prettify(s)
}

func (s DescribeDeliveryChannelsResponseBodyDeliveryChannels) GoString() string {
	return s.String()
}

func (s *DescribeDeliveryChannelsResponseBodyDeliveryChannels) SetConfigurationItemChangeNotification(v bool) *DescribeDeliveryChannelsResponseBodyDeliveryChannels {
	s.ConfigurationItemChangeNotification = &v
	return s
}

func (s *DescribeDeliveryChannelsResponseBodyDeliveryChannels) SetConfigurationSnapshot(v bool) *DescribeDeliveryChannelsResponseBodyDeliveryChannels {
	s.ConfigurationSnapshot = &v
	return s
}

func (s *DescribeDeliveryChannelsResponseBodyDeliveryChannels) SetDeliveryChannelAssumeRoleArn(v string) *DescribeDeliveryChannelsResponseBodyDeliveryChannels {
	s.DeliveryChannelAssumeRoleArn = &v
	return s
}

func (s *DescribeDeliveryChannelsResponseBodyDeliveryChannels) SetDeliveryChannelCondition(v string) *DescribeDeliveryChannelsResponseBodyDeliveryChannels {
	s.DeliveryChannelCondition = &v
	return s
}

func (s *DescribeDeliveryChannelsResponseBodyDeliveryChannels) SetDeliveryChannelId(v string) *DescribeDeliveryChannelsResponseBodyDeliveryChannels {
	s.DeliveryChannelId = &v
	return s
}

func (s *DescribeDeliveryChannelsResponseBodyDeliveryChannels) SetDeliveryChannelName(v string) *DescribeDeliveryChannelsResponseBodyDeliveryChannels {
	s.DeliveryChannelName = &v
	return s
}

func (s *DescribeDeliveryChannelsResponseBodyDeliveryChannels) SetDeliveryChannelTargetArn(v string) *DescribeDeliveryChannelsResponseBodyDeliveryChannels {
	s.DeliveryChannelTargetArn = &v
	return s
}

func (s *DescribeDeliveryChannelsResponseBodyDeliveryChannels) SetDeliveryChannelType(v string) *DescribeDeliveryChannelsResponseBodyDeliveryChannels {
	s.DeliveryChannelType = &v
	return s
}

func (s *DescribeDeliveryChannelsResponseBodyDeliveryChannels) SetDescription(v string) *DescribeDeliveryChannelsResponseBodyDeliveryChannels {
	s.Description = &v
	return s
}

func (s *DescribeDeliveryChannelsResponseBodyDeliveryChannels) SetNonCompliantNotification(v bool) *DescribeDeliveryChannelsResponseBodyDeliveryChannels {
	s.NonCompliantNotification = &v
	return s
}

func (s *DescribeDeliveryChannelsResponseBodyDeliveryChannels) SetOversizedDataOSSTargetArn(v string) *DescribeDeliveryChannelsResponseBodyDeliveryChannels {
	s.OversizedDataOSSTargetArn = &v
	return s
}

func (s *DescribeDeliveryChannelsResponseBodyDeliveryChannels) SetStatus(v int32) *DescribeDeliveryChannelsResponseBodyDeliveryChannels {
	s.Status = &v
	return s
}

type DescribeDeliveryChannelsResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeDeliveryChannelsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDeliveryChannelsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDeliveryChannelsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDeliveryChannelsResponse) SetHeaders(v map[string]*string) *DescribeDeliveryChannelsResponse {
	s.Headers = v
	return s
}

func (s *DescribeDeliveryChannelsResponse) SetStatusCode(v int32) *DescribeDeliveryChannelsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDeliveryChannelsResponse) SetBody(v *DescribeDeliveryChannelsResponseBody) *DescribeDeliveryChannelsResponse {
	s.Body = v
	return s
}

type DescribeDiscoveredResourceRequest struct {
	// This parameter is scheduled to be removed before 00:00:00, June 30, 2021. Account group-related APIs will be provided as an alternative before 00:00:00, May 30, 2021. If you are using this parameter, we recommend that you switch to account group-related APIs after 00:00:00, May 30, 2021. For more information, see [Account groups](~~211534~~).
	MemberId *int64 `json:"MemberId,omitempty" xml:"MemberId,omitempty"`
	// This parameter is scheduled to be removed before 00:00:00, June 30, 2021. Account group-related APIs will be provided as an alternative before 00:00:00, May 30, 2021. If you are using this parameter, we recommend that you switch to account group-related APIs after 00:00:00, May 30, 2021. For more information, see [Account groups](~~211534~~).
	MultiAccount *bool `json:"MultiAccount,omitempty" xml:"MultiAccount,omitempty"`
	// The ID of the region.
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The ID of the resource.
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The type of the resource.
	//
	// >  You can call the GetSupportedResourceTypes operation to query the resource types supported by Cloud Config. For more information, see [GetSupportedResourceTypes](~~169618~~).
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s DescribeDiscoveredResourceRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDiscoveredResourceRequest) GoString() string {
	return s.String()
}

func (s *DescribeDiscoveredResourceRequest) SetMemberId(v int64) *DescribeDiscoveredResourceRequest {
	s.MemberId = &v
	return s
}

func (s *DescribeDiscoveredResourceRequest) SetMultiAccount(v bool) *DescribeDiscoveredResourceRequest {
	s.MultiAccount = &v
	return s
}

func (s *DescribeDiscoveredResourceRequest) SetRegion(v string) *DescribeDiscoveredResourceRequest {
	s.Region = &v
	return s
}

func (s *DescribeDiscoveredResourceRequest) SetResourceId(v string) *DescribeDiscoveredResourceRequest {
	s.ResourceId = &v
	return s
}

func (s *DescribeDiscoveredResourceRequest) SetResourceType(v string) *DescribeDiscoveredResourceRequest {
	s.ResourceType = &v
	return s
}

type DescribeDiscoveredResourceResponseBody struct {
	// The information about the resource.
	DiscoveredResourceDetail *DescribeDiscoveredResourceResponseBodyDiscoveredResourceDetail `json:"DiscoveredResourceDetail,omitempty" xml:"DiscoveredResourceDetail,omitempty" type:"Struct"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDiscoveredResourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDiscoveredResourceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDiscoveredResourceResponseBody) SetDiscoveredResourceDetail(v *DescribeDiscoveredResourceResponseBodyDiscoveredResourceDetail) *DescribeDiscoveredResourceResponseBody {
	s.DiscoveredResourceDetail = v
	return s
}

func (s *DescribeDiscoveredResourceResponseBody) SetRequestId(v string) *DescribeDiscoveredResourceResponseBody {
	s.RequestId = &v
	return s
}

type DescribeDiscoveredResourceResponseBodyDiscoveredResourceDetail struct {
	// The ID of your Alibaba Cloud account.
	AccountId *int64 `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The zone where the resource resides.
	AvailabilityZone *string `json:"AvailabilityZone,omitempty" xml:"AvailabilityZone,omitempty"`
	// The configuration of the resource.
	Configuration *string `json:"Configuration,omitempty" xml:"Configuration,omitempty"`
	// The ID of the region.
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The timestamp when the resource was created.
	ResourceCreationTime *int64 `json:"ResourceCreationTime,omitempty" xml:"ResourceCreationTime,omitempty"`
	// Indicates whether the resource is deleted. Valid values:
	//
	// *   1: The resource is not deleted.
	// *   0: The resource is deleted.
	ResourceDeleted *int32 `json:"ResourceDeleted,omitempty" xml:"ResourceDeleted,omitempty"`
	// The ID of the resource.
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The name of the resource.
	ResourceName *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	// The status of the resource. The parameter value varies with the resource type and may be left empty. Examples:
	//
	// *   If ResourceType is set to ACS::ECS::Instance, the resource is an Elastic Compute Service (ECS) instance which has a specific state. In this case, valid values of this parameter include Running and Stopped.
	// *   If ResourceType is set to ACS::OSS::Bucket, the resource is an Object Storage Service (OSS) bucket which does not have a specific state. In this case, this parameter is left empty.
	ResourceStatus *string `json:"ResourceStatus,omitempty" xml:"ResourceStatus,omitempty"`
	// The type of the resource.
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The tags of the resource.
	Tags *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
}

func (s DescribeDiscoveredResourceResponseBodyDiscoveredResourceDetail) String() string {
	return tea.Prettify(s)
}

func (s DescribeDiscoveredResourceResponseBodyDiscoveredResourceDetail) GoString() string {
	return s.String()
}

func (s *DescribeDiscoveredResourceResponseBodyDiscoveredResourceDetail) SetAccountId(v int64) *DescribeDiscoveredResourceResponseBodyDiscoveredResourceDetail {
	s.AccountId = &v
	return s
}

func (s *DescribeDiscoveredResourceResponseBodyDiscoveredResourceDetail) SetAvailabilityZone(v string) *DescribeDiscoveredResourceResponseBodyDiscoveredResourceDetail {
	s.AvailabilityZone = &v
	return s
}

func (s *DescribeDiscoveredResourceResponseBodyDiscoveredResourceDetail) SetConfiguration(v string) *DescribeDiscoveredResourceResponseBodyDiscoveredResourceDetail {
	s.Configuration = &v
	return s
}

func (s *DescribeDiscoveredResourceResponseBodyDiscoveredResourceDetail) SetRegion(v string) *DescribeDiscoveredResourceResponseBodyDiscoveredResourceDetail {
	s.Region = &v
	return s
}

func (s *DescribeDiscoveredResourceResponseBodyDiscoveredResourceDetail) SetResourceCreationTime(v int64) *DescribeDiscoveredResourceResponseBodyDiscoveredResourceDetail {
	s.ResourceCreationTime = &v
	return s
}

func (s *DescribeDiscoveredResourceResponseBodyDiscoveredResourceDetail) SetResourceDeleted(v int32) *DescribeDiscoveredResourceResponseBodyDiscoveredResourceDetail {
	s.ResourceDeleted = &v
	return s
}

func (s *DescribeDiscoveredResourceResponseBodyDiscoveredResourceDetail) SetResourceId(v string) *DescribeDiscoveredResourceResponseBodyDiscoveredResourceDetail {
	s.ResourceId = &v
	return s
}

func (s *DescribeDiscoveredResourceResponseBodyDiscoveredResourceDetail) SetResourceName(v string) *DescribeDiscoveredResourceResponseBodyDiscoveredResourceDetail {
	s.ResourceName = &v
	return s
}

func (s *DescribeDiscoveredResourceResponseBodyDiscoveredResourceDetail) SetResourceStatus(v string) *DescribeDiscoveredResourceResponseBodyDiscoveredResourceDetail {
	s.ResourceStatus = &v
	return s
}

func (s *DescribeDiscoveredResourceResponseBodyDiscoveredResourceDetail) SetResourceType(v string) *DescribeDiscoveredResourceResponseBodyDiscoveredResourceDetail {
	s.ResourceType = &v
	return s
}

func (s *DescribeDiscoveredResourceResponseBodyDiscoveredResourceDetail) SetTags(v string) *DescribeDiscoveredResourceResponseBodyDiscoveredResourceDetail {
	s.Tags = &v
	return s
}

type DescribeDiscoveredResourceResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeDiscoveredResourceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDiscoveredResourceResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDiscoveredResourceResponse) GoString() string {
	return s.String()
}

func (s *DescribeDiscoveredResourceResponse) SetHeaders(v map[string]*string) *DescribeDiscoveredResourceResponse {
	s.Headers = v
	return s
}

func (s *DescribeDiscoveredResourceResponse) SetStatusCode(v int32) *DescribeDiscoveredResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDiscoveredResourceResponse) SetBody(v *DescribeDiscoveredResourceResponseBody) *DescribeDiscoveredResourceResponse {
	s.Body = v
	return s
}

type DescribeEvaluationResultsRequest struct {
	// The compliance evaluation result. Valid values:
	//
	// *   COMPLIANT: The resources are evaluated as compliant.
	// *   NON_COMPLIANT: The resources are evaluated as non-compliant.
	// *   NOT_APPLICABLE: The rule does not apply to the resources.
	// *   INSUFFICIENT_DATA: The resource data is insufficient.
	ComplianceType *string `json:"ComplianceType,omitempty" xml:"ComplianceType,omitempty"`
	// The rule ID. If you query the compliance evaluation results of resources by rule, you must specify this parameter.
	//
	// > You must specify the ResourceType or ConfigRuleId parameter.
	ConfigRuleId *string `json:"ConfigRuleId,omitempty" xml:"ConfigRuleId,omitempty"`
	// This parameter is unavailable since 00:00:00 on June 30, 2021. Account group-related operations are provided as an alternative after 00:00:00 on May 30, 2021. If you are using this parameter, we recommend that you switch to account group-related operations after 00:00:00 on May 30, 2021. For information about account groups, see [Overview](~~211534~~).
	MemberId *int64 `json:"MemberId,omitempty" xml:"MemberId,omitempty"`
	// This parameter is unavailable since 00:00:00 on June 30, 2021. Account group-related operations are provided as an alternative after 00:00:00 on May 30, 2021. If you are using this parameter, we recommend that you switch to account group-related operations after 00:00:00 on May 30, 2021. For information about account groups, see [Overview](~~211534~~).
	MultiAccount *bool `json:"MultiAccount,omitempty" xml:"MultiAccount,omitempty"`
	// The number of the page to return. Pages start from page 1.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return per page. Valid values: 1 to 100.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The resource ID. If you query the compliance evaluation results of resources by resource type, you must specify this parameter.
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The type of resource. If you query the compliance evaluation results of resources by resource type, you must specify this parameter.
	//
	// You can call the GetSupportedResourceTypes operation to query the resource types supported by Cloud Config. For more information, see [GetSupportedResourceTypes](~~169618~~).
	//
	// > You must specify the ResourceType or ConfigRuleId parameter.
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s DescribeEvaluationResultsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeEvaluationResultsRequest) GoString() string {
	return s.String()
}

func (s *DescribeEvaluationResultsRequest) SetComplianceType(v string) *DescribeEvaluationResultsRequest {
	s.ComplianceType = &v
	return s
}

func (s *DescribeEvaluationResultsRequest) SetConfigRuleId(v string) *DescribeEvaluationResultsRequest {
	s.ConfigRuleId = &v
	return s
}

func (s *DescribeEvaluationResultsRequest) SetMemberId(v int64) *DescribeEvaluationResultsRequest {
	s.MemberId = &v
	return s
}

func (s *DescribeEvaluationResultsRequest) SetMultiAccount(v bool) *DescribeEvaluationResultsRequest {
	s.MultiAccount = &v
	return s
}

func (s *DescribeEvaluationResultsRequest) SetPageNumber(v int32) *DescribeEvaluationResultsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeEvaluationResultsRequest) SetPageSize(v int32) *DescribeEvaluationResultsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeEvaluationResultsRequest) SetResourceId(v string) *DescribeEvaluationResultsRequest {
	s.ResourceId = &v
	return s
}

func (s *DescribeEvaluationResultsRequest) SetResourceType(v string) *DescribeEvaluationResultsRequest {
	s.ResourceType = &v
	return s
}

type DescribeEvaluationResultsResponseBody struct {
	// The information about the compliance evaluation results returned.
	EvaluationResults *DescribeEvaluationResultsResponseBodyEvaluationResults `json:"EvaluationResults,omitempty" xml:"EvaluationResults,omitempty" type:"Struct"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeEvaluationResultsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeEvaluationResultsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeEvaluationResultsResponseBody) SetEvaluationResults(v *DescribeEvaluationResultsResponseBodyEvaluationResults) *DescribeEvaluationResultsResponseBody {
	s.EvaluationResults = v
	return s
}

func (s *DescribeEvaluationResultsResponseBody) SetRequestId(v string) *DescribeEvaluationResultsResponseBody {
	s.RequestId = &v
	return s
}

type DescribeEvaluationResultsResponseBodyEvaluationResults struct {
	// Indicates whether remediation is enabled. Valid values:
	//
	// *   true
	// *   false
	EvaluationResultList []*DescribeEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList `json:"EvaluationResultList,omitempty" xml:"EvaluationResultList,omitempty" type:"Repeated"`
	// The page number of the returned page. Pages start from page 1.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page. Valid values: 1 to 100.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of compliance evaluation results.
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeEvaluationResultsResponseBodyEvaluationResults) String() string {
	return tea.Prettify(s)
}

func (s DescribeEvaluationResultsResponseBodyEvaluationResults) GoString() string {
	return s.String()
}

func (s *DescribeEvaluationResultsResponseBodyEvaluationResults) SetEvaluationResultList(v []*DescribeEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList) *DescribeEvaluationResultsResponseBodyEvaluationResults {
	s.EvaluationResultList = v
	return s
}

func (s *DescribeEvaluationResultsResponseBodyEvaluationResults) SetPageNumber(v int32) *DescribeEvaluationResultsResponseBodyEvaluationResults {
	s.PageNumber = &v
	return s
}

func (s *DescribeEvaluationResultsResponseBodyEvaluationResults) SetPageSize(v int32) *DescribeEvaluationResultsResponseBodyEvaluationResults {
	s.PageSize = &v
	return s
}

func (s *DescribeEvaluationResultsResponseBodyEvaluationResults) SetTotalCount(v int64) *DescribeEvaluationResultsResponseBodyEvaluationResults {
	s.TotalCount = &v
	return s
}

type DescribeEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList struct {
	// The annotation to the resource that is evaluated to be non-compliant.
	Annotation *string `json:"Annotation,omitempty" xml:"Annotation,omitempty"`
	// The compliance evaluation result of the resource. Valid values:
	//
	// *   COMPLIANT: The resource is evaluated to be compliant.
	// *   NON_COMPLIANT: The resource is evaluated to be non-compliant.
	// *   NOT_APPLICABLE: The rule does not apply to the resource.
	// *   INSUFFICIENT_DATA: The resource data is insufficient.
	ComplianceType *string `json:"ComplianceType,omitempty" xml:"ComplianceType,omitempty"`
	// The timestamp that was generated when the rule was triggered.
	ConfigRuleInvokedTimestamp *int64 `json:"ConfigRuleInvokedTimestamp,omitempty" xml:"ConfigRuleInvokedTimestamp,omitempty"`
	// The identifier of the compliance evaluation result.
	EvaluationResultIdentifier *DescribeEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifier `json:"EvaluationResultIdentifier,omitempty" xml:"EvaluationResultIdentifier,omitempty" type:"Struct"`
	// The trigger type of the rule. Valid values:
	//
	// *   ConfigurationItemChangeNotification: The rule is triggered by configuration changes.
	// *   ScheduledNotification: The rule is triggered as scheduled.
	InvokingEventMessageType *string `json:"InvokingEventMessageType,omitempty" xml:"InvokingEventMessageType,omitempty"`
	// Specifies whether to enable the remediation template. Valid values:
	//
	// * **true**
	// * **false**
	RemediationEnabled *bool `json:"RemediationEnabled,omitempty" xml:"RemediationEnabled,omitempty"`
	// The timestamp that was generated when the compliance evaluation result was recorded.
	ResultRecordedTimestamp *int64 `json:"ResultRecordedTimestamp,omitempty" xml:"ResultRecordedTimestamp,omitempty"`
	// The risk level of the non-compliant resource. Valid values:
	//
	// *   1: high risk
	// *   2: medium risk
	// *   3: low risk
	RiskLevel *int32 `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
}

func (s DescribeEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList) String() string {
	return tea.Prettify(s)
}

func (s DescribeEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList) GoString() string {
	return s.String()
}

func (s *DescribeEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList) SetAnnotation(v string) *DescribeEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList {
	s.Annotation = &v
	return s
}

func (s *DescribeEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList) SetComplianceType(v string) *DescribeEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList {
	s.ComplianceType = &v
	return s
}

func (s *DescribeEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList) SetConfigRuleInvokedTimestamp(v int64) *DescribeEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList {
	s.ConfigRuleInvokedTimestamp = &v
	return s
}

func (s *DescribeEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList) SetEvaluationResultIdentifier(v *DescribeEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifier) *DescribeEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList {
	s.EvaluationResultIdentifier = v
	return s
}

func (s *DescribeEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList) SetInvokingEventMessageType(v string) *DescribeEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList {
	s.InvokingEventMessageType = &v
	return s
}

func (s *DescribeEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList) SetRemediationEnabled(v bool) *DescribeEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList {
	s.RemediationEnabled = &v
	return s
}

func (s *DescribeEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList) SetResultRecordedTimestamp(v int64) *DescribeEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList {
	s.ResultRecordedTimestamp = &v
	return s
}

func (s *DescribeEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList) SetRiskLevel(v int32) *DescribeEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList {
	s.RiskLevel = &v
	return s
}

type DescribeEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifier struct {
	// The information of the rule and the evaluated resource.
	EvaluationResultQualifier *DescribeEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier `json:"EvaluationResultQualifier,omitempty" xml:"EvaluationResultQualifier,omitempty" type:"Struct"`
	// The timestamp that was generated when the compliance evaluation was performed.
	OrderingTimestamp *int64 `json:"OrderingTimestamp,omitempty" xml:"OrderingTimestamp,omitempty"`
}

func (s DescribeEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifier) String() string {
	return tea.Prettify(s)
}

func (s DescribeEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifier) GoString() string {
	return s.String()
}

func (s *DescribeEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifier) SetEvaluationResultQualifier(v *DescribeEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier) *DescribeEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifier {
	s.EvaluationResultQualifier = v
	return s
}

func (s *DescribeEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifier) SetOrderingTimestamp(v int64) *DescribeEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifier {
	s.OrderingTimestamp = &v
	return s
}

type DescribeEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier struct {
	// The Alibaba Cloud Resource Name (ARN) of the rule.
	ConfigRuleArn *string `json:"ConfigRuleArn,omitempty" xml:"ConfigRuleArn,omitempty"`
	// The ID of the rule.
	ConfigRuleId *string `json:"ConfigRuleId,omitempty" xml:"ConfigRuleId,omitempty"`
	// The name of the rule.
	ConfigRuleName *string `json:"ConfigRuleName,omitempty" xml:"ConfigRuleName,omitempty"`
	// The ID of the region.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the evaluated resource.
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The type of the evaluated resource.
	ResourceName *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	// The type of the evaluated resource.
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s DescribeEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier) String() string {
	return tea.Prettify(s)
}

func (s DescribeEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier) GoString() string {
	return s.String()
}

func (s *DescribeEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier) SetConfigRuleArn(v string) *DescribeEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier {
	s.ConfigRuleArn = &v
	return s
}

func (s *DescribeEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier) SetConfigRuleId(v string) *DescribeEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier {
	s.ConfigRuleId = &v
	return s
}

func (s *DescribeEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier) SetConfigRuleName(v string) *DescribeEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier {
	s.ConfigRuleName = &v
	return s
}

func (s *DescribeEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier) SetRegionId(v string) *DescribeEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier {
	s.RegionId = &v
	return s
}

func (s *DescribeEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier) SetResourceId(v string) *DescribeEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier {
	s.ResourceId = &v
	return s
}

func (s *DescribeEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier) SetResourceName(v string) *DescribeEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier {
	s.ResourceName = &v
	return s
}

func (s *DescribeEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier) SetResourceType(v string) *DescribeEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier {
	s.ResourceType = &v
	return s
}

type DescribeEvaluationResultsResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeEvaluationResultsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeEvaluationResultsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeEvaluationResultsResponse) GoString() string {
	return s.String()
}

func (s *DescribeEvaluationResultsResponse) SetHeaders(v map[string]*string) *DescribeEvaluationResultsResponse {
	s.Headers = v
	return s
}

func (s *DescribeEvaluationResultsResponse) SetStatusCode(v int32) *DescribeEvaluationResultsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeEvaluationResultsResponse) SetBody(v *DescribeEvaluationResultsResponseBody) *DescribeEvaluationResultsResponse {
	s.Body = v
	return s
}

type GetAggregateDiscoveredResourceRequest struct {
	// The ID of the account group.
	//
	// For more information about how to obtain the ID of an account group, see [ListAggregators](~~255797~~).
	AggregatorId *string `json:"AggregatorId,omitempty" xml:"AggregatorId,omitempty"`
	// The ID of the region where the resource resides.
	//
	// For more information about how to obtain the region ID of a resource, see [ListAggregateDiscoveredResources](~~265983~~).
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The ID of the resource.
	//
	// For more information about how to obtain the ID of a resource, see [ListAggregateDiscoveredResources](~~265983~~).
	ResourceId      *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The type of the resource.
	//
	// For more information about how to obtain the type of a resource, see [ListAggregateDiscoveredResources](~~265983~~).
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s GetAggregateDiscoveredResourceRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAggregateDiscoveredResourceRequest) GoString() string {
	return s.String()
}

func (s *GetAggregateDiscoveredResourceRequest) SetAggregatorId(v string) *GetAggregateDiscoveredResourceRequest {
	s.AggregatorId = &v
	return s
}

func (s *GetAggregateDiscoveredResourceRequest) SetRegion(v string) *GetAggregateDiscoveredResourceRequest {
	s.Region = &v
	return s
}

func (s *GetAggregateDiscoveredResourceRequest) SetResourceId(v string) *GetAggregateDiscoveredResourceRequest {
	s.ResourceId = &v
	return s
}

func (s *GetAggregateDiscoveredResourceRequest) SetResourceOwnerId(v int64) *GetAggregateDiscoveredResourceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetAggregateDiscoveredResourceRequest) SetResourceType(v string) *GetAggregateDiscoveredResourceRequest {
	s.ResourceType = &v
	return s
}

type GetAggregateDiscoveredResourceResponseBody struct {
	// The information about the resource.
	DiscoveredResourceDetail *GetAggregateDiscoveredResourceResponseBodyDiscoveredResourceDetail `json:"DiscoveredResourceDetail,omitempty" xml:"DiscoveredResourceDetail,omitempty" type:"Struct"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetAggregateDiscoveredResourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAggregateDiscoveredResourceResponseBody) GoString() string {
	return s.String()
}

func (s *GetAggregateDiscoveredResourceResponseBody) SetDiscoveredResourceDetail(v *GetAggregateDiscoveredResourceResponseBodyDiscoveredResourceDetail) *GetAggregateDiscoveredResourceResponseBody {
	s.DiscoveredResourceDetail = v
	return s
}

func (s *GetAggregateDiscoveredResourceResponseBody) SetRequestId(v string) *GetAggregateDiscoveredResourceResponseBody {
	s.RequestId = &v
	return s
}

type GetAggregateDiscoveredResourceResponseBodyDiscoveredResourceDetail struct {
	// The ID of the Alibaba Cloud account to which the resource belongs.
	AccountId *int64 `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The zone where the resource resides.
	AvailabilityZone *string `json:"AvailabilityZone,omitempty" xml:"AvailabilityZone,omitempty"`
	// The configuration of the resource.
	Configuration *string `json:"Configuration,omitempty" xml:"Configuration,omitempty"`
	// The ID of the region where the resource resides.
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The timestamp when the resource was created.
	ResourceCreationTime *int64 `json:"ResourceCreationTime,omitempty" xml:"ResourceCreationTime,omitempty"`
	// Indicates whether the resource is deleted. Valid values:
	//
	// *   1: The resource is retained.
	// *   0: The resource is deleted.
	ResourceDeleted *int32 `json:"ResourceDeleted,omitempty" xml:"ResourceDeleted,omitempty"`
	// The ID of the resource.
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The name of the resource.
	ResourceName *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	// The status of the resource. The parameter value varies with the resource type and may be left empty. Examples:
	//
	// *   If the ResourceType parameter is set to ACS::ECS::Instance, the resource is an Elastic Compute Service (ECS) instance that has a specific state. In this case, valid values of this parameter include Running and Stopped.
	// *   If the ResourceType parameter is set to ACS::OSS::Bucket, the resource is an Object Storage Service (OSS) bucket that does not have a specific state. In this case, this parameter is left empty.
	ResourceStatus *string `json:"ResourceStatus,omitempty" xml:"ResourceStatus,omitempty"`
	// The type of the resource.
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The tags of the resource.
	Tags *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
}

func (s GetAggregateDiscoveredResourceResponseBodyDiscoveredResourceDetail) String() string {
	return tea.Prettify(s)
}

func (s GetAggregateDiscoveredResourceResponseBodyDiscoveredResourceDetail) GoString() string {
	return s.String()
}

func (s *GetAggregateDiscoveredResourceResponseBodyDiscoveredResourceDetail) SetAccountId(v int64) *GetAggregateDiscoveredResourceResponseBodyDiscoveredResourceDetail {
	s.AccountId = &v
	return s
}

func (s *GetAggregateDiscoveredResourceResponseBodyDiscoveredResourceDetail) SetAvailabilityZone(v string) *GetAggregateDiscoveredResourceResponseBodyDiscoveredResourceDetail {
	s.AvailabilityZone = &v
	return s
}

func (s *GetAggregateDiscoveredResourceResponseBodyDiscoveredResourceDetail) SetConfiguration(v string) *GetAggregateDiscoveredResourceResponseBodyDiscoveredResourceDetail {
	s.Configuration = &v
	return s
}

func (s *GetAggregateDiscoveredResourceResponseBodyDiscoveredResourceDetail) SetRegion(v string) *GetAggregateDiscoveredResourceResponseBodyDiscoveredResourceDetail {
	s.Region = &v
	return s
}

func (s *GetAggregateDiscoveredResourceResponseBodyDiscoveredResourceDetail) SetResourceCreationTime(v int64) *GetAggregateDiscoveredResourceResponseBodyDiscoveredResourceDetail {
	s.ResourceCreationTime = &v
	return s
}

func (s *GetAggregateDiscoveredResourceResponseBodyDiscoveredResourceDetail) SetResourceDeleted(v int32) *GetAggregateDiscoveredResourceResponseBodyDiscoveredResourceDetail {
	s.ResourceDeleted = &v
	return s
}

func (s *GetAggregateDiscoveredResourceResponseBodyDiscoveredResourceDetail) SetResourceId(v string) *GetAggregateDiscoveredResourceResponseBodyDiscoveredResourceDetail {
	s.ResourceId = &v
	return s
}

func (s *GetAggregateDiscoveredResourceResponseBodyDiscoveredResourceDetail) SetResourceName(v string) *GetAggregateDiscoveredResourceResponseBodyDiscoveredResourceDetail {
	s.ResourceName = &v
	return s
}

func (s *GetAggregateDiscoveredResourceResponseBodyDiscoveredResourceDetail) SetResourceStatus(v string) *GetAggregateDiscoveredResourceResponseBodyDiscoveredResourceDetail {
	s.ResourceStatus = &v
	return s
}

func (s *GetAggregateDiscoveredResourceResponseBodyDiscoveredResourceDetail) SetResourceType(v string) *GetAggregateDiscoveredResourceResponseBodyDiscoveredResourceDetail {
	s.ResourceType = &v
	return s
}

func (s *GetAggregateDiscoveredResourceResponseBodyDiscoveredResourceDetail) SetTags(v string) *GetAggregateDiscoveredResourceResponseBodyDiscoveredResourceDetail {
	s.Tags = &v
	return s
}

type GetAggregateDiscoveredResourceResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetAggregateDiscoveredResourceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetAggregateDiscoveredResourceResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAggregateDiscoveredResourceResponse) GoString() string {
	return s.String()
}

func (s *GetAggregateDiscoveredResourceResponse) SetHeaders(v map[string]*string) *GetAggregateDiscoveredResourceResponse {
	s.Headers = v
	return s
}

func (s *GetAggregateDiscoveredResourceResponse) SetStatusCode(v int32) *GetAggregateDiscoveredResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAggregateDiscoveredResourceResponse) SetBody(v *GetAggregateDiscoveredResourceResponseBody) *GetAggregateDiscoveredResourceResponse {
	s.Body = v
	return s
}

type GetDiscoveredResourceCountsRequest struct {
	// The dimension by which resources are grouped. Valid values:
	//
	// *   ResourceType: specifies that resources are grouped by type.
	// *   Region: specifies that resources are grouped by region.
	GroupByKey *string `json:"GroupByKey,omitempty" xml:"GroupByKey,omitempty"`
	// This parameter has been removed from 00:00:00, June 30, 2021. Account group-related operations are provided as an alternative after 00:00:00, May 30, 2021. If you are using this parameter, we recommend that you switch to account group-related operations after 00:00:00, May 30, 2021. For more information about account groups, see [Overview](~~211534~~).
	MemberId *int64 `json:"MemberId,omitempty" xml:"MemberId,omitempty"`
	// This parameter has been removed from 00:00:00, June 30, 2021. Account group-related operations are provided as an alternative after 00:00:00, May 30, 2021. If you are using this parameter, we recommend that you switch to account group-related operations after 00:00:00, May 30, 2021. For more information about account groups, see [Overview](~~211534~~).
	MultiAccount *bool `json:"MultiAccount,omitempty" xml:"MultiAccount,omitempty"`
}

func (s GetDiscoveredResourceCountsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDiscoveredResourceCountsRequest) GoString() string {
	return s.String()
}

func (s *GetDiscoveredResourceCountsRequest) SetGroupByKey(v string) *GetDiscoveredResourceCountsRequest {
	s.GroupByKey = &v
	return s
}

func (s *GetDiscoveredResourceCountsRequest) SetMemberId(v int64) *GetDiscoveredResourceCountsRequest {
	s.MemberId = &v
	return s
}

func (s *GetDiscoveredResourceCountsRequest) SetMultiAccount(v bool) *GetDiscoveredResourceCountsRequest {
	s.MultiAccount = &v
	return s
}

type GetDiscoveredResourceCountsResponseBody struct {
	// The returned information about resources.
	GroupedResourceCounts *GetDiscoveredResourceCountsResponseBodyGroupedResourceCounts `json:"GroupedResourceCounts,omitempty" xml:"GroupedResourceCounts,omitempty" type:"Struct"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetDiscoveredResourceCountsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDiscoveredResourceCountsResponseBody) GoString() string {
	return s.String()
}

func (s *GetDiscoveredResourceCountsResponseBody) SetGroupedResourceCounts(v *GetDiscoveredResourceCountsResponseBodyGroupedResourceCounts) *GetDiscoveredResourceCountsResponseBody {
	s.GroupedResourceCounts = v
	return s
}

func (s *GetDiscoveredResourceCountsResponseBody) SetRequestId(v string) *GetDiscoveredResourceCountsResponseBody {
	s.RequestId = &v
	return s
}

type GetDiscoveredResourceCountsResponseBodyGroupedResourceCounts struct {
	// The dimension by which resources are grouped.
	GroupByKey *string `json:"GroupByKey,omitempty" xml:"GroupByKey,omitempty"`
	// The information about resources in a specific group.
	GroupedResourceCountList []*GetDiscoveredResourceCountsResponseBodyGroupedResourceCountsGroupedResourceCountList `json:"GroupedResourceCountList,omitempty" xml:"GroupedResourceCountList,omitempty" type:"Repeated"`
}

func (s GetDiscoveredResourceCountsResponseBodyGroupedResourceCounts) String() string {
	return tea.Prettify(s)
}

func (s GetDiscoveredResourceCountsResponseBodyGroupedResourceCounts) GoString() string {
	return s.String()
}

func (s *GetDiscoveredResourceCountsResponseBodyGroupedResourceCounts) SetGroupByKey(v string) *GetDiscoveredResourceCountsResponseBodyGroupedResourceCounts {
	s.GroupByKey = &v
	return s
}

func (s *GetDiscoveredResourceCountsResponseBodyGroupedResourceCounts) SetGroupedResourceCountList(v []*GetDiscoveredResourceCountsResponseBodyGroupedResourceCountsGroupedResourceCountList) *GetDiscoveredResourceCountsResponseBodyGroupedResourceCounts {
	s.GroupedResourceCountList = v
	return s
}

type GetDiscoveredResourceCountsResponseBodyGroupedResourceCountsGroupedResourceCountList struct {
	// The name of the resource group.
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The number of resources in the group.
	ResourceCount *int64 `json:"ResourceCount,omitempty" xml:"ResourceCount,omitempty"`
}

func (s GetDiscoveredResourceCountsResponseBodyGroupedResourceCountsGroupedResourceCountList) String() string {
	return tea.Prettify(s)
}

func (s GetDiscoveredResourceCountsResponseBodyGroupedResourceCountsGroupedResourceCountList) GoString() string {
	return s.String()
}

func (s *GetDiscoveredResourceCountsResponseBodyGroupedResourceCountsGroupedResourceCountList) SetGroupName(v string) *GetDiscoveredResourceCountsResponseBodyGroupedResourceCountsGroupedResourceCountList {
	s.GroupName = &v
	return s
}

func (s *GetDiscoveredResourceCountsResponseBodyGroupedResourceCountsGroupedResourceCountList) SetResourceCount(v int64) *GetDiscoveredResourceCountsResponseBodyGroupedResourceCountsGroupedResourceCountList {
	s.ResourceCount = &v
	return s
}

type GetDiscoveredResourceCountsResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetDiscoveredResourceCountsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetDiscoveredResourceCountsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDiscoveredResourceCountsResponse) GoString() string {
	return s.String()
}

func (s *GetDiscoveredResourceCountsResponse) SetHeaders(v map[string]*string) *GetDiscoveredResourceCountsResponse {
	s.Headers = v
	return s
}

func (s *GetDiscoveredResourceCountsResponse) SetStatusCode(v int32) *GetDiscoveredResourceCountsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDiscoveredResourceCountsResponse) SetBody(v *GetDiscoveredResourceCountsResponseBody) *GetDiscoveredResourceCountsResponse {
	s.Body = v
	return s
}

type GetDiscoveredResourceSummaryRequest struct {
	// This parameter is scheduled to be removed before 00:00:00, June 30, 2021. Account group-related APIs will be provided as an alternative before 00:00:00, May 30, 2021. If you are using this parameter, we recommend that you switch to account group-related APIs after 00:00:00, May 30, 2021. For more information, see [Account groups](~~211534~~).
	MemberId *int64 `json:"MemberId,omitempty" xml:"MemberId,omitempty"`
	// This parameter is scheduled to be removed before 00:00:00, June 30, 2021. Account group-related APIs will be provided as an alternative before 00:00:00, May 30, 2021. If you are using this parameter, we recommend that you switch to account group-related APIs after 00:00:00, May 30, 2021. For more information, see [Account groups](~~211534~~).
	MultiAccount *bool `json:"MultiAccount,omitempty" xml:"MultiAccount,omitempty"`
}

func (s GetDiscoveredResourceSummaryRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDiscoveredResourceSummaryRequest) GoString() string {
	return s.String()
}

func (s *GetDiscoveredResourceSummaryRequest) SetMemberId(v int64) *GetDiscoveredResourceSummaryRequest {
	s.MemberId = &v
	return s
}

func (s *GetDiscoveredResourceSummaryRequest) SetMultiAccount(v bool) *GetDiscoveredResourceSummaryRequest {
	s.MultiAccount = &v
	return s
}

type GetDiscoveredResourceSummaryResponseBody struct {
	// The statistics of the resources monitored by Cloud Config.
	DiscoveredResourceSummary *GetDiscoveredResourceSummaryResponseBodyDiscoveredResourceSummary `json:"DiscoveredResourceSummary,omitempty" xml:"DiscoveredResourceSummary,omitempty" type:"Struct"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetDiscoveredResourceSummaryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDiscoveredResourceSummaryResponseBody) GoString() string {
	return s.String()
}

func (s *GetDiscoveredResourceSummaryResponseBody) SetDiscoveredResourceSummary(v *GetDiscoveredResourceSummaryResponseBodyDiscoveredResourceSummary) *GetDiscoveredResourceSummaryResponseBody {
	s.DiscoveredResourceSummary = v
	return s
}

func (s *GetDiscoveredResourceSummaryResponseBody) SetRequestId(v string) *GetDiscoveredResourceSummaryResponseBody {
	s.RequestId = &v
	return s
}

type GetDiscoveredResourceSummaryResponseBodyDiscoveredResourceSummary struct {
	// The number of the regions involved in the monitoring.
	RegionCount *int32 `json:"RegionCount,omitempty" xml:"RegionCount,omitempty"`
	// The number of the resources monitored by Cloud Config.
	ResourceCount *int32 `json:"ResourceCount,omitempty" xml:"ResourceCount,omitempty"`
	// The number of the resource types monitored by Cloud Config.
	ResourceTypeCount *int32 `json:"ResourceTypeCount,omitempty" xml:"ResourceTypeCount,omitempty"`
}

func (s GetDiscoveredResourceSummaryResponseBodyDiscoveredResourceSummary) String() string {
	return tea.Prettify(s)
}

func (s GetDiscoveredResourceSummaryResponseBodyDiscoveredResourceSummary) GoString() string {
	return s.String()
}

func (s *GetDiscoveredResourceSummaryResponseBodyDiscoveredResourceSummary) SetRegionCount(v int32) *GetDiscoveredResourceSummaryResponseBodyDiscoveredResourceSummary {
	s.RegionCount = &v
	return s
}

func (s *GetDiscoveredResourceSummaryResponseBodyDiscoveredResourceSummary) SetResourceCount(v int32) *GetDiscoveredResourceSummaryResponseBodyDiscoveredResourceSummary {
	s.ResourceCount = &v
	return s
}

func (s *GetDiscoveredResourceSummaryResponseBodyDiscoveredResourceSummary) SetResourceTypeCount(v int32) *GetDiscoveredResourceSummaryResponseBodyDiscoveredResourceSummary {
	s.ResourceTypeCount = &v
	return s
}

type GetDiscoveredResourceSummaryResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetDiscoveredResourceSummaryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetDiscoveredResourceSummaryResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDiscoveredResourceSummaryResponse) GoString() string {
	return s.String()
}

func (s *GetDiscoveredResourceSummaryResponse) SetHeaders(v map[string]*string) *GetDiscoveredResourceSummaryResponse {
	s.Headers = v
	return s
}

func (s *GetDiscoveredResourceSummaryResponse) SetStatusCode(v int32) *GetDiscoveredResourceSummaryResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDiscoveredResourceSummaryResponse) SetBody(v *GetDiscoveredResourceSummaryResponseBody) *GetDiscoveredResourceSummaryResponse {
	s.Body = v
	return s
}

type GetResourceComplianceTimelineRequest struct {
	// The timestamp that specifies the end of the time range to query. By default, the value is the time when the GetResourceConfigurationTimeline operation is called.
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The number of entries to return on each page. Valid values: 1 to 100. Default value: 10.
	Limit *int32 `json:"Limit,omitempty" xml:"Limit,omitempty"`
	// This parameter is scheduled to be removed before 00:00:00, June 30, 2021. Account group-related APIs will be provided as an alternative before 00:00:00, May 30, 2021. If you are using this parameter, we recommend that you switch to account group-related APIs after 00:00:00, May 30, 2021. For more information, see [Account groups](~~211534~~).
	MemberId *string `json:"MemberId,omitempty" xml:"MemberId,omitempty"`
	// This parameter is scheduled to be removed before 00:00:00, June 30, 2021. Account group-related APIs will be provided as an alternative before 00:00:00, May 30, 2021. If you are using this parameter, we recommend that you switch to account group-related APIs after 00:00:00, May 30, 2021. For more information, see [Account groups](~~211534~~).
	MultiAccount *bool `json:"MultiAccount,omitempty" xml:"MultiAccount,omitempty"`
	// The token that is used to start the next query.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the region.
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The ID of the resource.
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The type of the resource.
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The timestamp that specifies the beginning of the time range to query. By default, Cloud Config retrieves the configuration changes in the last 30 days for the specified resource.
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s GetResourceComplianceTimelineRequest) String() string {
	return tea.Prettify(s)
}

func (s GetResourceComplianceTimelineRequest) GoString() string {
	return s.String()
}

func (s *GetResourceComplianceTimelineRequest) SetEndTime(v int64) *GetResourceComplianceTimelineRequest {
	s.EndTime = &v
	return s
}

func (s *GetResourceComplianceTimelineRequest) SetLimit(v int32) *GetResourceComplianceTimelineRequest {
	s.Limit = &v
	return s
}

func (s *GetResourceComplianceTimelineRequest) SetMemberId(v string) *GetResourceComplianceTimelineRequest {
	s.MemberId = &v
	return s
}

func (s *GetResourceComplianceTimelineRequest) SetMultiAccount(v bool) *GetResourceComplianceTimelineRequest {
	s.MultiAccount = &v
	return s
}

func (s *GetResourceComplianceTimelineRequest) SetNextToken(v string) *GetResourceComplianceTimelineRequest {
	s.NextToken = &v
	return s
}

func (s *GetResourceComplianceTimelineRequest) SetRegion(v string) *GetResourceComplianceTimelineRequest {
	s.Region = &v
	return s
}

func (s *GetResourceComplianceTimelineRequest) SetResourceId(v string) *GetResourceComplianceTimelineRequest {
	s.ResourceId = &v
	return s
}

func (s *GetResourceComplianceTimelineRequest) SetResourceType(v string) *GetResourceComplianceTimelineRequest {
	s.ResourceType = &v
	return s
}

func (s *GetResourceComplianceTimelineRequest) SetStartTime(v int64) *GetResourceComplianceTimelineRequest {
	s.StartTime = &v
	return s
}

type GetResourceComplianceTimelineResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the compliance timeline.
	ResourceComplianceTimeline *GetResourceComplianceTimelineResponseBodyResourceComplianceTimeline `json:"ResourceComplianceTimeline,omitempty" xml:"ResourceComplianceTimeline,omitempty" type:"Struct"`
}

func (s GetResourceComplianceTimelineResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetResourceComplianceTimelineResponseBody) GoString() string {
	return s.String()
}

func (s *GetResourceComplianceTimelineResponseBody) SetRequestId(v string) *GetResourceComplianceTimelineResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetResourceComplianceTimelineResponseBody) SetResourceComplianceTimeline(v *GetResourceComplianceTimelineResponseBodyResourceComplianceTimeline) *GetResourceComplianceTimelineResponseBody {
	s.ResourceComplianceTimeline = v
	return s
}

type GetResourceComplianceTimelineResponseBodyResourceComplianceTimeline struct {
	// The compliance evaluations in the compliance timeline.
	ComplianceList []*GetResourceComplianceTimelineResponseBodyResourceComplianceTimelineComplianceList `json:"ComplianceList,omitempty" xml:"ComplianceList,omitempty" type:"Repeated"`
	// The number of entries returned on each page. Valid values: 1 to 100.
	Limit *int32 `json:"Limit,omitempty" xml:"Limit,omitempty"`
	// The token that was returned for the next query.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The total number of the compliance evaluations for the specified resource.
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s GetResourceComplianceTimelineResponseBodyResourceComplianceTimeline) String() string {
	return tea.Prettify(s)
}

func (s GetResourceComplianceTimelineResponseBodyResourceComplianceTimeline) GoString() string {
	return s.String()
}

func (s *GetResourceComplianceTimelineResponseBodyResourceComplianceTimeline) SetComplianceList(v []*GetResourceComplianceTimelineResponseBodyResourceComplianceTimelineComplianceList) *GetResourceComplianceTimelineResponseBodyResourceComplianceTimeline {
	s.ComplianceList = v
	return s
}

func (s *GetResourceComplianceTimelineResponseBodyResourceComplianceTimeline) SetLimit(v int32) *GetResourceComplianceTimelineResponseBodyResourceComplianceTimeline {
	s.Limit = &v
	return s
}

func (s *GetResourceComplianceTimelineResponseBodyResourceComplianceTimeline) SetNextToken(v string) *GetResourceComplianceTimelineResponseBodyResourceComplianceTimeline {
	s.NextToken = &v
	return s
}

func (s *GetResourceComplianceTimelineResponseBodyResourceComplianceTimeline) SetTotalCount(v int64) *GetResourceComplianceTimelineResponseBodyResourceComplianceTimeline {
	s.TotalCount = &v
	return s
}

type GetResourceComplianceTimelineResponseBodyResourceComplianceTimelineComplianceList struct {
	// The ID of your Alibaba Cloud account.
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The ID of the zone.
	AvailabilityZone *string `json:"AvailabilityZone,omitempty" xml:"AvailabilityZone,omitempty"`
	// The timestamp when the compliance evaluation occurred.
	CaptureTime *int64 `json:"CaptureTime,omitempty" xml:"CaptureTime,omitempty"`
	// The information about the rules triggered for the resource and the compliance evaluation results.
	Configuration *string `json:"Configuration,omitempty" xml:"Configuration,omitempty"`
	// The information about the changes to the compliance evaluation results, including the information about the triggered rules.
	ConfigurationDiff *string `json:"ConfigurationDiff,omitempty" xml:"ConfigurationDiff,omitempty"`
	// The ID of the region.
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The timestamp when the resource was created.
	ResourceCreateTime *int64 `json:"ResourceCreateTime,omitempty" xml:"ResourceCreateTime,omitempty"`
	// The ID of the resource.
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The name of the resource.
	ResourceName *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	// The status of the resource. The parameter value varies with the resource type and may be left empty. Examples:
	//
	// *   If ResourceType is set to ACS::ECS::Instance, the resource is an Elastic Compute Service (ECS) instance which has a specific state. In this case, valid values of this parameter include Running and Stopped.
	// *   If ResourceType is set to ACS::OSS::Bucket, the resource is an Object Storage Service (OSS) bucket which does not have a specific state. In this case, this parameter is left empty.
	ResourceStatus *string `json:"ResourceStatus,omitempty" xml:"ResourceStatus,omitempty"`
	// The type of the resource.
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The tags of the resource.
	Tags *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
}

func (s GetResourceComplianceTimelineResponseBodyResourceComplianceTimelineComplianceList) String() string {
	return tea.Prettify(s)
}

func (s GetResourceComplianceTimelineResponseBodyResourceComplianceTimelineComplianceList) GoString() string {
	return s.String()
}

func (s *GetResourceComplianceTimelineResponseBodyResourceComplianceTimelineComplianceList) SetAccountId(v string) *GetResourceComplianceTimelineResponseBodyResourceComplianceTimelineComplianceList {
	s.AccountId = &v
	return s
}

func (s *GetResourceComplianceTimelineResponseBodyResourceComplianceTimelineComplianceList) SetAvailabilityZone(v string) *GetResourceComplianceTimelineResponseBodyResourceComplianceTimelineComplianceList {
	s.AvailabilityZone = &v
	return s
}

func (s *GetResourceComplianceTimelineResponseBodyResourceComplianceTimelineComplianceList) SetCaptureTime(v int64) *GetResourceComplianceTimelineResponseBodyResourceComplianceTimelineComplianceList {
	s.CaptureTime = &v
	return s
}

func (s *GetResourceComplianceTimelineResponseBodyResourceComplianceTimelineComplianceList) SetConfiguration(v string) *GetResourceComplianceTimelineResponseBodyResourceComplianceTimelineComplianceList {
	s.Configuration = &v
	return s
}

func (s *GetResourceComplianceTimelineResponseBodyResourceComplianceTimelineComplianceList) SetConfigurationDiff(v string) *GetResourceComplianceTimelineResponseBodyResourceComplianceTimelineComplianceList {
	s.ConfigurationDiff = &v
	return s
}

func (s *GetResourceComplianceTimelineResponseBodyResourceComplianceTimelineComplianceList) SetRegion(v string) *GetResourceComplianceTimelineResponseBodyResourceComplianceTimelineComplianceList {
	s.Region = &v
	return s
}

func (s *GetResourceComplianceTimelineResponseBodyResourceComplianceTimelineComplianceList) SetResourceCreateTime(v int64) *GetResourceComplianceTimelineResponseBodyResourceComplianceTimelineComplianceList {
	s.ResourceCreateTime = &v
	return s
}

func (s *GetResourceComplianceTimelineResponseBodyResourceComplianceTimelineComplianceList) SetResourceId(v string) *GetResourceComplianceTimelineResponseBodyResourceComplianceTimelineComplianceList {
	s.ResourceId = &v
	return s
}

func (s *GetResourceComplianceTimelineResponseBodyResourceComplianceTimelineComplianceList) SetResourceName(v string) *GetResourceComplianceTimelineResponseBodyResourceComplianceTimelineComplianceList {
	s.ResourceName = &v
	return s
}

func (s *GetResourceComplianceTimelineResponseBodyResourceComplianceTimelineComplianceList) SetResourceStatus(v string) *GetResourceComplianceTimelineResponseBodyResourceComplianceTimelineComplianceList {
	s.ResourceStatus = &v
	return s
}

func (s *GetResourceComplianceTimelineResponseBodyResourceComplianceTimelineComplianceList) SetResourceType(v string) *GetResourceComplianceTimelineResponseBodyResourceComplianceTimelineComplianceList {
	s.ResourceType = &v
	return s
}

func (s *GetResourceComplianceTimelineResponseBodyResourceComplianceTimelineComplianceList) SetTags(v string) *GetResourceComplianceTimelineResponseBodyResourceComplianceTimelineComplianceList {
	s.Tags = &v
	return s
}

type GetResourceComplianceTimelineResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetResourceComplianceTimelineResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetResourceComplianceTimelineResponse) String() string {
	return tea.Prettify(s)
}

func (s GetResourceComplianceTimelineResponse) GoString() string {
	return s.String()
}

func (s *GetResourceComplianceTimelineResponse) SetHeaders(v map[string]*string) *GetResourceComplianceTimelineResponse {
	s.Headers = v
	return s
}

func (s *GetResourceComplianceTimelineResponse) SetStatusCode(v int32) *GetResourceComplianceTimelineResponse {
	s.StatusCode = &v
	return s
}

func (s *GetResourceComplianceTimelineResponse) SetBody(v *GetResourceComplianceTimelineResponseBody) *GetResourceComplianceTimelineResponse {
	s.Body = v
	return s
}

type GetResourceConfigurationTimelineRequest struct {
	// The timestamp that specifies the end of the time range to query. By default, the value is the time when the GetResourceConfigurationTimeline operation is called.
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The number of entries to return on each page. Valid values: 1 to 100. Default value: 10.
	Limit *int32 `json:"Limit,omitempty" xml:"Limit,omitempty"`
	// This parameter is scheduled to be removed before 00:00:00, June 30, 2021. Account group-related APIs will be provided as an alternative before 00:00:00, May 30, 2021. If you are using this parameter, we recommend that you switch to account group-related APIs after 00:00:00, May 30, 2021. For more information, see [Account groups](~~211534~~).
	MemberId *int64 `json:"MemberId,omitempty" xml:"MemberId,omitempty"`
	// This parameter is scheduled to be removed before 00:00:00, June 30, 2021. Account group-related APIs will be provided as an alternative before 00:00:00, May 30, 2021. If you are using this parameter, we recommend that you switch to account group-related APIs after 00:00:00, May 30, 2021. For more information, see [Account groups](~~211534~~).
	MultiAccount *bool `json:"MultiAccount,omitempty" xml:"MultiAccount,omitempty"`
	// The token that is used to start the next query.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the region.
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The ID of the resource.
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The type of the resource.
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The timestamp that specifies the beginning of the time range to query. By default, Cloud Config retrieves the configuration changes in the last 30 days for the specified resource.
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s GetResourceConfigurationTimelineRequest) String() string {
	return tea.Prettify(s)
}

func (s GetResourceConfigurationTimelineRequest) GoString() string {
	return s.String()
}

func (s *GetResourceConfigurationTimelineRequest) SetEndTime(v int64) *GetResourceConfigurationTimelineRequest {
	s.EndTime = &v
	return s
}

func (s *GetResourceConfigurationTimelineRequest) SetLimit(v int32) *GetResourceConfigurationTimelineRequest {
	s.Limit = &v
	return s
}

func (s *GetResourceConfigurationTimelineRequest) SetMemberId(v int64) *GetResourceConfigurationTimelineRequest {
	s.MemberId = &v
	return s
}

func (s *GetResourceConfigurationTimelineRequest) SetMultiAccount(v bool) *GetResourceConfigurationTimelineRequest {
	s.MultiAccount = &v
	return s
}

func (s *GetResourceConfigurationTimelineRequest) SetNextToken(v string) *GetResourceConfigurationTimelineRequest {
	s.NextToken = &v
	return s
}

func (s *GetResourceConfigurationTimelineRequest) SetRegion(v string) *GetResourceConfigurationTimelineRequest {
	s.Region = &v
	return s
}

func (s *GetResourceConfigurationTimelineRequest) SetResourceId(v string) *GetResourceConfigurationTimelineRequest {
	s.ResourceId = &v
	return s
}

func (s *GetResourceConfigurationTimelineRequest) SetResourceType(v string) *GetResourceConfigurationTimelineRequest {
	s.ResourceType = &v
	return s
}

func (s *GetResourceConfigurationTimelineRequest) SetStartTime(v int64) *GetResourceConfigurationTimelineRequest {
	s.StartTime = &v
	return s
}

type GetResourceConfigurationTimelineResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information of the configuration timeline.
	ResourceConfigurationTimeline *GetResourceConfigurationTimelineResponseBodyResourceConfigurationTimeline `json:"ResourceConfigurationTimeline,omitempty" xml:"ResourceConfigurationTimeline,omitempty" type:"Struct"`
}

func (s GetResourceConfigurationTimelineResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetResourceConfigurationTimelineResponseBody) GoString() string {
	return s.String()
}

func (s *GetResourceConfigurationTimelineResponseBody) SetRequestId(v string) *GetResourceConfigurationTimelineResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetResourceConfigurationTimelineResponseBody) SetResourceConfigurationTimeline(v *GetResourceConfigurationTimelineResponseBodyResourceConfigurationTimeline) *GetResourceConfigurationTimelineResponseBody {
	s.ResourceConfigurationTimeline = v
	return s
}

type GetResourceConfigurationTimelineResponseBodyResourceConfigurationTimeline struct {
	// The configuration changes of the configuration timeline.
	ConfigurationList []*GetResourceConfigurationTimelineResponseBodyResourceConfigurationTimelineConfigurationList `json:"ConfigurationList,omitempty" xml:"ConfigurationList,omitempty" type:"Repeated"`
	// The number of entries returned on each page. Valid values: 1 to 100.
	Limit *int32 `json:"Limit,omitempty" xml:"Limit,omitempty"`
	// The token that was returned for the next query.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The total number of configuration changes for the specified resource.
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s GetResourceConfigurationTimelineResponseBodyResourceConfigurationTimeline) String() string {
	return tea.Prettify(s)
}

func (s GetResourceConfigurationTimelineResponseBodyResourceConfigurationTimeline) GoString() string {
	return s.String()
}

func (s *GetResourceConfigurationTimelineResponseBodyResourceConfigurationTimeline) SetConfigurationList(v []*GetResourceConfigurationTimelineResponseBodyResourceConfigurationTimelineConfigurationList) *GetResourceConfigurationTimelineResponseBodyResourceConfigurationTimeline {
	s.ConfigurationList = v
	return s
}

func (s *GetResourceConfigurationTimelineResponseBodyResourceConfigurationTimeline) SetLimit(v int32) *GetResourceConfigurationTimelineResponseBodyResourceConfigurationTimeline {
	s.Limit = &v
	return s
}

func (s *GetResourceConfigurationTimelineResponseBodyResourceConfigurationTimeline) SetNextToken(v string) *GetResourceConfigurationTimelineResponseBodyResourceConfigurationTimeline {
	s.NextToken = &v
	return s
}

func (s *GetResourceConfigurationTimelineResponseBodyResourceConfigurationTimeline) SetTotalCount(v int64) *GetResourceConfigurationTimelineResponseBodyResourceConfigurationTimeline {
	s.TotalCount = &v
	return s
}

type GetResourceConfigurationTimelineResponseBodyResourceConfigurationTimelineConfigurationList struct {
	// The ID of your Alibaba Cloud account.
	AccountId *int64 `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The ID of the zone.
	AvailabilityZone *string `json:"AvailabilityZone,omitempty" xml:"AvailabilityZone,omitempty"`
	// The timestamp when the configuration change occurred.
	CaptureTime *string `json:"CaptureTime,omitempty" xml:"CaptureTime,omitempty"`
	// The information of the changes to the configuration of the resource.
	ConfigurationDiff *string `json:"ConfigurationDiff,omitempty" xml:"ConfigurationDiff,omitempty"`
	// The ID of the region.
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The time when the resource was created.
	ResourceCreateTime *string `json:"ResourceCreateTime,omitempty" xml:"ResourceCreateTime,omitempty"`
	// The type of the resource change event. Valid values:
	//
	// *   DISCOVERED: A resource was created.
	// *   MODIFY: A resource was modified.
	// *   REMOVE: A resource is deleted.
	ResourceEventType *string `json:"ResourceEventType,omitempty" xml:"ResourceEventType,omitempty"`
	// The ID of the resource.
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The name of the resource.
	ResourceName *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	// The type of the resource.
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The tags of the resource.
	Tags *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
}

func (s GetResourceConfigurationTimelineResponseBodyResourceConfigurationTimelineConfigurationList) String() string {
	return tea.Prettify(s)
}

func (s GetResourceConfigurationTimelineResponseBodyResourceConfigurationTimelineConfigurationList) GoString() string {
	return s.String()
}

func (s *GetResourceConfigurationTimelineResponseBodyResourceConfigurationTimelineConfigurationList) SetAccountId(v int64) *GetResourceConfigurationTimelineResponseBodyResourceConfigurationTimelineConfigurationList {
	s.AccountId = &v
	return s
}

func (s *GetResourceConfigurationTimelineResponseBodyResourceConfigurationTimelineConfigurationList) SetAvailabilityZone(v string) *GetResourceConfigurationTimelineResponseBodyResourceConfigurationTimelineConfigurationList {
	s.AvailabilityZone = &v
	return s
}

func (s *GetResourceConfigurationTimelineResponseBodyResourceConfigurationTimelineConfigurationList) SetCaptureTime(v string) *GetResourceConfigurationTimelineResponseBodyResourceConfigurationTimelineConfigurationList {
	s.CaptureTime = &v
	return s
}

func (s *GetResourceConfigurationTimelineResponseBodyResourceConfigurationTimelineConfigurationList) SetConfigurationDiff(v string) *GetResourceConfigurationTimelineResponseBodyResourceConfigurationTimelineConfigurationList {
	s.ConfigurationDiff = &v
	return s
}

func (s *GetResourceConfigurationTimelineResponseBodyResourceConfigurationTimelineConfigurationList) SetRegion(v string) *GetResourceConfigurationTimelineResponseBodyResourceConfigurationTimelineConfigurationList {
	s.Region = &v
	return s
}

func (s *GetResourceConfigurationTimelineResponseBodyResourceConfigurationTimelineConfigurationList) SetResourceCreateTime(v string) *GetResourceConfigurationTimelineResponseBodyResourceConfigurationTimelineConfigurationList {
	s.ResourceCreateTime = &v
	return s
}

func (s *GetResourceConfigurationTimelineResponseBodyResourceConfigurationTimelineConfigurationList) SetResourceEventType(v string) *GetResourceConfigurationTimelineResponseBodyResourceConfigurationTimelineConfigurationList {
	s.ResourceEventType = &v
	return s
}

func (s *GetResourceConfigurationTimelineResponseBodyResourceConfigurationTimelineConfigurationList) SetResourceId(v string) *GetResourceConfigurationTimelineResponseBodyResourceConfigurationTimelineConfigurationList {
	s.ResourceId = &v
	return s
}

func (s *GetResourceConfigurationTimelineResponseBodyResourceConfigurationTimelineConfigurationList) SetResourceName(v string) *GetResourceConfigurationTimelineResponseBodyResourceConfigurationTimelineConfigurationList {
	s.ResourceName = &v
	return s
}

func (s *GetResourceConfigurationTimelineResponseBodyResourceConfigurationTimelineConfigurationList) SetResourceType(v string) *GetResourceConfigurationTimelineResponseBodyResourceConfigurationTimelineConfigurationList {
	s.ResourceType = &v
	return s
}

func (s *GetResourceConfigurationTimelineResponseBodyResourceConfigurationTimelineConfigurationList) SetTags(v string) *GetResourceConfigurationTimelineResponseBodyResourceConfigurationTimelineConfigurationList {
	s.Tags = &v
	return s
}

type GetResourceConfigurationTimelineResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetResourceConfigurationTimelineResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetResourceConfigurationTimelineResponse) String() string {
	return tea.Prettify(s)
}

func (s GetResourceConfigurationTimelineResponse) GoString() string {
	return s.String()
}

func (s *GetResourceConfigurationTimelineResponse) SetHeaders(v map[string]*string) *GetResourceConfigurationTimelineResponse {
	s.Headers = v
	return s
}

func (s *GetResourceConfigurationTimelineResponse) SetStatusCode(v int32) *GetResourceConfigurationTimelineResponse {
	s.StatusCode = &v
	return s
}

func (s *GetResourceConfigurationTimelineResponse) SetBody(v *GetResourceConfigurationTimelineResponseBody) *GetResourceConfigurationTimelineResponse {
	s.Body = v
	return s
}

type GetSupportedResourceTypesResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The types of resources that are supported by Cloud Config.
	ResourceTypes []*string `json:"ResourceTypes,omitempty" xml:"ResourceTypes,omitempty" type:"Repeated"`
}

func (s GetSupportedResourceTypesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSupportedResourceTypesResponseBody) GoString() string {
	return s.String()
}

func (s *GetSupportedResourceTypesResponseBody) SetRequestId(v string) *GetSupportedResourceTypesResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSupportedResourceTypesResponseBody) SetResourceTypes(v []*string) *GetSupportedResourceTypesResponseBody {
	s.ResourceTypes = v
	return s
}

type GetSupportedResourceTypesResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetSupportedResourceTypesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetSupportedResourceTypesResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSupportedResourceTypesResponse) GoString() string {
	return s.String()
}

func (s *GetSupportedResourceTypesResponse) SetHeaders(v map[string]*string) *GetSupportedResourceTypesResponse {
	s.Headers = v
	return s
}

func (s *GetSupportedResourceTypesResponse) SetStatusCode(v int32) *GetSupportedResourceTypesResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSupportedResourceTypesResponse) SetBody(v *GetSupportedResourceTypesResponseBody) *GetSupportedResourceTypesResponse {
	s.Body = v
	return s
}

type ListAggregateDiscoveredResourcesRequest struct {
	// The ID of the account group.
	//
	// For more information about how to obtain the ID of an account group, see [ListAggregators](~~255797~~).
	AggregatorId *string `json:"AggregatorId,omitempty" xml:"AggregatorId,omitempty"`
	// The number of the page to return.
	//
	// Pages start from page 1. Default value: 1.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page.
	//
	// Valid values: 1 to 100. A minimum of 1 entry can be returned on each page. Default value: 10.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the region in which the resource resides. Separate multiple region IDs with commas (,).
	//
	// For more information about how to obtain the region ID of a resource, see [ListAggregateDiscoveredResources](~~265983~~).
	Regions *string `json:"Regions,omitempty" xml:"Regions,omitempty"`
	// The status of the resource. Valid values:
	//
	// *   0: The resource is deleted.
	// *   1: The resource is retained. This is the default value.
	ResourceDeleted *int32 `json:"ResourceDeleted,omitempty" xml:"ResourceDeleted,omitempty"`
	// The ID of the resource.
	ResourceId      *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The type of the resource. Separate multiple resource types with commas (,).
	ResourceTypes *string `json:"ResourceTypes,omitempty" xml:"ResourceTypes,omitempty"`
}

func (s ListAggregateDiscoveredResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAggregateDiscoveredResourcesRequest) GoString() string {
	return s.String()
}

func (s *ListAggregateDiscoveredResourcesRequest) SetAggregatorId(v string) *ListAggregateDiscoveredResourcesRequest {
	s.AggregatorId = &v
	return s
}

func (s *ListAggregateDiscoveredResourcesRequest) SetPageNumber(v int32) *ListAggregateDiscoveredResourcesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListAggregateDiscoveredResourcesRequest) SetPageSize(v int32) *ListAggregateDiscoveredResourcesRequest {
	s.PageSize = &v
	return s
}

func (s *ListAggregateDiscoveredResourcesRequest) SetRegions(v string) *ListAggregateDiscoveredResourcesRequest {
	s.Regions = &v
	return s
}

func (s *ListAggregateDiscoveredResourcesRequest) SetResourceDeleted(v int32) *ListAggregateDiscoveredResourcesRequest {
	s.ResourceDeleted = &v
	return s
}

func (s *ListAggregateDiscoveredResourcesRequest) SetResourceId(v string) *ListAggregateDiscoveredResourcesRequest {
	s.ResourceId = &v
	return s
}

func (s *ListAggregateDiscoveredResourcesRequest) SetResourceOwnerId(v int64) *ListAggregateDiscoveredResourcesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListAggregateDiscoveredResourcesRequest) SetResourceTypes(v string) *ListAggregateDiscoveredResourcesRequest {
	s.ResourceTypes = &v
	return s
}

type ListAggregateDiscoveredResourcesResponseBody struct {
	// The information about the returned resources.
	DiscoveredResourceProfiles *ListAggregateDiscoveredResourcesResponseBodyDiscoveredResourceProfiles `json:"DiscoveredResourceProfiles,omitempty" xml:"DiscoveredResourceProfiles,omitempty" type:"Struct"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListAggregateDiscoveredResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAggregateDiscoveredResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *ListAggregateDiscoveredResourcesResponseBody) SetDiscoveredResourceProfiles(v *ListAggregateDiscoveredResourcesResponseBodyDiscoveredResourceProfiles) *ListAggregateDiscoveredResourcesResponseBody {
	s.DiscoveredResourceProfiles = v
	return s
}

func (s *ListAggregateDiscoveredResourcesResponseBody) SetRequestId(v string) *ListAggregateDiscoveredResourcesResponseBody {
	s.RequestId = &v
	return s
}

type ListAggregateDiscoveredResourcesResponseBodyDiscoveredResourceProfiles struct {
	// The information about the resource.
	DiscoveredResourceProfileList []*ListAggregateDiscoveredResourcesResponseBodyDiscoveredResourceProfilesDiscoveredResourceProfileList `json:"DiscoveredResourceProfileList,omitempty" xml:"DiscoveredResourceProfileList,omitempty" type:"Repeated"`
	// The page number of the returned page. Pages start from page 1.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page. Valid values: 1 to 100.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of resources.
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListAggregateDiscoveredResourcesResponseBodyDiscoveredResourceProfiles) String() string {
	return tea.Prettify(s)
}

func (s ListAggregateDiscoveredResourcesResponseBodyDiscoveredResourceProfiles) GoString() string {
	return s.String()
}

func (s *ListAggregateDiscoveredResourcesResponseBodyDiscoveredResourceProfiles) SetDiscoveredResourceProfileList(v []*ListAggregateDiscoveredResourcesResponseBodyDiscoveredResourceProfilesDiscoveredResourceProfileList) *ListAggregateDiscoveredResourcesResponseBodyDiscoveredResourceProfiles {
	s.DiscoveredResourceProfileList = v
	return s
}

func (s *ListAggregateDiscoveredResourcesResponseBodyDiscoveredResourceProfiles) SetPageNumber(v int32) *ListAggregateDiscoveredResourcesResponseBodyDiscoveredResourceProfiles {
	s.PageNumber = &v
	return s
}

func (s *ListAggregateDiscoveredResourcesResponseBodyDiscoveredResourceProfiles) SetPageSize(v int32) *ListAggregateDiscoveredResourcesResponseBodyDiscoveredResourceProfiles {
	s.PageSize = &v
	return s
}

func (s *ListAggregateDiscoveredResourcesResponseBodyDiscoveredResourceProfiles) SetTotalCount(v int32) *ListAggregateDiscoveredResourcesResponseBodyDiscoveredResourceProfiles {
	s.TotalCount = &v
	return s
}

type ListAggregateDiscoveredResourcesResponseBodyDiscoveredResourceProfilesDiscoveredResourceProfileList struct {
	// The ID of the Alibaba Cloud account to which the resource belongs.
	AccountId *int64 `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The ID of the region.
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The timestamp when the resource was created. Unit: milliseconds.
	ResourceCreationTime *int64 `json:"ResourceCreationTime,omitempty" xml:"ResourceCreationTime,omitempty"`
	// Indicates whether the resource is deleted. Valid values:
	//
	// *   0: The resource is deleted.
	// *   1: The resource is retained.
	ResourceDeleted *int32 `json:"ResourceDeleted,omitempty" xml:"ResourceDeleted,omitempty"`
	// The ID of the resource.
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The name of the resource.
	ResourceName *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	// The ID of the Alibaba Cloud account to which the resource belongs.
	ResourceOwnerId *int64 `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The status of the resource. The parameter value varies based on the resource type and may be left empty. Examples:
	//
	// *   If the value of the ResourceType parameter is ACS::ECS::Instance, the resource is an Elastic Compute Service (ECS) instance that has a specific state. In this case, the valid values of this parameter are Running and Stopped.
	// *   If the value of the ResourceType parameter is ACS::OSS::Bucket, the resource is an Object Storage Service (OSS) bucket that does not have a specific state. In this case, this parameter is left empty.
	ResourceStatus *string `json:"ResourceStatus,omitempty" xml:"ResourceStatus,omitempty"`
	// The type of the resource.
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The tags of the resource.
	Tags *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
}

func (s ListAggregateDiscoveredResourcesResponseBodyDiscoveredResourceProfilesDiscoveredResourceProfileList) String() string {
	return tea.Prettify(s)
}

func (s ListAggregateDiscoveredResourcesResponseBodyDiscoveredResourceProfilesDiscoveredResourceProfileList) GoString() string {
	return s.String()
}

func (s *ListAggregateDiscoveredResourcesResponseBodyDiscoveredResourceProfilesDiscoveredResourceProfileList) SetAccountId(v int64) *ListAggregateDiscoveredResourcesResponseBodyDiscoveredResourceProfilesDiscoveredResourceProfileList {
	s.AccountId = &v
	return s
}

func (s *ListAggregateDiscoveredResourcesResponseBodyDiscoveredResourceProfilesDiscoveredResourceProfileList) SetRegion(v string) *ListAggregateDiscoveredResourcesResponseBodyDiscoveredResourceProfilesDiscoveredResourceProfileList {
	s.Region = &v
	return s
}

func (s *ListAggregateDiscoveredResourcesResponseBodyDiscoveredResourceProfilesDiscoveredResourceProfileList) SetResourceCreationTime(v int64) *ListAggregateDiscoveredResourcesResponseBodyDiscoveredResourceProfilesDiscoveredResourceProfileList {
	s.ResourceCreationTime = &v
	return s
}

func (s *ListAggregateDiscoveredResourcesResponseBodyDiscoveredResourceProfilesDiscoveredResourceProfileList) SetResourceDeleted(v int32) *ListAggregateDiscoveredResourcesResponseBodyDiscoveredResourceProfilesDiscoveredResourceProfileList {
	s.ResourceDeleted = &v
	return s
}

func (s *ListAggregateDiscoveredResourcesResponseBodyDiscoveredResourceProfilesDiscoveredResourceProfileList) SetResourceId(v string) *ListAggregateDiscoveredResourcesResponseBodyDiscoveredResourceProfilesDiscoveredResourceProfileList {
	s.ResourceId = &v
	return s
}

func (s *ListAggregateDiscoveredResourcesResponseBodyDiscoveredResourceProfilesDiscoveredResourceProfileList) SetResourceName(v string) *ListAggregateDiscoveredResourcesResponseBodyDiscoveredResourceProfilesDiscoveredResourceProfileList {
	s.ResourceName = &v
	return s
}

func (s *ListAggregateDiscoveredResourcesResponseBodyDiscoveredResourceProfilesDiscoveredResourceProfileList) SetResourceOwnerId(v int64) *ListAggregateDiscoveredResourcesResponseBodyDiscoveredResourceProfilesDiscoveredResourceProfileList {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListAggregateDiscoveredResourcesResponseBodyDiscoveredResourceProfilesDiscoveredResourceProfileList) SetResourceStatus(v string) *ListAggregateDiscoveredResourcesResponseBodyDiscoveredResourceProfilesDiscoveredResourceProfileList {
	s.ResourceStatus = &v
	return s
}

func (s *ListAggregateDiscoveredResourcesResponseBodyDiscoveredResourceProfilesDiscoveredResourceProfileList) SetResourceType(v string) *ListAggregateDiscoveredResourcesResponseBodyDiscoveredResourceProfilesDiscoveredResourceProfileList {
	s.ResourceType = &v
	return s
}

func (s *ListAggregateDiscoveredResourcesResponseBodyDiscoveredResourceProfilesDiscoveredResourceProfileList) SetTags(v string) *ListAggregateDiscoveredResourcesResponseBodyDiscoveredResourceProfilesDiscoveredResourceProfileList {
	s.Tags = &v
	return s
}

type ListAggregateDiscoveredResourcesResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListAggregateDiscoveredResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListAggregateDiscoveredResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAggregateDiscoveredResourcesResponse) GoString() string {
	return s.String()
}

func (s *ListAggregateDiscoveredResourcesResponse) SetHeaders(v map[string]*string) *ListAggregateDiscoveredResourcesResponse {
	s.Headers = v
	return s
}

func (s *ListAggregateDiscoveredResourcesResponse) SetStatusCode(v int32) *ListAggregateDiscoveredResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAggregateDiscoveredResourcesResponse) SetBody(v *ListAggregateDiscoveredResourcesResponseBody) *ListAggregateDiscoveredResourcesResponse {
	s.Body = v
	return s
}

type ListConfigRulesRequest struct {
	// The ID of the compliance package to which the rule belongs.
	CompliancePackId *string `json:"CompliancePackId,omitempty" xml:"CompliancePackId,omitempty"`
	// The compliance evaluation result of the resources. Valid values:
	//
	// *   COMPLIANT: The resources are evaluated as compliant.
	// *   NON_COMPLIANT: The resources are evaluated as non-compliant.
	// *   NOT_APPLICABLE: The rule does not apply to the resources.
	// *   INSUFFICIENT_DATA: The resource data is insufficient.
	ComplianceType *string `json:"ComplianceType,omitempty" xml:"ComplianceType,omitempty"`
	// The name of the monitoring rule.
	ConfigRuleName *string `json:"ConfigRuleName,omitempty" xml:"ConfigRuleName,omitempty"`
	// The status of the rule. Valid values:
	//
	// *   ACTIVE: The rule is enabled.
	// *   EVALUATING: The rule is triggered and is being used to monitor resource configurations.
	// *   INACTIVE: The rule is disabled.
	ConfigRuleState *string `json:"ConfigRuleState,omitempty" xml:"ConfigRuleState,omitempty"`
	// This parameter is unavailable since 00:00:00 on June 30, 2021. Account group-related operations are provided as an alternative after 00:00:00 on May 30, 2021. If you are using this parameter, we recommend that you switch to account group-related operations after 00:00:00 on May 30, 2021. For information about account groups, see [Overview](~~211534~~).
	MemberId *int64 `json:"MemberId,omitempty" xml:"MemberId,omitempty"`
	// This parameter is unavailable since 00:00:00 on June 30, 2021. Account group-related operations are provided as an alternative after 00:00:00 on May 30, 2021. If you are using this parameter, we recommend that you switch to account group-related operations after 00:00:00 on May 30, 2021. For information about account groups, see [Overview](~~211534~~).
	MultiAccount *bool `json:"MultiAccount,omitempty" xml:"MultiAccount,omitempty"`
	// The page number of the returned page. A minimum of one entry can be returned on each page.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values: 1 to 100.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The risk level of the resources that are not compliant with the rule. Valid values:
	//
	// *   1: high
	// *   2: medium
	// *   3: low
	RiskLevel *int32 `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
}

func (s ListConfigRulesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListConfigRulesRequest) GoString() string {
	return s.String()
}

func (s *ListConfigRulesRequest) SetCompliancePackId(v string) *ListConfigRulesRequest {
	s.CompliancePackId = &v
	return s
}

func (s *ListConfigRulesRequest) SetComplianceType(v string) *ListConfigRulesRequest {
	s.ComplianceType = &v
	return s
}

func (s *ListConfigRulesRequest) SetConfigRuleName(v string) *ListConfigRulesRequest {
	s.ConfigRuleName = &v
	return s
}

func (s *ListConfigRulesRequest) SetConfigRuleState(v string) *ListConfigRulesRequest {
	s.ConfigRuleState = &v
	return s
}

func (s *ListConfigRulesRequest) SetMemberId(v int64) *ListConfigRulesRequest {
	s.MemberId = &v
	return s
}

func (s *ListConfigRulesRequest) SetMultiAccount(v bool) *ListConfigRulesRequest {
	s.MultiAccount = &v
	return s
}

func (s *ListConfigRulesRequest) SetPageNumber(v int32) *ListConfigRulesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListConfigRulesRequest) SetPageSize(v int32) *ListConfigRulesRequest {
	s.PageSize = &v
	return s
}

func (s *ListConfigRulesRequest) SetRiskLevel(v int32) *ListConfigRulesRequest {
	s.RiskLevel = &v
	return s
}

type ListConfigRulesResponseBody struct {
	// The rules.
	ConfigRules *ListConfigRulesResponseBodyConfigRules `json:"ConfigRules,omitempty" xml:"ConfigRules,omitempty" type:"Struct"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListConfigRulesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListConfigRulesResponseBody) GoString() string {
	return s.String()
}

func (s *ListConfigRulesResponseBody) SetConfigRules(v *ListConfigRulesResponseBodyConfigRules) *ListConfigRulesResponseBody {
	s.ConfigRules = v
	return s
}

func (s *ListConfigRulesResponseBody) SetRequestId(v string) *ListConfigRulesResponseBody {
	s.RequestId = &v
	return s
}

type ListConfigRulesResponseBodyConfigRules struct {
	// The tag key of the ENI.
	ConfigRuleList []*ListConfigRulesResponseBodyConfigRulesConfigRuleList `json:"ConfigRuleList,omitempty" xml:"ConfigRuleList,omitempty" type:"Repeated"`
	// The page number of the returned page. A minimum of one entry can be returned on each page.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values: 1 to 100.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of cluster defense rules.
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListConfigRulesResponseBodyConfigRules) String() string {
	return tea.Prettify(s)
}

func (s ListConfigRulesResponseBodyConfigRules) GoString() string {
	return s.String()
}

func (s *ListConfigRulesResponseBodyConfigRules) SetConfigRuleList(v []*ListConfigRulesResponseBodyConfigRulesConfigRuleList) *ListConfigRulesResponseBodyConfigRules {
	s.ConfigRuleList = v
	return s
}

func (s *ListConfigRulesResponseBodyConfigRules) SetPageNumber(v int32) *ListConfigRulesResponseBodyConfigRules {
	s.PageNumber = &v
	return s
}

func (s *ListConfigRulesResponseBodyConfigRules) SetPageSize(v int32) *ListConfigRulesResponseBodyConfigRules {
	s.PageSize = &v
	return s
}

func (s *ListConfigRulesResponseBodyConfigRules) SetTotalCount(v int64) *ListConfigRulesResponseBodyConfigRules {
	s.TotalCount = &v
	return s
}

type ListConfigRulesResponseBodyConfigRulesConfigRuleList struct {
	// The ID of the Alibaba Cloud account that owns the rule.
	AccountId *int64 `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The type of the remediation template. The value is fixed to LC.
	//
	// >  LC stands for Logic Composer.
	AutomationType *string `json:"AutomationType,omitempty" xml:"AutomationType,omitempty"`
	// The statistics about the compliance evaluation results based on the rule.
	Compliance *ListConfigRulesResponseBodyConfigRulesConfigRuleListCompliance `json:"Compliance,omitempty" xml:"Compliance,omitempty" type:"Struct"`
	// The ID of the compliance package to which the rule belongs.
	CompliancePackId *string `json:"CompliancePackId,omitempty" xml:"CompliancePackId,omitempty"`
	// The ARN of the rule.
	ConfigRuleArn *string `json:"ConfigRuleArn,omitempty" xml:"ConfigRuleArn,omitempty"`
	// The ID of the rule.
	ConfigRuleId *string `json:"ConfigRuleId,omitempty" xml:"ConfigRuleId,omitempty"`
	// The name of the rule.
	ConfigRuleName *string `json:"ConfigRuleName,omitempty" xml:"ConfigRuleName,omitempty"`
	// The status of the rule. Valid values:
	//
	// *   ACTIVE: The rule is enabled.
	// *   EVALUATING: The rule is triggered and is being used to monitor resource configurations.
	// *   INACTIVE: The rule is disabled.
	ConfigRuleState *string `json:"ConfigRuleState,omitempty" xml:"ConfigRuleState,omitempty"`
	// The information about how the rule is created.
	CreateBy *ListConfigRulesResponseBodyConfigRulesConfigRuleListCreateBy `json:"CreateBy,omitempty" xml:"CreateBy,omitempty" type:"Struct"`
	// The description of the rule.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The risk level of the resources that are not compliant with the rule. Valid values:
	//
	// *   1: high risk level
	// *   2: medium risk level
	// *   3: low risk level
	RiskLevel *int32 `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
	// The identifier of the rule.
	//
	// *   If the rule was created based on a managed rule, the value of this parameter is the name of the managed rule.
	// *   If the rule is a custom rule, the value of this parameter is the Alibaba Cloud Resource Name (ARN) of the relevant function in Function Compute.
	SourceIdentifier *string `json:"SourceIdentifier,omitempty" xml:"SourceIdentifier,omitempty"`
	// The method that is used to create the rule. Valid values:
	//
	// *   CUSTOM_FC: The rule is a custom rule.
	// *   ALIYUN: The rule is created based on a managed rule of Alibaba Cloud.
	SourceOwner *string `json:"SourceOwner,omitempty" xml:"SourceOwner,omitempty"`
	// The tags.
	Tags []*ListConfigRulesResponseBodyConfigRulesConfigRuleListTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s ListConfigRulesResponseBodyConfigRulesConfigRuleList) String() string {
	return tea.Prettify(s)
}

func (s ListConfigRulesResponseBodyConfigRulesConfigRuleList) GoString() string {
	return s.String()
}

func (s *ListConfigRulesResponseBodyConfigRulesConfigRuleList) SetAccountId(v int64) *ListConfigRulesResponseBodyConfigRulesConfigRuleList {
	s.AccountId = &v
	return s
}

func (s *ListConfigRulesResponseBodyConfigRulesConfigRuleList) SetAutomationType(v string) *ListConfigRulesResponseBodyConfigRulesConfigRuleList {
	s.AutomationType = &v
	return s
}

func (s *ListConfigRulesResponseBodyConfigRulesConfigRuleList) SetCompliance(v *ListConfigRulesResponseBodyConfigRulesConfigRuleListCompliance) *ListConfigRulesResponseBodyConfigRulesConfigRuleList {
	s.Compliance = v
	return s
}

func (s *ListConfigRulesResponseBodyConfigRulesConfigRuleList) SetCompliancePackId(v string) *ListConfigRulesResponseBodyConfigRulesConfigRuleList {
	s.CompliancePackId = &v
	return s
}

func (s *ListConfigRulesResponseBodyConfigRulesConfigRuleList) SetConfigRuleArn(v string) *ListConfigRulesResponseBodyConfigRulesConfigRuleList {
	s.ConfigRuleArn = &v
	return s
}

func (s *ListConfigRulesResponseBodyConfigRulesConfigRuleList) SetConfigRuleId(v string) *ListConfigRulesResponseBodyConfigRulesConfigRuleList {
	s.ConfigRuleId = &v
	return s
}

func (s *ListConfigRulesResponseBodyConfigRulesConfigRuleList) SetConfigRuleName(v string) *ListConfigRulesResponseBodyConfigRulesConfigRuleList {
	s.ConfigRuleName = &v
	return s
}

func (s *ListConfigRulesResponseBodyConfigRulesConfigRuleList) SetConfigRuleState(v string) *ListConfigRulesResponseBodyConfigRulesConfigRuleList {
	s.ConfigRuleState = &v
	return s
}

func (s *ListConfigRulesResponseBodyConfigRulesConfigRuleList) SetCreateBy(v *ListConfigRulesResponseBodyConfigRulesConfigRuleListCreateBy) *ListConfigRulesResponseBodyConfigRulesConfigRuleList {
	s.CreateBy = v
	return s
}

func (s *ListConfigRulesResponseBodyConfigRulesConfigRuleList) SetDescription(v string) *ListConfigRulesResponseBodyConfigRulesConfigRuleList {
	s.Description = &v
	return s
}

func (s *ListConfigRulesResponseBodyConfigRulesConfigRuleList) SetRiskLevel(v int32) *ListConfigRulesResponseBodyConfigRulesConfigRuleList {
	s.RiskLevel = &v
	return s
}

func (s *ListConfigRulesResponseBodyConfigRulesConfigRuleList) SetSourceIdentifier(v string) *ListConfigRulesResponseBodyConfigRulesConfigRuleList {
	s.SourceIdentifier = &v
	return s
}

func (s *ListConfigRulesResponseBodyConfigRulesConfigRuleList) SetSourceOwner(v string) *ListConfigRulesResponseBodyConfigRulesConfigRuleList {
	s.SourceOwner = &v
	return s
}

func (s *ListConfigRulesResponseBodyConfigRulesConfigRuleList) SetTags(v []*ListConfigRulesResponseBodyConfigRulesConfigRuleListTags) *ListConfigRulesResponseBodyConfigRulesConfigRuleList {
	s.Tags = v
	return s
}

type ListConfigRulesResponseBodyConfigRulesConfigRuleListCompliance struct {
	// The compliance evaluation result of the resources. Valid values:
	//
	// *   COMPLIANT: The resources are evaluated as compliant.
	// *   NON_COMPLIANT: The resources are evaluated as non-compliant.
	// *   NOT_APPLICABLE: The rule does not apply to the resources.
	// *   INSUFFICIENT_DATA: The resource data is insufficient.
	ComplianceType *string `json:"ComplianceType,omitempty" xml:"ComplianceType,omitempty"`
	// The number of resources with the specified compliance evaluation result.
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
}

func (s ListConfigRulesResponseBodyConfigRulesConfigRuleListCompliance) String() string {
	return tea.Prettify(s)
}

func (s ListConfigRulesResponseBodyConfigRulesConfigRuleListCompliance) GoString() string {
	return s.String()
}

func (s *ListConfigRulesResponseBodyConfigRulesConfigRuleListCompliance) SetComplianceType(v string) *ListConfigRulesResponseBodyConfigRulesConfigRuleListCompliance {
	s.ComplianceType = &v
	return s
}

func (s *ListConfigRulesResponseBodyConfigRulesConfigRuleListCompliance) SetCount(v int32) *ListConfigRulesResponseBodyConfigRulesConfigRuleListCompliance {
	s.Count = &v
	return s
}

type ListConfigRulesResponseBodyConfigRulesConfigRuleListCreateBy struct {
	// The ID of the compliance package.
	CompliancePackId *string `json:"CompliancePackId,omitempty" xml:"CompliancePackId,omitempty"`
	// The name of the compliance package.
	CompliancePackName *string `json:"CompliancePackName,omitempty" xml:"CompliancePackName,omitempty"`
}

func (s ListConfigRulesResponseBodyConfigRulesConfigRuleListCreateBy) String() string {
	return tea.Prettify(s)
}

func (s ListConfigRulesResponseBodyConfigRulesConfigRuleListCreateBy) GoString() string {
	return s.String()
}

func (s *ListConfigRulesResponseBodyConfigRulesConfigRuleListCreateBy) SetCompliancePackId(v string) *ListConfigRulesResponseBodyConfigRulesConfigRuleListCreateBy {
	s.CompliancePackId = &v
	return s
}

func (s *ListConfigRulesResponseBodyConfigRulesConfigRuleListCreateBy) SetCompliancePackName(v string) *ListConfigRulesResponseBodyConfigRulesConfigRuleListCreateBy {
	s.CompliancePackName = &v
	return s
}

type ListConfigRulesResponseBodyConfigRulesConfigRuleListTags struct {
	// The key of the tag.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListConfigRulesResponseBodyConfigRulesConfigRuleListTags) String() string {
	return tea.Prettify(s)
}

func (s ListConfigRulesResponseBodyConfigRulesConfigRuleListTags) GoString() string {
	return s.String()
}

func (s *ListConfigRulesResponseBodyConfigRulesConfigRuleListTags) SetKey(v string) *ListConfigRulesResponseBodyConfigRulesConfigRuleListTags {
	s.Key = &v
	return s
}

func (s *ListConfigRulesResponseBodyConfigRulesConfigRuleListTags) SetValue(v string) *ListConfigRulesResponseBodyConfigRulesConfigRuleListTags {
	s.Value = &v
	return s
}

type ListConfigRulesResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListConfigRulesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListConfigRulesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListConfigRulesResponse) GoString() string {
	return s.String()
}

func (s *ListConfigRulesResponse) SetHeaders(v map[string]*string) *ListConfigRulesResponse {
	s.Headers = v
	return s
}

func (s *ListConfigRulesResponse) SetStatusCode(v int32) *ListConfigRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListConfigRulesResponse) SetBody(v *ListConfigRulesResponseBody) *ListConfigRulesResponse {
	s.Body = v
	return s
}

type ListDiscoveredResourcesRequest struct {
	// This parameter has been removed from 00:00:00, June 30, 2021. Account group-related operations are provided as an alternative after 00:00:00, May 30, 2021. If you are using this parameter, we recommend that you switch to account group-related operations after 00:00:00, May 30, 2021. For more information about account groups, see [Overview](~~211534~~).
	MemberId *int64 `json:"MemberId,omitempty" xml:"MemberId,omitempty"`
	// This parameter has been removed from 00:00:00, June 30, 2021. Account group-related operations are provided as an alternative after 00:00:00, May 30, 2021. If you are using this parameter, we recommend that you switch to account group-related operations after 00:00:00, May 30, 2021. For more information about account groups, see [Overview](~~211534~~).
	MultiAccount *bool `json:"MultiAccount,omitempty" xml:"MultiAccount,omitempty"`
	// The number of the page to return. Pages start from page 1.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Valid values: 1 to 100.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the region in which the resource resides. Separate multiple regions IDs with commas (,).
	Regions *string `json:"Regions,omitempty" xml:"Regions,omitempty"`
	// The status of the resource. Valid values:
	//
	// *   0: The resource is deleted.
	// *   1: The resource is retained. This is the default value.
	ResourceDeleted *int32 `json:"ResourceDeleted,omitempty" xml:"ResourceDeleted,omitempty"`
	// The ID of the resource.
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The type of the resource. Separate multiple resource types with commas (,).
	ResourceTypes *string `json:"ResourceTypes,omitempty" xml:"ResourceTypes,omitempty"`
}

func (s ListDiscoveredResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDiscoveredResourcesRequest) GoString() string {
	return s.String()
}

func (s *ListDiscoveredResourcesRequest) SetMemberId(v int64) *ListDiscoveredResourcesRequest {
	s.MemberId = &v
	return s
}

func (s *ListDiscoveredResourcesRequest) SetMultiAccount(v bool) *ListDiscoveredResourcesRequest {
	s.MultiAccount = &v
	return s
}

func (s *ListDiscoveredResourcesRequest) SetPageNumber(v int32) *ListDiscoveredResourcesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListDiscoveredResourcesRequest) SetPageSize(v int32) *ListDiscoveredResourcesRequest {
	s.PageSize = &v
	return s
}

func (s *ListDiscoveredResourcesRequest) SetRegions(v string) *ListDiscoveredResourcesRequest {
	s.Regions = &v
	return s
}

func (s *ListDiscoveredResourcesRequest) SetResourceDeleted(v int32) *ListDiscoveredResourcesRequest {
	s.ResourceDeleted = &v
	return s
}

func (s *ListDiscoveredResourcesRequest) SetResourceId(v string) *ListDiscoveredResourcesRequest {
	s.ResourceId = &v
	return s
}

func (s *ListDiscoveredResourcesRequest) SetResourceTypes(v string) *ListDiscoveredResourcesRequest {
	s.ResourceTypes = &v
	return s
}

type ListDiscoveredResourcesResponseBody struct {
	// The information about the returned resources.
	DiscoveredResourceProfiles *ListDiscoveredResourcesResponseBodyDiscoveredResourceProfiles `json:"DiscoveredResourceProfiles,omitempty" xml:"DiscoveredResourceProfiles,omitempty" type:"Struct"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListDiscoveredResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDiscoveredResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *ListDiscoveredResourcesResponseBody) SetDiscoveredResourceProfiles(v *ListDiscoveredResourcesResponseBodyDiscoveredResourceProfiles) *ListDiscoveredResourcesResponseBody {
	s.DiscoveredResourceProfiles = v
	return s
}

func (s *ListDiscoveredResourcesResponseBody) SetRequestId(v string) *ListDiscoveredResourcesResponseBody {
	s.RequestId = &v
	return s
}

type ListDiscoveredResourcesResponseBodyDiscoveredResourceProfiles struct {
	// The information about the resource.
	DiscoveredResourceProfileList []*ListDiscoveredResourcesResponseBodyDiscoveredResourceProfilesDiscoveredResourceProfileList `json:"DiscoveredResourceProfileList,omitempty" xml:"DiscoveredResourceProfileList,omitempty" type:"Repeated"`
	// The page number of the returned page. Pages start from page 1.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page. Valid values: 1 to 100.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of resources.
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListDiscoveredResourcesResponseBodyDiscoveredResourceProfiles) String() string {
	return tea.Prettify(s)
}

func (s ListDiscoveredResourcesResponseBodyDiscoveredResourceProfiles) GoString() string {
	return s.String()
}

func (s *ListDiscoveredResourcesResponseBodyDiscoveredResourceProfiles) SetDiscoveredResourceProfileList(v []*ListDiscoveredResourcesResponseBodyDiscoveredResourceProfilesDiscoveredResourceProfileList) *ListDiscoveredResourcesResponseBodyDiscoveredResourceProfiles {
	s.DiscoveredResourceProfileList = v
	return s
}

func (s *ListDiscoveredResourcesResponseBodyDiscoveredResourceProfiles) SetPageNumber(v int32) *ListDiscoveredResourcesResponseBodyDiscoveredResourceProfiles {
	s.PageNumber = &v
	return s
}

func (s *ListDiscoveredResourcesResponseBodyDiscoveredResourceProfiles) SetPageSize(v int32) *ListDiscoveredResourcesResponseBodyDiscoveredResourceProfiles {
	s.PageSize = &v
	return s
}

func (s *ListDiscoveredResourcesResponseBodyDiscoveredResourceProfiles) SetTotalCount(v int32) *ListDiscoveredResourcesResponseBodyDiscoveredResourceProfiles {
	s.TotalCount = &v
	return s
}

type ListDiscoveredResourcesResponseBodyDiscoveredResourceProfilesDiscoveredResourceProfileList struct {
	// The ID of the Alibaba Cloud account to which the resource belongs.
	AccountId *int64 `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The ID of the region.
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The timestamp when the resource was created.
	ResourceCreationTime *int64 `json:"ResourceCreationTime,omitempty" xml:"ResourceCreationTime,omitempty"`
	// Indicates whether the resource is deleted. Valid values:
	//
	// *   1: The resource is retained.
	// *   0: The resource is deleted.
	ResourceDeleted *int32 `json:"ResourceDeleted,omitempty" xml:"ResourceDeleted,omitempty"`
	// The ID of the resource.
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The name of the resource.
	ResourceName *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	// The status of the resource. The parameter value varies based on the resource type and may be left empty.
	//
	// Examples:
	//
	// *   If the value of the ResourceType parameter is ACS::ECS::Instance, the resource is an Elastic Compute Service (ECS) instance that has a specific state. In this case, the valid values of this parameter are Running and Stopped.
	// *   If the value of the ResourceType parameter is ACS::OSS::Bucket, the resource is an Object Storage Service (OSS) bucket that does not have a specific state. In this case, this parameter is left empty.
	ResourceStatus *string `json:"ResourceStatus,omitempty" xml:"ResourceStatus,omitempty"`
	// The type of the resource.
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The tags of the resource.
	Tags *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
}

func (s ListDiscoveredResourcesResponseBodyDiscoveredResourceProfilesDiscoveredResourceProfileList) String() string {
	return tea.Prettify(s)
}

func (s ListDiscoveredResourcesResponseBodyDiscoveredResourceProfilesDiscoveredResourceProfileList) GoString() string {
	return s.String()
}

func (s *ListDiscoveredResourcesResponseBodyDiscoveredResourceProfilesDiscoveredResourceProfileList) SetAccountId(v int64) *ListDiscoveredResourcesResponseBodyDiscoveredResourceProfilesDiscoveredResourceProfileList {
	s.AccountId = &v
	return s
}

func (s *ListDiscoveredResourcesResponseBodyDiscoveredResourceProfilesDiscoveredResourceProfileList) SetRegion(v string) *ListDiscoveredResourcesResponseBodyDiscoveredResourceProfilesDiscoveredResourceProfileList {
	s.Region = &v
	return s
}

func (s *ListDiscoveredResourcesResponseBodyDiscoveredResourceProfilesDiscoveredResourceProfileList) SetResourceCreationTime(v int64) *ListDiscoveredResourcesResponseBodyDiscoveredResourceProfilesDiscoveredResourceProfileList {
	s.ResourceCreationTime = &v
	return s
}

func (s *ListDiscoveredResourcesResponseBodyDiscoveredResourceProfilesDiscoveredResourceProfileList) SetResourceDeleted(v int32) *ListDiscoveredResourcesResponseBodyDiscoveredResourceProfilesDiscoveredResourceProfileList {
	s.ResourceDeleted = &v
	return s
}

func (s *ListDiscoveredResourcesResponseBodyDiscoveredResourceProfilesDiscoveredResourceProfileList) SetResourceId(v string) *ListDiscoveredResourcesResponseBodyDiscoveredResourceProfilesDiscoveredResourceProfileList {
	s.ResourceId = &v
	return s
}

func (s *ListDiscoveredResourcesResponseBodyDiscoveredResourceProfilesDiscoveredResourceProfileList) SetResourceName(v string) *ListDiscoveredResourcesResponseBodyDiscoveredResourceProfilesDiscoveredResourceProfileList {
	s.ResourceName = &v
	return s
}

func (s *ListDiscoveredResourcesResponseBodyDiscoveredResourceProfilesDiscoveredResourceProfileList) SetResourceStatus(v string) *ListDiscoveredResourcesResponseBodyDiscoveredResourceProfilesDiscoveredResourceProfileList {
	s.ResourceStatus = &v
	return s
}

func (s *ListDiscoveredResourcesResponseBodyDiscoveredResourceProfilesDiscoveredResourceProfileList) SetResourceType(v string) *ListDiscoveredResourcesResponseBodyDiscoveredResourceProfilesDiscoveredResourceProfileList {
	s.ResourceType = &v
	return s
}

func (s *ListDiscoveredResourcesResponseBodyDiscoveredResourceProfilesDiscoveredResourceProfileList) SetTags(v string) *ListDiscoveredResourcesResponseBodyDiscoveredResourceProfilesDiscoveredResourceProfileList {
	s.Tags = &v
	return s
}

type ListDiscoveredResourcesResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListDiscoveredResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListDiscoveredResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDiscoveredResourcesResponse) GoString() string {
	return s.String()
}

func (s *ListDiscoveredResourcesResponse) SetHeaders(v map[string]*string) *ListDiscoveredResourcesResponse {
	s.Headers = v
	return s
}

func (s *ListDiscoveredResourcesResponse) SetStatusCode(v int32) *ListDiscoveredResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDiscoveredResourcesResponse) SetBody(v *ListDiscoveredResourcesResponseBody) *ListDiscoveredResourcesResponse {
	s.Body = v
	return s
}

type ListRemediationTemplatesRequest struct {
	// The identifier of the managed rule.
	ManagedRuleIdentifier *string `json:"ManagedRuleIdentifier,omitempty" xml:"ManagedRuleIdentifier,omitempty"`
	// The type of the remediation template. Valid value: OOS, which stands for Operation Orchestration Service.
	RemediationType *string `json:"RemediationType,omitempty" xml:"RemediationType,omitempty"`
}

func (s ListRemediationTemplatesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListRemediationTemplatesRequest) GoString() string {
	return s.String()
}

func (s *ListRemediationTemplatesRequest) SetManagedRuleIdentifier(v string) *ListRemediationTemplatesRequest {
	s.ManagedRuleIdentifier = &v
	return s
}

func (s *ListRemediationTemplatesRequest) SetRemediationType(v string) *ListRemediationTemplatesRequest {
	s.RemediationType = &v
	return s
}

type ListRemediationTemplatesResponseBody struct {
	// The information about the remediation template.
	RemediationTemplates []*ListRemediationTemplatesResponseBodyRemediationTemplates `json:"RemediationTemplates,omitempty" xml:"RemediationTemplates,omitempty" type:"Repeated"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListRemediationTemplatesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListRemediationTemplatesResponseBody) GoString() string {
	return s.String()
}

func (s *ListRemediationTemplatesResponseBody) SetRemediationTemplates(v []*ListRemediationTemplatesResponseBodyRemediationTemplates) *ListRemediationTemplatesResponseBody {
	s.RemediationTemplates = v
	return s
}

func (s *ListRemediationTemplatesResponseBody) SetRequestId(v string) *ListRemediationTemplatesResponseBody {
	s.RequestId = &v
	return s
}

type ListRemediationTemplatesResponseBodyRemediationTemplates struct {
	// The type of the remediation template. Valid value: OOS, which stands for Operation Orchestration Service.
	RemediationType *string `json:"RemediationType,omitempty" xml:"RemediationType,omitempty"`
	// The configuration of the remediation template.
	TemplateDefinition *string `json:"TemplateDefinition,omitempty" xml:"TemplateDefinition,omitempty"`
	// The ID of the remediation template.
	TemplateIdentifier *string `json:"TemplateIdentifier,omitempty" xml:"TemplateIdentifier,omitempty"`
	// The name of the remediation template.
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
}

func (s ListRemediationTemplatesResponseBodyRemediationTemplates) String() string {
	return tea.Prettify(s)
}

func (s ListRemediationTemplatesResponseBodyRemediationTemplates) GoString() string {
	return s.String()
}

func (s *ListRemediationTemplatesResponseBodyRemediationTemplates) SetRemediationType(v string) *ListRemediationTemplatesResponseBodyRemediationTemplates {
	s.RemediationType = &v
	return s
}

func (s *ListRemediationTemplatesResponseBodyRemediationTemplates) SetTemplateDefinition(v string) *ListRemediationTemplatesResponseBodyRemediationTemplates {
	s.TemplateDefinition = &v
	return s
}

func (s *ListRemediationTemplatesResponseBodyRemediationTemplates) SetTemplateIdentifier(v string) *ListRemediationTemplatesResponseBodyRemediationTemplates {
	s.TemplateIdentifier = &v
	return s
}

func (s *ListRemediationTemplatesResponseBodyRemediationTemplates) SetTemplateName(v string) *ListRemediationTemplatesResponseBodyRemediationTemplates {
	s.TemplateName = &v
	return s
}

type ListRemediationTemplatesResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListRemediationTemplatesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListRemediationTemplatesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListRemediationTemplatesResponse) GoString() string {
	return s.String()
}

func (s *ListRemediationTemplatesResponse) SetHeaders(v map[string]*string) *ListRemediationTemplatesResponse {
	s.Headers = v
	return s
}

func (s *ListRemediationTemplatesResponse) SetStatusCode(v int32) *ListRemediationTemplatesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRemediationTemplatesResponse) SetBody(v *ListRemediationTemplatesResponseBody) *ListRemediationTemplatesResponse {
	s.Body = v
	return s
}

type PutConfigRuleRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the value, but you must ensure that it is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the rule.
	ConfigRuleId *string `json:"ConfigRuleId,omitempty" xml:"ConfigRuleId,omitempty"`
	// The name of the rule.
	ConfigRuleName *string `json:"ConfigRuleName,omitempty" xml:"ConfigRuleName,omitempty"`
	// The description of the rule.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The settings of the input parameters for the rule.
	InputParameters *string `json:"InputParameters,omitempty" xml:"InputParameters,omitempty"`
	// This parameter is scheduled to be removed before 00:00:00, June 30, 2021. Account group-related APIs will be provided as an alternative before 00:00:00, May 30, 2021. If you are using this parameter, we recommend that you switch to account group-related APIs after 00:00:00, May 30, 2021. For more information, see [Account groups](~~211534~~).
	MemberId *int64 `json:"MemberId,omitempty" xml:"MemberId,omitempty"`
	// This parameter is scheduled to be removed before 00:00:00, June 30, 2021. Account group-related APIs will be provided as an alternative before 00:00:00, May 30, 2021. If you are using this parameter, we recommend that you switch to account group-related APIs after 00:00:00, May 30, 2021. For more information, see [Account groups](~~211534~~).
	MultiAccount *bool `json:"MultiAccount,omitempty" xml:"MultiAccount,omitempty"`
	// The risk level of the non-compliant resource. Valid values:
	//
	// *   1: high risk
	// *   2: medium risk
	// *   3: low risk
	RiskLevel *int32 `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
	// The ID of the resource to be evaluated.
	//
	// *   If you do not set this parameter, the rule evaluates resources of all the types specified by ScopeComplianceResourceTypes.
	// *   If you specify a resource ID, the rule evaluates the resource identified by the ID.
	ScopeComplianceResourceId *string `json:"ScopeComplianceResourceId,omitempty" xml:"ScopeComplianceResourceId,omitempty"`
	// The types of the resources to be evaluated.
	ScopeComplianceResourceTypes *string `json:"ScopeComplianceResourceTypes,omitempty" xml:"ScopeComplianceResourceTypes,omitempty"`
	// The trigger type of the rule. Valid values:
	//
	// *   ConfigurationItemChangeNotification: The rule is triggered by configuration changes.
	// *   ScheduledNotification: The rule is triggered as scheduled.
	SourceDetailMessageType *string `json:"SourceDetailMessageType,omitempty" xml:"SourceDetailMessageType,omitempty"`
	// The identifier of the rule.
	//
	// *   For a managed rule, the value is the name of the managed rule.
	// *   For a custom rule, the value is the Alibaba Cloud Resource Name (ARN) of the custom rule.
	SourceIdentifier *string `json:"SourceIdentifier,omitempty" xml:"SourceIdentifier,omitempty"`
	// The frequency at which the rule is executed. Valid values:
	//
	// *   One_Hour: 1 hour
	// *   Three_Hours: 3 hours
	// *   Six_Hours: 6 hours
	// *   Twelve_Hours: 12 hours
	// *   TwentyFour_Hours: 24 hours
	SourceMaximumExecutionFrequency *string `json:"SourceMaximumExecutionFrequency,omitempty" xml:"SourceMaximumExecutionFrequency,omitempty"`
	// Indicates whether you or Alibaba Cloud owns and manages the rule. Valid values:
	//
	// *   CUSTOM_FC: The rule is a custom rule that you own.
	// *   ALIYUN: The rule is a managed rule of Alibaba Cloud.
	SourceOwner *string `json:"SourceOwner,omitempty" xml:"SourceOwner,omitempty"`
}

func (s PutConfigRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s PutConfigRuleRequest) GoString() string {
	return s.String()
}

func (s *PutConfigRuleRequest) SetClientToken(v string) *PutConfigRuleRequest {
	s.ClientToken = &v
	return s
}

func (s *PutConfigRuleRequest) SetConfigRuleId(v string) *PutConfigRuleRequest {
	s.ConfigRuleId = &v
	return s
}

func (s *PutConfigRuleRequest) SetConfigRuleName(v string) *PutConfigRuleRequest {
	s.ConfigRuleName = &v
	return s
}

func (s *PutConfigRuleRequest) SetDescription(v string) *PutConfigRuleRequest {
	s.Description = &v
	return s
}

func (s *PutConfigRuleRequest) SetInputParameters(v string) *PutConfigRuleRequest {
	s.InputParameters = &v
	return s
}

func (s *PutConfigRuleRequest) SetMemberId(v int64) *PutConfigRuleRequest {
	s.MemberId = &v
	return s
}

func (s *PutConfigRuleRequest) SetMultiAccount(v bool) *PutConfigRuleRequest {
	s.MultiAccount = &v
	return s
}

func (s *PutConfigRuleRequest) SetRiskLevel(v int32) *PutConfigRuleRequest {
	s.RiskLevel = &v
	return s
}

func (s *PutConfigRuleRequest) SetScopeComplianceResourceId(v string) *PutConfigRuleRequest {
	s.ScopeComplianceResourceId = &v
	return s
}

func (s *PutConfigRuleRequest) SetScopeComplianceResourceTypes(v string) *PutConfigRuleRequest {
	s.ScopeComplianceResourceTypes = &v
	return s
}

func (s *PutConfigRuleRequest) SetSourceDetailMessageType(v string) *PutConfigRuleRequest {
	s.SourceDetailMessageType = &v
	return s
}

func (s *PutConfigRuleRequest) SetSourceIdentifier(v string) *PutConfigRuleRequest {
	s.SourceIdentifier = &v
	return s
}

func (s *PutConfigRuleRequest) SetSourceMaximumExecutionFrequency(v string) *PutConfigRuleRequest {
	s.SourceMaximumExecutionFrequency = &v
	return s
}

func (s *PutConfigRuleRequest) SetSourceOwner(v string) *PutConfigRuleRequest {
	s.SourceOwner = &v
	return s
}

type PutConfigRuleResponseBody struct {
	// The ID of the rule.
	ConfigRuleId *string `json:"ConfigRuleId,omitempty" xml:"ConfigRuleId,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s PutConfigRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PutConfigRuleResponseBody) GoString() string {
	return s.String()
}

func (s *PutConfigRuleResponseBody) SetConfigRuleId(v string) *PutConfigRuleResponseBody {
	s.ConfigRuleId = &v
	return s
}

func (s *PutConfigRuleResponseBody) SetRequestId(v string) *PutConfigRuleResponseBody {
	s.RequestId = &v
	return s
}

type PutConfigRuleResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *PutConfigRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PutConfigRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s PutConfigRuleResponse) GoString() string {
	return s.String()
}

func (s *PutConfigRuleResponse) SetHeaders(v map[string]*string) *PutConfigRuleResponse {
	s.Headers = v
	return s
}

func (s *PutConfigRuleResponse) SetStatusCode(v int32) *PutConfigRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *PutConfigRuleResponse) SetBody(v *PutConfigRuleResponseBody) *PutConfigRuleResponse {
	s.Body = v
	return s
}

type PutConfigurationRecorderRequest struct {
	// The types of the resources that you want Cloud Config to monitor. Separate multiple resource types with commas (,).
	ResourceTypes *string `json:"ResourceTypes,omitempty" xml:"ResourceTypes,omitempty"`
}

func (s PutConfigurationRecorderRequest) String() string {
	return tea.Prettify(s)
}

func (s PutConfigurationRecorderRequest) GoString() string {
	return s.String()
}

func (s *PutConfigurationRecorderRequest) SetResourceTypes(v string) *PutConfigurationRecorderRequest {
	s.ResourceTypes = &v
	return s
}

type PutConfigurationRecorderResponseBody struct {
	// The details of the configuration recorder that monitors resources.
	ConfigurationRecorder *PutConfigurationRecorderResponseBodyConfigurationRecorder `json:"ConfigurationRecorder,omitempty" xml:"ConfigurationRecorder,omitempty" type:"Struct"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s PutConfigurationRecorderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PutConfigurationRecorderResponseBody) GoString() string {
	return s.String()
}

func (s *PutConfigurationRecorderResponseBody) SetConfigurationRecorder(v *PutConfigurationRecorderResponseBodyConfigurationRecorder) *PutConfigurationRecorderResponseBody {
	s.ConfigurationRecorder = v
	return s
}

func (s *PutConfigurationRecorderResponseBody) SetRequestId(v string) *PutConfigurationRecorderResponseBody {
	s.RequestId = &v
	return s
}

type PutConfigurationRecorderResponseBodyConfigurationRecorder struct {
	// The ID of your Alibaba Cloud account.
	AccountId *int64 `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The status of the configuration recorder. Valid values:
	//
	// *   REGISTRABLE: The configuration recorder has not been registered.
	// *   BUILDING: The configuration recorder is being deployed.
	// *   REGISTERED: The configuration recorder has been registered.
	// *   REBUILDING: The configuration recorder is being redeployed.
	ConfigurationRecorderStatus *string `json:"ConfigurationRecorderStatus,omitempty" xml:"ConfigurationRecorderStatus,omitempty"`
	// The types of resources that are monitored by Cloud Config.
	ResourceTypes []*string `json:"ResourceTypes,omitempty" xml:"ResourceTypes,omitempty" type:"Repeated"`
}

func (s PutConfigurationRecorderResponseBodyConfigurationRecorder) String() string {
	return tea.Prettify(s)
}

func (s PutConfigurationRecorderResponseBodyConfigurationRecorder) GoString() string {
	return s.String()
}

func (s *PutConfigurationRecorderResponseBodyConfigurationRecorder) SetAccountId(v int64) *PutConfigurationRecorderResponseBodyConfigurationRecorder {
	s.AccountId = &v
	return s
}

func (s *PutConfigurationRecorderResponseBodyConfigurationRecorder) SetConfigurationRecorderStatus(v string) *PutConfigurationRecorderResponseBodyConfigurationRecorder {
	s.ConfigurationRecorderStatus = &v
	return s
}

func (s *PutConfigurationRecorderResponseBodyConfigurationRecorder) SetResourceTypes(v []*string) *PutConfigurationRecorderResponseBodyConfigurationRecorder {
	s.ResourceTypes = v
	return s
}

type PutConfigurationRecorderResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *PutConfigurationRecorderResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PutConfigurationRecorderResponse) String() string {
	return tea.Prettify(s)
}

func (s PutConfigurationRecorderResponse) GoString() string {
	return s.String()
}

func (s *PutConfigurationRecorderResponse) SetHeaders(v map[string]*string) *PutConfigurationRecorderResponse {
	s.Headers = v
	return s
}

func (s *PutConfigurationRecorderResponse) SetStatusCode(v int32) *PutConfigurationRecorderResponse {
	s.StatusCode = &v
	return s
}

func (s *PutConfigurationRecorderResponse) SetBody(v *PutConfigurationRecorderResponseBody) *PutConfigurationRecorderResponse {
	s.Body = v
	return s
}

type PutDeliveryChannelRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate a token, but you must make sure that it is unique among different requests. The token cannot exceed 64 characters in length and can only contain ASCII characters.
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The Alibaba Cloud Resource Name (ARN) of the role to be assumed by the delivery method. This parameter is required when you create a delivery method.
	//
	// >  If the delivery method assumes the service linked role for Cloud Config, you can specify the ARN in the format of the provided example and replace the account ID with the ID of your Alibaba Cloud account.
	DeliveryChannelAssumeRoleArn *string `json:"DeliveryChannelAssumeRoleArn,omitempty" xml:"DeliveryChannelAssumeRoleArn,omitempty"`
	// The rule attached to the delivery method. This parameter is applicable only to delivery methods of the MNS type.
	//
	// This parameter allows you to specify the lowest risk level for the events to subscribe to and the resource types for which you want to subscribe to events.
	//
	// *   To specify the lowest risk level for the events to subscribe to, use the following format:`{"filterType":"RuleRiskLevel","value":"1","multiple":false}.`
	//
	//     The value field indicates the lowest risk level and can be set to 1, 2, or 3, which indicates the high risk level, the medium risk level, and the low risk level, respectively.
	//
	// *   To specify the resource types for which you want to subscribe to events, use the following format:`{"filterType":"ResourceType","values":["ACS::ACK::Cluster","ACS::ActionTrail::Trail","ACS::CBWP::CommonBandwidthPackage"],"multiple":true}.`
	//
	//     The values field indicates the resource types. Its value is a JSON array.
	//
	//     Example: `[{"filterType":"ResourceType","values":["ACS::ActionTrail::Trail","ACS::CBWP::CommonBandwidthPackage","ACS::CDN::Domain","ACS::CEN::CenBandwidthPackage","ACS::CEN::CenInstance","ACS::CEN::Flowlog","ACS::DdosCoo::Instance"],"multiple":true}]`
	DeliveryChannelCondition *string `json:"DeliveryChannelCondition,omitempty" xml:"DeliveryChannelCondition,omitempty"`
	// The ID of the delivery method. This parameter is required when you modify a delivery method.
	DeliveryChannelId *string `json:"DeliveryChannelId,omitempty" xml:"DeliveryChannelId,omitempty"`
	// The name of the delivery method.
	DeliveryChannelName *string `json:"DeliveryChannelName,omitempty" xml:"DeliveryChannelName,omitempty"`
	// The ARN of the delivery destination. This parameter is required when you create a delivery method. The value must be in one of the following formats:
	//
	// *   `acs:oss:{RegionId}:{Aliuid}:{bucketName}` if your delivery destination is an Object Storage Service (OSS) bucket.
	// *   `acs:mns:{RegionId}:{Aliuid}:/topics/{topicName}` if your delivery destination is a Message Service (MNS) topic.
	// *   `acs:log:{RegionId}:{Aliuid}:project/{projectName}/logstore/{logstoreName}` if your delivery destination is a Log Service Logstore.
	DeliveryChannelTargetArn *string `json:"DeliveryChannelTargetArn,omitempty" xml:"DeliveryChannelTargetArn,omitempty"`
	// The type of the delivery method. This parameter is required when you create a delivery method. Valid values:
	//
	// *   OSS
	// *   MNS
	// *   SLS
	DeliveryChannelType *string `json:"DeliveryChannelType,omitempty" xml:"DeliveryChannelType,omitempty"`
	// The description of the delivery method.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The status of the delivery method. Valid values:
	//
	// *   0: The delivery method is disabled.
	// *   1: The delivery destination is enabled. This is the default value.
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s PutDeliveryChannelRequest) String() string {
	return tea.Prettify(s)
}

func (s PutDeliveryChannelRequest) GoString() string {
	return s.String()
}

func (s *PutDeliveryChannelRequest) SetClientToken(v string) *PutDeliveryChannelRequest {
	s.ClientToken = &v
	return s
}

func (s *PutDeliveryChannelRequest) SetDeliveryChannelAssumeRoleArn(v string) *PutDeliveryChannelRequest {
	s.DeliveryChannelAssumeRoleArn = &v
	return s
}

func (s *PutDeliveryChannelRequest) SetDeliveryChannelCondition(v string) *PutDeliveryChannelRequest {
	s.DeliveryChannelCondition = &v
	return s
}

func (s *PutDeliveryChannelRequest) SetDeliveryChannelId(v string) *PutDeliveryChannelRequest {
	s.DeliveryChannelId = &v
	return s
}

func (s *PutDeliveryChannelRequest) SetDeliveryChannelName(v string) *PutDeliveryChannelRequest {
	s.DeliveryChannelName = &v
	return s
}

func (s *PutDeliveryChannelRequest) SetDeliveryChannelTargetArn(v string) *PutDeliveryChannelRequest {
	s.DeliveryChannelTargetArn = &v
	return s
}

func (s *PutDeliveryChannelRequest) SetDeliveryChannelType(v string) *PutDeliveryChannelRequest {
	s.DeliveryChannelType = &v
	return s
}

func (s *PutDeliveryChannelRequest) SetDescription(v string) *PutDeliveryChannelRequest {
	s.Description = &v
	return s
}

func (s *PutDeliveryChannelRequest) SetStatus(v int32) *PutDeliveryChannelRequest {
	s.Status = &v
	return s
}

type PutDeliveryChannelResponseBody struct {
	// The ID of the delivery method that you created or the updated ID of the delivery method that you modified.
	DeliveryChannelId *string `json:"DeliveryChannelId,omitempty" xml:"DeliveryChannelId,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s PutDeliveryChannelResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PutDeliveryChannelResponseBody) GoString() string {
	return s.String()
}

func (s *PutDeliveryChannelResponseBody) SetDeliveryChannelId(v string) *PutDeliveryChannelResponseBody {
	s.DeliveryChannelId = &v
	return s
}

func (s *PutDeliveryChannelResponseBody) SetRequestId(v string) *PutDeliveryChannelResponseBody {
	s.RequestId = &v
	return s
}

type PutDeliveryChannelResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *PutDeliveryChannelResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PutDeliveryChannelResponse) String() string {
	return tea.Prettify(s)
}

func (s PutDeliveryChannelResponse) GoString() string {
	return s.String()
}

func (s *PutDeliveryChannelResponse) SetHeaders(v map[string]*string) *PutDeliveryChannelResponse {
	s.Headers = v
	return s
}

func (s *PutDeliveryChannelResponse) SetStatusCode(v int32) *PutDeliveryChannelResponse {
	s.StatusCode = &v
	return s
}

func (s *PutDeliveryChannelResponse) SetBody(v *PutDeliveryChannelResponseBody) *PutDeliveryChannelResponse {
	s.Body = v
	return s
}

type PutEvaluationsRequest struct {
	// The compliance evaluation results.
	Evaluations *string `json:"Evaluations,omitempty" xml:"Evaluations,omitempty"`
	// The token used to return the response.
	ResultToken *string `json:"ResultToken,omitempty" xml:"ResultToken,omitempty"`
}

func (s PutEvaluationsRequest) String() string {
	return tea.Prettify(s)
}

func (s PutEvaluationsRequest) GoString() string {
	return s.String()
}

func (s *PutEvaluationsRequest) SetEvaluations(v string) *PutEvaluationsRequest {
	s.Evaluations = &v
	return s
}

func (s *PutEvaluationsRequest) SetResultToken(v string) *PutEvaluationsRequest {
	s.ResultToken = &v
	return s
}

type PutEvaluationsResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the operation was successful. Valid values:
	//
	// *   true: The operation was successful.
	// *   false: The operation failed.
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s PutEvaluationsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PutEvaluationsResponseBody) GoString() string {
	return s.String()
}

func (s *PutEvaluationsResponseBody) SetRequestId(v string) *PutEvaluationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *PutEvaluationsResponseBody) SetResult(v bool) *PutEvaluationsResponseBody {
	s.Result = &v
	return s
}

type PutEvaluationsResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *PutEvaluationsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PutEvaluationsResponse) String() string {
	return tea.Prettify(s)
}

func (s PutEvaluationsResponse) GoString() string {
	return s.String()
}

func (s *PutEvaluationsResponse) SetHeaders(v map[string]*string) *PutEvaluationsResponse {
	s.Headers = v
	return s
}

func (s *PutEvaluationsResponse) SetStatusCode(v int32) *PutEvaluationsResponse {
	s.StatusCode = &v
	return s
}

func (s *PutEvaluationsResponse) SetBody(v *PutEvaluationsResponseBody) *PutEvaluationsResponse {
	s.Body = v
	return s
}

type StartConfigRuleEvaluationRequest struct {
	// The ID of the compliance package.
	//
	// For more information about how to obtain the ID of a compliance package, see [ListCompliancePacks](~~263332~~).
	//
	// >  You must set one of the `CompliancePackId` and `ConfigRuleId` parameters.
	CompliancePackId *string `json:"CompliancePackId,omitempty" xml:"CompliancePackId,omitempty"`
	// The ID of the rule.
	//
	// For more information about how to obtain the ID of a rule, see [ListConfigRules](~~169607~~).
	ConfigRuleId *string `json:"ConfigRuleId,omitempty" xml:"ConfigRuleId,omitempty"`
	// Specifies whether to reevaluate a non-compliant resource that is ignored. Valid values:
	//
	// *   true: reevaluates a non-compliant resource that is ignored.
	// *   false: continues to ignore a non-compliant resource.
	RevertEvaluation *bool `json:"RevertEvaluation,omitempty" xml:"RevertEvaluation,omitempty"`
}

func (s StartConfigRuleEvaluationRequest) String() string {
	return tea.Prettify(s)
}

func (s StartConfigRuleEvaluationRequest) GoString() string {
	return s.String()
}

func (s *StartConfigRuleEvaluationRequest) SetCompliancePackId(v string) *StartConfigRuleEvaluationRequest {
	s.CompliancePackId = &v
	return s
}

func (s *StartConfigRuleEvaluationRequest) SetConfigRuleId(v string) *StartConfigRuleEvaluationRequest {
	s.ConfigRuleId = &v
	return s
}

func (s *StartConfigRuleEvaluationRequest) SetRevertEvaluation(v bool) *StartConfigRuleEvaluationRequest {
	s.RevertEvaluation = &v
	return s
}

type StartConfigRuleEvaluationResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the operation is successful. Valid values:
	//
	// *   true: The operation is successful.
	// *   false: The operation fails.
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s StartConfigRuleEvaluationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StartConfigRuleEvaluationResponseBody) GoString() string {
	return s.String()
}

func (s *StartConfigRuleEvaluationResponseBody) SetRequestId(v string) *StartConfigRuleEvaluationResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartConfigRuleEvaluationResponseBody) SetResult(v bool) *StartConfigRuleEvaluationResponseBody {
	s.Result = &v
	return s
}

type StartConfigRuleEvaluationResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *StartConfigRuleEvaluationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StartConfigRuleEvaluationResponse) String() string {
	return tea.Prettify(s)
}

func (s StartConfigRuleEvaluationResponse) GoString() string {
	return s.String()
}

func (s *StartConfigRuleEvaluationResponse) SetHeaders(v map[string]*string) *StartConfigRuleEvaluationResponse {
	s.Headers = v
	return s
}

func (s *StartConfigRuleEvaluationResponse) SetStatusCode(v int32) *StartConfigRuleEvaluationResponse {
	s.StatusCode = &v
	return s
}

func (s *StartConfigRuleEvaluationResponse) SetBody(v *StartConfigRuleEvaluationResponseBody) *StartConfigRuleEvaluationResponse {
	s.Body = v
	return s
}

type StartConfigurationRecorderRequest struct {
	// Specifies whether to upgrade Cloud Config for Enterprise. Default value: false. Valid values:
	//
	// *   true
	// *   false
	//
	// >  Cloud Config for Enterprise is upgraded to the account group feature. For more information, see [Announcement: Enterprise Edition Cloud Config Upgrade to Account Group](~~213433~~).
	EnterpriseEdition *bool `json:"EnterpriseEdition,omitempty" xml:"EnterpriseEdition,omitempty"`
}

func (s StartConfigurationRecorderRequest) String() string {
	return tea.Prettify(s)
}

func (s StartConfigurationRecorderRequest) GoString() string {
	return s.String()
}

func (s *StartConfigurationRecorderRequest) SetEnterpriseEdition(v bool) *StartConfigurationRecorderRequest {
	s.EnterpriseEdition = &v
	return s
}

type StartConfigurationRecorderResponseBody struct {
	// The details of the configuration recorder that monitors resources.
	ConfigurationRecorder *StartConfigurationRecorderResponseBodyConfigurationRecorder `json:"ConfigurationRecorder,omitempty" xml:"ConfigurationRecorder,omitempty" type:"Struct"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StartConfigurationRecorderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StartConfigurationRecorderResponseBody) GoString() string {
	return s.String()
}

func (s *StartConfigurationRecorderResponseBody) SetConfigurationRecorder(v *StartConfigurationRecorderResponseBodyConfigurationRecorder) *StartConfigurationRecorderResponseBody {
	s.ConfigurationRecorder = v
	return s
}

func (s *StartConfigurationRecorderResponseBody) SetRequestId(v string) *StartConfigurationRecorderResponseBody {
	s.RequestId = &v
	return s
}

type StartConfigurationRecorderResponseBodyConfigurationRecorder struct {
	// The ID of your Alibaba Cloud account.
	//
	// >  This parameter is returned only when the operation was called by an ordinary account.
	AccountId *int64 `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The status of the configuration recorder. Valid values:
	//
	// *   REGISTRABLE: The configuration recorder has not been registered.
	// *   BUILDING: The configuration recorder is being deployed.
	// *   REGISTERED: The configuration recorder has been registered.
	// *   REBUILDING: The configuration recorder is being redeployed.
	ConfigurationRecorderStatus *string `json:"ConfigurationRecorderStatus,omitempty" xml:"ConfigurationRecorderStatus,omitempty"`
	// The upgrade status of Cloud Config for Enterprise. Valid values:
	//
	// *   REGISTRABLE: Cloud Config for Enterprise is not upgraded.
	// *   BUILDING: Cloud Config for Enterprise is being upgraded.
	// *   REGISTERED: Cloud Config for Enterprise is upgraded.
	OrganizationEnableStatus *string `json:"OrganizationEnableStatus,omitempty" xml:"OrganizationEnableStatus,omitempty"`
	// The ID of the management account.
	//
	// >  This parameter is returned only when the operation was called by a management account.
	OrganizationMasterId *int64 `json:"OrganizationMasterId,omitempty" xml:"OrganizationMasterId,omitempty"`
	// The types of resources that are monitored by Cloud Config.
	//
	// >  By default, Cloud Config returns all resource types. In the following sample responses, `ACS::ECS::Instance` and `ACS::ECS::NetworkInterface` are returned.
	ResourceTypes []*string `json:"ResourceTypes,omitempty" xml:"ResourceTypes,omitempty" type:"Repeated"`
}

func (s StartConfigurationRecorderResponseBodyConfigurationRecorder) String() string {
	return tea.Prettify(s)
}

func (s StartConfigurationRecorderResponseBodyConfigurationRecorder) GoString() string {
	return s.String()
}

func (s *StartConfigurationRecorderResponseBodyConfigurationRecorder) SetAccountId(v int64) *StartConfigurationRecorderResponseBodyConfigurationRecorder {
	s.AccountId = &v
	return s
}

func (s *StartConfigurationRecorderResponseBodyConfigurationRecorder) SetConfigurationRecorderStatus(v string) *StartConfigurationRecorderResponseBodyConfigurationRecorder {
	s.ConfigurationRecorderStatus = &v
	return s
}

func (s *StartConfigurationRecorderResponseBodyConfigurationRecorder) SetOrganizationEnableStatus(v string) *StartConfigurationRecorderResponseBodyConfigurationRecorder {
	s.OrganizationEnableStatus = &v
	return s
}

func (s *StartConfigurationRecorderResponseBodyConfigurationRecorder) SetOrganizationMasterId(v int64) *StartConfigurationRecorderResponseBodyConfigurationRecorder {
	s.OrganizationMasterId = &v
	return s
}

func (s *StartConfigurationRecorderResponseBodyConfigurationRecorder) SetResourceTypes(v []*string) *StartConfigurationRecorderResponseBodyConfigurationRecorder {
	s.ResourceTypes = v
	return s
}

type StartConfigurationRecorderResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *StartConfigurationRecorderResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StartConfigurationRecorderResponse) String() string {
	return tea.Prettify(s)
}

func (s StartConfigurationRecorderResponse) GoString() string {
	return s.String()
}

func (s *StartConfigurationRecorderResponse) SetHeaders(v map[string]*string) *StartConfigurationRecorderResponse {
	s.Headers = v
	return s
}

func (s *StartConfigurationRecorderResponse) SetStatusCode(v int32) *StartConfigurationRecorderResponse {
	s.StatusCode = &v
	return s
}

func (s *StartConfigurationRecorderResponse) SetBody(v *StartConfigurationRecorderResponseBody) *StartConfigurationRecorderResponse {
	s.Body = v
	return s
}

type StopConfigRulesRequest struct {
	// The IDs of the rules. Separate multiple rule IDs with commas (,). You can specify a maximum of 20 rule IDs at a time.
	ConfigRuleIds *string `json:"ConfigRuleIds,omitempty" xml:"ConfigRuleIds,omitempty"`
}

func (s StopConfigRulesRequest) String() string {
	return tea.Prettify(s)
}

func (s StopConfigRulesRequest) GoString() string {
	return s.String()
}

func (s *StopConfigRulesRequest) SetConfigRuleIds(v string) *StopConfigRulesRequest {
	s.ConfigRuleIds = &v
	return s
}

type StopConfigRulesResponseBody struct {
	// The result information about the operation.
	OperateRuleResult *StopConfigRulesResponseBodyOperateRuleResult `json:"OperateRuleResult,omitempty" xml:"OperateRuleResult,omitempty" type:"Struct"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StopConfigRulesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StopConfigRulesResponseBody) GoString() string {
	return s.String()
}

func (s *StopConfigRulesResponseBody) SetOperateRuleResult(v *StopConfigRulesResponseBodyOperateRuleResult) *StopConfigRulesResponseBody {
	s.OperateRuleResult = v
	return s
}

func (s *StopConfigRulesResponseBody) SetRequestId(v string) *StopConfigRulesResponseBody {
	s.RequestId = &v
	return s
}

type StopConfigRulesResponseBodyOperateRuleResult struct {
	// The result information about each rule.
	OperateRuleItemList []*StopConfigRulesResponseBodyOperateRuleResultOperateRuleItemList `json:"OperateRuleItemList,omitempty" xml:"OperateRuleItemList,omitempty" type:"Repeated"`
}

func (s StopConfigRulesResponseBodyOperateRuleResult) String() string {
	return tea.Prettify(s)
}

func (s StopConfigRulesResponseBodyOperateRuleResult) GoString() string {
	return s.String()
}

func (s *StopConfigRulesResponseBodyOperateRuleResult) SetOperateRuleItemList(v []*StopConfigRulesResponseBodyOperateRuleResultOperateRuleItemList) *StopConfigRulesResponseBodyOperateRuleResult {
	s.OperateRuleItemList = v
	return s
}

type StopConfigRulesResponseBodyOperateRuleResultOperateRuleItemList struct {
	// The ID of the rule.
	ConfigRuleId *string `json:"ConfigRuleId,omitempty" xml:"ConfigRuleId,omitempty"`
	// The error code.
	//
	// *   If the rule is disabled, no error code is returned.
	// *   If you fail to disable the rule, an error code is returned. For more information, visit the [API Error Center](https://error-center.alibabacloud.com/status/product/Config).
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// Indicates whether the rule is disabled. Valid values:
	//
	// *   true
	// *   false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s StopConfigRulesResponseBodyOperateRuleResultOperateRuleItemList) String() string {
	return tea.Prettify(s)
}

func (s StopConfigRulesResponseBodyOperateRuleResultOperateRuleItemList) GoString() string {
	return s.String()
}

func (s *StopConfigRulesResponseBodyOperateRuleResultOperateRuleItemList) SetConfigRuleId(v string) *StopConfigRulesResponseBodyOperateRuleResultOperateRuleItemList {
	s.ConfigRuleId = &v
	return s
}

func (s *StopConfigRulesResponseBodyOperateRuleResultOperateRuleItemList) SetErrorCode(v string) *StopConfigRulesResponseBodyOperateRuleResultOperateRuleItemList {
	s.ErrorCode = &v
	return s
}

func (s *StopConfigRulesResponseBodyOperateRuleResultOperateRuleItemList) SetSuccess(v bool) *StopConfigRulesResponseBodyOperateRuleResultOperateRuleItemList {
	s.Success = &v
	return s
}

type StopConfigRulesResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *StopConfigRulesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StopConfigRulesResponse) String() string {
	return tea.Prettify(s)
}

func (s StopConfigRulesResponse) GoString() string {
	return s.String()
}

func (s *StopConfigRulesResponse) SetHeaders(v map[string]*string) *StopConfigRulesResponse {
	s.Headers = v
	return s
}

func (s *StopConfigRulesResponse) SetStatusCode(v int32) *StopConfigRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *StopConfigRulesResponse) SetBody(v *StopConfigRulesResponseBody) *StopConfigRulesResponse {
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
	client.EndpointRule = tea.String("central")
	client.EndpointMap = map[string]*string{
		"cn-shanghai":    tea.String("config.cn-shanghai.aliyuncs.com"),
		"ap-southeast-1": tea.String("config.ap-southeast-1.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("config"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

/**
 * @deprecated : ActiveConfigRules is deprecated, please use Config::2020-09-07::ActiveConfigRules instead.
 *
 * @param request ActiveConfigRulesRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return ActiveConfigRulesResponse
 */
// Deprecated
func (client *Client) ActiveConfigRulesWithOptions(request *ActiveConfigRulesRequest, runtime *util.RuntimeOptions) (_result *ActiveConfigRulesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConfigRuleIds)) {
		query["ConfigRuleIds"] = request.ConfigRuleIds
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ActiveConfigRules"),
		Version:     tea.String("2019-01-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ActiveConfigRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * @deprecated : ActiveConfigRules is deprecated, please use Config::2020-09-07::ActiveConfigRules instead.
 *
 * @param request ActiveConfigRulesRequest
 * @return ActiveConfigRulesResponse
 */
// Deprecated
func (client *Client) ActiveConfigRules(request *ActiveConfigRulesRequest) (_result *ActiveConfigRulesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ActiveConfigRulesResponse{}
	_body, _err := client.ActiveConfigRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteConfigRulesWithOptions(request *DeleteConfigRulesRequest, runtime *util.RuntimeOptions) (_result *DeleteConfigRulesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConfigRuleIds)) {
		query["ConfigRuleIds"] = request.ConfigRuleIds
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteConfigRules"),
		Version:     tea.String("2019-01-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteConfigRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteConfigRules(request *DeleteConfigRulesRequest) (_result *DeleteConfigRulesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteConfigRulesResponse{}
	_body, _err := client.DeleteConfigRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeComplianceWithOptions(request *DescribeComplianceRequest, runtime *util.RuntimeOptions) (_result *DescribeComplianceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ComplianceType)) {
		query["ComplianceType"] = request.ComplianceType
	}

	if !tea.BoolValue(util.IsUnset(request.ConfigRuleId)) {
		query["ConfigRuleId"] = request.ConfigRuleId
	}

	if !tea.BoolValue(util.IsUnset(request.MemberId)) {
		query["MemberId"] = request.MemberId
	}

	if !tea.BoolValue(util.IsUnset(request.MultiAccount)) {
		query["MultiAccount"] = request.MultiAccount
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
		Action:      tea.String("DescribeCompliance"),
		Version:     tea.String("2019-01-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeComplianceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeCompliance(request *DescribeComplianceRequest) (_result *DescribeComplianceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeComplianceResponse{}
	_body, _err := client.DescribeComplianceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeComplianceSummaryWithOptions(request *DescribeComplianceSummaryRequest, runtime *util.RuntimeOptions) (_result *DescribeComplianceSummaryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MemberId)) {
		query["MemberId"] = request.MemberId
	}

	if !tea.BoolValue(util.IsUnset(request.MultiAccount)) {
		query["MultiAccount"] = request.MultiAccount
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeComplianceSummary"),
		Version:     tea.String("2019-01-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeComplianceSummaryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeComplianceSummary(request *DescribeComplianceSummaryRequest) (_result *DescribeComplianceSummaryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeComplianceSummaryResponse{}
	_body, _err := client.DescribeComplianceSummaryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeConfigRuleWithOptions(request *DescribeConfigRuleRequest, runtime *util.RuntimeOptions) (_result *DescribeConfigRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConfigRuleId)) {
		query["ConfigRuleId"] = request.ConfigRuleId
	}

	if !tea.BoolValue(util.IsUnset(request.MemberId)) {
		query["MemberId"] = request.MemberId
	}

	if !tea.BoolValue(util.IsUnset(request.MultiAccount)) {
		query["MultiAccount"] = request.MultiAccount
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeConfigRule"),
		Version:     tea.String("2019-01-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeConfigRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeConfigRule(request *DescribeConfigRuleRequest) (_result *DescribeConfigRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeConfigRuleResponse{}
	_body, _err := client.DescribeConfigRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * In this example, the Alibaba Cloud account `100931896542****` is used. The return result shows that two types of resources are monitored: `ACS::ECS::Instance` and `ACS::ECS::NetworkInterface`. The status of the configuration recorder is `REGISTERED`.
 *
 * @param request DescribeConfigurationRecorderRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DescribeConfigurationRecorderResponse
 */
func (client *Client) DescribeConfigurationRecorderWithOptions(runtime *util.RuntimeOptions) (_result *DescribeConfigurationRecorderResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("DescribeConfigurationRecorder"),
		Version:     tea.String("2019-01-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeConfigurationRecorderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * In this example, the Alibaba Cloud account `100931896542****` is used. The return result shows that two types of resources are monitored: `ACS::ECS::Instance` and `ACS::ECS::NetworkInterface`. The status of the configuration recorder is `REGISTERED`.
 *
 * @return DescribeConfigurationRecorderResponse
 */
func (client *Client) DescribeConfigurationRecorder() (_result *DescribeConfigurationRecorderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeConfigurationRecorderResponse{}
	_body, _err := client.DescribeConfigurationRecorderWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * @deprecated : DescribeDeliveryChannels is deprecated, please use Config::2020-09-07::ListConfigDeliveryChannels,Config::2020-09-07::ListAggregateConfigDeliveryChannels instead.
 *
 * @param request DescribeDeliveryChannelsRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DescribeDeliveryChannelsResponse
 */
// Deprecated
func (client *Client) DescribeDeliveryChannelsWithOptions(request *DescribeDeliveryChannelsRequest, runtime *util.RuntimeOptions) (_result *DescribeDeliveryChannelsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeliveryChannelIds)) {
		query["DeliveryChannelIds"] = request.DeliveryChannelIds
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDeliveryChannels"),
		Version:     tea.String("2019-01-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDeliveryChannelsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * @deprecated : DescribeDeliveryChannels is deprecated, please use Config::2020-09-07::ListConfigDeliveryChannels,Config::2020-09-07::ListAggregateConfigDeliveryChannels instead.
 *
 * @param request DescribeDeliveryChannelsRequest
 * @return DescribeDeliveryChannelsResponse
 */
// Deprecated
func (client *Client) DescribeDeliveryChannels(request *DescribeDeliveryChannelsRequest) (_result *DescribeDeliveryChannelsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDeliveryChannelsResponse{}
	_body, _err := client.DescribeDeliveryChannelsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDiscoveredResourceWithOptions(request *DescribeDiscoveredResourceRequest, runtime *util.RuntimeOptions) (_result *DescribeDiscoveredResourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MemberId)) {
		query["MemberId"] = request.MemberId
	}

	if !tea.BoolValue(util.IsUnset(request.MultiAccount)) {
		query["MultiAccount"] = request.MultiAccount
	}

	if !tea.BoolValue(util.IsUnset(request.Region)) {
		query["Region"] = request.Region
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
		Action:      tea.String("DescribeDiscoveredResource"),
		Version:     tea.String("2019-01-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDiscoveredResourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDiscoveredResource(request *DescribeDiscoveredResourceRequest) (_result *DescribeDiscoveredResourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDiscoveredResourceResponse{}
	_body, _err := client.DescribeDiscoveredResourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeEvaluationResultsWithOptions(request *DescribeEvaluationResultsRequest, runtime *util.RuntimeOptions) (_result *DescribeEvaluationResultsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeEvaluationResults"),
		Version:     tea.String("2019-01-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeEvaluationResultsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeEvaluationResults(request *DescribeEvaluationResultsRequest) (_result *DescribeEvaluationResultsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeEvaluationResultsResponse{}
	_body, _err := client.DescribeEvaluationResultsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * @deprecated : GetAggregateDiscoveredResource is deprecated, please use Config::2020-09-07::GetAggregateDiscoveredResource instead.
 * In this topic, the `new-bucket` resource in the `ca-5885626622af0008****` account group is used as an example.
 *
 * @param request GetAggregateDiscoveredResourceRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return GetAggregateDiscoveredResourceResponse
 */
// Deprecated
func (client *Client) GetAggregateDiscoveredResourceWithOptions(request *GetAggregateDiscoveredResourceRequest, runtime *util.RuntimeOptions) (_result *GetAggregateDiscoveredResourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetAggregateDiscoveredResource"),
		Version:     tea.String("2019-01-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAggregateDiscoveredResourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * @deprecated : GetAggregateDiscoveredResource is deprecated, please use Config::2020-09-07::GetAggregateDiscoveredResource instead.
 * In this topic, the `new-bucket` resource in the `ca-5885626622af0008****` account group is used as an example.
 *
 * @param request GetAggregateDiscoveredResourceRequest
 * @return GetAggregateDiscoveredResourceResponse
 */
// Deprecated
func (client *Client) GetAggregateDiscoveredResource(request *GetAggregateDiscoveredResourceRequest) (_result *GetAggregateDiscoveredResourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAggregateDiscoveredResourceResponse{}
	_body, _err := client.GetAggregateDiscoveredResourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetDiscoveredResourceCountsWithOptions(request *GetDiscoveredResourceCountsRequest, runtime *util.RuntimeOptions) (_result *GetDiscoveredResourceCountsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GroupByKey)) {
		query["GroupByKey"] = request.GroupByKey
	}

	if !tea.BoolValue(util.IsUnset(request.MemberId)) {
		query["MemberId"] = request.MemberId
	}

	if !tea.BoolValue(util.IsUnset(request.MultiAccount)) {
		query["MultiAccount"] = request.MultiAccount
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetDiscoveredResourceCounts"),
		Version:     tea.String("2019-01-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetDiscoveredResourceCountsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetDiscoveredResourceCounts(request *GetDiscoveredResourceCountsRequest) (_result *GetDiscoveredResourceCountsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetDiscoveredResourceCountsResponse{}
	_body, _err := client.GetDiscoveredResourceCountsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetDiscoveredResourceSummaryWithOptions(request *GetDiscoveredResourceSummaryRequest, runtime *util.RuntimeOptions) (_result *GetDiscoveredResourceSummaryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MemberId)) {
		query["MemberId"] = request.MemberId
	}

	if !tea.BoolValue(util.IsUnset(request.MultiAccount)) {
		query["MultiAccount"] = request.MultiAccount
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetDiscoveredResourceSummary"),
		Version:     tea.String("2019-01-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetDiscoveredResourceSummaryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetDiscoveredResourceSummary(request *GetDiscoveredResourceSummaryRequest) (_result *GetDiscoveredResourceSummaryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetDiscoveredResourceSummaryResponse{}
	_body, _err := client.GetDiscoveredResourceSummaryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetResourceComplianceTimelineWithOptions(request *GetResourceComplianceTimelineRequest, runtime *util.RuntimeOptions) (_result *GetResourceComplianceTimelineResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Limit)) {
		query["Limit"] = request.Limit
	}

	if !tea.BoolValue(util.IsUnset(request.MemberId)) {
		query["MemberId"] = request.MemberId
	}

	if !tea.BoolValue(util.IsUnset(request.MultiAccount)) {
		query["MultiAccount"] = request.MultiAccount
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.Region)) {
		query["Region"] = request.Region
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetResourceComplianceTimeline"),
		Version:     tea.String("2019-01-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetResourceComplianceTimelineResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetResourceComplianceTimeline(request *GetResourceComplianceTimelineRequest) (_result *GetResourceComplianceTimelineResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetResourceComplianceTimelineResponse{}
	_body, _err := client.GetResourceComplianceTimelineWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetResourceConfigurationTimelineWithOptions(request *GetResourceConfigurationTimelineRequest, runtime *util.RuntimeOptions) (_result *GetResourceConfigurationTimelineResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Limit)) {
		query["Limit"] = request.Limit
	}

	if !tea.BoolValue(util.IsUnset(request.MemberId)) {
		query["MemberId"] = request.MemberId
	}

	if !tea.BoolValue(util.IsUnset(request.MultiAccount)) {
		query["MultiAccount"] = request.MultiAccount
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.Region)) {
		query["Region"] = request.Region
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetResourceConfigurationTimeline"),
		Version:     tea.String("2019-01-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetResourceConfigurationTimelineResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetResourceConfigurationTimeline(request *GetResourceConfigurationTimelineRequest) (_result *GetResourceConfigurationTimelineResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetResourceConfigurationTimelineResponse{}
	_body, _err := client.GetResourceConfigurationTimelineWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetSupportedResourceTypesWithOptions(runtime *util.RuntimeOptions) (_result *GetSupportedResourceTypesResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("GetSupportedResourceTypes"),
		Version:     tea.String("2019-01-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetSupportedResourceTypesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetSupportedResourceTypes() (_result *GetSupportedResourceTypesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetSupportedResourceTypesResponse{}
	_body, _err := client.GetSupportedResourceTypesWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * @deprecated : ListAggregateDiscoveredResources is deprecated, please use Config::2020-09-07::ListAggregateDiscoveredResources instead.
 * In this example, the `ca-c560626622af0005****` account group is used. The return result shows a total of eight resources in this sample account group.
 *
 * @param request ListAggregateDiscoveredResourcesRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return ListAggregateDiscoveredResourcesResponse
 */
// Deprecated
func (client *Client) ListAggregateDiscoveredResourcesWithOptions(request *ListAggregateDiscoveredResourcesRequest, runtime *util.RuntimeOptions) (_result *ListAggregateDiscoveredResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AggregatorId)) {
		query["AggregatorId"] = request.AggregatorId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Regions)) {
		query["Regions"] = request.Regions
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceDeleted)) {
		query["ResourceDeleted"] = request.ResourceDeleted
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceTypes)) {
		query["ResourceTypes"] = request.ResourceTypes
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListAggregateDiscoveredResources"),
		Version:     tea.String("2019-01-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListAggregateDiscoveredResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * @deprecated : ListAggregateDiscoveredResources is deprecated, please use Config::2020-09-07::ListAggregateDiscoveredResources instead.
 * In this example, the `ca-c560626622af0005****` account group is used. The return result shows a total of eight resources in this sample account group.
 *
 * @param request ListAggregateDiscoveredResourcesRequest
 * @return ListAggregateDiscoveredResourcesResponse
 */
// Deprecated
func (client *Client) ListAggregateDiscoveredResources(request *ListAggregateDiscoveredResourcesRequest) (_result *ListAggregateDiscoveredResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListAggregateDiscoveredResourcesResponse{}
	_body, _err := client.ListAggregateDiscoveredResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListConfigRulesWithOptions(request *ListConfigRulesRequest, runtime *util.RuntimeOptions) (_result *ListConfigRulesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CompliancePackId)) {
		query["CompliancePackId"] = request.CompliancePackId
	}

	if !tea.BoolValue(util.IsUnset(request.ComplianceType)) {
		query["ComplianceType"] = request.ComplianceType
	}

	if !tea.BoolValue(util.IsUnset(request.ConfigRuleName)) {
		query["ConfigRuleName"] = request.ConfigRuleName
	}

	if !tea.BoolValue(util.IsUnset(request.ConfigRuleState)) {
		query["ConfigRuleState"] = request.ConfigRuleState
	}

	if !tea.BoolValue(util.IsUnset(request.MemberId)) {
		query["MemberId"] = request.MemberId
	}

	if !tea.BoolValue(util.IsUnset(request.MultiAccount)) {
		query["MultiAccount"] = request.MultiAccount
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RiskLevel)) {
		query["RiskLevel"] = request.RiskLevel
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListConfigRules"),
		Version:     tea.String("2019-01-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListConfigRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListConfigRules(request *ListConfigRulesRequest) (_result *ListConfigRulesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListConfigRulesResponse{}
	_body, _err := client.ListConfigRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListDiscoveredResourcesWithOptions(request *ListDiscoveredResourcesRequest, runtime *util.RuntimeOptions) (_result *ListDiscoveredResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MemberId)) {
		query["MemberId"] = request.MemberId
	}

	if !tea.BoolValue(util.IsUnset(request.MultiAccount)) {
		query["MultiAccount"] = request.MultiAccount
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Regions)) {
		query["Regions"] = request.Regions
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceDeleted)) {
		query["ResourceDeleted"] = request.ResourceDeleted
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceTypes)) {
		query["ResourceTypes"] = request.ResourceTypes
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListDiscoveredResources"),
		Version:     tea.String("2019-01-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListDiscoveredResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListDiscoveredResources(request *ListDiscoveredResourcesRequest) (_result *ListDiscoveredResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListDiscoveredResourcesResponse{}
	_body, _err := client.ListDiscoveredResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * @deprecated : ListRemediationTemplates is deprecated, please use Config::2020-09-07::ListRemediationTemplates instead.
 * In this topic, the `oss-bucket-public-write-prohibited` managed rule is used as an example. The return result shows the detailed information about the remediation template of the `OOS` type for the managed rule. OOS represents Operation Orchestration Service.
 *
 * @param request ListRemediationTemplatesRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return ListRemediationTemplatesResponse
 */
// Deprecated
func (client *Client) ListRemediationTemplatesWithOptions(request *ListRemediationTemplatesRequest, runtime *util.RuntimeOptions) (_result *ListRemediationTemplatesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ManagedRuleIdentifier)) {
		query["ManagedRuleIdentifier"] = request.ManagedRuleIdentifier
	}

	if !tea.BoolValue(util.IsUnset(request.RemediationType)) {
		query["RemediationType"] = request.RemediationType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListRemediationTemplates"),
		Version:     tea.String("2019-01-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListRemediationTemplatesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * @deprecated : ListRemediationTemplates is deprecated, please use Config::2020-09-07::ListRemediationTemplates instead.
 * In this topic, the `oss-bucket-public-write-prohibited` managed rule is used as an example. The return result shows the detailed information about the remediation template of the `OOS` type for the managed rule. OOS represents Operation Orchestration Service.
 *
 * @param request ListRemediationTemplatesRequest
 * @return ListRemediationTemplatesResponse
 */
// Deprecated
func (client *Client) ListRemediationTemplates(request *ListRemediationTemplatesRequest) (_result *ListRemediationTemplatesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListRemediationTemplatesResponse{}
	_body, _err := client.ListRemediationTemplatesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * @deprecated : PutConfigRule is deprecated, please use Config::2020-09-07::CreateConfigRule,Config::2020-09-07::CreateAggregateConfigRule instead.
 *
 * @param request PutConfigRuleRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return PutConfigRuleResponse
 */
// Deprecated
func (client *Client) PutConfigRuleWithOptions(request *PutConfigRuleRequest, runtime *util.RuntimeOptions) (_result *PutConfigRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MemberId)) {
		query["MemberId"] = request.MemberId
	}

	if !tea.BoolValue(util.IsUnset(request.MultiAccount)) {
		query["MultiAccount"] = request.MultiAccount
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.ConfigRuleId)) {
		body["ConfigRuleId"] = request.ConfigRuleId
	}

	if !tea.BoolValue(util.IsUnset(request.ConfigRuleName)) {
		body["ConfigRuleName"] = request.ConfigRuleName
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.InputParameters)) {
		body["InputParameters"] = request.InputParameters
	}

	if !tea.BoolValue(util.IsUnset(request.RiskLevel)) {
		body["RiskLevel"] = request.RiskLevel
	}

	if !tea.BoolValue(util.IsUnset(request.ScopeComplianceResourceId)) {
		body["ScopeComplianceResourceId"] = request.ScopeComplianceResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ScopeComplianceResourceTypes)) {
		body["ScopeComplianceResourceTypes"] = request.ScopeComplianceResourceTypes
	}

	if !tea.BoolValue(util.IsUnset(request.SourceDetailMessageType)) {
		body["SourceDetailMessageType"] = request.SourceDetailMessageType
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIdentifier)) {
		body["SourceIdentifier"] = request.SourceIdentifier
	}

	if !tea.BoolValue(util.IsUnset(request.SourceMaximumExecutionFrequency)) {
		body["SourceMaximumExecutionFrequency"] = request.SourceMaximumExecutionFrequency
	}

	if !tea.BoolValue(util.IsUnset(request.SourceOwner)) {
		body["SourceOwner"] = request.SourceOwner
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("PutConfigRule"),
		Version:     tea.String("2019-01-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &PutConfigRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * @deprecated : PutConfigRule is deprecated, please use Config::2020-09-07::CreateConfigRule,Config::2020-09-07::CreateAggregateConfigRule instead.
 *
 * @param request PutConfigRuleRequest
 * @return PutConfigRuleResponse
 */
// Deprecated
func (client *Client) PutConfigRule(request *PutConfigRuleRequest) (_result *PutConfigRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PutConfigRuleResponse{}
	_body, _err := client.PutConfigRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * @deprecated : PutConfigurationRecorder is deprecated, please use Config::2020-09-07::UpdateConfigurationRecorder instead.
 * This operation is valid only for the current Alibaba Cloud account. For more information, see [Alibaba Cloud services that support Cloud Config](~~127411~~).
 *
 * @param request PutConfigurationRecorderRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return PutConfigurationRecorderResponse
 */
// Deprecated
func (client *Client) PutConfigurationRecorderWithOptions(request *PutConfigurationRecorderRequest, runtime *util.RuntimeOptions) (_result *PutConfigurationRecorderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ResourceTypes)) {
		body["ResourceTypes"] = request.ResourceTypes
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("PutConfigurationRecorder"),
		Version:     tea.String("2019-01-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &PutConfigurationRecorderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * @deprecated : PutConfigurationRecorder is deprecated, please use Config::2020-09-07::UpdateConfigurationRecorder instead.
 * This operation is valid only for the current Alibaba Cloud account. For more information, see [Alibaba Cloud services that support Cloud Config](~~127411~~).
 *
 * @param request PutConfigurationRecorderRequest
 * @return PutConfigurationRecorderResponse
 */
// Deprecated
func (client *Client) PutConfigurationRecorder(request *PutConfigurationRecorderRequest) (_result *PutConfigurationRecorderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PutConfigurationRecorderResponse{}
	_body, _err := client.PutConfigurationRecorderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * @deprecated : PutDeliveryChannel is deprecated, please use Config::2020-09-07::CreateConfigDeliveryChannel,Config::2020-09-07::CreateAggregateConfigDeliveryChannel instead.
 * ****
 *
 * @param request PutDeliveryChannelRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return PutDeliveryChannelResponse
 */
// Deprecated
func (client *Client) PutDeliveryChannelWithOptions(request *PutDeliveryChannelRequest, runtime *util.RuntimeOptions) (_result *PutDeliveryChannelResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DeliveryChannelAssumeRoleArn)) {
		body["DeliveryChannelAssumeRoleArn"] = request.DeliveryChannelAssumeRoleArn
	}

	if !tea.BoolValue(util.IsUnset(request.DeliveryChannelCondition)) {
		body["DeliveryChannelCondition"] = request.DeliveryChannelCondition
	}

	if !tea.BoolValue(util.IsUnset(request.DeliveryChannelId)) {
		body["DeliveryChannelId"] = request.DeliveryChannelId
	}

	if !tea.BoolValue(util.IsUnset(request.DeliveryChannelName)) {
		body["DeliveryChannelName"] = request.DeliveryChannelName
	}

	if !tea.BoolValue(util.IsUnset(request.DeliveryChannelTargetArn)) {
		body["DeliveryChannelTargetArn"] = request.DeliveryChannelTargetArn
	}

	if !tea.BoolValue(util.IsUnset(request.DeliveryChannelType)) {
		body["DeliveryChannelType"] = request.DeliveryChannelType
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		body["Status"] = request.Status
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("PutDeliveryChannel"),
		Version:     tea.String("2019-01-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &PutDeliveryChannelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * @deprecated : PutDeliveryChannel is deprecated, please use Config::2020-09-07::CreateConfigDeliveryChannel,Config::2020-09-07::CreateAggregateConfigDeliveryChannel instead.
 * ****
 *
 * @param request PutDeliveryChannelRequest
 * @return PutDeliveryChannelResponse
 */
// Deprecated
func (client *Client) PutDeliveryChannel(request *PutDeliveryChannelRequest) (_result *PutDeliveryChannelResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PutDeliveryChannelResponse{}
	_body, _err := client.PutDeliveryChannelWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * @deprecated : PutEvaluations is deprecated, please use Config::2020-09-07::PutEvaluations instead.
 * ****
 *
 * @param request PutEvaluationsRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return PutEvaluationsResponse
 */
// Deprecated
func (client *Client) PutEvaluationsWithOptions(request *PutEvaluationsRequest, runtime *util.RuntimeOptions) (_result *PutEvaluationsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Evaluations)) {
		body["Evaluations"] = request.Evaluations
	}

	if !tea.BoolValue(util.IsUnset(request.ResultToken)) {
		body["ResultToken"] = request.ResultToken
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("PutEvaluations"),
		Version:     tea.String("2019-01-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &PutEvaluationsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * @deprecated : PutEvaluations is deprecated, please use Config::2020-09-07::PutEvaluations instead.
 * ****
 *
 * @param request PutEvaluationsRequest
 * @return PutEvaluationsResponse
 */
// Deprecated
func (client *Client) PutEvaluations(request *PutEvaluationsRequest) (_result *PutEvaluationsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PutEvaluationsResponse{}
	_body, _err := client.PutEvaluationsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * >  After you call this operation, the compliance evaluation is performed only once. To query the compliance evaluation results returned by the specified rule or compliance package, call the ListConfigRuleEvaluationResults operation. For more information, see [ListConfigRuleEvaluationResults](~~267503~~).
 * In this example, the `cr-bdb6626622af0012****` rule is used.
 *
 * @param request StartConfigRuleEvaluationRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return StartConfigRuleEvaluationResponse
 */
func (client *Client) StartConfigRuleEvaluationWithOptions(request *StartConfigRuleEvaluationRequest, runtime *util.RuntimeOptions) (_result *StartConfigRuleEvaluationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CompliancePackId)) {
		query["CompliancePackId"] = request.CompliancePackId
	}

	if !tea.BoolValue(util.IsUnset(request.ConfigRuleId)) {
		query["ConfigRuleId"] = request.ConfigRuleId
	}

	if !tea.BoolValue(util.IsUnset(request.RevertEvaluation)) {
		query["RevertEvaluation"] = request.RevertEvaluation
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("StartConfigRuleEvaluation"),
		Version:     tea.String("2019-01-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &StartConfigRuleEvaluationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * >  After you call this operation, the compliance evaluation is performed only once. To query the compliance evaluation results returned by the specified rule or compliance package, call the ListConfigRuleEvaluationResults operation. For more information, see [ListConfigRuleEvaluationResults](~~267503~~).
 * In this example, the `cr-bdb6626622af0012****` rule is used.
 *
 * @param request StartConfigRuleEvaluationRequest
 * @return StartConfigRuleEvaluationResponse
 */
func (client *Client) StartConfigRuleEvaluation(request *StartConfigRuleEvaluationRequest) (_result *StartConfigRuleEvaluationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StartConfigRuleEvaluationResponse{}
	_body, _err := client.StartConfigRuleEvaluationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Ordinary accounts and management accounts have the following differences when you use them to call the StartConfigurationRecorder operation:
 * *   If you are using an ordinary account, you can call the StartConfigurationRecorder operation to activate Cloud Config for the current account. You can also view the resources of the current account and manage the compliance rules of the current account.
 * *   If you are using the management account of a resource directory and set the `EnterpriseEdition` parameter to `true`, you can activate Cloud Config for all member accounts. In addition, a global account group that includes all member accounts is created by default. You can view the resources of all member accounts in the account group and manage the compliance rules of all member accounts.
 * In this topic, an ordinary account is used as an example to show how to activate Cloud Config and monitor the resources of the current account.
 *
 * @param request StartConfigurationRecorderRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return StartConfigurationRecorderResponse
 */
func (client *Client) StartConfigurationRecorderWithOptions(request *StartConfigurationRecorderRequest, runtime *util.RuntimeOptions) (_result *StartConfigurationRecorderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EnterpriseEdition)) {
		body["EnterpriseEdition"] = request.EnterpriseEdition
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("StartConfigurationRecorder"),
		Version:     tea.String("2019-01-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &StartConfigurationRecorderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Ordinary accounts and management accounts have the following differences when you use them to call the StartConfigurationRecorder operation:
 * *   If you are using an ordinary account, you can call the StartConfigurationRecorder operation to activate Cloud Config for the current account. You can also view the resources of the current account and manage the compliance rules of the current account.
 * *   If you are using the management account of a resource directory and set the `EnterpriseEdition` parameter to `true`, you can activate Cloud Config for all member accounts. In addition, a global account group that includes all member accounts is created by default. You can view the resources of all member accounts in the account group and manage the compliance rules of all member accounts.
 * In this topic, an ordinary account is used as an example to show how to activate Cloud Config and monitor the resources of the current account.
 *
 * @param request StartConfigurationRecorderRequest
 * @return StartConfigurationRecorderResponse
 */
func (client *Client) StartConfigurationRecorder(request *StartConfigurationRecorderRequest) (_result *StartConfigurationRecorderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StartConfigurationRecorderResponse{}
	_body, _err := client.StartConfigurationRecorderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StopConfigRulesWithOptions(request *StopConfigRulesRequest, runtime *util.RuntimeOptions) (_result *StopConfigRulesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConfigRuleIds)) {
		query["ConfigRuleIds"] = request.ConfigRuleIds
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("StopConfigRules"),
		Version:     tea.String("2019-01-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &StopConfigRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StopConfigRules(request *StopConfigRulesRequest) (_result *StopConfigRulesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StopConfigRulesResponse{}
	_body, _err := client.StopConfigRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
