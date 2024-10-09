// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type CheckRulesRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// acs:ecs:123***890:cn-shanghai:instance/i-001***90
	ResourceArn *string `json:"ResourceArn,omitempty" xml:"ResourceArn,omitempty"`
	// example:
	//
	// rule-000***dav
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s CheckRulesRequest) String() string {
	return tea.Prettify(s)
}

func (s CheckRulesRequest) GoString() string {
	return s.String()
}

func (s *CheckRulesRequest) SetResourceArn(v string) *CheckRulesRequest {
	s.ResourceArn = &v
	return s
}

func (s *CheckRulesRequest) SetRuleId(v string) *CheckRulesRequest {
	s.RuleId = &v
	return s
}

type CheckRulesResponseBody struct {
	Data *CheckRulesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 700683DE-0154-56D4-8D76-3B7A2C2C7DF9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CheckRulesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CheckRulesResponseBody) GoString() string {
	return s.String()
}

func (s *CheckRulesResponseBody) SetData(v *CheckRulesResponseBodyData) *CheckRulesResponseBody {
	s.Data = v
	return s
}

func (s *CheckRulesResponseBody) SetRequestId(v string) *CheckRulesResponseBody {
	s.RequestId = &v
	return s
}

type CheckRulesResponseBodyData struct {
	// example:
	//
	// t-0000e4w0u1v592zdf6s7
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CheckRulesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CheckRulesResponseBodyData) GoString() string {
	return s.String()
}

func (s *CheckRulesResponseBodyData) SetTaskId(v string) *CheckRulesResponseBodyData {
	s.TaskId = &v
	return s
}

type CheckRulesResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckRulesResponse) String() string {
	return tea.Prettify(s)
}

func (s CheckRulesResponse) GoString() string {
	return s.String()
}

func (s *CheckRulesResponse) SetHeaders(v map[string]*string) *CheckRulesResponse {
	s.Headers = v
	return s
}

func (s *CheckRulesResponse) SetStatusCode(v int32) *CheckRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckRulesResponse) SetBody(v *CheckRulesResponseBody) *CheckRulesResponse {
	s.Body = v
	return s
}

type CloseBdrcServiceResponseBody struct {
	// example:
	//
	// 663D8898-E0B5-5964-BF28-A191CE6A1825
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CloseBdrcServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CloseBdrcServiceResponseBody) GoString() string {
	return s.String()
}

func (s *CloseBdrcServiceResponseBody) SetRequestId(v string) *CloseBdrcServiceResponseBody {
	s.RequestId = &v
	return s
}

type CloseBdrcServiceResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CloseBdrcServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CloseBdrcServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s CloseBdrcServiceResponse) GoString() string {
	return s.String()
}

func (s *CloseBdrcServiceResponse) SetHeaders(v map[string]*string) *CloseBdrcServiceResponse {
	s.Headers = v
	return s
}

func (s *CloseBdrcServiceResponse) SetStatusCode(v int32) *CloseBdrcServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *CloseBdrcServiceResponse) SetBody(v *CloseBdrcServiceResponseBody) *CloseBdrcServiceResponse {
	s.Body = v
	return s
}

type DescribeCheckDetailsRequest struct {
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// cae**********699
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// acs:ecs:123***890:cn-shanghai:instance/i-001***90
	ResourceArn *string `json:"ResourceArn,omitempty" xml:"ResourceArn,omitempty"`
	// example:
	//
	// rule-000***dav
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s DescribeCheckDetailsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCheckDetailsRequest) GoString() string {
	return s.String()
}

func (s *DescribeCheckDetailsRequest) SetMaxResults(v int32) *DescribeCheckDetailsRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeCheckDetailsRequest) SetNextToken(v string) *DescribeCheckDetailsRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeCheckDetailsRequest) SetResourceArn(v string) *DescribeCheckDetailsRequest {
	s.ResourceArn = &v
	return s
}

func (s *DescribeCheckDetailsRequest) SetRuleId(v string) *DescribeCheckDetailsRequest {
	s.RuleId = &v
	return s
}

type DescribeCheckDetailsResponseBody struct {
	Data *DescribeCheckDetailsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 92793A50-0B97-59F1-BAEA-EAED83BA1998
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeCheckDetailsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeCheckDetailsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCheckDetailsResponseBody) SetData(v *DescribeCheckDetailsResponseBodyData) *DescribeCheckDetailsResponseBody {
	s.Data = v
	return s
}

func (s *DescribeCheckDetailsResponseBody) SetRequestId(v string) *DescribeCheckDetailsResponseBody {
	s.RequestId = &v
	return s
}

type DescribeCheckDetailsResponseBodyData struct {
	Content []*DescribeCheckDetailsResponseBodyDataContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Repeated"`
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// CAESGgoSChAKDGNvbXBsZXRlVGltZRABCgQiAggAGAAiQAoJAOTzWWYAAAAACjMDLgAAADFTNzMyZDMwMzAzMDM4NzA3NTcwMzY2MjMwNzY2ODcyMzAzMTY2Nzg3ODY5MzY=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// 100
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeCheckDetailsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeCheckDetailsResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeCheckDetailsResponseBodyData) SetContent(v []*DescribeCheckDetailsResponseBodyDataContent) *DescribeCheckDetailsResponseBodyData {
	s.Content = v
	return s
}

func (s *DescribeCheckDetailsResponseBodyData) SetMaxResults(v int32) *DescribeCheckDetailsResponseBodyData {
	s.MaxResults = &v
	return s
}

func (s *DescribeCheckDetailsResponseBodyData) SetNextToken(v string) *DescribeCheckDetailsResponseBodyData {
	s.NextToken = &v
	return s
}

func (s *DescribeCheckDetailsResponseBodyData) SetTotalCount(v int64) *DescribeCheckDetailsResponseBodyData {
	s.TotalCount = &v
	return s
}

type DescribeCheckDetailsResponseBodyDataContent struct {
	// example:
	//
	// PASSED
	CheckStatus *string `json:"CheckStatus,omitempty" xml:"CheckStatus,omitempty"`
	// example:
	//
	// 1701725715
	CheckTime *int64 `json:"CheckTime,omitempty" xml:"CheckTime,omitempty"`
	// example:
	//
	// {"ecsAutoSnapshotPolicyIds":[],"hbrBackupPlans":[{"planId":"po-xxxxxxxx","sourceType":"UDM_ECS"}]}
	Detail *string `json:"Detail,omitempty" xml:"Detail,omitempty"`
	// example:
	//
	// ecs
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	// example:
	//
	// acs:ecs:123***890:cn-shanghai:instance/i-001***90
	ResourceArn *string `json:"ResourceArn,omitempty" xml:"ResourceArn,omitempty"`
	// example:
	//
	// i-xxxxxxxx
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// example:
	//
	// test server
	ResourceName *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	// example:
	//
	// ACS::ECS::Instance
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// example:
	//
	// rule-xxxxxxxx
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// example:
	//
	// ecs-backup
	RuleTemplate *string `json:"RuleTemplate,omitempty" xml:"RuleTemplate,omitempty"`
}

func (s DescribeCheckDetailsResponseBodyDataContent) String() string {
	return tea.Prettify(s)
}

func (s DescribeCheckDetailsResponseBodyDataContent) GoString() string {
	return s.String()
}

func (s *DescribeCheckDetailsResponseBodyDataContent) SetCheckStatus(v string) *DescribeCheckDetailsResponseBodyDataContent {
	s.CheckStatus = &v
	return s
}

func (s *DescribeCheckDetailsResponseBodyDataContent) SetCheckTime(v int64) *DescribeCheckDetailsResponseBodyDataContent {
	s.CheckTime = &v
	return s
}

func (s *DescribeCheckDetailsResponseBodyDataContent) SetDetail(v string) *DescribeCheckDetailsResponseBodyDataContent {
	s.Detail = &v
	return s
}

func (s *DescribeCheckDetailsResponseBodyDataContent) SetProductType(v string) *DescribeCheckDetailsResponseBodyDataContent {
	s.ProductType = &v
	return s
}

func (s *DescribeCheckDetailsResponseBodyDataContent) SetResourceArn(v string) *DescribeCheckDetailsResponseBodyDataContent {
	s.ResourceArn = &v
	return s
}

func (s *DescribeCheckDetailsResponseBodyDataContent) SetResourceId(v string) *DescribeCheckDetailsResponseBodyDataContent {
	s.ResourceId = &v
	return s
}

func (s *DescribeCheckDetailsResponseBodyDataContent) SetResourceName(v string) *DescribeCheckDetailsResponseBodyDataContent {
	s.ResourceName = &v
	return s
}

func (s *DescribeCheckDetailsResponseBodyDataContent) SetResourceType(v string) *DescribeCheckDetailsResponseBodyDataContent {
	s.ResourceType = &v
	return s
}

func (s *DescribeCheckDetailsResponseBodyDataContent) SetRuleId(v string) *DescribeCheckDetailsResponseBodyDataContent {
	s.RuleId = &v
	return s
}

func (s *DescribeCheckDetailsResponseBodyDataContent) SetRuleTemplate(v string) *DescribeCheckDetailsResponseBodyDataContent {
	s.RuleTemplate = &v
	return s
}

type DescribeCheckDetailsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCheckDetailsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCheckDetailsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeCheckDetailsResponse) GoString() string {
	return s.String()
}

func (s *DescribeCheckDetailsResponse) SetHeaders(v map[string]*string) *DescribeCheckDetailsResponse {
	s.Headers = v
	return s
}

func (s *DescribeCheckDetailsResponse) SetStatusCode(v int32) *DescribeCheckDetailsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCheckDetailsResponse) SetBody(v *DescribeCheckDetailsResponseBody) *DescribeCheckDetailsResponse {
	s.Body = v
	return s
}

type DescribeProductsResponseBody struct {
	Data *DescribeProductsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 30FB202A-1D22-5394-AB02-4477CDFCF51F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeProductsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeProductsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeProductsResponseBody) SetData(v *DescribeProductsResponseBodyData) *DescribeProductsResponseBody {
	s.Data = v
	return s
}

func (s *DescribeProductsResponseBody) SetRequestId(v string) *DescribeProductsResponseBody {
	s.RequestId = &v
	return s
}

type DescribeProductsResponseBodyData struct {
	Content []*DescribeProductsResponseBodyDataContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Repeated"`
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// b4fd3cffcacafd65e3818a0b9b2ff9a2
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// 50
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeProductsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeProductsResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeProductsResponseBodyData) SetContent(v []*DescribeProductsResponseBodyDataContent) *DescribeProductsResponseBodyData {
	s.Content = v
	return s
}

func (s *DescribeProductsResponseBodyData) SetMaxResults(v int32) *DescribeProductsResponseBodyData {
	s.MaxResults = &v
	return s
}

func (s *DescribeProductsResponseBodyData) SetNextToken(v string) *DescribeProductsResponseBodyData {
	s.NextToken = &v
	return s
}

func (s *DescribeProductsResponseBodyData) SetTotalCount(v int64) *DescribeProductsResponseBodyData {
	s.TotalCount = &v
	return s
}

type DescribeProductsResponseBodyDataContent struct {
	// example:
	//
	// 1
	CheckFailedCount *int64 `json:"CheckFailedCount,omitempty" xml:"CheckFailedCount,omitempty"`
	// example:
	//
	// 1
	CheckFailedResourceCount *int64 `json:"CheckFailedResourceCount,omitempty" xml:"CheckFailedResourceCount,omitempty"`
	// example:
	//
	// 1
	DisableCheckResourceCount *int64 `json:"DisableCheckResourceCount,omitempty" xml:"DisableCheckResourceCount,omitempty"`
	// example:
	//
	// true
	EnableCheck *bool `json:"EnableCheck,omitempty" xml:"EnableCheck,omitempty"`
	// example:
	//
	// oss
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	// example:
	//
	// 90
	ProtectionScore             *int32                                                                `json:"ProtectionScore,omitempty" xml:"ProtectionScore,omitempty"`
	ProtectionScoreDistribution []*DescribeProductsResponseBodyDataContentProtectionScoreDistribution `json:"ProtectionScoreDistribution,omitempty" xml:"ProtectionScoreDistribution,omitempty" type:"Repeated"`
	// example:
	//
	// 1726036498
	ProtectionScoreUpdatedTime *int64 `json:"ProtectionScoreUpdatedTime,omitempty" xml:"ProtectionScoreUpdatedTime,omitempty"`
	// example:
	//
	// 1
	RiskCount *int64 `json:"RiskCount,omitempty" xml:"RiskCount,omitempty"`
	// example:
	//
	// 1
	RiskyResourceCount *int64 `json:"RiskyResourceCount,omitempty" xml:"RiskyResourceCount,omitempty"`
	// example:
	//
	// 100
	TotalResourceCount *int64 `json:"TotalResourceCount,omitempty" xml:"TotalResourceCount,omitempty"`
}

func (s DescribeProductsResponseBodyDataContent) String() string {
	return tea.Prettify(s)
}

func (s DescribeProductsResponseBodyDataContent) GoString() string {
	return s.String()
}

func (s *DescribeProductsResponseBodyDataContent) SetCheckFailedCount(v int64) *DescribeProductsResponseBodyDataContent {
	s.CheckFailedCount = &v
	return s
}

func (s *DescribeProductsResponseBodyDataContent) SetCheckFailedResourceCount(v int64) *DescribeProductsResponseBodyDataContent {
	s.CheckFailedResourceCount = &v
	return s
}

func (s *DescribeProductsResponseBodyDataContent) SetDisableCheckResourceCount(v int64) *DescribeProductsResponseBodyDataContent {
	s.DisableCheckResourceCount = &v
	return s
}

func (s *DescribeProductsResponseBodyDataContent) SetEnableCheck(v bool) *DescribeProductsResponseBodyDataContent {
	s.EnableCheck = &v
	return s
}

func (s *DescribeProductsResponseBodyDataContent) SetProductType(v string) *DescribeProductsResponseBodyDataContent {
	s.ProductType = &v
	return s
}

func (s *DescribeProductsResponseBodyDataContent) SetProtectionScore(v int32) *DescribeProductsResponseBodyDataContent {
	s.ProtectionScore = &v
	return s
}

func (s *DescribeProductsResponseBodyDataContent) SetProtectionScoreDistribution(v []*DescribeProductsResponseBodyDataContentProtectionScoreDistribution) *DescribeProductsResponseBodyDataContent {
	s.ProtectionScoreDistribution = v
	return s
}

func (s *DescribeProductsResponseBodyDataContent) SetProtectionScoreUpdatedTime(v int64) *DescribeProductsResponseBodyDataContent {
	s.ProtectionScoreUpdatedTime = &v
	return s
}

func (s *DescribeProductsResponseBodyDataContent) SetRiskCount(v int64) *DescribeProductsResponseBodyDataContent {
	s.RiskCount = &v
	return s
}

func (s *DescribeProductsResponseBodyDataContent) SetRiskyResourceCount(v int64) *DescribeProductsResponseBodyDataContent {
	s.RiskyResourceCount = &v
	return s
}

func (s *DescribeProductsResponseBodyDataContent) SetTotalResourceCount(v int64) *DescribeProductsResponseBodyDataContent {
	s.TotalResourceCount = &v
	return s
}

type DescribeProductsResponseBodyDataContentProtectionScoreDistribution struct {
	// example:
	//
	// 5
	Count *int64                                                                   `json:"Count,omitempty" xml:"Count,omitempty"`
	Range *DescribeProductsResponseBodyDataContentProtectionScoreDistributionRange `json:"Range,omitempty" xml:"Range,omitempty" type:"Struct"`
}

func (s DescribeProductsResponseBodyDataContentProtectionScoreDistribution) String() string {
	return tea.Prettify(s)
}

func (s DescribeProductsResponseBodyDataContentProtectionScoreDistribution) GoString() string {
	return s.String()
}

func (s *DescribeProductsResponseBodyDataContentProtectionScoreDistribution) SetCount(v int64) *DescribeProductsResponseBodyDataContentProtectionScoreDistribution {
	s.Count = &v
	return s
}

func (s *DescribeProductsResponseBodyDataContentProtectionScoreDistribution) SetRange(v *DescribeProductsResponseBodyDataContentProtectionScoreDistributionRange) *DescribeProductsResponseBodyDataContentProtectionScoreDistribution {
	s.Range = v
	return s
}

type DescribeProductsResponseBodyDataContentProtectionScoreDistributionRange struct {
	// example:
	//
	// 0
	From *int32 `json:"From,omitempty" xml:"From,omitempty"`
	// example:
	//
	// 60
	To *int32 `json:"To,omitempty" xml:"To,omitempty"`
}

func (s DescribeProductsResponseBodyDataContentProtectionScoreDistributionRange) String() string {
	return tea.Prettify(s)
}

func (s DescribeProductsResponseBodyDataContentProtectionScoreDistributionRange) GoString() string {
	return s.String()
}

func (s *DescribeProductsResponseBodyDataContentProtectionScoreDistributionRange) SetFrom(v int32) *DescribeProductsResponseBodyDataContentProtectionScoreDistributionRange {
	s.From = &v
	return s
}

func (s *DescribeProductsResponseBodyDataContentProtectionScoreDistributionRange) SetTo(v int32) *DescribeProductsResponseBodyDataContentProtectionScoreDistributionRange {
	s.To = &v
	return s
}

type DescribeProductsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeProductsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeProductsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeProductsResponse) GoString() string {
	return s.String()
}

func (s *DescribeProductsResponse) SetHeaders(v map[string]*string) *DescribeProductsResponse {
	s.Headers = v
	return s
}

func (s *DescribeProductsResponse) SetStatusCode(v int32) *DescribeProductsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeProductsResponse) SetBody(v *DescribeProductsResponseBody) *DescribeProductsResponse {
	s.Body = v
	return s
}

type DescribeResourcesRequest struct {
	// example:
	//
	// rule-000c***yc9
	FailedRuleTemplate *string `json:"FailedRuleTemplate,omitempty" xml:"FailedRuleTemplate,omitempty"`
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// cae**********699
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// i-0003***110
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// example:
	//
	// ACS::ECS::Instance
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// example:
	//
	// protectionScore
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// example:
	//
	// ASC
	SortOrder *string `json:"SortOrder,omitempty" xml:"SortOrder,omitempty"`
}

