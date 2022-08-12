// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type ResourceDirectoryFolderNode struct {
	AccountId      *string                        `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	Children       []*ResourceDirectoryFolderNode `json:"Children,omitempty" xml:"Children,omitempty" type:"Repeated"`
	DisplayName    *string                        `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	FolderId       *string                        `json:"FolderId,omitempty" xml:"FolderId,omitempty"`
	FolderName     *string                        `json:"FolderName,omitempty" xml:"FolderName,omitempty"`
	ParentFolderId *string                        `json:"ParentFolderId,omitempty" xml:"ParentFolderId,omitempty"`
}

func (s ResourceDirectoryFolderNode) String() string {
	return tea.Prettify(s)
}

func (s ResourceDirectoryFolderNode) GoString() string {
	return s.String()
}

func (s *ResourceDirectoryFolderNode) SetAccountId(v string) *ResourceDirectoryFolderNode {
	s.AccountId = &v
	return s
}

func (s *ResourceDirectoryFolderNode) SetChildren(v []*ResourceDirectoryFolderNode) *ResourceDirectoryFolderNode {
	s.Children = v
	return s
}

func (s *ResourceDirectoryFolderNode) SetDisplayName(v string) *ResourceDirectoryFolderNode {
	s.DisplayName = &v
	return s
}

func (s *ResourceDirectoryFolderNode) SetFolderId(v string) *ResourceDirectoryFolderNode {
	s.FolderId = &v
	return s
}

func (s *ResourceDirectoryFolderNode) SetFolderName(v string) *ResourceDirectoryFolderNode {
	s.FolderName = &v
	return s
}

func (s *ResourceDirectoryFolderNode) SetParentFolderId(v string) *ResourceDirectoryFolderNode {
	s.ParentFolderId = &v
	return s
}

type ActiveAggregateConfigRulesRequest struct {
	AggregatorId  *string `json:"AggregatorId,omitempty" xml:"AggregatorId,omitempty"`
	ConfigRuleIds *string `json:"ConfigRuleIds,omitempty" xml:"ConfigRuleIds,omitempty"`
}

func (s ActiveAggregateConfigRulesRequest) String() string {
	return tea.Prettify(s)
}

func (s ActiveAggregateConfigRulesRequest) GoString() string {
	return s.String()
}

func (s *ActiveAggregateConfigRulesRequest) SetAggregatorId(v string) *ActiveAggregateConfigRulesRequest {
	s.AggregatorId = &v
	return s
}

func (s *ActiveAggregateConfigRulesRequest) SetConfigRuleIds(v string) *ActiveAggregateConfigRulesRequest {
	s.ConfigRuleIds = &v
	return s
}

type ActiveAggregateConfigRulesResponseBody struct {
	OperateRuleResult *ActiveAggregateConfigRulesResponseBodyOperateRuleResult `json:"OperateRuleResult,omitempty" xml:"OperateRuleResult,omitempty" type:"Struct"`
	RequestId         *string                                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ActiveAggregateConfigRulesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ActiveAggregateConfigRulesResponseBody) GoString() string {
	return s.String()
}

func (s *ActiveAggregateConfigRulesResponseBody) SetOperateRuleResult(v *ActiveAggregateConfigRulesResponseBodyOperateRuleResult) *ActiveAggregateConfigRulesResponseBody {
	s.OperateRuleResult = v
	return s
}

func (s *ActiveAggregateConfigRulesResponseBody) SetRequestId(v string) *ActiveAggregateConfigRulesResponseBody {
	s.RequestId = &v
	return s
}

type ActiveAggregateConfigRulesResponseBodyOperateRuleResult struct {
	OperateRuleItemList []*ActiveAggregateConfigRulesResponseBodyOperateRuleResultOperateRuleItemList `json:"OperateRuleItemList,omitempty" xml:"OperateRuleItemList,omitempty" type:"Repeated"`
}

func (s ActiveAggregateConfigRulesResponseBodyOperateRuleResult) String() string {
	return tea.Prettify(s)
}

func (s ActiveAggregateConfigRulesResponseBodyOperateRuleResult) GoString() string {
	return s.String()
}

func (s *ActiveAggregateConfigRulesResponseBodyOperateRuleResult) SetOperateRuleItemList(v []*ActiveAggregateConfigRulesResponseBodyOperateRuleResultOperateRuleItemList) *ActiveAggregateConfigRulesResponseBodyOperateRuleResult {
	s.OperateRuleItemList = v
	return s
}

type ActiveAggregateConfigRulesResponseBodyOperateRuleResultOperateRuleItemList struct {
	ConfigRuleId *string `json:"ConfigRuleId,omitempty" xml:"ConfigRuleId,omitempty"`
	ErrorCode    *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ActiveAggregateConfigRulesResponseBodyOperateRuleResultOperateRuleItemList) String() string {
	return tea.Prettify(s)
}

func (s ActiveAggregateConfigRulesResponseBodyOperateRuleResultOperateRuleItemList) GoString() string {
	return s.String()
}

func (s *ActiveAggregateConfigRulesResponseBodyOperateRuleResultOperateRuleItemList) SetConfigRuleId(v string) *ActiveAggregateConfigRulesResponseBodyOperateRuleResultOperateRuleItemList {
	s.ConfigRuleId = &v
	return s
}

func (s *ActiveAggregateConfigRulesResponseBodyOperateRuleResultOperateRuleItemList) SetErrorCode(v string) *ActiveAggregateConfigRulesResponseBodyOperateRuleResultOperateRuleItemList {
	s.ErrorCode = &v
	return s
}

func (s *ActiveAggregateConfigRulesResponseBodyOperateRuleResultOperateRuleItemList) SetSuccess(v bool) *ActiveAggregateConfigRulesResponseBodyOperateRuleResultOperateRuleItemList {
	s.Success = &v
	return s
}

type ActiveAggregateConfigRulesResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ActiveAggregateConfigRulesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ActiveAggregateConfigRulesResponse) String() string {
	return tea.Prettify(s)
}

func (s ActiveAggregateConfigRulesResponse) GoString() string {
	return s.String()
}

func (s *ActiveAggregateConfigRulesResponse) SetHeaders(v map[string]*string) *ActiveAggregateConfigRulesResponse {
	s.Headers = v
	return s
}

func (s *ActiveAggregateConfigRulesResponse) SetStatusCode(v int32) *ActiveAggregateConfigRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *ActiveAggregateConfigRulesResponse) SetBody(v *ActiveAggregateConfigRulesResponseBody) *ActiveAggregateConfigRulesResponse {
	s.Body = v
	return s
}

type AttachAggregateConfigRuleToCompliancePackRequest struct {
	AggregatorId     *string `json:"AggregatorId,omitempty" xml:"AggregatorId,omitempty"`
	CompliancePackId *string `json:"CompliancePackId,omitempty" xml:"CompliancePackId,omitempty"`
	ConfigRuleIds    *string `json:"ConfigRuleIds,omitempty" xml:"ConfigRuleIds,omitempty"`
}

func (s AttachAggregateConfigRuleToCompliancePackRequest) String() string {
	return tea.Prettify(s)
}

func (s AttachAggregateConfigRuleToCompliancePackRequest) GoString() string {
	return s.String()
}

func (s *AttachAggregateConfigRuleToCompliancePackRequest) SetAggregatorId(v string) *AttachAggregateConfigRuleToCompliancePackRequest {
	s.AggregatorId = &v
	return s
}

func (s *AttachAggregateConfigRuleToCompliancePackRequest) SetCompliancePackId(v string) *AttachAggregateConfigRuleToCompliancePackRequest {
	s.CompliancePackId = &v
	return s
}

func (s *AttachAggregateConfigRuleToCompliancePackRequest) SetConfigRuleIds(v string) *AttachAggregateConfigRuleToCompliancePackRequest {
	s.ConfigRuleIds = &v
	return s
}

type AttachAggregateConfigRuleToCompliancePackResponseBody struct {
	OperateRuleResult *AttachAggregateConfigRuleToCompliancePackResponseBodyOperateRuleResult `json:"OperateRuleResult,omitempty" xml:"OperateRuleResult,omitempty" type:"Struct"`
	RequestId         *string                                                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AttachAggregateConfigRuleToCompliancePackResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AttachAggregateConfigRuleToCompliancePackResponseBody) GoString() string {
	return s.String()
}

func (s *AttachAggregateConfigRuleToCompliancePackResponseBody) SetOperateRuleResult(v *AttachAggregateConfigRuleToCompliancePackResponseBodyOperateRuleResult) *AttachAggregateConfigRuleToCompliancePackResponseBody {
	s.OperateRuleResult = v
	return s
}

func (s *AttachAggregateConfigRuleToCompliancePackResponseBody) SetRequestId(v string) *AttachAggregateConfigRuleToCompliancePackResponseBody {
	s.RequestId = &v
	return s
}

type AttachAggregateConfigRuleToCompliancePackResponseBodyOperateRuleResult struct {
	OperateRuleItemList []*AttachAggregateConfigRuleToCompliancePackResponseBodyOperateRuleResultOperateRuleItemList `json:"OperateRuleItemList,omitempty" xml:"OperateRuleItemList,omitempty" type:"Repeated"`
}

func (s AttachAggregateConfigRuleToCompliancePackResponseBodyOperateRuleResult) String() string {
	return tea.Prettify(s)
}

func (s AttachAggregateConfigRuleToCompliancePackResponseBodyOperateRuleResult) GoString() string {
	return s.String()
}

func (s *AttachAggregateConfigRuleToCompliancePackResponseBodyOperateRuleResult) SetOperateRuleItemList(v []*AttachAggregateConfigRuleToCompliancePackResponseBodyOperateRuleResultOperateRuleItemList) *AttachAggregateConfigRuleToCompliancePackResponseBodyOperateRuleResult {
	s.OperateRuleItemList = v
	return s
}

type AttachAggregateConfigRuleToCompliancePackResponseBodyOperateRuleResultOperateRuleItemList struct {
	ConfigRuleId *string `json:"ConfigRuleId,omitempty" xml:"ConfigRuleId,omitempty"`
	ErrorCode    *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AttachAggregateConfigRuleToCompliancePackResponseBodyOperateRuleResultOperateRuleItemList) String() string {
	return tea.Prettify(s)
}

func (s AttachAggregateConfigRuleToCompliancePackResponseBodyOperateRuleResultOperateRuleItemList) GoString() string {
	return s.String()
}

func (s *AttachAggregateConfigRuleToCompliancePackResponseBodyOperateRuleResultOperateRuleItemList) SetConfigRuleId(v string) *AttachAggregateConfigRuleToCompliancePackResponseBodyOperateRuleResultOperateRuleItemList {
	s.ConfigRuleId = &v
	return s
}

func (s *AttachAggregateConfigRuleToCompliancePackResponseBodyOperateRuleResultOperateRuleItemList) SetErrorCode(v string) *AttachAggregateConfigRuleToCompliancePackResponseBodyOperateRuleResultOperateRuleItemList {
	s.ErrorCode = &v
	return s
}

func (s *AttachAggregateConfigRuleToCompliancePackResponseBodyOperateRuleResultOperateRuleItemList) SetSuccess(v bool) *AttachAggregateConfigRuleToCompliancePackResponseBodyOperateRuleResultOperateRuleItemList {
	s.Success = &v
	return s
}

type AttachAggregateConfigRuleToCompliancePackResponse struct {
	Headers    map[string]*string                                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AttachAggregateConfigRuleToCompliancePackResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AttachAggregateConfigRuleToCompliancePackResponse) String() string {
	return tea.Prettify(s)
}

func (s AttachAggregateConfigRuleToCompliancePackResponse) GoString() string {
	return s.String()
}

func (s *AttachAggregateConfigRuleToCompliancePackResponse) SetHeaders(v map[string]*string) *AttachAggregateConfigRuleToCompliancePackResponse {
	s.Headers = v
	return s
}

func (s *AttachAggregateConfigRuleToCompliancePackResponse) SetStatusCode(v int32) *AttachAggregateConfigRuleToCompliancePackResponse {
	s.StatusCode = &v
	return s
}

func (s *AttachAggregateConfigRuleToCompliancePackResponse) SetBody(v *AttachAggregateConfigRuleToCompliancePackResponseBody) *AttachAggregateConfigRuleToCompliancePackResponse {
	s.Body = v
	return s
}

type AttachConfigRuleToCompliancePackRequest struct {
	CompliancePackId *string `json:"CompliancePackId,omitempty" xml:"CompliancePackId,omitempty"`
	ConfigRuleIds    *string `json:"ConfigRuleIds,omitempty" xml:"ConfigRuleIds,omitempty"`
}

func (s AttachConfigRuleToCompliancePackRequest) String() string {
	return tea.Prettify(s)
}

func (s AttachConfigRuleToCompliancePackRequest) GoString() string {
	return s.String()
}

func (s *AttachConfigRuleToCompliancePackRequest) SetCompliancePackId(v string) *AttachConfigRuleToCompliancePackRequest {
	s.CompliancePackId = &v
	return s
}

func (s *AttachConfigRuleToCompliancePackRequest) SetConfigRuleIds(v string) *AttachConfigRuleToCompliancePackRequest {
	s.ConfigRuleIds = &v
	return s
}

type AttachConfigRuleToCompliancePackResponseBody struct {
	OperateRuleResult *AttachConfigRuleToCompliancePackResponseBodyOperateRuleResult `json:"OperateRuleResult,omitempty" xml:"OperateRuleResult,omitempty" type:"Struct"`
	RequestId         *string                                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AttachConfigRuleToCompliancePackResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AttachConfigRuleToCompliancePackResponseBody) GoString() string {
	return s.String()
}

func (s *AttachConfigRuleToCompliancePackResponseBody) SetOperateRuleResult(v *AttachConfigRuleToCompliancePackResponseBodyOperateRuleResult) *AttachConfigRuleToCompliancePackResponseBody {
	s.OperateRuleResult = v
	return s
}

func (s *AttachConfigRuleToCompliancePackResponseBody) SetRequestId(v string) *AttachConfigRuleToCompliancePackResponseBody {
	s.RequestId = &v
	return s
}

type AttachConfigRuleToCompliancePackResponseBodyOperateRuleResult struct {
	OperateRuleItemList []*AttachConfigRuleToCompliancePackResponseBodyOperateRuleResultOperateRuleItemList `json:"OperateRuleItemList,omitempty" xml:"OperateRuleItemList,omitempty" type:"Repeated"`
}

func (s AttachConfigRuleToCompliancePackResponseBodyOperateRuleResult) String() string {
	return tea.Prettify(s)
}

func (s AttachConfigRuleToCompliancePackResponseBodyOperateRuleResult) GoString() string {
	return s.String()
}

func (s *AttachConfigRuleToCompliancePackResponseBodyOperateRuleResult) SetOperateRuleItemList(v []*AttachConfigRuleToCompliancePackResponseBodyOperateRuleResultOperateRuleItemList) *AttachConfigRuleToCompliancePackResponseBodyOperateRuleResult {
	s.OperateRuleItemList = v
	return s
}

type AttachConfigRuleToCompliancePackResponseBodyOperateRuleResultOperateRuleItemList struct {
	ConfigRuleId *string `json:"ConfigRuleId,omitempty" xml:"ConfigRuleId,omitempty"`
	ErrorCode    *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AttachConfigRuleToCompliancePackResponseBodyOperateRuleResultOperateRuleItemList) String() string {
	return tea.Prettify(s)
}

func (s AttachConfigRuleToCompliancePackResponseBodyOperateRuleResultOperateRuleItemList) GoString() string {
	return s.String()
}

func (s *AttachConfigRuleToCompliancePackResponseBodyOperateRuleResultOperateRuleItemList) SetConfigRuleId(v string) *AttachConfigRuleToCompliancePackResponseBodyOperateRuleResultOperateRuleItemList {
	s.ConfigRuleId = &v
	return s
}

func (s *AttachConfigRuleToCompliancePackResponseBodyOperateRuleResultOperateRuleItemList) SetErrorCode(v string) *AttachConfigRuleToCompliancePackResponseBodyOperateRuleResultOperateRuleItemList {
	s.ErrorCode = &v
	return s
}

func (s *AttachConfigRuleToCompliancePackResponseBodyOperateRuleResultOperateRuleItemList) SetSuccess(v bool) *AttachConfigRuleToCompliancePackResponseBodyOperateRuleResultOperateRuleItemList {
	s.Success = &v
	return s
}

type AttachConfigRuleToCompliancePackResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AttachConfigRuleToCompliancePackResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AttachConfigRuleToCompliancePackResponse) String() string {
	return tea.Prettify(s)
}

func (s AttachConfigRuleToCompliancePackResponse) GoString() string {
	return s.String()
}

func (s *AttachConfigRuleToCompliancePackResponse) SetHeaders(v map[string]*string) *AttachConfigRuleToCompliancePackResponse {
	s.Headers = v
	return s
}

func (s *AttachConfigRuleToCompliancePackResponse) SetStatusCode(v int32) *AttachConfigRuleToCompliancePackResponse {
	s.StatusCode = &v
	return s
}

func (s *AttachConfigRuleToCompliancePackResponse) SetBody(v *AttachConfigRuleToCompliancePackResponseBody) *AttachConfigRuleToCompliancePackResponse {
	s.Body = v
	return s
}

type CreateAggregateCompliancePackRequest struct {
	AggregatorId             *string                                            `json:"AggregatorId,omitempty" xml:"AggregatorId,omitempty"`
	ClientToken              *string                                            `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	CompliancePackName       *string                                            `json:"CompliancePackName,omitempty" xml:"CompliancePackName,omitempty"`
	CompliancePackTemplateId *string                                            `json:"CompliancePackTemplateId,omitempty" xml:"CompliancePackTemplateId,omitempty"`
	ConfigRules              []*CreateAggregateCompliancePackRequestConfigRules `json:"ConfigRules,omitempty" xml:"ConfigRules,omitempty" type:"Repeated"`
	Description              *string                                            `json:"Description,omitempty" xml:"Description,omitempty"`
	RiskLevel                *int32                                             `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
}

func (s CreateAggregateCompliancePackRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAggregateCompliancePackRequest) GoString() string {
	return s.String()
}

func (s *CreateAggregateCompliancePackRequest) SetAggregatorId(v string) *CreateAggregateCompliancePackRequest {
	s.AggregatorId = &v
	return s
}

func (s *CreateAggregateCompliancePackRequest) SetClientToken(v string) *CreateAggregateCompliancePackRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateAggregateCompliancePackRequest) SetCompliancePackName(v string) *CreateAggregateCompliancePackRequest {
	s.CompliancePackName = &v
	return s
}

func (s *CreateAggregateCompliancePackRequest) SetCompliancePackTemplateId(v string) *CreateAggregateCompliancePackRequest {
	s.CompliancePackTemplateId = &v
	return s
}

func (s *CreateAggregateCompliancePackRequest) SetConfigRules(v []*CreateAggregateCompliancePackRequestConfigRules) *CreateAggregateCompliancePackRequest {
	s.ConfigRules = v
	return s
}

func (s *CreateAggregateCompliancePackRequest) SetDescription(v string) *CreateAggregateCompliancePackRequest {
	s.Description = &v
	return s
}

func (s *CreateAggregateCompliancePackRequest) SetRiskLevel(v int32) *CreateAggregateCompliancePackRequest {
	s.RiskLevel = &v
	return s
}

type CreateAggregateCompliancePackRequestConfigRules struct {
	ConfigRuleId          *string                                                                `json:"ConfigRuleId,omitempty" xml:"ConfigRuleId,omitempty"`
	ConfigRuleName        *string                                                                `json:"ConfigRuleName,omitempty" xml:"ConfigRuleName,omitempty"`
	ConfigRuleParameters  []*CreateAggregateCompliancePackRequestConfigRulesConfigRuleParameters `json:"ConfigRuleParameters,omitempty" xml:"ConfigRuleParameters,omitempty" type:"Repeated"`
	Description           *string                                                                `json:"Description,omitempty" xml:"Description,omitempty"`
	ManagedRuleIdentifier *string                                                                `json:"ManagedRuleIdentifier,omitempty" xml:"ManagedRuleIdentifier,omitempty"`
	RiskLevel             *int32                                                                 `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
}

func (s CreateAggregateCompliancePackRequestConfigRules) String() string {
	return tea.Prettify(s)
}

func (s CreateAggregateCompliancePackRequestConfigRules) GoString() string {
	return s.String()
}

func (s *CreateAggregateCompliancePackRequestConfigRules) SetConfigRuleId(v string) *CreateAggregateCompliancePackRequestConfigRules {
	s.ConfigRuleId = &v
	return s
}

func (s *CreateAggregateCompliancePackRequestConfigRules) SetConfigRuleName(v string) *CreateAggregateCompliancePackRequestConfigRules {
	s.ConfigRuleName = &v
	return s
}

func (s *CreateAggregateCompliancePackRequestConfigRules) SetConfigRuleParameters(v []*CreateAggregateCompliancePackRequestConfigRulesConfigRuleParameters) *CreateAggregateCompliancePackRequestConfigRules {
	s.ConfigRuleParameters = v
	return s
}

func (s *CreateAggregateCompliancePackRequestConfigRules) SetDescription(v string) *CreateAggregateCompliancePackRequestConfigRules {
	s.Description = &v
	return s
}

func (s *CreateAggregateCompliancePackRequestConfigRules) SetManagedRuleIdentifier(v string) *CreateAggregateCompliancePackRequestConfigRules {
	s.ManagedRuleIdentifier = &v
	return s
}

func (s *CreateAggregateCompliancePackRequestConfigRules) SetRiskLevel(v int32) *CreateAggregateCompliancePackRequestConfigRules {
	s.RiskLevel = &v
	return s
}

type CreateAggregateCompliancePackRequestConfigRulesConfigRuleParameters struct {
	ParameterName  *string `json:"ParameterName,omitempty" xml:"ParameterName,omitempty"`
	ParameterValue *string `json:"ParameterValue,omitempty" xml:"ParameterValue,omitempty"`
}

func (s CreateAggregateCompliancePackRequestConfigRulesConfigRuleParameters) String() string {
	return tea.Prettify(s)
}

func (s CreateAggregateCompliancePackRequestConfigRulesConfigRuleParameters) GoString() string {
	return s.String()
}

func (s *CreateAggregateCompliancePackRequestConfigRulesConfigRuleParameters) SetParameterName(v string) *CreateAggregateCompliancePackRequestConfigRulesConfigRuleParameters {
	s.ParameterName = &v
	return s
}

func (s *CreateAggregateCompliancePackRequestConfigRulesConfigRuleParameters) SetParameterValue(v string) *CreateAggregateCompliancePackRequestConfigRulesConfigRuleParameters {
	s.ParameterValue = &v
	return s
}

type CreateAggregateCompliancePackShrinkRequest struct {
	AggregatorId             *string `json:"AggregatorId,omitempty" xml:"AggregatorId,omitempty"`
	ClientToken              *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	CompliancePackName       *string `json:"CompliancePackName,omitempty" xml:"CompliancePackName,omitempty"`
	CompliancePackTemplateId *string `json:"CompliancePackTemplateId,omitempty" xml:"CompliancePackTemplateId,omitempty"`
	ConfigRulesShrink        *string `json:"ConfigRules,omitempty" xml:"ConfigRules,omitempty"`
	Description              *string `json:"Description,omitempty" xml:"Description,omitempty"`
	RiskLevel                *int32  `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
}

func (s CreateAggregateCompliancePackShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAggregateCompliancePackShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateAggregateCompliancePackShrinkRequest) SetAggregatorId(v string) *CreateAggregateCompliancePackShrinkRequest {
	s.AggregatorId = &v
	return s
}

func (s *CreateAggregateCompliancePackShrinkRequest) SetClientToken(v string) *CreateAggregateCompliancePackShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateAggregateCompliancePackShrinkRequest) SetCompliancePackName(v string) *CreateAggregateCompliancePackShrinkRequest {
	s.CompliancePackName = &v
	return s
}

func (s *CreateAggregateCompliancePackShrinkRequest) SetCompliancePackTemplateId(v string) *CreateAggregateCompliancePackShrinkRequest {
	s.CompliancePackTemplateId = &v
	return s
}

func (s *CreateAggregateCompliancePackShrinkRequest) SetConfigRulesShrink(v string) *CreateAggregateCompliancePackShrinkRequest {
	s.ConfigRulesShrink = &v
	return s
}

func (s *CreateAggregateCompliancePackShrinkRequest) SetDescription(v string) *CreateAggregateCompliancePackShrinkRequest {
	s.Description = &v
	return s
}

func (s *CreateAggregateCompliancePackShrinkRequest) SetRiskLevel(v int32) *CreateAggregateCompliancePackShrinkRequest {
	s.RiskLevel = &v
	return s
}

type CreateAggregateCompliancePackResponseBody struct {
	CompliancePackId *string `json:"CompliancePackId,omitempty" xml:"CompliancePackId,omitempty"`
	RequestId        *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateAggregateCompliancePackResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateAggregateCompliancePackResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAggregateCompliancePackResponseBody) SetCompliancePackId(v string) *CreateAggregateCompliancePackResponseBody {
	s.CompliancePackId = &v
	return s
}

func (s *CreateAggregateCompliancePackResponseBody) SetRequestId(v string) *CreateAggregateCompliancePackResponseBody {
	s.RequestId = &v
	return s
}

type CreateAggregateCompliancePackResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateAggregateCompliancePackResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateAggregateCompliancePackResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateAggregateCompliancePackResponse) GoString() string {
	return s.String()
}

func (s *CreateAggregateCompliancePackResponse) SetHeaders(v map[string]*string) *CreateAggregateCompliancePackResponse {
	s.Headers = v
	return s
}

func (s *CreateAggregateCompliancePackResponse) SetStatusCode(v int32) *CreateAggregateCompliancePackResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAggregateCompliancePackResponse) SetBody(v *CreateAggregateCompliancePackResponseBody) *CreateAggregateCompliancePackResponse {
	s.Body = v
	return s
}

type CreateAggregateConfigDeliveryChannelRequest struct {
	AggregatorId                        *string `json:"AggregatorId,omitempty" xml:"AggregatorId,omitempty"`
	ClientToken                         *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	ConfigurationItemChangeNotification *bool   `json:"ConfigurationItemChangeNotification,omitempty" xml:"ConfigurationItemChangeNotification,omitempty"`
	ConfigurationSnapshot               *bool   `json:"ConfigurationSnapshot,omitempty" xml:"ConfigurationSnapshot,omitempty"`
	DeliveryChannelCondition            *string `json:"DeliveryChannelCondition,omitempty" xml:"DeliveryChannelCondition,omitempty"`
	DeliveryChannelName                 *string `json:"DeliveryChannelName,omitempty" xml:"DeliveryChannelName,omitempty"`
	DeliveryChannelTargetArn            *string `json:"DeliveryChannelTargetArn,omitempty" xml:"DeliveryChannelTargetArn,omitempty"`
	DeliveryChannelType                 *string `json:"DeliveryChannelType,omitempty" xml:"DeliveryChannelType,omitempty"`
	Description                         *string `json:"Description,omitempty" xml:"Description,omitempty"`
	NonCompliantNotification            *bool   `json:"NonCompliantNotification,omitempty" xml:"NonCompliantNotification,omitempty"`
	OversizedDataOSSTargetArn           *string `json:"OversizedDataOSSTargetArn,omitempty" xml:"OversizedDataOSSTargetArn,omitempty"`
}

func (s CreateAggregateConfigDeliveryChannelRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAggregateConfigDeliveryChannelRequest) GoString() string {
	return s.String()
}

func (s *CreateAggregateConfigDeliveryChannelRequest) SetAggregatorId(v string) *CreateAggregateConfigDeliveryChannelRequest {
	s.AggregatorId = &v
	return s
}

func (s *CreateAggregateConfigDeliveryChannelRequest) SetClientToken(v string) *CreateAggregateConfigDeliveryChannelRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateAggregateConfigDeliveryChannelRequest) SetConfigurationItemChangeNotification(v bool) *CreateAggregateConfigDeliveryChannelRequest {
	s.ConfigurationItemChangeNotification = &v
	return s
}

func (s *CreateAggregateConfigDeliveryChannelRequest) SetConfigurationSnapshot(v bool) *CreateAggregateConfigDeliveryChannelRequest {
	s.ConfigurationSnapshot = &v
	return s
}

func (s *CreateAggregateConfigDeliveryChannelRequest) SetDeliveryChannelCondition(v string) *CreateAggregateConfigDeliveryChannelRequest {
	s.DeliveryChannelCondition = &v
	return s
}

func (s *CreateAggregateConfigDeliveryChannelRequest) SetDeliveryChannelName(v string) *CreateAggregateConfigDeliveryChannelRequest {
	s.DeliveryChannelName = &v
	return s
}

func (s *CreateAggregateConfigDeliveryChannelRequest) SetDeliveryChannelTargetArn(v string) *CreateAggregateConfigDeliveryChannelRequest {
	s.DeliveryChannelTargetArn = &v
	return s
}

func (s *CreateAggregateConfigDeliveryChannelRequest) SetDeliveryChannelType(v string) *CreateAggregateConfigDeliveryChannelRequest {
	s.DeliveryChannelType = &v
	return s
}

func (s *CreateAggregateConfigDeliveryChannelRequest) SetDescription(v string) *CreateAggregateConfigDeliveryChannelRequest {
	s.Description = &v
	return s
}

func (s *CreateAggregateConfigDeliveryChannelRequest) SetNonCompliantNotification(v bool) *CreateAggregateConfigDeliveryChannelRequest {
	s.NonCompliantNotification = &v
	return s
}

func (s *CreateAggregateConfigDeliveryChannelRequest) SetOversizedDataOSSTargetArn(v string) *CreateAggregateConfigDeliveryChannelRequest {
	s.OversizedDataOSSTargetArn = &v
	return s
}

type CreateAggregateConfigDeliveryChannelResponseBody struct {
	DeliveryChannelId *string `json:"DeliveryChannelId,omitempty" xml:"DeliveryChannelId,omitempty"`
	RequestId         *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateAggregateConfigDeliveryChannelResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateAggregateConfigDeliveryChannelResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAggregateConfigDeliveryChannelResponseBody) SetDeliveryChannelId(v string) *CreateAggregateConfigDeliveryChannelResponseBody {
	s.DeliveryChannelId = &v
	return s
}

func (s *CreateAggregateConfigDeliveryChannelResponseBody) SetRequestId(v string) *CreateAggregateConfigDeliveryChannelResponseBody {
	s.RequestId = &v
	return s
}

type CreateAggregateConfigDeliveryChannelResponse struct {
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateAggregateConfigDeliveryChannelResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateAggregateConfigDeliveryChannelResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateAggregateConfigDeliveryChannelResponse) GoString() string {
	return s.String()
}

func (s *CreateAggregateConfigDeliveryChannelResponse) SetHeaders(v map[string]*string) *CreateAggregateConfigDeliveryChannelResponse {
	s.Headers = v
	return s
}

func (s *CreateAggregateConfigDeliveryChannelResponse) SetStatusCode(v int32) *CreateAggregateConfigDeliveryChannelResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAggregateConfigDeliveryChannelResponse) SetBody(v *CreateAggregateConfigDeliveryChannelResponseBody) *CreateAggregateConfigDeliveryChannelResponse {
	s.Body = v
	return s
}

type CreateAggregateConfigRuleRequest struct {
	AggregatorId              *string                `json:"AggregatorId,omitempty" xml:"AggregatorId,omitempty"`
	ClientToken               *string                `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	ConfigRuleName            *string                `json:"ConfigRuleName,omitempty" xml:"ConfigRuleName,omitempty"`
	ConfigRuleTriggerTypes    *string                `json:"ConfigRuleTriggerTypes,omitempty" xml:"ConfigRuleTriggerTypes,omitempty"`
	Description               *string                `json:"Description,omitempty" xml:"Description,omitempty"`
	ExcludeAccountIdsScope    *string                `json:"ExcludeAccountIdsScope,omitempty" xml:"ExcludeAccountIdsScope,omitempty"`
	ExcludeFolderIdsScope     *string                `json:"ExcludeFolderIdsScope,omitempty" xml:"ExcludeFolderIdsScope,omitempty"`
	ExcludeResourceIdsScope   *string                `json:"ExcludeResourceIdsScope,omitempty" xml:"ExcludeResourceIdsScope,omitempty"`
	FolderIdsScope            *string                `json:"FolderIdsScope,omitempty" xml:"FolderIdsScope,omitempty"`
	InputParameters           map[string]interface{} `json:"InputParameters,omitempty" xml:"InputParameters,omitempty"`
	MaximumExecutionFrequency *string                `json:"MaximumExecutionFrequency,omitempty" xml:"MaximumExecutionFrequency,omitempty"`
	RegionIdsScope            *string                `json:"RegionIdsScope,omitempty" xml:"RegionIdsScope,omitempty"`
	ResourceGroupIdsScope     *string                `json:"ResourceGroupIdsScope,omitempty" xml:"ResourceGroupIdsScope,omitempty"`
	ResourceTypesScope        []*string              `json:"ResourceTypesScope,omitempty" xml:"ResourceTypesScope,omitempty" type:"Repeated"`
	RiskLevel                 *int32                 `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
	SourceIdentifier          *string                `json:"SourceIdentifier,omitempty" xml:"SourceIdentifier,omitempty"`
	SourceOwner               *string                `json:"SourceOwner,omitempty" xml:"SourceOwner,omitempty"`
	TagKeyLogicScope          *string                `json:"TagKeyLogicScope,omitempty" xml:"TagKeyLogicScope,omitempty"`
	TagKeyScope               *string                `json:"TagKeyScope,omitempty" xml:"TagKeyScope,omitempty"`
	TagValueScope             *string                `json:"TagValueScope,omitempty" xml:"TagValueScope,omitempty"`
}

func (s CreateAggregateConfigRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAggregateConfigRuleRequest) GoString() string {
	return s.String()
}

func (s *CreateAggregateConfigRuleRequest) SetAggregatorId(v string) *CreateAggregateConfigRuleRequest {
	s.AggregatorId = &v
	return s
}

func (s *CreateAggregateConfigRuleRequest) SetClientToken(v string) *CreateAggregateConfigRuleRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateAggregateConfigRuleRequest) SetConfigRuleName(v string) *CreateAggregateConfigRuleRequest {
	s.ConfigRuleName = &v
	return s
}

func (s *CreateAggregateConfigRuleRequest) SetConfigRuleTriggerTypes(v string) *CreateAggregateConfigRuleRequest {
	s.ConfigRuleTriggerTypes = &v
	return s
}

func (s *CreateAggregateConfigRuleRequest) SetDescription(v string) *CreateAggregateConfigRuleRequest {
	s.Description = &v
	return s
}

func (s *CreateAggregateConfigRuleRequest) SetExcludeAccountIdsScope(v string) *CreateAggregateConfigRuleRequest {
	s.ExcludeAccountIdsScope = &v
	return s
}

func (s *CreateAggregateConfigRuleRequest) SetExcludeFolderIdsScope(v string) *CreateAggregateConfigRuleRequest {
	s.ExcludeFolderIdsScope = &v
	return s
}

func (s *CreateAggregateConfigRuleRequest) SetExcludeResourceIdsScope(v string) *CreateAggregateConfigRuleRequest {
	s.ExcludeResourceIdsScope = &v
	return s
}

func (s *CreateAggregateConfigRuleRequest) SetFolderIdsScope(v string) *CreateAggregateConfigRuleRequest {
	s.FolderIdsScope = &v
	return s
}

func (s *CreateAggregateConfigRuleRequest) SetInputParameters(v map[string]interface{}) *CreateAggregateConfigRuleRequest {
	s.InputParameters = v
	return s
}

func (s *CreateAggregateConfigRuleRequest) SetMaximumExecutionFrequency(v string) *CreateAggregateConfigRuleRequest {
	s.MaximumExecutionFrequency = &v
	return s
}

func (s *CreateAggregateConfigRuleRequest) SetRegionIdsScope(v string) *CreateAggregateConfigRuleRequest {
	s.RegionIdsScope = &v
	return s
}

func (s *CreateAggregateConfigRuleRequest) SetResourceGroupIdsScope(v string) *CreateAggregateConfigRuleRequest {
	s.ResourceGroupIdsScope = &v
	return s
}

func (s *CreateAggregateConfigRuleRequest) SetResourceTypesScope(v []*string) *CreateAggregateConfigRuleRequest {
	s.ResourceTypesScope = v
	return s
}

func (s *CreateAggregateConfigRuleRequest) SetRiskLevel(v int32) *CreateAggregateConfigRuleRequest {
	s.RiskLevel = &v
	return s
}

func (s *CreateAggregateConfigRuleRequest) SetSourceIdentifier(v string) *CreateAggregateConfigRuleRequest {
	s.SourceIdentifier = &v
	return s
}

func (s *CreateAggregateConfigRuleRequest) SetSourceOwner(v string) *CreateAggregateConfigRuleRequest {
	s.SourceOwner = &v
	return s
}

func (s *CreateAggregateConfigRuleRequest) SetTagKeyLogicScope(v string) *CreateAggregateConfigRuleRequest {
	s.TagKeyLogicScope = &v
	return s
}

func (s *CreateAggregateConfigRuleRequest) SetTagKeyScope(v string) *CreateAggregateConfigRuleRequest {
	s.TagKeyScope = &v
	return s
}

func (s *CreateAggregateConfigRuleRequest) SetTagValueScope(v string) *CreateAggregateConfigRuleRequest {
	s.TagValueScope = &v
	return s
}

type CreateAggregateConfigRuleShrinkRequest struct {
	AggregatorId              *string `json:"AggregatorId,omitempty" xml:"AggregatorId,omitempty"`
	ClientToken               *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	ConfigRuleName            *string `json:"ConfigRuleName,omitempty" xml:"ConfigRuleName,omitempty"`
	ConfigRuleTriggerTypes    *string `json:"ConfigRuleTriggerTypes,omitempty" xml:"ConfigRuleTriggerTypes,omitempty"`
	Description               *string `json:"Description,omitempty" xml:"Description,omitempty"`
	ExcludeAccountIdsScope    *string `json:"ExcludeAccountIdsScope,omitempty" xml:"ExcludeAccountIdsScope,omitempty"`
	ExcludeFolderIdsScope     *string `json:"ExcludeFolderIdsScope,omitempty" xml:"ExcludeFolderIdsScope,omitempty"`
	ExcludeResourceIdsScope   *string `json:"ExcludeResourceIdsScope,omitempty" xml:"ExcludeResourceIdsScope,omitempty"`
	FolderIdsScope            *string `json:"FolderIdsScope,omitempty" xml:"FolderIdsScope,omitempty"`
	InputParametersShrink     *string `json:"InputParameters,omitempty" xml:"InputParameters,omitempty"`
	MaximumExecutionFrequency *string `json:"MaximumExecutionFrequency,omitempty" xml:"MaximumExecutionFrequency,omitempty"`
	RegionIdsScope            *string `json:"RegionIdsScope,omitempty" xml:"RegionIdsScope,omitempty"`
	ResourceGroupIdsScope     *string `json:"ResourceGroupIdsScope,omitempty" xml:"ResourceGroupIdsScope,omitempty"`
	ResourceTypesScopeShrink  *string `json:"ResourceTypesScope,omitempty" xml:"ResourceTypesScope,omitempty"`
	RiskLevel                 *int32  `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
	SourceIdentifier          *string `json:"SourceIdentifier,omitempty" xml:"SourceIdentifier,omitempty"`
	SourceOwner               *string `json:"SourceOwner,omitempty" xml:"SourceOwner,omitempty"`
	TagKeyLogicScope          *string `json:"TagKeyLogicScope,omitempty" xml:"TagKeyLogicScope,omitempty"`
	TagKeyScope               *string `json:"TagKeyScope,omitempty" xml:"TagKeyScope,omitempty"`
	TagValueScope             *string `json:"TagValueScope,omitempty" xml:"TagValueScope,omitempty"`
}

func (s CreateAggregateConfigRuleShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAggregateConfigRuleShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateAggregateConfigRuleShrinkRequest) SetAggregatorId(v string) *CreateAggregateConfigRuleShrinkRequest {
	s.AggregatorId = &v
	return s
}

func (s *CreateAggregateConfigRuleShrinkRequest) SetClientToken(v string) *CreateAggregateConfigRuleShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateAggregateConfigRuleShrinkRequest) SetConfigRuleName(v string) *CreateAggregateConfigRuleShrinkRequest {
	s.ConfigRuleName = &v
	return s
}

func (s *CreateAggregateConfigRuleShrinkRequest) SetConfigRuleTriggerTypes(v string) *CreateAggregateConfigRuleShrinkRequest {
	s.ConfigRuleTriggerTypes = &v
	return s
}

func (s *CreateAggregateConfigRuleShrinkRequest) SetDescription(v string) *CreateAggregateConfigRuleShrinkRequest {
	s.Description = &v
	return s
}

func (s *CreateAggregateConfigRuleShrinkRequest) SetExcludeAccountIdsScope(v string) *CreateAggregateConfigRuleShrinkRequest {
	s.ExcludeAccountIdsScope = &v
	return s
}

func (s *CreateAggregateConfigRuleShrinkRequest) SetExcludeFolderIdsScope(v string) *CreateAggregateConfigRuleShrinkRequest {
	s.ExcludeFolderIdsScope = &v
	return s
}

func (s *CreateAggregateConfigRuleShrinkRequest) SetExcludeResourceIdsScope(v string) *CreateAggregateConfigRuleShrinkRequest {
	s.ExcludeResourceIdsScope = &v
	return s
}

func (s *CreateAggregateConfigRuleShrinkRequest) SetFolderIdsScope(v string) *CreateAggregateConfigRuleShrinkRequest {
	s.FolderIdsScope = &v
	return s
}

func (s *CreateAggregateConfigRuleShrinkRequest) SetInputParametersShrink(v string) *CreateAggregateConfigRuleShrinkRequest {
	s.InputParametersShrink = &v
	return s
}

func (s *CreateAggregateConfigRuleShrinkRequest) SetMaximumExecutionFrequency(v string) *CreateAggregateConfigRuleShrinkRequest {
	s.MaximumExecutionFrequency = &v
	return s
}

func (s *CreateAggregateConfigRuleShrinkRequest) SetRegionIdsScope(v string) *CreateAggregateConfigRuleShrinkRequest {
	s.RegionIdsScope = &v
	return s
}

func (s *CreateAggregateConfigRuleShrinkRequest) SetResourceGroupIdsScope(v string) *CreateAggregateConfigRuleShrinkRequest {
	s.ResourceGroupIdsScope = &v
	return s
}

func (s *CreateAggregateConfigRuleShrinkRequest) SetResourceTypesScopeShrink(v string) *CreateAggregateConfigRuleShrinkRequest {
	s.ResourceTypesScopeShrink = &v
	return s
}

func (s *CreateAggregateConfigRuleShrinkRequest) SetRiskLevel(v int32) *CreateAggregateConfigRuleShrinkRequest {
	s.RiskLevel = &v
	return s
}

func (s *CreateAggregateConfigRuleShrinkRequest) SetSourceIdentifier(v string) *CreateAggregateConfigRuleShrinkRequest {
	s.SourceIdentifier = &v
	return s
}

func (s *CreateAggregateConfigRuleShrinkRequest) SetSourceOwner(v string) *CreateAggregateConfigRuleShrinkRequest {
	s.SourceOwner = &v
	return s
}

func (s *CreateAggregateConfigRuleShrinkRequest) SetTagKeyLogicScope(v string) *CreateAggregateConfigRuleShrinkRequest {
	s.TagKeyLogicScope = &v
	return s
}

func (s *CreateAggregateConfigRuleShrinkRequest) SetTagKeyScope(v string) *CreateAggregateConfigRuleShrinkRequest {
	s.TagKeyScope = &v
	return s
}

func (s *CreateAggregateConfigRuleShrinkRequest) SetTagValueScope(v string) *CreateAggregateConfigRuleShrinkRequest {
	s.TagValueScope = &v
	return s
}

type CreateAggregateConfigRuleResponseBody struct {
	ConfigRuleId *string `json:"ConfigRuleId,omitempty" xml:"ConfigRuleId,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateAggregateConfigRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateAggregateConfigRuleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAggregateConfigRuleResponseBody) SetConfigRuleId(v string) *CreateAggregateConfigRuleResponseBody {
	s.ConfigRuleId = &v
	return s
}

func (s *CreateAggregateConfigRuleResponseBody) SetRequestId(v string) *CreateAggregateConfigRuleResponseBody {
	s.RequestId = &v
	return s
}

type CreateAggregateConfigRuleResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateAggregateConfigRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateAggregateConfigRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateAggregateConfigRuleResponse) GoString() string {
	return s.String()
}

func (s *CreateAggregateConfigRuleResponse) SetHeaders(v map[string]*string) *CreateAggregateConfigRuleResponse {
	s.Headers = v
	return s
}

func (s *CreateAggregateConfigRuleResponse) SetStatusCode(v int32) *CreateAggregateConfigRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAggregateConfigRuleResponse) SetBody(v *CreateAggregateConfigRuleResponseBody) *CreateAggregateConfigRuleResponse {
	s.Body = v
	return s
}

type CreateAggregateRemediationRequest struct {
	AggregatorId          *string `json:"AggregatorId,omitempty" xml:"AggregatorId,omitempty"`
	ClientToken           *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	ConfigRuleId          *string `json:"ConfigRuleId,omitempty" xml:"ConfigRuleId,omitempty"`
	InvokeType            *string `json:"InvokeType,omitempty" xml:"InvokeType,omitempty"`
	Params                *string `json:"Params,omitempty" xml:"Params,omitempty"`
	RemediationTemplateId *string `json:"RemediationTemplateId,omitempty" xml:"RemediationTemplateId,omitempty"`
	RemediationType       *string `json:"RemediationType,omitempty" xml:"RemediationType,omitempty"`
	SourceType            *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
}

func (s CreateAggregateRemediationRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAggregateRemediationRequest) GoString() string {
	return s.String()
}

func (s *CreateAggregateRemediationRequest) SetAggregatorId(v string) *CreateAggregateRemediationRequest {
	s.AggregatorId = &v
	return s
}

func (s *CreateAggregateRemediationRequest) SetClientToken(v string) *CreateAggregateRemediationRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateAggregateRemediationRequest) SetConfigRuleId(v string) *CreateAggregateRemediationRequest {
	s.ConfigRuleId = &v
	return s
}

func (s *CreateAggregateRemediationRequest) SetInvokeType(v string) *CreateAggregateRemediationRequest {
	s.InvokeType = &v
	return s
}

func (s *CreateAggregateRemediationRequest) SetParams(v string) *CreateAggregateRemediationRequest {
	s.Params = &v
	return s
}

func (s *CreateAggregateRemediationRequest) SetRemediationTemplateId(v string) *CreateAggregateRemediationRequest {
	s.RemediationTemplateId = &v
	return s
}

func (s *CreateAggregateRemediationRequest) SetRemediationType(v string) *CreateAggregateRemediationRequest {
	s.RemediationType = &v
	return s
}

func (s *CreateAggregateRemediationRequest) SetSourceType(v string) *CreateAggregateRemediationRequest {
	s.SourceType = &v
	return s
}

type CreateAggregateRemediationResponseBody struct {
	RemediationId *string `json:"RemediationId,omitempty" xml:"RemediationId,omitempty"`
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateAggregateRemediationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateAggregateRemediationResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAggregateRemediationResponseBody) SetRemediationId(v string) *CreateAggregateRemediationResponseBody {
	s.RemediationId = &v
	return s
}

func (s *CreateAggregateRemediationResponseBody) SetRequestId(v string) *CreateAggregateRemediationResponseBody {
	s.RequestId = &v
	return s
}

type CreateAggregateRemediationResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateAggregateRemediationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateAggregateRemediationResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateAggregateRemediationResponse) GoString() string {
	return s.String()
}

func (s *CreateAggregateRemediationResponse) SetHeaders(v map[string]*string) *CreateAggregateRemediationResponse {
	s.Headers = v
	return s
}

func (s *CreateAggregateRemediationResponse) SetStatusCode(v int32) *CreateAggregateRemediationResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAggregateRemediationResponse) SetBody(v *CreateAggregateRemediationResponseBody) *CreateAggregateRemediationResponse {
	s.Body = v
	return s
}

type CreateAggregatorRequest struct {
	AggregatorAccounts []*CreateAggregatorRequestAggregatorAccounts `json:"AggregatorAccounts,omitempty" xml:"AggregatorAccounts,omitempty" type:"Repeated"`
	AggregatorName     *string                                      `json:"AggregatorName,omitempty" xml:"AggregatorName,omitempty"`
	AggregatorType     *string                                      `json:"AggregatorType,omitempty" xml:"AggregatorType,omitempty"`
	ClientToken        *string                                      `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Description        *string                                      `json:"Description,omitempty" xml:"Description,omitempty"`
}

func (s CreateAggregatorRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAggregatorRequest) GoString() string {
	return s.String()
}

func (s *CreateAggregatorRequest) SetAggregatorAccounts(v []*CreateAggregatorRequestAggregatorAccounts) *CreateAggregatorRequest {
	s.AggregatorAccounts = v
	return s
}

func (s *CreateAggregatorRequest) SetAggregatorName(v string) *CreateAggregatorRequest {
	s.AggregatorName = &v
	return s
}

func (s *CreateAggregatorRequest) SetAggregatorType(v string) *CreateAggregatorRequest {
	s.AggregatorType = &v
	return s
}

func (s *CreateAggregatorRequest) SetClientToken(v string) *CreateAggregatorRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateAggregatorRequest) SetDescription(v string) *CreateAggregatorRequest {
	s.Description = &v
	return s
}

type CreateAggregatorRequestAggregatorAccounts struct {
	AccountId   *int64  `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	AccountType *string `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
}

func (s CreateAggregatorRequestAggregatorAccounts) String() string {
	return tea.Prettify(s)
}

func (s CreateAggregatorRequestAggregatorAccounts) GoString() string {
	return s.String()
}

func (s *CreateAggregatorRequestAggregatorAccounts) SetAccountId(v int64) *CreateAggregatorRequestAggregatorAccounts {
	s.AccountId = &v
	return s
}

func (s *CreateAggregatorRequestAggregatorAccounts) SetAccountName(v string) *CreateAggregatorRequestAggregatorAccounts {
	s.AccountName = &v
	return s
}

func (s *CreateAggregatorRequestAggregatorAccounts) SetAccountType(v string) *CreateAggregatorRequestAggregatorAccounts {
	s.AccountType = &v
	return s
}

type CreateAggregatorShrinkRequest struct {
	AggregatorAccountsShrink *string `json:"AggregatorAccounts,omitempty" xml:"AggregatorAccounts,omitempty"`
	AggregatorName           *string `json:"AggregatorName,omitempty" xml:"AggregatorName,omitempty"`
	AggregatorType           *string `json:"AggregatorType,omitempty" xml:"AggregatorType,omitempty"`
	ClientToken              *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Description              *string `json:"Description,omitempty" xml:"Description,omitempty"`
}

func (s CreateAggregatorShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAggregatorShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateAggregatorShrinkRequest) SetAggregatorAccountsShrink(v string) *CreateAggregatorShrinkRequest {
	s.AggregatorAccountsShrink = &v
	return s
}

func (s *CreateAggregatorShrinkRequest) SetAggregatorName(v string) *CreateAggregatorShrinkRequest {
	s.AggregatorName = &v
	return s
}

func (s *CreateAggregatorShrinkRequest) SetAggregatorType(v string) *CreateAggregatorShrinkRequest {
	s.AggregatorType = &v
	return s
}

func (s *CreateAggregatorShrinkRequest) SetClientToken(v string) *CreateAggregatorShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateAggregatorShrinkRequest) SetDescription(v string) *CreateAggregatorShrinkRequest {
	s.Description = &v
	return s
}

type CreateAggregatorResponseBody struct {
	AggregatorId *string `json:"AggregatorId,omitempty" xml:"AggregatorId,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateAggregatorResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateAggregatorResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAggregatorResponseBody) SetAggregatorId(v string) *CreateAggregatorResponseBody {
	s.AggregatorId = &v
	return s
}

func (s *CreateAggregatorResponseBody) SetRequestId(v string) *CreateAggregatorResponseBody {
	s.RequestId = &v
	return s
}

type CreateAggregatorResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateAggregatorResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateAggregatorResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateAggregatorResponse) GoString() string {
	return s.String()
}

func (s *CreateAggregatorResponse) SetHeaders(v map[string]*string) *CreateAggregatorResponse {
	s.Headers = v
	return s
}

func (s *CreateAggregatorResponse) SetStatusCode(v int32) *CreateAggregatorResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAggregatorResponse) SetBody(v *CreateAggregatorResponseBody) *CreateAggregatorResponse {
	s.Body = v
	return s
}

type CreateCompliancePackRequest struct {
	ClientToken              *string                                   `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	CompliancePackName       *string                                   `json:"CompliancePackName,omitempty" xml:"CompliancePackName,omitempty"`
	CompliancePackTemplateId *string                                   `json:"CompliancePackTemplateId,omitempty" xml:"CompliancePackTemplateId,omitempty"`
	ConfigRules              []*CreateCompliancePackRequestConfigRules `json:"ConfigRules,omitempty" xml:"ConfigRules,omitempty" type:"Repeated"`
	Description              *string                                   `json:"Description,omitempty" xml:"Description,omitempty"`
	RiskLevel                *int32                                    `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
}

func (s CreateCompliancePackRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateCompliancePackRequest) GoString() string {
	return s.String()
}

func (s *CreateCompliancePackRequest) SetClientToken(v string) *CreateCompliancePackRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateCompliancePackRequest) SetCompliancePackName(v string) *CreateCompliancePackRequest {
	s.CompliancePackName = &v
	return s
}

func (s *CreateCompliancePackRequest) SetCompliancePackTemplateId(v string) *CreateCompliancePackRequest {
	s.CompliancePackTemplateId = &v
	return s
}

func (s *CreateCompliancePackRequest) SetConfigRules(v []*CreateCompliancePackRequestConfigRules) *CreateCompliancePackRequest {
	s.ConfigRules = v
	return s
}

func (s *CreateCompliancePackRequest) SetDescription(v string) *CreateCompliancePackRequest {
	s.Description = &v
	return s
}

func (s *CreateCompliancePackRequest) SetRiskLevel(v int32) *CreateCompliancePackRequest {
	s.RiskLevel = &v
	return s
}

type CreateCompliancePackRequestConfigRules struct {
	ConfigRuleId          *string                                                       `json:"ConfigRuleId,omitempty" xml:"ConfigRuleId,omitempty"`
	ConfigRuleName        *string                                                       `json:"ConfigRuleName,omitempty" xml:"ConfigRuleName,omitempty"`
	ConfigRuleParameters  []*CreateCompliancePackRequestConfigRulesConfigRuleParameters `json:"ConfigRuleParameters,omitempty" xml:"ConfigRuleParameters,omitempty" type:"Repeated"`
	Description           *string                                                       `json:"Description,omitempty" xml:"Description,omitempty"`
	ManagedRuleIdentifier *string                                                       `json:"ManagedRuleIdentifier,omitempty" xml:"ManagedRuleIdentifier,omitempty"`
	RiskLevel             *int32                                                        `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
}

func (s CreateCompliancePackRequestConfigRules) String() string {
	return tea.Prettify(s)
}

func (s CreateCompliancePackRequestConfigRules) GoString() string {
	return s.String()
}

func (s *CreateCompliancePackRequestConfigRules) SetConfigRuleId(v string) *CreateCompliancePackRequestConfigRules {
	s.ConfigRuleId = &v
	return s
}

func (s *CreateCompliancePackRequestConfigRules) SetConfigRuleName(v string) *CreateCompliancePackRequestConfigRules {
	s.ConfigRuleName = &v
	return s
}

func (s *CreateCompliancePackRequestConfigRules) SetConfigRuleParameters(v []*CreateCompliancePackRequestConfigRulesConfigRuleParameters) *CreateCompliancePackRequestConfigRules {
	s.ConfigRuleParameters = v
	return s
}

func (s *CreateCompliancePackRequestConfigRules) SetDescription(v string) *CreateCompliancePackRequestConfigRules {
	s.Description = &v
	return s
}

func (s *CreateCompliancePackRequestConfigRules) SetManagedRuleIdentifier(v string) *CreateCompliancePackRequestConfigRules {
	s.ManagedRuleIdentifier = &v
	return s
}

func (s *CreateCompliancePackRequestConfigRules) SetRiskLevel(v int32) *CreateCompliancePackRequestConfigRules {
	s.RiskLevel = &v
	return s
}

type CreateCompliancePackRequestConfigRulesConfigRuleParameters struct {
	ParameterName  *string `json:"ParameterName,omitempty" xml:"ParameterName,omitempty"`
	ParameterValue *string `json:"ParameterValue,omitempty" xml:"ParameterValue,omitempty"`
}

func (s CreateCompliancePackRequestConfigRulesConfigRuleParameters) String() string {
	return tea.Prettify(s)
}

func (s CreateCompliancePackRequestConfigRulesConfigRuleParameters) GoString() string {
	return s.String()
}

func (s *CreateCompliancePackRequestConfigRulesConfigRuleParameters) SetParameterName(v string) *CreateCompliancePackRequestConfigRulesConfigRuleParameters {
	s.ParameterName = &v
	return s
}

func (s *CreateCompliancePackRequestConfigRulesConfigRuleParameters) SetParameterValue(v string) *CreateCompliancePackRequestConfigRulesConfigRuleParameters {
	s.ParameterValue = &v
	return s
}

type CreateCompliancePackShrinkRequest struct {
	ClientToken              *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	CompliancePackName       *string `json:"CompliancePackName,omitempty" xml:"CompliancePackName,omitempty"`
	CompliancePackTemplateId *string `json:"CompliancePackTemplateId,omitempty" xml:"CompliancePackTemplateId,omitempty"`
	ConfigRulesShrink        *string `json:"ConfigRules,omitempty" xml:"ConfigRules,omitempty"`
	Description              *string `json:"Description,omitempty" xml:"Description,omitempty"`
	RiskLevel                *int32  `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
}

func (s CreateCompliancePackShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateCompliancePackShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateCompliancePackShrinkRequest) SetClientToken(v string) *CreateCompliancePackShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateCompliancePackShrinkRequest) SetCompliancePackName(v string) *CreateCompliancePackShrinkRequest {
	s.CompliancePackName = &v
	return s
}

func (s *CreateCompliancePackShrinkRequest) SetCompliancePackTemplateId(v string) *CreateCompliancePackShrinkRequest {
	s.CompliancePackTemplateId = &v
	return s
}

func (s *CreateCompliancePackShrinkRequest) SetConfigRulesShrink(v string) *CreateCompliancePackShrinkRequest {
	s.ConfigRulesShrink = &v
	return s
}

func (s *CreateCompliancePackShrinkRequest) SetDescription(v string) *CreateCompliancePackShrinkRequest {
	s.Description = &v
	return s
}

func (s *CreateCompliancePackShrinkRequest) SetRiskLevel(v int32) *CreateCompliancePackShrinkRequest {
	s.RiskLevel = &v
	return s
}

type CreateCompliancePackResponseBody struct {
	CompliancePackId *string `json:"CompliancePackId,omitempty" xml:"CompliancePackId,omitempty"`
	RequestId        *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateCompliancePackResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateCompliancePackResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCompliancePackResponseBody) SetCompliancePackId(v string) *CreateCompliancePackResponseBody {
	s.CompliancePackId = &v
	return s
}

func (s *CreateCompliancePackResponseBody) SetRequestId(v string) *CreateCompliancePackResponseBody {
	s.RequestId = &v
	return s
}

type CreateCompliancePackResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateCompliancePackResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateCompliancePackResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateCompliancePackResponse) GoString() string {
	return s.String()
}

func (s *CreateCompliancePackResponse) SetHeaders(v map[string]*string) *CreateCompliancePackResponse {
	s.Headers = v
	return s
}

func (s *CreateCompliancePackResponse) SetStatusCode(v int32) *CreateCompliancePackResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateCompliancePackResponse) SetBody(v *CreateCompliancePackResponseBody) *CreateCompliancePackResponse {
	s.Body = v
	return s
}

type CreateConfigDeliveryChannelRequest struct {
	ClientToken                         *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	ConfigurationItemChangeNotification *bool   `json:"ConfigurationItemChangeNotification,omitempty" xml:"ConfigurationItemChangeNotification,omitempty"`
	ConfigurationSnapshot               *bool   `json:"ConfigurationSnapshot,omitempty" xml:"ConfigurationSnapshot,omitempty"`
	DeliveryChannelCondition            *string `json:"DeliveryChannelCondition,omitempty" xml:"DeliveryChannelCondition,omitempty"`
	DeliveryChannelName                 *string `json:"DeliveryChannelName,omitempty" xml:"DeliveryChannelName,omitempty"`
	DeliveryChannelTargetArn            *string `json:"DeliveryChannelTargetArn,omitempty" xml:"DeliveryChannelTargetArn,omitempty"`
	DeliveryChannelType                 *string `json:"DeliveryChannelType,omitempty" xml:"DeliveryChannelType,omitempty"`
	Description                         *string `json:"Description,omitempty" xml:"Description,omitempty"`
	NonCompliantNotification            *bool   `json:"NonCompliantNotification,omitempty" xml:"NonCompliantNotification,omitempty"`
	OversizedDataOSSTargetArn           *string `json:"OversizedDataOSSTargetArn,omitempty" xml:"OversizedDataOSSTargetArn,omitempty"`
}

func (s CreateConfigDeliveryChannelRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateConfigDeliveryChannelRequest) GoString() string {
	return s.String()
}

func (s *CreateConfigDeliveryChannelRequest) SetClientToken(v string) *CreateConfigDeliveryChannelRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateConfigDeliveryChannelRequest) SetConfigurationItemChangeNotification(v bool) *CreateConfigDeliveryChannelRequest {
	s.ConfigurationItemChangeNotification = &v
	return s
}

func (s *CreateConfigDeliveryChannelRequest) SetConfigurationSnapshot(v bool) *CreateConfigDeliveryChannelRequest {
	s.ConfigurationSnapshot = &v
	return s
}

func (s *CreateConfigDeliveryChannelRequest) SetDeliveryChannelCondition(v string) *CreateConfigDeliveryChannelRequest {
	s.DeliveryChannelCondition = &v
	return s
}

func (s *CreateConfigDeliveryChannelRequest) SetDeliveryChannelName(v string) *CreateConfigDeliveryChannelRequest {
	s.DeliveryChannelName = &v
	return s
}

func (s *CreateConfigDeliveryChannelRequest) SetDeliveryChannelTargetArn(v string) *CreateConfigDeliveryChannelRequest {
	s.DeliveryChannelTargetArn = &v
	return s
}

func (s *CreateConfigDeliveryChannelRequest) SetDeliveryChannelType(v string) *CreateConfigDeliveryChannelRequest {
	s.DeliveryChannelType = &v
	return s
}

func (s *CreateConfigDeliveryChannelRequest) SetDescription(v string) *CreateConfigDeliveryChannelRequest {
	s.Description = &v
	return s
}

func (s *CreateConfigDeliveryChannelRequest) SetNonCompliantNotification(v bool) *CreateConfigDeliveryChannelRequest {
	s.NonCompliantNotification = &v
	return s
}

func (s *CreateConfigDeliveryChannelRequest) SetOversizedDataOSSTargetArn(v string) *CreateConfigDeliveryChannelRequest {
	s.OversizedDataOSSTargetArn = &v
	return s
}

type CreateConfigDeliveryChannelResponseBody struct {
	DeliveryChannelId *string `json:"DeliveryChannelId,omitempty" xml:"DeliveryChannelId,omitempty"`
	RequestId         *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateConfigDeliveryChannelResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateConfigDeliveryChannelResponseBody) GoString() string {
	return s.String()
}

func (s *CreateConfigDeliveryChannelResponseBody) SetDeliveryChannelId(v string) *CreateConfigDeliveryChannelResponseBody {
	s.DeliveryChannelId = &v
	return s
}

func (s *CreateConfigDeliveryChannelResponseBody) SetRequestId(v string) *CreateConfigDeliveryChannelResponseBody {
	s.RequestId = &v
	return s
}

type CreateConfigDeliveryChannelResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateConfigDeliveryChannelResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateConfigDeliveryChannelResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateConfigDeliveryChannelResponse) GoString() string {
	return s.String()
}

func (s *CreateConfigDeliveryChannelResponse) SetHeaders(v map[string]*string) *CreateConfigDeliveryChannelResponse {
	s.Headers = v
	return s
}

func (s *CreateConfigDeliveryChannelResponse) SetStatusCode(v int32) *CreateConfigDeliveryChannelResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateConfigDeliveryChannelResponse) SetBody(v *CreateConfigDeliveryChannelResponseBody) *CreateConfigDeliveryChannelResponse {
	s.Body = v
	return s
}

type CreateConfigRuleRequest struct {
	ClientToken               *string                `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	ConfigRuleName            *string                `json:"ConfigRuleName,omitempty" xml:"ConfigRuleName,omitempty"`
	ConfigRuleTriggerTypes    *string                `json:"ConfigRuleTriggerTypes,omitempty" xml:"ConfigRuleTriggerTypes,omitempty"`
	Description               *string                `json:"Description,omitempty" xml:"Description,omitempty"`
	ExcludeResourceIdsScope   *string                `json:"ExcludeResourceIdsScope,omitempty" xml:"ExcludeResourceIdsScope,omitempty"`
	InputParameters           map[string]interface{} `json:"InputParameters,omitempty" xml:"InputParameters,omitempty"`
	MaximumExecutionFrequency *string                `json:"MaximumExecutionFrequency,omitempty" xml:"MaximumExecutionFrequency,omitempty"`
	RegionIdsScope            *string                `json:"RegionIdsScope,omitempty" xml:"RegionIdsScope,omitempty"`
	ResourceGroupIdsScope     *string                `json:"ResourceGroupIdsScope,omitempty" xml:"ResourceGroupIdsScope,omitempty"`
	ResourceTypesScope        []*string              `json:"ResourceTypesScope,omitempty" xml:"ResourceTypesScope,omitempty" type:"Repeated"`
	RiskLevel                 *int32                 `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
	SourceIdentifier          *string                `json:"SourceIdentifier,omitempty" xml:"SourceIdentifier,omitempty"`
	SourceOwner               *string                `json:"SourceOwner,omitempty" xml:"SourceOwner,omitempty"`
	TagKeyLogicScope          *string                `json:"TagKeyLogicScope,omitempty" xml:"TagKeyLogicScope,omitempty"`
	TagKeyScope               *string                `json:"TagKeyScope,omitempty" xml:"TagKeyScope,omitempty"`
	TagValueScope             *string                `json:"TagValueScope,omitempty" xml:"TagValueScope,omitempty"`
}

func (s CreateConfigRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateConfigRuleRequest) GoString() string {
	return s.String()
}

func (s *CreateConfigRuleRequest) SetClientToken(v string) *CreateConfigRuleRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateConfigRuleRequest) SetConfigRuleName(v string) *CreateConfigRuleRequest {
	s.ConfigRuleName = &v
	return s
}

func (s *CreateConfigRuleRequest) SetConfigRuleTriggerTypes(v string) *CreateConfigRuleRequest {
	s.ConfigRuleTriggerTypes = &v
	return s
}

func (s *CreateConfigRuleRequest) SetDescription(v string) *CreateConfigRuleRequest {
	s.Description = &v
	return s
}

func (s *CreateConfigRuleRequest) SetExcludeResourceIdsScope(v string) *CreateConfigRuleRequest {
	s.ExcludeResourceIdsScope = &v
	return s
}

func (s *CreateConfigRuleRequest) SetInputParameters(v map[string]interface{}) *CreateConfigRuleRequest {
	s.InputParameters = v
	return s
}

func (s *CreateConfigRuleRequest) SetMaximumExecutionFrequency(v string) *CreateConfigRuleRequest {
	s.MaximumExecutionFrequency = &v
	return s
}

func (s *CreateConfigRuleRequest) SetRegionIdsScope(v string) *CreateConfigRuleRequest {
	s.RegionIdsScope = &v
	return s
}

func (s *CreateConfigRuleRequest) SetResourceGroupIdsScope(v string) *CreateConfigRuleRequest {
	s.ResourceGroupIdsScope = &v
	return s
}

func (s *CreateConfigRuleRequest) SetResourceTypesScope(v []*string) *CreateConfigRuleRequest {
	s.ResourceTypesScope = v
	return s
}

func (s *CreateConfigRuleRequest) SetRiskLevel(v int32) *CreateConfigRuleRequest {
	s.RiskLevel = &v
	return s
}

func (s *CreateConfigRuleRequest) SetSourceIdentifier(v string) *CreateConfigRuleRequest {
	s.SourceIdentifier = &v
	return s
}

func (s *CreateConfigRuleRequest) SetSourceOwner(v string) *CreateConfigRuleRequest {
	s.SourceOwner = &v
	return s
}

func (s *CreateConfigRuleRequest) SetTagKeyLogicScope(v string) *CreateConfigRuleRequest {
	s.TagKeyLogicScope = &v
	return s
}

func (s *CreateConfigRuleRequest) SetTagKeyScope(v string) *CreateConfigRuleRequest {
	s.TagKeyScope = &v
	return s
}

func (s *CreateConfigRuleRequest) SetTagValueScope(v string) *CreateConfigRuleRequest {
	s.TagValueScope = &v
	return s
}

type CreateConfigRuleShrinkRequest struct {
	ClientToken               *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	ConfigRuleName            *string `json:"ConfigRuleName,omitempty" xml:"ConfigRuleName,omitempty"`
	ConfigRuleTriggerTypes    *string `json:"ConfigRuleTriggerTypes,omitempty" xml:"ConfigRuleTriggerTypes,omitempty"`
	Description               *string `json:"Description,omitempty" xml:"Description,omitempty"`
	ExcludeResourceIdsScope   *string `json:"ExcludeResourceIdsScope,omitempty" xml:"ExcludeResourceIdsScope,omitempty"`
	InputParametersShrink     *string `json:"InputParameters,omitempty" xml:"InputParameters,omitempty"`
	MaximumExecutionFrequency *string `json:"MaximumExecutionFrequency,omitempty" xml:"MaximumExecutionFrequency,omitempty"`
	RegionIdsScope            *string `json:"RegionIdsScope,omitempty" xml:"RegionIdsScope,omitempty"`
	ResourceGroupIdsScope     *string `json:"ResourceGroupIdsScope,omitempty" xml:"ResourceGroupIdsScope,omitempty"`
	ResourceTypesScopeShrink  *string `json:"ResourceTypesScope,omitempty" xml:"ResourceTypesScope,omitempty"`
	RiskLevel                 *int32  `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
	SourceIdentifier          *string `json:"SourceIdentifier,omitempty" xml:"SourceIdentifier,omitempty"`
	SourceOwner               *string `json:"SourceOwner,omitempty" xml:"SourceOwner,omitempty"`
	TagKeyLogicScope          *string `json:"TagKeyLogicScope,omitempty" xml:"TagKeyLogicScope,omitempty"`
	TagKeyScope               *string `json:"TagKeyScope,omitempty" xml:"TagKeyScope,omitempty"`
	TagValueScope             *string `json:"TagValueScope,omitempty" xml:"TagValueScope,omitempty"`
}

func (s CreateConfigRuleShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateConfigRuleShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateConfigRuleShrinkRequest) SetClientToken(v string) *CreateConfigRuleShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateConfigRuleShrinkRequest) SetConfigRuleName(v string) *CreateConfigRuleShrinkRequest {
	s.ConfigRuleName = &v
	return s
}

func (s *CreateConfigRuleShrinkRequest) SetConfigRuleTriggerTypes(v string) *CreateConfigRuleShrinkRequest {
	s.ConfigRuleTriggerTypes = &v
	return s
}

func (s *CreateConfigRuleShrinkRequest) SetDescription(v string) *CreateConfigRuleShrinkRequest {
	s.Description = &v
	return s
}

func (s *CreateConfigRuleShrinkRequest) SetExcludeResourceIdsScope(v string) *CreateConfigRuleShrinkRequest {
	s.ExcludeResourceIdsScope = &v
	return s
}

func (s *CreateConfigRuleShrinkRequest) SetInputParametersShrink(v string) *CreateConfigRuleShrinkRequest {
	s.InputParametersShrink = &v
	return s
}

func (s *CreateConfigRuleShrinkRequest) SetMaximumExecutionFrequency(v string) *CreateConfigRuleShrinkRequest {
	s.MaximumExecutionFrequency = &v
	return s
}

func (s *CreateConfigRuleShrinkRequest) SetRegionIdsScope(v string) *CreateConfigRuleShrinkRequest {
	s.RegionIdsScope = &v
	return s
}

func (s *CreateConfigRuleShrinkRequest) SetResourceGroupIdsScope(v string) *CreateConfigRuleShrinkRequest {
	s.ResourceGroupIdsScope = &v
	return s
}

func (s *CreateConfigRuleShrinkRequest) SetResourceTypesScopeShrink(v string) *CreateConfigRuleShrinkRequest {
	s.ResourceTypesScopeShrink = &v
	return s
}

func (s *CreateConfigRuleShrinkRequest) SetRiskLevel(v int32) *CreateConfigRuleShrinkRequest {
	s.RiskLevel = &v
	return s
}

func (s *CreateConfigRuleShrinkRequest) SetSourceIdentifier(v string) *CreateConfigRuleShrinkRequest {
	s.SourceIdentifier = &v
	return s
}

func (s *CreateConfigRuleShrinkRequest) SetSourceOwner(v string) *CreateConfigRuleShrinkRequest {
	s.SourceOwner = &v
	return s
}

func (s *CreateConfigRuleShrinkRequest) SetTagKeyLogicScope(v string) *CreateConfigRuleShrinkRequest {
	s.TagKeyLogicScope = &v
	return s
}

func (s *CreateConfigRuleShrinkRequest) SetTagKeyScope(v string) *CreateConfigRuleShrinkRequest {
	s.TagKeyScope = &v
	return s
}

func (s *CreateConfigRuleShrinkRequest) SetTagValueScope(v string) *CreateConfigRuleShrinkRequest {
	s.TagValueScope = &v
	return s
}

type CreateConfigRuleResponseBody struct {
	ConfigRuleId *string `json:"ConfigRuleId,omitempty" xml:"ConfigRuleId,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateConfigRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateConfigRuleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateConfigRuleResponseBody) SetConfigRuleId(v string) *CreateConfigRuleResponseBody {
	s.ConfigRuleId = &v
	return s
}

func (s *CreateConfigRuleResponseBody) SetRequestId(v string) *CreateConfigRuleResponseBody {
	s.RequestId = &v
	return s
}

type CreateConfigRuleResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateConfigRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateConfigRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateConfigRuleResponse) GoString() string {
	return s.String()
}

func (s *CreateConfigRuleResponse) SetHeaders(v map[string]*string) *CreateConfigRuleResponse {
	s.Headers = v
	return s
}

func (s *CreateConfigRuleResponse) SetStatusCode(v int32) *CreateConfigRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateConfigRuleResponse) SetBody(v *CreateConfigRuleResponseBody) *CreateConfigRuleResponse {
	s.Body = v
	return s
}

type CreateRemediationRequest struct {
	ClientToken           *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	ConfigRuleId          *string `json:"ConfigRuleId,omitempty" xml:"ConfigRuleId,omitempty"`
	InvokeType            *string `json:"InvokeType,omitempty" xml:"InvokeType,omitempty"`
	Params                *string `json:"Params,omitempty" xml:"Params,omitempty"`
	RemediationTemplateId *string `json:"RemediationTemplateId,omitempty" xml:"RemediationTemplateId,omitempty"`
	RemediationType       *string `json:"RemediationType,omitempty" xml:"RemediationType,omitempty"`
	SourceType            *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
}

func (s CreateRemediationRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateRemediationRequest) GoString() string {
	return s.String()
}

func (s *CreateRemediationRequest) SetClientToken(v string) *CreateRemediationRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateRemediationRequest) SetConfigRuleId(v string) *CreateRemediationRequest {
	s.ConfigRuleId = &v
	return s
}

func (s *CreateRemediationRequest) SetInvokeType(v string) *CreateRemediationRequest {
	s.InvokeType = &v
	return s
}

func (s *CreateRemediationRequest) SetParams(v string) *CreateRemediationRequest {
	s.Params = &v
	return s
}

func (s *CreateRemediationRequest) SetRemediationTemplateId(v string) *CreateRemediationRequest {
	s.RemediationTemplateId = &v
	return s
}

func (s *CreateRemediationRequest) SetRemediationType(v string) *CreateRemediationRequest {
	s.RemediationType = &v
	return s
}

func (s *CreateRemediationRequest) SetSourceType(v string) *CreateRemediationRequest {
	s.SourceType = &v
	return s
}

type CreateRemediationResponseBody struct {
	RemediationId *string `json:"RemediationId,omitempty" xml:"RemediationId,omitempty"`
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateRemediationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateRemediationResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRemediationResponseBody) SetRemediationId(v string) *CreateRemediationResponseBody {
	s.RemediationId = &v
	return s
}

func (s *CreateRemediationResponseBody) SetRequestId(v string) *CreateRemediationResponseBody {
	s.RequestId = &v
	return s
}

type CreateRemediationResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateRemediationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateRemediationResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateRemediationResponse) GoString() string {
	return s.String()
}

func (s *CreateRemediationResponse) SetHeaders(v map[string]*string) *CreateRemediationResponse {
	s.Headers = v
	return s
}

func (s *CreateRemediationResponse) SetStatusCode(v int32) *CreateRemediationResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateRemediationResponse) SetBody(v *CreateRemediationResponseBody) *CreateRemediationResponse {
	s.Body = v
	return s
}

type DeactiveAggregateConfigRulesRequest struct {
	AggregatorId  *string `json:"AggregatorId,omitempty" xml:"AggregatorId,omitempty"`
	ConfigRuleIds *string `json:"ConfigRuleIds,omitempty" xml:"ConfigRuleIds,omitempty"`
}

func (s DeactiveAggregateConfigRulesRequest) String() string {
	return tea.Prettify(s)
}

func (s DeactiveAggregateConfigRulesRequest) GoString() string {
	return s.String()
}

func (s *DeactiveAggregateConfigRulesRequest) SetAggregatorId(v string) *DeactiveAggregateConfigRulesRequest {
	s.AggregatorId = &v
	return s
}

func (s *DeactiveAggregateConfigRulesRequest) SetConfigRuleIds(v string) *DeactiveAggregateConfigRulesRequest {
	s.ConfigRuleIds = &v
	return s
}

type DeactiveAggregateConfigRulesResponseBody struct {
	OperateRuleResult *DeactiveAggregateConfigRulesResponseBodyOperateRuleResult `json:"OperateRuleResult,omitempty" xml:"OperateRuleResult,omitempty" type:"Struct"`
	RequestId         *string                                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeactiveAggregateConfigRulesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeactiveAggregateConfigRulesResponseBody) GoString() string {
	return s.String()
}

func (s *DeactiveAggregateConfigRulesResponseBody) SetOperateRuleResult(v *DeactiveAggregateConfigRulesResponseBodyOperateRuleResult) *DeactiveAggregateConfigRulesResponseBody {
	s.OperateRuleResult = v
	return s
}

func (s *DeactiveAggregateConfigRulesResponseBody) SetRequestId(v string) *DeactiveAggregateConfigRulesResponseBody {
	s.RequestId = &v
	return s
}

type DeactiveAggregateConfigRulesResponseBodyOperateRuleResult struct {
	OperateRuleItemList []*DeactiveAggregateConfigRulesResponseBodyOperateRuleResultOperateRuleItemList `json:"OperateRuleItemList,omitempty" xml:"OperateRuleItemList,omitempty" type:"Repeated"`
}

func (s DeactiveAggregateConfigRulesResponseBodyOperateRuleResult) String() string {
	return tea.Prettify(s)
}

func (s DeactiveAggregateConfigRulesResponseBodyOperateRuleResult) GoString() string {
	return s.String()
}

func (s *DeactiveAggregateConfigRulesResponseBodyOperateRuleResult) SetOperateRuleItemList(v []*DeactiveAggregateConfigRulesResponseBodyOperateRuleResultOperateRuleItemList) *DeactiveAggregateConfigRulesResponseBodyOperateRuleResult {
	s.OperateRuleItemList = v
	return s
}

type DeactiveAggregateConfigRulesResponseBodyOperateRuleResultOperateRuleItemList struct {
	ConfigRuleId *string `json:"ConfigRuleId,omitempty" xml:"ConfigRuleId,omitempty"`
	ErrorCode    *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeactiveAggregateConfigRulesResponseBodyOperateRuleResultOperateRuleItemList) String() string {
	return tea.Prettify(s)
}

func (s DeactiveAggregateConfigRulesResponseBodyOperateRuleResultOperateRuleItemList) GoString() string {
	return s.String()
}

func (s *DeactiveAggregateConfigRulesResponseBodyOperateRuleResultOperateRuleItemList) SetConfigRuleId(v string) *DeactiveAggregateConfigRulesResponseBodyOperateRuleResultOperateRuleItemList {
	s.ConfigRuleId = &v
	return s
}

func (s *DeactiveAggregateConfigRulesResponseBodyOperateRuleResultOperateRuleItemList) SetErrorCode(v string) *DeactiveAggregateConfigRulesResponseBodyOperateRuleResultOperateRuleItemList {
	s.ErrorCode = &v
	return s
}

func (s *DeactiveAggregateConfigRulesResponseBodyOperateRuleResultOperateRuleItemList) SetSuccess(v bool) *DeactiveAggregateConfigRulesResponseBodyOperateRuleResultOperateRuleItemList {
	s.Success = &v
	return s
}

type DeactiveAggregateConfigRulesResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeactiveAggregateConfigRulesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeactiveAggregateConfigRulesResponse) String() string {
	return tea.Prettify(s)
}

func (s DeactiveAggregateConfigRulesResponse) GoString() string {
	return s.String()
}

func (s *DeactiveAggregateConfigRulesResponse) SetHeaders(v map[string]*string) *DeactiveAggregateConfigRulesResponse {
	s.Headers = v
	return s
}

func (s *DeactiveAggregateConfigRulesResponse) SetStatusCode(v int32) *DeactiveAggregateConfigRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *DeactiveAggregateConfigRulesResponse) SetBody(v *DeactiveAggregateConfigRulesResponseBody) *DeactiveAggregateConfigRulesResponse {
	s.Body = v
	return s
}

type DeactiveConfigRulesRequest struct {
	ConfigRuleIds *string `json:"ConfigRuleIds,omitempty" xml:"ConfigRuleIds,omitempty"`
}

func (s DeactiveConfigRulesRequest) String() string {
	return tea.Prettify(s)
}

func (s DeactiveConfigRulesRequest) GoString() string {
	return s.String()
}

func (s *DeactiveConfigRulesRequest) SetConfigRuleIds(v string) *DeactiveConfigRulesRequest {
	s.ConfigRuleIds = &v
	return s
}

type DeactiveConfigRulesResponseBody struct {
	OperateRuleResult *DeactiveConfigRulesResponseBodyOperateRuleResult `json:"OperateRuleResult,omitempty" xml:"OperateRuleResult,omitempty" type:"Struct"`
	RequestId         *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeactiveConfigRulesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeactiveConfigRulesResponseBody) GoString() string {
	return s.String()
}

func (s *DeactiveConfigRulesResponseBody) SetOperateRuleResult(v *DeactiveConfigRulesResponseBodyOperateRuleResult) *DeactiveConfigRulesResponseBody {
	s.OperateRuleResult = v
	return s
}

func (s *DeactiveConfigRulesResponseBody) SetRequestId(v string) *DeactiveConfigRulesResponseBody {
	s.RequestId = &v
	return s
}

type DeactiveConfigRulesResponseBodyOperateRuleResult struct {
	OperateRuleItemList []*DeactiveConfigRulesResponseBodyOperateRuleResultOperateRuleItemList `json:"OperateRuleItemList,omitempty" xml:"OperateRuleItemList,omitempty" type:"Repeated"`
}

func (s DeactiveConfigRulesResponseBodyOperateRuleResult) String() string {
	return tea.Prettify(s)
}

func (s DeactiveConfigRulesResponseBodyOperateRuleResult) GoString() string {
	return s.String()
}

func (s *DeactiveConfigRulesResponseBodyOperateRuleResult) SetOperateRuleItemList(v []*DeactiveConfigRulesResponseBodyOperateRuleResultOperateRuleItemList) *DeactiveConfigRulesResponseBodyOperateRuleResult {
	s.OperateRuleItemList = v
	return s
}

type DeactiveConfigRulesResponseBodyOperateRuleResultOperateRuleItemList struct {
	ConfigRuleId *string `json:"ConfigRuleId,omitempty" xml:"ConfigRuleId,omitempty"`
	ErrorCode    *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeactiveConfigRulesResponseBodyOperateRuleResultOperateRuleItemList) String() string {
	return tea.Prettify(s)
}

func (s DeactiveConfigRulesResponseBodyOperateRuleResultOperateRuleItemList) GoString() string {
	return s.String()
}

func (s *DeactiveConfigRulesResponseBodyOperateRuleResultOperateRuleItemList) SetConfigRuleId(v string) *DeactiveConfigRulesResponseBodyOperateRuleResultOperateRuleItemList {
	s.ConfigRuleId = &v
	return s
}

func (s *DeactiveConfigRulesResponseBodyOperateRuleResultOperateRuleItemList) SetErrorCode(v string) *DeactiveConfigRulesResponseBodyOperateRuleResultOperateRuleItemList {
	s.ErrorCode = &v
	return s
}

func (s *DeactiveConfigRulesResponseBodyOperateRuleResultOperateRuleItemList) SetSuccess(v bool) *DeactiveConfigRulesResponseBodyOperateRuleResultOperateRuleItemList {
	s.Success = &v
	return s
}

type DeactiveConfigRulesResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeactiveConfigRulesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeactiveConfigRulesResponse) String() string {
	return tea.Prettify(s)
}

func (s DeactiveConfigRulesResponse) GoString() string {
	return s.String()
}

func (s *DeactiveConfigRulesResponse) SetHeaders(v map[string]*string) *DeactiveConfigRulesResponse {
	s.Headers = v
	return s
}

func (s *DeactiveConfigRulesResponse) SetStatusCode(v int32) *DeactiveConfigRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *DeactiveConfigRulesResponse) SetBody(v *DeactiveConfigRulesResponseBody) *DeactiveConfigRulesResponse {
	s.Body = v
	return s
}

type DeleteAggregateCompliancePacksRequest struct {
	AggregatorId      *string `json:"AggregatorId,omitempty" xml:"AggregatorId,omitempty"`
	ClientToken       *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	CompliancePackIds *string `json:"CompliancePackIds,omitempty" xml:"CompliancePackIds,omitempty"`
	DeleteRule        *bool   `json:"DeleteRule,omitempty" xml:"DeleteRule,omitempty"`
}

func (s DeleteAggregateCompliancePacksRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteAggregateCompliancePacksRequest) GoString() string {
	return s.String()
}

func (s *DeleteAggregateCompliancePacksRequest) SetAggregatorId(v string) *DeleteAggregateCompliancePacksRequest {
	s.AggregatorId = &v
	return s
}

func (s *DeleteAggregateCompliancePacksRequest) SetClientToken(v string) *DeleteAggregateCompliancePacksRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteAggregateCompliancePacksRequest) SetCompliancePackIds(v string) *DeleteAggregateCompliancePacksRequest {
	s.CompliancePackIds = &v
	return s
}

func (s *DeleteAggregateCompliancePacksRequest) SetDeleteRule(v bool) *DeleteAggregateCompliancePacksRequest {
	s.DeleteRule = &v
	return s
}

type DeleteAggregateCompliancePacksResponseBody struct {
	OperateCompliancePacksResult *DeleteAggregateCompliancePacksResponseBodyOperateCompliancePacksResult `json:"OperateCompliancePacksResult,omitempty" xml:"OperateCompliancePacksResult,omitempty" type:"Struct"`
	RequestId                    *string                                                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteAggregateCompliancePacksResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteAggregateCompliancePacksResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAggregateCompliancePacksResponseBody) SetOperateCompliancePacksResult(v *DeleteAggregateCompliancePacksResponseBodyOperateCompliancePacksResult) *DeleteAggregateCompliancePacksResponseBody {
	s.OperateCompliancePacksResult = v
	return s
}

func (s *DeleteAggregateCompliancePacksResponseBody) SetRequestId(v string) *DeleteAggregateCompliancePacksResponseBody {
	s.RequestId = &v
	return s
}

type DeleteAggregateCompliancePacksResponseBodyOperateCompliancePacksResult struct {
	OperateCompliancePacks []*DeleteAggregateCompliancePacksResponseBodyOperateCompliancePacksResultOperateCompliancePacks `json:"OperateCompliancePacks,omitempty" xml:"OperateCompliancePacks,omitempty" type:"Repeated"`
}

func (s DeleteAggregateCompliancePacksResponseBodyOperateCompliancePacksResult) String() string {
	return tea.Prettify(s)
}

func (s DeleteAggregateCompliancePacksResponseBodyOperateCompliancePacksResult) GoString() string {
	return s.String()
}

func (s *DeleteAggregateCompliancePacksResponseBodyOperateCompliancePacksResult) SetOperateCompliancePacks(v []*DeleteAggregateCompliancePacksResponseBodyOperateCompliancePacksResultOperateCompliancePacks) *DeleteAggregateCompliancePacksResponseBodyOperateCompliancePacksResult {
	s.OperateCompliancePacks = v
	return s
}

type DeleteAggregateCompliancePacksResponseBodyOperateCompliancePacksResultOperateCompliancePacks struct {
	CompliancePackId *string `json:"CompliancePackId,omitempty" xml:"CompliancePackId,omitempty"`
	ErrorCode        *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Success          *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteAggregateCompliancePacksResponseBodyOperateCompliancePacksResultOperateCompliancePacks) String() string {
	return tea.Prettify(s)
}

func (s DeleteAggregateCompliancePacksResponseBodyOperateCompliancePacksResultOperateCompliancePacks) GoString() string {
	return s.String()
}

func (s *DeleteAggregateCompliancePacksResponseBodyOperateCompliancePacksResultOperateCompliancePacks) SetCompliancePackId(v string) *DeleteAggregateCompliancePacksResponseBodyOperateCompliancePacksResultOperateCompliancePacks {
	s.CompliancePackId = &v
	return s
}

func (s *DeleteAggregateCompliancePacksResponseBodyOperateCompliancePacksResultOperateCompliancePacks) SetErrorCode(v string) *DeleteAggregateCompliancePacksResponseBodyOperateCompliancePacksResultOperateCompliancePacks {
	s.ErrorCode = &v
	return s
}

func (s *DeleteAggregateCompliancePacksResponseBodyOperateCompliancePacksResultOperateCompliancePacks) SetSuccess(v bool) *DeleteAggregateCompliancePacksResponseBodyOperateCompliancePacksResultOperateCompliancePacks {
	s.Success = &v
	return s
}

type DeleteAggregateCompliancePacksResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteAggregateCompliancePacksResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteAggregateCompliancePacksResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteAggregateCompliancePacksResponse) GoString() string {
	return s.String()
}

func (s *DeleteAggregateCompliancePacksResponse) SetHeaders(v map[string]*string) *DeleteAggregateCompliancePacksResponse {
	s.Headers = v
	return s
}

func (s *DeleteAggregateCompliancePacksResponse) SetStatusCode(v int32) *DeleteAggregateCompliancePacksResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAggregateCompliancePacksResponse) SetBody(v *DeleteAggregateCompliancePacksResponseBody) *DeleteAggregateCompliancePacksResponse {
	s.Body = v
	return s
}

type DeleteAggregateConfigRulesRequest struct {
	AggregatorId  *string `json:"AggregatorId,omitempty" xml:"AggregatorId,omitempty"`
	ConfigRuleIds *string `json:"ConfigRuleIds,omitempty" xml:"ConfigRuleIds,omitempty"`
}

func (s DeleteAggregateConfigRulesRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteAggregateConfigRulesRequest) GoString() string {
	return s.String()
}

func (s *DeleteAggregateConfigRulesRequest) SetAggregatorId(v string) *DeleteAggregateConfigRulesRequest {
	s.AggregatorId = &v
	return s
}

func (s *DeleteAggregateConfigRulesRequest) SetConfigRuleIds(v string) *DeleteAggregateConfigRulesRequest {
	s.ConfigRuleIds = &v
	return s
}

type DeleteAggregateConfigRulesResponseBody struct {
	OperateRuleResult *DeleteAggregateConfigRulesResponseBodyOperateRuleResult `json:"OperateRuleResult,omitempty" xml:"OperateRuleResult,omitempty" type:"Struct"`
	RequestId         *string                                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteAggregateConfigRulesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteAggregateConfigRulesResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAggregateConfigRulesResponseBody) SetOperateRuleResult(v *DeleteAggregateConfigRulesResponseBodyOperateRuleResult) *DeleteAggregateConfigRulesResponseBody {
	s.OperateRuleResult = v
	return s
}

func (s *DeleteAggregateConfigRulesResponseBody) SetRequestId(v string) *DeleteAggregateConfigRulesResponseBody {
	s.RequestId = &v
	return s
}

type DeleteAggregateConfigRulesResponseBodyOperateRuleResult struct {
	OperateRuleItemList []*DeleteAggregateConfigRulesResponseBodyOperateRuleResultOperateRuleItemList `json:"OperateRuleItemList,omitempty" xml:"OperateRuleItemList,omitempty" type:"Repeated"`
}

func (s DeleteAggregateConfigRulesResponseBodyOperateRuleResult) String() string {
	return tea.Prettify(s)
}

func (s DeleteAggregateConfigRulesResponseBodyOperateRuleResult) GoString() string {
	return s.String()
}

func (s *DeleteAggregateConfigRulesResponseBodyOperateRuleResult) SetOperateRuleItemList(v []*DeleteAggregateConfigRulesResponseBodyOperateRuleResultOperateRuleItemList) *DeleteAggregateConfigRulesResponseBodyOperateRuleResult {
	s.OperateRuleItemList = v
	return s
}

type DeleteAggregateConfigRulesResponseBodyOperateRuleResultOperateRuleItemList struct {
	ConfigRuleId *string `json:"ConfigRuleId,omitempty" xml:"ConfigRuleId,omitempty"`
	ErrorCode    *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteAggregateConfigRulesResponseBodyOperateRuleResultOperateRuleItemList) String() string {
	return tea.Prettify(s)
}

func (s DeleteAggregateConfigRulesResponseBodyOperateRuleResultOperateRuleItemList) GoString() string {
	return s.String()
}

func (s *DeleteAggregateConfigRulesResponseBodyOperateRuleResultOperateRuleItemList) SetConfigRuleId(v string) *DeleteAggregateConfigRulesResponseBodyOperateRuleResultOperateRuleItemList {
	s.ConfigRuleId = &v
	return s
}

func (s *DeleteAggregateConfigRulesResponseBodyOperateRuleResultOperateRuleItemList) SetErrorCode(v string) *DeleteAggregateConfigRulesResponseBodyOperateRuleResultOperateRuleItemList {
	s.ErrorCode = &v
	return s
}

func (s *DeleteAggregateConfigRulesResponseBodyOperateRuleResultOperateRuleItemList) SetSuccess(v bool) *DeleteAggregateConfigRulesResponseBodyOperateRuleResultOperateRuleItemList {
	s.Success = &v
	return s
}

type DeleteAggregateConfigRulesResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteAggregateConfigRulesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteAggregateConfigRulesResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteAggregateConfigRulesResponse) GoString() string {
	return s.String()
}

func (s *DeleteAggregateConfigRulesResponse) SetHeaders(v map[string]*string) *DeleteAggregateConfigRulesResponse {
	s.Headers = v
	return s
}

func (s *DeleteAggregateConfigRulesResponse) SetStatusCode(v int32) *DeleteAggregateConfigRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAggregateConfigRulesResponse) SetBody(v *DeleteAggregateConfigRulesResponseBody) *DeleteAggregateConfigRulesResponse {
	s.Body = v
	return s
}

type DeleteAggregateRemediationsRequest struct {
	AggregatorId   *string `json:"AggregatorId,omitempty" xml:"AggregatorId,omitempty"`
	RemediationIds *string `json:"RemediationIds,omitempty" xml:"RemediationIds,omitempty"`
}

func (s DeleteAggregateRemediationsRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteAggregateRemediationsRequest) GoString() string {
	return s.String()
}

func (s *DeleteAggregateRemediationsRequest) SetAggregatorId(v string) *DeleteAggregateRemediationsRequest {
	s.AggregatorId = &v
	return s
}

func (s *DeleteAggregateRemediationsRequest) SetRemediationIds(v string) *DeleteAggregateRemediationsRequest {
	s.RemediationIds = &v
	return s
}

type DeleteAggregateRemediationsResponseBody struct {
	RemediationDeleteResults []*DeleteAggregateRemediationsResponseBodyRemediationDeleteResults `json:"RemediationDeleteResults,omitempty" xml:"RemediationDeleteResults,omitempty" type:"Repeated"`
	RequestId                *string                                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteAggregateRemediationsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteAggregateRemediationsResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAggregateRemediationsResponseBody) SetRemediationDeleteResults(v []*DeleteAggregateRemediationsResponseBodyRemediationDeleteResults) *DeleteAggregateRemediationsResponseBody {
	s.RemediationDeleteResults = v
	return s
}

func (s *DeleteAggregateRemediationsResponseBody) SetRequestId(v string) *DeleteAggregateRemediationsResponseBody {
	s.RequestId = &v
	return s
}

type DeleteAggregateRemediationsResponseBodyRemediationDeleteResults struct {
	ErrorMessage  *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RemediationId *string `json:"RemediationId,omitempty" xml:"RemediationId,omitempty"`
	Success       *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteAggregateRemediationsResponseBodyRemediationDeleteResults) String() string {
	return tea.Prettify(s)
}

func (s DeleteAggregateRemediationsResponseBodyRemediationDeleteResults) GoString() string {
	return s.String()
}

func (s *DeleteAggregateRemediationsResponseBodyRemediationDeleteResults) SetErrorMessage(v string) *DeleteAggregateRemediationsResponseBodyRemediationDeleteResults {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteAggregateRemediationsResponseBodyRemediationDeleteResults) SetRemediationId(v string) *DeleteAggregateRemediationsResponseBodyRemediationDeleteResults {
	s.RemediationId = &v
	return s
}

func (s *DeleteAggregateRemediationsResponseBodyRemediationDeleteResults) SetSuccess(v bool) *DeleteAggregateRemediationsResponseBodyRemediationDeleteResults {
	s.Success = &v
	return s
}

type DeleteAggregateRemediationsResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteAggregateRemediationsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteAggregateRemediationsResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteAggregateRemediationsResponse) GoString() string {
	return s.String()
}

func (s *DeleteAggregateRemediationsResponse) SetHeaders(v map[string]*string) *DeleteAggregateRemediationsResponse {
	s.Headers = v
	return s
}

func (s *DeleteAggregateRemediationsResponse) SetStatusCode(v int32) *DeleteAggregateRemediationsResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAggregateRemediationsResponse) SetBody(v *DeleteAggregateRemediationsResponseBody) *DeleteAggregateRemediationsResponse {
	s.Body = v
	return s
}

type DeleteAggregatorsRequest struct {
	AggregatorIds *string `json:"AggregatorIds,omitempty" xml:"AggregatorIds,omitempty"`
	ClientToken   *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s DeleteAggregatorsRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteAggregatorsRequest) GoString() string {
	return s.String()
}

func (s *DeleteAggregatorsRequest) SetAggregatorIds(v string) *DeleteAggregatorsRequest {
	s.AggregatorIds = &v
	return s
}

func (s *DeleteAggregatorsRequest) SetClientToken(v string) *DeleteAggregatorsRequest {
	s.ClientToken = &v
	return s
}

type DeleteAggregatorsResponseBody struct {
	OperateAggregatorsResult *DeleteAggregatorsResponseBodyOperateAggregatorsResult `json:"OperateAggregatorsResult,omitempty" xml:"OperateAggregatorsResult,omitempty" type:"Struct"`
	RequestId                *string                                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteAggregatorsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteAggregatorsResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAggregatorsResponseBody) SetOperateAggregatorsResult(v *DeleteAggregatorsResponseBodyOperateAggregatorsResult) *DeleteAggregatorsResponseBody {
	s.OperateAggregatorsResult = v
	return s
}

func (s *DeleteAggregatorsResponseBody) SetRequestId(v string) *DeleteAggregatorsResponseBody {
	s.RequestId = &v
	return s
}

type DeleteAggregatorsResponseBodyOperateAggregatorsResult struct {
	OperateAggregators []*DeleteAggregatorsResponseBodyOperateAggregatorsResultOperateAggregators `json:"OperateAggregators,omitempty" xml:"OperateAggregators,omitempty" type:"Repeated"`
}

func (s DeleteAggregatorsResponseBodyOperateAggregatorsResult) String() string {
	return tea.Prettify(s)
}

func (s DeleteAggregatorsResponseBodyOperateAggregatorsResult) GoString() string {
	return s.String()
}

func (s *DeleteAggregatorsResponseBodyOperateAggregatorsResult) SetOperateAggregators(v []*DeleteAggregatorsResponseBodyOperateAggregatorsResultOperateAggregators) *DeleteAggregatorsResponseBodyOperateAggregatorsResult {
	s.OperateAggregators = v
	return s
}

type DeleteAggregatorsResponseBodyOperateAggregatorsResultOperateAggregators struct {
	AggregatorId *string `json:"AggregatorId,omitempty" xml:"AggregatorId,omitempty"`
	ErrorCode    *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteAggregatorsResponseBodyOperateAggregatorsResultOperateAggregators) String() string {
	return tea.Prettify(s)
}

func (s DeleteAggregatorsResponseBodyOperateAggregatorsResultOperateAggregators) GoString() string {
	return s.String()
}

func (s *DeleteAggregatorsResponseBodyOperateAggregatorsResultOperateAggregators) SetAggregatorId(v string) *DeleteAggregatorsResponseBodyOperateAggregatorsResultOperateAggregators {
	s.AggregatorId = &v
	return s
}

func (s *DeleteAggregatorsResponseBodyOperateAggregatorsResultOperateAggregators) SetErrorCode(v string) *DeleteAggregatorsResponseBodyOperateAggregatorsResultOperateAggregators {
	s.ErrorCode = &v
	return s
}

func (s *DeleteAggregatorsResponseBodyOperateAggregatorsResultOperateAggregators) SetSuccess(v bool) *DeleteAggregatorsResponseBodyOperateAggregatorsResultOperateAggregators {
	s.Success = &v
	return s
}

type DeleteAggregatorsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteAggregatorsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteAggregatorsResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteAggregatorsResponse) GoString() string {
	return s.String()
}

func (s *DeleteAggregatorsResponse) SetHeaders(v map[string]*string) *DeleteAggregatorsResponse {
	s.Headers = v
	return s
}

func (s *DeleteAggregatorsResponse) SetStatusCode(v int32) *DeleteAggregatorsResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAggregatorsResponse) SetBody(v *DeleteAggregatorsResponseBody) *DeleteAggregatorsResponse {
	s.Body = v
	return s
}

type DeleteCompliancePacksRequest struct {
	ClientToken       *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	CompliancePackIds *string `json:"CompliancePackIds,omitempty" xml:"CompliancePackIds,omitempty"`
	DeleteRule        *bool   `json:"DeleteRule,omitempty" xml:"DeleteRule,omitempty"`
}

func (s DeleteCompliancePacksRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteCompliancePacksRequest) GoString() string {
	return s.String()
}

func (s *DeleteCompliancePacksRequest) SetClientToken(v string) *DeleteCompliancePacksRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteCompliancePacksRequest) SetCompliancePackIds(v string) *DeleteCompliancePacksRequest {
	s.CompliancePackIds = &v
	return s
}

func (s *DeleteCompliancePacksRequest) SetDeleteRule(v bool) *DeleteCompliancePacksRequest {
	s.DeleteRule = &v
	return s
}

type DeleteCompliancePacksResponseBody struct {
	OperateCompliancePacksResult *DeleteCompliancePacksResponseBodyOperateCompliancePacksResult `json:"OperateCompliancePacksResult,omitempty" xml:"OperateCompliancePacksResult,omitempty" type:"Struct"`
	RequestId                    *string                                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteCompliancePacksResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteCompliancePacksResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCompliancePacksResponseBody) SetOperateCompliancePacksResult(v *DeleteCompliancePacksResponseBodyOperateCompliancePacksResult) *DeleteCompliancePacksResponseBody {
	s.OperateCompliancePacksResult = v
	return s
}

func (s *DeleteCompliancePacksResponseBody) SetRequestId(v string) *DeleteCompliancePacksResponseBody {
	s.RequestId = &v
	return s
}

type DeleteCompliancePacksResponseBodyOperateCompliancePacksResult struct {
	OperateCompliancePacks []*DeleteCompliancePacksResponseBodyOperateCompliancePacksResultOperateCompliancePacks `json:"OperateCompliancePacks,omitempty" xml:"OperateCompliancePacks,omitempty" type:"Repeated"`
}

func (s DeleteCompliancePacksResponseBodyOperateCompliancePacksResult) String() string {
	return tea.Prettify(s)
}

func (s DeleteCompliancePacksResponseBodyOperateCompliancePacksResult) GoString() string {
	return s.String()
}

func (s *DeleteCompliancePacksResponseBodyOperateCompliancePacksResult) SetOperateCompliancePacks(v []*DeleteCompliancePacksResponseBodyOperateCompliancePacksResultOperateCompliancePacks) *DeleteCompliancePacksResponseBodyOperateCompliancePacksResult {
	s.OperateCompliancePacks = v
	return s
}

type DeleteCompliancePacksResponseBodyOperateCompliancePacksResultOperateCompliancePacks struct {
	CompliancePackId *string `json:"CompliancePackId,omitempty" xml:"CompliancePackId,omitempty"`
	ErrorCode        *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Success          *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteCompliancePacksResponseBodyOperateCompliancePacksResultOperateCompliancePacks) String() string {
	return tea.Prettify(s)
}

func (s DeleteCompliancePacksResponseBodyOperateCompliancePacksResultOperateCompliancePacks) GoString() string {
	return s.String()
}

func (s *DeleteCompliancePacksResponseBodyOperateCompliancePacksResultOperateCompliancePacks) SetCompliancePackId(v string) *DeleteCompliancePacksResponseBodyOperateCompliancePacksResultOperateCompliancePacks {
	s.CompliancePackId = &v
	return s
}

func (s *DeleteCompliancePacksResponseBodyOperateCompliancePacksResultOperateCompliancePacks) SetErrorCode(v string) *DeleteCompliancePacksResponseBodyOperateCompliancePacksResultOperateCompliancePacks {
	s.ErrorCode = &v
	return s
}

func (s *DeleteCompliancePacksResponseBodyOperateCompliancePacksResultOperateCompliancePacks) SetSuccess(v bool) *DeleteCompliancePacksResponseBodyOperateCompliancePacksResultOperateCompliancePacks {
	s.Success = &v
	return s
}

type DeleteCompliancePacksResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteCompliancePacksResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteCompliancePacksResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteCompliancePacksResponse) GoString() string {
	return s.String()
}

func (s *DeleteCompliancePacksResponse) SetHeaders(v map[string]*string) *DeleteCompliancePacksResponse {
	s.Headers = v
	return s
}

func (s *DeleteCompliancePacksResponse) SetStatusCode(v int32) *DeleteCompliancePacksResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteCompliancePacksResponse) SetBody(v *DeleteCompliancePacksResponseBody) *DeleteCompliancePacksResponse {
	s.Body = v
	return s
}

type DeleteRemediationsRequest struct {
	RemediationIds *string `json:"RemediationIds,omitempty" xml:"RemediationIds,omitempty"`
}

func (s DeleteRemediationsRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteRemediationsRequest) GoString() string {
	return s.String()
}

func (s *DeleteRemediationsRequest) SetRemediationIds(v string) *DeleteRemediationsRequest {
	s.RemediationIds = &v
	return s
}

type DeleteRemediationsResponseBody struct {
	RemediationDeleteResults []*DeleteRemediationsResponseBodyRemediationDeleteResults `json:"RemediationDeleteResults,omitempty" xml:"RemediationDeleteResults,omitempty" type:"Repeated"`
	RequestId                *string                                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteRemediationsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteRemediationsResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteRemediationsResponseBody) SetRemediationDeleteResults(v []*DeleteRemediationsResponseBodyRemediationDeleteResults) *DeleteRemediationsResponseBody {
	s.RemediationDeleteResults = v
	return s
}

func (s *DeleteRemediationsResponseBody) SetRequestId(v string) *DeleteRemediationsResponseBody {
	s.RequestId = &v
	return s
}

type DeleteRemediationsResponseBodyRemediationDeleteResults struct {
	ErrorMessage  *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RemediationId *string `json:"RemediationId,omitempty" xml:"RemediationId,omitempty"`
	Success       *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteRemediationsResponseBodyRemediationDeleteResults) String() string {
	return tea.Prettify(s)
}

func (s DeleteRemediationsResponseBodyRemediationDeleteResults) GoString() string {
	return s.String()
}

func (s *DeleteRemediationsResponseBodyRemediationDeleteResults) SetErrorMessage(v string) *DeleteRemediationsResponseBodyRemediationDeleteResults {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteRemediationsResponseBodyRemediationDeleteResults) SetRemediationId(v string) *DeleteRemediationsResponseBodyRemediationDeleteResults {
	s.RemediationId = &v
	return s
}

func (s *DeleteRemediationsResponseBodyRemediationDeleteResults) SetSuccess(v bool) *DeleteRemediationsResponseBodyRemediationDeleteResults {
	s.Success = &v
	return s
}

type DeleteRemediationsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteRemediationsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteRemediationsResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteRemediationsResponse) GoString() string {
	return s.String()
}

func (s *DeleteRemediationsResponse) SetHeaders(v map[string]*string) *DeleteRemediationsResponse {
	s.Headers = v
	return s
}

func (s *DeleteRemediationsResponse) SetStatusCode(v int32) *DeleteRemediationsResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteRemediationsResponse) SetBody(v *DeleteRemediationsResponseBody) *DeleteRemediationsResponse {
	s.Body = v
	return s
}

type DetachAggregateConfigRuleToCompliancePackRequest struct {
	AggregatorId     *string `json:"AggregatorId,omitempty" xml:"AggregatorId,omitempty"`
	CompliancePackId *string `json:"CompliancePackId,omitempty" xml:"CompliancePackId,omitempty"`
	ConfigRuleIds    *string `json:"ConfigRuleIds,omitempty" xml:"ConfigRuleIds,omitempty"`
}

func (s DetachAggregateConfigRuleToCompliancePackRequest) String() string {
	return tea.Prettify(s)
}

func (s DetachAggregateConfigRuleToCompliancePackRequest) GoString() string {
	return s.String()
}

func (s *DetachAggregateConfigRuleToCompliancePackRequest) SetAggregatorId(v string) *DetachAggregateConfigRuleToCompliancePackRequest {
	s.AggregatorId = &v
	return s
}

func (s *DetachAggregateConfigRuleToCompliancePackRequest) SetCompliancePackId(v string) *DetachAggregateConfigRuleToCompliancePackRequest {
	s.CompliancePackId = &v
	return s
}

func (s *DetachAggregateConfigRuleToCompliancePackRequest) SetConfigRuleIds(v string) *DetachAggregateConfigRuleToCompliancePackRequest {
	s.ConfigRuleIds = &v
	return s
}

type DetachAggregateConfigRuleToCompliancePackResponseBody struct {
	OperateRuleResult *DetachAggregateConfigRuleToCompliancePackResponseBodyOperateRuleResult `json:"OperateRuleResult,omitempty" xml:"OperateRuleResult,omitempty" type:"Struct"`
	RequestId         *string                                                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DetachAggregateConfigRuleToCompliancePackResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DetachAggregateConfigRuleToCompliancePackResponseBody) GoString() string {
	return s.String()
}

func (s *DetachAggregateConfigRuleToCompliancePackResponseBody) SetOperateRuleResult(v *DetachAggregateConfigRuleToCompliancePackResponseBodyOperateRuleResult) *DetachAggregateConfigRuleToCompliancePackResponseBody {
	s.OperateRuleResult = v
	return s
}

func (s *DetachAggregateConfigRuleToCompliancePackResponseBody) SetRequestId(v string) *DetachAggregateConfigRuleToCompliancePackResponseBody {
	s.RequestId = &v
	return s
}

type DetachAggregateConfigRuleToCompliancePackResponseBodyOperateRuleResult struct {
	OperateRuleItemList []*DetachAggregateConfigRuleToCompliancePackResponseBodyOperateRuleResultOperateRuleItemList `json:"OperateRuleItemList,omitempty" xml:"OperateRuleItemList,omitempty" type:"Repeated"`
}

func (s DetachAggregateConfigRuleToCompliancePackResponseBodyOperateRuleResult) String() string {
	return tea.Prettify(s)
}

func (s DetachAggregateConfigRuleToCompliancePackResponseBodyOperateRuleResult) GoString() string {
	return s.String()
}

func (s *DetachAggregateConfigRuleToCompliancePackResponseBodyOperateRuleResult) SetOperateRuleItemList(v []*DetachAggregateConfigRuleToCompliancePackResponseBodyOperateRuleResultOperateRuleItemList) *DetachAggregateConfigRuleToCompliancePackResponseBodyOperateRuleResult {
	s.OperateRuleItemList = v
	return s
}

type DetachAggregateConfigRuleToCompliancePackResponseBodyOperateRuleResultOperateRuleItemList struct {
	ConfigRuleId *string `json:"ConfigRuleId,omitempty" xml:"ConfigRuleId,omitempty"`
	ErrorCode    *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DetachAggregateConfigRuleToCompliancePackResponseBodyOperateRuleResultOperateRuleItemList) String() string {
	return tea.Prettify(s)
}

func (s DetachAggregateConfigRuleToCompliancePackResponseBodyOperateRuleResultOperateRuleItemList) GoString() string {
	return s.String()
}

func (s *DetachAggregateConfigRuleToCompliancePackResponseBodyOperateRuleResultOperateRuleItemList) SetConfigRuleId(v string) *DetachAggregateConfigRuleToCompliancePackResponseBodyOperateRuleResultOperateRuleItemList {
	s.ConfigRuleId = &v
	return s
}

func (s *DetachAggregateConfigRuleToCompliancePackResponseBodyOperateRuleResultOperateRuleItemList) SetErrorCode(v string) *DetachAggregateConfigRuleToCompliancePackResponseBodyOperateRuleResultOperateRuleItemList {
	s.ErrorCode = &v
	return s
}

func (s *DetachAggregateConfigRuleToCompliancePackResponseBodyOperateRuleResultOperateRuleItemList) SetSuccess(v bool) *DetachAggregateConfigRuleToCompliancePackResponseBodyOperateRuleResultOperateRuleItemList {
	s.Success = &v
	return s
}

type DetachAggregateConfigRuleToCompliancePackResponse struct {
	Headers    map[string]*string                                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DetachAggregateConfigRuleToCompliancePackResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DetachAggregateConfigRuleToCompliancePackResponse) String() string {
	return tea.Prettify(s)
}

func (s DetachAggregateConfigRuleToCompliancePackResponse) GoString() string {
	return s.String()
}

func (s *DetachAggregateConfigRuleToCompliancePackResponse) SetHeaders(v map[string]*string) *DetachAggregateConfigRuleToCompliancePackResponse {
	s.Headers = v
	return s
}

func (s *DetachAggregateConfigRuleToCompliancePackResponse) SetStatusCode(v int32) *DetachAggregateConfigRuleToCompliancePackResponse {
	s.StatusCode = &v
	return s
}

func (s *DetachAggregateConfigRuleToCompliancePackResponse) SetBody(v *DetachAggregateConfigRuleToCompliancePackResponseBody) *DetachAggregateConfigRuleToCompliancePackResponse {
	s.Body = v
	return s
}

type DetachConfigRuleToCompliancePackRequest struct {
	CompliancePackId *string `json:"CompliancePackId,omitempty" xml:"CompliancePackId,omitempty"`
	ConfigRuleIds    *string `json:"ConfigRuleIds,omitempty" xml:"ConfigRuleIds,omitempty"`
}

func (s DetachConfigRuleToCompliancePackRequest) String() string {
	return tea.Prettify(s)
}

func (s DetachConfigRuleToCompliancePackRequest) GoString() string {
	return s.String()
}

func (s *DetachConfigRuleToCompliancePackRequest) SetCompliancePackId(v string) *DetachConfigRuleToCompliancePackRequest {
	s.CompliancePackId = &v
	return s
}

func (s *DetachConfigRuleToCompliancePackRequest) SetConfigRuleIds(v string) *DetachConfigRuleToCompliancePackRequest {
	s.ConfigRuleIds = &v
	return s
}

type DetachConfigRuleToCompliancePackResponseBody struct {
	OperateRuleResult *DetachConfigRuleToCompliancePackResponseBodyOperateRuleResult `json:"OperateRuleResult,omitempty" xml:"OperateRuleResult,omitempty" type:"Struct"`
	RequestId         *string                                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DetachConfigRuleToCompliancePackResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DetachConfigRuleToCompliancePackResponseBody) GoString() string {
	return s.String()
}

func (s *DetachConfigRuleToCompliancePackResponseBody) SetOperateRuleResult(v *DetachConfigRuleToCompliancePackResponseBodyOperateRuleResult) *DetachConfigRuleToCompliancePackResponseBody {
	s.OperateRuleResult = v
	return s
}

func (s *DetachConfigRuleToCompliancePackResponseBody) SetRequestId(v string) *DetachConfigRuleToCompliancePackResponseBody {
	s.RequestId = &v
	return s
}

type DetachConfigRuleToCompliancePackResponseBodyOperateRuleResult struct {
	OperateRuleItemList []*DetachConfigRuleToCompliancePackResponseBodyOperateRuleResultOperateRuleItemList `json:"OperateRuleItemList,omitempty" xml:"OperateRuleItemList,omitempty" type:"Repeated"`
}

func (s DetachConfigRuleToCompliancePackResponseBodyOperateRuleResult) String() string {
	return tea.Prettify(s)
}

func (s DetachConfigRuleToCompliancePackResponseBodyOperateRuleResult) GoString() string {
	return s.String()
}

func (s *DetachConfigRuleToCompliancePackResponseBodyOperateRuleResult) SetOperateRuleItemList(v []*DetachConfigRuleToCompliancePackResponseBodyOperateRuleResultOperateRuleItemList) *DetachConfigRuleToCompliancePackResponseBodyOperateRuleResult {
	s.OperateRuleItemList = v
	return s
}

type DetachConfigRuleToCompliancePackResponseBodyOperateRuleResultOperateRuleItemList struct {
	ConfigRuleId *string `json:"ConfigRuleId,omitempty" xml:"ConfigRuleId,omitempty"`
	ErrorCode    *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DetachConfigRuleToCompliancePackResponseBodyOperateRuleResultOperateRuleItemList) String() string {
	return tea.Prettify(s)
}

func (s DetachConfigRuleToCompliancePackResponseBodyOperateRuleResultOperateRuleItemList) GoString() string {
	return s.String()
}

func (s *DetachConfigRuleToCompliancePackResponseBodyOperateRuleResultOperateRuleItemList) SetConfigRuleId(v string) *DetachConfigRuleToCompliancePackResponseBodyOperateRuleResultOperateRuleItemList {
	s.ConfigRuleId = &v
	return s
}

func (s *DetachConfigRuleToCompliancePackResponseBodyOperateRuleResultOperateRuleItemList) SetErrorCode(v string) *DetachConfigRuleToCompliancePackResponseBodyOperateRuleResultOperateRuleItemList {
	s.ErrorCode = &v
	return s
}

func (s *DetachConfigRuleToCompliancePackResponseBodyOperateRuleResultOperateRuleItemList) SetSuccess(v bool) *DetachConfigRuleToCompliancePackResponseBodyOperateRuleResultOperateRuleItemList {
	s.Success = &v
	return s
}

type DetachConfigRuleToCompliancePackResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DetachConfigRuleToCompliancePackResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DetachConfigRuleToCompliancePackResponse) String() string {
	return tea.Prettify(s)
}

func (s DetachConfigRuleToCompliancePackResponse) GoString() string {
	return s.String()
}

func (s *DetachConfigRuleToCompliancePackResponse) SetHeaders(v map[string]*string) *DetachConfigRuleToCompliancePackResponse {
	s.Headers = v
	return s
}

func (s *DetachConfigRuleToCompliancePackResponse) SetStatusCode(v int32) *DetachConfigRuleToCompliancePackResponse {
	s.StatusCode = &v
	return s
}

func (s *DetachConfigRuleToCompliancePackResponse) SetBody(v *DetachConfigRuleToCompliancePackResponseBody) *DetachConfigRuleToCompliancePackResponse {
	s.Body = v
	return s
}

type GenerateAggregateCompliancePackReportRequest struct {
	AggregatorId     *string `json:"AggregatorId,omitempty" xml:"AggregatorId,omitempty"`
	ClientToken      *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	CompliancePackId *string `json:"CompliancePackId,omitempty" xml:"CompliancePackId,omitempty"`
}

func (s GenerateAggregateCompliancePackReportRequest) String() string {
	return tea.Prettify(s)
}

func (s GenerateAggregateCompliancePackReportRequest) GoString() string {
	return s.String()
}

func (s *GenerateAggregateCompliancePackReportRequest) SetAggregatorId(v string) *GenerateAggregateCompliancePackReportRequest {
	s.AggregatorId = &v
	return s
}

func (s *GenerateAggregateCompliancePackReportRequest) SetClientToken(v string) *GenerateAggregateCompliancePackReportRequest {
	s.ClientToken = &v
	return s
}

func (s *GenerateAggregateCompliancePackReportRequest) SetCompliancePackId(v string) *GenerateAggregateCompliancePackReportRequest {
	s.CompliancePackId = &v
	return s
}

type GenerateAggregateCompliancePackReportResponseBody struct {
	CompliancePackId *string `json:"CompliancePackId,omitempty" xml:"CompliancePackId,omitempty"`
	RequestId        *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GenerateAggregateCompliancePackReportResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GenerateAggregateCompliancePackReportResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateAggregateCompliancePackReportResponseBody) SetCompliancePackId(v string) *GenerateAggregateCompliancePackReportResponseBody {
	s.CompliancePackId = &v
	return s
}

func (s *GenerateAggregateCompliancePackReportResponseBody) SetRequestId(v string) *GenerateAggregateCompliancePackReportResponseBody {
	s.RequestId = &v
	return s
}

type GenerateAggregateCompliancePackReportResponse struct {
	Headers    map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GenerateAggregateCompliancePackReportResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GenerateAggregateCompliancePackReportResponse) String() string {
	return tea.Prettify(s)
}

func (s GenerateAggregateCompliancePackReportResponse) GoString() string {
	return s.String()
}

func (s *GenerateAggregateCompliancePackReportResponse) SetHeaders(v map[string]*string) *GenerateAggregateCompliancePackReportResponse {
	s.Headers = v
	return s
}

func (s *GenerateAggregateCompliancePackReportResponse) SetStatusCode(v int32) *GenerateAggregateCompliancePackReportResponse {
	s.StatusCode = &v
	return s
}

func (s *GenerateAggregateCompliancePackReportResponse) SetBody(v *GenerateAggregateCompliancePackReportResponseBody) *GenerateAggregateCompliancePackReportResponse {
	s.Body = v
	return s
}

type GenerateAggregateConfigRulesReportRequest struct {
	AggregatorId  *string `json:"AggregatorId,omitempty" xml:"AggregatorId,omitempty"`
	ClientToken   *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	ConfigRuleIds *string `json:"ConfigRuleIds,omitempty" xml:"ConfigRuleIds,omitempty"`
}

func (s GenerateAggregateConfigRulesReportRequest) String() string {
	return tea.Prettify(s)
}

func (s GenerateAggregateConfigRulesReportRequest) GoString() string {
	return s.String()
}

func (s *GenerateAggregateConfigRulesReportRequest) SetAggregatorId(v string) *GenerateAggregateConfigRulesReportRequest {
	s.AggregatorId = &v
	return s
}

func (s *GenerateAggregateConfigRulesReportRequest) SetClientToken(v string) *GenerateAggregateConfigRulesReportRequest {
	s.ClientToken = &v
	return s
}

func (s *GenerateAggregateConfigRulesReportRequest) SetConfigRuleIds(v string) *GenerateAggregateConfigRulesReportRequest {
	s.ConfigRuleIds = &v
	return s
}

type GenerateAggregateConfigRulesReportResponseBody struct {
	AggregatorId *string `json:"AggregatorId,omitempty" xml:"AggregatorId,omitempty"`
	ReportId     *string `json:"ReportId,omitempty" xml:"ReportId,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GenerateAggregateConfigRulesReportResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GenerateAggregateConfigRulesReportResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateAggregateConfigRulesReportResponseBody) SetAggregatorId(v string) *GenerateAggregateConfigRulesReportResponseBody {
	s.AggregatorId = &v
	return s
}

func (s *GenerateAggregateConfigRulesReportResponseBody) SetReportId(v string) *GenerateAggregateConfigRulesReportResponseBody {
	s.ReportId = &v
	return s
}

func (s *GenerateAggregateConfigRulesReportResponseBody) SetRequestId(v string) *GenerateAggregateConfigRulesReportResponseBody {
	s.RequestId = &v
	return s
}

type GenerateAggregateConfigRulesReportResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GenerateAggregateConfigRulesReportResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GenerateAggregateConfigRulesReportResponse) String() string {
	return tea.Prettify(s)
}

func (s GenerateAggregateConfigRulesReportResponse) GoString() string {
	return s.String()
}

func (s *GenerateAggregateConfigRulesReportResponse) SetHeaders(v map[string]*string) *GenerateAggregateConfigRulesReportResponse {
	s.Headers = v
	return s
}

func (s *GenerateAggregateConfigRulesReportResponse) SetStatusCode(v int32) *GenerateAggregateConfigRulesReportResponse {
	s.StatusCode = &v
	return s
}

func (s *GenerateAggregateConfigRulesReportResponse) SetBody(v *GenerateAggregateConfigRulesReportResponseBody) *GenerateAggregateConfigRulesReportResponse {
	s.Body = v
	return s
}

type GenerateCompliancePackReportRequest struct {
	ClientToken      *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	CompliancePackId *string `json:"CompliancePackId,omitempty" xml:"CompliancePackId,omitempty"`
}

func (s GenerateCompliancePackReportRequest) String() string {
	return tea.Prettify(s)
}

func (s GenerateCompliancePackReportRequest) GoString() string {
	return s.String()
}

func (s *GenerateCompliancePackReportRequest) SetClientToken(v string) *GenerateCompliancePackReportRequest {
	s.ClientToken = &v
	return s
}

func (s *GenerateCompliancePackReportRequest) SetCompliancePackId(v string) *GenerateCompliancePackReportRequest {
	s.CompliancePackId = &v
	return s
}

type GenerateCompliancePackReportResponseBody struct {
	CompliancePackId *string `json:"CompliancePackId,omitempty" xml:"CompliancePackId,omitempty"`
	RequestId        *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GenerateCompliancePackReportResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GenerateCompliancePackReportResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateCompliancePackReportResponseBody) SetCompliancePackId(v string) *GenerateCompliancePackReportResponseBody {
	s.CompliancePackId = &v
	return s
}

func (s *GenerateCompliancePackReportResponseBody) SetRequestId(v string) *GenerateCompliancePackReportResponseBody {
	s.RequestId = &v
	return s
}

type GenerateCompliancePackReportResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GenerateCompliancePackReportResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GenerateCompliancePackReportResponse) String() string {
	return tea.Prettify(s)
}

func (s GenerateCompliancePackReportResponse) GoString() string {
	return s.String()
}

func (s *GenerateCompliancePackReportResponse) SetHeaders(v map[string]*string) *GenerateCompliancePackReportResponse {
	s.Headers = v
	return s
}

func (s *GenerateCompliancePackReportResponse) SetStatusCode(v int32) *GenerateCompliancePackReportResponse {
	s.StatusCode = &v
	return s
}

func (s *GenerateCompliancePackReportResponse) SetBody(v *GenerateCompliancePackReportResponseBody) *GenerateCompliancePackReportResponse {
	s.Body = v
	return s
}

type GenerateConfigRulesReportRequest struct {
	ClientToken   *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	ConfigRuleIds *string `json:"ConfigRuleIds,omitempty" xml:"ConfigRuleIds,omitempty"`
}

func (s GenerateConfigRulesReportRequest) String() string {
	return tea.Prettify(s)
}

func (s GenerateConfigRulesReportRequest) GoString() string {
	return s.String()
}

func (s *GenerateConfigRulesReportRequest) SetClientToken(v string) *GenerateConfigRulesReportRequest {
	s.ClientToken = &v
	return s
}

func (s *GenerateConfigRulesReportRequest) SetConfigRuleIds(v string) *GenerateConfigRulesReportRequest {
	s.ConfigRuleIds = &v
	return s
}

type GenerateConfigRulesReportResponseBody struct {
	ReportId  *string `json:"ReportId,omitempty" xml:"ReportId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GenerateConfigRulesReportResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GenerateConfigRulesReportResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateConfigRulesReportResponseBody) SetReportId(v string) *GenerateConfigRulesReportResponseBody {
	s.ReportId = &v
	return s
}

func (s *GenerateConfigRulesReportResponseBody) SetRequestId(v string) *GenerateConfigRulesReportResponseBody {
	s.RequestId = &v
	return s
}

type GenerateConfigRulesReportResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GenerateConfigRulesReportResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GenerateConfigRulesReportResponse) String() string {
	return tea.Prettify(s)
}

func (s GenerateConfigRulesReportResponse) GoString() string {
	return s.String()
}

func (s *GenerateConfigRulesReportResponse) SetHeaders(v map[string]*string) *GenerateConfigRulesReportResponse {
	s.Headers = v
	return s
}

func (s *GenerateConfigRulesReportResponse) SetStatusCode(v int32) *GenerateConfigRulesReportResponse {
	s.StatusCode = &v
	return s
}

func (s *GenerateConfigRulesReportResponse) SetBody(v *GenerateConfigRulesReportResponseBody) *GenerateConfigRulesReportResponse {
	s.Body = v
	return s
}

type GetAggregateAccountComplianceByPackRequest struct {
	AggregatorId     *string `json:"AggregatorId,omitempty" xml:"AggregatorId,omitempty"`
	CompliancePackId *string `json:"CompliancePackId,omitempty" xml:"CompliancePackId,omitempty"`
}

func (s GetAggregateAccountComplianceByPackRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAggregateAccountComplianceByPackRequest) GoString() string {
	return s.String()
}

func (s *GetAggregateAccountComplianceByPackRequest) SetAggregatorId(v string) *GetAggregateAccountComplianceByPackRequest {
	s.AggregatorId = &v
	return s
}

func (s *GetAggregateAccountComplianceByPackRequest) SetCompliancePackId(v string) *GetAggregateAccountComplianceByPackRequest {
	s.CompliancePackId = &v
	return s
}

type GetAggregateAccountComplianceByPackResponseBody struct {
	AccountComplianceResult *GetAggregateAccountComplianceByPackResponseBodyAccountComplianceResult `json:"AccountComplianceResult,omitempty" xml:"AccountComplianceResult,omitempty" type:"Struct"`
	RequestId               *string                                                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetAggregateAccountComplianceByPackResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAggregateAccountComplianceByPackResponseBody) GoString() string {
	return s.String()
}

func (s *GetAggregateAccountComplianceByPackResponseBody) SetAccountComplianceResult(v *GetAggregateAccountComplianceByPackResponseBodyAccountComplianceResult) *GetAggregateAccountComplianceByPackResponseBody {
	s.AccountComplianceResult = v
	return s
}

func (s *GetAggregateAccountComplianceByPackResponseBody) SetRequestId(v string) *GetAggregateAccountComplianceByPackResponseBody {
	s.RequestId = &v
	return s
}

type GetAggregateAccountComplianceByPackResponseBodyAccountComplianceResult struct {
	AccountCompliances []*GetAggregateAccountComplianceByPackResponseBodyAccountComplianceResultAccountCompliances `json:"AccountCompliances,omitempty" xml:"AccountCompliances,omitempty" type:"Repeated"`
	CompliancePackId   *string                                                                                     `json:"CompliancePackId,omitempty" xml:"CompliancePackId,omitempty"`
	NonCompliantCount  *int32                                                                                      `json:"NonCompliantCount,omitempty" xml:"NonCompliantCount,omitempty"`
	TotalCount         *int32                                                                                      `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s GetAggregateAccountComplianceByPackResponseBodyAccountComplianceResult) String() string {
	return tea.Prettify(s)
}

func (s GetAggregateAccountComplianceByPackResponseBodyAccountComplianceResult) GoString() string {
	return s.String()
}

func (s *GetAggregateAccountComplianceByPackResponseBodyAccountComplianceResult) SetAccountCompliances(v []*GetAggregateAccountComplianceByPackResponseBodyAccountComplianceResultAccountCompliances) *GetAggregateAccountComplianceByPackResponseBodyAccountComplianceResult {
	s.AccountCompliances = v
	return s
}

func (s *GetAggregateAccountComplianceByPackResponseBodyAccountComplianceResult) SetCompliancePackId(v string) *GetAggregateAccountComplianceByPackResponseBodyAccountComplianceResult {
	s.CompliancePackId = &v
	return s
}

func (s *GetAggregateAccountComplianceByPackResponseBodyAccountComplianceResult) SetNonCompliantCount(v int32) *GetAggregateAccountComplianceByPackResponseBodyAccountComplianceResult {
	s.NonCompliantCount = &v
	return s
}

func (s *GetAggregateAccountComplianceByPackResponseBodyAccountComplianceResult) SetTotalCount(v int32) *GetAggregateAccountComplianceByPackResponseBodyAccountComplianceResult {
	s.TotalCount = &v
	return s
}

type GetAggregateAccountComplianceByPackResponseBodyAccountComplianceResultAccountCompliances struct {
	AccountId      *int64  `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	AccountName    *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	ComplianceType *string `json:"ComplianceType,omitempty" xml:"ComplianceType,omitempty"`
}

func (s GetAggregateAccountComplianceByPackResponseBodyAccountComplianceResultAccountCompliances) String() string {
	return tea.Prettify(s)
}

func (s GetAggregateAccountComplianceByPackResponseBodyAccountComplianceResultAccountCompliances) GoString() string {
	return s.String()
}

func (s *GetAggregateAccountComplianceByPackResponseBodyAccountComplianceResultAccountCompliances) SetAccountId(v int64) *GetAggregateAccountComplianceByPackResponseBodyAccountComplianceResultAccountCompliances {
	s.AccountId = &v
	return s
}

func (s *GetAggregateAccountComplianceByPackResponseBodyAccountComplianceResultAccountCompliances) SetAccountName(v string) *GetAggregateAccountComplianceByPackResponseBodyAccountComplianceResultAccountCompliances {
	s.AccountName = &v
	return s
}

func (s *GetAggregateAccountComplianceByPackResponseBodyAccountComplianceResultAccountCompliances) SetComplianceType(v string) *GetAggregateAccountComplianceByPackResponseBodyAccountComplianceResultAccountCompliances {
	s.ComplianceType = &v
	return s
}

type GetAggregateAccountComplianceByPackResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetAggregateAccountComplianceByPackResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetAggregateAccountComplianceByPackResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAggregateAccountComplianceByPackResponse) GoString() string {
	return s.String()
}

func (s *GetAggregateAccountComplianceByPackResponse) SetHeaders(v map[string]*string) *GetAggregateAccountComplianceByPackResponse {
	s.Headers = v
	return s
}

func (s *GetAggregateAccountComplianceByPackResponse) SetStatusCode(v int32) *GetAggregateAccountComplianceByPackResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAggregateAccountComplianceByPackResponse) SetBody(v *GetAggregateAccountComplianceByPackResponseBody) *GetAggregateAccountComplianceByPackResponse {
	s.Body = v
	return s
}

type GetAggregateCompliancePackRequest struct {
	AggregatorId     *string `json:"AggregatorId,omitempty" xml:"AggregatorId,omitempty"`
	CompliancePackId *string `json:"CompliancePackId,omitempty" xml:"CompliancePackId,omitempty"`
}

func (s GetAggregateCompliancePackRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAggregateCompliancePackRequest) GoString() string {
	return s.String()
}

func (s *GetAggregateCompliancePackRequest) SetAggregatorId(v string) *GetAggregateCompliancePackRequest {
	s.AggregatorId = &v
	return s
}

func (s *GetAggregateCompliancePackRequest) SetCompliancePackId(v string) *GetAggregateCompliancePackRequest {
	s.CompliancePackId = &v
	return s
}

type GetAggregateCompliancePackResponseBody struct {
	CompliancePack *GetAggregateCompliancePackResponseBodyCompliancePack `json:"CompliancePack,omitempty" xml:"CompliancePack,omitempty" type:"Struct"`
	RequestId      *string                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetAggregateCompliancePackResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAggregateCompliancePackResponseBody) GoString() string {
	return s.String()
}

func (s *GetAggregateCompliancePackResponseBody) SetCompliancePack(v *GetAggregateCompliancePackResponseBodyCompliancePack) *GetAggregateCompliancePackResponseBody {
	s.CompliancePack = v
	return s
}

func (s *GetAggregateCompliancePackResponseBody) SetRequestId(v string) *GetAggregateCompliancePackResponseBody {
	s.RequestId = &v
	return s
}

type GetAggregateCompliancePackResponseBodyCompliancePack struct {
	AccountId                *int64                                                             `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	AggregatorId             *string                                                            `json:"AggregatorId,omitempty" xml:"AggregatorId,omitempty"`
	CompliancePackId         *string                                                            `json:"CompliancePackId,omitempty" xml:"CompliancePackId,omitempty"`
	CompliancePackName       *string                                                            `json:"CompliancePackName,omitempty" xml:"CompliancePackName,omitempty"`
	CompliancePackTemplateId *string                                                            `json:"CompliancePackTemplateId,omitempty" xml:"CompliancePackTemplateId,omitempty"`
	ConfigRules              []*GetAggregateCompliancePackResponseBodyCompliancePackConfigRules `json:"ConfigRules,omitempty" xml:"ConfigRules,omitempty" type:"Repeated"`
	CreateTimestamp          *int64                                                             `json:"CreateTimestamp,omitempty" xml:"CreateTimestamp,omitempty"`
	Description              *string                                                            `json:"Description,omitempty" xml:"Description,omitempty"`
	RiskLevel                *int32                                                             `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
	Status                   *string                                                            `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetAggregateCompliancePackResponseBodyCompliancePack) String() string {
	return tea.Prettify(s)
}

func (s GetAggregateCompliancePackResponseBodyCompliancePack) GoString() string {
	return s.String()
}

func (s *GetAggregateCompliancePackResponseBodyCompliancePack) SetAccountId(v int64) *GetAggregateCompliancePackResponseBodyCompliancePack {
	s.AccountId = &v
	return s
}

func (s *GetAggregateCompliancePackResponseBodyCompliancePack) SetAggregatorId(v string) *GetAggregateCompliancePackResponseBodyCompliancePack {
	s.AggregatorId = &v
	return s
}

func (s *GetAggregateCompliancePackResponseBodyCompliancePack) SetCompliancePackId(v string) *GetAggregateCompliancePackResponseBodyCompliancePack {
	s.CompliancePackId = &v
	return s
}

func (s *GetAggregateCompliancePackResponseBodyCompliancePack) SetCompliancePackName(v string) *GetAggregateCompliancePackResponseBodyCompliancePack {
	s.CompliancePackName = &v
	return s
}

func (s *GetAggregateCompliancePackResponseBodyCompliancePack) SetCompliancePackTemplateId(v string) *GetAggregateCompliancePackResponseBodyCompliancePack {
	s.CompliancePackTemplateId = &v
	return s
}

func (s *GetAggregateCompliancePackResponseBodyCompliancePack) SetConfigRules(v []*GetAggregateCompliancePackResponseBodyCompliancePackConfigRules) *GetAggregateCompliancePackResponseBodyCompliancePack {
	s.ConfigRules = v
	return s
}

func (s *GetAggregateCompliancePackResponseBodyCompliancePack) SetCreateTimestamp(v int64) *GetAggregateCompliancePackResponseBodyCompliancePack {
	s.CreateTimestamp = &v
	return s
}

func (s *GetAggregateCompliancePackResponseBodyCompliancePack) SetDescription(v string) *GetAggregateCompliancePackResponseBodyCompliancePack {
	s.Description = &v
	return s
}

func (s *GetAggregateCompliancePackResponseBodyCompliancePack) SetRiskLevel(v int32) *GetAggregateCompliancePackResponseBodyCompliancePack {
	s.RiskLevel = &v
	return s
}

func (s *GetAggregateCompliancePackResponseBodyCompliancePack) SetStatus(v string) *GetAggregateCompliancePackResponseBodyCompliancePack {
	s.Status = &v
	return s
}

type GetAggregateCompliancePackResponseBodyCompliancePackConfigRules struct {
	ConfigRuleId          *string                                                                                `json:"ConfigRuleId,omitempty" xml:"ConfigRuleId,omitempty"`
	ConfigRuleName        *string                                                                                `json:"ConfigRuleName,omitempty" xml:"ConfigRuleName,omitempty"`
	ConfigRuleParameters  []*GetAggregateCompliancePackResponseBodyCompliancePackConfigRulesConfigRuleParameters `json:"ConfigRuleParameters,omitempty" xml:"ConfigRuleParameters,omitempty" type:"Repeated"`
	Description           *string                                                                                `json:"Description,omitempty" xml:"Description,omitempty"`
	ManagedRuleIdentifier *string                                                                                `json:"ManagedRuleIdentifier,omitempty" xml:"ManagedRuleIdentifier,omitempty"`
	RiskLevel             *int32                                                                                 `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
}

func (s GetAggregateCompliancePackResponseBodyCompliancePackConfigRules) String() string {
	return tea.Prettify(s)
}

func (s GetAggregateCompliancePackResponseBodyCompliancePackConfigRules) GoString() string {
	return s.String()
}

func (s *GetAggregateCompliancePackResponseBodyCompliancePackConfigRules) SetConfigRuleId(v string) *GetAggregateCompliancePackResponseBodyCompliancePackConfigRules {
	s.ConfigRuleId = &v
	return s
}

func (s *GetAggregateCompliancePackResponseBodyCompliancePackConfigRules) SetConfigRuleName(v string) *GetAggregateCompliancePackResponseBodyCompliancePackConfigRules {
	s.ConfigRuleName = &v
	return s
}

func (s *GetAggregateCompliancePackResponseBodyCompliancePackConfigRules) SetConfigRuleParameters(v []*GetAggregateCompliancePackResponseBodyCompliancePackConfigRulesConfigRuleParameters) *GetAggregateCompliancePackResponseBodyCompliancePackConfigRules {
	s.ConfigRuleParameters = v
	return s
}

func (s *GetAggregateCompliancePackResponseBodyCompliancePackConfigRules) SetDescription(v string) *GetAggregateCompliancePackResponseBodyCompliancePackConfigRules {
	s.Description = &v
	return s
}

func (s *GetAggregateCompliancePackResponseBodyCompliancePackConfigRules) SetManagedRuleIdentifier(v string) *GetAggregateCompliancePackResponseBodyCompliancePackConfigRules {
	s.ManagedRuleIdentifier = &v
	return s
}

func (s *GetAggregateCompliancePackResponseBodyCompliancePackConfigRules) SetRiskLevel(v int32) *GetAggregateCompliancePackResponseBodyCompliancePackConfigRules {
	s.RiskLevel = &v
	return s
}

type GetAggregateCompliancePackResponseBodyCompliancePackConfigRulesConfigRuleParameters struct {
	ParameterName  *string `json:"ParameterName,omitempty" xml:"ParameterName,omitempty"`
	ParameterValue *string `json:"ParameterValue,omitempty" xml:"ParameterValue,omitempty"`
	Required       *bool   `json:"Required,omitempty" xml:"Required,omitempty"`
}

func (s GetAggregateCompliancePackResponseBodyCompliancePackConfigRulesConfigRuleParameters) String() string {
	return tea.Prettify(s)
}

func (s GetAggregateCompliancePackResponseBodyCompliancePackConfigRulesConfigRuleParameters) GoString() string {
	return s.String()
}

func (s *GetAggregateCompliancePackResponseBodyCompliancePackConfigRulesConfigRuleParameters) SetParameterName(v string) *GetAggregateCompliancePackResponseBodyCompliancePackConfigRulesConfigRuleParameters {
	s.ParameterName = &v
	return s
}

func (s *GetAggregateCompliancePackResponseBodyCompliancePackConfigRulesConfigRuleParameters) SetParameterValue(v string) *GetAggregateCompliancePackResponseBodyCompliancePackConfigRulesConfigRuleParameters {
	s.ParameterValue = &v
	return s
}

func (s *GetAggregateCompliancePackResponseBodyCompliancePackConfigRulesConfigRuleParameters) SetRequired(v bool) *GetAggregateCompliancePackResponseBodyCompliancePackConfigRulesConfigRuleParameters {
	s.Required = &v
	return s
}

type GetAggregateCompliancePackResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetAggregateCompliancePackResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetAggregateCompliancePackResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAggregateCompliancePackResponse) GoString() string {
	return s.String()
}

func (s *GetAggregateCompliancePackResponse) SetHeaders(v map[string]*string) *GetAggregateCompliancePackResponse {
	s.Headers = v
	return s
}

func (s *GetAggregateCompliancePackResponse) SetStatusCode(v int32) *GetAggregateCompliancePackResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAggregateCompliancePackResponse) SetBody(v *GetAggregateCompliancePackResponseBody) *GetAggregateCompliancePackResponse {
	s.Body = v
	return s
}

type GetAggregateCompliancePackReportRequest struct {
	AggregatorId     *string `json:"AggregatorId,omitempty" xml:"AggregatorId,omitempty"`
	CompliancePackId *string `json:"CompliancePackId,omitempty" xml:"CompliancePackId,omitempty"`
}

func (s GetAggregateCompliancePackReportRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAggregateCompliancePackReportRequest) GoString() string {
	return s.String()
}

func (s *GetAggregateCompliancePackReportRequest) SetAggregatorId(v string) *GetAggregateCompliancePackReportRequest {
	s.AggregatorId = &v
	return s
}

func (s *GetAggregateCompliancePackReportRequest) SetCompliancePackId(v string) *GetAggregateCompliancePackReportRequest {
	s.CompliancePackId = &v
	return s
}

type GetAggregateCompliancePackReportResponseBody struct {
	CompliancePackReport *GetAggregateCompliancePackReportResponseBodyCompliancePackReport `json:"CompliancePackReport,omitempty" xml:"CompliancePackReport,omitempty" type:"Struct"`
	RequestId            *string                                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetAggregateCompliancePackReportResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAggregateCompliancePackReportResponseBody) GoString() string {
	return s.String()
}

func (s *GetAggregateCompliancePackReportResponseBody) SetCompliancePackReport(v *GetAggregateCompliancePackReportResponseBodyCompliancePackReport) *GetAggregateCompliancePackReportResponseBody {
	s.CompliancePackReport = v
	return s
}

func (s *GetAggregateCompliancePackReportResponseBody) SetRequestId(v string) *GetAggregateCompliancePackReportResponseBody {
	s.RequestId = &v
	return s
}

type GetAggregateCompliancePackReportResponseBodyCompliancePackReport struct {
	AccountId             *int64  `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	CompliancePackId      *string `json:"CompliancePackId,omitempty" xml:"CompliancePackId,omitempty"`
	ReportCreateTimestamp *int64  `json:"ReportCreateTimestamp,omitempty" xml:"ReportCreateTimestamp,omitempty"`
	ReportStatus          *string `json:"ReportStatus,omitempty" xml:"ReportStatus,omitempty"`
	ReportUrl             *string `json:"ReportUrl,omitempty" xml:"ReportUrl,omitempty"`
}

func (s GetAggregateCompliancePackReportResponseBodyCompliancePackReport) String() string {
	return tea.Prettify(s)
}

func (s GetAggregateCompliancePackReportResponseBodyCompliancePackReport) GoString() string {
	return s.String()
}

func (s *GetAggregateCompliancePackReportResponseBodyCompliancePackReport) SetAccountId(v int64) *GetAggregateCompliancePackReportResponseBodyCompliancePackReport {
	s.AccountId = &v
	return s
}

func (s *GetAggregateCompliancePackReportResponseBodyCompliancePackReport) SetCompliancePackId(v string) *GetAggregateCompliancePackReportResponseBodyCompliancePackReport {
	s.CompliancePackId = &v
	return s
}

func (s *GetAggregateCompliancePackReportResponseBodyCompliancePackReport) SetReportCreateTimestamp(v int64) *GetAggregateCompliancePackReportResponseBodyCompliancePackReport {
	s.ReportCreateTimestamp = &v
	return s
}

func (s *GetAggregateCompliancePackReportResponseBodyCompliancePackReport) SetReportStatus(v string) *GetAggregateCompliancePackReportResponseBodyCompliancePackReport {
	s.ReportStatus = &v
	return s
}

func (s *GetAggregateCompliancePackReportResponseBodyCompliancePackReport) SetReportUrl(v string) *GetAggregateCompliancePackReportResponseBodyCompliancePackReport {
	s.ReportUrl = &v
	return s
}

type GetAggregateCompliancePackReportResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetAggregateCompliancePackReportResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetAggregateCompliancePackReportResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAggregateCompliancePackReportResponse) GoString() string {
	return s.String()
}

func (s *GetAggregateCompliancePackReportResponse) SetHeaders(v map[string]*string) *GetAggregateCompliancePackReportResponse {
	s.Headers = v
	return s
}

func (s *GetAggregateCompliancePackReportResponse) SetStatusCode(v int32) *GetAggregateCompliancePackReportResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAggregateCompliancePackReportResponse) SetBody(v *GetAggregateCompliancePackReportResponseBody) *GetAggregateCompliancePackReportResponse {
	s.Body = v
	return s
}

type GetAggregateConfigDeliveryChannelRequest struct {
	AggregatorId      *string `json:"AggregatorId,omitempty" xml:"AggregatorId,omitempty"`
	DeliveryChannelId *string `json:"DeliveryChannelId,omitempty" xml:"DeliveryChannelId,omitempty"`
}

func (s GetAggregateConfigDeliveryChannelRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAggregateConfigDeliveryChannelRequest) GoString() string {
	return s.String()
}

func (s *GetAggregateConfigDeliveryChannelRequest) SetAggregatorId(v string) *GetAggregateConfigDeliveryChannelRequest {
	s.AggregatorId = &v
	return s
}

func (s *GetAggregateConfigDeliveryChannelRequest) SetDeliveryChannelId(v string) *GetAggregateConfigDeliveryChannelRequest {
	s.DeliveryChannelId = &v
	return s
}

type GetAggregateConfigDeliveryChannelResponseBody struct {
	DeliveryChannel *GetAggregateConfigDeliveryChannelResponseBodyDeliveryChannel `json:"DeliveryChannel,omitempty" xml:"DeliveryChannel,omitempty" type:"Struct"`
	RequestId       *string                                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetAggregateConfigDeliveryChannelResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAggregateConfigDeliveryChannelResponseBody) GoString() string {
	return s.String()
}

func (s *GetAggregateConfigDeliveryChannelResponseBody) SetDeliveryChannel(v *GetAggregateConfigDeliveryChannelResponseBodyDeliveryChannel) *GetAggregateConfigDeliveryChannelResponseBody {
	s.DeliveryChannel = v
	return s
}

func (s *GetAggregateConfigDeliveryChannelResponseBody) SetRequestId(v string) *GetAggregateConfigDeliveryChannelResponseBody {
	s.RequestId = &v
	return s
}

type GetAggregateConfigDeliveryChannelResponseBodyDeliveryChannel struct {
	AccountId                           *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	AggregatorId                        *string `json:"AggregatorId,omitempty" xml:"AggregatorId,omitempty"`
	ConfigurationItemChangeNotification *bool   `json:"ConfigurationItemChangeNotification,omitempty" xml:"ConfigurationItemChangeNotification,omitempty"`
	ConfigurationSnapshot               *bool   `json:"ConfigurationSnapshot,omitempty" xml:"ConfigurationSnapshot,omitempty"`
	DeliveryChannelAssumeRoleArn        *string `json:"DeliveryChannelAssumeRoleArn,omitempty" xml:"DeliveryChannelAssumeRoleArn,omitempty"`
	DeliveryChannelCondition            *string `json:"DeliveryChannelCondition,omitempty" xml:"DeliveryChannelCondition,omitempty"`
	DeliveryChannelId                   *string `json:"DeliveryChannelId,omitempty" xml:"DeliveryChannelId,omitempty"`
	DeliveryChannelName                 *string `json:"DeliveryChannelName,omitempty" xml:"DeliveryChannelName,omitempty"`
	DeliveryChannelTargetArn            *string `json:"DeliveryChannelTargetArn,omitempty" xml:"DeliveryChannelTargetArn,omitempty"`
	DeliveryChannelType                 *string `json:"DeliveryChannelType,omitempty" xml:"DeliveryChannelType,omitempty"`
	Description                         *string `json:"Description,omitempty" xml:"Description,omitempty"`
	NonCompliantNotification            *bool   `json:"NonCompliantNotification,omitempty" xml:"NonCompliantNotification,omitempty"`
	OversizedDataOSSTargetArn           *string `json:"OversizedDataOSSTargetArn,omitempty" xml:"OversizedDataOSSTargetArn,omitempty"`
	Status                              *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetAggregateConfigDeliveryChannelResponseBodyDeliveryChannel) String() string {
	return tea.Prettify(s)
}

func (s GetAggregateConfigDeliveryChannelResponseBodyDeliveryChannel) GoString() string {
	return s.String()
}

func (s *GetAggregateConfigDeliveryChannelResponseBodyDeliveryChannel) SetAccountId(v string) *GetAggregateConfigDeliveryChannelResponseBodyDeliveryChannel {
	s.AccountId = &v
	return s
}

func (s *GetAggregateConfigDeliveryChannelResponseBodyDeliveryChannel) SetAggregatorId(v string) *GetAggregateConfigDeliveryChannelResponseBodyDeliveryChannel {
	s.AggregatorId = &v
	return s
}

func (s *GetAggregateConfigDeliveryChannelResponseBodyDeliveryChannel) SetConfigurationItemChangeNotification(v bool) *GetAggregateConfigDeliveryChannelResponseBodyDeliveryChannel {
	s.ConfigurationItemChangeNotification = &v
	return s
}

func (s *GetAggregateConfigDeliveryChannelResponseBodyDeliveryChannel) SetConfigurationSnapshot(v bool) *GetAggregateConfigDeliveryChannelResponseBodyDeliveryChannel {
	s.ConfigurationSnapshot = &v
	return s
}

func (s *GetAggregateConfigDeliveryChannelResponseBodyDeliveryChannel) SetDeliveryChannelAssumeRoleArn(v string) *GetAggregateConfigDeliveryChannelResponseBodyDeliveryChannel {
	s.DeliveryChannelAssumeRoleArn = &v
	return s
}

func (s *GetAggregateConfigDeliveryChannelResponseBodyDeliveryChannel) SetDeliveryChannelCondition(v string) *GetAggregateConfigDeliveryChannelResponseBodyDeliveryChannel {
	s.DeliveryChannelCondition = &v
	return s
}

func (s *GetAggregateConfigDeliveryChannelResponseBodyDeliveryChannel) SetDeliveryChannelId(v string) *GetAggregateConfigDeliveryChannelResponseBodyDeliveryChannel {
	s.DeliveryChannelId = &v
	return s
}

func (s *GetAggregateConfigDeliveryChannelResponseBodyDeliveryChannel) SetDeliveryChannelName(v string) *GetAggregateConfigDeliveryChannelResponseBodyDeliveryChannel {
	s.DeliveryChannelName = &v
	return s
}

func (s *GetAggregateConfigDeliveryChannelResponseBodyDeliveryChannel) SetDeliveryChannelTargetArn(v string) *GetAggregateConfigDeliveryChannelResponseBodyDeliveryChannel {
	s.DeliveryChannelTargetArn = &v
	return s
}

func (s *GetAggregateConfigDeliveryChannelResponseBodyDeliveryChannel) SetDeliveryChannelType(v string) *GetAggregateConfigDeliveryChannelResponseBodyDeliveryChannel {
	s.DeliveryChannelType = &v
	return s
}

func (s *GetAggregateConfigDeliveryChannelResponseBodyDeliveryChannel) SetDescription(v string) *GetAggregateConfigDeliveryChannelResponseBodyDeliveryChannel {
	s.Description = &v
	return s
}

func (s *GetAggregateConfigDeliveryChannelResponseBodyDeliveryChannel) SetNonCompliantNotification(v bool) *GetAggregateConfigDeliveryChannelResponseBodyDeliveryChannel {
	s.NonCompliantNotification = &v
	return s
}

func (s *GetAggregateConfigDeliveryChannelResponseBodyDeliveryChannel) SetOversizedDataOSSTargetArn(v string) *GetAggregateConfigDeliveryChannelResponseBodyDeliveryChannel {
	s.OversizedDataOSSTargetArn = &v
	return s
}

func (s *GetAggregateConfigDeliveryChannelResponseBodyDeliveryChannel) SetStatus(v int32) *GetAggregateConfigDeliveryChannelResponseBodyDeliveryChannel {
	s.Status = &v
	return s
}

type GetAggregateConfigDeliveryChannelResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetAggregateConfigDeliveryChannelResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetAggregateConfigDeliveryChannelResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAggregateConfigDeliveryChannelResponse) GoString() string {
	return s.String()
}

func (s *GetAggregateConfigDeliveryChannelResponse) SetHeaders(v map[string]*string) *GetAggregateConfigDeliveryChannelResponse {
	s.Headers = v
	return s
}

func (s *GetAggregateConfigDeliveryChannelResponse) SetStatusCode(v int32) *GetAggregateConfigDeliveryChannelResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAggregateConfigDeliveryChannelResponse) SetBody(v *GetAggregateConfigDeliveryChannelResponseBody) *GetAggregateConfigDeliveryChannelResponse {
	s.Body = v
	return s
}

type GetAggregateConfigRuleRequest struct {
	AggregatorId *string `json:"AggregatorId,omitempty" xml:"AggregatorId,omitempty"`
	ConfigRuleId *string `json:"ConfigRuleId,omitempty" xml:"ConfigRuleId,omitempty"`
}

func (s GetAggregateConfigRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAggregateConfigRuleRequest) GoString() string {
	return s.String()
}

func (s *GetAggregateConfigRuleRequest) SetAggregatorId(v string) *GetAggregateConfigRuleRequest {
	s.AggregatorId = &v
	return s
}

func (s *GetAggregateConfigRuleRequest) SetConfigRuleId(v string) *GetAggregateConfigRuleRequest {
	s.ConfigRuleId = &v
	return s
}

type GetAggregateConfigRuleResponseBody struct {
	ConfigRule *GetAggregateConfigRuleResponseBodyConfigRule `json:"ConfigRule,omitempty" xml:"ConfigRule,omitempty" type:"Struct"`
	RequestId  *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetAggregateConfigRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAggregateConfigRuleResponseBody) GoString() string {
	return s.String()
}

func (s *GetAggregateConfigRuleResponseBody) SetConfigRule(v *GetAggregateConfigRuleResponseBodyConfigRule) *GetAggregateConfigRuleResponseBody {
	s.ConfigRule = v
	return s
}

func (s *GetAggregateConfigRuleResponseBody) SetRequestId(v string) *GetAggregateConfigRuleResponseBody {
	s.RequestId = &v
	return s
}

type GetAggregateConfigRuleResponseBodyConfigRule struct {
	AccountId                  *int64                                                                  `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	Compliance                 *GetAggregateConfigRuleResponseBodyConfigRuleCompliance                 `json:"Compliance,omitempty" xml:"Compliance,omitempty" type:"Struct"`
	ConfigRuleArn              *string                                                                 `json:"ConfigRuleArn,omitempty" xml:"ConfigRuleArn,omitempty"`
	ConfigRuleEvaluationStatus *GetAggregateConfigRuleResponseBodyConfigRuleConfigRuleEvaluationStatus `json:"ConfigRuleEvaluationStatus,omitempty" xml:"ConfigRuleEvaluationStatus,omitempty" type:"Struct"`
	ConfigRuleId               *string                                                                 `json:"ConfigRuleId,omitempty" xml:"ConfigRuleId,omitempty"`
	ConfigRuleName             *string                                                                 `json:"ConfigRuleName,omitempty" xml:"ConfigRuleName,omitempty"`
	ConfigRuleState            *string                                                                 `json:"ConfigRuleState,omitempty" xml:"ConfigRuleState,omitempty"`
	ConfigRuleTriggerTypes     *string                                                                 `json:"ConfigRuleTriggerTypes,omitempty" xml:"ConfigRuleTriggerTypes,omitempty"`
	CreateBy                   *GetAggregateConfigRuleResponseBodyConfigRuleCreateBy                   `json:"CreateBy,omitempty" xml:"CreateBy,omitempty" type:"Struct"`
	CreateTimestamp            *int64                                                                  `json:"CreateTimestamp,omitempty" xml:"CreateTimestamp,omitempty"`
	Description                *string                                                                 `json:"Description,omitempty" xml:"Description,omitempty"`
	ExcludeAccountIdsScope     *string                                                                 `json:"ExcludeAccountIdsScope,omitempty" xml:"ExcludeAccountIdsScope,omitempty"`
	ExcludeFolderIdsScope      *string                                                                 `json:"ExcludeFolderIdsScope,omitempty" xml:"ExcludeFolderIdsScope,omitempty"`
	ExcludeResourceIdsScope    *string                                                                 `json:"ExcludeResourceIdsScope,omitempty" xml:"ExcludeResourceIdsScope,omitempty"`
	FolderIdsScope             *string                                                                 `json:"FolderIdsScope,omitempty" xml:"FolderIdsScope,omitempty"`
	InputParameters            map[string]interface{}                                                  `json:"InputParameters,omitempty" xml:"InputParameters,omitempty"`
	ManagedRule                *GetAggregateConfigRuleResponseBodyConfigRuleManagedRule                `json:"ManagedRule,omitempty" xml:"ManagedRule,omitempty" type:"Struct"`
	MaximumExecutionFrequency  *string                                                                 `json:"MaximumExecutionFrequency,omitempty" xml:"MaximumExecutionFrequency,omitempty"`
	ModifiedTimestamp          *int64                                                                  `json:"ModifiedTimestamp,omitempty" xml:"ModifiedTimestamp,omitempty"`
	RegionIdsScope             *string                                                                 `json:"RegionIdsScope,omitempty" xml:"RegionIdsScope,omitempty"`
	ResourceGroupIdsScope      *string                                                                 `json:"ResourceGroupIdsScope,omitempty" xml:"ResourceGroupIdsScope,omitempty"`
	ResourceTypesScope         *string                                                                 `json:"ResourceTypesScope,omitempty" xml:"ResourceTypesScope,omitempty"`
	RiskLevel                  *int32                                                                  `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
	Source                     *GetAggregateConfigRuleResponseBodyConfigRuleSource                     `json:"Source,omitempty" xml:"Source,omitempty" type:"Struct"`
	TagKeyLogicScope           *string                                                                 `json:"TagKeyLogicScope,omitempty" xml:"TagKeyLogicScope,omitempty"`
	TagKeyScope                *string                                                                 `json:"TagKeyScope,omitempty" xml:"TagKeyScope,omitempty"`
	TagValueScope              *string                                                                 `json:"TagValueScope,omitempty" xml:"TagValueScope,omitempty"`
}

func (s GetAggregateConfigRuleResponseBodyConfigRule) String() string {
	return tea.Prettify(s)
}

func (s GetAggregateConfigRuleResponseBodyConfigRule) GoString() string {
	return s.String()
}

func (s *GetAggregateConfigRuleResponseBodyConfigRule) SetAccountId(v int64) *GetAggregateConfigRuleResponseBodyConfigRule {
	s.AccountId = &v
	return s
}

func (s *GetAggregateConfigRuleResponseBodyConfigRule) SetCompliance(v *GetAggregateConfigRuleResponseBodyConfigRuleCompliance) *GetAggregateConfigRuleResponseBodyConfigRule {
	s.Compliance = v
	return s
}

func (s *GetAggregateConfigRuleResponseBodyConfigRule) SetConfigRuleArn(v string) *GetAggregateConfigRuleResponseBodyConfigRule {
	s.ConfigRuleArn = &v
	return s
}

func (s *GetAggregateConfigRuleResponseBodyConfigRule) SetConfigRuleEvaluationStatus(v *GetAggregateConfigRuleResponseBodyConfigRuleConfigRuleEvaluationStatus) *GetAggregateConfigRuleResponseBodyConfigRule {
	s.ConfigRuleEvaluationStatus = v
	return s
}

func (s *GetAggregateConfigRuleResponseBodyConfigRule) SetConfigRuleId(v string) *GetAggregateConfigRuleResponseBodyConfigRule {
	s.ConfigRuleId = &v
	return s
}

func (s *GetAggregateConfigRuleResponseBodyConfigRule) SetConfigRuleName(v string) *GetAggregateConfigRuleResponseBodyConfigRule {
	s.ConfigRuleName = &v
	return s
}

func (s *GetAggregateConfigRuleResponseBodyConfigRule) SetConfigRuleState(v string) *GetAggregateConfigRuleResponseBodyConfigRule {
	s.ConfigRuleState = &v
	return s
}

func (s *GetAggregateConfigRuleResponseBodyConfigRule) SetConfigRuleTriggerTypes(v string) *GetAggregateConfigRuleResponseBodyConfigRule {
	s.ConfigRuleTriggerTypes = &v
	return s
}

func (s *GetAggregateConfigRuleResponseBodyConfigRule) SetCreateBy(v *GetAggregateConfigRuleResponseBodyConfigRuleCreateBy) *GetAggregateConfigRuleResponseBodyConfigRule {
	s.CreateBy = v
	return s
}

func (s *GetAggregateConfigRuleResponseBodyConfigRule) SetCreateTimestamp(v int64) *GetAggregateConfigRuleResponseBodyConfigRule {
	s.CreateTimestamp = &v
	return s
}

func (s *GetAggregateConfigRuleResponseBodyConfigRule) SetDescription(v string) *GetAggregateConfigRuleResponseBodyConfigRule {
	s.Description = &v
	return s
}

func (s *GetAggregateConfigRuleResponseBodyConfigRule) SetExcludeAccountIdsScope(v string) *GetAggregateConfigRuleResponseBodyConfigRule {
	s.ExcludeAccountIdsScope = &v
	return s
}

func (s *GetAggregateConfigRuleResponseBodyConfigRule) SetExcludeFolderIdsScope(v string) *GetAggregateConfigRuleResponseBodyConfigRule {
	s.ExcludeFolderIdsScope = &v
	return s
}

func (s *GetAggregateConfigRuleResponseBodyConfigRule) SetExcludeResourceIdsScope(v string) *GetAggregateConfigRuleResponseBodyConfigRule {
	s.ExcludeResourceIdsScope = &v
	return s
}

func (s *GetAggregateConfigRuleResponseBodyConfigRule) SetFolderIdsScope(v string) *GetAggregateConfigRuleResponseBodyConfigRule {
	s.FolderIdsScope = &v
	return s
}

func (s *GetAggregateConfigRuleResponseBodyConfigRule) SetInputParameters(v map[string]interface{}) *GetAggregateConfigRuleResponseBodyConfigRule {
	s.InputParameters = v
	return s
}

func (s *GetAggregateConfigRuleResponseBodyConfigRule) SetManagedRule(v *GetAggregateConfigRuleResponseBodyConfigRuleManagedRule) *GetAggregateConfigRuleResponseBodyConfigRule {
	s.ManagedRule = v
	return s
}

func (s *GetAggregateConfigRuleResponseBodyConfigRule) SetMaximumExecutionFrequency(v string) *GetAggregateConfigRuleResponseBodyConfigRule {
	s.MaximumExecutionFrequency = &v
	return s
}

func (s *GetAggregateConfigRuleResponseBodyConfigRule) SetModifiedTimestamp(v int64) *GetAggregateConfigRuleResponseBodyConfigRule {
	s.ModifiedTimestamp = &v
	return s
}

func (s *GetAggregateConfigRuleResponseBodyConfigRule) SetRegionIdsScope(v string) *GetAggregateConfigRuleResponseBodyConfigRule {
	s.RegionIdsScope = &v
	return s
}

func (s *GetAggregateConfigRuleResponseBodyConfigRule) SetResourceGroupIdsScope(v string) *GetAggregateConfigRuleResponseBodyConfigRule {
	s.ResourceGroupIdsScope = &v
	return s
}

func (s *GetAggregateConfigRuleResponseBodyConfigRule) SetResourceTypesScope(v string) *GetAggregateConfigRuleResponseBodyConfigRule {
	s.ResourceTypesScope = &v
	return s
}

func (s *GetAggregateConfigRuleResponseBodyConfigRule) SetRiskLevel(v int32) *GetAggregateConfigRuleResponseBodyConfigRule {
	s.RiskLevel = &v
	return s
}

func (s *GetAggregateConfigRuleResponseBodyConfigRule) SetSource(v *GetAggregateConfigRuleResponseBodyConfigRuleSource) *GetAggregateConfigRuleResponseBodyConfigRule {
	s.Source = v
	return s
}

func (s *GetAggregateConfigRuleResponseBodyConfigRule) SetTagKeyLogicScope(v string) *GetAggregateConfigRuleResponseBodyConfigRule {
	s.TagKeyLogicScope = &v
	return s
}

func (s *GetAggregateConfigRuleResponseBodyConfigRule) SetTagKeyScope(v string) *GetAggregateConfigRuleResponseBodyConfigRule {
	s.TagKeyScope = &v
	return s
}

func (s *GetAggregateConfigRuleResponseBodyConfigRule) SetTagValueScope(v string) *GetAggregateConfigRuleResponseBodyConfigRule {
	s.TagValueScope = &v
	return s
}

type GetAggregateConfigRuleResponseBodyConfigRuleCompliance struct {
	ComplianceType *string `json:"ComplianceType,omitempty" xml:"ComplianceType,omitempty"`
	Count          *int32  `json:"Count,omitempty" xml:"Count,omitempty"`
}

func (s GetAggregateConfigRuleResponseBodyConfigRuleCompliance) String() string {
	return tea.Prettify(s)
}

func (s GetAggregateConfigRuleResponseBodyConfigRuleCompliance) GoString() string {
	return s.String()
}

func (s *GetAggregateConfigRuleResponseBodyConfigRuleCompliance) SetComplianceType(v string) *GetAggregateConfigRuleResponseBodyConfigRuleCompliance {
	s.ComplianceType = &v
	return s
}

func (s *GetAggregateConfigRuleResponseBodyConfigRuleCompliance) SetCount(v int32) *GetAggregateConfigRuleResponseBodyConfigRuleCompliance {
	s.Count = &v
	return s
}

type GetAggregateConfigRuleResponseBodyConfigRuleConfigRuleEvaluationStatus struct {
	FirstActivatedTimestamp           *int64  `json:"FirstActivatedTimestamp,omitempty" xml:"FirstActivatedTimestamp,omitempty"`
	FirstEvaluationStarted            *bool   `json:"FirstEvaluationStarted,omitempty" xml:"FirstEvaluationStarted,omitempty"`
	LastErrorCode                     *string `json:"LastErrorCode,omitempty" xml:"LastErrorCode,omitempty"`
	LastErrorMessage                  *string `json:"LastErrorMessage,omitempty" xml:"LastErrorMessage,omitempty"`
	LastFailedEvaluationTimestamp     *int64  `json:"LastFailedEvaluationTimestamp,omitempty" xml:"LastFailedEvaluationTimestamp,omitempty"`
	LastFailedInvocationTimestamp     *int64  `json:"LastFailedInvocationTimestamp,omitempty" xml:"LastFailedInvocationTimestamp,omitempty"`
	LastSuccessfulEvaluationTimestamp *int64  `json:"LastSuccessfulEvaluationTimestamp,omitempty" xml:"LastSuccessfulEvaluationTimestamp,omitempty"`
	LastSuccessfulInvocationTimestamp *int64  `json:"LastSuccessfulInvocationTimestamp,omitempty" xml:"LastSuccessfulInvocationTimestamp,omitempty"`
}

func (s GetAggregateConfigRuleResponseBodyConfigRuleConfigRuleEvaluationStatus) String() string {
	return tea.Prettify(s)
}

func (s GetAggregateConfigRuleResponseBodyConfigRuleConfigRuleEvaluationStatus) GoString() string {
	return s.String()
}

func (s *GetAggregateConfigRuleResponseBodyConfigRuleConfigRuleEvaluationStatus) SetFirstActivatedTimestamp(v int64) *GetAggregateConfigRuleResponseBodyConfigRuleConfigRuleEvaluationStatus {
	s.FirstActivatedTimestamp = &v
	return s
}

func (s *GetAggregateConfigRuleResponseBodyConfigRuleConfigRuleEvaluationStatus) SetFirstEvaluationStarted(v bool) *GetAggregateConfigRuleResponseBodyConfigRuleConfigRuleEvaluationStatus {
	s.FirstEvaluationStarted = &v
	return s
}

func (s *GetAggregateConfigRuleResponseBodyConfigRuleConfigRuleEvaluationStatus) SetLastErrorCode(v string) *GetAggregateConfigRuleResponseBodyConfigRuleConfigRuleEvaluationStatus {
	s.LastErrorCode = &v
	return s
}

func (s *GetAggregateConfigRuleResponseBodyConfigRuleConfigRuleEvaluationStatus) SetLastErrorMessage(v string) *GetAggregateConfigRuleResponseBodyConfigRuleConfigRuleEvaluationStatus {
	s.LastErrorMessage = &v
	return s
}

func (s *GetAggregateConfigRuleResponseBodyConfigRuleConfigRuleEvaluationStatus) SetLastFailedEvaluationTimestamp(v int64) *GetAggregateConfigRuleResponseBodyConfigRuleConfigRuleEvaluationStatus {
	s.LastFailedEvaluationTimestamp = &v
	return s
}

func (s *GetAggregateConfigRuleResponseBodyConfigRuleConfigRuleEvaluationStatus) SetLastFailedInvocationTimestamp(v int64) *GetAggregateConfigRuleResponseBodyConfigRuleConfigRuleEvaluationStatus {
	s.LastFailedInvocationTimestamp = &v
	return s
}

func (s *GetAggregateConfigRuleResponseBodyConfigRuleConfigRuleEvaluationStatus) SetLastSuccessfulEvaluationTimestamp(v int64) *GetAggregateConfigRuleResponseBodyConfigRuleConfigRuleEvaluationStatus {
	s.LastSuccessfulEvaluationTimestamp = &v
	return s
}

func (s *GetAggregateConfigRuleResponseBodyConfigRuleConfigRuleEvaluationStatus) SetLastSuccessfulInvocationTimestamp(v int64) *GetAggregateConfigRuleResponseBodyConfigRuleConfigRuleEvaluationStatus {
	s.LastSuccessfulInvocationTimestamp = &v
	return s
}

type GetAggregateConfigRuleResponseBodyConfigRuleCreateBy struct {
	AggregatorId       *string `json:"AggregatorId,omitempty" xml:"AggregatorId,omitempty"`
	AggregatorName     *string `json:"AggregatorName,omitempty" xml:"AggregatorName,omitempty"`
	CompliancePackId   *string `json:"CompliancePackId,omitempty" xml:"CompliancePackId,omitempty"`
	CompliancePackName *string `json:"CompliancePackName,omitempty" xml:"CompliancePackName,omitempty"`
	CreatorId          *string `json:"CreatorId,omitempty" xml:"CreatorId,omitempty"`
	CreatorName        *string `json:"CreatorName,omitempty" xml:"CreatorName,omitempty"`
	CreatorType        *string `json:"CreatorType,omitempty" xml:"CreatorType,omitempty"`
}

func (s GetAggregateConfigRuleResponseBodyConfigRuleCreateBy) String() string {
	return tea.Prettify(s)
}

func (s GetAggregateConfigRuleResponseBodyConfigRuleCreateBy) GoString() string {
	return s.String()
}

func (s *GetAggregateConfigRuleResponseBodyConfigRuleCreateBy) SetAggregatorId(v string) *GetAggregateConfigRuleResponseBodyConfigRuleCreateBy {
	s.AggregatorId = &v
	return s
}

func (s *GetAggregateConfigRuleResponseBodyConfigRuleCreateBy) SetAggregatorName(v string) *GetAggregateConfigRuleResponseBodyConfigRuleCreateBy {
	s.AggregatorName = &v
	return s
}

func (s *GetAggregateConfigRuleResponseBodyConfigRuleCreateBy) SetCompliancePackId(v string) *GetAggregateConfigRuleResponseBodyConfigRuleCreateBy {
	s.CompliancePackId = &v
	return s
}

func (s *GetAggregateConfigRuleResponseBodyConfigRuleCreateBy) SetCompliancePackName(v string) *GetAggregateConfigRuleResponseBodyConfigRuleCreateBy {
	s.CompliancePackName = &v
	return s
}

func (s *GetAggregateConfigRuleResponseBodyConfigRuleCreateBy) SetCreatorId(v string) *GetAggregateConfigRuleResponseBodyConfigRuleCreateBy {
	s.CreatorId = &v
	return s
}

func (s *GetAggregateConfigRuleResponseBodyConfigRuleCreateBy) SetCreatorName(v string) *GetAggregateConfigRuleResponseBodyConfigRuleCreateBy {
	s.CreatorName = &v
	return s
}

func (s *GetAggregateConfigRuleResponseBodyConfigRuleCreateBy) SetCreatorType(v string) *GetAggregateConfigRuleResponseBodyConfigRuleCreateBy {
	s.CreatorType = &v
	return s
}

type GetAggregateConfigRuleResponseBodyConfigRuleManagedRule struct {
	CompulsoryInputParameterDetails map[string]interface{}                                                  `json:"CompulsoryInputParameterDetails,omitempty" xml:"CompulsoryInputParameterDetails,omitempty"`
	Description                     *string                                                                 `json:"Description,omitempty" xml:"Description,omitempty"`
	Identifier                      *string                                                                 `json:"Identifier,omitempty" xml:"Identifier,omitempty"`
	Labels                          []*string                                                               `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	ManagedRuleName                 *string                                                                 `json:"ManagedRuleName,omitempty" xml:"ManagedRuleName,omitempty"`
	OptionalInputParameterDetails   map[string]interface{}                                                  `json:"OptionalInputParameterDetails,omitempty" xml:"OptionalInputParameterDetails,omitempty"`
	SourceDetails                   []*GetAggregateConfigRuleResponseBodyConfigRuleManagedRuleSourceDetails `json:"SourceDetails,omitempty" xml:"SourceDetails,omitempty" type:"Repeated"`
}

func (s GetAggregateConfigRuleResponseBodyConfigRuleManagedRule) String() string {
	return tea.Prettify(s)
}

func (s GetAggregateConfigRuleResponseBodyConfigRuleManagedRule) GoString() string {
	return s.String()
}

func (s *GetAggregateConfigRuleResponseBodyConfigRuleManagedRule) SetCompulsoryInputParameterDetails(v map[string]interface{}) *GetAggregateConfigRuleResponseBodyConfigRuleManagedRule {
	s.CompulsoryInputParameterDetails = v
	return s
}

func (s *GetAggregateConfigRuleResponseBodyConfigRuleManagedRule) SetDescription(v string) *GetAggregateConfigRuleResponseBodyConfigRuleManagedRule {
	s.Description = &v
	return s
}

func (s *GetAggregateConfigRuleResponseBodyConfigRuleManagedRule) SetIdentifier(v string) *GetAggregateConfigRuleResponseBodyConfigRuleManagedRule {
	s.Identifier = &v
	return s
}

func (s *GetAggregateConfigRuleResponseBodyConfigRuleManagedRule) SetLabels(v []*string) *GetAggregateConfigRuleResponseBodyConfigRuleManagedRule {
	s.Labels = v
	return s
}

func (s *GetAggregateConfigRuleResponseBodyConfigRuleManagedRule) SetManagedRuleName(v string) *GetAggregateConfigRuleResponseBodyConfigRuleManagedRule {
	s.ManagedRuleName = &v
	return s
}

func (s *GetAggregateConfigRuleResponseBodyConfigRuleManagedRule) SetOptionalInputParameterDetails(v map[string]interface{}) *GetAggregateConfigRuleResponseBodyConfigRuleManagedRule {
	s.OptionalInputParameterDetails = v
	return s
}

func (s *GetAggregateConfigRuleResponseBodyConfigRuleManagedRule) SetSourceDetails(v []*GetAggregateConfigRuleResponseBodyConfigRuleManagedRuleSourceDetails) *GetAggregateConfigRuleResponseBodyConfigRuleManagedRule {
	s.SourceDetails = v
	return s
}

type GetAggregateConfigRuleResponseBodyConfigRuleManagedRuleSourceDetails struct {
	EventSource               *string `json:"EventSource,omitempty" xml:"EventSource,omitempty"`
	MaximumExecutionFrequency *string `json:"MaximumExecutionFrequency,omitempty" xml:"MaximumExecutionFrequency,omitempty"`
	MessageType               *string `json:"MessageType,omitempty" xml:"MessageType,omitempty"`
}

func (s GetAggregateConfigRuleResponseBodyConfigRuleManagedRuleSourceDetails) String() string {
	return tea.Prettify(s)
}

func (s GetAggregateConfigRuleResponseBodyConfigRuleManagedRuleSourceDetails) GoString() string {
	return s.String()
}

func (s *GetAggregateConfigRuleResponseBodyConfigRuleManagedRuleSourceDetails) SetEventSource(v string) *GetAggregateConfigRuleResponseBodyConfigRuleManagedRuleSourceDetails {
	s.EventSource = &v
	return s
}

func (s *GetAggregateConfigRuleResponseBodyConfigRuleManagedRuleSourceDetails) SetMaximumExecutionFrequency(v string) *GetAggregateConfigRuleResponseBodyConfigRuleManagedRuleSourceDetails {
	s.MaximumExecutionFrequency = &v
	return s
}

func (s *GetAggregateConfigRuleResponseBodyConfigRuleManagedRuleSourceDetails) SetMessageType(v string) *GetAggregateConfigRuleResponseBodyConfigRuleManagedRuleSourceDetails {
	s.MessageType = &v
	return s
}

type GetAggregateConfigRuleResponseBodyConfigRuleSource struct {
	Identifier    *string                                                            `json:"Identifier,omitempty" xml:"Identifier,omitempty"`
	Owner         *string                                                            `json:"Owner,omitempty" xml:"Owner,omitempty"`
	SourceDetails []*GetAggregateConfigRuleResponseBodyConfigRuleSourceSourceDetails `json:"SourceDetails,omitempty" xml:"SourceDetails,omitempty" type:"Repeated"`
}

func (s GetAggregateConfigRuleResponseBodyConfigRuleSource) String() string {
	return tea.Prettify(s)
}

func (s GetAggregateConfigRuleResponseBodyConfigRuleSource) GoString() string {
	return s.String()
}

func (s *GetAggregateConfigRuleResponseBodyConfigRuleSource) SetIdentifier(v string) *GetAggregateConfigRuleResponseBodyConfigRuleSource {
	s.Identifier = &v
	return s
}

func (s *GetAggregateConfigRuleResponseBodyConfigRuleSource) SetOwner(v string) *GetAggregateConfigRuleResponseBodyConfigRuleSource {
	s.Owner = &v
	return s
}

func (s *GetAggregateConfigRuleResponseBodyConfigRuleSource) SetSourceDetails(v []*GetAggregateConfigRuleResponseBodyConfigRuleSourceSourceDetails) *GetAggregateConfigRuleResponseBodyConfigRuleSource {
	s.SourceDetails = v
	return s
}

type GetAggregateConfigRuleResponseBodyConfigRuleSourceSourceDetails struct {
	EventSource               *string `json:"EventSource,omitempty" xml:"EventSource,omitempty"`
	MaximumExecutionFrequency *string `json:"MaximumExecutionFrequency,omitempty" xml:"MaximumExecutionFrequency,omitempty"`
	MessageType               *string `json:"MessageType,omitempty" xml:"MessageType,omitempty"`
}

func (s GetAggregateConfigRuleResponseBodyConfigRuleSourceSourceDetails) String() string {
	return tea.Prettify(s)
}

func (s GetAggregateConfigRuleResponseBodyConfigRuleSourceSourceDetails) GoString() string {
	return s.String()
}

func (s *GetAggregateConfigRuleResponseBodyConfigRuleSourceSourceDetails) SetEventSource(v string) *GetAggregateConfigRuleResponseBodyConfigRuleSourceSourceDetails {
	s.EventSource = &v
	return s
}

func (s *GetAggregateConfigRuleResponseBodyConfigRuleSourceSourceDetails) SetMaximumExecutionFrequency(v string) *GetAggregateConfigRuleResponseBodyConfigRuleSourceSourceDetails {
	s.MaximumExecutionFrequency = &v
	return s
}

func (s *GetAggregateConfigRuleResponseBodyConfigRuleSourceSourceDetails) SetMessageType(v string) *GetAggregateConfigRuleResponseBodyConfigRuleSourceSourceDetails {
	s.MessageType = &v
	return s
}

type GetAggregateConfigRuleResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetAggregateConfigRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetAggregateConfigRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAggregateConfigRuleResponse) GoString() string {
	return s.String()
}

func (s *GetAggregateConfigRuleResponse) SetHeaders(v map[string]*string) *GetAggregateConfigRuleResponse {
	s.Headers = v
	return s
}

func (s *GetAggregateConfigRuleResponse) SetStatusCode(v int32) *GetAggregateConfigRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAggregateConfigRuleResponse) SetBody(v *GetAggregateConfigRuleResponseBody) *GetAggregateConfigRuleResponse {
	s.Body = v
	return s
}

type GetAggregateConfigRuleComplianceByPackRequest struct {
	AggregatorId     *string `json:"AggregatorId,omitempty" xml:"AggregatorId,omitempty"`
	CompliancePackId *string `json:"CompliancePackId,omitempty" xml:"CompliancePackId,omitempty"`
}

func (s GetAggregateConfigRuleComplianceByPackRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAggregateConfigRuleComplianceByPackRequest) GoString() string {
	return s.String()
}

func (s *GetAggregateConfigRuleComplianceByPackRequest) SetAggregatorId(v string) *GetAggregateConfigRuleComplianceByPackRequest {
	s.AggregatorId = &v
	return s
}

func (s *GetAggregateConfigRuleComplianceByPackRequest) SetCompliancePackId(v string) *GetAggregateConfigRuleComplianceByPackRequest {
	s.CompliancePackId = &v
	return s
}

type GetAggregateConfigRuleComplianceByPackResponseBody struct {
	ConfigRuleComplianceResult *GetAggregateConfigRuleComplianceByPackResponseBodyConfigRuleComplianceResult `json:"ConfigRuleComplianceResult,omitempty" xml:"ConfigRuleComplianceResult,omitempty" type:"Struct"`
	RequestId                  *string                                                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetAggregateConfigRuleComplianceByPackResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAggregateConfigRuleComplianceByPackResponseBody) GoString() string {
	return s.String()
}

func (s *GetAggregateConfigRuleComplianceByPackResponseBody) SetConfigRuleComplianceResult(v *GetAggregateConfigRuleComplianceByPackResponseBodyConfigRuleComplianceResult) *GetAggregateConfigRuleComplianceByPackResponseBody {
	s.ConfigRuleComplianceResult = v
	return s
}

func (s *GetAggregateConfigRuleComplianceByPackResponseBody) SetRequestId(v string) *GetAggregateConfigRuleComplianceByPackResponseBody {
	s.RequestId = &v
	return s
}

type GetAggregateConfigRuleComplianceByPackResponseBodyConfigRuleComplianceResult struct {
	CompliancePackId      *string                                                                                              `json:"CompliancePackId,omitempty" xml:"CompliancePackId,omitempty"`
	ConfigRuleCompliances []*GetAggregateConfigRuleComplianceByPackResponseBodyConfigRuleComplianceResultConfigRuleCompliances `json:"ConfigRuleCompliances,omitempty" xml:"ConfigRuleCompliances,omitempty" type:"Repeated"`
	NonCompliantCount     *int32                                                                                               `json:"NonCompliantCount,omitempty" xml:"NonCompliantCount,omitempty"`
	TotalCount            *int32                                                                                               `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s GetAggregateConfigRuleComplianceByPackResponseBodyConfigRuleComplianceResult) String() string {
	return tea.Prettify(s)
}

func (s GetAggregateConfigRuleComplianceByPackResponseBodyConfigRuleComplianceResult) GoString() string {
	return s.String()
}

func (s *GetAggregateConfigRuleComplianceByPackResponseBodyConfigRuleComplianceResult) SetCompliancePackId(v string) *GetAggregateConfigRuleComplianceByPackResponseBodyConfigRuleComplianceResult {
	s.CompliancePackId = &v
	return s
}

func (s *GetAggregateConfigRuleComplianceByPackResponseBodyConfigRuleComplianceResult) SetConfigRuleCompliances(v []*GetAggregateConfigRuleComplianceByPackResponseBodyConfigRuleComplianceResultConfigRuleCompliances) *GetAggregateConfigRuleComplianceByPackResponseBodyConfigRuleComplianceResult {
	s.ConfigRuleCompliances = v
	return s
}

func (s *GetAggregateConfigRuleComplianceByPackResponseBodyConfigRuleComplianceResult) SetNonCompliantCount(v int32) *GetAggregateConfigRuleComplianceByPackResponseBodyConfigRuleComplianceResult {
	s.NonCompliantCount = &v
	return s
}

func (s *GetAggregateConfigRuleComplianceByPackResponseBodyConfigRuleComplianceResult) SetTotalCount(v int32) *GetAggregateConfigRuleComplianceByPackResponseBodyConfigRuleComplianceResult {
	s.TotalCount = &v
	return s
}

type GetAggregateConfigRuleComplianceByPackResponseBodyConfigRuleComplianceResultConfigRuleCompliances struct {
	ComplianceType *string `json:"ComplianceType,omitempty" xml:"ComplianceType,omitempty"`
	ConfigRuleId   *string `json:"ConfigRuleId,omitempty" xml:"ConfigRuleId,omitempty"`
	ConfigRuleName *string `json:"ConfigRuleName,omitempty" xml:"ConfigRuleName,omitempty"`
}

func (s GetAggregateConfigRuleComplianceByPackResponseBodyConfigRuleComplianceResultConfigRuleCompliances) String() string {
	return tea.Prettify(s)
}

func (s GetAggregateConfigRuleComplianceByPackResponseBodyConfigRuleComplianceResultConfigRuleCompliances) GoString() string {
	return s.String()
}

func (s *GetAggregateConfigRuleComplianceByPackResponseBodyConfigRuleComplianceResultConfigRuleCompliances) SetComplianceType(v string) *GetAggregateConfigRuleComplianceByPackResponseBodyConfigRuleComplianceResultConfigRuleCompliances {
	s.ComplianceType = &v
	return s
}

func (s *GetAggregateConfigRuleComplianceByPackResponseBodyConfigRuleComplianceResultConfigRuleCompliances) SetConfigRuleId(v string) *GetAggregateConfigRuleComplianceByPackResponseBodyConfigRuleComplianceResultConfigRuleCompliances {
	s.ConfigRuleId = &v
	return s
}

func (s *GetAggregateConfigRuleComplianceByPackResponseBodyConfigRuleComplianceResultConfigRuleCompliances) SetConfigRuleName(v string) *GetAggregateConfigRuleComplianceByPackResponseBodyConfigRuleComplianceResultConfigRuleCompliances {
	s.ConfigRuleName = &v
	return s
}

type GetAggregateConfigRuleComplianceByPackResponse struct {
	Headers    map[string]*string                                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetAggregateConfigRuleComplianceByPackResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetAggregateConfigRuleComplianceByPackResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAggregateConfigRuleComplianceByPackResponse) GoString() string {
	return s.String()
}

func (s *GetAggregateConfigRuleComplianceByPackResponse) SetHeaders(v map[string]*string) *GetAggregateConfigRuleComplianceByPackResponse {
	s.Headers = v
	return s
}

func (s *GetAggregateConfigRuleComplianceByPackResponse) SetStatusCode(v int32) *GetAggregateConfigRuleComplianceByPackResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAggregateConfigRuleComplianceByPackResponse) SetBody(v *GetAggregateConfigRuleComplianceByPackResponseBody) *GetAggregateConfigRuleComplianceByPackResponse {
	s.Body = v
	return s
}

type GetAggregateConfigRuleSummaryByRiskLevelRequest struct {
	AggregatorId *string `json:"AggregatorId,omitempty" xml:"AggregatorId,omitempty"`
}

func (s GetAggregateConfigRuleSummaryByRiskLevelRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAggregateConfigRuleSummaryByRiskLevelRequest) GoString() string {
	return s.String()
}

func (s *GetAggregateConfigRuleSummaryByRiskLevelRequest) SetAggregatorId(v string) *GetAggregateConfigRuleSummaryByRiskLevelRequest {
	s.AggregatorId = &v
	return s
}

type GetAggregateConfigRuleSummaryByRiskLevelResponseBody struct {
	ConfigRuleSummaries []*GetAggregateConfigRuleSummaryByRiskLevelResponseBodyConfigRuleSummaries `json:"ConfigRuleSummaries,omitempty" xml:"ConfigRuleSummaries,omitempty" type:"Repeated"`
	RequestId           *string                                                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetAggregateConfigRuleSummaryByRiskLevelResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAggregateConfigRuleSummaryByRiskLevelResponseBody) GoString() string {
	return s.String()
}

func (s *GetAggregateConfigRuleSummaryByRiskLevelResponseBody) SetConfigRuleSummaries(v []*GetAggregateConfigRuleSummaryByRiskLevelResponseBodyConfigRuleSummaries) *GetAggregateConfigRuleSummaryByRiskLevelResponseBody {
	s.ConfigRuleSummaries = v
	return s
}

func (s *GetAggregateConfigRuleSummaryByRiskLevelResponseBody) SetRequestId(v string) *GetAggregateConfigRuleSummaryByRiskLevelResponseBody {
	s.RequestId = &v
	return s
}

type GetAggregateConfigRuleSummaryByRiskLevelResponseBodyConfigRuleSummaries struct {
	CompliantCount    *int32 `json:"CompliantCount,omitempty" xml:"CompliantCount,omitempty"`
	NonCompliantCount *int32 `json:"NonCompliantCount,omitempty" xml:"NonCompliantCount,omitempty"`
	RiskLevel         *int32 `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
}

func (s GetAggregateConfigRuleSummaryByRiskLevelResponseBodyConfigRuleSummaries) String() string {
	return tea.Prettify(s)
}

func (s GetAggregateConfigRuleSummaryByRiskLevelResponseBodyConfigRuleSummaries) GoString() string {
	return s.String()
}

func (s *GetAggregateConfigRuleSummaryByRiskLevelResponseBodyConfigRuleSummaries) SetCompliantCount(v int32) *GetAggregateConfigRuleSummaryByRiskLevelResponseBodyConfigRuleSummaries {
	s.CompliantCount = &v
	return s
}

func (s *GetAggregateConfigRuleSummaryByRiskLevelResponseBodyConfigRuleSummaries) SetNonCompliantCount(v int32) *GetAggregateConfigRuleSummaryByRiskLevelResponseBodyConfigRuleSummaries {
	s.NonCompliantCount = &v
	return s
}

func (s *GetAggregateConfigRuleSummaryByRiskLevelResponseBodyConfigRuleSummaries) SetRiskLevel(v int32) *GetAggregateConfigRuleSummaryByRiskLevelResponseBodyConfigRuleSummaries {
	s.RiskLevel = &v
	return s
}

type GetAggregateConfigRuleSummaryByRiskLevelResponse struct {
	Headers    map[string]*string                                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetAggregateConfigRuleSummaryByRiskLevelResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetAggregateConfigRuleSummaryByRiskLevelResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAggregateConfigRuleSummaryByRiskLevelResponse) GoString() string {
	return s.String()
}

func (s *GetAggregateConfigRuleSummaryByRiskLevelResponse) SetHeaders(v map[string]*string) *GetAggregateConfigRuleSummaryByRiskLevelResponse {
	s.Headers = v
	return s
}

func (s *GetAggregateConfigRuleSummaryByRiskLevelResponse) SetStatusCode(v int32) *GetAggregateConfigRuleSummaryByRiskLevelResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAggregateConfigRuleSummaryByRiskLevelResponse) SetBody(v *GetAggregateConfigRuleSummaryByRiskLevelResponseBody) *GetAggregateConfigRuleSummaryByRiskLevelResponse {
	s.Body = v
	return s
}

type GetAggregateConfigRulesReportRequest struct {
	AggregatorId *string `json:"AggregatorId,omitempty" xml:"AggregatorId,omitempty"`
	ReportId     *string `json:"ReportId,omitempty" xml:"ReportId,omitempty"`
}

func (s GetAggregateConfigRulesReportRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAggregateConfigRulesReportRequest) GoString() string {
	return s.String()
}

func (s *GetAggregateConfigRulesReportRequest) SetAggregatorId(v string) *GetAggregateConfigRulesReportRequest {
	s.AggregatorId = &v
	return s
}

func (s *GetAggregateConfigRulesReportRequest) SetReportId(v string) *GetAggregateConfigRulesReportRequest {
	s.ReportId = &v
	return s
}

type GetAggregateConfigRulesReportResponseBody struct {
	ConfigRulesReport *GetAggregateConfigRulesReportResponseBodyConfigRulesReport `json:"ConfigRulesReport,omitempty" xml:"ConfigRulesReport,omitempty" type:"Struct"`
	RequestId         *string                                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetAggregateConfigRulesReportResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAggregateConfigRulesReportResponseBody) GoString() string {
	return s.String()
}

func (s *GetAggregateConfigRulesReportResponseBody) SetConfigRulesReport(v *GetAggregateConfigRulesReportResponseBodyConfigRulesReport) *GetAggregateConfigRulesReportResponseBody {
	s.ConfigRulesReport = v
	return s
}

func (s *GetAggregateConfigRulesReportResponseBody) SetRequestId(v string) *GetAggregateConfigRulesReportResponseBody {
	s.RequestId = &v
	return s
}

type GetAggregateConfigRulesReportResponseBodyConfigRulesReport struct {
	AccountId             *int64  `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	AggregatorId          *string `json:"AggregatorId,omitempty" xml:"AggregatorId,omitempty"`
	ReportCreateTimestamp *int64  `json:"ReportCreateTimestamp,omitempty" xml:"ReportCreateTimestamp,omitempty"`
	ReportId              *string `json:"ReportId,omitempty" xml:"ReportId,omitempty"`
	ReportStatus          *string `json:"ReportStatus,omitempty" xml:"ReportStatus,omitempty"`
	ReportUrl             *string `json:"ReportUrl,omitempty" xml:"ReportUrl,omitempty"`
}

func (s GetAggregateConfigRulesReportResponseBodyConfigRulesReport) String() string {
	return tea.Prettify(s)
}

func (s GetAggregateConfigRulesReportResponseBodyConfigRulesReport) GoString() string {
	return s.String()
}

func (s *GetAggregateConfigRulesReportResponseBodyConfigRulesReport) SetAccountId(v int64) *GetAggregateConfigRulesReportResponseBodyConfigRulesReport {
	s.AccountId = &v
	return s
}

func (s *GetAggregateConfigRulesReportResponseBodyConfigRulesReport) SetAggregatorId(v string) *GetAggregateConfigRulesReportResponseBodyConfigRulesReport {
	s.AggregatorId = &v
	return s
}

func (s *GetAggregateConfigRulesReportResponseBodyConfigRulesReport) SetReportCreateTimestamp(v int64) *GetAggregateConfigRulesReportResponseBodyConfigRulesReport {
	s.ReportCreateTimestamp = &v
	return s
}

func (s *GetAggregateConfigRulesReportResponseBodyConfigRulesReport) SetReportId(v string) *GetAggregateConfigRulesReportResponseBodyConfigRulesReport {
	s.ReportId = &v
	return s
}

func (s *GetAggregateConfigRulesReportResponseBodyConfigRulesReport) SetReportStatus(v string) *GetAggregateConfigRulesReportResponseBodyConfigRulesReport {
	s.ReportStatus = &v
	return s
}

func (s *GetAggregateConfigRulesReportResponseBodyConfigRulesReport) SetReportUrl(v string) *GetAggregateConfigRulesReportResponseBodyConfigRulesReport {
	s.ReportUrl = &v
	return s
}

type GetAggregateConfigRulesReportResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetAggregateConfigRulesReportResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetAggregateConfigRulesReportResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAggregateConfigRulesReportResponse) GoString() string {
	return s.String()
}

func (s *GetAggregateConfigRulesReportResponse) SetHeaders(v map[string]*string) *GetAggregateConfigRulesReportResponse {
	s.Headers = v
	return s
}

func (s *GetAggregateConfigRulesReportResponse) SetStatusCode(v int32) *GetAggregateConfigRulesReportResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAggregateConfigRulesReportResponse) SetBody(v *GetAggregateConfigRulesReportResponseBody) *GetAggregateConfigRulesReportResponse {
	s.Body = v
	return s
}

type GetAggregateDiscoveredResourceRequest struct {
	AggregatorId    *string `json:"AggregatorId,omitempty" xml:"AggregatorId,omitempty"`
	Region          *string `json:"Region,omitempty" xml:"Region,omitempty"`
	ResourceId      *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ResourceType    *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
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
	DiscoveredResourceDetail *GetAggregateDiscoveredResourceResponseBodyDiscoveredResourceDetail `json:"DiscoveredResourceDetail,omitempty" xml:"DiscoveredResourceDetail,omitempty" type:"Struct"`
	RequestId                *string                                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	AccountId            *int64  `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	AvailabilityZone     *string `json:"AvailabilityZone,omitempty" xml:"AvailabilityZone,omitempty"`
	Configuration        *string `json:"Configuration,omitempty" xml:"Configuration,omitempty"`
	Region               *string `json:"Region,omitempty" xml:"Region,omitempty"`
	ResourceCreationTime *int64  `json:"ResourceCreationTime,omitempty" xml:"ResourceCreationTime,omitempty"`
	ResourceDeleted      *int32  `json:"ResourceDeleted,omitempty" xml:"ResourceDeleted,omitempty"`
	ResourceId           *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ResourceName         *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	ResourceStatus       *string `json:"ResourceStatus,omitempty" xml:"ResourceStatus,omitempty"`
	ResourceType         *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	Tags                 *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
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

type GetAggregateResourceComplianceByConfigRuleRequest struct {
	AggregatorId    *string `json:"AggregatorId,omitempty" xml:"AggregatorId,omitempty"`
	ComplianceType  *string `json:"ComplianceType,omitempty" xml:"ComplianceType,omitempty"`
	ConfigRuleId    *string `json:"ConfigRuleId,omitempty" xml:"ConfigRuleId,omitempty"`
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s GetAggregateResourceComplianceByConfigRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAggregateResourceComplianceByConfigRuleRequest) GoString() string {
	return s.String()
}

func (s *GetAggregateResourceComplianceByConfigRuleRequest) SetAggregatorId(v string) *GetAggregateResourceComplianceByConfigRuleRequest {
	s.AggregatorId = &v
	return s
}

func (s *GetAggregateResourceComplianceByConfigRuleRequest) SetComplianceType(v string) *GetAggregateResourceComplianceByConfigRuleRequest {
	s.ComplianceType = &v
	return s
}

func (s *GetAggregateResourceComplianceByConfigRuleRequest) SetConfigRuleId(v string) *GetAggregateResourceComplianceByConfigRuleRequest {
	s.ConfigRuleId = &v
	return s
}

func (s *GetAggregateResourceComplianceByConfigRuleRequest) SetResourceOwnerId(v int64) *GetAggregateResourceComplianceByConfigRuleRequest {
	s.ResourceOwnerId = &v
	return s
}

type GetAggregateResourceComplianceByConfigRuleResponseBody struct {
	ComplianceResult *GetAggregateResourceComplianceByConfigRuleResponseBodyComplianceResult `json:"ComplianceResult,omitempty" xml:"ComplianceResult,omitempty" type:"Struct"`
	RequestId        *string                                                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetAggregateResourceComplianceByConfigRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAggregateResourceComplianceByConfigRuleResponseBody) GoString() string {
	return s.String()
}

func (s *GetAggregateResourceComplianceByConfigRuleResponseBody) SetComplianceResult(v *GetAggregateResourceComplianceByConfigRuleResponseBodyComplianceResult) *GetAggregateResourceComplianceByConfigRuleResponseBody {
	s.ComplianceResult = v
	return s
}

func (s *GetAggregateResourceComplianceByConfigRuleResponseBody) SetRequestId(v string) *GetAggregateResourceComplianceByConfigRuleResponseBody {
	s.RequestId = &v
	return s
}

type GetAggregateResourceComplianceByConfigRuleResponseBodyComplianceResult struct {
	Compliances []*GetAggregateResourceComplianceByConfigRuleResponseBodyComplianceResultCompliances `json:"Compliances,omitempty" xml:"Compliances,omitempty" type:"Repeated"`
	TotalCount  *int64                                                                               `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s GetAggregateResourceComplianceByConfigRuleResponseBodyComplianceResult) String() string {
	return tea.Prettify(s)
}

func (s GetAggregateResourceComplianceByConfigRuleResponseBodyComplianceResult) GoString() string {
	return s.String()
}

func (s *GetAggregateResourceComplianceByConfigRuleResponseBodyComplianceResult) SetCompliances(v []*GetAggregateResourceComplianceByConfigRuleResponseBodyComplianceResultCompliances) *GetAggregateResourceComplianceByConfigRuleResponseBodyComplianceResult {
	s.Compliances = v
	return s
}

func (s *GetAggregateResourceComplianceByConfigRuleResponseBodyComplianceResult) SetTotalCount(v int64) *GetAggregateResourceComplianceByConfigRuleResponseBodyComplianceResult {
	s.TotalCount = &v
	return s
}

type GetAggregateResourceComplianceByConfigRuleResponseBodyComplianceResultCompliances struct {
	ComplianceType *string `json:"ComplianceType,omitempty" xml:"ComplianceType,omitempty"`
	Count          *int32  `json:"Count,omitempty" xml:"Count,omitempty"`
}

func (s GetAggregateResourceComplianceByConfigRuleResponseBodyComplianceResultCompliances) String() string {
	return tea.Prettify(s)
}

func (s GetAggregateResourceComplianceByConfigRuleResponseBodyComplianceResultCompliances) GoString() string {
	return s.String()
}

func (s *GetAggregateResourceComplianceByConfigRuleResponseBodyComplianceResultCompliances) SetComplianceType(v string) *GetAggregateResourceComplianceByConfigRuleResponseBodyComplianceResultCompliances {
	s.ComplianceType = &v
	return s
}

func (s *GetAggregateResourceComplianceByConfigRuleResponseBodyComplianceResultCompliances) SetCount(v int32) *GetAggregateResourceComplianceByConfigRuleResponseBodyComplianceResultCompliances {
	s.Count = &v
	return s
}

type GetAggregateResourceComplianceByConfigRuleResponse struct {
	Headers    map[string]*string                                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetAggregateResourceComplianceByConfigRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetAggregateResourceComplianceByConfigRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAggregateResourceComplianceByConfigRuleResponse) GoString() string {
	return s.String()
}

func (s *GetAggregateResourceComplianceByConfigRuleResponse) SetHeaders(v map[string]*string) *GetAggregateResourceComplianceByConfigRuleResponse {
	s.Headers = v
	return s
}

func (s *GetAggregateResourceComplianceByConfigRuleResponse) SetStatusCode(v int32) *GetAggregateResourceComplianceByConfigRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAggregateResourceComplianceByConfigRuleResponse) SetBody(v *GetAggregateResourceComplianceByConfigRuleResponseBody) *GetAggregateResourceComplianceByConfigRuleResponse {
	s.Body = v
	return s
}

type GetAggregateResourceComplianceByPackRequest struct {
	AggregatorId     *string `json:"AggregatorId,omitempty" xml:"AggregatorId,omitempty"`
	CompliancePackId *string `json:"CompliancePackId,omitempty" xml:"CompliancePackId,omitempty"`
}

func (s GetAggregateResourceComplianceByPackRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAggregateResourceComplianceByPackRequest) GoString() string {
	return s.String()
}

func (s *GetAggregateResourceComplianceByPackRequest) SetAggregatorId(v string) *GetAggregateResourceComplianceByPackRequest {
	s.AggregatorId = &v
	return s
}

func (s *GetAggregateResourceComplianceByPackRequest) SetCompliancePackId(v string) *GetAggregateResourceComplianceByPackRequest {
	s.CompliancePackId = &v
	return s
}

type GetAggregateResourceComplianceByPackResponseBody struct {
	RequestId                *string                                                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResourceComplianceResult *GetAggregateResourceComplianceByPackResponseBodyResourceComplianceResult `json:"ResourceComplianceResult,omitempty" xml:"ResourceComplianceResult,omitempty" type:"Struct"`
}

func (s GetAggregateResourceComplianceByPackResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAggregateResourceComplianceByPackResponseBody) GoString() string {
	return s.String()
}

func (s *GetAggregateResourceComplianceByPackResponseBody) SetRequestId(v string) *GetAggregateResourceComplianceByPackResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAggregateResourceComplianceByPackResponseBody) SetResourceComplianceResult(v *GetAggregateResourceComplianceByPackResponseBodyResourceComplianceResult) *GetAggregateResourceComplianceByPackResponseBody {
	s.ResourceComplianceResult = v
	return s
}

type GetAggregateResourceComplianceByPackResponseBodyResourceComplianceResult struct {
	CompliancePackId  *string `json:"CompliancePackId,omitempty" xml:"CompliancePackId,omitempty"`
	NonCompliantCount *int32  `json:"NonCompliantCount,omitempty" xml:"NonCompliantCount,omitempty"`
	TotalCount        *int32  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s GetAggregateResourceComplianceByPackResponseBodyResourceComplianceResult) String() string {
	return tea.Prettify(s)
}

func (s GetAggregateResourceComplianceByPackResponseBodyResourceComplianceResult) GoString() string {
	return s.String()
}

func (s *GetAggregateResourceComplianceByPackResponseBodyResourceComplianceResult) SetCompliancePackId(v string) *GetAggregateResourceComplianceByPackResponseBodyResourceComplianceResult {
	s.CompliancePackId = &v
	return s
}

func (s *GetAggregateResourceComplianceByPackResponseBodyResourceComplianceResult) SetNonCompliantCount(v int32) *GetAggregateResourceComplianceByPackResponseBodyResourceComplianceResult {
	s.NonCompliantCount = &v
	return s
}

func (s *GetAggregateResourceComplianceByPackResponseBodyResourceComplianceResult) SetTotalCount(v int32) *GetAggregateResourceComplianceByPackResponseBodyResourceComplianceResult {
	s.TotalCount = &v
	return s
}

type GetAggregateResourceComplianceByPackResponse struct {
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetAggregateResourceComplianceByPackResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetAggregateResourceComplianceByPackResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAggregateResourceComplianceByPackResponse) GoString() string {
	return s.String()
}

func (s *GetAggregateResourceComplianceByPackResponse) SetHeaders(v map[string]*string) *GetAggregateResourceComplianceByPackResponse {
	s.Headers = v
	return s
}

func (s *GetAggregateResourceComplianceByPackResponse) SetStatusCode(v int32) *GetAggregateResourceComplianceByPackResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAggregateResourceComplianceByPackResponse) SetBody(v *GetAggregateResourceComplianceByPackResponseBody) *GetAggregateResourceComplianceByPackResponse {
	s.Body = v
	return s
}

type GetAggregateResourceComplianceGroupByRegionRequest struct {
	AggregatorId  *string `json:"AggregatorId,omitempty" xml:"AggregatorId,omitempty"`
	ConfigRuleIds *string `json:"ConfigRuleIds,omitempty" xml:"ConfigRuleIds,omitempty"`
}

func (s GetAggregateResourceComplianceGroupByRegionRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAggregateResourceComplianceGroupByRegionRequest) GoString() string {
	return s.String()
}

func (s *GetAggregateResourceComplianceGroupByRegionRequest) SetAggregatorId(v string) *GetAggregateResourceComplianceGroupByRegionRequest {
	s.AggregatorId = &v
	return s
}

func (s *GetAggregateResourceComplianceGroupByRegionRequest) SetConfigRuleIds(v string) *GetAggregateResourceComplianceGroupByRegionRequest {
	s.ConfigRuleIds = &v
	return s
}

type GetAggregateResourceComplianceGroupByRegionResponseBody struct {
	ComplianceResult *GetAggregateResourceComplianceGroupByRegionResponseBodyComplianceResult `json:"ComplianceResult,omitempty" xml:"ComplianceResult,omitempty" type:"Struct"`
	RequestId        *string                                                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetAggregateResourceComplianceGroupByRegionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAggregateResourceComplianceGroupByRegionResponseBody) GoString() string {
	return s.String()
}

func (s *GetAggregateResourceComplianceGroupByRegionResponseBody) SetComplianceResult(v *GetAggregateResourceComplianceGroupByRegionResponseBodyComplianceResult) *GetAggregateResourceComplianceGroupByRegionResponseBody {
	s.ComplianceResult = v
	return s
}

func (s *GetAggregateResourceComplianceGroupByRegionResponseBody) SetRequestId(v string) *GetAggregateResourceComplianceGroupByRegionResponseBody {
	s.RequestId = &v
	return s
}

type GetAggregateResourceComplianceGroupByRegionResponseBodyComplianceResult struct {
	ComplianceResultList []*GetAggregateResourceComplianceGroupByRegionResponseBodyComplianceResultComplianceResultList `json:"ComplianceResultList,omitempty" xml:"ComplianceResultList,omitempty" type:"Repeated"`
}

func (s GetAggregateResourceComplianceGroupByRegionResponseBodyComplianceResult) String() string {
	return tea.Prettify(s)
}

func (s GetAggregateResourceComplianceGroupByRegionResponseBodyComplianceResult) GoString() string {
	return s.String()
}

func (s *GetAggregateResourceComplianceGroupByRegionResponseBodyComplianceResult) SetComplianceResultList(v []*GetAggregateResourceComplianceGroupByRegionResponseBodyComplianceResultComplianceResultList) *GetAggregateResourceComplianceGroupByRegionResponseBodyComplianceResult {
	s.ComplianceResultList = v
	return s
}

type GetAggregateResourceComplianceGroupByRegionResponseBodyComplianceResultComplianceResultList struct {
	Compliances []*GetAggregateResourceComplianceGroupByRegionResponseBodyComplianceResultComplianceResultListCompliances `json:"Compliances,omitempty" xml:"Compliances,omitempty" type:"Repeated"`
	RegionId    *string                                                                                                   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetAggregateResourceComplianceGroupByRegionResponseBodyComplianceResultComplianceResultList) String() string {
	return tea.Prettify(s)
}

func (s GetAggregateResourceComplianceGroupByRegionResponseBodyComplianceResultComplianceResultList) GoString() string {
	return s.String()
}

func (s *GetAggregateResourceComplianceGroupByRegionResponseBodyComplianceResultComplianceResultList) SetCompliances(v []*GetAggregateResourceComplianceGroupByRegionResponseBodyComplianceResultComplianceResultListCompliances) *GetAggregateResourceComplianceGroupByRegionResponseBodyComplianceResultComplianceResultList {
	s.Compliances = v
	return s
}

func (s *GetAggregateResourceComplianceGroupByRegionResponseBodyComplianceResultComplianceResultList) SetRegionId(v string) *GetAggregateResourceComplianceGroupByRegionResponseBodyComplianceResultComplianceResultList {
	s.RegionId = &v
	return s
}

type GetAggregateResourceComplianceGroupByRegionResponseBodyComplianceResultComplianceResultListCompliances struct {
	ComplianceType *string `json:"ComplianceType,omitempty" xml:"ComplianceType,omitempty"`
	Count          *int64  `json:"Count,omitempty" xml:"Count,omitempty"`
}

func (s GetAggregateResourceComplianceGroupByRegionResponseBodyComplianceResultComplianceResultListCompliances) String() string {
	return tea.Prettify(s)
}

func (s GetAggregateResourceComplianceGroupByRegionResponseBodyComplianceResultComplianceResultListCompliances) GoString() string {
	return s.String()
}

func (s *GetAggregateResourceComplianceGroupByRegionResponseBodyComplianceResultComplianceResultListCompliances) SetComplianceType(v string) *GetAggregateResourceComplianceGroupByRegionResponseBodyComplianceResultComplianceResultListCompliances {
	s.ComplianceType = &v
	return s
}

func (s *GetAggregateResourceComplianceGroupByRegionResponseBodyComplianceResultComplianceResultListCompliances) SetCount(v int64) *GetAggregateResourceComplianceGroupByRegionResponseBodyComplianceResultComplianceResultListCompliances {
	s.Count = &v
	return s
}

type GetAggregateResourceComplianceGroupByRegionResponse struct {
	Headers    map[string]*string                                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetAggregateResourceComplianceGroupByRegionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetAggregateResourceComplianceGroupByRegionResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAggregateResourceComplianceGroupByRegionResponse) GoString() string {
	return s.String()
}

func (s *GetAggregateResourceComplianceGroupByRegionResponse) SetHeaders(v map[string]*string) *GetAggregateResourceComplianceGroupByRegionResponse {
	s.Headers = v
	return s
}

func (s *GetAggregateResourceComplianceGroupByRegionResponse) SetStatusCode(v int32) *GetAggregateResourceComplianceGroupByRegionResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAggregateResourceComplianceGroupByRegionResponse) SetBody(v *GetAggregateResourceComplianceGroupByRegionResponseBody) *GetAggregateResourceComplianceGroupByRegionResponse {
	s.Body = v
	return s
}

type GetAggregateResourceComplianceGroupByResourceTypeRequest struct {
	AggregatorId  *string `json:"AggregatorId,omitempty" xml:"AggregatorId,omitempty"`
	ConfigRuleIds *string `json:"ConfigRuleIds,omitempty" xml:"ConfigRuleIds,omitempty"`
}

func (s GetAggregateResourceComplianceGroupByResourceTypeRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAggregateResourceComplianceGroupByResourceTypeRequest) GoString() string {
	return s.String()
}

func (s *GetAggregateResourceComplianceGroupByResourceTypeRequest) SetAggregatorId(v string) *GetAggregateResourceComplianceGroupByResourceTypeRequest {
	s.AggregatorId = &v
	return s
}

func (s *GetAggregateResourceComplianceGroupByResourceTypeRequest) SetConfigRuleIds(v string) *GetAggregateResourceComplianceGroupByResourceTypeRequest {
	s.ConfigRuleIds = &v
	return s
}

type GetAggregateResourceComplianceGroupByResourceTypeResponseBody struct {
	ComplianceResult *GetAggregateResourceComplianceGroupByResourceTypeResponseBodyComplianceResult `json:"ComplianceResult,omitempty" xml:"ComplianceResult,omitempty" type:"Struct"`
	RequestId        *string                                                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetAggregateResourceComplianceGroupByResourceTypeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAggregateResourceComplianceGroupByResourceTypeResponseBody) GoString() string {
	return s.String()
}

func (s *GetAggregateResourceComplianceGroupByResourceTypeResponseBody) SetComplianceResult(v *GetAggregateResourceComplianceGroupByResourceTypeResponseBodyComplianceResult) *GetAggregateResourceComplianceGroupByResourceTypeResponseBody {
	s.ComplianceResult = v
	return s
}

func (s *GetAggregateResourceComplianceGroupByResourceTypeResponseBody) SetRequestId(v string) *GetAggregateResourceComplianceGroupByResourceTypeResponseBody {
	s.RequestId = &v
	return s
}

type GetAggregateResourceComplianceGroupByResourceTypeResponseBodyComplianceResult struct {
	ComplianceResultList []*GetAggregateResourceComplianceGroupByResourceTypeResponseBodyComplianceResultComplianceResultList `json:"ComplianceResultList,omitempty" xml:"ComplianceResultList,omitempty" type:"Repeated"`
}

func (s GetAggregateResourceComplianceGroupByResourceTypeResponseBodyComplianceResult) String() string {
	return tea.Prettify(s)
}

func (s GetAggregateResourceComplianceGroupByResourceTypeResponseBodyComplianceResult) GoString() string {
	return s.String()
}

func (s *GetAggregateResourceComplianceGroupByResourceTypeResponseBodyComplianceResult) SetComplianceResultList(v []*GetAggregateResourceComplianceGroupByResourceTypeResponseBodyComplianceResultComplianceResultList) *GetAggregateResourceComplianceGroupByResourceTypeResponseBodyComplianceResult {
	s.ComplianceResultList = v
	return s
}

type GetAggregateResourceComplianceGroupByResourceTypeResponseBodyComplianceResultComplianceResultList struct {
	Compliances  []*GetAggregateResourceComplianceGroupByResourceTypeResponseBodyComplianceResultComplianceResultListCompliances `json:"Compliances,omitempty" xml:"Compliances,omitempty" type:"Repeated"`
	ResourceType *string                                                                                                         `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s GetAggregateResourceComplianceGroupByResourceTypeResponseBodyComplianceResultComplianceResultList) String() string {
	return tea.Prettify(s)
}

func (s GetAggregateResourceComplianceGroupByResourceTypeResponseBodyComplianceResultComplianceResultList) GoString() string {
	return s.String()
}

func (s *GetAggregateResourceComplianceGroupByResourceTypeResponseBodyComplianceResultComplianceResultList) SetCompliances(v []*GetAggregateResourceComplianceGroupByResourceTypeResponseBodyComplianceResultComplianceResultListCompliances) *GetAggregateResourceComplianceGroupByResourceTypeResponseBodyComplianceResultComplianceResultList {
	s.Compliances = v
	return s
}

func (s *GetAggregateResourceComplianceGroupByResourceTypeResponseBodyComplianceResultComplianceResultList) SetResourceType(v string) *GetAggregateResourceComplianceGroupByResourceTypeResponseBodyComplianceResultComplianceResultList {
	s.ResourceType = &v
	return s
}

type GetAggregateResourceComplianceGroupByResourceTypeResponseBodyComplianceResultComplianceResultListCompliances struct {
	ComplianceType *string `json:"ComplianceType,omitempty" xml:"ComplianceType,omitempty"`
	Count          *int64  `json:"Count,omitempty" xml:"Count,omitempty"`
}

func (s GetAggregateResourceComplianceGroupByResourceTypeResponseBodyComplianceResultComplianceResultListCompliances) String() string {
	return tea.Prettify(s)
}

func (s GetAggregateResourceComplianceGroupByResourceTypeResponseBodyComplianceResultComplianceResultListCompliances) GoString() string {
	return s.String()
}

func (s *GetAggregateResourceComplianceGroupByResourceTypeResponseBodyComplianceResultComplianceResultListCompliances) SetComplianceType(v string) *GetAggregateResourceComplianceGroupByResourceTypeResponseBodyComplianceResultComplianceResultListCompliances {
	s.ComplianceType = &v
	return s
}

func (s *GetAggregateResourceComplianceGroupByResourceTypeResponseBodyComplianceResultComplianceResultListCompliances) SetCount(v int64) *GetAggregateResourceComplianceGroupByResourceTypeResponseBodyComplianceResultComplianceResultListCompliances {
	s.Count = &v
	return s
}

type GetAggregateResourceComplianceGroupByResourceTypeResponse struct {
	Headers    map[string]*string                                             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetAggregateResourceComplianceGroupByResourceTypeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetAggregateResourceComplianceGroupByResourceTypeResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAggregateResourceComplianceGroupByResourceTypeResponse) GoString() string {
	return s.String()
}

func (s *GetAggregateResourceComplianceGroupByResourceTypeResponse) SetHeaders(v map[string]*string) *GetAggregateResourceComplianceGroupByResourceTypeResponse {
	s.Headers = v
	return s
}

func (s *GetAggregateResourceComplianceGroupByResourceTypeResponse) SetStatusCode(v int32) *GetAggregateResourceComplianceGroupByResourceTypeResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAggregateResourceComplianceGroupByResourceTypeResponse) SetBody(v *GetAggregateResourceComplianceGroupByResourceTypeResponseBody) *GetAggregateResourceComplianceGroupByResourceTypeResponse {
	s.Body = v
	return s
}

type GetAggregateResourceComplianceTimelineRequest struct {
	AggregatorId    *string `json:"AggregatorId,omitempty" xml:"AggregatorId,omitempty"`
	EndTime         *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	MaxResults      *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken       *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	Region          *string `json:"Region,omitempty" xml:"Region,omitempty"`
	ResourceId      *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ResourceType    *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	StartTime       *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s GetAggregateResourceComplianceTimelineRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAggregateResourceComplianceTimelineRequest) GoString() string {
	return s.String()
}

func (s *GetAggregateResourceComplianceTimelineRequest) SetAggregatorId(v string) *GetAggregateResourceComplianceTimelineRequest {
	s.AggregatorId = &v
	return s
}

func (s *GetAggregateResourceComplianceTimelineRequest) SetEndTime(v int64) *GetAggregateResourceComplianceTimelineRequest {
	s.EndTime = &v
	return s
}

func (s *GetAggregateResourceComplianceTimelineRequest) SetMaxResults(v int32) *GetAggregateResourceComplianceTimelineRequest {
	s.MaxResults = &v
	return s
}

func (s *GetAggregateResourceComplianceTimelineRequest) SetNextToken(v string) *GetAggregateResourceComplianceTimelineRequest {
	s.NextToken = &v
	return s
}

func (s *GetAggregateResourceComplianceTimelineRequest) SetRegion(v string) *GetAggregateResourceComplianceTimelineRequest {
	s.Region = &v
	return s
}

func (s *GetAggregateResourceComplianceTimelineRequest) SetResourceId(v string) *GetAggregateResourceComplianceTimelineRequest {
	s.ResourceId = &v
	return s
}

func (s *GetAggregateResourceComplianceTimelineRequest) SetResourceOwnerId(v int64) *GetAggregateResourceComplianceTimelineRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetAggregateResourceComplianceTimelineRequest) SetResourceType(v string) *GetAggregateResourceComplianceTimelineRequest {
	s.ResourceType = &v
	return s
}

func (s *GetAggregateResourceComplianceTimelineRequest) SetStartTime(v int64) *GetAggregateResourceComplianceTimelineRequest {
	s.StartTime = &v
	return s
}

type GetAggregateResourceComplianceTimelineResponseBody struct {
	RequestId                  *string                                                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResourceComplianceTimeline *GetAggregateResourceComplianceTimelineResponseBodyResourceComplianceTimeline `json:"ResourceComplianceTimeline,omitempty" xml:"ResourceComplianceTimeline,omitempty" type:"Struct"`
}

func (s GetAggregateResourceComplianceTimelineResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAggregateResourceComplianceTimelineResponseBody) GoString() string {
	return s.String()
}

func (s *GetAggregateResourceComplianceTimelineResponseBody) SetRequestId(v string) *GetAggregateResourceComplianceTimelineResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAggregateResourceComplianceTimelineResponseBody) SetResourceComplianceTimeline(v *GetAggregateResourceComplianceTimelineResponseBodyResourceComplianceTimeline) *GetAggregateResourceComplianceTimelineResponseBody {
	s.ResourceComplianceTimeline = v
	return s
}

type GetAggregateResourceComplianceTimelineResponseBodyResourceComplianceTimeline struct {
	ComplianceList []*GetAggregateResourceComplianceTimelineResponseBodyResourceComplianceTimelineComplianceList `json:"ComplianceList,omitempty" xml:"ComplianceList,omitempty" type:"Repeated"`
	MaxResults     *int32                                                                                        `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken      *string                                                                                       `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s GetAggregateResourceComplianceTimelineResponseBodyResourceComplianceTimeline) String() string {
	return tea.Prettify(s)
}

func (s GetAggregateResourceComplianceTimelineResponseBodyResourceComplianceTimeline) GoString() string {
	return s.String()
}

func (s *GetAggregateResourceComplianceTimelineResponseBodyResourceComplianceTimeline) SetComplianceList(v []*GetAggregateResourceComplianceTimelineResponseBodyResourceComplianceTimelineComplianceList) *GetAggregateResourceComplianceTimelineResponseBodyResourceComplianceTimeline {
	s.ComplianceList = v
	return s
}

func (s *GetAggregateResourceComplianceTimelineResponseBodyResourceComplianceTimeline) SetMaxResults(v int32) *GetAggregateResourceComplianceTimelineResponseBodyResourceComplianceTimeline {
	s.MaxResults = &v
	return s
}

func (s *GetAggregateResourceComplianceTimelineResponseBodyResourceComplianceTimeline) SetNextToken(v string) *GetAggregateResourceComplianceTimelineResponseBodyResourceComplianceTimeline {
	s.NextToken = &v
	return s
}

type GetAggregateResourceComplianceTimelineResponseBodyResourceComplianceTimelineComplianceList struct {
	AccountId          *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	AvailabilityZone   *string `json:"AvailabilityZone,omitempty" xml:"AvailabilityZone,omitempty"`
	CaptureTime        *int64  `json:"CaptureTime,omitempty" xml:"CaptureTime,omitempty"`
	Configuration      *string `json:"Configuration,omitempty" xml:"Configuration,omitempty"`
	ConfigurationDiff  *string `json:"ConfigurationDiff,omitempty" xml:"ConfigurationDiff,omitempty"`
	Region             *string `json:"Region,omitempty" xml:"Region,omitempty"`
	ResourceCreateTime *int64  `json:"ResourceCreateTime,omitempty" xml:"ResourceCreateTime,omitempty"`
	ResourceId         *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ResourceName       *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	ResourceStatus     *string `json:"ResourceStatus,omitempty" xml:"ResourceStatus,omitempty"`
	ResourceType       *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	Tags               *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
}

func (s GetAggregateResourceComplianceTimelineResponseBodyResourceComplianceTimelineComplianceList) String() string {
	return tea.Prettify(s)
}

func (s GetAggregateResourceComplianceTimelineResponseBodyResourceComplianceTimelineComplianceList) GoString() string {
	return s.String()
}

func (s *GetAggregateResourceComplianceTimelineResponseBodyResourceComplianceTimelineComplianceList) SetAccountId(v string) *GetAggregateResourceComplianceTimelineResponseBodyResourceComplianceTimelineComplianceList {
	s.AccountId = &v
	return s
}

func (s *GetAggregateResourceComplianceTimelineResponseBodyResourceComplianceTimelineComplianceList) SetAvailabilityZone(v string) *GetAggregateResourceComplianceTimelineResponseBodyResourceComplianceTimelineComplianceList {
	s.AvailabilityZone = &v
	return s
}

func (s *GetAggregateResourceComplianceTimelineResponseBodyResourceComplianceTimelineComplianceList) SetCaptureTime(v int64) *GetAggregateResourceComplianceTimelineResponseBodyResourceComplianceTimelineComplianceList {
	s.CaptureTime = &v
	return s
}

func (s *GetAggregateResourceComplianceTimelineResponseBodyResourceComplianceTimelineComplianceList) SetConfiguration(v string) *GetAggregateResourceComplianceTimelineResponseBodyResourceComplianceTimelineComplianceList {
	s.Configuration = &v
	return s
}

func (s *GetAggregateResourceComplianceTimelineResponseBodyResourceComplianceTimelineComplianceList) SetConfigurationDiff(v string) *GetAggregateResourceComplianceTimelineResponseBodyResourceComplianceTimelineComplianceList {
	s.ConfigurationDiff = &v
	return s
}

func (s *GetAggregateResourceComplianceTimelineResponseBodyResourceComplianceTimelineComplianceList) SetRegion(v string) *GetAggregateResourceComplianceTimelineResponseBodyResourceComplianceTimelineComplianceList {
	s.Region = &v
	return s
}

func (s *GetAggregateResourceComplianceTimelineResponseBodyResourceComplianceTimelineComplianceList) SetResourceCreateTime(v int64) *GetAggregateResourceComplianceTimelineResponseBodyResourceComplianceTimelineComplianceList {
	s.ResourceCreateTime = &v
	return s
}

func (s *GetAggregateResourceComplianceTimelineResponseBodyResourceComplianceTimelineComplianceList) SetResourceId(v string) *GetAggregateResourceComplianceTimelineResponseBodyResourceComplianceTimelineComplianceList {
	s.ResourceId = &v
	return s
}

func (s *GetAggregateResourceComplianceTimelineResponseBodyResourceComplianceTimelineComplianceList) SetResourceName(v string) *GetAggregateResourceComplianceTimelineResponseBodyResourceComplianceTimelineComplianceList {
	s.ResourceName = &v
	return s
}

func (s *GetAggregateResourceComplianceTimelineResponseBodyResourceComplianceTimelineComplianceList) SetResourceStatus(v string) *GetAggregateResourceComplianceTimelineResponseBodyResourceComplianceTimelineComplianceList {
	s.ResourceStatus = &v
	return s
}

func (s *GetAggregateResourceComplianceTimelineResponseBodyResourceComplianceTimelineComplianceList) SetResourceType(v string) *GetAggregateResourceComplianceTimelineResponseBodyResourceComplianceTimelineComplianceList {
	s.ResourceType = &v
	return s
}

func (s *GetAggregateResourceComplianceTimelineResponseBodyResourceComplianceTimelineComplianceList) SetTags(v string) *GetAggregateResourceComplianceTimelineResponseBodyResourceComplianceTimelineComplianceList {
	s.Tags = &v
	return s
}

type GetAggregateResourceComplianceTimelineResponse struct {
	Headers    map[string]*string                                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetAggregateResourceComplianceTimelineResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetAggregateResourceComplianceTimelineResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAggregateResourceComplianceTimelineResponse) GoString() string {
	return s.String()
}

func (s *GetAggregateResourceComplianceTimelineResponse) SetHeaders(v map[string]*string) *GetAggregateResourceComplianceTimelineResponse {
	s.Headers = v
	return s
}

func (s *GetAggregateResourceComplianceTimelineResponse) SetStatusCode(v int32) *GetAggregateResourceComplianceTimelineResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAggregateResourceComplianceTimelineResponse) SetBody(v *GetAggregateResourceComplianceTimelineResponseBody) *GetAggregateResourceComplianceTimelineResponse {
	s.Body = v
	return s
}

type GetAggregateResourceConfigurationTimelineRequest struct {
	AggregatorId    *string `json:"AggregatorId,omitempty" xml:"AggregatorId,omitempty"`
	EndTime         *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	MaxResults      *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken       *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	Region          *string `json:"Region,omitempty" xml:"Region,omitempty"`
	ResourceId      *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ResourceType    *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	StartTime       *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s GetAggregateResourceConfigurationTimelineRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAggregateResourceConfigurationTimelineRequest) GoString() string {
	return s.String()
}

func (s *GetAggregateResourceConfigurationTimelineRequest) SetAggregatorId(v string) *GetAggregateResourceConfigurationTimelineRequest {
	s.AggregatorId = &v
	return s
}

func (s *GetAggregateResourceConfigurationTimelineRequest) SetEndTime(v int64) *GetAggregateResourceConfigurationTimelineRequest {
	s.EndTime = &v
	return s
}

func (s *GetAggregateResourceConfigurationTimelineRequest) SetMaxResults(v int32) *GetAggregateResourceConfigurationTimelineRequest {
	s.MaxResults = &v
	return s
}

func (s *GetAggregateResourceConfigurationTimelineRequest) SetNextToken(v string) *GetAggregateResourceConfigurationTimelineRequest {
	s.NextToken = &v
	return s
}

func (s *GetAggregateResourceConfigurationTimelineRequest) SetRegion(v string) *GetAggregateResourceConfigurationTimelineRequest {
	s.Region = &v
	return s
}

func (s *GetAggregateResourceConfigurationTimelineRequest) SetResourceId(v string) *GetAggregateResourceConfigurationTimelineRequest {
	s.ResourceId = &v
	return s
}

func (s *GetAggregateResourceConfigurationTimelineRequest) SetResourceOwnerId(v int64) *GetAggregateResourceConfigurationTimelineRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetAggregateResourceConfigurationTimelineRequest) SetResourceType(v string) *GetAggregateResourceConfigurationTimelineRequest {
	s.ResourceType = &v
	return s
}

func (s *GetAggregateResourceConfigurationTimelineRequest) SetStartTime(v int64) *GetAggregateResourceConfigurationTimelineRequest {
	s.StartTime = &v
	return s
}

type GetAggregateResourceConfigurationTimelineResponseBody struct {
	RequestId                     *string                                                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResourceConfigurationTimeline *GetAggregateResourceConfigurationTimelineResponseBodyResourceConfigurationTimeline `json:"ResourceConfigurationTimeline,omitempty" xml:"ResourceConfigurationTimeline,omitempty" type:"Struct"`
}

func (s GetAggregateResourceConfigurationTimelineResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAggregateResourceConfigurationTimelineResponseBody) GoString() string {
	return s.String()
}

func (s *GetAggregateResourceConfigurationTimelineResponseBody) SetRequestId(v string) *GetAggregateResourceConfigurationTimelineResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAggregateResourceConfigurationTimelineResponseBody) SetResourceConfigurationTimeline(v *GetAggregateResourceConfigurationTimelineResponseBodyResourceConfigurationTimeline) *GetAggregateResourceConfigurationTimelineResponseBody {
	s.ResourceConfigurationTimeline = v
	return s
}

type GetAggregateResourceConfigurationTimelineResponseBodyResourceConfigurationTimeline struct {
	ConfigurationList []*GetAggregateResourceConfigurationTimelineResponseBodyResourceConfigurationTimelineConfigurationList `json:"ConfigurationList,omitempty" xml:"ConfigurationList,omitempty" type:"Repeated"`
	MaxResults        *int32                                                                                                 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken         *string                                                                                                `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s GetAggregateResourceConfigurationTimelineResponseBodyResourceConfigurationTimeline) String() string {
	return tea.Prettify(s)
}

func (s GetAggregateResourceConfigurationTimelineResponseBodyResourceConfigurationTimeline) GoString() string {
	return s.String()
}

func (s *GetAggregateResourceConfigurationTimelineResponseBodyResourceConfigurationTimeline) SetConfigurationList(v []*GetAggregateResourceConfigurationTimelineResponseBodyResourceConfigurationTimelineConfigurationList) *GetAggregateResourceConfigurationTimelineResponseBodyResourceConfigurationTimeline {
	s.ConfigurationList = v
	return s
}

func (s *GetAggregateResourceConfigurationTimelineResponseBodyResourceConfigurationTimeline) SetMaxResults(v int32) *GetAggregateResourceConfigurationTimelineResponseBodyResourceConfigurationTimeline {
	s.MaxResults = &v
	return s
}

func (s *GetAggregateResourceConfigurationTimelineResponseBodyResourceConfigurationTimeline) SetNextToken(v string) *GetAggregateResourceConfigurationTimelineResponseBodyResourceConfigurationTimeline {
	s.NextToken = &v
	return s
}

type GetAggregateResourceConfigurationTimelineResponseBodyResourceConfigurationTimelineConfigurationList struct {
	AccountId          *int64  `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	AvailabilityZone   *string `json:"AvailabilityZone,omitempty" xml:"AvailabilityZone,omitempty"`
	CaptureTime        *string `json:"CaptureTime,omitempty" xml:"CaptureTime,omitempty"`
	ConfigurationDiff  *string `json:"ConfigurationDiff,omitempty" xml:"ConfigurationDiff,omitempty"`
	Region             *string `json:"Region,omitempty" xml:"Region,omitempty"`
	ResourceCreateTime *string `json:"ResourceCreateTime,omitempty" xml:"ResourceCreateTime,omitempty"`
	ResourceEventType  *string `json:"ResourceEventType,omitempty" xml:"ResourceEventType,omitempty"`
	ResourceId         *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ResourceName       *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	ResourceType       *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	Tags               *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
}

func (s GetAggregateResourceConfigurationTimelineResponseBodyResourceConfigurationTimelineConfigurationList) String() string {
	return tea.Prettify(s)
}

func (s GetAggregateResourceConfigurationTimelineResponseBodyResourceConfigurationTimelineConfigurationList) GoString() string {
	return s.String()
}

func (s *GetAggregateResourceConfigurationTimelineResponseBodyResourceConfigurationTimelineConfigurationList) SetAccountId(v int64) *GetAggregateResourceConfigurationTimelineResponseBodyResourceConfigurationTimelineConfigurationList {
	s.AccountId = &v
	return s
}

func (s *GetAggregateResourceConfigurationTimelineResponseBodyResourceConfigurationTimelineConfigurationList) SetAvailabilityZone(v string) *GetAggregateResourceConfigurationTimelineResponseBodyResourceConfigurationTimelineConfigurationList {
	s.AvailabilityZone = &v
	return s
}

func (s *GetAggregateResourceConfigurationTimelineResponseBodyResourceConfigurationTimelineConfigurationList) SetCaptureTime(v string) *GetAggregateResourceConfigurationTimelineResponseBodyResourceConfigurationTimelineConfigurationList {
	s.CaptureTime = &v
	return s
}

func (s *GetAggregateResourceConfigurationTimelineResponseBodyResourceConfigurationTimelineConfigurationList) SetConfigurationDiff(v string) *GetAggregateResourceConfigurationTimelineResponseBodyResourceConfigurationTimelineConfigurationList {
	s.ConfigurationDiff = &v
	return s
}

func (s *GetAggregateResourceConfigurationTimelineResponseBodyResourceConfigurationTimelineConfigurationList) SetRegion(v string) *GetAggregateResourceConfigurationTimelineResponseBodyResourceConfigurationTimelineConfigurationList {
	s.Region = &v
	return s
}

func (s *GetAggregateResourceConfigurationTimelineResponseBodyResourceConfigurationTimelineConfigurationList) SetResourceCreateTime(v string) *GetAggregateResourceConfigurationTimelineResponseBodyResourceConfigurationTimelineConfigurationList {
	s.ResourceCreateTime = &v
	return s
}

func (s *GetAggregateResourceConfigurationTimelineResponseBodyResourceConfigurationTimelineConfigurationList) SetResourceEventType(v string) *GetAggregateResourceConfigurationTimelineResponseBodyResourceConfigurationTimelineConfigurationList {
	s.ResourceEventType = &v
	return s
}

func (s *GetAggregateResourceConfigurationTimelineResponseBodyResourceConfigurationTimelineConfigurationList) SetResourceId(v string) *GetAggregateResourceConfigurationTimelineResponseBodyResourceConfigurationTimelineConfigurationList {
	s.ResourceId = &v
	return s
}

func (s *GetAggregateResourceConfigurationTimelineResponseBodyResourceConfigurationTimelineConfigurationList) SetResourceName(v string) *GetAggregateResourceConfigurationTimelineResponseBodyResourceConfigurationTimelineConfigurationList {
	s.ResourceName = &v
	return s
}

func (s *GetAggregateResourceConfigurationTimelineResponseBodyResourceConfigurationTimelineConfigurationList) SetResourceType(v string) *GetAggregateResourceConfigurationTimelineResponseBodyResourceConfigurationTimelineConfigurationList {
	s.ResourceType = &v
	return s
}

func (s *GetAggregateResourceConfigurationTimelineResponseBodyResourceConfigurationTimelineConfigurationList) SetTags(v string) *GetAggregateResourceConfigurationTimelineResponseBodyResourceConfigurationTimelineConfigurationList {
	s.Tags = &v
	return s
}

type GetAggregateResourceConfigurationTimelineResponse struct {
	Headers    map[string]*string                                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetAggregateResourceConfigurationTimelineResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetAggregateResourceConfigurationTimelineResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAggregateResourceConfigurationTimelineResponse) GoString() string {
	return s.String()
}

func (s *GetAggregateResourceConfigurationTimelineResponse) SetHeaders(v map[string]*string) *GetAggregateResourceConfigurationTimelineResponse {
	s.Headers = v
	return s
}

func (s *GetAggregateResourceConfigurationTimelineResponse) SetStatusCode(v int32) *GetAggregateResourceConfigurationTimelineResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAggregateResourceConfigurationTimelineResponse) SetBody(v *GetAggregateResourceConfigurationTimelineResponseBody) *GetAggregateResourceConfigurationTimelineResponse {
	s.Body = v
	return s
}

type GetAggregateResourceCountsGroupByRegionRequest struct {
	AggregatorId    *string `json:"AggregatorId,omitempty" xml:"AggregatorId,omitempty"`
	FolderId        *string `json:"FolderId,omitempty" xml:"FolderId,omitempty"`
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ResourceType    *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s GetAggregateResourceCountsGroupByRegionRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAggregateResourceCountsGroupByRegionRequest) GoString() string {
	return s.String()
}

func (s *GetAggregateResourceCountsGroupByRegionRequest) SetAggregatorId(v string) *GetAggregateResourceCountsGroupByRegionRequest {
	s.AggregatorId = &v
	return s
}

func (s *GetAggregateResourceCountsGroupByRegionRequest) SetFolderId(v string) *GetAggregateResourceCountsGroupByRegionRequest {
	s.FolderId = &v
	return s
}

func (s *GetAggregateResourceCountsGroupByRegionRequest) SetResourceOwnerId(v int64) *GetAggregateResourceCountsGroupByRegionRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetAggregateResourceCountsGroupByRegionRequest) SetResourceType(v string) *GetAggregateResourceCountsGroupByRegionRequest {
	s.ResourceType = &v
	return s
}

type GetAggregateResourceCountsGroupByRegionResponseBody struct {
	DiscoveredResourceCountsSummary []*GetAggregateResourceCountsGroupByRegionResponseBodyDiscoveredResourceCountsSummary `json:"DiscoveredResourceCountsSummary,omitempty" xml:"DiscoveredResourceCountsSummary,omitempty" type:"Repeated"`
	RequestId                       *string                                                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetAggregateResourceCountsGroupByRegionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAggregateResourceCountsGroupByRegionResponseBody) GoString() string {
	return s.String()
}

func (s *GetAggregateResourceCountsGroupByRegionResponseBody) SetDiscoveredResourceCountsSummary(v []*GetAggregateResourceCountsGroupByRegionResponseBodyDiscoveredResourceCountsSummary) *GetAggregateResourceCountsGroupByRegionResponseBody {
	s.DiscoveredResourceCountsSummary = v
	return s
}

func (s *GetAggregateResourceCountsGroupByRegionResponseBody) SetRequestId(v string) *GetAggregateResourceCountsGroupByRegionResponseBody {
	s.RequestId = &v
	return s
}

type GetAggregateResourceCountsGroupByRegionResponseBodyDiscoveredResourceCountsSummary struct {
	Region        *string `json:"Region,omitempty" xml:"Region,omitempty"`
	ResourceCount *int64  `json:"ResourceCount,omitempty" xml:"ResourceCount,omitempty"`
}

func (s GetAggregateResourceCountsGroupByRegionResponseBodyDiscoveredResourceCountsSummary) String() string {
	return tea.Prettify(s)
}

func (s GetAggregateResourceCountsGroupByRegionResponseBodyDiscoveredResourceCountsSummary) GoString() string {
	return s.String()
}

func (s *GetAggregateResourceCountsGroupByRegionResponseBodyDiscoveredResourceCountsSummary) SetRegion(v string) *GetAggregateResourceCountsGroupByRegionResponseBodyDiscoveredResourceCountsSummary {
	s.Region = &v
	return s
}

func (s *GetAggregateResourceCountsGroupByRegionResponseBodyDiscoveredResourceCountsSummary) SetResourceCount(v int64) *GetAggregateResourceCountsGroupByRegionResponseBodyDiscoveredResourceCountsSummary {
	s.ResourceCount = &v
	return s
}

type GetAggregateResourceCountsGroupByRegionResponse struct {
	Headers    map[string]*string                                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetAggregateResourceCountsGroupByRegionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetAggregateResourceCountsGroupByRegionResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAggregateResourceCountsGroupByRegionResponse) GoString() string {
	return s.String()
}

func (s *GetAggregateResourceCountsGroupByRegionResponse) SetHeaders(v map[string]*string) *GetAggregateResourceCountsGroupByRegionResponse {
	s.Headers = v
	return s
}

func (s *GetAggregateResourceCountsGroupByRegionResponse) SetStatusCode(v int32) *GetAggregateResourceCountsGroupByRegionResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAggregateResourceCountsGroupByRegionResponse) SetBody(v *GetAggregateResourceCountsGroupByRegionResponseBody) *GetAggregateResourceCountsGroupByRegionResponse {
	s.Body = v
	return s
}

type GetAggregateResourceCountsGroupByResourceTypeRequest struct {
	AggregatorId    *string `json:"AggregatorId,omitempty" xml:"AggregatorId,omitempty"`
	FolderId        *string `json:"FolderId,omitempty" xml:"FolderId,omitempty"`
	Region          *string `json:"Region,omitempty" xml:"Region,omitempty"`
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s GetAggregateResourceCountsGroupByResourceTypeRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAggregateResourceCountsGroupByResourceTypeRequest) GoString() string {
	return s.String()
}

func (s *GetAggregateResourceCountsGroupByResourceTypeRequest) SetAggregatorId(v string) *GetAggregateResourceCountsGroupByResourceTypeRequest {
	s.AggregatorId = &v
	return s
}

func (s *GetAggregateResourceCountsGroupByResourceTypeRequest) SetFolderId(v string) *GetAggregateResourceCountsGroupByResourceTypeRequest {
	s.FolderId = &v
	return s
}

func (s *GetAggregateResourceCountsGroupByResourceTypeRequest) SetRegion(v string) *GetAggregateResourceCountsGroupByResourceTypeRequest {
	s.Region = &v
	return s
}

func (s *GetAggregateResourceCountsGroupByResourceTypeRequest) SetResourceOwnerId(v int64) *GetAggregateResourceCountsGroupByResourceTypeRequest {
	s.ResourceOwnerId = &v
	return s
}

type GetAggregateResourceCountsGroupByResourceTypeResponseBody struct {
	DiscoveredResourceCountsSummary []*GetAggregateResourceCountsGroupByResourceTypeResponseBodyDiscoveredResourceCountsSummary `json:"DiscoveredResourceCountsSummary,omitempty" xml:"DiscoveredResourceCountsSummary,omitempty" type:"Repeated"`
	RequestId                       *string                                                                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetAggregateResourceCountsGroupByResourceTypeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAggregateResourceCountsGroupByResourceTypeResponseBody) GoString() string {
	return s.String()
}

func (s *GetAggregateResourceCountsGroupByResourceTypeResponseBody) SetDiscoveredResourceCountsSummary(v []*GetAggregateResourceCountsGroupByResourceTypeResponseBodyDiscoveredResourceCountsSummary) *GetAggregateResourceCountsGroupByResourceTypeResponseBody {
	s.DiscoveredResourceCountsSummary = v
	return s
}

func (s *GetAggregateResourceCountsGroupByResourceTypeResponseBody) SetRequestId(v string) *GetAggregateResourceCountsGroupByResourceTypeResponseBody {
	s.RequestId = &v
	return s
}

type GetAggregateResourceCountsGroupByResourceTypeResponseBodyDiscoveredResourceCountsSummary struct {
	ResourceCount *int64  `json:"ResourceCount,omitempty" xml:"ResourceCount,omitempty"`
	ResourceType  *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s GetAggregateResourceCountsGroupByResourceTypeResponseBodyDiscoveredResourceCountsSummary) String() string {
	return tea.Prettify(s)
}

func (s GetAggregateResourceCountsGroupByResourceTypeResponseBodyDiscoveredResourceCountsSummary) GoString() string {
	return s.String()
}

func (s *GetAggregateResourceCountsGroupByResourceTypeResponseBodyDiscoveredResourceCountsSummary) SetResourceCount(v int64) *GetAggregateResourceCountsGroupByResourceTypeResponseBodyDiscoveredResourceCountsSummary {
	s.ResourceCount = &v
	return s
}

func (s *GetAggregateResourceCountsGroupByResourceTypeResponseBodyDiscoveredResourceCountsSummary) SetResourceType(v string) *GetAggregateResourceCountsGroupByResourceTypeResponseBodyDiscoveredResourceCountsSummary {
	s.ResourceType = &v
	return s
}

type GetAggregateResourceCountsGroupByResourceTypeResponse struct {
	Headers    map[string]*string                                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetAggregateResourceCountsGroupByResourceTypeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetAggregateResourceCountsGroupByResourceTypeResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAggregateResourceCountsGroupByResourceTypeResponse) GoString() string {
	return s.String()
}

func (s *GetAggregateResourceCountsGroupByResourceTypeResponse) SetHeaders(v map[string]*string) *GetAggregateResourceCountsGroupByResourceTypeResponse {
	s.Headers = v
	return s
}

func (s *GetAggregateResourceCountsGroupByResourceTypeResponse) SetStatusCode(v int32) *GetAggregateResourceCountsGroupByResourceTypeResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAggregateResourceCountsGroupByResourceTypeResponse) SetBody(v *GetAggregateResourceCountsGroupByResourceTypeResponseBody) *GetAggregateResourceCountsGroupByResourceTypeResponse {
	s.Body = v
	return s
}

type GetAggregatorRequest struct {
	AggregatorId *string `json:"AggregatorId,omitempty" xml:"AggregatorId,omitempty"`
}

func (s GetAggregatorRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAggregatorRequest) GoString() string {
	return s.String()
}

func (s *GetAggregatorRequest) SetAggregatorId(v string) *GetAggregatorRequest {
	s.AggregatorId = &v
	return s
}

type GetAggregatorResponseBody struct {
	Aggregator *GetAggregatorResponseBodyAggregator `json:"Aggregator,omitempty" xml:"Aggregator,omitempty" type:"Struct"`
	RequestId  *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetAggregatorResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAggregatorResponseBody) GoString() string {
	return s.String()
}

func (s *GetAggregatorResponseBody) SetAggregator(v *GetAggregatorResponseBodyAggregator) *GetAggregatorResponseBody {
	s.Aggregator = v
	return s
}

func (s *GetAggregatorResponseBody) SetRequestId(v string) *GetAggregatorResponseBody {
	s.RequestId = &v
	return s
}

type GetAggregatorResponseBodyAggregator struct {
	AccountId                 *int64                                                   `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	AggregatorAccountCount    *int64                                                   `json:"AggregatorAccountCount,omitempty" xml:"AggregatorAccountCount,omitempty"`
	AggregatorAccounts        []*GetAggregatorResponseBodyAggregatorAggregatorAccounts `json:"AggregatorAccounts,omitempty" xml:"AggregatorAccounts,omitempty" type:"Repeated"`
	AggregatorCreateTimestamp *string                                                  `json:"AggregatorCreateTimestamp,omitempty" xml:"AggregatorCreateTimestamp,omitempty"`
	AggregatorId              *string                                                  `json:"AggregatorId,omitempty" xml:"AggregatorId,omitempty"`
	AggregatorName            *string                                                  `json:"AggregatorName,omitempty" xml:"AggregatorName,omitempty"`
	AggregatorStatus          *int32                                                   `json:"AggregatorStatus,omitempty" xml:"AggregatorStatus,omitempty"`
	AggregatorType            *string                                                  `json:"AggregatorType,omitempty" xml:"AggregatorType,omitempty"`
	Description               *string                                                  `json:"Description,omitempty" xml:"Description,omitempty"`
}

func (s GetAggregatorResponseBodyAggregator) String() string {
	return tea.Prettify(s)
}

func (s GetAggregatorResponseBodyAggregator) GoString() string {
	return s.String()
}

func (s *GetAggregatorResponseBodyAggregator) SetAccountId(v int64) *GetAggregatorResponseBodyAggregator {
	s.AccountId = &v
	return s
}

func (s *GetAggregatorResponseBodyAggregator) SetAggregatorAccountCount(v int64) *GetAggregatorResponseBodyAggregator {
	s.AggregatorAccountCount = &v
	return s
}

func (s *GetAggregatorResponseBodyAggregator) SetAggregatorAccounts(v []*GetAggregatorResponseBodyAggregatorAggregatorAccounts) *GetAggregatorResponseBodyAggregator {
	s.AggregatorAccounts = v
	return s
}

func (s *GetAggregatorResponseBodyAggregator) SetAggregatorCreateTimestamp(v string) *GetAggregatorResponseBodyAggregator {
	s.AggregatorCreateTimestamp = &v
	return s
}

func (s *GetAggregatorResponseBodyAggregator) SetAggregatorId(v string) *GetAggregatorResponseBodyAggregator {
	s.AggregatorId = &v
	return s
}

func (s *GetAggregatorResponseBodyAggregator) SetAggregatorName(v string) *GetAggregatorResponseBodyAggregator {
	s.AggregatorName = &v
	return s
}

func (s *GetAggregatorResponseBodyAggregator) SetAggregatorStatus(v int32) *GetAggregatorResponseBodyAggregator {
	s.AggregatorStatus = &v
	return s
}

func (s *GetAggregatorResponseBodyAggregator) SetAggregatorType(v string) *GetAggregatorResponseBodyAggregator {
	s.AggregatorType = &v
	return s
}

func (s *GetAggregatorResponseBodyAggregator) SetDescription(v string) *GetAggregatorResponseBodyAggregator {
	s.Description = &v
	return s
}

type GetAggregatorResponseBodyAggregatorAggregatorAccounts struct {
	AccountId      *int64  `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	AccountName    *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	AccountType    *string `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
	RecorderStatus *string `json:"RecorderStatus,omitempty" xml:"RecorderStatus,omitempty"`
}

func (s GetAggregatorResponseBodyAggregatorAggregatorAccounts) String() string {
	return tea.Prettify(s)
}

func (s GetAggregatorResponseBodyAggregatorAggregatorAccounts) GoString() string {
	return s.String()
}

func (s *GetAggregatorResponseBodyAggregatorAggregatorAccounts) SetAccountId(v int64) *GetAggregatorResponseBodyAggregatorAggregatorAccounts {
	s.AccountId = &v
	return s
}

func (s *GetAggregatorResponseBodyAggregatorAggregatorAccounts) SetAccountName(v string) *GetAggregatorResponseBodyAggregatorAggregatorAccounts {
	s.AccountName = &v
	return s
}

func (s *GetAggregatorResponseBodyAggregatorAggregatorAccounts) SetAccountType(v string) *GetAggregatorResponseBodyAggregatorAggregatorAccounts {
	s.AccountType = &v
	return s
}

func (s *GetAggregatorResponseBodyAggregatorAggregatorAccounts) SetRecorderStatus(v string) *GetAggregatorResponseBodyAggregatorAggregatorAccounts {
	s.RecorderStatus = &v
	return s
}

type GetAggregatorResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetAggregatorResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetAggregatorResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAggregatorResponse) GoString() string {
	return s.String()
}

func (s *GetAggregatorResponse) SetHeaders(v map[string]*string) *GetAggregatorResponse {
	s.Headers = v
	return s
}

func (s *GetAggregatorResponse) SetStatusCode(v int32) *GetAggregatorResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAggregatorResponse) SetBody(v *GetAggregatorResponseBody) *GetAggregatorResponse {
	s.Body = v
	return s
}

type GetCompliancePackRequest struct {
	CompliancePackId *string `json:"CompliancePackId,omitempty" xml:"CompliancePackId,omitempty"`
}

func (s GetCompliancePackRequest) String() string {
	return tea.Prettify(s)
}

func (s GetCompliancePackRequest) GoString() string {
	return s.String()
}

func (s *GetCompliancePackRequest) SetCompliancePackId(v string) *GetCompliancePackRequest {
	s.CompliancePackId = &v
	return s
}

type GetCompliancePackResponseBody struct {
	CompliancePack *GetCompliancePackResponseBodyCompliancePack `json:"CompliancePack,omitempty" xml:"CompliancePack,omitempty" type:"Struct"`
	RequestId      *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetCompliancePackResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetCompliancePackResponseBody) GoString() string {
	return s.String()
}

func (s *GetCompliancePackResponseBody) SetCompliancePack(v *GetCompliancePackResponseBodyCompliancePack) *GetCompliancePackResponseBody {
	s.CompliancePack = v
	return s
}

func (s *GetCompliancePackResponseBody) SetRequestId(v string) *GetCompliancePackResponseBody {
	s.RequestId = &v
	return s
}

type GetCompliancePackResponseBodyCompliancePack struct {
	AccountId                *int64                                                    `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	CompliancePackId         *string                                                   `json:"CompliancePackId,omitempty" xml:"CompliancePackId,omitempty"`
	CompliancePackName       *string                                                   `json:"CompliancePackName,omitempty" xml:"CompliancePackName,omitempty"`
	CompliancePackTemplateId *string                                                   `json:"CompliancePackTemplateId,omitempty" xml:"CompliancePackTemplateId,omitempty"`
	ConfigRules              []*GetCompliancePackResponseBodyCompliancePackConfigRules `json:"ConfigRules,omitempty" xml:"ConfigRules,omitempty" type:"Repeated"`
	CreateTimestamp          *int64                                                    `json:"CreateTimestamp,omitempty" xml:"CreateTimestamp,omitempty"`
	Description              *string                                                   `json:"Description,omitempty" xml:"Description,omitempty"`
	RiskLevel                *int32                                                    `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
	Status                   *string                                                   `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetCompliancePackResponseBodyCompliancePack) String() string {
	return tea.Prettify(s)
}

func (s GetCompliancePackResponseBodyCompliancePack) GoString() string {
	return s.String()
}

func (s *GetCompliancePackResponseBodyCompliancePack) SetAccountId(v int64) *GetCompliancePackResponseBodyCompliancePack {
	s.AccountId = &v
	return s
}

func (s *GetCompliancePackResponseBodyCompliancePack) SetCompliancePackId(v string) *GetCompliancePackResponseBodyCompliancePack {
	s.CompliancePackId = &v
	return s
}

func (s *GetCompliancePackResponseBodyCompliancePack) SetCompliancePackName(v string) *GetCompliancePackResponseBodyCompliancePack {
	s.CompliancePackName = &v
	return s
}

func (s *GetCompliancePackResponseBodyCompliancePack) SetCompliancePackTemplateId(v string) *GetCompliancePackResponseBodyCompliancePack {
	s.CompliancePackTemplateId = &v
	return s
}

func (s *GetCompliancePackResponseBodyCompliancePack) SetConfigRules(v []*GetCompliancePackResponseBodyCompliancePackConfigRules) *GetCompliancePackResponseBodyCompliancePack {
	s.ConfigRules = v
	return s
}

func (s *GetCompliancePackResponseBodyCompliancePack) SetCreateTimestamp(v int64) *GetCompliancePackResponseBodyCompliancePack {
	s.CreateTimestamp = &v
	return s
}

func (s *GetCompliancePackResponseBodyCompliancePack) SetDescription(v string) *GetCompliancePackResponseBodyCompliancePack {
	s.Description = &v
	return s
}

func (s *GetCompliancePackResponseBodyCompliancePack) SetRiskLevel(v int32) *GetCompliancePackResponseBodyCompliancePack {
	s.RiskLevel = &v
	return s
}

func (s *GetCompliancePackResponseBodyCompliancePack) SetStatus(v string) *GetCompliancePackResponseBodyCompliancePack {
	s.Status = &v
	return s
}

type GetCompliancePackResponseBodyCompliancePackConfigRules struct {
	ConfigRuleId          *string                                                                       `json:"ConfigRuleId,omitempty" xml:"ConfigRuleId,omitempty"`
	ConfigRuleName        *string                                                                       `json:"ConfigRuleName,omitempty" xml:"ConfigRuleName,omitempty"`
	ConfigRuleParameters  []*GetCompliancePackResponseBodyCompliancePackConfigRulesConfigRuleParameters `json:"ConfigRuleParameters,omitempty" xml:"ConfigRuleParameters,omitempty" type:"Repeated"`
	Description           *string                                                                       `json:"Description,omitempty" xml:"Description,omitempty"`
	ManagedRuleIdentifier *string                                                                       `json:"ManagedRuleIdentifier,omitempty" xml:"ManagedRuleIdentifier,omitempty"`
	RiskLevel             *int32                                                                        `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
}

func (s GetCompliancePackResponseBodyCompliancePackConfigRules) String() string {
	return tea.Prettify(s)
}

func (s GetCompliancePackResponseBodyCompliancePackConfigRules) GoString() string {
	return s.String()
}

func (s *GetCompliancePackResponseBodyCompliancePackConfigRules) SetConfigRuleId(v string) *GetCompliancePackResponseBodyCompliancePackConfigRules {
	s.ConfigRuleId = &v
	return s
}

func (s *GetCompliancePackResponseBodyCompliancePackConfigRules) SetConfigRuleName(v string) *GetCompliancePackResponseBodyCompliancePackConfigRules {
	s.ConfigRuleName = &v
	return s
}

func (s *GetCompliancePackResponseBodyCompliancePackConfigRules) SetConfigRuleParameters(v []*GetCompliancePackResponseBodyCompliancePackConfigRulesConfigRuleParameters) *GetCompliancePackResponseBodyCompliancePackConfigRules {
	s.ConfigRuleParameters = v
	return s
}

func (s *GetCompliancePackResponseBodyCompliancePackConfigRules) SetDescription(v string) *GetCompliancePackResponseBodyCompliancePackConfigRules {
	s.Description = &v
	return s
}

func (s *GetCompliancePackResponseBodyCompliancePackConfigRules) SetManagedRuleIdentifier(v string) *GetCompliancePackResponseBodyCompliancePackConfigRules {
	s.ManagedRuleIdentifier = &v
	return s
}

func (s *GetCompliancePackResponseBodyCompliancePackConfigRules) SetRiskLevel(v int32) *GetCompliancePackResponseBodyCompliancePackConfigRules {
	s.RiskLevel = &v
	return s
}

type GetCompliancePackResponseBodyCompliancePackConfigRulesConfigRuleParameters struct {
	ParameterName  *string `json:"ParameterName,omitempty" xml:"ParameterName,omitempty"`
	ParameterValue *string `json:"ParameterValue,omitempty" xml:"ParameterValue,omitempty"`
	Required       *bool   `json:"Required,omitempty" xml:"Required,omitempty"`
}

func (s GetCompliancePackResponseBodyCompliancePackConfigRulesConfigRuleParameters) String() string {
	return tea.Prettify(s)
}

func (s GetCompliancePackResponseBodyCompliancePackConfigRulesConfigRuleParameters) GoString() string {
	return s.String()
}

func (s *GetCompliancePackResponseBodyCompliancePackConfigRulesConfigRuleParameters) SetParameterName(v string) *GetCompliancePackResponseBodyCompliancePackConfigRulesConfigRuleParameters {
	s.ParameterName = &v
	return s
}

func (s *GetCompliancePackResponseBodyCompliancePackConfigRulesConfigRuleParameters) SetParameterValue(v string) *GetCompliancePackResponseBodyCompliancePackConfigRulesConfigRuleParameters {
	s.ParameterValue = &v
	return s
}

func (s *GetCompliancePackResponseBodyCompliancePackConfigRulesConfigRuleParameters) SetRequired(v bool) *GetCompliancePackResponseBodyCompliancePackConfigRulesConfigRuleParameters {
	s.Required = &v
	return s
}

type GetCompliancePackResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetCompliancePackResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetCompliancePackResponse) String() string {
	return tea.Prettify(s)
}

func (s GetCompliancePackResponse) GoString() string {
	return s.String()
}

func (s *GetCompliancePackResponse) SetHeaders(v map[string]*string) *GetCompliancePackResponse {
	s.Headers = v
	return s
}

func (s *GetCompliancePackResponse) SetStatusCode(v int32) *GetCompliancePackResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCompliancePackResponse) SetBody(v *GetCompliancePackResponseBody) *GetCompliancePackResponse {
	s.Body = v
	return s
}

type GetCompliancePackReportRequest struct {
	CompliancePackId *string `json:"CompliancePackId,omitempty" xml:"CompliancePackId,omitempty"`
}

func (s GetCompliancePackReportRequest) String() string {
	return tea.Prettify(s)
}

func (s GetCompliancePackReportRequest) GoString() string {
	return s.String()
}

func (s *GetCompliancePackReportRequest) SetCompliancePackId(v string) *GetCompliancePackReportRequest {
	s.CompliancePackId = &v
	return s
}

type GetCompliancePackReportResponseBody struct {
	CompliancePackReport *GetCompliancePackReportResponseBodyCompliancePackReport `json:"CompliancePackReport,omitempty" xml:"CompliancePackReport,omitempty" type:"Struct"`
	RequestId            *string                                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetCompliancePackReportResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetCompliancePackReportResponseBody) GoString() string {
	return s.String()
}

func (s *GetCompliancePackReportResponseBody) SetCompliancePackReport(v *GetCompliancePackReportResponseBodyCompliancePackReport) *GetCompliancePackReportResponseBody {
	s.CompliancePackReport = v
	return s
}

func (s *GetCompliancePackReportResponseBody) SetRequestId(v string) *GetCompliancePackReportResponseBody {
	s.RequestId = &v
	return s
}

type GetCompliancePackReportResponseBodyCompliancePackReport struct {
	AccountId             *int64  `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	CompliancePackId      *string `json:"CompliancePackId,omitempty" xml:"CompliancePackId,omitempty"`
	ReportCreateTimestamp *int64  `json:"ReportCreateTimestamp,omitempty" xml:"ReportCreateTimestamp,omitempty"`
	ReportStatus          *string `json:"ReportStatus,omitempty" xml:"ReportStatus,omitempty"`
	ReportUrl             *string `json:"ReportUrl,omitempty" xml:"ReportUrl,omitempty"`
}

func (s GetCompliancePackReportResponseBodyCompliancePackReport) String() string {
	return tea.Prettify(s)
}

func (s GetCompliancePackReportResponseBodyCompliancePackReport) GoString() string {
	return s.String()
}

func (s *GetCompliancePackReportResponseBodyCompliancePackReport) SetAccountId(v int64) *GetCompliancePackReportResponseBodyCompliancePackReport {
	s.AccountId = &v
	return s
}

func (s *GetCompliancePackReportResponseBodyCompliancePackReport) SetCompliancePackId(v string) *GetCompliancePackReportResponseBodyCompliancePackReport {
	s.CompliancePackId = &v
	return s
}

func (s *GetCompliancePackReportResponseBodyCompliancePackReport) SetReportCreateTimestamp(v int64) *GetCompliancePackReportResponseBodyCompliancePackReport {
	s.ReportCreateTimestamp = &v
	return s
}

func (s *GetCompliancePackReportResponseBodyCompliancePackReport) SetReportStatus(v string) *GetCompliancePackReportResponseBodyCompliancePackReport {
	s.ReportStatus = &v
	return s
}

func (s *GetCompliancePackReportResponseBodyCompliancePackReport) SetReportUrl(v string) *GetCompliancePackReportResponseBodyCompliancePackReport {
	s.ReportUrl = &v
	return s
}

type GetCompliancePackReportResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetCompliancePackReportResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetCompliancePackReportResponse) String() string {
	return tea.Prettify(s)
}

func (s GetCompliancePackReportResponse) GoString() string {
	return s.String()
}

func (s *GetCompliancePackReportResponse) SetHeaders(v map[string]*string) *GetCompliancePackReportResponse {
	s.Headers = v
	return s
}

func (s *GetCompliancePackReportResponse) SetStatusCode(v int32) *GetCompliancePackReportResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCompliancePackReportResponse) SetBody(v *GetCompliancePackReportResponseBody) *GetCompliancePackReportResponse {
	s.Body = v
	return s
}

type GetConfigDeliveryChannelRequest struct {
	DeliveryChannelId *string `json:"DeliveryChannelId,omitempty" xml:"DeliveryChannelId,omitempty"`
}

func (s GetConfigDeliveryChannelRequest) String() string {
	return tea.Prettify(s)
}

func (s GetConfigDeliveryChannelRequest) GoString() string {
	return s.String()
}

func (s *GetConfigDeliveryChannelRequest) SetDeliveryChannelId(v string) *GetConfigDeliveryChannelRequest {
	s.DeliveryChannelId = &v
	return s
}

type GetConfigDeliveryChannelResponseBody struct {
	DeliveryChannel *GetConfigDeliveryChannelResponseBodyDeliveryChannel `json:"DeliveryChannel,omitempty" xml:"DeliveryChannel,omitempty" type:"Struct"`
	RequestId       *string                                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetConfigDeliveryChannelResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetConfigDeliveryChannelResponseBody) GoString() string {
	return s.String()
}

func (s *GetConfigDeliveryChannelResponseBody) SetDeliveryChannel(v *GetConfigDeliveryChannelResponseBodyDeliveryChannel) *GetConfigDeliveryChannelResponseBody {
	s.DeliveryChannel = v
	return s
}

func (s *GetConfigDeliveryChannelResponseBody) SetRequestId(v string) *GetConfigDeliveryChannelResponseBody {
	s.RequestId = &v
	return s
}

type GetConfigDeliveryChannelResponseBodyDeliveryChannel struct {
	AccountId                           *int64  `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	ConfigurationItemChangeNotification *bool   `json:"ConfigurationItemChangeNotification,omitempty" xml:"ConfigurationItemChangeNotification,omitempty"`
	ConfigurationSnapshot               *bool   `json:"ConfigurationSnapshot,omitempty" xml:"ConfigurationSnapshot,omitempty"`
	DeliveryChannelAssumeRoleArn        *string `json:"DeliveryChannelAssumeRoleArn,omitempty" xml:"DeliveryChannelAssumeRoleArn,omitempty"`
	DeliveryChannelCondition            *string `json:"DeliveryChannelCondition,omitempty" xml:"DeliveryChannelCondition,omitempty"`
	DeliveryChannelId                   *string `json:"DeliveryChannelId,omitempty" xml:"DeliveryChannelId,omitempty"`
	DeliveryChannelName                 *string `json:"DeliveryChannelName,omitempty" xml:"DeliveryChannelName,omitempty"`
	DeliveryChannelTargetArn            *string `json:"DeliveryChannelTargetArn,omitempty" xml:"DeliveryChannelTargetArn,omitempty"`
	DeliveryChannelType                 *string `json:"DeliveryChannelType,omitempty" xml:"DeliveryChannelType,omitempty"`
	Description                         *string `json:"Description,omitempty" xml:"Description,omitempty"`
	NonCompliantNotification            *bool   `json:"NonCompliantNotification,omitempty" xml:"NonCompliantNotification,omitempty"`
	OversizedDataOSSTargetArn           *string `json:"OversizedDataOSSTargetArn,omitempty" xml:"OversizedDataOSSTargetArn,omitempty"`
	Status                              *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetConfigDeliveryChannelResponseBodyDeliveryChannel) String() string {
	return tea.Prettify(s)
}

func (s GetConfigDeliveryChannelResponseBodyDeliveryChannel) GoString() string {
	return s.String()
}

func (s *GetConfigDeliveryChannelResponseBodyDeliveryChannel) SetAccountId(v int64) *GetConfigDeliveryChannelResponseBodyDeliveryChannel {
	s.AccountId = &v
	return s
}

func (s *GetConfigDeliveryChannelResponseBodyDeliveryChannel) SetConfigurationItemChangeNotification(v bool) *GetConfigDeliveryChannelResponseBodyDeliveryChannel {
	s.ConfigurationItemChangeNotification = &v
	return s
}

func (s *GetConfigDeliveryChannelResponseBodyDeliveryChannel) SetConfigurationSnapshot(v bool) *GetConfigDeliveryChannelResponseBodyDeliveryChannel {
	s.ConfigurationSnapshot = &v
	return s
}

func (s *GetConfigDeliveryChannelResponseBodyDeliveryChannel) SetDeliveryChannelAssumeRoleArn(v string) *GetConfigDeliveryChannelResponseBodyDeliveryChannel {
	s.DeliveryChannelAssumeRoleArn = &v
	return s
}

func (s *GetConfigDeliveryChannelResponseBodyDeliveryChannel) SetDeliveryChannelCondition(v string) *GetConfigDeliveryChannelResponseBodyDeliveryChannel {
	s.DeliveryChannelCondition = &v
	return s
}

func (s *GetConfigDeliveryChannelResponseBodyDeliveryChannel) SetDeliveryChannelId(v string) *GetConfigDeliveryChannelResponseBodyDeliveryChannel {
	s.DeliveryChannelId = &v
	return s
}

func (s *GetConfigDeliveryChannelResponseBodyDeliveryChannel) SetDeliveryChannelName(v string) *GetConfigDeliveryChannelResponseBodyDeliveryChannel {
	s.DeliveryChannelName = &v
	return s
}

func (s *GetConfigDeliveryChannelResponseBodyDeliveryChannel) SetDeliveryChannelTargetArn(v string) *GetConfigDeliveryChannelResponseBodyDeliveryChannel {
	s.DeliveryChannelTargetArn = &v
	return s
}

func (s *GetConfigDeliveryChannelResponseBodyDeliveryChannel) SetDeliveryChannelType(v string) *GetConfigDeliveryChannelResponseBodyDeliveryChannel {
	s.DeliveryChannelType = &v
	return s
}

func (s *GetConfigDeliveryChannelResponseBodyDeliveryChannel) SetDescription(v string) *GetConfigDeliveryChannelResponseBodyDeliveryChannel {
	s.Description = &v
	return s
}

func (s *GetConfigDeliveryChannelResponseBodyDeliveryChannel) SetNonCompliantNotification(v bool) *GetConfigDeliveryChannelResponseBodyDeliveryChannel {
	s.NonCompliantNotification = &v
	return s
}

func (s *GetConfigDeliveryChannelResponseBodyDeliveryChannel) SetOversizedDataOSSTargetArn(v string) *GetConfigDeliveryChannelResponseBodyDeliveryChannel {
	s.OversizedDataOSSTargetArn = &v
	return s
}

func (s *GetConfigDeliveryChannelResponseBodyDeliveryChannel) SetStatus(v int32) *GetConfigDeliveryChannelResponseBodyDeliveryChannel {
	s.Status = &v
	return s
}

type GetConfigDeliveryChannelResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetConfigDeliveryChannelResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetConfigDeliveryChannelResponse) String() string {
	return tea.Prettify(s)
}

func (s GetConfigDeliveryChannelResponse) GoString() string {
	return s.String()
}

func (s *GetConfigDeliveryChannelResponse) SetHeaders(v map[string]*string) *GetConfigDeliveryChannelResponse {
	s.Headers = v
	return s
}

func (s *GetConfigDeliveryChannelResponse) SetStatusCode(v int32) *GetConfigDeliveryChannelResponse {
	s.StatusCode = &v
	return s
}

func (s *GetConfigDeliveryChannelResponse) SetBody(v *GetConfigDeliveryChannelResponseBody) *GetConfigDeliveryChannelResponse {
	s.Body = v
	return s
}

type GetConfigRuleRequest struct {
	ConfigRuleId *string `json:"ConfigRuleId,omitempty" xml:"ConfigRuleId,omitempty"`
}

func (s GetConfigRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s GetConfigRuleRequest) GoString() string {
	return s.String()
}

func (s *GetConfigRuleRequest) SetConfigRuleId(v string) *GetConfigRuleRequest {
	s.ConfigRuleId = &v
	return s
}

type GetConfigRuleResponseBody struct {
	ConfigRule *GetConfigRuleResponseBodyConfigRule `json:"ConfigRule,omitempty" xml:"ConfigRule,omitempty" type:"Struct"`
	RequestId  *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetConfigRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetConfigRuleResponseBody) GoString() string {
	return s.String()
}

func (s *GetConfigRuleResponseBody) SetConfigRule(v *GetConfigRuleResponseBodyConfigRule) *GetConfigRuleResponseBody {
	s.ConfigRule = v
	return s
}

func (s *GetConfigRuleResponseBody) SetRequestId(v string) *GetConfigRuleResponseBody {
	s.RequestId = &v
	return s
}

type GetConfigRuleResponseBodyConfigRule struct {
	AccountId                  *int64                                                         `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	Compliance                 *GetConfigRuleResponseBodyConfigRuleCompliance                 `json:"Compliance,omitempty" xml:"Compliance,omitempty" type:"Struct"`
	ConfigRuleArn              *string                                                        `json:"ConfigRuleArn,omitempty" xml:"ConfigRuleArn,omitempty"`
	ConfigRuleEvaluationStatus *GetConfigRuleResponseBodyConfigRuleConfigRuleEvaluationStatus `json:"ConfigRuleEvaluationStatus,omitempty" xml:"ConfigRuleEvaluationStatus,omitempty" type:"Struct"`
	ConfigRuleId               *string                                                        `json:"ConfigRuleId,omitempty" xml:"ConfigRuleId,omitempty"`
	ConfigRuleName             *string                                                        `json:"ConfigRuleName,omitempty" xml:"ConfigRuleName,omitempty"`
	ConfigRuleState            *string                                                        `json:"ConfigRuleState,omitempty" xml:"ConfigRuleState,omitempty"`
	ConfigRuleTriggerTypes     *string                                                        `json:"ConfigRuleTriggerTypes,omitempty" xml:"ConfigRuleTriggerTypes,omitempty"`
	CreateBy                   *GetConfigRuleResponseBodyConfigRuleCreateBy                   `json:"CreateBy,omitempty" xml:"CreateBy,omitempty" type:"Struct"`
	CreateTimestamp            *int64                                                         `json:"CreateTimestamp,omitempty" xml:"CreateTimestamp,omitempty"`
	Description                *string                                                        `json:"Description,omitempty" xml:"Description,omitempty"`
	ExcludeResourceIdsScope    *string                                                        `json:"ExcludeResourceIdsScope,omitempty" xml:"ExcludeResourceIdsScope,omitempty"`
	InputParameters            map[string]interface{}                                         `json:"InputParameters,omitempty" xml:"InputParameters,omitempty"`
	ManagedRule                *GetConfigRuleResponseBodyConfigRuleManagedRule                `json:"ManagedRule,omitempty" xml:"ManagedRule,omitempty" type:"Struct"`
	MaximumExecutionFrequency  *string                                                        `json:"MaximumExecutionFrequency,omitempty" xml:"MaximumExecutionFrequency,omitempty"`
	ModifiedTimestamp          *int64                                                         `json:"ModifiedTimestamp,omitempty" xml:"ModifiedTimestamp,omitempty"`
	RegionIdsScope             *string                                                        `json:"RegionIdsScope,omitempty" xml:"RegionIdsScope,omitempty"`
	ResourceGroupIdsScope      *string                                                        `json:"ResourceGroupIdsScope,omitempty" xml:"ResourceGroupIdsScope,omitempty"`
	ResourceTypesScope         *string                                                        `json:"ResourceTypesScope,omitempty" xml:"ResourceTypesScope,omitempty"`
	RiskLevel                  *int32                                                         `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
	Source                     *GetConfigRuleResponseBodyConfigRuleSource                     `json:"Source,omitempty" xml:"Source,omitempty" type:"Struct"`
	TagKeyLogicScope           *string                                                        `json:"TagKeyLogicScope,omitempty" xml:"TagKeyLogicScope,omitempty"`
	TagKeyScope                *string                                                        `json:"TagKeyScope,omitempty" xml:"TagKeyScope,omitempty"`
	TagValueScope              *string                                                        `json:"TagValueScope,omitempty" xml:"TagValueScope,omitempty"`
}

func (s GetConfigRuleResponseBodyConfigRule) String() string {
	return tea.Prettify(s)
}

func (s GetConfigRuleResponseBodyConfigRule) GoString() string {
	return s.String()
}

func (s *GetConfigRuleResponseBodyConfigRule) SetAccountId(v int64) *GetConfigRuleResponseBodyConfigRule {
	s.AccountId = &v
	return s
}

func (s *GetConfigRuleResponseBodyConfigRule) SetCompliance(v *GetConfigRuleResponseBodyConfigRuleCompliance) *GetConfigRuleResponseBodyConfigRule {
	s.Compliance = v
	return s
}

func (s *GetConfigRuleResponseBodyConfigRule) SetConfigRuleArn(v string) *GetConfigRuleResponseBodyConfigRule {
	s.ConfigRuleArn = &v
	return s
}

func (s *GetConfigRuleResponseBodyConfigRule) SetConfigRuleEvaluationStatus(v *GetConfigRuleResponseBodyConfigRuleConfigRuleEvaluationStatus) *GetConfigRuleResponseBodyConfigRule {
	s.ConfigRuleEvaluationStatus = v
	return s
}

func (s *GetConfigRuleResponseBodyConfigRule) SetConfigRuleId(v string) *GetConfigRuleResponseBodyConfigRule {
	s.ConfigRuleId = &v
	return s
}

func (s *GetConfigRuleResponseBodyConfigRule) SetConfigRuleName(v string) *GetConfigRuleResponseBodyConfigRule {
	s.ConfigRuleName = &v
	return s
}

func (s *GetConfigRuleResponseBodyConfigRule) SetConfigRuleState(v string) *GetConfigRuleResponseBodyConfigRule {
	s.ConfigRuleState = &v
	return s
}

func (s *GetConfigRuleResponseBodyConfigRule) SetConfigRuleTriggerTypes(v string) *GetConfigRuleResponseBodyConfigRule {
	s.ConfigRuleTriggerTypes = &v
	return s
}

func (s *GetConfigRuleResponseBodyConfigRule) SetCreateBy(v *GetConfigRuleResponseBodyConfigRuleCreateBy) *GetConfigRuleResponseBodyConfigRule {
	s.CreateBy = v
	return s
}

func (s *GetConfigRuleResponseBodyConfigRule) SetCreateTimestamp(v int64) *GetConfigRuleResponseBodyConfigRule {
	s.CreateTimestamp = &v
	return s
}

func (s *GetConfigRuleResponseBodyConfigRule) SetDescription(v string) *GetConfigRuleResponseBodyConfigRule {
	s.Description = &v
	return s
}

func (s *GetConfigRuleResponseBodyConfigRule) SetExcludeResourceIdsScope(v string) *GetConfigRuleResponseBodyConfigRule {
	s.ExcludeResourceIdsScope = &v
	return s
}

func (s *GetConfigRuleResponseBodyConfigRule) SetInputParameters(v map[string]interface{}) *GetConfigRuleResponseBodyConfigRule {
	s.InputParameters = v
	return s
}

func (s *GetConfigRuleResponseBodyConfigRule) SetManagedRule(v *GetConfigRuleResponseBodyConfigRuleManagedRule) *GetConfigRuleResponseBodyConfigRule {
	s.ManagedRule = v
	return s
}

func (s *GetConfigRuleResponseBodyConfigRule) SetMaximumExecutionFrequency(v string) *GetConfigRuleResponseBodyConfigRule {
	s.MaximumExecutionFrequency = &v
	return s
}

func (s *GetConfigRuleResponseBodyConfigRule) SetModifiedTimestamp(v int64) *GetConfigRuleResponseBodyConfigRule {
	s.ModifiedTimestamp = &v
	return s
}

func (s *GetConfigRuleResponseBodyConfigRule) SetRegionIdsScope(v string) *GetConfigRuleResponseBodyConfigRule {
	s.RegionIdsScope = &v
	return s
}

func (s *GetConfigRuleResponseBodyConfigRule) SetResourceGroupIdsScope(v string) *GetConfigRuleResponseBodyConfigRule {
	s.ResourceGroupIdsScope = &v
	return s
}

func (s *GetConfigRuleResponseBodyConfigRule) SetResourceTypesScope(v string) *GetConfigRuleResponseBodyConfigRule {
	s.ResourceTypesScope = &v
	return s
}

func (s *GetConfigRuleResponseBodyConfigRule) SetRiskLevel(v int32) *GetConfigRuleResponseBodyConfigRule {
	s.RiskLevel = &v
	return s
}

func (s *GetConfigRuleResponseBodyConfigRule) SetSource(v *GetConfigRuleResponseBodyConfigRuleSource) *GetConfigRuleResponseBodyConfigRule {
	s.Source = v
	return s
}

func (s *GetConfigRuleResponseBodyConfigRule) SetTagKeyLogicScope(v string) *GetConfigRuleResponseBodyConfigRule {
	s.TagKeyLogicScope = &v
	return s
}

func (s *GetConfigRuleResponseBodyConfigRule) SetTagKeyScope(v string) *GetConfigRuleResponseBodyConfigRule {
	s.TagKeyScope = &v
	return s
}

func (s *GetConfigRuleResponseBodyConfigRule) SetTagValueScope(v string) *GetConfigRuleResponseBodyConfigRule {
	s.TagValueScope = &v
	return s
}

type GetConfigRuleResponseBodyConfigRuleCompliance struct {
	ComplianceType *string `json:"ComplianceType,omitempty" xml:"ComplianceType,omitempty"`
	Count          *int32  `json:"Count,omitempty" xml:"Count,omitempty"`
}

func (s GetConfigRuleResponseBodyConfigRuleCompliance) String() string {
	return tea.Prettify(s)
}

func (s GetConfigRuleResponseBodyConfigRuleCompliance) GoString() string {
	return s.String()
}

func (s *GetConfigRuleResponseBodyConfigRuleCompliance) SetComplianceType(v string) *GetConfigRuleResponseBodyConfigRuleCompliance {
	s.ComplianceType = &v
	return s
}

func (s *GetConfigRuleResponseBodyConfigRuleCompliance) SetCount(v int32) *GetConfigRuleResponseBodyConfigRuleCompliance {
	s.Count = &v
	return s
}

type GetConfigRuleResponseBodyConfigRuleConfigRuleEvaluationStatus struct {
	FirstActivatedTimestamp           *int64  `json:"FirstActivatedTimestamp,omitempty" xml:"FirstActivatedTimestamp,omitempty"`
	FirstEvaluationStarted            *bool   `json:"FirstEvaluationStarted,omitempty" xml:"FirstEvaluationStarted,omitempty"`
	LastErrorCode                     *string `json:"LastErrorCode,omitempty" xml:"LastErrorCode,omitempty"`
	LastErrorMessage                  *string `json:"LastErrorMessage,omitempty" xml:"LastErrorMessage,omitempty"`
	LastFailedEvaluationTimestamp     *int64  `json:"LastFailedEvaluationTimestamp,omitempty" xml:"LastFailedEvaluationTimestamp,omitempty"`
	LastFailedInvocationTimestamp     *int64  `json:"LastFailedInvocationTimestamp,omitempty" xml:"LastFailedInvocationTimestamp,omitempty"`
	LastSuccessfulEvaluationTimestamp *int64  `json:"LastSuccessfulEvaluationTimestamp,omitempty" xml:"LastSuccessfulEvaluationTimestamp,omitempty"`
	LastSuccessfulInvocationTimestamp *int64  `json:"LastSuccessfulInvocationTimestamp,omitempty" xml:"LastSuccessfulInvocationTimestamp,omitempty"`
}

func (s GetConfigRuleResponseBodyConfigRuleConfigRuleEvaluationStatus) String() string {
	return tea.Prettify(s)
}

func (s GetConfigRuleResponseBodyConfigRuleConfigRuleEvaluationStatus) GoString() string {
	return s.String()
}

func (s *GetConfigRuleResponseBodyConfigRuleConfigRuleEvaluationStatus) SetFirstActivatedTimestamp(v int64) *GetConfigRuleResponseBodyConfigRuleConfigRuleEvaluationStatus {
	s.FirstActivatedTimestamp = &v
	return s
}

func (s *GetConfigRuleResponseBodyConfigRuleConfigRuleEvaluationStatus) SetFirstEvaluationStarted(v bool) *GetConfigRuleResponseBodyConfigRuleConfigRuleEvaluationStatus {
	s.FirstEvaluationStarted = &v
	return s
}

func (s *GetConfigRuleResponseBodyConfigRuleConfigRuleEvaluationStatus) SetLastErrorCode(v string) *GetConfigRuleResponseBodyConfigRuleConfigRuleEvaluationStatus {
	s.LastErrorCode = &v
	return s
}

func (s *GetConfigRuleResponseBodyConfigRuleConfigRuleEvaluationStatus) SetLastErrorMessage(v string) *GetConfigRuleResponseBodyConfigRuleConfigRuleEvaluationStatus {
	s.LastErrorMessage = &v
	return s
}

func (s *GetConfigRuleResponseBodyConfigRuleConfigRuleEvaluationStatus) SetLastFailedEvaluationTimestamp(v int64) *GetConfigRuleResponseBodyConfigRuleConfigRuleEvaluationStatus {
	s.LastFailedEvaluationTimestamp = &v
	return s
}

func (s *GetConfigRuleResponseBodyConfigRuleConfigRuleEvaluationStatus) SetLastFailedInvocationTimestamp(v int64) *GetConfigRuleResponseBodyConfigRuleConfigRuleEvaluationStatus {
	s.LastFailedInvocationTimestamp = &v
	return s
}

func (s *GetConfigRuleResponseBodyConfigRuleConfigRuleEvaluationStatus) SetLastSuccessfulEvaluationTimestamp(v int64) *GetConfigRuleResponseBodyConfigRuleConfigRuleEvaluationStatus {
	s.LastSuccessfulEvaluationTimestamp = &v
	return s
}

func (s *GetConfigRuleResponseBodyConfigRuleConfigRuleEvaluationStatus) SetLastSuccessfulInvocationTimestamp(v int64) *GetConfigRuleResponseBodyConfigRuleConfigRuleEvaluationStatus {
	s.LastSuccessfulInvocationTimestamp = &v
	return s
}

type GetConfigRuleResponseBodyConfigRuleCreateBy struct {
	CompliancePackId   *string `json:"CompliancePackId,omitempty" xml:"CompliancePackId,omitempty"`
	CompliancePackName *string `json:"CompliancePackName,omitempty" xml:"CompliancePackName,omitempty"`
	CreatorId          *string `json:"CreatorId,omitempty" xml:"CreatorId,omitempty"`
	CreatorName        *string `json:"CreatorName,omitempty" xml:"CreatorName,omitempty"`
}

func (s GetConfigRuleResponseBodyConfigRuleCreateBy) String() string {
	return tea.Prettify(s)
}

func (s GetConfigRuleResponseBodyConfigRuleCreateBy) GoString() string {
	return s.String()
}

func (s *GetConfigRuleResponseBodyConfigRuleCreateBy) SetCompliancePackId(v string) *GetConfigRuleResponseBodyConfigRuleCreateBy {
	s.CompliancePackId = &v
	return s
}

func (s *GetConfigRuleResponseBodyConfigRuleCreateBy) SetCompliancePackName(v string) *GetConfigRuleResponseBodyConfigRuleCreateBy {
	s.CompliancePackName = &v
	return s
}

func (s *GetConfigRuleResponseBodyConfigRuleCreateBy) SetCreatorId(v string) *GetConfigRuleResponseBodyConfigRuleCreateBy {
	s.CreatorId = &v
	return s
}

func (s *GetConfigRuleResponseBodyConfigRuleCreateBy) SetCreatorName(v string) *GetConfigRuleResponseBodyConfigRuleCreateBy {
	s.CreatorName = &v
	return s
}

type GetConfigRuleResponseBodyConfigRuleManagedRule struct {
	CompulsoryInputParameterDetails map[string]interface{}                                         `json:"CompulsoryInputParameterDetails,omitempty" xml:"CompulsoryInputParameterDetails,omitempty"`
	Description                     *string                                                        `json:"Description,omitempty" xml:"Description,omitempty"`
	Identifier                      *string                                                        `json:"Identifier,omitempty" xml:"Identifier,omitempty"`
	Labels                          []*string                                                      `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	ManagedRuleName                 *string                                                        `json:"ManagedRuleName,omitempty" xml:"ManagedRuleName,omitempty"`
	OptionalInputParameterDetails   map[string]interface{}                                         `json:"OptionalInputParameterDetails,omitempty" xml:"OptionalInputParameterDetails,omitempty"`
	SourceDetails                   []*GetConfigRuleResponseBodyConfigRuleManagedRuleSourceDetails `json:"SourceDetails,omitempty" xml:"SourceDetails,omitempty" type:"Repeated"`
}

func (s GetConfigRuleResponseBodyConfigRuleManagedRule) String() string {
	return tea.Prettify(s)
}

func (s GetConfigRuleResponseBodyConfigRuleManagedRule) GoString() string {
	return s.String()
}

func (s *GetConfigRuleResponseBodyConfigRuleManagedRule) SetCompulsoryInputParameterDetails(v map[string]interface{}) *GetConfigRuleResponseBodyConfigRuleManagedRule {
	s.CompulsoryInputParameterDetails = v
	return s
}

func (s *GetConfigRuleResponseBodyConfigRuleManagedRule) SetDescription(v string) *GetConfigRuleResponseBodyConfigRuleManagedRule {
	s.Description = &v
	return s
}

func (s *GetConfigRuleResponseBodyConfigRuleManagedRule) SetIdentifier(v string) *GetConfigRuleResponseBodyConfigRuleManagedRule {
	s.Identifier = &v
	return s
}

func (s *GetConfigRuleResponseBodyConfigRuleManagedRule) SetLabels(v []*string) *GetConfigRuleResponseBodyConfigRuleManagedRule {
	s.Labels = v
	return s
}

func (s *GetConfigRuleResponseBodyConfigRuleManagedRule) SetManagedRuleName(v string) *GetConfigRuleResponseBodyConfigRuleManagedRule {
	s.ManagedRuleName = &v
	return s
}

func (s *GetConfigRuleResponseBodyConfigRuleManagedRule) SetOptionalInputParameterDetails(v map[string]interface{}) *GetConfigRuleResponseBodyConfigRuleManagedRule {
	s.OptionalInputParameterDetails = v
	return s
}

func (s *GetConfigRuleResponseBodyConfigRuleManagedRule) SetSourceDetails(v []*GetConfigRuleResponseBodyConfigRuleManagedRuleSourceDetails) *GetConfigRuleResponseBodyConfigRuleManagedRule {
	s.SourceDetails = v
	return s
}

type GetConfigRuleResponseBodyConfigRuleManagedRuleSourceDetails struct {
	EventSource               *string `json:"EventSource,omitempty" xml:"EventSource,omitempty"`
	MaximumExecutionFrequency *string `json:"MaximumExecutionFrequency,omitempty" xml:"MaximumExecutionFrequency,omitempty"`
	MessageType               *string `json:"MessageType,omitempty" xml:"MessageType,omitempty"`
}

func (s GetConfigRuleResponseBodyConfigRuleManagedRuleSourceDetails) String() string {
	return tea.Prettify(s)
}

func (s GetConfigRuleResponseBodyConfigRuleManagedRuleSourceDetails) GoString() string {
	return s.String()
}

func (s *GetConfigRuleResponseBodyConfigRuleManagedRuleSourceDetails) SetEventSource(v string) *GetConfigRuleResponseBodyConfigRuleManagedRuleSourceDetails {
	s.EventSource = &v
	return s
}

func (s *GetConfigRuleResponseBodyConfigRuleManagedRuleSourceDetails) SetMaximumExecutionFrequency(v string) *GetConfigRuleResponseBodyConfigRuleManagedRuleSourceDetails {
	s.MaximumExecutionFrequency = &v
	return s
}

func (s *GetConfigRuleResponseBodyConfigRuleManagedRuleSourceDetails) SetMessageType(v string) *GetConfigRuleResponseBodyConfigRuleManagedRuleSourceDetails {
	s.MessageType = &v
	return s
}

type GetConfigRuleResponseBodyConfigRuleSource struct {
	Identifier    *string                                                   `json:"Identifier,omitempty" xml:"Identifier,omitempty"`
	Owner         *string                                                   `json:"Owner,omitempty" xml:"Owner,omitempty"`
	SourceDetails []*GetConfigRuleResponseBodyConfigRuleSourceSourceDetails `json:"SourceDetails,omitempty" xml:"SourceDetails,omitempty" type:"Repeated"`
}

func (s GetConfigRuleResponseBodyConfigRuleSource) String() string {
	return tea.Prettify(s)
}

func (s GetConfigRuleResponseBodyConfigRuleSource) GoString() string {
	return s.String()
}

func (s *GetConfigRuleResponseBodyConfigRuleSource) SetIdentifier(v string) *GetConfigRuleResponseBodyConfigRuleSource {
	s.Identifier = &v
	return s
}

func (s *GetConfigRuleResponseBodyConfigRuleSource) SetOwner(v string) *GetConfigRuleResponseBodyConfigRuleSource {
	s.Owner = &v
	return s
}

func (s *GetConfigRuleResponseBodyConfigRuleSource) SetSourceDetails(v []*GetConfigRuleResponseBodyConfigRuleSourceSourceDetails) *GetConfigRuleResponseBodyConfigRuleSource {
	s.SourceDetails = v
	return s
}

type GetConfigRuleResponseBodyConfigRuleSourceSourceDetails struct {
	EventSource               *string `json:"EventSource,omitempty" xml:"EventSource,omitempty"`
	MaximumExecutionFrequency *string `json:"MaximumExecutionFrequency,omitempty" xml:"MaximumExecutionFrequency,omitempty"`
	MessageType               *string `json:"MessageType,omitempty" xml:"MessageType,omitempty"`
}

func (s GetConfigRuleResponseBodyConfigRuleSourceSourceDetails) String() string {
	return tea.Prettify(s)
}

func (s GetConfigRuleResponseBodyConfigRuleSourceSourceDetails) GoString() string {
	return s.String()
}

func (s *GetConfigRuleResponseBodyConfigRuleSourceSourceDetails) SetEventSource(v string) *GetConfigRuleResponseBodyConfigRuleSourceSourceDetails {
	s.EventSource = &v
	return s
}

func (s *GetConfigRuleResponseBodyConfigRuleSourceSourceDetails) SetMaximumExecutionFrequency(v string) *GetConfigRuleResponseBodyConfigRuleSourceSourceDetails {
	s.MaximumExecutionFrequency = &v
	return s
}

func (s *GetConfigRuleResponseBodyConfigRuleSourceSourceDetails) SetMessageType(v string) *GetConfigRuleResponseBodyConfigRuleSourceSourceDetails {
	s.MessageType = &v
	return s
}

type GetConfigRuleResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetConfigRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetConfigRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s GetConfigRuleResponse) GoString() string {
	return s.String()
}

func (s *GetConfigRuleResponse) SetHeaders(v map[string]*string) *GetConfigRuleResponse {
	s.Headers = v
	return s
}

func (s *GetConfigRuleResponse) SetStatusCode(v int32) *GetConfigRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *GetConfigRuleResponse) SetBody(v *GetConfigRuleResponseBody) *GetConfigRuleResponse {
	s.Body = v
	return s
}

type GetConfigRuleComplianceByPackRequest struct {
	CompliancePackId *string `json:"CompliancePackId,omitempty" xml:"CompliancePackId,omitempty"`
}

func (s GetConfigRuleComplianceByPackRequest) String() string {
	return tea.Prettify(s)
}

func (s GetConfigRuleComplianceByPackRequest) GoString() string {
	return s.String()
}

func (s *GetConfigRuleComplianceByPackRequest) SetCompliancePackId(v string) *GetConfigRuleComplianceByPackRequest {
	s.CompliancePackId = &v
	return s
}

type GetConfigRuleComplianceByPackResponseBody struct {
	ConfigRuleComplianceResult *GetConfigRuleComplianceByPackResponseBodyConfigRuleComplianceResult `json:"ConfigRuleComplianceResult,omitempty" xml:"ConfigRuleComplianceResult,omitempty" type:"Struct"`
	RequestId                  *string                                                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetConfigRuleComplianceByPackResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetConfigRuleComplianceByPackResponseBody) GoString() string {
	return s.String()
}

func (s *GetConfigRuleComplianceByPackResponseBody) SetConfigRuleComplianceResult(v *GetConfigRuleComplianceByPackResponseBodyConfigRuleComplianceResult) *GetConfigRuleComplianceByPackResponseBody {
	s.ConfigRuleComplianceResult = v
	return s
}

func (s *GetConfigRuleComplianceByPackResponseBody) SetRequestId(v string) *GetConfigRuleComplianceByPackResponseBody {
	s.RequestId = &v
	return s
}

type GetConfigRuleComplianceByPackResponseBodyConfigRuleComplianceResult struct {
	CompliancePackId      *string                                                                                     `json:"CompliancePackId,omitempty" xml:"CompliancePackId,omitempty"`
	ConfigRuleCompliances []*GetConfigRuleComplianceByPackResponseBodyConfigRuleComplianceResultConfigRuleCompliances `json:"ConfigRuleCompliances,omitempty" xml:"ConfigRuleCompliances,omitempty" type:"Repeated"`
	NonCompliantCount     *int32                                                                                      `json:"NonCompliantCount,omitempty" xml:"NonCompliantCount,omitempty"`
	TotalCount            *int32                                                                                      `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s GetConfigRuleComplianceByPackResponseBodyConfigRuleComplianceResult) String() string {
	return tea.Prettify(s)
}

func (s GetConfigRuleComplianceByPackResponseBodyConfigRuleComplianceResult) GoString() string {
	return s.String()
}

func (s *GetConfigRuleComplianceByPackResponseBodyConfigRuleComplianceResult) SetCompliancePackId(v string) *GetConfigRuleComplianceByPackResponseBodyConfigRuleComplianceResult {
	s.CompliancePackId = &v
	return s
}

func (s *GetConfigRuleComplianceByPackResponseBodyConfigRuleComplianceResult) SetConfigRuleCompliances(v []*GetConfigRuleComplianceByPackResponseBodyConfigRuleComplianceResultConfigRuleCompliances) *GetConfigRuleComplianceByPackResponseBodyConfigRuleComplianceResult {
	s.ConfigRuleCompliances = v
	return s
}

func (s *GetConfigRuleComplianceByPackResponseBodyConfigRuleComplianceResult) SetNonCompliantCount(v int32) *GetConfigRuleComplianceByPackResponseBodyConfigRuleComplianceResult {
	s.NonCompliantCount = &v
	return s
}

func (s *GetConfigRuleComplianceByPackResponseBodyConfigRuleComplianceResult) SetTotalCount(v int32) *GetConfigRuleComplianceByPackResponseBodyConfigRuleComplianceResult {
	s.TotalCount = &v
	return s
}

type GetConfigRuleComplianceByPackResponseBodyConfigRuleComplianceResultConfigRuleCompliances struct {
	ComplianceType *string `json:"ComplianceType,omitempty" xml:"ComplianceType,omitempty"`
	ConfigRuleId   *string `json:"ConfigRuleId,omitempty" xml:"ConfigRuleId,omitempty"`
	ConfigRuleName *string `json:"ConfigRuleName,omitempty" xml:"ConfigRuleName,omitempty"`
}

func (s GetConfigRuleComplianceByPackResponseBodyConfigRuleComplianceResultConfigRuleCompliances) String() string {
	return tea.Prettify(s)
}

func (s GetConfigRuleComplianceByPackResponseBodyConfigRuleComplianceResultConfigRuleCompliances) GoString() string {
	return s.String()
}

func (s *GetConfigRuleComplianceByPackResponseBodyConfigRuleComplianceResultConfigRuleCompliances) SetComplianceType(v string) *GetConfigRuleComplianceByPackResponseBodyConfigRuleComplianceResultConfigRuleCompliances {
	s.ComplianceType = &v
	return s
}

func (s *GetConfigRuleComplianceByPackResponseBodyConfigRuleComplianceResultConfigRuleCompliances) SetConfigRuleId(v string) *GetConfigRuleComplianceByPackResponseBodyConfigRuleComplianceResultConfigRuleCompliances {
	s.ConfigRuleId = &v
	return s
}

func (s *GetConfigRuleComplianceByPackResponseBodyConfigRuleComplianceResultConfigRuleCompliances) SetConfigRuleName(v string) *GetConfigRuleComplianceByPackResponseBodyConfigRuleComplianceResultConfigRuleCompliances {
	s.ConfigRuleName = &v
	return s
}

type GetConfigRuleComplianceByPackResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetConfigRuleComplianceByPackResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetConfigRuleComplianceByPackResponse) String() string {
	return tea.Prettify(s)
}

func (s GetConfigRuleComplianceByPackResponse) GoString() string {
	return s.String()
}

func (s *GetConfigRuleComplianceByPackResponse) SetHeaders(v map[string]*string) *GetConfigRuleComplianceByPackResponse {
	s.Headers = v
	return s
}

func (s *GetConfigRuleComplianceByPackResponse) SetStatusCode(v int32) *GetConfigRuleComplianceByPackResponse {
	s.StatusCode = &v
	return s
}

func (s *GetConfigRuleComplianceByPackResponse) SetBody(v *GetConfigRuleComplianceByPackResponseBody) *GetConfigRuleComplianceByPackResponse {
	s.Body = v
	return s
}

type GetConfigRuleSummaryByRiskLevelResponseBody struct {
	ConfigRuleSummaries []*GetConfigRuleSummaryByRiskLevelResponseBodyConfigRuleSummaries `json:"ConfigRuleSummaries,omitempty" xml:"ConfigRuleSummaries,omitempty" type:"Repeated"`
	RequestId           *string                                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetConfigRuleSummaryByRiskLevelResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetConfigRuleSummaryByRiskLevelResponseBody) GoString() string {
	return s.String()
}

func (s *GetConfigRuleSummaryByRiskLevelResponseBody) SetConfigRuleSummaries(v []*GetConfigRuleSummaryByRiskLevelResponseBodyConfigRuleSummaries) *GetConfigRuleSummaryByRiskLevelResponseBody {
	s.ConfigRuleSummaries = v
	return s
}

func (s *GetConfigRuleSummaryByRiskLevelResponseBody) SetRequestId(v string) *GetConfigRuleSummaryByRiskLevelResponseBody {
	s.RequestId = &v
	return s
}

type GetConfigRuleSummaryByRiskLevelResponseBodyConfigRuleSummaries struct {
	CompliantCount    *int32 `json:"CompliantCount,omitempty" xml:"CompliantCount,omitempty"`
	NonCompliantCount *int32 `json:"NonCompliantCount,omitempty" xml:"NonCompliantCount,omitempty"`
	RiskLevel         *int32 `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
}

func (s GetConfigRuleSummaryByRiskLevelResponseBodyConfigRuleSummaries) String() string {
	return tea.Prettify(s)
}

func (s GetConfigRuleSummaryByRiskLevelResponseBodyConfigRuleSummaries) GoString() string {
	return s.String()
}

func (s *GetConfigRuleSummaryByRiskLevelResponseBodyConfigRuleSummaries) SetCompliantCount(v int32) *GetConfigRuleSummaryByRiskLevelResponseBodyConfigRuleSummaries {
	s.CompliantCount = &v
	return s
}

func (s *GetConfigRuleSummaryByRiskLevelResponseBodyConfigRuleSummaries) SetNonCompliantCount(v int32) *GetConfigRuleSummaryByRiskLevelResponseBodyConfigRuleSummaries {
	s.NonCompliantCount = &v
	return s
}

func (s *GetConfigRuleSummaryByRiskLevelResponseBodyConfigRuleSummaries) SetRiskLevel(v int32) *GetConfigRuleSummaryByRiskLevelResponseBodyConfigRuleSummaries {
	s.RiskLevel = &v
	return s
}

type GetConfigRuleSummaryByRiskLevelResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetConfigRuleSummaryByRiskLevelResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetConfigRuleSummaryByRiskLevelResponse) String() string {
	return tea.Prettify(s)
}

func (s GetConfigRuleSummaryByRiskLevelResponse) GoString() string {
	return s.String()
}

func (s *GetConfigRuleSummaryByRiskLevelResponse) SetHeaders(v map[string]*string) *GetConfigRuleSummaryByRiskLevelResponse {
	s.Headers = v
	return s
}

func (s *GetConfigRuleSummaryByRiskLevelResponse) SetStatusCode(v int32) *GetConfigRuleSummaryByRiskLevelResponse {
	s.StatusCode = &v
	return s
}

func (s *GetConfigRuleSummaryByRiskLevelResponse) SetBody(v *GetConfigRuleSummaryByRiskLevelResponseBody) *GetConfigRuleSummaryByRiskLevelResponse {
	s.Body = v
	return s
}

type GetConfigRulesReportRequest struct {
	ReportId *string `json:"ReportId,omitempty" xml:"ReportId,omitempty"`
}

func (s GetConfigRulesReportRequest) String() string {
	return tea.Prettify(s)
}

func (s GetConfigRulesReportRequest) GoString() string {
	return s.String()
}

func (s *GetConfigRulesReportRequest) SetReportId(v string) *GetConfigRulesReportRequest {
	s.ReportId = &v
	return s
}

type GetConfigRulesReportResponseBody struct {
	ConfigRulesReport *GetConfigRulesReportResponseBodyConfigRulesReport `json:"ConfigRulesReport,omitempty" xml:"ConfigRulesReport,omitempty" type:"Struct"`
	RequestId         *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetConfigRulesReportResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetConfigRulesReportResponseBody) GoString() string {
	return s.String()
}

func (s *GetConfigRulesReportResponseBody) SetConfigRulesReport(v *GetConfigRulesReportResponseBodyConfigRulesReport) *GetConfigRulesReportResponseBody {
	s.ConfigRulesReport = v
	return s
}

func (s *GetConfigRulesReportResponseBody) SetRequestId(v string) *GetConfigRulesReportResponseBody {
	s.RequestId = &v
	return s
}

type GetConfigRulesReportResponseBodyConfigRulesReport struct {
	AccountId             *int64  `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	ReportCreateTimestamp *int64  `json:"ReportCreateTimestamp,omitempty" xml:"ReportCreateTimestamp,omitempty"`
	ReportId              *string `json:"ReportId,omitempty" xml:"ReportId,omitempty"`
	ReportStatus          *string `json:"ReportStatus,omitempty" xml:"ReportStatus,omitempty"`
	ReportUrl             *string `json:"ReportUrl,omitempty" xml:"ReportUrl,omitempty"`
}

func (s GetConfigRulesReportResponseBodyConfigRulesReport) String() string {
	return tea.Prettify(s)
}

func (s GetConfigRulesReportResponseBodyConfigRulesReport) GoString() string {
	return s.String()
}

func (s *GetConfigRulesReportResponseBodyConfigRulesReport) SetAccountId(v int64) *GetConfigRulesReportResponseBodyConfigRulesReport {
	s.AccountId = &v
	return s
}

func (s *GetConfigRulesReportResponseBodyConfigRulesReport) SetReportCreateTimestamp(v int64) *GetConfigRulesReportResponseBodyConfigRulesReport {
	s.ReportCreateTimestamp = &v
	return s
}

func (s *GetConfigRulesReportResponseBodyConfigRulesReport) SetReportId(v string) *GetConfigRulesReportResponseBodyConfigRulesReport {
	s.ReportId = &v
	return s
}

func (s *GetConfigRulesReportResponseBodyConfigRulesReport) SetReportStatus(v string) *GetConfigRulesReportResponseBodyConfigRulesReport {
	s.ReportStatus = &v
	return s
}

func (s *GetConfigRulesReportResponseBodyConfigRulesReport) SetReportUrl(v string) *GetConfigRulesReportResponseBodyConfigRulesReport {
	s.ReportUrl = &v
	return s
}

type GetConfigRulesReportResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetConfigRulesReportResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetConfigRulesReportResponse) String() string {
	return tea.Prettify(s)
}

func (s GetConfigRulesReportResponse) GoString() string {
	return s.String()
}

func (s *GetConfigRulesReportResponse) SetHeaders(v map[string]*string) *GetConfigRulesReportResponse {
	s.Headers = v
	return s
}

func (s *GetConfigRulesReportResponse) SetStatusCode(v int32) *GetConfigRulesReportResponse {
	s.StatusCode = &v
	return s
}

func (s *GetConfigRulesReportResponse) SetBody(v *GetConfigRulesReportResponseBody) *GetConfigRulesReportResponse {
	s.Body = v
	return s
}

type GetDiscoveredResourceRequest struct {
	Region       *string `json:"Region,omitempty" xml:"Region,omitempty"`
	ResourceId   *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s GetDiscoveredResourceRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDiscoveredResourceRequest) GoString() string {
	return s.String()
}

func (s *GetDiscoveredResourceRequest) SetRegion(v string) *GetDiscoveredResourceRequest {
	s.Region = &v
	return s
}

func (s *GetDiscoveredResourceRequest) SetResourceId(v string) *GetDiscoveredResourceRequest {
	s.ResourceId = &v
	return s
}

func (s *GetDiscoveredResourceRequest) SetResourceType(v string) *GetDiscoveredResourceRequest {
	s.ResourceType = &v
	return s
}

type GetDiscoveredResourceResponseBody struct {
	DiscoveredResourceDetail *GetDiscoveredResourceResponseBodyDiscoveredResourceDetail `json:"DiscoveredResourceDetail,omitempty" xml:"DiscoveredResourceDetail,omitempty" type:"Struct"`
	RequestId                *string                                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetDiscoveredResourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDiscoveredResourceResponseBody) GoString() string {
	return s.String()
}

func (s *GetDiscoveredResourceResponseBody) SetDiscoveredResourceDetail(v *GetDiscoveredResourceResponseBodyDiscoveredResourceDetail) *GetDiscoveredResourceResponseBody {
	s.DiscoveredResourceDetail = v
	return s
}

func (s *GetDiscoveredResourceResponseBody) SetRequestId(v string) *GetDiscoveredResourceResponseBody {
	s.RequestId = &v
	return s
}

type GetDiscoveredResourceResponseBodyDiscoveredResourceDetail struct {
	AccountId            *int64  `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	AvailabilityZone     *string `json:"AvailabilityZone,omitempty" xml:"AvailabilityZone,omitempty"`
	Configuration        *string `json:"Configuration,omitempty" xml:"Configuration,omitempty"`
	Region               *string `json:"Region,omitempty" xml:"Region,omitempty"`
	ResourceCreationTime *int64  `json:"ResourceCreationTime,omitempty" xml:"ResourceCreationTime,omitempty"`
	ResourceDeleted      *int32  `json:"ResourceDeleted,omitempty" xml:"ResourceDeleted,omitempty"`
	ResourceId           *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ResourceName         *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	ResourceStatus       *string `json:"ResourceStatus,omitempty" xml:"ResourceStatus,omitempty"`
	ResourceType         *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	Tags                 *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
}

func (s GetDiscoveredResourceResponseBodyDiscoveredResourceDetail) String() string {
	return tea.Prettify(s)
}

func (s GetDiscoveredResourceResponseBodyDiscoveredResourceDetail) GoString() string {
	return s.String()
}

func (s *GetDiscoveredResourceResponseBodyDiscoveredResourceDetail) SetAccountId(v int64) *GetDiscoveredResourceResponseBodyDiscoveredResourceDetail {
	s.AccountId = &v
	return s
}

func (s *GetDiscoveredResourceResponseBodyDiscoveredResourceDetail) SetAvailabilityZone(v string) *GetDiscoveredResourceResponseBodyDiscoveredResourceDetail {
	s.AvailabilityZone = &v
	return s
}

func (s *GetDiscoveredResourceResponseBodyDiscoveredResourceDetail) SetConfiguration(v string) *GetDiscoveredResourceResponseBodyDiscoveredResourceDetail {
	s.Configuration = &v
	return s
}

func (s *GetDiscoveredResourceResponseBodyDiscoveredResourceDetail) SetRegion(v string) *GetDiscoveredResourceResponseBodyDiscoveredResourceDetail {
	s.Region = &v
	return s
}

func (s *GetDiscoveredResourceResponseBodyDiscoveredResourceDetail) SetResourceCreationTime(v int64) *GetDiscoveredResourceResponseBodyDiscoveredResourceDetail {
	s.ResourceCreationTime = &v
	return s
}

func (s *GetDiscoveredResourceResponseBodyDiscoveredResourceDetail) SetResourceDeleted(v int32) *GetDiscoveredResourceResponseBodyDiscoveredResourceDetail {
	s.ResourceDeleted = &v
	return s
}

func (s *GetDiscoveredResourceResponseBodyDiscoveredResourceDetail) SetResourceId(v string) *GetDiscoveredResourceResponseBodyDiscoveredResourceDetail {
	s.ResourceId = &v
	return s
}

func (s *GetDiscoveredResourceResponseBodyDiscoveredResourceDetail) SetResourceName(v string) *GetDiscoveredResourceResponseBodyDiscoveredResourceDetail {
	s.ResourceName = &v
	return s
}

func (s *GetDiscoveredResourceResponseBodyDiscoveredResourceDetail) SetResourceStatus(v string) *GetDiscoveredResourceResponseBodyDiscoveredResourceDetail {
	s.ResourceStatus = &v
	return s
}

func (s *GetDiscoveredResourceResponseBodyDiscoveredResourceDetail) SetResourceType(v string) *GetDiscoveredResourceResponseBodyDiscoveredResourceDetail {
	s.ResourceType = &v
	return s
}

func (s *GetDiscoveredResourceResponseBodyDiscoveredResourceDetail) SetTags(v string) *GetDiscoveredResourceResponseBodyDiscoveredResourceDetail {
	s.Tags = &v
	return s
}

type GetDiscoveredResourceResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetDiscoveredResourceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetDiscoveredResourceResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDiscoveredResourceResponse) GoString() string {
	return s.String()
}

func (s *GetDiscoveredResourceResponse) SetHeaders(v map[string]*string) *GetDiscoveredResourceResponse {
	s.Headers = v
	return s
}

func (s *GetDiscoveredResourceResponse) SetStatusCode(v int32) *GetDiscoveredResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDiscoveredResourceResponse) SetBody(v *GetDiscoveredResourceResponseBody) *GetDiscoveredResourceResponse {
	s.Body = v
	return s
}

type GetDiscoveredResourceCountsGroupByRegionRequest struct {
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s GetDiscoveredResourceCountsGroupByRegionRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDiscoveredResourceCountsGroupByRegionRequest) GoString() string {
	return s.String()
}

func (s *GetDiscoveredResourceCountsGroupByRegionRequest) SetResourceType(v string) *GetDiscoveredResourceCountsGroupByRegionRequest {
	s.ResourceType = &v
	return s
}

type GetDiscoveredResourceCountsGroupByRegionResponseBody struct {
	DiscoveredResourceCountsSummary []*GetDiscoveredResourceCountsGroupByRegionResponseBodyDiscoveredResourceCountsSummary `json:"DiscoveredResourceCountsSummary,omitempty" xml:"DiscoveredResourceCountsSummary,omitempty" type:"Repeated"`
	RequestId                       *string                                                                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetDiscoveredResourceCountsGroupByRegionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDiscoveredResourceCountsGroupByRegionResponseBody) GoString() string {
	return s.String()
}

func (s *GetDiscoveredResourceCountsGroupByRegionResponseBody) SetDiscoveredResourceCountsSummary(v []*GetDiscoveredResourceCountsGroupByRegionResponseBodyDiscoveredResourceCountsSummary) *GetDiscoveredResourceCountsGroupByRegionResponseBody {
	s.DiscoveredResourceCountsSummary = v
	return s
}

func (s *GetDiscoveredResourceCountsGroupByRegionResponseBody) SetRequestId(v string) *GetDiscoveredResourceCountsGroupByRegionResponseBody {
	s.RequestId = &v
	return s
}

type GetDiscoveredResourceCountsGroupByRegionResponseBodyDiscoveredResourceCountsSummary struct {
	Region        *string `json:"Region,omitempty" xml:"Region,omitempty"`
	ResourceCount *int64  `json:"ResourceCount,omitempty" xml:"ResourceCount,omitempty"`
}

func (s GetDiscoveredResourceCountsGroupByRegionResponseBodyDiscoveredResourceCountsSummary) String() string {
	return tea.Prettify(s)
}

func (s GetDiscoveredResourceCountsGroupByRegionResponseBodyDiscoveredResourceCountsSummary) GoString() string {
	return s.String()
}

func (s *GetDiscoveredResourceCountsGroupByRegionResponseBodyDiscoveredResourceCountsSummary) SetRegion(v string) *GetDiscoveredResourceCountsGroupByRegionResponseBodyDiscoveredResourceCountsSummary {
	s.Region = &v
	return s
}

func (s *GetDiscoveredResourceCountsGroupByRegionResponseBodyDiscoveredResourceCountsSummary) SetResourceCount(v int64) *GetDiscoveredResourceCountsGroupByRegionResponseBodyDiscoveredResourceCountsSummary {
	s.ResourceCount = &v
	return s
}

type GetDiscoveredResourceCountsGroupByRegionResponse struct {
	Headers    map[string]*string                                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetDiscoveredResourceCountsGroupByRegionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetDiscoveredResourceCountsGroupByRegionResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDiscoveredResourceCountsGroupByRegionResponse) GoString() string {
	return s.String()
}

func (s *GetDiscoveredResourceCountsGroupByRegionResponse) SetHeaders(v map[string]*string) *GetDiscoveredResourceCountsGroupByRegionResponse {
	s.Headers = v
	return s
}

func (s *GetDiscoveredResourceCountsGroupByRegionResponse) SetStatusCode(v int32) *GetDiscoveredResourceCountsGroupByRegionResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDiscoveredResourceCountsGroupByRegionResponse) SetBody(v *GetDiscoveredResourceCountsGroupByRegionResponseBody) *GetDiscoveredResourceCountsGroupByRegionResponse {
	s.Body = v
	return s
}

type GetDiscoveredResourceCountsGroupByResourceTypeRequest struct {
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
}

func (s GetDiscoveredResourceCountsGroupByResourceTypeRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDiscoveredResourceCountsGroupByResourceTypeRequest) GoString() string {
	return s.String()
}

func (s *GetDiscoveredResourceCountsGroupByResourceTypeRequest) SetRegion(v string) *GetDiscoveredResourceCountsGroupByResourceTypeRequest {
	s.Region = &v
	return s
}

type GetDiscoveredResourceCountsGroupByResourceTypeResponseBody struct {
	DiscoveredResourceCountsSummary []*GetDiscoveredResourceCountsGroupByResourceTypeResponseBodyDiscoveredResourceCountsSummary `json:"DiscoveredResourceCountsSummary,omitempty" xml:"DiscoveredResourceCountsSummary,omitempty" type:"Repeated"`
	RequestId                       *string                                                                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetDiscoveredResourceCountsGroupByResourceTypeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDiscoveredResourceCountsGroupByResourceTypeResponseBody) GoString() string {
	return s.String()
}

func (s *GetDiscoveredResourceCountsGroupByResourceTypeResponseBody) SetDiscoveredResourceCountsSummary(v []*GetDiscoveredResourceCountsGroupByResourceTypeResponseBodyDiscoveredResourceCountsSummary) *GetDiscoveredResourceCountsGroupByResourceTypeResponseBody {
	s.DiscoveredResourceCountsSummary = v
	return s
}

func (s *GetDiscoveredResourceCountsGroupByResourceTypeResponseBody) SetRequestId(v string) *GetDiscoveredResourceCountsGroupByResourceTypeResponseBody {
	s.RequestId = &v
	return s
}

type GetDiscoveredResourceCountsGroupByResourceTypeResponseBodyDiscoveredResourceCountsSummary struct {
	ResourceCount *int64  `json:"ResourceCount,omitempty" xml:"ResourceCount,omitempty"`
	ResourceType  *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s GetDiscoveredResourceCountsGroupByResourceTypeResponseBodyDiscoveredResourceCountsSummary) String() string {
	return tea.Prettify(s)
}

func (s GetDiscoveredResourceCountsGroupByResourceTypeResponseBodyDiscoveredResourceCountsSummary) GoString() string {
	return s.String()
}

func (s *GetDiscoveredResourceCountsGroupByResourceTypeResponseBodyDiscoveredResourceCountsSummary) SetResourceCount(v int64) *GetDiscoveredResourceCountsGroupByResourceTypeResponseBodyDiscoveredResourceCountsSummary {
	s.ResourceCount = &v
	return s
}

func (s *GetDiscoveredResourceCountsGroupByResourceTypeResponseBodyDiscoveredResourceCountsSummary) SetResourceType(v string) *GetDiscoveredResourceCountsGroupByResourceTypeResponseBodyDiscoveredResourceCountsSummary {
	s.ResourceType = &v
	return s
}

type GetDiscoveredResourceCountsGroupByResourceTypeResponse struct {
	Headers    map[string]*string                                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetDiscoveredResourceCountsGroupByResourceTypeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetDiscoveredResourceCountsGroupByResourceTypeResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDiscoveredResourceCountsGroupByResourceTypeResponse) GoString() string {
	return s.String()
}

func (s *GetDiscoveredResourceCountsGroupByResourceTypeResponse) SetHeaders(v map[string]*string) *GetDiscoveredResourceCountsGroupByResourceTypeResponse {
	s.Headers = v
	return s
}

func (s *GetDiscoveredResourceCountsGroupByResourceTypeResponse) SetStatusCode(v int32) *GetDiscoveredResourceCountsGroupByResourceTypeResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDiscoveredResourceCountsGroupByResourceTypeResponse) SetBody(v *GetDiscoveredResourceCountsGroupByResourceTypeResponseBody) *GetDiscoveredResourceCountsGroupByResourceTypeResponse {
	s.Body = v
	return s
}

type GetManagedRuleRequest struct {
	Identifier *string `json:"Identifier,omitempty" xml:"Identifier,omitempty"`
}

func (s GetManagedRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s GetManagedRuleRequest) GoString() string {
	return s.String()
}

func (s *GetManagedRuleRequest) SetIdentifier(v string) *GetManagedRuleRequest {
	s.Identifier = &v
	return s
}

type GetManagedRuleResponseBody struct {
	ManagedRule *GetManagedRuleResponseBodyManagedRule `json:"ManagedRule,omitempty" xml:"ManagedRule,omitempty" type:"Struct"`
	RequestId   *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetManagedRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetManagedRuleResponseBody) GoString() string {
	return s.String()
}

func (s *GetManagedRuleResponseBody) SetManagedRule(v *GetManagedRuleResponseBodyManagedRule) *GetManagedRuleResponseBody {
	s.ManagedRule = v
	return s
}

func (s *GetManagedRuleResponseBody) SetRequestId(v string) *GetManagedRuleResponseBody {
	s.RequestId = &v
	return s
}

type GetManagedRuleResponseBodyManagedRule struct {
	CompulsoryInputParameterDetails map[string]interface{}                                `json:"CompulsoryInputParameterDetails,omitempty" xml:"CompulsoryInputParameterDetails,omitempty"`
	ConfigRuleName                  *string                                               `json:"ConfigRuleName,omitempty" xml:"ConfigRuleName,omitempty"`
	Description                     *string                                               `json:"Description,omitempty" xml:"Description,omitempty"`
	HelpUrls                        *string                                               `json:"HelpUrls,omitempty" xml:"HelpUrls,omitempty"`
	Identifier                      *string                                               `json:"Identifier,omitempty" xml:"Identifier,omitempty"`
	Labels                          []*string                                             `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	OptionalInputParameterDetails   map[string]interface{}                                `json:"OptionalInputParameterDetails,omitempty" xml:"OptionalInputParameterDetails,omitempty"`
	RiskLevel                       *int32                                                `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
	Scope                           *GetManagedRuleResponseBodyManagedRuleScope           `json:"Scope,omitempty" xml:"Scope,omitempty" type:"Struct"`
	SourceDetails                   []*GetManagedRuleResponseBodyManagedRuleSourceDetails `json:"SourceDetails,omitempty" xml:"SourceDetails,omitempty" type:"Repeated"`
}

func (s GetManagedRuleResponseBodyManagedRule) String() string {
	return tea.Prettify(s)
}

func (s GetManagedRuleResponseBodyManagedRule) GoString() string {
	return s.String()
}

func (s *GetManagedRuleResponseBodyManagedRule) SetCompulsoryInputParameterDetails(v map[string]interface{}) *GetManagedRuleResponseBodyManagedRule {
	s.CompulsoryInputParameterDetails = v
	return s
}

func (s *GetManagedRuleResponseBodyManagedRule) SetConfigRuleName(v string) *GetManagedRuleResponseBodyManagedRule {
	s.ConfigRuleName = &v
	return s
}

func (s *GetManagedRuleResponseBodyManagedRule) SetDescription(v string) *GetManagedRuleResponseBodyManagedRule {
	s.Description = &v
	return s
}

func (s *GetManagedRuleResponseBodyManagedRule) SetHelpUrls(v string) *GetManagedRuleResponseBodyManagedRule {
	s.HelpUrls = &v
	return s
}

func (s *GetManagedRuleResponseBodyManagedRule) SetIdentifier(v string) *GetManagedRuleResponseBodyManagedRule {
	s.Identifier = &v
	return s
}

func (s *GetManagedRuleResponseBodyManagedRule) SetLabels(v []*string) *GetManagedRuleResponseBodyManagedRule {
	s.Labels = v
	return s
}

func (s *GetManagedRuleResponseBodyManagedRule) SetOptionalInputParameterDetails(v map[string]interface{}) *GetManagedRuleResponseBodyManagedRule {
	s.OptionalInputParameterDetails = v
	return s
}

func (s *GetManagedRuleResponseBodyManagedRule) SetRiskLevel(v int32) *GetManagedRuleResponseBodyManagedRule {
	s.RiskLevel = &v
	return s
}

func (s *GetManagedRuleResponseBodyManagedRule) SetScope(v *GetManagedRuleResponseBodyManagedRuleScope) *GetManagedRuleResponseBodyManagedRule {
	s.Scope = v
	return s
}

func (s *GetManagedRuleResponseBodyManagedRule) SetSourceDetails(v []*GetManagedRuleResponseBodyManagedRuleSourceDetails) *GetManagedRuleResponseBodyManagedRule {
	s.SourceDetails = v
	return s
}

type GetManagedRuleResponseBodyManagedRuleScope struct {
	ComplianceResourceTypes []*string `json:"ComplianceResourceTypes,omitempty" xml:"ComplianceResourceTypes,omitempty" type:"Repeated"`
}

func (s GetManagedRuleResponseBodyManagedRuleScope) String() string {
	return tea.Prettify(s)
}

func (s GetManagedRuleResponseBodyManagedRuleScope) GoString() string {
	return s.String()
}

func (s *GetManagedRuleResponseBodyManagedRuleScope) SetComplianceResourceTypes(v []*string) *GetManagedRuleResponseBodyManagedRuleScope {
	s.ComplianceResourceTypes = v
	return s
}

type GetManagedRuleResponseBodyManagedRuleSourceDetails struct {
	MaximumExecutionFrequency *string `json:"MaximumExecutionFrequency,omitempty" xml:"MaximumExecutionFrequency,omitempty"`
	MessageType               *string `json:"MessageType,omitempty" xml:"MessageType,omitempty"`
}

func (s GetManagedRuleResponseBodyManagedRuleSourceDetails) String() string {
	return tea.Prettify(s)
}

func (s GetManagedRuleResponseBodyManagedRuleSourceDetails) GoString() string {
	return s.String()
}

func (s *GetManagedRuleResponseBodyManagedRuleSourceDetails) SetMaximumExecutionFrequency(v string) *GetManagedRuleResponseBodyManagedRuleSourceDetails {
	s.MaximumExecutionFrequency = &v
	return s
}

func (s *GetManagedRuleResponseBodyManagedRuleSourceDetails) SetMessageType(v string) *GetManagedRuleResponseBodyManagedRuleSourceDetails {
	s.MessageType = &v
	return s
}

type GetManagedRuleResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetManagedRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetManagedRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s GetManagedRuleResponse) GoString() string {
	return s.String()
}

func (s *GetManagedRuleResponse) SetHeaders(v map[string]*string) *GetManagedRuleResponse {
	s.Headers = v
	return s
}

func (s *GetManagedRuleResponse) SetStatusCode(v int32) *GetManagedRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *GetManagedRuleResponse) SetBody(v *GetManagedRuleResponseBody) *GetManagedRuleResponse {
	s.Body = v
	return s
}

type GetResourceComplianceByConfigRuleRequest struct {
	ComplianceType *string `json:"ComplianceType,omitempty" xml:"ComplianceType,omitempty"`
	ConfigRuleId   *string `json:"ConfigRuleId,omitempty" xml:"ConfigRuleId,omitempty"`
}

func (s GetResourceComplianceByConfigRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s GetResourceComplianceByConfigRuleRequest) GoString() string {
	return s.String()
}

func (s *GetResourceComplianceByConfigRuleRequest) SetComplianceType(v string) *GetResourceComplianceByConfigRuleRequest {
	s.ComplianceType = &v
	return s
}

func (s *GetResourceComplianceByConfigRuleRequest) SetConfigRuleId(v string) *GetResourceComplianceByConfigRuleRequest {
	s.ConfigRuleId = &v
	return s
}

type GetResourceComplianceByConfigRuleResponseBody struct {
	ComplianceResult *GetResourceComplianceByConfigRuleResponseBodyComplianceResult `json:"ComplianceResult,omitempty" xml:"ComplianceResult,omitempty" type:"Struct"`
	RequestId        *string                                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetResourceComplianceByConfigRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetResourceComplianceByConfigRuleResponseBody) GoString() string {
	return s.String()
}

func (s *GetResourceComplianceByConfigRuleResponseBody) SetComplianceResult(v *GetResourceComplianceByConfigRuleResponseBodyComplianceResult) *GetResourceComplianceByConfigRuleResponseBody {
	s.ComplianceResult = v
	return s
}

func (s *GetResourceComplianceByConfigRuleResponseBody) SetRequestId(v string) *GetResourceComplianceByConfigRuleResponseBody {
	s.RequestId = &v
	return s
}

type GetResourceComplianceByConfigRuleResponseBodyComplianceResult struct {
	Compliances []*GetResourceComplianceByConfigRuleResponseBodyComplianceResultCompliances `json:"Compliances,omitempty" xml:"Compliances,omitempty" type:"Repeated"`
	TotalCount  *int64                                                                      `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s GetResourceComplianceByConfigRuleResponseBodyComplianceResult) String() string {
	return tea.Prettify(s)
}

func (s GetResourceComplianceByConfigRuleResponseBodyComplianceResult) GoString() string {
	return s.String()
}

func (s *GetResourceComplianceByConfigRuleResponseBodyComplianceResult) SetCompliances(v []*GetResourceComplianceByConfigRuleResponseBodyComplianceResultCompliances) *GetResourceComplianceByConfigRuleResponseBodyComplianceResult {
	s.Compliances = v
	return s
}

func (s *GetResourceComplianceByConfigRuleResponseBodyComplianceResult) SetTotalCount(v int64) *GetResourceComplianceByConfigRuleResponseBodyComplianceResult {
	s.TotalCount = &v
	return s
}

type GetResourceComplianceByConfigRuleResponseBodyComplianceResultCompliances struct {
	ComplianceType *string `json:"ComplianceType,omitempty" xml:"ComplianceType,omitempty"`
	Count          *int32  `json:"Count,omitempty" xml:"Count,omitempty"`
}

func (s GetResourceComplianceByConfigRuleResponseBodyComplianceResultCompliances) String() string {
	return tea.Prettify(s)
}

func (s GetResourceComplianceByConfigRuleResponseBodyComplianceResultCompliances) GoString() string {
	return s.String()
}

func (s *GetResourceComplianceByConfigRuleResponseBodyComplianceResultCompliances) SetComplianceType(v string) *GetResourceComplianceByConfigRuleResponseBodyComplianceResultCompliances {
	s.ComplianceType = &v
	return s
}

func (s *GetResourceComplianceByConfigRuleResponseBodyComplianceResultCompliances) SetCount(v int32) *GetResourceComplianceByConfigRuleResponseBodyComplianceResultCompliances {
	s.Count = &v
	return s
}

type GetResourceComplianceByConfigRuleResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetResourceComplianceByConfigRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetResourceComplianceByConfigRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s GetResourceComplianceByConfigRuleResponse) GoString() string {
	return s.String()
}

func (s *GetResourceComplianceByConfigRuleResponse) SetHeaders(v map[string]*string) *GetResourceComplianceByConfigRuleResponse {
	s.Headers = v
	return s
}

func (s *GetResourceComplianceByConfigRuleResponse) SetStatusCode(v int32) *GetResourceComplianceByConfigRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *GetResourceComplianceByConfigRuleResponse) SetBody(v *GetResourceComplianceByConfigRuleResponseBody) *GetResourceComplianceByConfigRuleResponse {
	s.Body = v
	return s
}

type GetResourceComplianceByPackRequest struct {
	CompliancePackId *string `json:"CompliancePackId,omitempty" xml:"CompliancePackId,omitempty"`
}

func (s GetResourceComplianceByPackRequest) String() string {
	return tea.Prettify(s)
}

func (s GetResourceComplianceByPackRequest) GoString() string {
	return s.String()
}

func (s *GetResourceComplianceByPackRequest) SetCompliancePackId(v string) *GetResourceComplianceByPackRequest {
	s.CompliancePackId = &v
	return s
}

type GetResourceComplianceByPackResponseBody struct {
	RequestId                *string                                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResourceComplianceResult *GetResourceComplianceByPackResponseBodyResourceComplianceResult `json:"ResourceComplianceResult,omitempty" xml:"ResourceComplianceResult,omitempty" type:"Struct"`
}

func (s GetResourceComplianceByPackResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetResourceComplianceByPackResponseBody) GoString() string {
	return s.String()
}

func (s *GetResourceComplianceByPackResponseBody) SetRequestId(v string) *GetResourceComplianceByPackResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetResourceComplianceByPackResponseBody) SetResourceComplianceResult(v *GetResourceComplianceByPackResponseBodyResourceComplianceResult) *GetResourceComplianceByPackResponseBody {
	s.ResourceComplianceResult = v
	return s
}

type GetResourceComplianceByPackResponseBodyResourceComplianceResult struct {
	CompliancePackId  *string `json:"CompliancePackId,omitempty" xml:"CompliancePackId,omitempty"`
	NonCompliantCount *int32  `json:"NonCompliantCount,omitempty" xml:"NonCompliantCount,omitempty"`
	TotalCount        *int32  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s GetResourceComplianceByPackResponseBodyResourceComplianceResult) String() string {
	return tea.Prettify(s)
}

func (s GetResourceComplianceByPackResponseBodyResourceComplianceResult) GoString() string {
	return s.String()
}

func (s *GetResourceComplianceByPackResponseBodyResourceComplianceResult) SetCompliancePackId(v string) *GetResourceComplianceByPackResponseBodyResourceComplianceResult {
	s.CompliancePackId = &v
	return s
}

func (s *GetResourceComplianceByPackResponseBodyResourceComplianceResult) SetNonCompliantCount(v int32) *GetResourceComplianceByPackResponseBodyResourceComplianceResult {
	s.NonCompliantCount = &v
	return s
}

func (s *GetResourceComplianceByPackResponseBodyResourceComplianceResult) SetTotalCount(v int32) *GetResourceComplianceByPackResponseBodyResourceComplianceResult {
	s.TotalCount = &v
	return s
}

type GetResourceComplianceByPackResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetResourceComplianceByPackResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetResourceComplianceByPackResponse) String() string {
	return tea.Prettify(s)
}

func (s GetResourceComplianceByPackResponse) GoString() string {
	return s.String()
}

func (s *GetResourceComplianceByPackResponse) SetHeaders(v map[string]*string) *GetResourceComplianceByPackResponse {
	s.Headers = v
	return s
}

func (s *GetResourceComplianceByPackResponse) SetStatusCode(v int32) *GetResourceComplianceByPackResponse {
	s.StatusCode = &v
	return s
}

func (s *GetResourceComplianceByPackResponse) SetBody(v *GetResourceComplianceByPackResponseBody) *GetResourceComplianceByPackResponse {
	s.Body = v
	return s
}

type GetResourceComplianceGroupByRegionRequest struct {
	ConfigRuleIds *string `json:"ConfigRuleIds,omitempty" xml:"ConfigRuleIds,omitempty"`
}

func (s GetResourceComplianceGroupByRegionRequest) String() string {
	return tea.Prettify(s)
}

func (s GetResourceComplianceGroupByRegionRequest) GoString() string {
	return s.String()
}

func (s *GetResourceComplianceGroupByRegionRequest) SetConfigRuleIds(v string) *GetResourceComplianceGroupByRegionRequest {
	s.ConfigRuleIds = &v
	return s
}

type GetResourceComplianceGroupByRegionResponseBody struct {
	ComplianceResult *GetResourceComplianceGroupByRegionResponseBodyComplianceResult `json:"ComplianceResult,omitempty" xml:"ComplianceResult,omitempty" type:"Struct"`
	RequestId        *string                                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetResourceComplianceGroupByRegionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetResourceComplianceGroupByRegionResponseBody) GoString() string {
	return s.String()
}

func (s *GetResourceComplianceGroupByRegionResponseBody) SetComplianceResult(v *GetResourceComplianceGroupByRegionResponseBodyComplianceResult) *GetResourceComplianceGroupByRegionResponseBody {
	s.ComplianceResult = v
	return s
}

func (s *GetResourceComplianceGroupByRegionResponseBody) SetRequestId(v string) *GetResourceComplianceGroupByRegionResponseBody {
	s.RequestId = &v
	return s
}

type GetResourceComplianceGroupByRegionResponseBodyComplianceResult struct {
	ComplianceResultList []*GetResourceComplianceGroupByRegionResponseBodyComplianceResultComplianceResultList `json:"ComplianceResultList,omitempty" xml:"ComplianceResultList,omitempty" type:"Repeated"`
}

func (s GetResourceComplianceGroupByRegionResponseBodyComplianceResult) String() string {
	return tea.Prettify(s)
}

func (s GetResourceComplianceGroupByRegionResponseBodyComplianceResult) GoString() string {
	return s.String()
}

func (s *GetResourceComplianceGroupByRegionResponseBodyComplianceResult) SetComplianceResultList(v []*GetResourceComplianceGroupByRegionResponseBodyComplianceResultComplianceResultList) *GetResourceComplianceGroupByRegionResponseBodyComplianceResult {
	s.ComplianceResultList = v
	return s
}

type GetResourceComplianceGroupByRegionResponseBodyComplianceResultComplianceResultList struct {
	Compliances []*GetResourceComplianceGroupByRegionResponseBodyComplianceResultComplianceResultListCompliances `json:"Compliances,omitempty" xml:"Compliances,omitempty" type:"Repeated"`
	RegionId    *string                                                                                          `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetResourceComplianceGroupByRegionResponseBodyComplianceResultComplianceResultList) String() string {
	return tea.Prettify(s)
}

func (s GetResourceComplianceGroupByRegionResponseBodyComplianceResultComplianceResultList) GoString() string {
	return s.String()
}

func (s *GetResourceComplianceGroupByRegionResponseBodyComplianceResultComplianceResultList) SetCompliances(v []*GetResourceComplianceGroupByRegionResponseBodyComplianceResultComplianceResultListCompliances) *GetResourceComplianceGroupByRegionResponseBodyComplianceResultComplianceResultList {
	s.Compliances = v
	return s
}

func (s *GetResourceComplianceGroupByRegionResponseBodyComplianceResultComplianceResultList) SetRegionId(v string) *GetResourceComplianceGroupByRegionResponseBodyComplianceResultComplianceResultList {
	s.RegionId = &v
	return s
}

type GetResourceComplianceGroupByRegionResponseBodyComplianceResultComplianceResultListCompliances struct {
	ComplianceType *string `json:"ComplianceType,omitempty" xml:"ComplianceType,omitempty"`
	Count          *int64  `json:"Count,omitempty" xml:"Count,omitempty"`
}

func (s GetResourceComplianceGroupByRegionResponseBodyComplianceResultComplianceResultListCompliances) String() string {
	return tea.Prettify(s)
}

func (s GetResourceComplianceGroupByRegionResponseBodyComplianceResultComplianceResultListCompliances) GoString() string {
	return s.String()
}

func (s *GetResourceComplianceGroupByRegionResponseBodyComplianceResultComplianceResultListCompliances) SetComplianceType(v string) *GetResourceComplianceGroupByRegionResponseBodyComplianceResultComplianceResultListCompliances {
	s.ComplianceType = &v
	return s
}

func (s *GetResourceComplianceGroupByRegionResponseBodyComplianceResultComplianceResultListCompliances) SetCount(v int64) *GetResourceComplianceGroupByRegionResponseBodyComplianceResultComplianceResultListCompliances {
	s.Count = &v
	return s
}

type GetResourceComplianceGroupByRegionResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetResourceComplianceGroupByRegionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetResourceComplianceGroupByRegionResponse) String() string {
	return tea.Prettify(s)
}

func (s GetResourceComplianceGroupByRegionResponse) GoString() string {
	return s.String()
}

func (s *GetResourceComplianceGroupByRegionResponse) SetHeaders(v map[string]*string) *GetResourceComplianceGroupByRegionResponse {
	s.Headers = v
	return s
}

func (s *GetResourceComplianceGroupByRegionResponse) SetStatusCode(v int32) *GetResourceComplianceGroupByRegionResponse {
	s.StatusCode = &v
	return s
}

func (s *GetResourceComplianceGroupByRegionResponse) SetBody(v *GetResourceComplianceGroupByRegionResponseBody) *GetResourceComplianceGroupByRegionResponse {
	s.Body = v
	return s
}

type GetResourceComplianceGroupByResourceTypeRequest struct {
	ConfigRuleIds *string `json:"ConfigRuleIds,omitempty" xml:"ConfigRuleIds,omitempty"`
}

func (s GetResourceComplianceGroupByResourceTypeRequest) String() string {
	return tea.Prettify(s)
}

func (s GetResourceComplianceGroupByResourceTypeRequest) GoString() string {
	return s.String()
}

func (s *GetResourceComplianceGroupByResourceTypeRequest) SetConfigRuleIds(v string) *GetResourceComplianceGroupByResourceTypeRequest {
	s.ConfigRuleIds = &v
	return s
}

type GetResourceComplianceGroupByResourceTypeResponseBody struct {
	ComplianceResult *GetResourceComplianceGroupByResourceTypeResponseBodyComplianceResult `json:"ComplianceResult,omitempty" xml:"ComplianceResult,omitempty" type:"Struct"`
	RequestId        *string                                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetResourceComplianceGroupByResourceTypeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetResourceComplianceGroupByResourceTypeResponseBody) GoString() string {
	return s.String()
}

func (s *GetResourceComplianceGroupByResourceTypeResponseBody) SetComplianceResult(v *GetResourceComplianceGroupByResourceTypeResponseBodyComplianceResult) *GetResourceComplianceGroupByResourceTypeResponseBody {
	s.ComplianceResult = v
	return s
}

func (s *GetResourceComplianceGroupByResourceTypeResponseBody) SetRequestId(v string) *GetResourceComplianceGroupByResourceTypeResponseBody {
	s.RequestId = &v
	return s
}

type GetResourceComplianceGroupByResourceTypeResponseBodyComplianceResult struct {
	ComplianceResultList []*GetResourceComplianceGroupByResourceTypeResponseBodyComplianceResultComplianceResultList `json:"ComplianceResultList,omitempty" xml:"ComplianceResultList,omitempty" type:"Repeated"`
}

func (s GetResourceComplianceGroupByResourceTypeResponseBodyComplianceResult) String() string {
	return tea.Prettify(s)
}

func (s GetResourceComplianceGroupByResourceTypeResponseBodyComplianceResult) GoString() string {
	return s.String()
}

func (s *GetResourceComplianceGroupByResourceTypeResponseBodyComplianceResult) SetComplianceResultList(v []*GetResourceComplianceGroupByResourceTypeResponseBodyComplianceResultComplianceResultList) *GetResourceComplianceGroupByResourceTypeResponseBodyComplianceResult {
	s.ComplianceResultList = v
	return s
}

type GetResourceComplianceGroupByResourceTypeResponseBodyComplianceResultComplianceResultList struct {
	Compliances  []*GetResourceComplianceGroupByResourceTypeResponseBodyComplianceResultComplianceResultListCompliances `json:"Compliances,omitempty" xml:"Compliances,omitempty" type:"Repeated"`
	ResourceType *string                                                                                                `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s GetResourceComplianceGroupByResourceTypeResponseBodyComplianceResultComplianceResultList) String() string {
	return tea.Prettify(s)
}

func (s GetResourceComplianceGroupByResourceTypeResponseBodyComplianceResultComplianceResultList) GoString() string {
	return s.String()
}

func (s *GetResourceComplianceGroupByResourceTypeResponseBodyComplianceResultComplianceResultList) SetCompliances(v []*GetResourceComplianceGroupByResourceTypeResponseBodyComplianceResultComplianceResultListCompliances) *GetResourceComplianceGroupByResourceTypeResponseBodyComplianceResultComplianceResultList {
	s.Compliances = v
	return s
}

func (s *GetResourceComplianceGroupByResourceTypeResponseBodyComplianceResultComplianceResultList) SetResourceType(v string) *GetResourceComplianceGroupByResourceTypeResponseBodyComplianceResultComplianceResultList {
	s.ResourceType = &v
	return s
}

type GetResourceComplianceGroupByResourceTypeResponseBodyComplianceResultComplianceResultListCompliances struct {
	ComplianceType *string `json:"ComplianceType,omitempty" xml:"ComplianceType,omitempty"`
	Count          *int64  `json:"Count,omitempty" xml:"Count,omitempty"`
}

func (s GetResourceComplianceGroupByResourceTypeResponseBodyComplianceResultComplianceResultListCompliances) String() string {
	return tea.Prettify(s)
}

func (s GetResourceComplianceGroupByResourceTypeResponseBodyComplianceResultComplianceResultListCompliances) GoString() string {
	return s.String()
}

func (s *GetResourceComplianceGroupByResourceTypeResponseBodyComplianceResultComplianceResultListCompliances) SetComplianceType(v string) *GetResourceComplianceGroupByResourceTypeResponseBodyComplianceResultComplianceResultListCompliances {
	s.ComplianceType = &v
	return s
}

func (s *GetResourceComplianceGroupByResourceTypeResponseBodyComplianceResultComplianceResultListCompliances) SetCount(v int64) *GetResourceComplianceGroupByResourceTypeResponseBodyComplianceResultComplianceResultListCompliances {
	s.Count = &v
	return s
}

type GetResourceComplianceGroupByResourceTypeResponse struct {
	Headers    map[string]*string                                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetResourceComplianceGroupByResourceTypeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetResourceComplianceGroupByResourceTypeResponse) String() string {
	return tea.Prettify(s)
}

func (s GetResourceComplianceGroupByResourceTypeResponse) GoString() string {
	return s.String()
}

func (s *GetResourceComplianceGroupByResourceTypeResponse) SetHeaders(v map[string]*string) *GetResourceComplianceGroupByResourceTypeResponse {
	s.Headers = v
	return s
}

func (s *GetResourceComplianceGroupByResourceTypeResponse) SetStatusCode(v int32) *GetResourceComplianceGroupByResourceTypeResponse {
	s.StatusCode = &v
	return s
}

func (s *GetResourceComplianceGroupByResourceTypeResponse) SetBody(v *GetResourceComplianceGroupByResourceTypeResponseBody) *GetResourceComplianceGroupByResourceTypeResponse {
	s.Body = v
	return s
}

type GetResourceComplianceTimelineRequest struct {
	EndTime      *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	MaxResults   *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken    *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	Region       *string `json:"Region,omitempty" xml:"Region,omitempty"`
	ResourceId   *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	StartTime    *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
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

func (s *GetResourceComplianceTimelineRequest) SetMaxResults(v int32) *GetResourceComplianceTimelineRequest {
	s.MaxResults = &v
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
	RequestId                  *string                                                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	ComplianceList []*GetResourceComplianceTimelineResponseBodyResourceComplianceTimelineComplianceList `json:"ComplianceList,omitempty" xml:"ComplianceList,omitempty" type:"Repeated"`
	MaxResults     *int32                                                                               `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken      *string                                                                              `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
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

func (s *GetResourceComplianceTimelineResponseBodyResourceComplianceTimeline) SetMaxResults(v int32) *GetResourceComplianceTimelineResponseBodyResourceComplianceTimeline {
	s.MaxResults = &v
	return s
}

func (s *GetResourceComplianceTimelineResponseBodyResourceComplianceTimeline) SetNextToken(v string) *GetResourceComplianceTimelineResponseBodyResourceComplianceTimeline {
	s.NextToken = &v
	return s
}

type GetResourceComplianceTimelineResponseBodyResourceComplianceTimelineComplianceList struct {
	AccountId          *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	AvailabilityZone   *string `json:"AvailabilityZone,omitempty" xml:"AvailabilityZone,omitempty"`
	CaptureTime        *int64  `json:"CaptureTime,omitempty" xml:"CaptureTime,omitempty"`
	Configuration      *string `json:"Configuration,omitempty" xml:"Configuration,omitempty"`
	ConfigurationDiff  *string `json:"ConfigurationDiff,omitempty" xml:"ConfigurationDiff,omitempty"`
	Region             *string `json:"Region,omitempty" xml:"Region,omitempty"`
	ResourceCreateTime *int64  `json:"ResourceCreateTime,omitempty" xml:"ResourceCreateTime,omitempty"`
	ResourceId         *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ResourceName       *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	ResourceStatus     *string `json:"ResourceStatus,omitempty" xml:"ResourceStatus,omitempty"`
	ResourceType       *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	Tags               *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
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
	EndTime      *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	MaxResults   *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken    *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	Region       *string `json:"Region,omitempty" xml:"Region,omitempty"`
	ResourceId   *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	StartTime    *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
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

func (s *GetResourceConfigurationTimelineRequest) SetMaxResults(v int32) *GetResourceConfigurationTimelineRequest {
	s.MaxResults = &v
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
	RequestId                     *string                                                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	ConfigurationList []*GetResourceConfigurationTimelineResponseBodyResourceConfigurationTimelineConfigurationList `json:"ConfigurationList,omitempty" xml:"ConfigurationList,omitempty" type:"Repeated"`
	MaxResults        *int32                                                                                        `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken         *string                                                                                       `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
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

func (s *GetResourceConfigurationTimelineResponseBodyResourceConfigurationTimeline) SetMaxResults(v int32) *GetResourceConfigurationTimelineResponseBodyResourceConfigurationTimeline {
	s.MaxResults = &v
	return s
}

func (s *GetResourceConfigurationTimelineResponseBodyResourceConfigurationTimeline) SetNextToken(v string) *GetResourceConfigurationTimelineResponseBodyResourceConfigurationTimeline {
	s.NextToken = &v
	return s
}

type GetResourceConfigurationTimelineResponseBodyResourceConfigurationTimelineConfigurationList struct {
	AccountId          *int64  `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	AvailabilityZone   *string `json:"AvailabilityZone,omitempty" xml:"AvailabilityZone,omitempty"`
	CaptureTime        *string `json:"CaptureTime,omitempty" xml:"CaptureTime,omitempty"`
	ConfigurationDiff  *string `json:"ConfigurationDiff,omitempty" xml:"ConfigurationDiff,omitempty"`
	Region             *string `json:"Region,omitempty" xml:"Region,omitempty"`
	Relationship       *string `json:"Relationship,omitempty" xml:"Relationship,omitempty"`
	RelationshipDiff   *string `json:"RelationshipDiff,omitempty" xml:"RelationshipDiff,omitempty"`
	ResourceCreateTime *string `json:"ResourceCreateTime,omitempty" xml:"ResourceCreateTime,omitempty"`
	ResourceEventType  *string `json:"ResourceEventType,omitempty" xml:"ResourceEventType,omitempty"`
	ResourceId         *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ResourceName       *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	ResourceType       *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	Tags               *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
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

func (s *GetResourceConfigurationTimelineResponseBodyResourceConfigurationTimelineConfigurationList) SetRelationship(v string) *GetResourceConfigurationTimelineResponseBodyResourceConfigurationTimelineConfigurationList {
	s.Relationship = &v
	return s
}

func (s *GetResourceConfigurationTimelineResponseBodyResourceConfigurationTimelineConfigurationList) SetRelationshipDiff(v string) *GetResourceConfigurationTimelineResponseBodyResourceConfigurationTimelineConfigurationList {
	s.RelationshipDiff = &v
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

type IgnoreAggregateEvaluationResultsRequest struct {
	AggregatorId *string                                             `json:"AggregatorId,omitempty" xml:"AggregatorId,omitempty"`
	ConfigRuleId *string                                             `json:"ConfigRuleId,omitempty" xml:"ConfigRuleId,omitempty"`
	IgnoreDate   *string                                             `json:"IgnoreDate,omitempty" xml:"IgnoreDate,omitempty"`
	Reason       *string                                             `json:"Reason,omitempty" xml:"Reason,omitempty"`
	Resources    []*IgnoreAggregateEvaluationResultsRequestResources `json:"Resources,omitempty" xml:"Resources,omitempty" type:"Repeated"`
}

func (s IgnoreAggregateEvaluationResultsRequest) String() string {
	return tea.Prettify(s)
}

func (s IgnoreAggregateEvaluationResultsRequest) GoString() string {
	return s.String()
}

func (s *IgnoreAggregateEvaluationResultsRequest) SetAggregatorId(v string) *IgnoreAggregateEvaluationResultsRequest {
	s.AggregatorId = &v
	return s
}

func (s *IgnoreAggregateEvaluationResultsRequest) SetConfigRuleId(v string) *IgnoreAggregateEvaluationResultsRequest {
	s.ConfigRuleId = &v
	return s
}

func (s *IgnoreAggregateEvaluationResultsRequest) SetIgnoreDate(v string) *IgnoreAggregateEvaluationResultsRequest {
	s.IgnoreDate = &v
	return s
}

func (s *IgnoreAggregateEvaluationResultsRequest) SetReason(v string) *IgnoreAggregateEvaluationResultsRequest {
	s.Reason = &v
	return s
}

func (s *IgnoreAggregateEvaluationResultsRequest) SetResources(v []*IgnoreAggregateEvaluationResultsRequestResources) *IgnoreAggregateEvaluationResultsRequest {
	s.Resources = v
	return s
}

type IgnoreAggregateEvaluationResultsRequestResources struct {
	Region            *string `json:"Region,omitempty" xml:"Region,omitempty"`
	ResourceAccountId *int64  `json:"ResourceAccountId,omitempty" xml:"ResourceAccountId,omitempty"`
	ResourceId        *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ResourceType      *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s IgnoreAggregateEvaluationResultsRequestResources) String() string {
	return tea.Prettify(s)
}

func (s IgnoreAggregateEvaluationResultsRequestResources) GoString() string {
	return s.String()
}

func (s *IgnoreAggregateEvaluationResultsRequestResources) SetRegion(v string) *IgnoreAggregateEvaluationResultsRequestResources {
	s.Region = &v
	return s
}

func (s *IgnoreAggregateEvaluationResultsRequestResources) SetResourceAccountId(v int64) *IgnoreAggregateEvaluationResultsRequestResources {
	s.ResourceAccountId = &v
	return s
}

func (s *IgnoreAggregateEvaluationResultsRequestResources) SetResourceId(v string) *IgnoreAggregateEvaluationResultsRequestResources {
	s.ResourceId = &v
	return s
}

func (s *IgnoreAggregateEvaluationResultsRequestResources) SetResourceType(v string) *IgnoreAggregateEvaluationResultsRequestResources {
	s.ResourceType = &v
	return s
}

type IgnoreAggregateEvaluationResultsShrinkRequest struct {
	AggregatorId    *string `json:"AggregatorId,omitempty" xml:"AggregatorId,omitempty"`
	ConfigRuleId    *string `json:"ConfigRuleId,omitempty" xml:"ConfigRuleId,omitempty"`
	IgnoreDate      *string `json:"IgnoreDate,omitempty" xml:"IgnoreDate,omitempty"`
	Reason          *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	ResourcesShrink *string `json:"Resources,omitempty" xml:"Resources,omitempty"`
}

func (s IgnoreAggregateEvaluationResultsShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s IgnoreAggregateEvaluationResultsShrinkRequest) GoString() string {
	return s.String()
}

func (s *IgnoreAggregateEvaluationResultsShrinkRequest) SetAggregatorId(v string) *IgnoreAggregateEvaluationResultsShrinkRequest {
	s.AggregatorId = &v
	return s
}

func (s *IgnoreAggregateEvaluationResultsShrinkRequest) SetConfigRuleId(v string) *IgnoreAggregateEvaluationResultsShrinkRequest {
	s.ConfigRuleId = &v
	return s
}

func (s *IgnoreAggregateEvaluationResultsShrinkRequest) SetIgnoreDate(v string) *IgnoreAggregateEvaluationResultsShrinkRequest {
	s.IgnoreDate = &v
	return s
}

func (s *IgnoreAggregateEvaluationResultsShrinkRequest) SetReason(v string) *IgnoreAggregateEvaluationResultsShrinkRequest {
	s.Reason = &v
	return s
}

func (s *IgnoreAggregateEvaluationResultsShrinkRequest) SetResourcesShrink(v string) *IgnoreAggregateEvaluationResultsShrinkRequest {
	s.ResourcesShrink = &v
	return s
}

type IgnoreAggregateEvaluationResultsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s IgnoreAggregateEvaluationResultsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s IgnoreAggregateEvaluationResultsResponseBody) GoString() string {
	return s.String()
}

func (s *IgnoreAggregateEvaluationResultsResponseBody) SetRequestId(v string) *IgnoreAggregateEvaluationResultsResponseBody {
	s.RequestId = &v
	return s
}

type IgnoreAggregateEvaluationResultsResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *IgnoreAggregateEvaluationResultsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s IgnoreAggregateEvaluationResultsResponse) String() string {
	return tea.Prettify(s)
}

func (s IgnoreAggregateEvaluationResultsResponse) GoString() string {
	return s.String()
}

func (s *IgnoreAggregateEvaluationResultsResponse) SetHeaders(v map[string]*string) *IgnoreAggregateEvaluationResultsResponse {
	s.Headers = v
	return s
}

func (s *IgnoreAggregateEvaluationResultsResponse) SetStatusCode(v int32) *IgnoreAggregateEvaluationResultsResponse {
	s.StatusCode = &v
	return s
}

func (s *IgnoreAggregateEvaluationResultsResponse) SetBody(v *IgnoreAggregateEvaluationResultsResponseBody) *IgnoreAggregateEvaluationResultsResponse {
	s.Body = v
	return s
}

type IgnoreEvaluationResultsRequest struct {
	ConfigRuleId *string                                    `json:"ConfigRuleId,omitempty" xml:"ConfigRuleId,omitempty"`
	IgnoreDate   *string                                    `json:"IgnoreDate,omitempty" xml:"IgnoreDate,omitempty"`
	Reason       *string                                    `json:"Reason,omitempty" xml:"Reason,omitempty"`
	Resources    []*IgnoreEvaluationResultsRequestResources `json:"Resources,omitempty" xml:"Resources,omitempty" type:"Repeated"`
}

func (s IgnoreEvaluationResultsRequest) String() string {
	return tea.Prettify(s)
}

func (s IgnoreEvaluationResultsRequest) GoString() string {
	return s.String()
}

func (s *IgnoreEvaluationResultsRequest) SetConfigRuleId(v string) *IgnoreEvaluationResultsRequest {
	s.ConfigRuleId = &v
	return s
}

func (s *IgnoreEvaluationResultsRequest) SetIgnoreDate(v string) *IgnoreEvaluationResultsRequest {
	s.IgnoreDate = &v
	return s
}

func (s *IgnoreEvaluationResultsRequest) SetReason(v string) *IgnoreEvaluationResultsRequest {
	s.Reason = &v
	return s
}

func (s *IgnoreEvaluationResultsRequest) SetResources(v []*IgnoreEvaluationResultsRequestResources) *IgnoreEvaluationResultsRequest {
	s.Resources = v
	return s
}

type IgnoreEvaluationResultsRequestResources struct {
	Region            *string `json:"Region,omitempty" xml:"Region,omitempty"`
	ResourceAccountId *int64  `json:"ResourceAccountId,omitempty" xml:"ResourceAccountId,omitempty"`
	ResourceId        *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ResourceType      *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s IgnoreEvaluationResultsRequestResources) String() string {
	return tea.Prettify(s)
}

func (s IgnoreEvaluationResultsRequestResources) GoString() string {
	return s.String()
}

func (s *IgnoreEvaluationResultsRequestResources) SetRegion(v string) *IgnoreEvaluationResultsRequestResources {
	s.Region = &v
	return s
}

func (s *IgnoreEvaluationResultsRequestResources) SetResourceAccountId(v int64) *IgnoreEvaluationResultsRequestResources {
	s.ResourceAccountId = &v
	return s
}

func (s *IgnoreEvaluationResultsRequestResources) SetResourceId(v string) *IgnoreEvaluationResultsRequestResources {
	s.ResourceId = &v
	return s
}

func (s *IgnoreEvaluationResultsRequestResources) SetResourceType(v string) *IgnoreEvaluationResultsRequestResources {
	s.ResourceType = &v
	return s
}

type IgnoreEvaluationResultsShrinkRequest struct {
	ConfigRuleId    *string `json:"ConfigRuleId,omitempty" xml:"ConfigRuleId,omitempty"`
	IgnoreDate      *string `json:"IgnoreDate,omitempty" xml:"IgnoreDate,omitempty"`
	Reason          *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	ResourcesShrink *string `json:"Resources,omitempty" xml:"Resources,omitempty"`
}

func (s IgnoreEvaluationResultsShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s IgnoreEvaluationResultsShrinkRequest) GoString() string {
	return s.String()
}

func (s *IgnoreEvaluationResultsShrinkRequest) SetConfigRuleId(v string) *IgnoreEvaluationResultsShrinkRequest {
	s.ConfigRuleId = &v
	return s
}

func (s *IgnoreEvaluationResultsShrinkRequest) SetIgnoreDate(v string) *IgnoreEvaluationResultsShrinkRequest {
	s.IgnoreDate = &v
	return s
}

func (s *IgnoreEvaluationResultsShrinkRequest) SetReason(v string) *IgnoreEvaluationResultsShrinkRequest {
	s.Reason = &v
	return s
}

func (s *IgnoreEvaluationResultsShrinkRequest) SetResourcesShrink(v string) *IgnoreEvaluationResultsShrinkRequest {
	s.ResourcesShrink = &v
	return s
}

type IgnoreEvaluationResultsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s IgnoreEvaluationResultsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s IgnoreEvaluationResultsResponseBody) GoString() string {
	return s.String()
}

func (s *IgnoreEvaluationResultsResponseBody) SetRequestId(v string) *IgnoreEvaluationResultsResponseBody {
	s.RequestId = &v
	return s
}

type IgnoreEvaluationResultsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *IgnoreEvaluationResultsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s IgnoreEvaluationResultsResponse) String() string {
	return tea.Prettify(s)
}

func (s IgnoreEvaluationResultsResponse) GoString() string {
	return s.String()
}

func (s *IgnoreEvaluationResultsResponse) SetHeaders(v map[string]*string) *IgnoreEvaluationResultsResponse {
	s.Headers = v
	return s
}

func (s *IgnoreEvaluationResultsResponse) SetStatusCode(v int32) *IgnoreEvaluationResultsResponse {
	s.StatusCode = &v
	return s
}

func (s *IgnoreEvaluationResultsResponse) SetBody(v *IgnoreEvaluationResultsResponseBody) *IgnoreEvaluationResultsResponse {
	s.Body = v
	return s
}

type ListAggregateCompliancePacksRequest struct {
	AggregatorId *string `json:"AggregatorId,omitempty" xml:"AggregatorId,omitempty"`
	PageNumber   *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize     *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListAggregateCompliancePacksRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAggregateCompliancePacksRequest) GoString() string {
	return s.String()
}

func (s *ListAggregateCompliancePacksRequest) SetAggregatorId(v string) *ListAggregateCompliancePacksRequest {
	s.AggregatorId = &v
	return s
}

func (s *ListAggregateCompliancePacksRequest) SetPageNumber(v int32) *ListAggregateCompliancePacksRequest {
	s.PageNumber = &v
	return s
}

func (s *ListAggregateCompliancePacksRequest) SetPageSize(v int32) *ListAggregateCompliancePacksRequest {
	s.PageSize = &v
	return s
}

func (s *ListAggregateCompliancePacksRequest) SetStatus(v string) *ListAggregateCompliancePacksRequest {
	s.Status = &v
	return s
}

type ListAggregateCompliancePacksResponseBody struct {
	CompliancePacksResult *ListAggregateCompliancePacksResponseBodyCompliancePacksResult `json:"CompliancePacksResult,omitempty" xml:"CompliancePacksResult,omitempty" type:"Struct"`
	RequestId             *string                                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListAggregateCompliancePacksResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAggregateCompliancePacksResponseBody) GoString() string {
	return s.String()
}

func (s *ListAggregateCompliancePacksResponseBody) SetCompliancePacksResult(v *ListAggregateCompliancePacksResponseBodyCompliancePacksResult) *ListAggregateCompliancePacksResponseBody {
	s.CompliancePacksResult = v
	return s
}

func (s *ListAggregateCompliancePacksResponseBody) SetRequestId(v string) *ListAggregateCompliancePacksResponseBody {
	s.RequestId = &v
	return s
}

type ListAggregateCompliancePacksResponseBodyCompliancePacksResult struct {
	CompliancePacks []*ListAggregateCompliancePacksResponseBodyCompliancePacksResultCompliancePacks `json:"CompliancePacks,omitempty" xml:"CompliancePacks,omitempty" type:"Repeated"`
	PageNumber      *int32                                                                          `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize        *int32                                                                          `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount      *int64                                                                          `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListAggregateCompliancePacksResponseBodyCompliancePacksResult) String() string {
	return tea.Prettify(s)
}

func (s ListAggregateCompliancePacksResponseBodyCompliancePacksResult) GoString() string {
	return s.String()
}

func (s *ListAggregateCompliancePacksResponseBodyCompliancePacksResult) SetCompliancePacks(v []*ListAggregateCompliancePacksResponseBodyCompliancePacksResultCompliancePacks) *ListAggregateCompliancePacksResponseBodyCompliancePacksResult {
	s.CompliancePacks = v
	return s
}

func (s *ListAggregateCompliancePacksResponseBodyCompliancePacksResult) SetPageNumber(v int32) *ListAggregateCompliancePacksResponseBodyCompliancePacksResult {
	s.PageNumber = &v
	return s
}

func (s *ListAggregateCompliancePacksResponseBodyCompliancePacksResult) SetPageSize(v int32) *ListAggregateCompliancePacksResponseBodyCompliancePacksResult {
	s.PageSize = &v
	return s
}

func (s *ListAggregateCompliancePacksResponseBodyCompliancePacksResult) SetTotalCount(v int64) *ListAggregateCompliancePacksResponseBodyCompliancePacksResult {
	s.TotalCount = &v
	return s
}

type ListAggregateCompliancePacksResponseBodyCompliancePacksResultCompliancePacks struct {
	AccountId                *int64  `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	AggregatorId             *string `json:"AggregatorId,omitempty" xml:"AggregatorId,omitempty"`
	CompliancePackId         *string `json:"CompliancePackId,omitempty" xml:"CompliancePackId,omitempty"`
	CompliancePackName       *string `json:"CompliancePackName,omitempty" xml:"CompliancePackName,omitempty"`
	CompliancePackTemplateId *string `json:"CompliancePackTemplateId,omitempty" xml:"CompliancePackTemplateId,omitempty"`
	CreateTimestamp          *int64  `json:"CreateTimestamp,omitempty" xml:"CreateTimestamp,omitempty"`
	Description              *string `json:"Description,omitempty" xml:"Description,omitempty"`
	RiskLevel                *int32  `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
	Status                   *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListAggregateCompliancePacksResponseBodyCompliancePacksResultCompliancePacks) String() string {
	return tea.Prettify(s)
}

func (s ListAggregateCompliancePacksResponseBodyCompliancePacksResultCompliancePacks) GoString() string {
	return s.String()
}

func (s *ListAggregateCompliancePacksResponseBodyCompliancePacksResultCompliancePacks) SetAccountId(v int64) *ListAggregateCompliancePacksResponseBodyCompliancePacksResultCompliancePacks {
	s.AccountId = &v
	return s
}

func (s *ListAggregateCompliancePacksResponseBodyCompliancePacksResultCompliancePacks) SetAggregatorId(v string) *ListAggregateCompliancePacksResponseBodyCompliancePacksResultCompliancePacks {
	s.AggregatorId = &v
	return s
}

func (s *ListAggregateCompliancePacksResponseBodyCompliancePacksResultCompliancePacks) SetCompliancePackId(v string) *ListAggregateCompliancePacksResponseBodyCompliancePacksResultCompliancePacks {
	s.CompliancePackId = &v
	return s
}

func (s *ListAggregateCompliancePacksResponseBodyCompliancePacksResultCompliancePacks) SetCompliancePackName(v string) *ListAggregateCompliancePacksResponseBodyCompliancePacksResultCompliancePacks {
	s.CompliancePackName = &v
	return s
}

func (s *ListAggregateCompliancePacksResponseBodyCompliancePacksResultCompliancePacks) SetCompliancePackTemplateId(v string) *ListAggregateCompliancePacksResponseBodyCompliancePacksResultCompliancePacks {
	s.CompliancePackTemplateId = &v
	return s
}

func (s *ListAggregateCompliancePacksResponseBodyCompliancePacksResultCompliancePacks) SetCreateTimestamp(v int64) *ListAggregateCompliancePacksResponseBodyCompliancePacksResultCompliancePacks {
	s.CreateTimestamp = &v
	return s
}

func (s *ListAggregateCompliancePacksResponseBodyCompliancePacksResultCompliancePacks) SetDescription(v string) *ListAggregateCompliancePacksResponseBodyCompliancePacksResultCompliancePacks {
	s.Description = &v
	return s
}

func (s *ListAggregateCompliancePacksResponseBodyCompliancePacksResultCompliancePacks) SetRiskLevel(v int32) *ListAggregateCompliancePacksResponseBodyCompliancePacksResultCompliancePacks {
	s.RiskLevel = &v
	return s
}

func (s *ListAggregateCompliancePacksResponseBodyCompliancePacksResultCompliancePacks) SetStatus(v string) *ListAggregateCompliancePacksResponseBodyCompliancePacksResultCompliancePacks {
	s.Status = &v
	return s
}

type ListAggregateCompliancePacksResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListAggregateCompliancePacksResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListAggregateCompliancePacksResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAggregateCompliancePacksResponse) GoString() string {
	return s.String()
}

func (s *ListAggregateCompliancePacksResponse) SetHeaders(v map[string]*string) *ListAggregateCompliancePacksResponse {
	s.Headers = v
	return s
}

func (s *ListAggregateCompliancePacksResponse) SetStatusCode(v int32) *ListAggregateCompliancePacksResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAggregateCompliancePacksResponse) SetBody(v *ListAggregateCompliancePacksResponseBody) *ListAggregateCompliancePacksResponse {
	s.Body = v
	return s
}

type ListAggregateConfigDeliveryChannelsRequest struct {
	AggregatorId       *string `json:"AggregatorId,omitempty" xml:"AggregatorId,omitempty"`
	DeliveryChannelIds *string `json:"DeliveryChannelIds,omitempty" xml:"DeliveryChannelIds,omitempty"`
}

func (s ListAggregateConfigDeliveryChannelsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAggregateConfigDeliveryChannelsRequest) GoString() string {
	return s.String()
}

func (s *ListAggregateConfigDeliveryChannelsRequest) SetAggregatorId(v string) *ListAggregateConfigDeliveryChannelsRequest {
	s.AggregatorId = &v
	return s
}

func (s *ListAggregateConfigDeliveryChannelsRequest) SetDeliveryChannelIds(v string) *ListAggregateConfigDeliveryChannelsRequest {
	s.DeliveryChannelIds = &v
	return s
}

type ListAggregateConfigDeliveryChannelsResponseBody struct {
	DeliveryChannels []*ListAggregateConfigDeliveryChannelsResponseBodyDeliveryChannels `json:"DeliveryChannels,omitempty" xml:"DeliveryChannels,omitempty" type:"Repeated"`
	RequestId        *string                                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListAggregateConfigDeliveryChannelsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAggregateConfigDeliveryChannelsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAggregateConfigDeliveryChannelsResponseBody) SetDeliveryChannels(v []*ListAggregateConfigDeliveryChannelsResponseBodyDeliveryChannels) *ListAggregateConfigDeliveryChannelsResponseBody {
	s.DeliveryChannels = v
	return s
}

func (s *ListAggregateConfigDeliveryChannelsResponseBody) SetRequestId(v string) *ListAggregateConfigDeliveryChannelsResponseBody {
	s.RequestId = &v
	return s
}

type ListAggregateConfigDeliveryChannelsResponseBodyDeliveryChannels struct {
	AccountId                           *int64  `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	AggregatorId                        *string `json:"AggregatorId,omitempty" xml:"AggregatorId,omitempty"`
	ConfigurationItemChangeNotification *bool   `json:"ConfigurationItemChangeNotification,omitempty" xml:"ConfigurationItemChangeNotification,omitempty"`
	ConfigurationSnapshot               *bool   `json:"ConfigurationSnapshot,omitempty" xml:"ConfigurationSnapshot,omitempty"`
	DeliveryChannelAssumeRoleArn        *string `json:"DeliveryChannelAssumeRoleArn,omitempty" xml:"DeliveryChannelAssumeRoleArn,omitempty"`
	DeliveryChannelCondition            *string `json:"DeliveryChannelCondition,omitempty" xml:"DeliveryChannelCondition,omitempty"`
	DeliveryChannelId                   *string `json:"DeliveryChannelId,omitempty" xml:"DeliveryChannelId,omitempty"`
	DeliveryChannelName                 *string `json:"DeliveryChannelName,omitempty" xml:"DeliveryChannelName,omitempty"`
	DeliveryChannelTargetArn            *string `json:"DeliveryChannelTargetArn,omitempty" xml:"DeliveryChannelTargetArn,omitempty"`
	DeliveryChannelType                 *string `json:"DeliveryChannelType,omitempty" xml:"DeliveryChannelType,omitempty"`
	Description                         *string `json:"Description,omitempty" xml:"Description,omitempty"`
	NonCompliantNotification            *bool   `json:"NonCompliantNotification,omitempty" xml:"NonCompliantNotification,omitempty"`
	OversizedDataOSSTargetArn           *string `json:"OversizedDataOSSTargetArn,omitempty" xml:"OversizedDataOSSTargetArn,omitempty"`
	Status                              *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListAggregateConfigDeliveryChannelsResponseBodyDeliveryChannels) String() string {
	return tea.Prettify(s)
}

func (s ListAggregateConfigDeliveryChannelsResponseBodyDeliveryChannels) GoString() string {
	return s.String()
}

func (s *ListAggregateConfigDeliveryChannelsResponseBodyDeliveryChannels) SetAccountId(v int64) *ListAggregateConfigDeliveryChannelsResponseBodyDeliveryChannels {
	s.AccountId = &v
	return s
}

func (s *ListAggregateConfigDeliveryChannelsResponseBodyDeliveryChannels) SetAggregatorId(v string) *ListAggregateConfigDeliveryChannelsResponseBodyDeliveryChannels {
	s.AggregatorId = &v
	return s
}

func (s *ListAggregateConfigDeliveryChannelsResponseBodyDeliveryChannels) SetConfigurationItemChangeNotification(v bool) *ListAggregateConfigDeliveryChannelsResponseBodyDeliveryChannels {
	s.ConfigurationItemChangeNotification = &v
	return s
}

func (s *ListAggregateConfigDeliveryChannelsResponseBodyDeliveryChannels) SetConfigurationSnapshot(v bool) *ListAggregateConfigDeliveryChannelsResponseBodyDeliveryChannels {
	s.ConfigurationSnapshot = &v
	return s
}

func (s *ListAggregateConfigDeliveryChannelsResponseBodyDeliveryChannels) SetDeliveryChannelAssumeRoleArn(v string) *ListAggregateConfigDeliveryChannelsResponseBodyDeliveryChannels {
	s.DeliveryChannelAssumeRoleArn = &v
	return s
}

func (s *ListAggregateConfigDeliveryChannelsResponseBodyDeliveryChannels) SetDeliveryChannelCondition(v string) *ListAggregateConfigDeliveryChannelsResponseBodyDeliveryChannels {
	s.DeliveryChannelCondition = &v
	return s
}

func (s *ListAggregateConfigDeliveryChannelsResponseBodyDeliveryChannels) SetDeliveryChannelId(v string) *ListAggregateConfigDeliveryChannelsResponseBodyDeliveryChannels {
	s.DeliveryChannelId = &v
	return s
}

func (s *ListAggregateConfigDeliveryChannelsResponseBodyDeliveryChannels) SetDeliveryChannelName(v string) *ListAggregateConfigDeliveryChannelsResponseBodyDeliveryChannels {
	s.DeliveryChannelName = &v
	return s
}

func (s *ListAggregateConfigDeliveryChannelsResponseBodyDeliveryChannels) SetDeliveryChannelTargetArn(v string) *ListAggregateConfigDeliveryChannelsResponseBodyDeliveryChannels {
	s.DeliveryChannelTargetArn = &v
	return s
}

func (s *ListAggregateConfigDeliveryChannelsResponseBodyDeliveryChannels) SetDeliveryChannelType(v string) *ListAggregateConfigDeliveryChannelsResponseBodyDeliveryChannels {
	s.DeliveryChannelType = &v
	return s
}

func (s *ListAggregateConfigDeliveryChannelsResponseBodyDeliveryChannels) SetDescription(v string) *ListAggregateConfigDeliveryChannelsResponseBodyDeliveryChannels {
	s.Description = &v
	return s
}

func (s *ListAggregateConfigDeliveryChannelsResponseBodyDeliveryChannels) SetNonCompliantNotification(v bool) *ListAggregateConfigDeliveryChannelsResponseBodyDeliveryChannels {
	s.NonCompliantNotification = &v
	return s
}

func (s *ListAggregateConfigDeliveryChannelsResponseBodyDeliveryChannels) SetOversizedDataOSSTargetArn(v string) *ListAggregateConfigDeliveryChannelsResponseBodyDeliveryChannels {
	s.OversizedDataOSSTargetArn = &v
	return s
}

func (s *ListAggregateConfigDeliveryChannelsResponseBodyDeliveryChannels) SetStatus(v int32) *ListAggregateConfigDeliveryChannelsResponseBodyDeliveryChannels {
	s.Status = &v
	return s
}

type ListAggregateConfigDeliveryChannelsResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListAggregateConfigDeliveryChannelsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListAggregateConfigDeliveryChannelsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAggregateConfigDeliveryChannelsResponse) GoString() string {
	return s.String()
}

func (s *ListAggregateConfigDeliveryChannelsResponse) SetHeaders(v map[string]*string) *ListAggregateConfigDeliveryChannelsResponse {
	s.Headers = v
	return s
}

func (s *ListAggregateConfigDeliveryChannelsResponse) SetStatusCode(v int32) *ListAggregateConfigDeliveryChannelsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAggregateConfigDeliveryChannelsResponse) SetBody(v *ListAggregateConfigDeliveryChannelsResponseBody) *ListAggregateConfigDeliveryChannelsResponse {
	s.Body = v
	return s
}

type ListAggregateConfigRuleEvaluationResultsRequest struct {
	AggregatorId     *string `json:"AggregatorId,omitempty" xml:"AggregatorId,omitempty"`
	CompliancePackId *string `json:"CompliancePackId,omitempty" xml:"CompliancePackId,omitempty"`
	ComplianceType   *string `json:"ComplianceType,omitempty" xml:"ComplianceType,omitempty"`
	ConfigRuleId     *string `json:"ConfigRuleId,omitempty" xml:"ConfigRuleId,omitempty"`
	MaxResults       *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken        *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	ResourceOwnerId  *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ListAggregateConfigRuleEvaluationResultsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAggregateConfigRuleEvaluationResultsRequest) GoString() string {
	return s.String()
}

func (s *ListAggregateConfigRuleEvaluationResultsRequest) SetAggregatorId(v string) *ListAggregateConfigRuleEvaluationResultsRequest {
	s.AggregatorId = &v
	return s
}

func (s *ListAggregateConfigRuleEvaluationResultsRequest) SetCompliancePackId(v string) *ListAggregateConfigRuleEvaluationResultsRequest {
	s.CompliancePackId = &v
	return s
}

func (s *ListAggregateConfigRuleEvaluationResultsRequest) SetComplianceType(v string) *ListAggregateConfigRuleEvaluationResultsRequest {
	s.ComplianceType = &v
	return s
}

func (s *ListAggregateConfigRuleEvaluationResultsRequest) SetConfigRuleId(v string) *ListAggregateConfigRuleEvaluationResultsRequest {
	s.ConfigRuleId = &v
	return s
}

func (s *ListAggregateConfigRuleEvaluationResultsRequest) SetMaxResults(v int32) *ListAggregateConfigRuleEvaluationResultsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListAggregateConfigRuleEvaluationResultsRequest) SetNextToken(v string) *ListAggregateConfigRuleEvaluationResultsRequest {
	s.NextToken = &v
	return s
}

func (s *ListAggregateConfigRuleEvaluationResultsRequest) SetResourceOwnerId(v int64) *ListAggregateConfigRuleEvaluationResultsRequest {
	s.ResourceOwnerId = &v
	return s
}

type ListAggregateConfigRuleEvaluationResultsResponseBody struct {
	EvaluationResults *ListAggregateConfigRuleEvaluationResultsResponseBodyEvaluationResults `json:"EvaluationResults,omitempty" xml:"EvaluationResults,omitempty" type:"Struct"`
	RequestId         *string                                                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListAggregateConfigRuleEvaluationResultsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAggregateConfigRuleEvaluationResultsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAggregateConfigRuleEvaluationResultsResponseBody) SetEvaluationResults(v *ListAggregateConfigRuleEvaluationResultsResponseBodyEvaluationResults) *ListAggregateConfigRuleEvaluationResultsResponseBody {
	s.EvaluationResults = v
	return s
}

func (s *ListAggregateConfigRuleEvaluationResultsResponseBody) SetRequestId(v string) *ListAggregateConfigRuleEvaluationResultsResponseBody {
	s.RequestId = &v
	return s
}

type ListAggregateConfigRuleEvaluationResultsResponseBodyEvaluationResults struct {
	EvaluationResultList []*ListAggregateConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList `json:"EvaluationResultList,omitempty" xml:"EvaluationResultList,omitempty" type:"Repeated"`
	MaxResults           *int32                                                                                       `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken            *string                                                                                      `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListAggregateConfigRuleEvaluationResultsResponseBodyEvaluationResults) String() string {
	return tea.Prettify(s)
}

func (s ListAggregateConfigRuleEvaluationResultsResponseBodyEvaluationResults) GoString() string {
	return s.String()
}

func (s *ListAggregateConfigRuleEvaluationResultsResponseBodyEvaluationResults) SetEvaluationResultList(v []*ListAggregateConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList) *ListAggregateConfigRuleEvaluationResultsResponseBodyEvaluationResults {
	s.EvaluationResultList = v
	return s
}

func (s *ListAggregateConfigRuleEvaluationResultsResponseBodyEvaluationResults) SetMaxResults(v int32) *ListAggregateConfigRuleEvaluationResultsResponseBodyEvaluationResults {
	s.MaxResults = &v
	return s
}

func (s *ListAggregateConfigRuleEvaluationResultsResponseBodyEvaluationResults) SetNextToken(v string) *ListAggregateConfigRuleEvaluationResultsResponseBodyEvaluationResults {
	s.NextToken = &v
	return s
}

type ListAggregateConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList struct {
	Annotation                 *string                                                                                                              `json:"Annotation,omitempty" xml:"Annotation,omitempty"`
	ComplianceType             *string                                                                                                              `json:"ComplianceType,omitempty" xml:"ComplianceType,omitempty"`
	ConfigRuleInvokedTimestamp *int64                                                                                                               `json:"ConfigRuleInvokedTimestamp,omitempty" xml:"ConfigRuleInvokedTimestamp,omitempty"`
	EvaluationResultIdentifier *ListAggregateConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifier `json:"EvaluationResultIdentifier,omitempty" xml:"EvaluationResultIdentifier,omitempty" type:"Struct"`
	InvokingEventMessageType   *string                                                                                                              `json:"InvokingEventMessageType,omitempty" xml:"InvokingEventMessageType,omitempty"`
	RemediationEnabled         *bool                                                                                                                `json:"RemediationEnabled,omitempty" xml:"RemediationEnabled,omitempty"`
	ResultRecordedTimestamp    *int64                                                                                                               `json:"ResultRecordedTimestamp,omitempty" xml:"ResultRecordedTimestamp,omitempty"`
	RiskLevel                  *int32                                                                                                               `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
}

func (s ListAggregateConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList) String() string {
	return tea.Prettify(s)
}

func (s ListAggregateConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList) GoString() string {
	return s.String()
}

func (s *ListAggregateConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList) SetAnnotation(v string) *ListAggregateConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList {
	s.Annotation = &v
	return s
}

func (s *ListAggregateConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList) SetComplianceType(v string) *ListAggregateConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList {
	s.ComplianceType = &v
	return s
}

func (s *ListAggregateConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList) SetConfigRuleInvokedTimestamp(v int64) *ListAggregateConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList {
	s.ConfigRuleInvokedTimestamp = &v
	return s
}

func (s *ListAggregateConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList) SetEvaluationResultIdentifier(v *ListAggregateConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifier) *ListAggregateConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList {
	s.EvaluationResultIdentifier = v
	return s
}

func (s *ListAggregateConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList) SetInvokingEventMessageType(v string) *ListAggregateConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList {
	s.InvokingEventMessageType = &v
	return s
}

func (s *ListAggregateConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList) SetRemediationEnabled(v bool) *ListAggregateConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList {
	s.RemediationEnabled = &v
	return s
}

func (s *ListAggregateConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList) SetResultRecordedTimestamp(v int64) *ListAggregateConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList {
	s.ResultRecordedTimestamp = &v
	return s
}

func (s *ListAggregateConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList) SetRiskLevel(v int32) *ListAggregateConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList {
	s.RiskLevel = &v
	return s
}

type ListAggregateConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifier struct {
	EvaluationResultQualifier *ListAggregateConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier `json:"EvaluationResultQualifier,omitempty" xml:"EvaluationResultQualifier,omitempty" type:"Struct"`
	OrderingTimestamp         *int64                                                                                                                                        `json:"OrderingTimestamp,omitempty" xml:"OrderingTimestamp,omitempty"`
}

func (s ListAggregateConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifier) String() string {
	return tea.Prettify(s)
}

func (s ListAggregateConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifier) GoString() string {
	return s.String()
}

func (s *ListAggregateConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifier) SetEvaluationResultQualifier(v *ListAggregateConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier) *ListAggregateConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifier {
	s.EvaluationResultQualifier = v
	return s
}

func (s *ListAggregateConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifier) SetOrderingTimestamp(v int64) *ListAggregateConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifier {
	s.OrderingTimestamp = &v
	return s
}

type ListAggregateConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier struct {
	CompliancePackId *string `json:"CompliancePackId,omitempty" xml:"CompliancePackId,omitempty"`
	ConfigRuleArn    *string `json:"ConfigRuleArn,omitempty" xml:"ConfigRuleArn,omitempty"`
	ConfigRuleId     *string `json:"ConfigRuleId,omitempty" xml:"ConfigRuleId,omitempty"`
	ConfigRuleName   *string `json:"ConfigRuleName,omitempty" xml:"ConfigRuleName,omitempty"`
	IgnoreDate       *string `json:"IgnoreDate,omitempty" xml:"IgnoreDate,omitempty"`
	RegionId         *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceId       *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ResourceName     *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	ResourceOwnerId  *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ResourceType     *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s ListAggregateConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier) String() string {
	return tea.Prettify(s)
}

func (s ListAggregateConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier) GoString() string {
	return s.String()
}

func (s *ListAggregateConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier) SetCompliancePackId(v string) *ListAggregateConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier {
	s.CompliancePackId = &v
	return s
}

func (s *ListAggregateConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier) SetConfigRuleArn(v string) *ListAggregateConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier {
	s.ConfigRuleArn = &v
	return s
}

func (s *ListAggregateConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier) SetConfigRuleId(v string) *ListAggregateConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier {
	s.ConfigRuleId = &v
	return s
}

func (s *ListAggregateConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier) SetConfigRuleName(v string) *ListAggregateConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier {
	s.ConfigRuleName = &v
	return s
}

func (s *ListAggregateConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier) SetIgnoreDate(v string) *ListAggregateConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier {
	s.IgnoreDate = &v
	return s
}

func (s *ListAggregateConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier) SetRegionId(v string) *ListAggregateConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier {
	s.RegionId = &v
	return s
}

func (s *ListAggregateConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier) SetResourceId(v string) *ListAggregateConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier {
	s.ResourceId = &v
	return s
}

func (s *ListAggregateConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier) SetResourceName(v string) *ListAggregateConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier {
	s.ResourceName = &v
	return s
}

func (s *ListAggregateConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier) SetResourceOwnerId(v int64) *ListAggregateConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListAggregateConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier) SetResourceType(v string) *ListAggregateConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier {
	s.ResourceType = &v
	return s
}

type ListAggregateConfigRuleEvaluationResultsResponse struct {
	Headers    map[string]*string                                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListAggregateConfigRuleEvaluationResultsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListAggregateConfigRuleEvaluationResultsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAggregateConfigRuleEvaluationResultsResponse) GoString() string {
	return s.String()
}

func (s *ListAggregateConfigRuleEvaluationResultsResponse) SetHeaders(v map[string]*string) *ListAggregateConfigRuleEvaluationResultsResponse {
	s.Headers = v
	return s
}

func (s *ListAggregateConfigRuleEvaluationResultsResponse) SetStatusCode(v int32) *ListAggregateConfigRuleEvaluationResultsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAggregateConfigRuleEvaluationResultsResponse) SetBody(v *ListAggregateConfigRuleEvaluationResultsResponseBody) *ListAggregateConfigRuleEvaluationResultsResponse {
	s.Body = v
	return s
}

type ListAggregateConfigRulesRequest struct {
	AggregatorId    *string `json:"AggregatorId,omitempty" xml:"AggregatorId,omitempty"`
	ComplianceType  *string `json:"ComplianceType,omitempty" xml:"ComplianceType,omitempty"`
	ConfigRuleName  *string `json:"ConfigRuleName,omitempty" xml:"ConfigRuleName,omitempty"`
	ConfigRuleState *string `json:"ConfigRuleState,omitempty" xml:"ConfigRuleState,omitempty"`
	PageNumber      *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize        *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RiskLevel       *int32  `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
}

func (s ListAggregateConfigRulesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAggregateConfigRulesRequest) GoString() string {
	return s.String()
}

func (s *ListAggregateConfigRulesRequest) SetAggregatorId(v string) *ListAggregateConfigRulesRequest {
	s.AggregatorId = &v
	return s
}

func (s *ListAggregateConfigRulesRequest) SetComplianceType(v string) *ListAggregateConfigRulesRequest {
	s.ComplianceType = &v
	return s
}

func (s *ListAggregateConfigRulesRequest) SetConfigRuleName(v string) *ListAggregateConfigRulesRequest {
	s.ConfigRuleName = &v
	return s
}

func (s *ListAggregateConfigRulesRequest) SetConfigRuleState(v string) *ListAggregateConfigRulesRequest {
	s.ConfigRuleState = &v
	return s
}

func (s *ListAggregateConfigRulesRequest) SetPageNumber(v int32) *ListAggregateConfigRulesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListAggregateConfigRulesRequest) SetPageSize(v int32) *ListAggregateConfigRulesRequest {
	s.PageSize = &v
	return s
}

func (s *ListAggregateConfigRulesRequest) SetRiskLevel(v int32) *ListAggregateConfigRulesRequest {
	s.RiskLevel = &v
	return s
}

type ListAggregateConfigRulesResponseBody struct {
	ConfigRules *ListAggregateConfigRulesResponseBodyConfigRules `json:"ConfigRules,omitempty" xml:"ConfigRules,omitempty" type:"Struct"`
	RequestId   *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListAggregateConfigRulesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAggregateConfigRulesResponseBody) GoString() string {
	return s.String()
}

func (s *ListAggregateConfigRulesResponseBody) SetConfigRules(v *ListAggregateConfigRulesResponseBodyConfigRules) *ListAggregateConfigRulesResponseBody {
	s.ConfigRules = v
	return s
}

func (s *ListAggregateConfigRulesResponseBody) SetRequestId(v string) *ListAggregateConfigRulesResponseBody {
	s.RequestId = &v
	return s
}

type ListAggregateConfigRulesResponseBodyConfigRules struct {
	ConfigRuleList []*ListAggregateConfigRulesResponseBodyConfigRulesConfigRuleList `json:"ConfigRuleList,omitempty" xml:"ConfigRuleList,omitempty" type:"Repeated"`
	PageNumber     *int32                                                           `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize       *int32                                                           `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount     *int64                                                           `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListAggregateConfigRulesResponseBodyConfigRules) String() string {
	return tea.Prettify(s)
}

func (s ListAggregateConfigRulesResponseBodyConfigRules) GoString() string {
	return s.String()
}

func (s *ListAggregateConfigRulesResponseBodyConfigRules) SetConfigRuleList(v []*ListAggregateConfigRulesResponseBodyConfigRulesConfigRuleList) *ListAggregateConfigRulesResponseBodyConfigRules {
	s.ConfigRuleList = v
	return s
}

func (s *ListAggregateConfigRulesResponseBodyConfigRules) SetPageNumber(v int32) *ListAggregateConfigRulesResponseBodyConfigRules {
	s.PageNumber = &v
	return s
}

func (s *ListAggregateConfigRulesResponseBodyConfigRules) SetPageSize(v int32) *ListAggregateConfigRulesResponseBodyConfigRules {
	s.PageSize = &v
	return s
}

func (s *ListAggregateConfigRulesResponseBodyConfigRules) SetTotalCount(v int64) *ListAggregateConfigRulesResponseBodyConfigRules {
	s.TotalCount = &v
	return s
}

type ListAggregateConfigRulesResponseBodyConfigRulesConfigRuleList struct {
	AccountId        *int64                                                                   `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	AutomationType   *string                                                                  `json:"AutomationType,omitempty" xml:"AutomationType,omitempty"`
	Compliance       *ListAggregateConfigRulesResponseBodyConfigRulesConfigRuleListCompliance `json:"Compliance,omitempty" xml:"Compliance,omitempty" type:"Struct"`
	ConfigRuleArn    *string                                                                  `json:"ConfigRuleArn,omitempty" xml:"ConfigRuleArn,omitempty"`
	ConfigRuleId     *string                                                                  `json:"ConfigRuleId,omitempty" xml:"ConfigRuleId,omitempty"`
	ConfigRuleName   *string                                                                  `json:"ConfigRuleName,omitempty" xml:"ConfigRuleName,omitempty"`
	ConfigRuleState  *string                                                                  `json:"ConfigRuleState,omitempty" xml:"ConfigRuleState,omitempty"`
	CreateBy         *ListAggregateConfigRulesResponseBodyConfigRulesConfigRuleListCreateBy   `json:"CreateBy,omitempty" xml:"CreateBy,omitempty" type:"Struct"`
	Description      *string                                                                  `json:"Description,omitempty" xml:"Description,omitempty"`
	RiskLevel        *int32                                                                   `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
	SourceIdentifier *string                                                                  `json:"SourceIdentifier,omitempty" xml:"SourceIdentifier,omitempty"`
	SourceOwner      *string                                                                  `json:"SourceOwner,omitempty" xml:"SourceOwner,omitempty"`
	Tags             []*ListAggregateConfigRulesResponseBodyConfigRulesConfigRuleListTags     `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s ListAggregateConfigRulesResponseBodyConfigRulesConfigRuleList) String() string {
	return tea.Prettify(s)
}

func (s ListAggregateConfigRulesResponseBodyConfigRulesConfigRuleList) GoString() string {
	return s.String()
}

func (s *ListAggregateConfigRulesResponseBodyConfigRulesConfigRuleList) SetAccountId(v int64) *ListAggregateConfigRulesResponseBodyConfigRulesConfigRuleList {
	s.AccountId = &v
	return s
}

func (s *ListAggregateConfigRulesResponseBodyConfigRulesConfigRuleList) SetAutomationType(v string) *ListAggregateConfigRulesResponseBodyConfigRulesConfigRuleList {
	s.AutomationType = &v
	return s
}

func (s *ListAggregateConfigRulesResponseBodyConfigRulesConfigRuleList) SetCompliance(v *ListAggregateConfigRulesResponseBodyConfigRulesConfigRuleListCompliance) *ListAggregateConfigRulesResponseBodyConfigRulesConfigRuleList {
	s.Compliance = v
	return s
}

func (s *ListAggregateConfigRulesResponseBodyConfigRulesConfigRuleList) SetConfigRuleArn(v string) *ListAggregateConfigRulesResponseBodyConfigRulesConfigRuleList {
	s.ConfigRuleArn = &v
	return s
}

func (s *ListAggregateConfigRulesResponseBodyConfigRulesConfigRuleList) SetConfigRuleId(v string) *ListAggregateConfigRulesResponseBodyConfigRulesConfigRuleList {
	s.ConfigRuleId = &v
	return s
}

func (s *ListAggregateConfigRulesResponseBodyConfigRulesConfigRuleList) SetConfigRuleName(v string) *ListAggregateConfigRulesResponseBodyConfigRulesConfigRuleList {
	s.ConfigRuleName = &v
	return s
}

func (s *ListAggregateConfigRulesResponseBodyConfigRulesConfigRuleList) SetConfigRuleState(v string) *ListAggregateConfigRulesResponseBodyConfigRulesConfigRuleList {
	s.ConfigRuleState = &v
	return s
}

func (s *ListAggregateConfigRulesResponseBodyConfigRulesConfigRuleList) SetCreateBy(v *ListAggregateConfigRulesResponseBodyConfigRulesConfigRuleListCreateBy) *ListAggregateConfigRulesResponseBodyConfigRulesConfigRuleList {
	s.CreateBy = v
	return s
}

func (s *ListAggregateConfigRulesResponseBodyConfigRulesConfigRuleList) SetDescription(v string) *ListAggregateConfigRulesResponseBodyConfigRulesConfigRuleList {
	s.Description = &v
	return s
}

func (s *ListAggregateConfigRulesResponseBodyConfigRulesConfigRuleList) SetRiskLevel(v int32) *ListAggregateConfigRulesResponseBodyConfigRulesConfigRuleList {
	s.RiskLevel = &v
	return s
}

func (s *ListAggregateConfigRulesResponseBodyConfigRulesConfigRuleList) SetSourceIdentifier(v string) *ListAggregateConfigRulesResponseBodyConfigRulesConfigRuleList {
	s.SourceIdentifier = &v
	return s
}

func (s *ListAggregateConfigRulesResponseBodyConfigRulesConfigRuleList) SetSourceOwner(v string) *ListAggregateConfigRulesResponseBodyConfigRulesConfigRuleList {
	s.SourceOwner = &v
	return s
}

func (s *ListAggregateConfigRulesResponseBodyConfigRulesConfigRuleList) SetTags(v []*ListAggregateConfigRulesResponseBodyConfigRulesConfigRuleListTags) *ListAggregateConfigRulesResponseBodyConfigRulesConfigRuleList {
	s.Tags = v
	return s
}

type ListAggregateConfigRulesResponseBodyConfigRulesConfigRuleListCompliance struct {
	ComplianceType *string `json:"ComplianceType,omitempty" xml:"ComplianceType,omitempty"`
	Count          *int32  `json:"Count,omitempty" xml:"Count,omitempty"`
}

func (s ListAggregateConfigRulesResponseBodyConfigRulesConfigRuleListCompliance) String() string {
	return tea.Prettify(s)
}

func (s ListAggregateConfigRulesResponseBodyConfigRulesConfigRuleListCompliance) GoString() string {
	return s.String()
}

func (s *ListAggregateConfigRulesResponseBodyConfigRulesConfigRuleListCompliance) SetComplianceType(v string) *ListAggregateConfigRulesResponseBodyConfigRulesConfigRuleListCompliance {
	s.ComplianceType = &v
	return s
}

func (s *ListAggregateConfigRulesResponseBodyConfigRulesConfigRuleListCompliance) SetCount(v int32) *ListAggregateConfigRulesResponseBodyConfigRulesConfigRuleListCompliance {
	s.Count = &v
	return s
}

type ListAggregateConfigRulesResponseBodyConfigRulesConfigRuleListCreateBy struct {
	AggregatorId       *string `json:"AggregatorId,omitempty" xml:"AggregatorId,omitempty"`
	AggregatorName     *string `json:"AggregatorName,omitempty" xml:"AggregatorName,omitempty"`
	CompliancePackId   *string `json:"CompliancePackId,omitempty" xml:"CompliancePackId,omitempty"`
	CompliancePackName *string `json:"CompliancePackName,omitempty" xml:"CompliancePackName,omitempty"`
	CreatorId          *string `json:"CreatorId,omitempty" xml:"CreatorId,omitempty"`
	CreatorName        *string `json:"CreatorName,omitempty" xml:"CreatorName,omitempty"`
	CreatorType        *string `json:"CreatorType,omitempty" xml:"CreatorType,omitempty"`
}

func (s ListAggregateConfigRulesResponseBodyConfigRulesConfigRuleListCreateBy) String() string {
	return tea.Prettify(s)
}

func (s ListAggregateConfigRulesResponseBodyConfigRulesConfigRuleListCreateBy) GoString() string {
	return s.String()
}

func (s *ListAggregateConfigRulesResponseBodyConfigRulesConfigRuleListCreateBy) SetAggregatorId(v string) *ListAggregateConfigRulesResponseBodyConfigRulesConfigRuleListCreateBy {
	s.AggregatorId = &v
	return s
}

func (s *ListAggregateConfigRulesResponseBodyConfigRulesConfigRuleListCreateBy) SetAggregatorName(v string) *ListAggregateConfigRulesResponseBodyConfigRulesConfigRuleListCreateBy {
	s.AggregatorName = &v
	return s
}

func (s *ListAggregateConfigRulesResponseBodyConfigRulesConfigRuleListCreateBy) SetCompliancePackId(v string) *ListAggregateConfigRulesResponseBodyConfigRulesConfigRuleListCreateBy {
	s.CompliancePackId = &v
	return s
}

func (s *ListAggregateConfigRulesResponseBodyConfigRulesConfigRuleListCreateBy) SetCompliancePackName(v string) *ListAggregateConfigRulesResponseBodyConfigRulesConfigRuleListCreateBy {
	s.CompliancePackName = &v
	return s
}

func (s *ListAggregateConfigRulesResponseBodyConfigRulesConfigRuleListCreateBy) SetCreatorId(v string) *ListAggregateConfigRulesResponseBodyConfigRulesConfigRuleListCreateBy {
	s.CreatorId = &v
	return s
}

func (s *ListAggregateConfigRulesResponseBodyConfigRulesConfigRuleListCreateBy) SetCreatorName(v string) *ListAggregateConfigRulesResponseBodyConfigRulesConfigRuleListCreateBy {
	s.CreatorName = &v
	return s
}

func (s *ListAggregateConfigRulesResponseBodyConfigRulesConfigRuleListCreateBy) SetCreatorType(v string) *ListAggregateConfigRulesResponseBodyConfigRulesConfigRuleListCreateBy {
	s.CreatorType = &v
	return s
}

type ListAggregateConfigRulesResponseBodyConfigRulesConfigRuleListTags struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListAggregateConfigRulesResponseBodyConfigRulesConfigRuleListTags) String() string {
	return tea.Prettify(s)
}

func (s ListAggregateConfigRulesResponseBodyConfigRulesConfigRuleListTags) GoString() string {
	return s.String()
}

func (s *ListAggregateConfigRulesResponseBodyConfigRulesConfigRuleListTags) SetKey(v string) *ListAggregateConfigRulesResponseBodyConfigRulesConfigRuleListTags {
	s.Key = &v
	return s
}

func (s *ListAggregateConfigRulesResponseBodyConfigRulesConfigRuleListTags) SetValue(v string) *ListAggregateConfigRulesResponseBodyConfigRulesConfigRuleListTags {
	s.Value = &v
	return s
}

type ListAggregateConfigRulesResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListAggregateConfigRulesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListAggregateConfigRulesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAggregateConfigRulesResponse) GoString() string {
	return s.String()
}

func (s *ListAggregateConfigRulesResponse) SetHeaders(v map[string]*string) *ListAggregateConfigRulesResponse {
	s.Headers = v
	return s
}

func (s *ListAggregateConfigRulesResponse) SetStatusCode(v int32) *ListAggregateConfigRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAggregateConfigRulesResponse) SetBody(v *ListAggregateConfigRulesResponseBody) *ListAggregateConfigRulesResponse {
	s.Body = v
	return s
}

type ListAggregateDiscoveredResourcesRequest struct {
	AggregatorId    *string `json:"AggregatorId,omitempty" xml:"AggregatorId,omitempty"`
	MaxResults      *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken       *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	Regions         *string `json:"Regions,omitempty" xml:"Regions,omitempty"`
	ResourceDeleted *int32  `json:"ResourceDeleted,omitempty" xml:"ResourceDeleted,omitempty"`
	ResourceId      *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ResourceTypes   *string `json:"ResourceTypes,omitempty" xml:"ResourceTypes,omitempty"`
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

func (s *ListAggregateDiscoveredResourcesRequest) SetMaxResults(v int32) *ListAggregateDiscoveredResourcesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListAggregateDiscoveredResourcesRequest) SetNextToken(v string) *ListAggregateDiscoveredResourcesRequest {
	s.NextToken = &v
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
	DiscoveredResourceProfiles *ListAggregateDiscoveredResourcesResponseBodyDiscoveredResourceProfiles `json:"DiscoveredResourceProfiles,omitempty" xml:"DiscoveredResourceProfiles,omitempty" type:"Struct"`
	RequestId                  *string                                                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	DiscoveredResourceProfileList []*ListAggregateDiscoveredResourcesResponseBodyDiscoveredResourceProfilesDiscoveredResourceProfileList `json:"DiscoveredResourceProfileList,omitempty" xml:"DiscoveredResourceProfileList,omitempty" type:"Repeated"`
	MaxResults                    *int32                                                                                                 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken                     *string                                                                                                `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	TotalCount                    *int32                                                                                                 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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

func (s *ListAggregateDiscoveredResourcesResponseBodyDiscoveredResourceProfiles) SetMaxResults(v int32) *ListAggregateDiscoveredResourcesResponseBodyDiscoveredResourceProfiles {
	s.MaxResults = &v
	return s
}

func (s *ListAggregateDiscoveredResourcesResponseBodyDiscoveredResourceProfiles) SetNextToken(v string) *ListAggregateDiscoveredResourcesResponseBodyDiscoveredResourceProfiles {
	s.NextToken = &v
	return s
}

func (s *ListAggregateDiscoveredResourcesResponseBodyDiscoveredResourceProfiles) SetTotalCount(v int32) *ListAggregateDiscoveredResourcesResponseBodyDiscoveredResourceProfiles {
	s.TotalCount = &v
	return s
}

type ListAggregateDiscoveredResourcesResponseBodyDiscoveredResourceProfilesDiscoveredResourceProfileList struct {
	AccountId            *int64  `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	Region               *string `json:"Region,omitempty" xml:"Region,omitempty"`
	ResourceCreationTime *int64  `json:"ResourceCreationTime,omitempty" xml:"ResourceCreationTime,omitempty"`
	ResourceDeleted      *int32  `json:"ResourceDeleted,omitempty" xml:"ResourceDeleted,omitempty"`
	ResourceId           *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ResourceName         *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ResourceStatus       *string `json:"ResourceStatus,omitempty" xml:"ResourceStatus,omitempty"`
	ResourceType         *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	Tags                 *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
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

type ListAggregateRemediationsRequest struct {
	AggregatorId  *string `json:"AggregatorId,omitempty" xml:"AggregatorId,omitempty"`
	ConfigRuleIds *string `json:"ConfigRuleIds,omitempty" xml:"ConfigRuleIds,omitempty"`
}

func (s ListAggregateRemediationsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAggregateRemediationsRequest) GoString() string {
	return s.String()
}

func (s *ListAggregateRemediationsRequest) SetAggregatorId(v string) *ListAggregateRemediationsRequest {
	s.AggregatorId = &v
	return s
}

func (s *ListAggregateRemediationsRequest) SetConfigRuleIds(v string) *ListAggregateRemediationsRequest {
	s.ConfigRuleIds = &v
	return s
}

type ListAggregateRemediationsResponseBody struct {
	Remediations []*ListAggregateRemediationsResponseBodyRemediations `json:"Remediations,omitempty" xml:"Remediations,omitempty" type:"Repeated"`
	RequestId    *string                                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListAggregateRemediationsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAggregateRemediationsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAggregateRemediationsResponseBody) SetRemediations(v []*ListAggregateRemediationsResponseBodyRemediations) *ListAggregateRemediationsResponseBody {
	s.Remediations = v
	return s
}

func (s *ListAggregateRemediationsResponseBody) SetRequestId(v string) *ListAggregateRemediationsResponseBody {
	s.RequestId = &v
	return s
}

type ListAggregateRemediationsResponseBodyRemediations struct {
	AccountId                    *int64  `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	AggregatorId                 *string `json:"AggregatorId,omitempty" xml:"AggregatorId,omitempty"`
	ConfigRuleId                 *string `json:"ConfigRuleId,omitempty" xml:"ConfigRuleId,omitempty"`
	InvokeType                   *string `json:"InvokeType,omitempty" xml:"InvokeType,omitempty"`
	LastSuccessfulInvocationId   *string `json:"LastSuccessfulInvocationId,omitempty" xml:"LastSuccessfulInvocationId,omitempty"`
	LastSuccessfulInvocationTime *int64  `json:"LastSuccessfulInvocationTime,omitempty" xml:"LastSuccessfulInvocationTime,omitempty"`
	LastSuccessfulInvocationType *string `json:"LastSuccessfulInvocationType,omitempty" xml:"LastSuccessfulInvocationType,omitempty"`
	RemediaitonOriginParams      *string `json:"RemediaitonOriginParams,omitempty" xml:"RemediaitonOriginParams,omitempty"`
	RemediationId                *string `json:"RemediationId,omitempty" xml:"RemediationId,omitempty"`
	RemediationSourceType        *string `json:"RemediationSourceType,omitempty" xml:"RemediationSourceType,omitempty"`
	RemediationTemplateId        *string `json:"RemediationTemplateId,omitempty" xml:"RemediationTemplateId,omitempty"`
	RemediationType              *string `json:"RemediationType,omitempty" xml:"RemediationType,omitempty"`
}

func (s ListAggregateRemediationsResponseBodyRemediations) String() string {
	return tea.Prettify(s)
}

func (s ListAggregateRemediationsResponseBodyRemediations) GoString() string {
	return s.String()
}

func (s *ListAggregateRemediationsResponseBodyRemediations) SetAccountId(v int64) *ListAggregateRemediationsResponseBodyRemediations {
	s.AccountId = &v
	return s
}

func (s *ListAggregateRemediationsResponseBodyRemediations) SetAggregatorId(v string) *ListAggregateRemediationsResponseBodyRemediations {
	s.AggregatorId = &v
	return s
}

func (s *ListAggregateRemediationsResponseBodyRemediations) SetConfigRuleId(v string) *ListAggregateRemediationsResponseBodyRemediations {
	s.ConfigRuleId = &v
	return s
}

func (s *ListAggregateRemediationsResponseBodyRemediations) SetInvokeType(v string) *ListAggregateRemediationsResponseBodyRemediations {
	s.InvokeType = &v
	return s
}

func (s *ListAggregateRemediationsResponseBodyRemediations) SetLastSuccessfulInvocationId(v string) *ListAggregateRemediationsResponseBodyRemediations {
	s.LastSuccessfulInvocationId = &v
	return s
}

func (s *ListAggregateRemediationsResponseBodyRemediations) SetLastSuccessfulInvocationTime(v int64) *ListAggregateRemediationsResponseBodyRemediations {
	s.LastSuccessfulInvocationTime = &v
	return s
}

func (s *ListAggregateRemediationsResponseBodyRemediations) SetLastSuccessfulInvocationType(v string) *ListAggregateRemediationsResponseBodyRemediations {
	s.LastSuccessfulInvocationType = &v
	return s
}

func (s *ListAggregateRemediationsResponseBodyRemediations) SetRemediaitonOriginParams(v string) *ListAggregateRemediationsResponseBodyRemediations {
	s.RemediaitonOriginParams = &v
	return s
}

func (s *ListAggregateRemediationsResponseBodyRemediations) SetRemediationId(v string) *ListAggregateRemediationsResponseBodyRemediations {
	s.RemediationId = &v
	return s
}

func (s *ListAggregateRemediationsResponseBodyRemediations) SetRemediationSourceType(v string) *ListAggregateRemediationsResponseBodyRemediations {
	s.RemediationSourceType = &v
	return s
}

func (s *ListAggregateRemediationsResponseBodyRemediations) SetRemediationTemplateId(v string) *ListAggregateRemediationsResponseBodyRemediations {
	s.RemediationTemplateId = &v
	return s
}

func (s *ListAggregateRemediationsResponseBodyRemediations) SetRemediationType(v string) *ListAggregateRemediationsResponseBodyRemediations {
	s.RemediationType = &v
	return s
}

type ListAggregateRemediationsResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListAggregateRemediationsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListAggregateRemediationsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAggregateRemediationsResponse) GoString() string {
	return s.String()
}

func (s *ListAggregateRemediationsResponse) SetHeaders(v map[string]*string) *ListAggregateRemediationsResponse {
	s.Headers = v
	return s
}

func (s *ListAggregateRemediationsResponse) SetStatusCode(v int32) *ListAggregateRemediationsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAggregateRemediationsResponse) SetBody(v *ListAggregateRemediationsResponseBody) *ListAggregateRemediationsResponse {
	s.Body = v
	return s
}

type ListAggregateResourceEvaluationResultsRequest struct {
	AggregatorId   *string `json:"AggregatorId,omitempty" xml:"AggregatorId,omitempty"`
	ComplianceType *string `json:"ComplianceType,omitempty" xml:"ComplianceType,omitempty"`
	MaxResults     *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken      *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	Region         *string `json:"Region,omitempty" xml:"Region,omitempty"`
	ResourceId     *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ResourceType   *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s ListAggregateResourceEvaluationResultsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAggregateResourceEvaluationResultsRequest) GoString() string {
	return s.String()
}

func (s *ListAggregateResourceEvaluationResultsRequest) SetAggregatorId(v string) *ListAggregateResourceEvaluationResultsRequest {
	s.AggregatorId = &v
	return s
}

func (s *ListAggregateResourceEvaluationResultsRequest) SetComplianceType(v string) *ListAggregateResourceEvaluationResultsRequest {
	s.ComplianceType = &v
	return s
}

func (s *ListAggregateResourceEvaluationResultsRequest) SetMaxResults(v int32) *ListAggregateResourceEvaluationResultsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListAggregateResourceEvaluationResultsRequest) SetNextToken(v string) *ListAggregateResourceEvaluationResultsRequest {
	s.NextToken = &v
	return s
}

func (s *ListAggregateResourceEvaluationResultsRequest) SetRegion(v string) *ListAggregateResourceEvaluationResultsRequest {
	s.Region = &v
	return s
}

func (s *ListAggregateResourceEvaluationResultsRequest) SetResourceId(v string) *ListAggregateResourceEvaluationResultsRequest {
	s.ResourceId = &v
	return s
}

func (s *ListAggregateResourceEvaluationResultsRequest) SetResourceType(v string) *ListAggregateResourceEvaluationResultsRequest {
	s.ResourceType = &v
	return s
}

type ListAggregateResourceEvaluationResultsResponseBody struct {
	EvaluationResults *ListAggregateResourceEvaluationResultsResponseBodyEvaluationResults `json:"EvaluationResults,omitempty" xml:"EvaluationResults,omitempty" type:"Struct"`
	RequestId         *string                                                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListAggregateResourceEvaluationResultsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAggregateResourceEvaluationResultsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAggregateResourceEvaluationResultsResponseBody) SetEvaluationResults(v *ListAggregateResourceEvaluationResultsResponseBodyEvaluationResults) *ListAggregateResourceEvaluationResultsResponseBody {
	s.EvaluationResults = v
	return s
}

func (s *ListAggregateResourceEvaluationResultsResponseBody) SetRequestId(v string) *ListAggregateResourceEvaluationResultsResponseBody {
	s.RequestId = &v
	return s
}

type ListAggregateResourceEvaluationResultsResponseBodyEvaluationResults struct {
	EvaluationResultList []*ListAggregateResourceEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList `json:"EvaluationResultList,omitempty" xml:"EvaluationResultList,omitempty" type:"Repeated"`
	MaxResults           *int32                                                                                     `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken            *string                                                                                    `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListAggregateResourceEvaluationResultsResponseBodyEvaluationResults) String() string {
	return tea.Prettify(s)
}

func (s ListAggregateResourceEvaluationResultsResponseBodyEvaluationResults) GoString() string {
	return s.String()
}

func (s *ListAggregateResourceEvaluationResultsResponseBodyEvaluationResults) SetEvaluationResultList(v []*ListAggregateResourceEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList) *ListAggregateResourceEvaluationResultsResponseBodyEvaluationResults {
	s.EvaluationResultList = v
	return s
}

func (s *ListAggregateResourceEvaluationResultsResponseBodyEvaluationResults) SetMaxResults(v int32) *ListAggregateResourceEvaluationResultsResponseBodyEvaluationResults {
	s.MaxResults = &v
	return s
}

func (s *ListAggregateResourceEvaluationResultsResponseBodyEvaluationResults) SetNextToken(v string) *ListAggregateResourceEvaluationResultsResponseBodyEvaluationResults {
	s.NextToken = &v
	return s
}

type ListAggregateResourceEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList struct {
	Annotation                 *string                                                                                                            `json:"Annotation,omitempty" xml:"Annotation,omitempty"`
	ComplianceType             *string                                                                                                            `json:"ComplianceType,omitempty" xml:"ComplianceType,omitempty"`
	ConfigRuleInvokedTimestamp *int64                                                                                                             `json:"ConfigRuleInvokedTimestamp,omitempty" xml:"ConfigRuleInvokedTimestamp,omitempty"`
	EvaluationResultIdentifier *ListAggregateResourceEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifier `json:"EvaluationResultIdentifier,omitempty" xml:"EvaluationResultIdentifier,omitempty" type:"Struct"`
	InvokingEventMessageType   *string                                                                                                            `json:"InvokingEventMessageType,omitempty" xml:"InvokingEventMessageType,omitempty"`
	RemediationEnabled         *bool                                                                                                              `json:"RemediationEnabled,omitempty" xml:"RemediationEnabled,omitempty"`
	ResultRecordedTimestamp    *int64                                                                                                             `json:"ResultRecordedTimestamp,omitempty" xml:"ResultRecordedTimestamp,omitempty"`
	RiskLevel                  *int32                                                                                                             `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
}

func (s ListAggregateResourceEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList) String() string {
	return tea.Prettify(s)
}

func (s ListAggregateResourceEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList) GoString() string {
	return s.String()
}

func (s *ListAggregateResourceEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList) SetAnnotation(v string) *ListAggregateResourceEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList {
	s.Annotation = &v
	return s
}

func (s *ListAggregateResourceEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList) SetComplianceType(v string) *ListAggregateResourceEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList {
	s.ComplianceType = &v
	return s
}

func (s *ListAggregateResourceEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList) SetConfigRuleInvokedTimestamp(v int64) *ListAggregateResourceEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList {
	s.ConfigRuleInvokedTimestamp = &v
	return s
}

func (s *ListAggregateResourceEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList) SetEvaluationResultIdentifier(v *ListAggregateResourceEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifier) *ListAggregateResourceEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList {
	s.EvaluationResultIdentifier = v
	return s
}

func (s *ListAggregateResourceEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList) SetInvokingEventMessageType(v string) *ListAggregateResourceEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList {
	s.InvokingEventMessageType = &v
	return s
}

func (s *ListAggregateResourceEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList) SetRemediationEnabled(v bool) *ListAggregateResourceEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList {
	s.RemediationEnabled = &v
	return s
}

func (s *ListAggregateResourceEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList) SetResultRecordedTimestamp(v int64) *ListAggregateResourceEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList {
	s.ResultRecordedTimestamp = &v
	return s
}

func (s *ListAggregateResourceEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList) SetRiskLevel(v int32) *ListAggregateResourceEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList {
	s.RiskLevel = &v
	return s
}

type ListAggregateResourceEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifier struct {
	EvaluationResultQualifier *ListAggregateResourceEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier `json:"EvaluationResultQualifier,omitempty" xml:"EvaluationResultQualifier,omitempty" type:"Struct"`
	OrderingTimestamp         *int64                                                                                                                                      `json:"OrderingTimestamp,omitempty" xml:"OrderingTimestamp,omitempty"`
}

func (s ListAggregateResourceEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifier) String() string {
	return tea.Prettify(s)
}

func (s ListAggregateResourceEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifier) GoString() string {
	return s.String()
}

func (s *ListAggregateResourceEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifier) SetEvaluationResultQualifier(v *ListAggregateResourceEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier) *ListAggregateResourceEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifier {
	s.EvaluationResultQualifier = v
	return s
}

func (s *ListAggregateResourceEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifier) SetOrderingTimestamp(v int64) *ListAggregateResourceEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifier {
	s.OrderingTimestamp = &v
	return s
}

type ListAggregateResourceEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier struct {
	ConfigRuleArn  *string `json:"ConfigRuleArn,omitempty" xml:"ConfigRuleArn,omitempty"`
	ConfigRuleId   *string `json:"ConfigRuleId,omitempty" xml:"ConfigRuleId,omitempty"`
	ConfigRuleName *string `json:"ConfigRuleName,omitempty" xml:"ConfigRuleName,omitempty"`
	IgnoreDate     *string `json:"IgnoreDate,omitempty" xml:"IgnoreDate,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceId     *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ResourceName   *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	ResourceType   *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s ListAggregateResourceEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier) String() string {
	return tea.Prettify(s)
}

func (s ListAggregateResourceEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier) GoString() string {
	return s.String()
}

func (s *ListAggregateResourceEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier) SetConfigRuleArn(v string) *ListAggregateResourceEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier {
	s.ConfigRuleArn = &v
	return s
}

func (s *ListAggregateResourceEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier) SetConfigRuleId(v string) *ListAggregateResourceEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier {
	s.ConfigRuleId = &v
	return s
}

func (s *ListAggregateResourceEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier) SetConfigRuleName(v string) *ListAggregateResourceEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier {
	s.ConfigRuleName = &v
	return s
}

func (s *ListAggregateResourceEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier) SetIgnoreDate(v string) *ListAggregateResourceEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier {
	s.IgnoreDate = &v
	return s
}

func (s *ListAggregateResourceEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier) SetRegionId(v string) *ListAggregateResourceEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier {
	s.RegionId = &v
	return s
}

func (s *ListAggregateResourceEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier) SetResourceId(v string) *ListAggregateResourceEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier {
	s.ResourceId = &v
	return s
}

func (s *ListAggregateResourceEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier) SetResourceName(v string) *ListAggregateResourceEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier {
	s.ResourceName = &v
	return s
}

func (s *ListAggregateResourceEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier) SetResourceType(v string) *ListAggregateResourceEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier {
	s.ResourceType = &v
	return s
}

type ListAggregateResourceEvaluationResultsResponse struct {
	Headers    map[string]*string                                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListAggregateResourceEvaluationResultsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListAggregateResourceEvaluationResultsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAggregateResourceEvaluationResultsResponse) GoString() string {
	return s.String()
}

func (s *ListAggregateResourceEvaluationResultsResponse) SetHeaders(v map[string]*string) *ListAggregateResourceEvaluationResultsResponse {
	s.Headers = v
	return s
}

func (s *ListAggregateResourceEvaluationResultsResponse) SetStatusCode(v int32) *ListAggregateResourceEvaluationResultsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAggregateResourceEvaluationResultsResponse) SetBody(v *ListAggregateResourceEvaluationResultsResponseBody) *ListAggregateResourceEvaluationResultsResponse {
	s.Body = v
	return s
}

type ListAggregatorsRequest struct {
	MaxResults *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken  *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListAggregatorsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAggregatorsRequest) GoString() string {
	return s.String()
}

func (s *ListAggregatorsRequest) SetMaxResults(v int32) *ListAggregatorsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListAggregatorsRequest) SetNextToken(v string) *ListAggregatorsRequest {
	s.NextToken = &v
	return s
}

type ListAggregatorsResponseBody struct {
	AggregatorsResult *ListAggregatorsResponseBodyAggregatorsResult `json:"AggregatorsResult,omitempty" xml:"AggregatorsResult,omitempty" type:"Struct"`
	RequestId         *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListAggregatorsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAggregatorsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAggregatorsResponseBody) SetAggregatorsResult(v *ListAggregatorsResponseBodyAggregatorsResult) *ListAggregatorsResponseBody {
	s.AggregatorsResult = v
	return s
}

func (s *ListAggregatorsResponseBody) SetRequestId(v string) *ListAggregatorsResponseBody {
	s.RequestId = &v
	return s
}

type ListAggregatorsResponseBodyAggregatorsResult struct {
	Aggregators []*ListAggregatorsResponseBodyAggregatorsResultAggregators `json:"Aggregators,omitempty" xml:"Aggregators,omitempty" type:"Repeated"`
	NextToken   *string                                                    `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListAggregatorsResponseBodyAggregatorsResult) String() string {
	return tea.Prettify(s)
}

func (s ListAggregatorsResponseBodyAggregatorsResult) GoString() string {
	return s.String()
}

func (s *ListAggregatorsResponseBodyAggregatorsResult) SetAggregators(v []*ListAggregatorsResponseBodyAggregatorsResultAggregators) *ListAggregatorsResponseBodyAggregatorsResult {
	s.Aggregators = v
	return s
}

func (s *ListAggregatorsResponseBodyAggregatorsResult) SetNextToken(v string) *ListAggregatorsResponseBodyAggregatorsResult {
	s.NextToken = &v
	return s
}

type ListAggregatorsResponseBodyAggregatorsResultAggregators struct {
	AccountId                 *int64  `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	AggregatorAccountCount    *int64  `json:"AggregatorAccountCount,omitempty" xml:"AggregatorAccountCount,omitempty"`
	AggregatorCreateTimestamp *int64  `json:"AggregatorCreateTimestamp,omitempty" xml:"AggregatorCreateTimestamp,omitempty"`
	AggregatorId              *string `json:"AggregatorId,omitempty" xml:"AggregatorId,omitempty"`
	AggregatorName            *string `json:"AggregatorName,omitempty" xml:"AggregatorName,omitempty"`
	AggregatorStatus          *int32  `json:"AggregatorStatus,omitempty" xml:"AggregatorStatus,omitempty"`
	AggregatorType            *string `json:"AggregatorType,omitempty" xml:"AggregatorType,omitempty"`
	Description               *string `json:"Description,omitempty" xml:"Description,omitempty"`
}

func (s ListAggregatorsResponseBodyAggregatorsResultAggregators) String() string {
	return tea.Prettify(s)
}

func (s ListAggregatorsResponseBodyAggregatorsResultAggregators) GoString() string {
	return s.String()
}

func (s *ListAggregatorsResponseBodyAggregatorsResultAggregators) SetAccountId(v int64) *ListAggregatorsResponseBodyAggregatorsResultAggregators {
	s.AccountId = &v
	return s
}

func (s *ListAggregatorsResponseBodyAggregatorsResultAggregators) SetAggregatorAccountCount(v int64) *ListAggregatorsResponseBodyAggregatorsResultAggregators {
	s.AggregatorAccountCount = &v
	return s
}

func (s *ListAggregatorsResponseBodyAggregatorsResultAggregators) SetAggregatorCreateTimestamp(v int64) *ListAggregatorsResponseBodyAggregatorsResultAggregators {
	s.AggregatorCreateTimestamp = &v
	return s
}

func (s *ListAggregatorsResponseBodyAggregatorsResultAggregators) SetAggregatorId(v string) *ListAggregatorsResponseBodyAggregatorsResultAggregators {
	s.AggregatorId = &v
	return s
}

func (s *ListAggregatorsResponseBodyAggregatorsResultAggregators) SetAggregatorName(v string) *ListAggregatorsResponseBodyAggregatorsResultAggregators {
	s.AggregatorName = &v
	return s
}

func (s *ListAggregatorsResponseBodyAggregatorsResultAggregators) SetAggregatorStatus(v int32) *ListAggregatorsResponseBodyAggregatorsResultAggregators {
	s.AggregatorStatus = &v
	return s
}

func (s *ListAggregatorsResponseBodyAggregatorsResultAggregators) SetAggregatorType(v string) *ListAggregatorsResponseBodyAggregatorsResultAggregators {
	s.AggregatorType = &v
	return s
}

func (s *ListAggregatorsResponseBodyAggregatorsResultAggregators) SetDescription(v string) *ListAggregatorsResponseBodyAggregatorsResultAggregators {
	s.Description = &v
	return s
}

type ListAggregatorsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListAggregatorsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListAggregatorsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAggregatorsResponse) GoString() string {
	return s.String()
}

func (s *ListAggregatorsResponse) SetHeaders(v map[string]*string) *ListAggregatorsResponse {
	s.Headers = v
	return s
}

func (s *ListAggregatorsResponse) SetStatusCode(v int32) *ListAggregatorsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAggregatorsResponse) SetBody(v *ListAggregatorsResponseBody) *ListAggregatorsResponse {
	s.Body = v
	return s
}

type ListCompliancePackTemplatesRequest struct {
	CompliancePackTemplateId *string `json:"CompliancePackTemplateId,omitempty" xml:"CompliancePackTemplateId,omitempty"`
	PageNumber               *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize                 *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListCompliancePackTemplatesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListCompliancePackTemplatesRequest) GoString() string {
	return s.String()
}

func (s *ListCompliancePackTemplatesRequest) SetCompliancePackTemplateId(v string) *ListCompliancePackTemplatesRequest {
	s.CompliancePackTemplateId = &v
	return s
}

func (s *ListCompliancePackTemplatesRequest) SetPageNumber(v int32) *ListCompliancePackTemplatesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListCompliancePackTemplatesRequest) SetPageSize(v int32) *ListCompliancePackTemplatesRequest {
	s.PageSize = &v
	return s
}

type ListCompliancePackTemplatesResponseBody struct {
	CompliancePackTemplatesResult *ListCompliancePackTemplatesResponseBodyCompliancePackTemplatesResult `json:"CompliancePackTemplatesResult,omitempty" xml:"CompliancePackTemplatesResult,omitempty" type:"Struct"`
	RequestId                     *string                                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListCompliancePackTemplatesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListCompliancePackTemplatesResponseBody) GoString() string {
	return s.String()
}

func (s *ListCompliancePackTemplatesResponseBody) SetCompliancePackTemplatesResult(v *ListCompliancePackTemplatesResponseBodyCompliancePackTemplatesResult) *ListCompliancePackTemplatesResponseBody {
	s.CompliancePackTemplatesResult = v
	return s
}

func (s *ListCompliancePackTemplatesResponseBody) SetRequestId(v string) *ListCompliancePackTemplatesResponseBody {
	s.RequestId = &v
	return s
}

type ListCompliancePackTemplatesResponseBodyCompliancePackTemplatesResult struct {
	CompliancePackTemplates []*ListCompliancePackTemplatesResponseBodyCompliancePackTemplatesResultCompliancePackTemplates `json:"CompliancePackTemplates,omitempty" xml:"CompliancePackTemplates,omitempty" type:"Repeated"`
	PageNumber              *int32                                                                                         `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize                *int32                                                                                         `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount              *int64                                                                                         `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListCompliancePackTemplatesResponseBodyCompliancePackTemplatesResult) String() string {
	return tea.Prettify(s)
}

func (s ListCompliancePackTemplatesResponseBodyCompliancePackTemplatesResult) GoString() string {
	return s.String()
}

func (s *ListCompliancePackTemplatesResponseBodyCompliancePackTemplatesResult) SetCompliancePackTemplates(v []*ListCompliancePackTemplatesResponseBodyCompliancePackTemplatesResultCompliancePackTemplates) *ListCompliancePackTemplatesResponseBodyCompliancePackTemplatesResult {
	s.CompliancePackTemplates = v
	return s
}

func (s *ListCompliancePackTemplatesResponseBodyCompliancePackTemplatesResult) SetPageNumber(v int32) *ListCompliancePackTemplatesResponseBodyCompliancePackTemplatesResult {
	s.PageNumber = &v
	return s
}

func (s *ListCompliancePackTemplatesResponseBodyCompliancePackTemplatesResult) SetPageSize(v int32) *ListCompliancePackTemplatesResponseBodyCompliancePackTemplatesResult {
	s.PageSize = &v
	return s
}

func (s *ListCompliancePackTemplatesResponseBodyCompliancePackTemplatesResult) SetTotalCount(v int64) *ListCompliancePackTemplatesResponseBodyCompliancePackTemplatesResult {
	s.TotalCount = &v
	return s
}

type ListCompliancePackTemplatesResponseBodyCompliancePackTemplatesResultCompliancePackTemplates struct {
	CompliancePackTemplateId   *string                                                                                                   `json:"CompliancePackTemplateId,omitempty" xml:"CompliancePackTemplateId,omitempty"`
	CompliancePackTemplateName *string                                                                                                   `json:"CompliancePackTemplateName,omitempty" xml:"CompliancePackTemplateName,omitempty"`
	ConfigRules                []*ListCompliancePackTemplatesResponseBodyCompliancePackTemplatesResultCompliancePackTemplatesConfigRules `json:"ConfigRules,omitempty" xml:"ConfigRules,omitempty" type:"Repeated"`
	Description                *string                                                                                                   `json:"Description,omitempty" xml:"Description,omitempty"`
	RiskLevel                  *int32                                                                                                    `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
}

func (s ListCompliancePackTemplatesResponseBodyCompliancePackTemplatesResultCompliancePackTemplates) String() string {
	return tea.Prettify(s)
}

func (s ListCompliancePackTemplatesResponseBodyCompliancePackTemplatesResultCompliancePackTemplates) GoString() string {
	return s.String()
}

func (s *ListCompliancePackTemplatesResponseBodyCompliancePackTemplatesResultCompliancePackTemplates) SetCompliancePackTemplateId(v string) *ListCompliancePackTemplatesResponseBodyCompliancePackTemplatesResultCompliancePackTemplates {
	s.CompliancePackTemplateId = &v
	return s
}

func (s *ListCompliancePackTemplatesResponseBodyCompliancePackTemplatesResultCompliancePackTemplates) SetCompliancePackTemplateName(v string) *ListCompliancePackTemplatesResponseBodyCompliancePackTemplatesResultCompliancePackTemplates {
	s.CompliancePackTemplateName = &v
	return s
}

func (s *ListCompliancePackTemplatesResponseBodyCompliancePackTemplatesResultCompliancePackTemplates) SetConfigRules(v []*ListCompliancePackTemplatesResponseBodyCompliancePackTemplatesResultCompliancePackTemplatesConfigRules) *ListCompliancePackTemplatesResponseBodyCompliancePackTemplatesResultCompliancePackTemplates {
	s.ConfigRules = v
	return s
}

func (s *ListCompliancePackTemplatesResponseBodyCompliancePackTemplatesResultCompliancePackTemplates) SetDescription(v string) *ListCompliancePackTemplatesResponseBodyCompliancePackTemplatesResultCompliancePackTemplates {
	s.Description = &v
	return s
}

func (s *ListCompliancePackTemplatesResponseBodyCompliancePackTemplatesResultCompliancePackTemplates) SetRiskLevel(v int32) *ListCompliancePackTemplatesResponseBodyCompliancePackTemplatesResultCompliancePackTemplates {
	s.RiskLevel = &v
	return s
}

type ListCompliancePackTemplatesResponseBodyCompliancePackTemplatesResultCompliancePackTemplatesConfigRules struct {
	ConfigRuleParameters  []*ListCompliancePackTemplatesResponseBodyCompliancePackTemplatesResultCompliancePackTemplatesConfigRulesConfigRuleParameters `json:"ConfigRuleParameters,omitempty" xml:"ConfigRuleParameters,omitempty" type:"Repeated"`
	Description           *string                                                                                                                       `json:"Description,omitempty" xml:"Description,omitempty"`
	ManagedRuleIdentifier *string                                                                                                                       `json:"ManagedRuleIdentifier,omitempty" xml:"ManagedRuleIdentifier,omitempty"`
	ManagedRuleName       *string                                                                                                                       `json:"ManagedRuleName,omitempty" xml:"ManagedRuleName,omitempty"`
	RiskLevel             *int32                                                                                                                        `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
}

func (s ListCompliancePackTemplatesResponseBodyCompliancePackTemplatesResultCompliancePackTemplatesConfigRules) String() string {
	return tea.Prettify(s)
}

func (s ListCompliancePackTemplatesResponseBodyCompliancePackTemplatesResultCompliancePackTemplatesConfigRules) GoString() string {
	return s.String()
}

func (s *ListCompliancePackTemplatesResponseBodyCompliancePackTemplatesResultCompliancePackTemplatesConfigRules) SetConfigRuleParameters(v []*ListCompliancePackTemplatesResponseBodyCompliancePackTemplatesResultCompliancePackTemplatesConfigRulesConfigRuleParameters) *ListCompliancePackTemplatesResponseBodyCompliancePackTemplatesResultCompliancePackTemplatesConfigRules {
	s.ConfigRuleParameters = v
	return s
}

func (s *ListCompliancePackTemplatesResponseBodyCompliancePackTemplatesResultCompliancePackTemplatesConfigRules) SetDescription(v string) *ListCompliancePackTemplatesResponseBodyCompliancePackTemplatesResultCompliancePackTemplatesConfigRules {
	s.Description = &v
	return s
}

func (s *ListCompliancePackTemplatesResponseBodyCompliancePackTemplatesResultCompliancePackTemplatesConfigRules) SetManagedRuleIdentifier(v string) *ListCompliancePackTemplatesResponseBodyCompliancePackTemplatesResultCompliancePackTemplatesConfigRules {
	s.ManagedRuleIdentifier = &v
	return s
}

func (s *ListCompliancePackTemplatesResponseBodyCompliancePackTemplatesResultCompliancePackTemplatesConfigRules) SetManagedRuleName(v string) *ListCompliancePackTemplatesResponseBodyCompliancePackTemplatesResultCompliancePackTemplatesConfigRules {
	s.ManagedRuleName = &v
	return s
}

func (s *ListCompliancePackTemplatesResponseBodyCompliancePackTemplatesResultCompliancePackTemplatesConfigRules) SetRiskLevel(v int32) *ListCompliancePackTemplatesResponseBodyCompliancePackTemplatesResultCompliancePackTemplatesConfigRules {
	s.RiskLevel = &v
	return s
}

type ListCompliancePackTemplatesResponseBodyCompliancePackTemplatesResultCompliancePackTemplatesConfigRulesConfigRuleParameters struct {
	ParameterName  *string `json:"ParameterName,omitempty" xml:"ParameterName,omitempty"`
	ParameterValue *string `json:"ParameterValue,omitempty" xml:"ParameterValue,omitempty"`
	Required       *bool   `json:"Required,omitempty" xml:"Required,omitempty"`
}

func (s ListCompliancePackTemplatesResponseBodyCompliancePackTemplatesResultCompliancePackTemplatesConfigRulesConfigRuleParameters) String() string {
	return tea.Prettify(s)
}

func (s ListCompliancePackTemplatesResponseBodyCompliancePackTemplatesResultCompliancePackTemplatesConfigRulesConfigRuleParameters) GoString() string {
	return s.String()
}

func (s *ListCompliancePackTemplatesResponseBodyCompliancePackTemplatesResultCompliancePackTemplatesConfigRulesConfigRuleParameters) SetParameterName(v string) *ListCompliancePackTemplatesResponseBodyCompliancePackTemplatesResultCompliancePackTemplatesConfigRulesConfigRuleParameters {
	s.ParameterName = &v
	return s
}

func (s *ListCompliancePackTemplatesResponseBodyCompliancePackTemplatesResultCompliancePackTemplatesConfigRulesConfigRuleParameters) SetParameterValue(v string) *ListCompliancePackTemplatesResponseBodyCompliancePackTemplatesResultCompliancePackTemplatesConfigRulesConfigRuleParameters {
	s.ParameterValue = &v
	return s
}

func (s *ListCompliancePackTemplatesResponseBodyCompliancePackTemplatesResultCompliancePackTemplatesConfigRulesConfigRuleParameters) SetRequired(v bool) *ListCompliancePackTemplatesResponseBodyCompliancePackTemplatesResultCompliancePackTemplatesConfigRulesConfigRuleParameters {
	s.Required = &v
	return s
}

type ListCompliancePackTemplatesResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListCompliancePackTemplatesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListCompliancePackTemplatesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListCompliancePackTemplatesResponse) GoString() string {
	return s.String()
}

func (s *ListCompliancePackTemplatesResponse) SetHeaders(v map[string]*string) *ListCompliancePackTemplatesResponse {
	s.Headers = v
	return s
}

func (s *ListCompliancePackTemplatesResponse) SetStatusCode(v int32) *ListCompliancePackTemplatesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCompliancePackTemplatesResponse) SetBody(v *ListCompliancePackTemplatesResponseBody) *ListCompliancePackTemplatesResponse {
	s.Body = v
	return s
}

type ListCompliancePacksRequest struct {
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Status     *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListCompliancePacksRequest) String() string {
	return tea.Prettify(s)
}

func (s ListCompliancePacksRequest) GoString() string {
	return s.String()
}

func (s *ListCompliancePacksRequest) SetPageNumber(v int32) *ListCompliancePacksRequest {
	s.PageNumber = &v
	return s
}

func (s *ListCompliancePacksRequest) SetPageSize(v int32) *ListCompliancePacksRequest {
	s.PageSize = &v
	return s
}

func (s *ListCompliancePacksRequest) SetStatus(v string) *ListCompliancePacksRequest {
	s.Status = &v
	return s
}

type ListCompliancePacksResponseBody struct {
	CompliancePacksResult *ListCompliancePacksResponseBodyCompliancePacksResult `json:"CompliancePacksResult,omitempty" xml:"CompliancePacksResult,omitempty" type:"Struct"`
	RequestId             *string                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListCompliancePacksResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListCompliancePacksResponseBody) GoString() string {
	return s.String()
}

func (s *ListCompliancePacksResponseBody) SetCompliancePacksResult(v *ListCompliancePacksResponseBodyCompliancePacksResult) *ListCompliancePacksResponseBody {
	s.CompliancePacksResult = v
	return s
}

func (s *ListCompliancePacksResponseBody) SetRequestId(v string) *ListCompliancePacksResponseBody {
	s.RequestId = &v
	return s
}

type ListCompliancePacksResponseBodyCompliancePacksResult struct {
	CompliancePacks []*ListCompliancePacksResponseBodyCompliancePacksResultCompliancePacks `json:"CompliancePacks,omitempty" xml:"CompliancePacks,omitempty" type:"Repeated"`
	PageNumber      *int32                                                                 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize        *int32                                                                 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount      *int64                                                                 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListCompliancePacksResponseBodyCompliancePacksResult) String() string {
	return tea.Prettify(s)
}

func (s ListCompliancePacksResponseBodyCompliancePacksResult) GoString() string {
	return s.String()
}

func (s *ListCompliancePacksResponseBodyCompliancePacksResult) SetCompliancePacks(v []*ListCompliancePacksResponseBodyCompliancePacksResultCompliancePacks) *ListCompliancePacksResponseBodyCompliancePacksResult {
	s.CompliancePacks = v
	return s
}

func (s *ListCompliancePacksResponseBodyCompliancePacksResult) SetPageNumber(v int32) *ListCompliancePacksResponseBodyCompliancePacksResult {
	s.PageNumber = &v
	return s
}

func (s *ListCompliancePacksResponseBodyCompliancePacksResult) SetPageSize(v int32) *ListCompliancePacksResponseBodyCompliancePacksResult {
	s.PageSize = &v
	return s
}

func (s *ListCompliancePacksResponseBodyCompliancePacksResult) SetTotalCount(v int64) *ListCompliancePacksResponseBodyCompliancePacksResult {
	s.TotalCount = &v
	return s
}

type ListCompliancePacksResponseBodyCompliancePacksResultCompliancePacks struct {
	AccountId                *int64  `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	CompliancePackId         *string `json:"CompliancePackId,omitempty" xml:"CompliancePackId,omitempty"`
	CompliancePackName       *string `json:"CompliancePackName,omitempty" xml:"CompliancePackName,omitempty"`
	CompliancePackTemplateId *string `json:"CompliancePackTemplateId,omitempty" xml:"CompliancePackTemplateId,omitempty"`
	CreateTimestamp          *int64  `json:"CreateTimestamp,omitempty" xml:"CreateTimestamp,omitempty"`
	Description              *string `json:"Description,omitempty" xml:"Description,omitempty"`
	RiskLevel                *int32  `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
	Status                   *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListCompliancePacksResponseBodyCompliancePacksResultCompliancePacks) String() string {
	return tea.Prettify(s)
}

func (s ListCompliancePacksResponseBodyCompliancePacksResultCompliancePacks) GoString() string {
	return s.String()
}

func (s *ListCompliancePacksResponseBodyCompliancePacksResultCompliancePacks) SetAccountId(v int64) *ListCompliancePacksResponseBodyCompliancePacksResultCompliancePacks {
	s.AccountId = &v
	return s
}

func (s *ListCompliancePacksResponseBodyCompliancePacksResultCompliancePacks) SetCompliancePackId(v string) *ListCompliancePacksResponseBodyCompliancePacksResultCompliancePacks {
	s.CompliancePackId = &v
	return s
}

func (s *ListCompliancePacksResponseBodyCompliancePacksResultCompliancePacks) SetCompliancePackName(v string) *ListCompliancePacksResponseBodyCompliancePacksResultCompliancePacks {
	s.CompliancePackName = &v
	return s
}

func (s *ListCompliancePacksResponseBodyCompliancePacksResultCompliancePacks) SetCompliancePackTemplateId(v string) *ListCompliancePacksResponseBodyCompliancePacksResultCompliancePacks {
	s.CompliancePackTemplateId = &v
	return s
}

func (s *ListCompliancePacksResponseBodyCompliancePacksResultCompliancePacks) SetCreateTimestamp(v int64) *ListCompliancePacksResponseBodyCompliancePacksResultCompliancePacks {
	s.CreateTimestamp = &v
	return s
}

func (s *ListCompliancePacksResponseBodyCompliancePacksResultCompliancePacks) SetDescription(v string) *ListCompliancePacksResponseBodyCompliancePacksResultCompliancePacks {
	s.Description = &v
	return s
}

func (s *ListCompliancePacksResponseBodyCompliancePacksResultCompliancePacks) SetRiskLevel(v int32) *ListCompliancePacksResponseBodyCompliancePacksResultCompliancePacks {
	s.RiskLevel = &v
	return s
}

func (s *ListCompliancePacksResponseBodyCompliancePacksResultCompliancePacks) SetStatus(v string) *ListCompliancePacksResponseBodyCompliancePacksResultCompliancePacks {
	s.Status = &v
	return s
}

type ListCompliancePacksResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListCompliancePacksResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListCompliancePacksResponse) String() string {
	return tea.Prettify(s)
}

func (s ListCompliancePacksResponse) GoString() string {
	return s.String()
}

func (s *ListCompliancePacksResponse) SetHeaders(v map[string]*string) *ListCompliancePacksResponse {
	s.Headers = v
	return s
}

func (s *ListCompliancePacksResponse) SetStatusCode(v int32) *ListCompliancePacksResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCompliancePacksResponse) SetBody(v *ListCompliancePacksResponseBody) *ListCompliancePacksResponse {
	s.Body = v
	return s
}

type ListConfigDeliveryChannelsRequest struct {
	DeliveryChannelIds *string `json:"DeliveryChannelIds,omitempty" xml:"DeliveryChannelIds,omitempty"`
}

func (s ListConfigDeliveryChannelsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListConfigDeliveryChannelsRequest) GoString() string {
	return s.String()
}

func (s *ListConfigDeliveryChannelsRequest) SetDeliveryChannelIds(v string) *ListConfigDeliveryChannelsRequest {
	s.DeliveryChannelIds = &v
	return s
}

type ListConfigDeliveryChannelsResponseBody struct {
	DeliveryChannels []*ListConfigDeliveryChannelsResponseBodyDeliveryChannels `json:"DeliveryChannels,omitempty" xml:"DeliveryChannels,omitempty" type:"Repeated"`
	RequestId        *string                                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListConfigDeliveryChannelsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListConfigDeliveryChannelsResponseBody) GoString() string {
	return s.String()
}

func (s *ListConfigDeliveryChannelsResponseBody) SetDeliveryChannels(v []*ListConfigDeliveryChannelsResponseBodyDeliveryChannels) *ListConfigDeliveryChannelsResponseBody {
	s.DeliveryChannels = v
	return s
}

func (s *ListConfigDeliveryChannelsResponseBody) SetRequestId(v string) *ListConfigDeliveryChannelsResponseBody {
	s.RequestId = &v
	return s
}

type ListConfigDeliveryChannelsResponseBodyDeliveryChannels struct {
	AccountId                           *int64  `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	ConfigurationItemChangeNotification *bool   `json:"ConfigurationItemChangeNotification,omitempty" xml:"ConfigurationItemChangeNotification,omitempty"`
	ConfigurationSnapshot               *bool   `json:"ConfigurationSnapshot,omitempty" xml:"ConfigurationSnapshot,omitempty"`
	DeliveryChannelAssumeRoleArn        *string `json:"DeliveryChannelAssumeRoleArn,omitempty" xml:"DeliveryChannelAssumeRoleArn,omitempty"`
	DeliveryChannelCondition            *string `json:"DeliveryChannelCondition,omitempty" xml:"DeliveryChannelCondition,omitempty"`
	DeliveryChannelId                   *string `json:"DeliveryChannelId,omitempty" xml:"DeliveryChannelId,omitempty"`
	DeliveryChannelName                 *string `json:"DeliveryChannelName,omitempty" xml:"DeliveryChannelName,omitempty"`
	DeliveryChannelTargetArn            *string `json:"DeliveryChannelTargetArn,omitempty" xml:"DeliveryChannelTargetArn,omitempty"`
	DeliveryChannelType                 *string `json:"DeliveryChannelType,omitempty" xml:"DeliveryChannelType,omitempty"`
	Description                         *string `json:"Description,omitempty" xml:"Description,omitempty"`
	NonCompliantNotification            *bool   `json:"NonCompliantNotification,omitempty" xml:"NonCompliantNotification,omitempty"`
	OversizedDataOSSTargetArn           *string `json:"OversizedDataOSSTargetArn,omitempty" xml:"OversizedDataOSSTargetArn,omitempty"`
	Status                              *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListConfigDeliveryChannelsResponseBodyDeliveryChannels) String() string {
	return tea.Prettify(s)
}

func (s ListConfigDeliveryChannelsResponseBodyDeliveryChannels) GoString() string {
	return s.String()
}

func (s *ListConfigDeliveryChannelsResponseBodyDeliveryChannels) SetAccountId(v int64) *ListConfigDeliveryChannelsResponseBodyDeliveryChannels {
	s.AccountId = &v
	return s
}

func (s *ListConfigDeliveryChannelsResponseBodyDeliveryChannels) SetConfigurationItemChangeNotification(v bool) *ListConfigDeliveryChannelsResponseBodyDeliveryChannels {
	s.ConfigurationItemChangeNotification = &v
	return s
}

func (s *ListConfigDeliveryChannelsResponseBodyDeliveryChannels) SetConfigurationSnapshot(v bool) *ListConfigDeliveryChannelsResponseBodyDeliveryChannels {
	s.ConfigurationSnapshot = &v
	return s
}

func (s *ListConfigDeliveryChannelsResponseBodyDeliveryChannels) SetDeliveryChannelAssumeRoleArn(v string) *ListConfigDeliveryChannelsResponseBodyDeliveryChannels {
	s.DeliveryChannelAssumeRoleArn = &v
	return s
}

func (s *ListConfigDeliveryChannelsResponseBodyDeliveryChannels) SetDeliveryChannelCondition(v string) *ListConfigDeliveryChannelsResponseBodyDeliveryChannels {
	s.DeliveryChannelCondition = &v
	return s
}

func (s *ListConfigDeliveryChannelsResponseBodyDeliveryChannels) SetDeliveryChannelId(v string) *ListConfigDeliveryChannelsResponseBodyDeliveryChannels {
	s.DeliveryChannelId = &v
	return s
}

func (s *ListConfigDeliveryChannelsResponseBodyDeliveryChannels) SetDeliveryChannelName(v string) *ListConfigDeliveryChannelsResponseBodyDeliveryChannels {
	s.DeliveryChannelName = &v
	return s
}

func (s *ListConfigDeliveryChannelsResponseBodyDeliveryChannels) SetDeliveryChannelTargetArn(v string) *ListConfigDeliveryChannelsResponseBodyDeliveryChannels {
	s.DeliveryChannelTargetArn = &v
	return s
}

func (s *ListConfigDeliveryChannelsResponseBodyDeliveryChannels) SetDeliveryChannelType(v string) *ListConfigDeliveryChannelsResponseBodyDeliveryChannels {
	s.DeliveryChannelType = &v
	return s
}

func (s *ListConfigDeliveryChannelsResponseBodyDeliveryChannels) SetDescription(v string) *ListConfigDeliveryChannelsResponseBodyDeliveryChannels {
	s.Description = &v
	return s
}

func (s *ListConfigDeliveryChannelsResponseBodyDeliveryChannels) SetNonCompliantNotification(v bool) *ListConfigDeliveryChannelsResponseBodyDeliveryChannels {
	s.NonCompliantNotification = &v
	return s
}

func (s *ListConfigDeliveryChannelsResponseBodyDeliveryChannels) SetOversizedDataOSSTargetArn(v string) *ListConfigDeliveryChannelsResponseBodyDeliveryChannels {
	s.OversizedDataOSSTargetArn = &v
	return s
}

func (s *ListConfigDeliveryChannelsResponseBodyDeliveryChannels) SetStatus(v int32) *ListConfigDeliveryChannelsResponseBodyDeliveryChannels {
	s.Status = &v
	return s
}

type ListConfigDeliveryChannelsResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListConfigDeliveryChannelsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListConfigDeliveryChannelsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListConfigDeliveryChannelsResponse) GoString() string {
	return s.String()
}

func (s *ListConfigDeliveryChannelsResponse) SetHeaders(v map[string]*string) *ListConfigDeliveryChannelsResponse {
	s.Headers = v
	return s
}

func (s *ListConfigDeliveryChannelsResponse) SetStatusCode(v int32) *ListConfigDeliveryChannelsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListConfigDeliveryChannelsResponse) SetBody(v *ListConfigDeliveryChannelsResponseBody) *ListConfigDeliveryChannelsResponse {
	s.Body = v
	return s
}

type ListConfigRuleEvaluationResultsRequest struct {
	CompliancePackId *string `json:"CompliancePackId,omitempty" xml:"CompliancePackId,omitempty"`
	ComplianceType   *string `json:"ComplianceType,omitempty" xml:"ComplianceType,omitempty"`
	ConfigRuleId     *string `json:"ConfigRuleId,omitempty" xml:"ConfigRuleId,omitempty"`
	MaxResults       *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken        *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListConfigRuleEvaluationResultsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListConfigRuleEvaluationResultsRequest) GoString() string {
	return s.String()
}

func (s *ListConfigRuleEvaluationResultsRequest) SetCompliancePackId(v string) *ListConfigRuleEvaluationResultsRequest {
	s.CompliancePackId = &v
	return s
}

func (s *ListConfigRuleEvaluationResultsRequest) SetComplianceType(v string) *ListConfigRuleEvaluationResultsRequest {
	s.ComplianceType = &v
	return s
}

func (s *ListConfigRuleEvaluationResultsRequest) SetConfigRuleId(v string) *ListConfigRuleEvaluationResultsRequest {
	s.ConfigRuleId = &v
	return s
}

func (s *ListConfigRuleEvaluationResultsRequest) SetMaxResults(v int32) *ListConfigRuleEvaluationResultsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListConfigRuleEvaluationResultsRequest) SetNextToken(v string) *ListConfigRuleEvaluationResultsRequest {
	s.NextToken = &v
	return s
}

type ListConfigRuleEvaluationResultsResponseBody struct {
	EvaluationResults *ListConfigRuleEvaluationResultsResponseBodyEvaluationResults `json:"EvaluationResults,omitempty" xml:"EvaluationResults,omitempty" type:"Struct"`
	RequestId         *string                                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListConfigRuleEvaluationResultsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListConfigRuleEvaluationResultsResponseBody) GoString() string {
	return s.String()
}

func (s *ListConfigRuleEvaluationResultsResponseBody) SetEvaluationResults(v *ListConfigRuleEvaluationResultsResponseBodyEvaluationResults) *ListConfigRuleEvaluationResultsResponseBody {
	s.EvaluationResults = v
	return s
}

func (s *ListConfigRuleEvaluationResultsResponseBody) SetRequestId(v string) *ListConfigRuleEvaluationResultsResponseBody {
	s.RequestId = &v
	return s
}

type ListConfigRuleEvaluationResultsResponseBodyEvaluationResults struct {
	EvaluationResultList []*ListConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList `json:"EvaluationResultList,omitempty" xml:"EvaluationResultList,omitempty" type:"Repeated"`
	MaxResults           *int32                                                                              `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken            *string                                                                             `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListConfigRuleEvaluationResultsResponseBodyEvaluationResults) String() string {
	return tea.Prettify(s)
}

func (s ListConfigRuleEvaluationResultsResponseBodyEvaluationResults) GoString() string {
	return s.String()
}

func (s *ListConfigRuleEvaluationResultsResponseBodyEvaluationResults) SetEvaluationResultList(v []*ListConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList) *ListConfigRuleEvaluationResultsResponseBodyEvaluationResults {
	s.EvaluationResultList = v
	return s
}

func (s *ListConfigRuleEvaluationResultsResponseBodyEvaluationResults) SetMaxResults(v int32) *ListConfigRuleEvaluationResultsResponseBodyEvaluationResults {
	s.MaxResults = &v
	return s
}

func (s *ListConfigRuleEvaluationResultsResponseBodyEvaluationResults) SetNextToken(v string) *ListConfigRuleEvaluationResultsResponseBodyEvaluationResults {
	s.NextToken = &v
	return s
}

type ListConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList struct {
	Annotation                 *string                                                                                                     `json:"Annotation,omitempty" xml:"Annotation,omitempty"`
	ComplianceType             *string                                                                                                     `json:"ComplianceType,omitempty" xml:"ComplianceType,omitempty"`
	ConfigRuleInvokedTimestamp *int64                                                                                                      `json:"ConfigRuleInvokedTimestamp,omitempty" xml:"ConfigRuleInvokedTimestamp,omitempty"`
	EvaluationResultIdentifier *ListConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifier `json:"EvaluationResultIdentifier,omitempty" xml:"EvaluationResultIdentifier,omitempty" type:"Struct"`
	InvokingEventMessageType   *string                                                                                                     `json:"InvokingEventMessageType,omitempty" xml:"InvokingEventMessageType,omitempty"`
	RemediationEnabled         *bool                                                                                                       `json:"RemediationEnabled,omitempty" xml:"RemediationEnabled,omitempty"`
	ResultRecordedTimestamp    *int64                                                                                                      `json:"ResultRecordedTimestamp,omitempty" xml:"ResultRecordedTimestamp,omitempty"`
	RiskLevel                  *int32                                                                                                      `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
}

func (s ListConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList) String() string {
	return tea.Prettify(s)
}

func (s ListConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList) GoString() string {
	return s.String()
}

func (s *ListConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList) SetAnnotation(v string) *ListConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList {
	s.Annotation = &v
	return s
}

func (s *ListConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList) SetComplianceType(v string) *ListConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList {
	s.ComplianceType = &v
	return s
}

func (s *ListConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList) SetConfigRuleInvokedTimestamp(v int64) *ListConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList {
	s.ConfigRuleInvokedTimestamp = &v
	return s
}

func (s *ListConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList) SetEvaluationResultIdentifier(v *ListConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifier) *ListConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList {
	s.EvaluationResultIdentifier = v
	return s
}

func (s *ListConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList) SetInvokingEventMessageType(v string) *ListConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList {
	s.InvokingEventMessageType = &v
	return s
}

func (s *ListConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList) SetRemediationEnabled(v bool) *ListConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList {
	s.RemediationEnabled = &v
	return s
}

func (s *ListConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList) SetResultRecordedTimestamp(v int64) *ListConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList {
	s.ResultRecordedTimestamp = &v
	return s
}

func (s *ListConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList) SetRiskLevel(v int32) *ListConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList {
	s.RiskLevel = &v
	return s
}

type ListConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifier struct {
	EvaluationResultQualifier *ListConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier `json:"EvaluationResultQualifier,omitempty" xml:"EvaluationResultQualifier,omitempty" type:"Struct"`
	OrderingTimestamp         *int64                                                                                                                               `json:"OrderingTimestamp,omitempty" xml:"OrderingTimestamp,omitempty"`
}

func (s ListConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifier) String() string {
	return tea.Prettify(s)
}

func (s ListConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifier) GoString() string {
	return s.String()
}

func (s *ListConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifier) SetEvaluationResultQualifier(v *ListConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier) *ListConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifier {
	s.EvaluationResultQualifier = v
	return s
}

func (s *ListConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifier) SetOrderingTimestamp(v int64) *ListConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifier {
	s.OrderingTimestamp = &v
	return s
}

type ListConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier struct {
	CompliancePackId *string `json:"CompliancePackId,omitempty" xml:"CompliancePackId,omitempty"`
	ConfigRuleArn    *string `json:"ConfigRuleArn,omitempty" xml:"ConfigRuleArn,omitempty"`
	ConfigRuleId     *string `json:"ConfigRuleId,omitempty" xml:"ConfigRuleId,omitempty"`
	ConfigRuleName   *string `json:"ConfigRuleName,omitempty" xml:"ConfigRuleName,omitempty"`
	IgnoreDate       *string `json:"IgnoreDate,omitempty" xml:"IgnoreDate,omitempty"`
	RegionId         *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceId       *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ResourceName     *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	ResourceOwnerId  *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ResourceType     *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s ListConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier) String() string {
	return tea.Prettify(s)
}

func (s ListConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier) GoString() string {
	return s.String()
}

func (s *ListConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier) SetCompliancePackId(v string) *ListConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier {
	s.CompliancePackId = &v
	return s
}

func (s *ListConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier) SetConfigRuleArn(v string) *ListConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier {
	s.ConfigRuleArn = &v
	return s
}

func (s *ListConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier) SetConfigRuleId(v string) *ListConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier {
	s.ConfigRuleId = &v
	return s
}

func (s *ListConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier) SetConfigRuleName(v string) *ListConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier {
	s.ConfigRuleName = &v
	return s
}

func (s *ListConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier) SetIgnoreDate(v string) *ListConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier {
	s.IgnoreDate = &v
	return s
}

func (s *ListConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier) SetRegionId(v string) *ListConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier {
	s.RegionId = &v
	return s
}

func (s *ListConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier) SetResourceId(v string) *ListConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier {
	s.ResourceId = &v
	return s
}

func (s *ListConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier) SetResourceName(v string) *ListConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier {
	s.ResourceName = &v
	return s
}

func (s *ListConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier) SetResourceOwnerId(v int64) *ListConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier) SetResourceType(v string) *ListConfigRuleEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier {
	s.ResourceType = &v
	return s
}

type ListConfigRuleEvaluationResultsResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListConfigRuleEvaluationResultsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListConfigRuleEvaluationResultsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListConfigRuleEvaluationResultsResponse) GoString() string {
	return s.String()
}

func (s *ListConfigRuleEvaluationResultsResponse) SetHeaders(v map[string]*string) *ListConfigRuleEvaluationResultsResponse {
	s.Headers = v
	return s
}

func (s *ListConfigRuleEvaluationResultsResponse) SetStatusCode(v int32) *ListConfigRuleEvaluationResultsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListConfigRuleEvaluationResultsResponse) SetBody(v *ListConfigRuleEvaluationResultsResponseBody) *ListConfigRuleEvaluationResultsResponse {
	s.Body = v
	return s
}

type ListDiscoveredResourcesRequest struct {
	MaxResults      *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken       *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	Regions         *string `json:"Regions,omitempty" xml:"Regions,omitempty"`
	ResourceDeleted *int32  `json:"ResourceDeleted,omitempty" xml:"ResourceDeleted,omitempty"`
	ResourceId      *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ResourceTypes   *string `json:"ResourceTypes,omitempty" xml:"ResourceTypes,omitempty"`
}

func (s ListDiscoveredResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDiscoveredResourcesRequest) GoString() string {
	return s.String()
}

func (s *ListDiscoveredResourcesRequest) SetMaxResults(v int32) *ListDiscoveredResourcesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListDiscoveredResourcesRequest) SetNextToken(v string) *ListDiscoveredResourcesRequest {
	s.NextToken = &v
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
	DiscoveredResourceProfiles *ListDiscoveredResourcesResponseBodyDiscoveredResourceProfiles `json:"DiscoveredResourceProfiles,omitempty" xml:"DiscoveredResourceProfiles,omitempty" type:"Struct"`
	RequestId                  *string                                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	DiscoveredResourceProfileList []*ListDiscoveredResourcesResponseBodyDiscoveredResourceProfilesDiscoveredResourceProfileList `json:"DiscoveredResourceProfileList,omitempty" xml:"DiscoveredResourceProfileList,omitempty" type:"Repeated"`
	MaxResults                    *int32                                                                                        `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken                     *string                                                                                       `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	TotalCount                    *int32                                                                                        `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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

func (s *ListDiscoveredResourcesResponseBodyDiscoveredResourceProfiles) SetMaxResults(v int32) *ListDiscoveredResourcesResponseBodyDiscoveredResourceProfiles {
	s.MaxResults = &v
	return s
}

func (s *ListDiscoveredResourcesResponseBodyDiscoveredResourceProfiles) SetNextToken(v string) *ListDiscoveredResourcesResponseBodyDiscoveredResourceProfiles {
	s.NextToken = &v
	return s
}

func (s *ListDiscoveredResourcesResponseBodyDiscoveredResourceProfiles) SetTotalCount(v int32) *ListDiscoveredResourcesResponseBodyDiscoveredResourceProfiles {
	s.TotalCount = &v
	return s
}

type ListDiscoveredResourcesResponseBodyDiscoveredResourceProfilesDiscoveredResourceProfileList struct {
	AccountId            *int64  `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	Region               *string `json:"Region,omitempty" xml:"Region,omitempty"`
	ResourceCreationTime *int64  `json:"ResourceCreationTime,omitempty" xml:"ResourceCreationTime,omitempty"`
	ResourceDeleted      *int32  `json:"ResourceDeleted,omitempty" xml:"ResourceDeleted,omitempty"`
	ResourceId           *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ResourceName         *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	ResourceStatus       *string `json:"ResourceStatus,omitempty" xml:"ResourceStatus,omitempty"`
	ResourceType         *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	Tags                 *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
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

type ListManagedRulesRequest struct {
	Keyword    *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RiskLevel  *int32  `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
}

func (s ListManagedRulesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListManagedRulesRequest) GoString() string {
	return s.String()
}

func (s *ListManagedRulesRequest) SetKeyword(v string) *ListManagedRulesRequest {
	s.Keyword = &v
	return s
}

func (s *ListManagedRulesRequest) SetPageNumber(v int32) *ListManagedRulesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListManagedRulesRequest) SetPageSize(v int32) *ListManagedRulesRequest {
	s.PageSize = &v
	return s
}

func (s *ListManagedRulesRequest) SetRiskLevel(v int32) *ListManagedRulesRequest {
	s.RiskLevel = &v
	return s
}

type ListManagedRulesResponseBody struct {
	ManagedRules *ListManagedRulesResponseBodyManagedRules `json:"ManagedRules,omitempty" xml:"ManagedRules,omitempty" type:"Struct"`
	RequestId    *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListManagedRulesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListManagedRulesResponseBody) GoString() string {
	return s.String()
}

func (s *ListManagedRulesResponseBody) SetManagedRules(v *ListManagedRulesResponseBodyManagedRules) *ListManagedRulesResponseBody {
	s.ManagedRules = v
	return s
}

func (s *ListManagedRulesResponseBody) SetRequestId(v string) *ListManagedRulesResponseBody {
	s.RequestId = &v
	return s
}

type ListManagedRulesResponseBodyManagedRules struct {
	ManagedRuleList []*ListManagedRulesResponseBodyManagedRulesManagedRuleList `json:"ManagedRuleList,omitempty" xml:"ManagedRuleList,omitempty" type:"Repeated"`
	PageNumber      *int32                                                     `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize        *int32                                                     `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount      *int64                                                     `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListManagedRulesResponseBodyManagedRules) String() string {
	return tea.Prettify(s)
}

func (s ListManagedRulesResponseBodyManagedRules) GoString() string {
	return s.String()
}

func (s *ListManagedRulesResponseBodyManagedRules) SetManagedRuleList(v []*ListManagedRulesResponseBodyManagedRulesManagedRuleList) *ListManagedRulesResponseBodyManagedRules {
	s.ManagedRuleList = v
	return s
}

func (s *ListManagedRulesResponseBodyManagedRules) SetPageNumber(v int32) *ListManagedRulesResponseBodyManagedRules {
	s.PageNumber = &v
	return s
}

func (s *ListManagedRulesResponseBodyManagedRules) SetPageSize(v int32) *ListManagedRulesResponseBodyManagedRules {
	s.PageSize = &v
	return s
}

func (s *ListManagedRulesResponseBodyManagedRules) SetTotalCount(v int64) *ListManagedRulesResponseBodyManagedRules {
	s.TotalCount = &v
	return s
}

type ListManagedRulesResponseBodyManagedRulesManagedRuleList struct {
	ConfigRuleName *string   `json:"ConfigRuleName,omitempty" xml:"ConfigRuleName,omitempty"`
	Description    *string   `json:"Description,omitempty" xml:"Description,omitempty"`
	HelpUrls       *string   `json:"HelpUrls,omitempty" xml:"HelpUrls,omitempty"`
	Identifier     *string   `json:"Identifier,omitempty" xml:"Identifier,omitempty"`
	Labels         []*string `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	RiskLevel      *int32    `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
}

func (s ListManagedRulesResponseBodyManagedRulesManagedRuleList) String() string {
	return tea.Prettify(s)
}

func (s ListManagedRulesResponseBodyManagedRulesManagedRuleList) GoString() string {
	return s.String()
}

func (s *ListManagedRulesResponseBodyManagedRulesManagedRuleList) SetConfigRuleName(v string) *ListManagedRulesResponseBodyManagedRulesManagedRuleList {
	s.ConfigRuleName = &v
	return s
}

func (s *ListManagedRulesResponseBodyManagedRulesManagedRuleList) SetDescription(v string) *ListManagedRulesResponseBodyManagedRulesManagedRuleList {
	s.Description = &v
	return s
}

func (s *ListManagedRulesResponseBodyManagedRulesManagedRuleList) SetHelpUrls(v string) *ListManagedRulesResponseBodyManagedRulesManagedRuleList {
	s.HelpUrls = &v
	return s
}

func (s *ListManagedRulesResponseBodyManagedRulesManagedRuleList) SetIdentifier(v string) *ListManagedRulesResponseBodyManagedRulesManagedRuleList {
	s.Identifier = &v
	return s
}

func (s *ListManagedRulesResponseBodyManagedRulesManagedRuleList) SetLabels(v []*string) *ListManagedRulesResponseBodyManagedRulesManagedRuleList {
	s.Labels = v
	return s
}

func (s *ListManagedRulesResponseBodyManagedRulesManagedRuleList) SetRiskLevel(v int32) *ListManagedRulesResponseBodyManagedRulesManagedRuleList {
	s.RiskLevel = &v
	return s
}

type ListManagedRulesResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListManagedRulesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListManagedRulesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListManagedRulesResponse) GoString() string {
	return s.String()
}

func (s *ListManagedRulesResponse) SetHeaders(v map[string]*string) *ListManagedRulesResponse {
	s.Headers = v
	return s
}

func (s *ListManagedRulesResponse) SetStatusCode(v int32) *ListManagedRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListManagedRulesResponse) SetBody(v *ListManagedRulesResponseBody) *ListManagedRulesResponse {
	s.Body = v
	return s
}

type ListRemediationTemplatesRequest struct {
	ManagedRuleIdentifier *string `json:"ManagedRuleIdentifier,omitempty" xml:"ManagedRuleIdentifier,omitempty"`
	RemediationType       *string `json:"RemediationType,omitempty" xml:"RemediationType,omitempty"`
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
	RemediationTemplates []*ListRemediationTemplatesResponseBodyRemediationTemplates `json:"RemediationTemplates,omitempty" xml:"RemediationTemplates,omitempty" type:"Repeated"`
	RequestId            *string                                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	RemediationType    *string `json:"RemediationType,omitempty" xml:"RemediationType,omitempty"`
	TemplateDefinition *string `json:"TemplateDefinition,omitempty" xml:"TemplateDefinition,omitempty"`
	TemplateIdentifier *string `json:"TemplateIdentifier,omitempty" xml:"TemplateIdentifier,omitempty"`
	TemplateName       *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
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

type ListRemediationsRequest struct {
	ConfigRuleIds *string `json:"ConfigRuleIds,omitempty" xml:"ConfigRuleIds,omitempty"`
}

func (s ListRemediationsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListRemediationsRequest) GoString() string {
	return s.String()
}

func (s *ListRemediationsRequest) SetConfigRuleIds(v string) *ListRemediationsRequest {
	s.ConfigRuleIds = &v
	return s
}

type ListRemediationsResponseBody struct {
	Remediations []*ListRemediationsResponseBodyRemediations `json:"Remediations,omitempty" xml:"Remediations,omitempty" type:"Repeated"`
	RequestId    *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListRemediationsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListRemediationsResponseBody) GoString() string {
	return s.String()
}

func (s *ListRemediationsResponseBody) SetRemediations(v []*ListRemediationsResponseBodyRemediations) *ListRemediationsResponseBody {
	s.Remediations = v
	return s
}

func (s *ListRemediationsResponseBody) SetRequestId(v string) *ListRemediationsResponseBody {
	s.RequestId = &v
	return s
}

type ListRemediationsResponseBodyRemediations struct {
	AccountId                    *int64  `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	ConfigRuleId                 *string `json:"ConfigRuleId,omitempty" xml:"ConfigRuleId,omitempty"`
	InvokeType                   *string `json:"InvokeType,omitempty" xml:"InvokeType,omitempty"`
	LastSuccessfulInvocationId   *string `json:"LastSuccessfulInvocationId,omitempty" xml:"LastSuccessfulInvocationId,omitempty"`
	LastSuccessfulInvocationTime *int64  `json:"LastSuccessfulInvocationTime,omitempty" xml:"LastSuccessfulInvocationTime,omitempty"`
	LastSuccessfulInvocationType *string `json:"LastSuccessfulInvocationType,omitempty" xml:"LastSuccessfulInvocationType,omitempty"`
	RemediationId                *string `json:"RemediationId,omitempty" xml:"RemediationId,omitempty"`
	RemediationOriginParams      *string `json:"RemediationOriginParams,omitempty" xml:"RemediationOriginParams,omitempty"`
	RemediationSourceType        *string `json:"RemediationSourceType,omitempty" xml:"RemediationSourceType,omitempty"`
	RemediationTemplateId        *string `json:"RemediationTemplateId,omitempty" xml:"RemediationTemplateId,omitempty"`
	RemediationType              *string `json:"RemediationType,omitempty" xml:"RemediationType,omitempty"`
}

func (s ListRemediationsResponseBodyRemediations) String() string {
	return tea.Prettify(s)
}

func (s ListRemediationsResponseBodyRemediations) GoString() string {
	return s.String()
}

func (s *ListRemediationsResponseBodyRemediations) SetAccountId(v int64) *ListRemediationsResponseBodyRemediations {
	s.AccountId = &v
	return s
}

func (s *ListRemediationsResponseBodyRemediations) SetConfigRuleId(v string) *ListRemediationsResponseBodyRemediations {
	s.ConfigRuleId = &v
	return s
}

func (s *ListRemediationsResponseBodyRemediations) SetInvokeType(v string) *ListRemediationsResponseBodyRemediations {
	s.InvokeType = &v
	return s
}

func (s *ListRemediationsResponseBodyRemediations) SetLastSuccessfulInvocationId(v string) *ListRemediationsResponseBodyRemediations {
	s.LastSuccessfulInvocationId = &v
	return s
}

func (s *ListRemediationsResponseBodyRemediations) SetLastSuccessfulInvocationTime(v int64) *ListRemediationsResponseBodyRemediations {
	s.LastSuccessfulInvocationTime = &v
	return s
}

func (s *ListRemediationsResponseBodyRemediations) SetLastSuccessfulInvocationType(v string) *ListRemediationsResponseBodyRemediations {
	s.LastSuccessfulInvocationType = &v
	return s
}

func (s *ListRemediationsResponseBodyRemediations) SetRemediationId(v string) *ListRemediationsResponseBodyRemediations {
	s.RemediationId = &v
	return s
}

func (s *ListRemediationsResponseBodyRemediations) SetRemediationOriginParams(v string) *ListRemediationsResponseBodyRemediations {
	s.RemediationOriginParams = &v
	return s
}

func (s *ListRemediationsResponseBodyRemediations) SetRemediationSourceType(v string) *ListRemediationsResponseBodyRemediations {
	s.RemediationSourceType = &v
	return s
}

func (s *ListRemediationsResponseBodyRemediations) SetRemediationTemplateId(v string) *ListRemediationsResponseBodyRemediations {
	s.RemediationTemplateId = &v
	return s
}

func (s *ListRemediationsResponseBodyRemediations) SetRemediationType(v string) *ListRemediationsResponseBodyRemediations {
	s.RemediationType = &v
	return s
}

type ListRemediationsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListRemediationsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListRemediationsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListRemediationsResponse) GoString() string {
	return s.String()
}

func (s *ListRemediationsResponse) SetHeaders(v map[string]*string) *ListRemediationsResponse {
	s.Headers = v
	return s
}

func (s *ListRemediationsResponse) SetStatusCode(v int32) *ListRemediationsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRemediationsResponse) SetBody(v *ListRemediationsResponseBody) *ListRemediationsResponse {
	s.Body = v
	return s
}

type ListResourceEvaluationResultsRequest struct {
	ComplianceType *string `json:"ComplianceType,omitempty" xml:"ComplianceType,omitempty"`
	MaxResults     *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken      *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	Region         *string `json:"Region,omitempty" xml:"Region,omitempty"`
	ResourceId     *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ResourceType   *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s ListResourceEvaluationResultsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListResourceEvaluationResultsRequest) GoString() string {
	return s.String()
}

func (s *ListResourceEvaluationResultsRequest) SetComplianceType(v string) *ListResourceEvaluationResultsRequest {
	s.ComplianceType = &v
	return s
}

func (s *ListResourceEvaluationResultsRequest) SetMaxResults(v int32) *ListResourceEvaluationResultsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListResourceEvaluationResultsRequest) SetNextToken(v string) *ListResourceEvaluationResultsRequest {
	s.NextToken = &v
	return s
}

func (s *ListResourceEvaluationResultsRequest) SetRegion(v string) *ListResourceEvaluationResultsRequest {
	s.Region = &v
	return s
}

func (s *ListResourceEvaluationResultsRequest) SetResourceId(v string) *ListResourceEvaluationResultsRequest {
	s.ResourceId = &v
	return s
}

func (s *ListResourceEvaluationResultsRequest) SetResourceType(v string) *ListResourceEvaluationResultsRequest {
	s.ResourceType = &v
	return s
}

type ListResourceEvaluationResultsResponseBody struct {
	EvaluationResults *ListResourceEvaluationResultsResponseBodyEvaluationResults `json:"EvaluationResults,omitempty" xml:"EvaluationResults,omitempty" type:"Struct"`
	RequestId         *string                                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListResourceEvaluationResultsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListResourceEvaluationResultsResponseBody) GoString() string {
	return s.String()
}

func (s *ListResourceEvaluationResultsResponseBody) SetEvaluationResults(v *ListResourceEvaluationResultsResponseBodyEvaluationResults) *ListResourceEvaluationResultsResponseBody {
	s.EvaluationResults = v
	return s
}

func (s *ListResourceEvaluationResultsResponseBody) SetRequestId(v string) *ListResourceEvaluationResultsResponseBody {
	s.RequestId = &v
	return s
}

type ListResourceEvaluationResultsResponseBodyEvaluationResults struct {
	EvaluationResultList []*ListResourceEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList `json:"EvaluationResultList,omitempty" xml:"EvaluationResultList,omitempty" type:"Repeated"`
	MaxResults           *int32                                                                            `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken            *string                                                                           `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListResourceEvaluationResultsResponseBodyEvaluationResults) String() string {
	return tea.Prettify(s)
}

func (s ListResourceEvaluationResultsResponseBodyEvaluationResults) GoString() string {
	return s.String()
}

func (s *ListResourceEvaluationResultsResponseBodyEvaluationResults) SetEvaluationResultList(v []*ListResourceEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList) *ListResourceEvaluationResultsResponseBodyEvaluationResults {
	s.EvaluationResultList = v
	return s
}

func (s *ListResourceEvaluationResultsResponseBodyEvaluationResults) SetMaxResults(v int32) *ListResourceEvaluationResultsResponseBodyEvaluationResults {
	s.MaxResults = &v
	return s
}

func (s *ListResourceEvaluationResultsResponseBodyEvaluationResults) SetNextToken(v string) *ListResourceEvaluationResultsResponseBodyEvaluationResults {
	s.NextToken = &v
	return s
}

type ListResourceEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList struct {
	Annotation                 *string                                                                                                   `json:"Annotation,omitempty" xml:"Annotation,omitempty"`
	ComplianceType             *string                                                                                                   `json:"ComplianceType,omitempty" xml:"ComplianceType,omitempty"`
	ConfigRuleInvokedTimestamp *int64                                                                                                    `json:"ConfigRuleInvokedTimestamp,omitempty" xml:"ConfigRuleInvokedTimestamp,omitempty"`
	EvaluationResultIdentifier *ListResourceEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifier `json:"EvaluationResultIdentifier,omitempty" xml:"EvaluationResultIdentifier,omitempty" type:"Struct"`
	InvokingEventMessageType   *string                                                                                                   `json:"InvokingEventMessageType,omitempty" xml:"InvokingEventMessageType,omitempty"`
	RemediationEnabled         *bool                                                                                                     `json:"RemediationEnabled,omitempty" xml:"RemediationEnabled,omitempty"`
	ResultRecordedTimestamp    *int64                                                                                                    `json:"ResultRecordedTimestamp,omitempty" xml:"ResultRecordedTimestamp,omitempty"`
	RiskLevel                  *int32                                                                                                    `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
}

func (s ListResourceEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList) String() string {
	return tea.Prettify(s)
}

func (s ListResourceEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList) GoString() string {
	return s.String()
}

func (s *ListResourceEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList) SetAnnotation(v string) *ListResourceEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList {
	s.Annotation = &v
	return s
}

func (s *ListResourceEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList) SetComplianceType(v string) *ListResourceEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList {
	s.ComplianceType = &v
	return s
}

func (s *ListResourceEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList) SetConfigRuleInvokedTimestamp(v int64) *ListResourceEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList {
	s.ConfigRuleInvokedTimestamp = &v
	return s
}

func (s *ListResourceEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList) SetEvaluationResultIdentifier(v *ListResourceEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifier) *ListResourceEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList {
	s.EvaluationResultIdentifier = v
	return s
}

func (s *ListResourceEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList) SetInvokingEventMessageType(v string) *ListResourceEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList {
	s.InvokingEventMessageType = &v
	return s
}

func (s *ListResourceEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList) SetRemediationEnabled(v bool) *ListResourceEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList {
	s.RemediationEnabled = &v
	return s
}

func (s *ListResourceEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList) SetResultRecordedTimestamp(v int64) *ListResourceEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList {
	s.ResultRecordedTimestamp = &v
	return s
}

func (s *ListResourceEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList) SetRiskLevel(v int32) *ListResourceEvaluationResultsResponseBodyEvaluationResultsEvaluationResultList {
	s.RiskLevel = &v
	return s
}

type ListResourceEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifier struct {
	EvaluationResultQualifier *ListResourceEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier `json:"EvaluationResultQualifier,omitempty" xml:"EvaluationResultQualifier,omitempty" type:"Struct"`
	OrderingTimestamp         *int64                                                                                                                             `json:"OrderingTimestamp,omitempty" xml:"OrderingTimestamp,omitempty"`
}

func (s ListResourceEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifier) String() string {
	return tea.Prettify(s)
}

func (s ListResourceEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifier) GoString() string {
	return s.String()
}

func (s *ListResourceEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifier) SetEvaluationResultQualifier(v *ListResourceEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier) *ListResourceEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifier {
	s.EvaluationResultQualifier = v
	return s
}

func (s *ListResourceEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifier) SetOrderingTimestamp(v int64) *ListResourceEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifier {
	s.OrderingTimestamp = &v
	return s
}

type ListResourceEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier struct {
	ConfigRuleArn  *string `json:"ConfigRuleArn,omitempty" xml:"ConfigRuleArn,omitempty"`
	ConfigRuleId   *string `json:"ConfigRuleId,omitempty" xml:"ConfigRuleId,omitempty"`
	ConfigRuleName *string `json:"ConfigRuleName,omitempty" xml:"ConfigRuleName,omitempty"`
	IgnoreDate     *string `json:"IgnoreDate,omitempty" xml:"IgnoreDate,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceId     *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ResourceName   *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	ResourceType   *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s ListResourceEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier) String() string {
	return tea.Prettify(s)
}

func (s ListResourceEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier) GoString() string {
	return s.String()
}

func (s *ListResourceEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier) SetConfigRuleArn(v string) *ListResourceEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier {
	s.ConfigRuleArn = &v
	return s
}

func (s *ListResourceEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier) SetConfigRuleId(v string) *ListResourceEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier {
	s.ConfigRuleId = &v
	return s
}

func (s *ListResourceEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier) SetConfigRuleName(v string) *ListResourceEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier {
	s.ConfigRuleName = &v
	return s
}

func (s *ListResourceEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier) SetIgnoreDate(v string) *ListResourceEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier {
	s.IgnoreDate = &v
	return s
}

func (s *ListResourceEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier) SetRegionId(v string) *ListResourceEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier {
	s.RegionId = &v
	return s
}

func (s *ListResourceEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier) SetResourceId(v string) *ListResourceEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier {
	s.ResourceId = &v
	return s
}

func (s *ListResourceEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier) SetResourceName(v string) *ListResourceEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier {
	s.ResourceName = &v
	return s
}

func (s *ListResourceEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier) SetResourceType(v string) *ListResourceEvaluationResultsResponseBodyEvaluationResultsEvaluationResultListEvaluationResultIdentifierEvaluationResultQualifier {
	s.ResourceType = &v
	return s
}

type ListResourceEvaluationResultsResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListResourceEvaluationResultsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListResourceEvaluationResultsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListResourceEvaluationResultsResponse) GoString() string {
	return s.String()
}

func (s *ListResourceEvaluationResultsResponse) SetHeaders(v map[string]*string) *ListResourceEvaluationResultsResponse {
	s.Headers = v
	return s
}

func (s *ListResourceEvaluationResultsResponse) SetStatusCode(v int32) *ListResourceEvaluationResultsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListResourceEvaluationResultsResponse) SetBody(v *ListResourceEvaluationResultsResponseBody) *ListResourceEvaluationResultsResponse {
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

type ListTagResourcesShrinkRequest struct {
	NextToken    *string   `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RegionId     *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceId   []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	ResourceType *string   `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	TagShrink    *string   `json:"Tag,omitempty" xml:"Tag,omitempty"`
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

func (s *ListTagResourcesShrinkRequest) SetResourceId(v []*string) *ListTagResourcesShrinkRequest {
	s.ResourceId = v
	return s
}

func (s *ListTagResourcesShrinkRequest) SetResourceType(v string) *ListTagResourcesShrinkRequest {
	s.ResourceType = &v
	return s
}

func (s *ListTagResourcesShrinkRequest) SetTagShrink(v string) *ListTagResourcesShrinkRequest {
	s.TagShrink = &v
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
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListTagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type RevertAggregateEvaluationResultsRequest struct {
	AggregatorId *string                                             `json:"AggregatorId,omitempty" xml:"AggregatorId,omitempty"`
	ConfigRuleId *string                                             `json:"ConfigRuleId,omitempty" xml:"ConfigRuleId,omitempty"`
	Resources    []*RevertAggregateEvaluationResultsRequestResources `json:"Resources,omitempty" xml:"Resources,omitempty" type:"Repeated"`
}

func (s RevertAggregateEvaluationResultsRequest) String() string {
	return tea.Prettify(s)
}

func (s RevertAggregateEvaluationResultsRequest) GoString() string {
	return s.String()
}

func (s *RevertAggregateEvaluationResultsRequest) SetAggregatorId(v string) *RevertAggregateEvaluationResultsRequest {
	s.AggregatorId = &v
	return s
}

func (s *RevertAggregateEvaluationResultsRequest) SetConfigRuleId(v string) *RevertAggregateEvaluationResultsRequest {
	s.ConfigRuleId = &v
	return s
}

func (s *RevertAggregateEvaluationResultsRequest) SetResources(v []*RevertAggregateEvaluationResultsRequestResources) *RevertAggregateEvaluationResultsRequest {
	s.Resources = v
	return s
}

type RevertAggregateEvaluationResultsRequestResources struct {
	Region            *string `json:"Region,omitempty" xml:"Region,omitempty"`
	ResourceAccountId *int64  `json:"ResourceAccountId,omitempty" xml:"ResourceAccountId,omitempty"`
	ResourceId        *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ResourceType      *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s RevertAggregateEvaluationResultsRequestResources) String() string {
	return tea.Prettify(s)
}

func (s RevertAggregateEvaluationResultsRequestResources) GoString() string {
	return s.String()
}

func (s *RevertAggregateEvaluationResultsRequestResources) SetRegion(v string) *RevertAggregateEvaluationResultsRequestResources {
	s.Region = &v
	return s
}

func (s *RevertAggregateEvaluationResultsRequestResources) SetResourceAccountId(v int64) *RevertAggregateEvaluationResultsRequestResources {
	s.ResourceAccountId = &v
	return s
}

func (s *RevertAggregateEvaluationResultsRequestResources) SetResourceId(v string) *RevertAggregateEvaluationResultsRequestResources {
	s.ResourceId = &v
	return s
}

func (s *RevertAggregateEvaluationResultsRequestResources) SetResourceType(v string) *RevertAggregateEvaluationResultsRequestResources {
	s.ResourceType = &v
	return s
}

type RevertAggregateEvaluationResultsShrinkRequest struct {
	AggregatorId    *string `json:"AggregatorId,omitempty" xml:"AggregatorId,omitempty"`
	ConfigRuleId    *string `json:"ConfigRuleId,omitempty" xml:"ConfigRuleId,omitempty"`
	ResourcesShrink *string `json:"Resources,omitempty" xml:"Resources,omitempty"`
}

func (s RevertAggregateEvaluationResultsShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s RevertAggregateEvaluationResultsShrinkRequest) GoString() string {
	return s.String()
}

func (s *RevertAggregateEvaluationResultsShrinkRequest) SetAggregatorId(v string) *RevertAggregateEvaluationResultsShrinkRequest {
	s.AggregatorId = &v
	return s
}

func (s *RevertAggregateEvaluationResultsShrinkRequest) SetConfigRuleId(v string) *RevertAggregateEvaluationResultsShrinkRequest {
	s.ConfigRuleId = &v
	return s
}

func (s *RevertAggregateEvaluationResultsShrinkRequest) SetResourcesShrink(v string) *RevertAggregateEvaluationResultsShrinkRequest {
	s.ResourcesShrink = &v
	return s
}

type RevertAggregateEvaluationResultsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RevertAggregateEvaluationResultsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RevertAggregateEvaluationResultsResponseBody) GoString() string {
	return s.String()
}

func (s *RevertAggregateEvaluationResultsResponseBody) SetRequestId(v string) *RevertAggregateEvaluationResultsResponseBody {
	s.RequestId = &v
	return s
}

type RevertAggregateEvaluationResultsResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RevertAggregateEvaluationResultsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RevertAggregateEvaluationResultsResponse) String() string {
	return tea.Prettify(s)
}

func (s RevertAggregateEvaluationResultsResponse) GoString() string {
	return s.String()
}

func (s *RevertAggregateEvaluationResultsResponse) SetHeaders(v map[string]*string) *RevertAggregateEvaluationResultsResponse {
	s.Headers = v
	return s
}

func (s *RevertAggregateEvaluationResultsResponse) SetStatusCode(v int32) *RevertAggregateEvaluationResultsResponse {
	s.StatusCode = &v
	return s
}

func (s *RevertAggregateEvaluationResultsResponse) SetBody(v *RevertAggregateEvaluationResultsResponseBody) *RevertAggregateEvaluationResultsResponse {
	s.Body = v
	return s
}

type RevertEvaluationResultsRequest struct {
	ConfigRuleId *string                                    `json:"ConfigRuleId,omitempty" xml:"ConfigRuleId,omitempty"`
	Resources    []*RevertEvaluationResultsRequestResources `json:"Resources,omitempty" xml:"Resources,omitempty" type:"Repeated"`
}

func (s RevertEvaluationResultsRequest) String() string {
	return tea.Prettify(s)
}

func (s RevertEvaluationResultsRequest) GoString() string {
	return s.String()
}

func (s *RevertEvaluationResultsRequest) SetConfigRuleId(v string) *RevertEvaluationResultsRequest {
	s.ConfigRuleId = &v
	return s
}

func (s *RevertEvaluationResultsRequest) SetResources(v []*RevertEvaluationResultsRequestResources) *RevertEvaluationResultsRequest {
	s.Resources = v
	return s
}

type RevertEvaluationResultsRequestResources struct {
	Region            *string `json:"Region,omitempty" xml:"Region,omitempty"`
	ResourceAccountId *int64  `json:"ResourceAccountId,omitempty" xml:"ResourceAccountId,omitempty"`
	ResourceId        *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ResourceType      *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s RevertEvaluationResultsRequestResources) String() string {
	return tea.Prettify(s)
}

func (s RevertEvaluationResultsRequestResources) GoString() string {
	return s.String()
}

func (s *RevertEvaluationResultsRequestResources) SetRegion(v string) *RevertEvaluationResultsRequestResources {
	s.Region = &v
	return s
}

func (s *RevertEvaluationResultsRequestResources) SetResourceAccountId(v int64) *RevertEvaluationResultsRequestResources {
	s.ResourceAccountId = &v
	return s
}

func (s *RevertEvaluationResultsRequestResources) SetResourceId(v string) *RevertEvaluationResultsRequestResources {
	s.ResourceId = &v
	return s
}

func (s *RevertEvaluationResultsRequestResources) SetResourceType(v string) *RevertEvaluationResultsRequestResources {
	s.ResourceType = &v
	return s
}

type RevertEvaluationResultsShrinkRequest struct {
	ConfigRuleId    *string `json:"ConfigRuleId,omitempty" xml:"ConfigRuleId,omitempty"`
	ResourcesShrink *string `json:"Resources,omitempty" xml:"Resources,omitempty"`
}

func (s RevertEvaluationResultsShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s RevertEvaluationResultsShrinkRequest) GoString() string {
	return s.String()
}

func (s *RevertEvaluationResultsShrinkRequest) SetConfigRuleId(v string) *RevertEvaluationResultsShrinkRequest {
	s.ConfigRuleId = &v
	return s
}

func (s *RevertEvaluationResultsShrinkRequest) SetResourcesShrink(v string) *RevertEvaluationResultsShrinkRequest {
	s.ResourcesShrink = &v
	return s
}

type RevertEvaluationResultsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RevertEvaluationResultsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RevertEvaluationResultsResponseBody) GoString() string {
	return s.String()
}

func (s *RevertEvaluationResultsResponseBody) SetRequestId(v string) *RevertEvaluationResultsResponseBody {
	s.RequestId = &v
	return s
}

type RevertEvaluationResultsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RevertEvaluationResultsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RevertEvaluationResultsResponse) String() string {
	return tea.Prettify(s)
}

func (s RevertEvaluationResultsResponse) GoString() string {
	return s.String()
}

func (s *RevertEvaluationResultsResponse) SetHeaders(v map[string]*string) *RevertEvaluationResultsResponse {
	s.Headers = v
	return s
}

func (s *RevertEvaluationResultsResponse) SetStatusCode(v int32) *RevertEvaluationResultsResponse {
	s.StatusCode = &v
	return s
}

func (s *RevertEvaluationResultsResponse) SetBody(v *RevertEvaluationResultsResponseBody) *RevertEvaluationResultsResponse {
	s.Body = v
	return s
}

type StartAggregateConfigRuleEvaluationRequest struct {
	AggregatorId     *string `json:"AggregatorId,omitempty" xml:"AggregatorId,omitempty"`
	CompliancePackId *string `json:"CompliancePackId,omitempty" xml:"CompliancePackId,omitempty"`
	ConfigRuleId     *string `json:"ConfigRuleId,omitempty" xml:"ConfigRuleId,omitempty"`
	RevertEvaluation *bool   `json:"RevertEvaluation,omitempty" xml:"RevertEvaluation,omitempty"`
}

func (s StartAggregateConfigRuleEvaluationRequest) String() string {
	return tea.Prettify(s)
}

func (s StartAggregateConfigRuleEvaluationRequest) GoString() string {
	return s.String()
}

func (s *StartAggregateConfigRuleEvaluationRequest) SetAggregatorId(v string) *StartAggregateConfigRuleEvaluationRequest {
	s.AggregatorId = &v
	return s
}

func (s *StartAggregateConfigRuleEvaluationRequest) SetCompliancePackId(v string) *StartAggregateConfigRuleEvaluationRequest {
	s.CompliancePackId = &v
	return s
}

func (s *StartAggregateConfigRuleEvaluationRequest) SetConfigRuleId(v string) *StartAggregateConfigRuleEvaluationRequest {
	s.ConfigRuleId = &v
	return s
}

func (s *StartAggregateConfigRuleEvaluationRequest) SetRevertEvaluation(v bool) *StartAggregateConfigRuleEvaluationRequest {
	s.RevertEvaluation = &v
	return s
}

type StartAggregateConfigRuleEvaluationResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *bool   `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s StartAggregateConfigRuleEvaluationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StartAggregateConfigRuleEvaluationResponseBody) GoString() string {
	return s.String()
}

func (s *StartAggregateConfigRuleEvaluationResponseBody) SetRequestId(v string) *StartAggregateConfigRuleEvaluationResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartAggregateConfigRuleEvaluationResponseBody) SetResult(v bool) *StartAggregateConfigRuleEvaluationResponseBody {
	s.Result = &v
	return s
}

type StartAggregateConfigRuleEvaluationResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *StartAggregateConfigRuleEvaluationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StartAggregateConfigRuleEvaluationResponse) String() string {
	return tea.Prettify(s)
}

func (s StartAggregateConfigRuleEvaluationResponse) GoString() string {
	return s.String()
}

func (s *StartAggregateConfigRuleEvaluationResponse) SetHeaders(v map[string]*string) *StartAggregateConfigRuleEvaluationResponse {
	s.Headers = v
	return s
}

func (s *StartAggregateConfigRuleEvaluationResponse) SetStatusCode(v int32) *StartAggregateConfigRuleEvaluationResponse {
	s.StatusCode = &v
	return s
}

func (s *StartAggregateConfigRuleEvaluationResponse) SetBody(v *StartAggregateConfigRuleEvaluationResponseBody) *StartAggregateConfigRuleEvaluationResponse {
	s.Body = v
	return s
}

type StartAggregateRemediationRequest struct {
	AggregatorId *string `json:"AggregatorId,omitempty" xml:"AggregatorId,omitempty"`
	ConfigRuleId *string `json:"ConfigRuleId,omitempty" xml:"ConfigRuleId,omitempty"`
}

func (s StartAggregateRemediationRequest) String() string {
	return tea.Prettify(s)
}

func (s StartAggregateRemediationRequest) GoString() string {
	return s.String()
}

func (s *StartAggregateRemediationRequest) SetAggregatorId(v string) *StartAggregateRemediationRequest {
	s.AggregatorId = &v
	return s
}

func (s *StartAggregateRemediationRequest) SetConfigRuleId(v string) *StartAggregateRemediationRequest {
	s.ConfigRuleId = &v
	return s
}

type StartAggregateRemediationResponseBody struct {
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StartAggregateRemediationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StartAggregateRemediationResponseBody) GoString() string {
	return s.String()
}

func (s *StartAggregateRemediationResponseBody) SetData(v bool) *StartAggregateRemediationResponseBody {
	s.Data = &v
	return s
}

func (s *StartAggregateRemediationResponseBody) SetRequestId(v string) *StartAggregateRemediationResponseBody {
	s.RequestId = &v
	return s
}

type StartAggregateRemediationResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *StartAggregateRemediationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StartAggregateRemediationResponse) String() string {
	return tea.Prettify(s)
}

func (s StartAggregateRemediationResponse) GoString() string {
	return s.String()
}

func (s *StartAggregateRemediationResponse) SetHeaders(v map[string]*string) *StartAggregateRemediationResponse {
	s.Headers = v
	return s
}

func (s *StartAggregateRemediationResponse) SetStatusCode(v int32) *StartAggregateRemediationResponse {
	s.StatusCode = &v
	return s
}

func (s *StartAggregateRemediationResponse) SetBody(v *StartAggregateRemediationResponseBody) *StartAggregateRemediationResponse {
	s.Body = v
	return s
}

type StartRemediationRequest struct {
	ConfigRuleId *string `json:"ConfigRuleId,omitempty" xml:"ConfigRuleId,omitempty"`
}

func (s StartRemediationRequest) String() string {
	return tea.Prettify(s)
}

func (s StartRemediationRequest) GoString() string {
	return s.String()
}

func (s *StartRemediationRequest) SetConfigRuleId(v string) *StartRemediationRequest {
	s.ConfigRuleId = &v
	return s
}

type StartRemediationResponseBody struct {
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StartRemediationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StartRemediationResponseBody) GoString() string {
	return s.String()
}

func (s *StartRemediationResponseBody) SetData(v bool) *StartRemediationResponseBody {
	s.Data = &v
	return s
}

func (s *StartRemediationResponseBody) SetRequestId(v string) *StartRemediationResponseBody {
	s.RequestId = &v
	return s
}

type StartRemediationResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *StartRemediationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StartRemediationResponse) String() string {
	return tea.Prettify(s)
}

func (s StartRemediationResponse) GoString() string {
	return s.String()
}

func (s *StartRemediationResponse) SetHeaders(v map[string]*string) *StartRemediationResponse {
	s.Headers = v
	return s
}

func (s *StartRemediationResponse) SetStatusCode(v int32) *StartRemediationResponse {
	s.StatusCode = &v
	return s
}

func (s *StartRemediationResponse) SetBody(v *StartRemediationResponseBody) *StartRemediationResponse {
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

type TagResourcesShrinkRequest struct {
	RegionId     *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceId   []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	ResourceType *string   `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	TagShrink    *string   `json:"Tag,omitempty" xml:"Tag,omitempty"`
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

func (s *TagResourcesShrinkRequest) SetResourceId(v []*string) *TagResourcesShrinkRequest {
	s.ResourceId = v
	return s
}

func (s *TagResourcesShrinkRequest) SetResourceType(v string) *TagResourcesShrinkRequest {
	s.ResourceType = &v
	return s
}

func (s *TagResourcesShrinkRequest) SetTagShrink(v string) *TagResourcesShrinkRequest {
	s.TagShrink = &v
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
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *TagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UntagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type UpdateAggregateCompliancePackRequest struct {
	AggregatorId       *string                                            `json:"AggregatorId,omitempty" xml:"AggregatorId,omitempty"`
	ClientToken        *string                                            `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	CompliancePackId   *string                                            `json:"CompliancePackId,omitempty" xml:"CompliancePackId,omitempty"`
	CompliancePackName *string                                            `json:"CompliancePackName,omitempty" xml:"CompliancePackName,omitempty"`
	ConfigRules        []*UpdateAggregateCompliancePackRequestConfigRules `json:"ConfigRules,omitempty" xml:"ConfigRules,omitempty" type:"Repeated"`
	Description        *string                                            `json:"Description,omitempty" xml:"Description,omitempty"`
	RiskLevel          *int32                                             `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
}

func (s UpdateAggregateCompliancePackRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateAggregateCompliancePackRequest) GoString() string {
	return s.String()
}

func (s *UpdateAggregateCompliancePackRequest) SetAggregatorId(v string) *UpdateAggregateCompliancePackRequest {
	s.AggregatorId = &v
	return s
}

func (s *UpdateAggregateCompliancePackRequest) SetClientToken(v string) *UpdateAggregateCompliancePackRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateAggregateCompliancePackRequest) SetCompliancePackId(v string) *UpdateAggregateCompliancePackRequest {
	s.CompliancePackId = &v
	return s
}

func (s *UpdateAggregateCompliancePackRequest) SetCompliancePackName(v string) *UpdateAggregateCompliancePackRequest {
	s.CompliancePackName = &v
	return s
}

func (s *UpdateAggregateCompliancePackRequest) SetConfigRules(v []*UpdateAggregateCompliancePackRequestConfigRules) *UpdateAggregateCompliancePackRequest {
	s.ConfigRules = v
	return s
}

func (s *UpdateAggregateCompliancePackRequest) SetDescription(v string) *UpdateAggregateCompliancePackRequest {
	s.Description = &v
	return s
}

func (s *UpdateAggregateCompliancePackRequest) SetRiskLevel(v int32) *UpdateAggregateCompliancePackRequest {
	s.RiskLevel = &v
	return s
}

type UpdateAggregateCompliancePackRequestConfigRules struct {
	ConfigRuleId          *string                                                                `json:"ConfigRuleId,omitempty" xml:"ConfigRuleId,omitempty"`
	ConfigRuleName        *string                                                                `json:"ConfigRuleName,omitempty" xml:"ConfigRuleName,omitempty"`
	ConfigRuleParameters  []*UpdateAggregateCompliancePackRequestConfigRulesConfigRuleParameters `json:"ConfigRuleParameters,omitempty" xml:"ConfigRuleParameters,omitempty" type:"Repeated"`
	Description           *string                                                                `json:"Description,omitempty" xml:"Description,omitempty"`
	ManagedRuleIdentifier *string                                                                `json:"ManagedRuleIdentifier,omitempty" xml:"ManagedRuleIdentifier,omitempty"`
	RiskLevel             *int32                                                                 `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
}

func (s UpdateAggregateCompliancePackRequestConfigRules) String() string {
	return tea.Prettify(s)
}

func (s UpdateAggregateCompliancePackRequestConfigRules) GoString() string {
	return s.String()
}

func (s *UpdateAggregateCompliancePackRequestConfigRules) SetConfigRuleId(v string) *UpdateAggregateCompliancePackRequestConfigRules {
	s.ConfigRuleId = &v
	return s
}

func (s *UpdateAggregateCompliancePackRequestConfigRules) SetConfigRuleName(v string) *UpdateAggregateCompliancePackRequestConfigRules {
	s.ConfigRuleName = &v
	return s
}

func (s *UpdateAggregateCompliancePackRequestConfigRules) SetConfigRuleParameters(v []*UpdateAggregateCompliancePackRequestConfigRulesConfigRuleParameters) *UpdateAggregateCompliancePackRequestConfigRules {
	s.ConfigRuleParameters = v
	return s
}

func (s *UpdateAggregateCompliancePackRequestConfigRules) SetDescription(v string) *UpdateAggregateCompliancePackRequestConfigRules {
	s.Description = &v
	return s
}

func (s *UpdateAggregateCompliancePackRequestConfigRules) SetManagedRuleIdentifier(v string) *UpdateAggregateCompliancePackRequestConfigRules {
	s.ManagedRuleIdentifier = &v
	return s
}

func (s *UpdateAggregateCompliancePackRequestConfigRules) SetRiskLevel(v int32) *UpdateAggregateCompliancePackRequestConfigRules {
	s.RiskLevel = &v
	return s
}

type UpdateAggregateCompliancePackRequestConfigRulesConfigRuleParameters struct {
	ParameterName  *string `json:"ParameterName,omitempty" xml:"ParameterName,omitempty"`
	ParameterValue *string `json:"ParameterValue,omitempty" xml:"ParameterValue,omitempty"`
}

func (s UpdateAggregateCompliancePackRequestConfigRulesConfigRuleParameters) String() string {
	return tea.Prettify(s)
}

func (s UpdateAggregateCompliancePackRequestConfigRulesConfigRuleParameters) GoString() string {
	return s.String()
}

func (s *UpdateAggregateCompliancePackRequestConfigRulesConfigRuleParameters) SetParameterName(v string) *UpdateAggregateCompliancePackRequestConfigRulesConfigRuleParameters {
	s.ParameterName = &v
	return s
}

func (s *UpdateAggregateCompliancePackRequestConfigRulesConfigRuleParameters) SetParameterValue(v string) *UpdateAggregateCompliancePackRequestConfigRulesConfigRuleParameters {
	s.ParameterValue = &v
	return s
}

type UpdateAggregateCompliancePackShrinkRequest struct {
	AggregatorId       *string `json:"AggregatorId,omitempty" xml:"AggregatorId,omitempty"`
	ClientToken        *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	CompliancePackId   *string `json:"CompliancePackId,omitempty" xml:"CompliancePackId,omitempty"`
	CompliancePackName *string `json:"CompliancePackName,omitempty" xml:"CompliancePackName,omitempty"`
	ConfigRulesShrink  *string `json:"ConfigRules,omitempty" xml:"ConfigRules,omitempty"`
	Description        *string `json:"Description,omitempty" xml:"Description,omitempty"`
	RiskLevel          *int32  `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
}

func (s UpdateAggregateCompliancePackShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateAggregateCompliancePackShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateAggregateCompliancePackShrinkRequest) SetAggregatorId(v string) *UpdateAggregateCompliancePackShrinkRequest {
	s.AggregatorId = &v
	return s
}

func (s *UpdateAggregateCompliancePackShrinkRequest) SetClientToken(v string) *UpdateAggregateCompliancePackShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateAggregateCompliancePackShrinkRequest) SetCompliancePackId(v string) *UpdateAggregateCompliancePackShrinkRequest {
	s.CompliancePackId = &v
	return s
}

func (s *UpdateAggregateCompliancePackShrinkRequest) SetCompliancePackName(v string) *UpdateAggregateCompliancePackShrinkRequest {
	s.CompliancePackName = &v
	return s
}

func (s *UpdateAggregateCompliancePackShrinkRequest) SetConfigRulesShrink(v string) *UpdateAggregateCompliancePackShrinkRequest {
	s.ConfigRulesShrink = &v
	return s
}

func (s *UpdateAggregateCompliancePackShrinkRequest) SetDescription(v string) *UpdateAggregateCompliancePackShrinkRequest {
	s.Description = &v
	return s
}

func (s *UpdateAggregateCompliancePackShrinkRequest) SetRiskLevel(v int32) *UpdateAggregateCompliancePackShrinkRequest {
	s.RiskLevel = &v
	return s
}

type UpdateAggregateCompliancePackResponseBody struct {
	CompliancePackId *string `json:"CompliancePackId,omitempty" xml:"CompliancePackId,omitempty"`
	RequestId        *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateAggregateCompliancePackResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateAggregateCompliancePackResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAggregateCompliancePackResponseBody) SetCompliancePackId(v string) *UpdateAggregateCompliancePackResponseBody {
	s.CompliancePackId = &v
	return s
}

func (s *UpdateAggregateCompliancePackResponseBody) SetRequestId(v string) *UpdateAggregateCompliancePackResponseBody {
	s.RequestId = &v
	return s
}

type UpdateAggregateCompliancePackResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateAggregateCompliancePackResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateAggregateCompliancePackResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateAggregateCompliancePackResponse) GoString() string {
	return s.String()
}

func (s *UpdateAggregateCompliancePackResponse) SetHeaders(v map[string]*string) *UpdateAggregateCompliancePackResponse {
	s.Headers = v
	return s
}

func (s *UpdateAggregateCompliancePackResponse) SetStatusCode(v int32) *UpdateAggregateCompliancePackResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAggregateCompliancePackResponse) SetBody(v *UpdateAggregateCompliancePackResponseBody) *UpdateAggregateCompliancePackResponse {
	s.Body = v
	return s
}

type UpdateAggregateConfigDeliveryChannelRequest struct {
	AggregatorId                        *string `json:"AggregatorId,omitempty" xml:"AggregatorId,omitempty"`
	ClientToken                         *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	ConfigurationItemChangeNotification *bool   `json:"ConfigurationItemChangeNotification,omitempty" xml:"ConfigurationItemChangeNotification,omitempty"`
	ConfigurationSnapshot               *bool   `json:"ConfigurationSnapshot,omitempty" xml:"ConfigurationSnapshot,omitempty"`
	DeliveryChannelCondition            *string `json:"DeliveryChannelCondition,omitempty" xml:"DeliveryChannelCondition,omitempty"`
	DeliveryChannelId                   *string `json:"DeliveryChannelId,omitempty" xml:"DeliveryChannelId,omitempty"`
	DeliveryChannelName                 *string `json:"DeliveryChannelName,omitempty" xml:"DeliveryChannelName,omitempty"`
	DeliveryChannelTargetArn            *string `json:"DeliveryChannelTargetArn,omitempty" xml:"DeliveryChannelTargetArn,omitempty"`
	Description                         *string `json:"Description,omitempty" xml:"Description,omitempty"`
	NonCompliantNotification            *bool   `json:"NonCompliantNotification,omitempty" xml:"NonCompliantNotification,omitempty"`
	OversizedDataOSSTargetArn           *string `json:"OversizedDataOSSTargetArn,omitempty" xml:"OversizedDataOSSTargetArn,omitempty"`
	Status                              *int64  `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s UpdateAggregateConfigDeliveryChannelRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateAggregateConfigDeliveryChannelRequest) GoString() string {
	return s.String()
}

func (s *UpdateAggregateConfigDeliveryChannelRequest) SetAggregatorId(v string) *UpdateAggregateConfigDeliveryChannelRequest {
	s.AggregatorId = &v
	return s
}

func (s *UpdateAggregateConfigDeliveryChannelRequest) SetClientToken(v string) *UpdateAggregateConfigDeliveryChannelRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateAggregateConfigDeliveryChannelRequest) SetConfigurationItemChangeNotification(v bool) *UpdateAggregateConfigDeliveryChannelRequest {
	s.ConfigurationItemChangeNotification = &v
	return s
}

func (s *UpdateAggregateConfigDeliveryChannelRequest) SetConfigurationSnapshot(v bool) *UpdateAggregateConfigDeliveryChannelRequest {
	s.ConfigurationSnapshot = &v
	return s
}

func (s *UpdateAggregateConfigDeliveryChannelRequest) SetDeliveryChannelCondition(v string) *UpdateAggregateConfigDeliveryChannelRequest {
	s.DeliveryChannelCondition = &v
	return s
}

func (s *UpdateAggregateConfigDeliveryChannelRequest) SetDeliveryChannelId(v string) *UpdateAggregateConfigDeliveryChannelRequest {
	s.DeliveryChannelId = &v
	return s
}

func (s *UpdateAggregateConfigDeliveryChannelRequest) SetDeliveryChannelName(v string) *UpdateAggregateConfigDeliveryChannelRequest {
	s.DeliveryChannelName = &v
	return s
}

func (s *UpdateAggregateConfigDeliveryChannelRequest) SetDeliveryChannelTargetArn(v string) *UpdateAggregateConfigDeliveryChannelRequest {
	s.DeliveryChannelTargetArn = &v
	return s
}

func (s *UpdateAggregateConfigDeliveryChannelRequest) SetDescription(v string) *UpdateAggregateConfigDeliveryChannelRequest {
	s.Description = &v
	return s
}

func (s *UpdateAggregateConfigDeliveryChannelRequest) SetNonCompliantNotification(v bool) *UpdateAggregateConfigDeliveryChannelRequest {
	s.NonCompliantNotification = &v
	return s
}

func (s *UpdateAggregateConfigDeliveryChannelRequest) SetOversizedDataOSSTargetArn(v string) *UpdateAggregateConfigDeliveryChannelRequest {
	s.OversizedDataOSSTargetArn = &v
	return s
}

func (s *UpdateAggregateConfigDeliveryChannelRequest) SetStatus(v int64) *UpdateAggregateConfigDeliveryChannelRequest {
	s.Status = &v
	return s
}

type UpdateAggregateConfigDeliveryChannelResponseBody struct {
	DeliveryChannelId *string `json:"DeliveryChannelId,omitempty" xml:"DeliveryChannelId,omitempty"`
	RequestId         *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateAggregateConfigDeliveryChannelResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateAggregateConfigDeliveryChannelResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAggregateConfigDeliveryChannelResponseBody) SetDeliveryChannelId(v string) *UpdateAggregateConfigDeliveryChannelResponseBody {
	s.DeliveryChannelId = &v
	return s
}

func (s *UpdateAggregateConfigDeliveryChannelResponseBody) SetRequestId(v string) *UpdateAggregateConfigDeliveryChannelResponseBody {
	s.RequestId = &v
	return s
}

type UpdateAggregateConfigDeliveryChannelResponse struct {
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateAggregateConfigDeliveryChannelResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateAggregateConfigDeliveryChannelResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateAggregateConfigDeliveryChannelResponse) GoString() string {
	return s.String()
}

func (s *UpdateAggregateConfigDeliveryChannelResponse) SetHeaders(v map[string]*string) *UpdateAggregateConfigDeliveryChannelResponse {
	s.Headers = v
	return s
}

func (s *UpdateAggregateConfigDeliveryChannelResponse) SetStatusCode(v int32) *UpdateAggregateConfigDeliveryChannelResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAggregateConfigDeliveryChannelResponse) SetBody(v *UpdateAggregateConfigDeliveryChannelResponseBody) *UpdateAggregateConfigDeliveryChannelResponse {
	s.Body = v
	return s
}

type UpdateAggregateConfigRuleRequest struct {
	AggregatorId              *string                `json:"AggregatorId,omitempty" xml:"AggregatorId,omitempty"`
	ClientToken               *string                `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	ConfigRuleId              *string                `json:"ConfigRuleId,omitempty" xml:"ConfigRuleId,omitempty"`
	ConfigRuleName            *string                `json:"ConfigRuleName,omitempty" xml:"ConfigRuleName,omitempty"`
	ConfigRuleTriggerTypes    *string                `json:"ConfigRuleTriggerTypes,omitempty" xml:"ConfigRuleTriggerTypes,omitempty"`
	Description               *string                `json:"Description,omitempty" xml:"Description,omitempty"`
	ExcludeAccountIdsScope    *string                `json:"ExcludeAccountIdsScope,omitempty" xml:"ExcludeAccountIdsScope,omitempty"`
	ExcludeFolderIdsScope     *string                `json:"ExcludeFolderIdsScope,omitempty" xml:"ExcludeFolderIdsScope,omitempty"`
	ExcludeResourceIdsScope   *string                `json:"ExcludeResourceIdsScope,omitempty" xml:"ExcludeResourceIdsScope,omitempty"`
	FolderIdsScope            *string                `json:"FolderIdsScope,omitempty" xml:"FolderIdsScope,omitempty"`
	InputParameters           map[string]interface{} `json:"InputParameters,omitempty" xml:"InputParameters,omitempty"`
	MaximumExecutionFrequency *string                `json:"MaximumExecutionFrequency,omitempty" xml:"MaximumExecutionFrequency,omitempty"`
	RegionIdsScope            *string                `json:"RegionIdsScope,omitempty" xml:"RegionIdsScope,omitempty"`
	ResourceGroupIdsScope     *string                `json:"ResourceGroupIdsScope,omitempty" xml:"ResourceGroupIdsScope,omitempty"`
	ResourceTypesScope        []*string              `json:"ResourceTypesScope,omitempty" xml:"ResourceTypesScope,omitempty" type:"Repeated"`
	RiskLevel                 *int32                 `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
	TagKeyLogicScope          *string                `json:"TagKeyLogicScope,omitempty" xml:"TagKeyLogicScope,omitempty"`
	TagKeyScope               *string                `json:"TagKeyScope,omitempty" xml:"TagKeyScope,omitempty"`
	TagValueScope             *string                `json:"TagValueScope,omitempty" xml:"TagValueScope,omitempty"`
}

func (s UpdateAggregateConfigRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateAggregateConfigRuleRequest) GoString() string {
	return s.String()
}

func (s *UpdateAggregateConfigRuleRequest) SetAggregatorId(v string) *UpdateAggregateConfigRuleRequest {
	s.AggregatorId = &v
	return s
}

func (s *UpdateAggregateConfigRuleRequest) SetClientToken(v string) *UpdateAggregateConfigRuleRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateAggregateConfigRuleRequest) SetConfigRuleId(v string) *UpdateAggregateConfigRuleRequest {
	s.ConfigRuleId = &v
	return s
}

func (s *UpdateAggregateConfigRuleRequest) SetConfigRuleName(v string) *UpdateAggregateConfigRuleRequest {
	s.ConfigRuleName = &v
	return s
}

func (s *UpdateAggregateConfigRuleRequest) SetConfigRuleTriggerTypes(v string) *UpdateAggregateConfigRuleRequest {
	s.ConfigRuleTriggerTypes = &v
	return s
}

func (s *UpdateAggregateConfigRuleRequest) SetDescription(v string) *UpdateAggregateConfigRuleRequest {
	s.Description = &v
	return s
}

func (s *UpdateAggregateConfigRuleRequest) SetExcludeAccountIdsScope(v string) *UpdateAggregateConfigRuleRequest {
	s.ExcludeAccountIdsScope = &v
	return s
}

func (s *UpdateAggregateConfigRuleRequest) SetExcludeFolderIdsScope(v string) *UpdateAggregateConfigRuleRequest {
	s.ExcludeFolderIdsScope = &v
	return s
}

func (s *UpdateAggregateConfigRuleRequest) SetExcludeResourceIdsScope(v string) *UpdateAggregateConfigRuleRequest {
	s.ExcludeResourceIdsScope = &v
	return s
}

func (s *UpdateAggregateConfigRuleRequest) SetFolderIdsScope(v string) *UpdateAggregateConfigRuleRequest {
	s.FolderIdsScope = &v
	return s
}

func (s *UpdateAggregateConfigRuleRequest) SetInputParameters(v map[string]interface{}) *UpdateAggregateConfigRuleRequest {
	s.InputParameters = v
	return s
}

func (s *UpdateAggregateConfigRuleRequest) SetMaximumExecutionFrequency(v string) *UpdateAggregateConfigRuleRequest {
	s.MaximumExecutionFrequency = &v
	return s
}

func (s *UpdateAggregateConfigRuleRequest) SetRegionIdsScope(v string) *UpdateAggregateConfigRuleRequest {
	s.RegionIdsScope = &v
	return s
}

func (s *UpdateAggregateConfigRuleRequest) SetResourceGroupIdsScope(v string) *UpdateAggregateConfigRuleRequest {
	s.ResourceGroupIdsScope = &v
	return s
}

func (s *UpdateAggregateConfigRuleRequest) SetResourceTypesScope(v []*string) *UpdateAggregateConfigRuleRequest {
	s.ResourceTypesScope = v
	return s
}

func (s *UpdateAggregateConfigRuleRequest) SetRiskLevel(v int32) *UpdateAggregateConfigRuleRequest {
	s.RiskLevel = &v
	return s
}

func (s *UpdateAggregateConfigRuleRequest) SetTagKeyLogicScope(v string) *UpdateAggregateConfigRuleRequest {
	s.TagKeyLogicScope = &v
	return s
}

func (s *UpdateAggregateConfigRuleRequest) SetTagKeyScope(v string) *UpdateAggregateConfigRuleRequest {
	s.TagKeyScope = &v
	return s
}

func (s *UpdateAggregateConfigRuleRequest) SetTagValueScope(v string) *UpdateAggregateConfigRuleRequest {
	s.TagValueScope = &v
	return s
}

type UpdateAggregateConfigRuleShrinkRequest struct {
	AggregatorId              *string `json:"AggregatorId,omitempty" xml:"AggregatorId,omitempty"`
	ClientToken               *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	ConfigRuleId              *string `json:"ConfigRuleId,omitempty" xml:"ConfigRuleId,omitempty"`
	ConfigRuleName            *string `json:"ConfigRuleName,omitempty" xml:"ConfigRuleName,omitempty"`
	ConfigRuleTriggerTypes    *string `json:"ConfigRuleTriggerTypes,omitempty" xml:"ConfigRuleTriggerTypes,omitempty"`
	Description               *string `json:"Description,omitempty" xml:"Description,omitempty"`
	ExcludeAccountIdsScope    *string `json:"ExcludeAccountIdsScope,omitempty" xml:"ExcludeAccountIdsScope,omitempty"`
	ExcludeFolderIdsScope     *string `json:"ExcludeFolderIdsScope,omitempty" xml:"ExcludeFolderIdsScope,omitempty"`
	ExcludeResourceIdsScope   *string `json:"ExcludeResourceIdsScope,omitempty" xml:"ExcludeResourceIdsScope,omitempty"`
	FolderIdsScope            *string `json:"FolderIdsScope,omitempty" xml:"FolderIdsScope,omitempty"`
	InputParametersShrink     *string `json:"InputParameters,omitempty" xml:"InputParameters,omitempty"`
	MaximumExecutionFrequency *string `json:"MaximumExecutionFrequency,omitempty" xml:"MaximumExecutionFrequency,omitempty"`
	RegionIdsScope            *string `json:"RegionIdsScope,omitempty" xml:"RegionIdsScope,omitempty"`
	ResourceGroupIdsScope     *string `json:"ResourceGroupIdsScope,omitempty" xml:"ResourceGroupIdsScope,omitempty"`
	ResourceTypesScopeShrink  *string `json:"ResourceTypesScope,omitempty" xml:"ResourceTypesScope,omitempty"`
	RiskLevel                 *int32  `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
	TagKeyLogicScope          *string `json:"TagKeyLogicScope,omitempty" xml:"TagKeyLogicScope,omitempty"`
	TagKeyScope               *string `json:"TagKeyScope,omitempty" xml:"TagKeyScope,omitempty"`
	TagValueScope             *string `json:"TagValueScope,omitempty" xml:"TagValueScope,omitempty"`
}

func (s UpdateAggregateConfigRuleShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateAggregateConfigRuleShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateAggregateConfigRuleShrinkRequest) SetAggregatorId(v string) *UpdateAggregateConfigRuleShrinkRequest {
	s.AggregatorId = &v
	return s
}

func (s *UpdateAggregateConfigRuleShrinkRequest) SetClientToken(v string) *UpdateAggregateConfigRuleShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateAggregateConfigRuleShrinkRequest) SetConfigRuleId(v string) *UpdateAggregateConfigRuleShrinkRequest {
	s.ConfigRuleId = &v
	return s
}

func (s *UpdateAggregateConfigRuleShrinkRequest) SetConfigRuleName(v string) *UpdateAggregateConfigRuleShrinkRequest {
	s.ConfigRuleName = &v
	return s
}

func (s *UpdateAggregateConfigRuleShrinkRequest) SetConfigRuleTriggerTypes(v string) *UpdateAggregateConfigRuleShrinkRequest {
	s.ConfigRuleTriggerTypes = &v
	return s
}

func (s *UpdateAggregateConfigRuleShrinkRequest) SetDescription(v string) *UpdateAggregateConfigRuleShrinkRequest {
	s.Description = &v
	return s
}

func (s *UpdateAggregateConfigRuleShrinkRequest) SetExcludeAccountIdsScope(v string) *UpdateAggregateConfigRuleShrinkRequest {
	s.ExcludeAccountIdsScope = &v
	return s
}

func (s *UpdateAggregateConfigRuleShrinkRequest) SetExcludeFolderIdsScope(v string) *UpdateAggregateConfigRuleShrinkRequest {
	s.ExcludeFolderIdsScope = &v
	return s
}

func (s *UpdateAggregateConfigRuleShrinkRequest) SetExcludeResourceIdsScope(v string) *UpdateAggregateConfigRuleShrinkRequest {
	s.ExcludeResourceIdsScope = &v
	return s
}

func (s *UpdateAggregateConfigRuleShrinkRequest) SetFolderIdsScope(v string) *UpdateAggregateConfigRuleShrinkRequest {
	s.FolderIdsScope = &v
	return s
}

func (s *UpdateAggregateConfigRuleShrinkRequest) SetInputParametersShrink(v string) *UpdateAggregateConfigRuleShrinkRequest {
	s.InputParametersShrink = &v
	return s
}

func (s *UpdateAggregateConfigRuleShrinkRequest) SetMaximumExecutionFrequency(v string) *UpdateAggregateConfigRuleShrinkRequest {
	s.MaximumExecutionFrequency = &v
	return s
}

func (s *UpdateAggregateConfigRuleShrinkRequest) SetRegionIdsScope(v string) *UpdateAggregateConfigRuleShrinkRequest {
	s.RegionIdsScope = &v
	return s
}

func (s *UpdateAggregateConfigRuleShrinkRequest) SetResourceGroupIdsScope(v string) *UpdateAggregateConfigRuleShrinkRequest {
	s.ResourceGroupIdsScope = &v
	return s
}

func (s *UpdateAggregateConfigRuleShrinkRequest) SetResourceTypesScopeShrink(v string) *UpdateAggregateConfigRuleShrinkRequest {
	s.ResourceTypesScopeShrink = &v
	return s
}

func (s *UpdateAggregateConfigRuleShrinkRequest) SetRiskLevel(v int32) *UpdateAggregateConfigRuleShrinkRequest {
	s.RiskLevel = &v
	return s
}

func (s *UpdateAggregateConfigRuleShrinkRequest) SetTagKeyLogicScope(v string) *UpdateAggregateConfigRuleShrinkRequest {
	s.TagKeyLogicScope = &v
	return s
}

func (s *UpdateAggregateConfigRuleShrinkRequest) SetTagKeyScope(v string) *UpdateAggregateConfigRuleShrinkRequest {
	s.TagKeyScope = &v
	return s
}

func (s *UpdateAggregateConfigRuleShrinkRequest) SetTagValueScope(v string) *UpdateAggregateConfigRuleShrinkRequest {
	s.TagValueScope = &v
	return s
}

type UpdateAggregateConfigRuleResponseBody struct {
	ConfigRuleId *string `json:"ConfigRuleId,omitempty" xml:"ConfigRuleId,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateAggregateConfigRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateAggregateConfigRuleResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAggregateConfigRuleResponseBody) SetConfigRuleId(v string) *UpdateAggregateConfigRuleResponseBody {
	s.ConfigRuleId = &v
	return s
}

func (s *UpdateAggregateConfigRuleResponseBody) SetRequestId(v string) *UpdateAggregateConfigRuleResponseBody {
	s.RequestId = &v
	return s
}

type UpdateAggregateConfigRuleResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateAggregateConfigRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateAggregateConfigRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateAggregateConfigRuleResponse) GoString() string {
	return s.String()
}

func (s *UpdateAggregateConfigRuleResponse) SetHeaders(v map[string]*string) *UpdateAggregateConfigRuleResponse {
	s.Headers = v
	return s
}

func (s *UpdateAggregateConfigRuleResponse) SetStatusCode(v int32) *UpdateAggregateConfigRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAggregateConfigRuleResponse) SetBody(v *UpdateAggregateConfigRuleResponseBody) *UpdateAggregateConfigRuleResponse {
	s.Body = v
	return s
}

type UpdateAggregateRemediationRequest struct {
	AggregatorId          *string `json:"AggregatorId,omitempty" xml:"AggregatorId,omitempty"`
	InvokeType            *string `json:"InvokeType,omitempty" xml:"InvokeType,omitempty"`
	Params                *string `json:"Params,omitempty" xml:"Params,omitempty"`
	RemediationId         *string `json:"RemediationId,omitempty" xml:"RemediationId,omitempty"`
	RemediationTemplateId *string `json:"RemediationTemplateId,omitempty" xml:"RemediationTemplateId,omitempty"`
	RemediationType       *string `json:"RemediationType,omitempty" xml:"RemediationType,omitempty"`
	SourceType            *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
}

func (s UpdateAggregateRemediationRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateAggregateRemediationRequest) GoString() string {
	return s.String()
}

func (s *UpdateAggregateRemediationRequest) SetAggregatorId(v string) *UpdateAggregateRemediationRequest {
	s.AggregatorId = &v
	return s
}

func (s *UpdateAggregateRemediationRequest) SetInvokeType(v string) *UpdateAggregateRemediationRequest {
	s.InvokeType = &v
	return s
}

func (s *UpdateAggregateRemediationRequest) SetParams(v string) *UpdateAggregateRemediationRequest {
	s.Params = &v
	return s
}

func (s *UpdateAggregateRemediationRequest) SetRemediationId(v string) *UpdateAggregateRemediationRequest {
	s.RemediationId = &v
	return s
}

func (s *UpdateAggregateRemediationRequest) SetRemediationTemplateId(v string) *UpdateAggregateRemediationRequest {
	s.RemediationTemplateId = &v
	return s
}

func (s *UpdateAggregateRemediationRequest) SetRemediationType(v string) *UpdateAggregateRemediationRequest {
	s.RemediationType = &v
	return s
}

func (s *UpdateAggregateRemediationRequest) SetSourceType(v string) *UpdateAggregateRemediationRequest {
	s.SourceType = &v
	return s
}

type UpdateAggregateRemediationResponseBody struct {
	RemediationId *string `json:"RemediationId,omitempty" xml:"RemediationId,omitempty"`
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateAggregateRemediationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateAggregateRemediationResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAggregateRemediationResponseBody) SetRemediationId(v string) *UpdateAggregateRemediationResponseBody {
	s.RemediationId = &v
	return s
}

func (s *UpdateAggregateRemediationResponseBody) SetRequestId(v string) *UpdateAggregateRemediationResponseBody {
	s.RequestId = &v
	return s
}

type UpdateAggregateRemediationResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateAggregateRemediationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateAggregateRemediationResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateAggregateRemediationResponse) GoString() string {
	return s.String()
}

func (s *UpdateAggregateRemediationResponse) SetHeaders(v map[string]*string) *UpdateAggregateRemediationResponse {
	s.Headers = v
	return s
}

func (s *UpdateAggregateRemediationResponse) SetStatusCode(v int32) *UpdateAggregateRemediationResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAggregateRemediationResponse) SetBody(v *UpdateAggregateRemediationResponseBody) *UpdateAggregateRemediationResponse {
	s.Body = v
	return s
}

type UpdateAggregatorRequest struct {
	AggregatorAccounts []*UpdateAggregatorRequestAggregatorAccounts `json:"AggregatorAccounts,omitempty" xml:"AggregatorAccounts,omitempty" type:"Repeated"`
	AggregatorId       *string                                      `json:"AggregatorId,omitempty" xml:"AggregatorId,omitempty"`
	AggregatorName     *string                                      `json:"AggregatorName,omitempty" xml:"AggregatorName,omitempty"`
	ClientToken        *string                                      `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Description        *string                                      `json:"Description,omitempty" xml:"Description,omitempty"`
}

func (s UpdateAggregatorRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateAggregatorRequest) GoString() string {
	return s.String()
}

func (s *UpdateAggregatorRequest) SetAggregatorAccounts(v []*UpdateAggregatorRequestAggregatorAccounts) *UpdateAggregatorRequest {
	s.AggregatorAccounts = v
	return s
}

func (s *UpdateAggregatorRequest) SetAggregatorId(v string) *UpdateAggregatorRequest {
	s.AggregatorId = &v
	return s
}

func (s *UpdateAggregatorRequest) SetAggregatorName(v string) *UpdateAggregatorRequest {
	s.AggregatorName = &v
	return s
}

func (s *UpdateAggregatorRequest) SetClientToken(v string) *UpdateAggregatorRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateAggregatorRequest) SetDescription(v string) *UpdateAggregatorRequest {
	s.Description = &v
	return s
}

type UpdateAggregatorRequestAggregatorAccounts struct {
	AccountId   *int64  `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	AccountType *string `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
}

func (s UpdateAggregatorRequestAggregatorAccounts) String() string {
	return tea.Prettify(s)
}

func (s UpdateAggregatorRequestAggregatorAccounts) GoString() string {
	return s.String()
}

func (s *UpdateAggregatorRequestAggregatorAccounts) SetAccountId(v int64) *UpdateAggregatorRequestAggregatorAccounts {
	s.AccountId = &v
	return s
}

func (s *UpdateAggregatorRequestAggregatorAccounts) SetAccountName(v string) *UpdateAggregatorRequestAggregatorAccounts {
	s.AccountName = &v
	return s
}

func (s *UpdateAggregatorRequestAggregatorAccounts) SetAccountType(v string) *UpdateAggregatorRequestAggregatorAccounts {
	s.AccountType = &v
	return s
}

type UpdateAggregatorShrinkRequest struct {
	AggregatorAccountsShrink *string `json:"AggregatorAccounts,omitempty" xml:"AggregatorAccounts,omitempty"`
	AggregatorId             *string `json:"AggregatorId,omitempty" xml:"AggregatorId,omitempty"`
	AggregatorName           *string `json:"AggregatorName,omitempty" xml:"AggregatorName,omitempty"`
	ClientToken              *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Description              *string `json:"Description,omitempty" xml:"Description,omitempty"`
}

func (s UpdateAggregatorShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateAggregatorShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateAggregatorShrinkRequest) SetAggregatorAccountsShrink(v string) *UpdateAggregatorShrinkRequest {
	s.AggregatorAccountsShrink = &v
	return s
}

func (s *UpdateAggregatorShrinkRequest) SetAggregatorId(v string) *UpdateAggregatorShrinkRequest {
	s.AggregatorId = &v
	return s
}

func (s *UpdateAggregatorShrinkRequest) SetAggregatorName(v string) *UpdateAggregatorShrinkRequest {
	s.AggregatorName = &v
	return s
}

func (s *UpdateAggregatorShrinkRequest) SetClientToken(v string) *UpdateAggregatorShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateAggregatorShrinkRequest) SetDescription(v string) *UpdateAggregatorShrinkRequest {
	s.Description = &v
	return s
}

type UpdateAggregatorResponseBody struct {
	AggregatorId *string `json:"AggregatorId,omitempty" xml:"AggregatorId,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateAggregatorResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateAggregatorResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAggregatorResponseBody) SetAggregatorId(v string) *UpdateAggregatorResponseBody {
	s.AggregatorId = &v
	return s
}

func (s *UpdateAggregatorResponseBody) SetRequestId(v string) *UpdateAggregatorResponseBody {
	s.RequestId = &v
	return s
}

type UpdateAggregatorResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateAggregatorResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateAggregatorResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateAggregatorResponse) GoString() string {
	return s.String()
}

func (s *UpdateAggregatorResponse) SetHeaders(v map[string]*string) *UpdateAggregatorResponse {
	s.Headers = v
	return s
}

func (s *UpdateAggregatorResponse) SetStatusCode(v int32) *UpdateAggregatorResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAggregatorResponse) SetBody(v *UpdateAggregatorResponseBody) *UpdateAggregatorResponse {
	s.Body = v
	return s
}

type UpdateCompliancePackRequest struct {
	ClientToken        *string                                   `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	CompliancePackId   *string                                   `json:"CompliancePackId,omitempty" xml:"CompliancePackId,omitempty"`
	CompliancePackName *string                                   `json:"CompliancePackName,omitempty" xml:"CompliancePackName,omitempty"`
	ConfigRules        []*UpdateCompliancePackRequestConfigRules `json:"ConfigRules,omitempty" xml:"ConfigRules,omitempty" type:"Repeated"`
	Description        *string                                   `json:"Description,omitempty" xml:"Description,omitempty"`
	RiskLevel          *int32                                    `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
}

func (s UpdateCompliancePackRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateCompliancePackRequest) GoString() string {
	return s.String()
}

func (s *UpdateCompliancePackRequest) SetClientToken(v string) *UpdateCompliancePackRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateCompliancePackRequest) SetCompliancePackId(v string) *UpdateCompliancePackRequest {
	s.CompliancePackId = &v
	return s
}

func (s *UpdateCompliancePackRequest) SetCompliancePackName(v string) *UpdateCompliancePackRequest {
	s.CompliancePackName = &v
	return s
}

func (s *UpdateCompliancePackRequest) SetConfigRules(v []*UpdateCompliancePackRequestConfigRules) *UpdateCompliancePackRequest {
	s.ConfigRules = v
	return s
}

func (s *UpdateCompliancePackRequest) SetDescription(v string) *UpdateCompliancePackRequest {
	s.Description = &v
	return s
}

func (s *UpdateCompliancePackRequest) SetRiskLevel(v int32) *UpdateCompliancePackRequest {
	s.RiskLevel = &v
	return s
}

type UpdateCompliancePackRequestConfigRules struct {
	ConfigRuleId          *string                                                       `json:"ConfigRuleId,omitempty" xml:"ConfigRuleId,omitempty"`
	ConfigRuleName        *string                                                       `json:"ConfigRuleName,omitempty" xml:"ConfigRuleName,omitempty"`
	ConfigRuleParameters  []*UpdateCompliancePackRequestConfigRulesConfigRuleParameters `json:"ConfigRuleParameters,omitempty" xml:"ConfigRuleParameters,omitempty" type:"Repeated"`
	Description           *string                                                       `json:"Description,omitempty" xml:"Description,omitempty"`
	ManagedRuleIdentifier *string                                                       `json:"ManagedRuleIdentifier,omitempty" xml:"ManagedRuleIdentifier,omitempty"`
	RiskLevel             *int32                                                        `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
}

func (s UpdateCompliancePackRequestConfigRules) String() string {
	return tea.Prettify(s)
}

func (s UpdateCompliancePackRequestConfigRules) GoString() string {
	return s.String()
}

func (s *UpdateCompliancePackRequestConfigRules) SetConfigRuleId(v string) *UpdateCompliancePackRequestConfigRules {
	s.ConfigRuleId = &v
	return s
}

func (s *UpdateCompliancePackRequestConfigRules) SetConfigRuleName(v string) *UpdateCompliancePackRequestConfigRules {
	s.ConfigRuleName = &v
	return s
}

func (s *UpdateCompliancePackRequestConfigRules) SetConfigRuleParameters(v []*UpdateCompliancePackRequestConfigRulesConfigRuleParameters) *UpdateCompliancePackRequestConfigRules {
	s.ConfigRuleParameters = v
	return s
}

func (s *UpdateCompliancePackRequestConfigRules) SetDescription(v string) *UpdateCompliancePackRequestConfigRules {
	s.Description = &v
	return s
}

func (s *UpdateCompliancePackRequestConfigRules) SetManagedRuleIdentifier(v string) *UpdateCompliancePackRequestConfigRules {
	s.ManagedRuleIdentifier = &v
	return s
}

func (s *UpdateCompliancePackRequestConfigRules) SetRiskLevel(v int32) *UpdateCompliancePackRequestConfigRules {
	s.RiskLevel = &v
	return s
}

type UpdateCompliancePackRequestConfigRulesConfigRuleParameters struct {
	ParameterName  *string `json:"ParameterName,omitempty" xml:"ParameterName,omitempty"`
	ParameterValue *string `json:"ParameterValue,omitempty" xml:"ParameterValue,omitempty"`
}

func (s UpdateCompliancePackRequestConfigRulesConfigRuleParameters) String() string {
	return tea.Prettify(s)
}

func (s UpdateCompliancePackRequestConfigRulesConfigRuleParameters) GoString() string {
	return s.String()
}

func (s *UpdateCompliancePackRequestConfigRulesConfigRuleParameters) SetParameterName(v string) *UpdateCompliancePackRequestConfigRulesConfigRuleParameters {
	s.ParameterName = &v
	return s
}

func (s *UpdateCompliancePackRequestConfigRulesConfigRuleParameters) SetParameterValue(v string) *UpdateCompliancePackRequestConfigRulesConfigRuleParameters {
	s.ParameterValue = &v
	return s
}

type UpdateCompliancePackShrinkRequest struct {
	ClientToken        *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	CompliancePackId   *string `json:"CompliancePackId,omitempty" xml:"CompliancePackId,omitempty"`
	CompliancePackName *string `json:"CompliancePackName,omitempty" xml:"CompliancePackName,omitempty"`
	ConfigRulesShrink  *string `json:"ConfigRules,omitempty" xml:"ConfigRules,omitempty"`
	Description        *string `json:"Description,omitempty" xml:"Description,omitempty"`
	RiskLevel          *int32  `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
}

func (s UpdateCompliancePackShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateCompliancePackShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateCompliancePackShrinkRequest) SetClientToken(v string) *UpdateCompliancePackShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateCompliancePackShrinkRequest) SetCompliancePackId(v string) *UpdateCompliancePackShrinkRequest {
	s.CompliancePackId = &v
	return s
}

func (s *UpdateCompliancePackShrinkRequest) SetCompliancePackName(v string) *UpdateCompliancePackShrinkRequest {
	s.CompliancePackName = &v
	return s
}

func (s *UpdateCompliancePackShrinkRequest) SetConfigRulesShrink(v string) *UpdateCompliancePackShrinkRequest {
	s.ConfigRulesShrink = &v
	return s
}

func (s *UpdateCompliancePackShrinkRequest) SetDescription(v string) *UpdateCompliancePackShrinkRequest {
	s.Description = &v
	return s
}

func (s *UpdateCompliancePackShrinkRequest) SetRiskLevel(v int32) *UpdateCompliancePackShrinkRequest {
	s.RiskLevel = &v
	return s
}

type UpdateCompliancePackResponseBody struct {
	CompliancePackId *string `json:"CompliancePackId,omitempty" xml:"CompliancePackId,omitempty"`
	RequestId        *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateCompliancePackResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateCompliancePackResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateCompliancePackResponseBody) SetCompliancePackId(v string) *UpdateCompliancePackResponseBody {
	s.CompliancePackId = &v
	return s
}

func (s *UpdateCompliancePackResponseBody) SetRequestId(v string) *UpdateCompliancePackResponseBody {
	s.RequestId = &v
	return s
}

type UpdateCompliancePackResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateCompliancePackResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateCompliancePackResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateCompliancePackResponse) GoString() string {
	return s.String()
}

func (s *UpdateCompliancePackResponse) SetHeaders(v map[string]*string) *UpdateCompliancePackResponse {
	s.Headers = v
	return s
}

func (s *UpdateCompliancePackResponse) SetStatusCode(v int32) *UpdateCompliancePackResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateCompliancePackResponse) SetBody(v *UpdateCompliancePackResponseBody) *UpdateCompliancePackResponse {
	s.Body = v
	return s
}

type UpdateConfigDeliveryChannelRequest struct {
	ClientToken                         *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	ConfigurationItemChangeNotification *bool   `json:"ConfigurationItemChangeNotification,omitempty" xml:"ConfigurationItemChangeNotification,omitempty"`
	ConfigurationSnapshot               *bool   `json:"ConfigurationSnapshot,omitempty" xml:"ConfigurationSnapshot,omitempty"`
	DeliveryChannelCondition            *string `json:"DeliveryChannelCondition,omitempty" xml:"DeliveryChannelCondition,omitempty"`
	DeliveryChannelId                   *string `json:"DeliveryChannelId,omitempty" xml:"DeliveryChannelId,omitempty"`
	DeliveryChannelName                 *string `json:"DeliveryChannelName,omitempty" xml:"DeliveryChannelName,omitempty"`
	DeliveryChannelTargetArn            *string `json:"DeliveryChannelTargetArn,omitempty" xml:"DeliveryChannelTargetArn,omitempty"`
	Description                         *string `json:"Description,omitempty" xml:"Description,omitempty"`
	NonCompliantNotification            *bool   `json:"NonCompliantNotification,omitempty" xml:"NonCompliantNotification,omitempty"`
	OversizedDataOSSTargetArn           *string `json:"OversizedDataOSSTargetArn,omitempty" xml:"OversizedDataOSSTargetArn,omitempty"`
	Status                              *int64  `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s UpdateConfigDeliveryChannelRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateConfigDeliveryChannelRequest) GoString() string {
	return s.String()
}

func (s *UpdateConfigDeliveryChannelRequest) SetClientToken(v string) *UpdateConfigDeliveryChannelRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateConfigDeliveryChannelRequest) SetConfigurationItemChangeNotification(v bool) *UpdateConfigDeliveryChannelRequest {
	s.ConfigurationItemChangeNotification = &v
	return s
}

func (s *UpdateConfigDeliveryChannelRequest) SetConfigurationSnapshot(v bool) *UpdateConfigDeliveryChannelRequest {
	s.ConfigurationSnapshot = &v
	return s
}

func (s *UpdateConfigDeliveryChannelRequest) SetDeliveryChannelCondition(v string) *UpdateConfigDeliveryChannelRequest {
	s.DeliveryChannelCondition = &v
	return s
}

func (s *UpdateConfigDeliveryChannelRequest) SetDeliveryChannelId(v string) *UpdateConfigDeliveryChannelRequest {
	s.DeliveryChannelId = &v
	return s
}

func (s *UpdateConfigDeliveryChannelRequest) SetDeliveryChannelName(v string) *UpdateConfigDeliveryChannelRequest {
	s.DeliveryChannelName = &v
	return s
}

func (s *UpdateConfigDeliveryChannelRequest) SetDeliveryChannelTargetArn(v string) *UpdateConfigDeliveryChannelRequest {
	s.DeliveryChannelTargetArn = &v
	return s
}

func (s *UpdateConfigDeliveryChannelRequest) SetDescription(v string) *UpdateConfigDeliveryChannelRequest {
	s.Description = &v
	return s
}

func (s *UpdateConfigDeliveryChannelRequest) SetNonCompliantNotification(v bool) *UpdateConfigDeliveryChannelRequest {
	s.NonCompliantNotification = &v
	return s
}

func (s *UpdateConfigDeliveryChannelRequest) SetOversizedDataOSSTargetArn(v string) *UpdateConfigDeliveryChannelRequest {
	s.OversizedDataOSSTargetArn = &v
	return s
}

func (s *UpdateConfigDeliveryChannelRequest) SetStatus(v int64) *UpdateConfigDeliveryChannelRequest {
	s.Status = &v
	return s
}

type UpdateConfigDeliveryChannelResponseBody struct {
	DeliveryChannelId *string `json:"DeliveryChannelId,omitempty" xml:"DeliveryChannelId,omitempty"`
	RequestId         *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateConfigDeliveryChannelResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateConfigDeliveryChannelResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateConfigDeliveryChannelResponseBody) SetDeliveryChannelId(v string) *UpdateConfigDeliveryChannelResponseBody {
	s.DeliveryChannelId = &v
	return s
}

func (s *UpdateConfigDeliveryChannelResponseBody) SetRequestId(v string) *UpdateConfigDeliveryChannelResponseBody {
	s.RequestId = &v
	return s
}

type UpdateConfigDeliveryChannelResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateConfigDeliveryChannelResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateConfigDeliveryChannelResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateConfigDeliveryChannelResponse) GoString() string {
	return s.String()
}

func (s *UpdateConfigDeliveryChannelResponse) SetHeaders(v map[string]*string) *UpdateConfigDeliveryChannelResponse {
	s.Headers = v
	return s
}

func (s *UpdateConfigDeliveryChannelResponse) SetStatusCode(v int32) *UpdateConfigDeliveryChannelResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateConfigDeliveryChannelResponse) SetBody(v *UpdateConfigDeliveryChannelResponseBody) *UpdateConfigDeliveryChannelResponse {
	s.Body = v
	return s
}

type UpdateConfigRuleRequest struct {
	ClientToken               *string                `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	ConfigRuleId              *string                `json:"ConfigRuleId,omitempty" xml:"ConfigRuleId,omitempty"`
	ConfigRuleName            *string                `json:"ConfigRuleName,omitempty" xml:"ConfigRuleName,omitempty"`
	ConfigRuleTriggerTypes    *string                `json:"ConfigRuleTriggerTypes,omitempty" xml:"ConfigRuleTriggerTypes,omitempty"`
	Description               *string                `json:"Description,omitempty" xml:"Description,omitempty"`
	ExcludeResourceIdsScope   *string                `json:"ExcludeResourceIdsScope,omitempty" xml:"ExcludeResourceIdsScope,omitempty"`
	InputParameters           map[string]interface{} `json:"InputParameters,omitempty" xml:"InputParameters,omitempty"`
	MaximumExecutionFrequency *string                `json:"MaximumExecutionFrequency,omitempty" xml:"MaximumExecutionFrequency,omitempty"`
	RegionIdsScope            *string                `json:"RegionIdsScope,omitempty" xml:"RegionIdsScope,omitempty"`
	ResourceGroupIdsScope     *string                `json:"ResourceGroupIdsScope,omitempty" xml:"ResourceGroupIdsScope,omitempty"`
	ResourceTypesScope        []*string              `json:"ResourceTypesScope,omitempty" xml:"ResourceTypesScope,omitempty" type:"Repeated"`
	RiskLevel                 *int32                 `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
	TagKeyLogicScope          *string                `json:"TagKeyLogicScope,omitempty" xml:"TagKeyLogicScope,omitempty"`
	TagKeyScope               *string                `json:"TagKeyScope,omitempty" xml:"TagKeyScope,omitempty"`
	TagValueScope             *string                `json:"TagValueScope,omitempty" xml:"TagValueScope,omitempty"`
}

func (s UpdateConfigRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateConfigRuleRequest) GoString() string {
	return s.String()
}

func (s *UpdateConfigRuleRequest) SetClientToken(v string) *UpdateConfigRuleRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateConfigRuleRequest) SetConfigRuleId(v string) *UpdateConfigRuleRequest {
	s.ConfigRuleId = &v
	return s
}

func (s *UpdateConfigRuleRequest) SetConfigRuleName(v string) *UpdateConfigRuleRequest {
	s.ConfigRuleName = &v
	return s
}

func (s *UpdateConfigRuleRequest) SetConfigRuleTriggerTypes(v string) *UpdateConfigRuleRequest {
	s.ConfigRuleTriggerTypes = &v
	return s
}

func (s *UpdateConfigRuleRequest) SetDescription(v string) *UpdateConfigRuleRequest {
	s.Description = &v
	return s
}

func (s *UpdateConfigRuleRequest) SetExcludeResourceIdsScope(v string) *UpdateConfigRuleRequest {
	s.ExcludeResourceIdsScope = &v
	return s
}

func (s *UpdateConfigRuleRequest) SetInputParameters(v map[string]interface{}) *UpdateConfigRuleRequest {
	s.InputParameters = v
	return s
}

func (s *UpdateConfigRuleRequest) SetMaximumExecutionFrequency(v string) *UpdateConfigRuleRequest {
	s.MaximumExecutionFrequency = &v
	return s
}

func (s *UpdateConfigRuleRequest) SetRegionIdsScope(v string) *UpdateConfigRuleRequest {
	s.RegionIdsScope = &v
	return s
}

func (s *UpdateConfigRuleRequest) SetResourceGroupIdsScope(v string) *UpdateConfigRuleRequest {
	s.ResourceGroupIdsScope = &v
	return s
}

func (s *UpdateConfigRuleRequest) SetResourceTypesScope(v []*string) *UpdateConfigRuleRequest {
	s.ResourceTypesScope = v
	return s
}

func (s *UpdateConfigRuleRequest) SetRiskLevel(v int32) *UpdateConfigRuleRequest {
	s.RiskLevel = &v
	return s
}

func (s *UpdateConfigRuleRequest) SetTagKeyLogicScope(v string) *UpdateConfigRuleRequest {
	s.TagKeyLogicScope = &v
	return s
}

func (s *UpdateConfigRuleRequest) SetTagKeyScope(v string) *UpdateConfigRuleRequest {
	s.TagKeyScope = &v
	return s
}

func (s *UpdateConfigRuleRequest) SetTagValueScope(v string) *UpdateConfigRuleRequest {
	s.TagValueScope = &v
	return s
}

type UpdateConfigRuleShrinkRequest struct {
	ClientToken               *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	ConfigRuleId              *string `json:"ConfigRuleId,omitempty" xml:"ConfigRuleId,omitempty"`
	ConfigRuleName            *string `json:"ConfigRuleName,omitempty" xml:"ConfigRuleName,omitempty"`
	ConfigRuleTriggerTypes    *string `json:"ConfigRuleTriggerTypes,omitempty" xml:"ConfigRuleTriggerTypes,omitempty"`
	Description               *string `json:"Description,omitempty" xml:"Description,omitempty"`
	ExcludeResourceIdsScope   *string `json:"ExcludeResourceIdsScope,omitempty" xml:"ExcludeResourceIdsScope,omitempty"`
	InputParametersShrink     *string `json:"InputParameters,omitempty" xml:"InputParameters,omitempty"`
	MaximumExecutionFrequency *string `json:"MaximumExecutionFrequency,omitempty" xml:"MaximumExecutionFrequency,omitempty"`
	RegionIdsScope            *string `json:"RegionIdsScope,omitempty" xml:"RegionIdsScope,omitempty"`
	ResourceGroupIdsScope     *string `json:"ResourceGroupIdsScope,omitempty" xml:"ResourceGroupIdsScope,omitempty"`
	ResourceTypesScopeShrink  *string `json:"ResourceTypesScope,omitempty" xml:"ResourceTypesScope,omitempty"`
	RiskLevel                 *int32  `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
	TagKeyLogicScope          *string `json:"TagKeyLogicScope,omitempty" xml:"TagKeyLogicScope,omitempty"`
	TagKeyScope               *string `json:"TagKeyScope,omitempty" xml:"TagKeyScope,omitempty"`
	TagValueScope             *string `json:"TagValueScope,omitempty" xml:"TagValueScope,omitempty"`
}

func (s UpdateConfigRuleShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateConfigRuleShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateConfigRuleShrinkRequest) SetClientToken(v string) *UpdateConfigRuleShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateConfigRuleShrinkRequest) SetConfigRuleId(v string) *UpdateConfigRuleShrinkRequest {
	s.ConfigRuleId = &v
	return s
}

func (s *UpdateConfigRuleShrinkRequest) SetConfigRuleName(v string) *UpdateConfigRuleShrinkRequest {
	s.ConfigRuleName = &v
	return s
}

func (s *UpdateConfigRuleShrinkRequest) SetConfigRuleTriggerTypes(v string) *UpdateConfigRuleShrinkRequest {
	s.ConfigRuleTriggerTypes = &v
	return s
}

func (s *UpdateConfigRuleShrinkRequest) SetDescription(v string) *UpdateConfigRuleShrinkRequest {
	s.Description = &v
	return s
}

func (s *UpdateConfigRuleShrinkRequest) SetExcludeResourceIdsScope(v string) *UpdateConfigRuleShrinkRequest {
	s.ExcludeResourceIdsScope = &v
	return s
}

func (s *UpdateConfigRuleShrinkRequest) SetInputParametersShrink(v string) *UpdateConfigRuleShrinkRequest {
	s.InputParametersShrink = &v
	return s
}

func (s *UpdateConfigRuleShrinkRequest) SetMaximumExecutionFrequency(v string) *UpdateConfigRuleShrinkRequest {
	s.MaximumExecutionFrequency = &v
	return s
}

func (s *UpdateConfigRuleShrinkRequest) SetRegionIdsScope(v string) *UpdateConfigRuleShrinkRequest {
	s.RegionIdsScope = &v
	return s
}

func (s *UpdateConfigRuleShrinkRequest) SetResourceGroupIdsScope(v string) *UpdateConfigRuleShrinkRequest {
	s.ResourceGroupIdsScope = &v
	return s
}

func (s *UpdateConfigRuleShrinkRequest) SetResourceTypesScopeShrink(v string) *UpdateConfigRuleShrinkRequest {
	s.ResourceTypesScopeShrink = &v
	return s
}

func (s *UpdateConfigRuleShrinkRequest) SetRiskLevel(v int32) *UpdateConfigRuleShrinkRequest {
	s.RiskLevel = &v
	return s
}

func (s *UpdateConfigRuleShrinkRequest) SetTagKeyLogicScope(v string) *UpdateConfigRuleShrinkRequest {
	s.TagKeyLogicScope = &v
	return s
}

func (s *UpdateConfigRuleShrinkRequest) SetTagKeyScope(v string) *UpdateConfigRuleShrinkRequest {
	s.TagKeyScope = &v
	return s
}

func (s *UpdateConfigRuleShrinkRequest) SetTagValueScope(v string) *UpdateConfigRuleShrinkRequest {
	s.TagValueScope = &v
	return s
}

type UpdateConfigRuleResponseBody struct {
	ConfigRuleId *string `json:"ConfigRuleId,omitempty" xml:"ConfigRuleId,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateConfigRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateConfigRuleResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateConfigRuleResponseBody) SetConfigRuleId(v string) *UpdateConfigRuleResponseBody {
	s.ConfigRuleId = &v
	return s
}

func (s *UpdateConfigRuleResponseBody) SetRequestId(v string) *UpdateConfigRuleResponseBody {
	s.RequestId = &v
	return s
}

type UpdateConfigRuleResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateConfigRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateConfigRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateConfigRuleResponse) GoString() string {
	return s.String()
}

func (s *UpdateConfigRuleResponse) SetHeaders(v map[string]*string) *UpdateConfigRuleResponse {
	s.Headers = v
	return s
}

func (s *UpdateConfigRuleResponse) SetStatusCode(v int32) *UpdateConfigRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateConfigRuleResponse) SetBody(v *UpdateConfigRuleResponseBody) *UpdateConfigRuleResponse {
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

func (client *Client) ActiveAggregateConfigRulesWithOptions(request *ActiveAggregateConfigRulesRequest, runtime *util.RuntimeOptions) (_result *ActiveAggregateConfigRulesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AggregatorId)) {
		query["AggregatorId"] = request.AggregatorId
	}

	if !tea.BoolValue(util.IsUnset(request.ConfigRuleIds)) {
		query["ConfigRuleIds"] = request.ConfigRuleIds
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ActiveAggregateConfigRules"),
		Version:     tea.String("2020-09-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ActiveAggregateConfigRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ActiveAggregateConfigRules(request *ActiveAggregateConfigRulesRequest) (_result *ActiveAggregateConfigRulesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ActiveAggregateConfigRulesResponse{}
	_body, _err := client.ActiveAggregateConfigRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AttachAggregateConfigRuleToCompliancePackWithOptions(request *AttachAggregateConfigRuleToCompliancePackRequest, runtime *util.RuntimeOptions) (_result *AttachAggregateConfigRuleToCompliancePackResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AggregatorId)) {
		query["AggregatorId"] = request.AggregatorId
	}

	if !tea.BoolValue(util.IsUnset(request.CompliancePackId)) {
		query["CompliancePackId"] = request.CompliancePackId
	}

	if !tea.BoolValue(util.IsUnset(request.ConfigRuleIds)) {
		query["ConfigRuleIds"] = request.ConfigRuleIds
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AttachAggregateConfigRuleToCompliancePack"),
		Version:     tea.String("2020-09-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AttachAggregateConfigRuleToCompliancePackResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AttachAggregateConfigRuleToCompliancePack(request *AttachAggregateConfigRuleToCompliancePackRequest) (_result *AttachAggregateConfigRuleToCompliancePackResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AttachAggregateConfigRuleToCompliancePackResponse{}
	_body, _err := client.AttachAggregateConfigRuleToCompliancePackWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AttachConfigRuleToCompliancePackWithOptions(request *AttachConfigRuleToCompliancePackRequest, runtime *util.RuntimeOptions) (_result *AttachConfigRuleToCompliancePackResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CompliancePackId)) {
		query["CompliancePackId"] = request.CompliancePackId
	}

	if !tea.BoolValue(util.IsUnset(request.ConfigRuleIds)) {
		query["ConfigRuleIds"] = request.ConfigRuleIds
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AttachConfigRuleToCompliancePack"),
		Version:     tea.String("2020-09-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AttachConfigRuleToCompliancePackResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AttachConfigRuleToCompliancePack(request *AttachConfigRuleToCompliancePackRequest) (_result *AttachConfigRuleToCompliancePackResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AttachConfigRuleToCompliancePackResponse{}
	_body, _err := client.AttachConfigRuleToCompliancePackWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateAggregateCompliancePackWithOptions(tmpReq *CreateAggregateCompliancePackRequest, runtime *util.RuntimeOptions) (_result *CreateAggregateCompliancePackResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateAggregateCompliancePackShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.ConfigRules)) {
		request.ConfigRulesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ConfigRules, tea.String("ConfigRules"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AggregatorId)) {
		body["AggregatorId"] = request.AggregatorId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.CompliancePackName)) {
		body["CompliancePackName"] = request.CompliancePackName
	}

	if !tea.BoolValue(util.IsUnset(request.CompliancePackTemplateId)) {
		body["CompliancePackTemplateId"] = request.CompliancePackTemplateId
	}

	if !tea.BoolValue(util.IsUnset(request.ConfigRulesShrink)) {
		body["ConfigRules"] = request.ConfigRulesShrink
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.RiskLevel)) {
		body["RiskLevel"] = request.RiskLevel
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateAggregateCompliancePack"),
		Version:     tea.String("2020-09-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateAggregateCompliancePackResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateAggregateCompliancePack(request *CreateAggregateCompliancePackRequest) (_result *CreateAggregateCompliancePackResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateAggregateCompliancePackResponse{}
	_body, _err := client.CreateAggregateCompliancePackWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateAggregateConfigDeliveryChannelWithOptions(request *CreateAggregateConfigDeliveryChannelRequest, runtime *util.RuntimeOptions) (_result *CreateAggregateConfigDeliveryChannelResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AggregatorId)) {
		query["AggregatorId"] = request.AggregatorId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.ConfigurationItemChangeNotification)) {
		query["ConfigurationItemChangeNotification"] = request.ConfigurationItemChangeNotification
	}

	if !tea.BoolValue(util.IsUnset(request.ConfigurationSnapshot)) {
		query["ConfigurationSnapshot"] = request.ConfigurationSnapshot
	}

	if !tea.BoolValue(util.IsUnset(request.DeliveryChannelCondition)) {
		query["DeliveryChannelCondition"] = request.DeliveryChannelCondition
	}

	if !tea.BoolValue(util.IsUnset(request.DeliveryChannelName)) {
		query["DeliveryChannelName"] = request.DeliveryChannelName
	}

	if !tea.BoolValue(util.IsUnset(request.DeliveryChannelTargetArn)) {
		query["DeliveryChannelTargetArn"] = request.DeliveryChannelTargetArn
	}

	if !tea.BoolValue(util.IsUnset(request.DeliveryChannelType)) {
		query["DeliveryChannelType"] = request.DeliveryChannelType
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.NonCompliantNotification)) {
		query["NonCompliantNotification"] = request.NonCompliantNotification
	}

	if !tea.BoolValue(util.IsUnset(request.OversizedDataOSSTargetArn)) {
		query["OversizedDataOSSTargetArn"] = request.OversizedDataOSSTargetArn
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateAggregateConfigDeliveryChannel"),
		Version:     tea.String("2020-09-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateAggregateConfigDeliveryChannelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateAggregateConfigDeliveryChannel(request *CreateAggregateConfigDeliveryChannelRequest) (_result *CreateAggregateConfigDeliveryChannelResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateAggregateConfigDeliveryChannelResponse{}
	_body, _err := client.CreateAggregateConfigDeliveryChannelWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateAggregateConfigRuleWithOptions(tmpReq *CreateAggregateConfigRuleRequest, runtime *util.RuntimeOptions) (_result *CreateAggregateConfigRuleResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateAggregateConfigRuleShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.InputParameters)) {
		request.InputParametersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.InputParameters, tea.String("InputParameters"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.ResourceTypesScope)) {
		request.ResourceTypesScopeShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ResourceTypesScope, tea.String("ResourceTypesScope"), tea.String("simple"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AggregatorId)) {
		body["AggregatorId"] = request.AggregatorId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.ConfigRuleName)) {
		body["ConfigRuleName"] = request.ConfigRuleName
	}

	if !tea.BoolValue(util.IsUnset(request.ConfigRuleTriggerTypes)) {
		body["ConfigRuleTriggerTypes"] = request.ConfigRuleTriggerTypes
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.ExcludeAccountIdsScope)) {
		body["ExcludeAccountIdsScope"] = request.ExcludeAccountIdsScope
	}

	if !tea.BoolValue(util.IsUnset(request.ExcludeFolderIdsScope)) {
		body["ExcludeFolderIdsScope"] = request.ExcludeFolderIdsScope
	}

	if !tea.BoolValue(util.IsUnset(request.ExcludeResourceIdsScope)) {
		body["ExcludeResourceIdsScope"] = request.ExcludeResourceIdsScope
	}

	if !tea.BoolValue(util.IsUnset(request.FolderIdsScope)) {
		body["FolderIdsScope"] = request.FolderIdsScope
	}

	if !tea.BoolValue(util.IsUnset(request.InputParametersShrink)) {
		body["InputParameters"] = request.InputParametersShrink
	}

	if !tea.BoolValue(util.IsUnset(request.MaximumExecutionFrequency)) {
		body["MaximumExecutionFrequency"] = request.MaximumExecutionFrequency
	}

	if !tea.BoolValue(util.IsUnset(request.RegionIdsScope)) {
		body["RegionIdsScope"] = request.RegionIdsScope
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupIdsScope)) {
		body["ResourceGroupIdsScope"] = request.ResourceGroupIdsScope
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceTypesScopeShrink)) {
		body["ResourceTypesScope"] = request.ResourceTypesScopeShrink
	}

	if !tea.BoolValue(util.IsUnset(request.RiskLevel)) {
		body["RiskLevel"] = request.RiskLevel
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIdentifier)) {
		body["SourceIdentifier"] = request.SourceIdentifier
	}

	if !tea.BoolValue(util.IsUnset(request.SourceOwner)) {
		body["SourceOwner"] = request.SourceOwner
	}

	if !tea.BoolValue(util.IsUnset(request.TagKeyLogicScope)) {
		body["TagKeyLogicScope"] = request.TagKeyLogicScope
	}

	if !tea.BoolValue(util.IsUnset(request.TagKeyScope)) {
		body["TagKeyScope"] = request.TagKeyScope
	}

	if !tea.BoolValue(util.IsUnset(request.TagValueScope)) {
		body["TagValueScope"] = request.TagValueScope
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateAggregateConfigRule"),
		Version:     tea.String("2020-09-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateAggregateConfigRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateAggregateConfigRule(request *CreateAggregateConfigRuleRequest) (_result *CreateAggregateConfigRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateAggregateConfigRuleResponse{}
	_body, _err := client.CreateAggregateConfigRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateAggregateRemediationWithOptions(request *CreateAggregateRemediationRequest, runtime *util.RuntimeOptions) (_result *CreateAggregateRemediationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AggregatorId)) {
		body["AggregatorId"] = request.AggregatorId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.ConfigRuleId)) {
		body["ConfigRuleId"] = request.ConfigRuleId
	}

	if !tea.BoolValue(util.IsUnset(request.InvokeType)) {
		body["InvokeType"] = request.InvokeType
	}

	if !tea.BoolValue(util.IsUnset(request.Params)) {
		body["Params"] = request.Params
	}

	if !tea.BoolValue(util.IsUnset(request.RemediationTemplateId)) {
		body["RemediationTemplateId"] = request.RemediationTemplateId
	}

	if !tea.BoolValue(util.IsUnset(request.RemediationType)) {
		body["RemediationType"] = request.RemediationType
	}

	if !tea.BoolValue(util.IsUnset(request.SourceType)) {
		body["SourceType"] = request.SourceType
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateAggregateRemediation"),
		Version:     tea.String("2020-09-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateAggregateRemediationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateAggregateRemediation(request *CreateAggregateRemediationRequest) (_result *CreateAggregateRemediationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateAggregateRemediationResponse{}
	_body, _err := client.CreateAggregateRemediationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateAggregatorWithOptions(tmpReq *CreateAggregatorRequest, runtime *util.RuntimeOptions) (_result *CreateAggregatorResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateAggregatorShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.AggregatorAccounts)) {
		request.AggregatorAccountsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AggregatorAccounts, tea.String("AggregatorAccounts"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AggregatorAccountsShrink)) {
		body["AggregatorAccounts"] = request.AggregatorAccountsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.AggregatorName)) {
		body["AggregatorName"] = request.AggregatorName
	}

	if !tea.BoolValue(util.IsUnset(request.AggregatorType)) {
		body["AggregatorType"] = request.AggregatorType
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateAggregator"),
		Version:     tea.String("2020-09-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateAggregatorResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateAggregator(request *CreateAggregatorRequest) (_result *CreateAggregatorResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateAggregatorResponse{}
	_body, _err := client.CreateAggregatorWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateCompliancePackWithOptions(tmpReq *CreateCompliancePackRequest, runtime *util.RuntimeOptions) (_result *CreateCompliancePackResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateCompliancePackShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.ConfigRules)) {
		request.ConfigRulesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ConfigRules, tea.String("ConfigRules"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.CompliancePackName)) {
		body["CompliancePackName"] = request.CompliancePackName
	}

	if !tea.BoolValue(util.IsUnset(request.CompliancePackTemplateId)) {
		body["CompliancePackTemplateId"] = request.CompliancePackTemplateId
	}

	if !tea.BoolValue(util.IsUnset(request.ConfigRulesShrink)) {
		body["ConfigRules"] = request.ConfigRulesShrink
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.RiskLevel)) {
		body["RiskLevel"] = request.RiskLevel
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateCompliancePack"),
		Version:     tea.String("2020-09-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateCompliancePackResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateCompliancePack(request *CreateCompliancePackRequest) (_result *CreateCompliancePackResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateCompliancePackResponse{}
	_body, _err := client.CreateCompliancePackWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateConfigDeliveryChannelWithOptions(request *CreateConfigDeliveryChannelRequest, runtime *util.RuntimeOptions) (_result *CreateConfigDeliveryChannelResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.ConfigurationItemChangeNotification)) {
		query["ConfigurationItemChangeNotification"] = request.ConfigurationItemChangeNotification
	}

	if !tea.BoolValue(util.IsUnset(request.ConfigurationSnapshot)) {
		query["ConfigurationSnapshot"] = request.ConfigurationSnapshot
	}

	if !tea.BoolValue(util.IsUnset(request.DeliveryChannelCondition)) {
		query["DeliveryChannelCondition"] = request.DeliveryChannelCondition
	}

	if !tea.BoolValue(util.IsUnset(request.DeliveryChannelName)) {
		query["DeliveryChannelName"] = request.DeliveryChannelName
	}

	if !tea.BoolValue(util.IsUnset(request.DeliveryChannelTargetArn)) {
		query["DeliveryChannelTargetArn"] = request.DeliveryChannelTargetArn
	}

	if !tea.BoolValue(util.IsUnset(request.DeliveryChannelType)) {
		query["DeliveryChannelType"] = request.DeliveryChannelType
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.NonCompliantNotification)) {
		query["NonCompliantNotification"] = request.NonCompliantNotification
	}

	if !tea.BoolValue(util.IsUnset(request.OversizedDataOSSTargetArn)) {
		query["OversizedDataOSSTargetArn"] = request.OversizedDataOSSTargetArn
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateConfigDeliveryChannel"),
		Version:     tea.String("2020-09-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateConfigDeliveryChannelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateConfigDeliveryChannel(request *CreateConfigDeliveryChannelRequest) (_result *CreateConfigDeliveryChannelResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateConfigDeliveryChannelResponse{}
	_body, _err := client.CreateConfigDeliveryChannelWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateConfigRuleWithOptions(tmpReq *CreateConfigRuleRequest, runtime *util.RuntimeOptions) (_result *CreateConfigRuleResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateConfigRuleShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.InputParameters)) {
		request.InputParametersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.InputParameters, tea.String("InputParameters"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.ResourceTypesScope)) {
		request.ResourceTypesScopeShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ResourceTypesScope, tea.String("ResourceTypesScope"), tea.String("simple"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.ConfigRuleName)) {
		body["ConfigRuleName"] = request.ConfigRuleName
	}

	if !tea.BoolValue(util.IsUnset(request.ConfigRuleTriggerTypes)) {
		body["ConfigRuleTriggerTypes"] = request.ConfigRuleTriggerTypes
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.ExcludeResourceIdsScope)) {
		body["ExcludeResourceIdsScope"] = request.ExcludeResourceIdsScope
	}

	if !tea.BoolValue(util.IsUnset(request.InputParametersShrink)) {
		body["InputParameters"] = request.InputParametersShrink
	}

	if !tea.BoolValue(util.IsUnset(request.MaximumExecutionFrequency)) {
		body["MaximumExecutionFrequency"] = request.MaximumExecutionFrequency
	}

	if !tea.BoolValue(util.IsUnset(request.RegionIdsScope)) {
		body["RegionIdsScope"] = request.RegionIdsScope
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupIdsScope)) {
		body["ResourceGroupIdsScope"] = request.ResourceGroupIdsScope
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceTypesScopeShrink)) {
		body["ResourceTypesScope"] = request.ResourceTypesScopeShrink
	}

	if !tea.BoolValue(util.IsUnset(request.RiskLevel)) {
		body["RiskLevel"] = request.RiskLevel
	}

	if !tea.BoolValue(util.IsUnset(request.SourceIdentifier)) {
		body["SourceIdentifier"] = request.SourceIdentifier
	}

	if !tea.BoolValue(util.IsUnset(request.SourceOwner)) {
		body["SourceOwner"] = request.SourceOwner
	}

	if !tea.BoolValue(util.IsUnset(request.TagKeyLogicScope)) {
		body["TagKeyLogicScope"] = request.TagKeyLogicScope
	}

	if !tea.BoolValue(util.IsUnset(request.TagKeyScope)) {
		body["TagKeyScope"] = request.TagKeyScope
	}

	if !tea.BoolValue(util.IsUnset(request.TagValueScope)) {
		body["TagValueScope"] = request.TagValueScope
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateConfigRule"),
		Version:     tea.String("2020-09-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateConfigRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateConfigRule(request *CreateConfigRuleRequest) (_result *CreateConfigRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateConfigRuleResponse{}
	_body, _err := client.CreateConfigRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateRemediationWithOptions(request *CreateRemediationRequest, runtime *util.RuntimeOptions) (_result *CreateRemediationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.ConfigRuleId)) {
		body["ConfigRuleId"] = request.ConfigRuleId
	}

	if !tea.BoolValue(util.IsUnset(request.InvokeType)) {
		body["InvokeType"] = request.InvokeType
	}

	if !tea.BoolValue(util.IsUnset(request.Params)) {
		body["Params"] = request.Params
	}

	if !tea.BoolValue(util.IsUnset(request.RemediationTemplateId)) {
		body["RemediationTemplateId"] = request.RemediationTemplateId
	}

	if !tea.BoolValue(util.IsUnset(request.RemediationType)) {
		body["RemediationType"] = request.RemediationType
	}

	if !tea.BoolValue(util.IsUnset(request.SourceType)) {
		body["SourceType"] = request.SourceType
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateRemediation"),
		Version:     tea.String("2020-09-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateRemediationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateRemediation(request *CreateRemediationRequest) (_result *CreateRemediationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateRemediationResponse{}
	_body, _err := client.CreateRemediationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeactiveAggregateConfigRulesWithOptions(request *DeactiveAggregateConfigRulesRequest, runtime *util.RuntimeOptions) (_result *DeactiveAggregateConfigRulesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AggregatorId)) {
		query["AggregatorId"] = request.AggregatorId
	}

	if !tea.BoolValue(util.IsUnset(request.ConfigRuleIds)) {
		query["ConfigRuleIds"] = request.ConfigRuleIds
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeactiveAggregateConfigRules"),
		Version:     tea.String("2020-09-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeactiveAggregateConfigRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeactiveAggregateConfigRules(request *DeactiveAggregateConfigRulesRequest) (_result *DeactiveAggregateConfigRulesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeactiveAggregateConfigRulesResponse{}
	_body, _err := client.DeactiveAggregateConfigRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeactiveConfigRulesWithOptions(request *DeactiveConfigRulesRequest, runtime *util.RuntimeOptions) (_result *DeactiveConfigRulesResponse, _err error) {
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
		Action:      tea.String("DeactiveConfigRules"),
		Version:     tea.String("2020-09-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeactiveConfigRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeactiveConfigRules(request *DeactiveConfigRulesRequest) (_result *DeactiveConfigRulesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeactiveConfigRulesResponse{}
	_body, _err := client.DeactiveConfigRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteAggregateCompliancePacksWithOptions(request *DeleteAggregateCompliancePacksRequest, runtime *util.RuntimeOptions) (_result *DeleteAggregateCompliancePacksResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AggregatorId)) {
		body["AggregatorId"] = request.AggregatorId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.CompliancePackIds)) {
		body["CompliancePackIds"] = request.CompliancePackIds
	}

	if !tea.BoolValue(util.IsUnset(request.DeleteRule)) {
		body["DeleteRule"] = request.DeleteRule
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteAggregateCompliancePacks"),
		Version:     tea.String("2020-09-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteAggregateCompliancePacksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteAggregateCompliancePacks(request *DeleteAggregateCompliancePacksRequest) (_result *DeleteAggregateCompliancePacksResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteAggregateCompliancePacksResponse{}
	_body, _err := client.DeleteAggregateCompliancePacksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteAggregateConfigRulesWithOptions(request *DeleteAggregateConfigRulesRequest, runtime *util.RuntimeOptions) (_result *DeleteAggregateConfigRulesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AggregatorId)) {
		query["AggregatorId"] = request.AggregatorId
	}

	if !tea.BoolValue(util.IsUnset(request.ConfigRuleIds)) {
		query["ConfigRuleIds"] = request.ConfigRuleIds
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteAggregateConfigRules"),
		Version:     tea.String("2020-09-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteAggregateConfigRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteAggregateConfigRules(request *DeleteAggregateConfigRulesRequest) (_result *DeleteAggregateConfigRulesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteAggregateConfigRulesResponse{}
	_body, _err := client.DeleteAggregateConfigRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteAggregateRemediationsWithOptions(request *DeleteAggregateRemediationsRequest, runtime *util.RuntimeOptions) (_result *DeleteAggregateRemediationsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AggregatorId)) {
		body["AggregatorId"] = request.AggregatorId
	}

	if !tea.BoolValue(util.IsUnset(request.RemediationIds)) {
		body["RemediationIds"] = request.RemediationIds
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteAggregateRemediations"),
		Version:     tea.String("2020-09-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteAggregateRemediationsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteAggregateRemediations(request *DeleteAggregateRemediationsRequest) (_result *DeleteAggregateRemediationsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteAggregateRemediationsResponse{}
	_body, _err := client.DeleteAggregateRemediationsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteAggregatorsWithOptions(request *DeleteAggregatorsRequest, runtime *util.RuntimeOptions) (_result *DeleteAggregatorsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AggregatorIds)) {
		body["AggregatorIds"] = request.AggregatorIds
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteAggregators"),
		Version:     tea.String("2020-09-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteAggregatorsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteAggregators(request *DeleteAggregatorsRequest) (_result *DeleteAggregatorsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteAggregatorsResponse{}
	_body, _err := client.DeleteAggregatorsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteCompliancePacksWithOptions(request *DeleteCompliancePacksRequest, runtime *util.RuntimeOptions) (_result *DeleteCompliancePacksResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.CompliancePackIds)) {
		body["CompliancePackIds"] = request.CompliancePackIds
	}

	if !tea.BoolValue(util.IsUnset(request.DeleteRule)) {
		body["DeleteRule"] = request.DeleteRule
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteCompliancePacks"),
		Version:     tea.String("2020-09-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteCompliancePacksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteCompliancePacks(request *DeleteCompliancePacksRequest) (_result *DeleteCompliancePacksResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteCompliancePacksResponse{}
	_body, _err := client.DeleteCompliancePacksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteRemediationsWithOptions(request *DeleteRemediationsRequest, runtime *util.RuntimeOptions) (_result *DeleteRemediationsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RemediationIds)) {
		body["RemediationIds"] = request.RemediationIds
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteRemediations"),
		Version:     tea.String("2020-09-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteRemediationsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteRemediations(request *DeleteRemediationsRequest) (_result *DeleteRemediationsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteRemediationsResponse{}
	_body, _err := client.DeleteRemediationsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DetachAggregateConfigRuleToCompliancePackWithOptions(request *DetachAggregateConfigRuleToCompliancePackRequest, runtime *util.RuntimeOptions) (_result *DetachAggregateConfigRuleToCompliancePackResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AggregatorId)) {
		query["AggregatorId"] = request.AggregatorId
	}

	if !tea.BoolValue(util.IsUnset(request.CompliancePackId)) {
		query["CompliancePackId"] = request.CompliancePackId
	}

	if !tea.BoolValue(util.IsUnset(request.ConfigRuleIds)) {
		query["ConfigRuleIds"] = request.ConfigRuleIds
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DetachAggregateConfigRuleToCompliancePack"),
		Version:     tea.String("2020-09-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DetachAggregateConfigRuleToCompliancePackResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DetachAggregateConfigRuleToCompliancePack(request *DetachAggregateConfigRuleToCompliancePackRequest) (_result *DetachAggregateConfigRuleToCompliancePackResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DetachAggregateConfigRuleToCompliancePackResponse{}
	_body, _err := client.DetachAggregateConfigRuleToCompliancePackWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DetachConfigRuleToCompliancePackWithOptions(request *DetachConfigRuleToCompliancePackRequest, runtime *util.RuntimeOptions) (_result *DetachConfigRuleToCompliancePackResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CompliancePackId)) {
		query["CompliancePackId"] = request.CompliancePackId
	}

	if !tea.BoolValue(util.IsUnset(request.ConfigRuleIds)) {
		query["ConfigRuleIds"] = request.ConfigRuleIds
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DetachConfigRuleToCompliancePack"),
		Version:     tea.String("2020-09-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DetachConfigRuleToCompliancePackResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DetachConfigRuleToCompliancePack(request *DetachConfigRuleToCompliancePackRequest) (_result *DetachConfigRuleToCompliancePackResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DetachConfigRuleToCompliancePackResponse{}
	_body, _err := client.DetachConfigRuleToCompliancePackWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GenerateAggregateCompliancePackReportWithOptions(request *GenerateAggregateCompliancePackReportRequest, runtime *util.RuntimeOptions) (_result *GenerateAggregateCompliancePackReportResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AggregatorId)) {
		body["AggregatorId"] = request.AggregatorId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.CompliancePackId)) {
		body["CompliancePackId"] = request.CompliancePackId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GenerateAggregateCompliancePackReport"),
		Version:     tea.String("2020-09-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GenerateAggregateCompliancePackReportResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GenerateAggregateCompliancePackReport(request *GenerateAggregateCompliancePackReportRequest) (_result *GenerateAggregateCompliancePackReportResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GenerateAggregateCompliancePackReportResponse{}
	_body, _err := client.GenerateAggregateCompliancePackReportWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GenerateAggregateConfigRulesReportWithOptions(request *GenerateAggregateConfigRulesReportRequest, runtime *util.RuntimeOptions) (_result *GenerateAggregateConfigRulesReportResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AggregatorId)) {
		body["AggregatorId"] = request.AggregatorId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.ConfigRuleIds)) {
		body["ConfigRuleIds"] = request.ConfigRuleIds
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GenerateAggregateConfigRulesReport"),
		Version:     tea.String("2020-09-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GenerateAggregateConfigRulesReportResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GenerateAggregateConfigRulesReport(request *GenerateAggregateConfigRulesReportRequest) (_result *GenerateAggregateConfigRulesReportResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GenerateAggregateConfigRulesReportResponse{}
	_body, _err := client.GenerateAggregateConfigRulesReportWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GenerateCompliancePackReportWithOptions(request *GenerateCompliancePackReportRequest, runtime *util.RuntimeOptions) (_result *GenerateCompliancePackReportResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.CompliancePackId)) {
		body["CompliancePackId"] = request.CompliancePackId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GenerateCompliancePackReport"),
		Version:     tea.String("2020-09-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GenerateCompliancePackReportResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GenerateCompliancePackReport(request *GenerateCompliancePackReportRequest) (_result *GenerateCompliancePackReportResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GenerateCompliancePackReportResponse{}
	_body, _err := client.GenerateCompliancePackReportWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GenerateConfigRulesReportWithOptions(request *GenerateConfigRulesReportRequest, runtime *util.RuntimeOptions) (_result *GenerateConfigRulesReportResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.ConfigRuleIds)) {
		body["ConfigRuleIds"] = request.ConfigRuleIds
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GenerateConfigRulesReport"),
		Version:     tea.String("2020-09-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GenerateConfigRulesReportResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GenerateConfigRulesReport(request *GenerateConfigRulesReportRequest) (_result *GenerateConfigRulesReportResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GenerateConfigRulesReportResponse{}
	_body, _err := client.GenerateConfigRulesReportWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetAggregateAccountComplianceByPackWithOptions(request *GetAggregateAccountComplianceByPackRequest, runtime *util.RuntimeOptions) (_result *GetAggregateAccountComplianceByPackResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetAggregateAccountComplianceByPack"),
		Version:     tea.String("2020-09-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAggregateAccountComplianceByPackResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetAggregateAccountComplianceByPack(request *GetAggregateAccountComplianceByPackRequest) (_result *GetAggregateAccountComplianceByPackResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAggregateAccountComplianceByPackResponse{}
	_body, _err := client.GetAggregateAccountComplianceByPackWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetAggregateCompliancePackWithOptions(request *GetAggregateCompliancePackRequest, runtime *util.RuntimeOptions) (_result *GetAggregateCompliancePackResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetAggregateCompliancePack"),
		Version:     tea.String("2020-09-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAggregateCompliancePackResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetAggregateCompliancePack(request *GetAggregateCompliancePackRequest) (_result *GetAggregateCompliancePackResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAggregateCompliancePackResponse{}
	_body, _err := client.GetAggregateCompliancePackWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetAggregateCompliancePackReportWithOptions(request *GetAggregateCompliancePackReportRequest, runtime *util.RuntimeOptions) (_result *GetAggregateCompliancePackReportResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetAggregateCompliancePackReport"),
		Version:     tea.String("2020-09-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAggregateCompliancePackReportResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetAggregateCompliancePackReport(request *GetAggregateCompliancePackReportRequest) (_result *GetAggregateCompliancePackReportResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAggregateCompliancePackReportResponse{}
	_body, _err := client.GetAggregateCompliancePackReportWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetAggregateConfigDeliveryChannelWithOptions(request *GetAggregateConfigDeliveryChannelRequest, runtime *util.RuntimeOptions) (_result *GetAggregateConfigDeliveryChannelResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AggregatorId)) {
		query["AggregatorId"] = request.AggregatorId
	}

	if !tea.BoolValue(util.IsUnset(request.DeliveryChannelId)) {
		query["DeliveryChannelId"] = request.DeliveryChannelId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetAggregateConfigDeliveryChannel"),
		Version:     tea.String("2020-09-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAggregateConfigDeliveryChannelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetAggregateConfigDeliveryChannel(request *GetAggregateConfigDeliveryChannelRequest) (_result *GetAggregateConfigDeliveryChannelResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAggregateConfigDeliveryChannelResponse{}
	_body, _err := client.GetAggregateConfigDeliveryChannelWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetAggregateConfigRuleWithOptions(request *GetAggregateConfigRuleRequest, runtime *util.RuntimeOptions) (_result *GetAggregateConfigRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AggregatorId)) {
		query["AggregatorId"] = request.AggregatorId
	}

	if !tea.BoolValue(util.IsUnset(request.ConfigRuleId)) {
		query["ConfigRuleId"] = request.ConfigRuleId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetAggregateConfigRule"),
		Version:     tea.String("2020-09-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAggregateConfigRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetAggregateConfigRule(request *GetAggregateConfigRuleRequest) (_result *GetAggregateConfigRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAggregateConfigRuleResponse{}
	_body, _err := client.GetAggregateConfigRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetAggregateConfigRuleComplianceByPackWithOptions(request *GetAggregateConfigRuleComplianceByPackRequest, runtime *util.RuntimeOptions) (_result *GetAggregateConfigRuleComplianceByPackResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetAggregateConfigRuleComplianceByPack"),
		Version:     tea.String("2020-09-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAggregateConfigRuleComplianceByPackResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetAggregateConfigRuleComplianceByPack(request *GetAggregateConfigRuleComplianceByPackRequest) (_result *GetAggregateConfigRuleComplianceByPackResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAggregateConfigRuleComplianceByPackResponse{}
	_body, _err := client.GetAggregateConfigRuleComplianceByPackWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetAggregateConfigRuleSummaryByRiskLevelWithOptions(request *GetAggregateConfigRuleSummaryByRiskLevelRequest, runtime *util.RuntimeOptions) (_result *GetAggregateConfigRuleSummaryByRiskLevelResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetAggregateConfigRuleSummaryByRiskLevel"),
		Version:     tea.String("2020-09-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAggregateConfigRuleSummaryByRiskLevelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetAggregateConfigRuleSummaryByRiskLevel(request *GetAggregateConfigRuleSummaryByRiskLevelRequest) (_result *GetAggregateConfigRuleSummaryByRiskLevelResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAggregateConfigRuleSummaryByRiskLevelResponse{}
	_body, _err := client.GetAggregateConfigRuleSummaryByRiskLevelWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetAggregateConfigRulesReportWithOptions(request *GetAggregateConfigRulesReportRequest, runtime *util.RuntimeOptions) (_result *GetAggregateConfigRulesReportResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AggregatorId)) {
		query["AggregatorId"] = request.AggregatorId
	}

	if !tea.BoolValue(util.IsUnset(request.ReportId)) {
		query["ReportId"] = request.ReportId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetAggregateConfigRulesReport"),
		Version:     tea.String("2020-09-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAggregateConfigRulesReportResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetAggregateConfigRulesReport(request *GetAggregateConfigRulesReportRequest) (_result *GetAggregateConfigRulesReportResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAggregateConfigRulesReportResponse{}
	_body, _err := client.GetAggregateConfigRulesReportWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

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
		Version:     tea.String("2020-09-07"),
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

func (client *Client) GetAggregateResourceComplianceByConfigRuleWithOptions(request *GetAggregateResourceComplianceByConfigRuleRequest, runtime *util.RuntimeOptions) (_result *GetAggregateResourceComplianceByConfigRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AggregatorId)) {
		query["AggregatorId"] = request.AggregatorId
	}

	if !tea.BoolValue(util.IsUnset(request.ComplianceType)) {
		query["ComplianceType"] = request.ComplianceType
	}

	if !tea.BoolValue(util.IsUnset(request.ConfigRuleId)) {
		query["ConfigRuleId"] = request.ConfigRuleId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetAggregateResourceComplianceByConfigRule"),
		Version:     tea.String("2020-09-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAggregateResourceComplianceByConfigRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetAggregateResourceComplianceByConfigRule(request *GetAggregateResourceComplianceByConfigRuleRequest) (_result *GetAggregateResourceComplianceByConfigRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAggregateResourceComplianceByConfigRuleResponse{}
	_body, _err := client.GetAggregateResourceComplianceByConfigRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetAggregateResourceComplianceByPackWithOptions(request *GetAggregateResourceComplianceByPackRequest, runtime *util.RuntimeOptions) (_result *GetAggregateResourceComplianceByPackResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetAggregateResourceComplianceByPack"),
		Version:     tea.String("2020-09-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAggregateResourceComplianceByPackResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetAggregateResourceComplianceByPack(request *GetAggregateResourceComplianceByPackRequest) (_result *GetAggregateResourceComplianceByPackResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAggregateResourceComplianceByPackResponse{}
	_body, _err := client.GetAggregateResourceComplianceByPackWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetAggregateResourceComplianceGroupByRegionWithOptions(request *GetAggregateResourceComplianceGroupByRegionRequest, runtime *util.RuntimeOptions) (_result *GetAggregateResourceComplianceGroupByRegionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AggregatorId)) {
		query["AggregatorId"] = request.AggregatorId
	}

	if !tea.BoolValue(util.IsUnset(request.ConfigRuleIds)) {
		query["ConfigRuleIds"] = request.ConfigRuleIds
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetAggregateResourceComplianceGroupByRegion"),
		Version:     tea.String("2020-09-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAggregateResourceComplianceGroupByRegionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetAggregateResourceComplianceGroupByRegion(request *GetAggregateResourceComplianceGroupByRegionRequest) (_result *GetAggregateResourceComplianceGroupByRegionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAggregateResourceComplianceGroupByRegionResponse{}
	_body, _err := client.GetAggregateResourceComplianceGroupByRegionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetAggregateResourceComplianceGroupByResourceTypeWithOptions(request *GetAggregateResourceComplianceGroupByResourceTypeRequest, runtime *util.RuntimeOptions) (_result *GetAggregateResourceComplianceGroupByResourceTypeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AggregatorId)) {
		query["AggregatorId"] = request.AggregatorId
	}

	if !tea.BoolValue(util.IsUnset(request.ConfigRuleIds)) {
		query["ConfigRuleIds"] = request.ConfigRuleIds
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetAggregateResourceComplianceGroupByResourceType"),
		Version:     tea.String("2020-09-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAggregateResourceComplianceGroupByResourceTypeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetAggregateResourceComplianceGroupByResourceType(request *GetAggregateResourceComplianceGroupByResourceTypeRequest) (_result *GetAggregateResourceComplianceGroupByResourceTypeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAggregateResourceComplianceGroupByResourceTypeResponse{}
	_body, _err := client.GetAggregateResourceComplianceGroupByResourceTypeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetAggregateResourceComplianceTimelineWithOptions(request *GetAggregateResourceComplianceTimelineRequest, runtime *util.RuntimeOptions) (_result *GetAggregateResourceComplianceTimelineResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetAggregateResourceComplianceTimeline"),
		Version:     tea.String("2020-09-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAggregateResourceComplianceTimelineResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetAggregateResourceComplianceTimeline(request *GetAggregateResourceComplianceTimelineRequest) (_result *GetAggregateResourceComplianceTimelineResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAggregateResourceComplianceTimelineResponse{}
	_body, _err := client.GetAggregateResourceComplianceTimelineWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetAggregateResourceConfigurationTimelineWithOptions(request *GetAggregateResourceConfigurationTimelineRequest, runtime *util.RuntimeOptions) (_result *GetAggregateResourceConfigurationTimelineResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetAggregateResourceConfigurationTimeline"),
		Version:     tea.String("2020-09-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAggregateResourceConfigurationTimelineResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetAggregateResourceConfigurationTimeline(request *GetAggregateResourceConfigurationTimelineRequest) (_result *GetAggregateResourceConfigurationTimelineResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAggregateResourceConfigurationTimelineResponse{}
	_body, _err := client.GetAggregateResourceConfigurationTimelineWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetAggregateResourceCountsGroupByRegionWithOptions(request *GetAggregateResourceCountsGroupByRegionRequest, runtime *util.RuntimeOptions) (_result *GetAggregateResourceCountsGroupByRegionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetAggregateResourceCountsGroupByRegion"),
		Version:     tea.String("2020-09-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAggregateResourceCountsGroupByRegionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetAggregateResourceCountsGroupByRegion(request *GetAggregateResourceCountsGroupByRegionRequest) (_result *GetAggregateResourceCountsGroupByRegionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAggregateResourceCountsGroupByRegionResponse{}
	_body, _err := client.GetAggregateResourceCountsGroupByRegionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetAggregateResourceCountsGroupByResourceTypeWithOptions(request *GetAggregateResourceCountsGroupByResourceTypeRequest, runtime *util.RuntimeOptions) (_result *GetAggregateResourceCountsGroupByResourceTypeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetAggregateResourceCountsGroupByResourceType"),
		Version:     tea.String("2020-09-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAggregateResourceCountsGroupByResourceTypeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetAggregateResourceCountsGroupByResourceType(request *GetAggregateResourceCountsGroupByResourceTypeRequest) (_result *GetAggregateResourceCountsGroupByResourceTypeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAggregateResourceCountsGroupByResourceTypeResponse{}
	_body, _err := client.GetAggregateResourceCountsGroupByResourceTypeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetAggregatorWithOptions(request *GetAggregatorRequest, runtime *util.RuntimeOptions) (_result *GetAggregatorResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetAggregator"),
		Version:     tea.String("2020-09-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAggregatorResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetAggregator(request *GetAggregatorRequest) (_result *GetAggregatorResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAggregatorResponse{}
	_body, _err := client.GetAggregatorWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetCompliancePackWithOptions(request *GetCompliancePackRequest, runtime *util.RuntimeOptions) (_result *GetCompliancePackResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetCompliancePack"),
		Version:     tea.String("2020-09-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetCompliancePackResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetCompliancePack(request *GetCompliancePackRequest) (_result *GetCompliancePackResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetCompliancePackResponse{}
	_body, _err := client.GetCompliancePackWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetCompliancePackReportWithOptions(request *GetCompliancePackReportRequest, runtime *util.RuntimeOptions) (_result *GetCompliancePackReportResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetCompliancePackReport"),
		Version:     tea.String("2020-09-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetCompliancePackReportResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetCompliancePackReport(request *GetCompliancePackReportRequest) (_result *GetCompliancePackReportResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetCompliancePackReportResponse{}
	_body, _err := client.GetCompliancePackReportWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetConfigDeliveryChannelWithOptions(request *GetConfigDeliveryChannelRequest, runtime *util.RuntimeOptions) (_result *GetConfigDeliveryChannelResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeliveryChannelId)) {
		query["DeliveryChannelId"] = request.DeliveryChannelId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetConfigDeliveryChannel"),
		Version:     tea.String("2020-09-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetConfigDeliveryChannelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetConfigDeliveryChannel(request *GetConfigDeliveryChannelRequest) (_result *GetConfigDeliveryChannelResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetConfigDeliveryChannelResponse{}
	_body, _err := client.GetConfigDeliveryChannelWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetConfigRuleWithOptions(request *GetConfigRuleRequest, runtime *util.RuntimeOptions) (_result *GetConfigRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConfigRuleId)) {
		query["ConfigRuleId"] = request.ConfigRuleId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetConfigRule"),
		Version:     tea.String("2020-09-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetConfigRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetConfigRule(request *GetConfigRuleRequest) (_result *GetConfigRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetConfigRuleResponse{}
	_body, _err := client.GetConfigRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetConfigRuleComplianceByPackWithOptions(request *GetConfigRuleComplianceByPackRequest, runtime *util.RuntimeOptions) (_result *GetConfigRuleComplianceByPackResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetConfigRuleComplianceByPack"),
		Version:     tea.String("2020-09-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetConfigRuleComplianceByPackResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetConfigRuleComplianceByPack(request *GetConfigRuleComplianceByPackRequest) (_result *GetConfigRuleComplianceByPackResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetConfigRuleComplianceByPackResponse{}
	_body, _err := client.GetConfigRuleComplianceByPackWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetConfigRuleSummaryByRiskLevelWithOptions(runtime *util.RuntimeOptions) (_result *GetConfigRuleSummaryByRiskLevelResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("GetConfigRuleSummaryByRiskLevel"),
		Version:     tea.String("2020-09-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetConfigRuleSummaryByRiskLevelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetConfigRuleSummaryByRiskLevel() (_result *GetConfigRuleSummaryByRiskLevelResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetConfigRuleSummaryByRiskLevelResponse{}
	_body, _err := client.GetConfigRuleSummaryByRiskLevelWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetConfigRulesReportWithOptions(request *GetConfigRulesReportRequest, runtime *util.RuntimeOptions) (_result *GetConfigRulesReportResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ReportId)) {
		query["ReportId"] = request.ReportId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetConfigRulesReport"),
		Version:     tea.String("2020-09-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetConfigRulesReportResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetConfigRulesReport(request *GetConfigRulesReportRequest) (_result *GetConfigRulesReportResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetConfigRulesReportResponse{}
	_body, _err := client.GetConfigRulesReportWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetDiscoveredResourceWithOptions(request *GetDiscoveredResourceRequest, runtime *util.RuntimeOptions) (_result *GetDiscoveredResourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetDiscoveredResource"),
		Version:     tea.String("2020-09-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetDiscoveredResourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetDiscoveredResource(request *GetDiscoveredResourceRequest) (_result *GetDiscoveredResourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetDiscoveredResourceResponse{}
	_body, _err := client.GetDiscoveredResourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetDiscoveredResourceCountsGroupByRegionWithOptions(request *GetDiscoveredResourceCountsGroupByRegionRequest, runtime *util.RuntimeOptions) (_result *GetDiscoveredResourceCountsGroupByRegionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetDiscoveredResourceCountsGroupByRegion"),
		Version:     tea.String("2020-09-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetDiscoveredResourceCountsGroupByRegionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetDiscoveredResourceCountsGroupByRegion(request *GetDiscoveredResourceCountsGroupByRegionRequest) (_result *GetDiscoveredResourceCountsGroupByRegionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetDiscoveredResourceCountsGroupByRegionResponse{}
	_body, _err := client.GetDiscoveredResourceCountsGroupByRegionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetDiscoveredResourceCountsGroupByResourceTypeWithOptions(request *GetDiscoveredResourceCountsGroupByResourceTypeRequest, runtime *util.RuntimeOptions) (_result *GetDiscoveredResourceCountsGroupByResourceTypeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetDiscoveredResourceCountsGroupByResourceType"),
		Version:     tea.String("2020-09-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetDiscoveredResourceCountsGroupByResourceTypeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetDiscoveredResourceCountsGroupByResourceType(request *GetDiscoveredResourceCountsGroupByResourceTypeRequest) (_result *GetDiscoveredResourceCountsGroupByResourceTypeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetDiscoveredResourceCountsGroupByResourceTypeResponse{}
	_body, _err := client.GetDiscoveredResourceCountsGroupByResourceTypeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetManagedRuleWithOptions(request *GetManagedRuleRequest, runtime *util.RuntimeOptions) (_result *GetManagedRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Identifier)) {
		query["Identifier"] = request.Identifier
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetManagedRule"),
		Version:     tea.String("2020-09-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetManagedRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetManagedRule(request *GetManagedRuleRequest) (_result *GetManagedRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetManagedRuleResponse{}
	_body, _err := client.GetManagedRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetResourceComplianceByConfigRuleWithOptions(request *GetResourceComplianceByConfigRuleRequest, runtime *util.RuntimeOptions) (_result *GetResourceComplianceByConfigRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetResourceComplianceByConfigRule"),
		Version:     tea.String("2020-09-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetResourceComplianceByConfigRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetResourceComplianceByConfigRule(request *GetResourceComplianceByConfigRuleRequest) (_result *GetResourceComplianceByConfigRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetResourceComplianceByConfigRuleResponse{}
	_body, _err := client.GetResourceComplianceByConfigRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetResourceComplianceByPackWithOptions(request *GetResourceComplianceByPackRequest, runtime *util.RuntimeOptions) (_result *GetResourceComplianceByPackResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetResourceComplianceByPack"),
		Version:     tea.String("2020-09-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetResourceComplianceByPackResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetResourceComplianceByPack(request *GetResourceComplianceByPackRequest) (_result *GetResourceComplianceByPackResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetResourceComplianceByPackResponse{}
	_body, _err := client.GetResourceComplianceByPackWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetResourceComplianceGroupByRegionWithOptions(request *GetResourceComplianceGroupByRegionRequest, runtime *util.RuntimeOptions) (_result *GetResourceComplianceGroupByRegionResponse, _err error) {
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
		Action:      tea.String("GetResourceComplianceGroupByRegion"),
		Version:     tea.String("2020-09-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetResourceComplianceGroupByRegionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetResourceComplianceGroupByRegion(request *GetResourceComplianceGroupByRegionRequest) (_result *GetResourceComplianceGroupByRegionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetResourceComplianceGroupByRegionResponse{}
	_body, _err := client.GetResourceComplianceGroupByRegionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetResourceComplianceGroupByResourceTypeWithOptions(request *GetResourceComplianceGroupByResourceTypeRequest, runtime *util.RuntimeOptions) (_result *GetResourceComplianceGroupByResourceTypeResponse, _err error) {
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
		Action:      tea.String("GetResourceComplianceGroupByResourceType"),
		Version:     tea.String("2020-09-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetResourceComplianceGroupByResourceTypeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetResourceComplianceGroupByResourceType(request *GetResourceComplianceGroupByResourceTypeRequest) (_result *GetResourceComplianceGroupByResourceTypeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetResourceComplianceGroupByResourceTypeResponse{}
	_body, _err := client.GetResourceComplianceGroupByResourceTypeWithOptions(request, runtime)
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
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetResourceComplianceTimeline"),
		Version:     tea.String("2020-09-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
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
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetResourceConfigurationTimeline"),
		Version:     tea.String("2020-09-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
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

func (client *Client) IgnoreAggregateEvaluationResultsWithOptions(tmpReq *IgnoreAggregateEvaluationResultsRequest, runtime *util.RuntimeOptions) (_result *IgnoreAggregateEvaluationResultsResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &IgnoreAggregateEvaluationResultsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Resources)) {
		request.ResourcesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Resources, tea.String("Resources"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AggregatorId)) {
		body["AggregatorId"] = request.AggregatorId
	}

	if !tea.BoolValue(util.IsUnset(request.ConfigRuleId)) {
		body["ConfigRuleId"] = request.ConfigRuleId
	}

	if !tea.BoolValue(util.IsUnset(request.IgnoreDate)) {
		body["IgnoreDate"] = request.IgnoreDate
	}

	if !tea.BoolValue(util.IsUnset(request.Reason)) {
		body["Reason"] = request.Reason
	}

	if !tea.BoolValue(util.IsUnset(request.ResourcesShrink)) {
		body["Resources"] = request.ResourcesShrink
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("IgnoreAggregateEvaluationResults"),
		Version:     tea.String("2020-09-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &IgnoreAggregateEvaluationResultsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) IgnoreAggregateEvaluationResults(request *IgnoreAggregateEvaluationResultsRequest) (_result *IgnoreAggregateEvaluationResultsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &IgnoreAggregateEvaluationResultsResponse{}
	_body, _err := client.IgnoreAggregateEvaluationResultsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) IgnoreEvaluationResultsWithOptions(tmpReq *IgnoreEvaluationResultsRequest, runtime *util.RuntimeOptions) (_result *IgnoreEvaluationResultsResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &IgnoreEvaluationResultsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Resources)) {
		request.ResourcesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Resources, tea.String("Resources"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConfigRuleId)) {
		body["ConfigRuleId"] = request.ConfigRuleId
	}

	if !tea.BoolValue(util.IsUnset(request.IgnoreDate)) {
		body["IgnoreDate"] = request.IgnoreDate
	}

	if !tea.BoolValue(util.IsUnset(request.Reason)) {
		body["Reason"] = request.Reason
	}

	if !tea.BoolValue(util.IsUnset(request.ResourcesShrink)) {
		body["Resources"] = request.ResourcesShrink
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("IgnoreEvaluationResults"),
		Version:     tea.String("2020-09-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &IgnoreEvaluationResultsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) IgnoreEvaluationResults(request *IgnoreEvaluationResultsRequest) (_result *IgnoreEvaluationResultsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &IgnoreEvaluationResultsResponse{}
	_body, _err := client.IgnoreEvaluationResultsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListAggregateCompliancePacksWithOptions(request *ListAggregateCompliancePacksRequest, runtime *util.RuntimeOptions) (_result *ListAggregateCompliancePacksResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListAggregateCompliancePacks"),
		Version:     tea.String("2020-09-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListAggregateCompliancePacksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListAggregateCompliancePacks(request *ListAggregateCompliancePacksRequest) (_result *ListAggregateCompliancePacksResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListAggregateCompliancePacksResponse{}
	_body, _err := client.ListAggregateCompliancePacksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListAggregateConfigDeliveryChannelsWithOptions(request *ListAggregateConfigDeliveryChannelsRequest, runtime *util.RuntimeOptions) (_result *ListAggregateConfigDeliveryChannelsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AggregatorId)) {
		query["AggregatorId"] = request.AggregatorId
	}

	if !tea.BoolValue(util.IsUnset(request.DeliveryChannelIds)) {
		query["DeliveryChannelIds"] = request.DeliveryChannelIds
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListAggregateConfigDeliveryChannels"),
		Version:     tea.String("2020-09-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListAggregateConfigDeliveryChannelsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListAggregateConfigDeliveryChannels(request *ListAggregateConfigDeliveryChannelsRequest) (_result *ListAggregateConfigDeliveryChannelsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListAggregateConfigDeliveryChannelsResponse{}
	_body, _err := client.ListAggregateConfigDeliveryChannelsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListAggregateConfigRuleEvaluationResultsWithOptions(request *ListAggregateConfigRuleEvaluationResultsRequest, runtime *util.RuntimeOptions) (_result *ListAggregateConfigRuleEvaluationResultsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListAggregateConfigRuleEvaluationResults"),
		Version:     tea.String("2020-09-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListAggregateConfigRuleEvaluationResultsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListAggregateConfigRuleEvaluationResults(request *ListAggregateConfigRuleEvaluationResultsRequest) (_result *ListAggregateConfigRuleEvaluationResultsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListAggregateConfigRuleEvaluationResultsResponse{}
	_body, _err := client.ListAggregateConfigRuleEvaluationResultsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListAggregateConfigRulesWithOptions(request *ListAggregateConfigRulesRequest, runtime *util.RuntimeOptions) (_result *ListAggregateConfigRulesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AggregatorId)) {
		query["AggregatorId"] = request.AggregatorId
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
		Action:      tea.String("ListAggregateConfigRules"),
		Version:     tea.String("2020-09-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListAggregateConfigRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListAggregateConfigRules(request *ListAggregateConfigRulesRequest) (_result *ListAggregateConfigRulesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListAggregateConfigRulesResponse{}
	_body, _err := client.ListAggregateConfigRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListAggregateDiscoveredResourcesWithOptions(request *ListAggregateDiscoveredResourcesRequest, runtime *util.RuntimeOptions) (_result *ListAggregateDiscoveredResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AggregatorId)) {
		query["AggregatorId"] = request.AggregatorId
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
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
		Version:     tea.String("2020-09-07"),
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

func (client *Client) ListAggregateRemediationsWithOptions(request *ListAggregateRemediationsRequest, runtime *util.RuntimeOptions) (_result *ListAggregateRemediationsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AggregatorId)) {
		query["AggregatorId"] = request.AggregatorId
	}

	if !tea.BoolValue(util.IsUnset(request.ConfigRuleIds)) {
		query["ConfigRuleIds"] = request.ConfigRuleIds
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListAggregateRemediations"),
		Version:     tea.String("2020-09-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListAggregateRemediationsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListAggregateRemediations(request *ListAggregateRemediationsRequest) (_result *ListAggregateRemediationsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListAggregateRemediationsResponse{}
	_body, _err := client.ListAggregateRemediationsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListAggregateResourceEvaluationResultsWithOptions(request *ListAggregateResourceEvaluationResultsRequest, runtime *util.RuntimeOptions) (_result *ListAggregateResourceEvaluationResultsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AggregatorId)) {
		query["AggregatorId"] = request.AggregatorId
	}

	if !tea.BoolValue(util.IsUnset(request.ComplianceType)) {
		query["ComplianceType"] = request.ComplianceType
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
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

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListAggregateResourceEvaluationResults"),
		Version:     tea.String("2020-09-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListAggregateResourceEvaluationResultsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListAggregateResourceEvaluationResults(request *ListAggregateResourceEvaluationResultsRequest) (_result *ListAggregateResourceEvaluationResultsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListAggregateResourceEvaluationResultsResponse{}
	_body, _err := client.ListAggregateResourceEvaluationResultsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListAggregatorsWithOptions(request *ListAggregatorsRequest, runtime *util.RuntimeOptions) (_result *ListAggregatorsResponse, _err error) {
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

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListAggregators"),
		Version:     tea.String("2020-09-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListAggregatorsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListAggregators(request *ListAggregatorsRequest) (_result *ListAggregatorsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListAggregatorsResponse{}
	_body, _err := client.ListAggregatorsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListCompliancePackTemplatesWithOptions(request *ListCompliancePackTemplatesRequest, runtime *util.RuntimeOptions) (_result *ListCompliancePackTemplatesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListCompliancePackTemplates"),
		Version:     tea.String("2020-09-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListCompliancePackTemplatesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListCompliancePackTemplates(request *ListCompliancePackTemplatesRequest) (_result *ListCompliancePackTemplatesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListCompliancePackTemplatesResponse{}
	_body, _err := client.ListCompliancePackTemplatesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListCompliancePacksWithOptions(request *ListCompliancePacksRequest, runtime *util.RuntimeOptions) (_result *ListCompliancePacksResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListCompliancePacks"),
		Version:     tea.String("2020-09-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListCompliancePacksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListCompliancePacks(request *ListCompliancePacksRequest) (_result *ListCompliancePacksResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListCompliancePacksResponse{}
	_body, _err := client.ListCompliancePacksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListConfigDeliveryChannelsWithOptions(request *ListConfigDeliveryChannelsRequest, runtime *util.RuntimeOptions) (_result *ListConfigDeliveryChannelsResponse, _err error) {
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
		Action:      tea.String("ListConfigDeliveryChannels"),
		Version:     tea.String("2020-09-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListConfigDeliveryChannelsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListConfigDeliveryChannels(request *ListConfigDeliveryChannelsRequest) (_result *ListConfigDeliveryChannelsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListConfigDeliveryChannelsResponse{}
	_body, _err := client.ListConfigDeliveryChannelsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListConfigRuleEvaluationResultsWithOptions(request *ListConfigRuleEvaluationResultsRequest, runtime *util.RuntimeOptions) (_result *ListConfigRuleEvaluationResultsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListConfigRuleEvaluationResults"),
		Version:     tea.String("2020-09-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListConfigRuleEvaluationResultsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListConfigRuleEvaluationResults(request *ListConfigRuleEvaluationResultsRequest) (_result *ListConfigRuleEvaluationResultsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListConfigRuleEvaluationResultsResponse{}
	_body, _err := client.ListConfigRuleEvaluationResultsWithOptions(request, runtime)
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
	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
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
		Version:     tea.String("2020-09-07"),
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

func (client *Client) ListManagedRulesWithOptions(request *ListManagedRulesRequest, runtime *util.RuntimeOptions) (_result *ListManagedRulesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Keyword)) {
		query["Keyword"] = request.Keyword
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
		Action:      tea.String("ListManagedRules"),
		Version:     tea.String("2020-09-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListManagedRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListManagedRules(request *ListManagedRulesRequest) (_result *ListManagedRulesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListManagedRulesResponse{}
	_body, _err := client.ListManagedRulesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

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
		Version:     tea.String("2020-09-07"),
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

func (client *Client) ListRemediationsWithOptions(request *ListRemediationsRequest, runtime *util.RuntimeOptions) (_result *ListRemediationsResponse, _err error) {
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
		Action:      tea.String("ListRemediations"),
		Version:     tea.String("2020-09-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListRemediationsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListRemediations(request *ListRemediationsRequest) (_result *ListRemediationsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListRemediationsResponse{}
	_body, _err := client.ListRemediationsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListResourceEvaluationResultsWithOptions(request *ListResourceEvaluationResultsRequest, runtime *util.RuntimeOptions) (_result *ListResourceEvaluationResultsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ComplianceType)) {
		query["ComplianceType"] = request.ComplianceType
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
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

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListResourceEvaluationResults"),
		Version:     tea.String("2020-09-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListResourceEvaluationResultsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListResourceEvaluationResults(request *ListResourceEvaluationResultsRequest) (_result *ListResourceEvaluationResultsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListResourceEvaluationResultsResponse{}
	_body, _err := client.ListResourceEvaluationResultsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListTagResourcesWithOptions(tmpReq *ListTagResourcesRequest, runtime *util.RuntimeOptions) (_result *ListTagResourcesResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ListTagResourcesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Tag)) {
		request.TagShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tag, tea.String("Tag"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		body["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		body["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		body["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.TagShrink)) {
		body["Tag"] = request.TagShrink
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTagResources"),
		Version:     tea.String("2020-09-07"),
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

func (client *Client) RevertAggregateEvaluationResultsWithOptions(tmpReq *RevertAggregateEvaluationResultsRequest, runtime *util.RuntimeOptions) (_result *RevertAggregateEvaluationResultsResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &RevertAggregateEvaluationResultsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Resources)) {
		request.ResourcesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Resources, tea.String("Resources"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AggregatorId)) {
		body["AggregatorId"] = request.AggregatorId
	}

	if !tea.BoolValue(util.IsUnset(request.ConfigRuleId)) {
		body["ConfigRuleId"] = request.ConfigRuleId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourcesShrink)) {
		body["Resources"] = request.ResourcesShrink
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RevertAggregateEvaluationResults"),
		Version:     tea.String("2020-09-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RevertAggregateEvaluationResultsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RevertAggregateEvaluationResults(request *RevertAggregateEvaluationResultsRequest) (_result *RevertAggregateEvaluationResultsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RevertAggregateEvaluationResultsResponse{}
	_body, _err := client.RevertAggregateEvaluationResultsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RevertEvaluationResultsWithOptions(tmpReq *RevertEvaluationResultsRequest, runtime *util.RuntimeOptions) (_result *RevertEvaluationResultsResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &RevertEvaluationResultsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Resources)) {
		request.ResourcesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Resources, tea.String("Resources"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConfigRuleId)) {
		body["ConfigRuleId"] = request.ConfigRuleId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourcesShrink)) {
		body["Resources"] = request.ResourcesShrink
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RevertEvaluationResults"),
		Version:     tea.String("2020-09-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RevertEvaluationResultsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RevertEvaluationResults(request *RevertEvaluationResultsRequest) (_result *RevertEvaluationResultsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RevertEvaluationResultsResponse{}
	_body, _err := client.RevertEvaluationResultsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StartAggregateConfigRuleEvaluationWithOptions(request *StartAggregateConfigRuleEvaluationRequest, runtime *util.RuntimeOptions) (_result *StartAggregateConfigRuleEvaluationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AggregatorId)) {
		query["AggregatorId"] = request.AggregatorId
	}

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
		Action:      tea.String("StartAggregateConfigRuleEvaluation"),
		Version:     tea.String("2020-09-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &StartAggregateConfigRuleEvaluationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StartAggregateConfigRuleEvaluation(request *StartAggregateConfigRuleEvaluationRequest) (_result *StartAggregateConfigRuleEvaluationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StartAggregateConfigRuleEvaluationResponse{}
	_body, _err := client.StartAggregateConfigRuleEvaluationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StartAggregateRemediationWithOptions(request *StartAggregateRemediationRequest, runtime *util.RuntimeOptions) (_result *StartAggregateRemediationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AggregatorId)) {
		query["AggregatorId"] = request.AggregatorId
	}

	if !tea.BoolValue(util.IsUnset(request.ConfigRuleId)) {
		query["ConfigRuleId"] = request.ConfigRuleId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("StartAggregateRemediation"),
		Version:     tea.String("2020-09-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &StartAggregateRemediationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StartAggregateRemediation(request *StartAggregateRemediationRequest) (_result *StartAggregateRemediationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StartAggregateRemediationResponse{}
	_body, _err := client.StartAggregateRemediationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StartRemediationWithOptions(request *StartRemediationRequest, runtime *util.RuntimeOptions) (_result *StartRemediationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConfigRuleId)) {
		query["ConfigRuleId"] = request.ConfigRuleId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("StartRemediation"),
		Version:     tea.String("2020-09-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &StartRemediationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StartRemediation(request *StartRemediationRequest) (_result *StartRemediationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StartRemediationResponse{}
	_body, _err := client.StartRemediationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) TagResourcesWithOptions(tmpReq *TagResourcesRequest, runtime *util.RuntimeOptions) (_result *TagResourcesResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &TagResourcesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Tag)) {
		request.TagShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tag, tea.String("Tag"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		body["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		body["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.TagShrink)) {
		body["Tag"] = request.TagShrink
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("TagResources"),
		Version:     tea.String("2020-09-07"),
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
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.All)) {
		body["All"] = request.All
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		body["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		body["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.TagKey)) {
		body["TagKey"] = request.TagKey
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UntagResources"),
		Version:     tea.String("2020-09-07"),
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

func (client *Client) UpdateAggregateCompliancePackWithOptions(tmpReq *UpdateAggregateCompliancePackRequest, runtime *util.RuntimeOptions) (_result *UpdateAggregateCompliancePackResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &UpdateAggregateCompliancePackShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.ConfigRules)) {
		request.ConfigRulesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ConfigRules, tea.String("ConfigRules"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AggregatorId)) {
		body["AggregatorId"] = request.AggregatorId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.CompliancePackId)) {
		body["CompliancePackId"] = request.CompliancePackId
	}

	if !tea.BoolValue(util.IsUnset(request.CompliancePackName)) {
		body["CompliancePackName"] = request.CompliancePackName
	}

	if !tea.BoolValue(util.IsUnset(request.ConfigRulesShrink)) {
		body["ConfigRules"] = request.ConfigRulesShrink
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.RiskLevel)) {
		body["RiskLevel"] = request.RiskLevel
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateAggregateCompliancePack"),
		Version:     tea.String("2020-09-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateAggregateCompliancePackResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateAggregateCompliancePack(request *UpdateAggregateCompliancePackRequest) (_result *UpdateAggregateCompliancePackResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateAggregateCompliancePackResponse{}
	_body, _err := client.UpdateAggregateCompliancePackWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateAggregateConfigDeliveryChannelWithOptions(request *UpdateAggregateConfigDeliveryChannelRequest, runtime *util.RuntimeOptions) (_result *UpdateAggregateConfigDeliveryChannelResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AggregatorId)) {
		query["AggregatorId"] = request.AggregatorId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.ConfigurationItemChangeNotification)) {
		query["ConfigurationItemChangeNotification"] = request.ConfigurationItemChangeNotification
	}

	if !tea.BoolValue(util.IsUnset(request.ConfigurationSnapshot)) {
		query["ConfigurationSnapshot"] = request.ConfigurationSnapshot
	}

	if !tea.BoolValue(util.IsUnset(request.DeliveryChannelCondition)) {
		query["DeliveryChannelCondition"] = request.DeliveryChannelCondition
	}

	if !tea.BoolValue(util.IsUnset(request.DeliveryChannelId)) {
		query["DeliveryChannelId"] = request.DeliveryChannelId
	}

	if !tea.BoolValue(util.IsUnset(request.DeliveryChannelName)) {
		query["DeliveryChannelName"] = request.DeliveryChannelName
	}

	if !tea.BoolValue(util.IsUnset(request.DeliveryChannelTargetArn)) {
		query["DeliveryChannelTargetArn"] = request.DeliveryChannelTargetArn
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.NonCompliantNotification)) {
		query["NonCompliantNotification"] = request.NonCompliantNotification
	}

	if !tea.BoolValue(util.IsUnset(request.OversizedDataOSSTargetArn)) {
		query["OversizedDataOSSTargetArn"] = request.OversizedDataOSSTargetArn
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateAggregateConfigDeliveryChannel"),
		Version:     tea.String("2020-09-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateAggregateConfigDeliveryChannelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateAggregateConfigDeliveryChannel(request *UpdateAggregateConfigDeliveryChannelRequest) (_result *UpdateAggregateConfigDeliveryChannelResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateAggregateConfigDeliveryChannelResponse{}
	_body, _err := client.UpdateAggregateConfigDeliveryChannelWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateAggregateConfigRuleWithOptions(tmpReq *UpdateAggregateConfigRuleRequest, runtime *util.RuntimeOptions) (_result *UpdateAggregateConfigRuleResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &UpdateAggregateConfigRuleShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.InputParameters)) {
		request.InputParametersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.InputParameters, tea.String("InputParameters"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.ResourceTypesScope)) {
		request.ResourceTypesScopeShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ResourceTypesScope, tea.String("ResourceTypesScope"), tea.String("simple"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AggregatorId)) {
		body["AggregatorId"] = request.AggregatorId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.ConfigRuleId)) {
		body["ConfigRuleId"] = request.ConfigRuleId
	}

	if !tea.BoolValue(util.IsUnset(request.ConfigRuleName)) {
		body["ConfigRuleName"] = request.ConfigRuleName
	}

	if !tea.BoolValue(util.IsUnset(request.ConfigRuleTriggerTypes)) {
		body["ConfigRuleTriggerTypes"] = request.ConfigRuleTriggerTypes
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.ExcludeAccountIdsScope)) {
		body["ExcludeAccountIdsScope"] = request.ExcludeAccountIdsScope
	}

	if !tea.BoolValue(util.IsUnset(request.ExcludeFolderIdsScope)) {
		body["ExcludeFolderIdsScope"] = request.ExcludeFolderIdsScope
	}

	if !tea.BoolValue(util.IsUnset(request.ExcludeResourceIdsScope)) {
		body["ExcludeResourceIdsScope"] = request.ExcludeResourceIdsScope
	}

	if !tea.BoolValue(util.IsUnset(request.FolderIdsScope)) {
		body["FolderIdsScope"] = request.FolderIdsScope
	}

	if !tea.BoolValue(util.IsUnset(request.InputParametersShrink)) {
		body["InputParameters"] = request.InputParametersShrink
	}

	if !tea.BoolValue(util.IsUnset(request.MaximumExecutionFrequency)) {
		body["MaximumExecutionFrequency"] = request.MaximumExecutionFrequency
	}

	if !tea.BoolValue(util.IsUnset(request.RegionIdsScope)) {
		body["RegionIdsScope"] = request.RegionIdsScope
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupIdsScope)) {
		body["ResourceGroupIdsScope"] = request.ResourceGroupIdsScope
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceTypesScopeShrink)) {
		body["ResourceTypesScope"] = request.ResourceTypesScopeShrink
	}

	if !tea.BoolValue(util.IsUnset(request.RiskLevel)) {
		body["RiskLevel"] = request.RiskLevel
	}

	if !tea.BoolValue(util.IsUnset(request.TagKeyLogicScope)) {
		body["TagKeyLogicScope"] = request.TagKeyLogicScope
	}

	if !tea.BoolValue(util.IsUnset(request.TagKeyScope)) {
		body["TagKeyScope"] = request.TagKeyScope
	}

	if !tea.BoolValue(util.IsUnset(request.TagValueScope)) {
		body["TagValueScope"] = request.TagValueScope
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateAggregateConfigRule"),
		Version:     tea.String("2020-09-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateAggregateConfigRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateAggregateConfigRule(request *UpdateAggregateConfigRuleRequest) (_result *UpdateAggregateConfigRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateAggregateConfigRuleResponse{}
	_body, _err := client.UpdateAggregateConfigRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateAggregateRemediationWithOptions(request *UpdateAggregateRemediationRequest, runtime *util.RuntimeOptions) (_result *UpdateAggregateRemediationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AggregatorId)) {
		body["AggregatorId"] = request.AggregatorId
	}

	if !tea.BoolValue(util.IsUnset(request.InvokeType)) {
		body["InvokeType"] = request.InvokeType
	}

	if !tea.BoolValue(util.IsUnset(request.Params)) {
		body["Params"] = request.Params
	}

	if !tea.BoolValue(util.IsUnset(request.RemediationId)) {
		body["RemediationId"] = request.RemediationId
	}

	if !tea.BoolValue(util.IsUnset(request.RemediationTemplateId)) {
		body["RemediationTemplateId"] = request.RemediationTemplateId
	}

	if !tea.BoolValue(util.IsUnset(request.RemediationType)) {
		body["RemediationType"] = request.RemediationType
	}

	if !tea.BoolValue(util.IsUnset(request.SourceType)) {
		body["SourceType"] = request.SourceType
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateAggregateRemediation"),
		Version:     tea.String("2020-09-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateAggregateRemediationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateAggregateRemediation(request *UpdateAggregateRemediationRequest) (_result *UpdateAggregateRemediationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateAggregateRemediationResponse{}
	_body, _err := client.UpdateAggregateRemediationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateAggregatorWithOptions(tmpReq *UpdateAggregatorRequest, runtime *util.RuntimeOptions) (_result *UpdateAggregatorResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &UpdateAggregatorShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.AggregatorAccounts)) {
		request.AggregatorAccountsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AggregatorAccounts, tea.String("AggregatorAccounts"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AggregatorAccountsShrink)) {
		body["AggregatorAccounts"] = request.AggregatorAccountsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.AggregatorId)) {
		body["AggregatorId"] = request.AggregatorId
	}

	if !tea.BoolValue(util.IsUnset(request.AggregatorName)) {
		body["AggregatorName"] = request.AggregatorName
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateAggregator"),
		Version:     tea.String("2020-09-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateAggregatorResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateAggregator(request *UpdateAggregatorRequest) (_result *UpdateAggregatorResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateAggregatorResponse{}
	_body, _err := client.UpdateAggregatorWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateCompliancePackWithOptions(tmpReq *UpdateCompliancePackRequest, runtime *util.RuntimeOptions) (_result *UpdateCompliancePackResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &UpdateCompliancePackShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.ConfigRules)) {
		request.ConfigRulesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ConfigRules, tea.String("ConfigRules"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.CompliancePackId)) {
		body["CompliancePackId"] = request.CompliancePackId
	}

	if !tea.BoolValue(util.IsUnset(request.CompliancePackName)) {
		body["CompliancePackName"] = request.CompliancePackName
	}

	if !tea.BoolValue(util.IsUnset(request.ConfigRulesShrink)) {
		body["ConfigRules"] = request.ConfigRulesShrink
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.RiskLevel)) {
		body["RiskLevel"] = request.RiskLevel
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateCompliancePack"),
		Version:     tea.String("2020-09-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateCompliancePackResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateCompliancePack(request *UpdateCompliancePackRequest) (_result *UpdateCompliancePackResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateCompliancePackResponse{}
	_body, _err := client.UpdateCompliancePackWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateConfigDeliveryChannelWithOptions(request *UpdateConfigDeliveryChannelRequest, runtime *util.RuntimeOptions) (_result *UpdateConfigDeliveryChannelResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.ConfigurationItemChangeNotification)) {
		query["ConfigurationItemChangeNotification"] = request.ConfigurationItemChangeNotification
	}

	if !tea.BoolValue(util.IsUnset(request.ConfigurationSnapshot)) {
		query["ConfigurationSnapshot"] = request.ConfigurationSnapshot
	}

	if !tea.BoolValue(util.IsUnset(request.DeliveryChannelCondition)) {
		query["DeliveryChannelCondition"] = request.DeliveryChannelCondition
	}

	if !tea.BoolValue(util.IsUnset(request.DeliveryChannelId)) {
		query["DeliveryChannelId"] = request.DeliveryChannelId
	}

	if !tea.BoolValue(util.IsUnset(request.DeliveryChannelName)) {
		query["DeliveryChannelName"] = request.DeliveryChannelName
	}

	if !tea.BoolValue(util.IsUnset(request.DeliveryChannelTargetArn)) {
		query["DeliveryChannelTargetArn"] = request.DeliveryChannelTargetArn
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.NonCompliantNotification)) {
		query["NonCompliantNotification"] = request.NonCompliantNotification
	}

	if !tea.BoolValue(util.IsUnset(request.OversizedDataOSSTargetArn)) {
		query["OversizedDataOSSTargetArn"] = request.OversizedDataOSSTargetArn
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateConfigDeliveryChannel"),
		Version:     tea.String("2020-09-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateConfigDeliveryChannelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateConfigDeliveryChannel(request *UpdateConfigDeliveryChannelRequest) (_result *UpdateConfigDeliveryChannelResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateConfigDeliveryChannelResponse{}
	_body, _err := client.UpdateConfigDeliveryChannelWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateConfigRuleWithOptions(tmpReq *UpdateConfigRuleRequest, runtime *util.RuntimeOptions) (_result *UpdateConfigRuleResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &UpdateConfigRuleShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.InputParameters)) {
		request.InputParametersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.InputParameters, tea.String("InputParameters"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.ResourceTypesScope)) {
		request.ResourceTypesScopeShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ResourceTypesScope, tea.String("ResourceTypesScope"), tea.String("simple"))
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

	if !tea.BoolValue(util.IsUnset(request.ConfigRuleTriggerTypes)) {
		body["ConfigRuleTriggerTypes"] = request.ConfigRuleTriggerTypes
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.ExcludeResourceIdsScope)) {
		body["ExcludeResourceIdsScope"] = request.ExcludeResourceIdsScope
	}

	if !tea.BoolValue(util.IsUnset(request.InputParametersShrink)) {
		body["InputParameters"] = request.InputParametersShrink
	}

	if !tea.BoolValue(util.IsUnset(request.MaximumExecutionFrequency)) {
		body["MaximumExecutionFrequency"] = request.MaximumExecutionFrequency
	}

	if !tea.BoolValue(util.IsUnset(request.RegionIdsScope)) {
		body["RegionIdsScope"] = request.RegionIdsScope
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupIdsScope)) {
		body["ResourceGroupIdsScope"] = request.ResourceGroupIdsScope
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceTypesScopeShrink)) {
		body["ResourceTypesScope"] = request.ResourceTypesScopeShrink
	}

	if !tea.BoolValue(util.IsUnset(request.RiskLevel)) {
		body["RiskLevel"] = request.RiskLevel
	}

	if !tea.BoolValue(util.IsUnset(request.TagKeyLogicScope)) {
		body["TagKeyLogicScope"] = request.TagKeyLogicScope
	}

	if !tea.BoolValue(util.IsUnset(request.TagKeyScope)) {
		body["TagKeyScope"] = request.TagKeyScope
	}

	if !tea.BoolValue(util.IsUnset(request.TagValueScope)) {
		body["TagValueScope"] = request.TagValueScope
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateConfigRule"),
		Version:     tea.String("2020-09-07"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateConfigRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateConfigRule(request *UpdateConfigRuleRequest) (_result *UpdateConfigRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateConfigRuleResponse{}
	_body, _err := client.UpdateConfigRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