func (s DescribeResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeResourcesRequest) GoString() string {
	return s.String()
}

func (s *DescribeResourcesRequest) SetFailedRuleTemplate(v string) *DescribeResourcesRequest {
	s.FailedRuleTemplate = &v
	return s
}

func (s *DescribeResourcesRequest) SetMaxResults(v int32) *DescribeResourcesRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeResourcesRequest) SetNextToken(v string) *DescribeResourcesRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeResourcesRequest) SetResourceId(v string) *DescribeResourcesRequest {
	s.ResourceId = &v
	return s
}

func (s *DescribeResourcesRequest) SetResourceType(v string) *DescribeResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *DescribeResourcesRequest) SetSortBy(v string) *DescribeResourcesRequest {
	s.SortBy = &v
	return s
}

func (s *DescribeResourcesRequest) SetSortOrder(v string) *DescribeResourcesRequest {
	s.SortOrder = &v
	return s
}

type DescribeResourcesResponseBody struct {
	Data *DescribeResourcesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 700683DE-0154-56D4-8D76-3B7A2C2C7DF9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeResourcesResponseBody) SetData(v *DescribeResourcesResponseBodyData) *DescribeResourcesResponseBody {
	s.Data = v
	return s
}

func (s *DescribeResourcesResponseBody) SetRequestId(v string) *DescribeResourcesResponseBody {
	s.RequestId = &v
	return s
}

type DescribeResourcesResponseBodyData struct {
	Content []*DescribeResourcesResponseBodyDataContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Repeated"`
	// example:
	//
	// 100
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// fb836242f4225fa0f0e0257362dfc6dd
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// 149
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeResourcesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeResourcesResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeResourcesResponseBodyData) SetContent(v []*DescribeResourcesResponseBodyDataContent) *DescribeResourcesResponseBodyData {
	s.Content = v
	return s
}

func (s *DescribeResourcesResponseBodyData) SetMaxResults(v int32) *DescribeResourcesResponseBodyData {
	s.MaxResults = &v
	return s
}

func (s *DescribeResourcesResponseBodyData) SetNextToken(v string) *DescribeResourcesResponseBodyData {
	s.NextToken = &v
	return s
}

func (s *DescribeResourcesResponseBodyData) SetTotalCount(v int64) *DescribeResourcesResponseBodyData {
	s.TotalCount = &v
	return s
}

type DescribeResourcesResponseBodyDataContent struct {
	// example:
	//
	// 0
	ArchiveDataSize *int64 `json:"ArchiveDataSize,omitempty" xml:"ArchiveDataSize,omitempty"`
	// example:
	//
	// 0
	CheckFailedCount *int64 `json:"CheckFailedCount,omitempty" xml:"CheckFailedCount,omitempty"`
	// example:
	//
	// 0
	ColdArchiveDataSize *int64 `json:"ColdArchiveDataSize,omitempty" xml:"ColdArchiveDataSize,omitempty"`
	// example:
	//
	// 1697798340
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 0
	EnableCheck *bool `json:"EnableCheck,omitempty" xml:"EnableCheck,omitempty"`
	// example:
	//
	// 0
	IaDataSize *int64 `json:"IaDataSize,omitempty" xml:"IaDataSize,omitempty"`
	// example:
	//
	// ecs
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	// example:
	//
	// 0
	ProtectionScore *int32 `json:"ProtectionScore,omitempty" xml:"ProtectionScore,omitempty"`
	// example:
	//
	// 0
	ProtectionScoreUpdatedTime *int64 `json:"ProtectionScoreUpdatedTime,omitempty" xml:"ProtectionScoreUpdatedTime,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// acs:ecs:cn-hangzhou:xxxxxxxx:instance/xxxxx
	ResourceArn *string `json:"ResourceArn,omitempty" xml:"ResourceArn,omitempty"`
	// example:
	//
	// i-xxxxxxxx
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// example:
	//
	// test server
	ResourceName *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	// example:
	//
	// ACS::ECS::Instance
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// example:
	//
	// 0
	RiskCount *int64 `json:"RiskCount,omitempty" xml:"RiskCount,omitempty"`
	// example:
	//
	// 0
	StandardDataSize *int64 `json:"StandardDataSize,omitempty" xml:"StandardDataSize,omitempty"`
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 0
	TotalDataSize *int64 `json:"TotalDataSize,omitempty" xml:"TotalDataSize,omitempty"`
	// vSwitch ID
	//
	// example:
	//
	// vsw-xxxxxxxx
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// vpc ID
	//
	// example:
	//
	// vpc-xxxxxxxx
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// example:
	//
	// cn-hangzhou-i
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeResourcesResponseBodyDataContent) String() string {
	return tea.Prettify(s)
}

func (s DescribeResourcesResponseBodyDataContent) GoString() string {
	return s.String()
}

func (s *DescribeResourcesResponseBodyDataContent) SetArchiveDataSize(v int64) *DescribeResourcesResponseBodyDataContent {
	s.ArchiveDataSize = &v
	return s
}

func (s *DescribeResourcesResponseBodyDataContent) SetCheckFailedCount(v int64) *DescribeResourcesResponseBodyDataContent {
	s.CheckFailedCount = &v
	return s
}

func (s *DescribeResourcesResponseBodyDataContent) SetColdArchiveDataSize(v int64) *DescribeResourcesResponseBodyDataContent {
	s.ColdArchiveDataSize = &v
	return s
}

func (s *DescribeResourcesResponseBodyDataContent) SetCreateTime(v int64) *DescribeResourcesResponseBodyDataContent {
	s.CreateTime = &v
	return s
}

func (s *DescribeResourcesResponseBodyDataContent) SetEnableCheck(v bool) *DescribeResourcesResponseBodyDataContent {
	s.EnableCheck = &v
	return s
}

func (s *DescribeResourcesResponseBodyDataContent) SetIaDataSize(v int64) *DescribeResourcesResponseBodyDataContent {
	s.IaDataSize = &v
	return s
}

func (s *DescribeResourcesResponseBodyDataContent) SetProductType(v string) *DescribeResourcesResponseBodyDataContent {
	s.ProductType = &v
	return s
}

func (s *DescribeResourcesResponseBodyDataContent) SetProtectionScore(v int32) *DescribeResourcesResponseBodyDataContent {
	s.ProtectionScore = &v
	return s
}

func (s *DescribeResourcesResponseBodyDataContent) SetProtectionScoreUpdatedTime(v int64) *DescribeResourcesResponseBodyDataContent {
	s.ProtectionScoreUpdatedTime = &v
	return s
}

func (s *DescribeResourcesResponseBodyDataContent) SetRegionId(v string) *DescribeResourcesResponseBodyDataContent {
	s.RegionId = &v
	return s
}

func (s *DescribeResourcesResponseBodyDataContent) SetResourceArn(v string) *DescribeResourcesResponseBodyDataContent {
	s.ResourceArn = &v
	return s
}

func (s *DescribeResourcesResponseBodyDataContent) SetResourceId(v string) *DescribeResourcesResponseBodyDataContent {
	s.ResourceId = &v
	return s
}

func (s *DescribeResourcesResponseBodyDataContent) SetResourceName(v string) *DescribeResourcesResponseBodyDataContent {
	s.ResourceName = &v
	return s
}

func (s *DescribeResourcesResponseBodyDataContent) SetResourceType(v string) *DescribeResourcesResponseBodyDataContent {
	s.ResourceType = &v
	return s
}

func (s *DescribeResourcesResponseBodyDataContent) SetRiskCount(v int64) *DescribeResourcesResponseBodyDataContent {
	s.RiskCount = &v
	return s
}

func (s *DescribeResourcesResponseBodyDataContent) SetStandardDataSize(v int64) *DescribeResourcesResponseBodyDataContent {
	s.StandardDataSize = &v
	return s
}

func (s *DescribeResourcesResponseBodyDataContent) SetStatus(v string) *DescribeResourcesResponseBodyDataContent {
	s.Status = &v
	return s
}

func (s *DescribeResourcesResponseBodyDataContent) SetTotalDataSize(v int64) *DescribeResourcesResponseBodyDataContent {
	s.TotalDataSize = &v
	return s
}

func (s *DescribeResourcesResponseBodyDataContent) SetVSwitchId(v string) *DescribeResourcesResponseBodyDataContent {
	s.VSwitchId = &v
	return s
}

func (s *DescribeResourcesResponseBodyDataContent) SetVpcId(v string) *DescribeResourcesResponseBodyDataContent {
	s.VpcId = &v
	return s
}

func (s *DescribeResourcesResponseBodyDataContent) SetZoneId(v string) *DescribeResourcesResponseBodyDataContent {
	s.ZoneId = &v
	return s
}

type DescribeResourcesResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeResourcesResponse) GoString() string {
	return s.String()
}

func (s *DescribeResourcesResponse) SetHeaders(v map[string]*string) *DescribeResourcesResponse {
	s.Headers = v
	return s
}

func (s *DescribeResourcesResponse) SetStatusCode(v int32) *DescribeResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeResourcesResponse) SetBody(v *DescribeResourcesResponseBody) *DescribeResourcesResponse {
	s.Body = v
	return s
}

type DescribeRulesRequest struct {
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// cae**********699
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// ACS::ECS::Instance
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s DescribeRulesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRulesRequest) GoString() string {
	return s.String()
}

func (s *DescribeRulesRequest) SetMaxResults(v int32) *DescribeRulesRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeRulesRequest) SetNextToken(v string) *DescribeRulesRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeRulesRequest) SetResourceType(v string) *DescribeRulesRequest {
	s.ResourceType = &v
	return s
}

type DescribeRulesResponseBody struct {
	Data *DescribeRulesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 86DEBAC9-AB6A-59AB-9E5C-A540E579ECC9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeRulesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRulesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRulesResponseBody) SetData(v *DescribeRulesResponseBodyData) *DescribeRulesResponseBody {
	s.Data = v
	return s
}

func (s *DescribeRulesResponseBody) SetRequestId(v string) *DescribeRulesResponseBody {
	s.RequestId = &v
	return s
}

type DescribeRulesResponseBodyData struct {
	Content []*DescribeRulesResponseBodyDataContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Repeated"`
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// 0975951c75d7b41464c8d08ae17043ca
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// 42
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeRulesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeRulesResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeRulesResponseBodyData) SetContent(v []*DescribeRulesResponseBodyDataContent) *DescribeRulesResponseBodyData {
	s.Content = v
	return s
}

func (s *DescribeRulesResponseBodyData) SetMaxResults(v int32) *DescribeRulesResponseBodyData {
	s.MaxResults = &v
	return s
}

func (s *DescribeRulesResponseBodyData) SetNextToken(v string) *DescribeRulesResponseBodyData {
	s.NextToken = &v
	return s
}

func (s *DescribeRulesResponseBodyData) SetTotalCount(v int64) *DescribeRulesResponseBodyData {
	s.TotalCount = &v
	return s
}

type DescribeRulesResponseBodyDataContent struct {
	// example:
	//
	// 0
	CheckFailedResourceCount *int64 `json:"CheckFailedResourceCount,omitempty" xml:"CheckFailedResourceCount,omitempty"`
	// example:
	//
	// PASSED
	CheckStatus *string `json:"CheckStatus,omitempty" xml:"CheckStatus,omitempty"`
	// example:
	//
	// 1704157635
	CheckTime *int64 `json:"CheckTime,omitempty" xml:"CheckTime,omitempty"`
	// example:
	//
	// ecs
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	// example:
	//
	// ACS::ECS::Instance
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// example:
	//
	// 0
	RiskyResourceCount *int64 `json:"RiskyResourceCount,omitempty" xml:"RiskyResourceCount,omitempty"`
	// example:
	//
	// rule-bp11ggd8wr762
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// example:
	//
	// ecs-backup
	RuleTemplate *string `json:"RuleTemplate,omitempty" xml:"RuleTemplate,omitempty"`
	// example:
	//
	// 1
	TotalResourceCount *int64 `json:"TotalResourceCount,omitempty" xml:"TotalResourceCount,omitempty"`
}

func (s DescribeRulesResponseBodyDataContent) String() string {
	return tea.Prettify(s)
}

func (s DescribeRulesResponseBodyDataContent) GoString() string {
	return s.String()
}

func (s *DescribeRulesResponseBodyDataContent) SetCheckFailedResourceCount(v int64) *DescribeRulesResponseBodyDataContent {
	s.CheckFailedResourceCount = &v
	return s
}

func (s *DescribeRulesResponseBodyDataContent) SetCheckStatus(v string) *DescribeRulesResponseBodyDataContent {
	s.CheckStatus = &v
	return s
}

func (s *DescribeRulesResponseBodyDataContent) SetCheckTime(v int64) *DescribeRulesResponseBodyDataContent {
	s.CheckTime = &v
	return s
}

func (s *DescribeRulesResponseBodyDataContent) SetProductType(v string) *DescribeRulesResponseBodyDataContent {
	s.ProductType = &v
	return s
}

func (s *DescribeRulesResponseBodyDataContent) SetResourceType(v string) *DescribeRulesResponseBodyDataContent {
	s.ResourceType = &v
	return s
}

func (s *DescribeRulesResponseBodyDataContent) SetRiskyResourceCount(v int64) *DescribeRulesResponseBodyDataContent {
	s.RiskyResourceCount = &v
	return s
}

func (s *DescribeRulesResponseBodyDataContent) SetRuleId(v string) *DescribeRulesResponseBodyDataContent {
	s.RuleId = &v
	return s
}

func (s *DescribeRulesResponseBodyDataContent) SetRuleTemplate(v string) *DescribeRulesResponseBodyDataContent {
	s.RuleTemplate = &v
	return s
}

func (s *DescribeRulesResponseBodyDataContent) SetTotalResourceCount(v int64) *DescribeRulesResponseBodyDataContent {
	s.TotalResourceCount = &v
	return s
}

type DescribeRulesResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRulesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRulesResponse) GoString() string {
	return s.String()
}

func (s *DescribeRulesResponse) SetHeaders(v map[string]*string) *DescribeRulesResponse {
	s.Headers = v
	return s
}

func (s *DescribeRulesResponse) SetStatusCode(v int32) *DescribeRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRulesResponse) SetBody(v *DescribeRulesResponseBody) *DescribeRulesResponse {
	s.Body = v
	return s
}

type DescribeTaskResponseBody struct {
	Data *DescribeTaskResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 14DFF801-A4E3-5136-AAB8-7D246012CD7A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeTaskResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTaskResponseBody) SetData(v *DescribeTaskResponseBodyData) *DescribeTaskResponseBody {
	s.Data = v
	return s
}

func (s *DescribeTaskResponseBody) SetRequestId(v string) *DescribeTaskResponseBody {
	s.RequestId = &v
	return s
}

type DescribeTaskResponseBodyData struct {
	// example:
	//
	// 1724983927
	CompleteTime *int64 `json:"CompleteTime,omitempty" xml:"CompleteTime,omitempty"`
	// example:
	//
	// too many requests.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	ExecutionId  *string `json:"ExecutionId,omitempty" xml:"ExecutionId,omitempty"`
	// example:
	//
	// 1719026680
	ExpireTime *int64 `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// example:
	//
	// 100
	Progress *int32 `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// example:
	//
	// 784076D6-BD6D-5564-9CEA-834EB11F0C62
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 1724983927
	StartTime       *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	TaskDescription *string `json:"TaskDescription,omitempty" xml:"TaskDescription,omitempty"`
	// example:
	//
	// {"resourceTypes":["ACS::ECS::Instance","ACS::OSS::Bucket","ACS::OTS::Instance","ACS::NAS::FileSystem"]}
	TaskDetail *string `json:"TaskDetail,omitempty" xml:"TaskDetail,omitempty"`
	// example:
	//
	// t-xxxxxxxx
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// example:
	//
	// test5566
	TaskName     *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	TaskPriority *string `json:"TaskPriority,omitempty" xml:"TaskPriority,omitempty"`
	// example:
	//
	// RUNNING
	TaskStatus *string `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	// example:
	//
	// UPDATE_RESOURCES
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s DescribeTaskResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeTaskResponseBodyData) SetCompleteTime(v int64) *DescribeTaskResponseBodyData {
	s.CompleteTime = &v
	return s
}

func (s *DescribeTaskResponseBodyData) SetErrorMessage(v string) *DescribeTaskResponseBodyData {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeTaskResponseBodyData) SetExecutionId(v string) *DescribeTaskResponseBodyData {
	s.ExecutionId = &v
	return s
}

func (s *DescribeTaskResponseBodyData) SetExpireTime(v int64) *DescribeTaskResponseBodyData {
	s.ExpireTime = &v
	return s
}

func (s *DescribeTaskResponseBodyData) SetProgress(v int32) *DescribeTaskResponseBodyData {
	s.Progress = &v
	return s
}

func (s *DescribeTaskResponseBodyData) SetRequestId(v string) *DescribeTaskResponseBodyData {
	s.RequestId = &v
	return s
}

func (s *DescribeTaskResponseBodyData) SetStartTime(v int64) *DescribeTaskResponseBodyData {
	s.StartTime = &v
	return s
}

func (s *DescribeTaskResponseBodyData) SetTaskDescription(v string) *DescribeTaskResponseBodyData {
	s.TaskDescription = &v
	return s
}

func (s *DescribeTaskResponseBodyData) SetTaskDetail(v string) *DescribeTaskResponseBodyData {
	s.TaskDetail = &v
	return s
}

func (s *DescribeTaskResponseBodyData) SetTaskId(v string) *DescribeTaskResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *DescribeTaskResponseBodyData) SetTaskName(v string) *DescribeTaskResponseBodyData {
	s.TaskName = &v
	return s
}

func (s *DescribeTaskResponseBodyData) SetTaskPriority(v string) *DescribeTaskResponseBodyData {
	s.TaskPriority = &v
	return s
}

func (s *DescribeTaskResponseBodyData) SetTaskStatus(v string) *DescribeTaskResponseBodyData {
	s.TaskStatus = &v
	return s
}

func (s *DescribeTaskResponseBodyData) SetTaskType(v string) *DescribeTaskResponseBodyData {
	s.TaskType = &v
	return s
}

type DescribeTaskResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeTaskResponse) GoString() string {
	return s.String()
}

func (s *DescribeTaskResponse) SetHeaders(v map[string]*string) *DescribeTaskResponse {
	s.Headers = v
	return s
}

func (s *DescribeTaskResponse) SetStatusCode(v int32) *DescribeTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeTaskResponse) SetBody(v *DescribeTaskResponseBody) *DescribeTaskResponse {
	s.Body = v
	return s
}

type DescribeTasksRequest struct {
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// cae**********699
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// RUNNING
	TaskStatus *string `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
}

func (s DescribeTasksRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeTasksRequest) GoString() string {
	return s.String()
}

func (s *DescribeTasksRequest) SetMaxResults(v int32) *DescribeTasksRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeTasksRequest) SetNextToken(v string) *DescribeTasksRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeTasksRequest) SetTaskStatus(v string) *DescribeTasksRequest {
	s.TaskStatus = &v
	return s
}

type DescribeTasksResponseBody struct {
	Data *DescribeTasksResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// AE43C4CB-8074-5EBD-9806-8CA6D12800B1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeTasksResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeTasksResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTasksResponseBody) SetData(v *DescribeTasksResponseBodyData) *DescribeTasksResponseBody {
	s.Data = v
	return s
}

func (s *DescribeTasksResponseBody) SetRequestId(v string) *DescribeTasksResponseBody {
	s.RequestId = &v
	return s
}

type DescribeTasksResponseBodyData struct {
	Content []*DescribeTasksResponseBodyDataContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Repeated"`
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// f4b8c2504545a3b41af5e75147d17d12e3818a0b9b2ff9a2
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// 100
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeTasksResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeTasksResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeTasksResponseBodyData) SetContent(v []*DescribeTasksResponseBodyDataContent) *DescribeTasksResponseBodyData {
	s.Content = v
	return s
}

func (s *DescribeTasksResponseBodyData) SetMaxResults(v int32) *DescribeTasksResponseBodyData {
	s.MaxResults = &v
	return s
}

func (s *DescribeTasksResponseBodyData) SetNextToken(v string) *DescribeTasksResponseBodyData {
	s.NextToken = &v
	return s
}

func (s *DescribeTasksResponseBodyData) SetTotalCount(v int64) *DescribeTasksResponseBodyData {
	s.TotalCount = &v
	return s
}

type DescribeTasksResponseBodyDataContent struct {
	// example:
	//
	// 1724983927
	CompleteTime *int64 `json:"CompleteTime,omitempty" xml:"CompleteTime,omitempty"`
	// example:
	//
	// device not online
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// example:
	//
	// empty
	ExecutionId *string `json:"ExecutionId,omitempty" xml:"ExecutionId,omitempty"`
	// example:
	//
	// 1724983927
	ExpireTime *int64 `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// example:
	//
	// 100
	Progress *int32 `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// example:
	//
	// AE43C4CB-8074-5EBD-9806-8CA6D12800B1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 1724983927
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// empty
	TaskDescription *string `json:"TaskDescription,omitempty" xml:"TaskDescription,omitempty"`
	// example:
	//
	// {"resourceTypes":["ACS::ECS::Instance","ACS::OSS::Bucket","ACS::OTS::Instance","ACS::NAS::FileSystem"]}
	TaskDetail *string `json:"TaskDetail,omitempty" xml:"TaskDetail,omitempty"`
	// example:
	//
	// t-0000e4w0u1v592zdf6s7
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// example:
	//
	// empty
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	// example:
	//
	// HIGH
	TaskPriority *string `json:"TaskPriority,omitempty" xml:"TaskPriority,omitempty"`
	// example:
	//
	// RUNNING
	TaskStatus *string `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	// example:
	//
	// UPDATE_RESOURCES
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s DescribeTasksResponseBodyDataContent) String() string {
	return tea.Prettify(s)
}

func (s DescribeTasksResponseBodyDataContent) GoString() string {
	return s.String()
}

func (s *DescribeTasksResponseBodyDataContent) SetCompleteTime(v int64) *DescribeTasksResponseBodyDataContent {
	s.CompleteTime = &v
	return s
}

func (s *DescribeTasksResponseBodyDataContent) SetErrorMessage(v string) *DescribeTasksResponseBodyDataContent {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeTasksResponseBodyDataContent) SetExecutionId(v string) *DescribeTasksResponseBodyDataContent {
	s.ExecutionId = &v
	return s
}

func (s *DescribeTasksResponseBodyDataContent) SetExpireTime(v int64) *DescribeTasksResponseBodyDataContent {
	s.ExpireTime = &v
	return s
}

func (s *DescribeTasksResponseBodyDataContent) SetProgress(v int32) *DescribeTasksResponseBodyDataContent {
	s.Progress = &v
	return s
}

func (s *DescribeTasksResponseBodyDataContent) SetRequestId(v string) *DescribeTasksResponseBodyDataContent {
	s.RequestId = &v
	return s
}

func (s *DescribeTasksResponseBodyDataContent) SetStartTime(v int64) *DescribeTasksResponseBodyDataContent {
	s.StartTime = &v
	return s
}

func (s *DescribeTasksResponseBodyDataContent) SetTaskDescription(v string) *DescribeTasksResponseBodyDataContent {
	s.TaskDescription = &v
	return s
}

func (s *DescribeTasksResponseBodyDataContent) SetTaskDetail(v string) *DescribeTasksResponseBodyDataContent {
	s.TaskDetail = &v
	return s
}

func (s *DescribeTasksResponseBodyDataContent) SetTaskId(v string) *DescribeTasksResponseBodyDataContent {
	s.TaskId = &v
	return s
}

func (s *DescribeTasksResponseBodyDataContent) SetTaskName(v string) *DescribeTasksResponseBodyDataContent {
	s.TaskName = &v
	return s
}

func (s *DescribeTasksResponseBodyDataContent) SetTaskPriority(v string) *DescribeTasksResponseBodyDataContent {
	s.TaskPriority = &v
	return s
}

func (s *DescribeTasksResponseBodyDataContent) SetTaskStatus(v string) *DescribeTasksResponseBodyDataContent {
	s.TaskStatus = &v
	return s
}

func (s *DescribeTasksResponseBodyDataContent) SetTaskType(v string) *DescribeTasksResponseBodyDataContent {
	s.TaskType = &v
	return s
}

type DescribeTasksResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeTasksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeTasksResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeTasksResponse) GoString() string {
	return s.String()
}

func (s *DescribeTasksResponse) SetHeaders(v map[string]*string) *DescribeTasksResponse {
	s.Headers = v
	return s
}

func (s *DescribeTasksResponse) SetStatusCode(v int32) *DescribeTasksResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeTasksResponse) SetBody(v *DescribeTasksResponseBody) *DescribeTasksResponse {
	s.Body = v
	return s
}

type DescribeTopRiskyResourcesRequest struct {
	// example:
	//
	// ACS::ECS::Instance
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s DescribeTopRiskyResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeTopRiskyResourcesRequest) GoString() string {
	return s.String()
}

func (s *DescribeTopRiskyResourcesRequest) SetResourceType(v string) *DescribeTopRiskyResourcesRequest {
	s.ResourceType = &v
	return s
}

type DescribeTopRiskyResourcesResponseBody struct {
	Data *DescribeTopRiskyResourcesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 34081B20-C4C0-514F-93F6-8EEC3D1A587E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeTopRiskyResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeTopRiskyResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTopRiskyResourcesResponseBody) SetData(v *DescribeTopRiskyResourcesResponseBodyData) *DescribeTopRiskyResourcesResponseBody {
	s.Data = v
	return s
}

func (s *DescribeTopRiskyResourcesResponseBody) SetRequestId(v string) *DescribeTopRiskyResourcesResponseBody {
	s.RequestId = &v
	return s
}

type DescribeTopRiskyResourcesResponseBodyData struct {
	Content []*DescribeTopRiskyResourcesResponseBodyDataContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Repeated"`
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// e557bc9a65fe22cb5e2a3b240f06b0de
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// 10
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeTopRiskyResourcesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeTopRiskyResourcesResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeTopRiskyResourcesResponseBodyData) SetContent(v []*DescribeTopRiskyResourcesResponseBodyDataContent) *DescribeTopRiskyResourcesResponseBodyData {
	s.Content = v
	return s
}

func (s *DescribeTopRiskyResourcesResponseBodyData) SetMaxResults(v int32) *DescribeTopRiskyResourcesResponseBodyData {
	s.MaxResults = &v
	return s
}

func (s *DescribeTopRiskyResourcesResponseBodyData) SetNextToken(v string) *DescribeTopRiskyResourcesResponseBodyData {
	s.NextToken = &v
	return s
}

func (s *DescribeTopRiskyResourcesResponseBodyData) SetTotalCount(v int64) *DescribeTopRiskyResourcesResponseBodyData {
	s.TotalCount = &v
	return s
}

type DescribeTopRiskyResourcesResponseBodyDataContent struct {
	// example:
	//
	// 0
	ArchiveDataSize *int64 `json:"ArchiveDataSize,omitempty" xml:"ArchiveDataSize,omitempty"`
	// example:
	//
	// 0
	CheckFailedCount *int64 `json:"CheckFailedCount,omitempty" xml:"CheckFailedCount,omitempty"`
	// example:
	//
	// 0
	ColdArchiveDataSize *int64 `json:"ColdArchiveDataSize,omitempty" xml:"ColdArchiveDataSize,omitempty"`
	// example:
	//
	// 1697798340
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// true
	EnableCheck *bool `json:"EnableCheck,omitempty" xml:"EnableCheck,omitempty"`
	// example:
	//
	// 0
	IaDataSize *int64 `json:"IaDataSize,omitempty" xml:"IaDataSize,omitempty"`
	// example:
	//
	// ecs
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	// example:
	//
	// 90
	ProtectionScore *int32 `json:"ProtectionScore,omitempty" xml:"ProtectionScore,omitempty"`
	// example:
	//
	// 1726036498
	ProtectionScoreUpdatedTime *int64 `json:"ProtectionScoreUpdatedTime,omitempty" xml:"ProtectionScoreUpdatedTime,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// acs:ecs:cn-hangzhou:xxxxxxxx:instance/xxxxx
	ResourceArn *string `json:"ResourceArn,omitempty" xml:"ResourceArn,omitempty"`
	// example:
	//
	// i-xxxxxxxx
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// example:
	//
	// test-server
	ResourceName *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	// example:
	//
	// ACS::ECS::Instance
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// example:
	//
	// 0
	RiskCount *int64 `json:"RiskCount,omitempty" xml:"RiskCount,omitempty"`
	// example:
	//
	// 0
	StandardDataSize *int64 `json:"StandardDataSize,omitempty" xml:"StandardDataSize,omitempty"`
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 0
	TotalDataSize *int64 `json:"TotalDataSize,omitempty" xml:"TotalDataSize,omitempty"`
	// vSwitch ID
	//
	// example:
	//
	// vsw-xxxxxxxx
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// vpc ID
	//
	// example:
	//
	// vpc-xxxxxxxx
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// example:
	//
	// cn-hangzhou-j
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeTopRiskyResourcesResponseBodyDataContent) String() string {
	return tea.Prettify(s)
}

func (s DescribeTopRiskyResourcesResponseBodyDataContent) GoString() string {
	return s.String()
}

func (s *DescribeTopRiskyResourcesResponseBodyDataContent) SetArchiveDataSize(v int64) *DescribeTopRiskyResourcesResponseBodyDataContent {
	s.ArchiveDataSize = &v
	return s
}

func (s *DescribeTopRiskyResourcesResponseBodyDataContent) SetCheckFailedCount(v int64) *DescribeTopRiskyResourcesResponseBodyDataContent {
	s.CheckFailedCount = &v
	return s
}

func (s *DescribeTopRiskyResourcesResponseBodyDataContent) SetColdArchiveDataSize(v int64) *DescribeTopRiskyResourcesResponseBodyDataContent {
	s.ColdArchiveDataSize = &v
	return s
}

func (s *DescribeTopRiskyResourcesResponseBodyDataContent) SetCreateTime(v int64) *DescribeTopRiskyResourcesResponseBodyDataContent {
	s.CreateTime = &v
	return s
}

func (s *DescribeTopRiskyResourcesResponseBodyDataContent) SetEnableCheck(v bool) *DescribeTopRiskyResourcesResponseBodyDataContent {
	s.EnableCheck = &v
	return s
}

func (s *DescribeTopRiskyResourcesResponseBodyDataContent) SetIaDataSize(v int64) *DescribeTopRiskyResourcesResponseBodyDataContent {
	s.IaDataSize = &v
	return s
}

func (s *DescribeTopRiskyResourcesResponseBodyDataContent) SetProductType(v string) *DescribeTopRiskyResourcesResponseBodyDataContent {
	s.ProductType = &v
	return s
}

func (s *DescribeTopRiskyResourcesResponseBodyDataContent) SetProtectionScore(v int32) *DescribeTopRiskyResourcesResponseBodyDataContent {
	s.ProtectionScore = &v
	return s
}

func (s *DescribeTopRiskyResourcesResponseBodyDataContent) SetProtectionScoreUpdatedTime(v int64) *DescribeTopRiskyResourcesResponseBodyDataContent {
	s.ProtectionScoreUpdatedTime = &v
	return s
}

func (s *DescribeTopRiskyResourcesResponseBodyDataContent) SetRegionId(v string) *DescribeTopRiskyResourcesResponseBodyDataContent {
	s.RegionId = &v
	return s
}

func (s *DescribeTopRiskyResourcesResponseBodyDataContent) SetResourceArn(v string) *DescribeTopRiskyResourcesResponseBodyDataContent {
	s.ResourceArn = &v
	return s
}

func (s *DescribeTopRiskyResourcesResponseBodyDataContent) SetResourceId(v string) *DescribeTopRiskyResourcesResponseBodyDataContent {
	s.ResourceId = &v
	return s
}

func (s *DescribeTopRiskyResourcesResponseBodyDataContent) SetResourceName(v string) *DescribeTopRiskyResourcesResponseBodyDataContent {
	s.ResourceName = &v
	return s
}

func (s *DescribeTopRiskyResourcesResponseBodyDataContent) SetResourceType(v string) *DescribeTopRiskyResourcesResponseBodyDataContent {
	s.ResourceType = &v
	return s
}

func (s *DescribeTopRiskyResourcesResponseBodyDataContent) SetRiskCount(v int64) *DescribeTopRiskyResourcesResponseBodyDataContent {
	s.RiskCount = &v
	return s
}

func (s *DescribeTopRiskyResourcesResponseBodyDataContent) SetStandardDataSize(v int64) *DescribeTopRiskyResourcesResponseBodyDataContent {
	s.StandardDataSize = &v
	return s
}

func (s *DescribeTopRiskyResourcesResponseBodyDataContent) SetStatus(v string) *DescribeTopRiskyResourcesResponseBodyDataContent {
	s.Status = &v
	return s
}

func (s *DescribeTopRiskyResourcesResponseBodyDataContent) SetTotalDataSize(v int64) *DescribeTopRiskyResourcesResponseBodyDataContent {
	s.TotalDataSize = &v
	return s
}

func (s *DescribeTopRiskyResourcesResponseBodyDataContent) SetVSwitchId(v string) *DescribeTopRiskyResourcesResponseBodyDataContent {
	s.VSwitchId = &v
	return s
}

func (s *DescribeTopRiskyResourcesResponseBodyDataContent) SetVpcId(v string) *DescribeTopRiskyResourcesResponseBodyDataContent {
	s.VpcId = &v
	return s
}

func (s *DescribeTopRiskyResourcesResponseBodyDataContent) SetZoneId(v string) *DescribeTopRiskyResourcesResponseBodyDataContent {
	s.ZoneId = &v
	return s
}

type DescribeTopRiskyResourcesResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeTopRiskyResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeTopRiskyResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeTopRiskyResourcesResponse) GoString() string {
	return s.String()
}

func (s *DescribeTopRiskyResourcesResponse) SetHeaders(v map[string]*string) *DescribeTopRiskyResourcesResponse {
	s.Headers = v
	return s
}

func (s *DescribeTopRiskyResourcesResponse) SetStatusCode(v int32) *DescribeTopRiskyResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeTopRiskyResourcesResponse) SetBody(v *DescribeTopRiskyResourcesResponseBody) *DescribeTopRiskyResourcesResponse {
	s.Body = v
	return s
}

type DisableCheckProductRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ecs
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
}

func (s DisableCheckProductRequest) String() string {
	return tea.Prettify(s)
}

func (s DisableCheckProductRequest) GoString() string {
	return s.String()
}

func (s *DisableCheckProductRequest) SetProductType(v string) *DisableCheckProductRequest {
	s.ProductType = &v
	return s
}

type DisableCheckProductResponseBody struct {
	// example:
	//
	// 89E3CBB7-16F3-52AE-BD32-31A43A2A807F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DisableCheckProductResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DisableCheckProductResponseBody) GoString() string {
	return s.String()
}

func (s *DisableCheckProductResponseBody) SetRequestId(v string) *DisableCheckProductResponseBody {
	s.RequestId = &v
	return s
}

type DisableCheckProductResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisableCheckProductResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisableCheckProductResponse) String() string {
	return tea.Prettify(s)
}

func (s DisableCheckProductResponse) GoString() string {
	return s.String()
}

func (s *DisableCheckProductResponse) SetHeaders(v map[string]*string) *DisableCheckProductResponse {
	s.Headers = v
	return s
}

func (s *DisableCheckProductResponse) SetStatusCode(v int32) *DisableCheckProductResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableCheckProductResponse) SetBody(v *DisableCheckProductResponseBody) *DisableCheckProductResponse {
	s.Body = v
	return s
}

type DisableCheckResourceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// acs:ecs:123***890:cn-shanghai:instance/i-001***90
	ResourceArn *string `json:"ResourceArn,omitempty" xml:"ResourceArn,omitempty"`
}

func (s DisableCheckResourceRequest) String() string {
	return tea.Prettify(s)
}

func (s DisableCheckResourceRequest) GoString() string {
	return s.String()
}

func (s *DisableCheckResourceRequest) SetResourceArn(v string) *DisableCheckResourceRequest {
	s.ResourceArn = &v
	return s
}

type DisableCheckResourceResponseBody struct {
	// example:
	//
	// 86DEBAC9-AB6A-59AB-9E5C-A540E579ECC9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DisableCheckResourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DisableCheckResourceResponseBody) GoString() string {
	return s.String()
}

func (s *DisableCheckResourceResponseBody) SetRequestId(v string) *DisableCheckResourceResponseBody {
	s.RequestId = &v
	return s
}

type DisableCheckResourceResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisableCheckResourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisableCheckResourceResponse) String() string {
	return tea.Prettify(s)
}

func (s DisableCheckResourceResponse) GoString() string {
	return s.String()
}

func (s *DisableCheckResourceResponse) SetHeaders(v map[string]*string) *DisableCheckResourceResponse {
	s.Headers = v
	return s
}

func (s *DisableCheckResourceResponse) SetStatusCode(v int32) *DisableCheckResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableCheckResourceResponse) SetBody(v *DisableCheckResourceResponseBody) *DisableCheckResourceResponse {
	s.Body = v
	return s
}

type EnableCheckProductRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ecs
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
}

func (s EnableCheckProductRequest) String() string {
	return tea.Prettify(s)
}

func (s EnableCheckProductRequest) GoString() string {
	return s.String()
}

func (s *EnableCheckProductRequest) SetProductType(v string) *EnableCheckProductRequest {
	s.ProductType = &v
	return s
}

type EnableCheckProductResponseBody struct {
	// example:
	//
	// 8724BC18-904D-5A0D-BFF4-F0554F0037E7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EnableCheckProductResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EnableCheckProductResponseBody) GoString() string {
	return s.String()
}

func (s *EnableCheckProductResponseBody) SetRequestId(v string) *EnableCheckProductResponseBody {
	s.RequestId = &v
	return s
}

type EnableCheckProductResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *EnableCheckProductResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableCheckProductResponse) String() string {
	return tea.Prettify(s)
}

func (s EnableCheckProductResponse) GoString() string {
	return s.String()
}

func (s *EnableCheckProductResponse) SetHeaders(v map[string]*string) *EnableCheckProductResponse {
	s.Headers = v
	return s
}

func (s *EnableCheckProductResponse) SetStatusCode(v int32) *EnableCheckProductResponse {
	s.StatusCode = &v
	return s
}

func (s *EnableCheckProductResponse) SetBody(v *EnableCheckProductResponseBody) *EnableCheckProductResponse {
	s.Body = v
	return s
}

type EnableCheckResourceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// acs:ecs:123***890:cn-shanghai:instance/i-001***90
	ResourceArn *string `json:"ResourceArn,omitempty" xml:"ResourceArn,omitempty"`
}

func (s EnableCheckResourceRequest) String() string {
	return tea.Prettify(s)
}

func (s EnableCheckResourceRequest) GoString() string {
	return s.String()
}

func (s *EnableCheckResourceRequest) SetResourceArn(v string) *EnableCheckResourceRequest {
	s.ResourceArn = &v
	return s
}

type EnableCheckResourceResponseBody struct {
	// example:
	//
	// E583A0FF-803C-51C4-9AC9-E029471ACD6A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EnableCheckResourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EnableCheckResourceResponseBody) GoString() string {
	return s.String()
}

func (s *EnableCheckResourceResponseBody) SetRequestId(v string) *EnableCheckResourceResponseBody {
	s.RequestId = &v
	return s
}

type EnableCheckResourceResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *EnableCheckResourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s EnableCheckResourceResponse) String() string {
	return tea.Prettify(s)
}

func (s EnableCheckResourceResponse) GoString() string {
	return s.String()
}

func (s *EnableCheckResourceResponse) SetHeaders(v map[string]*string) *EnableCheckResourceResponse {
	s.Headers = v
	return s
}

func (s *EnableCheckResourceResponse) SetStatusCode(v int32) *EnableCheckResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *EnableCheckResourceResponse) SetBody(v *EnableCheckResourceResponseBody) *EnableCheckResourceResponse {
	s.Body = v
	return s
}

type GetBdrcServiceResponseBody struct {
	Data *GetBdrcServiceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 5748C531-80B1-5C31-8421-63A1830B9E48
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetBdrcServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetBdrcServiceResponseBody) GoString() string {
	return s.String()
}

func (s *GetBdrcServiceResponseBody) SetData(v *GetBdrcServiceResponseBodyData) *GetBdrcServiceResponseBody {
	s.Data = v
	return s
}

func (s *GetBdrcServiceResponseBody) SetRequestId(v string) *GetBdrcServiceResponseBody {
	s.RequestId = &v
	return s
}

type GetBdrcServiceResponseBodyData struct {
	// example:
	//
	// 1726169608
	OpenTime *int64 `json:"OpenTime,omitempty" xml:"OpenTime,omitempty"`
	// example:
	//
	// 1726169608
	ProtectionScoreUpdatedTime *int64 `json:"ProtectionScoreUpdatedTime,omitempty" xml:"ProtectionScoreUpdatedTime,omitempty"`
	// example:
	//
	// SUCCESS
	ServiceInitializeStatus *string `json:"ServiceInitializeStatus,omitempty" xml:"ServiceInitializeStatus,omitempty"`
	// example:
	//
	// OPENED
	ServiceStatus *string `json:"ServiceStatus,omitempty" xml:"ServiceStatus,omitempty"`
}

func (s GetBdrcServiceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetBdrcServiceResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetBdrcServiceResponseBodyData) SetOpenTime(v int64) *GetBdrcServiceResponseBodyData {
	s.OpenTime = &v
	return s
}

func (s *GetBdrcServiceResponseBodyData) SetProtectionScoreUpdatedTime(v int64) *GetBdrcServiceResponseBodyData {
	s.ProtectionScoreUpdatedTime = &v
	return s
}

func (s *GetBdrcServiceResponseBodyData) SetServiceInitializeStatus(v string) *GetBdrcServiceResponseBodyData {
	s.ServiceInitializeStatus = &v
	return s
}

func (s *GetBdrcServiceResponseBodyData) SetServiceStatus(v string) *GetBdrcServiceResponseBodyData {
	s.ServiceStatus = &v
	return s
}

type GetBdrcServiceResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetBdrcServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetBdrcServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s GetBdrcServiceResponse) GoString() string {
	return s.String()
}

func (s *GetBdrcServiceResponse) SetHeaders(v map[string]*string) *GetBdrcServiceResponse {
	s.Headers = v
	return s
}

func (s *GetBdrcServiceResponse) SetStatusCode(v int32) *GetBdrcServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetBdrcServiceResponse) SetBody(v *GetBdrcServiceResponseBody) *GetBdrcServiceResponse {
	s.Body = v
	return s
}

type OpenBdrcServiceResponseBody struct {
	// example:
	//
	// 86DEBAC9-AB6A-59AB-9E5C-A540E579ECC9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s OpenBdrcServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s OpenBdrcServiceResponseBody) GoString() string {
	return s.String()
}

func (s *OpenBdrcServiceResponseBody) SetRequestId(v string) *OpenBdrcServiceResponseBody {
	s.RequestId = &v
	return s
}

type OpenBdrcServiceResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OpenBdrcServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OpenBdrcServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s OpenBdrcServiceResponse) GoString() string {
	return s.String()
}

func (s *OpenBdrcServiceResponse) SetHeaders(v map[string]*string) *OpenBdrcServiceResponse {
	s.Headers = v
	return s
}

func (s *OpenBdrcServiceResponse) SetStatusCode(v int32) *OpenBdrcServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *OpenBdrcServiceResponse) SetBody(v *OpenBdrcServiceResponseBody) *OpenBdrcServiceResponse {
	s.Body = v
	return s
}

type UpdateResourcesRequest struct {
	// example:
	//
	// ACS::ECS::Instance
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s UpdateResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateResourcesRequest) GoString() string {
	return s.String()
}

func (s *UpdateResourcesRequest) SetResourceType(v string) *UpdateResourcesRequest {
	s.ResourceType = &v
	return s
}

type UpdateResourcesResponseBody struct {
	Data *UpdateResourcesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 5B2F09BF-CEBD-5A7E-AC01-E7F86169A5E5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateResourcesResponseBody) SetData(v *UpdateResourcesResponseBodyData) *UpdateResourcesResponseBody {
	s.Data = v
	return s
}

func (s *UpdateResourcesResponseBody) SetRequestId(v string) *UpdateResourcesResponseBody {
	s.RequestId = &v
	return s
}

type UpdateResourcesResponseBodyData struct {
	// example:
	//
	// t-bp1ewftyzmeg3bl4dtd2
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s UpdateResourcesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s UpdateResourcesResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateResourcesResponseBodyData) SetTaskId(v string) *UpdateResourcesResponseBodyData {
	s.TaskId = &v
	return s
}

type UpdateResourcesResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateResourcesResponse) GoString() string {
	return s.String()
}

func (s *UpdateResourcesResponse) SetHeaders(v map[string]*string) *UpdateResourcesResponse {
	s.Headers = v
	return s
}

func (s *UpdateResourcesResponse) SetStatusCode(v int32) *UpdateResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateResourcesResponse) SetBody(v *UpdateResourcesResponseBody) *UpdateResourcesResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("bdrc"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

// @param request - CheckRulesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckRulesResponse
func (client *Client) CheckRulesWithOptions(request *CheckRulesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CheckRulesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ResourceArn)) {
		body["ResourceArn"] = request.ResourceArn
	}

	if !tea.BoolValue(util.IsUnset(request.RuleId)) {
		body["RuleId"] = request.RuleId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CheckRules"),
		Version:     tea.String("2023-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/rules/check"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CheckRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - CheckRulesRequest
//
// @return CheckRulesResponse
func (client *Client) CheckRules(request *CheckRulesRequest) (_result *CheckRulesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CheckRulesResponse{}
	_body, _err := client.CheckRulesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CloseBdrcServiceResponse
func (client *Client) CloseBdrcServiceWithOptions(headers map[string]*string, runtime *util.RuntimeOptions) (_result *CloseBdrcServiceResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("CloseBdrcService"),
		Version:     tea.String("2023-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/service/close"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CloseBdrcServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @return CloseBdrcServiceResponse
func (client *Client) CloseBdrcService() (_result *CloseBdrcServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CloseBdrcServiceResponse{}
	_body, _err := client.CloseBdrcServiceWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DescribeCheckDetailsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCheckDetailsResponse
func (client *Client) DescribeCheckDetailsWithOptions(request *DescribeCheckDetailsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeCheckDetailsResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.ResourceArn)) {
		query["ResourceArn"] = request.ResourceArn
	}

	if !tea.BoolValue(util.IsUnset(request.RuleId)) {
		query["RuleId"] = request.RuleId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeCheckDetails"),
		Version:     tea.String("2023-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/check-details"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeCheckDetailsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeCheckDetailsRequest
//
// @return DescribeCheckDetailsResponse
func (client *Client) DescribeCheckDetails(request *DescribeCheckDetailsRequest) (_result *DescribeCheckDetailsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeCheckDetailsResponse{}
	_body, _err := client.DescribeCheckDetailsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeProductsResponse
func (client *Client) DescribeProductsWithOptions(headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeProductsResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeProducts"),
		Version:     tea.String("2023-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/products"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeProductsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @return DescribeProductsResponse
func (client *Client) DescribeProducts() (_result *DescribeProductsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeProductsResponse{}
	_body, _err := client.DescribeProductsWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DescribeResourcesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeResourcesResponse
func (client *Client) DescribeResourcesWithOptions(request *DescribeResourcesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FailedRuleTemplate)) {
		query["FailedRuleTemplate"] = request.FailedRuleTemplate
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.SortBy)) {
		query["SortBy"] = request.SortBy
	}

	if !tea.BoolValue(util.IsUnset(request.SortOrder)) {
		query["SortOrder"] = request.SortOrder
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeResources"),
		Version:     tea.String("2023-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/resources"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeResourcesRequest
//
// @return DescribeResourcesResponse
func (client *Client) DescribeResources(request *DescribeResourcesRequest) (_result *DescribeResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeResourcesResponse{}
	_body, _err := client.DescribeResourcesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DescribeRulesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRulesResponse
func (client *Client) DescribeRulesWithOptions(request *DescribeRulesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeRulesResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeRules"),
		Version:     tea.String("2023-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/rules"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeRulesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeRulesRequest
//
// @return DescribeRulesResponse
func (client *Client) DescribeRules(request *DescribeRulesRequest) (_result *DescribeRulesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeRulesResponse{}
	_body, _err := client.DescribeRulesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeTaskResponse
func (client *Client) DescribeTaskWithOptions(TaskId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeTaskResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeTask"),
		Version:     tea.String("2023-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/tasks/" + tea.StringValue(openapiutil.GetEncodeParam(TaskId))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @return DescribeTaskResponse
func (client *Client) DescribeTask(TaskId *string) (_result *DescribeTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeTaskResponse{}
	_body, _err := client.DescribeTaskWithOptions(TaskId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DescribeTasksRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeTasksResponse
func (client *Client) DescribeTasksWithOptions(request *DescribeTasksRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeTasksResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.TaskStatus)) {
		query["TaskStatus"] = request.TaskStatus
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeTasks"),
		Version:     tea.String("2023-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/tasks"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeTasksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeTasksRequest
//
// @return DescribeTasksResponse
func (client *Client) DescribeTasks(request *DescribeTasksRequest) (_result *DescribeTasksResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeTasksResponse{}
	_body, _err := client.DescribeTasksWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DescribeTopRiskyResourcesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeTopRiskyResourcesResponse
func (client *Client) DescribeTopRiskyResourcesWithOptions(request *DescribeTopRiskyResourcesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeTopRiskyResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeTopRiskyResources"),
		Version:     tea.String("2023-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/resources/top-risky"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeTopRiskyResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeTopRiskyResourcesRequest
//
// @return DescribeTopRiskyResourcesResponse
func (client *Client) DescribeTopRiskyResources(request *DescribeTopRiskyResourcesRequest) (_result *DescribeTopRiskyResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeTopRiskyResourcesResponse{}
	_body, _err := client.DescribeTopRiskyResourcesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DisableCheckProductRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableCheckProductResponse
func (client *Client) DisableCheckProductWithOptions(request *DisableCheckProductRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DisableCheckProductResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ProductType)) {
		body["ProductType"] = request.ProductType
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DisableCheckProduct"),
		Version:     tea.String("2023-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/products/disable-check"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DisableCheckProductResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DisableCheckProductRequest
//
// @return DisableCheckProductResponse
func (client *Client) DisableCheckProduct(request *DisableCheckProductRequest) (_result *DisableCheckProductResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DisableCheckProductResponse{}
	_body, _err := client.DisableCheckProductWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DisableCheckResourceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableCheckResourceResponse
func (client *Client) DisableCheckResourceWithOptions(request *DisableCheckResourceRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DisableCheckResourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ResourceArn)) {
		body["ResourceArn"] = request.ResourceArn
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DisableCheckResource"),
		Version:     tea.String("2023-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/resources/disable-check"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DisableCheckResourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DisableCheckResourceRequest
//
// @return DisableCheckResourceResponse
func (client *Client) DisableCheckResource(request *DisableCheckResourceRequest) (_result *DisableCheckResourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DisableCheckResourceResponse{}
	_body, _err := client.DisableCheckResourceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - EnableCheckProductRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableCheckProductResponse
func (client *Client) EnableCheckProductWithOptions(request *EnableCheckProductRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *EnableCheckProductResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ProductType)) {
		body["ProductType"] = request.ProductType
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("EnableCheckProduct"),
		Version:     tea.String("2023-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/products/enable-check"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &EnableCheckProductResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - EnableCheckProductRequest
//
// @return EnableCheckProductResponse
func (client *Client) EnableCheckProduct(request *EnableCheckProductRequest) (_result *EnableCheckProductResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &EnableCheckProductResponse{}
	_body, _err := client.EnableCheckProductWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - EnableCheckResourceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableCheckResourceResponse
func (client *Client) EnableCheckResourceWithOptions(request *EnableCheckResourceRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *EnableCheckResourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ResourceArn)) {
		body["ResourceArn"] = request.ResourceArn
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("EnableCheckResource"),
		Version:     tea.String("2023-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/resources/enable-check"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &EnableCheckResourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - EnableCheckResourceRequest
//
// @return EnableCheckResourceResponse
func (client *Client) EnableCheckResource(request *EnableCheckResourceRequest) (_result *EnableCheckResourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &EnableCheckResourceResponse{}
	_body, _err := client.EnableCheckResourceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetBdrcServiceResponse
func (client *Client) GetBdrcServiceWithOptions(headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetBdrcServiceResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetBdrcService"),
		Version:     tea.String("2023-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/service"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetBdrcServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @return GetBdrcServiceResponse
func (client *Client) GetBdrcService() (_result *GetBdrcServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetBdrcServiceResponse{}
	_body, _err := client.GetBdrcServiceWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OpenBdrcServiceResponse
func (client *Client) OpenBdrcServiceWithOptions(headers map[string]*string, runtime *util.RuntimeOptions) (_result *OpenBdrcServiceResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("OpenBdrcService"),
		Version:     tea.String("2023-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/service/open"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &OpenBdrcServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @return OpenBdrcServiceResponse
func (client *Client) OpenBdrcService() (_result *OpenBdrcServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &OpenBdrcServiceResponse{}
	_body, _err := client.OpenBdrcServiceWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - UpdateResourcesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateResourcesResponse
func (client *Client) UpdateResourcesWithOptions(request *UpdateResourcesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		body["ResourceType"] = request.ResourceType
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateResources"),
		Version:     tea.String("2023-08-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/resources/update"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - UpdateResourcesRequest
//
// @return UpdateResourcesResponse
func (client *Client) UpdateResources(request *UpdateResourcesRequest) (_result *UpdateResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateResourcesResponse{}
	_body, _err := client.UpdateResourcesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
